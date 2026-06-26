---
milestone: M43
slug: cockpit-ux
version: v1.10 "method acting"
milestone_shape: section
status: planned
created: 2026-06-26
last_updated: 2026-06-26
complexity: small-medium
delivers: corpus/ops/demo/cockpit-spec.md (NEW — graduates the scattered M37/M38/clerkenstein presenter-cockpit mechanics into one standalone cockpit-UX spec — the restyled launcher, the FontAwesome icon set, the CTA-unification, the manifest download, and the staged login-progress overlay)
depends_on: none (parallel with M44 — different rext module)
spec_ref: .agentspace/scratch/roadmap-research-2026-06-26.md (the v1.10-extend research note; the presenter-cockpit strand)
---

# M43 — Cockpit UX polish

## Goal
Turn the **presenter cockpit** (a 253-line stdlib-Python static-HTML server, rext
`demo-stack/cockpit.py`, served at `:7700`+offset) from a dark, utilitarian panel into a **slick,
professional login launcher**: a clean light high-contrast design, FontAwesome icons next to each hero
+ org, **one** sensible CTA per hero (the unified **Log in as** that lands the hero on a per-role
screen), a **seed-manifest download**, and a **staged login-progress overlay** so the 2–5s
FAPI-handshake → next-web blank-load has feedback. **Zero next-web / platform-repo edits** — the
cockpit stays a standalone served panel (the v1.10 hard zero-edit line). This is the **presenter-grade**
extension of v1.10.

## Scope
**In:**
- **(1) Restyle** `cockpit.py` `_PAGE_CSS` (~lines 68–93) from the dark GitHub theme to a **clean LIGHT,
  high-contrast professional design** — a card hierarchy, deliberate spacing, an accent color.
- **(2) FontAwesome icons** — import the FontAwesome **FREE CDN** `<link>` in `render_page()`'s
  `<head>`, and render `fa-user-circle` before each hero name + `fa-building` before each org name (with
  icon CSS).
- **(3) CTA unification** — remove the **[Jump to section]** button; rename **[Login as]** →
  **[Log in as]** and route it to the hero's `jump_to` (a sensible per-role landing), **reusing**
  `cockpit.go`'s `defaultJumpForVantage`. One CTA per hero.
- **(4) Seed-manifest DOWNLOAD** — expose the existing `cockpit.py` `/manifest.json` endpoint as a
  **footer download link** (`Content-Disposition: attachment`).
- **(5) Staged JS login-progress overlay** — `'Signing you in…'` → `'Loading dashboard…'` →
  `'Initializing app…'`, with `localStorage` state carried across the redirect, so the FAPI-handshake →
  next-web blank-load shows progress (a generous final stage).

**Out:**
- Any **next-web / platform-repo edit** — the cockpit stays a standalone served panel (the hard
  zero-edit line).
- **Removing** the real handshake → next-web latency — only the **FEEDBACK** improves (the overlay), not
  the underlying ~2–5s.
- A **full bespoke design-system / brand pass** (clean professional light is the target, not a
  ground-up design system).
- **Presenter-feature deepening** (hero history, telemetry, note-taking) — UX polish only.

## Depends on / Parallel with
- **Depends on:** none.
- **Parallel with:** **M44** — a different rext module (M43 is `demo-stack/cockpit.py` + `cockpit.go`;
  M44 is `stack-seeding/seeders/`). The two run fully in parallel; both land before the close-release.

## Approach (default decisions — flagged ones are Open questions)
- **Aesthetic:** clean **LIGHT flat high-contrast** (card hierarchy + accent color) — decided default.
- **FontAwesome via the FREE CDN** (the user explicitly asked for the CDN `<link>`). Note: ant-academy
  vendors **FA Pro locally** (`code/public/assets/fontawesome/`) as an offline-safe precedent if a future
  offline-demo need ever surfaces — recorded, not adopted now.
- **Manifest download = the JSON** (`/manifest.json`, already available) — decided default.
- **Login progress = staged deterministic** with a generous final stage (no real progress signal from the
  cross-origin handshake; the overlay is feedback, not telemetry).

## Open questions
- The exact **design aesthetic** (default: clean light flat).
- Whether to **vendor FA** for offline-safe demos vs the CDN (default: CDN, per the user's ask).
- Whether **[Log in as]** shows the `jump_to`-label (e.g. "Log in as Maya → Profile") or stays bare.
- **Manifest as JSON** (`/manifest.json`) vs the source `stack.stories.yaml` (default: the JSON).

## KB dependencies
M43 reads these corpus docs as contract (it must not contradict them; it graduates + extends them):
- `corpus/ops/rosetta_demo.md` — the demo-stack lifecycle (Clerkenstein injection, the cockpit launch,
  the offset-port model) the cockpit is served within.
- `corpus/services/clerkenstein.md` — the `/v1/client/handshake` FAPI mechanic the **Log in as** CTA
  drives (**unchanged** by M43).
- `corpus/ops/demo/frontend-tier.md` — the demo UI tier the cockpit launches into.
- The scattered **M37/M38** cockpit mechanics (in `stories-spec.md` / `rosetta_demo.md`) — M43 graduates
  these into the new standalone `cockpit-spec.md`.

## Delivers →
- `corpus/ops/demo/cockpit-spec.md` **(NEW)** — graduate the scattered M37/M38/clerkenstein
  presenter-cockpit mechanics into **one standalone cockpit-UX spec**: the restyled light launcher, the
  FontAwesome icon set (+ the CDN-vs-vendored note), the CTA unification (`Log in as` →
  `defaultJumpForVantage`), the `/manifest.json` download, and the staged login-progress overlay.
- Reuse (no new mechanics): `cockpit.go`'s `BuildCockpitManifest` / `defaultJumpForVantage`, the
  Clerkenstein `/v1/client/handshake` (unchanged), and the existing `/manifest.json` endpoint.
