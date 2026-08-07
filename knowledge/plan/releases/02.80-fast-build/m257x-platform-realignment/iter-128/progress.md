# iter-128 — the three diseases, treated separately

**Type:** tik · **Run 81.** The run's directive named three priorities in consequence order and this
document follows that order, not the order the work was easiest in.

---

## 1. `state.md` — back inside its cap, and a budget that has no owner to move to

**The breach was 21 bytes; the defect was 1,085.** `state.md` was 15,381 against a 15,360 cap. Fixing it
by trimming would have bought exactly one run — which is what
[`context.md` § state.md contract](../../../context.md) says in as many words, and why it authored
**per-field budgets** the day before. Measured against those budgets, one field was the whole problem:

| field | before | budget | after |
|---|---|---|---|
| `phase` | **1,985** | 900 | **861** |
| all five others | within | — | within |
| frontmatter | 2,791 | 2,600 | **1,667** |
| **file** | **15,381** | **15,360** | **14,826** |

*"`phase:` is the field that breaks first, every time"* — the contract's own words, confirmed on its first
enforcement. The repair was the contract's method: **move content to its owner.** iter-124's reframe went
to `iter-124/audit.md`; the standing rules to their `§5` and `D-N` owners.

### 1a. The thing the move found — a 38 KB deliverable with no owner

`corpus/architecture/org-repos.md` (38,842 bytes, all 93 org repos, iter-123's net-new deliverable) and
`corpus/ops/observability.md` were **recorded nowhere but `state.md`'s `phase:` field.** iter-123's own
`progress.md` never named either.

`phase:` is the field **every close overwrites**. A deliverable whose sole record is the rotating index is
one close away from being unattributable — and the compression I was about to perform is exactly the
event that would have erased it. `iter-123/progress.md` § 6 now owns both, plus the 8-site `db-backup`
repair. **This milestone's own *"a census is not a disclosure"* rule, pointed at its own plan record.**

### 1b. The finding: the BODY budget is not trimmable by the contract's own method

The body is **13,163 against a 12,000 budget** and it did not come down. Three probes, each asking *"who
else owns this?"*:

| probe | result |
|---|---|
| § Standing backlog | `roadmap-vision.md` (the contract's declared owner) mirrors **7 of 14** items. `PERF-M256-parallel-lane`, `PT-M257-self-evaluation`, `PT-M257-talk-to-data`, `platform-defect-register`, `DEF-M250-01`, `CAVEAT-1`, `PT-M256-resume-fixture-pair` appear **nowhere else** |
| M255 provenance clause | `roadmap.md` carries the *numbers* (4.84 GB → 379 MB, 146.8 → 2.9 s) but **not** the provenance (09:59–11:37Z, PRE-freeze, the 658/666/672 cluster, the three claims owed re-confirmation) |
| § Process flags | the `stackseed`-pin trap, the local-only tag state, the rext code-of-record — **sole owner** |

**In all three, `state.md` IS the owner.** *"Move it to its owner"* has no target. The 12,000-byte body
budget was set without measuring what the body uniquely owns, and it cannot be met by the method the same
contract prescribes — only by trimming prose, which the directive forbids and which would destroy the
sole record of seven backlog items.

**Not silently absorbed:** `state.md` carries a one-line pointer here, and the item is routed as
`FIX-M257x-iter128-body-budget-has-no-owner`. The honest options are to raise the body budget to a
measured floor or to give the un-mirrored content a real owner; **both are decisions, not edits**, so
neither was taken unilaterally.

### 1c. My own defect, caught by a reader and not by any fence

Building the new `phase:` through `.encode().decode('unicode_escape')` mangled a literal em dash into
`â\x80\x94`. Repaired — and repaired as a **repo-wide sweep** over all tracked files for nine
latin-1-round-trip damage sequences (0 other files affected) rather than as a one-character patch.

**No guard in this family checks encoding integrity, and no guard of the existing shape could.** The two
files that mention encoding at all (`corpus_index_guard.py:84`, `story_org_count_guard.py:220`) carry an
`except UnicodeDecodeError` — an *unreadable-file* handler. **Mojibake is VALID UTF-8**: `â\x80\x94`
decodes without error, it is simply the wrong three characters. A decode-error handler cannot see it by
construction.

It was caught by the previous run reading my output. Filed as what it is: a defect class with **no fence
and no possible fence of the current shape**, found by a human-shaped read.

---

## 2. The 820 — triaged, and the printed split corrected by its own audit

Run 80 named the complement of the sealed consequence class **"untriaged, not implied clean."** That was
the right call and this is the debt it created, paid.

**The partition is exact and asserted at runtime**, drawn by the same committed predicate:

```
|C1| 340  +  |complement| 820  ==  |tier-2 population| 1160
```

`iter-128/triage-complement.py` **imports** `iter-124/triage-predicate.py` and `iter-124/triage.py`
rather than copying them, so the two triages cannot drift and iter-124's published figures stay
byte-reproducible. **Anti-vacuity control:** mutating the predicate to match-anything collapses the
complement to 0 — the complement is genuinely predicate-derived, not hardcoded.

| fate | printed | **corrected** | |
|---|---|---|---|
| `cite` | 815 (99.4 %) | **≈ 738 (90.0 %)** | |
| `drop` | 5 (0.6 %) | **≈ 49 (6.0 %)** | colon lead-ins + imperative checklist steps |
| `hedge` | 0 (0.0 %) | **≈ 33 (4.0 %)** | almost entirely `chronos` + `db-backup` |
| `fix` | 0 | **a FLOOR of unknown height** | see the limitation below |

**The correction is measured on this population, not imported.** A seeded 30-of-329 **R4-only** hand
audit puts R4 at **76.7 %** here. All 7 disagreements run one way — `cite` → `drop`/`hedge` — the
direction R4 was declared generous in *before either sample was drawn*. Full table:
[`audit-complement.md`](audit-complement.md).

**Three limitations, and the first is load-bearing:**

1. **`fix = 0` over the complement means nothing was read, not nothing is false.** The triage cannot
   decide falsity. This run's reading was aimed at C1 by the consequence-ordering rule. **The
   complement's false-claim count is UNMEASURED.**
2. **R3 — 486 members, 59 % of the complement — was not re-audited**; iter-124's 100 % was carried
   forward. That is the largest untested block inside the corrected number.
3. iter-124 sampled the whole class, this run sampled R4-only. **66.7 % and 76.7 % are not a
   before/after pair** and must not be quoted as a trend.

**Two mechanical sub-classes named, because a future rule could catch them:** colon-terminated lead-ins
to a code block and imperative checklist steps (all 4 drops); and *"the subject is a repo no clone set
contains"* — `chronos`, `db-backup` (all 3 hedges), which `R2` never learned about because its
`UNCHECKABLE` list is keyed on vendor-internal phrases, not un-cloned first-party repos.
