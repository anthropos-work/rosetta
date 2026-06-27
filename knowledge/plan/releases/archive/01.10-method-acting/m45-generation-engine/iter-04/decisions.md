# iter-04 decisions

| ID | Decision | Rationale |
|----|----------|-----------|
| D1 | `BatchHash` sorts the member prompts before hashing → the batch dir is INDEPENDENT of member order. | Two descriptors that expand to the same set of prompts in a different order should share cache entries. Per-member FILE selection still uses the index, so members stay addressable. |
| D2 | `Put` REJECTS non-JSON content (returns an error, writes nothing). | The cache must hold only valid envelopes; a malformed LLM response is re-rolled by cmd/gen-batch, never cached. This keeps the $0-reseed invariant honest (the cache never serves a half-formed member). |
| D3 | The lock is OS-atomic `O_CREATE\|O_EXCL` file creation (presence == held), with a `BreakLock` recovery. | Portable, no extra dep, atomic on the local FS the cache targets. A dead run's stale lock is recoverable via an explicit break (cmd/gen-batch will gate it behind --break-lock) rather than a hand-delete. |
