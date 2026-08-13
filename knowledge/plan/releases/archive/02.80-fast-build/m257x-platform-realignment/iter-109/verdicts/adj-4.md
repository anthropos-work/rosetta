# Adjudicator 4 — seats D and F (readings #27 and #28)

**Re-derived independently from the platform clones. No fetch, no git state change, read-only.**
Refs used, verified with `git rev-parse` at this adjudication's open, all matching the brief's table:
`platform 0c91421d` · `app ad9f3c49` · `graphql-wundergraph 60c229f3` · `ant-academy 22df69dd` ·
`sentinel f2c46190` · `messenger fa47850d` · `cms ca50c817` · `storage 4ce8ece5` ·
`jobsimulation 462343b0` · `roadrunner 87d8d443` · `next-web-app 8297c684` · `studio-desk 41ee3575`.

**Which rext tree:** no booking in my set turned on `rosetta-extensions` state, so no `wrong-tree`
question arose. Nothing in this file is graded against either rext clone.

**Note on the expected cross-reading collapse:** the caller assigned both readings of each seat so that
duplicates would collapse. They did not duplicate — **seat D's #27 booked `service_taxonomy.md` and its
#28 booked `security_compliance.md`; seat F's #27 booked `ant-academy.md` and its #28 booked
`sentinel.md`.** Four readings, six bookings, six distinct anchors, no cross-reading overlap. Reported
because a zero-collapse result is itself a measurement about the instrument's reproducibility.

---

## Verdicts

### D r27 B1 — `archiveStudioTask` cited to the queries SDL

```
D-r27 B1 | corpus/architecture/service_taxonomy.md:48-49 | UPHELD | IN-SCOPE | PREDICATE: archiveStudioTask is declared in app's cms_queries.graphqls at line 106.
```

- **evidence (opened myself):** `stack-demo/app` @ `ad9f3c49`.
  `git grep -n "archiveStudioTask\|studioTasks\|studioTask(" HEAD -- internal/web/backend/graphql/graph/schemas/`
  returns exactly three lines in **two** files:
  - `cms_queries.graphqls:106` → `studioTask(taskId: ID!, organizationId: ID): StudioTask!`
  - `cms_queries.graphqls:107` → `studioTasks(`
  - `cms_mutations.graphqls:22` → `archiveStudioTask(taskId: ID!, organizationId: ID): Boolean!`
  I read the corpus block itself: `:48` names all three operations, `:49` supplies **one** locator,
  `app/internal/web/backend/graphql/graph/schemas/cms_queries.graphqls:106`. `archiveStudioTask` does
  not occur anywhere in `cms_queries.graphqls`. This is a **wrong file**, not line drift — a mutation
  cannot live in a `_queries` SDL, and a reader chasing the archive operation finds nothing there.
- **what survives:** the sentence's load-bearing proposition — *these operations are `app`'s, not
  studio-desk's and not a standalone cms's* — is TRUE, and I verified it. Only the locator is false, and
  only for one of the three named constructs. I considered and rejected the "exemplar line" reading: the
  citation is a precise `file:line`, not a directory pointer.

### D r27 B2 — `d11a403` credited with re-pointing all four `*_RPC_ADDR`

```
D-r27 B2 | corpus/architecture/service_taxonomy.md:425 | UPHELD | IN-SCOPE | PREDICATE: d11a403 re-pointed all four of messenger's *_RPC_ADDR at http://backend:8083.
```

- **evidence (opened myself):** `stack-demo/platform`. `git show d11a403 -- docker-compose.yml` changes
  exactly two values in the messenger block:
  `CMS_RPC_ADDR http://cms:8091 → http://backend:8083` and
  `JOBSIMULATION_RPC_ADDR http://jobsimulation:8401 → http://backend:8083`.
  At the parent `d11a403^` (= `2adcf71`) I dumped the messenger block's four `*_RPC_ADDR` lines
  (`:255`, `:256`, `:258`, `:265`): `BACKEND_USERS_RPC_ADDR=http://backend:8083` and
  `SKILLER_RPC_ADDR=http://backend:8083` **already** read `backend:8083` before the commit. So the
  commit re-pointed the **middle two**, not four.
