---
active_release: "v2.0 opening night (Playthroughs); M201 + M202 + M203 shipped, M204 (last) next"
active_branch: "release/02.00-opening-night"
active_milestone: "M204 — manager-vantage coverage (iterative; next + LAST v2.0 milestone; not yet started)"
last_closed: "M203 — 2026-07-02 (Employee-vantage coverage; closed-on-gate; tag opening-night-m203)"
phase: "v2.0 opening night building — M201 + M202 + M203 (employee vantage) shipped; M204 (manager, last) next"
last_updated: "2026-07-02"
---

# State

**Active release:** **v2.0 "opening night"** (the **Playthroughs** pillar; branch `release/02.00-opening-night`).
The interposed v1.10b field-hardening backfill SHIPPED (tag `v1.10.1`); v2.0 is the active release. Governed by
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md). Execution graph:
**M201 ✅ ∥ M202 ✅ → { M203 ∥ M204 } → ship.**

**Last shipped:** **v1.10b "fit-up" — 2026-07-01, tag `v1.10.1`** (rext code-of-record @ `66a021e`). An **interposed
field-hardening backfill** (the v1.3b "dress rehearsal" lineage): a from-scratch `/demo-up` surfaced 8 bring-up issues +
a tail of v1.10 content gaps. **7 milestones M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53**, all merged to
`release/01.10b-fit-up` → `main`. The clones were found **current** (M47 — a trivial `make pull`, not the reported
5-week lag); the genuinely-stale surface was the **rosetta corpus** (re-grounded in M48). The release: snapshot
recaptured from current prod (M47), corpus re-grounded (M48), the 8 bring-up issues + v1.10 content gaps fixed
(M49/M50), a curated **AI-readiness showcase org** added (M51), **one auditable seed+gen manifest** consolidated (M52),
and **cold-rebuild acceptance proven 6/6 + academy F6 GREEN from cold** (M53). **Tooling + docs only — zero
platform-repo edits. Zero new third-party deps.** Records: [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).

```
M47 ──→ ┌ M48 corpus re-ground ───────────┐                (M48 ∥ M49 — disjoint clusters; M48 no-demo)
        └ M49 bring-up hardening ──────────┘ ──→ M50 ──→ M51 ──→ M52 ──→ M53 ✅ shipped
```

**Active milestone:** **M204 — manager-vantage coverage** (`iterative`; the **LAST v2.0 milestone**; **not yet
started**). The Playthroughs foundation (M202) + the employee vantage (M203, closed-on-gate 2026-07-02, 6/6 GREEN on
cold reset-to-seed, 5/5 deterministic) are DONE. M204 proves **Dan's** manager journeys — Workforce funnel + member
roster, member drill-down (activity-dashboard), succession/at-risk (Growth tab) — same gate shape as M203,
manager-vantage. It **imports the shared page-object layer** (M202 base + M203's employee surfaces — an *additive*
merge) + **runs on the reset-to-seed lifecycle**, driven by the `corpus/ops/demo/playthroughs.md` protocol. Start
with **`/developer-kit:work-mstone-iters`**; the release ships when M204's gate fires. Records:
[`releases/02.00-opening-night/`](releases/02.00-opening-night/).

**Phase:** **v2.0 "opening night" building — M201 + M202 + M203 (employee vantage) shipped; M204 (manager, LAST) next.**
M203 closed-on-gate (rext playthroughs Go 103 + TS 38 unit + 6 browser Playthroughs, flake 0; 11 close findings all
Fate-1; tooling + docs only, zero platform edits, zero new deps). rext authoring @ `fb94458`, tagged `opening-night-m203`.

**Next up:** **iterate M204 — manager-vantage coverage** (`/developer-kit:work-mstone-iters`); it is the last v2.0
milestone → then `/developer-kit:close-release`. _(The orchestrator still owes origin the pushes — see the push-gated
KEEP below.)_

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + the `v1.10` tag + the v1.10 ext tags +
the `fit-up-m47..m52` rext tags + `v1.10.1` + the rext **`opening-night-m202`** + **`opening-night-m203`** tags (M202/
M203 closes). Local closes deliberately do not push; this is the user's gate. The box-level re-pin (consumption clone +
`.agentspace/rext.tag`) stays at `v1.10.1` — M203's `playthroughs` section runs from the authoring copy against demo-1
(a coverage milestone needs no consumption re-pin). An administrative KEEP, not a deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`), **M314b** (prod frozen-read whole-org hydration — a
prod-team follow-up, not tooling); **future v2 milestones** M205 Hiring + tier gates · M206 AI-sim mirror tier · M207
Academy coverage. All tracked in [`roadmap-vision.md`](roadmap-vision.md); none scheduled.

## Recently shipped releases
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Interposed field-hardening backfill: re-sync + recapture,
  corpus re-ground, from-cold `/demo-up` hardening, content/AI-readiness-org seeding fill, one auditable seed+gen
  manifest, then a from-cold destroy-and-rebuild acceptance (**6/6 + academy F6 GREEN**). 7 milestones M47..M53.
  Tooling + docs only, zero platform edits, zero new deps. Records:
  [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The believable-profile release + presenter-grade /
  scalable-generation extension; Playwright SEMANTIC coverage gate at both vantages cold; 9 milestones (M39→M46). The
  **last v1.x release** — detail in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. Declarative Stories & Heroes seeding + presenter cockpit.
  Records: [`releases/archive/01.90-storytelling/`](releases/archive/01.90-storytelling/).

_(Earlier v1.x — v1.0 … v1.8 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Recently closed (v2.0 milestones)
- **M203 — Employee-vantage coverage** — **2026-07-02** (`iterative`, closed-on-gate; rext tag `opening-night-m203`
  @ `fb94458`). 6/6 employee Playthroughs GREEN on cold reset-to-seed (Profile · Skill Paths · AI Sims chat launch),
  5/5 deterministic; 6 iters (1 tok + 5 tiks). Go 103 + TS 38 unit + 6 browser; 11 close findings all Fate-1; 4
  non-gate edge UCs → Fate-3 M206. Full narrative in `roadmap.md` § M203.
- **M202 — Playthroughs Foundation** — **2026-07-01** (`section`, closed-complete; rext tag `opening-night-m202` @
  `b1e5528`). The `playthroughs` rext section + the `corpus/ops/demo/playthroughs.md` runbook; proof Playthrough
  GREEN on demo-1. 96 Go + 13 TS, 8 close findings all Fate-1. Full narrative in `roadmap.md` § M202.
- **M201 — Manifest corpus** — **2026-06-29** (`iterative`, closed-on-gate). 9 products · 26 stories · 28 use-cases,
  user-signed-off. Full narrative in `roadmap.md` § M201.
_(v1.10b milestones M47–M53 shipped as tag `v1.10.1` — see "Recently shipped releases" above +
[`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).)_

