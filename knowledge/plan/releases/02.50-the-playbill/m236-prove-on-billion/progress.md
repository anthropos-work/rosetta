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

## Next-iter queue

**Gate status — 3 of 4 components MET:**
- primary metric **29/29** landable (session × action) pairs rendering real content live on `billion` ✅
- academy grid renders real cards — **65 cards, 0 Draft chips** (Thread A) ✅
- **p95 click→ACCESS** hero vantages — employee **3.15 s**, manager **2.71 s** vs a 5 s gate ✅
- 0 platform-repo edits ✅
- **cold reset-to-seed reproduction — OUTSTANDING** ⬅ the only remaining component

**Next iter:**
- **The cold reset-to-seed reproduction** (handler `REPRO-M236-iterTBD-cold-cycle`). The gate requires
  reproducibility, and the current stack has had its manifest re-exported and its cockpit restarted in
  place across iters 07–08. Re-pin `billion` to the final tag, tear down, cold `/demo-up` with public-host
  default-on, then re-run the content-stories sweep + a latency battery against the fresh stack.
  Success = stack UP, fresh-green `autoverify.json`, cockpit serving `content-manifest.json` with all 4
  products, sweep 29/29, hero p95 < 5 s. **Expect 30–45 min** for the bring-up itself.
  *Note:* the cockpit currently binds `127.0.0.1` because `tailscale serve` had claimed `:17700` by
  restart time (iter-07); a cold bring-up restores the normal `0.0.0.0` ordering.

**Carried to release close (not gate-blocking):**
- **M230 carry-forward cluster 3** — the anonymous `/library` + `/free` academy routes render 0 cards
  (`getPublicCatalogView` takes the `new Set()` branch the M230 patch does not cover; the patch manifest
  names the gap itself). Handler `ACADEMY-M236-iter08-public-catalog-twin` — a twin manifest of the same
  FS-published transform. M230 cluster 1 is **discharged** (iter-08).
- **`apps/web` client GraphQL endpoint** — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql`,
  the **non-offset** port, while `apps/hiring` carries the correct offset origin. Did not manifest on any
  measured path (SSR uses `WUNDERGRAPH_SSR_ENDPOINT`); a demo-hygiene item. From iter-08 D5.
- **Standing carry:** 14 pre-existing demo-stack test failures (REPEAT) → v2.5 release close.

**Resolved:**
- **USER-BLOCKER-M236-01: RESOLVED** 2026-07-20 (B1–B5). Gate re-scoped in `overview.md`; Phase 0b
  re-verdicted RED → discharged (`spec-notes.md`).
- **`ACADEMY-M236-iterTBD-catalog-fill`** — closed by iter-08 (the CTA named a nonexistent route; the
  catalog was never the problem, and `academy-seed` is moot in a demo).
- **`SKILLPATH-M236-iter07-manager-hang`** — closed by iter-07 (a `networkidle` wait, over an
  unimplemented platform surface).
- **`LATENCY-M236-iterTBD-hero-p95`** — closed by iter-09.
