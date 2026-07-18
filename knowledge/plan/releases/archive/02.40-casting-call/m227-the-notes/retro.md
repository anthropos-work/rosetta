# M227 "the notes" — Retro

_Section milestone, closed 2026-07-17. v2.4 "casting call" (RE-OPENED for believability)._

## Summary
A believability-hardening backfill triggered by **live feedback after M226 shipped the working hiring demo**: the
demo functioned on `billion` but didn't fully *read* as real. Four seed/content realism fixes landed, all in the
`rosetta-extensions` seed tooling, **0 platform-repo edits**: **#1** hiring-only content (generic activity seeders
skip a hiring org), **#2** external candidate emails (role-keyed), **#3** 1 sim/candidate (~8/position) with the
compare gate retuned `≥40 → ≥6` everywhere, **#4** gender-consistent avatars across all orgs. Every fix is proven
**deterministically** by the unit+regression suite and write-path-fenced by the harden pass. The LOCAL live-render
re-prove (section 5) was environment-blocked and routed **Fate-2 → M228 "second night"** (already planned).

## Incidents This Cycle
- **P2 (environmental, not code) — the local Docker box wedged during section 5.** A `/demo-down --purge` removed the
  working demo's images; the cold rebuild hit **ENOSPC**; `docker builder prune` (to recover disk) evicted the
  `go-build` cache mount → ~35-min cold recompiles; `buildx` wedged under host CPU contention. Root-caused a real
  bug along the way: the demopatch **G6 refusal** (the bring-up had run the *authoring* clone's `up-injected` rather
  than the consumption clone + a registry `type:demo` row — both fixed). **Not a code defect**; the data correctness
  was proven deterministically instead, and the live render was routed to M228 (the billion venue). Box is recovering.
- **No test flakes, no regressions.** The 4 fixes are correct; harden surfaced 0 production bugs.

## What Went Well
- **Deterministic proof over live-render dependence.** Each fix has an exact-invariant unit/regression test
  (`TestGenericActivitySeeders_SkipHiringOrg`, `TestCandidateEmails_RosterMatchesSeed`, `TestHiringFunnelSeeder_Funnel`,
  `TestPhotoAvatarForName_GenderMatched`) — stronger than ad-hoc SQL, and unaffected by the wedged Docker box. The
  D17 keeper ("only an executable probe binds") held: the milestone shipped its correctness signal even when the
  environment couldn't render.
- **Harden found the load-bearing gap.** The build-phase fix#2/#4 tests re-derived the helpers and compared to the
  Clerkenstein roster — **neither ran the real `UsersSeeder`**, so a revert of the `users.go` call sites would have
  shipped SILENTLY (candidates → org domain, a "Sara" → a man's photo, all tests green). Harden added
  `users_m227_test.go` driving the actual seeder over the whole population — the one fence that matters, RED-proven.
- **The gate retune was threaded consistently** (`≥40 → ≥6`) across all 8 surfaces, and a Go↔render-probe drift fence
  now *enforces* the Go-const/TS-spec agreement the docs describe (`TestHiringComparableFloor_MatchesRenderProbeDefault`).
- **Round-robin 1-position split** gave the tightest, most predictable distribution (43 assessed → min 8 / max 9).

## What Didn't
- **The build-phase tests documented the mechanism but not the corpus.** Three of the four fix mechanisms (hiring-only,
  external emails, gender avatars) were undocumented in `knowledge/corpus` until the close Phase-5 blend — the
  `Delivers →` promise was only half-met during build. Caught and fixed at close (#M227-D1/#M227-D2/#M227-D4), but the
  lesson: land the corpus doc *with* the fix, not at close.
- **The local environment is a single point of failure for the live-render loop.** A `--purge` on a disk-tight box
  cost the whole section. M228's clean VM (warm cache) is the right venue; local live-render should stay a
  nice-to-have, never a gate.

## Carried Forward
- **Section 5 — the LOCAL live full-stack render/coverage/playthrough re-prove on the corrected data → M228
  "second night"** (Fate-2, already planned; M228's exit gate explicitly covers the corrected-data render). Data
  correctness proven here; only the live full-stack render is routed.
- **Inherited (→ v2.4 RELEASE close):** 8 pre-existing demo-stack test failures + the M204 assign-WRITE declared TODO.
- **DEF-M226-01** (pre-bind serve reap, Fate 3) → M228 / next prove-on-VM; self-resolves in the default flow.

## Metrics Delta
- Go test funcs **1888 → 1902** (+14). Flake **0** (5/5 gate). `go vet` clean; tsc `stack-verify/e2e` exit 0.
- **0 platform-repo edits.** Supply chain GREEN (0 net-new direct deps). Alignment 100%/100% (untouched).
- Full artifact: [`metrics.json`](metrics.json). Deferral audit: [`audit-deferrals/deferral-audit-2026-07-17-m227-close.md`](audit-deferrals/deferral-audit-2026-07-17-m227-close.md).
