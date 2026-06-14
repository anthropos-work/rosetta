# M27 — Decisions

_Implementation decisions with rationale, numbered `M27-D1`, `M27-D2`, … . Filled during build._

## M27-D1 — Milestone renumber: M26→M27 (secret-provisioning release shifted to M27→M30)

**Date:** 2026-06-14. **Status:** RESOLVED (user decision).

When the build sub-agent went to author this milestone's code in the `rosetta-extensions` authoring copy
(`.agentspace/rosetta-extensions/`), it found the flat milestone number **M26 already consumed** by an orphaned,
untracked ext effort:
- branch `m26/self-contained-demo` @ `25ab855`, **tagged `prop-room-m26`**, committed 2026-06-14 13:21 (after the
  v1.5 close 07:37, before the v1.6 design) — *"M26: make demo stacks self-contained (own GitHub clone set, like
  stack-dev)"*, +521/−141 across 12 files in `demo-stack/` + `stack-injection/` (notably `ensure-clones.sh +106`);
- local-only (never pushed), unmerged to ext `main`, and absent from the rosetta roadmap/state.

The v1.6 "stage door" secret-provisioning release had been designed (minutes later, from a state.md that read
v1.5 = M21→M25) as **M26→M29**, colliding on M26.

**Decision (user, via the work-milestone blocker escalation):**
1. **Keep `self-contained-demo` as the real M26** — its `prop-room-m26` tag + branch stay; it awaits its own
   `/developer-kit:design-roadmap` pass for a version + scope (a separate task, not part of v1.6).
2. **Renumber the secret-provisioning release to M27→M30** — M27 secret-coverage-dna (this milestone) → M28
   provisioning-engine → M29 docs+skill → M30 field-bake. Roadmap, state, context, vision, and the scaffold dirs
   were shifted accordingly.
3. **The stray uncommitted ext WIP** (`clerkenstein/knowledge/architecture.md`, +32 lines, browser-login handshake
   docs) was preserved by committing it on a dedicated ext branch (`wip/clerkenstein-browser-login`), leaving the
   authoring tree clean.
4. **This milestone (M27) is authored on a fresh ext branch off `main`** (`stage-door-m27`), tagged `stage-door-m27`
   on completion — never on top of the stale `m26/self-contained-demo` branch.

Note the future interaction: self-contained-demo touched `ensure-clones.sh`, which M28's provisioning engine plans
to extend (the single `cp` that M28 folds into `stacksecrets provision`). Whoever lands self-contained-demo's
roadmap home should coordinate with M28.
