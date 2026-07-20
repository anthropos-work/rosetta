# The demo-coverage protocol — the Playwright sweep + triage + fix loop

**The iteration protocol for the demo's prove-by-render gates.** Both are `iterative`: each iter **measures**
(run the sweep) → **triages** the failures → **fixes** them in `rosetta-extensions` (or a corpus doc) →
**re-sweeps**, until the gate is GREEN. The page set and the failure modes are **discovered by the sweep**,
not enumerable up front, which is why the milestone shape is iterative rather than sectioned.

**This doc governs TWO sweeps.** They share the loop, the harness foundation, the fix-surface routing table,
and the measurement conventions; they differ in what they enumerate and how a page is reached:

| | **Sweep 1 — hero vantage coverage** | **Sweep 2 — content stories (session × action)** |
|---|---|---|
| **added by** | v1.10 "method acting" M42e (employee) + M42m (manager) | v2.5 "the playbill" M236 |
| **proves** | *a vantage's pages are full* | *every seeded content story lands on a real result* |
| **unit** | a **page**, reached by BFS crawl from the hero's landing | a **(seat × exact URL)** pair — never crawled |
| **actor** | a roster **hero** (Maya, Dan) | a non-hero **MEMBER** seat (`content-player-<idx>`) + a manager hero |
| **enumerated by** | the sweep itself (discovered frontier) | the seeded `content-manifest.json` (a fixed denominator) |
| **section** | [Sweep 1](#for-pms--what-100-coverage-means) below | [Sweep 2](#sweep-2--content-stories-the-session--action-lands-sweep-v25-the-playbill-m236) below |

**Read them in order.** Sweep 1 establishes the loop, the gate vocabulary, and the harness; Sweep 2 assumes
all of it and documents only where a fixed-denominator, exact-path sweep diverges from a crawl. The
[cross-cutting protocol](#cross-cutting-protocol--rules-both-sweeps-owe-to-hard-experience) at the end holds
the rules that outlived either one.

> **The demo-patch mechanism is specified in [`demopatch-spec.md`](demopatch-spec.md).** It is the sanctioned **zero-platform-edit escape hatch**: patch the demo's own ephemeral clone before the image build, revert after — the canonical repos are never touched. Read it before adding or re-pinning a patch. Since M217 the gate is **self-healing**: the *anchor* is the contract, the whole-file sha is only a baseline.

> **Read first:** [`frontend-tier.md`](frontend-tier.md) (the UI tier the sweep crawls), [`verification.md`](../verification.md)
> (the offset/project/scope-aware probe net this harness sits beside + reuses), [`rosetta_demo.md`](../rosetta_demo.md)
> (the demo lifecycle + offset ports + Clerkenstein injection the cockpit login rides), and
> [`stories-spec.md`](stories-spec.md) (the roster hero — Maya for employee — the sweep logs in as).
>
> All the harness code lives in the gitignored `rosetta-extensions` monorepo (authored + tagged in the
> authoring copy, consumed per-stack at a pinned tag) — **zero platform-repo edits** (the v1.10 hard line).
>
> **Function sibling:** [`playthroughs.md`](playthroughs.md) (v2.0 "opening night" M202) proves what this sweep
> does not — that a hero can **do** the thing, not just that the page **shows** content. A Playthrough plays a
> real journey end-to-end and asserts the outcome (function); this sweep proves presence. Playthroughs extend
> this same M42 Playwright foundation (the `cockpit-login` handshake + the centralized-anchor discipline) into a
> mutating, reset-to-seed, serial suite with its own manifest + four-state map.

# Sweep 1 — hero vantage coverage (v1.10 "method acting" M42e / M42m)

**A coverage milestone proves that a hero of a given vantage** (employee/member, or manager), logged into a
demo stack via the presenter cockpit, **sees 100% of the pages that vantage can reach** rendered with **real
semantic content** and **zero out-of-demo escapes**.

## For PMs — what "100% coverage" means

After a presenter clicks **Login as {hero}** in the cockpit, every page that hero can click to should be
**full** — real text, real rows, a real chart — and every link should stay **inside the demo** (clicking
"Studio" lands on the demo's own studio, not the public prod site). The coverage sweep is a robot that logs in
as the hero, walks every reachable page, and flags any page that comes up **empty**, **errors**, or **links
out of the demo**. The milestone is done when the robot reports **zero** such pages.

## The gate (objective, machine-verifiable) — the re-scoped SEMANTIC gate (M42e iter-21)

> **Re-scope (2026-06-25).** The original gate measured **DOM text-density** (`textLen > 40` in `<main>`).
> It was too weak — it passed pages that render placeholder/empty-state cards ("add something here") + nav
> chrome, so the harness reported a green `(0,0)` while a logged-in presenter saw an empty profile, an empty
> AI-sim library, incoherent 3D-Dental skills for a backend dev, a silhouette avatar, and no org logo. The
> gate below is the **believability bar**: real semantic content + substantial per-section cardinality +
> persona self-consistency + no prod-eject, reproducible on a FRESH demo-up. It is measured by a
> **manifest-driven semantic harness**, not the text-density one.

A Playwright sweep, logged in as the vantage's hero via the cockpit handshake, of **every reachable demo
page** asserts — **per page AND per section/element**:
- **(a) Real semantic content** — actual seeded user/catalog content; **placeholder / empty-state copy and
  bare chrome do NOT count** (the `empty-states.ts` denylist). A section whose region selector is **not
  found** is a **FAIL** (an absent section can't silently escape the gate).
- **(b) Substantial cardinality** — each content section shows a **meaningful count** of items (≥ its
  manifest floor, not just 1), **except** documented exceptions where 0/1 is the genuinely-correct state.
- **(c) Persona self-consistency** — the hero's **role ↔ skills ↔ bio ↔ a real-photo avatar (consistent
  across the menu AND the profile) ↔ work history** cohere as one believable person; the **org has a name +
  a logo**.
- **(d) No prod-eject escape** — no in-app nav / menu / button ejects the presenter to a **prod anthropos
  surface** (e.g. left-menu "Studio" → `studio.anthropos.work`). Legitimate **external editorial / reference
  links inside content** (a `/chapter` citation; a LinkedIn-import help link) are **allowed but disclosed**
  in a presenter-notes list — they are not prod-ejects.
- **(d′) Cross-port demo-local outbound surfaces actually WORK (v1.10 "method acting" postfix).** A demo-local
  link to a **different port** than next-web's — e.g. the manager's **"Anthropos Studio"** left-nav →
  studio-desk on `:9000+offset` — is now **FOLLOWED** by the crawl with the logged-in context and asserted to
  land on a **real, non-blank** destination: on the destination `host:port`, **not** a `/login`·`/sign-in`·
  `/undefined` surface, **not** the un-offset `:3000`, HTTP 2xx (no `ERR_TOO_MANY_REDIRECTS`), with a
  destination DOM marker present (the studio-desk home's WelcomeSection / QuickActions). Previously the BFS only
  recursed **same-host-AND-port**, so these cross-port destinations were never visited and a blank /
  login-looping studio-desk slipped the gate. A failing follow is a gate **FAIL** (the studio-desk class).

**Gate = `0 failing sections + 0 persona failures + 0 prod-eject escapes + 0 not-reached manifest pages +
0 cross-port-follow failures`, over a FRONTIER-EXHAUSTED crawl (`cappedAtFrontier === false`), reproduced on a
FRESH demo-up.** A
**coverage-review.html** (per-section verdicts + screenshots + documentedExceptions[] + presenterNotes[]) is
emitted for human review — screenshot review is part of acceptance.

## The harness (where it lives, how it runs)

The coverage harness **extends the existing** [`stack-verify/e2e/`](../../../README.md) Playwright surface —
**it is not a from-scratch new dependency**. `stack-verify/e2e/` already pins `@playwright/test ^1.49.0` with a
`playwright.config.ts` keyed off `ROSETTA_E2E_BASE_URL` and an unauthenticated `smoke.spec.ts`. The coverage
sweep is the **authenticated, multi-page, semantic-content + escape-detection** evolution of that smoke test,
added as a sibling spec/runner under `stack-verify` so it reuses verify's offset/project/scope plumbing
(`STACK_PROJECT` / `STACK_OFFSET`, `lib/target.sh`).

> **Playwright is the first non-Go rext dev/test dependency** — sanctioned by the coverage requirement, pinned
> under `stack-verify/e2e/` with its own `package.json` + lockfile (already present). The Go rext tooling is
> untouched (supply-chain stays GREEN — no Go go.mod/go.sum change).

> ### ⚠️ WHERE you run the sweep is part of the test (M219, v2.3 "cue to cue")
>
> **Run the sweep from the vantage a PRESENTER has — a tailnet peer. Never from the demo host itself.**
>
> Until M219 both runners (`run-coverage.sh`, `run-playthroughs.sh`) hardcoded their app/FAPI bases to
> `localhost`, so a demo living on a remote tailnet VM **could not be swept at all**. That is a large part of
> *why* the AI-readiness asserts sat unrun for four releases. They now take `COVERAGE_HOST` /
> `COVERAGE_APP_SCHEME` (and `PT_HOST` / `PT_APP_SCHEME`).
>
> But pointing them at `localhost` **on the demo box** is not merely awkward — it is **wrong**, and it fails in
> a way that looks exactly like a product bug:
>
> - A `--public-host` demo **bakes the MagicDNS origin into the frontend build**, so the app's own GraphQL
>   client calls `https://<magicdns>:<15050+offset>/graphql`.
> - `docker-proxy` binds `0.0.0.0`, so a connection **from the demo host** to its own `100.x` tailscale IP hits
>   the kernel socket and **bypasses `tailscale serve`** — the thing that terminates TLS. Plain HTTP then
>   answers a TLS handshake: `ERR_SSL_PROTOCOL_ERROR` / *"wrong version number"*.
> - Every GraphQL call fails ⇒ every page is a permanent loading spinner ⇒ **every section reports
>   `region-not-found`** and the persona checks fail for want of an org name and an avatar.
>
> Measured on `billion`: from the host, https on `:13000`, `:15050` **and** `:18082` all fail TLS; **from a
> tailnet peer all three answer.** The demo was healthy throughout. The first M219 sweep run this way reported
> `failingSections=21, personaFailures=3` — a **systemic false-RED**, and exactly the sort that gets "fixed" by
> weakening asserts. From the correct vantage the same build reported `failingSections=0, personaFailures=0`.
>
> **A sweep that cannot reach the app does not report "broken" — it reports "empty", and empty is the one
> result this protocol forbids you to read as anything.** The `--reset-only` flag exists so the DB half (which
> needs docker + `stackseed`, i.e. the host) and the browser half (which needs the peer) can run on different
> machines.

The harness, against a **live** demo on offset ports:
1. **Logs in** as the vantage's roster hero via the **cockpit handshake** — the demo's fake FAPI deep-link
   `https://<fapi-host>/v1/client/handshake?…&__clerk_identity=<hero-key>` selects the hero's seat
   (`clerk-frontend/server.go::handleHandshake`) and establishes the RS256 session as that hero; the next-web
   app picks up the `__session` cookie. (`ignoreHTTPSErrors: true` for the openssl-fallback FAPI cert; the
   mkcert path is browser-trusted — see [`recipe-browser-login.md §B step 2`](recipe-browser-login.md).)
2. **Crawls** the in-app nav as that vantage — **pure in-app nav-link discovery** (BFS from the seed paths
   over same-origin links + the persistent nav chrome), NOT a static route manifest. The crawl is now
   **reachability + escape-classification ONLY** (`lib/crawl.ts`, M42e iter-21): it discovers which pages the
   vantage can navigate to (the gate's scope) and classifies every `<a href>` host as demo-local / prod-eject
   / allowed-external — the content VERDICT moved OUT of the crawl and INTO the manifest (step 3). Rationale
   for keeping discovery (not a manifest) for the scope: a route manifest can't catch a nav that **escapes**
   the demo. The frontier is capped + deduped; query-only variants of the same path collapse. **The BFS recurses
   same-host-AND-port only**; a **cross-port demo-local** link (a different port than next-web — e.g. the
   "Anthropos Studio" left-nav → studio-desk `:9000+offset`) is collected separately and **FOLLOWED after the
   BFS** (`onCrossPortFollow`, v1.10 postfix) — asserting the destination is non-blank / not a sign-in or
   `/undefined` surface / not the un-offset `:3000` / HTTP 2xx / carries a destination DOM marker (gate (d′)).
3. **Per page** runs the **manifest's per-section asserts** (`lib/section-assert.ts` against
   `lib/coverage-manifest.ts`): for each section the manifest declares for that page — resolve its region
   selector (0 matches → **region-not-found = FAIL**), reject error/skeleton/empty-state content
   (`lib/empty-states.ts`), and assert real text (mustInclude + a meaningful-length floor after stripping
   empty-state copy) and/or **cardinality ≥ floor**. One bounded re-assert after an extra settle distinguishes
   slow-paint from genuinely-empty. Captures a **screenshot** inline. (c) the crawl classifies links;
   (d)-failing prod-ejects are the off-demo links the allow-rule did NOT clear.
3a. **Persona self-consistency** (`lib/persona-assert.ts`): role ↔ skills coherence (the allow-set is derived
   from the hero's OWN rendered role skill-panel at sweep time — the platform resolving `job_role_skills` —
   not a hand-list; a junk-pool denylist catches the flat-pool head bug), avatar **menu == profile** + is-a-
   real-photo (a raster data-URI / image, NOT a silhouette placeholder SVG / initials), org **name + logo**.
4. **Emits** a coverage report JSON (`{reachable, failingSections, personaFailures, escapes, notReachedPages,
   cappedAtFrontier, gateMet, pages[], persona[], documentedExceptions[], presenterNotes[]}`) **and a
   `coverage-review.html`** — per-section verdicts + the per-page screenshot + the documented-exception list
   + the presenter-notes list, for human review.

### The semantic-gate assertion shape — the MANIFEST model (M42e iter-21)

The gate's (a)+(b) clauses are the false-pass/false-fail risk; the manifest (`lib/coverage-manifest.ts`)
replaces the old text-density floor with **per-page, per-section DESCRIPTORS**. Each descriptor is
`{ id, region selector, realContent assertion, minCount floor, exception?+reason }`:

- **region** — a CSS/text selector locating the section's container. **0 matches → `region-not-found` →
  FAIL** (the key re-scope property: an absent section can't silently escape).
- **realContent** — one of `text` (mustInclude substrings + a meaningful-length floor, measured AFTER
  stripping the empty-state phrases in `empty-states.ts`), `count` (cardinality of an itemSelector inside the
  region), or `both`.
- **minCount** — the **cardinality floor** ("substantial, not just 1"), calibrated against what the seed +
  set-dress actually produce (one calibration sweep) so the floors are achievable, not new false-fails.
- **exception + reason** — set where 0/1 is the **genuinely-correct** state (e.g. a terse Settings menu); the
  floor relaxes to "real, non-empty content" and the reason is surfaced in the review's
  `documentedExceptions[]` (honest disclosure, never a silent skip).

> **⚠️ AN EXCEPTION DESCRIBES A FINDING. IT MUST NEVER PASS ONE (M219 R-8).** The clause above says an
> exception is for where **0/1 is the genuinely-correct state**. That is the *only* legitimate use, and it was
> violated: `ai-readiness-interview-findings` carried an exception disclosing that **three findings blocks
> rendered their HEADINGS WITH NO CONTENT** and **a fourth did not render at all** — because *"the data they
> summarize is written by no seeder."* Empty was **not** the correct state. It was a **seeder gap**, honestly
> named, and then **shipped green for four releases** while the gate reported *"8/8 sections pass."*
>
> A disclosed empty is better than a hidden one. **It is still an empty section.** The discipline:
>
> | The exception is… | …when |
> |---|---|
> | ✅ **legitimate** | the surface is *genuinely* terse and 0/1 items is what a real user *should* see (a Settings menu with 2 entries) |
> | ❌ **a finding wearing a badge** | the section is empty because **something upstream is missing** (no seeder writes it, a serve-grant is absent, a link is unrewritten). **Fix the cause and DELETE the exception.** |
>
> The M219 R-8 test of the distinction, which is worth stating as a question: *"if this were fixed, would the
> section fill?"* If yes, it is **not** an exception — it is a **deferral with better prose**. The
> interview-findings exception was deleted, the four blocks are seeded
> (`corpus/services/ai-readiness.md` → `interview_aggregated_reports`), and the descriptor now asserts **all
> four blocks are FILLED** — including a **900-char content floor**, because the four empty headings measured
> ~120-200 chars and the old floor of **120** passed straight over them.
>
> **The same class, in the launcher:** the ant-academy liveness probe proved the port **SERVES** (200 to curl)
> and reported *"started + SERVING"* over a page that, in a browser, redirects to Clerk's keyless handshake and
> **renders nothing**. *"It serves"* ≠ *"it renders"* — the same sentence as *"the role classifies"* ≠ *"the pool
> is big enough"* and *"it resolves"* ≠ *"it has skills"*. Both that probe and
> `ANT_ACADEMY_HOME_SECTION` (which asserted a meaningless **40-char** floor with **no required markers** — the
> very `textLen>40` density check this protocol claims to have superseded) now assert the **render**. They report
> **RED until the render is fixed in M220 — which is intended.** An accurate red beats a comfortable green.

> **The ZERO-CELL blind spot — and the assert that closes it (M219, v2.3 "cue to cue").** A `text` descriptor
> asserts the section's **labels**. But a KPI tile's labels are **chrome**: they render whether the numbers
> behind them are `1,284` or `0`. So a section can clear a `mustInclude` + length floor while every value in it
> is zero — **an empty section wearing a hat**, and the sweep calls it a pass. It did: the AI-readiness
> *"✨ Handled for you this cycle"* tile shipped with a hard-`0` counter and nothing noticed, because nothing
> asserted a **value**.
>
> The fix is the existing **`textMatch`** kind, pointed at the values instead of the labels: a pattern requiring
> a **leading non-zero digit** in front of each counter label, with `minMatches` set to the **number of
> counters** — so `0 AI skills mapped` matches zero times and the section reports **empty**, and *one* dead
> counter among three still fails (`minMatches: 3`, not "at least one survived"). Its unit tests assert the
> pattern **rejects** a zeroed tile and **accepts** a filled one — an assert nobody has ever seen fail is not a
> fence.
>
> **Reach for `textMatch`-on-values wherever a section's content is a NUMBER.** `text` proves a section
> *rendered*; only a value assert proves it rendered something **true**.

There are **two manifest namespaces**: **employee** (M42e — Maya, the member vantage, fully calibrated) and
**manager** (M42m — Dan, the org-intelligence vantage; **fully calibrated** as of M42m iter-04). The manager
manifest covers the real **`/enterprise/*`** route surface (reconciled from the wrong `/workforce/*` guesses):
`/enterprise/workforce` (the M36 Workforce-Intelligence dashboard — ONE **tabbed SPA** route, NOT `?tab=`
sub-routes: the funnel + org-scale gap + the Growth / Skills & Verification / Talent Pool / Assignments /
Activity Log tabs render in-page), `/enterprise/members`, `/enterprise/assignments`,
`/enterprise/activity-dashboard`, `/enterprise/organization-feedback` (the ~2:1 feedback distribution — needs
the M42m FeedbackSeeder mirror fix to render; see `stories-spec.md`), and `/enterprise/settings` (a documented
terse exception). The manager vantage has **two manager-only fan-outs** (`/user/<uuid>` team roster +
`/enterprise/activity-dashboard/.../<uuid>` drill-downs) AND **inherits the employee Library families**
(`/sim/<slug>`, `/skill-path/<slug>(/chapter)`) because the manager nav links the Library — so the manager
`SAMPLE_RULES` are a **superset** (the 2 fan-outs + the 2 library families), or the crawl explodes + times out.

> **The manager manifest is ORG-CONDITIONAL (v1.10b "fit-up" M53 AB4; routes + shape CORRECTED in v2.3 M219).**
> The **`/ai-readiness`** page (its **8** descriptors + its seedPath) is a **showcase-org-only** prime: the
> AI-readiness cycles seed only for **Northwind Aviation** (the M51 showcase org — which since M219 carries
> **both** a `closed` and an `active` cycle), so the dashboard is LEGITIMATELY empty on any base-Workforce org
> (Cervato / Solvantis) — asserting it there is a false-fail. `manifestFor(vantage, expectedOrg, identityKey)`
> therefore returns the full **`MANAGER_MANIFEST`** (which primes + asserts the AI-readiness page) only when
> `expectedOrg` contains `AI_READINESS_SHOWCASE_ORG` (`'Northwind Aviation'`, case-insensitive substring); for
> any other manager org (or an empty/undefined org) it returns **`MANAGER_MANIFEST_BASE`** — the same surface
> MINUS the AI-readiness seedPath + descriptors. `coverage.spec.ts` threads `COVERAGE_EXPECTED_ORG` in.
>
> ⚠️ **M219 corrected three things this block used to say.**
> 1. **The route was the LEGACY one.** It said `/enterprise/workforce/ai-readiness` — an **unlinked orphan**
>    (no nav entry, no workforce tab, no redirect points at it). The CURRENT manager surface is **`/ai-readiness`**.
>    This protocol was the last place the dead pointer survived after M219 repointed every other one.
> 2. **"two descriptors"** → **8**.
> 3. **"The employee manifest is org-independent"** is now **FALSE**. `manifestFor` is **3-arg**: the employee
>    vantage is **IDENTITY-conditional**, because the member readiness surface **has no route of its own** — it is
>    embedded in `/home`. Seats `aria-completed` / `ben-started` get `EMPLOYEE_MANIFEST_READINESS_DONE` /
>    `_STARTED`, each asserting an `/home` readiness section (`ai-readiness-member-done` / `-progress`). Every
>    other seat gets the base employee manifest. *A member surface with no route is exactly why route-crawling
>    never found it.*
>
> The original M53 split fixed an M51 regression: an unconditional AI-readiness seedPath (M51 iter-05) had
> silently broken the M50 base-org manager gate (`dan-manager` @ Cervato) — surfaced only by M53's from-cold
> both-vantage assertion. The page's real proof (the funnel renders from real seeded data) still holds on the
> showcase org, so the split removes a false assertion, not a real one.

> **The HIRING vantage — org/identity-conditional dispatch, on a SECOND app (v2.4 "casting call" M225).** The
> recruiter/candidate hiring surfaces render in the demo's **second UI container — the real `apps/hiring` app**
> (the M224 two-app demo, offset **3001**-port), NOT next-web: `apps/web` **ejects** an all-hiring-orgs recruiter
> to the hiring app by design (`UserStatusContext`), and the hiring app's symmetric guard keeps her in. So
> `manifestFor` gains, via the **same 3-arg org/identity dispatch as the AI-readiness precedent**:
> - **manager @ the HIRING org** (`HIRING_ORG` = `'Meridian Talent'`, case-insensitive substring, checked
>   **before** the showcase-org branch) → **`MANAGER_MANIFEST_HIRING`** (the recruiter Rae): the Results
>   scoreboard `/enterprise/activity-dashboard` — the 5 shared positions render as **custom tanstack-table** rows
>   (`tbody.tbody > tr.tr`, NOT AntD — the M224 render-probe R4 finding) + the `isHiring` **"Results" re-skin**.
> - **the HIRING candidate seats** (`cara-assessed` / `cody-assigned`, employee vantage) →
>   `EMPLOYEE_MANIFEST_HIRING_ASSESSED` / `_ASSIGNED`: the candidate **`/home` self-views** (apps/hiring `/profile`
>   is admin-gated → the platform redirects a candidate to `/home`). Per-candidate role↔score self-consistency:
>   the **assessed** candidate shows a completed+scored position; the **assigned-only** candidate a pending one.
>
> Two knobs point the sweep at the hiring app: **`COVERAGE_APP_PORT_BASE=3001`** (`run-coverage.sh` → the hiring
> app base) + **`COVERAGE_PROFILE_GATED=1`**, which puts **`persona-assert` in `profileGated` mode**: because
> apps/hiring's `/profile` family admin-redirects, the role-skills + avatar checks read the **`/home`** self-view
> (assert **no flat-pool junk** + a **real-photo** avatar) instead of the next-web `/profile*` pages they would
> otherwise false-fail on — a **shared-lib extension, not a fork**. The cohort-level role↔skills↔score
> self-consistency (0 junk names + a non-degenerate score distribution across the ~43-candidate cohort) is proven
> by the M224 `render-hiring-comparison.spec.ts` render probe; this manifest asserts the compare-surface LIST
> believability + the candidate self-views. **0 prod-eject** is the sweep's own escape scan, unchanged.

#### The documented-exception table (where 0/1 is legitimately correct)

| Vantage | Page · section | Exception | Reason |
|---|---|---|---|
| employee | `/settings` · settings-menu | floor relaxed to real-content (no cardinality floor) | A thin account / security / subscription menu — terse by design; a substantial-content floor would be a false-fail. The section must still render real menu text (not a skeleton/error). |

(A new exception is added here + in the manifest descriptor's `exception.reason` whenever a section's 0/1
state is proven correct — it is disclosed in `coverage-review.html`, never a silent scope-out.)

## The iter loop (Phase A–E)

Each tik runs these phases. The protocol's **primary metric** is the pair `(failing-pages, escapes)`; the
gate is `(0, 0)`. "Progress" = a net reduction in either component (or a net increase in **pages-reached** that
doesn't add failures — a deeper crawl that holds coverage).

- **Phase A — Sweep (measure).** Run the harness against the live demo as the vantage's hero. Produces the
  coverage report (the pre-iter metric). On iter-02 (the first tik) this is the **baseline sweep**.
- **Phase B — Triage.** Classify each failing page / escape by **fix surface** (the routing table below). Pick
  the highest-leverage cluster (most pages unblocked per fix) as this iter's target. Record the triage in the
  iter's `progress.md`.
- **Phase C — Fix.** Land the fix in the routed `rosetta-extensions` surface (or a corpus doc). **Re-apply**
  the affected stack step against the live demo (re-seed / re-replay-snapshot / re-export-roster+restart-fapi
  / re-build-frontend), per the fix surface — the demo must reflect the fix before re-sweep.
- **Phase D — Re-sweep (re-measure).** Re-run the harness. Record the post-iter `(failing, escapes)` + the
  per-cluster delta. A claimed lift MUST be a re-sweep delta, never an un-run assertion.
- **Phase E — Close.** Grade the iter (`closed-fixed` / `closed-fixed-partial` / `closed-no-lift`) on whether
  the targeted cluster's failures cleared; route the rest forward; write the close section with the mandatory
  `**Gate:**` field.

### Fix-surface routing table (the triage map)

| Failure mode | Root cause | Fix surface (rext) | Re-apply step |
|---|---|---|---|
| **Empty section / missing seed** | the page reads data the seeder never wrote | `stack-seeding` (seed the rows the page reads) | re-seed the demo |
| **Entitlement/policy-gated empty page (a deny modal)** | the page gates on a Sentinel Casbin policy the seeder never wrote (e.g. `/sim/.../start` deny modal when the org lacks the `FEATURE_JOB_SIMULATIONS` `g3` grant) | `stack-seeding` (seed the `g3` feature grant per membership — `identity.go`/`users.go`) | re-seed the demo **+ reload Sentinel policy** (restart `<demo>-sentinel-1` — `LoadPolicy()` runs once at startup, no watcher) |
| **Federation / content error (403/500/panic)** | the page reads replayed content not serve-granted | `stack-snapshot` serve-grants (`directus/structure.go`) | re-replay snapshot into the demo |
| **Directus schema-drift content 500 — COLUMN class** (M46/DD, Option A) | a cms `SetFields("*", …)` content query (e.g. simulations) SELECTs a column the **captured Directus structure** lacks because the platform added it after the capture → `Directus 500: column <collection>.<col> does not exist` → the content fetch fails (the ~60–90 s "latency" is the router **retrying**). **Cache-masked** in a warm sweep; surfaces only on a **cold** federation tier (restart cms+router+directus). DIAGNOSE via `docker logs <stack>-directus-1 \| grep 'does not exist'` + diff the **full `*`-expanded SELECT** Directus generates against the replayed physical columns (the full SELECT lists every requested column before execution → the COMPLETE missing COLUMN set in one pass, not bounded by Postgres reporting only the first). | a reproducible **post-replay column backfill** — an idempotent `ALTER TABLE directus.<collection> ADD COLUMN IF NOT EXISTS <col> <type> [DEFAULT …]` in `demo-stack/up-injected.sh`'s `NO_SETDRESS` block (next to the FK indexes, on the demo's own offset Postgres, schema `directus`), gated on local content + `DEMO_NO_DIRECTUS_DRIFT_FIX`, non-fatal, values-blind (the FK-indexes mechanism class). DEMO-LOCAL DDL — the `cms`/`app` clones stay pristine. **Scope:** column drift ONLY — NOT the serve-grant closure row below. | re-up (the backfill runs post-replay) — verify on a **COLD** tier (restart cms+router+directus) so it isn't cache-masked |
| **Directus serve-grant CLOSURE gap — RELATION/COLLECTION class** (M46/DD → **CLOSED by M46 Path 2**) | a cms deep-fetch (`GetJobSimulation`: `sequences.knowledge.*`, `sequences.assets_files.directus_files_id.*`, `sim_features.*`, `translations.*`, …) traverses a target/junction collection (`knowledge_asset`, `sequences_files`/`_2`, `directus_files`, `sim_features`, `sim_translations`, `simulations_translations`, `sim_roles_tasks`) the M40 [serve-grant](#…) `servedCollections` set does NOT register/grant/relate → Directus drops the **whole parent alias** (e.g. `sequences`) → cms `s.Sequences[0]` **panics** (`index out of range`) → a federated non-nullable field (`jobSimulation.title`) is null → the whole section (e.g. activity-dashboard's activity-table) never hydrates. DIAGNOSE via `probe-empty.spec.ts` (the `insightsByJobSimulations.rows.@.jobSimulation` DOWNSTREAM_SERVICE_ERROR) + `docker logs <stack>-cms-1 \| grep 'index out of range'` + check `directus.directus_{collections,relations,permissions}` for the traversed collections. **NOT an Option-A column backfill** (an `ADD COLUMN` won't help) and the relation metadata must be CAPTURED from prod (never hand-fabricated — the M25 subtle-FK-bug risk). **⚠ cms caches `GetJobSimulation` per-id responses in Redis DB 5 (`simulations_<id>_<hash>`, 24 h TTL, cache-FIRST) — so a re-replay into an ALREADY-running demo can serve a poisoned EMPTY-sequences entry cached during the serve-grant settle. A FRESH `/demo-up` starts empty + provisions directus before cms queries it (no poison); to fix in place, clear DB 5 `simulations_*`.** | **The fix (M46 Path 2):** EXPAND `servedCollections` in `stack-snapshot/directus/structure.go` to the full deep-fetch closure (the 7 collections above) + a SYNTHESIZED `directus_files` SYSTEM read grant (`serveFilesCollection`/`serveFilesPermissionSQL`) + **RECAPTURE** the prod Directus structure (the relation/field metadata is captured, never fabricated). | **a FRESH `/demo-up`** off the regenerated cache (the capture-path live-acceptance pattern — re-capture + cache-bust + fresh up); on an already-running demo, re-replay the serve rows + **clear the cms Redis DB-5 `simulations_*` cache** |
| **Out-of-demo link (escape)** | a baked/rendered link host points at prod | the demo **injection + env link-rewriting** (`demo-stack/up-injected.sh` build-args / `stack-injection/gen_injected_override.py`) — rewrite the host to the offset port. **Precondition (M42m iter-02):** the platform must expose a **per-URL `NEXT_PUBLIC_<thing>_URL` override** for that host (rewritable in the gitignored `apps/web/.env.local` overlay or a build-arg, zero-edit — e.g. next-web's `ACADEMY_URL` reads `NEXT_PUBLIC_ACADEMY_URL`). If the host is instead behind a **coarse mode-flip** (`NEXT_PUBLIC_NODE_ENV`) or a **hardcode** with no per-URL knob (e.g. next-web's `STUDIO_URL` — a `NEXT_PUBLIC_NODE_ENV` ternary, wrong-port + side-effecting on flip), the host is **platform-bound** → this row does NOT apply; it's a **re-scope trigger** (the rewrite needs a platform-source edit). Diagnose: find the constant's source, check for a dedicated `NEXT_PUBLIC_<thing>_URL` read vs a mode-flip/hardcode. | re-build the frontend (baked URL) or re-emit the override + restart |
| **Platform-bound escape (no per-URL override)** | a baked link host is hardcoded / behind a coarse mode-flip with no `NEXT_PUBLIC_<thing>_URL` knob (e.g. next-web's `STUDIO_URL`) — the env-rewrite row above does NOT apply | the **demo-patch tool** (`demo-stack/patches/demopatch` + a content-anchored manifest, M42m iter-03): source-patch the demo's **EPHEMERAL gitignored clone** before the build to read `NEXT_PUBLIC_<thing>_URL` (a behavior-identical fallback ternary kept), then **trap-revert** after the image bakes — CANONICAL repos NEVER touched (**7** guards, incl. G7 the apply post-condition: hard path-assert demo-clone-only, drift-refuse, never-commit, idempotent, self-owned reversal, demo-only, apply post-condition). Wired into `up-injected.sh` (apply-before-build + RETURN-trap revert) with the offset value in the `.env.local` overlay; default-on + non-fatal (`DEMO_NO_PATCH=1` opts out). The clone is left git-clean; `ensure-clones.sh` **R1** pristine-reverts a crash-left patch + **R1b** sweeps a crash-left tooling `.dockerignore` (byte-identical + untracked guards). Resolved the Studio `studio.anthropos.work` escape demo-only (139→0). | re-build the frontend (the patch bakes into the image; revert is automatic) |
| **Cross-port demo-local surface blank / login-loops / wrong-eject** (studio-desk class, v1.10 postfix) | a demo-local link to a **different port** than next-web (e.g. "Anthropos Studio" → studio-desk `:9000+offset`) doesn't render the authenticated home for the **logged-in hero** — it opens a blank `/undefined`, a dead `:3000` redirect-loop, a `/login` loop, or ejects to `WEB_APP_URL` (the non-admin redirect when the hero's membership doesn't resolve) | **authenticate via Clerkenstein, never bypass**: studio-desk drives its **own** fake-FAPI handshake (per-app — the cross-port `__session` cookie is **not** needed; the FAPI holds the active seat server-side) verified networklessly via `CLERK_JWT_KEY`. The **injection override** (`gen_injected_override.py`) wires the runtime `CLERK_*` (← `DESK_CLERK_*` minted by `up-injected.sh`) + pins `CLERK_SIGN_IN_URL`/`WEB_APP_URL` at the offset next-web (the requireAuth fallback). For the **admin gate** (`checkEnterpriseAndAdmin` reads the fake BAPI's `getOrganizationMembershipList`), make the fake BAPI **roster-aware** (`cmd/fake-bapi` reads the same `FAKE_FAPI_ROSTER` + seeds each hero's `(org, user)→org_role`) so a manager passes and an employee is correctly redirected. The harness gate is closed by the crawl's **cross-port FOLLOW** (`crawl.ts` `onCrossPortFollow` → `coverage.spec.ts`) **in the logged-in context**: a blank/login-loop/un-offset-`:3000`/eject destination FAILS the gate | re-up (the roster re-seeds + the override re-emits, no rebuild) for the env/BAPI; **clear the cached image** (`docker image rm demo-N-studio-desk`) only to re-bake a stale pk/offset |
| **Org-scale grid perf wall (slow GraphQL, not empty)** (M46) | at org scale (~500 members) a heavy enterprise grid (`/enterprise/members`, `/enterprise/activity-dashboard`, `/enterprise/settings`) sits on a `…` spinner / skeleton through the whole warm-grid poll because its backing **federated GraphQL** takes 10–84 s — an over-broad fetch (`InsightsContext.tsx` loads `limit: 1000` = ALL members) AND a per-membership resolver fan-out (`jobRole`/`targetRole`/`tags` × a **per-object Sentinel RPC** in `app` `roles.go GetOrganizationTargetRole`, no DataLoader). DIAGNOSE it's perf-not-content by `docker logs <stack>-graphql-1` + `grep latency` (10 s+) while raw SQL is ms (so it's RPC-round-trip-count, not the DB). **Decompose it — part demo-patchable, part platform-bound:** | **The over-broad-fetch + missing-index part is demo-local (zero canonical edit), and clears the grids whose cost is fetch-width or DB-index-bound:** **(1)** a **next-web pagination demo-patch** (`patches/next-web-members-pagination`: `InsightsContext.tsx` `limit 1000→30` — the grids that already paginate, e.g. `/enterprise/members` at 20, are untouched; the CSV/email export uses a separate query so no data is lost) applied to the demo's ephemeral clone pre-build + trap-reverted (the demopatch 6-guard contract); **(2)** **post-seed FK indexes** (`CREATE INDEX IF NOT EXISTS` on `membership_skills(membership_skill_membership)` + `membership_tags(membership_tag_membership)`) on the demo's **own** Postgres via a rext-owned `up-injected.sh` step (idempotent, non-fatal — NOT a canonical ent/atlas change). On demo-3 (~500 members) (1)+(2) took graphql max latency **84 s → ~4 s** and cleared `/enterprise/activity-dashboard` + `/enterprise/settings`. **BUT the `/enterprise/members` per-row `targetRole` → `OrgCheckActionPermission` Sentinel RPC is `OrgActionAssignmentsWrite`, which is PER-OBJECT (per assignee), NOT an org-wide grant** — a manager legitimately can-write-assignments for some members and not others. CACHING that check is a **correctness bug** — keyed `(org, subject, action)` (object dropped to dedupe across rows) it returns the first row's allow/deny for all rows → `failed to get target role: forbidden` on legit members (~1744×/sweep) → the grid errors (this was T2, reverted); keyed correctly `(org, subject, OBJECT, action)` it cannot dedupe (every row = a different object). **The safe demo-patch is to DROP the read-gate, not cache it (Option B, the M46 close):** `roles.go RoleManager.checkPermission` short-circuits `return true, nil` BEFORE the per-member Sentinel RPC (mirroring its built-in `privacy.DecisionFromContext` bypass) — target roles still come from the DB so every member's REAL role renders (fast, fully-populated, 0 forbidden). Read-path relaxation ONLY (`patches/app-targetrole-authz-skip`, applied to the build-scratch app clone via a rext helper wired into the inject loop, svc=app, after `apply-authn`, before build, trap-reverted) — the assignment **mutations** still enforce via their own direct `OrgCheckActionPermission` calls. On demo-3 B took the members query **76.7 s → 0.51 s** and **cleared `/enterprise/members` → the manager gate is MET (no re-scope)**. The PLATFORM finding stays: prod needs a **DataLoader / batch `BulkCheckPermission` RPC**; B is a disclosed demo-perf relaxation, not a prod fix. **⚠ if you rebuild the injected `app` image, it MUST go through the inject loop so `apply-authn.sh` (the disarmed colony) is re-applied** — a rebuild without it ships a backend that hits real `api.clerk.com`, rejects every Clerkenstein token, and collapses the crawl to `reachable≈7` (a broken-auth artifact, not a content fail — grep `docker logs <stack>-backend-1` for `clerk`); and **never recreate one service with `--force-recreate` *without* `--no-deps`** (it recreates `postgresql` too and wipes the seeded org). | re-build the frontend (pagination bakes in) + re-build the injected `app` (the authz-skip bakes in) via the inject loop + re-up/re-seed (indexes apply post-seed) — all three demo-local, gate MET |
| **Replayed-CONTENT URL-field escape** (v1.10b "fit-up" M50) | a prod host (`https://[*.]anthropos.work/...`) is **baked into a replayed Directus content field**, not built from a JS constant — e.g. `directus.simulations.public_landing_page_url` / `read_more_link` (28 / 14 sims carried a prod URL), surfaced when the activity-dashboard sim drill-down renders the field as a link → prod-eject. **Distinct from BOTH the JS-constant "Platform-bound escape" row (a built-in-the-bundle constant — fixed by the demopatch; M50's `next-web-public-website-url` demopatch is an instance, killing the `PUBLIC_WEBSITE_URL`-built links) AND the serve-grant "Federation/content error" row (a 403/500, not a working link to the wrong host).** Diagnose: the escape URL matches a replayed row's *field value*, AND the JS-constant ejects are already 0 (the demopatch worked) — so the residual lives in the replayed DATA, not the code. | a **post-replay content-URL rewrite** in `demo-stack/up-injected.sh`'s `NO_SETDRESS` block — an idempotent demo-local `UPDATE … regexp_replace(<field>, 'https?://[a-z0-9.-]*anthropos\.work', '<demo next-web host:3000+offset>')` over the `anthropos.work`-bearing content fields (`directus.simulations.{public_landing_page_url,read_more_link}` + `directus.skill_paths.public_landing_page_url`) — the content-side analog of the injection link-rewriting for app constants. Same class as the M46 FK indexes / Directus column backfill: demo-local DDL on the per-stack Directus (the `cms`/Directus clones stay pristine), idempotent (a re-run matches 0 rows), non-fatal (M18/M19), gated on local content, `DEMO_NO_CONTENT_URL_REWRITE` opt-out. A **REGEX** over any `anthropos.work` subdomain (not a bare prefix) catches prod **and** `staging.anthropos.work`. | re-up (the rewrite runs post-replay; if fixing in place, clear the cms Redis DB-5 `simulations_*` cache like the serve-grant row) — then re-sweep → escapes → 0 |
| **Editorial citation in replayed content** (NOT an escape) | a real `<a href>` to an external article baked into replayed `/skill-path/.../chapter` body copy | the harness **citation allow-rule** (`coverage.spec.ts` `allowedExternalLink` → `crawl.ts`): classify the off-demo link on a `/chapter` path as a VALID citation, recorded as a **presenter note**, NOT counted as an escape (M42e iter-08) | (none — content fidelity; do NOT strip/rewrite the citation) |
| **Wrong identity / org on a surface** | roster/FAPI resource gap | `stack-seeding/seeders/roster.go` + `clerkenstein/clerk-frontend/resources.go` | re-export roster + restart `<demo>-fake-fapi-1` |
| **A documented gap** | the behavior is correct but undocumented | a corpus doc update | (none) |

### Re-scope trigger (the zero-edit line)

If a 100%-blocking failure can **ONLY** be closed by a **platform change** (next-web / app / cms /
jobsimulation are read-only this entire release — the skills/taxonomy domain is part of `app` since the
skiller→app merge), that is the milestone's **Re-scope trigger**:
**escalate** (exit `EXIT_REASON: re-scope-trigger`), **record it** in the milestone-root `decisions.md`, and
**do NOT edit the platform repo**. The user decides whether to (a) carve the page out of the vantage's gate
with a documented rationale, (b) own an upstream platform PR out-of-band, or (c) pivot strategy.

## Iter type selection (protocol refinements over the skill floor)

The generic `build-mstone-iters` tik/tok cadence applies. This protocol adds:
- **The bootstrap tok (iter-01)** authors this protocol doc + scaffolds the harness + resolves the overview's
  four open questions (harness home, crawl strategy, assertion shape, Playwright wiring) + takes the baseline
  reading framing. It does NOT terminate the call.
- **Tooling-iter** — if a tik closes blocked on a **missing harness capability** (the sweep can't reach or
  classify a page because the crawler/assertion lacks a feature), the next iter is a **tooling-iter**: ship
  the harness capability AND use it within the same iter for the coverage work. (Carries the
  `iter_shape: tooling` frontmatter.)
- The **3-consecutive-no-progress-tik** triggered-tok floor is unchanged (revise the triage/fix strategy).

## Measurement conventions

- **The sweep is the measurement.** `(failing-pages, escapes)` from the coverage report is the only metric.
  The roll-up is deterministic given a fixed seed/snapshot, so a re-sweep with no fix should reproduce the
  same numbers (a moved number with no fix = a flake to investigate, not a lift).
- **Raise the page cap until the BFS frontier EXHAUSTS before reading the residual (M42e iter-07 lesson).** The
  crawl is bounded by `COVERAGE_MAX_PAGES`. If the crawl stops because it HIT the cap while the queue was still
  non-empty (`cappedAtFrontier===true`, `reachable===maxPages`), the `(failing, escapes)` it reports are
  **FLOORS over a truncated slice**, not the true residual — unreached pages may carry more failures/escapes.
  A `(0,0)` over a truncated frontier is structurally **not gate-met** (the gate is over the FULL reachable
  set). The report carries `cappedAtFrontier` + `frontierRemaining` + `maxPages`, and the runner emits a CAP-HIT
  warning, so a cap-saturated `reachable===maxPages` can never be misread as a true page count (the run-1
  verification's exact mistake). **Raise `COVERAGE_MAX_PAGES` until `cappedAtFrontier===false` (queue empty),
  THEN quote the residual.** The vantage's reachable set is what its in-app nav actually LINKS — e.g. the
  employee/Maya vantage exhausts at ~87 pages (skill-paths + the sims linked from the library + profile +
  home), NOT all 300+ sims in the catalog (most sim detail pages aren't crawl-reachable nav links from that
  vantage). If a vantage genuinely links a huge template-identical set, a representative + boundary sample with
  a documented rationale is defensible — but the frontier where escapes/failures live MUST exhaust.
- **DIAGNOSE an empty page via a DOM + network + downstream-service-log probe BEFORE assuming a fix surface
  (M42e iter-07 lesson).** "The content wasn't replayed/seeded" is a tempting but often-wrong guess. A page can
  render empty (a perpetual loading `<main>`) while ALL its content IS replayed — because a **federation error**
  on a non-nullable GraphQL field nulls the whole query client-side. iter-07: two skill-paths rendered empty
  not because they weren't replayed (they were — published, full chapter_list) but because their chapters'
  job-simulations referenced a taxonomy skill node-id (`K-AIFUNX-E658`) ABSENT from the demo — the federated
  `getSkillPath.chapters.@.jobSimulations.@.simulation.skills.name` is non-nullable, so the one missing skill
  nulled the entire `getSkillPath` payload → empty page. The root cause was a **stale public-taxonomy cache**
  (prod gained 22 public skills after the cache's capture date), fixed by a `stack-snapshot` taxonomy
  re-capture (in-rext, zero platform edit). The diagnostic technique: log in as the hero, navigate the empty
  page, capture (a) the `<main>` innerHTML (spinner vs error vs empty), (b) every GraphQL response's
  operationName + whether it carried `errors` or null `data`, (c) the relevant subgraph container's logs for
  the same window. The federation-error string names the exact entity + field + the missing id. Distinguish at
  triage: a missing **referenced public-reference row** (skill/role/content) → `stack-snapshot` re-capture or
  serve-grant; a missing **tenant row the page reads** → `stack-seeding`; a **runtime-computed surface** (a
  result/start deep-link) → crawl-scope.
- **A persona RE-specialization (role change) needs a `--reset` before re-seed on a LIVE stack (M42e iter-15
  lesson).** The verified/claimed `user_skills` rows are written with `CopyRowsIdempotent` keyed on a
  DETERMINISTIC slot id; when a hero's ROLE changes (e.g. Maya: Backend Software Engineer → DevOps Engineer),
  the slot id is unchanged but the underlying `skill_id` differs — so an ADDITIVE re-seed (`ON CONFLICT (id) DO
  NOTHING`) keeps the OLD skill at that slot and the new role's skills never land. A fresh `/demo-up` runs a
  `--reset` implicitly (truncate-then-reload), so this NEVER affects reproducibility — it only bites a
  re-specialization MEASUREMENT on an already-seeded live stack. Re-apply step for a role change: `stackseed
  --reset --force` THEN the full re-seed. (A re-specialization that also touches the casbin g2/g3 grants then
  also needs the Sentinel reload — the entitlement re-apply note above.)
- **The demo must be live + at the consumed tag.** The sweep runs against `demo-N` on offset ports; the demo
  consumes `rosetta-extensions @ <tag>`. A harness/fix change is **authored** in the authoring copy; to be
  reflected in the sweep it must be **applied** to the live demo (re-seed/re-replay/re-build/re-export — or, for
  a harness-only change, run the authoring-copy harness against the live demo directly).
- **Scope = the vantage's reachable set.** The crawl is bounded by what the hero's seat can navigate to; pages
  no link reaches are out of scope (the gate is over *reachable* pages).
- **Never wait on `networkidle` — use `domcontentloaded` + a bounded settle (M42e iter-03 lesson).** next-web
  holds long-lived connections (websocket / long-poll / streaming), so a page's network **never goes idle**.
  A crawl that navigates with `waitUntil: 'networkidle'` eats the full per-page timeout on **every** page,
  exhausts the test budget, and false-scores perfectly-good `http=200` pages as empty/error (the M42e baseline
  reported 44 "failures" that were all this flake — the true count was 8). The harness navigates with
  `waitUntil: 'domcontentloaded'`, then races a **short bounded** `networkidle` settle (`.catch(()=>{})` on a
  ~1.5 s ceiling) for hydration + first data paint, and never blocks on never-idle. Screenshots are captured
  **inline** (an `onPage` hook in the crawl, while the page is already loaded) — never a 2nd full re-navigation
  pass, which would re-introduce the timeout and double the nav count. The crawl also guards every assert
  against a **closed page** (the test-timeout race can close the page mid-assert): `page.isClosed()` is checked
  before each assert and the read is wrapped so a close degrades to an "unmeasured" note + a clean break, never
  a crash that loses the whole partial report. Give the sweep a generous test timeout (the full reachable set
  can hit the page cap; ~20 min) so the gate measurement is never truncated.
- **Seed paths are guesses, not coverage commitments.** The seed list primes the BFS frontier; a seed that
  404s or redirects **away** (e.g. the authenticated root → the real landing, or a stale `/skills` → 404) is
  **dropped, not scored** — only pages reached via real in-app nav links are coverage commitments. The redirect
  target is enqueued and scored as the real page. (Otherwise a wrong seed guess false-inflates `failing`.)
- **Read `innerText`, never `textContent` — and the error-sentinel reads the *visible* text only (M42e
  iter-03/06 lesson).** `textContent` serializes hidden + inlined content, including next-web's inlined
  Next.js/i18n JSON payload on routes that ship the whole translation table in the SSR'd `<body>` (e.g. the
  root `/`). That table carries translation **values** like `"Something went wrong while…"` which the
  error-sentinel scan then false-matched — flagging a fully-rendered 189 KB page as an error. The assertion
  reads `innerText()` (only what the user actually sees), which structurally excludes the inlined table.
  Prefer `<main>`; when a route has no `<main>` (the root `/` is a pre-redirect/loading shell), it is **not a
  real content page** — drop it from the seed list (the real landing, `/home`, renders content and is reached
  via seed or nav), rather than scoring the empty shell. (`innerText` returns visible text, which can be
  shorter than `textContent`; re-sweep after switching to confirm no legitimately-terse page drops below the
  density floor → a Tier-2 per-section assertion for any that do.)
- **Seedable structural row vs runtime-computed artifact — not every empty page is a seed gap (M42e iter-04/05
  lesson).** Some surfaces are filled by a **runtime computation**, not a seedable row: e.g. a sim **result**
  page (`/sim/<slug>/result/<sessionId>`) reads `jobSimulationResult.evaluationStatus` — an AI evaluation the
  jobsimulation pipeline computes server-side from the session transcript, never written by a seed. A backdated
  **seeded** session lists in `/profile/activities` (its `jobsimulation.sessions` row exists) but its result
  `<main>` renders empty (there is no computed evaluation). Under the **zero-platform-edit line** a seeded demo
  cannot populate such a surface (the only path is to *run* the evaluation — a platform action). The correct
  resolution is **crawl-scope**, not seeding: per-session computed-result deep-links are excluded from the
  vantage's reachable set (a `skipPaths` rule, e.g. `/\/result\/[0-9a-f-]{8,}/`), because the gate is over the
  pages a vantage can **meaningfully** reach in a seeded demo — a presenter lands on `/profile/activities`
  (which renders fine), not on a specific historical session's runtime evaluation. Distinguish the two at
  triage: an empty page whose data is a structural row → `stack-seeding` (fix it); an empty page whose data is
  a runtime computation → crawl-scope (exclude it), or escalate as a re-scope-trigger **only if** the link is a
  load-bearing part of the vantage's demo and a platform change is the sole filler.

  > **⚠ CORRECTION (v2.5 "the playbill" M236 — the premise above is REFUTED for sim results).** The rule
  > survives; **its stated reason does not.** A `/sim/<slug>/result/<sessionId>` page is **NOT** a
  > runtime-computed artifact — **M231 established it is a PERSISTED READ**
  > (`jobsimulation/internal/graph/queries.resolvers.go:70` does plain Ent SELECTs of
  > `validation_attempt_results`; no engine/LLM recompute on render). **M232** writes that fan-out from
  > cloned real sessions, and **M236 proved it live on `billion`**: **13/13** seeded player results render a
  > real score, a real LLM feedback paragraph and the real evaluated skills.
  >
  > So *"a seeded demo cannot populate such a surface"* was wrong, and any future triage that cites this
  > bullet to skip a result page is citing a refuted premise. **What actually justifies the `skipPaths` rule
  > is CRAWL SCOPE alone:** a hero's BFS should not dive into arbitrary historical session deep-links. The
  > rule is therefore KEPT in `tests/coverage.spec.ts` (`RESULT_DEEP_LINK_SKIP`) with its comment corrected —
  > **not deleted**, because deleting it proves nothing extra and merely makes the hero sweep re-walk pages
  > that are already asserted far more precisely elsewhere.
  >
  > **Where result pages ARE proven: `tests/content-stories.spec.ts`** (the M236 content-stories sweep,
  > §"Content stories — the (session × action) LANDS sweep" below). It reaches them the only way they can be
  > reached — the exact session URLs from the seeded `content-manifest.json`, each entered through its own
  > owning cockpit seat — which is precisely why the crawl cannot and should not be the tool for the job.
  >
  > **The generalizable lesson:** this bullet stood unchallenged for four releases because it was written as
  > a *fact about the platform* ("an AI evaluation … never written by a seed") when it was really an
  > *inference from one observation* (a seeded session's result page was empty — because nothing had seeded
  > the result rows yet). When you record a skip, record **what you observed**, separately from **why you
  > think it happens** — the observation stays true, the explanation is what rots.
- **An entitlement-gated empty page is SEEDABLE, not runtime-computed — seed the entitlement, don't skip the
  page (M42e iter-09 lesson; corrects the iter-08 sim-start mis-triage).** Not every empty `/start`-style
  launch surface is runtime-computed. The per-sim `/sim/<slug>/start` page renders the **org-member deny
  modal** (an empty `<main>`) when the member's org lacks the `FEATURE_JOB_SIMULATIONS` entitlement. The page's
  `canStartAsOrganizationMember` reads `userMembership.organizationFeatures`, which `app` resolves via
  **Sentinel's Casbin grouping policy `g3`** (`g3 = _, _` → a `casbin_rules` row `p_type='g3', v0=org,
  v1=membership`), NOT the `app.organization_features` table (which is 0-rows even in normal operation — a
  red-herring symptom). The fix is a **`stack-seeding`** g3 feature grant per membership (mirroring the
  per-member g2 grant) — a demo employee SHOULD be able to start a sim — NOT a crawl-scope skip. Distinguish at
  triage: an empty page gated by a **missing entitlement/policy row** → `stack-seeding` (seed it); an empty page
  filled only by a **runtime server computation** (a sim/skill-path RESULT keyed by sessionId) → crawl-scope.
  Re-instating a skip on a seedable failure is a dishonest scope-out (the gate-honesty failure mode).
- **A casbin/policy seed applied to a LIVE stack needs a Sentinel policy RELOAD before it takes effect (M42e
  iter-09 lesson).** Sentinel's Casbin enforcer calls `LoadPolicy()` **once at startup** with **no watcher** —
  a raw INSERT into `casbin_rules` (the seeder's path) is invisible to the running in-memory enforcer until it
  reloads. On a fresh `/demo-up` the seed precedes Sentinel start, so this never bites; it bites only a
  **re-seed of a running stack** (a Phase C re-apply of a casbin-touching seeder). Re-apply step: **restart the
  demo's `<demo>-sentinel-1` container** (re-runs `LoadPolicy()` on startup) — or call the `Reload` RPC. The
  app's 1-min in-process feature cache also expires on its own. Demo-local container op, zero platform edit.
- **Some pages render their real content OUTSIDE `<main>` — fall back to `<body>` innerText when `<main>` is
  below the floor (M42e iter-09 lesson).** The sim `/start` launch UI (`AISimulationStartWithoutSession`) mounts
  a sibling region with an EMPTY `<main>` while the visible launch content (~625 chars) lives in `<body>`. A
  strict `<main>`-only density read FALSE-FAILS it. The harness prefers `<main>`, but when the `<main>` read is
  below the density floor it re-measures against the live `<body>` innerText (still VISIBLE-only, so the
  iter-03 inlined-i18n exclusion holds) and takes the larger. This is the Tier-2 escalation for out-of-`<main>`
  content; the `<main>`-preference (nav-chrome exclusion) still governs the common case.
- **Bound the per-page settle to the heaviest DATA GRID, not the first paint — under-settle COLLAPSES the BFS
  frontier (M42e iter-09 lesson).** The crawl extracts outbound links from what is PAINTED at settle-time, so a
  settle too short for a heavy catalog grid (the library's 22 skill-paths / 307 sims) extracts too few links and
  the whole BFS frontier collapses (iter-09: a 1.5s ceiling rendered only 1 of 22 skill-path cards after a cold
  start → the frontier fell from 93 pages to 8). `networkidle` returns as soon as the network quiets, so a
  generous ceiling (4s) costs fast pages nothing (they proceed early) and only the busy grids use it; a
  never-idle long-polling page just hits the ceiling and proceeds (the `.catch`). Set the ceiling correct-over-
  fast: full link discovery is a gate-correctness precondition. WARM the stack (or re-sweep) after a cold start
  (e.g. a sentinel restart that cleared GraphQL caches) before quoting the authoritative residual.
- **At ORG SCALE a single bounded re-assert is NOT enough for the heaviest data grids — use a BOUNDED
  re-assert POLL + warm the grid (v1.10 M46 iter-07 lesson).** The per-section slow-paint guard was a
  *single* bounded re-assert (one extra settle, re-measure once). It was calibrated against a ~221-member org;
  when the **generated supporting population** filled the org to ~1k members, the manager's enterprise data
  grids (`/enterprise/members`, `/enterprise/activity-dashboard`) grew ~4.5× and their server-query +
  serialize + client-render exceeded one re-assert's budget — so the harness captured a **skeleton** frame and
  false-failed three sections whose data was fully present (`200`, full chrome + org name/logo + table headers,
  but skeleton rows + a `0 / ∞` count). The screenshot is the tell: real chrome, skeleton rows. The fix is
  twofold and is the org-scale extension of the iter-09 settle lesson: (1) the slow-paint guard re-asserts in a
  **bounded poll** (re-settle + re-measure up to N times, return on the FIRST pass, keep polling ONLY while the
  verdict is a paint-timing kind — `skeleton`/`empty`; a genuinely-empty section still fails after the budget,
  so there is **no false-pass**); (2) the **warm set is vantage-aware** — the manager vantage warms its heavy
  enterprise grids (`/enterprise/members`, `/enterprise/activity-dashboard`, `/enterprise/workforce`) before
  the authoritative sweep, so the first visit isn't paying the cold-cache + heavy-grid cost together. A
  skeleton-frame false-fail on a populated org is a measurement bug, NOT a content gap — diagnose it by the
  screenshot (real chrome + skeleton rows) + the DB (the rows ARE there), and fix the harness, never the seed.
  - **A poll that EXHAUSTS on a never-resolving query is a PERF wall — diagnose N+1-vs-combination before
    re-scoping (M46).** The bounded poll above closes a *slow-paint* skeleton (the grid hydrates within a few
    re-asserts). It does **not** close a grid whose backing query takes far longer than the poll window. At org
    scale the manager's enterprise grids (`/enterprise/members`, `/enterprise/activity-dashboard`,
    `/enterprise/settings`) hit this: the **federated GraphQL** queries (`organizationMembers` +
    `membershipsCount` + the activity aggregation) fan out per-row resolvers
    (`jobRole`/`targetRole`/`tags`/`lastActivityDate` × a **Sentinel RPC** in `app` `roles.go
    GetOrganizationTargetRole`, no DataLoader), and the Cosmo router logged **10–84 s** latencies — the page sits
    on a `…` spinner / skeleton through the poll. Tell it apart from a slow-paint by **measuring the query, not
    the pixels**: (a) `docker logs <stack>-graphql-1 | grep '"latency"'` shows 10 s+ requests; (b) the raw SQL
    for the same shape is **milliseconds** — so the wall is the **RPC round-trip COUNT**, not the DB.
  - **Then decompose the wall — PART is demo-patchable, PART is a genuine re-scope trigger (M46).** iter-07
    called the whole wall a re-scope-trigger; M46 proved it's **two distinct costs** that must be split:
    **(A) demo-patchable** — (1) an **over-broad fetch**: `InsightsContext.tsx` (the activity-dashboard provider)
    loaded `useGetOrganizationMembers({ limit: 1000 })` = ALL ~500 members → **next-web pagination demo-patch**
    caps it to `limit: 30` (the `/enterprise/members` grid already paginates at 20, untouched; the CSV/email
    export uses a separate query, no data lost); (2) **2 missing FK indexes**
    (`membership_skills(membership_skill_membership)` + `membership_tags(membership_tag_membership)` were
    seq-scanned per row) → **post-seed `CREATE INDEX IF NOT EXISTS`** on the demo's own Postgres (a rext-owned
    `up-injected.sh` step, idempotent, non-fatal — NOT a canonical ent/atlas change). (A) took graphql max
    latency **84 s → ~4 s** and cleared `/enterprise/activity-dashboard` + `/enterprise/settings`.
    **(B) demo-patchable by DROPPING the read-gate (Option B, the M46 close).** `/enterprise/members`' per-row
    `targetRole` → `app` `roles.go` `RoleManager.checkPermission` → `OrgCheckActionPermission` Sentinel RPC
    checks `OrgActionAssignmentsWrite`, which is **PER-OBJECT** (per assignee) — a manager can-write-assignments
    for some members and not others. **Two failed/rejected approaches first, then the safe one:**
    **(T2, REVERTED — a correctness bug)** a cache/singleflight keyed by `(org, subject, action)` (dropping the
    object so it dedupes across rows) returns the first row's allow/deny for every row → `failed to get target
    role: forbidden` on legitimately-allowed members (~1744×/sweep) → the grid ERRORS, NOT just slows. Keyed
    correctly `(org, subject, OBJECT, action)` it can't dedupe (every row a different object). **Caching the
    per-object check is unsafe.**
    **(B, SHIPPED — the safe demo-patch) DROP the read-gate, don't cache it.** `checkPermission` short-circuits
    `return true, nil` BEFORE the per-member `OrgCheckActionPermission` RPC (the same shape as its built-in
    `privacy.DecisionFromContext` bypass). Target roles still come from the DB
    (`GetOrganizationTargetRoleByAssignee`, indexed) downstream, so every member's **REAL** role renders — fast
    (DB-only) AND fully-populated (no per-row "forbidden" sparse rows; the 741/min legit denials vanish). This is
    a **READ-path authz relaxation ONLY** (manifest `patches/app-targetrole-authz-skip`, applied to the demo's
    ephemeral build-scratch clone via a rext helper wired INTO `up-injected.sh`'s inject loop, svc=app, AFTER
    `apply-authn`, BEFORE the build, trap-reverted git-clean): the assignment **mutations** still enforce via
    their own direct `OrgCheckActionPermission` calls (`assignments.go`, `resolver_mutations.go`, and `roles.go`'s
    own `CreateOrganizationTargetRole`). On demo-3 (~500 members) B took the members query **76.7 s → 0.51 s**
    (~150×), 0 forbidden errors, and **cleared `/enterprise/members` → the manager gate is MET (failingSections=0,
    NO re-scope)**. **The PLATFORM finding stays documented:** prod still hits ~77 s at 500-member scale and
    genuinely needs a **DataLoader / batch `BulkCheckPermission` RPC** — B is a disclosed single-presenter
    demo-perf relaxation, NOT a prod fix. **Lesson: decompose a perf wall before judging it; and a per-OBJECT
    authz RPC can't be CACHED (object-blind = wrong answer) but it CAN be DROPPED where the read returns real
    DB data and the mutations stay enforced — dropping a read-gate is safe where caching-it-wrong wasn't.**
  - **A never-COMPLETING server response is a different wall than a slow-PAINT — check the backend request log
    for a completion BEFORE investing in a harness warm (v1.10b M51 iter-06).** The slow-paint poll + the
    `warmHeavyGrids` cache-primer both assume the backing query EVENTUALLY returns (the warm primes a result
    that WOULD complete; the poll waits for a paint that WILL happen). A grid whose server response **never
    completes in-budget** defeats both: `GET /api/workforce/ai-readiness` on the 200-member showcase org logged
    **ZERO completions** across the entire backend log (only OPTIONS preflights) — the React Query fetch is
    aborted every time, and the `ai_readiness_refresh` background worker (same compute) timed out
    `context deadline exceeded`. Tell this apart from a slow-paint / cold-cache wall by **grepping the backend
    request log for a COMPLETED request**: `docker logs <stack>-backend-1 | grep 'GET /api/<route>' | grep -v
    OPTIONS` — **0 hits = the server can't produce the response in-budget**, so no warm/poll will help (you'd be
    deepening a primer for a result that never arrives). The root cause here was NOT index-bound (every
    AI-readiness SQL query EXPLAINed at ms — jobsimulation.sessions fully indexed) and NOT the M46 members-grid
    fan-out: it was the **response-build live-recompute + a per-skill federated TRANSLATION fan-out**
    (`withSkillerLang` → backend `_entities` `get skill translation <uuid>/english`, one round-trip per skill in
    the aggregate's skill set — visible as a `context canceled` storm in `<stack>-backend-1` logs; the taxonomy
    translation surface is served by the backend subgraph since the skiller→app merge). That is the
    **same N+1 family as the M46 per-object Sentinel RPC**, in the translation path. **And a materialized
    snapshot mirror only helps if the read path CONSULTS it** — the default AI-readiness dashboard GET always
    takes the live-recompute branch (`buildLiveResponse`; the `ai_readiness_live_snapshots` read is gated behind
    a *closed* `CycleID`), so seeding the snapshot table would NOT short-circuit the default call — the
    short-circuit itself is a platform-read-path change. **Decompose like M46:** if a demo-patch can batch/relax
    the translation fan-out or make the default call read snapshots (the `app-targetrole-authz-skip` precedent),
    that's demo-local; if not, it's the milestone Re-scope trigger (`unimplementable-without-platform-edit`) —
    escalate, never edit the platform, and never widen the harness budget to mask a server that can't answer.
  - **A cycle-scoped FAST read-path only helps if the DEFAULT client call SELECTS it — confirm the FE's
    request SHAPE, not just the server branch (v1.10b M51 iter-07).** iter-06 root-caused the AI-readiness
    wall as a live-recompute + translation N+1 on the ACTIVE-cycle path; the M48 contract documents a fast
    alternative — a CLOSED cycle whose read takes `buildResponseFromSnapshots` (a pre-computed frozen read).
    iter-07 seeded the cycle closed + a frozen `ai_readiness_snapshots` row per member (DB-verified correct:
    199 snapshots, 78.4% stage-3, heroes right) — and the GATED sweep STILL held at the same failing count.
    The frozen branch EXISTS in code but the DEFAULT dashboard GET never takes it: `app
    GetAIReadinessWithOptions` reaches `buildResponseFromSnapshots` ONLY for `opts.CycleID != nil &&
    status=="closed"` — the nil-CycleID default is hardcoded to `buildLiveResponse`. An AUTHENTICATED network
    probe (log in as the hero, log every outbound backend request URL + query params + whether it completes)
    proved the demo FE fires the data GET **WITHOUT `?cycle=`** (the live path — it hangs) and never fires the
    `/cycles` list that would supply `latestClosedCycle.id`. **Lesson: a server-side fast branch is necessary
    but not sufficient — before assuming a data-shape change (a closed cycle, a materialized mirror) clears a
    wall, confirm BOTH sides: (1) the server branch exists AND (2) the FE's DEFAULT call fires the variant that
    hits it. Diagnose the FE side with an authenticated network probe (which request VARIANT the client fires),
    not just the backend completion log (which only tells you IF a request completed). When the fast branch is
    reachable only via a param the default FE omits, closing the gap is PLATFORM-bound (the FE must pass the
    param, or the backend default must prefer the closed cycle) → the milestone Re-scope trigger, or the
    disclosed-presenter-note (data proven-correct in the DB, slow-only via the default route) with the user's
    explicit sign-off.**
  - **A "frozen"/materialized read path can carry its OWN org-scale wall — measure the FROZEN read
    END-TO-END, not just confirm it is reachable (v1.10b M51 iter-08).** iter-07 assumed the closed-cycle
    frozen branch (`buildResponseFromSnapshots`) was *fast* and that the only gap was the FE not routing to it
    (the default GET omitting `?cycle=`). The user's chosen zero-edit fix was to **deep-link the demo entry** —
    make Dana's cockpit `jump_to` + the coverage manifest carry `?cycle=<latest-closed-cycle-id>` so the FE
    fires the cycle-scoped GET → the frozen branch. **Before touching the cockpit/manifest, an authenticated
    DUAL-ENDPOINT DIRECT probe** (lift the hero's bearer from a real outbound request, then hit `/cycles` AND
    the frozen data GET `?cycle=<closed>` **directly** via `page.request`, bypassing the FE's React Query
    orchestration) **falsified the premise**: `/cycles` returned **200 in 40 ms** (fast — the FE gate is fine),
    but the frozen data GET `?cycle=<closed>` **NEVER COMPLETED** (timed out the full 180 s budget), *identical*
    to the live-recompute default. Root cause: `buildResponseFromSnapshots` reads the frozen scores fast but
    then calls **`loadMembers(orgID, "")` — a full unbounded org-member hydration** (`hydrateMembers` with
    `memberIDs=nil, userIDs=nil` → whole-org tag/skill/sim aggregation) to attach current identity/tags to each
    snapshot; at 200 members that hydration is the SAME org-scale wall as the live path (and even the
    `ai_readiness_refresh` worker's parallel compute logs `context deadline exceeded`). So the frozen branch is
    NOT a fast path in this demo — the deep-link cannot clear the wall even in principle, and (crucially) it is
    NOT one of the demo-patchable costs: `queryBaseMembers` here reads `jobRole` from a SQL column (NOT the
    per-object targetRole Sentinel RPC that `app-targetrole-authz-skip` already drops), so the existing
    demo-patch does nothing for it. **Lesson: "frozen"/"pre-computed"/"materialized" names a *scores* freeze, not
    necessarily a *response* freeze — a snapshot read that re-joins live per-member identity re-incurs the
    org-scale member-load wall. Before betting a strategy on a fast branch, MEASURE THE BRANCH END-TO-END with a
    direct authenticated probe of the exact request the strategy will fire (here `?cycle=<closed>&includePeople=true`),
    not just confirm the branch is reachable / the DB rows are correct. When the frozen read is itself
    org-scale-slow and its cost is NOT a demo-patchable authz gate, the remaining zero-edit path is the
    disclosed-presenter-note (data proven-correct in the DB, slow-only) — which needs the user's EXPLICIT
    sign-off; a platform fix (bound `loadMembers` in the snapshot path / a frozen_tags column so it needn't
    re-join live members) is the Re-scope trigger. Reusable diagnostic:
    `stack-verify/e2e/tests/probe-aireadiness-deeplink.spec.ts`.**
    **Build pitfalls (each cost a full re-seed):** the injected Go images are built from a build-scratch clone
    AFTER `apply-authn.sh` vendors the **disarmed colony** (Clerkenstein token acceptance); a standalone `app`
    rebuild that SKIPS that step ships a backend that calls real `api.clerk.com`, rejects every demo token, and
    collapses the crawl to `reachable≈7` (looks like a content fail — it's broken auth: grep
    `docker logs <stack>-backend-1` for `clerk`). And never recreate a single service with `--force-recreate`
    *without* `--no-deps` — it recreates `postgresql` too and wipes the seeded org. Never widen the poll to mask a
    slow query, and never shrink the org below the org-scale premise just to pass.
  - **An unbounded-hydration perf wall often has an EXISTING id-restricted loader — swap, don't rewrite; and
    a demo-patch CAN bound a snapshot read-path where it can't cache an authz gate (v1.10b M51 iter-09, the
    fix that closed the wall).** iter-08 proved the frozen AI-readiness read (`buildResponseFromSnapshots`)
    times out because it calls `loadMembers(orgID, "")` — a full UNBOUNDED whole-org member hydration — to
    re-join current Tags/Name/Role onto each frozen snapshot. The fix is a **new app read-path demo-patch**
    (`patches/app-aireadiness-snapshot-loadmembers`, the `app-targetrole-authz-skip` precedent: a pinned
    anchor→replacement manifest + a rext-owned `apply-*.sh` helper mirroring the swap, wired INTO
    `up-injected.sh`'s inject loop svc=app after apply-authn + the authz-skip, before the build, trap-clean,
    `DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND=1` opt-out) that **bounds** the hydration: collect the ~199
    snapshot user-ids and call the EXISTING bounded sibling `loadMembersByUserIDs(orgID, "", snapUserIDs)`
    (indexed `memberships."user" = ANY(...)`) instead. It is a **PURE perf optimization, data-identical**:
    `buildResponseFromSnapshots` uses the members map ONLY keyed-by + looked-up-by each snapshot's UserID, so
    the returned `AIReadinessResponse` is byte-identical; only the member-load cost drops. The frozen
    `?cycle=<closed>` GET went **180 s-timeout → 19 ms** and the dashboard rendered the full funnel. **Two
    lessons:** (1) when a perf wall is an unbounded hydration, look for an id-restricted sibling loader in the
    codebase BEFORE writing a new query — the fix is often a one-call data-identical swap. (2) Unlike a
    per-OBJECT authz RPC (which can't be CACHED object-blind — the M46 T2 correctness bug — but CAN be
    DROPPED), an unbounded READ hydration can be **BOUNDED** by a demo-patch where the narrower set is already
    known (the snapshot user-ids), returning identical data faster — a third safe demo-patch shape alongside
    drop-the-read-gate (M46/B) and cap-the-fetch (M46/A). The prod finding stays: prod's frozen read still
    hydrates the whole org and needs `loadMembers` bounded in the snapshot path / a `frozen_tags` column
    (M314b) — a disclosed demo-perf relaxation, not a prod fix.
  - **A believability-gate section descriptor that requires TWO strings the FE renders MUTUALLY EXCLUSIVELY
    is a latent FALSE-FAIL — check EITHER/OR conditional headers before treating an absent substring as a
    content gap (v1.10b M51 iter-09).** The AI-readiness funnel section false-failed on "region missing
    required content: Stage breakdown" (re-asserted 6× — NOT a paint-timing skeleton) while its sibling
    section passed and the funnel WAS fully rendered. Root cause: `AIReadinessView.tsx` renders the funnel
    header as ONE of two mutually-exclusive strings — `t('stepsCompletionLink')` ("Steps completion", a link)
    when an `onStepsClick` handler is provided, ELSE `t('stageBreakdown')` ("Stage breakdown") — never both.
    The manager dashboard always provides the steps-completion drawer handler, so it renders "Steps
    completion" and NEVER "Stage breakdown"; the descriptor's `mustInclude: [..., 'Stage breakdown', 'Steps
    completion']` was impossible-in-manager-mode. The fix is a **harness** correction (drop the impossible
    alternative, keep the load-bearing proof — the three stage labels + the funnel header), NOT a seed/content
    fix and NOT a gate-loosening (the funnel's real proof still asserts). **Distinguish at triage:** a section
    that fails on ONE substring while its siblings render + the page's `main` dump shows a SIBLING string of an
    either/or pair → a mutually-exclusive-header descriptor bug (harness fix); a section that fails on ALL its
    substrings + a skeleton screenshot → a real empty/slow-paint (seed or warm/poll fix). The tell is the
    probe/`main` dump: if the missing string is the OTHER branch of a conditional you can SEE rendered, it's a
    descriptor false-fail.
- **An editorial citation in replayed content is VALID content, not a gate escape — disclose it, don't strip
  it (M42e iter-08 lesson).** Replayed `/skill-path/.../chapter` body copy can carry a real external `<a href>`
  citation (e.g. an `en.wikipedia.org` / `strategy-business.com` reference inside the course material). That is
  **content fidelity** — the gate's whole point is real content — so it must NOT be rewritten to an offset port
  and must NOT fail the escape check. The harness carries a narrow **allow-rule** (`allowedExternalLink`,
  anchored to `/chapter` paths): such a link is recorded as an `allowedCitation` → surfaced in a
  **presenter-notes** list (so the presenter doesn't click it live), but is **not** counted in the escape total.
  The rule is deliberately narrow so it CANNOT mask a real escape: a nav-chrome / baked-URL escape (a left-menu
  "Studio" → prod `studio.anthropos.work`) is not on a `/chapter` page and **still fails** the gate. Disclose,
  don't hide. Distinguish at triage: an off-demo link in **replayed editorial content** on a content page →
  allow-rule (presenter note); an off-demo link in **nav chrome / a baked app URL** → escape (rewrite it).
- **A long single-test sweep MUST stream per-page progress to stdout — fold the heartbeat into the harness, not
  the runner (M42e iter-08 lesson).** The coverage sweep is one long Playwright test (~13 min over the full
  reachable set); with the `list`/`line` reporter it otherwise emits NO stdout until the whole test finishes,
  which trips a >5-min output-watchdog (a run-2 attempt was killed mid-sweep by exactly this). The crawl loop
  emits one `console.log` per scored page (`[crawl] N/cap q=Q VERDICT status path`) so the sweep streams +
  stays observable. Pure observability — it does not touch the crawl frontier, the scoring, or the report JSON.
  When driving the sweep, line-buffer (`stdbuf -oL`) + `tee` so each line reaches stdout, and append a journal
  heartbeat every few minutes; never run the sweep as a silent foreground call.

---

# Sweep 2 — Content stories: the (session × action) LANDS sweep (v2.5 "the playbill" M236)

**The second sweep this protocol governs.** The hero sweep above proves *a vantage's pages are full*. This
one proves *every seeded content story actually lands on a real result* — the Thread-B half of v2.5.

### Why it is NOT a `VantageManifest` crawl

A content-story result page is reached by **cockpit seat + an exact URL**, never by crawling:

- the URL carries a **session uuid that only the seeded manifest knows** (`/sim/<slug>/result/<sessionId>`);
- the seat that owns it is a **non-hero MEMBER** (`content-player-<idx>`), not a roster hero;
- `VantageManifest.identityKey` is **singular**, so a manifest-per-seat model would mean **13 seats → 13
  manifests** for what is really one uniform assertion.

`cockpit-login.loginAs()` already accepts a **`landingPath`**, so the whole sweep is an **exact-path visit
per (seat × path)**. The BFS crawl machinery is simply the wrong tool and is deliberately unused. *(This
also shrinks the harness M235's carry-forward described: the work is a page-object + a runner, not 13
manifests.)*

### The pieces

| Piece | Path (in `rosetta-extensions`) | Role |
|---|---|---|
| Page object + assertions | `stack-verify/e2e/lib/content-result-page.ts` | Models the result surfaces; **authored from scratch** (see below) |
| The sweep | `stack-verify/e2e/tests/content-stories.spec.ts` | Enumerates landable pairs from the manifest, visits each |
| Runner | `stack-verify/e2e/run-content-stories.sh` | Resolves offset ports, fetches the **served** manifest, runs + aggregates |
| Aggregator | `stack-verify/e2e/aggregate-content.py` | Turns the per-pair ledger into the reading |

> **The `AISimulationResultContainer` trap.** M235's carry-forward said the harness would "reuse the shared
> `AISimulationResultContainer`". **It does not exist as a harness object** — it is a next-web `.tsx`
> component. The nearest harness object (`playthroughs/e2e/lib/simulation-page.ts`) stops at the *launch*
> boundary with no `/result/` locator. The result surface had **never** been modelled harness-side; M236
> authored it, **calibrated against a live seeded render** rather than blind.

### The SIX render shapes (calibrate; do not assume)

The roadmap called this "the result page." It is **six distinct surfaces across two apps**, and a single
"is it long enough?" check is wrong on all of them — it both false-fails and false-passes:

| Shape | Surface | What it looks like live | Assert on |
|---|---|---|---|
| `player-scored` | `/sim/<slug>/result/<id>` | ~1.9k chars: score, LLM feedback paragraph, "Evaluated Skills" | feedback **or** evaluated-skills present, `<main>` ≥ 300 |
| `player-interview` | same route, interview sims | **~205 chars by design** — an acknowledgement only; the player is *not* shown a report | the acknowledgement text |
| `player-skillpath` | `/skill-path/<id>` | 2.9k–11k chars: chapters + a progress signal ("Continue (45%)" / "100% complete") | chapter/path structure **and** a progress indicator |
| `player-academy` | `/courses/<slug>` (**ant-academy**, a different app) | ~3.7k chars: `COURSE · 12 CHAPTERS`, title, chapter list | course/chapter structure **and 0 Draft chips** |
| `manager-dashboard` | `/enterprise/activity-dashboard/ai-simulations/<simId>/<membershipId>` (`skill-paths` was REMOVED at iter-07 — that surface is unimplemented) | header (`<player>'s Results for <sim>`, "N skills measured") **plus an attempts TABLE** | the **table rows** — the header alone is chrome |
| `manager-interview` | `/enterprise/activity-dashboard/interviews/<simId>/<membershipId>` | ~590 chars: breadcrumb naming the player over an attempts table with "View Report" — **no `Results for` header at all** | an attempt row |

**Shape selection is by ROUTE, never by sniffing content.** Every shape after the first was discovered by
driving a route and *reading* it; every mis-grade came from inferring shape from keywords. In a system whose
surfaces share a design language, the URL is the only cheap, stable discriminator.

> **The calibration lessons, in the order they were paid for (M236 iters 04–08):**
> 1. **A terse page can be correct.** The interview player result is ~205 chars *and right*. A length floor
>    tuned to the scored shape reports a correct page as empty.
> 2. **A false PASS is worse than a false fail.** Grading skill-path pages with the scored-sim shape passed
>    the *completed* path — not because it rendered a report, but because 11k chars of legitimate content
>    happened to contain the word "feedback".
> 3. **A different surface is not a broken surface.** The interview *manager* view has no `Results for`
>    header, so `manager-dashboard` graded a perfectly-rendering page as broken (iter-06).
> 4. **A header can render from a different query than the payload.** The manager scoreboard's header comes
>    from the sim definition, so it populates even when the payload query has nulled entirely — which is how
>    a membership-id defect hid for an iter (iter-05), and how an *unimplemented* surface passed for two
>    (iter-07).
> 5. **A different app is a different shape.** The academy course page has no score and no feedback;
>    falling through to `player-scored` reported a correct 3.7k-char render as a failure (iter-08).
>
> **Running score across the milestone: five wrong assertions to one real product bug.** When a gate is new,
> disbelieve the gate first — probe the page before triaging the product.

### Two traps that make a working page look broken

- **Never let `loginAs` take its `networkidle` default on these routes.** next-web holds long-poll
  connections open, so `networkidle` resolves late or never; on the enterprise activity-dashboard it never
  resolves at all. The pair recorded `page.goto: Test timeout of 180000ms exceeded` for two iters and was
  read as a product **hang** — a heavy instance "hanging" while a lighter sibling passed, which is exactly
  the per-item fan-out signature `latency-budget.md` teaches. It was neither: instrumenting the navigation
  showed **134 completed legs, 0 pending, none over 800 ms, page painted in ~1 s**. Navigate with `commit`
  and let the page object's `settle()` own the wait. *(The rule was already written down in two places and
  still got inherited by default — hence `waitUntil` is now explicit at the call site.)*
  **Diagnostic:** `stack-verify/e2e/tests/probe-navigation.spec.ts` prints per-request wall times and
  everything still in flight at a deadline. Run it before ever blaming a route for hanging.
- **Read `body`, not just `<main>` — portals mount outside it.** The activity-dashboard drawers are antd
  `Drawer`s rendered through a **portal**, so `main.innerText` is the **empty string** on a fully-populated
  page. `main.length === 0` on a rendered page is a modal/drawer signature, not an empty one.

### The denominator (get this wrong and the gate lies)

**`has_manager_view` is per-SESSION, not per-product.** Reading it at product level silently under-counts
**31 → 18**. And **ai-labs is presence-only** (nil client, `grade_result` not GraphQL-exposed — M231): its
sessions have **no landable result surface**, so its player actions are **excluded from the denominator**
rather than counted as failures.

**A pair is landable only if the platform has BUILT the surface it points at.** M236 iter-07 drove the
skill-path manager route live and found it renders the literal string **"Coming soon"**: in
`InsightsBySkillPathStudentSimulationsContainer.tsx`, `userData` is hardcoded `null` and the results table
is **commented out**, so **no query touches the seeded session** and the page is identical whether or not
anything was seeded. Its 2 manager pairs are therefore **not landable** — the same ground on which
ai-labs is already excluded — and `skill-path-legacy` is projected **player-link-only**. On the corrected
manifest:

```
 18 sessions + 15 manager views = 33 raw pairs        (the manifest BEFORE iter-07)
                     − 2 skill-path manager views     (surface unimplemented — iter-07)
 18 sessions + 13 manager views = 31 raw pairs        (the CORRECTED manifest)
                     − 2 ai-labs presence-only player actions
                              = 29 LANDABLE
```

**Show the whole chain, because `31` names two different quantities.** It is both the *pre-iter-07 landable*
count and the *post-iter-07 raw* total — a coincidence, not a correspondence. A reader who subtracts the 2
skill-path manager pairs from the landable 31 gets 27 and concludes the doc is wrong; the 2 that leave the
landable count are the **ai-labs** pair, and the 2 skill-path pairs leave one line earlier, out of the *raw*
total. State which of **33 / 31 / 29 / 18** a number is, every time.

> **Correcting a denominator is not moving the goalposts — but it must be argued, never quietly applied.**
> 31 was never a count of *provable* pairs; it assumed a surface that does not exist. The controlling rule
> is M233's fail-closed projection contract: *a session that cannot form a real link is dropped with a
> reason, never linked anyway.* A CTA onto "Coming soon" is a fabricated CTA. Record the product-source
> evidence in the iter's `decisions.md` and surface the change in the close — the danger is not the
> correction, it is a correction nobody can audit later.
>
> **This one also hid a false PASS.** The lighter of the two skill-path sessions had scored as passing for
> two iters, because the definition-only header contains "Results for" and a "Coming soon" body contains
> no "No data". The false fail cost an investigation; the false pass would have hidden a missing surface
> indefinitely.

### A green unit test can defend a broken path

Three separate tests in three consecutive iters were found asserting the defect they should have caught:

| Iter | The test required… | Reality |
|---|---|---|
| 05 | the manager route built from a **user** id | the route takes a **membership** id — the query nulled |
| 07 | `has_manager_view` **true** for skill-path | the manager surface is unimplemented |
| 08 | the academy CTA to start `/library/` | **there is no `/library/[slug]` route** — it 404'd every time |

A test that encodes a **route** or a **contract** is only as good as the last time someone drove it. When a
live probe contradicts a green test, the test is the prime suspect — and the fix is to invert it so it
fails loudly if the defect is ever restored.

**This is a defect class with a name — *offline-authored, never driven* — and a fourth instance (the skill-path
version `"2"` guess) sits outside this table.** The standing form, **unit-proven ≠ route-proven**, is
[cross-cutting rule 3](#3-offline-authored-never-driven--unit-proven--route-proven).

### Running it

```bash
cd rosetta-extensions/stack-verify/e2e
./run-content-stories.sh 1                                   # a local demo-1
./run-content-stories.sh 1 --host billion.taildc510.ts.net   # a REMOTE stack over the tailnet
```

The runner fetches the manifest **from the cockpit the stack is actually serving**, not from the checked-in
preset — a sweep that read the preset would pass while the live cockpit served something else. Output lands
in `content-out/demo-<N>/` (`content-stories.json`, `pairs.jsonl`, `content-manifest.used.json`).

### Two harness invariants that are easy to get wrong

- **Not `mode: 'serial'`.** Serial makes tests *dependent*: the first failure **skips all the rest**, turning
  "how many of 29 land?" into "did pair #2 fail?". The pairs must be independent; the single-worker
  constraint the shared fake-FAPI seat needs comes from `--workers=1`, which serializes **without** coupling
  outcomes.
- **Persist per-pair results to disk, don't accumulate in memory.** Playwright **restarts its worker after
  every failure**, re-importing the spec module — so a module-level array resets and an `afterAll` summary
  fires once per restart seeing only a fragment. The first version of this sweep printed `LANDED 1 / 31`
  once per failure for exactly that reason. The spec appends to `pairs.jsonl`; the **runner** aggregates.

### The reading must be fail-CLOSED (M236 final harden)

The five-wrong-assertions-to-one-real-bug score above is a statement about the **grader**. The same defect
class lives one layer up, in the thing that turns per-pair verdicts into **the number the gate is graded
on** — and there it went unnoticed through all ten iters:

- **`0 / 0` is not a pass.** `aggregate-content.py` computed its denominator as "rows in the ledger" and
  **always exited 0**. A run in which nothing executed — playwright collecting no tests, a manifest fetch
  that produced an empty sweep — printed `LANDED 0 / 0` and reported **success**. Arithmetically that is
  also 100%. An empty ledger is now a **failed run**.
- **A dropped pair must stay in the denominator.** A pair the manifest cannot form writes no ledger line, so
  a denominator counted from survivors **silently shrinks and flatters the score**: every remaining pair
  landing reads as a clean sweep. The spec's own comment promised this could not happen — true of its
  console output, never of the machine-readable reading. Drops now go to `dropped.jsonl` and **fail the run**.
- **Pin the denominator from outside.** A sweep that runs 26 of 29 pairs and lands all 26 is **not** a pass,
  and no self-consistent ledger can detect that. `EXPECTED_PAIRS=29` makes the count itself an assertion.
- **The runner's exit code is the verdict.** `run-content-stories.sh` ended on the aggregator with its
  result swallowed, so the script exited 0 whether 29/29 or 0/29 landed. It now `exec`s the aggregator.
  *(`run-coverage.sh` already carries the same lesson in a comment — "swallowing it with `|| true` is what
  let a failed sweep exit 0 for four releases" — and the newer runner reintroduced it anyway. A lesson
  written down in one runner does not propagate to its siblings by itself.)*
- **A truncated ledger line must not cost the whole reading.** A run killed mid-append leaves a partial last
  line; bare `json.loads` threw before any report was written, so the operator got a traceback instead of
  "*29 rows read, 1 truncated*". Bad lines are now counted and reported — still fail-closed, but legibly.
- **Name the stack, or measure the wrong one.** Every runner computes `OFFSET=$(( N * 10000 ))`, and bash
  evaluates a non-numeric `N` to **0 with no error** — pointing the whole sweep at the **dev stack's** ports
  and reporting what it finds there as demo-N's. Both content-stories and latency runners now reject a
  non-integer `N`. `run-coverage.sh` / `run-hiring-render.sh` shared the hazard and were **guarded at the
  M236 close** — the lesson is now propagated to all four runners, not written in one and left in the others.

> **The generalisation worth carrying to the next gate:** the milestone's signature failure was *a check
> scoring green off a subject that proved nothing*. That is not a property of render shapes — it is a
> property of **every layer that reports a number**. Ask of each one: *what does it print when nothing
> happened?* If the answer is anything other than a loud failure, the gate can certify a vacuum.
>
> **Promoted to a standing rule** — see [Cross-cutting rule 1](#1-a-check-can-report-success-while-proving-nothing),
> which carries all nine instances and the pair-every-shape discipline that catches them.

*(The last bullet is [cross-cutting rule 2](#2-prose-does-not-propagate--only-a-shared-definition-or-an-executable-fence-does)
in miniature: a lesson written in one runner does not reach its siblings. Four runners, one fence.)*

### The directional rule for tuning a grader

> **A fix for a false-FAIL that creates a false-PASS is a net loss.**

**The two errors are not symmetric, and a grader author must hold that asymmetry explicitly.** A false-FAIL is
loud, lands in front of a human, and costs an investigation. A false-PASS is silent, certifies a vacuum, and
is discovered — if ever — releases later by something else entirely. **Trading one for the other is not
neutral even at one-for-one**, and a grader is *always* tuned under pressure from the false-FAIL side, because
that is the side that is currently shouting at you.

**The live instance.** `settle()` counts `main + body` so it can see antd-`Drawer` surfaces mounted through a
portal, while the length floors read `main.length` — so a portal-rendered page settles correctly and then
fails *"too short (0 chars)"*. A real false-FAIL. The obvious repair — `readable = main.length + body.length`
— **double-counts**, because `<main>` is a descendant of `<body>`. That silently **halves every floor**, and a
blank page carrying only nav chrome clears a 300-char gate: the exact false-PASS class six iters had been
spent eliminating. The landed fix is `main || body`.

**The check to run before any grader loosening:** *if this change is wrong, which way does it fail?* If the
answer is "it passes something it shouldn't", the change needs a **negative test proving the specifically-broken
page still grades not-ok** before it lands — not after. Loosening a floor is a false-PASS risk in every case;
loosening it to fix a false-FAIL is exactly when that risk is easiest to overlook.

### Prove the test fails (mutation, not coverage)

The harden pass added ~70 pure-logic tests over the grader, the denominator, and the reading — driven
through a scripted fake `Page`, no browser, no stack. The discipline that makes them worth having is not
line coverage; it is that **every shape gets a PAIR**: a good page graded `ok`, and a specifically-broken
page graded **not** `ok`. A happy-path-only test cannot tell a working grader from one that returns `true`
unconditionally — which is exactly what this milestone kept finding.

Where a fix is load-bearing, the harden pass **reintroduced the bug and confirmed the tests went red**,
then restored it: dropping `TZ=UTC` from the green gate turns 5 of 6 guards red; renaming `/courses/` in
`shapeFor` turns 2 of 8 route-contract tests red. **A regression test nobody has ever seen fail is a
hypothesis, not a guard.**

**The cross-language contract nobody was testing.** The Go projection (`content_manifest.go`,
`content_nonsim.go` — iters 05/07/08) and the TS grader (`shapeFor` — iters 04/06/07/08) agree on these
routes by **bare string prefix**. Four iters touched each side; no test touched the join. And `shapeFor`
**falls through to `player-scored`** for any prefix it does not recognise — so renaming `/courses/` on the
Go side throws nothing, fails no Go test and no TS test, and just grades every academy page against the
wrong shape. That *is* the iter-08 defect, and after iter-08 nothing prevented its return.
`stack-verify/e2e/tests/content-route-contract.unit.spec.ts` reads the **checked-in canonical manifest** and
asserts the grader understands every route in it — including that the landable count is still **29**.

---

# Cross-cutting protocol — rules both sweeps owe to hard experience

_These four outlived the milestone that found them. They are stated here as standing rules because each was
first learned as a one-off, written down in one place, and then **recurred** — which is itself the subject of
the second rule._

## 1. A check can report success while proving nothing

**The single most transferable finding of v2.5.** Nine distinct instances landed across M235–M236, and they
share no code:

| the layer | what it reported | what it had proved |
|---|---|---|
| a PII scrub | every test green | it had removed **zero** names |
| three route tests | green, three consecutive iters | they asserted the **defect** they should catch |
| the aggregator | `LANDED 0 / 0`, **exit 0** | nothing executed; 0/0 is also 100% |
| a test suite | passing | it collected **0 tests** — 61 offline for 8 iters |
| a grader | green on every page | it had **no negative test**; `return true` would score identically |
| a regression test | green | it was a **self-consistent tautology** (`len(x) != len(x)`) |

> **The rule, in the form worth memorizing:** **ask of every layer that reports a number — *what does it print
> when nothing happened?*** If the answer is anything other than a loud failure, that layer can certify a
> vacuum, and everything downstream of it inherits the certification.

**This is not a property of render shapes or of Playwright.** It is a property of **any layer that reports a
number**: a grader, an aggregator, a runner's exit code, a coverage floor, a manifest projection, a scrub. Each
needs the question asked of it separately — closing it in one layer says nothing about its siblings.

**The practical discipline that catches it: every shape gets a PAIR.** A good subject graded ok, and a
specifically-broken subject graded **not** ok. A happy-path-only test cannot distinguish a working check from
one that passes unconditionally — which is precisely what kept being found. Where a check is load-bearing, go
further and **reintroduce the bug to confirm the test goes red** (dropping `TZ=UTC` turns 5 of 6 green-gate
guards red; renaming `/courses/` turns 2 of 8 route-contract tests red). **A regression test nobody has ever
seen fail is a hypothesis, not a guard.**

## 2. Prose does not propagate — only a shared definition or an executable fence does

**The best available predictor of where the next defect lands.** A lesson written into one file does not reach
its siblings, however well written, however prominent:

- the **non-integer-`N` guard** was added to 2 of 4 runners during the milestone; the other two carried the
  identical hazard until the close;
- the **`networkidle` rule** was *already written down* in [`latency-budget.md`](latency-budget.md) — and the
  new sweep inherited the bad default anyway, producing a 180 s "hang" on a page that painted in ~1 s;
- **doc corrections** followed the same shape: each iter fixed the doc it touched, leaving three siblings
  asserting the claim just refuted — including the one that produced the inflated denominator;
- the **membership key** was a bare literal at **9 sites** — one writes the row, eight merely hope to match it.

> **The rule:** when a correction has more than one site, **do not fix the sites** — replace them with a single
> definition, or add a fence that fails when they disagree. A sweep that fixes N places leaves place N+1
> unfound, and the next reader cannot tell which places were swept.

**The class recurs during the very work that names it, so assume it is live.** v2.5's own close is the example:
the `/demo-up` knob count was corrected to **27 env knobs + 10 CLI flags** in `demo-up-defaults.md` and
`CLAUDE.md`, while [`README.md`](README.md) still said 25 + 9 and the `demo-up` skill still said 26. A second
instance in the same close: a correction sweep that covered `corpus/ops/**` and **no file under
`corpus/services/`**.

**An executable fence already exists for that particular claim** — `rext stack-core/demo_knob_guard.py` checks
the defaults table against the parsers **in both directions** (a doc-promised flag with no parser entry is a
*false promise*; a parser flag with no doc row is *undiscoverable*). Its scope is `demo-up-defaults.md` only.
**Every other doc restating the count is unfenced prose, and drifted for exactly that reason.** Prefer citing
`demo-up-defaults.md` over restating the number; if you must restate it, get it from the guard:

```bash
python3 rosetta-extensions/stack-core/demo_knob_guard.py <rosetta-root>
# demo-knob-guard: parsers expose 27 env knob(s) + 10 cli flag(s) across 2 entry point(s)
```

## 3. Offline-authored, never driven — **unit-proven ≠ route-proven**

**A defect class, not a run of bad luck.** Four artifacts shipped wrong for one reason — each was *authored
against a reading of the code* rather than *driven against the running system* — and **each was defended by a
green test**, because the test encoded the same reading:

| the artifact | the assumption | reality when driven |
|---|---|---|
| a manager route | built from a **user** id | the route takes a **membership** id — the query nulled |
| the academy CTA | `/library/<slug>` starts a path | **no such route exists** — 404, every time |
| `managerKind: skill-paths` | a manager surface exists | the page renders the literal **"Coming soon"** |
| a skill-path version | `"2"` looked right | a guess; never checked against a row |

**Why a test cannot save you here.** Unit tests prove the code does what the author *believed*. When the belief
is about an **external contract** — a route, a column set, an id's meaning, a rendered surface's existence —
the test inherits the belief and confirms it. The green is real; it is answering a different question.

> **The rule:** an assertion about a **route, an id's meaning, a schema, or a surface's existence** is not
> knowledge until it has been **driven live once**. Until then, mark it in the code as uncalibrated. When a
> live probe contradicts a green test, **the test is the prime suspect** — and the fix is to invert it so it
> fails loudly if the defect returns.

**Corollary for planning:** offline iters produce *risk*, not *score*. M236 iter-03 found the full seeded
substrate present — up to 26 pairs' worth of data — and still recorded the metric as **0/31**, because the gate
measures *renders live* and no render had been proven. Recording 26 would have been the same failure in the
opposite direction. **Read a substrate reading as "the remaining work is render work, not seed work" — never
as partial credit.**

## 4. Measurement hygiene — four ways a measurement lies

Each of these produced a confident number that was about something other than its subject:

- **A suppressed error channel measures the wrong thing silently.** A `git fetch` with stderr suppressed failed;
  the freshness check then compared **stale against stale** and reported agreement. *Never suppress the error
  channel of a command whose failure changes the meaning of the result.*
- **A guard that fails OPEN is worse than no guard**, because everything downstream trusts it. The verdict-age
  check parsed a UTC timestamp as local time: east of UTC it failed closed, but **west of UTC it inflated the
  window and read a STALE verdict as FRESH** — the exact hazard it existed to prevent, inverted, for half the
  world. *(It was itself introduced by a hardening pass. Code written to close a hazard is not exempt from it.)*
- **A denominator counted from survivors flatters every score.** A pair the manifest cannot form writes no
  ledger line, so counting "rows in the ledger" silently shrinks the denominator and every remaining pair
  landing reads as a clean sweep. **Pin the denominator from outside the thing being measured**, and make a
  drop fail the run.
- **A probe intended to discriminate between two hypotheses must not be constructed from the artifact under
  suspicion.** If the artifact is wrong, the probe is wrong in the same direction, and the run confirms the
  hypothesis it inherited rather than testing it.

> **The unifying question, and it is the same one as rule 1 from a different side:** *if the thing I am
> measuring were broken in the way I suspect, would this measurement look any different?* If not, you have
> built a mirror, not a probe.

---

## Related
- [Demo family index](README.md) · [Frontend tier](frontend-tier.md) · [Verification net](../verification.md)
- [Demo lifecycle](../rosetta_demo.md) · [Browser login recipe](recipe-browser-login.md) · [Stories & heroes](stories-spec.md)
- [Latency budget](latency-budget.md) — the perf gate beside these render gates (and the arithmetic-signature model)
- [Content stories spec](content-stories-spec.md) — the `content-manifest.json` Sweep 2 reads
- [Demo-up defaults](demo-up-defaults.md) — the fenced knob/flag contract rule 2 cites
