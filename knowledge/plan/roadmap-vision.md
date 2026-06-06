# Roadmap — Vision

Future versions and proposals, not yet in active development. The active version lives in
[`roadmap.md`](roadmap.md). When a version starts development, its section moves from here into `roadmap.md`
and a `release/{version}` branch is cut.

> **Promotion history:** **v1.0 "body double"** → `roadmap.md` 2026-06-02 (shipped 2026-06-03, tag `v1.0`).
> **v1.1 "show floor"** → `roadmap.md` 2026-06-03 (shipped 2026-06-05, tag `v1.1`).
> **v1.2 "set dressing"** → `roadmap.md` 2026-06-05 (in development on `release/01.20-set-dressing`; **refined
> 2026-06-06** — a **dedicated `stack-snapshot` extension** lifting M7c's `waived` taxonomy + content to 100%
> coverage, capturing **public** data only from a **safe non-primary source**, cached in `.agentspace`).
> **No future version (v1.3 / v2+) is scoped yet** — this file is the staging area for the next one.

---

## v1.3 seeds (carried forward from v1.2 scoping)

The next version will be drafted here (via `/developer-kit:design-roadmap`) once v1.2 "set dressing" is underway
and its successor's shape is clear. These seeds were **explicitly held out of v1.2** (user, 2026-06-05 — v1.2
stays the tight snapshot-spine release), and are the leading v1.3 candidates:

- **AI-generated rich content** — LLM-generated transcripts / AI-scored validation narratives /
  freshly-computed embeddings, **layered on the v1.2 snapshot foundation** (v1.2 replays *real* captured taxonomy
  + content; v1.3 would *generate* richer, fresher content on top). The natural iterative-shaped milestone (like
  M7c): non-deterministic outputs mean the determinism bar / caching / drift-tolerance path emerges from
  building. **Phase 0b note:** YELLOW — the Studio-Room generation pipeline is documented
  (`ai_architecture.md`, `cms.md`) but there's no *seeding-at-scale* spec; v1.3 would `Deliver →` an
  AI-content-seeding spec.
- **External shareability of demo stacks** — Tailscale-only (like staging) vs public ingress for
  customer-facing demos (a security-posture decision deferred since v1.1). **Phase 0b note:** **RED blind area**
  — only mentioned in this vision doc; no corpus anchor. A v1.3 milestone must `Deliver →` a demo-shareability
  architecture doc (the decision + access-control/TLS/domain model) before building.
- **Mirroring engines beyond Clerk** — the M0 alignment framework is engine-generic; v1.0 exercised it on Clerk
  (behavioral), v1.1's M7b extended it to **data** (structural), and v1.2 adds a **snapshot-fidelity** dimension.
  A future version could mirror another third-party dependency, or add further alignment dimensions.
- **Cloud snapshot store** — v1.2 caches snapshot payloads in a **local `.agentspace/snapshots/` store** (gitignored;
  no GB blobs in any git repo). It doesn't share across machines or scale. v1.3 swaps the `SnapshotStore` backend
  for **cloud object storage (S3)** — a backend change behind the existing interface (the manifest already addresses
  by location). **Explicitly named as a v1.2 → v1.3 follow-up** (user, 2026-06-06 note #4). **Phase 0b note:** the
  v1.2 `corpus/ops/snapshot-spec.md` anchors the store interface; v1.3 extends it with the cloud-backend section.
- **The deployment/injection CI gate** — currently a *local* gate (needs colony via GH_PAT); a future runner with
  org-secret access could wire it into CI (small tooling item — fold into a v1.3 B-milestone or leave local).

## Codename notes
- **v1.2 "set dressing"** (in development — codename still changeable until ship): alternatives were
  "full house" · "deep roots" · "true colors".
- _(v1.0 "body double" + v1.1 "show floor" shipped — their codenames are now permanent.)_

_Last updated: 2026-06-06 (v1.2 refined — dedicated `stack-snapshot` extension, public-only + safe-source capture,
`.agentspace` manifest cache; added the **cloud snapshot store** v1.3 seed (note #4). AI-content + shareability +
more-mirrors + cloud-store held as v1.3 seeds)._
