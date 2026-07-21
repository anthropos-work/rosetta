# M238 — Decisions

_(implementation choices with rationale accumulate here during build)_

## D1 — #3 fix: FS-as-published chapter-body demopatch (NOT backend-wiring). Resolves the open question.
**Open question (overview):** chapter-body FS-fallback demopatch vs wiring the academy backend for the demo — which
is revert-clean + sufficient?
**Decision: the FS-as-published chapter-body demopatch.** Rationale:
- **Consistency.** The catalog is ALREADY FS-published (the shipped `academy-fs-published-fallback` M230 patch renders
  the grid from the committed FS tree). Wiring the backend for bodies would mean the grid shows FS cards but bodies
  come from a DB that has no matching rows → mismatch. The body path must be FS-published TOO, gated on the SAME
  `ACADEMY_DEMO_FS_PUBLISHED` env var, so grid + body are one coherent behavior.
- **Revert-clean + zero platform edits.** A native-run demopatch on the demo's ephemeral clone (apply-before-launch,
  revert-on-stop, byte-exact pristine restore — proven in `test_academy_fs_published_body.py::TestLadder`). The
  canonical `anthropos-work/ant-academy` repo is never touched. Wiring the backend would require exporting the whole
  catalog into the demo DB (`content-export` + `academy-seed`) — heavy, and `ant-academy.md` already records those
  seeded rows are MOOT on a demo (no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` → FS fallback path, nothing reads them).
- **Minimal + safe.** The patch mirrors the existing dev-only `maybeResolveDraftBody` (which already serves the FS
  body unlocked) but gated on the demo env var instead of `draftsEnabled() ∧ isFsDraftSlug`, with `draft:false` (no
  chip). Behavior-identical when the env is unset — the pristine `return { notFound: true }` stands (safe even if
  upstreamed). Anchor `return { notFound: true }` occurs exactly once; pre/post sha pinned.
**Artifacts:** rext `demo-stack/patches/academy-fs-published-chapter-body/*.yaml` + `stack-injection/apply-academy-fs-published-body.sh` + `demo-stack/ant-academy.sh` wiring + `demo-stack/tests/test_academy_fs_published_body.py`.
**Live verify:** performed jointly with §2 in ONE billion session (both #3 Start→render and #2 language re-triage need the same fresh-academy bring-up) — not a deferral, executed this milestone.

## D2 — #2 language verdict: NOT a distinct code bug. Resolved by #3 + two documented findings.
Investigated the ant-academy i18n on the FRESH clone (`stack-demo/ant-academy` @ main). The language surface on
current code is robust — coerceLocale falls back to EN, translate.js returns the inline-English `defaultEn` on any
key miss, the switcher is a sound EN↔IT toggle `<Link>`. The M237 "language error" decomposes into three parts:

1. **`/it` returns 404** = **EXPECTED, not a bug.** Locale lives in the URL as a `?lang=<code>` QUERY PARAM only
   (`src/i18n/locale.js:6-8`, explicit by-design comment); there is **NO `/[locale]` path-prefix route** (confirmed:
   0 locale route dirs under `app/`). A bare `/it` URL 404s because it was never a route. The M237 triage tested a
   URL shape the app never supported.
2. **The switcher "shows no working language menu"** = it is a **2-way EN↔IT toggle `<Link>`** (`LocaleSwitch.jsx`),
   NOT a dropdown menu — clicking it navigates to `${pathname}?lang=it`. Sound on current code (no backend
   dependency; translate.js can't error). M237 observed it on a 5-behind clone; "no menu" describes a toggle, not a
   defect.
3. **Switching language on a CHAPTER reader → 404** = **the SAME backend-null path as #3**, and is **FIXED by the #3
   patch.** `chapters/[slug]/page.jsx:167,181` reads `?lang=` (`localeFromSearchParams`) and threads `locale` into
   `resolveServerChapterBody(slug, locale)`; a demo's backend is null → pre-fix 404. The #3 patch serves the
   **locale-aware** FS body (`safeLoadFsShape → loadChapterShape → resolveChapterPath` prefers the `it/` overlay,
   falls back to canonical EN). So switching to IT on a demo chapter now renders (IT overlay where it exists, else EN).

**Verdict:** #2 needs **NO new code/tooling.** Its one actionable defect (chapter-language 404) lands via #3 this
milestone (Fate 1); the `/it` 404 and the toggle-not-menu are FINDINGS, not defects. **No ant-academy code defect
surfaced** → nothing to escalate under zero-platform-edit. **To be confirmed on a fresh billion academy** (the brief
requires measuring fresh, not assuming) — jointly with the #3 Start-renders verify.

## D3 — §3 academy coverage sweep: a chapter-body fence + a general `mustNotInclude` negative marker
The sweep asserted only the academy HOME grid (`ANT_ACADEMY_HOME_SECTION`). Extended it so a hero must reach a
CHAPTER BODY + switch its language:
- **`mustNotInclude` (general).** Added a negative-marker to the `text`/`both` `realContent` kinds in
  `section-assert.ts` — a region carrying any forbidden substring FAILs `error`. The honest way to encode "this is
  NOT the 404": `isErrorText`'s sentinel list doesn't include the academy's whimsical "You wandered off the trail."
  copy, and a bare floor is the superseded density check. Unit-tested (both polarities).
- **`ANT_ACADEMY_CHAPTER_SECTION`.** `mustNotInclude: ['You wandered off the trail']` (the exact `not-found.jsx`
  `<h1>`) + `minMeaningfulLen: 400` (a real body is substantial; the 404 is ~130 chars).
- **The academy-chapter probe (`coverage.spec.ts`).** In the academy cross-port follow, after the home renders,
  fetch the PUBLIC FS-derived `/catalog.json` for a real chapter slug (deterministic, no fragile DOM discovery;
  prefer a slug with an `it` overlay), load `/chapters/<slug>/`, assert the body, then re-load under `?lang=it` and
  re-assert — the second is the **§2 language proof** (the language switch on a chapter reader re-runs the resolver
  with locale `it` → pre-fix the same backend-null 404; post-fix the locale-aware FS body). Runs post-BFS, so the
  extra navigations don't disturb the crawl. Drift-guarded in `coverage-manifest.unit.spec.ts`.

## D4 — full `coverage.spec.ts` billion sweep run → M244 (Fate 2, already owned)
The §3 academy-chapter probe is unit-proven (139 e2e unit tests) + its live premise validated directly on billion
(catalog.json slug → chapter renders → Playwright innerText distinguishes the fixed chapter from the 404). Running
the FULL `coverage.spec.ts` sweep on billion (authenticated hero crawl → the academy cross-port follow that
invokes the probe) is a heavy live-browser run. **M244 "prove-on-billion" already owns it** — its exit gate (c)
is "the 39 live-browser specs execute green (T-3)", which includes `coverage.spec.ts`. So this is Fate 2: covered
by a future milestone of this release; no new deferral, no plan edit. Recorded here as the cross-milestone
confirmation (build-milestone three-fate rule).

## KB-1 — demopatch-spec.md patch inventory omits `academy-fs-published-fallback` (Phase 0b, YELLOW)
The KB-fidelity audit (`kb-fidelity-audit.md`, 2026-07-21, YELLOW) found `corpus/ops/demo/demopatch-spec.md`
inventories "11 patches … 1 × ant-academy" listing only `ant-academy-dev-origins`, while
`academy-fs-published-fallback` is a real, wired native-run patch (14 patch dirs on disk).
**Fate (revised at §4):** fully reconciled the inventory to the **15 manifests on disk** — not just the academy
portion. A directory-vs-table sweep found the table stale in TWO ways: (a) the 3 native-run `ant-academy` patches
(`ant-academy-dev-origins` + the 2 `academy-fs-published-*`), and (b) the 2 M232 `next-web-interview-flag-*` patches
(`packages/ui`) — both classes were missing. Since I was already asserting a count, shipping a partial "13" would
have been a knowingly-wrong number, so I set it to the true **15** (10 × next-web-app · 2 × app · 3 × ant-academy)
and added all missing rows. Also noted the standing hygiene gap: no directory-driven fence guards the patch-manifest
inventory (a `demo_knob_guard`-style sweep would prevent this class of drift). That fence is a tooling task beyond
M238's academy subject — **surfaced, not built here** (no milestone currently owns a demopatch-inventory fence; a
low-priority hygiene item, not release-scope-breaking).
