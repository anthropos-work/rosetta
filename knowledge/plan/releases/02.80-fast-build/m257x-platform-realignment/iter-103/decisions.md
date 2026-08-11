# iter-103 decisions

## D-M257x-103-0 — the coordinator's mid-reading course correction: BOOKED, and two items DEFERRED TO AFTER THE READING because applying them now would contaminate it

A course correction arrived from the coordinator at ~12:33, with the reading in flight. It carries five
items. **The measuring pass does not repair the corpus**, and two of the items would have edited a file the
seats are *instructed by the frozen instrument* to read.

**`briefing-iter76-AS-RUN.md:84` tells every seat: *"Read `corpus/ops/platform-alignment.md` §5 in full
before you start."*** That file is outside clause 5's *scored* scope (`corpus/services/**` +
`corpus/architecture/**`) but it is squarely inside the *instrument*. Editing §5 at 12:35 would mean some
seats graded under the old rule set and some under the new, with nothing in either report saying which —
**the exact defect `§5 rule 41a` exists to forbid, one level up: the instrument is part of the ground truth
too.**

> **Decision.** The two `platform-alignment.md` §5 amendments (items 1 and 4) are **written after the last
> adjudicator returns**, in this same iter, and the iter records that they were deferred and why. Nothing
> else about them changes.

This is not a deferral against the three-fate rule — it is Fate 1, landed in this iter, sequenced correctly.

---

## D-M257x-103-1 — the "three iter-101 tags" ruling is REVERSED. Lane D was right; the record was wrong.

`D-M257x-102-6` recorded Lane D's escalation (*"three `iter-101` tags on origin"*) as a **FALSE POSITIVE**,
attributing it to miscounting `git ls-remote --tags` peeled `^{}` lines. **That ruling is wrong and it is in
the committed record.** Lane B has confirmed it had cut `-101b` and `-101c` and deleted both from origin and
locally. Three tags existed, two of them on one commit — exactly as reported.

**`D-M257x-102-6` is SUPERSEDED, not edited.** Lane D's finding is restored as **UPHELD**.

**Verified independently at this iter, by `ls-remote` (a remote read that moves no local ref, so §5 rule 41a
is not engaged):**

```
0011c10aba0ff0950341cb410265ee59d070afe3  refs/tags/fast-build-m257x-iter-101
09d06070fd99c742d7a671c468abf93074278575  refs/tags/fast-build-m257x-iter-101^{}
```

**Exactly ONE tag** — two `ls-remote` lines because a tag object and its peeled target are two lines. Local:
one tag. Current state matches the coordinator's.

**The error was not arithmetic, and that is the whole lesson.** A tag list was measured at **one instant**
on a surface **three lanes can write**, and the result was treated as a standing fact — then used to
overturn a correct report by another observer who had looked at a different instant. The peeled-`^{}`
miscount is a real mechanism and it is **demoted to a caveat**, not deleted: it is exactly why the
single-instant reading *looked* like a complete explanation. `D-M257x-102-6` even recorded that it was *"not
a complete explanation on its own"* — and the ruling was banked anyway.

**Rule to be written into `corpus/ops/platform-alignment.md` §5 (deferred per `D-M257x-103-0`):**

> **A measurement of a concurrently-mutated surface is timestamped, not standing.** To refute another
> observer's report of such a surface you need **their** timestamp or the surface's **history** — never your
> own later snapshot. A snapshot answers "what is true now"; the report you are refuting answered "what was
> true then". They are different questions, and only one of them was asked.

**This is the same class as `§5 rule 41a`**, one level out: 41a froze the clone refs a *reading* resolves
against; this freezes the evidentiary status of any surface a *concurrent lane* can move.

---

## D-M257x-103-2 — the gate moved to **4 of 5**, and clause 2's disclosure travels with it forever

Proven by the **concurrent lane, not by this reading**, and this iter states it that way everywhere.

- **Clause 1 — MET.** Five consecutive cold cycles, all `green:true` / 0 warnings, no restart
  (621 / 402 / 370 / 370 / 371 s) at platform `0c91421`, refs frozen and proven from the reflog. The
  predicted first-cycle defect never appeared — `7844e97` had already handled the aws-bind migration.
