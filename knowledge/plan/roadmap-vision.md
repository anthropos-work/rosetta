# Roadmap — Vision

Future versions and proposals, not yet in active development. The active version lives in
[`roadmap.md`](roadmap.md). When a version starts development, its section moves from here into `roadmap.md`
and a `release/{version}` branch is cut.

> **Promotion history:** **v1.0 "body double"** → 2026-06-02 (shipped 2026-06-03, tag `v1.0`).
> **v1.1 "show floor"** → 2026-06-03 (shipped 2026-06-05, tag `v1.1`).
> **v1.2 "set dressing"** → 2026-06-05 (shipped 2026-06-07, tag `v1.2`).
> **v1.3 "stack party"** → 2026-06-07 (shipped 2026-06-07, tag `v1.3`; the **dev/demo convergence** — dev stacks as
> first-class peers, a unified first-available-N stack registry, generic `stack-*` skills, a code-cited safety doc).
> **v1.3b "dress rehearsal"** → 2026-06-08 (shipped 2026-06-09, tag `v1.3.1`; the **field-hardening release** for
> the 14 issues the first real `/demo-up` run surfaced — `/demo-up` now produces a full/populated/verified/demoable
> stack, M16→M20; tooling + docs only, zero platform-repo edits).
> **v1.5 "prop room"** → 2026-06-11 (shipped 2026-06-14, tag `v1.5`; the **local-Directus release** — a real per-stack
> Directus serving the captured public content, demo-default + dev-opt-in, M21→M25). The first version staged after the
> v1.4 removal.
> **v1.6 "stage door"** → 2026-06-14 (shipped 2026-06-14, tag `v1.6`; the **secret-provisioning release** — one
> mechanism that ingests a secret source [dir/zip, default `.agentspace/secrets`] and provisions every repo of a stack,
> with a secret-coverage DNA that lists + keeps-listed the required secrets per repo, M27→M30). Requested directly by
> the user, not from prior backlog.
> **v1.7 "house lights"** → 2026-06-15 (shipped 2026-06-15, tag `v1.7`; the **demo-UI-hardening release** — a fresh
> browser at a demo's offset UI renders with zero manual steps: M31 a locally-trusted **mkcert** FAPI cert [so next-web
> stops blanking] + M32 the studio-desk single-port/production fix, M31→M32; tooling + docs only, zero platform-repo
> edits). Triggered by a live next-web blank-page defect, not from prior backlog.
> **v1.8 "understudy"** → 2026-06-15 (shipped 2026-06-15, tag `v1.8`; the **self-contained-demo
> release** — give `stack-demo/` its own platform clone set so a box with only `stack-demo/` runs a demo end-to-end:
> a single `section` milestone **M26** that re-implements the orphaned `m26/self-contained-demo` branch onto current
> `main`, preserving v1.6/v1.7; tooling + docs only, zero platform-repo edits). **Graduated from the unscheduled
> backlog** (the orphaned ext effort) on the user's "fill just that gap" go-ahead.
> **v1.9 "storytelling"** → 2026-06-22 (shipped 2026-06-23, tag `v1.9`; the **believable-demo-narrative release** —
> convert the placeholder seeder into a declarative **Stories & Heroes** engine: per-story org + a thriving/struggling/manager
> hero trio seeded via the real **verified-skill chain**, so the **skill profile** + the **Workforce dashboard**
> tell a story, plus a **presenter cockpit**; 5 `section` milestones M34→M38; tooling + docs only). Designed from
> the adversarially-verified spec [`.agentspace/seeding_gaps.md`](../../.agentspace/seeding_gaps.md). The first
> version since v1.5 to come from a substantive backlog/spec rather than a live defect.
> **v1.10 "method acting"** → 2026-06-24 (shipped 2026-06-27, tag `v1.10`; the
> **believable-profile release** — make each hero hold up under a close-up when a presenter *Logs in as* them:
> profile identity [org name + role + real-face avatar] + the content-surface unblock [library + activity feed via
> one Directus serve-grant] + profile depth [work/education/deep skills] + **100% per-vantage demo coverage** proven
> by Playwright, **EXTENDED** with cockpit-UX + whole-roster completeness + a cheap-LLM generation engine + org-scale
> fill; 9 milestones M39→M46 [section + iterative]; tooling + docs only). Designed from the
> live-demo review [`.agentspace/profile_gaps.md`](../../.agentspace/profile_gaps.md) — a hero (Maya Chen) read
> empty across the profile/library/activity surfaces. The second consecutive demo-believability release after v1.9,
> and the **last release of the v1.x major** — the v1.x history now lives in [`roadmap-legacy.md`](roadmap-legacy.md).
> **v2.0 "opening night"** → 2026-06-28 (IN DEVELOPMENT, branch `release/02.00-opening-night`; a **NEW MAJOR** —
> opens the **Playthroughs** pillar: functional-flow *testing*, a manifest-driven deterministic e2e suite that
> *pretends to be the human* and proves the platform's core user journeys **actually work** — distinct from the
> v1.x demo/seeding lineage. 4 milestones M201 ∥ M202 → { M203 ∥ M204 } [`Mxyy` numbering]: M201 Manifest corpus
> [`iterative`, user-guided] ∥ M202 Playthroughs Foundation [`section`] → M203 employee-vantage coverage ∥ M204
> manager-vantage coverage [both `iterative`]; tooling + docs only — zero platform-repo edits). Designed from the
> consolidated capability spec [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) v0.3. The
> **Playthroughs** capability **graduated from spec-draft to active development** here.
> **v3.0 "open house"** → 2026-07-01 (**PROPOSAL — awaiting review**; a **NEW MAJOR** proposal — opens the
> **Customer API + MCP** pillar: a stable, programmatic surface [REST façade + MCP server] that lets a
> customer's script or AI agent read/write against their tenant without going through the UI. Distinct from
> the v2.x Playthroughs lineage — v2.0 proves the *user* journeys work; v3.0 opens those same domain
> resources to *machines*. 4 milestones M301 → M302 → M303 → M304 [`Mxyy` numbering]: M301 discovery + identity
> seam [`section`] → M302 access primitive [`section`] → M303 REST reads gateway [`iterative`, exit gate =
> UC1–UC7 green with 0 cross-tenant leakage over 5 runs] → M304 customer surface + docs-lite [`section`];
> strictly sequential. R1 is READ-ONLY end-to-end; W1/W2 writes are R2+. Governed by the consolidated pillar
> spec [`spec-drafts/customer-api-mcp/spec.md`](spec-drafts/customer-api-mcp/spec.md) v0.1 + its companions
> [`spec-progress.md`, `next-release.md`, `vision.md`]).

