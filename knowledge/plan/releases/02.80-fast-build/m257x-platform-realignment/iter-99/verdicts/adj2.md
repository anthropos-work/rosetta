# Adjudicator 2 — verdicts for seats E / F / G (reading #21)

Re-derived from the platform clones at the ground-truth refs. No seat's evidence was taken as proof;
every anchor below was opened in the repo named. Corpus read at the working tree (`964b7a3`, whose
corpus content is `e858fd4`'s — the later commit adds only iter-99 plan files).

Instrument notes taken in-pass:
- All 2286 tracked `*.go` blobs at `app` `2035f9a4` are text to `git grep` (`git grep -Il ''` = 2286 of
  2286), so no NUL-bearing file can hide a Go match in the F-B1 predicate; `app/studio` (the only nested
  untracked repo under `app`) is Python and cannot contribute a `*.go` file either.
- `git grep` at an explicit ref sees tracked-but-gitignored blobs, so the `.gitignore` mechanism is
  neutralised for every ref-scoped measurement below.

---

## Seat E

### E B1 | `corpus/architecture/security_compliance.md:104-105` | UPHELD | IN-SCOPE
**Predicate:** the unpoliced remainder is *not* all "global reference data"; ≥4 members are tenant-scoped.

evidence: `stack-demo/app` @ `b948604f`, `internal/data/ent/schema/` —
`lab.go:63` `field.String("tenant_eid")`, `academy_chapter.go:84`, `academy_skill_path.go:78` (all
`tenant_eid`), `skill_path_session.go:43` `field.UUID("tenant_id")`. I grepped each of those four files
for `organization_id`, `OrganizationMixin`, `OrganizationIDMixin` and `func … Policy()`: **zero hits in
all four** (`skill_path_session.go:41` mentions `OrganizationMixin` only inside a comment explaining its
absence). So all four sit in the *remainder* the sentence characterises. Independently confirmed the
policy universe: `git grep -n 'func (.*) Policy()' b948604f -- 'internal/data/ent/schema/*.go'` returns
exactly five funcs in four files (`mixin.go:99`, `mixin.go:126`, `org_membership.go:172`,
`organization.go:56`, `user.go:116`) — none of the four can be policed. The same tenancy columns are
present at the paragraph's other named ref `5ba17044` (`lab.go`, `academy_skill_path.go`: 4 hits each),
so no pin rescues it. Corpus self-contradiction from a second in-scope file: `corpus/services/ai-labs.md:50`
calls the same table *"`tenant_eid` nullable = org-scoped else public"* and `:46-47` records its isolation
as *"fail-closed tenant-filtered reads"* in `ContentManager` — i.e. the application-code category the
sentence closes off. Set cardinality first: the remainder is 135 − 31 policed − 23 org_id-unpoliced = **81**
schemas, of which at least these 4 are per-tenant and one (`skill_path_session`) is per-user session state,
not reference data of any kind.

### E B2 | `corpus/architecture/service_taxonomy.md:405` | UPHELD | IN-SCOPE
**Predicate:** compose sets two service addresses and `backend → gotenberg` is a second live core edge.

evidence: `stack-demo/platform` @ the ref the claim itself names (`0c91421`),
`docker-compose.yml:48` `- AUTHORIZATION_ADDRESS=http://sentinel:8087` **and** `:57`
`- GOTENBERG_URL=http://gotenberg:3200`, both in `backend`'s `environment:` block. So *"Compose sets a
single service address"* is false at the named ref. The peer is a core-tier container in the default
selection by this same file's own roster (`service_taxonomy.md:81` lists **Gotenberg** as a Tier-1 row with
profile `core, backend, all`; `:65-67` counts it in the 5-container `core` start). The edge is live, not
vestigial: `app` @ `b948604f` `main.go:244` `GotenbergURL: os.Getenv("GOTENBERG_URL")`,
`internal/web/backend/coursebuilder/handler.go:244`, and the HTTP multipart client
`internal/converter/gotenberg.go:13,16`. Under the bullet's own heading (*Core Services ↔ Core Services*)
there are therefore two cross-process edges, not "exactly one" — a self-contradiction against `:81`/`:65-67`.
(The third clause, *zero `*_RPC_ADDR`*, is true and is not what is upheld.)

### E B3 | `corpus/architecture/security_compliance.md:185` | REJECTED | —
**Predicate (booked):** `README.md:21` resolves to the wrong construct in every candidate file.

evidence: `corpus/architecture/README.md:21` — printed in full — ends
*"…The doc covers what each provides and where its responsibilities begin and end (e.g. **cost tracking
lives in `app`, not the `ai` library**)."* That is precisely the construct the citing sentence corroborates.
The seat's candidate table truncated line 21 at the `authn` clause with an ellipsis and never read the tail.
class: mis-read — the anchor resolves to the supporting construct; only a *second*, likewise-supporting
occurrence exists at `:23`.

