---
milestone: M38
slug: presenter-cockpit
version: v1.9 "storytelling"
milestone_shape: section
status: planned
created: 2026-06-22
last_updated: 2026-06-22
complexity: medium
delivers: rosetta-extensions/demo-stack (a standalone served presenter panel reading stack.stories.yaml — [Login as] wired to M37 + [Jump to section] via a deep-link catalog) + corpus (the cockpit section of stories-spec.md + demo/README.md)
depends_on: M37
b_milestone_for: M37
spec_ref: .agentspace/seeding_gaps.md §4d (the presenter cockpit), §17
---

# M38 — Presenter cockpit (B-milestone for M37)

## Closes the gap after M37
M37 delivers the multi-identity *capability*; M38 turns it into a usable **demo-driving surface** — the
presenter's cockpit. It is the developer/presenter tooling that makes the whole Stories & Heroes engine
clickable, so it follows the B-milestone pattern (tooling that accelerates every future demo).

## Goal
A **standalone served panel** (rext `demo-stack`, on an offset port — **not** an in-app overlay, preserving
the hard zero-platform-repo-edit line) that reads the **same** `stack.stories.yaml` and renders each story →
its hero trio with **[Login as]** + **[Jump to section]**, so a demo-giver picks a hero and lands on the right
screen to present a flow live.

## Why section
The shape is known (a small served panel + two actions, reading the existing YAML). The only research is the
deep-link catalog (O9) and wiring login-as to M37 — both bounded. Fixed checklist.

## Scope
**In:**
- **The panel** — lists stories → heroes (with their `vantage`/`trajectory`/`annotation`), served on an
  offset port like the rest of the demo UI tier, torn down with the stack. Reads `stack.stories.yaml` (the
  single source — D9).
- **[Login as]** — wired to M37's active-user selection (mint + switch the active seat to that hero).
- **[Jump to section]** — redirect to the hero's `jump_to` deep-link in next-web.
- **The deep-link catalog (O9)** — the enumerated, stable set of next-web routes per vantage/use-case
  (profile, Spotlight, workforce tabs, talk-to-data, talent-pool/mobility), noting which need an id param.

**Out:** any change to platform repos (build/serve context only); any data seeding (that's M34–M36).

## Repo split
- **`rosetta-extensions`** `demo-stack` (the served panel + the launch wiring on the demo bring-up).
- **`rosetta`** corpus: the cockpit section of `stories-spec.md` + `demo/README.md` (the up→present flow).

## Open questions
- **O9** — the deep-link catalog (one UI pass over next-web routes).
- **O10/O11 (resolved/decided)** — standalone panel (D15); login-as mechanism inherited from M37.

## Acceleration effect
Every future demo becomes a click-driven, persona-anchored walkthrough — pick a story, log in as the
thriving/struggling/manager hero, jump to the surface that tells that part of the story.

## Done-when
The cockpit lists both stories' trios; [Login as] switches the active browser identity to the chosen hero;
[Jump to section] lands on the right next-web screen; it reads the same stories.yaml that seeded the data;
zero platform-repo edits.
