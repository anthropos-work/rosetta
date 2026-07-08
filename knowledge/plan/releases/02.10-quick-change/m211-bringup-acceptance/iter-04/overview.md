---
iter: 04
milestone: M211
iteration_type: tik
iter_shape: bringup-fix
status: closed-fixed
created: 2026-07-08
---

# iter-04 — tik: casbin policy load — fix the dev bring-up's missing sentinel-policy step

**Type:** tik (bring-up fix) — under **TOK-01**. Handler: `bringup-M211-iter04-casbin-migrate`
(the route iter-03 carried forward).

## Step 0 — Re-survey
iter-03 surfaced: seed's identity+users casbin-grant steps fail because `sentinel.casbin_rules` (and the
`sentinel` schema) is absent. Confirmed at iter-04 open: schema absent → the target is live + open.

## Active strategy reference
**TOK-01** — the warm inner-loop's authz-policy gate (sub-condition (d) start: `casbin_rules > 0`).

## Cluster / target
Get the Sentinel casbin policy loaded on the merged platform so (i) the seed's casbin-grants succeed and
(ii) sub-condition (d)'s cheap-win assert (`sentinel.casbin_rules > 0`) is satisfiable — and ROUTE the fix
into the documented cold `/dev-up` path so a cold bring-up reproduces it.

## Hypothesis
The dev cold bring-up creates the `sentinel` schema but never loads the casbin p-model policy
(`init_policy.sql`) — unlike the demo path (migrate-demo.sh) and `bootstrap-dev`. Sentinel auto-creates an
EMPTY `casbin_rules` on startup but does not seed it → blanket 403 + the seed's grant step has no p-model.
Loading `init_policy.sql` fixes it; the doc/skill gap must be closed for the cold gate.

## Phase plan (bring-up fix, multi-step)
1. Triage: create sentinel schema + restart → observe whether sentinel self-seeds casbin_rules.
2. Empirical finding: casbin_rules auto-created but EMPTY → the policy load is the missing step.
3. Route the fix to the corpus (setup_guide.md + dev-up SKILL) — add the `init_policy.sql` load.
4. Apply operationally on the warm stack; assert `casbin_rules > 0`.
5. Re-seed (--reset --force + stories-maya) → identity+users casbin-grants succeed; re-measure closure.

## Escalation conditions
- If the fix required a platform-repo edit → ESCALATE (`unimplementable-without-platform-edit`). **Did NOT
  fire** — the fix is a corpus doc/skill step + a chartered operational psql load (no platform edit).

## Outcome
**Target MET.** casbin_rules 0 → **170** (68 p-model p/p2/p3/p5 via init_policy.sql + 51 g2 + 51 g3 seeded
grants); backend `/api/health` OK; re-seed clean (identity ok=4, users ok=250, **0 seeder failures**, was
2); closure still PASS. Routed fix landed in `setup_guide.md` + `dev-up/SKILL.md`. See progress.md.
