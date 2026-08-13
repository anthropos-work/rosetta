# iter-288 — decisions

## D-M257x-288-1 — the academy guards the BUTTON and keeps the ROW

`UserMenu.jsx`'s logout row also holds a `connectionSlot`. Wrapping the row would have removed an
unrelated control along with Log out — **a fix whose blast radius exceeds its subject.** The anchor was
extended over the `<button>…</button>` only, derived **from the file's own bytes** rather than retyped:
the first hand-written attempt matched **0 occurrences** because the indentation was two spaces out, and
retyping source into a manifest is exactly how an anchor silently stops applying.

## D-M257x-288-2 — the studio CHAIN was re-pinned in the right order, and the chained patch survives

`studio-desk-logout-url.pre_sha256` **is** `studio-desk-back-to-cockpit.post_sha256`. Editing the first
re-pins both, in order: hash the pristine file → apply patch 1 → that hash **is** patch 2's `pre` → apply
patch 2 → its `post`. Checked first, before touching anything: the chained patch anchors on
**`handleLogout()`**, the method, not on the button — so removing the button leaves it applying. Proven
by a live apply-in-order → **LIFO** revert on the real clone, `diff` byte-identical and `git status`
clean.

## D-M257x-288-3 — fence the FAMILY, not the members

The new arm iterates the patch family and requires each member to make Log out conditional; a second arm
derives the family **from disk** (`*-back-to-cockpit`) and fails if the hand list stops matching. Three
per-patch arms would all have been green at iter-285 while two members were still unfixed — **which is
precisely how a routed partial becomes permanent.** RED-proven by neutralising the academy guard.
