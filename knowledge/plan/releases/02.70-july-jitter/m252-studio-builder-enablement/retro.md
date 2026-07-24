# M252 — Retro (studio-desk builder enablement)

## Summary
Fixed the demo studio builders' `POST /api/ai/completion` 500 — root cause a demo-**wiring** gap, not a missing
secret: base-compose studio-desk inherits ONLY `platform/.env` (no AI keys), while its provisioned
`AI_OPENAI_API_KEY` / `AI_ANTHROPIC_API_KEY` live in the studio-desk clone's own `.env`. The fix is a single
existence-guarded **`env_file`** at `<clone>/studio-desk/.env` in `gen_injected_override.py` — **`env_file`-ONLY,
NO `MOCK_CLERK`, NO provider-chain pin**. Added a demo-aware, non-fatal, values-blind **autoverify (g)** container
provider-key assert, and studio-desk's **first-ever Playthroughs-manifest entry** (Product "Studio", 2
builder-GENERATE Playthroughs, `16 → 18 live / 0 TODO`). Talk-to-data (M239 Bedrock) re-confirmed COMPLETE.
**Wiring PROVEN live** (op1, demo-2). Code-of-record `july-jitter-m252-studio-builder @ d80db9f` (on origin,
rung-zero verified); deferral audit GREEN; 0 platform-repo edits.

## Incidents This Cycle
- **P2 — a raw `curl` 302 misdiagnosed the demo studio as "unreachable" (D1, self-corrected in-milestone).** An
  early unauthenticated `curl :29000` returned `302 → /login`, first read as "studio entirely unreachable → needs
  `MOCK_CLERK=true` to disarm the gate." WRONG: the demo studio is **deliberately, test-enforcedly**
  Clerkenstein-authenticated — a logged-in org-admin hero 302s through the fake-FAPI handshake and passes
  `checkEnterpriseAndAdmin`; the `curl` 302 was the prod `clerkMiddleware()` behaving exactly as designed for a
  session-less browser. Adding `MOCK_CLERK` would have regressed the demo to the legacy bypass and **failed 2
  pinned regression tests**. Lane A correctly refused it; the fix stayed `env_file`-only. Caught before any code
  landed — the pinned tests were the backstop.
- **P2 — the close found the rosetta docs-commit MESSAGE imprecise (D6).** The `cedde09` commit summary (and the
  orchestrator's close-mandate prose) describe the fix as "env_file + a MOCK_CLERK/AI_PROVIDER_CHAIN overlay."
  The **shipped rext code-of-record is `env_file`-ONLY** — the Lane-A commit `4486cdd` says so verbatim ("No
  MOCK_CLERK, no provider-chain pin"), and all four corpus docs + D1 agree. The live boot line
  `ProviderHealth ... chain: azure-openai->openai->anthropic` is env_file mounting the clone's openai+anthropic
  keys beside platform's azure — not a pinned chain. Recorded as D6; the durable artifacts (docs) are correct, so
  the historical commit message is left as-is (no rewrite).
- **P3 — the build sub-agent froze mid-live-verify** (during the ~10-min async builder GENERATE). The orchestrator
  took over, committed the docs, and routed the live end-to-end drive to M254. No data loss; tree left clean.

## What Went Well
- **The pinned regression tests earned their keep.** `test_studio_desk_env_clerkenstein_no_mock_and_offset_sign_in`
  + `test_studio_desk_block_shape_single_port_clerkenstein_wired` turned a tempting-but-wrong `MOCK_CLERK` fix into
  a hard NO before it could ship — the misdiagnosis cost zero code churn.
- **`platform_dir=None` default kept the change byte-identical for existing callers** (`exposure_claim_guard`'s
  `frontend_lines(n, offset)`), so the env_file addition is purely additive and backward-safe.
- **Minimal, correct fix.** No chain-pin was needed because `aiService.getCompletion` loops every provider within
  one request — azure fast-fails on its non-studio key and falls through to the clone's real openai key in the same
  request. The env_file alone makes the builder work; verified live (op1).
- **Clean scope boundary.** studio-desk's first manifest entry reused the existing cockpit seat-switch + hero-login
  + the M42 e2e foundation; `studioBaseUrl` cleanly mirrors `hiringAppBaseUrl`. Zero platform edits.

## What Didn't
- **Flag/data-gated surfaces read poorly from a raw probe.** As in M248, a black-box `curl` gave a confident-but-wrong
  signal; the truth needed the auth model, not the HTTP status. The correction was cheap here (caught by tests before
  any code), but the pattern recurs across the demo's Clerkenstein-authed surfaces.
- **The ~10-min async builder GENERATE is flaky to drive locally** — it froze the build agent. Centralizing the live
  RUN at M254 (the billion-last design) is the right home, but it means M252 ships BUILD-proven + wiring-live-proven,
  with the end-to-end RUN deferred.

## Carried Forward
- **CARRY-M252-01 → M254 (Fate-2, gate g):** 5 pre-existing `stack-verify` unit failures (`TestAutoVerify` +
  `TestDirectusCheapWins`) from M245's academy `/library/` cheap-win added without updating 3 older curl stubs.
  Byte-identical pre/post M252 (M252 adds 0 failures); an academy-harness reconciliation, out-of-subject for a
  studio milestone. M254 gate (g) "the live/docker-gated test-health tests green" + close-release quality review own it.
- **CARRY-M252-02 → M254 (Fate-2, gate e+h):** the LIVE end-to-end builder-generate Playthrough RUN (the ~10-min
  async Generate). Wiring PROVEN + Playthrough artifacts complete (ptvalidate 18 live/0 TODO); the live RUN is M254's
  chartered domain (all cold-reset-to-seed live-proof is centralized there by the billion-last design).
- **D5 → M247-reconcile tail / close-release Phase 3b (Fate-2):** 2 stale playthrough-count mirrors
  (`README.md:206`, `CLAUDE.md:321` "16 live"). The authority (`playthroughs.md`) is correctly at 18; CLAUDE.md is
  M247-sole-owned (coordination rule #5), so both mirrors reconcile in the tail together — not split-brained in M252.

## Metrics Delta
- Playthroughs: 16 (M243) → **18 live / 0 TODO** (+2 studio: `pt-studio-advanced-generate` + `pt-studio-guided-generate`).
- rext tests: `test_injection` 156 OK (+3 env_file) · `exposure_claim_guard` 49 OK · 8 net-new studio autoverify (g)
  tests · 2 net-new playthroughs unit spec files (`studio-builder-locators` + `url-shapes`) + extended `stack-env` — unit-green.
- Flake: 0. Platform-repo edits: 0. Code-of-record: `july-jitter-m252-studio-builder @ d80db9f` (on origin, rung-zero verified).
- Full metrics: `metrics.json`.
