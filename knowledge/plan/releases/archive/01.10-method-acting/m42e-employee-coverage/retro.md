# M42e Retro — Employee 100% demo coverage

## Summary
M42e closed the FIRST `iterative` milestone of v1.10 **on-gate**: logged in as Maya (DevOps Engineer @
Cervato Systems) on a **fresh zero-manual demo-up**, the new Playwright semantic-coverage sweep reports
`gateMet:true` — 62 reachable pages, **0 failing sections, 0 persona failures, 0 prod-eject escapes,
frontier EXHAUSTED** — proven authoritatively from the consumed clone @ tag (committed code), reproduced
across 2 `demo-down --purge`+`demo-up` cycles. The milestone delivered BOTH the iteration protocol
(`corpus/ops/demo/coverage-protocol.md`) AND the Playwright semantic-coverage harness (rext
`stack-verify/e2e/` — the first non-Go rext dev/test dependency), then iterated the design-plan's 7 root
causes (P0–P8) into the rext seeders / snapshot surfaces / demo injection until the employee believability
bar was met and reproducible. **Zero platform-repo edits.** Code-of-record @ tag `method-acting-m42e` →
`53574ae`; the rosetta side carries the doc-half + the 23-iter plan archive. Closed in a near-clean
review: 8 findings, 0 blocking.

## Incidents This Cycle
- **Mid-milestone gate re-scope (2026-06-25, user live-review) — P1, not a defect.** The original gate
  measured DOM text-density (`textLen>40`), which green-lit pages rendering placeholder/empty-state cards +
  nav chrome — the harness reported a green `(0,0)` while a logged-in presenter saw a placeholder `/profile`,
  an empty library, incoherent 3D-dental skills for a backend dev, a silhouette avatar, and no org logo. The
  user re-scoped the gate to the **semantic believability bar** mid-milestone (commit `0eaab39`). Handled by
  the iteration protocol (a TOK re-scope at iter-10 ratified the new strategy); the lesson — a coverage gate
  must assert *semantic content + cardinality + persona consistency*, not text length — is folded into
  `coverage-protocol.md`. The iters 1–9 work (the harness skeleton, the crawl/wait-strategy fixes, the
  entitlement seeding) was NOT wasted: it carried forward to the rebuilt harness.
- **iter-16 user-blocker (avatar licensing) — P2, resolved in-milestone.** No image source is BOTH
  copyright-clean AND consent-clean for depicting a real identifiable person as a fictional employee. Surfaced
  to the user (per the hard avatar rule); the user chose option (a) **synthetic non-existent-person photoreal
  faces** (StyleGAN2 / "this-person-does-not-exist"-class). Landed iter-18 (`photo_avatar` generator);
  menu==profile real-photo persona assert PASSES.
- **2 reproducibility bugs found+fixed at the P8 fresh-demo-up acceptance — P2, the gate's whole point.**
  (1) iter-22: a STALE `stacksnap` binary surviving `demo-down --purge` in `$STACK/bin` (the `build_cli`
  skip-if-present) silently skipped the sim-embeddings replay on a fresh demo-up → `cms.similarities=0` →
  empty library. Fix: `build_cli` ALWAYS rebuilds from the consumed source. (2) iter-23: an avatar-consistency
  FALSE-fail — the harness `readMenuAvatarSrc` grabbed the org-switcher monogram SVG (`alt="company logo"`)
  instead of the user's real-photo data-URI. Fix: exclude the org img by `alt` + prefer a raster candidate.
  Both were exactly the class the fresh-demo-up gate exists to catch (a live-patched stack would have hidden
  the stale binary; a weaker harness would have mis-asserted the avatar).
- 0 flakes (5/5 shuffled `-race` at close); 0 regressions (supply-chain GREEN, go.mod/go.sum byte-identical;
  all 5 alignment gates 100%). The close's 3 Phase-2c adversarial scenarios surfaced 0 behavioural defects
  (the persona-assert precedence-trap was a future-edit hazard fixed defensively; the no-non-live-unit-test is
  the harness architecture; the curated-category precedence was already guarded).

## What Went Well
- **The fresh-demo-up reproducibility bar paid off.** Making "reproduces on a fresh zero-manual demo-up" part
  of the gate (not "works on the live-patched stack") caught two real bugs that live-patching would have
  masked. The bar is the milestone's most valuable design choice.
- **Root-cause-first over sweep-first (after the re-scope).** The TOK-10 design-plan had already triaged the 7
  believability roots on the live demo, so iters 10–23 LANDED fixes reproducibly rather than re-discovering
  them — each root in its own seeder (Go + tests), measured fast on demo-3, accepted on the fresh demo-up.
- **Honest residuals.** Each iter reported the TRUE state, never a cap-saturated floor as a win (iter-07
  corrected a `(2,3)` floor to a frontier-EXHAUSTED `(3,3)`; iter-09 corrected iter-08's dishonest sim-start
  skip). The gate fired only when it was genuinely met.
- **Zero platform edits held** across a milestone that touched skills, profile, activity, avatars, org logo,
  library, Sentinel reload, and the crawl — all via rext seeders / snapshot grants / injection.

## What Didn't
- **The original gate was too weak and shipped iters 1–9 against it** before the re-scope. A semantic bar from
  the start would have saved the text-density detour. Lesson captured in `coverage-protocol.md`.
- **The Playwright DOM-selector logic is the harness's bug-prone surface** (the iter-23 false-fail lived
  there) and has no non-live unit test — inherent to `page.evaluate` browser-context code under the
  Playwright-only supply-chain line. Mitigated by live calibration probes + the authoritative gate; the
  selector was hardened defensively at close.

## Carried Forward
- **The MANAGER vantage → M42m** (the sibling milestone, already scaffolded; NOT a repeat-defer — M42e never
  owned the manager vantage). The M42e manager smoke-sweep calibrated its input: **139 `studio.anthropos.work`
  escapes** (one baked left-nav prod link → link-rewriting, Fate-3 → M42m); **5 unreached `/workforce/*` M36
  dashboard pages** (the core manager content+nav work, Fate-2 → M42m); the **team-roster `/user/<id>` fan-out**
  (a representative-sample crawl rule + higher cap, Fate-3 → M42m). M42m's manager-manifest descriptors are
  authored (`calibrated:false` until the workforce surfaces render).
- **DEF-M40-01 (KPI "AI simulations completed" = 0)** — employee half RESOLVED in M42e (the `/profile`
  stats-row section passes); manager half (the Workforce-dashboard KPIs) Fate-2 → M42m.

## Metrics Delta
(from `metrics.json`) Go test funcs **1326 → 1373 (+47)**: stack-seeding 496→534 (+38), stack-snapshot
354→361 (+7), clerkenstein 264→266 (+2) — PLUS the NEW TypeScript Playwright harness (7 tests / 6 spec
files, the first non-Go test surface). Coverage: stack-seeding/seeders **97.0%** stmts. Flake: **0** (5/5
shuffled `-race`). Supply-chain: **GREEN** (0 new Go deps; go.mod/go.sum byte-identical; @playwright/test
^1.49.0 the only non-Go dep, pinned). Alignment gates: **5/5 at 100%**. Iters: **23** (1 bootstrap tok +
1 re-scope tok + 21 tiks; 0 triggered toks). Close status: **closed-on-gate**.
