# M51 — Retro (AI-readiness showcase org)

## Summary
Iterative milestone, **closed-on-gate** 2026-07-01. 9 iters (1 bootstrap tok + 8 tiks) drove the M42 manager-vantage
semantic coverage gate to **MET** on the 3rd org (Northwind Aviation, 200 members): `(failingSections, escapes) =
(0, 0)` frontier-exhausted on a fresh demo-up (reachable 70, personaFailures 0), dashboard ENABLED, 78.4%
all-3-complete, Ben STARTED (stage 1) + Aria COMPLETED (stage 3), cycle `closed` + 199 frozen snapshots. The bulk of
the milestone (iters 06→08) was a **strategy saga** against a *platform-side* AI-readiness read-path perf wall: three
successive read-fast strategies (active-cycle signals-true → closed-cycle frozen snapshots → deep-link the frozen
branch) were each falsified by a cheap probe, culminating in the iter-08 root cause **"frozen SCORES ≠ frozen
RESPONSE"** — `buildResponseFromSnapshots` re-joins CURRENT members via an unbounded whole-org `loadMembers`, so even
the "fast" frozen branch is org-scale-slow. The user (out-of-band, TOK-02) chose a NEW app read-path demo-patch;
iter-09 landed `app-aireadiness-snapshot-loadmembers` (a PURE data-identical bound of that hydration to the ~199
snapshot users), turning the frozen `?cycle=` GET from a 180s timeout into 19ms, and the gate fell 5→0.

## Incidents This Cycle
- **P2 — iter-04 dirty consumption-clone user-blocker (resolved out-of-band).** demo-1's `stack-demo/rosetta-extensions`
  carried leftover hand-modifications (a partial M50 application from a prior concurrency incident) that blocked the
  `git checkout fit-up-m50` re-pin — unblockable only via forbidden ops (`git clean`/`git checkout --`). The
  orchestrator reset the clone to a clean `fit-up-m50`; iter-04 resumed. Cost: one iter paused; no bad state shipped.
- **P2 — three falsified read-fast strategies (iters 06/07/08).** Each iter implemented a plausible read-fast strategy
  and the cheap-probe-first discipline falsified it before a full sweep or a shipped edit — active-signals never
  completes (per-skill translation N+1); the closed-cycle DEFAULT FE GET omits `?cycle=` so the frozen branch is
  never selected; the deep-link's frozen branch is itself org-scale-slow. Net: the strategy space was correctly
  exhausted (not blind guesses), but it took 3 tiks — which correctly triggered TOK-02. No inert code shipped (no
  hanging `?cycle=` deep-link was committed to the cockpit/manifest).
- **P3 — the C1 `--reset` guard shipped with a build error (close-caught).** The close-phase `to_regclass` guard used
  a pgx-style `conn.QueryRow(...).Scan(...)` signature; the repo's `Conn.QueryRow(ctx, sql, []any, dest...)` differs.
  Caught by the Phase-7 `go build`; fixed to the correct signature before commit. A CI build gate on the WIP would
  have caught it earlier.
- **P3 — the fix-queue was pre-checked before the work landed (resume-caught).** The prior close agent (which died
  mid-Phase-7) had marked C4/T1/T2 `[x]` in `progress.md` before authoring them. The resume verified ground truth
  (the test files didn't exist), authored them for real, and corrected the queue. Lesson: check-off must follow the
  landed artifact, not the intent.

## What Went Well
- **Cheap-probe-first saved real iters.** The dual-endpoint authed direct probe (`probe-aireadiness-deeplink.spec.ts`)
  falsified the deep-link premise in 40ms+180s-timeout BEFORE any cockpit/manifest edit or a ~13min sweep — the single
  highest-leverage protocol move of the milestone.
- **The tik→tok discipline worked as designed.** Three no-progress tiks (06/07/08) correctly surfaced the
  user-blocker rather than thrashing; the user reviewed the three falsifications and authored the TOK-02 pivot; the
  gate fell in one tik once the right (bounded-query) lever was identified.
- **CODE-owns-structure held.** Every AI-readiness skill/sim ref routes through the real replayed taxonomy resolvers
  (never fabricated); the closure gate stayed green; the frozen snapshots derive from the same signals the seeders
  write. The app read-path demo-patch is a pure, data-identical perf optimization (the response is byte-identical).
- **Zero-platform-edit held.** The read-path speedup is a demo-only in-inject-loop source swap (the
  `app-targetrole-authz-skip` precedent extended), not a canonical-repo edit.

## What Didn't
- Most of the milestone's iters (06→08) went to a *platform* perf characteristic, not to the seeding deliverable —
  the seed itself was structurally done by iter-03. The exit gate (a UI coverage sweep) coupled the milestone's
  success to a platform read-path wall the tooling could only work around, not fix.
- The residual (M314b — the prod frozen-read still hydrates the whole org) is a real prod-finding that the demo
  patch works around but the platform still carries; documented, not fixed (correctly out of the tooling's scope).

## Carried Forward
- **Academy F6 (course content + hero academy menu-link + non-anonymous academy session) → M53** (Fate-3, LAND-NEXT,
  user-decided, D-CLOSE-1). The repeat-defer (M50→M51) is resolved: M53 cold-rebuilds the demo, the natural place to
  seed + verify academy content on a clean build. AI chat stays documented-as-absent (AI-keys policy).
- **COLD reset-to-seed acceptance → M53** (Fate-2, user-decided at M50).
- **Consumption-clone re-pin to the release rext tag + `.agentspace/rext.tag` bump → M53** (push-gated KEEP).
- **M314b (prod frozen-read whole-org hydration / a `frozen_tags` column)** = a documented prod-finding in
  `coverage-protocol.md` + `services/ai-readiness.md`; NOT a demo fix (out of the tooling's scope).

## Metrics Delta (from metrics.json)
- rext stack-seeding Go tests **719 → 749** (+30 across the 9 iters + 5-pass final harden + close's +2); seeders pkg
  coverage **97.4% → 97.6%**.
- rext e2e TS unit **20 → 33** (+13: the NEW `section-assert.ts` no-browser verdict spec).
- Flake **0** (5/5 Go [seeders + cmd/stackseed] + 5/5 TS [both unit specs]).
- Close: 16 findings, all Fate-1; deferral audit RED→CLEARED (academy F6 → M53, Fate-3).

## Lessons (pinned)
1. **Measure the FAST branch end-to-end before committing to it — "frozen SCORES ≠ frozen RESPONSE."** A cached/frozen
   score path can still re-incur an org-scale cost downstream (a live member re-join). A cheap direct probe on the
   exact fast endpoint is worth more than a full sweep. (Blended into `coverage-protocol.md` + `services/ai-readiness.md`,
   #M51 iter-08/09.)
2. **A UI coverage gate can couple a seeding milestone to a platform read-path characteristic.** When the deliverable
   (data) is done but the gate (a render sweep) still fails, root-cause WHERE the failure lives before iterating — the
   fix surface may be the demo-up path (a read-path patch), not the seeder.
3. **Check off a fix-queue item only when its artifact has landed + built + passed**, never on intent — the resume
   caught three pre-checked-but-absent tests.
