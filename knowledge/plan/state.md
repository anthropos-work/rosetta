---
active_release: "v2.1 quick change — ALL MILESTONES CLOSED, awaiting close-release (branch release/02.10-quick-change; tag v2.1)"
active_branch: "release/02.10-quick-change"
active_milestone: "(between milestones — v2.1 COMPLETE: all 4 milestones M208 → M209 → M210 → M211 closed + merged to the release branch. Next: /developer-kit:close-release)"
last_closed: "M211 — 2026-07-08 (bring-up acceptance, iterative closed-on-gate — gate 6/6; the merged skiller-in-app platform stands up cold on BOTH stacks via the re-grounded tooling; cache-migration 42,790 taxonomy + build-scratch-resync root-fix + migrate-dev.sh cold DB-init; M42 coverage both vantages + Playthroughs 10/11 GREEN; 0 residual skiller; deferral audit GREEN)"
phase: "v2.1 COMPLETE — M211 closed-on-gate 2026-07-08; all 4 milestones (M208 re-sync → M209 rext re-ground → M210 corpus re-ground → M211 bring-up acceptance) merged to release/02.10-quick-change. The release is proven from cold. Next: /developer-kit:close-release (rolls the rext v2.1 tag, bumps .agentspace/rext.tag from v1.10.1, reconciles the rext READMEs [TEST-1 + DOC-1], merges → main, tags v2.1)."
last_updated: "2026-07-08"
---

# State

**Active release:** **v2.1 "quick change" — ALL MILESTONES CLOSED, awaiting `/developer-kit:close-release`.** The
**skiller-in-app re-ground** — a field-hardening release (v1.3b "dress rehearsal" / v1.10b "fit-up" lineage)
triggered by a **landed platform structural change**: the `skiller` service + its DB schema merged into `app`
(domain → the **`public`** schema, table names unchanged `skiller.X → public.X`; RPC → `backend`; the skiller
GraphQL subgraph gone → **4 subgraphs**; skiller repo/container removed). Designed 2026-07-08 via
`/developer-kit:design-roadmap`. Branch `release/02.10-quick-change` cut from `main`; tag `v2.1`; rext tag `v2.1`.
**4 milestones M208 → M209 → M210 → M211, strictly sequential — ALL `done`.** **Tooling + docs + stack-re-sync
only — zero platform-repo edits** (the platform already did its half). Detail:
[`roadmap.md`](roadmap.md) § In Development — v2.1.

**Active milestone:** **(between milestones).** v2.1 is **complete** — all four milestones closed + merged to the
release branch. M211 (the final milestone) closed-on-gate 2026-07-08 with **gate 6/6 MET**: the merged
(skiller-in-app) 4-subgraph platform stands up **end-to-end, cold, on BOTH stacks** via the re-grounded tooling,
with zero platform-repo edits. There is no active milestone; the release awaits close-release.

**Phase:** **M211 CLOSED-ON-GATE** (merged → `release/02.10-quick-change`) — proved the whole chain works
end-to-end on the merged platform. Delivered (all in the rext tooling, tag `quick-change-m211` = `2039103`): the
**cache-migration** (real 42,790-row taxonomy + 274 sims re-keyed `skiller.*→public.*`, replayed — the
no-prod-access recapture); the **root-cause fix** (the injected build-scratch was pinned pre-merge + survived
`--purge` → stale rebuilds → federation `_entities(Skill.name)` 422s; fixed to re-sync the scratch every bring-up);
the dev casbin `init_policy.sql` load; frontend offset-reuse guard; demo-local `ACADEMY_URL` bake + academy-aware
cross-port hook; demopatch URL re-pin (next-web v2.106.1); the Playthroughs reset-to-seed **roster-refresh**; and
the new **`dev-stack/migrate-dev.sh`** (dev cold DB-init, mirror of `migrate-demo.sh`, the M25-D9 path). Result:
cold `/demo-up` GREEN end-to-end; **M42 coverage GREEN both vantages**; **v2.0 Playthroughs 10/11 GREEN** (1
declared in-manifest TODO); dev cold DB-init cold-verified; 0 residual skiller. 17 iters + 4 stabilized harden
passes; rosetta close diff docs+plan-only → HARDEN N/A; close review 1 finding (DOC-1 → Fate-2 close-release);
deferral audit GREEN.

