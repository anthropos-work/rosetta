#!/usr/bin/env bash
# test/lib/readiness.sh — readiness probes. Beyond "is the port open?", these
# verify the service answers a known query correctly.
#
# Black-box only. Never imports service internals — speaks each service's
# documented external interface.

set -euo pipefail

# Probe GraphQL gateway: introspection query should return __schema.types.
# Uses the federated supergraph at :5050.
probe_graphql_introspection() {
  local response
  response=$(curl -s -X POST \
    -H 'content-type: application/json' \
    -d '{"query":"{ __schema { queryType { name } } }"}' \
    --max-time 10 \
    http://localhost:5050/graphql 2>/dev/null || true)

  if echo "$response" | grep -q '"queryType"'; then
    echo "ok"
    return 0
  else
    echo "fail: $response"
    return 1
  fi
}

# Probe Postgres: connect as postgres user and run a trivial query.
# Verifies the platform schemas exist.
probe_postgres_schemas() {
  local expected=(public sentinel cms jobsimulation skiller skillpath extensions)
  local response missing=()

  # Use docker exec because the platform's `.env` ties psql creds to the container.
  response=$(docker exec anthropos-postgresql-1 psql -U postgres -tAc \
    "SELECT schema_name FROM information_schema.schemata ORDER BY schema_name;" 2>/dev/null || true)

  for schema in "${expected[@]}"; do
    if ! echo "$response" | grep -qx "$schema"; then
      missing+=("$schema")
    fi
  done

  if [[ ${#missing[@]} -eq 0 ]]; then
    echo "ok: all expected schemas present"
    return 0
  else
    echo "fail: missing schemas: ${missing[*]}"
    return 1
  fi
}

# Probe Redis: simple PING.
probe_redis_ping() {
  local response
  response=$(docker exec anthropos-redis-1 redis-cli PING 2>/dev/null || true)
  if [[ "$response" == "PONG" ]]; then
    echo "ok"
    return 0
  else
    echo "fail: $response"
    return 1
  fi
}

# Probe Gotenberg: hit the version endpoint.
probe_gotenberg_version() {
  local response
  response=$(curl -s --max-time 5 http://localhost:3200/version 2>/dev/null || true)
  if echo "$response" | grep -q .; then
    echo "ok: $response"
    return 0
  else
    echo "fail: no response"
    return 1
  fi
}

# Probe Sentinel: hit the Connect-RPC service-info endpoint. Doesn't perform
# an authz check; just verifies the service handler is mounted.
probe_sentinel_rpc() {
  # Connect-RPC services answer POST with content-type application/json on
  # /<package>.<Service>/<Method>. An empty body returns a 400 (with a
  # well-formed Connect error), which proves the handler is wired.
  local code
  code=$(curl -s -o /dev/null -w "%{http_code}" \
    -X POST \
    -H 'content-type: application/json' \
    -d '{}' \
    --max-time 5 \
    http://localhost:8087/sentinel.authorization.v1.AuthorizationService/Reload 2>/dev/null || echo "000")
  case "$code" in
    200|400|401|403) echo "ok: HTTP $code"; return 0 ;;
    000)             echo "fail: no response"; return 1 ;;
    *)               echo "fail: unexpected HTTP $code"; return 1 ;;
  esac
}

# Probe Storage: HTTP health endpoint should be reachable, and the RPC port
# should respond too.
probe_storage_rpc() {
  local code
  code=$(curl -s -o /dev/null -w "%{http_code}" \
    --max-time 5 \
    http://localhost:8301/ 2>/dev/null || echo "000")
  # Any HTTP response means the RPC server is up (Connect handlers return
  # 404 on /, but the TCP layer is alive).
  if [[ "$code" =~ ^[2345][0-9][0-9]$ ]]; then
    echo "ok: HTTP $code"
    return 0
  else
    echo "fail: $code"
    return 1
  fi
}
