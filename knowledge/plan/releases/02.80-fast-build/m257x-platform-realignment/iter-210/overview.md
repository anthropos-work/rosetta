---
iteration_type: tik
status: in-flight
active_strategy: TOK-08
---

# iter-210 — iter-209 widened one of TWO identical source-set derivations, and nothing in the family compares them

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey before targeting

iter-209 routed *"the guard reads 94 of 2,560 markdown documents"* as sized-not-repaired. Re-surveyed
before continuing down that route, and a **sharper target created by iter-209 itself** appeared one
level over:

`grep -l "def collect_sources"` returns **two** modules — `corpus_citation_guard.py` and
`retracted_pin_guard.py` — carrying the **same four lines** verbatim:

```python
docs = sorted(p for p in (repo_root / "corpus").rglob("*.md"))
docs += [repo_root / n for n in EXTRA_SOURCES if (repo_root / n).is_file()]
```

Measured at today's tree: citation guard **114** sources, retracted-pin guard **94**, symmetric
difference **20 → 0** — exactly the skill documents, and **the divergence is four minutes old.** It was
created by iter-209 repairing one copy.

`§5` iter-190 states the rule this breaks, and states it about this exact failure: *two readers of one
construct must SHARE the derivation — agreement today is not the property.* Two matching literals are
not a shared derivation, and iter-209 proved it the hard way by moving one of them.

## Cluster / target identified

The family's answer to *"which documents are the corpus?"* is spelled per fence. `corpus_citation_guard`
and `retracted_pin_guard` each own a private copy; five other fences declare
`SCAN_FILES = ("CLAUDE.md", "README.md")`, which is **correct for their subject** (they fence those two
documents' prose) and must not be folded in. The two `collect_sources` copies are the pair that answer
the *same* question and answer it differently.

## Hypothesis

Widening the second copy is safe, and the durable repair is not to widen it: it is to leave **one**
derivation in the module the whole family already imports (`fence_provenance`), so a third fence cannot
fork the answer silently.

## Expected lift

- One shared `corpus_sources()`; both fences route through it; the family gains a fence against a third
  private copy appearing.
- Pre-measured: widening `retracted_pin_guard` to 114 sources takes its enumerated pins **2,193 →
  2,201** and its findings **3 → 3, zero new**. Adopted only because that reading is zero.

## Phase plan

Two planned lines:

1. Move the derivation into `fence_provenance.corpus_sources(repo_root)`; both fences call it.
2. Arms: the two fences agree by construction; **no module outside `fence_provenance` may spell the
   corpus-source construct**; plus a mutation control that re-forks one copy and requires RED.

## Escalation conditions

- If routing `retracted_pin_guard` through the shared derivation changes its **findings** (not merely
  its enumerated population), stop — that is a behaviour change, not a de-duplication, and it needs its
  own iter.

## Acceptable close-no-lift outcomes

- The two fences already agree after the move and no third copy exists → the family gains a registry it
  did not have, which is `§5`'s *a class is closed by an enumeration that keeps running*.
