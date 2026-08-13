# iter-84 — decisions

## D-M257x-84-1 — the union is adjudicated at 40/43 UPHELD, with a PER-ANCHOR ledger

93.0 %, against iter-80's 92.1 % on the pre-repair union. The pre-registered floor was ≥ 70 % and the
falsification condition (< 50 %, meaning the post-repair signal is mostly noise) did not fire. **The
instrument did not regress across iter-81's repair.**

`FIX-M257x-iter83-adjudication-has-no-per-anchor-ledger` is **discharged in the same act**: every anchor
carries its own verdict, severity and predicate, so the reach fence now has an `upheld` denominator
(**40**) rather than only a `booked` one. iter-76's gap is not repeated one iteration after naming it.

## D-M257x-84-2 — five of the eleven discharges stand; four are refuted; one is UNSETTLED, not clean

Re-derived by membership, not by re-reading. **P2, P3, P5, P6, P10 stand** — no live member on any swept
surface. **P4, P9, P11 are refuted.** **P1 is recorded UNSETTLED**: clean inside the instrument's 40
files, with ungraded candidates on the `corpus/ops/**` surface (`verification.md:650`, `safety.md:562`,
`snapshot-spec.md:468`), each of which may be true of the *tooling* while false of the *platform*.

**Recording P1 as unsettled rather than clean is the decision.** An unswept surface reported as clean is
the defect this milestone exists to end, and it would have been the easy call — P1 was the dominant
predicate and declaring it discharged would have read as progress.

## D-M257x-84-3 — clause 5's instrument reads 40 of 40; the corpus is larger BY THE CLAUSE'S OWN WORDING

> **This decision was rewritten after a peer review — see `D-M257x-84-6`.** It first read *"the instrument
> reads 40 of 112 published files, and this is a REACH limit"*. That was wrong.

**Clause 5's declared scope is `corpus/services/**` + `corpus/architecture/**`.** Measured:
`git ls-files -- 'corpus/services/*' 'corpus/architecture/*'` → **40**, of which **0** are non-`.md`; the
instrument's `corpus/architecture/*.md` + `corpus/services/*.md` glob is **exactly co-extensive** with the
declared `**` globs (`find … -mindepth 2 -name '*.md'` → 0). **The instrument reads 40 of 40 — complete.**

The corpus as a whole is **90** `.md` (`git ls-files -- 'corpus/*.md'`), of which `corpus/ops/**` is **46**.
That gap is a **SCOPE OBSERVATION**, not a reach defect: clause 5 is narrower than the corpus *because the
clause says so*. Wanting broader coverage would be a **re-cut of clause 5 — not on the table.**

**What survives, and it is still worth having:** live defects exist **outside** clause 5's scope. Of P4's
live members, **one** is inside the declared 40 and **≥16** are in `corpus/ops/**`, `CLAUDE.md` and
`.claude/skills/**`. That is a statement about **corpus quality**, not about the instrument — and it is
where the runnable commands live (`CLAUDE.md:285` was one). It also explains iter-81: the repair inherited
the **read's** 40-file partition, so no seat owned those surfaces (**§5 rule 19**).

**The limiter on what a zero establishes is the ~50 % per-pass RECALL** (iter-83), a within-scope property.
Correctly named, the file count is not a limiter at all.

## D-M257x-84-4 — the seat-ref rule is STATED WRONG, which is why it has failed five times

Occurrences 4 and 5 landed in this iter (`CLAUDE.md:203`, `hiring.md:73`); the escalation condition
declared in `overview.md` has **fired**.

The diagnosis is adjudicator D's and it is better than "seats are careless": *"grade at the ref the claim
names"* is silent on a sentence that asserts **currency**, so seats apply it unevenly — correctly
refusing to grade `hiring.md:73` against a newer checkout, then incorrectly extending the same courtesy
to `graphql-wundergraph.md:13`, which says *"survives"* and *"is now"* under a column headed *"origin
HEAD"*. Amended form:

