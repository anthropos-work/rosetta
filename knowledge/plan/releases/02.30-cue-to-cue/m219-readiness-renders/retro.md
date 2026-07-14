# M219 "readiness renders" — Retro

## Summary

**The user's bar was met on both counts** — *"make sure each element and sub section of readiness is filled spot
data + make sure to use the right dashboards/pages for readiness (both for managers and employees).. not the old
legacy ones"* — proven on **5 cold reset-to-seed rebuilds** at `cue-to-cue-m219-r8`, each independently evidenced.
**Zero platform-repo edits.**

The milestone's plan said the AI-readiness page was blank because of a read-path bug needing a demo-patch. **That
was wrong.** The page was blank because **every demo pointer aimed at an unlinked orphan** — a legacy route with no
nav entry, no tab, and no redirect — and because four of its sections read a table **no seeder had ever written**.
Both of the plan's named blockers were **REFUTED by measurement**, and the planned demo-patch was **WITHDRAWN**.

**Two premises of the milestone died on contact with a measurement. The user's report survived it intact.**

## Incidents this cycle

| # | What | Class |
|---|---|---|
| **F-1** | **All 3 demo pointers targeted `/enterprise/workforce/ai-readiness` — an UNLINKED ORPHAN** (no nav, no tab, no redirect; its hook takes no `cycle` param). The current surface is `/ai-readiness`. The **employee surface has no route of its own** — it is embedded in `/home`, which is *why* route-crawling never found it. | P1 · wrong vantage |
| **F-2** | **REFUTED: the `CycleID == nil` blocker does not exist.** The CURRENT frontend passes `?cycle=`; the original note was made against the **LEGACY** page. **The planned demo-patch was WITHDRAWN.** | P1 · false premise |
| **F-7** | **REFUTED: M217's `loadmembers` patch is not dead** — it **self-heals**. And *"the live recompute never completes"* is **false**: **2.09 s**, measured. | P2 · false premise |
| **D1** | **Junk skills ORG-WIDE.** The claimed-tail pool ran **DRY** and topped up from the flat pool's **alphabetical head**: Aria + 8 named members claimed **"24-hour dietary recall"** / `15Five` / `17Track`. | P1 · silent fallback |
| **D2** | **A hero rendered ROLE-LESS.** `Operations Analyst` **resolves** — and carries **0 `job_role_skills`**, so the seeder's own resolver rejected it ⇒ `job_role_id` NULL. | P1 · "resolves" ≠ "has skills" |
| **D3** | **The manager's 4 interview-findings blocks rendered HEADINGS OVER NOTHING.** They read `jobsimulation.interview_aggregated_reports` — which **NO SEEDER EVER WROTE** (`git grep` @ `ffc6ffe`: 0 refs). **The gate passed them under a disclosed EXCEPTION.** | P1 · disclosed empty |
| **F-13** | The bring-up reported the academy **"started"** while it was dying — a bare **502 for the life of every stack**. The node check tested **existence**, not `engines: >=22`; the liveness probe polled `kill -0 $pid`. | P1 · stale verdict |
| **F-14** | **The coverage sweep had never been EXECUTED against a live demo** — the harness could not even be *pointed* at a remote one. | P1 · unrun assert |
| **H-1** (harden) | **`run-coverage.sh` presented the PREVIOUS run's numbers as the current run's.** | P1 · **stale verdict** |
| **H-2** (harden) | **17 existing tests asserted the junk fallback AS THE CONTRACT.** | P1 · **the test was the bug** |
| **RED → M220** | The academy **poisons the demo session** (demo-BREAKING) · studio-desk **302s the presenter out of the demo**. | P1 · routed, fenced RED |

## The D17 thread — the spine of this milestone

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** Named in M218. It
has now bitten **~14 times across M217 → M219**. Five new instances landed **inside M219 — and the tooling found
several of them in itself:**

