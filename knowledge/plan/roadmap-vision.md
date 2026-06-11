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
> **v1.5 "prop room"** → 2026-06-11 (**in development**; the **local-Directus release** — a real per-stack Directus
> serving the captured public content, demo-default + dev-opt-in, M21→M25; promoted to [`roadmap.md`](roadmap.md),
> branch `release/01.50-prop-room` cut). The first version staged after the v1.4 removal.

---

## Staged: v1.5 "prop room" (in development — full detail in [`roadmap.md`](roadmap.md))

The **local-Directus release** is staged and in development. It closes the M10 collection-schema gap (NEW-1),
executes the per-stack-Directus recipe (NEW-2), and makes content↔taxonomy a referentially-closed captured pair
(NEW-3) — so every demo stack (default) and any opt-in dev stack serves the captured public content from its **own**
Directus instead of reading live from prod. Real images are preserved via prod public asset links.

## Unscheduled backlog (not a planned release)

Genuinely-deferred work, no target version, not scheduled:

- **DEF-M10-01 — cloud `SnapshotStore` backend + S3 media blob bytes.** Today the cache is the local
  `.agentspace/snapshots/` store and media replays **refs-only**. **Re-signed → backlog at v1.5 design (2026-06-11)**
  after its v1.4 destination was removed; its **user-facing sting is gone** — v1.5 "prop room" keeps the asset plane
  on prod public links so demos show **real images** without the blob-byte work. Real blob mirroring + the cloud
  store stay gated on **eu-west-1 S3 read access actually landing** (verified not wired). Replay-only to a per-stack
  isolated bucket, never the shared prod S3.

**Dropped from tracking (2026-06-11, user instruction — re-proposal requires a fresh `/developer-kit:design-roadmap` run):**
the former v1.4 seeds **AI-generated content**, **external stack shareability** (Tailscale/ingress), and **more
mirror engines**; the **deployment/injection CI gate** (a local-only alignment surface; gates nothing in the
demo/dev workflow); and the **`/dev-up` frontend-image pre-warm** question (a UX nicety with no owner).

## Codename notes
- _(v1.0 "body double" + v1.1 "show floor" + v1.2 "set dressing" + v1.3 "stack party" + v1.3b "dress rehearsal" shipped — their codenames are now permanent. "prop room" (v1.5, in dev) continues the stage-metaphor lineage: the prop room is where a show's real objects live — v1.5 gives each stack its own real content instead of borrowing prod's.)_

_Last updated: 2026-06-11 (**v1.5 "prop room" staged + in development** — the local-Directus release, M21→M25,
promoted to roadmap.md; branch `release/01.50-prop-room` cut. DEF-M10-01 re-signed → backlog; the ex-v1.4 seeds +
deploy-CI gate + dev-up pre-warm DROPPED from tracking. Prior: 2026-06-11 v1.4 removed.)_
