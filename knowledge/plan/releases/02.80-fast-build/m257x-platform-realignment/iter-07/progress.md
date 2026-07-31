---
milestone: M257x
iter: 07
---

# iter-07 — progress

**Type:** tik (under `TOK-01`, step 2 — *fix the mechanism, not the symptom*)

## What was found before anything was changed

The pre-compute (committed at `e3189bc`) held on every point. The **re-survey added one fact it did not
have, and that fact decided the design**.

### The snapshot cache key does not contain the schema name

`pg.SchemaVersionSQL` digests `table_name || '.' || column_name || ':' || data_type`. The schema is only the
`WHERE` filter — it never enters the digested string. So for a **narrowed row surface** the staleness key is
**schema-independent**. Measured on live `demo-1`:

    digest over public, narrowed to the 4 similarity tables : 032c99ea47678187631c59c31b4ef059
    digest over cms,    same 4 tables                       : <null — cms held 0 tables>
    cached manifest schema_version (captured 2026-06-29)    : 032c99ea47678187631c59c31b4ef059

An exact match. Three things follow, and the first two were open questions the pre-compute explicitly could
not settle:

1. **The 2026-06-29 capture is not stale.** The pre-compute could only compare column *names*; the digest
   covers name **and type**, for all four tables, order-independently. "Re-capture freshness is a separate
   question that cannot be settled from this box" is **answered — no re-capture is needed.**
2. **The cache HITS the moment the probe reads the right schema.** The whole failure was a *resolution*
   failure. Nothing was wrong with the data, the freshness, or the cache.
3. It therefore had to move exactly two things — the **probe's** schema and the **replay's** — and neither
   the capture's (`D-M257x-7`).

### The write surface, scanned and split live-vs-comment (§7 rule 1)

Exactly **one** live site named the schema: `stack-snapshot/simembeddings/simembeddings.go:44`
`const Schema = "cms"`. All four other occurrences in the tree (`dev-setdress.sh:140,340,342`,
`repos_yml.sh:97`) are comments. The smallest write surface of the three folds — and the hardest, because
that single constant is read by *both* halves of a system whose halves now disagree.

## The design decision, and why the easy answer was rejected

The pre-compute deliberately left this open, warning that the easy answer reproduces this milestone's own
defect. It does:

**Rejected — `ReplaySchema = "public"` on the surface.** Two lines. It is *the same hand-maintained constant
that has now been wrong three times*, and it would be wrong again at the v9.0 fold — silently, with the
bring-up reporting success.

**Adopted — a derived, replay-time resolver.** For a surface's declared schema `S` and its table set `T`,
ask the **target** (never the source):

| the target says | resolution |
|---|---|
| every table of `T` is in `S` | **identity** — taxonomy and directus are untouched |
| none in `S`, and **exactly one** other schema holds **all** of `T` | **remap**, announced LOUDLY |
| **no** schema holds all of `T` | fail loud — `ErrSurfaceNotOnTarget`, exit 4 |
| **two or more** do | fail loud, **naming them** — `ErrAmbiguousTargetSchema`, exit 1 |

No allow-list, no preference for `public`, no fallback to the declared value. Each of those is a form of the
same defect — a check reporting a state it did not measure — and an allow-list in particular is Trap A in
miniature (tune it until it stops catching what it exists to catch).

Landed as `replay.TargetSchema` + `replay.ResolveTargetSchema` (pure, DB-free) over
`pg.SchemasHoldingAllTablesSQL` (one catalog query). `replay.Run` gained an **explicit positional**
`TargetSchema` parameter rather than an option — so the compiler forced all 18 existing call sites to state
which they meant. A forgettable optional with a silent identity default is the shape this milestone exists
to end.

## Two findings that came out of doing it

### 1. A test's own comment caught a regression I had just introduced

The first cut resolved the schema **eagerly**, at the top of `replayCmd`. Three pre-existing tests went RED,
and their comments explained why better than I could have reconstructed:

> *"the DSN is parsed … but never DIALED: pgxpool connects lazily and the explicit `--schema-version` skips
> the probe, so no live DB is needed before the store resolve."*

