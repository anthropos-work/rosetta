---
title: "KB Fidelity Audit — M249 cross-app navigation"
date: 2026-07-23
scope: milestone:M249
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| demopatch mechanism (anchor→replace ladder, 7 guards, manifest schema, patch inventory) | `corpus/ops/demo/demopatch-spec.md` | `rext demo-stack/patches/{demopatch,manifest_loader.py}`, `stack-injection/apply-*.sh` | PAIRED |
| frontend image build (next-web + studio-desk + hiring + ant-academy offset-URL bake) | `corpus/ops/demo/frontend-tier.md` | `rext demo-stack/up-injected.sh` (`build_frontend_*`), `ant-academy.sh` | PAIRED |
| presenter cockpit (return-nav target; 7700+offset) | `corpus/ops/demo/cockpit-spec.md` | `rext demo-stack/cockpit.py` | PAIRED (return-nav SUBSECTION is BLIND — planned Delivers) |
| studio-desk service (offset-URL / prod-eject fix) | `corpus/services/studio-desk.md` | `stack-demo/studio-desk/app/core/scaffold/{userProfile.js,pageWrapper.js}` | PAIRED (prod-eject subsection is a planned Delivers update) |
| additive-UI injection pattern (inject a NEW menu element via demopatch) | — | (new manifests A/B/C) | BLIND-AREA — planned Delivers (`overview.md` Scope→In + Delivers) |

## Fidelity Findings

1. **demopatch-spec §5 patch inventory (16, breakdown `10 next-web-app · 2 app · 4 ant-academy`)** — ALIGNED with `demo-stack/patches/` (16 manifest dirs) and `test_patch_inventory.py`. **0 studio-desk source patches today** → confirms M249's "first-ever studio-desk source demopatch" framing. M249 will move this to 21 (`11 next-web-app · 2 app · 4+1 ant-academy · 3 studio-desk`) and update §5 + the fence together (in-scope, Delivers).
2. **frontend-tier.md:277 — "studio-desk … no source patch"** — ALIGNED as a description of the *current* state; M249 makes it STALE and updates it in the same milestone (Delivers → frontend-tier.md). Not a contract the implementation reads as truth; it is a doc M249 delivers an update to. **Fix owner: doc (this milestone).**
3. **cockpit-spec.md — 7700+offset port + deep-link contract present; NO return-nav section** — the return-nav subsection is BLIND but is an explicit Delivers (`overview.md`). **Fix owner: doc (this milestone, Docs section).**
4. **demopatch-spec §4 apply vehicles** — the studio-desk patch is a NEW image-baked `demopatch`-tool vehicle (target `stack-demo/studio-desk/…` is inside `DEMO_WS` → G1/G6 pass; applied in `build_frontend_studio_desk` before the docker build, reverted after — same class as next-web, NOT a native-run helper). ALIGNED; §4 will gain the studio-desk row.

## Completeness Gaps

1. `build_frontend_studio_desk` has **no patch-set fingerprint** today (only an endpoint check) — M249 adds one (reusing `next_web_patchset_fp`) + a `demo.patchset` label so cached studio images rebuild when a patch changes (the §5-bis "applying is not shipping" hazard). Documented as part of the frontend-tier.md / demopatch-spec.md updates. **In scope.**

## Applied Fixes
None applied inline — every finding is a planned Delivers of M249 (fixing them now would pre-empt the milestone's own Docs section). Triples recorded in `spec-notes.md`.

## Open Items (require user decision)
None.

## Gate Result
YELLOW: proceed with tracking. No blind area is unplanned (both the additive-UI injection pattern doc and the cockpit return-nav section are explicit `Delivers` in `overview.md`); no stale claim is a contract the implementation reads as truth (frontend-tier.md:277 is a doc M249 updates, not a spec it consumes). Findings tracked as `KB-1..KB-4` in `decisions.md`, addressed in the Docs section (Phase 5).
