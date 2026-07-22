# M243 — Decisions

## D1 — The assign-WRITE surface + flow (discovered live, per P3)
The assign WRITE is `/enterprise/assignments` (next-web `apps/web`), a tabbed surface (AI
Simulations / AI Interviews / Skill Paths). We drive the **Skill Paths** tab because its per-row
read-back is the cleanest, most unambiguous LANDS proof:
- Each member row's Skill Path cell renders EITHER a link to the assigned skill-path OR an inline
  **"Assign Skill Path"** affordance (`memberSkillPathColumn` → `AssignmentContainer`) for a member
  with **no active skill-path org-assignment**. Clicking it opens the assign modal ("Assign Skill
  Path to <member>") targeting that ONE member.
- Pick a skill path (the public-catalog `SkillPathSelector` antd Select) + accept the pre-filled
  deadline (the Form remounts with `initialValues` on selection, so `deadlineDate` is set) →
  **"Assign"**.
- The write is `app` `mutationResolver.CreateOrganizationAssignments` →
  `AssignmentManager.BulkCreateOrganizationAssignments` → a real `public.organization_assignments`
  row. The backend **refuses a duplicate active assignment** ("Assignment already exists"), which is
  exactly why the target must be an **unassigned** member.

## D2 — The read-back (anti-toothlessness): affordance-count delta, not a closed modal
The release thesis is that a test passing while proving nothing is the enemy. The final assertion is
that the count of "Assign Skill Path" affordances drops by **exactly one** after the confirm. The
members query key is prefixed `['assignments', ...]`, and the org-assign mutation's `onSuccess`
invalidates `['assignments']`, so the members table **refetches in place** — the target member's cell
FLIPS from the assign affordance to the assigned title. That count can only drop if a real row landed
AND is read back through the real members query; a closed modal with no write leaves it unchanged (a
red test). Reinforced by a specific-member check (the target's row no longer offers "Assign Skill
Path"). We assert the OUTCOME state (P2), never an exact skill-path title / member name / score.

## D3 — Preconditions declared in lockstep (no new seed data)
No seed CODE change was needed: Org A (Meridian Labs, 40 members) already materializes both
preconditions — the base stories seed pre-assigns skill paths to only a handful (8 org-assignments;
**34/40 members unassigned** at capture), and the public catalog (22 skill paths) is set-dressed. But
per the "declared-and-enforced" discipline, UC1 now names `seed.preconditions: [public-catalog,
org-unassigned-member]`, and `org-unassigned-member` is added to `seed-worlds.yaml`'s pt-world
capabilities in lockstep — so a future seed change that assigns skill paths to ALL members surfaces at
`ptvalidate`-time (a SETUP break), not as a mysterious capability failure.

## D4 — antd v6 Select is an rc-virtual-list: pick VISIBLE options
`getByRole('option').first()` resolves to antd's HIDDEN "measure" template node (it renders the raw
option VALUE / uuid, since `optionLabelProp` uses a custom React label). The visible option rows are
separate DOM nodes. `getByRole('option').filter({ visible: true }).first()` (Playwright ≥ 1.51) skips
the phantom and clicks the first real, user-visible option — role-anchored, no CSS/class/nth-child.
Recorded as a page-object-layer lesson for any future antd-Select surface.

## D5 — corpus_test.go M204 pin reversed (UC1: TODO → implemented)
`TestRealCorpus_ManagerCoverageIsPresent` pinned UC1 as a declared TODO ("if it was implemented,
update this pin"). M243 implemented it, so the pin now asserts UC1 is a **non-TODO** journey tagged
`pt-assignment-assign`, played by `pt-manager` — guarding against a silent regression back to TODO or
a deletion of the one net-new v2.6 journey. `pt-assignment-assign` also joined `wantManagerPTs`.

## D6 — dev-run roster refresh (not a code change)
demo-1 was up on the demo SHOWCASE roster (maya-thriving, …), so a `pt-manager` cockpit login 400s
`unknown_identity`. Fixed by re-exporting the pt-world roster to the fake-FAPI mount + restarting the
fake services — exactly what `run-playthroughs.sh --reset` does (M211 iter-16). A run-environment step,
not a harness change. The canonical gate is `run-playthroughs.sh N --reset`.
