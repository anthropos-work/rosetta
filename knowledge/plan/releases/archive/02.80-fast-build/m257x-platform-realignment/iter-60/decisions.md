# iter-60 — decisions

## `D-M257x-60-1` — the sibling guard is `platform_predicate_guard.py`, and its subject is PREDICATES

Built as a **new** `stack-core/*_guard.py` per `D-M257x-59-2`, not as a widening of
`platform_alignment_guard.py`. That guard's `FENCE_KIND = "standalone"` scoping and its
`repos_yml_path`-derived clone roots are what make assertion F trustworthy; re-targeting it at the
whole corpus would have broken the *same-reference* property. `compose_blocks()`'s parsing shape is
reused as a primitive. §8 rule 1's derived registry means the new guard self-registers — verified by a
test asserting it declares a legal `FENCE_KIND`.

**Six assertions, both directions**, on the `demo_knob_guard` contract. All six denominators
re-derived at platform `0dab54d` and **cross-checked against `docker compose --profile X config
--services`**, not just against a parser I wrote:

| # | derived | value |
|---|---|---|
| G1 | compose services (`docker-compose.yml` 8 + `common.yml` 2, `include:` resolved) | **10** |
| G1 | the **floor** — services declaring no `profiles:` key, so in *every* selection | **3** (`postgresql`, `redis`, `sentinel`) |
| G1 | legal profile tokens | **8** (`all backend core customerio-sync frontend messenger storage-legacy studio-desk`) |
| G2 | `repos.yml` entries | **6** |
| G3 | `Makefile` default (`PROFILE ?=`) and its selection | **`core`** → **5** |
| G4 | `*_RPC_ADDR` compose sets | **4**, all `http://backend:8083` |
| G5 | repos with `migrations: true` | **1** (`app`) |
| G6 | the mid-fold split, per variable | `STORAGE_RPC_ADDR` = unset config / 6 app reads |

## `D-M257x-60-2` — the failure taxonomy is THREE classes, not two, and one of them was unrecorded

TOK-05 briefed two: *works* and *silently no-ops*. Measured at `0dab54d`, `docker compose --profile
X config --services` splits eight ways into **three**:

| class | tokens | behaviour |
|---|---|---|
| **works** | `core` · `backend` · `all` · `storage-legacy` · `customerio-sync` | rc=0, selects beyond the floor |
| **silent no-op** | `graphql` · `cms` · `jobsimulation` · `roadrunner` · `storage` | **rc=0**, starts the floor and nothing else |
| **hard-fail** | `frontend` · `studio-desk` · `messenger` | **rc=1** — `service "X" depends on undefined service "backend": invalid compose project` |

The third class was in **no** corpus document. `make up PROFILE=frontend` and `make up
PROFILE=studio-desk` are documented commands in `setup_guide.md`'s command table and **both exit 1**:
each service declares `depends_on: backend`, which its own profile does not select, so compose rejects
the whole project. Of the six profile rows `CLAUDE.md` carried, **one** (`all`) was accurate.

## `D-M257x-60-3` — four RULES replaced after the first draft's own false positives (§4 Trap A)

The first draft produced **37 findings, of which 16 were its own**. Every one was fixed by replacing
the rule with one derived from the artifact's structure — never by excepting the offending names:

| its own finding | the rule that produced it | the rule that replaced it |
|---|---|---|
| `--profile billion`, `laptop`, `odysseus`, `names` | `--profile` read as a compose flag by spelling | a bare `--profile` counts only where the window names a **compose driver** — `buildbench --profile <host>` selects a checked-in host profile, a different tool entirely |
| `--profile P` | any token accepted | the token must match compose's **own token shape** (`^[a-z][a-z0-9-]*$`); `P` in an `argument-hint` is a metavariable |
| `§ 2 Clone repos` read as "2 repos" | `\d+ … repos` case-insensitive | the modifier slot is **case-sensitive** — a capitalised word after a numeral is a section number, and the document itself says so with the capital |
| `jobsimulation.sessions` read as a repo | word-boundary name match | a name adjacent to `.` or `/` is part of a longer identifier — a schema-qualified table or a path, both ubiquitous here |

