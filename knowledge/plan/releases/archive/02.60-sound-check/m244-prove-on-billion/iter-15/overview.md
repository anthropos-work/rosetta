---
iter: 15
milestone: M244
iteration_type: tik
status: closed-fixed
created: 2026-07-23
---

# iter-15 — gate (d): the anon academy FS-published fallback (unblocks the (c)↔(d) coupling)

**Type:** tik · **Active strategy:** TOK-02 (sweep the remaining gates live; iter-14 re-ordered gate (d) to the front as it unblocks the gate-(c) coverage cross-port)

## Active strategy reference
TOK-02, as re-ordered by iter-14's (c)↔(d) coupling finding: fix the academy-empty (gate d) first because it also fails the gate-(c) coverage cross-port to the academy home.

## Cluster / target identified
Gate (d): the anon academy `/library`+`/free` twin renders EMPTY (iter-09) and the workforce coverage sweep's cross-port follow to the academy home renders empty (iter-14). Root cause (iter-09): the anon routes read `serverTenant.js::getPublicCatalogView()` (`getBackendCatalogView(new Set())`) which the M230 `academy-fs-published-fallback` patch left OUT OF SCOPE (it only fixed the authed `getServerCatalogView`). So on a Clerk-free demo the public catalog view is empty.

## Hypothesis
A chained sibling demopatch (`academy-fs-published-public`) that adds the SAME FS-as-published fallback to `getPublicCatalogView` (with `new Set()`, public-only) makes the anon academy render real cards → gate (d) met AND the gate-(c) cross-port unblocked.

## Expected lift
Gate (d) → MET (4/8 → 5/8); the gate-(c) coverage cross-port passes.

## Phase plan
Author the chained patch + native helper + ant-academy.sh wiring → local chain test → push rext + move tag → re-pin billion + apply to the academy clone (native, no re-bake) → verify anon renders → re-run coverage.

## Escalation conditions
A platform-repo edit required → STOP. (None — a demopatch on the ephemeral clone.)

## Acceptable close-no-lift outcomes
If the anon academy still could not render real cards after the fix, close-no-lift with the falsification.
