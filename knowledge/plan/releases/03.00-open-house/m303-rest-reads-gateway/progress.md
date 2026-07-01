# M303 Progress

Iterative milestone — closed-per-resource-family loop (built by `/developer-kit:build-mstone-iters`). Scope
detail in [`overview.md`](overview.md). The gate + per-family close bar are the source of truth. Spec §4.2 (the
9-product / 35-resource / ~44-endpoint catalog at Talk-to-Data parity) is the authoritative resource list.

## Iters ledger

_(populated per-iter by `build-mstone-iters` — one row per family-attempt; each iter targets ONE resource-family
from the R1 catalog and closes on the 7-point per-family bar in `overview.md`.)_

| Iter | Date | Family | Endpoints | Status | Notes |
|------|------|--------|-----------|--------|-------|
| | | | | | |

## R1 read catalog — per-family gate tracker (Talk-to-Data parity)

The 7 FIRST-USABLE endpoints (∗) close first; the remainder close under the same per-family gate. Each row
closes when all 7 per-family bar points are green (spec.md §5.7 + overview.md).

### Product 1 — People
- [ ] **`GET /v1/people/organization`** ∗ (`people.organization`) — org + `max_level` (CR6 anchor).
- [ ] **`GET /v1/people/members`** ∗ + **`GET /v1/people/members/{member_id}`** ∗ (`people.member`) — roster.
- [ ] **`GET /v1/people/members/{member_id}/skills`** ∗ (`people.member.skill`) — mapped vs verified, org scale
  (CR5, CR6, CR7).
- [ ] `GET /v1/people/members/{member_id}/certifications` (`people.member.certification`).
- [ ] `GET /v1/people/members/{member_id}/educations` (`people.member.education`).
- [ ] `GET /v1/people/members/{member_id}/experiences` (`people.member.experience`).
- [ ] `GET /v1/people/members/{member_id}/languages` (`people.member.language`).
- [ ] `GET /v1/people/members/{member_id}/target-roles` (`people.member.target-role`).
- [ ] `GET /v1/people/members/{member_id}/tags` (`people.member.tag`).
- [ ] `GET /v1/people/members/{member_id}/profile-history` (`people.member.profile-history`) — CR13 self-scoping.
- [ ] `GET /v1/people/teams` + `GET /v1/people/teams/{team_id}` (`people.team`).
- [ ] `GET /v1/people/invitations` + `GET /v1/people/invitations/{invitation_id}` (`people.invitation`).
- [ ] `GET /v1/people/companies` (`people.company`).

### Product 2 — Assignments
- [ ] `GET /v1/assignments` + `GET /v1/assignments/{assignment_id}` (`assignments.assignment`).
- [ ] `GET /v1/assignments/{assignment_id}/sessions` (`assignments.assignment.session`).
- [ ] `GET /v1/assignments/organization-roles` (`assignments.organization-role`).
- [ ] `GET /v1/assignments/organization-target-roles` (`assignments.organization-target-role`).

### Product 3 — Simulations
- [ ] **`GET /v1/simulations/sessions`** ∗ + **`GET /v1/simulations/sessions/{session_id}`** ∗
  (`simulations.simulation-session`) — CR4 completed-definition, CR6 org-scale score.
- [ ] `GET /v1/simulations/sessions/{session_id}/recordings` (`simulation-session.recording`).
- [ ] `GET /v1/simulations/sessions/{session_id}/interactions` (`simulation-session.interaction`).
- [ ] `GET /v1/simulations/sessions/{session_id}/realtime-calls` (`simulation-session.realtime-call`).
- [ ] `GET /v1/simulations/sessions/{session_id}/code-submissions` (`simulation-session.code-submission`).
- [ ] `GET /v1/simulations/sessions/{session_id}/anticheat-results` (`simulation-session.anticheat-result`).
- [ ] `GET /v1/simulations/sessions/{session_id}/activity-events` (`simulation-session.activity-event`).
- [ ] `GET /v1/simulations/sessions/{session_id}/task-checks` (`simulation-session.task-check`; sub_checks embedded).
- [ ] `GET /v1/simulations/sessions/{session_id}/conversation-extractions` (`simulation-session.conversation-extraction`).
- [ ] `GET /v1/simulations/sessions/{session_id}/interview-extractions` (`simulation-session.interview-extraction`).
- [ ] `GET /v1/simulations/sessions/{session_id}/validation-results` (`simulation-session.validation-result`).
- [ ] `GET /v1/simulations/sessions/{session_id}/validation-attempts` (`simulation-session.validation-attempt`; skill/criterion/check results embedded).
- [ ] `GET /v1/simulations/feedback` (`simulations.simulation-feedback`).

