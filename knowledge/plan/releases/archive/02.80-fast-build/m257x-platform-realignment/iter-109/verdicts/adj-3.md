# Adjudicator 3 — seats E and G, readings #27 and #28

**Trees read, and which one settled what.** Every clone verified at this adjudication's open and matching
the brief's ground-truth table byte for byte: `platform 0c91421d` · `app ad9f3c49` · `next-web-app 8297c684`
· `sentinel f2c46190` · `studio-desk 41ee3575` · `ant-academy 22df69dd` · `cms ca50c817` ·
`jobsimulation 462343b0` · `messenger fa47850d` · `storage 4ce8ece5` · `roadrunner 87d8d443` ·
`graphql-wundergraph 60c229f3` · nested `app/studio` + `cms/studio` both `aeec036a` ·
`stack-demo/rosetta-extensions 09d06070` (pinned consumption clone) ·
`.agentspace/rosetta-extensions 680e8529` (authoring copy). **No fetch was run.**

None of the ten bookings in my partition is a claim about a fence's own verdict or configuration, so
**the authoring copy was not needed and was not read**; no rext claim appears in any booked blocker. Every
platform claim was re-derived at the ref the claim itself names, and where a claim named no ref it was
graded at the checkout.

Each verdict below was re-derived by opening the platform file myself. The seat's citation was treated as a
pointer only.

---

## Verdicts

### r27-E B1
`r27-E B1 | corpus/services/graphql-wundergraph.md:274 | UPHELD | IN-SCOPE | PREDICATE: graphql-wundergraph's own CLAUDE.md "Version Tracking" section is stale.`
   evidence: `graphql-wundergraph@60c229f3:CLAUDE.md:70-79` reads *"## Version Tracking / Service versions
   are tracked in `subgraphs.conf`. There is exactly **one** pin now: / ```BACKEND=v1.360.0``` / The `CMS`,
   `JOBSIMULATION`, `SKILLER` and `SKILLPATH` entries were removed as each of those services merged into
   `app`."* `subgraphs.conf` at the same ref is the single line `BACKEND=v1.360.0` — a byte match. So the
   section names `subgraphs.conf` as the source of truth, carries the correct single pin, and states the
   correct removal history: it *is* the current form of the very claim the corpus offers as its correction.
   `git log -- CLAUDE.md` shows the file was last rewritten by `60c229f3` itself, the checkout. The corpus
   sentence has a compound subject and predicates staleness of both halves; only the second half survives —
   `CLAUDE.md:85` (`wgc router compose -i supergraph-config-local.yaml`) is genuinely stale, since
   `ls supergraph-config-*.yaml` @ `60c229f3` returns `-compose`, `-dev`, `-prod` and no `-local`. Claim
   names no ref, so graded at the checkout. A reader is told to distrust an accurate section of the
   ground-truth repo.

### r27-E B2
`r27-E B2 | corpus/services/roadrunner.md:13-14 | UPHELD | IN-SCOPE | PREDICATE: roadrunner/terraform/main.tf:19 was last touched at 87d8d44 (2026-06-19).`
   evidence: `git -C stack-demo/roadrunner blame -L 19,19 87d8d443 -- terraform/main.tf` →
   **`84a4b4f1 (Mattia Sasso 2025-12-15)`**. `git show --stat 87d8d443` touches exactly one file,
   `.github/workflows/bump-version.yml | 3 +++` — it never goes near terraform, and it is the repo's HEAD,
   so "not touched since it" is vacuous by construction while the parenthetical `(2026-06-19, before the
   fold)` presents it as the date of the last touch. The file's own most recent touch is `e45eb61`
   (2026-05-27, a line-11 module-source swap). The subject of "has not been touched" in the corpus sentence
   is `main.tf:19`, the line — not the repo — so the sentence asserts that `87d8d44` is that line's last
   toucher, which it is not. The corpus's own fenced authority, which this doc points at two lines below,
   gives the correct provenance: `corpus/architecture/platform-migration-status.md:90` — *"last changed at
   **`84a4b4f` (2025-12-15)** … **That count is not a decision about the fold; it predates it by seven
   months** … `git blame -L 19,19` names `84a4b4f`; a file-level `git log` is not line provenance."*
   roadrunner.md never names `84a4b4f` anywhere. Materiality is modest — the sentence's own conclusion
   ("before the fold") survives — but the provenance sha is measurably wrong and the map warns against this
   exact error class.

