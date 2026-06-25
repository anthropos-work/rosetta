# State

**Active version:** **v1.10 "method acting" — IN DEVELOPMENT** (designed 2026-06-24 via
`/developer-kit:design-roadmap`; branch `release/01.10-method-acting`). The **believable-profile release** — v1.9
told the *story*; v1.10 makes each *character* hold up under a close-up: when a presenter clicks **Login as** a
hero, the individual's **profile** (org name, role+title, work history, education, a real face, deep role-aligned
skills) **and** the content surfaces (**library** + the **activity feed**) populate with real content, on **every**
page a hero of that vantage can reach — proven by a **Playwright** coverage sweep (DOM + screenshots) with **zero**
empty pages and **zero** out-of-demo escapes. 5 milestones **M39→M42m**: { M39 profile identity ∥ M40 Directus
serve-grant (library + activity feed) } → M41 profile depth → **M42e** employee 100% coverage → **M42m** manager
100% coverage (M42e/M42m `iterative`; the rest `section`). **Tooling + docs only — zero platform-repo edits.**
Grounded by the live-demo review [`.agentspace/profile_gaps.md`](../../.agentspace/profile_gaps.md) + the
root-cause workflow `w7t4wq2z4`.

**Active milestone:** **M42m** (manager 100% demo coverage, `iterative`) — the LAST v1.10 milestone. It REUSES the
M42e Playwright harness + `coverage-protocol.md`; its input is already calibrated (the M42e manager smoke-sweep:
139 `studio.anthropos.work` escapes from one baked left-nav prod link → link-rewriting; 5 unreached `/workforce/*`
M36 dashboard pages → the core content+nav work; the team-roster `/user/<id>` fan-out → a representative-sample
crawl rule). Drive with `/developer-kit:work-mstone-iters` / `/developer-kit:build-mstone-iters`, NOT
`build-milestone`.
**Next up:** **M42m** → then `/developer-kit:close-release` (v1.10 → main + tag `v1.10`).
**Last closed:** **M42e employee 100% coverage (iterative) — 2026-06-25** via `/developer-kit:close-milestone`
(GREEN, 0 blocking; closed **on-gate** — employee semantic believability gate `gateMet:true`, fresh zero-manual
demo-up, 62 reachable pages / 0 failing sections / 0 persona failures / 0 escapes / frontier EXHAUSTED). 23 iters;
delivered `coverage-protocol.md` + the Playwright semantic-coverage harness (rext `stack-verify/e2e/` — first non-Go
rext dev/test dep). 8 findings, 0 blocking (0 code must-fix; 3 Phase-2c adversarial recorded; 4 docs); deferral
re-audit GREEN (0 escape-hatch; DEF-M40-01 employee-half resolved, manager-half + manager-sweep findings clean
Fate-2/Fate-3 → M42m); zero platform edits; flake 0; supply-chain GREEN; all 5 alignment gates 100%. Merged into
`release/01.10-method-acting`; rext code-of-record @ tag `method-acting-m42e` → `53574ae`. **M42e demo acceptance
reminder:** login-as-Maya on a fresh `/demo-up` shows a believable fully-populated person + catalog on every
employee page (real-photo avatar menu==profile, org logo, DevOps-coherent skills, 274-sim library, no prod escapes).
**Phase:** **v1.10 in development — M39/M40/M41 (`section`) + M42e (`iterative`) closed; M42m (`iterative`) is the
last milestone. Release NOT merged to main (M42m remains; close-release follows).**
**Paused:** _(none)_

**Standing backlog (unscheduled, orthogonal to v1.10):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes),
DEF-M21-01 (replayCmd hermetic test), M25-D9 (dev taxonomy rc=4). None in v1.10 scope.

