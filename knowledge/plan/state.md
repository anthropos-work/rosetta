# State

**Active version:** **v1.10 "method acting" — IN DEVELOPMENT (EXTENDED with M43→M46)** (designed 2026-06-24,
**extended 2026-06-26**, via `/developer-kit:design-roadmap`; branch `release/01.10-method-acting`). The
**believable-profile release** (M39–M42m, all CLOSED) **+ the presenter-grade / scalable-generation extension**
(M43/M44/M45 CLOSED; M46 planned). The built half: when a presenter clicks **Login as** a hero, the individual's **profile** (org
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

**Active milestone:** **(between milestones — see Next Up).** The { M43 ∥ M44 } opening section-pair and **M45
(the generation engine)** are **CLOSED**; the next — and **ONLY remaining** — milestone is **M46 (org-scale fill
+ a gen-batch preview CLI — `iterative`)**, now unblocked (it reuses M45's engine + cache + the closure-safe
`GeneratedBatchSeeder`, and the M42 Playwright coverage harness for its population-believability gate).
**Next up:** **`/developer-kit:work-mstone-iters` for M46** (the **iterative** org-scale-fill gate — the
interactive driver; NOT `work-milestone`) → **THEN `/developer-kit:close-release`** (merge
`release/01.10-method-acting` → `main`, tag `v1.10`). _(Do NOT run close-release yet — M46 remains.)_
**Last closed:** **M45 generation engine (`iterative`, closed-on-gate) — 2026-06-26** via
`/developer-kit:close-milestone` (GREEN, 3 findings / 0 blocking). The cheap-LLM batch profile generator with the
CODE-owns-structure / AI-owns-content boundary, built inside-out over 7 iters (1 tok + 6 tiks). **Gate MET 5/5 on
a REAL Azure gpt-4o-mini run** (EU-first Sweden; the direct key was billing-dead), N=20 + demo-3 proof: valid-JSON
100% (33/33), 47/47 skills + 20/20 roles → real `skiller.*.node_id` (`datadna measure-closure` = `[PASS]`, 0
fabrication), 0 hero-collisions, $0.0059 ≤ $0.10, $0 byte-identical re-seed; 20/20 distinct multicultural names +
avatars, isolation CLEAN. **Supply-chain: the deliberate, user-acknowledged 1-new-dep inflection** —
`anthropos-work/ai@v1.40.1` (all-permissive; the FIRST new dep since v1.8, the milestone is ABOUT it; NOT a
regression). 5-pass final harden stabilized; coverage 88.5–98.9% across the 5 new pkgs; 0 flakes; `-race` clean.
NEW `corpus/ops/demo/ai-generation-spec.md` + `corpus/ops/demo/cache-spec.md` (indexed). Deferral GREEN (org-scale
→ M46 Fate-2 owned); alignment N/A (100%/100%); **zero canonical platform edits**. Merged `--no-ff` into
`release/01.10-method-acting`; rext code-of-record @ tag `method-acting-m45-harden-final`.
**Phase:** **v1.10 reopened / extended — 8 milestones CLOSED (M39/M40/M41 `section` + M42e/M42m `iterative` +
M43/M44 `section` + M45 `iterative`), 1 PLANNED (M46 `iterative`). Release NOT merged to main. Next: M46 →
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

