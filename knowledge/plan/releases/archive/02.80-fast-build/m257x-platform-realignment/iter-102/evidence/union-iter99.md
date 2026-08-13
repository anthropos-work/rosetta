# iter-99 union — the UPHELD in-scope blocker set, extracted verbatim

**Source:** `iter-99/adjudication.md` + `iter-99/verdicts/adj{1,2,3,4}.md`. Extraction only — nothing
here is repaired, re-graded, or re-derived. No corpus file was read or touched to produce it.

## Count reconciliation — **28 extracted, matching N**

| quantity | value |
|---|---|
| rows below (distinct in-scope upheld blocker units) | **28** |
| `adjudication.md` N | **28** |
| discrepancy | **none** |
| distinct predicates after grouping | **27** |

The arithmetic reproduces exactly. Per-adjudicator `IN-SCOPE-UPHELD-BLOCKERS` footers are
Adj1 = 9 · Adj2 = 9 · Adj3 = 9 · Adj4 = 7 → reading #21 `n₁ = 18`, reading #22 `n₂ = 16`, matched
`m = 6` → union **28**. The 6 matched are exactly the ones `adjudication.md:40-41` names:
`ai-readiness.md:305` · `ai-readiness.md:46` · `backend.md:33-34` · `dependency_map.md:59` ·
`hiring.md:38` · `service_taxonomy.md:405`.

