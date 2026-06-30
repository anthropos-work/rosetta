---
active_release: "v1.10b fit-up (interposed backfill; v2.0 opening night PAUSED)"
active_branch: "release/01.10b-fit-up"
active_milestone: "M51 — AI-readiness showcase org (iterative; next to build)"
last_closed: "M50 — 2026-06-30 (Content & seeding fill, closed-on-gate)"
phase: "v1.10b building — M47..M50 CLOSED; next M51 (AI-readiness showcase org, iterative)"
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

**Active milestone:** **M51 — AI-readiness showcase org** (`iterative`, **next to build**). **Exit gate:** a curated
**200-person 3rd org** with the AI-readiness manager dashboard **ENABLED**, **~80%** of members having completed all
**3** onboarding/evaluation steps, **1 hero started + 1 hero completed** — proven by the M42 coverage gate (manager
vantage), 0 prod-eject escapes. Iterative because the feature's data model (invisible to the stale clones) was mapped
fresh in M48 (the `ai-readiness.md` contract). The academy course-content + menu-link/non-anonymous-session (F6, from
M50's Fate-3 handoff) is annotated to its candidate scope. Records:
[`releases/01.10b-fit-up/m51-ai-readiness-org/`](releases/01.10b-fit-up/m51-ai-readiness-org/). _(M50 closed-on-gate
2026-06-30 — see Recently closed.)_

**Phase:** **v1.10b building — M47..M50 CLOSED; next M51 (iterative).** Foundation + corpus re-ground + bring-up
hardening + content/seeding fill are done (M50's M42 gate MET both vantages, warm, on the strengthened manifest; rext
@ `fit-up-m50`). The **single-demo serialization** (fix-on-live across M51→M52, then **M53 destroys + cold-rebuilds**
as the single acceptance proof — which also owns M50's deferred COLD reset-to-seed clause) is in effect.

**Next up:** **build M51** (`/developer-kit:work-mstone-iters` — the AI-readiness showcase org toward the manager
coverage gate). _(The orchestrator still owes origin the pushes: `main` + the `v1.10` tag + the v1.10 ext tags + now
the `fit-up-m47..m50` rext tags — the v1.10 LOCAL close did not push; the M201 close merged to `main` LOCALLY; this
v1.10b branch is cut from that local `main`. The consumption-clone re-pin to `fit-up-m50` is the push-gated KEEP
tracked here, authoritatively bumped at M53.)_

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
- **M50 — Content & seeding fill** — **2026-06-30** (`iterative`, **closed-on-gate**; merged →
  `release/01.10b-fit-up`; rext tag `fit-up-m50` @ `f0d984c` + close fix-commit `3c2de27`). M42 coverage gate MET
  BOTH vantages (warm, strengthened manifest, frontier-exhausted, (0,0)). NEW `MemberLanguagesSeeder` (langs 0→747)
  + `CertificatesSeeder` member-coverage (9→236) + `UsersSeeder` member-field backfill + a `PUBLIC_WEBSITE_URL`
  demopatch + a NEW post-replay Directus content-URL rewrite + manager-manifest strengthening (preAssert/textMatch).
  AI-keys policy DECIDED: documented-as-absent. Tests rext stack-seeding 719 (97.4%) · Python 108; flake 0.
  Carry-forward: COLD acceptance → M53 (Fate-2, user-decided); academy F6 → M51 (Fate-3). Full narrative in
  `roadmap.md` § M50.
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

## Headline numbers (v1.10b — M50 close)
- **Go test funcs (rext):** at M50 close the M50-touched module **stack-seeding = 719** (`Test`+`Fuzz`; seeders pkg
  349, +1 at close — `TestNativeLanguageByCity_CoversLocations`; 97.4% stmt coverage on the seeders pkg, stable from
  harden Pass 2). M49-close per-module baseline (unchanged this milestone): `alignment` 52 · clerkenstein 270 ·
  stack-snapshot 364 · stack-secrets 163; stack-seeding was 706 at M49 (M50's 6 iters + 6 harden/iter test files +
  the close test → 719).
- **Python / TS:** the `demo-stack` suite **299** (incl. `test_demopatch` **49** [+2 vs 47: the M50 content-URL
  rewrite tests] · `test_frontend_build` **59** [+4 vs 55: the M50 content-URL-rewrite contract tests] · `test_tooling`
  111 · `test_cockpit` 63 · `test_ant_academy` 17); `stack-injection` 117. The `@playwright/test ^1.49.0` coverage
  harness (M42e) is the e2e foundation M202 reuses (the first non-Go rext dev/test dep).
- **Flake:** **0** (M50 close flake gate 5/5 shuffled-sequential on the seeders pkg; demo-stack Python 108 OK).
- **Supply-chain:** the v1.10 close carried **1 new dep** (`github.com/anthropos-work/ai@v1.40.1`, M45). v2.0 has
  added none yet. The rosetta corpus is docs-only (no package manifest). Lockfile inherited from
  [`releases/archive/01.10-method-acting/dependencies.lock`](releases/archive/01.10-method-acting/dependencies.lock).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces at v1.10 close — v2.0 touches no contract
  surface so far.

## Branch model
**v1.10b IN DEVELOPMENT (active):** `release/01.10b-fit-up` cut from `main` 2026-06-29 (LOCAL — origin push is the
orchestrator's step). Milestone branches `m{47..53}/{slug}` branch from it at build time. rext code of record (a
SEPARATE repo) is authored in the `.agentspace/rosetta-extensions/` copy (cloned in M47) + tagged `fit-up-m47..m52`
per the tooling policy (`fit-up-m47..m50` cut; M50 close fixes on rext `main` @ `3c2de27`, tag unmoved), rolled to
the `v1.10.1` release tag at M53; the consumed tag is pinned via the new `.agentspace/rext.tag` source-of-truth
(M49 #1).
**v2.0 PAUSED:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL). M201 merged → `main` (LOCAL, no
`v2.0` tag); M202→M204 not started — resumes after v1.10b. A `playthroughs` rext section arrives at M202 build.
**v1.10 SHIPPED:** `release/01.10-method-acting` merged `--no-ff` → `main` + tagged `v1.10` at close (LOCAL).
rext code @ tags `method-acting-m39..m46-servegrant-closure`.
**Shipped:** **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` ·
**v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.
(Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-06-30 (**M50 "Content & seeding fill" CLOSED-on-gate** via `/developer-kit:close-milestone` —
the iterative content fill; M42 coverage gate MET BOTH vantages [warm demo-1, the manifest STRENGTHENED to PROVE the
gaps, frontier-exhausted, (failingSections,escapes)=(0,0)]. NEW `MemberLanguagesSeeder` [langs 0→747 across all 340
members] + `CertificatesSeeder` member-coverage [9→236] + `UsersSeeder` member-field backfill + a
`next-web-public-website-url` demopatch [JS-constant escape] + a NEW post-replay Directus content-URL rewrite
[replayed-content escape] + manager-manifest strengthening [preAssert/textMatch]. rext tag `fit-up-m50` @ `f0d984c`
+ close fix-commit `3c2de27`; merged `--no-ff` → `release/01.10b-fit-up`. Close review: 11 findings all Fate-1 [gofmt
2 files + co-derivation comment + locations-sync test + 2 ResourceWarning handles + the NEW coverage-protocol
escape-class row + seeding-spec/secrets-spec lines + a broken M51 backref]; deferral audit **GREEN** [AI-keys policy
RESOLVED: documented-as-absent; COLD acceptance → M53 Fate-2 user-decided; academy F6 → M51 Fate-3; re-pin →
push-gated KEEP]. Tests rext stack-seeding 706→719 · demo-stack Python 299; flake 0. Active milestone now **M51**
[iterative AI-readiness showcase org]; next: `/developer-kit:work-mstone-iters`. Prior: 2026-06-30 **M49
"Bring-up hardening + truth-up" CLOSED** [the 7 demo-up issues, live-verified from cold; rext `fit-up-m49`; tests
1552→1555; AI-keys → M50]. Prior: 2026-06-29 **M48 "Corpus re-ground" CLOSED** [NEW `corpus/services/ai-readiness.md`
+ drift reconciled]; **M47 "Re-sync & recapture" CLOSED** [clones found current; `pg.NormalizeDSN`; snapshots
recaptured]. Prior:
2026-06-29 **v1.10b "fit-up" DESIGNED + PROMOTED** — the interposed **field-hardening backfill**; 7 milestones
M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53 [v1.x flat counter re-opened]; branch `release/01.10b-fit-up` cut from
`main`; tag `v1.10.1`. Re-grounds demo + corpus to current prod, fixes the demo-up issues + content gaps, adds the
AI-readiness showcase org [M51], consolidates one inlined seed+gen manifest [M52], cold-rebuild acceptance [M53].
Tooling + docs only. Prior: 2026-06-29 **M201 "Manifest corpus" CLOSED-on-gate** [9 products · 26 stories · 28
use-cases, user-signed-off]; **v2.0 PAUSED** for this backfill. Full v2.0-design history in `roadmap.md`.)_
