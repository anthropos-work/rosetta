---
iter: 92
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
closed: 2026-08-05
---

# iter-92 — the M810 sweep: mostly already done, and the one place it was wrong was a FENCED claim restated UNFENCED

**Type:** tik, under `TOK-05`.

## Active strategy reference

`TOK-05`. The predicate here is *"M810 is one event"*, and its legal set is derivable from two platform
repos' terraform + CI — which is exactly TOK-05's *adjudicate against platform artifacts, never against
another document*.

## Step 0 — re-survey (and it changed the target)

The brief carries iter-89's finding that **~14 passages across 11 files** treat M810 as one future event.
Re-measured at HEAD before targeting: **15 files, 40 occurrences**, and the great majority **already state
it as uneven** — `corpus/README.md:16` (*"is **uneven**: **landed for jobsimulation**"*), `CLAUDE.md:189`,
`backend.md:36` (*"UNEVEN — do not state the two together"*), `cms.md:9`, `jobsimulation.md` throughout,
`services/README.md:17`, and both fenced map rows.

**So the sweep was substantially done, and re-doing it would have been re-landing landed work** (Phase 1
Step 0's whole purpose). The target was re-surveyed under the same strategy and substituted: not *"propagate
the split"* but *"where does the corpus still get M810 wrong, measured against the platform now?"*

## Cluster / target identified

Two things the re-survey found that the original framing would have missed entirely.

## Hypothesis

The residual M810 error is not un-propagated news; it is (a) news that arrived **after** the propagation,
and (b) a fenced claim being restated **more strongly** in unfenced prose.

## Expected lift

No reading. The deliverable is the corrected per-service M810 state and the named unmeasurable boundary.

## Phase plan

1. Measure M810's true per-service state from the platform repos (terraform + CI), not from the corpus.
2. Correct what is wrong; leave what is right alone and say so.
3. Name the unmeasurable class rather than asserting through it.

## Escalation conditions

- If the prod-side state turns out to be measurable after all, say so and measure it.

## Acceptable close-no-lift outcomes

Establishing that the sweep was already complete, with the measurement, is a complete iter.