- **Clause 2 — MET WITH DISCLOSURE.** `{passing: 30, failing: 0, unimplementable: 0, unimplemented: 1}`,
  `playwright_exit: 0 / ptreport_exit: 0`, `binding: true, scoped: false`, reproduced on two stacks.
  **BUT on a freshly built stack the first full run failed 29/1 in 2 of 2 attempts.**

> **Binding: any artifact that states clause 2 without the intermittency is wrong.** It is
> **MET-WITH-DISCLOSURE** and never a clean pass. The disclosure is part of the claim, not a footnote to it.

**Clause 5 is untouched by this.** It is still open, still met only by a reading that returns **zero**, and
clause 1/2 movement does not soften it, re-cut it, or license arguing it. Four user rulings.

### `FIX-M257x-iter56-assignment-flake` — the hold is retired, the repair is routed

It was held for lack of a failure rate. There is now a rate **and** a refuted hypothesis:

- `pt-assignment-assign` fails **iff** it is the first load of the assign surface in a full-suite run on a
  fresh stack (**2/2**), and passes once any earlier load has happened (**4/4**).
- **The write always lands** — exactly one row, never two — so the *baseline* is over-read by one.
- The harness's own `baseline-settle-fence` blames a still-filling table; **a 60-sample probe refutes that**
  on both cold and warm stacks.
- Surviving hypothesis: **cross-Playthrough session/org-context bleed** from the ~17 preceding specs.

**Routed to the NEXT repair iter as `FIX-M257x-iter103-assignment-context-bleed`, by CLAIM not by file.**
Not repaired here — this is a measuring pass.

---

## D-M257x-103-3 — three defects booked, and the middle one is booked as OPEN because a correction is a claim too

**DEF-2 — CONFIRMED. `stack-demo/autoverify.json` was a six-day-old ORPHAN.** It read
`green:false, warnings:2`; nothing writes it and nothing removes it, so a grader looking in the obvious
place reads **RED for a green stack**. Quarantined by Lane B.

**Book it for what makes it interesting: this is the milestone's own defect class appearing inside the
gate's own evidence path.** M257x exists because a stale artifact read exactly like a current one
(`|| echo 0` turning a deleted-table error into a plausible `0`). Here the same shape sat in the file the
gate is graded from. A stale green is famous; **a stale RED is worse, because nobody argues with it.**

**DEF-3 — OPEN. Both measurements recorded; NEITHER side asserted.** Lane B says `7844e97` mis-cites
`838d907` for the aws-bind move onto `backend`, and that it was really `d11a403`. Re-measured at this iter,
in `stack-demo/platform`:

| measurement | result |
|---|---|
| `git log -G'\.aws/credentials' -- docker-compose.yml` | `a2a3ee6` · `6daa67e` · `06275db` · `0b7968a` · `467965a` — **names neither commit** |
| bind present at `d11a403`? | **yes** — `docker-compose.yml:78` |
| bind present at `838d907`? | **yes** — `docker-compose.yml:100` |

So `838d907` is **wrong** *and* `d11a403` is **unverified**; the move predates both, possibly across the
compose split at `06275db`. **A correction is a claim too**, and this one is booked OPEN as
`DEF-M257x-iter103-aws-bind-provenance` rather than swapped in. Note for whoever settles it: **`-S` is the
wrong pickaxe for a MOVE** — it counts occurrences and a move keeps the count. Use `-G`.

**DEF-4 — CONFIRMED, and it is a correction to the MEASURING APPARATUS, which is why it was checked first.**
iter-102's brief and the milestone `progress.md` state that across the `app` clone advance
`terraform/main.tf` was **"byte-identical"** and that *"the entire residual is a LABEL."* Re-measured
`2035f9a4..ad9f3c49` in `stack-demo/app`:

| file | `git diff --numstat` | lines at `2035f9a4` → `ad9f3c49` |
|---|---|---|
| `terraform/main.tf` | **1 insertion / 1 deletion** | 786 → 786 |
| `terraform/variables.tf` | **37 insertions / 12 deletions** | 738 → **763** |

