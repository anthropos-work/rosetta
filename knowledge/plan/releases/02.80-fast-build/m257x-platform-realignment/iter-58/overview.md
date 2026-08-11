---
milestone: M257x
iter: 58
iteration_type: tik
status: closed-fixed
created: 2026-08-04
active_strategy: TOK-04
target_clause: "1, 2 (re-prove at advanced pins)"
refs:
  platform: 0dab54dfac6beacdef54a671e2500d3940fd7329   # origin/main, re-fetched at open (P3); clone level
  platform_source: stack-demo/platform
  app_at_open: v1.365.0 (bff61c91)                     # the committed pin
  app_origin_main: v1.366.0 (b948604f)                 # re-checked at open; has NOT moved since iter-56
  rext_head: 28c99d0e0fd89d4d8d62b2a93c84fdfb0ae4be72  # main, clean; tag fast-build-m257x-iter-57 (on origin)
  rext_pin_at_open: fast-build-m257x-iter-56           # .agentspace/rext.tag — one tag behind HEAD
  rosetta: 1937e1f9ead9cb9f0fc96297be11b604850c615a    # at open
  instrument_clonespin_sha256: e3a7227aabfb02903f3cc39638324e9a25fbc78273b095137200b9189853e6cc
  instrument_upinjected_sha256: 5ef0cc37af50f2ee80d1d95f3f726e8a1428175b43544dd61f062fc0d636d629
  taken_at: 2026-08-04T07:43:45Z
---

# iter-58 — advance the pins, and prove the advance cold

**Type:** tik
**Active strategy reference:** `TOK-04: pin the target, or stop calling it a measurement` — **P1** (every
measurement states its refs, in the artifact, at the moment it is taken), **P2** (every instrument is a
committed file), **P3** (the platform ref chosen/recorded/re-checked at open AND close; *the iter that
detects a move re-points in that iter, as its first act*), **P4** (derive, else fence, else declare).
Plus protocol **§7 rule 4** (*advance the pins deliberately and record what the advance contained*) and
**§7 rule 5** (*prove it cold*).

## Step 0 — Re-survey before targeting (mandatory)

**A machine reboot happened between iter-57's close and this open, and it moved the state.** Re-surveyed at
open rather than inherited:

| thing | inherited claim | measured at open | consequence |
|---|---|---|---|
| `demo-1` | orchestrator: **"GONE — 0 containers, 0 total"**; handoff: *"still UP, do not tear it down"* | **11 containers, all present** (`Exited (255)`, 2 `Restarting`) — the orchestrator's zero was a **dead-daemon false absence**, `D-M257x-58-1` | neither inherited claim was right; the stack exists but is not serving |
| docker daemon | (not mentioned by either) | **not running** post-reboot; started at open — and *this* is what produced the zero | §5 rule 1, live: an empty result from a failed command is not evidence of absence |
| platform | `0dab54d`, clone and origin level | **`0dab54d`, still level** (re-fetched) | P3 satisfied at open; no re-point owed |
| app origin | `v1.366.0` ahead of the `v1.365.0` pin | **`v1.366.0`, has NOT moved again** | the routed fix is still exactly the routed fix |
| rext pin | `iter-56` while HEAD/tag is `iter-57` | confirmed; `fast-build-m257x-iter-57` **is on origin** | the deliberate gap (`D-M257x-57-6`) is now actionable |

**iter-57's stated reason for declining the app advance survives the correction, but is spent anyway.**
It declined because *"it changes what a bring-up consumes and `demo-1` is live clause-1/2 evidence."*
`demo-1` does still exist — but it is `Exited (255)`, not serving, and clause 1's evidence was never the
running containers: it is the three checked-in `autoverify.json` verdicts and their refs blocks, which are
on disk and unaffected. Restarting a stopped stack would not reproduce a **cold** cycle either. So the
protective reason no longer applies, and the two deliberately-held-back pins are both held back **pending
an iteration that takes a cold cycle** — the only condition under which their effect is observable.

**Target selected: advance both pins and re-prove clauses 1 + 2 cold at the advanced refs.** This is a
substitution under the same TOK, not a re-scope: TOK-04's `Next-tik direction` still names the 81-site
sweep (clause 5), but TOK-04 **change 2** is explicit that the ref baseline is re-established *before* any
further reading, and both pins are exactly the "input that moved" class P3 governs. Clause 5 is iter-59.

## Cluster / target identified

Two routed handlers with one shared precondition (a cold cycle), plus one pre-registered question that
rides along free.

1. **`FIX-M257x-iter56-app-ref-moved`** — the committed pin `demo-stack/clones.pin.json` says
   `app: v1.365.0`; app origin/main is `v1.366.0`. Routed by iter-56, declined by iter-57 with a reason
   that no longer holds.
2. **`D-M257x-57-6`'s deferred half** — `.agentspace/rext.tag` = `fast-build-m257x-iter-56` while rext
   HEAD/tag = `fast-build-m257x-iter-57`. iter-57's reasoning was sound *and time-limited*: assertion F is
   a corpus-fidelity guard no bring-up consumes, so advancing the pin then would have been unobservable.
   **In an iteration that takes a cold cycle it is observable**, which is the condition iter-57 named.
3. **`FIX-M257x-iter56-assignment-flake`** — iter-56 routed it with an explicit instruction: *"Measure
   first whether app `850917d7` (`fix(assignments): …`) bears on it — the domains coincide and nothing has
   measured the link."* `850917d7` **is in the `v1.365.0..v1.366.0` range**. Taking clause 2 at the
   advanced pin answers the question at zero marginal cost.

### §7 rule 4 — what the app advance contains (`v1.365.0 → v1.366.0`, 5 commits), measured at open

