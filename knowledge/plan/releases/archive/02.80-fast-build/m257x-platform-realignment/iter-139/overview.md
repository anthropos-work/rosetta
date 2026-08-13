---
iter: 139
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
---

# iter-139 — audit the instrument before acting on it at scale

**Type:** tik
**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* iter-138
produced the census (**127 rotted pins / 222 decidable**). `TOK-08` says work the class; **this milestone
says verify the instrument first**, and it has said so nine times over.

## Step 0 — Re-survey before targeting

Re-ran the probe at HEAD (`461a4c4`): **127 rotted across 22 files**, unchanged from iter-138's reading
(iter-138's own 9 repairs removed their pins rather than moving them). Full census checked in as
`rot-census.txt`. Concentration: `org-repos.md` 16 · `ai-readiness.md` 14 · `ant-academy.md` 14 ·
`platform-alignment.md` 11 · `cms.md` 9 · `hiring.md` 9 · `jobsimulation.md` 8.

## Cluster / target identified

**The probe has never been audited, and iter-138 published its number.** That is the exact shape this
milestone has caught eight times: *an instrument's output treated as measurement before the instrument
was measured.* Three specific reasons to distrust it before repairing 127 sites from it:

1. **The largest deltas are implausible as same-file rot** — `+1557`, `+820`, `+295`. The probe's own
   declared floor is that it **cannot distinguish a same-file pin from a cross-file continuation pin**
   (`` `main.go:507`, `:509` ``), and a continuation pin resolved against the wrong file will match
   *some* line and be scored as rot.
2. **Unique-text matching is not unique-construct matching.** A repeated sentence, a moved-and-restored
   block, or a line whose text is coincidentally identical elsewhere all score as rot.
3. **`git blame` on the CITING line dates the line's last edit, not the citation's authorship.** A
   reflow, a typo fix, or one of this milestone's own repairs re-dates the line and changes which
   historical file the probe compares against.

## Hypothesis

**A meaningful fraction of the 127 are false positives, and the true rot count is lower than published.**
If so, iter-138's number needs a correction *before* `FIX-M257x-iter138-anchor-rot-fence` is built
against it as a baseline — a fence sized to a wrong denominator is the `D-M257x-134` failure again.

## Expected lift

No `N` movement claimed (no reading this iter). Deliverable: **a measured precision for the probe, with a
confidence interval and a stated denominator**, and iter-138's published figure either **upheld or
corrected in place**.

## Phase plan (declared 2-step shape)

1. **Priority 1 — precision audit.** A **stratified** sample of the 127 (large / medium / small delta,
   since the failure modes are delta-correlated), each opened by hand at the citing line and the target,
   classified `TRUE-ROT` / `FALSE-POSITIVE` with the reason. Publish precision + Wilson interval + the
   denominator. **Sample size and strata fixed in this file before the first case is opened.**
2. **Priority 2 — repair what the audit confirms**, prioritising by consequence, and correct iter-138's
   number in `iter-138/progress.md` and the milestone ledger **in place** if the audit moves it.

**Pre-registered strata (sealed before any case is opened):** 4 from `|Δ| ≥ 100`, 4 from `10 ≤ |Δ| < 100`,
4 from `|Δ| < 10` — **12 cases**. Every case in a stratum is taken in census order from the top, so the
selection is reproducible and not cherry-picked.

## Escalation conditions

- A 3rd unplanned line → tripwire; land the audit, route the repairs.
- **If precision is very low (< 50 %), the correct move is to RETRACT iter-138's number, not to repair
  from it** — and to say so in the ledger.

## Acceptable close-no-lift outcomes

**Precision measuring high (probe upheld) is a first-class result** — it converts iter-138's figure from
*published* to *audited* and lets the fence be built against a trusted baseline. Nothing needs to move
for the iter to have delivered.
