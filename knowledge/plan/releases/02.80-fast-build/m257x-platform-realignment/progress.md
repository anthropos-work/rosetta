---
milestone: M257x
---

# M257x — progress

## Running ledger

_(iter closeouts append here, newest last)_

- iter-01 (tok/bootstrap): all 5 open questions answered vs platform origin HEAD; authored the absent `corpus/ops/platform-alignment.md` and executed its procedure; found the class's root cause (**pinning disables drift detection** — 11/11 clones `behind: null` while the log says "provably fresh") and its local mechanism (**`migrate-demo.sh`'s hand-maintained 4-tuple** creates the legacy schemas itself, bypassing `repos.yml`); refuted 5 inherited/audited claims by measurement, one of which inverted a planned guard — see iter-01/progress.md
- iter-02 (tik): **the hand-maintained tuple is gone** — both migrate scripts now DERIVE the migration set from `repos.yml`'s machine-readable fields (origin HEAD says `app:public` alone; the tuple was wrong on **3 of 4** entries), the M810 silent-skip time bomb is disarmed, `skillpath` removed, and the non-derivable residual is declared debt behind a no-growth fence (14 tests, 4 mutations RED-proven). Re-survey **refuted TOK-01's next-tik direction** — the rext pin was already clean after the machine move (`D-M257x-4`); the real blocker was **no container runtime on this box** (Docker installed mid-iter by the user). Found a v2.1 test that **pinned the drift as a contract and still passed by reading its own refutation** (`D-M257x-6`) — see iter-02/progress.md

- iter-03 (tik): **the gate's instrument has a clone set on this host for the first time** — a pinned `stack-demo/rosetta-extensions` consumption clone (`fast-build-m257x-iter-02`) whose **pin guard MATCHED live**, plus all 10 `repos.yml` repos via `make init` in ~4.5 min (~590 MB). Phase 0d pre-flight green on every precondition (Docker 29.6.2 arm64, SSH, atlas, go, `GH_PAT`). **iter-02's derivation validated LIVE against the real cloned `repos.yml`** — `app:public` alone, exactly as its fixtures predicted, where the old tuple would have migrated four repos including one absent from the clone set. Bring-up not attempted (budget); clauses unchanged 0/5 — see iter-03/progress.md

- iter-04 (tik): **the first bring-up in this milestone to COMPLETE** — 18m 13s cold, 15 containers, autoverify 3-FAILED-but-UP. Two stacked defects made a clean host unbuildable and were fixed: `apply-authn.sh` cloned **private** colony from an anonymous HTTPS URL (exit 128 on any box without an ambient credential helper — every other rext acquisition already used SSH), and `up-injected.sh` invoked it with `>/dev/null 2>&1`, the **same masking class M217 fixed on the very next call in the same function**. 9 new tests, 4 mutations RED-proven. **The founding hypothesis then fired live: 7 seeders / 3 `jobsimulation.*` relations / 42P01** — gate clause 4 is now a measurement, and it fired *because* iter-02's derivation is correct. Side: two offline guards over these very files were RED-since-iter-02 and silently-skipping; `HOST-M257x-toolchain` partly closed (pytest + shellcheck) — see iter-04/progress.md

## Routes carried forward

| item | why | target |
|---|---|---|
| `DECIDE-M257-jobsim-schema-ownership` | The exit blocker that created this milestone: platform says cms/jobsimulation/roadrunner own no local schema; rext writes ~15 `jobsimulation.*` tables. **Inherited from M257 iter-03 — this milestone owns it now.** | iter-01+ |
| `FIX-M257-feedback-score-approximation` | Benign between a mirror and its source; **not** benign between two tables claiming to be the same row. | M257x |
| `DOC-M257-studio-in-app` | Corpus says studio-room is CMS-only in 5 places; nothing records `app` embeds it. | M257x |
| `FIX-M257-stacksnap-directus-sequences`, `FIX-M257-directus-coldstart-order` | Carried from M257 iter-02, both platform-shape-dependent. | M257x |
| ~~`HOST-M257x-stack-demo`~~ **DONE** | Bootstrap **completed after iter-03 closed**: `673 s` cold, lockfile written, 1.4 GB, "true peer of stack-dev". Do **not** re-run it to "finish trailing phases". Residual for iter-04 is only: provision secrets values-blind, then attempt the first `demo-up`. | — |
| `CHECK-M257x-demopatch-pristine` **NARROWED (iter-04)** | **23 of 23** manifests logged `⚠ pristine-ing skipped/failed` (not 5 — that was read off a truncated `tail`). Probably **benign**: on a minutes-old clone there is nothing for R1's pre-emptive revert to undo. The real finding is that the message **collapses `skipped` and `failed`** into one string at one severity — opposite conditions, one healthy and one the class that shipped a 76 s members grid for four releases (`next-web-members-pagination.yaml` is in the list). Re-scoped: split the two states in the log, then re-read. | iter-04 |
| ~~`CHECK-M257x-pin-state-on-fresh-clone`~~ **EXPLAINED (iter-04)** | `clones.lock.json` marks `platform` + `graphql-wundergraph` `pin_state: pin-drift` on **minutes-old** clones (all 11 are `ref: main`, `behind: 0`). `pin-drift` is escalated by `DEMO_FRESHNESS_STRICT=1`, so this could refuse a legitimately fresh stack. Observation, not a defect claim — state semantics unread. | iter-04 |
| `FIX-M257x-vmram-gib-unit` | `up-injected.sh:258-262` floors bytes to integer GiB, so a VM set to the documented "12 GB" (decimal) = 11.67 GiB → floors to 11 → trips the non-fatal `< 12 GiB` warning. A doc/code **unit mismatch**, never re-measured — this milestone's own subject matter. Non-fatal. | iter-04 |
| `HOST-M257x-toolchain` **PARTLY CLOSED (iter-04)** | No `pytest`, `gh`, `psql` or `tailscale` on this box. Two mutation batteries (`m220`, `m255`) cannot run at all — they shell out to `python3 -m pytest` and fail with zero named tests. | iter-04 |
| `REPOINT-M257x-jobsim-writes` **CONFIRMED FIRING (iter-04)** | ~12 `jobsimulation.*` tables (9 written) live in `stack-seeding/cmd/stackseed/main.go:45-105`. Until re-pointed, the transitional schema debt cannot shrink and **gate clause 4 cannot be met**. | later tik |
| `FIX-M257x-migrate-dev-swallows-atlas` | `migrate-dev.sh`'s atlas loop still `>/dev/null 2>&1`s every failure into "non-fatal migration warnings" — the M215-F8 masking class its demo twin already fixed. | later tik |

