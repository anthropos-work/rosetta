---
active_release: "v2.0 opening night (Playthroughs); ALL milestones (M201→M204) closed — ready for close-release"
active_branch: "release/02.00-opening-night"
active_milestone: "(between releases)"
last_closed: "M204 — 2026-07-02 (Manager-vantage coverage; closed-on-gate; tag opening-night-m204)"
phase: "between releases — v2.0 opening night COMPLETE; ready for /developer-kit:close-release"
last_updated: "2026-07-02"
---

# State

**Active release:** **v2.0 "opening night"** (the **Playthroughs** pillar; branch `release/02.00-opening-night`).
The interposed v1.10b field-hardening backfill SHIPPED (tag `v1.10.1`); v2.0 is the active release. Governed by
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md). Execution graph — **ALL CLOSED:**
**M201 ✅ ∥ M202 ✅ → { M203 ✅ ∥ M204 ✅ } → ship.**

**Last shipped:** **v1.10b "fit-up" — 2026-07-01, tag `v1.10.1`** (rext code-of-record @ `66a021e`). An **interposed
field-hardening backfill** (the v1.3b "dress rehearsal" lineage): a from-scratch `/demo-up` surfaced 8 bring-up issues +
a tail of v1.10 content gaps. **7 milestones M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53**, all merged to
`release/01.10b-fit-up` → `main`. Tooling + docs only — zero platform-repo edits. Zero new third-party deps.
Records: [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).

**Active milestone:** **(between releases).** All four v2.0 milestones are closed. M201 (manifest corpus) + M202
(Playthroughs foundation) + M203 (employee vantage) + **M204 (manager vantage, closed-on-gate 2026-07-02 — 4/4
manager Playthroughs GREEN on cold reset-to-seed, 5/5 deterministic)** are DONE and merged `--no-ff` →
`release/02.00-opening-night`. The corpus stands at **10 live Playthroughs, 1 TODO** (the assign-WRITE half, a
Fate-2 tracked manifest gap). No milestone is active. Records:
[`releases/02.00-opening-night/`](releases/02.00-opening-night/).

**Phase:** **between releases — v2.0 "opening night" COMPLETE; ready for `/developer-kit:close-release`.** All v2.0
milestones closed (M204 rext playthroughs Go 105 + TS 58 unit + 4 browser Playthroughs, flake 0; 5 close findings
all Fate-1; tooling + docs only, zero platform edits, zero new deps). rext authoring @ `c81c6dd`, tagged
`opening-night-m204`.

