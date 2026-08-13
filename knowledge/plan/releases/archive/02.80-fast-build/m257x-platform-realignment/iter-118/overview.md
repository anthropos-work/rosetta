---
iter: 118
milestone: M257x
iteration_type: tik
iter_shape: census
status: closed
opened: 2026-08-07
---

# iter-118 — `TOK-08` class 2: platform-source citation resolution

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— the user's re-scope, binding. This is the **second and last** class on iter-117's pre-registered class
list, which may only grow and did not.

## Step 0 — re-survey before targeting

| check | result |
|---|---|
| Is class 2 still the named next target? | **Yes** — iter-117 closed routing it, and it is #2 on the sealed class list |
| Does a fence already exist for it? | **Yes — `anchor_construct_guard`**, and `TOK-08` says *build **or extend***. So: extend, do not rebuild |
| Is it already green? | **Yes — 0 findings.** Which is exactly why the interesting question is not *is it green* but **over how much of its subject** |

**Type selection:** tik. The 3-no-prog tok-trigger cannot fire — of the last three tiks, **115 and 117
took no reading** and §9's refinement reads UNMEASURED as UNMEASURED, not as unmoved. Both said so in
their own closes, in those words, and `TOK-08` declared the read-last sequence in advance, which is that
rule's second guard-rail.

**Hypothesis.** A fence that is green over a subject it only partly reaches is a census over half a class.
The lever is **reach**, not findings.
