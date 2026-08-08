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
| ~~`SURVEY-M257x-iter143-bare-orphan-bucket`~~ **CLOSED by iter-143** | Censused (384 refusals: **88 % `orphan-no-referent`**, not ambiguity) and the head inference was built, hand-read at 100 % and **REFUSED at 77.3 % precision / 32.1 % recall**. The comment's *"where the ports are"* is retracted **as an explanation**: of 39 false admits, **21 ports (loud) vs 16 WRONG-HEAD (silent)**. Shipped instead: the census decomposition + a 7-suffix `_CODE_SUFFIX` widening, both no-inference (reach **861 → 875**). | iter-143 ✅ |
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

- iter-13 (tik): **the router re-point — clause 1 goes from *not attemptable* to *attemptable*, and the class is fenced.** Re-measured iter-12's routed list first (the standing rule) and found **~14 files, not 6** — the residual included `up-injected.sh`'s **six** `$((5050+OFFSET))/graphql` sites and, worse, `stack-verify/lib/readiness.sh`, whose probe introspects *"the federated supergraph at :5050"* and would therefore have reported the API **down while it was up**. **Two things moved, not one:** `graphql:8080` → `backend:8082` **and** `/graphql` → `/graphql/query` — a wrong host refuses loudly, a wrong path resolves, connects and 404s (the latency-budget *fast-failing fetch* signature), so a hostname-only re-point would have shipped green. The fix is therefore a **derivation, not a correction**: `browser_graphql_endpoint()` is the one definition all six sites call, `BACKEND_SERVICE`/`GRAPHQL_PATH`/`SSR_GRAPHQL_ENDPOINT` likewise, and the **two test files that had duplicated the literal now import the constant**; `gen_tailscale_serve`'s `("graphql", 5050)` row is **deleted rather than re-pointed** (backend was already fronted, and fronting a dead port yields a trusted-cert endpoint that always refuses — worse than absent, because it looks configured). **`FENCE-M257x-iter13-compose-service-exists`** — every emitted `depends_on` must name a service the platform compose at the ref in use defines — was **watched going RED first** (mutant exits 1 naming the vanished service; unmutated control exits 0), **reports what it checked** (*"2 depends_on target(s) … ['backend','postgresql']"*) and **fails closed** if it finds blocks but extracts zero targets. Its **placement was the real design decision**: the first cut sat inside `build_lines` and turned **16 unit tests into errors**, because it was asserting platform completeness against deliberately-truncated one-service fixtures — §8 rule 4 — so it moved to `main()`, where cfg *is* the real resolved compose. **iter-12's own live negative control flips RC=1 → RC=0** (16 services resolve). Attribution came within an inch of being wrong: the `git archive` scratch baseline reported demo-stack **3** failures where the same code in place reported **7**, and the tell was the **skip count, 26 vs 2** — live-clone tests resolve by relative path and *skipped* in scratch, so **a control that silently skips the tests you are attributing is not a control**; reconciled, the pre-existing set is exactly the 7 of `CHECK-M257x-live-clone-suites-red`. Mine were **26 injection · 8 core · 17 demo-stack · 2 dev-stack**, all fixed, **every section back to baseline exactly, zero regressions**. Paid the predicted side-effect (**22** restaled `demo-up-defaults.md` citations, repaired with the guard's `--fix`) and found `CLAUDE.md` stale on **two** counts — the router, and a *"3 subgraphs"* claim the cms-in-app merge had already reduced to one. rext **`fast-build-m257x-iter-13` (`4414527`) verified on origin**, tag file + consumption clone re-pinned. Not landed: the freshness-vs-origin fix (phase item 4, a 4th line — routed to iter-14), and **35 files / ~128 corpus hits** still describing the router (routed to clause 5) — see iter-13/progress.md

- iter-14 (tik): **GATE CLAUSE 1 IS MET — three consecutive cold cycles, 0 warnings, against platform origin HEAD.** `demo-down 1 --purge` → `demo-up 1`, three times, each teardown verified down to **0 containers** and each bring-up consuming rext at `fast-build-m257x-iter-13` **from origin**: `warnings:0 / green:true` at `11:43:02Z`, `11:53:04Z` and `12:03:33Z` (~11 min per cycle, no red in between). **The timestamps are load-bearing, not decoration** — a green verdict left on disk by an earlier cycle reads exactly like a fresh one, which is precisely M236's green-gate defect (a UTC `ts` parsed as local aged a stale verdict as fresh), so three distinct monotonically-advancing timestamps each following a purge is what makes "three consecutive" a measurement rather than a claim; checked in at `evidence/av-cycle{1,2,3}.json`. **The live run proved what iter-13's `compose config` RC=0 could not**: a valid project is not a stack that comes up, and iter-13 had re-pointed four surfaces no static check exercises — the SSR origin, the readiness probe's port **and path**, the tailscale front list, the clone pin. All held, and the arithmetic confirms the shape: **15 containers, not 16**, exactly the deleted router, so the re-point removed a service rather than renaming one. Two open items were **observed rather than assumed**: the first teardown's *"service `graphql` has neither an image nor a build context"* came from the override **left on disk by iter-12's killed cycle**, and the next two teardowns — against overrides regenerated by iter-13's code — were clean, confirming the diagnosis by the following observation instead of asserting it from the first; and `CHECK-M257x-demopatch-pristine` reproduced identically three times as `WARN … sha DRIFTED … anchor intact (1x)` then applying, which is `demopatch-spec.md`'s documented chain rule + self-healing freshness gate (*the anchor is the contract*), not a defect — the item narrows to how those three states are **reported**. **2 of the gate's 5 clauses now hold** (1 and 4); clauses 2 (Playthrough suite), 3 (migration-status map) and 5 (KB-fidelity) remain. `demo-1` is deliberately **left UP, green, 15 containers** — clause 2 must run on the clause-1 stack. See iter-14/progress.md

- iter-15 (tik): **gate clause 2 is measured for the first time — 20 live / 10 failing / 1 unimplemented, NOT MET — and its largest root cause is that the per-stack Directus was never actually serving.** The suite ran unscoped on purpose (the ptreport gate is **binding only on a full run**, `run-playthroughs.sh:300-307`, and its own `:292` residual says a broad `--grep '@pt'` is still graded advisory); the denominator reconciles exactly, 30 manifest `playthrough:` ids == 30 spec `@pt:` tags, `diff` identical, ptreport's 31st row being the one declared in-manifest `will-not-build` — though the first cut of that count returned **0**, from a leading-dash anchor the field does not carry, which is §5 rule 3 on *the same number the protocol already records a false absence for*. Every skill-path failure rendered **"Content not found."** while the row sat published in `directus.skill_paths`, and an anonymous read 403'd **with the grant row present** — so the answer was not permissions but the bring-up transcript, which had said the same thing on **all three of clause 1's cold cycles**: `directus=skipped(error)`, `advance identity sequences on directus.sequences: column "sequence_catalog" of relation "sequences" does not exist`. `boot_directus_step` runs **only on a successful replay**, so the failure silently cost the post-replay restart as well as the rows — one cause, two symptoms, and the second is what made it look like a permissions bug. **Two independent defects, both landed.** (1) A **plan-dependent** query that RAISES: `pg_get_serial_sequence(quote_ident($1)||…, a.attname)` references only `a.attname` and the parameters, so it is a restriction clause on `pg_attribute` the planner may push below the joins, evaluating it against columns of *other* relations where the function does not return NULL but raises — and literal-substituted in `psql` it **succeeds** (custom plan) while `PREPARE`d + `EXECUTE`d six times it **raises** (generic plan), which is how pgx sends it and precisely why reproducing it by hand "proved" the SQL fine. Fixed as a **barrier, not a correction**: the relation is resolved to one OID behind two `AS MATERIALIZED` fences and the function is handed `tgt.reloid::regclass::text`, the very OID the column list came from, so the two arguments cannot name different relations under any plan (§8 rule 4); 7 clean `EXECUTE`s after, and non-vacuous. (2) The structure script emits **two thirds of a `serial`** — `CREATE SEQUENCE` + `DEFAULT nextval(...)` and never `ALTER SEQUENCE … OWNED BY …`, so the edge `pg_get_serial_sequence` resolves through is absent and M256's refusal correctly fires; measured **8 provisioned user collections unowned** while every `directus_*` table Directus bootstrapped itself was owned, and reconciled on the **APPLY** side from the target's own catalog rather than in the capture, because a capture-side fix reaches only re-captured caches and this host cannot re-capture. The replay's refusal is deliberately **left alone** — healing inside the guard would make it a probe that satisfies itself. Live: replay **rc=1 → rc=0**, **0 → 11 986 rows**, two collections **403 → 200**; closes `FIX-M257-stacksnap-directus-sequences`, carried since M257 iter-02. **And the honest half: the clause-2 metric did not move.** A re-run scored 17/31, which is **not** a regression — it ran without `--reset` after a mutating run, so its three net-new failures are all `onboarding.*` negative controls asserting *"onboarding is INCOMPLETE"*, the stale-world state the runner's own header forbids; two runs of two different worlds are not comparable, and the comparable measurement is the one taken at the layer that changed. **At least three further causes remain, named separately rather than lumped**: a `directus_versions` 403 that still blocks both skill-path Playthroughs (though `/versions` answers 200 anonymously and the grants exist), a **106-occurrence** content-model shape drift (`cannot unmarshal string into … library_category of type struct{ID uuid.UUID…}` — the app reads an expanded relation, Directus returns the raw id), and 5 manager-side reads where the seeded hero is simply absent. Side, measured and **routed not landed**: `run-playthroughs.sh:161`'s readiness gate announces **`✓ fake-FAPI ready (HTTP 000000)`** on a connection that never happened, because `curl -w '%{http_code}'` writes its own `000` on failure *and* exits non-zero so `|| echo 000` appends a second — one **verdict-flipping** site plus three siblings where it only degrades the message. 6 mutants RED, control GREEN; §5 gains **rule 13** (*a catalog query correct in `psql` can be broken in the program — parameters change the plan*) — see iter-15/progress.md

### Routes opened by iter-15 (clause 2's remaining causes — four error strings, four handlers)

| item | why | target |
|---|---|---|
| `FIX-M257x-iter15-directus-versions-403` | The 4th cause and the one still blocking both skill-path Playthroughs. `/versions` answers **200 anonymously** and the read+create grants sit on the public policy — so the refusal is about **how** the cms domain asks, not whether the grant exists. | next tik |
| `FIX-M257x-iter15-library-category-expansion` | **106 occurrences**, the largest single class. `app` reads `library_category` as an EXPANDED object; Directus returns the raw id string. Suspect the replayed `directus_relations` / alias-field rows against what post-cms-in-app `app` expects. | next tik |
| `CHECK-M257x-iter15-manager-reads-empty` | 5 Playthroughs assert the seeded hero is among the results and get **0** (`workforce-funnel`, `workforce-succession`, `workforce-org-feedback`, `activity-drilldown`, `assign-and-track.UC2`). Do **not** assume it shares a cause with the content layer. | later tik |
| `CHECK-M257x-iter15-orgadmin-role-create` | A 60 s `waitForURL` timeout on the role-detail route; no error text captured yet. | later tik |
| `FIX-M257x-iter15-readiness-000000` | `run-playthroughs.sh:161` reports READY on curl's own failure code. One **verdict-flipping** site; `stack-verify/lib/services.sh:132` + `readiness.sh:151,185` carry the same double-source but only **degrade the message** — the `down` verdict still comes out right. The distinction is measured and paid for. | later tik |
| `DOC-M257x-iter15-autoverify-blind-to-content` | Clause 1 is met **as written** and the stack it certified had **no served content**. Record what a green autoverify does and does not assert, and make the Directus probe read an **item**, not a registry count. | clause 5 |
| `CHECK-M257x-iter15-stack-verify-red` | **`stack-verify` is RED and was not in anyone's baseline: 11 failures + 1 error of 224** — all in the classes the router deletion moved (`TestContainerLivenessM257`, `TestFapiProbeLadderM257`, `TestFrontendTierRegistration`, `TestOffsetAwareness`, `TestOffsetMatrixSweep`, `TestServiceScopeFilter`). **Attributed by measurement, not by reasoning:** the same six classes in the SAME consumption clone at the SAME path score `failures=11, errors=1` at **both** `fast-build-m257x-iter-13` and `fast-build-m257x-iter-15` — pre-existing, zero regressions from iter-15. The finding that survives is §5 rule 8 in the section that owns the bring-up's own probes: red in a suite nobody re-ran. Recorded here as the section's baseline. | next tik |

- iter-16 (tik): **the two bring-up verdicts that graded themselves green are fixed at the site that produced gate clause 1's false greens — and the suite was not silent about the defect, it was arguing for it.** Re-read both routed-forward findings at source first, and **the ledger's filing of RF-4 was wrong**: it is booked as a `dev-stack` finding, but `up-injected.sh` runs `dev-stack/dev-setdress.sh` **verbatim** via `--stack-type demo` (the M20 convergence invariant, with its own test forbidding a demo-only fork), so the `*)` arm that turns any unclassified replay rc into `skipped(error)` and returns 0 is not *like* the clause-1 signature — **it is the site**, and `up-injected.sh:2262`'s `if ! … "⚠ set-dressing did not fully complete"` had been **dead code for its entire life**, unreachable for the one state it was written for. The summary field had been honest since fix16 (`directus=skipped(error)`, on all three cold cycles); what lied was the **sentence it sat inside** and the **exit code after it**. Now rc=4 (unprovisioned) and rc=5 (cache-miss) stay documented degradations that still report `set-dressed` — correctly, the stack *is* dressed to the extent the environment allowed and each already prints a named operator fix — while **any other rc is an unclassified ERROR**: `replay FAILED`, counted, the surface **NAMED**, the verdict word `set-dress INCOMPLETE`, and a **new exit code 3** deliberately distinct from `die`'s 1, because the seed floor still ran (M20 atomicity) so this is a **report, not an abort**; the caller names rc=3 apart from the generic warning — whose named fixes, *provision the stack* / *capture the cache*, are exactly the wrong advice for a replay error — and says the consequence aloud: **the failed surface is EMPTY, which downstream reads as "the content layer 403s"**, the sentence that would have saved iter-15 its permissions detour. **Three existing tests were asserting the tolerant verdict on the intolerable case, and two said so in their own comments** (*"here a generic exit 1 — fix16 reserves 4/5 for the classified cases"*); the third's rationale, *"cache-miss vs firewall-error are the same exit 1; both degrade gracefully"*, **was true when written and stopped being true at fix16**, which split the graceful cases out into 4 and 5 — §8 rule 3 exactly, a fence pinning the current shape of the drift and converting the bug into a contract. The prod-safety never-capture test was **kept** at exit 1 on purpose: the unclassified error is the strongest place to pin it. **RF-1** ported `migrate-demo.sh`'s M215 F8 fix to the dev twin that never got it and was strictly worse — every non-zero atlas exit graded *"had migration warnings (non-fatal — see atlas output)"* while `>/dev/null 2>&1` had **thrown that output away**, and an absent clone logged `✗` and continued **recording no failure**, so *"the derived migration set applied"* printed after migrating nothing; the **parity fence could not have caught it** because it walks a hand-maintained guard enumeration missing all four of F8's guards (routed as `CHECK-M257x-iter16-parity-fence-hand-maintained` — §2's own class turned on the fences), and `test_absent_service_clone_is_skipped_not_fatal` was green over the defect on a rationale that **iter-02 had invalidated** (once the set became derived, the isolated cold proof's tree has no `repos.yml` either, so `MIG_PAIRS` is empty and the `continue` protects nothing there). **11 mutants, each `bash -n`-gated, control run before AND after: 10 declared-RED all killed, 1 declared-GREEN no-op survived** — the no-op is the one that makes the other ten mean something, and it is promoted to protocol **§8 rule 5** as a required positive control (ten REDs cannot distinguish a discriminating fence from a merely brittle one). The demo-stack suite's **+1** was `test_setdress_is_non_fatal`, which string-matched the **mechanism** (`if ! env `) rather than the **proposition**; re-written to assert that the invocation's exit is *captured* and never becomes `|| exit`, scoped to the extracted **block** because the surrounding script legitimately contains `|| exit` for fatal steps (§8 rule 4). Paid the predicted side-effect — **6** restaled `demo-up-defaults.md` citations, repaired with the guard's own `--fix`. Every section back to baseline exactly: dev-stack **OK 125** (+3 tests) · demo-stack **7F of 1030** (+1) · stack-core **14F** · stack-injection **OK 286**. Clause 1's re-prove now has its precondition — see iter-16/progress.md

- iter-17 (tik): **gate clause 1 was never met — the honest instrument fails it on the FIRST cold cycle, and its three checked-in verdicts are formally WITHDRAWN.** The re-pin was the precondition and it was real: `.agentspace/rext.tag` and the `stack-demo` consumption clone both still read `iter-15`, which **predates the harden pass**, so `probe_directus_serves_content` was not in what a stack consumed; both re-pinned to `fast-build-m257x-iter-16` (`c63d981`, verified on origin, `merge-base --is-ancestor` confirming it CONTAINS harden-1). **The positive control was not a formality** — run against the standing hand-repaired demo-1, autoverify showed no `directus-serves-content` line at all and the only textual match was the registry check's own cross-reference to it; running `verify.sh` directly resolved it (autoverify prints only the `✗` rows, so a passing probe is invisible **by design**) and the probe **ran and passed**: `anon GET /items/task_sub_checks served a real item`. *"I did not see it in the log"* and *"it did not run"* are different findings. A wrong entry point then cost a cycle and was caught by **arithmetic, not suspicion**: `rosetta-demo up 1` returned "up" in **30 s** against iter-14's measured ~11 min, its log reading `3 services, profile='base'` — the bare compose bring-up, not `up-injected.sh`; nothing about the resulting stack would have LOOKED wrong. Through the real path, cycle 1 read **`{"green":false,"warnings":1}`** at `16:03:35Z` on `✗ directus-serves-content — anon GET /items/task_sub_checks -> 403, the running Directus holds the content but serves it to nobody (public-role grants never applied)` — in the same run where `✓ directus.directus_collections = 21` passed. Both checks are correct; only one of them was ever asked before (§5 rule 14, **REGISTERED is not SERVED**). **The good half is real**: iter-15's replay fix **works through the real bring-up path** for the first time (`directus=replayed`, 11 986 rows, rc=0) — its own hand-off flagged this as unsettled because the auto-provision path fires only on a bootstrapped-GAP schema, which a *purged* stack has and the standing demo-1 no longer did. **The defect hiding behind it:** the per-stack Directus **bootstrap FAILED**, the pass announced *"skipping local content (the stack stays on the prod-read path)"* — and then replayed 11 986 rows into that same local Directus anyway, so the closing verdict asserts `content:prod-read` **and** `directus=replayed` in one sentence and exits 0. The 403 follows mechanically: content in the tables, but the system schema whose grants serve it to an anonymous reader was never bootstrapped. **And why it failed is currently unknowable — `dev-setdress.sh:264` runs `node cli.js bootstrap >/dev/null 2>&1`, the THIRD occurrence of RF-1's exact shape** (after `migrate-demo.sh`/M215 F8 and `migrate-dev.sh`/iter-16, fixed hours earlier); re-running the identical `docker run` by hand now **succeeds** (`Database already initialized, skipping install`) because the replay has since created the `directus_*` tables, so the state that failed no longer exists — the diagnosis existed for one moment and went to `/dev/null`. Cycles 2 and 3 deliberately **not run**: three-consecutive-green is a conjunction, cycle 1 is red, and running them to fill a table is the instinct that produced the withdrawn three. **Gate corrected 2 of 5 → 1 of 5** (clause 4 only) — see iter-17/progress.md

- iter-18 (tik): **gate clause 1 is MET again, this time by an instrument that can see served content — and the 403 turned out to be a RACE between two Directus bootstrappers.** The hand-off said the diagnosis could only arrive through an ~11-minute cold cycle because the failing state had healed; **measuring cost four minutes and refuted it on five points before a line was written**: the identical `docker run … node cli.js bootstrap` **exits 0** against a throwaway empty schema (so the command is fine and the failure is context); demo-1's `directus` schema held **all seven system tables and 87 migrations** (refuting iter-17's *"never bootstrapped"*); the public **policy, permission and `role IS NULL` access rows all existed** (the thing said to be missing was present); `docker restart demo-1-directus-1` **alone** flipped the anon read **403 → 200**; and `docker inspect` showed the image's own `CMD` is **`node cli.js bootstrap && pm2-runtime start`** — **the compose service is a second bootstrapper**, whose timestamped log crash-loops `schema "directus" does not exist` until set-dress's CREATE SCHEMA lands and then runs the 87 migrations itself at `16:02:28.98→29.40`, matching `directus_migrations` to the millisecond. **The 403 came one phase later and not from the bootstrap at all:** `boot_directus_step` — the post-replay restart that exists precisely so Directus re-introspects collections registered after it booted — was gated on `DIRECTUS_PROVISIONED`, i.e. on **a different step's exit code**, so a lost race cancelled a restart the replay had just made necessary (§5 rule 13's blast-radius note, third occurrence). Fixed as a **post-condition, not a winner**: the pass now asks the *database* whether the sentinel system table is there and reports PROVISIONED either way, **naming the winner** — the rejected alternative, overriding the image's `CMD` so our racer wins, is a hand-maintained contradiction of an upstream default, i.e. the class this milestone exists to end. The restart is **decoupled** and gated on its own leg; the bootstrap output is **captured, classified and echoed** (the **third** occurrence of the `>/dev/null 2>&1` masking class, after M215 F8 and iter-16) with its **sibling CREATE SCHEMA leg swept in the same pass** (§5 rule 9) and the sentinel probe now saying *why* it read 0 (§5 rule 12). `content_mode` gained a third state because two of them were a lie: `gen_injected_override.py:580` re-points `DIRECTUS_BASE_ADDR` at the stack's own Directus whenever the local-content service is emitted and a failed provision does not undo it — so *"the stack stays on the prod-read path"* was **false**, and the closing line printed `content:prod-read` **and** `directus=replayed` in one sentence (closing `CHECK-M257x-iter17-setdress-verdict-contradiction` by derivation, not by rewording). **Three existing tests required that false claim** — one of them said so in its own name — the second occurrence in this milestone of *"the suite was not silent about the defect, it was arguing for it"*; and the iter's own new fence went **red on its own explanation** until it stripped comments, §8's *allow comments unconditionally* catching the milestone that wrote it. **Live: three consecutive cold `--purge` cycles at `warnings:0 / green:true`** (`16:50:14Z`, `17:00:45Z`, `17:10:44Z`, anon read 200 after each, rext consumed at `fast-build-m257x-iter-18` from origin, platform origin HEAD `2adcf71` re-checked at open and close). **Cycle 3 is the causal proof and it arrived by luck — said so in writing:** cycles 1–2 won the race and would have been green on the unfixed code; only cycle 3 lost it, printed the losing run's `42P01` diagnosis (the one that went to `/dev/null` for two iters) and stayed green. **The race is nondeterministic — 2 of 3 one way — which dissolves the last puzzle in the thread:** it is why iter-14's three cycles read green and iter-17's single cycle read red on the same code. 7 new tests + 3 corrected; **7 mutants, 7/7 matching their declared expectation** (6 RED all killed + 1 GREEN no-op survived, each `bash -n`-gated, control green before and after). Every section at baseline: dev-stack **OK 132** (+7) · demo-stack **7F of 1030** · stack-core **14F** · stack-injection **OK 286**. §5 gains **rule 15** (*"it only reproduces in the full pipeline" is a claim about the failing step's INPUTS* — plus: a nondeterministic defect makes a green run weak evidence; record which path each run took). **Gate 1 of 5 → 2 of 5** — see iter-18/progress.md

- iter-19 (tik, closed-no-lift): **clause 2 re-measured through an instrument that now works — `20 / 10 / 1`, and the failing SET is byte-identical to iter-15's, which refutes the attribution this milestone made one iter earlier.** iter-18's close routed `FIX-M257x-iter15-directus-versions-403` and `FIX-M257x-iter15-library-category-expansion` forward with an explicit caveat that both had been diagnosed on a stack whose Directus served nothing and might be downstream of it; measured on the stack iter-18 proved green (three cold cycles, `anon GET /items/task_sub_checks` = **200**), with the real `--reset` path and the full suite (the ptreport gate is binding only on a full run), the answer is **no** — the caveat is **withdrawn** and both are independent defects the next tik can work directly. The comparison is a **`diff` of sorted ids**, not two summary lines: `20/10/1` twice could have been ten different failures. **Two findings the question did not ask for.** (1) **iter-15's number is now REPRODUCED rather than re-asserted** — its own re-run had scored `17/31` and the gap was *explained* by a missing `--reset` rather than tested; a real reset on a differently-built stack lands on the same ten ids, so the explanation has a confirming observation and the reset-vs-additive discipline is load-bearing rather than merely stated. (2) **The largest cause spans TWO fields and the count moved**: `119 × JobSimulation.data.library_category` **plus** `11 × JobSimulation.data.job_position` (iter-15 named the first alone, at 106), so a fix aimed at one field name leaves the other standing — whether they share a root cause is deliberately **not concluded**. Also: `run-playthroughs.sh:118` calls a bare `stackseed` that is not on PATH here (the bring-up builds it into the stack's own `bin/`), so the suite could not reset at all until the path was supplied by hand — the milestone's own §2 shape inside the instrument that measures clause 2, routed rather than fixed mid-measurement. No source change; gate unchanged at **2 of 5** — see iter-19/progress.md
- iter-20 (tik): **gate clause 3 MET (2/5 → 3/5)** — the migration-status map landed, completeness derived from git history (found 5 services the corpus never knew), fenced against `repos.yml` both ways and watched RED in each — see iter-20/progress.md
- iter-21 (tik): clause-5 sweep landed (~50 claims, 19 files, branch brought level with main) but the clause did NOT close — the first FULL read of all 40 in-scope files returns 21 blockers, refuting the 11→5→2 curve as a sampling artefact — see iter-21/progress.md
- iter-22 (tik): applied iter-21's 21 enumerated clause-5 blockers — **all 21 anchors verified, but TWO of the CORRECTIONS were false** (`JOBSIMULATION_RPC_ADDR`/`CMS_RPC_ADDR` still address the husks at origin `2adcf71`, deliberately, per `app/main.go:1196-1202` "until the M809 re-point") — traced to a false corpus line cited as authority (`backend.md:175`), which was the 22nd blocker. Then the **full-read re-audit (40 files / 5 sub-agents / 7,700+ lines, `wc -l` positive control) returned 53 blockers, not 21** — the inherited residual was 2.5x under-counted, and every batch's dominant cause was previously unnamed (the dropped `local_*` mirrors; "merged" over-applied to "gone from compose"; the post-M809 end state written as current; an uncaught Next 15->16 upgrade). **93 edits across 24 files, 93/93 anchors matched exactly once**, closing 34 of 53; ~19 routed. Clause 5 NOT met — see iter-22/progress.md
- iter-23 (tik): applied the enumerated clause-5 residual — **19 handed blockers + 2 found at re-derivation, 52 exactly-once edits across 13 files, 52/52 matched** — and re-derivation caught a shape the iter-22 rule does not: **a correction that is TRUE but INCOMPLETE** (the colony pin is split three ways across six live services, not two ways across four; applied verbatim it would have replaced one incomplete claim with another that read as freshly verified). `hiring.md` was re-grounded at its **thesis**, not only its rows — the doc's stated purpose is to name the one table that feeds the recruiter score and that table is DROPPED, so six row fixes would have left the argument pointing one way and the tables the other. **The iter's largest finding came from writing a correction rather than from auditing:** working out which service actually reads Directus at HEAD showed it is **`backend`, in-process** (`cms_reader_switch.go` — *"no internal traffic to a standalone cms"*; `main.go:971-973` fatals without `DIRECTUS_BASE_ADDR`), while rext's `--local-content` cutover re-points only `cms`, with a test **explicitly asserting `backend` must not carry it**. Measured on live `demo-1`: `cms` → `http://directus:8055`, **`backend` → `https://content.anthropos.work` with an empty token** — the per-stack Directus serves a consumer that no longer reads, and the reader is pointed at prod anonymously. That is the founding class wearing a new face — **a service NAME in a consumer list, not a schema name** — and nothing errors, because the named container is still real and still running; for the third time this milestone the suite was arguing for the defect. Routed (`FIX-M257x-iter23-backend-directus-not-repointed`), deliberately **not** attributed to clause 2's open causes without measurement. Second planned line landed too: **22 dead `:5050` sites** across 10 `corpus/ops/**` files incl. `run_guide.md`/`setup_guide.md`/`staging-bringup.md` (the onboarding path) and the `:15050` tailscale front row iter-13 deleted; `staging-clerk.md`'s allowed-origins lists annotated rather than edited (a transcript of external state is corrected by annotation). All 5 corpus guards green. §5 gains two rules: *re-derive the ENUMERATION, not just the values*, and *after a fold, grep the tooling for the folded service's NAME as a value*. Clause 5 still NOT MET — the list is exhausted, the corpus is not measured; the 40-file full re-read is next — see iter-23/progress.md
- iter-24 (tik): **the per-stack Directus was serving a consumer that no longer reads — `backend`'s Directus 403 class goes 96 → 0, proven live.** iter-23 routed this as a candidate; the re-survey proved it, and **one probe could not have done it**: `backend`'s log held 96 Directus lines, all 403 (`directus_versions` killing `publicSkillPaths`/`getSkillPath`/`getOrCreateSkillPathSession`, `library_categories` killing `libraryCategories`), while the **local** instance served `library_categories` **200** anonymously and **prod** served it **403** — backend's 403 set matched *prod's* answers, not the local one's. So not a missed grant: a client pointed at the wrong server (§5 rule 7 — iter-18's "does the local Directus serve?" is satisfiable by this broken world; asking **both ends** is what discriminates). It also explains rather than contradicts iter-19: the 403 **is** independent of the *serving* defect, because it is a **pointing** defect one layer up. Mechanism: `DIRECTUS_DATA_CONSUMERS = ("cms",)` in **both** twins, correct at M23 and dead at cms-in-app — `app/cms_reader_switch.go` swaps backend's content reader to the **in-process** cms server (*"no internal traffic to a standalone cms"*) and `main.go:971-973` fatals without `DIRECTUS_BASE_ADDR`. **Nothing errored, and that is the point:** a stale *schema* name fails loudly at 42P01; a stale **service name in a consumer list** fails silently, because the list still names a real, running container. `cms` **stays** in the list (its husk is still started until platform M810; messenger still calls it) — the list shrinks on an observable event, not a guess. **Third occurrence of the suite arguing for the defect:** `test_only_cms_is_repointed_not_other_services` asserted `backend` must NOT be re-pointed and would have failed on the fix (§8 rule 3); replaced by three tests that split the proposition. **And the dev twin's fixture could not name its own subject** — it called the app service `"app"` where the compose service is `backend`, so the dev list could have been wrong and every dev test still passed; `backend` added, `"app"` kept as a genuine non-consumer. **6-mutant battery, 6/6 matching declared expectation — 4 RED + TWO declared-GREEN tuple-reorder controls that survive** (M2, the dev-twin RED, would have survived without the fixture fix, and four demo REDs would have certified a fence that never looked at the other twin). rext **`fast-build-m257x-iter-24` (`f9ac72f`) verified on origin**, tag file + consumption clone re-pinned. Live on `demo-1` through the pinned clone: regenerated override adds exactly one line, `backend` recreated, **`libraryCategories` and `publicSkillPaths` now return real replayed data** and the Directus log line count is **0**. Closes `FIX-M257x-iter15-directus-versions-403` (58 occurrences, carried since iter-15) at a mechanism that was never `directus_versions` permissions. Clause-2 lift **deliberately not predicted** — it fixes one of at least four causes — see iter-24/progress.md
- iter-25 (tik, closed-fixed-partial): **clause 2's instrument can now reset itself — and the fix needed two passes, which is the finding.** `run-playthroughs.sh` called a **bare `stackseed`** that is not on PATH and never was: the bring-up BUILDS it into each stack's own `bin/`, deliberately, so a stack runs the tooling at ITS pinned tag. The failure shape is the dangerous one — `command not found` on the reset line, then the run **continues** into a suite measuring a world it did not reset (iter-15 lost a measurement to exactly this; iter-19 had to hand-supply the path). Fixed as a **derivation** from `N` like `OFFSET`, with the reset arm now **refusing** (`exit 2`, naming the path consulted) because *"the reset step did not run"* and *"it ran and did nothing"* must not print the same way (§5 rule 12); `PT_STACKSEED` stays as an escape hatch. Live negative control exits 2 before touching the world; the positive control resolves in a consumption clone and correctly reports absent in the authoring copy, which has no `stacks/`. **Then the live run refuted the fix's completeness within a minute:** *"line 180: stackseed: command not found"* and *"line 210"* — **two more bare calls** (roster export, cockpit manifest), both **non-fatal by design so they degrade SILENTLY**, leaving the roster the fake-FAPI serves and the manifest the cockpit reads describing the PREVIOUS world after a reset swapped it. The first pass had **cited §5 rule 9 in its own commit message** and swept one of three sites; its syntax check, negative control and positive control were all green over both defects, because all three exercised the same arm. Fixed behind an `-x` guard (fail closed where failure invalidates the measurement, degrade where it does not). **The re-measure did NOT land and no number is claimed.** The reset itself worked for the first time without a hand-supplied path (66 audited writes, 55 729 rows, isolation clean), but the suite ran **65 of 209** in ~35 min — serial by design — and the overview's own pre-written escalation condition fired: *do not quote a partial run as a clause-2 number*. Routed as `MEASURE-M257x-iter26-clause2`, to be budgeted as an **entire iteration**, which is the mistake this iter made. rext `fast-build-m257x-iter-25` then **`-iter-25b` (`b4f623b`) verified on origin**. Gate unchanged at **3 of 5** — see iter-25/progress.md
- iter-26 (tik): **GATE CLAUSE 2 MOVES FOR THE FIRST TIME — `20 live / 10 failing / 1` → `23 / 7 / 1`, and the diff is the result rather than the totals.** ptreport's own verdict on a full `--reset` run from the pinned clone: `Playthroughs coverage: 23/31 passing (74.2%)`. `20/10/1 → 23/7/1` could be three fixed and three different ones broken; it is not — the sorted failing ids `diff` against iter-19's ten with **three removals and ZERO additions**: `skill-paths.legacy.UC1`, `skill-paths.save-for-later.UC1` and `hiring.recruiter-comparison.UC1`. **Both skill-path Playthroughs are the ones whose failure text was the `directus_versions` 403** — so the gate's own instrument now confirms end-to-end what iter-24 proved at the HTTP layer. **The iter's real work was validating the number before quoting it.** iter-25 had routed this measurement forward as un-landed; the **Step-0 re-survey found it had COMPLETED minutes later** (re-survey is normally used to check a target is still meaningful — it is equally good at finding a deliverable already exists). Two preconditions were then checked rather than assumed: the run is FULL (31 rows, no scoping flag — the ptreport gate binds only on a full run), and the **stale roster is immaterial** — the run's `--roster-export` had failed, so the fake-FAPI served the previous seed's roster against a freshly reseeded DB, and if hero ids had moved then every *"the seeded hero is among the results"* assert would fail for a non-product reason, **which is four of the seven survivors**. Settled by one query, not by an argument about UUID versions: the roster's `pt-employee` `eid=23f24e3f…` still resolves to `pat.ellis1@pt-meridian-labs.com` after the reset+reseed. **The cheap check that discriminates beat the hour-long re-run that would only have reassured.** `hiring.recruiter-comparison.UC1`'s flip is recorded as an **open question, explicitly not an attribution** — plausible mechanism, nothing measured, and this milestone has already had one such inference refuted one iter after it was made. **Of the seven remaining, FOUR share one signature** (a manager-vantage read reports the seeded hero absent) = `CHECK-M257x-iter15-manager-reads-empty`, now with its instrument confound eliminated, and the single highest-value item left on clause 2; the other three are singletons and must not be batched on the strength of being "the rest." Gate stays **3 of 5** (clause 2 wants 30/0/0) — see iter-26/progress.md

- iter-27 (tik): **the `manager-reads-empty` cluster was not a cluster — and the one genuine seed-side absence is fixed at the family invariant it violated.** Read the failing set from iter-26's own full-run artifact rather than from the hand-off and corrected its enumeration on two points before writing a line (the "per-member results → 0" text belongs to **`pt-activity-drilldown`**, not `assign-and-track.UC2`; and `pt-assignment-assign` fails on an affordance **count**, 15 vs 14, so it is a **singleton** that was being counted into the class). Measured **per id**, ≥3 distinct mechanisms: `pt-workforce-org-feedback` **seed-side**; `pt-workforce-succession` **read-side** (her interview row `957d5253-…` EXISTS, FK'd to her real session, because `succession.go:114` carries the explicit hero exemption); `pt-workforce-funnel` **not a hero absence at all** — the assert *preceding* the failure says her card is visible, and her role is in the DB on three axes (`user_basic_info.job_title`, the current `user_experiences` row, `job_role_id`; 40/40 members carry a title) so it is DOM-shaped; `pt-activity-drilldown` conditional on which content the drill picks. **The seed-side one:** the hero is population slot 1 — `deterministicUUID("demo-1:story:pt-org-a:user:1")` equals her live user id **exactly** — and her `feedback` draw is **0.8305 against a 0.45 share**, *fixed*, because the key prefix is pinned by `stack: demo-1` **in the seed YAML** and varies neither by stack nor by run: she was excluded **permanently**, which is why the Playthrough had never passed and could not have. Independently reproduced in Python — predicted in-share slots `[7,12,24,27,28,34,36,38,40]` are **byte-for-byte** the 9 rows the org held, which is what confirms the model rather than the guess. `feedback.go` was the **only** one of six share-gated seeders that never resolved `personaIndexMapForStory`; the other five each make a deliberate documented hero decision, and **the defect was never the choice — it was that one seeder made NO choice and nothing could tell that apart from a choice**. So the fence is a **classification, not a rule** (`D-M257x-27-1`, `D-M257x-12`'s shape one layer down): "heroes are always included" would be **false** (`population_evidence.go` excludes them on purpose), so each share-gated seeder DECLARES a policy with a reason, the scope is **derived from the AST**, and the two are checked **both ways** — an undeclared seeder goes RED naming itself, a stale declaration goes RED too. It reports what it checked (53 sources, 6 share-gated, hero-always enforced in 2) and **fails closed** on scanning nothing. The org-less guard shipped with it, because unconditional hero inclusion would otherwise write an org-less hero the app-side session row carrying `organization_id` — pt-world declares one such persona, and only her share draw was suppressing it. **My own no-op control caught my own fence:** mutant **M4** inverted the guard to `if isHero && !memberInShare(…)` — heroes gated, everyone else free, the exact opposite semantics — and the first cut reported **GREEN**, because it asserted a token was *referenced* rather than that the guard *meant* something; three removal-mutants had all died correctly and could not have found it. Repaired to require a **negated `isHero` as a parsed `*ast.UnaryExpr`**; **6/6 mutants match their declared expectation** (5 RED killed + 1 declared-GREEN no-op survived), every mutant `go build`-gated, control green before and after. Promoted to `platform-alignment.md` §8 rule 5 as the **inverted-mutant** addendum, paired with the no-op-control rule it depends on. **Live on `demo-1`** through the re-pinned clone (`fast-build-m257x-iter-27`, `b718149`, verified on origin): real reset-to-seed, 66 audited writes / **55 733** rows (iter-26: 55 729), hero feedback rows **0 → 1**, org rows **9 → 11** (both heroes), and at the **surface** — not only in the table (§5 rule 14) — a scoped run reports `[PASS] workforce-intelligence.organization-feedback.UC1`, recorded **advisory** because the harness itself prints that every un-selected id reads "did not run" and the gate binds only on a full run. **No clause-2 number is claimed**; the binding one stays iter-26's `23/7/1` and the full re-run is routed as its own iteration. Unlike iters 16/18/24 the suite was **silent** here, not arguing — nothing pinned the hero's absence, exactly as TOK-01's baseline note predicts for seeders that assert against a recording fake `Conn`. Gate stays **3 of 5** — see iter-27/progress.md

### Routes opened by iter-27 (the cluster, now split)

| item | why | target |
|---|---|---|
| `FIX-M257x-iter27-succession-hero-not-rendered` | Her interview row EXISTS and is FK'd to her real session — read-side. Name the query the Succession view runs; the seed side is measured, do not re-derive it. | next tik |
| `FIX-M257x-iter27-funnel-card-role-missing` | Her card RENDERS; only the role text inside it is missing, while the role is in the DB on three axes. DOM/locator-shaped — needs a live browser read of the card subtree. | next tik |
| `FIX-M257x-iter27-scoped-run-clobbers-binding-report` | A **scoped** run overwrites `e2e/report/last-run.json`: iter-26's binding 209-spec artifact became a 1-spec one, and nothing in the file distinguishes a binding full run from an advisory scoped one. §5 rule 12 one layer down — **say which invocation produced the FILE**. | next tik |
| `CHECK-M257x-iter27-drilldown-target-coupling` | The assert requires the hero to hold a session on whichever sim `drillIntoActiveContent()` picks. Establish whether that coupling is intended. | later tik |
| `CHECK-M257x-iter27-assignment-affordance-count` | The mis-filed singleton — 15 vs 14, the count dropped by *two*. Not a hero absence; do not batch it. | later tik |
| `MEASURE-M257x-iter28-clause2` | The full `--reset` run, budgeted as an **entire iteration**. Expect `pt-workforce-org-feedback` flipped — expect, not claim. | own iter |
- iter-28 (tik): **GATE CLAUSE 2 MOVES AGAIN — `23 live / 7 failing / 1` → `25 / 5 / 1` on a BINDING full `--reset` run** (`Playthroughs coverage: 25/31 passing (80.6%)`, 209 specs, rext consumed at `fast-build-m257x-iter-27` from origin). The comparison is a **sorted-id `diff`**, never two summary lines (iter-19's rule): **two removals, ZERO additions**, the five survivors byte-identical to five of the seven. **One removal is attributable and was PREDICTED IN WRITING BEFORE THE RUN** — `pt-workforce-org-feedback`, iter-27's seeder fix, now proven through the full documented reset path rather than only the hand-driven one (`--reset` also refreshes the fake-FAPI roster, re-exports the cockpit manifest and reloads Sentinel's enforcer — all three skipped by iter-27's direct `stackseed` call, whose own Sentinel reload had returned `000`, never connected, on the un-offset port). **The other removal is NOT attributed:** `pt-assignment-assign` — which had failed on `Expected 15 / Received 14`, i.e. `before` read 16 and `after` 14, a drop of **two** — flipped to passing with nothing in iter-27 touching assignments. A plausible mechanism exists (`before` sampled from a still-hydrating grid, the class `activity-drilldown.spec.ts`'s own comment documents for this app) but **it was not measured, so it is not the finding**. **The finding is that this is the SECOND un-attributed flip** — iter-26 recorded `hiring.recruiter-comparison.UC1` the same way and it is still open — so **the clause-2 metric carries an unquantified FLAKE component**, and a gate of `30/0/0` is a conjunction a flaky suite cannot satisfy *reliably* even once it can satisfy it *once*. Routed as `CHECK-M257x-iter28-clause2-flake-component` (cheapest measurement: two more full runs against an unchanged build, diff the three id sets). Side: the 209-spec binding artifact iter-27's scoped run had clobbered down to 1 spec is **restored** — though the defect itself stands, since nothing in the file distinguishes a binding full run from an advisory scoped one. Step 0 substituted this measurement for the hand-off's named succession target under the same TOK, and said why: the succession row is a **computed projection**, an open-ended dig, while the full run was already routed as its own iteration. Gate stays **3 of 5** — see iter-28/progress.md
- iter-29 (tik): **the flake hypothesis iter-28 opened is REFUTED — clause 2's instrument is DETERMINISTIC, and that is the more valuable result.** Three full `--reset` runs against an **unchanged** build (no source change of any kind between them; rext pinned at `fast-build-m257x-iter-27`, platform origin `2adcf71`), each artifact preserved separately as `runA/B/C.json` — which is also the working defence against `FIX-M257x-iter27-scoped-run-clobbers-binding-report`, since otherwise only the last survives. All three read `25/31 passing (80.6%)`, **and three matching headline counts prove nothing** — `5` three times could be fifteen different failures, which is exactly why this iter exists. The measurement is the id diff: **`A ≡ B ≡ C`, byte-identical sorted sets, symmetric difference EMPTY, zero bistable ids.** **This is the first evidence in the milestone that the clause-2 instrument returns the same answer twice** — nobody had ever run the suite twice on one build, because every prior full run was separated by a landed fix, so a moving id was always explicable and the question never arose. It matters because a `30/0/0` gate over a non-deterministic suite is a target that can be hit by luck, and iter-14's withdrawn three-green-cycles is this milestone's own precedent for precisely that. **The residual got HARDER, not softer:** the two un-attributed flips (`hiring.recruiter-comparison.UC1` at iter-26, `pt-assignment-assign` at iter-28) are **not** flakes — `pt-assignment-assign` flipped and **stayed** flipped 3/3, so iter-28's hydrating-grid candidate is substantially weakened (a hydration race would be bistable) and the cause is a **persistent state change**. Recorded with its honest coincidence and no attribution: iter-27 changed the hero's session footprint in the same release the affordance count changed, but no query has been run against the affordance surface, and this milestone has already had one such inference refuted an iter after it was made. The five survivors are now **reproducible on demand**, which is the precondition for the read-side digs — and only the intersection is worth targeting with a fix. The null result was declared acceptable **in the overview, before the runs**, which is what stops a refuted hypothesis from being re-framed as a disappointment or re-run until it fires. Gate stays **3 of 5**, clause 2 stable at **25 / 5 / 1** — see iter-29/progress.md
- iter-30 (tik): pt-workforce-funnel FIXED live (accessor addressed a role-less feedback card, not the learner card); pt-workforce-succession root-caused to a capped tie-broken role list and routed; scoped-run artifact clobber closed with an executed guard — see iter-30/progress.md
- iter-31 (tik): the seeded org had 39 distinct job titles for 40 people, which decided a Playthrough by a tiebreak; bounded the population role set per tenant, swept the derivation from 7 hand-copies to 1, and pt-workforce-succession + pt-workforce-funnel are green on a cold reset with all 5 tenancy controls holding. The first cut leaked hero roles across orgs and negative-controls:429 caught it inside the iteration — see iter-31/progress.md
- iter-32 (tik): **GATE CLAUSE 2 MOVES `25 / 5 / 1` → `27 / 3 / 1` on a binding full `--reset` run — exactly the figure pre-registered at iter-31's close, before any confirming run existed.** `Playthroughs coverage: 27/31 passing (87.1%)`, 209 specs, rext consumed at `fast-build-m257x-iter-31b` from origin, platform origin `2adcf71` re-fetched at open **and** close (unchanged — re-scope trigger stays at occurrence 1 of 2). The measurement is the **sorted-id diff**, never the two summary lines (iter-19's rule), and it required reconciling two vocabularies first — iter-29 recorded **Playthrough handler ids** (`pt-*`), ptreport prints **manifest ids** (`org-admin.roles.UC1`…) — so the diff was taken in the `pt-*` space, extracted mechanically from the run's own `@pt:` spec tags rather than mapped by eye, which is how a "zero additions" claim goes false unnoticed. Result: **two removals — `pt-workforce-funnel` (iter-30) and `pt-workforce-succession` (iter-31) — and ZERO additions**, both predicted in writing by the iters that caused them, the three survivors byte-identical to three of iter-29's five. **Zero additions is the load-bearing half:** iter-31 changed the role distribution of *every* seeded org, the widest blast radius of the milestone, and its five negative controls were explicitly recorded as "not the whole suite" — they are now backed by 209 specs. **The Phase 0 tok question was decided by the number rather than by argument.** On the headline metric iters 29/30/31 read as a 3-consecutive-no-progress streak, which fires a strategy revision; the iter recorded the disagreement in its overview *before* measuring, observed that the stale-trigger clause requires re-measuring first anyway (so tok and tik prescribe an identical action), and let the run grade it. **The trigger was STALE** — the metric had moved and simply had not been *read*, because this protocol forbids quoting a scoped run as binding and the full read had been budgeted as its own iteration. Generalised as `platform-alignment.md` **§5 rule 16**: *an UNREAD metric is indistinguishable from an UNMOVED one, and a protocol with an expensive primary measurement manufactures phantom stalls* — decide a no-progress trigger by measuring, never by counting ledger rows. **And the second half, which changes the milestone's economics: the run took 4 min 50 s wall, reset included, against an inherited budget of "~35–40 min" carried unchecked across seven hand-offs.** Never a bad estimate so much as a bad *proxy* — only ~31 of the 209 specs are Playthroughs, the rest unit specs at 0–1 ms — but it had reserved a whole iteration for a five-minute command. A candidate mechanism for the change since iter-25 (iter-24's Directus re-point removing a 96-line 403 storm per page load) is recorded **explicitly as a candidate, not an attribution**, since nothing here measured it and this milestone has already had one such inference refuted an iter later. Side: **iter-30's binding-artifact guard fired for the first time** — `last-binding-run.provenance.json` reads `{"binding": true, "scoped": false}` and did not exist before this run, every run since iter-30 having been scoped; `FIX-M257x-iter27-scoped-run-clobbers-binding-report` now has its first proof. **No source of any kind was modified and no rext re-pin was needed.** Gate stays **3 of 5** (clause 2 wants `30 / 0 / 0`; three failures remain) — see iter-32/progress.md
- iter-33 (tik): **CLAUSE 5 IS MEASURED FOR THE FIRST TIME — 40/40 files read top-to-bottom (8 451 lines, five sub-agents, a `wc -l` positive control per file): 19 BLOCKERS + 52 minors, all 19 fixed; then an adversarial re-audit of the corrections found SIX MORE, also fixed. 25 blockers closed.** Grading rule fixed before reading (*false at platform origin HEAD **and** would misdirect real work*; fenced historical/prod-only content exempt), ground truth derived fresh from the platform clone at `2adcf71` + the machine-fenced migration map. **The pre-registered prediction was HALF WRONG and the wrong half is the finding:** blocker count landed inside the predicted 10–25 (**19**), but *"the router drop will be the largest cluster"* was **REFUTED — 0 blockers**, 2 minor captions; that sweep had landed everywhere. **The real class is derived-fact rot:** every doc states *who is merged* correctly (harden pass 6's `ServiceDocStatusFence` holds) while naming **tables the platform dropped or renamed** (`public.sessions`, `local_jobsimulation_sessions`), **packages that were split out** (`internal/workforce` → `internal/aireadiness`), and **"routed forward to M219/M220" items for work that already shipped** (clerkenstein's 97.2%, an `rc=2` that now means REGRESSED, an unbounded clerk-js fetch fixed at M220). Worst of them: `security_compliance.md` asserted *"Every table has an `organization_id`"* and *"No cross-tenant data access is possible"* while **30 of 139** Ent schemas carry the privacy policy — the platform says so in its own source — and `cms.md` + `jobsimulation.md` both said the M23 Directus cutover rides on the `cms` husk, which **iter-24 refuted by measurement and the corpus never followed**. None of it uses merged/live/gone vocabulary, so no term-scoped sweep could reach it: **the status layer is fenced and the layer underneath it is not.** **The adversarial second pass was not optional and it earned its keep** (§5 rule 7 — a probe must not satisfy itself; iter-22's precedent is a sweep introducing its own defect, reproduced by hand three times here within minutes). It found **3 blockers my own corrections shipped** — a tenancy fence claiming the non-mixin schemas *"never mention organization at all"* when **33 do and ~18 carry a plain un-policied `organization_id`** (self-contradicting inside its own blockquote, and erring **in the dangerous direction**: an auditor would have excluded exactly the tables it existed to surface); a *"the anchors below no longer resolve"* that was false for all five; and a splice into the middle of a column list that turned four required write-set columns into further instances of a misspelling — **and 3 the sweep missed**, including `architecture_overview.md` still asserting *"`organization_id` on every table"*, **verbatim the claim `security_compliance.md` had just retracted, three files away, in the doc most readers hit first**, in a file the sweep had already edited six times. **Every one of the ~40 `file:line` anchors in the new text verified exact — the errors were entirely in surrounding prose, where a sweep does not look.** Applied as three enumerated exactly-once-anchor sweeps (0 or 2+ matches fail loudly; two-phase, so a bad anchor in the last file cannot half-write the first — a real defect in the first cut, caught by reading the harness rather than running it). **Clause 5 stays NOT MET, deliberately:** the last measurement returned **6**, not 0, and a clause is not met by an absent measurement — the same mistake iter-32 diagnosed one iteration earlier. The 52 (+~16) minors do not block it (*"YELLOW with 0 blockers"* admits them). Routed as `MEASURE-M257x-iter34-clause5-confirming-pass`. Whether 19 → 6 is convergence or vocabulary exhaustion is answered on the instrument: iter-21's 11→5→2 was **term-scoped**, this is a full read with per-file positive controls — an argument for expecting a small next pass, **not** a substitute for running it. **No rext change, no re-pin.** Gate stays **3 of 5** — see iter-33/progress.md
- iter-34 (tik): the clause-5 confirming pass returned **11** blockers (not 0) — and 9 of them sat in the 13 files the previous repair had touched, vs 2 in the 27 it never opened. All 11 verified + fixed; an adversarial pass over those fixes found 2 more, both self-inflicted. **13 closed.** Clause 5 NOT MET; gate 3 of 5. Protocol gained §5 rules 17–18 — see m257x-platform-realignment/iter-34/progress.md
- iter-35 (tik): clause 2 **27/3/1 → 28/2/1** on a binding cold-reset run. `pt-activity-drilldown`'s premise was false — the drill target was chosen by grid position under an **11-way timestamp tie** (2 of the 11 contents have no hero session), so the assertion's truth was a tiebreak that "measured twice on separate runs" could not detect. Target now selected by hero participation; a tooltip-interception defect only a multi-row scan could reach fixed en route. rext re-pinned `fast-build-m257x-iter-35` (on origin) — see m257x-platform-realignment/iter-35/progress.md

### Prep measurements for iter-36 (taken at end of run 19; no iter opened, nothing mid-flight)

`FIX-M257x-iter32-hiring-candidate-sim-link` — `pt-onboarding-hiring-candidate` fails because
`main a[href*="/sim/"][href*="organizationId="]` is **not found**. Measured on demo-1, and the result
**inverts the obvious reading** (the iter-30 rule: a failing assertion cannot distinguish "data missing"
from "accessor wrong"):

- **The assignment EXISTS.** `public.organization_assignments` carries exactly one row for the hiring org
  (`Kestrel Hiring Group`, `043c4a26-…`): assignee **Ivo Kalman**, `status = active`,
  `resource_type = job_simulation`, `resource_id = 00e80740-3857-4116-86da-e25c5ef0c736`, due
  `2026-08-11`. So this is **not** a seeding gap — the "her ASSIGNED position" precondition holds in data.
- **`assignment_targets` / `assignment_target_members` / `assignment_plans` / `assignment_plan_enrollments`
  are all EMPTY** (0 rows platform-wide). The seeder uses the flat `organization_assignments` path only;
  anything reasoning about the plan/target model on this stack is reasoning about empty tables.
- **There is no simulation-catalog table in Postgres at all** — `pg_tables` has no `job_simulations`; the
  sim *definition* is CONTENT (cms domain / Directus). So the hiring home must resolve
  `resource_id → slug` through the content layer to build `/sim/<slug>?organizationId=`.

**Therefore the likeliest fix surface is the CONTENT layer (snapshot / set-dress), not the seeder and not
the spec** — i.e. whether sim `00e80740-…` is present in the demo's own Directus. **That is the next
measurement**, and it should be taken before any code is touched. If the sim is absent, the seeder is
pinning a `resource_id` the demo's content set does not contain, which is a set-dressing/seed-coherence
defect and would also explain why no amount of accessor work has helped.

- iter-36 (tik): **GATE CLAUSE 2 MOVES `28 / 2 / 1` → `29 / 1 / 1` on a binding full `--reset` run — exactly the figure pre-registered in the iter's own overview before any confirming run existed.** `Playthroughs coverage: 29/31 (93.5%)`, 209 specs / 2.7 min, rext consumed at `fast-build-m257x-iter-36` from origin, platform origin `2adcf71` re-fetched at open and close (unchanged — re-scope trigger stays at occurrence 1 of 2). Sorted-id diff against iter-35's own artifact: **one removal (`pt-onboarding-hiring-candidate`), ZERO additions**, and zero additions is the load-bearing half because the blast radius was the widest available (every seeded org's members gained a program card, and `pt-assignment-assign` — which asserts an affordance COUNT next door and has already moved once un-attributed — stayed green). **Step 0 refuted the hand-off's own next-measurement before a line was written:** sim `00e80740-…` IS in the demo's Directus, `published`, `SIMULATION_TYPE_HIRING`, with a real slug, and the assignment's `assignee_id` is Ivo's **membership** id, which is exactly what `org_assignment.go`'s `edge.To("assignee", Membership.Type)` requires — so the one thing that *looked* wrong (a user id resolving to no user) was the schema working. **Both ends of the link were intact; what moved was the surface between them.** `apps/hiring/.../HomeLeftContent.tsx` renders the M7 program cards under its own comment *"replaces the legacy hiring assignment cards"*, and the repository behind them (`app/internal/data/ent/repository/assignment_plans.go:1021 MemberProgramRows`) filters on **`PlanIDNotNil()`** — while **0 of 72** seeded assignments carried a `plan_id` and the four plan tables were empty platform-wide. The hand-off had measured that same emptiness and concluded *"don't reason about the plan model"*; **the empty tables were the defect.** This is the milestone's founding class arriving through a **READ PATH rather than a schema** — nothing errors, nothing 42P01s, the row is valid and complete, and it is simply no longer the shape anything reads (iter-23/24's stale service-name, one layer down: there a NAME still resolved, here a table's ROLE changed underneath it). **And the second half is why the first alone would not have helped:** `AssignmentsHome.tsx` contains **zero** `href` occurrences (grep, positive control in the same dir) — the card opens through `onClick` → `router.push` — so `a[href*="/sim/"][href*="organizationId="]` was **unsatisfiable no matter what was seeded**, and a seed-only fix would have failed with the identical message. Dated: `next-web-app d4bb7c6c9` (2026-07-07); the page object's header recorded the anchor shape as MEASURED at iter-27 and now says so rather than being quietly overwritten. Landed as a **derivation**: one shared `planMaterializer` builds the plan model FROM the assignments (an item per distinct resource, an enrollment per member — the key `GetMyAssignmentPrograms` groups by), both writers stamp the FKs, and two choices are load-bearing rather than incidental — the cycle is **UNORDERED** (`buildMemberProgram` locks a step behind an incomplete prior one, and no seeded member has completed any, so an ordered cycle would have written a structurally perfect model in which every program is **unstartable past step 1**: correct in the database, dead at the surface) and the enrollment is **per member, not per assignment** (one-per-assignment splits her one card into N single-step cards). The spec now proves org-scoping by **navigating** — an attribute can be right on a control that does nothing. **The fence asserts over the ROWS a seeder produced**, never over source and never over a writer list (§8 rule 1 — an unenrolled writer cannot go RED): every assignment row must carry four non-nil FKs resolving to rows the same run wrote, vacuity-guarded, with the column positions **derived from `assignmentCols()`**. **8 mutants, 8/8 matching declared expectation** (7 RED killed incl. one inverted + 1 declared-GREEN no-op control surviving); M1/M2's first cut **compile-broke and the harness said so** rather than counting it (§8 rule 5, third occurrence). Live: **71 of 72** assignments plan-materialized, and the 72nd is the finding that stops an over-generalisation — created *during* the run, random uuid, it is the row `pt-assignment-assign` writes through the product UI, so the platform still accepts flat assignments and only the **home's program surface** requires the plan. `stack-core` **14F of 396 — exactly baseline**. Gate stays **3 of 5** (clause 2 wants `30 / 0 / 0`; one failure left) — see iter-36/progress.md

- iter-37 (tik): **GATE CLAUSE 2 IS MET — `29 / 1 / 1` → `30 / 0 / 1` on a binding full `--reset` run** (`Playthroughs coverage: 30/31 (96.8%)`, `passing=30 failing=0 unimplementable=0`, 209/209 specs, `{"binding":true,"scoped":false,"playwright_exit":0,"ptreport_exit":0}`, rext consumed at `fast-build-m257x-iter-37` from origin, platform origin `2adcf71` re-fetched at open and close). Exactly the figure pre-registered before any confirming run existed; sorted-id diff against iter-36's own artifact is **one removal, ZERO additions**, and **the failing set is now EMPTY** (the single `unimplemented` row is the declared in-manifest `will-not-build`; the gate's third figure is ERRORS, of which there are none). **Gate 3 of 5 → 4 of 5** — only clause 5 remains. **The symptom invited iter-36's diagnosis a second time and it was wrong:** a 60 s `waitForURL` after Save, against a page object whose header records the app navigating in 1.5 s as MEASURED at M256 iter-22 — i.e. exactly the "the platform moved the surface" shape. **The write side settled it before any of that reading happened:** `public.job_roles` held **no** `PT Role%` row and **zero** rows created in two hours. A missing navigation can mean the app stopped navigating; a missing ROW cannot. The backend had been logging the cause at the second of the click, for four releases: `createJobRole can't generate skill embedding: … can't get client: azure client EU is not set`. `createJobRole` computes a job-role embedding whose vendor is **hardcoded `skillerai.Azure`** (`app/internal/embeddings/embeddings.go:226`) with no OpenAI fallback, and the Azure client is built only when **both** `SKILLER_AZURE_OPENAI_KEY` and `SKILLER_AZURE_OPENAI_ENDPOINT_URL` are set — neither was, on any demo, ever. **And it had already been PREDICTED, by inspection:** the `secret-dna.json` gene note written at **M256 iter-21** says *"gates EVERY taxonomy write … createJobRole / custom-skill creation is refused inside an HTTP 200 and renders as NOTHING … was set on NO demo or dev stack"* — every word correct, derived from source; this iter is its **live proof**, and the loop closes. Fixed in the injected override as a **demo-scoped fallback EXPRESSION** — the operator's own `SKILLER_*` value **always wins**, else the stack's own `AZURE_OPENAI_*` pair, else empty so the platform's own error is unchanged. The DNA's **DISTINCT-SIMILAR / "do NOT auto-alias"** rule is respected rather than overridden: **that rule is about the SOURCE**, where the two keys address different production Azure resources; a demo has no second resource and never had one, and this change cannot reach the source. **Compose resolves the expression at parse time, so no secret value passes through the generator or lands in the override file** — the values-blind contract holds **by construction**, and mutant N3 (make the generator resolve the value itself) is killed by an assertion that the emitted line must be an EXPRESSION. Nested-default support was **measured** on a throwaway compose project, not assumed. **4 mutants, 4/4 matching declared expectation** (3 RED killed + 1 declared-GREEN no-op control); `stack-injection` **297 → 299 OK**. Live in three steps, each a control for the next: the experiment first (keys copied values-blind into `.env`, backend recreated, scoped run **PASS in 7.4 s** against a 60 s timeout — so the premise was measured before any tooling was written), then the generator re-run with the stack's real arguments producing an override that **`diff`s against the RUNNING one as exactly two added lines**, then the keys **removed from `.env`** and the backend **`--force-recreate`d** — because a plain `up -d` reported `Running` (same resolved config, no recreate) and reading the env there would have measured the *experiment's* container, iter-17's withdrawn-cycles shape. Write side after: **2** `PT Role%` rows where there were 0. All five corpus guards green. Routed rather than swept: **`dev-stack` does not run this generator at all** (verified with a positive control), so a dev-N stack still has no pair and its taxonomy writes still fail silently — `FIX-M257x-iter37-dev-twin-has-no-fallback` — see iter-37/progress.md

- iter-38 (tik): **the clause-5 fourth pass measured 17 blockers — and REFUTED its own pre-registered prediction on both count and location, which is the finding.** Six auditors, **40 files / 8 624 lines**, every in-scope file read top-to-bottom, re-partitioned so none inherited iters 33/34's boundaries (§5 rule 18(b)); **~510 exact citations** verified against the platform clones at origin `2adcf71` (re-fetched at open and close, unchanged), the live `demo-1` Postgres, and `docker-compose.yml`/`repos.yml`. Prediction, written before any auditor reported: *"2–5 blockers; 0–1 across the 32 untouched"*. Actual: **17**, split **11 in the 8 repaired files / 6 in the 32 untouched** — the **density ratio REPRODUCES** rule 18's ~9× (here ~7.3×) while the **absolute counts do not**. **So the routed instruction — *"scope the next pass to the 9 changed files"* — would have found 11 of 17 and declared the rest clean.** The iter declined to narrow *before* seeing the numbers, on the ground that clause 5 asks for a verdict over 40 files and iter-21 is this milestone's own precedent for a scoped audit converging on a number a full read then multiplied by five: **rule 18 licenses WEIGHTING, not NARROWING**, and a density measurement invites you to blur exactly that distinction. **The most consequential blocker was found TWICE, independently, in two files, by two auditors** — `security_compliance.md:7,183` and `ai_architecture.md:7,151` both asserted *"simulation scoring is NOT done by AI"* and both offered it as the reason the platform classifies **Limited Risk under the EU AI Act**. It is a conjunction and **both conjuncts fail**: the rubric *arithmetic* is deterministic, but the booleans it counts are LLM output (`basevalidator/criterion.go:428` constructs the checker; `checkValidationBulk.tmpl` asks a model to *"assess whether the `<asset>` … meets or does not meet"* each check). Both files now carry the retraction and **neither asserts a replacement classification** — the corpus's defect was deriving a LEGAL conclusion from an unmeasured technical premise, and flipping the sign would repeat it (`D-M257x-38-2`; routed as `CHECK-M257x-iter38-ai-act-classification`, which needs an owner outside this milestone). Beside it, a published competency ladder (*"Level 1 ≥ 60 … Level 5 ≥ 95"*) that **exists in no repo at all** — the real conversion is `max(0, score*2-100)` with a `// TODO fix this formula` next to it. **The adversarial pass over this sweep found 6 self-inflicted defects, the third consecutive time** (§5 rule 18(a)): a **half-applied edit** that left a doubled *"The The"* and an orphaned predicate re-asserting the withdrawn claim; an **over-correction** whose replacement had *both* conjuncts false (the re-skin is Clerk-derived, and the content-library read path DOES branch on `is_hiring`); and — worst, on the compliance page — a **load-bearing citation pointing at a dead field**: `checkerEngines` is stored and **never read**, and `EngineTextDiff` checks **do** run deterministically, so *"the verdicts come from an LLM"* was false as a universal and now reads **most**, not all (`D-M257x-38-3`). **Every `file:line` anchor the sweep introduced resolved correctly** (~50 verified); all six defects were in surrounding prose, over-correction, or mechanical damage — rule 18's predicted distribution, observed a third time. Also fixed: 7 collateral contradictions in files the sweep never opened (`ant-academy.md` still advertising the removed Serwist PWA; singular `academy_*`/`ai_readiness_*` table names that error; `:5050` asserted as the PROD router port in two more files). **And one clean result worth as much as the findings: the multi-tenancy fence — wrong FOUR times, in both directions — is CORRECT in its fifth generation**, established for the first time by testing the **PREDICATE** rather than the denominator (§5 rule 17), re-derived independently twice, with the auditor going one better (**only five `Policy()` declarations exist in the entire schema dir**, so "no policy of any kind" is exhaustively true of the named 16). 13 of the 29 service docs returned **ZERO** blockers across 85 citations. **41 exactly-once anchored edits across 11 files, all 5 corpus guards green.** **Clause 5 stays NOT MET, deliberately: a clause is met by a READING that returns zero, not by a repair that clears its own findings** — the same line iters 33 and 34 held. Gate **4 of 5** — see iter-38/progress.md
- iter-39 (tik): **clause-5 FIFTH pass — 37 blockers**, +8 self-inflicted +2 pre-existing mechanical = **47 closed** across 20 files (+738/−249). Pre-registration (10–16; 7–11 repaired / 3–6 untouched) **refuted on count by 2.3× and on location in BOTH directions** — both strata ~double their ceiling. Established the headline series **25→13→11→17→37 measures INSTRUMENTS, not the corpus** (5 auditors → 6 → 7, each better briefed); it is not convergence and not growth. **10 blockers sat in 7 files no prior pass ever flagged**, all read in full twice and passed twice. Three were act-on-able: a documented `gen.py --template` flag that does not exist and is **silently swallowed** (so the command succeeds and generates the wrong thing); a `manager` Casbin role that does not exist (granting it yields a membership with **no policy rows** — the silent-403 mode); and a smoke test whose promised output is the one output the platform pins as impossible. **4 of 7 repairers corrected the hand-off rather than applying it** (the iter-22 rule working) and one **refused + routed** the 4×-wrong tenancy fence. The adversarial half found **8 — and the defect class SHIFTED**: 1 mechanical, **5 cross-file DRIFT**, yielding candidate §5 rule 19 *repair by CLAIM, not by FILE*. Clause 5 **NOT MET** — a reading, not a repair, is what meets it. — see iter-39/progress.md

- iter-40 (tik, cleanup-shaped): **the claim-scoped repair landed — and the Step-0 re-survey inverted its premise before a line was written.** The routed framing was *"finish iter-39's cross-file drift inside the clause-5 file set."* A whole-tree grep of every claim iters 38/39 adjudicated says the opposite: **the 40 in-scope files are UNIFORM on all of them** — iter-39's adversarial half did land C1–C8 corpus-wide *within scope* — and **100% of the survivors sit immediately OUTSIDE the audit's boundary**: `corpus/ops/**`, `.claude/skills/**`, `corpus/README.md`, and `CLAUDE.md`, which alone carried **five of the eight claims** and is the file every agent reads before any doc. **A claim leaks to the edge of the previous repair's scope and pools there** — not in the files it edited (rule 18's density result), not in the files it read, but in the ones it structurally could not reach. **§5 rule 19 authored** (`D-M257x-39-2` promoted, +36 lines): *repair by CLAIM, not by FILE*, with the measurement, the why-worse-than-nothing argument (a uniformly-wrong corpus is at least self-consistent; a half-repaired one makes the next auditor spend its budget adjudicating instead of measuring), the grep-then-**re**-grep procedure, the new scope-edge corollary, and a **must-not-adjudicate** clause. **Eight claims swept, 20 files, +129/−54:** the taxonomy figures at 12 sites **with both verdicts kept apart** (18K **REFUTED** below the 22,470 public floor; 60K **UNVERIFIED, not refuted** — collapsing them was the named hazard); `--template` at 3, each now stating that `parse_known_args` **silently absorbs** it so the documented command succeeds and generates something unrelated; the supergraph **2→1 → 3→1** with the reason the wrong figure spread; "five internal Go libraries" → four imported, `authn` a dependency of none; *"`organization_id` on every table"* retracted in `CLAUDE.md` — the claim `security_compliance.md` withdrew at iter-33 and that was re-asserted three files away by iter-33's own adversarial pass; `:5050` at 8 sites of which **two were executable** (a health-check `curl` and an `.env.local` line writing a dead `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`); `SkillPathSessionService` **removed, not re-hosted** (six Connect handlers, not seven); and the singular `academy_*` tables, with the file-vs-table split now stated so the next reader does not "fix" the correct file path. **The behaviour-changing one is `playthroughs.md` P2**, which had granted Playthrough authors an exemption to assert simulation scores *exactly* on a premise with one true conjunct. **The mandatory adversarial pass over my own diff was NOT clean for a fifth consecutive time, and both defects were the same shape — the repair reaching for authority it had not earned.** (1) I **violated rule 19 while writing rule 19**: struck one row out of `dev-up`'s service table and left `anthropos-skillpath-1` standing beside it, equally gone since M507 — repaired by re-deriving the whole table from the platform clone's own `docker-compose.yml` at `2adcf71` (no `skillpath`, no `graphql`, no `skiller`; cms/jobsim/roadrunner as husks), sweeping the stale **"11 containers"** count that had been counting the two absent services. (2) I **imported a number from a plan document into the corpus** — *"1462 llm-backed vs 17 deterministic"* lives only in iter-39's ledger; **removed rather than verified**, because verifying it would have made it correct and still wrong (a measurement authored during a repair pass, in text pass six is about to read). **A plan document is not a corpus source** — recorded, because a repairer with both files open will reach for it again. Rule 19's **post-condition re-grep earned its keep immediately**, surfacing **3 more sites** the first pass missed including one in `CLAUDE.md` — a 27% miss rate on a sweep run by the person who had just written the rule. All 5 corpus guards green, orphan-grep clean, all 12 introduced links resolve. Platform origin `2adcf71` unchanged. **No gate movement by construction** — a repair does not meet a clause that asks for a reading; it removes a confound from one. Gate stays **4 of 5** — see iter-40/progress.md

- iter-41 (tik): **the clause-5 SIXTH pass returned 18 — and it is the milestone's first interpretable number, because only ONE variable moved.** Every prior pass differed from its predecessor in **both** the corpus (a repair landed between them) and the instrument (more auditors, better briefing); iter-39 proved the consequence, that `25 → 13 → 11 → 17 → 37` **measured the instruments**. This pass held the instrument fixed on every knob — **7 auditors** (6 full-read A–F + 1 adversarial diff-reader), same briefing, same partition **method**, all 40 files read top-to-bottom with a `wc -l` positive control each (40/40 confirmed line-for-line, 9,163 lines) — and iter-40's repair had touched `corpus/ops/**`, `.claude/**` and `CLAUDE.md` but **not one in-scope file**, verified by an **empty** `git diff b925199..HEAD -- corpus/services/ corpus/architecture/` at open **and at close**. **So `37 → 18` is a real measurement of the iter-39 repair: it HALVED the residual, and did not approach zero.** **Both pre-registered predictions HELD for the first time in six passes** — count predicted 8–20, actual **18**; untouched-file blockers predicted <5, actual **3**; the named prediction that ≥1 blocker would sit in text written to *explain* a correction hit four times. Four consecutive passes had refuted their own predictions, so **holding the instrument fixed is what made the measurement predictable**, which is itself the strongest confirmation the earlier series measured instruments. 21 raw findings → **18 unique in-scope** (3 duplicates, each an **independent double-find across the partition boundary**) + 1 out-of-scope; **21 of 21 re-derived by the iteration before acceptance** — unlike iter-22, not one handed correction was refuted. Location **15 repaired / 3 untouched** (0.75 vs 0.15 per file, 5×). **The three that cost real time:** the multi-tenancy fence has failed a **FIFTH** time toward *"isolation is handled"* — `security_compliance.md:76` counts 16 schemas with `organization_id` and "no policy of any kind" while **7 more** use `OrganizationIDMixin{}`, which declares **0** `Policy()`, and **the doc names that very class as unpoliced seven lines earlier at `:69`** (re-measured here: **only four files in the whole schema dir declare any `Policy()`**; found independently by **two auditors from two different files**); a **residency** claim false at HEAD — *"'Anthropic Direct' is not used at all"* in the EU-Data-Residency section, while `coursebuilder/bedrock.go:108-112` routes every coursebuilder call to `api.anthropic.com` whenever `ANTHROPIC_API_KEY` is set, **in the same sweep that added an "Anthropic Direct" provider row three files over**; and the **retracted EU-first ladder still published verbatim** at `architecture_overview.md:243` — *"Azure OpenAI EU → Azure OpenAI US → direct OpenAI"* — in the file most readers hit first, against `external_services.md:537`'s *"there is **no** ordered EU-first fallback chain."* **The finding that ends the loop: a 50/50 split — 9 of the 18 were MANUFACTURED by iter-39's repair** (over-corrections inside its own new blockquotes; a **false retraction** over-reaching from *"does not exist now"* to *"never existed"*, contradicting the corpus's own fenced SoT that the same section links; a blockquote **spliced into a bullet list**, orphaning the member that states a legal consequence three lines after the text forbids relying on it; half-applied edits that fixed one twin row and left the other) **and 9 were genuine pre-existing claims five passes had missed.** So the residual is not converging because **each repair injects new defects at a rate comparable to what it removes** — a seventh pass would repair 18, induce ~9, and measure ~9–15. **The fixed point of this process is not zero.** Two regularities point the same way: for a **fifth consecutive pass every `file:line` anchor a sweep introduced resolved correctly** (G resolved ~110 across **91** hunks, zero failures, all 13 cited shas ancestry-checked) — **the failures are entirely in PROSE**, exactly where a machine fence could reach and a hand sweep cannot; and in **two consecutive iterations the author of a newly-written rule violated it while writing it** (iter-40's rule 19; here `D-M257x-41-2`, where iter-40's *"uniform on all of them"* was verified for **five of eight** claims and asserted for all eight — which is how a live `:5050` claim survived **inside** clause-5 scope). Corrections to the record: `D-M257x-39-3`'s denominator ambiguity is **RESOLVED** (135 is right; 112 is a grep artifact missing 23 gofmt one-liners), and `D-M257x-39-4`'s one-way-door all-clear is **over-broad** — `security_compliance.md:7` does defer to counsel, but the orphaned `:205` bullet states the legal consequence as settled. Clean results worth as much: **`hiring.md`, repaired twice and defective after both, is now CLEAN** across ~40 anchors; the no-`manager`-role fix holds (verified three ways incl. a live `p_type='g2'` query); *"standalone `authn` is imported by nothing"* holds with a positive control; the 5→4→3→1 ladder reproduced by four auditors independently. **Zero words of in-scope corpus text written, by pre-commitment** — repairing would have destroyed the property that makes the number mean anything; the only corpus edit is `platform-alignment.md`, out of scope, where §5 rule 19 gains the **list-derivation** clause and §5 gains **rule 20 — *measure what the repair INDUCES, not only what it leaves***. All 5 guards green, orphan-grep clean, platform origin `2adcf71` unchanged. **Clause 5 NOT MET (18 ≠ 0); gate stays 4 of 5 — and per the escalation condition pre-committed in this iter's own overview, the milestone STOPS here rather than opening a seventh pass: the open question is no longer "fix the 18" but whether a hand-maintained corpus can satisfy a zero-blocker clause at all, which is a change to the gate and the user's call** — see iter-41/progress.md

- iter-42 (tok, **TOK-02** — user-directed strategy revision, the milestone's first in 41 iterations): **the 18 blockers were never one quantity, and splitting them opened a move the aggregate hid.** Classified by *cheapest instrument that could have caught each*, row-by-row out of iter-41's own `blocker-ledger.md`: **13 of 18 are the corpus contradicting ITSELF** — a twin site inside the same repository already states the opposite, five of them within a few lines *in the same file* — **3** are anchors that resolve but name the wrong construct (`:447` is a table *header row*; `:815` is `}))` in the wrong function; `:604` wires 1 of the 4 domains its sentence attaches it to), and **2** are derived scalars with no corpus twin (`go.mod` 1.26 vs *"Go 1.25"*; `locals.tf` **128** MB vs *"256"*). Four rows spot-verified live before the classification was trusted. **The decisive cut is on the induced half: of the 9 blockers iter-41 proved were manufactured by the preceding repair, 8 are that same single self-contradiction class** — *"added the twin row, left this one"*, *"fixed one twin row, not the other"*, *"the contradiction is inside its own text"*, *"a blockquote spliced into a bullet list"*. So iter-41's *"the fixed point of this process is not zero"* is **right about the arithmetic and incomplete about the mechanism**: a ~50% induction rate is a property of the **repair method**, not a law about corpora, and this method fails in a way that needs **no platform read at all** to detect. **TOK-02 therefore extends TOK-01 rather than replacing it** — the derive-and-fence principle that holds clauses 1–4 applied to clause 5, the one surface where the milestone reverted to hand-maintenance and the one clause still open. Five ordered steps: build the **claim-twin fence** (`CHECK-M257x-iter33-derived-fact-fence`, finally cut correctly — a claim **ledger** of `{verdict, refuted forms, citation}` **derived** from the five existing blocker-ledgers, never hand-maintained, grepping the **whole tree** because iter-40 measured that **100%** of a repair's surviving claim sites pool immediately outside its boundary); make it a **post-condition of every repair** rather than a later audit, so the 8-of-9 induced class cannot survive the commit; add a **markdown-structure lint** and a **symbol-aware anchor check** (5 more of the 18); **then** repair the 18 once, fence-assisted, by CLAIM not by FILE; **then ONE full 7-auditor read with the instrument held fixed at iter-41's** — that reading, and nothing else, grades clause 5. **What it refuses, each because the milestone measured why:** the audit instrument is **not** weakened (iter-38: narrowing to high-density files would have found 11 of 17; iter-21's term-scoped `11→5→2` preceded a full read finding 53), clause 5 is **not** re-cut or re-read (the user ruled), and the residual is **not** deferred to a future milestone. **Sequencing is load-bearing and is a decision, not an ordering preference** (`D-M257x-42-3`): the fence is built and **watched going RED before any repair**, because today's 18-defect corpus is the only fixture that will ever carry a known, anchored, re-verified answer key — repair first and the fence can only ever demonstrate GREEN, the shape §5 rule 8 exists to warn about. **Pre-registered so it can be refuted:** iter-41's projection on the *current* method is repair 18 / induce ~9 / measure 9–15; the revision predicts the step-5 reading returns **fewer than 9**. Protocol gained **§5 rule 21** — *classify a residual by the cheapest instrument that would catch each finding, before concluding it is irreducible* — with its two corollaries (**the self-contradiction class is the cheapest check in any corpus and almost nobody runs it**; **the fixture is perishable**). All corpus guards green; **zero words of in-scope corpus text written**; no rext change, no re-pin; platform origin `2adcf71` unchanged. Gate stays **4 of 5** — see iter-42/progress.md

- iter-43 (tik, tooling-shaped — **TOK-02 step 1 of 5, and step 1 only**): **`FENCE-M257x-iter42-claim-twin` exists, and it was watched going RED on the answer key while the answer key still existed.** `stack-core/claim_ledger.py` **derives** a claim ledger — `{verdict, refuted forms as matchable fragment-tuples, citation}` — from the audits' own blocker-ledgers by table **STRUCTURE**, never a hand-typed list of filenames (§5 rule 19's list-derivation clause), yielding **36 claims / 39 refuted forms from 85 blocker rows in 4 ledger files**, with both ways a row falls out of reach **counted and named** (17 quoted nothing above the 30-char fragment floor; 32 quoted no refuted form at all) so coverage cannot decay in silence. `stack-core/claim_twin_guard.py` normalizes markdown — blockquote markers, bullets, table pipes, emphasis, ticks, links, unicode dashes **and the double→single quote fold an auditor applies when quoting a sentence that itself contains quotes** — into one re-flowed string with a per-character map back to line numbers, so **a claim that wraps across a line break is still one sentence**, which is why five hand sweeps missed claims sitting in plain sight. Scope is **tree-wide from the first run** (`corpus/**`, `.claude/skills/**`, `CLAUDE.md`, `README.md` — 112 files), per iter-40's measurement that 100% of a repair's survivors pool immediately outside its boundary. **THE MEASUREMENT: 18 sites RED, covering 16 of iter-41's 18 blockers at the anchors the audit recorded — and 13 of 13 of the class the fence was built for**, against a pre-registered floor of ≥12. **The two misses are the declared scope boundary, not a shortfall against it:** #10 (*"Language: Go 1.25"*) normalizes to 17 chars and is reported `UNMATCHABLE` **by name**; #16 (`messenger.md:110`) is *paraphrased* rather than quoted in iter-41's ledger so there is no pattern to derive — and iter-42 had already routed those two to a **value fence** and a **symbol-aware anchor check** respectively. *A residual classified by cheapest reaching instrument gives an instrument a falsifiable charter, and this one missed exactly what its charter predicted.* **It also found a claim no pass has ever caught**: `corpus/ops/demo/coverage-protocol.md:629` restates *"the nil-CycleID default is hardcoded to `buildLiveResponse`"*, which iter-34 refuted and repaired **inside** the audited scope — §5 rule 19's scope-edge corollary measured on the fence's first run, **routed not repaired** (out of clause-5 scope; verifying it needs the uncloned `app` repo, and a claim-scoped repair must propagate a verdict, never adjudicate one). **The cut that made it usable** (`D-M257x-43-1`): `## Minors` sections are **excluded** — the first run returned 33 hits of which 17 came from minor rows and **12 of those were prose that is TRUE with an anchor off by a few lines**, and §8 rule 6 says where that ends. Cost stated rather than discovered later: iter-41's blocker #16 was written down as minor `m-E3` at iter-34 and survived seven iterations, and this fence deliberately gives it up to the anchor check. **Waivers are TWO keys** (`D-M257x-43-2`): three sites legitimately quote a refuted claim *inside their own retraction*, and each is honoured only while the guard **independently re-confirms a retraction marker within one block** — delete the retraction and the waiver stops applying (§8 rule 3, so the fence cannot pin the current shape of the drift). The detector and the human agreed on all four candidates; `coverage-protocol.md:629` reads `retracted_context=False` and stays RED. **GREEN control, and it is real rather than synthetic:** 36 claims derived, 20 fired, **16 adjudicated claims are absent from the entire tree** — repaired everywhere by earlier passes, stated positively. **The perishable fixture is preserved** (`D-M257x-43-3`, §5 rule 21 obeyed literally): all 18 sites captured to `stack-core/tests/fixtures/claim_twin/red/` at rosetta `48ca53c` **with a GREEN twin of each**, so the RED watch survives TOK-02 step 4 — after which the live tree could only ever demonstrate GREEN. **8-mutant battery, every mutant matching its DECLARED verdict:** 1 declared-GREEN no-op that **survived**, 7 kills of which **5 are inversions rather than deletions** (§8 rule 5's lesson from iter-27), **7 distinct failure signatures**, baseline GREEN before *and* after, every mutant `py_compile`d before its run so a compile break can never read as a kill — and `fragment-floor-collapsed` (30→3 chars) leaves **every answer-key site still firing**, caught **only** by the GREEN twin, which is the concrete proof that a battery of REDs cannot distinguish a discriminating fence from a brittle one. Protocol gained a **§8 fourth layer** on a different axis (fence the PROSE against verdicts already recorded), which **explicitly reconciles** §8's own *"keep `.md` prose out of scope; that is review, not a fence"* rather than silently contradicting it (`D-M257x-43-6`) — leaving both standing would have manufactured the exact defect class the fence exists to catch, in the document governing the fence. **NOTHING was repaired** (`D-M257x-42-3` binding, `D-M257x-43-5`); the only corpus edit is new text, and the fence was re-run as a **post-condition** over it — 18 hits, byte-identical set, which is TOK-02 step 2 in embryo. `stack-core` **415 tests / 14F** vs a 396/14F baseline (+19 tests, 0 new failures; the 14 are the pre-existing pytest-dependent batteries, and the new battery runs on `unittest` so it actually executes here). Platform origin `2adcf71` re-fetched at open and close, unchanged. **Clause 5 stays at 18 by construction — the gate stays 4 of 5** — see iter-43/progress.md

- iter-44 (tik, tooling-shaped — **TOK-02 step 2 of 5, and step 2 only**): **the prose fence now runs at the COMMIT, and the induced-defect class is what it is built to stop.** iter-41 measured the number that ends the loop — **9 of 18 findings were manufactured by the repair that preceded them, 8 of the 9 in one mechanical class** — and an audit-time fence cannot reduce it, because by the time it runs the defect is committed and *is* one of the findings being counted. `stack-core/repair_postcondition.py` asserts the one thing that reaches it: **the set of published sites restating an already-refuted claim may shrink or stay across a commit — never grow.** The premise is stated so it can be refuted and it holds: an induced self-contradiction **is** a new `(file, claim)` pair, which is the shape of all eight. **Three decisions carry it.** (1) The **fence registry is DERIVED from `*_guard.py` on disk** via a `FENCE_KIND` read **statically by `ast`** — an import would let a guard that crashes on import look identical to a guard that declares nothing — and an undeclared guard makes the run **exit 2 naming its own filename**; a runner holding a hardcoded fence list would be §2's deleted hand-maintained tuple one level up, failing the same way, and iter-08 already measured the generalisation (*a fence only ever asserts about what it already scans*). All **7** stack-core guards now declare (1 `postcondition`, 6 `standalone`); a typo like `"postcondtion"` is refused like an absence. (2) A site's identity is **`(fence, path, claim_id)` with NO line number** — keyed on `file:line`, adding a paragraph above an accepted site would read as an induced defect, and §8 rule 6 says where crying wolf ends; lines are still printed on every RED, they are simply not the identity. (3) The baseline is a **ratchet**: `--accept` lowers, and **refuses to raise** for a fence already recorded, naming the sites that grew — while a fence **absent** from the baseline is the *other* state, a **registration** rather than a regression, admitted and **announced on stdout as a baseline rather than a pass** (the distinction iter-45 depends on, since it registers two more fences). **Watched going RED on the RED that did not previously exist:** the claim-twin suite already pins *"this corpus has 18 known defects"*; nothing pinned *"a repaired tree just grew a nineteenth"*. The anti-fixture is therefore **assembled from the answer key, never written** — the 18 GREEN twins plus one captured RED file dropped back in — because a hand-written induced defect encodes a guess about the class while this one **is** the class, byte for byte at rosetta `48ca53c`; and the repaired tree is asserted **silent first**, or a fence that reddens on everything would pass. **12-mutant battery, every mutant matching its declared verdict:** 1 declared-GREEN no-op that **survived**, 11 kills with **11 distinct failure signatures**, 5 of them inversions, baseline GREEN before *and* after, every mutant `py_compile`d first. **Four of the eleven do nothing but SILENCE a reporting path** — the direct lesson of harden passes 7–9, where two of the claim-twin fence's own honesty mechanisms turned out not to exist and one deleted clean with 15/15 green: *a reporting path with no mutant is a docstring*. **Two vehicles, and the weaker one is labelled** (`D-M257x-44-5`): the **suite** grades the live tree in every clone and is load-bearing; the `--install-hook` pre-commit is a latency optimisation, **per-clone and unversioned** — the same shape as iter-01's git-ignored `rext.tag` that *"never appears in a diff and drifts unseen"* — disclosed in the docstring, the protocol section and the installer's own output, and scoped to commits that stage a published path so it does not earn a `--no-verify` reflex. **Two defects were found in this iteration's own tests while writing them**, both the milestone's recurring shape: an exec-bit assertion that ran **after** its temp dir was deleted (fails for a reason unrelated to its subject; the inverted form would have passed forever), and a baseline built with an **absent** fence entry instead of an **empty** one — so the planted induced site graded as a *registration* and the test read 0. That is precisely the distinction `D-M257x-44-3` draws, violated by its own author within the hour: the **fourth consecutive occurrence** of *"the author of a newly-written rule violated it while writing it."* Protocol gained a §8 section (*run the fence at the COMMIT, or it cannot touch the number that matters*); `stack-core/README.md` gained 2 rows. **NOTHING was repaired** (`D-M257x-42-3` binding, `D-M257x-44-7`) — the only corpus edit is new text, and the post-condition was run over it: 18 sites, byte-identical set, exit 0. `stack-core` **450 tests / 14F** vs a 420/14F baseline (+30 tests, 0 new failures; the 14 are the pre-existing `test_buildbench`-rooted batteries). Platform origin `2adcf71` re-fetched at open and close, unchanged. **Clause 5 stays at 18 by construction — the gate stays 4 of 5** — see iter-44/progress.md

- iter-45 (tik, **TOK-02 step 3 of 5**): **the three mechanical fences exist, and their RED watch is built to outlive the repair that is about to erase the evidence.** `markdown_structure_guard.py` (orphaned list resumption · doubled function word · unbalanced code fence · table-row width), `anchor_construct_guard.py` (*the anchor resolves — but does it name a **construct**?* — blank line, bare closing delimiter, table separator, table **header** row, past EOF), and `derived_value_guard.py` (`**Language**: Go X.Y` vs the service's own `go.mod`; `<cpu> CPU / <mem> MB` vs `terraform/locals.tf`) — all three enrolled in iter-44's commit-time ratchet via a statically-read `FENCE_KIND`. Together they reach **5 more of the 18** (#6, #10, #11, #13, #16), which is exactly what iter-42's classification predicted these instruments would reach. **Every rule was measured tree-wide BEFORE adoption and the rejected drafts are recorded next to the rule that replaced them** (§5 rule 2 applied to lint rules): M1's first draft was **86% false positive** — 6 of 7 hits were an INDENTED blockquote inside a list item, which is correct markdown — M2 needed a hyphen guard on **both** sides of the pair (`in in-app`, then `default-ON on demo`), and the anchor self-reference rule went from **134 findings, essentially all of them ports**, to 2, both real, by requiring the `:NNN` to be *immediately* followed by above/below/earlier. **Blocker #17 was DROPPED rather than tuned in** (`D-M257x-45-3`): the enumeration rule that would have caught it fired on 6 of 7 tree-wide candidates and still missed it, and narrowing the window until #17 fires while its neighbours do not is Trap A — routed to step 4's hand repair, named. **The iteration's central finding is about its own tests.** The first draft of the behaviour suite asserted all five blockers against the LIVE corpus and passed — and **every one of those assertions would have failed at iter-46**, whose entire job is repairing those five sites, with the obvious repair being to edit the fence's own test to match. *A fence stops asserting anything by being maintained, not by being deleted.* So `tests/fixtures/mechanical/` captures two **LINE-FAITHFUL repo roots** — not the ±2-line neighbourhoods iter-43 used, because two of these five defects are **relationships between line numbers** (`:788` citing `:447` in the same file; `:110` citing line 815 of another) which an excerpt destroys — with the one large vendored platform source under **line-preserving elision** rather than compaction. `red/` fires all five at the exact anchors; `green/`, produced by **declared mechanical transforms** and never by hand, is silent **while still being resolved and measured**, because a fence that stopped reaching a site prints the same clean run as one that passed it. Generalised into `platform-alignment.md` as **§8 rule 7**. **The 20-mutant battery earned its keep on first run** — 1 declared-GREEN no-op that survived, 19 kills, ≥5 inversions, ≥5 distinct signatures, and **one mutant per reporting path in all three modules** (the harden 7–9 debt: *delete each new fence's reporting path and confirm a test fails*) — because **three mutants came back GREEN**, each naming a real hole in a suite that looked complete: M1 tested on only one of its two column-0 sides, a `corpus/` with zero scannable files reading as clean, and `measured` counting a doc no scalar was ever read from. All three closed with named tests. **Two defects were found in neighbouring fences and fixed at the fence, not worked around.** (1) **iter-44's own ratchet rewrote a record it had not moved**: `--accept --reason` overwrote `claim_twin_guard`'s registration sentence with one about three fences that postdate it — `registered_at` already preferred the prior value, `reason` did not, and *that asymmetry is what made the loss silent in review*. Reason restored verbatim, two regression tests (both directions — a fence that IS lowered must still take the new reason), one new mutant in iter-44's battery. **A record that silently rewrites itself is this milestone's own defect class, found this time in the instrument built to prevent it.** (2) **A captured fixture was read as this repository's source**: the vendored `assignments.go` made `test_write_target_schema_fence` report `stack-core` — a section that ships no Go — as an unclassified Go-bearing rext section, and the next step would have been scoring the platform's source for rext's schema writes. Fixed by pruning `fixtures` **directly under `tests`** from all three of its walks, with a two-directional regression test; the tempting alternative (classify `stack-core` as `n/a`) is the exact move that fence's own failure message forbids. **NOTHING was repaired** — clause 5 stays at **18** by construction and **the gate stays 4 of 5**. `stack-core` **491 tests / 14F**, exactly the pre-existing baseline (the one new failure this iteration caused was fixed, not accepted). **rext pin NOT moved** (`D-M257x-45-9`): everything here is offline guard/test code on no runtime path, so re-pinning would change what the next gate measurement runs against for no benefit. Platform origin `2adcf71` re-fetched at open and close, unchanged — see iter-45/progress.md

- iter-46 (tik, **TOK-02 step 4 of 5**): **all 18 of iter-41's blockers repaired in one pass, by CLAIM and tree-wide — and the fence caught three sites the repair itself left standing.** Every claim's *what is true* was re-derived from platform source at `app` @ `5ba17044` before a word was written, never from the ledger's prose (`D-M257x-46-1`), and **two came back stronger than their own adjudication**: `#2`'s uncounted exit is the `default:` arm reached by a **nullable-and-unset** `AIVendor` — the ordinary path, not a mistyped string — plus a **fifth** route nobody had counted (a caller explicitly selecting `Openai`); and `#4`'s Anthropic-direct path is selected by an **env var, not a feature flag**, so the `flag_use_azure_us` caveat four lines below it never covered it. **`#5`'s base count, which iter-41 deliberately left unsettled between two auditors (17+7 vs 16+7), was settled by counting** (`D-M257x-46-2`): 139 schema files, 30 `OrganizationMixin{}`, 7 `OrganizationIDMixin{}`, 18 plain-column, and **only FOUR files in the whole directory declare any `Policy()`** — E's 16 is right and the total is **23**; the defect was never the 16 but a closing sentence excluding the 7 **three lines after naming them as unpoliced**, the fifth consecutive failure of that fence and the second direction it has failed in. `#17` was repaired **by hand** exactly as `D-M257x-45-3` routed it — its anchor resolves *and* names a construct, just the wrong one for 3 of the 4 domains, which no fence in this family can see. **Then the measurement that justifies the whole of TOK-02:** with all 18 anchored sites fixed, `claim_twin_guard` still reported **three** — a correction *appended* while the original undercount sentence stayed, a second EU-residency site untouched because only the PM summary had been fixed, and a phrasing that survived a rewrite. **Every one is *"repaired at one site, left standing at another"*, the class iter-41 measured as 8 of the 9 repair-induced blockers.** Under the previous method all three would have been committed and counted by the next full read; here they were named before the commit and closed in the same pass. **This is the first pass in the series where that class was caught by a machine rather than by the next audit.** Two further live findings outside the 18 — a stray unterminated code fence and an anchor onto a blank line — were **repaired rather than baselined** (`D-M257x-46-4`), because a ratchet that accumulates what nobody wants to fix stops being one. **All four fences GREEN; the ratchet baseline lowered 25 → 0 sites across 4 fences.** And **GREEN was falsified rather than celebrated** (`D-M257x-46-5`, §5 rule 8 — four fences going 25→0 in a single pass is precisely the shape that rule warns about): reach did not shrink (`anchor_construct_guard` resolves **101** anchors across 112 files, *up*, since repaired anchors now resolve; `derived_value_guard` measures 5 service docs), **both perishable fixtures still go RED with their green twins silent** (53 tests), and `stack-core` holds at **491 tests / 14F**, exactly the pre-existing baseline. 17 files touched by claim rather than by file, spanning `corpus/architecture/**`, `corpus/services/**`, `corpus/ops/demo/**`, `.claude/skills/**` and `CLAUDE.md`. **Clause 5 is NOT graded here** (`D-M257x-46-7`) — only TOK-02 step 5's full 7-auditor read at iter-41's frozen instrument grades it, and reading a fence's GREEN as that number is the mistake iter-38 and iter-21 both paid for; TOK-02's pre-registered prediction (*the step-5 reading returns fewer than 9*) stands unmodified and testable. **Gate stays 4 of 5.** rext pin not moved (`D-M257x-46-8`) — corpus prose plus a baseline JSON, no runtime source. Platform origin `2adcf71` re-fetched at open and close, unchanged — see iter-46/progress.md

- iter-47 (tik, **TOK-02 step 5 of 5 — the last**): **the seventh pass returns 7 blockers, down from 18 on an identical instrument — and the DECOMPOSITION is the result: 0 pre-existing, 7 induced.** Seven auditors, **40 files / 9,243 lines**, instrument frozen at iter-41's on every knob (six full-read partitions + one adversarial diff-reader; same briefing; same size-sort snake-deal **method**, which re-dealt **20 of 40 files** into different hands because iter-46's repair moved the corpus 9,163→9,243 lines — §5 rule 18(b) satisfied *by* the method rather than despite it). Per-file `wc -l` positive control: **all 40 confirmed line-for-line**. 8 raw findings, `G1 ≡ B1` reached independently from two different seats → **7 unique**, and **7 of 7 HELD on this iteration's own re-derivation** (`adjudication.md`) — none refuted. **The headline is what six passes could never separate: six full-read auditors covering all 40 files found ZERO blockers in text iter-46 did not touch.** Every one of the 7 is in text it wrote or rewrote — **4 newly-written over-corrections** (a `nullable` pointer that is a **value**, normalized nil→`Openai` at the sole construction site so *unset* never reaches the `default:` arm it was said to define — inverting the very distinction the paragraph was written to fix, and contradicting the file its own next line calls *"the full per-line derivation"*; Studio-Room *"flipped off Bedrock"* when `grep -rin 'bedrock\|boto3' app/studio/` returns **0 hits**; *"the **two** US paths"* omitting the unconditional third; and an anchor transcribed from iter-41's ledger onto `external_services.md:489`, which is a **TypeScript codegen comment**) and **3 pure leaks** (the Directus *"has never existed"* over-correction repaired at `service_taxonomy.md` and **left standing at its twin** in `external_services.md:139`; the ordered EU-first arrow chain rewritten in three other files and surviving **68 lines below its own file's fence**; and the aireadiness *"always takes the live-recompute branch"* standing **13 lines above** the fence iter-46 added). **TOK-02's pre-registered *"fewer than 9"* is CONFIRMED at 7 — the first confirmed prediction in the series**, after four consecutive passes refuted their own; its sub-predictions on the self-contradiction class (**3** < 4) and on over-correction in explanatory text (**4** ≥ 1) also held, while the third — *the residual is NOT concentrated in the repaired files* — was **refuted absolutely, 7 of 7**. **iter-41's *"the fixed point of this process is not zero"* is now refutable in a specific way: the CORPUS term reached zero; the REPAIR term did not.** **All four fences correctly report 0 sites, and the gap is nameable rather than a failure** (`D-M257x-47-5`): `claim_twin_guard` matches **adjudicated** refuted forms and six of the seven are *new prose with no ledger entry*; `anchor_construct_guard` is silent on the bad anchor because `:489` **resolves and carries content** — the *wrong* construct, exactly the class `D-M257x-45-3` declined to tune a fence for. **But three of the seven need no new idea at all — each is a grep for a string still sitting beside its own repaired twin, and a leak-check over the repair's own diff (*for every claim this commit changed, grep the tree for the old form*) would have caught all three.** That is the highest-value mechanical finding here and is routed as `FENCE-M257x-iter47-leak-check`. The failure mode is narrow and specific (`D-M257x-47-6`): **iter-46 re-derived every QUANTITY rigorously — all four wiring anchors, the whole Ent chain (139/30/7/18/only-4-`Policy()`, with 18−2=16 and 16+7=23 and both subtractions justified), Go 1.26, 256/128, the `gpt-4o` slots, the Directus git evidence exact to the line including its hedge — and narrated MECHANISMS loosely.** **NOTHING was repaired** (`D-M257x-47-2`) — a clause is met by a reading that returns zero, never by a repair clearing its own findings, and repairing here would have made the number unfalsifiable. **64 minors** recorded, two worth promoting (iter-46's new `service_taxonomy.md:150-153` table cell spans four physical lines and **will not render as one row**; `hiring.md:189-196`'s "minimal write-set" omits a NOT NULL + UNIQUE column). **Clause 5 NOT MET at 7 — gate stays 4 of 5 — and TOK-02's five steps are now EXHAUSTED**, with the induced term it was built to attack now constituting the entire residual (`D-M257x-47-7`). rext pin not moved; no code written. Platform origin `2adcf71` re-fetched at open and close, unchanged — see iter-47/progress.md

- iter-48 (tik, **TOK-02's own pattern applied once more — fence, repair, read**): **the eighth pass returns 12 blockers, and iter-47's "0 pre-existing / 7 induced" is REFUTED by an inversion — 10 not induced, 2 induced, 7 predating the milestone entirely.** Three planned steps, all landed. (1) **`FENCE-M257x-iter48-repair-leak` built FIRST and watched RED** on the iter-46 fixture before the repair (§5 rule 21): 5 sites, 2 of them iter-47 blockers, **1 a real defect no auditor reported** (`external_services.md:565`, in a file six of seven seats read top-to-bottom), 2 benign and **named in a test rather than tuned away** — the 2-in-5 false-positive rate **pinned in both directions** so tuning it to zero fails as loudly as letting it grow. **K was MEASURED, not chosen** (`D-M257x-48-2`): the longest common token run a real editorial rewrite leaves is **8** (a rewrite that inserts one article and one verb), so a "comfortable round number" above it is blind to the ordinary shape of a leak. Twenty mutants, five inversions, one no-op control required to survive — and **one further mutant survived the first run** (`--json` emptied with the suite still green, the exact *a reporting path with no test is a docstring* shape harden 7–9 found twice inside `claim_twin_guard`), **closed with a test rather than booked as a known gap** (`D-M257x-48-5`). (2) **The seven repaired by CLAIM, not by file** (§5 rule 19), tree-wide, leak fence **GREEN on the repair's own diff**; the claim-twin ratchet went 12 sites → 0 and `stack-core` returned to its pre-existing baseline. (3) **The eighth reading** — seven seats, iter-41's instrument frozen on **every** knob, all 40 files top-to-bottom with per-file `wc -l` positive controls, plus the diff seat. **12 raw → 12 unique → 12 held**, each re-derived against `app @ 5ba17044` before acceptance. **The decisive result is not the count, it is WHERE the findings are:** seven sit in text **neither repair touched**, authored **2026-03-02 .. 2026-07-23**, *inside seats' own assigned file sets* — the LiveKit agent names (`anthropos-agent-eu` exists **nowhere**; EU is the bare `anthropos-agent`), a NOT NULL + UNIQUE + undefaulted `token` column missing from `hiring.md`'s "minimal write-set" (**iter-47 saw this exact passage and booked it a MINOR**), Storage's `dependency_map.md` row contradicted by its own twin doc, and **two found only now by seat B** — `keepStartedMembers` described as excluding members with no step-1 signal when it reads a **progress row** and no signal at all (wrong in both directions, and it is the stated justification for the seeding contract three lines below), and `flag_use_realtime_openai` described as gating LiveKit-vs-ElevenLabs for "new sessions" when engine choice is **per sequence from the CMS `voice_engine` field** and the flag only swaps a dispatched **endpoint** to `openai-hosted` from **inside** the LiveKit path already entered. **So iter-47's zero was a property of the READING, not of the corpus** — its two-term model is right about the arithmetic and wrong about one of its inputs. **Not one of the 7 was reachable by any shipped fence**, traced individually: `repair_leak_guard` is verbatim-only (`D-M257x-48-4` pins **paraphrase** out of reach; `D-M257x-48-9` pins **number-only** corrections out of reach *and* measures that lowering K buys two false positives without catching them, so K was **not** the binding constraint on the miss), `claim_twin_guard` fires only on claims already in a ledger, and `anchor_construct_guard` is silent because all 12 sit at **valid** anchors. **Three of five pre-registered predictions REFUTED, and the two that held were the least informative** — the heavily-repaired block predicted to carry the residual was verified clean line-by-line by two seats; the untouched files predicted fine were not. **The series is 25 → 13 → 11 → 17 → 37 → 18 → 7 → 12: every better instrument found MORE, and the last three readings used an instrument that did not change at all.** Run-to-run variance of the **frozen** instrument is **larger than the residual being chased**, so *a reading that returns zero is evidence about the reading* — which is exactly what iter-47's zero proved to be one pass later. **`claim_twin_iter48/` captures the perishable answer key** (18 red / 18 green, pinned at rosetta `cabc3b1`) **before** any repair can destroy it — the only artifact that can support *a full seven-auditor reading missed these while they sat in its own file sets*. **NOTHING repaired** (deliberate — the orchestrator asked for the honest number and split first). **Clause 5 NOT MET at 12 — gate stays 4 of 5** — and the escalation is the **stronger** version of the one `overview.md` pre-declared: not "clean except for the act of cleaning it", but a residual the cleaning barely touches. **`EXIT_REASON: user-blocker`.** Generalised into `platform-alignment.md` as **§5 rule 22** (*a frozen instrument is not a precise instrument*). `FENCE-M257x-iter49-numeric-leak` routed forward (Fate 3) — reaches 2 of the 7. Platform origin `2adcf71` re-fetched at open and close, unchanged — see iter-48/progress.md
- iter-49 (tik): **the two fences iter-48 named both shipped and both closed their gaps — and the ninth reading still returned 14, split 7 induced / 7 pre-existing, against a pre-registered 6 (2/4) that is refuted in every term.** `value_change_guard` (FENCE-M257x-iter49-numeric-leak) asks a *different question* from the verbatim leak fence — *did the old VALUE survive?* — by word-diffing each hunk's removed against added text and keeping the pairing `repair_leak_guard` discards; `D-M257x-48-9` had already measured that the miss was **structural, not a threshold** (lowering K bought two false positives and still missed). Watched RED on rosetta `301d61a`, a real already-committed incomplete repair: **5 sites, of which 3 were blockers two different audits had independently adjudicated** (iter-48 #3, iter-47 #5, iter-48 #9), and the gap it fills is **asserted against `repair_leak_guard` directly** rather than claimed. Its 2 false positives are kept and **pinned in both directions**, the stopword filter that would remove them having been measured and rejected — it costs a proven true positive (`D-M257x-49-2`). `--audit-commit` (FENCE-M257x-iter49-audit-commit-mode) closes `D-M257x-48-12`: it **writes nothing** — the ratchet's monotonicity is the contract — and admits a site only on a signature no repair can produce (the claim's **ledger row is a line this commit added** AND the site's file is one it **did not touch**), so laundering a repair would require not editing the file you were repairing while writing a genuine refutation of the claim you left standing, which is an audit. Then the 12 repaired by CLAIM tree-wide (11 files for 12 claims) — and **the post-condition caught this repair twice before the commit** (`D-M257x-49-5`: an anchor pushed onto a blank line by the repair's own insertions; a leak into `ai-readiness.md`), after which the **new** fence found a **third** site of that same claim (`seeding-spec.md:497`) which the verbatim fence had already passed GREEN, adjudicated a legitimate retraction and acknowledged with its reason. Ratchet **18 sites → 0**; `stack-core` **22F → 14F**, the 8 traced failures clearing exactly. **The reading's finding is the strategy-level one:** the 7 induced defects partition into **paraphrase leak (3)**, **overshoot in new text (3)** and **wrong-mechanism-correctly-cited (1)** — and **not one is mechanically reachable**; a paraphrase shares no token run (the limit `D-M257x-48-4` pinned and this iteration's own docstring re-pins), and an overshoot lives in prose that did not exist before the commit, so every diff-relative fence is silent by construction. **So TOK-02 step 2's premise is now true of a class that has stopped being the majority: mechanising the mechanical half did not lower the total, it changed what the remainder is made of**, and §5 rule 21's cheapest-instrument classification must be re-run *after each instrument lands*, not once. Seven pre-existing blockers surfaced in files six of seven seats had read top-to-bottom one iteration earlier (`roadrunner.md` ×2, `hiring.md` ×3, `shared_libraries.md`, and a 31-vs-32 count in the passage whose own text says it has been wrong four times), while Seat F read 1,498 lines across six files and found **zero** — so the corpus term is real and is not "auditors always find something." Nine readings now stand at `25 → 13 → 11 → 17 → 37 → 18 → 7 → 12 → 14`, variance ~±5 (§5 rule 22 reinforced). Also: five of this iteration's own tests were invalidated by this iteration's own repair — **the exact failure §8 rule 7 was written to prevent, at iter-45, by this milestone** — fixed with a hermetic temp-repo fixture and recorded as a recurrence corollary that generalizes the rule beyond fences (`D-M257x-49-9`). Gate NOT MET; the 14 route forward as `FIX-M257x-iter49-blocker-set`, with `FENCE-M257x-iter50-paraphrase-leak` and **`CHECK-M257x-iter49-overshoot-has-no-instrument`** named as the two handlers the classification now points at — see `iter-49/progress.md`.
- iter-50 (tik): **THE VARIANCE EXPERIMENT — §5 rule 22's own prescription, finally run, and it cost ONE reading.** The same tree read twice with no repair between (40 files byte-identical, 13 clone shas identical, the partition therefore dealing the **same hand**, seat G given the identical diff, and every seat **blind** — barred from `knowledge/plan/**`, so no partition, ground-truth or corpus confound survives): reading #9 returned **14**, reading #10 returns **7**, and they agree on **4**. **Union 18. Recall 29% and 57%. Chapman `N̂` ≈ 23 — and it is a FLOOR**, because heterogeneous detectability biases capture–recapture downward. **A single reading is a SAMPLE, not a census**: two full 7-seat passes named 18 findings between them and neither named more than 14; one of #10's four new findings (`dependency_map.md:19`) sits **inside a hunk seat G reviewed and passed at #9**. **This explains four iterations at once** — a repair pass can only repair what a reading NAMES, so with recall ≈ 0.43 and a non-zero induction rate the repair-then-read cycle has a **fixed point**, and it sits exactly where `18 → 7 → 12 → 14 → 7` has been sitting. iter-47's "zero pre-existing" was a low-recall draw, reproduced here under control. **Pre-registration graded 2 of 4** — the overlap (<7 of 14) and union (>14) predictions HELD; the count band [9,19] and the symmetry prediction were REFUTED. **And the pass indicts itself:** three seats independently cleared *"31 of 135 schemas auto-filter by organization"* as a **positively audited zero** and **all three are wrong** (`organization.go:56` declares its own `Policy()` with `FilterSameOrganizations()`, uses neither mixin → **32**) — each re-derived the arithmetic the document showed instead of the predicate it claimed, which is §5 rule 17 violated three times in one pass by auditors briefed on it. Reading #9 had it right. Ships **§5 rules 23 + 24**, `fixture-14.md` (the perishable key, captured first), and **`FIX-M257x-iter50-union-set` (18)**, which supersedes the 14 — see `iter-50/progress.md`
- iter-51 (tok, TRIGGERED — 3-no-prog streak, iters 48/49/50): **`TOK-03: repair the UNION, shrink the estimator, make the edits smaller`.** Step 0 verified the trigger is not stale (platform origin `2adcf714` unchanged; corpus unrepaired; union at 18 with `N̂` ≈ 23). **The conclusion: TOK-02 was optimising the wrong term.** Two strategies were spent making the READING better; iter-50 measured that a single 7-seat pass finds ~43% of what is there, so **a repair can only ever repair 43% of the pool** and repair-then-read has a **fixed point** — sitting exactly where the last five readings landed. Four moves, ordered by dependency: **(1) repair the UNION of two blind readings, never one** (18 findings / 78% coverage vs 14 / 61%; a cycle is now *read, read, repair*); **(2) drive `N̂` down FIRST and take clause 5's reading when it is small** — at `N̂` ≈ 23 a zero reading has probability ≈ 10⁻⁵, at `N̂` ≈ 2 it is reachable, and `N̂` has the floor-of-zero-by-construction property §5 rule 22 asks for, which the raw count never had; **(3) cut induction by shrinking the EDIT, not by adding fences** — the two unreachable induced classes (paraphrase leak, overshoot-in-new-text) are both properties of *rewriting*, so prefer DELETION > minimal scoping edit > rewrite and budget the added words; **(4) put two blind adversarial readers on the repair diff BEFORE the commit** — seat G has been the highest-yield seat in both readings and has always run one pass too late, which is `CHECK-M257x-iter49-overshoot-has-no-instrument` answered with the reader iter-49 said it would need. **Explicitly NOT done:** clause 5 not re-cut / narrowed / read met any other way (the user has ruled twice; TOK-03 changes only what happens *before* the reading), residual not deferred, instrument not weakened (it runs TWO frozen instruments, not one cheaper one), and every TOK-01/TOK-02 fence kept. **Strategy class `new-direction`** — the first revision in this milestone authored on a measured recall rather than an inferred pattern. **Pre-registered, refutable: `N̂` below 12 and the induced term below 4 at the next paired reading.** Next tik = `FIX-M257x-iter50-union-set` (the 18, by CLAIM, minimal-edit, two pre-commit diff readers), then the paired reading #11+#12 — see `iter-51/progress.md`
- iter-52 (tik): repaired the union of 18; two blind pre-commit readers found 5 blockers in the repair itself; claim_twin_guard 31 -> 0; N-hat re-derived ~23 -> ~14 (floor) — see iter-52/progress.md
- iter-53 (tik): **the paired reading #11+#12 returned 32 and 26, matching on 12 — union 46, recall 37.5%/46.2%, Chapman `N̂ ≈ 68` — and `N̂` went UP, not down.** TOK-03's two pre-registrations for this iteration are **both refuted** (`N̂ < 12` → ≈68; induced `< 4` → 9), as are three of the iteration's own five. **But the number is not the finding.** The instrument §5 rule 22 declares *"frozen at iter-41, never touched a knob again"* turned out to be a **git-ignored file** — `.agentspace/scratch/work-m257x/iter50-briefing.md`, in no commit and no iter dir — so every pass has "held it fixed" by **re-authoring it from a one-line summary in the previous iter's `overview.md`**. This iteration did exactly that, unknowingly, and the drift concentrated in one clause: the canonical rule resolves doubt **downward** (*"if you cannot cite the refutation, it is not a blocker"*, with undercount / omitted list member / line drift carved out as MINOR) while the re-authored rule resolved it **upward** (*"when in doubt, book it as a BLOCKER"*, no carve-outs). Re-grading the union against the canonical rule **verbatim** gives `23 / 23 / m=11 / union=35 / N̂ ≈ 47` — so **roughly half the jump over #9/#10 is grading drift and half is not, and neither half is a statement about the corpus.** The consequence is larger than this iteration: **the whole published series `25 → 13 → 11 → 17 → 37 → 18 → 7 → 12 → 14 → 7` is not a comparable series.** It is the same root cause iter-01 found for the *pins* — a git-ignored source of truth that never appears in a diff — sitting under the *audit* the entire time. **The one quantity that survives:** per-finding recall of a single 7-seat pass measured 43% (iter-50), 42% (as run) and 48% (re-graded) — **it replicates across two grading rules and two trees while the count does not.** Corrective action landed: the canonical briefing is now **committed** at `instrument/briefing-canonical-iter41.md` with the as-run one beside it as evidence, protocol **§5 rule 25** added (*an instrument that is described rather than stored is not frozen*), and rule 22's false sentence corrected in place. Fixture intact — **nothing repaired**, deliberately, so iter-54 inherits a live 46-row answer key. Exits **user-blocker**: the repair target (46 as-run vs 35 canonical) and whether the milestone's narrative needs re-baselining are decisions that must not be made on one seat's authority — that is exactly how the false `32` entered at iter-49. — see iter-53/progress.md
  - **AMENDMENT (close-fetch):** platform origin **moved mid-iteration**, `2adcf714 → ef32d4cd` — `d11a403 chore(compose): drop roadrunner, prune dead env, repoint messenger`, squarely this milestone's class. Reading #12 seat F's clearance that *"compose still starts cms/jobsimulation/roadrunner in the default `graphql` profile"* is **invalidated by name**. Phase 5 grading corrected: **(3) re-scope: y → `EXIT_REASON: re-scope-trigger`, occurrence 2 of 2**, which the milestone's `re_scope_trigger` field says must STOP and escalate to a **pinning-and-tracking policy**, not more alignment work. It compounds the iteration's own finding: **two of the three legs of every number this milestone has published — the instrument and the ground-truth reference — were free to move without appearing in any diff.** Status unchanged at `closed-fixed`; the measurement completed before the platform moved and is labelled as being against `2adcf714`.
- iter-54 (tok, TRIGGERED — the milestone's own `re_scope_trigger`, occurrence 2 of 2, plus a direct user ruling; NOT the 3-no-prog streak, which was checked and does not apply): **`TOK-04: pin the target, or stop calling it a measurement`.** Three jobs. **(1) The platform re-survey**: `stack-demo/platform` fast-forwarded `2adcf714 → ef32d4cd` (0 behind, zero platform-repo edits), and the clause-3 fence run **before any corpus edit** went **RED on 3 real direction-B departures** — cms, jobsimulation and roadrunner leaving `repos.yml` — then GREEN once the map was updated. **The first time a fence this milestone built caught a real departure it was not shown**, rather than a staged defect. `d11a403` deletes those three compose services *and* their clone entries (`make init` no longer clones them), restores onto `backend` the env the merged code still reads in-process, and repoints `messenger`'s last two RPC edges at the monolith; `6060315` changes the bring-up **timing contract** (`start_period: 120s`). **81 new drift sites across 21 files, from 3 commits inside one working day.** **(2) The honest reassessment**: **the gate is 2 of 5 against origin HEAD, not the booked 4 of 5** — nothing regressed, but clauses 1 (met at `2adcf71`, iter-14) and 2 (met at iter-37) are **STALE by the gate's own "against origin HEAD (never a pinned pre-drift commit)" wording**, and cost ~40 min of machine time to restore with a pre-registered expectation of green. Clause 3 **MET, re-met today**, with the caveat that the guard fences *membership* and not prose — and said so in its own header — while **this iter committed a false claim into the map 30 minutes earlier** (*"the armed failure is now armed"*, citing `migrate-demo.sh:81-85`/`:106`, code that this milestone's **own iter-02** deleted at rext `54bccf7`): quoted forward from iter-01 without re-measuring against its own repair, corrected in place and left visible. A second dead claim was found beside it — §5's `storage`/`messenger` watch-signal was already true and could never fire, the map committing Trap A from its own §1. **Clause 4 MET — and now met UNDER TEST**: verified by *running* the derivation against the new `repos.yml`, not by reading it (`app:public` / `extensions sentinel public` / transitional debt **empty**), byte-identical to the reading at `2adcf71` and identical *correctly*, because the three departing repos declared `migrations: false` and no `schema:` key. **The time bomb iter-01 named armed itself six weeks ahead of M810 and nothing happened, with zero human action** — the founding thesis surviving an event nobody arranged. **Clause 5 NOT MET and further away**: the series is non-comparable (unfrozen instrument, inverted tie-break), recall ≈43–48% survives, union 46 as-run / 35 canonical with 9 of 46 induced — and **+81 today**, so the cycle nets **−72**. **(3) TOK-04**: the pinning-and-tracking policy the trigger prescribed by name, generalized past the platform repo because the class has bitten **four** of this milestone's own instruments — the rext pin (git-ignored), the briefing (git-ignored), the clone (free to move), and **clause 2's gate-meeting run, which recorded no ref at all**. Four rules: **P1** every measurement states its refs in the artifact · **P2** every instrument is a committed file and nothing it depends on is git-ignored · **P3** the ref is chosen, recorded and re-checked at open *and* close, and the detecting iter re-points it *in that iter* · **P4** derive, else fence, else declare prose-under-review — **an order measured on a single event** (zero human action / unaided 3-for-3 catch / falsified in a day). TOK-03's three moves are all kept; its **premise** is refuted — the residual is a **flow, not a stock**, and the old metric could not go negative so it could never report that we were losing. Generalized into the protocol as **§5 rules 26 and 27**. Next: **iter-55 = re-establish the ref baseline** (clause 2 ~5 min, clause 1 ~35 min, each with a `refs:` block; pre-registered **both green**), then the **harden pass — due at 10 tiks but sequenced AFTER iter-55**, because its named residue is all Playthrough/seeder surface and hardening it against a stale ref would repeat the mistake this tok exists to name, then **iter-56 = the 81 sites as ONE derived-and-fenced class**. Exit `tok-fired` — see iter-54/progress.md
- iter-55 (tik): **the ref baseline, re-established — and TOK-04's pre-registration of "both green" refuted on both clauses.** P3 fired before the first measurement: origin had moved again (`ef32d4c → 0dab54d`, *"run without the standalone storage; rename graphql -> core"* — the v9.0 `support-in-app` step §1 already listed as IN FLIGHT), so the clone was fast-forwarded **inside the iteration**, second consecutive time. **Phase A replaced three hand-maintained topology tuples with one ref-independent derivation** (`stack-injection/platform_topology.py`): the profile name (5 literal sites — and a stale profile name is *not an error to compose*, it selects an empty service set and **starts zero containers**), `verify_svcs` (**5 of its 10 services no longer existed**), and `INJECT_SVCS` (2 of 3 deleted; it had not failed only because the pre-prune clones are still on disk — §2's time bomb, armed by the next clean checkout). Same code reads `graphql`+storage at `ef32d4c` and `core`−storage at `0dab54d` with no edit between, which is the property a corrected literal could never have. Two defects introduced **and caught**: 82 green tests turned red by deriving at source time (lib-only seam has no `log` yet), and a live fence that began **skipping** — caught only because the skip count was read (§5 rule 8). **Then the cold cycle found that the teardown removes nothing:** `down --purge` printed a clean teardown and exited **0 with eleven containers still running**, because a 45-h-old injected override named deleted services and compose refused the *whole project* as invalid — `|| true` discarded it, and `purge_data_dir` then deleted the database out from under eleven live containers. Fixed by asking **Docker's project label**, not the compose file; **it caught a real survivor on its first live run** (`demo-1-storage-1`) by a *second* mechanism — at `0dab54d` storage is declared under `storage-legacy`, so it is **not an orphan** while not being in `core`, so it is **not selected**: a container can fall through both. **Cycle A then ran to completion and went RED** — 3 checks, one cause: `backend` exits **0 in silence**; the compose deleted `STORAGE_RPC_ADDR` while the pinned `app v1.363.2` reads it at `main.go:446/516/983`, and `v1.364.0/v1.364.1/v1.365.0` exist on origin. **The compose half of v9.0 landed; the app half is not in the pinned release** — Trap D, and the exact break §9's watch bullet predicted by name, which still cost a full cold cycle because *a watch item with no fence is a note*. Everything else was green (casbin 1251, 21 directus collections, 42 790 skills, patches all applied, cockpit + FAPI + hiring + academy). Clause 2's binding reading is therefore impossible (a 30-fail run against a dead container measures nothing); the 44-h-old-stack reading is recorded as a **control**, not a restoration. **Clause 3 re-read and DOWNGRADED: the membership fence is GREEN and the map is FALSE** — 5 claims falsified by `0dab54d`, two citations now resolving to unrelated lines; **not repaired on one seat's reading**. Clause 4 **MET under test** (35 tests against the new `repos.yml`). **Gate: 1 of 5 at origin HEAD, not the booked 2 of 5 — a downgrade produced by looking.** Ships §8's *"A TEARDOWN is a write path too"*. Exits `user-blocker`: the fix for clause 1 is a deliberate `app` pin advance (§7 rule 4), the class that broke seeders at v2.1 and v2.7 — see iter-55/progress.md
- iter-56 (tik): **CLAUSES 1 AND 2 RESTORED — gate 1 of 5 → 3 of 5.** Three consecutive cold `down --purge` + `up` cycles at platform origin HEAD `0dab54d`, all `green:true / warnings:0` (`14:47:13Z`, `15:26:37Z`, `15:34:09Z`; 8m53s / 7m14s / 7m07s; teardown `survivors=0` each; 11/11 containers), and the full Playthrough suite `passing=30 failing=0 unimplemented=1 unimplementable=0` — each with a P1 `refs:` block. **The blocker was not what iter-55 named, and the pin advance it routed could not have fixed it:** `v1.365.0` reads `STORAGE_RPC_ADDR` at the same three sites and **IS** app origin/main (`rev-list --count` = 0), so no build exists in which that read is gone. Refuted before acting, by two commands. The real cause, by experiment on the same image and network: `~/.aws/credentials` does not exist on this host, **Docker creates a missing bind source as an empty DIRECTORY**, the AWS config load dies with `is a directory`, and the app **exits 0** — same env with no mounts starts and serves `:8082`; same env + that mount reproduces the dead 2-line signature; same env + a regular empty file starts. §5 Trap E on the new Mac. Fixed by **deriving** the host bind-mount precondition from the platform's own compose (`platform_topology.py` gains `host_bind_mounts()`/`check_host_mounts()`), scoped by real properties rather than a list, **watched RED on the real host then GREEN** — and the check tests for the *residue* of the failure (a path that exists as an **empty directory**), because the obvious existence-only assert reports GREEN on the exact state that produced the defect; that mutant is pinned as a negative control. The pin advance is **kept for the gate's own reason** (origin-HEAD compliance) and recorded per §7 rule 4: 37 commits, 2 migrations both `ADD COLUMN`, **0** DROP/RENAME, **0** new `log.Fatalf` — purely additive, and **the seeders survived unchanged**. Clause 2 took **two readings** (`29/1` then `30/0`) on a `pt-assignment-assign` flake asserting `toBe(before-1)` over a settling grid (`16 → 14`); reported, not buried — **PR-3 refuted as stated**. Lands `platform-alignment.md` §5 **rule 28** (*three true facts do not make a cause — join them with one experiment*) + its corollary (*check the remedy contains the fix before taking it*). P3 at close: platform **level**; **`app` moved to `v1.366.0`** during the iteration (5 commits, 0 migrations, one `fix(assignments)` whose bearing on the flake is **unmeasured**) — recorded and routed as the next iter's first act, deliberately not chased. `D-M257x-56-5`: the gate's own evidence artifacts are **git-ignored by default** (`*.log`, `knowledge/plan/**/*-report.json`) — a live instance of the routed iter-54 sweep. — see iter-56/progress.md
- iter-57 (tik): **CLAUSE 3 MET — gate 3 of 5 → 4 of 5.** The routed finding was *"a fence over membership says nothing about prose"*; the measurement says why, and it is not a wording problem: **37 of 52 citations in the map — 71% of its evidence — were invisible to every existing fence** (`anchor_construct_guard`'s regex requires a `/` or `.md`, so every bare `docker-compose.yml:90` / `repos.yml:18-20` / continuation `:161` is unseen; `platform_alignment_guard`'s assertion D checks only that a cell is NON-EMPTY). Both guards ran **GREEN over a map with five known-false claims** — a fence being green is a statement about the fence, not the file. **Assertion F** ships: every citation must resolve AND land in its own subject's compose block, with block boundaries and repo→service aliases **parsed out of compose itself** (`context: ${APP_BUILD_CONTEXT:-../app}` is how it learns `backend` IS `app` — it is not told). Its first draft returned **22 findings, 7 of them its own false positives**; narrowing until only the known-bad fired was available and **refused** as §5 Trap A, so the RULE was replaced rather than the threshold tuned — **22 → 8, false positives 7 → 0**, and the alias came free. Watched RED on the real map, then GREEN. **The hand reading that routed this work had found 2 of the 8**, in 3 rows of the 7 that were wrong. Repaired, plus 3 more real drifts the block rule structurally cannot see (declared as a limitation, not hidden), plus the substance `0dab54d` falsified: storage's v9.0 fold has LANDED (out of the default profile, `STORAGE_RPC_ADDR` deleted), messenger dropped from `all`, and the `graphql` profile is now `core`. Map §4 gains the standing **derived / fenced / range-only / prose-under-review** table P4 asks for — the third category only works if it is visible. 30/30 tests (was 19) incl. a no-op control, two inverted mutants and an exit-2 positive control — see iter-57/progress.md
- iter-58 (tik): both held-back pins advanced (app v1.366.0 + rext iter-58) and PROVEN COLD — clauses 1+2 re-established at current refs; the advance moved 22 of 23 `main.go` citations and the fence caught 1 — see iter-58/progress.md

- iter-59 (tok/triggered — by a **direct user directive**, not the 3-no-prog streak; iters 56/57/58 all closed `closed-fixed`, checked first): **TOK-05 authored — the unit of repair moves from the CLAIM to the PREDICATE.** The platform developer's PR scopes the fold as a 5-done / 2-pending program (**storage and messenger are next, not yet done**), and rosetta PR #14 — fetched read-only as `origin/pr-14`, reconciled at **92 absorbed / 30 superseded / 5 standing / 0 refuted / ZERO new information, verdict DO NOT MERGE** — turns out to be valuable only as **negative space**: all seven live defects sit where the PR and our corpus *agree*, which is structurally invisible to any method that diffs two documents. So: **adjudicate against platform artifacts, never against another doc.** The residual re-reads as **119 sites over 3 predicates** — the 81 drift sites (*three services are live-local husks*), 17 files / 30 occurrences (*a `graphql` profile exists*), 21 moved citations (*this line number names this construct*) — each with a legal set derivable from an artifact we already parse. A reading names *instances*; only a derivation names the *predicate*, which is why ten readings at 43–48% recall had a fixed point. **Step 0 re-derived all 12 denominators and corrected one, making the headline defect worse:** `docker-compose.yml` opens with `include: [common.yml]`, so there are **10** services not 8, and the two it adds declare no `profiles:` key — therefore `make up PROFILE=graphql` does **not** "start nothing", it starts **three** (`postgresql redis sentinel`). Postgres answers, `docker ps` is non-empty, and the *application* is absent: the silent no-op dressed as a partially-working stack. **Five decisions:** `D-M257x-59-1` predicate-scoping (union-set left as a **pending user decision**, subsumed but not answered; clause 5 **not** re-cut — still zero-or-nothing); `-2` a **new sibling guard** rather than widening `platform_alignment_guard.py` (whose `FENCE_KIND="standalone"` and repos_yml-derived roots are load-bearing), 6 derived assertions run **both ways**; `-3` **§7 rule 4 gains a citation-safety half** — iter-58 passed every schema dimension and still moved **22 of 23** citations at a **4.5%** catch rate, because rule 4's dimensions are about *removal* and this class is caused by *addition*; `-4` an 8th map state **`mid-fold`**, valid only with a **two-sided** citation — storage measured as compose-`storage-legacy` + `STORAGE_RPC_ADDR` set **nowhere** while app `v1.366.0` reads it at 3 sites and **hard-requires** it in 2 tools, recorded nowhere in the corpus; `-5` **fence first, then citations, then the map state, then read** — with recall at 43–48% a reading over an unfenced corpus samples a pool refilling faster than repair drains it (**net −72**). `stack-core` verified at its **1F/610** baseline before authoring — see iter-59/progress.md

- iter-60 (tik, **the first predicate fence**): **TOK-05 mechanised — `platform_predicate_guard.py` shipped
  GREEN, and the `graphql`-profile predicate class is closed by derivation rather than by 56 edits.** A **new
  sibling** of `platform_alignment_guard.py` (`D-M257x-59-2` — that guard's `FENCE_KIND="standalone"` scoping and
  its `repos_yml_path`-derived roots are what make assertion F trustworthy, so it was extended in *code*, not in
  *subject*). Six assertions, each run **both directions**, over `repos.yml` + `docker-compose.yml` (**`include:`
  resolved**) + `Makefile`. All six denominators re-derived at platform `0dab54d` **and cross-checked against
  `docker compose --profile X config --services`**: 10 services · floor **3** · 8 legal profiles · 6 repos ·
  default `core`→**5** · 4 `*_RPC_ADDR` all `http://backend:8083` · 1 migrating repo · the storage mid-fold split.
  Watched **RED at 37 → 21 → 2 → GREEN 0**, with **28 tests** including **three INVERSION mutants** (a guard
  hardcoding *"graphql is dead, core is alive"* survives all three; a derived one cannot — §8 rule 5's corollary
  that removal mutants do not catch inversion) and a **no-op positive control that survives** re-running with
  identical reach. **Of the 37 first-draft findings, 16 were the guard's own**, and every one was removed by
  **replacing the rule with one derived from the artifact's structure** rather than by excepting a name
  (`--profile` needs a compose driver in-window — `buildbench --profile billion` is a *host* profile; a token must
  match compose's own token shape; a repo-count's modifier slot is case-sensitive; a name adjacent to `.`/`/` is an
  identifier). One rule was **dropped rather than fixed** and the drop **reported**: per-name attribution of
  `migrations:` flags is not decidable in English prose, so G5 names **22 of 26** lines UNREACHED. **The briefed
  taxonomy was incomplete — there are THREE classes, not two:** works (5 tokens) · **silent no-op, rc=0** (5
  tokens, starting only the floor) · **hard-fail, rc=1** (`frontend`/`studio-desk`/`messenger`, `depends on
  undefined service "backend"`) — and `make up PROFILE=frontend` + `PROFILE=studio-desk` are **documented commands
  that exit 1**. Of `CLAUDE.md`'s six profile rows, **one** was accurate. **11 files repaired, +241/−74, to zero
  fence findings** — including two claims **fortified against repair**: `cms.md`'s *"the husk still starts …
  messenger is still pointed at it"* (M809 has landed) and, worse, `platform-alignment.md` §5 carrying iter-22's
  correct-at-`2adcf71` refutation as standing guidance, so **the protocol doc forbade the repair now required**.
  The protocol doc's own *"starts **zero** containers"* corrected to **three** — worse than zero, because a
  partially-working stack sends the reader debugging the application. **Six protocol-doc rules written** (the ones
  TOK-05 deliberately withheld): §5 **29** *a reading names instances; only a derivation names a predicate*, §5
  **30** *grade on does-it-still-SELECT*, §5 **31** *a refutation expires exactly like the claim it refuted*, §5
  **32** *re-derive the hand-off's numbers, including the orchestrator's*, §6 *the platform's CONFIG is its
  documentation of record*, and §7 rule 4's **citation-safety half** (`D-M257x-59-3`). `storage.md` carries the
  first **two-sided mid-fold record**, G6-fenced. **Two inherited denominators corrected**: "17 files / 30
  occurrences" measures **26 docs / 56 lines**; `academyImport/main.go:235` is the *return*, the `Getenv` is `:231`.
  rext tagged `fast-build-m257x-iter-60`, **verified on origin**; pin advanced. Gate **4 of 5** — see
  iter-60/progress.md

- iter-61 (tik, **the fence could not see its own class**): **iter-60's GREEN over-reported — the profile
  predicate's LARGER HALF was standing in a form the guard had no construct for.** iter-60 fenced the
  *command* (`PROFILE=x`, `--profile x`) and *table-first-cell* forms and read GREEN; a re-survey at the same
  ref found the **noun phrase** — *"the default `graphql` profile"* — at **34 raw sites across 17 files**.
  **A fence whose reach is narrower than its class over-reports its own GREEN**, and invisibly, because the
  fence is what you would check with. It also **vindicated the briefing**: *"17 files"* is exactly what the
  noun-phrase construct yields, so `D-M257x-60-7`'s "undercount" row is **withdrawn** (its two `main.go`
  line-number corrections stand). Two new constructs, both still constructs — the noun phrase (a backticked
  token adjacent to the literal word `profile`) and the **table row** (`| \`CMS_RPC_ADDR\` | \`http://cms:8091\` |`,
  the same binding with no `=`; `messenger.md` held two stale values in that shape the whole time the fence
  was green). *"GraphQL" the API stays unrepresentable.* **13 of 35 raw hits were the guard's own**, removed
  by two discriminators derived from the corpus's own writing: **adjacent negation** — and note the recursion,
  *"there is no `cms` profile"* is precisely how iter-60 wrote its repairs, so an undiscriminating widening
  **reads this milestone's own fixes as fresh defects** — and the **ref-pin** exemption G2/G4/G5 already had
  and G1 did not, plus the bare backticked sha that carries most of the historical narrative. **RED at 2
  findings / 22 sites / 12 files, enumerated in `iter-61/evidence/residual.md` with its regeneration command
  and routed WHOLE** (§5 rule 19's scope-edge corollary — a subset repair pools the residual at its own
  boundary). **RED and correct beats GREEN and narrow.** 35 tests (was 28), incl. a regression that iter-60's
  GREEN fixture survives the widening. The routed citation class was **re-measured before being routed on**:
  iter-58's "21 of 22" is, at app `v1.366.0`, **5 of 16 distinct `app/main.go:N` citations still landing on
  their claimed construct**. rext tagged `fast-build-m257x-iter-61`, **verified on origin**; pin advanced.
  Gate **4 of 5**, unchanged — see iter-61/progress.md

- iter-62 (tik, the repair half): **iter-61's enumerated 22-site / 12-file prose class repaired WHOLE to
  guard-GREEN** — and **12 of those sites carried a SECOND false predicate** the sweep caught: the husk
  containers. `dependency_map.md:31` asserted *"They are NOT gone from compose or `repos.yml`"* (they are gone
  from **both**); `service_taxonomy.md:97` answered **"YES — container still starts"** (it does not);
  `run_guide.md:88` listed **10** services incl. the Cosmo router for `make up` (it starts **five**); and
  `run_guide.md:203` documented `make up PROFILE=studio-desk` as starting Studio-Desk when it **exits 1**.
  **`D-M257x-62-1`: repair by predicate, but do not edit only the predicate** — the predicate is what finds
  the sites, the sentence is what you fix. **`D-M257x-62-2`: the ref-pin hole has a SECOND instance and is
  promoted to the next fence build** — `service_taxonomy.md:55-67`'s entire Services table is headed
  *"@ platform `2adcf71`"*, so every row listing Jobsimulation/CMS/Roadrunner as starting containers is
  exempt by the guard's own rule; `messenger.md:107-110`'s two stale RPC values were the first. Fix shape:
  a pin exempts only when its ref **is** the guard's ref. `markdown_structure_guard` clean at 112 files.
  Gate **4 of 5**, unchanged — see iter-62/progress.md

- iter-63 (tik, **a ref-pin is a DATE, not an exemption**): the routed citation class **re-measured
  before repair** — and **both** inherited figures were subset readings (`D-M257x-63-3`): iter-58's
  *"21 of 22"* counted **raw sites** of the string `main.go:N`, pooling three different files under one
  name; iter-61's *"5 of 16 distinct"* had the right unit but could not see the bare `:N`
  **continuation** construct. Derived over both constructs at `app` `b948604` v1.366.0: **104 sites / 86
  distinct citations across 22 files land in `app`, of which 18 are the app MAINLINE** — 5 held, **13
  moved and repaired**, each with its ref written beside it; the 68-citation remainder routed WHOLE.
  Adjudicating against artifacts (never against another doc) refuted four service docs at once
  (`D-M257x-63-5`): `docker-compose.yml:171-183` @ `0dab54d` sets **all four** `*_RPC_ADDR` to
  `http://backend:8083` under compose's own *"cms + jobsimulation are folded into app"* comment —
  **M809 has landed**, while `messenger.md`/`cms.md`/`jobsimulation.md`/`dependency_map.md`/
  `backend.md:195` asserted the two-of-four split as current, three of them emphatically (*"current,
  not stale"*). `platform-migration-status.md:76` was right the whole time — **the fenced map was
  right and the prose was wrong.** Then `CHECK-M257x-iter60-stale-pin-exemption`, **answered**: it is
  **three mechanisms wearing one name** (`D-M257x-63-1`) and only the third is a policy question —
  the pin crossed a **row** boundary (`shared_libraries.md:41-42`) and a **cell** boundary
  (`service_taxonomy.md:98`, one row two clauses), both plain window bugs; the third is the pin cited
  **as evidence of currency** (*"(**current** … @ platform `2adcf71`)"*). **Not an expiry — age is not
  the variable, tense is**; not a mandatory two-sided citation. A pin's scope is the claim's own
  block (a table **cell**); a pin naming **the guard's own ref** exempts nothing; a block asserting
  currency cannot be pinned into silence. Blast radius measured first: 19 blocks pinned-and-current,
  **1** carrying a checked construct. And the fence still was not enough — `D-M257x-63-2`:
  `service_taxonomy.md`'s Services table names its profile column **fourth** and G1 required it
  **first**, so six rows were **unreachable, not exempt** (*fifth GREEN-is-a-reach-limit of this
  milestone*). Widened: **17 illegal sites / 2 files, zero guard-own findings** (iter-60 was 16/37 its
  own, iter-61 13/35) → repaired to **GREEN**, incl. the Services table rewritten from the artifact
  (ten services, default `core` selects **five**, storage moved out of the default selection) and
  `shared_libraries.md`'s colony split corrected **three-way → two-way**. **54 tests** (was 35) with a
  **4-mutant inversion battery, all caught**. Two protocol rules written: §5 **33** *a ref-pin is a
  date, not an exemption* and §5 **34** *line numbers move when YOU edit too* — this iter's own repairs
  moved **9 intra-corpus citations across 6 files** (two in root `CLAUDE.md`) and the anchor guard
  caught **1**. Suites: `stack-core` **664 / 1F** (the perishable iter-48 fixture — baseline),
  `dev-stack` **151 / OK** — but only when run ALONE: beside `stack-core` it reported **6**
  `test_dev_public_host` failures that all vanish solo, because `stack-core`'s m220 battery spawns
  nested `dev-stack` runs. And `test_test_collection_fence` went RED on this iter's own test file
  (new classes appended *after* the `__main__` guard) — the fence caught it. rext tagged
  `fast-build-m257x-iter-63`, **verified on origin**; pin advanced.
  Gate **4 of 5**, unchanged — see iter-63/progress.md

- iter-64 (tik, **the map's eighth state**): `DOC-M257x-iter59-storage-mid-fold` closed — routed forward
  **five times**, and `D-M257x-64-1` names why it kept slipping: **the measurement had nowhere legal to
  go.** iter-59 measured the storage split and wrote it into `storage.md`; the *map* — the fenced
  artifact, policed by `platform_alignment_guard.py` assertion C — had a **seven-token** vocabulary with
  no token for it, so `storage` kept reading `live-standalone` on both sides. Not from belief: inventing
  a token in a fenced field turns the guard RED. **The fence had eight things to say and seven words** —
  the same shape as `D-M257x-63-2` one layer up, *the instrument reports agreement it never tested*.
  **`mid-fold` added as the eighth token**, `storage` re-stated with both halves cited, `ALLOWED_STATES`
  widened 7 → 8, and the protocol's own *"the seven-token vocabulary above has no token for mid-fold"*
  sentence retired. Re-deriving the split from artifacts (never from `storage.md`) found iter-59's
  consumer count **short by one**: **three** `cmd/` readers, not two — `academyImport:231` and
  `academy-asset-upload:129` hard-require it, while **`cmd/import/main.go:50` builds a storage client
  against the empty string without complaint**. Watched RED: the guard fired on the first attempt for
  the *wrong* reason — the new cell is bolded and `_state_head` read `**mid-fold**` literally — fixed by
  **stripping markdown emphasis** (`D-M257x-64-2`), derived from the format, rather than by un-bolding
  the cell and leaving the next author the same trap. **Two inversion mutants, both taking the LIVE map
  RED as well as the fixtures.** 35 alignment-guard tests (was 30); `stack-core` **669 / 1F** (the
  perishable iter-48 fixture). Side: the map's freshness pin advanced `ef32d4c` → **`0dab54d`**, and
  `service_taxonomy.md`'s *"CMS, Jobsimulation and Roadrunner are NOT out of local orchestration"* block
  retired — that phase closed at `d11a403`. rext tagged `fast-build-m257x-iter-64`, **verified on
  origin**; pin advanced. Gate **4 of 5**, unchanged — see iter-64/progress.md

- iter-65 (tik, **a citation must name its subject**): two routed `CHECK-*` items closed, both the same
  class — *an anchor that RESOLVES and still does not name the claim*, which
  `anchor_construct_guard`'s own docstring calls the line the fence family does not cross.
  `CHECK-M257x-iter60-g6-citation-subject`: G6's two-sided record was closed by `if site in all_text`,
  a **whole-corpus substring test** any document mentioning `main.go:446` for any reason satisfied.
  **For a known token the judgement is decidable** — the site and the variable must now co-occur in
  the same **block** (the unit `_pin_window` established at iter-63); no claim is parsed, two tokens
  must share a paragraph. Live corpus **GREEN**, which is the right outcome and is **not** evidence
  the rule works — the six fixtures are (`D-M257x-65-1`). And writing one of them found a **second**
  defect the tree could not show: `test_a_site_named_with_no_variable_anywhere` came back green where
  RED was expected because **G6's universe excluded the consumer side** — a variable the platform
  configures nowhere, the corpus names nowhere, and `app` READS had **no row at all**, i.e. the
  most-undocumented case was invisible to the assertion whose job is to catch it (`D-M257x-65-2`;
  **second time this milestone a FIXTURE surfaced a reach hole the live corpus lacked** — iter-61 was
  the first). `CHECK-M257x-iter64-pms-87-subject`: `service_taxonomy.md`'s Directus retraction appealed
  to `platform-migration-status.md:87` as *"the corpus's own fenced source of truth"* — **that map has
  no Directus row at all** (it maps repos; Directus is external), so the anchor resolved, carried
  content, passed the anchor guard, and named `anthropos-studio-room`. Re-adjudicated against the
  platform (`git show a2a3ee6^:docker-compose.yml`, already cited in the same paragraph). **It had been
  faithfully re-pointed TWICE by iters 63–64 as its target moved — a re-point preserves intent, it
  cannot audit it** (`D-M257x-65-3`). 60 predicate tests (was 54); mutants: revert-to-substring **2
  RED**, drop-consumer-side **1 RED**. `stack-core` **675 / 1F** (the perishable iter-48 fixture). rext
  tagged `fast-build-m257x-iter-65`, **verified on origin**; pin advanced. Gate **4 of 5**, unchanged —
  see iter-65/progress.md

- iter-66 (tik, **root `CLAUDE.md`**, corpus-only): the two facts iters 63–64 measured, carried into the
  file every agent reads first — and **neither was reachable by any fence in the family**, which is the
  point. (1) `Storage` was listed under *"In the default local profile (`core`)"*; at `0dab54d` it
  declares `profiles: [storage-legacy]` and `core` starts **five** containers. A reader would expect a
  storage container on a stock `make up`, not get one, and not know why the calls fail. Replaced with the
  two-sided **`mid-fold`** note, pointing at the fenced map row rather than restating it. (2) The
  RPC-edge inventory said *"backend → sentinel/storage and messenger → backend"* and was wrong **in both
  directions at once**: messenger's edge is now **all four** addresses (`d11a403`, **M809 landed**),
  while **`backend → storage` is mid-fold, not live** — nothing sets `STORAGE_RPC_ADDR`, `storage` is not
  started, so the client is built against the empty string and fails at **call time, not boot**. All five
  corpus guards OK; rext untouched (no tag, pin stays `fast-build-m257x-iter-65`). New route
  **`FENCE-M257x-iter66-tier-membership`** — *"service X is in selection Y"* has a derivable legal set
  (`compose.select(default_profile)`) and nothing checks it; same shape as G1 but about **membership**
  rather than the token. Gate **4 of 5**, unchanged — see iter-66/progress.md

- iter-67 (tik, **G7 — the service list beside a profile**): consumes `FENCE-M257x-iter66-tier-membership`
  **one iteration after iter-66 opened it**, so the rule is proven by use. iter-66's defect (`storage`
  placed in the default selection) was unreachable by every existing assertion — G1 checks the token is
  legal and selects *something*, G3 checks the default's *count*, **nothing checked the LIST**. G7
  asserts the services named beside a profile equal `compose.beyond_floor(tok)`, **both directions**
  (MISSING / NOT STARTED), with the services column found by its **header** (like the profile column at
  iter-63) and its cells read **by shape**; a prose cell yields no tokens and is **UNREACHED, never an
  empty claim**. Live: **22 membership rows, 12 checked, 10 UNREACHED — GREEN**, which is the right
  outcome for a fence built right after a hand-repair and is **not** the evidence; the fixtures and
  mutants are. Two findings on the way: **the corpus's most important row was invisible to G1 too** —
  `` | `core` *(default — `PROFILE ?= core`)* | `` defeated a `(default)`-only strip because the
  qualifier is emphasised and carries its own backticks (fixed by stripping a trailing parenthetical
  before the default mark; rows checked 10 → 12, profile sites 91 → 94, `D-M257x-67-2`); and **the
  mutation battery caught a weak TEST of mine, not a weak rule** — a `membership_rows >= 1` assertion
  that passes on the base corpus alone and therefore survived its own mutant, re-written to assert the
  **delta** (`D-M257x-67-3`). 67 predicate tests (was 60); `stack-core` **682 / 1F** (the perishable
  iter-48 fixture) — a second failure, `test_m220_mutation_battery`, was **my own contention** (a prior
  suite still running; solo re-run **10/10 OK**), the same mistake iter-63 recorded, repeated in the
  same session. rext tagged `fast-build-m257x-iter-67`, **verified on origin**; pin advanced.
  Gate **4 of 5**, unchanged — see iter-67/progress.md
- iter-68 (tik, closed-fixed-partial): the routed citation class re-derived — **64, not 68**: iter-63's
  own enumerator now reads **123 sites / 96 distinct / 22 files**, because iters 63–67 wrote ten
  net-new citations of their own and six of them died with the row they documented (a corpus repair
  **enlarges** its own citation class — the sibling of §5 rule 34). §7 rule 4b measured **before**
  spending the iter: `app` origin/main advanced **56 commits to `9d00a313` v1.367.0 at 10:56Z that
  morning**, and **25 of the 42 citations that hold at the demo's pinned `b948604` break at origin
  HEAD — 60% in one working day**. That redirected the repair to origin HEAD, the ref the gate names,
  and exposed the precondition nobody had seen: **three guards — G6's consumer side, assertion F,
  every anchor into a clone — were reading whatever the clone had checked out and none of them said
  so.** Same corpus: **GREEN at origin HEAD, 4-findings RED at the pinned ref.** All three now
  resolve *and* read at a named ref and print it (`auto` → origin/main → HEAD; `worktree` by name; a
  caller-named ref that does not resolve is UNMEASURED). **Release 09.00 landed while we were
  counting: storage and messenger are BOTH folded** — prod `service_desired_count = 0` on both,
  storage served in-process, and `app` **taking over messenger's own Redis consumer group** rather
  than merging its handlers. The `mid-fold` token built at iter-64 now has **no instance**; it stays,
  and the map says so. Seven of eight `* **Profile**:` bullets were wrong in a construct **no fence
  reaches** (seventh reach limit of this milestone). `CHECK-M257x-iter63-quoting-a-retired-token`
  closed — **the second window bug of this milestone wearing a policy's name**, one line too narrow
  for a denial that wrapped. 6 mutants across 3 guards all caught, one after its first test **survived**
  and had to be re-written against the call site. B2 (the non-fold citation remainder) routes forward.
  Gate **4 of 5**, unchanged — see iter-68/progress.md
- iter-69 (tik, closed-fixed): **B2 was 5 defects, not 64.** Graded at the ref the gate names, the
  126 (file × citation) pairs split **62 identical · 59 drifted-but-REF-PINNED · 2 unpinned-and-moved
  · 3 mis-rooted** — and every pinned mainline citation resolves *exactly at its own pin*
  (`backend.md:39`'s seven all do). A pass that "repaired" the 59 would have moved 59 correct claims
  onto a ref that moves again next week. The mechanical screen was watched RED first (pin-blind
  2→61, comparison-blind 2→0) with a **no-op control that survived** — and then **reading** found
  what the screen structurally could not: **`shared_libraries.md:70` contributes ZERO citations to
  every count this class has ever reported**, because it names `` `app/main.go` `` in backticks one
  clause away with no line number for the enumerator to latch onto. **Five of its six handler line
  numbers were wrong at every ref**, contradicting the pinned-and-correct `backend.md:39`. The
  hidden class is **23 citations across 14 lines** (derived; 17 routed forward). Also repaired:
  Judge0 at `wiring.go:118`→**`:123`** (unpinned, false at the gate's ref) and `ROADRUNNER_RPC_ADDR`
  *"(`docker-compose.yml:118`)"* — **zero occurrences** in the compose at `0dab54d`; rule **32**'s
  own worked example pinned (**rule 33 applied to rule 32's neighbour**); and `app/studio/**`
  documented as the **in-image** path, in no `app` commit at any ref. The repair **widened the
  fence**: `anchor_construct_guard`'s origin/main reach **43 → 49**. **G8 shipped** for the
  `* **Profile**:` bullet — the seventh reach limit of this milestone, where iter-68 found 7 of 8
  wrong while every fence read GREEN — as **G7 inverted** (a SERVICE, from the doc's own file stem,
  checked against the profiles beside it; three shapes, anything else UNREACHED): **8/8 reached, 0
  unreached, GREEN**, with **5 source mutants + 1 artifact inversion caught and a no-op control
  surviving**, every fixture copied verbatim from the corpus. Class re-measured at close: **141
  sites / 109 distinct** (from 135/105 — it grew by this iter's own corrections), **unpinned-moved
  2 → 0**. And the re-measurement crashed the enumerator: it cannot sort a path carrying both `X:N`
  and `X:N-M` at one start line — **every count this class has ever reported was correct by luck**.
  `stack-core` **753 / 1F** (the perishable iter-48 fixture, by IDENTITY) · `stack-injection` 332 OK
  · `dev-stack` 151 OK solo. rext NOT tagged (offline guard code only); pin stays
  `fast-build-m257x-iter-67`. Gate **4 of 5**, unchanged — see iter-69/progress.md
- iter-70 (tik, closed-fixed): **the class iter-69 routed as "23 citations across 14 lines" — and
  called a prerequisite for the graded read — is 4 citations across 3 lines.** Twelve of the
  seventeen are **PORT NUMBERS**: `repos.yml` has 31 lines and the citation is `:5050`;
  `docker-compose.yml` has 271 and it is `:8082`; `cors.go` has 110 and they are `:8000`/`:9000`.
  One mechanical discriminator settles it — *does the antecedent file even have that many lines?* —
  **and the discriminator's own defect is reported beside its result**: it sized the antecedent by
  `max()` over every same-basename file on disk, so `labsapi/client.go` resolved to a 23,437-line
  namesake and *"default `:7070`"* came back line-plausible. **§5 rule 32 against my own hand-off,
  off by a factor of six, one iteration later.** Three of the four hold — `org_membership.go:172-188`
  is exactly the fail-closed `Policy()` at `9d00a313`, and `ai-readiness.md:45`'s bare `:33` binds to
  the yaml already named on its line (the detector was wrong, not the corpus). The one repair is a
  dangling *"consistent with `:36` and `:261-266` **above**"* in `studio-room.md:388` where
  `services/ai.py` is cited **nowhere earlier in that document** — corpus-internal, no ref needed.
  Whether those are the right *lines* stays **UNMEASURED**: studio-room is not cloned on this box and
  the CI-pulled in-image copy is in no ref, so a verdict off it would not be a measurement —
  **recorded, not patched on speculation**. And `FENCE-M257x-iter69-citation-antecedent` is
  **falsified as designed**: the corpus uses the port sense 3× more often, so the rule would be a
  false-positive generator (§4 Trap A); replaced by `FENCE-M257x-iter70-line-or-port`, which adds a
  *decidable* side-condition (`N ≤` the file's length at the adjudication ref) rather than a better
  guess. **iter-69's claim that this class blocked the graded read is retracted.** Five corpus
  guards OK; no code touched, so iter-69's suite runs stand. Gate **4 of 5**, unchanged — see
  iter-70/progress.md
- iter-71 (tik, closed-fixed): **`FENCE-M257x-iter68-citation-resolution` lands — iter-68 gave three
  guards a ref and gave each of them ONE, and the corpus does not have one ref.** Measured at open:
  of **125** resolvable citations, **31 sit in a block naming exactly one ref that resolves in their
  own clone** — a quarter of the class, every one read at `origin/main` regardless of what its
  sentence said. `backend.md:39` pins to `app` `b948604` v1.366.0; iter-69 re-pointed
  `shared_libraries.md:79` to `9d00a313` v1.367.0; **no single knob can be right about both**, and
  iter-68 had already measured the consequence — the same corpus is GREEN at origin HEAD and
  **4-findings RED** at the pinned build ref. `anchor_construct_guard` now resolves **per citation**:
  exactly one resolvable sha in the block → read at it; **more than one → fall back AND COUNT IT as
  `ambiguous`** (a block naming two refs is *contrasting* them — `platform-alignment.md` rule 32 does
  exactly that — and guessing would be §4 Trap A, while a silent fallback would hide 12 citations
  inside `default`); none → the `CITE_REF` ladder unchanged. A sha that does not resolve in that
  citation's **own** clone is not a pin; a pin at which the **file does not exist** is **UNMEASURED,
  not clean** (§5 rule 7); and **`CITE_REF=worktree` still overrides every block pin** — verified
  live, still 1 finding under `worktree`, GREEN by default. Reach is now printed beside provenance:
  `ref chosen by default x57, block-pinned x31, no-clone x30, ambiguous x12`. **A mutant SURVIVED**:
  narrowing the window from the block to `lines[i]` passed the entire suite, because every fixture I
  wrote put the pin on the citation's own line and **the corpus does not write that way** — the
  **third one-line-window bug of this milestone and the first inside a test**; two corpus-shaped
  fixtures added (pin two lines above, in-block → found; pin one blank line away → not found, rule
  33) and both window mutants now die. **7 mutants caught, no-op control survived.** And the
  battery's own baseline caught a real breakage: one new reach counter grew `run()`'s return from a
  6- to a **7-positional tuple** and broke **four existing callers** — visible only because the
  baseline line read `4 errors` where `OK` was expected, which is why a battery is run against a
  *stated* baseline and not a remembered one. `tests/test_iter45_mechanical_fences.py` **46 → 55**.
  Five corpus guards OK. rext NOT tagged (offline guard code only); pin stays
  `fast-build-m257x-iter-67`. Gate **4 of 5**, unchanged — see iter-71/progress.md
- iter-72 (tik, closed-fixed): **`FIX-M257x-iter58-mainline-shift` closes with a DERIVED verdict —
  66 distinct `main.go` citations, 37 graded at a ref their own block names, 29 in ambiguous blocks,
  and ZERO out of range or absent.** Its *"21 of 22 outstanding"* joins iter-69's *"64"* and
  iter-70's *"23"* as a carried number that did not survive re-derivation — **four in four**.
  Probing the guard's reach inside that class then found the **eighth and largest reach limit of
  this milestone**, proven mechanically in both halves: `_QUALIFIED` requires a `/` or a `.md`, so a
  bare `` `main.go:1187` `` **never reaches `resolve()` at all** — **142 distinct citations are
  outside the guard's reach entirely**, led by **41 `docker-compose.yml:N`** and **32
  `up-injected.sh:N`**, the two most-cited artifacts in the ops corpus. And the resolver would miss
  them anyway: its service-doc rule maps `backend.md` → `stack-demo/backend/`, which does not exist,
  because the compose SERVICE is `backend` while the REPO is `app`. The fix is **designed and routed
  with both proofs** (`FENCE-M257x-iter72-bare-citation-reach`) rather than half-landed — widen the
  regex, **derive the doc-stem→clone map from compose's `build.context`** rather than listing it,
  and keep `AMBIGUOUS_BASENAMES`' discipline, since `main.go` exists in **seven** clones and a
  tree-wide basename search is the over-match the guard's own docstring records as *"134 findings,
  essentially all of them ports."* Also recorded: a **self-inflicted measurement bug** — the first
  derivation printed *"block-pinned: 0"* against the guard's own `block-pinned x31` because the
  script called `live.pop()` then tested `len(live) == 1` on the mutated set. **Two instruments
  disagreeing is a finding, and the one that agrees with nothing is usually the new one.** Five
  corpus guards OK; zero code touched in either repo, so iter-71's suite runs stand. Gate **4 of 5**,
  unchanged — see iter-72/progress.md
- iter-73 (tik, closed-fixed): **`FENCE-M257x-iter72-bare-citation-reach` lands one iteration after
  iter-72 opened it, and with it the eighth reach limit closes.** A bare `<name>.<ext>:N` now
  reaches the resolver, and the service→repo edge is **DERIVED from `docker-compose.yml`'s own
  `build.context`** rather than assumed to be the doc's filename — including the
  `${APP_BUILD_CONTEXT:-../app}` form (a bare-`../repo` parser drops **the one service that
  matters**) and excluding a non-local git-URL context (there is no clone to resolve into). The
  third regex alternative is a **CODE-SUFFIX allow-list, never a `\w+\.\w+` wildcard** — this
  guard's own docstring records what the wildcard costs: *"134 findings, essentially all of them
  ports."* **Reach 124 → 177 anchors.** Dry-run first with the guard untouched (136 newly
  resolvable / 92 still unresolvable / **12** findings) so "land and repair in one iter" was a
  decision at the start rather than a discovery halfway. Live it went **RED with 6**, all
  `docker-compose.yml:N` past the end of a **271-line** file — router-deletion debris — and **two
  were false CLAIMS, not stale numbers**: *"the container still starts locally"* said of
  **roadrunner**, which has **no compose service at all** at `0dab54d`, in `architecture_overview.md`
  **and** in `roadrunner.md`, unnoticed by every KB-fidelity pass in this milestone because nothing
  could resolve a bare `docker-compose.yml:281`. All six repaired. **Two battery defects, both
  mine:** a **mutant SURVIVED** — the wildcard `[a-z]*` passed because the anti-port fixtures
  (`:8082`, `:5050`, `localhost:3000`) have **no dot before the port** and could never have
  discriminated, while the corpus is full of dotted hosts one colon from being read as citations;
  and **the no-op control was not a no-op** (it split a `re.compile` call), which is worth a rule of
  its own — **if the control fails, the battery has not run, and every "caught" beside it is
  unearned**. Final: 5 mutants caught, control survived. `stack-core` **769 / 1F** (the perishable
  iter-48 fixture, by IDENTITY); `CITE_REF=worktree` still discriminates. rext NOT tagged (offline
  guard code only); pin stays `fast-build-m257x-iter-67`. Gate **4 of 5**, unchanged — see
  iter-73/progress.md
- iter-74 (tik, closed-fixed): **the ambiguous class's 12 → 39 is settled by PARTITION, not by
  judgement — and reading it anyway found the fence defect underneath.** Splitting the bucket along
  the one dimension iter-73's widening moved gives `bare-code 27 · md 0 · path 12`: the pre-existing
  partition is **still 12, to the citation**, so **all** the growth is a class that could not be
  counted at all the day before. The corpus did not move by one site. Then the class itself:
  **21 of the 39 sat in one document**, which turned out to be a property of the *guard* —
  `_block_of` walked to the nearest **blank line**, and a markdown table has none between its rows,
  so a citation inside a table took the **entire table** as its window and every sha named in any
  row chose the ref for the citations in **every** row. That contradicts **§5 rule 33**, derived in
  iter-63 and implemented in the sibling guard ever since: *a markdown CELL in a table, a wrapped
  sentence in prose.* **Two guards, two definitions of "block", both green, and the wrong one wrong
  silently.** Measured instance: `external_services.md`'s **AWS Bedrock** and **Mistral** rows carry
  no pin and were being read at `b948604` — a ref named by the **Anthropic Direct** row. Dry-run
  first (iter-73 lesson 1): predicted `39 → 24 / 0 findings`, reproduced exactly. Then **the first
  draft of the fix contained the defect it removes** — the row test, copied verbatim from the
  sibling, anchors at `|`, so `storage.md`'s **blockquoted** fold table (`> | side | … |`) failed it
  and fell back into the prose branch; admitted on a count, not a document (**75 quoted rows across
  14 files** vs 2197 plain). Final **ambiguous 39 → 20 · block-pinned 45 · default 82 · 177 resolved
  · 0 findings**: nothing in the corpus was wrong, and **19 citations stopped being adjudicated
  against a file their own document never named.** The residual 20 was then adjudicated rather than
  assumed — classified at **every** ref its own window names, **19 agree, 1 diverges**, and the 1 is
  a cell that names in words the ref it asserts, so **no repair and no fitted rule** (§4 Trap A).
  **5 mutants caught, no-op control SURVIVED**, every fixture derived from a document the corpus
  actually writes. §5 gains **rule 35** (*a count that grows in the same pass that reach grows is
  not a regression until you split it by the reach dimension* — rule 16's mirror — plus *when a rule
  is already derived, check every implementation of it*). Two hand-off corrections recorded: the
  routed *"92 unresolvable"* **count** reproduces (91 distinct / 103 sites) but its **head list came
  from a different instrument** (`gen.py` is cited 11× and **never** in `file:N` form — every one is
  a RANGE), and **88 of the 239 "unresolvable" sites are URLs, not citations**. P3 at close caught
  the adjudication ref moving *during* the iteration — a routine `git fetch` took `app`'s
  `origin/main` `9d00a313` → `7177374` (v1.367.0+4), silently re-pointing 82 default-adjudicated
  citations with no diff anywhere (§5 rule 26); both ref-aware guards re-run at the new ref, both
  GREEN, and the citation delta is **48 held / 1 moved / 0 dead** — the 1 being a claim **true at
  the ref it names**, i.e. iter-69's *a pin is a date* reappearing inside a script written the same
  day. Five corpus guards OK. rext NOT tagged (offline guard code only); pin stays
  `fast-build-m257x-iter-67`. Gate **4 of 5**, unchanged — see iter-74/progress.md
- iter-75 (tik, closed-fixed): **the routed "92 unrepaired citations" adjudicate to 0 defects — the
  FOURTH consecutive routed count in this milestone to collapse when someone finally derived it**
  (64 → 5, 23 → 1, 21 → 0, 92 → 0). Fated three ways against `git ls-files` over **7,265 tracked
  basenames across 13 clones** — the repository's own answer, not a directory walk: **UNIQUE 77
  sites / 26 basenames · MULTI 26 / 19 · ABSENT 0.** **Not one citation names a file that does not
  exist.** It is the **ninth reach limit**: iter-73 taught a bare `<name>.<ext>:N` to *reach* the
  resolver and did not teach the resolver to *find* it, because every route `resolve()` owns is
  positional and an ops doc citing `` `up-injected.sh:1487` `` supplies no position at all. Landed
  as a last-resort **unique-basename** route — `git ls-files` only, **bare citations only** (a
  path-qualified citation has already said where the file lives; resolving it by basename would
  override the document with a guess about its directory), and **exactly one path or nothing**.
  Reach **177 → 254, 0 findings**; `resolved via bare-unique-basename x77` printed beside the reach
  line, because `resolve()` returns a bare `Path` and a route that fires silently is a reach claim
  nobody can audit. **The 0 was earned rather than assumed** — 0 was the *surprising* answer one
  iteration after a comparable widening went RED with 6, so the same code path was fed `:99999`
  (out-of-range), a blank line (on-blank-line) and two valid lines first. The **26 that stay
  unresolved are the rule working**: `main.go` is **57** tracked files, `main.tf` 10, and
  `studioManager.go` is `app/internal/cms/studio/` **plus** `cms/internal/studio/` — the merged copy
  and the standalone husk, exactly the pair a directory guess gets wrong and exactly the fold the
  map exists to document. **My own instrument was wrong twice in one iteration, in two directions:**
  `rosetta-extensions` is cloned twice under this tree **with the same directory name** (the
  per-stack copy pinned at a tag, and the authoring copy), so the adjudication script collapsed them
  and the dry-run script split them — one reporting a file in *"2 places"* while printing the same
  path twice, the other silently finding 31 of 77. Fixed by deciding **which copy is the witness**
  (the authoring one — `resolve()`'s pre-existing rext fallback already preferred it) rather than by
  de-duplicating harder; pinned as a test and as mutant M4. §5 gains **rule 36** (*a universe built
  over CLONES must say which copy is the witness* + *a derivation that returns the convenient answer
  earns a positive control* — §5 rule 2 applied to measurements, not just searches). **6 mutants
  caught, no-op control SURVIVED.** `tests/test_iter45_mechanical_fences.py` **68 → 74**;
  `stack-core` **781 / 1F** (the perishable iter-48 fixture, baseline matched by IDENTITY, +6 =
  exactly this iter's tests). Five corpus guards OK; `CITE_REF=worktree` still discriminates.
  **Both known-bad citation classes are now closed**, which was the stated precondition for the
  graded read. rext NOT tagged (offline guard code only); pin stays `fast-build-m257x-iter-67`.
  Gate **4 of 5**, unchanged — see iter-75/progress.md
- iter-76 (tik, closed-fixed): **the graded READ was taken, and clause 5 is not met by a wide
  margin — 77 blockers in reading #13, 75 in #14, against a pre-registered ceiling of 10.** 14 blind
  seats, identical partition, instrument held fixed at iter-41's and STORED as
  `instrument/briefing-iter76-AS-RUN.md` (its only deltas: the ground-truth shas, and an explicit
  ref-selection rule — recorded, not applied silently). Taken only once iter-74 and iter-75 had
  closed both known-bad citation classes, because *a reading taken over a known-unrepaired class
  measures the instrument, not the corpus*. **The headline is not the count — it is that the count
  coexists with a green board.** Five corpus guards were OK throughout **and were right to be**:
  `platform_predicate_guard` prints on every single run that **G5 reaches 1 migration claim of 24
  and G2 reaches 3 repo-count claims**, so the corpus's false predicates live in the twenty-one it
  cannot see. *Every GREEN verdict is a statement about reach* — this iteration is that sentence
  with a measurement attached. Each dominant class was confirmed against the guard's **own derived
  denominators**, never against another document: **no compose service exists for cms /
  jobsimulation / roadrunner** while six documents say the container *"still starts locally"*;
  **`repos.yml` has 6 entries** (app · sentinel · storage · messenger · next-web-app · studio-desk)
  while the corpus says *"1 of 9"*; **`core` selects 5 containers** while the corpus says nine; and
  **`graphql` is not one of the 8 legal profile tokens** while `make up` is annotated with it in
  three documents. **Prediction 5 HOLDS — no seat booked a blocker inside the iter-74 or iter-75
  classes**, so both closes stand. Two of five pre-registrations were falsified, one by an order of
  magnitude, and the falsifications are the iteration's most useful output. **The first adjudicated
  finding is a FALSE POSITIVE with a systematic mechanism**: the *"past the end of a 271-line file"*
  class sits in blocks pinning `2adcf71`, where `docker-compose.yml` is **387 lines**, so §5 rule 33
  settles those citations as correct — **~150 is an upper bound, not a work item**, and the seats
  hesitated correctly (both booked `medium`). The instrument lesson is mine, not the seats': **a
  rule stated in a briefing is not a rule enforced by an instrument** — the same ref-selection
  predicate is already mechanical inside `anchor_construct_guard.block_ref`, one directory away, and
  the read had no access to it. Repair **routed whole** per the iter's own pre-registered escalation
  (*">~15 union blockers → measure and route"*), as `FIX-M257x-iter76-read-union`, under three
  binding conditions: **adjudicate before repairing** (four routed counts in a row have collapsed on
  adjudication — 64→5, 23→1, 21→0, 92→0), **repair by PREDICATE not by claim** (`D-M257x-59-1`; six
  predicates cover most of the ~150 and every one has a legal set the guard already derives), and
  **not closed until the G5/G2 reach hole is closed**, since 1-of-24 is *why* these claims survived
  seventy-five iterations of a milestone whose entire subject is this fold. One denominator (8 vs 9
  vs 10 compose services) is left **explicitly unsettled** rather than resolved by assertion. Gate
  **4 of 5**, unchanged — see iter-76/progress.md
- iter-77 (tik, closed-fixed): **the reach hole is closed at its cause, and the cause was not the
  twenty-one.** The briefing's design question — *can free prose be fenced, or should the corpus be
  restated?* — is answered **neither** (`D-M257x-77-1`), on measurements: three candidate prose
  fences reach 71% precision / 60% recall / a rejected third, and restating would have reshaped the
  corpus to fit a broken instrument. The hypothesis in `overview.md` was **confirmed live by
  counterfactual**: the guard's repo vocabulary was derived from what *currently exists*, so a repo
  left the vocabulary at the same commit it left `repos.yml` — `setup_guide.md:486` enumerated
  *"app, cms, jobsimulation"*, the resolver dropped the removed names, compared `{'app'}=={'app'}`
  and **passed a false claim**. *G5's effective reach was 0 of 24, not 1.* Vocabulary is now
  HISTORICAL (14 ever · 6 now · 8 removed, from the artifact's own 9 commits, re-derived
  independently), and **G9** is net-new reach into a construct no assertion could see — 19
  `repos.yml:N` citations enumerated, each read **at the ref its own block names and only if that
  ref resolves in the platform clone**. `FIX-M257x-iter76-read-union` adjudicated **before** repair:
  4 candidate defects → **3** (`jobsimulation.md:12` is a measurement at `2adcf71`, where lines
  17-19 *are* its block) — the **fifth** routed count in this milestone to shrink on adjudication.
  All 3 repaired **by predicate**, plus `setup_guide.md:486`; platform `d11a403` is derived as the
  commit that deleted both entries. A widening of `_pin_exempts` was built, measured (26 blocks, 1
  verdict changed) and **NOT shipped** — read rather than class-matched, that verdict is a false
  positive; reverted and pinned by 4 tests. The guard then went **RED on this iter's own repair**
  and taught `D-M257x-77-4`: a narrowing discriminator checked *before* the verdict cost a true
  claim and dropped graded 4→2; checked *after*, it spends no recall. §5 gains rules **37–39**
  (historical vocabulary · discriminator order · *committing is not pushing*). rext `main` pushed —
  13 commits had existed on one disk. 5 corpus guards GREEN, 140 guard tests, section suites at
  baseline. Gate **4 of 5**, unchanged — see iter-77/progress.md
- iter-78 (tik, closed-fixed): **the milestone's one explicitly-unsettled denominator is settled —
  and 8 vs 9 vs 10 was never three opinions about one number.** Derived across every ref the corpus
  cites: **8** is `docker-compose.yml` alone, **10** the effective topology once `include:
  common.yml` adds the always-on floor, and **9 is a count of nothing** — the history runs
  12 → 12 → 11 → 8 → 8 and never passes through it. Three documents asserted the nine. Six sites
  repaired so each states *which set it counts* (the unstated qualifier was the real defect),
  including `platform_repo.md:59`, which still listed a `graphql` service deleted two folds earlier,
  and the staging pair counting **14** running containers where `--profile all` now selects **8**
  (`messenger` and `storage` dropped from `all` — two consumers on one Redis group, two writers on
  one bucket). `platform-alignment.md:194`'s **14** adjudicated **correct**: a dated measurement,
  and at `b56d731` the effective count was 12 + 2. **G10** fences the predicate — and it asserts
  the **PAIR, never a value**, because both 8 and 10 name real sets and the corpus contains one
  correct document of each, so a fence picking a side would go RED on a correct doc and repair it
  into a different truth. The obvious construct was **measured at 44% precision** (this corpus
  counts the floor, the application services, *"the last two subgraph services"* and plain narrative
  with the same words) and **replaced, not thresholded** (§4 Trap A): scoped to a declaration verb
  with a compose subject in the same **block** it reaches 4 sites at **100%**. Number-words are in
  the pattern (all three false sites write *"nine"*; a digits-only rule reads none of them) and the
  block window recovered 2 of 4 that a line window missed — **the third window bug of this
  milestone**. iter-77's ref rule paid immediately: `external_services.md:296` claims *"platform
  `0dab54d`'s compose declares nine services"* while its leftmost pin is `b948604`, an **app** sha —
  **second confirmed instance** of a foreign sha laundering a platform claim.
  `CHECK-M257x-iter76-compose-service-count` **CLOSED**. 149 guard tests, 5 corpus guards GREEN.
  Gate **4 of 5**, unchanged — see iter-78/progress.md
- iter-79 (tik, closed-fixed): **`CHECK-M257x-iter77-cross-repo-pin` sized and CLOSED — 145 → 3, the
  sixth routed count in this milestone to collapse on derivation** (64→5, 23→1, 21→0, 92→0, 4→3,
  145→3). Of **390** pin-exempted blocks live, **145** are exempted by a sha that does not resolve in
  the platform repo (28 distinct shas) — but only **3** gate a platform-file assertion, and iter-77
  and iter-78 had already repaired 2. Both predictions registered in the overview held (*fewer than
  10*, *0 new findings*). **The finding is the duplication, not the count:** G9 hit this mechanism at
  iter-77 and G10 hit it independently at iter-78, and each fixed it *inside itself* — while G2/G4/G5,
  which assert claims about `repos.yml` and `docker-compose.yml` and are in exactly the same class,
  still took any sha in the block as a date. **A rule discovered twice and implemented twice is a rule
  not implemented.** Now one helper (`ref_resolves_in` / `pin_dates_a_platform_claim`) that the
  platform-file assertions call — deliberately NOT a widening of `_pin_exempts`, because the other 142
  are legitimate `app`-repo citations about `app` files and punishing them would be Trap A with 142
  false positives. Reach up one claim, findings unchanged at **0**. **And the generalisation inverted
  on first contact:** called unconditionally it made every sha unresolvable on a non-git platform dir,
  so no pin dated anything and the guard went RED across the whole corpus — iter-77's defect with the
  sign flipped, *blind there, hostile here*, both from reading "cannot answer" as an answer. **An
  existing silence-asserting test caught it within one run.** Fixed with `can_resolve_refs`, degrading
  to prior behaviour rather than to a verdict, and generalised: **every derived discriminator has
  three outcomes — yes, no, and cannot-tell.** Also wired `D-M257x-63-1`'s currency clause into G9/G10,
  which had silently opted out of all three by resolving refs instead of calling `_pin_exempts`.
  157 guard tests, 5 corpus guards GREEN, stack-core 1F/819. Gate **4 of 5**, unchanged —
  see iter-79/progress.md
- iter-80 (tik, measurement-shaped, + **harden pass 19**): **the seventh routed count is the first
  that did NOT collapse — and clause 5's instrument was never the fence.** The run opened on a
  framing question: if *free prose is not fenceable* and *the corpus should not be restated* both
  stand, then 20 of 23 migration claims are permanently outside the fence family's reach, and clause
  5 is either measured by something else or the milestone needs a TOK. **It is measured by something
  else, and this milestone has conflated the two before.** Clause 3's instrument is
  `platform_predicate_guard` + `platform_alignment_guard` (**MET**, 5 corpus guards GREEN); clause
  5's is the **graded READ** — 14 blind seats over the 40 files of `corpus/services/**` +
  `corpus/architecture/**`, frozen at iter-41's briefing and stored as
  `instrument/briefing-iter76-AS-RUN.md`. G5's printed reach (**1 enumerated + 22 free-prose
  UNREACHED + 1 ref-pinned of 24**) is **not clause 5's denominator and never was**: the fence stops
  a repaired claim from regressing, the READ measures the gate. Both prior conclusions stand and
  neither closes clause 5 off. **The measurement:** `FIX-M257x-iter76-read-union` adjudicated in
  full — **152 booked → 140 UPHELD / 12 REJECTED / 0 UNSETTLED, 92.1 %**, four parallel
  adjudicators each re-deriving from the clones rather than from any prior verdict. The six counts
  before it collapsed (**64→5, 23→1, 21→0, 92→0, 4→3, 145→3**), which is precisely why iter-76
  routed its 77/75 instead of repairing it and made *"adjudicate before repairing"* binding; the
  prior was *"it will collapse"* and **it did not**. That prior was well-earned and it now has a
  counter-example — the one that decides the gate. iter-76's own hedge (the *"past the end of a
  271-line file"* class making ~150 *"an upper bound, not a work item"*) is now measured: that
  systematic FP class is **4 of the 12 rejections, not most of the 152**. The other rejection
  mechanisms are new and worth carrying: a **pin in a subordinate clause**, **ref-relative truth read
  as self-contradiction** (two statements 58 lines apart, both true, at `9d00a313` and `b948604` —
  adjacency is not co-reference), and an **emphatic scalar scoped by its own argument**. **The path
  to clause 5 is now concrete and unblocked for the first time:** the 140 dedupe to **11 predicates**
  (all four adjudicators converged independently; P1 — *"the cms/jobsimulation/roadrunner containers
  still start"* — is ~47 findings across 6+ files), and all three binding conditions on the repair
  are discharged (adjudication done here; reach hole closed by iters 77–79; repair-by-predicate is
  what the 11 enable). **Repair the 11, then re-read. Not a TOK** — and the three-run instrument
  stretch was *on* the critical path, being the milestone's own booked precondition, while
  continuing it would not be. **Harden pass 19 RUN** (deferred three times, not a fourth), scope
  iter-69…79, closing `CHECK-M257x-iter79-three-valued-discriminators`: swept every
  subprocess-derived discriminator in `stack-core`; **most were already three-valued and said so** —
  iters 77–79 taught the module the rule and it mostly learned it — but **`_reads_at_ref` had not**,
  publishing two outcomes where `git grep` has three, so an unreadable object store reported *"the
  consumer side reads no `*_RPC_ADDR`"* with `app_consumer_side == "measured"`: the `|| echo 0`
  signature M257 opened on, one level down, inside the guard built to end it, **in the branch
  adjacent to the comment that states the rule** — the fifth *"author violated the rule while
  writing it."* Reachability **proven not argued**, and the upstream guard provably does not cover
  it: `rev-parse --verify` answers from the **commit** object (so it correctly catches a shallow
  clone) while `git grep` needs the **tree**, so a corrupt/unreadable tree object passes one and
  fails the other with rc 128 — built and reproduced. **3 mutants, 3 kills, 3 distinct signatures**,
  collected before running, including the rc-1 **control** that stops the fix from becoming *"treat
  every non-zero rc as unmeasurable"* and a provenance-silencing mutant (pass 18's *"a reporting path
  with no mutant is a docstring"*). **157 → 160 tests.** Also raised, and it is not a documentation
  defect: `storage.md:55,:154,:181` promise local private storage is sandboxed to `/tmp` while
  `docker-compose.yml:82` sets `STORAGE_S3_BUCKET=production-storage2024…` on **`backend`**, read
  straight into `NewManager` — **local private writes land in a production bucket**
  (`DEF-M257x-iter80-storage-prod-bucket`, high, escalated; not actioned, the platform half is a
  platform-repo question). 160 predicate tests, 5 corpus guards GREEN, stack-core's only non-green
  the known perishable iter-48 fixture; pin stays `fast-build-m257x-iter-67`. Gate **4 of 5**,
  unchanged — see iter-80/progress.md

- iter-82 (tik, measurement-shaped): **the re-read — clause 5 is NOT met, and the repair is
  incomplete.** The frozen instrument was re-run against the repaired corpus: 14 blind seats, two
  readings of one hand, all reports on disk under `iter-82/raw/`. **`N₁₅ = 29`, `N₁₆ = 30`;
  union 41 distinct anchors** — against **77 / 75 → 140 upheld** before the repair. Clause 5 is
  graded only by a reading that returns **zero**, so the **gate stays 4 of 5** and 140 → 41 is a
  *booked* union against an *unadjudicated* one, not 99 fixes. **The instrument was proven frozen,
  not asserted so:** briefing byte-identical (sha256 `3858ec53…`, one commit ever touched it), same
  40 files, same 7×2 seats, all **12** ground-truth clone shas re-derived and matched — and the
  partition **method** was re-executed at `012edd2` where it **reproduced iter-76's hand exactly**
  (40 files / 9,544 lines; all six seat totals and file lists identical). iter-81's +410/−251 moved
  line counts, so the fixed method deals a different hand over the same file set (now 9,712 lines) —
  the consequence iter-76 recorded at #11/#12, kept rather than engineered away. **Five of six
  pre-registered predictions graded, two falsified.** P5 — *"the 11 repaired predicates contribute
  zero"* — is **FALSIFIED and is the headline, as pre-registered**:
  `graphql-wundergraph.md:13` still says the `graphql` *"profile name survives in compose and is now
  simply the default profile"*, while at `0dab54d` the token appears in **no** `profiles:` key (the
  eight are core/backend/all/storage-legacy/customerio-sync/messenger/studio-desk/frontend) and the
  default is `core` — **both halves false**, inside **P4**, booked by **both** readings
  independently. P6 is also falsified, and in the useful direction: the held carve-out
  `storage.md:55,:154,:181` (`DEF-M257x-iter80-storage-prod-bucket`, escalated) was booked by
  **neither** reading — **it accounts for zero of the 41** and cannot explain any part of a non-zero
  N. **Recall remains the dominant term** — the two readings share only **15** of 41 anchors (≈51–55
  %), the same <60 % that has held across every paired measurement here, which is why a zero was
  never the honest expectation: repairing the union of two readings cannot repair what neither saw.
  Largest class in the 41 is the **stale cross-repo line anchor** (corpus → rext / `app`), led by
  `ai-readiness.md`'s rext anchors drifted a consistent **+4**, booked by two seats in both readings
  — which **measures `CHECK-M257x-iter77-cross-repo-pin`**, previously flagged unmeasured. One
  finding adjudicated in-run because three seats raised it and it decides the v9.0 fold's central
  fact: the `service_desired_count = 1` reading is a **FALSE POSITIVE** — `storage 63bffc8:…:38` = 0
  and `messenger a0ec933:…:29` = 0, exact at the refs the corpus names, the clones merely being
  older; **third occurrence of `CHECK-M257x-iter76-seat-ref-discipline`.** Two side-settlements:
  the **`stack-core` count is 822**, agreed by two independent methods (the runner's own
  `Ran 822 tests` from inside `stack-core/tests/`, and an AST enumeration of test methods — 822
  across 32 files), with the 1 known perishable iter-48 fixture the only non-green; and the
  `repos.yml :17-19` item is **narration, not documentation** — the corpus is clean at every
  `repos.yml:NN` citation and the three surviving `:17-19` mentions are each pinned to the older
  `2adcf71` where they are true. **Found while closing: iter-81 left NO iteration record at all** —
  empty `raw/`, no overview/progress/decisions, zero mentions in this file; the 11 predicates exist
  as a written list only inside its commit message. **No repair landed in this iter, by
  instruction** — see iter-82/progress.md
- iter-83 (tik, shape `tooling`): **why a discharged predicate had a surviving member — measured, and
  it is general.** iter-81 reported eleven predicates discharged; graded against **its own input
  ledger** (`iter-76/raw/`, 152 booked), it reached **109 of 147 gradeable findings = 74.1 %**, and
  **35 of the 38 misses sit in files it opened and edited**. So **H1 (partition gap) is REFUTED** —
  only 3 misses are in unopened files and all 3 are outside the read's own 40-file set — and **H2
  (estimated `~` membership) is REFUTED** too: the misses land on **exact**-count predicates
  (`external_services.md` ×5 = P8's 9; `storage.md` ×3 = P9's 3) as readily as on `~` ones. The
  decisive case is mechanical and is pinned as a regression test: in **one file**,
  `graphql-wundergraph.md`, the repair **rewrote `:177`** — a finding the adjudication had
  **REJECTED** — and **left `:13`**, booked as **B1** by **both** readings and **UPHELD**. **The
  discharge criterion was *"I have swept this file for this predicate"*** — not *"no member
  survives"*, not even *"every booked member is fixed"*. **Therefore all eleven verdicts are
  UNPROVEN** and are re-derived as membership questions at iter-84, not trusted. **Two deeper layers
  found while measuring:** `repair_leak_guard.py` — the one fence whose question is *"did this commit
  FINISH?"* — **exits 1 on `328ece5`** naming 3 sites (`CLAUDE.md:285`, a **live P4 member in a
  runnable command block**; `platform-alignment.md:1249`; `messenger.md:122`), all three still
  standing at HEAD, and it is **absent from the six guards the commit message lists**; and it is
  absent **structurally**, because it declares `FENCE_KIND = "standalone"` while the DERIVED registry
  selects only `postcondition`-kind fences — **10 of the 14 guards then standing were standalone** (4
  `postcondition`), i.e. the class you
  must remember, which is §2's hand-maintained tuple inside the machinery built to end it.
  **The arithmetic inconsistency is settled first, and it was a category error, not a slip:** 29/30
  count blocker **blocks**, 41 counted distinct **anchors**. Re-derived (extractor positive-controlled
  by reproducing iter-82's own per-seat table and both totals): **D₁₅ 28 · D₁₆ 29 · union 43 ·
  intersection 14**, and 28+29−14=43 ✓ — iter-82's own printed intersection list has **14** items
  while its prose said 15. **Recall follows:** Chapman `N̂` = **57**, **per-pass ≈ 49–51 %**, union
  ≈ 75 % — a *lower* bound on population and an *upper* bound on recall, because the two readings share
  a briefing/file-set/partition/model and correlated blind spots inflate the overlap. **Stated
  plainly, and it is the honest close-claim: a paired zero bounds the residual at roughly R ≤ 2 at
  95 %, not at none.** Clause 5 is **NOT** re-cut. **Shipped `FENCE-M257x-iter83-repair-reach`**
  (`repair_reach_guard.py`) — the third question, keyed on the **input ledger** where the other two are
  keyed on the diff and are blind by construction — **watched RED on a real answer key** (`iter-76/raw/`
  × `328ece5`), 16 behaviour tests + a **5-mutant battery, 5 kills, ≥3 distinct signatures**, no-op
  control survives; the load-bearing mutant `file-level-reach-accepted` is a mechanical statement of
  iter-81's actual criterion. **`FIX-M257x-iter82-iter81-has-no-record` DISCHARGED** — iter-81's record
  recovered from `git log` + the diff under an explicit **non-contemporaneous** banner, every field
  tagged `[git]`/`[msg]`/`[iter-83]`/`UNRECOVERABLE` (2 fields left unrecoverable rather than inferred),
  **re-grading iter-81 `closed-fixed` → `closed-fixed-partial`**. **No corpus claim repaired** —
  adjudication-before-repair is binding. `stack-core` 822 → **843**. **And rule 40 was tested on its author within the hour and FAILED** (`D-M257x-83-9`): the *"9 of 14 standalone"* figure was a **hand count** and is wrong — it is **10 of 14** (4 `postcondition`), caught by **the pre-commit hook's own output on the very commit that shipped rule 40**. Now DERIVED in one `ast` pass and always stated with its denominator and moment. *Derive, else fence, else declare* applies to the numbers a milestone states about **itself** — a hand-counted scalar in a plan doc is a **P11** waiting to happen, and P11 is one of the eleven. Corrected by follow-up commit, never by amend, so the evidence survives. Gate **4 of 5**, unchanged — see iter-83/progress.md
- iter-84 (tik): **the adjudication — 40 of 43 UPHELD (93.0 %), with a PER-ANCHOR ledger — and the
  eleven discharges re-derived by membership.** Four parallel adjudicators over disjoint packets, each
  re-deriving from the clones. 31 blocker · 9 minor · **0 unsettled**; pre-registered floor ≥ 70 %
  **HOLDS**, and the falsification condition (< 50 %, "the post-repair signal is mostly noise") did not
  fire — **the instrument did not regress across iter-81's repair** (iter-80 measured 92.1 % on the
  pre-repair union). `FIX-M257x-iter83-adjudication-has-no-per-anchor-ledger` **discharged in the same
  act**: the reach fence now has an `upheld` denominator (**40**), not only a `booked` one. Work list by
  predicate: **Q1** stale cross-repo anchor 13 · **Q2** present-tense claim about a **deleted** fact 7
  (re-anchoring is NOT the repair) · **Q3** wrong scalar/set 8 · **Q4** wrong predicate no line-checker
  could catch 7 · **Q5** 1. **The eleven discharges: 5 stand (P2 P3 P5 P6 P10), 4 refuted (P4 ≥17 live
  members in the published tree + 7 in rext source vs "~10 discharged"; P9 1; P11 ≥3), P1 recorded
  UNSETTLED** (clean inside the 40 files, ungraded candidates on the ops surface — *an unswept surface
  reported as clean is the defect this milestone exists to end*), P7/P8 folded into the adjudication.
  **A THIRD FINDING, CORRECTED AFTER PEER REVIEW (`D-M257x-84-6`) — it was first written as a
  "reach limit" on the instrument, and that was wrong:** clause 5's declared scope is
  `corpus/services/**` + `corpus/architecture/**` = **40 files** (0 of them non-`.md`), and **the
  instrument reads 40 of 40 — COMPLETE.** What is true is that **of P4's live members exactly ONE is
  inside those 40**; the other **≥16** are in `corpus/ops/**` (46 `.md`), `CLAUDE.md` and
  `.claude/skills/**`. That is a **SCOPE OBSERVATION and a corpus-quality finding — not an instrument
  defect**: clause 5 is narrower than the corpus (**90** `.md`, `git ls-files -- 'corpus/*.md'`) *by the
  clause's own wording*, and wanting more coverage would be a **re-cut of clause 5, which is NOT on the
  table**. The original framing measured the instrument against a scope the clause never claimed and
  reported a shortfall — an implicit re-cut, caught by a peer before it left the milestone; the
  undefined "112" is now stated with its command. **The real limiter on what a zero establishes is the
  ~50 % per-pass RECALL (iter-83), a within-scope property.** The finding still explains iter-81: the
  repair inherited the **read's** 40-file partition, so no seat owned the surfaces where most of P4
  lives — **§5 rule 19, the partition correct for reading is wrong for repairing.** Routed
  `CHECK-M257x-iter84-defects-outside-clause5-scope`.
  **All 3 rejections are one mechanism — `CHECK-M257x-iter76-seat-ref-discipline` at its 4th and 5th
  occurrences, and the declared escalation condition FIRED.** The diagnosis is that the rule is **stated
  wrong**: *"grade at the ref"* is silent on a sentence asserting **currency**, so seats apply it
  unevenly. Amended form — *grade at the ref the claim names UNLESS the sentence asserts currency* —
  which cleanly separates `graphql-wundergraph.md:13` (UPHELD) from `hiring.md:73` (REJECTED); plus the
  structural half, **the ground-truth table must carry each clone's `origin/main` sha beside its checkout
  sha**. Routed, **not written into §5 mid-run** (amending the frozen briefing's subject between readings
  is the one thing TOK-04 protects). **Adjudicate-before-repair earned itself again on a single anchor:**
  `ai_architecture.md:225` is **CORRECT** — the composited MP4 *is* in prod S3 — and the false statement
  is in `corpus/ops/demo/media-substrate-spec.md:33-35`, load-bearing for a **safety disposition**;
  repairing the booked anchor would have broken a true sentence and left the false one standing. **A LIVE
  REXT DEFECT, not a doc defect:** `dev-stack:186`/`:414` default `profile="graphql"`, so a bare
  `/dev-up N` runs `docker compose --profile graphql` — the token that **exits 0 and starts only the
  floor** (`FIX-M257x-iter84-dev-stack-default-profile`, high). The 3 leak-guard sites graded **2 real,
  1 benign** (67 % precision on a commit six guards passed clean) — `CLAUDE.md:285` and
  `platform-alignment.md:1305`, the latter being **the protocol doc contradicting the rule it teaches**.
  **§5 rule 32 fired twice in one run:** adjudicator B's own summary undercounted its own verdicts
  (9 vs 10). **No repair landed, by design.** Gate **4 of 5**, unchanged — see iter-84/progress.md
- iter-85 (tik): **the first repair since iter-81, and the first ever GRADED by a post-condition —
  reach 11/11 = 100 %** (iter-81: 74.1 %), measured by `repair_reach_guard` against this iter's own
  **declared input ledger** (`iter-85/ledger/`). Pre-registered *"0 unreached or it does not close
  `closed-fixed`"* — **0 unreached**. **Scope declared NARROW at open and that is the iter-83 lesson
  applied to itself:** Q2's 7 + the 2 confirmed leak sites + the rext defect; Q1/Q3/Q4/Q5 (29 upheld)
  routed to iter-86 with the ledger as their work list, because *a repair I cannot finish reproduces
  iter-81 exactly*. **Q2 — 7 present-tense claims about DELETED facts, restated or dropped, NOT ONE
  re-anchored** (§4 Trap A): `graphql-wundergraph.md:13` (**the run's centre** — the profile is gone,
  `PROFILE ?= core`, asking for the token exits 0) · `cms.md:8` (fourth, **not last** — v9.0 folded
  storage+messenger after it) · `backend.md:218` (**four of five** streams; nothing publishes to
  `skiller`) · `roadrunner.md:113` (**nothing** consumes the event) · `services/README.md:37` +
  `messenger.md:7` (**`MESSENGER_RPC_ADDR` exists in no repo and `git log -S` returns 0 commits ever**
  — repaired at **both** sites, §5 rule 19) · `architecture_overview.md:295`/`:311` (prod RPC list, and
  the retraction that scoped itself *"locally"* and thereby **affirmed** a dead prod edge) ·
  `alignment_testing.md:360` (**rc=3 since M219**, not rc=2 — it was telling readers to read a 2 as a
  missing Node module). Leak sites: `CLAUDE.md:285`'s runnable `make up # (graphql profile)` and
  `platform-alignment.md:1305`'s *"read by `main.go`"* → **read by nothing**. **🔴 The live rext defect
  FIXED — `dev-stack:186`/`:414` held the literal `profile="graphql"`, so a bare `/dev-up N` ran
  `docker compose --profile graphql` and brought up the floor with the application absent.** Fixed **by
  DERIVATION** (not by substituting `core`): both entry points now resolve via
  `platform_topology.default_profile()` and **import the same `FALLBACK_PROFILE` constant** rather than
  restating it. **The first cut was WRONG and the suite caught it** — a fatal derivation broke 13
  `test_dev_public_host` tests and took the M220 battery's **baseline** with it (*"the UNMUTATED subject
  fails its own suite"*); a stricter contract than the codebase's own is still a regression, and the
  correct shape already existed one file over. Proven three ways (real clone → `core`; synthetic dir →
  fallback **without dying**; no python3 → **dies loud**, because an *empty* profile selects only the
  floor). **`repair_leak_guard` — the guard iter-81 skipped — went RED on MY OWN repair**, naming §5
  rule 40 quoting the false sentence verbatim as its worked example; **waived with a written reason**
  (reported every run, never silent), not paraphrased. **A peer review caught a defect in iter-84's
  framing and it is corrected in place (`D-M257x-84-6`):** *"the instrument reads 40 of 112"* measured
  clause 5's instrument against a scope **the clause never claimed** and reported a shortfall — an
  implicit re-cut. **Clause 5's declared scope is 40 files and the instrument reads 40 of 40 —
  COMPLETE**; the ≥16 P4 members outside it are a **corpus-quality** finding, not an instrument defect,
  and every denominator now carries its command. `stack-core` **843**, only the known perishable
  fixture. Gate **4 of 5**, unchanged — see iter-85/progress.md
- iter-86 (tik): **the guard family ran as a family for the first time, and two of its sixteen members
  were RED.** `platform_predicate_guard` had **never been green since it was authored at iter-60** —
  25 iterations, across every record saying *"5 corpus guards GREEN"* / *"6 corpus guards exit 0"*,
  including the one that used that sentence to declare **clause 3 MET**. The reason is the sharper half:
  **G1 could not reach zero on a CORRECT corpus.** Its negation discriminator read only the text
  *before* the noun phrase, so the corpus's own correcting sentence — *"The `graphql` profile **is
  gone** too"* — read as a fresh claim that it lives. **A fence whose floor is above zero gets un-run,
  and it was.** iter-83's kind-filter diagnosis was one layer short: the census over iters 77–85 shows
  iters 77–79 captured a transcript of **one** guard beside an unenumerated count of 5; the count moved
  **5 → 6 with no record of what joined**; iters 83/84/85 assert *"6 exit 0"* with **no captured
  output**, and re-measurement with each iter's declared guard version and refs gives **rc=1 at all six
  of those points** — iter-85's own repair took the site count 2 → 3. **9 of 15 guards are covered by no
  green claim at all**, including `value_change_guard` (absent from the milestone outside iter-49) and
  **`derived_value_guard`, which is a `postcondition` guard** — so the filter cannot explain it. **§2's
  deleted 4-tuple returned as a human's remembered list, which a diff cannot catch.** Answer:
  **`guard_family.py`** — census DERIVED from disk, invocation map DECLARED and **reconciled both ways**
  (a guard with no entry exits 2 naming itself; a stale entry exits 2 too), refusing to read a guard's
  own *"Nothing was checked; this is not GREEN"* as a pass. **Both REDs adjudicated to opposite sides of
  the line:** 2 of the 3 G1 sites were a **CORPUS** defect — rule 40 **quoted the false sentence
  verbatim**, the protocol doc publishing a live-reading copy of the claim it kills (2nd occurrence) —
  repaired in the corpus, because `CLAUDE.md` already states the rule and rule 40 is the passage that
  teaches it; the 3rd was a **GUARD** defect, repaired in the guard, because contorting correct English
  to dodge a regex taxes every correct sentence after it. `_NEGATED_AFTER` denies **existence only** —
  never a bare *"is not"*, since *"is **not** started by default"* denies DEFAULTNESS — and **both false
  claims the guard was holding open survive it and stay RED**, which is the evidence it was fitted to
  English rather than to the answer key. **3 mutants, 3 kills, 3 signatures**; +7 tests (160 → **167**),
  +9 for the runner. **The seat-ref escalation settled by measurement, not preference** (`D-M257x-86-2`):
  the ground-truth sheet **already varies between the readings being compared**, so it cannot be the
  invariant that makes 140 → 43 comparable; and the class has contributed **zero to the graded count, 5
  times out of 5**, because adjudication was already filtering it. Adopted with an `origin/main` column
  (**6 of 14 clones are behind**; `app` by 60), and the **raw** series **declared discontinuous** at
  iter-86 while the adjudicated series is untouched — the third of the three honest options, taken out
  loud. **The repair: 30 enumerated predicate rows + the P4 sweep** (16 corpus/skills + 15 rext comment
  sites), seven disjoint file-partitioned packets, each re-deriving against the clones.
  **Adjudicate-before-repair earned itself FIVE more times** — B2's proposed range was the `return` and
  its brace (re-anchoring would have silently changed what the citation denotes); B11+B12 and B5+B6 were
  each **one** site double-counted; Q5's three supporting anchors landed on no S3 code (the conclusion
  held, every anchor had to be re-derived); and B23 was **two** switchers and **wrong when written**,
  not stale. **Two induced defects caught at COMMIT TIME by the new runner** — a repair re-anchored
  `app/main.go:212`, correct at the 60-behind checkout and a **closing brace** at origin/main (the
  seat-ref class arriving inside a *repair*, caught by an instrument instead of an adjudicator); and one
  packet's comment edits shifted another's re-derived anchor onto a blank line, which the postcondition
  ratchet refused. **Reach: raw 40/46 = 87.0 %** (iter-85: 11/11), 0 findings left unrepaired — **all
  six misses were defects in MY OWN LEDGER**: 4 already discharged by iter-85 (§5 rule 32, unapplied to
  my own input) and 2 anchor-drift inside iter-84's adjudication. **The fence graded the ledger, not the
  repair** — the first time it has said something about its own input, and the reason iter-85's 100 % is
  the weaker reading. Guard family **16/16 GREEN**. Gate **4 of 5**, unchanged — see iter-86/progress.md
- iter-87 (tik): **the platform moved a third time and the fence caught it unaided — the first entry in
  §1's fold table found by an instrument rather than by a breakage.** `838d907` (merged `0c91421`,
  2026-08-05) **deleted the `storage`, `messenger` and `customerio-sync` compose services outright** — not
  to a rollback profile, out of the file — and removed storage+messenger from `repos.yml`. Our clone was 2
  behind; advanced to origin HEAD in the iter that detected it (TOK-04 P3). Ground truth re-derived with a
  **control** (a detached worktree at the old ref): `platform_alignment_guard` assertion B gives **0**
  findings at `0dab54d` and **2** at `0c91421`, naming both departures in its own voice — the delta is the
  event, not accumulated debt. **The derived layer did better still: it needed no repair at all.** §2's
  time bomb was forecast to fire on *"the day they leave the clone set"* (13 write targets, 42P01, at
  once). This was that day, and the migration pairs (`app:public`) and CREATE SCHEMA set (`extensions
  sentinel public`) are **identical at both refs and identical correctly**, zero human action — third
  consecutive platform change absorbed unaided, and the bomb is retired rather than defused.
  **The hand-off's opening reading was refuted, and the refutation is the iter's best finding:** *"13
  GREEN · 0 RED"* is not reproducible — measured at the **identical** checkout after a `git fetch`, **10
  GREEN · 3 RED**. The citation guards read `origin/main` by iter-68's `CITE_REF=auto` ladder, so **the
  FETCH arms them, not the checkout**, and a citation fence pointed at an unfetched clone reads GREEN
  (**§5 rule 41**). That is what makes the clone-advance rule derivable rather than preferential:
  **fetch all; advance only what a derived set reads** (`D-M257x-87-1`) — and it **dissolves the deferral
  the hand-off anticipated**, because `app` at **93** commits behind (not iter-86's stale 60) carrying
  **65** citations was already being graded at origin HEAD and surfaced as **2** RED anchors, not a wave.
  **38 findings across SIX predicates repaired** by five disjoint packets, tree-wide by predicate and never
  by file: 3 dead profile tokens (28 sites), 3 no-such-service rows, the repo/service counts, **9
  unset-address claims** (the messenger block was the only thing setting all four `*_RPC_ADDR`, so
  `backend → sentinel` is now the ONLY cross-process edge), and 34 drifted citations → **13 GREEN · 0 RED ·
  3 not-run**. Three substantive state corrections, not re-anchoring: **`customerio-sync` moved
  `live-standalone` → `merged-into-app`** (a third service no plan doc had named); **`storage`'s prod ECS
  block is DELETED, not scaled to zero**, so the paired *"each `= 0`"* sentence was half false and was
  split; and **iter-86's own repair was among the falsified** — it had just written the `storage-legacy` /
  `messenger` profiles it is now false to claim, §5 rule 33 arriving as a live event rather than doctrine.
  **The re-scope trigger graded explicitly and NOT fired** (`D-M257x-87-2`), against a count the hand-off
  had wrong: `state.md` said *"occurrence 1 of 2"* while the milestone's own record shows it **already
  fired at iter-53** (`D-M257x-53-6`) and its prescribed remedy — a pinning-and-tracking policy — was built
  as TOK-04. Occurrence 3 is separated from occurrence 2 by **33 iters with no platform commit at all**, so
  *"two CONSECUTIVE"* is false on its own words; and the remedy performed on this very event. `state.md`
  repaired after **73 iters** of drift — an orchestration file is a claim like any other and nothing fences
  it. **Two instrument defects fixed:** rule 41 above, and the family runner reporting `lines[-1]` as a RED
  guard's headline, so the 21-finding alignment RED was summarised by a `gotenberg` citation nit while the
  two `[B departure]` lines were **invisible in the one view that speaks for the whole family** — repaired
  to *"N finding(s); first: …"*, derived from the producer's ordering (**§5 rule 42**, +5 tests).
  **Side-deliverable:** platform **M810 has already landed for `jobsimulation`** (`6092c6d2` destroyed the
  ECS service/task-def/ECR) while **`cms` has not moved** — the corpus asserted it as future work in ~14
  passages across 11 files. Landed as Fate 1 rather than routed, because the map had already been corrected
  and half-repairing is worse than not repairing. Carve-out held: `storage.md`'s three claims verbatim, the
  production-bucket escalation untouched. Gate **4 of 5**, unchanged — see iter-87/progress.md
- iter-88 (tik): **the class iter-87 named is real, and its second instance was a LIVE defect in the demo
  bring-up path.** iter-87 had only ever run `stack-core`; four rext sections read platform artifacts and
  had never been re-run at the advanced ref (§5 rule 8 — iter-04 found a test RED since iter-02 because
  only the new tests were run). Running them found: **the `$HOME/.aws/credentials` mitigation was keyed on
  `name == "jobsimulation"`.** That bind, unmounted on a fresh Linux box, is auto-created by Docker as an
  empty DIRECTORY, the AWS SDK opens it and fails `EISDIR`, and the container prints its whole cobra usage
  block and exits 1 — the symptom misread for a release cycle as a missing `serve` subcommand. `d11a403`
  deleted the service and **`838d907` moved the identical bind onto `backend`**, so **the hazard migrated
  to the stack's most important container and the mitigation did not follow.** Its tripwire said nothing:
  it looks the service up, misses, and calls `skipTest("jobsimulation not in the compose")` — a skip reads
  exactly like a pass. Its sibling **passed** while asserting *"exactly 1 `$HOME` bind (jobsimulation's AWS
  creds)"* — count right, claim false (§5 rule 17). **Repaired by derivation** (`D-M257x-88-1`):
  `services_with_only_home_binds()` over the **raw** compose text (after `docker compose config` expands
  them, the AWS bind and the stack's own postgres data dir are both just paths under `$HOME` — the intent
  survives only unexpanded), **ALL** volumes rather than ANY (so a service mixing a home bind with one it
  needs is excluded, not blanket-cleared — the caution the original comment asked for, now enforced), and
  **fail-closed on an empty derivation**. Returns exactly `{backend}` at `0c91421`. **Second instance
  (`D-M257x-88-2`): the demopatch anchor check has been SKIPPING since M254** — it restated its target path
  beside the manifest that declares it, M254 re-pointed the manifest into `internal/aireadiness/`, the copy
  did not follow, `isfile` failed, skip. So *"the app manifests were NEVER validated against a real clone"*
  — the thing that class exists to prevent — quietly became true again for four releases. Path now derived
  from the manifest; the skip split so **clone-present-but-path-absent FAILS**. Re-run: the anchor
  **resolves** (M254 was right, just unverified) and `stack-injection` went **2 skips → 0**. **Third
  (`D-M257x-88-3`): a test asserted the generator's source still contained `if name == "jobsimulation":`**
  — it would have failed anyone who tried to remove the dead literal. Re-pointed at the property, with an
  explicit `assertNotIn` on the old key. **§5 rule 43** generalises all three: key a mitigation on the
  PROPERTY that made the service special, never its name — the property outlives the fold, the name is what
  the platform keeps deleting; and **nothing in this family fences the test suites**, which is why they are
  where stale platform constants survive longest. The full sweep of the un-run sections came in at **41 files, 34 OK, 5 FAILED, 1 ERROR, 2 skips**, and
  **the residual is enumerated rather than summarised** — this iter's own finding is that unexercised
  checks go quiet. **Nothing in it is repaired here** (the tripwire fired on the third line): 3 are
  `pre_sha256` mismatches against live clones, which may be the **self-healing freshness gate working**
  rather than a defect — `demopatch-spec.md` says *the anchor is the contract; the whole-file sha is only
  a baseline* — so they are routed with an explicit **adjudicate-before-re-pinning** instruction; 1 is a
  **G5 self-revert failure** (`next-web-back-to-cockpit`), the most serious of the set; 4 look
  live-stack-dependent; 2 are unnamed skips. And the honest framing on all of them: `next-web-app` is 41
  commits behind and was **fetched** this run, so by **§5 rule 41** the *vantage* changed — "these have
  not been graded at origin HEAD before" is a different claim from "these broke", and establishing which
  is the first job of the iter that takes them. Gate **4 of 5**, unchanged — see iter-88/progress.md
- iter-89 (tik): **the four demo-stack failures iter-88 routed are ONE structural defect, and the joining
  experiment was `git status`.** `stack-demo/next-web-app` is dirty: the `next-web-back-to-cockpit`
  demo-patch is **left applied**, and a patched file cannot match a pristine `pre_sha256` — so the three
  sha mismatches were **downstream of the un-reverted patch**, not a separate class (§5 rule 28 again:
  three true facts do not make a cause). The revert refuses **correctly** (*"neither pre nor post — manual
  drift; refusing to guess"*, G2 doing its job); four shas explain why. The clone's **pristine** file
  (`48b6dd07`) is **not** the manifest's `pre` (`0c2c2ed2`) — `next-web-app` is 41 commits behind and the
  baseline predates that drift. So **apply** resolved its ANCHOR and self-healed onto the drifted base (by
  design — *"the anchor is the contract; the whole-file sha is only a baseline"*), producing a file whose
  whole-file sha is **necessarily** neither recorded value, and **revert**, which compares whole-file
  shas, refuses. **The asymmetry is the defect: APPLY is anchor-based and self-heals across base drift;
  REVERT is whole-file-sha-based and cannot — so on any clone whose base has drifted, which is the NORMAL
  state, the patch applies and will not come off, and the clone is left dirty.** That contradicts the
  mechanism's headline promise in its own spec's words (*"the clone is left git-clean"*): **G2 and G5 are
  in tension, structurally, and no test asserted their CONJUNCTION — each was verified alone.** Nothing
  repaired: the fix is a design choice on code that rewrites platform source inside a build (four options
  costed in `iter-89/decisions.md`; **(b) journal the observed pre-state at apply** is the reading offered,
  not taken), and the dirty clones are uncommitted state only the user may decide on. **Re-pinning the
  baseline was available and REFUSED** — it would have made one file match and hidden the asymmetry, which
  is exactly what iter-88's routing instruction and the spec both forbid. Clones left exactly as found so
  the defect reproduces without being re-created. `EXIT_REASON: user-blocker`. Gate **4 of 5**, unchanged —
  see iter-89/progress.md
- iter-90 (tik): **the demopatch asymmetry is repaired, and the CONJUNCTION is now what is tested.** The
  user's decision — **(b) journal the observed pre-state at apply; revert restores exactly it** — was
  re-derived and UPHELD, with one correction to its rationale: (b) does not *reconcile* G2 and G5, it
  **deletes the term they disagreed about** (revert's dependence on a recorded baseline that apply had
  already stopped consulting at M217). That distinction predicts why option (a) would have failed too. The
  reproduction was captured first, live, in two commands — `status` → `patched`, `revert` → *"neither pre
  nor post"* — on all three shipped next-web manifests, and the residue diff was independently re-derived as
  **exactly** the three demo-patches and no human work. Ordered as mandated: the **6-test conjunction
  battery** was written and shown **RED (4 of 5, negative control correctly green)** BEFORE the fix, and the
  dirty clones were not spent until it existed. The fix journals the observed pre-image in the **workspace
  root** (never in a clone), consumes it on revert, and removes the directory when it empties; **no journal
  ⇒ no guessing**, so an un-journalled drift still refuses — revert became *exact*, not *blind*. Proven
  **live on the real drifted clone**: apply→apply→LIFO revert→`git status` empty, and the chain's recomputed
  post `ebab9e7e…` is **exactly** the sha the failing tests reported when dirty, closing iter-89's diagnosis.
  The two dirty files were then cleaned via `--force-pristine`, and the limitation is **recorded, not
  hidden** — journalling cannot retroactively revert patches applied before it existed, so `apply` now WARNs
  when it meets that state. A first cut put the mutation control behind an env flag *inside* `demopatch`;
  it was removed — **a production code path that exists only for its own test is a backdoor** — and the
  control now rebuilds the tool with the journal blinded, test-side. **A double-revert test that failed
  against the fix was WITHDRAWN after measurement, not satisfied by bending the design**: `up-injected.sh:741`
  reverts once per `RETURN` trap, so it encoded a requirement that does not exist; it was replaced by the
  **chain** pair, which is on the shipped path and which no single-invocation test can see. Wider suite:
  **1054 tests, 6 failures, 0 of them mine** (3 need a live container; 3 are stale live-clone baselines
  across **two independent patch vehicles** — a systemic class, widened into one CHECK). The user's
  guard-fetch correction was applied: every `stack-demo` clone fetched, **nothing moved**, family reads
  **13 GREEN · 0 RED** before and after — so the reported 3 RED did **not** reproduce here, but **the class
  is real and worse**: `platform_alignment_guard` falls back to the **worktree** when a ref is missing and
  never says so (`auto` → GREEN/0-unresolvable vs `worktree` → **RED/8 findings**), and `unresolvable` is
  **printed but never graded**. Gate **4 of 5**, unchanged — see iter-90/progress.md
- iter-91 (tik): **the cannot-tell is now graded, and the user's question is answered with a measurement —
  a qualified NO.** *Should `guard_family.py` refuse against a stale clone?* **"Stale" is not locally
  decidable** (a clone fetched a minute ago can be behind; one fetched last week can be current), so
  answering it needs the network, and a fence that cannot run offline stops being run — the network check
  therefore exists and is **opt-in** (`--verify-remote`). **"Cannot see the objects it needs" IS locally
  decidable**, and that is what actually bit: `platform_alignment_guard` resolves at `origin/main` → `HEAD`
  → **silently the WORKTREE**, and its only positive control was `subject_checked == 0`, so **total**
  resolver failure tripped while **partial** blindness was folded into a verdict. Measured, the two
  references disagree — `auto` reads GREEN/0-unresolvable, the worktree fallback read **RED with 8 findings
  and 4 unresolvable UNGRADED**. That RED is the sharp case: it *looks* like diligence while 4 citations
  went unchecked. **A partial skip is worse than a total one, because it arrives with a verdict attached.**
  Both conditions now return **UNMEASURED (exit 2)**, with an `ALIGNMENT_ALLOW_UNMEASURED=1` hatch that
  RECORDS the gap. Fixed **at the point of use**, because only the guard knows which refs it needs — pushing
  it into the runner would rebuild §2's hand-maintained tuple in a new costume. The runner's own gap was
  smaller and worse: it printed `platform=<dir>` and **never a sha**, so every `13 GREEN` transcript in this
  milestone names a DIRECTORY, not a commit — which is exactly how a green reading gets quoted forward with
  no way to re-check it. The reference (corpus sha, platform sha, `origin/main`, in-sync, fetch age) is now
  printed on every run, and a platform-facing run against a clone with **no `origin/main`** exits 2 **before
  any guard runs**. **The 7-guard pair grid was enumerated** (Decision 1 item 3): of 21 pairs, **12 can
  interact, 5 were uncovered, 4 now have tests**. The two landed here were covered by **nothing** and are
  the safety-critical ones — **every** G1/G6 escape test drives `apply`, while **`revert` is the verb that
  writes AND runs `git checkout -- <path>`**; deleting its firewall call turns 3 subtests RED. Verification:
  `stack-core` **873 tests, 1 failure, PROVEN pre-existing** (the answer key re-run against a milestone copy
  with iter-90/91 removed gave the identical 2 hits / 101 claims) — and substantive: the claim it fires on
  asserts the `cms`/`jobsimulation` husks still run, which `838d907`/`0c91421` changed, so **the answer key
  is stale because the platform moved**. Folded into iter-92's M810 sweep. Guard family **15 GREEN · 0 RED ·
  0 could-not-check**, every clone fetched, `platform @ 0c91421df, in sync`. Gate **4 of 5**, unchanged —
  see iter-91/progress.md
- iter-92 (tik): **the M810 sweep was already done — and the one real defect was a FENCED claim restated
  UNFENCED, and stronger.** The brief's *"~14 passages across 11 files treat it as one future event"* was
  re-measured at HEAD before acting: **15 files / 40 occurrences, and the great majority already correct**
  (`corpus/README.md:16`, `CLAUDE.md:189`, `backend.md:36`, `cms.md:9`, `services/README.md:17`,
  `jobsimulation.md` 5 of 5, both fenced map rows). **A task description is a claim too**; re-running the
  sweep would have re-landed landed work, which is exactly what Step 0 exists to prevent. What the
  re-survey found instead is better. **(1)** The fenced map says of cms *"whether that rollback declaration
  still stands is **not something this map can see**"* — while `backend.md:36` and `CLAUDE.md:241` both
  asserted flatly that `module.cms_euwest1` **is still declared**. **Fencing a document does not fence its
  paraphrases**, and the unfenced copies had drifted UPWARD in confidence — a hedge that survives only where
  a guard reads it is worse than no hedge, because it implies the system checked. A real limit on the
  TOK-02/TOK-05 method, recorded as one. **(2)** cms **has** moved, opposite to what the corpus recorded:
  `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** the build-production workflow — *"the cms ECR
  repository is decommissioned (M810)"* — because it *"would try to push an image into a registry that no
  longer exists."* So the repo holds **two measured facts pointing opposite ways** (a module block still
  declared at `cms/terraform/main.tf:39`; a CI commit saying the registry is gone), while the destruction
  itself lands in `infrastructure/services.tf`, **in no clone set we have**. Now reported as UNMEASURABLE
  *with contrary evidence on both sides* — a better epistemic position than the corpus had, and the
  deferred-by-rule boundary respected rather than argued past. Ground truth re-measured per service:
  jobsimulation's module block is **deleted** (`6092c6d2`, service+task-def+ECR destroyed; the file survives
  only to own the LiveKit/Chime buckets), cms's is **declared at 0**. **And iter-91's new fence caught its
  own author within the hour**: the map edit first went in citing `` `main.tf:39` `` instead of
  `` `cms/terraform/main.tf:39` `` → `platform_alignment_guard CANNOT-CHECK rc=2`. **Before iter-91 that
  would have printed as unresolvable and exited 0 GREEN**, shipping a dead citation under a green fence —
  the first live catch of the third verdict, and a positive control nobody had to construct. **And the class
  was SIX documents deep:** `repair_leak_guard` went RED on the iter commit itself (the commit that had just
  claimed to fix it), naming 2 more sites; RED again on the first repair, naming 2 more; GREEN on the second.
  The single claim lived in `backend.md`, `CLAUDE.md`, `dependency_map.md`, `cms.md`, `external_services.md`
  and `storage.md`, and **only the fenced map carried the hedge** — the estimate when the class was named was
  "two files". **A claim's restatement count is not guessable**, and the guard's scope is the DIFF, so a
  repair commit surfaces sites the previous run could not see: **one green run is not the fixpoint, two
  consecutive greens are.** Gate **4 of 5**, unchanged — see iter-92/progress.md
- iter-93 (tik): **fence the HEDGE, not the sentence** — closing the class iter-92 discovered by committing
  it twice. iter-92 repaired six restatements of one claim by hand and **the repair leaked twice doing it**
  (`repair_leak_guard` RED on the iter commit, RED again on the first repair, GREEN on the second), which is
  the measured case for TOK-05 read literally: hand-repair does not hold; make the predicate
  unrepresentable. Shipped `stack-core/unreadable_repo_claim_guard.py` — **every corpus mention of a
  `module.*_euwest1` construct must carry an unmeasurable marker IN ITS OWN PARAGRAPH**, those modules being
  declared in `infrastructure/terraform/production/services.tf` and `infrastructure` never having been in a
  clone set. Registered in `guard_family.py` (**mandatory** — the family's reconciliation is bidirectional,
  so an unregistered guard on disk exits the whole family 2); **the family is now 17 members** and reads
  *"all 4 `module.*_euwest1` mention(s) are marked unmeasurable."* Four decisions, each tested not asserted:
  the marker is a **set of phrases, not a mandated token** (a fence requiring a magic string teaches people
  to type the string); the scope is the **paragraph**, with a wrapped blockquote counted as one, since a
  marker three screens away would launder a flat assertion while a fixed ±1 window would false-RED the real
  corpus (**both directions tested**); the guard **re-measures its own premise and RETIRES ITSELF** —
  if an `infrastructure` clone ever appears it prints *PREMISE LIFTED — go and MEASURE those declarations*,
  because a fence still demanding a hedge after the hedge became unnecessary is **pinning the current shape
  of our ignorance** (§8 rule 3 turned on a fence's own premise); and **no mentions ⇒ exit 2, never green**.
  That last rung proved itself inside the iter that wrote it: the live-corpus control was first written with
  a hardcoded `parents[3]` — `.agentspace/` in the authoring copy — so it **silently SKIPPED** and the suite
  printed `OK (skipped=1)`. **A check that skips reads exactly like a check that passes, including when it
  is the check on the guard that exists to say so** — ninth consecutive iteration in which the author of a
  rule broke it while writing the thing the rule governs. Gate **4 of 5**, unchanged — see iter-93/progress.md
- iter-94 (tik): **the guard family was reporting a green one of its own members had not earned.** iter-91's
  mutation run had shown the family against an empty tree reporting `2 GREEN · 2 RED · 9 could-not-check` —
  nine members correctly said *COULD NOT RUN — no corpus/*, two said GREEN. **This matters because the
  family's green is the evidence this milestone quotes**, so an unearned member contaminates every reading
  taken with it, including the clause-5 one still owed. Adjudicated **separately**, because a shared symptom
  is not a shared cause: `union_apply_guard` is **correct by design** (its subject IS rext's demopatch
  manifests; making it honour `--repo-root` would break it in the rext-only checkouts the family is consumed
  from per-stack) — **not changed, and that negative half is recorded, not silently skipped.**
  `story_org_count_guard` was a **real defect**: its `scan_roots()` returns the corpus **plus rext's own two
  directories**, and the guard *lives* in rext, so `if not roots: return 2` was **dead code that could never
  fire** — a run whose rosetta half was missing scanned only rext, found nothing to contradict, and printed
  *"and every doc agrees"* at `rc=0`. **"Every doc agrees" over ZERO corpus docs is vacuously true.** Fixed
  with a positive control on **the corpus half specifically** plus the cardinality: empty tree
  `rc=0 OK` → **`rc=2 CANNOT RUN — Nothing was checked; this is not GREEN`**; real tree → `OK — all **116**
  scanned doc(s) agree`. Both directions tested. **And the pattern is now the finding: three anti-vacuity
  defects in one session** (iter-91 graded total but not partial blindness; iter-93's own live-corpus test
  silently SKIPPED on a hardcoded `parents[3]`; iter-94's control could never fire) — three guards, three
  authors' assumptions, **one substitution: the control asked whether the guard had INPUTS when what
  mattered was whether it had reached its SUBJECT.** Generalised into the protocol doc §8. Gate **4 of 5**,
  unchanged — see iter-94/progress.md
- iter-95 (tik, MEASURING pass — **no repair inside it, deliberately**): **THE READING, TAKEN AT LAST — and at a platform ref no reading had ever used.** Preconditions **re-derived rather than inherited**: platform clone `== origin HEAD` by `ls-remote` (`0c91421`), guard family **14 GREEN · 0 RED** over 17 members (the run brief said 15; **the pass-22 ledger itself records 14**, so the brief carried a transcription slip and the instrument, the ledger and this re-derivation all agree), and the READ instrument **byte-identical** at sha256 `3858ec53…` with `git log --follow` showing **exactly one commit ever** (`012edd2`, iter-76). **The instrument was not touched**: the briefing went to the seats as a verbatim copy whose sha was re-checked *after* copying, with the stale ground-truth shas superseded by a **marked addendum** — only `platform` moved (`0dab54d → 0c91421`), every other checkout identical to iter-86's sheet. Partition re-executed by the fixed method: **40 files / 10,108 lines**, dealing a different hand than #13–#16 because the corpus grew, which is the method working rather than drifting. **14 blind seats, 55 blockers booked** (27 in #17, 28 in #18), **four independent adjudicators re-deriving from the clones: 51 upheld, 4 rejected — 92.7 %**, against iter-80's 92.1 % and iter-84's 93.0 %. **The graded number is `N` = 13 in-scope upheld BLOCKER anchors / 12 predicates** (`n₁=10 · n₂=7 · m=4` → Chapman **N̂ ≈ 16.6**, per-pass **60 %/42 %**, union **≈78 %**) — so **~4 more are estimated unfound**, and adjudicators named **≥8 unbooked propagation sites of already-upheld predicates**. **Clause 5 is NOT met; the gate stays 4 of 5.** All 4 rejections are the **ref-discipline class**, which has now run 9 occurrences across three readings and still contributes **zero** to a graded count. **The two findings that outrank the defect list:** (1) the environment's recursive `grep` is `ugrep --ignore-files`, so a `.gitignore` entry **silently hides tracked files** — `grep -rn mistralai app/studio/` returns 1 where `git grep` at the ref returns 2, verified independently; it produced a **false clearance in this reading** and very likely authored the false corpus claim originally, which makes **N a floor twice over** and yields the new rule *an absence is established only by `git grep` at a named ref*; and (2) `messenger.md:22`'s stated positive control measures **7**, not the 3 it claims — graded MINOR because the guarded claim is true, but a control that does not reproduce manufactures exactly the doubt it exists to remove. **Comparability stated rather than implied:** continuous in **instrument** and in **upheld rate**, a **declared re-baseline of the count** on two independent grounds (the gate-scoped basis, and the first-ever reading at `0c91421`); the raw series was already discontinuous at iter-86 and that stands. **`storage.md:58` is counted in `N`** — re-derivation found `:55` correct and `:156`/`:181` historically fenced, so `:58` is a **fourth anchor not among the three held lines**; `N = 12` is stated explicitly for the reading where the user extends the hold to the hazard class (**the user's item, the user's call**). Pre-registration graded **6 of 6 — recorded as a warning, not a win**: iter-76 and iter-53 each graded 2 of 5 and learned more, so the bands are to be tightened. Routed as **`FIX-M257x-iter95-read-union`** with iter-76's binding conditions inherited — **repair by PREDICATE, not by anchor**, and re-read after — see `iter-95/progress.md`
- iter-98 (tik): **the REPAIR of the iter-97 reading, by predicate + PARAPHRASE — 20 anchors / 21 predicates
  → 37 sites across 22 files, and 5 induced citation moves caught INSIDE the iter** (iter-96 caught 0 and
  shipped 2). The run brief asked whether to run twin expansion *ahead* of the read; answered **no, on
  measurement**: iter-96's string-twin pass already ran at ~0 escape and iter-97 measured what actually got
  through — **3 of 51, all paraphrases**, because `claim_twin_guard` matches quoted verbatim forms. So the
  effort went to the paraphrase axis, and it paid three times over. It found an unbooked third site of the
  messenger predicate; it found `frontend-tier.md` holding **both** readings of the demo-academy auth model
  nine lines apart; and it found that the booked anchor had named a **symptom** — the cockpit *does* still
  set `e2e_persona` at two live paths, and the real change is that the **bypass left the academy launch env**
  when the demo went Clerkenstein-wired, so the cookie is now **set and not honoured**. Rule 44's own recipe
  was corrected against its own worked example (it returned **2** where the prose publishes **22** — a file
  count labelled `hits=`, `-i` dropped, and its last line outside the loop with `$d` unbound), and its
  *"1,178 NULs"* is **1**; 1,178 is the file's line count. **The fence then caught what the repairer had
  not:** publishing the ledger turned `claim_twin_guard` RED on `seeding-spec.md:102`, a second site of the
  `s3-private` claim no sweep had reached — recorded plainly, because it is the whole argument for writing
  ledgers in the derived shape. Two guard interactions were resolved the hard way rather than the quiet one:
  a waiver **did not take** because the retraction said *"are both FALSE"* and `RETRACTION_MARKERS` knows
  `"is false"` — **the prose moved, not the marker list**, since widening it is "the direction that can
  hollow a fence out"; and the answer-key **green** fixture `claim_twin/green/17.md` was found to contain a
  sentence refuted 57 iters after its capture — repaired, with `test_01` (*all 18 known-bad must still fire*)
  untouched and still passing as the control that separates maintenance from tuning-to-green. Both are now
  §8 rules. **`DEF-M257x-iter80-storage-prod-bucket` was NOT resolved**: `safety.md:207` asserted the
  `s3-private` registry entry had been removed and `isolation.go:106` still carries it — the assertion is
  **withdrawn rather than made true**, because re-classing that store *is* the user's open question.
  **The discovery-pool question, answered with a measurement:** predicate WIDTH is collapsing — mean sites
  per predicate **3.64 → 1.76**, max width **11 → 4**, predicates of width ≥6 **4 → 0**, width-1 **29 % →
  52 %** — and width is a property of the corpus, not of the reading's recall. The classes split: draining
  and enumerable (platform-drift **7/13 → 1/20 → 0/17 predicates**, wrong-construct citations now a scan
  rather than a search, counts with a derivable source, currency pins) versus **open-ended** (scoping
  errors, model-changed-underneath from our *own* rext development, intra-document self-contradiction) —
  and **two of the three widest predicates this iter were open-ended**. **What it does not establish:** that
  `N` will fall. `N` counts what a ~68 %-recall instrument surfaces; two readings on different trees are not
  a series; and — the new point — **the shift from wide to narrow predicates should itself depress recall**,
  so a shrinking pool measured by a degrading instrument looks exactly like a growing one. Guard family
  **14 GREEN · 0 RED**; 21 new refuted forms fenced. Gate **4 of 5**, unchanged — no reading taken, and a
  measuring pass may not contain a repair — see iter-98/progress.md
- iter-99 (tik, the MEASURING pass — **no repair inside it**): **THE RE-READ RETURNS N = 28. Clause 5 is NOT
  met; the gate stays 4 of 5.** Taken at platform `0c91421` / corpus `e858fd4`, the tree iter-98's repair
  produced. Instrument **untouched and proven so** (sha256 `3858ec53…` re-checked AFTER copying, `git
  log --follow` on the file showing exactly one commit ever), the moved platform sha superseded by a
  **marked addendum** rather than an edit. Partition recomputed from current sizes — **40 files / 10,276
  lines**, 7 seats balanced 1431–1506 — dealing a different hand than iter-97 because the corpus grew, which
  is the method working. 14 blind seats booked **46**; four independent adjudicators re-deriving from the
  clones upheld **36 — 78.3 %**. `n₁=18 · n₂=16 · m=6` → union **28**, Chapman **N̂ ≈ 45.1**, per-pass recall
  **39.9 %/35.4 %**, union **62 %**, **≈17 estimated still unfound**. **The pre-registration was sealed in
  its own commit (`964b7a3`) BEFORE any seat reported** — that is what makes it a pre-registration — and
  graded **4 of 9**. **Three findings outrank the defect list. (1) Band #9 failed by ~7× and it indicts the
  INSTRUMENT, not the corpus:** `anchor_construct_guard` was **GREEN at the audited commit** — *"every
  **resolvable** anchor names a construct"* — while ≥7 upheld findings are citations resolving to the wrong
  construct, including a self-citation offered AS evidence that lands on a **blank line** and a
  `manager.go:485` that is a closing brace. The band was set at ≤1 precisely so an upheld member would mean
  a blind spot; the load-bearing word in the guard's green turned out to be *"resolvable."* **(2) Precision
  fell 93.1 % → 78.3 %, the first break in five readings** (92.1/93.0/92.7/93.1), with rejections rising
  4 → 10 while bookings *fell* 58 → 46. Three mechanisms fit and **this reading cannot separate them**: the
  residual genuinely hardened (iter-98 measured max predicate width collapsing 11 → 4); a **briefing gap** —
  **two independent seats made the identical wrong-tree error**, grading rext's authoring copy where the
  pinned per-stack clone was correct, and both bookings were rejected; and **adjudicator variance** —
  `hiring.md:80-82` was REJECTED by one panel and UPHELD by another, one disagreement in 46 but the first
  non-zero. Recorded, not resolved, because asserting the flattering explanation is the failure this
  milestone exists to prevent. **(3) Exactly 2 of the 28 were INDUCED by iter-98, both inside prose it
  rewrote** — `dependency_map.md:59` (the repair wrote *"6 occurrences across **4** files"*; both readings
  measured **3**) and `backend.md:33-34` (the repair removed `skiller` from the both-ends set and left it
  asserting exhaustiveness while **omitting `backend` itself**). Prediction #7 forecast [0,3] and got 2:
  **the mechanism model holds while the magnitude model fails**, the same split iter-97 recorded, and **0
  true recurrences of iter-98's 21 predicates** says the paraphrase axis worked. **What it does NOT
  establish: that `N` is rising.** 13 → 20 → 28 was measured at upheld rates of 92.7/93.1/**78.3 %**, so the
  series is no longer comparable on the axis iter-97 relied on; iter-98's discovery-pool §3 predicted recall
  would fall as the pool narrowed and it did (union 68 % → 62 %, both passes below 41 %, band #5 held). **A
  narrowing pool measured by a degrading instrument yields a rising `N`** — consistent with this data, and so
  is a genuinely growing residual. Routed as **`FIX-M257x-iter99-read-union`** plus three CHECKs, of which
  `CHECK-M257x-iter99-anchor-guard-blindspot` is the highest-value item in the reading — see
  `iter-99/progress.md`
- iter-100 (tik): the anchor fence was GREEN over 65% of its subject — blind spot closed + mutant-proven, 8 wrong-construct citations repaired (6 never named by any reading) — see iter-100/progress.md
- iter-101 (tik, the MEASURING pass — **no repair inside it**): **THE FIRST TRUE REPLICATE RETURNS N = 24, AND
  ITS ESTIMATOR BAND SAYS `N̂ = 45.1` WAS A FLOOR, NOT A CEILING.** Taken at platform `0c91421` (`== git
  ls-remote origin HEAD`, re-verified at the open) / corpus `8f04d3a`, over a pool iter-100 moved by **+2 net
  lines**, so the recomputed 7-seat partition came out **identical** — the instrument run twice over
  materially the same subject for the first time. **Ground truth re-derived, not inherited:** all 13 clone
  shas re-measured and matching the sealed sheet, instrument sha `3858ec53…` with `git log --follow` showing
  **exactly one commit ever** (`012edd2`) and the delivered copy proven a pure append (`diff` → `171a172`,
  nothing above the line edited). **Graded on 13 of 14 seats, disclosed rather than smoothed:** the fan-out
  died on a spend limit with `r24-D` unwritten; the 13 surviving reports were committed **verbatim,
  pre-adjudication** (`8b6d80f`) because untracked evidence is the defect class `evidence_visibility_guard`
  exists to catch. **`r24-D` was NOT re-run** — re-dealing it perturbs the replicate — so **reading #23 is a
  7-seat count (25 booked) and reading #24 a 6-seat count (11 booked), and `n₂` is never compared to `n₁`
  without that being said.** Four adjudicators re-deriving from the clones: **36 booked → 28 upheld / 8
  rejected**. **The upheld rate is reported TWICE as bound: 77.8 % raw, 80.0 % with the `wrong-tree`
  briefing-defect class separated** — against 92.1 / 93.0 / 92.7 / 93.1 and iter-99's 78.3 %. **That settles
  the question iter-99 could not at n=1: the ~15-point precision drop is STRUCTURAL, not adjudicator
  variance** — two readings, different adjudicators, same break. **`N` = 24 in-scope upheld anchors / 22
  predicates** (`n₁=20 · n₂=8 · m=4`), **a floor twice over** (13-seat union; a 14th seat can only add).
  **Clause 5 is NOT met.** **The headline is band #3, the first band on this milestone able to move the
  ESTIMATOR:** blind overlap with iter-99's published 28, matched on predicate, came out **6** against a band
  of **[14, 22]** — **FAILED LOW**, and the pre-registration pre-committed to what that means: the readings
  are **more independent than Chapman assumes**, so **`N̂ = 45.1` is a FLOOR**. Cross-reading Chapman over
  iter-99 × iter-101 (`28 · 24 · m=6`) gives **N̂ ≈ 102.6**, and heterogeneous catchability biases Chapman
  **downward**, so ~103 is itself a floor — **the residual is on the order of ~100, not ~45, and a zero
  reading is not near.** Stated with its assumptions (closed population: well supported; independence: now
  *better* supported; equal catchability: **dubious, direction of bias known**), never banked.
  ⚠ **RETIRED at iter-103, marked here at iter-104 (TOK-06).** The independence assumption this rests on was
  measured at **both extremes on one byte-identical instrument** — `m`/union **17 %** here, **61 %** at
  iter-103 — so independence is a property of *what is left to find*, not of the instrument. **`N̂ ≈ 102.6`
  is neither corroborated nor refuted; it is unestimable by this method.** The *conclusion* that a zero
  reading is not near survives on the floors alone (**≥ 24** at `8f04d3a`, **≥ 33** at `e6aed2e`); the point
  estimate does not, and must not be quoted forward.
  Pre-registration graded **5 of 9 held, 4 failed** — #1 split (`n₁=20` held; `n₂=8` failed low, and the
  6-seat caveat does **not** rescue it: normalized 9.3 < 10), #2 held (24), **#3 failed low (6)**, #4 held
  (77.8/80.0 %), #5 failed in an unpredicted direction (**83.3 % / 33.3 %** — the passes are wildly
  asymmetric, band #3's fact seen from the other side), #6 held at its bottom edge (**1** wrong-tree, vs
  iter-99's 4 — the unfixed briefing defect measured at n=2 readings and it is CHEAP), #7 held at **exactly 4**
  (iter-100's fence repair predicted ≤4 and the mechanical half really was ~half the class), #8 held (1–2 of
  24), #9 held (per-seat spread **4**, max 5 min 1 — the partition is not the variance source). **Three
  findings outrank the defect list:** (1) **iter-100 induced a defect inside the prose it rewrote** —
  `service_taxonomy.md:130-133` was exactly correct at `a229f8d^` until iter-100's own two-line parenthetical
  pushed the table down two rows and left the numbers unmoved, so the note now cites Chronos and Intelligence;
  found by **both** readings, and the ~2/cycle repair-induction rate held again while every magnitude band
  failed; (2) **seat D's entire in-scope yield is ONE predicate at three anchors in three files** — *"sentinel
  is the only cross-process edge"* — refuted by `GOTENBERG_URL=http://gotenberg:3200` at
  `docker-compose.yml:57` with `gotenberg` in the default `core` profile; the corpus already carries the
  correct qualified wording at `architecture_overview.md:321` and contradicts itself three more times, so it
  is **one repair propagated, not three fixes** — and it is **`CLAUDE.md`'s claim verbatim**; (3) the single
  wrong-tree rejection **inverts** iter-99's class, yielding a rule worth carrying: **the settling tree
  follows the claim's SUBJECT** — a local-stack claim is settled by the demo's build pin, a **production**
  claim by that repo's `origin/main`. **The ref-discipline class fired ZERO times** after 17 occurrences over
  five readings — three adjudicators reported it *structurally* absent. Routed as
  **`FIX-M257x-iter101-read-union`** (repair by PREDICATE, not by anchor; re-read after). **Gate re-graded
  honestly this iter: 2 of 5 proven, NOT 4 of 5** (`D-M257x-101-4`) — see `iter-101/adjudication.md`
- iter-102 (tik, the REPAIR pass — **no reading inside it**): **BOTH outstanding unions paid in one pass —
  `FIX-M257x-iter99-read-union` (28) + `FIX-M257x-iter101-read-union` (24) = 52 anchors — repaired BY
  PREDICATE across a TEN-seat file-disjoint fan-out**, wider than the habitual seven because *repair is the
  parallelizable half of this loop and the measuring pass is the half that must stay serial*. **76 anchor
  assignments → 98 sites found → 94 repaired**, plus **4 verified-correct-and-deliberately-not-rewritten**
  and **1 declined with evidence**. Guard family **14 GREEN · 0 RED** at open and at close;
  `claim_twin_guard`'s adjudicated-claim set went **134 → 264** because the ledgers are published in the
  shape `claim_ledger.py` derives from, so **completeness is fenced rather than claimed**.
  **Read the estimator correctly and do not let the record drift: the pool was probably always ~100. It is
  not growing — the estimator was wrong and iter-101's replicate fixed it.** The series
  16.7 → 29.4 → 45.2 → ~103 is four successive **corrections to an underestimate**, not four measurements of
  growth. ⚠ **HALF-CORRECTED at iter-103, marked here at iter-104 (TOK-06).** The second sentence survives —
  the series is still four corrections to an underestimate, not a growing pool. **The first does not:**
  *"the pool was probably always ~100"* quotes a point estimate from an estimator now **retired for this
  milestone**, on measured grounds (independence read 17 % then 61 % on one unchanged instrument).
  **Only floors survive** — ≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e` — and a floor is not a pool size. Track
  `N` and the predicate count directly; they need no assumption at all. **Five findings outrank the defect list. (1) The `app` clone move injected almost NOTHING, and
  that number is load-bearing for the ETA:** `origin/main` went `2035f9a4 → ad9f3c49` (5 commits) and
  `main.go` is **byte-identical** (1639 lines both refs, all five cited anchors identical). ⚠ **CORRECTED
  at iter-103 (`DEF-4`) — the two words that followed were false and they were MINE.** This sentence
  originally read *"as is `terraform/main.tf`"* and *"the entire residual is a LABEL"*. Re-measured
  `2035f9a4..ad9f3c49`: `terraform/main.tf` is **1 insertion / 1 deletion** (786 → 786; the changed line is
  the `atlas_sentinel_dev_url` `error_message` **prose string**, rewritten and lengthened), and
  `terraform/variables.tf` is **37 insertions / 12 deletions** (738 → **763**) — 49 lines that were never in
  the residual accounting at all. **The CONCLUSION survives and was re-verified** — `main.tf:181` is
  `service_desired_count = 1` at **both** refs, so the 5 commits still touched **no Go source and no cited
  terraform construct**, and the label really is where the defect lives. **The evidence sentence was stronger
  than the evidence, in the direction that made the conclusion cleaner** — this milestone's own class, in
  this milestone's own records. Read the residual as: **17 sites** call `2035f9a` *"origin/main"*
  (15 in `corpus/`, **2 in `CLAUDE.md`** that no reading is scoped to see), and that is the *cited* residual,
  not the whole diff. **A pin is a pin; only the labelled citations rot** — which is exactly why the label is the defect, and it confirms §5 rule 41's
  own corollary that *a large "behind" count is not a large repair.* **(2) `DEF-M257x-iter101-crosslane-fetch`,
  and it is MY defect:** I gave `stack-demo/**` exclusively to one lane while another lane's adjudicators
  were grading claims against the clones inside it. Measured: the fetch window was **11:18:16–11:20:51**,
  iter-101's adjudication commit landed **11:21:55**, and **five** clones advanced (`app` 98 commits/634
  files, `next-web-app` 41/192, `ant-academy` 5/86, `sentinel` 2/3, `studio-desk` 2/9) — not the one the
  brief named. `N = 24` stands, but **it cannot be PROVEN no adjudicator read post-fetch**, and finding (1)
  is what bounds the exposure rather than any argument. **Path ownership was necessary and NOT sufficient:
  the reading's SUBJECT is wider than the paths anyone declared.** Written into `platform-alignment.md` as
  **§5 rule 41a** — *a reading's ground truth includes the clone refs, so no lane may fetch while a reading
  is in flight.* **(3) CENTRALISING A WORDING CENTRALISES ITS DEFECTS, measured twice in one iter.** The
  canonical sheet exists so five seats do not word one predicate five ways; the price is that one error
  reaches every seat **by construction, with no independent re-derivation to catch it.** Its CANON-1 text
  ended `JUDGE0_BASE_URL (\`:59\`)` — a **bare** anchor, which resolves against the most recently named
  file, `gotenberg.go`, which is **53 lines** — and `repair_postcondition` caught it in the pre-commit hook
  at **four sites at once**. Its CANON-2 text was **weaker than the evidence** (*"not measurable"* alone,
  when the adjudicators had measured **0 hits over 44 tracked `.tf` across 13 clone dirs and 0 over 59 on
  disk**), and two seats had to be re-briefed mid-flight. **The ~2-defects-per-repair-cycle induction rate
  held for the fifth consecutive cycle — this time inside the iter that documents it.** **(4) The corpus
  inherited a false claim by quoting a PLATFORM COMMIT MESSAGE as authoritative** —
  `platform-migration-status.md:101` reported `838d907`'s *"`make up-all` started a second Brevo contact
  pusher alongside backend's own"* as *"the commit names the hazard."* Measured, the second half is false
  (`backend`'s in-process pusher is gated behind an unset `CUSTOMERIO_SYNC_ENABLED`), so `make up-all`
  started exactly **one**. Now quoted as the platform's wording and explicitly **not endorsed**: *a commit
  message is evidence of INTENT, never a measurement.* **(5) Predicate width is real and the booking
  under-counts it — `academy-subgraph-exists` was booked at 2 anchors and lives at 15 sites across 8 files**,
  8 of them in files no seat owned (`CLAUDE.md`, `run_guide.md`, `content-stories-routes.md`,
  `frontend-tier.md`), with the correct model already published at `academy-backend.md:83`. **Four seats
  independently reported their predicates were WIDER outside their partition than inside it**, which is the
  argument for a consolidation pass rather than a wider fan-out. **Also settled this iter, and three of the
  four were questions that should never have reached a user:** `DEF-M257x-iter80-storage-prod-bucket`
  **FILED** to `platform-defect-register.md` (`D-M257x-102-1`) — the register was built for this class by
  M256's audit and had **zero** M257x entries; `FIX-M257x-iter53-union-set` **DROPPED as subsumed 47 iters
  ago** (`D-M257x-102-2`); the gate re-grade **RATIFIED** with clauses 1–2 a **CLOSE BLOCKER**
  (`D-M257x-102-3`); and `CHECK-M257x-iter38-ai-act-classification` **DROPPED** (`D-M257x-102-4`) because
  the repair finished at iter-38 and the corpus asserts nothing — **what was carried for 36 iterations was
  an ASPIRATION, not a defect.** That yields **§5 rule 47** (*close a routed item when its DEFECT is
  repaired, not when its SUBJECT is understood; never route as "needs an owner"*) and **§5 rule 48**, the
  binding user decision **`D-M257x-102-5`: no legal/regulatory escalation during delivery** — route it,
  never ask. **The deferral audit's §8 now has ZERO open user questions.** Its one *urgent* escalation
  (F18, *"three iter-101 tags on origin"*) is a **FALSE POSITIVE** — re-derived, there is **exactly one**,
  no tag points at `4cb920a`, and the pin matches — recorded with **both** candidate mechanisms because a
  peeled-ref line-shape miscount cannot manufacture the specific names, so it is **not a complete
  explanation on its own** (`D-M257x-102-6`). ⚠ **THIS RULING IS REVERSED — see `D-M257x-103-1`. Lane D was
  RIGHT: three tags existed (`-101b`/`-101c` were cut and deleted between the two observations), and the
  refutation measured a concurrently-mutated surface at one instant and treated it as a standing fact. The
  peeled-`^{}` mechanism is real, was never the cause, and is demoted to a caveat. `D-M257x-102-6` is
  SUPERSEDED, not edited; Lane D's finding is restored as UPHELD; the lesson is now `§5 rule 49`.** F16/F17 closed: register entry landed, and `roadmap.md`
  § M257x was **`planned` at 101 iters** with **exit clause 1 still naming the retired `odysseus`** — this
  milestone's own class, in its own gate, for the second time. **REACH GRADED BY MACHINE, not claimed** (`repair_reach_guard` vs each reading's own raw seat reports,
  `iter-102/evidence/reach-grade.md`): **37/46 = 80.4 %** against iter-99's ledger and **29/36 = 80.6 %**
  against iter-101's — **and the entire unreached residue is the adjudicator-REJECTED findings**, which
  must NOT be repaired because a rejection is a claim that turned out TRUE. iter-101 makes it exact:
  **7 of 7 unreached anchors are 7 of its 8 rejections** (`ai_architecture.md:35`/`:141`,
  `security_compliance.md:185`, `backend.md:19`, `skiller.md:19`, `chronos.md:27`, `ai-labs.md:76`); the
  8th sits in a file edited for other reasons. iter-99's 9 line up the same way and include `hiring.md:80`
  twice — the one finding **seat 4 DECLINED with a written derivation** after re-measuring it as true.
  **So: 80.4/80.6 % against the FULL booked set, effectively 100 % against the UPHELD set — both published,
  the first not adjusted away.** **Gate: no clause moved in this pass and none could** — a repair pass
  contains no reading — see `iter-102/progress.md`

- iter-103 (tik, the MEASURING pass — **no repair inside it**): **`N = 33` — THE BURN-DOWN LEG DOES NOT
  REACH THE RESIDUAL**, on the `≥ 23` branch of a rule **sealed in its own commit (`04cbcfc`) before the
  first seat was dealt** and graded here exactly as written. 14 blind seats (readings **#25: 7 of 7** and
  **#26: 7 of 7** — `n₂` is a **SEVEN**-seat count, unlike iter-101's), identical recomputed partition,
  48 bookings → **47 upheld / 1 rejected**, 4 adjudicators grouped by seat LETTER so both readings of one
  file set land with one grader. **The number underneath the number is the finding: by PREDICATE the pool
  did not move at all — 22 at iter-101, 22 here — while anchors went 24 → 33 (1.09 → 1.50 per predicate).
  After a 52-anchor / 98-site repair the corpus carries the same count of false propositions in more
  places.** **Bands 4 HELD of 10.** **#3 HELD at 1 and it EXONERATES the repair**: exactly one of iter-101's
  22 predicates survives (`prod-terraform-8081` at `skiller.md:19` — an anchor iter-102's own repair map
  listed as a twin and flagged **`SEAT 9 (?)`**), so **21 of 22 closed, confirmed blind**, corroborating
  `repair_reach_guard`'s ~100 % upheld-set reach with an independent instrument. **#3b (m = 20 vs [1,7]),
  #4 (97.9 % raw / 100.0 % separated vs [74,88]), #5 (9.1 pts vs ≥15) and #8 (61 % vs ≤10 %) failed
  TOGETHER and are ONE finding: the instrument was byte-identical across the two readings and only the
  SUBJECT moved.** iter-102 repaired the residual's *subtle* half; what is left is **mechanically checkable
  drift** — a version literal, a `go.mod` pin, a symbol name, a line offset — which every competent pass
  finds (overlap explodes, recall spread collapses) and about which a seat has almost no room to be wrong
  (precision goes to 98 %). **Precision, overlap and inter-pass independence are properties of the
  RESIDUAL'S COMPOSITION, not of the instrument; three of those four moved in the direction that FLATTERS
  the reading and none is evidence it got better.** **#10 failed high by one (7/33 anchors in prose
  iter-102 wrote), and the count understates two shapes worth naming. (a) A canonical wording multiplies
  its own defects:** iter-102 closed `prod-terraform-8081` with a sentence saying the
  `backend.internal.anthropos:8081` literal has *"one occurrence anywhere in the clone set"* — **it has
  six**, five inside `rosetta-extensions`, which the same sentence's own 13-repo / 44-`.tf` denominator
  counts as one of its repos. **The replacement is self-refuting against its own stated denominator and it
  shipped to five anchors**, found independently by six seats. **(b) A repair rotted an anchor by inserting
  prose above it — the exact mechanism iter-101 booked against iter-100, one cycle later:**
  `architecture_overview.md:321` **was** the correct local-stack line at `8f04d3a`; iter-102 inserted a
  production-topology block above it, the wording moved to **`:331`**, and every citer stayed put, so
  `:321` now names the **production Cosmo Router** — the opposite topology. Measured corpus-wide: **4 sites
  cite `:321`, 0 cite `:331`**, and the reading found 2 of the 3 in-scope sites while **missing
  `backend.md:54` in both readings** — a detection miss inside its own file set. **Why `N` did not fall,
  as a mechanism: repair reaches its targets, but two inflows feed the residual that repair does not
  touch** — **clone advance** (61 % of `N`; neither platform guard fences version literals, `go.mod` pins
  or line offsets, since `platform_alignment_guard` fences `repos.yml` membership and
  `platform_predicate_guard` fences profile tokens) and **the repair's own induction**. **Inflow ≈ outflow;
  a loop with that property does not converge and running it faster does not help.** **CHAPMAN IS RETIRED
  for this milestone**: its independence assumption has now been measured at both extremes on one unchanged
  instrument — `m`/union **17 %** (iter-101) then **61 %** (here) — so `N̂ ≈ 103` and `N̂ ≈ 35` are both
  unusable, and **only the FLOOR survives (≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`)**. The series
  16.7 → 29.4 → 45.2 → ~103 stays **four corrections to an underestimate, not four measurements of a
  growing pool** — but stop quoting a point estimate from it. **#6 third measurement: 4 → 1 → 1**, so an
  ADDENDUM can carry ground truth a frozen instrument gets wrong **without editing it**; all fourteen seats
  stated which rext tree they read. **#7 HELD at 3** on a tree that grew +214 anchors, so the repair's new
  citations are as good as the old ones. **#9 HELD at 5**, clearing the recomputed partition. **§5 rule 41a
  PROVEN rather than asserted** — every clone HEAD, `origin/main` and fetch timestamp re-read at the close
  is identical to the open; nothing was fetched, no stack touched, no tag cut. **Landed at the close** (both
  deferred by `D-M257x-103-0` precisely because the frozen instrument tells every seat to read §5 in full,
  so editing it mid-read would have split the seats across two rule sets): **§5 rule 49** — *a measurement
  of a concurrently-mutated surface is timestamped, not standing; refuting another observer's report needs
  THEIR timestamp or the surface's HISTORY, never your own later snapshot* — and a **§5 rule 41a
  subsection** stating what it can and cannot enforce (`ensure-clones.sh` fetches unconditionally and moves
  `refs/remotes/*`, the exact surface a citation guard resolves against). **`DEF-4` corrected in this very
  file**: `terraform/main.tf` was NOT byte-identical across the `app` advance and *"the entire residual is
  a label"* was false as stated — the **conclusion** survives and was re-verified, the **evidence sentence**
  was stronger than its evidence. **Gate unchanged at 4 of 5; clause 5 is the only open one and was not
  re-cut, narrowed, reinterpreted or argued.** Routed: `FIX-M257x-iter103-read-union` (22 predicates / 33
  anchors, by CLAIM) and — outranking it — **`FIX-M257x-iter103-drift-fence-gap`**, because repairing 20
  drift anchors without a fence re-arms the class at the next clone advance. **Also `D-M257x-103-7`, found
  at the close and worth the space because the FALSE version was one commit from the record:** the guard
  family came back **2 RED**, contradicting this iter's own ground-truth sheet, reproducibly at `e6aed2e` —
  and the quotable conclusions (*"the sheet asserted a verdict it did not have"*, *"a fence names 8 in-scope
  sites the double reading missed, so `N ≥ 41`"*) were **both false**. The fence had been run from the
  **pinned per-stack clone** (`09d06070`); from the **authoring copy** (`944fc4a2`) both guards are GREEN at
  both subjects, the entire difference being `claim_twin_waivers.json` **+40 lines** (rext `944fc4a`, *"the 8
  acknowledged-site waivers"*) — **the 8 RED sites ARE the 8 waived sites**, and the family re-confirms
  **14 GREEN · 0 RED · 0 could-not-check · 3 not-run**. **A guard VERDICT is not stack behaviour**: §5 rule
  45 sends a claim about what the tooling *does on a stack* to the pinned clone, but a fence's verdict is a
  measurement taken with that fence's **configuration**, so it is settled by the tree the configuration lives
  in — run from the pinned clone you measure **last release's fence**, and every waiver added since reads as
  a fresh RED at sites nobody touched. `guard_family.py` prints the corpus sha and the platform sha and **not
  its own**, so the one input that decides the verdict is the one the output omits — routed as
  **`FIX-M257x-iter103-guard-tree-provenance`**. This is `DEF-M257x-iter101-briefing-rext-tree` **inverted**,
  the coordinator-facing half of the class band #6 measures at 4 → 1 → 1 seat-side. **Three times in one
  iteration the milestone's own class landed on the milestone's own apparatus** — `DEF-4`'s over-strong
  evidence sentence, `D-M257x-103-1`'s single-instant snapshot, and this — **and each time the thing that
  caught it was re-measuring rather than reasoning.** No corpus defect; `N` unchanged at 33; no band moves —
  see `iter-103/progress.md` and `iter-103/adjudication.md`
- iter-104 (tok, **DELIBERATE** — author-initiated, non-terminating): **`TOK-06: fence the inflows before
  repairing again`.** The streak was checked first and does **not** apply (iter-101 moved `N` 28 → 24), so
  this is an author-initiated revision on the TOK-04/TOK-05 precedent, and it does not terminate the call
  because **every element of its sequence is an item iter-103 already routed** — it orders routed work rather
  than opening territory (`D-M257x-104-2`, which also states the bound that keeps that reasoning from
  becoming a general escape). **The revision rests on iter-103's COMPOSITION, not on its number.** `N = 33`
  is the trigger; the finding is that **by predicate the pool did not move at all — 22 then, 22 now** — while
  **21 of iter-101's 22 predicates are CLOSED**, blind, so the repair leg demonstrably reaches what it aims
  at. `N` held up because two inflows feed the residual and nothing watches either: **clone advance, 61 % of
  `N`** (version literals, `go.mod` pins, symbol names, line offsets — the class a fence *can* reach, unlike
  intra-document self-contradiction) and **the repair's own induction, 21 %**, in two mechanical shapes that
  are **both repeats** (a canonical sentence self-refuting against its own stated denominator, shipped to 5
  anchors; and a repair inserting prose above a cited anchor so all 4 citers now name the opposite topology —
  the identical mechanism iter-101 booked against iter-100, one cycle later). **Inflow ≈ outflow: a loop with
  that property does not converge, and running it faster does not help.** So the sequence changes:
  **(0) guard-tree provenance → (1) the drift fence → (2) the induction checks → (3) repair the 33 → (4) read
  LAST.** Provenance goes first because steps 1–2 ship fences and a fence's verdict is settled by the tree its
  configuration lives in — ship the fence first and its founding green is unre-checkable in exactly the way
  that produced iter-103's two false quotable conclusions (`D-M257x-104-4`). Binding on every new fence: a
  mutation control **and** an anti-vacuity control that can actually fire — six fences in this milestone have
  been green over universes they never examined and one compared a string to itself. **Strategy class
  `new-direction`** — the first revision touching neither the instrument nor the unit of repair, but the ORDER
  of the loop. **Chapman sweep landed** (`D-M257x-104-3`): `state.md` was already clean; the two standing point
  estimates surviving in this ledger — iter-101's *"the residual is on the order of ~100"* and iter-102's *"the
  pool was probably always ~100"* — are **marked in place, not rewritten**, and the corrections are
  **asymmetric**, because iter-101's conclusion survives on the floors (≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`)
  and iter-102's *"four corrections to an underestimate, not a growing pool"* survives intact. Clause 5 **not**
  re-cut, narrowed, reinterpreted or argued; gate unchanged at **4 of 5**; zero code, zero platform edits —
  see iter-104/progress.md
- iter-105 (tik, `iter_shape: fence`, **`TOK-06` step 0**): **a guard verdict now states the TREE its
  configuration lives in — and the measured cost of it never having done so is 52 recorded family verdicts
  across 26 milestone artifacts, 0 of which name that tree.** `guard_family.py` printed the corpus sha and
  the platform sha and **not its own**, which is not cosmetic: the fence tree is the input that DECIDES the
  verdict. iter-103 ran the family from the pinned per-stack clone instead of the authoring copy, read **2
  RED** against its own sheet of **14 GREEN**, and drafted two conclusions from it — *"the sheet asserted a
  verdict it did not have"* and *"a fence names 8 sites the reading missed, so `N ≥ 41`"* — **both false**;
  the whole difference was `claim_twin_waivers.json` (+40 lines) and the 8 RED sites were exactly the 8
  waived sites. Neither the corpus nor the platform had moved, **so nothing in the transcript could have told
  the reader which verdict they held.** Shipped: net-new `stack-core/fence_provenance.py` (path · sha · dirty
  · describe, read from **where Python loaded the module** — a flag, a cwd or an env var each reintroduce the
  defect, `D-M257x-105-1`); `guard_family` states it **first**, refuses **`EXIT 2 — UNMEASURED`** when it
  cannot determine its own tree (`--allow-unknown-provenance` **records** the gap as `--allow-not-run` does),
  and puts the **DIRTY** caveat on the **summary** line because harden pass-20 measured that the summary line
  is what gets quoted forward; **all 17 members stamp on direct execution** so a standalone verdict carries it
  too, printed **first** so `run_one`'s `lines[-1]` reporting and `headline()`'s finding-shaped cut are both
  untouched (the iter-87 shape, checked not assumed). **TOK-06's binding control clause discharged:** the
  conformance check is **derived from `guard_family.census()`** so a new guard that does not stamp goes RED
  unbidden, **asserted over the parsed `__main__` AST** per §8, with **4 mutation shapes a grep would pass**
  (no stamp · imports-never-calls · mentions-in-a-comment · stamps-at-import-time) **plus a control on the
  controls**, and an anti-vacuity test written against the **subject** — the discovered set must be *exactly*
  the census, not merely non-empty. **19/19 new tests green.** `corpus/ops/platform-alignment.md` gains **§5
  rule 50** and **§8 rider 3** (*the reference is THREE trees, not two*). **The fence disclosed its own
  weakness on its first run** — every family run in this iter printed `fence tree 944fc4a21 is DIRTY`,
  because its own edits were uncommitted. **The 52 prior verdicts are re-graded provenance-UNSTATED, not
  void** — none is wrong, all are unre-checkable, and the re-grade is stated **once** in rule 50 as a reading
  instruction rather than stamped onto 26 artifacts, because inventing evidence to fix a lack of evidence is
  this milestone's class, not its cure (`D-M257x-105-5`). **`stack-core` 937 passed · 1 failed**, and the
  failure is **PROVEN pre-existing** — reproduced three ways by read-only `git archive` (pre-change rext vs
  live corpus; pre-change rext vs run-open corpus `22eaac4`; changed rext vs `22eaac4`), identical each time.
  It is `claim_twin_guard`'s **green-twin discrimination control**, not the guard: two fixtures fire from
  `iter-49/raw/C.md:57`, and the likely cause is **iter-102's 134 → 264 ledger growth** — repair-induction in
  the fence layer, one layer below where TOK-06 was looking. **Routed, not silenced**, as
  `FIX-M257x-iter105-claimtwin-green-twin-refire` to TOK-06 step 2. Guard family at close **14 GREEN · 0 RED ·
  0 could-not-check · 3 not-run — and for the first time the transcript says which fence tree said so.**
  Gate **unchanged at 4 of 5**; no `N` movement claimed (clause 3's instrument, never clause 5's); zero
  platform edits, no clone fetched, no tag cut — see iter-105/progress.md
- iter-106 (tik, `iter_shape: fence`, **`TOK-06` step 1**): **the 61 % inflow now has a watcher, and on its
  first committed run it fired ONCE — correctly, with zero false positives, having parsed no prose.**
  `stack-core/clone_drift_guard.py` names **`sentinel` as 2 commits past everything the corpus cites**, and
  those two commits are `chore(deps): update dependencies to latest versions` + a version bump, moving colony
  `v0.34.3 → v0.35.2` and proto `v1.200.0 → v1.210.0` — **both of iter-103's booked pin-drift predicates are
  downstream of that single advance.** The fence found the **CAUSE** of two findings that two full reading
  passes could only find the **EFFECTS** of. **The assertion it REFUSES to make is the design**
  (`D-M257x-106-1`): *"a version the corpus states must equal the clone's"* was built, measured and
  **rejected** — `shared_libraries.md:85` states `sentinel v1.200.0` citing `sentinel/go.mod:9 @ 88bc5592`,
  and `git show 88bc5592:go.mod` reads `proto v1.200.0`, so **at that ref the claim was TRUE**; §5 rules
  41/44 make it ref-scoped and a fence calling it false asserts what it never measured. **A fence that cries
  wolf gets suppressed, and a suppressed fence is worse than none** — this milestone has the receipt (a
  silently-refused perf patch shipped a 76 s grid for four releases). So D1 reports the **ADVANCE**, which
  accuses no sentence. **There is no baseline file** (`D-M257x-106-2`): a checked-in `repo → sha` map is §2's
  hand-maintained tuple in a new costume and its first value would have to be *asserted*, so instead every
  backticked sha in `corpus/**` is resolved with `git cat-file` against every clone — **a sha resolves in
  exactly the repo that contains it**, so attribution is exact with no convention and no list (live: **103
  distinct shas · 13 of 14 clones cited · 0 ambiguous**, asserted as a test). **20 tests, 7-mutant battery**
  (incl. *no corpus is `CANNOT RUN`, never clean* and *citing HEAD must CLEAR it, or the fence can never be
  satisfied and gets suppressed*) + **anti-vacuity written against the SUBJECT** per §8's iter-94 rule — the
  LIVE corpus and clone set must yield ≥20 shas / ≥5 attributed repos / 0 ambiguous; every fixture is a real
  git repo because the fence resolves shas with `cat-file`. **D2 (the pin check) ships CONSERVATIVE and its
  low yield is REPORTED, not dressed up** — 1 pin graded, 7 unmeasured; two widenings tried and dropped (a
  nearest-neighbour association is a coin flip dressed as a measurement), and **the reason it is empty is
  corpus-side**: the corpus writes pins as `<repo> <version>` with the module implied by a table heading, so
  the module token is not on the line — recorded as a convention in §8's new **fifth layer** section.
  **The family is committed RED — 14 GREEN · 1 RED over 18 members — deliberately** (`D-M257x-106-3`): clause
  4's own wording is *"asserted by a FENCE that is watched going RED, not by inspection"*, and TOK-06 puts
  repair at step 3 precisely so the repair has something watching it. **Read it correctly: clause 3's fence
  is GREEN, clause 4's schema fence is untouched, and the RED is a NEW member finding PRE-EXISTING drift —
  not a regression in either clause.** Its subject is 5 sites in 2 predicates, both already inside
  `FIX-M257x-iter103-read-union`, so **the fence is now step 3's answer key; a repair that leaves it RED has
  not finished.** `stack-core` **957 passed · 1 failed** (+20 this iter; the failure is iter-105's
  proven-pre-existing one). Gate **unchanged at 4 of 5**, no `N` movement claimed — clause 3's instrument,
  never clause 5's; zero platform edits, clones read but never fetched, no tag cut — see iter-106/progress.md
- iter-107 (tik, `iter_shape: fence`, **`TOK-06` step 2**): **the repair loop's own largest induction shape
  is now fenced AT THE COMMIT — and replaying iter-102's real commit surfaces all four `:321` citers,
  including `backend.md:54`, which the 14-seat double reading MISSED IN BOTH PASSES.**
  `stack-core/anchor_offset_guard.py` is commit-scoped (like `repair_leak_guard`) because the defect is only
  decidable at the commit: looking at `:321` today tells you what is on line 321, not that a citer meant
  something else. §5 rule 34 already names this mechanism and has failed to stop it **twice, one cycle
  apart** — iter-100 booked by iter-101, iter-102 booked by iter-103 — which is why it needed a fence and
  not a restatement. **The design changed TWICE in flight, both times because a control refuted it, and
  both cuts would have shipped green.** (1) The first waived any citation in a file the commit touched —
  sound reasoning, and it returned **GREEN on `cd16967`, the very commit that motivated it**, because
  iter-102 was a 98-site repair that modified all three citing service docs while editing *other* claims in
  them. The carve-out is now **line-level**. General form, worth more than the fix: **a waiver keyed on the
  unit the defect hides inside will waive the defect** (`D-M257x-107-1`). (2) The second went **RED on a
  CORRECT repair** — a synthetic control showed a citer correctly re-pointed `:7 → :9` is indistinguishable
  from iter-102's stale `:321`, because *"post-move and correct"* and *"pre-move and stale"* are both
  consistent with everything the diff records and **intent is not in the repository**. A third narrowing was
  tried and lost the real case. **Resolution: that class is REPORTED, COUNTED and excluded from the exit
  code**, with the OK line stating in its own words that the green does not cover it — §8's *grade the
  cannot-tell*. **A fence must not assert what it cannot decide**; the alternative is a RED that correct
  repairs trigger, i.e. a fence that gets turned off (`D-M257x-107-2`). **The answer key is the COMMIT, not
  a fixture** (`D-M257x-107-3`): `cd16967` → 5 ROT + 5 CANNOT-TELL (the four `:321` citers + a
  `service_taxonomy` self-cite); `a229f8d` → iter-101's booked `service_taxonomy.md:131 → :137/:139`
  induction. **And it surfaced 5 ROT findings NO reading has ever named**, incl. `hiring.md 93 → 107` and
  `storage.md 115 → 129`. Shas **pinned** in the test per §5 rule 25. **Citations are read at the range's END
  revision**, not the working tree (`D-M257x-107-4`) — the measured difference on `cd16967` was real (36 → 33
  seen, 7 → 5 ROT), and it is §5 rule 41a one level down: an instrument resolving against *now* cannot grade
  a measurement taken *then*. **18 tests**: two real commits + six synthetic shapes each separating a case
  the guard must distinguish + refusals (**an empty range is exit 2 with `Nothing was checked`, never a
  pass**) + anti-vacuity against the LIVE corpus (≥25 citations, **0 unresolved** — the `README.md` ×6
  ambiguity resolved by preferring the citer's own directory rather than dropping it). **TOK-06 step 2's
  SECOND shape is NOT taken and is routed with its measurement** (`D-M257x-107-5`): the canonical-wording
  defect is re-confirmed live — `:8081` has **1 occurrence in `app` + 3 in `stack-demo/rosetta-extensions`**,
  a repo the sentence's own 13-repo denominator counts. Gate **unchanged at 4 of 5**, no `N` movement claimed;
  zero platform edits, no clone fetched, no tag cut — see iter-107/progress.md
  **⚠ AND A FINDING AGAINST THIS RUN'S OWN iter-106 DELIVERABLE, recorded at iter-107's close:**
  `clone_drift_guard` went **GREEN** one iter after shipping RED, **with nothing repaired.** The §8 section
  written to *document* the RED contains *"`sentinel` at `f2c46190`, 2 commits past the newest sha the corpus
  cites"* — and that backticked sha **is** a corpus citation of sentinel's HEAD, so the fence's own assertion
  is now literally false while five stale sites remain stale. It is the stated reach behaving as documented,
  and it is the sharpest thing measured today: **writing about the drift satisfies the drift fence.** The
  honest reading of a D1 green is *no repo advanced past every sha the corpus MENTIONS* — mentions, not
  verifies. **Deliberately NOT patched** (an exclusion list of *"docs that are about fences"* is §2's
  hand-maintained tuple again, and inventing one at the end of a long session is how a fence acquires an
  exemption nobody can later justify): recorded in the guard's docstring, **pinned by a known-limitation
  test** (§8 rule 7 — the assertion is what expires), routed as
  `FIX-M257x-iter107-drift-fence-satisfiable-by-prose`. **Fourth time in three days the milestone's class has
  landed on its own apparatus, and the fourth time what caught it was re-running rather than reasoning.**
- iter-108 (tik, `iter_shape: repair`, **`TOK-06` step 3**): **the union is paid BY PREDICATE — 22
  predicates across 23 files — and BOTH new fences fired on the repair before it was allowed to stand,
  which is exactly what sequencing repair after them was for.** Machine-graded: `repair_reach_guard`
  **raw 46/47 = 97.9 %**, and **46/46 = 100 % over the UPHELD union** — the single unreached booking is
  `shared_libraries.md:128` / `r25-G B3`, which adjudicator 4 **REJECTED** (class `wrong-tree`: it graded
  app's post-fold in-tree fork instead of the `ai` module at the `v1.40.2` the section's own pin row names,
  readable in the same clone at `1e457fa70`). **iter-102's residue result reproduced exactly — the apparent
  miss is a claim that came out TRUE**, which is why reach is reported twice (`D-M257x-108-4`).
  `anchor_offset_guard` **OK** on the repair's own range (13 graded, 0 rotted, 7 CANNOT-TELL **each checked
  by hand and all correct post-move**); `clone_drift_guard` **OK with its 2 gradeable pins now MATCHING** —
  iter-106's RED is genuinely discharged, the drift repaired rather than written about.
  **The anchor list was DERIVED, never hand-assembled** (§5 rule 19) via `repair_reach_guard.read_ledger()`
  over `iter-103/raw/` — *the same code path that grades the repair* — 48 blocks / 14 seats → **31 primary
  anchors**; a hand list would have missed `shared_libraries.md:128`.
  **Both of iter-102's induction shapes were refused BY CONSTRUCTION.** The `:8081` canonical wording is
  **removed, not corrected** (`D-M257x-108-2`): re-derived first (**6 occurrences, 0 in any `.tf`, 44 `.tf`
  across 13 repos**), then stated **once** in `backend.md` with `cms.md`/`jobsimulation.md` **pointing at
  it** — *a pointer cannot carry a false cardinality to five places*, so the multiplier is gone rather than
  fixed. And the one line-adding edit was made FIRST, the target re-measured after (**`:335`**), and all
  four `:321` citers re-pointed — **including `backend.md:54`, which the 14-seat double reading missed in
  BOTH passes.**
  **⚠ THE CLASS LANDED ON THE APPARATUS AGAIN, TWICE, IN THIS ITER.** `repair_postcondition` **refused the
  repair commit itself** (§8's iter-102 post-mortem carried bare `:321`/`:331` numbers, which became
  citations onto a blank line once the wording moved); then `anchor_offset_guard` went **RED on the repair
  one commit old** — §8's write-up of the *drift fence* quoted the drift as **live corpus text**, and the
  repair had just fixed it. Fifth occurrence in four days; fifth time it was **re-running, not reasoning**,
  that caught it. Retrospective line numbers now carry the ref they were true at.
  **The 5 routed "rotted" citations were graded ONE AT A TIME, not bulk-bumped** (`D-M257x-108-3`): 3
  resolve correctly today, 1 was genuinely rotted (`backend.md` `:187`→`:241`→**`:302`** — the same anchor
  rotting **three times in three readings**, now marked *cite the claim, not the line*), and 1 is a
  **HISTORICAL** anchor (`hiring.md:93`) deliberately left, because re-pointing it would falsify the record
  of what iter-39 found. **A guard cannot tell a live citation from a record of where something once was.**
  **`FIX-M257x-iter107-drift-fence-satisfiable-by-prose` stays OPEN with no exclusion list** and the
  reasoning recorded (`D-M257x-108-5`): every candidate discriminator is a shape allow-list in a
  derivation's clothes, because *"is this sha dating a claim or being discussed"* is intent, and intent is
  not in the repository — the same wall `D-M257x-107-2` hit.
  **The 3-no-prog tok-trigger was graded out loud and did NOT fire** (`D-M257x-108-1`): iters 105–107 are
  three consecutive tiks with no `N` movement, but they took **no reading at all**, so the metric is
  **UNMEASURED, not unmoved**, and the trigger's precondition is unestablished. **Codified as a protocol
  refinement in §9** — with the floor preserved (three tiks that DID measure still fire it) and two
  mandatory guard-rails — so the next agent inherits a rule rather than re-deriving a judgement call.
  **Side-deliverable, rext `680e852`:** `anchor_offset_guard` **false-greened its own pinned answer key on a
  bare rev** — `git diff <sha>` is *sha vs the working tree*, so the bare form graded **0 of 33 citations
  and printed OK** where `<sha>^..<sha>` is **RED with 10 findings**. All 18 existing tests passed
  throughout, because every one used the explicit form; the defect lived in the invocation an operator types
  first. Normalized, +3 tests, mutation control verified firing. **Found by Phase 0d pre-flighting the fence
  that was about to grade this iter** — without it this iter's own verdict would have been a false green.
  Gate **unchanged at 4 of 5**; **no `N` claimed** — measuring here would be repair inside the measuring
  pass. **`TOK-06` step 4 (the read) is next and is deliberately unstarted.** Zero platform edits, clones
  read but never fetched, no tag cut — see iter-108/progress.md
  **⚠ AND A NET-NEW FINDING ABOUT THE SUITE ITSELF, found only because this close tried to report a
  total** (`D-M257x-108-6`, routed `FIX-M257x-iter108-stackcore-suite-hangs`): a plain `pytest tests/`
  in `stack-core` **BLOCKS INDEFINITELY** at
  `test_m220_mutation_battery.py::DevWiringMutationBattery::test_the_dev_fences_are_red_proven` (45 %) —
  **blocked, not slow** (12.6 s CPU over 3 m 43 s, frozen at 442 results, reproduced in two runs).
  **PROVEN PRE-EXISTING by read-only `git archive` of rext `adcf689`**, verified to contain 0
  `normalize_range` (i.e. without this iter's fix): it hangs identically there, and the module carries
  **0** references to `anchor_offset_guard`, the only module this iter changed. **The consequence is
  worth more than the bug: the standing *"975 pass / 1 fail"* figure CANNOT be produced by a plain
  full-suite run on this host**, because such a run never reaches the end — so whatever produced it used
  an unstated invocation. **A suite total whose invocation is unstated is the same defect class as a
  guard verdict whose tree is unstated** (§5 rule 50 / `fence_provenance`). **State the invocation with
  the count.** This iter therefore claims NO full-suite total: what was actually run is 118 passed over
  the 5 fence modules, 21 passed in the changed module (+3), and the documented pre-existing
  `claim_twin` `test_02` failure reproduced verbatim.

- iter-109 (tik, `reading`): **`TOK-06` step 4 — the read.** `P = 24` predicates / `N = 36` anchors against
  a rule sealed in `ac48e5b` before the first seat. **The pre-registered `P ≥ 15` branch fired: THE POOL
  DOES NOT DRAIN** (secondary `N ≥ 20` fired with it — both metrics agree). Series by predicate: **22 → 22
  → 24**; by anchor **24 → 33 → 36**. 14/14 blind seats, **zero lost**, all committed verbatim before
  adjudication; 35 booked → 32 upheld / 3 rejected; upheld **91.4 % raw and 91.4 % `wrong-tree`-separated**
  (they coincide because `wrong-tree` was **0** — that series is now 4 → 1 → 1 → 0). **7 of 13 bands held.**
  **The reading answers `TOK-06`'s own question and refutes its premise:** all 14 platform clones were held
  at the **identical sha**, so no drift could arrive — and drift is **still ~33 %** of the upheld residual
  (band #8, cut at ≤ 25 % on purpose, FAILED). iter-103's 61 % was a measurement of the residual's
  **composition** read as a measurement of its **flow**; only freezing the subject could separate
  **ARRIVAL** from **DETECTION**, and the answer is **DETECTION**. What *did* work is the induction half:
  band #10 measured **2 of 36** anchors in prose iters 104–108 wrote — **21 % → 5.6 %**, the lowest in the
  series. **Steps 0–2 are vindicated where they apply and re-ranked, not reverted.** The binding constraint
  is net-new and structural: **a repair scoped to a prior reading's DETECTIONS cannot close a predicate**
  — iter-108's 46/46 = 100 % reach is correct, but its ledger came from `iter-103/raw/`, and with per-pass
  recall at 33–83 % that leaves twins standing. Two measured directly, and **one is now a
  self-contradiction because the repair fixed one side of a pair**. Also: `shared_libraries.md:128`
  rejected a **second** time, by a different adjudicator two readings apart (`wrong-tree` → `ref-discipline`)
  — third independent confirmation the claim is TRUE. Clause 5 **not re-cut**; `P = 24` leaves it open.
  Gate unchanged at **4 of 5**. — see `iter-109/progress.md` + `iter-109/adjudication.md`

- iter-110 (tok, **deliberate** — non-terminating): **`TOK-07` authored — the repair's DENOMINATOR moves
  from a prior reading's detections to the corpus, per predicate.** The streak was graded out loud first
  and **could not** have fired (`D-M257x-110-1`): of the last three tiks only iter-109 measured, and §9's
  refinement reads UNMEASURED as UNMEASURED. **What iter-109 refuted is recorded, and so is what it
  KEPT** — `TOK-06`'s premise is gone (a frozen 14-clone subject still yields ~33 % drift, so the 61 %
  was composition read as flow) but its **induction leg took repair-induction 21 % → 5.6 %**, so steps
  0–2 are **re-ranked, not reverted**, and the one leg whose ranking rested entirely on the refuted
  premise — the drift fence — is **de-ranked and left open** rather than cancelled. The binding
  constraint is stated as a **verified** multiplier series, each figure re-opened at its source rather
  than quoted forward (`D-M257x-110-2`): iter-96 13 → **51 sites** (**3.92×**, and it counted the **38**
  an anchor-wise repair would have left) · iter-98 20 → 37 · iter-102 52 → 98 · **iter-108: no
  site-expansion figure at all.** That last row is the finding, and it is stronger than the framing this
  iter was handed — iter-108 did not report a *low* multiplier, it reported **none**, because the
  expansion step was absent, not poor. Four rules: a reading discovers **predicates**, not anchors ·
  enumerate every instance corpus-wide **before** repairing any · **never fix one side of a pair** ·
  **grade reach against the ENUMERATED set** — `100 % of the wrong set` being this milestone's signature
  defect arriving in the check meant to catch it. Order: enumerate → repair whole predicates → **read
  last** (`TOK-06`'s one correct half, kept). Step 0 is the two hardening items with teeth. **A
  falsification of the strategy itself is pre-registered before the next number exists**: another
  `P ≥ 15` after a full enumerate-then-repair cycle refutes repair-and-read as a path to clause 5 and the
  next tok is a **re-scope conversation**, not an eighth revision. Strategy class **`retry-with-evidence`**
  — the method is **iter-96's own**, restored after being silently dropped. Two lessons generalised into
  §9 in the same commit. Clause 5 **not** re-cut; Chapman stays retired, floors only. **No reading taken
  and no `N` movement is claimed.** Gate unchanged at **4 of 5**. Zero platform edits, no clone fetched,
  no stack touched, no tag cut — see `iter-110/progress.md` + `decisions.md` § TOK-07

- iter-111 (tik, `iter_shape: tooling`, **`TOK-07` step 0**): **the two hardening items with teeth —
  one DECIDED, one REFUTED.** (1) `FIX-M257x-harden23-json-polluted-by-provenance-stamp` is closed by
  dissolving its dilemma rather than choosing a side (`D-M257x-111-1`): iter-105's two stated reasons
  for *printed-FIRST-on-stdout* are about **order and shape**, never about the **stream**, so text mode
  is kept **byte for byte** and machine mode puts the tree **INSIDE the document** (`fence_tree`) with
  the human line on stderr — which is the doctrine **strengthened**, since an archived `verdict.json`
  now states its own tree without the terminal that produced it. Mode **derived from argv**, never from
  a flag and never from the environment. 19 guards → `stamp_main()` (mechanical, 19/0 misses), 11 sites
  → `emit_json`. **The workaround died with the defect** (`D-M257x-111-2`): all five sites setting
  `FENCE_PROVENANCE_STAMPED=1` — including the fifth the harden pass missed — now POP it, fenced by a
  disk-wide setter scan whose anti-vacuity control both proves it can see a setter **and** proves it
  does not fire on the removal helper. **Net-new** (`D-M257x-111-3`): `anchor_construct_guard`
  **declared `--json` and read it nowhere** — the false-promise class with its halves swapped —
  implemented and fenced by an AST walk. (2) **`FIX-M257x-iter108-stackcore-suite-hangs` is REFUTED as
  stated** (`D-M257x-111-4`): the suite **completes**, measured three times, ending at **`1 failed ·
  1011 passed` in 1090.88 s** with its invocation stated. The freeze is a **2 m 15 s silent stretch**
  watched directly at 45-second intervals, fully accounted for by a test pytest itself times at
  **132–136 s** whose module measures **142.38 s** alone and which prints nothing while running 8
  nested suites — **low parent CPU is expected, not diagnostic**. Retracted: *"blocks indefinitely"*,
  *"blocked, not slow"*, *"the total cannot be produced on this host"*. **Kept, and it was always the
  better half: STATE THE INVOCATION WITH THE COUNT.** ⚠ **And 16 of the 17 REDs those runs surfaced
  were THIS ITER'S OWN** (`D-M257x-111-5`) — five mutation batteries stage a hand-listed **subset**, a
  new module-scope import made the staged guards unimportable, and the batteries reported *"the fence
  is broken"* for *"you forgot a file"*; one had also pinned a literal this iter moved, silently
  turning a mutant into a **no-op**. Fixed inside the iter (30 passed · 1 failed, 664.44 s); the
  **absence direction is routed**, not faked. The 1 is
  `test_claim_twin_guard_iter48_answer_key::test_02`, **RE-ATTESTED at last** after two sessions of
  *"not re-run"* — run four times, fails exactly as documented (`D-M257x-111-6`). **Run C also found
  that the suite got 2.5× SLOWER by being fixed** (431 s → 1090 s), because a battery that dies on its
  baseline never runs its mutants — *a fast suite is not evidence of a healthy one*. The **second
  3-pass-harden-cap-without-stabilization** is recorded as a standing signal, not converted into a
  request for a fourth pass (`D-M257x-111-7`). Three §8 sections added in the same commit. **No reading
  taken and no `N` movement is claimed.** Gate unchanged at **4 of 5**. Zero platform edits, no clone
  fetched, no stack touched, no tag cut — see `iter-111/progress.md` + `decisions.md`

- iter-112 (tik, `iter_shape: tooling`, **`TOK-07` step 1**): **the enumerator — `FENCE-M257x-iter112`,
  `stack-core/predicate_enumerator.py` — the mechanism `FIX-M257x-iter109-repair-scope-is-detection-bounded`
  asked for.** The judgement/mechanism boundary is drawn OUT LOUD and the judgement half fenced
  (`D-M257x-112-1`): choosing a form is judgement (derived by default), enumerating it is mechanical
  and complete over raw lines **and** re-flowed prose; **seed recall is FAIL-CLOSED** — a form that
  cannot find the site it came from is a RED; **an underivable predicate is exit 2, never 0 sites**;
  and the **multiplier is printed per predicate with `NO-EXPANSION` NAMED**, because iter-108's whole
  failure was an absent step reading as a satisfied one. 18 tests, every control shown firing against a
  mutated copy — including one that restricts the scan to the seed's own file, i.e. **injects exactly
  what a detection-bounded repair does**, and watches the multiplier collapse to 1.0.
  **⚠ The fence caught its OWN derivation twice before any number shipped** (`D-M257x-112-2`): run 1
  padded each seed by a neighbouring line and pulled tokens off *adjacent* propositions (a speech-model
  claim derived `studio/tools/pdf2md.py`) → 1 refusal + 5 seed-recall REDs → fixed by reading **the
  booked range and not one line more**, with range seeds made first-class because collapsing a booked
  range is a narrowing; run 2 reported **36.07×**, *a number about the English language*, because an
  uncapped derivation enumerates a seed line's whole vocabulary → fixed by specificity ranking + a
  4-form cap. **Measured reach, and it decides how step 2 must be done** (`D-M257x-112-3`): **22 of 24
  predicates needed an AUTHORED form — a large share of this residual is PROSE, not citations**
  (`TTS v2 HD`, `Cosmo Router`, *"on the endpoint only, not on the agent name"*). That is the same
  boundary `claim_twin_guard` already draws, now **measured** for this residual rather than assumed.
  **The first per-predicate multiplier in this milestone: 29 seeds → 211 sites → 7.28×, seed recall
  100 %** — and it is **explicitly NOT yet trustworthy** (`D-M257x-112-4`): **12 of 24 read
  `NO-EXPANSION`**, which by `TOK-07`'s own guard-rail scores against **the forms**, not the
  predicates; 4 read as vocabulary (`Cosmo Router` ×37 is 37 mentions of a deleted component). The
  credible middle **is** real — `P10` ×10, `P09` ×6, `P21` ×6, `P15` ×4 — and **every one of those is a
  site iter-109 did not book**, the twin population `D-M257x-109-4` predicted and nothing had ever
  enumerated. **The instrument lands; the measurement does not**, and it is named per predicate rather
  than averaged into a headline. **Step 2 is BLOCKED on a form review** — repairing 211 sites of which
  some are vocabulary would be worse than repairing 46. **No reading taken and no `N` movement is
  claimed.** Gate unchanged at **4 of 5**. Zero platform edits, no clone fetched, no stack touched, no
  tag cut — see `iter-112/progress.md` + `decisions.md` + `enumeration.txt`

- iter-113 (tik, `iter_shape: tooling`, **`TOK-07` step 1, second pass**): **the ceiling —
  `NO-EXPANSION` stops being an assertion**, and the routed blocker
  `FIX-M257x-iter112-forms-need-a-second-pass` is CLOSED. iter-112's report obeyed `TOK-07` rule 2 and
  then had nowhere to go, because **"the form is too narrow" and "the class really is that small"
  produce the identical number** (`D-M257x-113-1`). The fix is a **second, broader form tier** — the
  SUBJECT — from which the ceiling falls out: `headroom` = subject sites − predicate sites, **named by
  `file:line`, never counted**; zero headroom settles the class, non-zero headroom is **RED until every
  candidate is folded in or excluded WITH A REASON**. A predicate with no subject tier is **exit 2
  UNMEASURED**, never a verdict — defaulting it to the predicate tier is the single cheapest way to make
  everything read SATURATED. **Four controls, each shown firing**: lexical refusal of an inverted tier,
  the file-granularity coverage invariant, the **aggregate** anti-vacuity RED (a subject tier that widens
  nothing anywhere was copied, not authored), and the stale-exclusion RED. Plus the paragraph rule
  (`D-M257x-113-2`): prose wraps at ~110 columns so a subject token routinely sits one line above its own
  claim — **one publication seen twice** — and the suppression's own risk is pinned in the same test,
  because the `ai`-fold twin at `external_services.md:554`/`:565` is **eleven lines apart** and must
  survive; a mutation control widens the paragraph to the whole file and watches the twin vanish.
  **30 tests (was 18); 64 passed with `test_fence_provenance.py` alongside, 82.83 s, invocation stated.**
  **The measurement, and it moved DOWN** (`D-M257x-113-3`): **29 seeds → 71 sites, 2.45×** against
  iter-112's 7.28× — because **162 of those 211 sites were four vocabulary forms** (P16 48, P18 58,
  P22 37, P24 19) which now enumerate **5 between them**. The expansion that was real surfaced where the
  flat forms hid it — **P21 6 → 22**, P10 10 → **11** (a **within-file** twin at `cms.md:171`), P15 4→5,
  P02 1→2, and **P12 1→2, a same-fact-different-pin pair (`:1594-1597` vs `:1594-1600`) whose one-sided
  repair would have manufactured exactly the self-contradiction `TOK-07` rule 3 forbids.** **The
  deliverable is the DENOMINATOR — 71 sites — not the multiplier.** **All 16 surviving `NO-EXPANSION`
  predicates are settled, and the verdict is SPLIT in the output because only part of it is mechanical**
  (`D-M257x-113-4`): **1 `SMALL-CLASS-PROVEN`** (zero headroom, nothing judged) and **15
  `SMALL-CLASS-ADJUDICATED`** (**368 candidates enumerated · 254 read and excluded by named reason · 2
  promoted into the enumerated set**). The fence warrants that the candidate set is complete and nothing
  went unexamined; **it does not warrant that the 254 reasons are right, and the close says so** —
  routed as `FIX-M257x-iter113-adjudication-is-judgement`. Three findings step 2 must not repair blind:
  **P08's pin is off by two** (`:496` → the block opens at `:498` — re-derive, never copy), **P13's
  "only" has a counter-example inside the corpus** (`external_services.md:495` records the router built
  from a `git+url` context too), and **P24 is one survivor against ten witnesses**. ⚠ **Two REDs, both
  this iter's own** (`D-M257x-113-5`), caught by controls written before the ledger existed: the coverage
  invariant fired when a subject-form tightening blinded two tiers to their own document (**the repair
  was the form, not the invariant**), and the stale-exclusion RED fired on three rows left behind by the
  two promotions. **No reading taken and no `N` movement is claimed** — §9's refinement applies, `P` is
  UNMEASURED rather than unmoved. **`TOK-07` step 2 is now UNBLOCKED.** Gate unchanged at **4 of 5**.
  Zero platform edits, no clone fetched, no stack touched, no tag cut — see `iter-113/progress.md` +
  `decisions.md` + `enumeration.txt` + `predicate-ledger.json`

- iter-114 (tik, `iter_shape: tooling`, **`TOK-07` step 2, first half**): **the reach metric names its
  DENOMINATOR, or prints no percentage** — rule 4 implemented rather than quoted. `repair_reach_guard`
  took one input, a `raw/` dir of seat reports — *what one reading DETECTED*, at a per-pass recall this
  milestone measured at **33–83 %** — and printed `reach t/N = P%` over it unconditionally. That is the
  instrument that graded **iter-108 at 46/46 = 100 %** while the same propositions stood false one file
  away. Now **exactly one declared denominator, named in the report** (`D-M257x-114-1`):
  `--enumeration` → `corpus-derived-per-predicate`, **may** carry a ratio; `--ledger` →
  `prior-reading-detections`, **may not**. **The refusal is the ABSENCE of the number, never a caveat
  beside it** — a percentage is what gets quoted in a close and a warning next to it is not — and in
  `--json` the **`reach_pct` key is OMITTED rather than nulled**, so a consumer raises instead of
  formatting a figure the run was not entitled to. Refused at exit 2: both inputs at once, neither, an
  empty enumeration, and a **malformed site (refused, not dropped** — a silently shrinking denominator is
  the one direction a reach number flatters itself). **And an UNSETTLED enumeration is not a denominator
  either** (`D-M257x-114-2`): a run reporting `seed_recall_failures` or `unsettled_headroom` is refused,
  because iter-113's ceiling makes "candidate list" and "population" distinguishable and grading against
  the former is the same defect one step earlier. **Measured live, both paths:** the detections path
  prints `reach 109/147 reached — NO PERCENTAGE IS AVAILABLE`; the corpus path prints
  `71 enumerated site(s) over 24 predicate(s) … reach 0/71 = 0.0%` — **step 2's pre-repair baseline,
  measured before the repair rather than assumed after it.** **The iter-81/76 known-answer fixture is
  UNCHANGED** (109 touched · 35 line-unreached · 3 file-unreached · 4 no-anchor · 1 out-of-tree = 152,
  exit 1, `graphql-wundergraph.md:13` named from both readings) — the control proving the extension did
  not soften the fence it extended. Ships with **both** halves of the pair: a positive control (an
  enumeration denominator DOES print a percentage — without it the refusal is satisfiable by a tool that
  never prints one) and an anti-vacuity control that loads **the artifact iter-113 actually checked in**
  (24 predicates / 71 sites), per §8's iter-94 rule. **21 → 30 tests in the file; 60 passed with the
  enumerator's suite, 46 passed for the reach mutation battery + `guard_family`, invocations stated.**
  **No reading taken and no `N` movement is claimed.** Gate unchanged at **4 of 5**. Zero platform edits,
  no clone fetched, no stack touched, no tag cut — see `iter-114/progress.md` + `decisions.md`
- iter-115 (tik, **`TOK-07` step 2, second half**): **the repair — 71/71 enumerated sites, and the
  claim that matters is the other one: 24 of 24 predicates closed at EVERY enumerated instance.**
  Baseline reproduced at `0/71` before a line was written; corpus verified unmoved since the
  enumeration (`git diff --stat 461b547 HEAD -- corpus/` empty); all 14 clones re-verified against
  iter-109's ground-truth table, no fetch. **Seven incremental commits**, so an abandoned run would
  have cost one predicate rather than ~20 files of stranded edits. Both promoted pair-halves landed
  **with** their twins (P10 `cms.md:171`+`:287`; P12 `ai_architecture.md:212`+`jobsimulation.md:160`,
  now the same path AND the same range) — the whole reason they were promoted. The three queued
  findings resolved against source: **P08's pin was off by two and the pin is DELETED, not moved to
  `:498`** (third generation of one same-file anchor; the construct was already named beside it);
  **P13's superlative fell to 18 distinct git-URL contexts** — it was the platform DEFAULT until
  `a2a3ee6` and `customerio-sync` was the LAST, never the only, with the corpus refuting it one file
  away; **P24 was one survivor against ten witnesses**, two of them in `sentinel.md` itself. Substantive
  corrections beyond citations: the `bash -c` claim **inverted a shipped security property** (P10 —
  right about the frozen `cms` repo, wrong about the shipped code); `d11a403` moved **two** of
  messenger's four `*_RPC_ADDR`, not four (P21, 22 sites over 9 files, the end-state claim true
  everywhere and only the agentive form false); storage's env block declares **seven** variables, the
  eight being the cardinality of the cited line RANGE (P16); studio-desk's Express backend holds **no**
  GraphQL client at all (P18); ant-academy **sells a $399/yr subscription to anonymous visitors** and
  calls itself a storefront 213 lines above the denial (P23). **Three sites publishing an enumerated
  predicate were NOT in the enumerated set** — P02's third instance, P22's second anchor (which an
  adjudicator had **booked** and the ceiling pass then excluded), P10's twin — all repaired anyway
  under rule 3 and all booked as **measured** evidence for
  `FIX-M257x-iter113-adjudication-is-judgement`, while the denominator was deliberately **not**
  renegotiated (D-M257x-115-2). **The repair induced anchor rot four times and the fence caught every
  one** — including §5 rule 22's own worked example, rotted a **FOURTH** time by the repair whose
  subject is that class; net-new `FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live` (a guard
  resolving a frozen fixture's path against the live tree). Tests **1 failed / 183 passed / 118.47 s**,
  invocation stated, the failure the **known pre-existing** iter-48 answer-key one over its own
  synthetic fixtures. **No reading was taken; `P` is UNMEASURED, not unmoved** — `TOK-07` step 3 is
  iter-116's entire content and is unblocked. Gate unchanged at **4 of 5** — see iter-115/progress.md

- iter-116 (tik · reading): **`TOK-07` step 3 — THE READ. `P = 37` / `N = 41`, against a rule sealed in
  `85f6f1c` before the first seat was dealt.** The pre-registered **`P ≥ 15` branch fired at more than
  double the threshold**, meeting **`TOK-07`'s OWN falsification: repair-and-read is REFUTED as a path
  to clause 5 under this instrument**, and routing the milestone to a **user re-scope conversation** —
  `TOK-08` is deliberately NOT authored. The reading separates *why*, and that separation is the
  deliverable: **iter-115's repair HELD** (band #3 blind-on-predicate = **3 re-found of 24**; band #11
  `N`/`P` **1.50 → 1.108**), so **DENOMINATOR is falsified and VOLUME is the answer** — **25 of the 37
  are standing pool no prior reading ever detected**, and **9 were induced by iter-115's own repair**,
  whose induction rate rose **5.6 % → 22.0 %** as volume went **+48 → +177** lines with the fences firing
  four times and still leaking. Series by predicate **22 → 22 → 24 → 37**; by anchor **24 → 33 → 36 →
  41**. **7 of 14 bands held** (6 of 8 mechanism, 0 of 3 magnitude); upheld rate **92.6 % raw / 92.6 %
  `wrong-tree`-separated** (`wrong-tree` **0**, series 4 → 1 → 1 → 0 → 0). Two failures are findings:
  **intra-corpus mis-citation is now the LARGEST class at 10 of 37** — the corpus mis-cites itself more
  often than it mis-describes the platform — and net-new band #12 shows **multi-pin blocks concentrate
  defects (6 of 41)**, the sharpest case being the guard-family **FALSE RED** disclosed at the open,
  which flagged `sentinel.md:5` for a proposition that is **true at the ref it names** while the line
  carried a **different, genuinely false** one. 14/14 seats, **0 lost**, each committed verbatim before
  adjudication; 4/4 verdicts committed unedited; 54 booked → 50 upheld / 4 rejected. All 14 clones
  identical at close and open, fetch times unchanged. **No repair taken inside the pass.** Gate **4 of
  5**; clause 5 **NOT met** — see iter-116/progress.md + iter-116/adjudication.md

- iter-117 (tik · `iter_shape: census`, **`TOK-08` class 1 — the USER's re-scope, sealed in this iter's
  FIRST commit `577446b` before any sweep work**): **the corpus stops mis-citing itself.**
  **1,520 intra-corpus citations enumerated over 92 source documents** — C1 path 1,337 · C2 anchor 179 ·
  C3 line-pin 4, denominator **`corpus-derived-per-arm`**, stated in the report and in `--json` per
  iter-114's rule. **8 already false, all 8 in the anchor arm, 8 of 8 repaired**; `corpus_citation_guard`
  ships green over **100 % of the enumerated population** and is registered in the postcondition ratchet
  at **zero** sites, not a tolerated count. **Zero** false of 1,337 path citations and **zero** of 4
  line-pins: the whole class lives in the one arm **no reader exercises**, because a broken `#fragment`
  still lands you on the right file — which is how 8 defects survived four graded readings. The eight:
  `platform-migration-status.md` ×4 (one extra hyphen in every deep link into the protocol doc — `## 6.`
  slugs the `.` to a SINGLE separator), `secrets-spec.md` (anchor stale by two milestones, M27–M28 for a
  heading reading M27–M30), `staging-bringup.md` (paraphrased from memory), `architecture_overview.md`
  (*"see AI Providers below"* for a heading that does not exist), and `coverage-protocol.md` — **a literal
  `[serve-grant](#…)` ellipsis placeholder that shipped**. **Four census passes ran before one line of
  prose was repaired**, three of them only to kill a false-positive class: basename-as-path (~180 false
  REDs), bare-`:NN`-as-corpus-pin (256), other repos' `knowledge/` (5), heading-only anchor model (22,
  including **CLAUDE.md's own retraction link**) — ~460 correct citations that a fence shipped on any one
  draft would have turned RED. Every exclusion is annotated **with the number it cost** and pinned by a
  regression test. **The mutation control earned its keep on its first run**, catching a **silent vacuity
  bug in this very guard** (an unresolved root made the cross-document arm enumerate nothing under a
  symlinked tree and report a clean pass) — the ninth vacuous fence caught here, and the first caught
  before shipping. **Reported against the pre-registration and before the reading that will grade it:
  the machine-reachable half of class 1 is largely DISJOINT from the 10 predicates iter-116 booked** —
  those are *construct* defects, and of 387 bare-pin lines only **4** are machine-resolvable — so little
  `P` movement should be expected from this class (`D-M257x-117-2`). Side finding, **pre-existing at rext
  HEAD** and proven so by read-only `git archive`: `test_repair_reach_guard.py` hid **14 tests** below its
  `__main__` guard since iter-114, printing OK over 16 of 30; `test_test_collection_fence` had been RED on
  it the whole time and the RED was invisible because the suite does not complete on this host. Fixed;
  30/30 collect. **`state.md` repaired by RELOCATION, not trimming** — `phase:` measured **2,230 B against
  its 900 B budget**, carrying a reading already written down here; now 895/900 and the file has **1,929 B
  of headroom** where it had ~200. Tests: 19/19 guard · 27/27 + 25/25 postcondition · 34/34 provenance ·
  41/41 guard_family · 8/8 collection-fence · guard_family live **14 GREEN / 0 RED**, every count naming
  its invocation. **No reading was taken and no `N`/`P` movement is claimed** (§9's UNMEASURED rule).
  Gate unchanged at **4 of 5** — see iter-117/progress.md + decisions.md

- iter-118 (tik · `iter_shape: census`, **`TOK-08` class 2 — platform-source citation resolution**):
  **the fence was green over a subject it only two-thirds reached.** `anchor_construct_guard` already
  reported *"675 resolved / 514 unresolvable"* and already NAMED its unresolvable heads — and nobody had
  read inside the head. **119 of the 514 are not citations at all**: `http://sentinel:8087` matches the
  qualified-anchor regex because `sentinel` is a path-ish head and `8087` is a line number (`http:` ×87,
  `AUTHORIZATION_ADDRESS=http:` ×11, `GOTENBERG_URL=http:` ×11, a tail of `VAR=http:`). The fence was
  grading its own coverage against a target inflated by things that could never resolve because they were
  never citations — **iter-114's denominator-provenance rule arriving one layer over, inside the guard
  that reports reach** — and it survived because the guard **printed no ratio at all**, so there was no
  number to distrust. Now: **reach 675/1,070 = 63.1 %, denominator `citation-candidates-minus-non-citations`**,
  with the 119 exclusions **counted and named** (a denominator that shrinks in silence is the same defect
  facing the other way). **0 findings** — the class was and is clean over everything it resolves. The
  genuine 395-site residual is now measured rather than hidden: `(bare)` ×276 (ports / continuations —
  the same undecidable shape class 1 measured), `main.go` ×27 (every Go repo has one), ~92 single-file
  basenames; the repo-disambiguation rule that would reach most of them is **routed, not attempted**
  (`FIX-M257x-iter118-bare-basename-needs-repo-disambiguation`). Controls are asymmetric on purpose:
  when a fix REMOVES from a denominator, **anti-vacuity is the load-bearing half** — six real citation
  shapes must STAY in, including `app/internal/httpclient/do.go`, a genuine path containing *http*.
  Side-deliverable: `test_iter45_mechanical_fences::test_21` hard-coded a four-fence baseline set and
  broke on every fence that joined the ratchet; it now DERIVES the set from disk (iter-44's own lesson
  applied to the test guarding it) with an anti-vacuity floor. Tests 83/83 · 6/6 · 19/19 · 41/41 · 27/27 ·
  8/8, guard_family live **14 GREEN / 0 RED**, every count naming its invocation. **The full mechanical
  sweep `TOK-08` pre-registered is COMPLETE** — class 1 at 100 % reach, class 2 at 63.1 %, both at 0
  findings, class list fixed in iter-117 and unchanged — **so iter-119 is the grading reading**, branch
  `P >= 19` refuted / `P <= 18` working against `P = 37` at `f581de09`. **No reading taken this iter and
  no `N`/`P` movement is claimed.** Gate unchanged at **4 of 5** — see iter-118/progress.md + decisions.md

- iter-119 (tik · `iter_shape: reading`, **`TOK-08`'s GRADING READING**): **`P = 22` / `N = 28` — the
  pre-registered `P >= 19` branch FIRES. `TOK-08` — the USER's re-scope — is REFUTED by its own sealed
  arithmetic**, and per that rule **no successor strategy is authored: there is no `TOK-09`.** The
  milestone returns to the user for a **scope decision**. Pre-registration re-sealed unchanged in the
  iter's FIRST commit `4d4530d` before a seat was dealt; instrument copied verbatim (sha `3858ec53`,
  `diff` empty after copy AND after the addendum, one commit ever); 14/14 seats, **0 lost**, each
  committed verbatim before adjudication; 4/4 verdicts unedited; **36 booked → 33 upheld / 3 rejected,
  all `ref-discipline`, 0 `wrong-tree`** (upheld rate **91.7 % raw = 91.7 % separated**). **The 40.5 %
  fall from 37 is SAMPLING, NOT DRAINAGE, and the same reading proves it:** the corpus moved **5 in-place
  lines, zero net**, and **none of iter-116's 37 predicates was repaired** — so band #3, the milestone's
  **first test-retest measurement**, re-found only **13 of 37 (35.1 %)** and added **9 net-new**, meaning
  **`P` FELL 37 → 22 while the measured floor ROSE 37 → ≥ 46.** It was measurable at all only because
  iter-119 is the milestone's **first true seat-level replicate** — partition bit-identical to iter-116's
  (reproduced at both refs as a control), 14 clones frozen for a **fourth** consecutive reading, and the
  low-overlap branch was named in the pre-registration in advance as the outcome that *"would bear on the
  refutation more than the primary itself does."* **Every `P` this milestone published (22, 22, 24, 37,
  22) samples about a third of a standing pool**; two passes recalled **63.6 % / 77.3 %** against their
  own union, and `P` is comparable across readings only to **±2** on adjudicator granularity alone
  (disclosed: iter-116 split `ai_architecture.md`'s three stale self-citations into 3 predicates, this
  panel collapsed them into 1 — split, `P = 24`; every granularity reading fires the same branch).
  **What `TOK-08` DID achieve, stated because refuted is not worthless:** class 1 censused and closed
  (1,520 citations / 92 docs, 8 false, 8 repaired, **100 %** reach, ratchet at zero) and class 2 censused
  (1,070 candidates, 0 false, **63.1 %** reach, denominator named) — **8 defects closed that four graded
  readings had missed**, all in the `#anchor` arm no reader exercises. **Why it still failed is
  `D-M257x-117-2`, recorded before either census closed: the census and the reader work OPPOSITE HALVES of
  one class.** Band #7 (predicted ≤ 6) **failed at 8** — the construct half grew **27 % → 36 %** of the
  pool, and **5 of the 8 were booked at iter-116 and are still false**. **Graded by consequence:**
  `clerk-integration.md:40` says Clerk sign-in tokens are minted *"only"* for app-native admin
  impersonation — **two other live minting sites exist**, a security-surface understatement now standing
  through its **second** reading. **The 16 small-class verdicts, audited as asked:** the sweep promoted
  **0 of 16** from judged to proven (a census measures citations; a small-class verdict claims a
  *ceiling*, which neither census measures), **P20 remains the only `SMALL-CLASS-PROVEN`**, and this
  reading **re-books 2 of the 15 judged** (P23 Ant Academy, P08 the `⚠⚠ M51` anchor — fifth generation)
  and **0 of the 1 proven** → measured error rate **>= 13.3 %**. **Bands 13 HELD of 15** (series 4/9 · 3/7
  · 5/9 · 4/10 · 7/13 · 7/14 · **13/15**; all four magnitude bands held for the first time), failures #7
  and #8 (the latter narrowly, boundary disclosed). All 14 clones **identical at close and open**, fetch
  times unchanged, **zero corpus edits inside the measuring pass**. Gate **4 of 5**; clause 5 **NOT met**
  — see iter-119/progress.md + adjudication.md + decisions.md

- harden pass 26 + iter-120 repairs (**no reading taken; `P` is UNMEASURED, not unmoved**): the
  consequence-graded defects closed, the 8 wrong-construct citations closed, and **two live fence defects
  found by attacking the young instruments rather than running them.** `clerk-integration.md:40`'s *"only"*
  is repaired **by enumeration** — three instruments over the whole clone set return **five** sign-in-token
  minting sites, and **two of the five appear in NO prior reading** (iter-116 and iter-119 both named
  three). A **second** security-surface understatement was found and repaired at **four** sites:
  *"Sentinel validates **every** API request"* — the platform's own source comment calls that blanket gate
  **FAILS OPEN** and six paths reach the resolver before the single Sentinel call. Fence side:
  FENCE-M257x-iter117's mutation AND anti-vacuity controls **could not be COLLECTED** on this host (PEP 604
  under a 3.9 pytest) and their collection error **aborted all 2,837 tests**; G10 was a **false RED on a
  correct corpus**, firing because iter-115's repair added a *more precise* ref. Both fixed, both
  mutation-proven. Gate **4 of 5**; clause 5 **NOT met** — see the hardening ledger, pass 26.

## Clause 5 — what the milestone measured about its own instrument

**Stated as measurement. This section makes no recommendation and argues for no change to the clause.**
Clause 5 stands exactly as written, as ruled four times: *KB-fidelity audit GREEN, or YELLOW with 0
blockers, over `corpus/services/**` + `corpus/architecture/**`.* Nothing below re-cuts, reinterprets or
narrows it.

### 1. `P` fell 37 → 22, and the fall is SAMPLING, not drainage

Between iter-116 and iter-119 the corpus moved **5 in-place lines, net zero**, and **none of iter-116's 37
predicates was repaired**. The headline `P` fell 37 → 22 while the **measured floor ROSE 37 → ≥ 46**. A
number that falls while its own lower bound rises is not measuring the quantity its name implies.

### 2. Every `P` this milestone published samples roughly a third of a standing pool

The series is **22, 22, 24, 37, 22**. iter-119 was the first true seat-level replicate — partition
bit-identical to iter-116's — and re-found **13 of iter-116's 37 (35.1 %)**. Within iter-119 the two
passes recalled **63.6 % and 77.3 %** against their own union, and **~35 %** against the floor.

**A third, independent measurement of the same thing, obtained this pass by ENUMERATION rather than by
replicate reading:** on the single predicate *"where are Clerk sign-in tokens minted"*, the true answer is
**five** sites. iter-116 booked three. iter-119 booked the same three. **Two readings, run three iters
apart, independently missed the same two sites** — a per-predicate recall of **3/5** with the misses
*correlated*, not random. Correlated misses are the case where more passes do not converge.

### 3. Therefore a gate phrased as *"a reading that returns zero"* cannot be REACHED by this instrument

Not *"is not yet met"* — **cannot be reached**, because the instrument cannot see most of what it must
return zero on. Two passes at ~35 % recall against a floor cannot demonstrate the absence of what neither
saw. **This is a statement about the instrument, not an argument about the clause.**

A fourth bound arrived this pass, from the fence side, and it bounds the *classes* too:
`anchor_construct_guard` detects **"resolves to blank"**, not **"resolves to the right construct"**. A
ninth wrong-construct citation sat in `platform-alignment.md` GREEN for the whole milestone and became
visible only when an unrelated edit shifted its target onto a blank line. So iter-119's *"8 of 22
wrong-construct"* is itself a **floor**, measured by an instrument with a known and now-quantified blind
side — `FIX-M257x-iter120-anchor-guard-detects-blank-not-wrong`.

### 4. 0 of 16 small-class verdicts were promoted from JUDGED to PROVEN

**A census measures citations; a small-class verdict claims a ceiling; neither census measures a ceiling.**
The two censuses (iter-117: 1,520 intra-corpus citations; iter-118: 1,070 class-2 candidates) closed the
**resolution** half of intra-corpus citation at 100 % reach — and the **construct** half, which no machine
in this family reads, GREW as a share of the pool, **27 % → 36 %**. Measured error rate on the judged set
is **≥ 13.3 %**.

### 5. What is NOT claimed here

- **No point estimate of the pool.** Chapman is retired; floors only (**≥ 46 at `194361e4`**).
- **No claim that the corpus got worse.** iter-120's repairs are real repairs against source.
- **No claim that the remaining pool is small, or large.** It is unmeasured above the floor.
- **No successor strategy.** `TOK-07` and `TOK-08` were each refuted by their own pre-registered
  arithmetic; per `TOK-08`'s sealed rule there is **no `TOK-09`**, and none is authored here.

---

## iter-121 — the work that is needed under EITHER scope outcome (2026-08-07)

**Type:** tik · `iter_shape: instrument + filing` · **No reading was taken and none is implied. `P` is
UNMEASURED, exactly as after pass 26.** No successor strategy is authored; clause 5 is not re-cut,
reinterpreted, narrowed or argued. **No fourth harden pass was requested** — see §3 below.

### 1. Two PLATFORM findings filed — they are not documentation defects, and the corpus repair did not address them

iter-120 repaired the corpus's flattering security claims. Two of the underlying facts are **platform**
concerns in their own right. Both are now in
[`platform-defect-register.md`](../../platform-defect-register.md) with `file:line` and their measurement,
which takes that file from 5 M257x-era entries to 7 — and it is worth naming that the register **had zero
M257x entries until iter-102**, having been created by M256's deferral audit for exactly this class.

- **`PLATFORM-M257x-graphql-authz-middleware-FAILS-OPEN-and-REST-has-no-blanket-gate`** — the six paths
  that reach the resolver before the single Sentinel call, the one hardcoded target variable, the REST
  surface, and the platform's own post-mortem comment. Re-derived line by line at `app` `ad9f3c49`.
  Carries, **without editorialising**, the impersonation mutation's gate:
  `ActionObjectTaxonomy` / `UserActionWrite` (`resolver_admin_audit.go:20-24`) — a taxonomy-write
  permission rather than a dedicated one; the sibling `adminAuditLogs` query uses the identical pair.
- **`PLATFORM-M257x-dev-login-routes-mint-a-full-session-for-any-email-behind-one-NODE_ENV-boolean`** —
  and **one premise of the routing was re-derived and is FALSE, which is recorded first in the entry
  rather than quietly dropped.**

  > The item came in as *"`ant-academy/…/login-as/route.js:78` has **no `NODE_ENV` gate**."* **It does** —
  > `route.js:34` refuses on `!DEV_LOGIN_ENABLED`, and `code/src/lib/devLogin.js:29` is
  > `NODE_ENV !== 'production'`. The **ungated** site is `next-web-app/e2e/auth.setup.ts:72`, and the
  > *"skips both factors"* comment is that file's (`:57-62`), not ant-academy's. **Two of the five minting
  > sites had been conflated** — which is itself an instance of the class this milestone measures, arriving
  > in the routing rather than in the corpus.

  What was filed instead is what is true, and it is still worth a platform engineer's attention: three of
  the five sites are made **public** (unauthenticated) in middleware by the same boolean that mounts them
  (`ant-academy/code/proxy.js:178`, `next-web-app/apps/web/src/proxy.ts:56`); the email is taken from the
  query string with **no allowlist and no shared secret**; and the whole control is one comparison against
  a **build-mode** variable, not a deployment gate.

### 2. A TENTH quantifier defect — in iter-120's own repair, pointing the other way

iter-120 replaced *"Sentinel validates **every** API request"* with *"the REST surface has **no authz
middleware at all**"*. **Also false.** `cbGate := courseBuilderAccessGate(authorizationManager)`
(`backend.go:227`, defined `internal/web/backend/gate.go:27-49`) IS a Sentinel-backed **group** middleware
(`OrgCheckFeaturePermission(OrgFeatureMembersEdit, orgID)`), applied to `/coursebuilder` (`:229-232`) and
`/credits` (`:273-276`). The repair's own citations — `:230-231` and `:274-275` — **each stop one line
short of it.**

Repaired at all three sites that carried it (`security_compliance.md`, `architecture_overview.md`,
`backend.md`), with the full six-group table replacing the assertion. **Two things this is evidence for,
and they point in opposite directions:** the repair *class* is right — the conclusion (no blanket gate)
survives — and **a repair of an absolute quantifier introduced a new absolute quantifier in one pass.**
Also: a citation that stops one line short of its own subject is precisely the wrong-construct class
`anchor_construct_guard` does not detect, found here by reading rather than by any fence.

### 3. Three instruments closed or floored — see the rext commit for the full derivations

| item | outcome |
|---|---|
| `FIX-M257x-iter108-stackcore-suite-hangs` | **closed as stated** — the ambiguity, not the runtime, was the defect. `progress_beacon.py` + 7 wired sites + an AST-derived wiring fence. §4 below records the standing invocation |
| `FIX-M257x-iter120-anchor-guard-detects-blank-not-wrong` | **DECLINED with a measurement, floor DISCLOSED in the instrument.** Only **28 of 511** backticked citations (5.5 %) supply their own expected content, and binding that content to the right anchor needs the sentence's claim — 1 of 20 measurably binds to a LATER anchor. Fourth decline by this fence family, first with numbers. `KNOWN_WEAKNESS` now prints on every run and rides in `--json` |
| `audit-deferrals` §8 read one field | **closed** — `blocking_state_guard.py` derives over all three blocking-capable fields. It found **8** blocking gradings across 109 graded iters where the audit named **3**; `deferrals-audit.md` §12 enumerates all eight with dispositions, and §8's banner is superseded with **Q5** written out |

**Standing rules added** (they outlive the milestone): `platform-alignment.md` §5 **rules 51, 52, 53**.

### 3b. The whole-suite claim, and it is now defensible — with the invocation it rests on

```
cd .agentspace/rosetta-extensions/stack-core
STACKCORE_PROGRESS_LOG=/tmp/m257x-beacon2.log \
  /usr/bin/python3 -m pytest tests/ -q --tb=line -p no:cacheprovider --no-header --durations=5
      ->  1 failed · 1125 passed  in  1032.57 s (17:12)   [rext 1bb64c3, corpus mid-iter-121]
```

`/usr/bin/python3` is **3.9.6** and is the only interpreter on this host with pytest. The single failure is
`test_claim_twin_guard_iter48_answer_key::test_02` — the standing, documented RED, **and it is now
genuinely re-attested by a full run** rather than carried as *"unchanged, not re-verified"* for a fourth
consecutive pass. **239 beacon lines** were emitted; the formerly-silent stretch is now eight timestamped
nested-run pairs.

**The expected wall time is ~1030 s and it reconciles with iter-111's 1090.88 s to within 6 %** — which is
the check rule 51 asks for, and it did not pass on the first attempt. **The FIRST run today returned
632.20 s with 4 failures, and the 460 s gap IS the defect** (§3c): a battery that dies on its baseline
never runs its mutants, so `test_01_every_mutant_matches_its_DECLARED_verdict` cost **45 s** while broken
and **405.69 s** once fixed. *A fast suite is not evidence of a healthy one* — iter-111's own words,
re-measured. **That first run was also tree-edited mid-run and is disclosed as confounded**, exactly as
iter-111 disclosed its run A; the 1032.57 s reading above is the clean one.

### 3c. The first whole-suite run found a RED that four iters of scoped runs could not

**Not introduced by this iter, and that was PROVEN rather than argued:** re-run in a detached worktree at
`b9bb2b6` (iter-121's parent) → identical 3 failures, identical named test. Worktree removed and pruned;
rext stayed on `main`.

`corpus_citation_guard.py` (FENCE-M257x-iter117) joined the participating baseline in
`repair_postcondition_baseline.json` and was **never added** to the mechanical-fences battery's
`_COPY_FILES`. `discover_fences()` globs `*_guard.py` on disk, so the staged tree found **4** participating
fences while its own staged baseline named **5** →
`test_21_the_shipped_baseline_records_EVERY_participating_fence` RED → **ANTI-THEATRE #1 reports a RED
BASELINE**, i.e. *"the fence is broken"* for *"you forgot a file."* **This is iter-111's defect recurring
four iters later on a different file** (§8 rule 7's recurrence corollary: the class, not the site).

Fixed, and the second half is the one that matters: `test_000_the_copy_list_stages_every_fence_the_baseline_names`
is the **ABSENCE direction** iter-111 booked as `FIX-M257x-iter111-staged-battery-dependency-is-underived`
and deferred as *"an import graph, not a grep."* **For this battery it is not an import graph** — the
staged tree carries its own baseline, the baseline names the fences, and discovery is a glob. Battery:
3 failed / 1 passed / 1 skipped → **6 passed in 433.27 s**.

**This is the argument for the beacon, made by the tree rather than by me.** The defect was live for four
iters and three harden passes, every one of which ran scoped suites and each of which recorded *"no
whole-suite total is quoted anywhere in this entry."* The first whole-suite run after the silence was made
readable found it in one pass.

### 3d. Guard family after the corpus edits

**21 members · 17 GREEN · 0 RED · 4 not-run** (`anchor_offset_guard`, `repair_leak_guard`,
`repair_reach_guard`, `value_change_guard` — the commit- and input-scoped members, no `--range`/`--ledger`
supplied). **Not a whole-family green, and the runner's own summary line says so.** The family grew 20 → 21
with `blocking_state_guard`. Invocation: `guard_family.py --repo-root <rosetta> --platform
<rosetta>/stack-demo/platform`.

### 4. The harden cap fired at 3-in-a-row (22, 25, 26) — recorded as a standing finding, NOT a fourth pass

Pass 26 gave the sound reason not to repeat in the same mode: *only mutating the named mechanism surfaces
this class, and running the suite never will.* That is now **rule 53**, together with the pass's
load-bearing method finding — **three of its own mutations silently failed to apply and each read as "the
controls survive"** — which is why every mutation added this iter asserts `count == 1` before its result
is interpreted, and why the beacon mutant removes **both** calls (removing one leaves the site still
beaconing, i.e. a mutant that does not isolate the property).

**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**a successor strategy is FORBIDDEN by
`TOK-08`'s sealed rule and none is authored**) — (3) re-scope: **y — unchanged and carried: the milestone
is still holding on the iter-119 scope decision. This iter neither advances nor argues it** —
(4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome:
**exit-3**
**Decisions:** `D-M257x-121-1` … `D-M257x-121-6` (see [`decisions.md`](decisions.md))

---

## iter-122 — the CLAIM CENSUS: the instrument the user funded (2026-08-07)

**Type:** tik · `iter_shape: instrument + census + adjudication` · **No reading was taken and none is
implied. `P` is UNMEASURED.** Clause 5 is not re-cut, reinterpreted, narrowed or argued, and **the gate
did not move — it stands at 4 of 5.**

**The user ruled.** Given the honest choice — close at 4 of 5 with clause 5 documented as
measured-and-unmet, or fund a materially different instrument — they chose the instrument. iter-122
built it, and **sealed its design and its six falsifications before a line of it existed**
(`1262ca7`), exactly as `TOK-07` and `TOK-08` were sealed. The full result is
[`iter-122/adjudication.md`](iter-122/adjudication.md); the sealed design is
[`iter-122/pre-registration.md`](iter-122/pre-registration.md).

### 1. The headline is a FAILED BAND, and it failed on the low side

`F5` sealed the tier-1 `DOES-NOT-SUPPORT` rate at **≥ 4 % and ≤ 25 %**, built from iter-119's
*"measured error rate on the judged set is ≥ 13.3 %"*. Measured over **427 decidable** adjudications:
**3 = 0.70 %**, and **1 = 0.23 %** once two verdicts are traced to the census's own stale substrate (§4).

**A hunted sample over-states the population rate by roughly nineteen-fold.** A reading stops on the
sentences that look wrong, so the rate among sentences a hunter stopped on was never the rate in the
population. **Every error rate this milestone published from a graded read is a rate over a hunted
sample** — that is a fact about the milestone's own prior measurements, and it was obtainable only by
censusing what a reading sampled.

The other two bands held: tier-2 count **1,151** (band 200–2,500); tier-1 pairs **2,603** (band
900–3,000).

### 2. TIER 1 — 525 of 525 adjudicated, exhaustive over the class it names

**2,603** (claiming unit × citation) pairs over the 40 files, across **932 of 2,485** claiming units.
**695** carry a line pin; **525** materialized; **all 525 were adjudicated** by 12 blind adjudicators,
**525 rows for 525 items, 0 seats lost**, verdicts committed verbatim before aggregation.

Clean tally (41 quarantined, §4): **`SUPPORTS` 407 · `PARTIAL` 17 · `DOES-NOT-SUPPORT` 3 ·
`UNRESOLVABLE` 57.**

**The dominant non-`SUPPORTS` class is not error — it is UNCHECKABILITY.** 57 of 484 (11.8 %) are pinned
to commits this tree is not at, with the construct alive at a different line (`app/main.go`'s handler
registrations pinned `@ b948604f` are ~110 lines adrift at `ad9f3c49`). **A re-pin backlog, not a
retraction backlog — and nineteen times the larger finding.**

**One genuine wrong-construct citation, found and repaired:** `academy-backend.md:62` cited
`app/main.go:471-472` for a construction at `:524-525` — **53 lines short of its own subject**, on a
clone level with `origin/main`, so no substrate excuse. `app`'s own `CLAUDE.md` names `:524`/`:525`.

**`F2` FIRES for tier 1 as a whole and is reported with arithmetic rather than a flattering number:**
525 of 2,603 = **20.2 %**; extrapolating ≈ **60 adjudicator runs / ~10.4 M subagent tokens**, and the
1,908 unadjudicated pairs are *harder* (826 bare shas, 525 unpinned files, 507 doc-links). **The honest
form is "100 % of a named 525-pair class", never "20 % of the corpus checked."**

### 3. TIER 2 — the half no fence in this family had ever looked at

Denominator named per `F6`: **3,292 assertion candidates** in the 40 files.

| | count | share |
|---|---|---|
| CITED (a citation in its block) | 2,117 | 64.3 % |
| UNCITED but **HEDGED** | **24** | **0.7 %** |
| UNCITED and **UNHEDGED** → the defect | **1,151** | **35.0 %** |

**1,175 factual assertions carry no citation; 1,151 of them — 98.0 % — carry no hedge either.** The
iter-093 principle is honoured in **24 places out of 1,175 opportunities**. They live in **975 blocks
across 39 of the 40 files**, enumerated line by line in
[`iter-122/raw/tier2-unevidenced-assertions.tsv`](iter-122/raw/tier2-unevidenced-assertions.tsv).

**UNEVIDENCED, never FALSE**, and **a FLOOR** — block scope exonerates a whole paragraph on one citation.
Both disclosures print in the guard's own output.

### 4. The census's THREE OWN defects, recorded first

| # | defect | caught by | consequence |
|---|---|---|---|
| 1 | the declared archived/external/rext artifact names were never unioned into the derived set — `redis`, `clerk`, `directus`, `skiller` invisible to the proxy | **mutation control 11**, pre-publication | tier 2 undercounted by **229** (922 → 1,151) |
| 2 | a **bare basename** resolved by probing clone dirs alphabetically — a silent guess | **control `test_22b`**, post-dispatch | **41 of 525** quarantined; independently flagged by **6 of 12** adjudicators |
| 3 | materialization read the clones' **WORKING TREES**; 6 of 13 are behind their own fetched `origin/main` | **the adjudicators**, not a control | **2 false `DOES-NOT-SUPPORT`** verdicts |

Defect 3 is the sharpest thing this iter learned. `storage` is **20** commits behind, `messenger` 7,
`jobsimulation` 4, `next-web-app` 4, `cms` 2, rext 1. Four independent adjudicators booked the corpus's
M810 claim as contradicted — and **the corpus is right**: `6092c6d2` **is an ancestor of `origin/main`
`82cb66ec` in the very clone the census read from**.

> **A stale substrate does not merely fail to confirm a claim — it manufactures evidence AGAINST a true
> one.** The guard cannot fix the checkout (the clone set belongs to a live demo stack this milestone may
> not touch), so it discloses: substrate on every materialization, a staleness table on every run, and
> `KNOWN_WEAKNESS` clause (5) naming the failure mode in the guard's own output.

### 5. `F1` — the blind recall audit does not fire

60 prose lines drawn with **seed 122 and sealed in the pre-registration commit, before the enumerator
existed**, adjudicated blind. Auditor booked **36 ASSERTION**; the census placed 24 in tier 1, 10 in
tier 2, and **2 in neither** → **enumeration recall 34/36 = 94.4 %** against a **≥ 90 %** floor; miss
rate 5.6 % against a **> 10 %** firing threshold.

Both misses are `UNCITED PLAIN`, which **sharpens** the FLOOR disclosure: **tier-2 recall against the
auditor's own uncited class is 10/13 = 76.9 %.** Precision (measured, not pre-registered): 1 of 24
NOT-ASSERTION rows booked = 4.2 %.

### 6. The whole-suite run found two more provenance defects — in the fence this very iter shipped

Third occurrence of one class, and this time inside the iter that argues for whole-suite runs:
`claim_census_guard` **stamped no tree at all** on direct execution, and its `--json` printed a bare
document followed by the text summary on the same stream — **unparseable stdout**, the exact regression
`fence_provenance` exists to fence. Green under every scoped run this iter took. Fixed on both halves;
the ratchet verdict now rides *inside* the document. A disclosure regression was caught while fixing it:
`KNOWN_WEAKNESS` was gated on `--census`, which is **not the verb `guard_family` uses** — a guard whose
qualifier is invisible in the family run is a guard whose green over-claims.

### 7. What the census establishes that a reading cannot

1. **A denominator** — 2,603 cited claims / 2,485 claiming units / 3,292 assertion candidates. Every `P`
   this milestone published was a numerator with nothing behind it.
2. **That the milestone's own published error rates were sample artifacts**, by about nineteen-fold.
3. **Exhaustive coverage of a named class** — 525 of 525, with the 41 it could not resolve **named**
   rather than guessed.
4. **A defect class no fence could see** — 1,151 unevidenced assertions, hedge discipline at 24/1,175.
5. **A ratchet** — the per-file count cannot rise without a guard going RED. The only mechanism in this
   milestone that acts on tier 2 at all.

**What it does NOT establish:** that the corpus is correct (a `SUPPORTS` verdict is about the citation,
not the world); that the residual pool is small (**1,908** tier-1 pairs unadjudicated, **1,151** tier-2
assertions unverified — the floor is now *larger* and *better named*); and **nothing whatever about
clause 5**.

### 8. The whole-suite claim, with the invocation and the expected wall time (§5 rule 51)

```
cd .agentspace/rosetta-extensions/stack-core
STACKCORE_PROGRESS_LOG=/tmp/m257x-iter122-beacon2.log \
  /usr/bin/python3 -m pytest tests/ -q --tb=line -p no:cacheprovider --no-header --durations=5
      ->  1 failed · 1146 passed  in  1055.54 s (17:35)   [rext 15b9454, corpus at iter-122]
```

`/usr/bin/python3` is **3.9.6**, the only interpreter on this host with pytest. The single failure is
`test_claim_twin_guard_iter48_answer_key::test_02` — **the standing, documented RED**, re-attested by a
full run rather than carried. **1,146 passed vs iter-121's 1,125 = +21**, which is exactly this iter's
control count, and **1055.54 s reconciles with iter-121's 1032.57 s to within 2.2 %** — the check rule 51
asks for.

**An EARLIER run of this suite is disclosed as CONFOUNDED and is not quoted as a result:** 3 failed /
1143 passed / 1128.21 s, taken while the tree was edited mid-run, exactly as iter-111 and iter-121
disclosed theirs. **It nonetheless did its job** — its two non-standing failures are §6's provenance
defects, and the run above is the clean one taken after they were fixed.

**Guard family: 22 members · 18 GREEN · 0 RED · 4 not-run** (`anchor_offset_guard`, `repair_leak_guard`,
`repair_reach_guard`, `value_change_guard` — commit- and input-scoped, no `--range`/`--ledger` supplied).
**Not a whole-family green, and the runner's own summary line says so.** The family grew 21 → 22 with
`claim_census_guard`, GREEN on its ratchet. Invocation: `guard_family.py --repo-root <rosetta> --platform
<rosetta>/stack-demo/platform`.

**Phase 5 grading:** (1) gate-met: **n — 4 of 5, unchanged** — (2) triggered-tok: n (**a successor
strategy remains FORBIDDEN by `TOK-08`'s sealed rule; the census is an INSTRUMENT, not a strategy, and
`F4` books any sentence treating it as the grader as a defect of this iter**) — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome:
**exit-3**
**Decisions:** `D-M257x-122-1` … `D-M257x-122-6` (see [`decisions.md`](decisions.md))
- iter-124 (tik): tier 2 triaged over its consequence class — cite 96.2 % printed / ≈86.6 % audited, fix a floor of 4; **the corpus is under-cited, not unfounded**. 27 false sentences repaired across 14 files, all one predicate iter-123 had measured and whose correction had reached 2 sites — see `iter-124/progress.md`
- iter-125 (tik): the directus register entry re-derived at source — 3 of 4 claims verbatim, the environment inventory WRONG and corrected, plus a net-new second exposure (`KEY` in clear in the task definition); the AKB contradiction placed in the three sites a reader meets it and filed with an owner; AKB's correctness on the router residue recorded — see `iter-125/progress.md`
- iter-126 (tik): `platform_alignment_guard` off exit 2 — `unclonable` split from `unresolvable`, disclosed in the verdict, gated by an anti-vacuity control; the widening caught a mis-bound citation the exit 2 was masking. Re-pin backlog ENUMERATED (89 → 7 at the census ref, reproducing iter-123; 6 closed, 1 refused) — and reading it at HEAD had inflated it 3.1× — see `iter-126/progress.md`
- iter-127 (tik): the stale thing was the VERDICT not the hedge token — 5 sites (2 fenced tables) still called `cms`'s prod state unsettled four days after it was measured; all repaired. The guard's 9/13 note did NOT move and is reported unchanged; the routed item stays open — see `iter-127/progress.md`
- iter-128 (tik): the three diseases treated separately — **Priority 1**: the consequence class read **exhaustively (340/340, every clone at a ref)** → **13 false claims repaired at 30 sites in 17 files**, 4 of them predicates swept corpus-wide; headline is the **3rd security-surface UNDERSTATEMENT of the milestone** — the REST surface has **11** Echo groups not 6, and **one has no authentication at all**, in a paragraph already corrected twice that both times re-derived from the file the wrong sentence named (→ `§5` **rule 57**). **Priority 2**: the 820 triaged by the same committed predicate, printed `cite` 99.4 % **corrected to ≈ 90.0 %** by an audit on *this* population, `fix = 0` published as **a floor because the complement was never read**. **Priority 3**: suite **counts attested and cross-checked** (1,158 collected, identical to run 80's independent run; 1 failed = the standing RED) and **timing declared unavailable** (+48.7 % vs baseline) with both contaminants measured — **my own six agents listed first**. `state.md` inside every budget it can meet; the one it cannot is disclosed with its number. `repair_postcondition` rejected the Priority-1 commit once and was satisfied by re-measuring, not silencing — see `iter-128/progress.md`
- iter-129 (tik): **the alarm was too loud, and the census was too narrow.** **Priority 1** — `/api/invitations` SETTLED at `app` `ad9f3c498` by reading the `RegisterRoutes` call site and the manager instead of the mount comment: `cors` really is the only middleware and Clerk's absence is deliberate (pre-login pages), **but a 256-bit HMAC-derived token IS required and IS checked before anything is returned** (`invite.go:159`, `:194`; a miss → `404 not_found`), the mechanism being a **stored bearer capability, not a re-verified signature** (`ValidateToken` is called by nothing outside its own test), with the source's own words *"token possession is the authorization"*. **Corrected to *token-authenticated, deliberately pre-login*; no defect filed.** Rule 57 bit its own worked example twice more: *eleven* is right **for groups** but `app` mounts **7 more routes on the root in no group** (two open by design), and the six-group predicate's **4th site** (`architecture_overview.md:405`) survived run 81's own sweep. **Priority 2** — the out-of-census surface enumerated with the census's OWN instrument: **39.7 % reach**, **724 unevidenced consequence assertions outside it, all 724 read**. Found the milestone's core drift class *in the ops docs*: **39 stale `jobsimulation.*`/`skillpath.*`/`local_*` tokens in 8 files**, two RUNNABLE (a `psql` block that 42P01s; a bring-up probe floor on a dropped relation) — repaired, mirror-writes **retracted not re-qualified**; plus a **15th site** of the refuted `@anthropos.work`-only predicate in `run_guide.md`, `CLAUDE.md`'s rext list **missing 2 of 11 sections**, the M23 Directus cutover naming a **deleted service at 7 sites**, and `toolchain_overview.md` naming **3 schemas no stack creates**. **The two accountings are kept apart: clause 5 is NOT re-cut and none of this is added to it (`F4`).** The sweep also found a rext defect I could fix and did: **the `$HOME/.aws/credentials` removal was keyed on the DELETED `"jobsimulation"` on the DEV path** (fixed on the demo path at iter-88), so a `dev-N` `backend` kept the operator's real AWS credentials mounted beside the hardcoded prod `STORAGE_S3_BUCKET` — now derived, + 5 regression tests with a mutation control. **Priority 3** — R3 re-audited **on the complement**: **28/30 = 93.3 %** (Wilson 78.7–98.2), both disagreements `cite→hedge`, both `chronos`; iter-124's imported 100 % overstated `cite` by ~4 points (**90.0 → 86.1 %**, `hedge` roughly doubles). The 820 complement was read **EXHAUSTIVELY** for falsity — **`fix` = 46/820 = 5.6 %**, no longer 0-because-unread — and its headline is a defect **in the fenced map**: `platform-migration-status.md`'s `ai` row still said `library | library` after the 2026-08-04 fold into `app` (`1e457fa70`; `app/go.mod` requires it no more, `go.sum` 0 hits). `shared_libraries.md` recorded that at **iter-102** and the neighbouring `authn` row is correct — **the `ai` row is the only library row that repair never reached**, and the map's guard could not have caught it because `ai` is a module, not a `repos.yml` clone. Rule 54, inside the artifact clause 3 is graded on. **Priority 4** — iter-128's *"no owner exists"* probe **searched two files and four of its seven items had owners**; three genuine orphans moved to `roadmap-vision.md`, M255's provenance to its own milestone, and the body budget **raised once against a measurement** (12,000→13,500; frontmatter 2,600→1,860 so the pair sums **exactly** to the 15,360 file cap, which the old triple did not) with a **re-raise guard** in the contract. Guards **18 GREEN / 0 RED** before and after; census **1,150 unchanged**; suite **not re-run** — load was checked FIRST and the external project was live at ~6 of 12 cores, and *re-running an instrument under the condition that invalidated it is not a second measurement*. **Gate unchanged at 4 of 5; no reading taken.** See `iter-129/progress.md`
- iter-130 (tik): **the routed set closed, and the library rows given the fence that would have caught the `ai` row.** **Priority 1** — `FIX-M257x-iter129-sweep-residual` closed, every finding re-verified at a ref (an upheld claim counted as a result: `demopatch-spec.md`'s 23-patch inventory held exactly). Headline is the milestone's **3rd security-surface misstatement**: `latency-budget.md` credited the fake Clerk FAPI with **validating `redirect_url`** — `clerk-frontend/server.go:414-423` redirects **verbatim**, an open redirect that mints a session with no credential; that is the *designed* disarm and the control is `safety.md` §3 + tailnet scope, so what was false was the claim of a check in front of it. Plus: *"the ent privacy layer is unaffected"* refuted by **30 `OrganizationMixin` schemas**; a **second Atlas pipeline owned by `app`** (`68272003`) that 4 sites denied; `CORS_EXTRA_ORIGINS` documented as unlanded at **6** sites when it landed in May; **password resets are Clerk's**, so blanking `BREVO_KEY` does not suppress them; the **Bunny provisioning path does not exist** (0 `BUNNY_RECORDING_*` in rext — widened to 39 hits, the number moved and the verdict did not), so the exhibit is blocked on **two** things and only one was disclosed; `secrets-spec.md`'s waived class **would have false-failed its own gate**. **The `@anthropos.work` predicate — width measured FIRST** (§5 rule 57): four regexes returned 28/11/41/31, union **63 lines/28 files → 10 false**, and **only 4 of the 10 carry the literal token** — the other 6 say *"internal learning portal"*, structurally invisible to every prior repair's regex, **which is why the count grew 4→14→15→16**. Refutation upgraded from *0 hits* to a mechanism: `d5875e34` *"replace @anthropos.work email gate with Clerk org-membership check"*. `anthropos-labs.md` is a **different subject**, in no clone set — **disclosed as UNCLONABLE, not repaired by analogy**. **Priority 3** — assertion **G** fences the map's LIBRARY rows against the **module graph**, the class A/B could never reach because *a library is a module, not a clone*; it fired on **3 real rows** at its first run, including **`ai`'s PROD cell, which iter-129 left wrong while repairing the local one** — rule 54 inside the row it cited as the rule-54 exemplar — and `authn` ×2, never previously caught. Ninth token `library-unimported` added; **the map now DECLARES which row classes each assertion reaches** (membership fenced for all, state fenced only for libraries). Controls: every mutant asserts it applied, a parser positive control that raises, a NOT-RUN disclosure, and a **meta-mutation that kills 4 tests**. Side-deliverable: the **front-door `README.md` tier table was false in all three rows** (CMS/Jobsimulation/Storage/Roadrunner listed as live Tier-1). Guards **18 GREEN/0 RED**; three went RED on **my own** edits mid-iter and were repaired by fixing the artifact (two citations re-pinned `:106`→`:118` at the construct, the census ratchet evidenced not loosened). Census 1141→1140. **No reading taken; no `N` movement claimed** — see `iter-130/progress.md`
- iter-131 (tik · `iter_shape: reading`): **THE READING — `P = 29` / `N = 47`.** `P` **ROSE** 22 → 29 (+32 %) and `N` 28 → 47 (+68 %) over a corpus that grew 9.7 % and absorbed eleven iters of repair (30 consequence sites, 13 C1 claims, 46 complement fixes, the `ai` row, the tier-1 re-pins). **The pre-registration barred the flattering reading of a fall and equally bars the despairing reading of a rise:** `P` is a FLOOR over a suspicion-selected sample, never a corpus error rate, and this reading **cannot** separate *new defects* from *the same pool sampled differently*. **What it DOES establish is the test-retest result, and it is sharper than the primary: the overlap with iter-119's 22 predicates is ~0** — two consecutive readings produced **almost disjoint** sets (iter-116→119 overlapped 13/37 = 35.1 %; 119→131 overlaps ~0). **And the metric turned out to be uncomputable as published**: `iter-119/adjudication.md` enumerates only 8 of its 22 predicates, so the overlap could be measured against just 9 — the milestone has quoted test-retest figures for three readings **without the substrate to compute them**. **Largest cluster: 19 of 80 blockers, six seats, one root cause** — the corpus says `infrastructure` *"has never been in any clone set"* and cms's prod state is *"NOT MEASURABLE"* at **11 sites** while **28 sites cite a READ of it @ `13c248e6`**; iter-123 cloned it and `org-repos.md:102` says **"cms M810: SETTLED — DESTROYED. It has now been read"**, a single read that settled **four** standing questions. **Rule 54 at scale — and `CLAUDE.md` publishes the retracted claim too**, so every agent starts from it. **Three of the 29 are defects in prose I wrote, two of them THIS RUN**: §1 never gained the `library-unimported` row while assertion C now says "nine" (P7), *"all three sites are the literal curl"* was false of one of its three (P19), and `architecture_overview.md:83` still lists `ai` among the imported modules my own rule-54 sweep missed (P5) — **so assertion G prints the true module set on every run while two corpus sentences contradict its own fence**. Bands **7 of 15**, six of the eight failures **high**; band #7 was deliberately cut DOWN to ≤5 and measured **10** (3rd consecutive failure, now a third of the pool) confirming `D-M257x-117-2`; band #5 **retired** after four consecutive failures; band #10 (induction) held at 3. Upheld **89.5 %**, reported twice, **identical** because `wrong-tree` = **0** for the seventh reading running despite the widest rext gap yet (33 commits) and a dirty `ant-academy` tree. **4 blockers left CANNOT-SETTLE, not laundered** — the root-mount count, disputed in three consecutive readings, where my own first three counts returned 0 from broken searches. ⚠ **Method deviation disclosed: the 200-subagent cap left ONE independent adjudicator; I adjudicated the other twelve seats myself while being the author of three upheld predicates** — routed for re-adjudication. Instrument verified on both sides of the copy, pre-registration sealed in its own commit **before any seat was dealt**, all 14 seats committed verbatim, no clone fetched mid-reading, **no repair taken inside the measuring pass**. **Gate 4 of 5; clause 5 NOT met and NOT re-cut** — see `iter-131/progress.md`
- iter-132 (tik): **the hedge the read retired, swept — and the premise turned out to be settleable, not just re-wordable.** iter-131's **P1**, the milestone's largest measured cluster (19 of 80 blockers, six seats), **repaired at 15 sites in 8 files**. Width measured first (§5 rule 57): four independent searches → a 22-line union in 11 files, of which **only 15 are the predicate** — the rest are jobsimulation's GitHub *archive* state, the router banners' Vercel line, already-corrected sites and two other repos. The eight false sites inferred *"cms's prod state is UNMEASURABLE"* from *"`infrastructure` has never been in any clone set"*; **both conjuncts are true and the inference is not**, which `adj-1` — the one independent adjudicator — had already corrected the coordinator on. **The seven OTHER sites hedged a different proposition on the same premise** (the production RPC address) and were not in the route at all. **Rather than write the honest hedge, iter-132 spent one `--depth 1 git clone`: `HEAD` = `13c248e6`, the exact sha the corpus already cites 28 times.** **Production DOES name `http://backend.internal.anthropos:8081` — exactly once**, as `module "backend_euwest1"`'s `cms_rpc_address` input (`terraform/production/services.tf:346` + `locals.tf:22`), confirmed by two independent searches: **M809 landed in production too, in the same shape as locally**, there is no `skiller_rpc_address` or `jobsimulation_rpc_address` in prod terraform at all (so `skiller.md`'s old flat assertion was wrong about the *variable*, not only the tense), and **`org-repos.md`'s 666 lines / 9 service modules / surviving `module "jobsimulation_euwest1"` all held exactly** (upheld claims counted as results). **Two riders, both the platform contradicting its own prose:** `services.tf:352-355` explains an absent input by naming `module.messenger_euwest1` while `:618-621` of the same file says that module is deleted, and **three `infrastructure` narrative docs describe that deleted module wiring four RPC addresses** — §6's *config is the documentation of record* met in the wild. **Then the fence fired on the repair, and it was right twice and blind once:** the NOTE went **9 → 11** hedges, so it was instrumented rather than argued — **8 of the 11 also carried a ref-pinned reading**, i.e. they are **retractions** the substring matcher cannot distinguish from live hedges (`architecture_overview.md:227` and `platform-migration-status.md:93`, the corpus's own model retractions, were being counted against it); **2 of the remaining 3 were paragraphs iter-132 had just written**, asserting a module deletion with no sha in the paragraph — **fixed in the prose, not the fence**. The blindness is fixed in the **instrument**: a disclosed third bucket (`hedged`/`mixed`/`measured`) + a `KNOWN_WEAKNESS` line, the NOTE now firing on live hedges only — **9 → 1**, the survivor being the protocol doc's own worked example. **The rejected alternative is named: re-wording until the substring vanished would have improved the number and nothing else** (`D-M257x-122-3`'s class). 3 tests added, **meta-mutation kills 2 of 3 — and the survivor is RENAMED off `test_MUTATION_…`** rather than left to look like a control; the real-corpus floor re-cut from one bucket to the sum, because a floor pinned to `hedged` would argue against the repair it exists to enable. §5 gains **rule 61** (*"not in the clone set" is a fact about our habits; it never entailed "not measurable"*). **The inherited route was also over-stated by one file and it is reported, not dropped:** it named `CLAUDE.md`, which iter-124 had already corrected. **Whole suite CLEAN: 1 failed / 1208 passed in 2077.16 s** — the 1 is the standing documented RED, re-attested not carried; **run 1 is disclosed as CONFOUNDED** (edited mid-run) and its second failure **did not reproduce**, proven by re-running that test alone on the committed tree. **Rule 51's timing leg FAILS and is reported as a failure**: +96.8 % against baseline, and the CLEAN run was 26 % slower than the contaminated one within the same hour — this host is not a stable timing substrate and no wall-time claim from it is a measurement (routed). **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5** — see `iter-132/progress.md`
- iter-133 (tik): **the route said three anchors; the width search said ten, and the width search won.** `FIX-M257x-iter131-my-three` closed — **and P5 was a predicate, not an anchor**. Two independent searches (rule 57) found the private-Go-module set misstated at **8 sites in 8 files**, wrong in **two opposite directions**: some still list **`ai`** (folded in-tree at `1e457fa70`, `app/go.mod` requires it no more), others say **"three — colony, proto, taxonomy"**, dropping **`analytics-go`** and **`storage`**, *both direct requires*. **Ground truth measured at source, not quoted:** `app` `ad9f3c498` `go.mod:14-18` = **`analytics-go`, `colony`, `proto`, `storage`, `taxonomy`** — five, all direct, no `// indirect`, no org `replace`; `sentinel` `f2c4619` adds no sixth (`taxonomy` there is `// indirect`). **The root is a conflation with a cardinality coincidence**: the **five historical shared libraries** (`colony`, `authn`, `proto`, `ai`, `taxonomy` — `shared_libraries.md`'s subjects) and the **five imported private modules** are *different sets that overlap in three and share only a number*, which is exactly why nobody noticed — **both counts were right**. `CLAUDE.md` already warns *"do not read this list as `app`'s dependency set"* and **nine sentences in eight other files did exactly that**; every repaired site now names **which five it means in the sentence itself**, because fixing the counts alone leaves the conflation to re-drift (the org-module block shrank **7 → 5** inside this milestone). `external_services.md:554` was **already correct and left alone** — an upheld claim counted as a result. **P7 closed:** §1 gains the missing **`library-unimported`** row — the guard's `ALLOWED_STATES` has had **nine** since iter-130 and §1 defined **eight**, so `platform_alignment_guard` was **GREEN for three iterations over a document that did not define what it enforced**; the row now says so, and this is **iter-131's lesson 1 running in mirror image** (there, a fence printed the right answer while the prose beside it was wrong; here, the prose a fence implements fell behind the fence — **neither direction is caught by running the fence**). **P19 closed at TWO of three, not three:** `staging-bringup.md:528` is a **prose bullet** (*"Quirk #13 … Bypass with `POST /v1/sign_in_tokens`"*) carrying neither `curl` nor the host, so the over-claim sat inside the sentence whose entire purpose was making a citation drift-proof — the robust re-derivation is restated as the **shared substring** that actually returns all three. Guards **18 GREEN / 0 RED / 4 not-run**; scoped fence suites **123 passed**; **the whole suite was NOT re-run and §5 rule 60 requires saying so out loud** — iter-132 ran it clean ~40 min earlier on **the same rext tree** (`223e4a6`), and iter-133 modified **zero** rext files, so the exposure is bounded and **stated as a gap rather than characterised as covered**. **Zero rext changes, deliberately: the fence was right and the prose was the defect.** **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5** — see `iter-133/progress.md`
- iter-134 (tik · `iter_shape: audit`): **the conjecture was wrong, and being wrong is the useful part.** `FIX-M257x-iter132-marker-fences-cannot-see-retractions` claimed the retraction blindness *"plausibly affects every marker-matching fence"* — acting on that as written is four fence fixes; **checking it cost one probe, and the branch was stated before the probe ran** (≥2 blind → pattern; ≤1 → refuted). **Measured: 1 of 4, and it is the one already fixed. REFUTED.** The four prose-classifying fences were **imported and interrogated, not grepped** (rule 22 — grepping `retract` hits all four, including comments *discussing* retraction, and returns the opposite verdict): `claim_twin_guard` **5/5** mechanisms, `claim_census_guard` **2/5**, `platform_predicate_guard` **1/5**, `unreadable_repo_claim_guard` **0/5** pre-iter-132. **The partial scores are coverage, not gaps, and the audit checked that before saying so** — `claim_census_guard` measures *unevidenced, never false*, so a ref-carrying retraction never enters its numerator; `platform_predicate_guard` grades prose against config and carries an explicit `not|never|no longer|formerly|was` exclusion (`:900`), a historical-vocabulary section (`:160`) and historical-sha handling (`:1092`) — for what it measures, the past-tense axis IS the retraction axis. **The audit's real product is the REUSE GAP it exposed: `claim_twin_guard` has shipped `RETRACTION_MARKERS` (14 tokens), `_looks_retracted`, a 320-char window and a waiver file that DECAYS if the retraction is deleted since iter-48 — with a docstring naming the exact hazard — and iter-132 met that hazard and built a coarser bespoke bucket beside it, in a fence the same runner loads.** iter-132 booked its independent arrival at `D-M257x-121-4` as a virtue; it is one, and **the same paragraph should say the family already held a tested, sharper answer** — a milestone that reports only the flattering half of its own pattern-matching is doing the thing it exists to stop. **The refactor is deliberately NOT taken**: sharing the predicate is a structural choice (cross-fence import vs shared module) and all eight vacuous fences on this milestone's record came from building under pressure — routed with both options named. Guards **18 GREEN / 0 RED / 4 not-run**; **zero files changed in `corpus/` and `rosetta-extensions`** (an audit), so no code-test gate applies and none is claimed; **whole suite not re-run and rule 60 requires saying so** — nothing executable has changed since iter-132's clean run. **⚠️ `FIX-M257x-iter131-adjudication-independence` is now UNTOUCHED for three consecutive iters** and is the oldest unactioned route on the milestone, named here rather than left to lapse quietly. **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5** — see `iter-134/progress.md`
- iter-135 (tik · `iter_shape: adjudication`): **the milestone's oldest route closed, and the disclosed deviation ran the OPPOSITE way to its own prediction.** The twelve seats iter-131's coordinator adjudicated *itself* were dealt to **six independent adjudicators** (B–G, one per seat-letter, `adjudicator-brief.md` unmodified, hard bars enforced: no `knowledge/plan/**` beyond the brief + their own two seats, no other adjudicator's output, read-only). **B–G: 60 claimed / 57 upheld / 3 rejected / 0 cannot-settle / 0 wrong-tree = 95.0 %**; with `adj-1`, **all 14 seats independently adjudicated: 80 / 76 / 4 / 0 / 0 = 95.0 %** — against the coordinator's own **89.5 %** (68 upheld, 8 rejected, **4 cannot-settle**). **iter-135 predicted self-adjudication had INFLATED the reading; measured, independents uphold MORE, reject half as many, and return zero `wrong-tree` for the eighth consecutive reading — the self-adjudication was CONSERVATIVE, and the hypothesis is refuted.** **All FOUR CANNOT-SETTLE blockers were SETTLED by opening the evidence**: `adj-C` located a genuine `infrastructure` clone at exactly `13c248e6` and read it (zero `module "cms"` org-wide, positive control 12 `module "` hits in the same file), and `adj-D`/`adj-F`/`adj-G` converged on the **root-mount count** the coordinator called *"disputed in three consecutive readings"* — **it is 8**, closing `FIX-M257x-iter131-root-mount-count-underived`. **What independence bought was FRAMING, not verdicts** (per-anchor truth barely moved; all six returned predicate corrections): `adj-C` **independently reproduced `adj-1`'s premise-vs-inference correction** — *"that premise is TRUE; the falsehood is the inference"* — and then found **the same defect in the BRIEF's own example predicate**, i.e. the instrument was teaching the error it exists to catch; `adj-F` **upheld a seat's number while refuting its diagnosis** — the missing 8th root route is **not** the one the seat named (already in the table) but **`/v1/labs/:slug/workspace.tar.gz`**, *inside* `internal/web/backend/`, inverting the seat's scope argument, so *"a repair driven by the seat's report would fix nothing"* — and the route is the **4th security-surface understatement of the milestone**: its own comment says *"OUTSIDE the write group — OPTIONAL auth"*, wired unconditionally at `backend.go:301`, serving a tarball; `adj-E` **overturned a causal story that would have misdirected the remedy** — five mis-anchors were **correct when written** and rotted (+2/+3/+8/+14) from unrelated insertions *above* them, so the fix is *fence the form*, not *repair harder*, and all five sit in `corpus_citation_guard.py`'s **declared** bare-`:NN` blind spot while the brief's repair-induced test **structurally cannot** attribute them (`git log -L` misses an edit made above the citing line). **Two blind confirmations of repairs made without sight of them**: `adj-D` and `adj-F` both surfaced the private-module 3-vs-5 defect **iter-133 had already repaired**, and `adj-C` booked the M810 cluster `UPHELD (since-repaired)`. **`P`/`N` are deliberately NOT re-cut** (`D-M257x-135-1`) — re-grading an existing seat set is not a fresh sample, and a milestone that re-grades its own old sample whenever it improves the method can walk `P` anywhere it likes. **The one added instruction is disclosed as load-bearing** (`D-M257x-135-2`): without *"grade as booked; mark `UPHELD (since-repaired)`"*, every claim iters 132–133 correctly repaired would have been REJECTED for no longer being visible — scoring the repair as a refutation; the since-repaired set is a minority, itemised per sheet, so the rate is recomputable without it. Guards **18 GREEN / 0 RED / 4 not-run**; **zero `corpus/**` and zero rext files changed** (an adjudication), so no code-test gate applies and none is claimed; **whole suite not re-run and rule 60 requires saying so**. **A large, adjudicated, cited work list of still-live blockers is routed as `FIX-M257x-iter135-adjudicated-live-defects` — a work list, not a sample. Gate unchanged at 4 of 5** — see `iter-135/progress.md`
- iter-136 (tik): **the 4th security-surface understatement corrected, and the instrument that taught the error fixed.** Two items taken off iter-135's work list **by CONSEQUENCE** (`TOK-08`'s carried finding), the rest left routed rather than half-done. **`app` mounts EIGHT routes on the root Echo, not seven** — the eighth is `/v1/labs/:slug/workspace.tar.gz`, serving a **workspace tarball** from outside the `/v1/labs` write group and therefore outside its `apiKeyAuthMiddleware(…, "labs:write")`, **wired unconditionally** at `backend.go:301` (no `colony.Development` guard, no flag). **Re-derived at source (`labs_admin.go:31-41` @ `app` `ad9f3c498`) rather than taken from the adjudicator — because the SEAT that first reported the miscount named the WRONG route** (`/ai-readiness/unsubscribe/:token`, which the table already contained), so *"a repair driven by the seat's report would fix nothing"*; **three independent adjudicators converged on the number and only one located the route.** → `D-M257x-136-1`: **a count claim is graded by re-enumerating, never by accepting the reporter's candidate** — the third time on this milestone a correct FINDING arrived with a wrong DIAGNOSIS attached. **Stated without over-claiming, deliberately:** it is *not* "an endpoint with no authentication" — the handler's own comment says a public Lab's workspace is public by design and a tenant-private one requires a key, i.e. **the tenancy decision moved inside the handler**, which is precisely what a group-level sentence cannot describe; so the repair ships an **enumeration of all eight** (2 open by design · 2 dev-only · 3 self-authenticating in-handler · 1 optional-auth) rather than an adjective — `D-M257x-121-2` records this milestone publishing a *new* absolute quantifier over a security surface inside a repair whose subject was absolute quantifiers. Repaired at `security_compliance.md:250` + its closing summary and `architecture_overview.md:406` (which is **now recorded as having been wrong about two different counts** — *"6 Echo groups"* until run 82, *"seven"* root mounts until now). **And the compounding one: `iter-131/adjudicator-brief.md:78`'s example predicate read *"cms's production ECS state is unmeasurable **/** infrastructure was never in a clone set"* — joining a FALSE proposition to a TRUE one with a slash and inviting adjudicators to book the conjunction.** `adj-1` (iter-131) and `adj-C` (iter-135), two blind adjudicators one reading apart, made the **same** premise-vs-inference correction and `adj-C` traced it to that line — **an instrument that models a conflated predicate teaches the error it exists to catch.** Corrected to the causal form + the rule: *state predicates so every conjunct is independently false, or split them.* Guards **18 GREEN / 0 RED / 4 not-run**; zero rext files changed; **whole suite not re-run and rule 60 requires saying so**. **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5** — see `iter-136/progress.md`
- iter-137 (tik): **roadrunner was wrong in TWO OPPOSITE DIRECTIONS at once, and both were live.** Two iter-135 adjudicators, blind to each other, hit the same subject from contradictory sides: `adj-F` P2 — *"retirement is unresolved; prod terraform still reads `service_desired_count = 1`"* (**says it is running in production**) — and `adj-B` P-2 — *"`roadrunner` is one of EIGHT domains folded into `app`"* (**says it was absorbed into the monolith**). A reader could open two files of this corpus and come away with mutually exclusive pictures, each stated confidently. **Ground truth re-derived at source under positive controls, never taken from the adjudicators** (`D-M257x-136-1`): **`app/internal/roadrunner/` exists at no ref and was NEVER added** — `git log --all --diff-filter=A` → **0 commits, ever**, in a full non-shallow 6,728-ref clone at `app` `ad9f3c498` (control: `jobsimwiring` → 3 paths); the seven that WERE folded all have packages; roadrunner's job is done **inside the jobsimulation domain** (`app/internal/jobsimwiring/wiring.go:123`, whose own comment says it *"replaces the removed roadrunner RPC edge"* — **the platform's own word is REMOVED**). **And production: `infrastructure` cloned and read at `13c248e6` (closing `FIX-M257x-iter132-infrastructure-is-cloneable-so-clone-it`) — `terraform/production/services.tf` declares EXACTLY TEN service modules and `module "roadrunner"` is not among them**; `roadrunner/terraform/main.tf` is 95 lines of module fed from **unbound `var.*`**, so `:19` is an input to something nothing instantiates — **the identical orphaned-dead-code class as `cms`, `messenger` and `graphql-wundergraph`, and § 3 was written at iter-123 to close exactly it, reached those three, and skipped roadrunner ONE ROW AWAY IN THE SAME TABLE.** `adj-F` found its sharpest form: `architecture_overview.md`'s **CMS row carries the full correction and the Roadrunner row one line below does not.** **The read also supplied the POSITIVE fact the corpus never had** — the `production_roadrunner_judge0_*` secrets feed `TF_VAR_judge0_*` (`infrastructure/.github/workflows/wf-terraform-deploy.yml:209-211`) into **`module "backend_euwest1"`** (`infrastructure/terraform/production/services.tf:384-385`): **production wiring Judge0 straight into `backend` under roadrunner-named keys IS the fold, at the config layer**, far better evidence than the count the corpus was reading; `infrastructure/knowledge/service-dependencies.md:119` says it in the platform's words (*"Judge0 … called directly now; `roadrunner` is off this path"*). **Repaired at 29 sites in 15 files**, both conjuncts, starting with `CLAUDE.md` (4 sites — the file every session loads). **Width measured first (§5 rule 57) with four independent searches — and the verification sweep still found THREE survivors**, which is `D-M257x-137-4`: two the planning vocabulary could not reach (a gRPC-hops paragraph; an index row) and **one no line-oriented instrument can reach at all — a Roadrunner bullet under a heading reading *"Domains inside Backend/App, not services"*, where the false predicate is asserted ONE LEVEL UP and inherited.** Grep, the anchor fences and the claim census all read *lines*; §5 rule 62 gains **(a′) then read the HEADINGS**. **The fences fired on my own repair and were right:** `anchor_construct_guard` + `repair_postcondition` went RED because `roadrunner.md` carried a bare `:NN` pin **as its own worked example of a bad pin**, and the repair shifted the file until that quoted pin hit a blank line — *the citation that rotted was inside the sentence warning about rotting citations*. **Fixed by DELETING the pin, not re-pinning it** (`D-M257x-137-3`; re-pinning restarts the clock) — the anchor-axis sibling of iter-132's hedge-marker blindness and iter-134's 1-of-4 measurement. `platform_alignment_guard` then went **exit 2** on two citations *I* had just written with **bare heads**, correctly classing them `unresolvable` (the citation's defect) rather than `unclonable` (the substrate's) — iter-126's split doing its job on its author; repo-qualified and green. **Upheld claims counted as results:** `org-repos.md:143`'s 7-hit `infrastructure` measurement **re-derived independently and upheld byte-for-byte**, and the fenced map's roadrunner row `:91` was **right all along** (only its *"a repo this map has never read"* clause was stale — repaired, and it is `adj-F`'s flagged site). **The 31 `roadrunner`-as-domain hits in `rosetta-extensions` are ALL frozen `repair_leak` test fixtures — deliberately NOT repaired, since editing them would corrupt what the fence measures.** **Side-deliverable, separate commit, does NOT grade the iter:** `corpus/README.md:18` — the **16th escape of the cms-M810 predicate, on the corpus front door**, still reading *"M810 … is uneven … not moved for cms"* four days after iter-123 measured it and after **two** corpus-wide sweeps (iter-127: 5 sites; iter-132: 15); width re-measured first — **1 live site**, every other match a retraction quoting the old wording. Rule 55 exactly. Guards **18 GREEN / 0 RED / 4 not-run** (and the runner's own line says that is NOT a whole-family green); scoped fence suites **367 passed / 0 failed** in 623.05 s; **whole suite NOT re-run and §5 rule 60 requires saying so** — zero `rosetta-extensions` files changed, iter-132's clean run stands on the same rext tree (`223e4a6`); **suite wall-time deliberately not quoted as a measurement**. **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5** — see `iter-137/progress.md`
- iter-138 (tik): **the anchors rotted, and a machine can see it after all.** `TOK-08` class 1 (intra-corpus citation resolution). **Priority 1** — the adjudicated citation set repaired at **9 sites in 6 files**, every target **re-derived by opening it** (`D-M257x-136-1`) and **7 of 7 adjudicator candidates held**, which is booked because iter-136's did not: `analytics-go` is wired at **`main.go:494-495`** (the file's only two `trackingManager` lines) not `:507-508` (the storage-in-app comment block, a *different* construct at all four refs the corpus reads) — repaired at `shared_libraries.md:77` **and `CLAUDE.md:288`**, while the `handler.go:302-316` half was **opened and is exact** (upheld claim counted as a result); the *"only"*-quantifier precedent lives in `clerk-integration.md`'s **Sign-in tokens** bullet, not `:40` (org invitations, no quantifier); and `adj-E`'s five rotted anchors all confirmed (`academy-backend.md` ×2 → Certificate-minting and the store.js/beacon blockquote; `graphql-wundergraph.md` ×3 → the no-hot-reload bullet, the Federation bullet, the profiles-warning opening line). **Every repair NAMES the construct and drops the number** (`D-M257x-138-3`) — because re-pinning restarts the clock: **that `5050` pointer has now rotted TWICE** (`:174-176` → iter-98 → `:193` → iter-138) and **one paragraph of `graphql-wundergraph.md` held three rotted pins**. **Priority 2 — the class MEASURED, not assumed.** `adj-E`'s reframing is that all five were **correct when written** and rotted from unrelated insertions *above* the target, so *"repair harder"* has no target; and `corpus_citation_guard.py` **declares this blind spot in its own docstring**, excluding bare `:NN` pins *"as not mechanically decidable."* **That is true of the CLAIM and false of ROT** (`D-M257x-138-1`) — rot needs no sentence read, only git: if the text that stood at the cited line when the citing line was authored now stands elsewhere, the pin rotted and the new line is the repair. **Branch pre-registered in `overview.md` before the probe ran** (≥5 → route a fence; ≤4 → refute). **Measured: 588 bare pins → 222 DECIDABLE → 127 ROTTED = 57.2 % — ⚠️ RETRACTED AT ITER-139, 0-for-12 on a stratified audit (Wilson95 [0.0, 24.3]); the repairs stand, the number does not**, against a positive control of **95 STABLE**, and **every delta is positive** (`+1` ×22 … `+135`) — `adj-E`'s mechanism reproduced at corpus scale. **Published with its floor** (iter-114): 127 is a claim about the **222**, never the 588; the undecidable buckets are disclosed — `out-of-range-then` **241** (largely **cross-file continuation pins**, which a same-file probe cannot distinguish and which a fence must resolve first), too-short 109, target-gone 16. **The fence is ROUTED, not built here** (`D-M257x-138-4`): tooling belongs in rext, and **all eight vacuous fences on this milestone's record were built under pressure at the end of an iter**; the probe's `STABLE=95` controls the *probe*, not the fence. `FIX-M257x-iter135-bare-pin-blind-spot` **CLOSES as measured**, superseded by `FIX-M257x-iter138-anchor-rot-fence`, which now has a denominator. **And a genuine RED, which was iter-137's:** `test_anchor_offset_guard.py::TestAntiVacuity` failed on `dependency_map.md:9` — a line iter-137 wrote one commit earlier, citing `services/README.md:39` when **six files here are named `README.md`**; fully qualified, `_ambiguous` → `[]`. **It escaped BOTH of iter-137's gates in the same direction** — `anchor_offset_guard` was **NOT-RUN** in the family (commit-scoped, no `--range`) *and* absent from its nine scoped suites, which were chosen **by topic when what that iter had rewritten, at 29 sites, was anchors** (`D-M257x-138-5`, → `§5` rule 63(d)): **a disclosed not-run bucket is not coverage.** Booked as the loop working too — the escape survived exactly one iter, was caught by the anti-vacuity control built for it, and cost one line. `§5` gains **rule 63** (an exclusion is only as narrow as the predicate that justified it; publish the undecidable buckets; name the construct, never re-pin; choose suites by what you changed). Guards **18 GREEN / 0 RED / 4 not-run**; scoped suites **102 passed / 0 failed** after the RED; **whole suite NOT re-run and rule 60 requires saying so** — zero rext files changed. **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5** — see `iter-138/progress.md`
- iter-139 (tik · `iter_shape: audit`): **the census was wrong, and the audit that found it was pre-registered.** iter-138 published **`127 rotted / 222 decidable = 57.2 %`** and routed a fence to be built against that baseline — **from a probe that had never been audited.** iter-139 audited it: **strata and selection rule sealed in `overview.md` before the first case was opened** (4 cases at `|Δ| ≥ 100`, 4 at `10 ≤ |Δ| < 100`, 4 at `|Δ| < 10`, each taken **in census order from the top of its stratum**, so the sample is reproducible and not cherry-picked). **Result: 12 of 12 are FALSE POSITIVES — precision 0.0 %, Wilson95 [0.0, 24.3].** The cause is a form the probe never modelled: **in this corpus a bare `` `:NN` `` is overwhelmingly a cross-file CONTINUATION pin** (`` `app/main.go:15`, `:62`, `:63` ``; `` (`isThrottlingError`, `:129` / `:166` / `:325`) ``) or a **quoted / historical / negated** pin (*"two of iter-45's five defects are relationships between line numbers — `:788` citing `:447`"*; *"there is **no** `:29` declaration at all"*) — **very rarely the same-file self-citation the probe was measuring.** **`127 / 57.2 %` is WITHDRAWN — not re-qualified, not narrowed, not restated with a caveat** (`D-M257x-122-3`'s class) — and **retracted IN PLACE at all three sites that published it** (rule 54): a banner above the retained numbers in `iter-138/progress.md`, this ledger's iter-138 entry, and `§5` **rule 63**, whose figures were replaced and whose **rules (a) and (b) were rewritten**. **The finding is worth more than the number** (`D-M257x-139-2`): iter-138 **disclosed** its 241-case `out-of-range-then` bucket and **named its cause exactly right** — *"largely cross-file continuation pins"* — and **that honest disclosure is what made the remaining 222 look clean.** It was not: a continuation pin lands in `out-of-range` only when the cited number exceeds the **citing** file's length, and in a 3,100-line doc almost none does, so **the same failure mode passed undisclosed into the "decidable" set and dominated it.** → **a disclosed limitation is quarantined only if you show the boundary holds; naming a floor is not bounding it, and the disclosure made the number MORE persuasive, not less. Sample the clean bucket for the disease you just disclosed.** Corollary (`D-M257x-139-3`): iter-138's `D-M257x-138-1` — *an exclusion is only as narrow as the predicate that justified it* — **survives as a RULE and is withdrawn as an APPLICATION**: content and rot are blocked by the **same** unresolved thing, the pin's **head**, and `corpus_citation_guard.py`'s exclusion is therefore better founded than iter-138 credited; `adj-E`'s five genuine rotted anchors were found **by a human reading five sentences**, a rare form no machine reaches before solving head resolution. **What STANDS is stated so the retraction is not over-read** (`D-M257x-139-4`): **all 9 of iter-138's citation repairs hold** — each came from `adj-E`/`adj-D`'s hand-verified list and each was **re-derived by opening it**, none from the probe — as do `D-M257x-138-3` (strengthened: a head a purpose-built probe cannot resolve is a head a reader cannot resolve) and `D-M257x-138-5`. **`FIX-M257x-iter138-anchor-rot-fence` is RE-SPECIFIED, not cancelled** — first deliverable is **head resolution**, and it has **no baseline until that exists**; `anchor_construct_guard`'s existing resolver (868/1469, with a `bare-continuation` strategy at 235) is the thing to reuse, which gives `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` its **first concrete consumer**. **`FIX-M257x-iter138-127-rotted-pins` is WITHDRAWN — there is no such work list.** **Nothing downstream had consumed the wrong number**: no repair was driven by it, the fence was not yet built, so the error cost **one iter** and skipping the audit would have cost a fence with a fabricated baseline plus 127 unnecessary edits (`D-M257x-139-5` — the milestone books its loop in both directions). Guards **18 GREEN / 0 RED / 4 not-run**; scoped suites **102 passed / 0 failed**; **whole suite NOT re-run and rule 60 requires saying so** — zero rext files changed. **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5** — see `iter-139/progress.md`
- iter-140 (tik): **the receipts — and the class that IS censusable, which is the point.** `TOK-08` class work on **published receipts**: a backticked command carrying a claimed count. **Population enumerated before any verdict — 22 across 15 files**; **check rule sealed in `overview.md` first** (run verbatim as published, in listed order, **no re-wording to make it pass**; *not-checkable* published as its own count, never folded in — iter-139's lesson). **9 were checkable on this box at the ref each names; 7 reproduce, 2 do not.** ✅ `security_compliance.md:95` (`OrganizationMixin{}` → **30**), `:235` (`ValidateToken` @ `ad9f3c498` → **8**, and **both named files correct**), `dependency_map.md:59` (`SKILLER_STREAM` @ `ad9f3c49` → **6 lines**, and **all three named files correct**), `studio-room.md:367` (`mistral` @ `aeec036a` → **22 / 3**), `safety.md:1082` (`BUNNY_RECORDING` → **0**), `build-budget.md:87` (`BRINGUP_ANCHORS` → **2**), `ai_architecture.md:54` (`bedrock|boto3` → **0**). ❌ **`sentinel.md:5`** — `adj-B`'s P-1 — claims *"returns **one unrelated hit**"*; run verbatim it returns **0 lines, exit 1**, as does the widened whole-tree `git grep -in authoriz fa47850d`, **with the positive control holding in the identical form** (`colony` → hits in messenger's own `cmd/` package, three files) so the absence is real and the **conclusion is STRENGTHENED, not weakened**. ❌ **`latency-budget.md:365`** — **found by this census and by nothing else** — claims *"returns **one** non-test occurrence"*; verbatim it returns **22 lines across 4 files, 3 of them non-test, and exactly ONE of them CODE** (`server.go:414`); the other two are **comments** (`:150`, `:155`) describing the same handshake bounce, so the conclusion is intact. **Both repaired with the real breakdown.** **The two failures share one authoring shape** (`D-M257x-140-1`): **the number was written from the CONCLUSION, not from the command's output** — you know the answer, you write the command that demonstrates it, and you fill in the count from what you know. **The defect is not the number; it is that the receipt no longer demonstrates anything** — a reader who runs the grep sees 22 where the page says one, and rationally distrusts the whole paragraph, *including the parts that are exactly right*. **A receipt is the strongest form of citation precisely because it is checkable; one that fails its own check is weaker than no receipt at all.** **And the meta-result that closes iter-139's open question** (`D-M257x-140-2`): iter-138 censused bare `:NN` pins and iter-139 retracted it at **0/12**; iter-140 censused receipts and got **7/9**. Same strategy, same author, one iter apart, opposite outcomes — **the variable was never the strategy, it was the SUBJECT'S DECIDABILITY. A class is censusable iff an instance carries its own HEAD**: a receipt names its command, pathspec and ref; a continuation pin names nothing. **So iter-139 must not be read as "censusing does not work here"** — it works, on a subject that can resolve itself. **The 13 not-checkable receipts are published as their own count** (`D-M257x-140-3`), neither passes nor failures — reporting *"7 of 9"* without the 22 would be the same over-claim iter-139 retracted one iter earlier; routed as `FIX-M257x-iter140-receipts-not-checkable-here`, alongside a new `FIX-M257x-iter140-receipt-fence` (population 22, buildable once pathspec head-resolution exists — **the same prerequisite iter-139 named, arriving from the other side**). **`adj-B`'s P-1 CLOSES.** Guards **18 GREEN / 0 RED / 4 not-run**; scoped suites **102 passed / 0 failed**; **whole suite NOT re-run and rule 60 requires saying so** — zero rext files changed. **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5** — see `iter-140/progress.md`
- iter-141 (tik): **three cross-references closed, and a pointer that names a RETRACTED TITLE — a shape no anchor fence can see.** Closes the **cross-reference half** of `FIX-M257x-iter135-adjudicated-live-defects`, each site re-derived at source before the plan was written. **(1) `ai-readiness.md:18-20`** (`adj-B` P-3, **upheld and widened**): *"the **ONLY** remaining dependency on `workforce` is the member directory … `LoadMembers`/`LoadMembersByUserIDs`, whose implementations **stayed** in `members.go`"* — measured at `app` `ad9f3c498`, the `WorkforceDirectory` interface declares **FOUR** methods (`aireadiness/manager.go:40-51`): `LoadMembers` · `LoadMembersByUserIDs` · **`BaseMembers`** · **`LevelsCount`** — **and the source's own doc comment four lines above the construct the sentence cites (`:36-39`) already says the seam is *"the active-member directory … AND the org's skill-scale setting"***; `LevelsCount` is an **org setting** (`readiness.go:770`), and it **did not stay in `members.go`** — it is at `internal/workforce/manager.go:90` (`git grep "func .*LevelsCount"` → three sites: unexported `:61`, exported `:90`, a test fake). **An absolute quantifier over a coupling seam, refuted by the doc comment on the interface it names** — the milestone's four security-surface understatements, on a different axis. **(2) `clerk-integration.md:126`** → `ant-academy.md:334` *"the `DEV_LOGIN_ENABLED` public-route pair"* — **rotted +4**; `:334` is the **AI-proxy** row and the real one is `:338`; repaired by **naming the row**. **(3) `backend.md:13`** → *"see the **M810 prod teardown is UNEVEN** bullet below"* — **that bullet was RETITLED at iter-127** to *"The M810 prod teardown has now LANDED for both"* and **its body retracts "UNEVEN" in its first sentence**. `D-M257x-141-2`: **a cross-reference that names its target by a retracted TITLE is invisible to every anchor fence, because the pointer still resolves** — the reader simply arrives at a paragraph that opens by contradicting the sentence that sent them. Sibling of `D-M257x-137-3` one level up (137 = a retracted *pin*; this = a retracted *name*), and **harder to catch precisely because nothing breaks. A title is a citation.** **And the session measured its own recurrence** (`D-M257x-141-1`): the house retraction idiom — *"it was `:274` at `<sha>`"*, *"this cited `:116-117` until iter-NN"* — **keeps the retracted number live in the text**, where the next insertion above its target moves it. **In five iters it turned fences RED three times, in three files, always on a pin whose own sentence existed to retract it**: `roadrunner.md` (iter-137), `graphql-wundergraph.md`'s `5050` pointer (**which rotted TWICE on its own**), and `ai-readiness.md`'s `:326`/`:274` note — **this iter, caused by this iter's own insertion.** → **retract by DESCRIBING the artifact, never by reproducing it**; *"this doc carried two different line numbers for it in successive iters"* says everything the quoted number said and **cannot rot**. Both land as `§5` **rule 63 (c′)** and **(c″)**, and the sweep is routed as `FIX-M257x-iter141-retraction-idiom-sweep` — **censusable by `D-M257x-140-2`, because the citing sentence carries its own head**, unlike the retracted iter-138 subject. Guards **18 GREEN / 0 RED / 4 not-run** *after* a genuine RED from this iter's own insertion (fixed by **removing** the quoted pin, not re-pinning it); scoped suites **102 passed / 0 failed** after 2 failures on the same defect — **the second consecutive iter in which the anchor suites caught a self-inflicted defect BEFORE the commit**, the mechanism `D-M257x-138-5` installed after iter-137 shipped one that survived a whole iter. **Whole suite NOT re-run and rule 60 requires saying so** — zero rext files changed in any of iters 133–141. **No reading taken; no `N` movement claimed. Gate unchanged at 4 of 5. 5-tik cap reached** — see `iter-141/progress.md`
- iter-142 (tik): **the retraction idiom censused — and the census's own blind spot, found by a guard on another axis.** `TOK-08`, closing `FIX-M257x-iter141-retraction-idiom-sweep` one iter after it opened. `§5` rule 63(c′) named the class from three incidents (*fences RED three times in five iters, always on a pin whose own sentence existed to retract it*); this iter **enumerated** it, because under `D-M257x-140-2` its subject is **the citing sentence**, which carries its own head — the thing iter-138's retracted census lacked. **The denominator, stated — and it had to be corrected mid-iter: 2,185 line pins over 94 source documents** (1,311 path-qualified, 874 bare; 23 excluded as ports, 3 in fenced blocks) → **50 live findings across 20 files, 44 bare + 6 path-qualified**; 2 remain, both in the teaching-exempt protocol doc. **Every finding was read in context before a line was repaired — 44/44 on the bare arm, 6/6 on the path arm.** ⚠️ **THE HEADLINE IS THE MISS, not the sweep** (`D-M257x-142-6`): the first pass matched only BARE pins, enumerated 922, repaired 44 and reported the class **clean at 0** — and then `repair_leak_guard`, run over this iter's own commit, went **RED** on `shared_libraries.md:77`, a twin of a `CLAUDE.md` site the new fence had just flagged, differing only in that its pin carried a path. **A fence's regex IS its denominator**, and no control inside the fence could see it: the mutation control fired correctly, the anti-vacuity arm passed correctly, both answering *is this guard working* while it worked on the wrong population. A guard on a **different axis** caught it. The three commit-scoped guards the family reports as `not-run` returned **3 RED on that commit, every one real** — a rotted cross-doc pin, the path-qualified twin, and a value-change twin. **Auditing every finding buys PRECISION and says nothing about RECALL.** The audit is the deliverable as much as the sweep — the first draft's *any-marker-anywhere* rule produced **2 false REDs of 46**, both from markers attaching to something other than the pin (`external_services.md:473`, a sha anchoring a **deletion**; `hiring.md:221`, a *"no longer"* about a deleted **FILE**), and demoting those to a **Tier B that must sit within 30 characters after the pin** took it to 44/44. Both ship as named regression tests. The repair rule is sharpened to **fence the TOKEN, not the digit** — *"rotted +8"*, *"+23 and +16"*, *"ten lines earlier"* keep every number the retraction was making a point with and are invisible to every resolver, so hygiene never costs evidence. **All 20 repaired files came out line-count FLAT: rewritten in place, added-minus-removed = 0 per file, `git diff --numstat`** — because iter-141 turned a fence RED with its *own* repair of this class, and a sweep against anchor rot must not induce any. Two sites are the class arguing with itself: `graphql-wundergraph.md:138` held **four** retracted pins in a paragraph that closes *"which is why all of them are now named rather than numbered"*, and `clerk-integration.md:126` was **iter-141's own repair**, reproducing two pins while making one. Ships `FENCE-M257x-iter142-retracted-pin` (`retracted_pin_guard.py`, registered in the family — **22 → 23 members**) with a firing 3-mutation control against a proven-green unmutated control, and an anti-vacuity arm written against the guard's SUBJECT that is unusually sharp: **`census.clauses` over the real corpus can never legitimately reach 0**, since the protocol doc must spell the idiom to teach it — so a broken matcher makes the guard's green impossible rather than merely wrong. Gates: commit-scoped guards **3 RED on the first commit, re-run GREEN** (2 real misses + 1 measured-false shingle collision, disclosed not obeyed); new module **21/0** (3 mutations + 2 anti-vacuity arms + 6 false-positive regressions across both arms); change-scoped suites (rule 63(d)) **209 passed / 0 failed**; family **19 GREEN · 0 RED · 4 not-run**, against a pre-iter baseline of 18 GREEN taken with the identical invocation, so the +1 is attributable. **Whole suite NOT re-run and said so (rule 60)** — and this is the first iter since 132 to touch `rosetta-extensions`, so iter-132's clean run **no longer stands** on this tree; a run is owed. **No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged — see iter-142/progress.md
- iter-143 (tik): **the `(bare)` orphan bucket censused; head inference built, hand-read at 100 %, and REFUSED — and the hand audit was itself wrong on 9 of 92.** `TOK-08`, closing `SURVEY-M257x-iter143-bare-orphan-bucket` (the route `4edad03` recorded as this iter's starting point) and `FIX-M257x-iter142-whole-suite-owed`. **The census, with its denominator, stated TWICE because this iter moved it: 384 at open (337 `orphan-no-referent` + 28 `ambiguous-address` + 15 `ambiguous-file-mention` + 4 `superseded-quote`) → 380 at close (333 + 26 + 17 + 4)**; the difference is this iter's own `_CODE_SUFFIX` widening plus one bare anchor its own rule-65 write-up quotes. **~88 % of the bucket is `orphan-no-referent`** — not ambiguity the guard is right to refuse, but referents it cannot see — which is the only arm any head inference could reach. **The inference was built from a regex the guard ALREADY HAS** (`_FILE_MENTION`, used only to BREAK a chain; it also describes how the corpus NAMES a file without citing it) and it places **129 of the 337 orphans, 92 resolving to a real file**. All 92 were read in full context before anything shipped, per `D-M257x-142-1`. ⚠️ **THE HEADLINE IS THE AUDIT, not the census** (`D-M257x-143-1`): the reader returned **62 true / 30 false** and the best predicate scored **90.2 %** — then those 62 "true" sites were pushed through the guard's own `classify()` and **nine came back `anchor-out-of-range` against files whose line count makes the citation impossible**. All nine are ONE mechanism a reader is structurally poor at seeing — a bare file **mention** sitting nearer the anchor than the **qualified citation** governing it (`cms.md:125`: *"`studio/gen.py` at `studioManager.go:119` and `studio/postgen.py` at `:1045`"* — `:1045` is a `studioManager.go` line; eight more on `demopatch-spec.md` belong to `up-injected.sh`, hijacked by the phrase *"the shared `urls.ts` pair"*). Corrected: **53 true / 39 false**, the predicate scores **74.5 %** — **the audit had inflated it by 15.7 points, in the direction that ships**. iter-142 taught *audit the predicate before the repair*; this iter adds the next turn — **an audit is a predicate too, and it needs a control that is not another reading** (`§5` rule 65). **The refusal, with its numbers: 57.6 % raw → 77.3 % at best (32.1 % recall). NOT SHIPPED.** That is an **answer** to `FIX-M257x-iter138-anchor-rot-fence`'s *head-resolution-first* re-specification, not a failure to attempt it. A numeric cut (`n < 3000`) separates the population perfectly and was **declined as a tuned constant on a 92-site denominator** — two are already routed open from iter-142, and a third is a pattern. **And the mechanism split is the durable part** (`D-M257x-143-3`): of the 39 false admits, **21 ports (LOUD — resolve out-of-range, show as a RED) vs 16 WRONG-HEAD (SILENT — a real line anchor booked against a file the sentence never named, which can land on a real construct and PASS)**. The guard's source comment justified refusing the bucket by naming **only the loud half**, for five iters; retracted in place as an explanation, with the numbers. **What DID ship is a different KIND of change** (`D-M257x-143-4`): `_CODE_SUFFIX` measured for the first time since iter-73 — **32** no-slash `name.EXT:NNN` citations were **invisible** to `_QUALIFIED`, not unresolved — **7 suffixes admitted** (`jsx` x8, `js` x8, `graphqls` x4, `ini`/`txt` x2, `hcl`/`sum` x1) and **3 declined**, the third being the counter-example that prices the list (**`de`**, `u422950.your-storagebox.de:23` — a HOST:PORT that would book a hostname as a file at a line number that is really a port). **Reach 861 → 875 anchors, all 14 newly-graded citations clean**; safe where the inference is not, because a qualified citation carries its own path and decides nothing the prose had not already said. Plus the census now REPORTED in the text run and `--json`, with `BARE_REFUSAL_REASONS` as one constant so the three spelling sites cannot drift. Gates: guard module **24 passed (+11 net-new)**; `run()`-tuple consumers **106 passed**; guard **GREEN before and after every edit**; **whole `stack-core` suite RE-RUN — the owed run from iter-142, so coverage is quoted rather than assumed**. **No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged — see iter-143/progress.md
- iter-144 (tik): **the wrapped retraction sites graded, repaired and explained — and the route's own count corrected.** `TOK-08`, closing `FIX-M257x-h30-crossline-repair`. Harden pass 30 routed *"the **8** true sites across 6 files"*; re-surveyed at open, `retracted_pin_guard`'s wrapped (SURVEY) arm holds **10 findings**. All 10 read in full context — reading says **7 true / 3 false** — and then, per `D-M257x-143-1` landed one iter earlier, **every pin the reading called LIVE was resolved against the source it claims, by machine**: `ai_architecture.md:303`'s `:98-99`/`:110-111` resolve at `app` `ad9f3c49` to the two `default: aiModel = …` lines the sentence describes, and `hiring.md:304`'s `:176-187` **is** the `job_position` bullet at HEAD. **Three for three — the control CONFIRMED this reading where at iter-143 it overturned nine of ninety-two** (`D-M257x-144-3`: a control earns its place by being run, not by firing). ⚠️ **THE FINDING IS A SUB-CLASS RULE 63(c′) DOES NOT COVER** (`D-M257x-144-1`): **a retraction clause holds the CORRECTED pin as often as the retracted one.** `ai_architecture.md:303` retracts the *absence of a filename*, not the numbers; `hiring.md:304`'s retracted value is *"an earlier range"*, left deliberately unnamed, while the pin actually present is the **fix**. Same shape, same sentence, same markers — **the token a retraction must not reproduce is the OLD value; the token a correction must publish is the NEW one**, and no form-matching fence can separate them. **Wrapped-arm precision 7/10 = 70 %**, and the remedy is **not** to tighten it — pass 30 was right to keep the arm non-gating; the remedy is to grade before repairing. **Repair: 7 sites across 4 files, every file line-count FLAT** (`git diff --numstat` 1/1, 2/2, 3/3, 4/4) — `shared_libraries.md:38`, `ai-readiness.md:113`/`:487`, `graphql-wundergraph.md:90` (two pins in one clause), `roadrunner.md:85`/`:228` (both the **ref-qualified** `D-M257x-142-2` class — *"at `2adcf71`"* makes them true statements about an immutable ref, and the fence still cannot read the qualification). Every sentence keeps the number it was making a point with (the **9** entries, the **seven** modules, the **twice**) and loses only the rot-prone token — *fence the token, not the digit*. `§5` gains **rule 67**. **`D-M257x-144-2`: grade a survey arm's findings before treating its count as a backlog — a routed count is an estimate of WORK, and quoting it makes it a measurement of DEFECTS**, a conversion this milestone has now watched twice. Gates: wrapped arm **10 → 3**, all 3 the machine-confirmed live pins; gating arm GREEN; `test_retracted_pin_guard` **51 passed**; guard family **19 GREEN · 0 RED · 4 not-run**, identical to iter-143's close so the delta is attributable. **Whole suite NOT re-run and saying so (rule 60)** — this iter touched **zero** `rosetta-extensions` files, which is precisely the condition under which iter-143's run still covers the tree. **No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged — see iter-144/progress.md
- iter-145 (tik): **the four never-run sections were RUN, and the 21 routed failures graded — 12 of them are OURS.** `TOK-08`, closing `FIX-M257x-h30-nonstackcore-suite`. Every "whole suite" figure in 144 iters + 32 harden passes meant `stack-core` alone (**1,280 of ~3,040 tests, 42 %**); run in full, `demo-stack` **9 failed**·1,047 passed, `dev-stack` 151 passed, `stack-injection` 335 passed, `stack-verify` **12 failed**·225 passed — **21, the routed count reproduced exactly**. Graded per-failure (`D-M257x-144-2`, one iter old): **12 real defects · 6 sha baseline drift · 3 host environment**, so passes 30/31/32's *"provably not ours … environment-coupled"* is **FALSE for 57 %** of what it described — and six of the twelve touch no clone and no container at all. ⚠️ **THE 12 ARE ONE CAUSE AND IT IS THIS MILESTONE'S OWN** (`D-M257x-145-1`): platform `2adcf71` deleted the GraphQL router, **M257x iter-13** (`4414527`) dropped its row from `stack-verify/lib/services.sh` and left the **test side's copy** — a `BASES` map plus **six independent count literals** (13·13·12·10·13·14). Twelve tests went RED on 2026-08-01 and stayed RED for **132 iters**, because no iter close and no harden pass has ever executed that section. **iter-13's own commit message names the defect it was fixing — *"six copies of a platform fact is the hand-maintained-tuple defect M257x exists to end"* — and left a seventh copy in the one place nothing watched.** `"not ours"` came from `git diff --name-only <5-iters>..HEAD`: a **window that opens after the breaking change can never see it** — bisect the failure, don't scope the diff. Repair (**+85/−31, one file**): ONE literal `REGISTRY_BASES` (hand-written on purpose — it is the offset sweep's anti-vacuity control, `§8`), **every count derived from it** (six literals → zero), two `graphql`-specific assertions re-pointed at rows that still exist (`roadrunner` 10400→40400, `storage`), the `"14 of 16 containers Up"` M256 figure **kept attached to its ref** instead of restated, and a **named-drift fence** — `test_the_test_side_registry_mirrors_services_sh`, whose message names the drifted rows in **both directions**. `D-M257x-145-2`: **the count is the copy that rots; the set is the control** — and the fence's value is not detection (six arithmetic assertions already detected it) but that it **says what drifted**, where `12 != 13` said nothing for four months. Anti-vacuity control RUN, not assumed: a re-drifted copy fires the fence by name and the sweep by arithmetic. **`stack-verify` 12 failed·225 passed → 0 failed·238 passed.** `D-M257x-145-3` states the standing **scope call** as an assumption pending the user's ruling — *"the suite" means all five sections* — with the cost argument measured and refuted: the four excluded sections cost ~11½ min **together**, less than half the one included section. `§5` gains **rule 68**; `§9` gains a net-new **measurement-preconditions** block (`/usr/bin/python3` 3.9.6 is the only interpreter here with pytest; no `timeout(1)` on macOS) — both facts were already in `hardening-ledger.md`, which the iter loop does not read. **No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged — see [`iter-145/progress.md`](iter-145/progress.md)
- iter-146 (tik): **iter-13's re-point censused end-to-end — 97.6 % complete, and both misses are EMITTERS on paths nothing exercises.** `TOK-08`, the question iter-145 routed rather than answered. Denominator stated: **84 `5050` references across 31 files** in `rosetta-extensions` — **62** guard prose / test fixtures (about the *corpus's* ports), **8** fences asserting the port's absence, **12** comments explaining the deletion, **1 LATENT**, **1 LIVE**. ⚠️ **THE LIVE ONE:** `dev-stack/dev-stack:285` printed, on every successful `--public-host` dev bring-up, `https://<host>:$((5050+n*OFFSET)) (graphql)` — a port with **no container** (`2adcf71` deleted the compose service), **no probe** (no `services.sh` row) and **no front** (iter-13 deleted its `tailscale serve` row). **And `gen_tailscale_serve.py:38-41` states that exact hazard in its own words** — *"fronting a port with no listener produced a trusted-cert HTTPS endpoint that always refused, **which is worse than no entry at all (it looks configured)**"* — while the file one directory over advertised it. **`D-M257x-146-1`: the repair removed the MECHANISM and left the ANNOUNCEMENT** — grep the *emitters* of a retired fact, not only its consumers, because a consumer fails loudly and an emitter is just a string that nothing in a suite, a health check or a bring-up can notice. The latent one: `hiring.Dockerfile:36` defaulted the client-bundle endpoint to `http://localhost:5050/graphql` — verified unreachable on the sanctioned path (`up-injected.sh:874`/`:1308` always pass `--build-arg`), repaired anyway because Next bakes `NEXT_PUBLIC_*` into the bundle and a hand-run `docker build` would ship a dead endpoint silently. Both re-points name **the path as well as the host** (`:5050/graphql` → `:8082/graphql/query`) — a host-only correction 404s rather than refusing. **`D-M257x-146-2`: a repair's completeness tracks EXERCISE, not care** — both misses sit on un-exercised paths (a never-run test section at iter-145; a `--public-host` branch needing tailscale *and* a public host here), and everywhere the code runs iter-13's work held. Fence: `stack-core/tests/test_deleted_router_endpoints.py` (3 tests) over **seven emitter files**, matching **executable content only** — **`D-M257x-146-3`: the comment carve-out is `§5` rule 67 in a second domain**, the same token carrying opposite obligations, except that in *code* the language marks the axis for free where in prose it does not (which is why that arm stays a 70 % SURVEY). Allowlist scope on purpose — a repo-wide sweep returns 84 hits to grade 2 — with `test_every_emitter_exists` so it cannot pass vacuously. **Control run against the REAL pre-fix text recovered from `HEAD`, not a reconstruction: 2 hits before, 0 after**, plus a synthetic pair proving the carve-out exempts a documenting comment and does *not* exempt an emitted URL. `§5` rule 68 gains sub-rule **(d)** with the census table. **No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged — see [`iter-146/progress.md`](iter-146/progress.md)
- iter-147 (tik): **every path that CHOOSES a compose profile censused — 7 sites, 5 correct, 2 live defects — and the finding is that a token census cannot see an ABSENT value.** `TOK-08`. The routed target (`SURVEY-M257x-iter146-other-retired-services-unaudited`) was **re-surveyed and answered with a measurement rather than executed**: the four ported retired services (`cms` 8090, `jobsimulation` 8400, `storage` 8300, `roadrunner` 10400) occupy **74 lines / 0 emitters** — registry rows, tests about those rows, port arithmetic. So the search was **inverted** (`D-M257x-147-1`): enumerate what the tooling *chooses*, not what it *says*, because iter-147's defect is an **empty** value and an empty value has no token to grep for. **Denominator: 7 profile-selecting compose sites across 3 entry points** — `up-injected.sh` (derived, die), `rosetta-demo` `cmd_down` (derived, non-fatal by the F-9 rule), both `dev-stack` verbs (derived, die — iter-85), `dev-stack` teardown (label sweep) → correct; **`rosetta-demo`'s `cmd_up` and `gen` → DEFECT.** ⚠️ **An empty compose profile is not "the base profile" — it is NO profile:** `docker compose up -d` then selects only services declaring no `profiles:` key (postgresql, redis, sentinel) and **exits 0**, `gen_override.py --profiles ""` resolves the same reduced set (its own help says *"empty = base/always-on only"*), and the run **announced success three times** — `profile='base'` to the operator, `"profile":"base"` into the unified registry that `/stack-list` reads, and `==> demo-N up.` **Compose has no `base` profile**; the word was this script's own name for *no profile at all* (`D-M257x-147-4`, both substitutions deleted). **And the twin lag is the expensive half** (`§5` rule 69): this defect was repaired three times before — iter-55 (demo teardown + injected-gen), iter-85 (both `dev-stack` verbs) — and **iter-85's own comment names the lag it was closing** (*"the dev path kept the literal for four more releases"*) while leaving the demo path's other two verbs for another **62 iters**. An observation about a twin is not a fence over it. Repair: `derive_profile()` calling the same `platform_topology.py profile` its sibling function already called, **no literal token in the file**, wired into both verbs, **refusing** on a derivation failure with the consequence named — the asymmetry against `cmd_down`'s deliberate fallback recorded as `D-M257x-147-2`, not left to be "cleaned up". Fence: `demo-stack/tests/test_profile_derivation_m257x.py` (**8 tests**), behavioural — the shipped `derive_profile`+`gen` **extracted from the real script with `awk`** and asserted on the **argv `gen_override.py` actually receives** — with a RED-proof against a guard-deleted mutant, an anti-vacuity control proving both structural predicates can fire *and* do not over-fire, and a **twin-drift arm naming which side drifted**. **Control run against the REAL pre-fix text recovered from `HEAD`: `--profiles` was `''` → now `'core'`; placeholder announcements 2 → 0.** `D-M257x-147-3`: the first suite run was **13 failed**, graded before quoted — **9 pre-existing (identical by name to iter-145's set) and 4 mine**, four `RosettaDemoRegistry` tests whose stub platform `docker-compose.yml` was created **empty** and which then asserted `up` returned 0; **the fixture was the defect encoded as expected behaviour** and was repaired, not bypassed. Docs re-pointed the same commit (README + GUIDE stopped teaching the by-hand derivation) and **`demo-up-defaults.md`'s `--profile` row gained the default it never had** — *an omitted default is a claim, not a gap*, inside the one document written to prevent exactly that. Gates: fence **8 passed**; `demo-stack` **9 failed · 1,055 passed** (the 9 = iter-145's baseline by name, +8 = this fence, **0 regressions**); `dev-stack` **151 passed**; guard family **19 GREEN · 0 RED · 4 not-run**, identical to iter-146's close so the delta is attributable. **No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged — see [`iter-147/progress.md`](iter-147/progress.md)
- iter-148 (tik): **the verify entry points censused — 3 sites, 1 unscoped, and it is the `/test-platform` driver.** `TOK-08`, taking iter-147's own `SURVEY-M257x-iter147-absent-value-class`. `stack-verify/lib/services.sh:33` **discloses this defect in its own comment** (*"running verify.sh with NO STACK_SERVICES set probes everything in the table and will false-`down` the merged-away rows"*) and had for as long as the rows existed. **Measured on the live `demo-1`, read-only, both arms: UNSCOPED ✗ 6 of 20 probes failed vs SCOPED ✗ 1 of 14** — the five-probe delta entirely `jobsimulation`/`cms`/`storage`/`roadrunner` liveness at `HTTP 000000` (no listener, and there never will be one) plus `storage-rpc` readiness on the same cause. **So `/test-platform` — the tool whose entire purpose is *"what is actually working"* — reported four services the platform deleted as DOWN and exited 1**, its documented invocation (`SKILL.md:76-78`) setting `STACK_ROOT` and `REPORT_DIR` and nothing else, while every other caller (`up-injected.sh:2681`, `dev-stack:347`) already derived its scope. **The remaining 1 is NOT booked as a defect** (`D-M257x-148-4`): `postgres-schemas` fails in *both* arms because its `repos.yml` candidate path resolves in a per-stack consumption copy and not in the `.agentspace` authoring copy this iter measured from — substrate, per `D-M257x-122-4`. Repair: `generate.sh` derives `STACK_SERVICES` from `$STACK_ROOT/platform` when unset and **DISCLOSES rather than refusing** when it cannot (`D-M257x-148-1` — *the disposition follows what the artifact is FOR*: a bring-up that cannot name its scope must not announce a stack; a **report** that cannot name one must still be produced, with the caveat printed **into the markdown** where stderr is what gets lost). Caller-supplied scope always wins; both branches exercised end-to-end. `SKILL.md` gained the scope note **and** the `STACK_PROJECT`/`STACK_OFFSET` requirement — unset, the probes silently target the main dev stack. Fence: `stack-verify/tests/test_probe_scope_m257x.py` (**6 tests**), load-bearing arm **derived-vs-declared in both directions** — the registry rows the platform no longer declares, minus the rows this tooling itself injects, must **equal** the set the disclosure names — so the warning goes RED at the next fold instead of reassuring a reader about the wrong four services (`D-M257x-148-2`; a hand-written set is what rotted as six count literals at iter-145). ⚠️ **THE RED-PROOF CAUGHT A REAL DEFECT IN THIS ITER'S OWN FENCE** (`D-M257x-148-3`): the "derivation precedes invocation" check used bare `src.find()` and **both tokens bound to COMMENTS** — `generate.sh`'s usage header names `verify.sh`, and the comment block explaining this very fix names `platform_topology.py` — so it would have kept passing after the guarded code was deleted. **The mutation control found it, not review.** Now bound to executable content only. `§5` rule 67/68(d)'s axis, reproduced inside the fence written to apply it, **twice in one iter**. Gates: fence **6 passed** with the RED-proof proven load-bearing; `stack-verify` **244 passed · 0 failed** (238 at iter-145's close + 6 new); guard family **19 GREEN · 0 RED · 4 not-run**, identical to iter-147's close. `stack-core`/`demo-stack`/`dev-stack`/`stack-injection` **not re-run and saying so** (rule 60) — zero files touched in them, and two of the four ran in full one iter ago. **No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged — see [`iter-148/progress.md`](iter-148/progress.md)
- iter-149 (tik): **the emitter census widened from one retired service to all twelve — 354 references, 17 in executable position, 0 emitters** — closing iter-146's route with a falsification rather than a repair. The defect the census found was **in the census's own denominator**: the raw run returned 50 executable-content hits, of which **33 (66 %) were `.m220-mutant-*` files** — copies the M220 mutation battery stages *beside* its subjects (deliberately: those scripts resolve siblings from `$HERE` and cannot run from `/tmp`) and deletes in `tearDown`, left behind by interrupted runs and **accumulating since 2026-08-04**. Every one carried, verbatim, the tailnet-URL line iter-146 repaired — **a stale mutant is a perfect forgery of the bug its own battery exists to catch**, and `.gitignore:32` had described the leak in as many words the whole time (iter-148's lesson 0 again: a file that discloses a hazard in a comment is a file nobody fixed). Repaired by a `setUpClass` sweep that is **age-gated at 1 h, never unconditional** — a mutant belongs to a live run while that run holds it, and deleting a concurrent sibling's staging surfaces as a *mutation result*, not as a missing file; proven on the real leftovers before the fixture was spent, 33 → 0 repo-wide. The fence then generalised (`test_deleted_router_endpoints.py` → `test_retired_service_endpoints.py`): the retired set **imported** from `claim_census_guard.ARCHIVED_SERVICE_NAMES` rather than re-declared (the first seam in the fence family to consume another guard's subject set — narrows `-iter134`), base ports **derived** from the tooling's own probe registry at the ref before the merges stripped its rows (`services.sh` @ `c95bce4`), iter-146's comment carve-out and pair control kept, and the `services.sh` rows carved out **with their owning route named** (`SURVEY-M257x-iter148-registry-is-hand-maintained`) rather than silently. **RED-proofed on a real answer key** — 1 hit against pre-iter-146 `dev-stack` (`1a44b97^`), 0 against the current file — and **every arm proven to fire** on synthetic content, because a census returning zero cannot tell a working arm from a decorative one. Side discovery, separate commit: `REXT_SECTION_NAMES` called itself *"derived from the monorepo's own layout"*, was declared, and had drifted — **`stack-secrets` missing, 10 declared against 11 on disk**, so claims naming the section behind `/stack-secrets` resolved to no known artifact and left the census silently (iter-129's class, inside a guard whose job is noticing omissions); fixed, comment made honest, and the asserted property fenced against the layout in both directions. Also a near-miss worth keeping: the pytest run failed under homebrew `python3` and the first draft booked a route for it — **`§9`'s first measurement-precondition bullet already said the interpreter is `/usr/bin/python3`**, and reading it deleted a fabricated finding. `§9` gains the exhaust-vs-subject precondition and a new subsection, *A census that returns ZERO must prove its instrument*. — see iter-149/progress.md
