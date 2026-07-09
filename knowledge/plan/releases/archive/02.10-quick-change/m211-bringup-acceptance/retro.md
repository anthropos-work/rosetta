# M211 Retro — Bring-up acceptance: dev-up + demo-up green on the merged platform

**Closed:** 2026-07-08 · `closed-on-gate` · `iterative` · complexity large. **The FINAL v2.1 milestone.**
**Outcome:** proved the merged (skiller-in-app) 4-subgraph platform stands up **end-to-end, cold, on BOTH stacks**
via the re-grounded tooling, with **zero platform-repo edits**. **Gate 6/6 MET.**

## Summary
17 iters (1 bootstrap tok + 16 tiks) + 4 stabilized final harden passes drove the composite exit gate to 6/6.
The rosetta close diff is docs+plan only (3 corpus/skill docs + plan records, 0 code) — the tooling code lives in
the rext repo, code-of-record at tag `quick-change-m211` = `2039103`. The substantive work, all in the rext
tooling: the **cache-migration** (real 42,790-row taxonomy + 274 sims re-keyed `skiller.*→public.*`, replayed —
the no-prod-access recapture); the **root-cause fix** for stale rebuilds (see Incidents); the dev casbin
`init_policy.sql` load; the frontend offset-reuse guard; the demo-local `ACADEMY_URL` bake + academy-aware
cross-port hook; the demopatch URL re-pin (next-web v2.106.1); the Playthroughs reset-to-seed **roster-refresh**;
and the new **`dev-stack/migrate-dev.sh`** (dev cold DB-init, mirror of `migrate-demo.sh`). Result: cold
`/demo-up` GREEN end-to-end; M42 coverage GREEN both vantages; v2.0 Playthroughs 10/11 GREEN; dev cold DB-init
cold-verified; 0 residual skiller.

## Incidents this cycle
- **P1 (investigation, fixed) — stale injected build-scratch.** A truly-cold `/demo-up` stood up but federation
  failed (`_entities(Skill.name)` 422 → empty library/profile). Root cause: the injected build-scratch
  (`$STACK/clones/app`) was pinned at pre-merge `v1.315.0` (0 `Skill.name` refs), never re-synced, and **survived
  `down --purge`** → every demo-up baked a stale pre-merge binary. Fixed (rext `0593cff`): re-sync the scratch to
  the source's current release tag every bring-up (+ a fence). Verified scratch `v1.315→v1.334`, `Skill.name` 0→5.
  Two iters (09 diagnosis, 10 fix) — the highest-value find of the milestone; a false-green in iter-08 (container
  count MET but composed supergraph not checked) was corrected by the diagnosis.
- **P2 (fixed) — reset-to-seed roster gap.** `run-playthroughs.sh --reset` swapped the DB but not the
  fake-FAPI/BAPI roster (baked at bring-up) → hero login `400 unknown_identity` on a demo brought up for something
  else. Fixed (rext `e822c70`): `--reset` re-exports the roster from the seed + restarts the fake services + waits
  for FAPI. Documented in `playthroughs.md`.
- **No flakes.** The final harden flake gate ran 3× clean on the only runtime-nondeterministic new test
  (`TestMigrateDevLive`, throwaway containers); the rest are deterministic.
- **CAVEAT-2 (pre-existing, not M211):** 33 `test_dev_stack.py` CLI failures from an incomplete local
  `.agentspace/secrets` source tripping the secret-coverage pre-flight — an environment condition, unrelated to
  M211 code (the standalone `migrate-dev.sh` tests are green).

## What went well
- **Warm-first, cold-prove (TOK-01) held the whole run** — no triggered tok, no re-scope. The strategy shortened
  fix→re-measure and dodged the docker-reap hazard, then the cold run confirmed the gate once the inner loop was
  green.
- **Symmetry over invention.** The dev cold DB-init (`migrate-dev.sh`) is a mirror of the demo's already-proven
  `migrate-demo.sh`, not a new mechanism — the un-editable-platform-Makefile constraint made the fix symmetric.
- **Cache-migration = no prod access needed.** The merge being a pure schema-prefix move meant the real captured
  taxonomy was faithful once re-keyed (empirical column-match gated; never fabricated).

## What didn't
- **iter-08's "cold" was image-warm** — the container-count check for sub-condition (a) passed while the composed
  supergraph was stale (a dead skiller subgraph + a pre-merge app image). Lesson: a compose gate must assert the
  composed supergraph, not just container presence. Cost two extra iters (09/10) to diagnose+fix — but that's
  exactly what the iterative loop is for.
- **A literal full destructive `/dev-up` couldn't run on this box** (committed to the user's native content-line
  dev). The dev-specific gate delta (M25-D9 DB-init) was cold-verified on a faithful throwaway + live harness — an
  environment-respecting interpretation, with a clean-box full `/dev-up` recorded as belt-and-suspenders backlog.

## Carried forward (three-fate — none escape-hatch)
- **TEST-1** (rext `stack-seeding/README` test-count drift) + **DOC-1** (rext `dev-stack/README` should index
  `migrate-dev.sh`) → **Fate-2 → `/developer-kit:close-release` rext roll** (rext frozen at the re-tag HEAD;
  README reconciliations land at the code-of-record roll).
- **CAVEAT-1** (clean-box literal full destructive `/dev-up`) → belt-and-suspenders backlog (`roadmap-vision.md`).
- **DEF-M208-02** (`INVITATION_HMAC_SECRET` dev `.env`) → confirmed-covered by `/stack-secrets`.
- **PT-TODO** (assign-WRITE Playthrough half) → Fate-2, inherited from v2.0 (reserved manager-write tier).
- **DEF-M208-01 / M25-D9 RESOLVED here** (iter-17 `migrate-dev.sh`).

## Metrics delta (from `metrics.json`)
- **Gate:** 6/6 MET (distance 0), `closed-on-gate`.
- **Iters:** 17 (1 tok + 16 tiks) + 4 stabilized final harden passes; 0 orphan iters/commits.
- **Tests (rext, spot-verified GREEN at the frozen tag):** Go `playthroughs/manifest` ok + `go vet` clean;
  demo-stack Python 114; dev-stack `migrate-dev` static+shellcheck 14; TS `coverage-manifest.unit` 32; live docker
  `TestMigrateDevLive` 4 (proven 3× in harden). **Flake:** 0.
- **Close findings:** 1 (DOC-1, docs → Fate-2 close-release). 0 must-fix / 0 should-fix / 0 test gaps / 0 blends.
- **Deferral audit:** GREEN (7 in scope, 0 blocking; 1 resolved, 1 aged→re-fated fresh, 0 repeat/chronic/escape).
- **Supply-chain:** 0 net-new deps.
