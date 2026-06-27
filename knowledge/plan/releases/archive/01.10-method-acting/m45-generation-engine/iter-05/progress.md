**Type:** tik

# iter-05 — cmd/gen-batch (the generation CLI)

Built component (4) of TOK-01's chain: the generation CLI that orchestrates expansion → cache → LLM with
the mandatory cost ceiling.

## Work done
- **NEW `cmd/gen-batch/main.go`** — the CLI: `--seed` (required), `--max-cost` (MANDATORY, >0, the run
  aborts BEFORE a call would breach it; --dry-run exempt), `--model` (gpt-4o-mini default),
  `--max-concurrent` (5), `--capture-version` (the cache-key extension), `--cache-root`, `--dry-run` (the
  offline preview — expands + reports the plan with NO LLM call, NO key), `--break-lock`, `--max-rerolls`.
  Flow: load+validate blueprint → `EffectiveBatches(reservedHeroNames)` → `batchcache.Open` → count cache
  hits → (dry-run? report+exit) → acquire the `.lock` → build the `services/ai` client values-blind (only
  if there's something to generate) → `generator.run` (the `--max-concurrent` semaphore; ceiling
  pre-check before each launch; re-roll on malformed JSON; post-gen hero-collision re-roll; atomic cache
  Put; member DROP after the budget = shallower batch, never an error) → write manifest → cost report +
  valid-JSON rate.
- The `generator` struct threads cache/completion/cost-tracker through the workers (no package globals;
  test-safe). `ValidJSONRate()` exposes the gate's pre-re-roll valid-JSON rate. `reservedHeroNames`
  collects the curated heroes (legacy personas + all stories' heroes) for the collision check.
- **NEW `cmd/gen-batch/main_test.go`** — 10 tests, ALL fixture-driven (no key, no network): mandatory
  `--max-cost`, `--seed` required, dry-run-makes-no-LLM-call (a panicking fixture), generate-then-$0-
  cached-reseed (run 2 makes ZERO calls), re-roll-on-malformed, drop-after-budget (not cached, not an
  error), hero-collision re-roll (the hero name never caches), ceiling abort (stops short of all 50),
  reserved-name collection, lock fence (a second run on a locked batch refuses).

## Measurements
- `go build ./...` + `go vet ./...` clean; `gofmt -l cmd/gen-batch/` clean.
- `go test ./cmd/gen-batch/...` GREEN; full `go test ./...` GREEN. Test funcs 613 -> **623** (+10).

## Close — 2026-06-26

**Outcome:** Component (4) landed — the generation CLI with the mandatory `--max-cost` ceiling, the
`--max-concurrent` semaphore, re-roll-on-malformed, hero-collision re-roll, the $0 cache reseed, dry-run,
and the lock fence — ALL fixture-proven (no key). The gate's measurement harness (valid-JSON rate / cost /
collision) is now computable. Gate still 0/5 EMPIRICALLY (the real valid-JSON rate needs a REAL run — the
seeder + the gate-proving batch are the next call's work), but every gate METRIC is now measurable.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (0/5 empirically — engine orchestration complete + fixture-proven; the GeneratedBatchSeeder + the REAL capped batch + the live $0-reseed remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (this is tik #4 of the session — iter-02..05; the cap is 5 tiks, not reached) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (--max-cost mandatory+>0), D2 (drop-not-error on malformed-past-budget), D3 (generator struct, no globals)
**Side-deliverables (if any):** none
**Routes carried forward:** iter-06 (THIS call, tik #5 → the cap fires after it) — component (5): the `GeneratedBatchSeeder` (surface 'generated-batch', DependsOn users+taxonomy, PerStackIsolated) reading the cache → users/persona/profile rows via the existing resolvers (drop-not-fabricate, closure green). Then a LATER call: the REAL gate-proving N=20 capped batch + the live $0 reseed on demo-3.
**Lessons:** threading per-run state through a `generator` struct (not package globals) keeps the
concurrent workers test-safe — a CLI that's a one-shot is still cleaner without globals. The ceiling guard
must pre-check BEFORE launching each call (not after) so a breach is PREVENTED, not merely detected; the
serialized `WouldExceed` + the break-on-first make the abort deterministic even under --max-concurrent.
