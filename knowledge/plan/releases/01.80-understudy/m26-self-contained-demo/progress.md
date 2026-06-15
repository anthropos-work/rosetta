# M26 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Sections (ext code → rosetta docs; ext tag `understudy-m26`)

### §1 — `ensure-clones.sh` (NEW) — the peer-clone bootstrapper
- [ ] port the orphan's `demo-stack/ensure-clones.sh` (~106 lines) with the D4 phase-(b) refinement
      (`.env` seed is copy-if-present, **skip non-fatally** if `stack-dev/.env` absent — defers to M30)
- [ ] phases (a) bootstrap-clone `stack-demo/platform` over SSH fail-loud · (b) [D4] non-fatal `.env` seed ·
      (c) `make -C $PLAT init` (no `|| true`) · (d) `make init-studio` (non-fatal) · (e) `clones.lock.json` provenance
- [ ] `set -euo pipefail`, `chmod +x`, shellcheck-clean (if-then-else, no `A && B || C`)

### §2 — `up-injected.sh` — DEMO alias + PLAT move + ensure-clones sequencing + build-ctx repoints + reuse flag
- [ ] [D2] add `DEMO="$DEMO_WS"` alias (keep `DEMO_WS`); [D-MAIN] move `PLAT="$DEMO_WS/platform"`
- [ ] [D-MAIN] re-word the M30 BASE_ENV "legacy stack-dev/platform/.env base" comments (base is now stack-demo)
- [ ] [D1] sequence `"$HERE/ensure-clones.sh"` (bare, no `|| true`) BELOW the L172 `UP_INJECTED_LIB_ONLY` seam,
      ABOVE the M28 pre-flight
- [ ] [D2] repoint build SOURCE: `build_frontend_next_web` ctx, `build_frontend_studio_desk` ctx, `src="$DEMO/$svc"`,
      cms-studio copy log → `$DEMO`
- [ ] [D5] add `rd_flag=(); [ "${DEMO_REUSE_DEV_IMAGES:-0}" = 1 ] && rd_flag=(--reuse-dev-images)` + append to the
      `gen_injected_override.py` call
- [ ] MUST-PRESERVE: M30 provision block (values-blind, non-fatal fallback), M31 mkcert branch, the seam, the
      injection-targets-COPY-only invariant

### §3 — `gen_injected_override.py` — `reuse_dev_images` gate [D5]
- [ ] `build_lines(... reuse_dev_images=False)` param + gate the dev-image-reuse `elif` behind it
- [ ] `--reuse-dev-images` CLI arg (default False) + `main()` wiring
- [ ] MUST-PRESERVE M32 (studio-desk `NODE_ENV=production` + `FRONTEND_PORT=9000` + dropped `:9100` CORS — disjoint)

### §4 — `migrate-demo.sh` / `ant-academy.sh` / `rosetta-demo` repoint [D3]
- [ ] migrate-demo.sh: `DEMO="$REPO_ROOT/stack-demo"; DEV="$DEMO"` (keep DEV alias, zero-churn); drop `anthropos-dev`
      fallback; port the atlas.hcl CAVEAT comment
- [ ] ant-academy.sh: same anchor repoint; re-word the clone-not-found log
- [ ] rosetta-demo: `PLATFORM_DIR` default → `stack-demo/platform` (env-overridable); `cmd_clone` dev_root → stack-demo;
      drop `anthropos-dev` fallbacks; the manual-`up`-presupposes-populated-stack-demo note. Preserve M17 race guards

### §5 — tests (ext) — port + retarget + fixture renames [D3, D4]
- [ ] `test_tooling.py`: append `TestEnsureClones` (with the **D4** non-fatal-`.env` adaptation, NOT the orphan's
      fail-loud) + `TestSelfContainedSource`; retarget `TestRenameDrift` (add ensure-clones.sh; split the
      dev-prefers-stack-dev test); `TestShellcheck` + ensure-clones; extend the GuideDocTruth formula class list.
      Do NOT touch `UpInjectedSecretPreflight`/`FapiCertStep`/`PreflightBehavior`
- [ ] `test_injection.py`: the `reuse_dev_images` opt-in tests (default-off builds-from-clones + probe-not-consulted +
      retarget the existing reuse tests to `reuse_dev_images=True`)
- [ ] `test_frontend_build.py` / `test_ant_academy.py` / `test_migrate_race_live.py`: `stack-dev`→`stack-demo` fixture
      renames (no strictness regression — `test_frontend_build.py`'s `set -uo` lib-only driver stays byte-identical)

### §6 — `demo-stack/GUIDE.md` — self-containment prose + the pinned [D6] count
- [ ] self-containment prose (step 0 = ensure-clones; build from stack-demo; injection on the per-demo COPY only)
- [ ] [D6] pin the advertised test count to the **live-recomputed** GuideDocTruth-formula sum (projected **41**),
      NOT the orphan's stale 40

### §7 — rosetta corpus doc-half (THIS worktree)
- [ ] `CLAUDE.md` — the "true peer (own clone set)" claim is now TRUE in code; reconcile any stale "borrows stack-dev"
      framing
- [ ] `corpus/ops/demo/README.md` — the up→… flow gains the ensure-clones step-0
- [ ] `corpus/ops/demo/frontend-tier.md` — build source is stack-demo's own clones
- [ ] `corpus/ops/rosetta_demo.md` — the from-scratch self-contained bring-up flow
- [ ] `.claude/skills/stack-secrets/SKILL.md:21` — the `/demo-up ensure-clones` reference is now real

## Notes
- ext authoring branch: `m26/self-contained-demo-reimpl` (the orphan branch name `m26/self-contained-demo` is
  taken by the orphan @ `25ab855`/`prop-room-m26`; the orchestrator finalizes ff-to-main + tag-repoint at close).
- Field-bake (a truly-empty `stack-demo/`) is **close-time**, not this build — build is static-verifiable only.
