# iter-288 — defect 3, closed as a FAMILY instead of left as a permanent partial

**Type:** tik — under `TOK-09`, closing `ROUTE-M257x-285-logout-swap-for-studio-and-academy`.

## Why this was worth an iter rather than a carried route

iter-285 landed the Log-out swap for `next-web` and routed the two siblings honestly, with named reasons.
**Nothing graded the route.** Three per-patch arms would all have been green while two of three apps
still showed a dead-end Log out — which is how *"three of four"* becomes permanent. So the deliverable is
the family fence as much as the two edits.

## ant-academy — guard the BUTTON, keep the ROW

The academy's logout **row** also holds a `connectionSlot`. Wrapping the row would have removed an
unrelated control along with Log out — a fix whose blast radius exceeds its subject. The anchor was
extended over the `<button>…</button>` only.

**The first hand-typed attempt matched 0 occurrences**, two spaces out on the inner indentation. The
anchor is now taken **from the file's own bytes** — retyping source into a manifest is exactly how an
anchor silently stops applying, and this fence family exists because of that class.

Proven live: `apply` → the guard is present at `UserMenu.jsx:156` → `revert` → **byte-identical**
(`diff -q` clean, `git status` clean).

## studio-desk — the chain, re-pinned in order

`studio-desk-logout-url.pre_sha256` **is** `studio-desk-back-to-cockpit.post_sha256`, so editing the
first re-pins both. Checked **before** touching anything: the chained patch anchors on `handleLogout()`
— the **method**, not the button — so removing the button leaves it applying.

Re-pinned in dependency order (pristine → apply 1 → that hash *is* 2's `pre` → apply 2 → its `post`), and
proven by a live **apply-in-order → LIFO revert** with the demopatch tool: *"demo clone left git-clean"*
twice, `diff` byte-identical.

The divider and the button collapse to `''` on the **same** `VITE_COCKPIT_URL` test that adds the item,
so off-demo the block is emitted verbatim.

## The fence is over the FAMILY

`TheLogOutSwapCoversEveryAppThatGotTheCockpitItem` iterates the three patches and requires each to make
Log out conditional; a second arm **derives the family from disk** (`*-back-to-cockpit`) and fails if the
hand list stops matching, so a fourth app cannot inherit the swap by omission. A third pins the academy's
row-vs-button distinction. `FAIL_CLOSED_CONTRACT` gained the two new expressions. **RED-proven** by
neutralising the academy guard.

## Verification

| scope | result |
|---|---|
| `demo-stack` whole section | **1085 passed · 9 failed · 1 skipped** (3 m 38 s) |
| the same section at iter-286 | **1082 passed · 9 failed · 1 skipped** |
| the 9 failures | **the identical set, module for module** — `test_migrate_race_live` (3, needs a live container; `demo-1` is down), `test_ssr_origin_chain` (3, `server.graphql.ts`), `test_demopatch` (2, `urls.ts`), `test_ant_academy` (1, `next.config`) |
| studio chain apply→LIFO revert on the real clone | **git-clean, byte-identical** |
| academy apply→revert on the real clone | **git-clean, byte-identical** |
| the family fence, academy guard neutralised | **2 failed** — RED-proven, restored |

**+3 passed and +0 failed** is the whole delta: the three new arms, and nothing disturbed.

**NOT COVERED:** no stack was brought up, re-seeded or restarted; the patches were applied and reverted
on the clones, never left applied. The 9 pre-existing failures are freshness gates against clones that
moved upstream — re-anchoring manifests this iter did not touch stays a deliberate pass
(`ROUTE-M257x-286-next-web-manifest-baselines-have-drifted`).

## Close — 2026-08-11

**Outcome:** defect 3 is closed across **all three** apps that carry the Back-to-Cockpit item, with the
family — not the members — fenced, and both hard cases handled on their own terms: the academy keeps its
logout row, and the studio chain is re-pinned in dependency order and proven by a live LIFO round-trip.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y (7 tiks, cap extended in-band by the user's recovery instruction — see iter-287) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**

**Decisions:** `D-M257x-288-1` … `D-M257x-288-3`.

**Routes carried forward:**
- **CLOSED:** `ROUTE-M257x-285-logout-swap-for-studio-and-academy`.
- Unchanged: `ROUTE-M257x-285-demo-2-cockpit-serves-a-stale-world`,
  `ROUTE-M257x-286-next-web-manifest-baselines-have-drifted`, `ROUTE-M257x-284-demo-2-is-live-and-uncontained`,
  and iter-282's prose-twin residuals.

**Lessons:**
1. **A routed partial needs a fence, or it is a permanent partial.** The route was honest and named; what
   it lacked was anything that would go RED while it stayed open.
2. **Take an anchor from the file's bytes, never from your fingers.** The first attempt was two spaces
   out and matched nothing — in a manifest that silently stops applying rather than failing loudly.
3. **Check the dependent before editing the dependency.** The studio chain survived because the chained
   patch anchors on the method, and that was established *first*, not discovered by a broken apply.