| dimension | finding | how measured |
|---|---|---|
| commits | 5 — one version chore, one merge, three functional (`feat(clerk)` force-join at signup, `fix(clerk)` review, `fix(assignments)` join fall-through) | `git log --oneline v1.365.0..v1.366.0` |
| **migrations** | **0** — the range touches no migration file at all | `git diff --stat … -- '*migrations*'` → empty |
| **destructive DDL** | **0** | grep for `DROP TABLE`/`DROP COLUMN`/`RENAME`/`DROP SCHEMA` in the migration diff → empty |
| **new hard-required config** | **0** — the one `+log.Fatalf` in the diff is the **pre-existing** `clerk events manager` fatal, *moved* below `orgManager`/`assignmentManager` so the webhook can force-join. Same line, same error, new position | full `main.go` diff read, not grepped |
| **new env reads** | **0** — the `+os.Getenv` line is that same relocated call; its three `CLERK_*` reads are unchanged | same |
| RPC addresses | `STORAGE_RPC_ADDR` unchanged (still read in `main.go`, still the same call) | same |
| blast radius | 11 files, all under `internal/{assignments,clerk/events,organization}` + `main.go`; 5 of 11 are new test files | `git diff --stat` |

**That is the safest shape an advance can have** — additive behaviour, zero schema movement, zero removed
contract. The class that broke at v2.1 and v2.7 was a *removed* table or schema; there is no removal here.

**§5 rule 28 / the iter-56 lesson applied — "check the remedy actually contains the fix before taking it."**
This advance is **not** aimed at a defect, and saying so is the point: iter-56's advance was justified by a
storage-fold hypothesis that turned out to exist at no ref. This one is justified by **P3 alone** (the
recorded ref should be the current one) and carries **one** falsifiable rider, pre-registered below.

## Hypothesis

Advancing both pins is **behaviour-neutral for the bring-up**: the app advance is additive-only, and the
rext advance contains only assertion F (a corpus guard) plus its tests. So three cold cycles come back
green and the Playthrough suite comes back `30 passing / 0 failing`, at refs that are — for the first time
in this milestone — **the current ones on every axis at once**.

## Expected lift

Clauses 1 and 2 move from *MET at superseded pins* to **MET at current pins**. Gate reading stays **4 of 5**
by count; what changes is that its evidence stops being stale-by-construction under the gate's own
"against origin HEAD" wording.

**Pre-registered, therefore refutable:**

1. **3 consecutive cold `demo-down --purge` + `demo-up` cycles reach `green:true / warnings:0`**, with
   distinct timestamps (the M236 stale-verdict guard). *If any cycle goes red, that is a real finding and
   the highest-value one available — it would mean a purely-additive app advance is not behaviour-neutral.*
2. **Container count is 15**, not 16 — the deleted WunderGraph router stays deleted.
3. **The Playthrough suite reads `passing=30 failing=0`.**
4. **`FIX-M257x-iter56-assignment-flake` is NOT fixed by `850917d7`.** iter-56 diagnosed it as a *test-side*
   race — `toBe(before - 1)` over a baseline sampled while the members grid is still settling, observed
   `16 → 14`. `850917d7` changes server-side join fall-through and already-member matching, which does not
   touch when the grid settles. *If the flake disappears at v1.366.0, iter-56's diagnosis was wrong and the
   route must be re-written, not closed.*

## Phase plan (4 planned lines — declared, per the scope-creep tripwire's multi-step carve-out)

- **A — verify the baseline first (the handoff's designated first act).** Full `stack-core` suite against
  the measured **1F / 599** baseline. The single expected red is the **iter-48 perishable answer-key
  fixture** (TOK-02 step 4) — *do not spend it*. Anything else is iter-57's regression and is fixed before
  any new work. Blast radius is `platform_alignment_guard.py` + its tests, but `repair_postcondition.py`
  reads `FENCE_KIND` statically from that module, so that path is checked explicitly.
- **B — advance both pins**, each with its §7-rule-4 record: `clones.pin.json` `app → v1.366.0`,
  `.agentspace/rext.tag` → `fast-build-m257x-iter-57` (verified **on origin**, per the CLAUDE.md
  *tagging-is-not-publishing* rung-zero rule).
- **C — prove it cold** (§7 rule 5): 3 consecutive `demo-down --purge` + `demo-up` cycles, each verdict read
  from the bring-up's **own** `autoverify.json` (not a standalone re-run — the iter-11 vantage lesson), with
  distinct timestamps.
- **D — clause 2 + the rider**: the Playthrough suite at the advanced pin; record the `pt-assignment-assign`
  outcome against pre-registration 4. Re-check the platform ref at close (P3).

## Escalation conditions

- **Platform ref moves mid-iter** → re-point **in this iteration, as its first act** (P3), not routed forward.
- **A cold cycle goes red** → that is the finding; do not advance past it and do not re-run until green.
  Diagnose at the first red, record it, and if it is not closable inside this iter's budget, **revert the
  pin advance** rather than shipping an unproven pin.
- **Budget exhaustion mid-cycle** → do NOT close, do NOT commit partial work; exit `user-blocker` per the
  skill's Phase 4 Step 0.

## Acceptable close-no-lift outcomes

- The advance is taken, proven cold, and reveals **nothing** — a null result at stated current refs is a
  complete iter under P1, because the milestone's whole thesis is that unstated refs make every number an
  anecdote.
- A cold cycle goes red and the iter closes on a **documented falsification** of "purely-additive advances
  are behaviour-neutral." That is worth more than a green cycle.
- The `stack-core` baseline comes back off 1F/599 and the iter spends itself repairing iter-57's
  regression. Fixing the instrument before using it is Phase A's whole reason for being first.
