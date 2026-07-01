---
milestone: M52
slug: seed-manifest
version: v1.10b "fit-up"
milestone_shape: section
status: archived
created: 2026-06-29
last_updated: 2026-07-01
complexity: medium
delivers: corpus/ops/demo/seed-manifest-spec.md (NEW — the consolidated single-file seed+generation contract)
issues: user note 3 (single auditable seed/generation manifest — one inlined file)
---

# M52 — Single auditable seed+gen manifest

## Goal
Consolidate the scattered seeding + generation direction into **one checked-in, auditable file** that drives all
seeding + generation (**all three orgs**, including the M51 AI-readiness story), and make the cockpit's
**[Download manifest]** serve **that** file — so a presenter (or auditor) can read the entire seed+gen intent in one
place, without reading Go code. **Cache + generated data excluded.**

## Why section
Concrete, enumerable deliverables: extract the Go prompts to YAML, author the consolidated manifest, repoint the
cockpit download, preserve/re-key the cache. The user picked the **"one inlined file"** option — the shape is fixed.

## The gap this closes (today's scattered state)
- **`stories.seed.yaml`** — population blueprint (orgs, heroes, role mix, activity span). *In a file ✓*
- **`gen-batch-*.seed.yaml`** — the M45 batch descriptors (count, role mix, narrative). *In a file ✓*
- **The "mother prompts"** — embedded in **Go templates** (`stack-seeding/blueprint/batch.go` → `EffectiveBatches()`).
  *NOT auditable without reading code ✗ — this is the core gap.*
- **The cockpit `/manifest.json` download** — a **stories→heroes projection only**: no population size, role mix,
  prompts, cost caps, or snapshot sources. *Incomplete ✗.*

## Repo split
- **`rosetta-extensions`** (authoring copy → tag `fit-up-m52` → consume per-stack): extract prompts to YAML in
  `stack-seeding/`; the consolidated-manifest reader/writer; `demo-stack/cockpit.py` download repoint.
- **`rosetta`** (this corpus): a **new** spec doc (the consolidated single-file contract — blind area ③).

## Scope
- **In:**
  - **Extract the mother prompts** from `stack-seeding/blueprint/batch.go` (`EffectiveBatches()`) into YAML — the
    generation prompt templates become file-resident + auditable, no recompile to change.
  - **Author one checked-in `seed-generation-manifest.yaml`** inlining everything: the **population blueprint** (all
    3 orgs incl. AI-readiness), the **generation prompt templates**, the **batch config** (`--max-cost` ceiling,
    concurrency, re-roll rules), and the **snapshot sources** (taxonomy + Directus capture versions). **EXCLUDE** the
    `.agentspace/.batchcache` + generated member envelopes (the user's cache/generated-data exclusion).
  - **Repoint the cockpit [Download manifest]** (`demo-stack/cockpit.py`) to serve the consolidated file (replacing
    the stories→heroes projection; keep the projection only if a consumer needs it).
  - **Cache integrity:** the M45 prompt-hash cache is keyed on the **mother prompt** — the extraction must yield the
    **same effective prompt** (cache stays valid) **or deliberately re-key** (a documented, intentional invalidation,
    coordinated with M47's recapture version bump).
- **Out:** new seeding behavior (M50/M51 own that — M52 only *expresses* it auditably).

## Depends on
**M50** + **M51** (all stories must be final before the single file can capture them). **Parallel with:** the
manifest code is authorable alongside M51 in the authoring copy, but its **live-verify** (the cockpit download)
serializes behind M51 on the single demo.

## Open questions (resolve during build)
- File location — `.agentspace/` vs a `presets/`-adjacent path. *Lean:* a checked-in path the cockpit + seeder both
  read (not git-ignored, unlike the cache).
- Keep the legacy stories→heroes projection endpoint for back-compat vs replace outright.

## KB dependencies (read as contract)
- `corpus/ops/demo/ai-generation-spec.md` (the CODE-owns-structure / AI-owns-content boundary + the mother prompt),
  `corpus/ops/demo/cache-spec.md` (the prompt-hash cache key — the integrity constraint),
  `corpus/ops/seeding-spec.md`, `corpus/ops/demo/cockpit-spec.md` (the download surface).

## Delivers
- **→ rosetta-extensions:** the extracted prompt YAML + the consolidated `seed-generation-manifest.yaml` + the
  cockpit download repoint, tagged `fit-up-m52`.
- **→ rosetta:** **NEW** `corpus/ops/demo/seed-manifest-spec.md` — the consolidated single-file seed+generation
  contract (what's inlined, what's excluded, how the cockpit serves it, the cache-key integrity rule).

## Risk
**(degrades-quality)** a careless prompt-extraction silently busts the M45 cache (a $0-re-seed regression).
*Mitigate:* the cache-integrity rule above — same effective prompt, or a documented deliberate re-key.
