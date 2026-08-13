# M258 iter-14 — progress

**Type:** tik · **Active strategy:** `TOK-02`, `TIK-B` (Class A — zero coupling to time)

Measured 2026-08-12 10:52–10:58Z on `macmini`, `load1` ~20.

## Phase A — is `-v` safe on the plain path? (measured, not assumed)

`--purge` already passed `down -v --remove-orphans` and `D56` exonerated it. The plain branch
(`rosetta-demo:446`) passed a bare `down`, so the two anonymous bitnami volumes outlived every
non-purge teardown and every container recreate — `D55`'s producer, 178 volumes / 5.297 GB over five days.

The fix is one flag, but the flag is only safe if no **named** volume exists to be destroyed. Census of
every container in the project:

```
demo-1-postgresql-1  VOLUME  9f227c2e92bc…  ->  /docker-entrypoint-initdb.d     ANONYMOUS
demo-1-postgresql-1  VOLUME  1c5dd2836cb6…  ->  /docker-entrypoint-preinitdb.d  ANONYMOUS
(every other mount in the stack is a bind)
```

**Zero named volumes** (`D69`). And the database itself is a host bind mount, which `-v` never touches.

## Phase B/C — the fix and its fence

`rosetta-demo`'s plain branch now runs `down -v`, carrying the measurement that makes it safe. Fenced by
`test_down_plain_removes_anonymous_volumes`, which asserts **both** branches pass `-v`, that no bare
`"${comp[@]}" down ||` survives anywhere, **and the rationale sentence** — so a future named volume
re-opens the decision instead of silently making a plain teardown destructive.

`bash -n` clean; `tests/test_purge.py` **5 passed, 1 skipped**.

## Phase D — the other half, priced and deliberately not taken

`purge_data_dir` is scoped to `$stack/data` (G1 path-assert). **≈276 MB per stack survives a full
`--purge`** — `clones/` 220 MB, `bin/` 37 MB, the two fake-Clerk trees 18.5 MB (`D70`). Routed rather
than rushed: widening a `rm -rf` whose safety rests on a path-assert, immediately before the milestone's
binding end state, is the wrong trade. `TIK-C` tears a stack down anyway and measures it for free.

## Close — 2026-08-12

**Outcome:** The orphaned-volume leak is stopped **at its producer**, not just reclaimed after the fact:
the plain teardown now passes `-v`, so the two undeclared bitnami anonymous volumes die with their
container. Safety was established by a live census (**zero named volumes anywhere in a demo stack**;
the DB is a bind mount `-v` cannot touch), and fenced so a future named volume re-opens the decision.
Cost on the time axis: **zero** — the defining property of `TOK-02` Class A. The second half, the
**276 MB per stack** of host-side residue that survives `--purge`, is priced and routed rather than
rushed in ahead of `END-M258-one-stack`.
**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — achieved by user ruling (`D52`); clause 3 NOT MET and never to be recorded as met.
**Phase 5 grading:** (1) gate-met: n *(never, by ruling)* — (2) triggered-tok: n — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n *(2 tiks)* — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**

**Decisions:** D69–D70

**Side-deliverables:** none.

**Routes carried forward:**

- **`FIX-M258-iter14-purge-leaves-276MB-of-stack-dir`** (`D70`, priced) — `purge_data_dir` clears only
  `data/`; `clones/` 220 MB + `bin/` 37 MB + fakes 18.5 MB survive. Measure empirically at `TIK-C`'s
  teardown, then widen the purge scope deliberately (the G1 path-assert must widen with it).
- Unchanged: `ROUTE-M258-iter13-dockerfile-not-in-cache-key` ·
  `TARGET-M258-iter13-browser-only-deps-in-the-runtime-image` (~200–260 MB) ·
  `SETTLE-M258-iter13-studio-desk-cold-time`.

**Lessons:**

- **Fix the producer, not the symptom.** iter-11 reclaimed 5.297 GB; that was housekeeping. One flag at
  the teardown means the 5.297 GB cannot accumulate again — and it cost nothing on either axis.
- **A one-flag fix can still need a measurement.** `-v` is safe here *because* a demo declares no named
  volumes. That is a property of today's compose, so the fence asserts the reasoning, not just the flag.
- **Don't widen an `rm -rf` under time pressure.** The 276 MB is real and can wait one iter; the
  milestone's binding end state cannot.
