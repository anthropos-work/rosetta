# M43 Spec Notes

Technical notes accumulate here during build. The authoritative design lives in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md)
(the presenter-cockpit strand). M43 is **tooling + docs only — zero next-web / platform-repo edits**; the
cockpit stays a standalone served panel (`rext demo-stack/cockpit.py`, served at `:7700`+offset).

## (1) Light restyle — `cockpit.py` `_PAGE_CSS` (~lines 68–93)
TODO: replace the dark GitHub-theme CSS with a clean LIGHT high-contrast professional design — card
hierarchy, deliberate spacing, an accent color. Keep the static-HTML / stdlib-only shape (no new deps).

## (2) FontAwesome icons — `render_page()` `<head>`
TODO: import the FontAwesome FREE CDN `<link>`; render `fa-user-circle` before each hero name +
`fa-building` before each org name (+ icon CSS). Note: ant-academy vendors FA Pro locally
(`code/public/assets/fontawesome/`) — the offline-safe precedent if ever needed.

## (3) CTA unification — remove [Jump to section], rename [Login as] → [Log in as]
TODO: drop the [Jump to section] button; rename the remaining CTA to **Log in as** and route it to the
hero's `jump_to` (per-role landing). REUSE `cockpit.go`'s `defaultJumpForVantage` (do not reinvent the
per-role landing map).

## (4) Seed-manifest download — the existing `/manifest.json`
TODO: expose `/manifest.json` as a footer download link with `Content-Disposition: attachment`. The
endpoint already exists (BuildCockpitManifest output) — wire the link + the header only.

## (5) Staged login-progress overlay (JS)
TODO: a staged deterministic overlay — `Signing you in…` → `Loading dashboard…` → `Initializing app…` —
with `localStorage` state carried across the FAPI-handshake → next-web redirect, so the 2–5s blank-load
has feedback. Generous final stage (no real cross-origin progress signal).

## Reuse (no new mechanics)
- `cockpit.go`: `BuildCockpitManifest`, `defaultJumpForVantage`.
- Clerkenstein `/v1/client/handshake` — UNCHANGED by M43.
- The existing `/manifest.json` endpoint.

## Delivers — `corpus/ops/demo/cockpit-spec.md` (NEW)
TODO: graduate the scattered M37/M38/clerkenstein cockpit mechanics into one standalone cockpit-UX spec
(the restyled launcher, the FontAwesome set + CDN-vs-vendored note, the CTA unification, the manifest
download, the staged overlay).
