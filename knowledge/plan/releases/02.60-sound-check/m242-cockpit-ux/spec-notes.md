# M242 — Spec notes

Topic → doc → code triples + cockpit-render findings accumulate here during build.

## (1) Row layout — regroup by requirement tuple
- Regroup by `(sim_type, modality)` → `target | passed login options | not-passed login options` on one row (render-only; fields exist).

## (2) Tab selector — into the white header
- Move into the white header, right, vertically centered (restructure `cockpit.py` header to flex).
- Preserve the byte-identical-when-no-content-manifest invariant.

## (3) Hero icon bg by user-type
- manager = orange / employee = indigo (reuse the badge palette); derive a candidate color = `is_hiring && vantage != manager`.

## Pre-flight audits — (1) row layout (session 1, first section)
- `/developer-kit:audit-kb-fidelity --milestone=M242` → **GREEN**. Report: `kb-fidelity-audit.md`.
- rext authoring-copy sha at audit: `17beede`.
- Topic → doc → code triples (audit reuse baseline for §2/§3 — same subsystem, one file):
  - Content-tab row render → `content-stories-spec.md §7.2/§7.6` → `cockpit.py::render_content_tab / _content_session_row / _content_session_actions`.
  - Cockpit header + tab selector → `cockpit-spec.md §"The 2nd tab"` → `cockpit.py::render_page` header (~L782) + `_TAB_JS`.
  - Hero icon bg by user-type → `cockpit-spec.md §"The UI surface"` → `cockpit.py` `.hero .avatar` CSS (L144) + `_badges` (L239) + `hero.is_hiring` (L715).
- Incidental (M244, not M242): `cockpit.py` module docstring L40-46 still narrates the removed per-hero `[Academy]` link.

## Test-staleness inventory (the 6 pre-existing fails — all OUTSIDE M242's 3 surfaces → M244)
- `TestAcademyLink::test_academy_link_renders_per_hero_when_base_set`, `TestAcademyCatalogEntryEdges::test_render_academy_entry_fields_are_escaped`, `TestAcademyCatalogEntryEdges::test_render_defaults_academy_path_persona_label_when_absent`, `TestServedPanelWithAcademy::test_root_serves_academy_link` — assert the per-hero `[Academy]` link REMOVED per user request 2026-07-15 (cockpit.py:717-721).
- `TestOverlayJs::test_inflight_window_is_30s`, `TestOverlayJs::test_localstorage_access_is_guarded` — assert the 30s in-flight window + 3 try-blocks REMOVED by M218/v2.3.1 (cockpit-spec.md:266-276).
- M242 must end with **0 NEW** failures; these 6 stay (M244 owns them).
