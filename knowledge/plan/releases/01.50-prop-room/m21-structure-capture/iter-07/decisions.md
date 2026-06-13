# M21 iter-07 — decisions (iter-local)

_(Cross-iter: M21-D11 in the milestone-root `decisions.md`.)_

- **iter-07-L1 — auto-provision is integrated INTO replay (not a separate subcommand).** The gate reads "stacksnap
  applies the captured structure → replay exits 0"; one `stacksnap replay` that auto-provisions-then-loads satisfies
  it + fixes the exit-5 the provision recipe hit. Gated strictly on cache-miss + a structure-bearing snapshot, so the
  existing taxonomy/reference + already-provisioned paths are untouched.
- **iter-07-L2 — StructureCapturer is a type-asserted OPTIONAL interface, not a Capturer method.** This keeps every
  existing capture fake (which never captures structure) compiling unchanged; capture.Run errors loudly only if a
  surface sets CapturesStructure but the capturer doesn't implement it.
- **iter-07-L3 — pg.ExecScript uses the simple protocol.** pgx's default extended protocol rejects multi-statement
  strings; the structure artifact is fixed DDL with no bind params, so the simple-protocol PgConn().Exec batch is the
  right tool. (A diverged/partially-provisioned target fails CREATE TABLE → the batch aborts → exit reported, no
  silent half-provision.)
- **iter-07-L4 — the apply/serve split.** iter-07 automates the SCHEMA (tables+PKs → digest converges → rows replay,
  stages 3-4). The SERVE rows (directus_collections registration + public read permissions, + the firewall
  structural-metadata admissibility class) are iter-08 → flips the gate met (stages 5-6 automated).
