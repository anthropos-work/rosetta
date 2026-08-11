# iter-97 decisions

## `D-M257x-97-1` — N = 20. The gate does not move, and the count going UP is reported as the result

51 sites repaired; N went **13 → 20**. Reported plainly, not framed. The three facts that make it
interpretable, in order: **only 3 of 20 are recurrences** (17 are newly surfaced, so the repair held);
the **platform-drift class fell 7-of-13 → 1-of-20**, which is the class M257x exists to fix; and the
**dominant class inverted** to the corpus's claims about *itself* — wrong-construct citations, stale
currency pins, enumerations short a member.

iter-95 said N was **a floor twice over**. It was. The true N at that tree was never 13, and this
reading's 20 is a floor too — N̂ ≈ 29.3, union recall ≈ 68 %, ~9 estimated unfound.

## `D-M257x-97-2` — the pre-registration failed 4 of 7, and that is the point of tightening it

iter-95 graded 6 of 6 with a `[0,12]` band and booked it as *a warning, not a win*. The bands were
narrowed until they could fail. **Failed: #1 (per-reading [2,7] → 12/13), #2 (union [3,9] → 20),
#3 (zero recurrence → 3). Held: #4 precision, #5 recall, #6 class composition, #7 induced defects.**

The split is informative in a way 6-of-6 was not: **every failure was a magnitude guess and every hold
was a mechanism claim.** The model of *how* the instrument behaves — 93 % precision, ~42 % per-pass
recall, class composition shifting with the repair, ~2 induced defects per repair pass — is good. The
model of *how many defects exist* was optimistic by 2×, because it assumed repairing 51 sites reduces N.
It does not: N measures what a 68 %-recall instrument surfaces, not what is there.

## `D-M257x-97-3` — the claim-twin fence is necessary and MEASURABLY not sufficient: 3 escapes in 51

`claim_twin_guard` is GREEN over all 14 refuted forms while three of them are live in the corpus. It is
not lying — it matches **quoted verbatim** forms, and all three escapes are **paraphrases**:
`backend.md:13` says *"still pending"* where the sweep enumerated `rollback path`; `service_taxonomy.md`
asserts an archive date the retraction is three rows away from; `backend.md:33` restates the skiller
predicate in different words **in the same file** whose `:127` was repaired.

iter-93 already recorded *fencing a document does not fence its paraphrases*. This iter puts a number on
it: **3 of 51, ~6 %**. That is the residual of the predicate-wise method, and it is not closable by
tuning the matcher — a paraphrase fence would have to understand the claim. Recorded as a **known,
measured limit**, not a bug to fix.

## `D-M257x-97-4` — I shipped the defect I had named one iter earlier

`D-M257x-96-5` states: *"a prose repair is a line-number edit, and only half of that is fenced."* The
repair then re-pointed the **three cross-file** citations to a construct it moved and left the **in-file
`:543` self-anchor** standing (`external_services.md:614`), so the corpus cites two different lines for
one construct. And `dependency_map.md:59` — a cell the repair **wrote** — names two `app` refs, which the
same commit forbids twice in its own prose (`platform-migration-status.md:92`, `storage.md:29`).

Naming a hazard in a decision record does not fence it. Both are routed to
`FIX-M257x-iter97-read-union` with the binding condition that inbound-citation re-derivation after a
line-count change must cover **in-file self-anchors**, which is the half that was missed.

## `D-M257x-97-5` — rule 44's own measurement is false, and it was caught by the adversarial seat in BOTH readings

`platform-alignment.md:1236` claims `useCoursebuilder.ts` holds **1,178 NUL bytes**. It holds **1**.
1,178 is the file's line count: `grep -c` counts matching *lines*, and the zsh `$'\x00'` pattern
degenerates to an empty pattern matching every line. The rule's predicate survives (`file(1)` calls it
`data`; both `grep -I` and `git grep` skip it), but **the rule about instruments lying was written with a
lying instrument** — and its printed recipe has the same shape of error, labelling a *file* count as
`hits=` and returning 2 where the rule's own worked example publishes 22.

Not repaired here: a measuring pass may not contain a repair. The standing rule it produces is the third
binding condition on the next repair: **do not write a measurement into the corpus without running the
measuring command exactly as printed and reading its output.**

## `D-M257x-97-6` — the ref-discipline rejection class is now 13 occurrences over four readings, still contributing zero

All 4 rejections were the same class, as in each of the three prior adjudications. It remains the
dominant false-positive generator, is caught every time by adjudicators applying rule 33, and has
contributed **zero** to any graded count in four readings. **Filtered, not fixed** — per standing
instruction, and now with enough data to say the filter is load-bearing rather than incidental.

## `D-M257x-97-7` — a pre-existing rext test failure was found, characterised, and NOT papered over

Installing pytest (absent on this box, so the Python suite had never run here) surfaced
**909 pass / 1 fail**. The failure is **pre-existing** — reproduced with this run's changes reverted.
Cause: `test_claim_twin_guard_iter48_answer_key::test_02` asserts an iter-48 fixture's "green twin" stays
silent, but iter-49's ledger later booked a claim whose refuted form is that fixture's sentence — and
**iter-49's claim was itself REFUTED at iter-52**. `claim_ledger.py` has **no supersession**: an
overturned claim stays armed forever, against the sentence that is now true. The live corpus escapes only
via a waiver.

I attempted the fixture fix, found it cannot be made green while an overturned claim stays armed,
**reverted it**, and routed the finding as `DEF-M257x-iter96-ledger-supersession`. Tuning an answer key
to green is Trap A by name; the design is what is wrong.
