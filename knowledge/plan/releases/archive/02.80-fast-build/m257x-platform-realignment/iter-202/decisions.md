# iter-202 — decisions

## `D-M257x-202-1` — `materialize(ref=…)`: a requested ref that cannot apply is REFUSED, never silently unapplied

**Decision.** `claim_census_guard.materialize` takes a `ref`. When set, a citation that resolves inside a
git clone of the clone set is read with `git show <ref>:<rel>` instead of the working tree. When the
resolved file is **not** inside such a clone — a `corpus/` file, a path outside the clone set, a clone-set
directory that is not a repo — the call returns `ref-not-applicable` with an **empty body**. Every return
carries `ref` and `ref_applied`.

**Why not fall back.** The fallback is invisible at the call site. A caller that asks for `origin/main`
and receives working-tree bytes cannot tell the difference from a successful read, and an entire
adjudication pass can then be run against the substrate it was explicitly built to avoid — which is the
mistake iter-122 made *by accident* and this would have made *by design*. The refusal is louder and
strictly cheaper: the caller knows immediately.

`clone_of()` is the precondition, factored out so the three `None` cases are one enumerated place rather
than three inline conditions.

## `D-M257x-202-2` — `_git` strips; file CONTENT reads through `_git_text`

**Found while building `--ref`.** `_git()` returned `stdout.strip()`, and `substrate_exposure` used it to
read `git show origin/main:<path>` — file content. `.strip()` removes **leading** blank lines as happily
as trailing ones, so a file whose first line is blank comes back shifted one line up **on the ref side
only**, while the working-tree side is read raw. Every pin into such a file then resolves one line early,
and `substrate_exposure` reports `differs` for a file that is byte-identical.

**Decision.** Split the helper: `_git` (stripped, metadata only — shas, counts, name lists) and
`_git_text` (verbatim, for content). Fenced with the minimal input on which the two must differ, plus an
end-to-end arm that shows the stripping read manufacturing a `differs` out of an identical file.

**Live here?** No — checked, not assumed: none of the nine exposed files opens with whitespace, so the
measured 19 are unaffected. That is exactly why it would have survived until a file that does.

## `D-M257x-202-3` — the 19 exposed pairs are ADJUDICATED: 18 false-RED, 0 false-GREEN, 1 not a staleness defect

With both excerpts side by side (`--exposure-adjudicate`), every one of the 19 was decided.

| | |
|---|---|
| corpus correct at `origin/main`, stale clone would contradict (**false-RED**) | **18** |
| corpus wrong at `origin/main`, stale clone would confirm (**false-GREEN**) | **0** |
| not a staleness defect at all | **1** |

The 18 verify as a group because they cite one region: `jobsimulation/terraform/main.tf:15-22` and
`:15-40` are, at `origin/main`, the M810 decommission comment the corpus quotes; `storage/terraform/main.tf`
is **18 lines** at `origin/main` exactly as the corpus says (100 in the stale checkout, still declaring the
ECS module); `storage/terraform/main.tf:9-11` is the *"The ECS service that used to live here is GONE"*
sentence; `messenger/terraform/main.tf:29` is `service_desired_count = 0`.

**This refutes iter-198's own guess — in the direction that matters.** iter-198 retracted the
single-direction wording and reasoned that false-GREEN was the *likelier* half *"because the corpus and the
clone fell behind together."* They did not: **the corpus was written against `origin/main` in the first
place**, several of these blocks disclosing the flip in their own prose (`storage.md` names the stale ref
and its opposite reading; `messenger.md` says which tree settles which row). Only one of the two substrates
fell behind. **The general claim is unchanged** — both directions remain reachable — but the observed ratio
now has a number instead of an intuition, over a named population of 19 terraform pairs in one migration
region.

## `D-M257x-202-4` — uniqueness is not correctness: a repo-qualified citation must not fall back to its basename

The 19th pair was not staleness. `db-backup.md` cited **`storage.tf:24-38`** — db-backup's own terraform —
and the resolver, finding no candidate, fell back to the bare basename, found **exactly one** `storage.tf`
in the index, and resolved it: `storage/terraform/storage.tf`, **a different repository**. No ambiguity
guard fired, because the ambiguity guard is built for ≥2 candidates.

The repair is two-sided:

* **Corpus** — the citation is now written repo-qualified, with the collision named in place so the next
  editor does not un-qualify it.
* **Instrument** — a **repo-qualified** target (rooted, not `./`-relative) that matches nothing may no
  longer fall back to its basename.

**The guard is as narrow as the defect, and two earlier drafts were not.** Guarding all rooted targets took
`./x.md` and `../services/y.md` with it — **434 pairs** to unresolved for a class of 1. Guarding every
repo-qualified miss re-labelled 21 already-honest `unresolved-ambiguous` pairs. The shipped rule fires only
where the old code would actually have **resolved**: repo-qualified, no path match, **exactly one** basename
hit. Measured corpus-wide, that is **1 pair**, and it is the whole class.

## `D-M257x-202-5` — the citation regex was truncating four extensions, and 35 citations were invisible

Found downstream: the wrong-repo guard flagged a **second** site, `ant-academy.md`'s
`code/public/catalog.js` → `ant-academy/code/ucourses/catalog.js`. But `code/public/catalog.json` **exists**.
The corpus was right and the **parser** was wrong.

`SOURCE_EXTS` is written in a human order — `js` before `json`, `ts` before `tsx`, `graphql` before
`graphqls` — and Python's alternation is **leftmost-first, never longest-match**. Four extensions were
truncated at every occurrence: `.json`→`.js`, `.jsx`→`.js`, `.tsx`→`.ts`, `.graphqls`→`.graphql`.

