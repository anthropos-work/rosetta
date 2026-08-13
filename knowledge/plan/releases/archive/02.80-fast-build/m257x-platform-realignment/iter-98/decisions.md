# iter-98 decisions

## `D-M257x-98-1` — twin expansion stays INSIDE the repair; the new axis is PARAPHRASE

The run brief asked whether to run twin expansion **ahead of** the read as a separate pass. **Answered no,
on measurement rather than preference.**

iter-96 already expanded string-twins inside repair (13 anchors → 51 sites). iter-97 then measured the
escape from that pass: **3 of 51, and all three paraphrases** — `claim_twin_guard` was GREEN over all 14
refuted forms while all three were live, because it matches **quoted verbatim forms**. So the string-twin
axis is near zero escape and a separate ahead-of-read pass would re-measure a solved class; the 6 % escape
is on the paraphrase axis, and that is where the extra sweep went (`git grep -niE` over meaning-alternations
per predicate, not per string).

**It paid, three times, in ways no string sweep of the booked anchors reaches:**
- `messenger.md:108` — an unbooked third site of P5 with a *different* wrong range.
- `frontend-tier.md:425-450` — the same document holding **both** readings of the academy auth model, its
  own retraction at `:48`/`:84` and the superseded model at `:427`.
- **P10 turned out to be bigger than booked**: not "the cockpit no longer sets the cookie" (false — two live
  paths) but *the demo academy's entire auth model had inverted* to real Clerkenstein keys, with the bypass
  removed from the launch env and fenced by two tests. The booked anchor named a symptom; the sweep found
  the change.

## `D-M257x-98-2` — a false claim ABOUT OUR OWN CODE is withdrawn, never made true

`safety.md:207` (and its unbooked twin `seeding-spec.md:102`) asserted that `s3-private` had been removed
from the `PerStackIsolated` registry. **It has not been** — `stack-seeding/isolation/isolation.go:106` still
registers it. Two ways to make the corpus true: change the doc, or change the code.

**Changing the code was refused.** Re-classing that store *is* the disposition of
`DEF-M257x-iter80-storage-prod-bucket` — an open, escalated, explicitly-undecided item belonging to the user.
A repair iter that silently resolved it would be deciding a user's question by editing a Go slice. So the
assertion is withdrawn and **the disagreement between doc and registry is now stated in both documents**,
which is the honest state until the user rules.

This is the general form: *when a doc and the code disagree, fixing the doc is always available; fixing the
code is only available when the change is yours to make.*

## `D-M257x-98-3` — the answer-key GREEN fixture was stale, and fixing it is not tuning to green

Arming the iter-98 ledger turned `tests/test_claim_twin_guard.py::test_02` RED: the **green** fixture
`claim_twin/green/17.md` — whose contract is *"the same neighbourhood with the refuted line removed"* —
still contained *"no other platform repo references roadrunner at all"*, which P7 has now refuted.

The fixture was green **with respect to iter-41's blocker only** (the *"still in `repos.yml` (1 of 9)"*
claim, which lives in its later lines). A claim refuted 57 iters later was never removed because nobody knew
it was false. **The fixture violated its own stated contract**, so it was repaired to satisfy it.

**Why this is not the anti-pattern the carried instruction warns about.** The banned move is loosening an
answer key so a fence stops detecting. Here:
- `test_01` — *all 18 known-bad sites must still fire* — is **untouched and still passes**; the RED fixture
  was not modified, and its blocker is in the lines below the edit.
- The edited line is **incidental context** in the red fixture, not its adjudicated claim.
- Detection power is therefore provably unchanged: 20/20 pass, including the positive control.

The alternative — weakening iter-98's P7 quote so it stops matching — *would* have been tuning, because it
would have hidden a true refutation to keep a stale fixture quiet. **The stale artifact was the fixture, and
the fixture is what moved.**

## `D-M257x-98-4` — a retraction must use the corpus's own retraction vocabulary, or the fence cannot see it

The P7 waiver (`claim_twin_waivers.json`, entry 9) did not take on its first run. The waiver is deliberately
**half a key**: `_looks_retracted` must independently find a retraction marker within 320 chars of the match.
The repair had written *"are both **FALSE**"* — and `RETRACTION_MARKERS` contains `"is false"` and
`"was false"` but **not** `"are both false"`.

Two ways to close it, and only one is safe. Widening `RETRACTION_MARKERS` is, in the guard's own recorded
words, *"the direction that can hollow a fence out"*. **So the prose moved, not the fence:** the sentence now
reads *"is RETRACTED — each is false"*, using a marker the guard already knows.

**Generalised:** a retraction that a fence cannot recognise is, to every automated reader, an unretracted
falsehood. Write retractions in the vocabulary the fence enumerates — and when tempted to widen the
vocabulary instead, note that you are proposing to make the fence accept a form you invented thirty seconds
ago.

## `D-M257x-98-5` — the induced-citation class was caught IN-ITER, five times, including one self-inflicted

Binding condition 2 was executed as a **post-edit** step, not a pre-edit plan, and that ordering is the whole
finding. Five citations moved; all five were caught before commit (table in the ledger). The sharpest case:
the P15 fix to `graphql-wundergraph.md` **added a line above the very target it was citing**, so the number
being written (`:192`) was invalidated by the act of writing it, and re-derivation returned `:193`.

**A citation repair is itself a line-count edit.** The re-derivation must therefore run *after* the edit and
be re-run after *every* edit, including the citation repairs themselves. iter-96 named this class
(`D-M257x-96-5`) and shipped it anyway because it re-derived from the pre-edit reading.
