# Release Review: v1.9 "storytelling"

**Date:** 2026-06-23
**Milestones:** M34 (verified-skill-chain) · M35 (stories-multi-org) · M36 (dashboard-surfaces) · M37 (clerkenstein-multi-identity) · M38 (presenter-cockpit)
**Branch:** `release/01.90-storytelling` → `main`, tag `v1.9`
**Code-of-record (separate repo):** `rosetta-extensions` @ tags `storytelling-m34` (`8eb603b`) · `m35` (`06d872c`) · `m36` (`11e15e3`) · `m37` (`52c1be0`) · `m38` (`237bede`). This close merges ONLY the rosetta doc-half branch; the ext tags ARE the code-of-record.

## Verdict: GREEN — 0 blocking findings

All 5 milestones reviewed as one PR. 9 review sweeps (supply-chain / scope / deferral re-audit / code-quality / docs / KB-consolidation / tests / metrics-regression / decisions) all clear. The release turns the placeholder seeder into a declarative Stories & Heroes demo engine + a presenter cockpit, proven live on demo-3. **Tooling + docs only — zero platform-repo edits** (verified: the rosetta release diff is confined to `corpus/` + `knowledge/plan/` + `CLAUDE.md` + `.claude/skills/`; no `.go`/`.py`/`.sh`/`stack-*` files).

## Scope (Phase 1)
- [x] All 5 roadmap milestones (M34→M38) delivered Fate-1; Version Plan table + roadmap detail confirm each `done`. No unaccounted items.
- [x] Both product Musts delivered: the individual **skill profile** (M34 verified-skill chain) + the org **Workforce dashboard** (M36 dashboard surfaces).
- [x] Fate-3 routed-AND-delivered within release: #M34-D7 → landed in full as D-M35-4 at M35 close (verified — the target milestone's ledger delivered it).
- [x] No Fate-3-undelivered. No escape-hatch. No drops beyond declared scope cuts.

## Supply Chain (Phase 0)
- [x] **GREEN** — `go.mod`/`go.sum` diff across all of v1.9 (`storytelling-m34~1..m38`) = EMPTY. Zero new third-party deps. Dep surface byte-identical to v1.8.
- [x] Go third-party deps unchanged vetted set: pgx/v5, yaml.v3, colony, clerk-sdk-go/v2, uuid, svix-webhooks (all MIT/BSD/Apache — no GPL/AGPL). Python stdlib-only + the pre-existing optional PyYAML test dep (M5).
- [x] Lockfile written: `releases/01.90-storytelling/dependencies.lock`.

