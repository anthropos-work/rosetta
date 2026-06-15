# M32 — Retro

_studio-desk single-port / production alignment + the `:9100` sweep. The 2nd and FINAL milestone of v1.7 "house
lights". Two-repo (rosetta `:9100` doc/CORS sweep + planning artifacts; ext code @ tag `house-lights-m32`).
Closed 2026-06-15._

## Summary

M32 fixed the sibling demo-UI defect surfaced alongside the M31 browser-login investigation: a fresh browser at
demo-N's studio-desk (e.g. `http://localhost:39000`) landed on a **302 to the dead `:9100`** instead of a live
page. Root cause is a three-layer Docker-env precedence chain (verified, not assumed): the studio-desk image
(`Dockerfile.dev`) is a *production* build that even bakes `ENV NODE_ENV=production`, BUT the base platform
`docker-compose.yml` studio-desk service sets `NODE_ENV=development` + `FRONTEND_PORT=9100` in its `environment:`
block — and a compose `environment:` value **overrides the image's baked `ENV`** (#M32-D4). Because the demo
override's per-frontend env block is deliberately **additive** (not `!override`, so inherited `PORT`/`VITE_*`
survive), that `development` survived into the demo → `src/index.ts` `isProduction=false` → the dev block ran
`res.redirect('http://localhost:9100/home')`, a dead port (the production image is single-port `9000`; `9100` is
the Vite dev port that exists only under `npm run dev`, never in the container).

