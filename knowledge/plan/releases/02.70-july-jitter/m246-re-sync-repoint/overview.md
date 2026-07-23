---
milestone_shape: section
milestone: M246
title: "re-sync & re-point (clean-stage barrier, HARD go/no-go)"
status: planned
release: v2.7 "july jitter"
depends_on: []
parallel_with: []
complexity: medium
created: 2026-07-23
---

# M246 — re-sync & re-point (clean-stage barrier, HARD go/no-go)

## Goal
The demo builds + comes up GREEN from the CONSOLIDATED platform (current `origin/main` — 3 subgraphs,
skillpath-in-app), and the rext seeder writes to `public.skill_path_sessions` — so every downstream fix is
scoped against reality and the imminent seeder break (the seeder is one `stack-update` away from breaking on
`skillpath.skill_path_sessions`) is defused. Emits a confirmed-drift ledger for M247.

## Shape (why this shape)
`section` — a HARD go/no-go barrier, the M237/M222 clean-stage pattern. It opens the release and gates every
other milestone: any fidelity fix authored on stale clone pins is untrustworthy, so bring-up on the
consolidated platform (skillpath decommissioned into `app`, jobsim mid-merge) must be proven GREEN first. The
load-bearing seeder re-point rides here too — it defuses the imminent break rather than letting a downstream
milestone trip on it. The barrier's payoff is a confirmed-drift ledger that scopes M247's doc reconciliation
against reality, not the pre-consolidation topology the corpus + rext tooling still describe.

This is Risk R1's mitigation (the 386-commit bump may break bring-up): prove bring-up first, do the
load-bearing seeder re-point, and scope the whole thing to the **disposable demo clone set — the dev native
worktrees are never touched**. If jobsim forces a change, re-scope downstream.

## Scope
### In
- Re-point rext `stack-seeding` writes `skillpath.skill_path_sessions → public.skill_path_sessions` in the
  **live** seeder code + tests: `cmd/stackseed/main.go:97`, `seeders/hero_activity.go:180`,
  `skillpath_sessions.go`, `content_nonsim.go`, `dna/data-dna.json`, + the in-package test assertions. Leave
  surface **names** (`skillpath-sessions`) and the mirror `public.local_skill_path_sessions` untouched.
- Author `stack-demo/clones.pin.json` + wire the `DEMO_ADVANCE_CLONES=pinned` advance path; bump the **demo**
  clone pins to current `origin/main` (jobsimulation stays standalone — still live).
- Fix the stale `stack-injection/gen_injected_override.py:16` skillpath comment (3 subgraphs).
- Prove ONE cold `/demo-up` GREEN on the consolidated platform; transcribe observed drift into the M247 ledger.

### Out
- The corpus doc reconciliation (M247).
- Any fidelity fix.
- Any platform-repo edit.
- Touching the dev native worktrees.

## Dependencies & parallelism
- **depends_on:** `[]` — none; opens the release.
- **parallel_with:** `[]` — none; a HARD go/no-go barrier that gates everything. M247–M254 fan out off the
  **post-M246 HEAD** (M246 touches `up-injected.sh`-adjacent pin machinery + the seeder package).
- **Intra-milestone lanes:** 4 concurrent prep lanes over disjoint files, then a single serial gate.
  - **Lane A — seeder-repoint-core:** `seeders/hero_activity.go:180`, `skillpath_sessions.go`,
    `content_nonsim.go` + their in-package test assertions.
  - **Lane B — driver + DNA:** `cmd/stackseed/main.go:97` + `dna/data-dna.json`.
  - **Lane C — pin-mechanism:** author `stack-demo/clones.pin.json` + wire `DEMO_ADVANCE_CLONES=pinned` + bump
    the demo clone pins to current `origin/main`.
  - **Lane D — injection-comment + ledger-scaffold:** fix `stack-injection/gen_injected_override.py:16` (3
    subgraphs) + scaffold the M247 drift ledger.
  - Then the **single serial `/demo-up` prove** — the go/no-go gate; it cannot be sharded (one bring-up).
- **Speedup:** ~1.3–1.5× (the serial gate is the floor). The seeder split is **core-vs-driver only** because
  `activity_seeders_test.go` couples `hero_activity` + `skillpath_sessions` — keep them together in Lane A
  rather than splitting per-file.
- **Recommended subagents:** up to 4 for the disjoint prep lanes (A / B / C / D), converging on **1 driver**
  for the serial `/demo-up` gate. Rung-zero: rext tags pushed to **origin** before billion re-pins.

## KB dependencies
- `corpus/ops/update_guide.md`
- `corpus/ops/rosetta_demo.md`
- `corpus/ops/demo/demopatch-spec.md`
- `corpus/services/skiller.md` (the redirect pattern)

## Delivers
- `corpus/ops/update_guide.md` — the consolidation re-sync note.
- A confirmed-drift ledger artifact for M247.

## Open questions
- Dual-schema-tolerant during the transition, or hard-cut to `public` + bump pins in the same milestone? (The
  deprecation alias suggests a clean window.)
- Does the 386-commit bump surface a migration/subgraph break that re-scopes the downstream milestones?
