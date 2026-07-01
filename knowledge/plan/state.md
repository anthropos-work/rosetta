---
active_release: "v2.0 opening night (Playthroughs); M201 + M202 shipped, M203 ∥ M204 next"
active_branch: "release/02.00-opening-night"
active_milestone: "M203 ∥ M204 — vantage coverage (both iterative; next up, not yet started)"
last_closed: "M202 — 2026-07-01 (Playthroughs Foundation; tag opening-night-m202)"
phase: "v2.0 opening night building — M201 + M202 (foundation) shipped; M203 ∥ M204 next"
last_updated: "2026-07-01"
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

**Active milestone:** **M203 ∥ M204 — vantage coverage** (both `iterative`; **next up, not yet started**). The
Playthroughs **foundation is DONE** (M202 closed-complete 2026-07-01): the `playthroughs` rext section + the
`corpus/ops/demo/playthroughs.md` runbook, proof Playthrough GREEN on demo-1. M203 (employee-vantage) ∥ M204
(manager-vantage) both **import M202's page-object layer** + **run on its reset-to-seed lifecycle**, driven by the
`playthroughs.md` iteration protocol (their `iteration_protocol_ref`). Start with
**`/developer-kit:work-mstone-iters`**. Records: [`releases/02.00-opening-night/`](releases/02.00-opening-night/).

**Phase:** **v2.0 "opening night" building — M201 + M202 (foundation) shipped; M203 ∥ M204 (vantage coverage) next.**
M202 closed-complete (Go 96 test/fuzz @ 98.5% + TS 13, 5/5 flake-clean; 8 close findings all Fate-1; tooling + docs
only, zero platform edits). rext authoring @ `b1e5528`, tagged `opening-night-m202`.

**Next up:** **iterate M203 ∥ M204 — vantage coverage** (`/developer-kit:work-mstone-iters`). _(The orchestrator
still owes origin the pushes — see the push-gated KEEP below.)_

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + the `v1.10` tag + the v1.10 ext tags +
the `fit-up-m47..m52` rext tags + `v1.10.1` + the new rext **`opening-night-m202`** tag (M202 close). Local closes
deliberately do not push; this is the user's gate. The box-level re-pin (consumption clone + `.agentspace/rext.tag` →
`v1.10.1`) is DONE and stays at `v1.10.1` (M202's `playthroughs` section runs from the authoring copy against demo-1
— a foundation milestone needs no consumption re-pin). An administrative KEEP, not a deferral.

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
- **M202 — Playthroughs Foundation** — **2026-07-01** (`section`, closed-complete; rext tag `opening-night-m202` @
  `b1e5528`). The `playthroughs` rext section (manifest+validator / page-object layer / dedicated `pt-world` seed /
  reset-to-seed serial runner / 4-state report) + the `corpus/ops/demo/playthroughs.md` runbook; proof Playthrough
  GREEN on demo-1. 96 Go + 13 TS, 98.5% section, 8 close findings all Fate-1. Full narrative in `roadmap.md` § M202.
- **M201 — Manifest corpus** — **2026-06-29** (`iterative`, closed-on-gate). 9 products · 26 stories · 28 use-cases,
  adversarially re-grounded (11-agent wf → 15/27 runnable), user-signed-off. Full narrative in `roadmap.md` § M201.
_(v1.10b milestones M47–M53 shipped as tag `v1.10.1` — see "Recently shipped releases" above +
[`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).)_

## Headline numbers (v2.0 — M202 close)
- **Go test funcs (rext):** **1736** across 6 modules (`Test`+`Fuzz`; +96 vs v1.10b's 1640 — the NEW
  `playthroughs` module). By module: stack-seeding **791** · stack-snapshot **364** · clerkenstein 270 ·
  stack-secrets **163** · **playthroughs 96** (M202-NEW: 92 Test + 4 Fuzz across manifest/report/ptvalidate/ptreport)
  · alignment 52. `go vet ./...` clean; playthroughs 5/5 shuffled -race clean.
- **Python / TS:** `demo-stack` Python **326** (unchanged since v1.10b). rext **e2e TS unit** **54** (coverage-manifest
  29 + section-assert 13 + **playthroughs 12** M202-NEW: stack-env resolveStackEnv/resolveWorkers). Playthroughs TS
  5/5 clean; the proof Playthrough (`profile-identity.spec.ts`) GREEN on demo-1.
- **Flake:** **0** (M202 close: Go 5/5 shuffled -race + TS unit 5/5, clean per stack).
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
**v2.0 PAUSED:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL). M201 merged → `main` (LOCAL, no
`v2.0` tag); M202→M204 not started — resumes next. A `playthroughs` rext section arrives at M202 build.
**Shipped tags:** **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` ·
**v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0**
`v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-01 (v1.10b "fit-up" SHIPPED — 7 milestones M47..M53, tag `v1.10.1`, cold-rebuild-accepted
6/6+F6 GREEN; close-release full suites + 3× flake gate clean, CVE cleared. Next: v2.0 "opening night" resumes.)_
