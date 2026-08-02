# iter-41 — clause-5 SIXTH pass: the blocker ledger

**18 unique in-scope blockers** (21 raw findings; 3 duplicates — `G1≡E1`, `G3≡A3`, `G5≡A1`, each counted once).
Plus **1 out-of-scope** (`G6`, four sites in `corpus/ops/demo/**`) — routed, not counted against clause 5.

Seven auditors, **40 files / 9,163 lines**, every in-scope file read top-to-bottom with a `wc -l` positive
control per file (all 40 confirmed line-for-line). Combined self-reported verification volume: **~800 exact
citations** plus G's 91 diff hunks and ~110 introduced anchors.

**Every blocker was re-derived by the iteration itself before acceptance** (see `adjudication.md`) —
**21 of 21 held**. Unlike iter-22, no handed correction was refuted on re-derivation.

---

## THE HEADLINE: the first like-for-like measurement in the series

| pass | iter | auditors | corpus read | blockers |
|---|---|---|---|---|
| 1 | 21/33 | grep-scoped → 5 | pre-repair | 25 |
| 2 | 34 | ~5 | post-33 repair | 13 |
| 3 | 38 | 6 | post-34 repair | 11 → 17 |
| 4 | 39 | **7** | post-38 repair | **37** |
| **5** | **41** | **7 — IDENTICAL instrument** | post-39 repair | **18** |

iter-39 established that `25 → 13 → 11 → 17 → 37` **measured five different instruments**, not the corpus.
**This pass is the first that changed only ONE variable.** The instrument was held fixed on every knob
(7 auditors, same briefing, same partition *method*, all 40 read in full); iter-40's repair touched
`corpus/ops/**`, `.claude/**` and `CLAUDE.md` but **not one in-scope file**, verified by an empty
`git diff b925199..HEAD -- corpus/services/ corpus/architecture/`.

**So `37 → 18` is a real measurement of iter-39's repair: it roughly HALVED the residual. It did not
approach zero.**

## The pre-registered prediction HELD — for the first time in this series

Registered in `overview.md` before any auditor reported:

| | predicted | actual | verdict |
|---|---|---|---|
| count | **8–20** | **18** | **HELD** |
| location — untouched 20 files | **< 5** | **3** | **HELD** |
| named #1: ≥1 blocker in text written to *explain* a correction | — | 4 (F1, E3, A1, A4) | **HIT** |
| consequent: clause 5 does NOT close | — | 18 blockers | **HELD** |

Four consecutive passes refuted their own predictions on count and/or location. **Holding the instrument
fixed is what made the measurement predictable** — which is itself the strongest evidence that the prior
series was measuring instruments.

## Location: 15 repaired / 3 untouched

**15 of 18 (83%)** sit in the 20 files iter-39 edited; **3** in the 20 it never opened
(**0.75 vs 0.15 per file — a 5× density ratio**, continuing the decay 9× → 7.3× → 4.4× → 5×).

---

## Blockers

### `corpus/architecture/ai_architecture.md` — 4 (A1/G5, A3/G3, G2, + shares G4)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 1 | `:104-105` | *"`gpt-4o` appears in **no `*_MODEL` slot** of any studio config"* | It is in two: `config_template.ini:39` `EXECUTION_AI_STABLE_MODEL = azure, gpt-4o, none` and `:40` `CREATIVE_AI_STABLE_MODEL`. **Self-contradicting inside the same blockquote** — `:102-103` says the template *"still carries … `gpt-4o`"*. Intended claim was "no *shipping* config". Found by **two** auditors |
| 2 | `:38-45` | *"Within the manager **exactly three** things can send a request outside the EU"*, then *"a **fourth** exit: an **unrecognised** `ai_vendor` string"* | Three defects in one enumeration: `ANTHROPIC_API_KEY` is **not** within the manager (it flips Course Builder / Studio-Room, which never touch `AIManager`); a caller-supplied `vendor = Openai` is uncounted; and the fourth exit is the **nil/unset** vendor, not an unrecognised one. Found by **two** auditors |
| 3 | `:59` | OpenAI row: *"direct US OpenAI **only on a 429 retry**"* | Also reached **on the first attempt, with no error condition**, when a sequence leaves `ai_vendor` unset: `jobsimulation.go:905` (`AIVendor *AIVendor`, nullable) → `:1302` `aiVendor := simulation.Openai` → `simulator/ai/ai.go:58-59` → `ai/ai.go:279` → `a.openaiClient`. **The sweep's own twin says so** (`external_services.md:488`) |

