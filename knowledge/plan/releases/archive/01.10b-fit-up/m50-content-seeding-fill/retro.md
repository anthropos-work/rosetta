# M50 — Retro (Content & seeding fill)

## Summary
Iterative milestone, **closed-on-gate** 2026-06-30. 6 iters (1 bootstrap tok + 5 tiks) drove the M42 semantic
coverage gate to **MET on BOTH vantages** (employee iter-02 + manager iter-06) on a WARM demo-1, on the manifest
**STRENGTHENED to PROVE the M50 fills** (frontier-exhausted, (failingSections,escapes)=(0,0), 0 persona, 0
cross-port). The sweep-first strategy (TOK-01) paid off: the employee vantage needed **no** seed fix (already
green on the fresh re-grounded demo), and the genuine gaps were the manager-side ones — member fields, spoken
languages (a NEW seeder), org-wide cert coverage, and two prod-eject escape classes. Final-harden Pass 2
(stabilized) pinned the cross-seeder shared-uuid-space invariant before close.

## Incidents This Cycle
- **P2 — iter-02 manager-baseline misread (corrected iter-03, no code impact).** iter-02 read the manager sweep
  as frontier-capped and routed a tooling-iter (raise/sample the BFS cap). iter-03's re-survey found the sweep
  actually EXHAUSTS (cap=300 → reachable=68, `cappedAtFrontier=false`, gate-VALID) — the tooling-iter was
  CANCELLED. Cost: one iter of mis-direction; caught by the protocol's "raise the cap until exhausted before
  quoting" discipline. No code shipped on the wrong assumption.
- **P2 — the run-1 gate passed BLIND (the D4/F1 reconciliation).** iter-05 reached `(0,0)` both vantages, but the
  manager manifest never ASSERTED languages/certs/member-fields — so a green gate co-existed with two genuinely
  empty M50-own surfaces (languages 0 rows DB-wide, certs hero-only 9/340). iter-06 FILLED them AND strengthened
  the manifest to assert them. Lesson pinned (below). No regression — the manifest simply under-asserted.
- **P3 — `run-coverage.sh` arg-forwarding footgun (side-fixed iter-06).** Consumed positional args leaked to
  Playwright as filename filters, so a new `calibrate-manager.spec.ts` matched "manager" and co-ran with the
  sweep, corrupting a run. Fixed with a guarded shift-consume loop.
- **P3 — gofmt slip (close-caught).** Two M50 files (`member_languages.go`, `users.go`) shipped not gofmt-clean
  (trailing-comment misalignment). `gofmt -w` at close; a CI fmt gate would have caught it.
- **P3 — a broken M51 backref (close-caught).** The Fate-3 academy annotation's `secrets-spec.md` backref was
  one `../` too shallow (resolved inside `knowledge/`). Fixed by the Phase 8 cross-ref check.

## What Went Well
- **Sweep-first (diagnose-before-fix) was correct.** The fresh re-grounded demo rendered the employee vantage
  green with zero seed work — exactly the "several empties may be demo-up-#7 artifacts that vanish once set-dress
  runs" hypothesis the iterative shape was chosen for. No speculative seeders built.
- **The two escape classes were correctly separated.** iter-04 proved the JS-constant escape (demopatch) and
  the replayed-content escape (Directus content-URL rewrite) are DIFFERENT classes — avoiding the trap of
  declaring victory when the demopatch left escapes at 1. The new escape class is now a routing-table row.
- **Zero-platform-edit held throughout.** The `MemberLanguagesSeeder` writes only `user_languages` and leans on
  the platform's existing AFTER-INSERT trigger; the content-URL rewrite is demo-local DDL on the per-stack
  Directus; the demopatch trap-reverts. No canonical repo touched.
- **The harden Pass 2 cross-iter sweep earned its place** — it pinned the shared per-member uuid-space invariant
  (every member-facing seeder derives the same uuid for member `i`), the load-bearing believability spine that
  per-seeder tests can't catch.

## What Didn't
- One iter (iter-02→03) was spent on a manager-baseline misread before the correct "it exhausts" reading. The
  protocol's measurement discipline (exhaust-before-quote) is the guard; it fired one iter late.
- The run-1 gate's blind-pass shows the gate is only as honest as its manifest asserts — a structural risk in
  presence-only coverage that M50 had to reconcile mid-flight rather than having been asserted up front.

## Carried Forward
- **COLD reset-to-seed acceptance → M53** (Fate-2, user-decided, D-CLOSE-2). All M50 seeders + fixes reproduce
  from the bring-up tooling on a fresh `/demo-up`; M53 owns the single from-cold acceptance truth.
- **Academy content + menu-link/non-anonymous-session (F6) → M51** (Fate-3, D-CLOSE-3). Annotated to M51's
  candidate scope; the academy AI chat stays documented-as-absent per the keys policy.
- **Consumption-clone re-pin to `fit-up-m50` → push-gated KEEP** (release-level; authoritative at M53).

## Lessons (pinned)
1. **A green `(0,0)` coverage gate is only as honest as its manifest ASSERTS.** Cross-check the gate's assertions
   against the milestone's INTENT, not just the (0,0) — the run-1 blind-pass → run-2 strengthened-pass is the
   canonical example. (Recorded in iter-06 + the coverage-protocol routing context.)
2. **Tab-gated / paginated content needs new harness primitives** (`preAssert` tab-click + `textMatch` OR-assert)
   — keep them additive (new branches gated on new fields) so other vantages stay provably unaffected.
3. **Two prod-eject escape classes, two fix surfaces:** a JS-constant baked link → the demopatch; a prod host in
   a REPLAYED CONTENT field → a post-replay content-URL rewrite. Diagnose by whether the JS-constant ejects are
   already 0. (Now a coverage-protocol routing-table row.)

## Metrics Delta (from metrics.json)
- **rext stack-seeding Go tests:** 706 → **719** (+13: M50's 6 iters' test files + 6 harden/iter files + the +1
  close test `TestNativeLanguageByCity_CoversLocations`). Seeders pkg 349, **97.4%** stmt (stable from harden Pass 2).
- **demo-stack Python:** `test_demopatch` 47→49 · `test_frontend_build` 55→59 (the content-URL-rewrite contract
  tests); suite total 299.
- **Flake:** **0** (5/5 shuffled-sequential seeders pkg; Python 108 OK).
- **Gate:** M42 coverage MET both vantages (warm, strengthened manifest). COLD → M53.
- **Supply-chain:** 0 new deps this milestone.
