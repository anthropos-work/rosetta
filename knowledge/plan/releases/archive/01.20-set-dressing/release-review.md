# Release Review: v1.2 "set dressing"

**Date:** 2026-06-06
**Milestones:** M9a (snapshot framework) · M9b (taxonomy) · M10 (Directus content) · M11 (recipes/presets/corpus)
**Verdict:** GREEN on scope / deferrals / code / tests / metrics / supply-chain. A handful of documentation fixes (2 must, 2 should, a few nice). No blocker.

## Scope (Phase 1) — GREEN
- [x] All 4 milestones delivered; the v1.2 thesis (taxonomy + content → **100% data-DNA coverage**, nothing waived) is structurally complete.
- [x] Fate-2 partitions (M9a's taxonomy/content/recipes Out-items) each landed in their target milestone. No Fate-3 routing needed.
- [x] Escape-hatch → v1.3: DEF-M10-01 (S3 media blob *bytes*) verified in `roadmap-vision.md` (cloud-store seed). DEF-M9a-D4 (cloud SnapshotStore backend) same v1.3 home.
- [x] Dropped: the offline pg_dump-FILE reader (M9b-D9, user) — cut + pinned-gone (`TestDroppedDumpFlagStaysGone`), not re-surfacing.
- [x] Unaccounted: NONE. 0 TODO/FIXME/HACK in `stack-snapshot` + `stack-seeding`.

## Deferrals (Phase 1b) — GREEN
- [x] `audit-deferrals --scope=release` GREEN — 1 open item (S3 blobs → v1.3, Fate-2, signed), 0 repeat/aged/chronic, 0 new from M11. Report: `audit-deferrals/deferral-audit-2026-06-06-release-close.md`.

## Supply chain (Phase 0) — GREEN
- [x] `go vet` clean on all 4 modules. Licenses all permissive (MIT / BSD-3 / Apache-2.0 / ISC); **no GPL/AGPL/LGPL**.
- [x] govulncheck: **0 called CVEs from any third-party dependency.** 12 called vulns are Go **stdlib** (go1.25.3), DoS/parsing-class, all fixed in go1.25.6–go1.25.11.
- [ ] [track] **Build/release with the go1.25.11+ toolchain** to clear the 12 stdlib findings (toolchain bump, not a code/dep change). Recorded in `dependencies.lock` + retro.
- [x] Lockfile written → `dependencies.lock`.

## Code Quality (Phase 2) — GREEN
- [x] gofmt + go vet clean (4 modules). No must-fix. Strong cross-milestone consistency (the M9b→M10 per-surface `PublicPredicate` firewall generalization keeps taxonomy byte-for-byte unchanged; defense-in-depth plan+post-capture firewall; no SQL-injection surface; `%w` error wrapping throughout).
- [ ] [should-fix → already tracked] Staged-ahead directus media/provision API (7 funcs, unit-tested, unreachable from any entrypoint) is the **intentional** DEF-M10-01 surface (blob bytes + per-stack provision, gated on S3 access) → v1.3. Confirm it's carried (it is). No code change.
- [ ] [nice-to-have] stack-seeding library API "deadcode" (9 funcs) is test-covered library surface beyond the CLI's current use — normal, not a defect.

## Tests & Benchmarks (Phase 4 / 4b) — GREEN
- [x] **708 Go test funcs** (693 test + 15 fuzz) — clerkenstein 218 · alignment 46 · stack-snapshot 212 · stack-seeding 232. All pass `-race`. **Flake 0** (2nd shuffled run of both v1.2 modules clean). Demo shell/py tooling 87 passed.
- [x] Metrics regression gate GREEN: test funcs **409 → 708 (+73%)**, no decrease; coverage up/flat on every continuing logic package (the single `pg` −2.7pp is live-DB-method denominator growth, no tests removed); flake 0. Aggregated → `metrics.json`; appended to `metrics-history.md`.

## Decision Consolidation (Phase 5) — CLEAN
- [x] All "blend-into-knowledge" decisions verified landed in the corpus. No cross-milestone contradictions (firewall generalization + the M10 "separate-store" self-correction are clean, not conflicts).
- [ ] [nice-to-have] A one-line methodology note in `snapshot-spec.md` ("toy/spike premises need prod re-validation before the real surface") — optional polish.

## Documentation (Phase 3) — 2 must-fix
- [ ] [must-fix] **F1** — `.claude/skills/db-query/SKILL.md` L177-179 still claims the M10-disproved "separate self-hosted Directus store / its own Postgres." Contradicts its own header + db-access.md + snapshot-spec.md. → rewrite to the `directus` schema in the same `postgres` DB (predicate `private=false AND tenant_id IS NULL AND status='published'`).
- [ ] [must-fix] **F2** — stray `</content>` tag at the tail of `corpus/ops/db-access.md` (L86) and `.claude/skills/db-query/SKILL.md` (L183). → delete the trailing line in both.
- [ ] [should-fix] **F3** — `README.md` status banner two releases stale (still "v1.0 shipped / v1.1 staged"). → v1.1 shipped + v1.2 complete (snapshot → 100% coverage).
- [ ] [should-fix] **F4** — `README.md` Quick Start omits `/demo-snapshot`. → add `/demo-snapshot N` between `/demo-up` and `/demo-seed`.
- [ ] [nice-to-have] **F5** — `roadmap.md` "In Development" narrative carries the pre-correction "separate Directus store" framing (the M10 close note records the correction; frozen planning record). → optional inline forward-note.
- [ ] [nice-to-have] **F6** — `corpus/architecture/README.md` alignment entry omits the data + snapshot-fidelity dimensions; index lists only 4 of ~9 docs.

## KB Consolidation (Phase 3b)
- [x] `corpus/ops/README.md` indexes both new docs; CLAUDE.md skill table + doc-locations updated; state.md/roadmap.md consistent (all 4 milestones done, v1.2 complete). No orphans, no oversized docs, all cross-refs + anchors resolve.
- [ ] [nice-to-have] `corpus/services/skiller.md` + `cms.md` carry no inbound pointer to the snapshot mechanism that now captures their schemas; the `corpus/ops/demo/` family isn't linked from the ops index. → add small cross-links.
</content>
