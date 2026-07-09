---
milestone: M208
slug: resync-groundtruth
version: v2.1 "quick change"
milestone_shape: section
status: archived
created: 2026-07-08
last_updated: 2026-07-08
complexity: medium
depends_on: none
delivers: the merge fact-sheet (corpus/services/backend.md + corpus/services/skiller.md stub) — the contract M209/M210/M211 grade against
issues: mixed pre/post-merge stack state (app post-merge @ v1.318 but platform 2 commits behind, still composes skiller); vestigial stack-*/skiller/ clones
---

# M208 — Re-sync & merged-schema ground-truth

## Goal
Bring both stacks (and the snapshot's target reality) current with the **merged platform**, and pin the
authoritative **merge fact-sheet** — so every downstream fix in this release grades against **current merged code**.
(Mirrors v1.10b M47's role: the load-bearing foundation.)

## Why section
The deliverables are enumerable up front: pull the stacks to the merged platform, drop the vestigial clones,
re-migrate against `public`, confirm the 4-subgraph compose, pin the fact-sheet. The *risk* (what the 86-commit
`app` pull + migration re-run surfaces) is real — flagged ⚠ in the risk section — but the work list is known.

## Ground truth (verified at design, 2026-07-08)
- **app** clone is already POST-merge (`a848cccb` / v1.318.0; ent schema has `skill.go`, `skill_embeddings.go`,
  `jobrole.go`, `category.go`, `specialization.go`, `skiller_mixins.go`; merge commit `1fc00c78 Deprecate skiller
  schema`) but **86 commits behind** `origin/main` (`c3c45e01` / v1.334.1).
- **platform** clone (both stacks) is **2 commits behind** `origin/main` (`0808b92`): `origin/main` carries
  `21429b7 remove skiller service from docker-compose and repos.yml` + `0808b92 remove skiller profile from Make
  Targets` — repos.yml has **no skiller** (11 repos), compose has **0** skiller services. So the platform-side of
  the merge **is landed upstream**; the clones are merely stale, not blocked.
- `stack-dev/skiller/` + `stack-demo/skiller/` clone dirs still exist on disk — **vestigial** (not in `repos.yml`).

## Repo split
- **`stack-dev/` + `stack-demo/` platform clones** (operational, not committed to rosetta): `make pull` to current
  prod refs; remove the vestigial `skiller/` dirs; rebuild images; re-run migrations against the merged `public`
  schema.
- **`rosetta`** (this corpus): the **merge fact-sheet** — a concise, authoritative statement of the merged shape
  that M209/M210/M211 grade against (anchored in `backend.md` + the `skiller.md` stub; the colleague's docs branch
  already drafts the stub).

## Scope
- **In:**
  - **Re-sync both stacks** — `make pull` `stack-dev` + `stack-demo` `platform` to `origin/main` (skiller gone
    from compose/repos.yml); pull `app` to current (v1.334, post-merge domain) + the sibling repo set. Capture
    before/after refs per repo.
  - **Remove the vestigial clones** — `stack-dev/skiller/` + `stack-demo/skiller/` (no longer in `repos.yml`; a
    lingering clone is harmless but a false signal).
  - **Rebuild + re-migrate against `public`** — rebuild Docker images; re-run migrations against the merged
    schema; confirm the **4-subgraph** compose (`backend`, jobsimulation, cms, skillpath), **no skiller
    container**, `SKILLER_RPC_ADDR=http://backend:8083`, `app` search_path no longer includes `skiller`.
  - **Pin the merge fact-sheet** — the moved tables now in `public` (names unchanged), the confirmed
    `organization_id IS NULL` public predicate, the ~42,763 public-skill count assertion, the re-pointed RPC, the
    4-subgraph list. The contract for M209/M210/M211.
  - **Opportunistic M25-D9 (Fate-1)** — the dev-`N` taxonomy replay `rc=4` migrate-ordering nuance lives on this
    re-migrate path; resolve it if it falls out (non-blocking; do not scope-creep if it doesn't).
- **Out:** rext code changes (M209); corpus body re-point (M210); live `/dev-up` + `/demo-up` bring-up acceptance
  (M211).

## Open questions / risks
- ⚠ **The 86-commit `app` pull + migration re-run** may surface a schema/migration issue (the fit-up M47 risk
  class). Bounded; capture before/after refs; do not proceed to M209 until the re-migrated stack composes clean.
- Confirm no *other* repo silently depended on the skiller container at compose-wire level (RPC addrs are the
  known one — re-pointed to `backend` upstream already).

## Done-bar
- Both stacks at current merged refs; no `skiller/` clone; no skiller container; 4-subgraph compose green; the
  merge fact-sheet written + cited from `backend.md`/`skiller.md`.
