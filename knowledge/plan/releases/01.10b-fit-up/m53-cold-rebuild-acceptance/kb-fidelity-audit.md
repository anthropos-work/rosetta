---
title: "KB Fidelity Audit — M53 Cold-rebuild acceptance"
date: 2026-07-01
scope: milestone:M53
invoked-by: build-milestone
---

## Verdict
**GREEN** (proceed). Every load-bearing KB contract doc the M53 acceptance bar reads is ALIGNED with the
rext code at authoring HEAD `36d7430`. No blind areas, no stale load-bearing claims. Two low-severity
operational caveats are recorded below and in `decisions.md` (KB-1, KB-2) — they refine how M53 asserts,
they do not block.

## Topic Inventory

| Topic (acceptance criterion) | Knowledge doc | Code path | Status |
|---|---|---|---|
| Demo lifecycle / bring-up sequence / #7 abort | `corpus/ops/rosetta_demo.md` (+ demo-up SKILL/GUIDE) | `demo-stack/up-injected.sh`, `rosetta-demo` | PAIRED — ALIGNED |
| `/demo-down` image cleanup (M49 #6) | `corpus/ops/rosetta_demo.md` | `demo-stack/rosetta-demo:162-171` (`--purge`) | PAIRED — ALIGNED |
| Cold-start MCP-DSN capture (AB2) | `corpus/ops/snapshot-cold-start.md` | `stack-snapshot/`, `pg/pg.go:114-134` | PAIRED — ALIGNED (see KB-1) |
| Re-run idempotency | `corpus/ops/idempotency.md` | replay/seed/casbin paths | PAIRED — ALIGNED |
| Auto-verify net (AB1/AB3) | `corpus/ops/verification.md` | `stack-verify/live/autoverify.sh`,`verify.sh` | PAIRED — ALIGNED |
| Both-vantage coverage (AB4) | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/run-coverage.sh` | PAIRED — ALIGNED |
| AI-readiness dashboard (AB5) | `corpus/services/ai-readiness.md`, `seeding-spec.md` | `internal/workforce/ai_readiness.go`, patch + seeders | PAIRED — ALIGNED (see KB-2) |
| Seed+gen manifest download (AB6) | `corpus/ops/demo/cockpit-spec.md`, `seed-manifest-spec.md` | `demo-stack/cockpit.py`, `stack-seeding/manifest/` | PAIRED — ALIGNED |
| Academy F6 (course content) | `corpus/services/ant-academy.md` | `stack-demo/ant-academy/code/public/content` (static JSON, 3250 files) | PAIRED — content ships with clone; F6(i) verify-only |
| Academy F6 (menu-link) | `corpus/ops/demo/cockpit-spec.md`, `ant-academy.md` | cockpit `DeepLinkCatalog` is next-web-only | BLIND — net-new in §1 |
| Academy F6 (authenticated session) | `frontend-tier.md § ant-academy` | `ant-academy.sh` runs anonymous `BENCHMARK_VISUAL_BYPASS` | BLIND (contradicts) — net-new in §1 |
| Academy AI chat absent-in-demo | `ant-academy.md` (Cosmo default-OFF) | `ant-academy.sh` sets no key/flag | DOC-thin — add one demo-contract line in §5 |

## Fidelity Findings (all ALIGNED; caveats are refinements, not divergences)

1. **AB2 — "cold-start auto-capture with NO prompt" is a REPLAY-FROM-FILLED-CACHE, not a capture-during-bring-up.**
   The docs+code are explicit: `/demo-up` on a cold box **replays only, never captures** (`up-injected.sh:665`,
   `snapshot-cold-start.md:110,198-211`). M47 makes the *operator's* one-time capture turnkey (MCP DSN, no
   `~/.pgpass`) — it does NOT auto-capture inside bring-up. The `.agentspace/snapshots/` cache is present and
   fully populated (1.4 GB; taxonomy + directus + sim-embeddings each with COPY files + manifest.json), so the
   cold rebuild set-dresses from cache prompt-free. **M53 asserts: cache present → cold replay is prompt-free +
   set-dress succeeds.** If M53 expected a truly-empty-cache auto-fill during `/demo-up`, that would contradict
   the (correct) doc — it does not; the cache is already filled by M47's turnkey capture. → KB-1.

2. **AB5 — AI-readiness shipped numbers are 78.4% over 199 frozen snapshots, not 80%/≈160.** `ai-readiness.md:106`
   uses the round "~80% / ≈160 of 200" contract-phase carryover; the shipped funnel + `seeding-spec.md:369-375`
   are **78.4% / 199 snapshots**. M53 asserts against the shipped 78.4%/199. Also load-bearing: the fast frozen
   read only fires on a `?cycle=<closed>` deep-link; the cockpit AI-readiness deep-link (`cockpit.go:74`) is the
   bare `/enterprise/workforce/ai-readiness` (no `?cycle=`). The M51 `loadMembers`-bound patch bounds hydration
   so even the live-path GET renders acceptably; M53 validates the dashboard renders (fast, not a 180s timeout)
   on whatever link the cockpit/coverage harness actually uses. → KB-2.

3. **Coverage runner invocations (AB4) — confirmed exact:** from `stack-verify/e2e/`:
   `./run-coverage.sh <N> employee` (defaults identity `maya-thriving`, org "Cervato Systems") and
   `./run-coverage.sh <N> manager` (defaults `dan-manager`). Both assert `gateMet` on a frontier-exhausted
   crawl (`cappedAtFrontier===false`); on a CAP-HIT warning, re-run prefixed with a higher `COVERAGE_MAX_PAGES`
   (default 150). The manager run exercises the M51 AI-readiness dashboard page.

## Completeness Gaps

- **Academy F6 is the milestone's one genuine build surface** (correctly scoped as new-code in overview.md):
  (i) content ships with the clone (verify-only), (ii) the menu-link is net-new (cockpit catalog is
  next-web-only — no academy entry today), (iii) the authenticated-session path is net-new and **inverts** the
  current anonymous `BENCHMARK_VISUAL_BYPASS` model. (iv) "AI chat absent in demo" is *implied* (Cosmo
  default-OFF + no key provisioned) but not stated as a demo contract — add one line in `frontend-tier.md`
  §ant-academy during §5. These are the milestone's planned deliverables, not blind areas that block — F6 is
  an explicit Fate-3-landed overview.md `In:` item.

## Applied Fixes
None inline. The two number/operational caveats are recorded as KB-1/KB-2 in `decisions.md` and honored when
M53 asserts; the academy F6 doc line is authored in §5 (Phase 5 of the academy section) as planned work.

## Open Items (require user decision)
None.

## Gate Result
**GREEN — proceed to Phase 1 (§1 academy F6 build).** All acceptance-bar contract docs align with code; the
two caveats refine assertion targets (78.4%/199, replay-from-cache) rather than blocking; academy F6 is
planned new-code, not a blind area.
