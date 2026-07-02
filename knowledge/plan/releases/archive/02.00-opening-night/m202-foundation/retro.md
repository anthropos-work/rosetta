# M202 Retro — Playthroughs Foundation

**Closed:** 2026-07-01 · `closed-complete` · `section` · complexity medium.
**Outcome:** the `playthroughs` rext section stood up on the shared M42 e2e foundation — manifest model + light
validator, per-surface page-object layer (1 surface), a dedicated decoupled `pt-world` seed, the reset-to-seed
serial runner, the 4-state reporting map, and one trivial proof Playthrough (login → /profile → assert hero
identity) **GREEN on demo-1** — plus the `corpus/ops/demo/playthroughs.md` runbook that graduates the spec-draft
and IS the M203/M204 `iteration_protocol_ref`. Tooling + docs only, **zero platform edits, zero new deps**.

## Summary
6-section foundation delivered end-to-end. Mixed Go+TS toolchain (M202-D1): Go for the manifest/validator/report
(matching the seeding module's `datadna`/`stackseed` CLI family), TS for the Playwright page-object layer (matching
the M42 e2e foundation). The datadna gate runs as a subprocess (M202-D2), preserving the zero-cross-section-import
module boundary. rext authoring @ `b1e5528`, tagged `opening-night-m202`.

## Metrics delta (from `metrics.json`)
- **Go:** 96 test/fuzz funcs (92 Test + 4 Fuzz) across manifest/report/ptvalidate/ptreport — a new module (+96 vs
  M201's prose-only 0). Section coverage **98.5%** (report 100 / manifest 99.4 / ptvalidate 97.6 / ptreport 94.8).
- **TS:** 13 (12 stack-env unit + 1 proof Playthrough).
- **Flake:** 0 (Go 5/5 shuffled -race + TS unit 5/5 at close; harden was 3/3).
- **Close findings:** 8, all Fate-1 (2 should-fix code, 2 nice-to-have, 2 docs, 1 decision-blend); 3 new regression tests.

## Incidents this cycle
- **P3 — M202-D4 layering collision (surfaced + fixed at build §6).** Seeding `pt-world` onto an already-seeded
  demo-1 collided: the stories model forces the FIRST story onto the single-tenant Clerkenstein default-org slot,
  which on a seeded demo is the showcase's default org → duplicate-key on pre-existing `user_skills`. The
  zero-platform-edit, zero-fork fix was a leading **anchor story** (size 0) that re-declares the default org so the
  real pt-orgs get their own `StoryOrgID`s. A genuine seeding-machinery constraint (the default-org slot is
  single-tenant) — now blended into `stories-spec.md` (#M202-D4) so a future second-world seeder sees the landmine.
- No P0/P1/P2. Zero flakes. Zero regressions. Zero bugs surfaced across the whole harden pass (the production code
  was well-behaved on every error/edge/branch arm).

## What went well
- **The reuse-not-fork discipline held.** `hero-login.ts` imports (never forks) the M37 `cockpit-login` handshake;
  the dedicated seed consumes `stackseed`/`datadna` unchanged (M202-D3). The section adds a 6th rext module without
  touching any sibling.
- **Harden found nothing to fix but the fuzz corpus proved the parsers total** — FuzzParsePlaywrightJSON /
  ExtractPTTag / Load / ValidateNeverPanics confirm the crash-free property (the value is the property, not a line bump).
- **The doc-drift guard did its job at close.** The CQ2 refactor (inline ternary → validated `resolveWorkers()`)
  tripped `runner_safety_test.go`'s serial-default assertion — exactly the drift the guard exists to catch; updated
  to assert the stronger validated invariant.

## What was hard / what to carry
- **CQ1 was a real diagnosis bug hiding in a happy-path.** `ptvalidate` collapsed datadna's exit-3 (usage/config
  error) into "gate FAILED" — an operator misconfiguration would have mis-pointed the diagnosis at the seed. Caught
  only by the cross-cutting close review, not the per-section build. Lesson: subprocess exit-code semantics deserve
  a dedicated arm, not an all-non-zero catch-all.
- **The per-unit index half of the handbook contract is easy to miss.** The section README existed but the section
  was absent from both rext section indexes (DOC1). Carry: the index row is part of "ship a new unit", not an afterthought.

## Carried forward
- **Nothing deferred** (deferral audit GREEN — 0 milestone-owned). The `Out:` scope (real product coverage, AI-sim
  mirror tier, cross-vantage) is original design-scope, owned Fate-2 by **M203 ∥ M204** (vantage coverage) + M206
  (roadmap-vision). Both import this page-object layer + run on the reset-to-seed lifecycle per `playthroughs.md`.
- **Administrative KEEPs:** the `opening-night-m202` tag is local-unpushed (the user's push gate); no consumption
  re-pin needed (the foundation runs from the authoring copy against demo-1).
