**Type:** tik — under `TOK-01`, carrying `TOK-02`'s `TIK-C` end state.

# iter-19 — converge onto a stack the FIXED tooling built

## Phase A — re-pin, and check the FEATURE, not the tag

Consumption clone `stack-demo/rosetta-extensions` re-pinned **`fast-build-m258-iter-16` →
`fast-build-m258-iter-18`**. Rung zero was then done the way M236 says it must be — against files,
not against a tag name:

| what must be there | where | present |
|---|---|---|
| the batch gate, hooked | `up-injected.sh:2871` | ✅ |
| invalidation after set-dress | `up-injected.sh:2501` | ✅ |
| invalidation after reset-to-seed | `run-playthroughs.sh:171` | ✅ |
| invalidation after the restore | `restore-presenter-world.sh:121` | ✅ |

## Phase B — the bring-up

`./up-injected.sh 4 --no-public-host`, started **13:24:21Z**, `macmini`, `load1` **2.31–2.50**
throughout, 180 GiB free, `demo-3` resident and untouched.

`--no-public-host` is required rather than preferred: `--public-host` is default-on and **turns the
batch gate off on its own host** (`D84`), so a bare `/demo-up 4` would have left the thing under test
unrun and still printed a clean bring-up.

The pin guard fired its **freshness warning** — `clones.pin.json` pins `0c91421`, platform is on
`766df6c` — which is the newest-mains state, correctly reported and non-fatal. `make init` named
**`app`, `next-web-app`, `studio-desk`** and no `sentinel`: the 3-repo clone set observed from the
bring-up rather than read from a doc.

## Phase C — the verdict

```
reconciled 31 use cases: {"failing": 0, "passing": 30,
                          "unimplementable-without-platform-edit": 0, "unimplemented": 1}
BATCH GATE: GREEN — red set EMPTY on demo-4.
batch 138s · restore 6s · presenter world restored.
UP, and every journey verified.
```

`batch-gate.json` → `verdict: green` · `red_count: 0` · `red_set: []` · `runner_exit: 0`.
`autoverify.json` → `green: true` · `warnings: 0`. *"all 12 cockpit seats resolve in the 35-identity
roster."*

⚠️ **Not a clause-3 measurement, and `D100` states the arithmetic**: the ~290 s cycle was a
**warm-cache** build (`#7 CACHED … #11 CACHED`, no image-export leg) on the quietest box of the
milestone. Clause 3 stays **NOT MET**; its waiter stays disarmed.

## Phase D — teardown, in the mandatory order

Heartbeat written naming the stack and the reason, **then** `rosetta-demo down 3 --purge` at
**13:32:28Z** — after `demo-4` had already returned an empty red set. `rc=0`, network removed, data
purged. At no point was the box without a working stack.

## Phase E — the survivor, re-verified after the teardown

| surface | port | result |
|---|---|---|
| presenter cockpit | **47700** | **200** |
| next-web | 43000 | 307 (login redirect) |
| studio-desk | **49000** | 302 — *not* 49100; see `D103` |
| backend health | 48082 | `"OK"` |

**`http://localhost:47700` is the URL to validate through.**

10 containers, one stack, `docker ps` names nothing but `demo-4-*`.

## Close — 2026-08-12

**Outcome:** **`END-M258-one-stack` re-established on the correct stack.** Exactly one stack is up —
**`demo-4`**, built by the **fixed** tooling (`fast-build-m258-iter-18`) from the newest platform
mains — and it proved itself in the same command: **`red_set: []`, `runner_exit: 0`, 30/31 passing,
`autoverify green: true / warnings: 0`, 12 of 12 cockpit seats**, ending *"UP, and every journey
verified."* `demo-3` — the stack the **buggy** tooling built, whose own `batch-gate.json` still read
`verdict: red, red_count: 15` — was torn down **after** that verdict, never before, and the survivor
was re-probed afterwards on all four surfaces. Cockpit: **`http://localhost:47700`**.

**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — the milestone's gate closed by user ruling (`D52`). **Clause 3 remains NOT MET and is
explicitly not recorded as met**: this cycle was warm-cache on a quiet box (`D100`), which is not a
clean cold measurement and is not offered as one.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n *(2 tiks)* — (6) protocol-stop: n — (7) budget-exhausted: **y** *(between iters,
tree clean — a corpus-wide sweep with a full-suite attribution plus a cold bring-up, verification and
teardown spent the session; every remaining route needs either fresh host time or its own rext change
with mutants)* — Outcome: **exit-7**

**Decisions:** D98–D103

**Side-deliverables:**

- **`anthropos-sentinel:latest` is still on the box** (`D102`) — an image that outlived its service,
  the physical counterpart of the corpus drift iter-18 swept. Recorded, not pruned.
- **A second data point for `FIX-M258-iter14-purge-leaves-276MB`** — `stacks/demo-3` went 2.1 GB →
  **131 MB** across a full `--purge` (`D101`).

**Routes carried forward:**

- **`ROUTE-M258-iter19-orphan-images-outlive-their-service`** (net-new, `D102`) —
  `anthropos-sentinel:latest` plus five `:probe` leftovers (`m257-l1-hiring`, `m257-l1-next-web`,
  `m257-old-next-web`, `m257-warmup-next-web`, `m258-studio-desk`). **Price the shared layers before
  quoting any of it as reclaimable** (`D53`).
- **`ROUTE-M258-iter18-g1-reads-host-profiles-as-compose-profiles`** — unchanged, still open; the one
  RED fence on the box, and it is the fence's defect, not the corpus's.
- **`ROUTE-M258-iter18-app-row-anchors-are-at-2035f9a`** — unchanged, mechanical, ~7 anchors.
- **`ROUTE-M258-iter19-studio-desk-frontend-port-is-not-published`** (net-new, `D103`) — a demo stack
  publishes studio-desk's backend port only; any doc or probe naming `9100+offset` on a demo is
  naming a port nothing binds.
- Unchanged and **not** re-verified this iter: `REPORT-M258-iter17-public-host-default-skips-the-batch`
  (re-observed in effect — it is why `--no-public-host` was used) ·
  `REPORT-M258-iter17-dev-ui-images-stay-pre-L1-fat` · `ROUTE-M258-iter17-batch-gate-has-no-dev-opt-in` ·
  `ROUTE-M258-iter17-registry-is-empty-while-a-stack-is-up` · `FIX-M258-iter14-purge-leaves-276MB` ·
  `TARGET-M258-iter13-browser-only-deps` · `SETTLE-M258-iter13-studio-desk-cold-time` (**still
  unmeasured**; this cycle was warm-cache, so it yielded nothing here either) ·
  `ROUTE-M258-iter13-dockerfile-not-in-cache-key` ·
  `ROUTE-M258-iter15-compose-down-cannot-parse-an-older-stack` (**did not reproduce** — `demo-3` was
  built at `766df6c`, so its compose parses; the failure was specific to a stack older than the
  sentinel deletion).

**Lessons:**

- **Check the feature, not the tag.** The pin said iter-18; what mattered was that four specific
  call sites existed in the clone. They did, and the check took thirty seconds.
- **A verdict before a teardown is not ceremony.** The order cost nothing here because the gate went
  green — but the escalation condition was written first, and it would have bound.
- **A warm number is not a fast number.** ~290 s looks like a clause-3 pass and is not one; the
  missing leg is the largest one. State the cache state with every cycle time.
- **A teardown between two same-ref stacks reclaims almost nothing in Docker** — the reclaim was on
  the host tree, which `docker system df` cannot see at all.
- **A probe that names an unpublished port reports a dead surface.** The control was one command
  away and settled it immediately.
