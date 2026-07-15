# M221 "prove-on-billion" — Retro

## Summary

**The FINAL milestone of v2.3 "cue to cue." Every release gate, re-proven live on a remote Linux VM
(`billion.taildc510.ts.net`), over the tailnet, with a DEFAULT `/demo-up` — no flags.** Gate **MET 8/8** on the
iter-06 cold reset-to-seed r4 cycle, browser-graded from a tailnet peer: login **p95 maya-thriving 2.11 s /
dan-manager 1.31 s** (both < 5 s, ACCESS 5/5); full catalog replayed (skills **42,790**, directus 21,
sim-embeddings); 3 orgs; **Dana `/ai-readiness` 900-char browser check PASSED** (1,745 / 1,629 / 2,305); Ben
STARTED; Aria COMPLETED (89/100); remote **default-on** no-flag; **0 platform-repo edits**. M219 readiness
fold-in all MET; F10 field-exercised; seed isolation CLEAN (49 writes / 71,783 rows / prod=false). **The `billion`
demo is LEFT LIVE** as an intentional final deliverable (cockpit `https://billion.taildc510.ts.net:17700`, app
`:13000`) — not torn down.

Direct-drive iterative (the M215 "prove-on-odyssey" analogue): live shared infra doesn't reward speculative
iteration. 6 iters in a day — 1 bootstrap tok + 5 tiks — with **host-isolation FIRST** (because M219 proved two
agents on one host corrupt the evidence).

**rext code-of-record:** `cue-to-cue-m221-final` (live-graded on `billion` at `-r4`). **Zero platform-repo edits.**

## Incidents this cycle

- **The live re-prove FALSIFIED the first fix (iter-05).** F1 shipped RED-fenced at r3; the live box then caught
  **two residuals no unit fence modelled** — an *empty* `snapshots/` subdir still shadowing the real cache
  (`public.skills` back to 0), and Next 16's **detached** `next-server` worker the reap's parent-walk couldn't
  reach (the academy respawned on a "freed" port). Corrected r3→r4 in-tik (non-empty-store resolver + `reap_procs`
  by port-anchored identity), re-fenced against the field-observed cases; the r4 re-prove ≈ met the gate. *The box
  did what a model could not.*
- **P-instance (bookkeeping, reconciled loudly):** the orchestrator prematurely committed an iter-05 close
  (`766c029`) telling the STALE r3-falsification story while the agent was still re-proving at r4. The host-lock
  prevented DATA corruption but not BOOKKEEPING corruption; corrected in `3c64af1`, never silently.
- **Harden false-green (P2, this close):** `test_reap.py` printed "Ran 21 tests … OK" on a direct `python3` run
  while `pytest` collected 41 — a mid-file `unittest.main()` silently omitted the 20 adversarial-error-path +
  suite-honesty classes below it (the exact fences guarding this milestone's reap work). Fixed inline (block → EOF;
  direct run 21→41, pytest unchanged).
- **No test regressions.** All suites green throughout; flake gate 3/3 on the new/edited targets.

## What went well

- **Host-isolation FIRST paid for itself.** The per-N host lock (`hostlock.sh`, 4 fences × mutants) landed in
  iter-02 and later protected iter-05's live cycle from exactly the concurrent-corruption M219 suffered.
- **Fix-then-prove kept the graded code = the shipped code.** Every off-box fence was RED/mutation-proven before
  the live battery, so the battery graded reality (the M219 graded≠shipped lesson, closed).
- **The D17 loud-store diagnostic FIRED live** and is what caught residual (i) — "the store is EMPTY … a nearer
  EMPTY .agentspace shadowing the real cache." A fence, not a comment.
- **The final harden caught its own theatre.** A false-green suite, a scheme-blind latency driver, and the F1
  "exists != populated" bug one directory deeper — all found and fixed/pinned before close.

## What didn't

- **F1 took two live cycles** — the shallow fix (require `snapshots/` exist) wasn't enough; only the box surfaced
  the empty-subdir depth. Unit fences modelled existence, not population. The harden then pinned a *third* depth
  (`snapshots/<surface>/`-empty) as loud-not-silent so the invariant is protected at every level.
- **F-M221-06b (`run-latency.sh` `http://` cockpit URL)** wasn't caught until the cockpit went HTTPS-fronted at
  M220 — a status artifact (the http scheme) that outlived the thing it described. Worked around by hand in
  iter-06, landed properly in this close's harden.

## Carried forward

Four **non-gate** tail carries → **v2.4** (cross-release; sign-off owed at `/developer-kit:close-release` Phase 1b):
- **F4** — academy grid renders 0 cards (catalog serves 2,705). Fix is a render-path change in the `ant-academy`
  **platform repo** → structurally out of v2.3's zero-platform-edit scope. Documented known cosmetic gap.
- **BURNIN-M221-dev-public-host** — the dev-path `--public-host` live burn-in.
- **F-M220-4** — ant-academy re-runnability on a live public-host demo.
- **PROBE-M218-c3-rerun** — router-403 re-check that needs the live box.

All recorded in `audit-deferrals/deferral-audit-2026-07-15-m221-close.md` (YELLOW; 0 blocking). The 5-milestone
`dev-stack`-suite chronic was LAND-NOW'd at M220 and is confirmed running (118 passed).

## Metrics delta (from metrics.json)

- **Gate:** 8/8 MET. Login p95 2.11 s / 1.31 s vs < 5 s (baseline 39.45 s / 38.30 s — a ~18× improvement).
- **Python tests:** 1341 (0 fail, 16 skip) — **+96 M221-attributable** (demo-stack +80, stack-injection +14,
  dev-stack +2). stack-core re-counted 152→182 (pre-existing drift *surfaced*, not M221's).
- **Go test funcs:** 1831 (+4 — the F1 store-resolver tests). 0 failures; `go vet` clean.
- **Flake:** 0 (gate 3/3 on new/edited targets). **Harden:** 2 passes → stabilized. **Platform-repo edits:** 0.
