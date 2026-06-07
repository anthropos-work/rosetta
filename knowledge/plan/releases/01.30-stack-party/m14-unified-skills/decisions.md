# M14 — Decisions

_Implementation decisions with rationale. ID scheme: M14-D1, M14-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M14-D1 | Hard rename, no back-compat aliases | user 2026-06-07 — clean break; update every in-repo reference | 2026-06-07 (user) |

## Open at design (to resolve during build)
- M14-Q1: how much first-time setup folds into `dev-up`.
- M14-Q2: target-detection UX (`stack-seed dev-1` vs `--stack dev-1`).
- M14-Q3: generic `stack-up`/`stack-down`? (lean: no).
