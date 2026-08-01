---
milestone: M257x
iter: 30
---

# iter-30 decisions

## D1 — the scoped-run artifact is now both LABELLED and NON-DESTRUCTIVE

`FIX-M257x-iter27-scoped-run-clobbers-binding-report`. Two mechanisms, because either alone is
insufficient: a provenance sidecar naming the invocation, AND a binding-only snapshot
(`last-binding-run.json` / `last-binding-report.json`) that a scoped run may not write. Labelling alone
would have told you the file was a probe *after* the binding measurement was already gone.

Chose a shell function taking explicit arguments over reading the globals, purely so
`TestRunnerSafety_RunProvenance` can execute it against a temp dir and assert what lands on disk. A
grep-shaped guard here would have passed on the label-only mutant (M3).

Not done: closing `CHECK-M257x-iter30-scoped-classifier-misses-filenames`. A bare spec filename narrows
the suite but classifies as unscoped, so it would now write a binding snapshot from a one-spec run. That
classification is an existing, deliberately tested contract (`notScoping` includes
`tests/profile-identity.spec.ts`), and changing it is a separate decision with its own blast radius.

## D2 — the individually-sufficient-clauses mutation rule

The funnel fix has two clauses (a `passed` discriminator, and `.first()` → `.last()`). Single-clause
mutants BOTH survived; only the full revert went RED. Reported rather than explained away, then
re-derived: each clause alone still lands on a node carrying the asserted text, so a single-clause mutant
is not a broken fix but a *different working one*.

Kept both clauses, and said so in the code comment — `.last()` alone addresses a bare name+role fragment
rather than a card, and the discriminator alone leaves the choice to DOM order, which is the original
defect. Redundant for today's assertion; jointly load-bearing for the accessor's meaning.

Generalised into `corpus/ops/platform-alignment.md` §8 rule 5 (same commit, per the protocol-evolution
rule).

## D3 — succession routed, not fixed, under the scope-creep tripwire

The cause is measured (28 roles tied at `risk 68`, 25 rendered, 39 distinct roles across 40 members, the
hero's role held by exactly 1 person). The fix is a **seed-shape** change — real incumbency for the proof
role — which is a third line of work in an iter with two planned lines, and it perturbs org-scale
aggregates that other specs read. Routed as `FIX-M257x-iter30-succession-role-tiebreak` with a named
target, per the three-fate rule's Fate 3.

Explicitly NOT chosen: weakening the assertion to match what renders. The assertion's intent — *this
manager's own tenant's projection* — is correct; it is the seeded world that cannot support it stably.
