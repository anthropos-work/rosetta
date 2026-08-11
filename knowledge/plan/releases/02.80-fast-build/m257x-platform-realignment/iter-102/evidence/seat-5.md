# seat-5 report

**Files owned:** `corpus/services/ant-academy.md` · `corpus/services/next-web-app.md` ·
`corpus/services/academy-backend.md`.
**Anchors booked:** 7 · **sites found:** 9 · **sites repaired:** 9 (7 booked + 2 unbooked moving-label sites
in my own files). **Twins outside my files:** 2, reported not edited.

**Settling trees used** (all re-derived at this iter's open, nothing inherited from a prior sheet):
`ant-academy` `22df69dd` · `next-web-app` `8297c684` · `app` `ad9f3c49` (== `origin/main`) ·
`platform` `0c91421` · rext = `stack-demo/rosetta-extensions` `09d06070`. No `git fetch` was run; no clone was
written.

---

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "progress persists via GraphQL mutations (`upsertChapterProgress[Batch]` / `setLastActivity`, posted from `code/app/api/academy/beacon/route.js`) to Ent tables `academy_chapter_progresses` / `academy_last_activities` / … in `app`" | `corpus/services/ant-academy.md:63` | The mutation NAMES are right; the **attribution is inverted**. @ `ant-academy` `22df69dd` every in-session write is fired from `code/src/progress/store.js`: `saveChapterProgress` (`:150`) calls the injected authed requester with `UPSERT_CHAPTER_PROGRESS` at `:162`; `saveLastActivity` (`:202`) with `SET_LAST_ACTIVITY` at `:210` — direct cross-origin calls to the supergraph with a Clerk Bearer token. `beacon/route.js` is the **on-unload last-ditch flush**, passed as the `beacon:` *option* at `store.js:169` / `:215` and reached only via `navigator.sendBeacon` / `fetch({keepalive:true})` on pagehide (`src/writeThrough/index.js:247`, `:259`); its own header comment (`route.js:1-18`) calls itself *"a best-effort last-ditch flush for a write that would otherwise be lost if the tab closes mid-retry"* and explains it exists **because** `sendBeacon` cannot set an `Authorization` header. The exception was published as the rule. | 1 |
| 2 | "so on **any** failure the catalog becomes `emptyCatalogView() = { chapters: [], skillPaths: {}, series: [] }` → **0 cards**." | `corpus/services/ant-academy.md:134-135` | The `=` asserts a **3-key** literal; measured is **5**. `code/src/lib/serverTenant.js:115-117` @ `22df69dd`: `return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }`. Neither extra key is empty — `PUBLIC_BUNDLES` (`code/ucourses/catalog.js:961`) is a populated exported array of curated bundle objects; `CATALOG_VERSION` (`:31`) is `'1.0'`. The function's own comment (`serverTenant.js:111-113`) says `bundles` *"carries no tenant metadata and is the one piece not yet modeled in the backend catalog"* and passes through verbatim. **Both other halves are TRUE and were kept**: the quoted one-liner is byte-exact at `serverTenant.js:145`, and `→ 0 cards` still holds (`AcademyClient.jsx:1363-1365` drops every bundle path whose `scopedChapters` filter is empty, so the stripes render zero path cards). | 1 |
| 3 | "but it is mounted **only in the public-storefront header** (`src/views/public/PublicHeader.jsx:20`, i.e. `/library` + `/free`)" | `corpus/services/ant-academy.md:229-230` | The *"only in the public-storefront header"* half is **TRUE and kept** (`git grep LocaleSwitch 22df69dd -- code/` → one mount, `PublicHeader.jsx:20`). The **`/free` half of the gloss is false**: `PublicHeader` has exactly one mount site, `code/app/(public)/library/page.jsx:28`, and `/free` renders no header of its own — `code/app/(public)/free/page.jsx:18` is the whole body, `redirect('/?tier=free')`, landing on the app-shell home, which serves the *other* switcher (`LanguageSelector`). The surface set is **`/library` alone**. | 1 |
| 4 | "and `TopBar` is on `/`, `/chapters/*`, `/latest`, `/bookmarks` and `/my-activity`." | `corpus/services/ant-academy.md:234` | Measured surface set @ `22df69dd` is **7**, not 5. `AcademyClient` serves **three** routes — `/` (`app/(authed)/page.jsx:151`), `/courses` (`courses/page.jsx:92`), `/courses/[slug]` (`courses/[slug]/page.jsx:219`) — and mounts `TopBar` at `AcademyClient.jsx:1906`; plus `/chapters/[slug]` (`CourseClient.jsx:2091`, `:2141`), `/latest` (`LatestClient.jsx:128`), `/bookmarks` (`BookmarksClient.jsx:508`), `/my-activity` (`MyActivityClient.jsx:161`). `/my-certificates` mounts **no** `TopBar`. The two omitted routes are the demo's landing routes. | 1 |
| 5 | "**NB the demo's recruiter candidate-comparison scoreboard is an `is_hiring` ORG-TYPE surface in the dockerized `apps/web`** (`/enterprise/activity-dashboard`), **not** this Vercel-only app" | `corpus/services/next-web-app.md:32` | False in both directions. `/enterprise/activity-dashboard` exists in **both** apps @ `8297c684`, so the route proves nothing; what decides it is a global product-boundary guard — `apps/web/src/context/UserStatusContext.tsx:144-148` computes `userHasAllHiringOrgs` from `membership.organization.publicMetadata.isHiring`, and on true `:168-172` sets `window.location.href = buildSwitchHandoffUrl({ targetProduct: 'hiring', clerkOrgId, next: '/home' })` — the recruiter is **ejected out of `apps/web`**. The comparison renders in `apps/hiring/src/components/containers/InsightsByMembersContainer.tsx:108`, mounted at `apps/hiring/…/activity-dashboard/@tabs/ai-simulations/[simId]/page.tsx:14`. **The true half kept:** the scoreboard *is* driven by the `is_hiring` org-type and *does* render from seedable data. Corroborated by `hiring.md:53` and `hiring.md:353`, which already said the opposite of `:32`. | 1 |
| 6 | "The hiring **org-type** (`is_hiring`) re-skins `apps/web` and exposes the recruiter **candidate-comparison read-model**" | `corpus/services/next-web-app.md:14` | Supporting site of row 5 — a conjunction whose first half is **TRUE** (`hiring.md:48`: the org-type *"re-skins the **Workforce** app (`apps/web`) for a recruiting buyer"*) and whose second half places the read-model in `apps/web`, which the ejection refutes. Repaired without weakening the re-skin clause. | 1 |
| 7 | "**`v1.369.0`** @ origin/main `2035f9a4`, measured 2026-08-06" **(the false span only — narrowed at consolidation).** The sentence's first half — the backend domain's own v1.0/v1.05 label line, distinct from `app` SemVer in the v1.3xx range — is **TRUE and was kept**; only the moving `origin/main` label expired, so quoting the whole sentence fenced a true clause. | `corpus/services/academy-backend.md:20` | CANON-3, routed to **move (2)** (a version claim). Re-derived at the new head: `git describe --tags ad9f3c49` → **`v1.369.0-7-gad9f3c498`**, so the version reading is **still `v1.369.0`** (tag dated 2026-08-05 in `CHANGELOG.md`) but the ref it is read at is now **`ad9f3c49`**, and `2035f9a4` is 5 commits behind (`git rev-list --count 2035f9a4..ad9f3c49` → 5). Re-stated with the new ref + a date. **`2035f9a4` was not silently swapped for `ad9f3c49`** — the label is what expired, and that is said in place. | 1 |
| 8 | "which does not exist at origin HEAD" (of `middleware.ts`) | `corpus/services/next-web-app.md:49` | Not booked — an **unbooked bare moving label** in my own file, the exact rot class CANON-3 names (*"Never leave a bare moving label"*), in a repo that moved 41 commits this cycle. The claim itself **verifies**: at `8297c684` there is no `middleware.ts` anywhere; `proxy.ts` exists in `apps/{web,hiring,integration}`; `CLAUDE.md:55` reads *"**Clerk middleware** lives in `src/proxy.ts` (Next 16 renamed the `middleware.ts` convention → `proxy.ts`)"*. Pinned to `8297c684` + dated; content unchanged. | 1 |
| 9 | "no `storage` script and no `.storybook/` dir exist at origin HEAD" | `corpus/services/next-web-app.md:72` | Same class as row 8. Verifies at `8297c684` (no `storybook` script, no `.storybook/`; the only trace is `configs/tailwind/storybooks.css`). Pinned + dated; content unchanged. | 1 |

---

## The `:134-135` verdict — the adjudicator's disclosed-weakest uphold

**VERDICT: UPHELD, and repaired.** The disclosure was warranted as a *materiality* note and is wrong as a
*falsity* note.

- **Literally false, unambiguously.** The sentence asserts the shape with `=`, not "roughly" or "e.g.":
  `emptyCatalogView() = { chapters: [], skillPaths: {}, series: [] }`. Measured at the settling tree
  (`serverTenant.js:116` @ `22df69dd`) the function returns five keys. An `=` between a named function call and a
  literal is an equality claim, and it does not hold.
- **Non-vacuous.** Had the two extra keys been `[]` / `null` the finding would arguably be cosmetic. They are
  not: `PUBLIC_BUNDLES` is a populated array (`catalog.js:961` ff.) and `catalogVersion` is `'1.0'`. A reader who
  trusts `:135` concludes the empty view carries **nothing**, and therefore that `bundles` must arrive from the
  backend — it does not; it is an FS constant that survives the DB cutover by design, which is exactly the
  distinction the surrounding section exists to draw.
- **Why the adjudicator's caution was right anyway.** The *conclusion* the sentence draws — `→ 0 cards` — is
  **unaffected**. I verified this rather than assuming it: `AcademyClient.jsx:1363-1365` builds each bundle's
  `pathChildren` from `scopedChapters` and returns `[]` for any path with no chapters, so with `chapters: []` the
  audience views render bundle chrome and **zero** path cards. The defect is in the stated shape, not in the
  inference drawn from it. That is materially smaller than most of this union.
- **How it was repaired.** The false `=` literal was removed from the assertion; the shape is restated as
  measured, with a one-line note that the extra keys do not disturb `0 cards` and an explicit instruction not to
  "fix" the code to match the old doc. Both TRUE halves — the byte-exact `getServerCatalogView()` one-liner and
  `→ 0 cards` — were kept, per binding rule 5.

## `next-web-app.md:32` — one defect or two?

**ONE defect, booked twice by two independent adjudicators in two readings; iter-101 additionally named a second
site.** They are not two defects:

| | iter-99 row 11 (`recruiter-scoreboard-in-apps-web`, Adj2 · r21 G B1) | iter-101 row 22 (`recruiter-scoreboard-app`, adj-4 · r23-G B1) |
|---|---|---|
| anchor | `next-web-app.md:32` | `next-web-app.md:32` (supported at `:14`) |
| sentence graded | the same one | the same one |
| refuting mechanism | `UserStatusContext.tsx:141-173` — the all-hiring-orgs ejection | the same ejection, **plus** the corpus-internal contradiction at `hiring.md:53` / `:352` and the render site `apps/hiring/.../InsightsByMembersContainer.tsx` |

Identical anchor, identical sentence, identical root mechanism. The predicate names differ only because the two
readings were seated independently. **The only real delta is scope**: iter-101 widened it by one site (`:14`).
I treated it as one predicate over two sites and repaired both, then swept the file for the third-order
consequence — see *Noticed* below, items 1–2, which I **did** repair because they are in my file and would
otherwise have left the document contradicting its own new note two screens later (binding rule 4c).

Line numbers drifted with the 41-commit move: iter-101 cited `InsightsByMembersContainer.tsx:359`; at
`8297c684` the exported component is at **`:108`**. I cited the measured value, not the inherited one.

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `academy-progress-write-path` | 1 | **4** (1 mine + 1 same-file non-defect + 2 outside) | 1 of 1 in scope | `git grep -n -i "beacon"` and `git grep -n "upsertChapterProgress\|setLastActivity"` over `corpus/ CLAUDE.md`. `ant-academy.md:438` also names both mutations but makes **no** source attribution — inspected, not a site, not touched. Two twins are outside my files (below). |
| `empty-catalog-view-shape` | 1 | 1 | 1 | `git grep -n "emptyCatalogView\|serverTenant"` corpus-wide → 18 hits; only `ant-academy.md:135` asserts a *shape*. Confirmed with a literal search for the 3-key form: `git grep -n "skillPaths: {}, series: \[\] }"` → 1 hit, mine. |
| `localeswitch-surface-set` | 1 | 1 | 1 | `git grep -n "LocaleSwitch\|PublicHeader"` corpus-wide → 2 hits, both inside my one bullet. Ground truth by `git grep -n LocaleSwitch 22df69dd -- code/` (mount enumeration) + `git ls-tree -r` of the `(public)` route group. |
| `topbar-surface-set` | 1 | 1 | 1 | `git grep -n "TopBar\|LanguageSelector"` corpus-wide → 4 hits, all inside the same bullet. Ground truth by enumerating `TopBar` importers at the ref, then mapping every `app/**/page.jsx` to its client component — the route→client map is where the "3 routes, not 1" for `AcademyClient` comes from. |
| `recruiter-scoreboard-app` (= `-in-apps-web`) | 1 (booked twice) | **4** in my file (`:14`, `:17`, `:32`, `:119-121`) | 4 | `git grep -ni "candidate-comparison\|scoreboard\|activity-dashboard"` over `corpus/ CLAUDE.md`, then `git grep -ni "apps/web"` inside my file for the consequential clauses. The `:17` and `:119-121` sites are **rule-4c** repairs: they asserted "only `apps/web` is containerizable / hiring is not containerized", which my new note contradicts. |
| CANON-3 currency pin | 1 | 1 | 1 | `git grep -n "2035f9a"` + `git grep -n "origin/main\|origin HEAD"` restricted to my three files. One `2035f9a` label (`academy-backend.md:20`) and two bare `origin HEAD` labels (`next-web-app.md:49`, `:72`). |

**Honest residue: none inside my three files that I know of.** Two twins are outside them and were deliberately
not edited. The width booked for this seat (7) came in at **9 repaired sites**, i.e. ~1.3×, well under the ~3×
the brief warns of — because four of my seven anchors are genuinely narrow, single-bullet claims about one repo's
component tree, with no paraphrase surface elsewhere in the corpus. I checked for that surface rather than
assuming it.

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| `academy-progress-write-path` | `corpus/architecture/service_taxonomy.md:268` | *"Writes: per-user progress, bookmarks, certificates and feedback **POST through** `code/app/api/academy/beacon/route.js:36,41-55` (`UPSERT_CHAPTER_PROGRESS`, `SET_LAST_ACTIVITY`, …)"* — the identical inversion, with cited line numbers. Note the anchors `:36,41-55` **do** resolve (the beacon route really does hold those mutation docs); it is the *"the writes go through here"* framing that is false, which makes this twin harder to catch than mine. |
| `academy-progress-write-path` | `corpus/ops/demo/content-stories-routes.md:383` | *"**Write path:** `ant-academy/code/app/api/academy/beacon/route.js` posts `UPSERT_CHAPTER_PROGRESS` / …"* — labelled literally "Write path", the strongest form of the claim in the corpus. |

Neither is in `corpus/services/**`, so neither is in this union; both are live. Routing them is the
orchestrator's call.

## Noticed, not repaired

1. **`corpus/services/next-web-app.md` heading `### Containerized (Workforce only)`** — still accurate for the
   `make up-frontend` platform path it introduces, but the word "only" now sits above a paragraph that names the
   demo's `apps/hiring` container. I left the heading and let the paragraph carry the qualifier rather than
   rewriting a section title I was not asked to touch. Flagging it as a wording-tension, not a defect.
2. **`ant-academy.md` `5b05b7d9` / `e22f3230`** (the "dropdown mounted since" / "7-locale since" pins) sit inside
   a bullet I edited. They are **sha pins**, and a pin is a pin, so I did not re-derive or touch them — but I did
   not verify them either, and I am saying so rather than implying I did.
3. **`academy-backend.md:15`'s self-citation** originally read `:80-83`, whose first line was blank after my
   edit's +4 shift. I re-pointed it at the **construct** (`:85-89` — the subgraph + endpoint block it means)
   rather than mechanically adding 4, since a mechanically-shifted number landing on whitespace is the
   wrong-construct-citation class this milestone keeps finding.
4. **`app` `2035f9a4` is not itself at `v1.369.0`.** `git describe --tags 2035f9a4` → `v1.369.0-2-g2035f9a40` —
   it was already 2 commits past the tag when the old sentence was written. The sentence's reading ("the newest
   tag reachable") was defensible and I preserved that reading at the new ref, but the same looseness is now
   7 commits wide at `ad9f3c49`, which is why I published the full `git describe` output instead of a bare
   version string. Anyone tightening the corpus's version convention should start here.

## What I could not settle, and why

Nothing in my assignment was left unsettled. Every one of the 7 booked anchors was re-derived from source at the
named settling tree, and all 9 sites were repaired. Two things I bounded rather than resolved:

- **The demo's `apps/hiring` container** is sourced from rext (`stack-demo/rosetta-extensions` `09d06070` —
  `demo-stack/frontend/hiring.Dockerfile`, `demo-stack/up-injected.sh:1076-1085`, image `demo-<N>-hiring`).
  I verified the mechanism exists and is wired, and cited it as such. I did **not** verify the port-offset
  arithmetic or that a live demo currently serves it — that is a reading, and this is a repair pass.
