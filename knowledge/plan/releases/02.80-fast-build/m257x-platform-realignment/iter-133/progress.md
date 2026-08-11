**Type:** tik

# iter-133 — the fence printed the true module set for three iterations while nine sentences contradicted it

Closes `FIX-M257x-iter131-my-three` — and the route's own framing was too small, which is the finding.

---

## 1. Width, measured before repairing (§5 rule 57)

The route names **three anchors**: P5 (`architecture_overview.md:83`), P7
(`platform-migration-status.md` §1) and P19 (`clerk-integration.md:107`). Two independent searches:

| search | hits |
|---|---|
| lines naming `colony` ∧ `proto` ∧ `taxonomy` | **16 lines / 16 files** |
| lines matching `private (go )?modules?` | **22 lines / 18 files** |

**P5 and P6 are one predicate with seven more anchors.** Substituted per Phase 1 Step 0 — the strategy
holds; only the named target was too narrow.

## 2. Ground truth, measured at source (not quoted from `CLAUDE.md`)

| repo | ref | org-private requires |
|---|---|---|
| `app` | `ad9f3c498` | `go.mod:14-18` — **`analytics-go`, `colony`, `proto`, `storage`, `taxonomy`** — five, all direct, no `// indirect`, no org `replace` |
| `sentinel` | `f2c4619` | `colony`, `proto` direct; **`taxonomy` `// indirect`** — adds no sixth |

**The corpus was wrong in two opposite directions at once, and the same paragraph often did both:**

| wrong direction | what it says | why false |
|---|---|---|
| **too many** | still lists `ai` | folded in-tree at `1e457fa70` (2026-08-04); `app/go.mod` requires it no more |
| **too few** | *"three — colony, proto, taxonomy"* | drops **`analytics-go`** and **`storage`**, both **direct** requires |

**The root is a conflation with a cardinality coincidence.** There are two five-member sets:

- the **five historical shared libraries** — `colony`, `authn`, `proto`, `ai`, `taxonomy` — the subject
  set of `shared_libraries.md`;
- the **five private modules a stack imports** — `analytics-go`, `colony`, `proto`, `storage`,
  `taxonomy`.

**They overlap in three and share only a cardinality.** Every repaired site now says which one it means.
`CLAUDE.md` already warns *"do not read this list as `app`'s dependency set; it is not one"* — **and
nine sentences elsewhere did exactly that**, which is the argument for the sweep rather than the warning.

## 3. The repair — 10 sites, 8 files

**The module-set predicate (8):** `corpus/README.md:29` · `corpus/architecture/README.md:22` ·
`architecture_overview.md:83` · `service_taxonomy.md:175` · `service_taxonomy.md:525` ·
`dependency_map.md:42` · `platform_repo.md:122` · `askengine.md:81`.

Each names `app` `ad9f3c498` + `app/go.mod:14-18`, and each states which of the two fives it means.
**`external_services.md:554` was already correct** and is left alone — an upheld claim counted as a
result.

**P7 (1)** — `platform-migration-status.md` §1 gains the missing **`library-unimported`** row. §1 now
defines **9** states; the guard's `ALLOWED_STATES` has **9**; assertion C's description says **nine**.
**All three agreed only after this edit** — before it the checker had nine and the document eight, so
the fence was green over a definition that did not contain the vocabulary it enforced. The row also
states why `library` is not its superset and `decommissioned` not its synonym.

**P19 (1)** — `clerk-integration.md:107`. *"All three sites are the literal `curl -s -X POST
…/sign_in_tokens`"* is false of one: `staging-bringup.md:528` is a **prose bullet** carrying neither
`curl` nor the host. Repaired to **two of three**, with the robust re-derivation restated as the
**shared substring** (`grep -n sign_in_tokens corpus/ops/*.md`, which returns all three).

## 4. Test gates

- **Guard family: 18 GREEN · 0 RED · 4 not-run** (`anchor_offset_guard`, `repair_leak_guard`,
  `repair_reach_guard`, `value_change_guard` — commit-/input-scoped, no `--range`/`--ledger`). **Not a
  whole-family green, and the runner says so.**
- Scoped fence suites — `platform_alignment_guard`, `corpus_citation_guard`, `claim_census_guard`,
  `corpus_index_guard`: **123 passed in 3.14 s**.
- **THE WHOLE SUITE WAS NOT RE-RUN, and §5 rule 60 requires that be said out loud.** iter-132 ran it
  clean **~40 minutes before this iter's edits** (`1 failed · 1208 passed in 2077.16 s`, the 1 being the
  standing RED) on **the same rext tree** (`223e4a6`, unchanged by this iter — iter-133 modified **zero**
  `rosetta-extensions` files). The exposure is therefore bounded to *corpus-reading tests outside the
  four scoped suites*, and the guard family covers the fences among them. **Stated as a gap, not
  characterised as covered** — this is exactly the disclosure rule 60 was written to force.

---

## Close — 2026-08-07

**Outcome:** the route's three named anchors are closed **and P5 turned out to be a predicate, not an
anchor** — the private-module set is misstated at **8 sites in 8 files**, wrong in **two opposite
directions**, because two different five-member sets were being conflated behind a shared cardinality.
Repaired against `app` `ad9f3c498` `go.mod:14-18`, measured at source. P7 closed by giving §1 the
`library-unimported` row it never got — **the guard's vocabulary and the document's definition had
disagreed for three iterations while the guard stayed green**. P19 closed at **two of three**, not
three. **No reading taken; no `N` movement claimed.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged; clause 5 is met only by a reading returning zero and none was
taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no successor strategy is authorable — `TOK-08`'s sealed refutation branch bars one; running under the user's direct brief**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-133-1` … `D-M257x-133-3` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none — this iter modified **zero** `rosetta-extensions` files, deliberately: the
fence was already right, and the prose was the defect.
**Routes carried forward:**
- `FIX-M257x-iter131-adjudication-independence` — **still the first item, still untouched by two
  consecutive iters.** It needs independent agents, not a repair pass, and it is now the oldest
  unactioned route on the milestone.
- `FIX-M257x-iter131-predicate-sets-not-enumerated` · `FIX-M257x-iter131-root-mount-count-underived` ·
  `FIX-M257x-iter132-infrastructure-is-cloneable-so-clone-it` ·
  `FIX-M257x-iter132-marker-fences-cannot-see-retractions` ·
  `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` — all open.
- **NEW — `FIX-M257x-iter133-two-fives-need-a-fence`:** the *historical shared libraries* and the
  *imported private modules* are different sets with the same cardinality, and **nothing fences the
  distinction.** Assertion G fences the map's library ROWS; it does not stop a prose sentence in another
  file from naming three modules. Every site repaired here is repaired by hand and can drift again.
**Lessons:**
1. **A route's anchor list is a sample, not a population.** The route named 3 sites; the predicate had
   10. Measuring width first (rule 57) is what turned a 3-site edit into a corpus-wide sweep.
2. **Two sets with the same cardinality will be conflated, and the coincidence hides it.** Nobody
   noticed *"five libraries"* and *"five imported modules"* were different fives, because both numbers
   were right.
3. **A vocabulary change must land in the definition and the checker, and the checker will not tell
   you.** P7's fence was green throughout, over a document that defined eight states while enforcing
   nine — the mirror image of iter-131's lesson 1 (*a fence that prints the right answer does not
   correct the prose beside it*).
