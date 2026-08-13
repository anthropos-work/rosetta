---
iter: 188
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-09
controlling_strategy: TOK-08
---

# iter-188 — a container-keyed exclusion has an environment-dependent SIZE and an environment-independent SHAPE

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey before targeting

iter-187 routed `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` **with a mechanical
selector**, which is the thing iter-185's residual lacked: *a registry keyed by a CONTAINER whose
justifying reason is a property of the container's CONTENTS.* This iter runs that selector rather than
re-surveying by judgement.

Applied over the 30 module-level registries in `stack-core`, the selector's strongest hit is
`claim_census_guard._SKIP_DIRS` — 12 **directory names** pruned from the `os.walk` that builds the
**basename index citations are resolved against**. Every member is justified by a property of the files
inside it (*build junk, not source*), and a file the walk never reaches is a file a citation cannot
resolve to.

Measured 2026-08-09, and the two environments disagree by construction. **The first reading of this iter
was wrong about which is the default** and is corrected here rather than quietly: `clones` falls back to
`root / "stack-demo"` when `--clones-root`/`CLONES_ROOT` is unset (`:768`), and that directory **exists**
on this box — so the default run is the large one, not the small one.

| roots walked | files indexed | files pruned | share pruned |
|---|---:|---:|---:|
| **`corpus/` + `stack-demo/` — the DEFAULT run here** | 21,610 | **50,357** | **70.0 %** |
| `corpus/` alone (`stack-demo` absent) | 92 | 0 | 0.0 % |

**70 % of the walk's candidate files never enter the index citations are resolved against, and no output
said so.** Of the 12 names, **2** did any pruning here (`node_modules` 3 directories, `test-results` 2),
**5** were inert (`vendor`, `dist`, `build`, `coverage`, `__pycache__`) and **5** can never fire at all.

## Cluster / target identified

Four properties, and the split between them is the point:

1. **Size is unprinted.** The guard prints `scope … -> 92 files` and never that the walk pruned 0 (or
   537). `§8`/iter-186's rule — *print the scope beside the total, derived, in the output* — was applied
   to a census and never to a **walk**, which is the same claim one layer down.
2. **5 of 12 members are subsumed by an unnamed adjacent rule.** The same line reads
   `d not in _SKIP_DIRS and not d.startswith(".")`, so `.git`, `.next`, `.venv`, `.turbo`,
   `.pnpm-store` cannot ever be the reason a directory was pruned. The registry over-states its own
   work, and this is **environment-independent** — provable from the source, no tree required.
3. **No member carries a reason** (`§5` r8), and there is no staleness arm. 8 of the 12 pruned nothing
   in either tree measured; nothing would ever say so.
4. **The consequence is a false defect, not a missed one.** This index is what `materialize()` resolves
   citations through. A cited file under a pruned directory is *unfindable*, which grades as
   unresolvable — the failure mode iter-122 already paid for once with stale substrate.

## Hypothesis

Assert the **shape** (every member reasoned; no member subsumed by the dot-rule; the registry is the only
prune rule, or the other rules are named too) and **print the size** (per-entry prune counts, derived, in
the guard's own output) — because the size is environment-dependent and an assertion over it would be
either vacuous in the default environment or wrong in a clone-shaped one.

## Expected lift

No `P`/`N` reading (`§9`). 1 unprinted 31.3 % exclusion made visible, 5 redundant members identified,
12 reasons supplied, ≥3 new arms mutation-proven RED, and a `§8` rule that names the shape/size split.

## Phase plan

- **A** — measure both environments (done above; re-derived in code).
- **B** — repair: `_SKIP_DIRS` → a reasoned dict; a `prune_census()` derivation; the guard prints it.
- **C** — fence the shape, in both directions.
- **D** — mutation-prove; run under both runners.
- **E** — publish the `§8` rule; route residuals.

## Escalation conditions

- If any current corpus citation resolves into a pruned directory, that is a **live false defect**, not a
  latent one — measure it and say so rather than folding it into the shape finding.
- If removing the dot-subsumed members changes what the walk prunes in any tree, the subsumption claim is
  wrong; keep the members and record the refutation.

## Acceptable close-no-lift outcomes

- The dot-rule subsumption turns out not to hold (e.g. the walk is reached by another path with a
  different rule) → the finding shrinks to the unprinted-size half; record the falsification and close.
