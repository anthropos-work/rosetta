---
iteration_type: tik
iter_shape: standard
status: closed
opened: 2026-07-29
---

# iter-17 — clause 3's landed half: the assign-tags Playthrough that four attempts could not reach

**Type:** tik · **Active strategy:** `TOK-01` move 3 ("org-admin … 4 mutating Playthroughs")

## Step 0 — re-survey

Both trees clean; `demo-2` up (16 containers); fixture `99e2f315`. Negative controls **21 of 24**.
Clause 3's landed half: **org-admin 2 of 4**, onboarding 1 of 5.

Of the two org-admin `TODO`s, `org-admin.members.UC1` was chosen over `org-admin.roles.UC1`: iter-05 left it
with **two specifically named untried routes** and a sanctioned verdict fallback, where the roles one needs a
live-LLM `Generate` leg (studio-lane budget) plus a product-defect report. Bounded beats open-ended.

## Cluster / target identified

`org-admin.members.UC1` — the assign-tags WRITE, `playthrough: TODO` for thirteen iters behind a wall iter-05
measured in four parts (the bulk-action dropdown stays mounted over the modal and swallows pointer events;
`Escape` closes the modal not the dropdown; an in-modal outside-click leaves 9 menu items open;
`check({force:true})` flips the DOM without driving antd's state).

## Hypothesis

iter-05 named two candidates (the `<label for=…>` element; a dropdown-free route). **A third, unlisted one is
stronger on first principles: a pointer-interception overlay cannot block a KEY event.** Probe all three.

## Expected lift

org-admin **2 of 4 → 3 of 4**; one more mutating Playthrough with a cross-surface read-back.

## Phase plan

A — probe the three routes. B — implement whichever works, in the page object. C — mutate: skip the
assignment (the final must go RED). D — full suite ×3 cold reset-to-seed + modules. E — close, commit, tag.

## Escalation conditions

If none of the three routes works, declare `unimplementable-without-platform-edit` with the evidence (1 such
entry is well inside the re-scope trigger's `> 3`).

## Acceptable close-no-lift outcomes

A measured refutation of all three routes plus the verdict.