`--schema-version` exists precisely to say *"do not ask the stack for the digest"*, and the store resolve
after it is a pure local-filesystem question. Resolving eagerly made a **cache-miss verdict require a
database** — and, worse, silently reclassified it from exit 5 (a capture fixes this) to exit 4 (a capture
cannot help). That is the exact conflation `fix16` split those two exit codes apart to stop. The resolution
is now lazy and memoized.

Worth naming because of *how* it was caught: not by a probe, but by a test that had written down the
contract it was protecting. The three-line comment above a fixture did the work.

### 2. The probe and the replay could not be allowed to read the schema separately

The pre-compute's "thing not to miss (a)" — that the pre-replay probe reads the manifest schema too, and the
surface would still skip at `rc=4` before any copy if only the replay moved. Rather than move both and rely
on review to keep them together, the two are now **one function**, `resolveThenProbe`, which computes the
probe's schema argument *inside itself* from its own resolution. There is no parameter for a caller to pass,
so there is no way for a caller to pass the wrong one. That is `§8 rule 1` — assert against the construct —
applied to the code rather than to a test.

## Fences — 5, each mutation-verified RED

| # | mutation | fence that fired |
|---|---|---|
| M1 | probe uses the **declared** schema (the half-done re-point) | `TestResolveThenProbe_ProbesTheRESOLVEDSchemaNotTheDeclaredOne` |
| M2 | `CopyIn` writes to the manifest schema, not the resolved one | `TestRun_RemapWritesEveryDBCallToTheResolvedSchema` (+1) |
| M3 | ambiguity resolved by picking `candidates[0]` | `TestResolveTargetSchema_AmbiguityFailsLoudAndNamesTheCandidates` |
| M4 | a failed catalog lookup degrades to `Identity()` | `TestResolveReplaySchema_LookupErrorIsNeverDegradedToIdentity` |
| M5 | candidate query matches **ANY** table instead of ALL | `TestSchemasHoldingAllTablesSQL_RequiresALLTablesNotAny` |

**M3's first run was a false RED and was re-run.** The mutation removed the last use of the `strings` import,
so the package failed to *compile* — and a compile break reports as a failing test run while proving nothing
about the fence. It was re-done with a compiling mutation (`_ = strings.Join(...)`), a build check was run
*before* the test run, and only then did the fence's RED count. Same family as `§5 rule 8`: a run that fails
for the wrong reason reads exactly like a run that fails for the right one.

None of the tests pin the string `"public"`. They pin **where the schema comes from** — `§8 rule 3`, pin the
mechanism not the contents. If the next fold moves these tables somewhere else, every one of them still
passes.

