# Adjudicator 2 — iter-101 verdict

Seats adjudicated: `r23-B` (4 blockers), `r24-B` (1 blocker), `r23-D` (4 blockers). **9 bookings, all adjudicated.**

Seat B was read twice; **seat D was read once** — `r24-D` does not exist and was not sought, reconstructed, or
substituted for. `r23-D`'s four bookings are adjudicated on their own merits and are subtotalled separately below.

## Refs re-verified at this adjudication's open

Every clone re-derived, all matching the brief's ground truth:
platform `0c91421d` · app `b948604f` · cms `ca50c817` · next-web-app `bb3313bc` · sentinel `88bc5592` ·
storage `4ce8ece5` · messenger `fa47850d` · graphql-wundergraph `60c229f3` · roadrunner `87d8d443` ·
jobsimulation `462343b0` · studio-desk `14a5442a` · ant-academy `9c3843cd`.
Nested `stack-demo/app/studio` + `stack-demo/cms/studio` both `aeec036a`.
`stack-demo/rosetta-extensions` = `ab81527a` (consumption pin) · `.agentspace/rosetta-extensions` = `09d0607` (authoring).

**Corpus state check.** Corpus HEAD is `8b6d80f`; `04cbcfc..8b6d80f` touches only `iter-101/raw/`, so the corpus
text I graded is byte-identical to what both seats read. (`r24-B` records its corpus HEAD as `1937e1f`, which is
an M257x/57 ancestor 60 corpus files behind — a mis-recorded instrument note in that report, not a difference
in what was read; every `r24-B` anchor I checked lands where it says at `8b6d80f`.)

## Rule applied to every tree choice: the settling tree follows the claim's SUBJECT

Three of these bookings turn on "unpinned anchor that resolves only at `origin/main`". They do not all resolve
the same way, and the discriminator is what the claim is *about*:

- a claim about **what a local stack runs** (FE code, the app binary a demo builds, compose) is settled by
  **the demo's pinned build ref** — that is the tree a reader has;
- a claim about **production infrastructure state** is settled by **the repo's `origin/main`** — a week-old
  demo build pin is not, and never was, evidence about prod.

Applied consistently below: it upholds `r23-B B3`/`B4` and rejects `r23-D B4`.

---

## Verdicts

### r23-B B1 | `corpus/services/ai-readiness.md:305` | UPHELD | IN-SCOPE | `AI_READINESS_URL` is declared at `urls.ts:52`

   evidence: `stack-demo/next-web-app` @ `bb3313bc`, `packages/core-js/src/constants/urls.ts` — `:50` is
   `export const AI_READINESS_URL = '/ai-readiness';`. **`:52` is
   `export const ORGANIZATION_FEEDBACK_URL = '/enterprise/organization-feedback';`** (`:49` `WORKFORCE_URL`,
   `:51` `TALK_TO_DATA_URL`, `:53` `INSIGHTS_URL`). At `origin/main 8297c684c` the constant is at `:51`, so
   `:52` is `TALK_TO_DATA_URL` there. I then swept the file's own history — across the 25 most recent commits
   touching it the constant sits at **41** (2 commits), **50** (3), **51** (2) and is absent in 18; it has
   **never** been at 52 at any ref reachable from this clone. The enclosing block names the ref itself
   (`:299-300`: *"the surrounding anchors below still resolve at HEAD"*), and HEAD of the clone is `bb3313bc`.
   A resolving line naming a different exported constant — an anchor-existence check passes it and a reader
   following it lands on the org-feedback route.
   tree-read: `stack-demo/next-web-app` at `bb3313bc`, at `origin/main 8297c684c`, and across its `urls.ts` history.

### r24-B B1 | `corpus/services/ai-readiness.md:305` | UPHELD | IN-SCOPE | same predicate, same anchor as r23-B B1

   evidence: identical re-derivation to r23-B B1 above. `r24-B` adds the history sweep independently and
   reaches the same result. **Duplicate anchor, duplicate predicate** — see DEDUPLICATION.
   tree-read: `stack-demo/next-web-app` at `bb3313bc` / `origin/main`.

