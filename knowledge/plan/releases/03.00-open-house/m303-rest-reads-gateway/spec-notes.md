# M303 Spec Notes

Technical detail during build. Post-v0.2 scope-correction — the catalog is at **Talk-to-Data data parity**
(9 products / 35 resources / ~44 endpoints / ~55 backing tables — spec §4.2). The `askengine.TableRegistry` +
`rules.md` in the platform backend are the authoritative parity source; every family the M303 handler builds
against those.

## Catalog entry template (draft)

```yaml
# app/internal/customerapi/catalog/catalog.yaml (extract)

- name: people.member.list
  product: people
  resource: member
  action: list
  action_type: read
  audience: [free, paying, enterprise, partner]
  scopes: [people:read]
  rate_limit_bucket: default-min
  rest:
    method: GET
    path: /v1/people/members
    query:
      - name: limit
        type: integer
        default: 50
        max: 200
      - name: page_token
        type: string
      - name: language
        type: string
        enum: [english, italian, spanish, french, german, dutch]
        default: english
    response:
      $ref: '#/components/schemas/MemberList'
  upstream:
    kind: connect-rpc
    service: app.OrganizationService
    method: ListMembers
    scope_from_principal: organization_id
  rules_applied: [CR1, CR2, CR3, CR9, CR10]
  audit:
    resource_id: people.member.list
    action: read
```

The `scope_from_principal: organization_id` is the load-bearing isolation contract — the adapter injects
`Principal.OrganizationID` into every upstream call; no handler code reads `organization_id` off the request.
`rules_applied:` is a v0.2 addition — enumerates the CR1–CR15 rules the family must honor; the CR-suite driver
uses it to derive which contract tests to run per family.

## CR-rule enforcement patterns (v0.2 — per-rule handler recipe)

For each rule from spec §4.5 that applies to a family, the handler wires the enforcement pattern below. These
are **shared** — one implementation per rule, referenced from every family that needs it.

| Rule | Enforcement layer | Shared helper |
|---|---|---|
| CR1 Principal-scoping | Handler middleware; upstream `scope_from_principal` | `customerapi.PrincipalMiddleware`, `Scope(principal, org)` |
| CR2 Soft-delete exclusion | Repository predicate | `repo.NotDeletedPredicate()` |
| CR3 Active-member definition | Shared predicate constant | `repo.ActiveMember = "status='active' AND deleted_at IS NULL"` |
| CR4 Completed-sim definition | Shared predicate | `repo.CompletedSession()` |
| CR5 Mapped ≠ Verified | Response shape (`mapped[]`, `verified[]`) | `skills.SplitByVerification(evidences)` |
| CR6 Org-scale + `max_level` | Level-normaliser resolves `max_level` from `organization_settings.options` | `levels.Normalise(rawScore, maxLevel)` (level-normaliser); every level-carrying DTO embeds `max_level` |
| CR7 Skill-level source column | Repository query | `repo.MemberSkills` reads `user_skill_evidences.level`; lint gate on any other column |
| CR8 Forbidden stale tables | CI grep-gate on `internal/customerapi/` | `.githooks/customerapi-lint.sh` |
| CR9 Person identifier = user UUID | Route contract binds `{member_id}` to `m."user"` | `paths.MemberID` type wrapper |
| CR10 Catalog resolution | Response transformer joins taxonomy/catalog labels | `catalog.Resolver` (skill/sim/skill-path/job-role) |
| CR11 Localization | Shared locale resolver + `?language=` param | `locale.Resolve(param, translations, defaultEN)` |
| CR12 AI Readiness live ≠ frozen | Distinct handlers per resource | `aireadiness.LiveHandler` vs `aireadiness.CycleSnapshotHandler` |
| CR13 Profile-history self-scoping | Middleware guarantee (not query param) | `profilehistory.ScopeForPrincipal(principal, pathMemberID)` |
| CR14 Academy visibility | Repository predicate | `academy.PublishedPredicate + academy.TenantScope` |
| CR15 Read-only R1 | HTTP-method allow-list on the router | `router.OnlyGET` |

## Isolation-test skeleton (draft — Go integration)

