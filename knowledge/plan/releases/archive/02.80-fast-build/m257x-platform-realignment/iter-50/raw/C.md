# Seat C — M257x clause-5 KB-fidelity reading

## 1. Header

**Corpus under audit:** `/Users/marco/workspace/anthropos/rosetta`, branch `m257x/platform-realignment`,
HEAD `57dfbfded8791fcb12a4651d747247ce9d04d7f0` (verified with `git rev-parse HEAD` + `git branch --show-current`).

**Ground-truth clones consulted** (sha verified with `git rev-parse HEAD` in each):

| clone | sha (measured) | matches briefing |
|---|---|---|
| `stack-demo/platform` | `2adcf714bd877a205e8948f59a23db49b884c054` | ✅ |
| `stack-demo/app` | `5ba1704482cf812b130c2d3673afd09f4f7f22e5` | ✅ |
| `stack-demo/app/studio` | `aeec036a51c8a4ae0c5b8f7d5d21cfa7086b658e` | ✅ |
| `stack-demo/cms` | `ca50c8170fefe1122d680efe54f7e56798a79d82` | ✅ |
| `stack-demo/sentinel` | `88bc55929dde7ba43913966ec3fc36372e4ff32a` | ✅ |
| `stack-demo/messenger` | `fa47850d9c507d1928da7a38f7b37bac1bb8fabc` | ✅ |
| `stack-demo/graphql-wundergraph` | `60c229f39adcbbe75c84cd58f0f45052b5423372` | ✅ |
| `stack-demo/next-web-app` | `bb3313bc0133ee5728ce83fda485e95bfea1a6c6` | ✅ |
| `.agentspace/rosetta-extensions` | `a91f8f78095d6725d9deb732c417661036550931` | (tooling monorepo) |

Also consulted (not sha-pinned in briefing but read): `stack-demo/jobsimulation/terraform`,
`stack-demo/roadrunner/terraform`.

### Positive control — `wc -l` on each assigned file

Single invocation: `cd /Users/marco/workspace/anthropos/rosetta && wc -l <7 files>`

| File | measured | briefing | ok |
|---|---|---|---|
| `corpus/architecture/alignment_testing.md` | 521 | 521 | ✅ |
| `corpus/architecture/architecture_overview.md` | 349 | 349 | ✅ |
| `corpus/services/cms.md` | 254 | 254 | ✅ |
| `corpus/services/sentinel.md` | 166 | 166 | ✅ |
| `corpus/services/messenger.md` | 128 | 128 | ✅ |
| `corpus/services/README.md` | 79 | 79 | ✅ |
| `corpus/services/db-backup.md` | 31 | 31 | ✅ |
| **total** | **1528** | 1528 | ✅ |

All seven read IN FULL, top-to-bottom, via `Read` with no offset/limit.

**Search-pipeline controls run** (method rule 3): every `grep` that returned zero was paired with a pattern
known to match in the same directory/pass. Recorded inline below at each zero result:
- `grep -rn 'REDIS_WORKER_INDEX' --include=*.go` in `messenger/` → **0**; control `REDIS_STREAMS_INDEX` → **2 hits** ⇒ pipeline live, absence real.
- `grep -rn 'CMS_RPC_ADDR' --include=*.go` in `app/` → 1 hit (a comment); control `STORAGE_RPC_ADDR` → 3 hits ⇒ pipeline live.
- `grep -rn 'cmsv1connect.NewCMSServiceClient'` in `app/` → **0**; same pass as the control above ⇒ absence real.
- `grep -rn 'ai_vendor\|AiVendor'` in `app/internal` → 2 hits; control `AIVendor` → 123 hits ⇒ pipeline live.
- `grep -c 'manager' sentinel/init_policy.sql` → **0**; control `"member"` across the repo → 22 hits ⇒ absence real.
- `grep -rn 'python-docx\|docx' app/studio/requirements.txt` → **0**; the same file was `cat`-ed in full in the same pass (9 lines shown) ⇒ absence real, not a regex miss.
- `grep -rn 'DIRECTUS_DATA_CONSUMERS'` in `stack-demo/app` → 0 (**wrong directory** — my error); re-run at the rext root → 6 hits. Recorded so the first zero is not mistaken for absence.

---

## 2. BLOCKERS

**None.** No claim in the seven assigned files was measured false in a way a reader would act on, and no
load-bearing `file:line` anchor failed to resolve to substantively the thing the text says is there.

