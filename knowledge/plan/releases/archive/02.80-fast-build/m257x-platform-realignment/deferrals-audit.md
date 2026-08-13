---
title: "Deferral Audit — M257x «platform re-alignment» (PRE-STAGED, not the close run)"
date: 2026-08-06
scope: milestone
invoked-by: pre-staged in parallel with the iter loop (Lane D); `close-milestone` must still run its own
release: v2.8 "fast build"
branch: m257x/platform-realignment
milestone: M257x
iters_at_audit: 101 (iter-101 IN FLIGHT at audit time — its verdicts dir was being written)
derived_at:
  rosetta: branch `m257x/platform-realignment` @ working tree, 2026-08-06
  rext (authoring): `4cb920a` (tag `fast-build-m257x-iter-101c`, on origin)
  platform: `0c91421d` == `git ls-remote origin HEAD` (level)
  app: `origin/main ad9f3c49` == remote `main` (level — see §5 change #3)
---

## Verdict

**RED — and deliberately pre-staged rather than terminal.**

> **⚠️ STATUS AT iter-102 (2026-08-06): reason 1 is DISCHARGED and reason 2 is PARTLY discharged.**
> Reason 1's item is **filed** (`D-M257x-102-1`) and §8 now has **zero** open user questions. Of reason 2's
> CHRONIC #2 block, **three of the four items are closed** — two dropped as subsumed or defect-less
> (`D-M257x-102-2`, `D-M257x-102-4`), one filed (`D-M257x-102-1`) — leaving
> `FIX-M257x-iter56-assignment-flake` alone. **The verdict stays RED on reason 3**, which is the one that
> was always the real one: the milestone is not near close. **The audit itself is not re-run here** — this
> is a status annotation by the iter lane, and `close-milestone` must still run its own audit.

RED for three independent reasons, any one of which is sufficient:

1. ~~**A HIGH-severity item has been escalated to the user for 21 iterations and is still undecided**~~
   **✅ DISCHARGED at iter-102** (`DEF-M257x-iter80-storage-prod-bucket`). It was a **blocking input to
   this audit** and was re-derived LIVE this pass: still true at platform origin HEAD — and it is now
   **filed to `platform-defect-register.md`** rather than escalated, per `D-M257x-102-1`. The audit's own
   finding §7 #1 (*"the register has zero M257x entries"*) is what resolved it: the destination existed and
   had never been used.
2. **Two CHRONIC repeat-deferral blocks** — the 11-item hardening routed-forward queue (`RF-2`, `RF-3`,
   `RF-7`…`RF-14`), restated as carried across **~33 iters and ≥5 harden passes**, and the four standing
   carry items iter-100 lists as *"carried, untouched, exactly as standing"* (restated in **28–36** iters
   each).
3. **The milestone is not near close.** `D-M257x-101-4` (written today, in flight) **withdraws the booked
   "4 of 5"** and re-grades the exit gate to **2 of 5 PROVEN** — clauses 1 and 2 were proven at `2adcf71`
   and are UNPROVEN at origin HEAD `0c91421` (6 commits / 281 changed `docker-compose.yml` lines behind).
   Clause 5's iter-101 reading returned **N = 24**, a floor.

**This is not a "something broke" RED.** Zero platform-repo edits hold; the platform clone is level with
origin HEAD; the rext pin and the per-stack consumption clone are level. It is the designed gate firing on
debt that was routed here honestly and never fated.

**What this document is NOT.** It does not fate anything unilaterally. Every Fate-1 candidate below names
its cost and its owner; **nothing was landed by this pass**, because almost all of it touches `corpus/**`,
`iter-*/`, or rext code that concurrent lanes own. `close-milestone` must still invoke
`developer-kit:audit-deferrals` for the signature; this exists so that run is a review, not a discovery.

## Summary

| | count |
|---|---|
| Distinct routed/deferred/unresolved tokens in the milestone tree (`FIX-`/`CHECK-`/`DEF-`/`RF-`/`MEASURE-`/`DOC-`/`FENCE-`/`HOST-`/`REPOINT-`/`DECIDE-`/`INVESTIGATE-`) | **255** |
| …of which **positively established OPEN** at audit time | **51** |
| …of which **positively established CLOSED / consumed / subsumed** | **64** with an explicit closure marker; the remainder are per-iter work tokens (`MEASURE-*`, `FENCE-*`, `FIX-*-blocker-set`) discharged by the named target iter — **not individually re-derived** (see §6 method limits) |
| **Blocking on a user decision** | **4** |
| **Fate-1 candidates** (landable, costed, owner named) | **19** |
| **Fate-2 candidates** (route to a named future milestone) | **17** |
| **Fate-3 candidates** (drop with a written reason) | **11** |
| **Repeat-deferrals** | **20** items across **3** chronic patterns |
| **CHRONIC patterns flagged** | **3** |
| **Status CHANGED on re-derivation** | **8** (§5) |
| **New findings made by this audit** | **5** (§7) |

---

## §1 — BLOCKING ON A USER DECISION (4)

These four cannot be fated by any lane. State them to the user as the questions in §8.

```yaml
- id: DEF-M257x-iter80-storage-prod-bucket
  severity: HIGH
  origin: iter-80 (2026-08-04), escalated to the user 2026-08-05
  status: CLOSED at iter-284 (2026-08-11) — DISCHARGED BY AN INCIDENT, not by a decision.
  re_derivation: |
    `stack-demo/platform` @ `0c91421d` (== `git ls-remote origin HEAD`), `docker-compose.yml:82`:
      - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    It sits INSIDE the `backend` service block — the service the default `core` profile starts — under a
    comment block that says these MUST be set. (:83 does the same for the public bucket.) The original
    booking derived this at `0dab54d`; the line number and the block hold at the newer ref.
  what_was_done: the false present-tense corpus claim was WITHDRAWN (iter-98), not made true.
    `safety.md`/`seeding-spec.md` now state the doc-vs-registry disagreement openly.
  what_is_still_open: NOTHING. Closed at iter-284.
  how_it_closed: |
    It was never decided; it was EXERCISED. On 2026-08-11 a demo's studio-desk attempted `s3:PutObject`
    against `s3://production-storage20240826131618541000000005/cms/<uuid>` and was refused **403** by an
    IAM policy on `user/claude-code-bedrock`. Nothing was written, and nothing in our design was the
    reason — the only thing between a demo and a production write was somebody else's permission
    boundary. The user then ruled it the top-priority item, which is the decision this entry had been
    waiting ten iters for.
    Landed at iter-284, on BOTH pointers a demo carries, because only one of them was ever in scope here:
      * the SEEDER's env — `PreflightEnv` now forces `STORAGE_S3_BUCKET` as well as the public bucket,
        on every target;
      * the RUNNING CONTAINER's env — the injected compose override strips both buckets on every emitted
        service, beside the `DIRECTUS_TOKEN` strip. This is the one the incident actually came through:
        `backend` reads its own compose env and never sees the seeder's.
    And the registry was re-classed `s3-private: PerStackIsolated -> SharedPollutionRisk`. iter-98's
    refusal ("fixing the code is only available when the change is yours to make") was right at the time
    and is now spent.
  what_this_entry_teaches: |
    The item was flagged CHRONIC_DEFER for ten iters and every restatement was accurate. What none of
    them did was make the hazard cheaper to hit or easier to see — it took a user clicking a button in a
    demo. **A correctly-recorded open hazard is still an open hazard**, and a ledger that keeps its
    entries honest is not the same as one that gets them closed.
  deferred_count: 10 iters restated it as held/undecided (80, 82, 83, 84, 85, 86, 87, 95, 96, 98, 100)
  flags: [WAS_REPEAT, WAS_CHRONIC_DEFER, CLOSED_AT_ITER_284, HIGH]

- id: FIX-M257x-iter53-union-set   # user decision D-M257x-53-5
  origin: iter-53 (2026-08-03)
  status: OPEN as booked — but SUBSUMED for scoping (see §5 change #5)
  question_as_booked: "is the clause-5 repair target 46 or 35?"
  re_derivation: |
    `D-M257x-59-1` (TOK-05) states in the milestone `decisions.md:886-892` that predicate-scoping
    SUBSUMES the question: "the union count stops being the SCOPING input and becomes a VALIDATION
    input... Whether that set is 46 or 35 changes the validation, not the work."
    So the decision no longer gates any repair; it gates only a post-sweep cross-check.
  deferred_count: 30 iters restated it
  flags: [REPEAT, CHRONIC_DEFER, BLOCKING_ON_USER, PREMISE_MOVED]

- id: CHECK-M257x-iter38-ai-act-classification
  severity: HIGH (legal exposure, not a documentation defect)
  origin: iter-38 (2026-08-02)
  status: OPEN — routed with NO OWNER inside this milestone, by explicit decision (`D-M257x-38-2`)
  substance: |
    `security_compliance.md` and `ai_architecture.md` both asserted "simulation scoring is NOT done by AI"
    and both offered it as the reason the platform classifies **Limited Risk under the EU AI Act**. The
    conjunction fails on both conjuncts (the rubric arithmetic is deterministic; the booleans it counts are
    LLM output). The premise is retracted in both docs. **Nobody has re-derived the classification.**
    D-M257x-38-2 deliberately declined to assert a replacement classification, because asserting
    "therefore High Risk" would repeat the defect with the sign flipped.
  deferred_count: 36 iters restated it
  flags: [REPEAT, CHRONIC_DEFER, BLOCKING_ON_USER, NO_OWNER, LEGAL]

- id: GATE-REGRADE-D-M257x-101-4
  origin: iter-101 (2026-08-06, IN FLIGHT)
  status: needs user ratification before it can drive scope
  substance: |
    The booked "4 of 5" is WITHDRAWN. Honest grade **2 of 5 PROVEN**: clause 3 and clause 4 re-verified at
    `0c91421`; clauses 1 and 2 UNPROVEN-at-origin-HEAD (proven at `2adcf71`, 6 commits and 281 changed
    `docker-compose.yml` lines behind); clause 5 NOT MET at N = 24.
    UNPROVEN is not REFUTED — the evidence was taken against a platform that no longer exists.
  consequence_for_this_audit: |
    `close-milestone` is not imminent. That does not reduce the value of pre-staging, but it does mean the
    close-path estimate in §9 is dominated by clause work, not by deferral work.
  flags: [BLOCKING_ON_USER, SCOPE]
```

---

## §2 — CHRONIC REPEAT-DEFERRAL PATTERNS (3)

### CHRONIC #1 — the hardening routed-forward queue (11 items)

**Booked:** pass 2, 2026-08-01, with a named source mutation each that leaves the suite green.
**Re-deferred:** explicitly restated as "unchanged" at harden passes **3, 7, 9, 12, 15, 19, 21, 22**, and
carried through **~33 iters** via the `RF-2/3/7–14` shorthand.
**Reason stable across every defer:** "source changes to bring-up scripts with real blast radius — belongs
to an iter, not a harden pass" (passes 2–3), then simply "not in this pass's scope" (9, 12, 15, 19).
**Composition + re-derivation at rext `4cb920a`:**

| id | subject | re-derived status |
|---|---|---|
| RF-2 | `demo-stack/ant-academy.sh` — the whole "SERVING is not RENDERING" block executed by **zero** positive tests; `render_body` uses `curl -fsSL … \|\| true`, so a 5xx collapses to silence | **STILL LIVE** (block at `:685-708`; `test_ant_academy.py` references are `assertNotIn` on `"started + SERVING"`) |
| RF-3 | `up-injected.sh` apply-authn FATAL path — "`authn_out`/`authn_rc` appear in no test file" | **BOOKED FALSE — see §5 change #1.** Three tests exist in `stack-injection/tests/test_apply_authn.py:527-550`, introduced by `21f2039` (2026-07-31 16:45), i.e. **before** the booking |
| RF-7 | `stack-seeding` 4 COPY targets with no offline assertion; `main_test.go`'s `mustCover` is a hand-maintained list | **STILL LIVE** (`main_test.go:229` `mustCover := []string{…}`) |
| RF-8 | `stackseed --reset` prints `reset complete` regardless of how many targets were silently skipped | **STILL LIVE** (`cmd/stackseed/main.go:801`) |
| RF-9 | `test_apply_authn.py` `skipTest`s on absent shellcheck — latent here, live on a clean CI host | **STILL LIVE** (`:610`) |
| RF-10 | the class docstring says "never a whole-file substring" and then uses three | **STILL LIVE** (`:534-542`) |
| RF-11 | `browser_graphql_endpoint` fenced by `count(...) >= 3` against 6 call sites, so two can be re-inlined | **STILL LIVE** (`test_frontend_build.py:913`) |
| RF-12 | write-target fence scope is Go-bearing sections only; 5 sections neither scored nor required to declare | **STILL LIVE** (`SCORED_SECTIONS` derived from `SECTION_COVERAGE`, Go-bearing only) |
| RF-13 | `security_compliance.md`'s `Derivation:` block cannot produce the claim it supports — it ends in a **hand** subtraction, and "that is exactly where all five failures have lived" | **corpus-owned; not re-derived by this lane** (`corpus/**` is another lane's, and a blind reading is in flight) |
| RF-14 | two byte-parallel container sweeps in `demo-stack` and `dev-stack` want one shared shell primitive | **not re-derived** (cosmetic/structural; lowest value in the queue) |

> RF-1, RF-4 landed at iter-16; RF-5, RF-6 landed at harden pass 3. Those four **leave the ledger.**

### CHRONIC #2 — the four "carried, untouched, exactly as standing" items

Named verbatim in `iter-100/decisions.md:59-64`. Restated per item in **28–36** iters:
`FIX-M257x-iter53-union-set` (30) · `FIX-M257x-iter56-assignment-flake` (28) ·
`CHECK-M257x-iter38-ai-act-classification` (36) · `DEF-M257x-iter80-storage-prod-bucket` (10, but from
iter-80 so a **higher rate**: 10 of the 21 iters since it opened).

`FIX-M257x-iter56-assignment-flake` is the only one of the four that is **not** blocked on a user and
**not** chronic-by-necessity — see §3, Fate-1.

### CHRONIC #3 — standing test debt on this host

| item | last measured | status |
|---|---|---|
| `demo-stack` full run | pass 22, 2026-08-05: **1058 tests, 6 failures** — 3 need a live container, 3 are stale live-clone baselines across two independent patch vehicles | **not re-measured this pass** (a full run risks colliding with the live iter-101 measurement) |
| `stack-core` full run | pass 22: **910 tests, 1 failure** (`test_claim_twin_guard_iter48_answer_key`) | **RE-DERIVED: still RED at `4cb920a`** — 6 tests, 1 failure, *"the fence fired on prose that no longer carries the refuted claim"*, naming `corpus/04.md:1` and `corpus/05.md:1` ← `iter-49/raw/C.md:57`. `5fb0915` ("repair a stale answer-key GREEN fixture") did **not** clear it |
| `CHECK-M257x-live-clone-suites-red` | iter-06: 7 live-clone/live-container tests RED, **reproduced on a pristine control clone** → pre-existing, not a regression | OPEN, 6 iters restated it |
| `FIX-M257x-iter100-suite-stall` | iter-100: full `python3 -m unittest discover -s tests` ran **> 55 min without terminating**, 53 lines emitted, stopped | OPEN — **this is the item that makes every "the suite is green" claim in this milestone scoped rather than whole** |

**The chronic-3 pattern is the one a close must not wave through:** a suite that cannot be run to
completion cannot be the evidence for a close. `FIX-M257x-iter100-suite-stall` is a **Fate-1 candidate**
(§3) precisely because it is cheap and it unblocks the close's own evidence.

---

## §3 — FATE-1 CANDIDATES (landable; costed; owner named)

**Prepared, not executed.** Each row says what landing costs and who owns it.

| # | item | cost | owner | why Fate-1 |
|---|---|---|---|---|
| F1 | `FIX-M257x-iter100-suite-stall` — exclude the mutation batteries from `unittest discover` and invoke them explicitly | **~1 h**, rext `stack-core`/`demo-stack` test runners; no product code | rext lane | The fix is named in the routing note ("a runner that excludes the batteries from discovery, not a timeout"). It is a precondition for any whole-suite claim at close |
| F2 | `RF-9` — `test_apply_authn.py:610` `skipTest` on absent shellcheck → fail, or assert-and-report | **~30 min** | rext lane | The commit whose message names that class shipped it. Latent here, live on a clean host |
| F3 | `RF-10` — replace the three whole-file substrings with construct-scoped assertions | **~1 h** | rext lane | The class docstring already declares the rule; the code violates its own docstring |
| F4 | `RF-11` — raise `browser_graphql_endpoint`'s fence from `>= 3` to the measured call-site count, derived not hardcoded; add the missing `build_frontend_hiring` stale-offset test | **~1–2 h** | rext lane | Same shape as the hand-maintained lists this milestone exists to end |
| F5 | `RF-8` — `stackseed --reset` must count skipped-as-absent targets and refuse (or report) rather than print `reset complete` | **~1–2 h** + a live reset to prove | rext lane | Verdict-flipping class; `main.go:801` |
| F6 | `RF-7` — derive `mustCover` from the AST instead of hand-maintaining it (iter-08 already removed this shape from the *other* fence) | **~2–3 h** | rext lane | The precedent exists in-tree; this is a port, not a design |
| F7 | Pass-22 #1 — `platform_predicate_guard` must **grade** its own `UNMEASURED`/`UNRESOLVED` reach, as iter-91 made `platform_alignment_guard` do | **~2 h** + a design call on whether partial blindness blocks the family | rext lane, **user ratifies the grading policy** | The sibling fix is written and tested; only the policy half is open |
| F8 | Pass-22 #2 — report every applied waiver on each run, and add stale-entry detection to `repair_leak_waivers.json` / `repair_reach_waivers.json` | **~2 h** | rext lane | Both files **claim in their own words** that waivers are reported; `find_leaks` (`repair_leak_guard.py:475`) silently filters them. A doc-vs-code lie in the fence family |
| F9 | Pass-22 #3 — `is_file()` guards + a typed error path in `story_org_count_guard.py:150` and the 11 unguarded `read_text()` sites in `platform_predicate_guard.py`; stop `headline()` counting traceback lines as findings | **~2 h** | rext lane | **Only 2 of the 3 named guards are live** — see §5 change #2 |
| F10 | Pass-22 #4 — a vacuity control for `demo_knob_guard` | **~1 h** | rext lane | Every sibling guard has one; this is the gap |
| F11 | Pass-22 #5 — `demopatch` `revert` must WARN on a corrupt/truncated journal (apply already does, `:430-435`); catch `OSError` in the G7 failure path so it does not escape `main`'s `except PatchError` | **~1–2 h** | rext lane | `_journal_read` (`:183`) returns `None` on `(OSError, ValueError)` — silent baseline fallback on the one vehicle that has shipped a four-release defect before |
| F12 | `FIX-M257x-iter56-assignment-flake` — `pt-assignment-assign` asserts `toBe(before - 1)` over a baseline sampled while the grid settles (observed `16 → 14`). Fix the assertion, not the count | **~2 h** + one full suite run; the routing note also says **measure first** whether app `850917d7` bears on it | playthroughs lane | Deterministic-by-construction fix; the domain link is un-measured and cheap to measure |
| F13 | `FIX-M257x-iter99-read-union` — the **28 unpaid anchors** from iter-99's reading | **1 repair iter** | the iter lane | Already enumerated with anchors; the repair-by-predicate method (`D-M257x-59-1`) is established |
| F14 | `CHECK-M257x-iter100-briefing-frozen-rext-ref` + `DEF-M257x-iter101-briefing-rext-tree` — the frozen briefing (`briefing-iter76-AS-RUN.md:37`) names the rext **authoring** copy where a rext claim settles in the **pinned per-stack clone**; it manufactured 4 of iter-99's 10 rejections | **~30 min** in the *next* reading's ground-truth addendum (the briefing itself must not be edited — comparability) | the iter lane | The remedy is already specified; it just has to be carried into the next sheet |
| F15 | `HOST-M257x-toolchain` residual — no `timeout(1)` on this box (**hit live by this audit**), and the `services/ai` green depends on `ai v1.40.1` already in `$GOMODCACHE` (Trap E unchanged cold) | **~30 min** to install `coreutils`; the cold-cache half is a documented caveat, not a fix | host owner / user | Half of it is a `brew install`; the other half should be **Fate-3 with a written reason** |
| F16 | Append M257x's platform-side findings to `knowledge/plan/platform-defect-register.md` (currently **0 M257x entries**) | **~1 h** | the lane that owns `knowledge/plan/**` | See §7 finding #1 — the register was created by the M256 deferral audit for exactly this, and M257x has not used it |
| F17 | `roadmap.md` § M257x — `**Status:** planned` while the milestone is at 101 iters; exit clause 1 still names **odysseus**, retired by `D-v28-15` | **~15 min** | the lane that owns `roadmap.md` | A stale exit clause naming a retired host is the milestone's own defect class in its own gate |
| F18 | Reconcile the **three** `fast-build-m257x-iter-101*` tags on origin (`iter-101`→`09d06070`, `iter-101b`→**the same commit**, `iter-101c`→`4cb920a`) | **~15 min** — decide the code-of-record, document the other two as superseded | rext lane | `D-M257x-101-3` explicitly warned: *"re-pin to it rather than cut a second tag, or the two tags will disagree about which tooling the release means."* Two more were cut |
| F19 | `DOC-M257x-iter22-ops-guides-5050` — 13 dead `:5050` refs in `corpus/ops/**` | **~1 h**, but blocked: the staging rows need a prior question answered | corpus lane | Mechanical once the staging question is answered |

---

## §4 — FATE-2 (route to a named milestone) and FATE-3 (drop with a reason) CANDIDATES

### Fate-2 candidates (17) — each needs a **named** destination, not "later"

> M255's close was bitten by routing four items to *"M255 harden resume"*, which was not a milestone and
> could not hold them. **A destination must be a milestone that exists.** The live candidates are
> **M257** (paused, resumes after M257x) and **M258**.

```yaml
- RF-2   → M258 (a bring-up/verify surface; M258 is the up-AND-self-proven milestone)
- RF-12  → M258 (fence scope widening; touches the same verify family)
- RF-13  → the clause-5 repair chain inside M257x, OR M258 — it is a corpus derivation block and
           the milestone that owns `security_compliance.md` prose owns it
- RF-14  → M258 (structural; lowest value)
- CHECK-M257x-iter35-seeder-writes-one-instant → **needs a destination.** Fate-3'd at harden passes
    9, 12 and 15 with the same reason each time ("a distribution change to seeded data, on a box whose
    demo-1 stack is live gate evidence"). Pass 9 called it "the highest-value item left". THREE
    same-reason defers is the pattern this gate exists to catch — it should be Fate-2'd to M258 with a
    date, or Fate-3'd once, permanently, with the believability cost written down.
- CHECK-M257x-live-clone-suites-red (7 tests) → M258
- DOC-M257x-iter33-corpus-minors  (~66)  ┐
- DOC-M257x-iter38-minors         (~45)  │  ≈ 320 corpus minors with exact anchors, accumulated across
- DOC-M257x-iter39-minors         (~60)  │  five clause-5 readings. NONE block clause 5 ("YELLOW with
- DOC-M257x-iter41-minors         (~85)  │  0 blockers" admits them) — but they have never been fated,
- DOC-M257x-iter47-minors         (~64)  ┘  and two in iter-47's set were flagged worth promoting:
    `service_taxonomy.md:150-153` (a table cell spanning four physical lines — will not render as a row)
    and `hiring.md:189-196` (the "minimal write-set" omits a NOT NULL + UNIQUE column, so a seeder built
    from it fails its INSERT).  → propose ONE destination for the whole block, not five.
- DOC-M257x-iter38-ops-collateral (4 `corpus/ops/**` docs asserting what iter-38 retracted)
- DOC-M257x-iter39-ops-collateral (the refuted "60K skills / 18K roles" survives in ~12 sites)
- DOC-M257x-iter41-ops-collateral (G6, the academy FS-fallback claim at 4 `corpus/ops/demo/**` sites)
- FIX-M257-feedback-score-approximation   → inherited from M257 iter-03; still unlanded in M257x
- DOC-M257-studio-in-app                  → inherited; corpus says studio-room is CMS-only in 5 places
- FIX-M257-stacksnap-directus-sequences   → inherited from M257 iter-02
- INVESTIGATE-M257-load1-48               → M257 on resume (peak load1 48.7 vs HEADROOM clause 1's 6);
    banked in state.md already, listed here so the audit's inventory is complete
```

### Fate-3 candidates (11) — drop, with the reason written

```yaml
- RF-3  → DROP. Re-derivation shows the booked claim was false at booking time (§5 #1). Record the
          reason: three tests cover the call site, the capture, and the exit-code measurement.
- FIX-M257x-iter101-app-clone-unfetched → DROP as discharged (§5 #3), but KEEP the generalised rule it
          produced: *"keeping a clone fetched is an ACTION, not a state, and nobody had been performing
          it"* + *"move a clone or a pin BETWEEN readings, never during one."*
- CHECK-M257x-iter100-briefing-frozen-rext-ref → folds into F14; drop as a separate id.
- FIX-M257x-iter100-read-union → superseded by iter-101's reading (in flight). Drop the id; the residual
          lives in whatever iter-101 books.
- HOST-M257x-toolchain (cold-module-cache half) → DROP with the reason: Trap E is a documented cold-start
          caveat of the box, not a tooling defect; the `timeout(1)` half is F15.
- FIX-M257x-iter53-union-set → **cannot be dropped by a lane** (§1), but the recommendation to the user is
          DROP-AS-SUBSUMED: TOK-05 turned 46-vs-35 from a scoping input into a validation cross-check.
- CHECK-M257x-iter86-value-change-weak-form → `value_change_guard` built the weak form of its own
          proposition (cross-reference verb `see` → `read`). Low value; drop or fold into F7's family sweep.
- FIX-M257x-iter100-semantic-anchor-class → `D-M257x-100-4` already answered this with a plain **no**:
          intra-document self-contradiction is NOT fenceable (quadratic in claim count; the relation is
          semantic entailment). Drop the id and keep the decision.
- RF-14 → alternatively DROP (structural tidiness with no defect behind it).
- DOC-M257x-iter31-role-concentration-believability, DOC-M257x-iter30-job-role-title-unfilled → drop as
          believability notes with no defect and no consumer.
- origin/pr-14 → standing verdict **DO NOT MERGE**, restated at TOK-05. Record it as a permanent
          disposition so it stops being re-stated every tok.
```

---

## §5 — ITEMS WHOSE STATUS CHANGED WHEN RE-DERIVED (8)

This milestone's defining defect class is *a check that reports a state without measuring it* — found ~20
times, including inside its own instruments. Six routed counts collapsed on derivation (64→5, 23→1, 21→0,
92→0, 4→3, 145→3) and two grew (152→140, 8→38). Every item below was re-derived against the tree, not
inherited from the note that routed it.

**1. `RF-3` — BOOKED FALSE, and it was false when booked.**
Booked at harden pass 2 (2026-08-01): *"iter-04's FATAL apply-authn path has no test; `authn_out`/`authn_rc`
appear in no test file."* Measured at `4cb920a`: `stack-injection/tests/test_apply_authn.py` contains
`authn_out`/`authn_rc` at `:536`, `:540`, `:548`, `:549`, inside three tests
(`test_the_call_site_does_not_discard_the_appliers_output`,
`test_the_call_site_captures_the_output_and_reports_it`,
`test_the_reported_exit_code_is_measured_not_assumed`). `git log` on that file shows exactly **two**
commits ever — `2813a5e` (2026-06-04) and **`21f2039` (2026-07-31 16:45)**, which is iter-04's own commit.
Pass 2 ran the next day. **The tests existed at booking time.** The likeliest mechanism is §5 rule 44: the
scan was section-scoped to `demo-stack`, and the tests live in `stack-injection`.
→ Fate-3, DROP with the reason recorded. **This is one of eleven items the queue has carried for 33 iters.**

**2. Pass-22 item #3 — 2 of 3, not 3 of 3.**
Booked: *"`story_org_count_guard`, `platform_predicate_guard` and `unreadable_repo_claim_guard` each
`read_text()` an input with no `is_file()` check."* Measured **at the pass's own commit `6130bfd`**,
`unreadable_repo_claim_guard.py` already had `is_file()` guards at `:106`, `:163`, `:164`, `:168`, covering
both of its `read_text()` sites. Live at `4cb920a`: `story_org_count_guard.py:150` unguarded (0 `is_file()`
in the file); `platform_predicate_guard.py` has **11 unguarded `read_text()` sites** of 18.
→ The item is real but **narrower on one axis and wider on another**. F9's scope should be re-cut.

**3. `FIX-M257x-iter101-app-clone-unfetched` — DISCHARGED since it was written (today).**
`D-M257x-101-1` recorded `stack-demo/app` local `origin/main` at `2035f9a4` against real remote
`ad9f3c49` — STALE, with the consequence that *"the citation guards are currently grading `app` anchors at
`2035f9a4`, which is itself no longer origin/main."* Measured now: local `origin/main` = **`ad9f3c49`** =
remote `main`. The fetch has happened.
→ Fate-3 as discharged; keep the rule.

**4. The `D-M257x-101-3` handoff — both "NOT DONE" steps are now DONE.**
Recorded mid-iter as owned by another lane: `.agentspace/rext.tag` still `iter-67`, `stack-demo/rosetta-extensions`
not level. Measured now: `rext.tag` = `fast-build-m257x-iter-101c`; the per-stack clone is at `4cb920a`,
which is exactly `fast-build-m257x-iter-101c` on origin.
→ Discharged. **But see §7 finding #3** — it was done by cutting two more tags, which the same decision
warned against.

**5. `FIX-M257x-iter53-union-set` — the premise moved under it.**
Booked as *"repair the 46 (or 35); the target count is a user decision."* `D-M257x-59-1` (milestone
`decisions.md:886-892`) states that predicate-scoping **subsumes** the question: the union count is now a
validation input, not a scoping input. Nothing in the intervening 47 iters withdrew that, and iters 76, 82,
86 and 100 all restate the item verbatim as *"PENDING USER DECISION"* without noting the subsumption.
→ The user's question should be re-put in its **current** form (§8 Q2), not its iter-53 form.

**6. The exit gate — 4 of 5 → 2 of 5 PROVEN.**
`D-M257x-101-4`, written today, withdraws a grade the milestone carried since iter-37. Clauses 1 and 2 were
proven at `2adcf71`; origin HEAD is `0c91421`, 6 commits and 281 changed `docker-compose.yml` lines later,
with 3 compose services deleted, 2 `repos.yml` entries removed and the default profile renamed
`graphql` → `core`. **The mechanical proof it cannot be carried:** platform `838d907` moved the
`$HOME/.aws/credentials` bind off the deleted `jobsimulation` service onto `backend` (`docker-compose.yml:100`
at `0c91421`); the demo override's mitigation was keyed on the literal `"jobsimulation"` and silently went
dead, **and its tripwire test skipped, which reads exactly like a pass**. The fix (`7844e97`) existed only
in the authoring copy until iter-101's tag.
→ Not a deferral, but it is the single largest change to the close-path estimate (§9).

**7. `DEF-M257x-iter80-storage-prod-bucket` — re-derived at the NEWER ref and still true.**
The original derivation was at `0dab54d`. Re-derived at `0c91421` (== origin HEAD): `docker-compose.yml:82`
still sets `STORAGE_S3_BUCKET` to the production private bucket inside the `backend` block. The line number
and the block both hold across the platform advance.
→ Status unchanged, but the evidence is now current, which is what the escalation needs.

**8. `stack-core`'s standing red — still red, after a commit that aimed at it.**
Pass 22 recorded `stack-core` at 910 tests / 1 failure (`test_claim_twin_guard_iter48_answer_key`). Commit
`5fb0915` is titled *"repair a stale answer-key GREEN fixture."* Run at `4cb920a`: `Ran 6 tests … FAILED
(failures=1)`, *"the fence fired on prose that no longer carries the refuted claim"*, naming `corpus/04.md:1`
and `corpus/05.md:1` ← `iter-49/raw/C.md:57`.
→ The standing red is unchanged in **count** and changed in **content**. It is a perishable answer-key
fixture against live corpus state, so it will keep moving while `corpus/**` moves.

---

## §6 — METHOD, AND WHAT IT DOES NOT ESTABLISH

**Instruments used (three, per §5 rule 44 — no single search tool is safe here):**
1. `/usr/bin/grep -r` over `knowledge/plan/**` (NOT the shell's `grep`, which is `ugrep --ignore-files`
   and hides tracked-but-gitignored files).
2. A Python walker over every `*.md` under `knowledge/plan/`, tokenising
   `(FIX|CHECK|DEF|RF|INVESTIGATE|MEASURE|DECIDE|DOC|REPOINT|FENCE|HOST|DROP|WAIVE)-…` with backticks
   stripped before matching (the first cut missed every `~~\`TOKEN\`~~` strikethrough because the
   backticks sat inside the strikethrough).
3. `git grep` / `git log -S` / `git show <ref>:<path>` in the rext authoring copy and the platform and
   `app` clones, for every code-side re-derivation.

**What this audit MEASURED:** the 4 blocking items, all 11 RF items (9 of them at `file:line` in the rext
tree), all 5 pass-22 items, the platform-side storage claim at origin HEAD, the `app` clone freshness, the
rext pin/clone/tag state, and one live test run.

**What it did NOT measure, stated rather than smoothed:**
- **The per-iter work tokens** (`MEASURE-M257x-*`, `FENCE-M257x-*`, `FIX-M257x-iterNN-blocker-set`,
  `-union-set` for iters other than 53, `DOC-M257x-iterNN-*-residual`) were classified as consumed **by
  structure** — each names a target iter that exists and closed — **not individually re-derived**. On this
  milestone's own evidence that is exactly the reasoning that produces a false green. If the close wants a
  number it can defend, ~40 of those need one grep each.
- **The full `demo-stack` and `stack-core` suites** were not re-run. A full run collides with the live
  iter-101 measurement, and `FIX-M257x-iter100-suite-stall` means `discover` does not terminate anyway.
- **`corpus/**` was not read or written.** A blind reading is in flight and any corpus write moves its
  subject. RF-13 and the ~320 corpus minors are therefore inventoried from the plan tree only.
- **`origin/pr-14`** was not fetched or read; the DO-NOT-MERGE verdict is carried, not re-derived.

---

## §7 — FINDINGS MADE BY THIS AUDIT (5)

1. **`platform-defect-register.md` has ZERO M257x entries.** It holds 4, all from M256. It was created
   *by the M256 deferral audit* because *"a defect recorded inside a closed milestone has been filed where
   it cannot be found."* M257x has found at least one platform-side defect that belongs in it —
   `DEF-M257x-iter80-storage-prod-bucket` — and it is filed only in `iter-80/progress.md` and this
   milestone's `decisions.md`, both of which flip to archived at close. **The register was built for this
   exact item and has not been used.** → F16.
2. **`roadmap.md` § M257x is stale in two load-bearing ways.** `**Status:** planned` at 101 iters, and
   **exit clause 1 still names `odysseus`** — retired from the project by `D-v28-15` (2026-07-31), which
   `state.md` records but the roadmap does not. A gate clause that names a retired host cannot be met as
   written; that is this milestone's own defect class, in this milestone's own gate. → F17.
3. ~~**Three `fast-build-m257x-iter-101*` tags exist on origin**, two of them pointing at the *same* commit
   (`iter-101` and `iter-101b` both dereference to `09d06070`; `iter-101c` → `4cb920a`).~~ → F18.

   > **⚠️ STRUCK — FALSE POSITIVE. Re-derived against origin at iter-102 (`D-M257x-102-6`): there is
   > EXACTLY ONE.**
   >
   > ```
   > git ls-remote --tags origin 'refs/tags/fast-build-m257x*' | grep -v '\^{}' | wc -l  ->  54
   > … | grep 'iter-101'                                                                 ->   1
   > fast-build-m257x-iter-101  ->  tag object 0011c10a, peeled commit 09d06070
   > ```
   >
   > No `iter-101b`, no `iter-101c`; **no origin tag points at `4cb920a`** at all. `.agentspace/rext.tag`
   > reads `fast-build-m257x-iter-101`, the clone HEAD is `09d0607`, and `git tag --contains 7844e97`
   > includes it. **The other lane's re-pin is done and correct**; its commit `b02150c` is titled, in terms,
   > *"one rext tag, not two"*.
   >
   > **Two mechanisms, and the evidence does not separate them — so both are recorded.** (1) A **peeled-ref
   > miscount**, which is demonstrable: `git ls-remote --tags` prints **two lines per annotated tag** (the
   > tag object and its `^{}` peel), so counting *line shapes* rather than *distinct refs* doubles every
   > annotated tag — for `iter-101` alone, **2 raw lines, 1 distinct ref**. This is the same mechanism run 57
   > fixed in the guard family (*"take the cardinality from the GUARD, not from line shapes"*), recurring in
   > a different instrument. (2) A **transient state since reconciled** by `b02150c`. Mechanism (1) cannot
   > manufacture the specific names `iter-101b`/`iter-101c`, so it is **not a complete explanation on its
   > own** — which is why (2) stands beside it rather than being dropped for the tidier story.
   >
   > **Booked regardless:** a peeled-ref line-shape miscount is a real `--tags` reading defect and will
   > recur. Any future tag count takes `| grep -v '\^{}'` first, or counts `refs/tags/` names.
4. **`repair_leak_waivers.json` and `repair_leak_guard.py` contradict each other in writing.** The guard's
   own prose says *"It can only ever make the fence quieter, so it is reported"* and the waiver file says
   *"every one is reported on each run"*; `find_leaks` (`repair_leak_guard.py:475`) returns
   `[lk for lk in leaks if not is_waived(lk, waivers)]` with no reporting path at all. Pass 22 booked this;
   this audit confirms it at `4cb920a` and adds that **the lie is in the artifact's own documentation**,
   which is the class the milestone is named after. → F8.
5. **`timeout(1)` is absent on this host** — hit live by this audit while trying to bound a test run.
   `HOST-M257x-toolchain` named it at iter-04 and it has never been installed. Any routed instruction that
   assumes `timeout` will silently not run. → F15.

---

## §8 — WHAT THE USER MUST DECIDE BEFORE `close-milestone` CAN RUN

> ## ⚠️ SUPERSEDED AT iter-121 — this banner said **"ZERO open user questions remain"** and it was
> ## measured wrong. **Q5 is open.** See [§12](#12--the-blocking-state-sweep-derived-over-every-field-that-can-block) for the derivation and the full enumeration.
>
> The banner below was **true of Q1–Q4 and false of the milestone**, because the sweep behind it read
> **one** grading field (`user-blocker`) and iter-119 graded its outcome **`re-scope: y` / `user-blocker:
> n`**. §11 spotted the single instance in prose; §12 replaces the prose with a **mechanized, multi-field
> derivation** (`rosetta-extensions/stack-core/blocking_state_guard.py`) and finds **8** blocking gradings
> across the milestone, of which **5** this file had never named.
>
> ## ✅ RESOLVED AT iter-102 — **Q1–Q4 are closed.**
>
> All four were decided or withdrawn on 2026-08-06. This section is kept in full, with each question's
> disposition inline, because the *reasoning* is the durable part — three of the four turned out to be
> questions that should never have reached a user.
>
> | | question | disposition | record |
> |---|---|---|---|
> | **Q1** | `DEF-M257x-iter80-storage-prod-bucket` | **(b) FILED** to `platform-defect-register.md`; the item moves *escalated-undecided → filed* | `D-M257x-102-1` |
> | **Q2** | `FIX-M257x-iter53-union-set`, 46-vs-35 | **DROPPED as subsumed** (Fate 3) — `D-M257x-59-1` absorbed it **47 iters ago**; the thing it was pending on no longer exists | `D-M257x-102-2` |
> | **Q3** | `CHECK-M257x-iter38-ai-act-classification` | **WITHDRAWN — not a user question at all.** The repair finished at iter-38; the corpus asserts no classification. What was carried was an *aspiration*, not a defect | `D-M257x-102-4` |
> | **Q4** | ratify the gate re-grade | **RATIFIED — 2 of 5 PROVEN**, and clauses 1–2 are a **CLOSE BLOCKER**, not an M258 route | `D-M257x-102-3` |
>
> **A standing rule now governs this section** (`D-M257x-102-5`, binding user decision): **no legal,
> regulatory, compliance or policy question is escalated during delivery.** Route it — close it, file it, or
> state exactly what it blocks. **Never "needs an owner."** That routing is what turned Q3, a finished
> repair, into a 36-iteration standing question.
>
> **The pattern across Q1–Q3 is the audit's most useful output and outranks its defect list:** two of the
> three were **not decidable questions at all** — one was already answerable by the register that existed
> for it, one had been subsumed 47 iters earlier, and one had no defect behind it. **Each survived because
> every pass restated it instead of re-reading it at source.** See `platform-alignment.md` §5 rules 47–48.

Four questions **as originally put**. **Q1 is a blocking input to this audit and has been open since 2026-08-05.**

> **Q1 — `DEF-M257x-iter80-storage-prod-bucket`: what is the disposition?**
> At platform origin HEAD `0c91421`, `docker-compose.yml:82` sets `STORAGE_S3_BUCKET` to the **production**
> private bucket inside the **`backend`** block — the service the default `core` profile starts — so local
> private writes on a default dev stack land in a production bucket. The corpus claim that said otherwise
> was **withdrawn**, not made true. `stack-seeding/isolation/isolation.go:106` still registers `s3-private`
> as `PerStackIsolated`, and re-classing it *is* the disposition.
> **Choose one:** (a) re-class `s3-private` in the isolation registry (a rext code change — ours to make);
> (b) file it to `platform-defect-register.md` as a platform defect and leave the registry as-is with the
> disagreement documented (the current state, made permanent); (c) something else.
> *(a) and (b) are not exclusive.*

> **Q2 — `FIX-M257x-iter53-union-set` / `D-M257x-53-5`, re-put in its current form.**
> The question as booked was *"is the clause-5 repair target 46 or 35?"* Since TOK-05 (`D-M257x-59-1`),
> predicate-scoping **subsumes** it: the union count is now a **validation** cross-check after the
> predicate sweep, not the scoping input. **Do you still want a ruling on 46-vs-35, or is the item
> dropped-as-subsumed?** If a ruling is still wanted, it changes which set the post-sweep validation must
> cover — nothing else.

> **Q3 — `CHECK-M257x-iter38-ai-act-classification`: who owns it, and when?**
> Both docs now state that the per-check verdicts are LLM-produced and that the stated basis for the
> **EU AI Act Limited-Risk** classification does not hold at platform HEAD. Neither asserts a replacement
> classification, by design (`D-M257x-38-2`). This is a **legal** question with **no owner inside this
> milestone**, carried across 36 iters. **Name an owner and a date, or record a written decision that the
> corpus will carry the retraction indefinitely without a classification.**

> **Q4 — ratify or reject the gate re-grade `D-M257x-101-4` (4 of 5 → 2 of 5 PROVEN), and set the scope
> that follows.**
> Clauses 1 and 2 are UNPROVEN-at-origin-HEAD (not refuted). Re-proving them means a cold
> `--purge` + `demo-up` × 3 plus a full Playthrough suite at `0c91421` with tooling that contains `7844e97`.
> **Is that in M257x's scope, or does M257x close at "2 of 5 proven + the map + the fence" with clauses 1/2
> re-proof routed to M258 (which needs a proven bring-up anyway)?**
> A concurrent lane already owns the re-run; this question decides whether it is a *close blocker*.

> **Q5 — the SCOPE DECISION, opened at iter-119 and OPEN. Added at iter-121; §8 could not see it.**
> Both `TOK-07` (repair-and-read) and `TOK-08` (the user's own enumerate-then-read re-scope) have been
> **refuted by their own pre-registered arithmetic** — `P = 37` vs `P ≥ 15` at iter-116, `P = 22` vs
> `P ≥ 19` at iter-119 — and `TOK-08`'s sealed rule **bars a successor strategy**, so there is no TOK-09.
> Clause 5 is met only by a reading that returns zero; the floor is **≥ 46** and a zero reading is not
> near. **What M257x closes as, and what clause 5 is re-scoped to (if anything), is the user's call.**
> `state.md` `phase:` has read **AWAITING USER SCOPE DECISION** since iter-119.
> **It is graded `re-scope: y`, `user-blocker: n`** — which is exactly why a `user-blocker`-keyed sweep
> reported zero. The recommendation is the milestone's; the decision is the user's, and this row does not
> pre-empt it.

**Also for signature, but not blocking (they can ride Q1–Q4):** a single destination for the ~320 corpus
minors (§4), and the Fate-3 list in §4 as a block.

---

## §9 — ESTIMATED REMAINING CLOSE-PATH WORK

Itemised. **The deferral gate itself is now ~0.5 day of the critical path instead of 0.5–1 day + an
unbounded user wait — provided Q1–Q4 are answered before `close-milestone` starts.**

| # | item | estimate | on the critical path? |
|---|---|---|---|
| 1 | User answers Q1–Q4 | unbounded — **start now** | **YES** (Q1 has already waited 21 iters) |
| 2 | Clause 5 to `N = 0`: at N = 24 (iter-101, a floor), with per-pass recall 43–51 % and union 62–78 % | **≥ 3–5 iters** (repair → read → repair → read), and the series has never converged monotonically | **YES** |
| 3 | `D-M257x-101-2` planted-defect positive control, payable at the moment of the first zero | **1 iter to author + 1 cycle** | **YES** (a user decision, already recorded) |
| 4 | Clauses 1 + 2 re-proof at `0c91421` with tooling ≥ `7844e97` — 3 cold `--purge`+`demo-up` cycles + a full Playthrough suite | **~1 day** (3 × ~11 min bring-up + suite + triage) — **owned by a concurrent lane** | **YES, unless Q4 routes it to M258** |
| 5 | `FIX-M257x-iter100-suite-stall` (F1) — without it no whole-suite claim at close is defensible | **~1 h** | **YES** |
| 6 | Fate-1 batch F2–F11 (RF-9/10/11/8/7 + the five pass-22 items) | **~1.5–2 days** for all ten; **~4 h** for the six cheapest (F2, F3, F10, F11, F9-narrowed, F14) | NO — can run in parallel |
| 7 | Fate-1 F12 (`assignment-flake`) + one confirming full Playthrough run | **~3 h** | NO — unless Q4 puts clause 2 in scope, then YES |
| 8 | Fate-1 F13 (`FIX-M257x-iter99-read-union`, 28 anchors) | **1 repair iter** | folded into #2 |
| 9 | Fate-1 F16–F18 (register entries · roadmap status/host · tag reconciliation) | **~1.5 h total** | NO |
| 10 | Fate-2 routing decisions (17 items) + Fate-3 write-ups (11 items) | **~2 h**, once Q1–Q4 land | **YES** (it is the gate's own output) |
| 11 | Final harden pass (`harden-mstone-iters --final`) | **~0.5–1 day** | **YES** |
| 12 | `close-milestone` itself | **~0.5 day** | **YES** |

**Critical-path total, assuming Q4 keeps clauses 1/2 in M257x:** roughly **4–7 days**, dominated by
clause 5's convergence (#2), which has no reliable estimate — the reading series is
`25 → 13 → 11 → 17 → 37 → 18 → 7 → 12 → 14 → 7 → 13 → 20 → 28 → 24` and *"every better instrument found
more."* **If Q4 routes clauses 1/2 to M258, subtract ~1 day.**

---

## §10 — WHAT NEEDS ANOTHER LANE

Nothing in this file was executed. These are handoffs, each with its owner.

| to | what |
|---|---|
| **the iter lane** (owns `iter-*/`, `progress.md`, `decisions.md`, `state.md`, `hardening-ledger.md`) | Book this audit in `progress.md`'s routes table so it is discoverable at close. Carry F13/F14 into the next reading's ground-truth addendum. Note in `decisions.md` that `D-M257x-101-1`'s app-clone precondition and `D-M257x-101-3`'s two handoff steps are now **discharged** (§5 #3, #4) |
| **the corpus lane** (owns `corpus/**`) | RF-13 (`security_compliance.md`'s `Derivation:` block ends in a hand subtraction); the ~320 minors; the three `*-ops-collateral` sets; `DOC-M257x-iter22-ops-guides-5050`. **Do not start until the blind reading in flight has closed** |
| **the rext lane** (owns rext code + tags + `.agentspace/rext.tag`) | F1–F11. ~~**F18 is the urgent one**~~ — **F18 is STRUCK as a false positive** (§7 #3, `D-M257x-102-6`): re-derived at iter-102, there is **exactly one** `iter-101` tag on origin, the pin matches it, and it contains `7844e97`. **Nothing is urgent for this lane.** **F1 (`FIX-M257x-iter100-suite-stall`) is the highest-value item** — until it lands, no whole-suite claim in this milestone is defensible, which is the close's own evidence |
| ~~**the lane that owns `knowledge/plan/**`**~~ **DONE at iter-102** | ✅ F16 — `platform-defect-register.md` has its first M257x entry (`PLATFORM-M257x-compose-points-local-backend-at-the-PRODUCTION-S3-buckets`). ✅ F17 — `roadmap.md` § M257x status and the retired-`odysseus` clause both corrected |
| **the user** | ✅ Q1–Q4 all decided or withdrawn at iter-102 — see the resolution block at the head of §8. ✅ **Q5 CLOSED at iter-283 by `TOK-09`** (the iter-119 `re-scope`; [§12](#12--the-blocking-state-sweep-derived-over-every-field-that-can-block), corrected at [§13](#13--close-run-re-audit-2026-08-11)). **Zero open user questions at close.** **⚠️ This row read *"DONE at iter-102 — zero open user questions"* until M257x iter-281.** §8's banner was corrected at iter-121 and this row was not, so the file asserted the claim in one section and retracted it in another — repair the site, leave the class, the same shape this milestone keeps paying for. It is not a harmless duplicate: §10 is the hand-off table a close gate reads to decide whether a lane is finished |

---

_Pre-staged 2026-08-06 by Lane D, in parallel with the M257x iter loop. Derived at rosetta branch
`m257x/platform-realignment`, rext `4cb920a`, platform `0c91421d`, app `ad9f3c49`. Zero platform-repo
edits; zero writes outside this file._

---

## §11 — RE-DERIVATION AT iter-120 (harden pass 26). What changed since this audit was written.

**This audit is staler than its own header implies.** `iters_at_audit: 101`, and it was last touched at
`cd16967` (iter-102). The milestone is now at **iter-119 + harden pass 26**. It is **17–18 iters stale**,
and it predates harden passes 23, 24 and 25 entirely. Every "cannot know about" below is wider than a
9-iter gap would suggest.

**Instruments (three, per §5 rule 44), each stated with its count.** Re-derived at corpus `f723101`:

```
I1  /usr/bin/grep -rnoE '(FIX|CHECK|DEF|RF|INVESTIGATE|MEASURE|DECIDE|DOC|REPOINT|FENCE|HOST|DROP|WAIVE)-[A-Za-z0-9._-]+' knowledge/plan/
      -> 4109 raw occurrences · 525 distinct (trailing `._-` normalised) · M257x tree: 290 distinct
I2  python3 os.walk over knowledge/plan/**/*.md, BACKTICKS STRIPPED per line before matching
      -> 2150 files · 240693 lines · 3839 occurrences · 522 distinct · M257x tree: 290 distinct
I3  git grep -ohE '<same>' HEAD -- 'knowledge/plan/*'          -> 525 distinct · M257x tree: 290
    git -C .agentspace/rosetta-extensions grep … @ 4304930     -> 36 distinct M257x tokens
```

**All three agree: 290 distinct tokens in the M257x tree** (FIX 105 · CHECK 84 · DOC 30 · FENCE 23 ·
RF 16 · DEF 12 · MEASURE 12 · HOST 3 · REPOINT 2 · DECIDE 1 · INVESTIGATE 1 · DROP 1).

**The audit's own Summary claims 255; re-derived at its own commit it is 259–260.** The gap is §6 rounding
plus **6 pseudo-ids the audit coined that were never booked anywhere** (`MEASURE-M257x`, `FENCE-M257x`,
`DOC-M257x-iterNN`, `FIX-M257x-iterNN-blocker-set` — glob placeholders; `DROP-AS-SUBSUMED`; and
`INVESTIGATE-M257-load1-48`, real but banked only in `state.md`).

### What the marker counts do and do not say

`STRICKEN 8 · OPEN 119 · CLOSED 29 · UNMARKED 134`. **UNMARKED is 46 %, and the tree is append-only** —
`git ls-tree` set-diff between the audit commit and HEAD returns **0 removed tokens**. Closure is recorded
in prose beside a token, never by deletion, so **absence proves nothing and this split is a MARKER count,
not a STATE count.** Exactly the limit §6 declared; restated because it is easy to read as a state count.

### Instrument blind spots, measured rather than assumed

| instrument | blind spot | measured |
|---|---|---|
| I1 | tracked-but-gitignored | **none exist** in this tree (0 lines from `git check-ignore` over `git ls-files`) |
| I1 | NUL-bearing files | **none.** 8 `.md` files read as binary; all 8 are **0-byte empty** |
| I2 | backtick-stripping **manufactures** tokens | **1 false positive** — `` `DEF-M239-01`-sibling `` welds into `DEF-M239-01-sibling`. The strikethrough fix has its own defect |
| I2 | trailing sentence punctuation | **114 phantom tokens** (636 → 522). `.` is inside the class, so `` `TOKEN`. `` yields `TOKEN.` — any un-normalised count is inflated ~22 % |
| I1+I2 | rext-resident tokens | **9 tokens exist ONLY in rext code** and are invisible to any `knowledge/plan` sweep (`FENCE-M257x-iter44/-iter77/-iter86/-iter93/-iter105/-iter106/-iter107/-iter117/-iter118`) |

### NEW since the audit — 30 tokens, 0 removed

Two are structural rather than routine:

- **`DEF-1`…`DEF-4` are a NEW, COLLIDING numbering scheme** — real defect ids in `gate-clauses-1-2/README.md`
  (Lane B, opened `ae192dd`), unqualified integers in the same namespace as `RF-1`…`RF-14`. Any future
  audit grepping `DEF-` gets four unscoped hits it cannot attribute to a milestone.
- **`FIX-M257x-harden23-json-polluted-by-provenance-stamp`** — routed, not fixed, at pass 25: 12 guards'
  `--json` is unparseable because `stamp()` writes to stdout, and *the suite does not see it because the
  suite works around it* (`FENCE_PROVENANCE_STAMPED=1`).

### Audit items that CLOSED

`DEF-M257x-iter80-storage-prod-bucket` **FILED** (`D-M257x-102-1`; now in `platform-defect-register.md`) ·
`FIX-M257x-iter53-union-set` **DROPPED as subsumed** (`D-M257x-102-2`) ·
`CHECK-M257x-iter38-ai-act-classification` **DROPPED** (`D-M257x-102-4`) ·
`GATE-REGRADE-D-M257x-101-4` ratified then **SUPERSEDED** — the gate is **4 of 5**, not 2 of 5, so §1, §5 #6
and the whole of §9 rest on a grade that no longer holds · `FIX-M257x-iter56-assignment-flake` hold retired,
repair routed · `FIX-M257x-iter100-suite-stall` **superseded by recurrence** as
`FIX-M257x-iter108-stackcore-suite-hangs`, still open.

### Audit items STILL OPEN — exactly one, and it is the strongest signal here

```
DEF-M257x-iter101-briefing-rext-tree  ->  restated in ALL TEN of iters 110-119
```

**The audit priced it at "~30 min"** as the cheap half of F14. Seventeen iters later it is still
*"open, delivered-unfixed"*. A 30-minute item that has outlived the audit that priced it is a pattern
signal in its own right.

### CHRONIC #1 went SILENT, which is not the same as closed

`RF-2` appears in **37** iter dirs, last at iter-107, and **not once in 110–119**. `RF-3` last iter-104;
`RF-7`/`RF-12` last iter-42; `RF-13` last iter-54; `RF-6/8/9/10/11/14` **never appear in any iter dir**
(ledger-only). The hardening ledger's last explicit restatement of the RF queue is **Pass 19**; passes
20–25 do not restate it at all. **F2–F6 (RF-9/10/11/8/7) have no evidence of landing and no evidence of
being carried.** Six further tokens at **18–24 restatements each** appear nowhere in this audit and went
quiet around iter-75/76 with no recorded fate — under the three-fate rule those are **unfated drops**.

### A SECOND chronic block, built entirely after this audit

| token | iter dirs | in 110–119 |
|---|---|---|
| `DEF-M257x-iter101-briefing-rext-tree` | 17 | **10/10** |
| `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` | 13 | **10/10** — **explicitly *de-ranked* at iter-110, then re-deferred nine more times** |
| `FIX-M257x-iter111-staged-battery-dependency-is-underived` | 9 | 9/10 |
| `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag` | 8 | 8/10 |
| `FIX-M257x-iter113-adjudication-is-judgement` | 7 | 7/10 |
| `FIX-M257x-iter108-stackcore-suite-hangs` | 7 | 5/10 |
| `FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live` | 5 | 5/10 |

**De-ranked rather than fated, then carried in every subsequent iter, is the shape the three-fate rule
exists to prevent.**

### §8 — CONFIRMED narrowly, REFUTED as a statement about the milestone

**Q1–Q4 have NOT reopened.** All ten of iters 110–119 grade `user-blocker: n`, and `D-M257x-102-5`'s
no-escalation rule is being honoured. **So §8's own four questions remain at zero.**

**But §8 is no longer a true statement about what the user must decide.** A larger question opened at
iter-119 and §8 has no row for it: `TOK-08` OUTCOME routes the milestone to **a user scope decision**, and
`state.md` reads **AWAITING USER SCOPE DECISION**. Note it is graded **`re-scope: y`**, not `user-blocker`
— **a sweep keyed on the `user-blocker` field alone reports zero and is wrong.** §9 row 1 (*"User answers
Q1–Q4 — start now"*) is likewise superseded.

### Stale-again inside the audit itself

- **§5 #8** — the standing `test_claim_twin_guard_iter48_answer_key` red: pass 25 states it was **not
  re-attested**. The audit's *"RE-DERIVED: still RED at `4cb920a`"* is 17 iters old.
- **F17 RE-ROTTED.** Landed at iter-102, marked DONE in §10. `roadmap.md:664-667` today still reads
  *"102 iters + 22 harden passes"* and *"Gate 2 of 5 PROVEN"*; truth is **119 iters + 26 harden passes,
  gate 4 of 5**. A Fate-1 item that was landed once and has already decayed again.

### What iter-120 adds that no ledger sweep could find

Three defects closed this pass were **booked and carried**, and two were **never booked at all**:

- carried: `FIX-M257x-iter119-clerk-signin-token-claim-understates-surface` (closed — **and the
  enumeration found 5 sites where both readings booked 3**), `FIX-M257x-iter116-intra-corpus-miscitation`
  (construct half closed, all 8 + a 9th), `FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block`
  (G10 half closed; the general half stays open).
- **never booked:** the *"Sentinel validates **every** API request"* understatement at 4 sites, and
  FENCE-M257x-iter117's **uncollectable** test module. Neither was a deferral, so **no deferral audit
  could ever have surfaced them.** That is the boundary of what this instrument measures.

**New token booked here:** `FIX-M257x-iter120-anchor-guard-detects-blank-not-wrong`.

---

## §12 — THE BLOCKING-STATE SWEEP, derived over EVERY field that can block

**§8's zero was a MEASUREMENT ERROR, not a stale note.** §11 caught the single instance in prose —
*"a sweep keyed on the `user-blocker` field alone reports zero and is wrong."* That sentence was right and
it stopped one layer short: it fixed the *reading* of one iter and left the *derivation* one-field-wide.
So §12 replaces the prose with an instrument.

**The instrument** (`rosetta-extensions/stack-core/blocking_state_guard.py`, FENCE-M257x-iter121). For the
milestone named in `state.md`'s `active_milestone:`, it parses the **Phase-5 grading** of every
`iter-*/progress.md` and asserts that every grading which routes *out of the iter loop to a decision the
loop cannot take* is **named in this file**. Three fields do that, and the list is checked in **both
directions** — a field no grading uses is `exit 2`, so it cannot hold a name that can never fire:

```
BLOCKING_FIELDS      re-scope · user-blocker · protocol-stop
NON-BLOCKING, by decision   gate-met · triggered-tok · cap-reached · budget-exhausted   (they end a
                            SESSION, not the milestone's ability to proceed)
```

**Invocation and result** — `/usr/bin/python3 blocking_state_guard.py --repo-root <rosetta>`, run from
`.agentspace/rosetta-extensions/stack-core`, at corpus `a95a356` + this commit:

```
109 graded iter(s) in m257x-platform-realignment
fields seen: budget-exhausted, cap-reached, gate-met, protocol-stop, re-scope, triggered-tok, user-blocker
```

**8 blocking gradings. This file named 3 of them.** The other **5** were invisible to every sweep this
audit has ever run — including §11's, which found one of the five by reading rather than by deriving.

| iter | field | what it routed to | disposition |
|---|---|---|---|
| iter-47 | `user-blocker` | the 7-blocker reading escalated for a ruling | **CLOSED** — answered in-session; the loop continued at iter-48 |
| iter-48 | `user-blocker` | *"the reading is mostly NOT repair-induced"* — the stronger form of iter-47's question (`EXIT_REASON: user-blocker`) | **CLOSED** — superseded by TOK-04, which re-scoped the unit of repair |
| iter-49 | `user-blocker` | the same question a third time | **CLOSED** — same route as iter-48 |
| iter-55 | `user-blocker` | clause 1 RED root-caused to a platform-side version skew | **CLOSED** — the concurrent lane re-proved clauses 1–2 at `0c91421` |
| iter-57 | `user-blocker` | *"the user paused the session"* — an exit, not a question | **CLOSED** — session resumed at iter-58. Recorded because the FIELD cannot tell a pause from a question, and a sweep that reads the field must not silently drop either |
| iter-68 | `protocol-stop` | the harden cadence hit 10 tiks; the loop stops so a harden pass can run | **CLOSED** — harden pass ran; 26 passes exist today |
| iter-116 | `re-scope` | `TOK-07` refuted by its own pre-registration (`P = 37` vs `P ≥ 15`) → a user re-scope conversation | **SUPERSEDED** by iter-119 — the conversation happened and produced `TOK-08` |
| iter-119 | `re-scope` | `TOK-08` refuted the same way (`P = 22` vs `P ≥ 19`); no successor strategy is permitted | **CLOSED at iter-283 by `TOK-09`** — corrected at the close re-audit, [§13](#13--close-run-re-audit-2026-08-11). This row read *"🔴 OPEN — this is Q5, the milestone is holding on it"* for 164 iters, and it was still saying so **after** the re-scope conversation had happened. `TOK-09` (`decisions.md`, iter-283, 2026-08-11) is a **USER-SUPPLIED re-scope** and is precisely the ruling this row was waiting for: it closed the scope to a fixed list and placed clause 5's sentence-level reading **out of scope**. ⚠️ **Out of scope is NOT met** — `TOK-09` says so in terms and this table must never be read as grading clause 5 clean |
| iter-259 | `user-blocker` | the `/dev-up` tooling puts a dev stack in `stack-dev/`, **another project's ACTIVE workspace** (367 commits on no remote, a live worktree) and `make init` is skip-if-present, so the bring-up would ADOPT their tree. Nothing there was touched; the dev half of clause 1 could not be proven | **CLOSED — the user LIFTED the prohibition mid-iter-261** (*"don't bother me and use that repo… don't search for exceptions"*), and iter-262 proved the dev half on the sanctioned path: 5/5 `core` containers from current `main`, `/api/health` 200, 172 migrations, 137 `public` tables, zero decommissioned schemas. **Recorded at M257x iter-269, and only because `blocking_state_guard` was RED on it** — the resolution happened in the iter stream and never reached the file a close gate reads, which is precisely the gap this table exists to close |

| iter-280 | `user-blocker` | the whole-section suite closed **RED (13 failed)**; everything attributable to that iter was found and fixed (22 → 13) and the residual sat in four files that pass standalone on both trees. The decision routed out was *spend an iter on the test-isolation defect plus `FIX-M257x-278`, or accept a RED section gate at close* | **CLOSED — ruled at build-iter run 42: we fix it.** A RED section gate at close is not acceptable, and the residual was not even attributable, because the control-tree full run aborted on the open `FIX-M257x-278`. iter-281 closed that fix **as a class** (6 sites in 3 files, fenced) and re-measured on both trees. **Recorded here at iter-281 — and again only because `blocking_state_guard` was RED on it.** That is the SECOND consecutive occurrence of this exact shape after iter-259, which was recorded ten iters late. The fence catches it every time; nothing runs the fence at close, which is what `ROUTE-M257x-280-the-31-minute-gate-is-skipped-because-it-is-31-minutes` is about |

**What the enumeration says that the count does not.** Eight of the nine are closed, and closing them is
not the finding. The finding is that **six of nine were unrepresented in the file a close gate reads**,
and that the one that is genuinely open is the one a `user-blocker`-keyed sweep is structurally least able
to see — because a *scope* decision is graded `re-scope`, and scope decisions are the ones that reach the
user latest and matter most.

**And the shape now has a rate, which is the part worth acting on (M257x iter-281).** Both of the last two
`user-blocker` gradings reached this table **only** because the fence went RED in a *later* iter — iter-259
ten iters late, iter-280 one iter late. The mechanism is structural, not carelessness: the grading is
written in the iter's `## Close` section, and the audit is a different file in a different lane, so the
write that creates the finding and the write that would represent it are never the same edit. **The fence
is not the gap; running it is.** A close that costs half an hour of whole-section suite gets skipped, so
the fence that would have caught this in the same iter is exactly the one nobody can afford to run.

**The class, named:** this is *green over something never checked* — the same class as the twelve
instruments harden pass 26 counted — landing in the milestone's **own close gate**. The audit was not
wrong about `user-blocker`. It was wrong about what it had measured.

**Controls** (`tests/test_blocking_state_guard.py`, 10 tests). The mutation that matters is the historical
bug itself: shrinking `BLOCKING_FIELDS` back to `("user-blocker",)` must **lose** the iter-119 finding —
a control that survives that mutant is not isolating the mechanism (harden pass 26's `_NOT_A_CITATION`
lesson). Anti-vacuity: breaking the grading header in all 25 fixture iters must exit 2, not sweep clean.
Every mutation asserts **it applied** before its result is read.

**New tokens booked here:** `FENCE-M257x-iter121-blocking-state-sweep` (landed) ·
`FIX-M257x-iter121-audit-sweep-was-one-field-wide` (closed by this section).

---

## §13 — CLOSE-RUN RE-AUDIT (2026-08-11)

**Invoked by:** `close-milestone` Phase 1b, `developer-kit:audit-deferrals --scope=milestone`.
**Derived at:** rosetta `m257x/platform-realignment` working tree · rext authoring `ddbe5a88c`
(`describe=fast-build-m257x-iter-288-3-gddbe5a8`) · milestone at **iter-288 + 73 harden passes**.
**Prior:** §1–§12 of this file, last touched at iter-284. This section brings it to close; it does not redo it.

### Verdict — **YELLOW (conditional)**

**The rule applied.** RED requires *a repeat-deferral without a resolution decision*. Repeat-deferral
patterns are real here and are named below — but every one has an **available in-release destination**
(M257 and M258 both exist and are unstarted), and **`TOK-09` supplies standing user authority** for the
disposition *"anything found outside the closed list is recorded and routed, never absorbed."* What is
missing is therefore not a **decision** but its **recording**: one block-level routing entry. That is a
close-agent write, not a user sign-off, so it does not reach the escape hatch.

> **The condition, stated so it can be graded:** if the close records **one block-level routing decision**
> for the carried-token population (§13.4), this is YELLOW. **If it records nothing, the correct grade is
> RED** — an unrecorded route is the unfated drop the three-fate rule exists to prevent, and this milestone
> has booked that exact shape twice already (§12).

**⚠️ Clause 5 is NOT graded here and must never be reported as measured clean.** `TOK-09` places its
sentence-level reading **out of scope by user ruling** and says in terms that this *"is a scope ruling, not
a declaration that the unmeasured reading came out clean."* Clauses 1–4 are met; clause 5 is out of scope.

### §13.1 — The machine fence, run at close (this is the point)

`ROUTE-M257x-280-the-31-minute-gate-is-skipped-because-it-is-31-minutes` records that this fence *"catches
it every time; nothing runs the fence at close."* It was run. **Exit code 0.**

```
fence-tree: /Users/marco/workspace/anthropos/rosetta/.agentspace/rosetta-extensions @ ddbe5a88c · describe=fast-build-m257x-iter-288-3-gddbe5a8
blocking-state-guard: 273 graded iter(s) in m257x-platform-realignment; fields seen: budget-exhausted, cap-reached, gate-met, protocol-stop, re-scope, triggered-tok, user-blocker
blocking-state-guard: blocking fields checked: re-scope, user-blocker, protocol-stop (non-blocking, by decision: gate-met, triggered-tok, cap-reached, budget-exhausted)
blocking-state-guard: iter-116 re-scope: y — represented in deferrals-audit.md
blocking-state-guard: iter-119 re-scope: y — represented in deferrals-audit.md
blocking-state-guard: iter-259 user-blocker: y — represented in deferrals-audit.md
blocking-state-guard: iter-280 user-blocker: y — represented in deferrals-audit.md
blocking-state-guard: iter-47 user-blocker: y — represented in deferrals-audit.md
blocking-state-guard: iter-48 user-blocker: y — represented in deferrals-audit.md
blocking-state-guard: iter-49 user-blocker: y — represented in deferrals-audit.md
blocking-state-guard: iter-55 user-blocker: y — represented in deferrals-audit.md
blocking-state-guard: iter-57 user-blocker: y — represented in deferrals-audit.md
blocking-state-guard: iter-68 protocol-stop: y — represented in deferrals-audit.md
blocking-state-guard: OK — every blocking grading is represented in the deferral audit
```

**10 blocking gradings across 273 graded iters; all 10 named in this file; zero un-represented.** The
population grew 109 → 273 graded iters since §12 and produced **no new blocking grading** — iters 281–288
all graded `re-scope: n` / `user-blocker: n` / `protocol-stop: n`.

**And the fence's own limit, found by running it: it checks PRESENCE, not DISPOSITION CURRENCY.** iter-119
was reported *"represented"* while this file recorded it as **🔴 OPEN — the milestone is holding on it**,
164 iters after `TOK-09` had closed it. A name can be present with a **wrong** disposition and the fence
stays green. That is *green over something never checked*, in the fence built to catch exactly that —
booked as **`ROUTE-M257x-289-blocking-guard-checks-presence-not-disposition`**. The row is corrected in §12.

### §13.2 — `TOK-09` verified; §12's one open row is closed

`decisions.md` **`TOK-09`** (iter-283, commit `6572d6ee`, 2026-08-11) is a **USER-SUPPLIED re-scope**, not
agent-authored and not a triggered tok. It is the ruling iter-119's `re-scope` routed out for. Chain intact:
`TOK-07` → refuted (iter-116) → `TOK-08` → refuted (iter-119) → **`TOK-09` (user)**. §12's iter-119 row and
§10's user row are corrected. **Zero open user questions at close.**

> ⚠️ `state.md` still reads `259 iters`, `62 harden passes`, `last_updated: 2026-08-10`, and asserts
> **"there is NO active strategy … so there is no `TOK-09`"** — contradicted by `TOK-09` existing.
> Not a deferral; a stale-state repair the close owns.

### §13.3 — Open deferrals at close (11), each with its fate

**No item is unfated, and no item is a process failure** — `TOK-09` mandates *record and route, never
absorb*, and the scope-creep tripwire firing 11 times is the tripwire working.

**Iter lane — routed under `TOK-09` authority (5)**

| token | fate | justification |
|---|---|---|
| `ROUTE-M257x-284-demo-2-is-live-and-uncontained` | **LAND-NEXT → M258** | The running `demo-2` predates the iter-284 fix; the strip takes effect at the next `/demo-up`. A bring-up is the user's call and M258 owns cold bring-up. |
| `ROUTE-M257x-285-demo-2-cockpit-serves-a-stale-world` | **LAND-NEXT → M258** | Restoring the demo world needs a re-seed of the live stack the user is validating on. Same owner as the row above. |
| `ROUTE-M257x-286-next-web-manifest-baselines-have-drifted` | **LAND-NEXT → M258** | 9 live-clone freshness gates doing their job against a moved clone; re-anchoring is a deliberate pass, not a side-effect. |
| `ROUTE-M257x-282-prose-twin-REPORT-tier-residual` | **LAND-NEXT → M258** | 12 edge divergences the RED tier cannot reach; each needs a source read — per-item work, published on every guard run so it cannot go silent. |
| `ROUTE-M257x-282-intra-tree-prose-twins` | **LAND-NEXT → M258** | Out of the fence's scope **by design**, measured not assumed. Widening scope is a design change. |

`ROUTE-M257x-285-logout-swap-for-studio-and-academy` **CLOSED at iter-288** (fenced over the family, not
the members). `ROUTE-M257x-282-readme-does-not-index-13-of-37-fences` is a measured-not-repaired side
deliverable → **LAND-NEXT → M258** with the block.

**Harden lane — passes 72–73, post-dating this audit's body entirely (6)**

| token | fate | justification |
|---|---|---|
| **`ROUTE-M257x-h72-dev-path-carries-the-prod-bucket`** | **LAND-NEXT → M258, and SURFACE TO THE USER AT CLOSE** | **The highest-stakes open item.** The container-side prod-bucket strip is **DEMO-only**: `/dev-up` applies only `gen_override.py`, which emits no strip, so a dev `backend` still holds `STORAGE_S3_BUCKET=production-storage…` — **measured live by `docker inspect` on the running stack, not inferred.** What bounds it is the **absence of credentials, not a pointer override**, so an operator whose `platform/.env` carries AWS keys re-opens on dev exactly what iter-284 closed for demos. **Disclosed, not silent** (`corpus/ops/safety.md:326-334`). Same hazard class the user ranked **#1** under `TOK-09`, on the path the user is on right now. **Not the escape hatch** — see §13.5. |
| `ROUTE-M257x-h72-gen-injected-override-cites-81-82` | **LAND-NOW (trivial)** or M258 | Off-by-one in a **safety-critical comment** (`:81-82` cited; measured `:82`/`:83`). Comment-only, minutes. |
| `ROUTE-M257x-h73-suite-census-unittest-path-is-dead-on-this-interpreter` | **LAND-NEXT → M258** | The two-runner cross-check is currently a **one-runner check reporting RED**. The fix is a contract change to `run_one` (or 5 `__init__.py` files that change pytest import semantics tree-wide) — a design decision, correctly not a harden pass's. |
| `ROUTE-M257x-h73-demo-stack-live-arms-red-while-a-demo-is-up` | **LAND-NEXT → M258** | 9 `demo-stack` failures are live arms asserting against an in-use clone. Arguably not a defect; the defect is that **nothing in the output distinguishes an in-use clone from a broken patch chain**. |
| `ROUTE-M257x-h73-battery-stage-stdlib-set-is-interpreter-versioned` | **LAND-NEXT → M258** | PEP 594 removals; the fence's own message asks for the allowance class. |
| `ROUTE-M257x-h73-readme-documents-a-python3-that-has-no-pytest` | **LAND-NOW (trivial)** or M258 | `stack-core/README.md:185` documents a command that exits *"No module named pytest"* on this host. |

> **Close-blocking test evidence, stated because a close gate will ask.** Pass 73 is the milestone's final
> whole-section run: `stack-core` **2,248 passed / 18 failed**, and the entry **self-corrects a `demo-stack`
> GREEN it had reported off a truncated tail** (`1,085 passed / 9 FAILED`). Of the 18: 7 did not reproduce
> at a different load, 8 are one `suite_census` interpreter defect, 1 is PEP 594, 2 were the real defect and
> are **fixed at `ddbe5a8`**. All 9 `demo-stack` reds are live-arm/in-use-clone. **No red is attributable to
> unfixed product or tooling logic** — but "green" is not the right word for it either, and the close should
> not use it.

### §13.4 — Repeat-deferrals (the pattern finding)

Mechanical sweep over the milestone tree (`ROUTE|FIX|CHECK|DEF|RF|INVESTIGATE|MEASURE|DECIDE|DOC|REPOINT|FENCE|HOST`,
backticks stripped, trailing `._-` normalised): **487 distinct tokens · 215 carried across ≥ 3 distinct
iter dirs.** Highest carriers, by distinct iter dirs:

| carried | token | disposition |
|---|---|---|
| 36 | `CHECK-M257x-iter38-ai-act-classification` | **DROPPED** `D-M257x-102-4` |
| 33 | `FIX-M257x-iter56-assignment-flake` | hold retired; repair routed |
| 32 | `FIX-M257x-iter135-adjudicated-live-defects` | closure markers present |
| 30 | `FIX-M257x-iter53-union-set` | **DROPPED as subsumed** `D-M257x-102-2` |
| 27 | `FIX-M257x-iter145-sha-baseline-drift` | closure marker present |
| 24 | `CHECK-M257x-iter52-second-ai-manager` | **quiet after iter-76, no closure marker** |
| 21 | `ROUTE-M257x-235-fence-scope-is-unread` · `ROUTE-M257x-235-runnable-block-has-two-halves` | **quiet after iter-256, no closure marker** |
| 20 | `ROUTE-M257x-236-disclosure-scope-is-document-level` | **quiet after iter-256, no closure marker** |
| 19 | `ROUTE-M257x-238-container-vs-native-is-undrawn` · `FIX-M257x-iter57-within-block-drift` | **quiet, no closure marker** |
| 17 | `DEF-M257x-iter80-storage-prod-bucket` | **CLOSED at iter-284** — the ten-iter chronic, discharged by being *exercised* |
| 17 | `DEF-M257x-iter101-briefing-rext-tree` | **quiet after iter-119, no closure marker** — §11 priced it at "~30 min" |

**The finding, and §11 already named it:** the failure mode here is not re-deferral, it is **going quiet**.
The tree is **append-only** — closure is recorded in prose beside a token, never by deletion — so a token
that stops being restated is indistinguishable from one that was fixed. **This is a MARKER count, not a
STATE count**, and no per-item state claim is made for the 215. (An honest per-item number costs ~1 grep
each; the prior audit said so at §6 and it is still true.)

**Fate for the block — one decision, not 215.** Under the three-fate rule these are **LAND-NEXT (Fate 3)**
to **M258**, recorded as a single routing decision in the milestone's `decisions.md`. **Do not route to
M257:** M257's own `exit_gate` still names **`odysseus`**, retired by `D-v28-15`, so routing there inherits
an unmeetable-as-written gate (already booked as `ROUTE-M257x-225/226`). M257 needs its gate re-cut before
it can hold anything.

### §13.5 — Escape hatch: **NONE. Nothing blocks the close.**

Conservatively assessed, item by item. The escape hatch requires *release-scope-breaking* work needing user
sign-off. **No open item qualifies**, for one structural reason: **M257 and M258 both exist, in this
release, unstarted** — so every routed item has an in-release home and nothing needs a cross-release punt.

`ROUTE-M257x-h72-dev-path-carries-the-prod-bucket` was tested hardest against the bar and **does not clear
it**: (1) `TOK-09` pre-authorises exactly this disposition for anything outside the closed list, and the dev
path is outside it — the user's item #1 named *the demo's* storage config; (2) it is **disclosed in the
corpus**, not silent; (3) it is **bounded by default** — the dev emitter clears the `~/.aws` bind
(`gen_override.py:195-196`) and there is no M239 env bridge on that path; (4) the remedy costs dev uploads
a broken-thumbnail trade, which the harden pass correctly called an **operator's** decision. It is a
**disclosure to surface at close**, not a sign-off gate.

**Recommended close actions (none require user sign-off):** record the one block-level M258 routing
decision (§13.4); repair `state.md` (§13.2); optionally land the two trivial LAND-NOWs; and **state the
dev-path prod-bucket residual to the user in the close report** — because the user is on a dev stack now.

_Close-run re-audit, 2026-08-11. Read-only except this file. Zero corpus edits, zero code edits, zero
platform-repo edits; no stack touched._
