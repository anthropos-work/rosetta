# Release Retro: v1.2 "set dressing"

**Shipped:** 2026-06-07 · tag `v1.2` · **Milestones:** M9a → M9b → M10 → M11 (all `section`)
**Theme:** the snapshot mechanism — capture the **public** reference library (skills taxonomy + Directus content) read-only from prod, replay per-stack, measured-faithful, lifting M7c's two `waived` surfaces to **100% data-DNA coverage**. Demo worlds get believable props; production is never touched.

## What shipped
A dedicated `rosetta-extensions/stack-snapshot/` extension: capture → `.agentspace` manifest-cache → per-stack replay, behind a tested **tenant-data firewall** (`AssertPublicOnly`, plan-time + post-capture, nothing persisted on a leak) and a **production-safe capture-source policy** (prod-`pg_dump` ingest → safe throttled `marco_read` primary read; restore/replica = v1.3 upgrades). Two real surfaces: the **public taxonomy** (M9b — 42.8K skills + embeddings + translations, pgvector index rebuilt on replay) and the **public Directus content** (M10 — 304 published sims / ~190 paths, multi-level parent-scoped, per-stack Directus store). Plus the `/db-query` port, the snapshot-fidelity data-DNA dimension, the `/demo-snapshot` skill, and the set-dressed recipe family.

## Incidents (P1) across the release
- **M9a framework gap surfaced by the first real surface (M9b-D2).** The framework recorded a parent name but applied an **empty** capture filter to column-less tables (embeddings/translations) — safe only for M9a's toy surface; on the real taxonomy it would have captured customer-parented rows. M9b closed it with a real `fk IN (SELECT id FROM <parent> WHERE org_id IS NULL)` predicate + a parent-aware leak probe. *Caught at build, fixed inline with regression + fuzz.*
- **Spike inference disproved by prod (M10-D2 / KB-1).** The M10 spike inferred a "separate self-hosted Directus store, its own Postgres"; a DB-side `marco_read` check found the content is the **`directus` schema of the same prod DB** — self-resolving the capture credential (nothing for the user to provide). *Corrected in-milestone + propagated to the corpus; a stale claim still leaked into the `/db-query` skill body and was caught at release-close (F1).*
- **M9b close blocked (legitimately) on a repeat-defer.** The offline pg_dump-FILE reader was an M9a→M9b aged-out carry-forward; close attempt 1 went RED (correct). The **user DROPPED it** (M9b-D9) — it added no capability over restore-then-`--dsn` and no reliable speed gain; docs/code were aligned to the truth and the `--dump` flag pinned-gone.
- **M10 close sub-agent returned prematurely.** Close attempt 1 ran only the deferral audit (GREEN) and returned without merging/tagging — a `/developer-kit:work-milestone` close-sub-agent reliability bug. **Post-flight caught it** (no merge/tag landed, tree clean); a re-run completed the full close. *No work lost; flagged as a tooling observation.*

## Cross-milestone patterns
- **"Toy/spike premises need prod re-validation before the real surface."** Recurred twice (M9b-D2 empty-filter; M10-D2 separate-store) — both caught at build/spike and corrected cleanly. The first *real* surface is what validates a framework's assumptions; budget for it.
- **Extensions tagging discipline.** Each close sets `stack-snapshot-mN` at the **hardened** HEAD (M9a caught a tag trailing HEAD by the harden commits and re-pointed it; the pattern held cleanly through M9b/M10/M11).
- **Post-flight verification earned its keep.** The orchestrator's branch/HEAD/tree post-flight check caught the M10 premature-return that the sub-agent's own "clear" report missed — trust-but-verify is load-bearing for autonomous pipelines.

## Carry-forward → v1.3 (all in `roadmap-vision.md`)
- **DEF-M10-01 — S3 media blob *bytes*.** The directus media/provision API (built + unit-tested, unreachable from any entrypoint) is gated on S3-read access not wired here; refs + structure shipped (the floor). Fate-2 → the cloud-store seed.
- **Cloud snapshot store (DEF-M9a-D4).** The `SnapshotStore` interface seam is built; the cloud/S3 backend is the v1.3 swap.
- **go1.25.11+ toolchain bump (supply-chain hygiene).** Clears the 12 Go-stdlib govulncheck findings (DoS-class); not a code/dep change. Recorded in `dependencies.lock`.
- (Also parked at design: AI-generated content, external shareability — v1.3 seeds.)

## Metrics delta (vs v1.1)
- Go test funcs **409 → 708 (+73%)**; flake **0**; both v1.2 modules `-race` + triple-clean. Coverage up/flat on every continuing logic package (the single `pg` −2.7pp is live-DB-method denominator growth, no tests removed).
- Data-DNA coverage **100% of the full catalog** (nothing waived — the thesis complete).
- Supply chain: **0 third-party CVEs**, all-permissive licenses.

## Stats delta
Phase 8c `/developer-kit:project-stats` snapshot saved at `knowledge/journal/stats/2026-06-07.json` (release tip; first snapshot — no prior to diff). **Caveat:** the script swept the gitignored cloned platform/extensions repos under `stack-dev/` + `.agentspace/`, so its `Source ~2M lines` figure is dominated by *cloned* code, not rosetta's own corpus; it also read `Docs 0` (layout-detection miss on rosetta's `knowledge/`/`corpus/` split). The accurate v1.2 engineering numbers are in this release's `metrics.json` (708 Go test funcs across the 4 extensions modules; rosetta itself tracks zero `.go`). Git velocity (205 commits, 136-day project age, 18/18 milestones done across v1.0–v1.2) is sound.
</content>
