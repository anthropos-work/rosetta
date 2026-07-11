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
> **v1.10b "fit-up"** → 2026-06-29 (IN DEVELOPMENT, branch `release/01.10b-fit-up`; tag `v1.10.1`; an **interposed
> field-hardening backfill** in the v1.3b "dress rehearsal" lineage — **v2.0 "opening night" is PAUSED** until it
> ships). A from-scratch `/demo-up` surfaced 8 bring-up issues + a tail of v1.10 content gaps, and the M201 close
> *reported* the `stack-demo` clones ~5 weeks / 115+ commits behind prod (M47 later found the clones **current** —
> the **corpus** is the stale surface, → M48). v1.10b recaptures the snapshot, re-grounds the corpus, fixes the
> bring-up + content issues, adds a curated **AI-readiness showcase org** (redeeming the M201
> member-AI-readiness false-negative), and consolidates **one auditable seed+gen manifest**. **7 milestones
> M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53** (the v1.x flat counter re-opened — backfill work is a `.1` patch of
> v1.10, not a v2 `Mxyy` milestone). Designed from `.agentspace/annotation.md` + the M201 stale-clone finding.
> Tooling + docs only — zero platform-repo edits.
> **v2.1 "quick change"** → 2026-07-08 (SHIPPED 2026-07-09, tag `v2.1`, branch `release/02.10-quick-change` merged →
> `main`; the
> **skiller-in-app re-ground** — a field-hardening release in the v1.3b "dress rehearsal" / v1.10b "fit-up" lineage,
> triggered by a **landed platform structural change**: the `skiller` service + its DB schema merged into
> `app`/`public` [table names unchanged, `skiller.X → public.X`], RPC → `backend`, the skiller GraphQL subgraph gone
> [**4 subgraphs**], skiller repo/container removed. Re-fits the rext tooling + corpus + stacks to the merged
> platform and **proves `dev-up` + `demo-up` still work**. **4 milestones M208 → M209 → M210 → M211**, strictly
> sequential; tooling + docs only, zero platform edits). Designed from the user's skiller-merge briefing + the
> colleague's `origin/docs/skiller-in-app-merge` corpus sweep (correct-but-incomplete). Takes **M208+** — the
> reserved Playthroughs futures **M205–M207 stay in vision** (M206 is a live Fate-3 destination from the M203 close),
> per the established "reserved-number-ships-later" precedent (M26).
> **v2.2 "panorama"** → 2026-07-11 (SHIPPED 2026-07-12, tag `v2.2`, branch `release/02.20-panorama` merged → `main`;
> rext code-of-record `v2.2` = `39e8013`); the **external-shareability
> release** — make dev/demo stacks reachable from other machines on a **Tailscale** tailnet (run a stack on a
> Tailscale VM, e.g. `billion.taildc510.ts.net` on the odyssey Proxmox host; a teammate with Tailscale up browses
> the demo end-to-end). **4 milestones M212 → { M213 ∥ M214 } → M215** (+ optional M216); opt-in default-off,
> HTTPS-everywhere under one MagicDNS origin, demo-first; **tooling + docs + an opt-in flag only — zero
> platform-repo edits** (a 2-item patch tail via the existing rext sha-pinned mechanism). The **sanctioned
> re-proposal** of the v1.4 seed "external stack shareability (Tailscale/ingress)" dropped 2026-06-11 — takes
> **M212+**, the next free band after v2.1's M211 (reserved Playthroughs futures M205–M207 untouched in vision).
> Designed from the user's Tailscale-serve briefing + the 5-agent feasibility workflow `wf_bea3be47` (config-only
> core confirmed).

---

> **v2.0 "opening night" SHIPPED 2026-07-02** (tag `v2.0`) + **v2.1 "quick change" SHIPPED 2026-07-09** (tag `v2.1`) +
> **v2.2 "panorama" SHIPPED 2026-07-12** (tag `v2.2`).
> **v2.2 "panorama" is now IN DEVELOPMENT** (branch `release/02.20-panorama`, designed 2026-07-11 — the external-shareability / Tailscale-serve release). (v2.0/v2.1 detail is in the
> `## Done` sections of [`roadmap.md`](roadmap.md) — the active roadmap holds the v2.x major; v1.x history is in
> [`roadmap-legacy.md`](roadmap-legacy.md).) v2.1 was a **field-hardening re-ground** (the skiller-in-app merge),
> **not** a Playthroughs release — it took **M208+**, so the reserved Playthroughs futures **M205–M207 below stay in
> vision** for the next release to design. The unscheduled backlog below is likewise unscheduled.

