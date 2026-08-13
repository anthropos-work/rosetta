---
iter: 44
milestone: M257x
iteration_type: tik
iter_shape: tooling
status: closed-fixed
opened: 2026-08-02
---

# iter-44 — `FENCE-M257x-iter44-repair-postcondition`: the fence runs at the commit, not at the next audit

**Active strategy:** [`TOK-02` — *fence the prose the way the anchors are fenced*](../decisions.md#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02).
**Step 2 of five, and step 2 only.** Step 1 (the claim-twin fence, watched RED on 16 of 18) closed at
iter-43. Step 3 (the two mechanical fences) is the next tik; step 4 (repair) and step 5 (the reading)
follow. **Nothing is repaired in this iteration** — `D-M257x-42-3` still holds, and for the same reason:
the 18-defect corpus is the only fixture with a known answer key, and it is perishable.

## Step 0 — re-survey (mandatory before targeting)

| check | result |
|---|---|
| fixture intact | `tests/fixtures/claim_twin/red/` — 18 files, byte-identical at open |
| live residual | claim-twin fence: **18 sites**, 36 claims derived, 112 files scanned — unchanged from iter-43 |
| platform origin | re-fetched at open: `2adcf71` — **unchanged**; re-scope trigger stays at occurrence 1 of 2 |
| TOK-02 target still current | yes — step 2 is named explicitly and nothing has absorbed it |

## Cluster / target identified

TOK-02's own arithmetic names this as the highest-value remaining move, and it is the only step that
attacks the **cost** term rather than the residual term:

> *"Of the 9 repair-induced blockers, **8 are self-contradiction** … a claim repaired at one site and left
> standing at another."*

iter-41 measured the process's fixed point: a pass repairs 18 and induces ~9. A fence that runs at the
**next audit** cannot change that number, because the induced defect is already committed and already
counted by the time it runs. The only place the induced term can be attacked is **between the repair and
the commit**.

## Hypothesis

Convert the claim-twin fence from an audit instrument into a **commit post-condition with a monotone
baseline**: the set of published sites restating an adjudicated claim may **shrink or stay**, never grow.
An induced self-contradiction is, by construction, a *new* `(file, claim)` pair — so the class TOK-02
measured at 8-of-9 becomes unrepresentable in a commit rather than merely detectable in the next pass.

Three properties decide whether this is a fence or a slogan:

1. **The registry is DERIVED, not hand-listed.** §2 deleted a hand-maintained tuple; a post-condition
   runner with a hardcoded list of fences is the same tuple wearing a new hat. Every guard on disk must
   declare its own `FENCE_KIND`, and an undeclared guard goes RED naming itself (iter-08's derived-scope
   pattern).
2. **The baseline is keyed line-number-free** — `(fence, path, claim_id)`. Keying on `file:line` would
   turn every ordinary edit above a known site into a fake NEW finding, and §8 rule 6 says where that
   ends: a fence that cries wolf gets disabled.
3. **It is watched going RED on an induced defect**, not only on today's 18 — the RED that matters is the
   one this instrument exists for and which no fixture currently contains.

## Expected lift

**Zero on clause 5, by construction** — this iteration repairs nothing, so the residual stays at 18.
The deliverable is the post-condition itself. Success criterion: the post-condition fires on a
synthetically induced self-contradiction in a repaired tree, with a surviving no-op control and an
inverted mutant, and the live corpus grades GREEN-against-baseline at 18.

## Phase plan

| phase | work |
|---|---|
| **A** | re-survey (above) — fixture, residual, platform origin |
| **B** | `FENCE_KIND` declaration on every `stack-core` guard + the derived registry |
| **C** | `repair_postcondition.py` — runner, monotone baseline, `--accept`, hook installer |
| **D** | **watch it go RED** on an induced defect in a repaired tree; GREEN control on the live tree |
| **E** | tests + mutation battery (no-op control that survives + inversions + a reporting-path deletion) |
| **F** | protocol doc §8 + `stack-core/README.md`; close |

## Escalation conditions

- Platform origin moves → `re-scope-trigger` (occurrence 2 of 2), stop.
- The post-condition cannot be made to fire on an induced defect → that falsifies the step; close-no-lift
  with the falsification recorded, and TOK-02 step 2 is re-planned rather than declared shipped.

## Acceptable close-no-lift outcomes

- The induced class turns out **not** to be expressible as a new `(file, claim)` pair — measured, not
  argued. That would refute TOK-02's step 2 premise and is a first-class result.
