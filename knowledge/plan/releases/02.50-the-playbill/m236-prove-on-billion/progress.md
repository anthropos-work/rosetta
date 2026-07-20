# M236 — Progress

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 "publish-then-prove" authored; baseline measured live — gate denominator is **31** landable (session × action) pairs, currently **0/31**, blocked by an unpublished-tooling gap (`billion` pins the M228 tag; 0 of 13 `playbill-*` tags are on origin) — see iter-01/progress.md

- iter-02 (tik): **closed-fixed** — Phase P complete. Published `rosetta-extensions` `main` (`1d97861..60eff14`) + all 13 `playbill-*` tags to origin; re-pinned `billion` → `playbill-m235-hardened`; M217 pin guard PASSES; **31 denominator confirmed against the shipped manifest**. Metric 0/31 unchanged by design (bought reachability, not numerator). Side-deliverable: 109 GB host cache reclaimed — see iter-02/progress.md

- iter-03 (tik): **closed-fixed** — Phase C complete. Stale M228-era demo-1 purged; cold rebuild at `playbill-m235-hardened` UP with **fresh-green autoverify**; cockpit serves `/content-manifest.json` HTTP 200 with 4 products / 18 sessions; substrate verified 13/13 sim sessions + attempt-results + manager mirror, ai-labs presence 2/2. Academy tables empty (in-scope, routed). Metric held honestly at 0/31 — no render proven yet. Blocker found + cleared: remote bring-up needs a LOGIN shell (Go on PATH only via profile) — see iter-03/progress.md

- iter-04 (tik): **closed-fixed** — Phase H. Calibrated the result render against the LIVE stack (3 shapes found, 2 of which a blind author would have gotten wrong in opposite directions), then authored the content-stories sweep (`lib/content-result-page.ts` + `tests/content-stories.spec.ts` + runner + aggregator, rext `playbill-m236-content-sweep`). **Metric 0/31 → 16/31** — all 13 simulation PLAYER pairs land. Fixed 2 harness bugs that produced confident wrong numbers (serial-abort; worker-restart amnesia) and 1 **false PASS** (skill-path graded as a scored sim). B4 + B5 landed: `coverage-protocol.md` premise corrected + content-stories section; `playthroughs.md` backfilled — see iter-04/progress.md

- iter-05 (tik): **closed-fixed** — the MANAGER vantage. Root-caused to ONE fact: the activity-dashboard route's last segment is a **MEMBERSHIP id**, not a user id (`GetMembership(membershipsID)` → `ent: membership not found` → the whole query nulls). The page's header comes from a *different* query, so it looked populated while proving nothing. Fixed in tooling (deterministic membership id, zero platform edits), test corrected (it had asserted the defective contract), honesty gate caught the stale preset. **16/31 → 27/31** — see iter-05/progress.md

## Next-iter queue

- **iter-06 — the residual 4** (27/31 → 31/31):
  - **2 INTERVIEW manager pairs** — fail with *no header at all* (a different mode from the 11 just fixed), on `/activity-dashboard/interviews/…`. Likely a distinct route/tab; cf. M231's `flag_interview_manager_report` demopatch. Handler `INTERVIEW-M236-iter06-manager-report`.
  - **1 skill-path manager route** times out (persisted through the membership fix, so independent). Handler `SKILLPATH-M236-iter06-manager-timeout`.
  - **1 academy pair** — `/library/<slug>` not-found; academy tables empty. Needs `app/cmd/academy-seed` wired into the cold bring-up. Handler `ACADEMY-M236-iterTBD-catalog-fill`.
- **Then:** the p95 click→ACCESS measurement for the HERO vantages (gate component, `run-latency.sh`), and a final cold reset-to-seed reproduction. at `playbill-m235-hardened`, public-host default-on. Success = stack UP, fresh-green `autoverify.json`, cockpit serving a non-404 `content-manifest.json` with all 4 products. Handler `PHASE-C-M236-iter03-cold-bringup`. **Expect 30–45 min** — the host build cache was pruned, so this is a genuinely cold UI-tier rebuild.
- **Phase L (iters 4+):** land the arms in descending evidence density — simulation (26) → skill-path-legacy (4) → skill-path-new/academy (1) → ai-labs (0 landable, prove 2 presence rows).
- **Phase H (interleaved, after the first live render):** author the content-stories seat-login coverage harness (13 seats → 13 manifests, page-object from scratch) + the B4/B5 doc work — `coverage-protocol.md` `skipPaths` `/result/` reversal WITH its amendment, and the content-stories backfill of `coverage-protocol.md` + `playthroughs.md`. Handler `DOC-M236-iterTBD-protocol-backfill`.
- **USER-BLOCKER-M236-01: RESOLVED** 2026-07-20 (B1–B5). Gate re-scoped in `overview.md`; Phase 0b re-verdicted RED → discharged (`spec-notes.md`).