## Future v2 milestones (Playthroughs pillar — NOT yet clustered into a minor version)

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
  these legs are parked as `later — needs a mirror engine`. **Also absorbs the non-gate employee-vantage
  coverage-DEEPENING legs the M203 gate legitimately left additional** (the M203 gate enumerated + proved the 3
  core employee journeys — Profile / Skill Paths / AI-sims chat-launch — GREEN; these are extra edge UCs beyond
  it, routed at M203 close 2026-07-02 via Fate-3): **`ai-simulations.code.UC1`** (Judge0-via-Roadrunner code sim
  — needs the external hardcoded Judge0 host as a live seed/stack precondition), **`ai-simulations.interview.UC1`**
  (text/non-voice interview — reuses the chat engine, needs an interview-typed catalog sim), **the Skill-Paths
  verify-skill end-to-end TERMINAL** (compose browse→learn→complete→verify with a NON-voice ASSESSMENT sim; the
  verify OUTCOME is already proven on the profile side by `pt-profile-verified`), and **`profile.self-evaluation.UC1`**
  (the Profile skill self-rate WRITE — a rate-modal click-intercept quirk that needs live browser iteration to
  stabilize). All four need a live demo + a browser drive to land (out of the docs-only M203 close's scope); the
  code-sim + interview additionally carry seed/stack preconditions. Provenance: M203 iter-05/iter-06 "routes carried
  forward" + `m203-employee-coverage/decisions.md` D-CLOSE-1.
- **M207 — Academy coverage** — Playthroughs over the **separate ant-academy deployment** (its own Vercel-deployed
  app, Clerk-only, not in the platform docker-compose). A distinct target environment from the demo-N hero stacks;
  a future surface for the capability.

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
- **M25-D9 — dev cold DB-init (extensions-schema bootstrap + PG-readiness) → RESOLVED (v2.1 M211).** The
  extensions-bootstrap-before-migrate class (a cold `make reset-db`/`make migrate` failing `schema "extensions"
  does not exist` because the un-editable platform `make migrate` doesn't create `extensions`) is now codified in
  **`dev-stack/migrate-dev.sh`** (rext `quick-change-m211`) — `wait_pg` → create schemas
  (`extensions`/`sentinel`/`cms`/`jobsimulation`/`skillpath`) + `CREATE EXTENSION vector/pgcrypto/pg_trgm SCHEMA
  extensions` → atlas-migrate the 4 merged services → load the casbin policy, a mirror of `demo-stack/migrate-demo.sh`
  for the main dev stack. Cold-verified on a faithful non-destructive throwaway (extensions + `gin_trgm_ops` + 89
  public tables + `cms.vector` + casbin, 0 skiller). Documented in `corpus/ops/setup_guide.md` + `.claude/skills/dev-up/SKILL.md`.
  _(Residual nuance, narrow: the original `dev-N≥1 --local-content` taxonomy-replay `rc=4` "target schema empty"
  symptom — orthogonal to the content-serve done-bar DB-2, which is GREEN — consumes the same hook; no separate
  work owed.)_
- **Clean-box literal full destructive `/dev-up` (v2.1 M211 belt-and-suspenders).** M211's gate proved the dev
  half via the M25-D9 cold DB-init cold-verified on a faithful throwaway + a live docker harness — a literal full
  all-services destructive `/dev-up` + verify-net was **deliberately not run** because this box is committed to the
  user's native-app content-line dev (`docker-compose.override.yml` → `backend:host-gateway` + an
  `app-01.10-content-line` worktree), which a full bring-up would clobber (and it can't go green without a v2.1
  native backend). An environment-respecting gate interpretation, not a gap. A clean-box full `/dev-up` remains a
  nice belt-and-suspenders confirmation on a box not committed to unrelated native-app work; unscheduled.
- **M314b — prod frozen-read whole-org AI-readiness hydration (a prod-team / PLATFORM follow-up, out of rosetta
  tooling scope).** Inherited from the v1.10b M51 AI-readiness showcase-org work: the prod AI-readiness snapshot
  read-path loads members whole-org from a frozen read, a hydration surface that belongs to the **platform** (which
  rosetta's read-only line does not edit) — so there is **no rosetta tooling work owed**. Context lives in
  `corpus/ops/demo/coverage-protocol.md` + `corpus/services/ai-readiness.md` + the v1.10b/v2.0 retros; tracked here
  as a standing prod-team backlog pointer so the state.md/context.md cross-reference resolves.
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
the former v1.4 seeds **AI-generated content** and **more mirror engines**; the **deployment/injection CI gate** (a
local-only alignment surface; gates nothing in the demo/dev workflow); and the **`/dev-up` frontend-image pre-warm**
question (a UX nicety with no owner).
> **RE-PROPOSED 2026-07-11 — `external stack shareability` (Tailscale/ingress):** brought back via the sanctioned
> path (a fresh `/developer-kit:design-roadmap` run) and **promoted to v2.2 "panorama"** (IN DEVELOPMENT, branch
> `release/02.20-panorama`). No longer dropped. The other three items above remain dropped (re-proposal requires
> their own fresh design-roadmap run).

## Codename notes
- _(v1.0 "body double" + v1.1 "show floor" + v1.2 "set dressing" + v1.3 "stack party" + v1.3b "dress rehearsal" + v1.5 "prop room" + v1.6 "stage door" + v1.7 "house lights" + v1.8 "understudy" shipped — their codenames are now permanent. **v1.8 "understudy"** continued the theatre lineage: an understudy is a fully self-contained substitute, ready to perform on its own without the lead — exactly the self-contained-demo thesis (`stack-demo/` becomes able to run with no `stack-dev/`). Chosen at the 2026-06-15 `/developer-kit:design-roadmap` run.)_
- **v1.9 "storytelling"** (shipped 2026-06-23, tag `v1.9` — codename now permanent) continues the theatre lineage and names the thesis directly: the release is about making the seeded world **tell a story** — declarative *stories*, each with a cast of *heroes* whose verified-skill histories the product surfaces narrate. Chosen by the user at the 2026-06-22 `/developer-kit:design-roadmap` run (over the proposed "method acting" / "dramatis personae").
- **v1.10 "method acting"** (shipped 2026-06-27, tag `v1.10` — codename now permanent; chosen 2026-06-24, the runner-up codename from the v1.9 round, now apt): continues the theatre lineage and names the thesis directly — *method acting* is the deep, immersive work that makes a single **character** believable up close, exactly v1.10's job (the hero you log in as must read as a real, fully-fleshed person on every page). Alternatives weighed: "in character", "close-up". **The last release of the v1.x major.**
- **v2.0 "opening night"** (IN DEVELOPMENT — chosen 2026-06-28, the **new-major** codename): the theatre lineage reaches its culmination — *opening night* is when the production is **proven before a live audience**, the moment the whole show must **actually work** end-to-end. Exactly v2.0's thesis: the **Playthroughs** pillar plays the platform's core user journeys through, start to finish, as a real person would, and proves they work. A fitting opener for the new major.
- **v1.10b "fit-up"** (IN DEVELOPMENT — chosen 2026-06-29 by the user, theatre lineage): the *fit-up* (a.k.a. the get-in) is the technical work of **building and rigging the set correctly in the venue** before the show can run — the crew assembles the world so it holds together under the lights. Exactly this backfill's job: re-ground the demo to current prod and fix the bring-up so the environment stands up cleanly from cold — the technical preparation that must happen **before** v2.0 "opening night" can resume. Sits in the same field-hardening lineage as **v1.3b "dress rehearsal"** (the prior demo-up-issue backfill). Alternatives weighed: "tech rehearsal", "house notes".
- **v2.2 "panorama"** (IN DEVELOPMENT — chosen 2026-07-11 by the user, over the proposed "on tour" / "road show" / "the transfer" / "guest house"): a *panorama* was a 19th-C immersive spectacle attraction — the whole scene, taken in from any vantage. Names the thesis directly: the whole environment made **viewable from anywhere on the tailnet**, no longer a single-seat show on the host's own `localhost`. Fits the spectacle/entertainment lineage even as it steps outside the strict backstage-of-a-play metaphor.

_Last updated: 2026-07-11 (**v2.2 "panorama" DESIGNED + PROMOTED to active development** — the **external-shareability
/ Tailscale-serve release**: make dev/demo stacks reachable from other machines on a Tailscale tailnet; **4 milestones
M212 → { M213 ∥ M214 } → M215** (+ opt M216), branch `release/02.20-panorama` cut from `main`, tag `v2.2`; opt-in
default-off, HTTPS-everywhere, demo-first; tooling + docs + an opt-in flag only, the 2-item patch tail via the rext
mechanism. The **sanctioned re-proposal** of the dropped v1.4 Tailscale/ingress seed; designed from the user's
briefing + the 5-agent feasibility workflow `wf_bea3be47`. Takes M212+; reserved Playthroughs futures M205–M207 stay
in vision. Prior: 2026-07-08 (**v2.1 "quick change" DESIGNED + PROMOTED to active development** — the **skiller-in-app
re-ground**, a field-hardening release [v1.3b "dress rehearsal" / v1.10b "fit-up" lineage] triggered by the landed
platform merge of `skiller` into `app`/`public`; **4 milestones M208 → M209 → M210 → M211**, strictly sequential;
branch `release/02.10-quick-change`, tag `v2.1`. Takes M208+; the reserved Playthroughs futures M205–M207 stay in
vision. Prior: 2026-07-02 **v2.0 "opening night" SHIPPED** [tag `v2.0`]. Prior: 2026-06-29 (**v1.10b "fit-up"
PROMOTED to active development** — the interposed **field-hardening backfill** [v1.3b "dress rehearsal" lineage]; **7 milestones M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53**
[v1.x flat counter re-opened], branch `release/01.10b-fit-up` cut from `main`, tag `v1.10.1`; designed from
`.agentspace/annotation.md` + the M201 stale-clone finding. Re-grounds demo + corpus to current prod, fixes the 8
demo-up issues + the v1.10 content gaps, adds the AI-readiness showcase org [M51], consolidates one inlined
seed+gen manifest [M52]. **v2.0 "opening night" PAUSED** until it ships. Codename "fit-up" chosen by the user.
Prior: 2026-06-28 (**v2.0 "opening night" PROMOTED to active development** — a NEW MAJOR opening the
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
