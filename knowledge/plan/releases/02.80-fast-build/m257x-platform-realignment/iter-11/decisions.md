---
milestone: M257x
iter: 11
---

# iter-11 — decisions

## D-M257x-12: the derivation WINS over an inherited `$STACK_DIR`, it does not merely default to it

`STACK_DIR="${STACK_DIR:-$derived}"` was the smaller change and it was rejected. It fixes only the caller
that passes *nothing* (`dev-stack:298`) and leaves untouched the caller that passes something *wrong* — the
one that actually cost this milestone an iter. Under that shape the failure stays expressible and the only
control is a message the reader must interpret correctly, which is precisely what did not happen.

So the derivation is authoritative whenever the project has a known layout, and an inherited value that
disagrees is **named** (a counted warning) before being discarded. An explicit value survives only where
there is nothing to derive.

The cost is real and accepted: an operator who deliberately wants the receipts read from somewhere else can
no longer ask for that on a `demo-N`/`dev-N` project. Nothing in rext wants it, and §8 rule 4 — *prefer a
design that cannot express the drift* — is the whole reason the parameter is being taken away.

## D-M257x-13: the receipt asserts are gated on the project's TYPE, not on whether a caller set a variable

Before this iter, "does this stack have demo-patch and frontend-build receipts?" was decided by
`[ -n "$STACK_DIR" ]` — an accident of the caller. That is why `dev-stack:298` skipped the block: not
because a dev stack has no patches (true, and the right reason), but because it forgot a variable (wrong
reason, right answer, by luck).

Deriving the dir removes the luck, and would have turned the right answer into two false warnings per dev
bring-up. `target_is_demo_project` states the actual proposition: `demopatch.log` and `buildfail.log` are
written **only** by `up-injected.sh`; a dev stack has no patch phase and no UI tier, so their absence there
is the correct state and asserting on it would fabricate a defect.

The transcript (`autoverify.log`) and the verdict (`autoverify.json`) are **not** demo-only and stay gated
on the dir alone — which is how dev acquires both for the first time.

## D-M257x-14: a missing derived stack dir is a NAMED SKIP, not a warning

The first cut warned. Ten targeted tests passed; the full suite went red in **18** pre-existing fixtures,
because every autoverify test uses a synthetic `--project demo-1` with no stack dir on disk. §8 rule 6: a
fence that cries wolf gets disabled, and a disabled fence is indistinguishable from never having written one.

The severity was reasoned rather than tuned:

- `up-injected.sh` `mkdir -p "$STACK"` at `:226` and truncates both logs by `:237`. **On any bring-up that
  got past its first screenful the dir exists.** A missing dir does not describe a real bring-up.
- What it does describe is running the script from a clone that did not bring the stack up — an operator
  diagnostic, and the exact confusion of §5 rule 12. Naming the **path consulted** is the entire fix for it.
- Every state that describes a real bring-up — dir present with receipts absent, or populated — remains a
  full warning. Nothing gate clause 1 depends on was downgraded.

The alternative considered and rejected: keep the warning and scaffold a stack dir in all 18 fixtures. That
buys a warning nobody can act on, on every synthetic run, and imposes the scaffold on every future
autoverify test — for a state that cannot occur on the path the gate measures.
