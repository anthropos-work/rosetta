# iter-89 — decisions

## `D-M257x-89-1` — four failures, one cause; the joining experiment was `git status`

iter-88 routed the demo-stack residual as up to three classes across two handlers, and every one of its
descriptions was accurate. They are one defect. The probe that collapsed them cost under a minute:
`git status` on `stack-demo/next-web-app` shows a demo-patch **left applied**, and a patched file cannot
match a pristine `pre_sha256`.

§5 rule 28 — *three true facts do not make a cause; join them with one experiment* — earned again, and
the generalisable form is worth carrying: **when several checks fail against the same external artifact,
read the artifact's state before classifying the failures.**

## `D-M257x-89-2` — do not re-pin, do not clean; escalate

Two things were available and both were refused.

**Re-pinning the baseline** would have made one file match again and **hidden the asymmetry** — the
mechanism would keep leaving clones dirty and the evidence would be gone. It is also precisely what
iter-88's own routing instruction forbade (*"adjudicate before touching; do not re-pin first"*), and what
`demopatch-spec.md` warns about in its own voice.

**Cleaning the clones** (`--force-pristine`, which runs `git checkout -- <path>`) is a decision about
uncommitted state. The user and the orchestrator are the only allowed deciders on that, and the fact that
the tool ships the escape hatch does not transfer the decision to me.

**The clones were left exactly as found**, and that is the deliverable's integrity: the next session can
reproduce the defect without re-creating it.

## The decision requested — four options, with the trade named

The repair is a design choice on a mechanism that **rewrites platform source inside a build**, so the
blast radius is real and it is not mine to pick:

| | option | what it costs |
|---|---|---|
| **(a)** | make revert **symmetric** — reverse the anchor transformation instead of comparing whole files | most faithful to *"the anchor is the contract"*; needs a reliable inverse for every patch shape, which not all edits have |
| **(b)** | **journal the observed pre-state at apply time**, revert restores exactly that | strongest and simplest to reason about — revert stops depending on baselines entirely. Adds per-apply state, which must itself be cleaned up |
| **(c)** | make apply **strict** — refuse on base drift | restores G2/G5 consistency by giving up self-healing, and base drift is the NORMAL state for these clones, so this would refuse constantly |
| **(d)** | accept it; have the runner call the existing `--force-pristine` | cheapest; makes *"the clone is left git-clean"* true by a `git checkout` rather than by the patch reverting itself, and quietly weakens G5 into a cleanup step |

My reading, offered but not acted on: **(b)**, because it makes revert independent of a baseline that is
guaranteed to go stale, and because it is the only option that keeps both G2 and G5 true at once. **(d)**
is a legitimate pragmatic answer if the demo path is the only consumer — but it should be chosen out
loud and written into the spec, not fallen into.

## The finding underneath, which outlives the fix

**Two guards can each be individually correct and jointly inconsistent.** G2 (*refuse on drift; never
guess*) and G5 (*always self-revert; leave the clone clean*) cannot both hold once the base is allowed to
move — and the base is always allowed to move. Every test asserted them **separately**, and each passed.
Nothing asserted their **conjunction**, which is where the defect lives.

This is §5 rule 17's shape (*verify the predicate, not the count*) lifted from a claim to a **contract**:
a specification with seven guards needs at least one test per *pair that can interact*, not one per
guard. Recorded here rather than in the protocol doc because the fix is not yet chosen; it belongs in
`demopatch-spec.md` alongside whichever option lands.
