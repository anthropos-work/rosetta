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

**Active milestone:** **(between milestones — all 5 v1.10 milestones closed)**. v1.10 is reproducibly gate-complete
across BOTH per-vantage coverage gates (employee M42e + manager M42m); the release branch
`release/01.10-method-acting` is NOT merged to main yet.
**Next up:** **`/developer-kit:close-release`** (review the whole v1.10 release, merge `release/01.10-method-acting`
→ `main`, tag `v1.10`) — the user's separate step, after their visual review.
**Last closed:** **M42m manager 100% coverage (iterative) — 2026-06-26** via `/developer-kit:close-milestone`
(GREEN, 0 blocking; closed **on-gate** — manager semantic believability gate `gateMet:true`, fresh zero-manual
demo-up: 70 reachable / 0 failing sections / 0 persona failures / 0 escapes / frontier EXHAUSTED — AND the M42e
employee gate HELD on the same fresh stack, no regression: 59 reachable / 0,0,0,0 / EXHAUSTED). 5 iters (1 bootstrap
tok + 4 tiks; 0 triggered toks). Delivered (rext, tagged): the **`demopatch` tool** (the sanctioned mechanism for the
platform-bound Studio left-nav escape — patch the demo's EPHEMERAL clone before build + trap-revert, 6 guards;
resolved **demo-only 139→0**), the **FeedbackSeeder org-feedback JOIN-mirror** (the `/enterprise/organization-feedback`
"No data" fix), and the **manager harness namespace** (`MANAGER_PAGES` reconciled to the real `/enterprise/*` route
model + `MANAGER_SAMPLE_RULES` superset, `calibrated:true`). RESCOPE-1 resolved demo-only via the demopatch tool —
**not a platform edit, not a deferral**. 3 findings, 0 blocking (corpus diff docs-only — code-of-record in rext; 2
Fate-1 docs + 1 triage-archive); deferral re-audit GREEN (0 escape-hatch; **DEF-M40-01 manager-half resolved
in-milestone** — route reconcile turned notReached=5 into 6 asserted dashboard pages rendering real M36 data);
**zero CANONICAL platform edits**; flake 0; supply-chain GREEN (0 dep/lockfile change in the whole M42m footprint);
alignment N/A (zero clerkenstein change) — 5 gates carry forward 100%. Merged into `release/01.10-method-acting`;
rext code-of-record @ tag `method-acting-m42m-harden-final`. **M42m demo acceptance reminder:** login-as-Dan-Rossi
(`dan-manager`) on a fresh `/demo-up` shows a believable fully-populated manager experience — the M36
Workforce-Intelligence dashboard (493 mapped / 262 verified / 53.1% coverage, 19 cards / 67 charts), the
`/enterprise/*` admin pages with real org/team data, the Studio left-nav resolves demo-local (`:39000`, no prod
eject).
**Phase:** **v1.10 fully built — all 5 milestones (M39/M40/M41 `section` + M42e/M42m `iterative`) closed; release
NOT merged to main. Next: `/developer-kit:close-release` (v1.10 → main + tag `v1.10`).**
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

