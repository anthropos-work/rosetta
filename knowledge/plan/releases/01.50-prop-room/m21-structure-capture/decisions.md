# M21 — Decisions

Implementation decisions with rationale, numbered `M21-D1`, `M21-D2`, … . Cross-iter decisions live here; per-iter
detail lives in each `iter-NN/decisions.md`. The strategy-of-record `TOK-NN` entries also live here (the milestone's
strategy-evolution chain).

## TOK-01: staged-pipeline build toward the binary serve-anonymously gate — 2026-06-11

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Treat M21's gate as a **6-stage end-to-end pipeline** (build → bootstrap → structure-apply →
replay-exit-0 → boot → serve-anonymously; full table in `spec-notes.md`) and drive **furthest-passing-stage (0–6)**
as the primary per-tik metric — converting the binary gate into a measurable signal. Each tik validates **live
against Docker** (the directus/directus:11.6.1 image is cached; Directus bootstrap/permission empiricism only breaks
live — the reason this milestone is iterative, not section). The central build is the **capture-side structure
artifact**: the 9 user-collection table DDL + the `directus_collections`/`directus_fields`/`directus_relations`
registry rows (filtered to the public content model) + the `directus_files` ref capture — produced as a **new
artifact keyed by the source schema digest** and applied to a fresh bootstrapped stack **before** the existing row
replay, so the target schema digest converges out of the exit-4/exit-5 "digest trap." Reuse the generic
capture/replay/manifest machinery and add structure **additively** (the `Predicate`-field precedent — no parallel
subsystem), honoring the user's simple/maintainable constraint. All capture stays read-only / public-only / behind
`AssertPublicOnly` (extended to admit structural metadata, never loosened) — the never-touch-prod constraint.
**Rationale:** the gap is empirically hard (fix16 spent +479 lines on Directus provision quirks) and the failure
modes (anonymous serving, registry-row carve-out, digest convergence) only surface live — so a staged pipeline with
a per-stage metric + live validation each tik is the honest shape. The structure-as-additive-artifact choice keeps
the change inside the existing well-tested machinery rather than spawning a second capture path.
**Strategy class:** new-direction (bootstrap).
**Distance-to-gate context:** gate metric = furthest pipeline stage passing; gate = stage 6 (serve a captured sim
anonymously over HTTP). Static baseline today = **stage 2 of 6** (build + bootstrap implemented; structure-apply is
the `provision.go:108` placeholder; replay exits 4; boot/serve blocked behind). The row half of the cache is
complete; only the structure half is missing.
**Next-tik direction:** iter-02 (first tik) — stand up the live baseline harness (throwaway Postgres + bootstrapped
Directus), confirm the live baseline (replay exits 4), resolve the structure-source question (lean: option (c) — a
self-contained reference Directus whose schema is exported via Directus `schema snapshot`, no prod access, doubles
as the test fixture), and produce the first structure artifact for the 9 collections. Target: stage 2 → 3.

---

## Cross-iter implementation decisions

### M21-D1 — admin email: `.local` TLD rejected by Directus 11.6.1 → `.example.com` (iter-02)
**Context:** the per-stack Directus bootstrap (`provision.go` `DefaultEnvContract.AdminEmail`) minted
`admin@<stack>.local`. **Live finding (iter-02, fresh `directus/directus:11.6.1` bootstrap):** the 11.6.1 email
validator rejects the **`.local` TLD** outright — `admin@dev-5.local` AND `admin@dev.local` both die
`FAILED_VALIDATION` in `createDefaultAdmin`, crashing bootstrap. Hyphens/digits in the label are fine
(`admin@dev-5.example.com` passes); the prior fix16 hyphen-vs-underscore framing was incomplete. **Decision:** mint
`admin@<stack>.example.com` (RFC-2606 reserved — never a real address, always format-valid), keeping the stack name
as the subdomain for provenance. Code + comment + tests updated (a `.local` guard pins it gone). The old value was
never executed (print-only recipe); M21 is its first live run. Committed: rosetta-extensions `98e51b4`.

### M21-D2 — structure-apply mechanism = Directus `schema apply` of a snapshot YAML (iter-02)
**Decision (mechanism validated live):** stage 3 (structure-apply) is implemented by Directus's own
`node cli.js schema apply <snapshot.yaml>`. A schema snapshot is a clean YAML
(`version / directus / vendor / collections / fields / relations`); applying it creates BOTH the user-collection
table AND its `directus_collections`/`directus_fields`/`directus_relations` registry rows in one step (proven on a
1-collection snapshot: table + registry rows appeared). This is the structure-artifact CARRIER — it supersedes a
hand-rolled DDL + raw registry-row COPY, and it's Directus-native (so the registry is internally consistent). The
M21 structure artifact = a Directus schema snapshot YAML scoped to the 9 public content collections.

