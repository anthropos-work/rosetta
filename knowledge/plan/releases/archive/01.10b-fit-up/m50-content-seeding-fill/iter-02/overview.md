---
iter: 2
milestone: M50
iteration_type: tik
status: in-progress
created: 2026-06-30
---

# iter-02 — tik (baseline sweeps + first seed-cluster fix)

**Type:** tik

## Active strategy reference
TOK-01 (sweep-driven seed-fill of the genuine empties; re-seed-to-iterate; COLD reset reserved for the
exit-gate proof).

## Step 0 — Re-survey
The iter-01 re-diagnosis gives the seed-level gap inventory; the rendered-sweep failing-section count is the
real metric and is unmeasured. So this tik's FIRST action is the baseline sweeps (employee Maya + manager Dan)
to fix the true starting `(failingSections, escapes)`. Target still untouched/meaningful — confirmed (no tik
has run yet).

## Cluster / target identified
TOK-01 named the member-field backfills (location/last_activity/joined_at) as the leading candidate. Final
pick deferred to the post-baseline triage (Phase B) — whichever cluster unblocks the most failing sections.

## Hypothesis
Filling the genuine seed gaps the baseline sweep confirms as failing sections will reduce `failingSections`
toward 0 for the affected vantage.

## Expected lift
A net reduction in `failingSections` for at least the targeted cluster's sections; baseline establishes the
denominator.

## Phase plan
A (baseline employee + manager sweeps) → B (triage by routing table) → C (fix highest-leverage seed cluster in
rext + re-seed demo-1) → D (re-sweep the affected vantage) → E (close).

## Escalation conditions
A failing section that can ONLY be closed by a platform edit → re-scope-trigger (escalate, do not edit the
platform). A surfaced unrelated bug → route forward (Fate-3).

## Acceptable close-no-lift outcomes
If the baseline sweep falsifies a presumed seed gap (the surface renders fine, or its empty is a
runtime-computed / federation artifact that's crawl-scope or a different fix surface), recording that
falsification + routing the real fix is a valid close.

## Baseline measurement (Phase A result)
- **Employee (Maya): GATE MET** — `reachable=59 frontier-EXHAUSTED (cappedAtFrontier=false) failingSections=0
  escapes=0 personaFailures=0 notReachedPages=0` (Playwright 1 passed, 12.0m). The annotation's employee gaps
  (academy link, /profile/activities xp, /home in-progress, /library/skill-paths, logout→cockpit) are ALL
  already closed by M47/M48 + the current seed — **the employee half of the gate needs NO fix.** This is the
  re-diagnosis dividend the orchestration warned about (annotation observed on the OLD stale demo).
- **Manager (Dan): baseline in progress** — all 7 `/enterprise/*` pages + /profile reach 200, 0 FAIL streamed;
  the verdict pends the full library-frontier exhaustion.

## Fix landed (Phase C — the member-field cluster)
The member-field backfill (memberships.joined_at / location / last_activity_date — 0/221 each on demo-1) is
authored + tested in rext `users.go` (COPY + NULL-only idempotent backfill + 2 unit tests; full seeders suite
GREEN incl the M17 re-seed-idempotency contract). It targets the manager `/enterprise/members` annotation gap.
Pending: re-seed demo-1 + the manager re-sweep to measure its effect (and the F1 manifest-strengthening, since
the current members-roster section asserts only names+emails, not these columns).
