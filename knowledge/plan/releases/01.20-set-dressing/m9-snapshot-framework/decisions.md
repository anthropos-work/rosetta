# M9 — Decisions

_Implementation decisions with rationale. ID scheme: M9-D1, M9-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| _(none yet)_ | | | |

## Open at design (to resolve during build)
- M9-Q1: snapshot source — committed golden vs live privileged read (lean: committed golden, M0 discipline).
- M9-Q2: embedding capture — verbatim vs recompute (lean: verbatim, offline/deterministic).
- M9-Q3: snapshot versioning on skiller schema drift (data-DNA `diff` flags it).
