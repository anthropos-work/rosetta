---
milestone: M51
slug: ai-readiness-org
version: v1.10b "fit-up"
milestone_shape: iterative
status: planned
created: 2026-06-29
last_updated: 2026-06-29
complexity: large
exit_gate: "On the live demo, a curated 200-person 3rd org renders the AI-readiness manager dashboard ENABLED with ~80% of members having completed all 3 onboarding/evaluation steps; one hero shows STARTED and one shows COMPLETED — proven by the coverage gate."
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md
delivers: corpus/ops/demo/stories-spec.md (the 3rd AI-readiness story + its seeder)
issues: user-requested AI-readiness showcase org (redeems the M201 member-AI-readiness false-negative)
---

# M51 — AI-readiness showcase org

## Goal
Add a **3rd Story** to the Stories & Heroes model — an organization **curated for AI readiness** — that
demonstrates the **member-AI-readiness flow** end-to-end on the manager dashboard. This is the curated showcase for
the very feature M201's verify reported as a false-negative (shipped in prod, invisible to the stale clones).

## Exit gate
On the live demo:
- a **200-person org** with the **AI-readiness manager dashboard ENABLED** (org setting / feature flag);
- **~80% of the 200 members have completed all 3 steps** of the onboarding/evaluation — the manager dashboard
  renders the high-completion AI-readiness funnel + scores;
- a **hero trio**: a **manager** hero (views the dashboard), an employee hero who has **COMPLETED** the readiness
  onboarding, and an employee hero who has **STARTED** it (in-progress on the dashboard's special onboarding
  element);
— **proven by the M42 semantic coverage gate** (manager vantage) on the 3rd org, 0 prod-eject escapes.

## Why iterative (not section)
The feature's **data model is unknown** — it was invisible to the stale clones and is being mapped fresh: how the
**dashboard is enabled** (org setting vs PostHog flag), which **tables** back the **3-step onboarding/evaluation**,
and how **started vs completed** states are stored. The seeder is not writable until that model is reverse-
engineered (the M48 contract). The goal is crisp; the path is exploratory.

## Iteration protocol
`corpus/ops/demo/coverage-protocol.md` — drives the same observe → fix → re-measure loop, against the
**M48-documented AI-readiness contract**.

## Candidate scope (emergent — gated on the M48 data-model map)
- A **3rd Story** in `stories.seed.yaml` / `stack.stories.yaml`: org "AI Readiness" (size 200), the hero trio, the
  narrative.
- The **AI-readiness enablement** (the org-level setting/flag the dashboard keys on).
- A **3-step onboarding/evaluation seeder** writing the funnel for 200 members at ~80% all-3-complete, with the two
  named heroes pinned to **started** and **completed** states.
- **Cockpit** jump-to links — the 3rd org's heroes land on the AI-readiness dashboard / their onboarding element.
- **ant-academy course content + the hero academy menu-link + a non-anonymous academy session** *(annotated by
  M50 close, 2026-06-30 — Fate-3 handoff; the F6 item from M50's candidate fix surface).* The field review flagged
  the academy as `0 chapters / 0 skill-paths`, no menu link, and an anonymous session when reached directly. M50's
  M42 gate is MET both vantages WITHOUT the academy (it's not on a coverage-gate path), so M50 routed it here — the
  academy content/wiring is a seeding/content surface, M51's domain. **The academy AI chat stays
  documented-as-absent** (gated by the AI-keys policy DECIDED at M50: documented-as-absent — see
  [`corpus/ops/secrets-spec.md`](../../../../../corpus/ops/secrets-spec.md)); the course-content + menu-link +
  non-anonymous-session are the seedable/wirable part. Decide the academy content surface (shared-Directus replay vs
  a dedicated academy snapshot surface) here.

## Open questions (resolve during build)
- Enablement mechanism (org setting vs feature flag) — from the M48 contract.
- The completion distribution shape beyond the 80% headline (per-step drop-off realism).
- Whether the manager hero is a distinct seat or the existing Dan persona re-homed — *lean:* a distinct 3rd-org
  manager (keeps the orgs cleanly separable for the cockpit).

## Depends on
**M48** (the documented member-AI-readiness contract — the seeder's spec) + **M50** (healthy seeding machinery +
the cold reset-to-seed lifecycle). **Parallel with:** none (single demo stack).

## Re-scope trigger
If enabling the dashboard or seeding the funnel **cannot** be done without a **platform edit** → **escalate**
(`unimplementable-without-platform-edit`), never edit the platform.

## KB dependencies (read as contract)
- **M48's member-AI-readiness contract** (the load-bearing input), `corpus/ops/demo/stories-spec.md` (the Stories &
  Heroes model + the 7-table chain), `corpus/ops/demo/coverage-protocol.md`, `corpus/ops/seeding-spec.md`.

## Delivers
- **→ rosetta-extensions:** the 3rd-org story + the AI-readiness onboarding/evaluation seeder + cockpit links,
  tagged `fit-up-m51`.
- **→ rosetta:** `corpus/ops/demo/stories-spec.md` — the 3rd AI-readiness story + its seeder documented.
