---
iter: 05
milestone: M21
iteration_type: tik
status: closed-fixed
created: 2026-06-11
---

# M21 iter-05 — tik (fourth tik under TOK-01): stages 5–6, serve a sim anonymously (the gate)

Under TOK-01. Target: advance furthest-passing-stage **4 → 6** (boot Directus on the stage-4 harness + serve a
captured simulation anonymously over HTTP — the exit gate). This is the milestone's flagged **live-only risk**
(Directus anonymous-permission model).

## Active strategy reference
TOK-01 + M21-D7 (option A). Builds directly on the iter-04 harness (26 tables + 9 row-tables, digest `6cd35278…`).

## Re-survey (Phase 1 Step 0)
furthest-passing-stage = 4 (iter-04: digest converged, replay exit 0, 10128 rows). Route `STRUCT-M21-iter05-serve`
untouched. Harness retained in the stage-4 state.

## Cluster / target identified
For Directus to serve `/items/simulations` anonymously it needs: (a) the `directus_collections` rows (a Directus only
serves REGISTERED collections); (b) the public-access chain — prod uses a `$t:public_label` policy
(`abf8a154-5b1c-4a46-ac9c-7300570f4f17`) linked via `directus_access` with `role=NULL, user=NULL`, granting `read`
on the content collections with `fields='*'`. The table rows already exist (stage 4).

## Hypothesis
Loading the `directus_collections` rows + replicating prod's public policy/access/read-permission rows into the
harness, then booting Directus on the stage-4 schema, makes `GET /items/simulations?limit=1` return **200 with a real
row to an anonymous reader** — the gate. (The `directus_fields` rows are likely NOT required for a raw GET — Directus
serves a registered collection's DB columns without explicit field metadata; load them only if the GET needs them.)

## Expected lift
furthest-passing-stage **4 → 6** (boot + serve anonymously). Partial-credit: **4 → 5** if Directus boots + is
reachable but anonymous read needs more permission wiring than expected (characterize + route).

## Phase plan
1. Load `directus_collections` (26) + the public policy + its `directus_access` link + the content read-permissions
   from prod (json_populate_recordset; harness system tables are byte-identical to prod's).
2. Boot Directus on the harness, published on an offset port.
3. `GET /items/simulations?limit=1` with NO auth → assert 200 + a real row. Also `/server/health`.
4. If 403/empty: diagnose the permission chain live (the flagged risk) → fix or characterize + route.

## Escalation conditions
- If anonymous read needs a running-instance API call or a permission shape no pre-staged rows satisfy (the
  `overview.md` Re-scope risk), record the falsification + route — that is a strategy finding, not a this-tik fix.

## Acceptable close-no-lift outcomes
- Directus boots + the exact anonymous-permission requirement is empirically pinned (even if the GET isn't 200 yet) —
  a first-class characterization of the milestone's hardest, live-only stage.

## Test discipline
Live harness only (throwaway). No stack-snapshot code changed this iter (serve-validation), so the Go suite is
unaffected; if registry/permission loading reveals a code need, route it to `STRUCT-M21-codeify`.
