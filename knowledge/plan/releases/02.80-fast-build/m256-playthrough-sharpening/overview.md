---
milestone_shape: iterative
milestone: M256
title: "playthrough sharpening"
status: planned
release: v2.8 "fast build"
exit_gate: "On billion, cold reset-to-seed, ALL MEASURED ON THE POST-COVERAGE SUITE (denominator pinned, D-v28-9): (1) FASTER — median per-Playthrough <= 5 s AND full post-coverage suite p50 <= 200 s wall-clock, 0 flake across 3 consecutive runs; the irreducibly LLM-bound studio lane is budgeted separately and excluded from the median (pt-studio-advanced-generate is a real ~2-3 min live-LLM round-trip, studio-builder.spec.ts:45 sets a 300 s timeout); negative-control runs are a SEPARATE execution and do not count toward the timed p50. (2) EFFECTIVE — every Playthrough passes a negative control (demonstrably RED when its outcome is absent) AND >= 5 mutating Playthroughs, where MUTATING means mutates state AND reads it back (the playthroughs.md:169-172 sense, which is what makes the 'from 1' baseline true) AND >= 1 `blocked` outcome (from 0). (3) COVERED — onboarding (5 UCs) + org-admin (4 UCs) LANDED, and every remaining uncovered curated UC carries a written verdict (named future milestone or drop) — zero silent gaps. Plus D-v28-5: the cockpit logout / Back-to-Cockpit double-click defect FIXED (no Playthrough added)."
iteration_protocol_ref: corpus/ops/demo/playthroughs.md
re_scope_trigger: "If > 3 of the un-homed curated UCs prove `unimplementable-without-platform-edit`, OR a negative control proves unimplementable for > 3 Playthroughs, escalate — that is a platform conversation, not a test one."
depends_on: [M255]
parallel_with: []
complexity: very-large
created: 2026-07-27
last_updated: 2026-07-27
---

# M256 — playthrough sharpening  (`iterative`)

**Status:** `planned` · **Shape:** `iterative` · **Complexity:** very-large · **Release:** v2.8 "fast build"
**Depends on:** M255

> **Revised 2026-07-27** after the adversarial plan review, which found the headline speed lever's premise
> **false twice over** (see the warning box below) and the speed clause **arithmetically impossible** as first
> written. The gate is re-cut against a pinned denominator and an LLM-bound floor (D-v28-9).

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

[`../evidence/playthrough-map.md`](../evidence/playthrough-map.md), compiled 2026-07-27 and
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

## ⚠️ The headline speed lever's premise is FALSE — corrected before iter-01 (review finding R1)

The first draft asserted: *"17 of 18 mutate nothing. A read-only lane at `workers: N` + a serial mutating lane
is safe **under the existing rationale**, with no seed partitioning."* **Both halves are wrong.**

**(a) The count is wrong.** `playthroughs/e2e/tests/skillpath-legacy.spec.ts:21-23` **self-declares**
*"MUTATION: Starting a path creates progress state → this Playthrough MUTATES real state, so the suite must run
under reset-to-seed"*, echoed independently at `playthroughs/e2e/lib/skill-path-page.ts:47-48`. The mechanism is
`getOrCreateSkillPathSession` — a **server-side create-on-read**. Three more specs are **unclassified**:
`studio-builder.spec.ts` ×2 (fires a real LLM generation, `:70-81`) and `aisim-chat-launch.spec.ts:61` (clicks
Start Simulation). Only **10 of 18** carry an explicit "(no mutation)" declaration. **The safe/unsafe partition
the lever consumes does not exist in any artifact.**
→ Correct statement: **17 of 18 are UNCLASSIFIED for mutation; ≥ 1 demonstrably mutates.**

