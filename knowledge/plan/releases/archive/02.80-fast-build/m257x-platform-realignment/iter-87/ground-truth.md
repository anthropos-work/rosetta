# iter-87 — ground truth, re-derived at open

**Every number below was measured in this iter.** Nothing is carried from the hand-off; where the
hand-off's figure and the measurement agree that is stated, and where they disagree the measurement wins
(§5 rule 32 — *re-derive the hand-off's numbers, including the orchestrator's*).

## Refs at open (§5 rule 26 P1 — state the refs in the artifact)

| ref | value |
|---|---|
| rosetta HEAD | `ae5c1db` (== `origin/m257x/platform-realignment`) |
| rext authoring copy | `ac30b9b` on `main` (== `origin/main`) |
| platform clone, **before** | `0dab54d` — **2 behind** `origin/main` |
| platform clone, **after** | `0c91421` — **0 behind**, working tree clean, nothing committed into it |
| instrument | `guard_family.py` @ rext `ac30b9b` |

## The move

```
0c91421 2026-08-05T16:21:19+02:00  Merge pull request #26 from anthropos-work/chore/drop-support-service-containers
838d907 2026-08-05T16:14:25+02:00  chore(compose): drop the storage, messenger and customerio-sync containers
```

Diff: `CLAUDE.md`, `README.md`, `docker-compose.yml` (−107), `repos.yml` (−21). No code, no migrations.

## §4's six signals, at `0c91421`

**Signal 1 — `repos.yml` repo set: 4** (was 6). Derived, not grepped:

```
app · sentinel · next-web-app · studio-desk
```

`storage` and `messenger` removed. The header comment moved with them and now reads: *"roadrunner,
jobsimulation, cms, messenger, storage and customerio-sync are all served in-process … `make init`
therefore does not clone them."*

**Signal 2 — migrating repos: `app` alone.** Unchanged.

**Signal 3 — declared schemas: `app -> public` alone.** Unchanged.

**Signal 4 — subgraphs: one** (`backend.graphqls`). Unchanged.

**Signal 5 — compose service set**, parsed as YAML (never grepped):

| | `0dab54d` | `0c91421` |
|---|---|---|
| declared in `docker-compose.yml` | 8 | **5** — backend, gotenberg, next-web-app, sentinel, studio-desk |
| via `include: [common.yml]` | 2 | **2** — postgresql, redis |
| **effective topology** | 10 | **7** |
| profiles present | 8 | **5** — `all`, `backend`, `core`, `frontend`, `studio-desk` |
| profiles removed | — | **`storage-legacy`, `messenger`, `customerio-sync`** |

`docker compose --profile core config --services` → **5**: backend, gotenberg, postgresql, redis,
sentinel. `PROFILE ?= core` unchanged at `Makefile:10`.

**`STORAGE_S3_BUCKET=production-storage20240826131618541000000005` is still on `backend` at
`docker-compose.yml:82`.** The `storage` *container* is gone, so `backend` is now the sole writer to that
production bucket — the standing escalation `DEF-M257x-iter80-storage-prod-bucket` is unchanged in
substance and cleaner in shape. **Not touched by this iter** (carve-out).

**Signal 6 — org repo census.** Not re-run; no repo entered or left the org in a 2-commit compose change,
and the signal costs an API sweep. Declared skipped rather than silently omitted (§5 rule 8).

## Clone set — 13 clones, all fetched

| clone | HEAD | behind `origin/main` |
|---|---|---|
| app | `b948604ff` | **93** |
| next-web-app | `bb3313bc0` | 41 |
| rosetta-extensions (consumption) | `ab81527` | 34 |
| storage | `4ce8ece` | 20 |
| messenger | `fa47850` | 7 |
| ant-academy | `9c3843cd` | 4 |
| jobsimulation | `462343b0` | 4 |
| cms | `ca50c81` | 2 |
| sentinel | `88bc559` | 2 |
| studio-desk | `14a5442` | 2 |
| graphql-wundergraph | `60c229f` | 0 |
| roadrunner | `87d8d44` | 0 |
| platform | `0c91421` | **0** (advanced this iter) |

`app` at **93** confirms the hand-off's figure and **refutes iter-86's 60**, which was measured 1 day and
33 commits ago and was carried forward unchecked.

## Citation exposure, by repo (the input to `D-M257x-87-1`)

Repo-prefixed code citations across `corpus/**` + `CLAUDE.md` + `.claude/**` (111 files searched;
positive control: `corpus/ops` present):

```
app 65 · jobsimulation 8 · graphql-wundergraph 7 · roadrunner 6 · cms 6 · messenger 6
storage 4 · next-web-app 2 · sentinel 1 · rosetta-extensions 1 · ant-academy 1
(unprefixed, resolved via the doc's own service map) 234
```

## The control — which assertions the checkout moved, and which the FETCH moved

A detached worktree at `0dab54d` was created in the same tree, the family run against it, then removed.

| | worktree @ `0dab54d` | clone @ `0c91421` |
|---|---|---|
| guard family | **10 GREEN · 3 RED · 3 not-run** | **9 GREEN · 4 RED · 3 not-run** |
| `platform_predicate_guard` | **GREEN** | **RED** — 17 findings |
| `platform_alignment_guard` assertion B | **0 findings** | **2 findings** |
| `anchor_construct_guard` | RED (15) | RED (15) — *identical* |

Two conclusions, and the second is the new one:

1. **The membership fence caught the departure unaided**, 0 → 2 across the checkout advance, with the
   control holding everything else fixed. This is the property the hand-off called decisive.
2. **The citation guards were armed by the `git fetch`, not by the checkout.** They read `origin/main`
   by the iter-68 `CITE_REF=auto` ladder, so their verdict is identical at both checkouts and changed
   only when the remote-tracking ref moved. The hand-off's *"13 GREEN · 0 RED"* was a reading taken
   against an **unfetched** clone; the same checkout, fetched, reads 10 GREEN · 3 RED.

## Findings enumerated at open — 38, all consequences of the move

| guard | n | predicate |
|---|---|---|
| `platform_alignment_guard` [B departure] | 2 | *service X is in `repos.yml`* |
| `platform_predicate_guard` [G1 dead-token] | 3 tokens / **28 sites** | *profile token X selects something* |
| `platform_predicate_guard` [G8 no-such-service] | 3 | *service X declares `profiles:`* |
| `platform_predicate_guard` [G10 service-count] | 1 | *compose declares N services* |
| `platform_predicate_guard` [G2 repo-count] | 1 | *`repos.yml` lists N repos* |
| `platform_predicate_guard` [G4 unset-address] | 9 | *the platform sets `<VAR>` locally* |
| `anchor_construct_guard` | 15 | *this line number names this construct* |
| `platform_alignment_guard` [F] | 19 | (same predicate, map-scoped) |

The **G4** class deserves its own line: the `messenger` compose service was the *only* thing that set
`BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` and `SKILLER_RPC_ADDR`. Deleting the
service deleted the four assignments, so **the last cross-process RPC edge besides `backend → sentinel`
is gone** — messenger is no longer a process that could hold one. `CLAUDE.md`'s communication-patterns
section asserts that edge in the present tense and cites `docker-compose.yml:174`/`:176`, two lines that
no longer exist.

Raw transcripts: `anchor-open.txt`, `palign-open.txt`, `ppred-open.txt` (scratch, reproduced in the close).
