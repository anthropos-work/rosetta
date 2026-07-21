# M238 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **#3 Start→404** — wire a chapter-body demo path. **DONE:** the FS-as-published `academy-fs-published-chapter-body` demopatch (rext `demo-stack/patches/` + `apply-academy-fs-published-body.sh` + `ant-academy.sh` wiring), gated on the same `ACADEMY_DEMO_FS_PUBLISHED` env var as the catalog patch. 18 tests green (16 at §1 + 2 harden Pass 2 revert-drift/locale). **Proven live on billion**: a chapter 404→200 (real title/body). (decision D1)
- [x] **#2 language error** — **DONE (verdict, not a code fix):** NOT a distinct code bug. Locale is a `?lang=` query param (no `/it` route); the switcher is a sound EN↔IT toggle; the chapter-language 404 is the SAME backend-null path as #3, **fixed by the #3 patch** (locale-aware FS body). Confirmed fresh + live on billion (`?lang=it` chapter → 200). (decision D2)
- [x] **Academy presence/coverage sweep** — extend to assert chapter-body render. **DONE:** `ANT_ACADEMY_CHAPTER_SECTION` + a general `mustNotInclude` negative marker + the academy-chapter probe (catalog.json slug → chapter → `?lang=it` re-render) in `coverage.spec.ts`; 139 e2e unit tests green; live premise validated on billion. Full billion sweep run → M244 (Fate 2, D4). (decision D3)
- [x] **Delivers** — `corpus/services/ant-academy.md` (chapter-body demo path + #2 language verdict) + `corpus/ops/demo/frontend-tier.md` (the BODY half) + `corpus/ops/demo/demopatch-spec.md` (KB-1: the 2 academy-fs-published rows, inventory 11→13) + `corpus/ops/demo/coverage-protocol.md` (the chapter-body + language sweep extension) updated.

## Live verify (billion, 2026-07-21)
#3 + #2 proven on demo-1: `/chapters/<slug>/` 404→200 (title flips "Not Found"→real chapter), `?lang=it` → 200, `/it` a non-route (expected). See spec-notes.md § LIVE VERIFY.

## Audits
Phase 0b KB-fidelity: **YELLOW** (`kb-fidelity-audit.md`) — central premise ALIGNED; one incidental gap KB-1 (fixed in §4).

## M238: Hardening

Code-of-record in **rosetta-extensions** (authoring copy, `main`); harden commits land there + the doc reconcile lands on the rosetta `m238` branch. Note for close: the consumed tag `sound-check-m238-ant-academy-reliability` should be re-pinned to the hardened rext HEAD (the M237 precedent — moving a consumed tag is a close-time decision).

**Scope manifest (M238-touched, `533c489..HEAD` in rext):** `ant-academy.sh` (launcher wiring, static-fenced), `demo-stack/patches/academy-fs-published-chapter-body/*.yaml` + `stack-injection/apply-academy-fs-published-body.sh` (the demopatch + helper — `test_academy_fs_published_body.py`), `stack-verify/e2e/{lib/section-assert.ts, lib/coverage-manifest.ts, tests/coverage.spec.ts, tests/coverage-manifest.unit.spec.ts, tests/section-assert.unit.spec.ts}` (the coverage sweep — its two unit specs).

**Coverage measure:** qualitative / behavior-path (the scope manifest — every touched function has a behavior or static-fence test). This rext tooling repo has **no line-coverage tool wired** and its convention is targeted behavior + static-fence tests (the `test:coverage` script is the *demo*-coverage sweep, not code coverage); instrumenting nyc/coverage.py is a repo-wide toolchain decision beyond M238's academy-reliability scope — not a silent skip, a convention-consistent choice (same as the M237 harden).

### Pass 1 — 2026-07-21 (the release-thesis work: mutation-verify the sweep goes RED on a broken academy)
**Tests added (TS, 64 → 72 unit specs green; `tsc` clean):**
- `coverage-manifest.ts` / `coverage.spec.ts`: extracted the untested `firstChapterSlug` slug-selection logic into the lib (the M211 `isAcademyPort` precedent — ONE shipped copy the sweep + its drift guard both consume).
- `coverage-manifest.unit.spec.ts`: +5 `firstChapterSlug` tests — PREFERS an it-overlay slug (the #2 proof), falls back to first-EN, treats missing `language` as en, filters non-string slugs, returns null (never throws) on empty/malformed/it-only.
- `section-assert.unit.spec.ts`: +3 mutation-verify tests running the **SHIPPED** `ANT_ACADEMY_CHAPTER_SECTION` (floor 400) through the assertion engine across the #3 states — patch ABSENT (short 404 → RED on the floor; padded 404 → RED on mustNotInclude, naming the phrase) vs patch PRESENT (a real ≥400-char body → GREEN). The prior mustNotInclude tests used an inline descriptor (floor 8); these exercise the exact shipped descriptor the live sweep asserts.

The full live billion sweep stays **M244**'s exit gate (c) (Fate 2, D4); Pass 1 is the unit-altitude proxy for "disable the patch, confirm red, restore." (rext `c10413d`)

### Pass 2 — 2026-07-21 (demopatch deepening + the directory-driven inventory fence)
**Tests added (Python, chapter-body 16 → 18; +5 inventory; all green):**
- `test_academy_fs_published_body.py`: +revert **DRIFT-REFUSE** (a file neither patched nor pristine must REFUSE exit-7 + be left byte-untouched — revert-clean is load-bearing, the F-M236-CLOSE-2 stranded-patch class) + **locale threading** (the shipped replacement passes BOTH slug AND locale to `safeLoadFsShape` → the #2 `it` request resolves the `it/` overlay, not an EN-only body).
- `test_patch_inventory.py` (**NEW**): the directory-driven exact-inventory fence — enumerates every `patches/<name>/<name>.yaml`, validates each via `manifest_loader` (scope=demo + id==dirname), and pins the EXACT total (**15**) + per-repo breakdown (`10 next-web-app · 2 app · 3 ant-academy`) against demopatch-spec.md §5. **Mutation-proven RED** on an added patch and on a mis-filed repo. Complements `TestR1SweepM237`'s ≥14 R1 floor with an exact + per-repo pin (the M237 R1 directory-driven precedent). This closes the build's D-KB-1 "standing hygiene gap — no directory-driven fence guards the patch-manifest inventory."
- `test_tooling.py`: de-staled the R1 fence's "all 14 today" comment (15 today; the ≥14 floor stays the F-M236-CLOSE-2 baseline).

**Bugs fixed inline:** none (no new test exposed a production bug — the subject was already correct; the value is the deepened fences + a caught doc contradiction).

**Flakes stabilized:** none observed. Flake gate: Python 23-test suite + TS 72-spec suite each ran **3× consecutively clean**.

### Knowledge backfill
- `corpus/ops/demo/demopatch-spec.md` reconciled (rosetta `b485f7c`): §5 "standing hygiene gap" note → the fence is **landed**; §2.1 present-tense "sweeps ALL 14 manifests" → "every manifest on disk (15 today, directory-driven)"; §2.1 + §6 now reference `TestPatchInventory`. **Also fixed a live contradiction M237's close left**: §6 step 7 still told authors to "register the manifest in R1's `PATCH_MANIFESTS` array" and claimed "no test fences the array" — but M237 removed that array (R1 is directory-driven). Rewritten to match §2.1. (The historical billion "swept 14" quote + the M237-era "14 total / 11 uncovered" narrative are left as accurate history; §5-bis line-230 "distinct-manifest total: 11" is a specialized hiring-build sub-count, not the inventory total — left untouched, out of M238 scope.)

### Stop condition
Two passes, then stabilized: the six-dimension Step-2b re-scan found no new high-value **unit** gap (the remaining candidates — `probeAcademyChapter`'s I/O orchestration + the `onCrossPortFollow` academy glue — are heavy Page-mock surfaces whose pieces are already covered by the Pass-1 tests + the live billion premise, and the full live sweep is M244/Fate-2). No flakes across 3× runs. `tsc` clean.

### Pre-existing baseline note (out of M238 scope — for close)
A full `demo-stack/tests` run (779 tests) surfaced **8 pre-existing failures, all in files M238 never touched**, so they are NOT M238 regressions and NOT harden-scope: `test_cockpit.py` ×6 (the M218 overlay-JS rewrite removed the "30000" 30s-window; the M234 cockpit academy-link render — stale-test drift), `test_host_prereqs_m215.py` ×1 (the serve-reset generator now emits port `13001`/hiring — M224 — but the test's `_UI_PORTS` wasn't updated), `test_purge.py` ×1 (self-documented macOS-environmental — the container-owned-0700 bug can't reproduce on Docker Desktop). Surfaced for close/re-fate; the M238-touched suites are all green.

## M238: Final Review

Close review (2026-07-21). Three parallel scans (code / docs / full test run) + the Phase 1b deferral re-audit.
Test run at close: **Python demo-stack 778/786** (the 8 known standing fails; **0 M238 regressions** — touched
files 183/183: `test_academy_fs_published_body.py` 18 · `test_patch_inventory.py` 5 · `test_tooling.py` 160);
**TS e2e unit 147/147**; **`tsc` clean**.

### Scope
- [x] All 4 In-list items delivered (#3 D1, #2 verdict D2, sweep D3, Delivers docs). No gaps; 0 TODO/FIXME in touched code.

### Code Quality
- [x] [should-fix, INHERITED — not an M238 regression] Shared-clone concurrent revert: all 3 native-run academy patches share `stack-demo/ant-academy`; a `demo-1 --stop` reverts the shared file while `demo-2` is live → documented as a known limitation (`frontend-tier.md`) + code fix routed to standing backlog (D6).
- [x] [nice-to-have] `firstChapterSlug` admits empty-string slug (theoretical false-pass); [nice-to-have] `mustNotInclude: ['']` footgun (author-error); [nice-to-have] non-atomic patch write (shared w/ siblings) → the 3 low-severity defensive edges in M238's own new rext code routed to the next `stack-verify`/`stack-injection` rext build-iter (D6) — NOT re-opening the finalized+flake-gated+live-proven code-of-record (tag pinned at hardened `3482a77`) for theoretical edges.
- [x] Adversarial scenarios A/B/C recorded in `decisions.md` (D6 § Adversarial review).

### Documentation
- [x] [should-fix] `demopatch-spec.md` §5 header "6 × apps/web" overcounted apps/web by 3 → corrected sub-split (3 apps/web · 2 apps/hiring · 2 packages/ui · 2 packages/core-js · 1 packages/graphql = 10; repo-level 10 next-web-app unchanged, test-pin intact).
- [x] [should-fix] `demopatch-spec.md` §5-bis stale present-tense "11" contradicted the 15 header → scoped as the M224-era figure + cross-ref the directory-fenced 15.
- [x] [nice-to-have] Added `(#M238-D1/D2/D3)` back-ref tags to the primary mechanism mentions.

### Tests & Benchmarks
- [x] No coverage gaps — touched surfaces have behavior/mutation/static-fence tests (2 harden passes). No benchmark surface (docs+tooling milestone). Handbook/progress count reconciled (§1 "16"→18).

### Decision Triage
- [x] D1/D2/D3 confirmed already-blended into `ant-academy.md`/`frontend-tier.md`/`coverage-protocol.md`/`demopatch-spec.md` (build §4) + ref-tagged. KB-1 fence landed (harden). D4/D5/D6 archive (planning) — stay in `decisions.md`.
