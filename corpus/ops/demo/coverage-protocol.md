# The demo-coverage protocol — the Playwright sweep + triage + fix loop

**The iteration protocol for the 100%-coverage milestones (v1.10 "method acting" M42e employee, M42m
manager).** A coverage milestone proves that a hero of a given **vantage** (employee/member, or manager),
logged into a demo stack via the presenter cockpit, sees **100% of the pages that vantage can reach** rendered
with **real semantic content** and **zero out-of-demo escapes**. The page set and the failure modes are
**discovered by the sweep**, not enumerable up front — so the milestone is `iterative`: each iter **measures**
(run the sweep) → **triages** the failures → **fixes** them in `rosetta-extensions` (or a corpus doc) →
**re-sweeps**, until the gate is GREEN.

> **Read first:** [`frontend-tier.md`](frontend-tier.md) (the UI tier the sweep crawls), [`verification.md`](../verification.md)
> (the offset/project/scope-aware probe net this harness sits beside + reuses), [`rosetta_demo.md`](../rosetta_demo.md)
> (the demo lifecycle + offset ports + Clerkenstein injection the cockpit login rides), and
> [`stories-spec.md`](stories-spec.md) (the roster hero — Maya for employee — the sweep logs in as).
>
> All the harness code lives in the gitignored `rosetta-extensions` monorepo (authored + tagged in the
> authoring copy, consumed per-stack at a pinned tag) — **zero platform-repo edits** (the v1.10 hard line).

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
| **Out-of-demo link (escape)** | a baked/rendered link host points at prod | the demo **injection + env link-rewriting** (`demo-stack/up-injected.sh` build-args / `stack-injection/gen_injected_override.py`) — rewrite the host to the offset port. **Precondition (M42m iter-02):** the platform must expose a **per-URL `NEXT_PUBLIC_<thing>_URL` override** for that host (rewritable in the gitignored `apps/web/.env.local` overlay or a build-arg, zero-edit — e.g. next-web's `ACADEMY_URL` reads `NEXT_PUBLIC_ACADEMY_URL`). If the host is instead behind a **coarse mode-flip** (`NEXT_PUBLIC_NODE_ENV`) or a **hardcode** with no per-URL knob (e.g. next-web's `STUDIO_URL` — a `NEXT_PUBLIC_NODE_ENV` ternary, wrong-port + side-effecting on flip), the host is **platform-bound** → this row does NOT apply; it's a **re-scope trigger** (the rewrite needs a platform-source edit). Diagnose: find the constant's source, check for a dedicated `NEXT_PUBLIC_<thing>_URL` read vs a mode-flip/hardcode. | re-build the frontend (baked URL) or re-emit the override + restart |
| **Platform-bound escape (no per-URL override)** | a baked link host is hardcoded / behind a coarse mode-flip with no `NEXT_PUBLIC_<thing>_URL` knob (e.g. next-web's `STUDIO_URL`) — the env-rewrite row above does NOT apply | the **demo-patch tool** (`demo-stack/patches/demopatch` + a content-anchored manifest, M42m iter-03): source-patch the demo's **EPHEMERAL gitignored clone** before the build to read `NEXT_PUBLIC_<thing>_URL` (a behavior-identical fallback ternary kept), then **trap-revert** after the image bakes — CANONICAL repos NEVER touched (6 guards: hard path-assert demo-clone-only, drift-refuse, never-commit, idempotent, self-owned reversal, demo-only). Wired into `up-injected.sh` (apply-before-build + RETURN-trap revert) with the offset value in the `.env.local` overlay; default-on + non-fatal (`DEMO_NO_PATCH=1` opts out). The clone is left git-clean; `ensure-clones.sh` **R1** pristine-reverts a crash-left patch + **R1b** sweeps a crash-left tooling `.dockerignore` (byte-identical + untracked guards). Resolved the Studio `studio.anthropos.work` escape demo-only (139→0). | re-build the frontend (the patch bakes into the image; revert is automatic) |
| **Cross-port demo-local surface blank / login-loops / wrong-eject** (studio-desk class, v1.10 postfix) | a demo-local link to a **different port** than next-web (e.g. "Anthropos Studio" → studio-desk `:9000+offset`) doesn't render the authenticated home for the **logged-in hero** — it opens a blank `/undefined`, a dead `:3000` redirect-loop, a `/login` loop, or ejects to `WEB_APP_URL` (the non-admin redirect when the hero's membership doesn't resolve) | **authenticate via Clerkenstein, never bypass**: studio-desk drives its **own** fake-FAPI handshake (per-app — the cross-port `__session` cookie is **not** needed; the FAPI holds the active seat server-side) verified networklessly via `CLERK_JWT_KEY`. The **injection override** (`gen_injected_override.py`) wires the runtime `CLERK_*` (← `DESK_CLERK_*` minted by `up-injected.sh`) + pins `CLERK_SIGN_IN_URL`/`WEB_APP_URL` at the offset next-web (the requireAuth fallback). For the **admin gate** (`checkEnterpriseAndAdmin` reads the fake BAPI's `getOrganizationMembershipList`), make the fake BAPI **roster-aware** (`cmd/fake-bapi` reads the same `FAKE_FAPI_ROSTER` + seeds each hero's `(org, user)→org_role`) so a manager passes and an employee is correctly redirected. The harness gate is closed by the crawl's **cross-port FOLLOW** (`crawl.ts` `onCrossPortFollow` → `coverage.spec.ts`) **in the logged-in context**: a blank/login-loop/un-offset-`:3000`/eject destination FAILS the gate | re-up (the roster re-seeds + the override re-emits, no rebuild) for the env/BAPI; **clear the cached image** (`docker image rm demo-N-studio-desk`) only to re-bake a stale pk/offset |
| **Org-scale grid perf wall (slow GraphQL, not empty)** (M46) | at org scale (~500 members) a heavy enterprise grid (`/enterprise/members`, `/enterprise/activity-dashboard`, `/enterprise/settings`) sits on a `…` spinner / skeleton through the whole warm-grid poll because its backing **federated GraphQL** takes 10–84 s — an over-broad fetch (`InsightsContext.tsx` loads `limit: 1000` = ALL members) AND a per-membership resolver fan-out (`jobRole`/`targetRole`/`tags` × a **per-object Sentinel RPC** in `app` `roles.go GetOrganizationTargetRole`, no DataLoader). DIAGNOSE it's perf-not-content by `docker logs <stack>-graphql-1` + `grep latency` (10 s+) while raw SQL is ms (so it's RPC-round-trip-count, not the DB). **Decompose it — part demo-patchable, part platform-bound:** | **The over-broad-fetch + missing-index part is demo-local (zero canonical edit), and clears the grids whose cost is fetch-width or DB-index-bound:** **(1)** a **next-web pagination demo-patch** (`patches/next-web-members-pagination`: `InsightsContext.tsx` `limit 1000→30` — the grids that already paginate, e.g. `/enterprise/members` at 20, are untouched; the CSV/email export uses a separate query so no data is lost) applied to the demo's ephemeral clone pre-build + trap-reverted (the demopatch 6-guard contract); **(2)** **post-seed FK indexes** (`CREATE INDEX IF NOT EXISTS` on `membership_skills(membership_skill_membership)` + `membership_tags(membership_tag_membership)`) on the demo's **own** Postgres via a rext-owned `up-injected.sh` step (idempotent, non-fatal — NOT a canonical ent/atlas change). On demo-3 (~500 members) (1)+(2) took graphql max latency **84 s → ~4 s** and cleared `/enterprise/activity-dashboard` + `/enterprise/settings`. **BUT the `/enterprise/members` per-row `targetRole` → `OrgCheckActionPermission` Sentinel RPC is `OrgActionAssignmentsWrite`, which is PER-OBJECT (per assignee), NOT an org-wide grant** — a manager legitimately can-write-assignments for some members and not others. CACHING that check is a **correctness bug** — keyed `(org, subject, action)` (object dropped to dedupe across rows) it returns the first row's allow/deny for all rows → `failed to get target role: forbidden` on legit members (~1744×/sweep) → the grid errors (this was T2, reverted); keyed correctly `(org, subject, OBJECT, action)` it cannot dedupe (every row = a different object). **The safe demo-patch is to DROP the read-gate, not cache it (Option B, the M46 close):** `roles.go RoleManager.checkPermission` short-circuits `return true, nil` BEFORE the per-member Sentinel RPC (mirroring its built-in `privacy.DecisionFromContext` bypass) — target roles still come from the DB so every member's REAL role renders (fast, fully-populated, 0 forbidden). Read-path relaxation ONLY (`patches/app-targetrole-authz-skip`, applied to the build-scratch app clone via a rext helper wired into the inject loop, svc=app, after `apply-authn`, before build, trap-reverted) — the assignment **mutations** still enforce via their own direct `OrgCheckActionPermission` calls. On demo-3 B took the members query **76.7 s → 0.51 s** and **cleared `/enterprise/members` → the manager gate is MET (no re-scope)**. The PLATFORM finding stays: prod needs a **DataLoader / batch `BulkCheckPermission` RPC**; B is a disclosed demo-perf relaxation, not a prod fix. **⚠ if you rebuild the injected `app` image, it MUST go through the inject loop so `apply-authn.sh` (the disarmed colony) is re-applied** — a rebuild without it ships a backend that hits real `api.clerk.com`, rejects every Clerkenstein token, and collapses the crawl to `reachable≈7` (a broken-auth artifact, not a content fail — grep `docker logs <stack>-backend-1` for `clerk`); and **never recreate one service with `--force-recreate` *without* `--no-deps`** (it recreates `postgresql` too and wipes the seeded org). | re-build the frontend (pagination bakes in) + re-build the injected `app` (the authz-skip bakes in) via the inject loop + re-up/re-seed (indexes apply post-seed) — all three demo-local, gate MET |
| **Editorial citation in replayed content** (NOT an escape) | a real `<a href>` to an external article baked into replayed `/skill-path/.../chapter` body copy | the harness **citation allow-rule** (`coverage.spec.ts` `allowedExternalLink` → `crawl.ts`): classify the off-demo link on a `/chapter` path as a VALID citation, recorded as a **presenter note**, NOT counted as an escape (M42e iter-08) | (none — content fidelity; do NOT strip/rewrite the citation) |
| **Wrong identity / org on a surface** | roster/FAPI resource gap | `stack-seeding/seeders/roster.go` + `clerkenstein/clerk-frontend/resources.go` | re-export roster + restart `<demo>-fake-fapi-1` |
| **A documented gap** | the behavior is correct but undocumented | a corpus doc update | (none) |

