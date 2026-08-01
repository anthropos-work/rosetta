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
| `HOST-M257x-toolchain` **PARTLY CLOSED (iter-04); MOVED AGAIN (iter-06); RE-MEASURED (iter-07)** | No `pytest` on the system `python3` (only the ephemeral `/tmp/rextvenv`), no `gh`/`psql`/`tailscale`/`timeout`. **Closed at iter-06:** Playwright/npm e2e (`stack-verify` 210/210). **Closed at iter-07:** `PyYAML` installed into the rext venv — it had been silently skipping **8** `stack-injection` tests (all 8 pass; the evidence hole is what mattered, §5 rule 8). **STILL OPEN, and the iter-07 green is misleading:** `stack-seeding`'s `services/ai` tests now PASS with `GOPRIVATE` **empty** and no `insteadOf` rewrite — only because `ai v1.40.1` already sits in `$GOMODCACHE` (confirmed with `GOPROXY=off`). On a cold module cache the **Trap E** failure is unchanged. Stopping at "the tests pass" would have closed this on evidence that does not support it. | later tik |
| ~~`REPOINT-M257x-jobsim-writes`~~ **DONE (iter-06)** | 12 write targets re-pointed at `public`, 3 duplicate session writes removed, proven live (7 failing seeders → 0). The debt list shrank to `cms`. | — |
| ~~`REPOINT-M257x-cms-similarity-writes`~~ **DONE (iter-07)** | Fixed by DERIVING the replay schema from the target, not by a second constant. `rc=4 skipped` → 1490 rows into `public`, proven with `cms` dropped from the stack. `REXT_TRANSITIONAL_SCHEMAS` is now **EMPTY** and the derived CREATE SCHEMA set is `extensions/sentinel/public`. | — |
| ~~`FENCE-M257x-write-fence-scans-one-section-of-nine`~~ **CLOSED (iter-08) — premise REFUTED, residual FIXED** | iter-07 opened this claiming the fence's scope limit was undocumented and that the fix was to widen the scored set. **Both refuted by measurement:** the rationale was ten lines above the constant iter-07 quoted, and widening would have scored **0** constructs (`stack-seeding` 92, every other Go section 0) while reporting GREEN. What was real and is now fixed: the exemption's rationale had gone **stale because of iter-07 itself** (it cited an `rc=4` signal iter-07 removed), and nothing mapped a section to which of §8's three layers covers it. `SECTION_COVERAGE` now declares `(layer, reason)` per Go-bearing section, `SCORED_SECTIONS` is **derived** from it, and a new section goes **RED naming itself**. Clause 4 claimable on written, machine-checked coverage. | — |
| ~~`CHECK-M257x-bringup-evidence-logs-absent`~~ → **`FIX-M257x-autoverify-evidence-log-path` (iter-09)** — **now the ONLY autoverify ✗ (iter-10 measured 2/2)**, i.e. the last thing between here and a clean `green:true / 0 warnings` | **Not missing evidence — autoverify looks in the wrong directory.** It reads `$STACK_DIR/{demopatch,buildfail}.log`; the bring-up writes them to `stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/`. Both exist. Its message then asserts a CAUSE from the absence (*"its ABSENCE means the phase never ran"*) — a conclusion drawn from looking in the wrong place, i.e. the dominant class inside the verifier that measures the gate. Second layer: at the correct path both are **0 bytes** while the patches demonstrably applied (their output went to stdout), so the fix must distinguish **absent / empty / populated** — three states the message collapses into one. | next tik |
| ~~`FIX-M257x-academy-not-serving`~~ **CLOSED (iter-10) — iter-09's MECHANISM REFUTED, fix landed + live-proven** | Not a resolution-order bug (`curl 127.0.0.1` fails identically; `--dns-result-order=ipv4first` changes nothing; `-H ::1` fails too). An **origin-STRING equality** across three next@16 files: the middleware normalizes every loopback host to `localhost`, the router keeps the raw `-H` string, and `relativize-url.js` compares with `===` — so the in-app rewrite is proxied externally to itself until http-proxy's `30_000` ms default. Fixed with **`-H localhost`**, exposure intact + re-proved three ways. `GET /` 500@30.077 s → 200@0.205 s. | — |
| `CHECK-M257x-live-clone-suites-red` **NEW (iter-06)** | 7 `demo-stack/tests` live-clone/live-container tests are RED on this host — `test_ant_academy` devorigins round-trip, `test_back_to_cockpit_m249` apply/revert, 2 × `test_demopatch` live-clone hash, 3 × `test_migrate_race_live`. **Reproduced identically on the pristine control clone**, so pre-existing, not an iter-06 regression. Two are demopatch apply/revert round-trips — the class that shipped a 76 s members grid for four releases. | later tik |
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

