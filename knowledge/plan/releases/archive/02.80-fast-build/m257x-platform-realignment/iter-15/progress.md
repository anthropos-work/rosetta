---
milestone: M257x
iter: 15
---

# iter-15 — progress

**Type:** tik

## Clause 2, measured for the first time in this milestone: **20 live / 10 failing / 1 unimplemented**

`run-playthroughs.sh 1 --reset` on the clause-1 stack `demo-1` (the stack iter-14 left up on purpose),
consuming rext at `fast-build-m257x-iter-13`. Serial, unscoped — deliberately, because the ptreport gate is
**binding only on a full run** (`run-playthroughs.sh:300-307`) and its own documented residual at `:292` says a
broad `--grep '@pt'` is still graded ADVISORY.

```
Playthroughs coverage: 20/31 passing (64.5%)   passing=20 failing=10 unimplemented=1 unimplementable=0
ptreport: GATE no-regressions FAILED
```

**Clause 2 is NOT MET.** The denominator reconciles: 30 manifest `playthrough:` ids == 30 spec `@pt:` tags,
`diff` identical; ptreport's 31 rows are those 30 plus the one declared-unimplemented use case
(`onboarding.enterprise-workforce-standard.UC1`, an in-manifest `will-not-build`).

> The first cut of that count returned **0** — I grepped `^\s+- playthrough:` with a leading-dash anchor the
> field does not carry. That is §5 rule 3, and it is the *same number* the protocol already records a false
> absence for. The rule is written down and it still cost a minute; the cheap defence is opening the file.

## The root cause: the per-stack Directus was never actually serving

Every skill-path failure said **"Content not found."** while `directus.skill_paths` held the row, published,
under the slug the spec asks for. The backend log named it:

```
Directus FORBIDDEN 403: You don't have permission to access collection "library_categories"
```

and an anonymous `GET /items/skill_paths` 403'd too — with the grant row **present in
`directus.directus_permissions`**. Restarting the container flipped it to `200`, which said *stale
permission cache*, not *missing grant*.

Reading the bring-up transcript rather than reasoning further found the actual line, in **all three of
clause 1's cold cycles**:

```
==> demo-1: set-dressed (… snapshot:taxonomy=replayed directus=skipped(error) sim-embeddings=replayed …)
stacksnap: replay failed: replay: advance identity sequences on directus.sequences:
  check for unowned sequence defaults on directus.sequences:
  pg: query row string: ERROR: column "sequence_catalog" of relation "sequences" does not exist (SQLSTATE 42703)
```

`boot_directus_step` runs **only on a successful directus replay** (`dev-setdress.sh:373-376`), so a failed
replay silently costs the restart as well as the rows. That is the whole mechanism: replay dies → content is
not committed → Directus is never restarted → every cms-domain read 403s.

**And autoverify reported `green:true / 0 warnings` over it, three times.** Its Directus probe counts
`directus.directus_collections` rows in Postgres — which the *structure* provision creates and which
therefore survive a failed *content* replay. It never asks the running Directus for an item.
`platform-alignment.md` §5 rule 11, in the verifier that measures the gate, for the second time this
milestone.

This closes `FIX-M257-stacksnap-directus-sequences`, carried since M257 iter-02.

### Half 1 — a plan-dependent query that RAISES, not one that answers wrong

`unownedSeqColumnsSQL` / `identitySeqColumnsSQL` both read, flat:

```sql
FROM pg_attribute a JOIN pg_class c … JOIN pg_namespace n …
WHERE n.nspname = $1 AND c.relname = $2
  AND pg_get_serial_sequence(quote_ident($1)||'.'||quote_ident($2), a.attname) IS [NOT] NULL
```

That last term references only `a.attname` and the parameters, so it is a **restriction clause on
`pg_attribute`** and the planner may push it below the joins — evaluating it against column names of *other
relations*. `pg_get_serial_sequence` does not return NULL for a foreign column; it **raises**.

**It is plan-dependent, which is exactly why it hid.** Measured on demo-1, same text:

