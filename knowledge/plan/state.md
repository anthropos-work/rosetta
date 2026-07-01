---
active_release: "v1.10b fit-up — SHIPPED 2026-07-01 (tag v1.10.1). v2.0 opening night PAUSED — resumes next (design-roadmap/user assigns)"
active_branch: "main"
active_milestone: "(between releases)"
last_closed: "v1.10b fit-up — 2026-07-01 (tag v1.10.1, 7 milestones M47..M53)"
phase: "between releases — v1.10b shipped; v2.0 opening night resumes next"
last_updated: "2026-07-01"
---

# State

**Last shipped:** **v1.10b "fit-up" — 2026-07-01, tag `v1.10.1`** (rext code-of-record @ `66a021e`). An **interposed
field-hardening backfill** (the v1.3b "dress rehearsal" lineage): a from-scratch `/demo-up` surfaced 8 bring-up issues +
a tail of v1.10 content gaps. **7 milestones M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53**, all merged to
`release/01.10b-fit-up` → `main`. The clones were found **current** (M47 — a trivial `make pull`, not the reported
5-week lag); the genuinely-stale surface was the **rosetta corpus** (re-grounded in M48). The release: snapshot
recaptured from current prod (M47), corpus re-grounded (M48), the 8 bring-up issues + v1.10 content gaps fixed
(M49/M50), a curated **AI-readiness showcase org** added (M51), **one auditable seed+gen manifest** consolidated (M52),
and **cold-rebuild acceptance proven 6/6 + academy F6 GREEN from cold** (M53). **Tooling + docs only — zero
platform-repo edits. Zero new third-party deps.** Records: [`releases/01.10b-fit-up/`](releases/01.10b-fit-up/).

```
M47 ──→ ┌ M48 corpus re-ground ───────────┐                (M48 ∥ M49 — disjoint clusters; M48 no-demo)
        └ M49 bring-up hardening ──────────┘ ──→ M50 ──→ M51 ──→ M52 ──→ M53 ✅ shipped
```

**Active milestone:** **(between releases).** v1.10b shipped 2026-07-01; there is no active milestone.

**Phase:** **between releases — v1.10b shipped; v2.0 "opening night" resumes next.** v1.10b's close-release ran to
completion: full suites green + a 3× flake gate clean (rext Go **1640** / demo-stack Python **326** / TS unit **42**,
flake **0**), the inherited HIGH `x/net` CVE cleared, `release/01.10b-fit-up` merged → `main` + tagged `v1.10.1`.

