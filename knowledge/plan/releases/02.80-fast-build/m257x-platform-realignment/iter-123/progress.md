# iter-123 — the org census follow-ups, re-derived

**Run 79.** The corpus half of the user's three-part instruction (corpus · skills · tooling aligned to
the platform as it is, plus which repos are worth keeping). Skills lane and repo census were already
finished; this is the corpus.

## The substrate, first — because it is why the numbers differ from the census's

Every fact below was re-derived from a **fresh clone at `origin` HEAD taken 2026-08-07**, never from a
stack working tree. The org enumeration is independent:

```
GET https://api.github.com/orgs/anthropos-work/repos?per_page=100&page=1..4&sort=pushed
    -> 93 repositories   ·   72 unarchived · 21 archived · 2 public · 1 fork
```

`app` `3eaadae6` · `infrastructure` `13c248e6` · `directus` `d6325731` · `judge0` `41eb75a3` ·
`metabase` `6f8cd5dd` · `db-backup` `6e1fb15b` (full clone) · `ant-singularity` `5d944e4a` (full).

## 1. What the census got wrong — five corrections, and they are the point

| Census said | Measured |
|---|---|
| `AI-Labs` is a control plane the corpus does not know exists | **The corpus knows** — `ai-labs.md:4-8` names it and records that the doc *previously* denied it. Zero contradicted statements |
| the five `livekit-agent*` repos are undocumented | **Enumerated** in the fenced census table. What was wrong was **two corpus sentences claiming no corpus document names them — self-refuting**, since they name all five while saying it |
| `watermill-redisstream` is upstream of the async backbone | **REFUTED — an inert mirror.** `app/go.mod:12` + `colony/go.mod:9` require **upstream** `ThreeDotsLabs/…v1.4.5`; **no `replace`** in either. The fork still declares the upstream module path, has 0 occurrences of "anthropos", pins `watermill v1.2.0` vs the platform's `v1.5.2` |
| `sim-qa` is a **standing** write-capable path, **unmarked** | Write-capable yes. **Standing: no** (`ls .github` → absent). **Unmarked: no** — `is_test=true` by default (`scenario.ts:156`, `:245`), honoured by `app` (`jobsimulations.graphqls:750`, `manager.go:445`). The stale source was **sim-qa's own README** |
| `analytics-go` absent from the corpus | Absent from the **library model**; `external_services.md:554` already had the right enumeration. **The defect was two corpus files disagreeing with nothing to reconcile them** |

And one it got **right and understated**: `ant-observability` — `git grep -i grafana -- corpus/` returned
**0 files**.

**`ant-singularity` was half-right in the most expensive direction.** The org fix repairs the `git clone`
in § 2 of the onboarding guide (measured `Repository not found` on both https and ssh — a day-one
failure). The other four instances deep-link a file that **has never existed**: 0 additions across 1,046
commits, and `total_count: 0` on an org-wide code search. **Repointing the org would have produced four
URLs that still 404** — a link that looks fixed. Retracted, not repaired.

## 2. The finding that outranks the repo list

Cloning `infrastructure` — never in any clone set, so never read — settled **four** standing questions
with one rule:

> **A service repo's own `service_desired_count` is not evidence of production state.** It is an input
> to a module that must be *instantiated* by `infrastructure/terraform/production/services.tf`.

`grep -n '^module "' terraform/production/services.tf` → **exactly ten** declarations. Four repos this
corpus quotes declare a count that instantiates nothing: `cms:39`, `roadrunner:19`,
`graphql-wundergraph:20`, `messenger:29`.

- **`cms` M810 — RESOLVED, DESTROYED.** The standing *"report both, assert neither"* is closed.
- **`graphql-wundergraph` — THIS CORPUS WAS WRONG**, in a fenced table, and the *other* org corpus was
  right. ECR hand-deleted 2026-08-05.
- **`roadrunner`** — settled; no ECS service exists.
- **`jobsimulation` + `storage`** — blocks survive **deliberately**, assets-only, and infrastructure says
  why: `prevent_destroy` is read from **configuration**, not state.