Measured over `corpus/services/**` + `corpus/architecture/**`:

| | before | after |
|---|---|---|
| tier-1 pairs | 3,090 | **3,125** |
| targets read with a truncated spelling | **63** | 0 |
| pairs that materialize | 582 | **623** |
| `unresolved` | 220 | **167** |
| wrong-repo-guarded (the real class) | 2 | **1** |

**+35 pairs is the sharper half**: a bare `AIReadinessClient.tsx` truncated to `.ts` failed the token
boundary and was **not seen as a citation at all** — the census was not merely mis-reading those, it was
not counting them.

**Decision.** `sorted(SOURCE_EXTS, key=len, reverse=True)` in the alternation, so the tuple's order becomes
cosmetic — which is what it looks like. Fenced three ways: the five truncatable spellings survive; a
**mutation control** re-compiles the declaration-order pattern and requires it to truncate **all** of them
(if it ever stops failing, the property the fix rests on is gone and nothing else would say so); and a
**derived** arm that walks every prefix-pair in `SOURCE_EXTS` itself, so the fence grows with the tuple
rather than with this iter's list.

**Why this landed inside iter-202 rather than routing forward.** It is a third line of investigation and
the tripwire is real — but the iter's *own* published number depended on it. Without the fix this iter
would have reported the wrong-repo class as **2**, one member of which is a parser artifact. A measurement
whose instrument you have just found to be truncating is not publishable as-is.

## `D-M257x-202-6` — the truncation class, censused across the instrument family: 3 alternations, 1 affected, and the discriminator is not ordering

`SOURCE_EXTS` is not the only extension alternation in `stack-core`. Enumerated, there are **four**:

| pattern | what follows the alternation | truncates? |
|---|---|---|
| `claim_census_guard.SOURCE_TOKEN` | `(?::(?P<lo>\d+)…)?` — **OPTIONAL** | **YES** |
| `anchor_construct_guard._QUALIFIED` | `:(\d{1,5})` — required | no |
| `anchor_construct_guard._FILE_MENTION` | `\b` — required | no |
| `anchor_subject_census.PATHY` | `(:[\d-]+)?$` — anchored | no |

**The discriminator is not the order of the alternation — all four are written `js` before `json`, `ts`
before `tsx`.** It is whether anything *after* the alternation can **fail**. A required follower (`\b`,
`:N`, `$`) makes the short alternative fail and the engine backtracks into the long one; leftmost-first
then costs nothing. `SOURCE_TOKEN`'s follower is **optional**, so `catalog.js` + no-line-pin is already a
complete match and there is nothing to backtrack from.

So the rule worth carrying, and it is not *"sort your alternations"*: **a leftmost-first alternation is
only safe while something downstream can reject the short match.** Sorting is the fix that does not depend
on the follower — which is why it is what shipped.

Measured, not assumed: all four patterns were run against `catalog.json`, `x.tsx`, `b.jsx`, `s.graphqls`.

**Routed, not fixed:** `anchor_subject_census.PATHY`'s extension list is a **different** defect — a
vocabulary gap, not a truncation. It omits `graphqls`, `mjs`, `proto`, `graphql`, `env`, `toml`, `lock`,
`conf`, `ini`, `cfg`, `tmpl`, `mod` and `txt`, so `s.graphqls` is not `PATHY` at all. Widening another
guard's vocabulary is a change to what that guard *measures*; it belongs in its own iter with its own
before/after. `SURVEY-M257x-iter202-anchor-subject-census-extension-vocabulary-is-narrower-than-the-census`.

**Direction of the routed `PATHY` gap, so the route carries its own grading:** `PATHY` is used as an
**exclusion** filter (`anchor_subject_census.py:232` — a path-shaped backtick is skipped as a literal
candidate). A vocabulary that is too NARROW therefore fails to *exclude*, which adds candidates — the
**conservative** direction, noisy rather than blind. It is a real gap and it is not a false-green.

## `D-M257x-202-7` — SIDE DISCOVERY: iter-201's conversion left two declarations naming a module it had just converted

Running the **complete** set of test modules that import the changed code — not just the census family —
turned two arms of `test_test_collection_fence.py` **RED**, and they were right:

* `PYTEST_ONLY_MODULES` still named `test_claim_census_guard.py` as *"25 module-level pytest-style test
  functions and NO `unittest.TestCase` subclass"*;
* `UNITTEST_BLOCKED` still carried it as **OPEN**, owned by
  `FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter`.

iter-201 **converted** that module (both runners now collect 29) and **closed** that route. Neither
declaration was removed, so the fences read *"the declaration outlived its subject"* and *"a stale
blocker reads as a live obligation"* — their own words, and both exactly correct. **iter-201's scoped
audit could not have caught it**: it ran the census family and `test_guard_family.py`, and this fence is
in neither.

**This is not a defect in iter-201's work; it is the cost of a scoped audit, priced.** `§5` r60 says a
scoped green is evidence about its scope alone — here is the invoice.

**Landed, not routed** (side-deliverable, separate commit): both entries removed, the shrink ratchet
tightened `≤1` → `== 0`, and — because iter-201's own lesson is *when a repair empties a registry, say so
where the registry is* — a new arm, `test_the_EMPTINESS_is_a_measured_closure_not_a_forgotten_dict`. The
two existing arms cannot tell a closure from an oversight (one iterates the dict, so an empty dict passes
**vacuously**; the other passes when the shape is absent for either reason). The new arm derives the zero:
scan the ≥100-module population, collect every module of the pytest-only *shape*, and require that set to
**equal** the declared set — both empty is a closure, and any other disagreement is rot.
