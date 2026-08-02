# iter-43 — `FENCE-M257x-iter42-claim-twin` watched going RED

**Measured against the corpus at rosetta `48ca53c`** — byte-identical to the tree iter-41 measured
(`git diff 103ad31..HEAD -- corpus/services/ corpus/architecture/` empty at open). **Nothing was
repaired in this iteration**, so the answer key is the one iter-41 recorded, unmodified.

Command:

    cd .agentspace/rosetta-extensions/stack-core && python3 claim_twin_guard.py

    claim-twin-guard: derived 36 claims (39 refuted forms) from 85 blocker rows in 4 ledger file(s);
                      17 row(s) quoted nothing longer than 30 chars, 32 row(s) quoted no refuted form at all
    claim-twin-guard: scanned 112 published file(s); 3 acknowledged site(s) skipped
    claim-twin-guard: RED — 18 published site(s) restate a claim an audit already refuted
    EXIT=1

---

## THE HEADLINE: 16 of the 18 blockers detected, at the anchors the audit recorded

| # | iter-41's anchor | fired at | class (iter-42) |
|---|---|---|---|
| 1 | `ai_architecture.md:104-105` | `ai_architecture.md:104` | self-contradiction |
| 2 | `ai_architecture.md:38-45` | `ai_architecture.md:38` | self-contradiction |
| 3 | `ai_architecture.md:59` | `ai_architecture.md:59` | self-contradiction |
| 4 | `security_compliance.md:175-176` | `security_compliance.md:175` | self-contradiction |
| 5 | `security_compliance.md:76`, `:83-84` | `security_compliance.md:76` | self-contradiction |
| 6 | `security_compliance.md:205` | `security_compliance.md:205` | self-contradiction |
| 7 | `service_taxonomy.md:288-289` | `service_taxonomy.md:287` | self-contradiction |
| 8 | `service_taxonomy.md:145` | `service_taxonomy.md:145` | self-contradiction |
| 9 | `service_taxonomy.md:136` | `service_taxonomy.md:136` | self-contradiction |
| **10** | `sentinel.md:12` | **not detected** | derived scalar |
| 11 | `sentinel.md:22` | `sentinel.md:22` | derived scalar |
| 12 | `graphql-wundergraph.md:79` | `graphql-wundergraph.md:79` | self-contradiction |
| 13 | `external_services.md:788` | `external_services.md:788` | wrong construct |
| 14 | `ai-readiness.md:37-43` | `ai-readiness.md:42` | self-contradiction |
| 15 | `roadrunner.md:23-25` | `roadrunner.md:23` | self-contradiction |
| **16** | `messenger.md:110` | **not detected** | wrong construct |
| 17 | `platform-migration-status.md:60` | `platform-migration-status.md:60` | wrong construct |
| 18 | `security_compliance.md:7` · `architecture_overview.md:243` | **both** | self-contradiction |

**13 of 13 of the class the fence was built for.** Plus 2 of the 3 wrong-construct blockers and 1 of the
2 derived scalars, which is more than its charter promised.

Anchors resolve exactly except #7 and #14, which fire one and five lines *earlier* — the normalized
match names the line the claim **starts** on, and in both cases the sentence begins on the line above the
one the auditor cited. The RED names the right block; it is not off-target.

### The two misses are the declared scope boundary, not a shortfall against it

- **#10** *"**Language**: Go 1.25"* — the quoted form normalizes to 17 characters, under the 30-character
  fragment floor, so the ledger row is reported `UNMATCHABLE` **by name** rather than dropped. iter-42
  assigned it to a **value fence** (TOK-02 step 3).
- **#16** `messenger.md:110` — iter-41's claim column *paraphrases* (*"skill-path read cited at
  `assignments.go:815`"*) instead of quoting, so there is no pattern to derive. Counted in the coverage
  line as one of the 32 rows that quoted no refuted form. iter-42 assigned it to a **symbol-aware anchor
  check** (TOK-02 step 3).

> **Both misses land in the two other instruments' territory, and neither is random.** That is the
> result to carry: the fence's blind spots are the ones its charter predicted, which is what makes the
> three-instrument split from iter-42 a plan rather than a hope.

**Cost recorded, not glossed:** the minors exclusion is what makes #16 unreachable here. `m-E3` in
iter-34's `audit-e.md` recorded that exact defect **seven iterations before** it was promoted to a
blocker, and a fence over minor rows would have caught it. That trade is taken deliberately — see below.

---

## The 18th site is a claim NO pass has caught, and it is outside the audited scope

    corpus/ops/demo/coverage-protocol.md:629
        refuted form : "the nil-CycleID default is hardcoded to buildLiveResponse"
        adjudicated  : iter-34/progress.md:52
        verdict      : readiness.go:307-312 takes the frozen path on the no-active/has-closed shape
                       — exactly what M51 seeds

