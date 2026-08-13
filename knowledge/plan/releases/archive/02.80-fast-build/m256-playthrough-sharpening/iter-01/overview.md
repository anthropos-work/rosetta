---
iter: 01
milestone: M256
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
opened: 2026-07-28
---

# M256 · iter-01 — bootstrap tok

**Type:** tok (bootstrap) · unconditional iter-01 of an `iterative` milestone (build-mstone-iters Phase 0
rule 1). Its job is to author the **first** strategy — there is no stalled strategy to revise.

## Inputs (no prior iters exist)

- `../overview.md` — scope, the gate as **re-cut 2026-07-28 by D-v28-12**, the two settled findings
  (parallel-lane premise FALSE, suite dominated by one LLM-bound test).
- `../../evidence/playthrough-map.md` §1–§7 — the pre-seeded, user-reviewed map (18 live Playthroughs by
  product × stream × proof depth; the 28-UC curated gap; the un-homed 12).
- `corpus/ops/demo/playthroughs.md` — the declared `iteration_protocol_ref` (declare → extend seed →
  page object + spec → run `--reset` → reconcile → triage → re-measure).
- `roadmap.md` § Active — v2.8, decisions **D-v28-3 / -4 / -5 / -9 / -12**.
- The code itself: `rosetta-extensions/playthroughs/`, `clerkenstein/clerk-frontend/`,
  `stack-verify/e2e/lib/cockpit-login.ts`, and the M201 curated corpus.

## Phase plan (bootstrap tok — a multi-step planned shape)

1. Phase 0b KB-fidelity audit (milestone-once gate) — **blocks** strategy authoring on its verdict.
2. Price each lever against the **re-cut** gate, from code, without a stack.
3. Extend `playthrough-map.md` into a **ranked triage** (§8) — do not re-derive §1–§7.
4. Stand up the **local demo stack** the whole milestone needs (infrastructure precondition; the
   baseline itself is iter-02's job).
5. Author **TOK-01** in the milestone-root `decisions.md`.

## What this tok must settle

- **Is the parallel-lane enabler on the critical path?** The overview says clause 1 is unreachable without
  it. Test that claim against the *re-cut* clause 1.
- **What actually moves a per-test median**, ranked, with `file:line` evidence.
- **Which cluster to land first** — the one that discharges the most gate clauses per unit of work.
- **Where the risk is** (the onboarding seed gap) and how to order around it.
- The **execution order** the first batch of tiks follows.

## Escalation conditions

- Audit verdict **RED** → `user-blocker`, do not author strategy against unverified knowledge.
- The local demo cannot be brought up → the milestone has **no measurement surface**; record it plainly
  and escalate rather than fabricate a baseline (the milestone's own instruction).

## Acceptable close-no-lift outcomes

A tok moves no metric by construction (Phase 3, tok branch). This iter closes `closed-fixed` when TOK-01
is authored, the ranked triage is landed, and the stack the tiks need is up — or `closed-fixed-partial` if
the stack does not come up and the strategy still lands.
