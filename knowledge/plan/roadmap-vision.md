# Roadmap — Vision

Future versions and proposals, not yet in active development. The active version lives in
[`roadmap.md`](roadmap.md). When a version starts development, its section moves from here into `roadmap.md`
and a `release/{version}` branch is cut.

> **Promotion history:** **v1.0 "body double"** → 2026-06-02 (shipped 2026-06-03, tag `v1.0`).
> **v1.1 "show floor"** → 2026-06-03 (shipped 2026-06-05, tag `v1.1`).
> **v1.2 "set dressing"** → 2026-06-05 (shipped 2026-06-07, tag `v1.2`).
> **v1.3 "stack party"** → 2026-06-07 (shipped 2026-06-07, tag `v1.3`; the **dev/demo convergence** — dev stacks as
> first-class peers, a unified first-available-N stack registry, generic `stack-*` skills, a code-cited safety doc).
> **v1.4 is staged below** — the former v1.3 seeds (cloud store / S3 blobs, AI content, shareability, more mirrors)
> moved here when the user scoped v1.3 = stack party (2026-06-07); now the next version to draft via
> `/developer-kit:design-roadmap`.

---

## v1.4 seeds (moved from v1.3 when the user scoped v1.3 = "stack party", 2026-06-07)

These were the leading v1.3 candidates, **moved to v1.4** when the user chose the dev/demo-convergence theme for
v1.3 (2026-06-07). They'll be drafted into v1.4 (via `/developer-kit:design-roadmap`) once v1.3 "stack party" is
underway. The leading v1.4 candidates:

- **Cloud snapshot store + S3 media blob bytes** — the **signed v1.2 escape-hatch** (DEF-M10-01 + M9a-D4). v1.2
  caches snapshot payloads in a **local `.agentspace/snapshots/` store** (gitignored) and replays Directus media
  **refs-only** (placeholder bytes). v1.4 swaps the `SnapshotStore` backend for **cloud object storage (S3)** behind
  the existing interface (the manifest already addresses by location) **and** captures the actual blob bytes (gated
  on S3-read access). **Phase 0b note:** `corpus/ops/snapshot-spec.md` anchors the store interface; v1.4 extends it
  with the cloud-backend + blob-bytes section.
- **AI-generated rich content** — LLM-generated transcripts / AI-scored validation narratives /
  freshly-computed embeddings, **layered on the v1.2 snapshot foundation** (v1.2 replays *real* captured taxonomy
  + content; v1.4 would *generate* richer, fresher content on top). The natural iterative-shaped milestone (like
  M7c): non-deterministic outputs mean the determinism bar / caching / drift-tolerance path emerges from
  building. **Phase 0b note:** YELLOW — the Studio-Room generation pipeline is documented
  (`ai_architecture.md`, `cms.md`) but there's no *seeding-at-scale* spec; v1.4 would `Deliver →` an
  AI-content-seeding spec.
- **External shareability of stacks** — Tailscale-only (like staging) vs public ingress for customer-facing demos
  (a security-posture decision deferred since v1.1). Now applies to **both** dev and demo (post-v1.3 convergence).
  **Phase 0b note:** **RED blind area** — a v1.4 milestone must `Deliver →` a stack-shareability architecture doc
  (the decision + access-control/TLS/domain model) before building.
- **Mirroring engines beyond Clerk** — the M0 alignment framework is engine-generic; v1.0 exercised it on Clerk
  (behavioral), v1.1's M7b extended it to **data** (structural), and v1.2 added a **snapshot-fidelity** dimension.
  A future version could mirror another third-party dependency, or add further alignment dimensions.
- **The deployment/injection CI gate** — currently a *local* gate (needs colony via GH_PAT); a future runner with
  org-secret access could wire it into CI (small tooling item — fold into a v1.4 B-milestone or leave local).

## Codename notes
- **v1.3 "stack party"** (in development — user-named; codename still changeable until ship). Note: it breaks the
  prior stage-metaphor lineage (body double → show floor → set dressing) in favor of the convergence feeling;
  stage-metaphor alternatives if wanted: "full ensemble" · "house lights" · "company call".
- _(v1.0 "body double" + v1.1 "show floor" + v1.2 "set dressing" shipped — their codenames are now permanent.)_

_Last updated: 2026-06-07 (**v1.3 "stack party" promoted** to roadmap.md — the dev/demo convergence; the former v1.3
seeds (cloud store / S3 blobs, AI content, shareability, more mirrors) moved to **v1.4 seeds** above)._
