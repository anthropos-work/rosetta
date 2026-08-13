# Seat E — M257x iter-49 KB-fidelity audit (ninth clause-5 reading)

## Provenance

| Repo / clone | SHA |
|---|---|
| `rosetta` (branch `m257x/platform-realignment`) | `2fc633a2c5c09a6034e5ab4e29d509dfcadcbd8a` |
| `stack-demo/app` | `5ba1704482cf812b130c2d3673afd09f4f7f22e5` |
| `stack-demo/platform` | `2adcf714bd877a205e8948f59a23db49b884c054` |
| `stack-demo/sentinel` | `88bc55929dde7ba43913966ec3fc36372e4ff32a` |
| `stack-demo/storage` | `4ce8ece52adb7c095e792e235da4a8913214d190` |
| `stack-demo/messenger` | `fa47850d9c507d1928da7a38f7b37bac1bb8fabc` |
| `stack-demo/cms` | `ca50c8170fefe1122d680efe54f7e56798a79d82` |
| `stack-demo/graphql-wundergraph` | `60c229f` (HEAD) — commit `915da06` inspected |
| `stack-demo/next-web-app` | `bb3313bc0133ee5728ce83fda485e95bfea1a6c6` |
| `stack-demo/studio-desk`, `stack-demo/ant-academy` | working clones (consulted for clerk-integration.md) |
| `.agentspace/rosetta-extensions` | `4d03b53a5e524e9abb020c1a4534ec968c25072b` |

## Coverage

| # | file | `wc -l` | lines read |
|---|------|--------:|-----------:|
| 1 | `corpus/services/hiring.md` | 398 | 398 (all) |
| 2 | `corpus/architecture/dependency_map.md` | 103 | 103 (all) |
| 3 | `corpus/services/cms.md` | 254 | 254 (all) |
| 4 | `corpus/services/storage.md` | 175 | 175 (all) |
| 5 | `corpus/services/sentinel.md` | 166 | 166 (all) |
| 6 | `corpus/services/messenger.md` | 128 | 128 (all) |
| 7 | `corpus/services/clerk-integration.md` | 128 | 128 (all) |
| 8 | `corpus/services/customerio-sync.md` | 75 | 75 (all) |
| 9 | `corpus/services/db-backup.md` | 31 | 31 (all) |
| | **total** | **1458** | **1458** |

Each file was read end-to-end with `Read` (no offset/limit), so every line above was in context.

---

## BLOCKERS

| # | site (file:line) | the false claim | what is true (with platform / rext `file:line`) |
|---|---|---|---|
| **B1** | `corpus/services/hiring.md:66-68` | `CreateOrganizationSimInvitationLink` is *"the very call the `HiringConfigSeeder` uses to write the 5 positions"* | The `HiringConfigSeeder` **never calls that manager**. It writes raw rows: `.agentspace/rosetta-extensions/stack-seeding/seeders/hiring_config.go:99` — `c.CopyRowsIdempotent(ctx, "public", "organization_sim_invitation_links", cols, rows, "id")`, with the column list built inline at `:98`. There is no `SimInvitationLinkManager`, no RPC and no GraphQL call anywhere in that seeder, and it never reads `organizations.is_hiring`; its only gate is the **blueprint** narrative predicate `st.IsHiringOrg()` at `hiring_config.go:65`. (The manager path being described lives at `stack-demo/app/internal/organization/siminvitationlink.go:56-64`.) |
| **B2** | `corpus/services/hiring.md:113-115` | Clerk-only (metadata `true`, column `false`) ⇒ *"`CreateOrganizationSimInvitationLink` hard-errors `\"organization is not hiring\"` … so the `HiringConfigSeeder` cannot write the 5 positions in the first place."* | The consequence is false and mis-routes an empty-positions debug. With `public.organizations.is_hiring = false` the seeder **still writes all 5 position rows** — it bypasses the hard-erroring manager entirely (`hiring_config.go:99`, same evidence as B1) and gates only on the blueprint narrative (`hiring_config.go:65`). The hard error at `app/internal/organization/siminvitationlink.go:63` is real, but nothing in the seed path reaches it. So the Clerk-only failure mode is **not** "no positions"; the actually-broken server-side half is only the content-library type-set (`app/internal/web/backend/graphql/graph/resolver_cms_queries.go:99-103`, verified) and `manager.go:485-487`. |
| **B3** | `corpus/services/hiring.md:293-296` (§ Local development) | In a paragraph scoped *"To make a hiring org's comparison scoreboard render on a **demo/dev stack**"*: *"`jobsimulation.sessions` still exists, frozen and unwritten, until M710"* | On any locally-built stack there is **no `jobsimulation` schema at all**. (a) `stack-demo/platform/repos.yml:17-19` lists `jobsimulation` with `migrations: false` and **no `schema:` key** — only `app` declares `schema: public` (`repos.yml:10-13`). (b) The **only** `CREATE SCHEMA` in app's entire migration set is `auth` (`app/terraform/migrations/20230817154747_supabase_baseline.sql:2`; exhaustive grep over `terraform/migrations/`, rc=0 with 1 hit). (c) Nothing in `platform/Makefile` or `platform/postgresql/` creates it. (d) rext's own ground truth says so verbatim: `.agentspace/rosetta-extensions/stack-seeding/seeders/persona_write.go:58-61` — *"AND THE `jobsimulation` SCHEMA IS GONE TOO … platform `repos.yml` @ 236771f103 stopped declaring the schema, so **a fresh stack never creates it**."* The M710-survival fact (`app/internal/askengine/registry.go:192`) is **production-scoped**; this doc states it unqualified and then applies it to a demo/dev stack. Same unqualified statement recurs at `hiring.md:33-34` and `hiring.md:157-158`. |

