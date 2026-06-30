---
active_release: "v1.10b fit-up (interposed backfill; v2.0 opening night PAUSED)"
active_branch: "release/01.10b-fit-up"
active_milestone: "M50 — Content & seeding fill (iterative; next to build)"
last_closed: "M49 — 2026-06-30 (Bring-up hardening + truth-up)"
phase: "v1.10b building — M47/M48/M49 CLOSED; next M50 (content & seeding fill, iterative)"
last_updated: "2026-06-30"
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

**Active milestone:** **M50 — Content & seeding fill** (`iterative`, **next to build**). **Exit gate:** Maya
(employee) + Dan (manager) render fully-populated across every annotation-listed surface on a COLD reset-to-seed
demo, proven by the **M42 semantic coverage gate** (both vantages), 0 prod-eject escapes. Iterative because the real
root-causes are **hypotheses until observed** on the clean bring-up (some need new seeders, e.g.
`MemberLanguagesSeeder`; some are demo-up #7 artifacts that vanish once set-dress runs). The **AI-provider-keys
policy** lands here (M49 routed it forward — which of OPENAI/ANTHROPIC/MISTRAL/ELEVENLABS/LIVEKIT become
throwaway/sandbox demo values vs documented-as-absent; gates the academy AI chat). Records:
[`releases/01.10b-fit-up/m50-content-seeding-fill/`](releases/01.10b-fit-up/m50-content-seeding-fill/). _(M49 closed
2026-06-30 — see Recently closed.)_

**Phase:** **v1.10b building — M47 + M48 + M49 CLOSED; next M50 (iterative).** Foundation (M47 re-sync/recapture) +
corpus re-ground (M48 — the new `ai-readiness.md` + the drift fixes) + bring-up hardening (M49 — the 7 demo-up fixes,
live-verified, rext @ `fit-up-m49`) are done. The first demo-touching milestone has shipped, so the M48 ∥ M49 window
has closed and the **single-demo serialization** (fix-on-live across M50→M52, then **M53 destroys + cold-rebuilds**
as the single acceptance proof) is now in effect.

**Next up:** **build M50** (`/developer-kit:work-mstone-iters` — iterative content & seeding fill toward the
both-vantage coverage gate). _(The orchestrator still owes origin the pushes: `main` + the `v1.10` tag + the v1.10
ext tags + now the `fit-up-m47..m49` rext tags — the v1.10 LOCAL close did not push; the M201 close merged to `main`
LOCALLY; this v1.10b branch is cut from that local `main`. The consumption-clone re-pin to `fit-up-m49` is the
push-gated KEEP tracked here, authoritatively bumped at M53.)_

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

## Recently closed (v1.10b milestones)
- **M49 — Bring-up hardening + truth-up** — **2026-06-30** (`section`; merged → `release/01.10b-fit-up`; rext tag
  `fit-up-m49` @ `ba586d6`). Closed the 7 remaining demo-up issues (rext.tag SoT + CRLF-tolerant reader, .env-guard
  order, `INVITATION_HMAC_SECRET` critical+auto-gen, ant-academy explicit clone, disk pre-flight + `down --purge`,
  true non-fatal frontend, demopatch re-anchor to v2.89.0) — **live-verified** from cold (demo-1 UP). Tests rext Go
  1552→1555 · demo-stack Python 299; flake 0. AI-keys policy → M50 (Fate-2). Full narrative in `roadmap.md` § M49.
- **M48 — Corpus re-ground** — **2026-06-29** (`section`; merged → `release/01.10b-fit-up`). NEW
  `corpus/services/ai-readiness.md` (the M51 seeder contract — Phase-2c-sharpened: active⇒signals-true,
  closed⇒frozen-snapshot) + drift reconciled + the false ant-academy "in repos.yml" claim corrected (M49 #5 owns
  the code fix). Docs-only; 3-agent investigation. Full narrative in `roadmap.md` § M48.
- **M47 — Re-sync & recapture** — **2026-06-29** (`section`; merged → `release/01.10b-fit-up`; rext `fit-up-m47`).
  The heavy re-sync was a no-op (clones already current); delivered `pg.NormalizeDSN` (the wired MCP DSN now works
  as a capture `--dsn`), recaptured all 3 snapshot surfaces (digests unchanged), confirmed the AI-readiness feature
  present (M201 false-negative resolved). Full narrative in `roadmap.md` § M47.

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

## Headline numbers (v1.10b — M49 close)
- **Go test funcs (rext):** **1555** total (`Test`+`Fuzz`) at M49 close (M47-close baseline 1552). Per-module:
  `alignment` 52 · clerkenstein 270 · stack-seeding 706 · stack-snapshot 364 · **stack-secrets 163** (+3 vs 160:
  the M49 #4 `INVITATION_HMAC_SECRET` gene asserts + the `TestMintedShape_DemoGeneratedKeyIsOpaque` harden probe).
- **Python / TS:** the `demo-stack` suite **299** (incl. `test_tooling` **111** · `test_demopatch` **47** ·
  `test_frontend_build` 55 · `test_cockpit` 63 · `test_ant_academy` 17); `stack-injection` 117. The
  `@playwright/test ^1.49.0` coverage harness (M42e) is the e2e foundation M202 reuses (the first non-Go rext
  dev/test dep).
- **Flake:** **0** (M49 close flake gate 5/5 on the M49-touched suite: 175 Python + the Go secretdna pkg).
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

_Last updated: 2026-06-30 (**M49 "Bring-up hardening + truth-up" CLOSED** via `/developer-kit:close-milestone` —
the 7 remaining demo-up issues, live-verified from cold [demo-1 UP, autoverify "verified-working"]; rext tag
`fit-up-m49` @ `ba586d6`; merged `--no-ff` → `release/01.10b-fit-up`. Close review: 6 findings all Fate-1 [2 stale
`make init` ant-academy claims in `service_taxonomy.md` retired; test-count literals reconciled]; deferral audit
**GREEN** [AI-keys policy → M50; consumption-clone re-pin → push-gated KEEP]. Tests rext Go 1552→1555 · demo-stack
Python 299; flake 0. Active milestone now **M50** (iterative content & seeding fill); next:
`/developer-kit:work-mstone-iters`. The M48 ∥ M49 parallel window is closed — single-demo serialization in effect.
Prior: 2026-06-29 **M48 "Corpus re-ground" CLOSED** [NEW `corpus/services/ai-readiness.md` + drift reconciled];
**M47 "Re-sync & recapture" CLOSED** [clones found current; `pg.NormalizeDSN`; snapshots recaptured]. Prior:
2026-06-29 **v1.10b "fit-up" DESIGNED + PROMOTED** — the interposed **field-hardening backfill**; 7 milestones
M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53 [v1.x flat counter re-opened]; branch `release/01.10b-fit-up` cut from
`main`; tag `v1.10.1`. Re-grounds demo + corpus to current prod, fixes the demo-up issues + content gaps, adds the
AI-readiness showcase org [M51], consolidates one inlined seed+gen manifest [M52], cold-rebuild acceptance [M53].
Tooling + docs only. Prior: 2026-06-29 **M201 "Manifest corpus" CLOSED-on-gate** [9 products · 26 stories · 28
use-cases, user-signed-off]; **v2.0 PAUSED** for this backfill. Full v2.0-design history in `roadmap.md`.)_