## Headline numbers (v2.0 — M203 close)
- **Go test funcs (rext):** **1743** across 6 modules (`Test`+`Fuzz`; +7 vs M202's 1736 — playthroughs 96→103). By
  module: stack-seeding **791** · stack-snapshot **364** · clerkenstein 270 · stack-secrets **163** ·
  **playthroughs 103** (M203: 99 Test + 4 Fuzz across manifest/report/ptvalidate/ptreport) · alignment 52.
  `go vet ./...` clean; playthroughs 5/5 shuffled clean.
- **Python / TS:** `demo-stack` Python **326** (unchanged). rext **e2e TS unit** **80** (coverage-manifest 29 +
  section-assert 13 + **playthroughs 38** — stack-env 12 + url-shapes 26 route/landmark predicate cases). Plus **6
  browser Playthroughs** (playthroughs/e2e/tests) — 5/5 cold reset-to-seed deterministic on demo-1 (M203 iter-06).
- **Flake:** **0** (M203 close: playthroughs Go 5/5 -shuffle + TS unit 5/5; browser 5/5 cold-reset iter-06).
- **Supply-chain:** **0 new deps** this release (`ai v1.40.1` from v1.10 M45 carried forward unchanged; `go:embed` is
  stdlib). The inherited HIGH **CVE-2026-39821** (`x/net` idna, called only in stack-seeding, disclosed post-v1.10) was
  **CLEARED** at close (`x/net v0.53.0→v0.55.0`; `govulncheck` → "No vulnerabilities found"). node audit 0 vulns. No
  GPL/AGPL. Lockfile: [`releases/archive/01.10b-fit-up/dependencies.lock`](releases/archive/01.10b-fit-up/dependencies.lock).
- **Coverage:** no >2pp drop on any measured surface. NEW `playthroughs` section **98.5%** stmt (report 100 /
  manifest 99.4 / ptvalidate 97.6 / ptreport 94.8).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces (carried from v1.10; v1.10b touched no contract
  surface — clerkenstein unchanged).

## Branch model
**v1.10b SHIPPED:** `release/01.10b-fit-up` cut from `main` 2026-06-29, all 7 milestone branches `m{47..53}/{slug}`
merged `--no-ff` + deleted, then the release merged `--no-ff` → `main` + tagged `v1.10.1` at close (LOCAL — origin push
is the user's step). rext code of record (a SEPARATE repo) authored in `.agentspace/rosetta-extensions/`, tagged
`fit-up-m47..m52` per milestone + rolled to the **`v1.10.1`** release tag (annotated → commit `66a021e`); consumed via
`.agentspace/rext.tag` (= `v1.10.1`) + the `stack-demo/rosetta-extensions` consumption clone (pinned to `v1.10.1`).
**v2.0 IN DEVELOPMENT:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL). M201/M202/M203 merged
`--no-ff` → the release branch (LOCAL, no `v2.0` tag yet — tags at close-release); **M204 (last) not started**. rext
authoring @ `fb94458`, tagged `opening-night-m201..m203` per milestone; consumed via `.agentspace/rext.tag` (=
`v1.10.1`, the demo-1 pin — the playthroughs section runs from the authoring copy).
**Shipped tags:** **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` ·
**v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0**
`v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-02 (M203 "Employee-vantage coverage" CLOSED-ON-GATE — 6/6 employee Playthroughs GREEN on
cold reset-to-seed, 5/5 deterministic; 6 iters; 11 close findings all Fate-1; 4 non-gate UCs → Fate-3 M206; rext
tag `opening-night-m203` @ `fb94458`. Next + LAST: M204 manager-vantage → then `/developer-kit:close-release`.)_
