# adj-4 — verdicts for seats F (r23, r24) and G (r23, r24)

**Adjudicated:** 10 blockers — r23-F (5), r24-F (2), r23-G (2), r24-G (1).

**Ground truth re-verified at this reading's open** (`git rev-parse --short=8 HEAD` per clone):
platform `0c91421d` · app `b948604f` · cms `ca50c817` · next-web-app `bb3313bc` · sentinel `88bc5592` ·
storage `4ce8ece5` · messenger `fa47850d` · graphql-wundergraph `60c229f3` · roadrunner `87d8d443` ·
jobsimulation `462343b0` · studio-desk `14a5442a` · ant-academy `9c3843cd` ·
`stack-demo/rosetta-extensions` `ab81527a` (pinned consumption clone) ·
`.agentspace/rosetta-extensions` `09d06070` (authoring copy).

**Corpus state.** Corpus HEAD is `8b6d80f5`. `git diff --quiet 8f04d3a..HEAD -- corpus/` is clean, so the
corpus I read is byte-identical to the one the seats read. (r23-G's header states corpus HEAD `1937e1f`;
that is iter-57, 2026-08-03 — a mis-stated header, not a different read. Every anchor the four seats quote
reproduced verbatim at the state I read, so nothing here turns on it.)

**Tooling-tree hazard: not engaged.** None of these 10 bookings turns on a `rosetta-extensions` claim, so
the two-clone split never had to be resolved. All 10 settle in a platform clone or in the corpus itself.

---

## Verdicts

### r23-F B1 | `corpus/services/ant-academy.md:229-230` | UPHELD | IN-SCOPE | LocaleSwitch claimed on `/library`+`/free`; `/free` renders no header at all

   evidence: `stack-demo/ant-academy` `code/app/(public)/free/page.jsx` @ `9c3843cd` — the default export is
   `export default function Page() { redirect('/?tier=free') }`, with an in-file comment at `:5-11` stating
   `/free` "now redirects to the real catalog filtered to the free tier". `code/app/(public)/layout.jsx` is a
   pass-through (`export default function PublicLayout({ children }) { return children }`) that renders
   nothing. `git grep -ln PublicHeader 9c3843cd -- 'code/'` returns exactly two paths —
   `code/app/(public)/library/page.jsx` and the component's own file. The `(public)` tree holds exactly three
   files (`free/page.jsx`, `layout.jsx`, `library/page.jsx`), so the enumeration is complete. `LocaleSwitch`
   is mounted at exactly one site, `PublicHeader.jsx:20` — the "only in the public-storefront header" half is
   TRUE; the "i.e. `/library` + `/free`" gloss is false for `/free`. A visitor at `/free` is redirected to
   `/`, which renders `AcademyClient` → `TopBar` → `LanguageSelector` — the 7-locale dropdown, the exact
   component this bullet exists to distinguish from `LocaleSwitch`.
   tree-read: `stack-demo/ant-academy @ 9c3843cd`
   note: the likely origin is `PublicHeader.jsx`'s own stale header comment, *"Server-rendered header for the
   public surface (/library, /free)"* — the corpus copied a source comment the source has outgrown. That
   explains the error; it does not make the claim true.

