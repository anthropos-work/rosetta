**Type:** tik (bring-up-fix) — under TOK-01. 0 production runtime code / 0 rext changes; corpus-only
pre-cold-run hardening. Protocol: triage → route fix to surface (corpus) → re-measure.

# iter-07 — tik progress

## Execution log
1. **Triage (Phase 0 re-survey).** Before spending a 20-45-min cold `/dev-up` (tik-11), closed the two
   KNOWN, still-open cold-`/dev-up` blockers surfaced in prior iters but not yet landed:
   - iter-04 routed-forward: the dev-up SKILL said "**12 containers**" (health check "confirm 12 healthy
     containers") — post-merge the `skiller` container is gone; warm `docker ps` = **11**. A cold `/dev-up`
     following the doc would fail its documented health assertion on a phantom 12th container.
   - M208 finding (overview.md Pre-surfaced requirement): `make reset-db` re-runs migrations against a wiped
     DB with **no `extensions` schema** → app/cms `vector(1536)` + gin-trigram migrations fail cold
     (`schema "extensions" does not exist`). The first-build flow creates extensions before `make migrate`,
     but the reset-db path (bundled auto-migrate) wasn't documented to re-bootstrap first.
2. **Verified container count = 11.** Warm `docker ps` = 11 `anthropos-*` (backend, graphql, jobsimulation,
   skillpath, cms, sentinel, storage, roadrunner, postgresql, gotenberg, redis) — no skiller. 11 is correct
   for the merged 4-subgraph graphql profile.
3. **Routed fix (corpus).**
   - `.claude/skills/dev-up/SKILL.md` (2 spots): "12 containers" → "11 containers post-merge" + explicit
     "skiller container gone, taxonomy merged into app's public schema" / "merged 4-subgraph platform (no
     skiller container)" annotations.
   - `corpus/ops/setup_guide.md` (Full Database Reset §): added a "⚠ Cold-reset ordering" block — reset-db's
     bundled auto-migrate fails cold without extensions; documented the re-create-schemas + reload-policy +
     re-`make migrate` recovery, tagged the M25-D9 class + why the first-build flow dodges it.
4. **Sweep.** Re-grepped bring-up docs (dev-up SKILL + setup_guide + run_guide) → 0 residual "12
   container/healthy" strings; run_guide never hardcoded a count (uses "all started containers Up").

## Re-measurement (gate sub-conditions)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (a) compose / no-skiller | MET (warm) | MET (warm) |
| (b) replay loads public.* | MET (warm) | MET (warm) |
| (c) seed closure green | MET (warm) | MET (warm) |
| (d) verify merged-assertion | MET (warm) | MET (warm) |
| (e) M42 coverage + Playthroughs | NOT MET | NOT MET (needs a live demo — tik-09/10) |
| (f) 0 residual skiller-schema refs | MET | MET |
**Metric:** fully-met sub-conditions **5/6 → 5/6** (no flip — this is prerequisite hardening). Lift: the
cold `/dev-up` proof (tik-11) is de-risked — two documented cold-failures removed (phantom-12th-container
health assertion + undocumented reset-db extensions race).

## Close — 2026-07-08

**Outcome:** Landed the dev cold-path corrections — dev-up SKILL container count 12→11 (skiller container
gone) + setup_guide reset-db "⚠ Cold-reset ordering" extensions-re-bootstrap block (M208 M25-D9 finding).
Corpus-only; 0 rext/platform edits. Metric 5/6 (no flip — pre-cold hardening); cold `/dev-up` de-risked.
**Type:** tik (bring-up-fix)
**Status:** closed-fixed
**Gate:** NOT MET (5/6 sub-conditions warm; (e) coverage/Playthroughs + the full COLD `/dev-up` + `/demo-up`
proofs remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1st tik of session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-07 D1 (dev-up SKILL 12→11 container count — skiller container removed by the merge), D2
(reset-db extensions re-bootstrap documented — M25-D9 class, corpus not rext: the platform Makefile is
un-editable + the main `/dev-up` drives docs, so the fix belongs in the doc the SKILL executes)
**Side-deliverables:** none.
**Routes carried forward:** OPTIONAL/out-of-critical-path — the rext `dev-stack` verify_svcs list (line 157)
still names `skiller`; only affects an additional `dev-N` verify scope (a spurious non-fatal `down` warning),
NOT the main `/dev-up` (N=0, uses `make up`) or `/demo-up` gate paths. Route to a future re-sync/harden pass
(needs a rext re-tag; not gate-blocking). tik-08 → cold `/demo-up` GREEN.
**Lessons:** A cold `/dev-up` health check is only as correct as its documented container count — the merge
silently invalidated the "12 containers" assertion. Pre-cold doc hardening (fixing known documented failures
BEFORE the expensive cold run) is cheap insurance against a mid-cold-run surprise that would burn a 20-45-min
bring-up.
