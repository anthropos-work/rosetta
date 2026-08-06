# Adjudicator 1 — M257x iter-101

Seats adjudicated: `r23-A`, `r24-A` (external_services · chronos · messenger · frontend_architecture ·
gotenberg) and `r23-C`, `r24-C` (alignment_testing · backend · ai_architecture · customerio-sync ·
skillpath · skiller). 10 blockers booked, 10 adjudicated.

**Refs re-verified at open** (`git rev-parse --short=8 HEAD` per tree):
platform `0c91421d` · app `b948604f` · app/studio `aeec036a` · cms `ca50c817` · cms/studio `aeec036a` ·
next-web-app `bb3313bc` · sentinel `88bc5592` · storage `4ce8ece5` · messenger `fa47850d` ·
graphql-wundergraph `60c229f3` · roadrunner `87d8d443` · jobsimulation `462343b0` · studio-desk `14a5442a` ·
ant-academy `9c3843cd` · **`stack-demo/rosetta-extensions` `ab81527a` (pinned consumption clone)** ·
`.agentspace/rosetta-extensions` `09d06070` (authoring copy).

Every verdict below was re-derived by opening the platform file myself. Nothing was taken from a seat's
evidence, and no other seat report, ledger or verdict was read.

---

## Verdicts

### r23-A B1 | `corpus/architecture/frontend_architecture.md:39` | **UPHELD** | IN-SCOPE | "29 direct REST/SSE calls" counts env-var mentions, not calls

   evidence: I derived the SET first, not the sum. `git grep -n NEXT_PUBLIC_BACKEND_API_URL bb3313bc -- '*.ts' '*.tsx'`
   minus `.test.ts`/`e2e/` = **29 lines across 21 files** — the published pair reproduces byte-exactly under
   the predicate *"occurrences of the literal string"*. That is not the set the sentence names. Its
   grammatical subject is unambiguous: *"there are direct REST/SSE calls, **29 of them**"*.
   Measured over those same 21 files at `bb3313bc`, `fetch(` / `new EventSource(` totals **43**.
   The four `packages/core-js` clients the sentence sizes at *"12 sites between them"* carry **3 mentions
   each** and **25 outbound calls between them**: `coursebuilder/api.ts` 3 mentions / **18** `fetch(`,
   `credits/api.ts` 3 / **3**, `talkToData/api.ts` 3 / **3**, `workforce/api.ts` 3 / **1**. The counted
   quantity is not even monotone in the claimed one. Worse, 8 of those 12 lines are not code at all —
   `coursebuilder/api.ts:2` is a header comment and `:26` an `'… is not set — required by @anthropos/core-js/…'`
   error string; only `:23` (`const BACKEND_URL = process.env.…`) is executable, and it is a read, not a call.
   The sentence exists to repair an undercount (it says so: *"The long-standing '~15 sites' undercounted by
   half"*) and is the corpus's only sizing of the non-GraphQL frontend surface; it undercounts that surface
   again, in the same direction, by ~33 %.
   tree-read: `stack-demo/next-web-app` @ `bb3313bc` (the ref the sentence pins).
   note: reading #24 **cleared** this same anchor as ENUMERATED (its item 9) by re-deriving the env-var
   occurrence count and finding 29/21 and the forward-looking 31/22 both exact. That is rule 4 in its pure
   form — the arithmetic is right and the predicate is wrong. The clearance does not survive; the booking does.

### r24-A B1 | `corpus/architecture/external_services.md:136` vs `:206`/`:210` | **UPHELD** | IN-SCOPE | Same file asserts the `--local-content` re-point has ONE target and BOTH cms+backend

   evidence: `:136` — *"`backend` is the only consumer left and per-service re-point tooling has **ONE
   target, not two** — see the ⚠️ under *Architecture* below."* `:206` (bolded headline) — *"**The
   `--local-content` re-point targets BOTH `cms` and `backend`.**"*, restated at `:210` as the tuple
   *"`("cms", "backend")`"*. Incompatible cardinalities about the same mechanism, 70 lines apart, and `:136`
   routes the reader to `:206` **as its own corroboration**, so the cross-reference lands on the refutation.
   Ground truth settles which half is literally right without settling the contradiction:
   `stack-injection/gen_injected_override.py:84` is `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` (two
   members) and `:669-670` is `if with_directus and name in DIRECTUS_DATA_CONSUMERS: env.append(f"DIRECTUS_BASE_ADDR=…")`.
   `:136` is defensible only about *effect* — `git show 0c91421d:docker-compose.yml` declares exactly
   `sentinel`(`:5`), `backend`(`:28`), `studio-desk`(`:112`), `next-web-app`(`:143`), `gotenberg`(`:170`), so
   `cms` never matches `name` on a live stack. Both halves are defensible in isolation; the document asserts
   both and reconciles neither. Upheld under the self-contradiction rule.
   tree-read: **`stack-demo/rosetta-extensions` @ `ab81527a`** — the pinned per-stack consumption clone, which
   is what settles a claim about the tooling a stack runs — plus `stack-demo/platform` @ `0c91421d`.

### r24-A B2 | `corpus/services/chronos.md:27` | **REJECTED** | IN-SCOPE | Ports `8080`/`8081` said to be false because compose ran chronos on `8500`/`8501`

   evidence: the compose half of the booking is real and I reproduced it — at every commit that ever touched
   the block (`3330ef6`, `081a350`, `b43b99a`, `045857c^`) the chronos service published `"8500:8500"` /
   `"8501:8501"` with `PORT=8500` / `RPC_PORT=8501`, and the consumer side agrees
   (`045857c^:docker-compose.yml:177` `CHRONOS_RPC_ADDR=http://chronos:8501`). But that does not refute the
   bullet. The corpus's own convention for a decommissioned colony service's **Ports** field is to give the
   binary's built-in defaults — `jobsimulation.md:96`, `roadrunner.md:77` and `cms.md:113` each state
   *"**8080 …, 8081 … — the binary's own defaults**"* and then record the compose pair as historical, and
   `messenger.md:62` states the same defaults explicitly. I verified those defaults in source rather than in
   the corpus: `messenger/cmd/root.go:63-64`, `jobsimulation/cmd/root.go:77-78`, `cms/cmd/root.go:77-78`,
   `roadrunner/cmd/root.go:84`/`:110` are all `cmp.Or(os.Getenv("PORT"), "8080")` / `cmp.Or(os.Getenv("RPC_PORT"), "8081")`,
   and `messenger/cmd/root.go:107` is `cmp.Or(os.Getenv("REDIS_STREAMS_INDEX"), "2")` — the third number the
   booking disputes, matching chronos.md's `2` exactly. Chronos is a colony service of the same generation;
   `8080`/`8081`/`2` is the colony default in four independently-read repos. The doc's env table column is
   headed **"Example"**, and its only runnable recipe is standalone `go run .` (`:163-179`), where defaults are
   the operative ports. The residual — chronos.md never records the historical `8500`/`8501` pair its five
   sibling docs all record — is an **omission**, not a false claim.
   tree-read: `stack-demo/platform` @ history reachable from `0c91421d`; `stack-demo/{messenger,jobsimulation,cms,roadrunner}`
   at their own HEADs. (`chronos` is in no clone set — measured.)
   class: mis-read — a compose-supplied value does not refute a bullet stating the binary's default, which is
   what the corpus's own convention puts in that field; the seat itself rated confidence **low**.

### r23-C B1 | `corpus/services/backend.md:19` | **REJECTED** | IN-SCOPE | "compose now declares five services" said to be false because the project resolves seven

   evidence: `git show 0c91421d:docker-compose.yml` opens `include:` / `- common.yml` (`:1-2`) and declares
   exactly five services — `sentinel:5`, `backend:28`, `studio-desk:112`, `next-web-app:143`, `gotenberg:170`;
   `common.yml` adds `postgresql:2` and `redis:24`, so the resolved **project** is seven. Both halves of the
   booking are measured correctly. The verdict turns on what "compose declares" names. The sentence is the
   trailing clause of *"The last three lost their **containers** a day later, at platform `838d907`"* — a
   statement about the delta in the compose **file**, paired in the same breath with *"`repos.yml` **four**
   entries"* (which I confirmed: `app`, `sentinel`, `next-web-app`, `studio-desk`), i.e. the same file-level
   register. The scalar is exact for `docker-compose.yml`. The corpus's canonical form carries the qualifier
   — `service_taxonomy.md:63-65`: *"**`docker-compose.yml` declares five services (seven in the effective
   topology, once `include: common.yml` adds the `postgresql`/`redis` floor)**"* — so the corpus knows and
   states the distinction; `backend.md:19` drops the parenthetical. Nor is it a self-contradiction:
   `backend.md:276` (*"also starts postgresql, redis, sentinel, gotenberg"*) and `customerio-sync.md:104` are
   both about what a profile **starts**, and the `include:` resolves the tension rather than creating one.
   The missing qualifier is a real defect of the MINOR class — which is exactly the grade the **other**
   reading of this same file gave it.
   tree-read: `stack-demo/platform` @ `0c91421d`.
   class: mis-read — the scalar is correct for the file the cited commit edited; an omitted qualifier is not a
   false claim.

