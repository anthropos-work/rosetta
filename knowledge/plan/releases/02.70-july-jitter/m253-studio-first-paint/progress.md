# M253 — Progress

Iterative milestone (perf, measure→patch→re-measure). Primary metric: first-meaningful-paint < 1000 ms + no blank > 1 s,
p95 over 5 consecutive cold loads on a cold demo (state the environment — laptop vs tailnet), gated on a fresh-green
`autoverify.json`.

## Running ledger

_(iter-NN/ directories are created by `/developer-kit:build-mstone-iters` on its first invocation — no iter dirs exist at scaffold time. Per-iter entries accumulate here during the iter loop.)_