And one rule was **dropped rather than fixed**: per-name attribution of `migrations: true|false` in
free prose. English puts the name list **before** *"are `migrations: false`"* and **after** *"have
`migrations: true`"*, so a nearest-marker rule is wrong roughly half the time — and a fence that is
right by accident teaches nothing. G5 now checks the one construct it *can* attribute (an explicit
`(currently: …)` enumeration) plus the direction-free count invariant, and **names the other 22 of 26
migration lines as UNREACHED** rather than passing them.

## `D-M257x-60-4` — G6 grades the CORPUS, not the platform

The first G6 fired whenever a variable was unset-but-read. That would have been **permanently RED**,
because a half-landed fold is a legitimate platform state and the developer is mid-program — and a
permanently-RED fence is one that gets disabled. `D-M257x-59-4`'s actual requirement is a **two-sided
record**: the guard *derives* the config side, and the corpus must *cite* the consumer side. G6 now
fires only when **no scanned document cites any of the variable's real read sites**. Closed this iter
by the two-sided block in `corpus/services/storage.md`.

Known limitation, stated rather than hidden: the citation test matches the read site's `path:line`, not
its subject, so a citation that lands on the right line for another reason would satisfy it. The one
that closed it (`app/main.go:992` `storage.NewClient(…)`) was checked by hand and is genuine.

## `D-M257x-60-5` — do not spell a dead command in runnable form, even to warn about it

The fence went RED **on this iteration's own repair text**: warnings written as
`` `make up PROFILE=cms` does NOT fail `` re-introduced the very construct G1 counts. That is the fence
being right, not over-eager — a copy-pasteable invocation for a silent no-op is indistinguishable from
an instruction to a reader in a hurry. All warnings were rewritten to **name the token without writing
the invocation**, and retired tokens were moved out of profile-table first cells into prose beneath.
Promoted to §5 rule 30's second corollary.

## `D-M257x-60-6` — a refutation expires exactly like the claim it refuted

iter-22 correctly refuted a proposed `*_RPC_ADDR` correction *at `2adcf71`*, and the write-up became
standing guidance: *"That address is **CURRENT, not stale text**"* (`jobsimulation.md:95`), with the
protocol doc adding *"Applying the correction would have replaced two true statements with false
ones."* At `0dab54d` **M809 has landed** — all four compose `*_RPC_ADDR` values read
`http://backend:8083` and there are no husk containers — so a passage written to forbid a repair now
forbids the repair that is required.

This is worse than ordinary staleness: **emphatic anti-repair language reads as already-adjudicated and
survives readings that would have caught plain prose.** Both passages are repaired and every
ref-dependent sentence now names its ref. Promoted to §5 rule 31.

Note the interaction with the fence, recorded because it is a hole and not a feature: a ref-pin
**exempts** a claim from G4. `jobsimulation.md:95` was exempted correctly by rule (it *is* a pinned
measurement) while still reading as current guidance. The guard now **prints the refs that bought each
exemption** so the hole is visible rather than silent; converting "pinned to a superseded ref" into a
finding is routed forward.

## `D-M257x-60-7` — three inherited numbers re-derived; two corrected

Per §5 rule 1 and the standing instruction to treat orchestrator facts as any other hand-off:

| inherited | measured at `0dab54d` | verdict |
|---|---|---|
| "17 files / 30 occurrences" assert a `graphql` profile | **26 live docs / 56 lines** by whole-tree grep; the fence's parsed-construct reach is 53 sites / 11 tokens | **corrected** (the briefed figure was an undercount) |
| `cmd/academyImport/main.go:235` hard-requires `STORAGE_RPC_ADDR` | the **`Getenv` is `:231`**; `:235` is the `is required` return | **both true, different lines** — the read site is what G6 needs |
| `cmd/academy-asset-upload/main.go:133` | `Getenv` at `:129`, return at `:133` | same shape, confirmed |
| six G1–G6 denominators | all six reproduced, and cross-checked against `docker compose config` | **held** |

Promoted to §5 rule 32.
