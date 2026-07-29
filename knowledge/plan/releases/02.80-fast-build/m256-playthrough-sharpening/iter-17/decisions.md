# iter-17 — decisions

## D83 — a pointer-interception overlay cannot block a KEY event

`org-admin.members.UC1` sat at `playthrough: TODO` for thirteen iters. iter-05 measured the wall in four
parts and named two untried candidates. Both of those were pointer routes; the one that worked was neither.

Measured, all three, on the live modal:

| route | result |
|---|---|
| click the `<label for=…>` | **REFUTED** — these labels have no `for` at all (14 labels, `labelsWithFor: 0`, inputs with no `id`; antd wraps the input *inside* the label), and the label click is intercepted exactly like the box |
| plain `click()` on the checkbox | **times out** — iter-05 finding 1 re-confirmed, unchanged |
| `focus()` + `keyboard.press('Space')` | **WORKS**: `checked: 0 → 1` **and** submit `"Assign Tags[DISABLED]" → "Assign Tags (1)[enabled]"` |

The dropdown is still open throughout (`openMenuitems: 12` during a passing run) — the route simply stops
needing it to close. Keyboard is also what a human would do, so it is a *more* faithful journey, not a
workaround. The dropdown-free route iter-05 also suggested was never needed.

**The assertion that proves it worked is the SUBMIT BUTTON'S OWN TALLY**, not `isChecked()`: `Assign Tags (1)`
and `enabled` come from React state, i.e. the application telling us it learned about the tick. `isChecked()`
is only the DOM's opinion.

## D84 — RETRACTION: iter-05 D19's `force: true` finding does not reproduce

**This corrects a claim that has been relayed and re-used as established fact, so it is recorded at milestone
level as well as here.**

iter-05 D19 concluded: *"`check({force:true})` DOES flip the DOM state but does **NOT** drive antd's React
handler, so the modal's [Assign Tags] submit stays **disabled** and the assignment never happens."* It was
generalised to *"`force: true` can manufacture a control the application does not know about"*, and this
milestone has cited it since.

**Repeated on this surface today, it does not reproduce.** `check({force:true})` alone yields:

```
before: {"checked":0,"submit":"Assign Tags|DISABLED"}
after : {"checked":1,"submit":"Assign Tags (1)|enabled"}
```

antd's state **is** driven. Found by accident and then confirmed deliberately: a mutant that swapped the
keyboard tick for `check({force:true})` was expected to go RED and **passed**, so the experiment was repeated
in isolation rather than the surprise being explained away.

Findings 1–3 of iter-05 all still hold — a **plain** `click()` still times out, re-measured — so the overlay
interception is real and only the `force`-specific conclusion is withdrawn. Why it differed at iter-05 cannot
be settled: that code is gone, and the likeliest candidate is that the box was located by accessible name and
resolved to a different element.

**What survives, and why keyboard still ships.** The transferable half is untouched: **`force: true` exists to
SKIP actionability checks, so it is the one interaction that CAN manufacture a state the application never
learns about** (iter-07's rule). A route that needs no `force` cannot have that failure mode. Choosing keyboard
is therefore a decision about *what a green run is evidence of*, not a claim about what `force` can do — and
the code comment that first said "`force: true` could never produce the tally" was overstated and has been
corrected in place.

*The general lesson, which is the reason this is a decision and not a footnote: a mutant that PASSES when you
expected RED is data. The first instinct was to adjust the mutant; the right move was to re-run the original
experiment.*

## D85 — the negative control came with the journey, by construction

The Playthrough **creates** the tag it assigns (a P1-sanctioned setup step, driven through the real UI), so
the tag's member tally starts at a known **0** and the final asserts `0 → ≥ 1` on a *different surface* after a
navigation. A run in which the assignment does not land reads 0 and fails — proven by mutant M1 (assignment
step removed → **RED**). This is iter-06 D22's pattern (a mutating Playthrough's pre-state read IS its
control), and it is why the Playthrough arrives already covered: negative controls went **21 of 24 → 22 of 25**,
the numerator and denominator both moving by one.
