---
iter: 06
milestone: M46
iteration_type: tik
status: closed-fixed
created: 2026-06-26
---

# M46 · iter-06 (tik #5) — REAL org-scale gate-proving (deliverable #5, part 1)

**Active strategy reference:** TOK-01. Tik #5 = deliverable #5, the real large-batch gate-proving sweep.

**Re-survey (Step 0):** all four code deliverables (auto-fill, per-story distribution, preview, the fence +
verified 429 backoff) landed in iters 02-05. The remaining gate face is the empirical "does a full org
rendered from one descriptor LOOK believable under the M42 semantic gate?" — answerable only by a real run.
Demo-3 is up (361 users baseline, taxonomy replayed); the harness is calibrated.

**Cluster / target identified:** fire a real ~500+ member supporting-population batch via Azure gpt-4o-mini,
seed demo-3, run the M42 semantic-coverage sweep on the generated org, prove the gate (believable populated
+ 0 collisions + closure GREEN + budget).

**Hypothesis:** the descriptor `gen-batch-org-fill.seed.yaml` (Cervato 500 + Solvantis 120, one `fill: true`
batch per story → ~614 generated members) fills both orgs believably; the M42 sweep passes on the generated
population.

**Phase plan:** A — build the ~500+ descriptor + preview (offline). B — fire the real capped Azure batch
(values-blind). C — seed demo-3 + run the M42 sweep. D — measure (collisions / closure / budget / coverage).
E — close.

**What actually happened (the iterative heart — the real run surfaced two org-scale bugs):**
The real 614-member batch SURFACED two scaling bugs invisible at M45's bounded N=20, both FIXED + tested:
1. **Multi-batch cache-index collision** — `genOne` cached at the batch-LOCAL `m.Index` (0..Count-1) while
   `Has`/`Get` use the GLOBAL slice index. With one batch they coincide; with two story batches, the
   Solvantis members (global 497-613, local 0-116) overwrote Cervato's files and the Solvantis slots stayed
   empty → the entire 2nd story's 117 members were LOST (497/614 cached, 0 Solvantis). Fixed (cache at the
   global index) + regression-tested (proven to fail on the old code).
2. **Name-distinctness at org scale** — gpt-4o-mini is a strong name-attractor: over 600 members it
   re-picked a handful of names hundreds of times (Zara Al-Mansoori ×20), so only 354/614 = 58% distinct.
   Fixed with a deterministic cost-free last-attempt disambiguator (keep first name + swap a distinct
   surname keyed on global idx) + a larger avoid-hint (40→120). Distinctness now guaranteed at any scale.
   2 tests.

**Empirical gate metrics PROVEN on the real run** (the faces the bounded M45 batch couldn't test at scale):
- **hero-collisions = 0** across 614 generated members (gate PASS).
- **valid-JSON rate 100%** (>> the 95% threshold).
- **$0 cache-hit reseed** — 579 cache-hits cost $0 (the reproducibility face).
- **--max-cost budget guard correctly aborted 3×** at its ceiling ($0.30/$0.35/etc.), never breaching — the
  mandatory-cap face proven on a real run. Total real spend across the capped runs ~$0.70.
- **614/614 members cached** end-to-end (one descriptor fills a ~600-member 2-org world).

**Escalation conditions:** the bugs were in-scope tooling fixes (not platform-bound), so no re-scope. The
full gate-proof TAIL — regenerate the now-distinct cache + seed demo-3 + the M42 manager+employee sweep +
datadna closure — needs more than this tik's budget; it is the next invocation's tik #6.

**Acceptable close-no-lift outcomes:** N/A — this tik LANDED two real fixes + proved 4 of the 5 gate faces
empirically (collisions/JSON/reseed/budget); the 5th face (the M42 semantic coverage on the seeded org)
needs the regenerate+seed+sweep tail.
