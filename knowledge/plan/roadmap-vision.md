# Roadmap — Vision

Future versions and proposals, not yet in active development. The active version lives in
[`roadmap.md`](roadmap.md). When a version starts development, its section moves from here into `roadmap.md`
and a `release/{version}` branch is cut.

> **Promotion history:** **v1.0 "body double"** → `roadmap.md` 2026-06-02 (shipped 2026-06-03, tag `v1.0`).
> **v1.1 "show floor"** → `roadmap.md` 2026-06-03 (now in development on `release/01.10-show-floor`).
> **No future version (v1.2 / v2+) is scoped yet** — this file is the staging area for the next one.

---

## No versions currently staged

The next version will be drafted here (via `/developer-kit:design-roadmap`) once v1.1 "show floor" is underway
and its successor's shape is clear. Likely seeds, carried forward from v1.1's open decisions:

- **External shareability of demo stacks** — Tailscale-only (like staging) vs public ingress for
  customer-facing demos (a security-posture decision deferred from v1.1; could anchor a "shareable demos"
  version).
- **AI-generated demo content** — LLM-generated transcripts / AI-scored validation narratives /
  freshly-computed embeddings. v1.1's seeding (M7a/M7b/M7c) draws a hard line excluding these — it ships
  **structural data only** (the data-DNA discipline + the seeder fleet make *that* robust + drift-proof). AI
  content is the natural v1.2 theme: richer, more believable demo worlds layered on the M7 seeding foundation.
- **Mirroring engines beyond Clerk** — the M0 alignment framework is engine-generic; v1.0 exercised it on Clerk
  (behavioral) and v1.1's M7b extends it to **data** (structural). A future version could mirror another
  third-party dependency, or add further alignment dimensions.

## Codename notes
- **v1.1 "show floor"** (in development — codename still changeable until ship): alternatives were
  "open house" · "set piece" · "dry run".
- _(v1.0 "body double" shipped — its codename is now permanent.)_

_Last updated: 2026-06-04 (v1.1 seeding redesigned into M7a/M7b/M7c — kept in v1.1; v1.2 seeds refined: AI-content layers on the M7 foundation, and M7b already extends M0 to a data dimension)._
