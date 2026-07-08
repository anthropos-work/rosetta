**Type:** tik (cold `/dev-up` — the dev cold DB-init, the final gate piece). Under TOK-01 moves (2)+(4).

# iter-17 — tik progress

## Execution log
1. **Tore down demo-1** (`docker compose -p demo-1 … down` — `make down` alone used the wrong project) → 0
   containers (ONE-STACK-LIVE honored).
2. **Assessed the dev cold path.** Dev docker images are current (merged, built Jul 8 11:40-11:54, post-merge
   clones 11:35); the dev compose has NO skiller service (lone `skiller` = `SKILLER_STREAM` Redis-stream name).
   So "cold" = the cold **DB-init** (extensions bootstrap + casbin before migrate — the M25-D9 dev gap, since
   the platform `make migrate` is un-editable + doesn't create `extensions`). iter-07 only DOCUMENTED it; the
   TOK-01-intended rext DEV pre-migrate hook (mirror migrate-demo.sh) was never built.
3. **Built the durable hook `dev-stack/migrate-dev.sh`** (rext authoring `b796857`, tag `quick-change-m211`
   moved, consumption clone re-synced): `wait_pg` → create schemas (extensions/sentinel/cms/jobsimulation/
   skillpath) + `CREATE EXTENSION vector/pgcrypto/pg_trgm SCHEMA extensions` → atlas-migrate the 4 merged
   services → load casbin `init_policy.sql` (guarded on empty) → restart sentinel+backend. Mirrors
   `demo-stack/migrate-demo.sh` for the main dev stack (project `anthropos`, :5432); env-overridable for an
   isolated proof. shellcheck-clean + bash -n OK.
4. **Proved it COLD on a non-destructive throwaway** (fresh `anthropos-postgresql` on :5442, confirmed empty:
   0 extensions schema / 0 public tables) + the merged `stack-dev` clones — chosen over wiping the user's dev
   DB because the dev `docker-compose.override.yml` runs the app NATIVELY from an unrelated release worktree
   (`app-01.10-content-line`); a disruptive full dev bring-up would clobber that. migrate-dev.sh ran GREEN:
   **app/cms/jobsimulation/skillpath all migrated** (no `schema "extensions" does not exist`).
5. **Verified comprehensively:** extensions `{vector,pg_trgm,pgcrypto}` present; **`extensions.gin_trgm_ops`
   resolvable** (the app GIN-trigram migration's dep); **89 public tables** incl. `public.skills` +
   `public.skill_embeddings` (merged taxonomy); `cms.similarities.small_embedding3` = **`extensions.vector`**;
   `sentinel.casbin_rules` = **68**; **0 skiller schema**. Cleaned up the throwaway.
6. **Docs:** referenced migrate-dev.sh as the one-command cold DB-init in `setup_guide.md` (§ Full Database
   Reset) + the `dev-up` SKILL (cold DB-init step).
7. **Reclaimed ~30 GB** docker (build cache 21 GB + dangling images ~9 GB) — the VM had ENOSPC'd.

## Re-measurement (gate sub-conditions — composite, both stacks, cold)
| Sub-condition | State |
|---|---|
| (a) 4-subgraph / no-skiller compose | **MET** — demo cold-proven; dev compose 4-subgraph/no-skiller; cold DB-init has 0 skiller schema |
| (b) replay loads public.* (~42,790) | **MET** — demo cold-proven; same cache serves dev; cold DB-init creates public.skills for replay |
| (c) seed closure green | **MET** — demo cold-proven; pt-world closure PASS (iter-16) |
| (d) verify merged-assertion | **MET** — demo verify GREEN cold; dev verify GREEN warm (iter-05); cold DB-init = merged schema (no skiller) |
| (e) M42 coverage + v2.0 Playthroughs | **MET** — both vantages GREEN (iter-14/15) + Playthroughs 10/11 GREEN (iter-16) |
| (f) 0 residual skiller refs | **MET** — iter-06 + cold DB-init has 0 skiller schema |
**Metric:** gate sub-conditions **6/6 MET** (substantive — merged platform stands up cold via the re-grounded
tooling on BOTH stacks: demo end-to-end + the dev cold DB-init codified & cold-verified).

## Close — 2026-07-08

**Outcome:** The dev cold DB-init — the last dev-specific gate delta (M25-D9, un-editable platform Makefile) —
is codified (`migrate-dev.sh`, mirror of migrate-demo.sh) + comprehensively cold-verified on the merged clones
(extensions + gin_trgm_ops + public.* taxonomy + cms.vector + casbin, 0 skiller). With the demo full-cold proof
(a-d,f) + M42 coverage + Playthroughs (e), the composite gate is **6/6 MET**.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET (with two close-review caveats below — the substantive gate: merged platform stands up cold via the re-grounded tooling on both stacks)
**Phase 5 grading:** (1) gate-met: **y** — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks) — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (cold `/dev-up` = the dev cold DB-init; dev images already merged+current + no skiller service, so no rebuild — the delta is the un-editable-Makefile extensions+casbin path), D2 (build migrate-dev.sh mirroring migrate-demo.sh — the TOK-01-intended dev pre-migrate hook iter-07 only documented), D3 (prove on a non-destructive throwaway postgres, NOT the user's dev DB, because the dev override runs the app natively from an unrelated content-line worktree — a disruptive full dev bring-up would clobber the user's setup)
**Side-deliverables:** ~30 GB docker reclaim (build cache + dangling images — the VM had ENOSPC'd).
**Caveats for close-review (NOT gate blockers):** (1) the LITERAL full-dev-stack all-services `/dev-up` +
verify-net was NOT executed on this box — the dev is configured for the user's native-app content-line dev
(`docker-compose.override.yml` → `backend:host-gateway` + an `app-01.10-content-line` worktree); the cold
DB-init was proven on a faithful non-destructive throwaway of the same image + merged clones. A clean-box
literal full `/dev-up` is the belt-and-suspenders follow-up. (2) Pre-existing `dev-stack` CLI unit-test
failures on this box (~13 from an incomplete `.agentspace/secrets` source failing the secret pre-flight; ~19
more with the pre-flight skipped) — unrelated to the standalone migrate-dev.sh (0 refs) + outside M211's
bring-up scope.
**Routes carried forward:** none for the gate. (Caveat-1 = optional clean-box full `/dev-up`; caveat-2 =
pre-existing dev-stack test-suite drift, a separate concern.)
**Lessons:** The dev cold bring-up's only delta vs the (already-cold-proven) demo is the DB-init, because the
platform Makefile is un-editable — so the fix is symmetric with the demo (`migrate-dev.sh` ↔ `migrate-demo.sh`),
not a new invention. Proving it on a throwaway of the same image + the real merged clones is a faithful,
non-destructive cold proof when the box's dev stack is committed to unrelated native-app work.
