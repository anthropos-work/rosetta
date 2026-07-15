# iter-05 (tik) — Phase C cycle 2: fix the dominant blocker + re-prove

**Type:** tik, under TOK-01 ("fix-then-prove, host-isolation FIRST").

## Close — 2026-07-15 (CORRECTED — supersedes the premature r3-only close)

> ⚠️ **This close was committed prematurely once (`766c029`) telling only the r3-falsification story, because the
> orchestrator graded the r3 cycle-1 and closed while the iter-05 agent was still working — correcting r3→r4 and
> re-proving. The r3 falsification below is TRUE for cycle-1, but the iter did not end there: the agent corrected
> the two live-caught residuals to r4 (D-05g) and the r4 re-prove ≈ MET the gate (D-05f). The real outcome is
> `closed-fixed`, not `closed-fixed-partial`.**

**Outcome:** Landed F1/F5/F5b (RED-fenced, r3), the live cycle-1 FALSIFIED r3 on the box (`skills=0`,
directus schema error — the empty-subdir + detached-supervisor residuals the unit fences couldn't model),
**corrected r3→r4 in-tik** (D-05g: F1 requires a NON-EMPTY snapshots dir; F5 reaps the supervisor by
port-anchored identity), and the **r4 re-prove ≈ MET the 8-condition gate on one cold cycle**: Gate 1 login
Maya p95 **2.29 s** / Dan **1.65 s** (both <5 s, over the tailnet); Gate 2 all 3 catalog surfaces replayed
(`skills 0→42,790`); Gate 3 three orgs; Gate 4 Dana filled (DB); Gate 5 Ben STARTED 1/3; Gate 6 Aria
COMPLETED 3/3; Gate 7 remote-by-default auto-discovered (no flag); Gate 8 zero platform edits. F5/F5b/F12
field-proven on the teardown. See D-05f for the full scoreboard.
**Type:** tik
**Status:** closed-fixed
**Gate:** ≈ MET on ONE cold r4 cycle (essentially 8/8) — **PROVISIONAL: the gate demands REPRODUCIBILITY**
(the M218/M219 lesson — one cold cycle is not a pass). A 2nd clean cold cycle + the browser-level 900-char
check on Dana's page remain.
**Phase 5 grading:** (1) gate-met: n (provisional only — reproducibility unproven) — (2) triggered-tok: n —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #4 of 5) — (6) protocol-stop: n —
Outcome: continue → iter-06 confirms reproducibility + finishes the browser grade.
**Decisions:** D-M221-05a..05h (iter-05/decisions.md). 05g = the r3→r4 correction; 05f = the r4 reprove grade.
**Side-deliverables:** F5 (port-anchored supervisor reap) + F5b (gate-8 dirt) + F12 (serve reset) — all
FIELD-proven on the r4 teardown (not just unit-fenced).
**Routes carried forward (Fate-3 → iter-06):**
- **Reproducibility:** one more clean cold `up-injected.sh 1` (no flags) confirming the same ≈8/8 — the gate's
  literal "reproducibly on a cold reset-to-seed."
- **The browser-level 900-char check on Dana's `/ai-readiness`** (gate 4 was DB-confirmed, not yet browser-graded).
- **F4** (academy client-side render — empty grid), **F10** (field-exercise freshness-abort + `assert_ports_free`),
  **BURNIN-M221-dev-public-host**, **F-M220-4**, **PROBE-M218-c3-rerun** — not reached; later tik.
**Lessons:** (1) "de-risked to certainty" via a manual pin proves the store is *loadable*, not that the *shipped
resolution finds it* — the live box caught the gap (empty-subdir shadow + detached supervisor) that no unit
fence modelled; the r4 correction closed it. (2) PROCESS: the orchestrator closed this iter prematurely on the
r3 story while the agent was still re-proving at r4 — a real instance of acting on a stale snapshot of a
concurrent worker's state. The host-isolation lock (iter-02) prevented *data* corruption, but not *bookkeeping*
corruption; the fix was to reconcile the record loudly (this correction), never silently.

## Ledger
- iter-05 (tik): landed F1/F5/F5b; r3 cycle-1 falsified on the box → corrected r3→r4 in-tik (empty-subdir +
  detached-supervisor) → **r4 re-prove ≈ MET the gate (Maya 2.29s / Dan 1.65s, all 3 surfaces, 3 orgs,
  Ben STARTED, Aria COMPLETED, no-flag) on ONE cold cycle — PROVISIONAL (reproducibility remains)** — see
  iter-05/progress.md
