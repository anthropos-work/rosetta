---
iter: 200
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-200 — is rule 77's hazard live in this repo, and is rule 77 the whole hazard?

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

Re-surveyed at HEAD `2fb13f7`. `route_disposition_guard` went **RED** during the survey — iters 198 and
199 had both written `SURVEY-M257x-h42-…` with an ellipsis, and a truncated stem is not an id. Repaired
first, in its own `fix(M257x/199)` commit, before any iter-200 work began; the milestone reads **395
route ids · 342 with a disposition · 1,422 dispositions · 46 closures · 0 malformed · 0 contradictions**.

The remaining harden-filed route is the oldest open one:

> `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` — filed at pass 42, and **explicitly NOT
> a claim that prior proofs were vacuous.**

## Cluster / target identified

Two mechanical questions, both censusable:

1. **Is rule 77's hazard live here?** The rule says a size-preserving mutation can be served from stale
   bytecode, and that only an mtime bump escapes it. That hazard requires a **re-imported `.py`** —
   a `.md` or `.yml` has no bytecode. So: which mutation controls rewrite Python source?
2. **Is rule 77 the whole hazard?** A mutation control is vacuous whenever the second read does not
   reach the disk. Bytecode is one way. **In-process memoisation is another**, and it is broader — it
   ignores mtime *and* size, and applies to every file type.

## Hypothesis

Both populations are AST-decidable: a repeated write to one target inside a function, and a
`path → content` memo in a non-test module.

## Expected lift

No `P`/`N` reading. Deliverable: both populations with their split, an answer to h42 that says *why*
rather than *whether*, and a fence that keeps the answer true.

## Phase plan

1. `mutation_rewrite_sites` — repeated in-place writes, split `py-no-utime` / `py-utime` / `data`.
2. `py_writing_tests` — every test writing a `.py` path at all, as the **upper bound** on (1)'s reach,
   because the repeat predicate cannot see a stage-then-mutate split across a helper.
3. `memoised_disk_readers` — `lru_cache`/`cache` decorated functions and hand-rolled module-dict memos.
4. Fence: the exposed shape at zero, both censuses proven non-vacuous, fire/no-fire controls on the memo
   shapes, and the repeat-predicate's own blind spot asserted rather than described.

## Escalation conditions

- A live `py-no-utime` repeat site → report it; do **not** rewrite other iters' mutation controls in
  this iter.
- The memo population is **not** asserted to zero. These caches are deliberate and load-bearing (one
  guard went from ~1 s to 10.9 s without one). Enumerating them is the deliverable, not removing them.

## Acceptable close-no-lift outcomes

Zero exposed sites is the expected and acceptable result — h42 said as much — **provided** the iter can
say why, and provided the instrument is shown able to find one.

## Explicitly NOT in scope

Changing any existing mutation control; adding cache-clearing to any guard.
