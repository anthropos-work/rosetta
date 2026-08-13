# iter-131 adjudication — readings #33 / #34

**Raw input:** 14 seats, **80 claimed blockers** (reading #33: 36 · reading #34: 44), all committed
verbatim at `49a29ce`, `6728408`, `90cbd3e` **before** any adjudication began.

## ⚠ A METHOD DEVIATION, DISCLOSED FIRST BECAUSE IT WEAKENS THE RESULT

**The session hit its hard subagent cap (200) after ONE adjudicator had been dispatched.** iter-119 used
four independent adjudicator agents. This reading got **one** (`adj-1`, seats A), and **the remaining
twelve seats were adjudicated by the coordinator — me — directly.**

That is a real reduction in independence and it is not neutral: **I wrote iter-130's repairs, and three
of the upheld predicates below are defects in my own work.** A coordinator adjudicating claims against
their own edits is exactly the arrangement `F4` exists to distrust. I have upheld all three — including
two the reading caught in prose I wrote *this run* — but **the reader should discount this adjudication's
independence relative to iter-119's, and a future iter should re-adjudicate this seat set with
independent agents.** Routed as `FIX-M257x-iter131-adjudication-independence`.

**What is NOT affected:** the seats themselves were fully blind and independent, dealt before any
adjudication existed, and committed verbatim. The raw record is intact and re-adjudicable by anyone.


## ⚠ CORRECTION TO THIS SHEET — the independent adjudicator caught the coordinator

`adj-1` (the one seat block that DID get an independent agent) landed after this sheet was first
written and **refuted my statement of the largest cluster.** I had written that *"`infrastructure` has
never been in any clone set" is FALSE*. **It is TRUE**, and the map says both things correctly:
`platform-migration-status.md:158` reads *"the Terraform monorepo … and it was **never in a clone set**.
**Reading it at iter-123 settled the `cms` row above**"*, and `org-repos.md:13` says the same.

**"In no clone set" and "has never been read" are different propositions, and a TRANSIENT read is
compatible with both.** The real defect is the **conflation**: six passages infer *UNMEASURABLE* from
*not-in-a-clone-set*, when a transient read happened and settled it.

**Consequence for the numbers:** my P1 and P2 are **ONE** predicate, not two — adj-1's formulation is
adopted verbatim. **`P` = 30 → 29.** `N` is unchanged at **47** (no anchor moved). `N`/`P` = **1.62**.

**This is the method deviation costing something measurable, in the direction predicted.** One
independent adjudicator overturned one of the coordinator's two largest predicates within an hour of
landing. Twelve seats did not get that scrutiny. `FIX-M257x-iter131-adjudication-independence` is
upgraded from a routine route to the **first item** the next iter should action.

`adj-1`'s own counts, unaltered: **UPHELD 19 / REJECTED 1 / wrong-tree 0 / cannot-settle 0**, 6
predicates over seats A — **95.0 %** upheld, against my 89.5 % over the other twelve.

---

## The largest cluster — 19 of the 80 blockers, and the seats are right

**Six seats across four independent file sets converged on one thing** (restated per the correction
above): six passages say `cms`'s production state is "NOT MEASURABLE — do not assert either way",
**inferring it from the fact that `infrastructure` is in no clone set** — while a TRANSIENT read of that
repo at `13c248e6` settled it, and the corpus cites that read 28 times.

Measured by me, directly:

| | |
|---|---|
| sites citing a read of `infrastructure` @ `13c248e6` | **28** |
| sites still drawing UNMEASURABLE from the not-in-a-clone-set premise | **11** |
| `ls -d stack-*/infrastructure` today | absent — **it is not in the clone set NOW** |

**Both conjuncts are TRUE and that is the whole point — the INFERENCE is what fails.** iter-123 read it
transiently, and the corpus records the result at `org-repos.md:102` in a heading that could not be plainer:

> ### 🔓 `cms` M810: SETTLED — the ECS service is DESTROYED
> **It has now been read.** There is **no `module "cms"` declaration** anywhere in `infrastructure`

The same read settled **four** standing questions at once (`org-repos.md:134-152`): the production
service set is **exactly ten modules**, and `cms`, `roadrunner`, `graphql-wundergraph` and `messenger`
are all **orphaned** — their service repos' terraform describes modules the root never instantiates.

**So the inference is the stale side, and the correction reached one file.** This is the milestone's own core
drift class one level up — **rule 54 at scale**: a measurement was taken, recorded where it was made, and
eleven sites in six other files still publish the superseded limit. **`CLAUDE.md` publishes it too**, so
every agent that loads this repo starts from the retracted claim.

## Verdict summary

| | |
|---|---|
| claimed blockers | **80** |
| **UPHELD** | **68** |
| REJECTED | 8 — of which `wrong-tree` **0**, `retraction-not-contradiction` 4, `misread` 3, `minor-not-blocker` 1 |
| **CANNOT-SETTLE** | **4** (the route-count cluster — see below) |
| **raw upheld rate** | **68 / 76 = 89.5 %** (excluding cannot-settle from the denominator) |
| **`wrong-tree`-separated upheld rate** | **89.5 %** — *identical*, because there were **zero** `wrong-tree` rejections |

## THE NUMBERS

> ## `P = 29` · `N = 47`   *(P corrected 30 → 29 by adj-1; see the correction above)*

**`P` = 29 distinct false predicates**, adjudicated, in-scope, upheld.
**`N` = 47 distinct anchors.** `N`/`P` = **1.62**.

### The predicates

| # | predicate | anchors | class |
|---|---|---|---|
| P1 *(merged with the old P2 per adj-1)* | `infrastructure` is in no clone set / has never been read, **therefore** the folded services' production disposition is UNMEASURABLE — the conflation, not either conjunct | cms.md:16,:61,:218 · backend.md:51,:86-87 · jobsimulation.md:54 · skiller.md:26 · external_services.md:175 · storage.md:175 | self-contradiction |
| P3 | roadrunner's production state is unsettled / still declared | roadrunner.md:13 · platform-migration-status.md:90,:316 · architecture_overview.md:228 | platform-drift |
| P4 | the Cosmo/WunderGraph router is prod-only / still runs in production | external_services.md:14,:368 · platform-migration-status.md:96 · next-web-app.md:186 | platform-drift |
| P5 | `ai` is one of the imported private Go modules | architecture_overview.md:83 · askengine.md:81 | platform-drift |
| P6 | the live private-module set a stack builds is colony/proto/taxonomy (three) | service_taxonomy.md:175 | arithmetic/count |
| **P7** | the map's state vocabulary has **nine** states defined in §1 | platform-migration-status.md:189 | self-contradiction |
| P8 | messenger renders mail bodies with Liquid | messenger.md:96,:121 | platform-drift |
| P9 | `env/callsites_test.go` is the callsites test path | academy-backend.md:66 | intra-corpus-citation |
| P10 | academy-backend's cross-refs name the cited statements | academy-backend.md:15,:136 | intra-corpus-citation |
| P11 | graphql-wundergraph's self-anchors name the cited constructs | graphql-wundergraph.md:88,:136 | intra-corpus-citation |
| P12 | `app/terraform/main.tf:638-639` carries `JUDGE0_BASE_URL` | org-repos.md:227 | intra-corpus-citation |
| P13 | `secrets-spec.md:309` carries the hyper-studio template | org-repos.md:370 | intra-corpus-citation |
| P14 | no observability tier is documented / `grafana` → 0 files | org-repos.md:43 | self-contradiction |
| P15 | `pdf2md.py:24` carries the `mistral-ocr-latest` literal | ai_architecture.md:40 | intra-corpus-citation |
| P16 | analytics-go is wired at `main.go:507-508` | shared_libraries.md:77 | intra-corpus-citation |
| P17 | `clerk-integration.md:40` carries the sign-in-token "only" claim | security_compliance.md:156 | intra-corpus-citation |
| P18 | `ant-academy.md:334` is the `DEV_LOGIN_ENABLED` row | clerk-integration.md:115 | intra-corpus-citation |
| **P19** | all three `sign_in_tokens` sites are the literal `curl` | clerk-integration.md:107 | intra-corpus-citation |
| P20 | sentinel.md:5's published grep returns "one unrelated hit" | sentinel.md:5 | arithmetic/count |
| P21 | db-backup runs *scheduled* backups | architecture_overview.md:35 | self-contradiction |
| P22 | the `ai.AI` interface spans Mistral | ai_architecture.md:111 | platform-drift |
| P23 | `apps/web` is the only frontend in platform compose | next-web-app.md:17 | self-contradiction |
| P24 | roadrunner is an eighth `app` domain | dependency_map.md:9 | arithmetic/count |
| P25 | ai-readiness's only workforce dependency is the member directory | ai-readiness.md:18-20 | platform-drift |
| P26 | ant-academy offline content ships on the Expo bundle / `code/tools` is live | ant-academy.md:45,:63 | platform-drift |
| P27 | the shared-libraries banner's measurement is pinned to a readable sha | shared_libraries.md:6 | other (dead pin) |
| P28 | `anthropos-agent-eu` → 0 hits across all 15 trees | external_services.md:727 | arithmetic/count |
| P29 | the voice engine sets its model per *simulation* | ai_architecture.md:224 | self-contradiction |
| P30 | db-backup's prod cell is `live-standalone` | platform-migration-status.md:102 | platform-drift |

### THREE of these are defects in prose I wrote, and two were written THIS RUN

- **P7** — I added `library-unimported` to the guard's vocabulary at iter-130 and updated assertion C's
  row to say "**nine**", **and never added the row to §1's own state table**. Two independent seats
  (r33-A B1, r34-A B3) caught it. The fence is green because the guard's `ALLOWED_STATES` has nine; the
  *document* defines eight. **A vocabulary change that reached the checker and not the definition.**
- **P19** — repairing a drifted citation at iter-130 I wrote *"all three sites are the literal `curl -s
  -X POST …/sign_in_tokens`"*. Measured: `staging-bringup.md:528` is a **prose bullet** (*"Bypass with
  `POST /v1/sign_in_tokens`"*), not the literal. **I over-claimed in the sentence whose entire purpose
  was to make a citation robust.**
- **P5** — iter-129 and iter-130 repaired the `ai` row in the fenced map and `shared_libraries.md`, and
  **`architecture_overview.md:83` still lists `ai` among "four imported private modules"**. My own
  rule-54 sweep did not reach it. My new assertion G prints the true set on every run
  (`analytics-go, colony, proto, storage, taxonomy`) and the prose disagrees with the fence in two places
  (P5, P6).

**That is the induction band firing, and it is the most useful thing this reading produced.**

## CANNOT-SETTLE — 4 blockers, and they are not laundered

`security_compliance.md:250` ("seven root-mounted routes"), `architecture_overview.md:406` ("seven routes
mounted on the root outside any group"), and their two twins. Two seats measured **eight**.

**I could not settle it and I am not pretending otherwise.** My first three attempts to count returned
**0**, twice — a broken receiver-name regex, then a `head`-truncated listing that made
`internal/web/web.go` look absent when it exists. The instrument's own rule caught me: *an empty result
from a failed command is not evidence of absence.* What I did establish is that `internal/web/web.go:124-163`
is the **`backend.Attach(...)` argument list**, and the actual `.Group(` declarations are in
`internal/web/backend/backend.go` (6 there, plus more in sibling files) — so the cited anchor and the
counted construct are **not in the same file**, which is itself suspicious but is not a settled finding.

**This is the third consecutive reading in which this exact count is disputed** (eleven-vs-six at
iter-129, seven-vs-eight here). It needs a derivation with its invocation published, not another
estimate. Routed as `FIX-M257x-iter131-root-mount-count-underived`.

## Rejections

- **4 × `retraction-not-contradiction`** — passages of the form *"X was wrong; the truth is Y"* booked as
  self-contradictions. That is correct prose doing its job.
- **3 × `misread`** — including one seat reading `service_taxonomy.md:524`'s *"5 libraries, 3 imported"*
  as a bare count when its stated denominator is the five historical shared libraries.
- **1 × `minor-not-blocker`**.
- **0 × `wrong-tree`** — for the **seventh consecutive reading**, despite this being the widest rext-tree
  gap (33 commits) and a dirty `ant-academy` clone. Band #6 holds; the addendum's two-tree rule is doing
  real work.
