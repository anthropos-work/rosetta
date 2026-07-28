---
iter: 05
milestone: M256
iteration_type: tik
status: closed-no-lift
opened: 2026-07-28
---

# M256 · iter-05 — unblock the two parked org-admin journeys (clause 2's mutating floor)

**Type:** tik · **Active strategy:** `TOK-01` move 3 (finish the org-admin cluster) feeding move 4.

## Step 0 — re-survey

iter-04 routed two gate-critical items forward with written diagnoses. Re-surveyed: `ptvalidate` reports
`9 product(s), 22 use case(s), 20 live, 2 TODO`; the two drafts are in `e2e/drafts/`; clause 2's mutating
count is **3 of the required 5** (**D17**). Nothing has changed since. Targets stand, and they are the same
two — this is the routed work, not a substitution.

## Phase 0d — pre-flight tooling check

The iter re-points manifest `playthrough:` keys and moves specs back into `tests/`, i.e. it wires artifacts
through the `ptvalidate` both-way-integrity pipeline. Re-run on the current tree before authoring; it passed
at iter-04 and nothing has changed since, so the check is a confirmation rather than a discovery.

## Cluster / target identified

| Handler | Use case | The diagnosis iter-04 recorded |
|---|---|---|
| `PT-M256-orgadmin-role-create` | `org-admin.roles.UC1` | `Save` enabled but apparently a **no-op**: no HTTP ≥ 400, no console error, no navigation, no new row, dialog open, `alert` region **empty**. Untried: **"Suggest skills"** |
| `PT-M256-orgadmin-member-tag` | `org-admin.members.UC1` | dropdown **pointer interception** fixed with `Escape`; `checkbox.check()` then still times out. The modal has a **filter box** and its own scroll container |

Both are **gate-critical**: clause 2 requires **≥ 5** mutating Playthroughs and stands at 3.

## Hypothesis

- **Roles:** the "Core skills" step is mandatory and `Save` fails validation **silently**. If picking a
  core-skills path satisfies it, the journey lands. If `Save` still no-ops with skills present, the defect is
  **platform-side** and the honest output is a *reported finding* plus a different fifth mutating journey.
- **Members:** the tag list is long and virtualised/scrolled, so the target checkbox is not reachable by
  `check()` alone. Filtering the list to the tag's unique name should make it directly actionable.

## Expected lift

Clause 2's mutating count **3 → 5**, and clause 3's org-admin half **2/4 → 4/4**. If the roles defect is
platform-side, the lift is **1** (members) plus a named platform finding, and the fifth mutating journey is
identified for the next iter.

## Phase plan

- **Phase A — probe both blockers** (the diagnoses are already written; this tests the two named hypotheses).
- **Phase B — land what the probes support**: move the draft back into `tests/`, re-point the use case's
  `playthrough:` **in the same change** (the manifest/test lockstep rule), extend the page object.
- **Phase C — run + reconcile**; then **re-measure** under D7's protocol if the denominator changed.

## Escalation conditions

- Roles `Save` no-ops **with** core skills present → that is a **platform defect**, not a test bug. Record it,
  keep the UC `TODO` with the sharpened finding, and name the replacement fifth mutating journey. Do **not**
  edit the platform (P3), and do **not** weaken the assertion to make it pass.
- Either journey turning red at batch end → per **D-v28-3** the batch runs to completion and one consolidated
  red set escalates; a journey that cannot land stays a declared `TODO` rather than standing red.

## Acceptable close-no-lift outcomes

A sharpened, evidenced verdict on the roles `Save` — even one that keeps the UC `TODO` — is a complete
outcome: it converts "we could not make it pass" into "the platform does not do this, and here is the
evidence". That is a finding the milestone exists to produce.
