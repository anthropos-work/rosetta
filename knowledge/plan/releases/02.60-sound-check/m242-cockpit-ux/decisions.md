# M242 — Decisions

_(implementation choices with rationale accumulate here during build)_

## D1 — the `.session` cell stays the atomic language-filter unit (why the regroup didn't churn `_LANG_JS`)
The M242 tuple regroup restructures the content tab from one-row-per-session to one-row-per-`(sim_type,
modality)` tuple with passed/not-passed columns. **The per-session login-options block kept its `class="session"`
+ `data-lang`/`lhide` markup** (now `_content_login_cell`, nested inside a column), so `_LANG_JS` (which filters
`.session[data-lang]`) needed **zero change** and the whole M241 `TestContentLanguageToggle` /
`TestContentToggleLangs` suite survived untouched. What moved: the **pass/fail** pill → the column header; the
**modality** pill → the tuple title. Alternative (a new `.ctopt` class + a `_LANG_JS` selector edit) was rejected
— more churn, no benefit.

## D2 — non-simulation products group by `label`, not by an empty tuple
A skill-path/academy/ai-labs session has empty `sim_type` AND `modality`, so `(sim_type, modality)` would merge
**every** non-sim session of a product into one `("","")` bucket. `_content_tuple_key` falls back to
`("label", label)` for those, so each distinct content item is its own row. (An ai-labs SAMPLE session with a
modality but no sim_type/label groups under `("sim","",modality)` and reads title "Session" — a synthetic-fixture
quirk; real ai-labs carry a `label` + empty modality.)

## D3 — an empty verdict column reads "No passing / No failing run" (the eyeball-artifacts guard)
A tuple with only one verdict (only a pass, or only a fail) shows the other column with an **explicit muted
marker**, not a blank cell — so an empty column can't be misread as "the tooling broke". Per the build-Phase-1
look-at-the-output guideline.

## D4 — the byte-identical-no-content header is a SEPARATE branch, not one flex template
Section 2 moves the tab bar into the header. To preserve the HARD byte-identical-when-no-content invariant, the
`else` (no-content) branch of `render_page` reproduces the **exact pre-M242 header** (2/4/6/6/4/2-space
indentation and all), rather than a single `.hwrap` flex template with an empty right slot (which would change the
no-content HTML). `test_no_content_header_is_byte_identical_to_today` pins a verbatim copy of today's header. The
two branches differ ONLY by the flex wrapper (`.hwrap`/`.hgroup`) + the tab bar the content tab needs.

## D5 — the candidate avatar colour = teal `#0f766e`/`#ccfbf1` (net-new), manager-wins order
Manager reuses the existing `.b-manager` orange (`#fef3e2`/`#b45309`); employee keeps the indigo accent. The
candidate needed a NEW colour (none existed): teal (`--cand`/`--cand-soft`), distinct from both indigo and orange
and ~4.8:1 on its soft bg (AA). `_avatar_class` checks **manager first**, so a hiring recruiter (a `MANAGER` seat
with `is_hiring`) reads as a manager, not a candidate — candidates are the non-manager hiring seats (EMPLOYEE +
`is_hiring`).

## D6 — the content-tab `.sicon` stays uniform (the "if it reads by type" clause is a no-op here)
The overview asked to "color the content-tab session icon `.sicon` consistently **if it reads by type**." The
`.sicon` reads by **sim_type** (clipboard-check/dumbbell/…​), not **user-type**, so the user-type tint doesn't
apply; it stays the uniform accent. Coloring it by user-type would be wrong (a tuple can mix seats).

## D7 — the 6 pre-existing `test_cockpit.py` failures: all outside M242's surface → left to M244
The 6 standing failures were assessed against M242's 3 touched surfaces (content-tab render / header / hero
avatar): **none fall inside.** The 4 academy-link tests assert the per-hero `[Academy]` link removed per user
request 2026-07-15 (cockpit.py:717-721); the 2 overlay-JS tests assert the 30s in-flight window + 3 try-blocks
removed by M218/v2.3.1 (both changes already in the docs) — a DIFFERENT cockpit area (M244 owns it, per the
milestone prompt). M242 ended with **0 NEW** failures. The only test churn M242 made is **Fate-1**: 3 stale
`<div class="stitle">` title asserts updated to `<div class="cttitle">` — the exact title container the tuple
regroup moved, in the surface M242 rewrote.