### r23-F B2 | `corpus/services/ant-academy.md:234` | UPHELD | IN-SCOPE | TopBar's surface set is 7 routes; the list names 5, dropping both `/courses` routes

   evidence: `stack-demo/ant-academy @ 9c3843cd`. `TopBar` is imported by five production views —
   `src/views/academy/AcademyClient.jsx:28` (rendered `:1906`), `bookmarks/BookmarksClient.jsx:28` (`:508`),
   `course/CourseClient.jsx:58` (`:2091`, `:2141`), `latest/LatestClient.jsx:7` (`:128`),
   `my-activity/MyActivityClient.jsx:5` (`:161`). **`AcademyClient` serves three routes, not one**:
   `app/(authed)/page.jsx:4`, `app/(authed)/courses/page.jsx:5` (renders `<AcademyClient` at `:92`) and
   `app/(authed)/courses/[slug]/page.jsx:8` (renders it at `:219`). Measured surface set = `/`, `/courses`,
   `/courses/[slug]`, `/chapters/[slug]`, `/latest`, `/bookmarks`, `/my-activity` = **7**. The list names 5.
   The doc is *correct* about what it excludes — `my-certificates/MyCertificatesClient` genuinely does not
   import TopBar — so the omission is not a scoping convention, it is two missing members.
   tree-read: `stack-demo/ant-academy @ 9c3843cd`
   grading basis (stated, because this is the weaker of my two ant-academy upholds): every member listed is
   true; the defect is an incomplete SET presented without hedge as the answer to "where is the app-shell
   header". I uphold on materiality supplied by the same document: `:292` instructs *"Link a course CTA at
   `/courses/<slug>`, never `/library/<slug>`"* — so the two omitted routes are precisely the demo's landing
   routes, and a reader triaging "the switcher shows no menu" there would conclude from this list that TopBar
   is absent. It is measurably present.

### r23-F B3 | `corpus/architecture/dependency_map.md:59` | UPHELD | IN-SCOPE | "SKILLER_STREAM: 6 Go occurrences across 4 files" — measured 3 Go files, 6 repo-wide; 4 is unsupportable at either scope

   evidence: measured at `app 2035f9a4`, the ref the cell itself names, with the `'*.go'` pathspec the
   sentence scopes to. `git grep -n SKILLER_STREAM 2035f9a4 -- '*.go'` → **6 lines**: `main.go:1532`,
   `main.go:1537`, `subscriber_merge_test.go:907`, `:1015`, `:1038`, `subscriber_wiring.go:165`.
   `git grep -ln … -- '*.go'` → **3 files**. Dropping the pathspec → **6 files** (adds `terraform/main.tf`,
   `knowledge/skiller-domain.md`, `knowledge/plan/releases/09.00-support-in-app/m905-subscriber-merge/
   decisions.md`). So the occurrence count 6 is exactly right and the file count 4 matches neither available
   scope. Three-instrument check per the brief: `git grep -a` (binary-as-text, defeating the NUL-skip) still
   returns 3 Go files and 6 repo-wide; the nested untracked repo `stack-demo/app/studio @ aeec036` returns 0
   for the predicate at its own ref; and `git grep` at a ref is unaffected by `.gitignore`, which is the
   mechanism that would hide a tracked file from this shell's `ugrep`. Positive control in the same pass:
   `git grep -ln CMS_STREAM 2035f9a4 -- '*.go'` → **4** files, so the pipeline does distinguish 3 from 4.
   tree-read: `stack-demo/app @ 2035f9a4`
   aggravating: this is a *correction* sentence — text written to repair an earlier miscount — and the two
   true numbers (6 occurrences, 6 repo-wide files) are the same digit at different places in the sentence,
   which is the most likely route by which the wrong one entered.

### r24-F B1 | `corpus/architecture/dependency_map.md:59` | UPHELD | IN-SCOPE | same predicate as r23-F B3 (SKILLER_STREAM file count)

   evidence: identical measurement, same anchor, same ref — see r23-F B3. Collapses onto one predicate.
   tree-read: `stack-demo/app @ 2035f9a4`

