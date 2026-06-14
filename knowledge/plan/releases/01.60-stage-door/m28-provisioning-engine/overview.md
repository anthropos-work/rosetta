---
milestone: M28
slug: provisioning-engine
version: v1.6 "stage door"
milestone_shape: section
status: planned
created: 2026-06-14
last_updated: 2026-06-14
complexity: large
delivers: rosetta-extensions/stack-secrets/ (the provision engine + check/measure gate) + the non-fatal pre-flight wiring into /dev-up + /demo-up; ext tag stage-door-m28
backlog_refs: (none)
---

# M28 — Provisioning engine + coverage/verify gate

## Goal
`stacksecrets provision` writes each repo's target `.env` from the source (correct exact key per repo,
alias-mapped per file), values-blind; `check`/`measure` computes coverage and is wired non-fatally into
`/dev-up` + `/demo-up` pre-flight.

## Why section
Both halves (write the env, then prove it's complete) share the M27 DNA + a per-repo target-file map that is
concrete and enumerable. The safety interactions (DIRECTUS_TOKEN, N=0, idempotency) are known constraints with
known mechanisms (the injection override, the `stackseed --reset` guard, `idempotency.md`). Build with
`/developer-kit:build-milestone`.

## Repo split
- **`rosetta-extensions`** (authoring → tag `stage-door-m28` → consume): the `provision` verb + per-repo
  target-file map, the idempotency/overwrite + N=0 guards, the compose-with-injection-override safety, the
  `PreflightEnv`-emitting path, the `check`/`measure` metrics + per-repo rollup + demo-awareness, the pre-flight
  hook the bring-up scripts call.
- **`rosetta`**: the `/dev-up` + `/demo-up` skill wiring is in `rosetta-extensions` scripts; any skill-doc
  reference lands in M29.

## Scope
- **In:**
  - `provision` verb — per-repo **target-file map**: `platform/.env`, `app/.env`, `studio-desk/.env`,
    `ant-academy/code/.env.local` **[exact filename pinned — not `code/.env`]**, `next-web-app/apps/web/.env`,
    `sentinel/.env`. One source value → all its per-file aliases.
  - **Idempotency + overwrite policy** — copy-if-absent default, `--force` to overwrite, never silently clobber;
    **N=0 main-dev-stack guard** (refuse without `--force`, mirroring `stackseed --reset`) so it can't clobber the
    operator's working `.env` (the secret source itself).
  - **Composes with + defers to the injection override** — `provision` runs BEFORE `gen_injected_override.py` and
    must NOT re-arm the stripped prod `DIRECTUS_TOKEN` on non-prod / `--local-content` stacks (the fix16/17 safety
    class). **[blocks-release safety — regression test required]**
  - Emit `PreflightEnv`-passing env (reuse the seeder's values-blind env-guard discipline, `safety.md:156-205`).
  - `check`/`measure` — Overall (weighted) + Critical (gate == 100%, unweighted) + **per-repo rollup** ("repo X is
    short key Y"); **demo-aware** (Clerk keys satisfiable by Clerkenstein minting, not the source dir); exit 1 if a
    critical key is missing.
  - Non-fatal pre-flight wiring into `/dev-up` + `/demo-up` (warn on standard-missing, fail on critical-missing —
    the `verification.md` convention).
  - Profile-scoping decision settled (v1 scopes the denominator to the default `graphql` profile, or a per-gene
    profile tag).
  - Hard safety: no verb ever reads/echoes/logs a value; extraction via `grep '^[A-Z_]...'` / cut-on-`=`.
- **Out:** the `/stack-secrets` skill + corpus doc (M29); the end-to-end build-from-stack-dev validation (M30).

## Depends on
M27 (the DNA + the ingestion reader).

## Parallel with
None.

## Estimated complexity
large.

## Open questions
- `check` measure source: static `.env` files (safe, names-by-grep) vs live container env (catches runtime-injected
  keys but risks touching values). **Default:** static, values-blind.

## KB dependencies
- `corpus/ops/safety.md` — `PreflightEnv`, never-write-prod, the 3-layer isolation guard.
- `corpus/ops/idempotency.md` — the run-it-twice contract.
- `corpus/ops/verification.md` — the non-fatal pre-flight convention.
- `corpus/services/clerkenstein.md` — the demo Clerk-minting path (for demo-aware coverage).
- `corpus/ops/rosetta_demo.md` — the injection override + the `DIRECTUS_TOKEN` strip (fix16/17).

## Delivers →
`rosetta-extensions/stack-secrets/` (the engine + gate; ext tag `stage-door-m28`) + the bring-up pre-flight wiring.