```go
// app/internal/customerapi/isolation_test.go

func TestIsolation_PeopleMemberList(t *testing.T) {
    // arrange: two orgs, each with a minted API key + a distinct roster
    orgA := seedOrgWithMembers(t, 3)
    orgB := seedOrgWithMembers(t, 5)
    keyA := mintKeyFor(t, orgA)
    keyB := mintKeyFor(t, orgB)

    // act: A calls with A's key, then B's key
    respA := getWithKey(t, "/v1/people/members", keyA)
    respB := getWithKey(t, "/v1/people/members", keyB)

    // assert: A sees only A's roster; B sees only B's roster; 0 overlap
    require.Equal(t, 3, len(respA.Members))
    require.Equal(t, 5, len(respB.Members))
    require.Empty(t, intersection(respA.MemberIDs(), respB.MemberIDs()))
}
```

Runs per family in the R1 catalog. The gate: **0 cross-tenant leakage over 5 consecutive integration runs across
all ~44 endpoints** (spec §5.7).

## CR-rule test skeleton (draft — one per rule per applicable family)

```go
// app/internal/customerapi/cr_test.go — example: CR6 org-scale
func TestCR6_OrgScale_MemberSkills(t *testing.T) {
    orgA := seedOrgWithLevelsCount(t, 7)          // levels_count=7 fixture
    member := seedMemberWithSkillLevel(t, orgA, /*raw*/ 85 /*of 100*/)
    keyA := mintKeyFor(t, orgA)

    resp := getWithKey(t, "/v1/people/members/"+member.ID+"/skills", keyA)

    require.Equal(t, 7, resp.MaxLevel)              // response carries max_level
    for _, s := range resp.Verified {
        require.LessOrEqual(t, s.Level, 7)          // on org scale, not 0..100
        require.GreaterOrEqual(t, s.Level, 0)
    }
    require.NotContains(t, mustJSON(resp), "\"raw_score\"")   // raw 0-100 NEVER leaks
}
```

## Audit-row expectation matrix (draft — filled per iter, populated across ~44 endpoints at close)

Per-endpoint rows land as families close. Illustrative rows for the FIRST-USABLE seven:

| Endpoint | resource_id | action | success | failure | notes |
|----|-------------|--------|--------------|--------------|-------|
| `GET /v1/people/organization` | people.organization.get | read | 200 | 401 / 403 / 429 | |
| `GET /v1/people/members` | people.member.list | read | 200 | 401 / 403 / 429 | |
| `GET /v1/people/members/{member_id}` | people.member.get | read | 200 | 401 / 403 / 404 / 429 | 404 = not-in-org (isolation) |
| `GET /v1/people/members/{member_id}/skills` | people.member.skill.list | read | 200 | 401 / 403 / 404 / 429 | `mapped`/`verified` split, `max_level` embedded (CR5+CR6) |
| `GET /v1/simulations/sessions` | simulations.simulation-session.list | read | 200 | 401 / 403 / 429 | |
| `GET /v1/simulations/sessions/{session_id}` | simulations.simulation-session.get | read | 200 | 401 / 403 / 404 / 429 | CR4 completed-def; CR6 score-on-org-scale |
| `GET /v1/learning/skill-path-sessions` | learning.skill-path-session.list | read | 200 | 401 / 403 / 429 | |
| `GET /v1/audit/events` | audit.audit-event.list | read | 200 | 401 / 403 / 429 | feature-flagged |

## Rate-limit budget defaults (draft)

- **`default-min`:** 60 req/min per `principal_id`.
- **`default-day`:** 10,000 req/day per `principal_id`.
- **`heavy-list`:** 20 req/min for the wide-nested reads (sim-session drill-down families) — decide at build.

Per-resource overrides recorded here as they land.

## Static-lint CI gates (CR7 + CR8)

A grep-based CI check on `internal/customerapi/` — fails the build on any reference to:
- `local_jobsimulation_sessions`
- `local_skill_path_sessions`
- `membership_skills.skill_level` (except inside a documented migration-shim block, which does not exist in R1)

Wired into `.githooks/` + the CI pipeline. Rule-fidelity is guarded, not hoped-for.
