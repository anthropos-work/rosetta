# iter-84 — progress

**Type:** tik, under `TOK-05`. Deliverables: the adjudication of iter-82's union, the **per-anchor**
ledger, and the eleven discharges re-derived as membership questions. **No repair landed** —
adjudication-before-repair stays binding (iter-80).

---

## THE HEADLINE

**40 of 43 upheld — 93.0 %.** The instrument did not regress across iter-81's repair (iter-80 measured
92.1 % on the pre-repair union). And the eleven discharge verdicts, re-derived by membership: **5 stand,
4 refuted, 1 unsettled, 2 folded into the adjudication.**

And a third finding, **corrected after a peer review** (`D-M257x-84-6`) — it was first written as a
"reach limit" on the instrument, which was wrong:

> **Clause 5's declared scope is `corpus/services/**` + `corpus/architecture/**` = 40 files, and the
> instrument reads 40 of 40 — COMPLETE. But of P4's live members, exactly ONE is inside those 40; the
> other ≥16 are in `corpus/ops/**`, `CLAUDE.md` and `.claude/skills/**`.**

That is a **scope observation and a corpus-quality finding — not a defect in the instrument.** Clause 5 is
narrower than the corpus *by its own wording*; the whole corpus is **90** `.md`
(`git ls-files -- 'corpus/*.md'`), of which `corpus/ops/**` is **46**. Wanting broader coverage would be a
re-cut of clause 5, which is **not on the table**. The real limiter on what a zero establishes is the
~50 % per-pass **recall** iter-83 measured — a within-scope property.

---

## 1. The adjudication — 40/43 upheld

Full per-anchor ledger: [`adjudication.md`](adjudication.md).

| | A | B | C | D | total |
|---|---|---|---|---|---|
| anchors | 11 | 11 | 11 | 10 | **43** |
| UPHELD | 10 | 10 | 11 | 9 | **40** (31 blocker · 9 minor) |
| REJECTED | 1 | 1 | 0 | 1 | **3** |
| UNSETTLED | 0 | 0 | 0 | 0 | **0** |

**Pre-registered ≥ 70 % — HOLDS at 93.0 %.** The falsification condition (collapse below 50 %, meaning
the post-repair signal is mostly noise) did not fire.

**Counts re-derived from the per-anchor verdicts, never from an adjudicator's summary line** (§5 rule
32) — and it paid immediately: **adjudicator B's summary said "9 UPHELD (7 blocker, 2 minor)" while its
own eleven verdicts give 10 UPHELD (8 blocker)**. Second occurrence of that rule firing in this run.

### The 40 by predicate — iter-85's work list

| class | n |
|---|---|
| **Q1** stale cross-repo line anchor | 13 |
| **Q2** present-tense claim about a **deleted** fact (re-anchoring is NOT the repair) | 7 |
| **Q3** wrong scalar / wrong set against source | 8 |
| **Q4** wrong predicate — no line-checker could catch it | 7 |
| **Q5** the booked anchor is CORRECT; the false statement is elsewhere | 1 |

### The three rejections are one mechanism, and the escalation condition FIRED

All three are `CHECK-M257x-iter76-seat-ref-discipline`, now at its **4th and 5th occurrences**
(`CLAUDE.md:203` — clones 6 and 3 commits behind `origin/main`; `hiring.md:73` — the doc's banner grounds
it at `app 5ba17044` and both anchors are byte-exact there). This iter's `overview.md` declared a 4th
occurrence escalated rather than absorbed. **It is escalated.**

**Adjudicator D found why the rule keeps failing: it is stated wrong.** Seats are not ignoring it, they
are applying it *unevenly*, because *"grade at the ref the claim names"* says nothing about a sentence
that asserts **currency**. The correct form:

> **Grade at the ref the claim names — UNLESS the sentence asserts currency, in which case no
> neighbouring pin rescues it.**

That single amendment separates `graphql-wundergraph.md:13` (**UPHELD** — *"survives"*, *"is now"*, under
a column headed *"origin HEAD"*) from `hiring.md:73` (**REJECTED** — a static citation under a
document-scope grounding banner). Routed as a §5 rule-33 amendment.

**Adjudicator A's structural fix:** the ground-truth table handed to seats must carry each clone's
**`origin/main` sha beside its checkout sha**. A seat given only the checkout cannot see it is stale —
which is exactly how occurrences 3, 4 and 5 happened.

---

## 2. The eleven discharges, re-derived as membership questions

Full sweep: [`membership.md`](membership.md).

| verdict | predicates |
|---|---|
| **stand** (no live member on any swept surface) | P2 · P3 · P5 · P6 · P10 |
| **refuted** | **P4** (≥17 live members in the published tree + 7 in rext source, vs "~10 discharged") · **P9** (1) · **P11** (≥3) |
| **unsettled** | P1 — clean inside the instrument's 40 files; ungraded candidates on the `corpus/ops/**` surface, recorded as *unsettled* rather than *clean* |
| deferred to the adjudication | P7 · P8 (grading them twice would double-count) |

