# Adjudicator 3 — verdicts for seats A · B · C · D (reading #22)

Every booking below was re-derived from the platform clones at the ground-truth refs, opening the file
myself and reading around the cited line. No seat's evidence and no prior verdict was used as a premise.

Refs re-verified at this adjudication's open (all match the ground-truth sheet):
`platform 0c91421d` · `app b948604f` (origin/main `2035f9a4`; `9d00a313` present) · `next-web-app bb3313bc`
(origin/main `8297c684`) · `sentinel 88bc5592` · `storage 4ce8ece5` · `messenger fa47850d` ·
`cms ca50c817` · `graphql-wundergraph 60c229f3` · `roadrunner 87d8d443` · `jobsimulation 462343b0` ·
`studio-desk 14a5442a` · `ant-academy 9c3843cd` · `app/studio` + `cms/studio` `aeec036a` ·
`.agentspace/rosetta-extensions 5fb0915e` (authoring, `main`) · `stack-demo/rosetta-extensions ab81527a`
(per-stack consumption clone, tag `fast-build-m257x-iter-58`).

rosetta tree: HEAD `964b7a3` — `git diff e858fd45 HEAD -- corpus/` is **empty**, so every corpus line
number below is the one the seats read.

---

## Seat A

### A B1 | `corpus/architecture/external_services.md:208-209` | **REJECTED** | IN-SCOPE
**Predicate booked:** the two rext anchors (`gen_injected_override.py:669-670` and `:84`) name the wrong
constructs.

**evidence (re-derived):** the passage names a path, `rosetta-extensions/stack-injection/…`, and **no ref
and no clone role** — and there are **two** clone roles for that repo, which the corpus's own CLAUDE.md
defines (authoring copy `.agentspace/…`; per-stack consumption copy `stack-<role>/rosetta-extensions @
<tag>`). Measured in **both**:

| anchor | authoring `5fb0915e` | consumption `ab81527a` (tag `fast-build-m257x-iter-58`) |
|---|---|---|
| `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` | `:86` | **`:84` — exact** |
| `if with_directus and name in DIRECTUS_DATA_CONSUMERS:` / `env.append(f"DIRECTUS_BASE_ADDR=…")` | `:698-699` | **`:669-670` — exact** |
| `def test_backend_the_actual_reader_is_repointed` | `:1109` — exact | `:1108` |

Both anchors the seat booked resolve **byte-exact** in the pinned per-stack clone that sits inside the very
workspace this audit reads (`stack-demo/rosetta-extensions`). The seat measured only the authoring copy on
a moving `main`, which has drifted +2 lines, and treated one tree as the tree.

**class: mis-read** — wrong tree. This is precisely the failure `platform-alignment.md` §5 rule 44 names
(*name the tree AND its ref, PER TREE*), applied here to a repo with two sanctioned checkouts rather than
to a nested one. The seat's own "what I could not settle" section concedes it did not rule this out.

### A B2 | `corpus/services/messenger.md:53` (twin `:149`) | **UPHELD** | IN-SCOPE
**Predicate:** two unpinned present-tense citations resolve at no ref the doc names; one names a file
absent from the graded tree.

**evidence (re-derived):**
- `git -C stack-demo/app ls-tree b948604f -- env_guards.go` → **empty**. The file **does not exist** at the
  graded checkout. `git -C stack-demo/app ls-tree b948604f -- internal/messenger` → **empty** too, and
  `git grep -n MESSENGER_ENABLED b948604f -- '*.go'` → **rc 1, 0 hits** (positive control
  `DIRECTUS_BASE_ADDR` at the same ref returns `cms_reader_switch.go:1` and others). The v9.0 messenger
  fold is **93 commits ahead** of the pinned checkout (`git rev-list --count b948604f..origin/main` = 93).
- `b948604f:main.go:295` = `if err != nil {` — the error branch of
  `openai.NewAzure(AZURE_OPENAI_KEY, …)` at `:294`. Not a `log.Fatalf`, not messenger.
