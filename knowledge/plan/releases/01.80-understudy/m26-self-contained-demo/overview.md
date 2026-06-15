---
milestone: M26
slug: self-contained-demo
version: v1.8 "understudy"
milestone_shape: section
status: planned
created: 2026-06-15
last_updated: 2026-06-15
complexity: medium
delivers: rosetta-extensions/demo-stack + stack-injection (ensure-clones.sh + the $DEV→stack-demo build-source repoints + the reuse_dev_images gate; ext tag understudy-m26) + the rosetta corpus doc-half (CLAUDE.md / demo/README.md / frontend-tier.md / rosetta_demo.md re-author + the stack-secrets SKILL.md:21 fix)
backlog_refs: M26 self-contained-demo (orphaned ext branch m26/self-contained-demo @ 25ab855, tag prop-room-m26, authored 2026-06-14 — re-implemented here against current main, NOT merged)
---

# M26 — Self-contained demo stacks

## Goal
A demo stack builds **entirely from its own clone set** in `stack-demo/` — so a box that has **only**
`stack-demo/` (no `stack-dev/`) can bring a demo up end-to-end. This closes a live **doc-vs-code gap**:
`CLAUDE.md` already claims `stack-demo` is *"a true peer of stack-dev (its own clone set)… not a
borrower,"* and M30's secret provisioner already writes to `stack-demo/platform/.env` assuming the
clones exist there — but on `main`, `up-injected.sh` still builds every demo image from `stack-dev`'s
repos (`$DEV` source paths) and resolves the compose topology from `stack-dev/platform`. M26 makes the
implementation match the documented model.

## Why this is a re-implementation, NOT a merge
The orphaned ext branch `m26/self-contained-demo` (@ `25ab855`, tag `prop-room-m26`, +521/−141 across 12
files, authored 2026-06-14) is the **spec**, not a mergeable change. It predates v1.6 "stage door" and
v1.7 "house lights", which **rewrote the same files** (`demo-stack/up-injected.sh`,
`stack-injection/gen_injected_override.py`, the test files). A merge today would silently revert the
`stack-secrets` module, the M30 BASE_ENV provision block, the M31 mkcert FAPI cert, and the M32
studio-desk single-port fix. So M26 **ports the orphan's intent onto current `main`**, preserving all
v1.6/v1.7 work. The port spec below was produced by a 3-agent fan-out + an adversarial completeness/
no-regression review (workflow `wf_212f3442-44e`, 2026-06-15) that verified all 12 orphan files are
accounted for, no spec step reverts M30/M31/M32, and reconciled 4 internal contradictions (recorded as
design decisions D1–D4 below).

## Why section
The full change surface is known and verified file-by-file against the orphan diff and current `main`.
One new file + repoints in 4 demo scripts + 1 python generator + the test/doc updates. Build with
`/developer-kit:build-milestone`.

## Repo split
- **`rosetta-extensions`** (authoring → tag `understudy-m26` → consume per-stack): `demo-stack/ensure-clones.sh`
  (new), the `$DEV`→`stack-demo` build-source repoints in `up-injected.sh` + `migrate-demo.sh` +
  `ant-academy.sh` + `rosetta-demo`, the `reuse_dev_images` gate in `gen_injected_override.py`, and the
  ported tests + `demo-stack/GUIDE.md`.
- **`rosetta`**: re-author the dropped doc-half — `CLAUDE.md` (the "true peer" claim is now TRUE),
  `corpus/ops/demo/README.md`, `corpus/ops/demo/frontend-tier.md`, `corpus/ops/rosetta_demo.md` (the
  from-scratch self-contained bring-up flow); and fix `.claude/skills/stack-secrets/SKILL.md:21` (it
  references a `/demo-up ensure-clones` that does not exist on `main` yet — M26 makes it real).

## Design decisions (settled at design time — resolved from the verified spec)