---

## MINORS

1. **`hiring.md:68` and `:114`** — anchor off by one. The hard error is `siminvitationlink.go:**63**` (`return nil, fmt.Errorf("organization is not hiring")`); `:62` is the `if !org.IsHiring {` guard.
2. **`hiring.md:188`** — `persona_write.go:69-71` is cited for the three `validation_*` writes. Those lines are the tail of a *comment* block; the mapping lines are `:66-68` and the actual write steps are `persona_write.go:92-94`.
3. **`hiring.md:205`** — `app/internal/data/ent/enum/jobsimulation.go:29-35` for the 5 completion-status members. The `const (` opens at `:28`, the five members are `:29-33`, `)` at `:34`. (`Values()` at `:37-43` is exact.)
4. **`hiring.md:169`** (read-path table row 5) — `intelligence.go:1728-1735` is cited for *"row_number() ORDER BY score DESC per candidate"*. `:1733` is the **call** to `usersBestOrFirstJobSimulationSession`; the actual `row_number()` window SQL is in that function at `intelligence.go:2124-2201` (`:2169`, `:2190`). Rows 3/4/6/7 and `:1820`/`:1844`/`:1846`/`:885-886` all verified exact.
5. **`hiring.md:166`** (read-path table row 1) — `simulationScoreColumn.tsx:54` is `accessorKey,`; the `row.score` render is `:95-99`. Also: the row names `apps/web/...` while § *The render path* (`:317-341`) establishes the reachable surface is `apps/hiring`. Both files exist and are byte-identical (168 lines each), so the claim is not false — but the table is the pre-M224 trace and reads inconsistently with the correction below it.
6. **`dependency_map.md:19` (the freshly-edited Storage row)** — the substantive claim is **verified true**: storage's compose env has no `DB_CONNECTION`/`REDIS_ADDR` (`docker-compose.yml:203-210`), `storage/go.mod` has no redis (grep rc=1, ran clean), `storage/internal/storage/storage.go` is S3/FS only. **But** the row's `Depends On (Direct)` = `-` and the flat *"No Postgres, no Redis"* omit that storage's compose block **does** declare `depends_on: redis {service_healthy}` + `postgresql {service_healthy}` at `docker-compose.yml:213-217` — the very declarations the table header (`:7`) names as its source. One row up, Sentinel's identical `depends_on: postgresql` **is** surfaced (as `Infrastructure: Postgres`). Recommend a parenthetical rather than a silent `-`.
7. **`dependency_map.md:18` (Roadrunner)** — *"orphaned, nothing calls it"*. True in effect (the caller is itself a husk), but the wiring is still present: the jobsimulation husk sets `ROADRUNNER_RPC_ADDR=http://roadrunner:10401` (`docker-compose.yml:118`) and declares `depends_on: roadrunner` (`:136`).
8. **`sentinel.md:44-45`** — *"[`manager`] appears … only as a fixture string in sentinel's own tests (`internal/authorization/casbin_test.go`, `internal/rpcsrv/rpc_test.go`)"*. A grep-driven reader will also hit `internal/authorization/manager.go`, `manager_test.go` and `internal/rpcsrv/rpc.go` — all the `Manager` **type**, not a role. The role claim itself is correct: `grep -in manager init_policy.sql` returns 0 hits (rc=1, clean).
9. **`db-backup.md` (whole file)** — none of its five claims (Go; 6-hour cadence; S3 + Azure + Hetzner; RDS all-schemas source; Docker scheduled) is checkable from any ground-truth clone in scope: there is no `db-backup` repo in `stack-demo/` and no compose entry. Reported as **unverifiable-in-scope**, not as a finding.

