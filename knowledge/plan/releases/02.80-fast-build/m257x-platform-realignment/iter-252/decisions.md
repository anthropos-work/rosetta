# iter-252 decisions

## `D-M257x-252-1` — the class boundary is the SUBJECT, not the mechanism

All four guards `ROUTE-M257x-250` named call `exists()` / `is_dir()` on the operator's tree. Three were
defective and one is correct, and the mechanism is identical in all four. The difference is **what claim
is being made**:

| guard | claim | operator's tree is… | verdict |
|---|---|---|---|
| `rext_path_guard` | *this corpus citation names a real tool* | a place it reached into for an answer | **defect** (iter-250, 1) |
| `corpus_citation_guard` | *this corpus citation resolves* | same | **defect** (iter-251, 21) |
| `fence_command_guard` | *this fenced command names a real target* | same — but it **refuses** rather than accusing (44 `workspace not provisioned` buckets) | reach gap, not this class |
| `clone_drift_guard` | *the clones have not advanced past what the corpus cites* | **the subject itself** | **correct** — refuses, exit 2 |

**Decision.** The rule to carry forward is not *"never touch the filesystem"*. It is: **a guard whose
claim is about the corpus must not let the operator's tree decide that claim** — and a guard whose claim is
about the operator's tree must **refuse** when it is absent, never substitute a corpus verdict. Same
mechanism, opposite obligations, decided by the subject.

## `D-M257x-252-2` — a route is closed by measuring its last member, never by analogy from the first three

Three confirmed instances in three consecutive iters made a fourth feel certain, and the pre-registration
was deliberately set **against** that thesis: 3 of the 4 claims predicted the class does not extend. All
four held. One command settled what a repair-by-analogy would have cost an iter to build and then have to
retract.

## `D-M257x-252-3` — when the measurement finds nothing to repair, the deliverable is the PIN

`clone_drift_guard` behaves correctly today, and nothing was holding it there. Two regression tests now
do: an absent clone set must exit **2** with `CANNOT RUN`, and the refusal must **name where it looked**.
Without them this iter's product would be a memory, and the next person re-derives it — the
"re-derived from scratch each time" waste this milestone exists to end (`overview.md`).

**Both fixture bugs these tests surfaced were mine**, caught before commit: `main()` takes argv **without**
the program name, and `clones()` returns a mapping rather than a list. A test written against an API you
did not read is a test of your assumptions.