**One unit-of-count caveat, stated rather than smoothed** (see `## Notes` #1): the `N = 28` unit is the
*deduplicated finding*, not the literal file:line. Several of these findings carry more than one anchor
(D B3 carries three; the terraform-address predicate carries two primaries plus a twin; four more carry a
twin each). Expanding every twin would give ~37 file:line sites, not 28. I kept the row unit at the
finding — the unit that produces 28 — and listed **every** anchor inside the anchor cell.

---

| # | predicate (short name) | the false claim (VERBATIM from the verdict) | anchor (file:line) | what is true (verbatim from the adjudicator's verdict) | adjudicator |
|---|---|---|---|---|---|
| 1 | `readme-folded-roadrunner-domain` | "the folded skiller …, skillpath, jobsimulation, cms and **roadrunner** domains" | `corpus/services/README.md:37` | "Ground truth: `git -C stack-demo/app ls-tree --name-only HEAD internal/ \| grep -i road` → **exit 1**, and identically at `origin/main` — `app/internal/roadrunner/` exists at neither ref. The fenced map says the same in the `app` row (`platform-migration-status.md:87`: *"**`app/internal/roadrunner/` does not exist**"*)." | Adj1 · r21 B B2 |
| 2 | `coursebuilder-sse-cost-event` | "`cost` documented on the SSE wire; the code filters it off by design" | `corpus/services/coursebuilder.md:77-79` | "`internal/web/backend/coursebuilder/handler.go:2709-2717` — `case cb.EventCost:` returns `\"\", nil` under a comment reading *\"the terminal cost readout is COGS and MUST NOT reach the customer SSE stream … filtered here at the wire boundary by returning an empty event name\"* … the live names are `text`, `score`, `patch_applied`, `patch_skipped`, `stage`, `outline`, `progress`, `preview_ready`, `draft_kept`, `error`, `translation_ready`, `rebuild_required`, `steering_received`, `steering_applied`, plus `session` (`:1458`) and `done` — **16**; the doc lists 13 and omits four" | Adj1 · r21 B B3 |
| 3 | `storage-env-compose-value` | "Table records `ENVIRONMENT` as never set by compose; the block set it" | `corpus/services/storage.md:215` | "the column's contract is explicit — header `:208` *\"Compose value\"*, banner `:203-204` *\"The middle column records what `docker-compose.yml` set on the `storage` service block.\"* The block set it at **both** refs the surrounding prose names: `platform@0dab54d docker-compose.yml:119` `- ENVIRONMENT=development`, and `platform@2adcf71 docker-compose.yml:206` `- ENVIRONMENT=development`." | Adj1 · r21 B B4 |
| 4 | `app-rpc-mux-universal` | "All of their Connect-RPC surfaces are on `app`'s mux" | `corpus/services/backend.md:29` | "enumerated the mux from `stack-demo/app` `main.go:1185-1228` @ `b948604f` — **six** handlers … Of the **eight** services the banner's own table (`:9-16`) quantifies over, only three are represented. The falsity is **non-vacuous** … `storage` declares a real `StorageService` Connect surface … and `messenger` a `MessengerService` … neither appears on `app`'s mux at either ref." | Adj1 · r21 C B1 |
| 5 | `jobsim-origin-main-mislabel` | "Labels `7177374` as `origin/main`; at the real origin/main the anchor names another construct" | `corpus/services/jobsimulation.md:203-204` | "`git rev-parse origin/main` → **`2035f9a4`**; `git rev-list --count 7177374..origin/main` → **33** … `git show origin/main:main.go \| sed -n 216p` → `if strings.Contains(dsn, \"://\") {`; `func main()` is at **`:229`** there. So a reader who resolves the label as written lands on an unrelated construct." | Adj1 · r21 D B2 |
| 6 | `cms-has-not-moved` | "**Do not generalise M810 from this row:** `cms` **has not moved**" (`:89`) · "while `cms` **sits untouched** at `service_desired_count = 0`" (`:270`) | `corpus/architecture/platform-migration-status.md:89` · `:270` · `corpus/services/jobsimulation.md:12` | "`:88` reads *\"**M257x iter-92 — cms has since taken an M810 step, and it points the OTHER way** … this repo now holds **two measured facts pointing opposite ways** … **report both, assert neither**.\"* … Both halves re-derived in the `cms` clone: `terraform/main.tf:39` is `service_desired_count = 0`, **and** `6efa1d5` … deletes `.github/workflows/build-production.yml`" | Adj1 · r21 D B3 |
| 7 | `unpoliced-remainder-is-reference-data` | "the unpoliced remainder is *not* all \"global reference data\"; ≥4 members are tenant-scoped" | `corpus/architecture/security_compliance.md:104-105` | "`lab.go:63` `field.String(\"tenant_eid\")`, `academy_chapter.go:84`, `academy_skill_path.go:78` (all `tenant_eid`), `skill_path_session.go:43` `field.UUID(\"tenant_id\")`. I grepped each of those four files for `organization_id`, `OrganizationMixin`, `OrganizationIDMixin` and `func … Policy()`: **zero hits in all four** … the remainder is 135 − 31 policed − 23 org_id-unpoliced = **81** schemas, of which at least these 4 are per-tenant" | Adj2 · r21 E B1 |
| 8 | `academy-subgraph-exists` | "there is no \"academy subgraph\"; the corpus elsewhere says so explicitly" | `corpus/architecture/service_taxonomy.md:37` (+ `:266`) | "`stack-demo/graphql-wundergraph` @ `60c229f3` — `supergraph-config-prod.yaml` declares exactly one subgraph (`- name: backend`), `schemas/` holds one file (`backend.graphqls`) … No `academy` subgraph, routing_url or SDL entry exists. Same-file contradiction: `service_taxonomy.md:387` *\"**`backend` alone (1)**\"*" | Adj2 · r21 E B4 |
| 9 | `messenger-not-opt-in` | "That one is not opt-in: it is the stock `core` selection." | `corpus/architecture/dependency_map.md:58` (vs `:21`) | "at `app` `origin/main` `2035f9a4`, `main.go:1445` is `if messengerEnabled {` wrapping the whole subscriber-server construction … and `env_guards.go:61` `envMessengerEnabled = \"MESSENGER_ENABLED\"`; `stack-demo/platform` @ `0c91421` `docker-compose.yml:84-92` deliberately does not set it" | Adj2 · r21 F B2 |
| 10 | `five-shared-library-repos` | "that shared plumbing lives in **five small repos that the services pull in like any third-party dependency**" | `corpus/architecture/shared_libraries.md:11` (and `:3`) | "colony **7/7**, proto **7/7**, taxonomy **6/7**, ai **3/7**, authn **0/7** ⇒ the set of repos any service pulls has cardinality **4** … Same file, `:167`: *\"**No checked-out service imports the standalone `github.com/anthropos-work/authn`**\"*" | Adj2 · r21 F B3 |
| 11 | `recruiter-scoreboard-in-apps-web` | "the recruiter scoreboard is not reachable in `apps/web` for a genuine hiring org" | `corpus/services/next-web-app.md:32` | "`apps/web/src/context/UserStatusContext.tsx:141-173` — `userHasAllHiringOrgs` is computed from `membership.organization.publicMetadata.isHiring` (`:144-149`), and when true the effect sets `window.location.href = buildSwitchHandoffUrl({ targetProduct: 'hiring', … })` (`:168-172`), i.e. the user is ejected out of `apps/web`." | Adj2 · r21 G B1 |
| 12 | `gwg-self-cite-84` | "the self-cited `:84` describes ports, not rebuild-on-SDL-change" | `corpus/services/graphql-wundergraph.md:134` | "`graphql-wundergraph.md:84` is the **Ports** bullet — *\"**Ports**: **8080 → 8080** (router `listen_addr 0.0.0.0:8080`, `graphql_path /graphql`). **There is no `5050` at platform HEAD**…\"*, running `:84-88`, entirely about ports … The construct the sentence needs is at `:114-117`." | Adj2 · r21 G B4 |
| 13 | `ai-readiness-urls-ts-52` | "`urls.ts:52` names `ORGANIZATION_FEEDBACK_URL`, not `AI_READINESS_URL`" | `corpus/services/ai-readiness.md:305` (r22: `:304-306`) | "`packages/core-js/src/constants/urls.ts:50` `export const AI_READINESS_URL = '/ai-readiness';`; `:52` is `export const ORGANIZATION_FEEDBACK_URL = '/enterprise/organization-feedback';` — a different constant for a different route … at next-web-app `origin/main` `8297c684` the constant is at `:51`. No ref makes `:52` correct." | **matched** — Adj1 · r21 B B5 + Adj3 · r22 B B1 |
| 14 | `ai-readiness-self-cite-458` | "contradicted `:458` of this same file, which is already in the past tense" | `corpus/services/ai-readiness.md:46` (r22: `:46-47`) | "`:456-457` close the *\"Route 2 …\"* paragraph, **`:458` is empty**, `:459` opens the `> **✅ CORRECTED M219 …**` blockquote. The sentence asserts a property of that line … and the line has no content to carry it." · Adj3: "The statement the sentence points at … is at **`:490`**. The sentence's entire evidentiary weight is a cross-reference it does not deliver." | **matched** — Adj1 · r21 B B6 + Adj3 · r22 B B2 |
| 15 | `app-stream-set-omits-backend` | "`skiller` is **NOT a fifth member**" | `corpus/services/backend.md:33-34` | "Both-ways set = **five**: backend, skillpath, jobsimulation, cms, ai_usage; `SKILLER_STREAM` consumer-only (`:1276`). `:33-34` names four, drops `backend`, and then closes the set … making the omission an explicit false exhaustiveness claim." · Adj3: "*\"`skiller` is **NOT a fifth member**\"* converts an illustrative list into a completeness claim, and gets the ordinal wrong — skiller would be a sixth." | **matched** — Adj1 · r21 C B2 + Adj3 · r22 C B2 |
| 16 | `skiller-stream-file-count` | "`SKILLER_STREAM` at `2035f9a4` spans 3 Go files, not 4" (booked claim: "`SKILLER_STREAM` has 6 Go occurrences across **4** files.") | `corpus/architecture/dependency_map.md:59` | "`git -C stack-demo/app grep -n SKILLER_STREAM 2035f9a4 -- '*.go'` → 6 lines … `git grep -l` at the same ref → **3** files. The occurrence count (6) is right; the file cardinality is 3." · Adj4: "Off the `*.go` pathspec the set is **6** files … dropping only the markdown gives 4 files but **7** occurrences. No reading of the corpus yields \"6 across 4\"." | **matched** — Adj2 · r21 F B1 + Adj4 · r22 F B1 |
| 17 | `hiring-twin-service-taxonomy-52` | "the cited twin `service_taxonomy.md:52` is about a different subject entirely" | `corpus/services/hiring.md:38` | "`corpus/architecture/service_taxonomy.md:52` reads *\"[`dependency_map.md`](./dependency_map.md)'s content-generation flow, which had it right all along.\"* — the closing line of a blockquote (`:44-52`) … It says nothing about schemas, `jobsimulation`, or local stacks. The intended twin is `service_taxonomy.md:62`" | **matched** — Adj2 · r21 G B3 + Adj4 · r22 G B3 |
| 18 | `single-service-address-gotenberg` | "Compose sets a single service address" | `corpus/architecture/service_taxonomy.md:405` | "`docker-compose.yml:48` `- AUTHORIZATION_ADDRESS=http://sentinel:8087` **and** `:57` `- GOTENBERG_URL=http://gotenberg:3200`, both in `backend`'s `environment:` block. So *\"Compose sets a single service address\"* is false at the named ref … Under the bullet's own heading (*Core Services ↔ Core Services*) there are therefore two cross-process edges, not \"exactly one\"" | **matched** — Adj2 · r21 E B2 + Adj4 · r22 E B1 |
| 19 | `single-service-address-gotenberg` | "*\"the only service address compose sets\"* / *\"the one cross-process edge a local stack has\"* is refuted nine lines after the line it cites — `GOTENBERG_URL`" | `corpus/services/sentinel.md:85` (twins `corpus/services/jobsimulation.md:145-146` · `corpus/architecture/platform-migration-status.md:93`) | "inside the **same** `backend` `environment:` block that carries `:48 AUTHORIZATION_ADDRESS=http://sentinel:8087`, line **`:57`** reads `- GOTENBERG_URL=http://gotenberg:3200` … That is a genuine cross-process HTTP hop to a second container on the default profile … the over-generalisation from \"no RPC edge\" to \"no cross-process edge\" is the defect. `architecture_overview.md:321` states the correctly-scoped version" | Adj3 · r22 D B2 |
| 20 | `messenger-unpinned-anchors` | "two unpinned present-tense citations resolve at no ref the doc names; one names a file absent from the graded tree" | `corpus/services/messenger.md:53` (twin `:149`) | "`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → **empty**. The file **does not exist** at the graded checkout … The **only other ref this file names** — `app` `9d00a313` v1.367.0 … fails too … Both resolve **only** at `app` origin/main `2035f9a4`" | Adj3 · r22 A B2 |
| 21 | `prod-terraform-backend-internal-8081` | "Production terraform still names `http://backend.internal.anthropos:8081`" | `corpus/services/cms.md:196` (twin `:55`) · `corpus/services/jobsimulation.md:49-50` | "`git grep` at each clone's own ref, across all 12 clones, restricted to `-- '*.tf'` → **0 files** … Working-tree recursive `grep -rn` … → **0** … **Self-contradiction inside the same file:** `cms.md:18` (*\"the deletion itself lands in `infrastructure`, **which has never been in any clone set we have**\"*) … A doc that states it cannot see the production terraform cannot report what it \"still names\"." | Adj3 · r22 B B3 + D B1 (one predicate, two anchors — enters `N` once) |
| 22 | `customerio-external-services-section` | "points at an `external_services.md` section that does not exist, and re-publishes the fossil the same file retracts" | `corpus/services/customerio-sync.md:140` | "`corpus/architecture/external_services.md` `##` headings, enumerated from source … **No Customer.io section, no Brevo section.** … `grep -ic brevo` → **0**. So *\"Customer.io as an integrated SaaS\"* is false in both directions. It also contradicts `customerio-sync.md:18-21` 122 lines above — *\"**The name is a fossil.** The destination has been **Brevo**, not Customer.io\"*" | Adj3 · r22 C B1 |
| 23 | `askengine-bedrock-disabled` | "a failed Bedrock init returns from `main()` — the whole backend exits; nothing is \"disabled\"" | `corpus/services/askengine.md:88-89` | "`411  if err != nil {` / `412      logger.Error(\"bedrock client unavailable; talk-to-data disabled\", \"error\", err)` / `413      return` … the `return` is at two tabs inside a one-tab `if`, i.e. the top level of `func main()`. So the process ends before the RPC mux, the GraphQL server, the Echo router, the Asynq pools and the Redis subscribers are started. The corpus took the platform's own misleading log string … as a description of behaviour" | Adj3 · r22 D B3 |
| 24 | `compose-service-census-26` | "the census says **26** service names ever in `docker-compose.yml`; the set is **25**, and the 26th token is the `app-network` **network**" | `corpus/architecture/platform-migration-status.md:78-81` | "Over all **80** commits `git log --follow` reports for `docker-compose.yml` … collecting only 2-space keys under `services:` … = **25**. The naive section-blind derivation … returns **26**; `comm -3` between the two sets yields exactly one token: **`app-network`** … the passage's own audit instruction — *\"a name they return that has no row is a gap\"* — turns the error into a false alarm for anyone who re-runs it." | Adj3 · r22 D B4 |
| 25 | `two-brevo-pushers-up-all` | "`make up-all` never ran two Brevo pushers; backend's own was never on locally" | `corpus/architecture/service_taxonomy.md:101-102` | "`env_guards.go:92-111` — `resolveSubsystemSwitch` returns `(false, nil)` for an empty value when `deployed == false`; `deployedEnvironment()` (`:37-44`) returns **false** for `ENVIRONMENT=development` … Platform `0dab54d:docker-compose.yml:56` sets `ENVIRONMENT=development` on `backend`, so the in-process pusher was OFF at every ref between the fold and the container's deletion at platform `838d907` … it self-contradicts three lines up in the same blockquote (`:98-100`)" | Adj4 · r22 E B2 |
| 26 | `academy-progress-write-path` | "Progress writes are posted from the client harness, not the beacon fallback route" | `corpus/services/ant-academy.md:63` | "`code/app/api/academy/beacon/route.js:1-18`, states the mechanism in the opposite order: *\"POST /api/academy/beacon — the on-unload last-ditch write flush … a best-effort last-ditch flush\"* … the real write path … `code/src/progress/store.js:26-27` … fires them at `:162` and `:210` … i.e. straight to the supergraph … Every in-session write is posted from there; the beacon route is the exception the sentence presents as the rule." | Adj4 · r22 F B2 |
| 27 | `hiring-manager-go-anchors` | "`manager.go:485` is `}` and `:448` a blank line; the constructs are at `:535-537` / `:450-453`" | `corpus/services/hiring.md:80-81` | "`stack-demo/app` @ `b948604f`, `internal/organization/manager.go` — `:448` is blank (`:450` is `switch org.IsHiring`, `:453` `antRole = enum.RoleCandidate`); `:485` is the closing `}` of `forceUserToOrganizationInClerk`. The hard-error is 50 lines further down, in a different function: `:535` `if !org.IsHiring {`" | Adj4 · r22 G B1 (**disputed** — REJECTED by Adj2 as r21 G B2 on ref-discipline) |
| 28 | `workforce-intelligence-hiring-gate` | "Nothing gates Workforce Intelligence on `isHiringOrg`; the gate is `isAdmin` + an `apps/hiring` prop" | `corpus/services/hiring.md:302-303` | "I enumerated **every** `isHiringOrg` site in `packages/ui/src/NavBar/useNavbarSections.tsx` … **None** touches the workforce entry … `orgSectionVisibility({ isAdmin, showStudio })` returns `intelligence: isAdmin` (`packages/ui/src/NavBar/orgGroups.ts:48-64`) — **no `isHiringOrg` parameter at all**. `showWorkforce` defaults to `true` … and is passed `false` in exactly two places, both in `apps/hiring`" | Adj4 · r22 G B2 |

---

## Predicate roll-up

27 distinct predicates over 28 rows. Only one predicate carries more than one row.

| predicate | rows | anchors |
|---|---|---|
| `single-service-address-gotenberg` | **2** (#18, #19) | `service_taxonomy.md:405` · `sentinel.md:85` · `jobsimulation.md:145-146` · `platform-migration-status.md:93` |
| `readme-folded-roadrunner-domain` | 1 | `services/README.md:37` |
| `coursebuilder-sse-cost-event` | 1 | `coursebuilder.md:77-79` |
| `storage-env-compose-value` | 1 | `storage.md:215` |
| `app-rpc-mux-universal` | 1 | `backend.md:29` |
| `jobsim-origin-main-mislabel` | 1 | `jobsimulation.md:203-204` |
| `cms-has-not-moved` | 1 | `platform-migration-status.md:89` · `:270` · `jobsimulation.md:12` |
| `unpoliced-remainder-is-reference-data` | 1 | `security_compliance.md:104-105` |
| `academy-subgraph-exists` | 1 | `service_taxonomy.md:37` · `:266` |
| `messenger-not-opt-in` | 1 | `dependency_map.md:58` (vs `:21`) |
| `five-shared-library-repos` | 1 | `shared_libraries.md:11` · `:3` |
| `recruiter-scoreboard-in-apps-web` | 1 | `next-web-app.md:32` |
| `gwg-self-cite-84` | 1 | `graphql-wundergraph.md:134` |
| `ai-readiness-urls-ts-52` | 1 | `ai-readiness.md:305` (r22 `:304-306`) |
| `ai-readiness-self-cite-458` | 1 | `ai-readiness.md:46` (r22 `:46-47`) |
| `app-stream-set-omits-backend` | 1 | `backend.md:33-34` |
| `skiller-stream-file-count` | 1 | `dependency_map.md:59` |
| `hiring-twin-service-taxonomy-52` | 1 | `hiring.md:38` |
| `messenger-unpinned-anchors` | 1 | `messenger.md:53` · `:149` |
| `prod-terraform-backend-internal-8081` | 1 | `cms.md:196` · `:55` · `jobsimulation.md:49-50` |
| `customerio-external-services-section` | 1 | `customerio-sync.md:140` |
| `askengine-bedrock-disabled` | 1 | `askengine.md:88-89` |
| `compose-service-census-26` | 1 | `platform-migration-status.md:78-81` |
| `two-brevo-pushers-up-all` | 1 | `service_taxonomy.md:101-102` |
| `academy-progress-write-path` | 1 | `ant-academy.md:63` |
| `hiring-manager-go-anchors` | 1 | `hiring.md:80-81` |
| `workforce-intelligence-hiring-gate` | 1 | `hiring.md:302-303` |

**Per-file concentration** (the repair-routing view): `hiring.md` 3 · `backend.md` 2 ·
`service_taxonomy.md` 3 · `platform-migration-status.md` 3 (two of them shared with other rows) ·
`ai-readiness.md` 2 · `dependency_map.md` 2 · `jobsimulation.md` 3 (two as twins).

---

## Notes

1. **Unit of count — "anchor" in `adjudication.md` means "deduplicated finding".** The task brief asks for
   one row per anchor and names N = 28. `adjudication.md`'s own N = 28 is derived as
   `n₁ (18) + n₂ (16) − m (6)`, where `n₁`/`n₂` are the adjudicators' `IN-SCOPE-UPHELD-BLOCKERS` footers —
   which are **finding** counts after each adjudicator's own dedup (Adj1 counts D B3's three anchors once;
   Adj3 explicitly collapses B B3 + D B1 to one and writes *"10 upheld bookings collapse to **9** distinct
   in-scope predicates"*). Expanding every named anchor and twin instead yields ~37 file:line sites. I kept
   the row unit at the finding so the extraction reconciles with N exactly, and put **every** anchor in the
   anchor cell so no repair target is lost. `adjudication.md:68`'s own prediction-#8 line
   (*"28 anchors / ~24 predicates ≈ 1.2"*) uses "anchor" in this same loose sense; my grouping lands on
   **27** predicates rather than ~24, and I did not force it downward. The 3-predicate gap is entirely a
   grouping-judgement difference, not a missing or extra finding — the 28 units are identical.

2. **Excluded — UPHELD but OUT-OF-SCOPE (1).** Adj1 r21 D B1, `corpus/services/jobsimulation.md:136` (the
   corpus still mandating seeding two DROPPED mirror tables). Adj1 upheld it as a real finding and then
   enumerated every corpus occurrence: *"Inside the audited partition the corpus is **uniformly on the
   correct side** … Every mandating site is in `corpus/ops/**` or `CLAUDE.md`."* It does not enter `N` and
   is not a row here — **but it is a real, live defect** whose repair surface is `corpus/ops/**` +
   `CLAUDE.md`, outside clause 5's scope. Worth carrying forward separately.

3. **Excluded — the 10 REJECTED.** Adj1: `chronos.md:27`, `external_services.md:208-211`,
   `services/README.md:21`, `backend.md:236`. Adj2: `security_compliance.md:185`, `hiring.md:80-82`.
   Adj3: `external_services.md:208-209`, `backend.md:50-52` (and `:12`), `backend.md:18-19`.
   Adj4: `security_compliance.md:158` + `:222-224`.

4. **Row 27 is the one adjudicator disagreement in the reading.** `hiring.md:80-82` was **REJECTED by Adj2**
   (r21 G B2, class *ref-discipline*: exact at `5ba17044`, the ref the doc's own re-grounding banner at
   `hiring.md:17` names) and **UPHELD by Adj4** (r22 G B1, which considered and rejected that shelter under
   §5 rule 33 — *"a pin's scope is the claim's own block"*, and the banner is 60+ lines away in a different
   section). `adjudication.md:100-101` records this as the first non-zero adjudicator variance in five
   readings. It enters `N` because reading #22's booking was upheld; a repair here should expect the
   ref-discipline counter-argument.

5. **Rows 18 + 19 are the strongest merge candidate and I made it.** Adj3's own dedup note declines to
   collapse its C B2 and D B2 (different predicates — correct), but says nothing about whether its D B2
   (`sentinel.md:85`) is the same predicate as Adj2's E B2 / Adj4's E B1 (`service_taxonomy.md:405`) —
   they were graded by different adjudicators on different seats. Both are *"compose sets a single service
   address / that is the only cross-process edge"*, both refuted by the identical evidence
   (`docker-compose.yml:57 GOTENBERG_URL`), so they share one repair. They are kept as **two rows** because
   `adjudication.md` counts them as two union members (`service_taxonomy.md:405` is one of the 6 matched;
   `sentinel.md:85` is reading-#22-only) — collapsing the rows would break the reconciliation with N.

6. **Where a verdict gave no direct quote of the corpus text**, the "false claim" cell carries the
   adjudicator's own verbatim headline or `**Predicate:**` line instead — per the brief's wording
   (*a verbatim quote from the verdict*). This applies to rows 2, 3, 5, 7, 8, 11, 12, 16, 20, 22, 23, 24,
   25, 26, 27, 28. Rows 1, 4, 6, 9, 10, 13, 14, 15, 17, 18, 19, 21 carry the adjudicator's direct quote of
   the corpus sentence itself.

7. **Two rows were INDUCED by iter-98's own repair** (`adjudication.md:106-113`): row 16
   (`dependency_map.md:59` — iter-98 pinned the cell and wrote the wrong file count in the same sentence)
   and row 15 (`backend.md:33-34` — iter-98's rewrite removed `skiller` and left an exhaustive-set claim
   that omits `backend`). Both sit inside prose iter-98 rewrote. Repairing them without re-reading the
   surrounding rewritten paragraph would risk a third induction.

8. **Verbatim fidelity.** Ellipses inside quoted evidence (` … `) are the adjudicators' own or mark where I
   elided mid-quote to fit a table cell; no wording inside a quote was altered, reordered, or corrected.
   One character was escaped for table safety: the literal pipe in row 1's shell pipeline
   (`ls-tree … | grep -i road`) is written `\|`.
