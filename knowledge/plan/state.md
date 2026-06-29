---
active_release: "v2.0 opening night (PAUSED — v1.10 backfill interposed)"
active_branch: "main"
active_milestone: "(between — M201 CLOSED; v2.0 paused for the v1.10 backfill)"
phase: "v2.0 PAUSED — v1.10 backfill next (user-driven)"
last_updated: "2026-06-29"
---

# State

**Active release:** **v2.0 "opening night" — IN DEVELOPMENT** (designed 2026-06-28 via
`/developer-kit:design-roadmap`; branch `release/02.00-opening-night` cut from `main`). A **NEW MAJOR** — opens
the **Playthroughs** pillar: functional-flow *testing*, a manifest-driven deterministic e2e suite that *pretends
to be the human* and proves the platform's core user journeys **actually work** end-to-end (the **functional**
sibling of v1.x's M42 **presence** coverage sweep). Governed by the consolidated capability spec
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (v0.3). v2+ uses **`Mxyy`** milestone
numbering. **Tooling + docs only — zero platform-repo edits.** 4 milestones:

```
M201 ──┐                 (manifest corpus — prose, user-guided)
M202 ──┼──→ M203 ─┐
            M204 ─┴──→ ship
```

**Active milestone:** **(between — M201 CLOSED 2026-06-29; v2.0 paused).** **M201 "Manifest corpus"**
(`iterative`, user-guided) **closed-on-gate** — the prose-intent Playthroughs manifest corpus (**9 products · 26
stories · 28 use-cases**) authored, **adversarially re-grounded** (11-agent workflow `wvpnpvozh` → 15/27 runnable),
and **signed off** by the user as the complete-enough v2.0 coverage contract. Deliverable:
[`releases/02.00-opening-night/m201-manifest-corpus/manifest-draft.yaml`](releases/02.00-opening-night/m201-manifest-corpus/manifest-draft.yaml).
The close surfaced the **stale-clone drift** (see Phase) → the v2.0 pause.

**Phase:** **v2.0 PAUSED for a v1.10 BACKFILL (user-directed).** M201's adversarial verify discovered the
stack-demo clones are **5 weeks / 115+ commits behind prod** (next-web @ v2.33.2; backends at early-June tags) and
the corpus likewise lags shipped features (the member-AI-readiness flow exists in prod + has live customers, but
is invisible to the stale clones — a proven false negative). So **everything downstream — demo, seeders, corpus,
and M201's own verify verdicts — is graded against stale code.** v2.0 is paused to interpose a dedicated backfill:
**re-sync the corpus + stack clones to current prod, then re-validate.** The M201 corpus is **preserved as the v2.0
spec**, resumable after the backfill.

**Next up:** **the user kicks off + runs the v1.10 backfill** (re-sync clones + re-ground corpus → re-validate the
M201 negative verdicts → fix the real gaps). The user owns the backfill design (`/developer-kit:design-roadmap`).
v2.0 "opening night" (Playthroughs) resumes after. _(The orchestrator still owes origin the pushes: `main` + the
`v1.10` tag + the v1.10 ext tags — the v1.10 LOCAL close did not push; this M201 close merges to `main` LOCALLY.)_

