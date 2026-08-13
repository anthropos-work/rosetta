# Auditor C — 7 files / 1505 lines (read-confirmed 1505) — 2 blockers, 13 minors

## B-C1 `corpus/services/roadrunner.md:23-25` — jobsimulation listed as removed from repos.yml + compose
FALSE. jobsimulation IS in both: `platform/repos.yml:17`, `platform/docker-compose.yml:83` with
`profiles: [graphql, jobsimulation, all]` at :140 — it starts on a bare `make up`. Only intelligence /
skiller / skillpath / chronos are absent from both. Contradicted by two sibling docs at HEAD
(`architecture_overview.md:188`, `services/README.md:20-21`) AND by roadrunner.md's own :26.
CLASS: repair-prose, and CROSS-FILE (a sibling says the opposite).

## B-C2 `corpus/services/messenger.md:110` — anchor only
`assignments.go:815` is `JobSimulationId:` inside `getEmailNotificationForSimulation` (a SIMULATION
email lookup). The skill-path CMS read is `:828` (`h.cms.GetSkillPath`). Substantive claim TRUE;
anchor lands 13 lines off in the wrong function.

## MEASUREMENT (C graded MINOR; flagged for re-grade) — the multi-tenancy fence, arch_overview:288-291
- "135 schemas": **135 ✓ — and BOTH derivations now agree at 135.** This REFUTES the "denominator is
  contested (135 vs 112)" premise of D-M257x-39-3.
- "30 using OrganizationMixin{}": 30 ✓
- "31 auto-filter": measured **32** — `organization.go:56` declares an org-filtering Policy() the doc
  omits. INDEPENDENT REPRODUCTION of the iter-39 repairer's refusal (D-M257x-39-3).
- "16 carry organization_id with no policy at all": measured **24** — 17 by plain field decl PLUS 7 via
  `OrganizationIDMixin{}` (category, jobrole, similarity, skill, specialization, studio_document,
  studio_task). That mixin is a PLAIN nullable column, and a naive `field.UUID("organization_id"` grep
  is BLIND to all 7. **Errs in the dangerous direction — understates the unpolicied set.**
- Home derivation lives in `security_compliance.md#layer-1-database` = auditor E's file. CROSS-CHECK E.

## UNVERIFIABLE: gh/archive status; alignment SCORES (runners can't build); prod runtime; db-backup
near-entirety (0 hits for db-backup in platform/, exit 0); chronos body (fenced historical); private Go modules.
