---
active_release: "v1.10b fit-up (interposed backfill; v2.0 opening night PAUSED)"
active_branch: "release/01.10b-fit-up"
active_milestone: "M52 — Single auditable seed+gen manifest (section; next to build)"
last_closed: "M51 — 2026-07-01 (AI-readiness showcase org, closed-on-gate)"
phase: "v1.10b building — M47..M51 CLOSED; next M52 (single auditable seed+gen manifest, section)"
last_updated: "2026-07-01"
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

**Active milestone:** **M52 — Single auditable seed+gen manifest** (`section`, **next to build**). Extract the Go
mother-prompts to YAML; author **one** checked-in `seed-generation-manifest.yaml` inlining the population blueprint
(all 3 orgs, incl. the M51 AI-readiness org) + prompts + batch config + snapshot sources (**cache + generated data
excluded**); the cockpit **[Download]** serves it. `section` — a fixed extract-and-inline scope. **Delivers →** NEW
`corpus/ops/demo/seed-manifest-spec.md`. Records:
[`releases/01.10b-fit-up/m52-seed-manifest/`](releases/01.10b-fit-up/m52-seed-manifest/). _(M51 closed-on-gate
2026-07-01 — see Recently closed.)_

**Phase:** **v1.10b building — M47..M51 CLOSED; next M52 (section).** Foundation + corpus re-ground + bring-up
hardening + content/seeding fill + the AI-readiness showcase org are done (M51's manager coverage gate MET, `(0,0)`
frontier-exhausted; rext @ `fit-up-m51`). The **single-demo serialization** (fix-on-live across M52, then **M53
destroys + cold-rebuilds** as the single acceptance proof — which also owns M50's deferred COLD reset-to-seed clause +
now M51's academy-F6 Fate-3 handoff) is in effect.

**Next up:** **build M52** (`/developer-kit:work-milestone` — extract the mother-prompts + inline the one manifest).
_(The orchestrator still owes origin the pushes: `main` + the `v1.10` tag + the v1.10 ext tags + now the
`fit-up-m47..m51` rext tags — the v1.10 LOCAL close did not push; the M201 close merged to `main` LOCALLY; this
v1.10b branch is cut from that local `main`. The consumption-clone re-pin to the release rext tag is the push-gated
KEEP tracked here, authoritatively bumped at M53.)_

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
- **M51 — AI-readiness showcase org** — **2026-07-01** (`iterative`, **closed-on-gate**; merged →
  `release/01.10b-fit-up`; rext tag `fit-up-m51` @ `a23f38d` + close fix-commit `1e958ac`). Manager coverage gate
  MET at iter-09: `(failingSections, escapes) = (0,0)` frontier-exhausted (reachable 70) on a fresh demo-up; org
  **Northwind Aviation** (200) ENABLED, **78.4%** all-3-complete, Ben STARTED + Aria COMPLETED, cycle `closed` + 199
  frozen snapshots. **9 iters** (1 bootstrap tok + 8 tiks). 3 net-new seeders (`OrgSettingsSeeder` +
  `AIReadinessConfigSeeder` + `AIReadinessFunnelSeeder`) + the `app-aireadiness-snapshot-loadmembers` read-path
  demo-patch (the iter-09 TOK-02 pivot bounding an unbounded whole-org `loadMembers` → the frozen `?cycle=` GET
  180s→19ms). Tests rext stack-seeding **719→749** (seeders pkg 97.6%) · e2e TS unit **20→33**; flake 0. Close: 16
  findings all Fate-1 (the 3rd AI-readiness story authored → `delivers` MET). Deferral audit **RED→CLEARED**:
  academy F6 fated LAND-NEXT → M53 (Fate-3, user-decided). Carry-forward: F6 + COLD acceptance + re-pin → M53. Full
  narrative in `roadmap.md` § M51.
- **M50 — Content & seeding fill** — **2026-06-30** (`iterative`, **closed-on-gate**; merged →
  `release/01.10b-fit-up`; rext tag `fit-up-m50` @ `f0d984c` + close fix-commit `3c2de27`). M42 coverage gate MET
  BOTH vantages (warm, strengthened manifest, frontier-exhausted, (0,0)). NEW `MemberLanguagesSeeder` (langs 0→747)
  + `CertificatesSeeder` member-coverage (9→236) + `UsersSeeder` member-field backfill + a `PUBLIC_WEBSITE_URL`
  demopatch + a NEW post-replay Directus content-URL rewrite + manager-manifest strengthening (preAssert/textMatch).
  AI-keys policy DECIDED: documented-as-absent. Tests rext stack-seeding 719 (97.4%) · Python 108; flake 0.
  Carry-forward: COLD acceptance → M53 (Fate-2, user-decided); academy F6 → M51 (Fate-3, **re-routed onward to M53
  at M51 close**). Full narrative in `roadmap.md` § M50.
- **M49 — Bring-up hardening + truth-up** — **2026-06-30** (`section`; merged → `release/01.10b-fit-up`; rext tag
  `fit-up-m49` @ `ba586d6`). Closed the 7 remaining demo-up issues (rext.tag SoT + CRLF-tolerant reader, .env-guard
  order, `INVITATION_HMAC_SECRET` critical+auto-gen, ant-academy explicit clone, disk pre-flight + `down --purge`,
  true non-fatal frontend, demopatch re-anchor to v2.89.0) — **live-verified** from cold (demo-1 UP). Tests rext Go
  1552→1555 · demo-stack Python 299; flake 0. AI-keys policy → M50 (Fate-2). Full narrative in `roadmap.md` § M49.
- **M48 — Corpus re-ground** — **2026-06-29** (`section`; merged → `release/01.10b-fit-up`). NEW
  `corpus/services/ai-readiness.md` (the M51 seeder contract — Phase-2c-sharpened: active⇒signals-true,
  closed⇒frozen-snapshot) + drift reconciled + the false ant-academy "in repos.yml" claim corrected (M49 #5 owns
  the code fix). Docs-only; 3-agent investigation. Full narrative in `roadmap.md` § M48.
_(M47 "Re-sync & recapture" closed 2026-06-29 — trimmed from this list; full narrative in `roadmap.md` § M47.)_

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

## Headline numbers (v1.10b — M51 close)
- **Go test funcs (rext):** at M51 close the M51-touched module **stack-seeding = 749** (`Test`+`Fuzz`; seeders pkg
  382, **97.6%** stmt coverage — +30 vs M50's 719 across M51's 9 iters + the 5-pass final harden + close's +2
  [`TestAIReadinessSimRefs_CoDerivationInvariant` + `TestBuildRegistry_AIReadinessConfigPrecedesFunnelInPlan`]).
  M49-close per-module baseline (unchanged this milestone): `alignment` 52 · clerkenstein 270 · stack-snapshot 364 ·
  stack-secrets 163.
- **Python / TS:** the `demo-stack` Python suite **299** (unchanged this milestone — M51 touched no Python). The rext
  **e2e TS unit** suite **33** (+13 at M51 close: the NEW `section-assert.ts` no-browser verdict spec, joining the
  20-test `coverage-manifest.unit.spec.ts`) — runs via the `@playwright/test` harness with no demo up. `stack-injection`
  117.
- **Flake:** **0** (M51 close flake gate 5/5 Go [seeders + cmd/stackseed, count=1, sequential] + 5/5 TS [both unit
  specs, 33 passed each]).
- **Supply-chain:** the v1.10 close carried **1 new dep** (`github.com/anthropos-work/ai@v1.40.1`, M45). v2.0 has
  added none yet. The rosetta corpus is docs-only (no package manifest). Lockfile inherited from
  [`releases/archive/01.10-method-acting/dependencies.lock`](releases/archive/01.10-method-acting/dependencies.lock).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces at v1.10 close — v2.0 touches no contract
  surface so far.

## Branch model
**v1.10b IN DEVELOPMENT (active):** `release/01.10b-fit-up` cut from `main` 2026-06-29 (LOCAL — origin push is the
orchestrator's step). Milestone branches `m{47..53}/{slug}` branch from it at build time. rext code of record (a
SEPARATE repo) is authored in the `.agentspace/rosetta-extensions/` copy (cloned in M47) + tagged `fit-up-m47..m52`
per the tooling policy (`fit-up-m47..m51` cut; M51 close fixes on rext `main` @ `1e958ac`, tag `fit-up-m51` @
`a23f38d` unmoved — the M50 tag-at-gate precedent: tag at the gate, test/fix commits on `main`), rolled to the
`v1.10.1` release tag at M53; the consumed tag is pinned via the new `.agentspace/rext.tag` source-of-truth (M49 #1).
**v2.0 PAUSED:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL). M201 merged → `main` (LOCAL, no
`v2.0` tag); M202→M204 not started — resumes after v1.10b. A `playthroughs` rext section arrives at M202 build.
**v1.10 SHIPPED:** `release/01.10-method-acting` merged `--no-ff` → `main` + tagged `v1.10` at close (LOCAL).
rext code @ tags `method-acting-m39..m46-servegrant-closure`.
**Shipped:** **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` ·
**v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.
(Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-01 (**M51 "AI-readiness showcase org" CLOSED-on-gate** via `/developer-kit:close-milestone` —
the iterative AI-readiness showcase; the M42 manager coverage gate MET at iter-09 [`(failingSections,escapes)=(0,0)`
frontier-exhausted on a fresh demo-up, reachable 70; org Northwind 200 ENABLED, 78.4% all-3-complete, Ben STARTED +
Aria COMPLETED, cycle `closed` + 199 frozen snapshots]. 9 iters [1 bootstrap tok + 8 tiks]; 3 net-new seeders +
the `app-aireadiness-snapshot-loadmembers` read-path demo-patch [the iter-09 TOK-02 pivot after 3 falsified
read-fast strategies — "frozen SCORES ≠ frozen RESPONSE"]. rext tag `fit-up-m51` @ `a23f38d` + close fix-commit
`1e958ac`; merged `--no-ff` → `release/01.10b-fit-up`. Close review: 16 findings all Fate-1 [the 3rd AI-readiness
story authored → `delivers` MET; C1/C3/C4 code + T1/T2 tests + D1–D5 docs]; deferral audit **RED→CLEARED** [academy
F6 repeat-defer fated LAND-NEXT → M53 Fate-3, user-decided; COLD acceptance + re-pin → M53]. Tests rext
stack-seeding 719→749 [seeders pkg 97.6%] · e2e TS unit 20→33; flake 0. Active milestone now **M52** [section, single
auditable seed+gen manifest]; next: `/developer-kit:work-milestone`. Prior closes [detail in `roadmap.md` §§ M47–M50]:
M50 "Content & seeding fill" CLOSED-on-gate 2026-06-30, M49/M48/M47 CLOSED 2026-06-30/29. Prior: 2026-06-29 **v1.10b
"fit-up" DESIGNED + PROMOTED** — the interposed **field-hardening backfill**; 7 milestones M47 → { M48 ∥ M49 } → M50 →
M51 → M52 → M53 [v1.x flat counter re-opened]; branch `release/01.10b-fit-up` cut from `main`; tag `v1.10.1`. Prior:
2026-06-29 **M201 "Manifest corpus" CLOSED-on-gate**; **v2.0 PAUSED** for this backfill. Full history in `roadmap.md`.)_
