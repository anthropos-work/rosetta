# M26 — Spec notes

_Technical detail. Stub at scaffold; seeded from the design + the verified orphan diff._

## The orphan is the spec (in the EXT repo, NOT this one)
- Orphan: ext repo `.agentspace/rosetta-extensions/` @ `25ab855` (branch `m26/self-contained-demo`, tag `prop-room-m26`).
- Read orphan files: `git -C .agentspace/rosetta-extensions show 25ab855:<path>`; diffs:
  `git -C .agentspace/rosetta-extensions diff 25ab855^ 25ab855 -- <path>`.
- Orphan touched 12 files (+521/−141): `demo-stack/{GUIDE.md,ant-academy.sh,ensure-clones.sh,migrate-demo.sh,rosetta-demo}`,
  `demo-stack/tests/{test_ant_academy.py,test_frontend_build.py,test_migrate_race_live.py,test_tooling.py}`,
  `stack-injection/{gen_injected_override.py,tests/test_injection.py}`.
- ext authoring branch for the re-impl: `m26/self-contained-demo-reimpl` (from `main` @ `7b17c39`); tag `understudy-m26`.

## up-injected.sh — current-main anchor points (the port targets)
- `DEV="$REPO_ROOT/stack-dev"` + `anthropos-dev` fallback (L18) — KEEP (dev side still uses it for `init_policy.sql`? no —
  up-injected reads INJECT_SVCS from `$DEV`). The build-SOURCE reads move to `$DEMO`.
- `DEMO_WS="$REPO_ROOT/stack-demo"` (L19, M30) — KEEP. Add `DEMO="$DEMO_WS"` (D2).
- `PLAT="$DEV/platform"` (L22) — MOVE to `PLAT="$DEMO_WS/platform"` (D-MAIN).
- `BASE_ENV="$PLAT/.env"` (L28) — once PLAT moves, default base = stack-demo/platform/.env; re-word the M30
  "legacy stack-dev/platform/.env base" comments (L24-27, L203-204, L218, L221, L224, L233, L236, L238).
- The early `[ -f "$PLAT/.env" ] ... exit 1` (L23): with PLAT on stack-demo + ensure-clones running below the seam
  (which now seeds `.env` only copy-if-present per D4), this guard sits ABOVE the seam at L23 currently. **Port note:**
  the orphan moved the `.env`/GH_PAT guards below the seam (gated on LIB_ONLY). On current main they are ALREADY below
  the seam (M28/M30 region, L246-247 GH_PAT guard reads from BASE_ENV). The L23 early `.env` guard must move below
  ensure-clones (or be removed — the L246 BASE_ENV GH_PAT guard already fails loud if no usable env). Verify the
  lib-only seam (L172) still returns before any `.env` read so `test_frontend_build.py` stays green.
- The seam: `[ "${UP_INJECTED_LIB_ONLY:-0}" = 1 ] && return 0` (L172). ensure-clones bare call goes just BELOW it,
  ABOVE the M28 pre-flight (L184) (D1).
- build-source reads to repoint to `$DEMO`: `build_frontend_next_web` ctx (L113), `build_frontend_studio_desk` ctx
  (L140), `src="$DEV/$svc"` (L254), cms-studio copy log (L262).
- gen call (L301-303): append `"${rd_flag[@]+"${rd_flag[@]}"}"`; add the `rd_flag` array (D5).
- **MUST-PRESERVE disjoint regions:** M30 provision block (L197-239), M31 mkcert FAPI branch (the FapiCertStep region),
  the injection-targets-COPY-only invariant (apply-authn/perl/build all on `$dst=$CLONES/$svc`, L266-273).

## ensure-clones.sh (the D4 refinement vs the orphan)
Port `25ab855:demo-stack/ensure-clones.sh` ~verbatim EXCEPT phase (b): the orphan hard-`exit 1`'d when neither
`.env` existed. D4 → non-fatal skip (copy-if-present, log a note, do NOT exit). The phase (a) `exit 1`s
(broken-partial-clone + clone-failure) stay. `set -euo pipefail`, `REPO_ROOT=$HERE/../../..`, `DEMO="$REPO_ROOT/stack-demo"`,
`PLAT="$DEMO/platform"`, `DEV_ENV="$REPO_ROOT/stack-dev/platform/.env"`. Five phases; provenance via the python3 heredoc.

