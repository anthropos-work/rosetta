# Hardening Ledger — M45 Generation engine

The M45 generation-engine CODE + tests live in the rext **authoring** copy
(`.agentspace/rosetta-extensions/stack-seeding/`); this ledger lives in the corpus on the
`m45/generation-engine` branch. The harden test commits + the tag bump are in rext (its own git);
this ledger records the passes. Final-mode (cumulative scope across the whole engine).

Footprint (all `stack-seeding/`): `services/ai/` (Azure multi-name EU-first routing + cost tracker +
`--max-cost` ceiling + per-call `--call-timeout` + values-blind redaction), `blueprint.Batch` +
`EffectiveBatches()`, `batchcache/`, `cmd/gen-batch`, `seeders/GeneratedBatchSeeder` — plus the 3
gate-proving fixes (per-call timeout / intra-batch name dedup / `user_skills` CHECK 23514).

All tests are DETERMINISTIC (fixture LLM responses, never real Azure calls — the gate is already
proven on a real Azure gpt-4o-mini run + demo-3 proof). rext tag at start:
`method-acting-m45-iter07-gate`; final tag: `method-acting-m45-harden-final`.

## Pass 1 — 2026-06-26 — final

**Iters hardened this pass:** all milestone-touched code (cumulative final scope)
**Tiks covered since prior pass:** all iters in milestone (first harden pass)
**Coverage delta on touched files:**
- services/ai: 83.5% -> 93.5% stmts
- cmd/gen-batch: 84.9% -> 89.2% stmts
**Tests added:**
- services/ai/ai_test.go: the EU->direct 429 fallback path (was 0%) via a fake `libai.AI` —
  primary-success / 429-recovers / non-429-propagates / 429-no-fallback; the transport-boundary
  values-blind key-leak grep (a provider error echoing the env key -> `[REDACTED]`); Spent +
  pct(0-ceiling). (5 unit + 2 accessor)
- cmd/gen-batch/main_test.go: the `--call-timeout` STALL-CLASS regression (a hung endpoint fails
  fast in ~100ms, never blocks the batch — the THE bug); `--break-lock` stale-lock recovery;
  zero-count no-op; avoidNamesHint determinism+cap; ValidJSONRate vacuous case. (7)
- seeders/generated_batch_test.go: the `user_skills` CHECK 23514 FK-provenance regression (every
  claimed-skill row ties to its member's ONE current-role experience; education edge stays NULL);
  the FK write-order (users -> experiences -> user_skills); current-role to=NULL. (3)
**Bugs surfaced + fixed inline:** none
**Flakes stabilized:** none
**Stop condition:** continue-to-next-pass — batchcache error paths + drop-not-fabricate deepening + cross-iter integration still uncovered

## Pass 2 — 2026-06-26 — final

**Iters hardened this pass:** all milestone-touched code
**Coverage delta on touched files:**
- batchcache: 84.6% -> 88.5% stmts
- cmd/gen-batch: 89.2% -> 90.9% stmts
- seeders: 96.9% -> 97.0% stmts
**Tests added:**
- batchcache/cache_test.go: the failure branches — Get on a missing member; Put / WriteManifest
  surfacing a write error (read-only dir); the lock holder-pid record; double-Unlock no-op; Open
  idempotency; a fresh Put overwriting a crashed `.tmp`. (7)
- seeders: drop-not-fabricate HARDENED — on a hallucination-heavy envelope every written
  role/skill node-id is a REAL one from the conn's pools, never a fabricated `J-…`/`K-…` (the
  highest-value invariant); email-local sanitization (invalid chars stripped, domain never
  LLM-injectable); reservedHeroNamesForSeed personas+stories de-dup. (3)
- cmd/gen-batch: reservedHeroNames stories-heroes path; the FULL-PIPELINE cross-iter integration —
  blueprint expansion -> generate+cache -> a second run reseeds at $0 with ZERO LLM calls -> a
  `--capture-version` bump invalidates and regenerates. Pins the cache-key × generator ×
  cost-tracker interaction no single iter test spanned. (2)
**Bugs surfaced + fixed inline:** none
**Flakes stabilized:** none
**Stop condition:** continue-to-next-pass — COPY-fault error paths, isolation assert, and the envelope boundary-fuzz dimension still uncovered

## Pass 3 — 2026-06-26 — final

