---
iter: 05
milestone: M211
iteration_type: tik
status: closed-fixed
created: 2026-07-08
---

# iter-05 — tik: verify net GREEN with the merged-platform assertion

**Type:** tik — under **TOK-01**.

## Step 0 — Re-survey
iter-04 loaded the casbin policy (casbin_rules=170) + re-seeded clean. The verify net (sub-condition (d))
is the open next thing; the casbin cheap-win is now satisfiable.

## Active strategy reference
**TOK-01** — the warm inner-loop's verify gate.

## Cluster / target
Sub-condition **(d)**: the `verification.md` auto-verify net passes with a **merged-platform assertion**
(no skiller schema/subgraph/container; `readiness.sh` schema probe GREEN) on the warm merged stack.

## Hypothesis
With taxonomy + seed + casbin loaded, the re-grounded stack-verify (M209 dropped skiller from the
expected-schemas + service list) runs GREEN on the merged backend stack, and its schema probe asserts the
merged shape (no skiller).

## Phase plan (executed)
1. `autoverify.sh --project anthropos` — cheap-wins + scoped verify.
2. Full `verify.sh` for per-service detail.
3. Inspect the expected-schemas list (merged-platform assertion content).
4. Scoped `verify.sh` (backend graphql-profile services) — confirm GREEN with UI/directus correctly skipped.

## Escalation conditions
- If verify still expected a skiller schema/subgraph → an M209 miss → route a rext fix. **Did NOT fire** —
  expected-schemas is `(public sentinel cms jobsimulation skillpath extensions)`, no skiller.

## Outcome
**Target MET.** Cheap-wins: backend `/api/health` 200 + `sentinel.casbin_rules = 170` (both GREEN — the
M18 silent-403 catcher passes, confirming iter-04). Scoped backend verify → **"✓ all live probes passed"**:
all 11 backend services live+ready; `readiness postgres-schemas: all expected schemas present` (the
merged-platform assertion — expected list has NO skiller, passes with skiller schema absent);
`graphql-introspection ok` (4-subgraph supergraph, no skiller subgraph); `sentinel-rpc ok`;
`gotenberg-version 8.34.0`. The 4 unscoped failures (next-web-app, studio-desk, directus,
directus-collections) are UI-tier/local-Directus services NOT started on a backend-only prod-read warm
stack — correctly skipped when scoped (verification.md scope model), deferred to the full UI/cold bring-up.
See progress.md.
