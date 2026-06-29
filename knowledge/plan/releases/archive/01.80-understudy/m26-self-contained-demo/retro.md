# M26 — Retro

_Self-contained demo stacks — a demo builds entirely from `stack-demo`'s OWN clone set. The single `section`
milestone of v1.8 "understudy". Two-repo (rosetta doc-half + planning; ext code @ tag `understudy-m26`).
A re-implementation of the orphaned `m26/self-contained-demo` branch onto current `main`. Closed 2026-06-15._

## Summary

M26 closed a live **doc-vs-code gap**: `CLAUDE.md` already called `stack-demo` *"a true peer of `stack-dev` (its
own clone set)"* and M30 already provisioned `stack-demo/platform/.env` — but on `main`, `up-injected.sh` still
built every demo image from `stack-dev`'s repos (`src="$DEV/$svc"`, `PLAT="$DEV/platform"`), so a box with only
`stack-demo/` could not run a demo. M26 made the implementation match the documented model.

The change re-implements the orphaned ext branch `m26/self-contained-demo` (@ `25ab855`, tag `prop-room-m26`,
authored 2026-06-14) — which predated v1.6/v1.7 and would have silently reverted the `stack-secrets` module + M30
provision + M31 mkcert + M32 studio-desk fix if merged — onto current `main`, preserving all of it. The 12-file
ext port: a NEW `demo-stack/ensure-clones.sh` (bootstrap-clone `stack-demo/platform` over SSH + seed `.env`
copy-if-present [D4] + `make init` the peer repos + `make init-studio` + a `clones.lock.json` provenance heredoc);
the `$DEV`→`stack-demo` build-source repoints across `up-injected.sh` / `migrate-demo.sh` / `ant-academy.sh` /
`rosetta-demo`; and the `reuse_dev_images=False` gate (opt back in via `DEMO_REUSE_DEV_IMAGES=1` →
`--reuse-dev-images`) in `gen_injected_override.py`. The rosetta doc-half reconciled `CLAUDE.md`'s "true peer"
claim + the demo/README flow + `frontend-tier.md` + `rosetta_demo.md`'s from-scratch section + the
`stack-secrets/SKILL.md` ensure-clones reference.

**D-MAIN was the load-bearing decision.** The orphan's "minimal" variant kept `PLAT` on `stack-dev`, which would
have left the compose file + the 4 non-Clerk services (`sentinel`/`storage`/`roadrunner`/`graphql`) bound to
`stack-dev` — the gap would NOT have been filled. D-MAIN moves `PLAT="$DEMO_WS/platform"`, so
`docker compose -f stack-demo/platform/docker-compose.yml` resolves the relative `build.context` of those services
against `stack-demo` → the WHOLE stack builds from `stack-demo`. M30's provision *behavior* (values-blind,
non-fatal fallback) was preserved; only its *paths* moved to match (provision target == compose dir).

The close-time observable-behavior gate (a demo builds with no `stack-dev` in the build graph) was satisfied **by
composition** — the accepted M31-D7 / M32-D5 pattern: the static self-containment assertions
(`TestSelfContainedSource`) + the functional ensure-clones harness (`TestEnsureClonesFunctional`, stubbed git/make,
mutation-verified) + the full green suites + the flake gate. The FULL LIVE field-bake on a freshly-emptied
`stack-demo/` is a user-authorized post-close follow-up (the on-disk `stack-demo/` is already populated from the
orphan run and would mask a from-scratch failure).

