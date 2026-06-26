# M43 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist

- [ ] **(1) Light restyle** — `cockpit.py` `_PAGE_CSS` (~lines 68–93) from the dark GitHub theme to a
  clean LIGHT high-contrast professional design (card hierarchy, spacing, accent color).
- [ ] **(2) FontAwesome icons** — import the FontAwesome FREE CDN `<link>` in `render_page()` `<head>`;
  `fa-user-circle` before each hero name + `fa-building` before each org name (+ icon CSS).
- [ ] **(3) CTA unification** — remove [Jump to section]; rename [Login as] → [Log in as], routed to the
  hero's `jump_to` per-role landing (reuse `cockpit.go`'s `defaultJumpForVantage`).
- [ ] **(4) Seed-manifest download** — expose the existing `/manifest.json` as a footer download link
  (`Content-Disposition: attachment`).
- [ ] **(5) Staged login-progress overlay** — `Signing you in…` → `Loading dashboard…` → `Initializing
  app…`, with `localStorage` state across the redirect (generous final stage).
- [ ] **Docs** — `corpus/ops/demo/cockpit-spec.md` **(NEW)**: graduate the scattered M37/M38/clerkenstein
  cockpit mechanics into one standalone cockpit-UX spec.

**Status:** `planned`. Zero next-web / platform-repo edits (the cockpit is a standalone served panel).
