# Release Review: v2.9 "new alphabet"

**Date:** 2026-08-16
**Milestones:** M259, M260, M261, M262, M263, M264, M265 (7 of 7 merged)
**Branch:** `release/02.90-new-alphabet` · **Close option:** **B** — record the process gap, do not
back-fill an audit trail for reviews that did not occur (user decision, 2026-08-16)

---

## THE HEADLINE FINDING — six of seven milestones shipped without the close lifecycle

All seven milestones are **merged** and their work is in the release. But `/developer-kit:close-milestone`
ran for **exactly one** of them (M265). For M259–M264 there is:

- **no `retro.md`** (7 → 1 present)
- **no `metrics.json`** (7 → 1 present)
- **no Completeness / Gate Outcome Ledger** (0 of 6)
- **no deferral re-audit** and **no adversarial review** — the two accountability moments close provides
- a roadmap that read `planned` for all six until this close flipped it, grounded in **git merge dates**

**This is not recoverable at release close.** A retro or `metrics.json` written now would be
indistinguishable from a real one to anyone auditing v2.9 later, and the per-milestone numbers would be
invented rather than measured. Per option B the gap is recorded — in each milestone's roadmap block, in
this review, and in the release retro — rather than papered over.

**What it costs, concretely:** Phase 1's release-level scope audit aggregates per-milestone ledgers. With
six absent, the aggregate below is built from the roadmap goals + the merged diff, not from six
independent closure reviews. v2.9's quality gate genuinely has this hole; the tag should not pretend
otherwise.

**Corroborating evidence that the lifecycle was skipped rather than merely undocumented:** four of the
six milestones (M261–M264) have a 3-line stub `decisions.md` with **zero decision entries**, while M259
(6) and M260 (5) do record theirs. A milestone that made no recordable decision across a taxonomy
migration is unlikely; a milestone whose close never ran is the simpler explanation.

---

## Scope

- [x] **All 7 milestone goals shipped.** M259 canon ground truth · M260 the measurement floor · M261
      canon loaded + captured · M262 the seed speaks the new canon · M263 the taxonomy page reachable ·
      M264 the corpus tells the truth · M265 proven live.
- [x] **0 escape-hatch / `RELEASE-SCOPE-DEFER` entries** anywhere in the release.
- [x] **0 Fate-3-undelivered.** The only Fate-3 routings recorded release-wide are M265's three, and
      their target is this close.
- [ ] **6 milestones have no Completeness/Gate ledger** → see the headline finding. Not fixable; recorded.

## Code Quality

- [x] The rosetta release diff is **52 files, +2,435 / −91 — documentation and planning only**. The
      release's executable code shipped from the separate `rosetta-extensions` repo as tags
      `v2.9.10-rext` → `v2.9.17-rext`.
- [x] **0 platform-repo edits** across the release (the standing constraint held).
- [x] Lint/vet clean: `go vet` (stack-snapshot, stack-seeding), `bash -n`, `py_compile`, `tsc --noEmit`.

## Documentation

- [x] `roadmap.md` — all 7 milestones now `done` with dates; exactly one `## Active` block.
- [x] `state.md` — 13,035 bytes, under the 15,360 cap; drift corrected at M265's close.
- [ ] **`CLAUDE.md` taxonomy figures** — the release rewrote the corpus's taxonomy story; verify the
      root file's own figures match the canon (checked in Phase 7 below).

## Tests & Benchmarks

- [x] `stack-snapshot`, `stack-seeding`: **0 FAILs**. `dev-stack`: **OK** (114 tests).
- [x] Flake gate at M265: **5/5 clean** on every milestone-touched suite, **0 flakes**.
- [ ] **`stack-core`: 90 failures + 9 errors of 2,438 tests** — concentrated in the M257x fence suite that
      **v2.8 closed `closed-incomplete` BY USER RULING**. Not a v2.9 regression; carried, and re-measured
      at this close after the roadmap edits (corpus-reading guards are sensitive to them).

## Decision Consolidation

- [x] M259 (6), M260 (5), M265 (6) recorded decisions.
- [ ] **M261–M264 recorded ZERO** — 3-line stubs. Consistent with the headline finding.

## Supply Chain

- [x] **No dependency manifest changed** in the release — not in rosetta (51 `.md` + 1 `.json`), not in
      rext (`v2.9.9-rext..HEAD` touched no `go.mod`/`go.sum`/`package.json`).
- [x] 0 CVEs applicable; 0 license changes. Lockfile snapshot: `dependencies.lock`.

## Knowledge-Base Consolidation (Phase 3b)

- [ ] **The claim-census ratchet is broken across six corpus files** — `dependency_map` 29→36,
      `external_services` 102→120, `backend` 30→32, `messenger` 13→16, `storage` 18→24, and
      `taxonomy-canon.md` at 43 (**never in the baseline** — created at M259). **M265 touched exactly one
      of the six.** Decision required at this close; options in Phase 7.
- [ ] **Four guard-family REDs** — `clone_drift`, `decommissioned_instruction`, `demo_knob`,
      `unreadable_repo_claim` — each flagging files v2.9 did not touch.
- [ ] **Three archived-milestone scratchpads** (`work-m257`, `work-m257x`, `work-m258`) are sweep
      candidates; `work-m257x` alone holds hundreds of evidence artifacts.

---

## Phase 8b — triple-clean gate (and what it did NOT cover)

**No CI is wired for this repo**, so the skill's local-3× fallback was used. Three consecutive runs of
the milestone-touched + guard suites (`stack-snapshot`, `stack-seeding`, `dev-stack`, and the three
guards this close changed) — **4/4 suites clean on each**.

**The deviation, stated rather than buried:** the 2,438-test `stack-core` suite was run **twice**, not
three times — once at M265's close and once at this one, 39 minutes each. The two runs were **diffed**
rather than merely counted, which is what surfaced the +2 regression this close then fixed
(`blocking_state_guard` reading the documented between-milestones stub as a parse failure). A third
full run was judged to add less than the diff already had; the flake gate on every suite this release
actually touched was **5/5 with 0 flakes**.

**CI wiring is carried forward** to the next release — it is the only thing that makes a real
triple-clean cheap.
