# M260 — Decisions

## `D-M260-1` — a comparison, not a constant

The floor's real question is *"is this capture broken?"*, and a constant answers a different one:
*"is this capture big?"*. Those coincide only while the population is stable. The replacement compares
against something **measured** — the surface's own capture history — so it keeps working across a
consolidation, and it catches partial captures the constant waved through.

**Why a ratio and not a source-count comparison.** Counting the source under the same predicate and
requiring equality would be the tightest check, but it cannot see the failure the floor exists for: an
unprovisioned source returns 0 rows and a source-count check happily agrees that 0 of 0 were captured.
The two rungs together cover both — non-empty catches the empty source, the ratio catches the collapse.

## `D-M260-2` — the acknowledgement is a REASON, not a `--force`

`AcceptShrink` takes a string and writes it into the manifest. A bare boolean would let a collapse
through while teaching nothing to whoever reads that manifest a release later; a reason travels with
the snapshot and can be audited.

## `D-M260-3` — `shrinkRatio = 0.5` is a separator, not a tuned value

It is the coarsest number that still separates the only two cases anyone has hit: an unprovisioned
source (orders of magnitude, always caught) and ordinary catalogue churn (single-digit percent, never
caught). It is **not** a claim that a 49 % drop is acceptable — it is the point past which the tool
stops guessing and asks a human.

## `D-M260-4` — the five net-new tables go to M261, with M260's measurements attached

D-M259-3 routed them to "M260/M261". They land in **M261**: declaring a capture surface whose column
lists no capture has ever validated is a confident guess, and validating it is exactly M261's job.
M260 did the measuring so M261 does not repeat it — the two redirect tables carry **zero** org-scoping
(`PureReference`), the two translation tables exist and are parent-scoped, and **`taxonomy_canon_state`
was NOT found** in the ent schema dir at `4bccda085` despite appearing in the commit range, so M261
must confirm what it is rather than declare it blind. M261's `overview.md` was edited to carry this.

## `D-M260-5` — two bugs this milestone's own tests caught, recorded because both are reusable

1. **A truncated int floor is not a ratio.** `int(11 * 0.5)` is 5, so `rows < 5` let 5-of-11 (45 %)
   pass. The boundary test caught it. Compare as a ratio.
2. **The fence missed the exact line it was written for.** Its first version blanked quoted strings
   before looking for the subject — and a table name IS a quoted string (`Table: "skills"`), so a
   re-introduced `MinRows: 40000` returned GREEN. The subject is now matched on the raw line and the
   magnitude on the blanked one. **A fence that misses its own motivating case is worse than no fence,
   because it is also a claim that the case is absent** — which is why RED-proving is not optional.
