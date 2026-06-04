---
milestone: M5
slug: demo-corpus-recipes
version: v1.1 "show floor"
milestone_shape: section
status: planned
created: 2026-06-03
last_updated: 2026-06-03
delivers: corpus/ops/demo/ (the demo-env corpus family + recipes) + curated seed presets
---

# M5 — Demo corpus + use-case recipes + skill polish

## Goal
Make demos **repeatable and discoverable**: author the demo-env corpus section, 2–3 end-to-end
"build a demo for use-case X" recipes, 2–3 curated + validated seed presets (small-200 / mid-500 / large-1k),
and polish the demo skill set so any teammate can find and run it. This is the consolidation milestone that
turns M3+M4's mechanisms into a usable, documented product.

## Scope
### In
- **The demo-env corpus section** — a third indexed family (e.g. `corpus/ops/demo/README.md`) with Purpose +
  When-to-Use + an index, matching the existing `staging-*` clustering convention.
- **Cross-linking** the new section into `corpus/README.md` (Ops subsection) + the root `README.md` / `CLAUDE.md`
  navigation.
- **2–3 end-to-end recipe docs** ("build a demo for use-case X": e.g. enterprise-org-onboarding,
  multi-month-skill-progression, hiring-pipeline) following the corpus doc template.
- **2–3 curated seed presets** (small-200 / mid-500 / large-1k) as instances of M4's `demo.seed.yaml`, each
  **validated to actually seed** a stack.
- **Skill polish**: the demo lifecycle skills (`/demo-up`, `/demo-down`, `/demo-status` from M3; `/demo-seed`
  wrapping M4's seeder) added to the root `CLAUDE.md` "Available Skills" table + made discoverable.
- **The v1.0 express-gate CI carry-forward** (default landing spot): wire the `@clerk/express` alignment gate
  into CI now that the demo stack materializes `node_modules/@clerk/express` (needs a Node CI step).
- **Release-boundary doc reconciliation**: ensure the M3 demo-env guide + M4 seeding spec are internally
  consistent and parent docs (README, CLAUDE.md) point at the demo-env story.

### Out
- The demo-env lifecycle guide (`corpus/ops/rosetta_demo.md`) — **M3's** deliverable; M5 only indexes/links it.
- The seeding spec + the seeder code — **M4's** deliverable.
- Any new runtime mechanism / compose change / skill behavior change — M5 is documentation + curation + discoverability.
- AI-generated rich transcripts/embeddings — at most a documented **STRETCH** (only pulled forward from M4's
  deferral if M3+M4 close under budget; trigger decided at M5 kickoff, not pre-committed).
- Any modification to platform repos under `anthropos-dev/`.

## Depends on
**M3** (stacks + lifecycle skills) and **M4** (seeder + `demo.seed.yaml`). Build-blocked until both ship — its
shape is concrete, but it curates M4's instances + validates against M3's stacks. **Parallel with:** none.

## Estimated complexity
**medium** — consolidation with a finite checklist (N recipes + N presets + the index + skill-table edits +
the CI pickup), no uncertain path.

## Open questions (resolve at kickoff)
- How many recipes is "enough" for v1.1 — 2 vs 3 (recommend 3 presets, 2–3 recipes).
- New `corpus/ops/demo/` subdir (cleanest, matches `staging-*` clustering) vs flat `corpus/ops/demo-*.md`.
- `/demo-seed` as a distinct skill vs a documented invocation of M4's seeder (recommend a thin wrapper for parity).
- STRETCH gate for AI content generation — define the trigger at kickoff; don't pre-commit.

## KB dependencies (read as contract)
- The M3 + M4 delivered docs (`corpus/ops/rosetta_demo.md`, `corpus/ops/seeding-spec.md`)
- `corpus/README.md` + `corpus/services/TEMPLATE.md` (doc conventions + where the section links in)
- `anthropos-demo/clerkenstein/` CI workflow (`.github/workflows/alignment.yml`) for the express-gate pickup

## Delivers → `corpus/ops/demo/` (net-new family) + curated seed presets
The demo-env corpus family + recipes + the validated preset configs + the polished, discoverable skill set.

## Exit (section)
All `In:` deliverables land: the demo-env corpus family + recipes + ≥2 validated seed presets exist, the demo
skills are in `CLAUDE.md`, the express-gate CI is wired, and parent docs point at the demo-env story.
