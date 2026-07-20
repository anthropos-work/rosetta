# M236 — Progress

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 "publish-then-prove" authored; baseline measured live — gate denominator is **31** landable (session × action) pairs, currently **0/31**, blocked by an unpublished-tooling gap (`billion` pins the M228 tag; 0 of 13 `playbill-*` tags are on origin) — see iter-01/progress.md

- iter-02 (tik): **closed-fixed** — Phase P complete. Published `rosetta-extensions` `main` (`1d97861..60eff14`) + all 13 `playbill-*` tags to origin; re-pinned `billion` → `playbill-m235-hardened`; M217 pin guard PASSES; **31 denominator confirmed against the shipped manifest**. Metric 0/31 unchanged by design (bought reachability, not numerator). Side-deliverable: 109 GB host cache reclaimed — see iter-02/progress.md

- iter-03 (tik): **closed-fixed** — Phase C complete. Stale M228-era demo-1 purged; cold rebuild at `playbill-m235-hardened` UP with **fresh-green autoverify**; cockpit serves `/content-manifest.json` HTTP 200 with 4 products / 18 sessions; substrate verified 13/13 sim sessions + attempt-results + manager mirror, ai-labs presence 2/2. Academy tables empty (in-scope, routed). Metric held honestly at 0/31 — no render proven yet. Blocker found + cleared: remote bring-up needs a LOGIN shell (Go on PATH only via profile) — see iter-03/progress.md

- iter-04 (tik): **closed-fixed** — Phase H. Calibrated the result render against the LIVE stack (3 shapes found, 2 of which a blind author would have gotten wrong in opposite directions), then authored the content-stories sweep (`lib/content-result-page.ts` + `tests/content-stories.spec.ts` + runner + aggregator, rext `playbill-m236-content-sweep`). **Metric 0/31 → 16/31** — all 13 simulation PLAYER pairs land. Fixed 2 harness bugs that produced confident wrong numbers (serial-abort; worker-restart amnesia) and 1 **false PASS** (skill-path graded as a scored sim). B4 + B5 landed: `coverage-protocol.md` premise corrected + content-stories section; `playthroughs.md` backfilled — see iter-04/progress.md

- iter-05 (tik): **closed-fixed** — the MANAGER vantage. Root-caused to ONE fact: the activity-dashboard route's last segment is a **MEMBERSHIP id**, not a user id (`GetMembership(membershipsID)` → `ent: membership not found` → the whole query nulls). The page's header comes from a *different* query, so it looked populated while proving nothing. Fixed in tooling (deterministic membership id, zero platform edits), test corrected (it had asserted the defective contract), honesty gate caught the stale preset. **16/31 → 27/31** — see iter-05/progress.md

- iter-06 (tik): **closed-fixed** — the residual. The 2 interview manager pairs were a **FALSE FAIL**: that route renders a *different surface* (breadcrumb + attempts table + "View Report", no `Results for` header) and rendered correctly all along. Added `manager-interview` as the 5th route-selected shape. **27/31 → 29/31; the simulation arm is 26/26 COMPLETE.** Residual 2 characterized: a skill-path manager route that HANGS (>180 s; heavy 13-chapter instance, light sibling passes → a per-item fan-out signature) and the academy pair — see iter-06/progress.md

- iter-07 (tik): **closed-fixed** — the "hang" was **`networkidle`**. A navigation probe found 0 requests in flight and the page painted in ~1 s; the sweep had inherited `loginAs`'s `networkidle` default, which never resolves on a long-polling surface — the rule `latency-budget.md` already states. Reading what then rendered produced the real finding: the skill-path per-user manager view is **unimplemented** ("Coming soon", table commented out, `userData` hardcoded null), so **no query touches the seeded session** — and its sibling had been scoring as a **false PASS** off the definition-only "Results for" header (the iter-05 trap on a second surface). Skill-path becomes player-link-only. **Denominator 31 → 29, numerator 29 → 28; gap to gate 2 → 1** — see iter-07/progress.md

