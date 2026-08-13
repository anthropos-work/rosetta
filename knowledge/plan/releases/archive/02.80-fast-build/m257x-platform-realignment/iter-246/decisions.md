# iter-246 — decisions

## `D-M257x-246-1` — the falsification's arithmetic fires; its conclusion is rejected, with evidence

`P-246-1` sealed: *"≤ 1 defect means six-for-six was coincidence, not structure, and the route closes as a
mis-diagnosis."* The census returned **0**. The arithmetic branch therefore fires — and the conclusion it
names is **wrong**, because it assumed the only two explanations were *structure* and *coincidence*. There
is a third, and it is the one the evidence supports: **the six repairs held, and nothing was left behind
to catch the seventh.**

The seal's purpose is to stop a number being re-argued after it lands. It is not a licence to publish an
inference the same iter can show is false. **State which half fired, and why the other is rejected.**

## `D-M257x-246-2` — workspace substitution, printed, and gated on the named workspace being real

A `stack-<X>/<repo>` path whose workspace this host has not provisioned is graded under any workspace that
carries the repo. Without it, **39 correct lines across 8 documents** were RED on the build host.

Two guards on the pardon, both from this iter's own tests:
* the **named** workspace must be a real directory here — otherwise a fabricated `stack-nowhere/platform`
  was silently substituted to a real one and reported green;
* the substitution is **printed on every run** (`stack-dev x32`), because a substitution nobody can see is
  a silent exclusion with a result attached.

## `D-M257x-246-3` — a repo-level miss is REFUSED; a sub-path miss is a FINDING

`cd stack-dev/chronos` (decommissioned) and `cd stack-dev/experiments` (a separate org repo cloned by
hand) are indistinguishable from typos with the evidence available. iter-244's rule — *a WRONG path and an
UNCHECKABLE one must not share a bucket* — cuts the **other** way here: the refusal bucket is honest
precisely because it does not claim to know.

Once the repo is located, its interior IS knowable, so a missing subdirectory beneath it fires.

## `D-M257x-246-4` — corpus-wide scope, not `CLAUDE.md`-only

The route names `CLAUDE.md` because that is where the defects were *noticed*, not where the cause lives.
The cause is that a fenced command body is in no guard's subject anywhere. Scoping the fence to one
document would have reproduced the original defect one level up: a repair to the instance, not the class.
Measured cost of the wider scope: **621 blocks**, of which 188 commands are gradeable and the rest refuse
by a named reason.
