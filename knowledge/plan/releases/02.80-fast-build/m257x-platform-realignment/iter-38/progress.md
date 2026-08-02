**Type:** tik

# iter-38 — the fourth clause-5 pass: 17 blockers, and the routing that would have found 11

## The measurement

Six auditors, **40 files / 8 624 lines**, every in-scope file read **top-to-bottom**, re-partitioned so no
auditor inherited iters 33/34's boundaries (§5 rule 18(b)). Combined verification volume, self-reported per
auditor: **~510 exact citations** checked against the platform clones at origin `2adcf71`, the live
`demo-1` Postgres, or `docker-compose.yml`/`repos.yml`.

**Result: 17 blockers.** Enumerated with anchors in `evidence/blocker-ledger.md`. All 17 fixed.

## The pre-registered prediction was REFUTED on both count and location — and that is the finding

Written into this iter's `overview.md` before any auditor reported: *"2–5 blockers total; at least 3 inside
the 8 repaired files; **0–1 across the 32 untouched**."*

|  | files | blockers | per file |
|---|---|---|---|
| repaired by iter-34 (in clause-5 scope) | 8 | **11** | 1.375 |
| never opened by iter-34 | 32 | **6** | 0.19 |

The **density ratio reproduces** §5 rule 18's ~9× (here ~7.3×). The **absolute counts do not**: 17, not
2–5, and 6 in the untouched set, not 0–1.

**So the routed instruction — *"scope the next pass to the 9 changed files"* — would have found 11 of 17
and reported everything else clean.** Rule 18 licenses **weighting**, not **narrowing**, and this iter
declined to narrow *before* seeing the numbers, on the ground that clause 5 asks for a verdict over
`corpus/services/**` + `corpus/architecture/**` and a pass reading 8 of 40 files cannot return one.
iter-21 is this milestone's own precedent: a scoped audit converged on a number a full read then multiplied
by five.

**And two of the six untouched-file blockers were in a file two prior full-read passes had already
cleared.** `ai_architecture.md` had been read cover-to-cover twice and passed twice; this pass found that
its EU-AI-Act premise was false and that a competency-threshold ladder it has published for releases
**does not exist in any repo**. What changed was not diligence but *partition* — which is exactly what rule
18(b) predicts: correlated blind spots are a property of how the corpus was divided, not only of who read
it.

## The one that matters most, and it was found twice independently

**`security_compliance.md:7,183` and `ai_architecture.md:7,151` both stated that simulation scoring is not
done by AI — and both offered that as the reason the platform classifies as Limited Risk under the EU AI
Act.** Two different auditors, reading two different files, refuted it from the same source within minutes
of each other.

It is a conjunction and **both conjuncts fail**. The rubric *arithmetic* is deterministic —
`calculateSkillScore` counts booleans and divides. **The booleans are LLM output**: the LLM checker is
constructed at `basevalidator/criterion.go:428` and fed
`basevalidator/templates/checkValidationBulk.tmpl`, a prompt asking a model to *"assess whether the
`<asset>` … meets or does not meet"* each check and return `{"check_id", "feedback", "success"}`. So *"AI
is used for conversation/generation only"* is false too.

Both files now carry the retraction, and both say the same thing about what to do with it: **this is a
question for counsel, not for this corpus.** A system judging workers and candidates sits near Annex III.
The corpus's job was to stop asserting a legal conclusion from a false technical premise, and it has.

The neighbouring blocker is of the same family and worse in one way: the published ladder *"Level 1 ≥ 60,
Level 2 ≥ 65, Level 3 ≥ 75, Level 4 ≥ 85, Level 5 ≥ 95"* **does not exist anywhere** — not in `app`, `cms`,
`jobsimulation` or `next-web-app`. The real conversion is `max(0, score*2-100)` with a `// TODO fix this
formula` comment beside it. A fabricated-looking precision had been standing in for a formula the platform
itself is unsure of.

## The adversarial pass over this sweep — 6 self-inflicted, third consecutive time

Mandatory under §5 rule 18(a), and it earned its keep again. The full table is in the ledger; the two worth
naming here:

- **A half-applied edit left a broken sentence.** `hiring.md` gained a doubled *"The The"* and an orphaned
  predicate — *"…What actually gates that scoreboard / won't treat the cohort as hiring"* — which both
  re-asserted the claim just withdrawn and promised an answer it never gave. Mechanical damage from an
  anchored replacement whose anchor cut a sentence in half.