| form | result |
|---|---|
| literal-substituted in `psql` (what a debugger types) | `''` — **succeeds** |
| `PREPARE` + `EXECUTE` ×6 (what pgx does) | raises `column "sequence_catalog" of relation "sequences" does not exist` |

`sequence_catalog` is a column of `information_schema.sequences` — a relation neither query mentions.
Postgres switches to a generic plan on the sixth execution of a prepared statement; the replay walks many
tables, so it always got there.

The fix is a **barrier, not a correction** (`platform-alignment.md` §8 rule 4). `seqTargetCTE` resolves
`$1.$2` to one OID and enumerates that relation's columns in two `AS MATERIALIZED` CTEs — a hard
optimization fence in PG12+ — and then hands `pg_get_serial_sequence` **`tgt.reloid::regclass::text`**, the
very OID the column list came from. The two arguments can no longer name different relations under any plan.
Re-measured: 7 consecutive `EXECUTE`s clean, and **not vacuous** — `directus.directus_permissions` still
yields `id`.

### Half 2 — the structure script emits two thirds of a `serial`

With half 1 in, the replay advanced one table and hit the M256 unowned-sequence refusal on the next:

```
column(s) id of directus.sequences_roles default to nextval() from a sequence that is NOT OWNED BY them …
```

The refusal was right. The captured structure script emits `CREATE SEQUENCE` (`structureSeqSQL`) and
`DEFAULT nextval(...)` (`structureDDLSQL`) and **never** `ALTER SEQUENCE … OWNED BY …`. Everything works —
inserts allocate ids, Directus serves — until something asks `pg_get_serial_sequence()`, which resolves
*only* through that edge.

Measured on demo-1: **8 of the provisioned user collections unowned** (`sequences_roles`, `sequences_files`,
`sequences_files_2`, `sim_translations`, `simulations_library_categories`, `simulations_translations`,
`skill_paths_curators`, `skill_paths_library_categories`) — while **every** `directus_*` table Directus
bootstrapped itself was properly owned. That split names the provision script, not Directus.

Fixed on the **APPLY** side, not the capture side: `reconcileSequenceOwnership` derives the missing edges
from the target's own catalog immediately after the structure lands. A capture-side fix would need prod
access (`HOST-M257x-toolchain`) and would leave every existing cache broken until a re-capture; this
self-heals for all of them, old and new — §2's *resolve it from the environment at the point of use*.

**The replay's refusal is deliberately left alone.** Healing inside the guard would make it a probe that
satisfies itself (§5 rule 7). It stays the fence: if the reconciliation stops running, the replay fails loud
again naming the exact column.

The interface change was the design: `provisionConn` gained `QueryRowString`, so the compiler forced every
fake to decide rather than inherit a silent default.

### The live proof

| | before | after |
|---|---|---|
| `stacksnap replay --surface directus` | **rc=1**, 0 rows committed | **rc=0** — 14 tables, **11 986 rows**, 3 sequences advanced |
| anon `GET /items/skill_paths` | 403 | **200** |
| anon `GET /items/library_categories` | `FORBIDDEN` | **200** |

## What the second run measured, and what it does NOT mean

A re-run after the fix reported **17/31**. That is **not** a regression and must not be read as one: it ran
**without `--reset`**, after a mutating run had already completed onboarding for the pt-world heroes. The
three net-new failures are all `onboarding.*`, whose negative controls assert *"onboarding is INCOMPLETE"* —
precisely the stale-world state `run-playthroughs.sh:9-12` forbids as a reset. Comparing the two runs' totals
is comparing two different worlds. **The comparable measurement is the one taken at the layer that changed**
(replay exit code, row count, HTTP status), and it moved.

**So the honest reading: the fix is proven at its own layer and the clause-2 metric has not moved.** The two
skill-path Playthroughs still fail — on a *different* 403 (`directus_versions`), which is a fourth cause, not
the one just fixed.

## The remaining failure classes, named rather than lumped

Triage of the run-1 failing set, by error, not by test name:

