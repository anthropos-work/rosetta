**Type:** tik — clause 3 (the migration-status map + its both-ways fence), under TOK-01 step 4.

## What was measured, before anything was written

Re-ran all six of `platform-alignment.md` §4's detection signals against origin HEAD, on this box:

| signal | reading, 2026-08-01 |
|---|---|
| 1 repo set | **9** repos in `repos.yml` @ `2adcf71` — `graphql-wundergraph` is gone |
| 2 migrating | **`app` alone** (`repos.yml:12`) |
| 3 declared schemas | **`app -> public` alone** (`repos.yml:13`) |
| 4 subgraphs | **1** — `supergraph-config-prod.yaml` lists `backend` alone; `schemas/` holds `backend.graphqls` alone |
| 5 compose services | 12 top-level + 2 from the **included** `common.yml`; cms/jobsimulation/roadrunner **still in the default `graphql` profile** |
| 6 org census | **93** repos — independently reproducing iter-01's count on a different box, with a different tool |

Every peer clone verified `behind=0` against its own origin first, so no row is cited to a stale checkout.

**`gh` is not installed on this box.** The census ran against the REST API with `GH_PAT` from
`platform/.env`, values-blind. That is now written into signal 6 rather than left as tribal knowledge.

## What landed

**`corpus/architecture/platform-migration-status.md`** — 32 service rows + a 19-row net-new census.
The doc `platform-alignment.md` §6 had linked to since iter-01 and which did not exist, so that link was dead.

**`rosetta-extensions/stack-core/platform_alignment_guard.py`** + `tests/test_platform_alignment_guard.py` —
layer 1 of §8's three fences, and the last of the three still unbuilt.

## The fence, watched going RED — in both directions, live

Not just unit-tested. Against **copies of the real map and the real `repos.yml`**:

| run | result |
|---|---|
| control, unmutated real pair | **GREEN**, rc=0 |
| direction **B** — delete `cms` from the `repos.yml` copy (the real class: a departure) | **RED**, rc=1, `[B departure] the map claims cms is in repos.yml, and it is not` |
| direction **A** — add `ant-observability` to the `repos.yml` copy (an arrival) | **RED**, rc=1 — **and it fired `[E census overlap]` too**, because that repo is in §3's census, which is a live cross-check of assertion E that was not designed for |
| control again, after the battery | **GREEN**, rc=0 |

The unit battery is 17 tests: a mandatory GREEN no-op control (`platform-alignment.md` §8 rule 5), one RED per
assertion A–E, a *"deleting the row must not satisfy the check"* mutant, and four parse/usage mutants that must
return **2** rather than 0. `test_real_map_is_green` runs the shipped map against the shipped `repos.yml` on
every suite run — so the fence has a driver, not just a definition.

`corpus_index_guard` was likewise **watched going RED** on the unindexed map (it named the exact file), then
GREEN after the `corpus/architecture/README.md` row — the iter-01 precedent, repeated.

## Findings the map itself produced

1. **Five services nobody on our side has ever named.** `git log -p --follow -- docker-compose.yml` returns 26
   service names ever; the corpus knows about 21. `nats` (removed `8770fe6`, four days after the first
   commit), `web-app`, `chromedp`, `simulator`, `realtime`. **`simulator` was replaced by `jobsimulations` at
   `84862d1` (2024-05-29)** — the first ancestor of what is now `app/internal/jobsimulation/`, i.e. the
   consolidation program has an earlier chapter than iter-01's "v2.0 skiller" origin story.
2. **The `archived` flag is the cheapest fold-confirmation there is.** jobsimulation + skillpath archived
   2026-07-31, graphql-wundergraph 2026-07-30, skiller 2026-07-01 — days after each fold. **And `chronos` is
   NOT archived**, while the corpus says it is.
3. **`roadrunner` is the one row where prod and the platform's own declaration contradict each other**
   (`repos.yml` "folded" vs `terraform:19` `= 1`). Recorded as a contradiction, not resolved (D-M257x-20-2).
4. **The router is `live-standalone` in prod and `decommissioned` locally** — the single sharpest illustration
   of why the protocol demands two states per row. One state would have been wrong for half its readers.
5. **19 unarchived, recently-pushed org repos the corpus has never named**, out of 46 undocumented of 93. Five
   are LiveKit agents; `ai_architecture.md` documents the LiveKit *engine* and none of the agents.

## Close — 2026-08-01

**Outcome:** **Gate clause 3 MET — 2 of 5 → 3 of 5.** The migration-status map is checked in, its row set
derived from git history rather than memory (which found 5 services the corpus never knew), every row cited,
and it is machine-fenced against `repos.yml` in both directions with the fence watched going RED in each,
live, with a GREEN control before and after.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (3 of 5 — clauses 1, 3, 4 met; 2 and 5 outstanding)
**Phase 5 grading:** (1) gate-met: n (3 of 5) — (2) triggered-tok: n (this tik moved the metric; the
iter-19 no-prog streak is reset) — (3) re-scope: n (platform origin HEAD `2adcf71` re-checked at open AND
close, unchanged — occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) —
(6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-20-1 … D-M257x-20-5.
**Side-deliverables:** two `platform-alignment.md` amendments made in the same commit per the protocol-evolution
rule — §4 signal 6 now names the `archived` flag and the `curl`/`GH_PAT` form for a box without `gh`, and §8's
layer-1 row records the fence as landed rather than as a precedent to follow.
**Routes carried forward:**
- **Clause 5 is now mechanical** — the map is the source of truth the corpus sweep reconciles against.
  `DOC-M257x-iter14-corpus-router-drop` (35 files / ~128 router hits) plus four map-derived corrections:
  `chronos` is not archived, `AI-Labs` is a live repo not just a subsystem, the five LiveKit agent repos are
  undocumented, and `studio-room` is not a repo name. → next tik.
- `FIX-M257x-iter15-library-category-expansion` (two fields: 119 `library_category` + 11 `job_position`),
  `FIX-M257x-iter15-directus-versions-403`, `FIX-M257x-iter19-playthrough-runner-path` → clause 2, unchanged.
- **New:** `DOC-M257x-iter20-net-new-repos` — the 19-row census names the gap; deciding which of those repos
  deserve a corpus doc is a milestone-sized question, not this iter's. The map records them so the question
  can be asked at all.

**Lessons:**
- **A completeness claim should name the command that generates it.** "Every service the platform has ever
  had" written from memory would have had 24 rows and looked finished. Two `git log -p --follow` invocations
  added 8, one of which (`simulator` → `jobsimulations`, 2024) predates the consolidation story we had been
  telling. The map now prints the commands so its own completeness is auditable rather than asserted.
- **A negative control can prove more than it was designed to.** Direction A's mutant used a *real* census
  repo name, so it fired assertion E as well — E was proven live by accident, and that is worth doing on
  purpose next time: mutate with real values, not placeholders.
- **The fence's exit codes carry the milestone's own lesson.** Refusing to default `PLATFORM_REPOS_YML`, and
  raising on an empty table, are both the same rule: this milestone exists because a stale local file was
  read as ground truth, and because a check that cannot run must never report OK.
