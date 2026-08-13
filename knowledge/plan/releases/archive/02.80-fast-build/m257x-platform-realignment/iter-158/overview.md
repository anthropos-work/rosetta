---
iter: 158
milestone: M257x
iteration_type: tik
status: closed-fixed
date: 2026-08-08
---

# iter-158 — the routed repair was wrong, and the disclosure shipped one iter earlier was too broad

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

**Step 0 — re-survey.** Took `FIX-M257x-iter156-cannot-run-sniff-reads-merged-stream`, routed two iters
ago with the note *"narrowing it needs its own evidence."* The re-survey produced that evidence in one
command — and it **refutes the narrowing**.

**Cluster / target identified.** `guard_family`'s rc-0 `CANNOT RUN` / `Nothing was checked` sniff, which
reads merged stdout+stderr; and — surfaced by the same measurement — the noise classifier iter-156 shipped,
which flags a stderr line when it does not `speaks_for` the guard.

**Hypothesis.** Both are authorship questions on the same stream, so both are answered by one classifier:
*which lines did the guard write, and which did the interpreter inject?*

**Phase plan (two planned steps — declared here so the tripwire counts against this shape).**
1. Measure whether guards speak on stdout or stderr, and settle the routed narrowing on that measurement.
2. Repair the sniff and the noise classifier through one authorship function; fence both.

**Escalation conditions.** If any guard's could-not-run message turns out to be unattributable, route
rather than guess.

**Acceptable close-no-lift outcomes.** A refutation of the routed repair, recorded with its number, is a
complete iter under this protocol even if no code changes.