### E B4 | `corpus/architecture/service_taxonomy.md:37` (+ `:266`) | UPHELD | IN-SCOPE
**Predicate:** there is no "academy subgraph"; the corpus elsewhere says so explicitly.

evidence: `stack-demo/graphql-wundergraph` @ `60c229f3` — `supergraph-config-prod.yaml` declares exactly
one subgraph (`- name: backend`), `schemas/` holds one file (`backend.graphqls`), and every `academy`
match in the repo is a *type* inside `backend.graphqls` (`:1` is literally
*"# Academy domain GraphQL surface (Federation v2)"*). No `academy` subgraph, routing_url or SDL entry
exists. Same-file contradiction: `service_taxonomy.md:387` *"**`backend` alone (1)**"* and `:392`
*"Backend (`app`) — the only subgraph left"*. Cross-file, the corpus denies the construct by name:
`corpus/services/academy-backend.md:83` — *"**There is no separate "academy subgraph"** — these types live
in the `app`/`backend` federation subgraph — the only subgraph left."* Booked claim survives re-derivation
on both grounds (false against ground truth; contradicts another corpus claim).

---

## Seat F

### F B1 | `corpus/architecture/dependency_map.md:59` | UPHELD | IN-SCOPE
**Predicate:** `SKILLER_STREAM` at `2035f9a4` spans 3 Go files, not 4.

evidence: measured at the ref the sentence itself names. `git -C stack-demo/app grep -n SKILLER_STREAM
2035f9a4 -- '*.go'` → 6 lines: `main.go:1532`, `main.go:1537`, `subscriber_merge_test.go:907`, `:1015`,
`:1038`, `subscriber_wiring.go:165`. `git grep -l` at the same ref → **3** files. The occurrence count (6)
is right; the file cardinality is 3. No alternative scope yields 4 either: repo-wide the file set is **6**
(adds `knowledge/plan/…/decisions.md`, `knowledge/skiller-domain.md`, `terraform/main.tf`). Absence-hiding
mechanisms ruled out in-pass (all 2286 `*.go` blobs at that ref are text; `git grep -al` also returns 3;
`app/studio` is Python). The remainder of the cell verifies — `main.go:1276` @ `2035f9a4` is
`apiKeyManager,` as stated.

### F B2 | `corpus/architecture/dependency_map.md:58` (vs `:21`) | UPHELD | IN-SCOPE
**Predicate:** same file asserts both "not opt-in, stock `core`" and "gated OFF behind `MESSENGER_ENABLED`".

evidence: two present-tense passages in one file. `dependency_map.md:58` — *"That one is not opt-in: it is
the stock `core` selection."* `dependency_map.md:21`, third cell, un-pinned on this clause — *"now `app`'s,
and gated **OFF** on a developer machine behind `MESSENGER_ENABLED`"*. They are incompatible (rule 5), and
the ground truth says `:21` is the right side: at `app` `origin/main` `2035f9a4`, `main.go:1445` is
`if messengerEnabled {` wrapping the whole subscriber-server construction (`:1444` declares
`messengerSubServer`), with the comment at `:1437` *"ALL OF IT is behind MESSENGER_ENABLED … merely
constructing this server and calling Subscribe() attaches app to messenger's LIVE consumer group"*, and
`env_guards.go:61` `envMessengerEnabled = "MESSENGER_ENABLED"`; `stack-demo/platform` @ `0c91421`
`docker-compose.yml:84-92` deliberately does not set it (*"which default to OFF on a developer machine"*).
I confirmed the pin half honestly: at `9d00a313`, the ref the cell names, `git grep MESSENGER_ENABLED`
returns **exit 1 / zero hits tree-wide** (positive control `BREVO` returns 6 files at the same ref) and
`main.go:1442`/`internal/messenger/flow/streams.go:65` both resolve — so the sentence *was* true there.
Upheld on the self-contradiction, which is independent of the pin, and in the direction compose's own
comment calls hazardous.

### F B3 | `corpus/architecture/shared_libraries.md:11` (and `:3`) | UPHELD | IN-SCOPE
**Predicate:** services pull in four library repos, not five; the same file says so at `:167`.

evidence: cardinality derived first, per repo `go.mod` at each clone's own ref (`app b948604f`,
`sentinel 88bc5592`, `storage 4ce8ece5`, `messenger fa47850d`, `cms ca50c817`, `jobsimulation 462343b0`,
`roadrunner 87d8d443`): colony **7/7**, proto **7/7**, taxonomy **6/7**, ai **3/7**, authn **0/7** ⇒ the set
of repos any service pulls has cardinality **4**. `shared_libraries.md:11` says *"that shared plumbing lives
in **five small repos that the services pull in like any third-party dependency**"*. Same file, `:167`:
*"**No checked-out service imports the standalone `github.com/anthropos-work/authn`**"*; `:148` calls the
standalone module *legacy*. Also contradicted by `architecture_overview.md:80` and root `CLAUDE.md`.

