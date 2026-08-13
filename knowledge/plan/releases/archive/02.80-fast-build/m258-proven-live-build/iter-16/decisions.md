# iter-16 — decisions

## D76 — The 15-red set is ONE defect in OUR tooling. Attributed, not narrowed.

`ESCALATE-M258-iter15-batch-red-15` is **resolved to a single root cause**, established from the
batch's own artifacts plus platform source. Neither candidate cause iter-15 carried survives.

**Root cause.** Platform `766df6c` (2026-08-11), *"chore(compose): remove sentinel service and
related configurations"*, folded **sentinel into `app`** — the **v11.0 "sentinel-in-app"** merge.
`stack-demo/platform/docker-compose.yml:85` says it in the platform's own words: *"sentinel removed
at v11.0: backend answers authorization in-process and does not talk to that container at all."*
`backend` gains `SENTINEL_DB_CONNECTION` (`:25`) and reads the `sentinel` schema directly.

Our three post-seed reload sites all drove the **standalone** sentinel's `AuthorizationService/Reload`
RPC on `8087+OFFSET`. On a v11.0 stack nothing listens there, so all three missed — and reported the
miss as *"non-fatal — a non-AI-sim run is unaffected."*

**Why that is catastrophic rather than narrow**, in `app`'s own words
(`app/internal/sentinel/watcher.go:5-11`):

> Casbin holds the entire policy set in memory. A write goes to the database AND to the writing
> process's own model… Every OTHER process holding an enforcer **keeps serving the policy it loaded
> at boot, forever. Nothing detects it, nothing logs it**, and the symptom is an authorization answer
> that is silently months stale.

The watcher fires on **casbin's own write path**. Our seeders write g2/g3 grants with **raw SQL**,
which is not on it. So after `--reset` + re-seed, the enforcer held a policy containing **no grant
belonging to any freshly-seeded user**, and every org-scoped check was refused.

**The measured chain, end to end** (`up-demo3-b.log`, batch 10:57–11:07Z):

| step | evidence |
|---|---|
| enforcer loads policy at boot | backend log `10:51:44 INFO policy invalidation active … instance=de4a740d` — **before** the reset |
| reset drops the grants | `deleted 1182 casbin grant(s) (g2 role + g3 feature)` (`:498`) |
| the world seeds **fine** | `Audit: 74 write attempt(s), **55,838 rows**` — `users` 928, `personas` 704, `taxonomy` 42,790, all `ok` |
| the reload misses | `⚠ sentinel reload failed (non-fatal — a non-AI-sim run is unaffected)` (`:540`), and **again** on the restore leg (`:1617`) |
| every org-scoped call is refused | trace bodies: `{"errors":[{"message":"forbidden","path":["organizationMembers"]}],"data":null}` — **HTTP 200** |

**Three enforcement sites, all denying**, from the traces: `forbidden` (lowercase) on
`organizationMembers` / `membershipsCount` / `organizationLicense` / `membershipTagSummary` /
`addTag`; `Forbidden` (capitalised) on `organizationFeatureUsageJobSimulations` /
`organizationFeatureCreditsJobSimulations`; and `user does not have permission to access` on
`enableOrganizationSettings`.

**The partition that proves it.** All **15 failing** Playthroughs are **org-scoped**; all **15
passing** are **user-scoped** (onboarding ×5, profile ×4, skill-paths ×2, studio ×2, ai-readiness
member-progress, ai-sim org-feature-blocked). **Both negative controls also failed** — a fact the
iter-15 close did not carry: `.last-run.json` lists **17** failed ids, not 15.

## D77 — Both of iter-15's candidate causes are REFUTED, and so is a third the traces killed.

- **Contention — refuted.** `forbidden` is a deterministic verdict, not a latency. The load1
  measurement iter-15 correctly demanded before attributing to load is **not recoverable**
  (`loadwatch.log` stops 09:14Z; the batch ran 10:57–11:07Z) — and it is **not needed**, because the
  evidence refutes load rather than quantifying it. The rule *"never attribute a red to load without
  measuring load1"* is satisfied in the direction that matters: nothing is attributed to load.
  **The four "timeouts" were the same defect wearing a different face** — `member-tag` and
  `tag-create` were waiting for the `Create Tag` modal to close, and it **resolved visible 122–123×
  across 60 s** because the mutation underneath was refused; `role-create` never navigated;
  `activity-drilldown` waited on a row of a refused query. A slow box does not do that.