### Re-scope trigger (the zero-edit line)

If a 100%-blocking failure can **ONLY** be closed by a **platform change** (next-web / app / cms /
jobsimulation / skiller are read-only this entire release), that is the milestone's **Re-scope trigger**:
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
  job-simulations referenced a skiller skill node-id (`K-AIFUNX-E658`) ABSENT from the demo — the federated
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
    **Build pitfalls (each cost a full re-seed):** the injected Go images are built from a build-scratch clone
    AFTER `apply-authn.sh` vendors the **disarmed colony** (Clerkenstein token acceptance); a standalone `app`
    rebuild that SKIPS that step ships a backend that calls real `api.clerk.com`, rejects every demo token, and
    collapses the crawl to `reachable≈7` (looks like a content fail — it's broken auth: grep
    `docker logs <stack>-backend-1` for `clerk`). And never recreate a single service with `--force-recreate`
    *without* `--no-deps` — it recreates `postgresql` too and wipes the seeded org. Never widen the poll to mask a
    slow query, and never shrink the org below the org-scale premise just to pass.
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

## Related
- [Demo family index](README.md) · [Frontend tier](frontend-tier.md) · [Verification net](../verification.md)
- [Demo lifecycle](../rosetta_demo.md) · [Browser login recipe](recipe-browser-login.md) · [Stories & heroes](stories-spec.md)