### M21-D3 — live baseline is exit **5** (digest miss), not exit 4 (iter-02)
**Correction to the static baseline (spec-notes iter-01):** the real provision pipeline bootstraps the directus
schema BEFORE replay, so `stacksnap replay --surface directus` runs against a **bootstrapped** schema (27 system
tables) → `pg.SchemaVersion` returns a non-nil digest → it is NOT `ErrEmptySchema` → replay reaches the cache
lookup and **exits 5** (cache miss at the bootstrapped digest `b4cb55bc…` ≠ prod key `6cd35278…`). Exit **4**
(`ErrEmptySchema`) occurs ONLY against a never-bootstrapped (empty) directus schema. Both confirmed live. The
stage-table in spec-notes is updated accordingly.

### M21-D4 — structure-source: pure option (c) can't provide prod types; resolve in iter-03 (iter-02)
**Finding:** TOK-01 leaned toward option (c) (a self-contained reference Directus) for the build/test loop. iter-02
validated that (c) proves the *format + mechanism*, but it **cannot invent prod-faithful column types** — and the
real artifact needs them, because stage 4 COPYs the cached **real** rows into the typed columns (a type mismatch
fails the COPY). **Decision:** the real structure artifact must be sourced from something carrying prod's actual
types + registry. iter-03 resolves the source — FIRST checking the platform repos (cms / migrations) for a committed
Directus schema snapshot or collection definitions (self-contained, prod-faithful, zero prod access); else weighing
(a) prod read / (b) restored dump / an MCP **structural** read (information_schema + directus_collections/fields/
relations as JSON — distinct from the cold-start row-capture ban, which is about COPY bytes). All options stay behind
the M9a capture-source policy + `AssertPublicOnly`.

### M21-D5 — the digest trap is full-schema-keyed; convergence is deeper than "apply before replay" (iter-02)
**Characterization:** `pg.SchemaVersionSQL()` digests every column of every table in the `directus` schema. The prod
cache key `6cd35278…` therefore encodes the WHOLE prod directus schema (system tables AT prod's Directus version +
ALL prod content collections + their exact types). A per-stack bootstrap converges that digest only if its entire
schema matches — which is fragile against (i) Directus version skew (system-table columns differ), and (ii) prod
having content collections beyond our 9. So "apply the 9-collection structure before row replay" (TOK-01) is
necessary but may be **insufficient** for a cache-HIT. The stage-4 resolution (a future M21-D or a tok) likely
chooses between: **(A)** capture/define the FULL prod content-model + pin the Directus version so the digest
converges exactly, or **(B)** re-key the cache per-surface over only the captured content tables (a surgical change
to the staleness key, shared with taxonomy — handle with care). Tracked as `STRUCT-M21-digest-keying`.

### M21-D6 — structure-source is operator-gated; self-contained options exhausted (iter-03 planning)
**Investigation (iter-03 Phase 1):** searched the stack workspaces for a self-contained, prod-faithful Directus
structure source. Found `stack-dev/cms/internal/directus/collections/*.go` — but these are the cms service's
**read-side application view**: JSON-tagged Go structs (uuid.UUID / *int / time.Time / json.RawMessage) mapping
Directus API responses to domain types, including **relational alias fields** (`roles.*`, `sequences.*`,
`library_categories.library_categories_id.*`) that are not real columns on the base table. They carry field *names*
but **not** authoritative Postgres column types nor the Directus field-registry metadata (interface/special/relation
defs). Reconstructing a faithful `schema apply` snapshot from them is lossy and would not guarantee the real cached
rows COPY in (stage 4).
**Conclusion:** the self-contained structure-source options are **exhausted** — TOK-01's option (c) (a hand-built
reference Directus) can't invent prod-faithful types (M21-D4), and the cms structs are a lossy app-view. The faithful
structure (real column types + the `directus_collections`/`fields`/`relations` registry rows for the 9 public
collections) exists only in the **prod directus schema**. Reading it is a **prod read**, which the milestone Risk
section mandates be **operator-confirmed** (read-only / bounded / public-only / behind the M9a capture-source policy +
`AssertPublicOnly`, extended to structural metadata). The milestone `overview.md` "Open questions" already leans
toward `pg_dump --schema-only -n directus` over the sanctioned `--dsn` path (option b) — but no such dump exists
locally, and a live prod read needs operator sign-off. **This is a user/operator decision; surfaced via
`build-mstone-iters` Phase 5 §4 (user-blocker).** Options put to the operator: (1) sanction a bounded read-only prod
**structural** read (information_schema types + the registry rows for the 9 public collections) via the wired
`postgres` MCP or a `--dsn`; (2) provide / point at a restored prod `pg_dump -n directus` (the cold-start (b) path);
(3) proceed self-contained with best-effort inferred types (lower fidelity — may not COPY the real rows; defers exact
prod-fidelity to a release-time (b) capture). Until chosen, the real 9-collection structure artifact
(`STRUCT-M21-iter03-artifact`) is blocked.