- The **only other ref this file names** — `app` `9d00a313` v1.367.0, at `:37-38` — fails too:
  `9d00a313:env_guards.go:61` is prose (`// GetBucketLocation rather than HeadBucket, deliberately.`) and
  `9d00a313:main.go:295` is `}`.
- Both resolve **only** at `app` origin/main `2035f9a4`: `env_guards.go:61` =
  `envMessengerEnabled = "MESSENGER_ENABLED"`; `main.go:295-299` = the
  `if (messengerEnabled || customerIOSyncEnabled) && os.Getenv("BREVO_KEY") == "" { log.Fatalf(…) }`.

Not the ref-discipline class: the blocks carry **no pin**, are **present tense**, and the contradicting
evidence is **older** than the claim, not newer — the inverse of rule 2. The convention exists and is
applied four lines away (`:43` writes "@ `app` origin/main" explicitly, and even flags that its anchors
moved), so this is a skipped pin, not a missing one. A reader on the stack's own `app` clone opens
`env_guards.go` and finds no such file.

---

## Seat B

### B B1 | `corpus/services/ai-readiness.md:304-306` | **UPHELD** | IN-SCOPE
**Predicate:** `AI_READINESS_URL` cited at `urls.ts:52`; it is at `:50`, and `:52` is a different live
constant.

**evidence (re-derived):** `next-web-app` @ `bb3313bc`,
`packages/core-js/src/constants/urls.ts:49-53` =
`WORKFORCE_URL` / **`:50 AI_READINESS_URL = '/ai-readiness'`** / `TALK_TO_DATA_URL` /
**`:52 ORGANIZATION_FEEDBACK_URL = '/enterprise/organization-feedback'`** / `INSIGHTS_URL`.
At `origin/main` (`8297c684`) it is `:51`. **`:52` at no ref in the clone.** The other three anchors in the
same bullet are exact (`useNavbarSections.tsx:4` import, `:398-400` `aiReadinessMenuItem` with
`key: AI_READINESS_URL` at `:400`, `:547` the gate) — so this is a single wrong-construct anchor, not file
drift. The passage at `:299-302` explicitly promises these anchors *"still resolve at HEAD … so they remain
checkable evidence"*; the anchor lands on a real, wrong answer of the same family, which is the failure
mode a reader cannot detect.

### B B2 | `corpus/services/ai-readiness.md:46-47` | **UPHELD** | IN-SCOPE
**Predicate:** the intra-corpus self-citation to `:458` resolves to a blank line.

