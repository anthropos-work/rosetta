# State

**Active version:** **v1.10 "method acting" — IN DEVELOPMENT (EXTENDED with M43→M46)** (designed 2026-06-24,
**extended 2026-06-26**, via `/developer-kit:design-roadmap`; branch `release/01.10-method-acting`). The
**believable-profile release** (M39–M42m, all CLOSED) **+ the presenter-grade / scalable-generation extension**
(M43/M44 CLOSED; M45/M46 planned). The built half: when a presenter clicks **Login as** a hero, the individual's **profile** (org
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

**Active milestone:** **(between milestones — see Next Up).** The { M43 ∥ M44 } opening section-pair of the
extension is **both CLOSED**; the next milestone is **M45 (the cheap-LLM generation engine — `iterative`)**, now
unblocked (it reuses M44's certificate/project + bulk-member surfaces + the trajectory-aware self-rating).
**Next up:** **`/developer-kit:work-mstone-iters` for M45** (the now-unblocked **iterative** LLM-generation gate —
the interactive driver; NOT `work-milestone`), then **M46** (`/developer-kit:work-mstone-iters` — iterative
org-scale fill) → **THEN `/developer-kit:close-release`** (merge `release/01.10-method-acting` → `main`, tag
`v1.10`). _(Do NOT run close-release yet — M45 + M46 remain.)_
**Last closed:** **M43 cockpit UX (`section`) — 2026-06-26** via `/developer-kit:close-milestone` (GREEN, 0
findings / 0 blocking). All five UX fills delivered Fate-1 + render-accepted on a live demo-3 (`:37700`) +
USER-APPROVED the design: light restyle, FontAwesome FREE-CDN icons, one unified `[Log in as]` CTA per hero →
`jump_to` (no `cockpit.go` change), seed-manifest download, staged login-progress overlay. NEW
`corpus/ops/demo/cockpit-spec.md` (+ the `stories-spec.md` two-CTA→unified reconciliation + README/CLAUDE index
rows). Hardening 2-pass / cockpit Python tests 27→63 / a meta-line double-escape bug fixed + pinned / 0 flakes.
Deferral re-audit GREEN (0 deferrals — the cockpit future-feature surface is a `cockpit-spec.md` expansion surface,
not a deferral); supply-chain GREEN (0 new deps — `cockpit.py` stdlib-only, FA = CDN `<link>`); alignment N/A;
**zero canonical platform edits**. Merged `--no-ff` into `release/01.10-method-acting`; rext code-of-record @ tag
`method-acting-m43-cockpit-ux-fix1`.
**Phase:** **v1.10 reopened / extended — 7 milestones CLOSED (M39/M40/M41 `section` + M42e/M42m `iterative` + M43/M44
`section`), 2 PLANNED (M45/M46 `iterative`). Release NOT merged to main. Next: M45 → M46 →
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

## Headline numbers (v1.10 built half — through M43+M44 close, 2026-06-26; M45/M46 not started)
- **Go test funcs:** **1406** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **266** · stack-seeding **567** · stack-snapshot **361** · stack-secrets 160.
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
- **Supply-chain:** **GREEN** through M43+M44 (`go.mod`/`go.sum` byte-identical to v1.8; 0 new deps across both
  footprints — M43's `cockpit.py` is stdlib-only Python + FontAwesome is a CDN `<link>`, not a build dep; all deps
  MIT/BSD/Apache). **M45 deliberately adds the first new third-party AI dep** (the
  `services/ai/` wrapper over the shared `ai` library) — an in-release decision recorded on the M45 block;
  supply-chain will be re-reviewed at M45/close.
- **Alignment gates (green since v1.0):** **100%/100%** on **all 5** Clerkenstein surfaces (Go 22/22, JS 9/9,
  multi 9/9, deploy 7/7, express 13/13) + drift 9/9 — re-verified at M42e close. M43 touched the cockpit
  (`demo-stack/cockpit.py`), not the Clerkenstein contract surfaces — gates carry forward (N/A change).

## Branch model
**v1.10 IN DEVELOPMENT (EXTENDED, not yet merged to main):** `release/01.10-method-acting` cut from `main`
2026-06-24. Milestone branches `m{39..42m}/{slug}` branched from it (all 5 merged + deleted); rext code-of-record
lands in the `rosetta-extensions` authoring copy (a SEPARATE repo) at v1.10 tags as milestones close. Closed-milestone
rext tags: `method-acting-m39` · `m40` · `m41` · `m42e` → `53574ae` · **`m42m-harden-final`** ·
**`m43-cockpit-ux-fix1`** · **`m44-profile-completeness-fix2`**. The `m43/cockpit-ux` + `m44/profile-completeness`
branches merged `--no-ff` + deleted; the M45/M46 branches (`m45/generation-engine`, `m46/org-scale-fill`) are cut by
`/developer-kit:build-mstone-iters` as work begins. The release branch is NOT merged to main yet
— `/developer-kit:close-release` follows AFTER M46.
**v1.9 SHIPPED:** `release/01.90-storytelling` merged `--no-ff` → `main` + tagged `v1.9`; pushed to origin
2026-06-24. The release branch is deleted. v1.9 rext code-of-record at tags `storytelling-m34..m38`.
**Shipped:** **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` ·
**v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-26 (**M43 cockpit UX CLOSED** via `/developer-kit:close-milestone` — GREEN, 0 findings / 0
blocking. The second of the M43→M46 extension closed; the cockpit-UX half of the { M43 ∥ M44 } opening pair — with
both closed the opening section-pair is DONE. All five UX fills delivered Fate-1 + render-accepted on a live demo-3
(`:37700`) + USER-APPROVED: light restyle, FontAwesome FREE-CDN icons, one unified `[Log in as]` CTA per hero →
`jump_to` (no `cockpit.go` change), seed-manifest download, staged login-progress overlay. NEW
`corpus/ops/demo/cockpit-spec.md` + the `stories-spec.md` two-CTA→unified reconciliation + README/CLAUDE index rows.
Hardening 2-pass / cockpit Python tests 27→63 / a meta-line double-escape bug fixed + pinned / 0 flakes. Deferral
GREEN (0 deferrals — the cockpit future-feature surface is a `cockpit-spec.md` expansion surface); supply-chain
GREEN (0 new deps — `cockpit.py` stdlib-only, FA = CDN `<link>`); alignment N/A; zero canonical platform edits.
Merged `--no-ff` into `release/01.10-method-acting`; rext code-of-record @ tag `method-acting-m43-cockpit-ux-fix1`.
Next: M45 (generation engine — `iterative`, via `/developer-kit:work-mstone-iters`) → M46 → close-release.
Prior: **M44 profile completeness CLOSED 2026-06-26** [`section`, GREEN; the { M43 ∥ M44 } profile-completeness
half — §A trajectory self-rating, §B NEW Certificates/ProjectsSeeder, §C manager personal data, §D bulk-member
career + the `memberships.picture_url` avatar fix; NEW `profile-completeness-spec.md`; stack-seeding 538→567; rext
tag `method-acting-m44-profile-completeness-fix2`; **M45 UNBLOCKED**]; **v1.10 EXTENDED with M43→M46** [2026-06-26,
via `/developer-kit:design-roadmap` — the presenter-grade / scalable-generation extension; { M43 ∥ M44 } → M45 → M46,
close-release AFTER M46; research note `.agentspace/scratch/roadmap-research-2026-06-26.md`]; **M42m manager 100%
coverage CLOSED 2026-06-26** [iterative on-gate, 5 iters; the `demopatch` tool + FeedbackSeeder mirror]; **M42e
employee 100% coverage CLOSED 2026-06-25** [first iterative milestone, 23 iters, Playwright harness].
Prior: **v1.9 "storytelling" SHIPPED 2026-06-23** [tag `v1.9`, M34–M38].)_
