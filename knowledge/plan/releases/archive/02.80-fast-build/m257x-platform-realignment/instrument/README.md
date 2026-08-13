# The audit instrument — now a durable artifact, because it was not one

**Read this before taking any reading.** §5 rule 22 records that M257x *"froze its instrument at iter-41 and
never touched a knob again."* **That claim was false in a way nobody could see**, and iter-53 found out how:

**the briefing — "the whole instrument", in its own words — was never captured in a versioned artifact.** It
lived in `.agentspace/scratch/work-m257x/iter50-briefing.md`, which is **git-ignored**, so it appeared in no
diff, no commit, and no iter directory. Every pass "held the instrument fixed" by **re-authoring the briefing
from its own description in the previous iter's `overview.md`.** That is not freezing an instrument; it is
re-building one from a summary each time.

iter-53 did exactly that, unknowingly, and its two readings came in at **32 and 26** against #9/#10's **14 and
7**. Re-grading iter-53's union under the canonical grading rule *verbatim* brings it to **23 and 23** — so
roughly half the apparent jump is grading drift and half is not. **Neither half is a statement about the
corpus**, which is precisely what an unfrozen instrument costs.

## Files

| file | what it is |
|---|---|
| `briefing-canonical-iter41.md` | **THE instrument.** The briefing readings #9 and #10 were taken with, recovered from the git-ignored scratchpad and committed here. **Use this file verbatim. Do not re-author it.** |
| `briefing-iter53-AS-RUN.md` | what iter-53 actually ran, preserved as the evidence of drift — **not** for reuse |

## The drift, itemized

| knob | canonical (`briefing-canonical-iter41.md`) | iter-53 as-run |
|---|---|---|
| undercount | **MINOR**, explicitly | not carved out; "wrong counts" listed under BLOCKER |
| omitted list member | **MINOR**, explicitly | not carved out |
| line drift | **MINOR**, explicitly | not carved out; anchor checks pushed toward BLOCKER |
| tie-break | *"if you cannot cite the refutation, it is not a blocker"* | *"when in doubt, book it as a BLOCKER"* — **inverted** |
| §5 rule 24 | not present (it postdates the briefing) | promoted to a standing instruction |
| per-file repair history | present | replaced by a generic "the corpus has been repaired many times" |

**The tie-break inversion is the load-bearing one.** The canonical rule resolves doubt *downward*; the as-run
rule resolved it *upward*. A grading rule that resolves doubt upward cannot produce a number comparable to
one that resolves it downward, and no amount of care in the reading recovers that.

## The rule this establishes

**An instrument that is described rather than stored is not frozen.** If a measurement's procedure lives only
in prose that the next run re-instantiates, then every "held fixed" is a re-authoring, and a rising series
measures the re-authoring. Store the instrument as a file, in the repository, and diff it.
