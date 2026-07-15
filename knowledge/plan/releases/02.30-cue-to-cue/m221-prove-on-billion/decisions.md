# M221 — Decisions

_Implementation + strategy choices with rationale (incl. tok strategy revisions)._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| — | _(none yet)_ | | |

## TOK-01: fix-then-prove, host-isolation first — 2026-07-15

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Fix-then-prove in a strict dependency order. Phase A: land `GUARD-M221-host-isolation`
(a stale-tolerant, identity-named lockfile on the demo host; a second concurrent cycle fails loud with the
holder's identity — never queues, never proceeds). Phase B: the off-box code fixes (academy loopback-bind +
extend `exposure_claim_guard` to host-native listeners; reap the native cockpit+academy; the academy
empty-catalog patch; a fence for the backend-api-url twin). Phase C: the live cold-reset-to-seed battery on
`billion`, driven from a tailnet peer, a DEFAULT `/demo-up N` with NO FLAGS, proving all 8 gate conditions +
folding in M219's five readiness gates re-proven at final code.
**Rationale:** The overview mandates the ordering twice — M221's own gate is a multi-cycle battery on the single
`billion` host, and M219 proved two agents on one host corrupt the evidence (a cycle purged the stack
mid-measurement; a gate went UNEXECUTED — a FINDING, not a pass). So host-isolation is a PREREQUISITE for the
battery, not a parallel item. The off-box fixes land before the battery so the battery grades the shipped code
(the graded≠shipped lesson from M219). Direct-drive shape per the M215 analogue: live shared infra does not
reward speculative iteration.
**Strategy class:** new-direction
**Distance-to-gate context:** Gate = 8 conditions on a cold reset-to-seed, NO FLAGS. Starting value UNMEASURED on
this code — the last real billion run (M215-era) skipped all three catalog surfaces and predates
M218/M219/M220. Phase C's first battery cycle establishes the baseline.
**Next-tik direction:** iter-02 (first tik) = Phase A, land `GUARD-M221-host-isolation`. Off-box; RED-prove that
a second concurrent cycle is refused with the holder's identity named.

## Close-time decisions (close-milestone, 2026-07-15)

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| D-M221-C1 | **F-M221-06b LANDED at close (Fate 1), not routed to v2.4.** `run-latency.sh` gains a backward-compatible `LATENCY_SCHEME` env (default `http`; `https` for the M220 HTTPS-fronted remote cockpit). | It fits the fixable-inline boundary (one env var, shellcheck-clean, construction-proven) and the task directed landing it in the harden pass if clean. Retires the iter-06 "drive the spec with https by hand" workaround. rext `a0f8615`. | 2026-07-15 |
| D-M221-C2 | **Four tail carries KEEP-DEFERRED-WITH-SIGNOFF → v2.4, finalized at close-release.** F4 (academy grid render defect), BURNIN-M221-dev-public-host, F-M220-4, PROBE-M218-c3-rerun. | M221 is the FINAL milestone → no in-release LAND-NEXT target. Each is non-gate with a concrete reason: F4's fix lives in the `ant-academy` **platform repo** (v2.3's zero-platform-edit constraint forbids it — structurally un-landable, not "no time"); the other three need live infra (dev-stack burn-in / live public-host re-run / the box) that repo-side close work can't provide. The cross-release escape-hatch sign-off is owed at `/developer-kit:close-release` Phase 1b (release scope), which runs next. See `audit-deferrals/deferral-audit-2026-07-15-m221-close.md` (verdict YELLOW). | 2026-07-15 |

## Adversarial review (close-milestone Phase 2c, 2026-07-15)

Two adversarial scenarios were constructed against the milestone-touched code during the final-harden pass; both
are recorded here as the *scenarios considered*, with their responses.

