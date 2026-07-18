---
iter: 02
milestone: M226
iteration_type: tik
status: closed-fixed
date: 2026-07-17
---

# iter-02 (tik-01) — substrate cutover + first cold bring-up + first 7-condition measurement

**Active strategy reference:** TOK-01 `reprove-hiring-on-billion` (the initial strategy; iter-02 is its
next-tik direction, verbatim).

**Step 0 — re-survey:** TOK-01 was just authored (iter-01); its next-tik target is fresh. Recon confirms billion
still runs the stale v2.3.2 demo-1 (nothing changed it). Target unchanged, no substitution.

## Cluster / target identified
Get a **casting-call** demo standing on billion from a cold reset-to-seed and take the first 7-condition
measurement. The stale v2.3 substrate + panorama rext tag is the gate blocker; the first tik's real cost is a
clean teardown + a full casting-call rebuild (2-app hiring image), not just a re-run.

## Hypothesis
A clean cold `up-injected.sh 1` (no flags) on billion at rext `casting-call-m225-harden` reproduces the local
M225 hiring proof (Meridian Talent 4th org + 3 workforce orgs + the recruiter comparison). Measuring the
7-condition gate from this Mac reveals what breaks cross-machine (R1 render / R4 latency / memory).

## Expected lift
0/7 → several conditions GREEN. Expect the structural ones (org present + 5/45 counts, reads-as-hiring,
coexistence, 0 platform edits) to pass on the first clean build; R1 (recruiter render / ≥40 rows) and R4
(recruiter p95 latency) are the live-only risks; billion's 7.3 GiB RAM (2-app demo) is the sharpest infra risk.

## Phase plan (protocol: verification.md + coverage-protocol.md + latency-budget.md)
1. Cold teardown of the stale v2.3.2 demo-1 (`rosetta-demo down 1 --purge` — reaps containers + cockpit +
   academy + serve); verify base ports freed + no survivor **from this Mac**.
2. Cutover: pin `.agentspace/rext.tag` + checkout the rext consumption clone to `casting-call-m225-harden`
   (confirm `sections`↔`harden` is test-only); ensure the platform clone carries `apps/hiring`.
3. Default cold `up-injected.sh 1` (NO FLAGS) — remote-reach default-on; run synchronously (tethered ssh),
   never detached-on-billion; heartbeat by file-write/process activity.
4. Measure the 7-condition gate FROM THIS MAC (peer): counts, recruiter comparison rows/sim, candidate profiles,
   reads-as-hiring, recruiter p95 click→ACCESS, cockpit coexistence, git-clean clone.
5. Attribute every failing condition to its surface before any fix.

## Escalation conditions
- An un-patchable platform surface (no env/config/compose/demo-patch seam) → ESCALATE (never a platform edit).
- OOM that can't be worked around under the memory budget → user-blocker.
- A demo-patch sha-drift needing re-pin → tooling fix within-iter or route forward.

## Acceptable close-no-lift outcomes
If the bring-up fails (build/OOM/wiring) and is attributed with falsification evidence (the surface + arithmetic
signature named, fix routed to the next iter), that is a complete diagnostic cycle → closed-no-lift.

## Close
See `progress.md`.
