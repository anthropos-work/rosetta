# iter-286 — decisions

## D-M257x-286-1 — defect 4's mechanism is SOURCE-SUPPORTED, not browser-confirmed, and it is labelled so

`mapItem` gives every dropdown row `onClick: () => handleClick(item)`, and `handleClick` does nothing at
all for an outbound item: it calls `setSelectedKey` only for `MenuType.link`, and routes only when
`item.key.startsWith('/')`. An absolute cockpit URL is neither. Navigation therefore came **only** from
the `<Link href>` that `navbarMenuItems` nests inside the icon and the label — while the antd row is
wider than both, carrying `paddingInlineStart`/`paddingInlineEnd`. A click in that padding runs the
no-op and closes the dropdown.

That accounts for *"the first time it doesn't work, the second time it does"* — it is a **hit-target**
defect, not a state-machine one. **It was not reproduced in a browser**, because the only stack available
is the one the user is validating on. The fix is stated with that limit attached rather than as a
confirmed diagnosis: it removes the dependence on hitting the nested link, which is correct whether or
not the padding is the exact path the user's clicks took.

## D-M257x-286-2 — the row override SPREADS the mapped item

`{ ...mapItem(item, 0), onClick: … }` rather than a hand-built row: the icon, the label, the padding and
the key all keep coming from `mapItem`. Hand-building would navigate correctly and lose the FontAwesome
icon and the depth padding — **a fix that works and looks wrong is still a defect.** Pinned as an arm.

## D-M257x-286-3 — defect 2 is ESCALATED, not patched, and the reason is that it is NOT demo-only

Two independent defects in three lines of `builderAssistant.js`, either sufficient on its own: the
`/default_avatar.png` fallback is overwritten by `getUserPicture()`, which returns `undefined` **without
throwing** for a user with no Clerk image (so the `catch` never runs and the guard never fires); and the
asset it names **does not exist in the repo**.

A demo-patch would have been easy and **wrong**: any user without a Clerk image hits this in production,
so patching the demo hides the symptom on the one surface where it does not matter. Recorded in
`knowledge/plan/platform-defect-register.md` with both mechanisms, the exact anchors, and why the bot
avatar loads while the user's does not.

## D-M257x-286-4 — the fail-closed CONTRACT registry was updated, not worked around

Changing the replacement turned `FAIL_CLOSED_CONTRACT`'s pinned fragment RED — the registry doing its
job. The entry was re-pinned to the new expression **and widened** to cover the log-out swap, so the
registry now fences both halves of the demo-only behaviour rather than the half that existed first.