1. **`run-coverage.sh` printed *"coverage report written to …"* UNCONDITIONALLY**, and its summary re-read whatever
   `coverage-report.json` was on disk. A spec that threw before assembling a report wrote **nothing** — so the
   script presented the **previous run's numbers, *"GATE: MET ✅"* and all, exiting 0**. **It nearly graded an M219
   rebuild on hours-old data from the old, broken stack.** Fixed by **deleting the report first**, so *"the file
   exists afterwards"* is **equivalent** to *"this run wrote it"* (rext `b5bf65b`).
2. **17 existing seeder tests asserted the junk-fallback as the contract** — *"unmatched role must fall back to flat
   pool"*; an expected value **literally containing `K-JUNK-1`**. **They were not missing the bug. They were
   guarding it.** Any correct fix was *guaranteed* to fail them. All inverted.
3. **The poisoned-pool fence's FIRST cut PASSED against the broken code** — it modelled no name attrition, so the
   junk tier never fired in the fixture. **Theatre, inside the fix for the very defect it could not see.**
   Strengthened until it went **RED pre-fix**.
4. **The interview-findings fence asserted the LEGACY page's strings — and passed.**
5. **The ant-academy launcher read *"a pid exists"* as *"the service is up."*** Its **test fixtures had encoded the
   broken behaviour** (an npm stub that `exec sleep 30` — alive, serving nothing), so **the suite was GREEN for four
   releases while the academy 502'd.**

**And the milestone did it to itself, twice, in its own status artifact**: `progress.md` claimed *"zero new
demo-patches"* while one had already landed (caught by the harden pass, not the close), and it claimed the battery
ran *"at final code"* — **true when written, false by the close** (see *Carried forward*).

### The generalized lessons — the sentences worth keeping

> **"The role classifies" ≠ "the pool is big enough."**
> **"It resolves" ≠ "it has skills."**
> **"It serves" ≠ "it renders."**
> **"The pool resolved" ≠ "the content is sane."**
> **An errored SQL query is not "zero rows."**
> **An unexecuted gate is a FINDING, not a pass.**

Every one of these is the same shape: **a proxy was checked, the proposition was not.** A green test proving a thing
*exists* is not a test proving it is *used* (**D-M219-14**) — and only a cold reset-to-seed ever exposed the
difference.

## What went well

- **Measurement beat the plan, twice.** Two of the milestone's own premises (F-2, F-7) were **refuted by looking**,
  and the planned demo-patch was **withdrawn** rather than built. A patch that fixes a bug that does not exist is
  indistinguishable, in a status report, from one that works.
- **The arithmetic closed exactly, and that is what proved the root cause.** D1: `want`=28, role pool=10, curated
  `data` supplied only **25 usable** ⇒ **exactly 3** junk tokens. **Ben was clean only because his `want` (16) fit
  his 30-name pool — and that asymmetry was the proof** the defect was pool **SIZE**, not pool **resolution**. No
  allow-list, however long, could have fixed it.
- **Fences graded by MUTATION, not by re-running them green.** 7 mutations, **7 REDs**. The release's own rule is
  that *a fence which passes against both the pre- and post-fix code is theatre* — and this milestone caught its own
  fence being exactly that (instance 3 above).
- **The disclosed EXCEPTION was DELETED, not re-worded.** A disclosed empty is better than a hidden one; **it is
  still an empty section.** The content floor went 120 → **900** — the empty headings measured ~120–200 chars, which
  is precisely *why* 120 passed over them for four releases.
- **Honest RED over convenient green.** The academy and studio-desk fences **deliberately report RED** until M220
  lands. A half-fix that reports green is worse than no fix.

## What didn't

- **We keep naming the hazard and then walking into it.** D17 was named in M218's close. M219 then hit it **five
  more times** — twice **in its own progress.md**, once **inside its own fence**, and once in **the harness convened
  to grade the milestone**. **Naming a hazard does not inoculate you against it. Only a probe does.**
