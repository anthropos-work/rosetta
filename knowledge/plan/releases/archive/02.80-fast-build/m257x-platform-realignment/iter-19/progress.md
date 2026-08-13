---
milestone: M257x
iter: 19
---

# iter-19 — progress

**Type:** tik (under `TOK-01`) — a measurement iter, single-step by design.

## What happened

**Clause 2 re-measured on the stack iter-18 proved green, and the result is a clean falsification of the
milestone's own most recent attribution.**

    ./run-playthroughs.sh 1 --reset          (full suite; the ptreport gate is binding only on a full run)
    invocation vantage: stack-demo/rosetta-extensions/playthroughs/e2e, rext @ fast-build-m257x-iter-18,
    stack demo-1 cold from iter-18 cycle 3, autoverify green:true / warnings:0 @ 2026-08-01T17:10:44Z,
    anon GET /items/task_sub_checks = 200

    summary: passing 20 · failing 10 · unimplemented 1 · unimplementable 0     total 31

**Identical to iter-15 — and not just in the totals.** The sorted list of the ten failing ids `diff`s
clean against iter-15's:

    assignment-monitoring.assign-and-track.UC1 · .UC2 · hiring.recruiter-comparison.UC1 ·
    onboarding.enterprise-hiring.UC1 · org-admin.roles.UC1 · skill-paths.legacy.UC1 ·
    skill-paths.save-for-later.UC1 · workforce-intelligence.organization-feedback.UC1 ·
    .skills-funnel.UC1 · .talent-pool.UC1

So `FIX-M257x-iter15-directus-versions-403` and `FIX-M257x-iter15-library-category-expansion` are **not**
downstream of the unserved content layer (`D-M257x-19-1`). iter-18's close attached that caveat to them
explicitly; it is withdrawn, and the next tik can work them directly instead of re-deriving whether they
still exist.

**Two things came out that were not the question asked.**

1. **iter-15's number is now REPRODUCED, not merely re-asserted** (`D-M257x-19-2`). Its own re-run had
   scored `17/31` and the difference was *explained* by the missing `--reset` rather than tested. A real
   `--reset` on a differently-built stack lands on `20/10/1` with the same ten ids, so the explanation has a
   confirming observation and the reset-vs-additive discipline is load-bearing rather than merely stated.

2. **The largest cause is at least TWO fields and the count moved** (`D-M257x-19-3`). Measured from the
   `backend` container's own log on the green stack:

       119 × cannot unmarshal string into … JobSimulation.data.library_category of type struct
        11 × cannot unmarshal string into … JobSimulation.data.job_position    of type simulation.…
        58 × directus_versions

   iter-15 named `library_category` alone at **106**. Same shape — `app` reads an expanded relation,
   Directus returns the raw id — but a fix aimed at one field name would leave the other standing. Whether
   they share one root cause is **not concluded**; it is likely, which is precisely why it should be
   measured.

**And the instrument needed a host fix to run at all** (`D-M257x-19-4`): `run-playthroughs.sh:118` calls a
bare `stackseed`, which is not on PATH here — the bring-up builds it into the stack's own
`stacks/demo-<N>/bin/`. The runner already derives its ports, bases and seed path from `N`; the binary
directory is derivable from the same `N`. Routed, not fixed — changing the instrument during the measurement
it is taking is the mistake this iter exists to avoid.

## Gate movement

**None, and that is the finding.** Clause 2 stands at `20 / 10 / 1`, NOT MET. **2 of 5 clauses hold**
(1 and 4).

## Verification

No source change landed, so no suite moved. The measurement's own vantage is recorded above (§5 rule 12:
*say which invocation produced the number*), and the comparison against iter-15 is a `diff` of sorted id
lists rather than a reading of two summary lines.

## Close — 2026-08-01

**Outcome:** Clause 2 re-measured on a stack whose content layer actually serves: **`20 / 10 / 1`,
byte-identical failing set to iter-15**. The "downstream of the Directus defect" attribution iter-18 attached
to two routed causes is **refuted**; iter-15's number is **reproduced**; and the largest cause turns out to
span **two** fields (119 `library_category` + 11 `job_position`), not the one it was routed under.
**Type:** tik
**Status:** closed-no-lift (documented falsification — the planned investigation completed and its declared
acceptable outcome, *"an unchanged 20/10/1 is a complete outcome"*, is the one that occurred)
**Gate:** NOT MET (2 of 5 — unchanged by this iter)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (1st no-prog tik of a streak; iter-18 moved the
metric) — (3) re-scope: n (platform origin HEAD `2adcf71` unchanged — occurrence stays 1 of 2) —
(4) user-blocker: n — (5) cap-reached: n (2 tiks this session, cap is 5) — (6) protocol-stop: n —
Outcome: continue
**Decisions:** D-M257x-19-1 … D-M257x-19-4.
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter15-library-category-expansion` **RE-SCOPED** → next tik. **Two** fields, not one:
  `JobSimulation.data.library_category` (119) and `JobSimulation.data.job_position` (11). Measure whether
  they share a root cause before fixing either.
- `FIX-M257x-iter15-directus-versions-403` → next tik, caveat **withdrawn** (independent of the Directus
  serving defect; 58 occurrences on the green stack).
- `FIX-M257x-iter19-playthrough-runner-path` → later tik. `run-playthroughs.sh:118` calls a bare
  `stackseed`; derive the stack's own `bin/` from `N` the way every other value in that script is derived.

**Lessons:**
- **Re-measuring an inherited cause is cheap next to working one that has moved.** The whole iter cost one
  suite run and it decided the shape of the next two. The alternative — assuming either direction — was a
  coin flip on a 100-plus-occurrence class.
- **Compare the SET, not the summary.** `20/10/1` twice could have been ten different failures; the
  `diff` of sorted ids is what makes "identical" a measurement.
