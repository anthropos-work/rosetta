# M261 — Progress

**Status: COMPLETE.** The canon is loaded, captured, and in the snapshot store. 2026-08-15.

## What was done

The milestone was scoped as *"re-capture from a safe prod source"*. That framing was wrong, and the
right route needs **no production access at all**: the canon is a **checked-in artifact** in `app`
(`taxonomy-canon/`, 13 files) and `cmd/taxonomy-load` takes only a `DB_CONNECTION` DSN. So the canon
was loaded directly, from the source of truth, into an isolated database built from the stack's own
Postgres image — **`demo-4` was not touched** (it is flagged do-not-reset and its app pin predates the
v2 migrations anyway).

1. Isolated Postgres from the stack's own `demo-4-postgresql` image on `:55432`.
2. `atlas migrate apply --env local` — **180 migrations**. One stop en route: `20260518125439` needs an
   `extensions` schema (`vector`/`pgcrypto`/`pg_trgm`) that the platform's own init creates and a bare
   container does not. Created it and resumed; **all 180 applied**.
3. `taxonomy-load -dry-run`, then the real load.

## Measured after the load

| table | rows |
|---|---:|
| `public.skills` | **3,562** |
| `public.job_roles` | **706** |
| `public.specializations` | 283 |
| `public.categories` | 25 |
| `public.skill_redirects` | **12,835** |
| `public.job_role_redirects` | **11,182** |
| `public.job_role_skills` | 14,106 |
| `public.skill_translations` | 7,124 (EN+IT over 3,562) |

**Redirect resolution works**: every redirect joins to a live skill, **0 dangling**, and
`K-INTPRE-0A43` — the sample quoted in `taxonomy-canon.md` — resolves.

## Cross-validation: three independent sources agree

M259 measured the **bundle files**. The loader read the same bundle and reported
`skills=3562 roles=706 profiles=15005 redirects=24017 drops=37207`, and the database now holds
12,835 + 11,182 = **24,017** redirects and 26,518 + 10,689 = **37,207** drops. Bundle → loader → DB,
all three agreeing.

**M259's disputed figure is confirmed by the platform's own tool.** M259 derived **61,224** total
retirements and flagged an 8-entry difference against the commit message's 61,216, attributing it to a
bundle regeneration. The loader's dry run reports `retired (soft delete) … 61224`. The derivation was
right and so was the explanation.

## Two things this surfaced that no downstream milestone had accounted for

1. **⚠️ The table names are PLURAL.** `skill_redirects`, `job_role_redirects`, `category_translations`,
   `specialization_translations` — not the singular names D-M260-4 routed here, which came from the ent
   schema *filenames* rather than from a database. Corrected here so M262/M263 do not query a table
   that does not exist.
2. **⚠️ `taxonomy_canon_states` does NOT exist.** It appears in the taxonomy-v2 commit range and in no
   migration; the loaded schema has no such table. M260 flagged this as "confirm rather than declare
   blind" — confirmed: **do not add it to the capture surface.**

## The caveat that matters for a demo

> `vectors not computed: the canon is loaded but does not take part in matching until this is re-run`
> — `reason="no embedding manager configured"`

The load is **structurally complete and semantically inert for matching**: no AI key was supplied, so
`skill_embeddings` were not computed. Every taxonomy surface that browses, lists or joins works; **AI
skill-matching against the new canon does not** until the load is re-run with an embedding manager.
This is a real gate on M262 (the seed's verified-skill chain) and M263 (the taxonomy page browses,
so it is unaffected).

## Production state — UNRESOLVED, and no longer asserted either way

An earlier note in this file claimed the canon "has never been loaded into production". **That was
read off a plan document, not measured, and it is withdrawn.** What is actually known:

- `app/knowledge/taxonomy-canon-migration.md` @ `4bccda085` *states* prod's last applied migration is
  `20260804160000` — and the repo carries **8 migrations newer than that**, all taxonomy-v2. That is
  evidence about what EXISTS to be applied, not about what IS applied.
- Prod's API is reachable (`api.anthropos.work/api/health` → 200) but `taxonomyCategories` returns
  `unknown viewer: Forbidden`, and **every** authenticated route 404s unauthenticated, so no HTTP
  signal distinguishes "not deployed" from "not logged in".
- Neither `db-query` path is available in this session: **no MCP servers are configured**
  (`mcpServers: []`), and there is no `psql` or `~/.pgpass`.

**It does not block this milestone** — the canon came from the checked-in bundle, which is the source
production itself loads from.


## The capture surface — 10 tables → 14, measured not guessed

Every column was read from a database with all 180 migrations applied and the canon loaded:

| table | scope |
|---|---|
| `skill_redirects` | parent-scoped via `skills` (the DESTINATION; `old_node_id` is a bare historical string with no FK, by the platform's design) |
| `job_role_redirects` | parent-scoped via `job_roles` |
| `category_translations` | parent-scoped via `categories` |
| `specialization_translations` | parent-scoped via `specializations` |

`ts_search` is excluded on both translation tables for the same reason it already was on the two
existing ones: it is a generated tsvector. Four pinning tests were updated deliberately (surface set,
parent scopes, `VersionTables` digest, and the stacksnap registry) — they were doing their job.

## Captured, and M260 proved itself live

`stacksnap capture --surface taxonomy` against the loaded database, into the REAL store:

```
capture: public.skills captured 3562 public rows against 42790 in the newest prior capture
(8% — below the 50% shrink floor). This is EITHER a broken source OR a real consolidation, and
the capture cannot tell them apart; only you can.
```

That is M260's shrink rung firing on a genuine 42,790 → 3,562 collapse, against real history, and
**refusing to guess**. Re-run with `--accept-shrink "<reason>"` (the flag this milestone wired to the
CLI — the Options field existed but nothing exposed it), it captured **14 tables / 50,848 rows /
public-only=true**, and the manifest records:

> `SHRINK ACCEPTED (taxonomy v2 canon — 43,584 skills consolidated to 3,562 …): public.skills 42790->3562`

## ⚠️ The cache PURGE was wrong, and is cancelled

M261's scope said *"PURGE rather than refresh the snapshot cache — node-ids moved, so a stale
artifact is not merely old, it is wrong."* **Measured, that is false, and the purge must not run.**
The store keys refs by `(surface, schemaVersion)`, so the canon captured as a NEW ref and the old
snapshots are untouched:

```
5afc0bccf1df   tables=10   skills=42790     (old, public schema)
c75ce94d6a80   tables=10   skills=42790     (old, skiller schema)
8b19b5b0a8ad   tables=14   skills=3562      (the canon)   ← new
```

A stack built from the old `app` resolves the old schema version and still gets a valid taxonomy; a
stack from the new `app` resolves the new one and gets the canon. **Purging would have destroyed a
working artifact to solve a collision that the store's own design prevents.**
