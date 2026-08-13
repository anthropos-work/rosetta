# iter-80 — the seventh routed count, and the one that did not collapse

**Shape:** tik (measurement-shaped) + **harden pass 19**.
**Gate:** 4 of 5, unchanged. Clause 5 remains the only open clause.

## What this iteration was asked to settle

Run 51 reached two conclusions that, taken together, looked like they might close off clause 5
entirely: *free prose is not fenceable* (three candidate fences measured 71 % precision / rejected)
and *the corpus should not be restated*. If both stand, 20 of 23 migration claims are permanently
outside the reach of the fence family — so either clause 5 is measured by a different instrument, or
the milestone needs a TOK.

**It is a different instrument, and this was conflated before.**

| | instrument | what it reads now |
|---|---|---|
| clause **3** | `platform_predicate_guard` (G1–G10) + `platform_alignment_guard` | **MET.** All 5 corpus guards GREEN |
| clause **5** | the **graded READ** — 14 blind seats over the 40 files of `corpus/services/**` + `corpus/architecture/**`, instrument frozen at iter-41's briefing, stored as `instrument/briefing-iter76-AS-RUN.md` | **NOT MET — 140 adjudicated blockers** |

`platform_predicate_guard`'s G5 reach (**1 enumerated + 22 free-prose UNREACHED + 1 ref-pinned of
24**, printed on every run) is **not clause 5's denominator and never was**. The fence's job is to
stop a repaired claim from silently regressing; the READ is what measures the gate. Both of run 51's
conclusions stand, and neither one closes clause 5 off — because neither one is about clause 5's
instrument.

## The measurement

`FIX-M257x-iter76-read-union` adjudicated **in full** — 152 booked blockers, 4 parallel adjudicators,
each re-deriving from the platform/service clones rather than from any prior verdict.

**152 booked → 140 UPHELD · 12 REJECTED · 0 UNSETTLED. 92.1 %.**

**This is the seventh routed count in this milestone and the first that did not collapse.** The six
before it: **64→5, 23→1, 21→0, 92→0, 4→3, 145→3**. That track record is exactly why iter-76 routed
its 77/75 instead of repairing it, and why "adjudicate before repairing" was made binding. The prior
was *"it will collapse."* It did not.

The prior was well-earned. It now has a counter-example, and the counter-example is the one that
decides the gate. **The instrument was not crying wolf.** iter-76's own hedge — that the *"past the
end of a 271-line file"* class made ~150 *"an upper bound, not a work item"* — is now measured: that
systematic false-positive class accounts for **4 of the 12 rejections, not most of the 152**.

Full verdict table, the 12 rejections by mechanism, and the 11-predicate repair unit:
`iter-76/adjudication.md`.

## The path to clause 5

Concrete, bounded, and **unblocked for the first time**. The repair was routed under three binding
conditions and all three are now discharged:

1. **adjudicate before repairing** — done this iteration (152/152).
2. **repair by PREDICATE, not by claim** (`D-M257x-59-1`) — the 140 dedupe to **11 predicates**, and
   all four adjudicators converged on substantially the same list independently. P1 alone
   (*"the cms/jobsimulation/roadrunner containers still start"*) is ~47 findings across 6+ files.
3. **not closed until the G5/G2 reach hole is closed** — discharged by iters 77–79.

So: **repair the 11 predicates, then re-read.** Not a TOK. The three-run stretch of instrument
defect-finding was *on* the critical path — it was the milestone's own booked precondition — and it
is now finished; continuing it would not be.

**Honest sizing:** ~140 findings over ~20 files is not one iteration. But it is bounded work against
a derived, deduplicated list, and the re-read that grades it already exists and is frozen.

## Harden pass 19 — RUN (deferred three times; not a fourth)

Scope **iter-69 … iter-79** (11 tiks; pass 18 closed at iter-68). Dimension: the routed
`CHECK-M257x-iter79-three-valued-discriminators`, now **CLOSED**.

Swept every subprocess-derived discriminator in the `stack-core` guard family. Most were already
three-valued and said so in their own prose — which is the pass's main finding, and a good one:
iters 77–79 taught this module the rule and the module mostly learned it.

**One sibling had not.** `_reads_at_ref` published two outcomes where `git grep` has three: rc 1
("no match", a real answer) and rc ≥ 2 ("could not look") both returned the same empty dict. A clone
whose object store cannot be read therefore reported **"the consumer side reads no `*_RPC_ADDR`"**
with `res.reach["app_consumer_side"] == "measured"` — the `|| echo 0` signature M257 opened on, one
level down, inside the guard built to end it, in the branch *adjacent to the comment that states the
rule*. Fifth occurrence this milestone of *"the author of a newly written rule violated it while
writing it."*

**Reachability proven, not argued, and the upstream guard does not cover it:** `app_ref_sha`'s
`rev-parse --verify` answers from the **commit** object and correctly rejects a shallow clone's
missing sha — but `git grep` needs the **tree**, so an unreadable/corrupt tree object resolves at
`rev-parse` and fails grep with rc 128. Built and reproduced (`chmod 000` on the loose tree object).

**3 mutants, 3 kills, 3 distinct signatures**, collected before running: the pre-fix collapse
(`'unmeasured' != 'measured'` — the defect named), the over-correction (`{} != None` — caught by the
rc-1 **control**), and a provenance-silencing mutant (pass 18's *"a reporting path with no mutant is
a docstring"*). **157 → 160 tests.**

## Also raised — an operational hazard, not a documentation defect

`storage.md:55,:154,:181` say local private storage is sandboxed to `/tmp`.
`docker-compose.yml:82` @ `0dab54d` sets `STORAGE_S3_BUCKET=production-storage2024…` **on `backend`**,
read straight into `NewManager` (`app 9d00a313 main.go:463→471`). **Local private writes land in a
production bucket**, while the doc reassures the reader about exactly that manager.

Routed `DEF-M257x-iter80-storage-prod-bucket`, severity **high**, escalated to the user. Not actioned
here — the platform half is a platform-repo question and this milestone is zero-platform-edit.

## State

5 corpus guards GREEN · `platform_predicate_guard` reports *"app consumer side measured @
origin/main@7177374"* (the new provenance path, live) · `stack-core` **160** tests in the predicate
suite, full-section run's only non-green is the known perishable iter-48 fixture (1F) · pin stays
`fast-build-m257x-iter-67` (guard code only; no bring-up consumes it) · zero platform-repo edits.

**Gate 4 of 5, unchanged.**
