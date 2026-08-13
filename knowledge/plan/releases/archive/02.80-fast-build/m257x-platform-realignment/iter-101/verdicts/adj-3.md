# ADJUDICATOR 3 — iter-101 — seats r23-E, r24-E

**Docket:** 7 bookings (r23-E: 4 · r24-E: 3). Seat E files: `service_taxonomy.md`,
`clerkenstein.md`, `security_compliance.md`, `ai-labs.md`, `clerk-integration.md`, `intelligence.md`.

**Refs re-verified at my open (`git rev-parse`), not taken from the seats:**
rosetta HEAD `8b6d80f5` (working tree clean under `corpus/`) ·
platform `0c91421d` · app `b948604f` · cms `ca50c817` · next-web-app `bb3313bc` · sentinel `88bc5592` ·
storage `4ce8ece5` · messenger `fa47850d` · graphql-wundergraph `60c229f3` · roadrunner `87d8d443` ·
jobsimulation `462343b0` · studio-desk `14a5442a` · ant-academy `9c3843cd` ·
`stack-demo/rosetta-extensions` `ab81527a` (pinned consumption clone) ·
`.agentspace/rosetta-extensions` `09d06070` (authoring copy).

**Two-clone-set handling.** No blocker in this docket rests on a `rosetta-extensions` claim — both
seats put every rext observation (`clerkenstein.md:21` section count, the `cmd/` `jwtkey` omission)
in *Minors*, not Blockers. I checked the trees anyway so the instrument note is on the record:
`diff -rq --exclude=.git` between the two clones shows **no source divergence** — every difference is
a runtime artifact (`.pytest_cache`, `demo-stack/stacks/demo-1/*.log|.pid|.env`, `.m220-mutant-*`
scratch files, a built `demopatch` binary). Both trees carry the same **12** top-level entries
(11 code sections + `knowledge/`) and the same `clerkenstein/cmd` = `{fake-bapi, fake-fapi, jwtkey,
mintpk}`. So **the known two-tree defect did not bite this docket**, and would not have changed a
verdict if it had. Zero `wrong-tree` rejections.

---

## Verdicts

### r23-E B1 | `corpus/architecture/service_taxonomy.md:130-133` | **UPHELD** | IN-SCOPE
`archive-state note's own row anchors are +2 off; names non-archive rows and calls an assertion a retraction`

   evidence: I opened `corpus/architecture/service_taxonomy.md` at rosetta HEAD `8b6d80f5` and
   enumerated the table rather than trusting the seat. `grep -n ARCHIVED` over the file returns
   exactly **four** hits, of which only **two** are table cells: `:139` (Skiller, `ARCHIVED
   2026-07-01`) and `:140` (Skillpath, `ARCHIVED 2026-07-31`). `:137` = Chronos and `:138` =
   Intelligence — I read both cells in full; each is `| Removed from local dev orchestration | **no**
   | Platform commit … |` and **contains no archive assertion of any kind**. `:141` = Jobsimulation,
   which carries *"repo archive state: report both, assert neither"* — that is the retraction. The
   note at `:130-133` says the flat form was published at `:137`/`:138` "two rows above `:139`, a
   cell retracting exactly that predicate". Every one of the three anchors is shifted by exactly
   **+2**: the flat form is at `:139`/`:140`, two rows above **`:141`**, and `:139` is one of the
   two cells *asserting* the predicate, not retracting it.

   I then established this is a live regression rather than a long-standing wording choice, which is
   the part that removes any ref-discipline defence. `git log -S'service_taxonomy.md:137'` returns
   **one** commit, `a229f8d` (iter-100). At `a229f8d^` the note read *"rows `:137`/`:138`"* and the
   table was: `:135` Chronos · `:136` Intelligence · **`:137` Skiller** · **`:138` Skillpath** ·
   **`:139` Jobsimulation**. The note was **exactly correct there**. `a229f8d` inserted the
   two-line parenthetical *("The file is named explicitly as of M257x iter-100…")* above the table,
   pushing every row down by two, and left the line numbers unmoved. The note carries no ref and
   names its own current file, so it is graded at the current file, where all three anchors are
   wrong.

   tree-read: rosetta corpus at HEAD `8b6d80f5`, plus rosetta git history at `a229f8d` / `a229f8d^`.

