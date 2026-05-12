#!/usr/bin/env bash
# test/census/inventory.sh — read-only inventory of each repo's test footprint.
#
# Does NOT execute anything. Walks each cloned repo and counts test files,
# checks for CI config, and surfaces "0 tests" as a development-health
# red flag.
#
# Output (machine-readable, one row per repo):
#   "<repo> <unit> <integration> <e2e> <ci_workflows> <flag>"
# where:
#   unit, integration, e2e:  counts of detected test files of each kind
#   ci_workflows:            count of YAML files under .github/workflows/
#   flag:                    "ok" | "no-tests" | "no-ci" | "no-tests,no-ci"
#
# Portability: targets bash 3.2 (macOS default) — no mapfile, no associative
# arrays, no readarray. Linux bash 5 also works.

set -uo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROSETTA="$(cd "$HERE/../.." && pwd)"
DEVDIR="$ROSETTA/anthropos-dev"
PLATFORM="$DEVDIR/platform"
REPOS_FILE="$PLATFORM/repos.yml"

if [ ! -f "$REPOS_FILE" ]; then
  echo "✗ $REPOS_FILE not found" >&2
  exit 2
fi

# Parse repos.yml into a plain newline-delimited list (bash 3.2 has no mapfile).
REPOS=$(awk '/^  - name:/{print $3}' "$REPOS_FILE")

count_unit_go() {
  find "$1" -name "*_test.go" -not -path "*/node_modules/*" -not -path "*/vendor/*" 2>/dev/null | awk 'END {print NR}'
}
count_unit_ts() {
  find "$1" \( -name "*.test.ts" -o -name "*.test.tsx" -o -name "*.spec.ts" -o -name "*.spec.tsx" \) \
    -not -path "*/node_modules/*" -not -path "*/.next/*" -not -path "*/dist/*" 2>/dev/null | awk 'END {print NR}'
}
count_unit_py() {
  find "$1" \( -name "test_*.py" -o -name "*_test.py" \) \
    -not -path "*/.venv/*" -not -path "*/__pycache__/*" 2>/dev/null | awk 'END {print NR}'
}
count_integration() {
  find "$1" -type d \( -name "integration" -o -name "integration_test" -o -name "integration-tests" \) \
    -not -path "*/node_modules/*" -not -path "*/vendor/*" 2>/dev/null \
    | while IFS= read -r dir; do
        find "$dir" -type f \( -name "*_test.go" -o -name "*.test.ts" -o -name "*.spec.ts" -o -name "test_*.py" \) 2>/dev/null \
          | awk 'END {print NR}'
      done | awk '{s+=$1} END {print s+0}'
}
count_e2e() {
  find "$1" -type d \( -name "e2e" -o -name "playwright" \) \
    -not -path "*/node_modules/*" -not -path "*/vendor/*" 2>/dev/null \
    | while IFS= read -r dir; do
        find "$dir" -type f \( -name "*.spec.ts" -o -name "*.test.ts" -o -name "*_test.go" \) 2>/dev/null \
          | awk 'END {print NR}'
      done | awk '{s+=$1} END {print s+0}'
}
count_ci_workflows() {
  if [ -d "$1/.github/workflows" ]; then
    find "$1/.github/workflows" -maxdepth 1 -type f \( -name "*.yml" -o -name "*.yaml" \) 2>/dev/null | awk 'END {print NR}'
  else
    echo 0
  fi
}

echo "# repo                       unit  integ  e2e   ci   flag"

echo "$REPOS" | while IFS= read -r repo; do
  [ -z "$repo" ] && continue
  dir="$DEVDIR/$repo"
  if [ ! -d "$dir" ]; then
    printf '%-26s  -     -      -     -    not-cloned\n' "$repo"
    continue
  fi

  unit_go=$(count_unit_go "$dir")
  unit_ts=$(count_unit_ts "$dir")
  unit_py=$(count_unit_py "$dir")
  unit=$((unit_go + unit_ts + unit_py))
  integ=$(count_integration "$dir")
  e2e=$(count_e2e "$dir")
  ci=$(count_ci_workflows "$dir")

  flags=""
  if [ "$unit" -eq 0 ] && [ "$integ" -eq 0 ] && [ "$e2e" -eq 0 ]; then
    flags="no-tests"
  fi
  if [ "$ci" -eq 0 ]; then
    [ -n "$flags" ] && flags="$flags,no-ci" || flags="no-ci"
  fi
  [ -z "$flags" ] && flags="ok"

  printf '%-26s  %-5s %-6s %-5s %-4s %s\n' "$repo" "$unit" "$integ" "$e2e" "$ci" "$flags"
done

exit 0