## Headline numbers (v1.10 built half — through M45 close, 2026-06-26; M46 not started)
- **Go test funcs:** **1516** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **266** · stack-seeding **677** · stack-snapshot **361** · stack-secrets 160.
  **M45 delta: stack-seeding 567→677 (+110)** — the 5 new pkgs (services/ai 32 + blueprint 45 + batchcache 21 +
  cmd/gen-batch 23) + the GeneratedBatchSeeder/drop-not-fabricate/boundary-fuzz tests in seeders/, across iter-02..07
  build + the 5-pass final harden. clerkenstein/snapshot/alignment/secrets UNCHANGED (M45 touched stack-seeding only —
  verified empty clerkenstein diff m44..m45). Coverage on the new pkgs: services/ai 97.8% / blueprint 98.9% /
  batchcache 88.5% / cmd/gen-batch 93.0% / seeders 97.3%. Prior milestone deltas (M39–M44) below.
  v1.10 deltas: M39 clerkenstein +5 / stack-seeding +18; M40 stack-snapshot +21 (directus pkg);
  M41 stack-seeding +34 (ProfileSeeder); M42e stack-seeding +38 (curated_pools/hero_activity/orglogo/
  photo_avatar/identity-casbin + harden) / stack-snapshot +7 (simembeddings pkg + directus categories) /
  clerkenstein +2 (image-threading invariant); **M42m** stack-seeding +3 (the FeedbackSeeder org-feedback
  mirror tests); **M44** stack-seeding **+29** (Certificates/ProjectsSeeder + §A/§C/§D build tests + the 17-test
  profile_completeness harden file; seeders stmt cov 96.5→97.5%). Verified at M44 close: `go test -race ./...`
  on stack-seeding GREEN; `go vet`+`gofmt` clean; flake gate 5× clean. v1.9-close baseline 1248 → **+158**.
  (Counting-method note: the M42m close quoted stack-seeding=537; ground-truth grep at the m42m tag is 538 — a
  1-off prior-close drift; the M44 +29 is measured against the authoritative 538.) (M43/M45/M46 will add: M43 a
  small Python `cockpit.py` delta; M45 a first new third-party AI dep [breaks the 0-new-deps streak — a
  deliberate in-release decision] + the gen-batch/cache/GeneratedBatchSeeder suites; M46 the org-scale + preview
  tests.)
  **M43 cockpit Python delta: `demo-stack/cockpit.py` tests 27→63 (+36)** across 2 harden passes (escaping
  depth incl. the `jump_to` href-injection invariant, CTA-unification, FA SRI well-formedness, manifest download,
  overlay JS, served-panel UTF-8 byte-length); Go test funcs unchanged (Python-only change, separate rext module).
  **PLUS the TypeScript Playwright harness:** the manager namespace gained a pure-logic unit spec
  (`coverage-manifest.unit.spec.ts`, +17) — the manifest decision logic now pinned in CI with no stack — on top
  of the M42e live-sweep specs (rext `stack-verify/e2e/`, @playwright/test ^1.49.0; M46 reuses it for the
  population-believability gate). **PLUS the Python demopatch suite:** `test_demopatch.py` 18→**43** (+25
  adversarial-guard + manifest-loader parser tests).
- **Python tests:** the `demo-stack/tests` suite grew with **M43** (cockpit Python tests 27→**63**, +36 across the
  2 harden passes — the suite that owns `test_cockpit.py`); `stack-injection/tests` **117** [8 opt-in skipped]
  unchanged. All green; no suite decreased (untouched rext suites carry forward at v1.8 counts).
- **Flake:** **0** — triple-clean release gate 3/3 (stack-seeding incl `-race` + clerkenstein, shuffled; Python
  re-verified).
- **Supply-chain:** **1 NEW DEP at M45 — deliberate + sanctioned, not a regression.** Through M44 it was GREEN
  (0 new deps since v1.8). **M45 adds exactly `github.com/anthropos-work/ai@v1.40.1`** (all-permissive transitive
  tree: Azure SDK + openai-go/v3, MIT/BSD/Apache) — the `services/ai/` wrapper's transport, the user-acknowledged
  in-release inflection the milestone is ABOUT (the design-roadmap approved adding the LLM engine + this dep).
  License-vetted, version-boundaried (pinned across consumers, Go 1.25.0); the 5-pass harden added 0 further deps,
  `go mod tidy` is a no-op. M43's `cockpit.py` is stdlib-only Python + FontAwesome is a CDN `<link>` (not a dep).
- **Alignment gates (green since v1.0):** **100%/100%** on **all 5** Clerkenstein surfaces (Go 22/22, JS 9/9,
  multi 9/9, deploy 7/7, express 13/13) + drift 9/9 — re-verified at M42e close. M43 (cockpit) + M45 (stack-seeding
  only — empty clerkenstein diff) touched neither the Clerkenstein contract surfaces — gates carry forward (N/A change).

