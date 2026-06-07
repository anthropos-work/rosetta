# M10 spike — Directus content snapshot recon (read-only)

_Read-only reconnaissance against prod Directus (`content.anthropos.work`, Directus 11.6.1). No writes. GET only._
_Date: 2026-06-06. Resolves M10-Q1/Q2/Q3 + the load-bearing content-store fork before any code is written._

## TL;DR recommendation

**Run a real per-stack Directus container, fed by replaying a captured public snapshot into the per-stack
Directus-backing Postgres.** Directus is mandatory (CMS + studio-desk integrate with it directly — it cannot be
mocked away) and Directus stores everything in Postgres, so its backing DB is **per-stack-isolated** by construction —
the snapshot framework's existing Postgres `COPY` replay path applies cleanly to it. **Do not** try to feed a per-stack
Directus via its write/import API from the seeder (slow, schema-fragile, and a write path the isolation guard would
have to specially trust). Capture the public subset; replay into the per-stack Directus Postgres; boot a per-stack
Directus pointed at it.

## Q1 — Reachability + what `DIRECTUS_BASE_ADDR` points at

- `DIRECTUS_BASE_ADDR="https://content.anthropos.work"` — **production** (not Tailscale, not localhost).
- **Reachable from here, no VPN**: `GET /server/ping` → `200 "pong"` (~175 ms). `GET /server/info` (auth) → `200`,
  project "Anthropos / Content Manager", **Directus 11.6.1**.
- The dev compose **does not run a Directus container** — it points CMS at prod (`docker-compose.yml` lines 237-238:
  `DIRECTUS_BASE_ADDR` + `DIRECTUS_PUBLIC_BASE_ADDR` = `https://content.anthropos.work`). There is no local Directus
  today → the per-stack Directus is genuinely new (the fork is real).

## Q2 — The content collections

59 collections total: **27 `directus_*` system** + **32 user collections**, grouped (`job_simulations`,
`simulations`, `paths`, `skill_paths`, `sequences`, `learning_resources`, `sim_validation`). Template core:
`simulations`, `skill_paths`, `sim_tasks`, `task_checks`/`task_sub_checks`, `sequences`/`sequences_roles`, `roles`,
`resource`, `learning_resources`, `paths`, `job_simulations`, `job_position`, `library_categories` /
`library_macro_categories`, `languages`, `curators`, `knowledge_asset`, `*_translations`, `sim_features`.

## Q3 — The public ↔ customer split (the capture filter)

The split field is **`private` (boolean)**, present on the top-level template collections, complemented by
`tenant_id` / `Tenant` / `status` / `Visibility`. Prod aggregate counts (read-only, count-only):

| Collection | total | `private=false` (PUBLIC) | `private=true` | `tenant_id IS NULL` | `status=published` |
|---|---|---|---|---|---|
| `simulations` | 2,597 | **647** | 1,950 | 638 | 1,041 (306 of the public) |
| `skill_paths` | 263 | **28** | 235 | 28 | 190 |

- **`private=false` is the public predicate.** For `skill_paths`, `private=false` and `tenant_id IS NULL` agree exactly
  (28/28). For `simulations` they nearly agree: of 647 public sims, **637** have `tenant_id IS NULL` and **10** carry a
  `tenant_id` — the firewall must DECIDE on those 10 (treat `private=false` as authoritative-public, OR intersect
  `private=false AND tenant_id IS NULL` for the strict set). Recommendation: use `private=false` as the public marker
  and flag the 10 for review; do not silently include tenant-bearing rows.
- Child collections (`sim_tasks`, `sequences`, `resource`, …) are **column-less w.r.t. `private`** (the token got
  `FORBIDDEN` on `private`/`tenant_id` there) — they are public **via their public parent** (the exact M9b
  `ParentScopes` pattern: a sim_task is public iff its `simulations` parent is public). The framework already models
  this.
- A "believable subset" lever (M10-Q3): only **306** public sims and **190** public paths are `status=published` —
  the demo can size down to the published public subset.

## Q4 — Where Directus data lives + how a per-stack Directus is fed

