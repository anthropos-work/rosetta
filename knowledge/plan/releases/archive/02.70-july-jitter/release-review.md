# Release Review: v2.7 "july jitter"

**Date:** 2026-07-25
**Milestones:** M246, M247, M248, M249, M250, M251, M252, M253, M254 (9)
**Branch:** `release/02.70-july-jitter` · **Diff:** 175 files, +9388/−534 — **docs-only in rosetta** (133 `knowledge/plan` bookkeeping + 39 `corpus` + root); code-of-record in the rext repo via 17 `july-jitter-*` tags on origin. **0 platform-repo edits.**

## Blocking-gate verdicts
- **Phase 0 supply-chain:** clean — **0 new third-party deps** (no go.mod/go.sum/package.json changes); rosetta has no code manifests.
- **Phase 1b deferral audit:** **YELLOW, 0 hard blockers** — 0 repeat/chronic patterns. 5 items are "Fate-3 follow-up" from the *terminal* milestone → need explicit escape-hatch sign-off (Phase 9).
- **Phase 3b KB consolidation:** no split needed (largest changed doc `safety.md` 998 lines); indexes reflect contents; `services/README.md` = 27 docs verified.
- **Phase 4b metrics regression:** **GREEN** — every suite grew (Go 2010→2018, TS unit 257→290, Python 865→909), coverage flat/in-tolerance (seeders −0.4pp), flake 0, 0 new deps, 0 platform edits. New perf gates pass (studio FCP p95 817ms <1s; click→ACCESS 1.43/1.41s <5s).

## Scope ledger (Phase 1)
All 6 field defects + the full re-ground **delivered and live-proven on billion** (M254 gate a–h). **0 unaccounted, 0 Fate-3-undelivered.** 2 principled drops with decision records (D18 `participants_filter`; `DEF-M215-03(a)/F11`). Detail: `audit-deferrals/` report.

---

## Code Quality (Phase 2)
- [x] **[should-fix / rext / RUNG-ZERO] `demo-stack/tests/test_frontend_build.py:46`** — `_STUDIO_PATCHSET_MANIFESTS` 3→5. **LANDED** (tag `july-jitter-v27-close-rext-cleanup @ e61b604`, on origin); `test_frontend_build` now green, demo-stack 910/0-fail confirms.
- [x] **[nice / rext] gofmt** — the 6 v2.7-touched stack-seeding Go files reformatted (same tag). `gofmt -l` clean, `go vet`/`go build` clean (verified triple-clean). Residual: `playthroughs/manifest/hiring_isolation_test.go` cosmetic drift is **pre-existing (M225/v2.4)**, out of v2.7 scope — left as-is.

