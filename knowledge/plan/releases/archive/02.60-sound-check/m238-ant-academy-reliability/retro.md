# M238 — Retro

## Summary
M238 was the **first post-barrier fix** in v2.6 "sound check". The headline result: **one** chapter-body FS-published
demopatch (`academy-fs-published-chapter-body`, gated on the same `ACADEMY_DEMO_FS_PUBLISHED` as the M230 catalog
patch) fixed **BOTH** confirmed defects — **#3 (Start→404)** and **#2 (the "lesson in Italian errors" symptom)** —
because #2 was **not a distinct code bug**: locale is a `?lang=` query param (no `/[locale]` route, so a bare `/it`
404 is expected), the switcher is a sound EN↔IT toggle, and the one actionable defect (chapter-language 404) is the
SAME backend-null path as #3, so the #3 patch's locale-aware FS body fixes it. **Both proven live on `billion`**
(chapter 404→200; `?lang=it` 200). Also extended the academy coverage sweep (`ANT_ACADEMY_CHAPTER_SECTION` + a general
`mustNotInclude` negative marker + the catalog.json→chapter→`?lang=it` probe), **mutation-verified to go RED on a
broken academy**, and built a **directory-driven demopatch-inventory fence** (`test_patch_inventory.py`, exact 15 +
per-repo). Section milestone, closed-complete; **0 ant-academy platform-repo edits**.

## Incidents This Cycle
- **P3 — demopatch-spec.md inventory counts had drifted in three places (fixed at close, Phase 7).** The §5 header's
  "6 × `apps/web`" **overcounted** apps/web by 3 (the 3 `packages/core-js`/`packages/graphql` patches were mislabeled
  as apps/web); §4's "eight `next-web-app`" undercounted the real ten and named only 1 of the 3 native-run academy
  helpers; and an M224-era "distinct-manifest total unchanged at **11**" read present-tense, contradicting the
  M238-rewritten "15" header in the same section. All reconciled to the directory-fenced ground truth (repo-level pin
  `10 next-web-app · 2 app · 3 ant-academy` unchanged → `TestPatchInventory` intact). *This is exactly the drift class
  the new directory-driven inventory fence exists to prevent at the repo-level — but the finer human-readable
  sub-splits (apps/web vs packages/*) are prose the fence doesn't pin, so they still drifted.*
- **No product regressions, no flakes.** M238-touched suites 5/5 sequential (Python touched 183, TS unit 147);
  full demo-stack 778/786 — the 8 pre-existing standing fails only, **0 M238 regressions**. Harden: 0 bugs fixed
  inline (the subject was already correct; the value was deepened fences + the caught doc contradiction).

## What Went Well
- **One patch, two defects — because the triage was honest.** Rather than build a separate "language fix", M238
  *investigated* #2 on a fresh clone and found it was the same dead-backend path as #3 plus two non-defects (a mistyped
  `/it` non-route + a toggle mistaken for a broken menu). The deliverable shrank to one coherent FS-as-published
  behavior (grid + body, same env gate) instead of two.
- **The sweep is mutation-verified, not just asserted.** The shipped `ANT_ACADEMY_CHAPTER_SECTION` descriptor was run
  through the assertion engine across the patch-absent / patch-present states and *watched go RED then GREEN* — the
  release thesis ("a check that proves nothing is worse than no check") applied to the milestone's own guard.
- **The inventory fence closed a real standing gap.** The KB-fidelity audit surfaced that no directory-driven fence
  guarded the patch-manifest inventory; the harden BUILT one (`test_patch_inventory.py`), converting a "surfaced, not
  built" hygiene note into a landed guard in the same milestone.
- **Zero platform edits held.** A dead demo surface (backend-null chapter body) was fixed with a sha-pinned,
  revert-clean native-run demopatch on the ephemeral clone — the canonical `ant-academy` repo untouched.

## What Didn't
- **Nothing blocking.** The only close friction was the inherited doc-count drift (P3 above) and one inherited latent
  code limitation (the shared-clone concurrent-revert), both cheap: the docs reconciled at close, the code limitation
  documented + its fix routed forward. None escaped the milestone.

## Carried Forward
- **The 8 standing demo-stack test failures → M244** (Fate-2, M238-D5, fresh-dated re-confirmation). Identical to the
  v2.5/M236 re-baselined set, 0 M238 regressions. M244 should discharge by **editing the tests** (6 of 8 need no live
  stack), not only via a live bring-up. **Now ridden ≥3 v2.6-adjacent milestones — M244 is the expiry point.**
- **Full live `coverage.spec.ts` billion sweep → M244** exit gate (c) (Fate-2, M238-D4). Probe logic is unit-proven +
  its live premise validated on `billion`.
- **The shared-clone concurrent-revert code fix → standing backlog** (M238-D6): per-demo academy clone or an
  applied-refcount before revert. Documented as a known limitation in `frontend-tier.md`; bites only on concurrent
  demos on one box.
- **3 nice-to-have defensive edges → the next `stack-verify`/`stack-injection` rext build-iter** (M238-D6):
  `firstChapterSlug` empty-string guard, `mustNotInclude['']` footgun, non-atomic patch write (fix across all 3
  sibling helpers together). Not re-opening the finalized+tagged rext code-of-record at close.
- No escape-hatch deferrals. Deferral audit YELLOW (repeat/aging pattern surfaced, not a blocker).

## Metrics Delta (from metrics.json)
- **Python demo-stack (M238-touched):** 183/183 — `test_academy_fs_published_body.py` 18 (+2 harden Pass 2),
  `test_patch_inventory.py` 5 (NEW), `test_tooling.py` 160 (unchanged; R1 comment de-staled).
- **TypeScript e2e unit:** 147/147 (M238 added the `firstChapterSlug` matrix + `mustNotInclude` polarity + the
  shipped-descriptor mutation-verify); `tsc --noEmit` clean.
- **Go:** unchanged — M238 touched no Go (bash `apply-*.sh` + Python demo-stack harness + TS e2e only).
- **rext code-of-record:** tag `sound-check-m238-ant-academy-reliability` @ `3482a77` (re-pinned to the hardened HEAD
  at close, force-pushed — a milestone-internal consumption-tag advance).
- **Flake:** 0 (5/5 both suites). **Platform-repo edits:** 0. **Supply chain:** 0 net-new deps.
