---
iter: 03
iteration_type: tik
status: closed-fixed
date: 2026-06-03
---

# iter-03 — tik (under TOK-01): the authn twin

**Active strategy:** TOK-01 (easy-offline-side first → authn before orgclient).
**Target:** build Clerkenstein's authn-provider twin + runner + authn goldens → the **first real
alignment score** on the 6 critical authn genes.
**Hypothesis:** the authn side is offline-capturable (local JWT), so a faithful claim-extraction twin
moves the score with no live-Clerk dependency.
**Expected lift:** 0% → the authn slice (4 VerifyToken genes aligned ≈ +21% overall / +31% critical).
**Phase plan:** build (HS256 mint/verify + `colony/authn.Provider` impl + runner) → hand-author authn
goldens → `alignctl run` → re-measure.

See `progress.md` for the close.
