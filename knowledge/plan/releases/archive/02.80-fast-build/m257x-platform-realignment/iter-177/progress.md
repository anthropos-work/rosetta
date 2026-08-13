**Type:** tik (standard shape — no protocol-codified iter shape selected; see
[`platform-alignment.md` §9](../../../../../../corpus/ops/platform-alignment.md) for the iter-type
refinements consulted).

# iter-177 — the retraction was wrong, and the correction is a LABEL

**Controlling strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.* The class here
is decidable by arithmetic over sets: **how many distinct derivations of "the fence family" are live, and
do the numbers published about it name which one they came from.**

## Phase A — the measurement (Step 0 re-survey, before any edit)

Pre-flight (`§0d`): `tests/test_fence_registry_population_m257x.py` → **8 passed in 6.07 s** under
`/usr/bin/python3` **3.9.6**, the only interpreter on this host with pytest (r75/76 — *name the runner*).
Green before the first edit.

Then the cheapest possible check — three lines of Python against the three owners:

| derivation | owner | rule | at rext `c7f4c3d` | at `5b108d0` | named in `README.md` |
|---|---|---|---|---|---|
| `union` | `guard_family.union` | spelled ∪ declared ∪ `EXTRA_CENSUS_MEMBERS` | 27 | 27 | **16** |
| `census` | `guard_family.census` | `union` − `CENSUS_EXCLUSIONS` | 26 | 26 | **16** |
| `declaring` | `repair_postcondition.declared_kind` | declares `FENCE_KIND` | 26 | 26 | **15** |

`5b108d0` is iter-175's own commit, reconstructed with `git archive` — deliberately the same method
harden pass 39 used, so the two readings are comparable rather than merely different.

