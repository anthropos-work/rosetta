# M43 Spec Notes

Technical notes accumulate here during build. The authoritative design lives in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md)
(the presenter-cockpit strand). M43 is **tooling + docs only — zero next-web / platform-repo edits**; the
cockpit stays a standalone served panel (`rext demo-stack/cockpit.py`, served at `:7700`+offset).

## Pre-flight audits — (1) Light restyle (first section)
- **Phase 0b KB-fidelity (M43): YELLOW.** Report: `kb-fidelity-audit.md`. No unplanned blind areas (the
  cockpit-UX BLIND-AREA is the explicit `Delivers → cockpit-spec.md` deliverable); today's docs accurately
  describe pre-M43 code; one tracking item KB-1 (reconcile `stories-spec.md`'s two-CTA description) addressed
  in Phase 5. Audited at sha `$(release/01.10-method-acting @ a2b0d45)`.
- **Topic→doc→code triples** (fast-start for later sections):
  - Cockpit UX → `corpus/ops/demo/cockpit-spec.md` (NEW) → `rext demo-stack/cockpit.py` + `stack-seeding/seeders/cockpit.go`.
  - M38 cockpit mechanics → `corpus/ops/demo/stories-spec.md` § "The presenter cockpit (M38)" → same code.
  - Handshake CTA seam → `corpus/services/clerkenstein.md` → clerkenstein (UNCHANGED by M43).
- **Key code facts** (confirmed by audit): the cockpit manifest ALREADY carries a resolved `jump_to` per hero
  (`cockpit.go::BuildCockpitManifest` via `defaultJumpForVantage`), so the CTA-unification reuses existing data
  — **no `cockpit.go` change is required**; M43 is a `cockpit.py`-only code change + the new doc. The live
  manifest (demo-3) has 2 stories / 6 heroes (Maya/Tom/Dan @ Cervato, Sara/Nick/Leah @ Solvantis), each with a
  resolved `jump_to` — ready for render-acceptance.

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

## Render-acceptance (the mandatory visual gate — M44 lesson)
Restarted the cockpit on `:37700` from the bumped consumption clone
(`stack-demo/rosetta-extensions` @ tag `method-acting-m43-cockpit-ux`) against the live demo-3 manifest, and
rendered it with Playwright/chromium (the rext `stack-verify/e2e/` harness). **All 5 deliverables verified
VISUALLY + via rendered-DOM facts** (screenshots in `.agentspace/scratch/work-m43/render-check/`):
- `cockpit-new.png` — light slick UI (bodyBg `rgb(244,246,251)`, white cards, indigo accent, pill badges, FA
  icons rendering), 6 `[Log in as]` CTAs, 0 jump buttons, footer download link.
- `cockpit-login-overlay.png` — the staged overlay on click (centered white card + indigo spinner over a
  dimmed backdrop), stages `Signing you in…` → `Loading your workspace…` → `Almost there…`.
- DOM facts: `faApplied=true` ("Font Awesome 6 Free"), 6 `fa-circle-user` + 2 `fa-building`, every CTA routes
  to the hero's `jump_to` (employees → `/profile`, managers → their workforce deep-link).

**SRI gotcha (caught + fixed during build):** the initial FA CDN integrity hash had a one-char typo
(`...SCWr3w6A==` vs the real `...SCWr3W6A==`) which would have silently broken the stylesheet load (no icons).
Fetched the real cdnjs SRI for FA 6.5.2 and corrected it; the render-acceptance (`faApplied=true`) is the proof
the hash is right.

## Tag / two-repo state
- rext authoring (`.agentspace/rosetta-extensions`): committed `4184650`, tagged `method-acting-m43-cockpit-ux`.
  **Harden bump:** commits `5d1b99e` (deepen tests 27→63 + meta-line double-escape fix) + `14e68fe`
  (served-panel depth), tagged `method-acting-m43-cockpit-ux-fix1`. See
  [`hardening-ledger.md`](hardening-ledger.md).
- consumption clone (`stack-demo/rosetta-extensions`): fetched + checked out the tag (was `…m44…fix1`).
- corpus m43 branch: the `Delivers → cockpit-spec.md` doc + KB-1 reconciliation + plan files.

## Reuse (no new mechanics)
- `cockpit.go`: `BuildCockpitManifest`, `defaultJumpForVantage`.
- Clerkenstein `/v1/client/handshake` — UNCHANGED by M43.
- The existing `/manifest.json` endpoint.

## Delivers — `corpus/ops/demo/cockpit-spec.md` (NEW)
TODO: graduate the scattered M37/M38/clerkenstein cockpit mechanics into one standalone cockpit-UX spec
(the restyled launcher, the FontAwesome set + CDN-vs-vendored note, the CTA unification, the manifest
download, the staged overlay).
