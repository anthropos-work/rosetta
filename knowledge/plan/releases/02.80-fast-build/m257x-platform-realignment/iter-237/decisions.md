# iter-237 — decisions

## `D-M257x-237-1` — the denominator is EVERY repo, and the first reading is recorded rather than discarded

The first census graded corpus-named variables against `platform` + `app` + `sentinel` and reported **221
distinct / 448 sites** orphaned. That number was **denominator, not evidence**: its top entries
(`STACK_PUBLIC_HOST`, `MOCK_CLERK`, `FAKE_FAPI_ROSTER`, `DEMO_*`, `FENCE_KIND`) are **`rosetta-extensions`**
variables, absent from the platform *by design* — the corpus documents them because the tooling is part of
the system it documents.

It also grepped **markdown**, which made every known-dead `*_RPC_ADDR` read as "still read by `app`" —
because `app/CLAUDE.md` names them.

Widened to seven repos and restricted to source extensions, orphans fell to **28 distinct / 37 sites**.

Recorded because this is the **fourth consecutive iter** whose first reading was a denominator error
(`D-M257x-235-1` one Makefile, `D-M257x-236-1` one prefix convention, iter-234's three literal-class
corrections, this). The generalisation is in the iter's Lessons and is worth more than the number.

## `D-M257x-237-2` — the 28 orphans are CLASSIFIED, and none is repaired

Inspected, they are: archived-doc references (`intelligence.md`'s `DB_CONNECTION_BACKEND` /
`DB_CONNECTION_SKILLER`, `messenger.md`'s `SKILLPATH_RPC_ADDR`, `README.md`'s `MESSENGER_RPC_ADDR`), psql
client variables that belong to `psql` and not to the platform (`PGHOST`, `PGDATABASE`), staging-script
locals (`SMOKE_EMAIL`, `SMOKE_PASSWORD`), third-party names (`ERR_PNPM_ABORTED_…`,
`UND_ERR_CONNECT_TIMEOUT`, `AZURE_OPENAI_DEPLOYMENT`, `OPENAI_ORG_ID`), and two **regex truncations of
prose** (`BACKEND_USERS_`, `JOBSIMULATION_` — the corpus writes `BACKEND_USERS_`/`JOBSIMULATION_RPC_ADDR`
across a line break).

**None is presented as a currently-required platform variable**, so `P-237-5` is refuted and no repair
follows. Reporting the class and declining to act on it is the same discipline as `D-M257x-236-1`.

## `D-M257x-237-3` — the repair PROMOTES `DIRECTUS_TOKEN` rather than only deleting the wrong name

Deleting `DIRECTUS_PUBLIC_BASE_ADDR` alone would have left the list with **no Directus entry at all**,
which is worse: the content surface has exactly one secret a reader must supply, `.env_example:92` ships it
blank, and an empty value is the *stack boots, catalog empty* failure the demo family has hit repeatedly.

So the row is replaced, not removed, and the removed name keeps a retraction banner stating **why** it is
not settable (compose hardcodes it at `docker-compose.yml:53`) — per `§8`'s rule that a retraction which
reaches the prose but not the reason has not landed.

## `D-M257x-237-4` — every verdict read from `git show HEAD:`, never the working tree

`.env_example`, `docker-compose.yml` and all seven read-sets come from `git` at each clone's HEAD. A local
`.env` — which by policy is never committed and always present — cannot influence any verdict.
