# iter-253 — decisions

## `D-M257x-253-1` — the fresh-checkout class is decidable DYNAMICALLY, and not statically at the file grain

The obvious fence is a text predicate: *a test file that names operator-local state and carries no skip
idiom*. It was built and measured **before** the instrument was designed, and it is the wrong instrument.

| instrument | population |
|---|---|
| static, file grain (`stack-dev`/`stack-demo`/`.agentspace` named, no `skipUnless`/`skipIf`/`pytest.skip`) | **8 files, 6 of them in `stack-core`** |
| dynamic, measured on a frozen clone pair | **13 files** |

The counter-example that settles it is `test_toolchain_floor_guard`: it **carried** a `skipUnless` and failed
anyway, because it declared **half** a precondition — the rosetta checkout, not the clone set the Node floor
is derived from. The defect lives at the **test-function** grain and is conditional on state a text scan
cannot evaluate. So the census RUNS the tests with the state absent; it does not read them.

## `D-M257x-253-2` — the live CONTROL is the load-bearing half, and this iter proved it the hard way

A frozen-tree failure list on its own says only *"these failed over there."* The classification into
*environmental* vs *real* requires re-running the same node-ids **here**, and the first run of this census
returned findings in **both** buckets. Without the control, **5 genuine REDs at HEAD would have been
published as fresh-checkout artifacts** — a real defect absorbed by the bucket built to excuse the
environment, which is the exact failure mode the three-way partition exists to prevent (`§5` rule 73).

`REAL` findings print **before** `BOX` findings, for the same reason.

## `D-M257x-253-3` — the repair is a `skipUnless` in the test, NEVER an `ENV_GATED` entry

`suite_census` already owns a declared environment bucket, and adding these 27 node-ids to it would turn
the census green in one edit. It is refused, and the refusal is **printed in the report where the reader is
standing when they are tempted**, plus pinned by a test:

> the census would be green and **every one of those tests would still FAIL** for a person who has just
> cloned the two repos.

`ENV_GATED` absolves the *instrument*; `skipUnless` fixes the *test*. Only the second changes what the
reader experiences. (`D-M257x-249-2` chose skip-over-fail; this decides where the declaration lives.)

## `D-M257x-253-4` — a reading's names must be DURABLE, and iter-249's were not

`ROUTE-M257x-249-a-reading-must-name-its-failures` reads as though the names were never written. They
**were** — all 29 of them, correctly. They were written to
`.agentspace/scratch/work-m257x/iter249-failing-ids.txt`, and `git check-ignore` answers
`.gitignore:138:.agentspace/`. **The names existed and could not survive the session.**

That is the same mechanism as the milestone's highest-leverage open item
(`ROUTE-M257x-h59-rext-edits-fire-no-fence-anywhere`: `.agentspace/` is git-ignored, so tooling edits fire
nothing). Recorded here because the distinction changes the fix: the discipline to add is not *"write the
names down"* but *"write them where the next reading can diff against them."* iter-249's 29 are rescued
into `evidence/iter249-frozen-failing-node-ids.txt` in this iter dir, which is how this iter's diff against
them was possible at all.

## `D-M257x-253-5` — a class-level `skipUnless` protects the CLASS, and the file keeps growing classes

`test_frozen_expectation_census_m257x.py` had **5** hostile tests at iter-249 (`LiveTree`), was repaired
with a class-level declaration that works, and returns **5** hostile tests now — from **four different
classes**, none of which inherits `LiveTree`'s declaration. Attribution refuted the tempting story: those
four classes are **older** than the repair (iter-207 and harden passes 52/55/61), so they were not
manufactured after it. They became RED for a different reason, measured in the same run — a breached
ratchet — which is `D-M257x-253-2`'s point restated: *the bucket a failure belongs in is not readable from
the file it lives in.*

## `D-M257x-253-6` — `--shape` ships only the tree it implements

iter-251 measured that *"a fresh checkout"* is **two trees that give opposite answers**. `pair` is
implemented; `rosetta-only` is declared in `FRESH_CHECKOUT_SHAPES` with `NOT IMPLEMENTED` and its route id,
and is refused at both the API and the CLI. **A flag that does not work is a false promise** — the milestone
has paid for that shape before. A test asserts every declared shape is either implemented or says it is not,
so the two lists cannot drift apart.
