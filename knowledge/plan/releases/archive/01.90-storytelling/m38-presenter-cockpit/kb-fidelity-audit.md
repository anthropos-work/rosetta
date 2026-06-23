---
title: "KB Fidelity Audit — M38 (Presenter cockpit)"
date: 2026-06-23
scope: milestone:M38
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Clerkenstein multi-identity consumer (registry + roster JSON + handshake `?__clerk_identity` + `/v1/demo/{identities,select}`) | `clerkenstein/knowledge/architecture.md` § Multi-identity (M37) | `clerkenstein/clerk-frontend/{registry,server}.go`, `cmd/fake-fapi/main.go` | PAIRED |
| Seeder id-derivation (the exact clerk-claim contract per hero the roster producer must single-source) | M37 `decisions.md` § ARCH + `stories-spec.md` § Stories & Heroes (M35) | `stack-seeding/seeders/{users,identity,helpers,userprofile,persona}.go`, `blueprint/stories.go` | PAIRED |
| Stories & Heroes blueprint (`stack.stories.yaml` — the single source the cockpit + seeder both read) | `corpus/ops/demo/stories-spec.md` § Stories & Heroes (M35) | `stack-seeding/blueprint/blueprint.go` (Persona: vantage/trajectory/annotation/login/jump_to), `presets/stories.seed.yaml` | PAIRED |
| Demo bring-up FAPI wiring (where `FAKE_FAPI_ROSTER` + the cockpit get launched) | `corpus/ops/demo/README.md` + `frontend-tier.md` (demo UI tier) | `demo-stack/up-injected.sh`, `stack-injection/gen_injected_override.py` (fake-fapi service block) | PAIRED |
| The presenter cockpit (panel + [Login as] + [Jump to section]) | `corpus/ops/demo/stories-spec.md` (flagged "M37–M38; this doc is the foundation those build on") | — (this milestone delivers it) | DOC-ONLY (expected) |
| The deep-link catalog (O9 — next-web routes per vantage) | — (this milestone delivers it) | next-web `jump_to` values in `presets/stories.seed.yaml` (`/profile`, `/enterprise/workforce?tab=…`) | DOC-ONLY (expected) |

## Fidelity Findings

1. **Roster JSON contract** — Doc (`architecture.md` § Multi-identity) says `cmd/fake-fapi` loads a `RosterEntry` JSON from `FAKE_FAPI_ROSTER`, the FAPI is single-`DefaultDemoUser()` without it. **Actual:** `registry.go` defines `Roster{Identities []RosterEntry}` with snake_case fields (`key,auth_id,eid,email,firstname,lastname,org_auth_id,org_eid,org_role`), `DisallowUnknownFields`, first entry = default seat; `cmd/fake-fapi/main.go` `buildServer` falls back to `DefaultDemoUser()` on empty path, fails loud on a set-but-bad path. **Verdict: ALIGNED.**
2. **Seat-switch mechanisms** — Doc says two: `?__clerk_identity=<key>` on the handshake (in-flow) + `GET /v1/demo/identities` / `POST /v1/demo/select` (control plane). **Actual:** `server.go` registers all three routes; `handleHandshake` selects the seat from `__clerk_identity` then redirects to `redirect_url` with the handshake token. **Verdict: ALIGNED** (and load-bearing: the cockpit's combined [Login as]+[Jump to] is one handshake redirect with both `__clerk_identity` and `redirect_url`).
3. **Seeder id-contract** — M37 `decisions.md` § ARCH states `Eid=deterministicUUID("<prefix>:user:<idx>")`, `AuthID="user_seed_<slug(prefix)>_<idx>"`, `Email=emailFor(...)`, `OrgEid=LegacyOrgID|StoryOrgID(id)`, `OrgAuthID=LegacyOrgClerkID|StoryOrgClerkID(id)`, `OrgRole=roleForIndex(idx,size,mix)`; idx = hero declaration-index + 1; prefix = stack (story 0) else `stack:story:<id>`. **Actual:** matches `users.go` (uid/clerkID/email/role derivation), `helpers.go::storyKeyPrefix/deterministicUUID/stackHost`, `userprofile.go::emailFor/splitName`, `persona.go::personaUserIndexFor`, `stories.go::Legacy*/StoryOrg*`. **Verdict: ALIGNED.** Note: `AuthID` uses `stackHost(prefix)` (= `slugify(prefix)`), so for story>0 the prefix `stack:story:<id>` is slugified — the producer must single-source from the seeder's own functions, not re-derive, to match exactly. Captured for build.

## Completeness Gaps

1. **Cockpit + deep-link catalog are undocumented in the corpus** — expected: they are M38's own deliverables. `stories-spec.md` already flags them ("The presenter cockpit + the Clerkenstein multi-identity seat-switch are M37–M38; this doc is the foundation those build on"). M38's Phase 5 adds the cockpit section to `stories-spec.md` + the up→present flow to `demo/README.md`. Not a blind area — the topic has a documented placeholder anchor and is in the milestone's `In:` scope.

## Applied Fixes
None required — no stale or load-bearing-wrong claims found. The consumer contract (M37) is fully documented and aligned; the producer (M38) is net-new and its anchor placeholders already exist.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. The M37 consumer contract the producer/cockpit build against is implemented and documented and aligned; the seeder id-derivation the roster producer must single-source is verified line-by-line; the cockpit + deep-link catalog are net-new M38 deliverables with existing doc anchors (DOC-ONLY, expected for an upcoming milestone).
