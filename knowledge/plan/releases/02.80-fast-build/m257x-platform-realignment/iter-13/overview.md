---
milestone: M257x
iter: 13
iteration_type: tik
status: closed-fixed-partial
opened: 2026-08-01
---

# iter-13 — re-point rext off the deleted GraphQL router, and fence the class

**Active strategy reference:** `TOK-01: instrument first, then follow` — still step 5/5, *"prove it cold,
three times."* iter-12 established that the proof cannot be attempted: the platform deleted the router and
`docker compose config` rejects the demo project. This iter removes that blocker and fences it, so step 5 is
attemptable in iter-14.

## Step 0 — Re-survey

Re-ran the measurement rather than trusting iter-12's routed-forward list, per the milestone's standing rule.
**The list was incomplete — 6 named handlers, ~14 real files.** The residual iter-12 did not name:

- `up-injected.sh` — **six** `$((5050+OFFSET))/graphql` sites (3 build-args + 3 image-reuse validators)
- `stack-verify/lib/readiness.sh:26-28` — the readiness probe introspects *"the federated supergraph at :5050"*
- `stack-verify/lib/services.sh:47` — a `graphql | anthropos-graphql-1 | localhost:5050` probe row
- `playthroughs/e2e/run-playthroughs.sh:77`, `stack-core/union_apply_guard.py`, the demopatch YAML headers

The **destination is already fronted**: `gen_tailscale_serve.py`'s `API_BROWSER_FACING` already serves
`("backend", 8082)`, so the browser's GraphQL endpoint moves onto a port that already has an HTTPS listener.
In-cluster, backend's compose sets `PORT=8082` and publishes `8082:8082` → `http://backend:8082/graphql/query`.

## Cluster / target identified

`FIX-M257x-iter13-router-drop-repoint` + `FENCE-M257x-iter13-compose-service-exists` +
`FIX-M257x-iter13-freshness-vs-origin`, all three routed forward by iter-12 (D4/D5, and the freshness
corollary now in protocol §3).

**The dangerous half is not the one compose caught.** `depends_on: graphql` is structural, so compose
rejected it loudly. The two `WUNDERGRAPH_SSR_ENDPOINT` values and the six `5050` build-args are **strings** —
nothing checks them, and they would present as the latency-budget *blackholing address* signature
(≈ 3 × 10.5 s + 6 s) rather than as an error. **And the path changed as well as the host**
(`/graphql` → `/graphql/query`), which is precisely what a hostname-only re-point misses.

## Hypothesis

Re-pointing every site from `graphql:8080|5050 + /graphql` to `backend:8082|8082 + /graphql/query`, plus a
fence that fails when rext emits a `depends_on` on a compose service the platform's own compose at the ref in
use does not define, takes `docker compose config` from **RC=1 → RC=0** at origin HEAD and makes the class
non-silent going forward.

## Expected lift

Clause 1 goes from *not attemptable* to *attemptable* — the gate metric itself (green cycles) stays 0 until
iter-14 runs them. The claimable deliverable here is the **flip of iter-12's own live negative control**, run
against the same origin-HEAD compose.

## Phase plan

1. Re-point the structural sites (`gen_injected_override.py`: `REUSE_DEV`, both SSR endpoints, the
   `depends_on` block).
2. Re-point the endpoint/port sites (`up-injected.sh` ×6, `gen_tailscale_serve.py`, `readiness.sh`,
   `services.sh`, `clones.pin.json`, `repos/run.sh`) and the comment/doc surfaces that name the old topology.
3. `FENCE-M257x-iter13-compose-service-exists` — **watched going RED against `2adcf71` before it is
   believed**, per §8 rule 5 and the standing "a check that reports without measuring" pattern.
4. `FIX-M257x-iter13-freshness-vs-origin` — measure distance to origin, and say which of checkout/pin/origin
   moved.
5. Tests + mutants; full rext suite; tag, **push the tag to origin**, re-pin `.agentspace/rext.tag` + the
   `stack-demo` consumption clone.
6. Re-run iter-12's negative control: `gen_injected_override.py` + `docker compose config` at origin HEAD.

## Escalation conditions

- A site that cannot be re-pointed without a platform-source change → `demopatch`, never a platform edit.
- A **second** platform commit invalidating this attempt → that is **occurrence 2 of 2** and the milestone's
  `re_scope_trigger` fires; stop and escalate rather than re-point again.

## Acceptable close-no-lift outcomes

A measured finding that the re-point is insufficient — e.g. the frontends need the router's *federation*
semantics and not merely its address — would be the iter's deliverable even with compose still RC=1.
