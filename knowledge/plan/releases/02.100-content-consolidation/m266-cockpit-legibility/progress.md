# M266 — progress

**Status: PLANNED. Not started.**

## Section checklist

- [ ] **A1 — half-row heroes.** Grid-wrap the hero cards emitted at `cockpit.py:1365-1377` / joined at
      `:1378-1383`; `.hero` (`:543`) stops being full-width by construction; the manager card gets
      `grid-column:1/-1`, split on the manifest's existing `vantage_label == "MANAGER"`.
      **Both test couplings intact** — `test_cockpit.py:151` (`class="btn login"` count == n_heroes) and
      `:684` (the `<a class="btn login" href=... data-login-as=...>` regex).
- [ ] **A2 — candidate labels.** `_badges()` (`:659-672`) reads `is_hiring` and relabels the hiring
      vantage (candidate / performing / under-performing). **ZERO Go change, ZERO re-seed** —
      `cockpit.go:246` (`IsHiring`, from `st.IsHiringOrg()` at `:315`) and `:232` (`Trajectory`) already
      carry it; `--cand` already exists at `cockpit.py:527`; `stories.seed.yaml:247` **untouched**.
- [ ] **A3 + A5 (render) — one rewrite of `_content_tuple_row()` (`:1117-1173`).** `.ctcol` gone
      (CSS `:629-631`, markup `:1147-1153`); the partition at `:1136-1141` no longer unconditional.
- [ ] **A5 (data) — `has_verdict` on `contentProductMeta`** (`content_manifest.go:149-157`, registry
      `:163-193`).
- [ ] **A5 (honesty gate) — canonical `presets/content-manifest.json` regenerated and
      `CanonicalFileMatchesProjection` re-passed IN THE SAME COMMIT.**
- [ ] **A4 — inline-SVG language flags** on `_LANG_PILL_LABEL` / `_LANG_TOGGLE_LABEL` (`:957-958`).
      No FontAwesome (no country flags in the free set), no emoji (does not render on Chrome/Windows);
      the panel stays stdlib-only and self-contained.
- [ ] **Open questions resolved in `spec-notes.md`** before the markup is written (label set, verdict
      membership, where the pass/fail signal goes, which English flag, odd-count/no-manager cases).
- [ ] **Line anchors re-resolved** against the tree at milestone start.
- [ ] **Corpus — `corpus/ops/demo/cockpit-spec.md`** revised (the M43/M242 card contract).
- [ ] **Corpus — `corpus/ops/demo/content-stories-spec.md`** revised (§7.2 columns, §7.6 language
      labels, §4 honesty gate).
- [ ] **Zero platform-repo edits confirmed** — every touched file rext-owned or corpus.

## What was done

_Nothing yet._
