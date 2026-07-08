---
active_release: "v2.1 quick change — IN DEVELOPMENT (branch release/02.10-quick-change; tag v2.1)"
active_branch: "release/02.10-quick-change"
active_milestone: "M210 — Corpus + skills re-ground (section, planned) — flip the rext-facing tooling-doc bodies to public.* in lockstep with M209's landed schema"
last_closed: "M209 — 2026-07-08 (rext tooling re-ground — skiller.*→public.* across snapshot+seeding+small modules; 6 Go modules GREEN, 0 skiller.<table> queries; rext tagged quick-change-m209@2f06e78; recapture Fate-3→M211)"
phase: "M209 CLOSED (merged → release/02.10-quick-change): rext re-grounded skiller.*→public.*, digest narrowed, MinRows floor, tagged quick-change-m209@2f06e78 (14 harden funcs, 0 bugs, 0 flakes). Recapture Fate-3→M211. M210 BUILT (m210/corpus-reground, 6 sections, 0 stale skiller.<table>). Next: /developer-kit:close-milestone M210"
last_updated: "2026-07-08"
---

# State

**Active release:** **v2.1 "quick change" — IN DEVELOPMENT.** The **skiller-in-app re-ground** — a field-hardening
release (v1.3b "dress rehearsal" / v1.10b "fit-up" lineage) triggered by a **landed platform structural change**:
the `skiller` service + its DB schema merged into `app` (domain → the **`public`** schema, table names unchanged
`skiller.X → public.X`; RPC → `backend`; the skiller GraphQL subgraph gone → **4 subgraphs**; skiller
repo/container removed). Designed 2026-07-08 via `/developer-kit:design-roadmap`. Branch
`release/02.10-quick-change` cut from `main`; tag `v2.1`; rext tag `v2.1`. **4 milestones M208 → M209 → M210 →
M211, strictly sequential** (the user's execution choice). **Tooling + docs + stack-re-sync only — zero
platform-repo edits** (the platform already did its half). Detail:
[`roadmap.md`](roadmap.md) § In Development — v2.1.

**Active milestone:** **M210 — Corpus + skills re-ground** (`section`, `planned`) — land the colleague's
`origin/docs/skiller-in-app-merge` sweep as the complete, internally-consistent corpus re-ground and **flip the
5-6 rext-facing tooling-doc bodies to `public.*` in lockstep with M209's landed schema** (`snapshot-spec.md`,
`seeding-spec.md`, `safety.md` firewall row, `stories-spec.md`, `recipe-snapshot-world.md`, `coverage-protocol.md`),
fix the missed `profile-completeness-spec.md`, sweep the skill files + `CLAUDE.md` (5→4 subgraphs, container counts).
**Next:** `/developer-kit:build-milestone` M210.

**Phase:** **M209 CLOSED** (merged → `release/02.10-quick-change`) — rext re-grounded `skiller.*→public.*` across
stack-snapshot (the `taxonomy.go` const flip re-grounding capture+replay; the Risk-1 `Surface.VersionTables()`
digest-narrowing; a one-sided `MinRows` under-capture floor; Risk-2 verified names-only, no column change),
stack-seeding (24 files, `organization_id IS NULL` preserved, `data-dna.json` golden, `isolation.go`), and the
small shell modules (5→4 services; stack-verify Python 104/104). 6 Go modules GREEN, `go vet` clean, 5× flake-clean,
**0 `skiller.<table>` queries in any production path**; rext tagged `quick-change-m209@2f06e78` (14 harden funcs, 0
bugs, 0 flakes). Close: 1 nice-to-have finding (pre-existing rext README count drift, routed), deferral audit GREEN.

**Next up:** **run `/developer-kit:build-milestone`** for **M210** (the corpus body-flip, lockstep with M209's
landed schema). Then M211 (iterative bring-up acceptance: `/dev-up` + `/demo-up` GREEN cold on the merged platform;
its first tik **recaptures** the `public.*` taxonomy — the M209-deferred data op — via a sanctioned COPY-byte source).

