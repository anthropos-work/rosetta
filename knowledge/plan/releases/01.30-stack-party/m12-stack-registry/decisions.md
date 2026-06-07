# M12 — Decisions

_Implementation decisions with rationale. ID scheme: M12-D1, M12-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| KB-1 | GUIDE.md ~L39 claims "the registry assigns N" but current code only *records* a caller-passed N. Non-load-bearing for the M12 implementer (M12 *builds* the allocator that makes this true). Reconcile the prose at Phase 5 once the allocator lands. | From Phase 0b KB-fidelity (YELLOW). The claim converges to truth at close — fixing it pre-build would make it true-but-unimplemented. | 2026-06-07 |
| M12-D1 | The registry is a **single shared module** `stack-core/stack_registry.py` (not an extension of demo's `registry.json`). N is keyed by docker project `"<type>-<N>"` but allocated from **one shared N-pool across both types** — `dev-1` and `demo-1` can never coexist. Runtime file at `stack-core/.stacks/registry.json` (gitignored), the identical path from both CLIs (both reach it as `../stack-core/stack_registry.py`). | The collision bug is that `dev-N`/`demo-N` both compute `base+N*OFFSET`. Making N a shared pool is the minimal fix; co-locating the runtime file with the module guarantees both CLIs share one registry. The demo `registry.json` stays as-is (M14 wires the skills); this is the new cross-type source of truth. | 2026-06-07 |
| M12-D2 (Q1+Q3) | **Registry-of-record + `docker ps` reconcile that only ADDS used-N, never subtracts.** A reserved row keeps its N until explicit `release()`; `docker ps` adopts unmanaged live stacks (so a manually-started stack reserves its N) but a lagging/empty `docker ps` never frees a reserved slot. | Resolves Q1 (registry is the record) and Q3 (manual stacks adopted). The "never subtract" rule is the race guard: a stack reserved-but-not-yet-started must not lose its N to a racing `up` the instant `docker ps` lags. Teardown (`release`) is the only free path → deterministic. | 2026-06-07 |
| M12-D3 (Q2) | **Concurrency via `fcntl.flock(LOCK_EX)` on a sidecar `.lock`** around the whole read-reconcile-pick-write in `allocate()`; portable macOS+Linux (unlike `flock(1)`). Registry writes are **atomic** (temp + `os.replace`). | Two concurrent `up`s must never pick the same N. flock on a persistent sidecar matches the existing demo `reg_set`/`reg_del` pattern (consistency). Atomic write + corrupt-JSON recovery in `_load` make a crash mid-write non-wedging. | 2026-06-07 |

## Open at design — RESOLVED during build
- ~~M12-Q1: registry-of-record vs docker-ps-derived~~ → **M12-D2**: registry is the record; docker-ps reconciles (adds-only).
- ~~M12-Q2: lock mechanism for concurrent up's~~ → **M12-D3**: fcntl.flock on a sidecar + atomic write.
- ~~M12-Q3: reconciling manually-started stacks~~ → **M12-D2**: live containers with no registry row are adopted (reserve their N).