## Deferral Re-Audit (Phase 1b)
- [x] **GREEN** — zero open deferrals, zero repeat-patterns, zero aged-out, zero escape-hatch. Both v1.9 items terminated Fate-1 in-release (#M34-D7→D-M35-4; M38-D7→M38-D8). Inherited backlog (M33/DEF-M10-01/DEF-M21-01../M25-D9) orthogonal, KEEP. Report: `audit-deferrals/deferral-audit-2026-06-23-release-close.md`.

## Code Quality (Phase 2)
- [x] stack-seeding `go test ./...` → 8 packages `ok`; `go vet` clean; `gofmt -l` clean.
- [x] clerkenstein `go test ./...` → 14 packages `ok` (jwtkey no tests); `go vet` clean.
- [x] Cross-milestone consistency: the EffectiveStories() normalization seam (D-M35-1) is the single code path all 8 seeders iterate; the roleForHero single-source (M38-D8) extends the same discipline to the role triple. No conflicting patterns, no dead code from iteration, no leaked seams across the M34→M38 chain.

## Documentation (Phase 3 + 3b)
- [x] **Clean** (Explore release-level coherence review): stories-spec.md (409L NEW) reads coherently end-to-end; all cross-refs resolve; "four → five surfaces" consistent across clerkenstein.md + alignment_testing.md; all corpus docs reflect shipped/past-tense state. No must/should/nice findings.
- [x] **KB consolidation:** stories-spec.md 409L (under split threshold); properly indexed from README.md/CLAUDE.md/seeding-spec.md/safety.md; `corpus/ops/demo/` cluster coherent. No structural debt.

## Tests & Benchmarks (Phase 4)
- [x] Go (rext, run at HEAD = `storytelling-m38`): alignment 52 · clerkenstein 259 · stack-seeding 444 · stack-snapshot 333 · stack-secrets 160 = **1248** (`Test`+`Fuzz`). All packages `ok` under `go test`.
- [x] Python: demo-stack 166/166 OK · stack-injection 117 OK (8 opt-in skipped) = **283** across M38-touched surfaces.
- [x] Flake 0 across all milestone gates (5/5 shuffled `-race` Go + Python). Integration tests opt-in behind `//go:build integration` + `STACKSEED_IT_DSN`.
- [x] Benchmarks: n/a (DB-IO-bound bulk COPY seeders + stdlib HTTP cockpit — no perf hot path).

## Metrics Regression (Phase 4b)
- [x] **GREEN** — vs v1.8 baseline: Go stack-seeding 259→444 (+185); rext total 1027→1248 (+221, no decrease). Python: demo-stack 110→166, stack-injection 113→117 (the two M38-touched suites grew; no decrease). Coverage no >2pp drop on any measured surface. Flake 0. Supply-chain unchanged. Aggregated metrics: `releases/01.90-storytelling/metrics.json`.

## Alignment Gates
- [x] **100%/100% on all 5 Clerkenstein surfaces** — clerk-multi-1 (NEW M37, 9 genes) 9/9 + Go clerk-2.6.0 22/22 + JS clerk-js-5 9/9 + deploy clerk-deploy-1 7/7. The `clerk-express-1` node-CI gate drives the genuine `@clerk/express` SDK (env prereq — needs installed npm modules — not a regression; M37/M38 never touched it). Re-verified at M37 + M38 closes.

## Decision Consolidation (Phase 5)
- [x] Load-bearing decisions blended into corpus with traceable tags: stories-spec.md carries `#M34-D2/D3`, `#M36-D1/D2/D3`, `#M38-D8`. Per-milestone metrics confirm the remainder blended at build or archived maintainer-only (D-M34-1..6, D-M35-*, D-M37 O11/ARCH/KB-1). No cross-milestone decision conflicts. The `wip/clerkenstein-browser-login` branch was reconciled (note folded into architecture.md) + retired at M37 close.

## Headline
**GREEN release, 0 blocking findings.** Tooling + docs only · zero platform-repo edits throughout · 5 alignment gates 100%/100% · supply-chain GREEN (stdlib-only, 0 new deps) · significant test growth (stack-seeding 259→444; rext total 1027→1248; +clerk-multi-1 5th alignment surface) · proven live on demo-3.

---

## Release Completeness Ledger: v1.9 (Phase 9)

All 5 milestones are `section` shape, terminally `archived`. **Clean release.**

### Delivered (Fate 1 across the release)
- **Verified-skill chain (the spine)** — delivered in M34; the 7-table fan-out per (hero × skill), real replayed taxonomy node-ids (never fabricated), G14 session fix, the seed-side closure gene (0 dangling). Verified by the integration test + live on demo-3 (Maya: 8 verified skills, 24 Spotlight datapoints).
- **Stories & Heroes multi-org engine** — delivered in M35; one `stack.stories.yaml` seeds multiple orgs each with a thriving/struggling/manager hero trio at vantage-appropriate fidelity; `EffectiveStories()` normalization keeps the legacy single-org path byte-identical. Verified live (2 orgs, 6 heroes, closure green across both).
- **Workforce dashboard surfaces (Must #2)** — delivered in M36; 6 dashboard seeders (membership_skills funnel, tags/teams, target-roles, succession, feedback, population_evidence) + 2 fixes. Verified live (Cervato org: every aggregate populated + distributed).
- **Clerkenstein multi-identity (seat-switch)** — delivered in M37; the users/orgs Registry + server-authoritative FAPI handshake selection + the 5th Alignment DNA `clerk-multi-1` (9 genes, 100%/100%). Single-identity path byte-identical.
- **Presenter cockpit** — delivered in M38; the standalone served panel reading the same `stack.stories.yaml`, [Login as]+[Jump] = one M37 handshake redirect, the roster-export + cockpit-manifest producers + the O9 deep-link catalog. Verified (cockpit keys == roster keys == seeded users).
- **Vantage-faithful `org_role`** — delivered as the M38 close fix (M38-D8); `roleForHero` single-source so the membership row + casbin grant + roster/JWT claim agree per hero (manager→admin, end-user→member).

### Iterative milestone gates
None — all 5 milestones are `section` shape.

### Fate 2/3 — Routed within this release (verified landed)
- **#M34-D7** (multi-hero index-collision guard + short-role-pool flat top-up) — Fate-3 from M34 → **delivered in full in M35 as D-M35-4** (both parts; the chosen fix was strictly better than the routed idea). Verified by the M35 multiorg + trajectory regression tests.

### Fate 3 — Routed within this release (UNDELIVERED)
None.

### Escape-hatch — Cross-release deferral (requires user sign-off)
None. Zero `RELEASE-SCOPE-DEFER:` decisions across v1.9.

### Dropped
None beyond the long-standing pre-v1.9 scope cuts (the former v1.4 seeds — AI content, shareability, more mirrors — dropped 2026-06-11, untouched by v1.9).

### Unaccounted
None. Every roadmap deliverable maps to a milestone ledger; every code item accounted for (no TODO/FIXME/HACK in product files — the two "XXX" hits are literal node-id format examples in comments).

### Non-Fate-1 highlight
```
✓ ALL RELEASE SCOPE ITEMS DELIVERED.
  Iterative milestones:        0 (all 5 are section shape)
  Fate-3 undelivered:          0
  Escape-hatch deferrals:      0
  Dropped (beyond declared):   0
  Unaccounted:                 0  = clean release
```
**All release scope items delivered. No iterative gate misses, no Fate-3-undelivered, no escape-hatch deferrals, no drops beyond declared scope cuts. Clean release.** Tags as a clean `v1.9` (not `v1.9-incomplete`).
