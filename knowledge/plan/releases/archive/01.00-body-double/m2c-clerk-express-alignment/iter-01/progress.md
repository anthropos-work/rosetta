# M2c / iter-01 — progress

**Type:** tok (bootstrap) — authors the first strategy (TOK-01). Protocol:
`corpus/architecture/alignment_testing.md` (the alignment loop: author-DNA → capture-goldens →
`alignctl run` → drive-score-to-gate).

## Close — 2026-06-03
**Outcome:** authored **TOK-01** (RS256-native, additive-first, real-`@clerk/express` Node runner); resolved
the 3 bootstrap unknowns (offline-available `1.7.79`, separate token domain → additive viable, Node runner).
**Type:** tok (bootstrap)
**Status:** closed-no-lift  (toks don't move the gate; the bootstrap authors strategy)
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (BOOTSTRAP, not triggered) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** TOK-01 (milestone `decisions.md`); iter-01-D1 (the 3 bootstrap findings, this iter's `decisions.md`)
**Routes carried forward:** iter-02 → author `clerk-express-1.json` + validate.
**Lessons:** `@clerk/express` is offline-available → verify against the **genuine** library (no Go fallback) — the svix discipline applies. A separate Clerk instance for studio-desk means additive RS256 should hold; re-survey confirms per-tik.