## Headline numbers (v1.10 fully built — through M42m close, 2026-06-26)
- **Go test funcs:** **1376** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **266** · stack-seeding **537** · stack-snapshot **361** · stack-secrets 160.
  v1.10 deltas: M39 clerkenstein +5 / stack-seeding +18; M40 stack-snapshot +21 (directus pkg);
  M41 stack-seeding +34 (ProfileSeeder); M42e stack-seeding +38 (curated_pools/hero_activity/orglogo/
  photo_avatar/identity-casbin + harden) / stack-snapshot +7 (simembeddings pkg + directus categories) /
  clerkenstein +2 (image-threading invariant); **M42m** stack-seeding +3 (the FeedbackSeeder org-feedback
  mirror tests). Verified at M42m close: `go test -race ./seeders/` GREEN; `go vet`+`gofmt` clean; flake gate
  3× clean. v1.9-close baseline 1248 → **+128**.
  **PLUS the TypeScript Playwright harness:** the manager namespace gained a pure-logic unit spec
  (`coverage-manifest.unit.spec.ts`, +17) — the manifest decision logic now pinned in CI with no stack — on top
  of the M42e live-sweep specs (rext `stack-verify/e2e/`, @playwright/test ^1.49.0). **PLUS the Python demopatch
  suite:** `test_demopatch.py` 18→**43** (+25 adversarial-guard + manifest-loader parser tests) — the demopatch
  tool patches platform source, so its REFUSE grid is the highest-value harden target.
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
**v1.10 FULLY BUILT (not yet merged to main):** `release/01.10-method-acting` cut from `main` 2026-06-24. Milestone
branches `m{39..42m}/{slug}` branched from it (all 5 merged + deleted); rext code-of-record lands in the
`rosetta-extensions` authoring copy (a SEPARATE repo) at v1.10 tags as milestones close. Closed-milestone rext tags:
`method-acting-m39` · `m40` · `m41` · `m42e` → `53574ae` · **`m42m-harden-final`**. All 5 milestones closed; the
release branch is NOT merged to main yet — `/developer-kit:close-release` follows (the user's separate step).
**v1.9 SHIPPED:** `release/01.90-storytelling` merged `--no-ff` → `main` + tagged `v1.9`; pushed to origin
2026-06-24. The release branch is deleted. v1.9 rext code-of-record at tags `storytelling-m34..m38`.
**Shipped:** **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` ·
**v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-26 (**M42m manager 100% coverage CLOSED** via `/developer-kit:close-milestone` —
GREEN, 0 blocking; closed **on-gate** [manager semantic believability gate `gateMet:true`, fresh zero-manual
demo-up: 70 reachable / 0 failing sections / 0 persona failures / 0 escapes / frontier EXHAUSTED — AND the M42e
employee gate HELD, no regression: 59 reachable / 0,0,0,0 / EXHAUSTED]; 5 iters [1 bootstrap tok + 4 tiks];
delivered the `demopatch` tool [the platform-bound Studio escape resolved demo-only 139→0, 6 guards] + the
FeedbackSeeder org-feedback JOIN-mirror + the manager harness namespace [`/enterprise/*` route reconcile +
sample-rules]; 3 findings [docs-only corpus diff: 2 Fate-1 docs + 1 triage-archive]; deferral re-audit GREEN
[0 escape-hatch; DEF-M40-01 manager-half resolved in-milestone; RESCOPE-1 demo-only, not a deferral]; zero
CANONICAL platform edits; Go 1373→1376 [+3] + the demopatch suite 18→43 + the TS manager unit spec +17; flake 0;
supply-chain GREEN; 5 alignment gates 100% [N/A change]. Merged into `release/01.10-method-acting`; rext
code-of-record @ tag `method-acting-m42m-harden-final`. **M42m is the LAST v1.10 milestone — all 5 closed; v1.10 is
reproducibly gate-complete across BOTH per-vantage gates [employee M42e + manager M42m]. Release NOT merged to main.
Next: `/developer-kit:close-release`** [v1.10 → main + tag `v1.10`]. Full M42m narrative in roadmap.md § M42m.
Prior: **M42e employee 100% coverage CLOSED 2026-06-25** [first iterative milestone, 23 iters, Go 1326→1373,
Playwright harness]; **M41 profile depth CLOSED 2026-06-25** [ProfileSeeder G3+G5, stack-seeding 462→496];
**M40 Directus serve-grant CLOSED 2026-06-24** [stack-snapshot 333→354]; **M39 profile-identity CLOSED
2026-06-24** + **v1.10 DESIGNED**. Prior: **v1.9 "storytelling" SHIPPED 2026-06-23** [tag `v1.9`, M34–M38]; v1.8
"understudy" SHIPPED 2026-06-15.
**Post-v1.9 demo-hardening pass SHIPPED** [`rosetta-extensions` @ `storytelling-postfix-1` + `-postfix-2`:
`DEMO_STORIES` default-on, M33 session-detach fix, Directus health-gate, ant-academy clone/token correction;
tooling + docs only, 5 alignment gates still 100%].)_