**(b) More decisive — Postgres is not the binding shared surface.** Clerkenstein's fake FAPI holds **one global
active seat, one `signedIn` flag, one `sessID` per stack** (`clerkenstein/clerk-frontend/registry.go:67,75`;
`server.go:100-105`). Every Playthrough login runs `hero-login.ts:44-53` → `cockpit-login.ts:58-73` →
`POST /v1/demo/select` → `handleSelectIdentity` (`server.go:573-586`), which re-points the seat **and** sets
`s.signedIn = false; s.sessID = ""` **globally**. Under `workers: N`, worker 2's login signs worker 1's browser
out mid-journey and its `/v1/me` 401s — and `server.go:454-466` reads `activeUserLocked()` **with no cookie
input**, so **`storageState` reuse does not isolate it either**. Two in-repo comments already record this
verdict: `stack-verify/e2e/tests/m224-candidate-heroes.spec.ts:10-15` ("MUST run serial… a later `selectSeat`
clobbers an earlier session") and `content-stories.spec.ts:128-130`.

**(c) The existing rationale sanctions only two paths** — **stack-per-worker** or **per-worker seed
partitions** (`spec-drafts/playthroughs/spec.md:447-450`; `corpus/ops/demo/playthroughs.md:441-443`). The draft
invented a third and declared it safe.

**So parallelism is an ENABLER to be priced at iter-01, not a free lever.** Candidate enablers, both rext-owned
and zero-platform-edit: a **cookie/`__client`-scoped Clerkenstein registry**, or **one fake-FAPI per worker**.

**Also deliver a machine-checked per-spec `MUTATES` / `READ-ONLY` / `UNKNOWN` tag** (greppable, fenced by a
test) that the lane consumes instead of an assumed 17.

**`storageState` reuse remains real but small.** All 18 log in from scratch across **6 distinct seats**
(`pt-employee`, `pt-manager`, `pt-ai-completed`, `pt-ai-started`, `pt-ai-manager`, `pt-recruiter`); reuse pays
the handshake ~6× instead of ~18× — order **~30 s** at the `latency-budget.md` ACCESS cost — with
`pt-profile-identity` retained as the one test that proves the handshake itself. It does **not** close the gap
on its own.

## Baseline — and why the first draft's 120 s was impossible

**228 s** (3.8 min) for 18 browser Playthroughs + ~99 unit specs; `retries: 0`; default per-test timeout 120 s,
`expect` 15 s. (Was 13 min before M254 iter-10's `networkidle` → `domcontentloaded` anti-deadlock fix.)

**But the suite is dominated by one test.** Three specs **override** the default timeout —
`studio-builder.spec.ts:45` = **300 s**, `:91` = **180 s**, `assignment-assign.spec.ts:43` = **240 s** —
because `pt-studio-advanced-generate` is a real **~2–3 min live-LLM round-trip**. 228 s is consistent with
studio-advanced at **~120 s** plus 17 tests at **~4.5 s** each. So the "~12.6 s per Playthrough" average in
[`../evidence/playthrough-map.md`](../evidence/playthrough-map.md) §6 is an artifact of that one outlier, and a
flat **suite** wall-clock gate of 120 s was **arithmetically impossible before a line was written**.

Hence the re-cut gate (D-v28-9): a **median per-Playthrough** target, a **post-coverage suite** ceiling, and
the **LLM lane budgeted separately**.

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
- **What does the parallel-lane enabler cost?** (cookie-scoped Clerkenstein registry vs one fake-FAPI per
  worker — both rext-owned, zero platform edits). This is the blocking surface, **not** Directus/Redis, which
  is what the first draft asked about. Answer at iter-01: without it, gate clause 1 is unreachable.
- **How is a negative control produced** — without a platform edit, and without mutating the shared world?
  Do negative-control runs execute against a separate stack, a fixture, or a temporarily-emptied surface?
  (They are excluded from the timed p50 either way.)

## KB dependencies

`corpus/ops/demo/playthroughs.md` (the iteration protocol) · `corpus/ops/demo/coverage-protocol.md` ·
`corpus/ops/demo/cockpit-spec.md` · `corpus/ops/seeding-spec.md` · `corpus/ops/demo/stories-spec.md` ·
`knowledge/plan/spec-drafts/playthroughs/spec.md`

**Delivers → `corpus/ops/demo/playthroughs.md`** (the count, the streams, the negative-control contract, the
batch-gate rule)