**Iters hardened this pass:** all milestone-touched code
**Coverage delta on touched files:**
- cmd/gen-batch: 90.9% -> 93.0% stmts
- seeders: 97.0% -> 97.3% stmts
**Tests added:**
- seeders: a COPY fault on any generated-batch table (users / user_experiences / user_skills)
  propagates wrapped with the seeder name + the failing table (operator-debuggability); the writes
  are ALL `PerStackIsolated` to the public schema — never a shared/prod write (every audit entry
  Allowed + scoped to the stack); an adversarial-envelope BOUNDARY FUZZ — empty/blank-name/array
  envelopes + on-disk raw corruption all DROP cleanly (no panic, no error), the well-formed members
  seed, and NO fabricated node-id leaks through. (3)
- cmd/gen-batch: the values-blind no-key error (no key material in the message); a fully-cached run
  needs NO key (the client is only built when there's something to generate); `--max-concurrent 0`
  is clamped to 1, never deadlocks. (3)
**Bugs surfaced + fixed inline:** none
**Flakes stabilized:** none
**Stop condition:** continue-to-next-pass — fixture accessors, bias-determinism, and the genEmail fallback edge still uncovered

## Pass 4 — 2026-06-26 — final

**Iters hardened this pass:** all milestone-touched code
**Coverage delta on touched files:**
- services/ai: 93.5% -> 97.8% stmts
- blueprint: 98.2% -> 98.9% stmts
**Tests added:**
- services/ai: the fixture accessors (custom ModelID through Model(), empty LastUser before any
  call). (2)
- blueprint: biasForIndex DETERMINISM (same index -> same bias, the $0-reseed invariant) + the
  all-zero-weights -> calibrated guard + a non-canonical bias key. (3)
- seeders: the genEmail ultimate-fallback regression. (1)
**Bugs surfaced + fixed inline:**
- genEmail degenerate-address edge: a separator-only email local part (an empty first+last yields
  ".") produced `.@domain`. `sanitizeEmailLocal` now collapses a separator-only result to "" so
  genEmail falls through to the name then to "member" — a valid address is ALWAYS produced (the
  CODE-owns-identity guarantee). Fix + regression land in the same commit (rext d00c606). Defensive
  reach (upstream `env.Name == ""` drop already prevents the empty-name path in production), <10
  lines, single subsystem — Fate 1.
**Flakes stabilized:** none
**Stop condition:** continue-to-next-pass — confirm the coverage delta has flattened (one more measuring pass)

## Pass 5 — 2026-06-26 — final

**Iters hardened this pass:** all milestone-touched code (confirmation pass)
**Coverage delta on touched files:** 0% across all 5 packages (services/ai 97.8%, blueprint 98.9%,
batchcache 88.5%, cmd/gen-batch 93.0%, seeders 97.3% — identical to Pass 4)
**Tests added:** none — the dimension scan (depth / edge / error-path / regression / fuzz /
perf) is exhausted for all REACHABLE surfaces. The residual uncovered are low-value, hard-to-reach
defensive branches: `main()` (the untestable OS entrypoint, 0% by design); the I/O-fault rename-fail
branches of Open/Put/WriteManifest (reachable only via a fragile filesystem race — the read-only-dir
tests already cover the write-fail arm); `NewFromEnv`'s `openai.New`/`NewAzure` constructor-error
branch (fires only on a lib-rejected malformed key — not values-blind-testable without fabricating a
bad key). Coverage-as-a-finder, not a goal: no shallow tests written to bump the number.
**Bugs surfaced + fixed inline:** none
**Flakes stabilized:** none — `go test -race ./...` clean across the whole suite; the 3x flake gate
(incl. the timing-sensitive `--call-timeout` test) is clean.
**Knowledge backfill:** none required — the engine's behaviors are documented in
`corpus/ops/demo/ai-generation-spec.md` + `cache-spec.md` (the kb-fidelity audit confirmed alignment
pre-milestone); hardening surfaced no new protocol-level or subsystem truth beyond the genEmail
defensive-edge fix (captured in the code comment + this ledger).
**Audits:** closure GREEN (seed-side: 8 no-fabrication assertions incl the adversarial-fuzz no-leak
check; runtime `datadna measure-closure --stack demo-3` was `[PASS]` on the real gate run) ·
isolation PerStackIsolated/public-only (asserted, zero shared/prod writes) · supply-chain = exactly
1 new dep (`ai v1.40.1`, the deliberate sanctioned M45 inflection; harden added ZERO further deps,
`go mod tidy` is a no-op) · alignment N/A (zero clerkenstein change -> carries 100%) · values-blind
(the transport-boundary + no-key key-leak greps).
**Stop condition:** stabilized
