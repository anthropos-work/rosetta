# iter-101 UPHELD in-scope blocker anchors — the extracted union

**Count landed: 24 rows / 22 distinct predicates.** This matches the adjudication's `N = 24` / 22 exactly.
No padding, no trimming; the derivation is reproduced in `## Notes` so the arithmetic can be re-checked
independently.

Sources: `iter-101/adjudication.md` + `iter-101/verdicts/adj-{1,2,3,4}.md`, read in full.
Corpus state graded: `8b6d80f` (= `8f04d3a` under `corpus/`). Platform ground truth: platform `0c91421d` ·
app `b948604f` (demo pin) / `2035f9a4` (origin/main) · next-web-app `bb3313bc` · ant-academy `9c3843cd` ·
`stack-demo/rosetta-extensions` `ab81527a`.

This file is an EXTRACTION. Nothing here is repaired, re-measured, or re-graded.

| # | predicate (short name) | the false claim (VERBATIM, as the adjudicator quoted it) | anchor (file:line) | what is true (verbatim from the adjudicator's verdict) | adjudicator | matched in both readings? |
|---|---|---|---|---|---|---|
| 1 | `frontend-rest-call-count` | "29 direct REST/SSE calls" — the sentence's "grammatical subject is unambiguous: *"there are direct REST/SSE calls, **29 of them**"*", with the four `packages/core-js` clients sized at *"12 sites between them"* | `corpus/architecture/frontend_architecture.md:39` | "Measured over those same 21 files at `bb3313bc`, `fetch(` / `new EventSource(` totals **43**." The four core-js clients "carry **3 mentions each** and **25 outbound calls between them**"; "The counted quantity is not even monotone in the claimed one." "it undercounts that surface again, in the same direction, by ~33 %." | adj-1 (r23-A B1) | no — reading #24 **cleared** this same anchor; "The clearance does not survive; the booking does." |
| 2 | `directus-repoint-target-count` | "`backend` is the only consumer left and per-service re-point tooling has **ONE target, not two** — see the ⚠️ under *Architecture* below." | `corpus/architecture/external_services.md:136` (refuted by `:206` / `:210` in the same file) | "`:206` (bolded headline) — *"**The `--local-content` re-point targets BOTH `cms` and `backend`.**"*, restated at `:210` as the tuple *"`("cms", "backend")`"*." Ground truth: "`stack-injection/gen_injected_override.py:84` is `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` (two members)". "Upheld under the self-contradiction rule." | adj-1 (r24-A B1) | no |
| 3 | `migrations-most-recent-may-2026` | "the most recent set of migrations (May 2026)" | `corpus/services/backend.md:301` | "the newest six are `20260724132049_cms_data_model.sql`, `20260724164346_ai_readiness_freeze_how_we_measure.sql`, `20260728103254_ai_readiness_snapshot_frozen_matched_sources.sql`, `20260729133514.sql`, `20260731131307.sql`, `20260731154527_academy_chapter_progress_completed_at.sql` — i.e. **2026-07-31**, two months past the claim". "None touches *"simulation-type definitions and content JSON defaults"*." "A reader trusting `:301` concludes the schema has been static since May." | adj-1 (r23-C B2) | no |
| 4 | `app-stream-set-cardinality` | "`:33` states the set as a bolded four-member equality and then makes the cardinality explicit and load-bearing (*"NOT a fifth member"*)" | `corpus/services/backend.md:33-34` | "So `app` is **both producer and consumer of the `backend` stream**, and the producer+consumer set is **five** — `backend`, `skillpath`, `jobsimulation`, `cms`, `ai_usage`; `skiller` would be a **sixth**, not a fifth." "The same file refutes it 231 lines later at `:264`" — and `:34` "cites `:264` as its own authority." | adj-1 (r23-C B3) | no (booked MINOR by r24-C) |
| 5 | `eight-folded-rpc-mux` | "All of their Connect-RPC surfaces are served on `app`'s single RPC mux" | `corpus/services/backend.md:29-30` | "So of the eight, **skillpath, roadrunner, storage, messenger and customerio-sync have no handler on the mux at all** — only skiller, jobsimulation and cms do. `SkillPathSessionService` = **0** occurrences in Go source at both refs." "The bullet's second clause (*"nothing outside the process calls them"*) is correct, which is what lends the false first clause its authority." | adj-1 (r24-C B1) | no |
| 6 | `ai-readiness-url-line` | "`AI_READINESS_URL` is declared at `urls.ts:52`" | `corpus/services/ai-readiness.md:305` | "`:50` is `export const AI_READINESS_URL = '/ai-readiness';`. **`:52` is `export const ORGANIZATION_FEEDBACK_URL = '/enterprise/organization-feedback';`**". Across 25 commits of that file's history the constant "has **never** been at 52 at any ref reachable from this clone." "an anchor-existence check passes it and a reader following it lands on the org-feedback route." | adj-2 (r23-B B1 + r24-B B1) | **YES** |
| 7 | `prod-terraform-8081` | "production terraform still names `http://backend.internal.anthropos:8081`" | `corpus/services/cms.md:196` (identical sentence also at `cms.md:55`) | "the `:8081` form appears in **exactly one file, a markdown KB page** — `stack-demo/app/knowledge/service-dependencies.md:46`". "Raw filesystem grep over **all 59 `.tf` files** in the workspace … **zero** name the literal". Newest prod truth: "*"**There are no external callers of app's RPC mux left.**"*". "the sentence is **present tense (`still names`) and names no ref**, so it claims currency rather than a date." | adj-2 (r23-B B2) | no (booked MINOR by r24-B) |
| 8 | `interview-questions-fe-line` | "`interviewQuestions` is in the FE type at `useAIReadiness.ts:326`" | `corpus/services/ai-readiness.md:595` | "`interviewQuestions: number;` is at **`:274`**, inside `export interface AIReadinessCycleTotals` (`:271-278`). **`:326` is `headers: {`**, inside `const res = await fetch(url.toString(), { ...init, headers: { Authorization: … } })` at `:324-331`." "it fails as a resolving line naming an unrelated construct." | adj-2 (r23-B B3) | no (booked MINOR by r24-B) |
| 9 | `storage-boot-guard-anchors` | **[quoted form deliberately withdrawn — the defect is an ABSENCE, not a false string.]** The HAZARD block's three `app` anchors (`main.go:518-523` / `:529-535` / `env_guards.go:37-44`) were cited **with no ref**; they name the right constructs at `ad9f3c49` and `2035f9a4` and fail only at the demo's *former* pin `b948604f`, where `env_guards.go` does not exist. Repaired by **pinning**, not by moving anchors — so the repaired text legitimately still contains the same anchor strings. A verbatim-quote fence cannot represent a missing-ref defect, and quoting the surrounding true headline — the HAZARD line about neither manager using local FS, and the nothing-warns clause, both TRUE and byte-unchanged — would fence a true sentence. Recorded in `claim_ledger`'s disclosed *"quoted no refuted form"* bucket rather than fenced wrongly. | `corpus/services/storage.md:73-75` (HAZARD block `:60-80`) | "`git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — `app/env_guards.go` does not exist … The entire "nothing warns" mechanism is uncheckable on the pinned clone." "`main.go:518-523` @ `b948604f` is the **public-storage clients** block …; `:529-535` is the **academy asset uploader** block. Both resolve; both name the wrong construct." Also: "`git ls-tree b948604f -- internal/storage/` is empty". | adj-2 (r23-B B4) | no |
| 10 | `sentinel-only-cross-process-edge` | "`AUTHORIZATION_ADDRESS` is the only service address compose sets; backend→sentinel the one cross-process edge" | `corpus/services/sentinel.md:85` | "`backend`'s `environment:` block carries four more cross-process service addresses: **`docker-compose.yml:57` `GOTENBERG_URL=http://gotenberg:3200`**, `:66` `REDIS_ADDR=redis:6379`, `:93` `SUPABASE_DB_CONN=…`, `:94` `COPILOT_DB_CONN=…`." "The `*_RPC_ADDR` half of the sentence is **true and verified** … It is the generalisation from *"the only RPC address"* to *"the only service address / the one cross-process edge"* that breaks." Correct wording already exists at `architecture_overview.md:321`. | adj-2 (r23-D B1) | no |
| 11 | `sentinel-only-cross-process-edge` | "sentinel is the only cross-process hop and the only service address backend's compose entry carries" | `corpus/services/jobsimulation.md:145-146` | "identical re-derivation to D B1, at the same platform ref. The block cites `docker-compose.yml:48` and names no platform sha; the checkout `0c91421d` is level with `origin/main`, so **there is no tree the claim could be true at**." "`backend → gotenberg` is a second live cross-process hop in the default `core` profile". Booked separately because "a claim-scoped repair that fixes one and not the other leaves this one standing." | adj-2 (r23-D B2) | no |
| 12 | `sentinel-only-cross-process-edge` | "the only cross-process service address left in a local stack is `AUTHORIZATION_ADDRESS`" | `corpus/architecture/platform-migration-status.md:93` (messenger row, **RPC-edge clause**) | "**the same file, twelve rows down, states the counter-evidence.** `platform-migration-status.md:105` (the `gotenberg` row) reads *"third-party image, `docker-compose.yml:170-171` … **default `core` profile**"* and grades its `fresh local stack` column **live-standalone**." "The clause that is true — *"all four `*_RPC_ADDR` are now set by no compose file at all"* — verifies … Only the trailing generalisation fails." | adj-2 (r23-D B3) | no |
| 13 | `archive-note-row-anchors` | the note "says the flat form was published at `:137`/`:138` "two rows above `:139`, a cell retracting exactly that predicate"" | `corpus/architecture/service_taxonomy.md:130-133` (r24 cites the same sentence as `:131-133`) | "Every one of the three anchors is shifted by exactly **+2**: the flat form is at `:139`/`:140`, two rows above **`:141`**, and `:139` is one of the two cells *asserting* the predicate, not retracting it." "`:137` = Chronos and `:138` = Intelligence … each **contains no archive assertion of any kind**." Regression proven: "At `a229f8d^` … The note was **exactly correct there**." | adj-3 (r23-E B1 + r24-E B1) | **YES** |
| 14 | `org-mixin-user-count` | "only 30 use OrganizationMixin{}" | `corpus/architecture/security_compliance.md:67-68` | "I then opened the 30th: `internal/data/ent/schema/user_resource.go:22` … reads `// OrganizationMixin{},  // We need to work on this` — commented out, not compiled. So **29 schemas use it; 30 mention it**, and the sentence's predicate is *"use"*." "The fence's opening sentence commits the exact error the fence exists to forbid." | adj-3 (r23-E B2 + r24-E B3) | **YES** |
| 15 | `base-services-floor-cardinality` | "Base services (no profile, always on with any `make up`)" heading "that exact predicate and lists **two** bullets" | `corpus/architecture/service_taxonomy.md:109-111` | "the set satisfying *"no profile, always on with any `make up`"* is `{postgresql, redis, sentinel}`, cardinality **3**." "Under `:109`'s enumeration the file's own arithmetic does not close: backend + gotenberg + 2 = **four**, against the **five** stated at `:68` and `:489`." The file says three at `:68`, `:465`, `:489`; "Root `CLAUDE.md` says three as well." | adj-3 (r23-E B4) | no |
| 16 | `localeswitch-surface-set` | "LocaleSwitch claimed on `/library`+`/free`" — the "i.e. `/library` + `/free`" gloss | `corpus/services/ant-academy.md:229-230` | "`LocaleSwitch` is mounted at exactly one site, `PublicHeader.jsx:20` — the "only in the public-storefront header" half is TRUE; the "i.e. `/library` + `/free`" gloss is false for `/free`. A visitor at `/free` is redirected to `/`, which renders `AcademyClient` → `TopBar` → `LanguageSelector` — the 7-locale dropdown, the exact component this bullet exists to distinguish from `LocaleSwitch`." | adj-4 (r23-F B1) | no |
| 17 | `topbar-surface-set` | "TopBar's surface set is the 5 named routes" | `corpus/services/ant-academy.md:234` | "**`AcademyClient` serves three routes, not one** … Measured surface set = `/`, `/courses`, `/courses/[slug]`, `/chapters/[slug]`, `/latest`, `/bookmarks`, `/my-activity` = **7**. The list names 5." "the two omitted routes are precisely the demo's landing routes". | adj-4 (r23-F B2) | no |
| 18 | `skiller-stream-file-count` | "SKILLER_STREAM: 6 Go occurrences across 4 files" | `corpus/architecture/dependency_map.md:59` | "`git grep -n SKILLER_STREAM 2035f9a4 -- '*.go'` → **6 lines** … `git grep -ln … -- '*.go'` → **3 files**. Dropping the pathspec → **6 files** … So the occurrence count 6 is exactly right and the file count 4 matches neither available scope." Positive control: "`CMS_STREAM … ` → **4** files, so the pipeline does distinguish 3 from 4." | adj-4 (r23-F B3 + r24-F B1) | **YES** |
| 19 | `no-line-number-resolves` | "**not one of those line numbers resolves**" | `corpus/architecture/dependency_map.md:59` (same cell — a **second, distinct** predicate) | "The seventh **holds**: `internal/roles/roles.go:791` = `func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {` — byte-identical to its value at `b948604f`. So the true split is 6 drift / 1 holds, and the cell's bolded universal … is falsified by a member of its own enumerated set." | adj-4 (r24-F B2) | no |
| 20 | `vendor-selection-path` | vendor selection sited at "each consumer's own `internal/ai/ai.go`" | `corpus/architecture/shared_libraries.md:130-131` | "**false at `b948604f`** — `git cat-file -e b948604f:internal/ai/ai.go` fails … no `internal/ai/` package at all. **Also false at `2035f9a4`** … is **21 lines** declaring `type AI interface {…}` … It contains **no** Azure client, **no** `flag_use_azure_us`, **no** 429 handling." The mechanics live at "`internal/jobsimulation/ai/ai.go:267`, `:344` and `internal/skillerai/ai.go:347`". | adj-4 (r23-F B4) | no |
| 21 | `empty-catalog-view-shape` | "`emptyCatalogView()` asserted by `=` as a 3-key literal" | `corpus/services/ant-academy.md:134-135` | "`code/src/lib/serverTenant.js:115-117`: `return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }`. Five keys, not three. `PUBLIC_BUNDLES` (`code/ucourses/catalog.js:956`) is a populated exported array of bundle objects — not empty". | adj-4 (r23-F B5) — flagged **"the weakest of my ten upholds"**; r24-F graded the identical fact a MINOR | no |
| 22 | `recruiter-scoreboard-app` | "**NB the demo's recruiter candidate-comparison scoreboard is an `is_hiring` ORG-TYPE surface in the dockerized `apps/web`** (`/enterprise/activity-dashboard`), **not** this Vercel-only app" | `corpus/services/next-web-app.md:32` (supported at `:14`) | "`hiring.md:53` *"**not in `apps/web`**. M222's dockerized-`apps/web` … premise was **falsified at M224**"*; `hiring.md:352` *"**M224 proved the render — and it does NOT land in `apps/web`.**"*". Ground truth: "A genuine hiring org's recruiter is ejected out of `apps/web`." "The drawer that renders the comparison is `apps/hiring/src/components/containers/InsightsByMembersContainer.tsx:359`." | adj-4 (r23-G B1) | no |
| 23 | `comparable-cohort-key` | "Comparable-cohort key given as `jobsimulation_id`, a column of the mirror the same document says was dropped" | `corpus/services/hiring.md:160` **and** `corpus/services/hiring.md:198` (one predicate, two anchors, one booking) | "**the live table has no such column.** `terraform/migrations/20260722104506.sql:2-26` is the full DDL of `job_simulation_sessions` — 23 columns, `sim_id` at `:7`, and **no `jobsimulation_id`**." "**it was the dropped mirror's column** … the table dies at `20260729133514.sql:62`". "**the real key is `sim_id` + `organization_id`**". | adj-4 (r23-G B2) | no |
| 24 | `hiring-twin-citation` | "Cites `service_taxonomy.md:52` as a corroborating twin" — the sentence asserts the twins *"already said"* it | `corpus/services/hiring.md:38` | "`corpus/architecture/service_taxonomy.md:52` reads, in full, `> [dependency_map.md](./dependency_map.md)'s content-generation flow, which had it right all along.` — the closing line of the `:44-52` blockquote whose entire subject is the **direction of the Studio-Desk → Backend → Studio-Room generation edge** … It says nothing about the `jobsimulation` schema, `public`, migrations, or M710". The one relevant hit is at **`:62`**. | adj-4 (r24-G B1) | no |