**Post-v1.9 tooling-hardening pass (between releases) — SHIPPED.** A demo-hardening effort landed on the
demo-stack / dev-setdress tooling across two `rosetta-extensions` tags: **`storytelling-postfix-1`** (the four
fixes) + **`storytelling-postfix-2`** (the ant-academy clone/token correction). At `storytelling-postfix-1`:
**`DEMO_STORIES` is now default-on** (a bare `/demo-up N` seeds the multi-org Stories & Heroes world + serves
the presenter cockpit; `DEMO_NO_STORIES=1` opts back out to the legacy structural small-200 + single-identity
fake-fapi + no-cockpit demo); the **M33** ant-academy + cockpit "dead on a later visit" bug is **resolved**
(both host-native daemons now launch session-detached via `demo-stack/detach.sh::launch_detached` instead of
bare `nohup`, so they survive the launching session/task ending); the per-stack **Directus boot now
health-gates** on the stack's own offset `/server/health` before returning so the bring-up-tail autoverify
can't race its ~30s re-introspect; and the **prod-Directus content note is guarded** (printed only on
`DEMO_NO_LOCAL_CONTENT=1`, since the default demo boots a per-stack Directus serving the captured catalog
locally). At **`storytelling-postfix-2`** the "ant-academy down in the demo" cause was corrected: it was
**not** a missing Font Awesome Pro token (ant-academy USES FA Pro icons, but the assets are
self-hosted/vendored in the repo — `code/public/assets/fontawesome/` — so `npm install` and running it need
**no token**; `FONTAWESOME_NPM_AUTH_TOKEN` in `.env.example` is vestigial). The real cause was a blocked
clone: an empty `stack-demo/ant-academy/` stub (holding only a gitignored `code/.env.local`) defeated `make
init`'s skip-if-present, so the source never landed. The fix: **`ensure-clones.sh` now sweeps incomplete
sibling stubs** (any `repos.yml` repo dir with no `.git`) before `make init`, and **`ant-academy.sh`
auto-runs `npm install`** (no token) when `node_modules` is absent — so a fresh `/demo-up` now brings
ant-academy up automatically (proven live on `:33077`). **Tooling + docs only — zero platform-repo edits; all
5 Clerkenstein alignment gates still 100%/100%.**

## Recently shipped releases
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. Believable-demo-narrative release: the placeholder
  seeder becomes a declarative **Stories & Heroes** engine (per-story org + a hero trio via the real
  verified-skill chain) so the skill profile + the Workforce dashboard tell a story, plus a presenter cockpit
  on Clerkenstein multi-identity. 5 `section` milestones **M34→M38**. Headline: zero platform-repo edits; all
  **5** Clerkenstein alignment gates 100%/100%; supply-chain GREEN (0 new deps); Go 1027→**1248** (stack-seeding
  259→444). Code: `rosetta-extensions` @ tags `storytelling-m34..m38`. Records:
  [releases/archive/01.90-storytelling/](releases/archive/01.90-storytelling/).
- **v1.8 "understudy"** — **2026-06-15**, tag `v1.8`. Self-contained-demo release: a demo builds **entirely from
  `stack-demo`'s own clone set** (a box with only `stack-demo/` runs a demo end-to-end). Single `section`
  milestone **M26**. Code: `rosetta-extensions` @ tag `understudy-m26`. Records:
  [releases/archive/01.80-understudy/](releases/archive/01.80-understudy/).
- **v1.7 "house lights"** — **2026-06-15**, tag `v1.7`. Demo-UI-hardening: M31 mkcert FAPI cert (next-web stops
  blanking) + M32 studio-desk single-port/production fix. Ext tags `house-lights-m31`/`m32`.

## Headline numbers (v1.10 in-progress — through M42e close, 2026-06-25)
- **Go test funcs:** **1373** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **266** · stack-seeding **534** · stack-snapshot **361** · stack-secrets 160.
  v1.10 deltas: M39 clerkenstein +5 / stack-seeding +18; M40 stack-snapshot +21 (directus pkg);
  M41 stack-seeding +34 (ProfileSeeder); **M42e** stack-seeding +38 (curated_pools/hero_activity/orglogo/
  photo_avatar/identity-casbin + harden) / stack-snapshot +7 (simembeddings pkg + directus categories) /
  clerkenstein +2 (image-threading invariant). Verified at M42e close: `go test -race ./...` GREEN on the 3
  touched modules; `go vet`+`gofmt` clean; flake gate 5/5 shuffled. v1.9-close baseline 1248 → **+125**.
  **PLUS the NEW TypeScript Playwright harness:** 7 tests / 6 spec files (rext `stack-verify/e2e/`) — the
  first non-Go test surface, @playwright/test ^1.49.0.
