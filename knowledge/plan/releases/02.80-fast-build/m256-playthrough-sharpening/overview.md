---
milestone_shape: iterative
milestone: M256
title: "playthrough sharpening"
status: planned
release: v2.8 "fast build"
exit_gate: "On billion, cold reset-to-seed: (1) FASTER — full live Playthrough suite p50 <= 120 s (from 228 s), 0 flake across 3 consecutive runs; (2) EFFECTIVE — every Playthrough passes a negative control (demonstrably RED when its outcome is absent) AND >= 5 mutating Playthroughs (from 1) AND >= 1 `blocked` outcome (from 0); (3) COVERED — onboarding (5 UCs) + org-admin (4 UCs) LANDED, and every remaining uncovered curated UC carries a written verdict (named future milestone or drop) — zero silent gaps. Plus D-v28-5: the cockpit logout / Back-to-Cockpit double-click defect FIXED (no Playthrough added)."
iteration_protocol_ref: corpus/ops/demo/playthroughs.md
re_scope_trigger: "If > 3 of the un-homed curated UCs prove `unimplementable-without-platform-edit`, escalate — that is a platform conversation, not a test one."
depends_on: [M255]
parallel_with: []
complexity: very-large
created: 2026-07-27
last_updated: 2026-07-27
---

# M256 — playthrough sharpening  (`iterative`)

**Status:** `planned` · **Shape:** `iterative` · **Complexity:** very-large · **Release:** v2.8 "fast build"
**Depends on:** M255

## Goal

Make the Playthrough suite a detector you can **trust** and **afford** — individually faster, actually proving
function rather than presence, and covering the journeys that are silently unwatched.

## The problem, stated precisely

The suite is **18/18 green** and the demo still has things that do not work. That is not a paradox; it is a
structural property, and it has five independent causes:

1. **Render ≠ function.** **1 of 18** Playthroughs proves a WRITE (`pt-assignment-assign`). The other 17 prove
   a page rendered populated — which, on a **seeded** demo, is the half that was never in doubt.
2. **No negative controls.** No Playthrough is proven to go RED when its outcome is absent. M219's lesson —
   *"a surface that renders is not the same as the RIGHT surface"* — is enforced in exactly one place
   (`LEGACY_AI_READINESS_URL`).
3. **Journeys stop at boundaries.** `pt-aisim-chat-launch` stops at `/start`. `pt-skillpath-legacy` stops at
   *"the completion control exists"*. `pt-studio-*` stops at *"the draft rendered"*. Defensible (P6), but it
   means a broken engine is invisible.
4. **One entitlement, one org shape.** Every actor is `entitlement: enterprise` on `pt-world`.
   **Outcome `blocked`: 0. Outcome `error`: 0** — nothing proves the platform correctly says *no*.
5. **Whole surfaces at zero.** ant-academy **0**, onboarding **0**, org-admin **0**, talk-to-data **0**.

## Shape (why iterative)

The coverage clusters are **unpriced until driven live**. `pt-world` seeds *post*-onboarding users, so an
onboarding Playthrough may need a seed state that does not exist. Org-admin is **four WRITE surfaces** that may
hit zero-edit walls. A fixed `In:` list would be speculative — the exit gate is the commitment.

## Bootstrap tok (iter-01) — already seeded

[`.agentspace/playthrough-map.md`](../../../../.agentspace/playthrough-map.md), compiled 2026-07-27 and
reviewed by the user: the 18 live Playthroughs by **product × stream × proof depth**, the 28-UC curated-corpus
gap, the **12 un-homed** use cases, and two speed levers visible from the config alone. iter-01 extends it into
a ranked triage + the first strategy — it does **not** re-derive the map.

## The coverage arithmetic (M201 curated corpus: 9 products / 28 UCs)

| | |
|---|---|
| Covered | **12** of 28 |
| Uncovered | **16** |
| — reserved by `M206` (vision) | 3 — `ai-sim.code`, `ai-sim.interview`, `profile.self-evaluation` |
| — reserved by `M207` (vision) | 1 — `skill-paths.academy` |
| — **un-homed (no milestone anywhere)** | **12** — onboarding ×5 · org-admin ×4 · `workforce.organization-feedback` · `profile-skills.import` · `talk-to-data.query` |

Per **D-v28-4**: **land onboarding (5) + org-admin (4)**; the other 3 plus the 5-release-old `M206`/`M207`
reservations each get a **written verdict** (named future milestone or drop). Zero silent gaps.

## Two levers already identified (not yet committed to)

- **The serial default is over-broad.** `workers: 1` / `fullyParallel: false` is pinned in
  `playwright.config.ts` because *"a Playthrough MUTATES real state against a single shared
  `organization_id`-scoped Postgres"* — but **17 of 18 mutate nothing**. A read-only lane at `workers: N` plus a
  serial mutating lane is safe **under the existing rationale**, with no seed partitioning.
- **Per-seat `storageState` reuse.** All 18 log in from scratch across **6 distinct seats**
  (`pt-employee`, `pt-manager`, `pt-ai-completed`, `pt-ai-started`, `pt-ai-manager`, `pt-recruiter`). Reuse pays
  the cockpit handshake ~6× instead of ~18×, with `pt-profile-identity` retained as the one test that proves
  the handshake itself.

Baseline: **3.8 min** (228 s) for 18 browser Playthroughs + ~99 unit specs; **~12.6 s per Playthrough**;
`retries: 0`; per-test timeout 120 s, `expect` 15 s. (Was 13 min before M254 iter-10's
`networkidle` → `domcontentloaded` anti-deadlock fix.)

## D-v28-5 — the cockpit logout defect

Logging out back to the cockpit **requires two-or-more clicks**. It is the same seat-switch machinery every
Playthrough drives (`hero-login.ts` / the M37 cockpit handshake), so it is **fixed here** — and, by the user's
explicit call, **deliberately gets no Playthrough**. (Should it later be wanted, it is a one-liner against the
existing page object.)

## Batch-gate rule (D-v28-3)

A run drives the **full batch to completion** — never halts at a step, never retries to hide a flake. At batch
end it emits **one consolidated red set**; a non-empty set **escalates to the user for renegotiation** (fix or
explicit written disposition). **Zero standing red**; nothing accumulates across runs.

## Open questions

- Does `pt-world` support a **pre-onboarding** user state at all?
- Do the org-admin writes have a **read-back surface**, or only a toast? (A write with no readable effect
  cannot be proven without a DB assert — which is a different, weaker proof shape.)
- Is a read-only parallel lane genuinely safe against the shared Directus/Redis surfaces, or only against
  Postgres?

## KB dependencies

`corpus/ops/demo/playthroughs.md` (the iteration protocol) · `corpus/ops/demo/coverage-protocol.md` ·
`corpus/ops/demo/cockpit-spec.md` · `corpus/ops/seeding-spec.md` · `corpus/ops/demo/stories-spec.md` ·
`knowledge/plan/spec-drafts/playthroughs/spec.md`

**Delivers → `corpus/ops/demo/playthroughs.md`** (the count, the streams, the negative-control contract, the
batch-gate rule)
