**Type:** tik · **Shape:** tooling (protocol's tooling-iter: ship the instrument, then use it in the
same iter — two planned lines, declared in [`overview.md`](overview.md))

# iter-202 — materialize at a REF, then adjudicate the nineteen

## Line 1 — the instrument iter-198 named and did not build

`claim_census_guard.materialize()` had exactly one substrate: whatever is checked out. Six clones in
this workspace are behind their own already-fetched `origin/main`, so every excerpt it handed an
adjudicator was a stale one — disclosed by `substrate_of`, counted by `substrate_exposure`, and
**unfixable by either**, because neither could offer the alternative. That is why the 19 sat identified
and unresolved from iter-198 to here.

`materialize(ref=…)` reads the cited bytes with `git show <ref>:<rel>`. **No network** — the ref is
already in the clone. `--exposure-adjudicate` renders every exposed pair as the citing corpus unit plus
**both** excerpts, which is the form a judgement can actually be made on.

Three things had to be impossible, and each is fenced:

* a requested ref silently **not applied** (`ref-not-applicable`, empty body, `ref_applied: False` —
  `D-M257x-202-1`);
* file content read through the **stripping** git helper (`D-M257x-202-2`);
* the counter and the renderer reading **different sets**.

The third one fired on its first joint run — **16 against 19** — because a whole-file citation
(`lo is None`) materializes to an *empty* body at both refs, so a renderer deciding `differs` by
comparing bodies read them as identical and dropped three pairs the counter had counted. Two agreeing
reconstructions would have been indistinguishable from a reading; two disagreeing ones announced
themselves. Repaired into **one derivation** (`cited_span_verdict`) that both call, with the CLI
asserting the totals match and **exiting 2** if they ever do not.

## Line 2 — the adjudication

| verdict over the 19 | n |
|---|---|
| corpus correct at `origin/main`; the **stale clone** is what would contradict it (**false-RED**) | **18** |
| corpus wrong at `origin/main`; the stale clone would **confirm** it (**false-GREEN**) | **0** |
| not a staleness defect at all | **1** |

The 18 are one region and verify as a group: at `origin/main`,
`jobsimulation/terraform/main.tf:15-22`/`:15-40` **is** the M810 decommission comment the corpus quotes;
`storage/terraform/main.tf` is **18 lines** exactly as the corpus says (**100** in the stale checkout,
still declaring the ECS module); `:9-11` is *"The ECS service that used to live here is GONE"*;
`messenger/terraform/main.tf:29` is `service_desired_count = 0`.

**iter-198's guess is refuted for this population** (`D-M257x-202-3`). It reasoned false-GREEN was the
*likelier* half *"because the corpus and the clone fell behind together."* They did not — the corpus was
written against `origin/main` in the first place, and several of these blocks disclose the flip in their
own prose. The general *both-directions-reachable* claim is untouched; what changed is that the observed
ratio now has a number.

## What the 19th turned out to be — and what it turned over

`db-backup.md` cited **`storage.tf:24-38`**, meaning db-backup's own terraform. The resolver found no
path match, fell back to the bare basename, found **exactly one** `storage.tf`, and resolved it —
`storage/terraform/storage.tf`, **a different repository**. **No ambiguity guard fired, because the
ambiguity guard is built for ≥2 candidates.** Uniqueness is not correctness: it was uniqueness of the
*filename*, and the citation had already said the filename was not enough (`D-M257x-202-4`).

Repaired on both sides — the citation is now repo-qualified with the collision named in place, and a
repo-qualified target that matches nothing may no longer fall back to its basename. **The guard is as
narrow as the defect and two earlier drafts were not**: guarding all rooted targets took `../x.md`
relative links with it (**434 pairs** unresolved for a class of 1); guarding every repo-qualified miss
re-labelled 21 already-honest `unresolved-ambiguous` pairs. The shipped rule fires only where the old
code would actually have **resolved**. Corpus-wide, that is **1 pair** — the whole class.

## The class it turned over: the parser was truncating four extensions

The wrong-repo guard flagged a **second** site — `ant-academy.md`'s `code/public/catalog.js`. But
`code/public/catalog.json` **exists**. The corpus was right and the **parser** was wrong: `SOURCE_EXTS`
is written in a human order (`js` before `json`, `ts` before `tsx`) and Python's alternation is
**leftmost-first, never longest-match** (`D-M257x-202-5`).

| | before | after |
|---|---|---|
| tier-1 pairs | 3,090 | **3,125** |
| targets read with a **truncated spelling** | **63** | 0 |
| pairs that **materialize** | 582 | **623** |
| pairs materializing from the clone set | 949 | **1,015** |
| `unresolved` | 220 | **167** |
| tier-2 unevidenced (ratchet ceiling 1,164) | 1,138 | **1,130** |

**The +35 is the sharper half.** A bare `AIReadinessClient.tsx` truncated to `.ts` failed the token
boundary and **was not seen as a citation at all** — the census was not mis-reading those, it was not
counting them. Every citation figure this milestone has published was taken through that parser.

Fenced three ways: the truncatable spellings survive; a **mutation control** re-compiles the
declaration-order pattern and requires it to truncate **all** of them; and a **derived** arm that walks
every prefix-pair in `SOURCE_EXTS` itself, so the fence grows with the tuple instead of with this iter's
list.

## Close — 2026-08-09

