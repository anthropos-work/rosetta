# iter-129 — the alarm was too loud, and the census was too narrow

**Type:** tik · **Run 82.** Four priorities in the run's own consequence order.

---

## 1. Priority 1 — `/api/invitations` settled at the ref: **token-authenticated, not unauthenticated**

**iter-128's count was right and its alarm was wrong.** The mechanism, measured at `app` `ad9f3c498` by
reading the `RegisterRoutes` **call site and the manager** rather than the mount comment:

| question | answer, with the anchor |
|---|---|
| what middleware is passed? | **`cors` and nothing else** — `RegisterRoutes(srv.e, cors.EchoCORSMiddleware(...))`, `internal/web/web.go:148`. No Clerk `authn`, deliberately: next-web-app renders the invite-landing and unsubscribe pages **before the caller has an account** (`web.go:145-146`) |
| is a credential required? | **Yes.** The path segment IS the credential: `base64url(HMAC-SHA256(email\|org_id\|invited_at, INVITATION_HMAC_SECRET))`, `internal/invitations/token.go:29-34`; `main.go:423-427` refuses to boot without the secret |
| is it checked **before** data is returned? | **Yes.** `GetInviteDetailsByToken` / `OptOutByToken` filter `Where(membershipinvitation.Token(token))` — `internal/invitations/invite.go:159`, `:194`. A miss returns `404 not_found` / `already_opted_out`: **no row, no email, no org name** |
| by what mechanism? | **A stored bearer capability, NOT a re-verified signature.** `TokenManager.ValidateToken` (`token.go:38`, constant-time) is called by **nothing outside its own test** — `git grep -n ValidateToken ad9f3c498 -- '*.go'` returns 8 hits, all in `token.go` + `token_test.go`. The HMAC supplies 256 bits of unguessability; the *check* is an equality match on an indexed column |
| what about the Ent privacy layer? | **Deliberately bypassed** — `privacy.DecisionContext(ctx, privacy.Allow)` at `invite.go:157`, `:190`, with the source stating the model in its own words: *"It backs a public endpoint — **token possession is the authorization**"* (`:154-155`, `:187-188`) |

**Verdict: no defect filed.** The design is stated in the source, the credential is required and checked,
and what the token gates (invited email, org name, inviter name) is the content of the invitation it was
mailed with. *"No Clerk middleware" is not "no authentication."*

Two observations recorded so nobody re-derives them, neither a defect: the token is written to application
logs on the miss and opt-out-success paths (`handlers.go:59`, `:62`, `:99`; `invite.go:203`, `:218`); and
`/api/invitations` is the second of **two** public token endpoints — the other, root-mounted, **does**
re-derive the HMAC and `401`s on mismatch.

### 1a. And rule 57 bit its own worked example a **second** way

*"Eleven groups"* is correct **for groups** — re-derived independently, repo-wide:
`git grep -nE '\.Group\("' ad9f3c498 -- '*.go' | grep -v _test.go` → exactly those 11. But `app` mounts
**seven more routes directly on the root `e`, in no group at all**, so no group-level statement reaches
them: `/graphql/query` (`backend.go:317`), `/api/webhook/directus` (`:324`),
`/ai-readiness/unsubscribe/:token` (`notifications/handlers.go:41`), `/api/schema.json` (`:117`, **no
gate**), `/content/catalog.json` (`content.go:23`, **open by design** — `content_admin.go:32` says so), and
two `colony.Development`-only mounts (`:309`, `:315`).

**Run 81 widened from one file to the whole service; it did not widen from *groups* to *routes*.** The
honest shape is **11 groups + 7 ungrouped root mounts**.

**And the count's FOURTH site.** `architecture_overview.md:405` still read *"2 of its **6** Echo groups"* —
run 81's own corpus-wide sweep of that very predicate did not reach it. Found by widening the regex, which
is the same lesson one layer up.

**Repaired:** `security_compliance.md` (the table row, two correction boxes, the honest-statement bullet),
`architecture_overview.md:405`, `platform-alignment.md` § 5 rule 57 (its worked example now carries both
corrections). *A comment is testimony; the call site is evidence* — rule 22's distinction, applied to code
comments instead of commit messages.