**Last shipped:** **v1.10 — 2026-06-27** (`method acting`, 9 milestones M39→M46, tag `v1.10`,
`release/01.10-method-acting` merged `--no-ff` → `main`). The **last release of the v1.x major**; its history +
the full shipped log now live in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
[`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).

**Paused:** **v2.0 "opening night" (Playthroughs)** — paused 2026-06-29 after M201 closed, to interpose the
**v1.10 backfill** (re-sync + re-ground + re-validate). M201 corpus preserved as the v2.0 spec; M202 ∥ M203 ∥ M204
not started. Resume after the backfill.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes),
DEF-M21-01 (`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`). Pre-existing, tracked in
[`roadmap-vision.md`](roadmap-vision.md); none scheduled. **Future v2 milestones** (Playthroughs pillar, NOT
pre-assigned to a minor): M205 Hiring + tier gates · M206 AI-sim mirror tier · M207 Academy coverage — also in
`roadmap-vision.md`.

## Recently shipped releases
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The **believable-profile release + the presenter-grade
  / scalable-generation extension**: a logged-in hero reads as a fully-fleshed person, proven by a **Playwright
  SEMANTIC coverage gate** at BOTH vantages cold (M42e employee / M42m manager), extended with M43 cockpit UX,
  M44 whole-roster completeness, M45 a cheap-LLM generation engine (first new dep, `ai@v1.40.1`), M46 org-scale
  fill. 9 milestones. Zero platform-repo edits; all 5 Clerkenstein gates 100%/100%. The **last v1.x release** —
  its detail + the full shipped log are in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. The declarative **Stories & Heroes** seeding engine + a
  presenter cockpit. 5 `section` milestones M34→M38.
  Records: [`releases/archive/01.90-storytelling/`](releases/archive/01.90-storytelling/).
- **v1.8 "understudy"** — **2026-06-15**, tag `v1.8`. Self-contained-demo release (a box with only `stack-demo/`
  runs a demo end-to-end). Single `section` milestone M26.
  Records: [`releases/archive/01.80-understudy/`](releases/archive/01.80-understudy/).
- **Earlier v1.x** (v1.0 … v1.7) — the full shipped table is in
  [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.

## Headline numbers (v2.0 — inheriting the v1.10-close baseline; no v2 work yet)
The v2.0 baseline is the v1.10-close inheritance — no v2.0 milestone has built yet, so these are the carried-over
totals (re-measured at first milestone close):
- **Go test funcs (rext):** **1551** total (`Test`+`Fuzz`) at the v1.10 close. Per-module: `alignment` 52 ·
  clerkenstein 270 · stack-seeding 706 · stack-snapshot 363 · stack-secrets 160. (A new `playthroughs` rext
  section arrives in M202; its first tests land at M202 build/close.)
- **Python / TS:** the cockpit `cockpit.py` suite 63 + the demopatch suite 43; `stack-injection` 117. The
  `@playwright/test ^1.49.0` coverage harness (M42e) is the e2e foundation M202 reuses (the first non-Go rext
  dev/test dep).
- **Flake:** **0** at v1.10 close.
- **Supply-chain:** the v1.10 close carried **1 new dep** (`github.com/anthropos-work/ai@v1.40.1`, M45). v2.0 has
  added none yet. The rosetta corpus is docs-only (no package manifest). Lockfile inherited from
  [`releases/archive/01.10-method-acting/dependencies.lock`](releases/archive/01.10-method-acting/dependencies.lock).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces at v1.10 close — v2.0 touches no contract
  surface so far.

## Branch model
**v2.0 IN DEVELOPMENT:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL — origin push is the
orchestrator's step). Milestone branches `m{201,202,203,204}/{slug}` will branch from it at build time. rext code
of record (a SEPARATE repo) gains a `playthroughs` section, authored + tagged per the tooling policy at M202 build.
**v1.10 SHIPPED:** `release/01.10-method-acting` merged `--no-ff` → `main` + tagged `v1.10` at close (LOCAL).
rext code @ tags `method-acting-m39..m46-servegrant-closure`.
**Shipped:** **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` ·
**v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.
(Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-06-29 (**M201 "Manifest corpus" CLOSED-on-gate** — 9 products · 26 stories · 28 use-cases,
adversarially re-grounded [wf `wvpnpvozh`], user-signed-off; closed on `release/02.00-opening-night` + merged →
`main` per the user, no `v2.0` tag. **v2.0 PAUSED** for a user-driven **v1.10 backfill** [re-sync + re-ground +
re-validate] triggered by the **stale-clone discovery** [next-web 115+ commits behind prod]. M201 corpus preserved
as the v2.0 spec. Prior: 2026-06-28 (**v2.0 "opening night" — M201 Manifest corpus INSERTED as the new first milestone**;
the prior 3 milestones renumbered M201→M202 [foundation], M202→M203 [employee], M203→M204 [manager]. 4 milestones
M201 ∥ M202 → { M203 ∥ M204 } [`Mxyy` numbering]: M201 `iterative`+user-guided manifest corpus is the prose
build+regression contract, authorable in parallel with the M202 `section` foundation. Active milestone now M201;
next: work it WITH THE USER via `/developer-kit:work-mstone-iters`. Future v2 milestones bumped to M205
Hiring/tier-gates, M206 AI-sim-mirror-tier, M207 Academy. Prior: 2026-06-28 v2.0 DESIGNED + PROMOTED — a NEW MAJOR
opening the Playthroughs pillar, branch `release/02.00-opening-night` cut from `main`, from
`spec-drafts/playthroughs/spec.md` v0.3. Headline numbers reset to the v1.10-close inheritance baseline.)_
