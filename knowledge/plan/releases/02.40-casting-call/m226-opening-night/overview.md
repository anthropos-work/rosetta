---
milestone: M226
slug: opening-night
version: v2.4 "casting call"
milestone_shape: iterative
status: archived
created: 2026-07-15
last_updated: 2026-07-17
depends_on: M222, M223, M224, M225
exit_gate: "On billion.taildc510.ts.net, a default /demo-up N (no flags) yields, reproducibly on a cold reset-to-seed: (1) the hiring org present, is_hiring=true, exactly 5 managers + 45 candidates; (2) the recruiter hero lands on the comparison surface and sees ≥40 comparable non-junk rows per EACH of the 5 positions; (3) the 2 candidate heroes render usable assessed profiles; (4) the org reads as hiring; (5) p95 click→ACCESS < 5 s for the recruiter vantage (v2.3's gate extended to a 3rd measured path); (6) coexists with the 3 workforce orgs on the cockpit; (7) 0 platform-repo edits."
iteration_protocol_ref: corpus/ops/verification.md (+ corpus/ops/demo/coverage-protocol.md + corpus/ops/demo/latency-budget.md — a remote-origin cold-reset-to-seed acceptance run, the M215/M221 prove-on-billion lineage)
delivers: the live-proof record; any latency finding folds into corpus/ops/demo/latency-budget.md (a hiring 3rd vantage)
---

# M226 — Opening night

## Goal
Every requirement of this release verified **on the remote VM, over the tailnet, default `/demo-up`, no flags.**

## Exit gate (measurable — the 7-condition live billion proof)
On `billion.taildc510.ts.net`, a **default** `/demo-up N` (no flags) yields, **reproducibly on a cold
reset-to-seed**:
1. the hiring org present, **`is_hiring=true`**, **exactly 5 managers + 45 candidates**;
2. the recruiter hero lands on the comparison surface and sees **≥40 comparable non-junk rows per each of the 5
   positions**;
3. the 2 candidate heroes render **usable assessed profiles**;
4. the org **reads as hiring**;
5. **p95 click→ACCESS < 5 s** for the recruiter vantage (v2.3's gate extended to a **3rd** measured path);
6. **coexists with the 3 workforce orgs** on the cockpit;
7. **0 platform-repo edits.**

## Why iterative
The direct analogue is **M215 "prove-on-odyssey"** and **M221 "prove-on-billion"** — the last breakages only surface
on a live cross-machine run over the tailnet; the path is discovered, not enumerated.

## Iteration protocol
`corpus/ops/verification.md` + `corpus/ops/demo/coverage-protocol.md` + v2.3's `corpus/ops/demo/latency-budget.md` —
a remote-origin cold-reset-to-seed acceptance run.

## Scope

### In
- Bring a default `/demo-up` up on `billion` over the tailnet (no flags) and drive the 7-condition gate to green,
  reproducibly on a cold reset-to-seed.
- **Latency is GATED here** (condition 5) for the recruiter vantage — the 3rd measured access path after v2.3's
  employee + manager vantages.

### Out
- Nothing net-new; this is the acceptance closer. Any un-patchable surface **escalates**.

## Depends on
**M222, M223, M224, M225.**

## KB dependencies
- `corpus/ops/demo/tailscale-serve.md` · `corpus/ops/demo/latency-budget.md` · `corpus/ops/verification.md` ·
  `corpus/ops/safety.md`

## Delivers → knowledge/corpus
The live-proof record; any latency finding folds into `corpus/ops/demo/latency-budget.md` (a hiring 3rd vantage).

## Demo-patch?
Whatever M224 pinned, **re-proven live at final code** (the M221 `REPROVE-…-at-final-code` discipline); a live-only
perf gap may pin a perf demo-patch here. **A platform-repo edit is never in bounds.**

## Risks carried
- **R4 (blocks-milestone)** — whole-org-hydration latency on the 45×5 compare table. **Latency is gated here** (was
  reported at M224). If real, a sha-pinned perf demo-patch, re-proven live at final code.
- **R1 (re-surface)** — a render gate that passed on the dev box may re-appear on the cold remote run; the M215/M221
  lesson is that the last breakages are cross-machine.
