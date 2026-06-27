# M43 Retro — Cockpit UX polish

## Summary
The presenter cockpit went from a dark, utilitarian panel to a slick light professional login
launcher — five fills (light restyle, FontAwesome CDN icons, one unified `[Log in as]` CTA per hero,
a seed-manifest download, a staged login-progress overlay), all render-accepted on a live demo-3
(`:37700`) and USER-APPROVED the design. Tooling + docs only — zero next-web / platform-repo edits;
the cockpit stays a standalone served panel. Code-of-record is the rext authoring copy at tag
`method-acting-m43-cockpit-ux-fix1`; this corpus close merges the doc-half (the NEW
`corpus/ops/demo/cockpit-spec.md` + the stories-spec two-CTA reconciliation + index rows).

## Incidents This Cycle
- **P2 (caught in build, no escape):** the FA CDN `<link>` integrity hash had a one-char SRI typo
  (`...SCWr3w6A==` vs the real `...SCWr3W6A==`) that would have silently broken the stylesheet load
  (no icons). Caught at render-acceptance (`faApplied=true` is the proof), corrected, and pinned by
  `TestFontAwesomeSri` (the hash must decode to a well-formed 64-byte sha512 digest).
- **P2 (surfaced + fixed in harden Pass 1):** the meta-line **double-escape** — the story
  `annotation` was `html.escape`d at read-time and again at the meta join, double-escaping `<`/`&`/`"`
  (so `<i>shtick</i>` rendered as literal entity text to the presenter). Fixed to escape exactly once;
  pinned by `test_story_annotation_in_meta_is_escaped_exactly_once` (verified to FAIL on the pre-fix
  code).
- 0 regressions. 0 flakes (3 consecutive clean sequential runs of both suites).

## What Went Well
- The CTA-unification reused existing manifest data (`cockpit.go`'s resolved `jump_to` per hero) — **no
  `cockpit.go` change required**, confirming the Phase-0b read. A `cockpit.py`-only code change.
- Render-acceptance (the M44 lesson) caught the SRI typo before it could ship silently.
- Supply-chain stayed GREEN: FontAwesome via the free CDN `<link>` is a runtime asset, not a build
  dep — `cockpit.py` stayed stdlib-only, zero lockfile change.

## What Didn't
- Two escaping bugs (the double-escape + the SRI typo) in a string-templating render path — both the
  class of bug a stdlib-string HTML builder invites. Hardening's escaping-depth + SRI-well-formedness
  test families now pin both invariants; a future cockpit edit is guarded.

## Carried Forward
- None. M43 added zero deferrals. The cockpit **future-feature expansion surface** (per-hero
  history/telemetry, note-taking/talk-track, search/filter, live seed status) is recorded in
  `cockpit-spec.md` as a documented home for a future milestone — an expansion surface explicitly out
  of v1.10 scope, not a punt of in-scope work.

## Metrics Delta
- cockpit.py tests: **27 → 63** (+36 across 2 harden passes).
- cockpit.py statement coverage: 97% (maxed; the one fixable partial branch `269->271` closed).
- Go test funcs: unchanged (M43 is a Python-only change in a separate rext module).
- Supply-chain: GREEN (0 new deps). Flakes: 0. Alignment gates: all 5 GREEN (N/A change).
- Source: [`metrics.json`](metrics.json).
