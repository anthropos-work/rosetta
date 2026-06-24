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

**Active milestone:** **M40 Directus serve-grant — BUILT (section), unclosed.** Implemented on
`m40/directus-serve-grant`: the per-stack Directus serves the library + activity feed anonymously via cms (live
acceptance demo-3: publicSkillPaths=22, publicJobSimulations=50, jobSimulation+sequences all >0) — the root cause
was a bigger-than-framed `directus_relations=0/directus_fields=0` gap, fixed in tooling (relations + fields +
synthesized read grants + directus_versions read+create + the library/resource/job_position closure). BOTH the
library AND activity-feed halves shipped in tooling (zero platform edits — the key-risk fork refuted). rext
code-of-record @ tag `method-acting-m40`. Next: `/developer-kit:close-milestone` M40 (or harden first).
**Next up:** **`/developer-kit:close-milestone`** M40 — then M41 (profile depth; shares `stack-seeding`, depends
on M39's `users.go`, which it has). Recommended order: M40 (closing) → M41 → M42e → M42m.
**Last closed:** **M39 profile identity (section) — 2026-06-24** via `/developer-kit:close-milestone` (GREEN, 0
blocking; 5 findings all decision-triage; deferral re-audit GREEN 0 escape-hatch; 3 offline alignment gates
100%/100%; flake 0). Merged into `release/01.10-method-acting`; rext code-of-record @ tag `method-acting-m39`.
**M39 LIVE ACCEPTANCE PASSED** (fresh `/demo-up 3`, logged in as Maya): G1 roster `org_name="Cervato Systems"`
threaded + served by the rebuilt fake-fapi; G2 `user_basic_info` role = "Backend Developer" (+ summary/location);
G4 offline real-face SVG avatar. **Acceptance surfaced + fixed a tooling defect** — `up-injected.sh` built the
shared `stackseed` CLI *build-if-missing*, and `$STACK/bin` survives `down --purge`, so a re-consume of new
seeder tooling silently reused the STALE binary (the acceptance no-op'd until the bin was removed by hand). Fixed
to ALWAYS rebuild stackseed: **rext tag `method-acting-m39-fix1`** (+ regression pin). Matters for every
remaining v1.10 milestone acceptance.
**Phase:** **v1.10 in development — M39 closed; M40 built (unclosed); M41–M42m planned, unstarted.**
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

## Headline numbers (v1.9 final close values — the v1.10 inheritance baseline, 2026-06-23)
- **Go test funcs:** **1248** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **259** (250 `Test` + 9 `Fuzz`) · stack-seeding **444** · stack-snapshot 333 · stack-secrets 160.
  Verified at close: stack-seeding `go test ./...` 8 pkgs `ok`, clerkenstein 14 pkgs `ok`; `go vet`+`gofmt`
  clean. Baseline v1.8 1027 → **+221**.
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

_Last updated: 2026-06-24 (**v1.10 "method acting" DESIGNED + PROMOTED** via `/developer-kit:design-roadmap` —
the believable-profile release; 5 milestones M39→M42m [3 `section` targeted fills: profile identity / Directus
serve-grant for library+activity-feed / profile depth — + 2 `iterative` per-vantage Playwright coverage gates,
M42e employee + M42m manager], grounded by the live-demo review `.agentspace/profile_gaps.md` + workflow
`w7t4wq2z4`; branch `release/01.10-method-acting` cut, milestone dirs scaffolded. **Next: `/developer-kit:build-milestone`
M40 ∥ M39.** Prior same day: **Post-v1.9 demo-hardening pass extended to `storytelling-postfix-2`** — the
ant-academy clone/token correction: `npm install` needs NO Font Awesome token [FA Pro icons are
self-hosted/vendored; `FONTAWESOME_NPM_AUTH_TOKEN` is vestigial], and the real "ant-academy down in the demo"
cause — an empty `stack-demo/ant-academy/` stub blocking `make init` — is fixed by `ensure-clones.sh`'s
incomplete-stub sweep + `ant-academy.sh` auto-`npm install`; proven live on `:33077`. Hardening pass now spans
`storytelling-postfix-1` [the four fixes] + `storytelling-postfix-2`.) Prior: 2026-06-23 — **v1.9
"storytelling" SHIPPED** via `/developer-kit:close-release` — reviewed M34–M38 as one PR, **GREEN/0
blocking**; release-level docs coherence 0 findings; deferral re-audit GREEN [0 open, 0 escape-hatch];
supply-chain GREEN [0 new deps]; all 5 Clerkenstein alignment gates 100%/100%; Go 1027→1248. Merged
`release/01.90-storytelling` → `main`, tagged `v1.9`, branch deleted — origin push pending [orchestrator's
step]. **Next: `/developer-kit:design-roadmap`.** v1.8 "understudy" SHIPPED 2026-06-15. **Post-v1.9 (between
releases): demo-hardening pass SHIPPED** (`rosetta-extensions` @ `storytelling-postfix-1` + `-postfix-2`) —
`DEMO_STORIES` default-on + `DEMO_NO_STORIES` opt-out; M33 ant-academy/cockpit session-detach fix; Directus
boot health-gate; guarded prod-Directus note; ant-academy clone/token correction. Tooling + docs only, zero
platform-repo edits, 5 alignment gates still 100%. M33 dropped from the backlog.)_