**Next up:** **run `/developer-kit:close-release`** — the release-level review + merge `release/02.00-opening-night`
→ `main` + tag `v2.0`. _(The orchestrator still owes origin the pushes — see the push-gated KEEP below.)_

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + the `v1.10` tag + the v1.10 ext tags +
the `fit-up-m47..m52` rext tags + `v1.10.1` + the rext **`opening-night-m201..m204`** tags (all v2.0 milestone
closes). Local closes deliberately do not push; this is the user's gate. The box-level re-pin (consumption clone +
`.agentspace/rext.tag`) stays at `v1.10.1` — the `playthroughs` section runs from the authoring copy against demo-1
(a coverage milestone needs no consumption re-pin). An administrative KEEP, not a deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`), **M314b** (prod frozen-read whole-org hydration — a
prod-team follow-up, not tooling); **future v2 milestones** M205 Hiring + tier gates · M206 AI-sim mirror tier · M207
Academy coverage. All tracked in [`roadmap-vision.md`](roadmap-vision.md); none scheduled.

## Recently shipped releases
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Interposed field-hardening backfill: re-sync + recapture,
  corpus re-ground, from-cold `/demo-up` hardening, content/AI-readiness-org seeding fill, one auditable seed+gen
  manifest, then a from-cold destroy-and-rebuild acceptance (**6/6 + academy F6 GREEN**). 7 milestones M47..M53.
  Records: [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The believable-profile release + presenter-grade /
  scalable-generation extension; Playwright SEMANTIC coverage gate at both vantages cold; 9 milestones (M39→M46). The
  **last v1.x release** — detail in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. Declarative Stories & Heroes seeding + presenter cockpit.
  Records: [`releases/archive/01.90-storytelling/`](releases/archive/01.90-storytelling/).

_(Earlier v1.x — v1.0 … v1.8 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Recently closed (v2.0 milestones)
- **M204 — Manager-vantage coverage** — **2026-07-02** (`iterative`, closed-on-gate; rext tag `opening-night-m204`
  @ `c81c6dd`). 4/4 manager Playthroughs GREEN on cold reset-to-seed (Workforce funnel+roster · activity-dashboard
  drill-down · succession/at-risk), 5/5 deterministic; 5 iters (1 tok + 4 tiks). Go 105 + TS 58 unit + 4 browser; 5
  close findings all Fate-1; the assign-WRITE UC1 → Fate-2 (in-manifest TODO). Full narrative in `roadmap.md` § M204.
- **M203 — Employee-vantage coverage** — **2026-07-02** (`iterative`, closed-on-gate; rext tag `opening-night-m203`
  @ `fb94458`). 6/6 employee Playthroughs GREEN on cold reset-to-seed, 5/5 deterministic; 6 iters. Go 103 + TS 38 +
  6 browser; 11 close findings all Fate-1; 4 non-gate edge UCs → Fate-3 M206. Full narrative in `roadmap.md` § M203.
- **M202 — Playthroughs Foundation** — **2026-07-01** (`section`, closed-complete; rext tag `opening-night-m202` @
  `b1e5528`). The `playthroughs` rext section + the `corpus/ops/demo/playthroughs.md` runbook; proof Playthrough
  GREEN on demo-1. 96 Go + 13 TS, 8 close findings all Fate-1. Full narrative in `roadmap.md` § M202.
- **M201 — Manifest corpus** — **2026-06-29** (`iterative`, closed-on-gate). 9 products · 26 stories · 28 use-cases,
  user-signed-off. Full narrative in `roadmap.md` § M201.

## Headline numbers (v2.0 — M204 close)
- **Go test funcs (rext):** **1745** across 6 modules (`Test`+`Fuzz`; +2 vs M203's 1743 — playthroughs 103→105). By
  module: stack-seeding **791** · stack-snapshot **364** · clerkenstein 270 · stack-secrets **163** ·
  **playthroughs 105** (M204: 101 Test + 4 Fuzz across manifest/report/ptvalidate/ptreport) · alignment 52.
  `go vet ./...` clean; playthroughs 5/5 shuffled clean.
- **Python / TS:** `demo-stack` Python **326** (unchanged). rext **e2e TS unit** **100** (coverage-manifest 29 +
  section-assert 13 + **playthroughs 58** — stack-env 12 + url-shapes 46 route/landmark predicate cases). Plus **10
  browser Playthroughs** (6 employee + 4 manager) — 5/5 cold reset-to-seed deterministic (M203 iter-06 + M204 iter-05).
- **Flake:** **0** (M204 close: playthroughs Go 5/5 -shuffle + TS unit 5/5; browser 4/4 manager 5/5 cold-reset iter-05).
- **Supply-chain:** **0 new deps** this release (`ai v1.40.1` from v1.10 M45 carried forward unchanged). node audit 0
  vulns. No GPL/AGPL. Lockfile: [`releases/archive/01.10b-fit-up/dependencies.lock`](releases/archive/01.10b-fit-up/dependencies.lock).
- **Coverage:** no >2pp drop on any measured surface. `playthroughs` manifest **100%** stmt (M204 harden held 100%
  across all 3 passes — invariant-pinning, not %-gains). TS manager url-shapes predicates ~100% via url-shapes.unit.
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces (carried from v1.10; v2.0 touched no contract surface).

## Branch model
**v1.10b SHIPPED:** `release/01.10b-fit-up` → `main` + tagged `v1.10.1` at close (LOCAL — origin push is the user's
step). rext tagged `fit-up-m47..m52` + rolled to **`v1.10.1`** (annotated → commit `66a021e`); consumed via
`.agentspace/rext.tag` (= `v1.10.1`) + the `stack-demo/rosetta-extensions` consumption clone.
**v2.0 — ALL MILESTONES CLOSED, ready for close-release:** `release/02.00-opening-night` cut from `main` 2026-06-28
(LOCAL). M201/M202/M203/**M204** all merged `--no-ff` → the release branch (LOCAL, no `v2.0` tag yet — tags at
`/developer-kit:close-release`). rext authoring @ `c81c6dd`, tagged `opening-night-m201..m204` per milestone;
consumed via `.agentspace/rext.tag` (= `v1.10.1`, the demo-1 pin — the playthroughs section runs from the authoring copy).
**Shipped tags:** **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` ·
**v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0**
`v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-02 (M204 "Manager-vantage coverage" CLOSED-ON-GATE — 4/4 manager Playthroughs GREEN on cold
reset-to-seed, 5/5 deterministic; 5 iters; 5 close findings all Fate-1; assign-WRITE UC1 → Fate-2 in-manifest; rext
tag `opening-night-m204` @ `c81c6dd`. v2.0 COMPLETE — ALL milestones closed; next: `/developer-kit:close-release`.)_
