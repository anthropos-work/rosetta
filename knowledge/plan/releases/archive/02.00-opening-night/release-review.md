# Release Review — v2.0 "opening night" (the Playthroughs pillar)

**Reviewed:** 2026-07-02 · **Branch:** `release/02.00-opening-night` @ `88c2be8` · **Diff vs `main`:** 66 files, +2799 / −141
**Milestones:** M201 (manifest corpus) ∥ M202 (foundation) → { M203 (employee) ∥ M204 (manager) } — all closed
**Prior release compared:** v1.10b "fit-up" (`releases/archive/01.10b-fit-up/metrics.json`)

**Consolidated verdict: GREEN — clean for merge.** No blockers. One required reconciliation (a stale M201
status label) + a small should-fix cluster (all Fate-1, land-now). Tooling + docs only; **zero platform-repo
edits; zero net-new third-party deps.**

---

## Phase 0 — Supply-chain · GREEN

- **Net-new deps: 0.** The new `playthroughs/` Go module needs only `gopkg.in/yaml.v3 v3.0.1` (already rext-wide)
  + stdlib; the TS e2e layer needs only `@playwright/test` (shared with `stack-verify/e2e`) + `@types/node` +
  `typescript`. `ai v1.40.1` stays confined to `stack-seeding` (unchanged).
- **Go CVEs (govulncheck, go1.25.11): 0 high / 0 critical.** `playthroughs/...` → "No vulnerabilities found";
  all 5 other rext modules re-scanned clean. **Carry-forward closed:** the v1.10b-flagged inherited HIGH
  GO-2026-5026 / CVE-2026-39821 (`x/net@v0.53.0` path in stack-seeding) is now **RESOLVED** — `x/net v0.55.0`.
- **Node CVEs (npm audit): 0 high / 0 critical** (`--omit=dev` + full) for `playthroughs/e2e` and `stack-verify/e2e`.
- **License:** all Apache-2.0 / MIT / BSD; **0 GPL/AGPL**. Internal tooling, not distributed.
- **Lockfile:** `dependencies.lock` written (uncommitted).

## Phase 1 — Release scope · clean (1 reconciliation)

- **DELIVERED:** all 4 milestones closed. **10 live `@pt:` Playthroughs + 1 TODO** — matches the roadmap claim exactly.
  M201 corpus (9 products · 26 stories · 28 use-cases, user-signed-off) → M202 validator/foundation →
  M203 6/6 employee GREEN cold reset-to-seed → M204 4/4 manager GREEN. The corpus→validator→coverage contract held.
- **Fate-2 CONFIRMED:** `assignment-monitoring.assign-and-track.UC1` (assign-WRITE, two-backend org-admin write) is a
  **declared in-manifest TODO** (`assignment-monitoring.yaml:50`), surfaced by `ptreport` as `unimplemented`,
  presence-pinned by `TestRealCorpus_ManagerCoverageIsPresent` (fail-red if removed). Recorded M204 D-CLOSE-1.
- **Fate-3 VERIFIED-LANDED (M203 → M206):** all 4 non-gate employee UCs carried in `roadmap-vision.md` M206
  (code-sim/Judge0, text-interview, verify-terminal, self-eval) with Fate-3 provenance. Recorded M203 D-CLOSE-1.
- **Fate-3 UNDELIVERED: NONE. Escape-hatch: NONE. Dropped: NONE. Unaccounted: NONE.**
- **⚠️ RECONCILIATION (Phase 7):** `m201-manifest-corpus/overview.md` frontmatter still reads **`status: planned`** —
  a stale label. M201 was closed-on-gate + merged at `1ccde8f` before the v1.10b pause; the close never flipped it
  (its own Gate Outcome Ledger + roadmap + state.md all record it done). M202/M203/M204 all read `archived`.
  **Flip `planned` → `archived`.** A genuinely-done milestone, not a scope gap.

## Phase 1b — Deferral re-audit (`--scope=release`) · GREEN

- Total in scope: **5** (4 M203 non-gate employee UCs → Fate-3 M206; 1 M204 assign-WRITE → Fate-2 in-manifest).
  All single. **Repeat: 0 · Chronic: 0 · Aged-out: 0 · Escape-hatch: 0.** SEVERITY=clear.
- All 3 prior per-milestone close audits (M202/M203/M204) were GREEN; the release re-audit confirms nothing slipped.
- Source grep: **0** stray FIXME/HACK/XXX + 0 non-sentinel TODO in milestone-authored rext source.
- Report: `audit-deferrals/deferral-audit-2026-07-02-release-close.md` (created).

## Phase 2 — Code quality · strong (no must-fix)

- **Build green across all 4 gates:** `go vet ./...` · `gofmt -l` · `go build ./...` · `npx tsc --noEmit` (e2e) — all clean.
- **Cross-milestone consistency verified:** the M203/M204 "additive not conflicting" claim **holds** — employee
  (`/skill-path`, `/sim`, `/profile`) and manager (`/enterprise/*`) route namespaces are disjoint; the two intentional
  parent/leaf overlaps (`WORKFORCE_URL`⊃`SUCCESSION_URL`; `ACTIVITY_DASHBOARD_URL`⊃`ACTIVITY_DRILLDOWN_URL`) are
  documented + unit-pinned. No O(tests) re-pin creep (1 pattern + 1 predicate + 1 agreement test per surface).
  The M204 reporter-override fix is a clean root-cause fix, not iteration residue.
