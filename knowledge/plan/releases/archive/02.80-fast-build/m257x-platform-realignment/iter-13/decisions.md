---
milestone: M257x
iter: 13
---

# iter-13 — decisions

## D1 — the routed-forward list was incomplete, as the standing rule says to assume

iter-12 routed 6 named sites. Re-measuring found **~14 files**. The residual it had not named included the
**six** `$((5050+OFFSET))/graphql` sites in `up-injected.sh`, `stack-verify/lib/readiness.sh` (whose probe
introspects *"the federated supergraph at :5050"* — it would have reported the API **down** while it was up),
and `stack-verify/lib/services.sh`'s probe row. Re-surveying cost minutes; trusting the hand-off would have
left a re-point that looked complete and was not.

## D2 — six hand-written copies became one derivation, not six corrected copies

The browser's GraphQL endpoint was written out at six sites (3 build-args + 3 image-reuse validators). Fixing
six copies is fixing the instance; the defect is that there were six. `browser_graphql_endpoint()` is now the
one definition, and the test asserts **the derivation exists and is called ≥3×** rather than that a port
literal appears somewhere — *prefer a design that cannot express the bug over a check that catches it*.

Same move in `gen_injected_override.py`: `BACKEND_SERVICE` / `GRAPHQL_PATH` / `SSR_GRAPHQL_ENDPOINT` are one
definition each, and the two test files that had **duplicated the literal** now import the constant.

## D3 — the path moved, and that is the half that would have survived a re-point

`/graphql` → `/graphql/query`. A wrong **host** refuses fast and loudly. A wrong **path** resolves, connects,
and 404s — the latency-budget *fast-failing fetch* signature (≈ 3 × 33 ms + 6 s), which reads as a slow page
rather than a broken one. Any re-point that had matched on hostname alone would have shipped green.

## D4 — the fence sits on the main() path, not inside build_lines

First cut asserted inside `build_lines`. That turned **16 unit tests into errors**, because those tests drive
the builder with deliberately truncated one-service fixtures — and the fence was then asserting platform
completeness against a fixture, which tests the fixture, not the platform (§8 rule 4: scope a fence to its
enclosing block). Moved to `main()`, where `cfg` **is** the real resolved compose. Every production caller
reaches the generator through `main()`; a direct `build_lines` import is a test path by definition. The four
main()-driving tests got a realistic `_main_cfg()` instead — a platform with no database and no backend could
not boot, so the old fixture was describing something impossible.

## D5 — watched RED before believed, and it fails closed when it measures nothing

A mutant restoring `depends_on: graphql` exits **1** naming the vanished service; the unmutated control exits
**0**. The fence also **reports what it checked** (*"2 depends_on target(s) … ['backend', 'postgresql']"*) and
raises if it finds `depends_on:` blocks but extracts **zero** targets — a parser defect would otherwise pass
against every input, which is the "reports without measuring" class this milestone keeps re-learning.
Emitting zero blocks is legitimate (`--no-ui --no-local-content` drops both); finding blocks and no targets
never is.

## D6 — a contaminated control, caught by arithmetic

The baseline was taken by `git archive`-ing rext HEAD into scratch. It reported demo-stack **3** failures;
in place the same code reported **7 pre-existing**. The difference is the **skip count — 26 vs 2**: the
live-clone tests resolve their clone by relative path, so in scratch they SKIPPED rather than ran. A control
that silently skips the tests you are trying to attribute is not a control. The in-place pre-existing set is
exactly the 7 of `CHECK-M257x-live-clone-suites-red`, so attribution held — but it held because the numbers
were reconciled, not because the control was sound.

## D7 — CLAUDE.md was stale on two independent counts

`CLAUDE.md:217` claimed *"Apollo Federation v2 gateway (3 subgraphs)"*. The router is deleted **and** the
subgraph count had already been wrong — the cms-in-app merge took the supergraph to one. Corrected in place.
The remaining corpus surface is **35 files / ~128 hits**, routed to clause 5 rather than swept here.

## D8 — scope: the freshness-vs-origin fix was NOT landed

Planned as phase item 4. It is a 4th line of investigation inside an iter that already ran three, and it sits
in `ensure-clones.sh`'s freshness subsystem — the wrong place to open with the remaining budget. Routed to
iter-14 with a named handler. Recorded here rather than quietly dropped: the iter closes
`closed-fixed-partial` **because** of this, not in spite of it.
