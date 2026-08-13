# iter-99 adjudication — readings #21 + #22 at platform `0c91421`, corpus `e858fd4`

**Status: COMPLETE.** All **46** booked blockers across 14 seats graded by **four parallel adjudicators**,
one per seat-group, each re-deriving from the platform/service clones rather than from any seat's evidence
or any prior verdict.

## Verdict

| | Adj1 · r21-A..D | Adj2 · r21-E..G | Adj3 · r22-A..D | Adj4 · r22-E..G | **total** |
|---|---|---|---|---|---|
| booked | 14 | 11 | 13 | 8 | **46** |
| **UPHELD** | 10 | 9 | 10 | 7 | **36** |
| REJECTED | 4 | 2 | 3 | 1 | **10** |
| in-scope upheld BLOCKERS | 9 | 9 | 9 | 7 | — |

**Upheld rate 78.3 %** — against 92.1 % (iter-80), 93.0 % (iter-84), 92.7 % (iter-95), 93.1 % (iter-97).
**This is the first adjudication to fall outside the 1.0-point band four readings had held**, and it is a
pre-registered failure (#4, band [88 %, 96 %]). See *What the precision drop means* below; it is the most
consequential number in this reading and it is **not** self-evidently a fact about the corpus.

## **N = 28** distinct in-scope upheld BLOCKER anchors

`n₁ = 18 · n₂ = 16 · m = 6` → union **28**.

| quantity | iter-95 | iter-97 | **iter-99** |
|---|---|---|---|
| booked (14 seats) | 55 | 58 | **46** |
| upheld | 51 — 92.7 % | 54 — 93.1 % | **36 — 78.3 %** |
| rejected | 4 | 4 | **10** |
| **graded N** | 13 | 20 | **28** |
| Chapman N̂ | ≈ 16.6 | ≈ 29.3 | **≈ 45.1** |
| per-pass recall | 60 / 42 % | 41 / 44 % | **39.9 / 35.4 %** |
| union recall | ≈ 78 % | ≈ 68 % | **≈ 62 %** |
| estimated still unfound | ≈ 4 | ≈ 9 | **≈ 17** |

**Clause 5 is NOT met. The gate stays 4 of 5.**

### Matched by both readings (6)

`ai-readiness.md:305` · `ai-readiness.md:46` · `backend.md:33-34` · `dependency_map.md:59` ·
`hiring.md:38` · `service_taxonomy.md:405`

### The 10 rejections

Four are the **ref-discipline** class (now **17 occurrences over five readings**, still contributing zero to
any graded count): `chronos.md:27` (8080/8081 are the binary's defaults, compose's 8500/8501 an override),
`backend.md:236` (true at the PR the dated bullet itself names), and — booked **independently by both
readings and rejected by both adjudicators** — `external_services.md:208-211`, where the rext anchors
resolve byte-exact in the **pinned per-stack clone** `ab81527a` and the seats graded them against the
authoring copy. That last one is worth naming: **two seats made the same wrong-tree error**, which is a
property of the briefing, not of the corpus.

The other six are mis-reads (`README.md:21` echoes the fenced map's own prod column; `backend.md:50-52` and
`:18-19` where the two "fives" answer different questions; `security_compliance.md:158` reconcilable;
`security_compliance.md:185`; `hiring.md:80-82` exact at the doc's declared ref for **one** adjudicator).

## The pre-registration graded — **4 of 9 held, 5 FAILED**

| # | prediction | band | outcome |
|---|---|---|---|
| 1 | per-reading count | [8, 18] each | **HELD** — 18 and 16 |
| 2 | union `N` | [12, 26] | **FAILED** — **28**, above |
| 3 | recurrence of the 21 iter-98 predicates | ≤ 2 | **HELD** — 0 true recurrences |
| 4 | upheld rate | [88 %, 96 %] | **FAILED** — **78.3 %** |
| 5 | per-pass recall, ≥1 pass below 41 % | [28 %, 55 %] | **HELD** — 39.9 % and 35.4 %, **both** below |
| 6 | platform-drift share | ≤ 10 % | **HELD** — ~1 of 28 |
| 7 | repair-induced upheld blockers | [0, 3] | **HELD** — **2** |
| 8 | mean sites/predicate of what is booked | < 2.5 | **HELD** — 28 anchors / ~24 predicates ≈ 1.2 |
| 9 | wrong-construct intra-corpus citations | ≤ 1 | **FAILED — badly.** At least **7** |

**Five of nine failed, and that is the pre-registration working.** iter-95 graded 6 of 6 and booked it as a
warning; iter-97 graded 3 of 7; this is 4 of 9 with the failures carrying the content again.

## The three findings that outrank the defect list

### 1. Band #9 failed by ~7×, and it is a finding about the INSTRUMENT, not the corpus

`anchor_construct_guard` was **GREEN at the commit under audit** — *"every resolvable anchor names a
construct"* — while the readings booked and both adjudication panels upheld at least **seven** citations
that resolve to the wrong construct: `ai-readiness.md:46` (a self-citation offered AS evidence landing on a
**blank line**), `ai-readiness.md:305` (`urls.ts:52` is `ORGANIZATION_FEEDBACK_URL`; the target is `:50`),
`graphql-wundergraph.md:134` (`:84` is the Ports bullet), `hiring.md:38` (twin drifted to `:52`),
`hiring.md:80-81` (`manager.go:485` is a closing brace), `messenger.md:53` (unpinned anchors resolving at no
named ref), `jobsimulation.md:203-204` (a ref *labelled* origin/main that is not).

The word doing the work in the guard's own green is **"resolvable."** The pre-registration set this band at
≤1 *precisely because* the guard now scans the whole anchor set, so an upheld member is evidence of a blind
spot. It found one, and this is the highest-value routed item in the reading.

### 2. Precision fell 93.1 % → 78.3 %, and the honest reading is NOT "the seats got worse"

Rejections rose 4 → 10 while bookings **fell** 58 → 46. Three candidate mechanisms, and this reading cannot
separate them:

- **The residual got harder.** iter-98 drained the wide, obvious predicates ([`iter-98/discovery-pool.md`](../iter-98/discovery-pool.md) measured max width 11 → 4). What is left is narrow and subtle, so a
  marginal booking is more likely to be wrong. Under this reading the precision drop is a *consequence* of
  the corpus improving.
- **A briefing gap.** Two independent seats made the identical wrong-tree error on the rext anchors. That is
  not seat noise; that is the instrument under-specifying which rext clone grades a claim.
- **Adjudicator variance.** `hiring.md:80-82` was **REJECTED by Adj2 and UPHELD by Adj4** on the same
  underlying construct. One disagreement in 46 is small, but it is non-zero for the first time.

**Recorded, not resolved.** Asserting the flattering explanation would be exactly the failure this milestone
exists to prevent.

### 3. Two of the 28 were INDUCED by iter-98's own repair, and both are in what it rewrote

- **`dependency_map.md:59`** — iter-98 pinned this cell to `b948604f` and wrote *"`SKILLER_STREAM` has 6 Go
  occurrences across **4** files."* Both readings measured **3 files**. The repair that fixed a two-ref
  ambiguity introduced a wrong file count in the same sentence.
- **`backend.md:33-34`** — iter-98 rewrote this to remove `skiller` from the both-ends set. Both readings
  now book the *rewritten* line: the set is **five, not four** — `backend` itself was omitted while the
  sentence asserts exhaustiveness.

Prediction #7 forecast [0,3] induced and got exactly 2. **The mechanism model keeps holding while the
magnitude model keeps failing** — the same split iter-97 recorded.

## What this reading establishes, and what it does not

**Establishes:** the corpus carries **at least 28** blocking falsehoods inside clause 5's scope at
`e858fd4`, with ≈17 more estimated unfound; the platform-drift class M257x was created to fix is ~1 of 28;
`anchor_construct_guard` has a demonstrated blind spot on intra-corpus citations; and iter-98's repair
induced exactly 2 upheld blockers, both inside prose it rewrote.

**Does not establish:** that `N` is rising. `N` went 13 → 20 → 28 across three readings **whose upheld rates
were 92.7 / 93.1 / 78.3 %** — the instrument's precision is no longer constant, so the series is no longer
comparable on the axis iter-97 relied on. [`iter-98/discovery-pool.md`](../iter-98/discovery-pool.md) §3
predicted recall would fall as the pool narrowed, and it did (union 68 % → 62 %, both passes below 41 %,
band #5 held). **A narrowing pool measured by a degrading instrument produces a rising `N`.** That is
consistent with this data and so is a genuinely growing residual; **this reading cannot tell them apart, and
does not claim to.**

**Comparability:** continuous in **instrument** (briefing byte-identical, sha `3858ec53…`, one commit ever;
14 seats; same partition method, grading rule and scope) and **discontinuous in upheld rate for the first
time in five readings**. The count is on the same basis as iter-95/97 but taken on a different tree.