1. **"Does the F1 `exists != populated` bug recur one directory level deeper?"** — `workspaceRootFrom`'s fix
   requires a NON-EMPTY `snapshots/` dir, but its heuristic is `len(ReadDir(snapshots)) > 0`. **Scenario:** an
   interrupted capture leaves `snapshots/<surface>/` present but empty — "populated" to the resolver, yet holding
   no real cache. **Response:** the load-bearing property is not "the resolver is perfect at every depth" but "a
   wrong/empty store can NEVER degrade the catalog SILENTLY." The replay-time `st.List()==0` net catches the
   deeper shadow and emits the D17 loud wrong-root diagnostic (non-fatal). Pinned by
   `TestReplay_StoreWithEmptySurfaceSubdirIsStillLoud` (rext `a0f8615`). A smarter recursive resolver would be a
   design change with no field evidence — routed out of scope, the invariant is protected regardless.
2. **"Does running a test suite directly actually run all its tests?"** — **Scenario:** `python3
   tests/test_reap.py` printed "Ran 21 tests ... OK" while `pytest` collected 41 — a mid-file
   `unittest.main()` silently omitted the 20 adversarial-error-path + suite-honesty classes defined below it (the
   exact fences that guard this milestone's reap work). **Response:** a false all-clear is the release's own D17
   signature hazard turned on the suite that fences it. Moved the block to EOF; direct run 21→41, pytest
   unchanged. Fixed inline (rext `a0f8615`).

---

## RELEASE-SCOPE-DEFER (v2.3 close-release Phase 1b — 2026-07-15)

_Recorded at `/developer-kit:close-release`. The M221 milestone audit
(`audit-deferrals/deferral-audit-2026-07-15-m221-close.md`) dispositioned four non-gate tail carries as
KEEP-DEFERRED-WITH-SIGNOFF → v2.4 and explicitly handed the cross-release escape-hatch sign-off to this phase.
**The user signed off (2026-07-15): accept all four → v2.4.** Two originate at M221; landing spot for all four is
`roadmap-vision.md` under the v2.4 heading. The other two are recorded at their originating milestones (F-M220-4 →
m220 `decisions.md`; PROBE-M218-c3-rerun → m218 `decisions.md`)._

**RELEASE-SCOPE-DEFER: F4 — academy grid renders 0 cards (DEF-M221-05).** The demo's ant-academy home renders
**0** skill-path/course cards even though the local catalog serves **2,705 entries** over HTTP 200 — a
client-side render defect in the academy UI. Origin M221 iter-04, reconfirmed iter-06 (D-M221-06e).
- **Fate-1 (land now) FAILS — structurally, not for lack of time.** The fix is a render-path change in the
  `ant-academy` **platform repo**, and v2.3's HARD constraint is **zero platform-repo edits**. It is out of
  bounds by release construction; a rext-owned demo-patch would first need an ant-academy render investigation
  that has not been done.
- **Fate-2 (drop) FAILS.** Real, reproducible, user-visible; a genuine (cosmetic) gap worth tracking — not
  noise to discard.
- **Fate-3 (defer) is correct.** Non-gate: the 8/8 headline gate is independent of it, and the demo's primary
  surfaces (presenter cockpit + next-web) are fully functional. → **v2.4** (documented known cosmetic gap; likely
  a new sha-pinned demo-patch once the render path is understood).

**RELEASE-SCOPE-DEFER: BURNIN-M221-dev-public-host (DEF-M221-06).** The `/dev-up --public-host` flag (built at
M220, consuming the retired M216 reservation) is fenced + mutation-proven byte-identical on the no-flag path, but
was **never brought up as a real remote dev stack** for a live stability burn-in. Carried M220 → M221 (iters
03–06); not reached.
- **Fate-1 (land now) FAILS.** A burn-in requires **repeated live dev-stack `--public-host` cycles on real
  infra** — not repo-side close work; there is no way to discharge it from the worktree.
- **Fate-2 (drop) FAILS.** The dev-path parity is a shipped capability; leaving it entirely un-exercised live is
  a real (if low) risk worth a tracked future check.
- **Fate-3 (defer) is correct.** Non-gate: v2.3's proven scope is the DEMO path on `billion` (8/8). → **v2.4**
  (live dev-path burn-in when infra is available).
