# M208 — Progress

Section checklist (closure = all boxes land).

- [x] `make pull` stack-dev + stack-demo platform to origin/main (skiller gone) + app to v1.334 + siblings; capture before/after refs
- [x] Remove vestigial `stack-dev/skiller/` + `stack-demo/skiller/` clones
- [x] Rebuild images + re-migrate against merged `public` schema — cold `make up` rc=0; clean-slate `make reset-db`+`make migrate` creates the full `public` taxonomy from scratch (no skiller schema on a clean DB), given the `extensions` prerequisite (Finding 1)
- [x] Confirm 4-subgraph compose / no skiller container / `SKILLER_RPC_ADDR=http://backend:8083` — all confirmed; router serves the absorbed taxonomy subgraph (existing-volume run); backend needs `INVITATION_HMAC_SECRET` (Finding 2, routed)
- [x] Pin the merge fact-sheet (moved tables, `org_id IS NULL` predicate, measured 42,790 count, RPC, 4-subgraph list, `extensions` prerequisite) — `corpus/services/backend.md` § + banner; `corpus/services/skiller.md` stub banner
- [x] Opportunistic M25-D9 — **surfaced on clean-slate** (not Fate-1); routed **Fate-3 to M211** (extensions-bootstrap + PG-readiness) + M209 Risk-2 cross-ref
