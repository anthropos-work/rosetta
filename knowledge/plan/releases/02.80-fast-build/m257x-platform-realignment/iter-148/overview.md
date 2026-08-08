---
iter: 148
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-08
---

# iter-148 — the absent value, one section over: an unscoped probe set

**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).
Tik. `TOK-08`'s sealed refutation branch bars an agent-authored successor.

## Cluster / target identified

iter-147's own new route, **`SURVEY-M257x-iter147-absent-value-class`**: the same *absent-value* question
at the tooling's other choice-points. iter-147 fixed a compose **profile** that defaulted to none; the
sibling scope variable is **`STACK_SERVICES`**, and `stack-verify/lib/services.sh:33` **discloses the
hazard in its own comment** — *"running verify.sh with NO STACK_SERVICES set probes everything in the
table and will false-`down` the merged-away rows."*

## Hypothesis

A verify entry point exists that supplies no scope, and it reports the platform's merged-away services
as failures. If so it is gate-adjacent: clause 1 is *"`autoverify green:true / 0 warnings`"*.

## Expected lift

No `N` reading planned → **no `N` movement claimed**. Deliverable: the verify entry points censused with
their denominator, any unscoped path repaired or made to disclose, and a fence that cannot rot.

## Phase plan

A census · B measure both arms against a live stack (read-only) · C repair · D fence + gates.

## Escalation conditions

A repair requiring a platform edit, or requiring `stack-demo/**` to be modified, escalates. **Probing**
`demo-1` read-only is not modifying it.

## Acceptable close-no-lift outcomes

Every verify path already scoping is a complete iter, provided the census states its denominator.
