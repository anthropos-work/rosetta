# M224 — Decisions

_Implementation + strategy choices with rationale (incl. tok strategy revisions)._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| — | _(intra-iter decisions live in `iter-NN/decisions.md`)_ | | |

---

## TOK-01: recruiter-render-first — 2026-07-16

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Drive the **render loop** the protocol prescribes — seed → render → **attribute** → fix →
re-render — against a **LOCAL** demo, ordered **recruiter-scoreboard-FIRST**. The exit gate IS the recruiter's
comparison surface (`/enterprise/activity-dashboard → AI-Simulations → [simId]`) painting ≥40 comparable non-junk
candidate rows per **each** of the 5 seeded HIRING positions, non-degenerate distribution, closure green, 0
prod-eject, over ≥3 cold reset-to-seed runs. So every early tik targets **rows-on-the-recruiter-scoreboard**; the
2 candidate `/profile` heroes (In-scope but not the gate) are layered **after** the scoreboard is green.

Each tik's fix is chosen by **attribution against `hiring.md`'s traced read-path**, in the M219-trap order:
  1. **Clerkenstein identity gap** — the client re-skin (`useGetClerkOrganization`) reads org
     `publicMetadata.isHiring` from the **FAPI**; without it the org renders as a normal Workforce org and the
     "Results" framing / hiring cohort treatment never appears. **Prefer the Clerkenstein FAPI wiring** (roster+
     resource, M39 `org_name`/`org_slug` precedent) over any patch (D-DESIGN-2). **Touches `clerk-frontend/` →
     carries the BLOCKING `/align-run` on `clerk-js-5`** (keep `Client/signed-in` critical/shape +
     `Me/universal-user` standard/shape green; the named `SessionToken/decoded-identity` exact gene is unaffected).
  2. **Render-gate bypass** — an M219-class gate (a `CycleID==nil→buildLiveResponse` analog, an undefined
     PostHog flag, or the silent-403 on the `OrgFeatureInsights` Casbin permission) that ignores the seed → route
     to a **sha-pinned demo-patch** on the demo's ephemeral clone (read `demopatch-spec.md` first) only if there is
     no env/config/Clerkenstein seam.
  3. **Seed gap** — a row the scoreboard reads that the M223 funnel didn't write (e.g. a missing membership, a
     NULL-bubbling federated `Session!` twin, a drill-down `validation_*` row for the assessed candidate hero) →
     `stack-seeding`.

**Rationale:** The opening move is "reach + measure + attribute" NOT "fix", because M223 **already** writes the
score source (the `local_jobsimulation_sessions` mirror pair, 45 candidates × 5 sims, verified ALIGNED in Phase
0b). The render risk is therefore concentrated in the render-GATE and the Clerkenstein re-skin identity — exactly
where a fix-before-measure would be a guess (the M219 trap the milestone is named against). Front-loading a real
baseline render makes the first fix land against a measured, attributed gap.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).

**Distance-to-gate context:** metric = `min(rows-painted across the 5 sims)` on a cold-seeded recruiter scoreboard
(+ the distribution / closure / eject sub-gates). Baseline = **UNMEASURED** (the surface has never been rendered on
a local stack — no recruiter seat, no isHiring wiring). Gate = **≥40** per each of 5 sims, non-degenerate, closure
green, 0 ejects, ≥3 cold runs.

**Next-tik direction:** iter-02 (first tik) = the **enabling-scaffold + baseline-render** tik. (a) Add the
recruiter (`vantage: manager`) hero seat to the Meridian Talent 4th story in `presets/stories.seed.yaml` with
`jump_to` at the comparison surface; (b) wire Clerkenstein FAPI `isHiring` (`clerk-frontend/registry.go` RosterEntry
`org_is_hiring` → `DemoUser.OrgIsHiring` → `resources.go::orgMemberships` `PublicMetadata["isHiring"]`) **+ the
`/align-run` gate**; (c) author a **recruiter render-probe** under `stack-verify/e2e/` (log in via the cockpit
handshake, reach each of the 5 `[simId]` scoreboards, count comparable rows + capture the score distribution + run
closure/eject checks); (d) tag rext; (e) bring up a **LOCAL** demo consuming the tag; (f) log in as the recruiter,
**measure the baseline**, and **attribute** the per-sim gap — no fix before the attribution. If reaching/measuring
is blocked on a missing probe capability, iter-03 becomes a **tooling-iter** (protocol refinement).

