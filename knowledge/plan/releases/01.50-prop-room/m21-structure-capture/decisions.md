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

**iter-03 evidence — surfaced the fork (operator chose A — see M21-D7).** _(The analysis below recommended B; the
operator decided **option A**. M21-D7 is the decision of record; this paragraph is kept for the reasoning trail.)_
A sanctioned prod structural read showed the prod `directus` schema
digest `6cd35278…` is computed over the **FULL 53-table schema**: 27 `directus_*` system tables + **26 user
collections**. The directus surface captures **9** of those 26. So a per-stack bootstrapped Directus (27 system
tables) + the 9-collection structure can **never** converge the `6cd35278…` digest — it is short 17 collections (and
would also need the system tables at prod's exact Directus version). Option A would require capturing/applying ALL 26
collections + pinning the version (heavy, drags in 17 unwanted collections). **Option B** (re-key the cache by a
digest over only the surface's captured content tables) is the surgical fix that makes the bootstrap+structure model
cache-hittable and generalizes to taxonomy. **Caveat (why operator-surfaced, not auto-decided):** `pg.SchemaVersionSQL`
+ the staleness key are SHARED with taxonomy and are load-bearing + well-tested; re-keying must not silently change
taxonomy's cache behavior. A vs B (and B's exact scoping) is the iter-04 architectural fork.

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

### M21-D7 — digest-keying RESOLVED: option A (capture all 26 collections + pin the Directus version) (operator decision, 2026-06-11)
**Decision (operator):** resolve the digest-convergence fork (M21-D5) with **option A**, not the per-surface re-key
(B). The cache stays keyed by the whole-`directus`-schema digest; convergence is achieved by making the per-stack
schema MATCH prod's: apply the structure for **all 26 user collections** (not just the 9 the row-surface captures) and
**pin the Directus version** so the 27 system tables also match. Rationale (operator): keep the keying untouched
(zero blast-radius on taxonomy's shared `pg.SchemaVersionSQL`), and a fully prod-faithful per-stack directus schema is
the more honest target.
**Implications for iter-04+ (reframes the structure scope):**
- The structure artifact must cover **all 26 user collections'** DDL + their `directus_collections`/`fields`/
  `relations` registry rows (the read is sanctioned, M21-D6) — superseding the 9-collection framing in TOK-01 / the
  earlier routes. The **row** cache stays at the 9 public-content collections (the other 17 tables exist but are
  empty — fine for the digest, which is over column structure, not rows; and the firewall/public-only predicate still
  governs which ROWS are captured).
- **Pin + verify the Directus version:** confirm prod's Directus version and that bootstrapping that exact image
  produces system tables whose columns match prod's (the bootstrapped-11.6.1 system digest was `b4cb55bc`; iter-04
  must verify prod's 27 system tables match the pinned image, else the system-table half of the digest won't
  converge). Record the source version in the manifest + pin the local image (the overview's version-skew lean).
- Referential closure of relations among the 26 (vs the 20-dangling-from-9 count) is largely subsumed — with all 26
  collections present, most intra-directus relations resolve; M23 still owns any remaining external refs.
`STRUCT-M21-digest-keying` is now "implement option A" (was "decide A/B").

### M21-D8 — Option A validated end-to-end through stage 4 (iter-04)
**Result:** the M21-D7 option-A path is proven on a live harness. On a fresh bootstrapped `directus/directus:11.6.1`
(27 system tables, digest `b4cb55bc…` = prod's system-only digest → no version skew), applying the **26-collection
structure** (real prod DDL + 8 junction sequences, `iter-04/structure.sql`) makes the directus-schema digest
`6cd35278…` — **exactly the prod cache key**. `stacksnap replay --surface directus --stack dev-5` then **exits 0**,
loading the 9 captured content tables (10128 rows; simulations=304). So **stages 3 (structure-apply + digest
converge) and 4 (replay exit 0) both pass**.
**Key structural insight:** the cache's staleness digest (`pg.SchemaVersionSQL`) is over **column structure**
(information_schema.columns), NOT registry rows. So digest convergence + replay (stage 4) need only the 26 collection
**tables** to exist with prod-faithful columns — the `directus_collections`/`fields`/`relations` registry ROWS are
needed only for SERVING (stages 5–6). This decouples stage 4 from the harder serve/permissions work.
**Caveat (what's NOT yet done):** the structure-apply here was a **hand-applied** real-DDL artifact (psql -f), not yet
produced+applied by `stacksnap` itself. The exit GATE says "stacksnap applies the captured structure", so the
remaining stage-3 work is the **code-ification** (`STRUCT-M21-codeify`): a stacksnap capture-side structure extension
that captures the 26-collection DDL + registry over `--dsn` into the snapshot and applies it before the row replay in
provision. iter-04 proves the target is reachable; the tooling that automates it is the next build.

### M21-D9 — the anonymous-serve recipe + the PRIMARY-KEY finding (iter-05)
**Result:** the milestone's flagged live-only risk is cracked — a booted per-stack Directus serves a captured public
simulation anonymously (`GET /items/simulations?limit=1` → 200 with a real published row). The full 6-stage path is
demonstrated end-to-end with the real captured structure.
**The load-bearing finding — PRIMARY KEY constraints:** the iter-04 `structure.sql` (column-only `pg_catalog` DDL)
created tables WITHOUT primary keys. The digest still converged (it's over column *types*, not constraints) and the
row COPY worked — but **Directus refuses to serve a collection with no detectable PK** (`"doesn't have a primary key
column and will be ignored"` → 403, even for admin). Adding the real PKs (`id` for all 26, `code` for `languages`)
fixed it; the digest stayed `6cd35278…` (PKs don't change column data_type). **So the structure artifact MUST capture
constraints (at least PKs), not just columns.**
**The serve recipe (what stages 5–6 need beyond stage 4):**
1. The tables exist with prod-faithful columns **+ PRIMARY KEYs** (the structure artifact).
2. A `directus_collections` registration row per served collection.
3. A `directus_permissions` `read` row on Directus's **hardcoded public policy** `abf8a154-5b1c-4a46-ac9c-7300570f4f17`
   (bootstrap creates the policy + its `(role=NULL,user=NULL)` access link), `fields='*'`, `status=published` filter
   for simulations/skill_paths.
   `directus_fields` rows are **NOT required** — Directus introspects DB columns once a collection is registered + has a PK.
**Gate status — demonstrated, not yet automated:** the structure + registry + permissions were applied BY HAND
(`structure.sql` + `iter-05/pks.sql` + `iter-05/serve.sql`). The exit_gate says "**stacksnap** applies the captured
structure", so the remaining deliverable is `STRUCT-M21-codeify`: make stacksnap capture (DDL+PKs+registration+
permissions over `--dsn`) and apply it in provision. That flips the gate from demonstrated → met.

### M21-D10 — privilege-visibility alignment: capture must match the digest's information_schema view (iter-06)
**Finding (capture-side code-ification):** the snapshot staleness digest (`pg.SchemaVersionSQL`) is computed over
`information_schema.columns`, which is **privilege-filtered** — it shows only relations the connecting role can see.
The structure capture must scope to the SAME view, else it captures tables that `pg_class` exposes but the read role
cannot, leaving the applied schema "ahead" of the digest so it never converges. **Verified on prod:**
`sim_tasks_criterion` + `sim_tasks_criterion_check` are `pg_class`-visible to the read user (`marco_read`) but absent
from its `information_schema` — so the cached digest `6cd35278…` counts **26** user collections, while a naive
`pg_class` capture finds **28**. **Decision:** every structure-capture catalog query intersects `pg_catalog` (for exact
types/PKs) with `information_schema.columns WHERE table_schema=$1` (for the digest's exact table set). This also means
a capture run as a DIFFERENT (more-privileged) role would see a different table set → a different digest/cache key —
consistent by construction as long as capture + the eventual replay-probe read the schema the same way. Generalizes to
any digest-keyed capture against a least-privilege read role.

### M21-D11 — capture sequences by DEFAULT-reference, not ownership (iter-07)
**Finding (the integration test surfaced it):** the first `structureSeqSQL` keyed `CREATE SEQUENCE` on the OWNERSHIP
dependency (`pg_depend`: sequence→table, the `OWNED BY` edge). Prod's junction-table sequences ARE owned (8 found), so
a prod capture worked — but a source whose sequences are NOT owned (a hand-built fixture, or in principle any source
where ownership was dropped) records only the `DEFAULT nextval(...)` reference, not ownership. Such sequences were
MISSED → the captured `CREATE TABLE`s referenced non-existent sequences → the apply would fail. **Decision:** key the
sequence capture on the **default-reference** dependency (`pg_attrdef` → sequence `pg_depend` edge), which a
`DEFAULT nextval('seq')` ALWAYS records regardless of ownership. This is the correct semantic (capture exactly the
sequences the table DDL references) and is robust to source setup. Live-validated: 8 on prod (owned) AND 8 on the
standalone-sequence integration source. A restored `pg_dump` preserves ownership too, so prod + cold-start paths are
unaffected; the change only ADDS robustness.

### M21-D12 — auto-provision is gated on the bootstrapped-GAP precondition (iter-07, review-driven)
**Context:** iter-07 wired `stacksnap replay` to AUTO-PROVISION a directus stack's schema (apply the captured
structure) on a cache miss. An adversarial review (the `m21-iter07-review` workflow) caught a regression: the apply
fired on ANY cache miss, not just a bootstrapped GAP. A diverged/schema-skewed target (already has user collections,
digest ≠ captured — e.g. a stack provisioned at release N then the cache re-captured at N+1) would hit the captured
`CREATE TABLE` (intentionally non-idempotent) → collide → raw **exit 1**, REGRESSING the pre-M21 clean **exit 5**
divergence path ("bring the stack to the captured shape first"). The apply itself stays atomic (simple-protocol
implicit transaction → no half-provision), so it's a UX/contract regression, not data corruption. **Decision:**
`tryAutoProvision` probes the target for any user collection (`information_schema` base tables minus the `directus_*`
system set) and ONLY provisions when that count is 0 (a true gap); a diverged target returns a no-op so the caller
falls through to the existing exit-5 message. This makes auto-provision strictly the gap→fill case its comments
already claimed. Live-confirmed: a diverged target (2 user tables) now → exit 5 (not exit 1). **General lesson:** any
auto-provision/auto-heal-on-cache-miss must gate the mutation on the precondition it assumes, or a skewed input
silently degrades a clean error into a raw failure.

### M21-D13 — the SERVE half: firewall structural-metadata admissibility + serve-row capture/apply (iter-08, GATE-COMPLETING)
**Result:** the M21 exit_gate is MET by tooling. `stacksnap` now captures the two serve-row system tables —
`directus_collections` (registration) + `directus_permissions` (public-policy read grant) — and applies them on
auto-provision, so a freshly bootstrapped + stacksnap-provisioned stack serves the captured public catalog to an
ANONYMOUS reader with NO hand SQL. The serve rows are byte-equivalent to iter-05's DEMONSTRATED hand-applied
`serve.sql`; iter-08 produces them BY TOOLING from the same source.
**The firewall structural-metadata admissibility class (the In-scope safety carve-out):** the serve rows live in
`directus_*` SYSTEM tables, outside `AssertPublicOnly`/`AssertPlan` (which govern user-collection ROW captures). A
new gate `firewall.AssertStructuralMetadata` admits a system table as "structure, not tenant data" ONLY if it
carries NONE of `firewall.TenantScopeColumns` (organization_id/tenant_id/private/user/owner/user_created/
user_updated). This EXTENDS the firewall (admits a previously-inadmissible class under a strict, explicit predicate)
without LOOSENING it (any tenant-scope column → reject; `AssertPlan`/`AssertCaptured` unchanged). Confirmed (sanctioned
prod read + the live fixture): `directus_collections`/`directus_permissions` carry zero tenant-scope columns →
admissible. `directus_access` is EXCLUDED (it has a `user` uuid column) and `directus_policies` is not captured —
both are bootstrap-provided (the hardcoded public policy `abf8a154` + its `(role=NULL,user=NULL)` access link exist
on any fresh bootstrap), so the anonymous access chain is already present on the target.
**Apply mechanism = additive INSERTs appended to the structure artifact** (NOT the TRUNCATE-COPY row replay, which
would wipe bootstrap's system rows). `directus_collections`: all columns, `ON CONFLICT (collection) DO NOTHING`
(idempotent + skips any bootstrap-provided registration). `directus_permissions`: OMIT the serial `id` (let it
auto-gen — a captured prod id would collide with bootstrap's own system-permission serials), plain INSERT. Both
rendered faithfully from the source via ONE server-side query each (`jsonb_each_text(to_jsonb(t))` joined to the
ordered column list + `quote_nullable`) so the column set is discovered dynamically (version-robust) and every value
round-trips. The serve rows ride the existing `Surface.CapturesStructure` capture path + the iter-07 `tryAutoProvision`
`ExecScript` apply path with ZERO apply-side code change (they ARE part of the structure SQL).
**Validation (iter-08, live):** the render SQL produces valid faithful INSERTs (apostrophe-escaped, served-only,
public-policy-only, id-omitted); applying them into a bootstrapped-shape target is idempotent + does not collide with
bootstrap's system rows; the full Go path `CaptureStructure(*pg.Conn)` → `ExecScript` into a fresh GAP target applies
cleanly (6 statements: 2 tables + 2 PKs + 2 serve INSERTs; collections created with PKs, registered, public-read
granted). The literal HTTP boot+serve against the REAL 26-collection prod structure was DEMONSTRATED in iter-05 with
these same rows; iter-08 rides that proven path. **General lesson:** a firewall carve-out for a new admissible class
should run ASSERT-THEN-READ (admissibility on the introspected column set BEFORE the capture read) so an unexpected
tenant column aborts before any row is materialized.
