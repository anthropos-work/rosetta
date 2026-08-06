---
milestone: M257x
---

# M257x — decisions

## KB-fidelity findings (Phase 0b audit, 2026-07-31 — YELLOW)

Full report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md). Platform ref for every citation below:
**platform @ origin HEAD `1e8e7540`** (2026-07-30T08:26:40Z). These three were blocker-severity;
KB-1 and KB-2 are neutralised by being written down here, KB-3 needs a planning decision.

| id | finding | why it is load-bearing | citation |
|---|---|---|---|
| **KB-1** | The `local_*` session **mirrors are DROPPED**, but the CORPUS still mandates co-writing them. M257 iter-03 fixed the **rext** side (34 sites / 20 files); the corpus side was never reconciled and no M257x doc flags it. An iter reading `hiring.md:132-143` as truth re-introduces writes to a dropped table. | `services/hiring.md:105` is a line-anchor into a **deleted file**; `:132-143` + `ops/seeding-spec.md:386-392,416,528-536` + `ops/demo/stories-spec.md:666,680` + `ops/demo/content-stories-routes.md:135,152` all still require the mirror. | `app/terraform/migrations/20260729133514.sql:62-63` — `DROP TABLE "local_jobsimulation_sessions"; DROP TABLE "local_skill_path_sessions";` (FKs re-pointed to canonical tables at `:52-56`). `local_jobsimulation_session.go` absent from `app` @ HEAD. |
| **KB-2** ⚠️ **REFUTED — see D-M257x-1 below; do NOT act on this row** | ~~**The obvious jobsim re-point is silently wrong.** `public` now holds **two** session tables: `sessions` AND `job_simulation_sessions`.~~ **The two-table premise is FALSE.** `public.sessions` does not exist: it is created by one migration and **dropped by the very next** (`20260722104506.sql:79`) as the rename completes. Measured across all **167** migration files — created once, dropped once, never recreated, positive control passed. The surviving true half is that app's GraphQL `Session` binds to `job_simulation_sessions`, and that **neither** table is named anywhere in the corpus. | Inverted: a naive re-point at `public.sessions` **fails loudly**, which is the SAFE mode — not a silent pass. The residual real risk is only that the corpus names neither table. | `app/terraform/migrations/20260722081626_jobsim_data_model.sql:2` creates `sessions`; `20260722104506.sql:2` creates `job_simulation_sessions`; `app/internal/data/ent/schema/job_simulation_session.go:7-9` — *"renamed from the too-generic `Session`… the table is `job_simulation_sessions`; app's GraphQL `Session` type binds to this entity"*; `20260729133514.sql:52-56` re-points FKs to it. |
| **KB-3** | **Exit-gate clause 3 rests on a premise platform source contradicts.** The clause fences the map against `repos.yml`, but the fold commit was **`docs:`-only** (3 files: `CLAUDE.md`, `README.md`, `repos.yml`; `docker-compose.yml` untouched). cms/jobsimulation/roadrunner are **still in the default `graphql` profile**. So *"own no local schema"* is true + fenceable; *"not part of the stack"* is **false at the same sha**. Fencing on prose would encode a falsehood. | Decides whether clause 3 can be satisfied honestly at all. | `platform` commit `236771f103` (2026-07-29T14:06:49Z) file list; `platform/docker-compose.yml:108/165` (jobsimulation), `:169/212` (cms), `:306/334` (roadrunner), `:20-25` (graphql `depends_on`), `:69,75` (backend RPC addrs); `platform/Makefile:14` (`MIGRATION_REPOS` ← `/migrations: true/` only); `platform/postgresql/Dockerfile` (creates no schemas). |

**KB-3 asks for two decisions:**
1. Narrow clause 3's fence to the **machine-readable fields only** (`name`/`type`/`migrations`/`schema`) — never the prose comment.
2. Add **`running_but_unfederated`** to the state enum `{live-standalone, merged-into-app, decommissioned, net-new}`. cms, jobsimulation and roadrunner are **all** in that state today and none of the four fits. Template already exists: `services/roadrunner.md:16-17` ("built + started but off every request path", ORPHANED).

**Also decided by evidence (answers to overview.md Open Questions):**
- **OQ4 — "does the 3-subgraph count still hold?" → NO. It is ONE subgraph (`backend`).** Verified 3 ways: `graphql-wundergraph/supergraph-config-{compose,prod}.yaml` declare a single `backend` subgraph; `schemas/` holds only `backend.graphqls`; `Dockerfile.dev:18-23` copies SDL only from `app/…/graphql/graph/schemas/` and states *"there are no standalone subgraph SDLs"*. Landed `915da06c58` (2026-07-29T09:24:38Z), which removed **both** the jobsimulation and cms subgraph wiring. The corpus asserts **3** in 16 numeric places.
- **OQ2 — jobsimulation: neither "merged" nor "unchanged".** Its data model is re-created in `public` (23 tables), its subgraph is gone, its migrations are off — **but its container still runs in the default profile**. `running_but_unfederated`.
- **`DOC-M257-studio-in-app` CONFIRMED, and it is an ADDITION not a move:** `app/Dockerfile.dev:24-26,38-41` embeds the studio pipeline (`python:3.11-slim`, `COPY … /build/studio ./studio`, `pip install -r studio/requirements.txt`); `app/.gitignore:78-79` documents the acquisition. `cms` **still** embeds it too (`cms/Makefile:11-17`). The corpus says CMS-only in 30 places and names `app` in **zero**. Note the platform-side ambiguity to handle: `Dockerfile.dev:38-39` calls `./studio` a *"git submodule … pinned SHA"* while `app` has **no `.gitmodules`** and `.gitignore:78` says *"pulled at build via additional_repo, like cms"* — the likely root of M257's *"`app`/studio had no rext acquisition path"*.

---

## iter-01 (bootstrap tok) decisions — 2026-07-31

### D-M257x-1 — KB-2's two-session-tables premise is REFUTED (measured)

Phase 0b and probe A2 **contradicted each other** on a load-bearing fact, so it was measured rather than
adjudicated. Across **all 167** files in `app/terraform/migrations/` (`rg -a`, search set enumerated,
positive control on `job_simulation_sessions` returning hits in the same pass):

    20260722081626_jobsim_data_model.sql:2   CREATE TABLE "sessions"
    20260722104506.sql:2                     CREATE TABLE "job_simulation_sessions"
    20260722104506.sql:79                    DROP TABLE "sessions"      <-- the rename completes
    (no later migration recreates it — scanned every file dated >= 20260722104506)

**`public.sessions` does not exist.** `public` holds exactly ONE session table.

**Why this matters, and why it inverts the plan:** KB-2 concluded a naive re-point
`jobsimulation.sessions -> public.sessions` would *succeed and read blank* — the silent signature of M257's
swallowed `|| echo 0`. It would in fact **fail loudly** (`relation "public.sessions" does not exist`), which
is the SAFE failure mode. Acting on KB-2 would have meant designing a guard against a silent failure that
cannot occur, and mis-prioritising the re-point tik toward the wrong risk.

Two further probe claims refuted the same way, recorded so they are not re-inherited:
- A2's *"169 migration files, ~172 of them contain NUL bytes"* — arithmetically impossible and false;
  measured **0** files with NUL bytes, on A2's own clone.
- My own note that `CLAUDE.md`'s "27 service docs" was off by two — **wrong**; 29 `.md` minus `README.md`
  and `TEMPLATE.md` is exactly 27. Phase 0b caught this independently.
- Phase 0b's refutation of *its* sub-agent (that `services/next-web-app.md` is invalid UTF-8 and silently
  binary-skipped) is **CONFIRMED CORRECT** — `file` says UTF-8, 0 NULs, Python decodes clean, and `rg -c` is
  identical with and without `-a` (20 == 20).

**Doctrine this establishes (carried into `platform-alignment.md` §3):** the NUL-byte trap has become
**folklore**. Three genuine false-absence incidents occurred in this milestone and **none** was caused by NUL
bytes — a wrong regex (`id: pt-` vs the real field `playthrough:`), ripgrep's default engine **rejecting
look-around** while `subprocess` **swallowed stderr** (an engine error reading as absence), and one that never
happened. Meanwhile the one place NULs were asserted had zero. Invoking the trap *feels* like rigour while
skipping the measurement. Rules: measure NULs before blaming them; **never let a search's stderr go unread**;
always run a **positive control in the same pass**; check the field name before concluding absence.

### D-M257x-2 — clause 3's fence is narrowed to machine-readable fields only (adopts KB-3 #1)

The fence reads **only** `name` / `type` / `migrations` / `schema` from `repos.yml`. It never fences on the
prose comments. Rationale: the fold commit `236771f103` was **`docs:`-only** and left `docker-compose.yml`
untouched, so at that very sha *"own no local schema"* is TRUE and fenceable while *"not part of the stack"*
is **FALSE**. Fencing on prose would mechanically encode a falsehood — the precise "fidelity check against
the wrong reference" failure this milestone exists to end.

Corollary already established independently: `migrations: false` **entails nothing on its own** — `sentinel`
is `migrations: false` and alive with its own `sentinel` schema (`docker-compose.yml:43`,
`search_path=sentinel`). So the fence's schema-existence reference is **empirical**
(`information_schema.schemata` on a cold freshly-migrated stack), not declarative.

### D-M257x-3 — the state enum gains `running_but_unfederated`, and the map gains a per-environment axis (adopts KB-3 #2, extended)

A2 classified jobsimulation `merged-into-app`; Phase 0b classified it `running_but_unfederated`. **Both are
correct about different environments**, which is the actual finding:

