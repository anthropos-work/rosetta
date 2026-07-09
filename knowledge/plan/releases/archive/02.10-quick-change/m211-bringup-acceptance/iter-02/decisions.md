# iter-02 — Decisions (tik under TOK-01)

## D1 — The cache-migration drift is empirically reconcilable (NOT a user-blocker)
Merged `public.skills` has 16 columns vs the cache's 15; the delta is exactly one added `ts_search`
(`GENERATED ALWAYS AS (...) STORED` tsvector). Sweeping all 10 tables, the only merged-but-not-cached
columns are the `ts_search` generateds (skills, job_roles, skill_translations, job_role_translations).
Every NOT-NULL-no-default column is present in the capture. A COPY MUST omit generated columns, and the
cached manifest column lists (which predate `ts_search`) already omit it → the replay's explicit-column
COPY loads the real 15/10 columns and PG auto-computes `ts_search`. Verified post-replay: `ts_search`
filled on all 42,790 skills. So the milestone's pre-surfaced "if column drift is unreconcilable →
user-blocker" branch did NOT fire — the drift is additive + self-healing.

## D2 — Zero-copy hardlink migration (bytes unchanged → checksums hold)
Migrated the cache by HARDLINKING the 10 payloads under `public.X.copy` names rather than copying
(~1.4 GB) or moving. Consequences: (a) zero extra disk (link count 2, shared inode); (b) the old
`c75ce94…` cache stays intact as a fallback; (c) the payload BYTES are unchanged, so the manifest's
per-table sha256 checksums remain valid — the replay's pre-write checksum verification passes without
re-hashing. Only the manifest LABELS (schema/payload/filter/public_via/schema_version) changed.

## D3 — Key the migrated cache under the REAL live-probed merged digest (`5afc0bcc…`), not an arbitrary key
The replay probes `conn.SchemaVersion(ctx, "public", VersionTables())` from the target stack's live schema
and requires `manifest.schema_version == that probe` (IsStaleFor). A cold `/dev-up` set-dress runs replay
WITHOUT a `--schema-version` override, so it will probe the same merged schema → `5afc0bcc…`. Keying the
migrated cache under exactly this probed digest makes the cold replay HIT automatically (no override flag,
no per-stack fixup). The digest is deterministic across reset-db + re-migrate (same DDL), so it survives a
cold rebuild. (Rejected the `--schema-version` override shortcut precisely because cold /dev-up won't use it.)

## D4 — Re-pin consumption now (not deferred to the cold tik)
`.agentspace/rext.tag` was v1.10.1 (stale, pre-merge). Re-pinned to `quick-change-m209` (rext HEAD
2f06e78, the re-grounded tooling) now — it's a planned TOK-01 prep step, one gitignored line, zero-risk,
and de-risks the cold /dev-up tik (which consumes rext at this tag). `git status` clean confirms
`.agentspace` is gitignored → no rosetta tree change. The final `v2.1` rext roll remains close-release's job.
