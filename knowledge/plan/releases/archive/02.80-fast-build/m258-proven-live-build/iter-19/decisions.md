# iter-19 — decisions

## D98 — `END-M258-one-stack` re-established on a stack the FIXED tooling built

Exactly **one** stack is up: **`demo-4`**, 10 containers, offset **40000**, single-box
(`--no-public-host`). `demo-3` is gone.

Built from the consumption clone at **`fast-build-m258-iter-18`** (re-pinned this iter from
iter-16, and the tag verified on origin earlier in the session), on the newest platform mains —
`make init` reported `app`, `next-web-app`, `studio-desk` *"already exists, skipping"* and **named no
`sentinel`**, which is the 3-repo clone set of `766df6c` observed from the bring-up rather than from
a doc.

**The verdict, from the stack's own artifacts:**

| artifact | value |
|---|---|
| `batch-gate.json` | `verdict: green` · `red_count: 0` · `red_set: []` · `runner_exit: 0` |
| reconciliation | 31 use cases → `passing: 30`, `failing: 0`, `unimplemented: 1` (the declared TODO) |
| `autoverify.json` | `green: true` · `warnings: 0` |
| cockpit seats | *"all 12 cockpit seats resolve in the 35-identity roster"* |
| the line it printed | **`UP, and every journey verified.`** |

Surfaces re-probed **after** the teardown, so the survivor is verified in its final state, not only in
its build state: cockpit `47700` **200** · web `43000` **307** · studio-desk `49000` **302** ·
`48082/api/health` **`"OK"`**.

**Pre-flight rung zero was done as a feature-present check, not a tag check.** The M236 failure is a
clone that does not carry the thing under test, so the assertion was made against files: the batch
gate is hooked at `up-injected.sh:2871`, and **all three** invalidation sites are present —
`up-injected.sh:2501` (set-dress), `run-playthroughs.sh:171` (reset-to-seed),
`restore-presenter-world.sh:121` (restore). Two of them printed
*"✓ policy invalidated via redis pub/sub (in-process enforcer reloaded — sentinel-in-app)"* during
this run.

## D99 — the order was honoured, and it is the reason to record it

Build → verify → **then** tear down. `demo-3` was left running and untouched until `demo-4` had
returned an empty red set and a green `autoverify`, with a heartbeat naming the stack and the reason
before the teardown command ran. At no point was the box without a working stack.

The escalation condition was written before the run and would have bound: a non-empty red set would
have been escalated to the user under `D-v28-3` **with `demo-3` left up**. It did not fire.

## D100 — ⚠️ this is NOT a clause-3 measurement, and the arithmetic says why

Cycle wall-clock **13:24:21Z → 13:29:11Z ≈ 290 s**, of which `batch_seconds` **138** and
`restore_seconds` **6**. That number must not be quoted against the 480 s gate.

**It was a warm-cache build.** `demo-3` had just built the same images from the same refs; the log
shows `#7 CACHED … #11 CACHED` and *"all images built"* without an image-export leg. `build-budget.md`
prices export/unpack at **46.2 %** of a cold cycle on the profiled host class, so a warm run omits the
single largest phase. `load1` **2.31–2.50** throughout — the quietest the box has been all milestone,
which flatters it further.

Clause 3 stays **NOT MET** and its waiter stays disarmed. What this run *is* evidence of is the thing
it was run for: **the composed bring-up ends in a green batch gate on a stack built by the fixed
tooling, cold on the database side** (`--purge` had wiped `demo-3`'s data dir; `demo-4` initdb'd a
fresh cluster, migrated, replayed and seeded).

## D101 — the teardown reclaimed almost no Docker space, and that is the correct result

Measured `docker system df` before/after: **Images 14.02 GB → 14.02 GB**, containers 12.29 MB →
5.845 MB, build cache 27.97 → 27.36 GB. Host free 180 → **181 GiB**.

`demo-3` and `demo-4` were built from the **same platform refs with the same Dockerfiles**, so their
images are the same layers. Removing *"this stack's images"* removed tags whose content another
running stack still uses. **A teardown reclaims what is unshared, and between two same-ref stacks that
is nearly nothing** — the `D53` family again (`docker images` SIZE overstates ~5× because layers are
shared), reaching the *teardown* side of the ledger rather than the listing side.

The real reclaim was on the **host tree**, which `docker system df` cannot see (`D53` sibling):
`stacks/demo-3` **2.1 GB → 131 MB**. That residue is `D70` — `purge_data_dir` is scoped to `data/`, so
`clones/`, `bin/` and the fake-Clerk tree survive `--purge`. Predicted ≈276 MB, measured **131 MB**
here, lower because parts of that tree are shared or absent on this slot. **`FIX-M258-iter14-purge-leaves-276MB`
stays open, with a second data point.**

## D102 — `anthropos-sentinel:latest` is still on the box, and it outlived the service by design

Net-new, small, and exactly on this release's theme. `docker images` still lists
**`anthropos-sentinel:latest`** — built by the dev stack before platform `766df6c` deleted the
service. Nothing references it: there is no `sentinel` compose service, no `repos.yml` entry, and no
stack that could start it.

Not deleted here. It is **~0 reclaimable** as a shared-layer question until measured, and this iter's
job was the end state, not a prune — but it is the physical counterpart of the corpus drift iter-18
just swept, and it is the kind of thing a reader finds on their own box and cannot explain.
Recorded as `ROUTE-M258-iter19-orphan-images-outlive-their-service`, together with the five leftover
probe images (`m257-l1-hiring`, `m257-l1-next-web`, `m257-old-next-web`, `m257-warmup-next-web`,
`m258-studio-desk` — all `:probe`).

## D103 — studio-desk publishes ONE port, and the 9100 probe was wrong on both stacks

A probe of `49100` returned **000** and read like a broken UI tier. It is not: `demo-4-studio-desk-1`
publishes **`0.0.0.0:49000->9000/tcp` only**, and answers **302** there. The control settles it —
`demo-3`'s `39100` returned **000 as well**, on a stack that had been serving all day.

Recorded because the failure mode is *"the surface is down"* when the truth is *"the probe named a
port nothing publishes"*, and this corpus has paid for that distinction before. The demo compose
publishes studio-desk's **backend** port; the vite frontend port is not exposed on a demo stack.
