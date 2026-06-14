# M28 — Spec notes

_Technical detail accumulated during build. Stub at scaffold; sections seeded from the overview scope._

## Per-repo target-file map
The map is the DNA's `target_file` per gene (no hardcoded table — provision groups measurable genes by
`(repo, target_file)` and writes each unit): `platform/.env`, `app/.env`, `studio-desk/.env`,
`ant-academy/code/.env.local` **[pinned exact filename]**, `next-web-app/apps/web/.env`, `sentinel/.env`.
Nested target dirs (`ant-academy/code`, `next-web-app/apps/web`) are `MkdirAll`'d before write.
**Alias-mapping:** one source value → every alias-family member's key. provision resolves a family's source
once (the first member present in the source — `resolveAliasSources`); a member whose OWN key is absent from
the source is filled from that sibling value (e.g. `GH_PAT` → `GH_ACCESS_TOKEN` + cross-repo `app/GH_TOKEN`).
A member whose own key IS present uses its own value (never the alias). Distinct-similar pairs
(`OPENAI_KEY` vs `OPENAI_API_KEY`, the Azure trio, `ANTHROPIC` vs `AI_ANTHROPIC`) carry no alias → never
auto-copied. **M28-D1:** the LiveKit key/secret were de-aliased (a credential pair = two distinct values).
Code: `provision/provision.go` (engine), `provision/io.go` (the value-carrying boundary), `cmd/stacksecrets/main.go`
(`provision` verb).

## Idempotency + overwrite + N=0 guard
_(copy-if-absent default, --force, refuse N=0 without --force — mirror stackseed --reset.)_

## Compose with the injection override (the DIRECTUS_TOKEN safety class)
_(provision runs BEFORE gen_injected_override.py; never re-arm the stripped prod DIRECTUS_TOKEN on non-prod /
--local-content; the regression test.)_

## PreflightEnv emission
_(reuse the seeder's values-blind env-guard discipline.)_

## check/measure metrics + per-repo rollup + demo-awareness
_(Overall weighted + Critical unweighted gate==100%; per-repo "short key Y"; Clerk-minted-OK for demo.)_

## Bring-up pre-flight wiring
Shared wrapper `stack-secrets/preflight.sh` (the verification.md non-fatal convention): runs `stacksecrets
check`, then:
- **exit 0** — critical coverage 100% (standard-missing keys WARNED via the per-repo `short:` lines, non-fatal);
- **exit 1** — a CRITICAL key missing → the bring-up STOPS with a clear, actionable message;
- **exit 2** — SKIP (no source dir = the documented manual-`.env` path; or any pre-flight bug) — never blocks
  a good bring-up.
Default source `.agentspace/secrets`; `--demo` makes it demo-aware (minted Clerk keys count). Wired into
`dev-stack` cmd_up (`DEV_NO_SECRET_PREFLIGHT=1` opts out) BEFORE allocation/build, and `demo-stack/up-injected.sh`
(`DEMO_NO_SECRET_PREFLIGHT=1`, `--demo`) BEFORE the heavy per-frontend build. Both capture the rc under `set -e`
(`|| pf_rc=$?`) so only a critical miss is fatal. shellcheck-clean.

## Demo-aware coverage
`secretdna/MeasureForStack(src, d, demo)` — on a demo stack a `mintedSource` overlay reports the
Clerkenstein-minted keys (`MintedKeys`: `CLERK_SECRET_KEY`, the 3 publishable-key variants,
`CLERK_WEBHOOK_SECRET`, `CLERK_JWT_KEY`) as present + non-empty + the right Clerk shape (pk/sk) so their
operators pass WITHOUT the source carrying them. Non-Clerk criticals (GH_PAT, OPENAI_KEY, DB_CONNECTION, …)
still require the source on a demo. Dev stacks get plain `Measure` (real keys required). Code: `secretdna/demo.go`.

## HARD safety — values-blind
The one value-carrying read is `provision/io.go::sourceValues`; its values flow ONLY into the `WriteString` to
the gitignored target `.env`, never to a log/error/return/Report/stdout/stderr. `readTargetKeys` (copy-if-absent)
is NAMES-only. Proven by `provision/provision_safety_test.go` (a reflection-walk asserts no sentinel value
surfaces in the Report / dry-run plan / errors) + the CLI-output no-leak test + the preflight no-leak test.

## Profile scoping
_(v1 default graphql profile, or per-gene profile tag.)_

## Pre-flight audits — provisioning engine (all M28 sections)
KB-fidelity audit `--milestone=M28` ran **GREEN** (report: `kb-fidelity-audit.md`, sha `3b528f8`). Topic→doc→code
triples (verified ALIGNED, no fixes needed):
- PreflightEnv / 3-layer isolation → `corpus/ops/safety.md` §2.1–2.5 (anchor `safety.md:156-205` exact) →
  `stack-seeding/isolation/`.
- N=0 guard precedent → `safety.md` §2.5 → `stackseed` `doReset` (`n == 0 && !force` refuse) +
  `dev-setdress.sh:96`.
- DIRECTUS_TOKEN strip (fix16/17) → `rosetta_demo.md` + `safety.md` §2.3 →
  `stack-injection/gen_injected_override.py:273` (`DIRECTUS_TOKEN=` on every emitted service) +
  `stack-core/gen_override.py:84` (dev data-plane consumer strip). The DNA gene `platform/DIRECTUS_TOKEN`
  is **key-present only (no nonempty)** so a deliberately-blanked non-prod value still passes coverage.
- Idempotency (copy-if-absent) → `corpus/ops/idempotency.md`.
- Non-fatal pre-flight (warn standard / fail critical) → `verification.md:48-49` → wired into `dev-stack` cmd_up,
  `dev-setdress.sh`, `demo-stack/up-injected.sh`.
- Demo Clerk minting → `clerkenstein.md` (mintpk + `inject.py` → PK_DEMO / `sk_test_<demo>` baked into
  `.env.demo-N`, NOT sourced from the secret dir).
- M27 base contract (DNA shape + source layout + exit codes) documented IN-CODE (package doc-comments in
  `dna.go` / `source.go` / `main.go`); `secrets-spec.md` is M29's `Delivers →`, not M28's.
