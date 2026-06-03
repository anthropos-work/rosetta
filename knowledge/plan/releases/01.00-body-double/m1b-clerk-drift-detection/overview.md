---
milestone: M1b
slug: clerk-drift-detection
version: v1.0 "body double"
milestone_shape: section
status: in-progress
started: 2026-06-03
---

# M1b — Clerk drift detection

## Goal
Make Clerk drift a **flagged, mechanical event** instead of a silent break: when the platform bumps
`clerk-sdk-go` / `@clerk/*`, re-`/align-dna` the new version (DNA diff = what changed) and re-`/align-run`
the existing Clerkenstein mirror (score drop = which genes broke), CI-gated on "alignment ≥ threshold".
This mechanizes the brief's *"follow platform updates within minutes"* requirement.

## Context (B-milestone — closes the gap after M1)
M1 shipped a Clerkenstein mirror aligned at `clerk@2.6.0`. It must *stay* aligned as Clerk moves. M1b
is **automation/config over M0** (it reuses `alignctl dna diff` + `alignctl run --gate` wholesale) —
no new measurement machinery. (Roadmap: `knowledge/plan/roadmap.md` § M1b.)

## Scope
### In
- A **drift-check** runnable: `alignctl dna diff` (old vs newly-authored DNA) + `alignctl run --gate`
  (re-score the mirror), with a combined report + a CI exit code (0 = no drift & gate met; 1 = DNA
  moved; 2 = gate regressed).
- A **CI gate**: a GitHub Actions workflow (+ a local gate script) that builds `alignctl` + the mirror
  runner and runs the alignment gate on push + a weekly schedule (the "follow updates" cadence).
- The **bump → DNA-diff → re-capture goldens → re-score → report** runbook, documented.

### Out
- Building the mirror (M1, done) · the JS surface + fake-API-server (M2) · stacks/seeding (v1.1).
- A live Clerk-version bump (none available now) — demonstrated with a **simulated** bumped DNA.

## Sections
See `progress.md`. S1 = the drift-check + CI gate (runnable, verified locally against a simulated
bump). S2 = the drift runbook documentation.

## Delivers → knowledge
- A **drift-detection runbook** section in `corpus/services/clerkenstein.md` (+ a pointer in
  `corpus/architecture/alignment_testing.md` § "How M1b consumes this").

## Where it lives
The drift-check + gate scripts + CI workflow live in the `clerkenstein` repo (they gate that mirror);
the runbook + this milestone's records live in rosetta.
