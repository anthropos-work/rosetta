---
iteration_type: tik
iter_shape: production-fix
status: closed-fixed
created: 2026-06-25
---

# iter-08 — land iter-07's three routed fixes; re-sweep to the post-fix residual

**Active strategy reference:** TOK-01 (`sweep-then-route-by-leverage`). This tik lands the three
Fate-3 items iter-07 routed forward, then re-sweeps the frontier-exhausted reachable set to quote
the post-fix residual `(failing, escapes)`.

**Cluster / target identified:** iter-07's close routed exactly three well-scoped fixes — the
TRUE residual was `(failing=3, escapes=3)` over the frontier-EXHAUSTED reachable set (87 pages):
- `SNAP-M42e-stale-taxonomy-recapture` — re-capture the public taxonomy (picks up the 22 missing
  public skills incl `K-AIFUNX-E658`) + cache swap + re-replay into demo-3 → the 2 empty skill-paths
  render. The highest-leverage production fix (expected failing 3 → 1).
- `SCOPE-M42e-sim-start-runtime` — a `skipPaths` rule for `/sim/.../start` (a runtime simulation-launch
  surface, crawl-scope class like the iter-04/05 `/result/` pages) → failing 1 → 0.
- `ESCAPE-M42e-skillpath-external-articles` — the editorial-citation allow-rule + presenter-notes list
  (full set known = 3) → escapes 3 → 0.

**Hypothesis:** with all three landed, the frontier-exhausted re-sweep reports `(failing=0, escapes=0)`
— the gate is MET for the employee vantage.

**Expected lift:** `(failing=3, escapes=3)` → `(0, 0)`.

**Phase plan (protocol Phase A–E):**
- The fixes are already APPLIED to the live demo-3 (taxonomy re-captured + replayed = 42790 skills incl
  `K-AIFUNX-E658`) and to the harness (the crawl.ts + coverage.spec.ts sim-start-skip + citation
  allow-rule). This tik's job is the AUTHORITATIVE re-measure + close.
- **Phase A (fast-confirm first):** re-probe just the previously-broken pages — the 2 empty skill-paths,
  the empty sim-start, and the 3 citation-bearing chapter pages — to confirm the taxonomy fix made the
  skill-paths render and the harness rules behave. Cheap, streaming.
- **Phase D (authoritative re-sweep):** run the full coverage sweep at cap=150 (frontier exhausts ~87),
  STREAMING per-page output + journal heartbeats. Quote the post-fix `(failing, escapes)`.
- **Phase E:** grade + close; commit the rext code-half (crawl.ts + coverage.spec.ts) + tag
  `method-acting-m42e-iter08`; commit the corpus doc-half.

**Escalation conditions:** if the re-sweep shows a NEW failure mode the three fixes didn't address →
triage + route forward (continue, not exit). If a 100%-blocker can ONLY be closed by a platform edit →
re-scope-trigger (the zero-edit line).

**Acceptable close-no-lift outcomes:** none expected — the fixes are landed; this is a re-measure tik.
If the re-sweep is unexpectedly flat with the fixes verified-applied, that is a flake/diagnosis gap to
investigate, not a close-no-lift.
