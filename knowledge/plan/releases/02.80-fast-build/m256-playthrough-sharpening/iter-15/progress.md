**Type:** tik · shape: standard, with a DECLARED three-step scope (see [`overview.md`](overview.md))

# iter-15 — the fence had missed a shape, and a half-dead stack proved why the assertion mattered

## Phase A — probe first (five passes; two of them refuted an earlier iter's conclusion)

### (1) The drill-down discriminator — confirmed

Drilling the first content row as Org A's manager yields `Pat Ellis / DevOps Engineer` + `Morgan Reyes /
Engineering Manager` (2 rows); drilling **the same content id** as a manager of another seeded tenant yields
`Mei Costa / Advanced Civil Engineer` + `Theo Lindqvist / Business Operations Analyst` (5 rows). The breakdown
names members **with their roles** and is org-scoped, so a cross-tenant control is available. Measured twice
on separate runs, which also settles the determinism question: the grid sorts by most-recent activity and the
seeded heroes' sessions are written at reset (dated today), so the first row is a hero-session content.

### (2) The method-panel scope — the first hypothesis was REFUTED by the DOM

The plan proposed scoping the step assertions inside a panel anchored on `methodHeading()` ("The three
inputs"). Measured: that string is a **leaf** `<div>` whose entire `textContent` is those three words
(length 16), and **0** of the step-name elements sit under it. It is a label, not a container.

The real container is a `<div>` whose `textContent` *starts with* `"How the score is built"` and runs
`"…Step 1AI Skill Mappingwith AI Framework98%completed Step 2AI Simulation…"`. On a NOT-enabled org that
string is **absent entirely** (its `/ai-readiness` renders a marketing upsell headed "How it works"). So:

| vantage | page-wide in `main` | scoped to the built-panel |
|---|---:|---:|
| ENABLED (has a cycle) | 2 / 3 / 3 | 2 / 3 / 3 |
| NOT ENABLED (upsell) | 1 / 1 / 1 | **0** |

The scope costs the Playthrough **nothing** where it runs and removes exactly the matches that made it
vacuous where it must fail. iter-12's finding is confirmed verbatim: the upsell renders
`"STEP 1 / AI Skill Mapping with AI Framework"`, so a substring match hits.

*(An intermediate probe pass counted only elements whose text was EXACTLY the step name and got 0 on the
upsell — which looked like a refutation of iter-12 and was not. The accessor under test matches a **regex
substring**; a probe must use the predicate the code uses, not a tidier one.)*

### (3) The finding the probe was not looking for: an unbounded click, 10 minutes

Driving `openHowWeMeasure()` on the non-enabled vantage **hung for a full 600 s** and then reported
`locator.click: Test timeout exceeded` at `ai-readiness-page.ts:130`. Playwright's action timeout defaults to
`0` — no timeout — so the click inherited the whole test budget while waiting for a tab that **does not
exist on that org** (0 elements, measured).

**This is the 7th site of a class the harden pass reported discharged at 6, with its exception boundary
written down** — and the boundary named this very line. Details in D74.

## Phase B — three deliverables

### B1 — the bounded-interaction fence, widened on its own stated condition

The fence's exception boundary read: *"straight-line interactions elsewhere … are OUT OF SCOPE … If a
straight-line site ever produces an opaque hang, it becomes evidence and this boundary moves — with the
measurement recorded."* Its self-test (b) then quoted **`getByText(/How we measure/i).first().click()`
verbatim** as its example of a safe straight-line click. That is the line that hung for ten minutes.

So the boundary moved, and the widening was **measured before it was adopted** rather than argued: the new
rule flags **7 sites across 5 files** — all bounded in this iter — where "bound everything in `lib/`" would
have been ~28 evidence-free edits, which is how a fence becomes noise and then gets switched off.

**The rule: every interaction inside a DISCLOSURE METHOD** (`open*` / `switchTo*` / `expand*` / `reveal*` /
`drill*`) must carry an explicit timeout. The harmful property is not straight-line-ness, it is that the
element **may legitimately not exist on some vantage** — not statically decidable in general, but a method
whose *name* declares "I reveal a surface" is exactly the place where "the control that reveals it is absent"
is a legitimate outcome, and a bound is the only way to report that outcome in bounded time.

Implemented as `unboundedDisclosureInteractions()` beside the loop scanner (same brace-balanced walk, exported
so the self-tests drive the real function), with a **fail-closed denominator** — the rule is decided by a
*name*, so a convention drift would empty it silently, and an empty scan is indistinguishable from a clean one
unless the count is asserted. Self-tests cover: the pre-iter-15 `openHowWeMeasure` verbatim (must be caught,
and the report must NAME the method), the fixed shape (must pass), the identical click in a NON-disclosure
method (must stay out of scope), and a one-line disclosure method (the shape the loop rule's own self-test had
to be fixed for). Self-test (b) is kept, correctly re-scoped, as the honest record of where the boundary was.

### B2 — `PT-M256-readiness-step-asserts`, discharged

`stepMethod()` re-scoped to `methodPanel()`, and the absence iter-12 measured but deliberately did not assert
is now asserted in the control (plus `methodPanel()` itself, the anchor the scope depends on).

### B3 — `pt-activity-drilldown` sharpened, both halves

- The grid and breakdown assertions now require rows that **carry text** (`contentfulContentRows` /
  `contentfulMemberResultRows`), because `contentRows().first()` visible + `count() > 0` are **both satisfied
  by 20 content-free rows**.
- The final now names the seeded hero **with her seeded role** in her own breakdown row, so it is
  tenant-specific.

## The detour that turned out to be the iter's best evidence

Mid-Phase-B the sharpened grid assertion went RED. Diagnosis, in order, each step measured:

1. `filter({hasText: /\S/})` on `main table tbody tr` → **0** while `contentRows()` → 20. Rows exist, carry
   no text: `row0 textContent = ""`, `innerText = "\t\n\t\n\t\n\t\n\t\n\t"`.
2. Watched for **40 s**: `contentful=0`, `mainLen=469`, unchanged. **Not** a hydration race, so iter-14's
   pass-2 reading was right all along and its pass-3 "it hydrates within 3 s" was a *different moment*.
3. Ran a fresh `--reset`: the seeder summary is **byte-identical** to iter-14's passing gate
   (`activity rows=600`, `feedback rows=70`, `jobsim-sessions rows=291`). The data is there.
4. Ran the **unmodified** Playthrough: it fails too — at `drillIntoActiveContent()`'s row-link wait. So the
   RED was **not** caused by this iter.
5. Read the logs. `demo-2-postgresql-1`: *"database system was not properly shut down; automatic recovery in
   progress"* at 14:38. `demo-2-jobsimulation-1`: *"DB too many ping failures, **shutting down**"* → and it
   was **gone from `docker ps`**, `Exited (0)`, together with `demo-2-cms-1`. Disk was fine (227 GiB free —
   **not** the documented ENOSPC trap). Two services had self-terminated cleanly after a Postgres blip and
   nothing restarts them.
6. `docker start` on those two containers (not a bring-up: no build, no teardown, the same class of action
   `run-playthroughs.sh --reset` already performs on the fake services). Both Playthroughs green in 7.3 s.

**Why this is the best argument for the sharpening rather than an interruption to it.** With `jobsimulation`
down, every one of its surfaces renders **20 blank rows and no error** — no empty state, no failure. The old
assertion (`rows > 0`) passes on that. The Playthrough only failed three steps later, on a row-link wait,
reporting a timeout that blames a locator. The **sharpened** assertion failed immediately and named the real
thing: *the rows carry no data*. An assertion that cannot tell "no data" from "a service is down" is the
same defect class as one that cannot fail — it just fails for the wrong reason instead of the right one.

## Phase C — the controls, each watched going RED

The drill-down control is batched onto the **same** `pt-ai-manager` login as iter-14's four (a fifth absence
for zero extra handshakes), with liveness through the same accessors — and the liveness poll is doing real
work now, since a half-dead stack renders exactly the rows it screens out. The liveness witness is the
**STARTED** hero, not the thriving one: measured, the thriving hero is *not* in that content's breakdown, and
picking the obvious hero would have shipped a control whose liveness half was simply absent.

**7 mutants, 7 RED** (harnesses in `$SCRATCH/iter15/`; `cp` backups only, all files verified byte-identical):

| mutant | result |
|---|---|
| un-bound `openHowWeMeasure` (the site that moved the boundary) | **RED** |
| un-bound `openSkillsTab` | **RED** |
| A1 `pt-activity-drilldown` driven on the contrast tenant | **RED** |
| A2 `pt-aireadiness-manager-howwemeasure` driven on a non-readiness org | **RED** |
| B1 drill-down hero absence re-aimed at the contrast tenant's own member | **RED** |
| **B2 UN-FIX the step scope (revert `stepMethod` to page-wide)** | **RED** |
| B3 method-panel absence re-aimed at a string the upsell DOES render | **RED** |

**B2 is the one worth having:** reverting the fix makes the control fire, which proves in one step that the
re-scope *was* the fix and that the control is what enforces it.

## Phase D — re-measure

**Gate: 3 consecutive cold reset-to-seed runs, `172 passed`, rc 0, 0 flake.**

| run | result | `ptreport` | wall |
|---|---|---|---|
| 1 | **172 passed**, rc 0 | 24 passing / 0 failing / 7 TODO / 0 unimplementable | 1.3 m |
| 2 | **172 passed**, rc 0 | same | 1.6 m |
| 3 | **172 passed**, rc 0 | same | 1.4 m |

- `@pt-negative-control` registry, **computed: 21 of 24** (8 self-declared + 13 via the control spec). Named
  uncovered: `pt-hiring-recruiter-compare` + the 2 studio (blocked behind `FIX-M256-studio-false-green`).
- `@pt-mutation` registry: **MUTATES=6 READ-ONLY=16 UNKNOWN=2**. `ptvalidate` **VALID**.
- Unit fences **146 passed** (the disclosure rule adds 2). `tsc --noEmit` clean; `gofmt -l` clean.
- Go: all **6** modules rc=0, **0 FAIL**.
- Python, one invocation per suite, rc captured into a variable: `demo-stack` **999 passed / 1 skipped**,
  `stack-core` **287**, `stack-verify` **171**, `stack-injection` **266 / 1 skipped** — **1723 passed, 0
  failed**.
- Drifted `demo-2` cockpit-manifest fixture backed up before the first `--reset`, restored after the last:
  sha **`99e2f315`**, verified.

### Reported (never gated) — clause 1's suite statistic, per D-v28-13

n=3, same host, cold reset-to-seed each: median per non-studio Playthrough **2.300 s = 0.6915×** of the
3.326 s baseline; per-run **0.6464× / 0.7817× / 0.7817×** (1.21× range). ORIGINAL-16 control subset
**0.5713×** — a sixth batch reading, joining 0.5281× · 1.0762× · 0.7517× · 0.9321× · 0.5562×. A gate at
0.79× still sits inside that. **iter-15 landed no speed mechanism, so clause 1's leg half has nothing new to
measure**; its flake half is **MET**.

**Environment:** `Kirality-Mac-Pro-6.local`, darwin 25.1.0, Docker VM 9.70 GiB (vs the documented 12 GB
floor); `demo-2` offset 20000, localhost/http. Per D-v28-12 no number here is comparable to billion's 228 s.

## Close — 2026-07-29

**Outcome:** **negative controls 20 → 21 of 24**, `PT-M256-readiness-step-asserts` **discharged**, and the
bounded-interaction fence **widened on its own stated condition** after its self-test's example of a *safe*
straight-line click turned out to hang for ten minutes. Along the way a half-terminated stack proved the
sharpening's worth: with `jobsimulation` down, every one of its surfaces renders 20 blank rows and no error,
which the old `rows > 0` assertion passes and the new one names.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET — clause 2 mutating **6/5 MET**, `blocked` **1/1 MET**, negative controls **21 of 24**
(`pt-hiring-recruiter-compare` routed with its investigation; 2 studio blocked behind
`FIX-M256-studio-false-green`); clause 3 verdict half **COMPLETE**, landed half short (org-admin 2/4,
onboarding 1/5); clause 1 leg half **N/A**, flake half **MET**; **D-v28-5 still unfixed** (unblocked, not
started).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik landed 2 of its 3 declared steps and moved the primary metric 20→21) — (3) re-scope: n (0 of 31 curated UCs `unimplementable`) — (4) user-blocker: n (the RED was diagnosed to a self-terminated container and recovered inside the iter with a non-destructive `docker start`; no decision was needed) — (5) cap-reached: n (2nd tik of this invocation) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D74 (**a fence whose exception boundary names a site as safe, and that site then produces the hang the boundary said would move it, is a boundary that must actually move** — widened to the DISCLOSURE class, 7 measured sites, not 28 argued ones), D75 (a probe must use the predicate the CODE uses — an exact-text count "refuted" iter-12's substring finding and was simply asking a different question), D76 (**an assertion that cannot tell "no data" from "a service is down" is the could-not-fail class wearing a different hat** — `rows > 0` passes on the 20 blank rows a half-dead stack renders), D77 (a *clean* `Exited (0)` is not a healthy container: two services self-terminated on their DB-health monitors after a Postgres blip and nothing restarted them, so the surfaces went quietly empty), D78 (pick a control's liveness witness by MEASUREMENT — the obvious thriving hero is not in that content's breakdown, and using her would have shipped a control whose liveness half was absent).
**Side-deliverables:**
- The **7 disclosure-method bounds** across `ai-readiness-page`, `assignments-page`, `org-admin-page` and
  `profile-page`. Strictly they are the fence's dependency rather than the iter's negative-control target, so
  they are recorded here — but they are the same body of work and share its commit.
**Routes carried forward:**
- `NEGCTL-M256-cross-vantage` → **21 of 24. One sharpenable Playthrough remains and it is the hard one.**
  `pt-hiring-recruiter-compare` needs a **same-vantage** control (iter-12 measured that its contrast vantage
  ejects the browser to production, `bodyLen 162` — out-of-demo, so an absence there is not evidence). What
  this iter adds is the shape of the honest option, not a solution: Org D's seed authors **5 shared positions
  / ~36 candidates** (`size: 40`, `role_mix 0.1 admin / 0.9 candidate`), so a *seed-pinned cardinality* final
  is available and would be a real strengthening — but the **absence** half is still unanswered, and the two
  candidate sources both need measuring first: (a) the hiring/workforce **framing pair** (the "Results"
  re-skin present on this vantage and the workforce "Activity" framing absent, with `pt-activity-drilldown`
  asserting the mirror image — the symmetric-pair pattern that made the iter-11 aisim pair contribute two
  controls), and (b) a candidate-vantage seat, which Org D does not currently roster. **A written verdict is
  the sanctioned outcome** if neither yields an honest absence — never a manufactured one.
- ~~`PT-M256-readiness-step-asserts`~~ → **DONE this iter.**
- `FIX-M256-demo2-service-self-termination` → **NEW.** `jobsimulation` and `cms` self-terminate (`Exited 0`)
  on their DB-health monitors after a Postgres restart, with no restart policy, and every jobsimulation
  surface then renders **20 content-free rows and no error**. A bring-up-time or verify-time guard belongs in
  `stack-verify` (the `autoverify` cheap-win class): a container-liveness assert would have named this in one
  line instead of costing an hour of Playthrough diagnosis. Not this milestone's target; **Fate 3 → M257/M258**,
  which compose the bring-up.
- `MEASURE-M256-clause1-sampling` → a sixth batch reading (0.5713×).
- Everything else from iter-14's list stands unchanged.
**Lessons:**
1. **A fence that needs a human to find its own misses is not yet a fence.** This one had an exception
   boundary, an enumerated out-of-scope set, and a stated trigger for moving — and the site that triggered it
   was quoted *inside the fence* as an example of safety. The remedy that worked was not "bound everything"
   but *find the property that distinguishes the harmful sites and measure how many there are*: 7, not 28.
2. **Diagnose the RED before disbelieving the change.** The instinct was that the new assertion was wrong.
   Five measurements later the assertion was right and the stack was half-dead — and the fastest step was the
   cheapest one (*run the unmodified version*), which should have been first, not fourth.
3. **A green suite does not prove a healthy stack, and a healthy-looking stack does not prove a green
   suite.** `docker ps` showed 14 of 16 containers "Up" with no error surfaced anywhere in the app; the two
   missing ones had exited cleanly with status 0.
