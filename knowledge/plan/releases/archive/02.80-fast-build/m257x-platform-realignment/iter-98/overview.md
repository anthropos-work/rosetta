---
iter: 98
milestone: M257x
iteration_type: tik
status: closed-fixed
date: 2026-08-06
---

# iter-98 — the REPAIR of the iter-97 reading, by predicate, with paraphrase expansion

**Active strategy reference:** `TOK-05` (*stop repairing claims; fence the predicates under them*).
Unchanged. This tik is `D-M257x-59-1` executed against `FIX-M257x-iter97-read-union`.

## Step 0 — re-survey before targeting

Re-run before committing to the target set, per Phase 1 Step 0:

| precondition | measured at open |
|---|---|
| platform clone == origin HEAD | `0c91421dfdb08dc75f17f1aabfb61394070e770b` == `ls-remote origin HEAD` ✓ |
| guard family | **14 GREEN · 0 RED · 3 not-run** (the 3 need `--range`/`--ledger`) |
| rosetta tree | clean at `a9f8ed4` |
| rext authoring copy | clean, on `main`, `bc2ee74` |
| the 20 anchors | **all 20 re-read line-by-line and still live** — no substitution needed |

`a9f8ed4` (installment 1, pre-iter) had already repaired **5** wrong-construct self-citations. Those are
*not* among the 20 booked anchors — they were sites the iter-97 ledger armed `claim_twin_guard` against.
The 20 stand at full count.

## Cluster / target identified

`FIX-M257x-iter97-read-union`: **20 upheld in-scope BLOCKER anchors / 17 predicates**, plus the named
unbooked twins, plus the out-of-scope-but-real set (rule 44's own false NUL count and its shell recipe,
`platform-alignment.md:1345`, `CLAUDE.md:280`, `safety.md:203/:207`).

## The one strategy choice this iter made, and why

The run brief asked whether to run **twin expansion AHEAD of the read** as a separate pass rather than only
inside repair. **Decision: no — expand twins inside the repair, and spend the new effort on PARAPHRASES
instead.** The reason is a measurement, not a preference:

- iter-96 already did string-twin expansion inside repair and got **13 anchors → 51 sites**.
- iter-97 then measured what ESCAPED that pass: **3 of 51, and all three were paraphrases**, not string
  twins. `claim_twin_guard` was GREEN over all 14 refuted forms while all three were live, because it
  matches **quoted verbatim forms**.

So the string-twin axis is already at roughly zero escape; a separate ahead-of-read pass would re-measure a
solved class. **The 6 %-escape axis is paraphrase**, and that is where this iter put the extra sweep: every
predicate was swept for its *meaning* (`git grep -niE` over paraphrase alternations), not only its strings.
That choice is what found `messenger.md:108`, `frontend-tier.md`'s self-contradiction, and the fact that the
demo academy's whole auth model had changed underneath the corpus (P10 below) — none of which any string
sweep of the booked anchors reaches.

## Hypothesis

Predicate-wise repair with paraphrase expansion discharges all 20 anchors, and the twin multiplier is
**lower** than iter-96's 3.9× because iter-96 drained the wide predicates (`mistralai` alone was 11 sites).

## Expected lift

Clause 5 is met only by a reading that returns **zero**; a repair iter cannot move the gate. The lift is
the discharge of the 20 + the twin reach, fenced rather than asserted.

## Phase plan

Re-derive each predicate from the clones → paraphrase-sweep → edit → re-derive inbound citations after any
line-count change (binding condition 2) → ledger in `claim_ledger.py` shape → guard family.

## Escalation conditions

A predicate whose correction requires an rext CODE change that would resolve the user's open
`DEF-M257x-iter80-storage-prod-bucket` escalation → withdraw the false claim, do **not** resolve the item.
(This fired once, at `safety.md:207`.)

## Acceptable close-no-lift outcomes

n/a — the input set was non-empty and every member was measurable.
