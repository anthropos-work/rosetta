# iter-39 — clause-5 fifth pass: the blocker ledger

**37 unique blockers** (38 raw findings; `hiring.md:86` was found independently by two auditors).
Seven auditors, **40 files / 8 674 lines**, every in-scope file read top-to-bottom under a partition that
shares no boundary with iters 33/34/38. Combined verification volume, self-reported: **~805 exact citations**
checked against the platform clones at origin `2adcf71`, the live `demo-1` Postgres, the read-only prod
taxonomy capture, or `docker-compose.yml`/`repos.yml`.

Owner column = the repairer assigned in Phase B.

## The measurement against the pre-registration

Registered in `overview.md` **before any auditor reported**: *"10–16 blockers; 7–11 in the 13 files iter-38
repaired; 3–6 in the 27 it never opened."*

| | files | blockers | per file | predicted |
|---|---|---|---|---|
| repaired by iter-38 | 13 | **~25** | 1.9 | 7–11 |
| never opened by iter-38 | 27 | **~12** | 0.44 | 3–6 |
| **total** | 40 | **37** | 0.93 | **10–16** |

**Refuted on count by 2.3×, and refuted on location in BOTH directions** — both strata came in roughly
double their predicted ceiling. The density ratio (repaired vs untouched) is **~4.4×**, down from iter-38's
7.3× and iter-34's 9×, while the absolute count rose 11 → 17 → **37**.

> **The honest reading, and it is a caution against the whole series:** these five numbers were produced by
> five different instruments. iter-33 ran 5 auditors, iter-38 ran 6, iter-39 ran 7 (6 full-read + 1
> adversarial diff) and briefed them with the accumulated §5 rules and each file's own repair history.
> **25 → 13 → 11 → 17 → 37 is not a corpus trend; it is four instrument changes.** Nothing in this series
> licenses a claim that the residual is converging — or that it is growing.

**Ten blockers sit in seven files no prior pass ever flagged** — `studio-desk.md` (2), `sentinel.md` (2),
`graphql-wundergraph.md` (2), `studio-room.md` (1), `clerkenstein.md` (1), `alignment_testing.md` (1),
`cms.md` (1). Every one of those files had been read in full at least twice and passed twice. This is
iter-38's `ai_architecture.md` finding reproduced at scale: **what changes between passes is the partition,
not the diligence.**

---

## Blockers

### `corpus/architecture/external_services.md` — 7 (owner R1)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 1 | `:395` | the compose `graphql` service built from the **production** `Dockerfile`, composing the **committed** SDL | it always built from **`Dockerfile.dev`** with `context: ..`, regenerating the SDL from `../app` (`show 1e8e754:docker-compose.yml:6-8`; `log -S` returns only `67ba772` + `360efd4`). `:434` in the same file already said `context: ..` — self-contradicting |
| 2 | `:401-410` | quoted `Dockerfile.dev` block shows `COPY cms/…` + `COPY jobsimulation/…` and an `awk … && …` | the archived file has **one** schema COPY (`:18`) and a comment at `:19-20` that cms + skillpath are folded in. The section is fenced as "what a reader will meet in the archived repo" — so a non-matching block defeats its own purpose |
| 3 | `:356` | GraphQL **Subscriptions** over `sse_post`, "served by `backend` now" | `sse_post` → `ws` in Feb 2026 (`bba862f`), then the whole jobsim subgraph deleted at `915da06`. `grep -rn sse *.yaml` → 0 (positive control: 3 files). **No `type Subscription` exists** in the backend SDL |
| 4 | `:508-511` | a strict EU-first **5-step fallback chain** Azure→Bedrock→Mistral→OpenAI→Anthropic | no ordered chain exists. The caller names the vendor; `Azure`→EU, US only behind `flag_use_azure_us`; Anthropic always Bedrock; the **only** automatic fallback is a **429-only** retry that goes to **direct OpenAI (US)**. Mistral is not in the manager at all — OCR only. EU-residency-relevant |
| 5 | `:503` | Bedrock runs **Claude 4.5/4 Sonnet** | no `claude-sonnet-4-5`/`claude-4` string exists. Real: `eu.anthropic.claude-sonnet-4-6`, plus `claude-opus-4-8` for Course Builder |
| 6 | `:202-203` | `gen_injected_override.py:598-599` is the `DIRECTUS_DATA_CONSUMERS` re-point | the re-point is `:636-637`; `:598-599` is an unrelated comment. (The claim itself is true) |
| 7 | `:211-212` | `test_injection.py:1005` is `test_backend_the_actual_reader_is_repointed` | it is `:1051`; `:1005` is a different test |

