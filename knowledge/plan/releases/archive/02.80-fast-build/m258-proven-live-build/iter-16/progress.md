**Type:** tik — under `TOK-01` (measure the composition before engineering it).

# iter-16 — attribute the 15-red batch verdict

## Phase A — retroactive attribution from the captured artifacts

The batch left a full artifact set per failure (`error-context.md`, `trace.zip`, screenshot, video)
under `playthroughs/e2e/test-results/`. **Nothing had to be re-run to attribute the reds**, which
matters because the load1 series needed to test the contention hypothesis stops at 09:14Z while the
batch ran 10:57–11:07Z — that measurement is gone for good. The artifacts are not.

First sweep, all 17 failure dirs: **every one carries an empty-state marker** — `0 / ∞`,
`0 Members`, `No data`, `EMPTY` — *including all four "timeouts"*. And `.last-run.json` lists
**17** failed test ids, not 15: **both negative controls failed too**, which the iter-15 close did
not record.

The 4/11 split iter-15 routed turned out to be the wrong partition. The right one:

| cluster | n | shape |
|---|---|---|
| org-scoped **reads** return nothing | 11 | workforce ×4, ai-readiness ×3, assignment ×2, hiring, activity-drilldown (+ both negative controls) |
| org-scoped **writes** never complete | 4 | setting-toggle, member-tag, tag-create, role-create |

And the partition that named the cause before any trace was opened: **all 15 failing Playthroughs
are org-scoped; all 15 passing ones are user-scoped.**

## Phase B — the traces decided the fork

The fork that had to be decided and not closed over: rows never written (**seed-side**) vs rows
present but not returned (**platform-side**). The traces answer it directly.

All GraphQL calls returned **HTTP 200** — no 403 anywhere. The bodies:

```
{"errors":[{"message":"forbidden","path":["organizationMembers"]}],"data":null}
{"errors":[{"message":"forbidden","path":["organizationLicense"]}],"data":null}
{"errors":[{"message":"forbidden","path":["addTag"]}],"data":null}
{"errors":[{"message":"user does not have permission to access ","path":["enableOrganizationSettings"]}]}
```

**Neither branch of the fork.** The reads were **refused**, not empty — the silent-403 class
`corpus/ops/verification.md` names, arriving as 200 + a GraphQL error. Three distinct enforcement
sites (`forbidden`, `Forbidden`, `user does not have permission to access`) all denying is the
signature of a seat with **no grants at all**, not of one new permission being added.

The same read killed the two "timeout" reds outright: the `Create Tag` modal
**resolved visible 122–123× across 60 s** because the mutation beneath it was refused. That is not
a slow box.

## Phase C — root cause, from platform source

`stack-demo/platform` HEAD is `766df6c` — ***"chore(compose): remove sentinel service and related
configurations"***. Sentinel is **folded into `app`** at **v11.0**; compose says so at `:85`, and
`backend` reads the `sentinel` schema in-process via `SENTINEL_DB_CONNECTION` (`:25`).

Our three post-seed reload sites all drove the **standalone** sentinel's Reload RPC on
`8087+OFFSET`. Nothing listens there any more, so all three missed — logging
*"non-fatal — a non-AI-sim run is unaffected"*. `app/internal/sentinel/watcher.go:5-11` explains why
that description is catastrophic rather than narrow: an enforcer that is not told to reload *"keeps
serving the policy it loaded at boot, forever. Nothing detects it, nothing logs it."* The seeders
write grants with **raw SQL**, which is not on casbin's write path, so no watcher fires.

The chain, every link measured — see [`decisions.md`](decisions.md) `D76`. Backend's enforcer loaded
at **10:51:44Z**, the reset dropped **1182** grants at ~10:57, the world re-seeded **fine**
(**55,838 rows**, `isolation: clean`), the reload missed twice, and every org-scoped call was refused.

**`D78`: the missing sentinel container is CORRECT, not the bug.** Stopping at *"sentinel is gone →
regression"* would have filed a platform bug against a deliberate platform merge.

## Phase D — the fix, proven live, then written

`app` subscribes to Redis Pub/Sub `sentinel:policy:invalidate`. Proven on `demo-3` **before** any
code changed (11:26:44Z, `load1` 3.12): `PUBLISH … → 1`, and backend logged
`policy changed elsewhere, reloading`.

