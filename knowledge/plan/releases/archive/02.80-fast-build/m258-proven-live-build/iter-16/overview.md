---
iter: 16
milestone: M258
iteration_type: tik
status: in-progress
created: 2026-08-12
---

# iter-16 — attribute the 15-red batch verdict

**Active strategy reference:** `TOK-01` (measure the composition before engineering it). `TOK-02` (space
by coupling-to-time) is standing but **not** this tik's controlling strategy — the target is the
`D-v28-3` escalation, which is a measurement question, not a space one.

## Step 0 — re-survey (mandatory)

`ESCALATE-M258-iter15-batch-red-15` re-verified **OPEN** at iter open: `last-report.json`
(mtime 13:07Z) reads `failing: 15 / passing: 15 / unimplemented: 1`, and `.last-run.json` lists
**17** failed test ids — the 15 Playthroughs **plus both negative controls**, a fact the iter-15 close
did not carry. Target still meaningful, and richer than routed.

## Cluster / target identified

iter-15 routed the red set with **two candidate causes, neither confirmed** — contention (4 timeouts at
`load1` 26–33, `retries: 0`) and a partial `pt-world` seed (11 data-shape assertions) — and forbade
attributing to load without measuring `load1` at the moment it happened. That measurement is **not
recoverable**: `loadwatch.log` stops at 09:14Z and the batch ran 12:57–13:07Z.

**So the load hypothesis cannot be settled forward, and does not need to be.** The batch left
per-failure artifacts — `error-context.md` (page snapshot at failure), `trace.zip` (network + console),
`test-failed-1.png`, `video.webm` — for all 17. Those are a **retroactive** record of what the browser
saw, and they are readable on a loud box at zero host cost (`L5`'s lesson, applied to attribution).

## Hypothesis

The 4/11 split is the wrong partition. A first read of all 17 snapshots shows **empty-state markers in
every one** (`0 / ∞`, `0 Members`, `No data`, `EMPTY`) — *including all four timeouts*, whose call logs
show they were waiting on things a populated org would provide, or on a **write that never completed**
(the `Create Tag` modal resolved **visible 122–123×** across 60 s). A stuck modal is not a slow box.

The sharper partition is **two clusters, one candidate common cause**:

- **Cluster A — org-scoped READS return nothing** (≈11): workforce roster/funnel/succession/feedback,
  ai-readiness ×3, assignment ×2, hiring, activity-drilldown — plus **both negative controls**.
- **Cluster B — org-admin WRITES do not complete** (4): setting-toggle (switch never flips),
  member-tag + tag-create (identical stuck `Create Tag` modal), role-create (navigation never fires).

**Candidate common cause:** the manager seat's **org-scoped authorization does not resolve** in
`pt-world` — the silent-403 class `corpus/ops/verification.md` names. The roster snapshot supports it:
the shell renders the org (`Meridian Labs`) and the manager (`Morgan`), the page reaches
`/enterprise/members` and draws the full column set — but the sidebar carries **employee-only** nav
(no Enterprise section) and the count chip reads `0 / ∞`. Reads empty + writes refused + manager nav
absent is one shape, not three.

**Every passing Playthrough is user-scoped** (onboarding ×5, profile ×4, skill-paths ×2, studio ×2,
ai-readiness member-progress, ai-sim org-feature-blocked); **every failing one is org-scoped.** That
partition is the hypothesis's strongest support and its sharpest falsifier.

## The fork that must be decided (and must not be closed over)

If the org-scoped surface is empty because rows were never written → **seed-side**, a `pt-world` defect.
If rows exist and the newest platform will not return them → **platform-side**, a genuine regression on
the newest mains and the finding this milestone exists to catch. **No `SQLSTATE 42P01` anywhere**, so
"a table moved" stays unproven and is not adopted.

`pt-world` itself is **gone** — the batch's restore leg rebuilt the presenter world over it — so the
fork cannot be decided by querying the seed that failed. It is decided by the two records that survive:

1. **The traces.** `trace.zip` carries the network log. HTTP/GraphQL **403 / authz error** on the
   org-scoped queries ⇒ authz, not emptiness. **200 with an empty array** ⇒ the data was not there.
2. **`demo-3`, read-only.** It runs the **same newest platform mains** and measured healthy. If its
   enterprise members grid renders members, the platform's org-scoped read path is fine and the cause
   is `pt-world`-side; if it is *also* empty, the finding is platform-side.

## Expected lift

No metric moves — the gate is closed by ruling and clause 3 stays NOT MET. The deliverable is an
**attribution**: each of the 17 reds assigned to a named cause with evidence, or explicitly declared
unattributable with the measurement that would settle it.

## Phase plan

- **A** — retroactive attribution from the 17 captured artifact sets (snapshots + call logs). *(largely
  done at re-survey)*
- **B** — extract the traces; read the network verdict on the org-scoped queries. Decide the fork.
- **C** — corroborate against `demo-3` **read-only** (no teardown, no reseed, no writes).
- **D** — record the attribution, route what remains, close.

## Escalation conditions

- Traces show a **tenant leak** (a manager seeing another org's rows) → user-blocker, immediately.
- The fork resolves **platform-side** → that is a finding, recorded and escalated in the report; it does
  not by itself block the iter.
- Any step would require writing to / reseeding / tearing down `demo-3` → refuse, record as a HEADROOM-
  class refusal, and name what a fresh slot would cost.

## Acceptable close-no-lift outcomes

- The traces are unreadable or carry no network log ⇒ close with the attribution that the snapshots
  alone support, and state precisely what a re-run on a quiet box would settle. **Saying "cannot
  attribute" with the settling measurement named is a result**, per the run brief.
