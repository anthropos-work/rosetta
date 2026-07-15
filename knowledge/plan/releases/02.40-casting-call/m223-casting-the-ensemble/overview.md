---
milestone: M223
slug: casting-the-ensemble
version: v2.4 "casting call"
milestone_shape: section
status: planned
created: 2026-07-15
depends_on: M222
delivers: updates to snapshot-spec.md (the directus.job_position replay surface) + seeding-spec.md/stories-spec.md (the 4th HIRING story + the is_hiring gate + the candidate-assessment funnel)
---

# M223 — Casting the ensemble

## Goal
The hiring org exists in the seed: **≤50 people (exactly 5 `admin` + 45 `candidate`, no `member`)**, distinct from
the 3 workforce orgs, with the **5 shared job positions** resolved to **real replayed content**. The 45 candidates
form a **realistic funnel** — MOST assessed on the 5 shared positions (so the comparison view is populated +
comparable), SOME in earlier states (assigned-not-taken) — with a **realistic non-degenerate score distribution**
(the M219 anti-junk / anti-flat-arc discipline), **NOT** a flat 225-session grid.

## Why section
A direct **M51 analog** — 1 preset org + narrative-gated seeders + reset/closure wiring + a content reader/replay —
all enumerable up front against the M222 contract. The research explicitly calls the seeder half section-shaped: 45
candidates × 5 shared sims is a direct generalization (2 shared → 5) of the `AIReadinessFunnelSeeder`.

## Scope

### In
1. A **4th `stories[]` entry** (`narrative: hiring`, `RoleMix` ≈ 0.1 admin / 0.9 candidate) with a hero-trio
   placeholder (heroes materialize as cockpit seats at M224).
2. A **`HiringConfigSeeder`** (the AI-readiness-config analog) defining the org's **5 shared HIRING sims**.
3. The **type-aware hiring-sim reader** — `type='SIMULATION_TYPE_HIRING' AND job_position IS NOT NULL` (the current
   `contentref` pool is type-blind: `SELECT id ORDER BY id LIMIT 50`; the `readAIReadinessSkillPool` precedent for a
   dedicated pattern query). Reserve the 5 shared sims **disjoint** in the reserved tail
   (`contentref.go reservedSimRefs/reservedAt`) so generic activity sessions can't collide/mis-score a candidate.
4. The **snapshot extension to REPLAY `directus.job_position` rows** (all 443 public) + guarantee the 5 chosen HIRING
   sims are captured — today `job_position` is structure + a synthesized public-read grant only ("no replayed rows,
   the expansion is simply NULL"). This makes the 5 positions **REAL replayed content** (D-DESIGN-2), pending M222's
   BA-6 confirm that the cold snapshot carries ≥5 usable HIRING sims.
5. **Wire new hiring rows into `resetTables`** + the closure gate (`datadna measure-closure`) + the isolation audit
   (`isolation.Guard`) — new hiring rows join `resetTables` + stay closure-measured across all orgs.
6. The **`candidate`-assessment funnel seeder** (the AI-readiness-funnel analog): resolve the org's 5 shared sim refs
   once, write each candidate a scored `SIMULATION_TYPE_HIRING` session against the shared sims — **MOST candidates
   on all 5, SOME assigned-not-taken** — each rolled into a **comparable-but-DIFFERENTIATED** per-candidate score (a
   realistic spread, not a flat arc). Reuse the M219 fix: skill ladder role→curated→general→**STOP, never pad**;
   every skill/role ref through the real resolvers (**closure green, never fabricated**).
7. Optionally seed `organization_sim_invitation_links` (faithful) — but the comparison reads **sessions, not links**,
   so this is a nice-to-have, not a gate.

### Out
- The render proof (M224) · cockpit heroes (M224) · coverage/playthrough (M225) · any latency work.

## Depends on
**M222** — the seeder-output contract (does the score need a per-session `validation_*`/eval row?) + the
real-vs-synth decision (BA-6) + the `is_hiring` gate + the `narrative: hiring` discriminator.

## KB dependencies
- `corpus/ops/demo/stories-spec.md` (the 7-table fan-out) · `corpus/ops/seeding-spec.md` ·
  `corpus/ops/snapshot-spec.md` · `corpus/services/ai-readiness.md` (the funnel contract) ·
  `corpus/services/hiring.md` (M222)

## Delivers → knowledge/corpus
Updates to `snapshot-spec.md` (the `job_position` replay surface) + `seeding-spec.md`/`stories-spec.md` (the 4th
story + the `is_hiring` gate + the candidate-assessment funnel).

## Demo-patch?
**Pure seeding** (+ a tooling-owned snapshot extension). No platform-render gap here — this is data. **If M222 found
the read-path needs a per-session `validation_*`/eval row, that extra row is added to the seeder (still data-only),
not a patch.**

## Risks carried
- **R3 (degrades-quality — the demo's whole point)** — 45 junk-or-identical assessments. Mitigation: the M219
  ladder + a realistic non-degenerate score DISTRIBUTION + closure green. (The believability manifest is M225's.)
- **R7 (degrades-quality)** — snapshot content starvation. If BA-6 came back thin, fall back to a dedicated
  hiring-sim pattern query or synthesized sims that **still resolve real skill/role refs** (closure preserved).
