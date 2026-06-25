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

**Active milestone:** **_(between section milestones — M42e is next, the first `iterative` milestone)_.** M41
profile depth closed 2026-06-25. The two remaining v1.10 milestones — **M42e** (employee 100% demo coverage) +
**M42m** (manager 100% demo coverage) — are `iterative` (per-vantage Playwright coverage gates); drive them with
`/developer-kit:work-mstone-iters` / `/developer-kit:build-mstone-iters`, NOT `build-milestone`.
**Next up:** **M42e** (employee 100% coverage, `iterative`; depends on M39+M40+M41 fills, all now landed; also
delivers the Playwright coverage harness M42m reuses) → **M42m** (manager 100% coverage, `iterative`).
**Last closed:** **M41 profile depth (section) — 2026-06-25** via `/developer-kit:close-milestone` (GREEN, 0
blocking; 8 findings — 1 code nice-to-have [kept], 1 Phase-2c adversarial [empty-`eduIDs` modulo guard, code
already correct], 5 docs [completeness, incl. README test-count 406→496], 1 decision-triage; deferral re-audit
GREEN, 0 escape-hatch; flake 0; supply-chain GREEN). The `ProfileSeeder` (work-history + education timeline + the
claimed-but-unverified tail) ships both G3+G5 in tooling, zero platform edits; stack-seeding 462→496 (+34), both
M41 files 100% per-function. Merged into `release/01.10-method-acting`; rext code-of-record @ tag
`method-acting-m41` → `0346113` (the tag moved for the close AR-1 test + README refresh). **Reminder for the M41
demo acceptance:** login-as-Maya on a fresh `/demo-up` should show a populated `/profile` timeline + a wide
claimed-vs-verified gap.
**Phase:** **v1.10 in development — M39 + M40 + M41 (all `section`) closed; M42e/M42m (`iterative`) next,
unstarted. Release NOT merged to main (M41 is not the final milestone).**
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

## Headline numbers (v1.10 in-progress — through M41 close, 2026-06-25)
- **Go test funcs:** **1326** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **264** · stack-seeding **496** · stack-snapshot **354** · stack-secrets 160.
  v1.10 deltas: M39 clerkenstein +5 / stack-seeding +18; M40 stack-snapshot +21 (directus pkg);
  M41 stack-seeding +34 (the ProfileSeeder — profile.go/profile_write.go + tests, both files 100% per-function).
  Verified at M41 close: stack-seeding `go test -race ./...` 8 pkgs `ok`; `go vet`+`gofmt` clean. v1.9-close
  baseline 1248 → **+78**.
- **Python tests:** **283** across the two v1.9-touched suites (demo-stack/tests **166**, stack-injection/tests
  **117** [8 opt-in skipped]). Both green; no suite decreased (untouched rext suites carry forward at v1.8
  counts).
- **Flake:** **0** — triple-clean release gate 3/3 (stack-seeding incl `-race` + clerkenstein, shuffled; Python
  re-verified).
- **Supply-chain:** **GREEN** — `go.mod`/`go.sum` diff `storytelling-m34~1..m38` EMPTY (byte-identical to v1.8);
  0 new third-party deps; all deps MIT/BSD/Apache.
- **Alignment gates (green since v1.0):** **100%/100%** on **all 5** Clerkenstein surfaces — M37 added the
  multi-identity `clerk-multi-1` (9 genes) and held the 4 existing (Go 22/22, JS 9/9, deploy 7/7) green (the
  `clerk-express-1` node-CI gate is an env prereq, not a regression).

## Branch model
**v1.10 IN DEVELOPMENT:** `release/01.10-method-acting` cut from `main` 2026-06-24 (the design-roadmap run lands the
v1.10 plan + milestone scaffolds here). Milestone branches `m{39..42m}/{slug}` branch from it; rext code-of-record
lands in the `rosetta-extensions` authoring copy (a SEPARATE repo) at new v1.10 tags as milestones close.
**v1.9 SHIPPED:** `release/01.90-storytelling` merged `--no-ff` → `main` + tagged `v1.9`; pushed to origin
2026-06-24. The release branch is deleted. v1.9 rext code-of-record at tags `storytelling-m34..m38`.
**Shipped:** **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` ·
**v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-25 (**M41 profile depth CLOSED** via `/developer-kit:close-milestone` — GREEN, 0
blocking, 8 findings [1 code nice-to-have kept, 1 Phase-2c adversarial empty-`eduIDs` modulo guard, 5 docs incl.
README test-count 406→496, 1 decision-triage]; the `ProfileSeeder` ships G3 work-history/education timeline + G5
verified-bump + the claimed-but-unverified tail, both halves in tooling, zero platform edits; stack-seeding
462→496 [+34], both M41 files 100% per-function, flake 0, supply-chain GREEN; deferral re-audit GREEN [M41 added
0; inherited KPI=0 Fate-2→M42e/M42m, not aged-out]. Merged into `release/01.10-method-acting`; rext
code-of-record @ tag `method-acting-m41` → `0346113`. **M41 is the LAST `section` milestone — M42e/M42m
(`iterative`) remain; release NOT merged to main. Next: M42e** (employee 100% Playwright coverage). Full M41
narrative in roadmap.md § M41. Prior: **M40 Directus serve-grant CLOSED 2026-06-24** [stack-snapshot 333→354,
directus pkg 100%, both library+activity-feed halves in tooling]. Prior: **M39 profile-identity CLOSED
2026-06-24** + **v1.10 "method acting" DESIGNED + PROMOTED** [5 milestones M39→M42m]. Prior: **v1.9
"storytelling" SHIPPED 2026-06-23** [tag `v1.9`, M34–M38, Go 1027→1248]; v1.8 "understudy" SHIPPED 2026-06-15.
**Post-v1.9 demo-hardening pass SHIPPED** [`rosetta-extensions` @ `storytelling-postfix-1` + `-postfix-2`:
`DEMO_STORIES` default-on, M33 session-detach fix, Directus health-gate, ant-academy clone/token correction;
tooling + docs only, 5 alignment gates still 100%].)_
