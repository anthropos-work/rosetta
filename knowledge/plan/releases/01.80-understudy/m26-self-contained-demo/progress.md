# M26 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Sections (ext code → rosetta docs; ext tag `understudy-m26`)

### §1 — `ensure-clones.sh` (NEW) — the peer-clone bootstrapper
- [x] port the orphan's `demo-stack/ensure-clones.sh` (~106 lines) with the D4 phase-(b) refinement
      (`.env` seed is copy-if-present, **skip non-fatally** if `stack-dev/.env` absent — defers to M30)
- [x] phases (a) bootstrap-clone `stack-demo/platform` over SSH fail-loud · (b) [D4] non-fatal `.env` seed ·
      (c) `make -C $PLAT init` (no `|| true`) · (d) `make init-studio` (non-fatal) · (e) `clones.lock.json` provenance
- [x] `set -euo pipefail`, `chmod +x`, shellcheck-clean (if-then-else, no `A && B || C`)

### §2 — `up-injected.sh` — DEMO alias + PLAT move + ensure-clones sequencing + build-ctx repoints + reuse flag
- [x] [D2] add `DEMO="$DEMO_WS"` alias (keep `DEMO_WS`); [D-MAIN] move `PLAT="$DEMO_WS/platform"`
- [x] [D-MAIN] re-word the M30 BASE_ENV "legacy stack-dev/platform/.env base" comments (base is now stack-demo)
- [x] [D1] sequence `"$HERE/ensure-clones.sh"` (bare, no `|| true`) BELOW the L172 `UP_INJECTED_LIB_ONLY` seam,
      ABOVE the M28 pre-flight
- [x] [D2] repoint build SOURCE: `build_frontend_next_web` ctx, `build_frontend_studio_desk` ctx, `src="$DEMO/$svc"`,
      cms-studio copy log → `$DEMO`
- [x] [D5] add `rd_flag=(); [ "${DEMO_REUSE_DEV_IMAGES:-0}" = 1 ] && rd_flag=(--reuse-dev-images)` + append to the
      `gen_injected_override.py` call
- [x] MUST-PRESERVE: M30 provision block (values-blind, non-fatal fallback), M31 mkcert branch, the seam, the
      injection-targets-COPY-only invariant

### §3 — `gen_injected_override.py` — `reuse_dev_images` gate [D5]
- [x] `build_lines(... reuse_dev_images=False)` param + gate the dev-image-reuse `elif` behind it
- [x] `--reuse-dev-images` CLI arg (default False) + `main()` wiring
- [x] MUST-PRESERVE M32 (studio-desk `NODE_ENV=production` + `FRONTEND_PORT=9000` + dropped `:9100` CORS — disjoint)

### §4 — `migrate-demo.sh` / `ant-academy.sh` / `rosetta-demo` repoint [D3]
- [x] migrate-demo.sh: `DEMO="$REPO_ROOT/stack-demo"; DEV="$DEMO"` (keep DEV alias, zero-churn); drop `anthropos-dev`
      fallback; port the atlas.hcl CAVEAT comment
- [x] ant-academy.sh: same anchor repoint; re-word the clone-not-found log
- [x] rosetta-demo: `PLATFORM_DIR` default → `stack-demo/platform` (env-overridable); `cmd_clone` dev_root → stack-demo;
      drop `anthropos-dev` fallbacks; the manual-`up`-presupposes-populated-stack-demo note. Preserve M17 race guards

### §5 — tests (ext) — port + retarget + fixture renames [D3, D4]
- [x] `test_tooling.py`: append `TestEnsureClones` (with the **D4** non-fatal-`.env` adaptation, NOT the orphan's
      fail-loud) + `TestSelfContainedSource`; retarget `TestRenameDrift` (add ensure-clones.sh; split the
      dev-prefers-stack-dev test); `TestShellcheck` + ensure-clones; extend the GuideDocTruth formula class list.
      Do NOT touch `UpInjectedSecretPreflight`/`FapiCertStep`/`PreflightBehavior`
- [x] `test_injection.py`: the `reuse_dev_images` opt-in tests (default-off builds-from-clones + probe-not-consulted +
      retarget the existing reuse tests to `reuse_dev_images=True`)
