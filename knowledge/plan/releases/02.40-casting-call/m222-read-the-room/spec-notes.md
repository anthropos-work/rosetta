# M222 — Spec notes

_Technical detail accumulated during the milestone: file:line maps, command transcripts, measured numbers._

## Hiring model (for `corpus/services/hiring.md`)
_`is_hiring`, the `candidate` role auto-assignment, hiring sims ↔ `job_position`, `SIMULATION_TYPE_HIRING`, the
comparison read-path, the `isHiringOrg` publicMetadata derivation, the `isEnterprise` divergence._

## Render-probe results (BA-1 / BA-2 / BA-3 / BA-6)
_Throwaway hand-seed against the live dockerized `apps/web` on the `billion` substrate; per-blind-area findings._

## The seeder-output contract (the GO/NO-GO record)
_Does the comparable score render from `sessions.score` alone, or does it need a `validation_*`/eval row per session?
The exact contract M223/M224 build against._

## `is_hiring` gate thread
_The blueprint `Org.IsHiring` field + `narrative: hiring` + `org.go` one-value change + the deterministic OrgID._
