# M27 — Spec notes

_Technical detail accumulated during build. Stub at scaffold; sections seeded from the overview scope._

## Per-repo target-file map
_(platform/.env, app/.env, studio-desk/.env, ant-academy/code/.env.local [pinned], next-web-app/apps/web/.env,
sentinel/.env; one source value → all per-file aliases.)_

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
_(non-fatal warn on standard-missing, fatal on critical-missing in /dev-up + /demo-up.)_

## Profile scoping
_(v1 default graphql profile, or per-gene profile tag.)_
