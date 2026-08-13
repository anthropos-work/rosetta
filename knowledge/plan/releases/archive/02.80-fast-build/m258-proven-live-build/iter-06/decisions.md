# M258 iter-06 — decisions

## D15 — the batch is default-ON, but it SKIPS on a `--public-host` stack, and the skip is recorded as `skipped`, never `green`

The gate's own text is *"one cold command brings the stack up AND drives the full Playthrough batch"*, and
the house style for this family is that **every feature knob is an opt-OUT** (`DEMO_NO_*`, default `0`).
So the batch is default-ON. That much was settled by the milestone.

What was **not** settled, and what this iter had to decide, is what happens on the **default** demo. And
the two defaults collide: `--public-host` is *also* default-on (`D-DESIGN-3`), and a `--public-host` demo
**cannot be browsed from its own host** — docker-proxy binds `0.0.0.0`, so a connection from the demo host
to its own tailscale IP bypasses `tailscale serve`, which is what terminates TLS
(`run-playthroughs.sh:92-105`; M255 spike (e)).

Composing the two naively would run the batch on a stack it cannot reach and **red all 30 Playthroughs for
a reason that has nothing to do with the product** — on *every default bring-up*. That is a false-RED
factory, and this release has already written down why that is the worse direction: M256 fixed a false-RED
in this very runner and recorded that **a false RED trains its operator to disbelieve the gate.** Shipping
one deliberately, by default, would be the same defect chosen on purpose.

**Decision:** on a public-host stack the gate skips itself, loudly, prints the peer path, and records
`verdict: skipped`. **`skipped` is not `green`** — the gate is never reported as met where it was never
taken. This is fail-closed on the *claim* while staying non-fatal to the *stack*, which is the same split
`autoverify` already draws.

**The consequence is stated rather than buried:** a bare `/demo-up N` **skips** the batch;
`/demo-up N --no-public-host` **gates** it. That is exactly the tension `overview.md` § *Open questions*
flagged and `TOK-01` disclosed — *"this proves the composition in a mode the presenter never uses"* — and
it is now enforced in code and printed at the point of use, instead of living in a plan document. The
alternative (drive the browser half from a tailnet peer) is a two-machine flow the bring-up cannot perform
by itself; the skip message names it in full.

## D16 — the batch gate and the restore leg are ONE deliverable, not `TOK-01` steps 2 and 3

`TOK-01` orders them as separate steps and the routing carried them as separate items. Building them
separately would have been wrong, and the reason is in the gate's own text.

A batch wired **without** a restore leg is not a partial delivery — it is a **regression**. The gate
requires the stack be left *"in a presenter-usable world"*, and `overview.md` § *The world contract* shows
the naive composition ending in a test world behind a cockpit full of **dead CTAs** — a state that fully
satisfies *"the stack is left UP regardless"*. M254 left `billion` in exactly that state. Wiring step 2
alone would have made that the outcome of **every** bring-up, and the milestone would have shipped the
defect its own overview was written to prevent.

So they landed together, declared as a planned two-step shape in this iter's `overview.md` so the
scope-creep tripwire graded them against the plan rather than as drift. The restore runs on **every** path
where the reset ran — **including a RED batch**, because a red test result must not *also* cost the
presenter the demo world.

## D17 — the batch needed a `buildbench` phase anchor, and without one the number would have been WRONG, not missing

`BRINGUP_ANCHORS` attributes sub-phases by *"each anchor's timestamp starts its phase and the next anchor's
ends it"*, so the **last** anchor absorbs everything to the end of the run. The last anchor was
`autoverify`.

Wiring the batch in after it — without an anchor — would therefore have billed the entire batch
(~129 s + the restore leg) to a phase that genuinely takes ~2-3 s. **The phase table would still have
summed**, `phases_complete` would still have read `true`, and nothing would have flagged it. This is the
distinction worth keeping: not a *missing* number, a **wrong** one, attributed to the phase least able to
explain it — and M258's whole method is per-phase attribution of a composed budget.

The anchor ships with a **requirement class of its own** (`batch`), because the phase is legitimately
absent in two independent cases (`DEMO_NO_BATCH=1`, and any public-host run where the gate skips). It
defaults to *not applicable*, so **no historical ledger turns red for lacking a phase that did not exist
when it was recorded** — and `test_the_batch_is_not_billed_to_autoverify` pins the mis-attribution itself,
asserting `autoverify` stays ~3 s while `batch_gate` takes the ~160 s.

The anchor is a **string coupling across two files in two languages**, so it is fenced in lockstep
against `batch-gate.sh`'s real source line (not against a fixture, which would keep passing forever while
the bring-up drifted away from it). Mutation-verified in both directions.

## D18 — `RESTORE-M258-world-contract` was routed as owed IN FACT, and that was refuted on re-survey

iter-04 routed it as *"now owed in fact, not in principle — `demo-1` is currently a Playthrough world
behind a cockpit projected from the stories preset."* Re-surveyed at this iter's open, `demo-1` held **4
story orgs** (Cervato Systems, Meridian Talent, Northwind Aviation, Solvantis), **591 users**, and a
`cockpit-manifest.json` advertising all four hero trios. **It was not a pt-world stack.**