## Branch model
**v1.10 IN DEVELOPMENT (EXTENDED, not yet merged to main):** `release/01.10-method-acting` cut from `main`
2026-06-24. Milestone branches `m{39..42m}/{slug}` branched from it (all 5 merged + deleted); rext code-of-record
lands in the `rosetta-extensions` authoring copy (a SEPARATE repo) at v1.10 tags as milestones close. Closed-milestone
rext tags: `method-acting-m39` · `m40` · `m41` · `m42e` → `53574ae` · **`m42m-harden-final`** ·
**`m43-cockpit-ux-fix1`** · **`m44-profile-completeness-fix2`** · **`m45-harden-final`**. The `m43/cockpit-ux` +
`m44/profile-completeness` + `m45/generation-engine` branches merged `--no-ff` + deleted; the M46 branch
(`m46/org-scale-fill`) is cut by `/developer-kit:build-mstone-iters` as work begins. The release branch is NOT
merged to main yet — `/developer-kit:close-release` follows AFTER M46.
**v1.9 SHIPPED:** `release/01.90-storytelling` merged `--no-ff` → `main` + tagged `v1.9`; pushed to origin
2026-06-24. The release branch is deleted. v1.9 rext code-of-record at tags `storytelling-m34..m38`.
**Shipped:** **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` ·
**v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-26 (**M45 generation engine CLOSED** via `/developer-kit:close-milestone` — `iterative`,
closed-on-gate, GREEN, 3 findings / 0 blocking. The third of the M43→M46 extension; the cheap-LLM batch profile
generator with the CODE-owns-structure / AI-owns-content boundary, built inside-out over 7 iters. **Gate MET 5/5
on a REAL Azure gpt-4o-mini run** (EU-first Sweden), N=20 + demo-3 proof: valid-JSON 100%, 47/47 skills + 20/20
roles → real `skiller.*.node_id` (`datadna measure-closure` = `[PASS]`, 0 fabrication), 0 hero-collisions,
$0.0059 ≤ $0.10, $0 byte-identical re-seed; 20/20 distinct multicultural names + avatars, isolation CLEAN.
**Supply-chain: the deliberate 1-new-dep inflection** `anthropos-work/ai@v1.40.1` (the FIRST since v1.8, the
milestone is ABOUT it — NOT a regression). 5-pass final harden stabilized; 0 flakes; `-race` clean; stack-seeding
567→677. NEW `ai-generation-spec.md` + `cache-spec.md` (indexed). Deferral GREEN (org-scale → M46 Fate-2 owned);
alignment N/A; zero canonical platform edits. Merged `--no-ff` into `release/01.10-method-acting`; rext @ tag
`method-acting-m45-harden-final`. **M46 is the ONLY remaining v1.10 milestone.** Next: M46 (org-scale fill —
`iterative`, via `/developer-kit:work-mstone-iters`) → close-release.
Prior: **M44 profile completeness CLOSED 2026-06-26** [`section`, GREEN; §A trajectory self-rating, §B NEW
Certificates/ProjectsSeeder, §C manager personal data, §D bulk-member career + the avatar fix; NEW
`profile-completeness-spec.md`; rext tag `…-m44-profile-completeness-fix2`]; **M43 cockpit UX CLOSED 2026-06-26**
[`section`, GREEN; the light presenter launcher — unified `[Log in as]` CTA + icons + manifest download + overlay;
NEW `cockpit-spec.md`; rext tag `…-m43-cockpit-ux-fix1`]; **v1.10 EXTENDED with M43→M46** [2026-06-26, via
`/developer-kit:design-roadmap` — the presenter-grade / scalable-generation extension]; **M42m manager 100%
coverage CLOSED 2026-06-26** [iterative on-gate, 5 iters; the `demopatch` tool + FeedbackSeeder mirror].
Prior: **v1.9 "storytelling" SHIPPED 2026-06-23** [tag `v1.9`, M34–M38].)_
