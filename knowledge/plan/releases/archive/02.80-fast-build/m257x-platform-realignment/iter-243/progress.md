**Type:** tik — under `TOK-08`, working `ROUTE-M257x-239-stackseed-sentinel-reload-is-demo-only`.

# iter-243 — `--reload-sentinel` restarted a `demo-N` container on every `dev-N` stack

## The defect

`stack-seeding/cmd/stackseed/main.go`:

* `shouldReloadSentinel` fires on **any** non-prod stack with `n > 0`, and its own doc says the `n > 0`
  clause is there to skip *the dev stack at N=0* — which puts **`dev-N` for N > 0 explicitly in scope**;
* `reloadStackSentinel` took **only `n`** and reconstructed `stackName := fmt.Sprintf("demo-%d", n)`.

The RPC leg was never wrong — it is offset-keyed (`8087+n`) and reaches whichever stack owns that
offset. The **`docker restart` fallback** named `demo-N-sentinel-1`, a container a dev-only box does not
have. And because the seed has already succeeded and the whole step is best-effort, the operator gets a
warning about **the wrong container** while the re-seeded `g2`/`g3` grants **never reach the running
enforcer** — the silent-403 class `corpus/ops/verification.md` exists to catch.

**Why it was invisible:** `blueprint.ParseStackN` returns only the **offset**; it takes everything after
the first `-` and discards the prefix. By the time `n` arrived at this function the stack's family had
been thrown away, so the function could not have known — it could only guess, and it guessed `demo`.

This is the third finding in this run to trace back to the same root: `ParseStackN` accepting and
flattening an unqualified target. iter-239 found the corpus **teaching** the unqualified spelling at 5
sites; this is the code that cannot recover from it.

## The repair

`sentinelContainerStack(stack, n)` — keep the name, do not re-derive it (`D-M257x-243-1`). A stack whose
name carries a `dev-`/`demo-` family keeps it; anything else (an unqualified `--stack 1`, an empty
target) falls back to the historical `demo-N`, so **the demo path is byte-identical**, which was this
iter's stated escalation condition (`D-M257x-243-2`).

**Two tests, and the second is the one that matters more.** `TestSentinelContainerStack_…` covers both
families, a high offset, and both fallbacks. `TestSentinelReload_GateAdmitsDevStacks` pins the *premise*:
the gate is family-blind, so the container name is the only place the distinction can be honoured — if a
future change makes the gate demo-only, that test fails loudly instead of leaving the new function
silently unreachable.

**The gate was already tested and the container NAME was not. That is precisely how the defect
survived** — the existing `TestShouldReloadSentinel_GatePredicate` has 8 cases, including
`high-n-demo-reloads`, and every one of them asserts the boolean while none asserts what the reload then
does.

## Pre-registration — scored 3 confirmed / 1 refuted

| claim | prediction | result |
|---|---|---|
| `P-243-1` the function receives no stack name | confirmed by signature | **CONFIRMED** — `reloadStackSentinel(n int, out *os.File)` |
| `P-243-2` the gate admits `dev-N`; no dev/demo predicate on the path | 0 predicates | **CONFIRMED** — now pinned by a test |
| `P-243-3` tests cover the gate, not the container name | ≥1 gate test, 0 name tests | **CONFIRMED** — 8 gate cases, 0 name assertions |
| `P-243-4` the hard-coded `demo-` shape recurs elsewhere | ≥ 1 more site | **REFUTED — exactly one site** in the whole section |

`P-243-4`'s refutation is a **stop** signal, not a disappointment: with a population of one there is no
class to census and no fence to build (`D-M257x-243-3`). `§5` iter-168 — *measure a hazard's size, or
"the same problem exists elsewhere" is only a mood* — used in the direction that says do less.

## Close — 2026-08-10

**Outcome:** the `--reload-sentinel` post-seed step now restarts the stack's **own** sentinel container
instead of always naming a `demo-N` one; the demo path is byte-identical and the dev path is correct for
the first time. The route that opened at iter-239 is closed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5
**Decisions:** `D-M257x-243-1` (keep the name, do not re-derive it) · `D-M257x-243-2` (an unqualified
target keeps the historical `demo-N`, so the demo path is byte-identical) · `D-M257x-243-3` (`P-243-4`
refuted; the class is one site and is not generalised).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — Go (`stack-seeding`, `go test ./...`): **12 packages, all ok, 0 failed**,
including the 3 net-new cases. Python unchanged from iter-242: `stack-core` guard family
**26 GREEN / 0 RED / 0 could-not-check / 5 not-run**.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-239-stackseed-sentinel-reload-is-demo-only` → **CLOSED** by this iter.
- `ROUTE-M257x-241-wider-citation-surface-is-ungraded` → open — **the strongest remaining lead**: 107
  corpus citations into the six frozen-legacy repos, outside the migration map, graded by nothing.
- `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` → open.
- `ROUTE-M257x-238-claude-md-fences-are-unmaintained` → open, six-for-six.
- `ROUTE-M257x-238-container-vs-native-is-undrawn` → open, three independent hits.
- `ROUTE-M257x-237-critical-env-list-is-unfenced` → open.
- `ROUTE-M257x-236-disclosure-scope-is-document-level` → open.
- `ROUTE-M257x-235-fence-scope-is-unread` → open.
- `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **A function that is handed an offset cannot honour a distinction the offset discards.** The defect
   was not a typo; it was the only thing the signature allowed. Fix the signature, not the string.
2. **Testing a gate is not testing what the gate admits.** Eight cases pinned the boolean — including a
   `high-n-demo-reloads` case that reads like coverage of exactly this path — and none asserted the
   action taken behind it.
3. **A refuted "it recurs elsewhere" prediction is a licence to stop.** Under a census strategy the pull
   is always toward sweeping; a measured population of one earns no sweep.
