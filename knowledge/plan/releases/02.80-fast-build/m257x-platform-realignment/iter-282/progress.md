# iter-282 — census the corpus/tooling prose copies, and grade the copies that disagree

**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Step 0 — the re-survey held, for once

Three consecutive iters have opened on a route that was already closed, so the route was re-verified
before any work: `guard_family.py`'s member list, the `stack-core/tests/` module list and a grep across
the tooling for a prose/copy/verbatim fence all return **nothing that grades this class**.
`ROUTE-M257x-h70-corpus-and-code-prose-are-copies-with-no-fence` is genuinely open, and it is the largest
remaining structural item on limb 3.

## Phase A — size the population BEFORE designing the fence, and let the triage pick the predicate

The route's inherited figure is **172 (module, doc) pairs sharing a verbatim 11-word run**. That number
is a measurement of **copying**, and copying is not a defect — a corpus that quotes its tooling is doing
the right thing. What the milestone has actually *paid for* three times is narrower: **the same sentence
in both trees carrying different numbers in the same slot.** So the probe measured that instead.

Four successive readings, each one narrowed by a false-positive class the previous reading exposed:

| reading | predicate | groups |
|---|---|---|
| 1 | any text, whole files | 936 — buried in Go table-test boilerplate |
| 2 | **prose only** (md body outside fences · comments · docstrings) | 90 |
| 3 | + divergent slot must be **interior** to the run | 65 |
| 4 | + **multiset** comparison · shared context on both sides | **39** (7 corpus↔code · 22 corpus↔corpus · 10 code↔code) |

**The corpus↔corpus bucket is dominated by legitimate TEMPLATE reuse** — `roadrunner.md`, `cms.md` and
`jobsimulation.md` share a paragraph and fill in their own ports — so the fence is scoped to the
**corpus↔code** direction the route actually names, and the other two buckets are reported with their
measured sizes rather than silently dropped.

## Phase B — `prose_twin_guard.py`, and the predicate is a parsed construct

The three exclusions were **measured, not anticipated** (`D-M257x-282-3`): an ordinal at a run's edge
belongs to the preceding sentence; an equation written the other way round is agreement; a template
instantiated for two subjects has two correct numbers and no repair that leaves both true. The first two
are excluded by construct; the third is waived **with a recorded reason**, keyed on the sentence template
rather than on `file:line` (`D-M257x-282-4` — an anchored waiver expires silently in the *accept*
direction, which is a ratchet).

**Two tiers, because a silent exclusion is a defect** (`D-M257x-282-5`). RED grades the exit code and
runs to zero; REPORT publishes the shared-context rule's own recall gap. The gap is not hypothetical —
one of the real defects below lives in it.

**Two instrument defects, both found by triage rather than by review** (`D-M257x-282-6`): the
ordinal-leader stripper turned a line-leading `136.5 s npm ci` into `5 s npm ci` and reported a file
disagreeing **with itself**; and one defect was counted once per overlapping window, so the population
was a count of windows wearing a count of defects' clothes.

## Phase C — the repairs, and BOTH copies were stale twice

**Every RED tier finding was a real defect except one, and that one is the waiver.** What the census
found that a reading would not:

| finding | what was true |
|---|---|
| `SCHEME`/`BIND_HOST` anchors into `up-injected.sh` | corpus said `:120`/`:118`, tooling said `:74`/`:76` — **the truth is `:154`/`:146` and NEITHER copy had it** |
| `docker-compose.yml` anchor for sentinel's `search_path` | corpus `:18` is right; `repos_yml.sh` said `:43` |
| `.env.example` coverage in `stack-secrets/README.md` | corpus says **7 of 8**, tooling still said **8 of 9** — the count the `skillpath` decommission moved, corrected corpus-side and never propagated |
| the mutation-battery maxim (`10` vs `18` REDs) | **both correct**, two different batteries → waived with the reason |

Two REPORT-tier findings were repaired in the same pass because they are the same class:

- **`run-playthroughs.sh` still described the DELETED Cosmo router** — `https://<magicdns>:<15050+offset>/graphql`,
  a port *and* a path that M257x iter-13 re-pointed. Squarely this milestone's subject, and it was sitting
  in the tooling's own runbook comment.
- **`up-injected.sh`'s advisory pre-flight anchors** — corpus `:280`,`:320`; tooling `:302`,`:341`; truth
  `:280`,`:335`. **Both copies stale again**, and differently.

> **The pattern the census made visible: in two of four cases neither copy was right.** A reading that
> compared a document against its source would have repaired one side and left the other. A reading that
> compared the two copies would have picked a winner. Only enumerating the *disagreement* and then going
> to the source gets both.

## Phase D — verification, and it is SCOPED — say so rather than implying a section run

The orchestrator re-scoped the milestone mid-iter with a hard wall-clock stop, so the whole-`stack-core`
section run (31–38 min on this box, measured three times last iter) was **deliberately not taken**. What
was run instead is named, with its result, and the gap is stated rather than left to be assumed:

| scope | result |
|---|---|
| `test_prose_twin_guard.py` (new), pytest **and** direct execution | **21 passed** both runners |
| `test_frozen_expectation_census_m257x.py` (the three literal ratchets + the derivation registry) | **102 passed** |
| `test_suite_census{,_collection,_population}.py` | passed inside the 199-passed batch |
| `test_fence_registry_completeness_m257x.py` + `test_fence_registry_population_m257x.py` | **28 passed** |
| `test_guard_family.py -k Reconciliation` (the enrolment arms) | **6 passed** |
| `test_fence_provenance.py -k "stamps or Reconcil"` | **2 passed** |
| `demo-stack/tests/test_public_host_flip.py` (edited docstring) | **23 passed** |
| the three ratchets, measured directly | **240 / 236 / 653 — `exact +0`, and 0 of the three literals is mine** |

