# M43 Hardening Ledger

The harden pass for **M43 — Cockpit UX polish** (release v1.10 "method acting"). Deepens test
coverage over the cockpit footprint in the rext authoring copy
(`.agentspace/rosetta-extensions/demo-stack/cockpit.py` + `tests/test_cockpit.py`; the Go cockpit
tests in `stack-seeding/seeders/cockpit_test.go` verified green — `cockpit.go` was UNCHANGED by M43).
The cockpit is rext tooling; the code + tests live in the rext repo, this ledger lives in the corpus.

**Footprint** (M43 commit `4184650`): `demo-stack/cockpit.py` (the stdlib-only static-HTML presenter
launcher) + `demo-stack/tests/test_cockpit.py`. `cockpit.go` was not touched (M43 reused
`BuildCockpitManifest` / `defaultJumpForVantage`).

**Audits:** supply-chain GREEN (zero `go.mod`/`go.sum`/`package.json`/lockfile change — cockpit.py is
stdlib-only Python; FontAwesome is a CDN `<link>`, not a dependency). Alignment N/A — zero
`clerkenstein/` change, carry 100%.

---

## Pass 2 — 2026-06-26 — final

(Two harden passes ran in this session; this is the final-mode entry covering both — the loop
stabilized after Pass 2's scan surfaced only behavioural-lock additions and no new bugs.)

**Coverage delta (milestone-touched files — `cockpit.py`):**
- Statements: 97% -> 97% (117 stmts, 3 missed: lines 422-423 = the `except KeyboardInterrupt: pass`
  handler, 430 = the `if __name__ == "__main__"` guard — both structurally untestable; statement
  coverage was already maxed at build time).
- Branches: the one fixable **partial branch** `269->271` (a story with a falsy `org_size`) was
  **closed** in Pass 1 — `BrPart` 2 -> 1 (the remaining partial is on the untestable interrupt
  handler). Branch coverage genuinely improved; the only gaps left are the interrupt/main-guard lines
  that no in-process test can reach.

**Tests added (27 -> 63 cockpit Python tests; +36):**

*Pass 1 (+32):*
- `TestHtmlEscapingDepth` (8): org/story-name/role/annotation/stack/fapi-host/app-base interpolation
  points all escape; the **`jump_to` href attribute-injection invariant** (a `"` in `jump_to` is
  doubly defended — `urlencode` → `%22` then `html.escape(quote=True)`, so it can't break out of
  `href="…"`); the `data-login-as` attribute quote-escaping.
- `TestCtaUnificationInvariant` (5): across a multi-story/multi-hero manifest — exactly one
  `[Log in as]` per hero, zero jump buttons, each CTA routes to **its own** hero's `jump_to`, with
  the app-root fallback when `jump_to` is empty OR the key is absent.
- `TestFontAwesomeSri` (4): the FA CDN `<link>` is present with the required attrs; the **SRI hash is
  a well-formed `sha512-<base64>` decoding to a 64-byte digest** (the regression for the mid-build
  one-char SRI typo); the four FA icon classes (avatar / building / CTA / download) render.
- `TestManifestDownloadEdges` (3): `/manifest.json` serves `Content-Disposition: attachment` + valid
  **pretty** JSON incl. the empty-manifest edge; the footer link's `download=` attr.
- `TestOverlayJs` (6): the embedded overlay JS has balanced braces/parens; the three stages present
  **in order**; the **30s in-flight window** + the localStorage key; localStorage access guarded by
  `try/catch` (graceful when unavailable); never `preventDefault`s the real handshake nav; the
  markup + `aria-live` in the rendered page.
- `TestRenderEdgeCases` (7): no-hero story, no-badge hero, no-org story, **zero `org_size`** (closes
  the `269->271` partial branch), `annotation == org` dedup, offset-port URL construction, a 4x3
  large manifest.

*Pass 2 (+4) — `TestServedPanelDepth`:*
- `/index.html` is a byte-identical alias for `/`; **Content-Length is the UTF-8 byte length** (not
  the char count — a `len(string)` refactor would truncate multi-byte hero/org names); unicode
  hero/org/role round-trip + decode as UTF-8; the page route sets `text/html; charset=utf-8`.

**Bugs fixed inline (regression test pins it):**
- **Meta-line double-escape** (`cockpit.py` `render_page`, commit `5d1b99e`). The story `annotation`
  was `html.escape`d at read-time and then escaped **again** when the meta line was joined
  (`html.escape(" · ".join(meta_bits))`), double-escaping any `<`/`&`/`"` — e.g. `<i>shtick</i>`
  rendered as the literal entity text `&lt;i&gt;shtick&lt;/i&gt;` to the presenter instead of the
  intended escaped-once display. Fix: keep the raw `org`/`annotation` for the dedup comparison and
  escape the meta string **exactly once** at the join. Pinned by
  `TestHtmlEscapingDepth.test_story_annotation_in_meta_is_escaped_exactly_once` (asserts the raw
  markup is gone, the single-escape form is present, AND the double-escape marker `&amp;lt;i&amp;gt;`
  is absent — verified to FAIL on the pre-fix code).

**Flakes stabilized:** none observed. 3 consecutive clean sequential runs of both suites (Python
63/63, Go cockpit pass).

**Knowledge backfill:** no KB-worthy findings to propagate. The double-escape was an
implementation-internal escaping bug (no external behavioural contract changed — the cockpit's
documented behaviour is "escape user-supplied manifest values"; the fix restores correct
single-escaping). The new tests pin existing documented invariants (CTA unification per
`cockpit-spec.md`, the SRI, the staged overlay) rather than discovering new ones. The
`cockpit-spec.md` authored in M43's build Phase 5 already documents the surfaces; nothing to add.

### Stop condition
**Stabilized.** Statement coverage maxed (only the structurally-untestable interrupt handler +
`__main__` guard remain); the one fixable partial branch closed in Pass 1; Pass 2's full
six-dimension scan surfaced only behavioural-lock additions (no coverage delta, no new bugs); zero
flakes across 3 sequential runs. Loop terminated at Pass 2, well under the 5-pass cap.

`/developer-kit:close-milestone` is still the next step — its Phase 4 audit runs independently as
defense-in-depth.
