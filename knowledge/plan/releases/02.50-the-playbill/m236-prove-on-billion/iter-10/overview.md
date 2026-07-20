---
milestone: M236
iter: 10
iteration_type: tik
status: closed-fixed
created: 2026-07-20
handler: REPRO-M236-iterTBD-cold-cycle
---

# iter-10 — the cold reset-to-seed reproduction (the last gate component)

**Type:** tik  ·  **Active strategy:** `TOK-01` (publish-then-prove) — the closing proof.

## Step 0 — re-survey

Three of four gate components are met (29/29 pairs · 65 academy cards, 0 chips · hero p95 3.15 / 2.71 s ·
0 platform edits). The outstanding one is **reproducibility on a cold reset-to-seed**, and it is the
component that retroactively validates the other three: everything measured so far was taken on a stack
that has been mutated in place across iters 07–09 (manifest re-exported twice, cockpit restarted three
times, rext clone re-pinned twice).

## Cluster / target identified

Handler `REPRO-M236-iterTBD-cold-cycle`. This is not a bug hunt — it is the proof that the whole chain
reproduces from nothing, which is what "prove on billion" means.

Known state to normalize, all self-inflicted by earlier iters:
- `billion`'s rext consumption clone is pinned to `playbill-m236-academy-course-route`; tooling has since
  been published through `playbill-m236-latency-tz-fix`. **Re-pin to the final tag before tearing down.**
- The cockpit currently binds `127.0.0.1` because `tailscale serve` had already claimed `:17700` at
  restart time (iter-07). A cold bring-up should restore the normal `0.0.0.0` ordering — and if it does
  not, that is a genuine finding about bring-up-vs-serve ordering, not a workaround to repeat.

## Hypothesis

The chain reproduces cold. Every fix this milestone landed is in published tooling consumed by tag, and
no fix depended on hand-mutated stack state — so a cold bring-up at the final tag should produce the same
readings without intervention.

## Expected lift

No numerator movement (29/29 already). The deliverable is the **gate**: the same numbers off a stack that
was built from nothing.

## Phase plan

- **Step 1** — re-pin `billion` to the final tooling tag; verify the M217 pin guard.
- **Step 2** — tear the stack down.
- **Step 3** — cold `/demo-up` (public-host default-on). **Expect 30–45 min.**
- **Step 4** — re-measure ALL THREE: the content-stories sweep (expect 29/29), the academy grid card
  count (expect non-zero, 0 chips), and a hero latency battery (expect p95 < 5 s).
- **Step 5** — grade the gate.

## Escalation conditions

- If the cold bring-up needs manual intervention to reach the measured state, the gate is **not** met —
  reproducibility is the component under test. Record what was needed and route it, do not paper over it.
- If a reading regresses cold, triage it as a real finding: the warm readings would then have been
  depending on hand-mutated state.

## Acceptable close-no-lift outcomes

A documented failure to reproduce, with the intervention required named precisely, is a complete iter —
it converts the gate from "believed met" to "met except for X", which is the honest state.