### r23-C B2 | `corpus/services/backend.md:301` | **UPHELD** | IN-SCOPE | "the most recent set of migrations (May 2026)" — the newest are 2026-07-31

   evidence: the sentence names no ref, so it grades at the checkout. `git ls-tree --name-only b948604f
   terraform/migrations/` holds 170 migrations + `atlas.sum`; the newest six are
   `20260724132049_cms_data_model.sql`, `20260724164346_ai_readiness_freeze_how_we_measure.sql`,
   `20260728103254_ai_readiness_snapshot_frozen_matched_sources.sql`, `20260729133514.sql`,
   `20260731131307.sql`, `20260731154527_academy_chapter_progress_completed_at.sql` — i.e. **2026-07-31**,
   two months past the claim, and the gap only widens at `origin/main`. Six May-2026 migrations exist
   (`20260505133528` … `20260529072659_add_lab_session`) but they are not the most recent set. The second
   half fails too: I opened the newest three. `20260731131307.sql` is
   *"Modify "course_builder_sessions" table … ADD COLUMN "brief" … "credits_spent""*;
   `20260731154527_…` adds `completed_at` to `academy_chapter_progresses` with a backfill;
   `20260729133514.sql` is *"Collapse the local_* session mirrors (M709c + skill-path equivalent)"*. None
   touches *"simulation-type definitions and content JSON defaults"*. The corpus contradicts itself here too:
   `skillpath.md:85-92` cites `20260729133514.sql:62-63` and names three migrations that post-date it.
   A reader trusting `:301` concludes the schema has been static since May.
   tree-read: `stack-demo/app` @ `b948604f` (and cross-checked at `origin/main 2035f9a4`).

