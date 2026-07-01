# M303 Progress

Iterative milestone — closed-per-resource loop (built by `/developer-kit:build-mstone-iters`). Scope detail in
[`overview.md`](overview.md). The gate + per-iter close bar are the source of truth.

## Iters ledger

_(populated per-iter by `build-mstone-iters` — one row per resource-attempt; each iter targets ONE UC from the
R1 catalog and closes on the 6-point per-iter bar in `overview.md`.)_

| Iter | Date | UC | Resource | Status | Notes |
|------|------|----|----------|--------|-------|
| | | | | | |

## R1 read catalog — per-resource gate tracker

FIRST-USABLE reads (UC1–UC4) close first; FIRST-USEFUL (UC5–UC7) after. Each row closes when all 6 per-iter
bar points are green.

- [ ] **UC1 — `GET /v1/people/members`** (`people.member.list`) — the roster read (FIRST-USABLE).
- [ ] **UC2 — `GET /v1/people/members/{id}`** (`people.member.get`) — the profile read (FIRST-USABLE).
- [ ] **UC3 — `GET /v1/learning/skill-paths`** (`learning.skill-path.list`) — the assigned-path list
  (FIRST-USABLE).
- [ ] **UC4 — `GET /v1/learning/skill-paths/{id}`** (`learning.skill-path.get`) — the path detail + progress
  (FIRST-USABLE).
- [ ] **UC5 — `GET /v1/learning/sessions`** (`learning.session.list`) — the session history read.
- [ ] **UC6 — `GET /v1/verification/verified-skills`** (`verification.verified-skill.list`) — the verified-skill
  chart source.
- [ ] **UC7 — `GET /v1/audit/events`** (`audit.event.list`) — the customer's own audit trail (feature-flagged
  in R1 — turns on when the customer opts in).

## Exit-gate acceptance (measured at close)

- [ ] All 7 UCs green on the per-iter bar (rows above).
- [ ] 5 consecutive integration runs with 0 cross-tenant leakage across all 7 UCs.
- [ ] `openapi.yaml` published + linted (Spectral) with 0 must-fix.
- [ ] The 7 audit-row shapes documented + verified against `customer_api.audit_events`.

**Status:** `planned` — not yet started. Next: `/developer-kit:work-mstone-iters` (M303) after M302 closes.
