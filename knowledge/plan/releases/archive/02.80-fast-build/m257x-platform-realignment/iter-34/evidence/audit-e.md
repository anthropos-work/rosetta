# Audit E — Clause-5 confirming pass (iter-34)

Platform clone read at `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform` (origin HEAD `2adcf71`);
sibling repos in `stack-demo/` — `app` @ `5ba17044` (v1.363.2), `next-web-app` @ `bb3313bc0` (v2.133.0),
`ant-academy` @ `9c3843cd` (v2.34.2). Read-only throughout; no file outside this report was created or modified.

## Positive control

| File | `wc -l` | Last line actually read |
|------|--------:|------------------------:|
| `corpus/architecture/shared_libraries.md` | 219 | 219 |
| `corpus/architecture/security_compliance.md` | 173 | 173 |
| `corpus/services/storage.md` | 175 | 175 |
| `corpus/architecture/ai_architecture.md` | 180 | 180 |
| `corpus/services/academy-backend.md` | 139 | 139 |
| `corpus/services/coursebuilder.md` | 132 | 132 |
| `corpus/services/messenger.md` | 128 | 128 |
| `corpus/services/clerk-integration.md` | 128 | 128 |
| `corpus/services/next-web-app.md` | 126 | 126 |
| `corpus/services/askengine.md` | 121 | 121 |
| `corpus/architecture/frontend_architecture.md` | 105 | 105 |
| `corpus/services/skillpath.md` | 92 | 92 |

All 12 files read top to bottom, in full. **0 UNREAD.**

---

## BLOCKERS

### B-E1 — `corpus/services/skillpath.md:68-75` — the manager-view "mirror table" no longer exists

**Verbatim (lines 68-74):**

> *"**The manager view reads an `app`-side MIRROR, not this runtime.** The **manager insights** surface
> (`insightsSkillPathByMemberships`, the `/enterprise/activity-dashboard/@tabs/skill-paths/[skillPathId]`
> scoreboard in `apps/web`) does **not** read the runtime session — it reads the `app`-side mirror table
> **`public.local_skill_path_session`** (`app/internal/organization/intelligence.go`; Ent schema
> `app/internal/data/ent/schema/local_skill_path_session.go` — `progress` 0-100, `status`, no `score`) …
> **Seeding only the runtime session rows renders an empty manager scoreboard** — the mirror row must be
> co-written."*

**What is actually true at HEAD:** the mirror tables were **DROPPED**, and the manager query now reads the
runtime session table directly.

- `app/terraform/migrations/20260729133514.sql:61-62` — `DROP TABLE "local_jobsimulation_sessions";` /
  `DROP TABLE "local_skill_path_sessions";` (with the trigger-function cleanup at `:63`). This is the **last**
  migration in `terraform/migrations/`; nothing re-creates them.
- There is **no** `local_skill_path_session.go` (nor `local_jobsimulation_session.go`) in
  `app/internal/data/ent/schema/` — `ls … | grep -i local` returns nothing across all 139 files.
- `app/internal/organization/intelligence.go:1159` — `InsightsSkillPathByMemberships` builds
  `m.ent.SkillPathSession.Query()` filtered on `skillpathsession.SkillPathID/StatusIn/HasUserWith`. It reads
  **`public.skill_path_sessions`**, the runtime table.

**Why it misdirects real work:** the operational instruction is now exactly inverted. Seeding tooling told to
"co-write the mirror row" will attempt an INSERT into a table that no longer exists (hard failure), and will
be told that seeding the runtime rows — which is now the *only* thing that populates this scoreboard — is
insufficient. (Note also the table name given is singular `local_skill_path_session`; the dropped table was
plural `local_skill_path_sessions`.)

**Grade: BLOCKER.**
**Suggested correction:** replace the bullet with: the manager insights surface reads the runtime
`public.skill_path_sessions` directly (`intelligence.go:1144-1186`); the `local_skill_path_sessions` /
`local_jobsimulation_sessions` mirrors were dropped at `20260729133514.sql:61-62` and there is no mirror row
to co-write. Keep the `apps/hiring` no-surface note and the "Coming soon" per-member note (both verified true —
see below).

---

### B-E2 — `corpus/architecture/security_compliance.md:70-72` and `:80` — the tenant-isolation fence names a *policed* schema as its flagship *unpoliced* example

**Verbatim (lines 70-73):**

