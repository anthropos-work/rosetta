# M23 — Decisions

_Implementation decisions with rationale, numbered `M23-D1`, `M23-D2`, … . Empty at scaffold; filled during build._

## KB-fidelity (Phase 0b, 2026-06-13) — verdict YELLOW

- **KB-1** — snapshot-spec.md § Media/blobs overstates `directus_files` as "always captured + replayed" while
  `media.go` is dead (not wired into `directus.Surface()`). Fate: code wins — M23 §4 wires it (DEF-M21-03 /
  Fate-3). Doc made true by the §4 landing; §6 reconciles the prose. Report: `kb-fidelity-audit.md`.
- **KB-2** — cms.md (~line 13) describes the M10 collection-schema gap as still open ("skips exit 4", "reads
  live from prod") though M21 closed the structure gap + M22 boots the per-stack Directus. Fate: doc wins —
  M23 §6 updates cms.md to the M21/M22/M23 truth (gap closed, instance booted, runtime cutover = M23).

## M23-D1 — the cms-only data-plane cutover (re-point one service, not the tier)

`cms` is the ONLY platform service that talks to Directus directly (HTTP API). jobsimulation reads Directus
THROUGH cms over RPC; next-web reads it through cms/the router. So re-pointing `cms`'s `DIRECTUS_BASE_ADDR` cuts
the WHOLE data plane over — no per-service env sprawl. studio-desk is the one exception (its own
`DIRECTUS_BASE_URL` + write path), handled separately in §3. The in-network address is `http://directus:8055`
(the compose service name on app-network), NOT `EnvContract.BaseAddr`'s `localhost:<offset>` (host-side only) —
the same two-address split provision.go documents for the DSN.

## M23-D2 — studio-desk's static token via bootstrap ADMIN_TOKEN

studio-desk Bearer-auths Directus with a STATIC token (verified vs the dev clone: `src/routes/skillpath.ts`).
Directus 11.6.1 bootstrap reads `ADMIN_TOKEN` and stamps it on the admin user's `token` field (verified vs the
pinned image `bootstrap/index.js:81`) — a Bearer-usable static token. So the per-stack `EnvContract.AdminToken`
(deterministic `local-directus-token-<stack>`) is stamped at bootstrap and handed to studio-desk as
`DIRECTUS_TOKEN`, with no runtime token fetch. `Validate` stays the prod-safety firewall (token-if-present
non-prod; empty OK for the `--check-env` use case); a new `ValidateProvisionable` adds the present-token gate for
the full provision recipe.

## M23-D3 — directus_files as a REFERENCED-SUBSET (reverse-reference closure)

directus_files has no scope column and no forward parent FK — it is referenced BY many public content rows
(simulations.cover, roles.avatar, …). So neither pure-reference (would capture all 10,340 incl. customer-
referenced) nor ParentScopes (a single forward FK) fits. Added a new capture admissibility kind:
`TableSpec.ReferencedSubsetFilter` carries the explicit closure predicate (`ReferencedFilesFilter` — an
OR-of-INs over the public file-ref columns); `firewall.AssertPlan` admits it iff `Filter == ReferencedSubsetFilter`
(mutually exclusive with the other kinds); the post-capture probe counts rows OUTSIDE the closure (0 by
construction). Captures REFS only (the asset plane serves blob bytes from prod); blob BYTES stay backlog
(DEF-M10-01, `BlobBytesAvailable()==false`).

## M23-D4 — replay-clear DELETE/TRUNCATE split (the directus_settings FK)

Discovered live: directus_files is FK-referenced by `directus_settings` (a Directus SYSTEM table OUTSIDE the
captured surface). The replay's bulk `TRUNCATE … RESTART IDENTITY` fails structurally on an FK-referenced table
even when the referrers are empty/NULL — and the referrer can't be co-TRUNCATEd (not in the surface) nor CASCADEd
(would clear the project config). Fix: a `ClearByDelete` flag (TableSpec → manifest → TableRef) makes replay
clear directus_files with `DELETE FROM` (permitted — the referencing columns are NULL on a fresh stack; no
RESTART IDENTITY needed — uuid PK) BEFORE the bulk TRUNCATE of the rest. Verified vs a fresh directus 11.6.1
bootstrap (2026-06-13).

## M23-D5 — referential closure: full-taxonomy capture + the MEASURED cross-surface gene (+ the 1 prod residual)

**The reference format (prod-verified 2026-06-13):** content references taxonomy not by a node-id FK column but
through `directus.sequences.skills` — a JSON array of `{node_id, name}` (e.g. `K-EFFSAL-23FF`) where `node_id`
targets `skiller.skills.node_id`. (`simulations.job_roles` is a JSON array of role NAME strings, not node-ids;
`simulations.skills` is mostly null at the sim level. The load-bearing cross-surface ref is the per-sequence
`skills` node-ids — what `publicJobSimulations.skills` resolves.)

**Full-taxonomy capture (the lean) is ALREADY the state.** The taxonomy surface captures `organization_id IS NULL`
— i.e. EVERY public skill/role, not a content-referenced subset. So the only way a content ref can dangle is if it
points at a node that isn't public (a customer-scoped skill) — which the firewall MUST NOT capture. No subsetting
change is needed; the lean is satisfied by construction.

**The cross-surface closure GENE (new, M23):** `OpSnapshotCrossSurfaceClosure` + `FidelityProbe.CrossSurfaceDangling`
measure — against the REPLAYED directus↔skiller pair — every content-referenced node-id resolves in the replayed
taxonomy; a non-zero dangling count fails the gene and names a sample node. Distinct from the WITHIN-surface
`OpSnapshotReferential`. Wired into the `content` gene (criticality `standard`, so it surfaces in the overall score
but does NOT block the CRITICAL gate, and `measure-snapshot` is not run in the bring-up so it never blocks UP).

**The 1 genuine prod residual (surfaced, not hidden):** prod has exactly 1 dangling node — `K-AIFUNX-E658`,
referenced by 2 public published sims ("Explaining Artificial Intelligence to a New Intern", "AI Consultant:
Multi-Department Prompt Coaching") but existing ONLY as a customer-scoped skill (org `f9e88e97…`). It CANNOT be
closed: capturing the customer node would breach the tenant firewall; editing prod is forbidden (zero platform/prod
edits). This is a prod DATA-QUALITY inconsistency (a public sim mis-referencing a customer skill), now MEASURED +
named by the gene rather than silently producing an empty picker. The honest resolution is the gene's report; the
fix (re-tagging or removing that skill ref) is a prod data correction the operator owns, outside tooling scope.