---

> **v2.0 "opening night" is IN DEVELOPMENT** (designed 2026-06-28; branch `release/02.00-opening-night`; full
> detail in the `## In Development — v2.0` section of [`roadmap.md`](roadmap.md) — the active roadmap now holds the
> v2.0 major; v1.x history is in [`roadmap-legacy.md`](roadmap-legacy.md)). The future v2 milestones + the
> unscheduled backlog below are orthogonal (not in v2.0 scope).

## Future versions

### v3.0 "open house" — PROPOSAL (2026-07-01), NOT yet promoted

A **NEW MAJOR** — opens the **Customer API + MCP** pillar: a stable, programmatic surface (REST façade + an
MCP server) that lets a customer's script or AI agent read/write against their tenant without going through
the UI. Distinct from the v2.x Playthroughs lineage — v2.0 proves the *user* journeys work; v3.0 opens those
same domain resources to *machines*. Governed by the consolidated capability spec
[`spec-drafts/customer-api-mcp/spec.md`](spec-drafts/customer-api-mcp/spec.md) v0.1 + its companions
(`spec-progress.md`, `next-release.md`, `vision.md`).

**Codename:** _open house_ continues the theatre lineage — after opening night proves the show works, the
venue opens its doors to guests (here: API consumers + AI agents).

**4 milestones, sequential** (Mxyy numbering, major digit 3):

- **M301 — Discovery + identity seam** (`section`): the `customer-api-mcp` package skeleton in
  `app/internal/customerapi/` + the OpenAPI-3.1-based `catalog.yaml` registry (with `x-anthropos-*` extension)
  + the loader/validator + the `Principal` DTO + the `IdentityProvider` adapter port + the first adapter
  `ClerkIdentityProvider` (with the `ClerkGuardrails` static-lint fence) + a smoke `/v1/access/whoami` handler.
  Proves the auth-vendor-swap boundary.
- **M302 — Access primitive** (`section`, depends on M301): the API-key store (Postgres, argon2id-hashed) +
  the `ApiKeyIdentityProvider` (2nd adapter) + mint/rotate/revoke via 3 admin-tier catalog entries + the
  append-only audit ledger + the shared audit-write middleware + the Redis token-bucket rate-limit middleware
  + a proof `/v1/access/whoami` + `/v1/access/api-keys` (list). No customer-data endpoint yet — the primitive
  is proven with an echo.
