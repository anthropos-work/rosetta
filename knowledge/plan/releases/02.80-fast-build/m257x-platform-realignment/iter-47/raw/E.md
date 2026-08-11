# AUDITOR E — 6 files / 1459 lines

**Positive control:** all 6 read to final line; counts match `wc -l`
(service_taxonomy 440 · hiring 378 · shared_libraries 242 · storage 175 · askengine 121 · dependency_map 103).

## BLOCKERS — 0

Every actionable claim checked held; every load-bearing anchor resolved to the construct named.

**All four iter-46 repairs in this hand verify independently:**

- **Tier-2 deployment** (`service_taxonomy.md:136-141`) — `docker-compose.yml:311` `studio-desk:`, `:342`
  `profiles: [studio-desk, all]`. Cross-anchors `frontend_architecture.md:11` + `studio-desk.md:21` agree.
  **No over-correction:** Ant Academy really is absent from compose and from `repos.yml` (exactly 9 entries).
- **Studio-Desk technology** (`:150-153`) — `package.json` has 0 react/vue/angular (express 4.18.2,
  vite 6.4.2, typescript 5.3.0); `find -name '*.tsx' -o -name '*.jsx'` returns nothing.
- **The Directus retraction-of-the-retraction** (`:290-305`) — `git show a2a3ee6^:docker-compose.yml` →
  `:384 image: directus/directus:10.10.1`, `:386 - 8055:8055`, `:409 - ADMIN_PASSWORD=password`, exactly as
  claimed; `platform-migration-status.md:86` agrees. The *current* half holds too: `grep -n DIRECTUS
  docker-compose.yml` returns only `:164/:165`, inside the `cms` block.
- **The `ai` correction** (`shared_libraries.md:118-128`) — mechanism matches
  `app/internal/skillerai/ai.go:336-359,161-182` and `app/internal/jobsimulation/ai/ai.go:263-276,87-92`.
  It does **not** contradict `service_taxonomy.md:110`'s "no ladder", and the
  `#routing-what-is-actually-implemented` anchor resolves.

Spot-verified clean beyond that: all 8 compose service/port/profile rows + both profile tables;
`repos.yml:10-19`; the **5→4→3→1 subgraph ladder** (`749dc86`/`7c17e63`/`915da06`, incl. the single commit
deleting both `cms.graphqls` and `jobsimulation.graphqls`); router port 8080 everywhere; all 5
shared-library version-pin tables against 7 `go.mod`s; `app/main.go`'s six Connect handlers with **no**
skillpath/roadrunner handler; every hiring `is_hiring` anchor; the five-member `completion_status` enum +
the no-CHECK/no-rejection chain; the whole of `askengine.md`; the whole of `storage.md`'s code map.

## MINORS — 7

| # | site | what is off |
|---|---|---|
| 1 | storage.md:158 | `ENVIRONMENT \| (empty)` under a **"Compose value"** column; `docker-compose.yml:206` sets `ENVIRONMENT=development`. Behaviourally inert (colony inits Sentry only at `production`) |
| 2 | hiring.md:189-196 | "Minimal write-set" omits `token` (NOT NULL + UNIQUE, no default) — a raw-SQL seeder from the list fails the INSERT, **loudly**. Same row calls `started_at`/`ended_at` "Non-null"; DDL says `timestamptz NULL` |
| 3 | hiring.md:159 | cites `intelligence.go:1728-1735` for the `row_number()` window; that range is the **call site** (`:1733`) — the window function is at `:2158-2169`. The adjacent `:1738-1751` cite is exact |
| 4 | hiring.md:21,146-148 | "back-fills it into the canonical entity" — the migration re-points FK **link ids** and drops the mirrors; no score/data copy (the score was already on `job_simulation_sessions`). Drop anchor + conclusion both correct |
| 5 | shared_libraries.md:126 | *"each consumer's own `internal/ai/ai.go` wrapper"* — there is **no `app/internal/ai/`**; in `app` the wrappers are `internal/skillerai/ai.go` and `internal/jobsimulation/ai/ai.go`. The literal path resolves only in the frozen `jobsimulation` husk. `external_services.md:539` names the correct path. **NB: adjacent to an iter-46 edit — see the ledger's induced-defect analysis** |
| 6 | service_taxonomy.md:66 + dependency_map.md:18 | roadrunner "Orphaned — nothing calls it": the still-started `jobsimulation` husk carries `ROADRUNNER_RPC_ADDR=http://roadrunner:10401` (`docker-compose.yml:118`). Husk is inert so no LIVE caller, but the edge is *configured*, not absent |
| 7 | dependency_map.md:50 | taxonomy consumers omit the `cms` and `jobsimulation` husk repos, which require it **directly** (`cms/go.mod:13`, `jobsimulation/go.mod:15`) — as `shared_libraries.md:181` already states. Cross-doc undercount |

## Files read clean

- **`askengine.md` — fully clean, 0 blockers and 0 minors.** Every const, route, table, test count, model
  id and version date reproduces against source.
- `service_taxonomy.md` · `storage.md` · `hiring.md` · `shared_libraries.md` · `dependency_map.md` — clean
  of blockers.

## Explicitly UNVERIFIED (not graded — uncheckable from this tree, and named rather than passed)

The internal package/symbol inventories of `colony`, `proto`, `ai`, `authn`, `taxonomy` (private Go
modules, not cloned — only their `go.mod` pins and import edges were checkable, and those all passed);
prod-terraform assertions in the uncloned `infrastructure` repo; GitHub archive states; and the taxonomy
row/skill counts, which are snapshot-provenanced and date-fenced.
