# M13 — Decisions

_Implementation decisions with rationale. ID scheme: M13-D1, M13-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M13-D1 | `dev-min` size = 10 users / 1 month / 1 org (vs demo small-200's 200/3mo). | Resolves M13-Q3. 10 is the floor that still exercises the role mix (~1 admin + ~6 members + ~3 candidates) so authz/memberships/activity render; smaller rounds the mix to noise. Keeps a fresh dev stack "never empty" without a demo-scale seed cost (<1s). Fixed admin = `dev@anthropos.test` (the local dev login identity → authorized routes 200). | 2026-06-07 |

## Open at design (to resolve during build)
- M13-Q1: dev bring-up heaviness (mitigate: cache-first snapshot, minimal seed).
- M13-Q2: auto-snapshot default-on vs opt-in (lean: default-on, `--no-snapshot`).
- M13-Q3: `dev-min` preset exact size.
