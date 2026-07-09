# Release Retro: v2.1 "quick change"

**Shipped:** 2026-07-09 · **Milestones:** M208 → M209 → M210 → M211 (strictly sequential) · **Tag:** `v2.1`
**Theme:** the skiller-in-app re-ground — re-fit the rext tooling + corpus + stacks to a landed platform structural change (skiller service + schema merged into `app`/`public`, 4 subgraphs, RPC→backend), and prove `/dev-up` + `/demo-up` still work cold. **Zero platform-repo edits.**

## Incidents this release
- **No P0/P1/P2. No regressions. No flakes.** The M211 iterative loop *surfaced and fixed* several real merged-platform bring-up bugs — a stale-image mishmash (a pre-merge router image with a dead skiller subgraph), a **pinned build-scratch** that survived `--purge` (the root cause of "cold" demos shipping pre-merge binaries), a missing dev casbin-policy load, and a reset-to-seed roster gap that failed all 10 playthroughs. These were the milestone's *work*, not incidents against a committed deliverable — the bring-up-acceptance gate exists precisely to drive them out, and it did.

## Cross-milestone patterns
- **The stale-clone / re-sync-at-use class recurred** (M208 stack re-sync + M211 build-scratch pinning). Durable lesson captured in `snapshot-spec.md` ("build-scratch freshness invariant"): the fix is always *re-sync at use*, never *trust a prior clone*. Worth a standing guard if it recurs a third time.
- **Cache-migration-as-recapture** — the substantive M211 mechanism (re-key a captured cache `skiller.*→public.*` when a merge preserves data + table names, no prod access) — now a durable corpus note (`snapshot-spec.md`). The sanctioned no-prod-capture-source path for a pure schema-prefix move.
- **rext-README reconciliations pile at the code-of-record roll** — three separate items (M209 + M211) route to the same close-release rext roll because rext is frozen per-milestone at its tag. Handled at the v2.1 roll (TEST-1 + DOC-1).

## Close-release findings (all should-fix / nice-to-have; 0 blocking)
- The one real cross-milestone contradiction: M210 flipped the schema token but left `42,763` as the current count in 3 tooling docs vs M208's authoritative `42,790` — **fixed**.
- `stack-snapshot/SKILL.md` self-contradiction (replay-target `skiller` vs its own merge banner) — **fixed**.
- Supply-chain: `go1.25.11→go1.25.12` cleared two inherited stdlib advisories (govulncheck clean all 6) — **fixed**.
- `stack-verify/e2e` missing typescript/tsconfig; 2 missed skiller live-service refs (`gen_injected_override.py`, `run.sh`); doc-comment double-words; classifier dedup — **all fixed**.

## Carry-forward
- **TEST-1** (rext `stack-seeding/README` test-count drift) + **DOC-1** (rext `dev-stack/README` index `migrate-dev.sh`) → reconciled at the v2.1 rext code-of-record roll (Phase 10).
- **CAVEAT-1** (clean-box literal full destructive `/dev-up`) → unscheduled backlog (`roadmap-vision.md`); the dev-cold gate delta was proven at the DB-init level on a non-destructive throwaway to protect the user's native-app content-line dev box. Belt-and-suspenders only.
- Standing backlog (unchanged): DEF-M10-01 (cloud store / S3 blob), DEF-M21-01 (replayCmd hermetic test), M314b (prod frozen-read hydration — prod-team, out of tooling scope).

## Metrics delta (vs v2.0)
- rext Go test funcs **1745 → 1764** (+19, every module flat-or-up); TS unit **100 → 103** (+3); Playthroughs **10/11 GREEN re-proven on the merged platform**; flake **0** (triple-clean 3/3); **0 net-new deps**; alignment **100%/100%** held. See `metrics.json`.
- **Housekeeping gap noticed:** `knowledge/plan/metrics-history.md` was missing the v1.10b + v2.0 rows (their close-releases skipped the append). v2.1's row is added; the v1.10b/v2.0 numbers live in their archived `metrics.json`. A backfill of those two rows is a low-priority hygiene follow-up.

## Stats delta reference
- Phase 8c project-stats snapshot: `knowledge/journal/stats/2026-07-09.json` (git velocity: 719 commits, +51 since the v2.0 close, ~4.3/day over 168 project-days). The code/docs/test cells read 0 — a known `stats.sh` layout-mismatch (rosetta docs live under `corpus/`, the rext code is a gitignored sibling repo); the authoritative v2.1 counts are in this release's `metrics.json`. Fixing the stats auto-detect for the two-repo layout is a separate tooling follow-up.
