---
iter: 243
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-243 — `--reload-sentinel` names a `demo-N` container on a `dev-N` stack

**Active strategy reference:** `TOK-08`.
**Route worked:** `ROUTE-M257x-239-stackseed-sentinel-reload-is-demo-only`, opened at iter-239 while
grading that iter's consequence and deliberately routed rather than worked (the scope-creep tripwire's
third line).

## Step 0 — re-survey

Still live, and it is a **working-stack** defect in the tooling rather than a corpus one — the half of
the user's redirect that the 235–242 series has served only indirectly.

`stack-seeding/cmd/stackseed/main.go`:

* `shouldReloadSentinel` fires on **any** non-prod stack with `n > 0` — its own doc says so, and the
  `n > 0` clause is there to skip the *dev stack N=0*, which makes the `dev-N > 0` case explicitly in
  scope;
* `reloadStackSentinel(n, out)` then takes **only `n`** and reconstructs
  `stackName := fmt.Sprintf("demo-%d", n)`.

So on a `dev-2` stack the RPC leg is correct (it is offset-keyed, `8087+n`), and the **`docker restart`
fallback** names `demo-2-sentinel-1` — a container that a dev-only box does not have. The seed already
succeeded and the whole step is best-effort, so the operator sees a warning about the wrong container
and the re-seeded `g2`/`g3` grants **never reach the running enforcer** — the silent-403 class
`corpus/ops/verification.md` exists to catch.

`ParseStackN` is the reason the bug is invisible: it discards the prefix and returns only the offset, so
by the time `n` reaches this function the stack's family has been thrown away.

## Pre-registered claims — SEALED IN THIS COMMIT

- **`P-243-1`.** `reloadStackSentinel` receives no stack name — only `n` — so the `demo-` prefix cannot
  be anything but hard-coded. **Predict: confirmed by signature.**
- **`P-243-2`.** `shouldReloadSentinel` admits `dev-N` for `n > 0`; there is **no** dev/demo predicate
  anywhere on this path. **Predict: confirmed, 0 such predicates.**
- **`P-243-3`.** The existing tests cover the *gate* (`shouldReloadSentinel`) and **not** the container
  name. **Predict: ≥ 1 test on the gate, 0 asserting the container name.**
- **`P-243-4`.** The same hard-coded-`demo-` shape occurs **elsewhere** in the seeding section.
  **Predict: ≥ 1 more site.** (Refutable: it may be the only one.)

## Phase plan

1. Seal this pre-registration.
2. Measure `P-243-1`…`P-243-4` against the source.
3. Repair: pass the stack NAME, not just its offset. Keep the RPC leg untouched — it was never wrong.
4. Regression-test the container name for both families, and grade whatever `P-243-4` finds.

## Escalation conditions

If the fix would change what `--reload-sentinel` does on a demo stack in any way, stop — the demo path is
the proven one and this is a dev-path correction.

## Acceptable close-no-lift outcomes

If `dev-N` cannot in fact reach this path, the route is closed by falsification and the finding is
downgraded in place.
