# M48 — spec notes

_Technical notes accumulate here during build (file:line surfaces, schema findings, the drift gap-list)._

## Pre-flight audits — S1/S2 (corpus re-ground)

**Phase 0b KB-fidelity verdict: GREEN-by-design** (recorded 2026-06-29).

M48 is the inverse of the usual gate: its *deliverable* is re-grounded docs, and the contract it reads is the
**current platform code** (the M47-confirmed-current clones), not knowledge docs. "Stale corpus docs" is not a
blind area that blocks M48 — it is M48's job. The one genuinely-undocumented surface (member-AI-readiness) is the
milestone's load-bearing **new** deliverable (S2), authored from the code. No blocker; proceed.

## Investigation (3 parallel read-only agents, launched 2026-06-29)
- **AI-readiness backend data model** (app/jobsimulation/skiller/skillpath/cms) — the seeding contract for M51.
- **AI-readiness frontend + manager dashboard** (next-web `ai-readiness/`) — the surface map.
- **Corpus drift survey** (corpus/architecture + corpus/services vs current clones) — the S1 material-lag gap list
  + the AI-readiness doc placement + the ant-academy stale-claim confirmation.

_(Findings land here as the agents report.)_
