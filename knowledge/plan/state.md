---
active_release: "v1.10b fit-up (interposed backfill; v2.0 opening night PAUSED)"
active_branch: "release/01.10b-fit-up"
active_milestone: "M47 — Re-sync & recapture (BUILD-COMPLETE — all sections S1-S6 done; ready for close-milestone)"
phase: "v1.10b building — M47 build-complete; next: close M47 → M48 corpus re-ground (the real staleness)"
last_updated: "2026-06-29"
---

# State

**Active release:** **v1.10b "fit-up" — IN DEVELOPMENT** (designed 2026-06-29 via
`/developer-kit:design-roadmap`; branch `release/01.10b-fit-up` cut from `main`). An **interposed field-hardening
backfill** (the v1.3b "dress rehearsal" lineage): a from-scratch `/demo-up` surfaced 8 bring-up issues + a tail of
v1.10 content gaps. **CORRECTION (M47 finding, 2026-06-29):** the M201 close *reported* the `stack-demo` clones
~5 weeks / 115+ commits behind prod (next-web @ v2.33.2), but **M47 found the clones actually CURRENT** (next-web @
v2.89.0, every repo ≤2 behind; the **AI-readiness feature is present** in `app`) — the re-sync was a trivial
`make pull`. The genuinely-stale surface is the **rosetta corpus** (e.g. the shipped AI-readiness feature is
**undocumented**), which **M48** re-grounds. So v1.10b: snapshot recaptured from current prod (M47), corpus
re-grounded (M48), the 8 bring-up issues + v1.10 content gaps fixed (M49/M50), a curated **AI-readiness showcase
org** added (M51), **one auditable seed+gen manifest** consolidated (M52), cold-rebuild proof (M53). The v1.x flat
counter re-opens at **M47**; tag **`v1.10.1`**. **Tooling + docs only — zero platform-repo edits.** 7 milestones:

```
M47 ──→ ┌ M48 corpus re-ground ───────────┐                (M48 ∥ M49 — disjoint clusters; M48 no-demo)
        └ M49 bring-up hardening ──────────┘ ──→ M50 ──→ M51 ──→ M52 ──→ M53
```

**Active milestone:** **M47 — Re-sync & recapture** (`section`, **BUILD-COMPLETE — ready for close-milestone**).
All sections S1-S6 done: rext authoring copy cloned to `.agentspace` (S1); the **sslmode `no-verify→require` fix**
(`pg.NormalizeDSN`, rext `c5323a1`, **tagged `fit-up-m47`**) so the wired MCP DSN works as a capture `--dsn` (S2,
demo-up #2, proven live); **clones confirmed current** (S3 — a trivial `make pull`, NOT a 5-week re-sync — see the
CORRECTION above); **all 3 snapshot surfaces recaptured** from current prod — directus + sim-embeddings + taxonomy,
all schema digests unchanged (S4); **AI-readiness feature confirmed present** in `app` (S5 — M201's false-negative
resolved); **`snapshot-cold-start.md` updated** (S6, KB-47-01). The ⚠ "biggest unknown" (heavy re-sync)
**evaporated** — the clones were already current. Consumption-clone re-pin deferred (needs the tag pushed; the fix
only affects capture). Records:
[`releases/01.10b-fit-up/m47-resync-recapture/`](releases/01.10b-fit-up/m47-resync-recapture/).

**Phase:** **v1.10b BUILDING — M47 in progress** (clones found current; see the CORRECTION above). Designed from the field review
[`.agentspace/annotation.md`](../../.agentspace/annotation.md) (8 demo-up issues + the content/hero/manager gaps) +
the M201 stale-clone finding, via 3 parallel research agents (fix surfaces at file:line, content/seeding gaps, KB
blind-areas). User decisions captured: **re-ground first** (the clones turned out **current** — a trivial pull; the
~1-month-stale **corpus** is the real re-ground → M48),
codename **"fit-up"**, the manifest as **one inlined file**, and a new **AI-readiness showcase org** (M51,
redeeming the M201 member-AI-readiness false-negative). The 4 KB blind-areas (cold-start MCP auto-capture,
ant-academy-in-demo, the unified seed+gen manifest, the rext tag-pin source-of-truth) are each homed via a milestone
`Delivers →` line.

