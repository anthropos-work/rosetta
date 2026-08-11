---
iteration_type: tik
status: archived
opened: 2026-08-07
corpus_at_open: bfb2660e48862890c1796cccdfd090c4e774b7ae
rext_at_open: 415240f6f48baa816e473f6cf9a0e2320225c830
platform_at_open: 0c91421df
---

# iter-130 — clear the routed set, then close the fence gap the `ai` row exposed

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them* (the
USER's re-scope, sealed at iter-117's `577446b`). This iter runs its **repair-and-fence** half. The
grading half — the reading — is deliberately the NEXT iter, because *the reading must not measure a tree
that is still being edited*.

## Step 0 — re-survey before targeting

`TOK-08`'s next-tik direction is *"continue sweeping classes"*, and run 83's brief refines the target
rather than replacing it: **close the routed residual first, stop repairing, then read.** Re-measured at
the open:

| | value |
|---|---|
| guards | **18 GREEN · 0 RED · 4 not-run** (needs `--range`/`--ledger`) |
| claim census | **1,141** (baseline 1,164) — down from the 1,150 iter-129 published at its close, because iter-129's post-close slices 3/4 landed after it |
| last reading | **iter-119**, `P = 22` / `N = 28` — **eleven iters old** |
| gate | **4 of 5** |

The target is current and meaningful: the routed set is enumerated with `file:line` on both sides and
has not been touched.

## Cluster / target identified

`iter-129/progress.md` § 2c and § 3c routed ~24 lower-consequence findings as
`FIX-M257x-iter129-sweep-residual` rather than half-repairing them. They are this iter's Priority 1.

## Hypothesis

Three things, in this order:

1. **Priority 1 — close the routed set.** Each finding RE-VERIFIED at a ref before anything is touched;
   an upheld claim is a result, not a miss. The `@anthropos.work` predicate is repaired **only after
   measuring the width of this run's own enumeration** (§5 rule 57 — the repair that found 14 sites had
   a first regex that reached 4).
2. **Priority 3 — the library-row fence gap.** The `ai` row sat wrong *inside clause 3's own fenced
   deliverable* and `platform_alignment_guard` could not have caught it: it fences the map against
   `repos.yml`, and **`ai` is a module, not a clone**. Either extend the fence to the module graph or
   make the map declare its unfenced class. **Do both** if the fence is cheap enough to ship with
   controls that fire.
3. **Then stop.** No further repair; the tree freezes for iter-131's reading.

## Expected lift

**No `N` movement is claimed and none is measured — this iter takes no reading.** (The protocol's
iter-108 refinement: an iter that took no reading has an UNMEASURED metric, not an unmoved one, and does
not count toward the 3-no-prog tok streak. `TOK-08` declared the sequence; this is its repair step and
the next iter is its grading step.)

The lift this iter is graded on is its **planned deliverables**: the routed set closed, and the library
rows fenced or declared.

## Phase plan

- **A** — re-verify + repair the routed set (parallel, disjoint file sets, every finding checked at a ref)
- **B** — the `@anthropos.work` predicate: measure enumeration width FIRST, then repair, then re-run
- **C** — the fence: assertion G against the module graph, with a mutation battery and an anti-vacuity
  control that can fire; plus the map declaring which row classes each assertion reaches
- **D** — guards + targeted tests; commit; **stop repairing**

## Escalation conditions

- A routed finding that turns out to require a platform-repo edit → route, never edit (v2.8 constraint).
- A security-surface claim found inverted → report plainly and first, grade by consequence not class,
  and **no legal/compliance/policy escalation** (§5 rule 48).
- The fence going RED on a row I cannot settle at a ref → disclose as unclonable, do not assert.

## Acceptable close-no-lift outcomes

If the routed findings were largely UPHELD on re-verification — i.e. the prior sweep over-reported —
that is a first-class result and would be reported as such, because it would mean the sweep's own
precision is the thing to fix rather than the corpus.
