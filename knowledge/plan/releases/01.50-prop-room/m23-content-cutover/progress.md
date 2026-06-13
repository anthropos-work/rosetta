# M23 — Progress

**Status:** in progress. **Shape:** section.

## Section checklist
_One checkbox per concrete deliverable from `overview.md` § Scope. Code lands in the `rosetta-extensions`
authoring copy (tag `prop-room-m23` at close); docs land in the `rosetta` worktree._

- [x] **§1 Dev env-emission plumbing** (ext) — grow `stack-core/gen_override.py` to emit per-service
  `environment:` blocks (today it emits only ports/volumes — the single genuinely-new bit of plumbing).
  _Landed: `directus_consumer_env()` + env-emission in `build_override`/`to_yaml`; dev `--with-directus`
  re-points cms `DIRECTUS_BASE_ADDR`→in-network instance + strips prod token; asset plane untouched. ext
  `3506976`. +9 tests, 19/19 green._
- [x] **§2 `DIRECTUS_BASE_ADDR` re-point — demo + dev** (ext) — inject the local-Directus data-plane address
  (`http://directus:8055`, the in-network compose service) into the Directus-consuming services; keep
  `DIRECTUS_PUBLIC_BASE_ADDR` on prod (asset plane stays real); extend the prod-token strip to opted-in dev.
  _Landed: dev side in §1 (ext `3506976`); demo side ext `b0f3945` (cms re-point gated on with_directus, asset
  plane untouched, prod-read path preserved on --no-local-content). Both bring-ups already thread the flag.
  82/82 demo + 19/19 dev green._
- [x] **§3 studio-desk local instance + minted admin token** (ext) — point studio-desk's `DIRECTUS_BASE_URL`
  at the per-stack instance + a locally-minted admin token so its skill-path writes target local, never prod.
  _Landed ext `5218953`: `EnvContract.AdminToken` (deterministic) → bootstrap `ADMIN_TOKEN` (stamps the
  admin's static token, verified vs the pinned directus image) → studio-desk `DIRECTUS_BASE_URL` + that token
  on local-content. `ValidateProvisionable` adds the present-token gate; `Validate` stays the prod firewall.
  +12 tests green._
- [x] **§4 `directus_files` ref capture** (ext, Fate-3 from M21/DEF-M21-03) — wire the dead `media.go`
  (FileRefColumns/ReferencedFilesFilter) into `directus.Surface()` via a `directus_files` TableSpec so captured
  content rows resolve their image-asset UUIDs to the prod-public asset-plane URLs (refs only; blob bytes stay
  backlog DEF-M10-01). _Landed ext `2b8e9a0`: new REFERENCED-SUBSET admissibility kind (reverse-reference
  closure, firewall-admitted); `ReferencedFilesColumns()` (26-col, verified vs fresh bootstrap); + the
  replay-clear DELETE/TRUNCATE split for the external directus_settings FK. +13 tests, 12/12 packages green._
- [x] **§5 Referential closure + cross-surface fidelity gene** (ext, DEF-M21-04/NEW-3) — make the taxonomy
  capture referentially closed against the content it serves (full-taxonomy capture, the simple fallback the
  corpus already names) + a measured cross-surface closure gene (no content row references a taxonomy node-id
  the captured subset lacks) + close the 20 dangling relations. _Landed ext `4cb8786`: full-taxonomy capture
  already the state (org_id IS NULL); new `OpSnapshotCrossSurfaceClosure` gene measures content→taxonomy
  node-id closure (standard crit, non-blocking); surfaces the 1 genuine prod residual (K-AIFUNX-E658, a public
  sim → customer-only skill, uncloseable w/o firewall breach). The 20 dangling relations were subsumed by M21's
  26-collection structure capture (M21-D7); M23 owns the external content→taxonomy refs — that's this gene._
- [x] **§6 Docs — env truth + safety + closure gene** (rosetta) — `corpus/services/{cms,studio-desk,
  jobsimulation,next-web-app}.md` (env/dependency truth), `corpus/ops/safety.md` (retire the live-prod-read
  notes; token-strip stays as the write-disarm), `corpus/ops/snapshot-spec.md` (the cross-surface closure gene).
  _Landed rosetta `e364b80`: all 6 docs updated; resolves KB-1 (directus_files now wired+documented) + KB-2
  (cms.md M10-gap retired). cms cutover, studio-desk local token, jobsim-via-RPC, next-web-no-direct-Directus,
  safety retire-live-read, snapshot-spec directus_files-true + closure gene + the 1 named prod residual._

## Build log
_(append per build session)_
