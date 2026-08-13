# iter-49 — decisions

## D-M257x-49-1 — the leak fence's blind spot was a QUESTION, not a threshold, and the fix is a second question

`D-M257x-48-9` measured the obvious move and refuted it: lowering `repair_leak_guard`'s K to 5 or 6 bought
**two false positives and still missed** the `16 → 23` twin. K was never the binding constraint.

The cause is the **added-text suppression**, which exists so a MOVE is not a leak. When a repair changes only
a value, the rest of the sentence is present in both the removed and the added text, so every shingle that
could locate the twin is suppressed as *"the commit rewrote this"* — and the one token that did change is far
below any distinctiveness floor.

`value_change_guard.py` therefore asks a **different question**: not *"does the old FORM still stand?"* but
*"does the old VALUE still stand in a sentence that otherwise says the same thing?"* It word-diffs each hunk
(removed against added, kept **paired** — the pairing `repair_leak_guard.read_diff` deliberately discards),
keeps the short `replace` opcodes, and searches for `(old value + surviving context)` in order within a
window.

**Watched RED before it was trusted**, on rosetta `301d61a` — a real, already-committed incomplete repair
whose answer key was written by seven auditors who had never seen this code. It reports 5 sites, of which
**three were independently adjudicated blockers by two different audits**: iter-48 #3
(`architecture_overview.md:293`, the motivating one), iter-47 #5 (`external_services.md:139`) and iter-48 #9
(`external_services.md:561`, found from a completely different angle — the `none of them → neither`
correction).

**The gap is asserted, not assumed.** `test_the_verbatim_fence_really_is_blind_to_it` runs
`repair_leak_guard` on the same commit and requires it NOT to report the site. If that ever changes, the
suite says so rather than leaving a redundant fence in place.

## D-M257x-49-2 — the fixture's 2 false positives were KEPT, and the tuning that would remove them was measured first

Both false positives (`safety.md:279`, `storage.md:32`) have one cause: `301d61a` also replaced the function
word `in` with `and`, and *"in the platform compose"* is ordinary English appearing in sentences about a
different subject.

A **stopword filter** would remove both. It would also remove `external_services.md:139` — **a real
adjudicated blocker** (iter-47 #5) found through that same function-word change. So the trade is one proven
true positive for two false ones, and it was measured rather than missed.

Following the discipline `repair_leak_guard` established one iteration earlier, the fence keeps them and
`test_the_false_positive_count_is_pinned` **pins the set in both directions**: growth is noise the fence will
be disabled for, and shrinkage to zero is usually a knob tuned to the answer key.

## D-M257x-49-3 — `--audit-commit` writes NOTHING, and that is what makes it safe

`D-M257x-48-12` recorded the ratchet hard-blocking iter-48's audit commit, the `--accept` escape correctly
refusing (a baseline may lower, never raise), and the commit landing with `--no-verify`.

The obvious repair — let `--accept` raise for an audit — is the wrong one: **the monotonicity IS the
contract**, and a mode that writes the baseline is `--accept` with a friendlier name. So `--audit-commit` is
a **hook-time classifier that touches nothing**. It admits a new site only on a signature no repair commit
can produce, prints it, and returns 0 — while the sites **stay RED in the suite**, the durable vehicle, until
they are repaired.

`TheBaselineIsNeverTouched` asserts this on **bytes**, before and after, for both the baseline and the waiver
file — not on intent. And `test_accept_still_refuses_to_raise` pins that the older escape still refuses, so a
new mode cannot quietly relax the old one.

## D-M257x-49-4 — the anti-laundering key is PER-SITE file scoping, and the test that proved it first asserted something false

The signature: a site is *audit-introduced* iff **(a)** its claim's ledger row is a line **this commit
added**, under `knowledge/**`, **and (b)** the site's own published file is one this commit **did not touch**.

Condition (b) is the anti-laundering key. A repair edits the file it repairs, so any site it leaves standing
there is refused — which is exactly the induced-defect class the fence exists for. To launder a repair you
would have to not edit the file you were repairing and write a genuine new refutation of the claim you were
leaving standing. That is an audit.

Condition (a) is **line-level on purpose**. A file-level test (*"the ledger was touched"*) would let a repair
append one unrelated line to any ledger and thereby admit every site in the tree.
`test_touching_the_ledger_FILE_is_not_enough` pins the difference.

**And the test battery refuted its own first assertion, which is worth recording.** It initially asserted
that a pure audit commit touches no published file. **False:** iter-48's audit commit edited
`corpus/ops/platform-alignment.md` to record §5 rule 22 — protocol-doc updates are part of an iter's close by
construction. Had the mode been built on a whole-commit *"did this touch `corpus/`"* test, it would have
refused all 18 sites and been useless. The per-site scoping is what makes it work, and the false assertion is
what surfaced that.

## D-M257x-49-5 — the post-condition caught the repairer in the act, twice, before the commit

This is the first iteration in which the fences caught **this iteration's own repair** rather than the next
audit finding it a pass later. On the first post-repair run:

| fence | finding |
|---|---|
| `anchor_construct_guard` (via the ratchet) | `platform-alignment.md:490` — the anchor `hiring.md:86` had drifted to a blank line, because this repair's `hiring.md` edits inserted 7 lines above it |
| `repair_leak_guard` | `ai-readiness.md:426-427` still published *"the live-recompute never completes"*, the very claim this repair had just fenced in `stories-spec.md` |

Both are **induced defects of this repair**, both were repaired before the commit, and neither would have
been visible to a hand sweep — the anchor drift is arithmetic, and the leak sits in a file the repair had no
reason to open. TOK-02 step 2's premise (*"the induced class cannot survive the commit"*) is doing the work
it was built for.

## D-M257x-49-6 — the new fence found a THIRD site of a claim the verbatim fence had already passed GREEN

After the two repairs above, `repair_leak_guard` reported GREEN. `value_change_guard` — run on the same diff,
minutes later — reported `seeding-spec.md:497`, a **third** site of the never-completes claim, reached
through the `completes → completed` tense change that the verbatim fence's added-text suppression hides.

Adjudicated: it is a **legitimate retraction** — the passage quotes the refuted wording in order to correct
it (*"⚠️ M219 FALSIFIED M51's headline strategy claim … The live recompute completes in 2.09 s"*), and
iter-48's own ledger cites this passage as where the refutation is already recorded. So it is a false
positive, and it is **acknowledged in the waiver file with its reason**, not tuned away.

The fence was right to raise it. Two fences asking two questions of one diff produced one real finding the
other could not see, which is the argument for building the second one.

## D-M257x-49-7 — widening a waiver key is the dangerous direction, so the answer keys were re-run

The waiver above could not apply, because `value_change_guard`'s second key (`looks_retracted`, shared with
`claim_twin_guard`) did not recognise this corpus's own retraction vocabulary: `seeding-spec.md` says
**"FALSIFIED"**, and neither `falsified` nor `refuted` was in `RETRACTION_MARKERS`.