---

## TOK-02: run-the-real-hiring-app (two-app demo, not a workforce re-skin) — 2026-07-16

**Tok type:** triggered — strategy revision (user-directed, after the iter-04/05 render evidence).

**Trigger (what falsified TOK-01's fix path):** the render loop reached + measured + attributed, and the
attribution pointed at a **product boundary, not a render-gate** — exactly the class TOK-01's fix ladder did not
anticipate.
- **iter-04 (first gate reading):** baseline `min(rows-per-sim) = 0`. Attribution = **SEED-GAP** (the M219 trap,
  caught by measure-first): the recruiter hero's role `"Technical Recruiter"` did not resolve to a real public
  `job_role` (no `job_role_skills`) → the users-seeder fail-fast killed the ENTIRE hiring seed (Meridian = 0
  members, blast-radius to the other orgs). **Fixed** (role → `"Talent Acquisition Specialist"`, a real resolvable
  role). After a cold reset-to-seed: **Meridian Talent = 50 members; each of the 5 shared sims = 43 comparable
  candidates; scores 27–100, non-degenerate. The DATA side of the gate is MET.**
- **iter-05 (diagnosis):** with the data present, the recruiter STILL cannot reach the scoreboard. **Root cause
  (code-cited):** `apps/web/src/context/UserStatusContext.tsx:141-173` — a user whose memberships are ALL
  hiring-orgs is **ejected out of the workforce app** to `hiring.anthropos.work` (`buildSwitchHandoffUrl`,
  `targetProduct:'hiring'`), **by design** (the platform enforcing the product boundary); `useGetClerkOrganization`
  filters hiring orgs out of the workforce list. The Hiring sub-app is **not in the demo**. So **"genuinely reads
  as hiring"** (needed for the re-skin — Scope.In#3 / D-DESIGN-1) and **"scoreboard reachable in apps/web"** are
  **mutually exclusive on the unmodified platform**. This **falsifies the M222 premise** (M222's "comparison lives
  in apps/web, reachable by the recruiter" held only because that spike org lacked client `publicMetadata.isHiring`
  — it read as workforce). Also landed a cheap probe fix (render-probe timeout 150s→300s; kept the fullPage-hang
  fix).

**Old strategy (TOK-01, superseded on the FIX target only):** render the comparison INSIDE `apps/web`, treating
the eject as a render-gate to neutralize with a single sha-pinned demo-patch + **re-skin the workforce Results
surface to imitate hiring**.

**Why the old strategy is inferior — not merely harder:** it keeps the recruiter in the **workforce** app wearing
hiring chrome — a **demo fiction** that actively suppresses the platform's real product routing. And most of its UI
work is **wasted**: the genuine hiring candidate-comparison screen **already ships in `apps/hiring`**; the
single-patch path rebuilds a fake of a screen that exists as-is.

**New strategy (TOK-02) — run the REAL Hiring app in the demo (a two-app demo).** Build `apps/hiring` into the
demo as a **second UI container** — the same recipe the demo already uses for `apps/web` + `studio-desk`, from the
**untouched clone** — on its own offset port, wired to the **same** fake-Clerk FAPI, the **same** Cosmo backend,
and the **same** seeded Postgres. The recruiter hero logs in via the cockpit **straight onto the Hiring app's
Results page**. The platform's **own** routing then does the work: the workforce app ejects a hiring-org recruiter,
and the Hiring app's **symmetric** guard (`apps/hiring/.../UserStatusContext.tsx:119-149`, `userHasAllWorkforceOrgs`
short-circuit → no redirect) **keeps her in**. No forcing, no fiction.

