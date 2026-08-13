# iter-128 — decisions

## `D-M257x-128-1` — a budget whose content has no other owner is not enforceable by moving content

`context.md § state.md contract` prescribes one repair method: **move content to its owner, never trim.**
Applied to the `phase:` field it worked exactly as designed — 1,985 → 861, and every frontmatter field is
inside its budget for the first time.

Applied to the **body** it has no target. Three probes each asked *"who else owns this?"* and each
returned **`state.md` itself**: seven § Standing backlog items `roadmap-vision.md` does not mirror, the
M255 provenance clause (`roadmap.md` carries the numbers but not the provenance), and the process flags.

**Decision: report the overage with its number and route it; do not trim, and do not silently raise the
budget.** Trimming would delete the sole record of seven backlog items — the same failure mode this iter
caught one paragraph earlier with `org-repos.md`. Raising the budget is a legitimate answer but it is a
**decision about the contract**, not an edit, and the contract is one day old and shared by eleven
skills; changing its numbers unilaterally on their first enforcement is not this iter's call.

**The general form worth keeping:** a per-field budget is enforceable only where the field is a
*duplicate*. Where the field is the **sole owner**, the budget is a statement about how much unique
content the document is allowed to hold — a different and much stronger claim, and one that was never
measured before being written down. **Budget the duplication, not the ownership.**

## `D-M257x-128-2` — re-measure a rule's accuracy on the population you are applying it to

iter-124 measured `R4` at **66.7 %** on C1. The cheap move was to reuse that rate to correct the
complement's split: same rule, same corpus, one multiplication.

**Decision: draw a fresh seeded sample and measure `R4` on the complement.** It came out **76.7 %** —
close enough that the imported number would not have embarrassed anyone, which is exactly why the
principle needs stating while the stakes are low. An imported rate is an **assumption wearing a
measurement's notation**, and this milestone's whole subject is that substitution.

**And the audit's own frame is disclosed rather than smoothed:** iter-124 sampled the whole class and
drew 9 R4 members; this run sampled **R4-only** and drew 30. The two accuracies are **not a before/after
pair** and the 66.7 → 76.7 movement is **not a trend**. Reporting it as improvement would have been a
free, flattering, false sentence.

## `D-M257x-128-3` — `fix = 0` over an unread population is a floor, and the sentence must say so

The complement triage prints `fix 0 = 0.0 %`. That number is **structurally incapable of being anything
else**: `fix` is a hand-adjudicated input (`R0`/`FIX_SITES`), the triage cannot decide falsity, and this
run's reading was aimed at **C1** by the consequence-ordering rule.

**Decision: publish it as `a FLOOR of unknown height` with the reason attached, never as `0.0 %`.** A
percentage implies a measurement was taken. None was. The complement's false-claim count is **UNMEASURED**
and is recorded that way in the split table itself — not in a footnote, because the table is what gets
quoted.

This is the same discipline iter-124 applied to its own `fix = 4`, one population over.

## `D-M257x-128-4` — two fence gaps found by reading, both unreachable by the current fence shape

Two defects this run introduced were caught by **reading**, not by the 22-member family:

1. **Mojibake** (`â\x80\x94` for an em dash) from a `unicode_escape` round-trip. The only two encoding
   mentions in the family are `except UnicodeDecodeError` handlers — and **mojibake is valid UTF-8**, so
   it decodes without error. A decode-error handler cannot see it *by construction*.
2. **A relative markdown link one level short** (`../` × 5 where the depth is 6). `corpus_citation_guard`
   was GREEN through it because it enumerates **citations**, not relative links in plan documents.

**Decision: record both as named gaps; build neither this iter.** A fence is in scope for this milestone,
but a fence built in the last third of a run, against a defect class of n = 1 each, would ship without
the mutation and anti-vacuity controls the standing rule requires of every fence in this family. **A
fence that cannot demonstrate it can fire is worth less than a recorded gap**, because it converts an
open question into a false green. Routed as `FIX-M257x-iter128-encoding-and-link-fences`.

Both are worth noting for the same reason: they are **author-side** defects, and the milestone's standing
question is whether the fences catch *the author*. On these two classes the answer is measured: **no.**
