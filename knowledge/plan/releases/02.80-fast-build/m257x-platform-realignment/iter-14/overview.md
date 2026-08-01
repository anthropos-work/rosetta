---
milestone: M257x
iter: 14
iteration_type: tik
status: in-progress
opened: 2026-08-01
---

# iter-14 — cold cycles against origin HEAD, now that one can start

**Active strategy reference:** `TOK-01: instrument first, then follow` — step 5/5, *"prove it cold."*

## Step 0 — Re-survey

- Platform `stack-demo/platform` is at **origin HEAD `2adcf71`** (fast-forwarded in iter-12; the gate's first
  clause names origin HEAD and iter-12 nearly spent three cycles 3 commits behind it).
- rext consumption clone + `.agentspace/rext.tag` both at **`fast-build-m257x-iter-13`**, verified on origin.
- `docker compose config` at that ref: **RC=0**, 16 services. Before iter-13 it was RC=1.
- No stack is up; iter-12 tore `demo-1` down with `--purge`.

## Cluster / target identified

Clause 1. **And a live control iter-13 does not have.** `docker compose config` proves the project is
*valid*; it does not prove the stack *comes up*. iter-13 changed the SSR origin, the readiness probe's port
**and path**, the tailscale front list and the clone pin — none of it exercised against a running stack. The
standing pattern of this session is that the real finding shows up in the live run, not the targeted work.

Two open items ride along on any cold cycle and cost nothing extra to observe:
`CHECK-M257x-demopatch-pristine` (the platform just changed 5 files, `docker-compose.yml` among them) and
`FIX-M257x-iter13-freshness-vs-origin`.

## Hypothesis

A cold `demo-down --purge` → `demo-up` now gets past compose validation and reaches
`autoverify green:true / 0 warnings`. If it does not, it fails at a **named, measured** point that the
re-point did not reach — which is the more likely and more useful outcome.

## Expected lift

Clause 1 goes 0 → N cycles, N ≥ 1. **Three consecutive greens is the clause; this iter closes on whatever N
it measures**, exactly as iter-12 did at N=0.

## Phase plan

1. `demo-down 1 --purge` → `demo-up 1`, cold. Heartbeat throughout — ~18 min, mostly waiting.
2. Read `autoverify.json` **and** the transcript; do not read the verdict alone (iter-11's whole lesson was
   that two vantages of the same verifier disagreed for five hours unnoticed).
3. Repeat while budget allows, up to 3.
4. Close on the measured count.

## Escalation conditions

- A failure needing platform source → `demopatch`, never a platform edit.
- A **second** platform commit invalidating the attempt → **occurrence 2 of 2**, the `re_scope_trigger`
  fires, stop and escalate.

## Acceptable close-no-lift outcomes

A cycle failing at a named mechanism is this iter's deliverable at N=0 — the same standard iter-12 closed on,
and the four highest-value iters of this milestone all closed that way.
