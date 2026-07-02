# iter-05 decisions

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| D1 | Run the gate with the demo-1 consumption-clone's `bin/` on PATH (`stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/bin`) so `run-playthroughs.sh --reset` finds `stackseed`. | The runner calls bare `stackseed` (the pinned tooling the demo consumes); it is not on the login PATH. Prepending the demo-1 bin is the correct, zero-code-change way to reach the pinned `stackseed`/`stacksnap` binaries. Recorded as a gate-run environment prereq for the next runner (harden/close, and any future 5-run re-verification). NOT a runner code change — the runner correctly delegates to the pinned CLI. | 2026-07-02 |