### r23-E B2 | `corpus/architecture/security_compliance.md:67-68` | **UPHELD** | IN-SCOPE
`"only 30 use OrganizationMixin{}" — 29 use it; contradicts :76, :84-85 and :135`

   evidence: re-derived the SET, not the sum, in `stack-demo/app` @ `b948604f`,
   `internal/data/ent/schema/`. Cardinality first: **139** `.go` files; **135** declare `ent.Schema`;
   the four that do not are `database_types.go`, `mixin.go`, `skiller_mixins.go`,
   `skillpath_mixins.go` (enumerated by `comm`, not counted). Files containing the literal
   `OrganizationMixin{}` = **30**, and bare `grep -l` and `git grep -l HEAD --` agree file-for-file
   (I diffed the two lists; identical — no `.gitignore`/NUL/nested-repo skew here). I then opened the
   30th: `internal/data/ent/schema/user_resource.go:22` sits inside `func (UserResource) Mixin()` and
   reads `// OrganizationMixin{},  // We need to work on this` — commented out, not compiled. So
   **29 schemas use it; 30 mention it**, and the sentence's predicate is *"use"*.

   The self-contradiction is inside one blockquote and is arithmetic, not stylistic: `:76` reads
   *"**So:** 31 schemas auto-filter by ORGANIZATION — **29** `OrganizationMixin{}` users, plus
   two…"*, which is unreachable from `:67`'s 30; `:84-85` and `:135` both state 29 explicitly. The
   decisive point is that `:67-68` is prefixed *"Measured at `app` HEAD"* (HEAD = my checkout, where
   the answer is 29) and that the same fence rules at `:89`: *"Re-derive the SET, not the sum, and
   exclude commented lines when you do — a `grep -c` over Go source counts code that does not compile
   into anything."* The fence's opening sentence commits the exact error the fence exists to forbid.

   Everything else in the fence re-derived clean and I state it so the upheld item is isolated:
   `OrganizationIDMixin{}` users = **7** and are set-identical to the seven named (`category`,
   `jobrole`, `similarity`, `skill`, `specialization`, `studio_document`, `studio_task`); the doc's
   own `comm`/`xargs` derivation returns **18**, and removing `org_membership.go` (own `Policy()`)
   and `academy_feedback.go` (`UserMixin{}`) leaves **16**, name-for-name identical to the doc's 16;
   files declaring any `Policy()` = exactly **4** (`mixin.go`, `org_membership.go`, `organization.go`,
   `user.go`); `mixin.go:126` is `func (OrganizationMixin) Policy()`, exact.

   tree-read: `stack-demo/app` @ `b948604f` (bare `grep` and `git grep` cross-checked).

