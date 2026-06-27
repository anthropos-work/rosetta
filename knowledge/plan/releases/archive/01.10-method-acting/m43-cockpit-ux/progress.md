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

**Status:** `archived` (completed 2026-06-26) — all 6 sections; render-accepted on demo-3. Zero next-web / platform-repo edits — the
cockpit is a standalone served panel; the code lives in rext (`demo-stack/cockpit.py`, tagged
`method-acting-m43-cockpit-ux`); cockpit.go UNCHANGED.

## M43: Hardening

Full ledger: [`hardening-ledger.md`](hardening-ledger.md). 2 passes, stabilized (under the 5-pass cap).

### Scope manifest
M43-touched source (commit `4184650`): `demo-stack/cockpit.py` (existing tests:
`demo-stack/tests/test_cockpit.py`). `cockpit.go` UNCHANGED by M43 — its Go cockpit tests
(`stack-seeding/seeders/cockpit_test.go`) verified green as a regression guard, not deepened.

### Pass 1 — 2026-06-26
**Coverage delta (`cockpit.py`):** statements 97%→97% (already maxed — 3 missed are the untestable
`KeyboardInterrupt` handler + `__main__` guard); branches: the fixable partial `269->271` (falsy
`org_size`) **closed** (BrPart 2→1).
**Tests added (+32):** HTML-escaping depth (8, incl. the `jump_to` href attribute-injection
invariant), CTA-unification across multi-story/hero (5, incl. empty/absent-`jump_to` app-root
fallback), FontAwesome SRI well-formedness (4), manifest download incl. empty-manifest (3), overlay
JS structure/stages/30s-window/localStorage-guard (6), edge grid (7).
**Bugs fixed inline:** meta-line **double-escape** — `annotation` escaped at read-time then again at
the meta join; fixed to escape once, dedup on raw values (commit `5d1b99e`; regression test verified
to fail pre-fix).

### Pass 2 — 2026-06-26 (final)
**Coverage delta:** none (statement coverage maxed; Pass 2 added behavioural locks line-coverage
can't see). **Tests added (+4):** served-panel depth — `/index.html` alias, Content-Length =
UTF-8 byte length, unicode round-trip, HTML content-type (commit `14e68fe`).
**Knowledge backfill:** none warranted (the fix restored a documented escaping invariant; new tests
pin existing `cockpit-spec.md` contracts).

### Stop condition
Stabilized — statement coverage maxed, the one fixable partial branch closed, Pass 2's six-dimension
scan surfaced only behavioural locks (no delta, no new bugs), zero flakes across 3 sequential runs of
both suites (Python 27→63 tests; Go cockpit green). Tags: rext `method-acting-m43-cockpit-ux-fix1`
(source-fix bump over the build tag).
