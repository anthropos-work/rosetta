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
> **v1.10 "method acting"** → 2026-06-24 (IN DEVELOPMENT, branch `release/01.10-method-acting`; the
> **believable-profile release** — make each hero hold up under a close-up when a presenter *Logs in as* them:
> profile identity [org name + role + real-face avatar] + the content-surface unblock [library + activity feed via
> one Directus serve-grant] + profile depth [work/education/deep skills] + **100% per-vantage demo coverage** proven
> by Playwright; 5 milestones M39→M42m [3 `section` + 2 `iterative`]; tooling + docs only). Designed from the
> live-demo review [`.agentspace/profile_gaps.md`](../../.agentspace/profile_gaps.md) — a hero (Maya Chen) read
> empty across the profile/library/activity surfaces. The second consecutive demo-believability release after v1.9.

---

> **v1.10 "method acting" is IN DEVELOPMENT** (designed 2026-06-24; branch `release/01.10-method-acting`; full
> detail in the `## In Development — v1.10` section of [`roadmap.md`](roadmap.md)). The unscheduled backlog below is
> orthogonal (not in v1.10 scope).

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
- **DEF-M46-01 — Directus serve-grant CLOSURE + schema RECAPTURE (Option B; the robust all-drift fix).** M46/DD landed
  a **targeted** column reconciliation (Option A): the captured per-stack Directus structure had **drifted** behind the
  platform (cms's `SetFields("*", …)` simulations query SELECTs `simulations.is_interview_validation_enabled`, a column
  added to prod Directus after the snapshot was captured → `Directus 500: column does not exist`). Option A is an
  idempotent **post-replay `ADD COLUMN IF NOT EXISTS`** backfill in `up-injected.sh` (the FK-indexes mechanism class;
  `DEMO_NO_DIRECTUS_DRIFT_FIX` opt-out) — reproducible + demo-local + zero canonical edit. **It fixed the column 500 (0
  `does not exist` on a cold stack), but the cold sweep surfaced a DEEPER, distinct blocker on `/enterprise/activity-dashboard`
  that Option A does NOT cover and is squarely Option B:** the M40 serve-grant `servedCollections` set (`stack-snapshot/directus/structure.go`)
  is **incomplete for the full cms `GetJobSimulation` deep-fetch closure**. cms requests `sequences.knowledge.*`,
  `sequences.assets_files.directus_files_id.*`, `sequences.collaborative_assets.*`, `sim_features.*`, `translations.*`,
  but the target/junction collections — `knowledge_asset`, `sequences_files`/`_2`, `directus_files`, `sim_features`,
  `sim_roles_tasks`, `sim_translations`, `simulations_translations` — are NOT registered/granted/related in the per-stack
  Directus (verified absent in the **current** cache `ea2e187`'s `_structure.sql` too → a FRESH `/demo-up` hits it). When
  a deep `*`-alias targets an unregistered/ungranted collection, Directus drops the WHOLE parent `sequences` alias → cms
  `jobsimulation.go:1097 s.Sequences[0]` panics (`index out of range`) → `jobSimulation.title` null → the federation
  fails the non-nullable field → the activity-table never hydrates → manager gate `failingSections=1`. **Option B** is the
  durable fix: (a) EXPAND `servedCollections` to the full deep-fetch closure, and (b) RECAPTURE the current prod Directus
  structure so the relation/field metadata for those collections is captured from prod (the M40 design captures relation
  wiring from the sanctioned prod structural read — junction_field/one_field/deselect — rather than fabricating it; a
  capture-path milestone, per the capture-path live-acceptance pattern: re-capture + cache-bust + a FRESH `/demo-up`,
  under the public-only firewall + capture-source policy). With the closure served and a current captured structure, the
  per-column backfill becomes a no-op and the activity-dashboard renders. See `corpus/ops/snapshot-spec.md` →
  "Schema-drift reconciliation (M46 …)". Tracked for a future `stack-snapshot` recapture milestone.

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
- **v1.10 "method acting"** (IN DEVELOPMENT — chosen 2026-06-24, the runner-up codename from the v1.9 round, now apt): continues the theatre lineage and names the thesis directly — *method acting* is the deep, immersive work that makes a single **character** believable up close, exactly v1.10's job (the hero you log in as must read as a real, fully-fleshed person on every page). Alternatives weighed: "in character", "close-up".

_Last updated: 2026-06-24 (**v1.10 "method acting" PROMOTED to active development** — the believable-profile
release, 5 milestones M39→M42m [3 `section` fills + 2 `iterative` Playwright per-vantage coverage gates], branch
`release/01.10-method-acting` cut, designed from `.agentspace/profile_gaps.md`. Standing backlog [DEF-M10-01,
DEF-M21-01, M25-D9] is orthogonal. Prior same day: **M33 ant-academy demo liveness FULLY RESOLVED post-v1.9** at rext tags `storytelling-postfix-1`
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
