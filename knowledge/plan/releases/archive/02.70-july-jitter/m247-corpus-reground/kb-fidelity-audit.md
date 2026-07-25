---
title: "KB Fidelity Audit — M247 corpus re-ground"
date: 2026-07-23
scope: milestone:M247
invoked-by: build-milestone
---

## Verdict
YELLOW

Proceed with tracking. This is a **doc-only re-ground** milestone: the corpus IS the deliverable, and the
platform source (`stack-demo/app` @ current `origin/main`, the M246 clone-pin bump) is the ground truth the
milestone reads. Every stale corpus claim in scope is a claim the milestone REWRITES — none is a dependency the
implementation reads as truth (it reads the source). Every net-new fact-sheet blind area is already promoted to a
milestone deliverable under `overview.md §Delivers`. No RED trigger.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| skillpath merged-into-app | `corpus/services/skillpath.md` (to convert → redirect); pattern: `skiller.md` | `stack-demo/app/internal/skillpath{,s}`; `public.skill_path_sessions` (561 rows, M246); repos.yml (0 skillpath); compose (0 skillpath svc) | PAIRED (stale — milestone rewrites) |
| 3-subgraph reclassification | `architecture_overview.md`, `service_taxonomy.md`, `graphql-wundergraph.md`, `backend.md`, `cms.md`, `dependency_map.md`, `external_services.md`, ops docs, `CLAUDE.md` | `graphql-wundergraph/supergraph-config-{compose,prod}.yaml` = **3 subgraphs** (backend, jobsimulation, cms) | PAIRED (stale — milestone rewrites) |
| coursebuilder fact sheet | — (BLIND, already a `Delivers`) | `stack-demo/app/internal/coursebuilder/` (SPEC.md, README.md, ~many .go) | BLIND-AREA → deliverable |
| AI Labs + credits fact sheet | — (BLIND, already a `Delivers`) | `app/internal/labs/{adapter,catalog,labsapi,session}`, `app/internal/{credits,payments,subscriptions}`, `app/stripe/` | BLIND-AREA → deliverable |
| askengine (Talk-to-Data) fact sheet | — (BLIND, already a `Delivers`) | `app/internal/askengine/` (executor.go, prompt.go, registry.go, bedrock.go, followups.go) | BLIND-AREA → deliverable |
| academy-backend fact sheet | — (BLIND, already a `Delivers`) | `app/internal/academy/` (academy.go, asset.go, body.go, bookmark.go, certificate.go, …) | BLIND-AREA → deliverable |
| ai-readiness refresh | `corpus/services/ai-readiness.md` (exists) | `app/internal/aireadiness/` (cycles.go, defaults.go, diagnosis.go, compare.go, csv.go, emailoverride/) | PAIRED (refresh) |
| roadrunner resolution | `corpus/services/roadrunner.md` (exists), `README.md` | repos.yml (**present**, 10 repos), docker-compose.yml `roadrunner:` block :306 profile graphql | PAIRED (alive — resolve ORPHANED flag) |
| fact-sheet shape (KB dep) | `corpus/services/TEMPLATE.md` | n/a | DOC-ONLY (contract — verified accurate) |
| redirect pattern (KB dep) | `corpus/services/skiller.md` | n/a | DOC-ONLY (contract — verified accurate) |
| index / archived-merged table (KB dep) | `corpus/services/README.md`, `corpus/services/backend.md` | n/a | DOC-ONLY (contract — verified present) |

## Fidelity Findings

1. **skillpath-is-live claim — STALE, milestone deliverable.** The corpus asserts skillpath a live Tier-1 service /
   4th subgraph across ~30 files (M246 drift ledger D-01). Ground truth (M246, empirical + verified this audit):
   repos.yml lists 10 repos with **0 skillpath**; docker-compose has **no skillpath service**; the supergraph has
   **3 subgraphs** (backend/jobsimulation/cms); runtime skill-path session state moved to `public.skill_path_sessions`
   (561 seeded rows). **Fix owner: doc — this milestone.** Not read as truth by the implementation.

2. **"4 subgraphs" claim — STALE, milestone deliverable.** 7 files literally say "4 subgraphs" (`architecture_overview.md`,
   `external_services.md`, `skillpath.md`, `README.md`, `cms.md`, `backend.md`, `CLAUDE.md`); the true count is **3**
   (verified against `supergraph-config-compose.yaml` + `-prod.yaml`). **Fix owner: doc — this milestone.**

3. **Redirect pattern (`skiller.md`) — ALIGNED, accurate exemplar.** The `skiller.md` merged-into-app redirect
   (⚠️ banner → where domain/RPC/GraphQL/infra/repo went → still-true domain knowledge → related docs) is a clean,
   current template for the skillpath redirect. Its own claims (skiller merged, `public` schema, subgraph removed)
   match the consolidated source. No fix needed.

4. **`TEMPLATE.md` — ALIGNED.** The fact-sheet shape (Role & Responsibility / Architecture & Code Map / Interface
   Discovery / Local Development / Testing) is intact; the 4 net-new fact sheets follow it.

5. **4 net-new domains — BLIND-AREA, resolved by promotion.** `coursebuilder`, AI Labs+credits, `askengine`,
   `academy` have **no** current corpus coverage. All four are already listed under `overview.md §Delivers` (the
   skill-sanctioned resolution: "add a `Delivers → knowledge/{path}` line"). Source confirmed present + substantial
   for each. **Not a RED blocker** — they are the milestone's deliverables.

6. **roadrunner "ORPHANED" flag — STALE-toward-alive (to resolve).** Section 8 says "resolve roadrunner
   ORPHANED→ARCHIVED **if dead**." Ground truth: roadrunner is **in repos.yml (10 repos) + docker-compose
   (`roadrunner:` service, profile graphql)** → **alive, not dead.** The resolution is the *negative*: confirm live,
   retire any stale ORPHANED framing. Verified this audit; resolved in build.

7. **ai-readiness doc PAIRED — refresh scope.** `corpus/services/ai-readiness.md` exists; `app/internal/aireadiness/`
   is the current (refactored into its own package) home. The refresh reconciles the aireadiness-package refactor.
   Not a blind area.

## Completeness Gaps

None load-bearing that block the build. The 4 blind-area domains are the completeness gaps this milestone exists to
close; each has confirmed source to read (`stack-demo/app/internal/{coursebuilder,labs,askengine,academy}` + the
credits/payments/subscriptions/stripe surface for AI Labs' shared purse).

## Applied Fixes

None inline (this is a pre-build audit of a milestone whose entire body is the fix). The stale claims + blind areas
are the milestone's Phase-1 work, not audit-time doc edits.

## Open Items (require user decision)

None. All findings route to this milestone's own sections or to already-planned siblings:
- rext-file drift (D-02/03/04 gen_injected_override/test_injection/exposure_claim_guard; D-06 up-injected.sh prose)
  is **out of M247's doc-only scope** — routes to M251 (test-health) / a rext-hygiene follow-up. Recorded in
  `decisions.md`, not an M247 blocker.

## Gate Result

**YELLOW — proceed with tracking.** No blind area is unpromoted; no stale claim is read as truth by the
implementation (it reads current source). Findings recorded as KB-1..KB-7 in `decisions.md`. Proceed to Phase 1.