- iter-08 (tik): **closed-fixed** — the academy CTA named a route that **does not exist** (`/library/[slug]` is not a route in ant-academy — only the index) and a slug that is not in its catalog, so it could only ever 404; the unit test *required* that prefix and so defended it. Corrected to `/courses/ai-foundations`, a real course in the FS catalog the demo academy serves. Needed a **sixth** render shape (`player-academy`) — it had been graded as a scored sim. Also established that **`academy-seed` is moot in a demo**: the academy runs with no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, so DB progress has no reader. **28/29 → 29/29 — primary metric MET.** Thread A verified live (65 cards, 0 Draft chips) → M230 cluster 1 **discharged** — see iter-08/progress.md

- iter-09 (tik): **closed-fixed** — the **p95 click→ACCESS** gate component, measured from the tailnet (the presenter's real vantage) against `billion`: employee **3.15 s**, manager **2.71 s**, 5/5 ACCESS each, against a 5 s gate — **MET**. Graded against billion's own fresh green verdict via the runner's already-documented remote seam (never the override). Side-deliverable: the green gate's age check parsed a UTC `ts` as **local** time on BSD — a verdict 121 s old aged as 7321 s; **west** of UTC that reads a STALE verdict as FRESH, the exact F-6 hazard the check exists to prevent, so the guard failed OPEN for half the world — see iter-09/progress.md

- iter-10 (tik): **closed-fixed — GATE MET.** The cold reset-to-seed reproduction. Re-pinned to the final tooling tag, `down --purge`, cold `/demo-up` — **no intervention required**. The corrected manifest arrived from published tooling with no hand-editing, and all three components reproduced: **29/29** pairs · **65** academy cards, 0 Draft chips · hero p95 **1.22 s / 1.51 s**. 0 platform edits verified per-clone. iter-07's cockpit-bind prediction confirmed (cold restores `0.0.0.0`). Also landed the protocol-evolution backfill owed by iters 05–09 — see iter-10/progress.md

## Gate — MET (2026-07-20, cold on billion)

| component | status |
|---|---|
| all landable (session × action) pairs render real content, both vantages | **MET** — 29/29 |
| the academy grid renders real cards (Thread A) | **MET** — 65 cards, 483 chapter links, 0 Draft chips |
| p95 click→ACCESS < 5 s, HERO vantages only (B2) | **MET** — employee 1.22 s · manager 1.51 s, 5/5 ACCESS |
| reproducible on a cold reset-to-seed | **MET** — all of the above on a stack built from nothing |
| 0 platform-repo edits | **MET** — canonical `anthropos-work` repos untouched |

**Next:** `/developer-kit:harden-mstone-iters --final`, then `/developer-kit:close-milestone`.

## Carried forward (none gate-blocking)

- **M230 carry-forward cluster 3** — the anonymous `/library` + `/free` academy routes render 0 cards
  (`getPublicCatalogView`'s `new Set()` branch is uncovered by the M230 patch; the patch manifest names
  the gap itself). Handler `ACADEMY-M236-iter08-public-catalog-twin` → v2.5 release close.
  *(M230 cluster 1 — the rendered-card count — is **discharged** by iters 08/10.)*
- **`apps/web` client GraphQL endpoint** — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql`,
  the non-offset port, while `apps/hiring` carries the correct offset origin. Never manifested on a measured
  path (SSR uses `WUNDERGRAPH_SSR_ENDPOINT`). Demo-hygiene → release close. From iter-08 D5.
- **Remaining v2.4-era method docs** flagged YELLOW by the KB-fidelity audit (F3/F4/F6/F7/F10 —
  `verification.md`, `demo-up-defaults.md`, `tailscale-serve.md`, `session-clone-spec.md`). Backfilled this
  milestone: `coverage-protocol.md`, `playthroughs.md`, `content-stories-spec.md`, `latency-budget.md`
  → the rest to release close.
- **Standing carry:** 14 pre-existing demo-stack test failures (REPEAT) → v2.5 release close.

## Resolved handlers

- `SKILLPATH-M236-iter07-manager-hang` — iter-07 (a `networkidle` wait over an unimplemented surface).
- `ACADEMY-M236-iterTBD-catalog-fill` — iter-08 (the CTA named a nonexistent route; the catalog was never
  the problem, and `academy-seed` is moot in a demo).
- `LATENCY-M236-iterTBD-hero-p95` — iter-09.
- `REPRO-M236-iterTBD-cold-cycle` — iter-10.
- **USER-BLOCKER-M236-01: RESOLVED** 2026-07-20 (B1–B5).

---

## Gate Outcome Ledger (close, 2026-07-20)

### Gate

- **Target (as re-scoped by USER-BLOCKER-M236-01, 2026-07-20):** both cockpit tabs work live on `billion` —
  all landable (session × action) pairs render real, non-empty content for player + manager vantages; the
  academy grid renders real cards (Thread A); reproducibly on a **cold reset-to-seed**; **p95 click→ACCESS
  < 5 s for the HERO vantages only**; **0 platform edits**.
- **Achieved:** **29 / 29** landable pairs · academy **65 cards / 483 chapter links / 0 Draft chips** ·
  hero p95 **1.22 s** (employee) / **1.51 s** (manager), 5/5 ACCESS both · reproduced cold with **no
  intervention** · **0** platform-repo edits, verified per-clone.
- **Distance:** **gate met**, every component.
- **Status:** **`closed-on-gate`** — but the milestone **did not merge at this close**; see "Close blocker".

> ### ⚠ The denominator was CORRECTED 31 → 29 mid-milestone. The target SHRANK.
>
> **This is not 31/31 achieved, and must never be reported as such.** The gate opened against **31**
> landable pairs and closed against **29**, because at iter-07 the 2 skill-path **manager** pairs were shown
> to point at a surface next-web **has not built**: `InsightsBySkillPathStudentSimulationsContainer`
> hardcodes `userData = null`, has its results table **commented out**, and renders the literal string
> **"Coming soon"**. No query touches the seeded session, so the page is byte-identical whether or not
> anything was seeded. Under M233's fail-closed rule (*a session that cannot form a real link is dropped
> with a reason, never linked anyway*) those pairs are **not landable** — on exactly the ground that already
> excludes AI-labs.
>
> **31 was never a count of provable pairs; it assumed a surface that does not exist.** The correction is
> argued inline in `overview.md` with product-source evidence and the 31 struck through rather than
> rewritten, so the shrink stays auditable.
>
> **The correction also exposed a FALSE PASS.** The lighter of the two skill-path manager pairs had been
> scoring **green** off a definition-only "Results for" header — chrome served by a *different* query than
> the one that failed. So the pre-correction reading was not merely optimistic, it was **wrong in both
> directions at once**: counting 2 pairs that could never land, while recording one of them as landed.
> Numerator and denominator moved together (29 → 28 of 29) and the gap to gate went 2 → 1.
>
> Arithmetic chain, stated in full because `31` names two different quantities:
> `18 sessions + 15 manager views = 33 raw` → −2 skill-path manager (unimplemented) → `31 raw` →
> −2 ai-labs (presence-only) → **29 LANDABLE**.

### Iter ledger summary

- **Total iters:** 10 — **1 tok** (bootstrap, iter-01 "publish-then-prove") + **9 tiks**. No triggered tok.
- **Duration:** 2026-07-20 (single day, 10 iters).
- **Decisions accumulated:** 30 iter-level + 1 bootstrap tok strategy + 1 user-blocker (5 sub-findings,
  all resolved) + 4 close decisions.
- **Hardening:** final `--final` pass, **4 passes**, stabilized (coverage delta 0% across 3→4); 19 files in
  scope, 7 defects fixed inline, harness unit tests **0 → 72**, both load-bearing fixes mutation-verified.
- **Metric trajectory:** 0/31 → (reachability, +0 by design) → 16/31 → 27/31 → 29/31 → **28/29** (denominator
  corrected) → **29/29** → reproduced cold.

### Routes carried forward

Gate met, so no carry-forward.md. Four items routed **Fate 3 → the v2.5 release close**, each with a named
handler (see `decisions.md` CLOSE-D1 for the full audit disposition):

- `ACADEMY-M236-iter08-public-catalog-twin` — anonymous `/library` + `/free` render 0 cards. Needs a 2nd
  demopatch manifest + a next-web rebuild + a live re-prove. `frontend-tier.md`'s unqualified "the empty
  grid is FILLED" was **scoped to signed-in** at this close so the doc no longer over-claims.
- `apps/web` client GraphQL endpoint on the non-offset `:5050` — never manifested on a measured path
  (SSR uses a different var). Batched with the above: same rebuild, one re-prove.
- `DEF-M235-03` — the M204 assign-WRITE declared in-manifest TODO (inherited, routed past M236 by design).
- The live re-prove of the harness as modified by this close (**CLOSE-D3**).

**Discharged at this close:** M230 carry-forward **cluster 1** (rendered-card count — iters 08/10) and
**cluster 2** (the next-web clone re-anchor), the latter having had **no closing entry anywhere** until now;
its diagnosis was falsified by the milestone's own Phase-0b audit (the clone was never drifted).

### Dropped

- The **"demo reachable only over the tailnet"** gate clause — dropped by user decision B1, not by attrition.
  It is false by construction (`safety.md`: every demo container publishes on `0.0.0.0` on every bring-up),
  so it was only ever demonstrable via an off-tailnet probe. Reaching the right people is the VM + VPN's
  job, not the demo stack's.
- **Content-seat p95 latency** — explicitly out of scope for v2.5 (user decision B2). The content CTA emits
  no `data-login-as`, which *is* the ACCESS predicate, and `run-latency.sh` hard-rejects non-hero vantages.
  The 29 content actions are proven for **CONTENT**, not formally timed.

### Protocol evolution

- **`coverage-protocol.md`** — gained the content-stories sweep as the **second** sweep it governs, and had
  its `skipPaths` `/result/` exclusion **deliberately reversed** (the pages M236 existed to prove were the
  pages the rule excluded), amended in the same change per the protocol-evolution rule. Plus two new
  subsections: *the reading must be fail-CLOSED* and *prove the test fails (mutation, not coverage)*.
- **`playthroughs.md`** — gained the Playthrough-vs-content-story delineation: a Playthrough plays forward;
  a content story is already played, so there is nothing to play and no Playthrough.
- **`verification.md`** — gained **PRE-FLIGHT RUNG ZERO** (*"tagging is not publishing"*): a remote stack
  consumes tooling only at a tag **fetched from origin**, so work that exists only in the authoring copy is
  unreachable to it. iter-01 found `billion` unable to obtain the feature under test at all.
- **`latency-budget.md`** — the green-gate age check, `LATENCY_SCHEME=https` for `--public-host`, and the
  non-integer-`N` guard now propagated to all four runners.

### Close blocker

The deferral audit returned **RED**. The standing 14-failure demo-stack carry is a genuine repeat-deferral
(**10 milestones, 2 releases**) whose declared destination — the v2.4 release close — **already fired once
without landing it**, an AGED_OUT trigger no audit had recorded. M236 is the FINAL v2.5 milestone, so there
is no further milestone to defer into. **Per Phase 1b, the milestone does not merge until the user records an
explicit fate** (LAND-NOW / DROP / KEEP-DEFERRED-WITH-SIGNOFF). Full argument: `decisions.md` **CLOSE-D2**.
