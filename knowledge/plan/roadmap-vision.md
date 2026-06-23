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

---

> **No version currently staged.** v1.9 "storytelling" shipped 2026-06-23 (tag `v1.9`; full detail in the
> `## Done — v1.9 "storytelling"` section of [`roadmap.md`](roadmap.md)). The next version awaits a
> `/developer-kit:design-roadmap` run — the unscheduled backlog below is candidate input.

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

**Resolved (no longer backlog):**

- **M33 — ant-academy demo liveness** (deferred from v1.7 design, 2026-06-15, repro-first) → **RESOLVED post-v1.9** at
  rext tag `storytelling-postfix-1` (a tooling-only post-v1.9 demo-hardening pass). The "dead on a later visit"
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
  **Residual ant-academy reality:** academy still needs the team Font Awesome Pro token (`FONTAWESOME_NPM_AUTH_TOKEN`)
  + a one-time `npm install` in `stack-demo/ant-academy/code` to actually run; without the token it's an intentional
  **non-fatal skip** (it's a Vercel-deployed, Clerk-only peripheral surface — the cockpit / next-web / studio-desk carry
  the demo). Provision the token (e.g. via `/stack-secrets`) then `npm install` once to enable it.

**Dropped from tracking (2026-06-11, user instruction — re-proposal requires a fresh `/developer-kit:design-roadmap` run):**
the former v1.4 seeds **AI-generated content**, **external stack shareability** (Tailscale/ingress), and **more
mirror engines**; the **deployment/injection CI gate** (a local-only alignment surface; gates nothing in the
demo/dev workflow); and the **`/dev-up` frontend-image pre-warm** question (a UX nicety with no owner).

## Codename notes
- _(v1.0 "body double" + v1.1 "show floor" + v1.2 "set dressing" + v1.3 "stack party" + v1.3b "dress rehearsal" + v1.5 "prop room" + v1.6 "stage door" + v1.7 "house lights" + v1.8 "understudy" shipped — their codenames are now permanent. **v1.8 "understudy"** continued the theatre lineage: an understudy is a fully self-contained substitute, ready to perform on its own without the lead — exactly the self-contained-demo thesis (`stack-demo/` becomes able to run with no `stack-dev/`). Chosen at the 2026-06-15 `/developer-kit:design-roadmap` run.)_
- **v1.9 "storytelling"** (shipped 2026-06-23, tag `v1.9` — codename now permanent) continues the theatre lineage and names the thesis directly: the release is about making the seeded world **tell a story** — declarative *stories*, each with a cast of *heroes* whose verified-skill histories the product surfaces narrate. Chosen by the user at the 2026-06-22 `/developer-kit:design-roadmap` run (over the proposed "method acting" / "dramatis personae").

_Last updated: 2026-06-23 (**M33 ant-academy demo liveness RESOLVED post-v1.9** at rext tag `storytelling-postfix-1`
— the session-reaping was reproduced + fixed via session-detach [`launch_detached`]; the same tooling-only
post-v1.9 demo-hardening pass also made `DEMO_STORIES` the default, added the Directus boot health-gate, and
guarded the prod-Directus note. Moved M33 out of backlog → resolved. Backlog now: DEF-M10-01, DEF-M21-01, M25-D9.
Prior: 2026-06-23 **v1.9 "storytelling" SHIPPED** [tag `v1.9`] via `/developer-kit:close-release` — reviewed
M34→M38 as one PR, GREEN/0 blocking, deferral re-audit GREEN [0 escape-hatch], merged `release/01.90-storytelling`
→ `main`; 2026-06-22 v1.9 DESIGNED + PROMOTED [5 `section` milestones M34→M38]; 2026-06-15 v1.8 "understudy"
SHIPPED [tag `v1.8`].)_