### New routes opened by iter-04's first completed bring-up

All three are `autoverify` ✗ checks on a stack that is UP — i.e. they are exactly what stands between here
and **gate clause 1** (`green:true / 0 warnings` × 3 consecutive cold cycles).

| item | why | target |
|---|---|---|
| `FIX-M257x-autoverify-skillpath-schema` | autoverify's `postgres-schemas` probe fails with `missing schemas: skillpath` — it demands the schema **iter-02 correctly removed** (absent from origin `repos.yml`). A stale check pinning the drift as a contract (`platform-alignment.md` §8 rule 3), and one of the 3 checks blocking clause 1. | next tik |
| `FIX-M257x-directus-container-exit1` | `demo-1-directus-1` **exited(1)**; `directus` probe returns HTTP 000000. The per-stack Directus is `--local-content`'s whole point, and this blocks clause 1. | next tik |
| `FIX-M257x-academy-not-serving` | ant-academy never answers on `:13077` within 120 s while its own log says `✓ Ready in 193ms` — a readiness probe disagreeing with the process. Not on clause 1's critical path but it is a ✗. | later tik |

**And the headline route is no longer a hypothesis:** `REPOINT-M257x-jobsim-writes` reproduces on demand as
**7 failing seeder surfaces** (`content-stories`, `hiring-funnel`, `jobsim-sessions`, `personas`, `activity`,
`succession`, `ai-readiness-funnel`) across **3 relations** (`jobsimulation.sessions`,
`.activity_events`, `.interview_extraction_results`), all `42P01`. That is gate clause 4's fence going RED,
which is the precondition for paying the debt down deliberately rather than by inspection.

- iter-05 (tik): **autoverify FAILED 3 → 2, containers 15 → 16/16, `verify live: all … probes passed`.** Two clause-1 blockers cleared: the `postgres-schemas` probe carried a **hand-written** expected list still demanding `skillpath` (the schema iter-02 correctly stopped creating) — now DERIVED from `repos.yml` via the same helper the migrator uses, fail-loud if the source is absent; and **`FIX-M257-directus-coldstart-order`, carried since M257 iter-02, is closed** — directus was the only service with neither a restart policy nor a readiness dependency, raced Postgres, and stayed `Exited(1)`; proven a pure ordering race (it serves 200 once started by hand) and fixed with the platform's own `restart: on-failure` + `depends_on: service_healthy`. 7 tests, both mutation-verified RED. Side: **iter-04's own edit shifted 7 `file:line` citations** in `demo-up-defaults.md` and the corpus fence caught it — repaired with the guard's `--fix` — see iter-05/progress.md

### Pre-computed input for iter-06 (`REPOINT-M257x-jobsim-writes`) — measured, do NOT re-derive

Measured against the live demo-1 database at the end of the iter-05 session, so the next iter starts from
evidence rather than from a fresh survey. **The re-point mapping is very nearly 1:1.**

    jobsimulation.<t>  ->  public.<t>     for 10 of 11:
      actors · interactions · activity_events · interview_extraction_results ·
      interview_aggregated_reports · validation_criterion_results ·
      validation_attempt_skill_results · validation_attempt_results ·
      code_submissions · collaborative_assets
                                          (each already EXISTS in `public`, same name)

    jobsimulation.sessions  ->  public.job_simulation_sessions      <- the ONE exception

The exception is already explained by `D-M257x-1`: app created `sessions` and renamed it to
`job_simulation_sessions` in the very next migration. `public.sessions` does not exist, which is why a naive
re-point fails LOUD rather than silently reading blank.

**Two things iter-06 must decide rather than assume** (neither is settled by the name mapping):

1. **The session PAIR.** `stackseed/main.go:524` shows the hiring funnel already writes **both**
   `jobsimulation.sessions` **and** `public.job_simulation_sessions` (the latter is where M257 re-pointed
   the dropped `local_*` mirror). So for `sessions` the fix may be *removing* the legacy write rather than
   re-pointing it — but `platform-alignment.md` §7 rule 2 forbids re-pointing to nothing without asserting
   the replacement. Check each pair before deleting either half.
2. **Column drift.** Same table NAME in `public` does not entail the same COLUMNS. The old
   `jobsimulation.*` shapes are three-plus releases old. Verify column-by-column before trusting the 1:1
   mapping — this is exactly the "a fidelity check against the wrong reference passes" trap (Trap A).

Once the writes land, `stack-core/lib/repos_yml.sh`'s `REXT_TRANSITIONAL_SCHEMAS` ("cms jobsimulation") can
shrink — and its no-growth fence is designed to fail on a SHRINK too, with "this failure is good news",
which is the deliberate act that closes **gate clause 4**.