I actively hunted for blockers in the highest-risk shapes the briefing names (retractions, "NB:", "**not**
X but Y", bolded corrections, banner blockquotes). Every one of those passages I could check against
ground truth **verified**, several of them to the digit. The four I most expected to break, and what
they measured:

| Repaired passage | Verdict |
|---|---|
| `architecture_overview.md:3` router-status banner ("two states", "no `:5050` on a local stack") | **TRUE.** `b56d731` + `360efd4` merged as `2adcf71` (2026‑07‑31) delete the router service + `repos.yml` entry; `grep -n '5050' docker-compose.yml` → 0 hits; frontends baked at `http://…:8082/graphql/query` (`docker-compose.yml:352,361`); prod `service_desired_count = 1` at `graphql-wundergraph/terraform/main.tf:20`. |
| `cms.md:36-39` / `architecture_overview.md:243` "**3 → 1**, not 2 to 1; the commit subject is wrong" | **TRUE, and the meta-correction is right.** `git show --stat 915da06` deletes **both** `schemas/cms.graphqls` **and** `schemas/jobsimulation.graphqls` in one commit; the commit's own subject reads *"supergraph 2→1"* — i.e. the corpus is correct and the commit message is the thing that is wrong. Ladder confirmed: `749dc86` removed skiller (5→4), `7c17e63` removed skillpath (its own message: *"composes 3 subgraphs (backend, jobsimulation, cms) instead of 4"*), `915da06` 3→1. `ls schemas/` → `backend.graphqls` alone. |
| `alignment_testing.md:327-343` the twice-corrected CI claim ("*'Inert' ≠ 'absent'*") | **TRUE.** `clerkenstein/.github/workflows/alignment.yml` exists and is git-tracked (`git ls-files` confirms); `ls -la .github` at the monorepo root → *No such file or directory*; `git ls-files '*.github/workflows/*'` returns that one file only. Its lines **10‑11** read verbatim: *"as a subdir workflow under clerkenstein/.github (not at the monorepo root), this is currently inert…kept illustrative of the per-mirror gate shape."* Both the original claim and the first M218 over-correction are correctly retracted. |
| `alignment_testing.md:174-217` the rewritten capability-coverage section | **TRUE in every enforceable particular** — see Audited zeros §4.1. Including the deliberately weakened claim *"It binds the **declared** surface. It cannot discover consumption nobody wrote down."* which is exactly what `dna.go:290-315` implements. |

---

## 3. MINORS

Eight. All are line drift, list-membership slips, or an internal-antecedent wobble — none changes what a
reader would do.

1. **`corpus/architecture/alignment_testing.md:193`** — *line drift.* Text: "`gate.sh:61` calls `alignctl dna
   coverage --dna … --if-declared`". The call is at **`gate.sh:69`** (`grep -n 'dna coverage' gate.sh` →
   `69:"$alignctl" dna coverage --dna "$base/$dna" --if-declared`; file is 87 lines). Line 61 lands inside the
   comment block *describing that same call* ("It now runs on every gate, BEFORE the score"), so a reader
   still finds it, and the **semantics are exactly right**: `dna.go:51` declares
   `--if-declared` as "no consumed_surface ⇒ warn (exit 0); a DECLARED hole still fails", and
   `coverage_test.go:70-82` pins both halves. Substance verified; number 8 lines stale.

2. **`corpus/services/cms.md:110`** — *false list member.* The `studio/` code-map row enumerates
   "openai, anthropic, mistralai, rich, pyyaml, **python-docx**, requests, jinja2, pytest, pytest-asyncio
   (see studio/requirements.txt)". `python-docx` is **not** in the file. Measured by `cat` on both copies:
   `stack-demo/app/studio/requirements.txt` and `stack-demo/cms/studio/requirements.txt` are byte-identical
   and contain exactly nine entries: openai, anthropic, rich, pyyaml, requests, jinja2, mistralai, pytest,
   pytest-asyncio. Ten claimed, nine real, one fabricated.

3. **`corpus/services/README.md:20`** — *internal antecedent inconsistency.* ":11" defines "**Four** services
   in this index — skiller, skillpath, jobsimulation and cms" and ":15" explicitly calls roadrunner "**the
   fifth**". Then ":20" says "**three of the four** (cms, jobsimulation, roadrunner)". Roadrunner is not one
   of "the four"; the set named is three of the **five**. The enumerated list itself is correct and verified
   (`docker-compose.yml:144`, `:83`, `:281`, all with `graphql` in `profiles:` at `:187`, `:140`, `:309`).

4. **`corpus/services/sentinel.md:44`** — *omitted list member.* "only as a fixture string in sentinel's own
   tests (`internal/authorization/casbin_test.go`, `internal/rpcsrv/rpc_test.go`)". It also appears at
   **`internal/authorization/manager_test.go:21`** (`manager := "manager"`). The claim's substance —
   *only* a test fixture, never a policy — is TRUE (`grep -c 'manager' init_policy.sql` → **0**); the
   parenthetical enumeration is one file short. (The only other hit, `manager.go:20`
   `logger.With("service","manager")`, is a logging label, correctly not a role.)

5. **`corpus/architecture/alignment_testing.md:512-519`** — *stale Layout block, two omissions.* The tree
   omits **`internal/canon`** (present: `ls internal/` → `canon compare dna outcome report`), and its
   `cmd/alignctl` gloss reads "run | capture | dna list|diff|validate", omitting **`coverage`** — which the
   same document documents at `:245` and `dna.go:26` implements. Self-inconsistent within one file.

6. **`corpus/architecture/alignment_testing.md:325-326`** — *mislabeled link.* Link **text** is
   `` `knowledge/alignment.md` `` but the **href** is `../services/clerkenstein.md`. The named file does
   exist (`clerkenstein/knowledge/alignment.md` ✓), so the text is truthful — but the link does not go
   there. (All seven files' relative links otherwise resolve: a full link-existence sweep over all `.md`
   targets in the seven files reported zero broken paths.)

7. **`corpus/services/messenger.md:110`** — *ambiguous scope.* Inside messenger's env-var table:
   "**Gone from docker-compose** … only the residual `SKILLPATH_STREAM=skillpath` remains." `SKILLPATH_RPC_ADDR`
   is indeed gone (`grep -n 'SKILLPATH' docker-compose.yml` → one hit, `64: - SKILLPATH_STREAM=skillpath`),
   but that surviving line is in the **`backend`** block (28‑81), **not** messenger's (240‑280). Read as
   "remains in docker-compose" it is true; read as "remains in messenger's env" — which the table context
   invites — it is false. Worth one clarifying clause.

8. **`corpus/architecture/alignment_testing.md:190,201`** — *paraphrased identifier.* The doc writes the
   endpoint as `GET /v1/users/{id}`; the DNA field reads `GET /v1/users/{userID}`. Cosmetic; the consumer
   string in the DNA matches the doc's gloss verbatim ("next-web @clerk/nextjs currentUser() — SSR, every
   authenticated render").

---

## 4. Audited zeros — read in full, measured clean

### 4.1 `corpus/architecture/alignment_testing.md` (521 ll.) — the most heavily verified file

Every mechanical claim I could execute, executed. **The worked example reproduces byte-for-byte.**

Ran (from `.agentspace/rosetta-extensions/alignment`):
`GOFLAGS=-mod=mod GOPROXY=off go run ./cmd/alignctl run --dna examples/toy/dna.json --runner "go run ./examples/toy/cmd/toyrun" --golden-dir examples/toy/golden`

Output matched `:296-306` **line for line**, including `Score: overall 86.7%   critical 100.0%   (5/6 genes
aligned)`, `Add 3/3 ok`, `Greet 2/3 DIVERGED`, and `FAIL Greet/padded-name  (exact, w2)` with the exact
source/mirror strings. `go test -tags alignment ./examples/toy/...` → `ok`; toy gate constants confirmed
`gateOverall = 80.0`, `gateCritical = 100.0` (`examples/toy/alignment_test.go:21-22`) — exactly as `:308`.

Structural verification, all confirmed:

| Claim (line) | Ground truth |
|---|---|
| Weights `3/2/1` by criticality (`:97`, `:153`) | `internal/dna/dna.go:38-45` |
| overall = Σ(w·aligned)/Σ(w); critical % is an **unweighted count ratio** (`:154-157`) | `compare.go:101-103` — `pct(alignedW,sumW)` vs `pct(critAligned,critTotal)` |
| id regex `^[A-Za-z0-9][A-Za-z0-9_-]*$` (`:104`) | `dna.go:21` — character-identical |
| explicit weight must be `1..1000000` (`:106`) | `dna.go:24,269-270` (`maxWeight = 1_000_000`) |
| `normalized` genes must list `normalize` paths (`:105`) | `dna.go:266` |
| Four operators exact/shape/normalized/error_class (`:113-118`) | `dna.go:55-58` — glosses match the source comments |
| Zero-critical guard is **two-layer**: `dna.Validate` rejects (authoritative) + `GateMet` refuses (scoring-time backstop) (`:161-167`) | `dna.go:277-282` and `compare.go:48-61`. The source comments use the *same words* ("LOAD/lint-time half"/"SCORING-TIME half", "authoritative"). |
| `critical_genes` carried in the report (`:165`) | `compare.go:43` `json:"critical_genes"` |
| `alignctl run` calls `Validate` **before it scores anything** (`:191`) | `run.go:38` `d.Validate()` precedes `run.go:70` `compare.Evaluate` |
| coverage exit codes: 0 covered / 2 uncovered **or no surface declared** (`:192`) | `dna.go:34-40` comment + `cmd/alignctl/dna.go` |
| `--if-declared` downgrades **exactly one** case (`:193`) | `dna.go:51` flag text + `coverage_test.go:70` (warn) / `:81` (still fails a declared hole) |
| `consumed_surface` = `{endpoint, consumer, capability \| covered_by}`; Validate rejects an endpoint naming no/unknown capability; `covered_by` **not** machine-verified (`:190,194`) | `dna.go:290-315` + `UncoveredEndpoints` at `:150-159` (`covered_by` short-circuits with `continue`) |
| **Only `clerk-2.6.0` declares a `consumed_surface`; the other four declare none** (`:201-204`) | Measured over all five DNAs: `clerk-2.6.0` → 15 entries; `clerk-deploy-1`/`clerk-express-1`/`clerk-js-5`/`clerk-multi-1` → **NONE**. Exactly as claimed. |
| `covered_by` example `clerk-express-1:ClerkClientBAPI/get-organization` (`:194`) | Present **verbatim** in `clerk-2.6.0.json` |
| `GET /v1/users/{userID}` → capability `GetUser`, consumer "next-web @clerk/nextjs currentUser() — SSR, every authenticated render" (`:182,190`) | Verbatim in the DNA |
| Exit **2 = REGRESSED** vs **3 = UNMEASURABLE**, "2 and 3 were the SAME code until M219" (`:237-239, :280-284`) | `run.go:121-135` const block + the loud banner at `:140-153`; `run_test.go:36-39` pins they may never be equal |
| All five DNA gene/capability counts (`:254-258`) | Measured: `clerk-2.6.0` 14 caps/27 genes · `clerk-js-5` 6/9 · `clerk-multi-1` 5/9 · `clerk-deploy-1` 3/7 · `clerk-express-1` 5/13. **All five exact.** |
| Five runners `clerkrun`/`jsfapirun`/`expressrun`/`deployrun`/`multirun` (`:375`) + five golden dirs | `ls alignment/cmd/` and `ls alignment/` — exact, no extras |
| `MembershipOrgIdentity/real-org-eid` is a **`standard`** gene and **stays in the DNA** (`:262-274`) | Present, `criticality: standard` |
| **"was 97.2% / 26-of-27"** (`:254, :262`) | **Arithmetic reproduced.** Σweights over the DNA = **71**; the org-eid gene's weight = **2**; (71−2)/71 = **97.18% → 97.2%**. 17 critical genes, none of them that one ⇒ critical stayed 100%. The number is a real computation, not a remembered figure. |
| `Store.SeedOrgIdentity`/`LookupOrgEid` wired from `cmd/fake-bapi` (`:269-271`) | `clerk-backend/store.go:138,151`; `cmd/fake-bapi/main.go:88` |
| The three `universal-user` genes are `ExtractIdentity`, `Me`, `DeployIdentity` (`:212-213`) | Measured across all five DNAs: exactly those three, in `clerk-express-1`, `clerk-js-5`, `clerk-deploy-1`. No fourth. |
| Two-sided `GetUser` fix — asks for hero **A** *and* hero **B** (`:216`) | `GetUser` variants are `hero-a`, `hero-b`; capability is `critical` |
| the reusable `alignment/` section **has no `scripts/`** and holds exactly `cmd/`, `internal/`, `examples/`, `go.mod`, `README.md` (`:321-322`) | `ls alignment/` — **exact set, no extras** |
| `gate.sh` defaults `RUNNER_PKG=./cmd/clerkrun`, `DNA=dna/clerk-2.6.0.json`, `ALIGN_DIR` default `../../alignment` (`:322-324`) | `gate.sh:30,33,35` — all three exact |
| gate `≥95 / =100` (`:254, :314`) | `gate.sh:31-32` `GATE_OVERALL:-95` / `GATE_CRITICAL:-100` |
| `clerkenstein/knowledge/{alignment,architecture}.md` exist (`:324,348`) | Both present |
| golden path `golden/<Capability>/<variant>.json` (`:146`) | `internal/outcome/outcome.go:41-43` |
| `--source live` re-runs the source (`:148-149`) | `run.go:52-56` |
| Stdlib-only, module `anthropos.dev/alignment` (`:521`) | `go.mod` has module + go/toolchain lines and **no `require` block** |

**Unmeasurable (see §5):** the five live *scores* — every gate build needs the private `colony` module.

### 4.2 `corpus/architecture/architecture_overview.md` (349 ll.)

Clean. Notable verifications:

- **The multi-tenancy split at `:298-303` is exact, including its own parenthetical correction.** Measured
  with an `awk` pass over `app/internal/data/ent/schema/*.go` counting struct types embedding `ent.Schema`
  (both the inline `struct{ ent.Schema }` and multi-line forms — the naive one-line grep undercounts to 112
  and would have been the cheap false answer):
  - total ent schemas = **135** ✅
  - `OrganizationMixin{}` users = **30**, each exactly once; `+ Membership` which declares its own
    `Policy()` at `org_membership.go:172` ⇒ **31 of 135** ✅
  - `OrganizationMixin` is the only org-scoping mixin carrying a `Policy()` (`mixin.go:126`); the only other
    policy-bearing mixin is `UserMixin` (`mixin.go:99`)
  - schemas with an `organization_id` and **no** policy: 17 direct-field files **minus** `academy_feedback`
    (which does carry a policy, via `UserMixin{}`) = **16** — precisely the doc's *"neither-mixin subset"* —
    plus the **7** using the policy-less `OrganizationIDMixin` (`skiller_mixins.go:152-166`: category,
    jobrole, similarity, skill, specialization, studio_document, studio_task) = **23** ✅.
    16 + 7 = 23, exactly as written, and the warning "16 is the *neither-mixin* subset of those 23, **not the
    total**" is exactly the trap I would otherwise have fallen into.
- `:15,60,189` jobsimulation husk at `docker-compose.yml:83`, `graphql` profile ✅ · `:60,190` cms at `:144` ✅
  · `:60,191` roadrunner at `:281` ✅ (profiles at `:140/:187/:309`).
- `:60` "a bare `make up` gives you **six Go services**, not three" ✅ — sentinel (no `profiles:` ⇒ always on),
  backend, jobsimulation, cms, storage, roadrunner. Gotenberg is separately labelled third-party.
- `:19` `app/internal/jobsimwiring/wiring.go:118` ⇒ `jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"),
  getenv("JUDGE0_BASE_URL"))` ✅ **exact line**.
- `:48` `next: ^16.2.7` across all four apps ✅ — web/hiring/integration/maintenance all `^16.2.7`.
  (`configs/i18n` is at `^15.5.16`, but that is a config package, not one of the four apps.)
- `:244-262` the AI-routing paragraph — **the most-repaired passage in the file, and it verifies end to end:**
  - `getClient` defaults to `azureClientEu`, swaps to `azureClientUs` on PostHog `flag_use_azure_us`:
    `ai.go:262-276` ✅ **exact range** (`:263` `case Azure:` → `:264` eu → `:267` the flag string → `:275` us).
  - `isThrottlingError` at `:129` ✅ exact; applied at `:166` ✅ and `:325` ✅ — both exact.
  - direct OpenAI **is** the 429 retry target: `ChatCompletion` sets `vendor = Openai` on `throttled`
    (`ai.go:152-155`) and `getClient(Openai)` returns `a.openaiClient`, built at `:80` from
    `openai.NewOpenAI(openaiKey)` ✅.
  - Bedrock pinned `eu-west-1` at `:85-88` ✅ exact (`config.WithRegion("eu-west-1")` is line 88).
  - **"an authored sequence with `ai_vendor` unset … reaching direct US OpenAI unconditionally, on the first
    attempt, with no flag and no 429"** — I chased this because it is the strongest claim on the page.
    **TRUE**: `app/internal/jobsimulation/simulator/ai/ai.go:51` `GetAIVendorAndModel` switches on
    `sequence.AIVendor` and its **`default:` branch (`:114`) sets `aiVendor = internalAi.Openai`**. So an
    unset vendor resolves to direct OpenAI before any flag or retry. (Note the manager's *own*
    `getClient` default is an error — the doc is right that this path is **outside the manager**.)
  - Mistral: `internal/cms/studio/markdownManager.go:11` (`ai/mistral` import) and `:19`
    (`mistral.NewMistral(nil, os.Getenv("MISTRAL_API_KEY"))`) ✅ both exact; called from
    `studioManager.go:583` ✅ exact.
- `:279-286` local request flow / no `:5050` ✅ · `:270` prod router `terraform/main.tf:20 = 1` ✅ exact.
- `:327-333` schema separation ✅ — `repos.yml` gives `schema:` to `app` alone (`public`), `sentinel`
  `migrations: false`; `extensions` schema really does host pgvector (`app/terraform/migrations/…:
  "embedding" extensions.vector(1536)`).
- The mermaid diagram (`:73-136`) is consistent with all of the above, including the `%%` comment correctly
  stating the 3→1 step and that "the jobsimulation subgraph outlived jobsim-in-app". Postgres/Redis in the
  diagram are real services — they arrive via the `include:` at `docker-compose.yml:1`.

### 4.3 `corpus/services/cms.md` (254 ll.)

Clean apart from MINOR #2. Verified:
`cms/terraform/main.tf:39 = 0` ✅ exact · `docker-compose.yml:144` + `repos.yml:14-16`
(`migrations: false # legacy — folded into app (cms-in-app v8.0)`) ✅ exact ·
`CMS_RPC_ADDR=http://cms:8091` at `docker-compose.yml:256` ✅ exact ·
`app/main.go:1196-1202` quote *"Additive + DORMANT: external callers (messenger) keep hitting the standalone
cms via CMS_RPC_ADDR until the M809 re-point"* ✅ **verbatim, in range** ·
**"`app` itself makes no outbound cms RPC"** ✅ (only a comment mentions `CMS_RPC_ADDR`; zero
`NewCMSServiceClient`, with a live control) ·
`20260724132049_cms_data_model.sql` creates **exactly** the six named tables in `public` ✅ ·
`app/internal/skillpath/session.go:205-207` — `:205-206` is the `// cms-in-app deseam: cms is in-process`
comment and `:207` is `u.cms.GetSkillPathDomain(...)`, with `u.cms` typed `contentread.CmsContentReader`
(`session.go:50`) ✅ **exact to the line** ·
`REDIS_CMS_CACHE_INDEX` default 5 (`app/main.go:988-989`) ✅ ·
`cms/go.mod:3 go 1.26.4` ✅ · `cms/Dockerfile:2` golang:1.26-bookworm, `:23` python:3.11-slim ✅ both exact ·
ports 8090/8091 ✅ · the whole `internal/` key-directory map ✅ (all 12 named dirs present, plus `ent/`,
`studio/`, `terraform/`, `cmd/`) · `simulation.graphqls`/`skills.graphqls`/`studio.graphqls` ✅ ·
six Directus collection names (`skill_paths`, `simulations`, `sequences`, `library_categories`,
`library_macro_categories`, `resource`) ✅ all six string-literals confirmed ·
webhook moved to `POST /api/webhook/directus` (`app/internal/web/backend/backend.go:324`) and **fails
closed**: `validWebhookSecret` returns `false` on an empty secret (`router.go:21-23`) → 401 before dispatch
(`:37-41`) — while the standalone was unauthenticated (`cms/cmd/root.go:242`, handler takes no secret) ✅
**both halves of the comparison verified** ·
`jobsimulation` has **no `DIRECTUS_BASE_ADDR` of its own` ✅ (the only two occurrences, `:164-165`, are in the
cms block 144‑187) · `app/cms_reader_switch.go` exists ✅ and `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")`
really does name both (`rosetta-extensions/stack-injection/gen_injected_override.py:53`) ✅ ·
`gen.py:484-492` registers **exactly nine** arguments, and they are exactly the nine named ✅ **exact range**;
`gen.py:18-28` is `parse_argument` using `parse_known_args` and merging leftovers ✅ **exact range**; there is
**no `--template` flag** ✅ (only a legacy blueprint *field*, `gen.py:213-233`) ·
venv auto-provision at `studio/studio-venv` invoked from the repo root ✅
(`cms/internal/studio/studioManager.go:98,951`).

### 4.4 `corpus/services/sentinel.md` (166 ll.)

Clean apart from MINOR #4. Verified:
`go.mod:3 go 1.26.0` ✅ exact · `Dockerfile:2` / `Dockerfile.dev:2` golang:1.26-bookworm ✅ both exact ·
Casbin **v3** (`v3.10.0`) ✅ · single table `casbin_rules`, **no Ent** (no `ent/` dir) ✅ ·
`terraform/locals.tf:4-5` → `service_cpu = 256`, `service_memory = 128` ✅ **exact lines** ·
binary default `PORT` 8080 (`cmd/root.go:47`) vs compose 8087 ✅ · no `profiles:` ⇒ always on ✅ ·
**the Casbin model shape is exactly right**: `casbin.go:14-44` declares **6** request types (r,r2..r6),
**6** policy types (p,p2..p6), **3** role groupings (g,g2,g3), **6** matchers (m,m2..m6) ✅ — and each of
the six matcher glosses in the table matches the actual expression, including `m`'s `TIER_FREE`
substitution from `Tier_TIER_FREE.String()` and `m6`'s "no tier logic" ✅ ·
four `MembershipRole` values at `app/internal/data/ent/enum/membership.go:8-15` ✅ **exact range** ·
`init_policy.sql:88-118` is precisely the `content_creator` block ✅ **exact range**, with
`casbin_content_creator_test.go` present ✅ · `init_policy.sql:63-66` is precisely the
"`org:feature:taxonomy:write` is NOT seeded as a default" NOTE ✅ **exact range** ·
**"There is no `manager` role"** ✅ — 0 occurrences in `init_policy.sql` (with a live control) ·
`AUTHORIZATION_ADDRESS` in **exactly three** blocks, `:45` backend / `:99` jobsimulation / `:160` cms ✅
**all three exact**, and `depends_on: sentinel` appears at `:75`/`:130`/`:183` — the same three blocks,
nowhere else ✅ · **`messenger` is not a caller** ✅ — no `AUTHORIZATION_ADDRESS`, no `depends_on: sentinel`,
and **zero** authorization-client imports in its Go source (the single `sentinel` hit is the English word
in a comment at `pkg/aireadinessemail/override.go:378`) ·
`make initdb` really does ignore `DB_CONNECTION` and hard-code
`postgresql://postgres@localhost:5432/postgres?sslmode=disable` ✅ (`Makefile:3`), and `init_policy.sql`
really is schema-qualified (`sentinel.casbin_rules`, 3×) ✅ ·
`local_superadmin_grants.sql` exists, grants p3 `org:feature:taxonomy:write` to `admin`, and carries the
LOCAL-ONLY warning ✅ · all 19 RPC methods in the Interface-Discovery table exist in
`internal/rpcsrv/rpc.go` ✅.

### 4.5 `corpus/services/messenger.md` (128 ll.)

Clean apart from MINOR #7. Verified:
`go 1.25.0` ✅ · `getbrevo/brevo-go v1.1.3` ✅ · `osteele/liquid v1.8.1` ✅ — all three exact versions ·
ports 8200/8201 and `messenger` profile only ✅ (`docker-compose.yml:250-251`, `:279`) ·
`internal/rpcsrv/rpcsrv.go:25-30` is exactly the two `CodeUnimplemented` stubs ✅ **exact range** ·
`internal/flow/jobsimulations.go:140-151` is exactly the >2h / >12h staleness guards ✅ **exact range**, and
the thresholds in prose match the code ·
`cmd/root.go:63` PORT→8080, `:64` RPC_PORT→8081, `:107` REDIS_STREAMS_INDEX→2 ✅ **all three exact** ·
`cmd/root.go:147` `READONLY_DB_CONNECTION` ✅ **exact line**, with the source comment "Copilot read-only DB
connection" at `:146` ·
**the whitelabel/v0.34.0 paragraph is fully substantiated by git**: `whitelabel.go` was introduced by
`d2d41e1` ("integrate Copilot DB … whitelabel support"), which `git tag --contains` places first in
**v0.34.0** ✅; `c0feaa2` "update whitelabel invitation rendering to **return subject and body separately**"
✅ matches the doc's phrasing; and `cda2b7f` "**rename Copilot DB connection to Readonly DB**" ✅ confirms the
"formerly `COPILOT_DB_CONNECTION`" lineage ·
`REDIS_WORKER_INDEX` is set in compose (`:261`) and **not read anywhere in the Go source** ✅ (with control) ·
`internal/flow/flow.go:72-87` is the `backend` subscriber block carrying the five `OrgSkillPath*` handlers ✅ ·
`internal/flow/assignments.go:828` is exactly `h.cms.GetSkillPath(...)` inside `getSkillPath` ✅ **exact line** ·
RPC clients are exactly CMS / BACKEND_USERS (×2, users + orgs) / SKILLER / JOBSIMULATION
(`cmd/root.go:120,125,130,135,140`) with **no skillpath client** ✅ ·
`depends_on` = backend + cms + jobsimulation ✅ (`:274,276,278`) · `SKILLER_RPC_ADDR=http://backend:8083`
✅ (`:265`) and `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401` at `:258` ✅ **exact** ·
`BREVO_KEY` empty ⇒ console sender ✅ (`cmd/root.go:74-75`) · the three named test files exist ✅.

### 4.6 `corpus/services/README.md` (79 ll.)

Clean apart from MINOR #3. Verified: `roadrunner/terraform/main.tf:19 = 1` ✅ **exact** ·
`cms/terraform/main.tf:39 = 0` ✅ **exact** · `jobsimulation/terraform/main.tf:40 = 0` ✅ **exact** —
all three anchors land on the digit. The "one row where prod and `repos.yml` contradict each other"
holds: `repos.yml:29-31` marks roadrunner `migrations: false # legacy — folded into app` while prod still
runs one task. Cms/jobsimulation/roadrunner all still start locally in `graphql` ✅. All 27 doc links in
the index resolve ✅.

### 4.7 `corpus/services/db-backup.md` (31 ll.)

Read in full. Nothing refutable and nothing refuted. The one claim I *could* check — "not in local
docker-compose" — holds: `db-backup` appears in no compose service block and in no `repos.yml` entry ✅.
Everything else (Go, 6-hourly, S3/Azure/Hetzner, RDS PITR) needs the un-cloned `db-backup` repo — see §5.
The doc is short and makes no `file:line` claims, so it carries no anchor risk.

---

## 5. Unverified — and why

Stated explicitly so no skip reads as a pass (method rule 6). Per the briefing these are **neither passes
nor blockers**.

1. **The five live alignment scores** (`alignment_testing.md:252-258`). Every gate build fails at
   `go: downloading github.com/anthropos-work/colony v0.34.3 … module lookup disabled by GOPROXY=off`.
   `colony` is not cloned. I ran `gate.sh` for the Go SDK, JS/FAPI and multi-identity surfaces — all three
   die identically at the runner build, so **no surface's live score was measured**. What I *did* verify is
   everything structural behind those numbers: gene/capability counts (all five exact), the criticality of
   the org-eid gene, and the 97.2% arithmetic (which reproduces from the DNA weights alone). Note this is
   itself an instance of the doc's own `UNMEASURABLE ≠ pass` rule.
2. **GitHub archive status** — `graphql-wundergraph` "ARCHIVED on GitHub (2026-07-30)"
   (`architecture_overview.md:3`, `services/README.md:43`), `jobsimulation` "repo is ARCHIVED (2026-07-31)"
   (`:15`), and `chronos` "**NOT** archived" (`services/README.md:70`). `gh` is unavailable; archive state
   is not derivable from a clone. I confirmed the *local* facts that accompany them (last commits, dates,
   desired_counts), but not the archive bit.
3. **`infrastructure/terraform/production/services.tf`** — `module.cms_euwest1` / `module.jobsimulation_euwest1`
   "still declared as the rollback path" (`cms.md:48-50`, and the same shape in `architecture_overview.md`).
   The `infrastructure` repo is not among the `stack-demo/` clones (`ls stack-demo` confirms). The per-repo
   `terraform/main.tf` desired-counts, which are the load-bearing half, **were** verified.
4. **Taxonomy magnitudes** — "≥42,790 skills / ≥22,470 job roles" (`architecture_overview.md:13,204`).
   Explicitly delegated by the text to `shared_libraries.md#taxonomy-figures` (another seat's file), and the
   underlying measurement needs a live DB, which I did not query.
5. **`db-backup` internals** — schedule, three storage targets, Go implementation
   (`db-backup.md:5-23`). Repo not cloned; service is production-only.
6. **Infra/ops assertions with no local artifact** — ECS 30-second health checks with automated rollback,
   VPC `10.0.0.0/16`, subnet placement, self-hosted EU runners, Better Stack/PostHog/Sentry monitoring
   (`architecture_overview.md:340-346`). All live in the un-cloned `infrastructure` repo.
7. **Sentinel crash-loop message** — "`pq: no schema has been selected`" without the `sentinel` schema
   (`sentinel.md:97`). A runtime behaviour; I did not stand a stack up to reproduce it. The static half
   (schema-qualified seed, `search_path=sentinel` in the DSN at `docker-compose.yml:18`) checks out.
8. **Shared-library internals** — `colony`, `proto`, `taxonomy` claims in
   `architecture_overview.md:200-204`. Not cloned, per the briefing.

---

## 6. Method notes

- All seven files read in full via `Read`, no offset/limit, before any grep — no narrowing to "high-risk
  sections", per method rule 1.
- Every quoted anchor was read **with surrounding context** (method rule 5), not grepped-to-and-quoted. This
  mattered twice: `session.go:205` looked off-by-one under a mis-counted `sed` window until re-derived, and
  the "16 neither-mixin" figure only resolves once you notice `academy_feedback` carries a `UserMixin`
  policy — reading the mixin list, not just the field.
- The `ent.Schema` count is the clearest case of method rule 4: the naive `grep 'ent.Schema$'` returns
  **112** because 23 schemas use the inline `struct{ ent.Schema }` form. The doc's **135** is right and the
  cheap grep would have produced a confident false blocker.
- No probe was allowed to satisfy itself: the toy alignment run compares tool output against corpus text
  neither of which I authored, and the 97.2% figure was recomputed from the DNA weights independently of
  any score the doc reports.
