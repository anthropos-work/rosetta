---
milestone: M225
slug: dress-the-set
version: v2.4 "casting call"
milestone_shape: section
status: archived
created: 2026-07-15
depends_on: M223, M224
delivers: the hiring sections of corpus/ops/demo/coverage-protocol.md + corpus/ops/demo/playthroughs.md (the hiring-vantage coverage manifest + the hiring playthrough)
---

# M225 — Dress the set

## Goal
The hiring org comes up **auto-set-dressed** on a default `/demo-up`, passes a **hiring-vantage coverage gate**, and
has **one GREEN playthrough** proving the recruiter journey end-to-end.

## Why section
Enumerable — extend the auto-set-dress bring-up to replay the HIRING-sim content, author a hiring coverage manifest
+ a hiring playthrough, wire the org into `pt-world`. **Reuses the M42 coverage + M202 playthrough machinery (never
forked).**

## Scope

### In
1. **Fold the HIRING-sim (`SIMULATION_TYPE_HIRING`) capture + replay into the auto-set-dress pass** (the default
   `/demo-up` bring-up), so the hiring org's positions + content come up real with no manual steps. **NB (M222
   BA-6 / M223 D4 — reconciled at M225 KB-1): there is NO `directus.job_position` replay** — it captured 0 rows and
   the scoreboard doesn't read it; the 5 "positions" ARE 5 real captured `SIMULATION_TYPE_HIRING` sims, resolved by
   `readHiringSimPool` from the auto-set-dress directus replay.
2. A **hiring coverage manifest** wired into `manifestFor(vantage, expectedOrg, identityKey)` (org-conditional
   dispatch, the AI-readiness precedent) — asserting **candidate persona self-consistency** (role↔skills↔score) +
   the compare-surface sections + **0 prod-eject**.
3. A **`playthroughs/manifest/hiring.yaml`** use case (a recruiter compares candidates on a shared sim; optionally a
   candidate completes a hiring assessment) + the hiring org into the **decoupled `pt-world` seed** (test data ≠
   demo data). One GREEN playthrough proving the recruiter journey end-to-end.

### Out
- The live cross-machine proof (M226).

## Depends on
**M223** (frozen seed shape) + **M224** (a rendering surface to sweep). **Note:** the manifest *authoring* can begin
once M223 freezes the seed shape — a partial overlap with M224's render loop — but the coverage/playthrough **gate
cannot pass until M224 is green.**

## KB dependencies
- `corpus/ops/demo/coverage-protocol.md` · `corpus/ops/demo/playthroughs.md` · `corpus/ops/demo/frontend-tier.md` ·
  `corpus/ops/snapshot-spec.md`

## Delivers → knowledge/corpus
The hiring sections of `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/demo/playthroughs.md`.

## Demo-patch?
**Pure tooling** (manifests + seed + set-dress). No platform-render gap.

## Risks carried
- **R3 (degrades-quality)** — believability. The **hiring believability manifest** in this milestone's coverage
  sweep enforces persona self-consistency (candidate role↔skills↔score), the concrete anti-junk gate for the seeded
  distribution.