## Predicate roll-up

22 distinct predicates. "anchors" = the corpus `file:line` sites the predicate is false at, as booked and upheld.

| predicate | anchors | anchor list |
|---|---|---|
| `sentinel-only-cross-process-edge` | **3** | `sentinel.md:85` · `jobsimulation.md:145-146` · `platform-migration-status.md:93` |
| `comparable-cohort-key` | **2** | `hiring.md:160` · `hiring.md:198` |
| `frontend-rest-call-count` | 1 | `frontend_architecture.md:39` |
| `directus-repoint-target-count` | 1 | `external_services.md:136` |
| `migrations-most-recent-may-2026` | 1 | `backend.md:301` |
| `app-stream-set-cardinality` | 1 | `backend.md:33-34` |
| `eight-folded-rpc-mux` | 1 | `backend.md:29-30` |
| `ai-readiness-url-line` | 1 | `ai-readiness.md:305` |
| `prod-terraform-8081` | 1 | `cms.md:196` (see Notes — the sentence has ≥3 further corpus sites, booked MINOR) |
| `interview-questions-fe-line` | 1 | `ai-readiness.md:595` |
| `storage-boot-guard-anchors` | 1 | `storage.md:73-75` |
| `archive-note-row-anchors` | 1 | `service_taxonomy.md:130-133` |
| `org-mixin-user-count` | 1 | `security_compliance.md:67-68` |
| `base-services-floor-cardinality` | 1 | `service_taxonomy.md:109-111` |
| `localeswitch-surface-set` | 1 | `ant-academy.md:229-230` |
| `topbar-surface-set` | 1 | `ant-academy.md:234` |
| `skiller-stream-file-count` | 1 | `dependency_map.md:59` |
| `no-line-number-resolves` | 1 | `dependency_map.md:59` (**same site, different predicate**) |
| `vendor-selection-path` | 1 | `shared_libraries.md:130-131` |
| `empty-catalog-view-shape` | 1 | `ant-academy.md:134-135` |
| `recruiter-scoreboard-app` | 1 | `next-web-app.md:32` |
| `hiring-twin-citation` | 1 | `hiring.md:38` |

