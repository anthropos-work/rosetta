---
milestone: M266
title: "Cockpit legibility"
milestone_shape: section
status: planned
release: "02.100-content-consolidation"
depends_on: "none"
parallel_with: "M270"
complexity: medium
last_updated: "2026-08-23"
---

# M266: Cockpit legibility

**Goal:** The presenter cockpit reads at a glance — hero cards pack two-up, the hiring vantage speaks
candidate language, and content-story cards lose the box-in-a-card.

Serves the five field annotation requests **A1–A5** (`.agentspace/annotations.md` § *cockpit menu*).

**The whole milestone lives in ONE file.** All of A1–A5 is
`.agentspace/rosetta-extensions/demo-stack/cockpit.py`, which is **server-rendered Python f-strings —
no template engine, no CSS file**. There is one `_PAGE_CSS` constant at **`:517`**. Nothing here is a
platform-repo change (see § Constraints).

## Scope

**In:**

  - **A1 — half-row heroes (org stories).** Non-manager hero cards should take **half a row**, so two
    heroes fit on one row. Cards are emitted at **`cockpit.py:1365-1377`** and joined into
    `<section class="story">` at **`:1378-1383`**. `.hero` at **`:543`** is `display:flex; margin:10px 0`
    — **full width by construction**. Wrap in a grid and give the manager card `grid-column:1/-1`. The
    manager/non-manager split is **ALREADY available** as `hero.get("vantage_label") == "MANAGER"`.
    ⚠️ **COUPLING — keep both attributes intact:** `demo-stack/tests/test_cockpit.py:151` asserts
    `page.count('class="btn login"') == n_heroes`, and **`:684`** regex-scrapes
    `<a class="btn login" href=... data-login-as=...>`.

  - **A2 — candidate labels (the hiring vantage).** In *"Candidate Hiring & Comparison"* the non-manager
    heroes are **candidates, not employees**, and their arc reads **performing / under-performing**, not
    thriving / struggling. `_badges()` at **`:659-672`** hardcodes `EMPLOYEE`/`MANAGER` +
    `THRIVING`/`STRUGGLING` and **never reads `is_hiring`**. Everything needed is **ALREADY in the
    manifest** — `stack-seeding/seeders/cockpit.go:246` (`IsHiring`, threaded from `st.IsHiringOrg()`
    at **`:315`**) and **`:232`** (`Trajectory`). **ZERO Go change, ZERO re-seed.** The teal `--cand`
    token already exists at **`cockpit.py:527`**. The section title *"Candidate Hiring & Comparison"* is
    **already correct** in `stack-seeding/presets/stories.seed.yaml:247` — **do not touch it**.

  - **A3 + A5 — ONE work item.** Both rewrite `_content_tuple_row()` at **`:1117-1173`**.
    - **A3 (no box within a card).** The "box within a card" is literally **`.ctcol`** (CSS at
      **`:629-631`**) — the two bordered pass/fail columns **M242** added inside `.ctuple`; the markup is
      at **`:1147-1153`**.
    - **A5 (Academy is consumed content, not pass/not-pass).** Needs the partition at **`:1136-1141`**
      to **stop being unconditional**.

  - **A5 (data half) — a new per-product `has_verdict`.** Pass/fail **is NOT a product property today.**
    It needs a new per-product `has_verdict` field on `contentProductMeta`
    (`stack-seeding/seeders/content_manifest.go:149-157`, registry **`:163-193`**).
    ⚠️ **THIS TRIPS THE HONESTY GATE:** the canonical `stack-seeding/presets/content-manifest.json` must
    be **regenerated and `CanonicalFileMatchesProjection` re-passed IN THE SAME COMMIT**
    ([`content-stories-spec.md` §4](../../../../../corpus/ops/demo/content-stories-spec.md)).

  - **A4 — a flag before each language copy.** `_LANG_PILL_LABEL` / `_LANG_TOGGLE_LABEL` at
    **`:957-958`**. ⚠️ **The cockpit's only external asset is the FontAwesome free CDN, which HAS NO
    country flags; and emoji flags do not render on Chrome/Windows. Use INLINE SVG** — it keeps the
    panel stdlib-only and self-contained.

  - **The two card contracts in the corpus move with the code** — see § Delivers →. The docs *are* the
    contract; a card-shape change that lands in `cockpit.py` alone leaves the corpus asserting the old
    shape.

**Out:**

  - any change to the **seeded story content itself** (A2 is a render-side relabel: zero Go change, zero
    re-seed)
  - the **EN/IT toggle behaviour** (v2.6 M241) beyond adding flags to its labels

## Depends on

none

## Parallel with

M270 — different repo, zero shared files.

## Open questions

