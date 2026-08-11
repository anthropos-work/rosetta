# adj-3 — verdicts for seats D (r29-D, r30-D) and F (r29-F, r30-F)

**Every verdict below was re-derived by opening the platform file myself at the ref the claim names.**
No verdict rests on a seat's quoted evidence.

**Trees read.** All platform claims: the `stack-demo/*` clones at the brief's refs (`platform 0c91421d`,
`app ad9f3c49`, `sentinel f2c46190`, `messenger fa47850d`, `ant-academy 22df69dd`, `studio-desk 41ee3575`,
`next-web-app 8297c684`, `cms ca50c817`, `jobsimulation 462343b0`, `storage 4ce8ece5`,
`graphql-wundergraph 60c229f3`, `roadrunner 87d8d443`). No fetch, no checkout, no git state change;
`git status --porcelain` empty at open and at close.

**rext tree used:** the one clerkenstein booking (r30-D B3) is a claim about *what the tooling does on a
stack*, so it was settled in the **pinned per-stack consumption clone `stack-demo/rosetta-extensions
@ 09d06070`** — not the authoring copy. No fence verdict or DNA configuration was graded, so the authoring
copy was not needed.

---

## Seat D — r29-D

```
r29-D B1 | corpus/services/sentinel.md:5 | UPHELD | IN-SCOPE | PREDICATE: The published messenger grep at fa47850d returns one hit; it returns zero.
   evidence: stack-demo/messenger @ fa47850d — ran the sentence's own command verbatim:
     `git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod` → rc=1, NO OUTPUT.
     Repo-wide case-insensitive `git grep -in "authorization" fa47850d` → rc=1, no output.
     All three §5-rule-44 instruments checked before booking an absence: (1) plain
     `grep -ril authorization --include='*.go' .` → nothing (nothing gitignored-but-tracked is hiding a
     hit); (2) NUL scan over every tracked blob (`git show <ref>:<f> | tr -dc '\000' | wc -c`) → no
     tracked file in `messenger` carries a NUL byte, so neither `grep -I` nor `git grep` is skipping a
     file; (3) `find . -name .git -maxdepth 3` → only the repo root, no nested checkout.
     Positive controls in the same pass: `git grep -c colony fa47850d -- go.mod` → 1; 42 tracked `.go`
     files at that ref. The pipeline, the ref and the pathspec all work — the stated result simply does
     not reproduce. The claim is FALSE at the ref the claim itself names.
```

```
r29-D B2 | corpus/architecture/service_taxonomy.md:403 | UPHELD | IN-SCOPE | PREDICATE: service_taxonomy.md:142 is the archive-state note; it is not — that note is :148-160.
   evidence: corpus/architecture/service_taxonomy.md read with line numbers. :139-146 is the blockquote
     "⚠️ Two different fates shared this table…", and :142 reads "> so a bare `make up` started all three
     as unfederated husks — the `running_but_unfederated` state in" — a claim about merged-into-app vs
     gone-from-compose. The archive-state note is the NEXT blockquote, :148-160, opening at :148
     "**Every `ARCHIVED <date>` in this table is a DATED SNAPSHOT, not a derived fact.**" (grep for
     "DATED SNAPSHOT" → single hit, line 148). The pointer resolves to a real line that names the wrong
     construct. Not a historical anchor (rule 7): the sentence is a live "see the archive-state note …
     `:142`" cross-reference, not a record of where something once was.
```

```
r29-D B3 | corpus/architecture/service_taxonomy.md:407 | UPHELD | IN-SCOPE | PREDICATE: service_taxonomy.md:67-68 states "There is no `graphql` profile"; that sentence is at :74.
   evidence: `grep -n "There is no \`graphql\` profile" corpus/architecture/service_taxonomy.md` returns
     exactly two hits: :74 (the assertion, inside the Services paragraph — "…`gotenberg` and the three
     always-on base services. **There is no `graphql` profile, and no cms /") and :407 (the citing
     sentence itself). :67 is "- **Communication**: HTTP/RPC + Redis Streams" and :68 is the Database
     bullet ("PostgreSQL — one schema, `public`, owned by `app`…") — two members of the Tier-1
     *Characteristics* bullet list, neither of which mentions profiles. Distinct proposition from B2:
     different pointer, different target construct, different repair.
```