## Notes

**1. The count reconciles exactly at 24, and it does so two independent ways.**
- *As the adjudication's union arithmetic:* `n₁ = 20` (reading #23, 7 seats) + `n₂ = 8` (reading #24, 6 seats)
  − `m = 4` matched = **24**. Reading #23's 20: adj-1 ×3 (rows 1, 3, 4) · adj-2 ×7 (rows 6–12) · adj-3 ×3
  (rows 13–15) · adj-4 ×7 (rows 16, 17, 18, 20, 21, 22, 23). Reading #24's 8: rows 2, 5, 6, 13, 14, 18, 19, 24
  — of which rows 6, 13, 14, 18 are the four matched.
- *As distinct `file:line` sites:* also **24** — the table's 24 rows resolve to 24 sites once
  `dependency_map.md:59` (rows 18+19) collapses to one site and `hiring.md:160`/`:198` (row 23) expands to two.
  The two accounting bases differ in composition and coincide in total; neither was fitted to the other.

**Predicate arithmetic:** 24 − 2 (the `sentinel-only-cross-process-edge` collapse, 3 anchors → 1 predicate)
= **22**. That is the adjudication's 22.

**2. `sentinel-only-cross-process-edge` is ONE repair, not three.** adjudication.md §2 states it directly:
*"That is one repair propagated to three files with a fourth as the model, not three fixes."* Refuted by a
single line — `GOTENBERG_URL=http://gotenberg:3200` on `backend` at `docker-compose.yml:57`, with `gotenberg`
declared at `:170-171` in the default `core` profile (`:183`) and reached over real HTTP at
`app/internal/converter/gotenberg.go:31`. The `*_RPC_ADDR`-is-zero half of every one of the three sentences is
**true**; only the generalisation breaks. The corpus already carries the correct qualified wording at
`architecture_overview.md:321` — *"the only cross-process **RPC** edge out of backend on a core stack"* — which
is the model for the repair. It contradicts itself at `gotenberg.md:50`, `dependency_map.md:103` and
`platform-migration-status.md:105`.

