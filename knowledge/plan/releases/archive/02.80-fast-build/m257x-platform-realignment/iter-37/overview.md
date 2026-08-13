---
milestone: M257x
iter: 37
iteration_type: tik
status: closed-fixed
opened: 2026-08-02
---

# iter-37 — `FIX-M257x-iter32-orgadmin-role-create-timeout`

**Active strategy:** `TOK-01`. The last clause-2 failure, and the one the milestone had never opened.

## Step 0 — re-survey

Confirmed against iter-36's own binding artifact rather than the hand-off: after iter-36 the failing set
is exactly `{pt-orgadmin-role-create}`, one row, `TimeoutError: page.waitForURL: Timeout 60000ms`. Still
meaningful, still the only thing between clause 2 and its gate.

## Hypothesis, and why the obvious one was wrong

The symptom is a 60 s `waitForURL` after Save, and the page object's own header records the app navigating
to `/enterprise/roles/<id>?setup=true` **1.5 s** after Save — MEASURED at M256 iter-22. So the shape
invites iter-36's diagnosis a second time: *the platform changed the post-save navigation*.

**Rejected by measurement, not by argument.** Playwright's failure snapshot shows the roles LIST, no
dialog, 12 roles. And the write side is decisive: `public.job_roles` holds **no** `PT Role%` row and
**zero** rows created in the last two hours. The role was never created, so no navigation was owed.

## Expected lift (pre-registered before any confirming run)

Clause 2 `29 / 1 / 1` → **`30 / 0 / 1`**, i.e. **30 live / 0 failing / 0 error — gate clause 2 MET**.
(The 31st manifest row is the declared in-manifest `will-not-build`, `unimplemented` by design; the gate's
third figure is ERRORS, of which there are none.)

## Escalation conditions

- If the fix needs a credential nobody has → close-no-lift with the falsification and escalate; do NOT
  leave a standing red (`D-v28-3`).
- If it needs a platform edit → STOP.

## Acceptable close-no-lift outcomes

A measured demonstration that the flow cannot succeed on a demo without a credential the project does not
hold is a complete iter under this protocol.
