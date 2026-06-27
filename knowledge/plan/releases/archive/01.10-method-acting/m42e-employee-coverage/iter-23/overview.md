---
iteration_type: tik
iter_shape: tooling
status: planned
created: 2026-06-25
---

# iter-23 — P8 harness fix: avatar-consistency false-fail (org-logo selector collision)

**Active strategy reference:** TOK-10 (persona-believability-by-root-cause). A P8-surfaced harness bug: the
employee sweep reported `avatar-consistency=FAIL` while the avatar was actually correct (menu==profile==real
JPEG). Routed here as a tooling-iter (the gate-measuring tool mis-measures).

**Cluster / target identified:** `persona-assert.ts readMenuAvatarSrc`. The iter-22 fresh-demo-up sweep
reached `failingSections=0 escapes=0 notReached=0 frontier=EXHAUSTED` but `personaFailures=1`
(avatar-consistency). A live diagnostic (diag-avatar spec) proved the data + threading are CORRECT — clerk
`user.imageUrl` = the real JPEG (8011 B, hasImage=true), FAPI `/v1/me` `image_url` = the real JPEG, and the
header DOM carries the user JPEG `<img>`. The harness's `readMenuAvatarSrc` returned the COMPANY-SWITCHER org
monogram (`alt="company logo"`, a `data:image/svg+xml` whose base64 carries no "logo" token, so the src-only
`/logo|wordmark/` skip missed it) instead of the user avatar that sits right after it → false FAIL.

**Hypothesis:** excluding the org/company img by its `alt` (CompanySwitcher sets `alt="company logo"`) AND
preferring a raster/remote candidate over a residual SVG makes `readMenuAvatarSrc` land on the user's real
photo → avatar-consistency PASSES, the employee gate is MET.

**Phase plan:** Phase C fix (persona-assert.ts `readMenuAvatarSrc` + harden `readProfileAvatarSrc` with the
same alt filter) → Phase D re-sweep (authoring harness vs live demo-3) → Phase E close.

**Expected lift:** `personaFailures 1 → 0` → the employee gate flips to MET (the other 3 components already
at 0 on the fresh demo-up).

**Acceptable close outcomes:** the selector fix lands + the re-sweep is gate-met (employee), proving the
avatar believability was always correct and only the measurement was wrong. The diagnostic spec is removed
(temporary).