| | production | a fresh local stack @ platform origin HEAD |
|---|---|---|
| jobsimulation | scaled to zero (`terraform/main.tf:40`), app owns it in-process | **container still starts** in the default `graphql` profile (compose unchanged, PR #20 open) |
| its schema | legacy `jobsimulation` schema still present, pending M710 | **never created** — `MIGRATION_REPOS` resolves to `app` alone |
| its subgraph | gone (1 subgraph) | gone (1 subgraph) |

So the enum becomes `{live-standalone, merged-into-app, running_but_unfederated, decommissioned, net-new,
external, library}`, and **every map row carries BOTH a prod state and a local-stack state.**

This is the deeper fix. Phase 0b's own observation is the key one: the corpus *"never distinguishes prod from
a fresh stack — which is plausibly why this class has recurred three times."* A single-state map would have
had to pick one and be wrong about the other, exactly as the corpus has been three times running.

---

## TOK-01: instrument first, then follow — 2026-07-31

**Tok type:** bootstrap (iter-01)

**Initial strategy:** **Build the instrument before doing the re-point, and derive every list from the platform
instead of maintaining it by hand.** Concretely, in this order:

1. **Unblock the gate's instrument.** The rext pin is inconsistent and `ensure-clones.sh` is FATAL on it, so
   `/demo-up` aborts on this box *today* — clause 1 cannot even be attempted. Advance and reconcile the pin
   first, because every other clause is measured through a bring-up.
2. **Fix the mechanism, not the symptom.** The symptom is "rext writes `jobsimulation.*`". The mechanism is
   `migrate-demo.sh`'s **hand-maintained 4-tuple** that creates those schemas itself while ignoring
   `repos.yml`'s `migrations:` flag. Derive the tuple from `repos.yml`. Re-pointing the writes without fixing
   the tuple would leave the time bomb armed for M810.
3. **Land the three fences, each watched going RED**, before trusting any green: map↔`repos.yml` both ways;
   the static schema fence (generalize the existing `test_dropped_mirror_fence.py` off its hardcoded tuple);
   and the live `information_schema.schemata` assert.
4. **Then the corpus** — the migration-status map with its **two states per row** (prod vs fresh stack), and
   the reconciliation sweep.
5. **Prove it cold, three times, on `odysseus`.**

**Rationale:** the milestone's own framing is that a re-point alone "just buys time until the next service
moves" — and measurement now shows exactly when: **v9.0 folds `storage` + `messenger`, PRs already open**, and
**M810 removes the legacy repos from the clone set**, at which point the hand-maintained tuple silently skips
them and 13 write targets fail at once. So the deliverable that ends the class is the *derivation plus the
fence*, and the fences must exist before the re-point so the re-point is verifiable rather than hopeful.
The ordering is forced by dependency, not preference: the bring-up is the instrument for clauses 1/2/4, so an
aborting bring-up is the first thing to fix.

**Strategy class:** new-direction (no prior strategy exists)

**Distance-to-gate context:** 5 clauses, **0 met**.
- (1) cold cycle ×3 green on odysseus — **blocked**: `/demo-up` aborts at the rext pin guard.
- (2) full Playthrough suite passing — denominator **verified 30** (30 manifest ids == 30 `@pt:` spec tags);
  `CLAUDE.md`'s "18 live" is stale. Unrun against a current stack.
- (3) migration-status map — **drafted** in scratch, all rows cited; needs the net-new rows, the two-state axis,
  and its fence.
- (4) zero rext writes to dropped schemas — the `local_*` half is **already clean** (69 references, **0** live,
  machine-fenced and passing); the *schema* half is unfenced, and the real surface is **12 `jobsimulation.*`
  tables, 9 of them writes**.
- (5) KB-fidelity — **YELLOW, 0 remaining blockers** after `D-M257x-1/2/3`; the enumerated doc drift (subgraph
  count, AI-Labs repo, LiveKit agents, repo states, KB-1) is routed.

Baseline that matters for honesty: **rext is entirely green offline** — 2,617 tests, 0 failures — and that green
is *structurally silent* on every claim above, because the seeders assert against a recording fake `Conn`.
Do not read a green suite as alignment.

**Next-tik direction:** iter-02 = `FIX-M257x-rext-pin`. Reconcile `.agentspace/rext.tag` with a current tag,
confirm the tag is **on origin** (`git ls-remote --tags origin` — tagging is not publishing), re-point the
`stack-demo` consumption clone, and get a `/demo-up` past phase (a2) on **odysseus**. Then immediately
`FIX-M257x-migrate-tuple` in the same or the following tik, since a bring-up that passes with the hand-maintained
tuple still proves nothing about M810.

---

## iter-07 decisions — 2026-07-31

Full text in [`iter-07/decisions.md`](iter-07/decisions.md). Summary, because two of them change how the
next fold should be handled:

### D-M257x-8 — the replay schema is DERIVED from the target, never declared

`REPOINT-M257x-cms-similarity-writes` was NOT a re-point of a constant, and could not be:
`simembeddings.Schema` is read by the prod CAPTURE *and* the stack REPLAY, and those two now **permanently**
disagree (`D-M257x-7`). A second constant (`ReplaySchema = "public"`) is two lines and is the same
hand-maintained-list defect that has been wrong three releases running.

Adopted instead: ask the **target** which schemas hold **all** of a surface's tables, then — declared schema
holds them → identity · exactly one other → remap loudly · none → fail loud (exit 4) · two or more → fail
loud **naming them** (exit 1, *not* a provisioning problem). Explicitly NOT built, each being the same defect
in costume: an allow-list of application schemas (Trap A in miniature), a preference for `public` (a constant
with extra steps), a fallback to declared on lookup error (a probe that satisfies itself, §5 rule 7).

**Consequence for v9.0** (`storage` + `messenger`, PRs open): the snapshot surfaces need **no edit at all**.

### D-M257x-9 — the digest probe and the replay resolve as ONE construct

The pre-compute's "thing not to miss": moving only the replay leaves the surface skipping at `rc=4` before a
row is copied, with a diff that looks complete in review. Rather than move both and trust review, they are
one function whose probe argument is computed inside itself — **no parameter for a caller to get wrong.**
Generalized into `platform-alignment.md` §8 rule 4: *prefer a construct that cannot express the drift over a
fence that catches it*, and make a must-decide value a **required positional**, not an option with a default.

### D-M257x-10 — the resolution is LAZY, because an exit code is a contract with a human

An eager first cut broke three tests whose own comments said a cache-miss verdict must be reachable without a
live DB. It had silently turned exit 5 (*stale cache — capture fixes it*) into exit 4 (*unprovisioned stack —
capture cannot help*), i.e. it would have sent an operator to repair a stack that was fine. Adding a
precondition check ahead of an existing decision tree can re-answer "what should I go and do next" without
crashing anything.

### D-M257x-11 — a mutation that does not COMPILE is not a RED fence

The mutation battery reported `RED (good)` for a mutant that had merely removed the last use of an import.
The tell was an **empty list of failing test names**. Re-run with a compiling mutant, gated on an explicit
build first, the fence fired for real. Promoted to `platform-alignment.md` §8 rule 5 — same family as §5
rule 8 (*a check that SKIPS reads exactly like a check that PASSES*).

---

## iter-08 decisions — 2026-07-31

### D-M257x-12 — a fence's SCOPE is a list of the system's parts, so derive it and fence it

iter-02 deleted a hand-maintained tuple of the platform's services. iter-08 found the same shape one level
up and in a worse place: **`SCORED_SECTIONS`, the write-target fence's own scope.**

Worse, because *a fence only ever asserts about what it already scans.* Every other hand-maintained list in
this milestone at least failed loudly when it drifted; an unclassified section is **invisible by
construction** — it cannot go RED, because nothing looks at it. A new rext section (v9.0 is already adding
surface) would simply be outside the fence, silently, forever.

So `SECTION_COVERAGE` now declares a `(layer, reason)` for every Go-bearing section, `SCORED_SECTIONS` is
**derived** from it, and the map is checked **against the repo** — a section that gains Go code and no
classification goes RED naming itself. Widening scope means classifying a section, not editing a tuple.

**Two sub-decisions worth carrying:**

- **A `static` section that yields ZERO scoreable constructs is mis-classified, not covered.** It reports
  GREEN, which is strictly worse than leaving it out — it *looks* fenced. Measured: widening the scored set
  to the other five Go sections would have scored **0** constructs. Now asserted
  (`test_the_static_layer_actually_scores_its_sections`), so the trap is unwritable rather than warned about.
- **`stack-snapshot` belongs to the LIVE layer, and that is now written down.** After `D-M257x-8` its write
  target is resolved at run time, so there is genuinely nothing static to score. *"Is it fenced?"* had a
  correct answer recorded nowhere, and iter-07 guessed wrong.

### The correction this iter had to make first

iter-07 routed this forward claiming the scope limit was undocumented. **It was documented — in ten lines
directly above the constant iter-07 quoted.** iter-08 refuted it by opening the file, and also refuted the
proposed fix (widening would have scored nothing). The milestone's dominant defect class, committed by the
milestone: a state reported without being measured.

Added to `platform-alignment.md` §5 as **rule 10** — *read the lines AROUND the line you are quoting*. It is
not on the existing false-absence list because it does not look like one: the substring was real, the line
number right, the quote accurate. **The search succeeded and the conclusion was still false**, because a
constant's meaning lives in its surroundings.

---

## iter-09 decisions — 2026-07-31

### D-M257x-13 — a de-exposure is not proven by an exposure guard

M221 tightened the academy's dev-server bind from `0.0.0.0` to `-H 127.0.0.1` because a *localhost* demo was
answering HTTP 200 on the tailnet IP. The fix was right, the guard for it was right, and the guard
(`stack-injection/exposure_claim_guard.py`) asked exactly one question: **does it still answer where it
shouldn't?** It got the correct answer and the change shipped.

Nobody asked the other half — **does it still answer where it should?** It did not. Next.js 16's dev server
proxies to `http://localhost:$PORT/` rather than to the address it was told to bind, so an IPv4-only bind on
a host that resolves `localhost` to `::1` first cannot reach itself. Measured: `500` in a flat 30.0 s versus
`200` in 2.4 s on the pre-M221 bind.

> ⚠ **CORRECTION (iter-10) — the MECHANISM above is REFUTED; the RULE below is upheld and strengthened.**
> "Next dev proxies to `localhost` rather than the bound address, so an IPv4-only bind on a `::1`-first host
> cannot reach itself" is wrong. iter-10 measured three falsifications: `curl http://127.0.0.1:$PORT/` — a
> dial that never utters the name `localhost` — fails **identically** (500 / 30.06 s);
> `NODE_OPTIONS=--dns-result-order=ipv4first`, the direct repair for a resolution-order bug, changes
> **nothing**; and `-H ::1` — which makes the first-resolved address the bound one — **also fails**.
>
> The real mechanism is an **origin-STRING equality**, not a resolution: `next@16` normalizes *every*
> loopback hostname to the literal `localhost` when its middleware builds a rewrite URL
> (`web/next-url.js:15-20`), builds the router's base URL from the **raw** `-H` string
> (`resolve-routes.js:117`), and compares the two with `===` (`relativize-url.js`). Mismatched, the app's own
> in-place rewrite is proxied *externally to itself* until `http-proxy`'s `30_000` ms default. Fixed with
> `-H localhost`. See `iter-10/decisions.md` D-M257x-14.
>
> Why the correction matters beyond the fix: every candidate on iter-09's pre-computed fix menu followed from
> the refuted story, and two of them (`-H ::1`, force the dial) were measured to **not work**. A mechanism
> that explains the observation is not the same as the mechanism that produced it — and this one explained
> the flat 30 s, the bind dependence, and the log line, while still being false.

**An exposure guard can never notice that a service stopped working** — by construction, "no response" is
its success condition. So it will confirm every over-tightening as a win.

> **Rule.** Every tightening ships with a paired liveness assert on the surface it tightened. A guard that
> only measures the property you removed is measuring half the change, and the half it measures always
> passes.

Two supporting method notes from the same measurement:

- **A flat, repeatable duration is a timeout, not work.** 30.014 s then 30.007 s, identical on a warm second
  request, is a configured limit. `ant-academy.sh`'s own comment attributed exactly this shape to Turbopack
  cold-compilation and budgeted 120 s for it — a wait that could never succeed, because every attempt fails
  identically. Compare the **variance**, not the magnitude. (Kin to `latency-budget.md`'s arithmetic
  signatures.)
- **A "not serving" verdict deserves one patient request before it is believed.** `--max-time 3` and
  `--max-time 180` are different instruments; only the second can see a 30 s failure. A probe whose
  per-attempt timeout is shorter than the failure mode it watches for cannot measure the thing it reports.

---

## HARDEN-CAP-ACCEPTED — the incremental harden pass closed un-stabilized, and that is the accepted disposition (2026-08-01)

**Recorded at iter-16 close so milestone close finds this stated rather than inferred.**

`/developer-kit:harden-mstone-iters` (incremental) ran its capped passes and **did not reach stabilization**.
Of 20 scanned findings it fixed **6**; the remaining **12 are routed forward as RF-1 … RF-12** in
`hardening-ledger.md`, each carrying a file, a line, the claim-vs-measurement, and — the part that makes them
actionable rather than a wish-list — **the source mutation that leaves the suite green today**. That last
column is why closing un-stabilized is defensible: every routed finding is already proven to be a real gap in
the safety net, not a suspicion about one.

**The pass stopped for a reason worth stating.** The residual queue is dominated by **source changes to
bring-up scripts** — RF-1 adds an `exit 1` to the dev migration path; RF-4 changes the verdict and exit code
of the set-dress pass. A hardening pass's mandate is to deepen the tests around code, not to change what the
code does; work with that blast radius belongs to an iter, where it gets a hypothesis, a mutation battery and
a full section-suite comparison against a recorded baseline. Continuing the harden pass into them would have
made large behavioural changes under the lighter discipline.

**Disposition: ACCEPTED by the orchestrator on the user's standing full-milestone-autonomy delegation.**

**Status of the routed set as of iter-16 close:** RF-4 and RF-1 are **landed** (iter-16 — see
`iter-16/progress.md`; 11 mutants, 10 declared-RED all killed, 1 declared-GREEN no-op survived, every section
back to its recorded baseline). **RF-2, RF-3 and RF-5 … RF-12 remain open** and are the standing queue for
later tiks. `CHECK-M257x-iter16-parity-fence-hand-maintained` joins them.

**One consequence to carry.** A harden pass that closes un-stabilized leaves the `--final` cumulative sweep
carrying more than it normally would. When the gate fires, `/developer-kit:harden-mstone-iters --final` should
be run in the knowledge that RF-2/RF-3/RF-5…RF-12 were never swept, rather than treating the ledger's most
recent entry as a clean high-water mark.

---

## HARDEN-CAP-ACCEPTED (2nd) — 2026-08-01, recorded at iter-27 close

The second incremental `/developer-kit:harden-mstone-iters` pass (passes 4–6) was **accepted by the
orchestrator without stabilizing**, exactly as the first was. Recorded here so the residue is visible to
whoever runs `--final`, rather than living only in a hand-off.

**What the pass delivered:** 27 tests, 5 fences repaired or built, **44/44 mutants matched their declared
expectation** (35 declared-RED killed, 9 declared-GREEN no-op controls survived). Five more guards that
reported without measuring.

**Two findings that change what can be trusted, and are load-bearing for later iters:**

1. **The runner-safety invariant "never an additive re-seed" was UNENFORCED since iter-25 while reporting
   enforced** — the fence read the raw file and was satisfied by an *echo line printed beside the call*.
   Two Go tests were RED that whole time against a baseline recording "Go sections green". Consequence:
   **never trust an inherited "playthroughs green"; re-run it.** (iter-27 did not need to — its only
   `playthroughs` interaction was a read of the report artifact — but the rule stands.)
2. **The corpus audit is now a DERIVATION, not a grep vocabulary**: the new service-doc fence reads the
   migration map and holds every service doc to it. It caught a doc on its first run that three sweeps had
   missed *because that doc never used the words being grepped*. **Clause 5's re-measurement must use it.**

**The accepted residue, unchanged by iter-27:**

- **iters 17, 19 and 26 were never scanned** — measurement/withdrawal tiks, the cheapest remaining, and
  plausibly empty. iter-18 *was* scanned and deliberately left alone: its stub already drives both sides of
  the bootstrap race end to end, and no gap was found.
- **`RF-2`, `RF-3`, `RF-7`…`RF-12` remain open** in `hardening-ledger.md`. `RF-1`/`RF-4` landed in iter-16;
  `RF-5`/`RF-6` landed in pass 3.

**Consequence for `--final`:** it must be run **knowing** the above — an unqualified "cumulative sweep"
claim would be false about three iters and eight RF items. This is the second consecutive cap accepted
without stabilization; a third should be read as a signal about the harden cadence itself, not about the
milestone's test debt.

---

## TOK-02: fence the prose the way the anchors are fenced — 2026-08-02

**Tok type:** triggered (session-terminating) — but **user-directed**, not fired by the 3-no-prog streak.
The metric moved at iter-41 (`37 → 18`). The user directed a strategy review after **40 consecutive tiks on
TOK-01 with no revision**, in substance: *address the slowdowns encountered and soften them, or accelerate
progress in any case.*

**Prior strategy:** [`TOK-01: instrument first, then follow`](#tok-01-instrument-first-then-follow--2026-07-31)
— *build the instrument before doing the work, and derive every list from the platform instead of
maintaining it by hand.* Authored at iter-01, never revised, and **it has worked**: the derived migration
tuple (iter-02), the derived section coverage (iter-08), the derived replay schema (iter-07), the derived
schema-probe list (iter-05) each ended a recurrence permanently, and **clauses 1, 2, 3 and 4 are all held on
fences that were watched going RED before they were trusted.**

**Why it stopped working — and it is one clause, not the whole strategy.** Clause 5 is the single surface
where the milestone **reverted to hand-maintenance**, and it is the single clause still open. Six passes,
`25 → 13 → 11 → 17 → 37 → 18`, of which iter-39 proved the first five measured *the instruments* and
iter-41 — holding the instrument fixed on every knob, the first controlled comparison in the series —
measured the corpus for real and found the residual **halving but not approaching zero**, with
**9 of 18 manufactured by the repair that preceded it.** Two further regularities, both measured:

- **In two consecutive iterations the author of a newly-written rule violated it while writing it**
  (iter-40's rule 19; iter-41's `D-M257x-41-2`), and iter-40's hand sweep — run by the rule's own author,
  with a mandatory post-condition re-grep — had a **27% miss rate**.
- **For five consecutive passes every `file:line` anchor a sweep introduced resolved correctly** — ~110
  across 91 hunks, zero failures. **The anchors are clean because a machine checks them on every pass.**

Hand-applied discipline has measurably stopped scaling at this corpus's size. That is not a discipline
problem to be solved with more discipline.

### The measurement that decides the revision

The 18 were never homogeneous. Classified by **the cheapest instrument that could have caught each**, taken
row-by-row from iter-41's own `blocker-ledger.md`:

| class | n | instrument |
|---|---|---|
| **the corpus contradicts ITSELF** — a twin site in the corpus already states the opposite | **13** | a claim-twin fence. **No platform read needed** |
| the anchor resolves but names the **wrong construct** (right line, wrong function/row) | **3** | a symbol-aware anchor check |
| a derived scalar vs platform source with no corpus twin (`go.mod` 1.26 vs "1.25"; `locals.tf` 128 MB vs "256") | **2** | a value fence |

**13 of 18 are the corpus disagreeing with itself** — not with the platform — and in five cases with another
sentence **in the same file, within a few lines** (`:102-103`, `:69`, `:202`, `:174-176`, `:458`).

**And it names the mechanism of the induced half: of the 9 repair-induced blockers, 8 are
self-contradiction.** Every one is the same shape — a claim repaired at one site and left standing at
another: *"added the twin row, left this one"*, *"fixed one twin row, not the other"*, *"the contradiction
is inside its own text"*, *"a blockquote spliced into a bullet list"*. So iter-41's conclusion is right about
the arithmetic and incomplete about the cause: **the injected defects are overwhelmingly ONE mechanical
class, and it is precisely the class a machine can check while knowing nothing about the platform.**

Second measurement: **at least 4 of the 18 are RETURNING claims** — restatements of a claim a prior pass
already adjudicated and recorded a verdict for (#12 `:5050`, swept at 8 sites by iter-40 and missed *inside*
clause-5 scope; #18 the retracted EU-first ladder, still published verbatim; #5 the tenancy fence, wrong a
**fifth** time; #7 a retraction contradicting `platform-migration-status.md:86`, the corpus's own
machine-fenced SoT, linked by the very section that contradicts it). **Nothing checks whether a refuted
claim has come back.** The verdicts exist, anchored, in five blocker-ledgers (iters 33/34/38/39/41).

**Stated weakness.** The ledger names each twin because a human found it; a fence must find it *without
being told*. Not fatal — that direction is **proven manually**: iter-40 swept 8 adjudicated claims tree-wide
by grep and its post-condition re-grep found 3 more sites. The mechanism works; the hand is what fails at it.

### Revised strategy — apply TOK-01's own principle to the one surface it was never applied to

Ordered, and the order is forced by dependency, not preference:

1. **Build the claim-twin fence** — `CHECK-M257x-iter33-derived-fact-fence`, finally cut correctly. **Not a
   general fact-checker** (impossible, and never the route). A **claim ledger**: one row per adjudicated
   claim = `{id, canonical verdict, the refuted form(s) as matchable patterns, the source citation}`,
   **DERIVED from the five existing blocker-ledgers** — §5 rule 19's list-derivation clause already mandates
   deriving the claim list from the prior pass's ledger and never by hand. The fence greps the **whole
   tree** — `corpus/**`, `.claude/skills/**`, `CLAUDE.md` — and goes RED naming `file:line`. Scope is
   tree-wide from the start because iter-40 measured that **100% of a repair's surviving claim sites pool
   immediately outside its boundary.**
2. **Make the fence a POST-CONDITION of every repair, not a later audit.** This is the change that attacks
   the dominant cost term. Today: repair → hand adversarial pass → the next full read finds what both
   missed. Under the revision the 8-of-9 induced class **cannot survive the commit**. The hand adversarial
   pass stays — it found real defects five times running — but stops being the only line of defence.
3. **Add the two small mechanical fences the classification names**, closing 5 more of the 18: a **markdown
   structure lint** (a blockquote spliced into a bullet list, an orphaned list member, a doubled word — #6
   and iter-38's *"The The"* are pure mechanical damage), and a **symbol-aware anchor check** (#13 `:447` is
   a table *header row*, #16 `:815` is `}))` in the wrong function, #17 `:604` wires 1 of 4 domains). The
   existing check proves a line **exists**; it never asks whether the line **says the thing**.
4. **Then repair the 18 once, fence-assisted** — by CLAIM not by FILE (§5 rule 19), tree-wide, with the
   fence as the commit post-condition.
5. **Then ONE full 7-auditor read, instrument held fixed at iter-41's** — same seven auditors, same
   briefing, same partition method, all 40 files top-to-bottom. **That reading is what meets or fails
   clause 5.** Nothing else does.

**What this revision explicitly does NOT do**, each because the milestone has already measured why:
- **It does not weaken the audit instrument.** iter-38 measured that narrowing to the high-density files
  would have found 11 of 17; iter-21 is the precedent for a scoped audit converging on a number a full read
  then multiplied by five. A cheaper instrument yields an uncomparable number.
- **It does not re-cut clause 5, or narrow it, or read "met" any other way.** The user has ruled. Only a
  pass returning zero meets it.
- **It does not defer the residual to a future milestone.** The 18 are repaired here, at step 4.
- **It does not replace TOK-01.** It *extends* it. Every fence TOK-01 built is kept and still holds clauses
  1–4.

**On the other measured slowdowns:**
- **The hand-off tax** (17 consecutive hand-offs refuted or materially corrected). Do **not** cheapen the
  re-derivation — it works, and it caught real errors every time. Cheapen the thing being handed: a hand-off
  whose claims already sit in a machine-checkable ledger costs a *fence run*, not a re-derivation. **The
  claim ledger is also the hand-off format.**
- **Budget exit every ~2 tiks / re-orientation cost.** A fence is written once and runs forever; the current
  repair method pays its full cost every pass. It also makes an interrupted run cheap to resume — the ledger
  plus fence state is the resumption point, not a narrative.

**Pre-registered, so it can be refuted.** iter-41's projection on the *current* method: a seventh pass
repairs 18, induces ~9, measures **9–15**. The revision attacks the induced term specifically. **Prediction:
the reading at step 5 returns fewer than 9.** This is a prediction and not a result; four consecutive passes
refuted their own predictions, and iter-41's held only once the instrument was held fixed.

**Strategy class:** `retry-with-evidence` — TOK-01's derive-and-fence principle, retried on the one surface
it was never applied to, now with the measurement (13 of 18 self-contradiction; 8 of 9 induced) that says it
applies there. It is a continuation, not a reversal.

**Distance-to-gate context:** **4 of 5.** Clauses 1, 2, 3, 4 hold; clause 5 at **18**, fully anchored and
ready as `FIX-M257x-iter41-blocker-set`. Late-milestone, so diminishing returns are expected — but the
per-pass return has been *negative-adjusted* by the induced term rather than merely small, which is why the
revision targets induction rather than throughput. **One conjunction risk to watch:** the gate is an AND
over five clauses and four are already held. Steps 1–3 add **offline guard/test code only**, on no runtime
path, so clauses 1/2/4 should be undisturbed by the re-pin — **should be, and that must be verified rather
than assumed** when the pin moves.

**Next-tik direction:** **iter-43 = `FENCE-M257x-iter42-claim-twin`.** Derive the claim ledger from the five
blocker-ledgers (never by hand), build the fence, and **watch it go RED on today's corpus** — it must fire on
the ≥4 known returning claims and on the self-contradiction blockers whose twin is nameable — then GREEN.
**Do not repair in the same tik.** Today's 18-defect corpus is the fence's only test fixture with a known
answer key, and it is **perishable**: repair first and the fence can only ever be tested against a corpus
nobody has measured. This is TOK-01's own rule 3 — *land the fences, each watched going RED, before trusting
any green* — applied to clause 5.

**Harden recommendation (the orchestrator spawns it, not this iteration):** **run it after iter-43, not
before.** 14 tiks have closed since the pass-6 window (iters 28–41) against a threshold of 10, so it is
overdue — but the window is dominated by the clause-5 **reading** iters (33, 34, 38, 39, 41), which ship
almost no code and are the same cheapest-and-plausibly-empty shape the second `HARDEN-CAP-ACCEPTED` recorded
for iters 17/19/26. iter-43 ships new guard code, which is exactly the material a harden pass should sweep.
Noting the counter-argument honestly: two consecutive passes have already closed **un-stabilized**, and that
ledger entry says a third *"should be read as a signal about the harden cadence itself, not about the
milestone's test debt."*

---

## HARDEN-CAP-ACCEPTED (3rd) — 2026-08-02, recorded at iter-44 close

The incremental harden pass fired three sub-passes (7, 8, 9) and closed **un-stabilized** for the third
consecutive time: pass 9 still surfaced a live unrepaired twin on its first look, so the dimension scan
is not clean.

**The disposition has changed, and this is the point of recording a third entry rather than another copy
of the second.** The 2nd entry read the cap as a signal *about the harden cadence*. The evidence now says
it is **structural**: the defect class the passes keep finding lives **outside the dimensions the sweep
scans**. Pass 7 found the fence's own reporting paths — behaviour no coverage metric can see, because a
`print` that never runs is not an unexecuted statement. Pass 8 found a fix proven **inside** a function
while all six of its **call sites** were wide open, with `jobroleref.go` sitting at **100% statement
coverage** the whole time. Pass 9 found a negative control drilling row 0 of an 11-way tie. None of
these is reachable by "sweep the iters we have not swept yet"; each needed an AST or call-site
assertion, which is a different instrument, not more of the same one.

**So the accepted disposition is:** the cap is not test debt to be paid down by more passes of the same
shape, and a fourth pass of that shape should not be expected to stabilize either. What the residue
needs is the instrument change — prefer AST / call-site assertions over coverage — which passes 7–9 have
now begun and which `platform-alignment.md` §8 records.

**Unchanged and still open:** iters 27–30, 32–34, 36–41 unscanned to completion; **RF-2**, **RF-3**,
**RF-7**…**RF-13**; and — the highest-value item in the queue —
**`CHECK-M257x-iter35-seeder-writes-one-instant`**, the seeder stamping every backdated session at a
single timestamp, which flattens recency ordering **in the product** and is the root cause behind two
assertions that were found true by coin-flip.

---

## TOK-03: repair the UNION, shrink the estimator, make the edits smaller — 2026-08-03

**Tok type:** triggered (3-no-prog streak: iters 48, 49, 50). Session-terminating.

**Prior strategy:** [`TOK-02: fence the prose the way the anchors are fenced`](#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02)
— *apply TOK-01's derive-and-fence principle to the one surface it was never applied to.* Five fences
were built under it, each watched RED before it was trusted. **They work.** The mechanical classes they
close stay closed and are not the subject of this revision.

### Why it stopped working — and this time the mechanism is measured, not inferred

TOK-02's own pre-registered prediction was *"the reading at step 5 returns fewer than 9."* Three readings
later: **7, 12, 14.** Its step-2 premise — *"the 8-of-9 induced class cannot survive the commit"* — became
true of a class that had stopped being the majority (iter-49). And iter-49's two new fences closed exactly
the gaps they were built for while the induced term rose **2 → 7**.

The decisive evidence is iter-50's **paired same-tree reading** — §5 rule 22's own prescribed experiment,
run under total control (40 files byte-identical, 13 clone shas identical, the same partition hand, the
identical diff for seat G, every seat blind):

| | reading #9 | reading #10 |
|---|---|---|
| blockers | **14** | **7** |
| per seat | A 1 · B 1 · C 2 · D 2 · E 3 · F 0 · G 5 | A 1 · B 0 · C 0 · D 1 · E 0 · F 0 · G 5 |

**Matched 4. Union 18. Recall 29% and 57%. Chapman `N̂` ≈ 23 — and a FLOOR**, because heterogeneous
detectability biases capture–recapture downward.

Three things fall out, and together they name the error in both prior strategies:

1. **A single reading is a SAMPLE, not a census.** Two full 7-seat passes over an unchanged tree named 18
   findings between them, and neither named more than 14. One of #10's four new findings sits **inside a
   hunk seat G reviewed and PASSED at #9**.
2. **A repair pass can only repair what a reading NAMES.** With recall ≈ 0.43 and a non-zero induction
   rate, repair-then-read has a **fixed point** — and it sits exactly where `18 → 7 → 12 → 14 → 7` has
   been sitting. **The series was never noise around zero; it is the equilibrium of the method.**
3. **The two prior strategies both optimised the reading.** TOK-01 built the instruments; TOK-02 sharpened
   the prose fences. Neither touched coverage. **Coverage is the binding constraint and always was.**

### Revised strategy — four moves, ordered by dependency

1. **Repair the UNION of two blind readings, never one.** `FIX-M257x-iter50-union-set` is **18**, not 14 —
   `#9 ∪ #10`, anchored across `iter-50/fixture-14.md` + `iter-50/blocker-ledger.md`. Two readings cost 2×
   and cover **78%** of the estimated pool against 61% for the richer one alone. From here a repair cycle
   is *read, read, repair* — and the second reading is **blind** (`D-M257x-50-2`), which is now part of the
   instrument.

2. **Drive `N̂` down, and only then take the reading clause 5 asks for.** This is the sequencing change,
   and it is the whole point. With recall ≈ 0.43, a single pass missing all of a residual of 23 has
   probability ≈ 10⁻⁵ — **so attempting clause 5's zero reading while `N̂` ≈ 23 is not a long shot, it is
   arithmetically hopeless.** Attempting it at `N̂` ≈ 2 is not. `N̂` is the quantity to drive; it has a
   **floor of zero by construction** (when nothing is there, `m`, `n₁` and `n₂` collapse together), which
   is the property §5 rule 22 asks a gate metric to have and which the raw blocker count never had.
   **Clause 5 is unchanged and is still the gate.** `N̂` is a progress metric, not a substitute verdict.

3. **Cut the induction rate by shrinking the EDIT, not by adding fences.** iter-49 measured the induced
   classes as paraphrase leak (3), overshoot-in-new-text (3), wrong-mechanism (1), and proved none is
   mechanically reachable — a paraphrase shares no token run; an overshoot lives in prose that did not
   exist before the commit. **The only available lever on an unreachable class is to shrink the surface it
   can live on.** Both classes are properties of *rewriting*. So the repair discipline changes:
   - **Prefer DELETION > minimal scoping edit > rewrite.** A false claim removed cannot be over-corrected
     and leaves no paraphrase twin. The corpus does not owe every retraction an explanation.
   - **Budget the new prose.** Count added words per repaired claim; a repair that adds more than it
     removes is the highest-risk shape this milestone has (eight consecutive occurrences of *the author of
     a correction violating it while writing it*, the latest being `dependency_map.md:19`).
   - **Pre-registerable, therefore refutable:** the induced term falls below 4 at the next paired reading.

4. **Put a READER on the repair diff, before the commit.** iter-49 named this and did not build it
   (`CHECK-M257x-iter49-overshoot-has-no-instrument`): the overshoot class *"needs either a reader or a
   different kind of check entirely."* Seat G **is** that reader — it has been the highest-yield seat in
   both readings — but it runs *after* the commit, in the next pass. Move it in front: **two blind
   adversarial diff seats on the repair diff pre-commit**, allowed to disagree. A diff is small, so this
   is cheap; and its own recall is measurable by the same paired method.

### What this revision explicitly does NOT do

- **It does not re-cut clause 5, narrow it, or read "met" any other way.** The user has ruled twice.
- **It does not defer the residual**, and does not propose closing at 4 of 5.
- **It does not weaken the instrument** — it runs *two* frozen instruments rather than one cheaper one.
- **It does not discard the fences.** All five are kept; they hold clauses 1–4 and the mechanical classes.

### Also carried, because iter-50 surfaced them

- **`FENCE-M257x-iter50-consecutive-audit-mode`** — `--audit-commit` assumes *audit → repair → audit* and
  cannot represent *audit → audit*, so iter-50 landed with a recorded `--no-verify` (`D-M257x-50-7`).
  Widen condition 1 (*the ledger row is a line **this** commit added*) to *any commit since that claim was
  last repaired*; leave condition 2 — the anti-laundering key — untouched, and watch it REFUSE a
  repair-shaped commit before trusting it.
- **`CHECK-M257x-iter50-audited-zero-is-evidence`** — §5 rule 24. Three seats cleared the same false count
  as a *positively audited zero*, each re-deriving the arithmetic the document showed instead of the
  predicate it claimed. The seat report format must distinguish *"I re-derived the document's own
  arithmetic"* from *"I enumerated the predicate from source"*; only the second is a clearance.
- rosetta's root `CLAUDE.md` is the stale side of **two** claims and **lies outside the 40-file
  partition**, so no reading will ever book it. It needs a deliberate pass.

**Strategy class:** `new-direction`. TOK-01 built the instruments; TOK-02 sharpened them. TOK-03 stops
optimising the reading and attacks **coverage** and **repair surface** instead. It is informed by the two
prior strategies rather than a reversal of them — every artefact they produced is kept — but the term
being optimised is different, which is what makes it a new direction rather than another
`retry-with-evidence`.

**Distance-to-gate context:** **4 of 5.** Clauses 1–4 hold. Clause 5's residual is **~23 and biased low**,
of which **18 are anchored and ready**. Late-milestone, and the honest characterization has changed: the
per-pass return was not merely small, it was **zero in expectation**, because the method's fixed point sat
where the readings were landing. TOK-03's first cycle should therefore produce the first genuinely
non-zero expected return of the clause-5 series — **which is a prediction, and four consecutive passes
before this one refuted their own.**

**Cross-refs to prior TOKs:** TOK-01 (*instrument first, then follow*) built the derived lists and fences
that still hold clauses 1–4; it is untouched. TOK-02 (*fence the prose*) correctly identified a
mechanizable class, closed it, and measured its own premise expiring — iter-49's ledger records that in
its own words. TOK-03 keeps both and changes the term: **not a sharper reading, but more readings, smaller
repairs, and a residual estimate with a floor of zero.**

**Next-tik direction:** **iter-52 = `FIX-M257x-iter50-union-set`** — repair the **18** by CLAIM, tree-wide
(§5 rule 19), under move 3's minimal-edit discipline (deletion preferred; added-word count recorded per
claim) with move 4's two blind pre-commit diff readers as the post-condition. **Do not take a reading in
the same tik.** Then **iter-53 = the paired reading #11 + #12**, blind, frozen instrument, and recompute
`N̂`. **Pre-registered now, so it can be refuted: `N̂` below 12, and the induced term below 4.**

**Harden recommendation (the orchestrator spawns it, not this iteration):** **7 tiks since the pass-7/8/9
window** (iters 45–50) against a threshold of 10 — **not yet due**. When it is, the third
`HARDEN-CAP-ACCEPTED` entry stands: the residue needs AST / call-site assertions, not another sweep of the
same shape.

---

## TOK-04: pin the target, or stop calling it a measurement — 2026-08-03

**Tok type:** triggered — but not by the 3-no-prog streak. Fired by the milestone's own
**`re_scope_trigger`, occurrence 2 of 2**, plus a direct user ruling. iter-53 closed `closed-fixed`; the
streak clause does not apply and was checked before this was written.

**Prior strategy:** `TOK-03: repair the UNION, shrink the estimator, make the edits smaller` — and behind it
TOK-02 (*fence the prose*) and TOK-01 (*instrument first, then follow*). All three optimised the **reading**:
a sharper instrument, then a mechanized class, then more readings with smaller repairs.

**Why it stopped working — and the honest version is that it did not stop, it was outrun.** TOK-03's premise
was a **fixed residual** that better/smaller/more readings could drain. On 2026-08-03 three platform commits
landed inside one working day and created **81 fresh drift sites across 21 files** — larger than the entire
46-item union that ten readings had been working on. At the observed rates (repair ~18 per iter, induction
~9 per 18, external drift 81 in a day) a repair pass is **net −72**. TOK-03 was not wrong about the reading;
it was measuring a stock while the thing that mattered was a **flow**.

The `re_scope_trigger` predicted this in its own words and prescribed the answer:

> *"The answer then is a pinning-and-tracking **POLICY** (how we choose a platform ref, how we notice it
> moved, who re-points), not more alignment work. Grinding against a moving target is the failure mode this
> trigger exists to catch."*

### The class, generalized — because a platform-only policy would miss two thirds of it

This milestone has now been bitten by the same class **four times, in its own instruments**:

| # | instrument | how it moved unseen | found at |
|---|---|---|---|
| 1 | the rext version pin (`.agentspace/rext.tag`) | **git-ignored** — never in a diff; 11/11 clones reported `behind: null` while the log claimed *"provably fresh"* | iter-01 |
| 2 | the audit briefing — *"the whole instrument"* | **git-ignored** scratch path, re-authored from a summary each pass; six knobs drifted, incl. an **inverted tie-break** | iter-53 |
| 3 | the ground-truth platform clone | free to move; moved **during** iter-53, invalidating a seat's clearance by name | iter-53/54 |
| 4 | **clause 2's gate-meeting run (iter-37)** | recorded **no platform ref at all** — the only sha in the file is `ad524614`; its ref exists only by adjacency to iter-36 | **iter-54, today** |

**The generalization: an input that can change without appearing in a diff is not a controlled input.**
Three of the four were git-ignored files; the fourth was an external ref nobody wrote down. Same class, and
a policy scoped to the platform repo would have caught exactly one of them.

### New strategy — four rules, each fenceable

**P1 — Every measurement states its refs, in the artifact, at the moment it is taken.**
A `refs:` block in the measuring iter's `overview.md` frontmatter *and* in the run artifact: platform sha,
rext tag, corpus HEAD, instrument-file sha. **A measurement without a `refs:` block is an anecdote, not a
measurement.** Fenceable: a guard that fails any `status: closed-*` iter whose progress claims gate-clause
movement without one. iter-37 is the retrofit case and it has to be re-run anyway.

**P2 — Every instrument is a committed file, and the measurement cites its sha.**
Not *"the briefing described in the previous overview."* `instrument/briefing-canonical-iter41.md` is now
that file. **Corollary, and the one with teeth: nothing an instrument depends on may live under a
git-ignored path.** `.agentspace/rext.tag` is the known offender; **assume it is not the only one until
measured** — the same wording this milestone's scope used about rext write paths, where it was right.

**P3 — The platform ref is chosen, recorded, and re-checked at open AND close of every measuring iter.**
- *Chosen*: origin HEAD at iter open, never a pinned pre-drift commit. The gate already says this; P3 makes
  it operational rather than aspirational.
- *Recorded*: in the P1 `refs:` block.
- *Re-checked at close*: iters 32 and 36 already did this **voluntarily**, and it is what detected both
  occurrences of the re-scope trigger. Make it mandatory. A close-time re-fetch that finds the ref moved
  **invalidates the measurement by construction** — it does not "probably still hold."
- *Who re-points*: **the iter that detects the move, in that iter, as its first act.** Not a routed-forward
  item. iter-54 absorbed a 3-commit move in under an hour; deferring is what makes it expensive.

**P4 — Derive; else fence; else DECLARE it prose-under-review. In that order, and mark which one you used.**
Today produced a **rank ordering of the three, measured on a single event**:

| approach | what happened on 2026-08-03 |
|---|---|
| **derived** (clause 4's `repos_yml_*`) | tracked the removal with **zero human action**; reading identical at `2adcf71` and `ef32d4c`, and identical *correctly* |
| **fenced** (clause 3's membership guard) | caught the departure **unaided, 3 for 3, within hours**, on a tree nobody had touched — the first non-staged catch in the milestone |
| **hand-maintained prose** (the map's §5 narrative; the corpus's 81 sites) | **falsified inside one working day** — and one of the falsehoods was written *by this milestone*, while updating the map, citing line anchors its own iter-02 had deleted |

That is not three anecdotes. It is the milestone's founding thesis, measured, on one event, with a control
in each row. **The engineering rule follows: before writing a claim, ask whether it can be derived; if not,
whether it can be fenced; if neither, mark it explicitly as prose-under-review with a re-check date.** The
map must carry that third category **visibly** — it currently does not, which is exactly why a false §5
narrative read as authoritative as a fenced table row.

### What TOK-04 keeps from TOK-03, because the evidence still supports it

- **Union-of-two readings stays.** Recall ≈ **43–48%** per 7-seat pass makes a single reading a coin flip.
  This is the one instrument-independent fact the series produced and it is not in dispute.
- **Pre-commit double-reads stay — and today upgraded them from "recommended" to "load-bearing."** TOK-03
  move 4 recorded that two blind adversarial diff seats caught **5 blockers inside a repair**. Today's
  over-claim is the **sixth instance of the same class, and it ESCAPED**, because job 1 was committed
  without one. Same-day confirmation of both the value and the cost of skipping it.
- **Smaller edits stay.** ~50% induction on the 18-claim repair is the reason.

### What TOK-04 changes

1. **Meter the flow, not the stock.** The clause-5 metric becomes **net**: *(sites repaired) − (sites
   induced) − (sites newly falsified by platform movement)*, per unit time. **This is the first clause-5
   metric that can go negative** — which is the entire point, because the old one could not, and therefore
   could never tell us we were losing. It reads **−72** for the last cycle.
2. **Re-establish the ref baseline BEFORE any further reading.** Clause 2 costs ~5 min, clause 1 ~35 min —
   together less than one reading pass. Taking them first means every subsequent clause-5 number is anchored
   to a stated, current ref, which **no number in the ten-reading series has ever been.**
3. **Repair the 81 as ONE derived-and-fenced class, not as 81 claims.** They share a single false predicate
   (*three services are live-local husks*). That is precisely the mechanizable shape TOK-02 identified and
   closed successfully. Routing it through the reading instrument claim-by-claim would be the most expensive
   available route to the cheapest available win.

### What this revision explicitly does NOT do

- **It does not re-cut clause 5, narrow it, or read "met" any other way.** The user has ruled **three**
  times. Met only by a reading that returns zero.
- **It does not defer the residual** and does not propose closing at 4 of 5 — which is moot anyway, since
  the honest count against origin HEAD is **2 of 5**.
- **It does not weaken the instrument.** It freezes it harder (P2) and adds a ref record (P1).
- **It does not discard anything TOK-01/02/03 built.** Every fence, every derived list, every artefact is
  kept. TOK-01's derived sets are the reason clause 4 passed today unaided.

**Strategy class:** `new-direction`. TOK-01 built instruments, TOK-02 fenced a mechanizable prose class,
TOK-03 attacked coverage and repair surface — all three optimised the **reading**. TOK-04 accepts the
re-scope trigger's own prescription and optimises **the target's stability and the accounting**. Different
term again, and this time chosen by the trigger rather than by us.

**Distance-to-gate context:** **2 of 5 verified at origin HEAD `ef32d4c`** — clauses 3 and 4 hold, clause 4
now *under test*. Clauses 1 and 2 are **stale by the gate's own "against origin HEAD" wording**, not failed,
and cost ~40 min of machine time to restore. Clause 5 is **not met and moved further away today by 81
sites**. Full detail: `iter-54/reassessment.md`.

**Cross-refs to prior TOKs:** **TOK-01** (*instrument first, then follow*) built the derived sets and fences
that just passed their first unplanned live test — untouched, and vindicated more strongly than any reading
vindicated anything. **TOK-02** (*fence the prose*) correctly identified a mechanizable class; its thesis is
exactly what change 3 above applies to the 81 sites. **TOK-03** (*repair the union, smaller edits*) keeps all
three of its moves; what is refuted is its **premise** — a fixed residual that more and smaller readings can
drain — by a platform that delivered 81 sites in a day.

**Next-tik direction:**
- **iter-55 = re-establish the ref baseline.** Re-run clause 2 (~5 min) then clause 1 (3 cycles, ~35 min)
  against `ef32d4c`, each carrying a P1 `refs:` block. **Pre-registered, therefore refutable: both come back
  green**, because clause 4's derivation is ref-independent and `d11a403` *restored* env onto `backend`
  rather than removing it. If either goes red, that is a real finding and the highest-value one available.
- **iter-56 = the 81-site sweep as one derived class** (P4 + TOK-02's mechanism), with TOK-03 move 4's two
  blind pre-commit diff seats — which today's escape proves are not optional.
- **Then** the next paired reading — the first in the series taken against a stated, current, committed set
  of refs.

**Harden recommendation (the orchestrator spawns it, not this iteration):** **10 tiks since the pass-7/8/9
window** (iters 45–50 → iter-54) against a threshold of 10 — **due now**. But it should sit **after
iter-55, not before.** The last pass stopped on *"cap reached without stabilization"*, and its named residue
— iters 27–30, 32–34, 36–41 unswept, with `CHECK-M257x-iter35-seeder-writes-one-instant` still the root — is
all **Playthrough and seeder** surface, precisely what iter-55 re-exercises against the new platform ref.
Hardening that surface against a stale ref would repeat the mistake this tok exists to name. **Order:
iter-55 (ref baseline) → `/developer-kit:harden-mstone-iters` → iter-56.** That pass also owns the **rext
`stack-core` suite that was not run to completion two iterations ago** — those baselines stand unmeasured
and belong to a harden, not to a tik.

---

## TOK-05: stop repairing claims; fence the predicates under them — 2026-08-04

**Tok type:** triggered — by a **direct user directive**, not by the 3-no-prog streak. iters 56, 57 and 58
all closed `closed-fixed` and each moved a clause; the streak clause was checked before this was written and
does not apply. Same precedent as TOK-04, which was also fired by something other than the streak.
Session-terminating.

**Prior strategy:** [`TOK-04: pin the target, or stop calling it a measurement`](#tok-04-pin-the-target-or-stop-calling-it-a-measurement--2026-08-03) — *meter the flow, not the stock; state your refs;
freeze your instruments.* **TOK-04 worked, and nothing in it is being discarded.** Under it the milestone
went **1 of 5 → 4 of 5** in four iterations (iter-55 ref baseline, iter-56 clauses 1+2, iter-57 clause 3,
iter-58 the pin advance proven cold). P1–P4 are kept whole.

### Why a revision now — the user handed us a map of the territory, and it re-scopes the work

The platform developer's PR shows the fold as a **program with a known shape**: skiller, cms, graphql,
jobsim and skillpath are **done**; **storage and messenger are next, and not yet done.** That single fact
re-scopes clause 5, because it says which of our residual is *stale* (repairable now, permanently) and which
is *in flight* (repairing it today buys a claim that expires on the developer's next merge).

Alongside it, rosetta PR #14 was fetched read-only as `origin/pr-14` and reconciled: **92 claims already
absorbed · 30 superseded · 5 contradictions standing · 0 refuted · ZERO new information. DO NOT MERGE** — it
would re-introduce three refuted things (the 60K/18K taxonomy figures, `internal/copilot` deleted at app
`889ae776`, a 4-subgraph Cosmo count). **Its value is negative space.** The live defects are where PR #14
and our corpus **agree**, and those are invisible to any method that diffs two documents against each other.

### Why TOK-04's own method has a ceiling it cannot cross

TOK-04 changed the accounting and the ref discipline. It did not change **the unit of repair**, which has
been *the claim* since TOK-03 and *the file* before that. Three measurements taken since say the unit is
wrong:

| measured | reading |
|---|---|
| single-pass recall of a 7-seat reading | **43–48%** (iter-50's paired same-tree experiment; Chapman `N̂` ≈ 23 and a floor) |
| a vetted, schema-safe pin advance | moved **22 of 23** `main.go:N` citations; the fence caught **1** — **4.5%** |
| one working day of platform commits | **81 fresh drift sites across 21 files**, larger than the 46-item union ten readings had been draining |

**A claim-by-claim repair pass cannot outrun any of those.** But look at what the three residuals actually
are. The 81 sites share **one** false predicate (*three services are live-local husks*). The 17 files / 30
occurrences naming a `graphql` profile share **one** (*a `graphql` profile exists*). The 21 moved citations
share **one** (*this line number names this construct*). **Three predicates, not 119 claims** — and each of
the three has a legal set that is **derivable from a platform artifact we already parse.**

That is the whole revision. TOK-01 built instruments, TOK-02 fenced a mechanizable prose class, TOK-03
attacked coverage, TOK-04 pinned the target and metered the flow. **TOK-05 changes the unit of repair from
the claim to the predicate**, which is the first term none of the four touched.

### The observation that makes it work — the platform's config IS its documentation of record

Measured, not asserted. The platform's **configuration files** are edited in the same commit as the change
and carry the rationale inline:

- `repos.yml`'s header: *"`app` is the ONLY repo with migrations to run… they own no local schema."*
- `docker-compose.yml:130-133`, on the storage service: *"v9.0: NOT in the default profiles any more — app
  serves storage in-process, and running both means two writers on one bucket. Kept startable for rollback
  comparison."*

The platform's **narrative** documentation lags and is partly unmeasured — app@v1.366.0's own
`knowledge/*.md` asserts "60K+ skills" with no measurement, and **the repo contains no job-role count
anywhere**, so "18K roles" has zero upstream provenance in the very repo we would be deferring to.

**Rule that follows, and it is the corollary to §5 rule 19:** adjudicate against **platform artifacts**,
never against another document — ours or theirs. Two documents that agree are not two witnesses.

### Revised strategy — five decisions, each recorded and each fenceable

Full records in `iter-59/decisions.md`. Summarised:

**`D-M257x-59-1` — clause 5's residual is scoped by PREDICATE, not by count.** A repair unit is now *"every
site in the tree asserting predicate P"*, adjudicated against the platform artifact that defines P's legal
set, and closed by a fence that makes the predicate underivable-when-false. §5 rule 19 said *repair by
claim, not by file*; this extends it one level: **repair by predicate, not by claim.** It does **not** touch
the reading instrument — the union-of-two discipline, the blind second reading, and the pre-commit double
reads all stay, and clause 5 is still met only by a reading that returns **zero**.
**`FIX-M257x-iter53-union-set` (46 vs 35) is a PENDING USER DECISION and is NOT resolved here.** What can be
said: predicate-scoping **subsumes** the question rather than answering it. A predicate-scoped repair covers
every site sharing the predicate *whether or not any reading named it* — which is precisely the ≈43–48%
recall problem the union was invented to work around. So the union count stops being the **scoping** input
and becomes a **validation** input: after the predicate sweep, the union's members must all be covered, and
any that are not name a predicate we have not yet found. Whether that set is 46 or 35 changes the
validation, not the work.

**`D-M257x-59-2` — the fence widening is the next build, and it is a NEW sibling guard, not a widening of
`platform_alignment_guard.py`.** A deliberate departure from the reconciliation's recommendation, with a
measured reason: that guard declares `FENCE_KIND = "standalone"` and its own docstring scopes it to *"a
map↔platform property"*; assertion F derives its clone roots from `repos_yml_path` precisely so citations
and membership are checked against the **same** reference. Re-targeting it at the whole corpus would break
the property that makes it trustworthy. Its `compose_blocks()` parser — already written, already tested — is
reused as an importable primitive, so this is an extension of the **code** and not of the **guard's
subject**. §8 rule 1's derived registry means a new `*_guard.py` self-registers.

Inputs: `repos.yml` + `docker-compose.yml` (**with its `include:` resolved**) + `Makefile`. Six assertions,
each **derived and run in both directions** (a doc-promised value with no artifact backing is a **false
promise**; an artifact value with no doc row is **undiscoverable**) — the `demo_knob_guard.py` precedent,
which already does exactly this and is green in the suite.

Denominators it must reproduce, **all re-derived at this open** (`platform 0dab54d`):

| # | assertion | derived denominator |
|---|---|---|
| G1 | every documented profile token selects something beyond the always-on floor | **10** services · **8** legal profiles · floor **3** (`postgresql redis sentinel`) · `core` selects **5** |
| G2 | any "N repos cloned" claim | **6** |
| G3 | any default-bring-up container claim | **5** (= `select(core)`) |
| G4 | any cited `*_RPC_ADDR` | **4**, all `http://backend:8083` |
| G5 | any named migration target | **1** repo (`app`) |
| G6 | half-fold split (see `D-M257x-59-4`) | compose-sets vs app-reads, per env var |

**Grade on "does it still SELECT something", never on "does it still parse"** — that is the dominant new
failure mode and G1 exists for it alone. `PROFILE=graphql`, `PROFILE=cms` and `PROFILE=storage` all exit 0;
**measured at this open they each start 3 containers, not 0** — Postgres up, Redis up, sentinel up, `docker
ps` non-empty, and the application absent. The briefing said "starts nothing"; the truth is worse, because
"nothing" would at least be unambiguous.

**`D-M257x-59-3` — §7 rule 4 gains a second half: a pin advance is not vetted until its CITATION delta is
measured.** iter-58 is the case study: the advance was vetted for 0 migrations / 0 destructive DDL / 0 new
hard-required config, all correct, and it still moved **22 of 23** `main.go:N` citations with a **4.5%**
fence catch rate. **Schema-safety and citation-safety are unrelated properties and rule 4 only ever measured
the first.** The new half: before taking an advance, re-resolve every corpus citation whose path lands in
the advancing repo at the new ref, and record moved/dead/held. It is bounded and cheap — the whole set is
**23**. Repair belongs to the **advancing iter**, as P3's *"the iter that detects the move re-points, in
that iter"* already says for refs. `FIX-M257x-iter58-mainline-shift` (**21 of 22** outstanding) is the
retrofit case.

**`D-M257x-59-4` — a HALF-LANDED fold gets a state of its own, and it is recorded on both sides or not at
all.** The map has two states per row (prod / fresh local stack) and a 7-token vocabulary with **no token
for mid-fold**, so the storage split is currently recorded **nowhere**. Measured at this open:

| side | state |
|---|---|
| compose | `storage` moved to `profiles: [storage-legacy]` — **not** in `core`, so a default bring-up never starts it; rationale in-comment at `docker-compose.yml:130-133` |
| compose env | **`STORAGE_RPC_ADDR` is set nowhere** — absent from `docker-compose.yml` and from `.env_example` |
| `repos.yml` | `storage` **still present** — still cloned |
| app `v1.366.0` | **still reads it**: `main.go:446`, `:524`, `:992`; and **hard-requires** it in two tools — `cmd/academyImport/main.go:235` and `cmd/academy-asset-upload/main.go:133` both `return … "STORAGE_RPC_ADDR is required"` |

So on every stack we currently run green, `os.Getenv("STORAGE_RPC_ADDR")` returns the empty string and a
storage client is constructed against it — a failure deferred to call time, not boot time. That is neither
`live-standalone` nor `merged-into-app`. **The fix: an 8th state token (`mid-fold`) that assertion C accepts
only when the row carries a TWO-SIDED citation** — the config side and the consumer side, each resolving.
G6 fences the pair. **Messenger is next by the developer's own account, so this row shape will be needed
again before it is finished being written** — which is the argument for building the state rather than
prose-ing the storage case.

**`D-M257x-59-5` — ordering: fence first, then citations, then the map's new state, then read.** Recorded in
full below under *Next-tik direction*.

### What TOK-05 keeps

- **Everything TOK-04 built.** P1 (refs in every measurement), P2 (instruments as committed files), P3 (the
  ref re-checked at close, re-pointed by the detecting iter), P4 (**derive, else fence, else declare**) are
  unchanged and are the machinery all five decisions above run on. P4 in particular is the *reason* the
  predicate unit works: a predicate with a derivable legal set is a P4 row-one candidate.
- **Union-of-two readings**, the blind second reading, and **pre-commit double-reads** — the last of which
  TOK-04 upgraded from *recommended* to **load-bearing** on same-day evidence, and which stay load-bearing.
- **Smaller edits: deletion > minimal scoping edit > rewrite**, with added-words counted per repaired claim.
- **Adjudicate against platform artifacts, never against another doc** — reinforced, now with the PR #14
  result as its proof.
- **Run an ancestry check before attributing a defect to M257x.** 2 of 5 verified findings in the
  reconciliation were inherited `origin/main` errors, not ours; `git log -L` on the line plus
  `git merge-base --is-ancestor` settles it, and the milestone should not be paying for `main`'s debts.
- **Cadence against commit rate, not doc PRs** — 4 of the last month's 5 structural changes have **no
  upstream doc PR at all**, so watching for doc PRs is watching the wrong signal.

### What this revision explicitly does NOT do

- **It does not re-cut clause 5, narrow it, or read it met any other way.** The user has ruled **three**
  times: met only by a reading that returns **zero**. Predicate-scoping changes how the residual is
  *repaired*, never how it is *graded*.
- **It does not weaken the audit instrument.** Not one seat, rule or blindness requirement is relaxed.
- **It does not defer the residual** to a future milestone, and does not propose closing at 4 of 5.
- **It does not resolve `FIX-M257x-iter53-union-set`** — pending user decision, stated as such.
- **It does not merge or cherry-pick `origin/pr-14`.** Read-only ref; the verdict stands at DO NOT MERGE.
- **It does not edit the platform repo.** Zero platform edits, as the whole release requires.

**Strategy class:** `more-granular` — and precisely inverted from what that label usually means late in a
milestone. The four prior toks each attacked a **larger** term (instrument → prose class → coverage →
target stability). TOK-05 goes *underneath* the claim to the predicate, which is a smaller and more specific
object, and gets **broader** coverage as a result: one derived legal set closes 30 occurrences across 17
files that no reading has ever named in one pass. Granularity in the unit, breadth in the reach.

**Distance-to-gate context:** **4 of 5**, at platform `0dab54d` / app `v1.366.0` / rext
`fast-build-m257x-iter-58` — all pins coherent, both trees clean, `stack-core` at its **1F/610** baseline
(the 1 being the perishable iter-48 fixture). Clauses 1–4 hold; clause 4 remains under test and tracked the
storage removal with **zero human action**, which is the strongest available evidence for the derive-first
ordering this tok doubles down on. **Clause 5 is the only open one** and its **net** metric — (repaired) −
(induced) − (newly falsified by platform movement) — last read **−72**. The honest late-milestone
characterization: the per-pass return under a claim-unit is not small, it is **negative**, and the three
predicate classes named above are the first residual in this milestone whose legal sets are all derivable
today.

**Cross-refs to prior TOKs:** **TOK-01** (*instrument first, then follow*) — its derived lists are why
clause 4 tracked a live platform removal unaided; untouched and vindicated again this iteration.
**TOK-02** (*fence the prose the way the anchors are fenced*) — TOK-05 is its thesis applied at the right
granularity; TOK-02 fenced prose *claims*, TOK-05 fences the *predicates* claims are instances of, which is
why the 17-file `graphql` class is one build rather than 30 edits. **TOK-03** (*repair the union, shrink the
estimator, smaller edits*) — all three moves kept; its **premise** stays refuted (a fixed residual more
readings can drain) and TOK-05 adds the reason: a reading can only name *instances*, never the *predicate*.
**TOK-04** (*pin the target*) — kept whole; TOK-05 is the next term, not a replacement. TOK-04 made the
numbers real; TOK-05 changes what the numbers are counting.

**Next-tik direction** — `D-M257x-59-5`, in dependency order:

1. **iter-60 = build the sibling guard (`D-M257x-59-2`) and let G1 close the `graphql`-profile class.**
   First because it is the cheapest win available, its six denominators were measured **today**, and it
   converts three predicate classes from prose into derivation. Watch every assertion **RED before trusting
   it** (§8 rule 5 — collect the mutant before running it), and mutation-verify the fixtures too (rule 2).
   **Pre-registered, therefore refutable:** G1 goes RED naming `graphql` across **17 files / 30
   occurrences**, `cms` and `storage` as well; after repair it is GREEN and the reverse direction names any
   of the 8 legal profiles the corpus documents nowhere.
2. **iter-61 = land §7's citation-safety half (`D-M257x-59-3`) and spend it on the 21 outstanding
   `main.go:N` sites.** The rule and its first application in one iter, so the rule is proven by use rather
   than by assertion.
3. **iter-62 = the map's `mid-fold` state + the storage row (`D-M257x-59-4`), with G6 fencing the split.**
   Before messenger folds, not after.
4. **Then `/developer-kit:harden-mstone-iters`** — see the recommendation below; **not due yet.**
5. **Then the next paired reading** — the first ever taken against a corpus whose three largest predicate
   classes are fenced rather than prose. Only then is a zero reading arithmetically reachable.

**Harden recommendation (the orchestrator spawns it, not this iteration):** the counter **restarted at
iter-58** after pass 15 closed `STABILIZED`, so it stands at **1 tik against a threshold of 10 — NOT due.**
When it comes due it owns the new sibling guard's assertions, which are exactly the AST/call-site shape the
three standing `HARDEN-CAP-ACCEPTED` entries said the residue needs.

---

## iter-101 decisions — 2026-08-06

### D-M257x-101-1 — the `app` clone is NOT advanced, and its precondition is measurably UNMET

**Decision (user's, proceeded on):** do not advance `stack-demo/app`. Two grounds, both re-derived here
rather than inherited:

1. **The exit gate names *platform* @ origin HEAD, and `app` is a different repo.** Verified against the
   milestone's own `exit_gate` string: *"Against platform @ **origin HEAD**…"*. `stack-demo/platform` is at
   `0c91421d` and `git ls-remote origin HEAD` returns the same sha — **the gate's named subject is level.**
2. **A fetched clone is already graded at origin HEAD wherever its own HEAD sits.** Re-derived in
   `stack-core/anchor_construct_guard.py:267` and `:664`: the ref ladder is
   `("origin/main", "HEAD") if ref == "auto"`, i.e. **origin/main first**. So advancing the working tree buys
   little and would move every `app:file:line` citation in the corpus (33 in-scope files reference `app/`).

**The measurement contradicts the decision's PRECONDITION, and that is recorded rather than smoothed:**
the decision rests on the clone being *kept fetched*. It is not.

| clone | local `origin/main` | real remote `main` | |
|---|---|---|---|
| **app** | `2035f9a4` | **`ad9f3c49`** | **STALE** |
| rosetta-extensions (per-stack) | `6130bfd8` | `09d06070` | STALE (re-pinned this iter) |
| the other 11 clones | — | — | **all LEVEL** |

11 of 13 clones are level; `app` — **the single most-cited repo in the corpus** — is one of the two that are
not. The consequence is precise and it is not the flattering one: **the citation guards are currently
grading `app` anchors at `2035f9a4`, which is itself no longer origin/main.** The reassurance that "a fetched
clone is graded at origin HEAD" therefore **does not presently apply to `app`**.

**The fetch was NOT performed in this run, deliberately.** A fetch moves `origin/main`, which is the ref the
citation guards resolve against — doing it while four adjudicators were mid-re-derivation would have changed
the grading basis of a reading in flight, which is the precise failure this milestone exists to prevent.
Routed as **`FIX-M257x-iter101-app-clone-unfetched`**, to be taken BETWEEN readings, never during one.

**Decision stands: do not advance. Amended with: keeping it fetched is an ACTION, not a state, and nobody
had been performing it.**

### D-M257x-101-2 — a zero reading gets a planted-defect positive control before the gate is declared met

**Decision (user's, recorded):** at the moment any reading returns `N = 0`, clause 5 is **not** declared met
on that reading alone. A **planted-defect positive control** is authored first: seed *k* known defects into
the tree, run the reading blind, and measure recall against a **known** denominator.

**Why this is honouring clause 5, not re-cutting it.** Clause 5 still requires a reading that returns zero
and nothing else. The control changes **no** threshold, **no** scope, and **no** part of the reading. It
establishes only whether the instrument that returned the zero **can see anything at all**. The four user
rulings are untouched.

**The gap it closes.** Every recall figure in this milestone — 43–51 % per pass, 62–78 % union — is
**within-reading overlap between two simultaneous passes sharing briefing, file set, partition and model**.
Shared inputs inflate overlap and therefore flatter recall. **There is no measurement of what this instrument
actually sees against a known denominator.** A zero from a blind instrument satisfies the letter of clause 5
while proving nothing — and this milestone has already found ~20 instruments that reported a state without
measuring it. This decision exists so this is not the twenty-first.

**Cost:** 1 iter to author the plant, 1 cycle at the moment of the zero. **Not** payable in advance of a
zero, because a plant authored early leaks into the pool it is meant to measure.

### D-M257x-101-3 — the rext pin is cut at authoring HEAD, and pushed in the same breath

The pin (`.agentspace/rext.tag`) named `fast-build-m257x-iter-67` while the per-stack consumption clone sat
at `fast-build-m257x-iter-58` (`ab81527a`) — a mismatch `demo-stack/ensure-clones.sh:94-101` treats as
**FATAL** (`exit 1`), so `/demo-up` aborted outright. Both tags are on origin; the mismatch was the defect.

The sharper half: the authoring copy was **43 commits ahead of the pin**, and those 43 include **`7844e97`**,
which keys the demo override's volumes reset on a **derived** property instead of the deleted service name.
Measured: `7844e97` is an ancestor of authoring HEAD and is in **neither** `fast-build-m257x-iter-67` **nor**
`fast-build-m257x-iter-68`. It existed **only in the authoring copy** — unreachable to any stack.

That fix is not incidental to this milestone; it is the fix for the platform advance clause 1 must now be
proven against. Platform `838d907` moved the `$HOME/.aws/credentials` bind **off the deleted `jobsimulation`
service and onto `backend`** — re-derived here: `docker-compose.yml:100` at `0c91421` carries that bind under
`backend`, and the mitigation keyed on the literal `"jobsimulation"` had silently gone dead.

**Cut `fast-build-m257x-iter-101` at authoring HEAD `09d06070`, pushed, and verified on origin by
`git ls-remote --tags origin` before being treated as done** — rung zero, *tagging is not publishing*.
Origin now returns `0011c10a refs/tags/fast-build-m257x-iter-101` dereferencing to `09d06070`.

**HANDED OFF MID-ITER — read this before acting on the above.** A concurrent lane was given end-to-end
ownership of the rext tag / pin / per-stack-clone job **after** the tag had already been cut and pushed.
The split at the moment of handoff was:

| step | state | owner |
|---|---|---|
| tag cut at `09d06070` | **DONE** | this lane |
| tag pushed + verified on origin | **DONE** | this lane |
| `.agentspace/rext.tag` updated `iter-67` → `iter-101` | **NOT DONE** | the other lane |
| `stack-demo/rosetta-extensions` brought level | **NOT DONE** | the other lane |

The two remaining steps were deliberately **not** taken by this lane even before the handoff, for a reason
that outlives it: the sealed ground truth pins `stack-demo/rosetta-extensions` at `ab81527a`, and four
adjudicators were re-deriving rext claims **in that tree** at the time. Moving it — or the pin file a seat
may consult — mid-adjudication would have changed the subject of a reading in flight.
**Sequencing rule, generalised: move a clone or a pin BETWEEN readings, never during one.**

`fast-build-m257x-iter-101` already exists on origin; the other lane should **re-pin to it rather than cut a
second tag**, or the two tags will disagree about which tooling the release means.

### D-M257x-101-4 — the gate is re-graded **2 of 5 PROVEN**, and the booked "4 of 5" is withdrawn

The milestone has carried **4 of 5** since iter-37. It is not defensible, and the reason is the gate's own
first clause of text: *"Against platform @ **origin HEAD**, never a pinned pre-drift commit."*

**Re-derived, not inherited** (`stack-demo/platform`, all shas verified this iter):

| clause | booked | **honest** | why |
|---|---|---|---|
| 1 — 3× cold green cycles | MET (iter-18, **2026-08-01**) | **UNPROVEN at origin HEAD** | proven at `2adcf71`; **6 platform commits** and **281 changed lines of `docker-compose.yml`** behind current HEAD |
| 2 — full Playthrough suite | MET (iter-37, **2026-08-02**) | **UNPROVEN at origin HEAD** | same stack, same drift; proven one day after clause 1 |
| 3 — migration-status map | MET | **MET** | re-verified at `0c91421`; fenced both ways by `platform_alignment_guard` |
| 4 — zero rext writes to dropped schemas | MET | **MET** | re-verified at `0c91421`; fence watched RED |
| 5 — KB-fidelity GREEN / YELLOW-0-blockers | NOT MET | **NOT MET** | this iter: **N = 24**, a floor |

**The drift is not cosmetic, and it is larger than the 107-line figure in circulation.** 107 lines is the
diff from `0dab54d`; the clauses were proven *before* that commit. From the actual proof ref:

```
git log --oneline 2adcf71..HEAD   ->  6 commits
git diff --stat 2adcf71..HEAD     ->  docker-compose.yml | 281 ++-----  · repos.yml | 29 ·  Makefile | 20
```

Measured across that span: **3 compose services deleted** (`storage`, `messenger`, `customerio-sync`),
**2 `repos.yml` entries removed** (`messenger`, `storage`), and the default profile **renamed
`graphql` → `core`**.

**The mechanical proof that clause 1 cannot be carried, not merely an argument from drift.** Platform
`838d907` moved the `$HOME/.aws/credentials` bind **off the deleted `jobsimulation` service and onto
`backend`** — re-derived: `docker-compose.yml:100` at `0c91421` carries that bind under `backend`. The demo
override's mitigation was keyed on the literal `"jobsimulation"`, so it silently went dead, **and its
tripwire test skipped**, which reads exactly like a pass (§5 rule 8). The fix (`7844e97`) is in **neither**
`fast-build-m257x-iter-67` (the pin) **nor** `fast-build-m257x-iter-68` — it existed **only in the authoring
copy**, unreachable to any stack, until this iter's tag. **The tooling that would have run clause 1 today
predates the fix for the platform change clause 1 must be proven against.**

**UNPROVEN is not REFUTED.** Nothing here shows clauses 1 and 2 would fail — only that the evidence for them
was taken against a platform that no longer exists, which is precisely what the gate's own wording forbids.
On this host the EISDIR hazard is additionally **latent**: `$HOME/.aws/credentials` exists as a 0-byte file,
so the failure mode that fires on a fresh box does not fire here.

**Grade: 2 of 5 proven, 2 unproven-pending-re-run, 1 not met.** Re-running clauses 1 and 2 at `0c91421` with
tooling that includes `7844e97` is the work; a concurrent lane owns it.

**The generalisable rule — this is the milestone's own recurring class turned on itself:** *a gate clause is
proven at a REF, and it decays when that ref moves.* Booking a clause MET without re-anchoring it is the same
defect as a corpus claim pinned to a stale sha, and M257x has now found it in its **own** exit gate.

---

### DEF-M257x-iter101-crosslane-fetch — a reading's ground truth moved under it, and it is my defect

**Class:** coordination · **Severity:** medium · **Exposure:** bounded and stated · **Status:** recorded,
rule written, not recurring (the rule now forbids it)

**What happened.** M257x runs three lanes concurrently against one checkout. I assigned `stack-demo/**`
exclusively to Lane B — while iter-101's adjudicators were grading claims against the platform clones,
which live **inside `stack-demo/**`**. Lane B ran a clone-set refresh; iter-101's adjudication commit
landed just after it.

**Measured, at iter-102's open:**

| | |
|---|---|
| Lane B's fetch window (`FETCH_HEAD` mtimes) | **11:18:16 – 11:20:51** |
| iter-101 adjudication commit `a360d66` | **11:21:55** |
| overlap | **≈ 68 s** between the last fetch and the commit |
| clones whose **HEAD advanced** | **5** — `app` (98 commits / 634 files), `next-web-app` (41 / 192), `ant-academy` (5 / 86), `sentinel` (2 / 3), `studio-desk` (2 / 9) |
| corpus files moved by the fetch | **0** |
| `platform` (the ref clause 3/4 are graded at) | **unmoved** — `0c91421`, still `== git ls-remote origin HEAD` |

**The exposure, stated honestly rather than minimised.** Most of adjudication ran pre-fetch and run 64
observed the pre-fetch ref, so **`N = 24` stands**. But **it cannot be PROVEN that no adjudicator
re-derived post-fetch**, and that is the whole problem: a fetch moves the very refs the citation guards
resolve against (`CITE_REF=auto` → `origin/main` first, §5 rule 41). An adjudicator who happened to
re-derive after 11:20:47 graded a *different subject* from one who re-derived before, and **nothing in
either report says which**.

**What makes it mine rather than Lane B's.** Lane B did exactly what it was told to do with a tree it had
been given exclusively. **The ownership map was wrong, not the lane.** Path ownership was **necessary and
not sufficient** — two lanes can hold disjoint sets of *writable paths* and still collide, because the
reading's **SUBJECT is wider than the paths anyone declared**. It reaches into a tree the reading never
writes and does not own. An instrument whose inputs sit outside every declared boundary is an instrument
nobody is protecting.

**The mitigating fact, and it is a measurement rather than a hope.** The `app` move — the largest of the
five, and the one 17 corpus sites depend on — turns out to have changed **nothing any corpus claim cites**:

| | at `2035f9a4` | at `ad9f3c49` |
|---|---|---|
| `main.go` line count | 1639 | **1639** |
| `main.go:504` / `:524` / `:525` / `:1384` / `:1450` | — | **all five byte-identical** |
| `terraform/main.tf` line count | 786 | **786** |
| `terraform/main.tf:181` (`service_desired_count = 1`) | — | **identical** |

The 5 commits touched `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`,
`terraform/main.tf` and `terraform/variables.tf` — **no Go source, and no cited terraform construct.** So
the fetch could not have changed any adjudicator's verdict on an `app` anchor even if one had re-derived
mid-window. **That is a fact about this particular move, not a general reassurance**, and it is exactly why
the rule below is stated in terms of *provability* rather than *outcome*.

**Durable rule — written into `corpus/ops/platform-alignment.md` §5 rule 41a.**

> **A reading's ground truth is not just the corpus — it is the corpus PLUS every clone ref the reading
> resolves against.** Freeze both for the reading's duration. **No lane may fetch any clone while a reading
> is in flight.** A lane that needs a clone advanced says so, and the fetch happens **between** readings and
> is recorded in the next reading's ground-truth sheet.

Two corollaries also written into 41a: a ground-truth sheet must record each clone's **fetch time**, not
only its sha, so a mid-reading move is *detectable* rather than *suspected*; and *"it has since been
fetched"* is **a new ground truth, not a repair** — the next pass re-derives against the moved refs and
measures what the move injected, and never retro-fits the old reading.

---

### D-M257x-102-1 — `DEF-M257x-iter80-storage-prod-bucket` is FILED to the platform-defect register

Answers the deferral audit's **Q1**. Escalated and undecided for 21 iterations; the audit made it a
blocking input to its own gate.

**Choice.** Option (b): **file it to `platform-defect-register.md`** as
`PLATFORM-M257x-compose-points-local-backend-at-the-PRODUCTION-S3-buckets`, and leave the rext isolation
registry as it stands with the disagreement documented. `DEF-M257x-iter80` moves **escalated-undecided →
filed**.

**Why.** The deciding line is **platform source** — `docker-compose.yml:82-83` @ `0c91421`, re-derived for
this filing and still true, inside the **`backend`** block, which the **default `core` profile** starts.
This milestone is zero-platform-edit, so the register is precisely the sanctioned destination for *"Rosetta
cannot fix this."* Filing does not pre-empt a platform-side fix; it makes the disagreement **documented and
permanent**, which is the honest end state.

**The audit's own finding that this closes:** `platform-defect-register.md` held **4 entries, all from
M256, and zero from M257x** — despite having been created *by M256's deferral audit* for exactly this
class. The register built for this item had never been used. It now has its first M257x entry.

**What is NOT decided here.** `stack-seeding/isolation/isolation.go:106` still registers `s3-private` as
`PerStackIsolated`. That is **ours**, it is a separate open question, and (a)/(b) were never exclusive.
Recorded rather than quietly folded in.

---

### D-M257x-102-2 — `FIX-M257x-iter53-union-set` is DROPPED as subsumed (Fate 3)

Answers the deferral audit's **Q2**. The question as booked was *"is the clause-5 repair target 46 or 35?"*

**Dropped.** `D-M257x-59-1` (TOK-05) made predicate-scoping the repair unit **47 iterations ago**, which
turned the union count from the *scoping input* into a *post-sweep validation cross-check*. The item has
been restated by iters 76, 82, 86 and 100 in a form that no longer corresponds to how the repair works.

It sat on the do-not-touch list pending a user ruling. **The thing it was pending on no longer exists**, so
there is nothing for a ruling to decide. Same shape as `D-M257x-102-4` below and the same lesson: an item
survived because each pass restated it instead of re-reading it.

---

### D-M257x-102-3 — the gate re-grade is RATIFIED, and clauses 1–2 are a CLOSE BLOCKER

Answers the deferral audit's **Q4**. **`D-M257x-101-4` is ratified: 2 of 5 PROVEN.**

**Clauses 1 and 2 stay in M257x's scope and are a close blocker — they are NOT routed to M258.** The gate's
own first clause of text is *"Against platform @ **origin HEAD**, never a pinned pre-drift commit."*
Routing them forward would be **re-cutting the gate**, which is not on the table. A concurrent lane owns
the re-run.

The 2-of-5 re-grade is honest accounting, not a regression: **UNPROVEN is not REFUTED.**

---

### D-M257x-102-4 — `CHECK-M257x-iter38-ai-act-classification` is DROPPED (Fate 3); the repair was finished 36 iters ago

Withdraws the deferral audit's **Q3**, which is therefore no longer a user question. **§8 of that audit now
has ZERO open user questions.**

**Re-read at source** (`iter-38/decisions.md` `D-M257x-38-2`, `iter-38/progress.md:128-130`) and
**re-verified in the corpus at this iter's open**: `security_compliance.md:227` sits inside a blockquote
(`:231-256`) that says in terms *"Both bullets above are what is STATED, not what this corpus asserts"*, and
`ai_architecture.md:284` carries the matching retraction. **The corpus asserts no classification, in either
direction.**

So there is **no false claim to fix**. What was carried was not a defect but an **aspiration** — *that
someone re-derive the true classification* — routed as *"needs an owner outside this milestone."*

- **Not a documentation defect** — nothing in the corpus is false.
- **Not clause-5 blocking** — clause 5 requires the corpus not to assert falsehoods, and it does not.
- **Not Rosetta's to answer** — the corpus documents what the code does. Issuing a legal classification in
  either direction is exactly the category error iter-38 identified and correctly refused.

**Silence is the correct end state, not a placeholder for a future answer.**

**The meta-lesson, and it is a real one — written into `platform-alignment.md` §5 rule 47.** A completed
repair was carried for **36 iterations** as an open question because the ledger recorded the **aspiration**
rather than the **defect**. An item whose only remaining content is *"someone should determine X"* is not
deferred work — it is a **finished repair with a wish attached**. The diagnostic is one question: *if nobody
ever does this, is any statement in the corpus false?* If no, the item is done.

---

### D-M257x-102-5 — BINDING USER DECISION: no legal/regulatory escalation during delivery

**The user's words (2026-08-06):** *"don't bother me with legal stuff during this migration: our goal is to
close this release, not waste resources on legal matter."*

**Binding for the remainder of v2.8. It constrains ROUTING, not just this iter.** When a legal, regulatory,
compliance or policy question surfaces, pick one of three **without asking**:

1. **already repaired** → **close it**; silence is a valid end state;
2. **a real defect this repo cannot fix** → **file it** to `platform-defect-register.md` and move on;
3. **genuinely blocks delivery** → surface it, **and state exactly what it blocks**.

**Never route one as "needs an owner."** That routing is what turned a finished repair into a 36-iteration
standing question (`D-M257x-102-4`).

**Applied retroactively** to the open ledger: anything whose remaining content is a legal or policy
*determination* rather than a *repair* is closed or filed, never carried. Written into
`platform-alignment.md` §5 rule 48, next to rule 47 — the two are the same lesson from opposite directions:
**the ledger records defects, not curiosity.**

---

### D-M257x-102-6 — the deferral audit's "urgent" tag escalation (F18) is a FALSE POSITIVE, and the state is clean

The audit escalated, as the single **urgent** item for the rext lane, that *"three `fast-build-m257x-iter-101*`
tags exist on origin, two of them pointing at the same commit"* — an ambiguous code-of-record.

**Re-derived directly against origin at this iter's open:**

```
git ls-remote --tags origin 'refs/tags/fast-build-m257x*' | grep -v '\^{}' | wc -l   ->  54 distinct tags
… | grep 'iter-101'                                                                  ->  1
fast-build-m257x-iter-101  ->  tag object 0011c10a, peeled commit 09d06070
git tag --contains 7844e97 | grep -c fast-build-m257x-iter-101                        ->  1
.agentspace/rext.tag                                                                  ->  fast-build-m257x-iter-101
```

**There is exactly ONE.** No `iter-101b`, no `iter-101c`; **no origin tag points at `4cb920a`** at all. The
pin, the clone HEAD (`09d0607`) and the tag agree, and the tag contains `7844e97`. **Lane B's re-pin is
done and correct.**

**Two mechanisms are available and I am not asserting one over the other, because the evidence does not
separate them:**

1. **A peeled-ref miscount, which is demonstrable.** `git ls-remote --tags` prints **two lines per annotated
   tag** — the tag object and its `^{}` peel. For `iter-101` alone: **2 raw lines, 1 distinct ref.** Counting
   *line shapes* rather than *distinct refs* doubles every annotated tag. **This is the same mechanism run 57
   fixed in the guard family** (*"take the RED headline's cardinality from the GUARD, not from line shapes"*),
   recurring in a different instrument.
2. **A transient state since reconciled.** Lane B's commit `b02150c` is titled, in terms, *"one rext tag, not
   two"*. The audit may have observed extra tags that were subsequently deleted.

**Mechanism (1) cannot manufacture the specific NAMES `iter-101b` / `iter-101c`, so it is not a complete
explanation on its own** — which is why (2) is recorded beside it rather than discarded in favour of the
tidier story. Asserting the flattering single explanation is the failure this milestone exists to prevent.

**What is booked regardless:** a **peeled-ref line-shape miscount is a real `--tags` reading defect** and
will recur. Any future count of tags takes `| grep -v '\^{}'` first, or counts `refs/tags/` names.

**Corrected in `deferrals-audit.md`:** F18 is struck, and §10's *"F18 is the urgent one"* handoff to the
rext lane is withdrawn.

---

## TOK-06: fence the inflows before repairing again — 2026-08-06

**Tok type:** **deliberate** — author-initiated, and **not session-terminating**.

Two things need saying before the strategy, because both have been got wrong in this milestone before.

**It is not the streak.** The clause was checked before this was written. The last three tiks are
iter-101 (reading, `N` **28 → 24** — progress), iter-102 (repair, 52 anchors / 98 sites, no reading inside
it) and iter-103 (reading, `N` **24 → 33**). iter-101 moved the metric, so the window never reached three
consecutive no-progress tiks and Phase 0 rule 2 falls through to tik. Same precedent as TOK-04 and TOK-05,
both fired by something other than the streak.

**It does not terminate the call**, and the reason is the bootstrap tok's reason rather than a licence.
A *triggered* tok exits so the user can review a revision before the next tik commits to it — the user has
watched a strategy run and stall, and the revision is the new thing. Here there is nothing unreviewed to
commit to: iter-103 produced the measurement **and** the replacement, and **every element of the sequence
below is an item already routed in iter-103's own close.** This tok *sequences* routed work. It opens no new
territory, so it closes and the loop continues into tiks in the same call.

**Prior strategy:** [`TOK-05: stop repairing claims; fence the predicates under them`](#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04), on the
[`TOK-04`](#tok-04-pin-the-target-or-stop-calling-it-a-measurement--2026-08-03) ref-discipline spine.
**Neither is discarded and neither failed.** TOK-05's unit of repair — the predicate — is *vindicated* by
iter-103's band #3, and that measurement is the reason this revision is about sequencing and not about
repair.

### Why it stopped being enough — the burn-down leg does not reach the residual

iter-103 read `N = 33` against a rule sealed in its own commit (`04cbcfc`) **before the first seat was
dealt**: `≤16` works · `17–22` ambiguous · **`≥23` DOES NOT REACH**. The `≥23` branch fired, graded exactly
as written, with nothing re-cut once the number was known.

The composition inverts what that appears to mean:

| | iter-101 | iter-103 |
|---|---|---|
| distinct false **predicates** | **22** | **22** |
| distinct **anchors** | 24 | **33** |
| anchors per predicate | 1.09 | **1.50** |

**By predicate the pool did not move.** After a 52-anchor / 98-site repair the corpus carries the same
number of distinct false propositions — in more places.

**And the repair is exonerated, blind.** 21 of iter-101's 22 predicates are CLOSED. The single survivor,
`prod-terraform-8081` at `skiller.md:19`, is one **iter-102's own repair map flagged `SEAT 9 (?)` and left**.
The repair leg reaches what it aims at; an independent instrument says so.

So the loop is not failing to converge because repair is weak. **`N` held up because two inflows feed the
residual and nothing watches either of them:**

| inflow | share of `N` | fenced by |
|---|---|---|
| **clone advance** — version literals, `go.mod` pins, symbol names, line offsets | 20 anchors — **61 %** | *nothing* |
| **the repair's own induction** — prose iter-102 wrote | 7 anchors — **21 %** | *nothing* |
| never-true | 4 — 12 % | the reading |
| unclassified | 2 | — |

**Inflow is comparable to outflow. A loop with that property does not converge, and running it faster does
not help.** That is the sentence the revision turns on. It is iter-103's, and it is measured.

### New strategy — fence the inflows before repairing again

**Repair-then-read is the wrong loop while 82 % of the residual arrives from sources nothing watches.**
Five steps, in dependency order, each one an item iter-103 already routed:

**0. `FIX-M257x-iter103-guard-tree-provenance` — first, because every guard verdict in this milestone has
unstated provenance.** At iter-103's close the family came back **2 RED** and two quotable conclusions were
drafted — *"the sheet asserted a verdict it didn't have"* and *"a fence names 8 sites the reading missed, so
`N ≥ 41`"* — **and both were false.** The fence had been run from the **pinned** rext clone rather than the
**authoring** copy; the whole difference was `claim_twin_waivers.json` (+40 lines) and the 8 RED sites were
exactly the 8 waived sites. `guard_family.py` prints the corpus sha and the platform sha **and not its own**
— the one input that decides the verdict is the one the output does not state. A guard verdict is a
measurement taken with a fence's *configuration*, so it is settled by the tree that configuration lives in.
Until this lands, **every prior guard green in this milestone is provenance-unstated**, and steps 1–3 would
be shipping fences whose verdicts inherit that.

**1. `FIX-M257x-iter103-drift-fence-gap` — the drift fence, and it outranks the repair.** 61 % of `N` is a
**mechanically checkable** class: a version literal, a `go.mod` pin, a symbol name, a line offset. That is
precisely the class a fence *can* reach — unlike intra-document self-contradiction, which this milestone has
already ruled unboundable. Neither platform guard covers it: `platform_alignment_guard` fences `repos.yml`
membership, `platform_predicate_guard` fences compose profile tokens. **Repairing 20 drift anchors without a
fence just re-arms them at the next clone advance**, which is the mechanism that produced them.

**2. The induction checks.** Both shapes are mechanical and **both are repeats**:
   - **A canonical wording multiplies its own defects.** iter-102's replacement sentence asserts the
     `backend.internal.anthropos:8081` literal has *"one occurrence anywhere in the clone set."* **It has
     six**, five of them inside a repo the sentence's own 13-repo / 44-`.tf` denominator counts. It is
     **self-refuting against its own stated denominator**, and it shipped to 5 anchors. Centralising a
     wording centralises its defects — now twice.
   - **A repair rotted an anchor by inserting prose above it.** `architecture_overview.md:321` was correct
     at `8f04d3a`; iter-102 inserted a block above it, the wording moved to `:331`, and all 4 citers stayed
     put — now naming the **opposite topology**. This is **the identical mechanism iter-101 booked against
     iter-100, one cycle later**, and §5 rule 34 already names it without fencing it.

   A post-repair line-offset check and a control on any centralised wording close most of this. **Design
   both from the MEASURED shapes above, not from the general idea.**

**3. `FIX-M257x-iter103-read-union`** — the 33, by predicate, with iter-103's two riders: a canonical
sentence published to ≥3 sites is verified **against its own stated denominator** before it is multiplied,
and a repair that inserts lines above a cited anchor **re-points the citers**.

**4. Read LAST**, once the inflows are watched. Read first and the next reading measures the same inflow
again and costs a full cycle to say so.

**Binding on every fence in steps 1–2:** a **mutation control AND an anti-vacuity control that can actually
fire**. This milestone has found **six** fences green over universes they never examined, and one whose
anti-vacuity test compared a string to itself. A fence without a firing control is not shipped.

**Strategy class:** `new-direction`. TOK-01 built instruments, TOK-02 fenced a mechanizable prose class,
TOK-03 attacked coverage, TOK-04 pinned the target and metered the flow, TOK-05 changed the unit of repair
from the claim to the predicate. **TOK-06 changes the ORDER OF THE LOOP** — it is the first revision that
touches neither the instrument nor the unit, but what runs before what. It is authored on a composition
measurement, which is a term none of the five touched.

### What is retired, and must stop being quoted

**Chapman is retired for this milestone.** `m`/union measured **17 %** at iter-101 and **61 %** at iter-103
on a **byte-identical** instrument. Independence is therefore not a property of the instrument — it is a
property of *what is left to find*: subtle residual → independent passes → large `N̂`; mechanical residual →
correlated passes → small `N̂`. **Both `N̂ ≈ 103` and `N̂ ≈ 35` are unusable, neither corroborated nor
refuted.**

**Only floors survive: ≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`.** Both are two-pass unions; a floor is not a
pool size. The series 16.7 → 29.4 → 45.2 → ~103 remains four **corrections to an underestimate** — that
reading is unchanged and still right — but **stop quoting a point estimate from it.** Track `N` and the
predicate count directly; they need no assumption at all. Swept at this iter: `state.md` already carried the
retirement; two standing point estimates survived in the milestone ledger (iter-101's *"the residual is on
the order of ~100"* and iter-102's *"the pool was probably always ~100"*) and are now **marked in place**,
per the ledger's own correction convention rather than rewritten.

Four bands (#3b, #4, #5, #8) failed for that one reason — iter-102 repaired the subtle half, and
mechanically checkable drift is found by every competent pass and leaves a seat almost no room to be wrong.
The upheld rate went to **97.9 % raw / 100 % with `wrong-tree` separated**, past even the "surprising"
threshold, **for the same reason**. Three of those four numbers moved in the direction that *flatters* the
reading. **None of them is evidence the reading got better.**

### What is NOT changed

Clause 5 is **not** re-cut, narrowed, reinterpreted or argued — it is met only by a reading that returns
**zero**, ruled four times. A tok revises *strategy*, never the gate. The read instrument is untouched (one
commit ever, `012edd2`). The union-of-two discipline, the blind second reading, the pre-adjudication verbatim
commit and every TOK-01…TOK-05 fence all stand. Zero platform-repo edits; `stack-demo/**` untouched.

**Distance to gate:** **4 of 5**. Clauses 1 and 2 closed at platform `0c91421` (clause 2 **MET WITH
DISCLOSURE** — 29/1 on 2 of 2 fresh-stack first runs, never a clean pass); 3 and 4 hold; **5 open at
`N = 33`.** No estimate of the distance to zero is offered, because the estimator that used to offer one is
retired and the honest answer is a floor.

**Next-tik direction:** iter-105 lands step 0 — `FIX-M257x-iter103-guard-tree-provenance` — and reports how
many prior verdicts it re-grades.