### r23-C B3 | `corpus/services/backend.md:33-34` | **UPHELD** | IN-SCOPE | App-owned stream set omits `backend`, so "skiller is NOT a fifth member" is a wrong ordinal

   evidence: enumerated every publisher and subscriber in Go source at `b948604f`.
   Publishers: `main.go:287` `pubsub.NewPublisher(serviceName, …)`, `:637` `SKILLPATH_STREAM`, `:1039`
   `CMS_STREAM`, `jobsimwiring/wiring.go:127` `AI_USAGE_STREAM`, `:180` `JOBSIMULATION_STREAM` — five.
   Subscribers: `main.go:1274`, `:1276` (`SKILLER_STREAM`), `:1285`, `:1303`, `:1305`, `:1320`
   `subServer.AddSubscriber(serviceName, backendSelfSub)` — six. `serviceName` is `os.Getenv("SERVICE_NAME")`
   defaulting to `"backend"` (`main.go:213-215`), and compose sets `SERVICE_NAME=backend`
   (`docker-compose.yml:70` @ `0c91421d`). So `app` is **both producer and consumer of the `backend` stream**,
   and the producer+consumer set is **five** — `backend`, `skillpath`, `jobsimulation`, `cms`, `ai_usage`;
   `skiller` would be a **sixth**, not a fifth. `:33` states the set as a bolded four-member equality and then
   makes the cardinality explicit and load-bearing (*"NOT a fifth member"*). The same file refutes it 231
   lines later at `:264`: *"app is both producer and consumer of **four of the five** application streams —
   **`backend`**, `skillpath`, `jobsimulation`, `cms` — plus the `AI`/`ai_usage` usage stream"* — and `:34`
   cites `:264` as its own authority. Material because the same bullet warns that *"a second `AddSubscriber`
   on one stream overwrites the first"*: the reader is handed an incomplete inventory of exactly the thing the
   warning is about.
   tree-read: `stack-demo/app` @ `b948604f`; `stack-demo/platform` @ `0c91421d`.

