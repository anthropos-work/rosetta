# iter-05 decisions

| ID | Decision | Rationale |
|----|----------|-----------|
| D1 | `--max-cost` is MANDATORY (required + must be > 0); a non-dry-run with no/zero ceiling errors out. | The user's hard requirement: no batch ever runs uncapped. Failing closed (error, not a default) makes the budget guard impossible to forget. --dry-run is exempt (it never spends). |
| D2 | A member that stays malformed (or un-de-collidable) past the re-roll budget is DROPPED, not an error. | The CODE-vs-AI boundary: a bad generation yields a SHALLOWER batch (one fewer generated member), never an invalid one or a failed run. The closure gate stays green; the run completes. |
| D3 | Per-run state (cache/completion/tracker) is threaded through a `generator` struct, NOT package globals. | The concurrent workers need shared access to the tracker + cache; a struct keeps it test-safe (each `run` is isolated) — important because gen-batch is unit-tested end-to-end with fixtures. |
| D4 | The ceiling guard pre-checks `WouldExceed` BEFORE launching each call (over-estimating the completion side). | A breach must be PREVENTED, not detected after spending. Pre-checking with an over-estimate errs toward caution; the serialized check + break-on-first makes the abort deterministic under --max-concurrent. |
