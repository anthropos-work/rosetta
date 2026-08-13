# Seat C — iter-48

Repo: `/Users/marco/workspace/anthropos/rosetta` @ `m257x/platform-realignment` HEAD `cabc3b1`.
Ground truth: `stack-demo/app` @ **`5ba17044`** (v1.363.2), `stack-demo/platform` @ **`2adcf71`**,
`stack-demo/graphql-wundergraph` @ `60c229f`, `stack-demo/{cms,jobsimulation,roadrunner,sentinel,messenger,next-web-app}`,
and the rext authoring copy `.agentspace/rosetta-extensions` @ `932554e`.

## Coverage (file, wc -l, lines read)

| file | `wc -l` | lines read |
|---|---|---|
| `corpus/architecture/alignment_testing.md` | 521 | **all 521** |
| `corpus/architecture/architecture_overview.md` | 346 | **all 346** |
| `corpus/services/cms.md` | 254 | **all 254** |
| `corpus/services/sentinel.md` | 166 | **all 166** |
| `corpus/services/messenger.md` | 128 | **all 128** |
| `corpus/services/README.md` | 79 | **all 79** |
| `corpus/services/db-backup.md` | 31 | **all 31** |

Every file was read top-to-bottom in one Read call each; no file was sampled. All seven files end with a
trailing newline (`tail -c1` = `0a`), so `wc -l` == line count — no off-by-one hole.

**Search hygiene (rule 1 + 2).** Every grep in this pass ran with stderr visible and `echo "EXIT=$?"`
appended; two rejections were caught and re-run (`--include=*.go` glob-expanded by zsh → re-run without it;
`docker-compose*` no-match → re-run with an explicit path). **Positive control:** I ran the toy alignment
harness live —
`go run ./cmd/alignctl run --dna examples/toy/dna.json --runner "go run ./examples/toy/cmd/toyrun" --golden-dir examples/toy/golden`
— and it reproduced `alignment_testing.md:296-306` byte-for-byte (`overall 86.7% / critical 100.0% / 5-of-6`,
`FAIL Greet/padded-name`). A second positive control: `grep -l 'ent.Schema' *.go | wc -l` in
`app/internal/data/ent/schema/` returned the expected non-zero 135.

**What I could NOT measure.** `clerkenstein/alignment/scripts/gate.sh` cannot run on this box —
`go build ./cmd/clerkrun` dies at `module lookup disabled by GOPROXY=off` fetching
`github.com/anthropos-work/colony v0.34.3`. So `alignment_testing.md:252-258`'s **score numbers**
(100%/100% Go, etc.) were verified structurally (gene counts, capability counts, golden-file counts, gate
defaults, exit-code constants — all exact) but **not re-measured live**. Stated rather than assumed.

## Blockers

