# iter-286 — the first click that did nothing, and the avatar that never could load

**Type:** tik — under `TOK-09`. The last two of the four named product defects.

## Defect 4 — the row's own click handler is a no-op for an outbound item

`mapItem` gives every account-dropdown row `onClick: () => handleClick(item)`. `handleClick` reads:

```ts
item.type === MenuType.link && setSelectedKey(item.key);
… item.key.startsWith('/') && onClick(item.key);
```

The Back-to-Cockpit item is `MenuType.outbound` with an **absolute** URL — **neither branch fires**. So
navigation came only from the `<Link href>` that `navbarMenuItems` nests inside the icon and the label,
while the antd row is wider than both (it carries `paddingInlineStart`/`paddingInlineEnd`). A click on
that padding ran the no-op and closed the dropdown.

That accounts for *"the first time it doesn't work; if I do it again it does"* — a **hit-target** defect,
not a state-machine one. **Fixed** by overriding the row's own handler:
`{ ...mapItem(backToCockpitMenuItem, 0), onClick: () => { window.location.href = cockpitUrl } }` — spread,
not hand-built, so the icon, label, padding and key still come from `mapItem` (`D-M257x-286-2`).

> **Stated as a limit, not buried:** this mechanism is **source-supported, not browser-confirmed**. The
> only stack available is the one the user is validating on, and reproducing a click defect on it was not
> a trade worth making. The fix removes the dependence on hitting the nested link, which is right whether
> or not the padding was the exact path their clicks took.

## Defect 2 — two independent defects in three lines, and neither is fixable here

`app/sim-advanced-builder/builderAssistant.js:741-745`:

1. **The fallback is overwritten by a getter that returns nothing without throwing.**
   `getUserPicture()` is `return this.clerk?.user?.imageUrl` — an optional chain. A user with no image
   yields `undefined`, the `catch` never runs, the unconditional assignment lands, and `avatar.src` is
   `undefined`. The **error** path is guarded; the **empty-success** path is not — *a check that skips
   reads exactly like a check that passes*, in assignment form.
2. **The declared fallback asset does not exist.** `/default_avatar.png` occurs **once** in the whole
   repo, at that line, and `app/public/` holds no such file. Even a correct guard lands on a 404.

The bot avatar loads because it is a literal, `/avatar_bot_nobg.png`, and that file **is** there — which
is exactly the asymmetry the user described.

**ESCALATED, not patched** (`D-M257x-286-3`). This is `studio-desk` platform source; v2.8 holds 0
platform edits. And it is **not demo-only**: any user without a Clerk image hits it in production, so a
demo-patch would hide the symptom on the one surface where it does not matter. Recorded in
`knowledge/plan/platform-defect-register.md` with both mechanisms and exact anchors.

## The manifest was RE-ANCHORED, because changing `post` while `pre` named an older file is worse

`next-web-back-to-cockpit`'s `pre_sha256` no longer matched the live clone — pre-existing upstream drift,
and per `demopatch-spec.md` the **anchor is the contract; the whole-file sha is only a baseline**. Since
this iter recomputed `post`, leaving `pre` pointing at a file it no longer describes would have been the
more confusing state. Both re-pinned to the clone this iter measured, with the old value and the reason
recorded in the manifest.

## Verification

| scope | result |
|---|---|
| `demo-stack` whole section, `--tb=no -rf` | **1082 passed · 9 failed · 1 skipped** (3 m 39 s) |
| `test_back_to_cockpit_m249.py` (+3 new arms) | **40 passed** |
| the row-navigation + log-out-swap arms, reverted | **RED**, then restored |

**The 9 failures are pre-existing and none is this iter's** — stated as an argument, since no
before-tree was run: they are 4 modules, every one of them a `*_against_live_clone` / `*Live` test.
`test_migrate_race_live` (3) needs a live container and `demo-1` was brought down at 09:09Z;
`test_ssr_origin_chain` (3) hashes `packages/graphql/src/server/server.graphql.ts`, `test_demopatch`'s
two hash `urls.ts` and the studio chain, and `test_ant_academy`'s round-trips `next.config`. **None of
them reads `NavbarTop.tsx`, the back-to-cockpit manifest, or `up-injected.sh`'s cockpit launch.** The
first run reported 13 failed but its captured summary named only 5, so it was **re-run to name all of
them** rather than reporting a number without its members.

**NOT COVERED:** no browser repro of defect 4; no stack brought up, re-seeded or restarted; the other
next-web manifests' upstream drift is left where it is — re-anchoring manifests this iter did not touch
would be re-pinning baselines against a clone whose freshness nothing here established.

## Close — 2026-08-11

**Outcome:** defect 4 fixed at the row level with its mechanism and its evidential limit both stated;
defect 2 diagnosed to two independent causes and **escalated** to the platform defect register rather
than patched, because patching it would have hidden a production defect on a demo.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**

**Decisions:** `D-M257x-286-1` … `D-M257x-286-4`.

**Routes carried forward:**
- **`ROUTE-M257x-286-next-web-manifest-baselines-have-drifted`** — 9 pre-existing live-clone failures
  across four modules. They are freshness gates doing their job against a clone that moved; re-anchoring
  them is a deliberate pass, not a side-effect of this one.
- Unchanged: `ROUTE-M257x-285-logout-swap-for-studio-and-academy`,
  `ROUTE-M257x-285-demo-2-cockpit-serves-a-stale-world`.

**Lessons:**
1. **Say when a mechanism is source-supported rather than reproduced.** The fix is right either way; the
   claim about *why* is the part that would have been over-stated.
2. **A demo-patch that hides a production defect is worse than no patch** — it removes the only surface
   where anyone was going to notice.
3. **A failure count without its members is not a result.** The first run said 13 and named 5; re-running
   to name all 9 cost four minutes and turned an assumption into an argument.
