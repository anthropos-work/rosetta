---
milestone: M201
slug: foundation
version: v2.0 "opening night"
milestone_shape: section
status: planned
created: 2026-06-28
last_updated: 2026-06-28
complexity: medium
delivers: corpus/ops/demo/playthroughs.md (NEW — graduates the playthroughs spec-draft into a corpus runbook: the capability, the manifest model, the page-object layer, the dedicated-seed + reset-to-seed lifecycle, the serial-default runner, the 4-state reporting map) + the new `playthroughs` rext section (the manifest + light validator + the per-surface locator/landmark page-object layer + the dedicated decoupled seed + the runner) built on the shared M42 e2e foundation
depends_on: none (reuses the M42 e2e harness + the seeding machinery)
spec_ref: knowledge/plan/spec-drafts/playthroughs/spec.md (the consolidated capability spec, v0.3)
---

# M201 — Playthroughs Foundation

## Goal
Stand up the **`playthroughs` rext section** on the **shared M42 e2e foundation**, proven by **one trivial
end-to-end Playthrough**. A Playthrough is an automated actor that **is the user** — it logs in as a seeded hero,
plays a real journey through the platform, and proves the platform delivered the outcome. M201 builds the
**plumbing** (the section, the manifest model + validator, the page-object layer, the dedicated seed + reset
lifecycle, the runner + reporting) — not the real product coverage (that is M202/M203). The proof of life is the
trivial Playthrough: **login → /profile → assert hero identity**.

## Scope
**In:**
- **(1) The manifest model + a light validator.** The manifest is the single source of truth (Products → Stories
  → Use Cases → Playthroughs, §2). The validator enforces:
  - **both-way id integrity** — every use case ↔ its Playthrough, traceable in **both** directions by a stable
    id (no orphan tests, no untested declared use cases mislabeled);
  - **precondition-coverage** — every declared `seed`/`preconditions` resolves to a **named seeded world** the
    Playthrough seed provides (no silent "ideally", §5.3);
  - **datadna-gated** — the Playthrough seed is covered by the **same `datadna` conformance gate** as the seeding
    machinery (§5.3).
- **(2) The per-surface locator/landmark page-object layer** — the §5.6 **shared registry every Playthrough
  imports** (semantic locators over the accessibility tree + a Rosetta-side landmark registry for ambiguous
  surfaces; a UI/antd/copy shift is absorbed by editing the per-surface registry, **not** N tests — re-pin is
  O(surfaces), not O(tests)). **1 surface to start.**
- **(3) The dedicated, decoupled seed preset** — test data ≠ demo data (§2/§5.4): the demo seed can be the
  *starting point*, but the Playthrough world is its own preset. **Spans entitlement tiers + multi-org-private.**
- **(4) The reset-to-seed lifecycle + serial-default runner** — per-suite **reset-to-seed** via the dedicated
  seed's own machinery (the real `--reset` path, honoring its contract + the N=0 `--force` guard; **additive
  re-seed is FORBIDDEN as a reset** — the M42e "green-but-wrong" trap, §5.7); the runner defaults to **serial**
  (`workers: 1`, `fullyParallel: false`, §5.7) against the single shared `organization_id`-scoped Postgres.
- **(5) The 4-state reporting map** (§5.5) — per use case: **`passing`** / **`failing`** / **`unimplemented`** /
  **`unimplementable-without-platform-edit`** (the last being the P3 zero-edit escape valve — it **escalates**,
  it does **not** edit the platform).
- **(6) One trivial proof Playthrough** — **login → /profile → assert hero identity** (the foundation smoke
  test; the cockpit-login handshake + the page-object layer + the seed + a single user-observable assertion).

**Out:** real product coverage (M202+); the AI-sim / integration mirror tier; cross-vantage.

## Why section (not iterative)
The foundation is **decomposable up front** — the manifest model, the validator, the page-object layer, the seed,
the reset lifecycle, the runner, the reporting map, and the one proof Playthrough are a known, enumerable
checklist. (The *coverage* of real journeys against the real antd UI — that is exploratory, and is M202/M203's
`iterative` job.) Build with `/developer-kit:build-milestone`.

## Depends on / Parallel with
- **Depends on:** **none** — it reuses the M42 e2e harness + the seeding machinery, both already shipped.
- **Parallel with:** none. M201 **lands first**; it gates the two `iterative` vantage-coverage milestones
  **M202 ∥ M203** (they import its page-object layer + run on its reset-to-seed lifecycle).

## Reuse paths (cite in spec-notes)
The shared e2e foundation it builds on (the §5.6 *built on a shared foundation it shares with `stack-verify`*):
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/cockpit-login.ts` — the existing cockpit-login helper (the
  M37 roster → fake-FAPI → `?__clerk_identity=` handshake); reused for hero login, **not** a generic
  Clerkenstein-login-anywhere (§5.6).
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/section-assert.ts` — the coverage sweep's per-section
  assertion helpers (the centralized-anchor discipline the page-object layer inherits).
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/empty-states.ts` — the empty-state/placeholder detection
  (so the proof Playthrough asserts real outcome state, not chrome).
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/coverage-manifest.ts` — the manifest-driven section model
  the playthroughs manifest extends (Products → Stories → Use Cases).
- `stack-demo/rosetta-extensions/stack-seeding/` — the seeding machinery the dedicated decoupled seed reuses
  (the `stackseed` / `--reset` path + the `datadna` gate).

## Re-scope trigger
A surface that can't be driven without a platform edit (the `unimplementable-without-platform-edit` state) →
**escalate, don't edit** (the P3 zero-edit line; the platform repos are read-only this release).

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) — the
  consolidated capability spec (v0.3) M201 builds the contract of (and graduates to a corpus runbook).
- [`corpus/ops/demo/coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) — the M42
  Playwright coverage sweep + the shared e2e foundation / locator discipline M201 extends from presence→function.
- [`corpus/ops/seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md) — the seeding machinery + the
  `stack.seed.yaml` blueprint + the production-isolation boundary the dedicated seed reuses.
- [`corpus/ops/idempotency.md`](../../../../../corpus/ops/idempotency.md) — the `--reset` contract + the N=0
  `--force` guard the reset-to-seed lifecycle honors (additive re-seed is NOT a reset).

## Delivers →
- **NEW** [`corpus/ops/demo/playthroughs.md`](../../../../../corpus/ops/demo/playthroughs.md) — the corpus runbook
  that **graduates the playthroughs spec-draft**: the capability, the manifest model + vocabulary, the per-surface
  page-object layer, the dedicated-seed + reset-to-seed lifecycle, the serial-default runner, and the 4-state
  reporting map. (Becomes the `iteration_protocol_ref` for M202/M203.)
- The **`playthroughs` rext section** — the manifest + light validator, the per-surface locator/landmark
  page-object layer (1 surface), the dedicated decoupled seed preset, the runner, and the one trivial proof
  Playthrough — authored + tagged per the tooling policy, consumed per-stack at a pinned tag.
