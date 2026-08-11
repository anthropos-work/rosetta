# Adjudicator 4 — seat E (readings #31 and #32)

## Trees read, and at which refs

Re-derived at this adjudication's open with **no fetch**; every ref below was read from the working clone
and matches the brief's ground-truth table exactly:

| repo | path | HEAD | origin/main |
|---|---|---|---|
| platform | `stack-demo/platform` | `0c91421d` | `0c91421d` |
| app | `stack-demo/app` | `ad9f3c49` | `ad9f3c49` |
| app/studio (nested) | `stack-demo/app/studio` | `aeec036a` | — |
| cms/studio (nested) | `stack-demo/cms/studio` | `aeec036a` | — |
| next-web-app | `stack-demo/next-web-app` | `8297c684` | `f97ba659` |
| sentinel | `stack-demo/sentinel` | `f2c46190` | in sync |
| studio-desk | `stack-demo/studio-desk` | `41ee3575` | in sync |
| ant-academy | `stack-demo/ant-academy` | `22df69dd` | in sync |
| cms | `stack-demo/cms` | `ca50c817` | `f38c0c4a` |
| jobsimulation | `stack-demo/jobsimulation` | `462343b0` | `82cb66ec` |
| messenger | `stack-demo/messenger` | `fa47850d` | `e9421c68` |
| storage | `stack-demo/storage` | `4ce8ece5` | `9f8cb532` |
| roadrunner | `stack-demo/roadrunner` | `87d8d443` | in sync |
| graphql-wundergraph | `stack-demo/graphql-wundergraph` | `60c229f3` | in sync |
| rosetta-extensions (pinned per-stack) | `stack-demo/rosetta-extensions` | `09d06070` | `4cb920aa` |
| rosetta-extensions (authoring) | `.agentspace/rosetta-extensions` | `43049308` | — |

Trees actually opened for these three verdicts: **`stack-demo/next-web-app` @ `8297c684`** (B1),
**`stack-demo/app` @ `ad9f3c49`, `b948604f` and `9ecade240`** (B2), **`stack-demo/platform` @ `0c91421d`,
`0dab54d`, `2adcf71`** (B3). No `rosetta-extensions` claim was in this seat's booked set, so no tree-choice
question arose; the `stack-demo` (consumption) clone is what I would have used.

`git status --porcelain` was **empty at my open**. At my close the only entry is this verdict file itself
(`?? knowledge/plan/.../iter-119/verdicts/adj-4.md`) — the one file I was permitted to create. No fetch, no
git state change, no other edit.

Seat E reading **#31 booked ZERO blockers**; I read it in full as evidence. Its enumerated clearances of the
`app` mux (six handlers), the `NEXT_PUBLIC_BACKEND_API_URL` set, and the profile-token ladder are consistent
with what I re-derived independently below, and its treatment of `backend.md:314` as "explicitly historical"
is corroborating (not decisive — I resolved the pin myself). All three bookings adjudicated below are from
reading **#32**.

---

## Bookings

```
E B1 | corpus/architecture/frontend_architecture.md:32 | UPHELD | IN-SCOPE | PREDICATE: packages/graphql's React Query hooks are codegen-generated; they are hand-authored.
   evidence: next-web-app@8297c684:packages/graphql/codegen.ts declares `preset: 'client'`, `plugins: []`,
   `generates: {'./src/__generated__/': …}` — the client preset, and NO `typescript-react-query` plugin is
   configured or even a dependency (packages/graphql/package.json lists only cli / client-preset /
   typescript / typescript-resolvers). `git ls-tree -r 8297c684 -- packages/graphql/src/__generated__`
   returns exactly four files — fragment-masking.ts, gql.ts, graphql.ts, index.ts — and
   `git grep -E 'useQuery|useMutation|useSuspenseQuery' 8297c684 -- packages/graphql/src/__generated__`
   exits rc=1 (zero hits). I also checked the working tree on disk (same four files, no hooks) and
   .gitignore (no `generated`/`graphql` rule), so no third instrument hides a generated hook. The real
   hooks are 256 TRACKED, hand-written files under packages/graphql/src/hooks/<domain>/ —
   e.g. hooks/academy/useAcademyProgress.tsx opens with a prose comment and imports
   `useQuery` from '@tanstack/react-query' plus `ACADEMY_PROGRESS` from '../../query/academy'. The same
   document states the truth 23 lines later at :55 ("React Query hooks are hand-authored on top of these
   typed documents"), so this is also a live self-contradiction (brief rule 5) — and :54 ("`packages/graphql`
   contains the generated types and hooks") is a SECOND anchor for the same predicate that the seat did not
   book. Both false anchors sit in corpus/architecture/**.
```

