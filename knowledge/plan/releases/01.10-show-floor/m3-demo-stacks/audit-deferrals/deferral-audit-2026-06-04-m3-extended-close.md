---
title: "Deferral Audit — M3 (extended-work close)"
date: 2026-06-04
scope: milestone
invoked-by: close-milestone
---

## Verdict
**YELLOW** — single deferrals only, all with clear in-release destinations now annotated. No repeat
deferrals, no chronic patterns, no aged-out items (M3 is v1.1's first milestone — a repeat is structurally
impossible; nothing is >3 months old). Not a blocker.

## Summary
- Total deferrals in scope: 7
- Single deferrals: 7 · Repeat: 0 · Chronic: 0 · Aged-out: 0
- Resolved (no longer deferred): 1 (M3-CF1)
- Fate-3 annotations applied this audit: 2 docs (M4, M5)

## Deferral Inventory + Fate verdicts

| # | Item | Origin | Fate | Destination |
|---|---|---|---|---|
| 1 | M3-CF1 running-app deployment proof | M3 S3 | **RESOLVED** (2026-06-04, proven live: 403-not-401) | — |
| 2 | `clerk-backend` api.clerk.com cert/redirect (the TLS trust for the orgclient) | M3 S3 `[~]` | **LAND-NEXT (Fate 3)** | M5 — annotated (finish the interactive demo recipe) |
| 3 | `clerk-webhook` live svix POST (feed identities into the sync pipeline) | M3 S3 `[~]` | **LAND-NEXT (Fate 3)** | M4 — annotated (post-seed identity feed mechanism) |
| 4 | **Demo identity = `user_clerkenstein`** must be seeded as the login user | M3 harden | **LAND-NEXT (Fate 3)** | M4 — annotated (prevents the seed-the-wrong-user 403 trap) |
| 5 | `casbin_rules` (plural) vs `casbin_rule` (singular) gorm gotcha | M3 migrate | **LAND-NEXT (Fate 3)** | M4 — annotated (alongside #4) |
| 6 | v1.0 express-gate CI carry-forward | v1.0 → M3-D | **LAND-NEXT (Fate 2)** | M5 — already in `In:` (confirmed, no edit) |
| 7 | Nightly auto-reaper of stale demos | M3-D2 | **KEEP** (deliberate design non-goal, not scope erosion) | — (add only if forgotten stacks become real) |

Also noted, not a deferral: 2-concurrent full stacks remains resource-gated (M3-D5 reduced acceptance to 1);
the mechanism is proven, only the bigger-box verification is pending → folded into M5's annotation as optional.

## Applied changes (Phase 5)
- `m4-declarative-seeding/overview.md` — In: scope: the login identity must match `DefaultDemoUser()`
  (`user_clerkenstein`), the casbin plural/singular gotcha, the clerk-webhook feed mechanism. (Items #3, #4, #5.)
- `m5-demo-corpus-recipes/overview.md` — In: scope: complete the clerk-backend cert/redirect + the
  browser-login walk-through as part of the recipe docs; optional 2-concurrent verification. (Item #2.)

## Blocking items
None. (No repeats, no aged-out, no chronic.) Close proceeds.
