# iter-86 — progress

**Type:** tik, under `TOK-05`. **Gate 4 of 5, unchanged.**

## The headline, and it is about the instrument stack, not the corpus

**`platform_predicate_guard` had never been green.** It was authored at iter-60 to fence the profile
predicate, and it exited 1 continuously from then until this iter — through 25 iterations, across every
record that says *"5 corpus guards GREEN"* and *"6 corpus guards exit 0"*, including the one that used
that sentence to declare **gate clause 3 MET**.

The reason it stayed RED is worth more than the fact: **G1 could not reach zero on a correct corpus.**
Its negation discriminator read only the text *before* the noun phrase, so the corpus's own correcting
sentence — *"The `graphql` profile **is gone** too"* — read to the guard as a fresh claim that the
profile lives. A fence whose floor is above zero on a correct tree is a fence that gets quietly un-run,
and it was.

That is the milestone's own disease at the top of its own instrument stack, and iter-83 diagnosed it one
layer short. The registry kind-filter is real — `repair_postcondition.py` derives its fence set from
`*_guard.py` on disk and then selects only `postcondition` kinds, hiding 11 standalone guards. But it
does not explain what actually happened, which the iter-86 census measures:

- iters **77–79** captured a genuine transcript of **one** guard, beside a bare *"5 corpus guards GREEN"*
  whose members are never listed anywhere.
- iter-**80** used that unenumerated 5 to declare **clause 3 MET**, naming 2 of them.
- The count moved **5 → 6** between iter-80 and iter-83 with **no record of what joined**. The only
  artifact that enumerates the six says *5* in its subject and lists *6* in its body.
- iters **83 · 84 · 85** assert *"6 corpus guards exit 0 at open and at close"* with **no captured
  output** — and re-measurement with the guard version each iter declared, against the tree each closed
  at, at the platform ref each declared, gives **rc=1 at every one of those six points**. iter-85's own
  repair took the site count 2 → **3**.
- **9 of the 15 guards are covered by no green claim anywhere in the window**, including
  `value_change_guard` (absent from the entire milestone outside iter-49) and **`derived_value_guard`,
  which is a `postcondition` guard** — so the kind-filter cannot explain it.

> **§2's deleted 4-tuple did return. Not as a runner's hardcoded list — `repair_postcondition.py`'s
> docstring guards against exactly that — but as a human's remembered list, which is worse, because a
> tuple in source can at least be diffed.**
>
> **The earlier `OK`s were real.** The claim did not become false when the guard grew assertions. It
> became false at the iteration where *"guards green"* stopped being a pasted transcript and became a
> sentence.

## What landed

**1. `guard_family.py`** (rext `stack-core/`) — one command that runs the family and **names every
member**. Census **derived** from `*_guard.py` on disk; invocation map **declared** (guards answer
genuinely different questions — tree-state · commit-scoped · needs-a-ledger — and no honest invocation
derives from a filename), and **reconciled against the census in BOTH directions**: a guard on disk with
no entry exits 2 naming itself; an entry naming a guard not on disk exits 2 too. It also refuses to read
a guard's own *"CANNOT RUN … Nothing was checked; this is not GREEN"* as a pass. **+9 tests.**

**First full-family run: 16 members · 14 GREEN · 2 RED · 0 could-not-check · 0 not-run.** Both REDs were
invisible to every green claim this milestone has made.

**2. Both REDs adjudicated and repaired, each on the correct side of the line** (`D-M257x-86-3`):