| # | site | the false claim (verbatim) | what is TRUE | citation |
|---|---|---|---|---|
| 1 | `corpus/architecture/architecture_overview.md:299` | "**16 schemas carry an `organization_id` with no policy at all**" | **23** schemas carry an `organization_id` with no policy at all. 16 is only the *neither-mixin* subset; the other **7** are the `OrganizationIDMixin{}` users, which by that mixin's own definition carry a plain `organization_id` and declare **no** `Policy()` either. The doc's own linked target says so in the same words, with the other number. | **Corpus twin:** `corpus/architecture/security_compliance.md:72-73` — *"and **23 carry an `organization_id` with no policy of any kind.** Sixteen of the 23 have neither mixin"* (identical predicate, different number; that file is the target of this very sentence's `see …` link). **Re-measured at `app` @ `5ba17044`, `internal/data/ent/schema/`:** `grep -l 'ent.Schema' *.go` → **135**; `grep -l 'OrganizationMixin{}' *.go` → **30**; `grep -l 'OrganizationIDMixin{}' *.go` → **7** (`category.go`, `jobrole.go`, `similarity.go`, `skill.go`, `specialization.go`, `studio_document.go`, `studio_task.go`); the `comm -23`-then-`grep '"organization_id"'` derivation → **18**, minus `org_membership.go` (self-policed) and `academy_feedback.go` (`UserMixin{}` owner filter) → **16**; and `grep -l 'func (.*) Policy() ent.Policy' *.go` returns **only four files** — `mixin.go`, `organization.go`, `org_membership.go`, `user.go` — so none of the 7 mixin users is policed. 16 + 7 = **23**. |

Why this is a blocker and not an anchor nit: the surrounding sentence is the *summary* of Layer 1 for every
reader who never opens `security_compliance.md`, and the error runs in the **"isolation is handled"**
direction — the exact failure direction that file's own four-times-wrong fence warns about. It understates
the unpoliced surface by 30%.

## Minors

1. **`corpus/architecture/alignment_testing.md:193` — drifted anchor.** *"`gate.sh:61` calls `alignctl dna
   coverage --dna … --if-declared`"*. The call is real and the described semantics are exactly right
   (verified in `cmd/alignctl/dna.go:51,73-80`: `--if-declared` returns 0 with a loud `NO COVERAGE CLAIM`
   only for a DNA that declares no surface; a declared-but-uncovered endpoint still returns 2). But the line
   is **69**, not 61 — `grep -n "dna coverage" clerkenstein/alignment/scripts/gate.sh` → `69:`.

2. **`corpus/architecture/alignment_testing.md:513` + `:511-519` — incomplete Layout block.** The `cmd/alignctl`
   line reads `run | capture | dna list|diff|validate` and omits `dna coverage`, which the same doc documents
   as real and binding at `:245` and which I confirmed works. (In fairness `alignctl`'s own `usage` text also
   omits it — a tooling gap, not the doc's invention.) The `internal/` listing likewise omits `internal/canon`
   (actual dirs: `canon`, `compare`, `dna`, `outcome`, `report`).

3. **`corpus/architecture/alignment_testing.md:3` — stale header.** *"**Last updated:** 2026-06-06"* on a doc
   whose body carries M218/M219 corrections and v2.x milestones. Harmless, but it is the first line a reader
   uses to decide whether to trust the page.

4. **`corpus/architecture/architecture_overview.md:37` — dead intra-document anchor.**
   `(see [AI Providers](#ai-providers) below)`. There is no `## AI Providers` heading and no
   `<a id="ai-providers">` anywhere in the file (`grep -n "^#" architecture_overview.md` lists 20 headings,
   none matching; `grep -n 'id="ai'` → no hits). The referent is a **list item** at `:244`, not a heading, so
   the link resolves nowhere. The *claim* it supports ("EU-resident by default, not an EU-first ladder") is
   fully correct — I verified it end to end (see note below).

5. **`corpus/services/README.md:20` — set-membership slip.** *"And **three of the four** (cms, jobsimulation,
   roadrunner) still start CONTAINERS locally"*. `roadrunner` is explicitly **not** one of "the four" — the
   same blockquote says so five lines earlier (`:15` *"`roadrunner` is the fifth, and it is different"*).
   Should be *"three of the five"*. The substance (which three containers start) is correct and is enumerated
   inline, and both underlying facts are verified: `docker-compose.yml:83/:144/:281` with
   `profiles: [graphql, …]` at `:140/:187/:309`, and `roadrunner/terraform/main.tf:19` = `1` vs
   `cms/…:39` = `0` and `jobsimulation/…:40` = `0`.

6. **`corpus/services/cms.md:110` — a package listed that is not in the file.** The `requirements.txt`
   gloss names *"openai, anthropic, mistralai, rich, pyyaml, **python-docx**, requests, jinja2, pytest,
   pytest-asyncio"*. `cms/studio/requirements.txt` (and the identical `app/studio/requirements.txt`) contains
   exactly: `openai, anthropic, rich, pyyaml, requests, jinja2, mistralai, pytest, pytest-asyncio` — **no
   `python-docx`**. Nine entries, not ten.

7. **`corpus/services/sentinel.md:30` — incomplete matcher description.** The `m` row describes the tier
   check and the `count <= max` check but omits the `r.feat == p.feat` conjunct. Actual:
   `m = ( g(r.sub, p.tier) || p.tier == 'TIER_FREE' ) && r.feat == p.feat && r.count <= parseFloat(p.max)`
   (`sentinel/internal/authorization/casbin.go`, `ModelConf`). The sibling `m6` row *does* name its `feat`
   match, so the omission reads as a difference where there is none.

8. **`corpus/services/db-backup.md:27` vs `corpus/services/README.md:37` / `architecture_overview.md:171-175`
   — framing mismatch.** db-backup.md says *"runs in production and staging environments"*; the other two
   file it under **"Production-only"**. Both are compatible if "production-only" is read as "not in local
   compose" (which is how the corpus uses it elsewhere), but the two words sit badly next to each other.
   No repo available to adjudicate — `db-backup` is not cloned under `stack-demo/` or `stack-dev/`, so
   nothing in this doc was verifiable against source at all.

---

### Verified-clean (so the next iteration does not re-spend budget here)

Everything below was checked against platform source and found **accurate**, including the anchors:

**`architecture_overview.md` — the newly-written AI-Providers paragraph (`:244-260`), read with extra
suspicion per the brief, is correct in mechanism AND in every anchor.** `getClient` at `internal/jobsimulation/ai/ai.go:259`,
`azureClientEu` default `:264`, `flag_use_azure_us` `:267`, `azureClientUs` `:275` (doc's `:262-276` brackets
it); `isThrottlingError` `:129`, applied `:166` and `:325`; Bedrock pinned `config.WithRegion("eu-west-1")`
`:87` + `anthropic.NewAnthropic` `:92` (doc's `:85-88`); Mistral `internal/cms/studio/markdownManager.go:11,19`
and `studioManager.go:583`. The *fourth* EU-exit — `ai_vendor` unset — I re-derived independently rather than
trusting the citation: `jobsimulation.go:905` `AIVendor *AIVendor` (nullable DTO field) → `:1302`
`aiVendor := simulation.Openai` when nil → `simulator/ai/ai.go:58-59` maps `simulation.Openai` →
`internalAi.Openai` → `getClient` → `openai.NewOpenAI(openaiKey)` at `ai.go:80`. Direct US OpenAI on the first
attempt, no flag, no 429 — as written. Cross-refs `external_services.md:545` and `:569` both land exactly.

**Router / compose / federation.** No `graphql` service in `docker-compose.yml` (services are exactly
sentinel, backend, jobsimulation, cms, storage, customerio-sync, messenger, roadrunner, studio-desk,
next-web-app, gotenberg); no `graphql-wundergraph` in `repos.yml`; **zero occurrences of `5050`** anywhere in
compose or `.env_example`; frontends at `:8082/graphql/query` (`:318,:334,:352,:361`). Prod router still
`service_desired_count = 1` at `graphql-wundergraph/terraform/main.tf:20`, `port = 8080` in its `locals.tf`.
`915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` (−762) **and** `schemas/jobsimulation.graphqls`
(−860) in one squashed commit whose subject says "2→1" — so the corpus's **3→1** reading is right and its
note that the subject undercounts is right. Platform `b56d731`+`360efd4` merged as `2adcf71`, 2026-07-31.
"Six Go services on a bare `make up`" checks out.

**cms.md.** `cms/terraform/main.tf:39` = 0; `docker-compose.yml:144` in `graphql`; `repos.yml:14-16`
`migrations: false # legacy`; `CMS_RPC_ADDR=http://cms:8091` at `:47/:104/:256`; `app/main.go:1196-1202`
carries *"Additive + DORMANT … until the M809 re-point"* (the phrase is on `:1199`, as messenger.md cites);
`app` makes **no** outbound cms RPC (the only `CMS_RPC_ADDR` hit in app's Go tree is that comment);
`20260724132049_cms_data_model.sql` creates exactly the six named tables with the same names;
`REDIS_CMS_CACHE_INDEX` default 5 (`main.go:988-989`); `CMS_STREAM` subscriber merged via `.AddHandler`
(`main.go:1279,1287,1294`); the Directus webhook **fails closed** — `validWebhookSecret` returns false on an
empty secret and `router` 401s before dispatch (`internal/cms/directus/webhooks/router.go`), while the
standalone `cms` receiver at `/webhooks/` (`cms/cmd/root.go:242`) has **no** auth at all; `cms/go.mod:3`
`go 1.26.4`; `cms/Dockerfile:2` golang, `:23` `FROM python:3.11-slim`; ports 8090/8091; `gen.py:18-28`
`parse_argument`/`parse_known_args` and exactly **nine** `add_argument` calls at `:484-492` with no
`--template`; `studio/CLAUDE.md:12-14`; `internal/skillpath/session.go:205-207` in-process
`GetSkillPathDomain`; the jobsimulation compose block has **no** `DIRECTUS_BASE_ADDR` (only cms, `:164-165`);
app v1.360.0/.1/.2/.3 all exist in `CHANGELOG.md` with the described contents; `app/studio/` is **not
git-tracked** (consistent with the CI `additional_repo` pull); the 23 jobsim tables are exactly 23
`CREATE TABLE`s in `20260722081626_jobsim_data_model.sql`.

**sentinel.md.** Casbin v3; **6** request defs / **6** policy defs / **3** role groupings (`g`/`g2`/`g3` with
2/3/2 args as documented) / **6** matchers, all in `internal/authorization/casbin.go`; `AUTHORIZATION_ADDRESS`
in **exactly three** compose blocks — `:45` backend, `:99` jobsimulation, `:160` cms — and nowhere else;
messenger imports no authorization client (its only "sentinel" hit is an English-language comment in
`pkg/aireadinessemail/override.go:378`); `terraform/locals.tf:4-5` = 256 / 128; `init_policy.sql:63-66` is the
taxonomy:write omission NOTE; the `content_creator` block runs `:88-118`; `local_superadmin_grants.sql` grants
p3 `org:feature:taxonomy:write` to `default/admin` with a commented-out p4 block; **no `manager` role** —
zero hits in `init_policy.sql`, and in tests only as the fixture string
(`casbin_test.go:17,24,76,83`; `rpc_test.go:113,300,439,482,483`); `make initdb` really is a hard-coded
`postgresql://postgres@localhost:5432/postgres?sslmode=disable` that never reads `DB_CONNECTION`
(`Makefile:3`); binary default `PORT=8080` (`cmd/root.go:47`) overridden to 8087 in compose; no `profiles:`
key on the sentinel service; `go.mod:3` `go 1.26.0`, `Dockerfile:2`/`Dockerfile.dev:2` `golang:1.26-bookworm`;
the four `MembershipRole` values at `app/internal/data/ent/enum/membership.go:8-15`.

**messenger.md.** `go 1.25.0`; `getbrevo/brevo-go v1.1.3` (`go.mod:11`), `osteele/liquid v1.8.1` (`:14`);
`Schedule`/`CancelScheduledMessage` both `CodeUnimplemented` at `internal/rpcsrv/rpcsrv.go:25-30`; binary
fallbacks `PORT=8080` `cmd/root.go:63`, `RPC_PORT=8081` `:64`, `REDIS_STREAMS_INDEX=2` `:107`;
`READONLY_DB_CONNECTION` at `cmd/root.go:147`; **`REDIS_WORKER_INDEX` is set in compose (=0) and has zero
occurrences in the messenger source** — the doc's claim exactly; the 2 h / 12 h staleness guards at
`internal/flow/jobsimulations.go:140-151`; the five `OrgSkillPath*` handlers inside
`internal/flow/flow.go:72-87`; `getSkillPath` reaching cms at `internal/flow/assignments.go:828`;
`depends_on: backend, cms, jobsimulation`; `SKILLPATH_RPC_ADDR` gone from compose with only
`SKILLPATH_STREAM=skillpath` left (`docker-compose.yml:64`).

**services/README.md.** The index really does enumerate **all 27** service docs (29 `.md` minus `README.md`
and `TEMPLATE.md`; table rows 10+5+8+4 = 27, none missing, none dangling).

**alignment_testing.md.** Gene/capability counts are exact for all five DNAs — `clerk-2.6.0` 14 caps/27
genes, `clerk-js-5` 6/9, `clerk-multi-1` 5/9, `clerk-deploy-1` 3/7, `clerk-express-1` 5/13 — and the golden
dirs hold exactly 27/9/9/7/13 files. **Only `clerk-2.6.0` declares a `consumed_surface`** (15 entries); the
other four declare none — precisely as `:201-204` states. `GetUser` really carries the two-sided
`hero-a`/`hero-b` variants and `MembershipOrgIdentity/real-org-eid` is still in the DNA as `standard`;
`Store.SeedOrgIdentity`/`LookupOrgEid`/`organizationWithEid`/`seedRosterMemberships` all exist
(`clerk-backend/store.go:126,138,151`; `cmd/fake-bapi/main.go:36,50,88`). `DNA.Validate()` does reject a
zero-critical-gene DNA (`internal/dna/dna.go:278-281`) and does call `validateConsumedSurface`; `alignctl run`
calls `Validate` **before** any scoring (`cmd/alignctl/run.go:38-43`); `ExitRegressed = 2` /
`ExitUnmeasurable = 3` are real named constants with the UNMEASURABLE banner (`run.go:133-135,142`).
The `alignment/` section has **no** `scripts/` dir (only `cmd`, `examples`, `go.mod`, `internal`,
`README.md`); module is `anthropos.dev/alignment`; `gate.sh` defaults are `RUNNER_PKG=./cmd/clerkrun`,
`DNA=dna/clerk-2.6.0.json`, `ALIGN_DIR=$base/../../alignment`. `clerkenstein/.github/workflows/alignment.yml`
**exists, is git-tracked** (`git ls-files` confirms), and its `:10-11` self-description matches the quoted
text verbatim. Five runners `clerkrun`/`jsfapirun`/`expressrun`/`deployrun`/`multirun` all present;
`deploy/colony-authn` present; `clerk-webhook/` genuinely uses `svix`;
`clerkenstein/knowledge/alignment.md` and `knowledge/architecture.md` both exist.
