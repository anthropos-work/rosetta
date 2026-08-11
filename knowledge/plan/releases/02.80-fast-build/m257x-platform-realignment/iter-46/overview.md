---
iter: 46
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-02
---

# iter-46 — `FIX-M257x-iter41-blocker-set`: repair the 18, fence-assisted, by CLAIM

**Active strategy reference:** [`TOK-02`](../decisions.md#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02) — **step 4** of five, verbatim: *"Then repair the 18 once,
fence-assisted — by CLAIM not by FILE (§5 rule 19), tree-wide, with the fence as the commit
post-condition."*

## Step 0 — re-survey before targeting

Platform origin HEAD re-fetched at open: `2adcf71`, **unchanged** (re-scope trigger stays at occurrence 1
of 2). Gate **4 of 5**; clause 5 at **18**. Steps 1–3 are complete and their instruments are live:

| instrument | iter | reaches |
|---|---|---|
| `claim_twin_guard` + `repair_postcondition` | 43, 44 | 16 of 18 detected; 13/13 of the self-contradiction class; RED at commit time |
| `markdown_structure_guard` | 45 | `#6` |
| `anchor_construct_guard` | 45 | `#13`, `#16` |
| `derived_value_guard` | 45 | `#10`, `#11` |

The answer key is captured in **two** perishable fixtures (`fixtures/claim_twin/`, `fixtures/mechanical/`),
so this iteration may finally spend the live one. **TOK-02's named target for this slot is unchanged.**

## Cluster / target identified

All 18, in one pass. The unit of repair is the **CLAIM**, not the file — §5 rule 19, and iter-40 measured
why: **100% of a repair's surviving claim sites pool immediately outside its boundary.** For each blocker
the work is: establish what is true from source, then find and fix **every** site tree-wide that publishes
the refuted form, not only the anchored one.

`claim_twin_guard` names 18 live sites across 11 files, each with its adjudicated verdict and citation.
Two of iter-41's blockers it cannot reach by design (`#10`, and `#17`) are covered by
`derived_value_guard` and by hand respectively.

## Hypothesis

The 18 are repairable in one pass **without inducing the ~9 that every prior pass induced**, because the
dominant induced class (8 of 9: *"repaired at one site, left standing at another"*) is now checked at the
commit rather than at the next audit. The fence cannot make a repair correct; it can only make an
incomplete one unrepresentable in a commit — which is precisely the term iter-41 measured.

## Expected lift

`claim_twin_guard` → **0 live sites**; `markdown_structure_guard` → blocker #6 gone; `anchor_construct_guard`
→ #13/#16 gone; `derived_value_guard` → #10/#11 gone; `repair_postcondition` GREEN with a **lowered**
baseline. Clause 5's own number is **not** measured here — only step 5's full 7-auditor read grades it,
and pre-empting that with a cheaper reading is what iter-38 and iter-21 both measured the cost of.

## Phase plan

1. **Per claim, re-derive what is TRUE from platform source** before writing a word. Every prior pass that
   repaired from the ledger's summary alone induced defects.
2. **Repair tree-wide per claim** — `corpus/**`, `.claude/skills/**`, `CLAUDE.md` — never file-by-file.
3. **`repair_postcondition` at the commit**, plus all four fences run explicitly.
4. **`--accept` the lowered baseline**, with a reason naming this iteration.
5. Two blockers need a decision rather than a transcription: **`#5`'s base count is deliberately
   unsettled** in the ledger (auditor C says 17+7, E says 16+7) and must be **re-measured**, not picked;
   **`#17`** needs an instrument that decides what a sentence claims, so it is repaired by hand
   (`D-M257x-45-3`).

## Escalation conditions

- A blocker whose "what is true" cannot be re-derived from source **is not repaired from the ledger's
  summary** — it is routed forward with a named handler. A repair written from a summary is exactly the
  mechanism that manufactured 9 of iter-41's 18.
- If a repair cannot be made without a platform-repo edit, it escalates: the v2.8 zero-platform-edit
  constraint is not negotiable.

## Acceptable close-no-lift outcomes

A blocker that turns out to be **correct as published** — the ledger refuted five inherited claims by
measurement in iter-01 alone — is a falsification, recorded and counted, not a repair. That is a complete
outcome for that claim, and the fence's waiver mechanism is where it lands.
