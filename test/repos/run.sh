#!/usr/bin/env bash
# test/repos/run.sh — invoke each platform repo's own test suite.
#
# Rosetta does NOT duplicate per-repo tests. This script shells into each
# cloned repo and runs the appropriate command for that stack, then captures
# the exit code + a short summary.
#
# Source of truth for which repos exist: anthropos-dev/platform/repos.yml
# (parsed at runtime — adding a repo there auto-includes it here).
#
# Output: machine-readable lines to stdout:
#   "<repo> <status> <duration_s> <detail>"
# Exit: 0 if every repo passed, 1 if any failed, 2 if any were skipped due
# to missing tools / missing checkout.
#
# Portability: targets bash 3.2 (macOS default).

set -uo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROSETTA="$(cd "$HERE/../.." && pwd)"
DEVDIR="$ROSETTA/anthropos-dev"
PLATFORM="$DEVDIR/platform"

if [ ! -d "$PLATFORM" ]; then
  echo "✗ anthropos-dev/platform not found at $PLATFORM" >&2
  echo "  run /setup-platform first" >&2
  exit 2
fi

REPOS_FILE="$PLATFORM/repos.yml"
if [ ! -f "$REPOS_FILE" ]; then
  echo "✗ $REPOS_FILE not found" >&2
  exit 2
fi

# Per-repo test command. Function instead of associative array (bash 3.2).
# Returns the command string for a repo name, or empty if none configured.
cmd_for_repo() {
  case "$1" in
    app)                 echo "go test ./..." ;;
    cms)                 echo "go test ./..." ;;
    jobsimulation)       echo "go test ./..." ;;
    skiller)             echo "go test ./..." ;;
    skillpath)           echo "go test ./..." ;;
    sentinel)            echo "go test ./..." ;;
    storage)             echo "go test ./..." ;;
    messenger)           echo "go test ./..." ;;
    roadrunner)          echo "go test ./..." ;;
    graphql-wundergraph) echo "echo 'no test suite declared'; exit 0" ;;
    next-web-app)        echo "pnpm -r run test --if-present" ;;
    studio-desk)         echo "npm test --if-present" ;;
    *)                   echo "" ;;
  esac
}

timeout_for_repo() {
  case "$1" in
    app)           echo 600 ;;
    skiller)       echo 600 ;;
    jobsimulation) echo 600 ;;
    next-web-app)  echo 900 ;;
    *)             echo 300 ;;
  esac
}

# Return non-empty reason string if we should skip this repo (missing toolchain).
skip_reason() {
  local repo="$1" cmd="$2"
  case "$cmd" in
    go*)   command -v go   >/dev/null 2>&1 || { echo "go not installed"; return; } ;;
    pnpm*) command -v pnpm >/dev/null 2>&1 || { echo "pnpm not installed"; return; } ;;
    npm*)  command -v npm  >/dev/null 2>&1 || { echo "npm not installed"; return; } ;;
  esac
  # No reason → empty
  echo ""
}

REPOS=$(awk '/^  - name:/{print $3}' "$REPOS_FILE")

total=0; passed=0; failed=0; skipped=0
fail_repos=""
skip_repos=""
status_overall=0

echo "$REPOS" | while IFS= read -r repo; do :; done   # no-op; iterate below

# We cannot rely on subshell variables in a piped `while`, so loop differently.
OLD_IFS=$IFS
IFS=$'\n'
for repo in $REPOS; do
  IFS=$OLD_IFS
  [ -z "$repo" ] && continue
  total=$((total + 1))
  cmd=$(cmd_for_repo "$repo")
  timeout_s=$(timeout_for_repo "$repo")

  if [ -z "$cmd" ]; then
    printf '%-25s skip  0      no test command configured\n' "$repo"
    skipped=$((skipped + 1))
    skip_repos="$skip_repos $repo(no-cmd)"
    status_overall=2
    continue
  fi

  repo_dir="$DEVDIR/$repo"
  if [ ! -d "$repo_dir" ]; then
    printf '%-25s skip  0      repo not cloned at %s\n' "$repo" "$repo_dir"
    skipped=$((skipped + 1))
    skip_repos="$skip_repos $repo(not-cloned)"
    status_overall=2
    continue
  fi

  reason=$(skip_reason "$repo" "$cmd")
  if [ -n "$reason" ]; then
    printf '%-25s skip  0      %s\n' "$repo" "$reason"
    skipped=$((skipped + 1))
    skip_repos="$skip_repos $repo($reason)"
    status_overall=2
    continue
  fi

  echo "▶ running tests in $repo: $cmd  (timeout ${timeout_s}s)" >&2
  start_ts=$(date +%s)
  log_file="$ROSETTA/test/.cache/repo-$repo.log"
  mkdir -p "$(dirname "$log_file")"

  set +e
  ( cd "$repo_dir" && timeout "$timeout_s" bash -c "$cmd" ) >"$log_file" 2>&1
  rc=$?
  set -e
  end_ts=$(date +%s)
  dur=$((end_ts - start_ts))

  if [ "$rc" -eq 0 ]; then
    printf '%-25s pass  %-4d   exit 0\n' "$repo" "$dur"
    passed=$((passed + 1))
  elif [ "$rc" -eq 124 ]; then
    printf '%-25s fail  %-4d   timeout after %ss (log: %s)\n' "$repo" "$dur" "$timeout_s" "$log_file"
    failed=$((failed + 1))
    fail_repos="$fail_repos $repo(timeout)"
    status_overall=1
  else
    tail_excerpt=$(tail -3 "$log_file" 2>/dev/null | tr '\n' ' ' | head -c 160)
    printf '%-25s fail  %-4d   exit %d (log: %s) … %s\n' "$repo" "$dur" "$rc" "$log_file" "$tail_excerpt"
    failed=$((failed + 1))
    fail_repos="$fail_repos $repo(exit$rc)"
    status_overall=1
  fi
  IFS=$'\n'
done
IFS=$OLD_IFS

echo "" >&2
echo "▶ repos: $total total, $passed pass, $failed fail, $skipped skip" >&2
[ -n "$fail_repos" ] && echo "  failures:$fail_repos" >&2
[ -n "$skip_repos" ] && echo "  skipped: $skip_repos" >&2

exit $status_overall