### `corpus/architecture/security_compliance.md` — 4 (E1/G1, E2, E3, + G4)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 4 | `:175-176` | *"**'Anthropic Direct' is not used at all**"* — in the **EU Data Residency** section | `coursebuilder/bedrock.go:108-112` routes **every** coursebuilder model call to `api.anthropic.com` whenever `ANTHROPIC_API_KEY` is set; `:43-45` says so in words; `ModelBackendName()` → `"anthropic-api"`. The same sweep **added an "Anthropic Direct" provider row** to `external_services.md:489` and `coursebuilder.md:48` calls it *"the shipped path"*. Found by **two** auditors |
| 5 | `:76`, `:83-84` | *"**16** carry an `organization_id` with **no policy of any kind**"*; *"the remainder … carry no org column by design"* | **Undercount.** 7 further schemas use `OrganizationIDMixin{}` (`category, jobrole, similarity, skill, specialization, studio_document, studio_task`); that mixin declares **0** `Policy()` and none of the 7 declares its own. **Re-measured by this iteration: only FOUR files in the entire schema dir declare any `Policy()`.** The doc **contradicts itself seven lines earlier** — `:69` already names `OrganizationIDMixin{}` as *"a plain nullable organization_id column with no policy"*, then excludes those 7 from the `:76` count. **Errs toward "isolation is handled" — the dangerous direction, and the FIFTH consecutive failure of this fence.** **DOUBLE-FIND: auditor C reached the same conclusion from `architecture_overview.md:288-291`, a different file.** Base count unsettled between auditors (C: 17+7=24, E: 16+7=23) — deliberately **not** settled here |
| 6 | `:205` | *"This classification means transparency obligations only, not the strict requirements of High Risk systems"* | An **orphaned trailing bullet**: the retraction blockquote was spliced into the middle of a bullet list, and the list resumes afterward drawing the operative **legal consequence** from the classification — three lines after `:202` says *"**Do not cite this section as evidence of a Limited-Risk classification**"*. Adjudicated precisely: `:7` **does** defer properly to counsel, so the corpus does not assert wholesale; the defect is this one spliced-orphan bullet. **This CORRECTS `D-M257x-39-4`**, which recorded that the one-way-door check passed and that the file "asserts no legal conclusion" |

