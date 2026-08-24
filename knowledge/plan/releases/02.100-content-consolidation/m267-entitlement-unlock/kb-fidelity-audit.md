---
title: "KB Fidelity Audit — M267 «The entitlement unlock»"
date: 2026-08-24
scope: milestone:M267
invoked-by: build-milestone
---

## Verdict

**RED on entry → GREEN after one inline fix.** One load-bearing stale claim, found and repaired before any
code landed. That is the gate working as designed rather than a formality passed.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Casbin PDP · the 6-matcher set · `m6` | `corpus/services/sentinel.md:88-97` | `app/internal/sentinel/casbin.go:20-45` | **PAIRED** |
| `'default'`-org policy semantics | `corpus/services/sentinel.md:201` | `casbin.go:21-23,44,45` | **PAIRED — was STALE** |
| Feature entitlement / `p6` (DEV path) | `corpus/ops/dev-identity.md:67,116,130` | `dev-stack/dev-identity.sh:205-208` | **PAIRED** |
| Seed contract · casbin grants (DEMO path) | `corpus/ops/seeding-spec.md:124-140` | `stack-seeding/seeders/identity.go:227-283` | **PAIRED — incomplete** |
| Demo authz-weakened posture | `corpus/ops/safety.md` §2.3, §3 | — (posture, not code) | **PAIRED** |

**No blind areas.** The feature-entitlement model *is* documented — `sentinel.md` enumerates all six matchers
and `dev-identity.md:67` names the `p6` row and its purpose. The milestone is not starting blind.

## Fidelity Findings

### 1. `'default'` org policies — the escape is not universal · **STALE · BLOCKER (fixed)**

- **Source:** `corpus/services/sentinel.md:201`
- **Expected (doc):** *"'default' org policies apply to all organizations unless overridden by org-specific
  entries."* Stated unqualified, as a property of the policy model.
- **Actual (code):** true for **three of six** matchers, false for the one this milestone depends on.
  Measured at `app/internal/sentinel/casbin.go`:
  - `:21` `m` — templated tier default (`p.tier == '%s'`)
  - `:22` `m2` — `('default' == p2.org || r2.org == p2.org)`
  - `:23` `m3` — `('default' == p3.org || r3.org == p3.org)`
  - `:44` `m5` — `(r5.org == p5.org || p5.org == 'default')`
  - **`:45` `m6` — `( g3(p6.org, r6.sub) ) && r6.feat == p6.feat && r6.count <= parseFloat(p6.max)`.
    ZERO occurrences of `default`.**
- **Verdict:** **STALE**, and stale in the direction that costs the most.
- **Why blocker:** M267's whole implementation is "write the missing `p6` row". A developer reading `:201`
  as an unqualified rule would write **one** row naming the org `default` and expect every org covered. That
  row authorizes **nobody** — and it fails silently, because the demo's raw-error path is PostHog-gated
  (`AISimulationStartWithoutSession.tsx:214-217`) and a demo has no PostHog, so the browser shows the generic
  message and hides the cause. The milestone would have "landed", the gate would still be shut, and the
  diagnosis would have started from a doc that said the fix was correct.
- **Fix owner:** **doc** — the code is right and is the contract. Applied inline (see below).

## Completeness Gaps

### 2. `seeding-spec.md` documents two grant families where the app needs three · **critical — already a milestone deliverable**

`corpus/ops/seeding-spec.md:124-140` § *"The minimum-proof identity + the casbin gotcha"* enumerates
*"Two casbin subtleties the live proof caught (both load-bearing)"* and documents the `g2 (org, user, role)`
grant — including the arg-order trap and the `casbin_rules`/`casbin_rule` table-name gotcha. It says nothing
about `p6`. Read as *the* seed contract, it reads complete.

**Not filed as a separate blocker: M267's `Delivers →` line already promises exactly this doc**, so the gap is
promoted rather than outstanding. It is recorded here so the Phase-5 documentation pass has a concrete target
and cannot mistake the omission for a deliberate scoping choice.

### 3. `sentinel.md`'s matcher table did not carry `m6`'s distinguishing property · **incidental (fixed)**

`:93` described `m6` as *"(no tier logic)"* — accurate but incomplete: the property that actually matters when
you go to write a `p6` row is the absent `'default'` escape. Added, because the matcher table is where someone
looks before writing a policy row.

## Applied Fixes

1. **`corpus/services/sentinel.md:201`** — the `'default'`-org sentence is now qualified per matcher, with
   `casbin.go` line citations for all five escapes and the measured absence on `m6`, plus a note that reading it
   unqualified sends you to a one-row fix that grants nothing.
2. **`corpus/services/sentinel.md:93`** — the `m6` row of the matcher table now records the no-`'default'`-escape
   property alongside "no tier logic".

Both are doc-side. **No code changed** — the code was correct in both cases.

## Open Items (require user decision)

**None.** Both findings were unambiguous (the code is the contract and the doc drifted), so both were applied
inline per Phase 6 rather than escalated.

Five open questions remain in the milestone's own `overview.md` — whether other features route through
`CanPerformFeatureAction`, whether `pt-world` needs a second insert site, whether DEV and DEMO converge on one
grant-writing shape, whether any v2.10 seeder writes `public.organization_features`, and how to read the real
error past the PostHog gate. **Those are implementation questions for this milestone to answer, not KB-fidelity
findings**, and they are deliberately left where they are.

## Gate Result

**GREEN — proceed to Phase 1.**

Entry verdict was RED on finding 1. It is repaired, and the repair is the audit's own justification: the stale
sentence would have been read as truth by the very next thing that happened in this milestone.
