# State

**Active version:** **v1.10 "method acting" — IN DEVELOPMENT (EXTENDED with M43→M46)** (designed 2026-06-24,
**extended 2026-06-26**, via `/developer-kit:design-roadmap`; branch `release/01.10-method-acting`). The
**believable-profile release** (M39–M42m, all CLOSED) **+ the presenter-grade / scalable-generation extension**
(M43–M46, planned). The built half: when a presenter clicks **Login as** a hero, the individual's **profile** (org
name, role+title, work history, education, a real face, deep role-aligned skills) **and** the content surfaces
(**library** + the **activity feed**) populate with real content, on **every** page a hero of that vantage can
reach — proven by a **Playwright** coverage sweep with **zero** empty pages and **zero** out-of-demo escapes (5
milestones M39→M42m). The extension: M43 **cockpit UX polish** (a slick light presenter launcher — one CTA, icons,
manifest download, login-progress overlay), M44 **profile completeness** (the whole roster — members **and**
managers — fully baked; DATA DENSITY ONLY, no UI widget), M45 a cheap-LLM **generation engine** (`cmd/gen-batch` +
a prompt-keyed cache + the CODE-owns-structure / AI-owns-content boundary), M46 **org-scale fill** + a gen-batch
preview CLI. Execution: **{ M43 ∥ M44 } → M45 → M46**, then close-release. **Tooling + docs only — zero
platform-repo edits.** Grounded by the live-demo review
[`.agentspace/profile_gaps.md`](../../.agentspace/profile_gaps.md) + workflow `w7t4wq2z4`, and the v1.10-extend
research note [`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../.agentspace/scratch/roadmap-research-2026-06-26.md).

**Active milestone:** **the M43 ∥ M44 opening pair** (both `section`, fully parallel — different rext modules).
**Recommend starting M44** (it gates M45) — or **M43 ∥ M44 in parallel** (a clean `/developer-kit:build-milestone`
section-pair).
**Next up:** build **M43 + M44** (`/developer-kit:build-milestone` — the section pair, parallel), then **M45**
(`/developer-kit:build-mstone-iters` — iterative LLM-generation gate) → **M46** (`/developer-kit:build-mstone-iters`
— iterative org-scale gate) → **THEN `/developer-kit:close-release`** (merge `release/01.10-method-acting` → `main`,
tag `v1.10`).
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
rext code-of-record @ tag `method-acting-m42m-harden-final`.
**Phase:** **v1.10 reopened / extended — 5 milestones CLOSED (M39/M40/M41 `section` + M42e/M42m `iterative`), 4 new
PLANNED (M43/M44 `section` + M45/M46 `iterative`). Release NOT merged to main. Next: build M43∥M44 → M45 → M46 →
`/developer-kit:close-release` (v1.10 → main + tag `v1.10`).**
**Paused:** _(none)_

**Standing backlog (unscheduled, orthogonal to v1.10):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes),
DEF-M21-01 (replayCmd hermetic test), M25-D9 (dev taxonomy rc=4). None in v1.10 scope.

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

## Headline numbers (v1.10 built half — through M42m close, 2026-06-26; M43→M46 not started)
- **Go test funcs:** **1376** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **266** · stack-seeding **537** · stack-snapshot **361** · stack-secrets 160.
  v1.10 deltas: M39 clerkenstein +5 / stack-seeding +18; M40 stack-snapshot +21 (directus pkg);
  M41 stack-seeding +34 (ProfileSeeder); M42e stack-seeding +38 (curated_pools/hero_activity/orglogo/
  photo_avatar/identity-casbin + harden) / stack-snapshot +7 (simembeddings pkg + directus categories) /
  clerkenstein +2 (image-threading invariant); **M42m** stack-seeding +3 (the FeedbackSeeder org-feedback
  mirror tests). Verified at M42m close: `go test -race ./seeders/` GREEN; `go vet`+`gofmt` clean; flake gate
  3× clean. v1.9-close baseline 1248 → **+128**. (M43–M46 will add: M44 the CertificatesSeeder/ProjectsSeeder +
  manager-unskip + bulk-member tests; M45 a first new third-party AI dep [breaks the 0-new-deps streak — a
  deliberate in-release decision] + the gen-batch/cache/GeneratedBatchSeeder suites; M46 the org-scale + preview
  tests.)
  **PLUS the TypeScript Playwright harness:** the manager namespace gained a pure-logic unit spec
  (`coverage-manifest.unit.spec.ts`, +17) — the manifest decision logic now pinned in CI with no stack — on top
  of the M42e live-sweep specs (rext `stack-verify/e2e/`, @playwright/test ^1.49.0; M46 reuses it for the
  population-believability gate). **PLUS the Python demopatch suite:** `test_demopatch.py` 18→**43** (+25
  adversarial-guard + manifest-loader parser tests).
- **Python tests:** **283** across the two v1.9-touched suites (demo-stack/tests **166**, stack-injection/tests
  **117** [8 opt-in skipped]). Both green; no suite decreased (untouched rext suites carry forward at v1.8
  counts). (M43 touches `demo-stack/cockpit.py` — Python; expect a small delta.)
- **Flake:** **0** — triple-clean release gate 3/3 (stack-seeding incl `-race` + clerkenstein, shuffled; Python
  re-verified).
- **Supply-chain:** **GREEN** through M42m (`go.mod`/`go.sum` byte-identical to v1.8; 0 new deps; all deps
  MIT/BSD/Apache). **M45 deliberately adds the first new third-party AI dep** (the `services/ai/` wrapper over the
  shared `ai` library) — an in-release decision recorded on the M45 block; supply-chain will be re-reviewed at
  M45/close.
- **Alignment gates (green since v1.0):** **100%/100%** on **all 5** Clerkenstein surfaces (Go 22/22, JS 9/9,
  multi 9/9, deploy 7/7, express 13/13) + drift 9/9 — re-verified at M42e close. M43 touches the cockpit
  (`demo-stack`), not the Clerkenstein contract surfaces — gates carry forward.

## Branch model
**v1.10 IN DEVELOPMENT (EXTENDED, not yet merged to main):** `release/01.10-method-acting` cut from `main`
2026-06-24. Milestone branches `m{39..42m}/{slug}` branched from it (all 5 merged + deleted); rext code-of-record
lands in the `rosetta-extensions` authoring copy (a SEPARATE repo) at v1.10 tags as milestones close. Closed-milestone
rext tags: `method-acting-m39` · `m40` · `m41` · `m42e` → `53574ae` · **`m42m-harden-final`**. M43→M46 branches
(`m43/cockpit-ux`, `m44/profile-completeness`, `m45/generation-engine`, `m46/org-scale-fill`) are cut by
`/developer-kit:build-milestone` / `build-mstone-iters` as work begins. The release branch is NOT merged to main yet
— `/developer-kit:close-release` follows AFTER M46.
**v1.9 SHIPPED:** `release/01.90-storytelling` merged `--no-ff` → `main` + tagged `v1.9`; pushed to origin
2026-06-24. The release branch is deleted. v1.9 rext code-of-record at tags `storytelling-m34..m38`.
**Shipped:** **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` ·
**v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-26 (**v1.10 EXTENDED with M43→M46** via `/developer-kit:design-roadmap` — the
presenter-grade / scalable-generation extension. Instead of close-release after M42m, v1.10 reopens IN DEVELOPMENT
and adds 4 planned milestones: M43 **cockpit UX polish** [`section`, the light presenter launcher + one CTA +
manifest download + login-progress overlay], M44 **profile completeness** [`section`, whole-roster members+managers,
DATA DENSITY only — no UI widget], M45 **generation engine** [`iterative`, cmd/gen-batch + cheap-LLM + prompt-keyed
cache + the CODE-owns-structure/AI-owns-content closure boundary; the first new third-party dep], M46 **org-scale
fill + gen-batch preview** [`iterative`, supporting-population auto-fill + dry-run preview]. Execution **{ M43 ∥ M44 }
→ M45 → M46**, close-release AFTER M46. Flat sequential numbering M43–M46. Active milestone: the M43∥M44 opening pair
[start M44 — it gates M45 — or both in parallel]. Tooling + docs only — zero platform-repo edits. Research note:
`.agentspace/scratch/roadmap-research-2026-06-26.md`. Full M43–M46 blocks in roadmap.md § Extension — M43→M46.
Prior: **M42m manager 100% coverage CLOSED 2026-06-26** [iterative, closed on-gate, 5 iters; delivered the
`demopatch` tool + the FeedbackSeeder mirror + the manager harness namespace]; **M42e employee 100% coverage CLOSED
2026-06-25** [first iterative milestone, 23 iters, Playwright harness]; **M41 profile depth CLOSED 2026-06-25**;
**M40 Directus serve-grant CLOSED 2026-06-24**; **M39 profile-identity CLOSED 2026-06-24** + **v1.10 DESIGNED**.
Prior: **v1.9 "storytelling" SHIPPED 2026-06-23** [tag `v1.9`, M34–M38].
**Post-v1.9 demo-hardening pass SHIPPED** [`rosetta-extensions` @ `storytelling-postfix-1` + `-postfix-2`:
`DEMO_STORIES` default-on, M33 session-detach fix, Directus health-gate, ant-academy clone/token correction;
tooling + docs only, 5 alignment gates still 100%].)_
