---
iter: 01
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
date: 2026-06-03
---

# iter-01 — bootstrap tok

The milestone's first iter: author the initial strategy (TOK-01) for building Clerkenstein toward the
alignment-score gate. No prior iters; no stalled strategy to revise — this authors the first one.

**Inputs:** M1 `overview.md` (scope, exit gate), `spec-notes.md` (the consumed Clerk surface + score
mechanics), the protocol doc `corpus/architecture/alignment_testing.md`, and the M0 framework.

**Output:** `TOK-01` in the milestone-root `decisions.md` — *mirror-by-score, easy-side-first*: author
the Clerk DNA, build the offline-capturable **authn** twin to 100% critical first, then the live-SaaS
**orgclient** twin; inject via `go.mod replace`. Plus the milestone's central **open decision (D1)**:
the orgclient golden source (live Clerk vs hand-authored vs hybrid) + workspace-setup availability.

See `progress.md` for the close.
