# iter-03 decisions

| ID | Decision | Rationale |
|----|----------|-----------|
| D1 | Use `text/template` (stdlib) for the mother-prompt expansion — NO new templating dep. | The 0-new-deps discipline (M45 already spends its ONE sanctioned new dep on `ai`); stdlib text/template is sufficient for the per-member prompt rendering. |
| D2 | Bias distribution is an INTERLEAVED largest-remainder assignment, not a block-run weighted sequence. | A block run ([50 under][50 over]) bunches one bias at the start of a small batch span (a 3-member batch would be all-under). Interleaving spreads the declared mix evenly across the index space, so even a tiny gate-proving batch (N=20) shows the distribution. Still deterministic → byte-identical reseed. |
| D3 | The taxonomy capture version is NOT folded into the prompt TEXT (prompts are taxonomy-agnostic) — only into the cache KEY (by cmd/gen-batch). | The mother prompt asks for plausible skill NAMES; it doesn't reference the taxonomy. Folding the capture version into the key (not the text) means a taxonomy re-replay invalidates the cache without changing what the LLM is asked — the cache-spec coarse-invalidation contract. |
