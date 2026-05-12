#!/usr/bin/env bash
# test/live/verify.sh — black-box verification of a running platform.
#
# Probes each service's external interface (HTTP/RPC/GraphQL/DB) to confirm
# it is alive and serving. Never imports service internals.
#
# Output: machine-readable lines to stdout, summary to stderr.
#   stdout lines: "<phase> <service> <status> <detail>"
#                 phase    in {liveness, readiness}
#                 status   in {ok, fail, unknown}
# Exit: 0 if all probes ok, 1 otherwise.

set -euo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
LIB="$(cd "$HERE/../lib" && pwd)"

# shellcheck source=../lib/services.sh
source "$LIB/services.sh"
# shellcheck source=../lib/readiness.sh
source "$LIB/readiness.sh"

# ─────────────────────────────────────────────────────────────────────────────
# Phase 1: liveness
# ─────────────────────────────────────────────────────────────────────────────

echo "▶ liveness (is each service reachable?)" >&2

fail_count=0
while read -r name container hp kind target; do
  read -r _ status detail < <(probe_service "$name" "$container" "$hp" "$kind" "$target")
  if [[ "$status" == "up" ]]; then
    printf 'liveness %-16s ok    %s\n' "$name" "$detail"
    printf '  ✓ %-16s %s\n' "$name" "$detail" >&2
  else
    printf 'liveness %-16s fail  %s\n' "$name" "$detail"
    printf '  ✗ %-16s %s\n' "$name" "$detail" >&2
    fail_count=$((fail_count + 1))
  fi
done < <(service_rows)

echo "" >&2

# ─────────────────────────────────────────────────────────────────────────────
# Phase 2: readiness (deeper probes — only if liveness passed for that service)
# ─────────────────────────────────────────────────────────────────────────────

echo "▶ readiness (does each service answer correctly?)" >&2

run_readiness() {
  local label="$1" fn="$2"
  local out status
  if out=$($fn 2>&1); then
    status=ok
  else
    status=fail
    fail_count=$((fail_count + 1))
  fi
  printf 'readiness %-32s %s  %s\n' "$label" "$status" "$out"
  printf '  %s %-32s %s\n' "$([[ $status == ok ]] && echo '✓' || echo '✗')" "$label" "$out" >&2
}

run_readiness "postgres-schemas"     probe_postgres_schemas
run_readiness "redis-ping"           probe_redis_ping
run_readiness "graphql-introspection" probe_graphql_introspection
run_readiness "gotenberg-version"    probe_gotenberg_version
run_readiness "sentinel-rpc"         probe_sentinel_rpc
run_readiness "storage-rpc"          probe_storage_rpc

echo "" >&2
if [[ $fail_count -eq 0 ]]; then
  echo "✓ all live probes passed" >&2
  exit 0
else
  echo "✗ $fail_count probe(s) failed" >&2
  exit 1
fi
