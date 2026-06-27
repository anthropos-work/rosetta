# State

**Active version:** **(between releases)** — v1.10 "method acting" **SHIPPED 2026-06-27** (tag `v1.10`,
`release/01.10-method-acting` merged `--no-ff` → `main`). No release is in development; the next step is
**`/developer-kit:design-roadmap`** to cut v2.

**Active milestone:** **(between releases).** Nothing in flight.
**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.**
**Next up:** **`/developer-kit:design-roadmap`** (design + cut the next release). _(The orchestrator pushes `main`
+ the `v1.10` tag + the ext tags to origin — the LOCAL close did not push.)_
**Last shipped:** **v1.10 — 2026-06-27** (`method acting`, 9 milestones M39→M46). close-release GREEN (0 blocking;
the single doc finding — an orphaned spec index — fixed at close). Records:
[`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).
**Paused:** _(none)_

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes),
DEF-M21-01 (`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`). Pre-existing, tracked in
[`roadmap-vision.md`](roadmap-vision.md); none scheduled. (DEF-M46-01 RESOLVED in v1.10.)

## Recently shipped releases
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The **believable-profile release + the
  presenter-grade / scalable-generation extension**: a logged-in hero reads as a fully-fleshed person (profile
  identity M39 + the Directus content-surface unblock M40 + profile depth M41), proven by a **Playwright SEMANTIC
  coverage gate** at BOTH vantages cold (M42e employee / M42m manager). Extended with M43 cockpit UX, M44
  whole-roster completeness, M45 a cheap-LLM **generation engine** (the first new dep, `ai@v1.40.1`), M46
  **org-scale fill** (~500/735-member org from one descriptor) + a gen-batch preview CLI. 9 milestones (3
  `section` + 4 `iterative` + 2 `section`). Headline: **zero platform-repo edits**; all 5 Clerkenstein gates
  100%/100%; rext Go 1248→**1551** (+303); 1 deliberate new dep; proven live on demo-3. The M46 manager
  org-scale render wall (a 5-layer activity-dashboard saga) cleared by a **demo-patch/recapture campaign** with
  ZERO canonical edits (resolving DEF-M46-01). Code: `rosetta-extensions` @ tags
  `method-acting-m39..m46-servegrant-closure`. Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. Believable-demo-narrative release: the placeholder
  seeder becomes a declarative **Stories & Heroes** engine (per-story org + a hero trio via the real
  verified-skill chain) + a presenter cockpit on Clerkenstein multi-identity. 5 `section` milestones
  **M34→M38**. Headline: zero platform-repo edits; all 5 Clerkenstein gates 100%/100%; 0 new deps; Go
  1027→**1248**. Records: [`releases/archive/01.90-storytelling/`](releases/archive/01.90-storytelling/).
- **v1.8 "understudy"** — **2026-06-15**, tag `v1.8`. Self-contained-demo release: a demo builds **entirely from
  `stack-demo`'s own clone set** (a box with only `stack-demo/` runs a demo end-to-end). Single `section`
  milestone **M26**. Code: `rosetta-extensions` @ tag `understudy-m26`. Records:
  [`releases/archive/01.80-understudy/`](releases/archive/01.80-understudy/).

## Headline numbers (v1.10 — through M46 close + close-release, 2026-06-27)
- **Go test funcs (rext):** **1551** total (`Test`+`Fuzz`). Per-module: `alignment` 52 · clerkenstein **270** ·
  stack-seeding **706** · stack-snapshot **363** · stack-secrets 160. (Ground-truth grep in the
  `.agentspace/rosetta-extensions` authoring copy at tag `method-acting-m46-servegrant-closure`.)
  **v1.9-close baseline 1248 → +303.** Drivers: stack-seeding 444→706 (+262; M45 `services/ai`/`blueprint`/
  `batchcache`/`cmd/gen-batch`/`GeneratedBatchSeeder` +110 is the biggest), stack-snapshot 333→363 (+30; M40
  serve-grant pkg +21 + M46 deep-fetch closure +2), clerkenstein 259→270 (+11; incl the recorded-vs-grep
  close-drift reconciled to the 270 ground-truth at the M45/M46 tags). alignment + stack-secrets unchanged.
- **Python / TS:** the cockpit `cockpit.py` suite 27→**63** (+36) + the demopatch suite 18→**43** (+25); the
  FIRST non-Go rext dev/test dep — the `@playwright/test ^1.49.0` coverage harness (M42e). `stack-injection` 117
  unchanged; no suite decreased.
- **Flake:** **0.** Per-milestone gates clean; M46 substituted a 4-sub-agent demo-patch/recapture verification
  campaign + a fresh `--purge /demo-up 3` reproducibility proof for a formal `--final` harden. The 1 demo-stack
  pytest non-pass is a PRE-EXISTING `ensure-clones` SC2015 shellcheck info (untouched) — not a flake.
- **Supply-chain:** **1 NEW DEP — deliberate + sanctioned** (`github.com/anthropos-work/ai@v1.40.1` at M45, the
  user-acknowledged in-release generation-engine inflection; v1.8→v1.9 was 0-new-deps). M46 reuses it unchanged.
  All deps MIT/BSD/Apache, no GPL/AGPL. The rosetta corpus is docs-only (no package manifest). Lockfile:
  [`releases/archive/01.10-method-acting/dependencies.lock`](releases/archive/01.10-method-acting/dependencies.lock).
- **Alignment gates:** **100%/100%** on **all 5** Clerkenstein surfaces (Go 22/22, JS 9/9, multi 9/9, deploy 7/7,
  express 13/13) + drift 9/9 — re-verified at M42e close; M40/M41/M43/M44/M45/M46 touched no contract surface.

## Branch model
**v1.10 SHIPPED:** `release/01.10-method-acting` cut from `main` 2026-06-24, merged `--no-ff` → `main` + tagged
`v1.10` at close (LOCAL — origin push is the orchestrator's step). All milestone branches `m{39..46}/{slug}`
merged `--no-ff` + deleted. rext code-of-record (a SEPARATE repo) @ tags `method-acting-m39` · `m40` · `m41` ·
`m42e` (`53574ae`) · `m42m-harden-final` · `m43-cockpit-ux-fix1` · `m44-profile-completeness-fix2` ·
`m45-harden-final` · `m46-servegrant-closure`.
**v1.9 SHIPPED:** `release/01.90-storytelling` merged → `main` + tagged `v1.9`; pushed 2026-06-24. rext code @
tags `storytelling-m34..m38`.
**Shipped:** **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` ·
**v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-27 (**v1.10 "method acting" SHIPPED** via `/developer-kit:close-release` — merged
`release/01.10-method-acting` `--no-ff` → `main` + tagged `v1.10`, LOCAL only. close-release GREEN: 9 review
sweeps clean, 0 blocking; the single doc finding — `profile-completeness-spec.md` orphaned from CLAUDE.md's index
— fixed at close. Scope clean (all 9 milestones Fate-1; M42e intra-milestone Fate-3 routes all RESOLVED;
DEF-M46-01 RESOLVED; 0 Fate-3-undelivered, 0 NEW escape-hatch). Deferral re-audit GREEN (standing backlog
pre-existing, tracked in roadmap-vision). 1 deliberate new dep; rext Go 1248→1551 (+303); all 5 Clerkenstein
gates 100%/100%; zero platform-repo edits. Next: `/developer-kit:design-roadmap`.)_
