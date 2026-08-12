# M258 iter-11 — progress

**Type:** tik · The **space axis**, opened by the user's ruling. Measured 2026-08-12 09:57–10:00Z on
`macmini`, `load1` ~20, three stacks resident (`demo-1` 11 · `demo-2` 11 · dev 5).

## Phase A — the measurement, and the trap in it

```
TYPE            TOTAL   ACTIVE   SIZE       RECLAIMABLE
Images          31      22       23.83 GB   1.754 GB (7 %)
Containers      27      27       279.8 MB   0 B
Local Volumes   184     6        5.297 GB   5.297 GB (100 %)
Build Cache     123     14       27.45 GB   19.28 GB
```

⚠️ **`docker images` SIZE is not reclaimable size, and reading it as such overstates the win by ~5×**
(`D53`). The four `m257-*:probe` leftovers read **8.88 GB** in that column, which invites "delete them
and get 8 GB back" — but they **share 5 of 10 layers** with `demo-2-next-web`, which is in use. The
authoritative figure is `system df`'s **1.754 GB** reclaimable across *all* unreferenced images.
*Two numbers for the same thing means the definitions differ.*

## Phase B — the reclaim, ownership verified per target

**178 of 184 volumes dangling, 5.297 GB, 100 % reclaimable, 6 active.** Before touching anything, the
in-use set was enumerated and intersected with the dangling set:

```
dangling: 178   in-use: 6   OVERLAP: 0
  1c5dd2836cb6 / 9f227c2e92bc  <- demo-1-postgresql-1
  5b981a73fb21 / a4ec3ae4a52d  <- demo-2-postgresql-1   ← the USER'S
  94dbcf828170 / c3bf040c376e  <- anthropos-postgresql-1 ← the USER'S
```

**Reclaimed: 5.297 GB. Volumes 184 → 6. All three stacks intact afterwards** (`demo-2` 11 · dev 5 ·
`demo-1` 11).

Build cache was **deliberately not pruned** (`D58`) — 19.28 GB is reclaimable and the campaign is
armed; evicting it buys space by spending time, which is backwards for this release.

## Phase C — the producer, named

**Every Postgres container start creates two orphan-able anonymous volumes.** The bitnami image
declares three `VOLUME`s and compose binds only one:

```
bind   /bitnami/postgresql          <- stacks/demo-1/data/postgresql   (the real DB data — on the HOST)
volume /docker-entrypoint-initdb.d     <- ANONYMOUS
volume /docker-entrypoint-preinitdb.d  <- ANONYMOUS
```

Nothing in compose declares those two, so a `down` **without `-v`** — or any `compose up` that
*recreates* the container — orphans them. Three stacks × repeated bring-ups over five days = 178.

**`--purge` is NOT the culprit and this is measured, not assumed:** it runs `down -v --remove-orphans`
(`rosetta-demo:413`) and today's 3-rep campaign (08:19–09:02Z, three `--purge` teardowns) created
**zero** orphans — the newest orphan predates it, at 03:00Z. The leaking path is the **plain
`down`** (`:446`, no `-v`) and container recreation.

## Phase C′ — the axis `docker system df` cannot see

The DB data is a **host bind mount**, so Docker's numbers never included it. Post-teardown disk has a
third component nobody was measuring:

