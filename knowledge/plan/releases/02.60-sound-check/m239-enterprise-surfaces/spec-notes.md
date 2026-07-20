# M239 — Spec notes

Topic → doc → code triples + Bedrock-secret / library-race findings accumulate here during build.

## talk-to-data — (a) flag + (b) real Bedrock creds
- (a) flag enablement: `NEXT_PUBLIC_DEMO_FLAGS_ALL` or a flag-gate demopatch (M219/M232 pattern).
- (b) real AWS Bedrock creds via `/stack-secrets`: key set `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` / `AWS_REGION` + `CLAUDE_CODE_USE_BEDROCK` (reference `../hyper-studio/.env.example`).
- Secret-coverage DNA extension for the `app` service; mount/env-wire into the `app` compose service.

## Library empty-first-load (#4)
- Client-fetch race + the open non-offset `:5050` `apps/web` endpoint carry.

## Hierarchical manager menu (#1)
- Confirm it renders for managers via the presence sweep (post-M237 fresh build).

## Secrets posture (safety.md)
- New present-not-absent secret class for `app` (same class as the AI-provider keys) — record the note.

_(will accumulate topic → doc → code triples during build)_
