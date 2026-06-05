---
milestone: M6
slug: dev-stack
version: v1.1 "show floor"
milestone_shape: section
status: done
completed: 2026-06-04
created: 2026-06-04
last_updated: 2026-06-04
delivers: rosetta-extensions/dev-stack/ + rosetta-extensions/stack-core/
---

# M6 — `dev-stack`: tooled local dev environment

## Goal
Bring the **local dev environment** under the same tooled, AI-guided pattern as `demo-stack` — a
`rosetta-extensions/dev-stack/` section that can spin up isolated dev stacks (port-shift, per-stack
project/data), **with Clerkenstein injection OPTIONAL (default OFF — dev uses real Clerk)**, the heavy build
still guided by the rosetta `/setup|start|update-platform` skills.

## Scope
### In
- **`dev-stack/` tooling** reusing the shared multi-instance engine (port-offset, `!override`, per-stack
  project/data isolation, the registry) — proven on demo-stack, generalized here.
- **Optional injection** via `stack-injection` (M5): a dev stack can opt INTO Clerkenstein (e.g. offline dev,
  no Clerk app needed) but **defaults OFF** (real Clerk from `platform/.env`).
- **Skills**: the dev lifecycle remains rosetta's `/setup|start|update-platform` (they stay in rosetta); dev-stack
  provides the multi-instance + optional-injection tooling they can drive.
### Out
- Multiple concurrent dev boxes as a *requirement* — the value is the tooling + optional-injection; N-concurrent
  is a capability, not a goal (demand unproven). Seeding (M7), recipes (M8). No platform-repo change.

## Depends on
M5 (the inject toggle + shared engine). **Parallel with:** none load-bearing.

## Open questions
- Shared platform clones vs per-dev-stack clones (dev workflow differs from demo's per-stack clones).
- Which dev workflows actually want a second stack (testing? a clean-room? offline?) — scope to proven need.

## Exit (section)
`dev-stack/` can bring up an isolated dev stack on offset ports, dev untouched, injection OFF by default and
opt-in-able; documented in `rosetta-extensions/knowledge/`.