**Premise VERIFIED end-to-end (adversarial, against the demo clone — see the 2026-07-16 two-app feasibility
workflow):**
- **The comparison screen ships in `apps/hiring`:** route
  `apps/hiring/src/app/(authenticated)/(verified)/enterprise/activity-dashboard/@tabs/ai-simulations/[simId]/page.tsx`
  ("Enterprise Results | AI Simulations by Members") → `InsightsByMembersContainer` → a `Table` with
  `simulationScoreColumn` rendering each candidate `{score}/100` (pass/fail colored), plus `CandidateSearchInput`,
  sorting, and all/shortlisted/archived tabs — a genuine ranked candidate comparison per sim.
- **Same data + backend, no new plumbing:** the `insightsJobSimulationByMemberships` field is in the **app
  subgraph** SDL with **no feature gate** (`app/.../graphql/graph/schemas/queries.graphqls:188-192`), federates the
  **same Cosmo router the demo already bakes**; the resolver (`intelligence.go:1681`) reads
  `public.local_jobsimulation_sessions` filtered by `jobSimulationId` + org + memberships, default sort DESC by
  score — **exactly where M223 seeded 5 sims × 43 candidates**. So: **no re-skin, no new resolver, no data
  migration.**

**Faithfulness:** materially MORE faithful than the patch — it shows the ACTUAL product (hiring recruiters really
use the separate Hiring site) and exercises the platform's genuine routing.

**Constraint compliance (zero platform-repo edits):** the Hiring app builds from the **untouched clone** (build
context only). A **rext-owned** `hiring.Dockerfile` (the platform `Dockerfile.dev` hardcodes
`--filter=@anthropos/web-app` / `start:web` / `EXPOSE 3000`, so the demo cannot reuse it verbatim) + the demo-stack
tooling own everything. **At most ONE sanctioned demo-patch** (`core-js/src/constants/urls.ts` `HIRING_APP_URL`,
**chained** onto the existing `studio-url`/`public-website-url` patch on that same shared file) to point the in-app
"go to Hiring" link at the demo-local hiring app — **or bypass it entirely** via the cockpit deep-link.

**Strategy class:** new-direction on the FIX target. The render-loop **measure → attribute → fix → re-measure**
spine (TOK-01) is unchanged; TOK-02 changes only WHERE the fix lands, once the attribution resolved to a
product-boundary eject rather than a render-gate.

**Distance-to-gate context:** metric unchanged = `min(rows-painted across the 5 sims)` on the recruiter's Results
scoreboard, now measured **in the Hiring app**. **Data side MET** (43 × 5 in `local_jobsimulation_sessions`, scores
27–100). **Render side UNMEASURED** in `apps/hiring` (the app is not built into the demo yet). Gate unchanged: **≥40
per each of 5 sims, non-degenerate, closure green, 0 ejects, ≥3 cold reset-to-seed runs.**

**Next-tik direction (the build tiks under TOK-02):**
- **tik A (iter-07) — the hiring-app UI container.** A rext-owned `hiring.Dockerfile` (filter the hiring package,
  its `start` cmd + `EXPOSE`) + a `FRONTENDS` row in `demo-stack/…/gen_injected_override.py` (offset port,
  fake-Clerk pk + FAPI URL bake, `CORS_EXTRA_ORIGINS`, `UI_BROWSER_FACING`). Bring up demo-1; confirm the hiring app
  boots and reaches the same Cosmo backend + fake FAPI.