- **M303 — REST reads gateway** (`iterative`, depends on M302, exit gate: UC1–UC7 green with 0 cross-tenant
  leakage over 5 consecutive runs, iteration protocol = the pillar spec): the R1 READ catalog — 7 resources
  across 4 products (People `member.list`+`.get`, Learning `skill-path.list`+`.get`+`session.list`,
  Verification `verified-skill.list`, Audit `event.list`), each closed-per-resource on: contract test +
  cross-tenant isolation test + audit row + rate-limit fire.
- **M304 — Customer surface + docs-lite** (`section`, depends on M303): the Workforce API-key page
  (list/mint/rotate/revoke) + the `docs.anthropos.work/api/v1/` static site (OpenAPI-rendered reference +
  4 Quickstarts for UC1–UC4 + Principles page + entitlement-tier page). R1 becomes visible + usable.

**Execution graph:** `M301 → M302 → M303 → M304` (strictly sequential — every downstream milestone reads
outputs from the prior one; no parallelism in R1). The **MCP server** is a fast-follow after R1 (see
`next-release.md` — parked for R2), not in-scope here.

**Scope boundary:** R1 is **read-only** end-to-end. All writes (W1 safe writes, W2 advanced writes) are R2+.
The write staging model + a mutation-gap inventory are in the spec (§4 + Appendix A), enforcing the
"real-mutations-only, escalate-never-shim" principle (P1) when writes land.

**Milestone dirs scaffolded** under [`releases/03.00-open-house/`](releases/03.00-open-house/) — awaiting
review and promotion. Once promoted, `release/03.00-open-house` branch cuts from `main` post-v2.0 ship, and
this section moves into `roadmap.md`.

## Future v2 milestones (Playthroughs pillar — NOT yet clustered into a minor version)

### Future v2 milestones (Playthroughs pillar — NOT yet clustered into a minor version)

The Playthroughs capability has **graduated from spec-draft to active development** in v2.0 "opening night"
(M201 ∥ M202 → { M203 ∥ M204 }). The milestones below are the **declared-but-deferred** Playthroughs surfaces —
real future work, **not** pre-assigned to a minor version (per the `Mxyy` rule, a future major's milestone gets
its number at *design* time, not before). They are governed by the same capability spec
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md).

- **M205 — Hiring + tier gates.** The **recruiter** vantage: candidate-pipeline journeys (post a role → review
  applicants → advance a candidate) + the **free→paid entitlement gates** (a flow that's gated by tier proves the
  gate fires). Needs a **new `HiringSeeder`** (a recruiter persona + a candidate pipeline the seed populates) +
  a **Stripe test-mode / assertion-boundary** for the paid-tier gate (Stripe is one of the integration third
  parties with no mirror today — spec §5.8; test-mode is the likely mirror).
- **M206 — AI-sim mirror tier.** The signature **voice / recording** AI-simulation journey — **voice (LiveKit)**
  + **recording (Chime)** — driven at the **launch / completion assertion boundary** (the flow launched, reached
  an interactive state, the outcome artifact materialized), NOT turn-by-turn. Needs **mirror engines**:
  **Clerkenstein only mocks Clerk** — there is no LiveKit/Chime mirror yet (spec §5.8). Until a mirror lands,
  these legs are parked as `later — needs a mirror engine`.
- **M207 — Academy coverage** — Playthroughs over the **separate ant-academy deployment** (its own Vercel-deployed
  app, Clerk-only, not in the platform docker-compose). A distinct target environment from the demo-N hero stacks;
  a future surface for the capability.

### Future v3 minors (Customer API + MCP pillar — NOT yet clustered)

