# M236 — Progress

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 "publish-then-prove" authored; baseline measured live — gate denominator is **31** landable (session × action) pairs, currently **0/31**, blocked by an unpublished-tooling gap (`billion` pins the M228 tag; 0 of 13 `playbill-*` tags are on origin) — see iter-01/progress.md

- iter-02 (tik): **closed-fixed** — Phase P complete. Published `rosetta-extensions` `main` (`1d97861..60eff14`) + all 13 `playbill-*` tags to origin; re-pinned `billion` → `playbill-m235-hardened`; M217 pin guard PASSES; **31 denominator confirmed against the shipped manifest**. Metric 0/31 unchanged by design (bought reachability, not numerator). Side-deliverable: 109 GB host cache reclaimed — see iter-02/progress.md

- iter-03 (tik): **closed-fixed** — Phase C complete. Stale M228-era demo-1 purged; cold rebuild at `playbill-m235-hardened` UP with **fresh-green autoverify**; cockpit serves `/content-manifest.json` HTTP 200 with 4 products / 18 sessions; substrate verified 13/13 sim sessions + attempt-results + manager mirror, ai-labs presence 2/2. Academy tables empty (in-scope, routed). Metric held honestly at 0/31 — no render proven yet. Blocker found + cleared: remote bring-up needs a LOGIN shell (Go on PATH only via profile) — see iter-03/progress.md

## Next-iter queue

- ~~iter-03 — Phase C~~ **DONE (closed-fixed)**. Superseded by: **iter-04 — Phase H: author the content-stories seat-login harness** against the now-live seeded render (handler `HARNESS-M236-iter04-seat-login`) — unblocked for the first time. Original entry: Phase C cold bring-up on `billion` at `playbill-m235-hardened`, public-host default-on. Success = stack UP, fresh-green `autoverify.json`, cockpit serving a non-404 `content-manifest.json` with all 4 products. Handler `PHASE-C-M236-iter03-cold-bringup`. **Expect 30–45 min** — the host build cache was pruned, so this is a genuinely cold UI-tier rebuild.
- **Phase L (iters 4+):** land the arms in descending evidence density — simulation (26) → skill-path-legacy (4) → skill-path-new/academy (1) → ai-labs (0 landable, prove 2 presence rows).
- **Phase H (interleaved, after the first live render):** author the content-stories seat-login coverage harness (13 seats → 13 manifests, page-object from scratch) + the B4/B5 doc work — `coverage-protocol.md` `skipPaths` `/result/` reversal WITH its amendment, and the content-stories backfill of `coverage-protocol.md` + `playthroughs.md`. Handler `DOC-M236-iterTBD-protocol-backfill`.
- **USER-BLOCKER-M236-01: RESOLVED** 2026-07-20 (B1–B5). Gate re-scoped in `overview.md`; Phase 0b re-verdicted RED → discharged (`spec-notes.md`).
