# M253 — Progress

Iterative milestone (perf, measure→patch→re-measure). Primary metric: first-meaningful-paint < 1000 ms + no blank > 1 s,
p95 over 5 consecutive cold loads on a cold demo (state the environment — laptop vs tailnet), gated on a fresh-green
`autoverify.json`.

## Running ledger

- iter-01 (tok/bootstrap): authored TOK-01 (shell-before-awaits + no-thirdparty demopatches + FCP runner); baseline skeleton-visible 4669 ms (demo-2, laptop); dominant await = canAccess (~3.9 s), not clerk.load — see iter-01/progress.md