---

## What I checked hardest (and where the audited zeros are)

The instruction singled out three areas as freshly repaired. I re-derived all three from the SQL rather than the prose, and **all three now read correctly**:

- **`20260729133514.sql`** — read in full (64 lines). `:58` is the *"5. Drop the mirrors."* comment, `:62` is `DROP TABLE "local_jobsimulation_sessions"`; the re-point of the *referencing* `organization_assignment_sessions` link ids is `:15-23`; there is **no** back-fill — `grep -rn 'SET "score"'` over `terraform/migrations/` returns 0 hits with a clean exit (rc=1, command succeeded). `hiring.md:19-23` and `:151-159` are accurate.
- **`20260722104506.sql:79`** — is a bare `DROP TABLE "sessions"`, and `atlas.hcl:8` pins `search_path=public`, so it dropped `public.sessions`; `:2` creates `public.job_simulation_sessions`. **No `app` migration references the `jobsimulation` schema** (`grep -rn 'jobsimulation\.' terraform/migrations/` → rc=1, clean). `hiring.md:28-35` is accurate *as a statement about the migration set* — B3 is about the separate, prod-only survival claim.
- **The `public.job_simulation_sessions` write-set** — verified column-by-column against the DDL and the Ent schema: `token` is `NOT NULL` (`:13`), `UNIQUE` (`:29`) and undefaulted — the only such column; `score real NOT NULL DEFAULT 0` (`:17`) ↔ `ent/schema/job_simulation_session.go:45`; `completion_status` is a plain `varchar` with no CHECK (`:12`); `anticheat_summary` is genuinely absent from the table and was a column of the dropped mirror (`20250416091037.sql:5`); `assignValues` casts unconditionally (`ent/jobsimulationsession.go:181-186`); the gqlgen marshals are bare passthroughs (`graph.go:129546-129554` and the proto twin `:129392-129400`) against 5-member SDL enums (`jobsimulations.graphqls:14`, `:128`). Every one of these anchors is exact. `sessionCols()` at `persona_write.go:152` carries `token`; `localSessionCols()` no longer exists.

Other exhaustive checks that came back clean (audited zeros): every `docker-compose.yml` line anchor in all four docs that cite one (`:45/:70-80/:83/:99/:118/:136/:144/:160/:210/:255/:256/:258/:265/:281/:324/:337-341`) — all exact; `AUTHORIZATION_ADDRESS` appears in **exactly three** blocks; sentinel's Casbin model really is 6 request / 6 policy / 3 grouping / 6 matchers (`internal/authorization/casbin.go:14-44`) and `locals.tf:4-5` really is 256/128; `init_policy.sql:63-66` and the `content_creator` block `:88-118` are exact; messenger reads no `REDIS_WORKER_INDEX` (rc=1, clean) and no authorization client (rc=1, clean); sentinel imports no Clerk/authn (rc=1, clean); `app/go.mod:31` is `clerk-sdk-go/v2 v2.7.0`; Clerk emits exactly **12** webhook event types (`internal/clerk/events/events.go:121-190`); all four `@clerk/nextjs` pins are `^6.39.2` and the expo pair is `~2.6.18` vs `~2.19.36`; `gen.py:484-492` registers exactly nine args with no `--template` and `parse_argument` (`:18-28`) swallows the leftovers; `20260722081626_jobsim_data_model.sql` creates exactly **23** tables; the taxonomy/ai/colony/proto import matrix (`dependency_map.md:46-50`) matches all four `go.mod`s exactly; and **cms.md's contested "3 → 1, not 2 → 1"** is correct — `git ls-tree 915da06^ schemas/` shows three files and `915da06` `D`s both `cms.graphqls` and `jobsimulation.graphqls` in one commit, even though the commit *title* says "2→1".
