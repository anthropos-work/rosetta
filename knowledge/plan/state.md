# State

**Active version:** **v1.10 "method acting" — FEATURE-COMPLETE (all 9 milestones M39→M46 CLOSED; awaiting
close-release)** (designed 2026-06-24, **extended 2026-06-26**, via `/developer-kit:design-roadmap`; branch
`release/01.10-method-acting`). The **believable-profile release** (M39–M42m) **+ the presenter-grade /
scalable-generation extension** (M43→M46), **all CLOSED**. The built whole: when a presenter clicks **Login as** a
hero, the individual's **profile** (org name, role+title, work history, education, a real face, deep role-aligned
skills) **and** the content surfaces (**library** + the **activity feed**) populate with real content, on **every**
page a hero of that vantage can reach — proven by a **Playwright** coverage sweep with **zero** empty pages and
**zero** out-of-demo escapes (5 milestones M39→M42m). The extension: M43 **cockpit UX polish**, M44 **profile
completeness** (the whole roster fully baked), M45 a cheap-LLM **generation engine** (`cmd/gen-batch` + a
prompt-keyed cache + the CODE-owns-structure / AI-owns-content boundary), M46 **org-scale fill** + a gen-batch
preview CLI (a whole ~500/735-member org fills from one supporting-population descriptor, believable under the M42
sweep cold). **Tooling + docs only — zero platform-repo edits.** Grounded by the live-demo
review [`.agentspace/profile_gaps.md`](../../.agentspace/profile_gaps.md) + workflow `w7t4wq2z4`, and the
v1.10-extend research note [`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../.agentspace/scratch/roadmap-research-2026-06-26.md).

**Active milestone:** **(between milestones — v1.10 feature-complete, ready for close-release).** All 9 v1.10
milestones (M39→M42m + the M43→M46 extension) are **CLOSED**; nothing remains to build.
**Next up:** **`/developer-kit:close-release`** (merge `release/01.10-method-acting` → `main`, tag `v1.10`) — the
**user's next step, after their visual review** of the org-scale demo (demo-3 is up as the review stack). _(This
close did NOT run close-release — that is the user's call.)_
**Last closed:** **M46 — 2026-06-27** (`org-scale-fill`, `iterative`, `closed-on-gate`) via
`/developer-kit:close-milestone`. Org-scale fill + a gen-batch preview CLI: scale the M45 engine to filling an
ENTIRE org from one supporting-population descriptor (per-story, deterministic auto-fill); 7 iters (1 tok + 6 tiks).
**Gate MET 5/5, robustly COLD** — a ~500/735-member org fills believably; the **M42 Playwright SEMANTIC sweep
PASSES on BOTH vantages cold** (employee `(0,0,0,0,0)`; manager `failingSections=0, gateMet=true`); 0
hero-collisions on a real ~600-member Azure batch; closure GREEN; cost/throughput in budget; $0 cache-hit reseed.
The **5th gate face** (the manager enterprise grids — a platform-bound org-scale render wall) was cleared by a
**demo-patch / recapture campaign** (ZERO canonical edits): T1 (next-web pagination + FK indexes) · B (`roles.go`
authz-skip read-gate drop) · DD (a column-drift backfill) · SG/Path 2 (the Directus serve-grant deep-fetch
**CLOSURE** + a prod-structure RECAPTURE over `marco_read`, firewall public-only/0 tenant rows — **resolving
`DEF-M46-01`**). iter-07's `exit-3 (re-scope-trigger)` was **superseded** by the campaign. Harden-equivalent: the
4-sub-agent demo-patch/recapture campaign + a fresh `--purge /demo-up 3` proof. **0 new deps**; alignment N/A;
**zero canonical edits**. Merged `--no-ff` into `release/01.10-method-acting`; rext @ tag
`method-acting-m46-servegrant-closure`. (Full breakdown in the closing block below.)
**Phase:** **v1.10 FEATURE-COMPLETE — all 9 milestones CLOSED (M39/M40/M41 `section` + M42e/M42m `iterative` +
M43/M44 `section` + M45/M46 `iterative`). Release NOT merged to main. Next: `/developer-kit:close-release` (v1.10 →
main + tag `v1.10`) — the user's step after visual review.**
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

## Headline numbers (v1.10 — through M46 close, 2026-06-27; FEATURE-COMPLETE)
- **Go test funcs:** **1551** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 52 ·
  clerkenstein **270** · stack-seeding **706** · stack-snapshot **363** · stack-secrets 160. (Ground-truth grep in
  the `.agentspace/rosetta-extensions` authoring copy at tag `method-acting-m46-servegrant-closure`.)
  **M46 delta: stack-seeding 677→706 (+29)** — the org-scale + per-story + preview + `--gen-batches`-fence
  deliverables + the 5 org-scale regression-test classes (multi-batch cache-index collision · seed-time name
  distinctness · email distinctness · the fence · 429-fallback · preview · per-story routing · auto-fill);
  **stack-snapshot 361→363 (+2)** — the serve-grant deep-fetch closure tests (`IncludesDeepFetchClosure` +
  `SynthesizesGuardedReadGrant`). clerkenstein/alignment/secrets UNCHANGED (M46 did not touch them — verified empty
  clerkenstein diff m45..m46). **NB clerkenstein reconciled 266→270:** ground-truth is 270 at BOTH the M45 and M46
  tags; the prior headline carried 266 (a recorded-vs-grep close-drift)
  — reconciled to the grep ground-truth here. **M46 harden-equivalent:** no formal `--final` pass; the demo-patch /
  recapture verification campaign (4 adversarial sub-agents, each full-suite + M42 sweep, + a fresh `--purge
  /demo-up 3` reproducibility proof) substitutes. **M45 delta (prior): stack-seeding 567→677 (+110)** — the 5 new
  pkgs (services/ai + blueprint + batchcache + cmd/gen-batch) + the GeneratedBatchSeeder tests + the 5-pass final
  harden (coverage 88.5-98.9% across the new pkgs). Prior milestone deltas (M39–M44) below.
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
  re-verified). **M46:** stack-seeding + stack-snapshot suites GREEN across the build + the 4-pass demo-patch
  campaign; the 1 demo-stack pytest non-pass is a PRE-EXISTING `ensure-clones` SC2015 shellcheck info (line 154,
  untouched — not introduced by M46, not a flake).
- **Supply-chain:** **1 NEW DEP at M45 — deliberate + sanctioned, not a regression; M46 = 0 NEW DEPS (GREEN).**
  Through M44 it was GREEN (0 new deps since v1.8). **M45 adds exactly `github.com/anthropos-work/ai@v1.40.1`**
  (all-permissive transitive tree: Azure SDK + openai-go/v3, MIT/BSD/Apache) — the `services/ai/` wrapper's
  transport, the user-acknowledged in-release inflection the milestone is ABOUT (the design-roadmap approved adding
  the LLM engine + this dep). License-vetted, version-boundaried (pinned across consumers, Go 1.25.0); the 5-pass
  harden added 0 further deps. **M46 reuses the `ai` dep unchanged** — `go.mod`/`go.sum` unchanged across its
  footprint, `go mod tidy` a no-op. M43's `cockpit.py` is stdlib-only Python + FontAwesome is a CDN `<link>` (not a
  dep).
- **Alignment gates (green since v1.0):** **100%/100%** on **all 5** Clerkenstein surfaces (Go 22/22, JS 9/9,
  multi 9/9, deploy 7/7, express 13/13) + drift 9/9 — re-verified at M42e close. M43 (cockpit) + M45 + M46
  (stack-seeding / stack-snapshot only — empty clerkenstein diff) touched none of the Clerkenstein contract
  surfaces — gates carry forward (N/A change).

## Branch model
**v1.10 FEATURE-COMPLETE (all 9 milestones closed, release NOT yet merged to main):** `release/01.10-method-acting`
cut from `main` 2026-06-24. All milestone branches `m{39..46}/{slug}` branched from it, **all merged `--no-ff` +
deleted**; rext code-of-record lands in the `rosetta-extensions` authoring copy (a SEPARATE repo) at v1.10 tags as
milestones close. Closed-milestone rext tags: `method-acting-m39` · `m40` · `m41` · `m42e` → `53574ae` ·
**`m42m-harden-final`** · **`m43-cockpit-ux-fix1`** · **`m44-profile-completeness-fix2`** · **`m45-harden-final`** ·
**`m46-servegrant-closure`** (the M46 code-of-record; iter tags `m46-iter02..iter07` + the close-path tags
`m46-{gridperf,authz-skip,directus-drift-fix,servegrant-closure}` also exist). The `m46/org-scale-fill` branch
merged `--no-ff` + deleted at this close. The release branch is **NOT merged to main yet** —
`/developer-kit:close-release` (the user's next step, after visual review) follows.
**v1.9 SHIPPED:** `release/01.90-storytelling` merged `--no-ff` → `main` + tagged `v1.9`; pushed to origin
2026-06-24. The release branch is deleted. v1.9 rext code-of-record at tags `storytelling-m34..m38`.
**Shipped:** **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` ·
**v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-27 (**M46 org-scale-fill CLOSED** via `/developer-kit:close-milestone` — `iterative`,
`closed-on-gate`. The FOURTH + FINAL of the M43→M46 extension → **v1.10 is now FEATURE-COMPLETE.** Org-scale fill +
a gen-batch preview CLI: scale the M45 engine from a bounded batch to filling an ENTIRE org from one
supporting-population descriptor (per-story, deterministic auto-fill), built fixtures-first over 7 iters (1
bootstrap tok + 6 tiks). **Gate MET 5/5, robustly COLD:** a ~500/735-member org fills believably; the **M42
Playwright SEMANTIC sweep PASSES on BOTH vantages cold** (employee `(0,0,0,0,0)`; manager `failingSections=0,
gateMet=true, personaFailures=0, escapes=0`); 0 hero-collisions on a real ~600-member Azure gpt-4o-mini batch;
closure GREEN (0 fabrication); cost/throughput in budget; $0 byte-distinct cache-hit reseed. The 5th gate face (the
manager enterprise grids — a platform-bound org-scale render wall) was cleared by a **demo-patch / recapture
campaign** (ZERO canonical edits): T1 (next-web pagination + FK indexes, 84s→~4s) · B (`roles.go` authz-skip read-
gate drop, members 76.7s→0.51s) · DD (a reproducible column-drift backfill) · SG/Path 2 (the Directus serve-grant
deep-fetch CLOSURE — `servedCollections` +7 + a prod-structure RECAPTURE over the sanctioned `marco_read` DSN,
firewall public-only/0 tenant rows — **resolving `DEF-M46-01`**). iter-07's `exit-3 (re-scope-trigger)` was
**superseded** by the campaign (the grids WERE demo-patchable). Harden-equivalent: the 4-sub-agent
demo-patch/recapture verification campaign + a fresh `--purge /demo-up 3` reproducibility proof (exceeds a standard
harden). 0 new deps; alignment N/A; zero canonical platform edits; stack-seeding 677→706, stack-snapshot 361→363.
Merged `--no-ff` into `release/01.10-method-acting`; rext @ tag `method-acting-m46-servegrant-closure`. **Next:
`/developer-kit:close-release` (v1.10 → main + tag `v1.10`) — the user's step after visual review.**
Prior: **M45 generation engine CLOSED 2026-06-26** [`iterative` on-gate, 7 iters; the cheap-LLM batch profile
generator + the CODE-owns-structure / AI-owns-content boundary; gate MET 5/5 on a real Azure gpt-4o-mini run; the
deliberate 1-new-dep inflection `ai@v1.40.1`; NEW `ai-generation-spec.md` + `cache-spec.md`; rext tag
`…-m45-harden-final`]; **M44 profile completeness CLOSED 2026-06-26** [`section`, GREEN; whole-roster baked +
Certificates/ProjectsSeeder + the avatar fix; NEW `profile-completeness-spec.md`]; **M43 cockpit UX CLOSED
2026-06-26** [`section`, GREEN; the light presenter launcher; NEW `cockpit-spec.md`]; **M42m manager 100% coverage
CLOSED 2026-06-26** [iterative on-gate, 5 iters; the `demopatch` tool + FeedbackSeeder mirror].
Prior: **v1.10 EXTENDED with M43→M46** [2026-06-26]; **v1.9 "storytelling" SHIPPED 2026-06-23** [tag `v1.9`, M34–M38].)_