**Design inputs / evidence:** the user's skiller-merge briefing + the colleague's unmerged
`origin/docs/skiller-in-app-merge` corpus sweep (correct-but-incomplete). A 7-agent research workflow
(`wf_08b6bf4a`) mapped the per-module blast radius, adversarially confirmed the snapshot firewall public-predicate
(`organization_id IS NULL`) **survives** the merge (no data-leak risk), and confirmed the docs branch cannot land
present-tense before the rext re-ground + stack re-sync. The two non-obvious risks it surfaced (folded into M209):
the **cache-key digest regression** (post-merge `SchemaVersionSQL` digests the whole app monolith → taxonomy cache
thrash — narrow the digest) and the **capture column-mapping** (`embedding→small_embedding3`, `extensions.`-opclasses
— verify vs merged-prod).

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + tags `v1.10.1` + `v2.0` + the rext
tags. Local closes deliberately do not push; this is the user's gate. The box-level re-pin (`.agentspace/rext.tag`)
stays at `v1.10.1` until M209 tags `v2.1` and re-pins the consumption stacks. An administrative KEEP, not a deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), **M25-D9** (dev taxonomy `rc=4` — **surfaced at M208 on the clean-slate re-migrate as
the `extensions`-schema-bootstrap + PG-readiness bring-up requirement; did NOT fall out as a trivial Fate-1 →
routed Fate-3 to M211**), **rext `stack-seeding/README` test-count drift** (says 496 / 8 pkgs, actual ~788 / 13;
pre-existing since M41, cross-release — reconcile at the v2.1 rext roll / next rext re-tag; rext was frozen at
`2f06e78` for the M209 close — see M209 `decisions.md` D-close-2), M314b (prod frozen-read whole-org hydration —
a prod-team follow-up). All tracked in [`roadmap-vision.md`](roadmap-vision.md). The reserved **Playthroughs futures** (M205 Hiring/tier-gates · M206
AI-sim-mirror-tier + M203-carried edge UCs · M207 Academy) stay reserved in vision — v2.1 takes M208+ per the
established "reserved-number-ships-later" precedent (M206 is a live Fate-3 destination from the M203 close).

## Recently shipped releases
- **v2.0 "opening night"** — **2026-07-02**, tag `v2.0`. The **Playthroughs** pillar: manifest corpus (M201) →
  foundation (M202) → employee (M203) + manager (M204) coverage. **10 live Playthroughs GREEN on cold reset-to-seed**
  + 1 in-manifest TODO. 4 milestones M201..M204. **The first v2.x release.** Records:
  [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Interposed field-hardening backfill: re-sync + recapture,
  corpus re-ground, from-cold `/demo-up` hardening, content/AI-readiness-org seeding fill, one auditable seed+gen
  manifest, then a from-cold destroy-and-rebuild acceptance. 7 milestones M47..M53. Records:
  [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The believable-profile release + presenter-grade /
  scalable-generation extension; Playwright SEMANTIC coverage gate at both vantages cold; 9 milestones (M39→M46).
  The **last v1.x release** — detail in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).

_(Earlier v1.x — v1.0 … v1.9 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Headline numbers (inherited from v2.0 — final; v2.1 baseline)
- **rext Go test funcs:** **1763** across 6 modules (playthroughs the 6th). `go vet ./...` clean. — v2.1 M209
  re-pointed the seeding/snapshot tests (net +18: the ~111 `skiller.*→public.*` fake-Conn matcher renames are flat;
  +14 harden funcs on the two new non-mechanical risk items + a few build-phase matcher additions).
- **Live Playthroughs:** **10** (6 employee + 4 manager) GREEN on cold reset-to-seed + 1 in-manifest TODO. v2.1
  M211 keeps this suite GREEN as a bring-up-acceptance gate on the merged platform.
- **Supply-chain:** **0 net-new deps** target for v2.1 (a schema re-point adds none). `ai v1.40.1` unchanged.
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces (carried from v1.10; v2.1 touches no Clerk
  contract surface — the skiller merge is a taxonomy-schema/RPC move, not a Clerk change).

## Branch model / shipped tags
**v2.1 IN DEVELOPMENT:** `release/02.10-quick-change` cut from `main` 2026-07-08. Milestones `m208/…`, `m209/…`,
`m210/…`, `m211/…` branch off it (strictly sequential). rext authoring copy currently @ `v2.0`; M209 rolls it to
`v2.1`. Consumption pin (`.agentspace/rext.tag`) stays `v1.10.1` until M209.
**Shipped tags:** **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` ·
**v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` ·
**v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-08 (M209 "rext tooling re-ground" CLOSED — rext re-grounded `skiller.*→public.*` across
snapshot+seeding+small modules, 6 Go modules GREEN + `go vet` clean + 5× flake-clean, 0 `skiller.<table>` queries in
production; rext tagged `quick-change-m209@2f06e78` (14 harden funcs, 0 bugs); merged → `release/02.10-quick-change`.
1 nice-to-have finding (pre-existing rext README drift, routed), deferral audit GREEN. Next: `/developer-kit:build-milestone` M210.)_
