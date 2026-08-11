---
iter: 54
milestone: M257x
iteration_type: tok
tok_flavor: triggered
status: closed-fixed
opened: 2026-08-03
platform_ref_at_open: ef32d4cd8e0ceecf528a74c37d5e2ae5804ce021
---

# iter-54 — TOK-04: pin the target, or stop calling it a measurement

**Type:** tok (triggered). Session-terminating by protocol.

## Why this iter is a tok

Two independent triggers fired on iter-53, and the user has ruled on what happens next.

1. **The `re_scope_trigger` fired, occurrence 2 of 2.** The milestone's own trigger reads: *"If TWO
   consecutive full-alignment attempts are invalidated by new platform commits landing mid-milestone — i.e.
   the target moves faster than we can track it — STOP and escalate. The answer then is a pinning-and-tracking
   POLICY (how we choose a platform ref, how we notice it moved, who re-points), not more alignment work."*
   It fired exactly as written: platform moved `2adcf714 → ef32d4cd` **during iter-53**, invalidating a seat's
   clearance by name.

2. **The audit instrument was never frozen.** It lived at a git-ignored scratch path and was re-authored from
   a summary on every pass. The published series is therefore not a comparable series.

The user's instruction for this iteration is explicit and has three parts, in order: **re-survey the platform
(before → now), reassess the milestone honestly, then author TOK-04.** This iter does those three things.

## Step 0 — re-survey before authoring (mandatory, §Phase 1)

Run and recorded before any strategy was written:

- Platform clone brought from `2adcf71` to origin HEAD `ef32d4c` (fetch + fast-forward; **zero platform-repo
  edits**, nothing committed into it).
- The clause-3 fence re-run against the new HEAD **before** any corpus edit: **RED, 3 findings, direction B.**
- The 3-no-prog premise re-checked: not applicable — iter-53 closed `closed-fixed`, and the trigger here is
  the re-scope trigger plus a user ruling, not the 3-tik streak.

The trigger is **not** stale. It is the opposite of stale: the platform moved again between iter-53's close
and this iter's open being written down.

## Planned scope

1. **The before → now service inventory**, cited to platform source, landed by updating
   `corpus/architecture/platform-migration-status.md` against origin HEAD rather than starting over.
2. **An honest reassessment** of all five gate clauses, including which met clauses are now stale.
3. **TOK-04** in the milestone-root `decisions.md`: a pinning-and-tracking policy, plus the next-tik
   direction and the harden placement.

## Escalation conditions

A tok terminates the session by protocol. Exit `tok-fired` after the close; do not continue into tiks — the
revised strategy is surfaced to the user before anything commits to it.
