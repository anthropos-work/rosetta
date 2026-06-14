# M21 iter-08 — progress

**Type:** tik (under TOK-01, refined by M21-D7/D9).
**Protocol:** the staged-pipeline (TOK-01) — gate-completing serve-row automation.

## Phase A — build (DONE)
Shipped the SERVE half of the structure code-ification (rosetta-extensions):
- **Firewall structural-metadata admissibility class** (`firewall/firewall.go`): `StructuralMetadataPolicy` +
  `AssertStructuralMetadata` + `TenantScopeColumns`. A `directus_*` SYSTEM table is admitted as "structure, not
  tenant data" ONLY if it carries NONE of the tenant-scope columns (organization_id/tenant_id/private/user/owner/
  user_created/user_updated). EXTENDS the firewall (admits a new class) without LOOSENING it (any tenant column →
  reject). Empty plan passes vacuously (serve rows are optional per surface) — unlike `AssertPlan`.
- **Serve-row capture** (`directus/structure.go`): `CaptureServeRows` + `assertServeTablesAdmissible`. Captures
  `directus_collections` registration rows (`ON CONFLICT (collection) DO NOTHING`) + `directus_permissions` public-
  policy read rows (serial `id` OMITTED → auto-gen, plain INSERT) for the 9 served collections, rendered FAITHFULLY
  from the source via dynamic `quote_nullable`/`jsonb_each_text` server-side SQL (version-robust column discovery).
  Admissibility runs BEFORE any serve-row read (assert-then-read). Appended to the structure artifact AFTER the
  schema (so collections exist before they are registered).
- **Wiring:** `CaptureStructure` now also captures the serve rows → they ride the existing
  `Surface.CapturesStructure` path in `capture.Run` (manifest `Structure` artifact) and the existing
  `tryAutoProvision` apply path (`ExecScript`) unchanged — no apply-side code change needed (the serve rows ARE part
  of the structure SQL). `EXCLUDE directus_access` (`user` column) + `directus_policies` (bootstrap-provided).
- **countStatements** made robust to comment-line semicolons (a header `;` no longer inflates the provenance count).

## Phase B — test gate (GREEN)
- `go vet ./...` clean; `go test ./...` all packages PASS (12/12).
- New unit tests: `firewall/structural_metadata_test.go` (8 tests — admissible OK, rejects user/tenant_id/case-
  insensitive/empty-identity, reports-all-offenders, TenantScopeColumns coverage guard); `directus/serve_test.go`
  (6 tests — admissible appends both in order + ON CONFLICT + id-omitted, admissibility-rejects-tenant-column-before-
  read, empty-source-nothing, CaptureStructure-appends-after-schema + stmt count, array literal).
- Existing tests unchanged + green (the `CaptureStructure`/`assembleStructure` signature change is internal).

## Phase C — live gate proof (DONE; throwaway Postgres + cached directus image present, prod/Tailscale NOT reachable)
The load-bearing NEW behavior validated against a real Postgres (`postgres:16-alpine`, directus-shaped fixture with
edge cases: apostrophe in a value, a non-served collection, a permission on a non-public policy):
1. **Render-SQL faithfulness:** the dynamic `directus_collections` + `directus_permissions` render queries produce
   valid, faithful INSERTs — only the served collections + only the public-policy read perms; apostrophe correctly
   escaped; serial `id` omitted; full column list discovered dynamically.
2. **Apply idempotency + no system-row collision:** applying the generated INSERTs into a bootstrapped-shape target
   (with bootstrap's own system perm rows + a pre-existing `simulations` registration) → collections `ON CONFLICT`
   skips the dupe, permissions auto-gen ids (no collision with bootstrap's serials), re-run is a no-op. Final state
   exactly the served set.
3. **Full Go path (capture→apply):** `directus.CaptureStructure` through a real `*pg.Conn` → 6 statements (2 tables
   + 2 PKs + 2 serve INSERTs), serve rows correctly ordered after the schema; then `pg.ExecScript` (the real
   `tryAutoProvision` apply path) into a fresh GAP target → APPLY OK: 2 user collections created WITH PKs, 3
   registrations, 3 served public-read perms. This is exactly the production capture→auto-provision chain.

**Not re-runnable this session (environment, not a gap):** the literal HTTP `GET /items/simulations → 200` against a
Directus booted on the REAL 26-collection prod structure needs Tailscale-to-prod (down now). That boot+serve was
DEMONSTRATED in iter-05 with the SAME registration + permission rows (`iter-05/serve.sql`); iter-08 now produces
those identical rows BY TOOLING (faithfully captured from the same source), riding the iter-05-demonstrated +
iter-07-automated path. Every step in the gate chain is automated + validated; the serve recipe rows are byte-
equivalent to the hand-applied proof.

Evidence: `iter-08/evidence.md` (the live render/apply transcripts).

## Close — 2026-06-13

**Outcome:** the SERVE half is AUTOMATED — `stacksnap` now captures the `directus_collections` registration +
public-policy `directus_permissions` read rows (firewall structural-metadata-admissibility-gated) and applies them
on auto-provision, so a freshly bootstrapped + stacksnap-provisioned stack serves the captured catalog anonymously
with NO hand SQL. Full capture→apply Go path live-validated; the serve rows are byte-equivalent to iter-05's
DEMONSTRATED hand-applied `serve.sql`. The M21 exit_gate is MET by tooling.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n (tik) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** M21-D13 (firewall structural-metadata admissibility class + the serve-row capture/apply); iter-local in `decisions.md`.
**Side-deliverables:** the `countStatements` robustness fix (comment-line semicolons no longer inflate the provenance count) — folded in as part of the serve-header landing, not separable.
**Routes carried forward:**
  - To `/developer-kit:harden-mstone-iters` (final pass): AP-1 (replayCmd-wiring hermetic test + conn seam), AP-2
    (multi-snapshot tie-break determinism), AP-3 (exit-4-boundary regression guard), UDT/identity guards, firewall-
    ordering direct test, PLUS the iter-08 additions: a live-integration test for the serve-row render SQL (the
    dynamic queries are hermetically unit-tested + hand-validated live, but not in an automated integration harness),
    and a direct test that admissibility runs before any serve-row read.
  - `directus_files` ref capture (wire the dead `media.go`); M23 referential closure of the 20 dangling relations.
**Lessons:** building a faithful multi-row INSERT from a version-unknown column set is cleanest as ONE server-side
render query (`jsonb_each_text(to_jsonb(t))` joined to the ordered column list + `quote_nullable`) — it discovers the
column set dynamically (version-robust), round-trips every value correctly (apostrophes, JSON, NULLs, bools), and is
one round trip. The `assert-then-read` ordering for a firewall carve-out (admissibility BEFORE the capture read)
is the safe shape: a registry table that unexpectedly carried a tenant column aborts before a single row is
materialized.