The changed `main.tf` line is an `error_message` **prose string** (the `atlas_sentinel_dev_url` validation
message, rewritten and lengthened). **The CONCLUSION survives and was re-verified**: `main.tf:181` is
`service_desired_count = 1` at **both** refs, so no *cited* construct moved. **The WORDS do not survive.**
"Byte-identical" is false, and "the entire residual is a label" is false as stated — 49 lines of
`variables.tf` moved and were never in the residual accounting at all.

> **This is the milestone's own class, in the milestone's own records, for the second time in two iters** —
> iter-102 found the corpus inheriting a false claim from a platform commit message, and here iter-102's own
> record over-states a measurement in the direction that made its conclusion cleaner. **The conclusion was
> right; the evidence sentence was stronger than the evidence.** That is the same shape as iter-102's CANON-2
> defect (*a verdict weaker than its evidence*) with the sign flipped.

**Checked FIRST, and the answer is that this reading is clean.** The only occurrence of "byte-identical" in
any `iter-103/` artifact is `pre-registration.md:124`, describing the **briefing instrument**, whose sha
`3858ec53…` was verified after copying. **No iter-103 ground-truth, briefing, or band statement carries the
false claim, so the measurement is uncontaminated.** The false sentences live in `iter-102/` records and the
milestone `progress.md`; those are milestone **records**, not the corpus under audit, and correcting a
record is bookkeeping rather than corpus repair — so it is in scope for this pass and is done at the close.

---

## D-M257x-103-4 — `§5 rule 41a` is PARTIALLY UNENFORCEABLE, and a rule believed enforced is worse than a disclosed gap

`ensure-clones.sh` runs `git fetch` on **every** bring-up as a freshness assertion and **cannot be
suppressed**. It moves only `refs/remotes/*`, never a working tree (`DEMO_ADVANCE_CLONES` defaults to `0`).
**But `refs/remotes/*` is precisely the surface a citation guard resolves against** — `CITE_REF=auto`'s
ladder is `origin/main` first. One such fetch caught `next-web-app`'s `origin/main` advancing **4 commits**
past the frozen ref mid-run.

**This reading is unaffected: Lane B is finished, no bring-up ran during it, and the fetch times recorded in
`ground-truth.md` are re-read at the close and published.** But 41a as written forbids something the tooling
does unconditionally.

**Amendment to be written into `platform-alignment.md` §5 (deferred per `D-M257x-103-0`):** state what 41a
**can** and **cannot** enforce — it binds *lanes*, and it cannot bind `ensure-clones.sh`, so a reading that
overlaps a bring-up must **record the fetch and treat the affected refs as moved**, never assume the rule
held.

---

## D-M257x-103-5 — no tag is cut in this run; both rext commits fold into the NEXT cut

rext `main` is 2 commits past `fast-build-m257x-iter-101`: `4cb920a` (advances
`demo-stack/clones.pin.json` to the proven topology) and `944fc4a` (the 8 acknowledged-site waivers). Lane B
verified `4cb920a` is **not** needed for clauses 1–2, **but a fresh box will not reproduce this clone set
without it.**

**Both fold into the next tag cut. No tag is cut now, or at any point in this run** — the standing
constraint is unchanged, and there remains exactly one `iter-101` tag on origin.

---

## D-M257x-103-6 — the reading itself: what was changed before it opened, and what deliberately was not

**Changed** (instrument-adjacent, in the ADDENDUM only, never above the frozen line): the addendum now names
**which** `rosetta-extensions` tree settles a claim, per §5 rule 45. Band #6 measures whether an addendum can
repair a defect in a frozen instrument without editing it.

**Not changed, on purpose:** `briefing-iter76-AS-RUN.md:37` still names the authoring copy as "the tooling",
for the third reading running. Editing it would break the comparability the series exists to establish.

**Disclosed, not smoothed:** this reading is **not a replicate**. iter-102 grew the in-scope corpus
10,278 → 10,646 lines (+3.6 %), so the greedy LPT partition was recomputed and differs from iter-101's. The
partitioning **algorithm** was proven unchanged — the same script reproduces iter-101's published partition
**exactly** when run over the file sizes at `8f04d3a`.

---

## D-M257x-103-7 — a guard verdict depends on WHICH rext tree ran it, the two trees disagree by 8 sites today, and I nearly published the false half of that

**Found at the close, by running the guard family and getting a verdict that contradicted this iter's own
ground-truth sheet.**

