# M223 — Spec notes

_Technical detail accumulated during the milestone: file:line maps, command transcripts, measured numbers._

## The 4th story entry
_`narrative: hiring`, `RoleMix` ≈ 0.1 admin / 0.9 candidate, hero-trio placeholder._

## HiringConfigSeeder + the type-aware hiring-sim reader
_The 5 shared HIRING sims; the `type=HIRING AND job_position NOT NULL` pattern query; the disjoint reserved tail._

## Snapshot extension — directus.job_position replay
_Replaying all 443 public rows + pinning the 5 chosen HIRING sims; the digest/capture-column changes._

## The candidate-assessment funnel seeder
_Resolve 5 shared sim refs once; MOST candidates on all 5, SOME assigned-not-taken; the differentiated score
distribution; the M219 skill-ladder + closure wiring._

## reset / closure / isolation wiring
_New hiring rows into `resetTables`; `datadna measure-closure`; the `isolation.Guard` audit._
