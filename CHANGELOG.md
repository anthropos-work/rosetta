# Changelog

All notable user-facing changes to Project Rosetta. Format: [Keep a Changelog](https://keepachangelog.com/), semver-aware.

## [v2.2] "panorama" — 2026-07-12

**External shareability over Tailscale.** Make a dev/demo stack reachable from **another machine on a Tailscale tailnet** — run a demo on a Tailscale VM (e.g. `billion.taildc510.ts.net` on the odyssey Proxmox host) and a teammate with Tailscale up browses it end-to-end over **one trusted HTTPS origin**. External access is **opt-in, default-off** (an explicit `/demo-up N --public-host <magicdns>` flag); a bare `/demo-up N` is byte-identical to before. This release ships the whole surface as **tooling + docs + one opt-in flag — zero platform-repo edits, 0 net-new deps** — and is the **first live remote Linux-VM deploy**, proven end-to-end. (rext code-of-record @ tag `v2.2` = `39e8013`.)

### Added
- **`/demo-up N --public-host <magicdns>`** (M212) — the opt-in host knob (`STACK_PUBLIC_HOST`) that makes a demo bind + advertise a public MagicDNS origin instead of `localhost`. Default-off, unset ⇒ byte-identical to today.
- **HTTPS-everywhere over the tailnet** (M213) — one trusted **`tailscale cert`** origin (real Let's Encrypt, no CA install) fronted by a per-offset-port `tailscale serve` reverse proxy (`gen_tailscale_serve.py`), so Clerk gets the secure context it requires. Dotted-publishable-key host validation; host-agnostic token verify; path-only cert mount.
- **`corpus/ops/demo/tailscale-serve.md`** (M214) — the full **remote Linux-VM deploy runbook**: Step-0 host prereqs (Go + atlas + tailscale operator) with install commands → PAT clone → workspace → secrets → snapshot cache → `--public-host` bring-up → `tailscale serve` → verify (exact curls + cockpit login, both vantages) → teardown. Plus the F1–F12 host-deploy finding set + safety framing. Indexed from `CLAUDE.md` + `corpus/ops/README.md`.
- **Host pre-flight + auto-handling for a fresh Linux VM** (M215, in rext) — `preflight_host_prereqs` (Go/atlas/tailscale-operator → fail loud with exact fix lines; `DEMO_NO_HOST_PREFLIGHT=1` opt-out); a keyless ssh-agent auto-start for buildx; Linux bind-mount data-dir pre-creation; `migrate-demo.sh` fails loud on a genuine atlas failure.

### Changed
- **CORS + origins** (M214) — `CORS_EXTRA_ORIGINS` emits the https origin trio; the ant-academy `allowedDevOrigins` rides the existing rext **sha-pinned patch mechanism** (drift-refuse fails loud on an upstream change; reverted on teardown) — never a canonical platform-repo edit.
- **Teardown resets `tailscale serve`** (M215) — `/demo-down` clears THIS demo's serve ports (offset-scoped per-port `off`) + a defensive up-path pre-reset, so a re-deploy no longer port-conflicts on a leftover listener. Non-fatal, no-op on localhost.
- **rext README index reconcile** (close) — demo-stack test-count, the `gen_tailscale_serve.py` + `apply-ant-academy-dev-origins.sh` rows, and the F12 ADV-1 comment (D-CLOSE-1/-2/-3 + ADV-1).

### Fixed
- **`git tag --list | head` SIGPIPE → exit 141** aborting bring-up on a many-tag repo (`app` ~337 v-tags) — replaced with a pipe-less `git for-each-ref --count=1` (M215/F3).

### Verified
- **First live remote Linux-VM deploy** on `billion` (M215, `closed-on-gate`): a browser on a **different** tailnet machine completed a full journey for **both** hero vantages — employee `maya-thriving` → `/profile` and manager `dan-manager` → `/enterprise/workforce` — on a **genuinely trusted** cert (`ignoreHTTPSErrors:false`, `verify=0`, no CA install), 0 console errors, 0 localhost/prod ejects, assets rendering, **reproducibly on a clean cold reset-to-seed** one-shot. Unset knob byte-identical (regression-safe). Triple-clean demo-stack suite 3/3; Go `go test ./...` exit 0 all 6 modules; shellcheck clean.

### Supply chain
- **0 net-new dependencies** (a network/scheme reconfiguration adds no imports). 0 reachable vulnerabilities (govulncheck, go1.25.12); npm 0; Python 0 third-party deps; 0 GPL/AGPL. 13 pre-existing dependabot alerts (all the **unreachable** `x/crypto` ssh-subpackage in the clerkenstein Clerk mock) → cleared next rext roll with `go get x/crypto@v0.52.0`.

### Known limitations
- **Standing-backlog residuals** (documented, non-blocking, off the proven journey path): F5 two `app` demopatches sha-drift-refuse (demo works, slower per-member fan-out) → demopatch re-anchor; F9 a fresh VM has no snapshot cache (taxonomy/library sparse; identity/profile/dashboard/workforce render fully) → cache pre-stage/auto-sync; F11 a cosmetic seed hero-name mismatch; F13 the jobsimulation container exits(1) on startup (AI-Simulations surface, would hit any demo). Optional **M216** (dev-path Tailscale parity) stays roadmap-only until promoted.

## [v2.1] "quick change" — 2026-07-09

**The skiller-in-app re-ground.** The platform merged the standalone `skiller` service + its DB schema into `app` — the skills taxonomy, embeddings, and job-roles now live in the `public` schema (table names unchanged, only `skiller.X → public.X`), the RPC surface is served by `backend`, and the skiller GraphQL subgraph is gone (**4 subgraphs**). This release re-fits the environment-builder tooling, the corpus, and the local stacks to that merged platform and **proves `/dev-up` + `/demo-up` still work end-to-end, cold**. A field-hardening release (v1.3b/v1.10b lineage) — **tooling + docs + stack-re-sync only; zero platform-repo edits.** (rext code-of-record @ tag `v2.1`.)

### Changed
- **rext tooling re-grounded `skiller.*` → `public.*`** (M209): the snapshot capture/replay taxonomy surface (one const flip re-grounds both capture and replay) + the cache-key staleness digest narrowed to the surface's own tables (no post-merge cache-thrash) + the seeding taxonomy resolvers — all keeping the `organization_id IS NULL` public predicate; the `skiller` service probe dropped from verify + demo bring-up.
- **Corpus + skills re-ground** (M210): the merged-platform architecture (4 subgraphs, RPC→backend, taxonomy in `public`) documented + 6 rext-facing tooling-doc bodies flipped to `public.*` (0 stale `skiller.<table>` tooling refs corpus-wide).
- **Both local stacks re-synced** to the merged platform (M208); the vestigial `skiller/` clones removed.

### Added
- **`dev-stack/migrate-dev.sh`** — a dev cold DB-init (bootstraps the `extensions` schema + pgvector/pg_trgm/pgcrypto + casbin before migrate; mirror of the demo path), resolving the standing M25-D9 cold-bring-up gap.
- **Cache-migration recapture** — re-key a captured snapshot cache `skiller.*→public.*` when a merge preserves data + table names (no prod capture source needed); used to recapture the 42,790-skill public taxonomy + the sim-embeddings, cold.
- **A build-scratch freshness guard** — the demo bring-up now re-syncs its injected build-scratch to the current source ref on every bring-up (a pinned scratch had been silently shipping pre-merge binaries through `--purge`).

### Fixed
- The cold `/dev-up` missing casbin-policy load (the M18 silent-403 class); a reset-to-seed roster gap that failed all 10 Playthroughs; a stale-image mishmash that made "cold" demos ship pre-merge code.
- Two inherited Go-stdlib advisories cleared by `go1.25.11 → go1.25.12` (govulncheck clean on all 6 modules).

### Verified
- **M42 coverage GREEN at both vantages + the v2.0 Playthroughs 10/11 GREEN** on the merged platform (M211, gate 6/6, closed-on-gate); triple-clean 3/3; Clerkenstein alignment gates 100%/100% held.

### Supply chain
- **0 net-new dependencies** (a schema re-point adds no imports). Go toolchain `go1.25.11 → go1.25.12`. npm 0 vulnerabilities; 0 GPL/AGPL.

### Known limitations
- The dev-cold gate was proven at the **DB-init level on a non-destructive throwaway** (to protect the box's native-app content-line dev setup), not a literal full destructive `/dev-up`; a clean-box full run is tracked as belt-and-suspenders backlog (CAVEAT-1).

## [v2.0] "opening night" — 2026-07-02

**The Playthroughs pillar — the platform's core user journeys, proven to actually work.** A **Playthrough** is an automated actor that *is the user*: it logs in as a seeded hero, plays a real journey across the platform end-to-end, and proves the platform delivered the outcome. Where the v1.x coverage sweep proves **presence** (every page *shows* real content), a Playthrough proves **function** (the hero can *do* the thing) — it breaks only when a capability breaks, not when pixels shift. A new MAJOR (v2.x, `Mxyy` milestone numbering); the demo/seeding lineage carries forward as the foundation. 4 milestones M201→M204. **Tooling + docs only — zero platform-repo edits, zero net-new third-party deps.** (rext code-of-record @ tag `v2.0`)

### Added
- **The Playthroughs manifest corpus** — a prose-intent contract of the platform's products × their must-work user journeys (9 products · 26 stories · 28 use-cases), each use case carrying goal + actor + flow + expectations. The build+regression contract every coverage milestone implements against. (M201)
- **The `playthroughs` tooling pillar** (a new `rosetta-extensions` section) — a manifest model + validator (both-way id integrity + precondition coverage, gated on the seed-closure check), a per-surface page-object/locator layer (semantic-by-default; re-pinning a UI shift is O(surfaces), not O(tests)), a dedicated decoupled test world (`pt-world`, two private orgs — test data ≠ demo data), a **reset-to-seed serial runner**, and a 4-state coverage report (`passing` / `failing` / `unimplemented` / `unimplementable-without-platform-edit`). (M202)
- **10 live Playthroughs, GREEN on a cold reset-to-seed demo, deterministic over 5 consecutive runs:**
  - **Employee vantage (6):** Profile (identity + verified-skill Spotlight + claimed-vs-verified gap + growth trajectory + work/education timeline), Skill Paths (browse → open → start → progress), AI Simulations (chat launch). (M203)
  - **Manager vantage (4):** Workforce funnel + member roster, per-member activity-dashboard drill-down, succession / at-risk. (M204)
- New runbook: [`corpus/ops/demo/playthroughs.md`](corpus/ops/demo/playthroughs.md) — the functional-flow e2e protocol + the M203/M204 iteration protocol.

### Changed
- The presenter cockpit's seat-switch (M37) and the M42 e2e foundation are now **reused** (never forked) as the Playthrough login + browser-drive substrate.

### Known limitations
- **1 declared in-manifest TODO:** `assignment-monitoring.assign-and-track` UC1 — the assign-**write** half (a two-backend org-admin write flow) — is declared and surfaced as `unimplemented` (a first-class tracked state, presence-pinned so it can't silently vanish), routed to a future manager-write tier.
- **4 non-gate employee edge journeys** (Judge0 code-sim, text-interview, the skill-path verify-skill terminal, profile self-evaluation write) are routed to the future **M206** AI-sim-mirror tier.
- Voice/recording AI-simulation journeys and the separate Ant Academy deployment are out of v2.0 scope by design (future M206 / M207). The read-only-platform line carries over: an un-drivable surface **escalates**, it never edits the platform.

## [v1.10.1] "fit-up" — 2026-07-01

The **field-hardening backfill.** A from-scratch `/demo-up` surfaced bring-up issues + a tail of v1.10 content gaps; v1.10b re-grounds the demo + corpus to current prod, fixes them, adds the AI-readiness showcase org, consolidates the seed+gen intent into one auditable file, and — the proof — **destroys the demo and cold-rebuilds it from scratch**, verified end-to-end (6/6 acceptance bars + the academy). 7 milestones M47→M53. **Tooling + docs only — zero platform-repo edits.** (rext code-of-record @ tag `v1.10.1`)

### Added
- **AI-readiness showcase org.** A curated 200-person org (Northwind Aviation) with the manager AI-readiness dashboard **enabled**, ~80% of members having completed all 3 onboarding/evaluation steps, and hero personas — one who *started* the readiness onboarding and one who *completed* it. (M51)
- **One auditable seed+generation manifest.** A single checked-in file (`seed-generation-manifest.yaml`) inlining the entire seed+gen intent — all 3 orgs' population blueprint, the generation prompt templates, the batch config (cost ceiling, concurrency), and the snapshot sources — **cache + generated data excluded**. The cockpit **[Download manifest]** serves it. (M52)
- **Academy content on the demo.** Real course content + a hero academy menu-link + a non-anonymous academy session, seeded + verified on the cold build. (M53)
- New docs: `corpus/services/ai-readiness.md` (the previously-undocumented member-AI-readiness feature) + `corpus/ops/demo/seed-manifest-spec.md`.

### Changed
- **The demo now cold-rebuilds from scratch** and is verified end-to-end as the single acceptance proof — a `stack-demo`-only box brings the whole stack up with one `/demo-up`, no manual steps. (M53)
- Snapshot recaptured from current prod; the corpus re-grounded to match (the shipped AI-readiness feature was undocumented). (M47/M48)

### Fixed
- The 8 from-scratch `/demo-up` bring-up issues (rext.tag source-of-truth, env-guard order, an auto-generated invitation secret, ant-academy explicit clone, disk pre-flight + `demo-down` image purge, non-fatal frontend, demopatch re-anchor) + v1.10 content/seeding gaps (member languages, certificates, member-field backfill). (M49/M50)
- **The AI-readiness dashboard now renders at 200-member scale** — an app read-path demo-patch bounds the unbounded member hydration (180s timeout → 19ms), a data-identical perf swap. (M51)
- A manager coverage-gate assertion that M51 made unconditional (breaking the base-Workforce org's gate) is now org-conditional — caught from cold by the acceptance gate and fixed at the gate. (M53)

### Supply chain
- Cleared the inherited HIGH **CVE-2026-39821** by bumping `golang.org/x/net` v0.53.0 → v0.55.0 in `stack-seeding` (the CVE was disclosed after the v1.10 close; called only in one low-blast-radius path).

### Known limitations
- Tooling + docs only — zero platform-repo edits. The AI-readiness dashboard's org-scale perf fix is a **demo-injection patch**; the production read-path (`loadMembers`) remains a documented prod-team follow-up. The academy's AI chat stays absent-by-design (no keys provisioned in the demo).

## [v1.10] "method acting" — 2026-06-27

The **believable-profile release + the presenter-grade / scalable-generation extension.** v1.9 told the *story*; v1.10 makes each *character* hold up under a close-up. When a presenter clicks **Login as** a hero, that hero reads as a fully fleshed, believable person — the individual's **profile** (org name, role + title, work history, education, a real face, deep role-aligned skills) **and** the content surfaces they land on (**library** + the **activity feed**) populate with real semantic content, on **every** page a hero of that vantage can reach — proven by a **Playwright semantic coverage sweep** with **zero** empty pages and **zero** out-of-demo escapes, at both the employee and the manager vantage. Then it scales: a presenter-grade cockpit, a whole roster baked to depth, a cheap-LLM generation engine, and a whole **~500/735-member org filled from one descriptor**. **Tooling + docs only — zero platform-repo edits.** (rext code-of-record @ tags `method-acting-m39`..`m46-servegrant-closure`)

### Added
- **Believable hero profiles.** A logged-in hero now shows the right **org name**, a real **role + title**, a real **face**, and a deep **profile** — work history + education timeline + a claimed-but-unverified skill tail that widens the visible claimed-vs-verified gap. (M39 identity + M41 depth)
- **The content surfaces stop emptying.** The per-stack Directus **serve-grant** lets the library **and** the activity feed render real content (the anonymous public-policy deep-fetch no longer panics cms) — the two surfaces that read empty in the live review now populate. (M40)
- **A per-vantage demo-coverage gate.** A **Playwright** sweep crawls every page a hero of a vantage can reach and asserts a **semantic believability gate** (real seeded content + per-section cardinality + persona self-consistency [role↔skills, menu==profile real-photo avatar, org name + logo] + 0 prod-eject escapes). Both vantages pass cold (employee + manager). (M42e + M42m)
- **A slicker presenter cockpit.** The login launcher goes **light** + professional: a card per hero, one unified **[Log in as]** CTA (logs in *and* lands on her per-role screen — no separate Jump), a seed-manifest download, and a staged login-progress overlay. (M43)
- **The whole roster baked to depth.** Not just the heroes — **every** seeded member and manager gets trajectory-aware self-ratings, certificates + projects, an avatar, and a career, so an org grid is believable end-to-end (DATA DENSITY). (M44)
- **A cheap-LLM generation engine.** `cmd/gen-batch` turns a YAML **batch descriptor** into realistic per-member profiles via a cheap model (gpt-4o-mini), with a **prompt-keyed cache** (an unchanged descriptor re-seeds byte-identical at **$0**), a mandatory `--max-cost` ceiling, and the **CODE-owns-structure / AI-owns-content** boundary (every generated role/skill routes through the real resolvers; non-resolving names **drop**, never fabricated). (M45)
- **Org-scale fill + a preview CLI.** One supporting-population descriptor fills a whole believable **~500/735-member org** (per-story, deterministic auto-fill); a `gen-batch --preview` renders the expanded plan + an estimated-cost line without seeding. (M46)
- **Five new corpus specs** — `cockpit-spec.md` (M43), `coverage-protocol.md` (M42), `profile-completeness-spec.md` (M44), `ai-generation-spec.md` + `cache-spec.md` (M45) — plus `snapshot-spec.md` / `stories-spec.md` / `seeding-spec.md` / `frontend-tier.md` / `demo/README.md` updates.

### Supply chain
- **One new dependency — deliberate + user-acknowledged.** M45 adds `github.com/anthropos-work/ai@v1.40.1` (the `services/ai` wrapper transport, EU-first Azure; transitive Azure SDK + openai-go/v3, all MIT/BSD/Apache) — the in-release inflection the generation-engine milestone is *about* (v1.8→v1.9 was 0-new-deps). M46 reuses it unchanged. License-vetted, no GPL/AGPL. Lockfile: `knowledge/plan/releases/archive/01.10-method-acting/dependencies.lock`.

### Known limitations
- The **manager enterprise grids** at org scale render through a few **`demopatch`** layers applied to the demo's ephemeral clone before build (a next-web pagination patch + 2 FK indexes + a `roles.go` read-gate drop + a Directus column-drift backfill + the serve-grant closure) — **zero canonical platform edits**, every fix reproducible on a fresh `/demo-up`, but they are demo-time patches, not platform changes.
- **A regenerated snapshot cache needs `--purge`.** A plain `down`/`up` keeps the demo's Postgres volume and no-ops the structure replay; the serve-grant recapture only lands on a fresh `--purge /demo-up`.
- **The org-scale fill is opt-in.** A bare `/demo-up` seeds the default Stories trio (~341); the ~500/735-member generated org is a separate `/stack-seed --gen-batches` (a $0 cache-hit reseed once the batch is cached).
- The `clerk-express-1` alignment gate drives the genuine `@clerk/express` SDK and needs installed npm modules (a node-CI env prerequisite) — unrunnable in the authoring copy; not a regression.

## [post-v1.9] demo-hardening — `storytelling-postfix-1`, `storytelling-postfix-2` — 2026-06-23

The **post-v1.9 demo-hardening patch**: the Stories & Heroes world becomes the **default** demo, and the M33-class "dead on a later visit" failures are run down and fixed. A bare `/demo-up` now seeds the narrative + serves the cockpit out of the box, the host-native daemons survive their launching session, ant-academy clones + installs itself so it comes up automatically, and two transient false-down warnings on a healthy stack are closed. **Tooling + docs only — zero platform-repo edits.** (ext tags `storytelling-postfix-1`, `storytelling-postfix-2`)

### Changed
- **`DEMO_STORIES` is now default-ON.** A bare `/demo-up N` now seeds the multi-org **Stories & Heroes** world (2 orgs × a thriving/struggling/manager hero trio) **and** serves the **presenter cockpit** by default — the M38 opt-in became opt-out. `DEMO_NO_STORIES=1` (or the explicit `DEMO_STORIES=0`) restores the legacy structural **small-200** single-identity demo (no cockpit), mirroring the `DEMO_NO_*` family.

### Fixed
- **M33 ant-academy + presenter cockpit "dead on a later visit" is resolved.** Both host-native daemons were launched via `nohup` alone, which does **not** detach from the launcher's process group — so when a backgrounded `/demo-up` task's process tree was reaped (or its launching session ended), the daemon died with it. They now launch **session-detached** via a shared `demo-stack/detach.sh::launch_detached` (`setsid` where present; a portable `python3 os.setsid` double-fork on macOS, which has no `setsid`), so they survive the launching session/task ending.
- **ant-academy "down in the demo" is resolved — and the Font Awesome token was a red herring.** The real cause was a **blocked clone**: an empty `stack-demo/ant-academy/` stub (holding only a gitignored `code/.env.local`) tripped `make init`'s skip-if-present, so the ant-academy source never landed. Fixed two ways: `ensure-clones.sh` now **sweeps incomplete sibling stubs** (any `repos.yml` repo dir with no `.git`) before `make init`, so the clone actually happens; and `ant-academy.sh` now **auto-installs deps** (`npm install`, **no token**) when `node_modules` is absent. A fresh `/demo-up` brings ant-academy up automatically (proven live on `:33077`). The `FONTAWESOME_NPM_AUTH_TOKEN` in `code/.env.example` is **vestigial** — ant-academy *does* use Font Awesome Pro icons, but the assets are **self-hosted/vendored** in the repo (`code/public/assets/fontawesome/…`), so a token-less `npm install` succeeds and the icons render; it is **never** a prerequisite to install or run, nor a reason ant-academy is skipped. (ext tag `storytelling-postfix-2`)
- **The per-stack Directus boot now health-gates before autoverify.** `boot_directus_step` used to `docker restart` the Directus container and return immediately, so the bring-up-tail autoverify raced Directus's ~30s re-introspect and **false-reported it "down"** on a stack that was actually fine. It now waits (bounded by `DEV_SETDRESS_DIRECTUS_BOOT_TIMEOUT`, default 90s; non-fatal on timeout/curl-absent) for the stack's own offset `/server/health` to answer `200` before returning.
- **The prod-Directus content note is now guarded.** `up-injected.sh` used to print "reads PUBLIC content LIVE from prod Directus (no per-stack Directus yet)" **unconditionally**, but since v1.5 "prop room" the default demo boots a per-stack Directus serving the captured catalog locally. The note now prints **only** on `DEMO_NO_LOCAL_CONTENT=1` (the genuine prod-read opt-out), reworded to that reality; the default local-served case says nothing.

## [v1.9] "storytelling" — 2026-06-23

The **believable-demo-narrative release**: the placeholder seeder becomes a declarative **Stories & Heroes** engine. Each *story* is one org with a thriving/struggling/manager **hero** trio, seeded via the real **verified-skill chain**, so the individual **skill profile** and the org **Workforce dashboard** tell one coherent story (the claimed-vs-verified gap is the "aha"). A standalone **presenter cockpit** lets a demo-giver *log in as* a hero and *jump to* the right screen, wired on **Clerkenstein multi-identity**. **Tooling + docs only — zero platform-repo edits.**

### Added
- **Stories & Heroes seeding model.** One `stack.stories.yaml` seeds **multiple orgs**, each with a thriving/struggling/manager **hero trio** at vantage-appropriate fidelity (thriving = dense verified, under-claim; struggling = sparse, over-claim; manager = rides aggregates). Drives `/stack-seed`. The legacy single-org preset/`dev-min` path stays byte-identical. (M35)
- **The verified-skill chain.** The seeder writes the real **7-table fan-out** per (hero × skill) (`jobsimulation.sessions` → validation results → `public.user_skills` → `user_skill_evidences`), drawing real replayed skiller node-ids (role-coherent; **never fabricated** — empty-pool skips), so a hero's **skill profile + Skill Spotlight chart** render with real, closure-verified data (0 dangling node-ids). (M34)
- **Org Workforce-Intelligence dashboard data.** Six seeders make every dashboard aggregate render believably for a seeded story: the mapped→verified verification funnel, business-unit teams + a mentor tag, role-gap + internal-mobility target-roles, succession interview coverage, ~2:1 positive feedback, and the org-scale claimed-vs-verified gap + AI-readiness. (M36)
- **Clerkenstein multi-identity (seat-switch).** A demo can present as any seeded hero — a users/orgs Registry fed a roster JSON + a server-authoritative FAPI handshake selection (`?__clerk_identity=<key>`). A **5th Alignment DNA** (`clerk-multi-1`, 9 genes) measures the multi-session FAPI surface, scoring 100%/100% alongside the existing four. (M37)
- **The presenter cockpit.** A standalone served panel (offset port, gated on `DEMO_STORIES=1`) that reads the same `stack.stories.yaml` that seeded the data and lists each story → its hero trio with **[Login as]** + **[Jump to section]** — the two collapse into one FAPI handshake redirect, so a demo-giver picks a hero and lands logged-in on her screen in one move. Ships the roster-export + cockpit-manifest producers + the deep-link catalog. (M38)
- **`corpus/ops/demo/stories-spec.md`** (net-new) — the Stories & Heroes + verified-skill-chain + dashboard + cockpit reference; plus updates to `seeding-spec.md`, `safety.md`, `demo/README.md`, `clerkenstein.md`, `alignment_testing.md`, `rosetta_demo.md`, and the `/stack-seed` skill.

### Changed
- **A demo seat's `org_role` is now vantage-faithful.** A hero's JWT `org_role` follows her vantage (end-user → `member`, manager → `admin`), single-sourced so the membership row + the casbin grant + the roster/JWT claim agree per hero — an "employee" demo seat reads as `member`, not org-admin. (M38, close-fix M38-D8)
- The `clerkenstein.md` / `alignment_testing.md` narrative now reflects **five** measured Clerkenstein surfaces (was four).

### Supply chain
- **No new dependencies.** The `go.mod`/`go.sum` surface is **byte-identical to v1.8** across all of v1.9 (the work is seeder/cockpit/clerkenstein `.go` + stdlib Python + docs). All deps remain MIT/BSD/Apache. Lockfile: `knowledge/plan/releases/archive/01.90-storytelling/dependencies.lock`.

### Known limitations
- The cockpit + Stories layer is **opt-in** (`DEMO_STORIES=1`); default-off keeps every existing demo byte-identical. The literal browser-pixels end-to-end needs a `DEMO_STORIES=1` re-deploy — a deliberate demo step.
- The `clerk-express-1` alignment gate drives the genuine `@clerk/express` SDK and so needs installed npm modules (a node-CI env prerequisite) — unrunnable in the authoring copy; not a regression.

## [v1.8] "understudy" — 2026-06-15

The **self-contained-demo release**: `stack-demo/` gets its **own platform clone set** so a box with **only** `stack-demo/` (no `stack-dev/`) can run a demo end-to-end. A single `section` milestone (M26) re-implements the orphaned `m26/self-contained-demo` effort onto current `main`, preserving v1.6/v1.7. **Tooling + docs only — zero platform-repo edits** (`stack-demo/platform` is a build *context* only).

### Added
- **`demo-stack/ensure-clones.sh`** (net-new) — bootstrap-clones `stack-demo/platform` from GitHub + `make init`s the peer repos, so a box with only `stack-demo/` brings a demo up end-to-end. Seeds the shared `platform/.env` copy-if-present from `stack-dev` (same Clerk app + GH_PAT; non-fatal if absent — `/stack-secrets` provisions the real one). (M26)

### Changed
- **A demo now builds entirely from its own `stack-demo/` clone set.** `up-injected.sh` moves the build **source** + the compose dir (`PLAT`, `D-MAIN`) to `stack-demo`, so every service builds from `stack-demo` — it no longer borrows `stack-dev`'s repos or built images for the build source. Dev-image reuse is gated behind an opt-in `--reuse-dev-images` flag (OFF by default). (M26)

### Supply chain
- **No new dependencies.** The ext Go `go.mod`/`go.sum` diff is EMPTY (M26 is shell + Python + docs); Python stdlib-only + the pre-existing optional PyYAML test dep. Lockfile: `knowledge/plan/releases/archive/01.80-understudy/dependencies.lock`.

### Known limitations
- The live field-bake on a freshly-emptied `stack-demo/` is a user-authorized post-close follow-up (M26 satisfied the observable-behavior gate by composition, the M31/M32 precedent).

## [v1.7] "house lights" — 2026-06-15

The **demo-UI-hardening release**: when the house lights come up, the audience can see the show — a fresh browser at a demo's offset UI renders the working app with **zero manual steps**. Two browser-facing demo defects fixed (next-web blank page · studio-desk dead redirect). **Tooling + docs only — zero platform-repo edits.**

### Fixed
- **The demo next-web blank page is gone.** A fresh browser at a demo's offset next-web (e.g. `http://localhost:33000`) previously rendered **blank** because clerk-js's handshake to the fake FAPI (`https://127.0.0.1:35400`) hit an **untrusted self-signed cert** (`net::ERR_CERT_AUTHORITY_INVALID`) → clerk-js aborted. The demo bring-up now mints a **locally-trusted (mkcert) FAPI cert**, so clerk-js completes the handshake and the app renders with no manual cert-trust / proceed-anyway. (M31)
- **The demo studio-desk dead redirect is gone.** A fresh browser at a demo's offset studio-desk previously hit a **302 to a dead `:9100`** (the container ran its `NODE_ENV=development` redirect path). It now serves the **production single-port path** on `9000`+offset, landing on a live page. (M32)

### Changed
- **Demo bring-up auto-mints a locally-trusted mkcert FAPI cert.** Idempotent `mkcert -install` + cert mint for `127.0.0.1 localhost ::1`, with an **openssl self-signed fallback** (byte-compatible, when mkcert is absent or minting fails) and a **`DEMO_NO_MKCERT=1` opt-out** (forces openssl for operators who won't add a dev CA to their trust store). Non-fatal throughout — a cert step never aborts a good bring-up. (M31)
- **studio-desk now runs single-port `9000`+offset** in demos — the injection override pins `NODE_ENV=production` (+ `FRONTEND_PORT=9000`), so the production `sendFile` path wins and the cross-port `:9100` redirect never fires. The dead un-offset `:9100` CORS origin is dropped; the `:9100`→`9000` story is swept through the docs (`frontend-tier.md`, the demo-up SKILL). (M32)

### Supply chain
- **No new dependencies.** Both fixes are bring-up-script + injection-override changes (`demo-stack/up-injected.sh`, `stack-injection/gen_injected_override.py`) in `rosetta-extensions`; no new Go/Python/JS deps. Lockfile: `knowledge/plan/releases/archive/01.70-house-lights/dependencies.lock`.

### Known limitations
- **mkcert auto-trust is per the bring-up machine.** On a fresh machine, the first `mkcert -install` may **prompt once for the OS password** (to write the dev CA into the system trust store). **Remote/VM demos** still need the proceed-anyway fallback — `-install` trusts the bring-up machine, not the browsing machine. **Firefox** needs `certutil` for the dev CA to be trusted. (M31)

## [v1.6] "stage door" — 2026-06-14

The **secret-provisioning release**: one mechanism that ingests a secret source (a directory **or** zip, default `.agentspace/secrets`) and **provisions every repo of a stack** from it — **values-blind** (no verb ever reads, echoes, or logs a secret value) — verified by a **secret-coverage DNA** that *lists and keeps listed* the required secrets per repo. Retires the manual `.env` hand-copy. **Tooling + docs only — zero platform-repo edits; `.env` never committed; never writes prod.**

### Added
- **`/stack-secrets`** — provision a stack's secrets (`dev-N` / `demo-N`): write every repo's target `.env` from one secret source and verify coverage, values-blind. Drives the `stacksecrets` CLI (`check` / `provision` / `status`). Mirrors `/stack-seed`. (M29)
- A **secret-coverage DNA** — gene = repo × KEY (**6 repos / 55 genes**); `introspect` rebuilds the required set from a hybrid source, `diff` is a two-tier keep-listed gate (fatal only on an already-tracked secret omitted for a repo → catches vacuously-green coverage). (M27)
- **Directory + zip source provisioning** — one source value → each repo's correct per-file key (alias-mapped, e.g. the `gh-token` family across 3 files; distinct-similar pairs never auto-copied), with an explicit source-dir **layout contract** so a `zEnvs/` backup mirror or a stray `.env` can never be silently ingested. Per-repo target-file map pinned (`ant-academy → code/.env.local`, `next-web-app → apps/web/.env`). (M27/M28)
- A **demo-aware secret pre-flight** wired non-fatally into `/dev-up` + `/demo-up` (warn on standard-missing, fail on critical-missing; Clerk keys satisfiable by Clerkenstein minting on demo stacks; `DEV_/DEMO_NO_SECRET_PREFLIGHT=1` opt-outs). `provision` is idempotent (copy-if-absent, `--force` to overwrite) and **N=0 main-dev-stack-guarded** so it can't clobber the operator's working `.env`. (M28)
- **`corpus/ops/secrets-spec.md`** (net-new) — the secret-provisioning source-of-truth: the source-dir/zip layout contract, the 6-repo/55-gene DNA, the per-repo target-file map, the values-blind safety statement, the alias-family vs distinct-similar rules, the waived class, and the `0/1/3` exit contract. Plus the `/stack-secrets` skill, the CLAUDE.md skill-table + doc-index rows, and a `safety.md` clause. (M29)

### Changed
- **The manual `.env` hand-copy is retired** in favor of `/stack-secrets` — `setup_guide.md`'s hand-copy prose for studio-desk / ant-academy / next-web-app now points to the skill, and the in-tree **`setup_guide.md:447` TODO is deleted** (the per-repo key lists stay as reference). (M29)

### Fixed
- The **field-bake** (M30) proved the whole mechanism live on a fresh **demo-3** (17 containers UP, coverage Critical **100%**, prod `DIRECTUS_TOKEN` armed in **ZERO** containers) and caught + fixed **2 real release bugs**: (1) the demo secret pre-flight **silently skipped** — its source path resolved one level too shallow (doubled `.agentspace/.agentspace/secrets`), so the gate exited 2 instead of scoring; (2) the demo bring-up only *checked* coverage but **never provisioned** — added a non-fatal provision step (`DEMO_NO_PROVISION=1` opt-out). (M30)

### Supply chain
- The new `stack-secrets` Go module is **stdlib-only** — no `require` block, no `go.sum`, **0 new third-party deps** → a minimal values-blind audit surface. The 4 prior ext modules untouched; all-permissive licenses; `govulncheck` clean. Lockfile: `knowledge/plan/releases/archive/01.60-stage-door/dependencies.lock`.

### Known limitations
- The ~10–15% of keys that don't pass coverage are **entirely waived-class, by design — zero critical short**: AWS credentials supplied via the `~/.aws` mount, profile-gated keys (BREVO/messenger, customerio-sync), and optional Bunny/GCloud. (M30)
- **Encrypted-zip source** (age/gpg) was dropped as a deliberate v1 boundary — plain dir + plain zip cover the need; re-proposal requires a fresh design-roadmap pass.

## [v1.5] "prop room" — 2026-06-14

The **local-Directus release**: every stack now serves its **own captured public catalog** from a per-stack Directus (data plane local, asset plane prod → real images) on `--local-content` (demo default-on, dev opt-in). Prod-read remains the documented fallback. **Tooling + docs only — zero platform-repo edits.**

### Added
- A **per-stack local Directus** serving the captured public catalog locally, content-self-contained (M21–M23). `/demo-up` default-on; `/dev-up N --local-content` opt-in; `N=0` manual recipe.
- **Structure-bearing snapshot capture** — `stacksnap` captures the content-model DDL + primary keys + serve-row registration atomically with the rows, so a bootstrapped Directus auto-provisions and serves with **zero hand SQL** (M21).
- **Per-stack Directus as a compose service** — offset port, idempotent re-provision, torn down with the stack, verify-net probes (M22).
- A **cross-surface referential-closure gene** + the `directus_files` referenced-subset capture (M23).
- Corpus hygiene guards: a **README-index-row lint**, an **alignment zero-critical-genes guard**, and a **`/project-stats` scope fix** (M24).

### Changed
- The data plane (`DIRECTUS_BASE_ADDR`) re-points to the per-stack Directus on `--local-content`; the **asset plane stays on prod** (`content.anthropos.work`) so browser images stay real (M23).
- Go toolchain pinned to **go1.25.11** (clears the 12 stdlib advisories) (M24).
- The corpus tells the self-contained-content truth — retired the print-only / exit-4 / live-from-prod framing and a fictional local-Directus docker-service that never existed (M24 + close-release docs fix).

### Fixed
- A **`directus_files` tenant-data over-capture** that the firewall caught **fail-closed** under live prod capture — fixed in the capture filter; the firewall was never weakened (M25).
- A **dangling-FK class** — captured public content referencing uncaptured admin/library UI tables (M25).
- The **offline `GOTOOLCHAIN`** regression that aborted `/demo-up`'s clerkenstein build under the new Go pin (M25).

### Supply chain
- go1.25.11 toolchain pin on all 4 ext modules + the clerkenstein CI; `govulncheck`: **0 called CVEs**; all licenses permissive.

### Known limitations
- Media is served by **prod asset links**, not local blob bytes (DEF-M10-01, backlog).
- The local-content path is **opt-in for dev** (`--local-content`); `N=0` stays prod-read.
- 1 prod data-quality residual (a public sim referencing a customer-only skill, K-AIFUNX-E658) is operator-owned, uncloseable by tooling.

## [v1.3.1] "dress rehearsal" — 2026-06-09

Field-hardening of the demo/dev tooling: `/demo-up` now produces a **full, populated, verified, demoable** stack — closing the gaps the first real `/demo-up` run surfaced (14 field issues). **Tooling + docs only — zero platform-repo edits.**

### Added
- `/demo-up` brings up the **full UI tier** — next-web-app + studio-desk (per-demo *cached* image built from the **unmodified** platform Dockerfile, offset ports, minted Clerk pk + offset URLs baked) + ant-academy natively (Clerk-free); `--no-ui` to skip; a non-fatal 12 GB Docker-VM pre-flight. (M19)
- `/demo-up` now **auto-set-dresses by default** (cache-first snapshot replay → a `small-200` seed), exactly like `/dev-up` — reusing one shared set-dress engine; `--no-setdress` to skip. (M20)
- A **post-bring-up verification net**: every bring-up ends with a scoped, **non-fatal** check on the stack's own offset ports (`/api/health` + a Sentinel-policy-loaded assert) then the full probe set — so "UP" means *verified-working*, not just *containers-started*. (M18)
- New operator docs: `corpus/ops/idempotency.md`, `corpus/ops/verification.md`, `corpus/ops/demo/frontend-tier.md`, `corpus/ops/snapshot-cold-start.md` (the fresh-box capture runbook).

### Changed
- `stack-verify` is now **offset-/scope-aware** — it targets an individual `demo-N`/`dev-N` on its shifted ports + containers and checks only the services actually brought up (no more wall-of-false-"down"). (M18)
- The dev workspace rename **`anthropos-dev → stack-dev`** is now the documented default everywhere (legacy `anthropos-dev` retained as a single back-compat fallback). (M16)

### Fixed
- **Re-running a bring-up is now safe.** Snapshot-replay and seed are idempotent (a 2nd run replaces, never silently doubles or aborts mid-surface); `--reset` now clears the full data set; the `set -e` first-run race that could silently ship a 403-ing stack is fenced + regression-tested. (M17)
- The migrate-time **Sentinel-policy race** (a demo could report "UP" while authz silently failed to load → every authorized route 403'd) is fixed. (M16/M17)
- Stale `demo-stack` GUIDE/README facts corrected (test counts, "no remote", renamed skill names, version label). (M16)

### Supply chain
- No dependency changes. 0 called third-party CVEs; all-permissive licenses; lockfile recorded (`releases/archive/01.3b-dress-rehearsal/dependencies.lock`). The Go stdlib advisories clear with a go1.25.11+ toolchain.

### Known limitations
- A **new** `demo-N` pays one ~3-min *cached* frontend build per frontend (re-ups reuse the image). True zero-rebuild would need an optional, **user-owned upstream platform PR** — out of scope (the tooling never edits platform repos).
- Snapshot **media blob bytes + a cloud snapshot store** remain deferred to **v1.4** (DEF-M10-01, signed).

## [v1.3] "stack party" — 2026-06-07

dev + demo stacks become **first-class peers**: a dev stack gets the demo treatment (its own local Directus,
an auto-snapshot of the real reference data, a light default seed), one unified registry keeps dev and demo
from colliding on ports, and **one converged skill set** operates any stack — plus a single authoritative,
code-cited safety doc.

### Added
- **`/dev-up`** / **`/dev-down`** — the dev-stack lifecycle, mirroring `/demo-up` / `/demo-down`. `/dev-up`
  builds-or-starts the main dev stack (consolidating the former `setup-platform` + `start-platform`), and for
  an additional `dev-N` set-dresses it by default (local Directus + cache-first snapshot replay + a light
  `dev-min` seed). `/dev-down` tears a `dev-N` down and frees its registry slot. (M13/M14)
- **Dev stacks are now full-fidelity peers of demo for data** — a fresh `dev-N` is never empty: it gets a
  per-stack local Directus, an auto-snapshot replay of the real public taxonomy + content, and a light
  `dev-min` seed (~1 org + ~10 users), all default-on (escapes: `--no-snapshot`, `--no-setdress`). Capture is
  never run against prod from a dev stack (replay-only). (M13)
- **A unified stack registry + first-available-N allocation** — one shared N-pool spans dev *and* demo, so a
  bring-up always claims the lowest free slot and `dev-N`/`demo-N` can never collide on ports (e.g. building
  dev, demo, dev, demo, demo yields `dev-1, demo-2, dev-3, demo-4, demo-5`). `/stack-list` surfaces it. (M12)
- **`corpus/ops/safety.md`** — the authoritative, code-cited safety contract of the stack tooling: it **never
  reads private/customer data** (the tenant firewall + public predicates + bounded read-only capture) and
  **never touches production** (the 3-layer isolation guard + never-write shared Directus/prod-S3 + the audit-
  proven zero-pollution assertion). Every load-bearing claim is pinned to the source by a fail-closed drift
  guard. (M15)

### Changed
- **The stack-operation skills were hard-renamed to generic `stack-*` forms (no aliases)** — each accepts a
  `dev-N` or `demo-N` target: `/demo-status` → **`/stack-list`**, `/demo-seed` → **`/stack-seed`**,
  `/demo-snapshot` → **`/stack-snapshot`**, `/update-platform` → **`/stack-update`**. `/demo-up` / `/demo-down`
  stay as the demo lifecycle (now aligned with `/dev-up` / `/dev-down`). (M14)

### Removed
- The old skill names `/setup-platform`, `/start-platform`, `/update-platform`, `/demo-status`, `/demo-seed`,
  `/demo-snapshot` (and their skill dirs) — a clean break, no back-compat shims. Update any saved invocations
  to the converged names above. (M14)

### Supply chain
- No dependency changes; all deps permissive (MIT / BSD-3 / Apache-2.0 / ISC); **0 called third-party CVEs**.
  **Recommendation:** build with the **go1.25.11+** toolchain to clear 12 Go-stdlib (parsing/DoS-class)
  govulncheck findings (same class as v1.2). Lockfile: `knowledge/plan/releases/archive/01.30-stack-party/dependencies.lock`.

### Known limitations
- Demo/dev media still renders structure + file **references** (placeholder bytes); the actual S3 media blob
  **bytes** and a **cloud snapshot store** are now deferred to **v1.4** (DEF-M10-01 — gated on eu-west-1 S3-read
  access not wired here). *(This corrects the v1.2 changelog note below, which named v1.3 as the destination
  before the item was re-scoped to v1.4.)* AI-generated rich content (transcripts/embeddings) and external
  stack shareability are also v1.4.

## [v1.2] "set dressing" — 2026-06-07

The **snapshot mechanism**: *set-dress* a disposable demo stack with the **real public reference library** — the actual skills taxonomy and the Directus simulation/skill-path templates — so the catalog and the content behind seeded sessions are real, not placeholders. Everything is captured **read-only** from production; **customer data is never copied** (a tested tenant-data firewall).

### Added
- **`/demo-snapshot N`** — replay the captured public taxonomy + Directus content into a demo/dev stack (drives the `stacksnap` CLI: capture / replay / status). New demo flow: `/demo-up → /demo-snapshot → /demo-seed`. (M11)
- **`/db-query`** — read-only production DB investigation skill (the wired `postgres` MCP tool, or Tailscale + `~/.pgpass`), with the public-vs-customer data boundary documented. (M9a)
- **The `stack-snapshot` extension** — capture a public reference surface once from a safe source, manifest-cache it under `.agentspace`, replay per-stack — behind a firewall that hard-fails on any captured customer row. (M9a)
- **Real public skills taxonomy** in demo stacks (42.8K skills + roles + embeddings + translations; pgvector index rebuilt on replay). (M9b)
- **Real public Directus content** in demo stacks (published simulation/skill-path templates), served by a per-stack Directus; seeded sessions/assignments now resolve to real templates. (M10)
- A **snapshot-fidelity** alignment dimension and **100% data-DNA coverage** (the two formerly-`waived` surfaces promoted).
- New docs: `corpus/ops/snapshot-spec.md`, `corpus/ops/db-access.md`; the set-dressed `corpus/ops/demo/` recipe family (incl. `recipe-snapshot-world.md`).

### Changed
- Seeding spec: `taxonomy` + `content` promoted `waived` → `snapshot-seeded` (100% catalog coverage; nothing waived).
- Capture-source policy: default **ingest an existing prod `pg_dump`**, fallback a **safe throttled read-only primary read** (PostgreSQL MVCC = a read never blocks prod writes).

### Removed
- The planned offline `pg_dump`-**FILE** reader was dropped — restore-then-`--dsn` covers the need with no new capability or speed gain; the `--dump` flag is gone.

### Supply chain
- No dependency changes; all deps permissive (MIT / BSD-3 / Apache-2.0 / ISC); **0 third-party CVEs**. **Recommendation:** build with the **go1.25.11+** toolchain to clear 12 Go-stdlib (DoS/parsing-class) govulncheck findings.

### Known limitations
- Demo media renders structure + file references; the actual S3 blob **bytes** and a **cloud snapshot store** are deferred to v1.3 (gated on S3/AWS access not wired here).

## [v1.1] "show floor" — 2026-06-05

The platform-operations extension framework: spin up a **disposable, Clerk-free, realistically-populated** copy of the platform — for a demo, screenshots, or QA — alongside the dev stack, **without touching production or any read-only platform repo**.

### Added
- **Disposable demo stacks** (`/demo-up`, `/demo-down`, `/demo-status`) — bring up `demo-N` isolated on offset ports, Clerkenstein-wired (Clerk-free), with its own data; killable cleanly; the dev stack never touched. (M3)
- **Declarative, production-safe seeding** (`/demo-seed` + the `stackseed` tool) — backfill a stack from one `stack.seed.yaml` (or a curated preset): an org + 1,000 users + the real `user_clerkenstein` login identity + months of **backdated** job-sim / skill-path sessions, assignments, and activity — a believable world a stakeholder can log into (authorized routes return **200**). It connects **directly to the stack's Postgres** (`COPY`; ~0.7s for ~9,500 rows) behind a **3-layer production-isolation guard** that makes it *structurally impossible* for a non-prod run to write a shared/prod store (Directus, the prod S3-public bucket, live Clerk, marketing/AI SaaS), and proves zero pollution with an audit log. (M7a)
- **The data-DNA** — the alignment framework extended to a third dimension, **data**: the `datadna` CLI enumerates the seedable surfaces, **conformance-gates** each seeder's output against the platform's current schema, and **detects drift** when that schema moves (`measure` 100% / `diff` flags a changed column). (M7b)
- **The seeder fleet** — backdated-activity seeders for the believability core (job-sim + skill-path sessions, assignments, activity events), driven to a data-DNA coverage gate. (M7c)
- **`dev-stack`** — the same multi-instance tooling for isolated *dev* stacks (`dev-N`), real Clerk by default, optional Clerkenstein injection. (M6)
- **Demo-env corpus family** (`corpus/ops/demo/`) — a family index + 3 end-to-end recipes (enterprise onboarding · skill progression · interactive browser login) + 3 curated seed presets (small-200 / mid-500 / large-1k). (M8)
- **`@clerk/express` alignment gate wired into CI** — the v1.0 carry-forward; clerkenstein's CI now materializes the SDK + runs the gate (validated 9/9). (M8)

### Changed
- **Repo consolidation** — the standalone `clerkenstein` + `rosetta-demo` repos collapsed into one private `rosetta-extensions` monorepo (history preserved via `git subtree`); the old org repos were removed; `rosetta` thinned to documentation + the alignment framework + pointers. The reusable Clerk-mock injection layer (`stack-injection`) and the shared port-offset engine (`stack-core`) were extracted as sections. (M4/M5/M6)
- **Clerkenstein gained a 4th measured surface** — deployment/injection (`clerk-deploy-1`, 7/7): the disarmed `colony/authn` drop-in compiles against the platform's real `colony` and satisfies its contract. All four gates held 100%/100% throughout v1.1. (M3 extended)

### Known limitations
- Seeding ships **structural data only** by design. Two surfaces are **waived** to v1.2: the skill *taxonomy* (needs a pre-embedded skiller snapshot) and Directus *content* (the shared instance — snapshot-replay only). AI-generated rich content (transcripts/embeddings) is also out of scope. Data-DNA coverage is 100% over the 8 reachable surfaces; the waived surfaces are recorded as `waived-m7c` in the manifest.
- The deployment/injection alignment gate stays a **local** gate (it needs the platform's `colony` via a private token); the other three surfaces run in CI.

### Supply chain
- No new runtime deps beyond the Postgres driver (`jackc/pgx/v5 v5.9.2`) + `gopkg.in/yaml.v3` for the seeder. All deps permissive (BSD/MIT/Apache); lockfile at `knowledge/plan/releases/archive/01.10-show-floor/dependencies.lock`.

## [v1.0] "body double" — 2026-06-03

The first release under the developer-kit planning lifecycle. Rosetta gains a measurement discipline and its first product: a drop-in Clerk mock you can *prove* is faithful.

### Added
- **Alignment-testing framework** (`test/alignment/`) — a third test class that scores how faithfully a *mirror* engine reproduces a *source* library as a 0–100% number (overall + a separate critical gate). Ships `alignctl` (stdlib-only Go, runs offline) with `run`/`capture`/`dna list|diff|validate`, engine-agnostic via a pluggable `--runner`; four equivalence operators; record/replay goldens; a `//go:build alignment` test class; and a toy reference that proves end-to-end divergence detection. Plus the `/align-dna` and `/align-run` skills and the canonical doc `corpus/architecture/alignment_testing.md`. (M0)
- **Clerkenstein** — a *measured* drop-in mock of the Clerk libraries the platform uses, so demos run Clerk-free with **zero platform-code change**. Verified at **100% / 100% on all three measured surfaces**: the Go SDK (`clerk-sdk-go/v2`, 22 genes — M1), the JS/FAPI browser surface (`@clerk/clerk-js`+`@clerk/nextjs`, 9 genes — M2), and the `@clerk/express` Node backend (9 genes, RS256/JWKS — M2c). Lives in its own repo; injected via build-time `go.mod replace` + a fake Clerk API server.
- **Drift detection** (M1b) — a `gate.sh`/`drift-check.sh` exit-code contract + a weekly CI alignment gate + a drift regression harness, so a Clerk version bump becomes a flagged, mechanical event instead of a silent break.
- **Webhook + browser-session coherence** (M2) — a fake FAPI (browser login via a minted publishable key, config-only), a fake BAPI that disarms the platform's networked org client, and an svix-signed webhook injector.

### Changed
- The `clerkenstein` repo was consolidated into a **library-named** layout (one dir per mocked dependency) with a self-contained `knowledge/` base (M2b). No behavior change.
- `corpus/services/clerkenstein.md`, `corpus/architecture/alignment_testing.md`, and `README.md` now describe the alignment framework + the measured mock, with an explicit provenance note on what "100%" means.

### Fixed
- **`@clerk/express` alignment gate regression** — the bad-signature scenario was misclassified (`malformed` vs `bad-signature`) by a flawed signature-tamper introduced in the M2c close; the gate had silently dropped to 88.0%/85.7%. Corrected to a byte-level signature corruption with a property-based regression test; gate restored to 9/9 100%/100%. (caught + fixed at release close)
- A cross-milestone documentation/decision drift sweep (canonical framework doc, planning context, resolved decisions, metrics field).

### Known limitations
- "100%" means *indistinguishable from hand-authored / hybrid source goldens* (the reference behavior derived from the real libraries) — **not** a byte-diff against a live, network-connected Clerk tenant. Re-capturing goldens on a Clerk version bump is the drift loop's job. This is the right bar for a demo mock, not a production-Clerk conformance certificate.
- The `@clerk/express` gate is the only non-pure-Go gate (it drives the genuine Node SDK), so it runs locally / at close rather than in the pure-Go CI. CI-wiring is staged for v1.1.

### Supply chain
- The rosetta-committed code (`test/alignment/`) is **stdlib-only** — zero external modules. See `releases/archive/01.00-body-double/dependencies.lock`.