### r23-B B2 | `corpus/services/cms.md:196` (same claim at `:55`) | UPHELD | IN-SCOPE | production terraform still names `http://backend.internal.anthropos:8081`

   evidence: **three-instrument absence check.** (a) `git grep 'backend.internal.anthropos'` at each of the 12
   clones' own HEAD *and* at `origin/main` for the six behind clones: the `:8081` form appears in **exactly one
   file, a markdown KB page** — `stack-demo/app/knowledge/service-dependencies.md:46` @ `b948604f` / `:52` @
   `origin/main`. Positive control good: `graphql-wundergraph@60c229f3` returns `CLAUDE.md:18` and
   `supergraph-config-prod.yaml:6`, both at **:8080**. (b) Raw filesystem grep over **all 59 `.tf` files** in
   the workspace (catches `.gitignore`-hidden and untracked files that `git grep` would miss): **zero** name
   the literal; positive control — 13 `.tf` files match `anthropos`. No NUL-bearing `.tf` exists (byte-counted
   with `tr -dc '\000'`, not `grep -c`). (c) Nested `app/studio` + `cms/studio` grepped at their own ref
   `aeec036a`: 0.
   The only terraform that could carry it does not: `messenger/terraform/main.tf:74-75` sets
   `"name": "CMS_RPC_ADDR"` / `"value": "${var.cms_rpc_address}"` at **both** `fa47850d` and `e9421c68`, and
   `terraform/variables.tf:77-80` declares `variable "cms_rpc_address"` with **no default** — the literal is
   supplied by `infrastructure`, which `cms.md:82-84` itself declares *"not visible to this corpus — the
   `infrastructure` repo has never been in the clone set … report both, assert neither."* Nine lines later the
   same file asserts it anyway, present tense, unpinned.
   Newest prod ground truth points the other way: `app/knowledge/service-dependencies.md:50-53` @ `2035f9a` —
   *"**There are no external callers of app's RPC mux left.** `messenger` was the last one — it **used to**
   reach … at `http://backend.internal.anthropos:8081`, and folding it in at v9.0 **closed that edge**"* — and
   `:58-59`: *"the `messenger` and `customerio-sync` ECS services and their **terraform modules are gone**."*
   Not ref-discipline: the sentence is **present tense (`still names`) and names no ref**, so it claims
   currency rather than a date. Unsupportable at every tree, contradicted by the newest, and self-contradictory
   inside its own file (rule 5).
   tree-read: all 12 `stack-demo` clones at HEAD + `origin/main`, both nested `studio` checkouts at `aeec036a`,
   and the raw worktree filesystem.

### r23-B B3 | `corpus/services/ai-readiness.md:595` | UPHELD | IN-SCOPE | `interviewQuestions` is in the FE type at `useAIReadiness.ts:326`

   evidence: `stack-demo/next-web-app` @ `bb3313bc`, `apps/web/src/hooks/useAIReadiness.ts` — `interviewQuestions: number;`
   is at **`:274`**, inside `export interface AIReadinessCycleTotals` (`:271-278`). **`:326` is `headers: {`**,
   inside `const res = await fetch(url.toString(), { ...init, headers: { Authorization: … } })` at `:324-331`.
   At `origin/main 8297c684c` the field is at exactly `:326`. The blockquote (`:589-599`) names no next-web ref
   and opens *"the **current** dashboard"*; the file's only stated convention is `:300`'s *"still resolve at
   HEAD"*, i.e. the checkout. This is a claim about **FE code a reader/demo runs**, so the demo pin settles it.
   One conjunct of a three-conjunct sentence is false at the grading tree, and it fails as a resolving line
   naming an unrelated construct. Not ref-discipline — the inverse: an unpinned claim resolving only at a
   *newer* ref, not a pinned claim contradicted by newer evidence.
   (Recorded: `r24-B` booked this same passage as a MINOR rather than a blocker, reasoning that it resolves at
   one available ref. I grade the booking as booked; the anchor is false at the tree that settles it.)
   tree-read: `stack-demo/next-web-app` at `bb3313bc` and `origin/main 8297c684c`.

