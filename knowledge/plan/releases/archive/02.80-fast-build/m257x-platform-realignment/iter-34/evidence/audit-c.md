# iter-34 confirming pass — audit report, group C

Auditor: group C. Method: full top-to-bottom read of every assigned file, no grep-scoping.
Platform clone verified at `stack-demo/platform` HEAD `2adcf71` ("Merge pull request #23 …
chore/drop-wundergraph", 2026-07-31); peer clones read at `app` `5ba17044` (v1.363.2),
`graphql-wundergraph` `60c229f`. rext authoring copy read at `.agentspace/rosetta-extensions/`.

---

## 1. Positive control — every assigned file read to EOF

| # | file | `wc -l` | read to | last line I actually read (verbatim tail) |
|---|---|---|---|---|
| 1 | `corpus/services/studio-room.md` | 414 | EOF | `- [External Services](../architecture/external_services.md) - AI provider details` |
| 2 | `corpus/services/clerkenstein.md` | 357 | EOF | `- [Webhook setup](../ops/webhook_setup.md) — the real Clerk webhook path the `clerk-webhook/` injector replays into.` |
| 3 | `corpus/architecture/architecture_overview.md` | 326 | EOF | `For AI model inventory, provider routing, and voice/recording architecture, see [AI Architecture](./ai_architecture.md).` |
| 4 | `corpus/services/chronos.md` | 245 | EOF | ` ``` ` (close of the *Handling the event* Go block) |
| 5 | `corpus/architecture/platform-migration-status.md` | 189 | EOF | `- [`corpus/services/README.md`](../services/README.md) — the per-service docs this map is the index of truth for.` |
| 6 | `corpus/services/gotenberg.md` | 82 | EOF | `* [Service Taxonomy](../architecture/service_taxonomy.md)` |
| 7 | `corpus/services/customerio-sync.md` | 75 | EOF | `The "build from GitHub URL" pattern is intentional: … clone it as a sibling of `platform/` and add it to `repos.yml` temporarily.` |

**0 files UNREAD.** Total 1688 lines.

---

## 2. BLOCKERS

### B1 — `corpus/architecture/architecture_overview.md:242` · fabricated AI-provider fallback cascade

**Verbatim false text:**

> `*   **AI Providers**: EU-first routing — Azure OpenAI (EU) → AWS Bedrock (EU) → Mistral (EU) → OpenAI Direct (US fallback)`

**What is actually true at HEAD.** There is no four-tier cascade, and Mistral is not in the routing
path at all. The real selection, in `app`:

- **Azure is the default and the *only* flag-switched tier**, and it switches **EU → US**, not
  EU → Bedrock: `app/internal/jobsimulation/ai/ai.go:262-276` — `case Azure: client :=
  a.azureClientEu; … IsFeatureEnabled(nil, "flag_use_azure_us", …); if isAzureUsFlagEnabled {
  client = a.azureClientUs }`. Same construct again at `:340-352` for `AudioTranscriptions`, and
  at `internal/skillerai/ai.go:347`.
- **The only *automatic* fallback is a 429 retry onto direct OpenAI**, not onto Bedrock:
  `app/internal/jobsimulation/ai/ai.go:127-137` (`isThrottlingError` → `StatusTooManyRequests`),
  exercised by `internal/skillerai/ai_test.go:24` — *"vendor 429 detection driving the
  retry->OpenAI fallback."*
- **Bedrock is not a tier, it is a vendor.** `AnthropicAws` is a discrete `AIVendor`
  (`internal/jobsimulation/ai/ai.go:32`, dispatched at `:233`, `:280`) chosen per call-site.
  Nothing ever falls *from* Azure *to* it.
- **Mistral is not in the platform routing at all.** The only Mistral in `app` is an *indirect*
  go.mod entry (`go.mod:127`) and one SSM key whose own description scopes it:
  `app/terraform/ssm.tf:291` — *"Mistral API Key (cms studio markdown manager)"*. Grep of `app`
  outside `studio/` returns nothing else.

**Why this misdirects real work.** It fails in the dangerous direction on a compliance-sensitive
path. A reader answering *"where does our data go when Azure EU is degraded?"* is told the next
two hops are EU (Bedrock EU, Mistral EU) and that US is a last resort. In reality **one PostHog
flag (`flag_use_azure_us`) sends the very next request to Azure US**, and a 429 sends it to direct
OpenAI — both without touching an EU tier. Anyone provisioning, auditing, or removing a provider
from this doc's chain would work on infrastructure that does not exist and miss the one that does.

**Grade: BLOCKER.**

**Suggested correction (one line):** replace with
`**AI Providers**: EU-first — Azure OpenAI **EU** by default, switched to Azure **US** by the PostHog flag `flag_use_azure_us`, with direct OpenAI as the HTTP-429 retry fallback; Anthropic is always AWS Bedrock `eu-west-1` (a separate vendor, not a fallback tier). Mistral is used only by the cms Studio markdown manager. See [AI Architecture](./ai_architecture.md).`

---

## 3. Minors

### M1 — `corpus/services/clerkenstein.md:317-319` · "routed forward" work that already shipped, contradicting its own paragraph 14 lines above

> `**Routed forward to M220** (vendor the bundle; serve from disk; keep the CDN proxy only as a *bounded* fallback). Until then, treat a slow/blocked jsdelivr as a **plausible cause of an arbitrarily long demo login**.`

M220 **shipped**: `.agentspace/rosetta-extensions/clerkenstein/clerk-frontend/server.go:35-67` —
`clerkJSFetchTimeout = 15 * time.Second`, `clerkJSMaxBytes = 32 << 20`, and
`var clerkJSClient = &http.Client{Timeout: clerkJSFetchTimeout}` commented *"Explicitly NOT
http.DefaultClient"*, with `clerkjs_cache_test.go` present. The doc says so itself at `:297`
(*"✅ FIXED"*) and `:304` (*"NO LONGER a plausible cause … look elsewhere"*). The tail paragraph
then reinstates the retracted directive verbatim.
This is exactly the *work-described-as-routed-forward-that-already-shipped* rot class. Kept at
**minor** only because the block header and `:304` both retract it first; a reader who lands on
the tail (e.g. searching "jsdelivr") still gets a false debugging directive.
**Fix:** delete `Until then, treat …` and past-tense the rest ("was routed forward to M220, where
it landed").

### M2 — `corpus/services/clerkenstein.md:71` · wrong gene count for the `@clerk/express` surface

> `the **`@clerk/express`** Node-backend surface (9/9 genes, `@clerk/express` ^1.3.47, M2c …)`

`clerkenstein/alignment/dna/clerk-express-1.json` holds **5 capabilities / 13 genes** —
`ExpressAuth` (5), `ExpressRequest` (4), `ExtractIdentity` (1), `JWKS` (1), `ClerkClientBAPI` (2).
Every other count on this page verifies exactly (`clerk-2.6.0` = 14 caps / **27** genes,
`clerk-js-5` = 9, `clerk-multi-1` = 9, `clerk-deploy-1` = 7), which makes the one wrong number
harder to spot. **Fix:** `9/9` → `13/13`.

### M3 — `corpus/services/clerkenstein.md:18` · the monorepo section list is 6 of 11

> `` `rosetta-extensions` is ONE private monorepo with sections (`clerkenstein`, `demo-stack`, `stack-injection`, `stack-core`, `stack-seeding`, `alignment`) ``

Actual sections at the authoring copy: `alignment`, `clerkenstein`, `demo-stack`, `dev-stack`,
`playthroughs`, `stack-core`, `stack-injection`, `stack-secrets`, `stack-seeding`,
`stack-snapshot`, `stack-verify`. Reads as enumerative and is 5 short. **Fix:** add the five, or
mark the list "e.g.".

### M4 — `corpus/services/clerkenstein.md:100` · `cmd/` binary list omits one

> `` `cmd/` | — | standalone binaries: `mintpk` … · `fake-fapi` / `fake-bapi` … ``

`clerkenstein/cmd/` contains `fake-bapi`, `fake-fapi`, **`jwtkey`**, `mintpk`. **Fix:** add `jwtkey`.

### M5 — `corpus/services/studio-room.md:336-348` · the orchestrator is described as a separate "CMS service"

> `Orchestration is performed by the **CMS Go code**, not by studio-room itself.` … `#### With CMS Service` / `The CMS service drives the full lifecycle:` … `**Workflow**: Desk designs → CMS enqueues → Room generates → CMS imports into Directus`

There is no CMS service. The orchestrator is the **cms domain inside `app`** —
`app/internal/cms/studio/studioManager.go:119`
(`return s.runCommand(ctx, pyBin, append([]string{"studio/gen.py"}, tokens...))`) driven by
`app/internal/cms/worker/`. The `cms` container that still starts locally
(`docker-compose.yml:144`) is an unfederated husk and does **not** drive this.
Held at **minor** because the file's own ⚠ banner (`:3-11`) and its install block (`:287-290`)
both name `app/internal/cms/studio/` correctly — but this is the same shape that made this doc the
archetype of the audit: five sections of live-service prose, correct only at the top.
**Fix:** "CMS service" → "the cms domain in `app`" throughout `### Integration Points`.

### M6 — `corpus/services/studio-room.md:61` · project tree rooted at a path that does not exist

The `### Project Structure` block is rooted `studio-room/`. The doc itself corrects this at
`:287` — *"studio-room's root IS `app/studio/` … There is no studio/studio-room path."* Verified:
`app/studio/{gen.py,postgen.py,console.py,format.py,errors.py,cert.py,agents,services,configs,benchmark,knowledge}`
all present; `workspace/` is `.gitignore`d (created at runtime), and there is correctly **no**
`templates/` (confirming `:244`). **Fix:** re-root the block `app/studio/`.

### M7 — `corpus/services/studio-room.md:290` · stale line anchor for the venv constants

> `against the managed venv at `studio/studio-venv` (studioManager.go:92-94)`

The consts are at `internal/cms/studio/studioManager.go:**94-96**`
(`studioVenvDir` / `studioVenvPython` / `studioVenvPip`). The sibling anchor on the same line
(`studioManager.go:119` for the `gen.py` invocation) is **exact**. **Fix:** `92-94` → `94-96`.

### M8 — `corpus/services/chronos.md:9` · the successor is named as the standalone jobsimulation

> `have moved to **in-process [Asynq](…)** running inside jobsimulation`

The standalone jobsimulation is merged into `app` and its repo is **ARCHIVED (2026-07-31)**
(migration map `:62`). The Asynq pools that replaced Chronos now run inside `app` — `app/main.go`
around `:604-610`: *"jobsim-in-app (full merge): app OWNS jobsim — start the dual Asynq worker
pools unconditionally. The standalone jobsim service is decommissioned."* A reader chasing session
timeouts is sent to an archived repo instead of `app/internal/jobsimulation/`.
**Fix:** "inside jobsimulation" → "inside the jobsimulation domain in `app`".
(The rest of `chronos.md` is correct, including the harder call: it says the **GitHub repo is NOT
archived**, which matches the map `:72`.)

### M9 — `corpus/architecture/architecture_overview.md:16` · CMS is the only default-profile bullet with no husk annotation

Its siblings carry one — jobsimulation `:15` (*"a standalone container still starts here …
unfederated husk, until M810"*), roadrunner `:18` (*"orphaned husk"*). The CMS bullet reads as a
live content-layer service in the local `graphql` profile. It **is** in that profile
(`docker-compose.yml:144-187`, `profiles: [graphql, cms, all]`) but as a husk; the content layer
runs in `backend`. The doc gets this right in three other places (`:59`, `:148`, `:177-188`), so:
**minor**. **Fix:** append the same husk clause used at `:15`.

### M10 — `corpus/architecture/architecture_overview.md:175-190` · the "Archived / merged" table omits `graphql-wundergraph`

The table enumerates Chronos, Intelligence, Skiller, Jobsimulation, CMS, Roadrunner, Skillpath —
but not the router, which is the **newest** local removal (`b56d731` + `360efd4`, merged
`2adcf71`) and the one a reader is most likely to be looking for. The top banner `:3` covers it, so
this is completeness, not falsity. **Fix:** add a `graphql-wundergraph` row (decommissioned
locally; prod terraform still `= 1` at `graphql-wundergraph/terraform/main.tf:20`; repo ARCHIVED
2026-07-30).

### M11 — `corpus/architecture/architecture_overview.md:209` · Tier-2 table row for Studio-Room carries no "embedded in app" note

The same doc says it three times elsewhere (`:27`, `:65`, `:86`). A reader who reaches only the
table sees a standalone Tier-2 pipeline. **Fix:** mirror the `*(now a domain in `app`)*` treatment
used for CMS/Jobsimulation at `:148`/`:150`.

### M12 — `corpus/architecture/platform-migration-status.md:76` · `anthropos-studio-room` is labelled with a state whose own definition it fails

§1 `:27` defines `merged-into-app` as *"`app` owns the code and calls it unconditionally, **the
tables live in `public`**, and **the standalone is scaled to zero** — all three"*. studio-room has
no tables and was never a deployment (the row itself says *"Not a deployment, not in
`repos.yml`"*), and it is fetched at image-build time by CI, verified:
`app/.github/workflows/build-production.yml:29` `additional_repo: "anthropos-studio-room:studio"`,
with `git ls-files studio` → **0 tracked files** in `app`. The state label is a stretch by the
file's own vocabulary; guard assertion C only checks membership in the seven, so nothing catches
it. **Fix:** either widen the `merged-into-app` definition to cover "embedded in the image", or
give it a distinct label.

### M13 — `corpus/architecture/architecture_overview.md:23` · "Archived" used for Chronos

> `Archived (removed from local orchestration): Chronos, Intelligence.`

The parenthetical defines the term as *removed from local orchestration*, which is true. But the
migration map `:72` records the trap explicitly — *"**Repo is NOT archived on GitHub** … the corpus
called it archived; the org disagrees"* — and this line re-uses the exact word. **Fix:** "Removed
from local orchestration: Chronos, Intelligence (the Intelligence repo is archived on GitHub;
Chronos is not)."

---

## 4. Verified-clean — claims I checked and found TRUE (recorded so the next pass need not redo them)

**`platform-migration-status.md` — audited against the platform, not against other docs, as
instructed. Every checkable citation in it verified EXACT. Zero findings beyond M12 (a vocabulary
stretch, not a false fact).**

- Row anchors: `docker-compose.yml` `:5` sentinel · `:28` backend · `:83` jobsimulation · `:144`
  cms · `:189` storage · `:220-222` customerio-sync git-URL context · `:238` its profile · `:240`
  messenger · `:281` roadrunner · `:311` studio-desk · `:344` next-web-app · `:352` the
  `…:8082/graphql/query` endpoint · `:371-372` gotenberg — **all exact**. `repos.yml` `:10-13`,
  `:14-16`, `:20-22`, `:23-25`, `:26-28`, `:29-31`, `:34-36`, `:37-39` — **all exact**.
- Terraform desired counts, all 8 checked: app `main.tf:44 = 1` · cms `:39 = 0` · jobsimulation
  `:40 = 0` · roadrunner `:19 = 1` · sentinel `:19 = 1` · storage `:19 = 1` · messenger `:19 = 1`
  · graphql-wundergraph `:20 = 1` — **every line number and value exact**.
- Trap A: `docker-compose.yml:18` `search_path=sentinel` with `repos.yml:20-22` `migrations:
  false` — exact.
- `postgresql` / `redis` really are in the *included* file: `common.yml:2` / `common.yml:20`,
  `docker-compose.yml:1-2` `include: - common.yml` — exact.
- Every cited sha resolves, with matching date and subject: `a2a3ee6` (2026-02-27, and its diff
  really does delete the `directus:` service) · `045857c` (2026-04-17 chronos) · `fdfa189`
  (2026-04-17 intelligence) · `21429b7` (2026-07-07 skiller) · `a4db680` (2026-07-21 skillpath
  M507) · `b56d731` + `360efd4` (2026-07-31 router drop) · platform `236771f` (2026-07-29
  cms-in-app v8.0) · graphql-wundergraph `915da06` (2026-07-29 supergraph 2→1) · `8770fe6`
  (2023-05-04 nats) · `cb6ebf5` (2023-04-30 first commit) · app `5ba17044` v1.363.2.
- The **measured completeness claim reproduces exactly**: `git log -p --follow -- repos.yml` →
  **14** distinct names (app · chronos · cms · graphql-wundergraph · intelligence · jobsimulation ·
  messenger · next-web-app · roadrunner · sentinel · skiller · skillpath · storage · studio-desk),
  all 14 have rows; same command on `docker-compose.yml` → **26** names. 9 repos in `repos.yml`.
- Supergraph: `supergraph-config-prod.yaml` lists `backend` alone; `schemas/` holds
  `backend.graphqls` alone; `subgraphs.conf` = `BACKEND=v1.360.0` — exact.
- `app/internal/roadrunner/` does not exist; the runner is constructed at
  `app/internal/jobsimwiring/wiring.go:118`
  (`jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`) — exact.

**`architecture_overview.md` — the tenancy claim the brief flagged is now CORRECT and measured.**
`:274-282` says `organization_id` on *org-scoped* tables (**not** every table) and that Ent privacy
policies auto-filter *"only the 30 schemas using `OrganizationMixin{}`"*. Measured in
`app/internal/data/ent/schema/`: **139** schema files, **exactly 30** contain `OrganizationMixin{}`
(70 mention "organization" in some form — the number the sibling-file defect got wrong), and the
mixin really does carry a filtering policy: `mixin.go:126-142`
(`privacy.QueryPolicy{ rule.DenyIfNoOrganizationInContext(), rule.FilterOrganizationEdges(), … }`).
**No finding.**

Also verified true in `architecture_overview.md`: six Go services on a bare `make up`
(`Makefile:10` `PROFILE ?= graphql`; sentinel has **no** `profiles:` key so it always starts;
backend/jobsimulation/cms/storage/roadrunner all carry `graphql`) · the 23 jobsim run-state tables
(`app/terraform/migrations/20260722081626_jobsim_data_model.sql` has **exactly 23** `CREATE TABLE`)
· `extensions` schema holds pgvector (`…/20260615130000_skiller_taxonomy.sql:53` `extensions.vector(1536)`)
· messenger really does hold `CMS_RPC_ADDR=http://cms:8091` (`docker-compose.yml:256`) and reaches
backend (`:255` `BACKEND_USERS_RPC_ADDR=http://backend:8083`) · Next.js **`^16.2.7` across all four**
Next apps (web/hiring/integration/maintenance; `mobile` is Expo) and React `^19.2.7` · every one of
the ~39 relative doc links in the file resolves to an existing file.

**`studio-room.md`** — the numbers are right where it would be easiest to be wrong:
`Concurrency: 5`, `ai_video` weight **7**, `studio` weight **3**
(`app/internal/cms/worker/worker.go:29-34`) · `python:3.11-slim` (`app/Dockerfile:28`) ·
`studioManager.go:119` for the `gen.py` invocation · no `templates/` dir · requirements list
matches `app/studio/requirements.txt` including `mistralai`, with no `aiohttp` · the taxonomy client
really is the single egress (`app/studio/services/taxonomy.py:11`
`BASE_URL = "https://api.anthropos.work/api"`) · `ENVIRONMENT` really defaults to `local`
(`gen.py:31`) and `configs/local_*` really is gitignored, so `local_config.ini`'s absence from the
tree is correct, not drift.

**`clerkenstein.md`** — the version-drift ⚠ at `:266-271` is **exactly right, including the part
that is easy to get wrong**: `app/go.mod` = `colony v0.35.2` **and** `clerk-sdk-go/v2 v2.7.0`,
while `sentinel/go.mod` and `storage/go.mod` are both still `v0.34.3` (roadrunner too; cms and
jobsimulation are on `v0.35.1`). All five named `server_test.go` tests exist at `:256`, `:286`,
`:390`, `:427`, `:461`; `meorgmemberships_test.go` exists; `clerk-backend/store.go` has
`SeedOrgIdentity` (`:138`) and `LookupOrgEid` (`:151`); `alignment/cmd/alignctl/run.go:133-135`
really defines `ExitRegressed = 2` / `ExitUnmeasurable = 3`; the `clerk-2.6.0.json`
`MembershipOrgIdentity` gene carries the M219 fix note. The `alignment/` split (`:78` sibling
section vs `:101` in-repo dir) is **not** a contradiction — both exist and hold what the doc says.

**`gotenberg.md` — clean, every claim verified.** `docker-compose.yml:371-384` (image
`gotenberg/gotenberg:8`, `--api-port=3200`, `--api-timeout=60s`, `--libreoffice-restart-after=50`,
`profiles: [graphql, backend, all]`, `3200:3200`); `GOTENBERG_URL=http://gotenberg:3200` appears
**once** in the whole compose file, on `backend` (`:51`) — so *"the only consumer"* is measured, not
assumed; `app/internal/converter/gotenberg.go:13` `Timeout: 90 * time.Second`, `:16` the exact
`ConvertToPDF(ctx, gotenbergURL, document, filename)` signature, `:31`
`gotenbergURL+"/forms/libreoffice/convert"`. **0 findings.**

**`customerio-sync.md` — clean, every claim verified.** `docker-compose.yml:220-238` matches the
quoted YAML including `context: git@github.com:anthropos-work/customerio-sync.git#main`,
`ssh: ["default"]`, `8080:8080`, the `search_path=public` DSN, and `profiles: [customerio-sync,
all]`; absent from `repos.yml`, matching migration map `:74`. **0 findings.**

---

## 5. Counts

**1 BLOCKER · 13 minors.**

**How the group read: mixed, and it splits cleanly by prior treatment — but not in the direction
the brief's two warnings predicted.**

- The **swept** files (`architecture_overview.md`, `clerkenstein.md`) carry **all of the density**:
  1 blocker + 8 minors between them. Notably, the specific tenancy defect flagged in my brief has
  been repaired *correctly and with a measurement that reproduces* (30 of 139 schemas) — so the
  repair pass was not lazy. What it left behind is the class the addendum predicted: the
  **surrounding prose**. The blocker is a one-line summary bullet in a *different section* from
  anything the sweep anchored; M1 is a stale tail paragraph that the same block's own header
  retracts. Both are "read the paragraph, not the citation" misses.
- The **never-edited** files split hard. `platform-migration-status.md` is the strongest file I
  read anywhere — I audited it against the platform rather than against other docs, as instructed,
  and **every sha, every `file:line`, every terraform value and both measured completeness counts
  (14 and 26) reproduced exactly**; its single finding is a vocabulary stretch, not a false fact.
  `gotenberg.md` and `customerio-sync.md` are likewise clean end-to-end. `chronos.md` is
  well-fenced and gets the hardest call right (repo NOT archived). `studio-room.md` — the archetype
  — now opens with a correct ⚠ banner and is technically exact on every number I could check, but
  its `### Integration Points` section is still written as though a standalone CMS service existed,
  which is the same failure shape that made it the archetype, one section further down the page.
  So: the low pass-1 density on unswept files here was **partly real cleanliness, partly
  under-detection** — 5 of my 13 minors are in files nobody has ever edited.