### `corpus/architecture/service_taxonomy.md` — 3 (F1, F2, F3)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 7 | `:288-289` | *"a `directus/directus:10.10.1` compose service on port 8055 with an `admin@example.com` / `password` login — that service **has never existed** in the platform compose"* | **It did exist, with exactly that image tag, port and password**, until `a2a3ee6` (2026-02-27): `git show a2a3ee6^:docker-compose.yml` → `:384 image: directus/directus:10.10.1`, `:386 8055:8055`, `:409 ADMIN_PASSWORD=password`. **A FALSE RETRACTION that contradicts the corpus's own fenced SoT which the same section links** (`platform-migration-status.md:86`). Only the `admin@example.com` *email* is unfound in history. Over-correction: *"does not exist now"* → *"never existed"* |
| 8 | `:145` | Studio-Desk **Technology**: *"TypeScript, Vite, Express.js, **React**"* | **No React.** 0 react/vue/angular entries in `package.json`, 0 `.tsx`/`.jsx` in the repo. `studio-desk.md:20` (same auditor's partition) says *"vanilla TS frontend, no framework"* |
| 9 | `:136` | Tier 2: *"**Deployment**: Standalone processes (**not in main docker-compose**)"* | Studio-Desk, the first Tier-2 member, **is** in compose: `docker-compose.yml:311` with `profiles: [studio-desk, all]` at `:342`. Contradicted by `:75` in the same file, by `frontend_architecture.md:11`, and by `studio-desk.md:21` |

### `corpus/services/sentinel.md` — 2 (D1, D2)

| # | anchor | the false claim | what is true |
|---|---|---|---|
| 10 | `:12` | *"**Language**: Go 1.25"* | **Go 1.26** — `go.mod:3` `go 1.26.0`, `Dockerfile:2`/`Dockerfile.dev:2` `golang:1.26-bookworm`, `sentinel/CLAUDE.md:9`. **Actionable**: the same doc's `:115` says `go run main.go`, so a reader who provisions 1.25 gets a hard `go.mod requires go >= 1.26.0` failure |
| 11 | `:22` | *"cheap to operate (**256 CPU / 256 MB** on ECS)"* | `terraform/locals.tf:4-5` → `service_cpu = 256`, `service_memory = **128**`. CPU right, memory wrong by 2× — §5 rule 17's exact shape: a conjunction false while one conjunct measures correctly |

### One each — 8

| # | file:anchor | the false claim | what is true | auditor |
|---|---|---|---|---|
| 12 | `graphql-wundergraph.md:79` | *"**Ports**: host **5050** → container 8080"* — un-fenced, present tense, in *Architecture & Code Map* | No `5050` at HEAD (`grep -c 5050 docker-compose.yml` = 0, rc=1); prod is 8080→8080. The **same doc** at `:174-176` says `localhost:5050` refuses the connection. **This is the claim iter-40 swept at 8 sites outside the scope — see `D-M257x-41-2`** | A |
| 13 | `external_services.md:788` | *"Consistent with **`:447`** above, where the same correction is already recorded"* | `:447` is the **header row** of the *Subgraph routing URLs* table. The correction is at `:512`. 65 lines off, onto unrelated content | A |
| 14 | `ai-readiness.md:37-43` | *"The patch **must re-anchor** — this is the M246 drift-ledger **D-07** item, owned by **v2.7 M250**"* | **The re-anchor already LANDED, at v2.7 M254.** The manifest reads `path: internal/aireadiness/readiness.go` and its own header says *"v2.7 M254 RE-POINT"*. The doc contradicts itself at `:458` (present tense). States completed work as outstanding | B |
| 15 | `roadrunner.md:23-25` | *"Unlike intelligence / skiller / skillpath / **jobsimulation** (which were removed from `repos.yml` + `docker-compose.yml`…)"* | **jobsimulation is in BOTH** — `repos.yml:17`, `docker-compose.yml:83`, `profiles: [graphql, jobsimulation, all]` → it starts on a bare `make up`. Contradicted by `architecture_overview.md:188` and `services/README.md:20-21`, **and by roadrunner.md's own `:26`** | C |
| 16 | `messenger.md:110` | skill-path read cited at `assignments.go:815` | `:815` is `}))` inside `getEmailNotificationForSimulation` — a *simulation* lookup. The skill-path read is `:828` (`h.cms.GetSkillPath`). Claim true, anchor 13 lines off into the wrong function | C |
| 17 | `platform-migration-status.md:60` | *"Owns cms · jobsimulation · skiller · skillpath in-process … **wired at `app/main.go:604`**"* | `:604` wires **jobsimulation only**. skiller `:573`, skillpath `:634`, cms `:1034`. The anchor resolves for 1 of the 4 domains the sentence attaches it to | F |
| 18 | `security_compliance.md:7` · `architecture_overview.md:36`, `:201`, `:243` · `architecture/README.md:23` | *"AI providers are routed through **EU endpoints first**"* / *"EU-first routing"* / **`:243`: *"EU-first routing — Azure OpenAI EU → Azure OpenAI US → direct OpenAI"*** | The sweep's own canonical statement is *"There is **no** ordered EU-first fallback chain"* (`external_services.md:537`). **`architecture_overview.md:243` still publishes the retracted LADDER verbatim**, in the file most readers hit first. The **identical shared-library table row** was rewritten in `service_taxonomy.md:110` and left verbatim in `architecture_overview.md:201`. At `security_compliance.md:7` the sweep **edited the second half of that very sentence and left the retracted clause in the first half** | G |

### Out of clause-5 scope — routed, not counted

| # | site | what is wrong |
|---|---|---|
| G6 | `content-stories-routes.md:23`, `:456`; `content-stories-spec.md:301`, `:385` | *"A demo academy runs with no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, **so it serves its committed FS catalog**."* There is **no FS-as-published fallback** — `serverTenant.js:145` falls back to `emptyCatalogView()`, and its own doc-comment says the cutover is *"not reversible-on-error"*. A demo grid renders 65 cards **only** because the default-on rext demo-patch `academy-fs-published-fallback` restores it. iter-39 retracted this in `ant-academy.md:66-72`; the ops twins survived — **and iter-40's claim-scoped sweep did not have this claim on its list** (`D-M257x-41-2`) |

---

## How much of this is repair-induced? — the number that matters for the escalation

Classified per blocker by whether iter-39's sweep **created** it (new text, over-correction, half-applied
edit, mechanical splice) or merely **failed to catch** a pre-existing claim:

