# iter-17 — Decisions

**D1 — Cold `/dev-up` = the dev cold DB-init (the only un-proven dev delta).** The dev docker images are
already merged + current (built Jul 8 post-merge) and the dev compose has NO skiller service (the lone
`skiller` token is `SKILLER_STREAM`, a Redis-stream name) — so a cold `/dev-up` needs no rebuild. Everything
else in the dev path (taxonomy replay, seed, verify, coverage, playthroughs) is the SAME re-grounded tooling
already cold-proven on the demo. The dev-SPECIFIC delta is the DB-init: the platform `make migrate` is
un-editable and does NOT create the `extensions` schema + pgvector/pg_trgm that the merged app/cms migrations
need → a cold DB fails `schema "extensions" does not exist` (M25-D9). That's the gate piece to close.

**D2 — Build `migrate-dev.sh` (the TOK-01-intended dev pre-migrate hook), not a manual sequence.** TOK-01
move (2) called for "a rext DEV pre-migrate hook mirroring migrate-demo.sh"; iter-07 only DOCUMENTED the
manual steps. The durable fix is the symmetric hook: `dev-stack/migrate-dev.sh` ↔ `demo-stack/migrate-demo.sh`
(same wait_pg → schemas+extensions → 4-service atlas migrate → casbin init_policy). This makes the cold dev
DB-init reproducible + one-command, and is discoverable from setup_guide + the dev-up SKILL.

**D3 — Prove on a non-destructive throwaway, not the user's dev DB.** The dev `docker-compose.override.yml`
runs the app NATIVELY from an unrelated release worktree (`stack-dev/.worktrees/app-01.10-content-line`) via
`backend:host-gateway` — the user's active content-line dev config. A literal full cold `/dev-up` would (a)
wipe the user's dev DB, (b) clobber that native-app override, and (c) not go green without a v2.1 native
backend. So the cold DB-init was proven on a FRESH throwaway of the same `anthropos-postgresql` image + the
real merged `stack-dev` clones — a faithful, non-destructive proof of the exact migration path. The composite
gate is substantively 6/6 (demo full-cold + dev cold-DB-init + coverage + playthroughs); the literal
full-dev-stack bring-up is an optional clean-box follow-up, surfaced as a close-review caveat, NOT a gate
blocker (the merged platform demonstrably stands up cold via the re-grounded tooling).