### r24-F B2 | `corpus/architecture/dependency_map.md:59` | UPHELD | IN-SCOPE | Cell asserts "not one of those line numbers resolves", then names one that does — at both refs

   evidence: I resolved all seven line anchors the cell enumerates, at both refs, by
   `git show <ref>:<file> | sed -n '<n>p'`.
   At `2035f9a4` six drift: `main.go:287` → `logger.Info("subsystem switches",`; `:637` → `)`; `:1039` →
   `serverContext,`; `jobsimwiring/wiring.go:127` → `asynqClient := jsworkerclient.NewClient(...)`; `:180` →
   blank; `main.go:1276` → `apiKeyManager,` (the cell's own parenthetical, correct).
   The seventh **holds**: `internal/roles/roles.go:791` = `func (r *RoleManager) SkillerSubscriber()
   *pubsub.Subscriber {` — byte-identical to its value at `b948604f`. So the true split is 6 drift / 1 holds,
   and the cell's bolded universal ("**not one of those line numbers resolves**") is falsified by a member of
   its own enumerated set — a member the same cell then names three sentences later as surviving
   (*"and `internal/roles/roles.go:791` is `SkillerSubscriber()` at each"*).
   tree-read: `stack-demo/app @ 2035f9a4` and `@ b948604f`
   grading basis: the one available defence is a narrow reading in which "those line numbers" excludes the
   parenthetical handler anchor. I rejected it: the text carries no such restriction, `roles.go:791` is
   enumerated in the immediately preceding sentence, and the later clause introduces it with a bolded
   corrective (*"**The consumer-only finding itself holds at both**"*) — which is only needed if the earlier
   universal over-reached. This is brief rule 5 with the additional strength that I can say which side is
   right: `roles.go:791` resolves. Distinct predicate from r23-F B3/r24-F B1 — same line, different claim.

### r23-F B4 | `corpus/architecture/shared_libraries.md:130-131` | UPHELD | IN-SCOPE | Vendor selection sited at "each consumer's own `internal/ai/ai.go`"; false at the live consumer at BOTH refs

   evidence: **false at `b948604f`** — `git cat-file -e b948604f:internal/ai/ai.go` fails; `git ls-tree -r`
   shows no `internal/ai/` package at all (only `internal/aireadiness/…`).
   **Also false at `2035f9a4`, and this is where I depart from the seat's evidence.** The path *does* exist at
   origin/main — but `git show 2035f9a4:internal/ai/ai.go` is **21 lines** declaring
   `type AI interface { ChatCompletion, Response, CreateEmbeddings, CreateSpeech, OCRProcess,
   AudioTranscriptions, Tokenize, GetEndpoint }` plus `type TokenEncoder` — i.e. an in-repo copy of the *`ai`
   library's own interface*, the very thing the table at `:113-117` describes. It contains **no** Azure
   client, **no** `flag_use_azure_us`, **no** 429 handling. At both refs the three quoted mechanics live
   elsewhere: `git grep -n flag_use_azure_us -- '*.go'` returns `internal/jobsimulation/ai/ai.go:267`, `:344`
   and `internal/skillerai/ai.go:347`, and `internal/jobsimulation/ai/ai.go:262-276` is `getClient`'s Azure
   arm (`client := a.azureClientEu` → PostHog `flag_use_azure_us` → `azureClientUs`).
   So the seat's booking is correct and its predicate is *stronger* than the seat argued: the sentence is
   unsupportable at the checkout ref (path absent) and at origin/main (path present, wrong contents).
   Compounding it, the same table's **Imported by** row two lines above (`:110`) restricts the live consumer
   set to *"`app` alone among the services a stack runs"* — and the corpus's own twins name the real path:
   `architecture_overview.md:292` (*"Measured at `app/internal/jobsimulation/ai/ai.go`"*) and
   `external_services.md:581` (*"all in `app/internal/jobsimulation/ai/ai.go`"*). Cross-file
   self-contradiction, both anchors opened and read.
   tree-read: `stack-demo/app @ b948604f` and `@ 2035f9a4`; corroborated against `stack-demo/jobsimulation
   @ 462343b0`, where `internal/ai/ai.go` genuinely *does* carry the mechanics (`:44`, `:65`, `:129`, `:267`)
   — that frozen repo is the only place the corpus's path is true, and it is not a repo any stack builds.

### r23-F B5 | `corpus/services/ant-academy.md:134-135` | UPHELD | IN-SCOPE | `emptyCatalogView()` asserted by `=` as a 3-key literal; source returns 5, one of them populated

   evidence: `stack-demo/ant-academy @ 9c3843cd`, `code/src/lib/serverTenant.js:115-117`:
   `return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion:
   CATALOG_VERSION }`. Five keys, not three. `PUBLIC_BUNDLES` (`code/ucourses/catalog.js:956`) is a populated
   exported array of bundle objects — not empty — and the function's own doc-comment at `:111-113` singles it
   out (*"`bundles` (PUBLIC_BUNDLES) carries no tenant metadata … so it passes through verbatim"*).
   `CATALOG_VERSION = '1.0'` (`catalog.js:31`).
   tree-read: `stack-demo/ant-academy @ 9c3843cd`
   **This is the weakest of my ten upholds and I am flagging it as such.** The conclusion the passage draws
   (`→ 0 cards`) is correct and unaffected — `bundles` is not a grid card. The sibling reading r24-F measured
   the identical fact and graded it a **MINOR** ("a truncated quotation, not a wrong mechanism"). I uphold on
   the brief's stated bar and nothing more: the passage uses `=` to assert an exact value of a named
   function, and the asserted value is not the code's value. None of the five rejection classes fits — the
   sentence carries no ref pin (not ref-discipline), the tree is right (not wrong-tree), the seat's
   measurement reproduces exactly (not mis-read), and it is not already-true. If this milestone applies any
   materiality floor above "literally false", this is the one booking that falls below it.

### r23-G B1 | `corpus/services/next-web-app.md:32` | UPHELD | IN-SCOPE | Recruiter scoreboard sited in dockerized `apps/web`; `hiring.md` retracts it three times and ground truth agrees with `hiring.md`

   evidence — **(a) intra-corpus contradiction, with the twin naming this file.** `next-web-app.md:32`:
   *"**NB the demo's recruiter candidate-comparison scoreboard is an `is_hiring` ORG-TYPE surface in the
   dockerized `apps/web`** (`/enterprise/activity-dashboard`), **not** this Vercel-only app"* (supported at
   `:14`). Against it: `hiring.md:53` *"**not in `apps/web`**. M222's dockerized-`apps/web`
   (`/enterprise/activity-dashboard`) premise was **falsified at M224**"*; `hiring.md:352` *"**M224 proved the
   render — and it does NOT land in `apps/web`.**"*; and `hiring.md:422-424`, the Cross-references entry **for
   `next-web-app.md`**: *"⚠️ M222 inferred the scoreboard was reachable in `apps/web`; M224 rendering proved
   it is not for a *genuine* hiring org — the demo serves the real `apps/hiring` as a second container."*
   One file points at the other and warns the reader its content is retracted; the other still asserts it.
   **(b) ground truth, opened independently at `next-web-app bb3313bc`.**
   `apps/web/src/context/UserStatusContext.tsx:141-174` — a `useEffect` computes `userHasAllHiringOrgs` at
   `:143-148` over `membership?.organization?.publicMetadata?.isHiring`, returns early if false, and otherwise
   sets `window.location.href = buildSwitchHandoffUrl({ targetProduct: 'hiring', … })`. A genuine hiring org's
   recruiter is ejected out of `apps/web`. `apps/web/src/hooks/useGetClerkOrganization.tsx:16-18` filters
   `!organization.publicMetadata?.isHiring`, removing hiring orgs from the workforce list. The symmetric guard
   `apps/hiring/src/context/UserStatusContext.tsx:120-148` (`userHasAllWorkforceOrgs` → `targetProduct:
   'workforce'`) keeps the recruiter *in* `apps/hiring`. The drawer that renders the comparison is
   `apps/hiring/src/components/containers/InsightsByMembersContainer.tsx:359` (`<Drawer`).
   The claim carries no ref pin and no past tense, so ref-discipline does not reach it.
   tree-read: `stack-demo/next-web-app @ bb3313bc`
   noted hesitation, and why it does not save the sentence: `apps/web` *does* still ship a byte-identical
   `InsightsByMembersContainer.tsx` and a full `/enterprise/activity-dashboard` route tree (`git ls-tree`
   confirms both apps carry it), and the client re-skin really is an `apps/web` mechanism. But the sentence's
   load-bearing clause is *where the demo's recruiter scoreboard is*, and the platform's own guard makes that
   the one app the recruiter cannot stay in.