### r23-B B4 | `corpus/services/storage.md:73-75` (HAZARD block `:60-80`) | UPHELD | IN-SCOPE | app's two boot guards at `main.go:518-523` / `:529-535` / `env_guards.go:37-44`

   evidence: the HAZARD blockquote's only pin is `@ platform 0c91421` (`:63-64`), attached to the compose
   claim; under rule 1 a pin's scope is its own block, so the three `app` anchors are unpinned. This is a claim
   about **the binary a stock local stack runs** ("on a stock stack NEITHER manager uses local FS … Nothing
   warns"), so the demo's build pin `b948604f` settles it. There:
   - `git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — `app/env_guards.go` does
     not exist; `git grep 'func deployedEnvironment' b948604f` exits 1 (control: the same grep at `origin/main`
     returns `env_guards.go:37`). The entire "nothing warns" mechanism is uncheckable on the pinned clone.
   - `main.go:518-523` @ `b948604f` is the **public-storage clients** block (`publicStorageClients, storagePublicClient := publicstorage.NewClients(os.Getenv("STORAGE_RPC_ADDR"), …)`);
     `:529-535` is the **academy asset uploader** block. Both resolve; both name the wrong construct.
   - All three are exact at `origin/main 2035f9a`: `:518-523` is the empty-bucket `log.Fatalf` guard,
     `:529-535` the `verifyBucketAccess` block, `env_guards.go:37-44` `func deployedEnvironment()` with
     `case "", "development", "dev", "local", "test": return false`.
   Also confirmed the same-family sibling the seat folded in rather than booking twice: `storage.md:44`'s
   *"the live code is `app/internal/storage/`"* — `git ls-tree b948604f -- internal/storage/` is empty; only
   `internal/publicstorage` and `internal/storagens` exist at the pin.
   Aggravating and decisive: this document polices its own ref discipline everywhere else — `:29` (*"every
   `app` path in this cell resolves at `2035f9a` only"*, *"a block that names two refs makes every anchor in it
   ungradeable"*), `:58` (`@ app 2035f9a`), `:34-35` (which tells the reader outright that `b948604` is *"the
   demo's pinned build ref"*) — and drops it precisely on the one safety-critical claim in the file, a local
   stack writing into the **production** bucket.
   tree-read: `stack-demo/app` at `b948604f` and `origin/main 2035f9a`.

### r23-D B1 | `corpus/services/sentinel.md:85` | UPHELD | IN-SCOPE | `AUTHORIZATION_ADDRESS` is the only service address compose sets; backend→sentinel the one cross-process edge

   evidence: the claim **pins itself** — *"at platform `0c91421`"* — which is the ground-truth ref, so there is
   no tree ambiguity. At `stack-demo/platform` `0c91421d`, `backend`'s `environment:` block carries four more
   cross-process service addresses: **`docker-compose.yml:57` `GOTENBERG_URL=http://gotenberg:3200`**,
   `:66` `REDIS_ADDR=redis:6379`, `:93` `SUPABASE_DB_CONN=postgresql://…@postgresql:5432/…`, `:94`
   `COPILOT_DB_CONN=…@postgresql:5432/…`.
   The gotenberg edge is live on a stock stack, not theoretical: `gotenberg` is a declared compose service
   (`:170-171`, `gotenberg/gotenberg:8`) in `profiles: [core, backend, all]` (`:183`) — the same default
   profile as `backend` (`:110`) — and `Makefile:10` is `PROFILE ?= core`, so `make up` starts both. `app`
   reads the variable at two live sites @ `b948604f` (`main.go:244`, `internal/web/backend/coursebuilder/handler.go:244`)
   and `internal/converter/gotenberg.go:31` issues `http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", …)`,
   executed at `:37` — a genuine cross-process HTTP hop over the compose network.
   The `*_RPC_ADDR` half of the sentence is **true and verified**: 0 occurrences across `docker-compose.yml`,
   `common.yml`, `.env_example`; positive control `AUTHORIZATION_ADDRESS` = 1. It is the generalisation from
   *"the only RPC address"* to *"the only service address / the one cross-process edge"* that breaks.
   Corroborated by two other corpus passages asserting the opposite (rule 5): `corpus/services/gotenberg.md:50`
   — *"`GOTENBERG_URL=http://gotenberg:3200` (injected via the backend's compose `environment:`)"* — and
   `corpus/architecture/dependency_map.md:103` — *"`GOTENBERG_URL=…` is injected via the backend's compose env."*
   `corpus/architecture/architecture_overview.md:321` carries the correct, qualified form (*"the only
   cross-process **RPC** edge out of backend on a core stack"*), which is the shape the repair wants.
   (One narrowing of the seat's argument: the sentence's trailing *"none has the env"* clause refers to
   `AUTHORIZATION_ADDRESS` specifically and is true, so the "contradicts itself inside the same sentence"
   reading is looser than the seat states. The over-claim stands on its own without it.)
   tree-read: `stack-demo/platform` at `0c91421d`; `stack-demo/app` at `b948604f` and `origin/main`.

### r23-D B2 | `corpus/services/jobsimulation.md:145-146` | UPHELD | IN-SCOPE | sentinel is the only cross-process hop and the only service address backend's compose entry carries

   evidence: identical re-derivation to D B1, at the same platform ref. The block cites `docker-compose.yml:48`
   and names no platform sha; the checkout `0c91421d` is level with `origin/main`, so there is no tree the
   claim could be true at. `backend`'s compose entry carries `GOTENBERG_URL` (`:57`), `REDIS_ADDR` (`:66`),
   `SUPABASE_DB_CONN` (`:93`) and `COPILOT_DB_CONN` (`:94`); `backend → gotenberg` is a second live
   cross-process hop in the default `core` profile (`:170-171`, `:183`, `Makefile:10`), reached over HTTP at
   `app/internal/converter/gotenberg.go:31`.
   Booked and upheld separately from D B1 because it is a distinct anchor in a distinct corpus file with
   distinct wording — a claim-scoped repair that fixes one and not the other leaves this one standing. Same
   predicate; see DEDUPLICATION.
   tree-read: `stack-demo/platform` at `0c91421d`; `stack-demo/app` at `b948604f`.

### r23-D B3 | `corpus/architecture/platform-migration-status.md:93` (messenger row, RPC-edge clause) | UPHELD | IN-SCOPE | the only cross-process service address left in a local stack is `AUTHORIZATION_ADDRESS`

   evidence: identical re-derivation to D B1/B2 at platform `0c91421d` — `GOTENBERG_URL=http://gotenberg:3200`
   at `docker-compose.yml:57`, plus `REDIS_ADDR`/`SUPABASE_DB_CONN`/`COPILOT_DB_CONN`.
   Sharper here, and independently sufficient under rule 5: **the same file, twelve rows down, states the
   counter-evidence.** `platform-migration-status.md:105` (the `gotenberg` row) reads *"third-party image,
   `docker-compose.yml:170-171` (`gotenberg/gotenberg:8`), **default `core` profile** (`:183` …)"* and grades
   its `fresh local stack` column **live-standalone**. A service the map itself records as live in the default
   profile, which `backend` addresses by compose-network name over HTTP, is a cross-process service address by
   any reading. Two passages in one file assert incompatible things.
   The clause that is true — *"all four `*_RPC_ADDR` are now set by no compose file at all"* — verifies (0
   occurrences across the three files; control `AUTHORIZATION_ADDRESS` = 1). Only the trailing generalisation
   fails.
   tree-read: `stack-demo/platform` at `0c91421d`; the corpus file itself at `8b6d80f`.

### r23-D B4 | `corpus/architecture/platform-migration-status.md:93` (messenger row, prod-terraform clause) | REJECTED | — | `messenger/terraform/main.tf:29` `= 0` in an otherwise-intact 121-line module

   evidence: I confirmed every measurement the seat reports. At the demo pin `fa47850d`,
   `terraform/main.tf` is **111 lines**, `:29` is `container_definitions = <<EOF`, `:19-28` are ordinary module
   arguments (`service_cpu`, `service_memory`, `health_check_path`, `ecs_cluster_id`, `vpc_id`,
   `vpc_cidr_block`, `private_subnets_ids`, `service_discovery_namespace_id`, `monitoring_sns_topic_arn`), and
   `service_desired_count` sits at `:19` with the value **`1`**. At `origin/main e9421c68` the file is
   **121 lines**, `:29` is `service_desired_count = 0`, `:19-25` is exactly the quoted comment block (down to
   *"Follows the cms precedent (cms v0.255.2 terraform/main.tf: service_desired_count = 0)"*), and `:27-28` is
   verbatim *"The image and task definition stay declared: this is the rollback path."* **Every anchor,
   including the whole-file line count, is exact at `origin/main` and none is exact at the pin.**
   The measurements are right; the conclusion drawn from them is not. This clause is a claim about
   **production ECS state**, and a production claim is not settled by a demo's build pin — `messenger` is not
   cloned by `make init` since `838d907`, nothing local builds from it, and a checkout 7 commits behind is a
   stale artifact rather than evidence about prod. The tree that settles it is the repo's `origin/main`, where
   the claim is true in full. The file's own header (`:17-19`) declares exactly that basis — *"Re-measured
   2026-08-05 against platform origin HEAD `0c91421` … `app` @ `2035f9a`"* — and **every explicit repo pin in
   the file is that repo's `origin/main`** (`app` `2035f9a`, `storage` `9f8cb53`, `jobsimulation` `82cb66e`),
   which is strong evidence about the tree the author measured and the one a reader is directed to.
   This is the mirror image of the ref-discipline class rather than a member of it — an unpinned claim graded
   against an *older* tree — which is what the `wrong-tree` label is for.
   (Not booked by the seat, so not adjudicated, but recorded for the repairer: at `app origin/main 2035f9a`,
   `knowledge/service-dependencies.md:58-59` says messenger's ECS service *and its terraform module* are
   **gone** and *"all three rollback paths are closed"*, while `messenger`'s own repo at `origin/main` still
   declares the module and calls it the rollback path. That tension is between two platform sources, not a
   corpus defect, and it is a different predicate from the one booked.)
   tree-read: `stack-demo/messenger` at both `fa47850d` and `origin/main e9421c68`; `stack-demo/app` at
   `origin/main 2035f9a`.
   class: **wrong-tree** — a production-terraform claim graded against the demo build pin instead of the
   `origin/main` that settles prod state, where all four anchors and the line count are exact.

---

## DEDUPLICATION

Nine bookings collapse onto **five distinct predicates**.

**P1 — "`AI_READINESS_URL` is declared at `packages/core-js/src/constants/urls.ts:52`."**
Anchors: `r23-B B1` and `r24-B B1`, both at `corpus/services/ai-readiness.md:305`.
**One anchor, two readings.** A true duplicate — same seat, same corpus site, same predicate, same refutation,
reached independently by both readings. This is the expected seat-B overlap and it is the cleanest signal in
the set: the one booking both readings of seat B agree is a blocker.

**P2 — "Production terraform still names `http://backend.internal.anthropos:8081`."**
Anchor booked: `r23-B B2` at `corpus/services/cms.md:196` (the identical sentence also stands at `cms.md:55`,
inside the same file's fold banner — one repair must take both).
Not collapsed with anything, but worth flagging for the repairer: **the same sentence appears at two further
corpus sites outside this predicate's booked anchor** — `r24-B` books it at `cms.md:196` as a MINOR rather
than a blocker (severity disagreement between the two readings of seat B, not a second predicate), and
`r23-D` books it at `corpus/services/jobsimulation.md:49-50` as a MINOR, noting a third instance at
`corpus/services/backend.md:241`. Only the blocker enters the count; the predicate has **at least four corpus
anchors** and a claim-scoped repair will leave three standing.

**P3 — "`interviewQuestions` sits in the FE type at `apps/web/src/hooks/useAIReadiness.ts:326`."**
Anchor: `r23-B B3` at `corpus/services/ai-readiness.md:595`. Booked MINOR by `r24-B` — severity disagreement
within seat B, same predicate. **Not collapsed with P1** despite both being next-web line anchors: different
file, different construct, and P1 is wrong at *every* ref while P3 is wrong only at the demo pin.

**P4 — the storage HAZARD block's app boot-guard anchors (`main.go:518-523`, `:529-535`, `env_guards.go:37-44`).**
Anchor: `r23-B B4` at `corpus/services/storage.md:73-75`. **Not collapsed with P3** — same *mechanical class*
(an unpinned anchor resolving only at `origin/main`, which `r23-B` itself identifies as the through-line of
three of its four blockers), but a different file, a different claim, and a different failure mode: here one
cited **file does not exist at all** at the grading ref, not merely a shifted line.

**P5 — "sentinel/`AUTHORIZATION_ADDRESS` is the only service address compose sets / the only cross-process
edge or hop a local stack has."**
**Three anchors, three corpus files, one predicate:**
- `r23-D B1` — `corpus/services/sentinel.md:85`
- `r23-D B2` — `corpus/services/jobsimulation.md:145-146`
- `r23-D B3` — `corpus/architecture/platform-migration-status.md:93`
One refutation kills all three (`GOTENBERG_URL=http://gotenberg:3200` on `backend` at `docker-compose.yml:57`,
`gotenberg` declared at `:170-171` in the default `core` profile at `:183`). All three are upheld as separate
blockers because each is an independently false sentence at its own anchor, and the corpus already carries the
*correct* qualified form at `architecture_overview.md:321` — proof that a claim-scoped repair here has
historically fixed one site and left the others. The repairer should treat P5 as **one edit propagated to
three files, with a fourth site (`architecture_overview.md:321`) as the model wording**.

`r23-D B4` shares a corpus **line** with `r23-D B3` (both sit in the single long `messenger` table row at
`platform-migration-status.md:93`) but is a **different clause and a different predicate** — prod ECS state
vs. local compose addressing. They are deliberately not collapsed. B4 is rejected on its own terms.

---

## Summary

BOOKED=9 UPHELD=8 REJECTED=1 IN-SCOPE-UPHELD-BLOCKERS=8 DISTINCT-PREDICATES=5 WRONG-TREE-REJECTIONS=1

SEAT-D-ONLY: BOOKED=4 UPHELD=3 REJECTED=1 IN-SCOPE-UPHELD-BLOCKERS=3 DISTINCT-PREDICATES=1 WRONG-TREE-REJECTIONS=1

(Seat-B-only, for completeness: BOOKED=5 UPHELD=5 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=5 DISTINCT-PREDICATES=4
WRONG-TREE-REJECTIONS=0 — of which one predicate is the `r23`/`r24` duplicate at `ai-readiness.md:305`.)

All 8 upheld blockers are IN-SCOPE: 6 in `corpus/services/**` (`ai-readiness.md` ×3 incl. the duplicate,
`cms.md`, `storage.md`, `sentinel.md`, `jobsimulation.md` — 7 counting by file-hit) and 1 in
`corpus/architecture/**` (`platform-migration-status.md`). Zero out-of-scope.

**Ref-discipline rejections this reading: 0.** The class did not fire once across these nine bookings, and the
reason is structural rather than lucky: no seat here booked a pinned, past-tense, or dated claim because newer
evidence contradicted it. Every ref-shaped booking in this group is the *inverse* — an **unpinned** claim that
resolves only at a **newer** ref than the tree it grades against. That inverse is real (P3, P4) and it is also
where the one rejection lives (D B4), separated from the upheld cases by the subject test rather than by the
direction of the drift.
