---
iter: 02
milestone: M224
iteration_type: tik
status: closed-fixed
date: 2026-07-16
---

# iter-02 — tik (Clerkenstein org publicMetadata.isHiring wiring)

**Type:** tik — under **TOK-01** (recruiter-render-first). Protocol: `corpus/ops/demo/coverage-protocol.md`
(measure→attribute→fix→re-measure); this tik lands the **fix-half enabling scaffold** for the Clerkenstein-identity
axis of the attribution ladder.

## Step 0 — Re-survey
TOK-01's iter-02 direction named "enabling-scaffold + baseline-render" as one tik. Re-survey (reading the roster /
funnel / Clerkenstein code + Docker at 9.7 GiB) shows that is genuinely **two iters of work** (a careful
align-guarded wiring + hero-seat/funnel-awareness change, THEN a heavy Docker bring-up + a new render-probe). Per
the scope-creep discipline I **decompose under the same TOK-01 strategy** (not a re-scope — the strategy holds):
- **iter-02 (this):** the **Clerkenstein org `publicMetadata.isHiring` wiring** end-to-end (seeder roster export →
  fake FAPI resource) + lockstep/behaviour tests + the **BLOCKING `/align-run`** + rext tag. The single
  highest-risk, most-guarded change; self-contained + independently verifiable.
- **iter-03:** the **hero seats** (recruiter + 2 candidate exemplars) + funnel hero-awareness (pin one candidate
  assessed / one assigned-only; recruiter=admin).
- **iter-04:** the heavy **LOCAL demo bring-up** at the tag + the **render-probe** + the **baseline measurement +
  attribution** (the first gate reading).

## Active strategy reference
**TOK-01** (recruiter-render-first). This tik lands the Clerkenstein-identity fix-half so that, once the hiring-org
heroes exist (iter-03), the org re-skins as hiring and the recruiter's "Results"/insights cohort treatment appears
— the precondition the attribution ladder's step 1 (Clerkenstein identity) needs.

## Cluster / target identified
The client re-skin (`useGetClerkOrganization` → `Boolean(organization.publicMetadata.isHiring)`) reads the org
`public_metadata.isHiring` from the **fake FAPI** (Phase 0b pinned this: `clerk-frontend/resources.go::orgMemberships`,
fed by the `RosterEntry`→`DemoUser` roster thread; the M39 `org_name`/`org_slug` precedent). Today it is unset → a
hiring org renders as a normal Workforce org. Target: wire `isHiring` through the roster + resource.

## Hypothesis
Emitting org `public_metadata.isHiring=true` for a hiring org (and only then) makes `isHiringOrg` true client-side,
which is the documented precondition for the hiring re-skin / "Results" framing. **Metric-neutral this iter** (no
render yet) — it removes attribution-ladder step 1 as a future blocker.

## Expected lift
Gate metric (rows-per-sim) does NOT move this iter (no bring-up). Deliverable = the wiring landed + **`/align-run`
green** + **rext tagged**. Graded on planned-scope landing, not gate-metric.

## Phase plan
1. Clerkenstein: `resources.go` (DemoUser.OrgIsHiring + conditional org publicMetadata) + `registry.go` (RosterEntry
   `org_is_hiring` + `toDemoUser`). **Emit `isHiring` ONLY when true** → the `{eid}`-only golden shape is unchanged
   for every non-hiring identity → `/align-run` stays green.
2. Seeder: `roster.go` (RosterIdentity `org_is_hiring` + `BuildRoster` sets it from `st.IsHiringOrg()`).
3. Tests: extend the lockstep `rosterEntryMirror` (roster_test.go); add a hiring-org BuildRoster test + a Clerkenstein
   round-trip/behaviour test (true→isHiring present; false→key absent, shape preserved).
4. `go test ./...` (clerkenstein + stack-seeding); `gofmt`/`go vet`.
5. **`/align-run`**: clerk-js-5 (jsfapirun) + clerk-multi-1 (multirun) gates — both must stay green (identity genes
   critical/shape; `SessionToken/decoded-identity` exact is unaffected — JWT claims).
6. Commit + **tag** the rext authoring clone; record the tag in spec-notes/decisions.

## Escalation conditions
- `/align-run` RED (any identity gene perturbed) → **user-blocker** (the BLOCKING guard; do not commit a
  gene-perturbing change).
- A shape change that would require re-capturing goldens → surface (a bigger DNA decision).

## Acceptable close-no-lift outcomes
N/A shape — this is a fix-half landing; it closes-fixed on the wiring+tests+align+tag, Gate NOT MET (no render yet).
The BAPI (`clerk-backend`) isHiring is **out of scope** (the server derives hiring from the `is_hiring` DB column,
not Clerk BAPI metadata — Phase 0b finding); documented as optional.
