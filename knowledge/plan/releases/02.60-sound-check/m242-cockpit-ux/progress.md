# M242 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **(1) row layout** — regroup by requirement tuple `(sim_type, modality)` → `target | passed | not-passed` on one row (render-only) — rext `df82fea` (+`bcbec32` review fix)
- [x] **(2) tab selector** — move into the white header, right, vertically centered; preserve the byte-identical-when-no-content-manifest invariant — rext `a5f96e8`
- [x] **(3) hero icon bg by user-type** — manager=orange / employee=indigo / candidate=teal (net-new `--cand`); manager-wins order — rext `ccc13f3`
- [x] **Extend the cockpit specs** — `cockpit-spec.md` (header layout + role-color + regroup) + `content-stories-spec.md §7.2/§7.6` (tuple-regrouped row model)
- [x] **Delivers** — `corpus/ops/demo/cockpit-spec.md` + `corpus/ops/demo/content-stories-spec.md`

## Result
- All 3 UX changes landed, render-layer only (no data/seed change, no manifest schema change, zero platform-repo edits).
- Tests: **152 pass / 6 standing** in `demo-stack/tests/test_cockpit.py` — **0 NEW failures**. +16 net-new M242 tests (5 tuple-regroup, 3 header-placement, 8 avatar-by-role). 3 stale `.stitle` title asserts updated to `.cttitle` (Fate-1).
- The 6 standing failures (4 academy-link + 2 overlay-JS) are outside M242's surface — left to M244 (see `decisions.md` D7).
- KB-fidelity Phase 0b: **GREEN** (`kb-fidelity-audit.md`).
- rext authoring copy: 4 commits `df82fea` → `a5f96e8` → `ccc13f3` → `bcbec32` on `main`, tagged (see close report).

## M242: Hardening

### Pass 1 — 2026-07-22
**Method:** mutation verification (stdlib-only — the demo-stack section takes no third-party dev deps, so
coverage was read via stdlib `trace`, not coverage.py; a temp-dir mutation harness kept the real tree
untouched). Every M242 render branch was mutation-tested against its target test.

**Mutation results (all four prompt priorities):**
- **Byte-identical-no-content invariant — STRONG.** Breaking the no-content header branch to emit the new
  flex header turns `TestHeaderTabPlacement::test_no_content_header_is_byte_identical_to_today` RED. The
  test is a true verbatim byte-for-byte pin (not a substring) — it already had teeth; left as-is.
- **Role-color derivation — STRONG.** Manager-wins precedence (candidate-before-manager mutant → RED),
  employee-default (mistint mutant → RED), and the candidate branch are all caught.
- **Row-regroup — TWO toothless tests found + closed:**
  1. `test_same_tuple_pass_and_fail_share_one_row_in_two_columns` proved presence + pass-before-fail order
     but NOT column placement — a mutant filing the FAILING session into the PASSED column SURVIVED.
     Strengthened to pin the verdict→column mapping (pass CTA between the headers, fail CTA after the
     not-passed header). Mutant now RED.
  2. The non-sim tuple TITLE label escaping was unpinned — a mutant dropping `html.escape` on the free-text
     `label` SURVIVED (the only test hitting that path used benign labels). New escaping test added; RED.
- **Two untested branches closed:** the symmetric only-failing "No passing run" empty marker (the
  only-passing fixture never reached it) and the defensive presence-session-inside-a-verdict-group slot
  (`cockpit.py:620`, "nothing silently dropped").

**Tests added:** `test_cockpit.py`: 4 unit (only-failing marker, presence-in-verdict-group, non-sim title
escape, tuple icon escape) + 1 existing test strengthened (column placement). 158 → 162.

**Bugs fixed inline:** none — both mutation survivors were TEST weaknesses, not production bugs (cockpit.py
render logic was correct; the tests just couldn't see a break). No production code changed this pass.

**Coverage delta (branch-level, milestone-touched fns):** closed the 2 genuinely-uncovered real code
branches (`_content_tuple_row` presence-in-verdict at :620; the only-failing `pass_body` empty ternary
branch). Remaining `trace` "misses" are docstring / f-string-continuation / bare-`else:` artifacts.

**Flakes stabilized:** none observed (3 consecutive clean sequential runs of the touched classes).

**Knowledge backfill:** no KB-worthy findings — the surfaced facts (verdict→column mapping, the symmetric
empty markers, the AA teal) are already in `content-stories-spec.md §7.2` + `cockpit-spec.md` + the code
comments; the survivors were test-quality gaps, not new system behavior.

### Pass 2 — 2026-07-22
**Tests added:** 1 unit — `test_candidate_avatar_meets_aa_contrast` pins priority-3's accessibility claim:
parse the rendered `--cand`/`--cand-soft` and compute the WCAG ratio (4.86:1), assert ≥ 4.5 (AA-normal).
A bare hex-string assert cannot see an inaccessible palette change; this can (mutation-verified — a
pale-on-pale teal → RED). 162 → 163.

**Bugs fixed inline:** none. **Flakes:** none. **Knowledge backfill:** none (the ~4.8:1 AA claim is
already stated in the `--cand` CSS comment + the role-color doc section; now regression-guarded).

### Stop condition
Loop stopped after Pass 2: the full six-dimension scan found nothing further worth adding (the only
remaining candidates are low-value CSS-string change-detectors for the flex-gap review polish — style,
not behavior — deliberately not pinned); branch coverage on the touched functions is complete; 0 flakes.
Final: **157 pass / 6 standing** (`test_cockpit.py`), **+5 net-new harden tests** (162→…→163 total; 158
pre-harden), **0 NEW failures** (the 6 standing = academy-link + overlay-JS, owned by M244). rext harden
commits: `db9e229` (pass 1) + `b0e76b3` (pass 2) on `main` — **the tag `sound-check-m242-cockpit-ux`
must be re-pinned to the hardened HEAD at close** (M237–M241 precedent).
