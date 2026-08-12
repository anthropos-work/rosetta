# M258 iter-04 — decisions

## D11 — the gate instrument refused, so the cycle was driven as an OPERATOR, and the split is deliberate

`buildbench run` **cannot execute a cycle** while headroom fails: `D-M255-1` makes the pre-rep assert
abort before rep 1, on purpose, because *a gate number measured on a host without headroom is not a
number*. Two bounded waiters hit their caps and the load **rose** in between (third-party: Spotlight,
the user's own `anima8` build, node) — this is the permanently-contended host, and *"do not wait for
quiet"* is the standing rule.

The alternative to stopping was available and is sanctioned by the same module: **`up-injected.sh`'s
pre-flights are advisory by design** (`buildbench.py:36`, *"never block a genuinely good bring-up on a
soft signal"*). The gate refuses to **quote**; it does not forbid **work**. So the cold cycle ran
directly, and the iter harvested exactly what contention cannot corrupt:

- **booleans** — single-box mode engaged, `.env.demo-1` 24 → 1, ISOLATION green on fresh images, the
  minted-host line present in the log, `BATCH_RC=0`;
- **counts** — 215 specs, 30 passing, 0 failing, **red set empty**;
- and **timings that are labelled contended and explicitly NOT gate numbers** (781 s / 129 s / 910 s).

**The reason this is not a loophole:** nothing measured this way is being offered as a gate reading. The
milestone's `exit_gate` needs a p50 over 3 cold cycles; this iter contributes **zero** samples to it and
says so. What it contributes is the answer to *"is the batch half 60 s or 600 s?"* — which is the
question `TOK-01` step 1 was written to answer, and which no ceiling-arithmetic could settle.

## D12 — 910 s over a 480 s ceiling did NOT fire the re-scope trigger, and firing it would have been wrong

The composed figure exceeds the gate by 430 s. It is tempting — and would look diligent — to escalate it
as the declared `re_scope_trigger`. **It does not qualify, on the trigger's own text:**

> *"If the composed **p50** exceeds **600 s** after **3 tiks**, split the suite into a fast smoke lane …
> and renegotiate the gate with the user."*

Three conditions, and this run satisfies none of them: it is **n=1**, not a p50; it is **one tik**, not
three; and both halves were taken in the one condition the release's own gate instrument refuses to
measure in. A trigger fired on a contended single sample would renegotiate a gate that has **not been
shown to be missed** — the exact inversion of the discipline that made M256 re-cut `D-v28-12` rather
than grade inside its own noise floor.

What the number *does* support is reported as an inference and flagged as one: the batch half is **small**
(~14 % of this cycle), so the ceiling is dominated by the bring-up half, and M257's proven **286.99 s**
gateable bring-up plus a ~129 s batch would land near **416 s** — inside 480 s. That is the first
evidence the composition is reachable, and it is **not** a measurement of the composition.

**Escalation deferred, not skipped.** If a *gateable* p50 over 3 cycles exceeds 600 s, the trigger fires
and goes to the user **with measurements attached**, which is what the milestone asked for.
