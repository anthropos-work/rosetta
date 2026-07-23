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

## D8 — the language toggle now maintains the empty-column marker (close Phase 2c fix)
The M242 verdict-column split (§(1)) crosses the M241 per-cell EN/IT toggle. `_content_tuple_row` renders the
"No passing / No failing run" marker **only for a column that is cell-less at RENDER time** — but the toggle
hides cells at VIEW time. So an **unbalanced bilingual tuple** (a verdict present in only one language — e.g. a
requirement passed only in english, failed only in italian, with no english-fail / italian-pass counterpart)
would, when toggled to the other language, leave a verdict column showing its header over a **blank body with
no marker** — the exact "misread as broken" failure **D3** forbids. **Fix:** `_LANG_JS.syncEmpty()` recomputes
per-`.ctcol` emptiness after every `apply()` and once on load, injecting a `ctempty-dyn` "No passing/failing
run" marker into any column whose cells are all language-hidden (and removing it when re-filled). **Pure
client-side — zero server-markup change**, so every existing render test (byte-identical header, the exact
static empty-marker asserts, the placement + coexist tests) stays green. Guarded by
`test_lang_toggle_syncs_empty_column_marker_for_unbalanced_bilingual_tuple` (render-shape + call-site ordering,
mutation-verified: dropping either call site → RED).

### Adversarial review (close Phase 2c)
- **Scenario considered:** the column split × the per-cell language toggle on an UNBALANCED bilingual tuple →
  a language-emptied verdict column with no marker (D3 violation). **Reachable by design** (the pass/fail and
  language axes are independent), but the **canonical `content-manifest.json` seed is currently BALANCED** —
  every cross-verdict bilingual tuple carries both verdicts in both languages, so the precise phantom-empty
  condition has **0 live occurrences** today. Classified as a **latent robustness gap** M242's split
  introduced, not a live demo defect. Landed the fix (D8) rather than accept-with-risk: the risk is in
  just-shipped render code and the guard is cheap + self-contained. Live browser proof rides M244's e2e
  execution (the language-toggle spec is already in `stack-verify/e2e/`, unexecuted this release).
- **Other angles checked (no change needed):** `_content_tuple_row(sessions[0])` — `sessions` is always ≥1
  (built via `dict.setdefault`, keyed per appended session). All manifest free-text (title/label/modality/
  icon_key/language) routes through `html.escape` (the harden pass pinned the two previously-unpinned paths).
  `_avatar_class` missing-key defaults are safe (absent `vantage_label`/`is_hiring` → employee default).