These are genuinely open at scaffold time. **Resolve them in `spec-notes.md` before writing code; do not
resolve them by guessing in the markup.**

  - **A2's exact replacement label set.** The annotation asks for *"candidate"* instead of *"employee"*
    and *"performing" / "under-performing"* instead of *"thriving" / "struggling"*. Unresolved: whether
    the candidate label keys off `is_hiring` **alone** or `is_hiring` **AND** non-manager. Note the
    precedent already in the file — `_avatar_class()` documents that a hiring **manager** seat must let
    the manager branch win first *"so she reads as a manager, not a candidate"*. `_badges()` should
    almost certainly mirror that ordering, but it is not yet a decision.
  - **A5's `has_verdict` membership.** The annotation names **Academy** explicitly. Which of the other
    registry products (`simulation`, `skill-path-legacy`, `ai-labs`) carry a verdict is **not settled by
    the annotation** and must be derived from what each product's result surface actually shows.
    `ai-labs` is already presence-only, so it is not the same case as Academy.
  - **A3 — where the pass/fail SIGNAL goes once `.ctcol` is gone.** Removing the two bordered columns
    removes the only place the verdict is currently stated as a heading. Does the signal move onto the
    per-cell pill, onto the row, or somewhere else? The `.p-pass` / `.p-fail` pill styles already exist
    and are used by `.ctcol-hd`; reusing them per-cell is the obvious route but is a layout decision
    M266 has not taken.
  - **A4 — which flag for "english".** A language is not a country. `_LANG_PILL_LABEL` maps
    `english → "English"` and `italian → "Italiano"`; the Italian flag is unambiguous, the English one
    is not (GB vs US vs a non-flag glyph). A real design call, not a lookup.
  - **A1 — the odd-count and no-manager cases.** What a story with an odd number of non-manager heroes
    renders (trailing half-width card vs stretched), and what a story with **no** manager renders, are
    unspecified by the annotation.
  - **Line anchors will need re-resolving at milestone start.** A scaffold-time read of the authoring
    copy found every construct above present and unambiguous, but several anchors resolve a few lines
    off the cited value (e.g. `--cand` reads at `:524`, `.ctcol` CSS at `:623`, the `.ctcol` markup at
    `:1152-1154`, `_badges` at `:659`). **The citations above are the milestone contract and are carried
    verbatim; re-resolve them against the tree before editing** rather than trusting either number.

## KB dependencies

- [`cockpit-spec.md`](../../../../../corpus/ops/demo/cockpit-spec.md) — the **M43/M242 card contract**
  (A1/A2). § *The UI surface (v1.10 M43)* and § *The v2.6 "sound check" UX pass (M242)* describe the card
  shape this milestone changes; § *The hiring vantage — the recruiter + 2 candidate seats (v2.4 "casting
  call" M224)* is where the hiring-seat model already lives.
- [`content-stories-spec.md` §7.6 + §4](../../../../../corpus/ops/demo/content-stories-spec.md) — **the
  card columns and the honesty gate** (A3/A4/A5). §7.6 is the EN/IT language toggle (A4); §4 is the
  honesty gate + fail-closed resolver the `has_verdict` change trips. **Also read §7.2** (*the
  tuple-regrouped row + the two-action contract — M234 contract, M242 layout*), which is the section
  that documents the `.ctcol` columns A3 removes.

## Delivers →

- [`corpus/ops/demo/cockpit-spec.md`](../../../../../corpus/ops/demo/cockpit-spec.md) — the org-stories
  card contract changes (A1's half-row grid + the manager full-bleed row; A2's candidate/performing badge
  set for the hiring vantage).
- [`corpus/ops/demo/content-stories-spec.md`](../../../../../corpus/ops/demo/content-stories-spec.md) —
  the content-story card contract changes (A3's flattened row, A5's per-product `has_verdict` and the
  conditional partition, A4's flagged language labels) **and** §4's honesty-gate note now covers a
  net-new projected field.

Both are revisions, not net-new docs. **The docs are the contract** — this is why they are in scope.

## Constraints (standing, non-negotiable)

- **Zero platform-repo edits.** Every file M266 touches is **rext-owned** (`demo-stack/cockpit.py`,
  `demo-stack/tests/test_cockpit.py`, `stack-seeding/seeders/content_manifest.go`,
  `stack-seeding/presets/content-manifest.json`) or corpus, so **no platform source is expected to be in
  scope at all**. If any part of A1–A5 turns out to require a platform-source change, it routes through
  the sha-pinned **demopatch** mechanism
  ([`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)) — never a repo edit, never an
  ad-hoc file change inside a stack dir.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed
  to origin**, then consumed per-stack at a pinned tag. *Tagging is not publishing* — a tag that exists
  only locally is unreachable to a remote stack.
- Secrets handled values-blind.
