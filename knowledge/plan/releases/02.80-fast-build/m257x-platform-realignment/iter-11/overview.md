---
milestone: M257x
iter: 11
iteration_type: tik
status: in-progress
opened: 2026-08-01
---

# iter-11 — `FIX-M257x-autoverify-evidence-log-path`

**Active strategy reference:** `TOK-01: instrument first, then follow` — step 3/5 (*"land the fences, each
watched going RED, before trusting any green"* + *"prove it cold, three times"*). This iter clears the last
standing autoverify warning class so clause 1's three cold cycles are measurable.

## Step 0 — Re-survey (mandatory) — the inherited pre-compute is REFUTED on 3 of its 5 points

Measured before writing a line, against the live `demo-1` stack and the rext authoring copy.

| pre-compute claim | measurement |
|---|---|
| *"The path is wrong — autoverify reads `$STACK_DIR/…`; the bring-up writes elsewhere"* | **REFUTED.** `up-injected.sh:2550` calls autoverify with `STACK_DIR="$STACK"`, and `:75` sets `STACK="$HERE/stacks/demo-$N"`. The bring-up path is correct **by construction**. |
| *"iter-10 measured 2 FAILED on demo-1, both this"* | **Measured through a DIFFERENT invocation.** `stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/autoverify.json` = `warnings:1 … 15:49:36Z` (the bring-up's own verdict). `stack-demo/autoverify.json` = `warnings:2 … 20:37:55Z` — a standalone re-run with `STACK_DIR=stack-demo`, the workspace root, which holds neither log. The 2 warnings are the re-run's, not the bring-up's. |
| *"the fix must distinguish absent / empty / populated — today all three read as 'the phase never ran'"* | **REFUTED.** autoverify has distinguished all three since **v2.8 M256 harden pass 2**: `[ ! -e ]` → absent-warn, `elif [ -s ]` → populated-warn, `else` → `✓`. Four tests in `TestEvidenceAbsenceIsNotEvidenceOfSuccess` pin it. **Empty is the HEALTHY state by design** (M217: the log is written only from failure branches). |
| *"the message asserts a CAUSE from the absence"* | **Partly refuted.** The message already names the alternative — `"(or STACK_DIR is not the bring-up's \$STACK)"` — which is precisely what happened. iter-10's pre-compute quoted the message **truncated at the em-dash**. Third occurrence of §5 rule 10 in this milestone. |

## Cluster / target identified

The residue after the refutation is a real defect, and a bigger one than the pre-compute described:

**`STACK_DIR` is a hand-supplied path with no derivation and no validation**, in a script that already
derives its offset from `--project` and cross-checks it against the registry. Three consequences, all
measured:

1. **A wrong value produces two warnings indistinguishable from real defects** and flips `green:false`.
   That is what produced iter-10's false pre-compute, i.e. it has already cost this milestone one iter.
2. **`dev-stack:298` omits it entirely.** Every `dev-N` bring-up therefore skips the whole cheap-win block:
   no demopatch check, no buildfail check, **no `autoverify.log`, and no `autoverify.json` at all** — the
   machine-readable verdict every grader reads. A skip that reads exactly like a pass (§5 rule 8).
3. **The trap is documented instead of removed.** `CLAUDE.md`'s latency-budget row already records
   *"`autoverify.sh` needs `STACK_DIR`"* as a gotcha to remember. *Prefer a design that cannot express the
   bug over a check that catches it.*

## Hypothesis

Derive `STACK_DIR` from `--project` at the point of use — the same move §2 prescribes and the same one
`target_resolve_offset` already makes for the offset — and the wrong-path failure becomes unexpressible.
Separately, gate the two demo-only receipt asserts on the project's TYPE rather than on whether the caller
happened to set a variable, so deriving the dir for `dev-N` does not manufacture two new false warnings.

## Expected lift

- A standalone `autoverify.sh --project demo-1` (no `STACK_DIR`, or a wrong one) returns the **same verdict**
  as the bring-up's own tail run.
- `dev-N` bring-ups gain `autoverify.log` + `autoverify.json` (currently absent) without gaining warnings.
- The `demo-1` warning count measured through the standalone path: **2 → 0**.

## Phase plan

1. Re-survey (done above) + baseline the `stack-verify` suite.
2. `target_resolve_stack_dir` in `stack-verify/lib/target.sh`; wire it in `autoverify.sh`.
3. Type-gate the two receipt asserts.
4. Regression tests, each mutation-verified RED with a **compiling** mutant (§8 rule 5).
5. **Live negative control** on the running `demo-1`: no `STACK_DIR`; wrong `STACK_DIR`; correct `STACK_DIR`.
6. Close.

## Escalation conditions

- If the derivation cannot be made authoritative without breaking the `anthropos` (main dev stack) path →
  route forward rather than special-case it silently.
- If the type-gate would suppress a check that a demo genuinely needs → stop and re-scope.

## Acceptable close-no-lift outcomes

If the live negative control shows the derivation changes nothing observable (e.g. every real caller already
passes the right dir and dev stacks turn out to write their verdict elsewhere), that falsification — with the
`dev-N` measurement to back it — is the iter's deliverable.
