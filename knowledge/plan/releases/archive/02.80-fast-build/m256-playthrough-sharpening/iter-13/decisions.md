# iter-13 — decisions

## D61 — a structural final's problem is the ASSERTION, not the vantage

iter-12 measured that a *structural* predicate (a stat LABEL visible, a chart count ≥ 1, a "Work" section
present) renders for every member of `pt-world`, because M44's profile-completeness seeder gives every member
a career and skills. That measurement is correct and stands. What it does **not** establish — but was read as
establishing — is that the nine Playthroughs holding such finals cannot have a negative control. Their finals
were structural *because they were written structurally*. Phase A measured three independent hero-specific
facts on the same two surfaces, each reading **0** for the contrast seat. **A bound on a mechanism is only as
durable as the thing it was measured against; when the assertion is ours to change, re-test it.**

## D62 — a sharpened final's magnitudes must be machine-linked to the seed, fail-closed

Sharpening introduces a failure mode the suite did not have: a spec carrying a NUMBER is making a claim about
`seed/pt-world.seed.yaml` with no link back to it. Renumber the seed and three Playthroughs go RED naming a
product regression that never happened — a RED that misattributes its cause, which is the same class of
dishonesty as a green over an absent outcome, pointing the other way. Hence `lib/seed-facts.ts` (declare once)
+ `tests/seed-facts-fence.unit.spec.ts` (parse the YAML and reconcile).

The fence's FIRST test asserts the parse is **not vacuous**, before any comparison runs. The parser is
regex-based (the harness has no YAML dependency and this is not worth adding one for), so a reformat could
make it match nothing — and a reconciliation over an empty parse passes every comparison silently. That is
the milestone's signature defect, found 17 times across two harden passes. Mutant-proven both ways: a drifted
magnitude → RED naming the field; a broken indent → the vacuity guard fires first.

Corollary: `allSkillsTotal()` **throws** for a hero with no seeded skills block rather than returning `NaN`,
and the fence asserts the contrast hero *has* no such block. If she ever gained one, her magnitudes could
coincide with the hero's and the magnitude control would go quietly vacuous.

## D63 — `\b` in a `hasText` regex is unreliable; `textContent` concatenates sibling nodes

Both conjunction locators shipped in this iter read **0 on a page that plainly rendered the thing**.
`hasText` matches against `textContent`, which joins sibling text nodes with **no separator**: the work card
reads `…Meridian LabsFeb 2024 - Present (2 years)…`, so a `\b` before `\w{3}` has no boundary to find and the
pattern matches nothing. Same for `^My Closest Roles\b` ("Roles" then "Based").

`TIMELINE_DATE_ENTRY` keeps its `\b` safely only because it is consumed through `getByText`, which resolves to
leaf-ish elements whose text is not a concatenation. **Same regex, different consumer, different rules.** Both
constants now carry the warning, and the regression test asserts the *concatenated* shape while asserting the
old `\b` version fails on it — the test holds the bug.

This was a **false RED**, caught only because the sharpened finals were run and watched. A false red is exactly
as dishonest as a false green (iter-06's rule), and this is that rule earning its keep in the other direction.

## D64 — a control's liveness floor should be the contrast hero's OWN equivalent, same accessor

`assertPageIsAlive` (bodyLen + link count) rules out the iter-07 dead page. It does not rule out "the stats
did not load", which for a magnitude control is the failure mode that matters. Every absence in the new
control is therefore paired with the contrast hero's own equivalent asserted PRESENT **through the same
accessor** — her current position, her role-context line, her verified count and claimed total. The absence is
then unambiguously about *whose* profile this is, and it also proves the accessor works on this vantage (which
is what makes "and yet no closest roles" meaningful rather than merely unread).

## D65 — the assign flake: a ROW-SCOPED modal unmounted by a members-table re-render, against a ladder that could not re-establish its subject

Three hypotheses refuted by measurement first: iter-11's bloated-policy (measured `g3 = 171` for 191
memberships, **0 orphans** — exactly as designed), the antd `maskClosable` re-click (the re-click **throws** on
the mask; the modal survives), and `press('Enter')` with the dropdown closed (`aria-expanded` stays true,
`dialogCount` stays 1). A fourth: the modal **survived 151 s unattended**.

The trace named it. The modal opened healthy at t+3.79 s (title read, submit resolved to a visible disabled
`<button type="submit">`), the Select's inner input went "not stable" ×3 then "detached from the DOM" at
t+4.15 s, and the **dialog never returned**. The remaining 80 s decompose exactly as the ladder's own bounds:
3 × 15 s of combobox clicks against a locator that cannot resolve + 20 s diagnostic + 15 s expect = 84 s,
the reported duration. The modal's accessible title is *"Assign Skill Path to `<member>`"* — it is **row-scoped**,
rendered by the member row's action cell, so a members-table re-render unmounts it. It had opened 2.2 s after
the first row painted, i.e. while the table was still settling.

**Every bound in that ladder was correct. Bounding makes a stuck attempt yield to the next; it does not make a
dead subject detectable.** Fixed in three parts: `dialogIsOpen()` + re-open at the top of every attempt
(recovery), `waitForMembersTableSettled()` (do not race in the first place), and the composed
`openBuilderAndPickSkillPath()` (see D66). Recovery proven deterministically by reaching the exact state with a
real user action (the modal's own Cancel; `Escape` is disabled on it, measured) — the ladder re-opened the
builder, completed the pick, ended with the submit ENABLED. Regression-pinned in the bounded-interaction fence,
mutant-proven.

Supersedes `FLAKE-M256-assign-under-bloated-policy`, whose hypothesis was wrong.

## D66 — recovery creates a correctness obligation: name the member the ACCEPTED builder targets

Once the ladder can re-open a dead modal, the spec's old shape — `open()` → `targetMemberName()` →
`pick()` — reads the target from a modal that is not necessarily the one that completes the pick. A re-opened
builder could target a different member after a table re-order, and the read-back would then assert against
the wrong row and fail for a reason unrelated to the platform. `openBuilderAndPickSkillPath()` returns the name
read **after** the pick was accepted, so *the member I name is the member I assigned to* holds by construction
rather than by two calls happening to see the same modal. Pinned by the fence (pick must precede the name read).

## D67 — three refuted hypotheses cost less than one guessed fix

Each refutation was a bounded probe (minutes), and the fourth explanation came from the trace's own arithmetic
rather than from reasoning. Shipping any of the first three would have closed a handler and changed nothing —
which iter-07 D31 already recorded as the worst outcome available, and iter-10 declined to do for the same
reason. **Read the artifact the failure already produced before proposing a mechanism.**
