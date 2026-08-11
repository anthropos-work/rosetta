# iter-64 — decisions

## `D-M257x-64-1` — a vocabulary gap is a claim the map cannot make

`DOC-M257x-iter59-storage-mid-fold` had been routed forward five times, and the reason it kept
slipping is worth naming: **the measurement had nowhere legal to go.** iter-59 measured the storage
split and wrote it into `storage.md` as a two-sided record. The *map* — the fenced artifact, the one
`platform_alignment_guard.py` assertion C polices — had a seven-token vocabulary with no token for it,
so `storage` kept reading `live-standalone` on both sides. Not because anyone believed it, but because
inventing a token in a fenced field turns the guard RED.

**The fence had eight things to say and seven words.**

The rule this yields: **widen the vocabulary in the same iter that takes the measurement.** Otherwise
the measurement lands in a service doc — unfenced, and therefore free to rot — while the fenced
artifact keeps the old answer and keeps passing. A guard that is GREEN because its vocabulary cannot
express the truth is the same failure shape as a fence whose reach is narrower than its class
(`D-M257x-63-2`), one layer up: **in both cases the instrument reports agreement it never tested.**

## `D-M257x-64-2` — markdown emphasis is presentation, not part of the token

The first row to *bold* its state produced `[C vocabulary] storage: fresh local stack state
'**mid-fold**' is not one of [...]` — a finding that reads like an invented vocabulary word and is
actually a formatting artifact. `_state_head` already tolerated a parenthetical qualifier
(`external (Vercel)`); it now strips `*` and backticks too. Derived from the format, not an exception
for one row, and tested both ways: emphasis is stripped, **and** stripping it does not launder an
illegal token (`**totally-made-up**` is still refused).

Worth stating because the alternative was tempting and wrong: un-bolding the cell would have made the
guard green while leaving the next author the same trap.

## `D-M257x-64-3` — the split is re-derived, and iter-59's consumer count was short by one

Adjudicated against artifacts at platform `0dab54d` / `app` `b948604` v1.366.0, not against
`storage.md` (§5 — never against another document):

| side | measurement |
|---|---|
| config | `STORAGE_RPC_ADDR`: **0** occurrences across `docker-compose.yml`, `common.yml`, `.env_example` |
| compose | `storage` moved to `profiles: [storage-legacy]` (`docker-compose.yml:135`), rationale in-comment at `:131-134` — *two writers on one bucket* |
| `repos.yml` | `storage` **still an entry** (`:18-20`, `migrations: false`) — still cloned |
| consumer | `app` reads it at `main.go:446`, `:524`, `:992` and in **three** `cmd/` tools |

iter-59 recorded **two** `cmd/` readers. There are three: `cmd/academyImport/main.go:231` and
`cmd/academy-asset-upload/main.go:129` hard-require it (`:235` / `:133`), and **`cmd/import/main.go:50`
builds a storage client against the empty string without complaint** — the same silent-deferral shape
as `main.go`, in a tool nobody had counted. The map now names all three.

## Routed forward

- **`CHECK-M257x-iter64-pms-87-subject`** — `service_taxonomy.md`'s Directus-retraction passage cites
  `platform-migration-status.md:87` as *"the corpus's own fenced source of truth"*, and that row is
  `anthropos-studio-room`, which is not about Directus. The line number was re-pointed mechanically
  (the row it named moved +1 with the new vocabulary row) and is therefore **as correct as it was**;
  the *subject* looks wrong and predates this iter. Sibling of
  `CHECK-M257x-iter60-g6-citation-subject` — the same class: an anchor that resolves and still does
  not name the claim.
