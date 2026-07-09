---
iter: 17
milestone: M211
iteration_type: tik
status: planned
created: 2026-07-08
---

# iter-17 — tik: cold `/dev-up` GREEN (the DEV half of the cold headline)

## Active strategy reference
**TOK-01** move (2) "close the dev-side M25-D9 gap — mirror migrate-demo.sh as a rext DEV pre-migrate hook so a
cold `make reset-db`/`make migrate` on the main dev stack succeeds (the platform Makefile is un-editable)" +
move (4) "prove cold". This is the FINAL composite-gate piece.

## Step 0 — Re-survey
Gate is 5/6: (a) 4-subgraph/no-skiller, (b) taxonomy replay public.*, (c) seed closure, (d) verify
merged-assertion, (f) 0 skiller refs — all cold-proven on demo-1; (e) M42 coverage + Playthroughs COMPLETE
(iter-14/15/16). Remaining: **cold `/dev-up`**. The dev docker images are current (merged, built Jul 8
11:40-11:54, post-merge clones 11:35); the dev compose has NO skiller service (lone `skiller` ref =
`SKILLER_STREAM` Redis-stream name). So "cold" = the cold **DB-init** path (extensions bootstrap + casbin load
before migrate on a wiped DB) — the M25-D9 dev gap, since the platform `make migrate` is un-editable and does
NOT create the `extensions` schema/extensions itself.

## Cluster / target identified
iter-07 DOCUMENTED the cold-reset extensions ordering + iter-04 fixed the casbin load, but the TOK-01 intent —
a rext DEV pre-migrate hook mirroring `demo-stack/migrate-demo.sh` — was never built (iter-07 was corpus-only).
Build a minimal rext `dev-stack/migrate-dev.sh` (main dev stack, postgres :5432): `wait_pg` → create schemas
(extensions/sentinel/cms/jobsimulation/skillpath) + `CREATE EXTENSION vector/pgcrypto/pg_trgm SCHEMA extensions`
→ casbin `init_policy.sql` load (guarded on empty) → atlas migrate the 4 merged services (app:public / cms /
jobsimulation / skillpath). Then run a cold dev bring-up using it + verify GREEN.

## Hypothesis
On a wiped dev DB + the current merged images, the extensions-first + casbin + 4-service migrate sequence
(mirror migrate-demo.sh) makes a cold `make up` → migrate succeed (no `schema "extensions" does not exist`, no
`gin_trgm_ops` unresolved) + casbin_rules > 0 → 4-subgraph compose (no skiller container) + verify net GREEN.

## Expected lift
Cold `/dev-up` GREEN → gate 5/6 → **6/6 → GATE MET** (both /dev-up + /demo-up GREEN cold on the merged
platform). EXIT_REASON: gate-met.

## Phase plan
1. Tear down demo-1 (`cd stack-demo/platform && make down`) — ONE-STACK-LIVE.
2. Build `dev-stack/migrate-dev.sh` in the rext authoring copy (mirror migrate-demo.sh, main-dev-targeted) +
   a self-test; commit + re-tag `quick-change-m211`.
3. Cold dev bring-up: `make down` + wipe the DB volume (truly cold) → `make up` (merged images) →
   `migrate-dev.sh` (extensions + casbin + migrate) → verify.
4. Verify GREEN: migrate rc 0 (extensions resolved), casbin_rules > 0, 4-subgraph/no-skiller `docker ps`,
   the verification.md auto-verify net (backend + readiness merged-schema probe) GREEN.

## Escalation conditions
- If migrate fails cold for a reason NOT covered by extensions/casbin (a genuine merge defect) →
  `unimplementable-without-platform-edit` (the platform did the merge; v2.1 is tooling-only) — escalate.
- If the cold bring-up surfaces a multi-loop fix that can't land in one tik → route forward / cap.

## Acceptable close-no-lift outcomes
If the migrate-dev hook lands + the cold DB-init verifies GREEN but a downstream set-dress step needs a
separate fix, close-fixed-partial with the DB-init cold-proven + the residual routed.
