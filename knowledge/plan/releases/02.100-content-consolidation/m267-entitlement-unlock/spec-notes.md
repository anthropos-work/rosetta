# M267 — Spec notes

## Pre-flight audits — section 1 (`Seed one p6 row per seeded org`)

**Phase 0b KB-fidelity: RED on entry → GREEN after one inline fix.** Run 2026-08-24 at rosetta `def78548`.
Report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md).

⚠️ **AUDIT REUSE ON RESUME.** Per `/developer-kit:build-milestone` §"Audit reuse across sections", a resumed
session may **skip Phase 0b and go straight to Phase 1** if all three conditions still hold:

1. this block exists (it does);
2. `git diff def78548...HEAD` shows no change to the load-bearing docs — `corpus/services/sentinel.md`,
   `corpus/ops/seeding-spec.md`, `corpus/ops/dev-identity.md`, `corpus/ops/safety.md`;
3. the section's touched files stay inside `stack-seeding/` + those docs.

If any has moved, re-run it. The audit is cheap to repeat and expensive to skip wrongly.

**What it found, because it changes how section 1 is implemented:** `corpus/services/sentinel.md:201` claimed —
unqualified — that *"'default' org policies apply to all organizations unless overridden by org-specific
entries."* Measured in `app/internal/sentinel/casbin.go`, that holds for `m2` (`:22`), `m3` (`:23`), `m5` (`:44`)
and the templated tier default on `m` (`:21`), and **`m6` (`:45`) has ZERO occurrences of `default`.** So the
`p6` row **must name the real org id** and there is no one-row way to cover every org. Both `sentinel.md:201`
and the `m6` row of the matcher table at `:93` were corrected; **no code changed — the code was the contract.**

**Why that mattered enough to block:** the obvious reading sends you to a single `p6 default …` row, which
authorizes nobody — and it fails *silently*, because the demo's raw-error path is PostHog-gated
(`AISimulationStartWithoutSession.tsx:214-217`) and a demo has no PostHog. The milestone would have reported
landed with the gate still shut.

---

_Sections below are derived from the scope in [`overview.md`](overview.md). Nothing implemented yet — Phase 1
had not started when the milestone was paused (2026-08-24)._

## The `p6` insert — shape, call site, and per-org loop

_None yet._

## Convergence with the DEV path (`dev-identity.sh:205-208`)

_None yet._

## The `Used = 0` safety argument, re-checked against every v2.10 seeder

_None yet._

## Policy reload — observing the existing `sentinel:policy:invalidate` publish

_None yet._

## Live proof — vantages, stack, and how the browser check is run

_None yet._

## The inert limiters (recorded so they are not chased)

_None yet._

## The PostHog diagnostic trap — how a residual failure gets diagnosed

_None yet._

## `seeding-spec.md` — what the delivered doc change says

_None yet._