- **iter-101's union header records the rext pin as `ab81527a`;** the per-stack tree I settled against is
  `09d06070`. I used the tree the brief names (`stack-demo/rosetta-extensions`), not the union's inherited
  value, and note the divergence rather than reconciling it.

## Guard-safety note on the retraction wording

Each repair quotes the claim it retracts, which is this corpus's convention — but a retraction that reproduces
the false string **verbatim and adjacently** would re-register as a fresh site under `claim_twin_guard`. I broke
the adjacency in three places on purpose: the 3-key literal is quoted without its `emptyCatalogView() = ` prefix;
the `/library` + `/free` gloss is described, not quoted; and `academy-backend.md` attaches *"the moving label
`origin/main`"* to the sha in a separated construction rather than reproducing the `@ origin/main 2035f9a4`
form. The false sentences quoted verbatim in the ledger above appear **nowhere** in the repaired corpus.

**Verified, not asserted.** I ran a fixed-string `git grep -F` for all 7 ledger quotes over `corpus/ CLAUDE.md`
after the edits. Six are **zero-hit tree-wide**. The seventh (`@ origin/main `2035f9a4``) still returns exactly
**two** hits — `corpus/services/ai-labs.md:18` and `corpus/services/coursebuilder.md:132` — which are precisely
the two sites CANON-3 assigns to **seat 6**, and no others. My three files are clear. That the residue lands
exactly on the assigned seat and nowhere unassigned is the useful signal here: for this predicate the
CANON-3 anchor table is complete over `corpus/services/**`, at least as far as my sweep can see.

**Not committed.** `git status --porcelain` over my three files shows ` M` only — no `git add`, `stash`,
`checkout --`, `reset`, or `fetch` was run at any point.
