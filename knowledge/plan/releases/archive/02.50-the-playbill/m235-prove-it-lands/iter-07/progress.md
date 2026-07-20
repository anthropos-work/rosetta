# iter-07 — progress

**Type:** tik (under TOK-01, Track A step 2 — non-simulation product sections)

## What landed
The **academy (skill-path-new)** content-story section (`seeders/content_nonsim.go`) — the 3rd + most
live-coupled non-sim section:
- An `AcademySlug` field + 1 academy exhibit (a REAL public course "Foundation of Artificial Intelligence",
  slug from the captured snapshot).
- The academy arm in `resolveNonSimSession` — app_base=academy, a `/library/<slug>` player route (a real,
  resolvable course), NO manager view. The cockpit renders the DIRECT academy-origin CTA (e2e_persona seam).
- The seeder academy arm is a documented no-op — the `academy_chapter_progress` write is the `academy-seed`
  PLATFORM binary (live, needs M230's catalog), routed to M236.
- Regenerated `presets/content-manifest.json` (18 sessions; all 4 products resolve); honesty gate GREEN.

## Proof (unit — the academy progress-write + exact chapter route + M230 catalog are M236)
- Go: full `go test ./...` GREEN. New: `TestBuildNonSimProducts_AcademyDirectOriginNoManager` (app_base
  academy, /library/<slug>, no manager, label+seat).
- Python: 47 content/academy render tests GREEN incl. a new academy render test (course title, direct-origin
  CTA + e2e_persona seam, NOT a FAPI handshake, no manager CTA). Full test_cockpit.py: 122 passed, the SAME 6
  pre-existing failures (Academy/OverlayJs) — no new failures.
- rext tag: `playbill-m235-nonsim-academy`.

## Close — 2026-07-20

**Outcome:** the academy section RESOLVES (app_base academy, real course CTA, no manager), unit-proven.
Readiness: **3 of 3** non-simulation sections — ALL offline-buildable non-sim sections land.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the academy progress-write + the progress-bearing route + the M230 catalog fill are LIVE →
M236; iter-07 unit-proves the section shell + a real course link).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (academy direct-origin, no manager), D2 (real /library/<slug> course), D3 (academy-seed is
a live binary — no in-process write) + the M236 academy live checklist — see iter-07/decisions.md.
**Side-deliverables:** none.
**Routes carried forward:** the M236 Fate-3 handoff (the live proof + the M230 carry-forward live items + the
per-section live-calibration checklists from iters 05/06/07) → iter-08. After iter-08 the offline-buildable
scope is exhausted (all 3 non-sim sections land; the live proof legitimately routes to M236 per the ruling).
**Lessons:** the academy is the most live-coupled content product — catalog fill + progress-seed + the exact
route are ALL live. The honest offline floor is "the section resolves + points at a REAL course"; everything
that needs a live render or a platform-binary write is M236. Don't fabricate an academy chapter slug the
snapshot can't supply.
