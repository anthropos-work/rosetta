---
iter: 137
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-137 — roadrunner: two false predicates, pulling in opposite directions

**Type:** tik
**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* This tik
works the **adjudicated work list**, which is the census output `TOK-08` asks for: a list of live blockers
with anchors, not a sample. Selected **by consequence** per `TOK-08`'s carried findings and iter-136's
lesson 3.

## Step 0 — Re-survey before targeting

`TOK-08`'s next-target is carried by iter-136's close: the remainder of
`FIX-M257x-iter135-adjudicated-live-defects`. Re-surveyed at HEAD (`6a872c0`):

- `roadrunner.md:12-14`, `:30-32`, `:53`, `:74` — **all four live, unrepaired.**
- `architecture_overview.md:228` — **live**, and still the half-applied repair `adj-F` described (the CMS
  row one line above at `:227` carries the full iter-123/127 correction; the Roadrunner row does not).
- `dependency_map.md:9` — **live** (verified `awk 'NR==9'`).

Target still current. No substitution.

## Cluster / target identified

Two adjudicators, blind to each other, hit **roadrunner** from opposite directions in iter-135:

| | adjudicator | predicate | direction |
|---|---|---|---|
| **Q1** | `adj-F` P2 | *roadrunner's production retirement is unresolved/pending — prod terraform still reads `service_desired_count = 1`* | says it is **still live in production** |
| **Q2** | `adj-B` P-2 | *`roadrunner` is one of EIGHT domains folded INTO `app`* | says it was **absorbed into the monolith** |

**Both are false, both are live in the corpus at the same time, and they contradict each other.** That is
the shape worth taking: a reader can find, in two files of the same corpus, that roadrunner is running in
production and that it is a package inside `app`. Neither is true.

Per iter-136's `D-M257x-136-1` the conjuncts are **stated separately** so each is independently falsifiable
— the brief-teaches-the-error defect iter-136 just fixed.

## Hypothesis

The settled verdict already exists — `org-repos.md:140-146` and `platform-migration-status.md:90` carry the
iter-123 measurement — and **never propagated**. This is the half-repair class (iter-102 corrected
`services/README.md` and the fenced map; the twins no partition owned were left standing). Repairing at
every site the width search returns closes both conjuncts.

## Expected lift

No `N` movement claimed — **no reading is taken this iter** (`§9`'s iter-type refinement: an iter that took
no reading has an UNMEASURED metric, not an unmoved one). The deliverable is the predicate closed at every
site, with ground truth **re-derived at source** rather than taken from the adjudicators.

## Phase plan

1. **Width first (`§5` rule 57)** — ≥2 independent searches per conjunct before any repair. *(done at open)*
2. **Ground truth re-derived at source** (`D-M257x-136-1`: re-enumerate; never accept the reporter's
   candidate). Four checks with positive controls: `app/internal/roadrunner/` existence at every ref; the
   Judge0 wiring's actual home; `roadrunner/terraform/main.tf`'s module-vs-root shape; `infrastructure`
   @ `13c248e6` org-wide.
3. **Repair** every site both conjuncts reach.
4. **Guard family + scoped fence suites**; whole-suite decision stated out loud (`§5` rule 60).

## Escalation conditions

- If a third distinct line of investigation opens → tripwire; land the predicate, route the rest.
- If the ground-truth checks **contradict** the adjudicators → report the refutation, do not repair to the
  adjudicator's story (iter-136's precedent, where the seat's number was right and its candidate wrong).

## Acceptable close-no-lift outcomes

If measurement shows the corpus's current wording is defensible at most sites, the iter closes on the
falsification with the width recorded — an upheld claim counted as a result, this milestone's practice.