---

## Seat G

### G B1 | `corpus/services/next-web-app.md:32` | UPHELD | IN-SCOPE
**Predicate:** the recruiter scoreboard is not reachable in `apps/web` for a genuine hiring org.

evidence: `stack-demo/next-web-app` @ `bb3313bc`,
`apps/web/src/context/UserStatusContext.tsx:141-173` — `userHasAllHiringOrgs` is computed from
`membership.organization.publicMetadata.isHiring` (`:144-149`), and when true the effect sets
`window.location.href = buildSwitchHandoffUrl({ targetProduct: 'hiring', … })` (`:168-172`), i.e. the user
is ejected out of `apps/web`. The corpus records the refutation in the very doc `:32` points at:
`corpus/services/hiring.md:53-58` (*"M222's dockerized-`apps/web` … premise was **falsified at M224** … are
**mutually exclusive**"*), `:359-364` and `:423-424`. Straight cross-file contradiction with ground truth on
the refuting side.

### G B2 | `corpus/services/hiring.md:80-82` | REJECTED | —
**Predicate (booked):** `manager.go:448` / `:485` resolve to the wrong construct.

evidence: true at the checkout `b948604f` (`:448` blank; `:485` the closing `}` of
`forceUserToOrganizationInClerk`; the hard error at `:536-537` guarded at `:535`) — but this document
declares its app-side grounding ref in its own re-grounding banner at `hiring.md:17`
(*"RE-GROUNDED — v2.8 M257x iter-23, against platform origin `2adcf71` / `app` @ `5ba17044`"*), and at
`5ba17044` **both anchors are exact**: `:448` is `switch org.IsHiring {` (the candidate branch the clause
describes, `:451` `antRole = enum.RoleCandidate`) and `:485` is `if !org.IsHiring {` with the hard error at
`:486-487`. `5ba17044` is an ancestor of the checkout, 42 commits back. The seat's tiebreaker — *"every
other `app` anchor in this document resolves at the checkout"* — does not hold as evidence: I diffed the
sampled anchors across the two refs (`intelligence.go:885-886/1700/1820`, `resolver_queries.go:1034-1035`,
`resolver_cms_queries.go:95,99-103`, the generated `graph.go:129546-129549`) and every one is
**byte-identical at both refs**, so they are silent about which ref the doc was anchored to; `manager.go` is
simply the one file whose lines moved.
class: ref-discipline — a dated claim, correct at the ref its document names, booked because newer
evidence contradicts it.

### G B3 | `corpus/services/hiring.md:38` | UPHELD | IN-SCOPE
**Predicate:** the cited twin `service_taxonomy.md:52` is about a different subject entirely.

evidence: `corpus/architecture/service_taxonomy.md:52` reads *"[`dependency_map.md`](./dependency_map.md)'s
content-generation flow, which had it right all along."* — the closing line of a blockquote (`:44-52`) about
Desk → Backend → Room generation and `studioManager.go:119` exec'ing `studio/gen.py`. It says nothing about
schemas, `jobsimulation`, or local stacks. The intended twin is `service_taxonomy.md:62`
(*"the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks"*). The paired citation
`dependency_map.md:78` **is** correct (it ends *"…or directly to the **`public`** schema (the legacy
`jobsimulation` schema is non-authoritative)."*), which isolates the defect to the single site.

### G B4 | `corpus/services/graphql-wundergraph.md:134` | UPHELD | IN-SCOPE
**Predicate:** the self-cited `:84` describes ports, not rebuild-on-SDL-change.

evidence: `graphql-wundergraph.md:84` is the **Ports** bullet — *"**Ports**: **8080 → 8080** (router
`listen_addr 0.0.0.0:8080`, `graphql_path /graphql`). **There is no `5050` at platform HEAD**…"*, running
`:84-88`, entirely about ports and the absent `5050`. The construct the sentence needs is at `:114-117`
(*"Adding/changing a subgraph **or a single field** requires re-running `wgc compose` and **rebuilding +
restarting**…"* / *"It *used to* rebuild whenever any subgraph schema changed, because the build context is
the parent dir (`..`)…"*). Wrong-construct self-citation; the file already carries a recorded instance of the
same class in the very bullet cited (`:86-88`).

---

## Dedup

No two bookings across E/F/G share a predicate. G-B3 and G-B4 are both *wrong-construct self-citation*
class but are distinct predicates at distinct anchors (`service_taxonomy.md:52` vs
`graphql-wundergraph.md:84`) and are not collapsed. E-B4 and G-B1 both concern a construct the corpus
corrects elsewhere, but the underlying predicates (no academy subgraph / scoreboard not in `apps/web`) are
unrelated.

BOOKED=11 UPHELD=9 REJECTED=2 IN-SCOPE-UPHELD-BLOCKERS=9