## migrate-demo.sh / ant-academy.sh / rosetta-demo (D3 — current-main anchors)
- migrate-demo.sh L13: `DEV="$REPO_ROOT/stack-dev"; [ -d "$DEV" ] || DEV=".../anthropos-dev"` → `DEMO="$REPO_ROOT/stack-demo"; DEV="$DEMO"`
  (keep DEV alias → the `$DEV/$r` atlas loop + `$DEV/sentinel/init_policy.sql` reads are zero-churn). Port the atlas.hcl CAVEAT comment.
- ant-academy.sh L31: same repoint; `ACADEMY="$DEV/ant-academy/code"`; re-word the clone-not-found log; the
  `$DEV/platform/.env` Clerk-keys read moves to stack-demo too.
- rosetta-demo L23-24: `PLATFORM_DIR` default → `stack-demo/platform` (drop anthropos-dev fallback, keep env-override);
  L164 `cmd_clone` dev_root → stack-demo; add the manual-`up`-presupposes-populated-stack-demo note (L112 region).
  Preserve the M17 race guards (untouched).

## Tests (D3, D4)
- `test_tooling.py`: SECTION/REPO_ROOT/ROSETTA_DEMO defined; `shutil`/`os`/`subprocess` imported. Append
  `TestEnsureClones` (6 methods — **adapt** `test_env_copy_is_copy_if_absent_and_fails_loud` to the D4 non-fatal
  behavior per M26-D4-impl) + `TestSelfContainedSource` (4 methods). Retarget `TestRenameDrift` (add
  `demo-stack/ensure-clones.sh` to `FUNCTIONAL_SCRIPTS`; replace `test_stack_dev_is_the_preferred_default` with
  `test_dev_stack_still_prefers_stack_dev` + `test_demo_scripts_resolve_stack_demo_as_the_build_source`; add
  `test_ensure_clones_reads_stack_dev_only_for_secrets`). `TestShellcheck` += `test_ensure_clones_is_shellcheck_clean`.
  Extend the GuideDocTruth formula class list with `TestEnsureClones, TestSelfContainedSource`. Do NOT touch
  `UpInjectedSecretPreflight`/`FapiCertStep`/`PreflightBehavior`.
- `test_injection.py`: `_lines` gains `reuse_dev_images=False`; add `_sentinel_block` helper +
  `test_reuse_dev_off_by_default_builds_from_clones` + `test_image_probe_not_consulted_when_reuse_off`; retarget the
  existing 4 reuse tests to pass `reuse_dev_images=True`.
- fixture renames (stack-dev→stack-demo): `test_frontend_build.py` (TestFrontendBuildBehaviour + TestZeroPlatformRepoEdit),
  `test_ant_academy.py`, `test_migrate_race_live.py`. The `set -uo` lib-only driver in test_frontend_build.py stays byte-identical.

## D6 — the live count
Main GuideDocTruth-formula sum = 28. After M26: +1 (shellcheck) +2 (renamedrift) +6 (ensureclones) +4 (selfcontained)
= **41**. GUIDE advertises 41; the test recomputes. Re-verify the actual collected sum at the test commit.

## rosetta corpus doc-half (this worktree)
`CLAUDE.md`, `corpus/ops/demo/README.md`, `corpus/ops/demo/frontend-tier.md`, `corpus/ops/rosetta_demo.md`,
`.claude/skills/stack-secrets/SKILL.md:21`. The "true peer / own clone set" claim is now TRUE in code.

## Pre-flight audits — §1 (ensure-clones) + the whole milestone
Phase 0b KB-fidelity: **GREEN** (2026-06-15). Report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md).
Reconciliation milestone — implementation contract is the verified orphan diff (ext `25ab855`/`prop-room-m26`)
+ the doc-half §7. 3 DOC-ONLY/DOC-AHEAD findings = the planned §7 deliverables, not blockers. No blind area.
Reused across all sections (single subsystem, knowledge docs unchanged during the build).
