---
iter: 205
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-205 — the third site-kind: `#` comments

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* Fourth
consecutive iter on the measurement-literal class, at the last site-kind neither census could reach.

**Step 0 — re-survey (mandatory).** `SURVEY-M257x-iter204-the-superset-scan-is-string-literals-only`
was routed one iter ago. Measured before any change: **101 measurement-shaped numbers across 89 sites**
live in `#` comments in rext non-test modules — invisible to **both** existing censuses, because a
comment is a *token*, discarded before an AST exists.

**Cluster / target identified.** The comment is where this repo writes its provenance — *"Measured:
N of M"*, *"iter-122 booked 2 false verdicts"* — so it holds the sentences most likely to be read as
evidence, and it was the one site-kind with no census at all. A visibly stale instance was already in
hand: `claim_census_guard.py:1275` reads *"a wholesale warning over 949 pairs"*, and the census now
prints **1,015** — moved by iter-202's parser fix.

**Hypothesis.** `tokenize` reaches comments; the classifier is already written; the honest shape is the
one iter-203 settled — three classes, report all, assert only the size. The open risk is drift: two
censuses reconstructing one classification rule is what cost iter-202 a 16-against-19 disagreement, so
the rule must be **one function with two callers**, not two copies.

**Expected lift.** The third site-kind censused and ratcheted; the vocabulary's residual re-driven to
zero over the widened superset; the reach audit extended to comments so it stops carrying the blind spot
it exists to find.

**Phase plan.** A: census comments via `tokenize`, sharing the classifier. B: extend the vocabulary for
the comment site-kind. C: extend the reach audit to comments. D: fence + ratchet from the census. E: close.

**Escalation conditions.** If sharing the classifier changes the docstring census's output, that is a
drift the two copies were already hiding — a Fate-1 finding, not a refactor to abandon.

**Acceptable close-no-lift outcomes.** If comments turn out to hold only dated provenance, that is a
complete iter: the site-kind is then measured rather than unmeasured, which is the claim the route
disputes.
