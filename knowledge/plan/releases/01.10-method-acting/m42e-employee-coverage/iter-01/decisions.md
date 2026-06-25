# iter-01 decisions

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| D1 | Harness home = under `stack-verify/e2e/` (not a new `stack-coverage` section) | `stack-verify/e2e/` already pins `@playwright/test ^1.49.0` + a working config + an unauthenticated smoke test; the coverage sweep is the authenticated evolution of it and reuses verify's offset/project/scope plumbing. No new section needed. | 2026-06-25 |
| D2 | Crawl = pure in-app nav-link BFS (not a static route manifest) | The gate requires escape-detection; a route manifest can't see a nav that escapes the demo (external links are invisible to a route list). The crawl observes actual rendered same-origin links + nav chrome. | 2026-06-25 |
| D3 | Non-empty assertion = two-tier (Tier-1 text-density floor + error-sentinel default; Tier-2 per-section selectors escalated per-page) | Tier-1 is cheap and catches the dominant empty mode without over-fitting brittle selectors to the whole sweep; Tier-2 is added per-page only when Tier-1 false-passes/-fails, keeping the sweep robust. | 2026-06-25 |
| D4 | Login seat for employee vantage = `maya-thriving` (member); manager = `dan-manager` (admin) | Confirmed against the live demo-3 roster (`/roster/roster.json`): `maya-thriving` is `org_role: member`, the default seat (order[0]); `dan-manager` is `admin`. The harness selects the seat via `POST /v1/demo/select` + the handshake `__clerk_identity`. | 2026-06-25 |
