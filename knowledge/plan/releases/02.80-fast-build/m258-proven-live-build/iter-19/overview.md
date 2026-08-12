---
iteration_type: tik
status: in-flight
milestone: M258
iter: 19
opened: 2026-08-12
---

# iter-19 — converge onto a stack the FIXED tooling built

**Type:** tik — under `TOK-01`, carrying `TOK-02`'s `TIK-C` end state (`END-M258-one-stack`).

## Step 0 — re-survey before targeting

Re-measured, not assumed:

- **One stack is up: `demo-3`**, 10 containers, offset **30000**, cockpit live at
  `http://localhost:37700` (HTTP 200). Its frontends carry `NEXT_PUBLIC_*=…localhost…`, so it is a
  **single-box** stack, not `--public-host` — the mode in which the batch gate actually runs.
- **`demo-3` was built by the BUGGY tooling.** Its own artifacts still say so:
  `batch-gate.json` → `verdict: red, red_count: 15, runner_exit: 1`; `autoverify.json` →
  `green: false, warnings: 1`. iter-16's live probe reloaded its enforcer by hand, so it *answers*
  today — but **anything that re-seeds it silently returns it to all-`forbidden`**, because the fix
  is in tooling the stack does not carry.
- The **consumption clone is pinned `fast-build-m258-iter-16`**, one tag behind
  `fast-build-m258-iter-18` (which is on origin, verified this session).
- Host is quiet: `load1` **2.64** (1-, 5-, 15-min all ≈ 2.6–2.8), 180 GiB free.

So the user's own requirement — *"only one stack up, and it's built with the new process/mechanism and
the newest repos of the platform"* — is **not** satisfied by `demo-3`. It is one stack, and it is the
wrong one.

## Active strategy reference

`TOK-01` (measure before engineering) for the verification half; `TOK-02` `TIK-C` for the end state,
whose **order is mandatory and is not an optimisation**: build and verify the new stack FIRST, tear
the other down LAST.

## Cluster / target identified

`END-M258-one-stack`, re-established on a stack built by the fixed tooling. iter-17 proved the fix
cold on a `demo-4` and then **tore that stack down**, keeping `demo-3` — so the proof exists and the
artifact does not.

## Hypothesis

A cold `up-injected.sh 4 --no-public-host` from the consumption clone at
`fast-build-m258-iter-18`, on the newest platform mains, ends in `BATCH GATE: GREEN — red set EMPTY`
and a presenter-usable world; `demo-3` can then be retired and the box left with exactly one stack,
correctly built.

`--no-public-host` is required, not preferred: `--public-host` is default-on and **turns the batch
gate off on its own host** (`D84`), so a bare `/demo-up 4` would leave the thing under test unrun.

## Expected lift

Binary deliverables:

- `batch-gate.json` on the new stack: `verdict: green`, `red_count: 0`, `runner_exit: 0`.
- `autoverify.json`: `green: true`, `warnings: 0`.
- 12 of 12 cockpit seats resolve; the presenter world is intact after the restore leg.
- Exactly **one** stack up at close, and it is the new one.
- The surviving stack's **cockpit URL reported** — the offset differs per slot, so the user needs it.

## Phase plan

- **A** — re-pin the consumption clone to `fast-build-m258-iter-18`; confirm the batch-gate files and
  the iter-16 invalidation fix are present in it (the M236 "the feature under test was not in the
  clone" shape).
- **B** — cold bring-up of the free slot, foreground-polled, never backgrounded-and-yielded.
- **C** — verify: batch gate, autoverify, cockpit seats, container census.
- **D** — heartbeat naming the stack and why, **then** tear `demo-3` down with `--purge`.
- **E** — re-verify the survivor after the teardown, report its cockpit URL, close.

## Escalation conditions

- Batch gate non-empty → **escalate the red set to the user** per `D-v28-3`; do NOT tear `demo-3`
  down. The user must never be left with only a stack that failed its own gate.
- Bring-up fails → `demo-3` stays up, untouched; report and stop.
- Headroom refusal → that is a RESULT, reported with `load1`, not routed around.

## Acceptable close-no-lift outcomes

A bring-up that fails for a named, measured reason with `demo-3` left intact and presenter-usable is a
complete iter — the end state is not worth a box with no working stack on it.