| host path | size |
|---|---|
| `stack-demo/…/demo-stack/stacks/` | **4.2 GB** (demo-1 2.2 · demo-2 2.0) |
| `.agentspace/…/demo-stack/stacks/` | **264 MB** (the authoring clone's leftovers) |
| `stack-dev/` | 3.8 GB |

And a live instance of the already-open `ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`:
**`stacks/demo-4/` still exists** for a stack that has not run in this milestone. Stack dirs survive
`--purge`; that is the same F-9 defect whose other symptom was the 24 accumulated Clerkenstein blocks
in iter-02.

Free disk: **173 GiB**.

## Close — 2026-08-12

**Outcome:** The space axis has its first measurement, and post-teardown was indeed the defect:
**5.297 GB in 178 orphaned volumes, reclaimed with zero overlap against the user's stacks**, and the
producer named exactly — two undeclared anonymous volumes per Postgres container, orphaned by the
non-`--purge` teardown path and by container recreation. `--purge` was **exonerated by measurement**.
Two things the headline numbers hide are recorded so nobody re-derives them wrong: image SIZE
overstates reclaimable ~5× via shared layers, and **host-side stack dirs (4.2 GB) are invisible to
`docker system df` entirely**.
**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — the user has ruled the milestone achieved (`D52`); clause 3 remains NOT MET and
unmeasured under load, and is never to be recorded as met.
**Phase 5 grading:** (1) gate-met: n *(never, by ruling)* — (2) triggered-tok: n — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n *(3 tiks)* — (6) protocol-stop: n — (7) budget-exhausted:
**y** *(between iters, tree clean — context spent; the new scope needs a fresh agent)* —
Outcome: **exit-7**

**Decisions:** D53–D58

**Routes carried forward — the new scope, unstarted and handed off:**

- **`TOK-02` — SPACE, a new goal of this milestone (user ruling, `D57`).** Must be authored as a
  strategy, not worked as a side route. Constraint, in the user's own framing: *as little disk as
  possible on `up`/`down`, **not** at the cost of time* — so **any cache policy must be argued on BOTH
  axes with measurements** (`D58`), `--filter until=24h` never `-af`.
- **`END-M258-one-stack`** (hard requirement, `D57`) — at milestone close there must be **exactly ONE
  stack up**, built with the new mechanism from the newest platform repos. Currently three.
  ⚠️ **Order is mandatory: build-and-verify the new stack FIRST, then tear the others down.** The user
  must never be left without a working stack. Heartbeat before each teardown, naming the stack and why.
  This is the one sanctioned exception to "never touch `demo-2`" — **only at the end, only in that
  order**. `demo-2` is *not* the stack to keep: it is on **pre-L1** images.
- **`TARGET-M258-studio-desk`** — the two axes converge here: **1.7 GB × 2 and 115.35 s cold**, and L1
  never touched it. On a post-L1 stack it is now the **largest** UI image (1.7 GB vs next-web's
  417 MB) *and* the largest UI time leg. Highest-value remaining item on either axis.
- **`FIX-M258-iter11-postgres-anonymous-volumes`** (net-new) — declare
  `/docker-entrypoint-initdb.d` and `/docker-entrypoint-preinitdb.d` in compose (bind or tmpfs), or
  make every teardown path use `-v`. Time-neutral, so it is free under `TOK-02`'s constraint.
- **`ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`** — upgraded from a curiosity to a **space**
  finding: 4.2 GB of host-side stack dirs, invisible to `docker system df`, with an orphan `demo-4/`
  still present.
- Unchanged: `SPLIT-M258-iter09-copy-vs-reindex` (needs one run) ·
  `ROUTE-M258-iter10-hand-rolled-path-filters` · `RATCHET-M257-literal-ceilings-breached` (honest debt
  +8 / +4) · the iter-02/07 routes.
- **Clause 3 is still armed** (`campaign-iter09/`, `fast-build-m258-iter-09` pin, waiter alive since
  09:40:37Z, `load1` minimum 14.21 vs a threshold of 5.0). Opportunistic bonus only.

**Lessons:**

- **The phase nobody measures is where the leak is.** Time had a budget, a baseline and a gate; space
  had a free-disk *floor* and nothing else — and that is precisely where 5.3 GB was quietly
  accumulating.
- **Verify ownership by intersection, not by naming convention.** The volumes were anonymous shas with
  no compose label; the only safe check was enumerating what every container actually mounts and
  proving the overlap with the delete set was **0**.
- **Exonerate as carefully as you accuse.** `--purge` was the obvious suspect and is innocent — the
  evidence is that a 3-rep `--purge` campaign created zero orphans while the newest orphan predates it.
- **A space number that ignores host bind mounts is not a space number.** The DB data — the largest
  per-stack artifact — is on the host and `docker system df` never sees it.
