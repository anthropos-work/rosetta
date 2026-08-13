---
iter: 15
milestone: M258
iteration_type: tik
status: closed-fixed
created: 2026-08-12
---

# iter-15 — `TIK-C`: `END-M258-one-stack`

**Type:** tik · **Active strategy:** `TOK-02`, `TIK-C` — the milestone's **binding end state**, and
simultaneously the largest **Class A** reclaim available.

## Step 0 — re-survey

Three stacks resident at open: `demo-1` (11 containers, mine), `demo-2` (11, the user's), `anthropos`
dev (5, the user's). Target state: **exactly one**, built with the new mechanism from the newest
platform repos. Unchanged from `D57`; nothing has absorbed it.

## Cluster / target identified

`D57`, verbatim from the user: *"by end of this milestone there is only one stack up, and it's built with
the new process/mechanism and the newest repos of the platform."*

## Hypothesis

A cold `demo-3` built from the re-pinned consumption clone (`fast-build-m258-iter-14` — carrying the
multi-stage studio-desk and the `-v` teardown) against freshly-pulled platform mains will come up and
**prove itself** via the batch gate; the other three can then stand down, reclaiming their images and
data.

## Expected lift

- **End state:** 3 stacks → 1.
- **Space:** the largest single Class A reclaim of the milestone — two stacks' images and data, including
  `demo-2`'s **pre-L1** pair (next-web 4.04 GB + hiring 3.94 GB) that the release's biggest win never
  reached. To be reported as a **`system df` before/after**, never from the SIZE column (`D53`).
- **Free settlement:** `ui_studio_desk` cold, against iter-05's 115.35 s — discharging
  `SETTLE-M258-iter13-studio-desk-cold-time` at zero extra host cost.

## Phase plan

A. Disarm the clause-3 waiter (a campaign firing mid-transition would tear `demo-1` down under us).
B. Re-pin the consumption clone; **verify the feature is present, not merely the tag** (M236).
C. Pull every platform clone to newest `main`.
D. **BUILD AND VERIFY `demo-3` FIRST.**
E. Only then tear down — heartbeat before each, naming the stack and why.
F. Measure the reclaim; settle the studio-desk cold time.

## ⚠️ Order is mandatory and is not an optimisation

**Build-and-verify first; teardown last.** The user must never be left without a working stack. If D
fails, **nothing is torn down** and the iter closes with three stacks up and a finding — that is a
correct outcome, not a failed one.

`demo-2` is explicitly *not* the stack to keep: it is on pre-L1 images, so the user's own stack never
received the release's biggest win. This is the one sanctioned exception to "never touch `demo-2`",
**only at the end, and only in this order.**

## Escalation conditions

If `demo-3` cannot come up green from newest `main`, that is a **finding about the platform**, and the
end state is not forced: keep the three stacks, report, escalate.

## Acceptable close-no-lift outcomes

A failed bring-up from newest `main`, fully diagnosed and with all three stacks left intact, closes this
iter honestly — the user keeps a working environment and learns his newest platform does not build.
