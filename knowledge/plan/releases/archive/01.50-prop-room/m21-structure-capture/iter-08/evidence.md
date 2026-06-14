# M21 iter-08 — live evidence (serve-row capture + apply)

Environment: Docker up; `directus/directus:11.6.1` + `postgres:16-alpine` cached. Prod/Tailscale NOT reachable
this session (postgres MCP timed out; `~/.pgpass` absent) — so the live validation used a throwaway
`postgres:16-alpine` with a directus-shaped fixture (edge cases below), exercising the EXACT render/apply SQL +
the real Go code path. The HTTP boot+serve against the real 26-collection prod structure was DEMONSTRATED in
iter-05 with the same serve rows; not re-run here (environment).

## Fixture (edge cases)
- `directus_collections` + `directus_permissions` with representative 11.6.x columns.
- Registration rows for `simulations`, `skill_paths`, `roles` (served) + `internal_only` (NOT served).
- `roles.note` = `role's note with apostrophe` (quote-escape test).
- Permissions: 3 served read rows on the public policy `abf8a154…`; 1 row on a DIFFERENT policy (must skip);
  1 row for `internal_only` on the public policy (must skip — not a served collection).

## 1. Render-SQL faithfulness (the dynamic queries)
`serveCollectionsRowsSQL` → one INSERT with the full 20-column list discovered dynamically; only the 3 served
collections; `'role''s note with apostrophe'` correctly escaped; NULLs as `NULL`, bools as `'false'`/`'true'`;
`ON CONFLICT (collection) DO NOTHING`. `internal_only` excluded.
`servePermissionsRowsSQL` → one INSERT WITHOUT the serial `id` column; only the 3 served public-policy read rows;
the `status='published'` filter JSON preserved; the non-public-policy row + `internal_only` row excluded.

## 2. Apply idempotency + no system-row collision
Applied the generated INSERTs into a bootstrapped-shape target (`tgt` schema: system tables + 2 bootstrap system
perm rows + a pre-existing `simulations` registration):
- collections: `INSERT 0 2` (the pre-existing `simulations` registration skipped by `ON CONFLICT`).
- permissions: `INSERT 0 3` (serial ids auto-genned — no collision with bootstrap's serials).
- re-run collections: `INSERT 0 0` (idempotent).
- final: 3 served registrations (no dupes), 3 served public-read perms, 5 total perms (2 bootstrap + 3 served).

## 3. Full Go capture→apply path
`directus.CaptureStructure(ctx, *pg.Conn)` against the fixture → `=== CaptureStructure OK: 6 statements ===`
(2 CREATE TABLE + 2 ADD PRIMARY KEY + 2 serve INSERTs), serve rows ordered AFTER the schema, admissibility passed.
`pg.ExecScript` (the real `tryAutoProvision` apply path) into a fresh GAP target (`tgt2` DB: directus_* system
tables only) → `=== APPLY OK (6 statements applied) ===`. Verified in `tgt2`:
- user collections created: 2 (`simulations`, `roles`) — `simulations has PK: true`.
- collections registered: 3; served public-read perms: 3.

This is exactly the production chain: capture (admissibility-gated) → manifest `Structure` artifact → replay
`tryAutoProvision` `ExecScript` on a bootstrapped GAP. The serve rows ride the iter-07-automated apply path
unchanged; iter-08 only added WHAT is captured (the two serve-row system tables) + the firewall carve-out that
admits them.

(All throwaway containers + the temp Go harness were removed after the run.)
