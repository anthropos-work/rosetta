---
active_release: "(between releases) — v2.1 quick change SHIPPED (tag v2.1); awaiting /developer-kit:design-roadmap"
active_branch: "main"
active_milestone: "(between releases)"
last_closed: "v2.1 quick change — 2026-07-09 (tag v2.1, 4 milestones M208..M211) — the skiller-in-app re-ground"
phase: "between releases — awaiting /developer-kit:design-roadmap"
last_updated: "2026-07-09"
---

# State

**Active release:** **(between releases).** **v2.1 "quick change"** — the **skiller-in-app re-ground** — SHIPPED
2026-07-09 (tag `v2.1`). No release is active; the next v2.x release awaits `/developer-kit:design-roadmap`.

**Last shipped:** **v2.1 "quick change" — 2026-07-09, tag `v2.1`** (rext code-of-record rolled to `v2.1` =
`quick-change-m211`; `.agentspace/rext.tag` bumped `v1.10.1` → `v2.1`). The **skiller-in-app re-ground**: the landed
platform merge (skiller service + DB schema → `app`'s `public` schema, table names unchanged; RPC → `backend`; the
skiller GraphQL subgraph gone → **4 subgraphs**; skiller repo/container removed) re-fit across the rext tooling +
corpus + both local stacks, and **`/dev-up` + `/demo-up` proven to work cold on the merged platform**. **4
milestones M208 → M209 → M210 → M211, strictly sequential**, all merged `--no-ff` → `release/02.10-quick-change` →
`main`. Zero platform-repo edits, 0 net-new deps. Records:
[`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/).

**Active milestone:** **(between releases).** All four v2.1 milestones are closed + merged. No milestone is active.

**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.** v2.1 shipped clean (close-release: both
blocking gates GREEN + adversarially verified [deferral 0 Fate-3-undelivered; metrics 1745→1764, flake 0]; 0
must-fix; all should-fix landed — the 42,790 count reconcile + the `go1.25.11→go1.25.12` toolchain bump [govulncheck
clean] + the doc-flip fidelity; triple-clean 3/3; user sign-off on the clean release incl. the dev-cold DB-init
interpretation).

**Next up:** **run `/developer-kit:design-roadmap`** to open the next release. Reserved Playthroughs-futures
candidates in `roadmap-vision.md`: **M205** Hiring + tier gates · **M206** AI-sim mirror tier + the M203 carried edge
UCs · **M207** Academy coverage.

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + tags `v1.10.1` + `v2.0` + **`v2.1`** +
the rext tags (incl. the new rext `v2.1` roll + `quick-change-m208..m211`). Local closes deliberately do not push;
this is the user's gate. `.agentspace/rext.tag` is now `v2.1` (the box-level consumption re-pin). An administrative
KEEP, not a deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), **M25-D9 — RESOLVED (v2.1 M211** via `dev-stack/migrate-dev.sh`**)**, **CAVEAT-1**
(clean-box literal full `/dev-up` — belt-and-suspenders; the dev-cold gate delta was proven at the DB-init level on a
non-destructive throwaway to protect the box's native-app content-line dev), M314b (prod frozen-read hydration — a
prod-team follow-up). All tracked in [`roadmap-vision.md`](roadmap-vision.md).

## Recently shipped releases
- **v2.1 "quick change"** — **2026-07-09**, tag `v2.1`. The **skiller-in-app re-ground**: re-sync + live de-risk
  (M208) → rext re-point `skiller.*→public.*` + recapture (M209) → corpus flip (M210) → merged-platform cold
  bring-up proven, gate 6/6 (M211). M42 coverage both vantages + v2.0 Playthroughs 10/11 GREEN on the merged
  platform. 4 milestones M208..M211. Records:
  [`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/).
- **v2.0 "opening night"** — **2026-07-02**, tag `v2.0`. The **Playthroughs** pillar: manifest corpus (M201) →
  foundation (M202) → employee (M203) + manager (M204) coverage. 10 live Playthroughs GREEN on cold reset-to-seed.
  The first v2.x release. Records: [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Interposed field-hardening backfill: re-sync + recapture,
  corpus re-ground, from-cold `/demo-up` hardening, seeding fill, one auditable manifest. 7 milestones M47..M53.
  Records: [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).

_(Earlier v1.x — v1.0 … v1.10 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Headline numbers (v2.1 — final)
- **rext Go test funcs:** **1764** across 6 modules (+19 vs v2.0's 1745; every module flat-or-up). `go vet ./...`
  clean; triple-clean 3/3 (-shuffle, `go1.25.12`).
- **rext TS unit specs:** **103** (+3 vs v2.0's 100). **Live Playthroughs:** **10** (6 employee + 4 manager) +
  1 in-manifest TODO — re-proven GREEN on the merged platform.
- **Flake:** **0**. **Coverage:** <2pp delta (rename-not-removal + added harden tests).
- **Supply-chain:** **0 net-new deps**. Go toolchain `go1.25.11 → go1.25.12` (cleared 2 inherited stdlib advisories;
  govulncheck clean all 6). npm 0 vulns; 0 GPL/AGPL. Lockfile:
  [`releases/archive/02.10-quick-change/dependencies.lock`](releases/archive/02.10-quick-change/dependencies.lock).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces (v2.1 touched no Clerk contract surface).

## Branch model / shipped tags
**v2.1 SHIPPED:** `release/02.10-quick-change` cut from `main` 2026-07-08 → merged `--no-ff` → `main` + tagged
`v2.1` at close (LOCAL — origin push is the user's step). M208/M209/M210/M211 all merged `--no-ff` → the release
branch. rext code-of-record rolled to **`v2.1`** (= `quick-change-m211`); per-milestone tags `quick-change-m208..m211`.
Consumption pin (`.agentspace/rext.tag`) bumped `v1.10.1` → **`v2.1`**.
**Shipped tags:** **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` ·
**v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` ·
**v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-09 (v2.1 "quick change" SHIPPED — tag `v2.1`, 4 milestones M208..M211; the skiller-in-app
re-ground, merged platform proven cold on both stacks; close-release GREEN, both gates verified. Next:
`/developer-kit:design-roadmap`.)_
