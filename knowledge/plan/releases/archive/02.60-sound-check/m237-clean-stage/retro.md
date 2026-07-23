# M237 — Retro

## Summary
M237 was the **HARD go/no-go barrier** that opens v2.6 "sound check". It fixed the two clone-freshness findings
carried out of the v2.5 close: **F-M236-CLOSE-1** (a **fetch-verified** freshness assertion in
`ensure-clones.sh` — rc-checked fetch, stderr never suppressed, a real **7-state pin model** in `clones.lock.json`,
opt-in `DEMO_ADVANCE_CLONES` + `DEMO_FRESHNESS_STRICT`) and **F-M236-CLOSE-2** (the R1 pristine sweep now
**directory-driven** across all 14 manifests, was a hard-coded 3). Both were **dogfooded live on `billion`**. The
§3/§4 **confirmed-defect ledger** re-triaged the ambiguous v2.5 UI defects on a correctly-built demo — and its
headline finding **refuted the release's own premise**: billion's clones were **0–2 commits behind** (frontend
current), not the "~202 behind" the design assumed; the "202" reading was itself the suppressed-fetch artifact §1
eliminates. #1 (menu) RESOLVED, #4 (library) does-not-reproduce → M239, #2 (academy language) SURVIVES → M238.
Section milestone, closed-complete; **0 platform-repo edits.**

## Incidents This Cycle
- **P2 — the test harness had the milestone's own defect class (fixed inline, harden Pass 3, rext `533c489`).** The
  ensure-clones test sandbox constrained `PATH` to a stub bindir but never symlinked `grep`/`sed`, so phase (c-pre)
  (the repos.yml stub-sweep, a `grep … | sed …`) printed `grep: command not found` and ran **INERT in every test** —
  a sandbox that "runs the script" while a whole phase proves nothing is the exact clean-stage defect class, turned
  inward. Wired `grep`+`sed`, gave the `cms` fixture a real `.git`, and locked it with a c-pre sweep test.
- **P2 — a log line claimed a step was skipped while it ran (fixed inline, harden Pass 2, rext `2ef8f43`).** The
  "demopatch not found" branch logged `skipping R1/R2`, but R2 (no-push) + R1b run unconditionally after it — a
  confident-wrong log, precisely the class this release exists to kill. Message now names only R1. **This fix was
  the reason the consumption tag had to be re-pinned at close** (7847473 → 533c489), so consumers get the honesty fix.
- **P3 — a stale doc cross-reference (fixed at close, Phase 7).** M237 renamed the `rosetta_demo.md` §"Clone freshness"
  heading and corrected the 249/202-behind figures, but the sibling `frontend-tier.md` callout still linked the old
  anchor and cited the refuted figures as current fact ("that half is open"). Swept and reconciled.
- No regressions; no flakes (5/5 on the M237-touched suites); full `test_tooling.py` 160 green.

## What Went Well
- **The barrier proved its own worth.** The §1 fetch-verified method contradicted the old suppressed-fetch method on
  the very first live run (202-behind → 0-behind), and the refutation was independently confirmed by raw `git rev-parse`.
  A fix that immediately falsifies the assumption that motivated it is the strongest possible validation.
- **Honesty invariants are mutation-verified, not just asserted.** The 12-vs-202 anti-regression (a behind-count off a
  failed fetch is `null`, never fabricated) and the log-honesty fix were each proven by watching a test go RED when the
  guard is removed — twice over across the harden passes.
- **Scoped-as-a-barrier held.** M237 fixed the build-freshness mechanism and *re-triaged* the defects; it did NOT reach
  downstream to fix them. The survivors routed cleanly to owners that already held them (Fate-2), so the fan-out is
  scoped against reality with zero plan churn.

## What Didn't
- **Nothing blocking.** The friction was self-inflicted-then-caught: two harden-pass bugs in the test infra / log
  strings (both fixed inline, both mutation-verified) and one stale sibling doc at close. All cheap, all landed —
  none escaped the milestone.

## Carried Forward
- **#2 academy language-switch → M238** (Fate-2; already in M238's `In:` list — "Fix #2 (language error — re-triaged in
  M237)"). M238 must re-verify on a **fresh (advanced) academy** — billion's academy is 5-behind; the M237
  `DEMO_ADVANCE_CLONES` tooling is the lever.
- **#4 library empty-first-load → M239** (Fate-2; already in M239's `In:` list), **re-scoped down** from "library is
  empty" to "verify there is no cold-first-load empty flash / client-fetch race" — the library itself renders populated.
- **ant-academy 5-behind clone → M238/M244** — advance it with `DEMO_ADVANCE_CLONES` on the next billion rebuild.
- No deferrals (deferral audit GREEN, 0 items). No escape-hatch.

## Metrics Delta (from metrics.json)
- **Python demo-stack `test_tooling.py`:** 146 → **160** (+14). M237-touched surface (CloneFreshness/R1Sweep/Functional):
  **46** tests, all pass.
- **Go:** unchanged — M237 touched no Go (bash `ensure-clones.sh` + its python subprocess harness only).
- **Fences:** shellcheck rc0; demo_knob_guard OK both-directions (29 env knobs incl. the 2 new). **Flake:** 0 (5/5).
- **rext code-of-record:** tag `sound-check-m237-clean-stage` @ `533c489` (re-pinned to the hardened HEAD at close).
- **Platform-repo edits:** 0. **Supply chain:** 0 net-new deps.