| | blockers | which |
|---|---|---|
| **INDUCED by the iter-39 repair** | **9** | 1 (over-correction in its own new blockquote) · 2 (its own new enumeration) · 4 (added the twin row, left this one) · 5 (rewrote this fence; the `:69`↔`:76` contradiction is inside its own text) · 6 (blockquote spliced into a list, bullet orphaned) · 7 (a false retraction it authored) · 13 (a cross-ref it introduced) · 18 (edited half a sentence, left the other half; fixed one twin row, not the other) · 3 (its table) |
| **GENUINE pre-existing, missed by five passes** | **9** | 8, 9 (`service_taxonomy` rows) · 10, 11 (`sentinel` rows) · 12 (`:5050`) · 14 (`ai-readiness` staleness) · 15, 16, 17 (all three in files never opened) |

**A clean 50/50 split.** Half of what the sixth pass found was **manufactured by the fifth pass's repair.**

> **This is the finding that ends the loop.** The residual is not converging to zero because **each repair
> injects new defects at a rate comparable to what it removes.** `37 → 18` is real improvement, but 9 of
> the 18 are the repair's own children — so a seventh pass would repair 18, induce ~9, and measure ~9–15.
> **The fixed point of this process is not zero.**

---

## What five passes have now established as CLEAN (the negative results are load-bearing)

Files repaired in earlier passes and **confirmed holding** this pass, often against adversarial re-grep:
`hiring.md` — **repaired twice, defective after both, now CLEAN** across ~40 exact anchors · the
*"there is no `manager` Casbin role"* fix (verified **three ways** incl. a live `p_type='g2'` query) ·
*"standalone `authn` is imported by nothing"* (0 hits across every `go.mod`/`go.sum`/`*.go`, with a positive
control finding `colony` in 8) · *"`gen.py` registers exactly nine arguments, `--template` has zero
consumers"* · the **5→4→3→1** subgraph ladder and *"the commit's own `2→1` subject is what's misleading"*
(reproduced independently by **four** auditors) · `bba862f` correctly reported **not** an ancestor · no
`type Subscription` in the backend SDL · taxonomy **42,790 / 22,470** live · all 11 `academy_*` names
plural-verified · **the 135-vs-112 denominator ambiguity RESOLVED in the doc's favour** — 112 is a grep
artifact (`grep '^\tent.Schema$'` misses 23 gofmt one-liners), so **`D-M257x-39-3`'s second stated reason
for refusing the tenancy edit no longer holds**.

**G's positive control:** all **91** diff hunks read in full with file context; **~110** introduced anchors
resolved; **zero** failed to resolve; every one of 13 cited shas ancestry-checked. **For a fifth
consecutive pass, the `file:line` anchors a sweep introduces are correct and its PROSE is not.**

## UNVERIFIABLE classes (unchanged; nine auditors have now independently confirmed the boundary)

`gh` absent → every GitHub archive-status claim · private Go module internals (colony/proto/ai/authn/
taxonomy not cloned) · alignment **scores** (runners cannot build) · production-only infra and **legal**
facts · `db-backup.md` in near-entirety (no repo, 0 references in `platform/`) · prod state of the legacy
schemas · historical demo-run measurements. **An unreachable ground truth is not a false claim.**
