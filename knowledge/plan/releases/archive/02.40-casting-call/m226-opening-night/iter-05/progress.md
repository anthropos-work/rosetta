**Type:** tik (under TOK-01). Fix Finding-4 (C2 harness race) + a 2nd clean default cold cycle → firm
reproducibility. Protocol: verification.md + coverage-protocol.md + latency-budget.md.

# iter-05 — work log

1. **Finding-4 fix (rext `stack-verify`):** `render-hiring-comparison.spec.ts` — after the DOM-rows poll, added a
   bounded 30 s poll for the `insightsByJobSimulations` LIST op to be CAPTURED before deriving the 5 sims. The DOM
   rows hydrate from SSR/RSC and settle before the client insights POST fires → deriving raced the capture
   (cold-stack ~3/5 runs lost → dom-links fallback empty → false "0 sims"). Non-fatal (SSR-only render falls
   through). **Validated live: C2 3/3 on the iter-04b demo** (was 3/5). Committed+tagged
   `casting-call-m226-c2-race-fix` (4bd68ff — the final code-of-record).
2. **Finding-3 self-heal confirmed:** consumed the final tag on billion; the teardown CLEARED :13001 automatically
   (the iter-04b default bring-up owned the reset plan) — 13001 free, no orphan. The default flow is clean.
3. **2nd clean DEFAULT cold cycle** (`down 1 --purge` → `up-injected.sh 1` NO FLAGS): clean :13001 bind,
   autoverify all-green, hiring org 5 positions + 284 sessions.
4. **Re-measured ALL 7 from this Mac** — C2 now reliably automated.

## The 7-condition scoreboard (iter-05 — 2nd default cold cycle, fixed harness)

| # | condition | verdict |
|---|-----------|---------|
| 1 | is_hiring + exactly 5 mgr / 45 cand | ✅ GREEN — 5 admin + 45 candidate (**reproducible**) |
| 2 | recruiter comparison ≥40 rows/each 5 positions | ✅ GREEN — 42×5, **2/2 first-try** (Finding-4 fix works on a FRESH stack) |
| 3 | 2 candidate heroes usable assessed profiles | ✅ GREEN — Cara (Completed) / Cody (Assigned) / Rae (no regression) |
| 4 | reads as hiring | ✅ GREEN — Rae nav "Results \| AI Simulations \| …" |
| 5 | recruiter p95 click→ACCESS < 5 s | ✅ GREEN — p95 2.36 s, ACCESS 5/5 |
| 6 | coexists with 3 workforce orgs | ✅ GREEN — 12 heroes / 4 orgs on cockpit |
| 7 | 0 platform-repo edits | ✅ GREEN — clones clean (cms `?? studio/` disclosed) |

**7/7 GREEN on the 2nd clean default cold cycle, C2 reliably automated.**

## GATE MET

**Two clean DEFAULT cold reset-to-seed cycles (iter-04b + iter-05), both 7/7 on `billion` over the tailnet, a bare
`/demo-up 1` (no flags), measured from this Mac (the peer). 0 platform-repo edits.** The M221 "one cycle is
provisional; a 2nd confirms reproducibly" precedent is satisfied, and the C2 measurement is now robust across a
warm demo (3/3) + two fresh stacks (iter-04b eventually + iter-05 2/2 first-try).

## Corpus deliverables (the milestone `Delivers → knowledge/corpus`)
- `corpus/ops/demo/latency-budget.md` — the recruiter 3rd vantage (gate + baseline: p95 1.09 s / 2.36 s over 2
  cycles) + the Finding-1 serve prerequisite.
- `corpus/ops/demo/tailscale-serve.md` — the apps/hiring port (`3001+off`) added to the serve front (Finding-1).

## Close — 2026-07-17

**Outcome:** Finding-4 fixed + validated (C2 reliably automated on a fresh stack); the 2nd clean default cold cycle
proved 7/7 GREEN. **The 7-condition exit gate is MET — reproducibly (2 cold cycles) on billion, default no-flags,
0 platform edits.** Corpus deliverables folded in.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: **y** — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks) — (6) protocol-stop: n — Outcome: exit-1 (gate-met)
**Decisions:** iter-05 D1 (Finding-4 poll-for-capture fix); D2 (final tag consumed on billion = code-of-record).
**Side-deliverables:** rext `casting-call-m226-c2-race-fix` (4bd68ff); corpus fold-ins (latency-budget.md + tailscale-serve.md).
**Routes carried forward:** Finding-3 hardening nice-to-have (pre-bind reap clears stale serve fronts on offset ports) — to harden/close-milestone; NOT gate-blocking.
**Lessons:** The billion proof surfaced four cross-machine breakages invisible on localhost (serve gap F1, count displacement F2, surgical-orphan F3, harness race F4) — three tooling/harness, one seed — exactly the M215/M221 pattern. Every fix was rext/tooling; 0 platform edits throughout. The gate held across 2 cold cycles with the measurement harness now robust.