- I also killed the charitable temporal reading ("by the time of `d11a403`, all four had been
  re-pointed"). `git log -S 'BACKEND_USERS_RPC_ADDR' --all -- docker-compose.yml` back to its
  introduction at `3e85fce` shows the value was `http://backend:8081` from birth and only ever moved
  ports — it **never addressed a service other than `backend`**, so there was nothing to re-point.
  The end-state reading fails on that variable too.
- **corpus-internal contradiction:** root `CLAUDE.md:282` states the precise form —
  *"`d11a403` had re-pointed the **middle two** at `backend` — M809"* — and this same file at `:160`
  names only `CMS_RPC_ADDR` for that commit. The corpus knows the distinction and `:425` states it
  wrong. The clause *"— the M809 re-point landed —"* forces the agentive reading.

### D r28 B1 — the Cosmo Router placed in the public subnets

```
D-r28 B1 | corpus/architecture/security_compliance.md:22 | UPHELD | IN-SCOPE | PREDICATE: The Cosmo Router is deployed in the VPC's public subnets.
```

- **evidence (opened myself):** `stack-demo/graphql-wundergraph` @ `60c229f3`, `terraform/main.tf`
  read in full for the `module "graphql"` block. It passes
  `private_subnets_ids = var.platform_private_subnets_ids` (`:31`) **and no public-subnet argument of
  any kind**. I re-derived the SET rather than sampling it: across **all eight** service terraform trees
  in the clone set (`app`, `sentinel`, `graphql-wundergraph`, `messenger`, `cms`, `roadrunner`,
  `storage`, `jobsimulation`) the token `public_subnet` occurs **0 times**, and **every one of the eight**
  passes the identical `private_subnets_ids = var.platform_private_subnets_ids`. Cardinality first:
  8 modules, 8 private-only, 0 public.
- **the router is not distinguished by anything the bullet could mean.** It uses
  `modules/services/base_service` — the same module as `app` (`app/terraform/main.tf:172`) — and `app`
  passes the same `private_subnets_ids` (`:194`) and the same ALB wiring. `backend` is filed by the very
  next line (`:23`, *"Private subnets: All microservices (no direct internet access)"*) as private. So
  the two bullets single out the router for a placement it shares with `backend`, which they place in
  the other subnet class.
- **residual, stated:** `infrastructure` (which defines `base_service`) has never been in a clone set,
  and `use_fargate = false` means the tasks land on cluster container instances I cannot see. I cannot
  prove the module does not place the task elsewhere. What I *can* measure — the only declaration that
  speaks to placement — says private, and the corpus asserts public without a ref. Upheld as
  contradicted-by-the-only-readable-evidence, not merely as unverifiable prose (which I would not have
  booked; the rest of that VPC block — CIDR, data subnets, NACLs — I left alone for exactly that reason).
- **second anchor, same predicate:** `corpus/architecture/architecture_overview.md:423`
  (*"public subnets (ALB, Cosmo Router), private subnets (all microservices)"*), which I opened and
  confirmed carries the identical claim. Seat F flagged that line as unsettleable rather than booking it.

### D r28 B2 — "Both bullets above" over a three-bullet list

```
D-r28 B2 | corpus/architecture/security_compliance.md:252 | REJECTED | — | PREDICATE (as booked): The EU-AI-Act retraction leaves the Limited-Risk classification bullet outside its fence.
   class: mis-read — the fence covers the classification bullet twice over, and the booking's harm-claim misquotes :7.
```

- **evidence (opened myself):** I read `security_compliance.md:1-12` and `:220-262` in full.
  The `### EU AI Act` list is indeed three bullets (`:227` classification, `:228` *"Stated reason:"*,
  `:229` *"Stated consequence of that classification:"*). So "Both" over three is loose. But the booking's
  consequence — that `:227` therefore reads as the corpus asserting a legal classification — does not
  survive:
  1. `:228` and `:229` both **self-label** with the word *"Stated"*. The two bullets that need an
     external fence are `:227` and `:229` (the latter named explicitly by the gloss *"including the
     consequence bullet"*, because the repair note at `:253-256` records that it had been sitting
     *outside* the blockquote). Under that resolution "Both" picks out `:227` + `:229` and the
     classification **is** inside the fence.
  2. Independently, `:248` — sixteen lines above, inside the same blockquote — already states
     *"**Do not cite this section as evidence of a Limited-Risk classification** — re-derive it."*
     That is a blanket fence over the classification bullet regardless of how "Both" resolves.
  3. The booking asserts `:7` *"states the classification unhedged in the summary."* It does not. The
     full sentence at `:7` reads *"AI Simulations are classified as **Limited Risk** under the EU AI Act
     — **but the stated reason for that classification does not hold at platform HEAD** … **The legal
     classification itself is a question for counsel; this corpus only records that the stated technical
     premise is false.**"* The seat quoted the first clause only. The corroborating site cuts the other
     way.
- **what is left** is a loose quantifier in a hedging blockquote — a wording nit of MINOR shape, not a
  false proposition about the platform or a live self-contradiction. Rule 5's carve-out applies in
  spirit: this is a retraction doing its job, not two live incompatible assertions.

### F r27 B1 — Ant Academy as an internal `@anthropos.work`-gated employee portal

```
F-r27 B1 | corpus/services/ant-academy.md:474 | UPHELD | IN-SCOPE | PREDICATE: Ant Academy is an internal @anthropos.work-domain-gated portal that external users cannot enter.
```

- **evidence (opened myself):** `stack-demo/ant-academy` @ `22df69dd`, files read with `git show HEAD:…`
  (the worktree is dirty with demopatch residue; every read below is at HEAD).
  - `code/src/lib/platformUrls.js:1-32` — *"**The Academy is a storefront in front of the Anthropos
    platform** (`app.anthropos.work`…). Account + billing live on the platform, so every conversion CTA
    leaves the Academy and lands on a platform page."* It then defines **FLOW A — Account gate**
    (*"The PRIMARY CTA for this flow — **a new visitor registers**"*) and **FLOW B — Checkout gate**
    (*"Checkout contextually registers → signs in → subscribes **for an anonymous visitor**"*).
  - `code/src/lib/pricing.js` — `STANDARD_YEARLY = { usd: 399, eur: 349 }` (`:22`),
    `LAUNCH_PROMO_ENABLED = true` (`:15`), `LAUNCH_COUPON_CODE = 'aiacademylaunch'` (`:19`), and a live
    *"$199 for everyone"* **community offer** with per-currency Stripe coupons. The header says the site
    *"ADVERTISES the price; the real charge happens at next-web-app checkout."*
  - `code/src/components/TopBar.jsx:77-88` — for `anonymous` visitors the top bar renders a
    `topbar-buy-cta` reading *"Buy AI Academy"* with a live price, firing a Meta `Lead` event and
    navigating to `platformCheckoutUrl(window.location.href)`.
  - `code/src/lib/schema.js:3` — `SITE_URL = 'https://aiacademy.anthropos.work'`, an
    `EducationalOrganization` schema.org block; `code/src/lib/landingSeo.js` +
    `landingSeoOverlay.js` exist as public marketing/SEO copy decks.
  - `knowledge/user-types.md` (in the ant-academy repo, **not** under any `knowledge/plan/` path) — *"The
    four user types"* are **Anonymous · Signed-in (free) · Subscriber · Enterprise/Org member**, each
    with an authoritative server-side signal. **No `@anthropos.work` predicate appears in the detection
    list at all.**
  A product that sells a $399/yr subscription to an anonymous visitor through a public checkout funnel
  is not one that *"external users cannot enter."*
- **self-contradiction inside the same corpus file** (I confirmed both sides): `ant-academy.md:316-317`
  documents anonymous browsing of `/`, `/latest`, `/chapters/*`, `/courses/*` and a `/library`, `/free`
  *"Phase-1 public launch"*, `:298` says *"**The public surface is much wider than 'a few auth pages'**"*,
  and `:261`/`:266` literally use the phrase **"public-storefront"** — 213 lines before `:474` denies that
  any external user can enter.
- **anchors collapsing onto this one predicate** (all opened and confirmed to carry it):
  `ant-academy.md:5` (*"internal learning portal … to anyone with an `@anthropos.work` email"*), `:24`
  (*"Clerk (`@anthropos.work` domain gate + org-membership gate)"*), `:31`
  (*"**Internal-only** learning portal … to `@anthropos.work` employees"*), `:474`; and
  `architecture_overview.md:40` + `:260` (*"Internal learning portal … for `@anthropos.work` employees"*).
- **the mitigation I checked and discarded:** the ant-academy repo's own `CLAUDE.md:11` still says
  *"AI Academy is an internal learning portal for Anthropic employees (`@anthropos.work` domain)"* — so
  the corpus is echoing a stale upstream doc rather than inventing it. That explains the error; it does
  not make it true. The corpus states it in its own voice, present tense, with no ref and no attribution,
  so rule 1 settles it at ground-truth HEAD. Likewise `ant-academy.md:298`'s *"enforced in the Clerk
  app"* is unreadable from any clone — but a Clerk-side domain restriction cannot coexist with a
  checkout flow whose own comment says it *"contextually registers … for an anonymous visitor."*

### F r28 B1 — `sentinel.md:5` asserts a `messenger` compose block that no longer exists

```
F-r28 B1 | corpus/services/sentinel.md:5 | UPHELD | IN-SCOPE | PREDICATE: The platform's docker-compose still contains a messenger service block.
```

- **evidence (opened myself):** `stack-demo/platform` @ `0c91421d`.
  `git show 0c91421d:docker-compose.yml | grep -nE '^  [a-z0-9_-]+:$'` enumerates **five** services —
  `sentinel` (`:5`), `backend` (`:28`), `studio-desk` (`:112`), `next-web-app` (`:143`),
  `gotenberg` (`:170`). `git grep -n messenger 0c91421d -- docker-compose.yml common.yml repos.yml`
  returns **only comments**: `docker-compose.yml:84`, `:91`, `:102` on the `backend` block and
  `repos.yml:5`. There is no messenger block. `838d907` (*"chore(compose): drop the storage, messenger
  and customerio-sync containers"*, 2026-08-05) deleted it; I read the commit body.
  The corpus sentence — *"its compose block sets no `AUTHORIZATION_ADDRESS` and declares no
  `depends_on: sentinel`"* — is present-tense and presupposes a block that does not exist.
- **ref-discipline check, done explicitly because this is where the class lives.** The parenthetical in
  the *preceding* sentence pins `0dab54d`. Per brief rule 1 a pin's scope is *the claim's own block — a
  wrapped sentence*; that pin sits in a different sentence, about a different subject (which blocks set
  `AUTHORIZATION_ADDRESS`). The messenger sentence carries no ref of its own and no past tense. It is
  therefore not the ref-discipline class: nothing here is *"pinned, past-tense, or dated"*. I verified
  the historical form separately — at `0dab54d` the messenger block does begin at `docker-compose.yml:156`
  and sets no `AUTHORIZATION_ADDRESS` — so the sentence *was* true and has silently expired.
- **live self-contradiction in the same file**, both sides opened: `sentinel.md:85` — *"The blocks that
  used to carry it are **gone** rather than corrected: … then `storage`, **`messenger`** and
  `customerio-sync` at `838d907`"* — and `:89` — *"`messenger` and `storage` never called it, and
  **neither is a compose service any more** (deleted at `838d907`)."* Neither is framed as a retraction
  of `:5`; both are asserted as current state alongside it. That is rule 5's "book it" branch, not its
  carve-out.
- **what survives:** the headline *"`messenger` is not a caller"* is TRUE and the third conjunct is TRUE.
  I re-derived it: `git -C stack-demo/messenger grep -n "authorization\|AUTHORIZATION_ADDRESS" fa47850d --
  '*.go' go.mod` returns one unrelated hit (`pkg/aireadinessemail/override.go`, *"fallbackName is the
  sentinel used when…"*), against `colony` present in `go.mod` as a positive control. The defect is the
  evidence clause, not the conclusion.

---

## PREDICATE ROLL-UP

```
P1 | archiveStudioTask is declared in app's cms_queries.graphqls at line 106.                        | anchors: D-r27 B1 @ corpus/architecture/service_taxonomy.md:48-49
P2 | d11a403 re-pointed all four of messenger's *_RPC_ADDR at http://backend:8083.                   | anchors: D-r27 B2 @ corpus/architecture/service_taxonomy.md:425
P3 | The Cosmo Router is deployed in the VPC's public subnets.                                       | anchors: D-r28 B1 @ corpus/architecture/security_compliance.md:22, D-r28 B1 @ corpus/architecture/architecture_overview.md:423
P4 | Ant Academy is an internal @anthropos.work-gated portal that external users cannot enter.       | anchors: F-r27 B1 @ corpus/services/ant-academy.md:474, :5, :24, :31; F-r27 B1 @ corpus/architecture/architecture_overview.md:40, :260
P5 | The platform's docker-compose still contains a messenger service block.                         | anchors: F-r28 B1 @ corpus/services/sentinel.md:5
```

No predicate in this set collapses with another: P1–P5 are five distinct false propositions across four
corpus files. **6 anchors booked → 8 anchor sites named → 5 distinct predicates.** The anchor:predicate
ratio is driven entirely by P3 (2 sites) and P4 (6 sites); the other three are one site each.

BOOKED=6 UPHELD=5 REJECTED=1 IN-SCOPE-UPHELD-BLOCKERS=5 DISTINCT-IN-SCOPE-PREDICATES=5
