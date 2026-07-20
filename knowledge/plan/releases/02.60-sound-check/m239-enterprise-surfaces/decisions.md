# M239 — Decisions

_(implementation choices with rationale accumulate here during build)_

- **User decision (2026-07-20): talk-to-data → FULL.** Real AWS Bedrock creds via `/stack-secrets` + secret-coverage DNA extension for `app` (reference `../hyper-studio/.env.example`), not just a flag. Recorded at design time; carried here for build traceability.
