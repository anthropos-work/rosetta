# M256 · iter-05 — decisions

## D18 — The create-role flow has a HIDDEN mandatory step, so iter-04's "Save no-ops" was incomplete

Clicking **"Suggest skills"** *transforms* the create-role dialog: two scope options appear
("Core + complementary skills" / "Core skills only") and the primary button changes from **`Save` to
`Generate`**. So **"Core skills" is a mandatory step**, and a role cannot be created without completing it.
"Start from scratch" does **not** transform the dialog and does not unblock `Save`.

**Two consequences, in opposite directions:**

1. **A real product defect, now evidenced.** `Save` is **enabled** with the form incomplete; clicking it
   produces no HTTP ≥ 400, no console error, no navigation, no row — and the `alert` region is **empty**. A
   user clicks Save and *nothing whatsoever happens, with no explanation*. That is a finding to report, not a
   test bug, and it is exactly the kind of thing a journey-level test surfaces that a render check cannot.
2. **The journey has a live-LLM leg.** `Generate` prefills skills via the model, so under the protocol's
   integration-dependent assertion-boundary rule this journey belongs **with the studio lane** — asserted at
   its completion boundary and **excluded from the timed median** — rather than treated as a normal write.
   That is a design consequence the owning tik must honour, not a detail.

`org-admin.roles.UC1` stays `playthrough: TODO`, with the finding sharpened from *"Save appears to no-op"* to
*"Save is enabled while a mandatory step is incomplete, and fails silently"*.

## D19 — The assign-tags modal is unreachable through four measured routes — and `force: true` LIES

Four attempts, each measured, each failing. Recorded on `MembersBulkActionsPage.openAssignTags()` so the
next attempt starts from the evidence:

| # | Attempt | Result |
|---|---|---|
| 1 | click the tag label normally | the bulk-action **dropdown stays mounted over the modal it opened** and intercepts its pointer events. **A hit-target interception is invisible to actionability checks** — Playwright reported "visible, enabled and stable" and retried **454 times** to the 240 s budget |
| 2 | `Escape` to dismiss the dropdown | **closes the MODAL, not the dropdown** (antd Modal defaults to `keyboard: true`) — this was iter-04's own "fix", and it made things worse |
| 3 | an outside-click landing *inside* the modal | dropdown still open — **9** `menuitem`s counted afterwards |
| 4 | `checkbox.check({ force: true })` | **`isChecked()` → true, and the modal's submit stays DISABLED.** The DOM flipped; antd's React state did not; nothing was assigned (the tag's member tally stayed **0**) |

**Attempt 4 is the finding that generalises, and it is the sharpest thing this iter produced.**
**`force: true` can manufacture a control that LOOKS checked to a test and is unknown to the application.**
A weaker Playthrough would have asserted the visual state, gone green, and proved nothing — the same
false-green class as iter-02 D6's studio test, arrived at from a completely different direction. The reason
this one was caught is that the assertion was a **cross-surface read-back** (the tag's member tally) rather
than the state of the thing just clicked.

**Disposition:** `org-admin.members.UC1` stays `TODO`. `unimplementable-without-platform-edit` is now a **live
possibility** — an overlay that traps pointer events over its own modal is a platform defect, and P3's escape
valve exists for exactly that — but it is not claimed yet, because two candidates remain untried (driving the
`<label for=…>` element; reaching the surface by a route that never opens a dropdown). The claim will be made
only when it is earned.

## D20 — Clause 2's fifth mutating journey must come from somewhere else

Mutating count stands at **3** (`pt-assignment-assign`, `pt-orgadmin-tag-create`,
`pt-orgadmin-setting-toggle`); clause 2 needs **≥ 5**. Both org-admin candidates are now blocked on
platform-side behaviour rather than on test effort (D18, D19), so **the next iter must not assume they will
land**. It should identify one or two further write-and-read-back surfaces in parallel with one last attempt
at these — candidates visible from work already done in this milestone:

- the **`Remove Tags`** bulk action (the same menu, but it operates on a member who already carries a tag, so
  it may avoid the modal entirely);
- the **profile self-evaluation / skills claim** write (`profile-skills.self-evaluation`, an M206 reservation
  that needs a verdict anyway under clause 3);
- an **onboarding completion** write, if the seed work clause 3 needs for onboarding lands first.

Recorded so the fifth journey is a *chosen* target with a rationale rather than whatever happens to work.
