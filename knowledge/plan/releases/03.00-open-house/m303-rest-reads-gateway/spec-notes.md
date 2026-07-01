# M303 Spec Notes

Technical detail during build.

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
    response:
      $ref: '#/components/schemas/MemberList'
  upstream:
    kind: connect-rpc
    service: app.OrganizationService
    method: ListMembers
    scope_from_principal: organization_id
  audit:
    resource_id: people.member.list
    action: read
```

The `scope_from_principal: organization_id` is the load-bearing isolation contract — the adapter injects
`Principal.OrganizationID` into every upstream call; no handler code reads `organization_id` off the request.

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

Runs per resource in the R1 catalog. The gate: **0 cross-tenant leakage over 5 consecutive integration runs**.

## Audit-row expectation matrix (draft — filled per iter)

| UC | resource_id | action | success rows | failure rows | notes |
|----|-------------|--------|--------------|--------------|-------|
| UC1 | people.member.list | read | 200 | 401 / 403 / 429 | |
| UC2 | people.member.get | read | 200 | 401 / 403 / 404 / 429 | 404 = not-in-org (isolation) |
| UC3 | learning.skill-path.list | read | 200 | 401 / 403 / 429 | |
| UC4 | learning.skill-path.get | read | 200 | 401 / 403 / 404 / 429 | 404 = not-assigned-in-org |
| UC5 | learning.session.list | read | 200 | 401 / 403 / 429 | |
| UC6 | verification.verified-skill.list | read | 200 | 401 / 403 / 429 | |
| UC7 | audit.event.list | read | 200 | 401 / 403 / 429 | feature-flagged |

## Rate-limit budget defaults (draft)

- **`default-min`:** 60 req/min per `principal_id`.
- **`default-day`:** 10,000 req/day per `principal_id`.

Per-resource overrides recorded here as they land (rare in R1).
