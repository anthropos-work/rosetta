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

## The gate (objective, machine-verifiable)

A Playwright sweep, logged in as the vantage's hero via the cockpit handshake, of **every reachable demo
page** asserts BOTH:
- **(a)** non-empty **semantic content in the DOM** — real text/rows per section, not just a shell; AND
- **(b)** a populated **screenshot**,

for **100% of pages**, with **ZERO pages empty/error** AND **ZERO nav links escaping the demo platform** —
every in-app link/nav resolves to a **demo-local surface on its offset port** (e.g. left-menu "Studio" → the
local studio-desk, NOT prod `studio.anthropos.work`). An external link is **NOT valid filler**.

**Gate = the sweep's coverage report shows `0 failing pages + 0 escapes`.**

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
2. **Crawls** the in-app nav as that vantage — **pure in-app nav-link discovery** (BFS from the landing page
   over same-origin links + the persistent nav chrome), NOT a static route manifest. Rationale: a manifest
   cannot catch a nav that **escapes** the demo (an external link is invisible to a route list); the gate
   requires escape-detection, so the crawl must observe the actual rendered links. The frontier is capped +
   deduped; query-only variants of the same path collapse.
3. **Per page** asserts: (a) DOM **non-emptiness** — real text/rows per section, above a per-page semantic
   floor (see below); (b) captures a **screenshot**; (c) asserts **every link host is demo-local** — the host
   resolves to a `localhost:<base+offset>` surface, never a prod host (`*.anthropos.work`, real Clerk, etc.).
4. **Emits** a coverage report: each page → `{path, status, empty?, error?, escapes[]}`, plus the roll-up
   `{pages, failing, escapes}` the gate reads.

### The "non-empty semantic content" assertion shape

The gate's (a) clause is the false-pass/false-fail risk. The protocol uses a **two-tier** assertion, tuned
per-iter against the rate the sweep surfaces:
- **Tier 1 — generic text-density floor (default):** the page's main content region must carry more than a
  shell's worth of real text (a per-page minimum visible-text length, excluding nav chrome/footer), AND must
  not contain an error sentinel ("Something went wrong", a stack trace, a bare "No data", a 404/500 body).
  Cheap, catches the dominant empty-page mode.
- **Tier 2 — per-section DOM selectors (escalated when Tier 1 false-passes/-fails a specific page):** for a
  page where the density floor is wrong (a legitimately terse page, or a shell that's text-heavy but
  data-empty), assert the page's known content sections by selector (e.g. the activity feed's row list, the
  library shelf's card grid). Added per-page only as the sweep proves the floor insufficient — avoids
  over-fitting the whole sweep to brittle selectors up front.

The choice between tiers is an **iter decision**, recorded in the iter's `decisions.md` when a page escalates
to Tier 2.

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
| **Federation / content error (403/500/panic)** | the page reads replayed content not serve-granted | `stack-snapshot` serve-grants (`directus/structure.go`) | re-replay snapshot into the demo |
| **Out-of-demo link (escape)** | a baked/rendered link host points at prod | the demo **injection + env link-rewriting** (`demo-stack/up-injected.sh` build-args / `stack-injection/gen_injected_override.py`) — rewrite the host to the offset port | re-build the frontend (baked URL) or re-emit the override + restart |
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
- **The demo must be live + at the consumed tag.** The sweep runs against `demo-N` on offset ports; the demo
  consumes `rosetta-extensions @ <tag>`. A harness/fix change is **authored** in the authoring copy; to be
  reflected in the sweep it must be **applied** to the live demo (re-seed/re-replay/re-build/re-export — or, for
  a harness-only change, run the authoring-copy harness against the live demo directly).
- **Scope = the vantage's reachable set.** The crawl is bounded by what the hero's seat can navigate to; pages
  no link reaches are out of scope (the gate is over *reachable* pages).

## Related
- [Demo family index](README.md) · [Frontend tier](frontend-tier.md) · [Verification net](../verification.md)
- [Demo lifecycle](../rosetta_demo.md) · [Browser login recipe](recipe-browser-login.md) · [Stories & heroes](stories-spec.md)