**Five of eleven stand. That is worth saying plainly: the repair's *content* was largely correct where
it reached.** The failure was never competence; it was that four verdicts were **reported complete
without being measured**.

### Why the five that stand are not evidence the criterion was sound

P1 was the dominant predicate (~47 sites) and its members sat in `corpus/services/**` and
`corpus/architecture/**` — **inside the file set the repair's seats owned**. P4's members are mostly in
`corpus/ops/**` and `.claude/skills/**`, which **no seat owned**, because iter-81 inherited its partition
from the **read's** 40-file partition.

**§5 rule 19 says the partition that is correct for reading is wrong for repairing.** iter-81 was
partitioned by the reading's file set anyway. That is the fourth mechanism, and it is one iter-83 could
not see from the diff alone: iter-83 measured *reach against the ledger* (74.1 %); this measures *reach
against the predicate*, and for P4 **the ledger itself only ever covered 1 of ≥17 members** — because the
other 16 sit outside clause 5's declared scope, which no seat was ever asked to read or repair.

### 🔴 A live tooling defect, not a documentation defect

**`rosetta-extensions/dev-stack/dev-stack:186` and `:414` initialise `profile="graphql"`.** A bare
`dev-stack up N` — what `/dev-up N` runs with default flags — executes `docker compose --profile graphql
up -d`, and that token **exits 0 and starts only the always-on floor**. Postgres answers, `docker ps` is
non-empty, the application is absent. `SKILL.md:147` documents the default faithfully; **the tool is
wrong.** Routed `FIX-M257x-iter84-dev-stack-default-profile`, severity **high**, live proof required
rather than a source reading.

---

## 3. The three `repair_leak_guard` sites

Full grading: [`leak-sites.md`](leak-sites.md). **2 real, 1 benign — 67 % precision on a commit six other
guards passed clean.**

- `CLAUDE.md:285` — **UPHELD blocker.** `(graphql profile)` in a **runnable command block**, 16 lines
  above the same file's own warning about that exact hazard.
- `corpus/ops/platform-alignment.md:1305` — **UPHELD blocker.** *"`STORAGE_RPC_ADDR` is read by
  `main.go`"* at `9d00a313`: measured **3 hits, all comments, 0 env lookups**. It describes a middle
  state that never existed. **The protocol doc contradicting the rule it teaches.**
- `corpus/services/messenger.md:122` — **REJECTED, benign.** The header says `Value (compose)` and
  compose sets `4`. A shared-vocabulary shingle match with the `roadrunner.md` twin iter-81 legitimately
  rewrote.

---

## State at close

- **Gate 4 of 5, unchanged.** No reading was taken; clause 5 moves only on one that returns zero.
- 6 corpus guards exit 0 at open and at close. `stack-core` unchanged at **843** (1 known perishable).
- Zero platform-repo edits. `storage.md:55,:154,:181` **held** (`DEF-M257x-iter80-storage-prod-bucket`).
- `FIX-M257x-iter53-union-set`, `FIX-M257x-iter56-assignment-flake`,
  `CHECK-M257x-iter38-ai-act-classification`, RF-2/3/7–14 untouched.
- `CHECK-M257x-iter82-commit-message-narration` stays **separate** from `CHECK-iter77`.

## Close — 2026-08-05

**Outcome:** iter-82's union adjudicated **40/43 upheld (93.0 %)** with a **per-anchor** ledger; the
eleven discharges re-derived by membership (**5 stand, 4 refuted, 1 unsettled, 2 deferred**); and clause
5's instrument measured at **40 of 40 of its declared scope — complete**, with **≥16** live P4 members
sitting *outside* that scope (a corpus-quality finding, **not** an instrument defect — corrected after
peer review, `D-M257x-84-6`). One live rext defect found. **No repair landed, by design.**

**Type:** tik
**Status:** closed-fixed — all three declared lines landed
**Gate:** NOT MET — 4 of 5, unchanged
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue

**Decisions:** D-M257x-84-1 … D-M257x-84-5 (iter-84/decisions.md)

**Side-deliverables:** none.

**Routes carried forward:** `FIX-M257x-iter85-repair-the-40` (the Q1–Q5 work list) ·
`FIX-M257x-iter85-p4-membership` (≥17 corpus + 7 rext members) ·
`FIX-M257x-iter84-dev-stack-default-profile` (**high**) ·
`FIX-M257x-iter83-leak-guard-3-sites` (2 of 3 confirmed) ·
`CHECK-M257x-iter84-rule33-currency-amendment` · `CHECK-M257x-iter84-ground-truth-needs-origin-sha` ·
`CHECK-M257x-iter84-defects-outside-clause5-scope` (corpus quality; **NOT** a clause-5 re-cut) ·
`CHECK-M257x-iter83-standalone-is-the-forgettable-class`

**Lessons:** (1) adjudicate-before-repair earned itself again — `ai_architecture.md:225` is **correct**
and repairing it would have broken a true sentence while leaving the false one standing in a different
file. (2) A rule that fails five times is usually **stated wrong**, not ignored. (3) Re-derive the
hand-off's numbers *including the orchestrator's* — it fired twice in one run.
