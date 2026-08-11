---
milestone: M257x
iter: 13
---

# iter-13 — progress

**Type:** tik

## What was actually wrong, and how much of it iter-12 had not seen

iter-12 proved clause 1 was **not attemptable**: platform `2adcf71` deleted the GraphQL router and
`docker compose config` rejected the demo project outright. It routed **6** named sites forward. Re-measuring
first — the milestone's standing rule — found **~14 files**, including three iter-12 never named:

- `up-injected.sh`, **six** `$((5050+OFFSET))/graphql` sites (3 build-args + 3 image-reuse validators)
- `stack-verify/lib/readiness.sh`, whose probe introspects *"the federated supergraph at :5050"* — after
  `2adcf71` that is a port with **no listener**, so the probe would have reported the API **down while it was
  up**
- `stack-verify/lib/services.sh`, a `graphql | anthropos-graphql-1 | localhost:5050` census row

## The re-point, and the shape it took

**Two things moved, not one:** `graphql:8080` → `backend:8082`, **and** `/graphql` → `/graphql/query`. The
address change fails loudly (a refused connection). The **path** change does not — it resolves, connects and
404s, which is the latency-budget *fast-failing fetch* signature (≈ 3 × 33 ms + 6 s) and reads as a slow page
rather than a broken one. A re-point matching on hostname alone would have shipped green.

So the fix is a **derivation, not a correction**:

- `up-injected.sh` gains `browser_graphql_endpoint()` — one definition, called by all six sites. Six
  hand-written copies of a platform fact is the hand-maintained-tuple defect this milestone exists to end
  (`platform-alignment.md` §2), and correcting six copies would have left the defect intact.
- `gen_injected_override.py` gains `BACKEND_SERVICE` / `GRAPHQL_PATH` / `SSR_GRAPHQL_ENDPOINT`, one each. The
  two **test files that had duplicated the literal** now import the constant instead of re-typing it — that
  duplication is exactly what would have had to be found by hand next time.
- `gen_tailscale_serve.py`'s `("graphql", 5050)` row is **deleted, not re-pointed**: `backend` (8082) was
  already in `API_BROWSER_FACING`. Fronting a port with no listener yields a trusted-cert HTTPS endpoint that
  always refuses — worse than no entry, because it looks configured.

## The fence, watched RED before it was believed

`FENCE-M257x-iter13-compose-service-exists`: every emitted `depends_on` must name a service the platform
compose **at the ref in use** defines. Same class as clause 4's schema fence, one axis over — clause 4 covers
what rext *writes*, this covers what rext *depends on*.

- **Mutant** (restore `depends_on: graphql`) → exit **1**, naming the vanished service and listing what the
  platform does define. **Unmutated control** → exit **0**.
- It **reports what it checked** — *"2 depends_on target(s) all defined by the platform compose:
  ['backend', 'postgresql']"* — and **fails closed** if it finds `depends_on:` blocks and extracts zero
  targets. A parser defect there would pass against every input; that is the class this session has now hit
  9+ times.
- **Placement was the real design decision.** The first cut asserted inside `build_lines` and turned **16
  unit tests into errors** — those tests drive the builder with truncated one-service fixtures, so the fence
  was asserting platform completeness against a *fixture*. Moved to `main()`, where `cfg` is the real
  resolved compose (§8 rule 4). The four main()-driving tests got a realistic fixture instead: a platform
  with no database and no backend could not boot.

## Attribution, and a control that was quietly lying

Baseline taken by `git archive`-ing rext HEAD into scratch: demo-stack **3** failures. In place, the same
code: **7**. The tell was the **skip count — 26 vs 2**. Live-clone tests resolve their clone by relative
path, so in scratch they *skipped* rather than ran. **A control that silently skips the very tests you are
attributing is not a control.** The in-place pre-existing set turned out to be exactly the 7 of
`CHECK-M257x-live-clone-suites-red`, so the attribution held — but it held because the numbers were
reconciled, not because the control was sound.

Attributed: **26** injection · **8** core · **17** demo-stack · **2** dev-stack were mine. All fixed.

## Measured

| | before | after |
|---|---|---|
| `docker compose config` @ origin HEAD | **RC=1** *(undefined service `graphql`)* | **RC=0**, 16 services |
| `gen_injected_override.py` @ origin HEAD | RC=0, silently wrong | RC=0 + fence reports 2 targets checked |
| stack-injection | 0 fail (baseline) → 26 mine | **OK**, 277 |
| stack-core | 14 (pre-existing batteries) | **14** — my 8 fixed |
| demo-stack | 7 pre-existing + 17 mine | **7** — the known live-clone set |
| dev-stack | 0 (baseline) → 2 mine | **OK**, 122 |

**Zero regressions; every section back to its baseline exactly.**

Side-effect paid, as predicted: editing `up-injected.sh` restaled **22** `file:line` citations in
`demo-up-defaults.md`; repaired with the guard's own `--fix`, re-verified *"OK — both directions."*

`CLAUDE.md` was stale on **two** independent counts — the router, and a *"3 subgraphs"* claim the cms-in-app
merge had already reduced to one. Corrected. The remaining corpus surface is **35 files / ~128 hits**, routed
to clause 5.

## Close — 2026-08-01

**Outcome:** clause 1 goes from **not attemptable** to **attemptable** — iter-12's own live negative control
flips **RC=1 → RC=0** against the same origin-HEAD compose. The gate metric itself (green cold cycles) is
still **0 of 3**; iter-14 runs them. rext `fast-build-m257x-iter-13` (`4414527`) is **verified on origin** and
both the tag file and the `stack-demo` consumption clone are re-pinned to it.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (still occurrence 1 of 2) —
(4) user-blocker: n — (5) cap-reached: n (2 tiks this session, cap is 5) — (6) protocol-stop: n —
Outcome: continue
**Decisions:** D1–D8 (iter-13/decisions.md)
**Side-deliverables:** the 22 restaled `demo-up-defaults.md` citations repaired; `CLAUDE.md`'s router and
subgraph-count claims corrected.
**Routes carried forward:**
- `FIX-M257x-iter13-freshness-vs-origin` → **iter-14**. Planned as this iter's phase item 4 and **not
  landed** — it is a 4th line of investigation in `ensure-clones.sh`'s freshness subsystem. The defect
  stands: the check compares the checkout to `clones.pin.json` rather than to **origin**, so a stale *clone*
  is reported as a stale *pin*. It is why iter-12 nearly spent three 18-minute cycles on the wrong ref.
- `DOC-M257x-iter14-corpus-router-drop` → **clause 5**. **35 files / ~128 hits** still describe the Cosmo
  router, `:5050`, or `graphql-wundergraph` — measured, not estimated.
- `CHECK-M257x-live-clone-suites-red` — **unchanged at 7**, and now independently confirmed as the exact
  in-place pre-existing set.
**Lessons:**
1. **Re-measure the inherited list.** iter-12's hand-off named 6 of ~14. Fifth consecutive iter in which an
   inherited hand-off was materially incomplete or wrong.
2. **When a platform fact moves, count the copies before fixing them.** Six copies in one file and two more
   in test files that had duplicated the literal. Fixing copies leaves the defect.
3. **A fence's placement is a design decision, not a detail.** Inside the pure builder it asserted platform
   completeness against unit fixtures and produced 16 false errors; on the `main()` path it asserts the same
   invariant against the real thing and produces none.
4. **Check what your control actually ran.** The scratch baseline differed from the in-place run by 24 SKIPS,
   not by 24 outcomes — and a skip reads as "fine" in every summary line.
