#!/usr/bin/env bash
# test/reports/generate.sh — drive a /test-platform run and produce a
# single timestamped markdown report under .agentspace/test-platform/.
#
# Usage:
#   ./test/reports/generate.sh <scope>
# where scope ∈ {live, repos, census, full}
#
# Each scope:
#   live    — run test/live/verify.sh (liveness + readiness probes)
#   repos   — run test/repos/run.sh   (each repo's own test suite)
#   census  — run test/census/inventory.sh (test-file inventory)
#   full    — run all three sequentially
#
# Exit code is the worst exit code of the underlying scripts (1 if any fail,
# 2 if anything was skipped due to missing tools / missing checkout, else 0).
#
# Writes:
#   .agentspace/test-platform/op_YYYYMMDD_HHMMSS_<scope>.md
#   .agentspace/test-platform/op_YYYYMMDD_HHMMSS_<scope>.raw.txt  (raw stdout)

set -uo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROSETTA="$(cd "$HERE/../.." && pwd)"
OUTDIR="$ROSETTA/.agentspace/test-platform"
mkdir -p "$OUTDIR"

scope="${1:-live}"
case "$scope" in
  live|repos|census|full) ;;
  *) echo "✗ unknown scope: $scope (expected live|repos|census|full)" >&2; exit 2 ;;
esac

ts=$(date +"%Y%m%d_%H%M%S")
out_md="$OUTDIR/op_${ts}_${scope}.md"
out_raw="$OUTDIR/op_${ts}_${scope}.raw.txt"

# Tally worst exit code
worst=0
bump_worst() { local rc=$1; [[ $rc -gt $worst ]] && worst=$rc; }

# Collect outputs in temp files so we can interleave into one report.
live_out=$(mktemp)
repos_out=$(mktemp)
census_out=$(mktemp)
trap 'rm -f "$live_out" "$repos_out" "$census_out"' EXIT

run_live=false; run_repos=false; run_census=false
case "$scope" in
  live)   run_live=true ;;
  repos)  run_repos=true ;;
  census) run_census=true ;;
  full)   run_live=true; run_repos=true; run_census=true ;;
esac

if $run_live; then
  echo "═══ live probes ═══" >&2
  bash "$ROSETTA/test/live/verify.sh" > "$live_out" 2>>"$out_raw"
  bump_worst $?
fi

if $run_repos; then
  echo "═══ repo test suites ═══" >&2
  bash "$ROSETTA/test/repos/run.sh" > "$repos_out" 2>>"$out_raw"
  bump_worst $?
fi

if $run_census; then
  echo "═══ test census ═══" >&2
  bash "$ROSETTA/test/census/inventory.sh" > "$census_out" 2>>"$out_raw"
  bump_worst $?
fi

# ─────────────────────────────────────────────────────────────────────────────
# Render markdown report
# ─────────────────────────────────────────────────────────────────────────────

git_sha=$(cd "$ROSETTA" && git rev-parse --short HEAD 2>/dev/null || echo "unknown")
git_branch=$(cd "$ROSETTA" && git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown")
hostname_str=$(hostname)
overall_status=$([[ $worst -eq 0 ]] && echo "✓ pass" || ([[ $worst -eq 1 ]] && echo "✗ fail" || echo "⚠ partial"))

{
  echo "# Test Platform Report — $scope"
  echo
  echo "**Date**: $(date '+%Y-%m-%d %H:%M:%S %Z')  "
  echo "**Scope**: \`$scope\`  "
  echo "**Overall**: $overall_status (worst exit code: $worst)  "
  echo "**Host**: \`$hostname_str\`  "
  echo "**Rosetta git**: \`$git_branch\` @ \`$git_sha\`"
  echo
  echo "---"
  echo

  if $run_live; then
    echo "## Live verification"
    echo
    echo "Black-box probes against the running platform. Rosetta speaks each service's external interface only — never imports service internals."
    echo

    echo "### Liveness"
    echo
    echo "| Service | Status | Detail |"
    echo "|---|---|---|"
    grep '^liveness ' "$live_out" | while read -r _ name status detail; do
      icon=$([[ "$status" == "ok" ]] && echo "✓" || echo "✗")
      echo "| $name | $icon $status | $detail |"
    done
    echo

    echo "### Readiness"
    echo
    echo "| Probe | Status | Detail |"
    echo "|---|---|---|"
    grep '^readiness ' "$live_out" | while read -r _ name status rest; do
      icon=$([[ "$status" == "ok" ]] && echo "✓" || echo "✗")
      echo "| $name | $icon $status | ${rest} |"
    done
    echo
  fi

  if $run_repos; then
    echo "## Repo test suites"
    echo
    echo "Each platform repo's own tests are invoked via its native runner. Rosetta does not duplicate test logic — it captures pass/fail and a log location."
    echo
    echo "| Repo | Status | Duration | Detail |"
    echo "|---|---|---|---|"
    # repos/run.sh output: "<repo> <status> <duration_s> <detail...>"
    while read -r repo status dur rest; do
      [[ -z "$repo" || "$repo" == "#"* ]] && continue
      icon=$([[ "$status" == "pass" ]] && echo "✓" || ([[ "$status" == "skip" ]] && echo "○" || echo "✗"))
      echo "| $repo | $icon $status | ${dur}s | ${rest} |"
    done < "$repos_out"
    echo
    echo "_Per-repo logs in_ \`test/.cache/repo-<name>.log\` _(gitignored)._"
    echo
  fi

  if $run_census; then
    echo "## Test census (development-health)"
    echo
    echo "Read-only inventory of test files per repo. \`no-tests\` and \`no-ci\` are red flags worth investigating."
    echo
    echo "| Repo | Unit | Integ | E2E | CI | Flag |"
    echo "|---|---:|---:|---:|---:|---|"
    while read -r repo unit integ e2e ci flag; do
      [[ -z "$repo" || "$repo" == "#"* ]] && continue
      case "$flag" in
        ok)         icon="✓" ;;
        not-cloned) icon="—" ;;
        *)          icon="⚠" ;;
      esac
      echo "| $repo | $unit | $integ | $e2e | $ci | $icon $flag |"
    done < "$census_out"
    echo
  fi

  echo "---"
  echo
  echo "## Raw output"
  echo
  echo "Full stderr/stdout of every probe and runner: \`$(basename "$out_raw")\` (same directory)."
  echo
  echo "## Notes"
  echo
  if [[ $worst -eq 0 ]]; then
    echo "All probes passed. No action required."
  elif [[ $worst -eq 1 ]]; then
    echo "At least one probe or test suite failed. Failures above with ✗. Inspect raw output and per-repo logs."
  elif [[ $worst -eq 2 ]]; then
    echo "Probes ran without failure but some checks were skipped (missing tool, repo not cloned, etc.). Inspect ○ rows."
  fi
} > "$out_md"

echo "" >&2
echo "▶ report written to: $out_md" >&2
echo "▶ raw output:        $out_raw" >&2
echo "▶ overall: $overall_status" >&2

exit $worst
