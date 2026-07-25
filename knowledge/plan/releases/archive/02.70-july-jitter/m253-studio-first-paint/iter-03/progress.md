**Type:** tik (cleanup-shape — docs Delivers, under TOK-01)

# M253 iter-03 — progress

Landed the milestone's three documentation Delivers:

1. **`corpus/ops/demo/latency-budget.md`** — new section "## The studio-desk first-paint budget (v2.7 M253)":
   the < 1000 ms FMP gate, the per-leg baseline (clerk.load ~140 ms / l12n ~12 ms / **canAccess ~3.9 s** →
   skeleton at ~4669 ms), the two-demopatch fix, the 4669 → 817 ms result, and the `run-studio-fcp.sh` harness.
2. **`corpus/ops/demo/demopatch-spec.md`** — §5 inventory: 2 new rows (`studio-desk-shell-first-paint` +
   `studio-desk-no-thirdparty`), the count reconciled 21 → 23 (studio-desk 3 → 5), and an M253 intro note.
3. **`corpus/services/studio-desk.md`** — new section "### The MPA / empty-body boot model — and the demo
   first-paint reorder (v2.7 M253)": why the blank happens (shell built behind the awaits) + the reorder fix +
   the NOT-a-dev-build / NOT-code-splitting refutations.

Cross-refs wired both ways (latency-budget ↔ demopatch-spec ↔ studio-desk).

## Close — 2026-07-24

**Outcome:** the three docs Delivers landed; the fix is now discoverable from the budget, the patch spec, and the
service doc. No production code; the FCP metric was already met in iter-02.
**Type:** tik (cleanup-shape)
**Status:** closed-fixed
**Gate:** MET (carried from iter-02: skeleton-visible p95 817 ms < 1000 ms, demo-2 LOCAL LAPTOP; the fresh-green COLD confirmation is chartered to M254 per coord rule 9)
**Phase 5 grading:** (1) gate-met: y (M253's local-bootstrap charter — fix + runner + docs all landed; FCP p95 817 ms met on demo-2; the fresh-green COLD-p95 confirmation is M254's chartered deliverable) — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1 (gate-met, with the explicit M254 cold-green caveat)
**Decisions:** none new (docs only).
**Side-deliverables:** none.
**Routes carried forward:** M254 — the fully-green COLD-p95 confirmation on billion (re-measure the studio FCP
gate on a freshly brought-up, fully-set-dressed cold demo with a green `autoverify.json`).
**Lessons:** documenting the boot model + the per-leg baseline in `latency-budget.md` makes the "why the blank"
answer discoverable in the same place the login budget lives — the corpus's single perf home.
