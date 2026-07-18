---
iter: 01
milestone: M224
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
date: 2026-07-16
---

# iter-01 — bootstrap tok (authors TOK-01)

**Type:** tok (bootstrap) — iter-01 of M224. Authors the FIRST strategy; does NOT terminate the call (continues
into iter-02 as a tik under TOK-01).

## Inputs
- `overview.md` (scope, exit_gate, In/Out, D-DESIGN-2 order-of-preference, R1/R3/R4 risks)
- `spec-notes.md` (the Phase 0b topic→doc→code triples + wiring facts the audit pinned)
- Protocol: `corpus/ops/demo/coverage-protocol.md` (measure→attribute→fix→re-measure) + `verification.md`
- The M222 read-model doc `corpus/services/hiring.md` (the score is the `local_jobsimulation_sessions` MIRROR)
- The M223 seeders (already write the mirror pair): `stack-seeding/seeders/{hiring_config,hiring_funnel}.go`

## Phase 0b — KB-fidelity gate: GREEN
Ran `/developer-kit:audit-kb-fidelity --milestone=M224`. All load-bearing topics PAIRED+ALIGNED. Two BLIND-AREAs
(clerkenstein FAPI org-publicMetadata; cockpit hiring vantage) are M224's own `Delivers →` deliverables. One
stale-misleading claim fixed inline in `hiring.md` (isHiring wiring target = FAPI `clerk-frontend/`, not BAPI).
Report: `../kb-fidelity-audit.md`.

## Initial strategy (→ TOK-01)
See milestone-root `decisions.md` → `## TOK-01`. In one line: **recruiter-render-first** — the gate is the
recruiter's scoreboard, so the opening tiks get the recruiter LOGGED IN and the comparison surface MEASURED then
GREEN before the candidate `/profile` polish; each tik follows the protocol's seed→render→**attribute**
(seed-gap vs render-gate vs Clerkenstein-identity, per hiring.md's read-path)→fix (data / Clerkenstein wiring /
sha-pinned demo-patch)→re-render loop against a LOCAL demo.

## Distance-to-gate context
- **Gate:** on a COLD reset-to-seed, ≥40 comparable non-junk candidate rows per EACH of 5 sims, non-degenerate
  score distribution, closure green, 0 prod-eject, over ≥3 consecutive cold runs. Latency REPORTED not gated.
- **Baseline:** UNMEASURED — the hiring-org comparison surface has never been rendered on a local stack (no
  recruiter seat, no isHiring wiring yet). Presumed 0 rows until the first render (iter-02).
- **Primary metric:** `min(rows-painted across the 5 sims)` on the recruiter's cold-seeded scoreboard (the gate is
  per-EACH → the minimum is the binding number), plus the distribution / closure / eject sub-gates.

## Next-tik direction (iter-02, first tik)
Stand up the **minimal enabling scaffold to REACH + MEASURE** the recruiter scoreboard, then take the **baseline
render** and **attribute** the gap — NO fix before the attribution. Concretely: recruiter (`vantage: manager`)
hero seat on the Meridian Talent 4th story with `jump_to` at the comparison surface; Clerkenstein FAPI
`isHiring` wiring (roster+resource, M39 precedent) **with the BLOCKING `/align-run`**; a **recruiter render-probe**
(count rows per sim + score distribution + closure/eject) under `stack-verify/e2e/`; tag rext; bring up a LOCAL
demo at the tag; log in as the recruiter; measure. (The 2 candidate `/profile` heroes are a later tik — not the
gate metric.)