**NOT COVERED, stated:** the whole `stack-core` section was not re-run at this HEAD, nor were
`stack-seeding`, `stack-snapshot`, `stack-verify`, `demo-stack` (beyond the one edited module),
`stack-injection` or `playthroughs`. `run-playthroughs.sh`'s edit is a comment.

### Enrolling the guard turned three registries RED, which is those registries working

1. **`FENCE_KIND` + the provenance stamp** were missing — the family's *"a verdict states the tree it
   was taken with"* contract. Added.
2. **The fence-registry's disclosed limit moved** — `24 of 36` → `24 of 37` (`union`), and the two
   `census`/`declaring` figures with it. The **numerator did not move**: `stack-core/README.md` does not
   index the new guard, and it does not index 13 others either, so the honest edit is the denominator
   alone. Published in **both** places the figure lives, as that arm requires.
3. **Two derivations were unclassified** — `collect_sides` and `prose_tokens` — and the completeness
   fence named them by id. Both graded `DECLINE`, `collect_sides` as a **tree-scan on purpose**: a corpus
   doc or a tooling module added tomorrow must enter this fence's population with no edit anywhere, or
   the class reopens exactly where nobody is looking.

**The ratchets were held by rephrasing, never by bumping** (`D-M257x-282-8`): the new test module
breached `TEST_MODULE_LITERAL_CEILING` by +2 with two staged fixture numbers, removed by `%d`-formatting
them out of the literal text.

## Close — 2026-08-11

**Outcome:** `ROUTE-M257x-h70-corpus-and-code-prose-are-copies-with-no-fence` is **fenced and at zero**.
`prose_twin_guard.py` enumerates the corpus↔tooling prose-twin class over **324,549 numeric prose
windows** across **114 corpus files and 763 tooling files**, is enrolled in the guard family, and reports
**RED 0 · 1 waiver with a recorded reason · REPORT-tier 12**. Six real divergences were repaired at
**eight sites in two trees** — including a tooling runbook still describing the **deleted Cosmo router**
— and **in two of the four RED-tier cases NEITHER copy was right**, which is the finding a document-vs-
source reading cannot produce. The population was measured before the fence was designed and narrowed
four times, each narrowing caused by a false-positive class the previous reading exposed; the two
instrument defects found on the way are recorded rather than quietly fixed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**

**Decisions:** `D-M257x-282-1` … `D-M257x-282-9`, including two recorded self-defects (an ordinal
stripper that made a file disagree with itself, and exclusion arms that could not tell exclusion from
refusal) and one escalation held out of scope by the tripwire.

**Side-deliverables:**
- `stack-core/README.md`'s fence-index gap **measured, not repaired**: 13 of 37 family members are
  unindexed. Reported here because the disclosed-limit arm made it visible; repairing it is not this
  iter's scope.

**Routes carried forward:**
- **`ROUTE-M257x-282-prose-twin-REPORT-tier-residual`** — 12 edge-divergence findings the RED tier's
  shared-context rule cannot reach, published by the guard on every run. Several are real (a node
  version, a cockpit offset, a build-cache figure); each needs a source read, which is per-item work.
- **`ROUTE-M257x-282-intra-tree-prose-twins`** — measured and **out of this fence's scope by design**:
  22 corpus↔corpus and 10 code↔code groups. The corpus↔corpus bucket is dominated by legitimate
  template reuse (three service docs sharing a paragraph with their own ports), so a fence over it would
  need the waiver ledger to carry most of its population — a different design, not this one widened.
- **`ROUTE-M257x-282-readme-does-not-index-13-of-37-fences`** — new, from the disclosed-limit arm.
- Unchanged and still open: `ROUTE-M257x-281-rext-tag-SoT-has-no-fence`,
  `ROUTE-M257x-280-the-31-minute-gate-is-skipped-because-it-is-31-minutes`,
  `ROUTE-M257x-281-suite-census-is-structurally-inside-its-subject`,
  `ROUTE-M257x-h70-quotation-verification-instrument-is-unreliable`,
  `ROUTE-M257x-279-durations-are-unclassified-measurement-nouns`,
  `ROUTE-M257x-278-thirteen-unpinned-rext-anchors-are-on-undecidable-clocks`,
  `ROUTE-M257x-274-successor-half-is-uncovered`, `ROUTE-M257x-274-tie-order-is-unstable`,
  `FIX-M257x-269`, `ROUTE-M257x-270-directus-consumer-cms-key`, `FIX-M257x-266`, `FIX-M257x-265`,
  `ROUTE-M257x-h59`, `ROUTE-M257x-h65`, the fence half of `ROUTE-M257x-277`.
- **ESCALATED, not routed:** the demo's production-S3 write path (`D-M257x-282-9`). The next iter owns it.

**Lessons:**
1. **A copy is not a defect; a copy that DISAGREES is.** Fencing the 172-pair copying population would
   have been a fence against quotation. Fencing the disagreement is 4 findings, all real.
2. **In half the cases NEITHER copy was right.** Compare the copies to find the pair, then go to the
   source to settle it — a reading that checks a document against its source repairs one side and leaves
   the other standing.
3. **A waiver keyed on `file:line` expires silently in the ACCEPT direction.** Key it on the claim.
4. **Every false-positive class here was measured, not predicted.** The predicate came out of triage;
   nothing designed up front would have found the ordinal, the reversed equation or the template.
5. **A fixture that cannot tell its two outcomes apart is not evidence for either** — third occurrence
   in three iters, and this time the refusal arm was masquerading as the exclusion arm.
