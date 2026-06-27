# M40 Retro — Per-stack Directus public-policy serve-grant

## Summary
M40 closed the highest single-surface gap in v1.10: on a fresh `/demo-up`, a hero's `/library/*` + `/profile/activities`
surfaces now render real content, anonymously via cms, entirely from the snapshot replay — **zero seeding, zero
platform-repo edits**. The work landed in rext `stack-snapshot/directus/structure.go` (@ tag `method-acting-m40` →
`5e53301`); the rosetta side carries the doc-half (`snapshot-spec.md`) + plan records. Built + hardened (3 passes) +
closed in one near-clean review: 2 findings, 0 blocking, both docs/triage.

## Incidents This Cycle
None. 0 production bugs across the build + 3 harden passes; 0 flakes (5/5 shuffled at close); 0 regressions
(supply-chain GREEN, go.mod/go.sum byte-identical). The 2 test-fixture bugs surfaced during harden Pass 2 were fixed
inline within that pass (whitespace-strict closure assertion + a missing junction in a fixture map) — both confirmed
the production closure logic is correct as specified, not product defects.

## What Went Well
- **The root cause was diagnosed bigger than the spec framing — and the fix collapsed the key risk fork.** The
  original a/b/c framing feared the `simulations.sequences` O2M might need a platform nil-guard (read-only — would
  have forced an escalation + split shipping). Live investigation (demo-3) found the real cause: the per-stack Directus
  had `directus_relations=0` / `directus_fields=0`, so the nested aliases were UNKNOWN to Directus, not "stripped under
  the public policy". Registering the relational metadata fixed BOTH halves in tooling — the fork was refuted, not
  navigated. (M40-D2, M40-D4.)
- **Dynamic capture over a hardcoded list.** The fix synthesizes the relational web from the sanctioned source under a
  both-endpoints closure (off-stack aliases dropped), matching structure.go's existing version-robust philosophy — no
  brittle enumerated collection list. (M40-D3.)
- **Clean firewall extension.** `directus_fields`/`directus_relations` carry zero tenant-scope columns (re-verified via
  the sanctioned prod structural read), so they slot into the existing `AssertStructuralMetadata` carve-out — the table
  set was extended, the predicate untouched. The adversarial review's tenant-column-bypass scenario is gated before any
  serve row is read.
- **Harden reached the 100%-statement ceiling and then deepened SQL-semantics behavior** (Go statement coverage can't
  measure SQL-template correctness — the both-endpoints AND-closure Go-mirror grid + the relational-special regex tests
  cover what coverage % can't).

## What Didn't
- Nothing material. The single nit: the build's `serveVersionsActions` slice was momentarily flagged at close as
  "unused at runtime" before being recognized as the readable drift-guard source the opaque SQL array-literal is pinned
  against — a 30-second resolution, kept as-is.

## Carried Forward
- **KPI "AI simulations completed" = 0** → **M42e + M42m** (Fate-2, M40-D7). Its source
  `public.local_jobsimulation_sessions` (=21 seeded) has no CMS dependency, so it's a separate frontend/auth-context
  surface genuinely owned by the per-vantage coverage milestones — their exit gate ("every reachable demo page renders
  non-empty content, 0 failing") already encompasses it; no plan edit. Deferral re-audit GREEN.

## Metrics Delta
(from `metrics.json`)
- Go test funcs: stack-snapshot **333 → 354** (+21; directus pkg). Release total 1248 → 1292 (incl M39 +23).
- Coverage: directus pkg **100.0% of statements** (`CaptureServeRows` 88.9% → 100% at harden).
- Flake: **0** (5/5 sequential shuffled at close).
- Supply-chain: **GREEN** (0 new deps; go.mod/go.sum byte-identical).
- Platform-repo edits: **0** (the key-risk fork refuted — both halves ship in tooling).
- Live acceptance (build-recorded, demo-3, anonymous cms): publicSkillPaths=22, publicJobSimulations=50,
  jobSimulation→sequences[].scenarioIntro — all >0.
