# M237 — Progress

Section milestone. Checklist from the roadmap In-list. **All sections landed (Fate-1).**

## Sections

- [x] **Clone-freshness fix** (`ensure-clones.sh`) — fetch-verified freshness assertion (rc-checked, never suppressed-stderr) + opt-in advance-to-`origin/main`-or-pinned-tag (`DEMO_ADVANCE_CLONES`) + a real 7-state pin model (`clones.pin.json`; pinned vs stale-by-neglect distinguishable in `clones.lock.json`) + `DEMO_FRESHNESS_STRICT`. Tests `TestCloneFreshnessM237` (incl. 12-vs-202 anti-regression); shellcheck clean; demo-knob-guard green.
- [x] **F-M236-CLOSE-2** — R1 pristine sweep now globs `patches/*/*.yaml` (all 14), not the hard-coded 3. Tests `TestR1SweepM237`; harness refactored into `EnsureClonesHarness` mixin (no triple-run).
- [x] **Fresh-clone demo on billion** — re-pinned rext to `sound-check-m237-clean-stage`; §1 freshness + §2 R1(14) dogfooded live; verified-current clones (frontend 0-behind — the "202-behind" premise REFUTED).
- [x] **Confirmed-defect ledger** (decisions.md) — #1 menu **RESOLVED** (hierarchical for managers) · #4 library **does-not-reproduce** (populated, 0 errors) → M239 re-scoped · #2 academy language **SURVIVES** (`/it` 404 on the 5-behind empty academy) → M238. Verified via the e2e cockpit-login harness from the tailnet presenter vantage.
- [x] **Delivers** — `corpus/ops/rosetta_demo.md` (clone-freshness mechanism + 202-refutation correction) + `corpus/ops/demo/demopatch-spec.md` (R1 all-14) + `corpus/ops/demo/demo-up-defaults.md` (the two new knobs).

## rext code-of-record
Tag `sound-check-m237-clean-stage` (origin/main + tag pushed). Commits `8661444` (§1) + `7847473` (§2).

## Rosetta commits (branch `m237/clean-stage`)
`d6aadf1` (knob docs) · `c19c7a6` (ledger) · `230e247` (Delivers docs).
