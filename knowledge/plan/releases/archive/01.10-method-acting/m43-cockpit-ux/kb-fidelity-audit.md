---
title: "KB Fidelity Audit — M43 cockpit-ux"
date: 2026-06-26
scope: milestone:M43
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Cockpit UX (restyle / icons / CTA / manifest-download / login-overlay) | `corpus/ops/demo/cockpit-spec.md` (NEW — the M43 `Delivers →` target) | `rext demo-stack/cockpit.py`, `stack-seeding/seeders/cockpit.go` | BLIND-AREA (planned: the milestone's explicit deliverable) |
| Presenter cockpit (M38 mechanics: served panel, manifest single-source, handshake CTA) | `corpus/ops/demo/stories-spec.md` § "The presenter cockpit (M38)" | `rext demo-stack/cockpit.py`, `stack-seeding/seeders/cockpit.go` | PAIRED |
| Clerkenstein FAPI handshake (the CTA's seat-switch seam) | `corpus/services/clerkenstein.md` | (separate rext module; UNCHANGED by M43) | PAIRED |
| Demo UI tier (the cockpit launches INTO next-web/studio-desk/ant-academy) | `corpus/ops/demo/frontend-tier.md` | `rext demo-stack/up-injected.sh` | PAIRED |
| Demo-stack lifecycle (cockpit launch + offset-port model) | `corpus/ops/rosetta_demo.md` | `rext demo-stack/up-injected.sh` | PAIRED |

## Fidelity Findings

1. **stories-spec.md § cockpit (M38) — the two-CTA description.**
   - **Source:** `corpus/ops/demo/stories-spec.md:418` + `:461` + `:473`.
   - **Expected (doc):** two actions per hero — `[Login as]` (lands on app root) + `[Jump to section]` (lands on the hero's `jump_to`).
   - **Actual (code, pre-M43):** `cockpit.py` `login_as_url()` → app root; `jump_url()` → handshake to `jump_to`; `render_page()` renders BOTH buttons. **ALIGNED** with the pre-M43 code.
   - **Verdict:** ALIGNED (today). M43's CTA-unification deliverable (section 3) deliberately SUPERSEDES this: one unified `[Log in as]` → `jump_to`. This is planned evolution the milestone owns, not pre-existing drift.
   - **Fix owner:** doc — reconciled in Phase 5 (the new `cockpit-spec.md` documents the unified CTA; `stories-spec.md` gets a pointer / supersession note). Tracked as KB-1.

2. **clerkenstein.md — the `[Login as]` handshake seam.**
   - **Source:** `corpus/services/clerkenstein.md:74`.
   - **Expected/Actual:** `?__clerk_identity=<key>` on `/v1/client/handshake` selects the hero's seat. UNCHANGED by M43 (the CTA still drives the same handshake; only the redirect target + button count change on the cockpit side).
   - **Verdict:** ALIGNED. No edit needed.

3. **cockpit.go `defaultJumpForVantage` / `BuildCockpitManifest` — already emit `jump_to` per hero.**
   - **Source:** code (`stack-seeding/seeders/cockpit.go::defaultJumpForVantage`, `::BuildCockpitManifest`).
   - **Actual:** the manifest ALREADY carries a resolved `jump_to` for every hero (declared, or the vantage default via `defaultJumpForVantage`). So M43's CTA-unification reuses existing manifest data — no `cockpit.go` change is REQUIRED for the unified CTA (the data is already there). cockpit.go stays read-only unless a manifest-shape need surfaces.
   - **Verdict:** ALIGNED. Confirms section-3's "reuse `defaultJumpForVantage`" is data-already-present, not a new mechanic.

## Completeness Gaps

1. **(incidental, planned)** No standalone cockpit-UX spec exists today — the cockpit mechanics are scattered across `stories-spec.md` (M38), `clerkenstein.md` (the handshake), and `rosetta_demo.md` (the launch). This IS the M43 `Delivers →`: `corpus/ops/demo/cockpit-spec.md` graduates them into one place. Not a blocker — it is the milestone's documentation deliverable.

## Applied Fixes
None inline. The one doc-reconciliation (stories-spec.md two-CTA → unified, + the new cockpit-spec.md) is in-scope Phase-5 milestone work (the `Delivers →` + a supersession pointer), not a pre-flight fix — applying it now would document code that does not yet exist.

## Open Items (require user decision)
None.

## Gate Result
**YELLOW: proceed with tracking.** No unplanned blind areas (the cockpit-UX BLIND-AREA is the explicit `Delivers → cockpit-spec.md` deliverable). No stale load-bearing claim the milestone's implementation would read as truth and be misled by — today's docs accurately describe the pre-M43 code. Tracking item KB-1 (reconcile stories-spec.md's two-CTA description after the CTA unification) is recorded in `decisions.md` and addressed in Phase 5.
