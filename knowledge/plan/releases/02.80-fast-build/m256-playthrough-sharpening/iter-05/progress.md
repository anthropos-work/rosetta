# M256 · iter-05 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 3 (finish the org-admin cluster).
**Phase 0d:** re-run on the current tree — `9 product(s), 22 use case(s), 20 live, 2 TODO`, valid.

## Phase A — both named hypotheses TESTED and REFUTED

The iter existed to unblock two gate-critical journeys using the diagnoses iter-04 wrote. Both
hypotheses were wrong, and being wrong produced better evidence than being right would have.

**Roles (D18).** Hypothesis: *the mandatory "Core skills" step fails validation silently; the "Suggest
skills" path unblocks Save.* **Half right.** "Suggest skills" **transforms the dialog** — two scope options
appear and the primary button changes from **`Save` to `Generate`** — confirming "Core skills" is mandatory.
But that also means the journey carries a **live-LLM leg**, so it belongs with the studio lane under the
integration-dependent assertion-boundary rule, not in the timed median. And the enabled-`Save`-with-an-empty-
`alert` is confirmed as a **real product defect**, not a test bug.

**Members (D19).** Hypothesis: *the modal's filter box makes the target checkbox actionable.* **Refuted, four
ways** — including a refutation of **this milestone's own iter-04 fix**: `Escape` closes the **modal**, not
the dropdown. The decisive finding is that **`check({ force: true })` flips the DOM but not antd's React
state**: `isChecked()` reported `true` while the submit stayed **disabled** and the tag's member tally stayed
**0**. *`force: true` can manufacture a control that looks checked to a test and is unknown to the
application.*

## Phase B — what landed

No coverage. One **correction** and a body of recorded evidence:

- **`openAssignTags()`'s `Escape` removed** — iter-04's own fix, now known to close the dialog the next step
  waits for. The reason is recorded at the call site so it is not re-added.
- **The four measured attempts documented on the page object**, next to the code, so the next attempt starts
  from the evidence instead of re-deriving it.
- **`e2e/drafts/README.md` sharpened** with both diagnoses.
- **The draft spec left calling a helper that no longer exists**, deliberately, so it cannot be moved back
  into `tests/` without reading the note first.

rext commit `38d21e9`, tag **`fast-build-m256-orgadmin-diagnoses`** — pushed to origin.

## Phase C — the suite is unchanged and green

`136 passed`; `ptreport` **20/22 passing, 2 `[TODO]`, 0 failing, 0 unimplementable**. No re-measure: the
denominator did not change, so iter-04's **0.5434×** stands.

## Close — 2026-07-28

**Outcome:** both hypotheses refuted with evidence; one of this milestone's own fixes reverted as wrong; two
findings that generalise well beyond org-admin (a platform overlay that traps pointer events over its own
modal, and `force: true` manufacturing a false UI state). No coverage lift, by an honest margin.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET — clause 1 met (0.5434×, unchanged); **clause 2** stands at **3 of 5** mutating, 0 negative controls, 0 `blocked`; **clause 3** at 2/4 org-admin + 0/5 onboarding + no verdicts; **D-v28-5** unstarted.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this is **1** no-progress tik, not 3; iters 02–04 all progressed) — (3) re-scope: n (a `closed-no-lift` with documented falsification does **not** count toward the trigger — the protocol working, not failing; and 0 surfaces are yet claimed unimplementable) — (4) user-blocker: n (no red at batch end; both journeys remain declared `TODO`, so nothing accumulates and D-v28-3 has nothing to escalate) — (5) **cap-reached: y — this is the 5th tik of the invocation** — (6) protocol-stop: n — Outcome: exit-5
**Decisions:** D18 (the roles flow's hidden mandatory `Generate` step + the confirmed silent-Save product defect), D19 (the assign-tags modal unreachable through four measured routes, and `force: true` lies), D20 (clause 2's fifth mutating journey must be chosen from elsewhere).
**Side-deliverables:** none.
**Routes carried forward:**
- `PT-M256-orgadmin-role-create` → **next iter.** Drive the `Generate` path and assert at its **completion
  boundary**; budget it with the studio lane, not the median. **And report the silent-Save defect.**
- `PT-M256-orgadmin-member-tag` → **next iter.** Two untried candidates (the `<label for=…>` element; a route
  that never opens a dropdown). If both fail, declare `unimplementable-without-platform-edit` **with the
  four-attempt evidence** — earned, not assumed.
- `PT-M256-clause2-fifth-write` → **next iter (D20).** Choose the 5th/6th mutating journey deliberately:
  the `Remove Tags` bulk action, the profile self-evaluation write, or an onboarding completion.
- All pre-existing routes in `../progress.md` § Next-iter routing still stand.
**Lessons:**
1. **A refuted hypothesis that produces a mechanism is worth more than a passing test.** Neither journey
   landed, and the iter still produced the milestone's sharpest finding — that `force: true` can create a
   control the application does not know about. Both of this milestone's false-greens (iter-02 D6, this one)
   were caught by an assertion that read state **through a different surface** than the one being driven.
2. **Check your own fixes against the mechanism, not the symptom.** iter-04's `Escape` was plausible,
   untested against the actual overlay, and made the failure worse. One line of verification —
   *"is the dropdown still open?"* — would have caught it immediately.
3. **When a control resists a real click, do not reach for `force`.** `force` skips the very check that was
   telling you the truth. The interception is information.