**Next up:** **close M47** (`/developer-kit:close-milestone` — merges `m47/resync-recapture` →
`release/01.10b-fit-up`), then **M48 — corpus re-ground** (`/developer-kit:build-milestone`): document the
**AI-readiness feature** + reconcile the ~1-month-stale `corpus/architecture` + `corpus/services` docs against the
now-confirmed-current clones — the **genuine** staleness. The 1-demo-stack constraint serializes verification:
fix-on-live across M48→M52, then **M53 destroys + cold-rebuilds** as the single acceptance proof. _(The orchestrator still owes origin the pushes: `main` + the `v1.10` tag + the v1.10 ext tags —
the v1.10 LOCAL close did not push; the M201 close merged to `main` LOCALLY; this v1.10b branch is cut from that
local `main`.)_

**Last shipped:** **v1.10 — 2026-06-27** (`method acting`, 9 milestones M39→M46, tag `v1.10`,
`release/01.10-method-acting` merged `--no-ff` → `main`). The **last release of the v1.x major**; its history +
the full shipped log now live in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
[`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).

**Paused:** **v2.0 "opening night" (Playthroughs)** — paused 2026-06-29 after M201 closed, to interpose the
**v1.10b "fit-up"** backfill (re-sync + re-ground + re-validate + fix). M201 corpus preserved as the v2.0 spec;
M202 ∥ M203 ∥ M204 not started. Resume after v1.10b ships (tag `v1.10.1`).

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
**v1.10b IN DEVELOPMENT (active):** `release/01.10b-fit-up` cut from `main` 2026-06-29 (LOCAL — origin push is the
orchestrator's step). Milestone branches `m{47..53}/{slug}` branch from it at build time. rext code of record (a
SEPARATE repo) is authored in the `.agentspace/rosetta-extensions/` copy (cloned in M47) + tagged `fit-up-m47..m52`
per the tooling policy, rolled to the `v1.10.1` release tag at M53; the consumed tag is pinned via the new
`.agentspace/rext.tag` source-of-truth (M49 #1).
**v2.0 PAUSED:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL). M201 merged → `main` (LOCAL, no
`v2.0` tag); M202→M204 not started — resumes after v1.10b. A `playthroughs` rext section arrives at M202 build.
**v1.10 SHIPPED:** `release/01.10-method-acting` merged `--no-ff` → `main` + tagged `v1.10` at close (LOCAL).
rext code @ tags `method-acting-m39..m46-servegrant-closure`.
**Shipped:** **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` ·
**v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.
(Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-06-29 (**v1.10b "fit-up" DESIGNED + PROMOTED** via `/developer-kit:design-roadmap` — the
interposed **field-hardening backfill**; **7 milestones M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53** [v1.x flat
counter re-opened]; branch `release/01.10b-fit-up` cut from `main`; tag `v1.10.1`. Designed from
`.agentspace/annotation.md` + the M201 stale-clone finding [3 research agents]. Re-grounds demo + the ~1-month-stale
corpus to current prod; fixes the 8 demo-up issues + the v1.10 content gaps; adds the **AI-readiness showcase org**
[M51]; consolidates **one inlined seed+gen manifest** [M52]; cold-rebuild acceptance [M53]. User decisions:
re-ground first, codename "fit-up", one inlined manifest, +AI-readiness org. Active milestone now **M47**; next:
`/developer-kit:build-milestone`. Tooling + docs only. Prior: 2026-06-29 **M201 "Manifest corpus" CLOSED-on-gate** —
9 products · 26 stories · 28 use-cases, adversarially re-grounded [wf `wvpnpvozh`], user-signed-off; closed on
`release/02.00-opening-night` + merged → `main` per the user, no `v2.0` tag. **v2.0 PAUSED** for this backfill,
triggered by the **stale-clone discovery** [next-web 115+ commits behind prod — *later corrected: M47 found the
clones current; the stale surface is the corpus, → M48*]. M201 corpus preserved as the v2.0
spec. Prior: 2026-06-28 (**v2.0 "opening night" — M201 Manifest corpus INSERTED as the new first milestone**;
the prior 3 milestones renumbered M201→M202 [foundation], M202→M203 [employee], M203→M204 [manager]. 4 milestones
M201 ∥ M202 → { M203 ∥ M204 } [`Mxyy` numbering]: M201 `iterative`+user-guided manifest corpus is the prose
build+regression contract, authorable in parallel with the M202 `section` foundation. Active milestone now M201;
next: work it WITH THE USER via `/developer-kit:work-mstone-iters`. Future v2 milestones bumped to M205
Hiring/tier-gates, M206 AI-sim-mirror-tier, M207 Academy. Prior: 2026-06-28 v2.0 DESIGNED + PROMOTED — a NEW MAJOR
opening the Playthroughs pillar, branch `release/02.00-opening-night` cut from `main`, from
`spec-drafts/playthroughs/spec.md` v0.3. Headline numbers reset to the v1.10-close inheritance baseline.)_
