**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-33 — clause 5, the full read

## What ran

**The first actual measurement of clause 5 in this milestone.** 40 files
(`corpus/services/*.md` + `corpus/architecture/*.md`), **8 451 lines**, five read-only sub-agents,
~1 700 lines each, every file read top-to-bottom with a **`wc -l` positive control per file**.
**40/40 files read to their last line; none skimmed, none sampled, none unread.**

Ground truth was **derived fresh this session** from the platform clone at `2adcf71` and from
`platform-migration-status.md` (whose services table is machine-fenced against the platform's own
`repos.yml`), then handed to all five agents as one shared brief — checked in at
`iter-33/iter33-groundtruth.md`. The grading rule was fixed **before** any file was read:

> **BLOCKER = false at platform origin HEAD *and* acting on it would misdirect real work.**

Launched on iter-32's otherwise-idle wall-clock while the binding run was in flight.

## Pass 1 — 19 blockers, 52 minors

| group | files | blockers | minors |
|---|---|---|---|
| 1 — the big architecture docs | 5 | 5 | 13 |
| 2 — ai-readiness, studio-desk, studio-room, shared_libraries | 4 | 3 | 7 |
| 3 — alignment, ant-academy, clerkenstein, ai/security/frontend | 6 | 4 | 11 |
| 4 — hiring, backend, chronos, cms, jobsimulation, graphql, the map | 7 | 6 | 11 |
| 5 — the remaining 18 service docs | 18 | 1 | 10 |
| **total** | **40** | **19** | **52** |

Every blocker was **verified against platform source before being acted on** — the milestone's standing
re-derive rule. A sample re-derived by hand: `completion_status` is spelled correctly in the Ent schema;
the mirrors really are dropped; `sessions` really was renamed; `app/main.go` really registers six handlers
with no `SkillPathSessionService`; `AIReadinessContainer` really is absent from next-web source; 30 of 139
Ent schemas really carry `OrganizationMixin{}`.

## The prediction was HALF wrong, and the wrong half is the result

Two predictions were recorded in `overview.md` before any report was read:

| prediction | outcome |
|---|---|
| **10–25** blockers | **19** ✅ inside the range |
| the **router drop** will be the largest single cluster | ❌ **REFUTED — 0 blockers**, 2 minor captions |

The router sweep had already landed everywhere: banners, inline prod-only fences, HISTORICAL blocks. Four
of five groups reported it clean unprompted.

