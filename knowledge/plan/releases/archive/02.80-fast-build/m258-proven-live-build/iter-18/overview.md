---
iteration_type: tik
status: in-flight
milestone: M258
iter: 18
opened: 2026-08-12
---

# iter-18 — the 8th merge reaches the corpus, and the fence that watches it goes green

**Type:** tik — under `TOK-01` (*measure the composition before engineering it*), extended by the
user's ruling that the milestone's remaining scope is **debt paydown to zero**.

## Step 0 — re-survey before targeting

`TOK-01`'s next-tik direction is spent (the composition is measured, the batch gate is wired, the
15-red set is closed). The controlling direction for this iter is the **user's own ruling**, relayed
through the milestone: *"i want no debt"*, with `CORPUS-M258-iter16-sentinel-in-app` named as the
item. Re-surveyed before targeting, three ways, and all three confirm the target is live:

1. **The platform moved and the corpus does not know.** `stack-demo/platform` is at `766df6c`
   (*"chore(compose): remove sentinel service and related configurations"*), whose `repos.yml` lists
   **three** repos — `app`, `next-web-app`, `studio-desk`. `CLAUDE.md` says *"the 4 repos in
   repos.yml"* and names `sentinel` among them.
2. **The fence is RED, measured directly (no pipe).**
   `PLATFORM_REPOS_YML=…/stack-demo/platform/repos.yml python3 platform_alignment_guard.py` → **`rc=1`,
   17 findings**: 1 `[B departure]` (*"the map claims sentinel is in repos.yml, and it is not"*) + **16
   citation failures**, because `repos.yml` went 28 → **13** lines and `docker-compose.yml` 190 → **164**.
3. **The census is large but the wrong claims are a subset.** 71 files mention `sentinel`; most
   mention the DB **schema** (`sentinel.casbin_rules`), which is still correct — `SENTINEL_DB_CONNECTION`
   at `docker-compose.yml:25` still carries `search_path=sentinel`. The wrong claims are the
   **structural** ones: a container, a `repos.yml` entry, an always-on floor member, an RPC edge, a
   Tier-1 service.

## Active strategy reference

`TOK-01` — with the standing rider from the user's ruling (`iter-10/decisions.md` `D52`,
`iter-11/decisions.md` `D57`): the gate is closed by ruling; remaining scope is debt to zero, then
`END-M258-one-stack` on a stack built by the **fixed** tooling.

## Cluster / target identified

`CORPUS-M258-iter16-sentinel-in-app`, routed unchanged from iter-16 and iter-17 and named by the user
as the "no debt" item. The drift is **structural, not cosmetic** — the corpus describes a service
topology the platform no longer has.

## Hypothesis

The 8th merge (`sentinel` → `app`, platform v11.0) can be landed in the corpus **completely** in one
iter, because the wrong claims cluster into six mechanical classes, and because the map that the
machine fence watches is the single canonical statement the rest can point at:

1. the **always-on floor** (three → two: `postgresql`, `redis`; `sentinel` is not in `common.yml`)
2. the **`core` container count** (five → four)
3. the **`repos.yml` membership** (four → three)
4. the **cross-process Connect-RPC edge** (`backend → sentinel` → **zero**, because `app` deleted its
   RPC listener entirely)
5. the **Tier-1 service list** (`app` + `sentinel` → `app` alone)
6. the **`sentinel` service doc** (live service → merged/redirect, the `skiller`/`skillpath` shape)

## Expected lift

Not a metric-lift iter. The deliverables are binary and checkable:

- `platform_alignment_guard.py` exits **0** against `stack-demo/platform/repos.yml` (from `rc=1`, 17).
- The six classes above are corrected at every load-bearing site, each cited to `766df6c` /
  `app c52dbc51e`.
- No prose that is *correct* is edited to make a fence green (the release's own rule).

## Phase plan

- **A — measure the platform** at `766df6c` + `app c52dbc51e`: compose shape, `common.yml` include,
  profiles, ports, `repos.yml`, `app/internal/sentinel/`, the Redis invalidation channel.
- **B — the fenced map**: rewrite the `sentinel` row (both cells), repair all 16 citations, re-run the
  guard to green.
- **C — the sweep**: `CLAUDE.md` + the architecture + services + ops docs, by class.
- **D — gates**: re-run the alignment guard, the citation/anchor guards, and the `stack-core` suite
  against a pre-measured baseline (pre-existing failures must not be attributed to this iter).
- **E — close.**

## Escalation conditions

- If a fence is wrong and correct prose would have to be edited to satisfy it → **fix the fence**, and
  say so. Never the reverse.
- If the residual census cannot be closed within the iter → land the canonical statement, route the
  named remainder with a file-level census so it needs no re-discovery, and close
  `closed-fixed-partial` rather than claiming a complete sweep.

## Acceptable close-no-lift outcomes

None expected — the fence is measurably RED, so a green fence is a floor. A close-no-lift would only
be honest if the RED turned out to be the fence's defect rather than the corpus's, which is itself a
reportable finding.