### r23-G B2 | `corpus/services/hiring.md:160` and `corpus/services/hiring.md:198` | UPHELD | IN-SCOPE | Comparable-cohort key given as `jobsimulation_id`, a column of the mirror the same document says was dropped

   evidence, all at `app b948604f`:
   **(a) the live table has no such column.** `terraform/migrations/20260722104506.sql:2-26` is the full DDL
   of `job_simulation_sessions` — 23 columns, `sim_id` at `:7`, and **no `jobsimulation_id`**.
   `internal/data/ent/jobsimulationsession/jobsimulationsession.go:28` is `FieldSimID = "sim_id"`.
   `git grep -c jobsimulation_id b948604f -- 'internal/data/ent/schema/'` exits 1 (zero matches).
   **(b) it was the dropped mirror's column.** `terraform/migrations/20240527131926.sql:7` creates
   `"jobsimulation_id" uuid NOT NULL` on `local_jobsimulation_sessions`, indexed at `:21`; the table dies at
   `20260729133514.sql:62` (`DROP TABLE "local_jobsimulation_sessions";`) under the `:58` comment
   `-- 5. Drop the mirrors.` That is exactly the table `hiring.md:168-176` and `:262-266` spend two
   blockquotes telling the reader is gone.
   **(c) the real key is `sim_id` + `organization_id`.** `internal/organization/intelligence.go:1700-1709` —
   `m.ent.JobSimulationSession.Query().Where(jobsimulationsession.SimID(jobSimulationId), …,
   jobsimulationsession.OrganizationID(organizationID))`; the best-attempt window at `:2156-2173` is
   `ROW_NUMBER() OVER (PARTITION BY %s, %s …)` formatted with `FieldSimID, FieldOwnerID`, with the attempts
   count over the same partition.
   **(d) the document contradicts itself.** Its own minimal write-set names the right column twice —
   `hiring.md:225` (`owner_id`, `sim_id`, `sim_type`, **`token`**) and `:233` (`sim_id` `:7`). So one section
   tells a seeder author to write `sim_id` while another gives `jobsimulation_id` as the comparison key.
   Corpus-wide, `jobsimulation_id` occurs in exactly two places, both of them these two anchors.
   tree-read: `stack-demo/app @ b948604f`
   scope: one predicate, two anchors — booked as a single blocker by the seat, counted once.

