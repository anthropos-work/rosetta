# seat-4 report

**File owned:** `corpus/services/hiring.md` (sole edit surface; nothing else was written).
**Anchors booked:** 6 across 4 predicates (2 of the 6 are the same anchor booked twice, see row 1).
**Sites repaired:** 8. **Anchors declined with evidence:** 1 (`:80-81`).
**Ground truth used:** `app` **`ad9f3c49`**, `next-web-app` **`8297c684`**, platform `0c91421`.
**Guards after the edit:** `markdown_structure_guard` OK · `claim_twin_guard` **0 `hiring.md` rows**.

> **The single most important measurement in this report:** `stack-demo/clones.pin.json` now pins `app` at
> **`ad9f3c49`**, which is **also `origin/main`**. The demo build pin and `origin/main` are **the same tree**.
> Every `app`-side claim in this file therefore settles identically whichever way the "settling tree follows
> the claim's SUBJECT" rule is applied — which is what makes the `:80-81` dispute decidable rather than a
> matter of adjudicator preference.

---

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "— as the twins [`service_taxonomy.md:52`](../architecture/service_taxonomy.md)" | `corpus/services/hiring.md:38` | `service_taxonomy.md:52` reads, in full, `> [dependency_map.md](./dependency_map.md)'s content-generation flow, which had it right all along.` — the **closing line of the `:44-52` blockquote correcting the Studio-Desk → Backend → Studio-Room generation edge** (re-derived in the working tree at this iter's open; iter-101's adjudicator recorded the same construct, and it is **still** the studio-room correction). It says nothing about schemas, `jobsimulation`, `public`, or migrations. The real twin is the **Tier-1 `Database` characteristic bullet at `:62`** — *"one schema, `public`, owned by `app`, which is the only repo with migrations (`repos.yml:14-17`) … the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks"*. The **second** twin in the same sentence, `dependency_map.md:78`, is **TRUE and was preserved** — *"or directly to the **`public`** schema (the legacy `jobsimulation` schema is non-authoritative)"*, § *2. Job Simulation*. | 1 |
| 2 | "share the same `jobsimulation_id` — the sim IS the position for scoreboard purposes." | `corpus/services/hiring.md:160` | **No such column exists on the live table.** `public.job_simulation_sessions` has **23** columns (`terraform/migrations/20260722104506.sql:2-27` @ `app` `ad9f3c49`); the sim reference is **`sim_id`** (DDL `:7`, Ent `internal/data/ent/schema/job_simulation_session.go:33`). `git grep 'jobsimulation_id' ad9f3c49 -- 'internal/data/ent/schema/*.go' 'internal/organization/intelligence.go'` → **0 hits**. `jobsimulation_id` was a column of the **dropped mirror** `local_jobsimulation_sessions` — created `20240527131926.sql:7`, dropped with the table at `20260729133514.sql:62`. | 1 |
| 3 | "Candidates are **comparable** when they share the same `jobsimulation_id` + `organization_id` — that pair defines **one comparable cohort** (one scoreboard)." | `corpus/services/hiring.md:198` | Same column error; the **`+ organization_id` half is TRUE and was preserved.** The cohort is the resolver's own two predicates: `jobsimulationsession.SimID(jobSimulationId)` (`internal/organization/intelligence.go:1702`) + `jobsimulationsession.OrganizationID(organizationID)` (`:1709`), and the best-attempt window partitions on the same sim column — `ROW_NUMBER() OVER (PARTITION BY sim_id, owner_id ORDER BY score DESC …)` (`:2158-2160`, inside `usersBestOrFirstJobSimulationSession`, declared `:2124`). All @ `ad9f3c49`. | 1 |
| 4 | "the nav trims the library to **AI-Simulations**, hides some member surfaces for non-admins, and gates **Workforce Intelligence off**" | `corpus/services/hiring.md:302-303` | **Two of three clauses are TRUE and were kept**; the third is retracted. Library trim: `packages/ui/src/NavBar/useNavbarSections.tsx:340-343` (`isHiringOrg` selects `[librarySimulationsMenuItem]`, `:249-256`, label `tNavbar('aiSimulations')`). Member surfaces: `:329-331` (`!isHiringOrg \|\| isAdmin`). **Workforce Intelligence is NOT gated on `isHiringOrg` anywhere.** `enterpriseWorkforceMenuItem` (`tNavbar('workforceIntelligence')`, `:391-398`) sits in the `intelligence` group whose visibility is `orgSectionVisibility({ isAdmin, showStudio })` → `intelligence: isAdmin` (`packages/ui/src/NavBar/orgGroups.ts:48-65`, field at `:61`) — **the function takes no `isHiringOrg` parameter at all** — and the item is gated on `showWorkforce` (`:568`), which defaults `true` (`:161`) and is passed `false` in exactly **two** places, **both in `apps/hiring`** (`apps/hiring/src/app/(authenticated)/(verified)/template.tsx:167`, `:248`). All @ `next-web-app` `8297c684`. | 1 |
| 5 | *(same-predicate twin of #4's neighbourhood, not booked)* "`packages/ui/src/NavBar/useNavbarSections.tsx:460`, inside `enterpriseInsightsMenuItem` (`:459-466`)" | `corpus/services/hiring.md:292` (+ the paraphrase at `:117`) | The relabel is at **`:476`**, `enterpriseInsightsMenuItem` at **`:475-482`** @ `8297c684`. `:460` is a comment (*"Legacy assignments surface…"*). **These were EXACT at `bb3313bc`** (iter-101's ref) — pure ref drift, and the citations carried **no ref at all**, so they simply rotted. Both sites re-pinned to `8297c684` with the old value recorded. | 2 |
| 6 | *(same-predicate twin of the `:80-81` anchor set, not booked)* "`CreateOrganizationSimInvitationLink` hard-errors `"organization is not hiring"` (`siminvitationlink.go:62`)" | `corpus/services/hiring.md:130` | `:62` is the guard `if !org.IsHiring {`; the quoted string is at **`:63`** (`return nil, fmt.Errorf("organization is not hiring")`) @ `ad9f3c49`. iter-100 fixed the `:82` site to `:63` and left this one at `:62` — the document contradicted itself. Now `:63`, guard `:62`. | 1 |
| 7 | *(induced-defect check spillover, not booked)* "\| 5 \| `intelligence.go:1728-1735` \| best-attempt: `row_number() ORDER BY score DESC` per candidate \|" | `corpus/services/hiring.md:208` (read-path table row 5) | `:1728-1735` holds the `onlyAssignments` branch and the **call site** (`:1733`); the `ROW_NUMBER()` is at **`:2158-2160`**, inside `usersBestOrFirstJobSimulationSession` (`:2124`). Repaired **because row 3 above now cites `:2158-2160` eight lines away** — leaving it would have made the document contradict itself (rule 4c). Also widened the sort-block cite `:1738-1751` → **`:1738-1764`** (the three sort fields run to `:1764`). | 1 |

**Row 8 — the in-file self-citations, recomputed.** The note at `:290` cited **`:170-175`** ("the History
blockquote") and **`:157-159`** ("the `job_position` bullet"). My repairs shifted those blocks by **+23** and
**+16**. Both were re-derived and rewritten to **`:193-198`** / **`:173-175`**, with the construct **named**
so the next shift is self-correcting. This is precisely the class iter-100 induced (`service_taxonomy.md:130-133`)
and it was checked twice — once after the main repairs, once after the banner amendment moved everything a
further +1.

---

## The `:80-81` verdict — **DECLINED. I did not edit it.**

**Verdict: the sentence at (then) `hiring.md:80-81` is TRUE at the settling tree, and there was nothing to
repair. It had already been repaired by iter-100.**

### Derivation

**1. The booked false text is not in the corpus.** The union's claim is *"`manager.go:485` is `}` and `:448` a
blank line"*. `git log -p -- corpus/services/hiring.md` shows iter-100 (`a229f8d`) already rewrote it:

```
-   … as do `organization/manager.go:448` (a forced Clerk membership
-   is created with role `candidate` instead of `member`) and `:485` + `siminvitationlink.go:62` (both
+   … as do `organization/manager.go:450` (a forced Clerk membership
+   is created with role `candidate` instead of `member`, `:453`) and `:537` + `siminvitationlink.go:63` (both
```

iter-99 graded the pre-`a229f8d` state; iter-100 repaired it. The union carried the finding forward across a
repair that had already landed.

**2. The current text is exact at `ad9f3c49`.** Measured with `git show <ref>:internal/organization/manager.go`:

| anchor | value @ `ad9f3c49` | doc's claim |
|---|---|---|
| `manager.go:450` | `switch org.IsHiring {` | branches on the column ✓ |
| `manager.go:453` | `antRole = enum.RoleCandidate` | role `candidate` not `member` ✓ |
| `manager.go:537` | `return fmt.Errorf("organization is not hiring")` | hard-errors that string ✓ |
| `siminvitationlink.go:63` | `return nil, fmt.Errorf("organization is not hiring")` | hard-errors that string ✓ |

**3. Which tree settles it — and why the question turns out not to matter.** The claim's SUBJECT is `app`
source behaviour reachable on a local stack, so by the brief's rule the demo build pin settles it. The pin is
`stack-demo/clones.pin.json` → **`app: ad9f3c498e9c…`**, and `git -C stack-demo/app rev-parse origin/main` is
**the same sha**. **The two candidate settling trees are one tree.** I additionally checked the two historical
refs the dispute was argued over: the current numbers are byte-identical at **`b948604f`** (iter-101's demo pin)
and **`2035f9a4`** (the then-`origin/main`). There is no reachable ref among {`ad9f3c49`, `b948604f`,
`2035f9a4`} at which the current sentence is false.

**4. Where the two adjudicators actually stood, re-derived.** Both were right about the *old* text:

- **Adj2 (REJECT, ref-discipline)** — the old `:448`/`:485` were **exact at `5ba17044`**, the ref
  `hiring.md:17`'s re-grounding banner names. Confirmed: at `5ba17044`, `:448` is `switch org.IsHiring {` and
  `:485` is `if !org.IsHiring {`. Adj2's reading of the *evidence* was correct.
- **Adj4 (UPHOLD, §5 rule 33)** — the old `:448`/`:485` were **wrong at `b948604f` and `2035f9a4`**, both of
  which are 2 lines offset from `5ba17044`. Also correct.

The disagreement was never about the code; it was about **whether a banner 60+ lines away pins the anchors
below it.** That question is now moot for the *values* (all live refs agree) but was live for the *shelter*.

**5. What I did instead of editing the sentence.** I amended the banner at `:17` — the shelter itself — because
it was the one live hazard left: `5ba17044` is now **the only ref at which the current, correct
`manager.go:450`/`:453`/`:537` anchors DO NOT resolve** (each off by −2). A future reading applying §5 rule 33
with `:17` as the governing pin would have re-booked a correct sentence. The amendment states that `5ba17044`
is the historical iter-23 re-grounding ref, **not a governing pin**; that anchors re-derived at iter-102 are
measured at `ad9f3c49` (= `origin/main` **and** the demo pin); and that the two refs are not interchangeable,
with the −2 offset named. **This is an addition to settle ref-discipline, not a repair of a false claim, and it
is disclosed as such.**

---

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `hiring-twin-service-taxonomy-52` **+** `hiring-twin-citation` (**one defect, two bookings**) | 2 (same anchor `:38` from both unions) | 1 | 1 | `grep -n 'service_taxonomy' corpus/services/hiring.md` → 1 hit. Corpus-wide: `git grep -n 'service_taxonomy.md:5[0-9]\|:6[0-9]' HEAD -- 'corpus/**' 'CLAUDE.md'` → the only line-numbered citation of that file anywhere is this one. **No twin outside my file.** |
| `comparable-cohort-key` | 2 (`:160`, `:198`) | 2 | 2 | `grep -n 'jobsimulation_id\|comparabl\|sim_id'` in-file → exactly the 2 booked sites. Corpus-wide: `git grep -n 'jobsimulation_id' HEAD -- 'corpus/**' 'CLAUDE.md'` → **0 hits outside `hiring.md`**. The booking's width of 2 was, unusually, exact. |
| `workforce-intelligence-hiring-gate` | 1 (`:302-303`) | 1 | 1 | `grep -niE 'workforce intelligence\|intelligence off\|gates .*off'` in-file → 1 hit. Corpus-wide `git grep -il 'workforce intelligence' HEAD -- 'corpus/**'` → 4 other files, all inspected: `frontend-tier.md:380`, `tailscale-serve.md:394`, `staging-bringup.md:421/585/597/599`, `staging-sync.md:16` — **none makes an `is_hiring`/`isHiringOrg` gating claim**. No twin. |
| `hiring-manager-go-anchors` | 1 (`:80-81`) | 1 | **0 — declined, see above** | `grep -n 'manager\.go'` → 1 hit; re-derived at 4 refs. **Its un-booked twin at `:130` (`siminvitationlink.go:62`) WAS repaired** — that is the paraphrase axis the booking missed. |
| *(un-booked, same-neighbourhood)* nav-anchor currency `useNavbarSections.tsx:460` | 0 | 2 (`:117`, `:292`) | 2 | `grep -n 'useNavbarSections'` in-file. Measured drift `bb3313bc` → `8297c684`. |
| *(un-booked, rule-4c spillover)* read-path row 5 `row_number()` site | 0 | 1 (`:208`) | 1 | Surfaced while verifying row 3's new `:2158-2160` cite against the table 8 lines up. |

**Totals: 6 anchors booked → 6 sites found on the booked predicates → 5 repaired + 1 declined with evidence;
plus 3 un-booked same-file twin/currency sites repaired. 8 sites edited overall.**

Honest note on width: three of the four booked predicates turned out to be **exactly as wide as booked** — the
2–3× under-count the brief warns about did **not** reproduce here. What *did* reproduce is the other half of
rule 2: **everything extra I found was a paraphrase or a stale ref, not a twin** (the `siminvitationlink.go:62`
twin, the two `:460` sites, the row-5 range). Searching for the *fact* found them; searching for the *phrasing*
would not have.

---

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| `hiring-twin-service-taxonomy-52` | **`corpus/architecture/service_taxonomy.md:62`** — **SEAT 3's FILE** | ⚠️ **CROSS-SEAT COUPLING — the orchestrator must re-verify after both seats land.** My repair at `hiring.md:38` now cites `service_taxonomy.md`'s **Tier-1 `Database` characteristic bullet** at **`:62`**. Seat 3 is editing that file concurrently and **`:62` can move.** I cited the construct **by NAME** (*"Tier-1 **Database** characteristic bullet"*) **and quoted its text** (*"one schema, `public`, owned by `app`, which is the only repo with migrations"*) so the citation survives a shift, and I said so in the corpus text itself. **I did not edit `service_taxonomy.md`.** Re-derive `:62` after seat 3 commits. *(I also cite `service_taxonomy.md`'s `:44-52` blockquote in the retraction note — same coupling, same mitigation: named as the Studio-Desk → Backend → Studio-Room generation-edge correction.)* |
| `hiring-twin-service-taxonomy-52` (the surviving TRUE half) | `corpus/architecture/dependency_map.md:78` — **seat 7's file** | The second twin named in the same sentence. **Verified TRUE and preserved** — § *2. Job Simulation*: *"or directly to the **`public`** schema (the legacy `jobsimulation` schema is non-authoritative)"*. **No edit needed; do not let a `dependency_map.md` repair delete it**, or my citation at `hiring.md:38` breaks. |
| `recruiter-scoreboard-app` (iter-101 row 22) | `corpus/services/next-web-app.md:32` — **not mine** | `hiring.md:53` and `:352` are the adjudicator's own cited **counter-evidence** for that booking. My edits did not touch either sentence; both still say *"**not in `apps/web`**"* / *"**it does NOT land in `apps/web`**"*. My new `:302-303` text is consistent with them (the recruiter loses Workforce Intelligence **by being handed to `apps/hiring`**), so the two files now agree rather than merely not-conflicting. |

---

## Noticed, not repaired

1. **`corpus/services/hiring.md:208-210` read-path rows 6 & 7 verify** (`intelligence.go:1820` =
   `score := RoundFloat(float64(ls.Score), 0)`; `:1846` = `Score: &score`; `job_simulation_session.go:45` =
   `field.Float32("score").Default(0).Min(0).Max(100)`) — all exact at `ad9f3c49`. Recorded because I read them
   while repairing row 5; nothing to do.
2. **`hiring.md:224-226` (the silent-403 substrate) cites `resolver_queries.go:1035` / `:1089` / `:1085`, and
   `:80-84` cites `:1034-1080` / `:1035` / `:1053`.** I did **not** re-derive these — they are outside my four
   predicates and outside the `:80-81` anchor set proper. They are `app` anchors and `app` moved 5 commits this
   iter, so they are candidates for a currency check by whoever holds `resolver_queries.go`-shaped claims.
3. **`hiring.md` carries several `app`/`next-web-app` anchors with no ref attached at all** (e.g.
   `useGetClerkOrganization.tsx:16-18` / `:20-21`, `template.tsx:90`, `FreeTrialContainer.tsx:29`,
   `UserStatusContext.tsx:125,144-145`, `resolver_cms_queries.go:95,210,258,295`). Unpinned anchors are what
   produced defect-row 5 in this report — they do not *break*, they *rot silently*. The banner amendment tells a
   reader to read the ref that travels with the anchor; **most anchors in this file still carry none.** This is a
   structural observation, not a booking.
4. **The banner at `:17` still names platform `2adcf71` / `0dab54d`** while this iter's platform ground truth is
   `0c91421`. I deliberately left the platform half alone — it already carries its own re-anchoring parenthetical,
   and rewriting platform refs here would ripple into clauses I did not measure.

## What I could not settle, and why

1. **`service_taxonomy.md:62` is a moving target while seat 3 is live.** I re-derived it at this iter's open and
   it is the Tier-1 `Database` bullet. I mitigated by naming + quoting the construct, but I **cannot** guarantee
   the bare number `:62` after seat 3 commits, and I was instructed not to edit that file. **Flagged above for
   orchestrator re-verification.** This is the one thing in my partition that is not fully settled by
   measurement.

   **Live re-check at my hand-off** (seat 3's file already shows as modified in `git status`): both lines still
   hold in the working tree — `grep -n '^- \*\*Database\*\*' corpus/architecture/service_taxonomy.md` → **`62`**,
   and the studio-room correction's closing line → **`52`**. So the coupling is **currently satisfied**; it just
   is not *guaranteed* if seat 3 writes again after this. **One `grep -n '^- \*\*Database\*\*'` re-run at commit
   time is the whole verification.**
2. **Nothing else.** All four booked predicates were settled against a named ref, in a clone, at a sha, with the
   negative findings cross-checked by a second mechanism (`git grep <ref>` for the corpus-wide absence of
   `jobsimulation_id` and of any other `service_taxonomy.md` line citation; `git grep -il` + per-file inspection
   for the Workforce-Intelligence paraphrase sweep).

## Discipline notes

- **No commit, no `git add`, no fetch, no clone write.** Read-only against `stack-demo/**` throughout.
- **Zero platform-repo edits.** One corpus file touched: `corpus/services/hiring.md` (+64 / −17).
- **Rule 5 honoured explicitly at three sites** — the `dependency_map.md:78` twin (row 1), the
  `+ organization_id` half of the cohort key (row 3), and two of the three clauses in the `is_hiring` nav bullet
  (row 4) were **true, are named as true in the repaired text, and were preserved verbatim in substance.**
- **Rule 4 caught one defect I induced and fixed it in-iter**: the banner amendment initially wrapped
  `` `internal/organization/manager.go:450` `` across a line break inside the code span, which renders as a
  broken path (`internal/organization/ manager.go:450`). Re-wrapped. A whole-file odd-backtick scan now shows
  only the file's pre-existing multi-line spans, none of which break a path.
- **Line-number stability re-verified twice** (after the main repairs, and again after the banner amendment
  shifted everything +1): the two in-file self-citations are correct at `:193-198` and `:173-175`.
