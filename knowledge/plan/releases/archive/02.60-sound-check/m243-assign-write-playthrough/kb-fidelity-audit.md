---
title: "KB Fidelity Audit — M243 assign-WRITE Playthrough"
date: 2026-07-22
scope: milestone:M243
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Playthroughs pillar (model, manifest, page-objects, seed, runner, 4-state map) | `corpus/ops/demo/playthroughs.md` | `rext playthroughs/` (manifest/ · e2e/ · seed/ · report/) | PAIRED |
| assign-WRITE UC1 (the sole TODO) | `playthroughs.md:105-107` + `assignment-monitoring.yaml` UC1 | `playthroughs/manifest/assignment-monitoring.yaml:34-50` | PAIRED (declared `TODO`) |
| pt-world dedicated seed + seed-worlds index | `playthroughs.md` §"The dedicated, decoupled seed" + `seeding-spec.md` | `playthroughs/seed/pt-world.seed.yaml`, `seed/seed-worlds.yaml` | PAIRED |
| per-surface page-object / locator layer | `playthroughs.md` §"The per-surface page-object / locator layer" | `playthroughs/e2e/lib/*.ts`, `url-shapes.ts` | PAIRED |
| `/enterprise/assignments` platform surface (the assign WRITE) | — (discovered live by M243, per P3 + iteration-protocol step 3) | `next-web-app` (platform, read-only) | CODE-ONLY (platform surface — not a corpus topic; investigated live, never pre-documented) |

## Fidelity Findings

1. **Playthroughs pillar — "15 live Playthroughs, 1 TODO" (`playthroughs.md:104-107`)**
   - Source: `corpus/ops/demo/playthroughs.md:104-107`
   - Expected: 15 live Playthroughs; 1 TODO = `assignment-monitoring.assign-and-track.UC1`.
   - Actual: 15 `@pt:` spec files, 15 distinct `@pt:` tags, 1 `playthrough: TODO` (assignment-monitoring.yaml:50, UC1), 16 total declared use cases.
   - Verdict: ALIGNED.

2. **assign-WRITE UC1 declared as a build-reference gap (`assignment-monitoring.yaml:34-50`)**
   - Source: `playthroughs/manifest/assignment-monitoring.yaml:34-50`
   - Expected (doc): the sole TODO is the assign-WRITE half, "a two-backend org-admin WRITE flow", out of M204's declared manager journeys.
   - Actual: UC1 declares `goal` (assign content with a deadline), `actor: pt-manager/enterprise`, `seed: pt-world`, `outcome: success`, a 5-step flow (open `/enterprise/assignments` → pick content → pick target → set deadline → confirm), an intermediate `assignment-builder-accepts`, a final "the assignment is created — appears in the manager's view AND as assigned for the target", `playthrough: TODO`.
   - Verdict: ALIGNED. The manifest already declares the UC per P5 (manifest-first); M243 builds the Playthrough that fills it.

3. **pt-world seed / seed-worlds single-sourcing + precondition-coverage (`seed-worlds.yaml`)**
   - Source: `playthroughs.md` §"The dedicated, decoupled seed" + §"A world's shape is DECLARED".
   - Expected: `ptvalidate` precondition-coverage resolves every UC's `seed.world`/`actor.hero`/`actor.entitlement`/`seed.preconditions[]` against `seed-worlds.yaml`, which is single-sourced with `pt-world.seed.yaml`.
   - Actual: `seed-worlds.yaml` lists `pt-world` with roster (incl. `pt-manager`), tiers (incl. `enterprise`), and a capabilities list. UC1 uses only `world: pt-world` + `hero: pt-manager` + `entitlement: enterprise` (all present) + no extra `preconditions[]`. Any NEW precondition M243 adds must land in lockstep in both files.
   - Verdict: ALIGNED.

4. **Cross-references** — `playthroughs.md` links to `coverage-protocol.md`, `stories-spec.md`, `seeding-spec.md`, `idempotency.md`, `clerkenstein.md`, `safety.md`, `ai_architecture.md`; all are indexed in CLAUDE.md and exist. Verdict: ALIGNED.

## Completeness Gaps

None blocking. The `/enterprise/assignments` platform surface is intentionally undocumented in the corpus — Playthroughs are built by investigating the live UI semantically (P3) and by the iteration protocol's step 3 ("add the page object for any new surface"). The METHOD contract (page-object discipline, locator rules, reset-to-seed, the honesty/anti-toothlessness thesis) is fully covered by `playthroughs.md`. The surface specifics are a build-time discovery, not a pre-req doc. The milestone will ADD documentation (the count bump 15→16 + any new page-object/precondition note) in Phase 5.

## Applied Fixes
None needed — all in-scope doc claims ALIGNED.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to build Phase 1. The anchor doc (`playthroughs.md`) is accurate to code; the assign-WRITE UC is manifest-declared and ready to be filled; the seed index is single-sourced. The platform surface is discovered live by design.
