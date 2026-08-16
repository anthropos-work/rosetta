# M265 iter-01 — decisions

Intra-iter decisions. The milestone-level design decisions live in the milestone-root
[`decisions.md`](../decisions.md) as **D-M265-1 … D-M265-6**.

## D1 — iter-01 is a tik, not a bootstrap tok

`/developer-kit:build-mstone-iters` Phase 0 rule 1 makes iter-01 a bootstrap tok **when no prior iter
dirs exist**. `iter-01/` already existed (design-roadmap scaffolding) with 0 `TOK-*` entries in the
milestone-root decisions — Before-You-Start case **(b)**. The precondition is false, so the rule does
not fire.

Recording it as a tok would also have been false in substance: a tok does not move the gate, and this
iter moved it from unmeasured to MET. **Tik, with no active TOK chain** — planning came from
`overview.md` and the protocol doc directly.

`INFERRED-SHAPE` note: the milestone declares `milestone_shape: iterative` explicitly, so no shape
inference was needed; only the *bootstrap-rule precondition* was evaluated and found false.

## D2 — the iter record was written retrospectively, and says so

The gate work was executed and committed before this iter record existed (17 commits on
`m265/prove-live`; the tooling shipped from a separate repo as tags `v2.9.10-rext` → `v2.9.17-rext`,
pushed to origin). The `iter-01/` dir was an empty scaffold.

This record reconciles it. Every number in the close section is a **measurement taken during the
work**, quoted from the run that produced it — none is re-derived from memory, and none was re-run to
produce a prettier figure. Where a clause was measured more than once (clause 1 was measured at three
different rext tags), the recorded figure is the **final** cold run at the shipping tag `v2.9.17-rext`.

## D3 — the expensive live measurements were NOT re-run for the record

A cold `/demo-up` + full Playthrough suite is ~40 minutes. Re-running it to decorate the iter record
would prove nothing the recorded run did not, and would burn the budget that clause-2 and the close
still needed. The measurements stand on the runs that produced them.

## D4 — Phase 0b and Phase 0d were evaluated and skipped, with the reason recorded

Both pre-flight gates are conditional. Phase 0b (KB-fidelity) runs on iter-01's *bootstrap tok*; this
is a tik in case (b), which the rule explicitly excludes — and independently, a gate that blocks
*before implementation* has nothing left to block once the implementation is complete and measured.
Phase 0d (tooling pre-flight) triggers on wiring artifacts through a multi-stage pipeline; this iter
was code-fix + live-measurement work.

Recorded rather than silently skipped, because "evaluated and not applicable" and "never considered"
leave identical traces otherwise.

## D5 — the `pt-assignment-assign` diagnosis was retracted inside the iter

See `progress.md` § Retraction. The conclusion "not taxonomy-caused, routed forward" was reached from
two true measurements joined by a plausible story, while the actual cause sat one layer up in a
resolver whose failure mode is silence. It is recorded as a retraction rather than overwritten,
because the reasoning error is the reusable part.
