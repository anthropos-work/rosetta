#!/usr/bin/env bash
# test/lib/services.sh — service registry for live probes.
#
# Single source of truth for what services rosetta knows about and where to
# reach them. Sourced by live/*.sh; do not execute directly.
#
# Each service is declared as a row in the SERVICES array. Format:
#   "<name>|<container>|<host_port>|<probe_kind>|<probe_target>"
#
# probe_kind values:
#   http        - simple GET; any 2xx/3xx/4xx response = up (not 5xx, not refused)
#   http-200    - GET must return 200
#   tcp         - raw TCP connect (no HTTP) — for redis, postgres
#   docker      - `docker compose ps` says container is healthy
#
# probe_target:
#   for http:   the path (e.g. /health). Host+port comes from host_port.
#   for tcp:    unused (host_port is enough)
#   for docker: unused

set -euo pipefail

# Registry. Edit this file to add/remove services.
# Source of truth: anthropos-dev/platform/docker-compose.yml graphql profile.
SERVICES=(
  # name              container                       host:port           probe_kind    probe_target
  "postgresql       | anthropos-postgresql-1        | localhost:5432    | docker      | -"
  "redis            | anthropos-redis-1             | localhost:6379    | docker      | -"
  "sentinel         | anthropos-sentinel-1          | localhost:8087    | http        | /"
  "backend          | anthropos-backend-1           | localhost:8082    | http        | /health"
  "skiller          | anthropos-skiller-1           | localhost:8085    | http        | /"
  "skillpath        | anthropos-skillpath-1         | localhost:8100    | http        | /"
  "jobsimulation    | anthropos-jobsimulation-1     | localhost:8400    | http        | /"
  "cms              | anthropos-cms-1               | localhost:8090    | http        | /"
  "storage          | anthropos-storage-1           | localhost:8300    | http        | /"
  "roadrunner       | anthropos-roadrunner-1        | localhost:10400   | http        | /"
  "graphql          | anthropos-graphql-1           | localhost:5050    | http-200    | /health"
  "gotenberg        | anthropos-gotenberg-1         | localhost:3200    | http-200    | /health"
)

# Returns each service as: name container hostport kind target
# (whitespace-separated, trimmed)
service_rows() {
  local row name container hp kind target
  for row in "${SERVICES[@]}"; do
    IFS='|' read -r name container hp kind target <<<"$row"
    # trim whitespace
    name=$(echo "$name" | xargs)
    container=$(echo "$container" | xargs)
    hp=$(echo "$hp" | xargs)
    kind=$(echo "$kind" | xargs)
    target=$(echo "$target" | xargs)
    echo "$name $container $hp $kind $target"
  done
}

# Probe a single service. Echoes one line:
#   <name> <status> <detail>
# where status is one of: up | down | unknown
probe_service() {
  local name="$1" container="$2" hp="$3" kind="$4" target="$5"
  local host="${hp%%:*}" port="${hp##*:}"
  local status detail

  case "$kind" in
    docker)
      if docker inspect -f '{{.State.Health.Status}}' "$container" 2>/dev/null | grep -q healthy; then
        status=up
        detail="docker health: healthy"
      elif docker inspect -f '{{.State.Running}}' "$container" 2>/dev/null | grep -q true; then
        status=up
        detail="container running (no healthcheck declared)"
      else
        status=down
        detail="container not running"
      fi
      ;;
    tcp)
      # bash /dev/tcp; redirect stderr so timeouts are quiet
      if timeout 3 bash -c "echo > /dev/tcp/$host/$port" 2>/dev/null; then
        status=up
        detail="tcp connect ok"
      else
        status=down
        detail="tcp connect failed"
      fi
      ;;
    http|http-200)
      local code
      code=$(curl -s -o /dev/null -w "%{http_code}" --max-time 5 "http://$hp$target" 2>/dev/null || echo "000")
      if [[ "$code" == "000" ]]; then
        status=down
        detail="connection refused / timeout"
      elif [[ "$kind" == "http-200" && "$code" != "200" ]]; then
        status=down
        detail="HTTP $code (expected 200)"
      elif [[ "$code" =~ ^[2345][0-9][0-9]$ ]]; then
        status=up
        detail="HTTP $code"
      else
        status=down
        detail="HTTP $code (unexpected)"
      fi
      ;;
    *)
      status=unknown
      detail="unknown probe kind '$kind'"
      ;;
  esac

  echo "$name $status $detail"
}