### r23-C B4 | `corpus/services/skiller.md:19` | **REJECTED** | IN-SCOPE | `skiller_rpc_addr = http://backend.internal.anthropos:8081` said to name a non-existent identifier

   evidence: the booking's decisive half — *"the identifier is wrong … There is no `skiller_rpc_addr` variable
   in `app`, `messenger`, `cms`, `jobsimulation`, `storage` or `sentinel` (searched all six at HEAD)"* — is
   measurably false. `git -C cms grep -n skiller_rpc_addr HEAD -- '*.tf'` → `terraform/variables.tf:76`
   `variable "skiller_rpc_addr" {` and `terraform/main.tf:107` `"value": "${var.skiller_rpc_addr}"`;
   `git -C jobsimulation grep …` → `terraform/variables.tf:163` and `terraform/main.tf:121`. Two of the six
   trees the seat says it searched carry the exact identifier at their own HEADs. Only `messenger` spells it
   `skiller_rpc_address` (`variables.tf:87`), which is the single instance the seat generalised from.
   The value is corroborated too: `app/knowledge/service-dependencies.md:46` gives
   `http://backend.internal.anthropos:8081`, and the namespace is independently confirmed outside `app` at
   `graphql-wundergraph@60c229f3:supergraph-config-prod.yaml:6`
   (`routing_url: http://backend.internal.anthropos:8080/graphql/query`). Nor does `backend.md:118-121`
   contradict `skiller.md:19` — it asserts the *same* address (`:112`) and adds that the literal is set in the
   un-cloned `infrastructure` root module because every cloned module declares the variable with no default
   (which I confirmed: both declarations above are default-less). "In production terraform" is where the
   assignment is; it is simply in the repo we cannot clone.
   tree-read: `stack-demo/{cms,jobsimulation,messenger,app,storage,sentinel,roadrunner,graphql-wundergraph}`
   at their own HEADs.
   class: mis-read — the seat's own absence measurement failed; the identifier exists in two of the trees it
   swept, and the qualification it wanted is already present in the corpus's canonical statement.

### r23-C B5 | `corpus/architecture/ai_architecture.md:35` | **REJECTED** | IN-SCOPE | `mistral-ocr-latest` parenthetical anchored at `pdf2md.py:24`; the literal is at `:127`

   evidence: at `app/studio`'s **own** ref `aeec036a`, `tools/pdf2md.py:24` is `from mistralai import Mistral`
   and the literal `mistral-ocr-latest` is at `:127` (`model="mistral-ocr-latest",` inside
   `client.ocr.process(`) — its only occurrence in the tree. Both measurements hold. But `:24` is not the wrong
   construct for the proposition the sentence makes. The sentence's anchor convention is *use-sites*, not
   model literals, and its two Go anchors prove it: `internal/cms/studio/markdownManager.go:19` is
   `ai, err := mistral.NewMistral(nil, os.Getenv("MISTRAL_API_KEY"))` and `studioManager.go:583` is
   `markdownManager, err := NewMarkdownManager(os.Getenv("MISTRAL_API_KEY"))` — neither carries a model
   literal either. `:24` is the Python-side counterpart. The parenthetical names the model in play; the
   substantive proposition (*"Every use of Mistral in `app` is OCR, never generation"*) is true and
   independently verified. A 103-line drift on a parenthetical inside the right file is the MINOR class —
   which is the grade the other reading of this same file gave it (its `:36` is itself one line off; the
   sentence is at `:35`).
   tree-read: `stack-demo/app/studio` @ `aeec036a` (the nested checkout, grepped at its own ref);
   `stack-demo/app` @ `b948604f`.
   class: mis-read — the anchor resolves to a real Mistral use consistent with the sentence's own convention;
   the claim it supports is true.

