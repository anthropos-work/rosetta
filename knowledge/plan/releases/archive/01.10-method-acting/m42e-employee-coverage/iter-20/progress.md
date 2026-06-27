**Type:** tik (production-fix; design-plan P6 replay-wiring)

# iter-20 — P6 replay-wiring: sim-embeddings into set-dress + verify the library

## Phase 0 — type selection
Plain **tik** under TOK-10. Iter dirs exist (no bootstrap); the last several iters (12-19) all closed-fixed (no
3-no-prog tok trigger). iter-18/19 were committed as code+docs without iter dirs (orchestrator-driven); this run
resumes the iter-dir convention at iter-20.

## Phase 0b — KB-fidelity
SKIP (plain tik; inherits the milestone's standing GREEN verdict — the touched surfaces (snapshot-spec.md,
coverage-protocol.md) were authored by this milestone's own iters).

## Phase 0d — pre-flight tooling (replay pipeline)
The wiring depends on the `stacksnap replay --surface sim-embeddings` pipeline. Built the CLI from the authoring
copy (`.agentspace/rosetta-extensions @ iter-19`); ran the replay against demo-3's `cms` schema from the captured
cache BEFORE committing the wiring — it loaded + reindexed cleanly (see Phase C). Pipeline confirmed working.

## Phase A — measure (baseline on demo-3, BEFORE the replays)
- `cms.similarities` = **0 total / 0 public-sim** (design-plan root #3 — the AI-sim library is empty).
- All 4 cms similarity tables + 4 directus library-category tables EXIST on demo-3 (schema provisioned).

## Phase C — fix + re-apply
**The wiring (rext):** added `sim-embeddings` to `dev-stack/dev-setdress.sh` `snapshot_step`'s replay loop
(`for s in taxonomy directus sim-embeddings`) — covers BOTH the demo `up-injected.sh` and the dev path (one shared
engine). The surface targets the stack's `cms` schema; the existing per-surface exit-code handling is
surface-agnostic + non-fatal (rc 4 cms-missing / rc 5 cache-miss → skip, seed floor still runs); the
`boot_directus_step` guard is keyed `[ "$s" = directus ]` so sim-embeddings doesn't trip it. Updated the
set-dress log line + the `up-injected.sh` comment header for honesty. `bash -n` + shellcheck clean.

**Re-apply to demo-3 (authoring-copy CLI vs the live offset DSN — the sanctioned harness-change apply path):**
- `stacksnap replay --surface sim-embeddings --stack demo-3` → **4 tables, 1490 rows loaded, reindexed
  [cms.similarities.small_embedding3]**.
- `stacksnap replay --surface directus --stack demo-3` (re-replay, now carries categories) → 14 tables, 11982 rows.

## Phase D — re-measure (probe-confirmed on demo-3, AFTER the replays)
**DB:** `cms.similarities` = **274 total / 274 public-sim** (+ 278 categories / 274 features / 664 skills children);
directus library tables = **17 categories / 7 macro / 310 sim-junctions / 21 path-junctions**. pgvector
nearest-neighbour smoke (cosine `<=>` over a real row's `small_embedding3`) ranks correctly (self 0.0000, then
0.2211, 0.2256…) — the `searchSimulations` path is functional.

**UI (authenticated probe as maya-thriving, demo-3):**
- `/library/ai-simulations` (was "0 simulations"): http 200, innerTextLen 5692, **real CATEGORIES with counts**
  (Technology & Engineering 55, Cybersecurity 6, AI/ML/GenAI 7, Coding/Software 33, Cloud/DevOps/IT 10…);
  GraphQL `libraryMacroCategories=[7]`, `libraryCategories=[17]`, **`searchSimulations=obj`** (was null/empty); no
  failures, no console errors.
- `/library/skill-paths` (was no categories): http 200, innerTextLen 17422, categories + **22 public skill paths**
  (`publicSkillPaths=[22]`); no failures, no console errors.

## Close — 2026-06-25

**Outcome:** P6 sim-embeddings replay wired into the set-dress loop (both lifecycles via the shared engine) +
verified on demo-3 — `/library/ai-simulations` now renders real sims + categories (`searchSimulations=obj`,
cms.similarities 0→274), `/library/skill-paths` renders categories + 22 paths. The wiring (not the surface) was
the gap; the surface code existed since iter-19.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (P6 of P0–P8; the authoritative gate is P8's fresh-demo-up acceptance + the P7 semantic harness,
both later — this iter does NOT claim gate-met)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** wiring rides the shared `dev-setdress.sh` engine (one edit covers demo+dev); non-fatal per the
existing exit-code handling; cms-schema target needs no new boot step (only directus does).
**Side-deliverables:** none.
**Routes carried forward:** P7 (semantic-coverage harness rebuild) → iter-21+; P8 (fresh-demo-up acceptance +
manager sweep) → a later run (the authoritative gate). The fresh-demo-up reproducibility of P6 is a CAPTURE-path
change (per `capture-path-milestone-live-acceptance` memory: needs re-capture + clear-old-cache + FRESH demo-up to
prove) — the orchestrator already re-captured; P8 proves it on a fresh bring-up.
**Lessons:** A snapshot SURFACE existing ≠ it being replayed — the set-dress loop's surface list is the wiring
seam; a new surface must be ADDED to `for s in …` or a fresh stack silently skips it (the AI-sim-library-empty
gap). Documented in snapshot-spec.md (the "replay leg is wired into the set-dress loop" note).
