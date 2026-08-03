---
iter: 52
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-03
---

# iter-52 — repair the UNION of 18, at the smallest edit that makes each claim true

**Active strategy:** [`TOK-03: repair the UNION, shrink the estimator, make the edits smaller`](../decisions.md#tok-03-repair-the-union-shrink-the-estimator-make-the-edits-smaller--2026-08-03).
This is TOK-03's own pre-registered `Next-tik direction`, executed verbatim: **`FIX-M257x-iter50-union-set`**,
under move 3's minimal-edit discipline, with move 4's two blind pre-commit diff readers as the post-condition,
and **no reading taken in this tik**.

## Step 0 — re-survey before targeting (mandatory)

The TOK-directed target is **not stale**, and each leg was measured at open rather than assumed:

- **Platform origin re-fetched: `2adcf714`, unchanged.** Re-scope trigger stays at **occurrence 1 of 2**.
- **All thirteen ground-truth clones re-read and byte-identical** to the ones readings #9 and #10 recorded.
  The union key is therefore valid against the same ground truth that produced it.
- **The corpus is still unrepaired.** `git diff --stat 47c9b7d..HEAD -- corpus/ CLAUDE.md .claude/` is empty;
  every commit since is `knowledge/plan/**`. The 18 stand where the readings left them.
- **The fence agrees, independently of the ledger prose.** `claim_twin_guard --report` at open:
  **RED — 31 published sites, 19 unique**, derived from 100 claims across 24 ledger files. It reaches most
  of the union and not all of it (it has no purchase on #15's variable-name block or #14's cross-file
  synthesis), which is why the union key — not the fence — is the repair target.

No substitution. The target named by TOK-03 is the target.

## Cluster / target identified

The **18** of [`fixture-18.md`](fixture-18.md), captured before this iter edits a byte. Not iter-49's 14:
**8 of the 18 were named by exactly one of the two readings**, so a single-reading repair leaves 8 standing
by construction. That is the arithmetic TOK-03 was authored on.

## Hypothesis

The metric has been stuck because **repair coverage**, not reading sharpness, is the binding constraint —
and because **the repair itself manufactures roughly as much as it removes** (the induced term ran
9 → 7 → 2 → 7 → 4, and iter-49's rose 2 → 7). So this iter changes two things at once and expects each to
move a different term:

1. **Coverage** — repair 18 rather than 14 (78 % of `N̂ ≈ 23` rather than 61 %).
2. **Induction** — shrink the surface new defects can live on. Every induced class iter-49 measured
   (paraphrase leak, overshoot-in-new-text, wrong-mechanism-correctly-cited) is a property of **rewriting**,
   and none is mechanically reachable. The only lever on an unreachable class is to make less of the thing
   it lives in. **Prefer DELETION > minimal scoping edit > rewrite**, and count the words each repair adds.

## Expected lift

Not a blocker count — this tik takes no reading, deliberately. The measurable outputs are:

| output | expectation, pre-registered so it can be refuted |
|---|---|
| union rows repaired | **18 of 18**, each by claim and tree-wide, or named as routed with a reason |
| `claim_twin_guard` | falls from **31** hits / 19 sites; every survivor named and explained, none silent |
| **net words added** | **≤ 0 across the pass** — the repair should be a net deletion |
| pre-commit reader findings | **> 0.** A pass that finds nothing has not been read; two readers who agree on nothing are also a signal |

The `N̂` re-estimate this iter owes is a **derivation from the coverage arithmetic**, not a new reading —
readings #11/#12 are iter-53's job, and taking one here would not be blind to this iter's own work.

## Phase plan

- **A** — capture the perishable fixture (done first, committed before any repair byte).
- **B** — repair the 18, by claim, tree-wide, smallest edit first; per-claim added-word ledger.
- **C** — **two blind adversarial readers on the repair diff, pre-commit** (TOK-03 move 4). They are given
  the diff and the ground truth, and are barred from the answer key and from each other.
- **D** — triage the readers' findings, fix, re-run the fence as post-condition.
- **E** — re-estimate `N̂` from the coverage arithmetic and state the derivation; close.

## Escalation conditions

- A union row whose correct form is **not already adjudicated** in a ledger → **route it, do not repair it.**
  Deriving a fresh verdict during a repair pass is precisely how rule 18's highest-risk text gets written
  (§5 rule 19, *"what a claim-scoped repair must NOT do: adjudicate"*).
- The ratchet refusing the repair commit → **record why and fix the fence or the repair**, never `--no-verify`
  silently. iter-50's bypass is a recorded decision (`D-M257x-50-7`), not a precedent.
- A second platform commit invalidating an alignment attempt → `EXIT_REASON: re-scope-trigger`, escalate.

## Acceptable close-no-lift outcomes

- The readers finding that the repair induced **more** than it removed would be a real result and would close
  this iter honestly at `closed-no-lift` — it would refute TOK-03 move 3 on its first outing, which is
  exactly what a pre-registration is for.

## What this iter does NOT do

- **It does not take a clause-5 reading**, and does not claim clause 5 is nearer than the estimate supports.
- **It does not re-cut clause 5.** The user has ruled three times.
- **It does not touch the existing fixtures.** `iter-50/fixture-14.md`, `claim_twin/red/`, `claim_twin_iter48/`
  and the iter-45 key stay byte-identical; they are supposed to contain false claims.
- **Zero platform-repo edits.**