> **⚠️ This is also `CLAUDE.md`'s claim, VERBATIM, in this repo's own root instructions** (adjudication.md §2:
> *"It is also `CLAUDE.md`'s claim, verbatim, in this repo's own root instructions."*). The repo-root
> instructions file is **outside** the `corpus/services/**` + `corpus/architecture/**` scope this union
> enumerates, so it is **not** a 25th row — but a predicate-scoped repair that leaves `CLAUDE.md` asserting the
> refuted generalisation has not finished the job. Handle it as part of the same edit.

**3. `dependency_map.md:59` carries TWO independent upheld defects** — a count (`skiller-stream-file-count`)
and a self-contradicting universal (`no-line-number-resolves`). adj-4 deliberately did not collapse them:
*"One markdown cell, two independent defects."* Both must be repaired; fixing the file count alone leaves the
bolded universal standing.

**4. `service_taxonomy.md:130-133` is a REPAIR-INDUCED defect** — iter-100's own two-line parenthetical
(`a229f8d`) pushed the table down two rows and left the line numbers unmoved. adj-3 verified the note was
*"exactly correct"* at `a229f8d^`. Both readings found it independently; it is one of the four matched
predicates.

**5. `prod-terraform-8081` under-counts its own blast radius.** adj-2's DEDUPLICATION records that the same
sentence stands at **at least four corpus anchors**: the booked blocker `cms.md:196`, plus `cms.md:55` (same
file's fold banner), `jobsimulation.md:49-50` and `backend.md:241` (both booked MINOR, so out of this union).
*"Only the blocker enters the count; the predicate has at least four corpus anchors and a claim-scoped repair
will leave three standing."* Repair by predicate.

**6. Excluded, and why.** The 8 REJECTED bookings are out: `backend.md:19`, `skiller.md:19`,
`ai_architecture.md:35`, `ai_architecture.md:141`, `security_compliance.md:185` (5 mis-read); `chronos.md:27`,
`ai-labs.md:75` (2 wrong-convention); and seat D's messenger-row **prod-terraform clause** at
`platform-migration-status.md:93` (1 wrong-tree). **Note the collision:** that wrong-tree rejection shares the
corpus *line* `platform-migration-status.md:93` with row 12, which IS upheld. They are different clauses of one
long table row — prod ECS state (rejected) vs local compose addressing (upheld). Do not let the shared line
number merge them. There were **zero** out-of-scope upheld bookings.

**7. Row 21 (`empty-catalog-view-shape`) is disclosed as marginal by its own adjudicator** — *"If this
milestone applies any materiality floor above 'literally false', this is the one booking that falls below it."*
It is included because the milestone's grading bar is literal falsity, and the sibling reading's MINOR grade is
recorded rather than smoothed.

**8. Row 1 (`frontend-rest-call-count`) is a cross-reading verdict inversion, not a duplicate.** Reading #24
opened the identical anchor and **cleared** it by re-deriving the env-var occurrence count (29/21, exact).
adj-1: *"That is rule 4 in its pure form — the arithmetic is right and the predicate is wrong. The clearance
does not survive; the booking does."*

**9. `N = 24` is a FLOOR.** It is a **13-seat** union (`r24-D` was never produced and was deliberately not
re-run); a 14th seat can only add. adjudication.md's cross-reading estimator puts the residual inside clause
5's scope *"on the order of ~100, not ~45"*.

**10. Binding repair conditions**, inherited from iter-76 and restated in adjudication.md: routed as
**`FIX-M257x-iter101-read-union`** — **repair by PREDICATE, not by anchor**, and **re-read after**.
