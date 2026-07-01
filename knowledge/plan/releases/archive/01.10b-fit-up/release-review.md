# Release Review: v1.10b "fit-up"

**Date:** 2026-07-01
**Milestones:** M47, M48, M49, M50, M51, M52, M53 (7)
**Tag target:** `v1.10.1` (rext code-of-record @ `v1.10.1`)
**Verdict:** GREEN — zero blockers, zero must-fix. Cold-rebuild acceptance passed 6/6 + academy F6 (M53).

**Phase-7 (fix pass, 2026-07-01):** all 6 user-approved review fixes LANDED (2 should-fix + 1 security + 3 nice-to-have; every checkbox below ticked). The inherited HIGH CVE-2026-39821 is CLEARED (`govulncheck`: no vulnerabilities). rext suites re-verified green (stack-seeding Go + TS unit 42/42); other Go modules + Python untouched. `v1.10.1` re-rolled to the post-fix rext HEAD.

## Supply-chain (Phase 0)
- [x] [should-fix / security] **CVE-2026-39821** — `golang.org/x/net@v0.53.0` (idna hostname-validation bypass), CALLED HIGH, path exists only in `stack-seeding` (`reloadStackSentinel → idna.ToASCII`). **Inherited** (in the tree since v1.10; CVE disclosed after the v1.10 close — NOT a v1.10b regression). Indirect via `ai v1.40.1`. Fix: `go get golang.org/x/net@v0.55.0 && go mod tidy` in stack-seeding. Internal-tooling / not-distributed context ⇒ non-blocking, but cheap + a called HIGH → recommend landing before the tag. **[Phase-7 LANDED]** bumped `golang.org/x/net v0.53.0 → v0.55.0` (+ `x/text v0.36.0 → v0.37.0`) in `stack-seeding`; `go build`/`go vet`/`go test ./...` all green; `govulncheck ./...` now reports **"No vulnerabilities found"** — CVE-2026-39821 no longer CALLED.
- Node (`npm audit`, stack-verify/e2e): **0 vulns**. `ai v1.40.1` present + verified. Licenses clean (Apache/MIT/BSD; 0 GPL/AGPL). `dependencies.lock` written.

## Scope (Phase 1 + 1b)
- All 7 milestones delivered Fate-1. Fate-3 chains all verified-landed (academy F6 M50→M51→M53 ✔; up-injected.sh M52→M53 ✔; AB4 fixed-at-gate ✔). Zero Fate-3-undelivered, zero escape-hatch, zero dropped, zero unaccounted. Deferral audit **GREEN** (one historical repeat — academy F6 — resolved by execution). Origin pushes = push-gated user step, not a blocker.

## Code Quality (Phase 2) — zero must-fix
- [x] [should-fix] `gofmt -w` two test files (`stack-seeding/cmd/gen-batch/main_test.go`, `stack-seeding/seeders/ai_readiness_harden_test.go`) — comment-column drift only. **[Phase-7 LANDED]** both gofmt'd; `gofmt -l stack-seeding/` clean.
- [x] [should-fix] Abandoned iter-08 diagnostic probe `stack-verify/e2e/tests/probe-aireadiness-deeplink.spec.ts` hard-asserts a FALSIFIED precondition (`frozenFast` — the deep-link-fast claim the app-patch disproved). Not gate-run, but collected by a bare `playwright test`. Delete, or relabel as a soft/forensic diagnostic. **[Phase-7 LANDED]** deleted via `git rm` (cleaner than relabel — it served its diagnostic purpose); TS unit suites still 42/42 (coverage-manifest 29 + section-assert 13).
- [x] [should-fix] `warmHeavyGrids` (`coverage.spec.ts`) unconditionally warms `/enterprise/workforce/ai-readiness` for every manager org, but M53 made the manifest org-conditional — burns the ~25s ceiling polling a legitimately-empty page on base orgs. Derive warm paths from the resolved `manifest.seedPaths`. **[Phase-7 LANDED]** warm paths now = the heavy-grid set ∩ `manifest.seedPaths`, so ai-readiness is warmed ONLY for the showcase org (where the manifest primes it); base orgs no longer poll the empty page. TS unit suite green.
- [ ] [nice-to-have] `member_languages` surface token uses underscore vs the fleet hyphen-case (cosmetic; not referenced in any DependsOn).
- Confirmed clean: the app `loadMembers` demo-patch (reversible, hash-pinned to app v1.315, data-identical); 5 new seeders isolation/scoping intact; the M53 org-conditional manifest cleanly supersedes M51's unconditional version.

## Documentation (Phase 3 + 3b) — coherent + discoverable
- [x] [nice-to-have] PostHog phrasing: `ai-readiness.md` says the UI needs org-setting AND `flag_ai_readiness`; `stories-spec.md` says "org setting, not a PostHog flag." Not contradictory (seeder-writes vs UI-gate), but a one-line clarify would help. **[Phase-7 LANDED]** added a "two gates are different layers — not a contradiction" note in `ai-readiness.md` (§Org enablement) + a matching gate-1/gate-2 clarification with cross-link in `stories-spec.md` (the `OrgSettingsSeeder` row).
- [x] [nice-to-have] Blind area: no demo doc explains how the demo next-web satisfies the FE `flag_ai_readiness` gate (the M42m gate provably renders the dashboard, but the mechanism is undocumented). One-line note closes the loop. **[Phase-7 LANDED]** documented in `ai-readiness.md`: the demo bakes **no** `NEXT_PUBLIC_POSTHOG_KEY`, so the client-side flag check has no PostHog backend to consult + doesn't block; empirically proven by M53 AB5 (dashboard renders from cold on Northwind). Honest caveat included (exact in-SDK default-through inferred from "no key + AB5 renders", not separately FE-traced).
- Both new docs (`ai-readiness.md`, `seed-manifest-spec.md`) indexed + referenced (8 / 6 inbound); 78.4%/199 reconciles across all docs; no stale branch/WIP refs; no split candidates (max 619 lines).

## Tests & Benchmarks (Phase 4 + 4b) — GREEN
- Full suites pass, flake 0: rext Go **1640** (5 modules), Python **326**, TS unit **42**. (Playwright coverage sweep out of scope — cold-accepted at M53.)
- Metrics regression **GREEN**: Go 1551→1640 (+89), TS 17→42, no coverage drop >2pp, flake 0, 0 new deps. `metrics.json` aggregated + written. (Apples-to-oranges caveat: v1.10b is a docs+tooling backfill; the load-bearing "no test regression + flake 0" holds.)

## Decision Consolidation (Phase 5) — clean
- Zero unblended decisions (all triaged-to-knowledge entries verified landed). Zero unresolved conflicts (the seed-strategy saga converged coherently — docs document the shipped closed-cycle + app-patch). Both release-level lessons captured (perf saga → `coverage-protocol.md`; sub-agent sweep-fragility → M51 retro as a process lesson).
