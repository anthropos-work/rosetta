---
milestone: M239
slug: enterprise-surfaces
version: v2.6 "sound check"
milestone_shape: section
status: planned
created: 2026-07-20
last_updated: 2026-07-20
depends_on: M237
delivers: corpus/ops/secrets-spec.md + corpus/ops/safety.md
---

# M239 — enterprise surfaces

**Status:** `planned`  ·  **Shape:** `section`  ·  **Complexity:** medium (large if the Bedrock wiring balloons)  ·  **Depends on:** M237

## Goal
talk-to-data works live; the library grid loads first-time; the hierarchical manager menu is confirmed.

## User decision baked in (2026-07-20) — talk-to-data → FULL
**Wire REAL AWS Bedrock creds** via the `/stack-secrets` provisioning mechanism — not just a flag. Reference `../hyper-studio/.env.example` for the key set: `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` / `AWS_REGION` + `CLAUDE_CODE_USE_BEDROCK`. **Extend the secret-coverage DNA for the `app` service.** This introduces a new present-not-absent secret class for `app` (same class as the AI-provider keys) — a `safety.md` secrets-posture note records it.

## Scope
### In
- talk-to-data **(a)** flag enablement (`NEXT_PUBLIC_DEMO_FLAGS_ALL` or a flag-gate demopatch, the M219/M232 pattern) + **(b) real AWS Bedrock creds** provisioned via `/stack-secrets` + the **secret-coverage DNA extension for `app`** (hyper-studio template) + mounted/env-wired into the `app` compose service.
- Fix **#4 (library empty-first-load)**: the client-fetch race + the open non-offset `:5050` `apps/web` endpoint carry.
- Confirm **#1 hierarchical menu** renders for managers (presence sweep).

### Out
- The media-porting content-fidelity work (M240).
- Academy (M238).

## Open questions
- The demo secrets-posture for AWS creds (safety.md note — same class as AI-provider keys, now present-not-absent for `app`).
- Is #4 (library) a pure client-fetch race, or does the non-offset `:5050` `apps/web` endpoint carry contribute?

## Delivers
`corpus/ops/secrets-spec.md` (the Bedrock cred class for `app`) + a `corpus/ops/safety.md` secrets-posture note.

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.6 "sound check" for the authoritative milestone design + the release-level decisions/risks (design spec: `releases/02.60-sound-check/design-notes.md`).
