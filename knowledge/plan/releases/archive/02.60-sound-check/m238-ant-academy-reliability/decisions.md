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
**Fate update at close (D5):** the harden Pass 2 **BUILT the fence** (`demo-stack/tests/test_patch_inventory.py`) —
KB-1's "surfaced, not built" hygiene gap **converted deferred → LAND-NOW (done)**. No longer a deferral.

## D5 — Close-time deferral re-audit (Phase 1b): the standing-8 demo-stack failures re-fated Fate 2 → M244
The M238 harden full-sweep (779 demo-stack tests) re-surfaced **8 pre-existing failures**, all in files M238 never
touched → **0 M238 regressions**. The deferral re-audit (`audit-deferrals/deferral-audit-2026-07-21.md`, **YELLOW**)
re-fated every close-time deferral:

- **The standing 8** = the SAME set characterised at v2.5/M236 (`releases/archive/02.50-the-playbill/
  m236-prove-on-billion/rebaseline-standing-failures.md`, 2026-07-20) — composition UNCHANGED: `test_cockpit.py`
  ×6 (4 M234-academy-link-semantics + 2 M218 overlay-JS 30 s-window), `test_host_prereqs_m215.py` ×1 (M224 hiring
  port 13001 absent from `_UI_PORTS`), `test_purge.py` ×1 (macOS-environmental, self-documented). **All test-side
  debt, 0 product defects, 0 pin drift.** REPEAT/AGED-OUT (ridden M236→M237→M238; demo-stack test area touched by
  M238). **Re-fated Fate 2 → M244** (state.md standing backlog already owns it; M244 runs the full suite live).
  Landing them in M238 would be scope-bleed (unrelated to ant-academy). **Elevated note for M244:** these are cheap
  LAND-able TEST edits — discharge by EDITING the tests, not only "via a live bring-up"; they have now ridden three
  v2.6-adjacent milestones and M244 is the expiry point. **Fresh dated re-confirmation: 2026-07-21.** NOT
  release-scope-breaking → not an escape-hatch, no user wake (YELLOW, surfaced to orchestrator).
- **D4** (full `coverage.spec.ts` billion sweep) → **Fate 2 confirmed** — owned by M244 exit gate (c). No edit.
- **Inherited #4** (library-flash) → **Fate 2 confirmed still owned by M239** (open). No M238 action.

**Verdict:** YELLOW, `SEVERITY=warning` → close proceeds. Blocking items: none.

## D6 — Close code-review fates (Phase 2/2c): 1 inherited should-fix + 3 nice-to-have defensive edges
The cross-cutting code review found **0 must-fix**. The rext code-of-record is finalized + flake-gated + live-proven
and pinned at the hardened HEAD `3482a77` (the sole sanctioned close-time force op was the consumption-tag re-pin).
Re-opening it for the low-severity defensive edges below would exceed that op + require re-running the flake gate on
new code — disproportionate for theoretical/author-error edges in already-proven tooling. Fates:

- **[should-fix, INHERITED] Shared-clone concurrent revert.** All 3 native-run academy patches share the
  `stack-demo/ant-academy` clone (path is `demo-N`-independent); `ant-academy.sh N --stop` reverts the shared files
  unconditionally → tearing down `demo-1` while `demo-2` is live re-404s `demo-2`'s chapter route on its next HMR.
  **Pre-existing** (M238's body patch follows the pattern, doesn't introduce it); bites only on concurrent demos on
  one box. **Fate:** DOCUMENTED as a known limitation now (`frontend-tier.md` §"The BODY half"); the code fix
  (per-demo academy clone / applied-refcount before revert) → **standing backlog** (low-priority hygiene; not
  release-scope-breaking; no milestone owns demo-clone topology).
- **[nice-to-have] `firstChapterSlug` admits empty-string slug** → `/chapters//` theoretical false-pass. Unreachable
  from the real FS-derived `catalog.json` (an empty slug needs an empty chapter dir); the function already
  null-guards every other malformed shape. **Fate:** → next `stack-verify` rext build-iter (tighten to
  `.length > 0` + a unit case).
- **[nice-to-have] `mustNotInclude: ['']` footgun** — an empty forbidden entry makes `includes('')` always-true →
  unconditional `error`. Author-error only; no shipped descriptor triggers it. **Fate:** → same build-iter (guard/validate).
- **[nice-to-have] Non-atomic patch write** (`open("w")`, no `.tmp`→rename). Shared with both sibling `apply-*.sh`
  helpers; the clone is ephemeral so impact is bounded (worst case a re-clone). Fixing only M238's helper would break
  the byte-identical-siblings consistency the review praised. **Fate:** → the next `stack-injection` rext build-iter,
  fixed across all three helpers together (out of M238's academy-body scope).

None re-touch the tagged rext code-of-record at close. All are precisely recorded so they're findable by the owning
build-iter. Not disguised deferrals of the deliverable (the deliverable — #3 fix, #2 verdict, sweep — is DONE +
live-proven); these are polish on already-working code.

### Adversarial review (Phase 2c — scenarios considered)
Three non-trivial modules, one non-obvious failure mode each (recorded per the Phase 2c contract — the scenario, not
just the fix):

- **A — `apply-academy-fs-published-body.sh`, concurrent-teardown state.** demo-1 + demo-2 both up (idempotent
  second apply = correct no-op → both patched); `demo-1 --stop` reverts the **shared** clone file while demo-2 serves.
  The file-level guard ladder (G1–G7) is individually correct + idempotent, but there is **no cross-demo
  coordination** → **GAP (real, inherited)**. Same for all 3 native-run academy patches; M238 adds no new instance.
  Response: documented limitation + backlog (D6 above).
- **B — `firstChapterSlug`, input beyond bounds (empty-string slug).** `[{ slug: '', language: 'en' }]` → returns
  `''` (does NOT throw — the cross-port follow stays protected), yielding `/chapters//`, a theoretical false-pass on
  the catalog index. **GAP (low, unreachable from real data).** Every other malformed shape is handled + unit-fenced.
  Response: routed to next build-iter (D6).
- **C — `section-assert.ts` `mustNotInclude`, missing/empty list.** No `mustNotInclude` / `[]` → `(… ?? [])` →
  no-fail: **HANDLED** (unit-fenced + the shipped-descriptor floor-vs-`mustNotInclude` mutation-verify). The
  non-obvious edge — an empty-string *entry* `['']` → unconditional `error` — is a **GAP (footgun, author-error
  only)**. Response: routed to next build-iter (D6).