---

## 2. Priority 2 — the out-of-census consequence surface, swept, with **two accountings kept apart**

### 2a. The denominator, re-derived — and iter-128's numbers do not reproduce

`iter-129/out-of-census.py` **imports** `claim_census_guard` (the census) and `iter-124/triage-predicate.py`
(the sealed `C1` regex) rather than copying either, so the two enumerations cannot drift.

```
/usr/bin/python3 iter-129/out-of-census.py <rosetta-root>
```

| | census SCOPE (`corpus/services/*.md` + `corpus/architecture/*.md`) | OUTSIDE it (`CLAUDE.md` + `corpus/ops/**` + `corpus/tools/**` + `corpus/*.md`) |
|---|---|---|
| files | 41 | **52** |
| consequence-class sentences, ALL | 1,324 | **2,012** |
| consequence-class **unevidenced assertions** | 327 | **724** |

**Census reach: 39.7 % of all consequence-class sentences, 31.1 % of the unevidenced ones.**

⚠️ **iter-128 published 37.1 %, from 1,598 in-scope vs 2,714 out.** Re-derived with the census's own
splitter, block model, heading filter and imperative/subject-token filters, the pair is **1,324 / 2,012**.
Likewise *"`CLAUDE.md` alone holds 108"* — measured here, `CLAUDE.md` holds **99** (108 is
`corpus/ops/secrets-spec.md`'s count, one row above it in the same table). **The substance is unchanged and
the arithmetic is not**: the census reads well under half the surface either way. The reproducible pair is
this one, because the invocation is stated and the instrument is the census's own.

### 2b. Two accountings, and they are not merged

- **Clause 5** names `corpus/services/**` + `corpus/architecture/**`. **NOT re-cut, not reinterpreted, not
  widened by this work.** Its in-scope unevidenced consequence set is 327 — down from 340 because
  Priority 1's repairs *added evidence*, not because anything was re-scoped.
- **The user's standing ask** — *the corpus, the skills and rext aligned to the platform* — has no such
  boundary. A false security sentence in `corpus/ops/**` or in `CLAUDE.md` misleads a reader exactly as
  much as one inside clause 5's scope. **724 sentences, 47 files.** This is the accounting the sweep
  below belongs to, and **no number from it is added to clause 5's.** `F4` binds: the census is an
  instrument, not clause 5's grader.

### 2c. The sweep — 724/724 read, in six slices, every clone at a ref

**Found and repaired this iter (the classes, not a raw list):**

**(i) The milestone's own core drift class, in the ops docs — 39 stale schema-qualified tokens across 8
files.** `app` migration `20260729133514.sql` dropped `local_jobsimulation_sessions` +
`local_skill_path_sessions` (*"Collapse the `local_*` session mirrors"*), jobsim-in-app re-created the
fan-out in `public`, and **a fresh stack never creates the `jobsimulation` or `skillpath` schema at all.**
The tooling moved — **44 of 44 `CopyRows*` call sites in `stack-seeding/seeders/` pass `"public"`** — and
the ops corpus did not. Repaired in `seeding-spec.md`, `stories-spec.md`, `playthroughs.md`,
`recipe-skill-progression.md`, `session-clone-spec.md`, `media-substrate-spec.md`,
`content-stories-spec.md`, `demo/README.md`, `cockpit-spec.md`, plus **`safety.md`** (the write-set
enumeration in the safety contract) and **`verification.md`** (a bring-up *probe floor* written against a
dropped relation). Two of these were **runnable defects, not prose**: `recipe-skill-progression.md`'s
verification `psql` block would have `42P01`'d on all four of its counts.
**Mirror-write instructions are RETRACTED, not re-qualified — there is no mirror to write.**

**(ii) The refuted `@anthropos.work`-only predicate, again, in a *guide*.** `run_guide.md:284` told the
reader they *"should land on the Clerk sign-in page"* and `:291` diagnosed a *"domain not allowed"* failure
as *"account not on `@anthropos.work`"*. Measured at `ant-academy` `22df69dd8`: `/` is an **explicitly
public** route (`code/proxy.js:170`, in the matcher at `:112`, *"M4 public catalog — the front door"*), and
there is **no email-domain allowlist anywhere in the repo** (0 hits). The same page contradicts itself 267
lines earlier. **iter-128 swept 14 sites of this predicate; this is a 15th its enumeration did not reach.**

**(iii) `CLAUDE.md`'s rext section list omitted two of eleven sections** — `stack-secrets` and
`playthroughs`, both with their own `go.mod`, and `stack-secrets` is the one `CLAUDE.md:561` itself depends
on. **This is the enumeration every session loads**, so an agent looking for either tool concluded it did
not exist.

**(iv) The M23 Directus cutover names a service that does not exist — 7 sites.** `directus-local.md` (the
dedicated doc), `snapshot-spec.md` ×4 and `safety.md` ×2 all say the cutover re-points **`cms`'s**
`DIRECTUS_BASE_ADDR`. There is no `cms` service; `app` reads it in-process. The tooling recorded the
correction and the corpus did not — `gen_injected_override.py:79` states the `cms` branch *"never matches
it on a current clone"*, and `:70-75` records the measured cost: **96 Directus lines on `demo-1`, all
403**. The correction had reached `cms.md` and `jobsimulation.md` at iter-24 and stopped there — **rule 54,
four docs short.**

**(v) `toolchain_overview.md` said Atlas manages `public, cms, jobsimulation, skillpath`.** Measured: two
pipelines, both in `app` — `env "local"` → `public` and `env "sentinel"` → `sentinel`
(`app/atlas.hcl:6-19`, `:50-64`, the latter added 2026-08-04).

**Found, evidenced, NOT yet repaired — routed as `FIX-M257x-iter129-sweep-residual`.** The sweep returned
more than one run can repair with the verification each deserves, and a half-verified repair is the defect
this milestone exists to catch. Each carries a `file:line` on both sides in the slice reports:
`staging-bringup.md` (sentinel-Atlas, the landed `CORS_EXTRA_ORIGINS` PR ×5 sites, the *"GraphQL is
unaffected"* claim refuted by 30 `OrganizationMixin` schemas, the colony bug fixed at v0.34.4, the
self-contradictory step-6 note), `staging_from_dump.md` (3 of 4 notification examples do not exist —
**password resets are Clerk's, so blanking `BREVO_KEY` does not suppress them**), `staging-clerk.md`
(Next.js 15.5 → 16.2; the *"subgraphs warm serially"* mechanism), `frontend-tier.md` (ant-academy
*"depends only on Clerk"*; the studio-desk profile), `demopatch-spec.md` (four patches → seven; LIFO),
`latency-budget.md` (**the fake FAPI does not validate `redirect_url`** — `clerk-frontend/server.go:414-423`
redirects verbatim; the doc credits the mock with a security property it does not have),
`build-budget.md` (phase attribution; eleven → twelve sub-phases), `setup_guide.md` (the
`env_file`-already-carries-it premise), `platform-alignment.md` (`:46-47` live anchors on rewritten code;
`:308-314` a mechanism refuted by both routes being public), `recipe-browser-login.md` (names the `cms`
container; the live path uses a network alias, not `extra_hosts`), `safety.md` (**the `Bunny` recording-key
provisioning path does not exist — 0 occurrences of `BUNNY_RECORDING_*` in rext**; `AssertClean`'s
*"every attempted write is recorded"* is true only of the **blocked** path, which `audit.go:82-96` says in
as many words; the attacker-gain enumeration omits the demo's bridged AWS credentials),
`secrets-spec.md` (*"the one secret"* is two; `platform/OPENAI_KEY` is `required·critical`, so the waived
class as written **false-fails** its own gate), `db-access.md` (the `274/733` split was measured on
`cms.similarities`, which the tooling still captures, not on `public.similarities`).

### 2d. And one rext defect the sweep found that I **could** fix, so I did

**The `$HOME/.aws/credentials` removal was fixed on the demo path at iter-88 and left dead on the dev
path.** `stack-core/gen_override.py:154` still read `if name == "jobsimulation"` — a service `d11a403`
deleted, whose identical bind `838d907` moved onto **`backend`** (`docker-compose.yml:99-100`). So on a
`dev-N` stack the mitigation has been a **no-op**, and `backend` kept the operator's real AWS credentials
mounted next to the hardcoded production `STORAGE_S3_BUCKET` (`:82`) that
`DEF-M257x-iter80-storage-prod-bucket` already tracks. **No test failed, because there is nothing named
`jobsimulation` left for a test to assert on.** Now derived from the resolved compose (ALL-not-ANY), with
**5 regression tests including a mutation control** that re-creates the literal and proves it selects
nothing. *The same stale-literal defect, unrepaired one family over.*

---

## 3. Priority 3 — the two blocks iter-128 named untested

### 3a. R3 re-audited **on this population** — 486 members, the largest untested block

`iter-129/sample-complement.py` imports iter-124's predicate + triage and iter-128's partition; the
fail-closed partition assert (`|C1| + |complement| == |tier-2|`) is retained.

```
/usr/bin/python3 sample-complement.py /tmp/tier2-129.json --audit 30 --seed 129 --rule R3
```

**28 of 30 agree = 93.3 %** (Wilson 95 %: **78.7 – 98.2 %**). Both disagreements run one way, `cite →
hedge`, and both are `chronos` — *"the subject is a repo no clone set contains"*, the same mechanical
sub-class iter-128 named in its R4 audit. **iter-124's 100 % was measured on C1 and does not hold here.**

Two would-be disagreements survived scrutiny and are recorded because assuming them would have inflated the
error: `studio-room.md`'s `{simid}_usage.json` **is** reachable (`stack-dev/studio-room/gen.py:458`), and
`shared_libraries.md`'s `colony/rpc` **is** reachable (`~/go/pkg/mod/…/colony@v0.34.3/rpc/intercepters.go:10`)
— neither is cloned by `make init`, and both are readable anyway.

**The complement's split, re-derived with R3 measured here rather than imported:**

| fate | iter-128 published | **iter-129, R3 measured on this population** |
|---|---|---|
| `cite` | ≈ 738 (90.0 %) | **≈ 706 (86.1 %)** — band 70.3 – 93.6 % |
| `hedge` | ≈ 33 (4.0 %) | **≈ 65 (8.0 %)** — roughly doubled |
| `drop` | ≈ 49 (6.0 %) | **≈ 49 (6.0 %)** |
| `fix` | 0, a floor | see § 3b |

Carrying a 100 % measured on a different frame **overstated `cite` by ~4 points**. That is the whole reason
the re-audit was owed.

### 3b. The `fix = 0` complement — read for falsity

The 820 were sliced by file into four balanced sets (`comp-slice-1…4.txt`, `sum == 820` asserted) and read
against source at refs. **`fix = 0` was 0-because-unread. It is no longer zero.**

| | read | **FALSE / materially misleading** | UNCHECKABLE |
|---|---|---|---|
| slice 1 (`service_taxonomy`, `ai_architecture`, `shared_libraries`, `dependency_map`, `jobsimulation`, `coursebuilder`, `cms`, `security_compliance`) | 205 | **6** | 18 |
| slice 2 (`external_services`, `studio-desk`, `architecture_overview`, `org-repos`, `frontend_architecture`, `hiring`, `gotenberg`, `clerk-integration`) | 204 | **8** | 20 |
| slice 4 (`ai-readiness`, `sentinel`, `studio-room`, `backend`, `askengine`, `storage`, `next-web-app`, `clerkenstein`, `academy-backend`) — **returned AFTER the close commit; folded in rather than dropped** | 206 | **11** | 1 |
| slice 3 (`alignment_testing`, `ant-academy`, `graphql-wundergraph`, `ai-labs`, `chronos`, `roadrunner`, `messenger`, `customerio-sync`, `platform-migration-status`, `skiller`, `db-backup`) — also **post-close** | 205 | **21** | 22 |
| **820 of 820 — EXHAUSTIVE** | **820** | **46 = 5.6 %** | 61 |

**`fix` over the complement is 46 of 820 = 5.6 %, measured, not a floor of unread.** All four slices
returned; the two that landed after the close commit were **folded in rather than dropped**, and the record
says which. The `chronos`/`db-backup` un-cloned class dominates `UNCHECKABLE` (19 of 61 are chronos
internals alone) — the same class § 3a's R3 audit found, now counted rather than inferred.

Four of slice 1's six cluster in **one** place — the *Embeddings & RAG* section of `ai_architecture.md`,
which is that document's own EU-residency argument: the embeddings client is **Azure EU, not OpenAI** (17
call sites, `skillerai.Openai` returns **0** repo-wide); the folded `ai.AI` interface is **not**
*"unchanged"* (9 methods → 8, and **Mistral is no longer an implementation** — it moved to
`internal/cms/studio/mistralocr`); there is **no Anthropic/Bedrock path in the skiller domain** at all
(`getClient` rejects that vendor); and the match cache is **Postgres, not Redis**. Slice 2's sharpest is a
**runnable instruction for a component destroyed in both states** (`external_services.md:806`, under a
heading with no retraction while the same page's banner at `:3` retracts it) and a **Directus webhook setup
line that produces a webhook rejecting every delivery** — the endpoint is **fail-closed on a shared secret**
the line never mentions, and the sender is a Directus **Flow**, not the Webhooks module (deleted in
Directus 11). All findings carry `file:line` on both sides and join the routed residual, **in the
complement's own accounting, never added to clause 5's**.

**Slice 4 landed one finding that indicts THIS iter's own repair, and it is repaired.** `clerkenstein.md:21`
carried the **same** rext-section under-count I fixed in `CLAUDE.md` under § 2c(iii) — *six* sections where
there are **eleven**. **My repair reached one cell.** `§5` rule 54 says a correction that reaches one cell
is not a correction, and I committed the violation inside the run that was applying the rule — the third
time this milestone has caught that exact shape (run 81's ant-academy regex, run 81's group count, this).
Both sites now name all eleven. Two further slice-4 findings repaired in the same pass, both in
`next-web-app.md`: *"`apps/web` is the only frontend in platform compose"* (**`studio-desk` is a second
compose frontend**, `docker-compose.yml:112`) and `GRAPHQL_SCHEMA_FOR_GEN` *"used by `graphql-codegen`"*
(**read by nothing** — 4 hits, all `.env.example`; `codegen.ts:9` hardcodes the endpoint).

### 3c. And slice 3's headline is a defect **in the fenced map** — clause 3's own deliverable

**`platform-migration-status.md:108`, the `ai` row, still read `library | library`.** Measured at
`ad9f3c498`: `app` **folded the library in-tree as `internal/ai` at `1e457fa70`** (2026-08-04) and dropped
the requirement — `app/go.mod`'s `anthropos-work` requires are `analytics-go`, `colony`, `proto`, `storage`,
`taxonomy`, **no `ai`**, and `go.sum` has **0** hits. By the map's own § 1 vocabulary (*a `library` is
imported as a private Go module*), **no repo a stack builds imports it**; only the frozen `cms` and
`jobsimulation` still require it, and nothing builds either. `internal/ai/module_import_guard_test.go:18-38`
is a one-way door against re-acquiring it.

**The corpus already knew and the map did not.** `shared_libraries.md` recorded the fold at **iter-102**, and
the *neighbouring* `authn` row in this very table carries exactly the right caveat — **the `ai` row is the
only library row iter-102's repair never reached.** `§5` rule 54 again, and this time inside the artifact
clause 3 is graded on. **The map's guard could not have caught it**: `platform_alignment_guard` fences the
map against `repos.yml` in both directions, and `ai` **is not in `repos.yml`** — it is a module, not a
clone. *A fence is green over its reach, and a library row is outside this one's.*

Four more repaired from slice 3, each verified at the ref before touching anything: `messenger.md:13-15`
(*"sends **and schedules**"* — 0 scheduler hits, `Schedule` is `CodeUnimplemented`, which the page said 4
lines down; and *"Liquid templating for the bodies"* — **Brevo** renders), `chronos.md:122-124` (the only
present-tense block on the page with no local fence, refuted by `wiring.go:192` and 0 platform hits) and
`:214`, `ant-academy.md:19` (the mobile bundler ships **nothing** — its deep-dive already said so) and
`:314` (**31** e2e specs and ~2,700 Vitest cases, not "~26" and "1000+" — the naive grep returns 38 by
catching `-snapshots/*.png`, so the invocation is stated), and `graphql-wundergraph.md:182` (the historical
subgraph set was **five**, not two — contradicting its own ladder 137 lines up).

**The remaining ~24 findings across slices 2–4 join the routed residual**, each with `file:line` on both
sides — including `ant-academy.md:356`, a **16th** site of the `@anthropos.work` predicate.

---

## 4. Priority 4 — the two structural items, decided rather than papered over

### 4a. `org-repos.md` has an owner — and it already had three

Measured before acting: `org-repos.md` is indexed at `corpus/architecture/README.md:12` **and**
`CLAUDE.md:439`; `observability.md` at `corpus/ops/README.md:57` and `CLAUDE.md:440`; both gained a plan
owner at `iter-123/progress.md` § 6. **The gap iter-128 named was real and iter-128 closed it.** Recorded
here rather than re-closed, because re-doing a discharged item is how a ledger loses its meaning.

### 4b. The body budget — raised **once, against a measurement**, and the probe that said it couldn't be

**iter-128's three probes had targets in two of three cases. Its search was narrower than its conclusion.**

| iter-128's probe | re-measured |
|---|---|
| § Standing backlog — *"7 of 14 appear nowhere else"* | **3, not 7.** `PERF-M256-parallel-lane`, `PT-M257-self-evaluation`, `PT-M257-talk-to-data` are each in `roadmap.md`; `platform-defect-register` is its own file. The probe searched `roadmap-vision.md` and concluded *nowhere else* |
| M255 provenance — *"`roadmap.md` has the numbers, not the provenance"* | **true of `roadmap.md`, and it asked the wrong file.** The contract's own table names *that milestone's `progress.md`* as the home for per-milestone measurement |
| § Process flags | **genuinely sole-owned.** No target, correctly |

**That is rule 57 pointed at a plan probe.** A mirror-search finds only content that is *already*
duplicated — exactly the content that does not need moving. The question is not *who else already owns
this?* but *who does the contract say should own it?*

**Done:** the three orphaned backlog items moved to `roadmap-vision.md` § Unscheduled backlog; the M255
provenance moved to `m255…/progress.md` § Provenance; § Standing backlog rewritten as a four-row index.

**Decided:** body `12,000 → 13,500`, frontmatter `2,600 → 1,860`. The old triple was **arithmetically
incoherent** — `12,000 + 2,600 = 14,600 ≠ 15,360`, so 760 bytes belonged to no budget; the new pair sums
**exactly** to the file cap. 13,500 is the measured floor (13,069 across the 11 sections after both moves;
13,281 with this iter's own disclosure line) plus ~1.7 %. **A re-raise guard is written into the contract**:
a future breach re-runs the ownership probe against **every** file under `knowledge/plan/`, not two of
them, and the per-field budgets — `phase:` above all — remain the real control. **Every budget is now met.**

---

## 5. Guards, suite, gate

**Guards: 22 members · 18 GREEN · 0 RED · 0 could-not-check · 4 not-run**, run before and after the
Priority-1 commit. `repair_postcondition` accepted the commit on the first attempt. Census **1,150,
unchanged** — the Priority-1 repairs added evidence to sentences already in scope. Invocation as § 5 of
iter-128.

**Suite: NOT RE-RUN, and not for want of trying.** Load was checked **first** this time, per iter-128's own
amendment: the external `hyperspace/anima8` project was live at **~6 of 12 cores** (`load1 7.44`) at the
start of this run and never quiet. **Re-running an instrument under the condition that invalidated it is
not a second measurement** — iter-128's words, and they apply to its successor. The counts stand at
**1,158 collected / 1,157 passed / 1 failed** (the standing documented RED), attested by two independent
runs on different trees. `FIX-M257x-iter128-suite-timing-unattested` stays open. The rext change here
carries its own targeted run instead: **26 passed** on the emitter's existing tests, **5 passed** on the
net-new regression file.

**Gate: unchanged at 4 of 5. No reading was taken and `P` is unmeasured.** Clause 5 is met only by a
reading that returns zero; **repair is not a reading, a triage is not a grader, and an out-of-scope sweep
is neither** (`F4`). This iter removes confounds and repairs ~60 sites; it moves no clause.
