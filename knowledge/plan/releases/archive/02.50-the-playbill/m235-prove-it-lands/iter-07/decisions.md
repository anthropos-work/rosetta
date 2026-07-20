# iter-07 — decisions

## D1 — academy is app_base=academy, a direct-origin CTA, NO manager view
The registry marks `skill-path-new` (academy) `appBase="academy"`, `managerKind=""`. The projection academy
arm builds a link-bearing row (a `/library/<slug>` player route) with NO manager view (M231: no academy
manager result route). The cockpit renders it as a DIRECT academy-origin link (the M53 `e2e_persona=member`
cookie seam), NOT a FAPI handshake — proven by `TestBuildNonSimProducts_AcademyDirectOriginNoManager` +
`test_academy_content_story_renders_direct_origin_cta_with_persona_seam`.

## D2 — the offline route is a REAL public course `/library/<slug>` (resolves, non-fabricated)
The academy chapter route `/chapters/<slug>/` uses an academy-app-specific CHAPTER slug that is NOT
offline-derivable (the snapshot's `chapter_list` carries chapter id+title but no slug). `/library/<slug>` is
the only offline-derivable REAL academy route — a public course that resolves in M230's filled catalog. The
exhibit pins a real public course slug (`foundation-of-artificial-intelligence`, from the captured snapshot's
`directus.skill_paths.slug`). So the CTA lands on a REAL course (not a dead link); upgrading it to the
progress-bearing chapter route is M236.

## D3 — the academy seeder arm writes NO rows here (academy-seed is a live platform binary)
Per-member `academy_chapter_progress` / `academy_last_activity` are written by the `app/cmd/academy-seed`
PLATFORM binary (M231 §6) — a live DB write needing M230's filled catalog. That is NOT an in-process rext
Conn write, so `ContentStoryNonSimSeeder`'s academy arm seeds nothing; it advances the flat index (owner
stays aligned) and defers to the M236 bring-up invocation.

## M236 ACADEMY LIVE CHECKLIST (all academy substance is live — route to M236)
1. **academy_chapter_progress write.** Wire `app/cmd/academy-seed` into the M236 cold bring-up:
   `academy-seed --user-id <the academy content-player owner's user-id> --fixture in-progress` (or
   `completed`), so the academy renders REAL progress for the seeded member. INVOKE the built binary (within
   the zero-edit wall) — do NOT edit it; if a platform-source change is needed → demopatch or escalate.
2. **the progress-bearing route.** `/library/<slug>` is the anonymous course preview (no progress); calibrate
   the CTA to the authed progress-bearing route (the academy home resume-row or `/chapters/<chapterSlug>/`)
   against a live render, using the real chapter slug from the filled catalog.
3. **M230 catalog-fill dependency.** The academy grid must be filled (the `playbill-m230-academy-fs-published`
   demo-patch applied at bring-up) or the course/chapters don't render — the M230 carry-forward cluster 1
   (already routed to M235/M236). This is the ANT_ACADEMY coverage descriptor's prerequisite.
