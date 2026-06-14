# M21 iter-04 — decisions (iter-local)

_(Cross-iter: M21-D8 in the milestone-root `decisions.md`.)_

- **iter-04-L1 — Harness retained for iter-05.** The throwaway harness (`m21test-pg`, 26 tables + 9 row-tables, digest
  `6cd35278…`) is kept alive into iter-05 (same session) since stages 5–6 (boot + serve) start exactly from this
  stage-4 state. Torn down at session end.
- **iter-04-L2 — structure.sql is a validation artifact, not the deliverable.** It proves the Option-A approach; the
  real deliverable is the stacksnap capture-side extension that produces an equivalent artifact over `--dsn` and
  applies it in provision (routed `STRUCT-M21-codeify`).