**Next up:** **resume v2.0 "opening night" (Playthroughs).** It was paused 2026-06-29 after M201 closed, to interpose
this backfill. The M201 corpus is preserved as the v2.0 spec; M202 ∥ M203 ∥ M204 are not started. `design-roadmap` /
the user assigns the next active release + milestone. _(The orchestrator still owes origin the pushes — see the
push-gated KEEP below.)_

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + the `v1.10` tag + the v1.10 ext tags +
the `fit-up-m47..m52` rext tags + `v1.10.1`. Local closes deliberately do not push; this is the user's gate. The
box-level re-pin (consumption clone + `.agentspace/rext.tag` → `v1.10.1`) is DONE. An administrative KEEP, not a
deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`), **M314b** (prod frozen-read whole-org hydration — a
prod-team follow-up, not tooling); **future v2 milestones** M205 Hiring + tier gates · M206 AI-sim mirror tier · M207
Academy coverage. All tracked in [`roadmap-vision.md`](roadmap-vision.md); none scheduled.

## Recently shipped releases
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Interposed field-hardening backfill: re-sync + recapture,
  corpus re-ground, from-cold `/demo-up` hardening, content/AI-readiness-org seeding fill, one auditable seed+gen
  manifest, then a from-cold destroy-and-rebuild acceptance (**6/6 + academy F6 GREEN**). 7 milestones M47..M53.
  Tooling + docs only, zero platform edits, zero new deps. Records:
  [`releases/01.10b-fit-up/`](releases/01.10b-fit-up/).
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The believable-profile release + presenter-grade /
  scalable-generation extension; Playwright SEMANTIC coverage gate at both vantages cold; 9 milestones (M39→M46). The
  **last v1.x release** — detail in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. Declarative Stories & Heroes seeding + presenter cockpit.
  Records: [`releases/archive/01.90-storytelling/`](releases/archive/01.90-storytelling/).

_(Earlier v1.x — v1.0 … v1.8 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Recently closed (v1.10b milestones)
- **M53 — Cold-rebuild acceptance** — **2026-07-01** (`section`; the FINAL v1.10b milestone; rext release tag
  `v1.10.1`). demo-1 purged + cold-rebuilt from `v1.10.1` by a single `/demo-up`; **6/6 acceptance criteria + academy
  F6 GREEN from cold.** AB4 surfaced + fixed an M51-owned gate regression at the gate (org-conditional manager
  manifest, `117fe41`). Full narrative in `roadmap.md` § M53.
- **M52 — Single auditable seed+gen manifest** — **2026-07-01** (`section`; rext `fit-up-m52` @ `36d7430`). ONE
  checked-in `seed-generation-manifest.yaml` drives all seed+gen intent, projected + honesty-gated from the presets,
  served by the cockpit [Download]. Full narrative in `roadmap.md` § M52.
- **M51 — AI-readiness showcase org** — **2026-07-01** (`iterative`, closed-on-gate; rext `fit-up-m51` @ `a23f38d`).
  Manager coverage gate MET at iter-09 (Northwind 200, 78.4% all-3-complete, closed cycle + 199 frozen snapshots). 3
  net-new seeders + the `app-aireadiness-snapshot-loadmembers` read-path demo-patch (the perf-saga fix). Full narrative
  in `roadmap.md` § M51.
- **M50 — Content & seeding fill** — **2026-06-30** (`iterative`, closed-on-gate; rext `fit-up-m50` @ `f0d984c`). M42
  coverage gate MET both vantages; new member-language/certificate/user-field seeders + Directus content-URL rewrite.
  Full narrative in `roadmap.md` § M50.
_(M49/M48/M47 closed 2026-06-30/29 — full narratives in `roadmap.md` §§ M47–M49.)_

## Headline numbers (v1.10b — release close)
- **Go test funcs (rext):** **1640** across 5 modules (`Test`+`Fuzz`; +89 vs v1.10's 1551, no decrease). By module:
  stack-seeding **791** · stack-snapshot **364** · stack-secrets **163** · alignment 52 · clerkenstein 270 (the last
  two untouched all release). `go vet ./...` clean; all modules green in a 3× shuffled run.
- **Python / TS:** `demo-stack` Python **326** (whole-dir; every touched suite grew, none decreased). rext **e2e TS
  unit** **42** (coverage-manifest 29 + section-assert 13). Both 3/3 clean at close.
- **Flake:** **0** (per-milestone 5/5 shuffled gates + a release-close **3× shuffled** gate, 3/3 clean per stack).
- **Supply-chain:** **0 new deps** this release (`ai v1.40.1` from v1.10 M45 carried forward unchanged; `go:embed` is
  stdlib). The inherited HIGH **CVE-2026-39821** (`x/net` idna, called only in stack-seeding, disclosed post-v1.10) was
  **CLEARED** at close (`x/net v0.53.0→v0.55.0`; `govulncheck` → "No vulnerabilities found"). node audit 0 vulns. No
  GPL/AGPL. Lockfile: [`releases/01.10b-fit-up/dependencies.lock`](releases/01.10b-fit-up/dependencies.lock).
- **Coverage:** no >2pp drop on any measured surface (seeders ~97.5→97.6%; NEW M52 manifest pkg 100% stmt).
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
