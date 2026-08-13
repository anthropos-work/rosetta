---
iter: 23
iteration_type: tik
iter_shape: standard
status: closed-fixed
opened: 2026-07-30
---

# iter-23 — the refusal nobody sees, captured on purpose

**Active strategy reference:** `TOK-01` move 4 (*"close the honesty items last, deliberately, not as
leftovers"*). This is an honesty item in the strictest sense: a **product** defect this milestone
discovered, which no gate clause rewards recording, and which is therefore exactly the kind of thing a
milestone drops.

## Step 0 — re-survey (mandatory)

| checked | reading |
|---|---|
| `DEFECT-M256-silent-forbidden-mutation` in the routing table | present, targeted **iter-23**, still owed |
| the mechanism, as recorded at iter-20 D98 | *"the refusal is surfaced to the user as NOTHING AT ALL — no alert, no toast, no inline message, the dialog simply stays open"* |
| is the path still exercised on demo-2? | **NO** — iter-21 seeded the grant, so the write now succeeds. It must be **reproduced deliberately** |
| `stackseed --policy-check --stack demo-2` | `live=18 expected=18 · OK` (verified after iter-22's three resets) |
| demo-2 | 16 containers Up, 0 exited |

**Not absorbed, and not stale.** iter-22 deliberately declined to open it because the reproduction
disables the write path iter-22 was landing.

## Cluster / target identified

`DEFECT-M256-silent-forbidden-mutation`. The defect is the **reason the create-role UC was misdiagnosed
for fifteen iters**: a mutation that the server refuses looks, from the UI, exactly like a mutation that
was never sent. iter-20 found it while root-causing something else and recorded it in one paragraph.
That paragraph is the entire asset, and it is attached to a UC that is now green — so it is one
manifest edit away from disappearing.

## Hypothesis

1. The silence is **systemic, not per-form**: a GraphQL error riding inside an HTTP 200 has no
   user-visible surface *anywhere* in the org-admin write set. If so this is **one** platform defect
   with four instances, not four defects — a materially different report.
2. The three org-admin writes that DO work today (tags-create, member-tag, settings-toggle) would show
   the same silence if their own mutation were refused. **Measure it; do not infer it from (1).**

## Expected lift

**No gate clause moves.** The deliverable is a defect record good enough for someone else to act on
without re-deriving it, plus the sweep that establishes its true scope. Recording it is the point;
iter-20 escalated rather than forcing green, and this is the other half of that decision.

## Phase plan

- **A — reproduce it deliberately.** Revoke the `p3 admin → org:feature:taxonomy:write` row on
  **demo-2 only** (a demo-DB write, permitted; never prod), reload Sentinel, drive the real create-role
  journey in a browser, and capture the *whole* observable surface: the response body, the dialog state,
  every `role=alert` / toast / inline message, and the catalog count. Then **restore the row** and
  re-verify with `--policy-check`.
- **B — read the error path** in `stack-demo/next-web-app` (**read-only**; zero platform edits) to
  establish whether the silence is systemic or local to one form.
- **C — sweep the other three org-admin writes** for the same shape, by measurement.
- **D — write the record and route it**, then close.

## Escalation conditions

- The revoke does not reproduce the refusal (e.g. a cached enforcer decision) → the defect's stated
  mechanism is wrong and that is the finding; record it and stop rather than manufacture the symptom.
- Restoring the grant fails to re-enable the write → **stop and escalate**: the stack is left in a state
  a later iter depends on, and that outranks this iter's deliverable.

## Acceptable close-no-lift outcomes

Finding that the refusal **does** surface somewhere I had not looked (a console error, a dev-tools-only
warning, an off-screen `alert` region) would retract iter-20 D98. That is a first-class outcome: D98 is
currently an unverified single-observation claim, and this iter's job is to make it evidence either way.

## Safety

Every write is to **demo-2's own Postgres**. The production DB is read-only and is not touched at all in
this iter. No platform repo is edited — the next-web read is a read. `billion` is not touched.
