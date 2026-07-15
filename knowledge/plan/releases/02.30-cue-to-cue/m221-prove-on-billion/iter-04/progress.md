# iter-04 (tik) — progress

**Type:** tik, under TOK-01 (Phase C, cycle 1). The FIRST live cross-machine battery cycle on `billion`: a
DEFAULT `up-injected.sh 1` (NO FLAGS), cold reset-to-seed, driven from a tailnet peer. Baseline + diagnose.
Full capture in `findings.md`.

## What happened
- **Pin roll:** clone m221-r2 but `rext.tag` still m220-r5 -> first bring-up fast-failed `UP_RC=1` (clean guard).
  Fixed `rext.tag`; second bring-up `UP_RC=0` (F0).
- **Cold proof:** PG_VERSION mtime inside the container `03:41:32Z` > T0 `03:35:22Z` -> genuinely cold.
- **Gate measured (from this Mac, tailnet peer):**
  - **1 MET** — maya p95 2.23 s, dan p95 2.08 s (ACCESS 5/5, hero identity present), over the tailnet.
  - **2 NOT MET** — all 3 snapshot surfaces skipped (cache-miss). Root cause F1 (taxonomy+sim-embeddings) + F2 (directus).
  - **3 MET** — 3 orgs (Cervato/Solvantis/Northwind), `ai_readiness_cycles=2`.
  - **4/5/6 NOT MET** — cascade of taxonomy=0 (F3): `ai-readiness-funnel=0`, `interview_aggregated_reports=0`, step-progress=0.
  - **7 MET** — auto-discovered billion (6 rungs, cert minted), no flag, `tailscale serve` fronting 6 ports HTTPS.
  - **8 NOT MET (at-risk)** — 12/13 clones clean; ant-academy dirtied by its own predev hooks (F5b).
- **Dominant root cause IRONCLAD-confirmed (F1):** `STACKSNAP_STORE` pinned -> taxonomy replay loaded 330,261
  rows, `public.skills 0 -> 42,790`. The bug is store-root resolution at `dev-setdress.sh:323`; the cache is valid.
- **Carries captured:** academy render-defect re-characterised (F4); native-academy teardown reap FAILED, killed
  manually (F5); F-7 measured (fast connection-refused, not the feared blackhole — F6); c-3 403s not reproduced
  this cycle (F7); C-6 RAM decided sufficient (F8); demopatch self-heal fired (F9); M217 items partial (F10).

## Left clean
Host pristine: nothing running, 0 containers, lock released, `tailscale serve` empty (F12 works), registry `{}`,
all clones 0-dirty, rext pinned m221-r2.

## Close — 2026-07-15

**Outcome:** First live baseline cycle complete. A DEFAULT no-flag `up-injected.sh 1` on `billion`, cold
reset-to-seed, driven from a tailnet peer. Baseline distance to the 8-condition gate: **3 MET (1,3,7) / 4 NOT MET
(2,4,5,6) / 1 at-risk (8)** — and the 4 NOT-MET conditions **collapse to ONE root cause**, the snapshot
store-root resolution bug (F1), IRONCLAD-confirmed (`public.skills 0 -> 42,790` with the store pinned) with an
exact seam. No code fix landed (both candidates routed with named handlers — the academy carry disproven, the
store-root fix de-risked for iter-05). Host left pristine, lock released.
**Type:** tik
**Status:** closed-no-lift (baseline-with-full-findings)
**Gate:** the 8-condition milestone gate — **MET: 1 (p95 2.23/2.08 s < 5 s, both heroes, tailnet), 3 (3 orgs incl.
AI-readiness), 7 (remote-by-default, no flag, serve fronting).** **NOT MET: 2 (all 3 catalog surfaces skipped —
F1 store-root bug for taxonomy+sim-embeddings, F2 genuine digest miss for directus), 4/5/6 (Dana/Ben/Aria — one
cascade of the taxonomy=0 from F1).** **AT-RISK: 8 (ant-academy clone dirtied by its own predev-hook generated
files — F5b).** Baseline established; gate not passed on cycle 1 (as expected).
**Phase 5 grading:** (1) gate-met: n (baseline cycle — 3/8 MET; a clean pass on cycle 1 was explicitly not
expected) — (2) triggered-tok: n (tik made measurable progress: the full baseline is captured, the dominant root
cause is IRONCLAD-diagnosed + de-risked, all carries triaged) — (3) re-scope: n (TOK-01 holds — Phase C is the
live battery, exactly as planned; iter-05 continues it) — (4) user-blocker: n — (5) cap-reached: n (3 tiks this
milestone) — (6) protocol-stop: n — **Outcome: continue** (iter-05 lands F1 + F2 and re-cycles).
**Decisions:** D-M221-04a..f (iter-04/decisions.md).
**Side-deliverables:** the snapshot store-root fix is de-risked to certainty for iter-05 (a known-good fix, not a
hypothesis); the academy carry is re-characterised (a client-side render defect, not empty local-content);
`FIX-M221-reap-native-academy` is field-confirmed still-broken; C-6 RAM is decided (sufficient); the run-latency
harness's http-cockpit gap + the pin-roll rext.tag gotcha are documented.
**Routes carried forward (Fate-3, iter-05+):** **F1** store-root fix at `dev-setdress.sh:323` (#1, de-risked) /
**F2** directus surface re-capture / **F4** academy client-side render defect (re-characterised) / **F5**
native-academy teardown reap (port+identity, not pidfile-only) + **F5b** gate-8 generated-file dirt / **F7** c-3
router-403 re-check (ride with F2) / **F9** optional demopatch baseline re-pin / **F10** freshness-abort +
`assert_ports_free` field-exercise. Not reached: `BURNIN-M221-dev-public-host`, `F-M220-4`.
**Lessons:** a "cache miss" is not always a missing cache — here the exact-hash entry existed and the *lookup
root* was wrong (a shadowing empty `.agentspace` in the consumption clone), which only ground truth
(`stacksnap status` at two roots + a pinned replay) could distinguish from a genuine miss. And a carry's stated
root cause is a hypothesis, not evidence: the academy "empty local-catalog" theory dissolved the moment the live
`/catalog.json` returned 2705 courses — the fix the carry proposed would have shipped against a phantom.

## Next iter
- iter-05: **land F1** (pin `STACKSNAP_STORE`/`--store` to the workspace cache root in `dev-setdress.sh`, RED-fence
  the store-root divergence, re-tag rext) **+ F2** (re-capture the directus surface), then **re-cycle** the live
  battery on `billion` and re-measure gates 2/4/5/6. The store-root fix is de-risked (the manual replay proved it).