**Outcome:** the two iter-198 routes are closed together — `materialize()` can now read a citation at a
named ref (`git show`, no network, refusal rather than silent fallback), and with both excerpts in front
of an adjudicator all **19** exposed pairs were decided: **18 false-RED, 0 false-GREEN, 1 not a staleness
defect at all** — which refutes iter-198's own guess that false-GREEN was the likelier half, for this
population. The 19th was a citation resolving into **the wrong repository** because after the resolver
discarded the directory there was exactly one file with that basename; repaired on both sides, and the
guard is as narrow as the defect (two wider drafts cost 434 pairs and 21 re-labels). Chasing its
apparent second member found the class underneath: the citation regex's extension alternation is
**leftmost-first**, so `.json`/`.jsx`/`.tsx`/`.graphqls` were truncated at every occurrence —
**63 targets mis-spelled and 35 citations never seen as citations at all** (tier-1 **3,090 → 3,125**,
materializing **582 → 623**). Censused across the instrument family: **4 alternations, 1 affected**, and
the discriminator is not ordering but whether anything after the alternation can fail.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-fourth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
the two RED arms were a **pre-existing** stale declaration from iter-201, not a regression from this
iter's planned scope, and the repair is mechanical (`D-M257x-202-7`) — (5) cap-reached: n — **counted:**
iter 202 = **one** tik this run — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-202-1` … `D-M257x-202-7` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **280 passed · 1 skipped** across the
**ten** modules that import the changed code or its symbols (`test_claim_census_guard` ·
`test_claim_census_substrate_m257x` · `test_claim_census_skip_registry_m257x` ·
`test_retired_service_endpoints` · `test_test_collection_fence` · `test_frozen_expectation_census_m257x` ·
`test_guard_family_verdict_line_m257x` · `test_suite_census_collection` · `test_suite_census_population` ·
`test_guard_family`, which drives every guard end-to-end). Both changed modules green under **both**
runners (unittest 3.9.6: `Ran 51 … OK`). Guards run directly and green: `claim_census_guard --check`
(ratchet holds, **1,130** unevidenced against a 1,164 baseline), `corpus_citation_guard` (**1,468** path
resolutions, OK), `markdown_structure_guard`, `route_disposition_guard`, `derivation_registry`.
*Scope: `stack-core` only, Python only, importer reach (`§5` r60) — no Go, no TypeScript, and the other
ten rext sections were not run. A whole-section run was started and **STOPPED at ~10 minutes, not
completed**; it is reported here as discarded, never as green. Harden pass 47's **1,699** remains the
last whole-section `stack-core` figure, and it predates both iter-201's runner-gap closure and this
iter's parser fix.*

**Side-deliverables:** `D-M257x-202-7` — iter-201's conversion left `test_claim_census_guard.py` declared
in both `PYTEST_ONLY_MODULES` and `UNITTEST_BLOCKED`; both arms went RED and both were right. Entries
removed, shrink ratchet tightened `≤1` → `== 0`, and a new arm derives the emptiness instead of trusting
it. Separate commit; does not change this iter's close status.

**Routes carried forward:**
- `SURVEY-M257x-iter198-materialization-reads-the-working-tree-by-construction` — **CLOSED.** `--ref` /
  `materialize(ref=…)` ships, proven against a branch, a sha, `HEAD~5`, an unknown ref and a
  not-under-a-clone path, with refusal rather than fallback in the last two.
- `SURVEY-M257x-iter198-the-nineteen-exposed-pairs-are-unadjudicated` — **CLOSED.** All 19 decided;
  18/0/1, the one repaired.
- `SURVEY-M257x-iter202-published-citation-figures-predate-the-truncation-fix` — **NEW.** Every tier-1
  and citation figure this milestone has published was taken through a parser that mis-spelled 63 targets
  and could not see 35 citations. **9 files in this milestone name a tier-1 total.** The figures are not
  wrong for their parser; none of them describes the same population as one taken after this iter.
- `SURVEY-M257x-iter202-anchor-subject-census-extension-vocabulary-is-narrower-than-the-census` — **NEW.**
  `anchor_subject_census.PATHY` omits 13 extensions the claim census knows. Graded: it is used as an
  **exclusion** filter, so a narrow vocabulary fails to exclude — the **conservative** direction. A real
  gap, not a false green; widening it changes what that guard measures and belongs in its own iter.
- `SURVEY-M257x-iter202-the-eighteen-false-RED-pairs-remain-substrate-dependent` — **NEW.** They are
  adjudicated, not *fixed*: `materialize` still defaults to the working tree, so the next reader still
  gets the stale excerpt unless they pass `--ref`. Making `origin/main` the default is a change to what
  every consumer reads and was not attempted here.
- Unchanged and still open: `SURVEY-M257x-iter201-published-suite-totals-predate-the-runner-gap-closing` ·
  `SURVEY-M257x-h45-printed-measurement-literals-uncensused` ·
  `SURVEY-M257x-iter200-battery-stagers-are-safe-by-isolation-not-by-discipline` ·
  `SURVEY-M257x-iter200-only-one-test-module-ever-clears-a-memo` ·
  `SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only` ·
  `SURVEY-M257x-iter199-the-noun-list-is-a-declared-vocabulary` ·
  `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` · and the standing queue.

**Lessons:**
- **Build the other side before you adjudicate.** The 19 sat unresolved for four iters not because
  `F4` forbade the judgement but because nobody could see the second excerpt. The instrument was one
  parameter wide.
- **A leftmost-first alternation is only safe while something downstream can reject the short match.**
  Three sibling patterns in this repo are written in the same unsafe order and are all safe — a `\b`, a
  required `:N`, an anchoring `$`. The one with an *optional* follower was the one that truncated.
- **Repairing a citation is a way of testing the resolver.** Qualifying the path changed nothing, and
  that non-effect — not the original defect — is what exposed the fallback and then the parser.
- **A scoped green is evidence about its scope, and the invoice arrives later.** iter-201's audit was
  correct and its scope was honest; two stale declarations still shipped, and the first run that touched
  the right module found them immediately.
