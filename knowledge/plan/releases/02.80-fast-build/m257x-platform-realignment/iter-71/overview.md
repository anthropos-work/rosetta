---
iter: 71
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
closed: 2026-08-04
---

# iter-71 — one knob, and the corpus does not have one ref

**Active strategy reference:** `TOK-05`, step 1 (**fence**) applied to the predicate iters 68–70
measured: *a citation is graded at the ref its own block names.* Carries
`FENCE-M257x-iter68-citation-resolution`, the largest fence route still open.

## Step 0 — re-survey before targeting

`FENCE-M257x-iter68-citation-resolution` was routed as *"re-resolve every corpus citation into a
clone at the adjudication ref and grade HELD/MOVED/GONE/DEAD/UNREACHED."* iter-69 re-scoped it by
measurement to *"every citation is HELD at the ref **its block** names."* That refinement is what
makes it buildable, and `D-M257x-69-1` is the rule underneath it.

Re-derived at this open, against the live corpus:

| | n |
|---|---|
| resolvable citations | **125** |
| block names **exactly one** ref that resolves in that citation's own clone | **31** |
| block names **more than one** resolvable ref (a contrast, not a pin) | **12** |
| block names no resolvable ref | 82 |

**A quarter of the class was being graded against a file it never claimed.** `CITE_REF` is a single
process-wide knob; `backend.md:39` pins to `b948604` v1.366.0 and `shared_libraries.md:79` was
re-pointed to `9d00a313` v1.367.0 in iter-69. Read at one ref, one of them is always wrong — and
iter-68 measured the consequence directly: the same corpus is **GREEN at origin HEAD and 4-findings
RED at the pinned build ref.**

## Cluster / target identified

Make `anchor_construct_guard` resolve **per citation**: read at the ref that citation's own block
names, fall back to the `CITE_REF` ladder when it names none, and refuse to guess when it names two.

## Hypothesis

Per-block refs make the guard's verdict a measurement rather than an artifact of one env var, and the
31 stop being graded against the wrong file — without weakening `CITE_REF=worktree`, the escape
hatch that asks *what does the checkout say*.

## Expected lift

- 31 citations adjudicated at the ref they name, reported as such.
- The ambiguous 12 **counted**, not hidden inside the default.
- The corpus stays GREEN, and `CITE_REF=worktree` still goes RED — the demonstration iter-68 used.

## Phase plan

- **A** — measure the per-block class (**done at open**).
- **B** — build `block_ref()` + thread it through the anchor loop; report where each ref came from.
- **C** — RED-before-trusted: mutants over every decision the new code makes, plus a no-op control.
- **D** — gates.

## Escalation conditions

If the fence cannot be made to go RED on an inverted mutant, it does not ship. If per-block
resolution turns the live corpus RED, the findings are adjudicated as corpus defects in this iter —
they would be claims that are false at the ref they themselves name, which is the strongest kind.

## Acceptable close-no-lift outcomes

That all 31 grade identically at their pin and at the default would falsify the premise and be a
complete iter: the deliverable would then be the **measurement that the knob never mattered**.
