# M38 — Spec notes

Authoritative design: [`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md) §4d (the
presenter cockpit), §17. Lives in `rosetta-extensions/demo-stack` (a standalone served panel, offset port —
NOT an in-app overlay, preserving the zero-platform-repo-edit line, D15).

## The single-source property
_(the panel reads the SAME `stack.stories.yaml` that seeded the data, D9 — annotations describing a hero are
the same ones that scoped her seed. No drift between data and cockpit menu.)_

## The two actions
_([Login as] → M37 active-user selection. [Jump to section] → the hero's `jump_to` next-web deep-link.)_

## O9 — deep-link catalog
_(enumerate the next-web routes per vantage: profile, Spotlight, workforce tabs, talk-to-data,
talent-pool/mobility; note which need a hero/skill id param. One UI pass.)_

## Pre-flight audits — The panel (first section)
KB-fidelity audit M38: **GREEN** (0 stale, 0 blind areas). Report:
[`kb-fidelity-audit.md`](kb-fidelity-audit.md). The M37 consumer contract (registry + roster JSON +
handshake `?__clerk_identity` + `/v1/demo/{identities,select}`) is implemented + documented + aligned in
`clerkenstein/knowledge/architecture.md`; the seeder id-derivation the roster producer single-sources is
verified line-by-line (`stack-seeding/seeders/{users,helpers,userprofile,persona}.go`,
`blueprint/stories.go`). Cockpit + deep-link catalog are net-new M38 deliverables (DOC-ONLY, expected).

## Topic → doc → code triples (for fast re-audit)
- Clerkenstein consumer → `clerkenstein/knowledge/architecture.md` § Multi-identity → `clerk-frontend/{registry,server}.go` + `cmd/fake-fapi/main.go`
- Seeder id-contract → M37 `decisions.md` § ARCH → `stack-seeding/seeders/{users,identity,helpers,userprofile,persona}.go` + `blueprint/stories.go`
- Stories blueprint → `corpus/ops/demo/stories-spec.md` § Stories & Heroes (M35) → `stack-seeding/blueprint/blueprint.go` + `presets/stories.seed.yaml`
- Demo bring-up → `corpus/ops/demo/README.md` → `demo-stack/up-injected.sh` + `stack-injection/gen_injected_override.py`

## Roster id-contract (single-source from the seeder; NEVER re-derive in Clerkenstein)
Per hero, indexed by population user-index `idx` (= hero declaration order + 1), prefix = stack (story 0) else `stack:story:<id>`:
- `key`        = hero `id` (stories.yaml, e.g. `maya-thriving`) — the seat-switch handle (NOT a token claim)
- `auth_id`    = `fmt.Sprintf("user_seed_%s_%d", stackHost(prefix), idx)`  (stackHost = slugify)
- `eid`        = `deterministicUUID(fmt.Sprintf("%s:user:%d", prefix, idx))`
- `email`      = `emailFor(first,last,storyEmailDomainFor(st),idx)`  (first/last = splitName(hero.Name))
- `firstname`/`lastname` = `splitName(hero.Name)`
- `org_auth_id`= story 0 → `LegacyOrgClerkID` else `StoryOrgClerkID(st.ID)`
- `org_eid`    = story 0 → `LegacyOrgID` else `StoryOrgID(st.ID)`
- `org_role`   = `roleForIndex(idx, st.Size, st.RoleMix)`
First roster entry = the default active seat. The producer lives in the `seeders` package (where these
unexported helpers live) and is exercised by `stackseed --roster-export`.

## Login-as = ONE handshake redirect (the elegant M37 seam)
The cockpit's [Login as] + [Jump to section] collapse into a single redirect to the FAPI:
`https://<fapi-host>/v1/client/handshake?__clerk_identity=<key>&redirect_url=<jump_to>` — `handleHandshake`
selects the seat from `__clerk_identity` THEN establishes the session + redirects to `redirect_url`, so the
chosen hero is active everywhere AND the browser lands on her screen in one move.
