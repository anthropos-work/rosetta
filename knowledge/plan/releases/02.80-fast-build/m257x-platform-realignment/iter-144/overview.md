---
iter: 144
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
active_strategy: TOK-08
route_closed: FIX-M257x-h30-crossline-repair
---

# iter-144 — the wrapped retraction sites, repaired; and a sub-class the fence cannot see

**Active strategy reference:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them*. This iter closes `FIX-M257x-h30-crossline-repair`, routed by harden pass 30 as *"the
8 true sites across 6 files — editorial prose judgement in six documents is an iter's work, not a
harden inline fix (Fate 3)."*

**Step 0 — re-survey (mandatory).** Re-ran `retracted_pin_guard --repo-root .` at iter open. The
wrapped (SURVEY) arm reports **10 findings across 6 files**, not 8. The route's figure is not
re-quoted; the measured one is used, and the discrepancy is resolved by the audit below rather than
assumed away. Gating arm (bare, single-line) **GREEN**, as iter-142 left it.

## Cluster / target identified

Harden pass 30 landed `§5` rule 64 (*a fence over wrapped prose must state its line reach*) and gave
`retracted_pin_guard` a wrapped arm. That arm is **SURVEY, not gating** — deliberately, after the
pass found the path arm's constant had been silently gating. So its findings sit unrepaired by
construction, and the repair is this iter's job.

## Hypothesis

The 10 wrapped findings are the same class iter-142 repaired on the single-line arm, and the same
remedy applies: **retract by describing the artifact, never by reproducing it** (rule 63(c′)), with
every file line-count FLAT so the repair cannot induce the rot it removes (`D-M257x-142-4`).

## Expected lift

The wrapped arm's live population goes to zero, or every survivor is explained.

## Phase plan (three planned lines — declared, per the scope-creep carve-out)

1. **Audit all 10 in full context BEFORE repairing** (`D-M257x-142-1`).
2. **Control the audit with a machine, not a second reading** (`D-M257x-143-1`, this milestone's
   newest rule) — resolve every pin the audit calls LIVE against the source it claims, and let the
   resolver, not the reader, settle it.
3. **Repair what the controlled audit supports**, line-count flat; explain every survivor.

## Escalation conditions

- A survivor that is neither repairable nor explainable → route forward, do not repair blind.
- The audit and the machine control disagree → the machine wins, and the disagreement is the finding
  (this is exactly how iter-143 caught its own reader).

## Acceptable close-no-lift outcomes

If the audit shows the wrapped findings are predominantly false, the deliverable is the **measured
precision of the wrapped arm** plus the mechanism — which is what harden pass 30 could not supply,
having routed the sites without grading them.