**What the corpus actually had is a different class — derived-fact rot.** Every doc states *who is merged*
correctly (harden pass 6's `ServiceDocStatusFence` holds) while naming **tables the platform dropped or
renamed**, **packages that were split out**, and **"routed forward to M219/M220" items for work that
already shipped. Representative:

- `security_compliance.md` asserted *"Every table has an `organization_id`"* and *"No cross-tenant data
  access is possible"* — while **30 of 139** Ent schemas carry the privacy policy, and the platform says
  so in its own source (`job_simulation_session.go:5`: *"L2: NO Ent privacy Policy"*).
- `jobsimulation.md` told readers to seed `public.local_jobsimulation_sessions` and `public.sessions`.
  Both were dropped or renamed by `app` migrations.
- `cms.md` **and** `jobsimulation.md` both said the M23 Directus cutover rides on the `cms` husk. iter-24
  measured that as 96 all-403 lines in `backend`'s log — **the corpus never followed its own fix.**
- `clerkenstein.md` carried a 97.2% score, an `rc=2` that now means *REGRESSED* rather than *unmeasurable*,
  and an unbounded-clerk-js perf defect — all three closed by M219/M220.
- `ai-readiness.md`'s *"there are TWO manager dashboards"* section: the legacy one was deleted at next-web
  `dae0fb2f7` and its route 404s.

**None of it uses merged/live/gone vocabulary.** The status layer is now fenced; the layer underneath it
is not. That is the studio-room archetype, generalised.

## Pass 2 — the adversarial verification, and why it was not optional

**Clause 5 is not graded on "19 found, 19 fixed."** That is a probe satisfying itself (§5 rule 7), and
iter-22's precedent is that a corrective sweep introduces defects its own pass cannot see. **This
iteration reproduced that by hand three times within minutes of applying the sweep** — a mid-sentence
anchor leaving *"The manager view does / reads the same table"*, a *"false on **two** counts"* that became
false once one count was fixed, and a present-tense *"fabricates"* under a ✅ RESOLVED heading.

So the 13 changed files went to a **second, adversarial read-only audit** briefed to catch new false
claims **the corrections themselves introduced**. It found **6 more blockers**:

**Three self-inflicted by my sweep:**

1. `security_compliance.md` — the new tenancy fence claimed the non-mixin schemas *"never mention
   organization at all"*. **33 do, and ~18 declare a plain un-policied `organization_id`.** It
   contradicted itself inside its own blockquote, and it erred **in the dangerous direction**: an auditor
   would have excluded precisely the un-policied org-scoped tables the fence exists to surface.
2. `ai-readiness.md` — I wrote that the anchors below *"no longer resolve."* **All five still resolve**;
   only the three components are gone.
3. `hiring.md` — the `completion_status` correction was **spliced into the middle of a column list**, so
   four required write-set columns read as further places the misspelling survives.

**Three the sweep missed:**

4. `architecture_overview.md:274-278` still asserted *"`organization_id` on every table; Ent ORM policies
   auto-filter queries"* — **verbatim the claim `security_compliance.md` had just retracted with
   measurements**, three files away, in the doc most readers hit first. The sweep edited that file six
   times and walked past it.
5. `hiring.md:126` still called the score column *"the mirror's"*.
6. `hiring.md:144`/`:229` still scored from *"the 2-table pair"* — inside the file whose own banner says
   the mirrors were dropped.

All six fixed. Notably, **every one of the ~40 `file:line` anchors in my new text verified exact** — the
errors were entirely in surrounding prose, which is where a sweep does not look.

**25 blockers found and fixed across the two passes.**

## Clause 5 is NOT MET — and saying so is the point

The gate wants *"KB-fidelity audit **GREEN**, or **YELLOW with 0 blockers**."* The last measurement taken
returned **6 blockers**, not 0. Those 6 are fixed, but **the confirming re-measurement has not been run**,
and this milestone does not grade a clause on the absence of a measurement — that is exactly the
`25 → 27` mistake iter-32 diagnosed one iteration ago.

The 52 (+~16) **minors do not block the gate** — *"YELLOW with 0 blockers"* admits them. The single
outstanding action is **one more full-read pass that returns zero blockers.**

**Is the curve converging or exhausting itself?** iter-21's 11→5→2 looked like convergence and was an
artifact — the instrument was term-scoped and running out of vocabulary. **This instrument is not
vocabulary-bound**: it is a full top-to-bottom read with a per-file positive control, so 19 → 6 is a
different kind of number. That is an argument for expecting the next pass to be small — **not** a
substitute for running it.

## Evidence

`iter-33/iter33-groundtruth.md` (the shared brief) · `iter-33/evidence/audit-g{1..5}.md` (the five pass-1
reports, with their positive-control tables) · `evidence/sweep33{,b,c}.py` (the three enumerated
exactly-once-anchor sweeps).

## Close — 2026-08-02

**Outcome:** clause 5 **measured for the first time** — 40/40 files read in full, **19 blockers + 52
minors**, all 19 fixed; then an adversarial re-audit of the corrections found **6 more** (3 self-inflicted,
3 missed), also fixed. **25 blockers closed.** Clause 5 stays **NOT MET** pending one confirming
zero-blocker pass.
**Type:** tik
**Status:** closed-fixed (planned scope was *measure clause 5, then fix by evidence rank*; both landed,
plus the verification pass the milestone's own rules required)
**Gate:** NOT MET (**3 of 5**. Clause 5 is now *measured and largely repaired* rather than unmeasured —
but a clause is not met by an absent measurement)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-fetched at open, unchanged) — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) —
(6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-33-1 … D-M257x-33-6 (`iter-33/decisions.md`).
**Side-deliverables:** none. **No rext change and no re-pin** — clause 5 is entirely a rosetta-corpus
clause; the pin stays `fast-build-m257x-iter-31b`.
**Routes carried forward:**
- `MEASURE-M257x-iter34-clause5-confirming-pass` — **the one action that closes clause 5.** A third full
  read; expected small, must not be assumed empty.
- `DOC-M257x-iter33-corpus-minors` — 52 (+~16) minors: stale `file:line` refs widened by the sweep's own
  line shifts, imprecise counts (`139` is a *file* count, 135 are schemas; `9` ai-readiness tables are 13),
  residual "CMS service"/"Jobsimulation service" phrasing in table cells, one missing blockquote prefix.
- `CHECK-M257x-iter33-derived-fact-fence` — the status layer is fenced and the layer under it is not. A
  fence over *derived facts* (table names, package paths) is the structural answer; unclear one exists.
- `CHECK-M257x-iter27-drilldown-target-coupling` — unchanged, still the best-evidenced clause-2 target.
- Unchanged: the other two clause-2 survivors, and the standing ledger.

**Lessons:**
- **A corrective sweep is text, and text is audited.** Six of 25 blockers were created or left by the
  repair pass itself — a 24% self-inflicted rate. Fixing without re-auditing would have shipped a
  tenant-isolation fence that excluded the very tables it existed to surface.
- **The dangerous direction is the one to check.** Both tenancy defects — the original absolute and my
  over-correction — failed *toward* "isolation is handled." A finding's direction matters more than its
  count.
- **A fence closes a layer, not a subject.** `ServiceDocStatusFence` made merged-status claims reliable
  and thereby made the *next* layer down — table names, package paths, shipped-vs-routed status — the
  place drift accumulates unobserved.
- **Convergence means something different when the instrument is not vocabulary-bound.** iter-21's
  11→5→2 was an artifact of grep vocabulary; 19→6 from full reads is evidence. It still is not zero.
- **Predict two things, so one can be wrong.** The blocker-count prediction held and the cluster
  prediction was refuted — and the refutation is what corrected the milestone's model of its own corpus.