- **2 of the 3 G1 sites were a CORPUS defect** — `platform-alignment.md` rule 40 **quoted the false
  sentence verbatim** as its worked example, publishing a live-reading copy of the claim it exists to
  kill. Second occurrence (`:1305` was iter-84's). Repaired in the corpus, not waived: `CLAUDE.md`
  already states the governing rule — *"no retired token is spelled here in runnable form"* — and rule
  40 is the passage that teaches it. **A worked example does not need a working copy of the defect.**
- **1 was a GUARD defect** — repaired in the guard, because contorting a correct English sentence to
  dodge a regex taxes every correct sentence written after it. `_NEGATED_AFTER` completes a
  discriminator that was half-implemented. The predicate list denies **existence only** (*gone /
  removed / renamed / retired / dropped / decommissioned / no longer / does not exist*), never a bare
  *"is not"* — because *"the `storage-legacy` profile is **not** started by default"* denies a
  profile's DEFAULTNESS, not the profile. **Both false claims the guard was holding open survive the
  change and stay RED**, which is the evidence the rule was fitted to English and not to the answer key.
  **3 mutants · 3 kills · 3 distinct signatures**; the over-correction mutant is killed by the CONTROL
  tests, which is why they exist. **+7 tests** (160 → 167).
- `value_change_guard`'s RED was `tailscale-serve.md:463`, a real unfinished repair from iter-85's own
  commit — repaired in the P4 sweep.

**3. The seat-ref escalation settled with two measurements, not a preference** (`D-M257x-86-2`) — see
the decision. Short form: **the ground-truth sheet is not the instrument**, because it *already varies
between the readings being compared*; and the class has contributed **zero to the graded count, 5 times
out of 5**, because adjudication was already removing it. Adopted, with the **raw** series declared
discontinuous at iter-86 and the **adjudicated** series untouched. The §5 rule-33 amendment stays
**routed, not written** — `TOK-04` protects the frozen briefing's subject between readings.

**4. The repair.** 30 enumerated predicate rows + the P4 membership sweep, seven disjoint file-partitioned
packets, each **re-deriving every correction against the clones before editing**.

## Adjudicate-before-repair earned itself again, five times

Every one of these was a proposed correction that measurement refused:

| what the ledger said | what the clone said |
|---|---|
| Q1 B2: re-anchor to `:184-190` | those are the `return` and its closing brace; the citation's own convention is the five id lines, `:185-189`. Re-anchoring would have **silently changed what the citation denotes** |
| Q1 B11 + B12: two sites at `:217`/`:218` | **one** site at `:219`. `grep -rn ":1051" corpus/` returns exactly one hit — a double-count with off-by-two anchors |
| Q1 B5 + B6: "same construct, same fix" | **one** site, not two — `SHOW_SECONDARY_TABS` appears once in the file. The quoted source text was also wrong (the declaration carries a type annotation) |
| Q5: `chime.go:188`, `:264-266`, `:341-352` | none of those land on S3 code at this checkout. The **conclusion** held — the MP4 *is* in prod S3 — but every anchor supporting it had to be re-derived |
| Q4 B23: "a mounted 7-locale dropdown" | **two** switchers, not one: a genuine 2-way EN↔IT toggle on the public storefront *and* a 7-locale dropdown in the authed shell. And it was **wrong when written**, not stale — the dropdown has been mounted since 2026-05-05 |

## Findings the repair produced that no ledger predicted

- 🔴 **`platform_repo.md:103–107` and `staging-bringup.md:152` contradict `repos.yml`** — they list
  `cms`/`jobsimulation`/`roadrunner` as entries and show a clone tree containing `graphql-wundergraph/`.
  Six entries exist. This is the `platform_alignment_guard` predicate, not P4 — **routed**, not repaired,
  because it is a different subject and half-repairing it is §5 rule 19 backwards.
- 🔴 **`CLAUDE.md:463`** documents `make up PROFILE=studio-desk`, which **exits 1** —
  *"service studio-desk depends on undefined service backend: invalid compose project"*. A runnable
  command that cannot work, contradicting the profile table 140 lines above it in the same file. Same
  *class* as the dead-token defect but not a P4 member (the token is legal), so G1 cannot see it.
  **Routed.**
- **The rext secret-DNA artifact itself hardcodes `"profile": "graphql"`** (`secret-dna.json`). It is
  **inert** — never resolved against compose, only printed to operators in the catalog banner — and it
  is disclosed in `secrets-spec.md` rather than silently corrected, because re-labelling it is a rext
  change with its own blast radius. **Routed.**
- **The gene scalar had THREE published values** (56 / 61 / 64) across six files. Derived from the
  artifact: **64**, version `fast-build-m256`. The delta is traceable to one commit (`ef41a19`, M256
  iter-21, +3 platform genes) — **the corpus was exactly one release stale, not wrong**. Settled to 64
  in all six.
- **A second induced defect, caught at commit time by the family runner**: a repair re-anchored
  `app/main.go:212` — correct at the 60-behind checkout, a closing brace at `origin/main`. Caught by
  `anchor_construct_guard`, repaired to `:216` with both refs named. **This is the seat-ref class
  arriving inside the repair rather than inside a reading**, and it is the first time it has been caught
  by an instrument instead of by an adjudicator.

## Reach: **raw 40/46 = 87.0 %**, and the six misses were all defects in MY OWN LEDGER

Both numbers, because only reporting the second would be the thing this milestone exists to stop.

| | |
|---|---|
| **raw reach** | **40 of 46 = 87.0 %** (iter-85: 11/11 = 100 %) |
| after adjudicating the 6 | 46/46, recorded as **written dispositions** in `rext stack-core/repair_reach_waivers.json` — never a silent skip; the guard refuses an empty reason |
| findings actually left unrepaired | **0** |

**The post-condition graded the LEDGER, not the repair — and it was right to.** All six misses split into
two causes, neither of them a repair that missed:

- **4 were ALREADY DISCHARGED before this iter opened** (`dev-up/SKILL.md:74`, `:147`, `:175`,
  `dev-up/reference.md:38`). I built the ledger from iter-84's `membership.md` **without re-checking it
  against iter-85's repair**, which had already fixed all four. That is **§5 rule 32 — re-derive a
  hand-off's numbers — and I failed to apply it to my own input** while applying it to everyone else's.
- **2 were ANCHOR DRIFT inside iter-84's adjudication** (`alignment_testing.md:463` names the paragraph
  opener; the false claim is at `:470`. `cms.md:67` names the section heading; the claim is at `:71`).
  Both claims **are** repaired. An adjudication that anchors the **block** instead of the **sentence**
  reads as a miss to a line-scoped post-condition, and at tolerance 3 a seven-line offset is a miss.