### r23-E B3 | `corpus/architecture/security_compliance.md:185` | **REJECTED** | —
`"consistent with README.md:21" allegedly resolves to an unrelated sentence in every candidate file`

   evidence: **the citation resolves, exactly.** `corpus/architecture/README.md:21` is the
   `shared_libraries.md` index row — the seat stopped there — but the row does not end there. Its
   final clause reads: *"The doc covers what each provides and where its responsibilities begin and
   end (**e.g. cost tracking lives in `app`, not the `ai` library**)."* That is verbatim the
   proposition `security_compliance.md:185` cites it for (*"cost tracking in `app/internal/aiusage`
   — **not** the shared `ai` library"*). I confirmed the line at HEAD two ways (`grep -n 'cost
   tracking' corpus/architecture/README.md` → `:21` and `:23`; `git show HEAD:corpus/architecture/
   README.md | awk 'NR==21'` tail), and confirmed `corpus/` is clean in the working tree, so both
   seats read the same bytes. The bare `README.md` resolves to the sibling
   `corpus/architecture/README.md` by the same convention as `ai_architecture.md` in the very same
   parenthetical.

   The seat's own instrument produced the miss: it searched for the token `aiusage`, found it only at
   `:23`, and concluded `:21` was unrelated — but `:21` states the claim without naming the package
   path, and `ai_architecture.md` (the second half of the same citation) is what carries the path.
   Seat r24 opened the whole line and cleared it, recording that it *"nearly booked this as a
   wrong-construct anchor before reading the whole line"*. r24 is right.

   tree-read: rosetta corpus at HEAD `8b6d80f5` (worktree and `git show HEAD:` both).
   class: **mis-read** — the seat truncated the cited line at the index-row boundary; the
   corroborating clause is the tail of that same line.

### r23-E B4 | `corpus/architecture/service_taxonomy.md:109-111` | **UPHELD** | IN-SCOPE
`"Base services (no profile, always on)" enumerates 2; the predicate admits 3, and the file says 3 thrice`

   evidence: re-derived the predicate from source, not from the seat. In `stack-demo/platform` @
   `0c91421d`: `docker-compose.yml` declares **five** services — `sentinel` (`:5`), `backend`
   (`:28`), `studio-desk` (`:112`), `next-web-app` (`:143`), `gotenberg` (`:170`) — and exactly
   **four** carry a `profiles:` key (`:110`, `:141`, `:168`, `:183`). **`sentinel` declares none.**
   `include: common.yml` (`:1-2`) adds `postgresql` and `redis`, and I read `common.yml` end to end
   (36 lines) — neither declares `profiles:` either. So the set satisfying *"no profile, always on
   with any `make up`"* is `{postgresql, redis, sentinel}`, cardinality **3**. `service_taxonomy.md:
   109-111` heads that exact predicate and lists **two** bullets.

   This is a self-contradiction on the file's own defined term, and I cite both sides: `:68` — *"`core`
   starts five: `backend`, `gotenberg` and **the three always-on base services**"*; `:465` —
   *"postgresql, redis, sentinel only — **the floor**, the three services that declare no `profiles:`
   key"*; `:489` — *"**the three always-on base services** (`postgresql`, `redis`, `sentinel`)"*.
   Root `CLAUDE.md` says three as well. Under `:109`'s enumeration the file's own arithmetic does not
   close: backend + gotenberg + 2 = **four**, against the **five** stated at `:68` and `:489`.

   I weighed the seat's own hesitation (sentinel is documented one table up at `:80` as *"(always on
   — declares no `profiles:` key)"*, so "Base services" could be read as *base infrastructure*). It
   does not survive, because `:489` uses the identical phrase *"base services"* for all three — the
   two readings are the file disagreeing with itself about the extension of its own term — and
   because the floor's cardinality is load-bearing for the `:474-479` mechanism (*a retired profile
   token exits 0 and starts the 3-service floor*), where a mis-sized floor is precisely what makes a
   dead stack look alive.

   tree-read: `stack-demo/platform` @ `0c91421d` for the source; rosetta corpus at HEAD for the
   three corroborating passages.

### r24-E B1 | `corpus/architecture/service_taxonomy.md:131-133` | **UPHELD** | IN-SCOPE
`same predicate as r23-E B1 — archive-note anchors +2 off, retraction/assertion inverted`

   evidence: identical re-derivation to r23-E B1 above (the enumerated `ARCHIVED` set, the four cells
   opened, and the `a229f8d` / `a229f8d^` history that shows the note was correct before iter-100's
   two-line insertion). r24's extra content check also holds: `045857c` and `fdfa189` are real, both
   2026-04-17, both touching `docker-compose.yml` + `repos.yml`, and the Jobsimulation retraction is
   itself correct. **The defect is purely the anchor set** — which is r24's own framing and is the
   right one.

   tree-read: rosetta corpus at HEAD `8b6d80f5` + rosetta git history; platform `0c91421d` for the
   two commits.

### r24-E B2 | `corpus/services/ai-labs.md:75` (seat cited `:76`; the wrapped sentence spans `:75-79`) | **REJECTED** | —
`"course.build=5/chapter" alleged wrong unit — it is exactly right at the named ref`

   evidence: I opened `stack-demo/app/internal/credits/cost.go` @ `b948604f` and read the **whole**
   file rather than the 4 lines cited. The package doc at `cost.go:22-28` settles it outright:

   > "The launch actions are (**Review M5 — the per-COURSE build price is per-CHAPTER × the tier's
   > chapter count, NOT a flat 5**): `course.build` → **5 credits per CHAPTER** — a course holds
   > maxChapters × 5 up front (simple 3 → 15, advanced 8 → 40) and refunds undelivered chapters.
   > `course.build` = the per-chapter unit."

   Corroborated three more ways inside the same package: `cost.go:259-267` — *"`CourseBuild`…are the
   PER-CHAPTER unit price — NOT the price of a whole course"*; `cost.go:285-292` —
   `BuildPriceCredits(maxChapters) = maxChapters * Cost(ActionCourseBuild)`; and the **live billing
   path**, which is the decisive one: `internal/web/backend/coursebuilder/handler.go:543-544`
   `perChapterCredits() = envCreditPrice("COURSEBUILDER_CREDITS_PER_CHAPTER", 5)`, `:549-550`
   `buildMaxChapters` → `PlanChapterCapsFor` (`internal/coursebuilder/planner.go:56,58` — advanced 8,
   simple 3), `:581-587` `buildCreditUnits(depth) = buildMaxChapters(depth) * perChapterUnits()`, and
   the debit at `:677` `h.debit(ctx, orgID, credits.ActionCourseBuild, units, …)`. A simple build
   holds 3 × 5 = 15 credits, an advanced one 8 × 5 = 40 — exactly `cost.go:26`.

   The seat's evidence is a **stale comment inside the file it quotes**. `cost.go:78-81` (*"for
   course.build one unit is one build"*) is contradicted 60 lines above it by a block explicitly
   labelled `Review M5` and explicitly saying *"NOT a flat 5"*. The same package doc is stale on a
   second line too — `:29` says refine = 5 while the map at `:88` says 1 — so no single comment block
   in `cost.go` is authoritative; the map + the handler are. The seat's two doc citations are stale
   for a checkable reason: `GO-LIVE-RUNBOOK.md:36-37` and `SPEC.md:2672` describe a superseded
   depth-tiered *flat* scheme keyed on `COURSEBUILDER_CREDITS_SIMPLE` / `_ADVANCED`, and
   `git grep -n COURSEBUILDER_CREDITS -- '*.go'` returns **only**
   `COURSEBUILDER_CREDITS_PER_CHAPTER` (3 hits, all in `handler.go`) — the old variables are read by
   no Go code at `b948604f`. The `1.60/5 × 1.40 = 0.45` derivation the seat offers as independent
   arithmetic is neutral: `$1.60` is the measured Bedrock COGS used to price *one credit*, and the
   whole-course price is recovered as `maxChapters × 5 × $0.45` (`BuildPriceUSD`, `cost.go:296-297`).
   The seat's *"20× error"* is inverted — the corpus, which also gets `refine`=1 (map) and
   `translate`=1/locale right, is more accurate than any single comment block in the source.

   tree-read: `stack-demo/app` @ `b948604f` (Go source + the app-repo `internal/coursebuilder/`
   SPEC/runbook the seat cited).
   class: **mis-read** — the seat quoted the one stale comment in `cost.go` and did not read the
   package doc 60 lines above it that pre-emptively refutes it by name.

### r24-E B3 | `corpus/architecture/security_compliance.md:67-68` | **UPHELD** | IN-SCOPE
`same predicate as r23-E B2 — "30 use OrganizationMixin{}" where 29 use it`

   evidence: identical independent re-derivation to r23-E B2 above (139 → 135 → 30 mentions → 29
   uses, with `user_resource.go:22` opened in place and confirmed commented out inside
   `func (UserResource) Mixin()`, and the 7 / 18 / 16 / 4 sub-sets enumerated by name). r24's added
   observation is correct and worth keeping: read from `:68`, `:76`'s "So:" would give `30 + 2 = 32`
   — the very total this document records at `:86-88` as the wrong answer iter-52 shipped and had
   refuted by two blind readers.

   tree-read: `stack-demo/app` @ `b948604f`.

---

## DEDUPLICATION

Seat E was read twice over the same six files, so duplicates were expected. Two collapses, both on
an identical predicate at an identical (or overlapping) anchor:

| predicate | anchors | bookings | verdict |
|---|---|---|---|
| **P1** — the archive-state note's three row anchors are +2 off: `:137`/`:138` carry no archive assertion, and `:139` *asserts* the flat form the note calls a retraction (which is `:141`) | `service_taxonomy.md:132-133` (r23) and `:131-133` (r24) — the **same sentence**, which spans `:130-133` | r23-E B1 + r24-E B1 | UPHELD, IN-SCOPE |
| **P2** — `"only 30 use OrganizationMixin{}"` states the containment count as the use count; 29 use it, and the same blockquote says 29 three times | `security_compliance.md:67-68` (both seats, identical anchor) | r23-E B2 + r24-E B3 | UPHELD, IN-SCOPE |
| **P3** — `"Base services (no profile, always on with any make up)"` enumerates 2 where the predicate admits 3 | `service_taxonomy.md:109-111` | r23-E B4 only | UPHELD, IN-SCOPE |
| **P4** — `"consistent with README.md:21"` names an unrelated sentence | `security_compliance.md:185` | r23-E B3 only | REJECTED (mis-read) |
| **P5** — `course.build = 5/chapter` is the wrong unit | `ai-labs.md:75` | r24-E B2 only | REJECTED (mis-read) |

**7 bookings → 5 distinct predicates. Of the 5 upheld bookings, the distinct upheld predicates are
P1, P2 and P3 — all three IN-SCOPE (`corpus/architecture/**`).**

Note that the two seats also *disagreed* on P4 in opposite directions: r23 booked it as a blocker,
r24 cleared it explicitly in its "Positively cleared" list. Re-derivation sides with r24. Neither
seat's clerkenstein/rext observations reached blocker status, so no rext-tree adjudication was
required.

---

BOOKED=7 UPHELD=5 REJECTED=2 IN-SCOPE-UPHELD-BLOCKERS=5 DISTINCT-PREDICATES=5 WRONG-TREE-REJECTIONS=0

> Disambiguation for the gate arithmetic: `IN-SCOPE-UPHELD-BLOCKERS=5` counts **bookings**
> (r23-B1, r23-B2, r23-B4, r24-B1, r24-B3) and `DISTINCT-PREDICATES=5` counts distinct predicates
> across the whole 7-booking docket. **Deduplicated, the in-scope upheld predicates number 3**
> (P1, P2, P3). Use 3 if `N` is a predicate count; use 5 if `N` is a booking count.
