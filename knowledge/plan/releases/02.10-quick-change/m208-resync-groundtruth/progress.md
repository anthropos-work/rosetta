# M208 — Progress

Section checklist (closure = all boxes land).

- [ ] `make pull` stack-dev + stack-demo platform to origin/main (skiller gone) + app to v1.334 + siblings; capture before/after refs
- [ ] Remove vestigial `stack-dev/skiller/` + `stack-demo/skiller/` clones
- [ ] Rebuild images + re-migrate against merged `public` schema
- [ ] Confirm 4-subgraph compose / no skiller container / `SKILLER_RPC_ADDR=http://backend:8083`
- [ ] Pin the merge fact-sheet (moved tables, `org_id IS NULL` predicate, ~42,763 count, RPC, subgraph list)
- [ ] Opportunistic M25-D9 dev migrate-ordering fix (non-blocking)