```
r29-D B4 | corpus/architecture/security_compliance.md:23-25 | UPHELD | IN-SCOPE | PREDICATE: The clone set holds eight service terraform trees; it holds ten, nine of them ECS services.
   evidence: I enumerated the SET first, not the sum (rule 4). Iterating every `stack-demo/*` clone for a
     `terraform/` directory: app · sentinel · graphql-wundergraph · messenger · cms · roadrunner ·
     storage · jobsimulation · **next-web-app** · **studio-desk** = TEN (ant-academy and platform have
     none). The passage names eight and asserts "all eight service terraform trees **in the clone set**"
     — a completeness claim over the clone set. `studio-desk` is squarely on-predicate, not an outlier:
     `git -C studio-desk show 41ee3575:terraform/main.tf` — `:11`
     `source = "github.com/anthropos-work/infrastructure.git//modules/services/base_service?ref=main"`
     (the same module the passage cites for the router and for `app`), `:13` `use_fargate = false`,
     `:33` `private_subnets_ids = var.platform_private_subnets_ids`. So the true set of ECS `base_service`
     terraform trees in the clone set is NINE. (next-web-app's terraform is a Vercel project — `:11`
     `resource "vercel_project"` — genuinely off-predicate, no subnets at all.)
     Measured across all ten: `public_subnet` = 0 files in every tree, so the passage's CONCLUSION is
     unaffected and in fact reinforced; what is false is the enumerated set the word "all" ranges over.
```

## Seat D — r30-D

```
r30-D B1 | corpus/architecture/service_taxonomy.md:175 (+ :518) | REJECTED | — | PREDICATE: The private Go modules a stack's Docker build pulls are exactly colony, proto, taxonomy.
   evidence: The measurement the seat reports is correct — `git -C stack-demo/app show ad9f3c49:go.mod`
     lines 14-18 are `analytics-go v0.3.1`, `colony v0.35.2`, `proto v1.210.0`, `storage v0.15.2`,
     `taxonomy v1.2.0`, and `sentinel f2c46190:go.mod` adds none of the missing two. But I opened
     service_taxonomy.md around :175 and the sentence's domain is not "all anthropos-work modules": the
     paragraph OPENS with its own domain — "**Shared Libraries** (imported as private Go modules — not
     cloned by `make init`; pulled at Docker build via `GH_PAT`/`GOPRIVATE`)" — and heads the five-row
     Shared Libraries table (colony, proto, ai, authn, taxonomy) immediately below it. The clause
     immediately before the booked one binds the pronoun: "**`ai` is NO LONGER among them**". Read as
     "of the five shared libraries, three are still built", the sentence is TRUE, and that is exactly how
     its three siblings phrase it — `:518` "**5 libraries, 3 imported by a service a stack builds**",
     `corpus/architecture/dependency_map.md:42` "Five repos exist; the two Go repos a stack clones and
     builds … pull in **THREE of them**", `corpus/architecture/README.md:21` "only **three are imported
     as private modules by a service a stack builds**". `storage` and `analytics-go` sit outside that
     domain, and the corpus states them correctly where the domain IS all of `app/go.mod`
     (`external_services.md:554`, verified verbatim) — so there is no cross-document contradiction either.
   class: mis-read — the enumeration's domain is the five shared-library repos, fixed by the paragraph's
     own header and table and by three sibling sites; it is not a universal over every private module.
```

```
r30-D B2 | corpus/architecture/service_taxonomy.md:403 | UPHELD | IN-SCOPE | PREDICATE: service_taxonomy.md:142 is the archive-state note; it is not — that note is :148-160.
   evidence: same re-derivation as r29-D B2 (:139-146 vs :148-160, "DATED SNAPSHOT" single hit at :148).
     Same anchor, same predicate — collapses onto P2.
```

```
r30-D B3 | corpus/services/clerkenstein.md:171-176 | UPHELD | IN-SCOPE | PREDICATE: clerkenstein's handleMeOrganizationMemberships discards or ignores the *http.Request.
   evidence: pinned per-stack clone stack-demo/rosetta-extensions @ 09d06070 (the tree that settles what
     the tooling does on a stack). `clerkenstein/clerk-frontend/server.go` signatures:
     `:241 handleClient(w, _ *http.Request)`, `:467 handleToken(w, _ *http.Request)`,
     `:488 handleMe(w, _ *http.Request)` — three genuinely discard it — but
     `:512 func (s *Server) handleMeOrganizationMemberships(w http.ResponseWriter, r *http.Request)`
     BINDS `r` and reads it twice: `:527 strconv.Atoi(r.URL.Query().Get("offset"))` and
     `:531 strconv.Atoi(r.URL.Query().Get("limit"))`, which page the response at :530/:532. The hedge
     "(or ignore)" does not save it — it parses request input and changes the response body on it.
     Self-contradiction confirmed in the same document: :131-134 asserts as a correctness requirement
     that "**`limit`/`offset` are honoured** — clerk-js sends `limit=10&offset=0`…", which is only
     possible by reading the request. Both sides are asserted live, so rule 5's retraction exemption
     does not apply.
```

```
r30-D B4 | corpus/services/sentinel.md:85 → corpus/architecture/architecture_overview.md:335 | REJECTED | — | PREDICATE: architecture_overview.md:335 holds the quoted "only cross-process RPC edge" line.
   evidence: I opened architecture_overview.md. `:335` is "**On a local stack** (platform `2adcf71`
     deleted the router — **there is no `:5050`**):" — the lead-in that INTRODUCES the fenced model; the
     fence opens at :337 and the quoted string is at :339, inside it (grep "only cross-process RPC edge"
     → single hit, line 339). The citing sentence reads "The correctly-scoped form is **the model** at
     `architecture_overview.md:335`" — it names a CONSTRUCT (the model), and :335 is where that very
     construct begins; the quote is four lines into the block it opens. That is line-level imprecision
     INSIDE the right construct, not a pointer onto a different one — materially unlike B2/B3/F-B3,
     where the cited line sits in a different blockquote / different bullet list about a different
     proposition. Seat r29-D, reading the identical anchor, graded it a MINOR for exactly this reason.
   class: mis-read — same construct, not a wrong one.
```

```
r30-D B5 | corpus/architecture/security_compliance.md:197 → README.md:21 | REJECTED | — | PREDICATE: The README line cited for "cost tracking is in app, not the ai library" does not support it.
   evidence: I opened the cited target under the natural sibling-relative resolution
     (`corpus/architecture/security_compliance.md` → `corpus/architecture/README.md`). Line 21 is the
     `shared_libraries.md` bullet and it CLOSES with: "…where its responsibilities begin and end (e.g.
     **cost tracking lives in `app`, not the `ai` library**)." That is precisely the proposition the
     citation cross-references ("cost tracking in `app/internal/aiusage` — **not** the shared `ai`
     library"). The seat's objection is that the literal token `aiusage` appears only at :23 — but :23 is
     the `ai_architecture.md` bullet, and the citation's own second half is "+ `ai_architecture.md`", so
     the pair "(consistent with `README.md:21` + `ai_architecture.md`)" resolves correctly on both legs.
     Seat r29-D opened the same line and cleared it for the same reason.
   class: mis-read — the cited line does state the proposition it is cited for.
```

## Seat F — r29-F

```
r29-F B1 | corpus/services/ant-academy.md:31 (co-anchors :5, :298) | UPHELD | IN-SCOPE | PREDICATE: Ant Academy is an internal-only portal restricted to `@anthropos.work` employees.
   evidence: re-derived at ant-academy 22df69dd, reading blobs at the ref (the working tree is dirty from
     an applied demo-patch, so `git show <ref>:<path>` throughout):
     · `code/src/lib/pricing.js:21` `export const STANDARD_YEARLY = { usd: 399, eur: 349 };` (+ a live
       launch coupon at :18, discount at :24).
     · `code/src/lib/platformUrls.js:4` "The Academy is a **storefront** in front of the Anthropos
       platform"; :17-21 "FLOW B — Checkout gate … Checkout contextually registers → signs in →
       subscribes **for an anonymous visitor**".
     · `code/src/components/TopBar.jsx:77` `{anonymous ? (` guarding :80-87, a "Buy AI Academy" CTA with
       price that navigates an ANONYMOUS visitor to `platformCheckoutUrl()`.
     · `code/src/lib/schema.js:3` `const SITE_URL = 'https://aiacademy.anthropos.work';`
     · `knowledge/user-types.md` — the four user types are Anonymous / Signed-in (free) / Subscriber /
       Enterprise-Org-member, detected by `auth().userId`, `isInOrg()`, `billingPremium()`. I read the
       whole "How each type is detected" block: NO `@anthropos.work` predicate anywhere in it.
     · `git grep -in "anthropos\.work" 22df69dd -- code/src code/proxy.js` returns only host URLs
       (`app.anthropos.work`, `aiacademy.anthropos.work`), dev-login help text and e2e fixture emails —
       no email-domain gate.
     So a paid public storefront, not an internal-only portal. And the SAME FILE books the proposition
     false and live at :474: "⚠️ **'Domain-gated to `@anthropos.work` so external users cannot enter' is
     FALSE and was removed at M257x iter-115.** … **This document contradicted itself 213 lines
     earlier**". Rule 5's retraction exemption does not apply — the retraction and the retracted claim
     are BOTH asserted live (:5 and :31 are present-tense, un-hedged). :298's "`@anthropos.work` domain
     restriction is enforced in the Clerk app" is the same proposition on an artifact outside every clone
     (Clerk dashboard config); it is unsupportable from source and refuted by :474, and folds in as a
     co-anchor rather than a separate predicate.
```

```
r29-F B2 | corpus/architecture/architecture_overview.md:40 (co-anchor :260) | UPHELD | IN-SCOPE | PREDICATE: Ant Academy is an internal-only portal restricted to `@anthropos.work` employees.
   evidence: :40 "**Ant Academy** (`ant-academy`): Internal learning portal (Next.js 16 + Expo mobile)
     for `@anthropos.work` employees."; :260 the same in the Frontend Applications table. Identical
     source re-derivation as B1 (pricing.js:21, platformUrls.js:17-21, TopBar.jsx:77-87, schema.js:3,
     knowledge/user-types.md, the zero-hit domain-predicate grep, all @ 22df69dd). Same underlying false
     proposition at a different anchor → collapses onto P6, which is what the roll-up records. I also
     enumerated the predicate's full in-scope footprint myself
     (`grep -rn "[Ii]nternal.*learning portal\|[Ii]nternal-only" corpus/architecture/ corpus/services/`):
     six live sites in five files — architecture_overview.md:40, :260, service_taxonomy.md:290,
     frontend_architecture.md:9, ant-academy.md:5, :31 — plus corpus/services/README.md:58. All one
     predicate.
```

```
r29-F B3 | corpus/services/roadrunner.md:130 | UPHELD | IN-SCOPE | PREDICATE: roadrunner.md:124 holds "Upstream consumers: none (orphaned)"; that line is :134, and it is above :130, not below.
   evidence: corpus/services/roadrunner.md read with line numbers. :124 is "* The repo contains an
     experimental WebSocket LSP proxy (`internal/lsp/lsp.go`) that is NOT wired into any running server
     — there is no reachable LSP endpoint today." `grep -n "Upstream consumers"` returns two hits: :130
     (the citing sentence's quotation) and :134, "* **Upstream consumers**: **none (orphaned — see the
     banner at the top).**" So the pointer resolves onto a real line naming a different construct, and
     the direction word "below" is wrong in the bargain (:124 is six lines ABOVE the citing sentence).
     The substantive claim it supports (roadrunner is orphaned) is separately true and was not booked.
```

## Seat F — r30-F

```
r30-F B1 | corpus/architecture/architecture_overview.md:80 | UPHELD | IN-SCOPE | PREDICATE: `ai` is one of the private Go modules a stack's Docker build pulls.
   evidence: the sentence names NO ref, so it grades at the checkouts. The two Go repos `repos.yml`
     @ platform 0c91421 lists and a stack builds are `app` and `sentinel`.
     `git -C stack-demo/app show ad9f3c49:go.mod` — the anthropos-work requires are exactly
     `analytics-go v0.3.1`, `colony v0.35.2`, `proto v1.210.0`, `storage v0.15.2`, `taxonomy v1.2.0`;
     **no `github.com/anthropos-work/ai`**. `git -C stack-demo/sentinel show f2c46190:go.mod` — `colony
     v0.35.2`, `proto v1.210.0`, `taxonomy v1.2.0 // indirect`; no `ai`. So of the five Shared Libraries,
     three are pulled at Docker build and `ai` is not one of them — the claim "**four** imported private
     modules — colony, proto, **ai**, taxonomy" is false on the member it adds.
     Corroborated as a live self-contradiction against four siblings that all measure the other way:
     `dependency_map.md:42` ("pull in **THREE** of them — colony, proto, taxonomy … `app` **dropped** the
     `ai` module at `1e457fa70`") and its `ai` row at :48 ("**No repo a stack builds**");
     `service_taxonomy.md:175` ("**`ai` is NO LONGER among them**"); `architecture/README.md:21`;
     `external_services.md:554`. `:80` is the lone survivor.
     (Distinct from the r30-D B1 predicate I rejected: that one alleged the enumeration is short by
     `storage`/`analytics-go`; this one alleges it contains a member that is gone. Different
     propositions, different repairs, so they are not deduplicated together.)
```

```
r30-F B2 | corpus/services/roadrunner.md:130 | UPHELD | IN-SCOPE | PREDICATE: roadrunner.md:124 holds "Upstream consumers: none (orphaned)"; that line is :134, and it is above :130, not below.
   evidence: same re-derivation as r29-F B3 (:124 = the LSP bullet; :134 = the Upstream-consumers
     bullet). Same anchor, same predicate — collapses onto P7.
```

```
r30-F B3 | corpus/services/ant-academy.md:63 | UPHELD | IN-SCOPE | PREDICATE: The academy persists chapter progress via the upsertChapterProgressBatch mutation.
   evidence: ant-academy @ 22df69dd, `code/src/graphql/query/academyProgress.js` read in full.
     The module exports exactly three documents: `GET_ACADEMY_PROGRESS` (:23), `UPSERT_CHAPTER_PROGRESS`
     (:54) and `SET_LAST_ACTIVITY` (:73). Between the last two, :68-70 reads verbatim: "(v0.5 M2 §6:
     UPSERT_CHAPTER_PROGRESS_BATCH removed — it was the FE outbox flush, consumed only by the deleted
     sync transport. Progress is now an immediate per-chapter write through UPSERT_CHAPTER_PROGRESS
     above.)". `git grep -in "upsertChapterProgressBatch\|UPSERT_CHAPTER_PROGRESS_BATCH" 22df69dd`
     returns, outside that removal note, only the module's own STALE header at :18-19 ("progress flushes
     via the BATCH mutation") — refuted three comment-lines later at :50-53 — plus
     `code/src/graphql/schema.graphql:2207`, the BACKEND's SDL, and `knowledge/**` planning prose. There
     is no client call site. So the `[Batch]` half of the corpus sentence is false of the subject the
     sentence is about (how the academy persists progress).
     Self-contradiction inside the same corpus file, both live: the blockquote at :65-77 enumerates the
     write path as exactly two mutations (`UPSERT_CHAPTER_PROGRESS` at store.js:162, `SET_LAST_ACTIVITY`
     at store.js:210), and :483-484 restates it with no `[Batch]`.
```

```
r30-F B4 | CLAUDE.md:294 | UPHELD | OUT-OF-SCOPE | PREDICATE: The academy React app reads env only from `code/.env.local`.
   evidence: root `CLAUDE.md:294` — "…(not the repo root — **the React app reads only from
     `code/.env.local`**)". Measured at ant-academy 22df69dd: the repo's own guide says the opposite —
     `CLAUDE.md:164` "Copy `.env.example` to `.env` and fill in.", `CLAUDE.md:129` "…with
     `CLERK_SECRET_KEY` set in `code/.env`." `code/.gitignore:22-24` ignores `.env`, `.env.local` AND
     `.env.*.local`, i.e. all three are expected to exist locally. And the in-scope corpus file states
     the correct form: `corpus/services/ant-academy.md:391` "The **app's** env file is `code/.env`, not
     the repo root", with `cp .env.example .env` at :395 and `REQUIRE_ORGANIZATION_MEMBERSHIP=0` "in
     `code/.env`" at :412. The claim is false and contradicts the service doc — but the anchor is the
     repo-root `CLAUDE.md`, which is neither `corpus/services/**` nor `corpus/architecture/**`, so under
     rule 8 it does not enter N or P. (`corpus/ops/setup_guide.md` carries the same `.env.local` recipe
     at root CLAUDE.md:494's pointer — also out of scope.)
```

---

## PREDICATE ROLL-UP

```
P1 | The published messenger grep at fa47850d returns one hit; it returns zero. | anchors: r29-D B1 @ corpus/services/sentinel.md:5
P2 | service_taxonomy.md:142 is the archive-state note; it is not — that note is :148-160. | anchors: r29-D B2 @ corpus/architecture/service_taxonomy.md:403, r30-D B2 @ corpus/architecture/service_taxonomy.md:403
P3 | service_taxonomy.md:67-68 states "There is no `graphql` profile"; that sentence is at :74. | anchors: r29-D B3 @ corpus/architecture/service_taxonomy.md:407
P4 | The clone set holds eight service terraform trees; it holds ten, nine of them ECS services. | anchors: r29-D B4 @ corpus/architecture/security_compliance.md:23-25
P5 | clerkenstein's handleMeOrganizationMemberships discards or ignores the *http.Request. | anchors: r30-D B3 @ corpus/services/clerkenstein.md:171-176
P6 | Ant Academy is an internal-only portal restricted to `@anthropos.work` employees. | anchors: r29-F B1 @ corpus/services/ant-academy.md:31 (+:5, +:298), r29-F B2 @ corpus/architecture/architecture_overview.md:40 (+:260); same predicate also live at corpus/architecture/service_taxonomy.md:290, corpus/architecture/frontend_architecture.md:9, corpus/services/README.md:58
P7 | roadrunner.md:124 holds "Upstream consumers: none (orphaned)"; that line is :134, above :130 not below. | anchors: r29-F B3 @ corpus/services/roadrunner.md:130, r30-F B2 @ corpus/services/roadrunner.md:130
P8 | `ai` is one of the private Go modules a stack's Docker build pulls. | anchors: r30-F B1 @ corpus/architecture/architecture_overview.md:80
P9 | The academy persists chapter progress via the upsertChapterProgressBatch mutation. | anchors: r30-F B3 @ corpus/services/ant-academy.md:63
```

**Out-of-scope upheld (not counted in N or P):** `CLAUDE.md:294` — "the React app reads only from
`code/.env.local`" (r30-F B4).

**Rejections by class:** mis-read ×3 (r30-D B1, r30-D B4, r30-D B5). No ref-discipline rejection arose in
this seat group — no booked claim was pinned, past-tense or dated in a way the booking ignored. No
wrong-tree rejection arose: the one rext booking was correctly settled in the pinned per-stack clone and
the verdict does not turn on which tree was read (both agree on that file).

---

BOOKED=16 UPHELD=13 REJECTED=3 IN-SCOPE-UPHELD-BLOCKERS=12 DISTINCT-IN-SCOPE-PREDICATES=9