**evidence (re-derived):** `ai-readiness.md:458` is **empty** (`:457` = *"…The current manager dashboard
passes the cycle id anyway."*, `:459` = the `> **✅ CORRECTED M219 …**` blockquote opener, which is about the
*cycle-param* demo-patch, a different subject). A case-insensitive grep for `loadmembers` over the file
returns `:19, 37, 38, 40, 41, 42, 44, 484, 490, 491, 493, 495` — **nothing at or adjacent to 458**. The
statement the sentence points at ("*In the demo*, M51 iter-09 bounds it with the
`app-aireadiness-snapshot-loadmembers` app read-path demo-patch … 180 s → 19 ms") is at **`:490`**. The
sentence's entire evidentiary weight is a cross-reference it does not deliver.

### B B3 | `corpus/services/cms.md:196` (twin `:55`) | **UPHELD** | IN-SCOPE
**Predicate:** *"Production terraform still names `http://backend.internal.anthropos:8081`"* — no terraform
in any clone names it, and the same file declares that terraform unreadable.
**(Same predicate as D B1, different anchor — see the dedup note.)**

**evidence (re-derived), three instruments:**
1. `git grep` at each clone's own ref, across all 12 clones, restricted to `-- '*.tf'` → **0 files**
   (tracked `.tf` counts: app 5 · sentinel 4 · storage 5 · messenger 4 · cms 4 · gwg 4 · roadrunner 4 ·
   jobsimulation 6 · next-web 2 · studio-desk 4 — the pipeline is live, not empty).
2. Working-tree recursive `grep -rn --binary-files=text --include='*.tf'` over all of `stack-demo/`
   (sees untracked files and the nested `app/studio` + `cms/studio` checkouts) → **0**.
3. Unrestricted working-tree grep for `backend.internal.anthropos` → the literal exists in exactly two
   non-fixture places: `app/knowledge/service-dependencies.md` (a **markdown** file) and, at port **8080**
   not 8081, `graphql-wundergraph/supergraph-config-prod.yaml:6` + `CLAUDE.md:18`.
4. The only prod terraform touching these values is messenger's, and it names **variables, not literals**:
   `messenger/terraform/main.tf:74/78/82/86` → `"${var.cms_rpc_address}"` etc., and
   `terraform/variables.tf:77-95` declares `cms_rpc_address` / `backend_users_rpc_address` /
   `skiller_rpc_address` / `jobsimulation_rpc_address` **all with no default**. The literal is supplied by
   `infrastructure`.
5. **Self-contradiction inside the same file:** `cms.md:18` (*"the deletion itself lands in `infrastructure`,
   **which has never been in any clone set we have**"*) and `cms.md:82-84` (*"Whether
   `infrastructure/terraform/production/services.tf` still declares `module.cms_euwest1` is **not visible to
   this corpus**"*). A doc that states it cannot see the production terraform cannot report what it "still
   names". Brief rule 5 applies verbatim.

The *address* may well be right (app's own KB asserts it, present-tense at `b948604f:…/service-dependencies.md:46`
and **past**-tense at origin/main `:50-53` — *"it used to reach … folding it in at v9.0 closed that edge"*).
What is unsupportable is the source attribution, which is the part a reader would act on.

---

## Seat C

### C B1 | `corpus/services/customerio-sync.md:140` | **UPHELD** | IN-SCOPE
**Predicate:** points at an `external_services.md` section that does not exist, and re-publishes the fossil
the same file retracts.

**evidence (re-derived):** `corpus/architecture/external_services.md` `##` headings, enumerated from source:
`High-Level Summary` · `Clerk` · `Directus` · `GraphQL Gateway — WunderGraph Cosmo Router` · `AI Providers` ·
`LiveKit` · `AWS Chime SDK` · `Development Setup Summary` · `Production Deployment` · `Troubleshooting` ·
`Related Documentation`. **No Customer.io section, no Brevo section.** `grep -ni 'customer\.io\|customerio'`
→ **5 hits, all the `customerio-sync` service name** in compose/`repos.yml` prose (`:174, :306, :850, :859,
:860`). `grep -ic brevo` → **0**. So *"Customer.io as an integrated SaaS"* is false in both directions.
It also contradicts `customerio-sync.md:18-21` 122 lines above — *"**The name is a fossil.** The destination
has been **Brevo**, not Customer.io"* — which I confirmed against the platform's own package doc
(`app/internal/customeriosync/doc.go` @ `2035f9a`). Small blast radius; a verified-false claim about the
corpus's own contents plus a same-file self-contradiction.

### C B2 | `corpus/services/backend.md:33-34` | **UPHELD** | IN-SCOPE
**Predicate:** the app-owned-stream enumeration omits `backend` while asserting completeness at four.

**evidence — I derived the SET first, then compared cardinality (brief rule 4).** Every publisher/subscriber
constructor in `app` Go source @ `b948604f`, `_test` excluded
(`git grep -n "NewPublisher\|AddSubscriber" b948604f -- '*.go'`):

| stream | publisher | subscriber |
|---|---|---|
| `serviceName` (= `"backend"`, `main.go:214-215`) | `main.go:287` | `main.go:1320` |
| `SKILLPATH_STREAM` | `main.go:637` | `main.go:1274` |
| `CMS_STREAM` | `main.go:1039` | `main.go:1303` |
| `JOBSIMULATION_STREAM` | `internal/jobsimwiring/wiring.go:180` | `main.go:1285` |
| `AI_USAGE_STREAM` | `internal/jobsimwiring/wiring.go:127` | `main.go:1305` |
| `SKILLER_STREAM` | **none** | `main.go:1276` |

The only other publisher constructor in the tree is `internal/deadletterqueue/dead_letter_queue.go:38`, a
bare `redisstream.NewPublisher` that re-publishes onto a message's original topic and owns no stream.
**Cardinality of "app is both producer and consumer" = 5**, not 4; the banner omits `backend`.
*"`skiller` is **NOT a fifth member**"* converts an illustrative list into a completeness claim, and gets the
ordinal wrong — skiller would be a sixth. The same file's `:264` (pinned to `b948604f`, the graded ref)
states the correct set: *"four of the five application streams — `backend`, `skillpath`, `jobsimulation`,
`cms` — plus the `AI`/`ai_usage` usage stream"*. Two passages, one predicate, incompatible enumerations —
brief rule 5.

### C B3 | `corpus/services/backend.md:50-52` (and `:12`) | **REJECTED** | IN-SCOPE
**Predicate booked:** *"jobsimulation's ECS service is already destroyed"* is asserted under a standard the
same bullet forbids for `cms`, and contradicts root `CLAUDE.md`.

**evidence (re-derived):**
- The `CLAUDE.md` leg is **false at the audited tree**. `grep -n "jobsimulation_euwest1" CLAUDE.md` → **rc 1,
  0 hits** (`cms_euwest1` also 0). The current text is `CLAUDE.md:189-191` — *"teardown is **M810** — and it
  is **uneven**: it has **LANDED for jobsimulation** (`6092c6d2` deleted the ECS service *and* the ECR
  repository) and has **not moved for cms** … Do not generalise one row to the other"* — and `:240`
  *"**M810 has LANDED for the ECS service**"*. The seat quoted a superseded `CLAUDE.md` (repaired at
  `e858fd4`, iter-98) rather than the tree it was auditing.
- The asymmetry leg does not hold either: the distinction is stated and defended in **four** places —
  `backend.md:36-49` (cms's block *"has not moved"*), `platform-migration-status.md:89` (*"**Do not
  generalise M810 from this row:** `cms` has not moved"*), `jobsimulation.md:65` (*"the ECS service is
  destroyed; this is no longer a rollback path"*) and `CLAUDE.md:189-191`/`:241`. The corpus is internally
  consistent, and the evidence classes genuinely differ: I opened `jobsimulation` `6092c6d2` and its tree —
  the commit subject is *"remove the jobsimulation ECS service and ECR repository (M810)"*, the body
  enumerates exactly what the deletion destroys, and `terraform/main.tf:15-22` at that ref is the
  decommission comment where the `module "jobsimulation"` block used to be. `cms/terraform/main.tf:39` still
  declares its module. Code-declared-vs-not is a real, measured difference.
- What remains is only that *"destroyed"* states an **applied** state from a **code** deletion. That is one
  inference step, it is taken identically by the fenced map (which this bullet cites as authoritative), and
  the corpus discloses the remaining M810 step in the same breath.

**class: mis-read** — the contradiction the booking rests on does not exist at the audited tree.

### C B4 | `corpus/services/backend.md:18-19` | **REJECTED** | IN-SCOPE
**Predicate booked:** *"compose now declares **five** services"* is an unresolved-`include:` count and
collides with a different "five" in `CLAUDE.md`.

**evidence (re-derived):** at platform `0c91421d`, `docker-compose.yml` top-level keys are exactly
`include:` `services:` `networks:`; the `services:` block declares **5** (`sentinel:5`, `backend:28`,
`studio-desk:112`, `next-web-app:143`, `gotenberg:170`), and `common.yml` adds `postgresql:2` + `redis:24`
→ 7 effective. The sentence's subject is what **`838d907`** changed, and it pairs the number with
"`repos.yml` **four** entries" — both are file-scoped counts of the two files that commit edited, and both
are **exact**. The corpus states the qualifier where the project total is the subject (`cms.md:53-54` and
`jobsimulation.md:52-53`: *"the **five** services compose declares … (**seven** effective, once
`include: common.yml` adds the `postgresql`/`redis` floor)"*) — same five, no conflict.
`CLAUDE.md`'s "five containers" answers a different question (what `core` **starts**); both are true, and
two true answers that coincide numerically are not a contradiction.

**class: already-true** — the count is exact for the predicate the sentence states.

---

## Seat D

### D B1 | `corpus/services/jobsimulation.md:49-50` | **UPHELD** | IN-SCOPE
**Predicate:** identical to **B B3** — *"Production terraform still names
`http://backend.internal.anthropos:8081`"*, asserted about an artifact the same file declares unreadable.

**evidence (re-derived):** the same three-instrument measurement recorded under B B3 (0 `.tf` hits at any
ref, in any clone, by `git grep` **and** by a working-tree grep that sees untracked files and both nested
`studio` checkouts; the literal exists only in `app/knowledge/service-dependencies.md`, and at port 8080 in
graphql-wundergraph). Same-file self-contradiction at **`jobsimulation.md:74-75`**: *"Whether
`infrastructure/terraform/production/services.tf` still declares `module.jobsimulation_euwest1` is not
something this corpus can see — `infrastructure` has never been in the clone set."*
`messenger/terraform/main.tf:82` names `"${var.jobsimulation_rpc_address}"`, declared **defaultless** at
`variables.tf:92-95`.

### D B2 | `corpus/services/sentinel.md:85` | **UPHELD** | IN-SCOPE
**Predicate:** *"the only service address compose sets"* / *"the one cross-process edge a local stack has"*
is refuted nine lines after the line it cites — `GOTENBERG_URL`.

**evidence (re-derived), platform `0c91421d`:** inside the **same** `backend` `environment:` block that
carries `:48 AUTHORIZATION_ADDRESS=http://sentinel:8087`, line **`:57`** reads
`- GOTENBERG_URL=http://gotenberg:3200`. `gotenberg` is a declared compose service (`:170-183`,
`image: gotenberg/gotenberg:8`, `ports: 3200:3200`) carrying **`:183 profiles: [core, backend, all]`** —
the **default** profile, so it starts on a bare `make up`. `app` @ `b948604f` reads it at
`main.go:244` (`GotenbergURL: os.Getenv("GOTENBERG_URL")`) and
`internal/web/backend/coursebuilder/handler.go:244`, and `internal/converter/gotenberg.go:31` does
`http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)` on a 90 s client
(`:13`). That is a genuine cross-process HTTP hop to a second container on the default profile. The narrow
reading does not save it: the same sentence enumerates `gotenberg` among the services it checked, and scored
it only on *not carrying `AUTHORIZATION_ADDRESS`*. The RPC half of the sentence **is** right — I confirmed
`git grep '_RPC_ADDR' 0c91421 -- docker-compose.yml common.yml .env_example` → 0; the over-generalisation
from "no RPC edge" to "no cross-process edge" is the defect. `architecture_overview.md:321` states the
correctly-scoped version (*"the only cross-process **RPC** edge"*), which is what these three sites were
reaching for. Twins recorded by the seat and confirmed by me: `jobsimulation.md:145-146` and
`platform-migration-status.md:93`.

### D B3 | `corpus/services/askengine.md:88-89` | **UPHELD** | IN-SCOPE
**Predicate:** a failed Bedrock init returns from `main()` — the whole backend exits; nothing is "disabled".

**evidence (re-derived), `app` @ `b948604f`** (the doc names no ref; identical at origin/main `:467-471`):
```
410  bedrockClient, err := askengine.NewBedrockClient(serverContext)
411  if err != nil {
412      logger.Error("bedrock client unavailable; talk-to-data disabled", "error", err)
413      return
414  }
```
Enclosing construct checked three ways, per §5 rule 10: (1) `grep -n '^func '` over `b948604f:main.go`
returns **only** `newEntClientFromDriver:186`, `withStatementTimeout:194`, `main:212` — nothing between
`:212` and `:410`; (2) a scan of lines 212-414 for `func(` / `go func` / `wg.Go` / `errgroup` returns
**nothing**; (3) the `return` is at two tabs inside a one-tab `if`, i.e. the top level of `func main()`.
So the process ends before the RPC mux, the GraphQL server, the Echo router, the Asynq pools and the Redis
subscribers are started. The corpus took the platform's own misleading log string
(`"talk-to-data disabled"`) as a description of behaviour — the exited-0-in-silence class.

### D B4 | `corpus/architecture/platform-migration-status.md:78-81` | **UPHELD** | IN-SCOPE
**Predicate:** the census says **26** service names ever in `docker-compose.yml`; the set is **25**, and the
26th token is the `app-network` **network**.

**evidence — I re-derived the SET, not the sum (brief rule 4).** Over all **80** commits
`git log --follow` reports for `docker-compose.yml` (I verified `git cat-file -e $c:docker-compose.yml`
succeeds for **80/80** — no rename gap), tracking the current top-level YAML section and collecting only
2-space keys under `services:`:

```
backend chromedp chronos cms customerio-sync directus gotenberg graphql intelligence jobsimulation
messenger nats next-web-app postgresql realtime redis roadrunner sentinel simulator skiller skillpath
storage studio-desk web-app wundergraph
```
= **25**. The naive section-blind derivation over the same history
(`git log -p --follow | grep -oE '^[+-]  [a-z0-9_-]+:'`) returns **26**; `comm -3` between the two sets
yields exactly one token: **`app-network`**. Independently confirmed that the file has only three top-level
keys in its whole history (`include:` `networks:` `services:`) and that `networks/app-network` is the
**only** 2-space key ever to sit outside `services:`.

Two consequences: the stated count is wrong for the stated predicate (*"every **service**"*), and the
passage's own audit instruction — *"a name they return that has no row is a gap"* — turns the error into a
false alarm for anyone who re-runs it. The companion figure in the same sentence is **correct** and I
enumerated it independently: `repos.yml` has had exactly **14** names ever, and the doc's list is
set-identical. This is the one passage whose thesis is *"Completeness is measured, not asserted"*.

---

## Deduplication

| predicate | anchors |
|---|---|
| *"Production terraform still names `http://backend.internal.anthropos:8081`"* — unsupportable + same-file self-contradiction | **B B3** `cms.md:196` (twin `:55`) · **D B1** `jobsimulation.md:49-50` |

**One predicate, two anchors.** Both bookings are reported and both upheld; they enter `N` **once**.

No other pair collapses. In particular **C B2** (`backend.md:33-34`, the Redis-stream enumeration) and
**D B2** (`sentinel.md:85`, the cross-process-edge claim) only *look* adjacent — one is about which streams
`app` both publishes to and subscribes to, the other about which service addresses compose sets; different
sets, different sources, not collapsed.

The three sites of D B2's predicate (`sentinel.md:85`, `jobsimulation.md:145-146`,
`platform-migration-status.md:93`) were booked by the seat as **one** finding with twins recorded, and I
graded it that way — one anchor booked, one predicate.

---

BOOKED=13 UPHELD=10 REJECTED=3 IN-SCOPE-UPHELD-BLOCKERS=9

(10 upheld bookings collapse to **9** distinct in-scope predicates, per the dedup table above.)
