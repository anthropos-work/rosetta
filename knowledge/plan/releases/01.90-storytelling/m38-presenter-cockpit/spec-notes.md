# M38 — Spec notes

Authoritative design: [`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md) §4d (the
presenter cockpit), §17. Lives in `rosetta-extensions/demo-stack` (a standalone served panel, offset port —
NOT an in-app overlay, preserving the zero-platform-repo-edit line, D15).

## The single-source property
_(the panel reads the SAME `stack.stories.yaml` that seeded the data, D9 — annotations describing a hero are
the same ones that scoped her seed. No drift between data and cockpit menu.)_

## The two actions
_([Login as] → M37 active-user selection. [Jump to section] → the hero's `jump_to` next-web deep-link.)_

## O9 — deep-link catalog
_(enumerate the next-web routes per vantage: profile, Spotlight, workforce tabs, talk-to-data,
talent-pool/mobility; note which need a hero/skill id param. One UI pass.)_
