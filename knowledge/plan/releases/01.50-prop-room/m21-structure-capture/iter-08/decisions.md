# M21 iter-08 — decisions (iter-local)

_(Cross-iter decisions promote to the milestone-root `decisions.md` as M21-DNN.)_

- **iter-08-L1 — the serve rows ride the structure artifact, not a new manifest field.** Appending the
  registration + permission INSERTs to the structure SQL (after the schema) means the existing `Surface.
  CapturesStructure` capture path + the existing `tryAutoProvision` `ExecScript` apply path carry them with ZERO
  apply-side code change. The serve rows ARE part of `manifest.Structure` — no format bump, no second artifact.
- **iter-08-L2 — assert-then-read for the firewall carve-out.** `assertServeTablesAdmissible` runs the
  structural-metadata admissibility check on the serve-row system tables' column sets BEFORE any serve row is read.
  A registry table that unexpectedly carried a tenant column aborts the capture before a single row is materialized
  (extend, never loosen — proven by `TestCaptureServeRows_AdmissibilityRejectsTenantColumn`: serveRead stays false).
- **iter-08-L3 — faithful multi-row INSERT via one server-side render.** The directus_* system-table shape is
  version-dependent, so the render queries discover the column set dynamically (information_schema.columns) and build
  each row's value tuple with `quote_nullable(jsonb_each_text(to_jsonb(t)).value)` joined back to the ordered column
  list. One round trip, every value round-trips (apostrophes/JSON/NULL/bool), no hardcoded column list. Permissions
  exclude `id` (`to_jsonb(t) - 'id'`) so the serial auto-gens.
- **iter-08-L4 — countStatements ignores comment-line semicolons.** A serve-header line carried a `;`, which the
  raw `strings.Count(script, ";")` provenance probe counted as a statement. Made `countStatements` skip `--` comment
  lines so the count is robust to header wording (the count is provenance-only, but a silent off-by-one is avoidable).
