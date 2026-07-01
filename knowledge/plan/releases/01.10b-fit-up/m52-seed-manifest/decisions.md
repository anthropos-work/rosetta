# M52 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## D1 — Prompt extraction mechanism: `go:embed`, byte-identical
- **Context:** S1 must make the generation mother prompt file-resident (auditable, no-recompile-to-read)
  WITHOUT busting the M45 prompt-hash cache (keyed on the exact rendered mother prompt).
- **Options:** (a) `go:embed` the exact const bytes into a checked-in `.tmpl` file; (b) load the template
  from a YAML/JSON at runtime (os.ReadFile); (c) deliberately re-word + re-key the cache.
- **Choice:** (a) — `//go:embed prompts/default_batch_prompt.tmpl` into the same-named var, bytes extracted
  from the former const via a one-shot Go writer so they are byte-identical.
- **Why:** (a) keeps the cache key UNCHANGED (existing caches stay valid, $0 reseed holds) — the exact
  cache-integrity invariant the overview's Risk flags — while still making the prompt a plain file an
  auditor reads without Go. (b) adds a runtime read + a failure mode (missing file at gen time); embed
  fails at BUILD instead (safer). (c) needlessly invalidates a valid cache. Same-name var means zero
  consumer/test churn. `embed` is stdlib → no supply-chain change.