Widening a waiver key is precisely the edit that can hollow a fence out, so it was **not** done on the
strength of the words looking right. Both were added and **all three answer-key suites re-run** — iter-41's
18, iter-47's 12, iter-48's 18 — and every known blocker is **still detected**. The widening is recorded in
the marker list itself, with that verification named.

**Second key, and why it is not Trap A:** the waiver is only half a key. `looks_retracted` is recomputed on
every run, so deleting the retraction and leaving the old value standing makes the waiver **stop applying**
with the waiver file unchanged (§8 rule 3). `test_a_waiver_DECAYS_when_the_retraction_is_deleted` pins it —
and that test exists because the first version of the waiver test asserted only the path-AND-form half and
passed while the decay property was absent.

## D-M257x-49-8 — the reporting paths were deleted and the suite was watched failing

Per the protocol's standing instruction (*a fence's self-reporting is as likely to be fictional as the claims
it checks*), each new fence's reporting path was mechanically emptied and the suite re-run:

- `value_change_guard.render_report`'s site list → **2 tests failed**
- `repair_postcondition`'s `--audit-commit` admitted-site list → **1 test failed**

Both were then restored and the suites re-confirmed green. Recent precedent made this non-optional: a
mutation battery returned three surviving mutants on first run against a suite that looked complete, and
`repair_postcondition --reason` once overwrote its own audit trail.

## D-M257x-49-9 — five of this iteration's own tests were invalidated by this iteration's own repair

`--audit-commit`'s five end-to-end tests were driven against the LIVE corpus at rosetta `2fc633a`, the
commit `D-M257x-48-12` records as blocked. All five passed. The same iteration then repaired the 12
blockers, the ratchet fell to **0 new sites**, and all five failed **in the same session** — with nothing
RED there is no admission to make, so `main()` returns before the audit path is reached.

**`corpus/ops/platform-alignment.md` §8 rule 7 already forbids this**, in almost these words, and was
written by this milestone at iter-45: *"if you cannot state why an assertion will still be true after the
defect is fixed, it belongs in the fixture."* It was violated by the author of the fence it protects.
This is the **eighth consecutive iteration** in which a rule's own author broke it while writing the thing
the rule governs.

**The fix is a hermetic fixture, not a skip.** A skip is a check that does not run, and §5 rule 8 is
explicit that a check which skips reads exactly like one that passes. `TheModeWorksEndToEnd` now builds a
temp git repo — a corpus file carrying the claim, an audit ledger adjudicating it, and two
differently-shaped HEAD commits — and drives the whole pipeline (`commit_line_map` → `claim_twin_guard` →
`grade_audit_commit` → the report) through the real CLI. It runs in **2.6 s**, never goes stale, and
carries its own positive control: `test_the_fixture_really_does_go_RED_without_the_flag` fails loudly if
the fixture stops being RED, because every other assertion in the class would then be vacuous.

**What the live-tree run keeps.** It is still the evidence that the mode works on the real blocked commit,
and it is reported in `D-M257x-49-3`. It simply cannot be the durable assertion.

## D-M257x-49-10 — the reading refuted the pre-registration, and the fences are not why

Pre-registered in `overview.md` before any report was read: **6 blockers, 2 induced / 4 pre-existing.**
Measured: **14, split 7 / 7.** Refuted in every term, and hardest on the one this iteration was built to
move.

**The fences are not the failure.** Both closed the gap they were named for, both were watched RED first,
and the commit-time post-condition caught this repair **twice before the commit** (`D-M257x-49-5`). The
cause is that the seven induced findings partition into three classes and **none is mechanically
reachable**: **paraphrase leak** (3 — the twin says it in different words, sharing no token run: the limit
`D-M257x-48-4` pinned), **overshoot in new text** (3 — no old form to leak, no value to diff; the prose
did not exist before the commit), and **wrong mechanism correctly cited** (1 — semantic).

> **So TOK-02 step 2's premise is now true of a class that has stopped being the majority.** Mechanising
> the mechanical half did not lower the total; it changed what the remainder is made of.

**The consequence for the protocol:** §5 rule 21's *classify-the-residual-by-cheapest-catching-instrument*
must be re-run **after each instrument lands**, not once at strategy time. A classification is a snapshot
of a distribution the instruments themselves change.
