# M210 — Spec notes

## Pre-flight audits — Section 1 (Adopt + validate)
- **Verdict: YELLOW** (proceed). Report: `kb-fidelity-audit.md`.
- Ground truth = M209 landed rext (`.agentspace/rosetta-extensions`): **89 `public.*` query refs, 1 residual skiller.***; state.md records M209 CLOSED, 0 skiller.<table> queries, tagged `quick-change-m209@2f06e78`.
- Release-branch corpus has **33 stale `skiller.<table>` refs** (snapshot-spec 19, db-query SKILL 8, db-access 7, stories-spec 3, safety 1, directus-local 1) — these ARE the milestone's fix-list, not a blocker.
- Applies to all sections (single subsystem = the docs corpus); reuse across Sections 2-6 per §"Audit reuse".

## Topic → doc → code triples (fast-start for future audits)
- taxonomy schema → snapshot-spec/safety/stories-spec/seeding-spec/recipe-snapshot-world/coverage-protocol → M209 rext (public.*)
- subgraph/container/RPC → CLAUDE.md + service_taxonomy + skill files → M208-verified merged compose (4 subgraphs, no skiller container, SKILLER_RPC_ADDR=backend:8083)
- member surface → profile-completeness-spec.md → seeders (public.*)

## Adopt/validate the docs branch (architecture half)
## profile-completeness-spec.md missed-file fix
## Rext-facing tooling-doc body flips (6 docs)
## db-access ↔ tooling reconcile
## Skill-file sweep + CLAUDE.md catalog