- **Python tests:** **283** across the two v1.9-touched suites (demo-stack/tests **166**, stack-injection/tests
  **117** [8 opt-in skipped]). Both green; no suite decreased (untouched rext suites carry forward at v1.8
  counts).
- **Flake:** **0** — triple-clean release gate 3/3 (stack-seeding incl `-race` + clerkenstein, shuffled; Python
  re-verified).
- **Supply-chain:** **GREEN** — `go.mod`/`go.sum` diff `storytelling-m34~1..m38` EMPTY (byte-identical to v1.8);
  0 new third-party deps; all deps MIT/BSD/Apache.
- **Alignment gates (green since v1.0):** **100%/100%** on **all 5** Clerkenstein surfaces (Go 22/22, JS 9/9,
  multi 9/9, deploy 7/7, express 13/13) + drift 9/9 — re-verified at M42e close (which touched clerk-frontend
  for the avatar/org-logo `image_url` threading; the image-threading JSON invariant gained a dedicated test).

## Branch model
**v1.10 IN DEVELOPMENT:** `release/01.10-method-acting` cut from `main` 2026-06-24. Milestone branches
`m{39..42m}/{slug}` branch from it; rext code-of-record lands in the `rosetta-extensions` authoring copy (a SEPARATE
repo) at new v1.10 tags as milestones close. Closed-milestone rext tags: `method-acting-m39` · `m40` · `m41` ·
**`m42e`** → `53574ae`. M42m is the last milestone; release NOT merged to main yet (close-release follows M42m).
**v1.9 SHIPPED:** `release/01.90-storytelling` merged `--no-ff` → `main` + tagged `v1.9`; pushed to origin
2026-06-24. The release branch is deleted. v1.9 rext code-of-record at tags `storytelling-m34..m38`.
**Shipped:** **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` ·
**v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-25 (**M42e employee 100% coverage CLOSED** via `/developer-kit:close-milestone` —
GREEN, 0 blocking; closed **on-gate** [employee semantic believability gate `gateMet:true`, fresh zero-manual
demo-up: 62 reachable / 0 failing sections / 0 persona failures / 0 escapes / frontier EXHAUSTED]; 23 iters;
delivered `coverage-protocol.md` + the Playwright semantic-coverage harness [rext `stack-verify/e2e/`, first
non-Go rext dev/test dep]; 8 findings [0 code must-fix, 3 Phase-2c adversarial, 4 docs]; deferral re-audit GREEN
[0 escape-hatch; DEF-M40-01 employee-half resolved, manager-half + manager-sweep findings clean Fate-2/Fate-3 →
M42m]; zero platform edits; Go 1326→1373 [+47] + 7 Playwright tests; seeders 97%; flake 0; supply-chain GREEN; 5
alignment gates 100%. Merged into `release/01.10-method-acting`; rext code-of-record @ tag `method-acting-m42e` →
`53574ae`. **M42e is the FIRST iterative milestone closed — M42m [manager, iterative] is the last; release NOT
merged to main. Next: M42m** [reuses the M42e harness; its input is calibrated]. Full M42e narrative in
roadmap.md § M42e. Prior: **M41 profile depth CLOSED 2026-06-25** [ProfileSeeder G3+G5, stack-seeding 462→496];
**M40 Directus serve-grant CLOSED 2026-06-24** [stack-snapshot 333→354]; **M39 profile-identity CLOSED
2026-06-24** + **v1.10 DESIGNED**. Prior: **v1.9 "storytelling" SHIPPED 2026-06-23** [tag `v1.9`, M34–M38]; v1.8
"understudy" SHIPPED 2026-06-15.
**Post-v1.9 demo-hardening pass SHIPPED** [`rosetta-extensions` @ `storytelling-postfix-1` + `-postfix-2`:
`DEMO_STORIES` default-on, M33 session-detach fix, Directus health-gate, ant-academy clone/token correction;
tooling + docs only, 5 alignment gates still 100%].)_
