---
iter: 264
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-10T15:31:00Z
closed: 2026-08-10T15:34:00Z
---

# iter-264 — the corpus half of `D-M257x-262-2`: the guide calls a build dependency optional

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
under the user's binding `D-M257x-256-1`.

## Step 0 — Re-survey before targeting

iter-262 proved the documented dev bring-up cannot build `backend` on a fresh clone; iter-263 corrected the
secret half of that finding down to an unrun check. **The `app/studio` half stands entirely**, and iter-263
noted why it is the more serious of the two: **no instrument would have caught it.**

The re-survey locates the corpus defect precisely. `corpus/ops/setup_guide.md:308-315`, *"Initialize CMS
Studio Submodule"*, currently reads:

> *"there is no `make init-studio` step for a local stack any more. To run the pipeline by hand, clone
> `anthropos-studio-room` yourself and run it directly."*

Every clause of that is true, and together they are **misleading in the one way that matters**: they frame
the clone as an **optional** step for someone who wants to run the generation pipeline, when in fact
**`make up` cannot build the `backend` image without it**. A reader following this guide top to bottom
reaches `make up` and gets `UP_EXIT=2`.

## Cluster / target identified

`FIX-M257x-262-dev-path-needs-the-studio-acquisition`, **corpus half only**. The tooling half (hoisting
`demo-stack/lib/studio.sh` so the dev path shares it) is a `rosetta-extensions` change requiring a tag and
a pin bump, which would destroy the frozen-pin control `D-M257x-258-1` holds for this milestone. It stays
routed; the corpus fix does not depend on it and is in the gate by construction (clauses 3/5 are rosetta).

## Hypothesis

The guide can be made to produce a working stack by stating the dependency where the reader meets it —
before `make up` — with the exact command, derived from the Dockerfile rather than from service names.

## Expected lift

A reader following `setup_guide.md` on a clean box reaches a running `backend`. No metric moves; the
corpus becomes correct where it was actively wrong, which under `TOK-08` is the deliverable.

## Pre-registrations — sealed in this iter's FIRST commit

| | claim | prediction |
|---|---|---|
| PR-1 | the guide's studio section is the **only** place the corpus mentions acquiring `anthropos-studio-room` for a local build | **AT RISK** — `grep -ci studio` returns 43 hits in this file alone |
| PR-2 | no corpus file states that `app`'s image **hard-fails** without `studio/` | **HOLDS** — iter-262 found the failure by hitting it, not by reading |
| PR-3 | `corpus/services/studio-room.md` describes the pipeline but not the build dependency | **HOLDS** |
| PR-4 | the corpus fences accept the edit (`repair-postcondition` green, no new adjudicated claim) | **HOLDS** |
| PR-5 | `CLAUDE.md`'s Studio-Room block — which iter-236 already corrected once — **also** omits the build dependency | **HOLDS** |

## Phase plan

- **Phase A** — seal. **Phase B** — measure PR-1…PR-3, PR-5 across the corpus. **Phase C** — edit the
  guide (and any other site the measurement names). **Phase D** — fences, grade, close.

## Escalation conditions

No platform edit; no tooling edit; no pin bump. If the measurement shows the claim is stated correctly
somewhere already, **say so and narrow the edit** rather than adding a duplicate.

## Acceptable close-no-lift outcomes

If PR-2 is refuted — some corpus file does state the hard dependency — that is a complete iter: the defect
becomes *discoverability*, not absence, and the fix changes shape accordingly.
