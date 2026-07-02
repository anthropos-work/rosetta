# M302 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in [`overview.md`](overview.md).

## Section checklist

- [ ] **(1) API-key store** — Postgres schema + migration + repo layer; argon2id hashing (parameters decided +
  recorded in `decisions.md`).
- [ ] **(2) `ApiKeyIdentityProvider`** — the second `IdentityProvider` adapter; contract tests (given a key,
  produce a `Principal`; unknown/revoked → 401).
- [ ] **(3) Mint / rotate / revoke path** — three admin-tier catalog entries + handlers; a Clerk-authenticated
  admin Principal is the only caller; every lifecycle event lands in the audit ledger.
- [ ] **(4) Audit ledger** — the `customer_api.audit_events` schema + the shared middleware; contract tests
  (every request writes a row, success + failure; W2 shape carries only `input_hash`).
- [ ] **(5) Rate-limit middleware** — Redis token-bucket + `X-RateLimit-*` headers + 429 + `Retry-After`;
  contract tests (default budgets fire; exhaustion returns 429).
- [ ] **(6) `/v1/access/whoami` + `/v1/access/api-keys` (list)** — the proof handlers; feature-flagged,
  internal-only.
- [ ] **Docs** — the two openapi entries generated from the catalog; `decisions.md` records the
  argon2id-vs-bcrypt call (spec-progress #1) and the Redis-vs-Postgres rate-limit call.

**Status:** `planned` — not yet started. Next: `/developer-kit:build-milestone` (M302) after M301 closes.