## 3. The re-pin backlog — the census's dominant class is 85 % its own substrate

**Denominator: 89 rows** (68 `UNRESOLVABLE` + 17 `PARTIAL` + 4 `DOES-NOT-SUPPORT`, the raw verdict TSVs;
the adjudication's deduped figures are 57/17/3). Each row's citation was re-resolved **at a sha the
corpus block itself names**, against full-history clones, with the corpus read at `afe58ac` so the
backlog's own anchors are read at the ref they were taken at:

| | count | what it means |
|---|---|---|
| **RESOLVES-AT-PIN** | **76 / 89 = 85.4 %** | the citation is correct and permanently verifiable. The census could not decide it **only** because its resolver read a clone's working tree instead of the pinned ref |
| NO-SHA-IN-BLOCK | 7 | the real re-pin class — an unqualified `path:line` that drifts silently |
| PIN-DOES-NOT-RESOLVE | 6 | candidate rot; several are blocks citing a second repo, not decay |

**So the backlog is at most 13, not 74** — and *"11.8 % are pinned to commits the tree is not at"* is a
statement about the census's substrate, not about the corpus. **This is the milestone's own
stale-substrate rule, one level up.** Re-pinned this iter: **3** (`clerk-integration.md` 461→468 —
caught by `repair_postcondition`, not by inspection; the two `observability.md` anchors). Retired: **0**
— no anchor in the 13 has a third-generation history, so iter-115's retire-rather-than-re-derive
precedent does not fire. **The other 76 need no work, and saying so is the deliverable.**

## 4. The hunted-sample rule — checked, and there was nothing to repair

`D-M257x-122-3` is already carried correctly in `state.md` (`phase:` and the Chapman paragraph). Measured:
the **corpus publishes no error rate at all** — `git grep -E '[0-9]+% of (the )?(corpus|claims|citations)'
-- corpus/ CLAUDE.md` returns nothing. No place in the record quotes a hunted rate as corpus-wide.

## 5. Guards

**22 members · 17 GREEN · 0 RED · 1 could-not-check · 4 not-run** after repair (from **15 GREEN · 2 RED**).
**Not a whole-family green, and the runner's own summary says so** — the family exits **2**, not 0.

> **This line said "20 GREEN" until the closing re-run.** I wrote the total from arithmetic on the
> before-state instead of reading the after-run's own summary — the exact substitution of inference for
> measurement this milestone exists to catch, committed in the paragraph reporting the guards. Corrected
> against `guard-family: 17 GREEN · 0 RED · 1 could-not-check · 4 not-run`.
Invocation: `guard_family.py --repo-root <rosetta> --platform <rosetta>/stack-demo/platform`.

- `platform_alignment_guard` is the **could-not-check**: 5 citations to `infrastructure` / `db-backup`
  resolve in no clone set. That is the honest verdict and it is new — the corpus now cites a repo no
  stack clones.
- `unreadable_repo_claim_guard` **extended**: a mention passes on an unmeasurable marker **or** a
  ref-pinned reading (repo **and** sha, as a conjunction). Prints the split; live **9 = 4 hedged + 5
  measured**. 12 → 18 tests; the `and`→`or` mutant turned both new controls RED (2 failed / 16 passed)
  and restoring returned 18/18.
- `claim_census_guard` ratchet: **7 existing files went DOWN (−13 unevidenced assertions)** — the repairs
  added evidence — and the one new file added **27**. Rebaselined, delta disclosed. **Zero regressions
  across the 20 existing files edited.**
- Three guards caught three defects that inspection did not: a blank-line anchor, a wrong-construct
  anchor, and — the sharpest — **the same anchor a second time because the fix named two refs in one
  cell**, which makes every anchor in that block ambiguous to the resolver (M257x run-53).

## 6. Gate

**Unchanged at 4 of 5.** No reading was taken; `P` is unmeasured. Clause 5 is met only by a reading that
returns zero, and repair is not a reading — this iter removes confounds from one, as iter-40 recorded.
