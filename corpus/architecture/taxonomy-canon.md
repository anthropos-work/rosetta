# The taxonomy canon (taxonomy v2)

**What the platform's skill/job-role vocabulary IS after the v2 consolidation, where it lives, and what
consuming it costs this project.** Authored M259 / v2.9 "new alphabet" — the corpus's **first** doc anchor for
the taxonomy *source pipeline*. Until now the corpus documented the taxonomy's **size** (contested, and
correctly) and its **runtime home** (`app`'s `public` schema), but nothing about where the data comes from or
how it is governed.

> **Measured 2026-08-14** against `app` **`4bccda085`** (v2.3.2) and `next-web-app` **`20a410d7d`** (v2.144.1),
> both `origin/main`, in `stack-dev/`. Every count below is either read from a file in that tree or quoted from
> the platform's own source comments — and each is labelled with which.

> **Update 2026-08-26 — production ADOPTED the canon, and it is now measured rather than inferred from a plan
> document.** Everything below stands as the M259 reading of the checked-in artifact; **§8** records the fact
> that changes its status: prod's `public.*` tables now **serve** the canon (`organization_id IS NULL AND
> deleted = false` → 3,562 skills / 709 roles / 32 categories), with the legacy set soft-deleted **in place**
> and the redirect tables loaded. §6's snapshot-cost reading is likewise superseded where marked.

---

## 1. Where it lives

**The canon is a checked-in artifact inside `app`**, not a separate dataset repo and not a runtime import:

```
app/taxonomy-canon/
  canon/     canone_skill.jsonl · canone_entita.jsonl · indice.csv · traduzioni_it.jsonl
  bundle/    canonical_roles.csv · skill_redirects.csv · role_redirects.csv
             skills_drop.csv · roles_drop.csv · role_skill_profiles.csv
             role_pages.jsonl · role_provenance.jsonl · role_tool_usedby.csv
```

Supporting code, all net-new in the `feat/taxonomyv2` program:

| Package | Role |
|---|---|
| `app/cmd/taxonomy-load/main.go` | the loader |
| `app/internal/taxonomyredirect/redirect.go` | resolves a retired `node_id` → what it became |
| `app/internal/taxonomyguard/chokepoint_test.go` | **closes the taxonomy to runtime minting** (§5) |
| `app/internal/taxonomyapi`, `taxonomyload`, `skilltaxonomy` | API / load / domain surfaces |
| `app/knowledge/taxonomy-canon-migration.md`, `…-release-checklist.md` | the platform's own migration plan |

> **This is the third thing to be folded into `app` in a year.** `app` `e72f18199` — *"chore(deps): fold
> taxonomy in — app has no first-party module left"* — retires the `taxonomy` **Go module** into the monolith.
> [`shared_libraries.md`](shared_libraries.md) states `app/go.mod` requires **five** org-private modules; after
> this it requires **zero first-party**. That correction is M264's, not this doc's, but it belongs on the record
> here because it is the same commit range.

---

## 2. The numbers — and *which* numbers, because three sets are in circulation

The bundle has been **regenerated more than once**, so the platform's own comments quote figures that its
current files no longer match. They are not contradictions to resolve; they are dated snapshots. Never merge
them into one figure.

| Source | Skills kept | Roles kept | Skill redirects | Role redirects | Total retired |
|---|---|---|---|---|---|
| **the bundle, measured 2026-08-14** (this doc) | **3,562** | **706** | **12,835** | **11,182** | 61,224 *(derived)* |
| `internal/taxonomyredirect/redirect.go` pkg doc | 3,590 | 740 | 12,785 | 11,121 | — |
| `next-web-app` `e883efd37` commit message | — | — | — | — | **61,216** *(21,871 roles)* |

The `redirect.go` row is the **oldest**; the commit message sits between it and the current files. All three
describe the same consolidation at three moments of its regeneration.

> **A FOURTH set exists since 2026-08-26 — the LIVE one.** Production's adopted counts are **3,562 skills /
> 709 roles / 32 categories** (§8). Same rule as the other three: a dated snapshot, never merged — note the
> live roles figure (**709**) is three above the bundle's 706, exactly the class of regeneration delta the
> paragraph above says to record rather than average away.

Measured from the bundle, 2026-08-14:

| Quantity | Count | File |
|---|---:|---|
| canon skills | **3,562** | `canon/canone_skill.jsonl` |
| canon entities | 289 | `canon/canone_entita.jsonl` |
| canonical roles | **706** | `bundle/canonical_roles.csv` |
| skills with a redirect | 12,835 | `bundle/skill_redirects.csv` |
| skills **dropped, no successor** | **26,518** | `bundle/skills_drop.csv` |
| roles with a redirect | 11,182 | `bundle/role_redirects.csv` |
| roles **dropped, no successor** | **10,689** | `bundle/roles_drop.csv` |
| role↔skill profiles | 14,106 | `bundle/role_skill_profiles.csv` |

### The arithmetic closes, and it vindicates this corpus's contested figures

Retired **roles** = 11,182 + 10,689 = **21,871** — matching `e883efd37`'s *"21,871 of them roles"* **exactly**.
Retired **skills** = 12,835 + 26,518 = **39,353**, so total retired = **61,224** against that commit's
**61,216** — a difference of **8**, which is the bundle regenerating after the comment was written. Record the
delta; do not average it away.

Pre-consolidation totals therefore read **42,915 skills / 22,577 roles** from the bundle, while
`redirect.go` states *"collapses **43,584** skills onto 3,590 and **22,511** job roles onto 740"*.

> **This settles the long-running "60K / 18K" dispute in this corpus's favour, and supplies the number the
> corpus could not measure.** [`shared_libraries.md`](shared_libraries.md#taxonomy-figures) held **≥42,790
> public skills / ≥22,470 public job roles** (2026-06-29) and insisted the public figure was a **floor, never a
> total**, because a public-only capture cannot see org-private rows. The platform's own pre-consolidation
> total is **43,584 / 22,511** — above the floor by **794 skills / 41 roles**, exactly the private remainder the
> floor language predicted. **"18K roles" stays REFUTED; "60K skills" is now REFUTED too**, not merely
> unverified: the platform counted 43,584 before it removed anything.

### The third lineage: ESCO, and why `skills-and-job-roles` is not the source

A separate figure circulates for "the old taxonomy": **12,201 skills / 1,893 job roles**, from
`anthropos-work/skills-and-job-roles` (`skillsandjobroles/skill.csv`, `job_role.csv`; last commit
**2024-04-08**). It is real, and it is **a different thing** — an early **ESCO** import, two years stale, and
never the same population as the 43,584 production rows the canon consolidated.

ESCO survives as **provenance on the canon, not as its source**: `canonical_roles.csv` carries `esco_uri` and
`isco` on **454 of 706** canonical roles (64 %) and `onet_code` on 495 (70 %). So the canon cites ESCO/ISCO/O*NET
where a role maps to them and stands on its own where it does not.

**Practical consequence:** `skills-and-job-roles` is **not** a thing to pull, diff or track for this work, and
its "dormant ≥18 months" grading in [`org-repos.md`](org-repos.md) is correct and unchanged. Anyone reaching for
"the old taxonomy repo" is reaching two years and one whole lineage sideways.

---

## 3. The redirect map — partial in coverage, sound in quality

Every retired id falls into one of two sets, and the distinction is the single most consequential fact for
anything holding old ids:

- **redirected** — 12,835 skills / 11,182 roles have a successor. The map is a **clean total function on its
  domain**: 12,835 distinct old ids, zero duplicates, **zero empty destinations**, and (per `redirect.go`) no
  chains at all — no destination is itself a source, and there are no self-redirects. `maxHops = 3` exists to
  turn a future cycle into an error, not to support depth.
- **dropped** — **26,518 skills / 10,689 roles have no successor at all.** `skills_drop.csv` carries a `motivo`
  (reason), and the dominant one is *"nessun riferimento: utenti, simulazioni, contenuti, profili di ruolo"* —
  **nothing in production referenced them.** So the drop set was chosen against *production* usage, which is
  **not** the same population a demo seeds from.

### ⚠️ A correction this doc makes about its own first reading

An early sample of `skill_redirects.csv` suggested the map was semantically warped — generic skills like
*"Financial Planning and Management"* mapping into agriculture. **That reading was wrong, and it is recorded
here because the mistake is instructive:** the CSV is **ordered by target category** and begins with
Agriculture, so the first N rows are a biased sample of exactly one domain. Re-sampled **randomly**, both
classes read sensibly — `AJAX → Frontend Development`, `CRM Systems in Sales → Sales and CRM Technology
Fundamentals`, `Litigation Process Management → Litigation & Trial Advocacy` — and the target-category
distribution is an ordinary professional spread (IT 2,445 · Engineering 1,879 · Safety & Security 1,208 ·
Education 1,147 · Legal 522 · Marketing 518 · Design 493 · HR 409).

**`review` is a confidence flag, not a quality verdict.** 8,456 of 12,835 skill redirects (66 %) are
`review=true` — machine re-matched (`rematch`, `rematch-binario`) and not human-confirmed; 4,370 are
`review=false`. For roles it inverts: 7,499 confirmed vs 3,670 flagged. Sampled randomly, `review=true` rows
are broadly reasonable. **Treat the flag as a ranking signal, never as a filter that implies the rest are
wrong.**

Cosmetic data wrinkles, recorded so they are not rediscovered as bugs: 299 rows have a blank `old_name` and 28
a blank `new_skill`. **Both are name columns; every `node_id` is populated.**

---

## 4. The retired-id contract: 404, and a UI that says so

Retired ids do not resolve. `next-web-app`'s taxonomy pages state it directly — a retired `node_id` yields *"a
404"*, and *"every saved URL, cached response or export"* pointing at one is affected. The frontend ships
`MovedNotice.tsx` for the redirected case.

**`taxonomyredirect` is INERT on purpose.** Its own doc: *"Nothing calls this yet, and that is the point of
shipping it first: the tables go to production and start being populated while every read path still behaves
exactly as it does today."* So the mapping **exists as data before it exists as behaviour** — which is precisely
why a consumer like this project can use it while the platform has not yet adopted it.

---

## 5. The taxonomy is closed to runtime minting

`internal/taxonomyguard` *"closes the global taxonomy to runtime minting"*. Two paths used to create nodes on a
lookup miss — `skilltaxonomy.CreateSkill` and `jobrole.CreateJobRole` (via `GetOrCreateJobRole`, which import
and matching flows call whenever a keyword does not resolve). The guard asserts the chokepoint rather than
assuming it: exactly two files may call the create methods, both call the guard first, and **a third caller
fails a test** rather than passing review as "just another get-or-create".

> **For this project this is the rule change that matters most.** Any tooling that fed a skill or role **name**
> to the platform and relied on get-or-create to mint it will now silently resolve to nothing. **Resolve against
> the canon, or accept the miss — never expect creation.**

---

## 6. What it costs this project

### 6.1 Five net-new tables sit OUTSIDE the snapshot capture surface — ✅ four of the five have since JOINED it

`feat/taxonomyv2` adds these ent schemas. **The table below was authored at M259 all-❌ and is updated in
place**: M260/M261 (this same release) grew the surface **10 → 14 tables**, and the 2026-08-26 live capture
(§8) proves each landed row with a real row count.

| Table | Captured today? | Consequence on a replayed stack |
|---|---|---|
| `skill_redirect` | ✅ M260/M261 (12,835 rows live, 2026-08-26) | redirect resolution works |
| `job_role_redirect` | ✅ M260/M261 (11,182 rows live) | same, for roles |
| `category_translation` | ✅ M260/M261 (29 rows live) | the EN/IT category level survives replay |
| `specialization_translation` | ✅ M260/M261 (283 rows live) | …and the specialization level |
| `taxonomy_canon_state` | ❌ still outside the surface | the `/taxonomy` page's canon-state panel has nothing to read |

The capture surface (`rosetta-extensions/stack-snapshot/taxonomy/taxonomy.go`) covered **ten** tables when this
doc was authored — `categories`, `job_role_categories`, `job_role_embeddings`, `job_role_skills`,
`job_role_translations`, `job_roles`, `skill_embeddings`, `skill_translations`, `skills`, `specializations` —
and covers **fourteen** now (the ten plus the four ✅ rows above). Only `taxonomy_canon_state` remains out, so
the canon-state-panel consequence stands and the two language-axis consequences no longer do.

### 6.2 The capture floor aborts before any of that matters — ✅ and then it didn't, for a reason worth naming

`stack-snapshot/taxonomy/taxonomy.go:104` pins `MinRows: 40000` on `public.skills`, enforced at
`capture/capture.go:392` — *"refusing to persist a broken snapshot"*. Against a **3,562**-row canon the capture
**aborts**, so there is no snapshot, no set-dress, no demo and no `dev-N`. It fails loudly, which is correct
behaviour for the under-capture case it was built for; it is simply now asserting a size nobody measured.

> **✅ Measured live 2026-08-26 — the capture COMPLETES against post-adoption prod** (14 tables, 309,545 rows,
> `primary-read`, digest `8b19b5b0…` — §8). Two independent things closed the gap this section predicted.
> M260 («the floor comes down», this release) owned the floor. And **production's adoption itself defused it**:
> the consolidation is a **soft-delete in place**, so `public.skills` still holds **42,908 physical rows**
> (3,562 canon + 39,346 `deleted = true`) — above the old 40,000 floor regardless. The row that scared this
> section — *"a 3,562-row canon"* — never became the physical row count, because prod deleted nothing
> physically.

### 6.3 The seed's exposure is a coverage problem, not a quality one

A seeded skill ref pointing at a retired id has a **~33 %** chance of being redirectable (12,835 of 39,353) and
a **~67 %** chance of having no successor at all. The drop set was chosen by *production* reference-counts, and
a demo seeds from the full replayed public taxonomy — so demo refs are **not** protected by that reasoning.

### 6.4 The refs the redirect map structurally cannot help: pins by **NAME**

§6.3 measures the exposure of refs held **by node-id**. There is a second, smaller population the redirect map
cannot reach *in principle*: refs held **by name**.

Both redirect tables key on `old_node_id`. A seed preset that pins `role: Business Operations Analyst` carries no
node-id at all, so no lookup — automated or manual — can follow it. It does not 404 the way §6.4's retired-id
contract promises; the resolver simply returns nothing, and what happens next depends entirely on the consumer:

| Consumer | Behaviour when the name stops resolving |
|---|---|
| `PersonaSeeder` enrichment | **silent** — the hero seeds, minus their skill chain |
| the users seeder | **loud** — `hero role(s) [...] do not resolve`, the whole run fails |

The loud one is the lucky case. The silent one produces a stack that comes up green and is *hollow* — the exact
failure the row-count floor in §6.2 also cannot see, for the same underlying reason: **a count proves rows exist,
never that they resolve.**

Measured at v2.9 M265: **10** name-pins across the seed presets, of which **6** were retired by the
consolidation — a far higher hit rate than §6.3's 33 %/67 % node-id split, because named roles skew to the
common, heavily-consolidated end of the taxonomy. All six were repaired from the canon's own
`role_redirects.csv` (`old_title` → `new_role`), which *does* carry the old title and is therefore the right
lookup for this class even though the redirect *tables* are not.

Fenced by `rosetta-extensions/stack-core/seed_role_guard.py`, which walks every seed YAML in the repo rather
than the directory a defect was last found in — the M262→M265 lesson: **two fixes for this class were applied
where the bug surfaced, and the second surfaced only because a different section had never been scanned.**

### 6.5 The one that actually broke a demo: node-ids embedded in replayed CONTENT

§6.3 and §6.4 are about the *seed*. The largest exposure is somewhere neither of them looks: the
**Directus content**, which pins skills by node-id inside JSON documents.

Taxonomy and content are two **separate** snapshot surfaces. Replaying the taxonomy swaps
`public.skills` wholesale; the content is replayed unchanged, still carrying the ids it was captured
with. Measured on a cold demo-5 at v2.9 M265: **302 distinct skill node-ids** referenced by simulation
sequences, **187 of them retired** — plus more in two columns nested a level down.

What makes this severe rather than cosmetic is the resolver's non-null contract:

```
ERROR graphql resolver error error="input:38:7: publicJobSimulations[5].skills[1].name
                                    ent: skill not found"
```

`skills[].name` is non-null, so **one** unresolvable id nulls the **entire list**. The observed
result was an AI-simulations library rendering **zero cards** and every sim detail page returning an
empty HTTP 200 — while `/api/health` was 200, every container was up, and `public.skills = 3562`
was green. **A row count cannot catch a hollow row, and a liveness probe cannot catch an empty
list.** This is the [M236 iter-05 shape](../ops/demo/content-stories-spec.md) — the whole query nulls
while the header still renders — with a taxonomy cause.

**The repair** is `rosetta-extensions/stack-snapshot/realign`, which runs on **every** replay (the
two surfaces are replayed by separate invocations in either order, so running it unconditionally is
what makes the outcome order-independent). It rewrites dead ids to their successors via
`public.skill_redirects`, and it is **discovery-driven, not list-driven** — for a reason worth
keeping:

> Its first cut carried a hand-maintained list of four columns. It repaired all four, verified
> clean, exited 0 — and the next page load still failed, because ids also live in
> `sequences.validation_evaluation_criteria` and `skill_paths.chapter_list`, **nested below the top
> level**. A hand-maintained list of the places a value can hide is wrong the moment content changes
> shape.

So it enumerates every json column in the content schema and substitutes by **exact token**, which
is what makes nesting irrelevant — it never has to know a document's shape to repair it. Live proof:
**26 columns scanned, 257 dangling before, 0 after, 257 repaired**, and `skill not found` in the
backend log went **258 → 0**.

An id with **no** successor cannot be repaired and **fails the bring-up loudly**, naming the columns
— a half-realigned content set looks repaired and still nulls the resolver, so it must never exit 0.
At M265 there were none: **187 of 187 were redirectable**. That is far above §6.3's ~33 %, and the
reason generalises — *content references the skills people actually use, and those are the ones a
consolidation gives a successor*. Do not quote the 100 % as a property of the map.

**Operational note:** the cms domain caches Directus content. Realigning an already-warm stack leaves
the stale payload served until the cache turns over; a bring-up is unaffected (the cache is cold at
that point), but a manual re-realign wants a `backend` restart.

---

## 7. Verdict — **GO**

The consolidation is legible, the bundle is checked in and measurable without touching production, the redirect
map is clean and sound, and every risk to this project is in **our** tooling rather than the platform's. The
two facts that resize downstream work:

1. **The redirect map covers a third of retirements.** A remap-everything design is not viable; the seed needs a
   resolve-or-drop path for the other two thirds.
2. **Five tables must join the capture surface**, or the demo gets a taxonomy that cannot explain its own
   history and loses two levels of the language axis.

---

## 8. LIVE — prod's `public.*` IS the canon (measured 2026-08-26)

**The question §1–§7 could not answer is now measured: production adopted the canon.** At M259 this doc — and
the `CLAUDE.md` banner that cites it — could only say *"whether production has adopted the canon is NOT
measured by us; the platform's plan says not as of 2026-08-14, but that is a document."* On 2026-08-26 it was
read from prod itself (the sanctioned `primary-read` DSN; the counts below are from the capture manifest that
read produced — `.agentspace` digest `8b19b5b0…`, 14 tables, 309,545 rows, `public_only: true`).

**The serving model is a soft-delete in place, and every number closes:**

| Quantity | Count | How it's read |
|---|---:|---|
| canon skills (`organization_id IS NULL AND deleted = false`) | **3,562** | prod live, 2026-08-26 |
| canon job roles (same predicate) | **709** | prod live |
| canon categories | **32** | prod live |
| legacy skills, soft-deleted **in place** (`deleted = true`) | **39,346** | derived: 42,908 public rows − 3,562 |
| legacy roles, soft-deleted in place | 21,846 | derived: 22,555 public rows − 709 |
| `public.skill_redirects` | **12,835** | live table — matches the bundle exactly |
| `public.job_role_redirects` | **11,182** | live table — matches the bundle exactly |
| `public.skill_embeddings` | 3,562 | **canon rows only** — the legacy set has no embeddings |
| `public.job_role_embeddings` | 707 | two canon roles lack an embedding; recorded, not explained |

Three consequences, in descending order of how often they will be needed:

1. **The platform's `public.*` tables ARE the live canonical taxonomy.** Canon rows are
   `organization_id IS NULL AND deleted = false`; the legacy set was retired by **flag, not by DELETE**; the
   redirect data §3 read from CSVs is loaded as live tables. The app's `/taxonomy` surface
   (`internal/taxonomyapi`, §1's table) serves exactly this predicate — what a prod user sees *is* the canon.
2. **A replayed stack serves the live canon identically to prod — no extra step.** The stacksnap taxonomy
   surface (14 tables, §6.1) captures the whole thing **including the `deleted` flag column** and both
   redirect tables, so replay carries the serving predicate's inputs verbatim: same rows, same flags, same
   redirects. There is no post-replay canon-activation step to run, and nobody should go looking for one.
3. **The figures dispute ends in a two-figure world, permanently.** *"How many skills?"* now has two correct
   answers that must never be merged: **3,562** (what the platform serves) and **42,908** (what
   `public.skills` physically holds). The soft-delete is why the ≥42,790 floor language of
   [`shared_libraries.md`](shared_libraries.md#taxonomy-figures) still describes the physical table, and why
   §6.2's 40,000-row capture floor never tripped.

> ⚠️ **`canon_live` is a MIRROR, not a source — never capture from it.** The taxonomy-rebuild workspace's
> Postgres (on the Ithaca dev server) carries a **`canon_live` schema refreshed nightly FROM prod's live
> canon** — it exists so the rebuild tooling has something to diff against, downstream of prod. It is **not**
> the serving source: it is up to a day stale, it is not what `internal/taxonomyapi` serves, and its schema
> digest is not the platform's. A capture `--dsn` must point at prod's `public` schema (or a dump of it) — the
> only places the serving predicate, the `deleted` flags and the redirect tables are authoritative. See the
> source-selection warning in [`../ops/snapshot-cold-start.md`](../ops/snapshot-cold-start.md).

---

## See also

- [`shared_libraries.md`](shared_libraries.md#taxonomy-figures) — the figures adjudication this doc closes
- [`../ops/snapshot-spec.md`](../ops/snapshot-spec.md) — the capture surface + the tenant-data firewall
- [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md) — the seeded verified-skill chain that holds the refs
- [`org-repos.md`](org-repos.md) — `skills-and-job-roles`, `skill-index-data`, `taxonomy-generation-tool`,
  `blueprint-skill-taxonomy` (all four graded "dormant ≥18 months, DECIDE"; the canon lives in `app`, so that
  grading stands — the revamp did **not** wake them)
