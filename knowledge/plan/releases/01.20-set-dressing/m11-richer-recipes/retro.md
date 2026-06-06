# M11 — Retro (Richer-world recipes, presets + corpus polish) — the LAST v1.2 milestone

## Summary
M11 is the M8-analog product/discoverability layer that closes v1.2 "set dressing": it curates the
100%-coverage full-fidelity world (shipped by M9a/M9b/M10) into something a demo curator can actually
use and discover — with **zero new production code path**. Five sections: (1) the 3 seed presets
gained a documented **FULL-FIDELITY PREREQUISITE** comment-header (replay `taxonomy` + `directus`
before seeding for a set-dressed world; graceful structural-only degradation without) while staying
**purely structural** — snapshots are stack-global reference data, not org-scoped, so no schema field
was added (M11-D3); (2) the **set-dressed `corpus/ops/demo/` recipe family** — both org-onboarding +
skill-progression recipes gained a `/demo-snapshot` replay step and call out the real catalog/templates,
the FALSE "waived/future-v1.2" note in `recipe-skill-progression.md` was rewritten to the shipped
100%-coverage reality, and a **new `recipe-snapshot-world.md`** documents the capture→replay→set-dressed
flow end-to-end; (3) the **`/demo-snapshot` skill** (M11-D1, resolving M11-Q2 — a NEW skill rather than
a `/demo-seed` extension, because capture is a privileged prod READ with its own firewall contract,
distinct from seeding's per-stack WRITEs; `replay` is the headline verb, `capture` the rare maintenance
op); (4) corpus cross-links + the CLAUDE.md skill-table row + data-DNA-100% throughout; (5) the
release-close hygiene carry — the stale `stacksnap --help` text fixed (M11-D4: framework tag
M9a/M9b → M9a/M9b/M10 + the registered-but-unlisted `directus` surface now listed). Build (5 sections)
→ 1 harden pass → close clean (single attempt, 0 findings).

**With M11 closed, all 4 v1.2 milestones are done + merged — v1.2 is complete.** The next step is
`/developer-kit:close-release` (release-level review, then merge `release/01.20-set-dressing` → `main`
+ tag `v1.2`).

## Incidents This Cycle
- **None.** No production-code path was introduced, so there were no runtime defects to surface. The
  single harden pass pinned two real Fate-1 *test* gaps (the unpinned `--help` contract + the
  never-loaded shipped presets, plus a docs↔parser flag-drift guard) — no bug, just missing
  regression coverage. Flake gate 0 (5/5 both M11-touched packages); both modules `-race` + gofmt +
  `go vet` clean. No P2s, no regressions. Close found 0 findings.

## What Went Well
- **The `/demo-snapshot`-as-separate-skill call (M11-D1)** kept the UX honest: `replay|capture|status`
  maps 1:1 to the `stacksnap` subcommands, capture's prod-read blast radius stays distinct from
  `/demo-seed`'s per-stack writes, and the curator flow chains cleanly (`up → snapshot replay → seed →
  login`). The spec already modeled them as sibling extensions, so the skill split followed the grain.
- **Presets stayed structural (M11-D3).** Resisting the temptation to encode a snapshot field in
  `stack.seed.yaml` avoided a false coupling — a preset describes an *org*, not the platform's
  reference library. The prerequisite is a documented header + the README flow, and the seeder
  degrades gracefully (empty catalog, free refs) when no snapshot is replayed.
- **The harden tests pin real cross-repo contracts.** The docs↔parser drift guard
  (`TestDocsFlagsExistInParser`, `TestDroppedDumpFlagStaysGone`, `TestDocumentedSourceKindsAreReal`)
  and the registry-driven `--help` pins (`TestHelp_NamesEveryRegisteredSurface`) catch the exact
  drift class M11 §5 itself fixed — a future registered-but-unlisted surface, or a doc promising a
  flag the parser lacks, fails the suite. Each was mutation-verified to bite.

## What Didn't
- **The `--help` text was a documented contract that nothing pinned** — the M10 `directus` surface
  shipped registered but unlisted in `--help`, and the drift sat undetected until the M11 KB-fidelity
  audit caught it. Root cause: no test asserted the help string against the surface registry. Fixed
  at the root (registry-driven `TestHelp_*`, not a string literal) so the *class* can't recur, not
  just the one instance. Minor, doc-side (the surface always worked; the help under-advertised it).
- **The shipped preset files were never loaded by the suite** — verified only by hand via `stackseed
  --dry-run` during build. A stray field would have surfaced only at a curator's runtime past the
  strict `KnownFields(true)` loader. Now pinned by `TestShippedPresets_ParseStrictAndValidate`.

## Carried Forward
- **DEF-M10-01 (S3 media blob BYTES) → v1.3** — inherited unchanged; Fate-2, confirmed-covered by the
  `roadmap-vision.md` cloud-store / S3-backend seed. M11 touched no media surface; its ref-floor-vs-bytes
  boundary wording honors the routing. Re-confirmed GREEN at the M11 close deferral re-audit.
- **No M11-originated carry-forward.** Every M11 scope item landed Fate-1; the milestone added zero
  deferrals.

## Metrics Delta (from metrics.json)
- **Go test funcs:** 701 → **708** (+7) — stack-snapshot 207→212 (+5: the `--help` contract pins +
  the docs↔parser flag-drift guard); stack-seeding 230→232 (+2: shipped-preset strict-parse/validate +
  size-order). alignment 46, clerkenstein 218 unchanged.
- **Flake:** 0 (5/5 shuffled, both M11-touched packages). **Race:** both modules green. **gofmt +
  go vet:** clean. **Coverage:** data-DNA 100% (unchanged — M11 is curation, not a new surface).
- **Findings at close:** 0. **Harden passes:** 1 (expected short pass for a docs/discoverability
  milestone). **Deferral re-audit:** GREEN (0 new, 0 repeat, 0 aged-out).
- **Extensions tag:** `stack-snapshot-m11` @ `1e18df6` (annotated, pushed).