- iter-06 (tik): **the milestone's headline route landed and was proven live — `stackseed: 7 seeder(s) failed` → 0, autoverify FAILED 2 → 1.** All **12** `jobsimulation.*` write targets re-pointed at the `public` tables jobsim-in-app created (the pre-compute listed 11 — `validation_check_results` was missing, found because the re-survey measured instead of trusting); the **three co-written session PAIRs** collapsed to their canonical half (`D-M257x-6`: for a pair, removal IS the re-point, and §7 rule 2's assertion is the very next step in the diff); the FK ordering that was free across schemas became load-bearing within one and was fixed in both flush slices and `resetTables`. `REXT_TRANSITIONAL_SCHEMAS` shrank `cms jobsimulation` → `cms` and the no-growth fence's **shrink branch fired as designed** (*"this failure is GOOD NEWS"*), so the paydown is a claim in writing. New **derived** fence `test_write_target_schema_fence.py` names no dead schema at all — it reads the legal set from the same `repos_yml.sh` the migrator and the verifier use, so the next fold goes RED offline; 2 mutations RED-proven, its fixture mutation-verified too. Side: a `shellcheck` guard RED since **iter-02** in a suite nobody re-ran — see iter-06/progress.md

- iter-07 (tik): **the last transitional schema is paid off, and `REXT_TRANSITIONAL_SCHEMAS` is now EMPTY — gate clause 4 is claimable.** The `sim-embeddings` surface went `rc=4 skipped` → **1490 rows** into `public` on demo-1, and still 1490 after `DROP SCHEMA cms`. Fixed by **deriving** the replay target from the stack's own catalog (`replay.ResolveTargetSchema` over `pg.SchemasHoldingAllTables`) rather than declaring a second constant — the rejected two-line fix was the same hand-maintained-list defect that has been wrong three releases running. The digest probe and the replay are now **one construct** (`resolveThenProbe`), so the half-done re-point the pre-compute warned about is unwritable, not merely fenced; `replay.Run` takes the resolution as a **required positional**, forcing all 18 call sites to decide. Re-survey **answered a question the pre-compute recorded as unanswerable**: the cache key omits the schema name, so the 2026-06-29 `cms` capture hits a `public` stack byte-exact — no re-capture. An eager first cut was caught by three tests whose own comments documented the contract it broke (a cache-miss verdict must not need a live DB; it had silently turned exit 5 into exit 4). Derived CREATE SCHEMA set is now `extensions/sentinel/public` — rext creates **no schema the platform does not own**; `verify live: all probes passed` from the stack's own re-pointed clone. 5 fences mutation-verified RED — **one of which first "passed" as a compile break** with an empty failing-test list, re-run with a compiling mutant and promoted to `platform-alignment.md` §8 rule 5. Side: the wider sweep found **8 `stack-injection` tests silently skipping on missing PyYAML** — installed, all 8 green, evidence hole closed — see iter-07/progress.md

