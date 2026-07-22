# M242 — Retro

## Summary
M242 was the **fifth post-barrier fix** in v2.6 "sound check" (cockpit UX), a `section` milestone serial after
M241 (it wanted the pass/fail + language variants to regroup against). Three render-layer changes made the
presenter cockpit read more clearly, **all in `cockpit.py`, 0 data/seed change, 0 manifest schema change, 0
platform-repo edits:** **(1)** the Content-stories rows **regroup by requirement tuple** `(sim_type, modality)`
(non-sim → `label`) into one row per tuple — `target | passed login options | not-passed login options` side by
side, with symmetric "No passing/failing run" empty markers and a presence-only inline slot; the old
`_content_session_row` split into `_content_tuple_row` + `_content_login_cell` so the M241 EN/IT toggle's atomic
`.session` cell was untouched. **(2)** the **tab selector moved into the white header** (flex `.hwrap`, right,
vertically centered), `_TAB_JS` placement-agnostic, the **byte-identical-when-no-content invariant preserved** via
a separate verbatim else-branch (mutation-STRONG). **(3)** the **hero avatar is tinted by user-type** —
manager orange / employee indigo / candidate net-new teal (~4.86:1 AA), **manager-wins** order. **2 harden passes**
mutation-verified every render branch and fixed **2 toothless tests** + added an AA-contrast pin. The **close's
adversarial pass** found + landed **1 latent D3-invariant gap** (D8). Delivered docs: `cockpit-spec.md` (the v2.6
UX-pass section + the avatar glance note) + `content-stories-spec.md` (§7.2 tuple row + render-helpers + the
empty-marker-under-toggle note). Clean complete 3-section close.

## Incidents This Cycle
- **P3 (latent, caught at close) — the empty-column marker did not hold under the language toggle.** M242's
  verdict-column split crosses M241's per-cell EN/IT toggle: an **unbalanced bilingual tuple** (a verdict present
  in only one language) would, when toggled, show a verdict header over a **blank body with no marker** — a silent
  violation of D3's own "an empty column always carries a marker, never misread as broken" invariant. **Reachable
  by design** (the pass/fail and language axes are independent), though the canonical `content-manifest.json` seed
  is currently **balanced** (0 live occurrences). Fixed inline at close (D8): `_LANG_JS.syncEmpty()` re-derives the
  per-column marker on every toggle + on load (pure client-side, **0 server-markup change**), guarded by a
  mutation-verified test. Not a shipped user-visible defect (the seed never triggers it today), but a real
  invariant hole in just-shipped code — the adversarial pass's job.
- **P3 (harden) — 2 toothless tests.** Harden pass 1 found a wrong-column mutant (a FAILING session filed into the
  PASSED column) and an unescaped non-sim title mutant both SURVIVED; both tests were strengthened to bite. No
  production bug — the render logic was correct; the tests just couldn't see a break.
- **No regressions.** The full demo-stack Python suite ran **839 pass / 9 fail** — the 9 are the identical standing
  set (6 academy+overlay `test_cockpit.py` + `test_host_prereqs_m215` + `test_purge` + `test_reap` reap-17700),
  **0 new from M242** (matches the M241-close 9-fail baseline; the pass count grew 838→839 by the one close test).

## What Went Well
- **The byte-identical-no-content invariant was honoured structurally, not hoped for.** Section (2) moved the tab
  bar into the header behind a HARD back-compat line; rather than a single flex template with an empty slot, the
  no-content path reproduces the exact pre-M242 header as a separate branch, and `test_no_content_header_is_byte_
  identical_to_today` pins it verbatim (mutation-STRONG at harden). A load-bearing invariant was fenced by a test
  with teeth, not a comment.
- **The adversarial pass earned its keep on a "cosmetic" milestone.** A render-only UX polish still hid a latent
  cross-feature invariant break (M242's columns × M241's toggle). Constructing the input, confirming the phantom
  render, and checking the seed's actual balance (0 live) turned a vague worry into a precise, fated fix + a guard —
  exactly the "sound check" discipline applied to the cockpit.
- **The regroup didn't churn M241.** Keeping the per-session `.session` cell as the atomic language-filter unit
  (D1) meant the whole M241 `TestContentLanguageToggle`/`TestContentToggleLangs` suite survived untouched — the
  restructure moved the pass/fail + modality pills out of the cell without touching the toggle's selector.

## What Didn't
- **The build under-tested its own column split — twice.** Harden had to fix 2 toothless tests (wrong-column +
  title-escaping) that proved presence but not placement/escaping, and the close's adversarial pass found a THIRD
  gap the build's tests didn't cover (the empty marker under the toggle). The pattern: a render change that
  introduces a new *layout* needs tests that pin *which slot* content lands in and *what happens when a sibling
  feature hides a cell* — presence-only asserts are the toothless default. Not a shipped defect (all closed before
  close), but a reminder that layout-change tests must assert placement + cross-feature interaction, not presence.
- **The close had to re-disposition inherited test noise again.** 6 of the 9 standing demo-stack failures surface
  through `test_cockpit.py` (the file M242 touched), so every M242 run shows red that isn't M242's — the close
  re-confirmed provenance (identical set, 0 new). The standing debt has now ridden **≥5 v2.6 milestones**; **M244
  is the named expiry** and should discharge it by editing the tests (6 of 9 need no live stack).

## Carried Forward
- **None new from M242** — a clean complete section close, 0 new deferrals.
- **(Inherited, confirmed) → M244:** the **9 standing demo-stack test failures** (Fate-2, M238-D5 / M239-D13
  reap-17700) — 6 surfaced here via `test_cockpit.py` — already homed; M244 is the named expiry.
- **(Follow-on proof) → M244:** the live browser behaviour of `_LANG_JS.syncEmpty` (per-column marker inject on
  toggle) is proven structurally + by render-shape here; the live click-swap proof rides M244's e2e execution (the
  language-toggle spec is already in `stack-verify/e2e/`, unexecuted this release).

## Metrics Delta
- **Tests:** `test_cockpit.py` 142 (M241) → **164** (+22 M242: 16 build + 5 harden + 1 close/adversarial), 158
  pass / 6 pre-existing standing. Full demo-stack suite **839 pass / 9 fail** (848 collected). Go **2005** + TS
  **151** unchanged (M242 is Python-render-only). Flake **0** (5/5). Platform-repo edits **0**.
- **rext code-of-record:** tag `sound-check-m242-cockpit-ux` @ **`73d37d5`** (re-pinned at close to the close-fix
  HEAD; build/harden HEAD was `b0e76b3`).
