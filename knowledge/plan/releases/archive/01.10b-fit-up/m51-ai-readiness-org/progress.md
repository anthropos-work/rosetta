# M51 — progress

## Running ledger
_Appended after each iter (tik/tok). Each entry: what was attempted, what moved, gate distance._

- iter-01 (tok bootstrap): KB-fidelity gate GREEN (AI-readiness contract verified vs app code); authored TOK-01 (active-cycle signals-true additive-to-stories seed); 3 doc fixes + seeder survey — see iter-01/progress.md
- iter-02 (tik): strand 1 — appended the 3rd AI-readiness story (Northwind, 200, narrative:ai-readiness, hero trio) + net-new OrgSettingsSeeder (the organization_settings ai_readiness enablement gate); seeded+verified in demo-1 (gate scoped to Northwind only), roster/cockpit re-exported. Baseline manager-vantage sweep = (failing=6, escapes=0), persona OK, frontier-exhausted 52pp — the gate's starting point. closed-fixed. See iter-02/progress.md
- iter-03 (tik): strands 2+3 — the ai_readiness_* CONFIG seeder (active cycle + 3 steps + 8 real-node-id AI skills + 2 sims) + the 200-member signals-true FUNNEL seeder (78.4% all-3-complete, Aria→stage3 COMPLETED / Ben→stage1 STARTED / Dana excluded); adopted from a stopped parallel agent, verified-green (7 new unit tests + full suite), demo-1 re-seeded CLEANLY to a known reproducible state. GATED manager sweep held at (failing=6, escapes=0) UNCHANGED — the 6 are an UNRELATED cluster (base-Workforce org-scale PERF-WALL skeleton false-fails: data-in-DB, real chrome + skeleton rows + 11.4s GraphQL), a demo-UP/consumed-tag fix not a seed fix; + the manager manifest doesn't yet ASSERT the AI-readiness funnel. closed-fixed-partial; perf-wall + manifest routed to iter-04. See iter-03/progress.md
- iter-04 (tik, run-2 RESUME after the cleared consumption-clone blocker): re-up demo-1 at fit-up-m50 (all 3 M46/M50 perf demo-patches baked: pagination + targetRole-authz-skip + FK indexes, verified) + re-seed the AI-readiness showcase org from the authoring copy (Northwind 200, ENABLED, 78.4% all-3, heroes Aria→3/Ben→1/Dana→0) + re-export the 9-hero roster/cockpit. GATED manager sweep HELD at (failing=6, escapes=0) frontier-exhausted — the m50 perf-patches reduced the wall 76.4s→~11.6s but the residual COLD members-grid query still exceeds the harness measurement budget → the 6 skeleton false-fails persist (data-in-DB, real-chrome-over-skeleton, slow-not-erroring). HYPOTHESIS "m50 patches alone clear all 6" FALSIFIED. closed-no-lift; residual perf (harness warm/poll deepening) + manifest AI-readiness assertion + cockpit jump_to routed to iter-05. See iter-04/progress.md
- iter-05 (tik, tooling-iter): the content-presence `warmHeavyGrids` (real-rows warm replacing the 4s networkidle bail for the org-scale manager grids) + TOK-01 strand-4 landed — the gate-proving AI-readiness manifest descriptor (`/enterprise/workforce/ai-readiness` org-score + 3-step funnel assertions + seedPath prime), the cockpit `ai-readiness` deep-link + Dana `jump_to` re-point, and the org-agnostic corporate-email regex CORRECTING the shared manifest's hardcoded `cervato-systems.com` (false-failed Northwind). GATED manager sweep moved (failing=6→5, escapes=0), reachable 49→65, persona GREEN — clearing the stale-Sentinel-policy reachability root cause. closed-fixed-partial; residual 3 perf-wall grids + the 2 cervato false-fails (resolved by the committed manifest correction) + bake-Sentinel-reload-after-seed routed to iter-06. See iter-05/progress.md
- iter-06 (tik): re-swept with the corrected iter-05 harness — the org-agnostic email fix CLEARED the 2 cervato false-fails; deepened warmHeavyGrids (added the ai-readiness route + skeletons-cleared hydration) landed clean but did NOT lift the metric (5→5). Root-caused the residual: a PLATFORM-side AI-readiness response-build perf wall — GET /api/workforce/ai-readiness NEVER completes in-budget (0 backend completions; refresh-worker deadline-exceeded), all SQL ms-fast, the cause is the live-recompute + a per-skill federated translation N+1 (withSkillerLang→skiller _entities), same class as the M46 per-object RPC. closed-no-lift; USER-BLOCKER: the gate can't reach (0,0) without a platform-read-path edit OR a substantial new app-translation demo-patch OR a disclosed-residual sign-off — surfaced for the user's decision. Protocol lesson added (never-completing-GET vs slow-paint). See iter-06/progress.md
- iter-07 (tik): applied the user-chosen M48 CLOSED-CYCLE strategy — cycle active→closed + frozen ai_readiness_snapshots per member (platform-model-scored, heroes preserved) + the 8 ai_readiness_* tables added to stackseed --reset + baked --reload-sentinel (FIX-M51-iter07). Re-seeded demo-1 CLEANLY: DB VERIFIED closed cycle + 199 frozen snapshots (78.4% stage-3, Aria=3/Ben=1/Dana=none). GATED manager sweep HELD at (failing=5, escapes=0) frontier-exhausted — HYPOTHESIS PARTIALLY FALSIFIED: the platform FE's DEFAULT dashboard GET fires the LIVE path (no ?cycle=, never completes) + never fires /cycles, so the frozen path (buildResponseFromSnapshots, which EXISTS + is fast) is never taken. Closing the gap is PLATFORM-bound. closed-no-lift; USER-BLOCKER surfaced (a: disclosed-presenter-note [needs sign-off] / b: escalate platform edit / c: new app read-path demo-patch [not chosen]). See iter-07/progress.md
- iter-08 (tik, tooling-iter): applied the user-chosen run-5 DEEP-LINK strategy — resolved the Northwind closed cycle id (95d9fc3d…c97; DB: 199 frozen snaps, 78.4% s3, Aria s3/Ben s1/Dana none) + confirmed the FE reads ?cycle= from the URL. But the mandated CHEAP PROBE FIRST (authed dual-endpoint DIRECT probe, probe-aireadiness-deeplink.spec.ts) FALSIFIED the premise: /cycles is fast (40ms) but the frozen data GET ?cycle=<closed> NEVER completes (180s timeout) — buildResponseFromSnapshots re-loads the WHOLE org's members (loadMembers, hydrateMembers(nil,nil)) → the SAME org-scale wall, NOT the demo-patchable targetRole authz RPC (jobRole is a SQL column). So the deep-link is INERT — no cockpit/manifest edit made (would ship a hanging ?cycle=), no sweep run (would reconfirm 5). closed-no-lift; USER-BLOCKER surfaced (deep-link FALSIFIED — stronger than iter-07's; residual reachable only via a platform edit [bound loadMembers in the snapshot path / frozen_tags column] OR the disclosed-presenter-note which needs the user's EXPLICIT sign-off). Protocol lesson added (frozen SCORES ≠ frozen RESPONSE; measure the fast branch end-to-end). See iter-08/progress.md
- iter-09 (tik, tooling-iter): app-aireadiness-snapshot-loadmembers demo-patch bounds the frozen-read member hydration (loadMembers whole-org → loadMembersByUserIDs over the ~199 snapshot users) → the frozen ?cycle=<closed> GET 180s-timeout → 19ms, dashboard renders the full funnel. + a coverage-manifest funnel-descriptor correction (the header is one of two mutually-exclusive strings). GATED manager sweep 5 → 0, escapes 0, persona 0, frontier EXHAUSTED, on a fresh demo-up. GATE MET. rext tagged fit-up-m51. closed-fixed. — see iter-09/progress.md

