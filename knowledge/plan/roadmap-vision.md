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

---

> **No version is currently staged.** v1.7 "house lights" shipped 2026-06-15 (tag `v1.7`; full detail in the
> `## Done — v1.7` section of [`roadmap.md`](roadmap.md)). The next version is **unplanned** — run
> `/developer-kit:design-roadmap` to scope it.

## Unscheduled backlog (not a planned release)

Genuinely-deferred work, no target version, not scheduled:

- **M33 — ant-academy demo liveness** (deferred from v1.7 design, 2026-06-15, repro-first). The ant-academy demo
  surface runs **native** (nohup + pidfile + a kill-0 relaunch guard in `ant-academy.sh`), not a container — so on a
  later visit `:33077+offset` can be dead. Root cause is **unconfirmed** (most likely the launching `/demo-up` session
  reaping the nohup'd process tree on session-end, NOT a tooling bug — re-running `/demo-up` relaunches it). **Scope
  only after reproducing the exact "dead on later visit" scenario**; likely doc-only (document the "native — re-run
  `/demo-up` to relaunch" reality in `frontend-tier.md`) or a minimal liveness loop if repro proves a real gap. Smaller
  payoff (academy is the least-central, Vercel-native, Clerk-only demo surface) — deliberately left out of v1.7's firm scope.
- **M26 — self-contained demo stacks.** An orphaned `rosetta-extensions` effort: branch `m26/self-contained-demo`
  @ `25ab855`, tag `prop-room-m26` (ext, local-only, **unmerged + unpushed**), "make demo stacks self-contained
  (their own GitHub clone set, like stack-dev)" (+521/−141 in `demo-stack/` + `stack-injection/`, authored
  2026-06-14). It consumed the flat milestone counter's **M26** slot before v1.6 "stage door" was designed; on
  discovering the collision the user kept self-contained-demo as M26 and renumbered the secret-provisioning release
  to M27→M30 (see [`context.md`](context.md) + [`roadmap.md`](roadmap.md) "Why v1.6 starts at M27"). **Awaits its
  own `/developer-kit:design-roadmap` pass for a version + scope** — not yet placed in any release.
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

**Dropped from tracking (2026-06-11, user instruction — re-proposal requires a fresh `/developer-kit:design-roadmap` run):**
the former v1.4 seeds **AI-generated content**, **external stack shareability** (Tailscale/ingress), and **more
mirror engines**; the **deployment/injection CI gate** (a local-only alignment surface; gates nothing in the
demo/dev workflow); and the **`/dev-up` frontend-image pre-warm** question (a UX nicety with no owner).

## Codename notes
- _(v1.0 "body double" + v1.1 "show floor" + v1.2 "set dressing" + v1.3 "stack party" + v1.3b "dress rehearsal" + v1.5 "prop room" + v1.6 "stage door" + v1.7 "house lights" shipped — their codenames are now permanent. "house lights" continued the theatre lineage: when the house lights come up, the audience can see the show — v1.7 made the demo's browser UI actually render [the blank page goes away] instead of staying dark. No version is currently staged; the next codename will be chosen at the next `/developer-kit:design-roadmap` run.)_

_Last updated: 2026-06-15 (**v1.7 "house lights" SHIPPED** — the demo-UI-hardening release, M31→M32, merged
`--no-ff` → `main`, tag `v1.7`; full detail in the `## Done — v1.7` section of roadmap.md. No version currently
staged — next unplanned. Backlog unchanged: M33 ant-academy liveness [repro-first], M26 self-contained-demo,
DEF-M10-01, DEF-M21-01, M25-D9. Prior: 2026-06-15 v1.7 staged + in development.)_