- **D-MAIN — Full self-containment: PLAT (the compose dir) moves to `stack-demo/platform`.** This is the
  decision that actually *fills the gap* the user asked for. With `docker compose -f
  stack-demo/platform/docker-compose.yml`, the relative `build.context` of the non-Clerk compose-built
  services (`../sentinel`, `../storage`, …) resolves against `stack-demo`, so **all** services build from
  `stack-demo` — not just the injected/copied ones. Keeping PLAT on `stack-dev` (the spec's "minimal"
  variant) would leave the compose file + 4 non-Clerk services bound to `stack-dev`, so the demo still
  could not run without `stack-dev` → the gap would NOT be filled. **M30 reconciliation:** M30's provision
  target is already `stack-demo/platform/.env`, so once PLAT moves it equals `$PLAT/.env` and provision
  becomes consistent (write target == compose dir); the fallback base is the ensure-clones-seeded
  `stack-demo/platform/.env` (re-word the "legacy stack-dev/platform/.env base" comments). M30's *behavior*
  (provision values-blind, repoint `--env-file`, non-fatal fallback) is preserved — only its paths move.
- **D1 — ensure-clones placement: BELOW the `UP_INJECTED_LIB_ONLY` lib-only seam (`up-injected.sh:172`),
  ABOVE the M28 secret pre-flight (`:184`).** ensure-clones is a real-execution ACTION (network `git
  clone` + `make init`); a bare call (no `|| true` — fail-loud) just below the seam keeps the seam the
  single gate between sourced function-defs and actions, so the lib-only unit tests (`test_frontend_build.py`)
  never trigger a real clone. **NOT at the head of the file** (the orphan's pre-v1.6 idiom) — that would
  repeat the exact "above-the-seam fired at source time and crashed 20 tests" regression M28's close-review
  fixed.
- **D2 — var token: add `DEMO="$DEMO_WS"` alias.** Keep `DEMO_WS` (M30's provision-target var, referenced
  at `up-injected.sh:215/228/230`) AND add `DEMO` (the orphan's build-source name) as a one-line alias.
  Repoint the build-SOURCE reads to `$DEMO`; the orphan's ported tests assert the literal `$DEMO` token, so
  neither subsystem's references churn.
- **D3 — all-4-script repoint, no drops.** Repoint all four demo scripts (`up-injected.sh`,
  `migrate-demo.sh`, `ant-academy.sh`, `rosetta-demo`) and port BOTH conditional test files
  (`test_ant_academy.py`, `test_migrate_race_live.py`) — do NOT take the spec's "leave migrate/academy on
  `$DEV`, drop those tests" branch (it would false-fail the orphan's 4-script `TestRenameDrift`).
- **D4 — ensure-clones phase (b) `.env` copy is best-effort / non-fatal (copy-if-present, skip-if-absent),
  deferring to M30 for the real `.env`.** This is the one **refinement** of the orphan (which predated M30
  and hard-`exit 1`'d when `stack-dev/.env` was absent). M30 provisions `stack-demo/platform/.env` from
  `.agentspace/secrets`; ensure-clones only seeds a base if `stack-dev/.env` happens to be present, then M30
  overlays values-blind. **Result:** a box with no `stack-dev` but with `.agentspace/secrets` is fully
  supported — true self-containment. (A box with neither degrades loudly at the existing M28 pre-flight.)
- **D5 — `reuse_dev_images` OFF by default** (full independence: build from `stack-demo` clones). Opt back
  into dev-image reuse with `DEMO_REUSE_DEV_IMAGES=1` → `--reuse-dev-images`, gating the currently
  *unconditional* dev-image reuse at `gen_injected_override.py:252`.
- **D6 — GUIDE advertised test count pinned to the live-recomputed collection sum (~41), NOT the orphan's
  stale `40`** (the orphan's number was computed against a pre-v1.6/v1.7 ancestor; `TestGuideDocTruth`
  recomputes dynamically — trust the formula, not the literal).

## Scope
- **In — `demo-stack/ensure-clones.sh` (NEW):** port the orphan's 106-line bootstrapper ~verbatim
  (`git show 25ab855:demo-stack/ensure-clones.sh`), with the D4 phase-(b) refinement. `set -euo pipefail`;
  no args/flags (a fixed bring-up prelude); `REPO_ROOT=$HERE/../../..`. Five fail-loud idempotent phases:
  (a) bootstrap-clone `stack-demo/platform` from `git@github.com:anthropos-work/platform.git` over SSH if
  `$PLAT/.git` absent (a dir-without-`.git` = broken partial clone → `exit 1`; clone failure → `exit 1`
  with the SSH-access hint; **never** falls back to stack-dev for SOURCE; uses if-then-else not `A && B || C`
  so shellcheck stays SC2015-clean); (b) **[D4]** `.env` seed — copy `stack-dev/platform/.env` →
  `stack-demo/platform/.env` *if present and target absent*, **skip non-fatally if stack-dev is absent**
  (M30 provisions the real `.env`); (c) `make -C $PLAT init` (NOT `|| true` — Make's clone loop fails loud);
  (d) `make init-studio` for cms (non-fatal); (e) `clones.lock.json` provenance via a `python3` heredoc
  (`repo<TAB>ref<TAB>sha` per cloned repo, `json.dump` indent=2 sort_keys=True). `chmod +x`.
- **In — `up-injected.sh`:** (1) add `DEMO="$DEMO_WS"` alias [D2] and move `PLAT="$DEMO_WS/platform"`
  [D-MAIN]; re-word the M30 fallback base comment; (2) sequence `"$HERE/ensure-clones.sh"` below the L172
  seam, above the M28 pre-flight [D1]; (3) repoint frontend build ctx (`build_frontend_next_web` ~L113,
  `build_frontend_studio_desk` ~L140) and the INJECT_SVCS source (`src="$DEMO/$svc"` ~L254) + the
  cms-studio copy log to `$DEMO`; (4) add `rd_flag=(); [ "${DEMO_REUSE_DEV_IMAGES:-0}" = 1 ] &&
  rd_flag=(--reuse-dev-images)` and append it to the `gen_injected_override.py` call [D5]. **Injection
  still mutates only the per-demo COPY `$CLONES/$svc`** (`apply-authn.sh`, the Dockerfile patch, `docker
  build` all on `$dst`) — the shared `stack-demo` clone stays git-clean.
- **In — `stack-injection/gen_injected_override.py`:** add `reuse_dev_images=False` to `build_lines()` +
  the `--reuse-dev-images` CLI arg + main() wiring; gate the dev-image-reuse `elif` (L252) behind it [D5].
  **Preserve M32** (the studio-desk `NODE_ENV=production` + `FRONTEND_PORT=9000` override + the dropped
  `:9100` CORS origin — disjoint regions, untouched).
- **In — `migrate-demo.sh` / `ant-academy.sh` / `rosetta-demo`:** repoint the workspace anchor to
  `stack-demo` (keep the `DEV` alias where it makes the rest of the script zero-churn); drop the
  `anthropos-dev` legacy fallback on the demo scripts. `rosetta-demo` default `PLATFORM_DIR` →
  `stack-demo/platform` (env-overridable). Preserves the M17 race guards (untouched).
- **In — tests:** port `TestEnsureClones` + `TestSelfContainedSource` (`test_tooling.py`, append-only —
  do NOT touch `UpInjectedSecretPreflight`/`FapiCertStep`/`PreflightBehavior`); retarget `TestRenameDrift`
  (add `ensure-clones.sh`; split the dev-prefers-stack-dev test so dev-stack keeps stack-dev while the
  demo scripts resolve stack-demo as build source); `TestShellcheck` + ensure-clones; the
  `reuse_dev_images` opt-in tests (`test_injection.py`); the `stack-dev`→`stack-demo` fixture renames
  (`test_frontend_build.py`, `test_ant_academy.py`, `test_migrate_race_live.py`). Pin the `GUIDE.md` count
  to the recomputed sum [D6]. **No strictness regression** (the spec verified `test_frontend_build.py`'s
  `set -uo` lib-only driver is byte-identical to main and intentional).
- **In — `demo-stack/GUIDE.md`:** self-containment prose (step 0 = ensure-clones; build from stack-demo;
  injection on the per-demo copy only) + the pinned test count.
- **In — rosetta corpus doc-half:** re-author `CLAUDE.md` (reconcile the now-true "true peer" claim),
  `corpus/ops/demo/README.md`, `corpus/ops/demo/frontend-tier.md`, `corpus/ops/rosetta_demo.md` (the
  from-scratch self-contained bring-up); fix `.claude/skills/stack-secrets/SKILL.md:21`.
- **Out:** `stack-dev` (already self-contained on its own clones — unchanged); any platform-repo edit
  (`stack-demo/platform` is a build *context* only, never edited — the zero-platform-repo-edit line holds);
  new third-party deps (shell + python-stdlib + docs).

## After ship
Delete the orphan tag `prop-room-m26` + branch `m26/self-contained-demo` (superseded by this
re-implementation).

## Field-bake (close-time, observable-behavior gate)
Prove on a **truly empty** `stack-demo/`: the on-disk `stack-demo/` is already populated (a
`clones.lock.json` from the orphaned run) which would **mask a from-scratch failure** — move/rename it
first. Then `/demo-up <N>` from scratch: ensure-clones bootstraps `stack-demo`'s own clone set → M30
provision → all images build from `stack-demo` → verified UP. Confirm **no `stack-dev` path in the build
graph** (the self-containment assertion). Re-run once (idempotency: clones reused, not re-cloned).

**Depends on:** none (single milestone of v1.8).
**Parallel with:** none.
**Estimated complexity:** medium (one new file + repoints across 6 code files + 5 test files + the doc-half;
the PLAT-move [D-MAIN] is the one part that touches M30's path wiring → re-point the M28/M30 pre-flight test
expectations, behavior preserved).
**Open questions:** none blocking — D-MAIN (PLAT moves), D1–D6 all settled at design.
**KB dependencies:** `corpus/ops/rosetta_demo.md` (the demo lifecycle); `corpus/ops/demo/README.md` +
`frontend-tier.md` (the demo family + UI tier); `corpus/ops/secrets-spec.md` + `.claude/skills/stack-secrets/SKILL.md`
(the M30 provision the ensure-clones `.env` seed layers under); `corpus/ops/safety.md` (the sanctioned
stack-dev read is `.env`-copy-only, never SOURCE).
**Delivers →** `rosetta-extensions/demo-stack` + `stack-injection` (ext tag `understudy-m26`) + the rosetta
corpus doc-half.

## Risks
| Risk | Severity | Mitigation |
|---|---|---|
| PLAT-move [D-MAIN] touches M30's path wiring → M28/M30 pre-flight tests may assert `stack-dev/platform/.env` as the BASE_ENV default | degrades-quality | re-point those test expectations; M30 *behavior* (values-blind provision, non-fatal fallback) preserved — only paths move |
| ensure-clones runs a real `git clone` + `make init` (network/SSH) → must never fire in static unit tests | breaks-tests | placement BELOW the L172 lib-only seam [D1] + `TestEnsureClones` static-only fences (assert the script text, never execute) |
| on-disk `stack-demo/` already populated (orphan's `clones.lock.json`) masks a from-scratch failure | gives-false-green | field-bake on a freshly-emptied `stack-demo/` (rename the existing one first) |
| ensure-clones `.env` seed vs M30 provision both write `stack-demo/platform/.env` | nice-to-resolve | D4 (copy-if-present, non-fatal) + sequencing (ensure-clones below-seam-above-M28, before M30 provision) — they layer: seed then values-blind overlay |
| manual `rosetta-demo up` now presupposes a populated `stack-demo` (it neither clones nor provisions) | behavior-change (accepted) | documented: operator runs `up-injected.sh`/ensure-clones first; the auto `/demo-up` path handles it |
