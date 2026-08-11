**Type:** tik — under `TOK-08`, working `ROUTE-M257x-236-host-is-the-unreliable-witness` (open since
iter-236, never worked).

# iter-241 — the clone set is an instrument, and the same green meant two different things

## The measurement

`platform_alignment_guard` derives its unclonable set from disk presence
(`not (clones_root / head).is_dir()`), so *"the repos no stack clones"* is a fact about **the machine
that ran it**. Run against the **same corpus**, on two clone sets:

| clone set | citations NOT checked | repos excused | exit |
|---|---|---|---|
| this laptop — **13** clones | **11 of 109** | 2 (`db-backup`, `infrastructure`) | **0** |
| a fresh bring-up's set — **7** clones | **27 of 109** | 7 (+ `cms`, `jobsimulation`, `messenger`, `roadrunner`, `storage`) | **0** |

> ### Both GREEN. The guard checked **16 fewer citations** on the fresh box, and the only thing that said
> ### so was a number nobody was comparing — in a sentence with **no denominator and no roster**.

The fresh set was not imagined: it is `repos.yml`'s four repos + the two extras `clone_pin_guard`
sanctions + the rext consumption clone, symlinked into a temp root and passed to the guard's own CLI.

## Why the two boxes differ — and why the laptop is the anomaly

`stack-demo/` holds **13** git clones. Seven are accounted for. The other **six** — `cms`,
`graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner`, `storage` — are exactly the repos
`repos.yml`'s own header comment describes as frozen legacy that *"`make init` therefore does not clone…
clone them by hand if you need to read the pre-merge source."* They are **leftovers from before
`838d907`**, and `ensure-clones.sh` adds only `ant-academy` + `platform` on top of `make init`.

**So the workspace that has been used longest verifies the most, and a clean machine — the one the exit
gate actually describes — verifies the least. Seniority of a workspace reads as coverage.**

## The exposure is much larger than the guard's subject

The guard grades **one file** (the migration map, 109 citations). Corpus-wide, `CLAUDE.md` + `corpus/**`
carry **107 citations into those same six repos** — `cms` 41 · `jobsimulation` 21 · `messenger` 15 ·
`roadrunner` 12 · `graphql-wundergraph` 10 · `storage` 8 — including the merge-banner anchors this
milestone leans on most (`cms/terraform/main.tf:39`, `messenger/terraform/main.tf:29`,
`jobsimulation/terraform/main.tf:15-22`). **Nothing grades those at all**, and their verifiability
depends on which box asks.

## What landed

**The verdict now states its reach** — numerator, denominator, and the clone-set roster:

> `OK OVER ITS REACH — … 11 of 109 citation(s) — into 2 repo(s) … — were NOT checked. Clone set read
> (13): ant-academy, app, cms, … A different clone set gives a different reach under the same green.`

That sentence is the one `guard_family.run_one` reports for a green member, which is why the qualifier
had to ride *inside* it rather than in a line above (the guard's own iter-91 comment says so; the reach
clause simply had not been finished).

**Two regression tests** (`ReachIsStated`, file 64 → 66, all green): one pins the shape; the other builds
the restricted clone set by symlink and requires the unchecked count to **rise** while the denominator —
a property of the map, not of the box — **stays put** (`D-M257x-241-3`). A regression making `M`
clone-set-dependent would pass a naive "the number got bigger" check and fail this one.

**The corpus disclosure** — `corpus/ops/platform-alignment.md` §8 gains
*"A fence's REACH is a property of the CLONE SET"*, carrying both readings, why the laptop is the
anomaly, the general rule, and the gate-clause-3 consequence (the map is fenced; the wider 107-citation
surface is not).

**The six leftovers are NOT deleted** (`D-M257x-241-2`) and the clone set is **not** grown
(`D-M257x-241-1`): curating a clone set to make a fence green is `§8`'s fetch rule one level out —
measuring a memory of the platform rather than the platform.

## Pre-registration — scored 3 confirmed / 1 refuted / 1 partly confirmed