**The retraction is falsified.** Pass 39 recorded *"the disclosed limit said `16 of 27`; it was `15 of
26` the day it was written… both operands were wrong at publication."* **`16 of 27` is an exact reading
of `union`** at both refs — the function iter-175 authored *in that very commit*. And which set iter-175
enumerated is not an inference: its routed text lists the **11** missing members **by name**, and that
list contains `guard_family`, which `census` excludes and `declaring` includes. Only `union` fits.

**What iter-175 actually got wrong was the population LABEL** — it called a `union` reading *"the
census"*. Pass 39 corrected the *number* onto a **third** population without noticing there were three,
and its replacement is equally correct and equally unreadable on its own. → `D-M257x-177-1`,
`D-M257x-177-2`.

## Phase B — the repair, at the site that carries the live claim

`stack-core/tests/test_fence_registry_population_m257x.py`:

* the module docstring and the disclosure test's own docstring now publish the **triple**, each figure
  labelled with the function that produced it, and both carry the retraction-of-the-retraction;
* `_derivations()` reads all three **through their owners** — a private copy of any rule here would be a
  **fourth** derivation, authored by the test that exists to count them (iter-175's rule);
* `test_the_disclosed_limit_is_STATED_not_assumed` asserts the whole claim rather than a fragment of it:
  every `**N of M**` must carry a parenthesised derivation name, each is checked against **that** derivation's live
  value, an unknown label is RED, and **all three** must appear — so a fourth cannot enter the tree the
  way the third one did, by being nobody's declared subject. → `D-M257x-177-3`.

The assert's history is the argument for it:

| version | assert | could not detect |
|---|---|---|
| iter-176 | `len(named) >= 2` | an error of one in either operand |
| harden 39 | one unlabelled `(N, M)` equality | that **three** populations answer to one name |
| **iter-177** | one labelled figure per derivation, all three required | — |

## Phase C — fencing the coincidence

`census` and `declaring` are **both 26** and differ by **one member in each direction** —
`guard_family` (in `declaring`, out of `census`) and `repair_postcondition` (in `census`, out of
`declaring`). Both dispositions are correct and both are now written down. **The defect is that nothing
said they differ, and every comparison by COUNT reads green** — which is how one population came to
publish `15 of 26`, `16 of 26` and `16 of 27` at the same time with nothing going RED.

Three new tests:

1. `test_the_two_26_member_derivations_differ_by_MEMBERSHIP` — the symmetric difference must equal the
   disposition table; a third difference is RED in either direction. → `D-M257x-177-4`.
2. `test_the_three_derivations_form_ONE_algebra` — `union == census | declaring`, so a fourth cannot hide
   between them. → `D-M257x-177-5`.
3. `test_mutation_control_a_COUNT_comparison_reads_GREEN_on_this_disagreement` — the control that makes
   (1) non-obvious: it asserts the counts **do** agree while the sets do not. If the coincidence ever
   ends it goes RED **on purpose**, so the hazard is re-characterised rather than quietly deleted.

## Phase D — RED-proof, then run

Every new assert was mutation-proved in a **scratch copy** of `stack-core` (`/private/tmp/**`), never in
the tree — the tree was not edited while anything ran:

| # | mutation | result |
|---|---|---|
| M1 | strip one derivation label from the docstring | RED — *"publishes 1 `**N of M**` pair with NO derivation label"* |
| M2 | **restore harden pass 39's `15 of 26` onto the `union` label** | RED — publishes 15 of 26 for `union`; measured 16 of 27 |
| M3 | delete the `declaring` row | RED — *"does not publish one figure per derivation"* |
| M4 | add a third, undocumented divergence (`EXTRA_CENSUS_MEMBERS` grows) | RED ×3, incl. the membership assert |
| M5 | break the algebra (an exclusion outside `declaring`) | RED ×4, incl. `…form_ONE_algebra` |
| M6 | converge the two derivations to one set | RED — the count-coincidence control, with its re-derive-don't-delete message |

Baseline restored between mutations; the pristine copy is green.

**Suite (counts, never wall-time — `§5` rule 51's timing leg fails on this host):** see the close below.

## Phase E — routes

The hardening ledger's own Finding 2 still carries the retracted retraction. `hardening-ledger.md` is
owned exclusively by `/developer-kit:harden-mstone-iters` and this skill is **forbidden** from writing
it, so the correction is **routed to its owner** and the live instrument carries the current reading
instead. → `D-M257x-177-6`.

## Close — 2026-08-09

**Outcome:** harden pass 39's retraction of iter-175's `16 of 27` is **itself retracted, measured**:
`16 of 27` is an exact reading of `guard_family.union()` at HEAD *and* at `5b108d0`, and iter-175's own
missing-member list (11 names, including `guard_family`) proves that is the set it enumerated. **One
population, three live derivations — `union` 27, `census` 26, `declaring` 26 — and two of them return
the same count over different members**, so every count-based comparison of them reads green. The
defect was never arithmetic; it was the missing population LABEL, in the claim *and in its retraction*.
The disclosure now publishes one labelled figure per derivation and asserts each against its own owner;
the two 26-member sets are compared by **membership** against a table of written dispositions; and
`union == census | declaring` pins all three into one algebra. **6 mutations RED-proof the new asserts**
(including *restore pass 39's figure* → RED). `FIX-M257x-h39-survey-id-embeds-retracted-figure` is
**closed as refuted** — the id embeds a correct figure.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (ninth consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9` iter-type refinement, and
`TOK-08` declares the class-by-class sweep order in advance) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-177-1` … `D-M257x-177-6` (see [`decisions.md`](decisions.md))

**Suite (counts, never wall-time as a verdict — `§5` rule 51's timing leg fails on this host):**

| suite | runner | result |
|---|---|---|
| `stack-core/tests` (whole section, 1,550 collected) | `/usr/bin/python3` 3.9.6 | **1,548 passed · 2 skipped · 0 failed** (1314.44 s) |
| `tests/test_fence_registry_population_m257x.py` (pre-flight, before any edit) | same | 8 passed (6.07 s) |

**The delta reconciles exactly.** The standing baseline is **1,545 passed / 2 skipped / 0 failed**; this
run is **1,548 / 2 / 0** — `+3`, which is this iter's three new tests and nothing else. **Scope stated
(r60/66):** this is `stack-core` only. The other four `rosetta-extensions` sections were **not** run and
nothing is claimed about them; the edit's blast radius was measured first — `test_fence_registry_
population_m257x` is referenced by **no** other file in the monorepo, and no fence module was added, so
no ratchet baseline or battery seed list changes.

**Side-deliverables:** none.

**Routes carried forward:**
- `FIX-M257x-iter177-ledger-carries-a-retracted-retraction` — **NEW.** `hardening-ledger.md` pass 39
  Finding 2 and its routed-forward list still publish *"both operands were wrong at publication"* and
  the rename route built on it. This skill may not write that file (it is owned by
  `/developer-kit:harden-mstone-iters`); **handler = the next harden pass.** The measurement and both
  refs are in `D-M257x-177-1`.
- `FIX-M257x-iter177-derivation-registry-decline-rationale-is-false` — **NEW, measured not asserted.**
  `derivation_registry.py:224-227` justifies declining `guard_family::declaring_modules` and `::union`
  on the ground that *"each returns the same population `census` does (modulo `CENSUS_EXCLUSIONS`)."*
  **True for `union`** (`census ∪ CENSUS_EXCLUSIONS` = 27 = `union`); **false for `declaring_modules`**
  — it differs from `census` by two members, and adding the exclusions back yields `union`, still one
  member away. Not landed here: it is a second line of investigation and, more concretely, an edit to
  that file after the suite ran would have invalidated the run this close reports.
- `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` — **narrowed, and the id is now CORRECT.** What
  remains open is the part neither iter stated: **which derivation the README is meant to be complete
  against.** The gap is 11 (`union`) / 10 (`census`) / 11 (`declaring`) unnamed members.
- `SURVEY-M257x-iter175-census-vs-discover_fences-classified-differently` — unchanged as a
  classification question, but its factual premise is now measured (see the `derivation_registry` route).
- `FIX-M257x-iter174-accept-registers-one-registry-of-two` — unchanged (member open inside the fence).
- `FIX-M257x-iter173-ledger-denominator` / the observed half of
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` — unchanged; open.
- The standing queue, unchanged.

**Next-iter lead (read-only, taken while the suite ran and the tree was frozen):** the `N of M` prose
class this iter's rule is about is **censusable and small** — **18** bolded `**N of M**` constructs in
`corpus/**` + `CLAUDE.md`, **61** including unbolded, and **8 inside clause 5's scope**
(`corpus/architecture/**` + `corpus/services/**`). That is the unmeasured half of
`SURVEY-M257x-iter173-derived-count-guard-reach`, and unlike this iter it is scoped where `P` lives.

**Lessons:** **a count about a population is unreadable until it names the derivation that produced it —
and a retraction is a count too.** A retraction inherits every weakness of the claim it retracts: this
one omitted exactly what the original omitted, so it did not correct the error, it doubled it and added
authority. Two mechanisms follow, and the second generalises further than the first: publish **one
labelled figure per derivation and require all of them**, and when two derivations return the **same
count**, compare them by **membership** — with a control asserting the counts *do* agree, so the day the
coincidence ends the instrument says so instead of quietly becoming vacuous. Written into
`platform-alignment.md` §8 in this iter's commit.
