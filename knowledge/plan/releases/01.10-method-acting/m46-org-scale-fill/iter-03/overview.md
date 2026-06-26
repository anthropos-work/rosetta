---
iter: 03
milestone: M46
iteration_type: tik
status: closed-fixed
created: 2026-06-26
---

# M46 · iter-03 (tik #2) — per-story batch distribution (deliverable #2)

**Active strategy reference:** TOK-01 — build the three deliverables fixtures-first. Tik #2 = deliverable
#2 (per-story batch distribution).

**Re-survey (Step 0):** TOK-01 named per-story distribution. The KB-fidelity audit confirmed
`GeneratedBatchSeeder` hardcodes `stories[0]` (org/prefix/domain resolved once, all members written there).
Still untouched (iter-02 only touched `batch.go`/`blueprint.go`), still meaningful — the multi-org Stories
model needs each org's supporting population filled from its own batch.

**Cluster / target identified:** `GeneratedBatchSeeder.Seed` — the org/prefix/domain are resolved from
`stories[0]` once, outside the member loop, so every generated member lands in the first org regardless of
which story's batch produced it.

**Hypothesis:** carry the story index on each member (`BatchMember.StoryIndex`, set at expansion from
`st.Index`; NOT in the mother prompt → cache unchanged), then resolve org/prefix/domain/company PER member
from `m.StoryIndex` inside the loop. Each story's members then land in its own org.

**Expected lift:** the structural capability "a multi-org blueprint fills EACH org's supporting population
from its own batch" lands + is unit-proven, and it composes with the iter-02 fill (each org fills to its
OWN Size). (Gate metric on a generated org unmoved until the real-run tik — a build-toward-the-gate
deliverable.)

**Phase plan:** code (`BatchMember.StoryIndex` in batch.go; per-member story resolution in
generated_batch.go) → fixtures-first unit tests (per-story routing → distinct orgs/domains/ids; fill +
per-story composition) → full stack-seeding suite green.

**Escalation conditions:** none expected. A cache-key change from carrying StoryIndex would break the $0
reseed → would need StoryIndex kept OUT of the mother prompt (it is).

**Acceptable close-no-lift outcomes:** N/A — deterministic code deliverable.
