# M224 — Spec notes

_Iterative milestone: this file accumulates iteration-protocol-specific technical notes (attribution traces,
render-gap findings, demo-patch pin records). Per-iter detail lives in `iter-NN/`._

## Cockpit hero trio

**Recruiter seat LANDED iter-03 (rext tag `casting-call-m224-iter03`)** — content for the `cockpit-spec.md` M224 delivery:

- **Recruiter** (`rae-recruiter`, Rae Ramirez, Technical Recruiter, `vantage: manager` → admin, slot 1,
  `jump_to: /enterprise/activity-dashboard`) — added to the Meridian Talent 4th story in `presets/stories.seed.yaml`.
  Admin → inherits `org:feature:insights` (no net-new grant); slot 1 = admin band so the `HiringFunnelSeeder`
  correctly skips her (recruiters read the scoreboard, they don't audition).
- **A new manager hero role needs a curated skill family** (the iter-03 lesson): a manager renders a *modest
  personal `/profile`* (persona.go gives a manager 3-8 verified + a claimed tail from her role's curated pool), so
  her role must classify to a curated family or `TestShippedPresets_EveryHeroRoleClassifies` fails. iter-03 added
  the `curatedTalent` (recruiting) family in `seeders/curated_pools.go`.
- **`jump_to` is a raw path** (catalog membership cosmetic) — `/enterprise/activity-dashboard` works with a generic
  label; a per-`[simId]` `NeedsID` DeepLinkCatalog entry is optional polish (decide in iter-04 once the seeded sim
  ids are known — the real comparison view is per-position).
- **The 2 candidate exemplars** (one assessed, one only-assigned → `/profile`) are a **post-gate tik** — they need
  candidate-role + funnel-stage hero-awareness (a hero at an admin-band slot, end-user vantage, in a candidate-only
  org — `roleForHero`/`hiringCandidateStage`/the funnel's candidate-detection all need hero-awareness). Not the
  gate metric, so deferred until the recruiter render is proven.

## Clerkenstein isHiring wiring

**LANDED iter-02 (rext commit `d8d9846`, tag `casting-call-m224-iter02`).** The FAPI half of the `is_hiring`
dual-write — content for the `clerkenstein.md` + `hiring.md` M224 delivery:

- **Mechanism (the M224 delivery fills the clerkenstein.md BLIND-AREA):** org `public_metadata.isHiring` is emitted
  by the **fake FAPI** in `clerk-frontend/resources.go::orgMemberships()` (the org resource `@clerk/clerk-js`
  `useOrganization().publicMetadata` reads), fed by the roster thread `RosterEntry.org_is_hiring` →
  `DemoUser.OrgIsHiring`. The seeder produces it: `RosterIdentity.org_is_hiring` ← `ResolvedStory.IsHiringOrg()` in
  `BuildRoster`. Only a hiring story's heroes carry `true`.
- **`/align-run` record (the clerk-frontend BLOCKING guard):** clerk-js-5 **100.0%/100.0%** (9/9), clerk-multi-1
  **100.0%/100.0%** (9/9, incl. Roster 2/2). GREEN — no identity gene perturbed.
- **The align-safety lesson (generalizes — for the clerkenstein.md doc):** when extending a Clerkenstein FAPI
  resource, **emit a new field only when its non-default value applies** (omit it otherwise). The goldens are
  captured from the existing identities; adding a key to a `shape`-graded response (`Client/signed-in`,
  `Me/universal-user`) for the default case would flag the gate or force a golden re-capture. Conditional-emit
  (`if u.OrgIsHiring { pm["isHiring"]=true }`) keeps every non-hiring org's public_metadata byte-identically
  `{eid}`. Generalizes the `Picture`/`OrgLogo` `omitempty` pattern to non-string additions.
- **BAPI (`clerk-backend`) — intentionally NOT wired** (D2): the server derives hiring from the
  `public.organizations.is_hiring` DB column, not Clerk BAPI metadata; a BAPI change adds Go-SDK align surface for no
  render benefit. Optional, only if a server-side consumer reads `organization.publicMetadata.isHiring`.
- **Inert until iter-03:** `BuildRoster` emits identities only for heroes; Meridian Talent still has `heroes: []`.

## Render-gap attribution (the measure→attribute→fix loop)
_Per-iter: what painted, what didn't, the attributed cause (seed-data vs render-gate vs Clerkenstein wiring), the
fix surface chosen._

## Demo-patch pins (if any)
_`next-web-hiring-flag-gate` and/or a perf patch — the sha-pin + the anchor record._

## Pre-flight audits — iter-01

**KB-fidelity gate (Phase 0b): GREEN.** Report: `kb-fidelity-audit.md`. All load-bearing topics PAIRED + ALIGNED;
the two BLIND-AREAs (clerkenstein FAPI org-publicMetadata; cockpit hiring vantage) are M224's own `Delivers →`
deliverables, not un-anchored gaps. One stale-misleading claim fixed inline in `hiring.md` (the isHiring wiring
target is the FAPI `clerk-frontend/`, not the BAPI `clerk-backend/`).

### Topic → doc → code triples (fast-start for later audits)

| Topic | Doc | Code (rext authoring clone) |
|---|---|---|
| Render read-model (score source) | `corpus/services/hiring.md` | `stack-seeding/seeders/{hiring_funnel,persona_write,hiring_config}.go` |
| M223 hiring seed chain | `corpus/ops/demo/stories-spec.md`, `seeding-spec.md` | `stack-seeding/seeders/{hiring_config,hiring_funnel}.go`, `presets/stories.seed.yaml` |
| Clerkenstein org `public_metadata.isHiring` (the FAPI target) | `corpus/services/clerkenstein.md` | `clerkenstein/clerk-frontend/{resources.go::orgMemberships,registry.go::RosterEntry}` (browser-visible); `clerk-backend/{resources.go,store.go}` (BAPI, optional) |
| Cockpit hero trio + DeepLinkCatalog | `corpus/ops/demo/cockpit-spec.md` | `demo-stack/cockpit.py`, `stack-seeding/seeders/cockpit.go` (`DeepLink` struct + `DeepLinkCatalog()`), `presets/stories.seed.yaml` (heroes + `vantage`/`jump_to`) |
| Alignment DNA (clerk-frontend guard) | `corpus/services/clerkenstein.md` §alignment | `alignment/dna/clerk-js-5.json` (genes: `SessionToken/decoded-identity` critical/exact; `Client/signed-in` critical/shape; `Me/universal-user` standard/shape), `alignment/scripts/gate.sh` |

### Wiring facts pinned by the audit (so tiks don't re-trace)

- **`jump_to` is a raw next-web path string** (`cockpit.go::BuildCockpitManifest` → `defaultJumpForVantage`); catalog
  membership is **cosmetic (display label only)** — a hero can jump to `/enterprise/activity-dashboard` with **zero
  catalog change**. Adding a `DeepLinkCatalog()` entry is optional polish (a nicer label / a `NeedsID` for a specific
  `[simId]`). No `NeedsID:true` entry exists anywhere today.
- **The real comparison view is per-`[simId]`** (`/enterprise/activity-dashboard → AI-Simulations → [simId]`). The bare
  `/enterprise/activity-dashboard` shell may land the recruiter on the tab index, not a populated scoreboard — decide in
  a tik whether the recruiter's `jump_to` needs a concrete position id (a `NeedsID`-style entry) vs the shell.
- **vantage enum is only `end-user` | `manager`** (`blueprint.go`). Recruiter → `manager` (like Dan); the 2 candidate
  exemplars → `end-user` (like Maya). No new vantage primitive needed.
- **The 4th story `hiring` (Meridian Talent)** already exists in `presets/stories.seed.yaml` with `heroes: []`
  (intentionally none at M223). M224 adds the recruiter + 2 candidate seats there.
- **Clerkenstein isHiring slot-in:** `clerk-frontend/registry.go` `RosterEntry` (add `org_is_hiring`) →
  `toDemoUser()` → `DemoUser.OrgIsHiring` → `clerk-frontend/resources.go` `orgMemberships()` `PublicMetadata["isHiring"]`
  — the M39 `org_name`/`org_slug` precedent. **Touches `clerk-frontend/` → BLOCKING `/align-run` on `clerk-js-5`.**
- **Drill-down `validation_*` rows are NOT seeded by M223** — if the assessed candidate hero's `/profile` or a
  `[simId]/[userId]` drill-down needs them, that is net-new M224 seeding (the `PersonaSeeder` `persona_write.go:69-71`
  pattern). The comparison SCOREBOARD itself needs only the 2-table mirror pair (already seeded).