```
E B2 | corpus/services/backend.md:314 | REJECTED | (would be IN-SCOPE) | PREDICATE: n/a — claim is true at the ref it names.
   evidence: The bullet is dated by its own head parenthetical — "**AI Labs LabSession** (Phase B PR 2,
   #896)" — inside the section heading "## Recent Feature Additions (Q1-Q2 2026)" (backend.md:303), a
   changelog whose sibling bullets are likewise pinned ("v1.266.0+, May 2026", "v1.266.2", "v1.267.1",
   "`feat/company-context-m1m2` branch"). That pin resolves to exactly one commit:
   `git log -S 'NewLabSessionServiceHandler' -- main.go` in stack-demo/app returns the single commit
   `9ecade240` "feat(labsession): add LabSession Ent entity + Connect-RPC service (Phase B PR 2) (#896)"
   (2026-05-29). I opened main.go AT `9ecade240`: the mux carries exactly three registrations —
   `:240` UsersService, `:243` OrganizationsService, `:249` LabSessionService. LabSession IS the third
   handler, immediately after Users and Organizations. The claim is TRUE at its named ref.
   I confirmed the seat's current-state measurement independently and it is right about TODAY: at
   `ad9f3c49` the mux is `main.go:1297` Users, `:1298` Organizations, `:1306` Skiller, `:1314` JobSimulation,
   `:1323` CMS (inside `if cmsRPCServer != nil`), `:1338` LabSession — sixth; identically shaped at
   `b948604f` (`:1187/:1188/:1196/:1204/:1213/:1228`). But that is newer evidence against a dated claim.
   Nor is it a live self-contradiction with backend.md:30 / :106 (both of which I read and both of which
   correctly say six, LabSession last, at named refs): those describe the current mux, this describes what
   PR #896 shipped — two times, not two incompatible assertions (brief rules 5 and 7). The bullet's own
   "**The real HTTP client has since LANDED** … @ `app` `b948604` v1.366.0" is an explicit update marker
   that fixes the rest of the bullet's baseline at #896, and the sibling AI-Readiness bullet says "at HEAD"
   where it means now — the section distinguishes dated from current on purpose.
   class: ref-discipline — a pinned changelog entry booked because newer evidence contradicts it; it is the
   pin working, exactly the class brief rule 2 says to expect, reject, and name.
```

```
E B3 | corpus/services/graphql-wundergraph.md:177 (clause of the :172-178 bullet) | UPHELD | IN-SCOPE | PREDICATE: NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT's compose env line was :236 at 0dab54d and :352 at 2adcf71.
   evidence: The bullet separates two constructs at HEAD — the runtime-env pair (`docker-compose.yml:135`
   VITE_GRAPHQL_ENDPOINT, `:160` NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT) and the build-arg pair (`:119`, `:151`)
   — then says "Those line numbers … were `:220`/`:236` at `0dab54d` and `:334`/`:352` at `2adcf71`".
   Measured by `git show <ref>:docker-compose.yml` in stack-demo/platform:
     · `0c91421d`: VITE arg `:119` / env `:135`; NEXT_PUBLIC arg `:151` / env `:160`  → the HEAD pairs are correct.
     · `0dab54d`: next-web-app service opens `:228`, `args:` `:235` → `:236` is
       `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://${PUBLIC_HOST:-localhost}:8082/graphql/query` under **args**;
       `environment:` `:243` → the env line is `:245`. VITE: arg `:204`, env `:220`.
     · `2adcf71`: next-web-app opens `:344`, `args:` `:351` → `:352` is the **build arg**; `environment:` `:359`
       → the env line is `:361`. VITE: arg `:318`, env `:334`.
   So each historical "pair" is one env line plus one build-arg line: it is neither the env pair
   (`:220`/`:245`, `:334`/`:361`) nor the build-arg pair (`:204`/`:236`, `:318`/`:352`). Under EITHER reading
   of the antecedent, exactly one of the two offsets at each ref names the wrong construct — and it is always
   the next-web half. This is not ref-discipline (the claim names its own refs and is false AT them) and not
   a historical-record anchor under brief rule 7 (it asserts what a file contained at a named ref, and the
   file at that ref says otherwise). A line-existence check cannot catch it because both candidate lines
   carry the same env-var name — which is what makes it load-bearing in a passage whose entire moral is
   "grade the construct, not the offset".
```

---

## PREDICATE ROLL-UP

```
P1 | packages/graphql's React Query hooks are codegen-generated; they are hand-authored. | anchors: E B1 @ corpus/architecture/frontend_architecture.md:32  (unbooked sibling anchor for the same predicate, noted not counted: frontend_architecture.md:54)
P2 | NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT's compose env line was :236 at 0dab54d and :352 at 2adcf71. | anchors: E B3 @ corpus/services/graphql-wundergraph.md:177
```

No two upheld anchors collapse onto one predicate: P1 is a next-web-app packaging claim in
`corpus/architecture/**`, P2 a platform-compose offset claim in `corpus/services/**`. They are distinct.

BOOKED=3 UPHELD=2 REJECTED=1 IN-SCOPE-UPHELD-BLOCKERS=2 DISTINCT-IN-SCOPE-PREDICATES=2
