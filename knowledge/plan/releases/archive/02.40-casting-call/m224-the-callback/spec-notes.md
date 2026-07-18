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

### iter-05 — THE ATTRIBUTION: a platform cross-product redirect ejects the recruiter (DESIGN WALL)

**Symptom.** The `rae-recruiter` render-probe never settles — it hangs on the **real** Clerk sign-in
(`app.anthropos.work`), whose after-sign-in target decodes to
`hiring.anthropos.work/switch?to=hiring&org=org_8ff48bf32d61&next=/home`.

**Isolation (empirical, on the warm demo-1 — both seats driven through the SAME `loginAs`):**

| Seat | Login lands | Direct `goto /enterprise/activity-dashboard` | External eject | Surface |
|---|---|---|---|---|
| `dan-manager` (workforce org) | `localhost:13000/home` | **STAYS local** — `localhost:13000/enterprise/activity-dashboard` | none (only a harmless `support.anthropos.work` **script** widget) | RENDERS (main len 469) |
| `rae-recruiter` (all-hiring org) | briefly `/home`, then | **`document` nav to `hiring.anthropos.work/switch?...&org=org_8ff48bf32d61`** (aborted by the probe so it can't hang) | **top-frame document eject, fires TWICE** (post-login AND post-direct-goto) | **RENDERS NOTHING** (main len 0) |

So the demo login + `apps/web /enterprise/activity-dashboard` **work in general** (dan proves it). The recruiter
divergence is **NOT the seat, NOT Clerkenstein, NOT a seed gap** — the fake session IS established (the app knows
`org_8ff48bf32d61` = Meridian and reads `isHiring=true`; the M224 iter-02 FAPI wiring WORKED).

**Root cause (code-cited, `stack-demo/next-web-app`):** an all-hiring-orgs user is hard-redirected OUT of the
Workforce app to the standalone Hiring product:

- `apps/web/src/context/UserStatusContext.tsx:141-173` — a `useEffect` in `UserStatusProvider` (which wraps the
  **entire** `(authenticated)/(verified)/` subtree, incl. `/enterprise/activity-dashboard`): if
  `memberships.every(m => m.organization.publicMetadata.isHiring)` → `window.location.href =
  buildSwitchHandoffUrl({ targetProduct: 'hiring', … })`. It has an explicit comment: *"this still redirects a
  Workforce-less user to Hiring."* **A GLOBAL guard — direct-navigating to the scoreboard does NOT bypass it**
  (empirically the eject fires on the direct goto too).
- `apps/web/src/hooks/useGetClerkOrganization.tsx:16-18` — `regularOrganizationList = filter(!isHiring)`, so a
  hiring org is also **excluded from the Workforce org list** (the verified layout would treat it as foreign for a
  mixed user).

**The M222-premise tension (the crux).** M222 traced the scoreboard rendering in `apps/web /enterprise` and wrote
*"survives the `is_hiring` flip, no route guard."* That held on the **billion** substrate **only because the org
had NO client `publicMetadata.isHiring`** — client-side it was a *workforce* org (`isHiring=false`), stayed active,
never tripped the all-hiring eject; the scoreboard rendered because the **server** derives the hiring cohort from
the `public.organizations.is_hiring` **DB column**, independent of Clerk. **M224 Scope.In #3 wired client
`isHiring=true` (required for the "Results" re-skin / "reads as hiring end-to-end") — the SAME flag the platform
uses to eject the user.** ⇒ **On the unmodified platform, "reads as hiring in the browser" and "scoreboard
reachable in apps/web" are MUTUALLY EXCLUSIVE.** The milestone wants both.

**Attribution class:** RENDER-GATE / platform routing wall (R1 realized), **NOT** seed-gap. The M223 data side is
MET (Meridian `is_hiring=true`, 50 members, 5 sims × 43 candidates, scores 27–100). The block is 100% the
cross-product redirect.

## Demo-patch pins (if any)
_`next-web-hiring-flag-gate` and/or a perf patch — the sha-pin + the anchor record._

### iter-05 — RECOMMENDED (escalated, NOT yet applied — needs a ~20-min next-web rebuild)

**`next-web-hiring-cross-product-redirect`** — a sha-pinned demo-patch (per `corpus/ops/demo/demopatch-spec.md`)
that neutralizes the `userHasAllHiringOrgs` eject in the demo's OWN ephemeral `apps/web/src/context/
UserStatusContext.tsx` clone (early-return before the `window.location.href` hop), so an all-hiring recruiter STAYS
in the demo's `apps/web`. **A SINGLE file suffices** — verified by reading the downstream guards:
`useResolveActiveOrg` early-returns on an empty workforce list (never auto-picks; `:93-97`) and
`isActiveOrgForeign(Meridian, [])` returns `false` (empty `orgIds` short-circuits `:99-104`), so with only the
redirect removed the verified layout becomes ready, `isHiringOrg=true` drives the "Results" re-skin, and the
scoreboard query scopes to Meridian's `eid`. This keeps the M224 genuine-hiring-org requirement AND makes the
surface reachable. **Cost: a next-web image REBUILD (~20 min) — the reason this was escalated, not auto-applied.**
This is exactly the D-DESIGN-2 order-of-preference outcome (Clerkenstein-wiring-first was tried in iter-02 and is
correct; the residual is a platform routing wall with no env/config/Clerkenstein seam → sha-pinned demo-patch).

**Alternatives handed to the orchestrator** (see the iter-05 hand-off): (A) drop the client `isHiring` emit — no
rebuild, scoreboard renders DB-driven, but LOSES the re-skin (regresses Scope.In #3 / D-DESIGN-1); (C) reframe the
recruiter as a manager of a MIXED org set — does not work (a hiring org can never be the active org in apps/web, so
the scoreboard can't scope to it). Option B (the demo-patch) is the only path that satisfies both halves.

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
