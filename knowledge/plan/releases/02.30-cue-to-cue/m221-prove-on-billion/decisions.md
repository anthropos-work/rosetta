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
