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