- **tik B (iter-08) — cockpit + probe re-point + first Hiring-app render reading.** The recruiter seat gets a
  hiring base + `jump_to` the Results page; confirm the cockpit login lands the recruiter **on** the Hiring Results
  scoreboard (the platform's guards keep her in). Re-point the render-probe at the hiring app; **measure**
  `min(rows-per-sim)`.
- **tik C (iter-09+) — drive to green.** Attribute + fix any residual: the enterprise/admin role gate on the hiring
  routes (confirm the seeded recruiter has an admin/appropriate role in the hiring org, not merely a member); the
  optional `urls.ts HIRING_APP_URL` chain-patch (only if we want the in-app eject to land demo-local rather than
  relying on the cockpit deep-link); a possible `hiring-layout no-thirdparty` patch for self-containment. Gate: ≥40
  × 5 over ≥3 cold runs.
- **Post-gate:** the 2 candidate `/profile` heroes (unchanged — still layered after the scoreboard is green).

**Open risks carried (from the feasibility workflow):**
1. **In-app eject target is baked** (`core-js/urls.ts HIRING_APP_URL` → prod `hiring.anthropos.work` or dev
   `localhost:3001`; neither matches an offset/tailnet demo). Mitigation: the chained `urls.ts` demo-patch, OR steer
   the presenter in via the cockpit deep-link (bypasses the eject).
2. **Auth/role gate:** the hiring app's `enterprise`/Results routes likely require an enterprise/admin membership
   (analogous to studio-desk's `checkEnterpriseAndAdmin`) — the seeded Meridian recruiter must hold an admin role in
   the hiring org, not just be a member.
3. **GraphQL resolver coverage:** confirm `insightsJobSimulationByMemberships` (+ `useGetSimulationDetails`) is in
   the demo's running federated schema (same `@anthropos/graphql` + router the demo bakes — likely fine, verify).
4. **Hiring layout third-party scripts:** `apps/hiring` has its own `layout.tsx`; if it loads egress scripts, a
   self-contained demo may need a parallel `hiring-layout no-thirdparty` demo-patch (quick read to confirm).
5. **Build residual:** a 2nd cached ~3-min frontend image per new demo-N (documented in `frontend-tier.md`).
6. **Cockpit manifest:** the recruiter's `jump_to` + a hiring/vantage flag must be authored (`stackseed
   --cockpit-export`) so the cockpit picks the hiring base for that hero; else she jumps to the web base and is
   ejected.

---

## GATE-DECISION D1: keep the platform-native 20-per-page (faithful pagination) — 2026-07-16

**Context.** iter-07→iter-10 drove the two-app render loop to the payoff: the recruiter logs into the demo's real
`apps/hiring`, is kept in (not ejected), reads as admin, and her Results scoreboard **renders real ranked
candidates for ALL 5 shared sims** (`insightsJobSimulationByMemberships` 200, 0 errors, 0 403). Two sha-pinned
demo-patches landed it — `next-hiring-role-remap` (iter-09, the admin remap) + `next-hiring-members-pagination`
(iter-10, the unbounded-fetch unblock); the per-member authz was already covered by the existing
`app-targetrole-authz-skip`. **Zero platform-repo edits.**

**The 20-vs-43 finding.** Each of the 5 shared sims returns exactly **20 rows on page 1** — the platform's Results
table paginates at 20 (`apps/hiring/.../InsightsByMembersContainer.tsx` `useTablePagination`, default 20). All **43**
seeded candidates per sim EXIST, are non-degenerate (scores 27–100), and are **reachable by paging**; page 1 shows
the platform-native 20. The exit gate's "≥40 rows per sim" was written assuming all candidates on one page — before
we knew the scoreboard paginates at 20.

**User decision (2026-07-16): keep the platform-native 20/page** — the faithful option (the real product paginates
at 20; a page-size bump would diverge from production). **Rejected:** a `next-hiring-page-size` demo-patch to force
~43 on one page.

**Gate re-interpretation (consequent).** The exit gate is satisfied as **"≥40 comparable candidates seeded +
reachable + rendering per sim, under the platform-native pagination (page 1 = 20)"** — NOT "≥40 rows painted on a
single page." The data-completeness sub-gates are UNCHANGED and already met (≥43 comparable, non-degenerate,
closure-green, 0-eject per sim); only the on-screen page-1 count reflects the real 20/page. This keeps TOK-02's
faithfulness thesis intact (show the real product, real routing, real pagination). The **≥3-cold-run** requirement
and the **render-probe fixes (R1–R4)** remain open iter-close work (below / carry-forward).

> **SUPERSEDED by v2.4 M227 fix #3 (2026-07-17, live-feedback believability backfill).** The `≥40`-per-sim gate
> assumed **every candidate auditioned on all 5 positions** (~43/position). Live feedback found that unbelievable —
> a real applicant auditions for the ONE role they applied to. M227 changed the funnel so each candidate takes
> **exactly one** position (round-robined evenly → **~8 per position**), and **retuned the compare gate `≥40 → ≥6`**
> (`hiringComparableFloor` / the render-probe `RENDER_GATE_FLOOR`, a small margin below the seeded min of ~8). The
> "page 1 = 20" pagination nuance of D1 is now moot for the demo (a ~8-candidate cohort fits on one page); the
> platform still paginates at 20, so D1's faithfulness point stands. **M228 re-proves the retuned gate live on
> billion.**

---

## Close decisions (2026-07-16 — `/developer-kit:close-milestone`)

Deferral re-audit `audit-deferrals/deferral-audit-2026-07-16-m224-close.md` → **GREEN** (no repeat/chronic/aged-out).
Scope review surfaced three overlooked `Delivers → knowledge/corpus` sections in addition to the known demopatch
fold-in; all landed **Fate-1** at close.

| # | Decision | Fate | Rationale |
|---|----------|------|-----------|
| **D2** | **`demopatch-spec.md` §4/§5 fold-in landed.** The HIRING image (`build_frontend_hiring`) bakes **4** patches — 2 net-new `apps/hiring` (`next-hiring-role-remap`, `next-hiring-members-pagination`) + the 2 chained shared `urls.ts` (`next-web-studio-url` → `next-web-public-website-url`, the iter-13 Studio-eject kill via the shared `packages/ui` NavBar); the chain runs on **both** frontend builds; distinct-manifest total unchanged at 11. | **Fate-1** | The known progress.md fold-in (a); a small in-domain doc update done properly at close. |
| **D3** | **`cockpit-spec.md` hiring-vantage section landed.** The hero trio (Rae/Cara/Cody), the `CockpitHero.IsHiring` two-app-base routing, the candidate `/home` faithful landing, and the optional-polish DeepLinkCatalog `NeedsID` call. | **Fate-1** | Named in `overview.md` `Delivers →` + the KB-fidelity audit; had **zero** M224 content pre-close — a silent `Delivers` gap, landed now. |
| **D4** | **`clerkenstein.md` isHiring-FAPI section landed.** The org `public_metadata.isHiring` emission (`resources.go::orgMemberships` ← `RosterEntry.org_is_hiring`), the CONDITIONAL-EMIT align-safety rule, the `/align-run` 100/100 record, and the BAPI-intentionally-not-wired call. | **Fate-1** | The KB-fidelity audit named the clerkenstein FAPI org-publicMetadata BLIND-AREA an "M224 `Delivers →` deliverable"; content was verified in `spec-notes.md` — landed now. |
| **D5** | **`hiring.md` render-path corrected.** Added § *The render path (M224 — the two-app demo)* (TOK-02: render lives in the real `apps/hiring`, not `apps/web`; same field/router/mirror-table) and fixed the stale "scoreboard is in the dockerized `apps/web`" cross-ref. | **Fate-1** | Phase-3 stale-claim fix; the shipped reality contradicted the doc (the milestone's own `Delivers → render-path into hiring.md`). |
| **D6** | **8 pre-existing test failures CARRIED (not fixed).** 6 × `demo-stack/tests/test_cockpit.py` (4 removed-academy-CTA + 2 v2.3.1 overlay-JS) + `test_purge` + `test_reap`. Verified HEAD-identical (matched the harden-pass Phase-5 record byte-for-byte: `650 passed / 8 failed`), in files M224 never touched, predating v2.4 (v2.3.1/v2.3.2 cockpit hotfixes). | **Carry (inherited known-issue)** | Outside M224's hiring render-loop domain; fixing = scope-bleed. Routed to the standing test-debt backlog (a future demo-stack test-debt harden pass). Non-blocking to the gate (a live render probe, GREEN). Flagged in the Gate Outcome Ledger. |

**DeepLinkCatalog `NeedsID` per-`[simId]` entry (scope sub-item, `overview.md` Scope.In #2).** Judged **optional
polish** and consciously **not added** — the recruiter's raw `jump_to` to `/enterprise/activity-dashboard` suffices
and the render gate was met without it (spec-notes iter-03). Not a silent drop: a deliberate scope call, recorded
in the Gate Outcome Ledger.

