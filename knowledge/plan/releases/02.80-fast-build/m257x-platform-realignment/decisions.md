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