- iter-08 (tik): **the inherited finding was REFUTED by measurement, and the narrower real gap was fixed.** iter-07 claimed the write-target fence's scope limit was undocumented — it was documented in **ten lines directly above the constant iter-07 quoted** — and claimed the fix was to widen the scored set, which measurement shows would score **0** constructs (`stack-seeding` 92, every other Go section 0) and report GREEN. The milestone's own dominant defect, committed by the milestone; corrected in `iter-07/progress.md` too, and added to `platform-alignment.md` §5 as **rule 10** (*read the lines AROUND the line you are quoting* — the search succeeded and the conclusion was still false). What survived is real and was landed: the exemption's stated rationale had gone **stale because of iter-07 itself** (it cited an `rc=4` signal iter-07 removed), and nothing mapped a section to which of §8's **three layers** covers it. Now `SECTION_COVERAGE` declares a `(layer, reason)` per Go-bearing section, `SCORED_SECTIONS` is **derived** from it, and the map is checked against the repo — so **a new section goes RED naming itself** instead of sitting outside the fence silently, which is the one list that could never announce its own staleness (*a fence only asserts about what it already scans*). The subtlest new test: a section classified `static` that yields **zero** constructs is **mis-classified, not covered** — it reports GREEN and *looks* fenced. 5 fences, all mutation-verified RED, each mutant **collected before being run** (iter-07's new §8 rule 5). **Gate clause 4 is now claimable on written, machine-checked coverage** rather than a reader's guess — see iter-08/progress.md

- iter-09 (tik, closed-no-lift): **the academy's four-iter-old symptom has a proven mechanism, and it is not a probe bug — a SECURITY TIGHTENING broke it.** `ant-academy.sh:466` (M221/F-M220-5) tightened the localhost bind `0.0.0.0` → `-H 127.0.0.1`, a real de-exposure fix. But Next.js 16's dev server runs an internal proxy that dials **`http://localhost:$PORT/`** — `localhost`, not the address it was told to bind — so on a host where `localhost` resolves to `::1` first the server listens IPv4-only and the self-proxy hangs. Measured both ways, same code, only the bind changed: `-H 127.0.0.1` → **`500` in a flat 30.014 s** (reproduced, identical warm; the log names it: `Failed to proxy http://localhost:13077/ … ECONNRESET`); `-H 0.0.0.0` → **`200` in 2.378 s**. **The flat 30.0 s is the tell** — the launcher's own comment had attributed exactly this shape to Turbopack cold-compile and budgeted 120 s for it, a wait that could never succeed. **No fix landed, deliberately:** the naive repair re-opens the exposure M221 closed, and the real fix must be re-proved on both host families (they resolve `localhost` differently — that IS the mechanism). Pre-computed for iter-10. Side: `CHECK-M257x-bringup-evidence-logs-absent` **resolved — an autoverify WRONG-PATH bug**; the logs exist under `.../demo-stack/stacks/demo-1/`, and autoverify concludes *"the phase never ran"* from looking in the wrong place — the dominant class, in the verifier that measures the gate — see iter-09/progress.md

- iter-10 (tik): **the academy's landing page is fixed at its REAL mechanism, and the two checks that should have caught it were both measuring something other than what they reported.** iter-09's inherited mechanism was **refuted by measurement** before a line was written — `curl 127.0.0.1:PORT/` (which never utters `localhost`) fails identically, `--dns-result-order=ipv4first` changes nothing, and `-H ::1` fails too, so it is not a resolution-order bug. It is an **origin-STRING equality**: `next@16` normalizes every loopback hostname to the literal `localhost` when its middleware builds a rewrite URL (`web/next-url.js:15-20`), builds the router's base URL from the **raw** `-H` string (`resolve-routes.js:117`), and compares the two origins with `===` (`relativize-url.js`) — so a `127.0.0.1` bind makes the app's own in-place rewrite look **external**, the dev server proxies to itself, and http-proxy's **30_000 ms default** fires. The flat 30.0 s was a constant in next's source, not a cold compile. Fix: **`-H localhost`**, the one loopback literal that is its own normalized form — **resolver-independent by construction**, because the comparison is on the string we supply. **The de-exposure is intact and re-proved three ways**, not argued: the exposure guard reads `academy:localhost` and passes, the routable address answers `000`, and the process listens on `[::1]` alone. `GET /` **500 in 30.077 s → 200 in 0.205 s**; the launcher went `✗ NEVER ANSWERED` → `started + SERVING`. **The bigger finding: autoverify never detected any of this** — it probed only `/library/` (200 in 9 ms, spared because it short-circuits in Clerk's middleware before the loop) and printed *"✓ AI Academy renders its catalog"* over a demo whose Academy link was a 500. iter-09's "the academy is the only ✗ blocking clause 1" was wrong in both directions: the baseline 2 ✗s were both the evidence-log-path bug. Both checks now measure what they claim, and **the iter's own first cut of the new check repeated the defect one level down** (it kept the LAST attempt, so a 500-then-timeout sequence read as "does not answer at all") — invisible to three unit tests, obvious on the first **live** negative control. 7 fences, 5 mutations RED-proven — see iter-10/progress.md

### Pre-computed input for iter-10 (`FIX-M257x-academy-not-serving`) — measured 2026-07-31, do NOT re-derive

**The mechanism is settled; what is open is the FIX, and it is a design.**

    ant-academy.sh:466
      bind_args=(-H 127.0.0.1); [ -n "${STACK_PUBLIC_HOST:-}" ] && bind_args=(-H 0.0.0.0)

Next.js 16's dev server proxies to **`http://localhost:$PORT/`** — `localhost`, not the bind address it was
given. Bound IPv4-only, on a host resolving `localhost` to `::1` first, the self-proxy never connects.

    -H 127.0.0.1  ->  GET / = 500 in 30.014 s   (twice; identical warm; log: "Failed to proxy
                                                  http://localhost:13077/ … ECONNRESET")
    -H 0.0.0.0    ->  GET / = 200 in  2.378 s

**Three things iter-10 must decide rather than assume:**

1. **Do NOT revert to `0.0.0.0`.** M221 tightened this for a real reason — a *localhost* demo was answering
   200 on the tailnet IP ("the S0 exposure lie"), and `safety.md` §3 Part 3 treats exposure as load-bearing.
   Any candidate fix must be re-run against `stack-injection/exposure_claim_guard.py`, not just against the
   readiness probe.
2. **The two host families disagree, and that disagreement IS the mechanism.** This Mac and the `billion`
   Linux VM resolve `localhost` differently. A fix proven on one is not proven. Candidates worth measuring
   before choosing: bind `-H ::1`; bind `-H localhost` and let Node resolve it the same way the proxy does;
   or force the proxy's dial. **Measure which of these actually listens where, on both hosts** — do not
   reason it out from resolver rules.
3. **Fix the READINESS PROBE too, and separately.** It polls `/` — the single route that 500s — with
   `curl -fsS --max-time 3`, so it can never observe a 30 s failure however long the outer 120 s budget is.
   `/library` answers `308` in 2 ms even while "not serving". A probe whose per-attempt timeout is shorter
   than the failure it is watching for is a probe that reports a state it cannot measure.

**Also ready, and cheap:** `FIX-M257x-autoverify-evidence-log-path` — autoverify reads
`$STACK_DIR/{demopatch,buildfail}.log`; the bring-up writes them to
`stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/`. Both exist. **iter-10 confirmed it is the whole remaining warning count: 2 FAILED, both this.** Fix the path **and** split the three
states the message conflates: **absent** (the phase never ran) / **empty** (it ran and said nothing — the
current true state, since patch output goes to stdout) / **populated**. Today all three read as "the phase
never ran", which is a cause asserted from a wrong-path absence.

### Pre-computed input for iter-08 (`FENCE-M257x-write-fence-scans-one-section-of-nine`) — CONSUMED by iter-08; NOTE: iter-08 REFUTED its central premise

Measured at the end of the iter-07 session so iter-08 starts from evidence. **This one is the gate itself:
clause 4's condition now holds, but the fence that is supposed to ASSERT it does not cover what it claims.**

    stack-core/tests/test_write_target_schema_fence.py:92
        SCORED_SECTIONS = ("stack-seeding",)

**One section of nine.** Its docstring opens *"Every schema a rext artifact WRITES TO must be a schema
rext's own migrate step CREATES"*. `stack-snapshot` — where iter-07's entire `cms` write surface lived — is
never scanned, nor are `demo-stack`, `dev-stack`, `stack-verify`, `stack-injection`, `stack-secrets`,
`clerkenstein`, `alignment`. **The fence would not have caught the defect iter-07 just fixed.**

**What iter-08 must decide rather than assume** (this is not a one-line constant edit):

1. **Widen, or narrow the claim?** The limit may have been deliberate at iter-06 — the seeders are where all
   three previous occurrences bit, and `SCORED_SECTIONS` is a tuple *designed* to be extended. But nothing
   records the choice, and clause 4 rests on it. Either outcome is acceptable; an unstated boundary is not.
2. **Widening will surface new constructs, and §8 rule 6 applies to every one.** iter-06's three recognisers
   are tuned to the seeders' shapes (`CopyRows(ctx, "<schema>", "<table>"`, the `[]struct{schema, table
   string}` literal, `var resetTables`). `stack-snapshot` writes through
   `{Schema: <const>, Table: "<name>"}` keyed literals and through `replay.Run`'s resolved value — **a
   const, not a string literal**, so a naive widening scores nothing and reports GREEN, which is worse than
   not widening. Measure what each new section actually looks like before extending the tuple.
3. **The resolved-schema case is now legitimately dynamic.** Post-`D-M257x-8` the snapshot replay's schema is
   *derived at run time*, so there may be nothing static to fence there at all — in which case the honest
   answer is that `stack-snapshot`'s write target is covered by the LIVE layer (§8's third row), and the
   docstring should say which layer covers which section. Check before building.

**Also open, and cheaper:** `FIX-M257x-academy-not-serving` is now the ONLY genuine ✗ in autoverify (the
other two are `CHECK-M257x-bringup-evidence-logs-absent`, evidence-absence rather than defects). It is
therefore the last thing between here and a green cold cycle for **clause 1**, which needs three of them.

### Pre-computed input for iter-07 (`REPOINT-M257x-cms-similarity-writes`) — CONSUMED by iter-07; kept for provenance

Measured against the live demo-1 DB + the snapshot cache at the end of the iter-06 session, so iter-07
starts from evidence. **This one is NOT a re-point of a constant — it is a capture/replay ASYMMETRY**,
and that is the whole finding.

**1. The surface, and where its four tables actually are.** `stack-snapshot/simembeddings` declares
`const Schema = "cms"` and replays `similarities` (274 rows) · `similarity_categories` (278) ·
`similarity_features` (274) · `similarity_skills` (664). On demo-1 **all four exist in `public`**, all
four hold **0 rows**, and the `cms` schema holds **0 tables**. That empty library IS the surface's own
stated root cause #3 ("a fresh demo has 0 rows there, so /library/ai-simulations is EMPTY") recurring
through a different door.

**2. Column drift: NONE.** The cached manifest's column lists match `public.*` exactly, table for table
(order differs; the manifest names its columns, so COPY is order-independent). Verified against
`information_schema.columns` on demo-1 — the same check iter-06 ran, same result.

**3. `Schema` is used for BOTH capture and replay, and they now disagree.** Capture reads **prod**,
where the legacy `cms` schema survives pending platform **M810/M710**; replay writes a **fresh stack**,
where the platform never creates it. This is exactly `D-M257x-3`'s two-states-per-row axis, and
`D-M257x-7` (iter-06) already set the precedent: **the WRITE side moves, the prod-READ side stays.**
The cached capture is dated **2026-06-29** — before cms-in-app — so re-capture freshness is a separate
question that cannot be settled from this box (no prod access; `HOST-M257x-toolchain`).

**4. The mechanism to change, and its size.** `replay/replay.go` reads `tb.Schema` straight off the
manifest at `:135/152/163/173/184`. Only **three** types implement `Replayer` (`cmd/stacksnap/adapters.go`
+ two test fakes), so widening the interface is bounded.

**The design decision iter-07 must make (do not default to the easy one).** A declared
`ReplaySchema = "public"` on the surface is the two-line fix and is *the same hand-maintained-constant
defect this milestone exists to end* — it would be wrong again at the next fold. The alternative is a
**derived replay-time resolver**: for each manifest table, if `<manifest schema>.<table>` exists on the
target use it unchanged (so taxonomy/directus are untouched); else if **exactly one** schema on the
target holds that table, remap to it and say so LOUDLY; else fail loud naming the candidates. That is
"follow the platform when it moves" implemented once, generically, and it self-heals for every future
surface. It is more work and it changes a shared engine — which is why it is a decision, not a detail.

**Two things not to miss.** (a) `stacksnap`'s pre-replay precondition probe fails first with
`probe stack schema: pg: schema "cms": schema has no columns (empty digest)` — that gate reads the
manifest schema too and must move with the resolver, or the surface still skips at rc=4 before any copy
is attempted. (b) The firewall `ParentScope` predicates embed `"cms"."similarities"` inside the capture
FILTER SQL; those run against **prod** and must NOT be re-pointed (`D-M257x-7`).

**When it lands,** `REXT_TRANSITIONAL_SCHEMAS` can go `"cms"` → empty — the second and last firing of
the debt fence's "this failure is GOOD NEWS" branch, and the act that makes **gate clause 4** claimable.

### Pre-computed input for iter-06 (`REPOINT-M257x-jobsim-writes`) — CONSUMED by iter-06; kept for provenance

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

### Pre-computed input for iter-11 (`FIX-M257x-autoverify-evidence-log-path`) — measured 2026-07-31, do NOT re-derive

**It is now the ENTIRE remaining autoverify warning count.** iter-10 measured demo-1 post-fix: `2 check(s)
FAILED`, and **both are this**. Clause 1 wants `green:true / 0 warnings` on three consecutive cold cycles;
this is what stands in the way.

```
⚠ NO demo-patch evidence at <STACK_DIR>/demopatch.log — … its ABSENCE means the phase never ran …
⚠ NO frontend-build evidence at <STACK_DIR>/buildfail.log — … its ABSENCE means the build phase never ran …
```

1. **The path is wrong.** autoverify reads `$STACK_DIR/{demopatch,buildfail}.log`; the bring-up writes them
   to `stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/`. **Both files exist there.** The message then
   asserts a *cause* from the absence — the dominant class of this milestone, inside the verifier that
   measures the gate.
2. **At the correct path both are 0 bytes**, while the patches demonstrably applied (their output went to
   stdout — visible in every `ant-academy.sh` run iter-10 captured). So the fix must distinguish **absent**
   (the phase never ran) / **empty** (it ran and said nothing — the current true state) / **populated**.
   Today all three read as "the phase never ran".
3. **Do not stop at the path.** Deriving the path from the same helper the bring-up writes it with is the
   §2 "derive it at the point of use" move; a second hand-written path constant would be the defect this
   milestone exists to end, in a third place.
4. **Carry iter-10's D-M257x-11 across:** the three states must be *reported* distinctly, not collapsed into
   one string at one severity — the same fix shape as the academy readiness probe, and the same reason
   (`CHECK-M257x-demopatch-pristine` is still open for exactly this on `pristine-ing skipped/failed`).
5. **Run the check against the broken state, live.** iter-10's own first cut of its new probe repeated the
   defect it was fixing and passed three unit tests; only a live negative control caught it.

- iter-11 (tik): **the last standing autoverify warning class is closed at its real mechanism, and the inherited hand-off was refuted on 3 of its 5 points before a line was written.** iter-10 handed forward *"2 FAILED on demo-1, both the evidence-log path"*; the bring-up's OWN verdict, on disk the whole time, said `warnings:1`. The 2 came from a **standalone re-run of the same script pointed at the workspace root** — same tool, same stack, different vantage, different answer, five hours apart, timestamps never compared. The path was **not** wrong (`up-injected.sh:2550` passes `STACK_DIR="$STACK"`, correct by construction); autoverify has distinguished **absent / empty / populated since M256** with four tests pinning it (and **empty is the HEALTHY state** — M217 writes the log only from failure branches); and the message already named the alternative, *"(or STACK_DIR is not the bring-up's `$STACK`)"* — quoted truncated at the em-dash, §5 rule 10's third occurrence. What was real is bigger: **`STACK_DIR` was a hand-supplied path with no derivation and no validation**, in a script that already derives its offset from `--project`. All three failure modes had already happened — one caller right, `dev-stack:298` passing **nothing** so every dev-N bring-up silently skipped the cheap-wins **and wrote no `autoverify.json` at all**, and any standalone run reporting missing evidence for a stack that has it. Now **derived** (`target_resolve_stack_dir`, §2 at the point of use), derivation **wins** over an inherited value with a named mismatch, and the two receipt asserts are gated on the project's **TYPE** (`target_is_demo_project`) rather than on whether a caller set a variable — so dev gains the transcript + verdict it never had. Live on `demo-1` from its own pinned clone: a run given **no path at all** now reads `✓ demo-patches` + `✓ frontend builds` → **`warnings:0 / green:true`** (was 2 / `green:false`); the bring-up's own correct call stays **silent**, which clause 1 requires. **The first cut cried wolf** — it warned on a missing derived dir, passed 10 targeted tests, and went red across **18 pre-existing fixtures**; reasoned down to a named `·` skip that prints the path consulted, because `mkdir -p` at `:226` means a missing dir never describes a real bring-up. 10 tests, 6 mutants RED (M6 **two-point**, since the no-fabrication invariant is enforced twice and a single-point mutant is a no-op rather than a survivor) **plus an unmutated control that goes GREEN**; 224/224. Three protocol rules promoted: §5 rule 12 (*say which INVOCATION produced the number*) and two §8 rule 5 addenda (*read the collected count with the exit code — pytest exits 5 on an empty `-k` match*; *a mutant that changes nothing is not a survivor*) — see iter-11/progress.md

- iter-12 (tik): **clause 1 is 0 of 3, and the reason is that the platform deleted the GraphQL router while we were iterating.** The iter opened on `FIX-M257x-vmram-gib-unit` (real in two halves — the pre-flight reported the integer **floor**, `11` for a measured 11.67 GiB, and its remediation told operators to set **12 GB decimal**, ~11.2 GiB, which can never clear a 12 GiB floor; the *comparison* defect I had also drafted was **refuted by mutation N2** — for an integer floor, `floor(x) >= m` iff `x >= m` — and the refutation is written into the code comment rather than dropped), cost the predicted 14 restaled `file:line` citations and repaired them with the guard's own `--fix`, and closed `DOC-M257x-claude-md-knob-count` **REFUTED** (`CLAUDE.md:307` and `demo/README.md:153` already say 30). Cold cycle 1 then went RED on directus `exit(1)` — root cause: **iter-05's cold-start fix had been applied to the DEV twin, not the DEMO twin it was measured on**, and its test passed *because it tested the twin that was fixed* (the "reports without measuring" class, 9th occurrence this session); fixed + fenced at rext `fast-build-m257x-iter-12b`. Before letting the rebuilt cycle stand, the gate's own first clause — *"against platform @ **origin HEAD**"* — was checked against the clone, which was **3 commits behind**, while the bring-up's freshness check printed `PIN-DRIFT` naming a staler pin: **a freshness check that compares to a pin cannot detect a stale clone, and reports the two identically** (promoted to protocol §3). Origin HEAD `2adcf71` (2026-07-31 15:58, mid-milestone) had **dropped the WunderGraph/Cosmo federation router outright** — service, `repos.yml` entry and clone — with GraphQL now served straight from `backend` at `:8082/graphql/query` (**the path changed too**, `/graphql` → `/graphql/query`, the half a hostname-only re-point would silently miss). The `graphql` *profile* survives on seven services, so nothing about the profile wiring warns. Measured, not inferred: `gen_injected_override.py` against the origin-HEAD compose returns **RC=0** — it does not notice — and still emits `depends_on: graphql` plus two `WUNDERGRAPH_SSR_ENDPOINT=http://graphql:8080/graphql`; `docker compose config` then returns **RC=1**, *"service \"hiring-app\" depends on undefined service \"graphql\": invalid compose project"*. **Clause 1 was not attemptable at origin HEAD at all** — the bring-up dies at project validation before a single build, so no amount of cycling could have produced a green one. This is the **fourth** occurrence of the milestone's founding class (skiller → app, skillpath → app, jobsimulation → app, now router → app) and the **first caught before it caused a mystery**, because the gate named a ref. Scope-creep tripwire fired on the third line of investigation; the re-point routed whole to iter-13 with three named handlers, and the two-consecutive-invalidations re-scope trigger recorded at **occurrence 1 of 2** — see iter-12/progress.md
