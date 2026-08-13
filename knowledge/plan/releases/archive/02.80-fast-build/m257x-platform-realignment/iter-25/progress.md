**Type:** tik

# iter-25 — the runner could not reset itself, and the fix needed two passes to say so

## Line 1 — `FIX-M257x-iter19-playthrough-runner-path` (LANDED)

`run-playthroughs.sh` called a **bare `stackseed`**. It is not on PATH here and never was: the bring-up
**builds it into each stack's own `bin/`**, deliberately, so a stack runs the tooling at *its* pinned tag
rather than whatever happens to be installed on the box.

The failure shape is the dangerous one. `command not found` on the reset line, and then the run **continues**
into a suite that measures a world it did not reset. iter-15 lost a measurement to exactly this (`17/31` vs
`20/31`, *explained* by a missing `--reset` rather than tested); iter-19 had to hand-supply the path to
measure at all. This is the milestone's own §2 class — a hand-supplied value with no derivation — sitting
**inside the instrument that measures gate clause 2**, the same shape iter-11 found in autoverify's
`STACK_DIR`.

Fixed as a **derivation**: `STACK_BIN` is computed from `N` like `OFFSET` and every other per-N value in the
file, with `PT_STACKSEED` as an explicit escape hatch. And the reset arm now **refuses** — `exit 2`, naming
the path it consulted and saying why — because *"the reset step did not run"* and *"the reset step ran and
did nothing"* must not print the same way (§5 rule 12).

**Controls, live:**
- **negative** — `PT_STACKSEED=/nonexistent … --reset-only` exits **2** with the three-line diagnosis,
  *before* touching the world.
- **positive** — the derivation resolves to the real executable in a consumption clone, and correctly reports
  *absent* in the authoring copy, which has no `stacks/`. That asymmetry is the right semantics: the runner is
  driven from a stack's own clone.

### The second pass, and it is the more interesting half

The first pass fixed the reset arm **and stopped there** — while citing §5 rule 9 (*sweep the sibling leg in
the same pass*) in its own commit message. The live run said so within a minute:

```
run-playthroughs.sh: line 180: stackseed: command not found    (roster export)
run-playthroughs.sh: line 210: stackseed: command not found    (cockpit manifest)
```

Two more bare calls, both **non-fatal by design**, so they degrade **silently**: after a reset-to-seed swaps
the world, the roster the fake-FAPI serves and the manifest the cockpit reads are left describing the
**previous** one. Both now call `"$STACKSEED"` behind an `-x` guard, keeping their documented non-fatal
warning instead of a shell error.

**It was found by running the thing.** The first pass's syntax check passed, its negative control passed, and
its positive control passed — over two live defects in the same file. A control that exercises one arm says
nothing about the other two.

## Line 2 — the clause-2 re-measure (NOT LANDED — routed)

Launched full (`--reset`, no scoping — the ptreport gate is binding only on a full run) from the pinned
consumption clone. The reset **worked, for the first time in this milestone without a hand-supplied path**:

```
▶ reset-to-seed: stackseed --reset --stack demo-1
    stackseed: …/demo-stack/stacks/demo-1/bin/stackseed
Audit: 66 write attempt(s), 55729 row(s) on target "demo-1" (prod=false)
isolation: clean (no shared/external writes landed; 66 audited write attempts covering every surface)
```

The suite then ran **65 of 209** tests in ~35 minutes and was still going, held up on the known 60 s
`waitForURL` timeout class. **The overview's escalation condition fired as written — *do not quote a partial
run as a clause-2 number* — so no number is claimed.** For the record only, and explicitly **not** a
clause-2 measurement: 3 distinct failing ids had appeared (`pt-activity-drilldown`, `pt-assignment-assign`,
`pt-onboarding-hiring-candidate`), and the skill-path Playthroughs that iter-15/19 saw fail on the
`directus_versions` 403 had **not** appeared among them at that point. That is suggestive of iter-24's fix
having landed and is worth **nothing** as evidence until the full run's sorted-id `diff` exists.

The run was left in flight; its log is `.agentspace/scratch/work-m257x/iter25-pt.log`. **iter-26 should
re-run rather than read it** — it ran under `fast-build-m257x-iter-25`, i.e. *before* the second-pass fix, so
its roster and cockpit manifest describe the pre-reset world.

## Close — 2026-08-01

**Outcome:** clause 2's instrument can now reset itself — derived from `N`, refusing loudly when it cannot —
and the fix took two passes because the first swept one of three sibling call sites. The re-measure itself is
routed, unquoted.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close, unchanged; occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n
— (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-25-1, D-M257x-25-2 (this iter's `decisions.md`)
**Side-deliverables:** none.
**Routes carried forward:**
- **`MEASURE-M257x-iter26-clause2`** (next tik, and it is the milestone's remaining substantial work) — full
  suite with `--reset` from a clone pinned at `fast-build-m257x-iter-25b` or later. Compare **sorted failing
  ids** against iter-19's set. Budget it as the iteration's whole scope: the run is ~1 h serial by design
  (`workers:1`, one shared org-scoped Postgres), so an iteration that also plans other work will not finish
  it. That is the mistake this iter made.
- `DOC-M257x-iter23-rext-stale-session-comment` — still open.

**Lessons:**
1. **A long measurement is an iteration, not a phase.** The suite is serial by design and takes about an
   hour; planning it as line 2 of a two-line iter guaranteed the partial. Route the measurement to its own
   iter and let it be the only thing.
2. **Citing a rule is not applying it.** The first pass named §5 rule 9 in its commit message and swept one
   of three sites. The rule is cheap to quote and the sweep is the work.
3. **A control proves the arm it exercised.** Syntax check, negative control and positive control were all
   green over two live defects in the same file — because all three exercised the reset arm. Enumerate the
   call sites, then control each.
