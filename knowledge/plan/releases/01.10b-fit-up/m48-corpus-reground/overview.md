---
milestone: M48
slug: corpus-reground
version: v1.10b "fit-up"
milestone_shape: section
status: planned
created: 2026-06-29
last_updated: 2026-06-29
complexity: medium-large
delivers: corpus/architecture/* + corpus/services/* re-ground to current prod; a documented member-AI-readiness contract
issues: corpus ~1-month staleness (user-flagged) + the M201 member-AI-readiness false-negative
---

# M48 — Corpus re-ground

## Goal
Bring the **rosetta documentation corpus current** with the freshly-synced prod (it lags ~a month): document
shipped-but-undocumented features the stale clones hid — above all the **member-AI-readiness flow** — reconcile
drift, and retire stale "RESOLVED" claims. The corpus must tell the truth before the seeding milestones build on it.

## Why section
The target file clusters are enumerable (`corpus/architecture/`, `corpus/services/`), and the sweep is a known
`/update-knowledge`-style pass against the current clones. The load-bearing deliverable — the member-AI-readiness
contract — has a concrete consumer (M51's seeder).

## Repo split
- **`rosetta`** (this corpus only): `corpus/architecture/*` + `corpus/services/*` re-ground; the member-AI-readiness
  feature doc; parent-index cross-links.
- **No rext code** — this is a docs milestone. (Keeps the file surface **disjoint from M49**, which owns
  `corpus/ops/` + rext.)

## Scope
- **In:**
  - **Sweep `corpus/architecture/` + `corpus/services/`** against the current (M47-synced) clones; reconcile drift —
    services added/removed/changed, signature/flow changes since early June.
  - **Document the member-AI-readiness flow** (the M201 false-negative — ships in prod with live customers, was
    invisible to the stale clones): its surfaces, **data model** (tables / org-enablement flag / the 3-step
    onboarding-evaluation), and the manager dashboard. **This is the contract M51's seeder builds against** — it is
    the milestone's load-bearing deliverable.
  - **Retire stale claims:** the ant-academy "fully RESOLVED at `storytelling-postfix-2`" claim
    (`roadmap-vision.md:139-148` — it is NOT; `repos.yml` still lacks ant-academy → M49 #5); the stale rext-tag
    references in prose (→ M49 #1 makes them a single source-of-truth).
  - **Cross-link** new/updated docs from their parent indexes so they're discoverable.
- **Out:** `corpus/ops/` bring-up docs (M49 truth-up); the actual AI-readiness seeding (M51); the manifest spec (M52).

## Depends on
**M47** (needs the current-prod clones to read). **Parallel with:** **M49** — *yes* (disjoint file clusters:
`architecture`+`services` here vs `ops`+rext there; M48 **never touches the live demo**, M49 monopolizes it → no
contention; additive merge — coordinate only shared parent-index cross-links).

## Open questions (resolve during build)
- Sweep breadth — exhaustive vs material-lag-first. *Lean:* **material-lag-first** (the AI-readiness contract + the
  drifted services), time-boxed; a full corpus rewrite is out of scope.
- Where the member-AI-readiness contract lives — a new `corpus/services/` doc vs a section in an existing one.
  *Decide during the read.*

## KB dependencies (read as contract)
- The **current (M47-synced) clones** — the code of record being re-documented.
- `CLAUDE.md` (the architecture overview + service taxonomy the corpus mirrors) + the `/update-knowledge` method.

## Delivers
- **→ rosetta:** `corpus/architecture/*` + `corpus/services/*` re-ground; **a documented member-AI-readiness
  contract** (the M51 seeder's spec); retired stale "RESOLVED"/tag claims.

## Risk
**(degrades-quality if rushed)** the corpus is large; an exhaustive sweep could balloon. *Mitigate:* material-lag-
first, with the **member-AI-readiness contract as the non-negotiable deliverable** (M51 is blocked without it).
