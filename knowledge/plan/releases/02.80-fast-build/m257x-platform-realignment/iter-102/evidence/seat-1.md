# seat-1 report

**File owned:** `corpus/services/backend.md` (only file edited). 371 lines after repair, was 319.
**Anchors booked:** 11. **Sites found:** 13. **Sites repaired:** 13. **Left standing: 0.**

**Ground truth re-measured at this seat's open** (no `git fetch` issued; every clone read at its own HEAD):
platform `0c91421d` · `app` `ad9f3c498e9c…` (== `origin/main`, 5 commits past `2035f9a40c57…`) ·
demo pin `b948604ff861…` · `storage` `4ce8ece5` · `messenger` `fa47850d` · `roadrunner` `87d8d443`.
**Load-bearing measurement for CANON-3:** `git diff --stat 2035f9a4 ad9f3c49` touches **5 files, none of
them Go** (`.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`,
`terraform/main.tf`, `terraform/variables.tf`), and `git diff --quiet 2035f9a4 ad9f3c49 -- main.go` →
**`main.go` is byte-identical**. So every `main.go` line number this corpus pinned to `2035f9a4` still
resolves at `ad9f3c49`; only the **label** expired. `terraform/*.tf` did move, which is why those
citations were left pinned rather than re-derived.

---

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "**All of their Connect-RPC surfaces are served on `app`'s single RPC mux — and nothing outside the process calls them.**" | `backend.md:29` (`app-rpc-mux-universal`) · `backend.md:29-30` (`eight-folded-rpc-mux`) | Of the **eight** folded services, only **three** have a handler on `app`'s mux. Enumerated from `main.go` @ `ad9f3c49`: `UsersService` (`:1297`), `OrganizationsService` (`:1298`), `SkillerService` (`:1306`), `JobSimulationService` (`:1314`), `CMSService` (`:1323`, inside `if cmsRPCServer != nil`), `lab.v1.LabSessionService` (`:1337-1338`) — **six** handlers, of which **three are `app`'s own** and were never folded in. `skillpath`, `roadrunner`, `storage`, `messenger`, `customerio-sync`: **none**. `SkillPathSessionService` → **0** occurrences in `app` Go source at `ad9f3c49` (M506 removed it). `storage@4ce8ece5:sdk/storage/v1/service.go` and `messenger@fa47850d:internal/rpcsrv/rpcsrv.go` declare real Connect services **in their own repos, on their own muxes**; `app` imports `storagev1connect` only in two *test* files. The **second clause is TRUE and was preserved verbatim in substance**. | 1 |
| 2 | "**`app` owns the `skillpath`, `jobsimulation`, `cms` and `ai_usage` Redis Streams** — both producer and consumer in-process. **`skiller` is NOT a fifth member: it is consumer-only**" | `backend.md:33-34` (`app-stream-set-omits-backend` **and** `app-stream-set-cardinality` — two predicates, one site, both upheld, both fixed, not merged) | The both-ways set is **five**, not four: `backend` is missing. Publishers @ `ad9f3c49`: `main.go:325` (`pubsub.NewPublisher(serviceName, …)`, `SERVICE_NAME` default `backend` at `:230-232`), `:746` `SKILLPATH_STREAM`, `:1149` `CMS_STREAM`, `jobsimwiring/wiring.go:132` `AI_USAGE_STREAM`, `:185` `JOBSIMULATION_STREAM`. Subscribers: `subscriber_wiring.go` builds a **6-entry map** — `subs[d.Streams.Backend]` at `:248`, whose own comment is *"Backend is app's OWN self-stream (SERVICE_NAME, "backend"): events app publishes and also consumes in-process"* (`:112-113`). So **six subscribers, five publishers**, and `skiller` would be a **sixth** stream, never a fifth member. | 1 (+1 coherence, row 3) |
| 3 | *(no false claim — the true statement at the sibling site, edited only to stop the counts reading as a contradiction)* | `backend.md:305` → now `:317` | `four` (application-stream subtotal) / `five` (both-ways total) / `six` (subscriber count) are three different partitions. Made explicit, and the `ad9f3c49` publisher anchors + map-based subscriber shape added so the `b948604f`-pinned `main.go:1276` sentence cannot be read as current. | 1 |
| 4 | "the most recent set of migrations (May 2026) touched simulation-type definitions and content JSON defaults" | `backend.md:301` (`migrations-most-recent-may-2026`) | **170** `terraform/migrations/*.sql` at `ad9f3c49`, **169** at `b948604f`. **46** landed after 2026-05-31. Head at `ad9f3c49` = `20260803143844_ai_readiness_recommendation_path.sql` (**2026-08-03**); head at the demo pin = `20260731154527_academy_chapter_progress_completed_at.sql`. Last May migration = `20260529072659_add_lab_session.sql`, **46 back from the head**. None of the newest five touches simulation-type definitions or content JSON defaults. *(Note: my head is one migration newer than adj-1's iter-101 reading — `20260803143844` postdates the demo pin adj-1 measured at.)* | 1 |
| 5 | "In production terraform the re-pointed pair is at `http://backend.internal.anthropos:8081`." | `backend.md:241` (CANON-2 `prod-terraform-8081`; booked MINOR at iter-101, repaired anyway) | **Assertion DROPPED, not softened.** Two independent mechanisms, 2026-08-06: `git grep` at each clone's own HEAD over the **44** tracked `.tf` files in the 13-repo `stack-demo` clone set → **0 files**; raw filesystem `find -name '*.tf'` sweep, **59** files → **0** (positive control `service_discovery_namespace_id` → 25 files). The literal's sole occurrence anywhere is `app/knowledge/service-dependencies.md:52` @ `ad9f3c49` — a **markdown KB page, not terraform** — and it is **past tense**: *"it used to reach … and folding it in at v9.0 closed that edge"*, under *"**There are no external callers of app's RPC mux left.**"* And `infrastructure`, the only tree that could settle it, is in no clone set. **TRAP A honoured: not re-anchored at `service-dependencies.md`**, and the doc now says so explicitly. | 1 |
| 6 | "In prod terraform the address is `http://backend.internal.anthropos:8081` —" | `backend.md:112` — **a 5th CANON-2 site, unbooked**, found by predicate sweep | Same repair. The block's *derivation* (`locals.tf:6` `project = "backend"`, `main.tf:58` namespace **id** variable, `locals.tf:8` `rpc_port = 8081`, `main.tf:185-186` port mapping — all verified at `b948604f`) is TRUE and was **kept**, re-framed as a derivation rather than a reading. Added: the namespace *name* `internal.anthropos` is in **no `.tf` in any clone**, and the `app/terraform/variables.tf:197,230` pin **has moved** (`cms_rpc_address` → `:309`, `storage_rpc_addr` **deleted** for `storage_s3_bucket`/`storage_s3_public_bucket` at `ad9f3c49`) — stated as a pin, not silently swapped. | 1 |
| 7 | "re-derived at `app` origin/main `2035f9a`" | `backend.md:70` (CANON-3) | CANON-3 **move (2)** — a currency claim (*"`app`'s OWN docs **still** list it"*). Re-derived at `ad9f3c49`: `app/CLAUDE.md:109` and `app/knowledge/architecture.md:28` **both still list `SkillPathSessionService`**, unchanged from `2035f9a4` (the one CLAUDE.md commit in the 5 added a line at `:345`, below the anchor). | 1 |
| 8 | "no `NewPublisher` names `SKILLER_STREAM` at `b948604f` or at origin/main `2035f9a4`" | `backend.md:138` (CANON-3) | The **negative holds at all three refs** (`b948604f`, `2035f9a4`, `ad9f3c49`) — re-derived. `SKILLER_STREAM` in `*.go`: **1** occurrence at `b948604f` (`main.go:1276`, an `AddSubscriber`), **6 across 3 files** at `ad9f3c49`, none a publisher. Also recorded: at `ad9f3c49` that registration **moved** into `buildStreamSubscribers` (`subscriber_wiring.go:209`) + one loop at `main.go:1579-1581`, so there is no standalone skiller `AddSubscriber` line to cite there. Label dropped, sha kept, drift disclosed. | 1 |
| 9 | "re-derived at `app` origin/main `2035f9a`" | `backend.md:254` (CANON-3) | CANON-3 **move (1)** — pinned historical state. `main.go:524`/`:525` still are `internalstorage.NewManager` / `NewPublicManager`; `main.go:504` still reads *"STORAGE_RPC_ADDR is gone"*; `STORAGE_RPC_ADDR` still has **3** Go occurrences, all comments. Label replaced with `ad9f3c49` + the byte-identical-`main.go` fact that makes every `2035f9a4` `main.go` pin still resolve. | 1 |
| 10 | "as it is at origin/main `2035f9a4`" | `backend.md:299` (CANON-3) | `git ls-tree <ref> migrations/` is empty at `b948604f`, `2035f9a4` **and** `ad9f3c49`. Label replaced; both shas kept; `2035f9a4` explicitly re-labelled *a pin, not the tip, 5 commits behind*. | 1 |
| 11 | "The one cross-process edge left on a local stack is `backend → sentinel`." | `backend.md:31-32` — **CANON-1 twin inside my file, not on CANON-1's anchor list** | Canonical CANON-1 form applied verbatim in substance. `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48` @ `0c91421`) — the RPC-edge clause and the zero-`*_RPC_ADDR` clause are **TRUE and preserved**; the generalisation is false. `GOTENBERG_URL=http://gotenberg:3200` (`:57`), `gotenberg` `profiles: [core, backend, all]` (`:183`), consumed by a plain `http.NewRequestWithContext(ctx, "POST", …)` at `app/internal/converter/gotenberg.go:31`; `JUDGE0_BASE_URL` (`:59`). | 1 |
| 12 | "the only cross-process edge a local stack has is `backend → sentinel` (`AUTHORIZATION_ADDRESS`, `docker-compose.yml:48`)" | `backend.md:241` — **second CANON-1 twin in my file**, same line as row 5 | Same canonical form. Two distinct predicates on one line (CANON-1 + CANON-2); both repaired, **not merged**. | 1 |
| 13 | *(induction prevention, not a booked defect)* | `backend.md:13` — the self-reference `see `:36`` | My row-1/row-11 edit lengthened the banner and pushed the *M810 prod teardown is UNEVEN* bullet from `:36` to `:68`. The numeric self-ref was replaced with a **name**, not a new number, so it cannot rot again. This is the exact failure iter-100 induced (a parenthetical that moved a table and left the numbers behind). | 1 |

---

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `app-rpc-mux-universal` + `eight-folded-rpc-mux` (one statement, two bookings) | 2 | 1 false + 1 verified-correct (`:66`/now `:92`, which already enumerated 5 unconditional + CMS conditional) | 1 | `grep -n 'mux' backend.md` (5 hits, all classified); mux enumerated from source at **both** refs via `git grep 'mux.Handle' <ref> -- '*.go'` and `git grep 'v1connect.New.*ServiceHandler'` |
| `app-stream-set-omits-backend` + `app-stream-set-cardinality` (two predicates, one anchor) | 2 | 3 (`:33-34` false; `:264`→`:317` true-but-count-ambiguous; `:136`→`:185` true) | 2 | `git grep 'NewPublisher\|AddSubscriber\|_STREAM' ad9f3c49 -- '*.go'`, then read `subscriber_wiring.go:108-250` and `main.go:1520-1590` in full |
| `migrations-most-recent-may-2026` | 1 | 1 | 1 | `git ls-tree --name-only <ref> terraform/migrations/`; corpus sweep `grep -rn 'most recent set of migrations\|simulation-type definitions' corpus/ CLAUDE.md` → **1 site, mine** |
| CANON-2 `prod-terraform-8081` | 1 (`:241`) | **2** — booking under-counted by 1× within my file | 2 | `grep -n '8081\|internal\.anthropos' backend.md` (5 hits: 2 claims, 2 derivation, 1 unrelated port line). Ground truth by two mechanisms, both with a positive control |
| CANON-3 currency pin | 4 | 4 — `grep -n 'origin/main\|2035f9a' backend.md` returns **exactly** `:70`, `:138`, `:254`, `:299`; no fifth | 4 | Per-site classification into move (1) *pinned state* vs move (2) *currency claim*, then re-derivation at `ad9f3c49` for each |
| CANON-1 `sentinel-only-cross-process-edge` | **0 booked to me** — backend.md is not on CANON-1's anchor list | **2** in my file (`:31-32`, `:241`) | 2 | `grep -n 'cross-process\|only cross\|one cross' backend.md`. Repaired because **no other seat may touch my file**, so leaving them would have left backend.md contradicting the canonical repair 12 files over. Canonical wording applied — I did not compose my own |
| **TOTAL** | **11** | **13** | **13** | |

**Booked-to-found ratio 11 → 13 (1.18×).** Lower than the 3× the brief warned of, and I believe that is a
real property of this file rather than a shallow sweep: `backend.md` was already the most heavily
re-ground file in the corpus, so most of its paraphrases had been converged in earlier iters. The one
genuine widening is CANON-2 (`:112`), which the union itself flagged as under-counted.

---

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| `app-stream-set-omits-backend` | **`CLAUDE.md:282`** (orchestrator-owned) | *"The `skillpath`, `jobsimulation` and `cms` streams have `app` on **both** ends."* Same omission family — drops `backend` **and** `ai_usage` from the both-ways set. **Weaker than the booked defect**: it does not close the set (no *"NOT a fifth member"* clause), so it is incomplete rather than false. Reporting it because it is the last site of the pattern I repaired. |
| `eight-folded-rpc-mux` | `corpus/services/jobsimulation.md:46` · `corpus/services/cms.md:48` | *"`JobSimulationService` / `CMSService` is served on `app`'s single RPC mux."* **VERIFIED TRUE — these are NOT twins.** Both services *do* have a handler (`main.go:1314`, `:1323` @ `ad9f3c49`). Recorded so a later sweep does not "repair" a correct sentence into a wrong one. The false form is the **universal** over all eight, which existed only in my file. |
| CANON-2 `prod-terraform-8081` | `corpus/services/cms.md:196` (seat 8) | Already repaired by seat 8 as I read it, with the **same** two-mechanism measurement (44 tracked / 59 filesystem `.tf`, 0 hits). Recorded as independent convergence, not as an outstanding site. |

---

## Induced and corrected

Two defects reached my file **from the canonical sheet, not from my own derivation**. Both were flagged
mid-seat by the orchestrator and both are fixed. Recorded separately because the distinction is
load-bearing for this iter's induction-rate measurement.

| # | defect | origin | where it landed | fix |
|---|---|---|---|---|
| I-1 | CANON-1's replacement text ended `` …via `JUDGE0_BASE_URL` (`:59`) ``. A bare `:59` resolves against the **most recently named file**, `app/internal/converter/gotenberg.go`, which is **53 lines** → `anchor-construct-guard` `[anchor-out-of-range]` | `canonical-repairs.md` §CANON-1 as issued (corrected at source since) | `backend.md:48` and `backend.md:282` | Anchor made explicit: `` (`docker-compose.yml:59`) ``. I also pre-emptively made the adjacent `` `:183` `` explicit at both sites — same class, one hop further along the bare-continuation chain |
| I-2 | CANON-2's replacement said only *"not measurable from this repo"* — half the verdict, and the weaker half. Stating only unmeasurability reads as *"we didn't check"* when the adjudicators measured **zero** | `canonical-repairs.md` §CANON-2 as issued (corrected at source since) | `backend.md:139-155` and `backend.md:294` | Both halves now present at both sites, **measured zero first**, unmeasurability second, plus an explicit TRAP A instruction not to repoint at `service-dependencies.md` |

**Self-induced defects: 1, caught by me before publication.** My first draft of the mux bullet wrote
*"roadrunner and customerio-sync never had a Connect surface."* Checking it before committing the
sentence: `git -C roadrunner grep -n 'v1connect' 87d8d443 -- '*.go'` → `cmd/root.go:87`
`rpcMux.Handle(roadrunnerv1connect.NewRoadRunnerServiceHandler(…))`. **False for roadrunner.** And for
`customerio-sync` it was unmeasurable in principle — that repo has never been in a clone set, so the
claim would have been a TRAP A invention. The published sentence now names roadrunner's real surface in
its own repo and says outright that nothing can be said about customerio-sync's from here.

---

## Guard results (post-repair, run from `.agentspace/rosetta-extensions/stack-core`)

| guard | result | `corpus/services/backend.md` |
|---|---|---|
| `anchor_construct_guard.py` | RED ×4 — `dependency_map.md:7` ×2, `dependency_map.md:58`, `sentinel.md:85` | **0 sites — CLEAN** (was 2 before the I-1 fix) |
| `claim_twin_guard.py` | RED ×8 | **0 sites — CLEAN** |
| `markdown_structure_guard.py` | **OK — no structural damage** | clean |
| `unreadable_repo_claim_guard.py` | **OK — all 7 `module.*_euwest1` mentions marked unmeasurable** | clean |

The 4 remaining anchor REDs and all 8 twin REDs are in seat 3 / 7 / 8 / 9 / 10 files, not mine.
`claim_twin_guard` transiently flagged `backend.md:344` mid-seat: my first migrations rewrite quoted the
refuted sentence **verbatim** in order to retract it. The waiver mechanism requires an entry in a waiver
file inside `rosetta-extensions` — outside my editable scope — so I **paraphrased the retraction** rather
than republish the string. Green since.

---

## Noticed, not repaired

1. **`backend.md:294` — *"Services include `lab.v1.LabSessionService`, `SkillerService`, `JobSimulationService` and `CMSService`."*** Omits `UsersService` and `OrganizationsService`, which my `:29` repair names. **Not false** — the verb is *"include"*, not *"are"* — so under rule 6 I left it. Flagging it because a future sweep comparing the two passages may read a contradiction where there is only an open list.
2. **`backend.md:354` — *"The `public` schema is the largest in the platform."*** Unbooked, and after eight folds it is close to tautological (it is nearly the only application schema left). Not measured, not touched.
3. **`backend.md:19` — *"compose now declares **five** services and `repos.yml` **four** entries."*** Verified in passing and **correct**: `docker-compose.yml` @ `0c91421` declares `sentinel`, `backend`, `studio-desk`, `next-web-app`, `gotenberg`; `repos.yml` lists `app`, `sentinel`, `next-web-app`, `studio-desk`. Worth knowing that `postgresql` and `redis` come from the **included `common.yml`**, not from `docker-compose.yml` — so "five" and the always-on floor of three are counting different files, and a reader could take the two as contradictory.
4. **`app`'s own KB now contradicts `app`'s own code more sharply than before.** `app/CLAUDE.md:109` and `app/knowledge/architecture.md:28` still list `SkillPathSessionService` on the RPC mux at `ad9f3c49`, while `SkillPathSessionService` has **0** occurrences in that same tree's Go source. This is the corpus's Trap C, already documented at `backend.md:102` — recorded only because it is now measured at the newest ref, not because it needs a corpus change.

---

## What I could not settle, and why

1. **The production value of `backend`'s RPC address.** Settled as *unsettleable*, which is the point.
   The deciding declaration is in `infrastructure`, in no clone set. I did **not** repoint at
   `app/knowledge/service-dependencies.md:52` — a markdown page is not terraform, and TRAP A says a
   correctly-cited false statement is worse than a stale one. What I *can* state is bounded and measured:
   0 hits over 44 tracked `.tf` files (git, per-ref) and 0 over 59 filesystem `.tf` files, both with a
   positive control, plus the platform's own KB putting the edge in the past tense.
2. **Whether `customerio-sync` ever had a Connect-RPC surface of its own.** Its repo has never been
   cloned. The doc now says so rather than guessing; only the `app`-side fact (no handler on the mux) is
   asserted.
3. **`CLAUDE.md:282`'s three-stream list.** Inside the same predicate family as my `:33-34` repair, but
   `CLAUDE.md` is the orchestrator's file. Reported in *Twins outside my files*, not edited.

**No commit made. No `git add`, no `git fetch`, no clone written. Zero platform-repo edits.**