The fix is a 1-line-class change: pin `NODE_ENV=production` (+ `FRONTEND_PORT=9000` belt-and-suspenders) in the
studio-desk override's env dict, winning the additive merge back to the production `sendFile` path. Route coverage
was the load-bearing open question — verified by code-read of `stack-demo/studio-desk/src/index.ts` (#M32-D1, "NO
GAP"): the production block serves every dev-block route via `sendFile` + an `express.static(dist/public)` mount
(serving the `.html`-extension targets the dev block redirected to) + an `index.html` SPA `*` fallback (strictly
better than the dev `*` catch-all, which only ever bounced to the now-absent Vite dev server). The un-offset
`:9100` CORS origin was dropped (dead now; #M32-D2), and the `:9100`→single-port-`9000` sweep ran across
`frontend-tier.md` (port row + example + the new mechanism block + CORS note + verify registry) + the demo-up SKILL.

The close-time observable verify was satisfied by **composition** (#M32-D5, mirroring M31-D7): demo-3 had been
torn down, so rather than re-spin a fresh `/demo-up`, the claim was proven as the production pin is set (the
regression test, mutation-checked 4 ways) → `isProduction=true` → the production code path serves on the single
`9000` port, no dead-`:9100` 302, no 404 (the code-read route-coverage verdict). At close, this composition was
*strengthened with a live merge-probe* (Phase 2c adversarial): the demo override was built, parsed via the repo's
`!override`-aware loader, and Docker's additive list-merge against base `[development, 9100, PORT=9000]` was
simulated — the override's `production`/`9000` pins WIN (last-`VAR=`-wins), exactly one `environment:` key, no
`9100` anywhere. The #M32-D4 precedence chain is now confirmed on the real generated artifact, not just code-read.

The close review found **4 findings, all Fate-1**: 3 decision-tag blends (D1/D2/D4 reference-tagged into
`frontend-tier.md`'s single-port + CORS notes) + 1 adversarial-scenario record (the additive-merge scenario,
defended live). No code changed beyond the doc tags — the milestone shipped clean from build + harden.

## Incidents This Cycle

None at close. No P0/P1/P2 incidents, no regressions, no flakes (5/5 randomized-sequential runs, all 88/88).

Two **latent env-masked test bugs** were surfaced + fixed Fate-1 during build (P2-class, NOT production bugs, both
pre-existing on the `house-lights-m31` tag; #M32-D3): (1) `test_frontend_blocks_parse_to_valid_compose` asserted
studio-desk ports `["29000:9000","29100:9100"]` while the generator only ever emitted the single `["29000:9000"]`;
(2) `test_a_plain_service_parses_to_ports_only` predated the universal fix16/17 prod-`DIRECTUS_TOKEN=` strip. Both
were masked because the YAML-structural tests `@skipIf(yaml is None)` and PyYAML is absent in the default `python3`
env — they only fail when PyYAML IS present. The lesson: an env-gated test tier can silently rot; running the suite
under both modes (py3.11+PyYAML and py3.14-no-PyYAML) at build is what surfaced them.

One **test-rigour gap** was closed during harden (commit `7b17c39`, P2-class): the CORS test had only membership
asserts (`19000` present, `19100` absent) and did NOT pin the full surviving origin set — so over-removing the kept
`3001` next-web origin passed the entire 88-test suite (mutation-confirmed). The new exact-set assertion (`origins
== [3000,3001,9000]+offset` in emit order) catches over-removal of any kept origin AND an accidental re-add.

## What Went Well

- **Code-read-before-trust on the load-bearing question.** The "does production cover all routes?" risk was the
  one that could have shipped a 404. Reading `src/index.ts` 148-272 directly (the byte-pristine platform copy) gave
  a definitive NO-GAP verdict before flipping the flag — no guesswork, no "the smoke will tell us."
- **The composition-close discipline held for the 2nd time in v1.7.** No demo was up; standing up
  studio-desk-in-production standalone would have been disproportionate. The necessary+sufficient chain (pin set →
  production path → covered routes) plus the live merge-probe at close gave real confidence without docker.
- **Dual-PyYAML-mode runs caught the latent bugs.** Running both interpreters at build (not just the default) is
  what un-masked the two env-gated stale assertions that had been green-by-skip since before M31.
- **Small, well-scoped surface.** Two ext files + a doc sweep; the harden + close passes each found exactly one
  genuine gap and a re-scan surfaced nothing further.

## What Didn't

- **The two latent test bugs should have been caught earlier** — they predated M32 and were only surfaced because
  M32 happened to run the PyYAML tier. A periodic "run the full suite under every supported interpreter" gate would
  have caught them at their introducing milestone. (Carried as an observation, not a milestone-blocking item.)
- **`junit_tally.py` couldn't run under the box's default `python3` (3.14)** — its `pyexpat` is broken
  (`Symbol not found: _XML_SetAllocTrackerActivationThreshold`). Worked around by parsing the JUnit XML with the
  same `python3.11` that ran the suite. Not M32's to fix; noted for the toolchain.

## Carried Forward

- **None M32-originated.** Every In-list item landed Fate-1; the ant-academy demo-liveness boundary was routed to
  M33/roadmap-vision at the v1.7 design pass (not deferred at M32).
- **Release-level (for `/developer-kit:close-release`, not an M32 deferral):** push the unpushed ext tags
  (`prop-room-m21..m25`, `stage-door-m27`/`m28`/`m30`, `house-lights-m31`/`m32`) to `origin`; the orphaned
  `m26/self-contained-demo` + `wip/clerkenstein-browser-login` ext branches still await their own design-roadmap
  home. Tracked in state.md.

## Metrics Delta

(Sourced from `metrics.json`.)

- **Python tests:** stack-injection `test_injection.py` **87 → 88** test functions (+1 regression:
  `test_studio_desk_env_pins_node_env_production`). Full suite **88/88** PASS (0 skipped under PyYAML, authoritative
  JUnit tally; 8 env-skipped without PyYAML). The harden CORS exact-set + the 2 latent fixes were assertions on
  existing tests, not new test functions.
- **Go test funcs:** **1027** unchanged (M32 touched zero `.go` files).
- **Flake count:** **0** (5/5 randomized-sequential, `PYTHONHASHSEED`-varied — catches dict/order deps in the
  dict-keyed env generator).
- **Coverage:** mutation evidence (no per-file instrument in this behavior-pinned `unittest` layer) — the regression
  test is mutation-checked 4 ways; the CORS exact-set is mutation-confirmed; the additive-merge defense is live-probed.
- **Review findings:** 4, all Fate-1 (3 decision-tag blends + 1 adversarial record). **Field bugs:** 0. **Deferral
  audit:** GREEN (0 in-milestone punts, 0 repeat, 0 aged-out, 0 blockers).
- **ext:** tag `house-lights-m32` @ `107599c` (build tip); harden `7b17c39`.

_M32 is the final v1.7 milestone — v1.7 "house lights" is now ready for `/developer-kit:close-release`._
