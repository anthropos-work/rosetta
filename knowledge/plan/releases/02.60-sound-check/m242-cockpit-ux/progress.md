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
