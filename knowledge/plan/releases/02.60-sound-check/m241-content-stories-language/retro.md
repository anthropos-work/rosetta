# M241 — Retro

## Summary
M241 was the **fourth post-barrier fix** in v2.6 "sound check" (content-stories language), a `section` milestone that
**opened with a read-only prod pool-count go/no-go query** (user decision 3, R2). It gave each content-story session its
**real played language**. The pool query returned **GO** and immediately surfaced the core defect: **11 of the 13 pinned
sessions were actually Italian**, yet `content_stories_write.go` hard-coded every clone's `sessions.language` to
`english`. **§2** added `s.language` to the `sourcing.go` SELECT + an optional language filter, a `Language` field on the
fixture + the `content_manifest.go` projection, and flipped the write to **`cs.Language`**. **§3** sourced **10 EN/IT
counterparts** (fixture **13 → 23**, denominator **29 → 49**), so 11 of 12 requirement tuples carry both languages;
**INTERVIEW is Italian-only** (EN interview passes all out-of-band, EN interview fails = 0 — R2 realized, the
EN-only-fallback-per-tuple decision). **§4** the fail-closed **`ValidateLanguageConsistency`** gate (a `lang_toggle`
disagreeing with its own coverage FAILS `--content-export`) + a TS mirror, with teeth. **§5** the cockpit **EN|IT
segmented toggle** (`_LANG_JS`, raw-string injection-free; per-row language pill; byte-clean when the manifest has no
language axis). Delivered docs: `content-stories-spec.md` (§2/§4/§7.6) + `session-clone-spec.md` §2.1. Clean complete
5-section close; **0 platform-repo edits**; assertions are STRUCTURE / PRESENCE / language-LABEL only, never a translated
value (P2 copy-immunity).

## Incidents This Cycle
- **None.** The close's scope + code-quality + adversarial + test reviews found **0 fix-required findings**; the one
  documentation fix (a stale `29 → 49` denominator line) was cosmetic drift, not a defect. Flake gate **5/5** all three
  stacks (Go shuffle, TS, Python — the 6 pre-existing cockpit failures are deterministic 5/5, not flaky).
- **No regressions.** The full demo-stack Python suite ran **808 pass / 9 fail** — the 9 are the identical standing set
  (6 academy+overlay `test_cockpit.py` + `test_host_prereqs_m215` + `test_purge` + `test_reap` reap-17700), **0 new from
  M241** (matches the M239-close 9-fail baseline).

## What Went Well
- **The pool-query-FIRST discipline paid off exactly as designed.** Running the read-only counts before any code both
  cleared the go/no-go (R2 interview scarcity) *and* exposed the real defect that no one had noticed: the seeder was
  mislabeling 11 of 13 sessions. The measurement drove the scope, not the other way around.
- **The harden caught the CORE-BUG write-side gap — the exact failure class v2.6 exists to kill.** The build shipped the
  `cs.Language` write but **no test asserted the seeded `sessions.language` column carried it**, so reverting the write to
  the pre-M241 hard-coded `english` passed **every** Go suite (proven: mutated → all green). `TestContentStorySeeder_
  WritesRealLanguage` closes it (column resolved by NAME, asserts the real language + a valid non-blank label + that the
  set spans BOTH languages so it can't pass vacuously). A "sound check" milestone catching its own silent-pass gap is the
  release thesis working on itself.
- **Fail-closed on BOTH sides.** `ValidateLanguageConsistency` (Go, wired into export) + the TS canonical mirror refuse a
  manifest whose `lang_toggle` disagrees with its coverage — a toggle that would swap to an unseeded language is caught at
  export, and a language drift that slips one side is caught by the other. Both mutation-verified to bite.
- **PII discipline held.** Customer content never entered agent context; only scrubbed fixtures + the read-only counts-only
  pool query; every assertion is structure/presence/label — the translated copy itself is never asserted (P2).

## What Didn't
- **The build under-tested its own headline fix.** The whole milestone is about writing the *real* language, yet the
  build's own tests didn't fence the write column — the guard came only at harden. Not a shipped defect (harden closed it
  before close), but a reminder that a fix's regression test belongs in the section that lands it, not the hardening pass.
- **The close had to disposition inherited test noise.** 6 of the 9 standing demo-stack failures surface through
  `test_cockpit.py` (the file M241 touched), so every M241 run shows red that isn't M241's — the close spent effort
  re-confirming provenance (identical at base `ae0e869`, 0 new). The standing debt has now ridden ≥4 v2.6 milestones;
  **M244 is the named expiry** and should discharge it by editing the tests.

## Carried Forward
- **None new from M241** — a clean complete section close, 0 new deferrals.
- **(Inherited, confirmed) → M244:** the **9 standing demo-stack test failures** (Fate-2, M238-D5 / M239-D13 reap-17700) —
  6 surfaced here via `test_cockpit.py` — already homed; M244 is the named expiry (discharge by editing the tests).

## Metrics Delta
- **Tests:** rext whole-repo Go test funcs 1999 (M240) → **2005** (release-cumulative); M241 added **6** (the 4 build-time
  language-axis tests + the 2 harden write-side regression/fallback tests). `test_cockpit.py` 142 (136 pass / 6
  pre-existing). TS **151** unit specs. Full stack-seeding module GREEN. Flake **0**.
- **Coverage (touched language fns):** `SelectionSQL` 100%, `bilingualTuples` 100%, `resolveSession` 100%,
  `ValidateLanguageConsistency` 96%, `seedContentStorySession` 91.3% (the language write now covered).
- **Deferral audit:** YELLOW (0 new M241 deferrals; the standing debt = tracked repeat, Fate-2 → M244, named expiry).
  0 RED blockers. **Platform-repo edits:** 0. **PII into context:** 0. (Full machine-readable delta: `metrics.json`.)
