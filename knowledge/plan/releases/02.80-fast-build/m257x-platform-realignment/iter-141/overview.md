---
iter: 141
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-141 — the last three adjudicated cross-references, and a pointer that names a retracted title

**Type:** tik
**Active strategy reference:** `TOK-08`. Closing the cross-reference remainder of
`FIX-M257x-iter135-adjudicated-live-defects` — the sample side, now that iters 138–140 have worked the
census side.

## Step 0 — Re-survey (all three live at HEAD `34d1170`, all three verified before planning)

| target | adjudicator | verified |
|---|---|---|
| `ai-readiness.md:18-20` — *"the **only** remaining dependency on `workforce` is the member directory … whose implementations **stayed** in `members.go`"* | `adj-B` P-3 | **confirmed at `app` `ad9f3c498`**: the interface declares **FOUR** methods (`manager.go:40-51`), and **`LevelsCount` lives in `internal/workforce/manager.go:90`, not `members.go`** |
| `clerk-integration.md:126` → `ant-academy.md:334` *"the `DEV_LOGIN_ENABLED` public-route pair"* | work list | **rotted +4** — `:334` is the **AI-proxy** row; `DEV_LOGIN_ENABLED` is at `:338` |
| `backend.md:13` → *"see the **M810 prod teardown is UNEVEN** bullet below"* | work list | the bullet exists but is **titled *"The M810 prod teardown has now LANDED for both"*** and its body **retracts** *"UNEVEN"* |

## Cluster / target identified

The third is the interesting one and it is a **new shape**: a cross-reference that names its target **by a
title the target itself has retracted**. `D-M257x-137-3` covered quoting a retracted *pin*; this is
quoting a retracted *name*. It is worse in one respect — the pointer still resolves, so no anchor fence
can see it, and the reader arrives at a bullet that opens by contradicting the sentence that sent them.

## Hypothesis

All three repair cleanly by **naming what the target now says** rather than what it used to be called or
where it used to sit.

## Expected lift

No `N` movement claimed (no reading). Deliverable: the cross-reference remainder of the adjudicated work
list closed.

## Phase plan

1. Repair the three, each re-derived at source (done at re-survey).
2. Gates: guard family + the anchor/citation scoped suites (**chosen by what changed** — `D-M257x-138-5`,
   which caught a self-inflicted defect at iter-140 before the commit).

## Escalation conditions

3rd unplanned line → tripwire. This is the **5th tik** of the session, so the cap fires at close.

## Acceptable close-no-lift outcomes

Any of the three proving already-correct is a result; all three were verified at source before this file
was written, and all three are live.