The close review found **2 findings, both Fate-1, both documentation/comment-accuracy** (no code-logic or test
gap): (1) two stale "legacy stack-dev/platform/.env base" comments in `up-injected.sh` (L36-37 + L417) — the code
was already correct (D-MAIN moved BASE_ENV to `stack-demo`), only the comments lied; re-worded (ext `773184f`).
(2) `safety.md` did not document the M26 sanctioned cross-stack read (the sole `stack-dev` read by the demo tooling
is ensure-clones' `.env` *seed*, never the build SOURCE) — blended into §2.7 (#M26-D4).

## Incidents This Cycle

None. No P0/P1/P2 incidents, no regressions, no flakes (5/5 randomized sequential on both touched suites; demo-stack
138/138, stack-injection 113/113 every run). Both harden-pass "fixes" were **test-fixture-only**, not production
bugs: the macOS `/var`→`/private/var` realpath + the `$HERE/../../..` sandbox-depth in the `TestEnsureClonesFunctional`
harness (Pass 1), and symlinking `bash` into the clean stub PATH so a stub python3's `#!/usr/bin/env bash` shebang
resolves (Pass 2). `ensure-clones.sh` and `up-injected.sh` were byte-identical before/after the harden harness fixes.

## What Went Well

- **The orphan-as-spec, re-implement-onto-main discipline paid off.** Rather than attempt an unmergeable merge of a
  pre-v1.6/v1.7 branch (which would have reverted 4 milestones of work), the design treated the orphan diff as a
  *spec* and a 3-agent fan-out + adversarial no-regression review pre-settled D-MAIN + D1–D6, verifying all 12
  orphan files were accounted for and no step reverted M30/M31/M32. The build then had a clean contract to hit.
- **The must-preserve invariants were verified at close, not assumed.** The cross-cutting code review confirmed all
  4 (M30 BASE_ENV values-blind+non-fatal, M31 mkcert, M32 studio-desk `NODE_ENV=production`+`:9000`+dropped-`:9100`,
  injection-COPY-only) intact — and the adversarial PLAT-consumer audit walked *every* `$PLAT` reference in
  `up-injected.sh` to prove the D-MAIN move re-pointed only what should move (and surfaced the stale comments).
- **The D4 refinement extended the orphan correctly.** The orphan hard-`exit 1`'d when `stack-dev/.env` was absent;
  D4 made the `.env` seed non-fatal copy-if-present (defer to M30's provisioner) so a box with only `stack-demo/`
  and `.agentspace/secrets` is fully supported — true self-containment. `TestEnsureClonesFunctional` mutation-verified
  the never-clobber + non-fatal-skip behavior (re-introducing the orphan's fail-loud fails the test).
- **The bash-3.2 empty-array hazard was caught proactively.** The new `--reuse-dev-images` shell seam is a second
  instance of the M28 empty-array-under-`set -u` class; `TestReuseFlagArrayExpansion` pins it across all 8 flag
  combinations on bash 3.2, mutation-verified against the bare-array crash. Recorded in `safety.md` §2.8 (#M26-harden).

## What Didn't

- **Nothing went wrong in the milestone.** The two close findings were both comment/doc accuracy, not behavior — a
  sign the build + harden were thorough. The one mild process note: the stale "legacy stack-dev base" comments were
  introduced *at build time* (the D-MAIN repoint moved the code but left two narrative comments behind); a tighter
  build-time comment sweep alongside a path-move would have caught them before close. Low-severity; caught at the
  Phase-2c adversarial PLAT audit exactly as intended.

## Carried Forward

- **The FULL LIVE field-bake on a freshly-emptied `stack-demo/`** — a user-authorized post-close follow-up. The
  on-disk `stack-demo/` carries a `clones.lock.json` from the orphan run; the live run must move/rename it first,
  then `/demo-up N` from scratch (ensure-clones bootstraps → M30 provision → all images build from `stack-demo` →
  verified UP, no `stack-dev` path in the build graph), then re-run for idempotency. Composition satisfied the close
  gate; this is the optional live confirmation (consistent with the M25 / M30 live field-bakes run on explicit go-ahead).
- **Orchestrator ext-side finalize (post-close):** ff ext `main` → the reimpl HEAD, re-point the `understudy-m26` tag
  forward over the 2 harden + 1 close commit, delete the orphan tag `prop-room-m26` + orphan branch
  `m26/self-contained-demo` (superseded by the reimpl). Tracked in state.md.
- **v1.8 release closure:** M26 is the only milestone — after this close the user runs `/developer-kit:close-release`
  for the release-level review + ff `release/01.80-understudy` → `main` + tag `v1.8`.
- **Release-level git carry-over** (push the ext tags `understudy-m26` + `house-lights-m31`/`m32` +
  `stage-door-m27`/`m28`/`m30` + `prop-room-m21..m25` to `origin`; `wip/clerkenstein-browser-login` awaits its own
  design-roadmap pass) — surfaces at v1.8 close-release.

## Metrics Delta

(from `metrics.json`)
- **Findings:** 2 (0 scope · 0 code-quality · 2 docs [1 = the D4 decision-triage blend] · 0 tests · 2
  adversarial-records) — both Fate-1.
- **Field/production bugs:** 0 (the 2 harden fixes were test-fixture-only).
- **Go tests:** 1027 → **1027** (+0 — M26 touched no Go).
- **Python tests (touched suites):** demo-stack **110 → 138** (+28); stack-injection **111 → 113** (+2); +30 total.
  GUIDE advertised count **41** reconciles (`TestGuideDocTruth` green). `gen_injected_override.py` 99%.
- **Flake count:** 0 (5/5 randomized sequential, both suites; 138/138 + 113/113 every run).
- **Observable-behavior gate:** MET by composition — `TestSelfContainedSource` + `TestEnsureClonesFunctional` +
  green suites + flake gate. Live field-bake = user-authorized follow-up.
- **Lint:** shellcheck clean on all 5 touched shell scripts; cross-references resolve; README-index guard exit 0.
- **Supply-chain:** GREEN (0 new deps — shell + python-stdlib + docs). **Alignment:** 100%/100% (untouched).
- **Ext code:** `understudy-m26` @ `17971c1` (reimpl, where the tag currently sits); harden `2bcaf49` (+12) +
  `0eac424` (+3); close review-fix `773184f` — all on `m26/self-contained-demo-reimpl`; the orchestrator finalizes
  the ext side (ff main + re-point tag + delete orphan tag/branch).
- **Deliverables:** the `ensure-clones.sh` peer-clone bootstrapper (ext) + the `$DEV`→`stack-demo` build-source
  repoints + the `reuse_dev_images` gate + the ported/retargeted tests + `GUIDE.md` + the rosetta doc-half + the
  planning record.
