**Type:** tik — under TOK-01. Protocol: run bring-up phase (seed) → triage → route → re-measure.

# iter-03 — tik progress

## Execution log
1. **Build:** stackseed + datadna from authoring copy (quick-change-m209) → rc 0 each.
2. **Dry-run** stories-maya: `--stack ""` resolved to the preset's `demo-1` (port 15432) — the empty
   `--stack` does NOT override the blueprint stack. Re-ran with `--stack anthropos` → correctly resolves
   **N=0, port 5432, prod=false** (`blueprint.ParseStackN`: ""/"anthropos" → 0). The DAG + isolation
   preview validated cleanly against the merged schema: `public` schema hosts the merged taxonomy + pgvector
   (M209 comments), NO `skiller` schema, all writes `per-stack-isolated`, external stores `[BLOCK]`.
3. **Baseline:** warm tenant data all 0 (orgs/users/skills) — clean slate.
4. **Real seed** `stackseed --stack anthropos --seed presets/stories-maya.seed.yaml`: **45,710 rows, 45
   write attempts, isolation clean, prod=false.** Per-surface: org=1, taxonomy=42,790, users=150,
   personas=750, membership-skills=307, profiles=916, population-evidence=70, jobsim-sessions=105,
   assignments=65, activity=200, certificates=19, feedback=23, succession=17, etc. — the verified-skill
   fan-out landed.
5. **`datadna measure-closure --stack anthropos`** → **`[PASS] seed-verified-skill-closure: every seeded
   verified-skill node-id resolves in the replayed taxonomy`.** Sub-condition (c) core GREEN.

## Surfaced mid-iter (routed to iter-04)
2 seeders FAILED: `identity` + `users` → "casbin grant: resolve casbin table: pg: query row: no rows in
result set." **Root diagnosed:** the `sentinel` schema is entirely ABSENT from the warm DB (schemas
present: auth, cms, extensions, jobsimulation, public, skillpath). `sentinel.casbin_rules` never created —
`make migrate` was incomplete during the M208 container-only warm bring-up (sentinel container is UP +
listening :8087, but its policy schema never materialized). This is the **M18 silent-403 / sub-condition
(d) `casbin_rules > 0`** class, independent of the closure target. **Fix route (iter-04):** run
`make migrate` on the warm stack (chartered stack-op; idempotent per idempotency.md) → creates sentinel
schema + casbin_rules + loads policy → re-run the casbin-grant seeders → assert casbin_rules > 0. If migrate
reveals a merged-platform ordering defect (M25-D9 class), route the tooling hook to rext.

## Re-measurement (gate sub-conditions, warm)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (a) compose / no-skiller | MET | MET |
| (b) replay loads public.* | MET | MET |
| (c) seed closure green | NOT MET | **closure GREEN (PASS)**; full clean seed blocked on casbin (routed iter-04) |
| (d) verify merged-assertion | NOT MET | NOT MET — casbin_rules absent surfaced (iter-04 target) |
| (e) M42 coverage + Playthroughs | NOT MET | NOT MET |
| (f) 0 residual skiller refs | clean | seed DAG + isolation confirm public.* end-to-end |
**Metric:** closure GREEN proven (the milestone's seed-side merge-correctness gate). Fully-met sub-conditions
2/6; (c) closure-half green, seed-clean pending casbin (iter-04).

## Close — 2026-07-08

**Outcome:** `datadna` seed-side verified-skill closure GREEN — re-pointed public.* taxonomy resolvers
resolve real public node-ids (0 dangling) after a 45,710-row stories-maya seed. Surfaced + routed the
absent-casbin-policy bring-up gap to iter-04.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (closure sub-condition core proven; casbin/verify + coverage + playthroughs + cold proof remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (casbin is a routed-forward Fate-3 stack-op, not an architectural blocker) — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-03 D1 (closure GREEN validates the M209 seed-side re-ground), D2 (casbin gap = incomplete warm migrate, routed not folded)
**Side-deliverables:** none
**Routes carried forward:** casbin/sentinel-policy absent → **iter-04** (run `make migrate`, load policy, re-seed casbin-grants, assert casbin_rules>0 — sub-condition (d) start). Handler: `bringup-M211-iter04-casbin-migrate`.
**Lessons:** The re-grounded seeder resolves real public node-ids cleanly — the M209 skiller→public seed-side re-ground is proven correct against a live merged stack. A container-only warm bring-up (M208) leaves the sentinel policy unloaded; the seed's casbin-grant step is an early, cheap detector of that gap (same signal as the M18 verify cheap-win).
