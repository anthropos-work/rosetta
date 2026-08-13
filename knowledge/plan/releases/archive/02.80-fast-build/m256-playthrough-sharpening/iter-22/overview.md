---
iter: 22
iteration_type: tik
iter_shape: standard
status: closed-fixed
opened: 2026-07-30
---

# iter-22 — the fourth org-admin write, landed with its own control

**Active strategy reference:** `TOK-01` move 3 (*"land org-admin before onboarding, because org-admin
discharges two clauses with one body of work"*). This is that move's **last** unit of work: the fourth and
final org-admin use case, whose three config-fidelity blockers iter-21 cleared.

## Step 0 — re-survey (mandatory)

Re-measured before targeting; the routed target is **still live and still meaningful**:

| checked | reading |
|---|---|
| `manifest/org-admin.yaml` `org-admin.roles.UC1` | **`playthrough: TODO`** — not absorbed by any iter |
| the drafted spec | `e2e/drafts/orgadmin-role-create.spec.ts.draft` present, 70 lines, unlanded |
| `OrgRolesPage` page object | **already exists** (`org-admin-page.ts:78-133`) — heading, count, dialog, create, read-back |
| `stackseed --policy-check --stack demo-2` | **`live=18 expected=18 · OK`** — iter-21's grant survives on the running stack |
| demo-2 | 16 containers Up, **0 exited** (the iter-15 D77 liveness check, run before any diagnosis) |

**No substitution needed.** TOK-01's named target for this slot is exactly what is in front of me.

## Cluster / target identified

`PT-M256-orgadmin-role-create` — the **Playthrough half**. iter-21 cleared its three blockers (the
platform's sanctioned `p3` grant, the `SKILLER_AZURE_OPENAI_*` genes, the un-restored identity sequences)
and **proved the write path end-to-end by hand**, but the crash landed between the fix and the test, so the
Playthrough itself was never landed.

**Landed WITH its negative control, deliberately.** The coupling the milestone learned at iter-17 is that
adding a live Playthrough moves clause 2's *denominator* — controls went 23/24 → 23/25 when
`pt-orgadmin-member-tag` landed. Landing the control in the same iter keeps clause 3 advancing without
clause 2 regressing.

## Hypothesis

1. With the three iter-21 fixes in place, the drafted spec's journey **completes** and its read-back
   (`catalog grew by exactly one` + `the title is listed`) holds. **Probe the real journey FIRST** — iter-21
   recorded that the app *navigated to the new role's detail page* after Save, which the draft's
   "dialog hides, then re-read the list underneath" shape does not anticipate.
2. A negative control exists **without a contrast tenant**: the created title is per-run unique, so its
   absence is provable on the same vantage before the write. Per iter-06 D22 a mutating Playthrough's
   **pre-state read IS its control** when the final is a strict delta — but that only counts if the
   pre-state read is asserted, not merely taken as a baseline.

## Expected lift

- clause 3: org-admin **3/4 → 4/4** (the product complete; 8 of 9 M201 clusters' landings done)
- clause 2: mutating **7 → 8**; controls **23/25 → 24/26** (numerator and denominator both move)
- clause 1: no speed mechanism landed, so the leg half has nothing new to measure; the flake half is
  re-verified on the grown denominator (3× cold reset-to-seed)

## Phase plan

- **A — probe the live journey before writing.** Drive the real create-role flow on demo-2 in a real
  browser and measure: what the app does after `Save` (stays / navigates), whether the roles table
  paginates (a `+1` read-back is a bet on sort order if it does), and where the pre-state absence is
  readable. Refute or confirm the draft's shape *before* landing it.
- **B — land the spec + the control**, each assertion watched RED (mutants), then flip the manifest
  `TODO → declared` with the diagnosis comment reduced to the landed truth.
- **C — gate.** `run-playthroughs.sh 2 --reset` ×3 cold, rc captured per run, fixture backed up + sha-verified.
- **D — close**, with the clause standing restated.

## Escalation conditions

- The journey completes but its outcome is **not readable on any surface** → the final would be a toast or a
  DB assert; declare the weaker proof shape explicitly rather than dressing it up.
- The read-back is only satisfiable by a **paginated first page** → sharpen to a name lookup (the
  `OrgTagsPage` precedent), never a row-count bet on sort order.
- Any assertion cannot be watched RED → it is not an assertion; do not claim it.

## Acceptable close-no-lift outcomes

The journey failing *after* iter-21's three fixes would mean a **fourth** blocker in the series, which is
worth more than a spec written on the assumption there were three. Record it measured, route it, close no-lift.

## Deliberately NOT in this iter

`DEFECT-M256-silent-forbidden-mutation` (the capture) — it needs the grant **revoked** on demo-2 to
reproduce, which would break the very write path this iter lands. Sequenced to iter-23. **Its evidence is
not perishable in the way it was feared**: we hold the grant, so the refusal can be reproduced on demand
(revoke → observe → restore). What changed is that the path is no longer exercised *by accident*, which is
a reason to schedule it, not to rush it into an iter whose planned scope it would undermine.
