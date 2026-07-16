# M224 — Spec notes

_Iterative milestone: this file accumulates iteration-protocol-specific technical notes (attribution traces,
render-gap findings, demo-patch pin records). Per-iter detail lives in `iter-NN/`._

## Cockpit hero trio
_1 recruiter/manager (`vantage: manager` → jump_to the comparison surface) + 2 candidate exemplars (one assessed,
one only-assigned) → /profile. DeepLinkCatalog entries + labels._

## Clerkenstein isHiring wiring
_FAPI/BAPI emission of org `public_metadata.isHiring=true`; the `/align-run` record for the clerk-frontend change._

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
