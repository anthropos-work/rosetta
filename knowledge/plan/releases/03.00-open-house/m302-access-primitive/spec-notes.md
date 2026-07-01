# M302 Spec Notes

Technical detail during build.

## API-key store schema (draft)

```sql
CREATE SCHEMA IF NOT EXISTS customer_api;

CREATE TABLE customer_api.api_keys (
    id                UUID PRIMARY KEY,
    organization_id   UUID NOT NULL,
    hashed_key        BYTEA NOT NULL,        -- argon2id output
    key_prefix        TEXT NOT NULL,          -- "ak_live" | "ak_test" + first N chars, for display
    label             TEXT,                   -- human-visible name
    scopes            TEXT[] NOT NULL,        -- e.g. {"people:read","learning:read"}
    entitlement_tier  TEXT NOT NULL,          -- "free" | "paying" | "enterprise" | "partner"
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_by_user   UUID NOT NULL,          -- the Clerk user who minted it
    last_used_at      TIMESTAMPTZ,
    revoked_at        TIMESTAMPTZ,
    rotated_from_id   UUID REFERENCES customer_api.api_keys(id)
);

CREATE INDEX ON customer_api.api_keys (hashed_key) WHERE revoked_at IS NULL;
CREATE INDEX ON customer_api.api_keys (organization_id);
```

## Audit ledger schema (draft)

```sql
CREATE TABLE customer_api.audit_events (
    id               BIGSERIAL PRIMARY KEY,
    ts               TIMESTAMPTZ NOT NULL DEFAULT now(),
    principal_id     TEXT NOT NULL,
    organization_id  UUID NOT NULL,
    resource_id      TEXT NOT NULL,          -- catalog id, e.g. "people.member.list"
    action           TEXT NOT NULL,          -- "read" | "w1" | "w2" | "admin"
    status           INT NOT NULL,           -- HTTP status
    input_hash       TEXT,                   -- for W2 writes only
    latency_ms       INT,
    client_ip        INET,
    user_agent       TEXT
);

CREATE INDEX ON customer_api.audit_events (organization_id, ts DESC);
CREATE INDEX ON customer_api.audit_events (principal_id, ts DESC);
```

## argon2id parameters (draft — confirm at build)

OWASP 2024 baseline: memory 47 MiB, iterations 1, parallelism 1 — tuned so hash+verify < 100ms on `app`
container size. Confirm parameters at build; record chosen values in `decisions.md`.

## Rate-limit bucket keying

- **Key format:** `rl:{principal_id}:{bucket_id}:{window}`.
- **Buckets:** `default-min` (60/min), `default-day` (10k/day). Per-resource buckets can override (rare).
- **Store:** Redis (already in the stack).