### r24-C B1 | `corpus/services/backend.md:29-30` | **UPHELD** | IN-SCOPE | "All of their Connect-RPC surfaces are served on `app`'s single RPC mux" — half the eight are not

   evidence: `"their"` is the eight-row table headed *"**Eight former microservices now run inside `app`**"*
   (`:5`), and the neighbouring bullets use the same referent (`:27` *"All of their tables live in `public`"*).
   I enumerated the mux at both legitimate refs. There is exactly **one** `http.NewServeMux()` in `main.go`
   (`:1185` @ `b948604f`), passed to `rpc.NewServer(mux, cfg.RPCPort, …)` at `:1243`, and every `mux.Handle`
   on it is: `usersv1connect.NewUsersServiceHandler` (`:1187`), `organizationsv1connect.NewOrganizationsServiceHandler`
   (`:1188`), `skillerv1connect.NewSkillerServiceHandler` (`:1196`), `jobsimulationv1connect.NewJobSimulationServiceHandler`
   (`:1204`), `cmsv1connect.NewCMSServiceHandler` (conditional, `:1213`), `labv1connect.NewLabSessionServiceHandler`
   (`:1231`). At `origin/main 2035f9a4` — after the storage/messenger/customerio-sync fold — the set is
   **identical** (`:1297`, `:1298`, `:1306`, `:1314`, `:1323`, `:1338`), and a tree-wide sweep for
   `New[A-Za-z]+ServiceHandler\(` at that ref returns no other registration anywhere in Go source (the only
   other hits are four comment lines in `internal/messenger/adapters/doc.go` + `handlers.go` and
   `internal/storagens/callsites_test.go`). So of the eight, **skillpath, roadrunner, storage, messenger and
   customerio-sync have no handler on the mux at all** — only skiller, jobsimulation and cms do.
   `SkillPathSessionService` = **0** occurrences in Go source at both refs.
   The same document refutes the banner three times: `:66` correctly enumerates the mux as five unconditional
   handlers plus conditional `CMSService`; `:68` states *"**There is no `SkillPathSessionService`** — measured:
   **0** occurrences in Go source"*; `:254` says object storage is *"in-process … not a service hop"*; `:256`
   says Judge0 is *"called directly (`JUDGE0_BASE_URL`)"*. `skillpath.md:30` puts it strongest: *"It was not
   re-hosted in `app`; it was DROPPED."* The bullet's second clause (*"nothing outside the process calls
   them"*) is correct, which is what lends the false first clause its authority.
   tree-read: `stack-demo/app` @ `b948604f` **and** `origin/main 2035f9a4`.

