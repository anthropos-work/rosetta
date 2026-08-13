# Auditor F — 6 files / 1486 lines — 4 blockers, 8 minors

## B-F1 `service_taxonomy.md:288-289` — FALSE RETRACTION [file EDITED by iter-39]
Says the `directus/directus:10.10.1` compose service on :8055 with admin@example.com/password
"has NEVER existed in the platform compose". It DID: `git show a2a3ee6^:docker-compose.yml` →
:383 directus: / :384 image directus/directus:10.10.1 / :386 8055:8055 / :409 ADMIN_PASSWORD=password.
Deleted at `a2a3ee6` (2026-02-27). Contradicts the corpus's OWN fenced SoT that the same section links:
`platform-migration-status.md:86` ("Removed from compose at a2a3ee6").
(Only the admin@example.com EMAIL is unfound in history — 0 hits.)
CLASS: retraction-as-new-claim + over-correction ("does not exist now" -> "never existed").

## B-F2 `service_taxonomy.md:145` — Studio-Desk tech listed as including **React** [EDITED]
No React: 0 react/vue/angular in package.json, 0 .tsx/.jsx in repo. The doc under the SAME auditor
contradicts it: `studio-desk.md:20` "vanilla TS frontend, no framework", `:30` "no React/Vue/Angular".

## B-F3 `service_taxonomy.md:136` — Tier 2 "Standalone processes (not in main docker-compose)" [EDITED]
Studio-Desk IS in platform docker-compose.yml:311 with `profiles: [studio-desk, all]` at :342.
Contradicted by :75 in the same file, by `frontend_architecture.md:11`, and by `studio-desk.md:21`.

## B-F4 `platform-migration-status.md:60` — anchor attached to 4 domains, resolves for 1 [UNTOUCHED]
"wired at app/main.go:604" — :604 wires jobsimulation ONLY. skiller :573, skillpath :634, cms :1034.

## MINORS: 8
## UNVERIFIABLE: gh/archive; GitHub org census (93 repos); private Go modules; legacy DB schemas;
## studio-room not cloned; production infra repo not cloned; demo-measured latency + directus image internals.

## NOTABLE CONFIRMATION: the supergraph ladder 5->4->3->1 verified — `git show 915da06` deletes BOTH
## cms.graphqls AND jobsimulation.graphqls in one commit, so "3->1" is right and the commit's own
## "2->1" subject is the misleading thing. (Independently agrees with auditor C.)
