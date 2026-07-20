---
iter: iter-07
milestone: M235
type: tik
iteration_type: tik
status: planned
created: 2026-07-20
active_strategy: TOK-01 (two-track) — Track A step 2 (non-simulation product sections)
---

# iter-07 — academy content-story section (the 3rd + most live-coupled non-sim section)

**Type:** tik · **Active strategy:** TOK-01 Track A step 2.

## Step 0 — re-survey
iter-05 (skill-path) + iter-06 (ai-labs) landed 2/3. Academy is the last offline surface. M231 §6 ruled it
**IN / real progress** (backend-authoritative `academy_chapter_progress` via `app/cmd/academy-seed`), but ALL
its substance is LIVE: the catalog fill is M230's demo-patch, the per-member progress is the `academy-seed`
platform binary, and the exact `/chapters/<slug>/` route needs the filled catalog. Confirmed offline: the
snapshot's `chapter_list` carries chapter id+title but NO slug (the academy chapter slug is app-specific), and
`/library/<slug>` is the only offline-derivable REAL academy route (a real public course, resolves). Target
meaningful; the offline floor is the manifest section + a real course CTA; the substance is M236.

## Cluster / target identified
The **skill-path-new (academy)** section: app_base=academy, a real public course deep-link
(`/library/<slug>` — a resolvable non-fabricated route), Label=course title, NO manager view (M231: no
academy manager route). The cockpit already renders the direct academy-origin CTA (e2e_persona=member seam).

## Hypothesis
An academy exhibit + the academy arm in `resolveNonSimSession` (app_base academy, `/library/<slug>` player
route, no manager) makes the academy section RESOLVE, unit-proven; the cockpit renders the direct academy CTA.
The `academy-seed` progress-write + the progress-bearing chapter route + the M230 catalog dependency are the
M236 live set (documented).

## Expected lift
Readiness 3/3 non-sim sections (all three content-product non-sim sections resolve).

## Phase plan
- `content_nonsim.go`: an `AcademySlug` field + 1 academy exhibit (a real public course) + the academy arm in
  `resolveNonSimSession` (app_base academy, `/library/<slug>`, no manager). The seeder academy arm is a
  documented no-op — the progress-write is `academy-seed` (a live platform binary), routed to M236.
- Regenerate `presets/content-manifest.json`; honesty gate green.
- Unit tests: academy projection (app_base academy, no manager, real course route) + a cockpit academy render
  test (direct academy-origin CTA, e2e_persona seam).

## Escalation / close-no-lift
- The academy progress-write + the exact progress-bearing chapter route + the M230 catalog fill are LIVE →
  M236 (documented, NOT faked here). If the offline manifest section can't be made honest (a real course link)
  it closes-no-lift with the constraint recorded — but `/library/<slug>` IS a real resolvable route, so the
  section is offline-honest.