- **A load-bearing citation pointed at a dead field, on the compliance page.** The retraction cited the
  `checkerEngines` map as the mechanism. That map is **stored and never read** — dispatch is a hardcoded
  switch. And `EngineTextDiff` checks **do** run deterministically (`criterion.go:168`, `:450-475`), so
  *"the verdicts come from an LLM"* was false as a universal: the honest claim is **most**, not all. Both
  files now say so.

Also over-corrected: `hiring.md` swung from *"the insights path requires `is_hiring`"* to *"`is_hiring`
drives the re-skin, not the read path"* — **both conjuncts of the replacement were false**, since the
re-skin is Clerk-derived and the *content-library* read path does branch on the column. Corrected to name
which read path does and which does not.

**Every `file:line` anchor the sweep introduced resolved correctly** (~50 verified). All six defects were
in surrounding prose, over-correction, or mechanical damage — the distribution rule 18 predicts, observed
now for the third time.

## Also verified, and worth recording as a clean result

The **multi-tenancy fence** — wrong four times, in both directions — is **CORRECT in its fifth
generation**, and this is the first pass to establish that by testing the **predicate** rather than the
denominator (§5 rule 17). Re-derived independently twice (by me and by the auditor who owns the file):
139 `.go` / 135 schemas / 30 / 7 / 18 all reproduce, **and** a per-file test of both conjuncts confirms
the two exclusions. The auditor went one better than my check: only **five** `Policy()` declarations exist
in the whole schema directory, so "no policy of any kind" is exhaustively true of the named 16. Full
working in `evidence/tenancy-fence-rederivation.md`.

Also clean: **13 of the 29 service docs returned ZERO blockers across 85 verified citations**, and the
8-file remainder of `corpus/architecture/**` returned 3 across ~95.

## Clause 5 — NOT MET, deliberately

**A clause is met by a READING, not by a repair.** This pass returned 17, then 6 more from its own
adversarial half — 23 found and fixed. That is a productive pass and it is not a zero-blocker reading.
iters 33 and 34 both refused to claim the clause on exactly this ground; so does this one.

## Close — 2026-08-02

**Outcome:** clause-5 fourth pass measured at **17 blockers** (+6 self-inflicted, all fixed; 23 closed
across 26+15 = 41 exactly-once anchored edits in 11 files). Pre-registered prediction refuted on count and
location. Clause 5 stays open pending a zero-blocker reading.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (4 of 5 — clauses 1, 2, 3, 4 hold; clause 5 outstanding)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-38-1 (run it wide anyway — rule 18 weights, it does not narrow) · D-M257x-38-2 (retract the compliance premise, do not re-classify) — see `decisions.md`
**Side-deliverables:** none.
**Routes carried forward:**
- `MEASURE-M257x-iter39-clause5-fifth-pass` — the confirming reading. **Re-partition again** (rule 18(b)
  has now paid twice). Weight toward the 11 files this pass edited; do NOT narrow to them.
- `DOC-M257x-iter38-ops-collateral` — **four `corpus/ops/**` docs assert what this sweep just retracted**,
  out of clause-5 scope but wrong now: `playthroughs.md:521-523` builds a **test-design rule** (principle
  P2's exemption for simulation scores) on *"scoring is deterministic … NOT AI-scored"* and links to the
  section that now retracts it; `platform_repo.md:111` still lists `SkillPathSessionService` in the RPC
  mux; `content-stories-routes.md` + `content-stories-spec.md` use the singular `academy_*` table names
  that now error. The first is the one that changes behaviour.
- `CHECK-M257x-iter38-ai-act-classification` — the corpus no longer asserts the false premise, but nobody
  has re-derived the classification. **This is a legal question, not a documentation one**; it needs an
  owner outside this milestone.
- `DOC-M257x-iter38-minors` — ~45 minors with exact anchors across the six auditor reports (line drift,
  undercounts, omitted list members). None block clause 5 (*"YELLOW with 0 blockers"* admits them).
**Lessons:**
- **A routing instruction is a hypothesis, not an instruction.** *"Scope to the 9 changed files"* was
  well-founded, derived from a real 9× measurement — and following it would have missed 6 of 17 blockers
  including the two most consequential in the corpus. Weighting and narrowing are different operations and
  a density measurement licenses only the first.
- **Two auditors finding the same false claim in different files is the strongest signal a partition can
  produce.** It cannot happen when one auditor owns both files, which is how it survived two passes.
- **A repair pass damages text mechanically, not only semantically.** An anchored replacement that cuts a
  sentence in half leaves a fragment that reads as prose. Diff-read the sweep, not just its anchors.