| class | evidence | routed to |
|---|---|---|
| **`directus_versions` 403** — the cms domain's version lookup is refused although `/versions` answers 200 anonymously and the read+create grants exist on the public policy | 19 backend errors post-fix; blocks `getSkillPath`, `publicSkillPaths`, `getOrCreateSkillPathSession` | `FIX-M257x-iter15-directus-versions-403` |
| **content-model shape drift** — `cannot unmarshal string into Go struct field JobSimulation.data.library_category of type struct{ID uuid.UUID…}`: the app expects an EXPANDED relation, Directus returns the raw id, i.e. the replayed `directus_relations`/alias-field set no longer matches what post-cms-in-app `app` reads. **106 occurrences** — the largest single class | backend log, 8-minute window | `FIX-M257x-iter15-library-category-expansion` |
| **seeded hero absent from manager reads** — 5 Playthroughs assert "the org's seeded hero (Pat Ellis) is among the results" and get **0** (`workforce-funnel`, `workforce-succession`, `workforce-org-feedback`, `activity-drilldown`, `assign-and-track.UC2`) | run 1 | `CHECK-M257x-iter15-manager-reads-empty` |
| **hiring board type-purity** — expects 5 `hiring` rows, renders 1 | run 1 | folded into the above until measured |
| **org-role create never navigates** — `waitForURL` 60 s timeout on the role-detail route | run 1 | `CHECK-M257x-iter15-orgadmin-role-create` |

None of these is assumed to share a cause with the others. Four distinct error strings, four handlers.

## Side discovery (NOT landed — routed)

`run-playthroughs.sh:161` reads

```bash
code="$(curl -sk -o /dev/null -w '%{http_code}' "$FAPI_BASE/v1/client" --max-time 5 2>/dev/null || echo 000)"
[ "$code" != "000" ] && { echo "  ✓ fake-FAPI ready (HTTP $code)"; break; }
```

**curl writes its own `000` on a failed transfer *and* exits non-zero**, so `|| echo 000` appends a second
one: the captured value is `000000`, which `!= "000"` — and the readiness gate announces **`✓ fake-FAPI ready
(HTTP 000000)`** on a connection that never happened. Observed live in this iter's own run; reproduced
against a dead port (`old shape → [000000] len=6`, `|| true → [000] len=3`).

A sibling sweep (§5 rule 9) found the same double-source at `stack-verify/lib/services.sh:132` and the same
`|| echo "000"` shape at `readiness.sh:151,185` — but at those three sites the concatenation only **degrades
the message** (`HTTP 000000 (unexpected)` instead of *connection refused*); the verdict still comes out
`down`. **Only the playthrough site flips the verdict.** Routed with that distinction attached rather than
landed, because this iter had already opened its third line — see the close.

## Close — 2026-08-01

