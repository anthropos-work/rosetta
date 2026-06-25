# M42e Progress

## Running ledger
Iter closeouts append here (one line each — the tik/tok, what the sweep measured, what was fixed, the
remaining failing-page + escape count vs the gate).

<!-- iter-NN/ dirs are created by /developer-kit:build-mstone-iters on first run. -->

- iter-01 (tok/bootstrap): authored `coverage-protocol.md` + scaffolded the Playwright coverage harness under `stack-verify/e2e/` (compile-validated); resolved the 4 open questions; KB-fidelity GREEN; TOK-01 recorded — see iter-01/progress.md
- iter-02 (tik): baseline sweep -- `(failing=44, escapes=1)` over reachable=80 (employee/`maya-thriving`); all 44 failures probe-confirmed as a `networkidle` harness flake (pages are http=200 w/ real content), routed to iter-03 (tooling-iter); 1 real content escape + maxPages saturation routed fwd -- see iter-02/progress.md
- iter-03 (tik, tooling): wait-strategy fix (`networkidle`->`domcontentloaded` + bounded settle + inline screenshots + seed-vs-discovered scoring); re-sweep `(failing=8, escapes=1)` down from `(44,1)` -- 36 false failures removed; TRUE residual (5 sim-result empties, 2 empty skill-paths, 1 `/` sentinel false-positive, 1 escape) routed to iter-04; lesson folded into coverage-protocol.md -- see iter-03/progress.md
- iter-04 (tik, closed-no-lift): investigated the 5 sim-result empties; FALSIFIED the planned seed -- the result is a runtime-computed AI-evaluation artifact (no `jobSimulationResult` for a backdated seeded session), not a seedable row; zero-edit line blocks the platform path. Correct fix = crawl-scope (exclude per-session result deep-links), routed to iter-05 (tooling-iter). Metric unchanged `(8,1)` -- see iter-04/progress.md
- iter-05 (tik, tooling): result-deep-link `skipPaths` rule (crawl-scope the runtime-computed sim-result pages); re-sweep `(failing=3, escapes=2)` down from `(8,1)` -- cleared all 5 sim-result empties; escapes 1->2 is a deeper-crawl discovery (2nd external chapter link). Residual: 1 `/` sentinel FP + 2 empty skill-paths + 2 external chapter links -- see iter-05/progress.md
