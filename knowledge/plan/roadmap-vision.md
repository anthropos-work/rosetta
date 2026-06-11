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
>
> **No version is currently staged for development** (v1.4 was removed, 2026-06-11). When the next version is
> chosen, its section is drafted here (via `/developer-kit:design-roadmap`) and a `release/{version}` branch is cut.

---

## No staged next version

There is **no version currently staged** (v1.4 was removed on 2026-06-11). Genuinely-deferred work is tracked as
**unscheduled backlog**, not as a planned release — chiefly **DEF-M10-01** (the cloud `SnapshotStore` backend + S3
media blob bytes; today the cache is the local `.agentspace/snapshots/` store and media replays refs-only). It has
no target version and is not scheduled. The per-stack-Directus **content** replay (the M10 collection-schema gap —
the `directus` surface still skips on every stack) is likewise unscheduled backlog; see
[`../../corpus/ops/snapshot-spec.md`](../../corpus/ops/snapshot-spec.md) § the per-stack Directus store fork.

## Codename notes
- _(v1.0 "body double" + v1.1 "show floor" + v1.2 "set dressing" + v1.3 "stack party" + v1.3b "dress rehearsal" shipped — their codenames are now permanent. "dress rehearsal" re-opened the stage-metaphor lineage: a demo *is* a show, and v1.3b was the run-through that made it actually perform.)_

_Last updated: 2026-06-11 (**v1.4 removed** — no version currently staged; the former v1.4 seeds, incl. DEF-M10-01
cloud-store/S3-blobs, are unscheduled backlog, not a planned release. Prior: 2026-06-09 v1.3b "dress rehearsal"
SHIPPED, tag `v1.3.1`.)_
