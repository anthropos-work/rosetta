# M37 — Spec notes

Authoritative design: [`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md) §4f
(Clerkenstein single→multi identity, G20). Clerkenstein lives in `.agentspace/rosetta-extensions/clerkenstein/`
(its own `knowledge/` is the source of truth; start at `knowledge/kb-index.md`).

## The constraint
_(single identity today: `clerk-frontend/resources.go:21` "the single universal identity every fake-FAPI
session resolves to"; `SingleSessionMode: true` at `:86`; the `Server` holds one `user DemoUser`,
`NewServer(DefaultDemoUser())` at `cmd/fake-fapi/main.go:18`.)_

## What's cheap vs the real work
_(cheap: per-identity JWT minting already exists — `sessionClaimsFor(u,…)` + `MintRS256`, universal HS256 key,
claim shape `{AuthID,Eid,Org*,OrgRole}`. Real work: (a) the registry, (b) active-user selection [O11].)_

## O11 — seat-switch mechanism
_(spike + record: token-injection [set the cookie/storage clerk-js reads] vs a parameterized FAPI handshake.)_

## Alignment
_(new measured surface — author a DNA + goldens; `/align-run` must keep all 4 existing surfaces at 100%.)_