## Proven live on `demo-1`

    BEFORE  sim-embeddings replay skipped (rc=4)
            public.similarities 0 / similarity_categories 0 / similarity_features 0 / similarity_skills 0

    AFTER   stacksnap: ⚠ surface "sim-embeddings" was captured from schema "cms", but on demo-1 its
            tables live in "public" — replaying into "public".
            replayed "sim-embeddings" into demo-1: 4 table(s) cleared, 4 table(s), 1490 row(s) loaded,
            schema cms -> public, reindexed [public.similarities.small_embedding3]

            public.similarities 274 / similarity_categories 278 / similarity_features 274
            / similarity_skills 664          (= 1490, exactly the manifest's counts)

Re-run immediately: still 1490, not 2980 — the M17 re-run guard holds through a remap.

Then the stronger proof: **`DROP SCHEMA cms RESTRICT`** on demo-1 (non-CASCADE, so it would have failed loud
had anything been in it — it held 0 tables), and the replay run again. **Still 1490 rows.** The surface no
longer needs the schema to exist at all.

## The debt, paid

`REXT_TRANSITIONAL_SCHEMAS` went `"cms"` → **empty**, and the no-growth fence's shrink branch fired exactly
as designed:

    AssertionError: Debt paid down (['cms'] re-pointed) — update _EXPECTED_TRANSITIONAL in this fence
    to lock the win in. This failure is GOOD NEWS and the fix is a one-line edit.

Second and last firing. The derived CREATE SCHEMA set is now **`extensions · sentinel · public`** — infra
(2, both justified) plus exactly what `repos.yml` declares. **rext no longer creates a single schema the
platform does not own**, and that is asserted by the fence rather than narrated here.

The variable and the fence are both KEPT, empty. v9.0 folds `storage` + `messenger` with PRs already open;
the next debt entry should have to argue with a fence, not land as a one-word diff.

## The verifier went RED in between, correctly — and then GREEN

With `cms` dropped but the debt not yet paid, `autoverify` on demo-1 reported
`postgres-schemas fail: missing schemas: cms` — the derived expected set still demanded it. That is iter-05's
derived-expected-list machinery working: the two halves must move together, and the verifier says so.

After the paydown, tagged (`fast-build-m257x-iter-07`, **verified on origin**) and the `stack-demo`
consumption clone re-pointed to it — i.e. measured through the same instrument a real bring-up uses, not the
authoring copy:

    ✓ verify live: all liveness + readiness probes passed

`postgres-schemas` is GREEN **with the `cms` schema physically absent from the stack**. autoverify's failed
count went **4 → 3**, and the residual 3 are: 2 × the evidence-absence warnings routed below, and
`FIX-M257x-academy-not-serving` — the one already-known clause-1 blocker.

**Measurement-hygiene note.** The first autoverify run of this iter was made from the **authoring** clone and
reported a *different* `postgres-schemas` failure — `cannot derive the expected schema set (missing …
platform/repos.yml)`. That is a path artifact of the authoring copy having no sibling `platform/`, not a
stack defect. Comparing it against iter-06's baseline would have been a fidelity check against the wrong
reference (Trap A). The instrument for a stack claim is **the stack's own pinned clone**.

## The wider suite run (the standing protocol note — three consecutive iters found their finding here)

| section | result |
|---|---|
| `stack-core` | **354 passed, 0 skipped** |
| `stack-snapshot` (Go) | all packages green; `go vet` + `gofmt` clean |
| `demo-stack` | 1011 passed, **7 failed** — bit-identical to `CHECK-M257x-live-clone-suites-red` (already reproduced on the pristine control clone). No regression. |
| `dev-stack` | 122 passed |
| `stack-verify` | 210 passed (matches iter-06) |
| `stack-injection` | 267 passed, **9 skipped** → see below |
| `stack-seeding` / `alignment` (Go) | green |

*(Sections with no `tests/` dir — `stack-secrets`, `clerkenstein`, `alignment`, `stack-seeding`,
`stack-snapshot`, `playthroughs` — were enumerated explicitly rather than assumed absent, so the sweep's
coverage is a measured claim.)*

**And it did produce a finding, again — in the skip count.** 8 of `stack-injection`'s 9 skips read
`PyYAML not available`: eight compose/YAML-shape tests had been silently not running on this host. `§5 rule
8` says a skip is a hole in the evidence, not a pass. Installing PyYAML into the rext venv turned them into
**275 passed, 1 skipped** — all 8 green, so no defect was hiding, but the evidence now exists rather than
being assumed. Another slice of `HOST-M257x-toolchain`, and one nobody had counted.

The 1 remaining skip is structural: `test_apply_patch_selfheal.py` needs an `app` clone at
`stack-demo/app/…`, which a demo deliberately does not keep (it builds `app` from a scratch clone). Named
here so it is a known hole rather than a silent one — and it is in the demopatch family, the class that
shipped a 76 s members grid for four releases.

`demo-stack`'s **2** skips were also read rather than counted, and both are legitimate and self-describing:
`test_interview_flag_patch_m232.py` (needs `stack-dev/next-web-app`, absent by design here) and
`test_purge.py` (*"the container-owned UID-1001 0700 layout this regression needs is unreachable here; the
defect and this test are Linux-host-only — verified live on billion"*). **Full sweep skip ledger: 11 skips,
11 named, 8 of them closed.**

## Why gate clause 4 is NOT being claimed, despite the debt list being empty

This is the iter's most important finding and it came out of the follow-up sweep, not the planned work —
the third consecutive iter for which that is true.

Clause 4 reads: *"zero rext writes to a schema the platform no longer creates, **asserted by a FENCE that is
watched going RED**, not by inspection."* The debt list is empty and the live probe is green, so the
*condition* holds. But the fence that is supposed to assert it does not cover what its name says:

    stack-core/tests/test_write_target_schema_fence.py:92
        SCORED_SECTIONS = ("stack-seeding",)

**One section of nine.** Its own docstring opens *"Every schema a rext artifact WRITES TO must be a schema
rext's own migrate step CREATES"* — and `stack-snapshot`, which is where **this entire iter's `cms` write
surface lived**, is never scanned. Nor are `demo-stack`, `dev-stack`, `stack-verify`, `stack-injection`,
`stack-secrets`, `clerkenstein`, `alignment`.

So the fence would **not** have caught the defect this iter just fixed — and would not catch its equivalent
in seven other sections at the next fold. It is a correct fence with an unstated boundary, which makes it
exactly the milestone's dominant class: **a check that reports a state it did not measure.** The scope
limit is real and may even have been deliberate at iter-06 (the seeders are where the previous three
occurrences bit); what is missing is that nothing says so, and clause 4 currently rests on it.

Claiming clause 4 on this fence today would be claiming a measurement that was not taken. Routed as
`FENCE-M257x-write-fence-scans-one-section-of-nine`, and clause 4 stays **unclaimed** until it is closed —
which is a one-iter job, not a milestone-sized one.

## Routes carried forward

| item | why | target |
|---|---|---|
| `FENCE-M257x-write-fence-scans-one-section-of-nine` **NEW — and it gates clause 4** | See the section below. `test_write_target_schema_fence.py:92` is `SCORED_SECTIONS = ("stack-seeding",)`. **`stack-snapshot` — where this iter's entire `cms` write surface lived — is not scanned at all**, nor are the other seven sections. The docstring claims *"Every schema a rext artifact WRITES TO"*. Widen the scope (or narrow the claim) before clause 4 is called met. | **next tik** |
| `HOST-M257x-toolchain` **NOT closed — the green is a warm cache** | `stack-seeding`'s `services/ai` tests now PASS with `GOPRIVATE` **empty** and no `insteadOf` rewrite — because `ai v1.40.1` is already in `$GOMODCACHE` (`/Users/marco/go/pkg/mod/github.com/anthropos-work/ai@v1.40.1`). Confirmed with `GOPROXY=off`. On a cold module cache the Trap-E failure is unchanged. Stopping at "the tests pass" would have closed this deferral on evidence that does not support it. | later tik |
| `CHECK-M257x-bringup-evidence-logs-absent` **NEW** | `autoverify` on demo-1 warns that **neither** `demopatch.log` nor `buildfail.log` exists under `$STACK_DIR`, so it "cannot claim the demo is patched" or "that the running images are this run's". Confirmed absent anywhere under the workspace. Pre-existing (this stack was brought up in iter-04) and NOT caused by this iter — but it means two of autoverify's four current ✗s are *evidence-absence*, not defects, and the next cold cycle should either produce the files or the check should say which. Bears directly on reading clause 1's "0 warnings". | next cold cycle |

## Close — 2026-07-31

**Outcome:** the last transitional schema is paid off — `sim-embeddings` replay went `rc=4 skipped` → **1490
rows** into `public` on demo-1 with the `cms` schema **dropped from the stack entirely**, fixed by a
**derived** replay-time resolver rather than a second declared constant; `REXT_TRANSITIONAL_SCHEMAS` is now
**EMPTY** and rext creates no schema the platform does not own — **gate clause 4 is claimable**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-8` (derive the replay schema, do not declare it), `D-M257x-9` (the probe and the replay resolve as one construct)
**Side-deliverables:** none — every change served the planned target.
**Routes carried forward:** `CHECK-M257x-write-fence-blind-to-const-schema` · `HOST-M257x-toolchain` (re-opened with better evidence) · `CHECK-M257x-bringup-evidence-logs-absent`
**Lessons:**
- **A cache key that omits the schema name makes a schema re-point free.** Worth checking for *before*
  designing a migration: the question "is the cached artifact still valid after the platform moved this?"
  had a measurable answer (an md5 comparison) and it was `yes`. The pre-compute had assumed it was
  unanswerable from this box.
- **A mutation that does not compile is not a RED fence.** Build the mutant before you trust its failure —
  otherwise a compile error signs off on a fence that may not fence at all. Added to `platform-alignment.md`
  §8 as rule 5.
- **When two consumers must read the same derived value, make them one construct** rather than moving both
  and relying on review. `resolveThenProbe` has no parameter to get wrong.
- **A test comment can be the best documentation of a contract in the repo.** The lazy-resolution regression
  was caught, and *explained*, by three lines above a fixture.
