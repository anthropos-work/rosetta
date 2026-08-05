---
iter: 90
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
closed: 2026-08-05
---

# iter-90 — the demopatch asymmetry: journal the observed pre-state, and test the CONJUNCTION

**Type:** tik, under `TOK-05`.

## Active strategy reference

`TOK-05: stop repairing claims; fence the predicates under them`. This iter is squarely inside it: the
repair unit is a **predicate** — *"a demo-patch reverts itself and leaves the clone git-clean"* — and it is
closed by making the false case **unrepresentable**, not by re-pinning the two baselines that happen to be
stale today.

## Step 0 — re-survey (mandatory)

TOK-05's next-tik direction is superseded by an explicit user decision resolving iter-89's `user-blocker`.
Re-surveyed at open, and the target is confirmed still live and still meaningful:

- `stack-demo/next-web-app` is still dirty in exactly the two files iter-89 named
  (`packages/core-js/src/constants/urls.ts`, `packages/ui/src/NavBar/NavbarTop.tsx`).
- The reproduction is intact and now captured verbatim (three manifests, live):
  `status` → `patched`, `revert` → `REFUSE: … is neither pre nor post`.
- The demopatch suite is **50 pass / 2 fail**, both in `TestRealManifest`, both reading the live dirty clone.

## Cluster / target identified

iter-89 root-caused four failures to **one** structural defect and escalated the repair as a design choice.
The user has now decided: **option (b) — journal the observed pre-state at apply time; revert restores
exactly that.** The blocker is resolved, so this iter lands it.

The finding that outlives the fix, and the reason the test matters more than the patch:

> **G2 (refuse on drift) and G5 (always self-revert) cannot both hold once the base is allowed to move.**
> Every test asserted them **separately** and each passed. Nothing asserted their **CONJUNCTION**, and that
> is where the defect lived.

Confirmed in the suite by inspection: `test_g2_drifted_sha_with_an_INTACT_anchor_SELF_HEALS`
(tests/test_demopatch.py:196) applies onto a drifted base and **never reverts**;
`test_g5_revert_on_drifted_refuses_without_force` (:426) reverts a **manually** drifted target and asserts
the refusal. Neither composes apply-onto-drift with revert. Both pass. The defect sits in the gap.

## Hypothesis

Recording the **observed** pre-image at apply time makes revert independent of the recorded baseline, so
revert becomes exact on a drifted base while apply keeps its self-healing. Both G2 and G5 then hold at once,
and the conjunction test that fails today passes.

## Expected lift

No movement on the clause-5 reading (this iter takes none). The deliverable is a **repaired mechanism plus
the conjunction test that would have caught it**, and the removal of a live dirty-clone condition that is
currently failing 2 tests in the demopatch suite.

## Phase plan (declared multi-step shape — the scope-creep tripwire counts against THIS list)

1. **Capture the reproduction as a regression test FIRST** — write the G2∧G5 conjunction test and
   demonstrate it RED against unfixed code; record the mutant signature. (Sequencing is mandated by the
   user's Decision 2: the dirty clones were deliberately left so the defect reproduces without being
   re-created; that state is not to be spent before it has bought a test.)
2. **Land the (b) fix** — journal the observed pre-state at apply; revert restores exactly it; the journal
   is itself cleaned up on successful revert (no leaked per-apply state).
3. **THEN clean the two dirty files** in `stack-demo/next-web-app`, and record plainly that journaling
   **cannot** retroactively revert them — they were applied before any journal existed. That is a real
   limitation of (b), not a workaround to hide.

## Escalation conditions

- If (b) measurably cannot hold both G2 and G5 at once, **say so with the measurement** and re-escalate
  rather than shipping a fix that only appears to work.
- If the journal cannot be written outside every git clone, escalate: a journal that dirties the clone
  would defeat the promise it exists to keep.

## Acceptable close-no-lift outcomes

A demonstration that (b) is the wrong reading — with the measurement that shows it — is a complete iter.
The user asked for the decision to be re-derived, not assumed.

## Routed forward at open (NOT this iter's scope)

- The guard-family **stale-clone freshness fence** (the user's carried-forward correction). Pre-flight
  measurement already taken and recorded in `decisions.md`; the fix is iter-91.
- The **7-guard conjunction-pair sweep** (Decision 1 item 3) — iter-91.
- The **M810 `cms`-vs-`jobsimulation` split sweep** — iter-92.
- The clause-5 **reading** at `0c91421` — iter-93, and only if the guard family is green on a fetched clone.