`ground-truth.md` records the open as **`14 GREEN · 0 RED · 0 could-not-check · 3 not-run`**. Re-running the
family after the §5 amendments, I got **2 RED** — `claim_twin_guard` (10 published sites restating an
already-refuted claim) and `repair_postcondition` (the same 10). Reproduced at the reading's own subject
`e6aed2e`, materialised read-only with `git archive`, so it was not my edit.

**The obvious conclusion was available, quotable, and wrong.** It was: *"the ground-truth sheet asserted a
guard verdict it did not have — the milestone's own class, in the reading's own apparatus"*, plus *"a machine
fence names 8 in-scope sites the 14-seat double reading missed, so `N ≥ 41`."* Both sentences were drafted.
**Neither is true.**

**What was actually happening — I ran the fence from the wrong tree.**

| tree | ref | `claim_twin_guard` | `repair_postcondition` |
|---|---|---|---|
| `stack-demo/rosetta-extensions` (**pinned per-stack**) | `09d06070` | **rc=1 RED**, 8 sites | **rc=1 RED**, 8 sites |
| `.agentspace/rosetta-extensions` (**authoring**) | `944fc4a2` | **rc=0 GREEN** | **rc=0 GREEN** |

Identical at both subjects (`e6aed2e` and the working tree). The whole difference is one file:
`git diff 09d06070..944fc4a2 -- stack-core/` is **`claim_twin_waivers.json`, +40 lines, 1 file changed** —
rext `944fc4a`, *"the 8 acknowledged-site waivers"*. **The 8 RED sites ARE the 8 waived sites.**
`ground-truth.md` was right; the family is **14 GREEN · 0 RED**, re-confirmed after the §5 edits.

### The finding that survives, and it is not small

**§5 rule 45 says *name the tree that settles a claim*, and the settling tree follows the SUBJECT.** For a
corpus claim about **what the tooling does on a stack**, that is the **pinned per-stack clone** — the code a
stack executes. **A guard VERDICT is not stack behaviour.** It is a measurement of the corpus taken with a
fence's *configuration*, and it is settled by the tree that configuration lives in — for an authoring-time
fence, the **authoring copy**. Running a fence from the pinned clone measures **last release's fence**, and
any waiver or rule added since reads as a **fresh RED at sites nobody touched**.

> **Rule.** A guard verdict must name the rext tree that produced it, exactly as a seat must name the tree
> that settles a corpus claim. `guard_family.py` prints the corpus sha and the platform sha and **not its own
> sha** — so the one input that decides the verdict is the one input the output does not state.

**This is `DEF-M257x-iter101-briefing-rext-tree` inverted.** That defect is a briefing naming the *authoring*
copy where the *pinned* clone settles it. This is an operator reaching for the *pinned* clone where the
*authoring* copy settles it. **The defect is not "the wrong tree" — it is that neither the instrument nor the
tooling makes you SAY which tree, so the choice is made silently and the answer changes.** Band #6 measured
the seat-facing half at 4 → 1 → 1 across three readings. This is the coordinator-facing half, measured once,
at 8 sites.

### Why this is booked rather than quietly fixed

**The false version was one commit away from the record.** What stopped it was checking a surprising
measurement against a second tree instead of banking it — the discipline `D-M257x-103-1` was written about
this same iter, where a single-instant snapshot was used to overturn a correct report, and `DEF-4`, where an
evidence sentence ran ahead of its evidence in the direction that made the conclusion cleaner.

**Three for three, in one iteration: the milestone's own class keeps landing on the milestone's own
apparatus, and each time the thing that caught it was re-measuring rather than reasoning.** A guard verdict
that contradicts a sheet is a *disagreement between two observers*, and §5 rule 49 — written earlier in this
same iter — says the first hypothesis is that the surfaces differ, not that one observer erred.

**Routed as `FIX-M257x-iter103-guard-tree-provenance`** — `guard_family.py` and each member print the rext
tree path **and sha** they ran from, and the milestone's ground-truth sheets record it beside the corpus and
platform shas. Not repaired here: this pass repairs nothing, and rext stays on `main` with no tag cut.

**No corpus defect was found or fixed by this.** The 8 sites are acknowledged waivers, `N` is unchanged at
**33**, and no band moves.