### r27-E B3
`r27-E B3 | corpus/architecture/dependency_map.md:19 | UPHELD | IN-SCOPE | PREDICATE: Storage's compose environment block at platform 0dab54d declares eight variables.`
   evidence: I enumerated the SET before recomputing the sum, per rule 4. At `platform 0dab54d`,
   `docker-compose.yml:116` is `environment:` — the **key** — and `:117-123` are the members:
   `AWS_DEFAULT_REGION`, `AWS_REGION`, `ENVIRONMENT`, `PORT`, `RPC_PORT`, `SERVICE_NAME`,
   `STORAGE_S3_PUBLIC_BUCKET`. Cardinality of the set = **7**, over 8 lines. The `8` is the cardinality of
   the corpus's own cited line range, not of the set the sentence ranges over. The rest of the sentence
   verifies exactly: `depends_on` at `:126-131` is `redis` + `postgresql`, both `service_healthy`, and
   neither `DB_CONNECTION` nor `REDIS_ADDR` is among the seven — i.e. the predicate the count *serves*
   survives, and only the cardinality is false. Booked as a measurement ("Measured in the platform's own
   compose"), which is what makes the number load-bearing rather than colour.

### r27-E B4
`r27-E B4 | corpus/services/backend.md:37-41 | UPHELD | IN-SCOPE | PREDICATE: storage's StorageService Connect handler is declared in storage:sdk/storage/v1/service.go.`
   evidence: `storage@4ce8ece5:sdk/storage/v1/service.go` is 47 lines and is the **consumer-side SDK**: it
   declares `type Service interface {GetObject, PutObject, GetPresignedUrl}` (`:13-17`) and the constructors
   `NewService(c storagev1connect.StorageServiceClient, namespace string)` (`:19`) /
   `NewPublicService(...)` (`:34`). Every `StorageService` token in the file is
   `storagev1connect.StorageService**Client**` — the type an external caller imports to *dial*.
   `git grep -n 'Handler\|ServeMux' 4ce8ece5 -- sdk/storage/v1/service.go` returns **nothing**: no handler,
   no mux. Storage's real registration is `cmd/root.go:62-66` —
   `rpcMux := http.NewServeMux(); rpcMux.Handle(storagev1connect.NewStorageServiceHandler(rpcsrv.New(...)))`
   — and its server implementation is `internal/rpcsrv/rpcsrv.go`, which exists and is the exact sibling of
   the messenger citation standing beside it in the same sentence
   (`messenger@fa47850d:internal/rpcsrv/rpcsrv.go:26,:29`, verified) and the exact analogue of the
   roadrunner one (`roadrunner@87d8d443:cmd/root.go:87` = `rpcMux.Handle(roadrunnerv1connect.NewRoadRunner
   ServiceHandler(...))`, verified). Of three parallel citations for a sentence whose whole point is
   *"on their own muxes"*, two land on the construct and one lands on its mirror image. The underlying
   proposition is true and independently verified at `cmd/root.go:63`; the anchor is what fails, and a
   line-existence check can never catch it because the file name contains the service's name.

### r28-E B1
`r28-E B1 | corpus/services/backend.md:360 | REJECTED | IN-SCOPE | PREDICATE: (not established) — app's Atlas configuration declares a single migration directory.`
   evidence: The net-new construct is real. `app@ad9f3c49:atlas.hcl` declares **two** envs: `env "local"`
   (`:6-21`, `dir = "file://terraform/migrations"`, `src = "ent://internal/data/ent/schema"`,
   `revisions_schema = "public"`) and `env "sentinel"` (`:50-66`, `dir =
   "file://terraform/migrations-sentinel"`, `src = "file://terraform/sentinel/schema.sql"`,
   `revisions_schema = "sentinel"`), the second added under an emphatic *"sentinel-in-app v10.0 / M1001 —
   the SECOND Atlas pipeline"* comment block. `git ls-tree ad9f3c49 terraform/migrations-sentinel/` returns
   `20260804151548_adopt_casbin_rules.sql` + `atlas.sum`; the same path is **empty at `b948604f`**. But every
   clause the corpus sentence actually asserts verifies at the ref it names: versioned Atlas migrations *do*
   live in `terraform/migrations/`; the quoted `dir` and `src` are `env "local"`'s verbatim; and the
   emphasised half — *"There is no top-level `migrations/` dir"* — is about a **different** directory and is
   true. The recipe is also correct rather than misleading: `platform@0c91421d:Makefile:87,:95` runs
   `atlas migrate apply --env local`, which is exactly what the doc documents, and the seat's own
   "could-not-settle" list concedes it cannot show the second env is *supposed* to run there. `casbin_rules`
   is still created by the live `sentinel` container, so no reader following the section gets a broken stack
   today.
   class: other — an omission of a net-new construct, not a false proposition; no clause fails at the named
   ref and the documented command matches what `make migrate` actually does. Worth a corpus addition next
   cycle; it is not a false predicate this reading.

### r28-E B2
`r28-E B2 | corpus/architecture/dependency_map.md:19 | UPHELD | IN-SCOPE | PREDICATE: Storage's compose environment block at platform 0dab54d declares eight variables.`
   evidence: same anchor, same falsehood, re-derived once — see r27-E B3. Collapses onto **P3**; this is the
   same seat booking the same line in a second reading, so it adds an anchor-instance but **no** predicate.

### r28-E B3
`r28-E B3 | CLAUDE.md:248 (repository root) | UPHELD | OUT-OF-SCOPE | PREDICATE: The ai module is imported as a private Go module and pulled at Docker build.`
   evidence: I enumerated `anthropos-work/*` requires across all seven on-disk Go repos at their own HEADs:
   colony **7/7**, proto **7/7**, taxonomy **6/7** (roadrunner absent), **ai 2/7**, authn **0/7**.
   `app@ad9f3c49:go.mod:14-18` has `analytics-go`, `colony`, `proto`, `storage`, `taxonomy` and **no**
   `anthropos-work/ai`; `1e457fa70` (2026-08-04, *"refactor(ai): fold the ai library into app as
   internal/ai"*) removed it. `sentinel@f2c46190:go.mod` never had it. The only two requires left are
   `cms@ca50c817:go.mod:9` and `jobsimulation@462343b0:go.mod:11` (both `v1.40.2`) — frozen repos with no
   compose service and no `repos.yml` entry, so `make init` never clones them and no Docker build resolves
   them. Two corpus documents already carry the correction —
   `corpus/architecture/dependency_map.md:48` (*"**No repo a stack builds**"*) and
   `corpus/architecture/shared_libraries.md:126` (same verdict, with the `module_import_guard_test.go`
   one-way door) — so the root instruction file is the surviving instance of a falsehood the corpus has
   otherwise repaired. **Out of scope**: the anchor is the repository-root `CLAUDE.md`, not
   `corpus/services/**` or `corpus/architecture/**`, so it enters neither `N` nor `P`.

### r27-G B1
`r27-G B1 | corpus/architecture/shared_libraries.md:128-130, :137, :163-164 | REJECTED | IN-SCOPE | PREDICATE: (not established) — the ai library at v1.40.2 exposes ChatCompletionStream, a mistral constructor, and panicking Anthropic methods.`
   evidence: The section's **Module** row is `github.com/anthropos-work/ai` and its **Version pin** row is
   `v1.40.2` — the claim names its subject *and* its version, so it is settled there (rule 1). The module is
   in no clone set, but it is readable: `app`'s fold commit `1e457fa70` states in its own body *"Source
   identity: merged from tag **v1.40.2** (`df05d720…`) — **byte-for-byte** the version `go.mod:14` pinned …
   The copy is byte-identical to the tag modulo import paths … no conversion of the six
   `panic("not implemented")` stubs."* Read at that commit, all three booked statements hold exactly:
   `internal/ai/ai.go:8-18` declares `type AI interface` with **nine** methods **including
   `ChatCompletionStream`** (`:10`); `internal/ai/mistral/completion.go` exists with
   `func NewMistral(...)` at `:25` and `panic("not implemented")` on `ChatCompletion` `:40`,
   `ChatCompletionStream` `:44`, `CreateEmbeddings` `:48`, `CreateSpeech` `:52`, with `OCRProcess`
   implemented — *"Mistral — OCR only (chat/embeddings/speech `panic`)"* verbatim; and
   `internal/ai/anthropic/completion.go` panics on exactly `ChatCompletionStream` (`:179-181`) and
   `CreateSpeech` (`:187-190`), the two the corpus names, while `CreateEmbeddings` `:184` and `OCRProcess`
   `:194` return errors — the corpus names those two and only those two.
   class: ref-discipline — the seat graded a version-pinned description of the standalone module against
   `app`'s post-fold in-tree copy at `ad9f3c49`, which has since deleted the streaming path (`9048ce1b4`)
   and re-homed mistral (`2b3a65cf0`). The pin is a date, and the claim is true at it. (The seat's own
   reading #28 reached the opposite conclusion on the same passage and cleared it as *"not a defect today"*;
   the fold commit is the evidence that settles it.) The latent tension the seat identifies — that
   `:126`'s *"`app/internal/ai/` carries the library in-tree"* invites reading the surface table as
   describing the in-tree copy — is real and worth a future clarifying edit, but it is not a false predicate:
   the section states its module and its version pin in the two rows immediately above the table.

### r28-G B1
`r28-G B1 | corpus/services/studio-desk.md:38 and :46 | UPHELD | IN-SCOPE | PREDICATE: studio-desk's Express backend holds the GraphQL client and calls the platform GraphQL API.`
   evidence: measured at `studio-desk 41ee3575` (the passage names no ref, so graded at the checkout).
   `git grep -in graphql 41ee3575 -- 'src/*'` returns exactly **two** lines, both comments saying the
   opposite of the claim: `src/routes/skillpath.ts:374` *"We do NOT route this through the platform's
   `privateSkillPaths` GraphQL"* and `:405`. `git grep -n 8082 41ee3575 -- 'src/*'` → **0**. `src/index.ts`
   mounts exactly four API routers — `/api/dev` `:150`, `/api/ai` `:158`, `/api/skillpath` `:161`,
   `/api/youtube` `:164` — none of them GraphQL. Every `new GraphQLClient(...)` in the repo is in the
   **frontend**: `app/services/userService.ts:20`, `app/services/taxonomyService.ts:43`,
   `app/services/userPreferencesService.js:13`, `app/services/content/simulationContentService.js:325`, all
   fed by `app/services/config.ts:6` reading the **`VITE_`-prefixed, browser-baked**
   `VITE_GRAPHQL_ENDPOINT`. The backend's real remote dependency is Directus over REST —
   `src/routes/skillpath.ts:13` *"Data: All data stored in Directus CMS (REST API with Bearer token)"*, with
   `DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN` read at `:44-47` and `src/index.ts:303-310`. So the mermaid at `:46`
   asserts a non-existent edge (`Backend --> GraphQL`), omits the real one (`Frontend --> GraphQL`), and
   omits `Backend --> Directus`, routing Directus through `CMS` instead. **It is also a self-contradiction
   inside one file** (rule 5, and not a retraction — both sides are asserted as live): `:107` (*"Integrates
   directly with Directus … the skill-path **writes** … go to Directus as a `Bearer ${DIRECTUS_TOKEN}`
   static token (`src/routes/skillpath.ts`)"*), `:135` (*"Example from `app/services/graphql/`"*), `:140`
   (*"Configured via `VITE_GRAPHQL_ENDPOINT`"*) and `:283-285` all place the client and the Directus edge
   correctly. Both anchors are one predicate.

### r28-G B2
`r28-G B2 | corpus/services/studio-desk.md:107 | UPHELD | IN-SCOPE | PREDICATE: studio-desk declares GCLOUD_SERVICE_ACCOUNT at .env.example:120.`
   evidence: `git -C stack-demo/studio-desk grep -n GCLOUD_SERVICE_ACCOUNT 41ee3575` → `.env.example:**119**`
   (plus `CHANGELOG.md:158` and `terraform/main.tf:129`, nothing under `src/`). Printing the file's own
   numbered lines: `:117 YOUTUBE_API_KEY=…`, `:118 # Google Cloud service account JSON (used for YouTube
   Data API)`, `:119 GCLOUD_SERVICE_ACCOUNT=`, **`:120` empty**, `:121` the `LEGACY VARIABLES` banner. The
   file is 131 lines, so `:120` is in range and resolves — to nothing. This is a *point* citation offered as
   the evidence for the claim, and it lands on a blank line. The other two thirds of the same sentence are
   exact: `terraform/main.tf:129` is `"name": "GCLOUD_SERVICE_ACCOUNT"`, and no file under `src/` reads it.
   Not a historical anchor (rule 7 — the sentence is present-tense about the current file) and not a
   pointer-to-a-derivation (rule 6). The passage names no ref, so it is graded at the checkout, where it is
   false.

---

## PREDICATE ROLL-UP

Distinct upheld **in-scope** predicates, each with every anchor that collapses onto it:

```
P1 | graphql-wundergraph's own CLAUDE.md "Version Tracking" section is stale.
     anchors: r27-E B1 @ corpus/services/graphql-wundergraph.md:274

P2 | roadrunner/terraform/main.tf:19 was last touched at 87d8d44 (2026-06-19).
     anchors: r27-E B2 @ corpus/services/roadrunner.md:13-14

P3 | Storage's compose environment block at platform 0dab54d declares eight variables.
     anchors: r27-E B3 @ corpus/architecture/dependency_map.md:19,
              r28-E B2 @ corpus/architecture/dependency_map.md:19   (same seat, second reading — same anchor)

P4 | storage's StorageService Connect handler is declared in storage:sdk/storage/v1/service.go.
     anchors: r27-E B4 @ corpus/services/backend.md:37-41

P5 | studio-desk's Express backend holds the GraphQL client and calls the platform GraphQL API.
     anchors: r28-G B1 @ corpus/services/studio-desk.md:38,
              r28-G B1 @ corpus/services/studio-desk.md:46          (prose + mermaid, one booking)

P6 | studio-desk declares GCLOUD_SERVICE_ACCOUNT at .env.example:120.
     anchors: r28-G B2 @ corpus/services/studio-desk.md:107
```

Upheld but **out of scope**, recorded and not counted:

```
X1 | The ai module is imported as a private Go module and pulled at Docker build.
     anchors: r28-E B3 @ CLAUDE.md:248 (repository root)
```

**Anchor-vs-predicate note.** 8 upheld anchor-instances reduce to **6** distinct in-scope predicates plus 1
out-of-scope: P3 absorbs a same-anchor re-booking across the seat's two readings, and P5 absorbs a prose
sentence and the mermaid edge that restates it. No two bookings from *different* seats collapsed — seats E
and G share no predicate this reading.

**Rejections, by class.** One `ref-discipline` (r27-G B1 — a version-pinned library-surface description
graded against a post-fold in-tree copy; the pin is a date and the claim is true at `v1.40.2`, readable
byte-for-byte through the fold commit) and one `other` (r28-E B1 — an omission of a net-new second Atlas
pipeline, where every asserted clause verifies at the named ref). Neither is a wrong-tree rejection: no
booked blocker in my partition rests on a `rosetta-extensions` claim, so the instrument's line-37 defect did
not bite here.

BOOKED=10 UPHELD=8 REJECTED=2 IN-SCOPE-UPHELD-BLOCKERS=7 DISTINCT-IN-SCOPE-PREDICATES=6
