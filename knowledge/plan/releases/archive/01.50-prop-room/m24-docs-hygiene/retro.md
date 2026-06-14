# M24 — Retro

## Summary
M24 made the rosetta corpus tell the new truth after the M21–M23 cutover and absorbed the four aged-out hygiene
items the deferral audit had surfaced. The load-bearing correction: the corpus carried a **fictional local-Directus
docker service** (image `directus/directus:10.10.1` + admin/password creds + a `docker-compose` snippet) that
**never existed** in the platform compose — verified against `stack-dev/platform/docker-compose.yml` (no `directus`
service, only the `DIRECTUS_BASE_ADDR` prod pointer). M24 corrected it across `external_services.md` /
`service_taxonomy.md` / `quick_ops.md`, rewrote the known-state + `directus-local.md` + `snapshot-spec.md` on the
M23 cutover, and swept the print-only/exit-4/reads-live-from-prod framing across 5 skills + `CLAUDE.md` + 5 demo
docs (the real two-path posture: prod-read default; per-stack local Directus on `--local-content`). The hygiene
strand: a `toolchain go1.25.11` pin (clears the 12 called-stdlib advisories), a corpus README-index-row guard, the
alignment zero-critical-genes guard, and the `/project-stats` scope fix. 7 sections, all Fate-1.

## Incidents This Cycle
- **1 P2 bug fixed inline (harden Pass 1) — 0 regressions / 0 flakes.** The new `corpus_index_guard.py` shipped from
  build with a **prefix-collision false-negative**: its raw-substring reference test let an unindexed `setup.md`
  slip through whenever a different, indexed `dev-setup.md` was present (`setup.md` is a substring of
  `dev-setup.md`) — i.e. exactly the miss class the guard exists to catch. Caught + fixed in harden with a
  token-bounded regex match (ext `191d650`), with a regression test that fails on the pre-fix code. Live corpus
  still exits 0. Caught before close, by the milestone's own hardening — the system worked.
- **The close adversarial pass found 0 fixable findings.** 1 scenario probed (the guard's incidental-mention
  lenient-by-filename behaviour) was confirmed by-design and recorded; 0 code change.

## What Went Well
- **The README-index guard dog-fooded itself.** Running the new guard against the live corpus surfaced **7
  pre-existing index gaps** (5 architecture + 2 ops docs) that 3 straight releases had missed — all backfilled so
  the guard passes clean. The tool earned its keep on its first run, on the exact failure class it was built for.
- **The milestone's own harden caught its own bug.** The one real defect (the prefix-collision) was found by the
  hardening pass, not in the field — a behavioural gap that line-coverage was blind to (coverage was green while
  the correctness property went untested). Exactly the value a harden pass is supposed to add.
- **The stale-claim correction was verified against code, not assumed.** The overview flagged the risk ("verify
  against the *actual* platform compose, not assumed"); the KB-fidelity audit + the build both checked
  `docker-compose.yml` directly and confirmed no `directus` service — so the correction states a verified truth,
  not a guessed one.
- **The deferral audit shrank the backlog.** M24's charter was to CLEAR 4 aged-out hygiene items (NEW-5/6/11/14),
  and it landed all 4 Fate-1 while introducing 0 new deferrals — the ledger got cleaner, not longer.

## What Didn't
- **The "Repo split" mis-located one hygiene item.** The overview assumed §7 (`/project-stats` scope fix) was a
  `rosetta-extensions` item; investigation found the script is the shared **developer-kit plugin** `stats.sh` (no
  stats tooling exists in rosetta-extensions). Landed correctly at the real source (cross-repo `825cdce`) and
  recorded in M24-D3 — but the plan's repo-split guess was wrong, a reminder to locate a tool before assuming its
  home repo.
- **A pre-existing foreign-tool limitation surfaced mid-fix.** While fixing §7, `stats.sh`'s doc-counter was found
  to report 0 for rosetta (it only scans `knowledge/`; rosetta's docs live under `corpus/`). Out of M24 scope and
  not a v1.5 deliverable — recorded as an observation (M24-D3) so it isn't silently lost, not fixed here.

## Carried Forward
- **DEF-M21-02** (automated serve-live-integration harness) → **M25 field-bake** (unchanged — M24 is docs/hygiene,
  touches no serve-row render).
- **DEF-M21-01** (replayCmd conn-seam) → tracked tooling-debt follow-up (untouched by M24 — its ext diff is
  alignment + stack-core + go.mod pins, not `replay.go`).
- **DEF-M10-01** (S3 blob bytes + cloud store) → backlog (unscheduled); M24 touches no asset-plane code.
- **M24-OBS** (the `stats.sh` `corpus/`-blind doc-counter) → a foreign developer-kit-plugin limitation, recorded
  (M24-D3), not an M24 deferral (§7 landed in full).

## Metrics Delta
- **Go test funcs:** 844 → **850** (M24-own **+6**, all in `alignment`: the zero-critical-genes guard — +3
  `dna.Validate`, +3 `compare.GateMet`/`CriticalGenes`). The other 3 modules untouched. Coverage: `dna`
  98.5→**100%**; the touched `compare` code fully exercised.
- **Python (touched suite):** stack-core 77 → **85** (+8, the new `corpus_index_guard.py` lint — 8 build + 8
  harden). The other 4 suites untouched. Live corpus guard **exit 0**.
- **Flake:** 0 (5/5 — Go shuffled, Python sequential). **go vet / gofmt / py_compile:** CLEAN.
- **Supply chain:** the 12 stdlib advisories @go1.25.3 now **pinned out** by the `toolchain go1.25.11` directive.
- **Review:** 5 findings (0 scope / 0 must-fix code-quality / 0 docs / 0 tests / 1 adversarial-record / 3
  decision-triage). **Deferral audit:** GREEN (0 new deferrals, 4 chartered items cleared, 0 repeat / 0 aged / 0
  blockers).
- Full machine-readable record: [`metrics.json`](metrics.json).