- Directus 11 stores **all** content in **its own Postgres** (a database SEPARATE from the app Postgres). The app
  Postgres `cms` schema is unrelated (and 100% customer — excluded, note #3).
- **No Directus-Postgres DSN is exposed locally** — `.env`/compose give only the HTTP API base + a token; there is no
  `DIRECTUS_DB_*` DSN here. So an app-style `pg_dump`/`--dsn` capture of the Directus DB is **not driveable from this
  machine today** without a DSN (the dump-ingest/primary-read source kinds need a Postgres endpoint).
- **Load mechanism for the per-stack Directus (recommended):** capture the public subset, replay it into the
  **per-stack Directus-backing Postgres** via the framework's bulk `COPY` path, then boot a per-stack Directus
  container pointed at that DB. Because Directus's data IS Postgres, this stays inside the **per-stack-isolated** class
  — no special guard exception. (Alternative — feeding a running per-stack Directus through its REST/import API — is
  slower, schema-version-fragile, and introduces a write path; rejected.)

## Q5 — Is the `DIRECTUS_TOKEN` read-only?

**Cannot prove read-only without a forbidden write test** — but strong read-only signals: the token is a **scoped user
role** (role `186d…fd24`, `policies: []`), **NOT admin** — it is `FORBIDDEN` on `GET /roles/{id}` (its own role),
`GET /fields/job_simulations`, `GET /fields/paths`, and on whole collections (`job_simulations`, `paths`,
`learning_resources`) and on `private`/`tenant_id` fields of child collections. So it is a partially-scoped read role,
not a write-all admin. **However it CAN read customer/private rows** (a `private=true` row with a `tenant_id` set
returned 200) → **the token is NOT a pre-filtered public-only credential**; capture MUST apply the `private=false`
filter itself + run the post-capture firewall. **Follow-up:** before the real build, obtain/confirm a **dedicated
read-only, public-scoped capture token** (or a Directus-Postgres read DSN) rather than reusing this app token — its
inconsistent per-collection permissions (403s) would make an API-based capture lossy.

## Q6 — Prod-write-block guard confirmed

`stack-seeding/isolation/isolation.go` already registers `directus` as **`SharedPollutionRisk`** with note _"ONE
global content instance (content.anthropos.work), visible on prod; content must come via snapshot-replay — block
direct writes"_. `Guard.CheckWrite` **always blocks** SharedPollutionRisk writes on a non-prod target (opt-in does not
apply). M10 fits the guard cleanly: **capture = READ** (governed by the snapshot extension's `firewall.AssertPublicOnly`,
the read-side analog), **replay = WRITE only to the per-stack Directus Postgres** (class `postgres` =
`PerStackIsolated` = always allowed). The shared prod Directus is never written.

## Q7 — Media / blobs

Content references files: **10,340 `directus_files`, 100% on the `s3` storage adapter** (Directus's own S3 bucket,
distinct from the app's `STORAGE_S3_PUBLIC_BUCKET`). Templates reference S3-backed media. For the demo MVP, **refs-only
is viable** (the per-stack Directus can keep S3 file rows as references / point at the same read-only S3, or a
local-storage adapter with placeholder assets); full blob mirroring is a follow-on (S3-private is per-stack-isolated,
so blobs CAN be replayed later). Recommend **refs-only for MVP, blobs deferred**.

## Architectural gap M10 must close (for the build)

The existing snapshot framework (`stack-snapshot/`) is **Postgres-DSN + `COPY` oriented** (`pg/`, `source/` — every
source reads over a DSN) and the firewall predicate is **`organization_id IS NULL`**. Directus's public predicate is
**`private = false`** (+ parent-scope for children), and there is no local Directus DSN. M10 therefore needs ONE of:

1. **Directus-Postgres capture (preferred if a DSN can be provisioned):** get a read DSN to the Directus backing
   Postgres, add a `directus` surface enumerating its tables in FK order, with the firewall predicate `private = false`
   (parent-scoped for children) — maximal reuse of the existing `pg`/`COPY`/`replay` machinery. Replay → per-stack
   Directus Postgres → boot per-stack Directus.
2. **API-based capture (fallback if no DSN):** a new `source` kind that reads the public subset over the Directus REST
   API (with a proper public-scoped read token), serialized into the same manifest/payload contract, replayed into the
   per-stack Directus Postgres. More framework work; depends on a clean read token.

## Blockers / decisions for the user

- **DECIDE the content-store fork** (recommended: real per-stack Directus container + replay into its Postgres). ✅ recommended above.
- **DECIDE capture path** (preferred: Directus-Postgres DSN capture → needs a read DSN provisioned; fallback: API
  capture → needs a clean public-scoped read token). **Need:** a Directus-Postgres read DSN OR a dedicated read-only
  public-scoped Directus token — the current app token reads customer rows and has inconsistent 403s.
- **DECIDE the 10 tenant-bearing public sims** + the believable subset size (`private=false` vs
  `private=false AND status=published`).
- **DECIDE media:** refs-only for MVP (recommended), blobs deferred.
