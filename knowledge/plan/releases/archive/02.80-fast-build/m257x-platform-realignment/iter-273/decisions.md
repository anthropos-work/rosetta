# iter-273 — decisions

## D-M257x-273-1 — a SCOPED suite run cannot grade clause 2, and the harness says so itself

**Context.** Gate clause 2 reads *30 live / 0 failing / 0 error* over the **full** Playthrough suite. Every
recent grading of it — including iter-272's — came from a `--grep`'d run.

**The harness prints its own disqualification on every scoped invocation:**

```
ℹ this run was SCOPED — its artifacts are advisory. The last BINDING run is preserved at
  report/last-binding-run.json
⚠ ptreport gate not met — ADVISORY on a SCOPED run … Re-run unscoped for a binding verdict.
```

That text is not a warning about noise; it is a statement that the artifact **is not gradeable**, because
every un-selected id correctly reports *"did not run"* and the gate cannot distinguish that from a pass.

**Decision.** Clause 2 is graded only from an **unscoped** run. Measured here for the first time at the
shipping pin: **31 total / 29 passing / 1 failing / 1 unimplemented / 0 unimplementable / 0 error** →
**30 live, 1 failing**. The one-failure prior is now measured rather than inherited, and its identity
(`talent-pool.UC1`) is the same across three runs.

**Why it is worth a decision entry rather than a footnote.** This is the third time in four iters that a
belief steered work while resting on an artifact that could not support it — after *"empty projection
tables"* (refuted, iter-272) and *"the trace carries the response"* (refuted, iter-272). The common shape
is **an artifact read past its own declared scope**. The harness was the only one of the three that
announced the limit in its own output, and it was still read past.
