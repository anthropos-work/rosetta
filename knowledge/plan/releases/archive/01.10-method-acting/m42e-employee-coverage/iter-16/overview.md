---
iteration_type: tik
iter_shape: production-fix
status: planned
created: 2026-06-25
---

# iter-16 — P4: REAL-photo avatars (menu==profile) + org logo

**Active strategy reference:** TOK-10 (persona-believability-by-root-cause). P4 closes design-plan root
causes #5 (org logo broken-glyph) + #6 (avatar menu≠profile + not real).

**Step 0 re-survey:** confirmed live on demo-3 — Maya's `users.picture` is the generated SVG cartoon
data URI (avatar.go, the M39 G4 illustrated face); the menu reads Clerk `user.imageUrl` which is ABSENT
(no image field on the FAPI userRes) → a silhouette; org `logo_url` is NULL for both orgs and the
Clerkenstein orgRes carries no image → a broken glyph in the top menu. Still the before-state. Same
TOK-10 strategy; P4 is a routed-forward target from iter-14.

**Cluster / target identified (design-plan root #5 + #6):**
1. **Avatar** — replace the SVG cartoon with a LICENSE-CLEAN REAL human portrait (user decision:
   licensed real-person stock photos). Source a small set offline, vendor into
   `stack-seeding/assets/avatars/` with a per-file LICENSE/SOURCE note, assign deterministically per
   hero (by uuid), and thread the SAME image bytes to BOTH `public.users.picture` (the /profile avatar)
   AND the Clerkenstein roster→userRes image (`image_url` + `has_image`) so the MENU avatar == the
   PROFILE avatar and neither is a silhouette/SVG.
2. **Org logo** — seed `organizations.logo_url` (a deterministic offline data-URI MONOGRAM — the orgs
   are fictional, so a generated mark is the honest choice; avatar.go data-URI precedent) + thread an
   org image through the Clerkenstein orgRes so the top-menu broken-image glyph becomes a real logo.

**Hypothesis:** with a real-photo data URI on both `users.picture` and the FAPI userRes image, and a
monogram logo on `organizations.logo_url` + the orgRes, /profile + the top-menu show the SAME real
photo for Maya and the org shows a logo (not a broken glyph).

**License gate (HARD):** use ONLY clearly license-clean REAL photos (CC0 / public-domain / explicitly
free-for-commercial-use, no attribution-or-consent encumbrance). If no clearly license-clean real-photo
source can be found, STOP and surface a user-blocker — do NOT use questionable images.

**Expected lift:** qualitative — menu avatar == profile avatar (both a real photo); org logo renders.

**Phase plan:** Phase B (source + license-vet the photo set) → Phase C (vendor + avatar.go real-photo
selector + thread to users.picture + roster + clerkenstein userRes/orgRes image + org logo_url) →
Phase D (re-seed demo-3, re-export roster, restart fake-fapi, probe both surfaces). Clerkenstein touched
→ keep the 5 alignment gates green.

**Escalation conditions:** license-unclean photo source with no clean alternative → user-blocker
(SEVERITY blocker per the run prompt).

**Acceptable close-no-lift:** if the only blocker is licensing and no clean source exists → user-blocker
(not a close).
