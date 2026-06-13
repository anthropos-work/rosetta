# M23 — Spec notes

_Technical detail that doesn't belong in `overview.md` (code maps, contracts, edge cases). Accumulates during build._

## Pre-flight audits — §1 (first section of session)

- **Phase 0b KB-fidelity:** YELLOW. Report: `kb-fidelity-audit.md` (2026-06-13). KB-1 (directus_files capture
  wiring) + KB-2 (cms.md M10-gap staleness) tracked in `decisions.md`; both close inside M23 (§4 code + §6 docs).
  No blind area, no implementation-blocking stale claim.

## Topic → doc → code triples (for fast future audits)

| Topic | Knowledge doc | Code |
|---|---|---|
| Demo override env | `corpus/ops/safety.md`, `corpus/ops/snapshot-spec.md` | `stack-injection/gen_injected_override.py` |
| Dev override env-emission | `corpus/ops/snapshot-spec.md` | `stack-core/gen_override.py` |
| Per-stack Directus EnvContract / BaseAddr | `corpus/ops/snapshot-spec.md` | `stack-snapshot/directus/provision.go` |
| directus_files ref capture | `corpus/ops/snapshot-spec.md § Media/blobs` | `stack-snapshot/directus/{media.go,directus.go}` |
| Cross-surface referential closure gene | `corpus/ops/snapshot-spec.md § fidelity gate` | `stack-seeding/dna/snapshot.go` |
| Service Directus dependency truth | `corpus/services/{cms,studio-desk,jobsimulation,next-web-app}.md` | (platform; doc-only) |

## Code map — the M23 surfaces (from the pre-flight read)

- **Two override emitters, opposite conventions** (kept in step):
  - `stack-injection/gen_injected_override.py` (demo) — `build_lines()` already emits a per-service
    `environment:` block (it strips `DIRECTUS_TOKEN=` from every service). Adding `DIRECTUS_BASE_ADDR` is one
    line in the collected `env` list. The UI tier (next-web/studio-desk) has its own per-service env dict in
    `FRONTENDS`.
  - `stack-core/gen_override.py` (dev) — `build_override()` + `to_yaml()` emit ONLY ports/volumes. Needs new
    env-emission plumbing (§1): collect per-service env in `build_override`, render an `environment:` block in
    `to_yaml`. This is the single genuinely-new bit.
- **The in-network Directus address** is `http://directus:8055` (compose service name `directus`, container
  `<stack>-directus-1`, on `app-network`) — NOT `EnvContract.BaseAddr`'s `localhost:<offset>` (that is the
  HOST-side address). cms/jobsimulation/studio-desk containers reach Directus by the in-network service name
  (the same two-address distinction `provision.go` documents for the DSN). DIRECTUS_PUBLIC_BASE_ADDR stays
  `content.anthropos.work` (asset plane).
- **directus_files capture** (§4): `directus_files` is a Directus *system* table, but here we capture only the
  rows the public content references (`media.go ReferencedFilesFilter()` = an OR-of-INs over the public content
  file-ref columns). It is neither scope-bearing (no private/tenant_id) nor pure-reference (capturing it whole
  would pull all 10,340 incl. customer-referenced) nor single-FK parent-scoped. It needs an EXPLICIT raw filter
  on its TableSpec → a new capture mechanism (a `RawFilter`/referenced-files capture kind) that BuildPlan honors
  and the firewall admits. Refs only (blob bytes stay backlog DEF-M10-01).
- **Referential closure** (§5): lean = **full-taxonomy capture** (no node-id subsetting → no dangling problem,
  the simple fallback the corpus already names). The cross-surface gene measures: every taxonomy node-id the
  captured content references (`simulations.skills/job_roles`, `sequences.skills`, etc.) is present in the
  captured taxonomy — distinct from the existing within-surface `OpSnapshotReferential`.
