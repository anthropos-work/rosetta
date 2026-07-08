---
active_release: "v2.1 quick change — IN DEVELOPMENT (branch release/02.10-quick-change; tag v2.1)"
active_branch: "release/02.10-quick-change"
active_milestone: "M211 — Bring-up acceptance (iterative, IN PROGRESS — PAUSED 2026-07-08 for user commute): 8 iters committed (iter-01..08); gate ~5.5/6 (5/6 proven warm + cold /demo-up GREEN); remaining: M42 coverage + v2.0 Playthroughs + cold /dev-up. Resume: /developer-kit:work-mstone-iters M211"
last_closed: "M210 — 2026-07-08 (corpus + skills re-ground — adopted the colleague's arch/subgraph/service half + flipped the 6 rext-facing tooling-doc bodies skiller.*→public.*; 0 stale skiller.<table> refs corpus-wide; docs-only 50 .md, HARDEN N/A; deferral audit GREEN, KB-1/2/3 resolved)"
phase: "M211 IN PROGRESS — PAUSED (user commute 2026-07-08) at run-2 iter-08. Iterative bring-up acceptance: TOK-01 warm-first-cache-migrate/cold-prove-both; the CACHE-MIGRATION workaround SUCCEEDED (real 42,790-row taxonomy cache re-keyed skiller.*→public.* at .agentspace/snapshots/taxonomy/5afc0bcc…; cold replay hits it); rext.tag→quick-change-m209; fixed dev casbin-policy-load gap; verify GREEN; 0 residual skiller; dev cold-path extensions-bootstrap prep landed; **cold /demo-up GREEN on merged platform** (demo half of cold headline). GATE ~5.5/6. REMAINING: (e) M42 coverage + v2.0 Playthroughs (vs the UP demo-1) + cold /dev-up. demo-1 stack UP (16 containers; `cd stack-demo/platform && make down` to free resources). 8 iters committed (durable). Resume: /developer-kit:work-mstone-iters M211."
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

**Active milestone:** **M211 — Bring-up acceptance: `dev-up` + `demo-up` green on the merged platform**
(`iterative`, `closed-on-gate`, `planned`) — the **FINAL v2.1 milestone**. Prove the whole chain works end-to-end
on the merged 4-subgraph platform with the re-grounded tooling. **Exit gate:** from a re-synced state, `/dev-up`
AND `/demo-up` both go **GREEN cold** — 4-subgraph compose / no skiller container; snapshot **recapture→replay**
loads `public.*` (taxonomy replay exits 0, ~42,763 public skills); **seed** resolves real public node-ids (closure
green); **verify** passes with a merged-platform assertion (no skiller schema/subgraph/container); the M42 coverage
sweep + the v2.0 Playthroughs suite stay GREEN; **0 residual skiller-schema references** in any queried path.
**Its first tik recaptures the `public.*` taxonomy** — the M209-deferred data op — via a sanctioned COPY-byte
source. **Next:** `/developer-kit:work-mstone-iters` M211 (then `/developer-kit:close-release`).

**Phase:** **M210 CLOSED** (merged → `release/02.10-quick-change`) — made the corpus internally consistent with the
merged platform + M209's landed `public.*` rext code. Adopted the colleague's correct architecture/subgraph/service
half (28 files, reconciled vs the M208 fact-sheet — no duplicate merge section); fixed the profile-completeness
node-id prose (**verified NO literal "43/44" exists** — did not fabricate a phantom count); flipped the **6
rext-facing tooling-doc bodies + directus-local** `skiller.*→public.*` and deleted the interim disclosure notes;
reconciled db-access ↔ tooling on `public.*`; swept the 4 skill files + `CLAUDE.md` to the verified merged compose
(no skiller container, **4 subgraphs**, `SKILLER_RPC_ADDR=http://backend:8083`; superseded the colleague's stale
exit-4 note). **0 stale `skiller.<table>` tooling-query refs corpus-wide**; docs-only (50 `.md`, 0 code/test) →
HARDEN N/A; close review 0 must-fix / 1 nice-to-have no-change-needed; deferral audit **GREEN** (KB-1/2/3 resolved).

**Next up:** **run `/developer-kit:work-mstone-iters`** for **M211** — the iterative bring-up acceptance closer.
Its exit gate stands both `/dev-up` and `/demo-up` GREEN cold on the merged platform; the first tik **recaptures**
the `public.*` taxonomy (the M209/M208-deferred data op) via a sanctioned COPY-byte source, then bring-up +
set-dress + seed + verify + the M42 coverage + v2.0 Playthroughs gates. M211 is the last v2.1 milestone; after it,
**`/developer-kit:close-release`** rolls the rext `v2.1` tag, bumps `.agentspace/rext.tag`, and merges → `main`.

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
stays at `v1.10.1` until close-release tags `v2.1` and re-pins the consumption stacks. An administrative KEEP, not
a deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), **DEF-M208-01 / M25-D9** (dev taxonomy `rc=4` — the `extensions`-schema bootstrap +
PG-readiness bring-up requirement; Fate-3 → M211), **DEF-M208-02** (`INVITATION_HMAC_SECRET` dev `.env` gap →
M211 / `/stack-secrets`), **rext `stack-seeding/README` test-count drift** (says 496 / 8 pkgs, actual ~788 / 13;
pre-existing since M41, cross-release — reconcile at the v2.1 rext roll / close-release; rext frozen at `2f06e78`),
M314b (prod frozen-read whole-org hydration — a prod-team follow-up). All tracked in
[`roadmap-vision.md`](roadmap-vision.md). The reserved **Playthroughs futures** (M205 Hiring/tier-gates · M206
AI-sim-mirror-tier + M203-carried edge UCs · M207 Academy) stay reserved in vision.

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
  +14 harden funcs on the two new non-mechanical risk items + a few build-phase matcher additions). M210 = docs-only
  (0 code/test) → no change.
- **Live Playthroughs:** **10** (6 employee + 4 manager) GREEN on cold reset-to-seed + 1 in-manifest TODO. v2.1
  M211 keeps this suite GREEN as a bring-up-acceptance gate on the merged platform.
- **Supply-chain:** **0 net-new deps** target for v2.1 (a schema re-point adds none). `ai v1.40.1` unchanged.
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces (carried from v1.10; v2.1 touches no Clerk
  contract surface — the skiller merge is a taxonomy-schema/RPC move, not a Clerk change).

## Branch model / shipped tags
**v2.1 IN DEVELOPMENT:** `release/02.10-quick-change` cut from `main` 2026-07-08. Milestones `m208/…` … `m211/…`
branch off it (strictly sequential); M208+M209+M210 **CLOSED** (merged into the release branch). rext authoring
copy @ `quick-change-m209` (`2f06e78`); the `v2.1` rext roll + consumption re-pin (`.agentspace/rext.tag` stays
`v1.10.1`) are close-release's job.
**Shipped tags:** **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` ·
**v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` ·
**v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-08 (M210 "Corpus + skills re-ground" CLOSED — corpus re-grounded to M209's landed `public.*`
code; adopted the colleague's arch/subgraph/service half + flipped the 6 rext-facing tooling-doc bodies + swept the
skill files/CLAUDE.md to the 4-subgraph/no-skiller compose; 0 stale `skiller.<table>` refs corpus-wide; docs-only
(50 `.md`, 0 code) → HARDEN N/A; 0 must-fix; deferral audit GREEN (KB-1/2/3 resolved, 4 defers confirm-only →
M211/close-release); merged → `release/02.10-quick-change`. Next: `/developer-kit:work-mstone-iters` M211.)_
