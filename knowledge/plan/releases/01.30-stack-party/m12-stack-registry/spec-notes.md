# M12 — Spec notes

Technical notes accumulate here during build.

## Unified registry schema
_{type: dev|demo, N, ports, status, created} — one record per live stack._

## First-available-N allocator
_Scan registry + `docker ps`, return lowest free N; locked write; explicit-N override._

## Up/teardown wiring
_dev-stack + demo-stack consume the allocator; teardown frees the slot._
