# M42e Progress

## Running ledger
Iter closeouts append here (one line each — the tik/tok, what the sweep measured, what was fixed, the
remaining failing-page + escape count vs the gate).

<!-- iter-NN/ dirs are created by /developer-kit:build-mstone-iters on first run. -->

- iter-01 (tok/bootstrap): authored `coverage-protocol.md` + scaffolded the Playwright coverage harness under `stack-verify/e2e/` (compile-validated); resolved the 4 open questions; KB-fidelity GREEN; TOK-01 recorded — see iter-01/progress.md
- iter-02 (tik): baseline sweep -- `(failing=44, escapes=1)` over reachable=80 (employee/`maya-thriving`); all 44 failures probe-confirmed as a `networkidle` harness flake (pages are http=200 w/ real content), routed to iter-03 (tooling-iter); 1 real content escape + maxPages saturation routed fwd -- see iter-02/progress.md
- iter-03 (tik, tooling): wait-strategy fix (`networkidle`->`domcontentloaded` + bounded settle + inline screenshots + seed-vs-discovered scoring); re-sweep `(failing=8, escapes=1)` down from `(44,1)` -- 36 false failures removed; TRUE residual (5 sim-result empties, 2 empty skill-paths, 1 `/` sentinel false-positive, 1 escape) routed to iter-04; lesson folded into coverage-protocol.md -- see iter-03/progress.md