## Documentation (Phases 3 + 3b) — MUST-FIX
- [x] `CLAUDE.md:321` — "16 live Playthroughs … `playthroughs.md:105`" → **18** (`playthroughs.md:108`) + add M252's studio pair.
- [x] `corpus/ops/demo/README.md:206` — "16 live Playthroughs" → **18**.
- [x] `corpus/architecture/service_taxonomy.md:378` — "8 Go services" → **6** (contradicts its own `:55-62` table).
- [x] `corpus/ops/platform_repo.md:23` — "12 app service definitions" → **11** (vs `:59`).
- [x] `corpus/ops/platform_repo.md:44` — "5 migration repos" → **3** (vs `:92`).
- [x] `corpus/ops/staging-bringup.md:312` — delete the live `skillpath` migration-apply row (now folded into `app`'s `public`).
- [x] `corpus/ops/demo/demopatch-spec.md:52` — G5 "15 today" → **23** (vs `:62`/`:197`).

## Documentation — SHOULD-FIX
### Re-ground residuals (skillpath decommission / new domains)
- [x] `corpus/services/ai-readiness.md:185` — retired KPI ids (`avg_frequency/breadth/depth/context_fit`, "4 tiles") → the current 5 (`avg_adoption/transformation/originality/depth/ownership`, per `:533-536`). *(source-confirm vs `how_we_measure.go`.)*
- [x] `corpus/architecture/shared_libraries.md:77` — "12 Connect-RPC" enumerates 11 → **add `LabSessionService`** (the 12th; documented in new `ai-labs.md:100` + `backend.md:151`) — makes "12" correct.
- [x] `corpus/ops/platform_repo.md:11` — "~13 sibling repos" → **10** (vs `skillpath.md:39` + `roadrunner.md:14`; reconcile `repos.yml`).
- [x] `corpus/services/jobsimulation.md:63`/`:79` — Roadrunner listed live → mark orphaned (vs `roadrunner.md:5`).
- [x] `corpus/services/clerk-integration.md:99` — studio gating `org:admin` → also `content_creator` (`STUDIO_ACCESS_ROLES`).
- [x] `corpus/ops/setup_guide.md:302` — add the skillpath decommission note (row deleted above, note silent).
- [x] `corpus/ops/staging-sync.md:38` — "15 repos" → **14** (vs `:99`).
- [x] `corpus/ops/staging-sync.md:192` — "skiller subgraph" → merged app/backend subgraph.
- [x] `corpus/ops/staging-bringup.md:40` — "15 service-repo sibling clones" → correct count.
- [x] `corpus/ops/staging-bringup.md:246` — move `skillpath` to the legacy-schema clause.
- [x] `corpus/ops/staging-bringup.md:371` — "all 15 services" → **14**.
- [x] `corpus/ops/staging-bringup.md:598` — drop the "skillpath subgraph 422" clause (keep the `language`-arg drift).
- [x] `corpus/ops/secrets-spec.md:131` — "8 of 9 Go repos" → re-verify (likely 7 of 8). *(Resolved to 7 of 8 — skillpath was a Go repo shipping no `.env.example`, so its removal drops both counts by one.)*
- [x] `corpus/ops/seeding-spec.md` — add the seeder re-point to `public.skill_path_sessions`.
### Field-defect doc seams (M248–M253)
- [x] `corpus/ops/demo/frontend-tier.md:181` — blanket "studio-desk needs no demopatch" → scope to auth (5 studio patches exist).
- [x] `corpus/ops/demo/frontend-tier.md:294` — "THREE" studio source patches → **FIVE**.
- [x] `corpus/ops/demo/frontend-tier.md:489` — "All three ant-academy patches" → **five**. *(Also fixed the mirror drift at `demopatch-spec.md:205` "The 3 ant-academy patches" → 5.)*
- [x] `corpus/ops/demo/cockpit-spec.md:226` — "four demo-patches" → **three** (shared next-web covers 4 menus).
- [x] `corpus/ops/demo/content-stories-spec.md:109` — manager-surface cell → the M248 `/sim/<slug>/<userId>/result/<sessionId>` route (interview excepted).
- [x] `corpus/ops/demo/content-stories-routes.md:67-69`/`:168-169` — reconcile matrix cells to the M248 route.
- [x] `corpus/ops/demo/content-stories-routes.md:133`/`:421` — clarify mirror-row scope (M248 non-interview reads jobsim by `sessionId`; mirror still needed for org scoreboards).
- [x] skillpath-schema residuals → `public.skill_path_sessions`: `content-stories-routes.md:154`/`:379`, `content-stories-spec.md:294`.
- [x] `corpus/ops/demo/latency-budget.md:381-384` — studio FCP "chartered to M254 / local only" → fold in the billion p50 637–726 ms.
- [x] `CLAUDE.md` skills table — omits `dev-for-dummies` (on disk) → add a row or document the exclusion.

## Documentation — NICE-TO-HAVE
- [x] `ai-readiness.md:546` note M254 re-proof; `demopatch-spec.md:92` "(v2.6 count)"; `latency-budget.md:52-53` "(later 49 at M241)"; `platform_repo.md:104` skillpath RPC note; `roadrunner.md:144` "historical (severed) consumer"; stale stamps (`secrets-spec.md:393` 56→61-gene, `setup_guide.md:286`, `staging_from_dump.md:323`, `update_guide.md:161`).
- [x] **Release-level notes:** (C1) a "mirrored-count discipline" note — 3 count-drifts this release (patch-inventory shipped RED, playthroughs, KPI tiles); any count mirrored in >1 doc or backed by a fence must move with all mirrors + its fence in one commit. (C2) studio-desk graduated to a first-class demopatch target (M249 ladder + M252 wiring + M253 pair). *(Both added to `demopatch-spec.md` §5.)*
- [ ] `roadmap.md:138`/`:154` "IN DEVELOPMENT" → resolved by Phase 10 rotation (expected pre-rotation state). *(Left as-is — this is expected pre-rotation state, resolved by Phase 10, not a Phase 7 doc fix.)*

## Decision Consolidation (Phase 5)
- Blends landed + coherent (M248/M249/M250/M252/M253). B1 (demopatch doc-vs-fence RED across M253→M254) resolved by M254. No cross-milestone contradictions on content-stories link shape / AI-readiness seeder handoff / demo defaults.

---

## Escape-hatch items — need explicit fate at Phase 9 (the YELLOW driver)
Terminal-milestone Fate-3 follow-ups with no in-release home:
1. **FIX-M254-c-academy-durable** — native academy loses Back-to-Cockpit on a *long-running* demo (fresh demo renders fine; gate (c) MET). rext `stack-injection` reapply lifecycle.
2. **FIX-M254-g-testhealth** — 6–7 host-sensitive demo-stack test-harness tests, **0 demo-runtime impact** (real academy serves 200). 2/8-class already fixed live.
3. **(b)-voice `manager_presence_only`** re-seed — honesty fix (denominator 47→45 + flag); underlying blocker is accepted vision `DEF-M240-01` (Bunny-keyless box).
4. **`verify.sh` stale-`skillpath` default** — 1-line rext `stack-verify` edit (still present; not gate-blocking). *Candidate to LAND-NOW.*
5. **rext-hygiene inert set** — dormant `INJECTED["skillpath"]` + guard `skillpath:8095` + fixtures; aged-out at M247; **0 functional impact**. *Candidate to LAND-NOW or standing-note.*
