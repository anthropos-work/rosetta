# M266 — spec notes

_Technical notes accumulate here during the build. Section headers are derived from the milestone's
`In:` list; nothing below is decided yet._

## Pre-flight audits

_None yet._

## The one-file surface — `demo-stack/cockpit.py`

_Server-rendered Python f-strings; no template engine, no CSS file; one `_PAGE_CSS` at `:517`.
Re-resolve every anchor from `overview.md` § Open questions against the tree before editing._

## A1 — the half-row hero grid

### The emit sites (`:1365-1377`, `:1378-1383`) and `.hero` (`:543`)

_None yet._

### The manager full-bleed row (`grid-column:1/-1`) via `vantage_label`

_None yet._

### Test coupling — `test_cockpit.py:151` (count) and `:684` (anchor regex)

_None yet._

## A2 — the candidate label set

### `_badges()` (`:659-672`) and the `is_hiring` / `trajectory` inputs

_None yet._

### Manifest fields already present — `cockpit.go:246` / `:315` / `:232`

_None yet._

### The `--cand` token (`:527`) and the existing `av-candidate` precedent

_None yet._

## A3 + A5 — the `_content_tuple_row()` rewrite (`:1117-1173`)

### Removing `.ctcol` — CSS `:629-631`, markup `:1147-1153`

_None yet._

### Making the partition conditional (`:1136-1141`)

_None yet._

### Where the pass/fail signal goes instead

_None yet._

## A5 (data half) — `has_verdict` on `contentProductMeta`

### The field + registry (`content_manifest.go:149-157`, `:163-193`)

_None yet._

### The honesty gate — regenerate `presets/content-manifest.json` + re-pass `CanonicalFileMatchesProjection`, same commit

_None yet._

## A4 — inline-SVG language flags

### `_LANG_PILL_LABEL` / `_LANG_TOGGLE_LABEL` (`:957-958`)

_None yet._

### Why not FontAwesome free (no country flags) and not emoji (Chrome/Windows)

_None yet._

## Corpus revisions

### `cockpit-spec.md` — the M43/M242 card contract

_None yet._

### `content-stories-spec.md` — §7.2 columns, §7.6 language labels, §4 honesty gate

_None yet._