> *"**and a further ~18 declare a plain `organization_id` field with no mixin and no policy at all**
> (`org_membership.go`, `org_subscription.go`, `organization_settings.go`, … ). **Those are the rows most
> likely to be missed by an audit**: they look org-scoped and are not policed."*

**and line 80:**

> *"- Ent privacy policies auto-filter by organization **only on the 30 schemas using `OrganizationMixin{}`**"*

**What is actually true at HEAD:** `org_membership.go` — the **first** name in the list — declares its own
org-scoping, fail-closed privacy policy:

```
app/internal/data/ent/schema/org_membership.go:172-188
func (Membership) Policy() ent.Policy {
    Mutation: { OnMutationOperation(rule.DenyMismatchedOrganization(), ent.OpCreate),
                rule.AllowCurrentOrgEdgesOrSkipRule(), privacy.AlwaysDenyRule() },
    Query:    { rule.AllowCurrentOrgEdgesOrSkipRule(),
                rule.AllowCurrentUserOwnedOrgEdgesOrSkipRule(), privacy.AlwaysDenyRule() },
}
```

`Membership` is not merely policed — it is the **only** entity in that list of 18 that *default-denies*
(`AlwaysDenyRule`), while the OrganizationMixin schemas end in `AlwaysAllowRule` (`mixin.go:126-142`).

