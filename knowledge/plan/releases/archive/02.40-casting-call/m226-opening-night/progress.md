# M226 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | Phase 0b KB-fidelity GREEN; billion recon (stale v2.3 demo up, prereqs green, C-6 mem risk, rext cutover needed); authored TOK-01 `reprove-hiring-on-billion` | 0/7 (no lift — tok) | closed-fixed — see iter-01/progress.md |
| iter-02 | tik | teardown stale v2.3 + substrate cutover (rext→m225-harden, 12 clones pinned to lock SHAs) + default cold `up-injected.sh 1` GREEN (~10.5 min, rc=0) + FIRST 7-condition measurement from this Mac | **3/7 GREEN (C4,6,7); C1 discrepancy 3+47 vs 5+45; C2/3/5 blocked by the :13001 serve gap** | closed-fixed — see iter-02/progress.md |
| iter-03 | tik | applied Finding-1 fix (consume `casting-call-m226-serve-hiring` + surgical serve re-apply → :13001 peer-reachable); measured C2/C3/C5 from this Mac | **6/7 GREEN — C2 (44×5, junk=0), C3 (Cara/Cody usable), C5 (recruiter p95 1.50 s <5 s); C4 fully confirmed. Only C1 (count) open** | closed-fixed — see iter-03/progress.md |
| iter-04 | tik | Finding-2 fix (`role_mix.admin 0.1→0.14`, tag `casting-call-m226-count-5-45`) + full DEFAULT cold re-bring-up at the fixed tag + re-measured all 7 | **7/7 GREEN on one default cold cycle (PROVISIONAL — C1 now 5+45; Finding-1 proven in default path). Findings 3 (orphan, self-resolved) + 4 (C2 harness race) surfaced** | closed-fixed — see iter-04/progress.md |
| iter-05 | tik | Finding-4 fix (C2 harness insights-capture poll, tag `casting-call-m226-c2-race-fix`) + 2nd clean DEFAULT cold cycle + re-measured all 7 (C2 reliable) + corpus fold-ins | **7/7 GREEN — 2nd cold cycle. GATE MET (2 clean cold cycles, both 7/7, 0 platform edits)** | closed-fixed (GATE MET) — see iter-05/progress.md |

## GATE MET — 2026-07-17

**The M226 7-condition exit gate is MET.** On `billion.taildc510.ts.net`, a **default `/demo-up 1` (no flags)**
yielded, **reproducibly across 2 clean cold reset-to-seed cycles** (iter-04b + iter-05), measured from this Mac
(the tailnet peer): (1) hiring org `is_hiring=true`, **exactly 5 admin + 45 candidate**; (2) recruiter comparison
**42 comparable candidates × each of the 5 positions** (≥40), junk=0, prod-ejects=0; (3) both candidate heroes
(Cara Completed / Cody Assigned) render usable profiles; (4) reads as hiring (nav "Results"); (5) recruiter **p95
click→ACCESS 1.09 s / 2.36 s < 5 s** (the 3rd measured vantage); (6) coexists with the 3 workforce orgs (12 heroes
/ 4 orgs on the cockpit); (7) **0 platform-repo edits**. rext code-of-record `casting-call-m226-c2-race-fix`
(4bd68ff). Four cross-machine findings surfaced + fixed (all tooling/harness/seed, 0 platform edits): F1 serve
gap, F2 count displacement, F3 surgical-orphan (self-resolved), F4 harness race.

**Next:** `/developer-kit:harden-mstone-iters` (final pass) → `/developer-kit:close-milestone`. Optional harden:
Finding-3 (pre-bind reap clears stale serve fronts on offset ports — a nice-to-have, not gate-blocking).

## M226: Final Review

_Close review (2026-07-17). Iterative milestone, `closed-on-gate`. Deterministic-only signal (go/python/tsc +
5/5 flake) — the live gate was already proven on `billion` (2 cold cycles) + independently orchestrator-re-verified;
billion was NOT re-brought-up (left UP as the live-proof artifact). **1 finding total, fixed.**_

### Scope
- [x] Gate-distance review → **7/7 MET, reproducible** (2 cold cycles + orchestrator re-verify). No gap.
- [x] Iter-ledger audit → all 5 iters closed (iter-01 tok + iter-02..05 tik); 6 iter/harden commits map 1:1; 0 orphans.
- [x] Finding-3 (pre-bind serve reap) → Fate 3, routed forward (non-gate-blocking, self-resolving) — see Gate Outcome Ledger + deferral audit.

### Code Quality
- [x] rext code harden-swept (3 deterministic regressions already fixed at harden; `TestTailscaleServe` port-count set corrected). No new cross-cutting findings. Deterministic suites GREEN.

### Documentation
- [x] `latency-budget.md` fold-in (R4 warm-up-transient + recruiter 3rd vantage + orchestrator re-verify) — accurate, cross-refs resolve.
- [x] `tailscale-serve.md` fold-in (apps/hiring :3001+off serve front, Finding-1) — accurate.

