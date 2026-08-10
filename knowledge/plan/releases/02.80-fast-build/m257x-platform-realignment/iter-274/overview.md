---
iter: 274
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: FIX-M257x-272-succession-hero-has-no-qualifying-surface
---

# iter-274 — the threshold that dropped the hero, measured per role

**Type:** tik, under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Step 0 — re-survey (mandatory, before targeting)

iter-272 established **that** the hero appears on no succession surface; iter-273 established that she is
the **only** thing between the suite and clause 2 (`30 live / 1 failing / 0 error`, binding).

Re-survey at open changes the target's shape, and the correction belongs here rather than in the fix:
**iter-272's account of the at-risk exclusion was incomplete.** It said she is excluded *because she is
well-matched* (`fit 87`). Reading `buildAtRisk` (`succession.go:892`) shows fit is **not** the gate. Both
signals that named her in the M256 iter-14 measurement — *"Rare skill held only by this person"* and
*"In fragile role: DevOps Engineer"* — are drawn from role sets built behind the **same** guard, stated
twice:

```go
if role.RiskScore < 50 { continue }
```

So a member is eligible for those two signals only through a role whose **`RiskScore ≥ 50`**. Her fit
governs whether she is a *successor*; it does not govern at-risk membership. **The at-risk question is a
per-role threshold question**, and it has never been measured.

The payload iter-272 captured already carries `riskScore`, `structuralRisk`, `criticalityMultiplier` and
`isKey` for **all 12 roles** — so this is a read of an artifact already on disk, not a new run.

## Cluster / target identified

The **per-role risk figures for Org A**, and specifically where `DevOps Engineer` — the hero's role — sits
relative to the hard-coded `50`. This names the lever the fix must move, and distinguishes a lever that is
**ours** (seed-side, if `structuralRisk` is driven by seeded composition) from one that is **not** (a
platform constant, which v2.8 forbids editing).

## Hypothesis

`DevOps Engineer` sits **below 50**, and the roles that do produce at-risk signals sit above it — i.e. the
hero's role fell under the threshold while others stayed over. No mechanism for *why* is predicted; the
distribution is the measurement.

## Expected lift

- The 12 roles ranked by `riskScore`, with the `≥ 50` line drawn and `DevOps Engineer` located on it.
- The fix lever **named and classified** ours-vs-platform, so iter-275 implements rather than investigates.
- No suite-metric movement promised. Clause 2 moves when the fix lands, not when it is specified.

## Phase plan (declared multi-step — the tripwire counts UNPLANNED lines only)

1. Seal these pre-registrations (first commit).
2. Extract per-role `riskScore` / `structuralRisk` / `criticalityMultiplier` / `isKey` from **both**
   captured payloads (iter-272 runs 1 and 2) — two independent captures of the same seed.
3. Locate `DevOps Engineer`; identify which roles clear 50 and what separates them.
4. Classify the lever; record the fix decision with its alternative.

## Out of this iter's planned scope (declared, so the tripwire is clean)

- **Implementing** the fix (seed change or expectation change), its tag, pin bump and re-run. iter-275.
- Gate clause 5 and the inherited route queue.

## Escalation conditions

- **No platform edit** under any finding (v2.8). If the only lever is a platform constant, that is a
  finding to report, not a change to make.
- No new stack runs are needed; if one becomes necessary the iter says so rather than quietly widening.
- If the two captured payloads disagree, the disagreement is the result — a projection that differs
  across two resets of the same seed would invalidate the determinism iter-272 recorded.

## Acceptable close-no-lift outcomes

The lever proving to be a platform constant — i.e. the fix cannot be seed-side — is a complete result. It
would settle the seed-vs-expectation fork by elimination rather than by preference, which is stronger than
the recommendation iter-272 recorded.

## Pre-registrations (sealed in this iter's FIRST commit, before any measurement)

- **PR-1 — `DevOps Engineer` is below the line.** Its `riskScore < 50`. *Refuted by:* `≥ 50`, which would
  mean the guard is not what excludes her and iter-272's account is wrong a second time.
- **PR-2 — some roles ARE above it.** At least one of the 12 has `riskScore ≥ 50` — otherwise no member
  could carry the two role-derived signals at all, yet 27 do. *Refuted by:* all 12 below 50.
- **PR-3 — the two payloads agree exactly.** Every per-role figure is identical across iter-272's two
  captures. *Refuted by:* any difference.
- **PR-4 — `riskScore = round(structuralRisk × criticalityMultiplier)`.** The relation holds for all 12
  roles (the `Administrator` pair `53 → 45` at `0.85` is the reason to suspect it). *Refuted by:* any role
  where it does not hold — which would mean the score has an input not in the payload.