### r24-C B2 | `corpus/architecture/ai_architecture.md:141` | **REJECTED** | IN-SCOPE | Embeddings counts said to be derived by `organization_id IS NULL`, a column neither table has

   evidence: the column half is true and I confirmed it three ways —
   `app/internal/data/ent/schema/skill_embeddings.go` declares one field (`skill_id`) under a doc comment
   saying *"Ported GLOBAL from skiller … no UserMixin/OrganizationMixin and no Policy(), so the row is not
   org-filtered"*; the checked-in capture manifest records the columns as exactly
   `["id","small_embedding3","skill_id"]` / `["id","small_embedding3","job_role_id"]`; and the doc's own DDL
   block at `:157-169` prints three columns each. But the sentence does not state that filter *on the
   embeddings tables*. It reads *"…rows **over the public taxonomy** (`organization_id IS NULL`, measured
   2026-06-29)"* — the parenthetical is appositive to **"the public taxonomy"**, a defined term the corpus
   uses identically where the column does exist (`backend.md:96` *"Public predicate `organization_id IS NULL`
   (the public taxonomy…)"*, `skiller.md:44-46`), and it is compound with a provenance stamp
   ("measured 2026-06-29"), i.e. a gloss and not a SQL predicate. The set so named is exactly the set the
   capture ranged over: `.agentspace/snapshots/taxonomy/5afc0bcc…/manifest.json` filters
   `"skill_id" IN (SELECT id FROM "public"."skills" WHERE organization_id IS NULL)` with
   `"public_via": "public.skills"` → `row_count 42790`, and the job-role twin → `18919`. Both published
   numbers reproduce exactly, over the correctly-named set. The `backend.md:99` "collision" is not one either:
   that figure is dated **2026-07-08** and sits beside *"42,790 (43,584 total incl. 794 org-private)"*, i.e.
   whole table vs public subset, and `:100-101` explicitly reconciles the drift against the 2026-06-29 capture.
   The residual — the derivation is not directly runnable on the embeddings table without the join — is a
   MINOR-grade imprecision, which is how the other reading of this same file graded the neighbouring
   anchors.
   tree-read: `stack-demo/app` @ `b948604f`; `.agentspace/snapshots/taxonomy/5afc0bccf1df7ef538b643321fc6362f/manifest.json`
   (read-only).
   class: mis-read — the predicate names the parent-entity set, which is the set the numbers were measured
   over; the counts and the set both re-derive.

---

## DEDUPLICATION

**No two BLOCKERS collapse onto one predicate.** All ten bookings assert distinct propositions at distinct
anchors, so **DISTINCT-PREDICATES = 10**. Checked pairwise; the near-misses and why they do not collapse:

- `backend.md:29-30` (r24-C B1, the **RPC mux** banner bullet) and `backend.md:33-34` (r23-C B3, the
  **Redis-stream** banner bullet) are consecutive bullets of the same banner and both are upheld, but they
  are different mechanisms and different false propositions. **Two predicates, two anchors.**
- `frontend_architecture.md:39` (r23-A B1) has no twin: reading #24 opened the identical anchor and
  **cleared** it. Same anchor, opposite verdicts — not a duplicate, a severity disagreement, resolved above
  in favour of the booking.
- `ai_architecture.md:35` (r23-C B5) and `ai_architecture.md:141` (r24-C B2) are both `ai_architecture.md`
  bookings and both rejected, but the first is an anchor-drift claim about `pdf2md.py` and the second a
  set-derivation claim about embeddings row counts. **Two predicates.**

**Cross-reading severity disagreements** (same predicate, one seat BLOCKER / the other MINOR — recorded
because they are the shape of a duplicate without being one, and all three resolve the same way I did):

| predicate | booked BLOCKER | booked MINOR | my verdict |
|---|---|---|---|
| `backend.md:19` five-vs-seven compose services | r23-C B1 | r24-C minor #1 | REJECTED (minor is the right grade) |
| `backend.md:33-34` app-owned stream set drops `backend` | r23-C B3 | r24-C minor #2 | UPHELD (explicit wrong cardinality) |
| `ai_architecture.md:35` `pdf2md.py:24` parenthetical | r23-C B5 | r24-C minor (`:36`) | REJECTED (minor is the right grade) |

`WRONG-TREE-REJECTIONS = 0`. None of the five rejections was graded against the wrong checkout: r23-C B5 was
correctly measured at the nested `app/studio @ aeec036a`, r24-A B1 correctly at the pinned
`stack-demo/rosetta-extensions @ ab81527a`, and r23-C B4's failure was an absence measurement that missed
hits present in the very trees it named, not a ref error.

---

`BOOKED=10 UPHELD=5 REJECTED=5 IN-SCOPE-UPHELD-BLOCKERS=5 DISTINCT-PREDICATES=10 WRONG-TREE-REJECTIONS=0`
