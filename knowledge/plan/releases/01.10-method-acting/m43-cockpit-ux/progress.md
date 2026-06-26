# M43 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist

- [x] **(1) Light restyle** — `cockpit.py` `_PAGE_CSS` rewritten from the dark GitHub theme to a clean
  LIGHT professional design (CSS custom properties, card-per-hero hierarchy, indigo accent, subtle
  shadows + hover, high-contrast typography). RENDER-VERIFIED: bodyBg `rgb(244,246,251)`, white cards.
- [x] **(2) FontAwesome icons** — cdnjs FA6.5.2 free `<link>` (correct SRI) in `render_page()` `<head>`;
  `fa-circle-user` per hero avatar, `fa-building` per org, `fa-arrow-right-to-bracket` on the CTA,
  `fa-download` in the footer (+ icon CSS). RENDER-VERIFIED: `faApplied=true`, 6 hero icons + 2 org icons.
- [x] **(3) CTA unification** — dropped [Jump to section]; ONE `[Log in as]` per hero routed to the hero's
  `jump_to` per-role landing (reuses the manifest's resolved `jump_to`, single-sourced via
  `cockpit.go`'s `defaultJumpForVantage` — no `cockpit.go` change needed). RENDER-VERIFIED: 6 CTAs, 0
  jump buttons, each routing to the right per-role screen.
- [x] **(4) Seed-manifest download** — footer `[Download seed manifest]` → `/manifest.json`, now served
  with `Content-Disposition: attachment` (pretty JSON). RENDER-VERIFIED: link present + `download=` attr.
- [x] **(5) Staged login-progress overlay** — `Signing you in…` → `Loading your workspace…` → `Almost
  there…`, with `localStorage` in-flight state (30s window) across the FAPI-handshake redirect (generous
  final stage). RENDER-VERIFIED: overlay shows on click + stages through all three.
- [x] **Docs** — `corpus/ops/demo/cockpit-spec.md` **(NEW)** authored; reconciled the M38 two-CTA
  description in `stories-spec.md` (KB-1) + updated `demo/README.md` index + `CLAUDE.md` docs list.

**Status:** `done` (all 6 sections; render-accepted on demo-3). Zero next-web / platform-repo edits — the
cockpit is a standalone served panel; the code lives in rext (`demo-stack/cockpit.py`, tagged
`method-acting-m43-cockpit-ux`); cockpit.go UNCHANGED.