The explanation is benign and was already recorded one iter later: iter-05's three campaign bring-ups each
re-seeded the presenter world, restoring it **incidentally**. iter-04's routing was true when written and
stale by the time it was read.

This is the fourth time this release a routed item did not survive contact with the evidence, and it is
worth separating from the other three: the earlier ones were **wrong diagnoses**, this one was a **true
observation that expired**. The corrective is different — not "verify the diagnosis" but *"re-measure the
state a routed item asserts, because a state claim has a shelf life."* The item itself was **not** dropped:
the restore leg is owed as a **mechanism the wiring must carry**, which is a stronger obligation than the
one-off repair the routing described, and it landed in this iter.

## D19 — the restore wrote the presenter menu into a STALE clone, and only the live run could have found it

**This is the iter's most important finding, and no unit test could have produced it**, because every
fixture mirrors ONE tree and the defect exists only on a box with **two** of them.

`rosetta-extensions` has two clone roles (CLAUDE.md § *where stack tooling lives*): the **authoring copy**
(`.agentspace/rosetta-extensions`) and the **per-stack consumption clone**
(`stack-demo/rosetta-extensions`). `up-injected.sh` derives its stack dir from `$HERE`, so
`demo-stack/stacks/demo-1/` exists in **both**, and only one of them is the stack's. Measured this iter:
the live dir was the consumption clone's (mtimes 08:08–08:12, iter-05's campaign); the authoring copy's
was **stale since Aug 11 23:15** (iter-04) — and it still contained a working `bin/stackseed`, a parseable
`cockpit-manifest.json` and plausible logs. **Nothing about it looks stale.**

`restore-presenter-world.sh` derived its stack dir from `$EXT_ROOT` — i.e. from where the *script* lives.
Run from the authoring copy against a stack brought up from the consumption clone, it:

| leg | path source | outcome |
|---|---|---|
| DB reset + seed | docker | ✅ correct (4 orgs / 591 users restored) |
| roster | **`docker inspect`** mount | ✅ correct (35 stories identities, live path) |
| cockpit + content manifests | **`$EXT_ROOT`** | ❌ **written into the stale clone** |

So the live pair became a **stories roster beside a pt-world menu** — the cockpit advertising 11
`pt-*` seats no roster identity could serve. That is the *precise* stale-menu failure the restore exists
to prevent, **reintroduced by the restore itself**, while the batch gate printed *"presenter world
restored"* and exited **0**. Measured, not inferred: the live `cockpit-manifest.json` held
`['pt-employee','pt-manager','pt-free',…]` at 08:39 beside a 35-identity stories roster at 08:42.

`run-playthroughs.sh` already had this right and **says so in a comment** — *"the mount path is discovered
via docker inspect so it works regardless of which rext clone drives the run."* The rule existed; it was
simply not reused. It is now `stack-paths.sh::resolve_stack_dir`, applied in both scripts.

**The rule, stated generally: ask DOCKER where a stack's files are.** The container is the only party that
cannot be wrong about which directory the stack actually reads; a path derived from the script's own
location is an *assumption*, and the mount is *evidence*. The fallback remains for a single-clone box, but
it is never preferred.

## D20 — the restore is now SELF-VERIFYING, because the defect passed every step-level check

D19's failure is not caught by checking the steps. **Every export reported success** — the reset seeded,
the roster exported, the manifest exported — and each was individually true. The damage lived in the
*relationship between two artifacts*, which no step can see.

So the restore gained a post-condition that asks the only question that matters afterwards: **can every
seat the cockpit advertises actually be logged into?** It compares the two **files on disk** (not the
exports that produced them) at the paths the running stack actually uses, and fails loud on any orphan.

It is a separate, directly-testable script (`check-cockpit-roster.py`) rather than inline shell,
specifically so the rule can be exercised against fixtures — including one that **reproduces the live
defect verbatim** (a stories roster beside a pt-world menu) and one that refuses an **empty** menu as a
vacuous pass. Live on `demo-1` it now reports *"all 12 cockpit seats resolve in the 35-identity roster."*

The severity is worth restating, because it justifies the check's existence: an orphaned seat does **not**
404. The fake FAPI establishes a session for an **unknown** identity, so the visitor is signed in as
**whoever was last active** — a successful-looking **wrong login**, with no error and no log line.

## D18 addendum — the manifest half of D18's evidence was read off the STALE file

D18 recorded that at iter open `demo-1` held the presenter world, citing two things: the DB (4 story orgs,
591 users) and `cockpit-manifest.json` advertising all four hero trios. **D19 shows the manifest was read
from the authoring copy's stale dir** — a file last written on Aug 11.

**The conclusion is unchanged and the DB evidence is untouched**: `demo-1` did hold the presenter world,
and iter-04's "owed in fact" routing was genuinely stale. But one of the two supporting facts was read
from the wrong file, and it is corrected here rather than left standing — the same failure it documents,
committed while documenting it. *A path is evidence about a claim only if it is the path the system reads.*
