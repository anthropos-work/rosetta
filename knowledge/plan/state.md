---
active_release: "v2.1 quick change — IN DEVELOPMENT (branch release/02.10-quick-change; tag v2.1)"
active_branch: "release/02.10-quick-change"
active_milestone: "M208 — Re-sync & merged-schema ground-truth (section, planned) — next up"
last_closed: "v2.0 opening night — 2026-07-02 (tag v2.0, 4 milestones M201..M204) — the Playthroughs pillar"
phase: "designed — release branch cut + milestones scaffolded; awaiting /developer-kit:build-milestone (M208)"
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

**Active milestone:** **M208 — Re-sync & merged-schema ground-truth** (`section`, `planned`) — the load-bearing
foundation; everything downstream grades against the current merged code it produces. **Next:**
`/developer-kit:build-milestone` M208.

**Phase:** **designed** — release branch cut + `m208…m211` milestone dirs scaffolded; awaiting
`/developer-kit:build-milestone`.

**Next up:** **run `/developer-kit:build-milestone`** to build **M208** (`make pull` both stacks to the merged
platform, remove the vestigial `stack-*/skiller/` clones, re-migrate against `public`, pin the merge fact-sheet).
Then M209 (rext re-ground + recapture) → M210 (corpus re-ground, lockstep with M209) → M211 (iterative bring-up
acceptance: `/dev-up` + `/demo-up` GREEN cold on the merged platform).

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
(`replayCmd` hermetic test), **M25-D9** (dev taxonomy `rc=4` — **now on the M208 re-sync migration path →
opportunistic Fate-1**), M314b (prod frozen-read whole-org hydration — a prod-team follow-up). All tracked in
[`roadmap-vision.md`](roadmap-vision.md). The reserved **Playthroughs futures** (M205 Hiring/tier-gates · M206
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
- **rext Go test funcs:** **1745** across 6 modules (playthroughs the 6th). `go vet ./...` clean. — v2.1 M209 will
  re-point the seeding/snapshot tests (net count roughly flat; a lockstep rename, not new surface).
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

_Last updated: 2026-07-08 (v2.1 "quick change" DESIGNED + PROMOTED — the skiller-in-app re-ground; 4 milestones
M208 → M209 → M210 → M211, strictly sequential; branch `release/02.10-quick-change` cut from `main`, tag `v2.1`.
Next: `/developer-kit:build-milestone` M208.)_
