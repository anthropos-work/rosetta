**Type:** tik (P8 harness fix, under TOK-10)

# iter-23 — P8 harness fix: avatar-consistency false-fail + the authoritative gate + manager smoke-sweep

## Phase A/B — the false-fail diagnosis
The iter-22 fresh-demo-up employee sweep reached `failingSections=0 escapes=0 notReached=0
frontier=EXHAUSTED` but `personaFailures=1` (avatar-consistency=FAIL: "menu avatar is not a real photo
kind=svg-data-uri"). A live diagnostic (a throwaway `diag-avatar` spec, since removed) proved the data +
threading are CORRECT: clerk `user.imageUrl` = the real JPEG (8011 B, `hasImage=true`), FAPI `/v1/me`
`image_url` = the real JPEG, and the header DOM carries the user JPEG `<img>` (3rd of 3: wordmark, org
monogram SVG, user JPEG). The harness's `readMenuAvatarSrc` returned the COMPANY-SWITCHER org monogram
(`alt="company logo"`, a `data:image/svg+xml` whose base64 carries no "logo" token, so the src-only
`/logo|wordmark/` skip missed it) instead of the user avatar sitting right after it → false FAIL.

## Phase C — fix (rext, harness-only)
`persona-assert.ts`: `readMenuAvatarSrc` excludes the org/company img by its `alt` too (CompanySwitcher sets
`alt="company logo"`) AND prefers a raster/remote candidate over a residual SVG so the user's real photo
always wins. Same `alt` filter hardens `readProfileAvatarSrc`. Zero platform edit.

## Phase D — re-measure (the AUTHORITATIVE employee gate, on the consumed tag)
Committed + tagged `iter23`, bumped the consumed clone, then ran the gate from the CONSUMED clone (committed
code) vs the fresh zero-manual demo-3:
`reachable=62/150 failingSections=0 personaFailures=0 escapes=0 notReached=0 frontier=EXHAUSTED → GATE MET`
(`gateMet: true`; all 3 persona checks PASS: role-skills-coherence, avatar-consistency [menu==profile==real
photo], org-identity). **THE EMPLOYEE GATE IS MET.**

## Manager smoke-sweep (M42m input — NOT fixed here)
`dan-manager` vs the same fresh demo-3 (consumed iter23):
`reachable=150/150 failingSections=0 personaFailures=0 escapes=139 notReached=5 frontier=CAPPED(+79) → NOT MET`.
- **persona PASS (all 3)** — the M42e identity machinery generalizes to the manager hero for free.
- **escapes=139, ALL `studio.anthropos.work`** — ONE root cause: the baked left-nav "Studio" prod link on
  every authenticated page (the manager/enterprise nav renders it; the employee nav doesn't → employee had
  0). The fix-surface routing table maps this to demo injection + env link-rewriting — a single fix clears
  all 139.
- **notReached=5 = ALL `/workforce/*` M36 dashboard pages** (Workforce landing / teams / role-readiness /
  succession / mobility) — the manager nav doesn't link them as seeded; the core M42m content + nav work.
- **frontier CAPPED(+79)** — the manager's team-roster fan-out (`/user/<id>/skills` + `/user/<id>/activities`
  per team member) exceeds 150; M42m needs a representative-sample rule for the template-identical
  `/user/<id>` set + a higher cap to EXHAUST before grading.
- **failingSections=0 is over the REACHED set only** — the workforce content gate is UNMEASURED until those
  5 pages are reachable.

## Close — 2026-06-25

**Outcome:** the avatar-consistency FAIL was a harness selector collision (org logo vs user avatar), not a
data/platform gap. Fixed; the AUTHORITATIVE employee gate (consumed iter23 vs fresh zero-manual demo-3) is
MET — M42e employee believability is reproducibly proven. The manager smoke-sweep calibrated the M42m
residual.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET (employee — `gateMet: true`, 0 failingSections / 0 personaFailures / 0 escapes / 0 notReached /
frontier EXHAUSTED, against committed code on a fresh demo-up)
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1 (gate-met)
**Decisions:** the alt-based org-img exclusion + raster-preference in the menu/profile avatar selectors.
**Side-deliverables:** the diag-avatar diagnostic (temporary, removed after use).
**Routes carried forward (→ M42m, the separate milestone — NOT this milestone's work):**
- M42m-escapes: rewrite the baked `studio.anthropos.work` left-nav link (demo injection / env link-rewriting)
  → clears all 139 manager escapes.
- M42m-workforce: make the 5 `/workforce/*` M36 dashboard pages reachable + seeded (the manager content gate).
- M42m-frontier: a representative-sample rule for the `/user/<id>` team-roster fan-out + raise the cap so the
  manager frontier EXHAUSTS before grading.
- M42m-manifest: the manager manifest pages are `calibrated:false` — the section selectors/floors need a
  manager calibration pass once the workforce pages render.
**Lessons:** a believability gate's persona checks must distinguish co-located chrome (the org logo) from the
asserted element (the user avatar) by a stable attribute (`alt`), not just a src-substring — a data-URI src
carries no semantic token to skip on. The employee identity machinery generalizes to other vantages for free
(manager persona PASSED unchanged).
