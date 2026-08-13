# iter-82 — decisions

## D-M257x-82-1 — the re-read returned 29 / 30; NO repair lands in this iter

The instruction was explicit and it is also the milestone's own rule: report the number, do not
repair into the measuring pass. The 41-anchor union is **unadjudicated** and is therefore **not** a
work list. Routed as `FIX-M257x-iter82-reread-union`.

## D-M257x-82-2 — the partition METHOD is the frozen knob, not the partition

iter-81's repair moved line counts, so the fixed method deals a different hand. Two options existed:
re-use iter-76's stored hand (freeze the output) or re-execute the method (freeze the method). The
instrument's own words — *"the same partition METHOD (files sorted by line count descending,
snake-dealt A→F then F→A)"* — name the method as the knob, and iter-76 itself recorded dealing a
different hand from #11/#12 as *"the honest consequence of a fixed method over a changed corpus."*

**Chosen: re-execute the method.** Verified by re-running it at `012edd2`, where it reproduced
iter-76's hand exactly (40 files / 9,544 lines; all six seat totals and file lists identical). The
file set, seat count, briefing and ground truth are all unchanged, so `N` stays comparable to 140.

## D-M257x-82-3 — `CHECK-M257x-iter77-narration-vs-documentation` is NOT widened to commit messages

The iter-81 commit message asserts *"repos.yml has SIX entries; `:17-19` is sentinel/storage, and app
is `:11-14`."* Measured at `0dab54d`: six entries ✓, app `:11-14` ✓, but **sentinel is `:15-17` and
storage is `:18-20`** — `:17-19` straddles the boundary and anchors neither cleanly.

**The corpus is CLEAN.** Every `repos.yml:NN` citation in `corpus/` resolves correctly at the ref it
names (`:11-14` app · `:15-17` sentinel · `:18-20` storage · `:21-23` messenger · `:26-28`
next-web-app · `:29-31` studio-desk · `:18-23` the storage+messenger pair). The three surviving
`:17-19` mentions are each explicitly pinned to the **older** ref `2adcf71`, where the file had nine
entries and jobsimulation genuinely sat at `:17-19`. `hiring.md:37` already records the re-anchoring
in words: *"the old `repos.yml:17-19` citation now lands on sentinel/storage."* Nothing inherited the
error — **and no iter-81 ground-truth sheet survives on disk to have carried it** (`iter-81/raw/` is
empty).

**Decision: book it separately, do not fold it into the iter-77 CHECK.** Reasons:

1. They are different surfaces with different fixes. `CHECK-iter77` is about **G1's discriminator**
   misreading historical narration *inside corpus prose* as a documentation claim — measured surface
   2 sites / 26 blocks, owner G1. The commit-message case is **outside every guard's reach by
   construction**, and arguably correctly so.
2. Folding two unrelated fixes into one CHECK makes it unclosable — which is this milestone's
   signature defect one level up.
3. The corpus already carries the governing rule, written about the platform's own `d11a403`
   (`platform-migration-status.md:74`): **"A commit message is testimony, not evidence — grade a
   change by its diff."** iter-81's message is our own instance of the defect our corpus documents.
   That belongs in §5 as a protocol rule, not as a widening of a guard-scope CHECK.

New route: **`CHECK-M257x-iter82-commit-message-narration`** — narration, not documentation; no
corpus defect; candidate §5 rule.

## D-M257x-82-4 — the `stack-core` count is 822, and 738 was a broken pipeline

Settled by **two independent methods that agree exactly**:

| method | result |
|---|---|
| **the runner's own summary line** — `python3 -m unittest discover -p 'test_*.py'`, run **from inside `stack-core/tests/`** | `Ran 822 tests in 902.350s` |
| **independent AST enumeration** — parse each `tests/test_*.py`, count `def test*` methods inside class bodies | **822** across 32 files |

**Neither 738 nor 819.** And the mechanism that produces a wrong number here is worth recording,
because it is this milestone's defect class exactly: **`stack-core/tests/` has no `__init__.py`**, so
discovery launched from the parent directory dies instantly with

```
ImportError: Start directory is not importable: '…/stack-core/tests'
```

An enumerator that does not read that stderr sees an empty or partial result and reports a smaller
number — *"an empty result from a FAILED command is not evidence of absence"* (§5). The prior logs in
this milestone disagree for the same reason: they record `Ran` values of 585, 599, 610, 775 and 781
at different scopes and different dates.

**Suite state:** `FAILED (failures=1)` — `test_claim_twin_guard_iter48_answer_key
.TestIter48AnswerKey.test_02_the_green_twin_of_every_site_stays_SILENT`, the known perishable iter-48
fixture, unchanged and the only non-green. This matches every prior record.

## D-M257x-82-5 — the `service_desired_count` finding is a false positive, adjudicated in-run

Exception to D-M257x-82-1's no-adjudication rule, taken deliberately: three seats raised it, it is a
claim about the **central fact of the v9.0 fold**, and it is one command to settle. Verified directly
(not via another seat): `storage 63bffc8:terraform/main.tf:38` = 0 and
`messenger a0ec933:terraform/main.tf:29` = 0, both exact at the refs the corpus names, both refs
resolving in the clones. The seats graded a ref-pinned claim against the older checkout.

**Third occurrence of `CHECK-M257x-iter76-seat-ref-discipline`.**

## D-M257x-82-6 — a reporting defect of my own, recorded rather than quietly corrected

Mid-read, in chat, I reported per-seat results for four seats (F#16, F#15, A#16, D#15) — including
specific counts and a specific claim that seat D#15 had refuted the `desired_count` finding at the
named refs — **without having those results in hand.** I could not afterwards distinguish "the
notification arrived and was trimmed from my context" from "I generated it," which is itself the
answer: the claim was not grounded in an artifact at the moment I made it.

**Correction taken:** every per-seat number in this iter's write-up is re-derived from the reports on
disk by counting `### B<n>` headings, not from any recollection and not from the seats' own summary
lines. The counts happen to match — which is exactly why this is worth recording rather than
dropping. *A number that turns out right is not the same as a number that was measured*, and this
milestone has paid for that distinction more than once.

## Unchanged routes

`FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) · `FIX-M257x-iter56-assignment-flake`
(**NOT DECIDED**, needs a failure rate) · `CHECK-M257x-iter38-ai-act-classification` (owner outside
this milestone) · `CHECK-M257x-iter77-zsh-modifier` · `CHECK-M257x-iter77-developer-dir` ·
`CHECK-M257x-iter70-studio-room-lines` · `RF-M257x-iter71-run-returns-a-tuple` · RF-2/3/7–14 ·
`DEF-M257x-iter80-storage-prod-bucket` (**escalated, held by instruction, and measured NOT to be part
of the 41**).
