---
title: "KB Fidelity Audit — M224 the-callback"
date: 2026-07-16
scope: milestone:M224
invoked-by: build-mstone-iters (Phase 0b, iter-01 bootstrap tok)
---

## Verdict

**GREEN** — proceed. No RED blocker. One stale-misleading claim fixed inline (hiring.md wiring pointer);
the two BLIND-AREAs are both explicitly carried as M224 `Delivers → knowledge/corpus` items in `overview.md`
(the skill's sanctioned resolution for a blind area is exactly a `Delivers →` line, already present) — so
development does not start against an un-anchored topic.

## Topic Inventory

| Topic | Knowledge doc | Code paths (rext authoring clone) | Status |
|---|---|---|---|
| Hiring render read-model (the score source) | `corpus/services/hiring.md` | `stack-seeding/seeders/hiring_funnel.go`, `persona_write.go`, `hiring_config.go` | PAIRED — ALIGNED |
| M223 hiring seed chain | `corpus/ops/demo/stories-spec.md`, `seeding-spec.md` | `stack-seeding/seeders/hiring_config.go`, `hiring_funnel.go`, `presets/stories.seed.yaml` | PAIRED — ALIGNED |
| Clerkenstein org `public_metadata` (isHiring wiring target) | `corpus/services/clerkenstein.md` | `clerkenstein/clerk-frontend/{resources.go,registry.go}` (FAPI), `clerk-backend/{resources.go,store.go}` (BAPI) | PAIRED body / BLIND on FAPI org-publicMetadata mechanism — **M224 deliverable** |
| Cockpit hero trio + DeepLinkCatalog | `corpus/ops/demo/cockpit-spec.md` | `demo-stack/cockpit.py`, `stack-seeding/seeders/cockpit.go`, `presets/stories.seed.yaml` | PAIRED (single-CTA/jump_to/profile ALIGNED); BLIND on hiring vantage — **M224 deliverable** |
| Iteration protocol (render loop) | `corpus/ops/demo/coverage-protocol.md`, `verification.md` | `stack-verify/e2e/` (Playwright) | PAIRED — current (M219) |
| Demo-patch escape hatch | `corpus/ops/demo/demopatch-spec.md` | `demo-stack/patches/` | PAIRED (only consulted if a render-gate needs it) |

## Fidelity Findings

1. **hiring.md — the isHiring wiring pointer is STALE-misleading (FIXED inline).** hiring.md §"the gate" said
   "Today Clerkenstein emits `{eid}` only (`clerkenstein/clerk-backend/resources.go:38-47`)". The citation resolves,
   but it points at the **BAPI** (server-side); the client re-skin (`useGetClerkOrganization.tsx:20-21`) reads the
   **FAPI** org `publicMetadata`, emitted independently at `clerk-frontend/resources.go` `orgMemberships()` (fed by
   the `RosterEntry`→`DemoUser` roster thread — the M39 `org_name`/`org_slug` precedent). Following the old pointer
   verbatim would send an M224 author to the wrong file. **Fix applied:** added a correction callout naming the FAPI
   as the browser-visible emission point + the roster+resource slot-in + the `/align-run` guard + the BAPI's optional
   status. Verdict: ALIGNED after fix.
2. **hiring.md read-model vs M223 seeders — ALIGNED (6/6 claims).** The mirror pair (`jobsimulation.sessions` +
   `public.local_jobsimulation_sessions`, score on the mirror, misspelled `completition_status`), the funnel shape
   (50-size org → 5 admin/45 candidate, ~10% assigned-only, differentiated 30–95 score spread ≥20 distinct), the
   no-net-new-insights-grant claim (admins inherit `org:feature:insights` via the standard g2 admin grant), the
   drill-down `validation_*` gap (NOT written by HiringFunnelSeeder — M224+), and the 5-position `readHiringSimPool`
   selection all check out against the code + its unit tests. Cross-ref anchors in the "IMPLEMENTED as of M223"
   paragraph resolve.
3. **clerkenstein.md — the FAPI org-publicMetadata mechanism is a BLIND-AREA (expected M224 deliverable).** The doc
   narrates the same `orgMemberships()` plumbing for M39 org-name/slug but never says that function also builds
   `public_metadata`. Nothing false; a genuine documentation gap that hiring.md's own cross-ref (line 197) already
   promises M224 fills. Alignment DNA confirmed real (`alignment/dna/clerk-js-5.json`): the named
   `SessionToken/decoded-identity` critical/exact gene scores JWT claims (unaffected by an org-metadata field); the
   org resource sits behind `Client/signed-in` (critical, **shape**) + `Me/universal-user` (standard, shape) — those
   must stay green under `/align-run`.
4. **cockpit-spec.md — hiring vantage is a BLIND-AREA (expected M224 deliverable); everything it claims is ALIGNED.**
   Single-CTA model, `jump_to`-is-a-raw-path (catalog membership is cosmetic label-only), and the `/profile`
   employee precedent (Maya) all check out. No `NeedsID:true` entry exists anywhere; no `/enterprise/activity-dashboard`
   catalog entry exists (confirms overview's "none exist today"). The 4th story "hiring" (Meridian Talent) exists in
   `stories.seed.yaml` with `heroes: []` — intentionally none at M223; the trio materializes at M224.

## Completeness Gaps

- The two blind areas above (FAPI org-publicMetadata; cockpit hiring vantage) are the milestone's own delivery
  surface — not pre-existing debt. No critical undocumented behavior exists in code M224 will read as a contract.

## Applied Fixes

- `corpus/services/hiring.md` — added the FAPI-vs-BAPI wiring correction callout (finding 1).

## Open Items (require user decision)

- None. (The runtime stat "393/393 local rows on billion carry the pair" is unverifiable without DB access; it is
  M222 live-probe evidence, not an M223-seeder claim, and load-bearing on nothing M224 does.)

## Gate Result

**GREEN: proceed.** The bootstrap tok may author its strategy against verified docs. The two blind areas are the
milestone's declared deliverables (into `clerkenstein.md`/`hiring.md` and `cockpit-spec.md`); the one stale pointer
is fixed.