**Next up:** **run `/developer-kit:close-release`** — the last milestone is closed, so the release-level review +
merge is next. It treats all v2.1 commits (M208..M211) as one PR, runs the release-level quality sweep, then:
**rolls the rext `v2.1` release tag** (rolling up `quick-change-m208..m211`; the authoring copy is at
`quick-change-m211` = `2039103`), **bumps `.agentspace/rext.tag`** from `v1.10.1` → `v2.1` (the box-level
consumption re-pin), **reconciles the two rext-README items** (TEST-1 test-count drift + DOC-1 `migrate-dev.sh`
index — both land at the rext code-of-record roll), and **merges `release/02.10-quick-change` → `main` + tags
`v2.1`**. Origin push stays the user's manual gate (push-gated KEEP).

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + tags `v1.10.1` + `v2.0` + the rext
tags. Local closes deliberately do not push; this is the user's gate. An administrative KEEP, not a deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), **M25-D9 — RESOLVED (v2.1 M211)** via `dev-stack/migrate-dev.sh` (dev cold DB-init),
**CAVEAT-1** (clean-box literal full destructive `/dev-up` — belt-and-suspenders; the gate was met via an
environment-respecting interpretation on the user's native-content-line box), M314b (prod frozen-read whole-org
hydration — a prod-team follow-up). **Routed to close-release's rext roll:** **TEST-1** (rext
`stack-seeding/README` test-count drift, ~788/13 actual vs 496/8 quoted; pre-existing since M41) + **DOC-1** (rext
`dev-stack/README` should index `migrate-dev.sh`). **Confirmed-covered:** DEF-M208-02 (`INVITATION_HMAC_SECRET`
dev `.env`) by `/stack-secrets`; PT-TODO (assign-WRITE Playthrough half) by the reserved manager-write tier. All
tracked in [`roadmap-vision.md`](roadmap-vision.md).

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

## Headline numbers (inherited from v2.0; v2.1 baseline — M211 added rext harness tests, code-of-record @ quick-change-m211)
- **rext Go test funcs:** **~1764** across 6 modules — v2.1 M209 re-pointed the seeding/snapshot tests (net +18);
  M211 added a `playthroughs/manifest` RosterRefreshGate; M211's other harness additions are Python/TS/shell
  (migrate-dev live docker 4, TS coverage-manifest +3, next-web unreadable-reuse edge +1). `go vet ./...` clean.
- **Live Playthroughs:** **10** (6 employee + 4 manager) GREEN on cold reset-to-seed + 1 in-manifest TODO — M211
  re-proved this suite GREEN on the merged platform (a bring-up-acceptance gate).
- **Supply-chain:** **0 net-new deps** in v2.1 (a schema re-point adds none). `ai v1.40.1` unchanged.
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces (carried from v1.10; v2.1 touches no Clerk
  contract surface).

## Branch model / shipped tags
**v2.1 ALL MILESTONES CLOSED:** `release/02.10-quick-change` cut from `main` 2026-07-08. Milestones `m208/…` …
`m211/…` all **CLOSED** (merged into the release branch); the milestone branches deleted. rext authoring copy @
`quick-change-m211` (`2039103`). The `v2.1` rext roll + consumption re-pin (`.agentspace/rext.tag` stays
`v1.10.1`) + merge → `main` are close-release's job.
**Shipped tags:** **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` ·
**v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` ·
**v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-08 (M211 "Bring-up acceptance" CLOSED-ON-GATE — the merged platform stands up cold on both
stacks via the re-grounded tooling; gate 6/6; cache-migration + build-scratch-resync root-fix + migrate-dev.sh;
M42 coverage both vantages + Playthroughs 10/11 GREEN; 0 residual skiller; deferral audit GREEN; merged →
release/02.10-quick-change; the milestone branch deleted. v2.1 is COMPLETE → /developer-kit:close-release.)_
