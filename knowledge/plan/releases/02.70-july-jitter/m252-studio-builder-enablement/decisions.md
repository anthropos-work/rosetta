# M252 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

## KB-fidelity tracked items (pre-flight audit 2026-07-24, YELLOW — see `kb-fidelity-audit.md`)

- **KB-1 (secrets-spec.md M50 imprecision).** The M50 "AI-provider keys — documented-as-absent" box
  (`corpus/ops/secrets-spec.md:269-278`) says AI keys "remain in the waived/optional class." studio-desk's
  `AI_OPENAI_API_KEY`/`AI_ANTHROPIC_API_KEY` are actually **required·standard** (warn, not waived) — the SAME
  posture as the M239 Bedrock class below it, and the studio-builders live-AI surface is undocumented. **Land as
  part of the M252 secrets-spec.md Deliverable** (a studio-builders carve-out mirroring the M239 Bedrock carve-out).
  Deferred (not fixed inline) to avoid pre-empting the env_file-vs-bridge decision.

## Implementation decisions

- **D1 — the fix is env_file ONLY, pointed at `<clone>/studio-desk/.env`, existence-guarded (NO MOCK_CLERK, NO
  provider-chain pin).** The AI-key wiring is added to the studio-desk overlay of
  `stack-injection/gen_injected_override.py` `frontend_lines()` (thread `platform_dir` in; keep default `None` so
  `exposure_claim_guard.py:142`'s `frontend_lines(n,offset)` stays byte-identical). It emits an existence-guarded
  absolute `env_file: ["<abs>/studio-desk/.env"]` (the studio-desk clone is a sibling of platform/), which mounts
  the clone's own provisioned `AI_OPENAI_API_KEY` (len 164) + `AI_ANTHROPIC_API_KEY` (len 108) into the container.
  A literal "mirror hiring-app" would point at platform/.env (no AI keys) and fix nothing. The explicit
  `environment:` block still wins (compose precedence `environment > env_file`); env_file lists concatenate so the
  clone `.env` keys win over platform/.env. Existence-guarded so a stack without a provisioned `studio-desk/.env`
  gets no `env_file` (non-fatal). Values-blind — the file is only referenced.

  **D1 — why NO MOCK_CLERK (a corrected misdiagnosis).** An early raw-`curl` of `:29000` (no session) returned
  `302 → /login`, which I first read as "the demo studio is entirely unreachable, so `MOCK_CLERK=true` is needed to
  disarm the gate." That was WRONG. The studio-desk demo overlay is **deliberately, test-enforcedly**
  Clerkenstein-authenticated (`gen_injected_override.py:85,155-176` + the pinned tests
  `test_studio_desk_block_shape_single_port_clerkenstein_wired` :1499 exact-block `assertEqual` and
  `test_studio_desk_env_clerkenstein_no_mock_and_offset_sign_in` :1560 "NO MOCK_CLERK line"): under prod `NODE_ENV`,
  `clerkMiddleware()` 302s an *unauthenticated* browser into the fake-FAPI `/v1/client/handshake` (networkless
  `CLERK_JWT_KEY` RS256 verify), then `requireAuth` + `checkEnterpriseAndAdmin` (fake-BAPI
  `getOrganizationMembershipList`, an admin/content_creator membership required — the manager hero qualifies). So a
  *logged-in admin hero* reaches the studio; my `curl` 302 was an unauthenticated request behaving exactly as
  designed. Adding `MOCK_CLERK=true` would REGRESS the "actual logged-in hero" demo to the legacy bypass, contradict
  the design comments, and FAIL two regression tests. Lane A correctly refused it. The provisioned clone `.env`
  leaves `MOCK_CLERK=` empty anyway (it is config, not a secret DNA gene, so `/stack-secrets` never fills it — the
  secrets SOURCE template's `MOCK_CLERK=true` is not copied), so env_file could not have carried it regardless.

  **D1 — why NO `AI_PROVIDER_CHAIN` pin (considered, rejected).** platform/.env carries a real `AZURE_OPENAI_KEY`
  (len 32) + endpoint (len 44), so `config.ts`'s legacy fallback (`AI_PROVIDER_CHAIN` empty) tries **azure first**
  using platform's non-studio key. I considered pinning `AI_PROVIDER_CHAIN=openai,anthropic` in the environment
  overlay to force the clone's real keys. Rejected: `aiService.getCompletion` (`src/services/aiService.ts:118-175`)
  loops **every configured provider within a single request** (`for i in orderedProviders … try/catch continue`),
  so azure fast-fails on its bad key (a reachable-endpoint 401/404, not a 90 s hang) and the loop falls through to
  the clone's real openai key **in the same request** — the builder works on env_file alone. Pinning would modify a
  high-care, exact-block-pinned invariant (`:1499`) for a marginal skip-the-fast-fail-detour benefit not needed for
  correctness. Verified live (see the M252 live-verify). If live evidence had shown an azure hang, the pin was the
  ready fallback.

- **CARRY-M252-01 (Fate 2 — pre-existing, covered by M254 gate-g).** Running the full `stack-verify` unit suite
  surfaced **5 PRE-EXISTING failures** — `TestAutoVerify.{test_healthy_stack_exits_zero,
  test_health_only_failure_still_warns_and_is_non_fatal}` +
  `TestDirectusCheapWins.{test_cheap_wins_skip_when_directus_container_absent,
  test_no_prod_read_assert_skips_when_dsn_unreadable, test_registered_collections_nonzero_passes}`. Root cause:
  **M245** added the academy `(f)` cheap-win to `autoverify.sh` (greps `/library/` for a course card) but did NOT
  update the pre-M245 curl stubs in those classes to serve `/library/`, so the academy assert warns and the
  `warnings==0` tests fail. **NOT introduced by M252** — Lane A proved the clean-tree baseline failure set (HEAD
  `584f1fe`) is **byte-identical** to the post-change set (M252 adds **0** failures); M252's own 8 studio-desk
  `(g)` tests all pass (Lane A's `_run_studio` harness correctly stubs `/library/`). Fate 2: **M254 "prove on
  billion" exit-gate part (g)** already commits to "the … test-health tests green" and close-release runs a
  release-level quality review — the fix (mirror the `*"/library/")` stub arm into the 3 older stubs) is a
  contained academy-harness reconciliation that belongs there, not smuggled into the studio-builder milestone
  (out-of-subject; the cross-class stub gating carries its own risk). Recorded so it is not lost.

- **D2 — the container-carries-provider-key assertion lives in `stack-verify/live/autoverify.sh`, not
  `stack-secrets`.** The roadmap phrased it as "DNA hardening (rext stack-secrets)", but `stack-secrets` is a pure
  **source-dir vs DNA** harness — VALUES-BLIND, zero docker — and adding container inspection would violate its
  contract. The correct home is the live-verify layer, which ALREADY inspects a container's `.Config.Env`
  (`autoverify.sh:124-133`, the directus `DB_CONNECTION_STRING` check — the exact idiom to mirror). The assert is
  demo-aware, **non-fatal** (warn, never abort), values-blind (presence-and-nonempty of a provider-key NAME:
  `AI_OPENAI_API_KEY` | `OPENAI_KEY` | `AI_ANTHROPIC_API_KEY` | `AI_AZURE_KEY`). This closes the exact
  `.env`-vs-container gap the wiring fixes. secrets-spec.md gets a cross-ref note (the container-side proof lives in
  live-verify).

- **D3 — the builder Playthrough uses `pt-manager` (the seeded org-admin manager hero) + the assertion boundary.**
  studio-desk is studio-desk's FIRST playthroughs-manifest entry. Hero = `pt-manager` (Morgan Reyes,
  `vantage: manager` → org admin — the "Dan"-equivalent; workforce end-user heroes lack a Studio-eligible role).
  Two Use Cases (advanced + guided GENERATE). Because these are real-LLM flows, the spec asserts at the **completion
  boundary** (a result landmark rendered / no 500), never the generated content (P6 discipline,
  `playthroughs.md:490-495`) — cheap tier + bounded prompt.

- **D4 — M239 Bedrock talk-to-data: CONFIRMED COMPLETE, no work owed.** 5 genes on the `app` repo in
  `secret-dna.json` (`AWS_ACCESS_KEY_ID` + `AWS_SECRET_ACCESS_KEY` required·standard; `AWS_REGION` +
  `AWS_SESSION_TOKEN` + `CLAUDE_CODE_USE_BEDROCK` optional·config), provenance line records M239 explicitly,
  structure + behaviour pinned by `secret_dna_json_test.go:175-222` + `bedrock_measure_test.go`, provisioned +
  bridged by `bridge_bedrock_creds()` (`up-injected.sh:1244-1271`, called `:1420`). Recorded for the audit trail.
