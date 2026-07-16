# iter-02 — progress

**Type:** tik — under TOK-01 (recruiter-render-first). Protocol: `corpus/ops/demo/coverage-protocol.md`. This is
the **Clerkenstein-identity fix-half** of the enabling scaffold (decomposed from TOK-01's iter-02 direction per
the scope-creep discipline — see overview Step 0).

## What this iter did
Wired org `public_metadata.isHiring` end-to-end so a demo hiring org re-skins in the browser (removes
attribution-ladder step 1 as a future render blocker), verified by the BLOCKING `/align-run` + unit tests, and
tagged the rext authoring clone.

- **`clerk-frontend/resources.go`** — `DemoUser.OrgIsHiring`; `orgMemberships()` emits org
  `public_metadata.isHiring=true` **only when true** (a false value omits the key → the `{eid}`-only shape is
  byte-identical for every non-hiring identity → alignment goldens unchanged).
- **`clerk-frontend/registry.go`** — `RosterEntry.org_is_hiring` (omitempty), threaded via `toDemoUser()`.
- **`stack-seeding/seeders/roster.go`** — `RosterIdentity.org_is_hiring`, set from `ResolvedStory.IsHiringOrg()`
  in `BuildRoster` (only Meridian Talent's heroes carry true).
- **Tests (3 new):** the lockstep mirror extended (`roster_test.go`); `TestBuildRoster_HiringStorySetsOrgIsHiring`
  (true for the recruiter hero, false+omitted for Workforce heroes, exactly-once in the JSON); the Clerkenstein
  conditional-emit test (`TestDemoUser_orgMemberships_IsHiringConditional` — true→present, false→absent, shape=={eid});
  the roster round-trip (`TestRoster_ThreadsOrgIsHiring`).

## Verification
- `go test ./...` GREEN on both modules (clerkenstein full suite + stack-seeding seeders); `gofmt`/`go vet` clean.
- **BLOCKING `/align-run` GREEN:** clerk-js-5 **100.0%/100.0%** (9/9 genes, no divergences); clerk-multi-1
  **100.0%/100.0%** (9/9, incl. Roster 2/2). No identity gene perturbed — the conditional emit preserves every
  golden's org-resource shape.
- rext commit **d8d9846**, tag **`casting-call-m224-iter02`**.

## Attribution note (for the render loop)
This tik lands the **Clerkenstein-identity axis** of the attribution ladder (hiring.md read-path step 1). It is
**inert until a hiring-org hero exists in the roster** (`BuildRoster` only emits identities for heroes; Meridian
Talent still has `heroes: []`) — iter-03 adds the seats, iter-04 brings up + measures.

## Close — 2026-07-16

**Outcome:** Clerkenstein org `publicMetadata.isHiring` wired end-to-end + `/align-run` GREEN (100/100 ×2) +
rext tagged `casting-call-m224-iter02`. Planned scaffold scope landed fully.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (no render measured yet — this is a fix-half/scaffold tik under TOK-01; the first render
reading is iter-04)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this is a tik; and only 1 tik exists so the 3-no-prog window can't fire) — (3) re-scope: n — (4) user-blocker: n (`/align-run` GREEN — the milestone's hard guard passed) — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue (→ iter-03)
**Decisions:** D1 (conditional-emit for align-safety); D2 (BAPI out of scope); D3 (scope-decompose iter-02→02/03/04). See iter-02/decisions.md.
**Side-deliverables:** none.
**Routes carried forward:** iter-03 = hero seats (recruiter + 2 candidate exemplars) + funnel hero-awareness
(pin one candidate assessed / one assigned-only; recruiter=admin) + re-tag rext. iter-04 = LOCAL demo bring-up at
the tag + render-probe + baseline measurement + attribution (the first gate reading).
**Lessons:** For any Clerkenstein FAPI resource extension, **emit the new field only when its non-default value
applies** (omit otherwise) so the golden shapes — captured from the existing (non-hiring) identities — are
unchanged and `/align-run` stays green without a golden re-capture. Generalizes the Picture/OrgLogo `omitempty`
pattern to non-string additions. (Recorded in `spec-notes.md` — the iterative milestone's protocol-notes
accumulator — as content for the `clerkenstein.md` M224 delivery; not folded into the Playwright-focused
coverage protocol, which the lesson doesn't concern.)
