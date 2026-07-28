---
milestone: M255
title: "build-bench & host-headroom (the barrier)"
release: v2.8 "fast build"
closed: 2026-07-28
close_status: closed-complete
---

# M255 — retro

## Summary

v2.8's HARD go/no-go barrier, and it **PASSED — verdict GO**. Two deliverables landed: the **measurement
floor** (`buildbench` + two measured host profiles + a hard headroom assert; gated baseline **n=3 p50
666.29 s**) and the **barrier verdict** (the multi-stage `.next/standalone` prototype takes the hiring image
**4.84 GB → 379 MB** and its export leg **146.8 s → 2.9 s**, so L1 does not collapse).

The milestone's real value was not satisfying its own scope — it was **reshaping the release**. Three findings
changed what M257 will do, before M257 started.

## Incidents This Cycle

- **P1 — a zombie sub-agent ran for 10.6 h after reporting dead, and clobbered a working tree.** The build
  agent reported `API 529 — Agent terminated early` at ~11:30. That was treated as terminal; it wasn't. It kept
  running while two replacement agents worked the same trees, then landed uncommitted edits over ~140 lines of
  the harden pass's work — including a `read_verdict` freshness guard — leaving 7 tests red. **Recovery:** git
  history was intact (the clobber was working-tree only), the zombie had preserved its own version as patches,
  and the tree was restored to the harden commit. **Nothing was lost.** Root cause: a crash notification was
  read as proof of death rather than verified. The harness even said *"the same task-id may notify more than
  once"* — read as boilerplate. **Lesson: verify a sub-agent is dead before spawning its replacement.**
- **P1 (consequence of the above) — `billion` was used for ~6.5 h during a user-imposed freeze.** The freeze
  landed at 14:00Z and was sent to the agent believed to be alive; the zombie never got it and kept ssh-ing to
  the box until ~20:36Z. Disclosed to the user with the commit-timestamp evidence. **Lesson: a standing
  constraint must be broadcast to every live agent, and "believed dead" is not a filter.**
- **P2 — three API 529s across the session** cost ~40 min of wall-clock across build and harden. Zero work
  lost each time (both trees clean, the remote campaign survived because it was launched detached). Resolution
  the user chose and that worked: wait ~12 min, then re-spawn.
- **P2 — a rext suite had been RED since v2.6 and two closes missed it.** `stack-snapshot`'s directus-surface
  test hardcoded 14 tables; M244 (`e74e563`) added a 15th without moving the count. The v2.6 **and** v2.7 closes
  both reported clean because their test rosters named `stack-seeding` and `playthroughs` but **never**
  `stack-snapshot`. Fixed here. **The process gap is the finding: a rext module can sit outside the close's
  roster indefinitely and nothing notices.**
- **P3 — the academy readiness probe's timeout does not bound wall-clock.** `ACADEMY_READY_TIMEOUT_S=120`
  increments its counter by 2 per iteration on the assumption each iteration costs ~2 s; against a server that
  **hangs** rather than refuses, each iteration also pays its own curl timeout, so the nominal 120 s bound ran
  ~5 real minutes. Non-fatal, observed on a local demo bring-up. Not fixed (out of M255's scope) — recorded
  because it is the same class the milestone spent its day on: a guard whose accounting does not match reality.

## What Went Well

- **Measure-first earned its keep, twice.** Spike (a) took ~15 minutes and replaced a planned 2.5–3 h
  truly-cold campaign that would have tested the wrong hypothesis. Spike (d) then **refuted both** standing
  hypotheses and re-priced L2 from ~200 s to ≲45 s — a correction that would otherwise have been discovered
  inside M257, after the work was planned around the wrong number.
- **The headroom model was validated, not merely applied.** It predicted 5,400 MiB; the three reps peaked at
  5,446 / 5,579 / 5,398 MB (~3 %). That is what makes its refusals credible — including the one it issued
  against the laptop, which was reported honestly rather than overridden.
- **The adversarial plan review paid for itself before implementation.** It cut a deliverable with no consumer
  (`hostprofile` the auto-planner — it would have measured and changed nothing, the very defect this release
  retracts), and it caught that M256's 120 s speed clause was arithmetically impossible.
- **Fences were tested for their ability to fail.** The harden pass found `union_apply_guard`'s clause (i)
  shipped as `_sha(p) != _sha(p)` — a tautology that could never fire — and the `Read at` guard reporting green
  while **28 of 29** citations had rotted. Both are now covered by a 10-mutant RED-proof battery.

## What Didn't

- **Concurrency discipline across sub-agents.** Three agents on two trees with no lock and no liveness check.
  The clobber was recoverable only because git history held and the zombie preserved its own work — neither was
  by design.
- **`state.md` breached its 15 KB cap four separate times** during this milestone and had to be trimmed each
  time. The cap is real and enforced at close; the writing habit that pushes past it is to narrate in `state.md`
  what belongs in `roadmap.md` / `progress.md`.
- **The mirror fence should have existed before the numbers were mirrored.** `666.29` was quoted across 8 prose
  sites while living only inside a prose evidence string in `billion.json` — no comparable source at all. Three
  separate count-drift incidents surfaced in a single day (the directus 14/15, the 28 drifted citations, the
  27-vs-30 knob count) before the fence was written.
- **"M255 harden resume" was accepted as a Fate-3 destination.** It is not a named milestone, so four items
  were routed somewhere that could not survive the close. Caught at close and re-fated to M257; should have
  been rejected when written.

## Carried Forward

**Fate 3 → M257** (recorded in `m257-first-light-build/overview.md` § Inherited from the M255 close), each
attached there because M257 is the milestone that exercises it:

| item | why M257 |
|---|---|
| `run_campaign` rep-body coverage | M257 runs the campaign on every gate cycle |
| `demo_knob_guard` anchor-fence mutants | M257's levers add/rename `DEMO_*` knobs |
| `_manifest_lists` silent truncation | M257 flips union-apply on, which reads those lists |
| the `laptop` profile's provisional field | M257 is the first to quote host-profile numbers in a gate |

**Fate 1 landed at close:** the plan-number mirror fence (`test_baseline_mirror_fence.py`).

**Not carried, recorded only:** the academy readiness-probe wall-clock defect (P3 above) — it belongs to the
demo bring-up path, not to v2.8's scope. Left as a named observation rather than a silent omission.

**Zero** escape-hatch (cross-release) deferrals. **Zero** dropped items.

## Metrics Delta

| | v2.7 close | M255 close | Δ |
|---|---|---|---|
| Python (rext) | — | **1505 pass / 2 skip / 0 fail** (1507 tests) | stack-core alone **226 → 272** |
| Go test funcs (rext) | 2019 | **2023** | +4 |
| Go modules failing | (not measured) | **0 of 6** | — |
| Flake count | 0 | **0** (3 sequential runs) | — |
| Platform-repo edits | 0 | **0** | — |
| Net-new deps | 0 | **0** | — |

Full artifact: [`metrics.json`](metrics.json).