> **Why iter-85 scored 100 % and this iter scored 87 %, stated plainly:** iter-85 graded against a
> ledger it authored from **one** clean adjudication of 11 anchors. iter-86 assembled **46** anchors
> from **two** sources and inherited both sources' defects. **The reach number is a measurement of the
> ledger as much as of the repair**, and a 100 % on a ledger you wrote yourself the same hour is the
> weaker of the two readings. This is the first time the fence has said something about its own input.

## One more finding, routed rather than tuned

`CHECK-M257x-iter86-value-change-weak-form`. `value_change_guard` built the form
**`('graphql', 'backend', 'stack')`** — one changed token plus the two most ubiquitous nouns in this
corpus — and matched it, in order within a 90-character window, against **two sentences that state the
CORRECTED fact**. **0 of 2 precision on that form.** The floor `MIN_CONTEXT = 2` is derived and
defended in the module, so it is not the defect: **distinctiveness is measured in token COUNT, not in
token RARITY.**

It is **not waived** — and the waiver mechanism is right to refuse it. `is_waived` requires a retraction
marker, so a waiver can only cover a site *quoting a value in order to retract it*. These sites are not
retracting; they are correct. A mechanism that cannot express *"false positive"* is what stops a waiver
file from pinning drift (§8 rule 3), so the answer is a precision fix, not a JSON entry. **Routed.**

What cleared the instance was a **better sentence, not a tuning**: the repaired cell was genuinely
under-specified, and rewriting it into a real rewrite handed the question to `repair_leak_guard`, where
it belongs and is GREEN. That fixes one instance and not the class, and this iter says so.

## The denominator, stated because the source states two

iter-84's verdict table says **40 UPHELD** — an *anchor* count. Its by-predicate tables enumerate **37
rows** — a *predicate-row* count. Separately **Q4's heading reads `(7)` against 8 rows**. iter-85
expanded Q2's 7 rows to 9 anchors. So *"the remaining 29"* is an artifact of the mis-stated Q4 heading;
the enumerated remainder is **30 rows**, and that — not 29, not 33 — is what
[`ledger/r86.md`](ledger/r86.md) declares and what reach is graded against.