### Tests & Benchmarks
- [x] **[Phase 4 handbook reconciliation]** stack-seeding README test-func count drift **832→855** (13 pkgs, 83→86 files) — fixed in-place in the rext authoring clone (`7032aea`, tag `casting-call-m226-close`). Stale since the v2.1 roll; drifted across the whole v2.4 release.
- [x] Deterministic suites re-run GREEN (Go seeders+blueprint, Python injection 145/8, tsc exit 0); 5/5 flake gate on the two M226-touched deterministic suites.

### Decision Triage
- [x] Knowledge-worthy insights (Finding-1 serve front → `tailscale-serve.md`; R4 + recruiter vantage → `latency-budget.md`) already blended during iters/harden. Remaining decisions (billion ssh, substrate cutover, surgical-vs-full bring-up, harness race) = maintainer-archive → stay in `decisions.md`.

## M226: Gate Outcome Ledger (Phase 9-iter)

### Gate
- **Target:** On `billion.taildc510.ts.net`, a default `/demo-up N` (no flags) yields, reproducibly on a cold
  reset-to-seed: (1) hiring org present, `is_hiring=true`, exactly 5 managers + 45 candidates; (2) recruiter hero
  lands on the comparison surface, ≥40 comparable non-junk rows per each of the 5 positions; (3) 2 candidate heroes
  render usable assessed profiles; (4) org reads as hiring; (5) p95 click→ACCESS < 5 s for the recruiter vantage
  (3rd measured path); (6) coexists with the 3 workforce orgs; (7) 0 platform-repo edits.
- **Achieved:** **7/7** — C1 exactly 5 admin + 45 candidate · C2 42 comparable × each of 5 positions, junk=0,
  ejects=0 (orchestrator re-verify: 5/5 sims × 42) · C3 Cara/Cody usable profiles · C4 reads as hiring · C5 recruiter
  p95 **1.09 s / 2.36 s** (orchestrator re-verify **1.74 s**) < 5 s · C6 12 heroes / 4 orgs on the cockpit · C7 0
  platform edits.
- **Distance:** **gate met** (reproducibly — 2 clean default cold reset-to-seed cycles [iter-04b + iter-05], both
  7/7, measured from this Mac + independently orchestrator-re-verified).
- **Status:** `closed-on-gate`.

### Iter ledger summary
- **Total iters:** 5 (tiks: 4, toks: 1 bootstrap).
- **Duration:** 2026-07-17 → 2026-07-17 (single day).
- **Decisions accumulated:** TOK-01 (bootstrap strategy) + ~16 intra-iter decisions across iter-01..05.
- **Hardening passes embedded:** the final `harden-mstone-iters` pass (`hardening-ledger.md` Pass 1 — final,
  stabilized): the RED-proven net-count fence + 3 deterministic serve-regression fixes + the hiring serve fence.

### Routes carried forward — three-fate dispositions

#### Fate 3 — routed forward (non-gate-blocking)
- **Finding-3 — the pre-bind serve reap** (clear stale `tailscale serve` fronts on the demo's offset ports before
  bind; the M215 F12 out-of-band-serve window). **Why not Fate 1 at close:** a bring-up-path behavioral change on a
  **live-only** surface (no `tailscaled` in the build/test env → not deterministically testable here) needing a
  **live re-prove on billion**, forbidden at close (billion left UP; no re-bring-up). **Self-resolves in the default
  flow** (teardown already emits the reset). **Attached to:** a follow-up build-iter with a live re-prove, or the
  next `prove-on-<VM>` milestone (M226 is the final v2.4 milestone, so no later v2.4 milestone owns it — inherited by
  the v2.4 release close for routing). Recorded in `audit-deferrals/deferral-audit-2026-07-17-m226-close.md` (DEF-M226-01).

#### Escape-hatch — release-scope-breaking deferral
- **None.**

### Dropped (cut from goal entirely)
- **None.**

### Inherited carries (pre-existing, NOT M226 work — routed to v2.4 release close)
- 8 pre-existing demo-stack test failures (6× `test_cockpit.py` + `test_purge` + `test_reap`) — HEAD-identical, in
  files M226 never touched; predate v2.4. → v2.4 release close (Phase 1b) → future demo-stack test-debt harden pass.
- M204 `assign-and-track.UC1` assign-WRITE declared in-manifest TODO (reports `unimplemented`). → v2.4 release close.

### Protocol evolution
- The `iteration_protocol_ref` (`verification.md` + `coverage-protocol.md` + `latency-budget.md`) was extended by a
  **remote-origin cold-reset-to-seed acceptance run** with a **3rd measured vantage**: `latency-budget.md` gained the
  recruiter vantage + the **R4 warm-up-transient** section (the compare-drawer cold first render is a transient, NOT
  a gate violation — C2 gates on data-present-and-renders, C5 on login→ACCESS) + the **`RENDER_TEST_TIMEOUT_MS`
  cold/tailnet probe-budget rule** (a slow-but-correct cold render can't false-fail the C2 probe); `tailscale-serve.md`
  gained the **apps/hiring :3001+off serve front** (the recruiter-reachability prerequisite). The harness gained the
  `recruiter` vantage (`run-latency.sh`) + remote-capable candidate/render specs.

**Gate met: 7/7 ≥ 7/7, reproducibly. No carry-forward.md; clean close.** (Finding-3 is a non-gate-blocking
hardening nice-to-have routed forward, not a gate shortfall.)
