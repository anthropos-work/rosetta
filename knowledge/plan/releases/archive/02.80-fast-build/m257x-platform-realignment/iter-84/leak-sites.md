# iter-84 — the three `repair_leak_guard` sites, adjudicated

`FIX-M257x-iter83-leak-guard-3-sites`. iter-83 measured that `repair_leak_guard.py` exits **1** on
iter-81's repair commit naming 3 sites, and that it was **not among the six guards iter-81 ran**. Those
sites were routed rather than repaired, because adjudication-before-repair is binding. This grades them.

**Result: 2 of 3 are real. 1 is a benign shared-vocabulary match — the guard's own documented
false-positive class.** A fence whose false-positive rate is undocumented is a fence nobody can decide
whether to trust (`test_repair_leak_guard.py`'s own words), so the benign one is recorded, not hidden.

---

## 1. `CLAUDE.md:285` — **UPHELD, blocker**

```
make up                # Build from local code and start (graphql profile)
```

**Predicate:** P4 — *"`graphql` is a live profile / the default"*.

**Measured** at `platform 0dab54d`: `Makefile:10` is `PROFILE ?= core`; the token `graphql` appears in no
`profiles:` key. Asking for it exits **0** and starts only the always-on floor.

**Why it is the worst instance in the tree:** it is a **runnable command with an explanatory comment**, in
a fenced Quick-Start block, in the repository's most-read file — and the same file warns about exactly
this hazard 16 lines later at `:301`. The corpus contradicts itself across 16 lines, and the wrong half
is the copy-pasteable one.

**Correction:** drop the parenthetical or say `(core profile — the default)`. Restate, do not re-anchor.

---

## 2. `corpus/ops/platform-alignment.md:1305` (was `:1249` before iter-83's rule-40 edit) — **UPHELD, blocker**

> *"At `app` `9d00a313` v1.367.0 — 56 commits and one working morning later — `STORAGE_RPC_ADDR` is read
> by `main.go` and by **none** of the three CLIs"*

**Predicate:** P9 — *"`STORAGE_RPC_ADDR` is read by `main.go` at `9d00a313`"*. This is the predicate
iter-81's commit message singled out as **load-bearing**, because it is the evidence that moved `storage`
off `mid-fold`.

**Measured at the ref the sentence itself names:**

```
$ git -C stack-demo/app grep -n 'STORAGE_RPC_ADDR' 9d00a313 -- '*.go'
9d00a313:internal/jobsimwiring/wiring.go:101:  // replacing the STORAGE_RPC_ADDR edge. Threaded in rather than constructed here
9d00a313:internal/storagens/callsites_test.go:189: // …the standalone storage service is scaled to 0 and STORAGE_RPC_ADDR is
9d00a313:main.go:451:                              // now — the standalone service takes no traffic and STORAGE_RPC_ADDR is gone.
rc=0     (positive control on the same pipeline: `func main` → 3+ hits)
```

**3 hits, every one a comment. Zero env lookups.** The `main.go` hit is a comment that says *in words*
that the variable is gone.

**The sentence describes a middle state that never existed.** At the older `b948604` it is read by
`main.go` **and** by all three CLIs (7 env lookups); at `9d00a313` it is read by **nothing**. "Read by
`main.go` but by none of the CLIs" is neither.

**This is the sharpest instance in the milestone of `platform-migration-status.md:74`'s own rule.** The
migration map states this correctly and ref-relatively; the **protocol doc that teaches the rule** states
it wrongly, in the very passage explaining why both sides must be recorded. iter-81 repaired the map and
left the teaching text — the leak, exactly as the guard names it.

**Correction:** *"…`STORAGE_RPC_ADDR` is read by **nothing** — 3 hits, all comments"*. Restate; the fact
did not move, it was deleted.

---

## 3. `corpus/services/messenger.md:122` — **REJECTED (benign)**

```
| Variable | Value (compose) | Description |
| `REDIS_STREAMS_INDEX` | `4` | Redis DB index for streams |
```

**Measured:** `docker-compose.yml` @ `0dab54d` sets `REDIS_STREAMS_INDEX=4` on `messenger` (`:180`) and on
`backend` (`:67`). The table's column header is **`Value (compose)`**, so `4` is exactly what the column
claims to hold.

**Why the guard fired, and why it is right to have fired.** iter-81 rewrote the *twin* table in
`roadrunner.md`, whose defect was real: it printed **compose-supplied** values under a header that read
as **binary defaults** (`10400`/`10401` against the binary's own `8080`/`8081`). The removed rows shingle-
match messenger's surviving rows because the two tables share their vocabulary. **`messenger.md` had
already made the distinction its header states** — it is not the same defect.

The neighbouring row is also correct on measurement: `REDIS_WORKER_INDEX` is *"set in docker-compose (=0)
but NOT read by the code"* — `git grep REDIS_WORKER_INDEX` in `messenger fa47850d` returns **0** Go hits,
while compose sets it at `:68` and `:179`. And `REDIS_STREAMS_INDEX` really is consumed at
`cmd/root.go:107` (`cmp.Or(os.Getenv(…), "2")`), plus a second site at `cmd/trigger.go:22`.

**One nuance worth recording rather than suppressing:** the code's *built-in* default is `2`, not `4`.
The table is honest because its header says `Value (compose)` — but the roadrunner repair shows how thin
that margin is, and any future edit that widens this header to *"Default"* re-creates the defect
iter-81 just removed one file over.

---

## What the three say together

**The one guard that asks *"did this commit FINISH?"* found 2 real leaks out of 3 candidates — a 67 %
precision on a commit that six other guards passed clean.** It was not run. The cost of not running it
was that both survivors stood at HEAD for two iterations, and one of them is the protocol doc
contradicting the rule it exists to teach.

Routed to **iter-85** for repair, by predicate, graded by `repair_reach_guard`.