iter-34 refuted this claim and it was repaired in `ai-readiness.md`. **The identical sentence survived in
`corpus/ops/**`** — outside clause-5 scope, and therefore invisible to all six passes. This is §5 rule
19's corollary (*"a claim leaks to the EDGE of the previous repair's scope and stops there"*) measured
rather than argued, and it is the fence earning its keep on its first run.

**Not adjudicated here.** Verifying it needs the `app` repo, which is not cloned. It is routed as
`FIX-M257x-iter43-coverage-protocol-livepath`, out of clause-5 scope, and it is **not waived** — the
fence stays RED on it. *(This is the same shape as iter-41's `G6`.)*

---

## The GREEN control, which is a real measurement and not a synthetic one

**36 claims derived; 20 fired; 16 adjudicated claims are ABSENT from the entire tree.** Those 16 are
claims iters 33/34/38/39 refuted and earlier passes genuinely repaired everywhere — the fence says so
positively, in its coverage line, rather than by silence.

Plus a synthetic control that survives the repair: `tests/fixtures/claim_twin/green/` holds a twin of
every one of the 18 sites with the offending line removed. **Zero fire.** Without it, "all 18 fired" is
equally consistent with a fence that fires on anything.

## The three acknowledged sites, and why acknowledging them is safe

| site | why |
|---|---|
| `ai_architecture.md:7` | the page's own *"⚠️ The often-repeated claim that "simulation scoring is NOT done by AI" **is false at platform HEAD**"* |
| `service_taxonomy.md:139` | *"(This page previously called Ant Academy "fully independent of the backend"; that framing was **retired** at v2.5 M231…)"* |
| `skillpath.md:81` | inside an explicit *"⚠️ RETRACTION — this bullet previously said the opposite"* blockquote |

A waiver is **two keys**: the acknowledgement in `claim_twin_waivers.json`, **and** the guard's own
re-confirmation that a retraction marker is still within one block of the quote. Delete the retraction
and leave the sentence standing, and the waiver stops applying on the next run. The detector agreed with
the human on all four candidates — the three above read `retracted_context=True`, and
`coverage-protocol.md:629` reads **False**.

And the real control against Trap A is elsewhere: `test_01` asserts all 18 answer-key sites still fire,
so a waiver that suppressed one turns the suite RED.

---

## The cut that made the fence usable: `## Minors` sections are excluded

First run, before the cut: **33 hits**, of which 17 came from minor sections and **12 of those 17 were
prose that is TRUE** with an anchor off by a few lines (eight in `ai-readiness.md` alone). A fence that
reddens on correct sentences is disabled within a week — §8 rule 6, which this milestone has already
paid for once.

The cut is also correct by construction rather than merely convenient: a minor row records *"the claim
itself is TRUE"*; a blocker row records a claim that is **false**. They are different kinds of finding,
not two severities of one kind, and iter-42 routed them to different instruments.

## The mutation battery — 8 mutants, all matching their DECLARED verdict

| mutant | declared | actual | killed by |
|---|---|---|---|
| `noop-comment` (**positive control**) | **GREEN** | **GREEN — survived** | — |
| `minors-exclusion-inverted` | RED | RED | `test_01`, `test_04`, `test_07`, `test_13` |
| `elided-quote-becomes-disjunction` | RED | RED | `test_05` |
| `waiver-ignores-the-retraction` | RED | RED | `test_11` |
| `pure-anchor-filter-inverted` | RED | RED | `test_01`, `test_05`, `test_06`, `test_09`, `test_10`, `test_13` |
| `quote-fold-removed` | RED | RED | `test_01`, `test_09` |
| `fragment-floor-collapsed` | RED | RED | **`test_02` — the GREEN twin** |
| `empty-derivation-reports-clean` | RED | RED | `test_15` |

**Five of the seven kills are inversions, not deletions** (§8 rule 5: removal-mutants could not have
found iter-27's flipped guard). **Seven distinct failure signatures**, so the suite is discriminating
rather than reporting a constant. The baseline is GREEN before *and* after the battery. Every mutant is
`py_compile`d before its run, so a compile break can never be misread as a kill.

`fragment-floor-collapsed` is the one worth naming: loosening the floor from 30 to 3 characters leaves
**every answer-key site still firing** — only the GREEN twin catches it. A fence measured solely by REDs
would have shipped it.

## Suite

`stack-core`: **415 tests, 14 failures** — the 14 are the pre-existing pytest-dependent batteries
(`test_m220_*`, `test_m255_*`; no pytest on this host). Baseline was **396 / 14F**; this iteration adds
**+19 tests and 0 new failures**. The new battery runs on `unittest`, so it actually executes here.