- **A partial `pt-world` seed — refuted.** 55,838 rows landed across 32 seeders with `isolation:
  clean`. The data was there; the reads were refused before reaching it.
- **"The newest platform moved a table" — stays refuted**, and now with a mechanism: **no
  `SQLSTATE 42P01`** because the queries never reached a table. Correlation was the starting point,
  not the report.

## D78 — The missing `sentinel` container is CORRECT. The dangling reload is the defect.

Worth stating separately because the first reading was the opposite. `demo-3` runs **10** containers
and none is `sentinel`; `docker-compose.injected.yml` declares no `sentinel` service and no
`AUTHORIZATION_ADDRESS`; `backend`'s `AUTHORIZATION_ADDRESS` is **empty**. Every one of those facts
is **expected on v11.0** and none is a defect. Had this iter stopped at "sentinel is missing → that's
the bug", it would have filed a platform regression against a deliberate platform merge.

`app/internal/sentinel/` is the fold, and it makes **sentinel the 8th service merged into `app`** —
a change this corpus documents nowhere. `CLAUDE.md` still calls sentinel a Tier-1 always-on service
and still calls `backend → sentinel` *"the only cross-process Connect-RPC edge left in a local
stack"*. Routed as a corpus finding, not fixed here.

## D79 — The fix is the platform's own mechanism, and it was proven live before being written.

`app` subscribes to the Redis Pub/Sub channel **`sentinel:policy:invalidate`**
(`watcher.go:56`, `DefaultInvalidationChannel`) and calls `LoadPolicy()` on any payload that is not
its own instance-id (`:129-137`). So the correct post-seed action is a **PUBLISH on the stack's own
redis**, not an HTTP call to a container that no longer exists.

**Proven live on `demo-3` before a line was changed** (11:26:44Z, `load1` 3.12):

```
$ docker exec demo-3-redis-1 redis-cli PUBLISH sentinel:policy:invalidate stackseed-iter16-probe
1
$ docker logs demo-3-backend-1
11:26:44 INFO policy changed elsewhere, reloading component=sentinel.watcher from=stackseed-iter16-probe
```

`PUBLISH` returning the **subscriber count** is a strictly better signal than the RPC's 2xx: `>= 1`
is positive proof the enforcer was reached. Shipped in all **three** sites (`run-playthroughs.sh`,
`restore-presenter-world.sh`, `up-injected.sh`) with the RPC — and up-injected's `docker restart`
rung — retained **beneath** it, so a stack built from a pre-v11.0 ref is unaffected. Non-fatal by
contract, but the failure text no longer understates its blast radius.

**Side effect, deliberate and benign:** that probe reloaded `demo-3`'s enforcer from the DB, where
the restored presenter world's **591 g2 + 591 g3** grants sit. It is a repair if the cache was stale
and a no-op if it was not. Nothing was torn down, reseeded or reset.

## D80 — Zero platform edits, and the fences stayed green without being edited.

The fix drives a channel `app` documents and subscribes to; no platform repo was touched. Fences
(1)–(4) of `TestRunnerSafety` and the `up-injected` control-plane fence
(`test_frontend_build.py:1966`) all stay green **untouched**, because the RPC rung was *retained*
rather than replaced. Two fences were **added**, mirroring the existing intent rather than relaxing
it: (5) the runner must drive the invalidation channel, (6) it must do so **per-stack**
(`${STACK}-redis-1`), which is the isolation rule (3) already enforces for the RPC. 105
demo-stack tests + the full `playthroughs` suite pass; `shellcheck -S warning` clean on both edited
e2e scripts.

## D81 — Scope: the orchestrator's ruling expanded this iter mid-flight, deliberately.

The iter opened planning **attribution only** (`overview.md` Phase A–D). The user's ruling arrived
mid-iter — *the 15 reds are not deferrable; "unattributed" is not an outcome* — which makes the
**fix** part of planned scope, not creep. Recorded so the close status grades against the expanded
scope. What is **not** in this iter, and is routed rather than dropped: **proving the fix live** on
a demo stack and a dev stack. That needs a bring-up, which needs a fresh slot, and it is iter-17's
whole job.