> **Grade at the ref the claim names — UNLESS the sentence asserts currency, in which case no
> neighbouring pin rescues it.**

Plus adjudicator A's structural half: **the ground-truth table must carry each clone's `origin/main` sha
beside its checkout sha.** A seat handed only the checkout has no way to see it is stale, and that is
mechanically how 3, 4 and 5 happened. Routed `CHECK-M257x-iter84-rule33-currency-amendment` +
`CHECK-M257x-iter84-ground-truth-needs-origin-sha`.

**Not written into §5 in this iter.** The rule text is the frozen briefing's subject matter and amending
it mid-run would change the instrument between readings — the one thing this milestone has protected
since TOK-04. It is routed with its evidence.

## D-M257x-84-5 — adjudicate-before-repair earned itself again, on a single anchor

**`ai_architecture.md:225` is CORRECT against source.** The composited MP4 *is* written to and read from
prod S3 (`chime.go:188`, `:264-266`, `:341-352`, with no delete). The false statement is at
**`corpus/ops/demo/media-substrate-spec.md:33-35`**, where it is load-bearing for a **safety
disposition**.

**Repairing the booked anchor would have broken a true sentence and left the false one standing.** No
reading could have caught this — both seats booked the right *pair* and the wrong *half* — and only an
adjudicator re-deriving from source could separate them. It is also the third independent confirmation
of `D-M257x-84-3`: the real defect is in `corpus/ops/**`, outside the instrument's reach.

## D-M257x-84-6 — I measured the instrument against a scope clause 5 never claimed, and a peer caught it

**Recorded, not quietly fixed** — the second time in this run (`D-M257x-83-9` was the first), and the same
class both times: **a number stated without the command that produces it.**

I wrote that the instrument *"reads 40 of 112 published files"* and classified the difference as a **reach
limitation** that weakens what a zero establishes. Two defects in one sentence:

1. **`112` was never defined.** It is reproducible — `find corpus .claude/skills -name '*.md'` (110) plus
   `CLAUDE.md` and `README.md` — but an undefined denominator in a gate-critical claim is exactly the
   disease this milestone treats. (`git ls-files -- '*.md' | grep -v '^knowledge/plan/'` gives **113**, the
   same set plus `CHANGELOG.md`; the corpus alone is **90**, not 112.)
2. **The classification was wrong, in the direction that matters.** Clause 5's declared scope is 40 files
   and the instrument reads **40 of 40**. Measuring it against the whole tree and reporting a shortfall is
   an **implicit re-cut of clause 5** — the one thing the user has ruled out three times. It is a **scope
   observation**, and it must not be written in a form that reads as an argument for widening.

**Corrected in `membership.md`, `progress.md`, `D-M257x-84-3` and the milestone ledger**, each denominator
now carrying its command. The substantive finding is unchanged and is restated in its honest form: **live
defects exist outside clause 5's scope**, which is a corpus-quality fact and an explanation of iter-81's
partition, not a deficiency of the gate's instrument.

**Why it is worth a decision entry.** I had just written §5 rule 40 about checks that report a state
without measuring it, and `D-M257x-83-9` about hand-counted scalars — and then stated a denominator I had
not defined and drew a gate-relevant conclusion from it. The rule does not get easier to follow because
you wrote it.

## Unchanged routes

`FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) · `FIX-M257x-iter56-assignment-flake`
(**NOT DECIDED**) · `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
`CHECK-M257x-iter77-zsh-modifier` · `CHECK-M257x-iter77-developer-dir` ·
`CHECK-M257x-iter70-studio-room-lines` · `RF-M257x-iter71-run-returns-a-tuple` · RF-2/3/7–14 ·
`CHECK-M257x-iter82-commit-message-narration` (**stays SEPARATE from `CHECK-iter77`**) ·
`CHECK-M257x-iter83-recall-lift-options` · `DEF-M257x-iter80-storage-prod-bucket` (**escalated, held;
`storage.md:55,:154,:181` unchanged, and measured NOT to be part of the 43**).
