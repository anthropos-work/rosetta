---
milestone: M28
slug: provisioning-engine
version: v1.6 "stage door"
milestone_shape: section
status: archived
created: 2026-06-14
last_updated: 2026-06-14
complexity: large
delivers: rosetta-extensions/stack-secrets/ (the provision engine + the demo-aware extension of check/measure + the non-fatal pre-flight wiring into /dev-up + /demo-up); ext tag stage-door-m28
scope_note: "The base check/measure SCORER (Overall/Critical/per-repo-rollup + exit-1-if-critical) shipped early in M27 (see M27-D3.2). M28 owns only the DEMO-AWARE extension (Clerkenstein-minted Clerk keys satisfy coverage) + the pre-flight WIRING into /dev-up + /demo-up + the whole provision engine."
backlog_refs: (none)
---

# M28 — Provisioning engine + coverage/verify gate

## Goal
`stacksecrets provision` writes each repo's target `.env` from the source (correct exact key per repo,
alias-mapped per file), values-blind; the M27 `check`/`measure` scorer is extended to be **demo-aware** and
wired non-fatally into `/dev-up` + `/demo-up` pre-flight.

> **Scope adjusted 2026-06-14** (M27-D3.2): the base `check`/`measure` scorer (Overall/Critical/per-repo rollup +
> exit-1-if-critical-missing) already **shipped in M27** alongside the DNA. M28 no longer builds it from scratch —
> it adds the **demo-aware** variant (Clerkenstein-minted Clerk keys count as satisfied) and the **pre-flight
> wiring**, on top of the whole `provision` engine.

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
  - **Demo-aware `check`/`measure`** — the base scorer (Overall weighted + Critical gate==100% + per-repo rollup +
    exit-1-if-critical-missing) **already landed in M27**; M28 adds the **demo-aware** variant only: on a demo stack,
    Clerk keys (and any Clerkenstein-minted credential) count as satisfied without being in the source dir.
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
