# M208 — Spec notes

_Technical details accumulate here as the milestone is built._

## Pre-flight audits — Re-sync & merged-schema ground-truth
- Phase 0b KB-fidelity: **YELLOW** — report `kb-fidelity-audit.md` (sha at audit `e319d2f`). No blind
  area; pre-merge corpus staleness tracked as KB-1/2/3 (all Fate-2 → M210). Proceed.
- Topic → doc → code triples:
  - Merged shape → `corpus/services/{backend,skiller}.md` → `stack-dev/app/internal/data/ent/schema/{skill,jobrole,category,specialization,*_embeddings,skiller_mixins}.go` + `internal/rpc/skillerrpc/` (merge commit `1fc00c78`)
  - RPC re-point → `backend.md` + `dependency_map.md` → `platform@origin/main:docker-compose.yml` (`SKILLER_RPC_ADDR=http://backend:8083`)
  - 4-subgraph federation → `graphql-wundergraph.md` + `CLAUDE.md` → `platform@origin/main` router config
  - Re-sync ops → `update_guide.md`/`setup_guide.md`/`run_guide.md` → `stack-dev/platform/Makefile`

## Re-sync (make pull / refs)

Both stacks pulled to `origin/main` on 2026-07-08. `make pull` iterates the SIBLING repos from
`repos.yml` and does NOT touch `platform` itself, so platform was pulled directly first (which also
swaps in the merged `repos.yml`/compose before the sibling sweep runs).

**Before → After (short refs):**

| repo | stack-dev before → after | stack-demo before → after |
|---|---|---|
| **platform** | `5e1ae6b` → **`0808b92`** (2 ahead: rm skiller from compose+repos.yml+Make) | `5e1ae6b` → **`0808b92`** |
| **app** | `a848cccb` (v1.318) → **`c3c45e01`** (v1.334.1) — **86 commits** | `158a8394` → **`c3c45e01`** (v1.334.1) |
| cms | `57297a6` → `770ec3a` | `57297a6` → `770ec3a` |
| jobsimulation | `9f40604a` → `5d3003f9` | `9f40604a` → `5d3003f9` |
| graphql-wundergraph | `38f5d0a` → `c284453` (**`schemas/skiller.graphqls` deleted**; backend.graphqls +259) | `7ffe4f8` → `c284453` |
| next-web-app | `d689ecdea` → `23bdbb5db` | `928cc8e32` → `23bdbb5db` |
| studio-desk | `7a9ad78` → `f6320f8` | `7a9ad78` → `f6320f8` |
| sentinel / storage / messenger / roadrunner / skillpath | already current (unchanged) | already current (unchanged) |
| skiller (**vestigial**) | `b7a8950` (not in repos.yml — removed in §2) | `b7a8950` (removed in §2) |

Post-pull `platform` (both stacks): `repos.yml` has **0** skiller entries, `docker-compose.yml` has
**0** skiller services, and all four `SKILLER_RPC_ADDR` values are `http://backend:8083`.

Out of scope / untouched: `stack-dev/ant-academy` is **13 behind** but **not in `repos.yml`** (a
Clerk-free UI-tier native app, not part of the skiller merge) so `make pull` skips it — a UI-tier
concern for M211, not this merged-platform de-risk. `rosetta-extensions` (stack-demo) is a pinned-tag
clone, also not in `repos.yml` — M209's concern.

## Vestigial clone removal

`stack-dev/skiller` + `stack-demo/skiller` (both `b7a8950`, 8 MB each) — verified clean (no
uncommitted work) and referenced by **0** entries in either `repos.yml` — `rm -rf`'d. Post-removal
scan of `docker-compose.yml` confirms **no residual skiller container wiring** (no `context: ../skiller`,
no `http://skiller:8086`; all consumer RPC addrs point at `backend:8083`). A lingering clone was a
false signal only; nothing built or wired against it.

## Re-migrate against public
## Merge fact-sheet
## M25-D9 (opportunistic Fate-1)