- **Should-fix (1):** dead base method `region()` — `e2e/lib/page-object.ts:59-65`. **Zero call sites** (every surface
  uses `this.main()`); holds the layer's only xpath step; docstring claims a `<main>` fallback the body doesn't
  implement. Exactly the posture the project's own M203-close F3 rule says to prune. **Couples to a Phase-3 docs finding.**
- **Nice-to-have:** `stepHeading()` bypasses the `byRole` helper; Sentinel port `8087+OFFSET` hardcoded inline vs the
  shared offset helper; `.gitignore` lists a never-produced `report/last-report.html`; `playwright install … || true`
  swallows failure; 3 forward-declared seed axes (`pt-free`, `entitlement-gated`, `multi-org-private`) with no
  consumer yet (intentional forward-provisioning — worth tracking so they don't become permanent orphans).

## Phase 3 — Docs · accurate (2 minor inaccuracies)

- `corpus/ops/demo/playthroughs.md` (389 lines) reads coherently, accurately reflects what shipped (10 live + 1 TODO
  verified 1:1), Sentinel-reload + reporter-override lessons captured, 4-state names/gates match `report.go` exactly,
  all cross-refs resolve (25/25 rext + 8/8 corpus), no stale in-progress/branch refs. state.md + roadmap current.
- **Inaccuracy 1:** `region(headingText)` listed as a live base primitive at **`playthroughs.md:183`** — it's dead code
  (Phase-2 should-fix). If `region()` is pruned, this line drops with it.
- **Inaccuracy 2:** `fixtures/` "populated by M203/M204's upload flows" at **`playthroughs.md:373-376`** — the dir stayed
  **empty** through M204 (no shipped spec uses `setInputFiles`). Reword to reserved / still-empty.

## Phase 3b — KB consolidation · clean (1 index nit)

- `playthroughs.md` indexed in CLAUDE.md (line 296) + demo/README (line 146); reciprocally backlinked from
  stories-spec.md + coverage-protocol.md ("function sibling"). No orphan. 389 lines — no split needed. No blind area.
- **Should-fix (minor):** the **CLAUDE.md** index blurb (line 296) tags the pillar "M202" and describes the iteration
  protocol but — unlike demo/README — omits the shipped outcome (10 live / 1 TODO, M202–M204). Sync the two blurbs.

## Phase 4 — Tests · GREEN

- **Playthroughs Go:** `go test ./...` EXIT 0 (manifest/report/ptvalidate/ptreport). **105 test/fuzz funcs** (101 + 4
  fuzz) — matches the M204 recorded total. Flake gate: **5/5 shuffled `-race` clean.**
- **Playthroughs TS unit:** **58 passed** (url-shapes 46 + stack-env 12); `tsc --noEmit` clean; flake **5/5 clean**.
- Live Playthroughs gate-verified per milestone (M203 6/6 5/5, M204 4/4 5/5); the ~13-min sweep not re-run at release
  close by design. No untracked coverage gap (the 1 Fate-2 TODO + 4 Fate-3 UCs all tracked).

## Phase 4b — Metrics regression · GREEN

- Aggregated `m201..m204/metrics.json` → `metrics.json` (uncommitted). playthroughs Go **105** · TS unit **58** ·
  live Playthroughs **10** + 1 TODO · coverage 94.8–100% new module · flake **0** · deps **0 net-new**.
- vs v1.10b: new pillar (new-vs-old). All blocking rules pass — no test-count decrease, no coverage drop >2pp, no
  flake, no unsanctioned dep (x/net advisory resolved = a supply-chain improvement).

## Phase 5 — Decision consolidation · clean

- No unblended material decisions (all M201–M204 decisions verified landed in `playthroughs.md`), no cross-milestone
  conflicts (pt-world seed, §5.8 assertion boundary, serial runner, page-object discipline all consistent), both
  deferrals correctly routed + tracked (M203→M206 vision; M204 assign-WRITE Fate-2 in-manifest). Three release-level
  patterns (Sentinel-reload, reporter-override, segment-anchored URL discipline) already in the runbook.
- **Nice-to-have (2 optional runbook additions):** M204 iter-05 D1 (consumption-clone PATH gate-run prereq);
  M204 iter-03 D2 (in-main vs out-of-main table-scoping locator example).

---

## Phase 7 fix queue (all Fate-1 — land-now before tag)

| # | Fate | Item | Surface |
|---|------|------|---------|
| 1 | Required (reconcile) | Flip M201 `overview.md` `status: planned` → `archived` | `m201-manifest-corpus/overview.md:8` |
| 2 | Should-fix (code+docs) | Prune dead `region()` method + its docs twin | `e2e/lib/page-object.ts:59-65` + `playthroughs.md:183` |
| 3 | Should-fix (docs) | Reword `fixtures/` claim → reserved/still-empty | `playthroughs.md:373-376` |
| 4 | Should-fix (docs, minor) | Sync CLAUDE.md index blurb → shipped outcome (10 live/1 TODO, M202–M204) | `CLAUDE.md:296` |
| — | Nice-to-have (optional) | 2 runbook additions (PATH gate-run prereq; in/out-of-main scoping example) | `playthroughs.md` |

Fix #2 touches the rext authoring copy (`.agentspace/rosetta-extensions/`) → will be re-tagged as part of the v2.0
rext roll (Phase 10). Fixes #1/#3/#4 are rosetta-corpus only.

**Push-gated KEEP (unchanged):** origin still owes `main` + all tags (`v1.10.1`, `v2.0`, rext `fit-up-*` +
`opening-night-m201..m204` + the v2.0 rext roll). The user's manual step — not a defer, not a blocker.