Shipped in all three sites, publish-first with the RPC (and up-injected's restart rung) retained
beneath as the pre-v11.0 path. `PUBLISH`'s subscriber count is positive proof the enforcer was
reached — a better signal than the RPC's 2xx. Fences (5)+(6) added; (1)–(4) and the up-injected
control-plane fence stay green **untouched**. Full `playthroughs` suite + 105 demo-stack tests pass;
`shellcheck -S warning` clean.

`rosetta-extensions` `fcdc651`, tagged **`fast-build-m258-iter-16`**, **pushed and verified on
origin** (`git ls-remote --tags`) — tagging is not publishing.

## Close — 2026-08-12

**Outcome:** **`ESCALATE-M258-iter15-batch-red-15` is ATTRIBUTED AND FIXED — 15 reds + 2 failed
negative controls are ONE defect in our own tooling, not fifteen product failures, not contention,
and not a moved table.** Platform `766df6c` folded **sentinel into `app`** (v11.0 — the **8th**
service merge, documented nowhere in this corpus); our three post-seed reload sites still drove the
deleted container's RPC and called the miss *"non-fatal — a non-AI-sim run is unaffected"*. It is
not: a stale enforcer refuses **every** org-scoped read and write with `forbidden` at **HTTP 200**.
Fixed by publishing to `sentinel:policy:invalidate`, the channel `app` itself subscribes to —
**proven live on `demo-3` before the code was written**, with the pre-v11.0 RPC retained beneath.
**Zero platform edits; no fence edited to go green.**
**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — closed by user ruling (`D52`); clause 3 remains **NOT MET** and is never recorded as
met. The clause-3 waiter stays disarmed (`D72`) — not re-armed here, since this iter ran no batch.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
*(the escalation is resolved, not escalated onward; the corpus finding is a route)* — (5)
cap-reached: n *(1 tik)* — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**

**Decisions:** D76–D81

**Side-deliverables:** `demo-3`'s enforcer was reloaded from the DB by the live proof — a repair if
the cache was stale, a no-op if not. Nothing torn down, reseeded or reset.

**Routes carried forward:**

- **`PROVE-M258-iter17-policy-fix-on-demo-and-dev`** (iter-17, the whole job) — the fix is proven at
  the *mechanism* level (publish → reload) but **not yet end-to-end**: nobody has re-run the batch
  and seen the 15 reds clear. Needs a **fresh slot**, never `demo-3`. Carries the user's ruling #2:
  prove the release's new functionality on a demo stack **and** a dev stack, per piece answering
  *does it apply to dev · is it wired there · if not, is that correct?*
- **`CORPUS-M258-iter16-sentinel-in-app`** (net-new, and large) — **sentinel is the 8th service
  merged into `app`** (`app/internal/sentinel/`, platform `766df6c`). `CLAUDE.md` still lists it as
  a Tier-1 always-on service and still calls `backend → sentinel` *"the only cross-process
  Connect-RPC edge left in a local stack"* with `AUTHORIZATION_ADDRESS=http://sentinel:8087`; the
  always-on floor is now **postgresql + redis** only. Also net-new and undocumented: the Redis
  invalidation channel, and `app`'s own `W1`/`W4` notes. Route to `/update-knowledge`.
- **`ROUTE-M258-iter16-compose-comment-outlives-its-block`** (small) — `docker-compose.yml:86-87`
  says sentinel *"stays defined above as the rollback target until M1103 decommissions it"*, but
  `766df6c` **deleted** the 30-line block. The rollback target the comment promises does not exist.
- Unchanged and re-verified **open**: `FIX-M258-iter15-hiring-under-set-dressed` (confirmed live —
  `autoverify.json` reads `green:false, warnings:1`) · `FIX-M258-iter14-purge-leaves-276MB` ·
  `TARGET-M258-iter13-browser-only-deps` · `SETTLE-M258-iter13-studio-desk-cold-time` (`D75`) ·
  `ROUTE-M258-iter13-dockerfile-not-in-cache-key` ·
  `ROUTE-M258-iter15-compose-down-cannot-parse-an-older-stack`.

**Lessons:**

- **Ask what the red set is a verdict on — then keep asking.** iter-15 correctly refused to read it
  as a broken product and located it on `pt-world`. That was right and it was not the end: the cause
  was in neither world, it was in the tooling that hands worlds to the enforcer.
- **A "non-fatal" branch's own description is a claim, and claims decay.** *"A non-AI-sim run is
  unaffected"* was true when written and false the day sentinel moved. The line kept printing,
  reassuringly, while it took a whole batch down. **Audit the reassurance, not just the assertion.**
- **The failing set's shape names the cause before any tool is opened.** 15 org-scoped red / 15
  user-scoped green is a partition, and a partition points at a boundary — here, authorization.
- **Read what the platform says about itself.** `watcher.go`'s header described this exact failure
  mode, in advance, in the platform's own words. The evidence was authored before the incident.
- **A missing thing is not automatically a broken thing.** No sentinel container looked like the
  defect and was the *design*. One more file — compose line 85 — separated them.
