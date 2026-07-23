---
title: "KB Fidelity Audit — M242 cockpit-UX"
date: 2026-07-22
scope: milestone:M242
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Content-tab row render (regroup by tuple) | `corpus/ops/demo/content-stories-spec.md §7.2 / §7.6` | `demo-stack/cockpit.py::render_content_tab / _content_session_row / _content_session_actions` | PAIRED |
| Cockpit header + tab selector | `corpus/ops/demo/cockpit-spec.md §"The 2nd tab (M234)" / §"The UI surface (M43)"` | `demo-stack/cockpit.py::render_page` (header block ~L782) + `_TAB_JS` | PAIRED |
| Hero icon bg by user-type | `corpus/ops/demo/cockpit-spec.md §"For PMs" / §"The UI surface"` | `demo-stack/cockpit.py` `.hero .avatar` CSS (L144) + `_badges` (L239) + `is_hiring` (L715) | PAIRED |

## Fidelity Findings

1. **`.hero .avatar` uniform indigo** — Doc (`cockpit-spec.md`) describes a per-hero `fa-circle-user` avatar with a single indigo accent; code CSS L144-146 (`background: var(--accent-soft); color: var(--accent)`) is uniform. **Verdict: ALIGNED.** M242 §3 will introduce role-tinting.
2. **`is_hiring` per-hero flag** — required to derive the candidate color (`is_hiring && vantage != manager`). Present at `render_page` L715 (`hero.get("is_hiring")`). **Verdict: ALIGNED.**
3. **`_badges` emits `MANAGER`/`EMPLOYEE`** — the `vantage_label` source for §3's role tint. L241-245. **Verdict: ALIGNED.**
4. **Content-tab per-session-row model** — `content-stories-spec.md §7.2` describes one row per played session (icon + descriptor + up-to-two CTAs); code `render_content_tab` matches. §7.6 **explicitly forward-references M242**: *"M242 owns the row-REGROUP by tuple that will pair the EN/IT variants onto one row"*. **Verdict: ALIGNED** — doc is the intended contract M242 now realizes.
5. **Tab conditional-render / byte-identical-when-no-content** — `cockpit-spec.md §"The 2nd tab"` states "Absent `--content-manifest` ⇒ no tab bar (byte-identical to today)"; code gates on `has_content` (L745-758). **Verdict: ALIGNED** — the invariant M242 §2 must preserve.

## Completeness Gaps

1. **(incidental)** `cockpit.py` module docstring L40-46 still narrates the per-hero `[Academy]` link ("each hero card also renders an [Academy] link") which was **removed** at L717-721 per user request 2026-07-15. This is a **code-comment** staleness, not a knowledge-doc claim (the `cockpit-spec.md` render diagram shows only `[Log in as]`, correctly). Not load-bearing for M242's three sections; owned by the M244 cockpit-cleanup surface (same area as the 4 stale academy-link tests). Recorded, not fixed here (out of M242 scope).

## Applied Fixes
None required — every doc claim in scope is ALIGNED. Triples recorded in `spec-notes.md`.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. The two delivered docs both exist (no blind area) and accurately describe current code; no stale load-bearing claim the implementation would misread. The 6 pre-existing `test_cockpit.py` failures are TEST-staleness (academy-link removed / overlay 30s-window removed per M218+v2.3.1 — both changes already reflected in the docs), owned by M244, outside M242's three touched surfaces.
