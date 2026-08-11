---
milestone: M257x
iter: 01
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
date: 2026-07-31
---

# iter-01 — bootstrap tok: find where the platform actually is, and write the procedure down

**Type:** tok (bootstrap). Authors the milestone's first strategy (`TOK-01`) and lands its
`iteration_protocol_ref`, which did not exist.

## Why a bootstrap tok had real work to do here

The milestone's five open questions were all *unknowns about the platform*, and a strategy authored without
answering them would have committed to the stale picture — which the overview explicitly warns against
("committing to a checklist now would be committing to the stale picture"). So this tok **measured first**:
five parallel probes against **platform origin HEAD**, never a pinned clone, plus independent re-verification
by the tok itself of every load-bearing claim.

## Plan

- Run the blocking Phase 0b KB-fidelity gate.
- Answer all five open questions, each cited to a sha or `file:line`.
- Re-verify every load-bearing claim independently before acting on it (the milestone's own warning: audits and
  reviewers have been wrong five times in three days).
- Author `corpus/ops/platform-alignment.md` — the `iteration_protocol_ref`, absent by design; its absence *was*
  the gap.
- Draft the clause-3 migration-status map and the clause-4 fence design.
- Author `TOK-01`.

## Escalation conditions

- Phase 0b RED → user-blocker (it returned **YELLOW**; proceeded).
- Evidence that two consecutive alignment attempts have been invalidated by mid-milestone platform commits →
  `re-scope-trigger`. **Not fired:** zero alignment attempts have been made yet, so the trigger's precondition
  is unmet. The target *is* moving (§ platform PR #20 open, `app` PR #1103 open), which is recorded as the
  standing risk TOK-01 is built around rather than a trigger firing.

## Acceptable close-no-lift outcomes

A bootstrap tok has no metric to move. Its deliverable is a strategy grounded in measurement plus the protocol
doc. Falsifying inherited premises counts as success, not as absence of progress.
