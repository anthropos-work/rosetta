---
milestone: M256
iter: 11
iteration_type: tik
status: closed-fixed
created: 2026-07-28
---

# M256 · iter-11 — the `blocked` outcome

**Type:** tik · **Active strategy:** `TOK-01` move 4 ("close the honesty items last, deliberately") ·
**Handler:** `BLOCKED-M256-refusal-surface`.

## Step 0 — re-survey (mandatory)

Read the current gate state rather than trusting iter-10's routing verbatim:

- `ptvalidate`: **VALID** — 10 products, 30 use cases, **23 live Playthroughs**, 7 TODO. Unchanged.
- `ptreport`: reads **`e2e/report/last-run.json`**, which iter-10's ad-hoc click-counting probe
  **CLOBBERED** — the file now holds **1 spec** (`how many clicks to switch from hero A to hero B?`),
  so every use case currently reconciles as `failing / no test outcome`. That is a **stale artifact,
  not a regression**: no runtime code changed in iter-10. Recorded as a finding; the gate figure is
  only meaningful after a fresh full run, which Phase C does.
- Clause 2 live counts, from the DB on `demo-2`: **every** `pt-*` membership has its g3 grant
  (`pt-halcyon-retail` 20/20, `pt-kestrel-hiring` 40/40, `pt-meridian-labs` 40/40,
  `pt-vertex-logistics` 40/40). So **`blocked` is still 0 and there is no refusal anywhere in the
  seeded world** — the target is untouched and still meaningful.

Target confirmed: **clause 2's `>= 1 blocked` outcome**, the one sub-clause of the gate that is still
at zero after 10 iters, and the only one whose absence means *nothing in the suite proves the platform
correctly says no*.

## Cluster / target identified

`BLOCKED-M256-refusal-surface` (routed since iter-01, sharpened by iter-01 D4's refutation of the
entitlement mechanism). The routing names the surface and the locator already exists:
`SimulationPage.orgMemberCannotStartModal()`, which `pt-aisim-chat-launch` asserts **ABSENT** today.

## Hypothesis

**H1 (the refusal is real and reachable).** The g3 `FEATURE_JOB_SIMULATIONS` casbin grant is written
**per membership, unconditionally**, by the users seeder (`stack-seeding/seeders/users.go:236`). An org
whose memberships carry **no** g3 grant is the platform's own real state for *"this organization has not
enabled AI Simulations"* (the app grants it via `OrgAllowUserToUseFeature` — an org-admin action). So
withholding the grant for one seeded org produces a **genuine Sentinel/Casbin refusal**, not a
simulated one, and the deny surface becomes the Playthrough's asserted outcome → `outcome: blocked`.

**H2 (the pairing is free).** The same locator, asserted **PRESENT** for a member of the no-feature org
and **ABSENT** for `pt-employee` (who has the grant), is exactly `NEGCTL-M256-cross-vantage`'s
mechanism: one locator, two vantages, opposite verdicts, both live. If H1 holds, the `blocked`
Playthrough **is** `pt-aisim-chat-launch`'s negative control and vice versa.

**The risk H1 must clear first, by probe, before any code is written.** The in-repo evidence about what
the deny surface actually *renders* is **contradictory**: `simulation-page.ts:90-96` names a text modal
(*"cannot start AI Simulations in this organization"*), while `identity.go:250` says the same condition
renders *"the org-member deny modal (empty `<main>`)"*. If the deny surface is an EMPTY main and not that
text, a spec asserting the text would be a **false RED** — the iter-06 `/sim/<slug>/session-list`
trap in the other direction. **Probe the rendered DOM before writing a line of spec.**

## Expected lift

Clause 2's `blocked` count **0 → ≥ 1** (the sub-clause discharged), plus **+1 negative control**
(cross-vantage, on the same locator) and a 24th live Playthrough. No speed claim: a new Playthrough
changes the clause-1 denominator, so clause 1 is **re-verified**, never "improved".

## Phase plan

- **Phase A — probe.** Withhold the g3 grant for one seeded membership on `demo-2`, Sentinel-Reload,
  drive the real browser to `/sim/<slug>` → Start Simulation, and **read what renders**. Restore the
  grant. Decide H1 on measured DOM, not on either comment.
- **Phase B — land** (only if Phase A confirms a real, assertable refusal): the seed opt-out + the
  `seed-worlds.yaml` capability + the manifest use case (`outcome: blocked`) + the spec, plus the
  cross-vantage negative-control assertion.
- **Phase C — re-measure.** Full suite, cold `--reset`, n=3; re-verify clause 1 on the grown
  denominator; refresh `last-run.json` (which also repairs the stale-artifact finding above).
- **Phase D — close.**

## Escalation conditions

- Phase A finds the deny surface is **not assertable** (empty main / no stable landmark) → do **not**
  ship a text assert; record the falsification, route the `blocked` outcome to a different refusal
  surface with the measured evidence attached, close `closed-no-lift`.
- Withholding the grant breaks an **unrelated** Playthrough (a suite red outside the planned scope) →
  user-blocker per Phase 5 §4.
- The seed opt-out cannot be expressed without a platform edit → escalate (there are none in this
  release; the seeder is rext-owned, so this is not expected).

## Acceptable close-no-lift outcomes

A measured refutation of H1 — *"the deny surface renders X, not the modal text; a `blocked` assert on
it would be a false RED"* — with the replacement refusal surface named and priced. That is a real
deliverable: it is the same shape as iter-06's `session-list` finding, which stopped a false-RED
negative control from shipping.
