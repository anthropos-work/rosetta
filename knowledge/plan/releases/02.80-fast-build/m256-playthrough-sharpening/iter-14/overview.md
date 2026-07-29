---
iteration_type: tik
status: in-progress
opened: 2026-07-29
---

# iter-14 — sharpen the four WORKFORCE finals the same way iter-13 sharpened the profile three

**Type:** tik · **Active strategy:** `TOK-01` move 4 ("close the honesty items … negative controls")

## Step 0 — re-survey (mandatory)

- `ptvalidate --manifest-dir=manifest --e2e-dir=e2e` → **VALID**, 24 live Playthroughs / 31 use cases /
  7 TODO. Tree clean at iter open, both repos.
- `demo-2` up (16 containers, 26 h), seeded `pt-world`, suite green at iter-13's close (166 passed / 0 flake).
- The gate's live gaps per `overview.md`:
  - **clause 2** — negative controls **16 of 24**; mutating 6/5 MET; `blocked` 1/1 MET.
  - **clause 3** — LANDED half short: org-admin 2 of 4, onboarding 1 of 5.
  - **clause 1** — flake half MET; leg half N/A unless a speed mechanism lands.
  - **D-v28-5** — unblocked (harden pass shipped the cockpit/roster authority fix), not started.
- `NEGCTL-M256-cross-vantage` is routed to **iter-14+** and names the residual explicitly: six
  Playthroughs reachable by iter-13's move — `pt-workforce-{roster,funnel,succession,org-feedback}`,
  `pt-activity-drilldown`, `pt-hiring-recruiter-compare` — with the hiring one flagged as needing a
  same-vantage control (its contrast vantage ejects the browser to production).

**TOK-01's named target is still current** — no substitution.

## Cluster / target identified

The **four WORKFORCE finals** on the `pt-manager` seat: `pt-workforce-roster`, `pt-workforce-funnel`,
`pt-workforce-succession`, `pt-workforce-org-feedback`.

Why these four as one batch, and why the other two are NOT in it:

- All four are **org-aggregate surfaces read by the same hero** (`pt-manager`, Org A / Meridian Labs), so
  one contrast login covers all four absences — the same batching economy iter-12 established and iter-13
  reused. The marginal cost per Playthrough is a page-object accessor and an assertion.
- `pt-activity-drilldown` is **routed to iter-15 with its measurement attached**: Phase A found its grid
  renders **20 rows with no cell content** during hydration, so its `contentRows().count() > 0` assertion
  passes over a skeleton — that is a *defect fix* plus a two-step drill, not a sharpening, and mixing it
  into a four-way batch is the scope-creep the tripwire names. Its discriminator is measured and recorded
  so iter-15 is cheap.
- `pt-hiring-recruiter-compare` needs a **different mechanism** (same-vantage), also iter-15.

## Hypothesis

Identical in shape to iter-13, one product over: these four finals are structural (*"the Role column is
present"*, *"a chart drew"*, *"a stat label is visible"*, *"the table has rows"*) and therefore pass for
**any** populated org — so they would be green if the manager were shown **another tenant's** aggregates.
The seed pins each org's identity and each hero's own facts, so re-aiming the finals at those makes all four
discriminating, at which point a manager in a **different seeded org** falsifies every one of them.

**Phase A measured this before a line was written** (both seats, all five surfaces, `demo-2`, settled):

| surface | Org A / `pt-manager` (main) | Org C / `pt-ai-manager` (main) |
|---|---|---|
| `/enterprise/members` | `Pat Ellis` 1 · `DevOps Engineer` 1 · `pt-meridian-labs` **15** | 0 · 0 · 0 |
| `/enterprise/workforce` | `Pat Ellis` 1 · `DevOps Engineer` 2 · `Overall Members` = **40** | 0 · 0 · `Overall Members` = 40 |
| `/enterprise/workforce/succession` | `Pat Ellis` 1 · `DevOps Engineer` **3** | 0 · 0 |
| `/enterprise/organization-feedback` | `Pat Ellis` 1 · `pt-meridian-labs` **15** | 0 · 0 |

And two negative measurements that shaped the design:

- **The org NAME is chrome, not outcome** (`Meridian Labs` main=**0**, body=1 — it is the sidebar org
  switcher). Asserting it would prove the switcher label, not that the aggregate is this org's. Rejected.
- **`Overall Members` = 40 for BOTH orgs** (both seeded `size: 40`), so it is a seed-linked *magnitude*
  strengthening — it catches a stat card rendering a constant or an accessor reading the wrong card — but
  it is **not** a discriminator, and it is not used as one.

The discriminator that generalises across three of the four surfaces is the **org email domain**
(`…@pt-meridian-labs.com`, 15 of 20 rendered rows on both table surfaces): derived from the seed's
`org.slug`, present on every page of the roster, so unlike a single hero row it is **pagination-proof**.

## Expected lift

- Negative controls **16 → 20 of 24**, computed by the fence.
- Four Playthroughs that today prove "an org dashboard rendered" prove "**this org's** dashboard rendered"
  — i.e. they become multi-tenancy proofs, which is a strengthening independent of the control.
- No new Playthrough, no seed change, no new spec file in the gated median (controls live in
  `negative-controls.spec.ts`, outside it — iter-12 D53), so the median is untouched by construction.

## Phase plan

- **Phase A — probe first. DONE at iter open** (three passes; the table above). Pass 2's reading that Org C's
  activity grid was permanently empty was **refuted by pass 3** — it is a transient un-hydrated state, and
  the probe's own settle condition (`rows > 5`) matched the skeleton. Recorded, because the *assertion* it
  exposed is real even though the *absence* is not.
- **Phase B — sharpen** the four finals; keep the structural ones as intermediates (iter-13's pattern).
  Extend `lib/seed-facts.ts` with the org facts + the contrast manager, and the fail-closed fence with them.
- **Phase C — land the controls** in `negative-controls.spec.ts` on ONE `pt-ai-manager` login, each absence
  paired with the contrast manager's OWN equivalent asserted PRESENT (iter-13 D64), and **watch every one
  go RED**.
- **Phase D — re-measure:** full suite ×3 cold reset-to-seed, the fence's computed control count, `ptreport`
  four-state, `ptvalidate`, the unit fences, all six Go modules, the four Python suites.
- **Phase E — close:** commit both repos, tag + push, doc backfill if the pattern generalises.

## Escalation conditions

- If a sharpened final proves **flaky** across the 3× gate (a magnitude or a row that drifts between
  resets), revert that one final and keep the stable ones — a false RED is exactly as dishonest as a false
  green (iter-06's rule). The known risk is the roster's **pagination**: 20 of 40 rows render, and the
  hero's row is on page 1 only because the grid sorts alphabetically by surname and hers is "Ellis"
  (measured: Andersen, Becker, Costa, Dubois, **Ellis**, Esposito, …). The org-domain assertion is the
  pagination-proof half and does not depend on it.
- If the contrast manager turns out **not** to be able to reach a surface (an authz difference between the
  two manager seats), that absence is not evidence — record it and drop that control rather than bank it.

## Acceptable close-no-lift outcomes

A measured refutation — e.g. the org-domain anchor turning out to be generated rather than seed-derived, or
the contrast seat reading Org A's data (which would be a real tenancy defect and a far more valuable
finding than four controls).
