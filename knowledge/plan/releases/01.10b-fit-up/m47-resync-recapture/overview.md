---
milestone: M47
slug: resync-recapture
version: v1.10b "fit-up"
milestone_shape: section
status: done
created: 2026-06-29
last_updated: 2026-06-29
complexity: large
delivers: corpus/ops/snapshot-cold-start.md (the sanctioned MCP-DSN auto-capture path)
issues: Thread-A stale-clone drift (M201 false-negatives) + demo-up #2 (cold-start not turnkey)
---

# M47 — Re-sync & recapture

## Goal
Bring the `stack-demo` clone set **and** the captured snapshot **current with prod**, clone the absent rext
authoring copy, and re-validate the M201 false-negatives — so every downstream fix in this release is graded
against **current code**.

## FINDING (2026-06-29) — the clones were already current
The M201 close *reported* the clones ~5 weeks / 115+ commits behind prod (next-web @ v2.33.2). **M47 measured them
and found them CURRENT:** next-web @ **v2.89.0** (2 commits behind, fast-forwarded by `make pull`), every other repo
at `origin/main` (0 behind), and the **member-AI-readiness feature present** in `app` (v1.315). So the heavy "re-sync
5 weeks + rebuild" was a **trivial `make pull`** — the ⚠ risk did not materialize. The genuinely-stale surface is the
**rosetta corpus** (AI-readiness undocumented) — that is **M48**'s job, not M47's. The snapshot recapture confirmed
**both schema digests unchanged** (taxonomy `c75ce94…`, directus `ea2e187…`) → a clean in-place data refresh.

## Why section
The deliverables are enumerable up front: clone the authoring copy, sync N repos, rebuild + re-migrate, implement
the sslmode-normalizing MCP-DSN auto-capture, recapture the snapshot, re-grade the M201 negatives. The *risk* (what
breaks when 5 weeks of prod lands) is real, but the work list is known — flagged ⚠ in the risk section.

## Repo split
- **`rosetta-extensions`** (authoring copy → tag `fit-up-m47` → consume per-stack): the `stack-snapshot` capture
  fix (sslmode normalize + MCP-DSN-as-sanctioned-source + drop the cold-cache prompt) + the `up-injected.sh`
  set-dress wire-in.
- **`stack-demo/` platform clones** (operational, not committed to rosetta): `make pull` / checkout current prod
  refs across the repo set; rebuild images; re-migrate.
- **`rosetta`** (this corpus): the `snapshot-cold-start.md` update (the auto-capture path is now sanctioned +
  automated).

## Scope
- **In:**
  - **Clone the rext authoring copy** — `.agentspace/rosetta-extensions/` is **absent**; clone it (the prerequisite
    for ALL rext work this release, per the note-2 policy). Tag the first change `fit-up-m47`.
  - **Re-sync the platform clones to current prod** — bring the `stack-demo/` repo set current (current `main` /
    latest stable per-repo tag), rebuild Docker images, re-run migrations. Capture the before/after refs.
  - **Cold-start MCP-DSN auto-capture (demo-up #2)** — normalize `sslmode=no-verify → require` in
    `stack-snapshot/pg/pg.go:54` (`DSNForOffset`); teach `stack-snapshot/source/source.go` to accept the configured
    `postgres` MCP DSN (`marco_read`) as a **sanctioned primary-read** source under `AssertPublicOnly`; **remove the
    cold-cache prompt** (just do it); wire the auto-capture into `up-injected.sh` set-dress.
  - **Recapture the snapshot from current prod** — recapture taxonomy + Directus (public-only firewall, 0 tenant
    rows) into `.agentspace/snapshots/`; bump the capture version. (Coordinate the version bump with M52 — the M45
    batch cache is keyed on the taxonomy capture version.)
  - **Re-validate the M201 negatives** — re-grade the M201 verify verdicts (esp. **member-AI-readiness**) against
    current code; record which were stale false-negatives (feeds M48 docs + M51 seeder).
- **Out:** bring-up ordering/secret/frontend fixes (M49); corpus architecture/service re-ground (M48); content
  seeding (M50/M51); the manifest consolidation (M52).

## Depends on
None — the **foundation**. **Parallel with:** none (everything else builds on the current-code stack it produces).

## Open questions (resolve during build)
- Which prod ref to pin the clones at — latest `main` vs latest per-repo release tag. *Lean:* latest stable
  per-repo tag where one exists, else `main`.
- The recapture changes the snapshot digest → cache invalidation. *Lean:* accept the bump; coordinate with M52 so
  the batch cache re-keys cleanly (cache-spec invalidates on taxonomy capture-version change).

## KB dependencies (read as contract)
- `corpus/ops/snapshot-spec.md` + `corpus/ops/snapshot-cold-start.md` (capture/replay + the cold-start runbook)
- `corpus/ops/safety.md` (the `AssertPublicOnly` firewall + capture-source policy — the MCP-DSN must satisfy it)
- `corpus/ops/secrets-spec.md` (the secret source the recapture/bring-up reads)

## Delivers
- **→ rosetta-extensions:** the `stack-snapshot` MCP-DSN auto-capture (sslmode-normalized, sanctioned, no-prompt),
  tagged `fit-up-m47`.
- **→ rosetta:** `corpus/ops/snapshot-cold-start.md` — the wired `postgres` MCP DSN is now a **sanctioned**
  cold-start primary-read source (was explicitly NOT, per M20-D4); the auto-capture is part of the bring-up flow.

## Risk — RETIRED
The flagged ⚠ "biggest unknown" (pulling 5 weeks × the repo set → migrations/build breaks; a digest-changing
recapture) **did not materialize**: the clones were already current (trivial pull) and both snapshot digests were
unchanged (clean in-place refresh). See the FINDING above. The release's real risk now lives in the *other*
milestones (M50 content root-causes, M51's freshly-mapped AI-readiness data model), not here.
