# iter-35 decisions

## D-M257x-35-1 — measure the target before touching the test (the standing route, honoured)

The hand-off's instruction was explicit: *"Measure which content the first row actually is before touching
anything."* It was followed, and it changed the fix. The inherited hypothesis — iter-27's new hero
sessions reordered the grid — predicts a *reordering*; the measurement found a **tie**: 11 contents at an
identical `max(started_at)`, 2 of them without the hero.

Had the reordering hypothesis been implemented instead (e.g. re-pinning the expected content, or dating
the hero's session later), the test would have gone green **and remained a coin-flip**, because nothing
about the tie would have changed. The distinction was only visible from the data.

## D-M257x-35-2 — the fix selects on the assertion's own property, and fails loudly when it cannot

`drillIntoContentContaining(name)` returns **-1** when no scanned content carries the member, and the
spec asserts `>= 0`. This is deliberate: a scan that silently falls back to row 0 would convert a real
tenancy failure into a pass, which is the failure mode the whole assertion exists to prevent (§5 rule 7 —
a probe must not be able to satisfy itself).

The scan is bounded at 8 rows so a broken grid cannot burn the test budget, and the settled-page
`heroRow.count()` check is explicitly relabelled in the spec as a re-assert rather than the load-bearing
claim — because after this change the selection guarantees it, and a check that cannot fail should not be
dressed as a final.

## D-M257x-35-3 — `force: true` was available and was refused

The tooltip interception (`ant-tooltip-container` covering row *i+1* after `goBack()`) would have been
silenced by `click({ force: true })`. That was rejected: `force` makes a click pass by **skipping the
actionability check**, so it would assert that the element is clickable by declining to find out. Parking
the pointer and awaiting `detached` fixes the actual condition and leaves the check intact.

Worth recording because the defect is real product behaviour a user would hit — a tooltip covering the
next row after a back-navigation — and `force` would have hidden it permanently.
