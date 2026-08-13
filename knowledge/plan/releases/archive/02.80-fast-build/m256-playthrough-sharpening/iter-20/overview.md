---
iter: 20
milestone: M256
iteration_type: tik
status: closed
opened: 2026-07-29
---

# iter-20 — the last org-admin use case, and it was never a form problem

**Active strategy:** `TOK-01` move 3 — org-admin, the cluster that discharges two clauses with one body of
work. **3 of 4 landed**; `org-admin.roles.UC1` is the last.

## Cluster / target identified

`PT-M256-orgadmin-role-create`, parked since **iter-05** — fifteen iters. Its recorded blocker (iter-05 D18,
re-routed by iter-17):

> "Suggest skills" transforms the create-role dialog and the primary button becomes **`Generate`** (a
> live-LLM leg → budget it with the studio lane). Separately: **report the product defect** — `Save` is
> enabled with the form incomplete and fails silently with an EMPTY `alert` region.

iter-17 added: *try the KEYBOARD route (D83) first, and do not rely on iter-05 D19's `force` claim
(retracted).* The complete draft has been sitting in `e2e/drafts/orgadmin-role-create.spec.ts.draft`.

## Hypothesis

The same pattern as iter-17 and iter-18: **a blocker recorded once and never re-driven is a hypothesis.**
Drive the dialog, watch the network, and find out what `Save` actually does before choosing a route.

## Expected lift

Clause 3 landed-half **org-admin 3 → 4 of 4** (complete), and clause 2 mutating **7 → 8**.

## Phase plan

- **A — drive the dialog and capture the NETWORK** (the half iter-05 never looked at).
- **B/C/D** — implement, mutation-verify, 3× cold gate, *conditional on what A finds.*

## Escalation conditions

- If the write is refused by **authorization** rather than by the form, then granting ourselves the missing
  permission would **manufacture the capability under test** — the `force: true` failure mode (iter-07,
  sharpened iter-17) one layer down. That is a **user decision**, not an implementation choice.
- The milestone holds **0** `unimplementable-without-platform-edit`; the re-scope trigger fires above 3. One
  is not the trigger, but it is a first, and the protocol says **escalate, never absorb**.

## Acceptable close-no-lift outcomes

A measured root cause that replaces the fifteen-iter-old diagnosis, with the disposition question stated
precisely enough for the user to answer in one reading.
