# M221 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| 01 | tok/bootstrap | authored TOK-01 (fix-then-prove, host-isolation first) | N/A (tok) | continue |
| 02 | tik | Phase A: landed `GUARD-M221-host-isolation` — a per-N host lock (`hostlock.sh`) wired into `up-injected.sh` + `rosetta-demo down`/`up`; rext `cue-to-cue-m221-r1` | sub-gate MET: 4 DoD fences RED-proven (18 failed pre-fix → 18 passed) | closed-fixed |
| 03 | tik | Phase B: 3 off-box demo-hygiene fixes + fences — (1) academy binds `-H 127.0.0.1` on localhost (F-M220-5) + `exposure_claim_guard` extended to host-native listeners; (2) fenced the `ant-academy.sh` pre-bind reap (shipped M217+M220, ran in 0 tests); (3) new `backend_api_url_server_reader_guard` fences the F-7 server-side blackhole. rext `cue-to-cue-m221-r2` | sub-gate MET: 3 fences RED-proven pre-fix; suites 892+152 green; M220 S3 HARD INVARIANT undisturbed | closed-fixed |
| 04 | tik | Phase C cycle 1 — FIRST live battery on `billion` (DEFAULT `up-injected.sh 1`, NO FLAGS, cold reset-to-seed, driven from a tailnet peer). Baseline + diagnose. IRONCLAD-confirmed the dominant root cause (snapshot store-root bug: `public.skills 0→42,790` with the store pinned); re-characterised the academy carry (client-side render defect, not empty local-content); field-confirmed the native-academy teardown reap is still broken. No code fix landed (both candidates routed, de-risked). | **3/8 MET** (1 latency p95 2.23/2.08 s, 3 orgs, 7 remote-by-default) · 4 NOT MET (2,4,5,6 — ONE root cause: taxonomy cache-skip) · 1 at-risk (8) | closed-no-lift |
| 05 | tik | Phase C cycle 2 — landed F1/F5/F5b (RED-fenced, r3), the live cycle-1 FALSIFIED r3 on the box (skills=0, directus schema err — empty-subdir + detached-supervisor residuals), **corrected r3→r4 in-tik**, and the r4 re-prove ≈ MET the gate on ONE cold cycle. rext `cue-to-cue-m221-r4`. | ≈8/8 on one r4 cycle — PROVISIONAL (reproducibility + browser Dana grade remained) | closed-fixed |
| 06 | tik | **FINAL live demo cycle** — ONE DEFAULT cold no-flag `up-injected.sh 1` at r4 on `billion`, browser-graded from the peer (incl. Dana's `/ai-readiness` 900-char floor), **LEFT LIVE** (not torn down). F1 cascade fixed live (taxonomy 0→42,790, all 3 catalog surfaces replayed, directus self-unblocked). M219 readiness all MET; F10 fully field-exercised (assert_ports_free + freshness-abort); F4 confirmed known-gap + F-M221-06b (run-latency http-cockpit-scheme) routed forward. | **8/8 gate MET** (maya p95 2.11 s / dan 1.31 s; Dana browser >>900; Ben STARTED; Aria COMPLETED) — **Gate: MET** | closed-fixed |

## Next iter
_(set by the previous iter's closeout)_

- iter-05: **Phase C cycle 2 — land F1 + F2, then re-cycle.** (1) **F1** — pin `STACKSNAP_STORE`/`--store` to the
  workspace-level cache root in `dev-setdress.sh:323` (the walk-up resolves a shadowing empty
  `stack-demo/rosetta-extensions/.agentspace`), RED-fence the store-root divergence, re-tag rext. **De-risked:
  iter-04 proved the fix works** (`public.skills 0→42,790` with the store pinned). (2) **F2** — re-capture the
  directus surface (genuine digest divergence). Then a fresh cold no-flag battery on `billion` and re-measure
  gates 2/4/5/6 (F1 cascade-unblocks all four). Also route-forward this cycle: **F5** native-academy teardown reap
  (port+identity, not pidfile-only) + **F5b** gate-8 generated-file dirt; **F4** academy render defect; **F7** c-3
  router-403 re-check (with F2); **F10** freshness-abort + `assert_ports_free` field-exercise. Not yet reached:
  `BURNIN-M221-dev-public-host`, `F-M220-4`. Full baseline capture: `iter-04/findings.md`.

- iter-05 (tik): F1/F5/F5b landed; r3 cycle-1 falsified on the box → corrected r3→r4 in-tik → **r4 re-prove ≈ MET the gate on ONE cold cycle** (Maya p95 2.29s / Dan 1.65s, all 3 catalog surfaces, 3 orgs, Ben STARTED, Aria COMPLETED, no-flag) — PROVISIONAL (reproducibility remains) — see iter-05/progress.md

- iter-06 (tik) — **GATE MET; the milestone's exit gate is reached.** The FINAL DEFAULT cold no-flag r4 cycle on
  `billion` came up GREEN and **8/8** conditions held, browser-graded from the tailnet peer (maya p95 2.11 s /
  dan 1.31 s; Dana `/ai-readiness` filled >>900 ch; Ben STARTED; Aria COMPLETED); M219 readiness all MET; F10
  fully field-exercised. The stack is **LEFT LIVE** as the final demo (cockpit `https://billion.taildc510.ts.net:17700`,
  app `:13000`). **No further tik is required** for the gate. Residuals routed forward (Fate-3, next
  session/harden or close-milestone): **F-M221-06b** (`run-latency.sh` http-cockpit-scheme fix), **F4** (academy
  client-side grid render defect), and the earlier-carried `BURNIN-M221-dev-public-host` / `F-M220-4` /
  `PROBE-M218-c3-rerun`. Next lifecycle step: `/developer-kit:harden-mstone-iters --final` then
  `/developer-kit:close-milestone`.

## Gate Outcome Ledger

_close-milestone Phase 9-iter · iterative shape · **closed-on-gate** · 2026-07-15_

### Gate
- **Target:** on `billion.taildc510.ts.net`, a DEFAULT `/demo-up N` (**no flags**) yields, reproducibly on a cold
  reset-to-seed: (1) p95 click→ACCESS < 5 s for BOTH `maya-thriving` and `dan-manager` over the tailnet origin;
  (2) full replayed catalog — taxonomy + directus + sim-embeddings, no skipped surface; (3) all 3 story orgs incl.
  AI-readiness; (4) Dana sees a FILLED AI-readiness page; (5) Ben's from-scratch STARTED workflow visible; (6)
  Aria's COMPLETED state renders; (7) remote access came up BY DEFAULT, no flag; (8) ZERO platform-repo edits.
- **Achieved:** **8/8** on the iter-06 FINAL cold r4 cycle (a DEFAULT `up-injected.sh 1`, NO FLAGS), cold-proven
  (PG_VERSION mtime inside the container > T0) + green (`autoverify green:true, warnings:0`), browser-graded from
  a tailnet peer:
  1. **Login p95 — maya-thriving 2.11 s / dan-manager 1.31 s** (both < 5 s; ACCESS 5/5 each; tailnet HTTPS origin).
  2. Full catalog replayed — skills **42,790**, `directus_collections` **21**, sim-embeddings present (the F1
     store-root fix, live-confirmed; was 0 in iter-04).
  3. 3 orgs (Cervato / Solvantis / Northwind), cycles=2.
  4. **Dana `/ai-readiness` browser 900-char check PASSED** — all 3 tabs >> 900 chars (1,745 / 1,629 / 2,305).
  5. Ben STARTED on `/home`. 6. Aria COMPLETED (89/100). 7. remote-by-default auto-discovered, no flag. 8. zero
     platform edits.
  - **M219 readiness fold-in all MET** (0 junk of 132 distinct claimed skills; hero roles resolve; cycles=2;
    interview report present; frozen 62 / live 66 agree — the single-round score fix holds). **F10 field-exercised**
    (`assert_ports_free` + demopatch freshness G2-REFUSE abort). Seed isolation CLEAN (49 writes / 71,783 rows /
    prod=false).
- **Distance:** **gate met** (all 8 conditions hold).
- **Status:** `closed-on-gate`.
- **Reproducibility basis (disclosed):** the gate demands "reproducibly." Under the user's one-cycle pragmatic
  mandate, "reproducibly" is evidenced by **two independent cold reset-to-seed r4 cycles**: iter-05's r4 cycle
  (Maya p95 2.29 s / Dan 1.65 s, all 3 catalog surfaces, 3 orgs, Ben STARTED, Aria COMPLETED) **and** iter-06's
  FINAL cold cycle (the numbers above), both at rext r4, both DEFAULT no-flag, both green. Not a 5×-battery
  reproducibility (that would re-corrupt the single-host evidence the M219 lesson warns of and re-consume the
  left-live demo); two clean cold cycles at the same code is the honest, mandate-scoped bar. The stack is **LEFT
  LIVE** as the final deliverable (not torn down).

### Iter ledger summary
- **Total iters:** 6 — tiks: 5 (iter-02…06), toks: 1 (iter-01 bootstrap TOK-01).
- **Duration:** 2026-07-15 (single-day direct-drive, per the M215 live-infra analogue).
- **Decisions accumulated:** TOK-01 + per-iter D-M221-04b/05f/05g/06b/06d/06e (+ the reconciled orchestrator
  fallback note) + 2 close-time decisions (D-M221-C1/C2) + the Phase 2c adversarial subsection.
- **Hardening passes embedded:** the final `--final` harden pass (2 passes: continue → stabilized) —
  `hardening-ledger.md`. Per-iter regression fences were RED/mutation-proven inline as each fix landed.

### Routes carried forward — three-fate dispositions
Most in-milestone routes **resolved as Fate 1 in this milestone**: F1 (iter-05 + harden depth-2 edge), F5/F5b
(iter-05), F10 (iter-06 field-exercise), **F-M221-06b (this close's harden — `LATENCY_SCHEME`)**. The
five-milestone `dev-stack`-suite chronic was LAND-NOW'd at M220 and is confirmed running (harden: 118 passed).
Four **tail carries** remain — being at the FINAL milestone they cannot LAND-NEXT in-release:

#### Escape-hatch — cross-release deferral (sign-off owed at close-release)
- **F4** — academy grid renders 0 cards (catalog serves 2,705). Fix lives in the `ant-academy` **platform repo**;
  v2.3's zero-platform-edit constraint forbids it → **v2.4** (documented known cosmetic gap). D-M221-06e / D-M221-C2.
- **BURNIN-M221-dev-public-host** — live dev-path `--public-host` burn-in (non-gate; needs live cycling) → **v2.4**.
- **F-M220-4** — ant-academy re-runnability on a live public-host demo (non-gate) → **v2.4**.
- **PROBE-M218-c3-rerun** — router-403 re-check that needs the live box (non-gate) → **v2.4**.
- All four are recorded with fresh reasons in `audit-deferrals/deferral-audit-2026-07-15-m221-close.md` (YELLOW);
  their cross-release `RELEASE-SCOPE-DEFER` sign-off is `/developer-kit:close-release` Phase 1b's to obtain (it
  runs next). This close does not fabricate that sign-off.

### Dropped
- None.

### Protocol evolution
- The milestone's `iteration_protocol_ref` (`corpus/ops/verification.md` + `coverage-protocol.md` +
  `playthroughs.md`, the remote-origin cold reset-to-seed gates) held unchanged. M221's own contribution to the
  protocol lineage: the **host-isolation lock** (`hostlock.sh`) makes the single-host live battery
  self-non-corrupting (the M219 "two agents on one host corrupt the evidence" lesson, now fenced), and the D17
  "loud-not-silent / a status artifact read as evidence" hazard was exercised end-to-end (the snapshot-cache
  shadow, the M217 reap that ran in zero tests, the run-latency http scheme, and — a process instance — the
  premature iter-05 close on a stale r3 snapshot, reconciled loudly in `3c64af1`).

## M221: Final Review (close-milestone Phases 1–8)

### Scope (iterative)
- [x] Gate MET 8/8 (closed-on-gate); all 6 iters closed; every commit maps to an iter (iter-05 has the
  reconciliation commit `3c64af1`). Carry-forward queue audited → 4 non-gate tail carries to v2.4.

### Code Quality
- [x] `go vet` clean; shellcheck clean on all 7 M221-touched shell scripts; all touched `.py` compile.
- [x] Cross-cutting review done in the final-harden pass: consistent D17 "loud-not-silent / finding-not-pass"
  discipline; no dead code (the one dead path — a mid-file `unittest.main()` — was caught + fixed).

### Adversarial (Phase 2c)
- [x] 2 scenarios recorded in `decisions.md` (depth-2 F1 shadow; direct-run suite false-green) — both fixed in harden.

### Documentation
- [x] [fixed] `hostlock.sh` indexed in `demo-stack/README.md` (was missing though the other 2 M221 tools were listed).
- [x] [fixed] demo-stack test count reconciled 576 → 663 against the JUnit XML.
- [x] No new top-level rext unit needs a handbook (M221 added files within existing sections, not a new package).

### Tests & Benchmarks
- [x] Full suites green (JUnit-authoritative): demo-stack 663 (659 pass/4 skip), dev-stack 122 (118/4),
  stack-injection 260 (252/8), stack-core 182, stack-verify 114; Go all 13 stack-snapshot pkgs + 1831 funcs, 0 fail.
- [x] Flake gate 3/3 on the new Go edge test + the reordered `test_reap.py`.
- [x] Every M221 fence carries a RED/mutation regression test (audited in the hardening-ledger).

### Decision Triage
- [x] The durable mechanisms (host-lock, F1 store resolver, native/detached reap, host-native exposure guard,
  backend-url server-reader guard) are documented in `demo-stack/README.md` + the tools' own header docstrings —
  the rext section README IS their subsystem doc. No separate knowledge-doc blend needed beyond the README fixes.
- [x] Close-time decisions D-M221-C1 (F-M221-06b landed Fate 1) + D-M221-C2 (4 tail carries → v2.4) recorded.

### Deferral audit (Phase 1b)
- [x] `/developer-kit:audit-deferrals --scope=milestone` → **YELLOW**, 0 blocking. Report in `audit-deferrals/`.