### r24-G B1 | `corpus/services/hiring.md:38` | UPHELD | IN-SCOPE | Cites `service_taxonomy.md:52` as a corroborating twin; that line is about Studio-Room's `gen.py`

   evidence: `corpus/architecture/service_taxonomy.md:52` reads, in full,
   `> [`dependency_map.md`](./dependency_map.md)'s content-generation flow, which had it right all along.`
   — the closing line of the `:44-52` blockquote whose entire subject is the **direction of the Studio-Desk →
   Backend → Studio-Room generation edge** (`studio-desk/.env.example:45`, `cms_queries.graphqls:106`,
   `app/internal/cms/studio/studioManager.go:119` exec'ing `studio/gen.py`). It says nothing about the
   `jobsimulation` schema, `public`, migrations, or M710 — which is what `hiring.md:30-38` asserts the twins
   *"already said"*.
   Absence control, enumerated not sampled: `grep -nE 'M710|non-authoritative|legacy husk'` over
   `service_taxonomy.md` (496 lines) returns **one** hit, at **`:62`** — *"**Database**: PostgreSQL — **one
   schema, `public`, owned by `app`**, which is the only repo with migrations (`repos.yml:14-17`) … the `cms`,
   `jobsimulation` and `skillpath` schemas are legacy husks"* — and **nothing at `:52`**. Positive control,
   same pattern in the same invocation over `dependency_map.md`: hits at `:31` and `:78`. So the pipeline was
   working and the absence at `:52` is real. The citation is off by ten lines onto an unrelated construct.
   The sibling citation in the same sentence, `dependency_map.md:78`, **does** resolve onto a relevant
   construct (*"or directly to the **`public`** schema (the legacy `jobsimulation` schema is
   non-authoritative)"*) — which is what lets the pair survive a casual reading.
   tree-read: corpus @ `8b6d80f5` (unchanged from the seats' read since `8f04d3a`)
   grading basis: the sentence is not a bare cross-reference, it is an **assertion of corroboration**
   (*"as the twins … already said"*), so the false half manufactures the appearance of a second check that was
   never made. The substantive claim it dresses is independently true, which is the hesitation — but the
   document applies exactly this standard to itself at `hiring.md:270`: *"It cited `:157-159` until M257x
   iter-98 — that is the `job_position` bullet, a different construct entirely."*

---

## DEDUPLICATION

Nine distinct false predicates across ten bookings. One collapse.

| Predicate | Anchors | Bookings |
|---|---|---|
| **P1** — `SKILLER_STREAM` Go **file count** is 4 (measured: 3 Go / 6 repo-wide) | `dependency_map.md:59` | **r23-F B3 + r24-F B1 → COLLAPSE (1 predicate, 1 anchor, 2 bookings)** |
| **P2** — "not one of those line numbers resolves" at `2035f9a4`, contradicted in-cell by `roles.go:791` | `dependency_map.md:59` | r24-F B2 |
| **P3** — `LocaleSwitch` is on `/library` + `/free` | `ant-academy.md:229-230` | r23-F B1 |
| **P4** — `TopBar`'s surface set is the 5 named routes | `ant-academy.md:234` | r23-F B2 |
| **P5** — vendor selection lives in each consumer's `internal/ai/ai.go` | `shared_libraries.md:130-131` | r23-F B4 |
| **P6** — `emptyCatalogView()` = a 3-key object | `ant-academy.md:134-135` | r23-F B5 |
| **P7** — the demo's recruiter scoreboard is a dockerized-`apps/web` surface | `next-web-app.md:32` | r23-G B1 |
| **P8** — the comparable-cohort key is `jobsimulation_id` | `hiring.md:160` **and** `hiring.md:198` | r23-G B2 (**1 predicate, 2 anchors, 1 booking**) |
| **P9** — `service_taxonomy.md:52` corroborates the `public`/`jobsimulation`-schema claim | `hiring.md:38` | r24-G B1 |

**Collapses:** r23-F B3 ≡ r24-F B1 (P1) — same anchor, same ref, same measurement, booked once per reading.
This is the expected duplicate: seat F was read twice over the same files, and both readings independently
landed the same miscount in a cell whose whole purpose is to correct an earlier miscount.

**Deliberately NOT collapsed:** r24-F B2 shares the *anchor* `dependency_map.md:59` with P1 but asserts a
different thing — P1 is a count, P2 is a self-contradicting universal over line numbers. One markdown cell,
two independent defects.

**Within-booking multi-anchor:** P8 (r23-G B2) is one predicate at two anchors (`hiring.md:160`, `:198`) and
counts once.

---

BOOKED=10 UPHELD=10 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=9 DISTINCT-PREDICATES=9 WRONG-TREE-REJECTIONS=0

`IN-SCOPE-UPHELD-BLOCKERS` is the **post-deduplication** figure: 10 upheld bookings collapse to 9 distinct
predicates (r23-F B3 ≡ r24-F B1), and all 9 sit inside `corpus/services/**` or `corpus/architecture/**`, so
the in-scope count and the distinct-predicate count coincide at 9. Pre-deduplication the figure is 10.
