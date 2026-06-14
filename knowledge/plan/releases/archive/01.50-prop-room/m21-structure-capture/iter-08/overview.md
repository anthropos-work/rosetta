---
milestone: M21
iter: 08
iteration_type: tik
created: 2026-06-13
status: closed-fixed
---

# M21 iter-08 — the SERVE half: registration + permissions capture/apply + firewall admissibility

## Active strategy reference
**TOK-01** (staged-pipeline build toward the binary serve-anonymously gate), as refined by **M21-D7**
(option A — full-content-model + version-pin) and **M21-D9** (the anonymous-serve recipe). This is the
gate-completing tik: stages 5-6 ("serve anonymously") were DEMONSTRATED by hand in iter-05; the apply side
of the SCHEMA (stages 3-4) was AUTOMATED in iter-07. iter-08 automates the remaining serve rows so the gate
is met **by tooling** (no hand SQL).

## Step 0 — re-survey (TOK-directed target still current?)
TOK-07's route-forward named `STRUCT-M21-iter08-serve` as the next target. Re-survey confirms it is still the
exact residual: iter-07's `stacksnap` captures + auto-provisions the schema (tables+PKs → digest converges →
rows replay, exit 0) but NOT the `directus_collections` registration rows nor the public-policy
`directus_permissions` read rows — those were hand-applied in iter-05 (`serve.sql`). Target current; no substitution.

## Cluster / target identified
The gate's clause "a booted Directus serves a captured public simulation over HTTP to an ANONYMOUS reader" is
DEMONSTRATED but NOT yet tooling-automated. The residual (M21-D9 serve recipe, beyond the schema iter-07 ships):
1. a `directus_collections` registration row per served collection, and
2. a `directus_permissions` `read` row on Directus's hardcoded public policy `abf8a154` per served collection.
Both live in `directus_*` SYSTEM tables, today outside `AssertPublicOnly` — so a **firewall structural-metadata
admissibility class** is needed (extend, never loosen). The iter-08 design is pre-de-risked in
`spec-notes.md` § "iter-08 design" (sanctioned prod structural read confirmed: these two tables carry 0
tenant-scope columns; `directus_access`/`directus_policies` are EXCLUDED — bootstrap provides them).

## Hypothesis
Capturing the two serve-row registry tables (filtered to the served collections + the public policy) as
additive INSERTs appended to the structure artifact, gated by a firewall admissibility check that asserts they
carry no tenant-scope column, then applied by the existing `tryAutoProvision` path — will let a freshly
bootstrapped + stacksnap-provisioned stack serve the captured catalog anonymously with NO hand SQL. That flips
the M21 exit_gate from DEMONSTRATED → MET-by-tooling.

## Expected lift
Furthest pipeline stage passing stays 6 (already demonstrated), but the **gate automation clause** flips from
"schema half done (iter-07)" to **MET** — the binary gate is satisfied by `stacksnap` alone. Success = the
live gate proof (capture → fresh bootstrap → `stacksnap replay` → boot → anonymous 200) runs with zero hand SQL.

## Phase plan (staged-pipeline protocol, TOK-01)
- **Phase A — build:** the serve-row capture (registry SELECTs → faithful additive INSERTs) + the firewall
  structural-metadata admissibility class + wire the serve rows into the structure artifact + apply (the existing
  auto-provision path runs them).
- **Phase B — test gate:** `go vet` + `go build ./...` + `go test ./...` (capture/firewall/directus/cmd) green;
  new unit tests for the serve-row SQL assembly + the firewall admissibility (carries-tenant-column → reject).
- **Phase C — re-measure (live):** capture against the sanctioned prod `--dsn`/MCP, fresh-bootstrap a Directus,
  `stacksnap replay` (auto-provision), boot, anonymous `GET /items/simulations?limit=1` → 200 + real row. Confirm
  NO hand SQL ran. Re-confirm the exit matrix (empty→4, gap+no-structure→5, diverged→5, gap+structure→0) holds.
- **Phase D — close:** grade, write close section, commit (rosetta-extensions + the milestone state).

## Escalation conditions
- If the serve rows cannot be captured public-only (a registry table unexpectedly carries a tenant-scope
  column the sanctioned read did not show) → the firewall admissibility check fails LOUD; surface as user-blocker
  (the firewall must not be loosened).
- If the additive INSERT apply collides irrecoverably with bootstrap's system rows in a way `ON CONFLICT`/omit-id
  can't resolve → route forward + surface.
- Otherwise: land + route any non-gate residual (directus_files, M23) forward.

## Acceptable close-no-lift outcomes
If the live serve proof surfaces a NEW load-bearing requirement beyond registration+permissions (e.g. a
directus_fields row IS required after all for a typed column), characterize it, route the automation forward,
and close-no-lift with the falsification recorded — the gate stays demonstrated, the automation gap is named.