**Outcome:** Gate clause 2 measured for the first time — **20 live / 10 failing / 1 unimplemented, NOT MET** —
and its largest root cause found and fixed: the per-stack Directus content replay had failed on **every** cold
cycle of clause 1 (`directus=skipped(error)`), taking the post-replay container restart with it, so the whole
content layer 403'd while autoverify graded the stack green three times. Two independent defects, both landed:
a **plan-dependent** catalog query that raises under a generic plan (fenced with a MATERIALIZED barrier + OID
identity so it cannot be expressed again), and a structure script that emits two thirds of a `serial` (the
missing `OWNED BY` now reconciled from the target's own catalog at apply time). Proven live: replay
**rc=1 → rc=0**, 0 → **11 986 rows**, two collections **403 → 200**. Closes
`FIX-M257-stacksnap-directus-sequences`, carried since M257 iter-02.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n (clause 2 measured and NOT met; 2/5 clauses still hold) — (2)
triggered-tok: n (this iter moved a measurement and landed a fix) — (3) re-scope: n (still occurrence 1 of 2 —
no platform commit invalidated anything this iter) — (4) user-blocker: n — (5) cap-reached: n (1 tik this
session; the cap is 5) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1–D5 (iter-15/decisions.md)
**Side-deliverables:** none landed. The `000000` readiness defect was measured and routed, not fixed — see
above and the routes below.
**Routes carried forward:**
- `FIX-M257x-iter15-directus-versions-403` — the 4th cause, and the one still blocking both skill-path
  Playthroughs. `/versions` answers 200 anonymously and the read+create grants sit on the public policy, so
  the refusal is about HOW the cms domain asks, not whether the grant exists. Start there.
- `FIX-M257x-iter15-library-category-expansion` — 106 occurrences, the largest class. The app reads
  `library_category` as an expanded object; Directus returns a string. Suspect the replayed
  `directus_relations` / alias-field rows against what post-cms-in-app `app` expects.
- `CHECK-M257x-iter15-manager-reads-empty` — 5 Playthroughs; the seeded hero is absent from every manager-side
  aggregate. Do NOT assume it shares a cause with the content layer; measure the query.
- `CHECK-M257x-iter15-orgadmin-role-create` — a 60 s `waitForURL` timeout, no error text yet.
- `FIX-M257x-iter15-readiness-000000` — the readiness gate that announces ready on curl's own failure code;
  one verdict-flipping site + three message-degrading siblings, all enumerated above.
- **`DOC-M257x-iter15-autoverify-blind-to-content`** — clause 1 is met *as written* and the stack it certified
  had no served content. The clause is not wrong, but the corpus should record what a green autoverify does
  and does not assert, and the Directus probe should read an ITEM, not a registry count.
**Lessons:**
1. **Read the bring-up transcript before theorising about the stack.** The 403 invited a permissions
   investigation; the answer was one line in a log written three times, saying the replay had failed. Two
   probes' worth of reasoning were spent before opening it.
2. **A query that is correct in `psql` can be broken in the program.** Literals get a custom plan, parameters
   get a generic one, and a qual that can be pushed down is only pushed down in the second. Reproducing a
   catalog query "the way I'd type it" is not reproducing it. Promoted to §5 rule 13.
3. **When a step's success gates a side effect, a failure costs both.** `boot_directus_step` runs only on a
   successful replay, so the failed replay silently also cancelled the restart — a second symptom with no
   second cause, and the one that made the wreckage look like a permissions bug.
4. **Comparing two runs means comparing two worlds.** Run 2 without `--reset` scored lower for a reason the
   runner's own header documents. State the world with the number, like §5 rule 12 says to state the
   invocation.

## Addendum — the wider-suite sweep, and a baseline nobody had

Every section re-run at close. Four were exactly at their recorded baseline: `stack-injection` **OK**
(277, 1 skip) · `stack-core` **14 failures** (the m255/m220 batteries, pre-existing) · `demo-stack`
**7 failures** (`CHECK-M257x-live-clone-suites-red`, 1029 tests) · `dev-stack` **OK** (122). Every Go
section green, including `stack-snapshot` after the change.

`stack-verify` was **not** in the inherited baseline set, and it is **RED: 11 failures + 1 error of 224**
(`TestContainerLivenessM257`, `TestFapiProbeLadderM257`, `TestFrontendTierRegistration`,
`TestOffsetAwareness`, `TestOffsetMatrixSweep`, `TestServiceScopeFilter` — all about the expected service
set and the offset port matrix, i.e. the surface the router deletion moved).

**Attributed by measurement, not by reasoning.** "My diff only touched `stack-snapshot/cmd/stacksnap/*.go`,
so these cannot be mine" is exactly the argument this milestone has been wrong with before. Ran the six
classes in the **same consumption clone, at the same path**, first at `fast-build-m257x-iter-13` and then at
`fast-build-m257x-iter-15`:

| vantage | result |
|---|---|
| `stack-demo/rosetta-extensions` @ `fast-build-m257x-iter-13` | `Ran 53 — FAILED (failures=11, errors=1)` |
| `stack-demo/rosetta-extensions` @ `fast-build-m257x-iter-15` | `Ran 53 — FAILED (failures=11, errors=1)` |

Identical. **Pre-existing; zero regressions from this iter.** Same clone, same path, same vantage — iter-13's
lesson that *a control that silently skips the tests you are attributing is not a control* applies to path
depth as much as to skips.

The finding that survives is that **`stack-verify` has been red in a suite nobody re-ran** — the §5 rule 8
class, in the section that owns the bring-up's own probes. Routed as `CHECK-M257x-iter15-stack-verify-red`,
and recorded as the section's baseline so the next iter can attribute against a number rather than a memory.