Plus `:349` (port row, see #34) and `:540` (taxonomy figures, see #35).

### `corpus/services/hiring.md` — 6 (owner R3) — twice-repaired, defective after both

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 8 | `:86` | *"Clerk-only → the browser doesn't re-skin either (the re-skin reads Clerk, not the column)"* | **the parenthetical refutes its own predicate.** Clerk-only means Clerk `publicMetadata.isHiring` IS set — and the re-skin reads Clerk, so it **does** re-skin (`useGetClerkOrganization.tsx:17,20`). It is **DB-only** that fails to re-skin. **Found independently by two auditors.** Second consecutive wrong answer in this slot |
| 9 | `:142-143` | the `OrgFeatureInsights` gate is at `resolver_queries.go:1089` | `:1089` is inside a *different* resolver; the gate is `:1035`. The same doc has it right at `:56` |
| 10 | `:129` | read-path table step 3 cites `:1088,:1134` | `:1088` is blank, `:1134` is unrelated. Correct: `:1034`, `:1080`. **The headline table** |
| 11 | `:204` | the Results relabel is at `useNavbarSections.tsx:300-307` | that is `settingsMenuItem`; the relabel is `:460` (~`:459-466`) — 153 lines off |
| 12 | `:60` | `publicMetadata.isHiring` (`:197-198`) | bare line range, no filename → resolves against a 127-line file. Dead, and redundant with the correct citation 8 lines below |
| 13 | `:165` | an out-of-range enum row "**vanishes at Ent scan**" | Ent's `assignValues` casts **unconditionally** (`jobsimulationsession.go:181-186`) — the row survives the scan. It fails at the **GraphQL enum marshal** |

### `corpus/services/backend.md` — 3 (owner R2)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 14 | `:73` | `schemas/skiller.graphqls` deleted at `graphql-wundergraph@c284453` | `c284453` is a CHANGELOG-only version bump. The deletion is **`749dc86`** (2026-06-24) |
| 15 | `:16`, `:73` | cms-in-app was the **2→1** step; at the skiller merge it was **3** | `915da06^` composes **three** subgraphs and `915da06` deletes **both** cms and jobsimulation SDLs → a **3→1** step (so jobsim-in-app did *not* remove the jobsim subgraph). At the skiller merge it was **4**. `:92` in the same file says 4 and is right — self-contradicting |
| 16 | `:173` | the "four FAMILIES a 9 omits" — then lists **five**, two of which were in the original 9 | 13 − 9 = 4. The four actually added are `recommendations` (M219), `email_overrides` (M408), `notification_logs` + `notification_optouts` (M400/M403). `live_snapshots` and `text_translations` were already there; `recommendations` — the one genuinely missing — is not named. The **13** is correct |

### `corpus/architecture/ai_architecture.md` — 3 (owner R6, #37 shared)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 17 | `:56-66` | the Studio-Room generation-slot table — **wrong on every row** | `production_config.ini:26-36`: FAST/STRICT `gpt-5-mini`, EXECUTION/CREATIVE/REASONING `gpt-5.4`, stable == experimental. `gpt-5.2` is in **no** studio config; `gpt-4o` in no `*_MODEL` slot. `studio-room.md:220-223` quotes it correctly — self-contradicting across files |
| 18 | `:148` | *"default: `gpt-5` via Azure"* | `gpt-5` is nowhere a default. Three different defaults: engine fallback **gpt-4.1 on direct OpenAI** (not Azure); validation pinned **Azure + gpt-4.1**; content-side **gpt-5.1 + openai** |

### `corpus/architecture/service_taxonomy.md` — 2 (owner R6)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 19 | `:203`, `:137` | Ant Academy has **"None at runtime"** platform dependencies and **"No GraphQL"**; "fully independent of the backend" | it reads its catalog **and writes per-user progress** over GraphQL to the platform academy subgraph at runtime. `ant-academy.md:63` **explicitly retired** this framing at v2.5 M231 — the taxonomy page re-asserts the documented root cause of the "empty academy" bug |
| 20 | `:307` | router **Port 5050 — *prod only*** | prod is **8080** (`terraform/locals.tf:8`); `5050` was **only** the deleted local compose host mapping. iter-38 corrected this in three sibling files, one of which now says "never a production port" — this row survived and says the opposite |

### `corpus/services/sentinel.md` — 2 (owner R4)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 21 | `:5`, `:78`, `:82` | live callers are `app` **and `messenger`**; `AUTHORIZATION_ADDRESS` is in **"every other service's"** compose env | `AUTHORIZATION_ADDRESS` appears **exactly 3 times**: backend `:45`, jobsimulation husk `:99`, cms husk `:160`. messenger has no such env, no `depends_on: sentinel`, and **zero** authorization RPC imports. `clerk-integration.md:75` already says messenger has no auth — corpus self-contradiction |
| 22 | `:40` | `g2` roles are `admin`/`member`/**`manager`**/`candidate` | the roles are `admin`/`member`/`candidate`/**`content_creator`** (`enum/membership.go:8-15`). **There is no `manager` role** — it appears only as a test fixture. Granting it yields a membership with **no policy rows at all**: the silent-403 mode this corpus warns about elsewhere |

### `corpus/services/graphql-wundergraph.md` — 2 (owner R7)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 23 | `:91` | *"Before that it built from the production `Dockerfile`"* | same false claim as #1, independently found in a second file by a second auditor. Compose always built `Dockerfile.dev`. `:84` in the same file describes the `Dockerfile.dev` behaviour correctly |
| 24 | `:146-152` | the smoke test promises `{ __typename }` → `{"data":{"__typename":"Query"}}` | it returns **HTTP 200** with `{"errors":[{"message":"unknown viewer: Forbidden"}],"data":null}`. The app pins **that exact query** in its own regression test (`gqlauthz_test.go:176-190`). The documented "healthy" output is the one output the platform guarantees you cannot get |

### `corpus/services/studio-desk.md` — 2 (owner R5)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 25 | `:65` | code map lists `app/designer-sim/` | it does not exist. The three real dirs are `simulation-builder`, `sim-advanced-builder`, `sim-guided-builder` — named correctly by the same doc's prose at `:34`. The map contradicts the prose on the flagship feature |
| 26 | `:271` | *"sync Clerk users to local DB via Tailscale funnel"* | studio-desk has **no database of any kind** (no driver in `package.json`, no DSN in `src/`). The only Tailscale funnel in the repo is the **GlitchTip error-ingest** endpoint. A false capability claim |

### `corpus/services/ai-readiness.md` — 2 (owner R2)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 27 | `:558-560` | `HowWeMeasureTab.tsx:2773-2797` | the file is **1989 lines** — the anchor is ~784 past EOF. The claim (`interviewQuestions` never rendered) is TRUE; the anchor is unfollowable. Real block `:1879-1930` |
| 28 | `:295-296` | `useNavbarSections.tsx:253-260` consumes `AI_READINESS_URL` | that is `librarySkillPathsMenuItem`, ~145 lines off. Real: `:398-400` + gate `:547`. Also `urls.ts:50` is a comment; the constant is `:52` |

### One each — 9

| # | file:anchor | the false claim | what is true | owner |
|---|---|---|---|---|
| 29 | `architecture/README.md:21` | *"the **five** internal Go modules **imported as private dependencies** (…, authn, …)"* | `authn` is **not** imported — zero hits in any `go.mod`/`go.sum` (positive control passes). It ships **inside colony**. Four modules are required. `shared_libraries.md` is correct and fenced on this; **the index drops the caveat and re-asserts the retracted claim** — and an index is where a reader stops | R2 |
| 30 | `cms.md:73,:76` | Go **1.25**; built in `golang:1.25-bookworm` | `go.mod:3` → **1.26.4**; `Dockerfile:2` → `golang:1.26-bookworm`. Building with 1.25 fails outright | R1 |
| 31 | `alignment_testing.md:320` | scripts at `alignment/scripts/{gate,drift-check}.sh` | `alignment/` has **no `scripts/` dir**; they live at `clerkenstein/alignment/scripts/`. Load-bearing because the doc's own "where things live" section draws the harness/mirror line — and these are on the mirror side | R4 |
| 32 | `clerkenstein.md:134-135` | *"dropping the prefix would bounce every hero"* | `STUDIO_ACCESS_ROLES` accepts **both** forms (`['admin','org:admin','content_creator','org:content_creator']`), so an unprefixed `admin` passes. The prefix is a **fidelity** choice, not a gate requirement | R4 |
| 33 | `studio-room.md:181,197,202,328` | `gen.py -t/--template NAME`, with four example commands using it | **there is no `--template` flag** — `gen.py:483-492` registers nine args and none is it. `template` survives only as a legacy *blueprint field*, popped with a deprecation warning, and only for three legacy names. **Worse than a hard failure:** `parse_known_args` **silently absorbs** `--template x`, so every documented command runs and produces something unrelated | R5 |
| 34 | `service_taxonomy.md:307` + `external_services.md:349` | port `5050` presented as production | see #20 — the same retraction missed two rows | R6/R1 |
| 35 | `skiller.md:36`, `external_services.md:540`, `shared_libraries.md:186`, `architecture_overview.md:202,:13` | *"60K skills, 18K roles"* | **"18K roles" is REFUTED from below**: the prod public capture gives `public.job_roles` = **22,470**, and public ⊆ total. 18,919 is the **job-role-embedding** count, mis-transcribed onto the role count. **"60K skills" is NOT refuted — it is UNSUPPORTED**: public skills measure 42,790 and org-scoped rows could close the gap. Two different verdicts; do not collapse them | R7/R1 |
| 36 | `ant-academy.md:233` | public routes enumerated, *"(other `/api/*` stay gated)"* | **false** — `/api/_meta*`, `/api/meta*` and `/api/verify/*` are public — and the list omits the **largest** public surfaces: `/`, `/latest*`, `/chapters/*`, `/courses`, `/courses/*`, plus robots/sitemap/llms/.well-known | R7 |
| 37 | `ai_architecture.md:7` + `security_compliance.md:7` | the per-check verdicts **"come from an LLM"** (unqualified) | **both files' own bodies call that universal "the opposite error"** and state the honest claim as **most**; `EngineTextDiff` checks are deterministic (live: llm 1462 / text_diff 17). The summary lines assert what the detail sections retract — in both files | R6 |

### Mechanical damage from the previous repair pass — 1

| # | file:anchor | what is wrong | owner |
|---|---|---|---|
| 38 | `coursebuilder.md:66-67` | an anchored replacement left the tail of the old sentence dangling behind *"Historically this read as"* — ungrammatical, **and it traps the true, current assertion ("no half-working surface", `main.go:758-766`) inside a historical frame**, so the doc no longer states it as fact | R5 |

*(#38 is the 37th unique blocker; #8 and the duplicate finding of it are counted once.)*

---

## What this ledger does NOT contain

**~60 minors** with exact anchors across the seven auditor reports — line drift, undercounts, omitted list
members, stale "last updated" headers. Clause 5's *"YELLOW with 0 blockers"* admits them. They are routed as
`DOC-M257x-iter39-minors`.

**Nine UNVERIFIABLE classes**, reported honestly rather than guessed: everything internal to the private Go
modules (colony/proto/taxonomy/authn — not cloned, not in the module cache); every GitHub archive-status
claim (`gh` is not installed); production-only infrastructure and legal facts (VPC CIDR, sub-processor count,
DPA version, backup topology); the five alignment scores (runners cannot build — `GOPROXY=off` and colony
absent); `db-backup.md` in near-entirety (no repo on disk, no reference anywhere in `platform/`); the prod
state of the legacy `cms`/`jobsimulation`/`skillpath` schemas (a fresh local stack has only `auth`,
`directus`, `extensions`, `public`, `sentinel`); the 200-person AI-readiness showcase figures (the running
demo is a different, smaller preset); several timing figures the docs themselves flag as re-measure-first;
and `ai-labs`' remote control-plane repo. **None of these is a blocker — an unreachable ground truth is not a
false claim.** They are the honest edge of what this environment can settle, and they are worth writing down
so a future pass does not spend budget re-discovering the boundary.

---

## Phase C — the adversarial pass over iter-39's OWN sweep: **8 self-inflicted**

Fourth consecutive non-clean adversarial pass (24 % → 2 → 6 → **8**). 122 hunks read in full with ±20
lines of file context each; ~95 introduced anchors resolved.

**The defect class has SHIFTED, and this is the pass's most useful output.** Three prior passes were
dominated by *mechanical damage* — half-applied edits, doubled words, orphaned predicates. This one found
**one** mechanical defect and **five cross-file DRIFT** defects: a claim corrected in the file its owner
held, while the identical claim survived in a twin file owned by somebody else, or in no partition at all.

| # | site | class | what happened |
|---|---|---|---|
| C1 | `ai_architecture.md:15-23` | drift | still published the **EU-first 5-step ladder** that `external_services.md` had just refuted — in the corpus's own AI-plane doc, in a file this sweep edited 7 times without touching this section |
| C2 | `external_services.md:405-410` | **over-correction** | *"Three things — and **only three** — can send a request outside the EU"*. There is a fourth: a sequence with **no `ai_vendor` set** defaults to `openai` (`jobsimulation.go:1302`) → **direct US OpenAI**. The universal was derived by reading `jobsimulation/ai/ai.go` alone, which structurally cannot see a default applied in the cms content layer |
| C3 | `external_services.md:356,:434` | dead citation | cited `bba862f` for an `sse_post`→`ws` switch. `merge-base --is-ancestor` → **rc=1**; it lives only on `origin/feat/use-web-socket`. Mainline **never carried `ws`** — `915da06~1` still reads `sse_post` |
| C4 | `graphql-wundergraph.md:41` | drift | still sold the jobsimulation **subscriptions** the same sweep withdrew two files over |
| C5 | seven sites | drift | `backend.md` was corrected to **3 → 1**; `service_taxonomy.md` ×2, `external_services.md`, `graphql-wundergraph.md` ×3 and `architecture_overview.md` kept **2 → 1**. Five of the seven are in files this sweep edited |
| C6 | five sites | drift | the taxonomy verdicts reached 4 canonical sites correctly and **missed** `ai_architecture.md:80,:51`, `service_taxonomy.md:112`, `backend.md:9,:36` — two of which still assert the **refuted** "18K roles" |
| C7 | three sites | drift | `studio-room.md` now warns `--template` does not exist and is **silently swallowed** — while `service_taxonomy.md:186` and `cms.md:210` still tell the reader to run it, and `cms.md:126` carries it in a **mermaid edge label** |
| C8 | `ant-academy.md:65` | drift | the FS-catalog-fallback premise is false (`serverTenant.js:115-145`: *"there is NO FS-as-published fallback … not reversible-on-error"*) and the file **already contradicted it 47 lines later**. A demo grid renders only via the `academy-fs-published-fallback` rext demo-patch |

**The `file:line` asymmetry held for a fourth pass: every introduced anchor that was resolved was correct.**
All eight defects are prose-level.

### What this changes about how a repair pass should be run

A partition that is correct for **reading** is wrong for **repairing**. Six auditors each owning a disjoint
file set is exactly what surfaces independent double-finds (§5 rule 18(b), and it worked again here —
`hiring.md:86` was found twice). But a *claim* does not respect a file partition, and a repairer who owns
`external_services.md` cannot fix the same sentence in `graphql-wundergraph.md`. **Five of eight
self-inflicted defects are the direct, predictable consequence of file-scoped ownership over a
claim-scoped problem.**

> **Rule earned (candidate §5 rule 19): repair by CLAIM, not by FILE.** Before editing, grep the whole
> corpus for the claim being corrected and fix every instance in one pass — or the repair *manufactures* a
> contradiction where there was previously a uniform falsehood. A uniformly-wrong corpus is at least
> self-consistent; a half-repaired one teaches the reader that the corpus disagrees with itself, and the
> next auditor spends its budget adjudicating rather than measuring.

**One-way-door check that passed:** the adversarial pass confirmed neither `ai_architecture.md` nor
`security_compliance.md` now settles the EU-AI-Act classification — both defer explicitly to counsel. That
was the single most dangerous thing this sweep could have done and it did not do it.
