---
iteration_type: tik
status: archived
opened: 2026-07-29
---

# iter-13 — sharpen the structural PROFILE finals so they have a contrast vantage again

**Type:** tik · **Active strategy:** `TOK-01` move 4 ("close the honesty items … negative controls")

## Step 0 — re-survey (mandatory)

Re-read before targeting:

- `ptvalidate --manifest-dir=manifest --e2e-dir=e2e` → **VALID**, 24 live Playthroughs / 31 use cases / 7 TODO.
  The tree is clean at iter open; the pipeline this iter wires through accepts what is already there.
- The gate's two live gaps, per `overview.md` (clause 1 re-cut by **D-v28-13** yesterday, so the suite-ratio
  question is no longer a gate and `MEASURE-M256-clause1-sampling` no longer blocks progress):
  - **clause 2** — negative controls **13 of 24**.
  - **clause 3** — LANDED half short: org-admin **2 of 4**, onboarding **1 of 5**.
  - **D-v28-5** — unblocked as of the harden pass (the cockpit-manifest drift fix shipped, rext `0f36f71`).
- `NEGCTL-M256-cross-vantage` is routed to **iter-13+** and names this iter's target class explicitly:
  9 of the 11 uncovered are **structural** finals whose route is *"sharpen the final to name real seeded
  data"*.

**TOK-01's named target is still current** — no substitution. Move 4 is the live one; moves 1–3 are closed
(baseline iter-02, the `networkidle` leg iter-03, org-admin iter-04/05, onboarding iter-08).

## Cluster / target identified

The **three structural PROFILE finals**: `pt-profile-verified`, `pt-profile-growth`, `pt-profile-timeline`.

Why these three and not the workforce five: they share **one page object** (`profile-page.ts`), **one route**
(`/profile` + `/profile/skills`), and **one already-wired contrast vantage** (`pt-manager`, live in
`negative-controls.spec.ts` since iter-12 for `pt-profile-identity`). So the marginal cost of the batch is the
sharpening itself, not new login plumbing — and if the mechanism works here it is a template the workforce
quartet can consume in a later tik. It is also the batch iter-12 measured and named: it reads
`verifiedSkillsStat` 1 / `skillCharts` 10 / `workSection` 1 for `pt-manager`, i.e. the exact numbers that
prove the CURRENT finals do not discriminate.

## Hypothesis

**iter-12's refutation closed a shortcut, not the path.** Its measured finding is that a *structural*
predicate — a stat LABEL is visible, a chart count ≥ 1, a "Work" section exists — renders for **any**
populated member, because M44's profile-completeness seeder gives every member a career and skills. That is
true, and it is why no *suppression* mechanism can exist (there is no vantage for whom the outcome is
legitimately absent).

But the finals are structural **because they were written structurally**, not because the surface carries
nothing hero-specific. `pt-world.seed.yaml` pins each hero's own data deterministically and by name:
`pt-employee` = Pat Ellis, **DevOps Engineer**, `skills: {verified: 8, mapped: 12}`; `pt-manager` = Morgan
Reyes, **Engineering Manager**; `pt-free` = Sam Okafor, **Account Executive**, `{verified: 2, mapped: 8}`.

So: **re-aim each final at the hero's OWN seeded data** and the final becomes hero-specific — at which point
the contrast vantage that does not exist for a structural predicate **does** exist for a specific one
(`pt-manager`'s profile shows *her* role and *her* magnitudes, not Pat's).

This is a genuine strengthening independent of the control: today these three Playthroughs would pass
against **any** member's profile, which is the M219 lesson unlearned — *a surface that renders is not the
same as the RIGHT surface*.

## Expected lift

- Negative controls **13 → 16 of 24** (clause 2's live gap), computed by the fence, not narrated.
- Three Playthroughs that today prove "a profile rendered" prove "**this hero's** profile rendered".
- No new Playthrough, no seed change, no new page object file → the gated median is untouched by
  construction (controls live outside the Playthroughs, iter-12 D53).

## Phase plan

- **Phase A — probe first (mandatory; this milestone has refuted its own plan in 6 of 12 iters).** Drive a
  live probe on `demo-2` that dumps, for `pt-employee` **and** `pt-manager` **and** `pt-free`: the
  `/profile/skills` stat magnitudes as rendered, and the `/profile` career-tab role/timeline text. Only then
  decide which seeded fact each sharpened final names. **Do not write an assertion whose discriminating power
  has not been measured on both vantages.**
- **Phase B — sharpen** the three finals (spec + page object accessors as needed, semantic locators only, P3).
- **Phase C — land the controls** in `negative-controls.spec.ts` with the `@pt-control-for:` link, and
  **watch each one go RED**: point the sharpened final at the contrast hero and show it fails.
- **Phase D — re-measure:** full suite ×3 on cold reset-to-seed, the fence's computed control count, `ptreport`
  four-state, `ptvalidate`, the unit fences, the Go modules.
- **Phase E — close:** commit both repos, doc backfill if the pattern generalizes (protocol-evolution rule).

## Escalation conditions

- If the probe shows the profile surfaces render **no hero-specific fact at all** on either the skills or
  career tab, the hypothesis is refuted: record the falsification (`closed-no-lift`), and route the 9
  structural controls to a **declared verdict** rather than a mechanism — that would be a real change to
  clause 2's reachability and it goes to the user, not into a workaround.
- If a sharpened final proves **flaky** (a magnitude that drifts between resets), revert that one final and
  keep the stable ones — a false RED is exactly as dishonest as a false green (iter-06's rule).

## Acceptable close-no-lift outcomes

A measured refutation of the "sharpen for discrimination" mechanism, with the per-vantage numbers attached —
i.e. the same shape iter-05, iter-07 and iter-12 produced. That would be worth more than three vacuous
controls, and it is what the escalation above routes.
