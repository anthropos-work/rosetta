---
milestone: M257x
iter: 09
iteration_type: tik
status: closed-no-lift
opened: 2026-07-31
---

# iter-09 — `FIX-M257x-academy-not-serving`

**Active strategy reference:** `TOK-01` step 5 — *prove it cold*. This is the last genuine `✗` between the
stack and a green cold cycle, and **gate clause 1 needs three of them**.

## Cluster / target identified

After iter-07/08, autoverify on demo-1 reports 3 failures. Two are
`CHECK-M257x-bringup-evidence-logs-absent` (evidence-absence, not defects — and iter-09 resolved that too,
see below). The academy is the only remaining real one.

## Hypothesis at open

The bring-up log says the academy *"is alive but NEVER ANSWERED on :13077 within 120 s"* while its own
output says `✓ Ready in 193ms`. A readiness probe disagreeing with the process it watches is this
milestone's signature shape, so the prior was a **probe** defect (wrong host form, too-short per-attempt
timeout, or a self-matching check).

## Escalation conditions

If the root cause turns out to be a deliberate security tightening rather than a probe bug, do **not**
revert it to make the check pass — characterize it, and design the fix so the security property is
re-proved rather than traded away.

That is exactly what happened. See `progress.md`.