- [x] `test_frontend_build.py` / `test_ant_academy.py` / `test_migrate_race_live.py`: `stack-dev`→`stack-demo` fixture
      renames (no strictness regression — `test_frontend_build.py`'s `set -uo` lib-only driver stays byte-identical)

### §6 — `demo-stack/GUIDE.md` — self-containment prose + the pinned [D6] count
- [x] self-containment prose (step 0 = ensure-clones; build from stack-demo; injection on the per-demo COPY only)
- [x] [D6] pin the advertised test count to the **live-recomputed** GuideDocTruth-formula sum (projected **41**),
      NOT the orphan's stale 40

### §7 — rosetta corpus doc-half (THIS worktree)
- [x] `CLAUDE.md` — the "true peer (own clone set)" claim is now TRUE in code; reconcile any stale "borrows stack-dev"
      framing
- [x] `corpus/ops/demo/README.md` — the up→… flow gains the ensure-clones step-0
- [x] `corpus/ops/demo/frontend-tier.md` — build source is stack-demo's own clones
- [x] `corpus/ops/rosetta_demo.md` — the from-scratch self-contained bring-up flow
- [x] `.claude/skills/stack-secrets/SKILL.md:21` — the `/demo-up ensure-clones` reference is now real

## Notes
- ext authoring branch: `m26/self-contained-demo-reimpl` (the orphan branch name `m26/self-contained-demo` is
  taken by the orphan @ `25ab855`/`prop-room-m26`; the orchestrator finalizes ff-to-main + tag-repoint at close).
- Field-bake (a truly-empty `stack-demo/`) is **close-time**, not this build — build is static-verifiable only.

## Build outcome (2026-06-15)
All 7 sections landed. **Two-repo split:**
- **ext** (`m26/self-contained-demo-reimpl` @ `17971c1`, tagged **`understudy-m26`**): the 12-file port —
  `ensure-clones.sh` (new) + `up-injected.sh` / `gen_injected_override.py` / `migrate-demo.sh` / `ant-academy.sh`
  / `rosetta-demo` repoints + the ported/retargeted tests + `GUIDE.md`. Phase 0b GREEN; **demo-stack 123/123 +
  stack-injection 113/113** pass (python3.11/PyYAML); all 5 touched scripts shellcheck-clean. D6 GUIDE count =
  live-recomputed **41** (not the orphan's stale 40).
- **rosetta** (this worktree, commit on `m26/self-contained-demo`): the doc-half — `CLAUDE.md` /
  `corpus/ops/demo/README.md` / `corpus/ops/demo/frontend-tier.md` / `corpus/ops/rosetta_demo.md` +
  `.claude/skills/stack-secrets/SKILL.md` + the milestone planning files.
- **Must-preserve verified intact:** M30 BASE_ENV provision (values-blind, non-fatal fallback — paths repointed
  to stack-demo per D-MAIN, behavior preserved), M31 mkcert FAPI branch, M32 studio-desk single-port override.
- **PR review:** CLEAN — 0 issues; D1–D6 + D-MAIN all faithfully applied; the `--reuse-dev-images` flag agrees
  across parser / caller / docs.
- **Remaining for close:** field-bake on a freshly-emptied `stack-demo/`; the orchestrator's ff-to-main +
  `understudy-m26` tag-repoint + orphan-tag (`prop-room-m26`) / orphan-branch deletion.

## M26: Hardening

### Pass 1 — 2026-06-15

**Scope manifest (M26-touched, ext-repo `7b17c39..17971c1`, 12 files):**
| File | Stack | Existing tests | Coverage / gap |
|---|---|---|---|
| `demo-stack/ensure-clones.sh` (NEW) | shell | `TestEnsureClones` (8 STATIC text-pins only) | **the gap** — 116-line script, fail-loud + .env-seed (D4) + provenance + idempotency, ZERO functional execution coverage |
| `demo-stack/up-injected.sh` | shell | `TestSelfContainedSource` + `UpInjectedSecretPreflight` + `FapiCertStep` (static + functional cert block) | source-repoint static-fenced; M30/M31/M32 preserved |
| `stack-injection/gen_injected_override.py` | python | `TestGenInjectedOverride` + `TestFrontendTier` + reuse-flag tests | **99%** (only the `if __name__` guard line 393 missed — `main()` IS covered); reuse_dev_images gate fully tested |
| `demo-stack/migrate-demo.sh` | shell | `TestMigrateRaceGuard` + `TestSetEraceGuards` | repoint only (DEV=DEMO alias); race guards intact |
| `demo-stack/ant-academy.sh` | shell | `test_ant_academy.py` | repoint only |
| `demo-stack/rosetta-demo` | shell | `RosettaDemoRegistry` (functional) | repoint only |
| `demo-stack/GUIDE.md` | doc | `TestGuideDocTruth` (live-recomputed) | count pinned dynamically |
| 4 test files | python | self | retargeted/ported per D3/D4 |

**Coverage measure:** `gen_injected_override.py` **99%** (1 miss = the unreachable `if __name__` guard). Shell
scripts are exercised via subprocess (coverage.py can't instrument them — the in-house pattern is functional
extract-and-run harnesses, as `FapiCertStep` does for the cert block).

**Identified deepening target (Fate 1 — LAND NOW):** `ensure-clones.sh` is the milestone's one NEW unit and is
covered only by static body-pins. Build a functional harness (stub `git`/`make`/`python3` on PATH, the
`RosettaDemoRegistry`/`FapiCertStep` pattern) exercising: the broken-partial-clone exit-1 guard; the D4 .env seed
(copy-if-present + non-fatal-skip-if-absent + never-clobber — the riskiest behavior, the orphan's divergence);
the provenance lockfile (valid sorted JSON); `make init-studio` non-fatal; idempotency (re-run reuses, never
re-clones); the fail-loud-on-source-clone-failure path. All reachable without a real network clone.

**Tests added (Pass 1, ext `2bcaf49`): `TestEnsureClonesFunctional` (+12, `demo-stack/tests/test_tooling.py`)**
— a sandboxed functional harness that copies `ensure-clones.sh` into a depth-matched temp tree (so its
`REPO_ROOT=$HERE/../../..` resolves into the sandbox), stubs `git`/`make` on a constrained PATH (no live network —
a successful stub clone materialises `$PLAT/.git`+`repos.yml` so re-runs are idempotent + phase (e)'s `awk` runs),
and uses the REAL `python3` for the provenance heredoc:
- (a) broken-partial-clone → exit 1; clone-failure → exit 1 + SSH hint + no-stackdev-fallback + no make/lockfile;
  fresh bootstrap clones-then-make-init; present-platform reused (idempotent, no re-clone).
- (b) [D4] .env seed copies byte-for-byte from stack-dev when present+target-absent; **non-fatal skip (exit 0,
  not the orphan's exit 1) when stack-dev absent**; **never-clobbers** a pre-existing/provisioned stack-demo .env.
- (d) `make init-studio` failure non-fatal (exit 0 + warn); skipped when `cms/studio` already present.
- (e) provenance lockfile is valid **sorted** JSON `{repo:{ref,sha}}`; skips `.git`-less repos; full re-run
  idempotent (no re-clone, stable lockfile + .env).
- **Mutation-verified:** re-introducing the orphan's fail-loud .env (exit 1), defeating the never-clobber guard,
  and deleting the broken-partial-clone abort each fail the corresponding test (not shallow box-tickers).

**Bugs fixed inline:** none (no production bug surfaced — the script's behavior matched the D-decisions; the two
harness iterations were test-fixture bugs: the macOS `/var`→`/private/var` realpath + the `$HERE/../../..`
depth, both fixed in the harness, not the script).

### Pass 2 — 2026-06-15

**Coverage:** `gen_injected_override.py` unchanged at **99%** (no Python touched; the deepening is shell-side,
which coverage.py cannot instrument — covered functionally + mutation-verified instead).

**Tests added (ext `0eac424`): `TestReuseFlagArrayExpansion` (+3, `demo-stack/tests/test_frontend_build.py`)**
— the `--reuse-dev-images` SHELL seam (the Python `reuse_dev_images` arg was already saturated in
`test_injection.py`; this covers the OTHER half — up-injected.sh's `rd_flag` assembly + the
`"${rd_flag[@]+"${rd_flag[@]}"}"` expansion passed to the generator). The hazard: up-injected.sh runs
`set -euo pipefail`, and on **bash 3.2** (the macOS system bash) a BARE `"${rd_flag[@]}"` on an EMPTY array trips
`set -u` 'unbound variable' and aborts the whole bring-up at override-generation — the same empty-array class that
crashed the M28 pre-flight. The tests extract the real flag-assembly + call block, stub `python3` to record argv,
and run it under `set -euo pipefail` on bash 3.2:
- empty-default (reuse OFF, all 3 flag arrays empty) → exit 0, no 'unbound variable', no opt-in flags in argv,
  but the always-present generator args DID reach it;
- `DEMO_REUSE_DEV_IMAGES=1` → `--reuse-dev-images` in argv; unset → absent (the D5 opt-in seam, shell→argv);
- all 8 (ui, lc, reuse) combinations: each exit 0 + carries exactly the flags whose env var is set.
- **Mutation-verified:** regressing the `+`-guard to the bare `"${rd_flag[@]}"` makes the empty-default case fail
  with the genuine **runtime** `ui_flag[@]: unbound variable` crash (the `_block()` extractor re-anchored on
  `rd_flag[@]` in any form so the RUNTIME crash — not a missing literal — is the signal).

**Bugs fixed inline:** none (the production guard is correct; the one harness fix was symlinking `bash` into the
clean stub PATH so the stub python3's `#!/usr/bin/env bash` shebang resolves — a fixture artifact).

### Stop condition
Stopped after Pass 2 (re-measure as Pass 3). The full Step 2b scan found nothing new worth adding: the two
highest-risk M26 surfaces (the NEW `ensure-clones.sh` + the M26-new `--reuse-dev-images` shell seam) now have
mutation-verified functional coverage; `gen_injected_override.py` is at 99% (the 1 miss is the unreachable
`if __name__` guard — `main()` IS covered); the repoint-only scripts (`migrate-demo.sh`, `ant-academy.sh`,
`rosetta-demo`) are covered by the live docker migrate-race harness (retargeted to stack-demo), the
`TestRenameDrift` no-stack-dev-source static fence (verified: every code-level `stack-dev` is confined to
ensure-clones.sh's `DEV_ENV` .env-seed), and the existing behavioral `test_ant_academy.py`/`RosettaDemoRegistry`
suites. Python coverage delta 0 (saturated); no flakes (3 clean runs below). demo-stack **123 → 138** (+15
functional); stack-injection 113 (unchanged — its M26 surface was already saturated).

## M26: Final Review (close — 2026-06-15)

Review fan-out (Phases 1–5): scope CLEAN (all 7 sections + the 2 harden passes); deferral re-audit **GREEN**
(0 deferrals — M26 is the only milestone of v1.8, nothing inherited); code-quality CLEAN (the Explore cross-cut
verified all 4 must-preserve invariants — M30 BASE_ENV / M31 mkcert / M32 studio-desk / injection-COPY-only —
intact; no dead code; reuse flag agrees parser↔caller↔docs; lib-only seam returns before ensure-clones);
shellcheck 5/5; demo-stack **138/138** + stack-injection **113/113** (python3.11, JUnit-authoritative).
Adversarial (Phase 2c): 2 scenarios probed — both clean — one surfaced the stale-comment finding below.

### Scope
- [x] No gaps — overview `In:` list fully delivered; no silent drops; 0 TODO/FIXME in any touched file.

### Code Quality
- [x] [must-fix] none — review CLEAN, 4 must-preserve invariants verified intact.

### Documentation
- [x] [should-fix] up-injected.sh L36-37 + L417 still call the BASE_ENV fallback "the legacy
      stack-dev/platform/.env base"/"otherwise" — but post-D-MAIN the actual fallback base is the
      ensure-clones-seeded **stack-demo**/platform/.env (`BASE_ENV="$PLAT/.env"`, `PLAT="$DEMO_WS/platform"`).
      Code correct; comments stale. Re-word per D-MAIN. (ext fix)
- [x] [should-fix] `corpus/ops/safety.md` does not document the M26 **sanctioned cross-stack read** invariant
      (the sole sanctioned `stack-dev` read by the demo tooling is ensure-clones' `.env` *seed* — copy-if-present,
      never the build SOURCE). It's a new read-side boundary + a user-facing safety claim (spec-notes flagged
      safety.md as a KB dependency for exactly this). Blend into Part 1 / §2.7. (= Decision-triage D4 blend below.)

### Tests & Benchmarks
- [x] No gaps — the milestone's one NEW unit (`ensure-clones.sh`) has mutation-verified functional coverage
      (`TestEnsureClonesFunctional` +12); the M26-new `--reuse-dev-images` shell seam (`TestReuseFlagArrayExpansion`
      +3, bash-3.2 mutation-verified); GUIDE advertised count **41** reconciles (`TestGuideDocTruth` green).

### Decision Triage
- [x] D-MAIN, D1, D2, D3, D5, D6 → already blended into knowledge during build (rosetta_demo.md self-contained
      section + CLAUDE.md "true peer" + README flow + safety.md §2.8 #M26-harden). Verified accurate; reference
      tags present. No new blend.
- [x] **D4 (the sanctioned `.env`-seed-only stack-dev read) → blend into `safety.md`** (= the doc finding above).
- [x] M26-D2-impl / M26-D4-impl / M26-D6-impl → **archive** (maintainer-only: alias reconciliation, the ported
      test-assertion change, the live count=41). Stay in decisions.md.

## Completeness Ledger (close — 2026-06-15)

### Done (Fate 1 — landed in this milestone, properly and completely)
- **§1 `demo-stack/ensure-clones.sh` (NEW)** — the 115-line peer-clone bootstrapper (5 phases, D4 non-fatal
  `.env` seed) — verified in `TestEnsureClones` (static pins) + `TestEnsureClonesFunctional` (+12, mutation-verified).
- **§2 `up-injected.sh`** — `DEMO=$DEMO_WS` alias [D2] + `PLAT`→stack-demo [D-MAIN] + ensure-clones below the
  L172 seam [D1] + build-SOURCE repoints to `$DEMO` + the `rd_flag`/`--reuse-dev-images` seam [D5] — verified in
  `TestSelfContainedSource` + `TestRenameDrift` + `TestReuseFlagArrayExpansion`; M30/M31/M32 invariants intact.
- **§3 `gen_injected_override.py`** — `reuse_dev_images=False` param + `--reuse-dev-images` CLI + the gated reuse
  `elif` [D5] — verified in `test_injection.py` reuse-flag suite (99% coverage; M32 region untouched).
- **§4 `migrate-demo.sh` / `ant-academy.sh` / `rosetta-demo`** — workspace anchor repointed to stack-demo [D3],
  `anthropos-dev` fallback dropped on the demo scripts, M17 race guards preserved — verified by the retargeted
  `test_migrate_race_live.py` + `test_ant_academy.py` + `RosettaDemoRegistry`.
- **§5 tests** — `TestEnsureClones`/`TestSelfContainedSource` appended, `TestRenameDrift`/`TestShellcheck`
  retargeted, the reuse opt-in tests, the `stack-dev`→`stack-demo` fixture renames — verified: full suites green.
- **§6 `demo-stack/GUIDE.md`** — self-containment prose (step 0 = ensure-clones) + the pinned [D6] count **41** —
  verified `TestGuideDocTruth` reconciles dynamically (green).
- **§7 rosetta corpus doc-half** — `CLAUDE.md` "true peer" reconcile, `demo/README.md` flow, `frontend-tier.md`
  academy path, `rosetta_demo.md` from-scratch section, `safety.md` §2.7 sanctioned-read blend [#M26-D4],
  `stack-secrets/SKILL.md:21` — verified accurate; all cross-references resolve.
- **Field-bake (close-time observable-behavior gate) — satisfied BY COMPOSITION** (the accepted M31-D7 / M32-D5
  pattern): (1) static self-containment assertions (`TestSelfContainedSource`); (2) the functional ensure-clones
  harness (`TestEnsureClonesFunctional`); (3) full unit suites green (138 + 113) + flake gate 5/5 each. The FULL
  LIVE field-bake on a freshly-emptied `stack-demo/` is a **user-authorized post-close follow-up** (the
  orchestrator offers it after close; consistent with how the M25 / M30 live field-bakes ran only on explicit
  go-ahead). NOT a closed-incomplete condition — composition satisfies the section gate.

### Confirmed-covered by another milestone in this release (Fate 2)
- None (M26 is the only milestone of v1.8).

### Annotated to a milestone of this release (Fate 3)
- None.

### Dropped (cut from roadmap entirely)
- None.

### Release-scope-breaking deferral (escape hatch — requires user sign-off)
- None.

**All scope items delivered in this milestone (Fate 1). Nothing routed, dropped, or escape-hatch-deferred.**