### Product 4 — Learning
- [ ] **`GET /v1/learning/skill-path-sessions`** ∗ (`learning.skill-path-session`) — the training-pulse read.

### Product 5 — Catalog
- [ ] `GET /v1/catalog/simulations` + `GET /v1/catalog/simulations/{simulation_id}` (`catalog.simulation-template`; CR11 `?language=`).
- [ ] `GET /v1/catalog/skill-paths` + `GET /v1/catalog/skill-paths/{skill_path_id}` (`catalog.skill-path-template`).

### Product 6 — Taxonomy
- [ ] `GET /v1/taxonomy/skills` (`taxonomy.skill`; public + org-custom; CR10 + CR11).
- [ ] `GET /v1/taxonomy/job-roles` (`taxonomy.job-role`; public + org-custom; CR10 + CR11).
- [ ] `GET /v1/taxonomy/world-languages` (`taxonomy.world-language`).

### Product 7 — Academy
- [ ] `GET /v1/academy/series` (`academy.series`).
- [ ] `GET /v1/academy/skill-paths` (`academy.skill-path`; CR14 `lifecycle=published`).
- [ ] `GET /v1/academy/chapters` + `GET /v1/academy/chapters/{slug}` (`academy.chapter`; `?locale=`).
- [ ] `GET /v1/academy/progress` (`academy.progress`).

### Product 8 — AI Readiness
- [ ] `GET /v1/ai-readiness/live` (`ai-readiness.live`; CR12 the "right now" read).
- [ ] `GET /v1/ai-readiness/cycles` + `GET /v1/ai-readiness/cycles/{cycle_id}` (`ai-readiness.cycle`).
- [ ] `GET /v1/ai-readiness/cycles/{cycle_id}/snapshots` (`ai-readiness.cycle.snapshot`; CR12 frozen).

### Product 9 — Audit
- [ ] **`GET /v1/audit/events`** ∗ (`audit.audit-event`) — the customer's own API-usage audit trail.

## Read-contract rule matrix (CR1–CR15 — spec §4.5)

- [ ] **CR1 Principal-scoping** — middleware in place; every endpoint asserted.
- [ ] **CR2 Soft-delete exclusion** — repository predicate + fixture test on member/team/skill.
- [ ] **CR3 Active-member definition** — shared predicate; contract test.
- [ ] **CR4 Completed-sim definition** — shared predicate; fixture matrix (running/ended/timedout/discarded).
- [ ] **CR5 Mapped ≠ Verified** — response shape + `verified ⊆ mapped` invariant test.
- [ ] **CR6 Org-scale + `max_level`** — level-normaliser + `levels_count = 7` fixture.
- [ ] **CR7 Skill-level source column** — repository lint + contract test.
- [ ] **CR8 Forbidden stale tables** — CI grep-gate on `internal/customerapi/`.
- [ ] **CR9 Person identifier = user UUID** — route contract + 404 test for a membership-PK call.
- [ ] **CR10 Catalog resolution (human labels)** — response shape test.
- [ ] **CR11 Localization** — `?language=` fixture with partial translations.
- [ ] **CR12 AI Readiness live ≠ frozen** — distinct-handler test + shape divergence assert.
- [ ] **CR13 Profile-history self-scoping** — non-admin principal test.
- [ ] **CR14 Academy visibility** — `draft` fixture invisible.
- [ ] **CR15 Read-only R1** — HTTP-method allow-list; `POST /v1/people/members` returns 405.

## Exit-gate acceptance (measured at close)

- [ ] All ~44 endpoints green on the per-family bar (rows above).
- [ ] **5 consecutive integration runs with 0 cross-tenant leakage across the full ~44-endpoint surface.**
- [ ] All 15 CR1–CR15 rule tests green; CR7 + CR8 static-lint CI gates enforced.
- [ ] The 7 FIRST-USABLE ∗ endpoints exercised end-to-end in one integration script under one minted API key.
- [ ] `openapi.yaml` published + linted (Spectral) with 0 must-fix.
- [ ] The ~44 audit-row shapes documented + verified against `customer_api.audit_events`.

**Status:** `planned` — not yet started. Next: `/developer-kit:work-mstone-iters` (M303) after M302 closes.