- **The lesson was never propagated one directory over.** `run-playthroughs.sh` had already hit and fixed the
  identical stale-report class in **M204**. `run-coverage.sh`, one directory away, shipped the same bug and **nearly
  graded this milestone's rebuild on the old stack's numbers**.
- **The battery's audit trail was corrupted by an ORCHESTRATION error, not a demo defect.** Two batteries were run
  **concurrently against the single demo host**; one purged the stack **mid-measurement**. It cost hours, and it
  means the 5 greens — each individually evidenced — **are not one uncontested consecutive run**. The tooling treats
  a **singleton shared resource** as if it were private. → `GUARD-M221-host-isolation`.
- **The code that graded is not the code that shipped.** `c6648d1` is a **seed-path** change that landed **after**
  the graded tag. Per **D13** that restarts the battery count. The delta is small and strictly corrective — **and it
  is still not allowed to be rounded away.** → `REPROVE-M221-battery-at-final-code`.
- **Checking the heroes proved the wrong proposition** (the fifth instance of that shape). R-2 declared the Step-1
  arithmetic *"matches `computeTier1` exactly"* — because it checked **the two heroes**, and both happen to round
  **identically under both formulas**. The other ~198 members do not.

## Carried forward

| Item | Destination | Severity |
|---|---|---|
| **The academy POISONS the demo session** — cookies scope by **HOST, not PORT**; a presenter who clicks the academy link is **logged OUT of the demo** into `ERR_TOO_MANY_REDIRECTS`, and every employee coverage sweep **aborts** | **M220** — item **(i)**, severity **RAISED** | 🔴 **demo-BREAKING** |
| **studio-desk `:19000` → 302 → `:13000/login`** — clicking *"Anthropos Studio"* **ejects the presenter from the demo** | **M220** — item **(j)**, added at close | 🔴 **demo-BREAKING** |
| **`GUARD-M221-host-isolation`** — two agents can run cycles against ONE demo host and nothing stops them. **A prerequisite for M221's own gate**, which is itself a multi-cycle battery on that same singleton host | **M221** | P1 |
| **`FIX-M221-reap-native-academy`** — `down --purge` doesn't reap the host-native academy; the OLD process keeps answering `:13077` while the log says it *"DIED"* | **M221** | P1 |
| **`REPROVE-M221-battery-at-final-code`** — the shipped code is **not** the code the 5 greens graded (`c6648d1`, seed-path, post-`r8`) | **M221** | P1 |

**Every receiving `overview.md` was EDITED** — Fate 3 means the sibling's *plan* changes, not that a string changes
in ours. **Zero escape-hatch deferrals; nothing leaves v2.3.**

## Metrics delta

| | M217 | M218 | **M219** |
|---|---|---|---|
| Python tests (passing) | 867 | 887 | **903** (0 fail, 12 skip) |
| Go test funcs | 1750 | 1784 | **1821** (+37, same method) |
| Go failures | 0 | 0 | **0** (6 modules) |
| Alignment (Go surface) | 100% / 100% *(hollow)* | 97.2% / 100% *(honest)* | **100% / 100% critical (27/27)** — the RED gene **fixed**, and **retained as a fence** |
| Alignment surfaces measurable | 4 of 5 *(silently)* | 4 of 5 *(disclosed)* | **an unmeasurable surface now FAILS LOUD** (exit `3` ≠ exit `2`) |
| Demo-patches | 5 | 6 | **7** (`next-web-aireadiness-flag-gate`) |
| Platform-repo edits | 0 | 0 | **0** |
| Bugs fixed | 32 | 21 | **13** (build/R-8 9 · harden 3 · close 1) |

## The one line worth keeping

**A test that proves a thing exists is not a test that proves it is used** — and a test suite can be **green,
comprehensive, and entirely wrong about what it is protecting.** Seventeen of them were guarding the bug.
