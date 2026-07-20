# M231 "content-stories feasibility spike" — Retro

## Summary
The hard go/no-go barrier for Thread B — **GO**. Discovered the per-product player+manager result-route map, proved-by-
render (against billion) which land from seedable rows, confirmed the prod-session sourcing+anonymization mechanism (via
the read-only postgres MCP), catalogued public sims by modality, and ruled AI-labs OUT / academy IN. Deliverable:
`corpus/ops/demo/content-stories-routes.md`.

## Incidents this cycle
- **A stale corpus claim caught + corrected:** the M228 "recruiter drawer is a Next.js intercepting route" explanation
  was a MISDIAGNOSIS (it's a plain AntD Drawer). The RENDER_ONLY_SIM fix still worked, but the documented *why* was
  wrong; fixed in hiring.md (KB-3). Two more stale facts (skillpath manager mirror, ant-academy backend read/WRITE)
  fixed inline (KB-4/7). Lesson: a spike that prove-by-renders naturally re-audits the docs it reads.

## What went well
- **The central release risk resolved cleanly + early:** the sim result page reads a persisted row (plain SELECTs), so
  a cloned session renders — no demo-patch/escalation forced for the main path. This is the finding that makes Thread B
  buildable as designed.
- **Prove-by-render + read-only prod scouting** measured the RIGHT things (real routes, real modality counts, the mirror
  trap) rather than assuming — exactly what a go/no-go spike should do.
- The Fate-3 routings tightened M232/M234's scope with code-cited specifics (flag-gate, public-anchored sourcing,
  presence-only AI-labs, academy real-progress) — the downstream milestones inherit a sharp spec.

## What didn't
- Two content products degrade from the ideal: interview needs a demo-patch (flag-gate), AI-labs can't do a played
  result at all (presence-only). Honest scoping, but the tab won't be uniform across all four products.

## Carried forward
- D3 interview flags → M232 · D4 AI-labs presence-only → M234 · D5 academy real-progress → M234 (all user-accepted).
  KB-2/5/6/8 → an architecture-doc pass. See decisions.md + content-stories-routes.md.

## Metrics delta
- Deliverable 349 lines / 23 code-citations · modality catalog 77 voice / 65 code / 30 document · 0 platform edits ·
  KB-fidelity YELLOW. See metrics.json.
