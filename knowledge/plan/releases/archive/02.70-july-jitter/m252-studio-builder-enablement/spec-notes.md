# M252 — Spec notes

Topic → doc → code triples + studio-builder wiring / Playthrough findings accumulate here during build.

> **Pre-flight KB-fidelity audit (2026-07-24): YELLOW.** Full report: `kb-fidelity-audit.md`. Two stale
> load-bearing claims in `studio-desk.md` FIXED inline (F1 phantom `designer-sim.html`; F2 admin-only role gate).
> KB-1 tracked (see `decisions.md`). Verified triples below.

## Topic → doc → code triples (verified)
- **studio builder pages + AI endpoint** → `corpus/services/studio-desk.md` → `stack-dev/studio-desk/vite.config.ts:28-42`
  (entries incl. `sim-advanced-builder` / `sim-guided-builder` / `simulation-builder`), `src/index.ts:96,99-128,191-250`
  (routes + `STUDIO_ACCESS_ROLES` gate), `src/routes/ai.ts:36` (`POST /api/ai/completion`),
  `src/services/aiService.ts:178` + `src/services/ai/config.ts` (500-on-no-provider; chain/tier defaults).
- **container wiring** → `corpus/ops/demo/frontend-tier.md:560-566` → `stack-injection/gen_injected_override.py:145-181`
  (studio-desk static block), `:237-248` (`studio_desk_local_directus_env`), `:273-290` (`frontend_lines` overlay),
  `:309-356` (`hiring_lines` — the explicit `env_file` reference impl; note it targets **platform/.env**, not a per-repo file).
- **secret DNA** → `corpus/ops/secrets-spec.md:94-115,269-320` → `stack-secrets/secretdna/secret-dna.json`
  (studio-desk 7 genes; `AI_OPENAI_API_KEY`/`AI_ANTHROPIC_API_KEY` = required·standard; split 42/11/8 + 13 crit; version `sound-check-m239`).
- **bedrock bridge (M239 double-check)** → `corpus/ops/secrets-spec.md:311-320` → `demo-stack/up-injected.sh:1244-1271` (call `:1420`). **COMPLETE, no gap.**
- **Playthroughs** → `corpus/ops/demo/playthroughs.md:106` (16 live / 0 TODO) → `playthroughs/e2e/lib/stack-env.ts`
  (`hiringAppBaseUrl` 3001+offset — the pattern to mirror; **no `studioBaseUrl` yet**), `playthroughs/manifest/*.yaml` (no studio entry — this is its first).

## AI-key demo-container wiring (env_file vs bridge)
- **Root cause CONFIRMED:** studio-desk container inherits `env_file: .env` = platform/.env (BASE_ENV), which carries
  no AI key; the provisioned studio AI keys land in `studio-desk/.env` (DNA `target_file`) → `/api/ai/completion` 500s.
- **Design nuance (from audit):** "mirrors hiring-app" is imprecise — `hiring_lines()` env_file → **platform/.env**
  (no AI keys). The fix must point env_file at **`<clone>/studio-desk/.env`**, OR add `bridge_studio_ai_creds()`
  copying the studio AI class `studio-desk/.env → platform/.env` (mirroring `bridge_bedrock_creds`).
- **DECISION → env_file ONLY at `<clone>/studio-desk/.env`, existence-guarded (NO MOCK_CLERK, NO chain pin)** (D1).
  Thread `platform_dir` into `frontend_lines()` (`gen_injected_override.py:251`, call site `:491`; keep default
  `None` so `exposure_claim_guard.py:142`'s `frontend_lines(n,offset)` stays byte-identical). Path =
  `os.path.abspath(os.path.join(platform_dir, os.pardir, "studio-desk", ".env"))` (the studio-desk clone is a
  **sibling of platform**). Emit `env_file: ["<abs>"]` in the studio-desk overlay only when the file exists. That
  mounts the clone's provisioned `AI_OPENAI_API_KEY` (164) + `AI_ANTHROPIC_API_KEY` (108) — the whole fix.
- **CORRECTION — the studio is NOT "unreachable"; NO MOCK_CLERK.** An early raw-`curl :29000` (no session) got
  `302 → …/login`; I first read that as "entirely unreachable → needs `MOCK_CLERK`." WRONG. studio-desk's demo
  overlay is **deliberately, test-enforcedly** Clerkenstein-authed (`gen_injected_override.py:85,155-176`; tests
  `:1499` exact-block + `:1560` "NO MOCK_CLERK line"): under prod `NODE_ENV`, `clerkMiddleware()` 302s an
  *unauthenticated* browser into the fake-FAPI handshake (networkless `CLERK_JWT_KEY` verify), then `requireAuth` +
  `checkEnterpriseAndAdmin` (fake-BAPI `getOrganizationMembershipList`, admin/content_creator required — the
  manager hero qualifies). A logged-in admin hero reaches the studio; the `curl` 302 was an unauthenticated request
  behaving as designed. `MOCK_CLERK=true` would regress the "actual logged-in hero" demo to the legacy bypass +
  fail 2 tests. (The provisioned clone `.env` leaves `MOCK_CLERK=` empty anyway — config, not a gene.)
- **CORRECTION — NO `AI_PROVIDER_CHAIN` pin.** platform/.env has a real `AZURE_OPENAI_KEY` (32) + endpoint (44), so
  the empty-chain legacy fallback tries azure first with a non-studio key. But `aiService.getCompletion`
  (`aiService.ts:118-175`) loops EVERY configured provider **within one request**, so azure fast-fails (reachable
  endpoint → 401/404, not a 90 s hang) and falls through to the clone's real openai key in the same request — the
  builder works on env_file alone. Pinning would disturb the exact-block invariant (`:1499`) for no correctness
  gain; considered + rejected (ready fallback if live shows an azure hang).
- **Precedence:** compose CONCATENATES `env_file` (base platform `.env` → override list); the studio-desk clone
  `.env` loads LAST so its keys win over platform `.env`; the explicit `environment:` block wins over both
  (Clerkenstein `CLERK_*`, `DIRECTUS_TOKEN=""`, `NODE_ENV=production`). **Verify the merged container env live.**

## DNA hardening — container-carries-provider-key assertion
- No existing container-side AI-key assertion (net-new). The `.env`-vs-container gap is exactly what F8 documents.

## Builder Playthrough (studio-desk enters the playthroughs manifest)
- Confirmed: 0 studio references in `playthroughs/` today. Add `studioBaseUrl(9000+offset)` to `stack-env.ts`
  (mirror `hiringAppBaseUrl`), a `studio-builder-page.ts`, `manifest/studio-builders.yaml`, admin/content_creator precondition.

## Talk-to-data double-check (M239 Bedrock)
- **CONFIRMED COMPLETE, no gap** — `bridge_bedrock_creds()` present + wired + tested. No work owed.