Counted independently this session over `app/internal/data/ent/schema/` (139 `.go` files):
- `OrganizationMixin{}`: **30** files (the doc's 30 ✓, `mixin.go:126` ✓).
- `OrganizationIDMixin{}`: **7** files (`category, jobrole, similarity, skill, specialization, studio_task,
  studio_document`) — no policy, confirmed at `skiller_mixins.go:148-153` ✓.
- Plain `field.UUID("organization_id")` with neither mixin: **18** schema files. Of those, **17 have no
  `Policy()`; `org_membership.go` has one.**
- Schemas declaring a hand-rolled `Policy()` outside the mixins: exactly **3** — `org_membership.go`,
  `organization.go`, `user.go`.

So line 80's "**only** on the 30 schemas using `OrganizationMixin{}`" is false: **31** schemas auto-filter by
organization at the ORM (the 30 + `Membership`).

**Why it misdirects real work:** this is the platform's tenant-isolation fence, and it has now been wrong three
times. An engineer acting on it would add caller-side scoping to membership reads that already fail closed at
the ORM — and would be baffled when an unscoped `Membership` query returns nothing (`AlwaysDenyRule`) rather
than leaking. It also, again, hands a reviewer a self-contradicting fence: the flagship example refutes the
sentence it illustrates. (Direction note: unlike the two prior versions, this error errs toward *understating*
isolation, not overstating it.)

**Grade: BLOCKER.**
**Suggested correction:** "…and a further **17** declare a plain `organization_id` field with no mixin and no
policy at all (`org_subscription.go`, `organization_settings.go`, `organization_feature.go`, `api_key.go`,
`lab_session.go`, `interview_aggregated_report.go`, `admin_audit_log.go`, `job_simulation_session.go`, …).
**`org_membership.go` is the exception and is *not* one of them** — `Membership` declares its own fail-closed
org policy (`org_membership.go:172-188`), as do `organization.go` and `user.go`." And at line 80: "auto-filter
by organization on the **31** schemas that carry an org policy — the 30 using `OrganizationMixin{}` plus
`Membership`'s hand-rolled one."

---

### B-E3 — `corpus/services/clerk-integration.md:103` — the "SDK versions are aligned" note is false for `@clerk/clerk-expo`

**Verbatim:**

> *"**SDK versions:** the JS Clerk SDKs are **aligned** across `next-web-app` and `ant-academy` — both on
> `@clerk/nextjs ^6.39.2` and `@clerk/clerk-expo ~2.6.18`."*

**What is actually true at HEAD:**

| Declared in | Package | Version |
|---|---|---|
| `next-web-app/apps/web/package.json:10`, `apps/hiring:10`, `apps/integration:9` | `@clerk/nextjs` | `^6.39.2` ✓ |
| `ant-academy/code/package.json:52` | `@clerk/nextjs` | `^6.39.2` ✓ |
| `next-web-app/apps/mobile/package.json:6` | `@clerk/clerk-expo` | `~2.6.18` |
| **`ant-academy/mobile/package.json:18`** | `@clerk/clerk-expo` | **`~2.19.36`** |

The `nextjs` half is aligned; the `clerk-expo` half is **13 minor versions apart**, not aligned.

**Why it misdirects real work:** the paragraph exists specifically to tell the reader which pins must be
re-verified before trusting a Clerkenstein alignment score (`CHECK-M257x-iter22-clerk-sdk-drift`). It asserts
alignment on the one axis that has actually drifted, so the re-verification it prescribes would skip the
drifted pin.

**Grade: BLOCKER.**
**Suggested correction:** "`@clerk/nextjs` is aligned at `^6.39.2` across `next-web-app` and `ant-academy`.
`@clerk/clerk-expo` is **not**: `next-web-app/apps/mobile` pins `~2.6.18`, `ant-academy/mobile` pins
`~2.19.36`. The Go side has drifted too: `app/go.mod:31` @ `5ba17044` reads `clerk-sdk-go/v2 v2.7.0`…"
(the `v2.7.0` claim itself verified exact at `app/go.mod:31`).

---

## Minors

| # | Anchor | Issue | Evidence |
|---|--------|-------|----------|
| m-E1 | `security_compliance.md:67` | *"of **139** Ent schemas"* — 139 is the **file** count of `app/internal/data/ent/schema/*.go`. Only **135** declare `ent.Schema`; `mixin.go`, `database_types.go`, `skiller_mixins.go`, `skillpath_mixins.go` are helper files. Ratio impact negligible; the number is presented as a schema count. | `ls *.go \| wc -l` = 139; `grep -l "ent.Schema" *.go \| wc -l` = 135 |
| m-E2 | `frontend_architecture.md:39` and `:44` | *"the Cosmo Router at `:5050/graphql` in prod"* (twice). `:5050` was only ever the **local** compose host-port of the now-deleted `graphql` service. Prod is `https://gql.anthropos.work/graphql/query`. The local half of both sentences (`backend` at `:8082/graphql/query`) is correct. | `platform/CLAUDE.md` "No router… (prod `https://gql.anthropos.work/graphql/query`)"; no `5050` anywhere in `docker-compose.yml` |
| m-E3 | `messenger.md:110` | Anchor `internal/flow/assignments.go:815` cited for *"skill-path data is read via the CMS client"*. Line 815 is inside `getEmailNotificationForSimulation` (a `GetEmailNotifications` call). The actual `cms.GetSkillPath` call is at `:828`. | `messenger/internal/flow/assignments.go:828` |
| m-E4 | `ai_architecture.md:141` | *"Both recordings are stored in S3 and linked to the simulation session."* True for the write path, but the Chime video that actually **renders** is resolved by reference — `ChimeRecording.bunny_video_id` (`app/internal/data/ent/schema/chime_recording.go:25,42`) → a Bunny.net Stream pull-zone. The corpus treats that as load-bearing (`ops/demo/media-substrate-spec.md`); this section never mentions Bunny, so a reader hunting the video bytes is pointed only at S3. | `chime_recording.go:20-27` |
| m-E5 | `ai_architecture.md:119-128` | The voice-engine inventory (active `livekitgptrealtime`; legacy `elevenlabs`, `gptrealtime`) omits the 4th enum value **`livekitchain`** — which is the engine matching the `livekit-agent-chain` repo the same doc newly names at `:110`. | `app/internal/web/backend/graphql/graph/schemas/cms_simulation.graphqls:257-263` (`enum SimulationVoiceEngine { gptrealtime elevenlabs livekitgptrealtime livekitchain }`) |
| m-E6 | `next-web-app.md:124` | Related-doc line reads *"[GraphQL Gateway](./graphql-wundergraph.md) — the federated endpoint this app consumes"*. Not a dead link (file exists), but stale against `:14`, `:47` and `:96` of the same doc, which correctly state the app now consumes `backend:8082/graphql/query` locally. | `corpus/services/graphql-wundergraph.md` exists; `docker-compose.yml:352,361` |

---

## Verified-clean spot checks (recorded so they are not re-audited)

Everything below I checked against the clone and found **exact**, so it is deliberately *not* a finding:

- **`shared_libraries.md`** — every version pin: colony `app/messenger v0.35.2`, `cms/jobsimulation v0.35.1`,
  `sentinel/storage v0.34.3`; proto `1.210.0 / 1.207.0 / 1.205.0 / 1.200.0 / 1.196.0`; ai `v1.40.2`; taxonomy
  `v1.2.0` (indirect in sentinel/storage) — all read from the seven `go.mod`s. `repos.yml:14-19` = cms +
  jobsimulation ✓. `docker-compose.yml:83,144` = jobsimulation + cms build blocks ✓.
  `ROADRUNNER_RPC_ADDR` at `docker-compose.yml:118` (in the *jobsimulation* block, not backend's) ✓.
  **Six** Connect handlers in `app/main.go` — Users `:1178`, Organizations `:1179`, Skiller `:1187`,
  JobSimulation `:1195`, CMS `:1204`, LabSession `:1218-1219` — and **no** `SkillPathSessionService` or
  RoadRunner handler ✓. `skillpaths.go:27-29` "the drop-in for the **removed** skillpath RPC client" ✓.
  `jobsimwiring/wiring.go:118` = `jsrunner.NewRunnerManager(JUDGE0_*)` ✓. `app/internal/askengine/bedrock.go`,
  `app/internal/aiusage/ai_usage.go`, `app/cmd/createTaxonomy`, `app/internal/taxonomy` all exist ✓.
  The `graphql` **profile** still exists in compose (`:81,140,187,218,309,384`) even though the
  `graphql-wundergraph` **service** is gone — line 42's wording is correct on that point.
- **`storage.md`** — `docker-compose.yml:210` is exactly the `STORAGE_S3_PUBLIC_BUCKET=production-…` line and
  `:324` is inside the studio-desk block ✓. `storage` repo has **0** `*_test.go` files and `Dockerfile:18` is
  `RUN go test -v ./...` ✓. Go 1.25.0 ✓. Ports 8300/8301 ✓. `storage.go:121-123` empty-string presigned URL
  when `S3Bucket == ""` ✓. `app/main.go:983` `storage.NewClient(…, storagens.CMS)` ✓;
  `jobsimulation/recording/recording.go:12` and `…/anticheat.go:34` `storagev1` imports ✓.
- **`messenger.md`** — `rpcsrv.go:25-30` = the two `CodeUnimplemented` stubs ✓; `cmd/root.go:63/64/107/147` =
  `PORT` 8080 / `RPC_PORT` 8081 / `REDIS_STREAMS_INDEX` 2 / `READONLY_DB_CONNECTION` ✓; compose `:256` CMS,
  `:258` JOBSIMULATION, `:265` SKILLER ✓; `app/main.go:1199` = *"…until the M809 re-point"* ✓;
  `flow/flow.go:73-77` the five `OrgSkillPath*` handlers ✓; `flow/jobsimulations.go:142,148` the 2 h / 12 h
  staleness guards ✓; brevo `v1.1.3` + liquid `v1.8.1` ✓; `depends_on` = backend + cms + jobsimulation, no
  skillpath ✓; Go 1.25.0 ✓.
- **`askengine.md`** — `ask_conversations` / `ask_messages` / `ask_query_examples` / `ask_query_lessons` /
  `ask_auto_rules` table consts ✓; `DefaultModelID = "eu.anthropic.claude-sonnet-4-6"` (`bedrock.go:25`) ✓;
  `maxAgenticIterations = 15` / `loopTimeout = 10 * time.Minute` (`ask/handler.go:34,41`) ✓;
  `askEngineMaxConns = 6` (`main.go:149,335`) ✓; `rules.md` = **146 219 bytes** ("~146 KB") ✓;
  test counts **49 / 13 / 10 / 14** in sandbox/executor/prompt/followups — exact ✓; registry ~60 `TableDef`s ✓.
- **`coursebuilder.md`** — Go **1.26.4** (`app/go.mod:3`) ✓; migration `terraform/migrations/20260717151144.sql`
  exists ✓; `coursebuildersession.go` Ent schema ✓; 98 `.go` files in `internal/coursebuilder/` ("~100") ✓;
  `internal/web/backend/coursebuilder/handler.go` + the named engine files all present ✓;
  `cmd/coursebuilder-{e2e,liverun}` ✓; app `v1.363.2` @ `5ba17044` ✓.
- **`academy-backend.md`** — all 11 `academy_*` Ent schemas present ✓; `internal/academy/` file set matches
  (`academy.go, content.go, body.go, content_import.go, embeddings.go, asset.go, certificate.go`, plus
  `sqlite_test.go`, `certificate_mint_test.go`, `entitlement_parity_test.go`, `merge_test.go`, `time_test.go`,
  `progress_aggregate_test.go`) ✓; `internal/web/backend/graphql/graph/schemas/academy.graphqls` ✓;
  `cmd/academy-seed`, `cmd/academyImport` ✓; app version line `v1.363.2` @ `5ba17044` ✓.
- **`next-web-app.md`** — Next `^16.2.7` in all four apps ✓, React `^19.2.7` ✓, `pnpm@10.30.3` ✓,
  `engines.node ">=24.0.0"` ✓, turbo `^2.9.6` ✓, TypeScript `^5.9.3` ✓; `apps/web/src/proxy.ts` exists and
  `middleware.ts` does **not** ✓; repo `CLAUDE.md:15` = "Next.js 16 App Router" and `:55` = the proxy.ts rename
  note — both exact ✓; **no** `.storybook/` and no storybook script, only `configs/tailwind/storybooks.css` ✓;
  exactly one `Dockerfile.dev` at repo root ✓; **8** locale dirs (de en es fr it ja nl pt) ✓;
  `docker-compose.yml:352` = the `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` build arg ✓; AI-readiness gating —
  `AI_READINESS_FLAG = 'flag_ai_readiness'` (`components/ai-readiness/aiReadiness.constants.ts:26`) and
  `aiReadinessEnabled` (`ai_readiness.graphqls:73`) are genuinely two different gates ✓.
- **`frontend_architecture.md`** — `repos.yml` holds **exactly 9** entries and ant-academy is not one ✓;
  `docker-compose.yml:311` = studio-desk ✓; `docker-compose.yml:362` = `NEXT_PUBLIC_BACKEND_API_URL` ✓;
  `pnpm-workspace.yaml` carries `'!apps/mobile'` ✓.
- **`ai_architecture.md`** — `app/internal/ai/ai.go` genuinely **does not exist**; the EU-first wrappers are
  `internal/jobsimulation/ai/ai.go` (`flag_use_azure_us` at `:267` and `:344`) and `internal/skillerai/ai.go`
  (`:345-347`) ✓; `flag_use_realtime_openai` at `internal/jobsimulation/calls/livekit.go:133` ✓;
  `job_role_embeddings` + `skill_embeddings` Ent schemas present ✓.
- **`clerk-integration.md`** — exactly **12** webhook event types (`internal/clerk/events/events.go`) ✓;
  `clerk-sdk-go/v2 v2.7.0` (`app/go.mod:31`) + `svix-webhooks v1.99.1` ✓; `metabase/route.ts:35` checks
  `authData.orgRole !== 'admin'` (bare form) ✓; `STUDIO_ACCESS_ROLES = ['admin','org:admin','content_creator',
  'org:content_creator']` (`studio-desk/src/index.ts:96`) ✓; sentinel has no Clerk import ✓.
- **`skillpath.md`** (the non-mirror parts) — `internal/skillpath/session.go:204-206` "cms-in-app deseam: cms is
  in-process" → `GetSkillPathDomain` ✓; `internal/skillpaths/skillpaths.go:88-95` same ✓; `repos.yml` 9 repos,
  0 skillpath ✓; `public.skill_path_sessions` + `skillpath_mixins.go` ✓; the per-member drill-down really does
  render `t('enterprise.insights.comingSoon')` with `userData` hardcoded `null` and the results table commented
  out (`apps/web/src/components/containers/InsightsBySkillPathStudentSimulationsContainer.tsx:31-33,138-160`) ✓.
- **No local-router drift found in this group.** Every file that touches the GraphQL endpoint —
  `next-web-app.md:14,47,96`, `academy-backend.md:13-15,75-78`, `frontend_architecture.md:32,39,44`,
  `shared_libraries.md:42` — already states the post-`2adcf71` reality. The only residue is m-E2's `:5050`
  mislabelled as a *prod* address.

---

## Counts

**3 BLOCKERS, 6 minors.**

**Character of the group: mixed, and the split is not where I expected it.** Two of the three blockers are in
the *swept* files (`security_compliance.md`), or in text the sweep authored, and the third is in a never-edited
file (`clerk-integration.md`) whose surrounding paragraph is *about* version drift — i.e. the drift was
detected on one axis and asserted-away on the other. The heaviest blocker (`skillpath.md` B-E1) is in a
never-edited file and is pure derived-fact rot: the doc's merge/status layer is perfectly correct, and what
rotted underneath is a table name, an Ent-schema path, and an operational instruction — none of which uses
merged/live/gone vocabulary. The eight remaining files are genuinely dense with exact, re-verifiable anchors
(`askengine.md`'s 49/13/10/14 test counts and 146 KB `rules.md`; `messenger.md`'s four `cmd/root.go` line
numbers; `storage.md`'s `:210`-vs-`:324` disambiguation) and I could not break them.
