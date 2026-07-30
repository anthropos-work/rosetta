---
iteration_type: tik
status: archived
opened: 2026-07-30
---

# iter-32 — `individual.UC1`, the org-less user: the last curated use case

**Type:** tik · **Active strategy:** `TOK-01` move 3 — land the coverage clusters; onboarding is the second,
and this is its fifth and final use case.

## Step 0 — re-survey (mandatory)

| fact | reading |
|---|---|
| `ptvalidate --manifest-dir ./manifest --stack demo-2` | manifest VALID · **29 live / 2 TODO** |
| `datadna measure-closure --stack demo-2` | **PASS** — all 279 seeded verified-skill node-ids resolve |
| `demo-2` | 16 containers Up / 0 exited |
| suite | `199 passed` × 3 cold, rc 0, 0 flake (iter-31) |
| the target | `onboarding.individual.UC1` — `verdict: not-yet-built`, handler `ONBOARD-M256-orgless-seat` |

Still the named target, still meaningful, and it is now the **only** `not-yet-built` verdict in the corpus.

## Cluster / target identified

`ONBOARD-M256-orgless-seat`, priced at iter-30 **from measurement**: an org-less user **reaches the app** (she
logs in, `/onboarding` serves the flow, `/home` renders — measured by DELETING a hero's membership row on
demo-2), and the cost is **≥ 5 seeders**, not the "a seeder change" carried in prose since iter-08.

## Hypotheses

- **H1 — the capability is a Persona field with ONE recognition point per consumer, not a special case.**
  `memberships` ids are a deterministic `membershipUUID(prefix, i)` and every membership-keyed seeder loops
  `for i := 1; i <= n`, so "she has no membership" is one predicate consulted at each `membershipUUID` call
  site. **That makes the audit machine-checkable** — the call sites ARE the FK surface — which is a better
  answer than reasoning about 12 constraints.
- **H2 — an org-less hero must be BLANK, and the seed should enforce that rather than permit it.** A solo day-0
  user has no verified skills, no career, no activity. So `verified: 0` should be **required** of her (inverting
  the end-user `verified > 0` rule) rather than merely tolerated — a non-zero count is a contradiction the seed
  can catch. If it is required, the whole persona/profile/activity chain skips her *because she declares
  nothing*, instead of because fifteen seeders each learned a new special case.
- **H3 — the org-scoped tail must be MEASURED, not reasoned.** iter-30's warning is that FK breakage fails
  loudly and names its consumer, while a seeder writing an `organization_id`-bearing row for a user with no
  membership does not break at all and is simply wrong invisibly (the iter-28 **D112** silence class). So after
  seeding, **sweep the live DB for her uuid across every table** and fix what appears — never audit by reading.

## Expected lift

- Clause 3's onboarding half **4 landed → 5 landed**, i.e. the curated corpus's last un-homed use case.
- Its verdict block must be **removed** by the same change — iter-31's bidirectional fence forces it, which is
  the first live exercise of that half.
- Clause 2 must not regress: the new Playthrough arrives **with** a negative control.

## Phase plan

- **Phase A — measure first.** Confirm the seat's population slot is unique and is not slot 1 (`identity.go`
  pins the org-admin identity at `membershipUUID(prefix, 1)`), and re-confirm the four surfaces iter-30 measured.
- **Phase B — the capability**: `Persona.OrgMembership` (closed enum), the `validate()` rules, the per-story
  org-less index set, the skip at every `membershipUUID` site + the persona chain, and a **fence** that fails
  when a new membership-keyed writer appears without the check.
- **Phase C — seed + SWEEP.** Reseed, then sweep the live DB for her uuid across every uuid column in `public`
  and `jobsimulation`; fix every org-scoped row that appears.
- **Phase D — the Playthrough + its negative control**, then the mutants.
- **Phase E — the gate**: `run-playthroughs.sh 2 --reset` × 3 cold, rc per run into a variable, `ptvalidate` +
  `ptreport` + `datadna` + `--policy-check`, gofmt, the Go suites, and the DRIFTED cockpit fixture restored
  **sha-verified** after every reset.

## Escalation conditions

- If the seat's slot collides with another hero's, or lands on slot 1, **move the seat** (iter-28's precedent —
  a seat placement cost a red gate) rather than weakening the collision fence.
- If the DB sweep shows org-scoped rows that cannot be suppressed without a platform edit, the UC gets a
  **written verdict** (now a first-class, machine-checked artifact) rather than a Playthrough over a
  half-org-less user.
- Org B is the host story: `seed-facts.ts` names Halcyon Retail only for `PT_FREE`'s per-hero facts and
  **anchors nothing on its member count** (checked). If any Org B count assertion surfaces, move the seat.

## Acceptable close-no-lift outcomes

A measured refutation of H1/H2 — an org-less hero cannot be expressed without a platform edit — recorded with
the evidence and the UC's verdict re-cut accordingly, is a complete iter.
