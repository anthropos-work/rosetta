---
milestone: M257x
iter: 12
---

# iter-12 — decisions

## D1 — the VM-RAM pre-flight's two halves, one real and one refuted

The unit defect was real in the **report** (integer-floored GiB) and, worse, in the **remediation** (a
decimal-GB instruction that can never clear a GiB floor). The *comparison* defect I had also drafted was
**refuted by mutation N2**: for an integer floor `m`, `floor(x) >= m` iff `x >= m`. The comparison was always
correct. Recorded in the code comment rather than dropped, per §8 rule 3 — a comment claiming a fix that
fixed nothing is how a false contract gets pinned.

## D2 — `DOC-M257x-claude-md-knob-count` closed REFUTED, not fixed

`CLAUDE.md:307` and `corpus/ops/demo/README.md:153` already say 30. The only surviving "27" is
`coverage-protocol.md:1056`, a historical account of what v2.5's close corrected — accurate as written.

## D3 — iter-05's cold-start fix had been applied to the DEV twin, not the DEMO twin

Cold cycle 1 went RED on directus `exit(1)`. Root cause: the fix landed on the twin that was **not** the one
the defect was measured on, and the test passed **because it tested the twin that was fixed**. This is the
"a check that reports without measuring" class in its purest form — the 9th occurrence this session — and it
is why the standing rule is *prefer a live negative control*. Fixed and fenced; rext `fast-build-m257x-iter-12b`.

## D4 — clause 1 is not attemptable at origin HEAD, and this is measured, not inferred

Platform origin/main moved to `2adcf71` (2026-07-31 15:58) and **dropped the WunderGraph/Cosmo federation
router outright** — the `graphql` service, its `repos.yml` entry, and its clone. GraphQL is now served
directly by `backend` at `:8082/graphql/query` (note the **path change**, `/graphql` → `/graphql/query`).

Proven live, not asserted:

1. `gen_injected_override.py` run against the origin-HEAD compose returns **RC=0** — it does not notice —
   and still emits `depends_on: graphql` plus two `WUNDERGRAPH_SSR_ENDPOINT=http://graphql:8080/graphql`.
2. `docker compose config` on base+override returns **RC=1**:
   `service "hiring-app" depends on undefined service "graphql": invalid compose project`.

So a cold cycle at origin HEAD **fails at project validation, before a single build**. Clause 1 could not
have been satisfied at this ref by any amount of cycling. Evidence: `evidence/iter-12-router-drop.md`.

## D5 — the fence gap this exposes, routed forward

Clause 4 fences the schemas rext **writes** that the platform no longer creates. Nothing fences the compose
**services rext depends on** that the platform no longer defines — which is the same class, one axis over,
and it is why the generator returned RC=0 above. `FENCE-M257x-iter13-compose-service-exists` is the named
handler; it belongs to iter-13 with the re-point, because a fence written without the RED it must catch is
the anti-pattern this milestone keeps re-learning.

## D6 — re-scope trigger: occurrence 1 of 2, recorded so the next one fires

The milestone's `re_scope_trigger` fires when **TWO consecutive** full-alignment attempts are invalidated by
new platform commits landing mid-milestone. This is **occurrence 1** — cold cycle 1 was killed mid-flight
because origin moved under it. Not a trigger yet. Recorded explicitly so occurrence 2 is recognised as the
second and escalates rather than being re-derived as another first.

## D7 — scope-creep tripwire fired; iter closed rather than absorbing the re-point

Three distinct lines of investigation in this iter: the RAM pre-flight, the directus twin, and now the router
drop. The tripwire fires on the third. The re-point is ≥3 rext sites plus its fence plus three ~18-min cycles
— routed to iter-13 whole rather than half-landed here.