| claim | prediction | result |
|---|---|---|
| `P-241-1` exactly 6 unaccounted clones, by name | 6 | **CONFIRMED** — `cms`, `graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner`, `storage` |
| `P-241-2` excused-repo count rises 2 → 8 | 8 | **REFUTED — 7.** See below; the miss is the useful part |
| `P-241-3` ≥ 20 corpus citations into the 6 | ≥ 20 | **CONFIRMED**, understated **5×**: 107 |
| `P-241-4` the verdict does not name its clone set | absent | **PARTLY CONFIRMED** — it named the *excused repos*, but neither the clone set read nor a denominator |
| `P-241-5` `ensure-clones.sh` clones 0 of the 6 | 0 of 6 | **CONFIRMED** — `make init` (repos.yml) + `ant-academy` + `platform` only |

**`P-241-2` was wrong for a reason worth more than the prediction.** I predicted 8 by adding all six
leftovers to the two already-excused repos. The answer is 7 because **`graphql-wundergraph` is cited 10
times corpus-wide and not once in the migration map** — the guard's population is the map's citations,
not the corpus's. The off-by-one *is* the fourth denominator error of this run's lineage: I compared a
count derived from one population against a guard whose population is a different, smaller one. **Ask
what the number is a number OF — even when both numbers are yours.**

## Close — 2026-08-10

**Outcome:** the alignment fence's reach is now stated (numerator, denominator, clone-set roster) and
regression-tested, and the fact it discloses is measured: the same corpus reads **11 of 109** unchecked on
this box and **27 of 109** on a fresh bring-up's clone set, both GREEN.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-241-1` (disclose, do not grow the clone set) · `D-M257x-241-2` (the 6 leftovers
stay) · `D-M257x-241-3` (the denominator must be clone-set-invariant, and it is tested) ·
`D-M257x-241-4` (the wider 107-citation surface is disclosed, not fenced, in this iter).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — `stack-core` (pytest 8.4.2 / CPython 3.9.6, Python):
`tests/test_platform_alignment_guard.py` **66 passed / 0 failed** (was 64; +2 this iter). iter-239's 17
and iter-240's 14 still green. Guard family at platform reach: **26 GREEN / 0 RED / 0 could-not-check /
5 not-run**.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-241-wider-citation-surface-is-ungraded` → **new.** **107** corpus citations into the six
  frozen-legacy repos, outside the migration map, graded by nothing. Fencing them needs the
  base-directory model this iter's path census showed is not mechanical — a real piece of work, now
  stated with its number in `platform-alignment.md` §8 rather than left implicit.
- `ROUTE-M257x-236-host-is-the-unreliable-witness` → **worked, and narrowed rather than closed.** The
  alignment guard now discloses its clone-set dependence; **the other 30 family members were not
  audited for the same shape**, and at least `unreadable_repo_claim_guard` is explicitly built on a
  disk-presence premise it re-measures every run.
- `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` → open.
- `ROUTE-M257x-239-stackseed-sentinel-reload-is-demo-only` → open.
- `ROUTE-M257x-238-claude-md-fences-are-unmaintained` → open, six-for-six.
- `ROUTE-M257x-238-container-vs-native-is-undrawn` → open.
- `ROUTE-M257x-237-critical-env-list-is-unfenced` → open.
- `ROUTE-M257x-236-disclosure-scope-is-document-level` → open.
- `ROUTE-M257x-235-fence-scope-is-unread` → open.
- `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **Seniority of a workspace reads as coverage.** The box that has been used longest has the most
   clones and therefore verifies the most — while the gate describes a cold machine, which verifies the
   least. Any guard that derives a capability from disk inherits this backwards.
2. **A fraction needs both numbers pinned, and only one of them is about the box.** Reporting `N of M`
   is only an improvement if `M` cannot move with the clone set; otherwise it is two moving numbers
   pretending to be a rate. That is a test, not a comment.
3. **The off-by-one in my own prediction was a denominator error.** `P-241-2` derived 8 from the corpus's
   population and compared it to a guard whose population is one file. Fourth of this lineage — and the
   first where both numbers were mine.