After v3.0 R1 ships, the pillar's post-R1 arc is spec'd in
[`spec-drafts/customer-api-mcp/next-release.md`](spec-drafts/customer-api-mcp/next-release.md): **R2** the MCP
server + W1 safe writes + interactive OpenAPI + SDKs, **R3** W2 advanced writes + entitlement enforcement,
**R4–R6** the developer-portal + billing/quotas + partner tier. No pre-assigned Mxyy numbers (per the rule, a
future minor's milestone gets its number at *design* time, not before).

## Unscheduled backlog (not a planned release)

Genuinely-deferred work, no target version, not scheduled:

- **DEF-M10-01 — cloud `SnapshotStore` backend + S3 media blob bytes.** Today the cache is the local
  `.agentspace/snapshots/` store and media replays **refs-only**. **Re-signed → backlog at v1.5 design (2026-06-11)**
  after its v1.4 destination was removed; its **user-facing sting is gone** — v1.5 "prop room" keeps the asset plane
  on prod public links so demos show **real images** without the blob-byte work. Real blob mirroring + the cloud
  store stay gated on **eu-west-1 S3 read access actually landing** (verified not wired). Replay-only to a per-stack
  isolated bucket, never the shared prod S3.
- **DEF-M21-01 — `replayCmd` conn-seam hermetic test.** A hermetic `replayCmd`-wiring test needs an injectable
  connector seam (>50 lines, touches the load-bearing replay path). Tracked KEEP across the M21→M25 close-audits;
  **landed here at v1.5 close-release (2026-06-14)** so it survives the release-branch merge. Pick up in a future
  `stack-snapshot` build iter when the replay path is next opened.
- **M25-D9 — dev-`N` taxonomy replay `rc=4` ("target schema empty").** A pre-existing dev-stack migrate-ordering
  nuance on opt-in `dev-N≥1 --local-content` stacks (non-fatal, orthogonal to the content-serve path — the directus
  content-serve done-bar DB-2 is GREEN). Surfaced by the M25 field-bake; tracked dev migrate-ordering follow-up.
- **DEF-M46-01 — Directus serve-grant CLOSURE + schema RECAPTURE (Option B) → RESOLVED (M46 Path 2,
  `method-acting-m46-servegrant-closure`).** M46/DD first landed a **targeted** column reconciliation (Option A): the
  captured per-stack Directus structure had **drifted** behind the platform (cms's `SetFields("*", …)` simulations query
  SELECTs `simulations.is_interview_validation_enabled`, a column added to prod Directus after the snapshot was captured →
  `Directus 500: column does not exist`); Option A is an idempotent **post-replay `ADD COLUMN IF NOT EXISTS`** backfill in
  `up-injected.sh` (the FK-indexes mechanism class; `DEMO_NO_DIRECTUS_DRIFT_FIX` opt-out) — kept. Option A fixed the
  column 500, but the cold sweep surfaced a DEEPER, distinct blocker on `/enterprise/activity-dashboard`: the M40
  serve-grant `servedCollections` set (`stack-snapshot/directus/structure.go`) was **incomplete for the full cms
  `GetJobSimulation` deep-fetch closure**. cms requests `sequences.knowledge.*`,
  `sequences.assets_files.directus_files_id.*`, `sequences.collaborative_assets.*`, `sim_features.*`, `translations.*`,
  but the target/junction collections — `knowledge_asset`, `sequences_files`/`_2`, `directus_files`, `sim_features`,
  `sim_roles_tasks`, `sim_translations`, `simulations_translations` — were NOT registered/granted/related → Directus
  dropped the WHOLE parent `sequences` alias → cms `jobsimulation.go:1097 s.Sequences[0]` panicked (`index out of range`)
  → `jobSimulation.title` null → the federation failed the non-nullable field → the activity-table never hydrated.
  **CLOSED by M46 Path 2 (the Option-B durable fix):** (a) EXPANDED `servedCollections` to the 7 deep-fetch closure
  collections + a SYNTHESIZED `directus_files` SYSTEM public-read grant (`serveFilesCollection`/`serveFilesPermissionSQL`,
  read-only); (b) RECAPTURED the prod Directus structure over the sanctioned `marco_read` DSN (firewall public-only,
  `public_only=true`, 0 tenant rows; the relation/field metadata is captured from prod, never hand-fabricated — the digest
  was unchanged so the capture overwrote the cached `_structure.sql` in place: relations 35→45, fields 239→294, perms +8).
  A fresh `/demo-up` replays the regenerated cache and self-applies the closure; the anonymous deep-fetch now preserves the
  `sequences` alias (no panic) and the activity-dashboard renders real per-sim content. See `corpus/ops/snapshot-spec.md`
  → "The GetJobSimulation deep-fetch closure (M46 …)". rext tag `method-acting-m46-servegrant-closure`.

**Resolved (no longer backlog):**

- **M33 — ant-academy demo liveness** (deferred from v1.7 design, 2026-06-15, repro-first) → **RESOLVED post-v1.9** at
  rext tags `storytelling-postfix-1` (session-reaping fix) + `storytelling-postfix-2` (the blocked-clone / token-residual
  fix — ant-academy now auto-comes-up on a fresh `/demo-up`), both tooling-only post-v1.9 demo-hardening passes. The "dead on a later visit"
  reaping was **REPRODUCED + FIXED**: the host-native daemons were launched via `nohup` alone, which does **not**
  detach from the launcher's process group — so when a backgrounded `/demo-up` task's process tree was reaped on
  completion (or the launching session ended), the daemon died with it (the exact M33 hypothesis). Both ant-academy
  and the presenter cockpit now launch **session-detached** via a shared `demo-stack/detach.sh::launch_detached`
  (`setsid` where present; a portable `python3 os.setsid` double-fork on macOS, which has no `setsid`), so they
  survive the launching session/task ending. The same `storytelling-postfix-1` pass also made **`DEMO_STORIES` the
  default** (a bare `/demo-up N` now seeds the multi-org Stories & Heroes world + serves the cockpit; `DEMO_NO_STORIES=1`
  restores the legacy small-200 structural demo), added the **per-stack Directus boot health-gate** (the bring-up tail
  waits for the stack's own offset `/server/health` before returning, so autoverify can't race the ~30s re-introspect),
  and **guarded the prod-Directus content note** (it now prints only on the genuine `DEMO_NO_LOCAL_CONTENT=1` opt-out).
  **ant-academy demo liveness — fully RESOLVED at rext tag `storytelling-postfix-2`:** there is *no token residual*.
  ant-academy *uses* Font Awesome Pro icons, but the FA Pro assets are **self-hosted / vendored** in the repo
  (`code/public/assets/fontawesome/webfonts/*.woff2` + `css/all.min.css`, rendered as `<i class="fa-solid …">`) — they
  are **not** pulled from the Font Awesome npm registry, so `npm install` (and running the app) needs **no** token. The
  `FONTAWESOME_NPM_AUTH_TOKEN` in `code/.env.example` is **vestigial**. The real "academy down in the demo" cause was a
  **blocked clone**: an empty `stack-demo/ant-academy/` stub (holding only a gitignored `code/.env.local`) tripped
  `make init`'s skip-if-present, so the source never landed. Fixed at `storytelling-postfix-2` — `ensure-clones.sh` now
  sweeps incomplete sibling stubs (any `repos.yml` repo dir with no `.git`) before `make init`, and `ant-academy.sh`
  **auto-runs** a token-less `npm install` when `node_modules` is absent. A fresh `/demo-up` now brings ant-academy up
  automatically (**proven live on :33077**).

**Dropped from tracking (2026-06-11, user instruction — re-proposal requires a fresh `/developer-kit:design-roadmap` run):**
the former v1.4 seeds **AI-generated content**, **external stack shareability** (Tailscale/ingress), and **more
mirror engines**; the **deployment/injection CI gate** (a local-only alignment surface; gates nothing in the
demo/dev workflow); and the **`/dev-up` frontend-image pre-warm** question (a UX nicety with no owner).

## Codename notes
- _(v1.0 "body double" + v1.1 "show floor" + v1.2 "set dressing" + v1.3 "stack party" + v1.3b "dress rehearsal" + v1.5 "prop room" + v1.6 "stage door" + v1.7 "house lights" + v1.8 "understudy" shipped — their codenames are now permanent. **v1.8 "understudy"** continued the theatre lineage: an understudy is a fully self-contained substitute, ready to perform on its own without the lead — exactly the self-contained-demo thesis (`stack-demo/` becomes able to run with no `stack-dev/`). Chosen at the 2026-06-15 `/developer-kit:design-roadmap` run.)_
- **v1.9 "storytelling"** (shipped 2026-06-23, tag `v1.9` — codename now permanent) continues the theatre lineage and names the thesis directly: the release is about making the seeded world **tell a story** — declarative *stories*, each with a cast of *heroes* whose verified-skill histories the product surfaces narrate. Chosen by the user at the 2026-06-22 `/developer-kit:design-roadmap` run (over the proposed "method acting" / "dramatis personae").
- **v1.10 "method acting"** (shipped 2026-06-27, tag `v1.10` — codename now permanent; chosen 2026-06-24, the runner-up codename from the v1.9 round, now apt): continues the theatre lineage and names the thesis directly — *method acting* is the deep, immersive work that makes a single **character** believable up close, exactly v1.10's job (the hero you log in as must read as a real, fully-fleshed person on every page). Alternatives weighed: "in character", "close-up". **The last release of the v1.x major.**
- **v2.0 "opening night"** (IN DEVELOPMENT — chosen 2026-06-28, the **new-major** codename): the theatre lineage reaches its culmination — *opening night* is when the production is **proven before a live audience**, the moment the whole show must **actually work** end-to-end. Exactly v2.0's thesis: the **Playthroughs** pillar plays the platform's core user journeys through, start to finish, as a real person would, and proves they work. A fitting opener for the new major.
- **v3.0 "open house"** (PROPOSAL 2026-07-01, awaiting review, **new-major** codename): continues the theatre lineage past opening night — an *open house* is when the venue throws its doors open to guests who weren't in the original audience. Exactly v3.0's thesis: the **Customer API + MCP** pillar opens the platform to programmatic consumers (customer scripts + AI agents), not just human users at the UI. Chosen at the 2026-07-01 `/developer-kit:design-roadmap` run. Alternatives considered: "curtain call", "matinee", "backstage pass".

_Last updated: 2026-07-01 (**v3.0 "open house" PROPOSED** — a NEW-MAJOR proposal opening the **Customer API + MCP**
pillar [programmatic access], **4 milestones M301 → M302 → M303 → M304** [`Mxyy` numbering; M301 `section`
identity seam ∥ M302 `section` access primitive → M303 `iterative` REST reads gateway → M304 `section` customer
surface + docs-lite; strictly sequential], governed by the consolidated pillar spec
`spec-drafts/customer-api-mcp/spec.md` v0.1 + its companions. Milestone dirs scaffolded under
`releases/03.00-open-house/`. R1 is **READ-ONLY**; W1/W2 writes staged for R2/R3. v2.0 opening night remains in
development — this v3.0 proposal is orthogonal and awaits review. Prior: 2026-06-28 (**v2.0 "opening night" PROMOTED to active development** — a NEW MAJOR opening the
**Playthroughs** pillar [functional-flow e2e *testing*], **4 milestones M201 ∥ M202 → { M203 ∥ M204 }** [`Mxyy`
numbering; M201 `iterative`+user-guided manifest corpus ∥ M202 `section` foundation → M203/M204 `iterative`
per-vantage coverage], branch `release/02.00-opening-night` cut from `main`, designed from
`spec-drafts/playthroughs/spec.md` v0.3. The Playthroughs capability **graduated from spec-draft to active
development**. Future v2 milestones [M205 Hiring/tier-gates, M206 AI-sim-mirror-tier, M207 Academy] added — NOT
pre-assigned to minor versions. v1.x history rotated to `roadmap-legacy.md`. Standing backlog
[DEF-M10-01, DEF-M21-01, M25-D9] is orthogonal + intact. Prior: 2026-06-27 **v1.10 "method acting" SHIPPED**
[tag `v1.10`] — the last release of the v1.x major. Prior: 2026-06-24 **v1.10 PROMOTED to active development** — the
believable-profile release, M39→M42m extended to M46, designed from `.agentspace/profile_gaps.md`. Prior same day:
**M33 ant-academy demo liveness FULLY RESOLVED post-v1.9** at rext tags `storytelling-postfix-1`
+ `storytelling-postfix-2`. postfix-1: the session-reaping was reproduced + fixed via session-detach [`launch_detached`];
the same pass also made `DEMO_STORIES` the default, added the Directus boot health-gate, and guarded the prod-Directus
note. postfix-2: corrected the academy residual — there is **no FA-token prerequisite** (FA Pro icons are vendored, so
`npm install` needs no token); the real cause was a **blocked clone** (empty `stack-demo/ant-academy/` stub), fixed by
`ensure-clones.sh`'s incomplete-stub sweep + `ant-academy.sh`'s token-less auto-`npm install` — a fresh `/demo-up` now
brings ant-academy up automatically (proven live on :33077). Moved M33 out of backlog → resolved. Backlog now: DEF-M10-01, DEF-M21-01, M25-D9.
Prior: 2026-06-23 **v1.9 "storytelling" SHIPPED** [tag `v1.9`] via `/developer-kit:close-release` — reviewed
M34→M38 as one PR, GREEN/0 blocking, deferral re-audit GREEN [0 escape-hatch], merged `release/01.90-storytelling`
→ `main`; 2026-06-22 v1.9 DESIGNED + PROMOTED [5 `section` milestones M34→M38]; 2026-06-15 v1.8 "understudy"
SHIPPED [tag `v1.8`].)_
