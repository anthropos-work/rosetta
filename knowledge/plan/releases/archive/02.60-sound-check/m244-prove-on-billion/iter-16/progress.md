**Type:** tik (run 6, tik 4). Active strategy: TOK-02.

# iter-16 — progress

## The fix (harness-only, like iter-13 — no billion build input)
`ANT_ACADEMY_HOME_SECTION` recalibrated from `{ kind: 'text', mustInclude: ['AI Academy'], minMeaningfulLen: 400 }` to `{ kind: 'both', mustInclude: ['Academy'], minMeaningfulLen: 400, itemSelector: 'a[href*="/chapters/"]' }` + `minCount: 12`. Rationale: the literal "AI Academy" is only in the non-visible `<title>` (the visible heading is "Academy"); the browser gate reads what a human sees, so the token false-FAILED a home rendering 483 cards. The card-count is the M42e cardinality the coverage protocol calls for — it STRENGTHENS the teeth (an empty catalog OR a keyless blank both render 0 cards → both FAIL, which the old "AI Academy" token could NOT catch: the pre-fix empty `/library` header read "The AI Academy library" with 0 cards). `coverage-manifest.unit.spec.ts` updated to pin the `both`-kind + the ≥10 card floor + the dropped stale literal.

## Verification
- `tsc --noEmit` clean; `coverage-manifest.unit.spec.ts` **54/54** green.
- Live on billion (both hero vantages, `COVERAGE_HOST` + https):
  - **employee (maya-thriving): GATE MET ✅** — reachable 62/150, failingSections 0 / personaFailures 0 / escapes 0 / notReached 0 / crossPortFailures 0 / frontier EXHAUSTED (4.0m).
  - **manager (dan-manager): GATE MET ✅** — reachable 70/150, all 0 / crossPort 0 / frontier EXHAUSTED (5.7m).
- **⇒ the stack-verify COVERAGE half of gate (c) is GREEN on billion.**

rext `2a71e08`, consumption tag moved (peels `^{}` → 2a71e08 on origin). Billion's rext clone stays at 8391843 — the coverage harness runs from the local authoring copy driving billion's unmodified app, so it is not a billion build input (no re-pin).

## Gate (c) status
The coverage sweep (the core stack-verify live-browser proof) is GREEN both vantages. Gate (c) as a whole ("40 live-browser specs = 24 stack-verify + 16 Playthroughs") is NOT yet ticked: the DISCRETE stack-verify specs (calibrate/persona/verify/m220/m224/talk-to-data/enterprise-surfaces/smoke/…) + the 16 Playthroughs remain. Playthroughs run LAST (pt-world reset destroys the demo seed — iter-14 D1). So the primary metric (gate parts) stays 5/8.

## Close — 2026-07-23

**Outcome:** The stale academy-home coverage marker recalibrated to card-count; both hero coverage sweeps now GATE MET ✅ on billion — the coverage half of gate (c) proven green. Gate (c) not yet ticked (discrete specs + Playthroughs remain). Metric stays **5/8**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (gate part (c) — coverage half green; discrete specs + Playthroughs remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 4/5) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1 (iter-16/decisions.md).
**Side-deliverables:** none.
**Routes carried forward:**
- gate (c) REMAINING: the discrete stack-verify specs (calibrate/persona/verify/m220/m224/talk-to-data/enterprise-surfaces/smoke/…) + the 16 Playthroughs LAST (pt-world reset). Then gate (c) ticks.
- gate (f) 3 drift-carries · (h) v2.6 fixes + p95<5s · DEF-M239-01 — demo-seed gates, run before the Playthroughs reset.
**Lessons:** a coverage gate's TEXT markers rot when the app's copy is redesigned (M211's "AI Academy" → "Academy") — a CARDINALITY marker (does the grid render N cards?) is both more durable (immune to copy) and stronger (catches the empty-catalog false-pass the token missed). Prefer counting real items over matching a heading string.