## M51: Final Review

_Close review (2026-07-01). Phases 1–5 ran as parallel scans (deferral audit blocking-gate; code-quality + test-coverage in the rext authoring copy; docs review in a parallel agent). Iterative shape → Phase 9-iter (Gate Outcome Ledger). Default fix-everything. One BLOCKING deferral surfaced (academy F6 repeat-defer — user fate decision needed)._

### Scope
- [x] **S1 [CRITICAL] `delivers` commitment UNMET** — `stories-spec.md` has no 3rd AI-readiness story → author the Northwind/Aria·Ben·Dana section + the 3 seeders + the closed-cycle frozen-snapshot model (Fate-1).
- [x] **S2 [BLOCKING] Academy F6 repeat-defer** (M50 Fate-3 → M51, not done; gate MET without it) → **RESOLVED: user fated LAND-NEXT → M53** (Fate-3 override of M53's `No new feature code` guard — the cold rebuild is the natural place to seed + verify academy content). Recorded in `decisions.md` (D-CLOSE-1) + `m53-cold-rebuild-acceptance/overview.md` Scope `In:`; audit RED cleared. See `audit-deferrals/deferral-audit-2026-07-01-m51-close.md`.

### Code Quality
- [x] [should-fix] C1 — `cmd/stackseed/main.go` `resetTables`: guard the 2 platform-only AI-readiness tables (`ai_readiness_live_snapshots`, `ai_readiness_diagnose_narratives`, which no seeder writes) with `to_regclass` so `--reset` doesn't hard-abort on a stack whose app schema predates them.
- [x] [should-fix→accept] C2 — funnel `flush` is non-transactional (a package-wide pattern; deterministic-id re-seed heals partial state). Document the re-run-idempotency reliance in a code comment; not a per-seeder-tx rewrite (matches the population_evidence sibling).
- [x] [nice] C3 — stale `keepStartedMembers` comment symbol in `ai_readiness_funnel.go` (no such local symbol) → re-point the comment at the real gate.
- [x] [nice] C4 — the config↔funnel sim-ref value co-derivation was duplicated inline (a desync hazard the ordering-contract test doesn't cover) → **extracted the shared `aiReadinessSimRefs(pool, org)` helper** (both seeders call it → co-derive by construction, can't drift) + a co-derivation invariant test asserting the config's persisted `ai_readiness_sims.sim_ref` equals what the funnel derives (populated pool + free-fallback). Fixed the C1 build-bug (`Conn.QueryRow` signature) landed with it.
- [x] [nice] C5 — LIKE-metachar fragility / hand-rolled JSON encoder / reload-RPC observability: document as accepted (code-owned inputs, correct for the domain) in the review record; no change.

### Documentation
- [x] D1 — `stories-spec.md`: author the 3rd AI-readiness story section (= S1).
- [x] D2 — `seeding-spec.md`: add the M51 seeder-inventory + Status entry (3 new seeders, `fit-up-m51`).
- [x] D3 — `services/ai-readiness.md`: add the iter-08/09 `loadMembers` whole-org-hydration finding + M314b + record that M51 SHIPPED the closed-cycle/frozen strategy (not active/signals).
- [x] D4 — `CLAUDE.md`: extend the stories-spec index entry past M41 to the M51 AI-readiness 3rd story.
- [x] D5 — `stories-spec.md`: correct the stale "locked 2-stories × 3-heroes" claim → 3 stories.

### Tests & Benchmarks
- [x] T1 — add a resolved-DAG integration assertion that `ai-readiness-config` precedes `ai-readiness-funnel` in the rendered plan (currently pinned only at the metadata-contract level).
- [x] T2 — add a no-browser unit test for `section-assert.ts` (the new ~21KB TS lib currently exercised only by the live-demo sweep).

### Decision Triage
- [x] TOK-01/TOK-02 (active→closed→deep-link→app-patch strategy arc) + the `app-aireadiness-snapshot-loadmembers` demo-patch model + the "frozen SCORES ≠ frozen RESPONSE" loadMembers finding → already blended into `coverage-protocol.md` (iter-08/09 lessons) + being blended into `stories-spec.md`/`ai-readiness.md` (D1/D3). Add `(#M51-*)` ref tags where blended. The USER-BLOCKER decisions (iter-06/07/08) stay archived in `decisions.md` (maintainer-only history).

## M51: Gate Outcome Ledger (iterative — Phase 9-iter)

**Close status: `closed-on-gate`.** Gate met at iter-09; no `carry-forward.md` (not closed-incomplete).

### Gate
| Field | Value |
|---|---|
| **Metric** | `(failingSections, escapes)` manager-vantage on the 3rd org (Northwind Aviation), frontier-exhausted, on a fresh demo-up |
| **Target** | `(0, 0)` + dashboard ENABLED + ~80% all-3-complete + 1 hero STARTED + 1 COMPLETED |
| **Achieved** | `(0, 0)` — reachable 70, personaFailures 0, escapes 0, frontier EXHAUSTED; dashboard ENABLED, 78.4% all-3-complete, Ben STARTED (stage 1), Aria COMPLETED (stage 3), cycle `closed`, 199 frozen snapshots |
| **Distance** | 0 (gate MET, iter-09) |
| **Status** | ✅ `closed-on-gate` |

### Iter ledger summary
9 iters closed — 1 bootstrap tok (iter-01) + 8 tiks (iter-02…iter-09). Every iter has an `iter-NN/progress.md` with a close section; every iter maps to a commit (one-per-iter). Baseline manager sweep (iter-02) = `(failing=6, escapes=0)`; the frontier moved 6→5 (iter-05, harness/manifest) then held 5 across iter-06/07/08 (the AI-readiness read-path perf wall — three falsified strategies: active-signals → closed-cycle-snapshot → deep-link) then 5→0 at iter-09 (the app read-path demo-patch bounding `loadMembers`). TOK-01 (bootstrap, active-cycle signals-true) authored iter-01; TOK-02 (triggered by the iter-06/07/08 no-prog streak, user-authored out-of-band) pivoted to the app read-path demo-patch for the run-6 iter-09 tik that met the gate.

### Routes carried forward (Fate-2/3 — from the deferral re-audit + close)
| Item | Fate | Target | Note |
|---|---|---|---|
| Academy F6 (course content + hero menu-link + non-anon session) | **Fate-3 (LAND-NEXT)** | **M53** | User-decided at close (D-CLOSE-1); M53 cold-rebuilds the demo → the natural place to seed + verify academy content. AI chat documented-as-absent (AI-keys policy). |
| COLD reset-to-seed acceptance | Fate-2 | M53 | M53 `In:` owns the from-cold rebuild acceptance (user-decided at M50). |
| Re-pin consumption clone to the release rext tag + `.agentspace/rext.tag` bump | Fate-2 / push-gated KEEP | M53 | Box-level authoritative bump at M53 (v1.10.1). |

### Dropped
None.

### Protocol evolution
- **"Frozen SCORES ≠ frozen RESPONSE" (iter-08).** A closed-cycle frozen snapshot freezes the per-member scores, but `buildResponseFromSnapshots` still re-joins CURRENT members via an unbounded `loadMembers(orgID, "")` whole-org hydration → the fast branch is itself org-scale-slow. Lesson: measure the fast branch END-TO-END with a cheap probe before committing to a strategy. Blended into `coverage-protocol.md` + `services/ai-readiness.md` (#M51 iter-08/09).
- **Cheap-probe-first for a perf strategy (iter-08).** A dual-endpoint authed direct probe (`probe-aireadiness-deeplink.spec.ts`) falsified the deep-link premise BEFORE any cockpit/manifest edit or a full sweep — avoided shipping an inert `?cycle=`.
- **App read-path demo-patch as a pure, data-identical perf optimization (iter-09, TOK-02).** `app-aireadiness-snapshot-loadmembers` bounds the unbounded hydration to the ~199 snapshot users via the existing bounded sibling `loadMembersByUserIDs` (the members map is keyed+looked-up by snapshot UserID → byte-identical response). The `app-targetrole-authz-skip` precedent extended to a second app read-path swap.
