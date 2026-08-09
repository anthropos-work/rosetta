**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them.*

# iter-219 — the by-effect recorder perturbed its subject, and its own sealed rule caught it

## What this iter set out to do

`SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` has been open and re-listed **unchanged
through eleven harden passes**. `test_mutation_proof_cache_hazard_m257x.py` **proves** the hazard in
three arms — it is real on this interpreter, the mtime bump defeats it, clearing `__pycache__` does not
— and **enumerates nothing.** The mechanism is proven; the population has never been measured.
`§5` iter-184: *a fence's POPULATION is a registry too.*

## V1 — the static approach, refuted before the pivot (sealed at `33b3129`)

| heuristic | answer |
|---|---|
| test modules containing any `.write_text(` | **54** of 77 |
| `write_text` calls whose receiver does not look temp-ish | **369** |
| `write_text` calls with a `.py` path in the surrounding lines | **74** |

Three heuristics, three answers, **none of them a population** — and only **3 of the 54** handle mtime
at all. Sealed as a refutation *before* the by-effect recorder existed, so the pivot cannot be narrated
afterwards as the plan. iter-218's rule, one iter old: *a reading SAMPLES; a fence CENSUSES* — and the
census must be by **effect**, because no spelling of a call site tells you what it writes.

## ⚠️ V3 FIRED — the pre-registered stop condition, and it is the iter's result

> *"The instrumented run's pass/fail/skip counts must equal the uninstrumented run's. A recorder that
> perturbs its subject is refused, not corrected."*

The recorder ran the whole suite — **1,869 passed · 1 failed · 3 skipped in 26 m 02 s** against an
uninstrumented expectation of **1,870 · 0 · 3** — and **recorded 9,813 writes**. Exactly one test
flipped, and it reproduces in **0.19 s** on one module: `test_stack_registry.py` is **62 passed**
uninstrumented, **61 passed / 1 failed** instrumented.

### The mechanism, which is worth more than the census would have been

The recorder wrapped `builtins.open` and, to learn the final size, rebound `fh.close` to a closure that
captured `fh` — **a closure stored on the handle it captures.** The handle acquires a reference cycle
and is no longer released by refcounting. The subject writes its fixture as

```python
json.dump({"dev-1": "garbage"}, open(self.reg, "w"))     # test_stack_registry.py:321 — never closed
```

relying on CPython dropping the `TextIOWrapper` at statement end to flush it. Under the recorder the
flush never happens, the next reader sees an **empty** file, and a test about *malformed* input fails on
*empty* input.

> **An instrument that changes its subject's LIFETIME changes its subject's behaviour** — the sharper
> form of what the bytecode hazard says about caching: the subject you measured is not the subject that
> ships.

### And the non-perturbing half is the blind half — measured, not argued

Removing the `open` hook removes the perturbation (**62 passed**) **and the visibility**: the `Path`-only
recorder sees **0 of that module's 140 writes**, because the module writes exclusively through `open`.
**The half that perturbs is the half that measures.** That is why the repair is a redesign and not a
patch, and why it is routed rather than smuggled into this iter.

## The reading the perturbed run produced — reported, NOT claimed

Of **9,813** recorded writes: **1** is inside the repo (a `.yaml`), **0** are `.py` inside the repo, and
**0** match the exposed shape (existing `.py`, size-preserving, mtime not forced). Every mutation
control writes into a temporary directory.

**This is not a claimable measurement.** The instrument failed its own validity condition in the same
run, and *an instrument that states its own invalidity must not exit 0*. It is recorded because a
suppressed reading is worse than a disclosed one — but `SURVEY-M257x-h42` stays **OPEN**, now with a
method, a named obstacle and a designed fix instead of eleven identical re-listings.

## Scope, stated rather than implied (`§5` r60)

`/usr/bin/python3 -m pytest` (**pytest 8.4.2 / CPython 3.9.6**), **Python**, `stack-core` only.
Whole-section, **instrumented**: 1,869 passed · 1 failed · 3 skipped (26 m 02 s) — the single failure is
the perturbation itself. Post-change scoped: **144 passed / 1 skipped** across three modules (43 s);
`--ceilings` exit **0**, all three `exact +0` after re-pinning **559 → 563**. **No uninstrumented
whole-section run after the change.** No Go, no TypeScript, no non-`stack-core` Python section.

## Close — 2026-08-09

**Outcome:** the iter's own pre-registered validity condition fired. The by-effect write recorder
perturbs its subject; the mechanism is a reference cycle that defeats an implicit flush, reproducible in
0.19 s and now fenced by three arms. The static approach is refuted (54 / 369 / 74, none a population),
the exposed-write population remains **unmeasured**, and `SURVEY-M257x-h42` stays open — for the first
time in eleven passes with a method and a named obstacle rather than a re-listing.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — **this is the FIRST `closed-no-lift`
after forty-nine consecutive `closed-fixed`; the trigger needs three consecutive no-progress tiks, so
one cannot fire it** — (3) re-scope: n — a `closed-no-lift` with documented falsification does not count
toward it — (4) user-blocker: n — (5) cap-reached: n — **counted, not felt: iters 217, 218, 219 = three
tiks this run against a cap of five** — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-219-1` … `D-M257x-219-3` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none this iter. The recorder ships as a **diagnostic**, explicitly not as a
fence — that is the iter's finding, not a deliverable beside it.

**Routes carried forward:**
- `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` — **still OPEN**, and no longer a bare
  re-listing: the method is by-effect recording, the obstacle is named, the fix is designed (record the
  path at open time, `stat` once at process exit, never touch the handle).
- `SURVEY-M257x-iter219-write-recorder-must-not-touch-the-handle` — **NEW.** The redesign above, with
  the re-run cost stated (26 min instrumented).
- `SURVEY-M257x-iter219-a-test-fixture-depends-on-refcount-timing` — **NEW.**
  `test_stack_registry.py:321` passes only under CPython refcounting. Not wrong today, not portable,
  and it took an unrelated instrument to make it visible.
- All routes from iters 207–218 unchanged, plus the standing queue.

**Lessons:**
- **Pre-register the validity condition, not just the expected number.** V3 was a sentence about the
  instrument, and it is the only reason this iter has a result instead of a 9,813-row table nobody
  could trust.
- **An instrument that changes its subject's lifetime changes its subject's behaviour.**
- **When the perturbing half and the measuring half are the same half, the answer is a redesign** —
  and saying so is a finding, not a deferral.
