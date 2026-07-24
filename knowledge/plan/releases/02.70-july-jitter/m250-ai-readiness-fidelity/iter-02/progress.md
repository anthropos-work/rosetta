**Type:** tik — Lane A, the arithmetic-spine atomic edit (under TOK-01).

# iter-02 — progress

Landed the 8→31 arithmetic spine as one atomic, compile-and-fence-green change in rext `stack-seeding`.

## What landed
- **`ai_readiness_config.go`** — new shared `aiReadinessDefaultSkills` (31 verbatim from platform `defaults.go`,
  19 core @1.0 + 12 enabling @0.5) + `aiReadinessDefaultSims` (3 pinned track-keyed named uuids) +
  `aiReadinessDefaultSkillPool()`. The config seeder now writes the 31 defaults unconditionally (mirrors
  `provision.go`; closure-safe) and 3 sims with the `track` column, row-id keyed by (org, step_type, track).
  Constants `5/3 → 19/12`.
- **`ai_readiness_funnel.go`** — `aiSkills` pool built from `aiReadinessDefaultSkillPool()` (31, deterministic,
  single-sourced with config); Step-2/3 sim refs pinned to the defaults (Tech sim for Step-2, interview for
  Step-3); `aiReadinessStartedHeroSkills 3 → 9` (→ 11/30, re-derived at 25.0).
- **Fences re-derived (5 test blocks across 3 files):** `ai_readiness_m219_test.go` (6.5→25.0, loop 13→50,
  "14"→"51", MirrorsComputeTier1 cases, Aria/Ben heroes, and the double-round divergence sub-test → a **live
  invariant** proving the divergence class is UNREACHABLE at denom 25.0 [held/25×100 = held×4 is always
  integral for half-step weights]); `ai_readiness_config_test.go` (`isAISkillNodePresent`→the 31; sims 2→3 +
  distinct tracks + the pinned refs; the co-derivation invariant re-pointed at the pins; the no-fabrication
  test re-derived to "writes the 31 real defaults"); `ai_readiness_funnel_test.go` (no-fabrication re-derived
  to "every evidence node-id ∈ the 31 defaults"); `ai_readiness_harden_test.go` (NoTaxonomy comment).

## Measurement
- `go build ./...` GREEN; `go test ./seeders/` GREEN; full `go test ./...` (stack-seeding module) GREEN incl.
  `dna` (closure) + `cmd/stackseed`. No collateral breakage in rext (cockpit/test_cockpit `6.5.2` = FontAwesome
  CDN, not arithmetic; stackseed truncate-list already carries the tables).
- **Gate at LIVE render:** still 0/5 (no reset-to-seed render this iter — deferred to iter-05 per TOK-01). The
  arithmetic/config foundations for parts 1, 2, 5 are landed + Go-test-validated (the fences ARE gate part 5's
  frozen-vs-live arithmetic check). Honest read: real progress, not a claimable live gate-part pass yet.

## Close — 2026-07-24

**Outcome:** the arithmetic-spine atomic edit landed — demo now seeds the platform's 31 real defaults + 3
track-keyed named sims; all fences re-derived green at 25.0/31. Unblocks Lanes B + join.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (live render pending — iter-05)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue (→ iter-03, Lane B)
**Decisions:** D4 (write the 31 defaults unconditionally, mirroring provision.go — closure-safe), D5 (started-hero 3→9 → 11/30, keep the >=10 believable floor), D6 (double-round divergence is arithmetically unreachable at denom 25.0 → converted the stale hardcoded triple into a live invariant), D7 (defer participants_filter track-tagging + business-sim session routing to the render loop).
**Routes carried forward:** iter-03 = Lane B Directus set-dress; iter-04 = evidence-distribution join; iter-05 = live reset-to-seed render; participants_filter + business-sim session routing → render tik.
**Lessons:** three MORE fences beyond m219/harden encoded the OLD "no-fabrication-by-taxonomy-starvation" premise (config `NoTaxonomyWritesNoFabricatedSkills`, funnel `NoTaxonomyNoFabrication`, the SimRefs co-derivation). M250 replaces it with "no-fabrication-by-construction" (the 31 are the platform's own real defaults). When a design moves from pool-derived to constant-list, EVERY starvation-premised fence inverts — re-derive them as "writes the real defaults" not "writes nothing".
