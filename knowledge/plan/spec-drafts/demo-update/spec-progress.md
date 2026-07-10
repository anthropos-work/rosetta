# `/demo-update` — Spec Progress (open points tracker)

> **Status:** Draft · spec-draft · 2026-07-10 (tracks [`spec.md`](spec.md) `v0.1`)
> Tracker + decision log for the spec, [`spec.md`](spec.md). Decisions are recorded one at a time; open
> questions surface in the "Open" rows and either resolve here or move to
> [`next-release.md`](next-release.md).

**Legend:** 🔴 not decided · 🟡 discussing / proposed · ✅ decided · ⏭️ deferred (→ [`next-release.md`](next-release.md))

| # | Topic | Status | Decision |
|---|---|:------:|---|
| A | What is the capability? | ✅ | `/demo-update <N>` — refresh a running `demo-N` to prod-latest **in place** (no teardown-reseed), pulling in new features + populating them via existing seeders + gated by a fatal verify. [`spec.md`](spec.md) §1. |
| B | Extend `/stack-update` or build new? | ✅ | **Build new peer.** `/stack-update` stays dev-only + as-is. Different target, source model, verification bar, data-preservation contract. [`spec.md`](spec.md) §1.2. |
| C | What "prod-latest" means concretely | ✅ | **Highest semver `v*` tag** per repo (reuses `up-injected.sh`'s existing rule + `demo-stack/lib/clone_repos.py:pick_latest`). Per-repo `--ref svc=ref` overrides via `resolve_ref()`. [`spec.md`](spec.md) §1.1, §3.2. |
| D | Ref recording semantics | ✅ | **SHA-pin the record**, tag-pick the resolution. A mid-update tag move never rewrites history. [`spec.md`](spec.md) P2, §3.2. |
| E | Data preservation contract | ✅ | Data survives every update path. `--reset` is double-gated (`--force-reset` + `--force-reset-confirm=demo-N`) + N=0-guarded. Additive seeders only in the default flow. [`spec.md`](spec.md) P1, §3.7. |
| F | Feature-parity mechanism | ✅ | `coverage-manifest.ts` (in rext) is the declarative section catalog. Post-update coverage sweep (T3) detects present-but-empty sections. Rext-bump moves the catalog forward. [`spec.md`](spec.md) §2.1. |
| G | Synthetic-data reuse policy | ✅ | Reuse `stacksnap replay` + `stackseed --additive` + optional `gen-batch`. One net-new artifact: `update-routing.yaml` (section→seeder). Layout test keeps it honest. [`spec.md`](spec.md) §2.2. |
| H | Verify gate contract | ✅ | **3 tiers, fatal-by-default** (T1 autoverify, T2 test-platform live, T3 coverage sweep). `--no-verify` skips T3 only. Optional T4 = playthroughs. [`spec.md`](spec.md) §2.3, P4. |
| I | Behavior on verify failure | ✅ | **Demo stays UP**; print diff + suggested remediation; offer `--rollback`. Never auto-roll-back on T3 failure. [`spec.md`](spec.md) §2.3, §3.9. |
| J | Rollback semantics | ✅ | `--rollback [op_id]` re-checks pre-update SHAs + re-injects + rebuilds. **Schema NOT rolled back** (limitation printed every rollback). `--snapshot-db-first` = presenter insurance. [`spec.md`](spec.md) §3.9. |
| K | Where the CLI lives | ✅ | `rosetta-demo update N [flags]` (new verb on the existing lifecycle CLI) + `rosetta-extensions/demo-stack/update-injected.sh` (mirrors `up-injected.sh` phasing). [`spec.md`](spec.md) §3. |
| L | Where the skill lives | ✅ | `.claude/skills/demo-update/SKILL.md`; guide at `corpus/ops/demo-update.md`; CLAUDE.md skill table entry. Added in M-E. [`spec.md`](spec.md) §5 M-E. |
| M | Milestone breakdown | ✅ | **M-A** code refresh + rolling rebuild + migrate → **M-B** data refresh → **M-C** fatal verify gate → **M-D** section→seeder routing → **M-E** rollback + logs + docs + skill. [`spec.md`](spec.md) §5. |
| N | Close-on-gate for M-A | ✅ | `test-platform N live` GREEN on Ithaca demo-N post-update, `docker inspect` + `git rev-parse HEAD` prove the new SHAs. [`spec.md`](spec.md) §5 M-A. |
| O | Cost enforcement for `gen-batch` | ✅ | `--max-cost-usd` from the manifest is re-asserted; `--reconfirm-max-cost` required for non-cache-hit runs. [`spec.md`](spec.md) P7, §3.7. |
| P | Secrets handling | ✅ | Values-blind; `--reprovision` delegates to `/stack-secrets`, never duplicates. [`spec.md`](spec.md) P6, §1.3. |

## Open questions

| Q | Topic | Status | Note |
|---|---|:------:|---|
| Q1 | Does the routing-coverage layout test need a schedule or only rext-bump trigger? | 🟡 | Rext-bump triggers naturally when a new coverage-manifest section lands. A schedule feels redundant. Lean: **rext-bump only.** Confirm at M-D. |
| Q2 | Snapshot-then-resolve vs resolve-then-freeze at Phase 2 | 🟡 | The SHA-pin record (P2) makes this moot for history; the *plan-vs-execute* mismatch window is what's at stake. Lean: **resolve-then-freeze** (record SHAs at Phase 2, ignore later tag moves). Confirm at M-A. |
| Q3 | Auto-require `--reconfirm-max-cost` on mother-prompt digest change | 🟡 | Yes — that's the honest signal. Implementation detail for M-B. |
| Q4 | Cross-N update lock granularity | 🟡 | Docker network + rolling-restart Port coordination may need a box-wide lock. Lean: **coarse advisory lock** on `.agentspace/updates/box.lock` in addition to the per-N lock. Confirm at M-A. |
| Q5 | Phase-4 rolling-restart checkpointing for deterministic rollback | 🟡 | Per-service checkpoint file. Confirm at M-A. |

## Deferred

| # | Item | → |
|---|---|---|
| D1 | Continuous / scheduled `demo-update` (cron) | [`next-release.md`](next-release.md) |
| D2 | Multi-`N` batch update | [`next-release.md`](next-release.md) |
| D3 | Update from `demo-down` state (would require a resurrect verb first) | [`next-release.md`](next-release.md) |
| D4 | Full schema rollback (auto pg_restore) | [`next-release.md`](next-release.md) |
| D5 | `--with-playthroughs` as default | [`next-release.md`](next-release.md) |

## Decision log

- **2026-07-10** — All rows A..P decided in initial draft. Q1..Q5 opened for milestone-time resolution.
- **2026-07-10 (M-B live-run)** — Closure gate fired RED on demo-1 against two pre-existing substrate defects — both **out of the `/demo-update` lane**, both recorded durably in [`known-findings.md`](known-findings.md):
  - **Finding A** (stacksnap): taxonomy capture stale vs post-seed FK indexes (schema fingerprint drift). Not user-visible. Fix in `rosetta-extensions/stack-snapshot`.
  - **Finding B** (stack-seeding): `stories.seed.yaml` gen-batch mints fabricated `K-*` verified-skill node-ids instead of resolving via `TaxonomyRefs`. **User-visible on demo-1 today** — 168 `is_verified=true` rows on `user_skills` + 113 distinct fabricated IDs on `membership_skills` (3 522 rows) render on Skill Spotlight / org-workforce / member listings. Fix in `rosetta-extensions/stack-seeding` (`stories.seed.yaml` preset + `GeneratedBatchSeeder` gen-batch path). Prioritise independently of `/demo-update`.
  - **`/demo-update` decision:** accept RED as spec P1 validation. Gate is doing its job. Ship M-B and move to M-C.
- **2026-07-10 (M-C live-run)** — Verify gate (T1 + T2 + T3) landed on `agent/demo-update` at commit `1618c6b`, hardened at `6bc207f`. Live-run against demo-1 with `--no-snapshot --no-seed` (to reach Phase 8 past the M-B RED):
  - **T1 (`phase8_t1_autoverify`)** GREEN — backend `/api/health` 200 on :18082 (after 1s readiness wait), `sentinel.casbin_rules = 1150`.
  - **T2 (`phase8_t2_scoped_live`)** RED — surfaced **Finding C**: `demo-1-directus-1` is in `Exited(1)` since the original `/demo-up 1` bring-up, so `probe_postgres_schemas` correctly expects a `directus` schema that was never created. Same pattern as A/B: pre-existing substrate defect, not in the `/demo-update` lane. Recorded in [`known-findings.md`](known-findings.md) with owner surface (`demo-stack` bring-up-side + operator-side box remediation) and the "not user-visible in the content sense" note (demo-1 falls back to prod-read content without local Directus).
  - **T3 (`phase8_t3_coverage_sweep`)** not reached (T2's `exit 1` halted the flow, per the fatal-always contract of §2.3).
  - **In-lane hardening on M-C (commit `6bc207f`):** T1 bounded 20×1s readiness retry to tolerate Phase 5 migrate-restart; T2 enumerates the RUNNING compose services via `docker ps --filter label=com.docker.compose.project=demo-N` and passes `STACK_SERVICES` to `verify.sh` so its M18 scope filter is honoured. Neither loosens the fatal semantic.
  - **`/demo-update` decision:** accept RED as spec §2.3 validation. Gate mechanics proven (fatal-always T1 gated by readiness; fatal-always T2 scope-aware; T3 unit-tested but not live-exercised on this substrate). Ship M-C and move to M-D.
