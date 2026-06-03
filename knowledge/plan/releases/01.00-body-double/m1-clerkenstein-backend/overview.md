---
milestone: M1
slug: clerkenstein-backend
version: v1.0 "body double"
milestone_shape: iterative
status: in-progress
started: 2026-06-03
exit_gate: "alignctl run reports 100% alignment on the platform-consumed Clerk Go surface (critical capabilities) AND >=95% overall, with any waived genes documented + justified"
iteration_protocol_ref: corpus/architecture/alignment_testing.md
---

# M1 — Clerkenstein backend mirror (Go)

## Goal
The first real **mirror**, built *by* the M0 alignment process: a drop-in Go stand-in for
`colony/authn`'s provider + the Clerk `orgclient`, injected into the demo build via `go.mod replace`
(zero platform-repo edits), so backend services authenticate with one universal credential and
locally-minted JWTs — no real Clerk. (Roadmap: `knowledge/plan/roadmap.md` § M1.)

## Exit gate (the commitment; the path is open → iterative)
`alignctl run` reports **100% alignment on the platform-consumed Clerk Go surface (critical
capabilities) and ≥95% overall**, with any waived genes documented + justified in the divergence
report. The score is measured by the M0 framework against the **Clerk Alignment DNA** this milestone
authors.

## Iteration protocol
`corpus/architecture/alignment_testing.md` — the measure → fix-diverging-genes → re-measure loop.
Each tik closes a batch of diverging genes (or authors/extends the DNA + mirror) and re-scores via
`/align-run`; the gate is the alignment score.

## Why iterative (not section)
The deliverables are writable, but *which genes diverge and how costly each is to close* only emerges
from measuring the mirror against the real Clerk surface — a fixed up-front checklist would be
speculative. The score is the commitment; the path to it is open.

## Scope
### In
- The dedicated **`clerkenstein` repo** (cloned into the gitignored `anthropos-demo/`, mirroring the
  `anthropos-dev/` pattern), with a gitignored `.agentspace/` holding the pinned official Clerk SDK
  (`clerk-sdk-go/v2 @ v2.6.0`) for interface extraction.
- The **Clerk Alignment DNA** (`clerk@2.6.0` genome) authored via `/align-dna` from the consumed
  surface: the `colony/authn` provider claim shape (`eid`, `email`, `firstname`, `lastname`,
  `org.eid`, `org_id`, `org_role`) + the `orgclient` methods the platform calls
  (Invite/Create/Delete membership, Create org, BulkInvite, RevokeInvitation, ChangeRole,
  Update{User,Membership}Metadata, UpdateClerkOrganizationWithExternalId).
- The Go mirror: an authn-provider twin (local JWT mint/verify; one universal credential) + an
  `orgclient` twin; its `--target source|mirror` **runner** for `alignctl`.
- Source **goldens** + the `go.mod replace` + skip-worktree injection recipe.
- Driving `/align-run`'s score to the exit gate.

### Out
- The JS/browser surface + webhook injector → **M2**.
- Drift CI wiring across Clerk version bumps → **M1b**.
- Multi-instance demo stacks, data seeding → v1.1.

## Re-scope trigger
If 5 consecutive triggered toks fail to produce a viable new strategy, OR a critical capability proves
genuinely un-mirrorable offline (no viable golden-capture path), escalate to user-strategic-replan.

## KB dependencies (contract)
`corpus/architecture/alignment_testing.md` (the iteration protocol + the framework M0 shipped),
`corpus/services/clerk-integration.md`, `corpus/architecture/shared_libraries.md` (§ authn/colony),
`corpus/ops/staging-clerk.md`, `corpus/ops/webhook_setup.md`.

## Delivers → knowledge
- `corpus/services/clerkenstein.md` (the mirror design + injection mechanism — net-new) + the Clerk
  Alignment DNA authored in the `clerkenstein` repo.

## Note on repo footprint
The mirror's **code, DNA, goldens, alignment tests, and runner live in the `clerkenstein` repo**
(gitignored `anthropos-demo/`), not in rosetta. This milestone's rosetta footprint is the planning /
iter docs (here) + the eventual `corpus/services/clerkenstein.md`. The build measures progress via the
alignment score, recorded per-iter in this milestone's `progress.md`.
