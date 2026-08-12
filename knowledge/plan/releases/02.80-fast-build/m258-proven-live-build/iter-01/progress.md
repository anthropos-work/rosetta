# M258 iter-01 — progress

**Type:** tok (bootstrap)

Opened 2026-08-12T04:27:55Z on branch `m258/proven-live-build`, created from
`release/02.80-fast-build` @ `b8ae62c6`.

## Phase 0b — pre-flight KB-fidelity gate

Run before the strategy was authored, per the skill's contract. Verdict **YELLOW**; full report at
[`../kb-fidelity-audit.md`](../kb-fidelity-audit.md), triples table at
[`../spec-notes.md`](../spec-notes.md).

**Rung zero first** (`corpus/ops/verification.md` § PRE-FLIGHT RUNG ZERO — *tagging is not publishing*):
the authoring copy is clean at `679a5f7` = `origin/main` = tag `fast-build-m257-close`, and that tag
**is on origin** (`15c1352…`). But `.agentspace/rext.tag` names **`fast-build-m257-iter-09`**
(`8956e69`) — **the pin is one tag behind the published tooling**, and the gap is not cosmetic: the
close tag carries the M257 close's three fail-open repairs **to `buildbench`, the instrument M258 reads
its own p50 from**. Recorded as `R0`; tik 1 opens by re-pinning.

**Inventory:** 8 topics, **8 PAIRED, 0 BLIND-AREA.** Both `Delivers →` targets exist and cover their
subject.

**What was probed rather than assumed** — five load-bearing claims, all **ALIGNED**:

| claim | how it was settled |
|---|---|
| `FIX-M257-content-stories-pair-count` still open | read both counters: `content-pairs.ts:115` has the `manager_presence_only` branch, `run-content-stories.sh:145-152` does not → 47 vs a pinned 45 → `exit 2`. **Open exactly as written** (see `decisions.md` D4 — it first read as fixed) |
| `ptvalidate` invoked nowhere outside its own tests | swept every `.sh`/`.go`/`.py`/`.ts`; every hit outside `cmd/ptvalidate/` and `*_test.go` is prose in a comment |
| 30 live Playthroughs / 31 use cases / 10 products | ran the validator: `manifest VALID: 10 product(s), 31 use case(s), 30 live Playthrough(s), 1 TODO`. A raw `grep '@pt:'` says **35** — the wrong instrument, and it would have overstated the batch |
| `autoverify` check (h) container-liveness landed at M257 | present at `autoverify.sh:564` |
| the `DEMO_STORIES_PRESET` seam resolution (a) rests on | present at `up-injected.sh:261`, with all four downstream exports reading `$STORIES_PRESET` |

**Findings, all applied:**

- **F7 — 13 stale line anchors in M258's own `overview.md`.** Every one **in range** and landing on
  unrelated content, so no out-of-range lint could see them; in every case the **substance holds**.
  That is what makes them costly rather than cosmetic: an agent following `up-injected.sh:218` for the
  seam that resolution (a) rests on finds a comment about `INJECT_SVCS` and may conclude the seam does
  not exist. Repaired in place — a stale reference is a repair, not a re-plan, so no clause, target or
  number was touched. *(`demo-up-defaults.md:47` cites the same seam **correctly** at `:261` — the
  corpus was fresher than the milestone plan.)*
- **C1 — the batch half of the composed budget has no published wall-clock.** `overview.md` asked M256
  to measure the reset-to-seed leg; it did not, and no corpus doc carries a suite wall-clock. The only
  figure anywhere is **56.6 s** (M256 `progress.md` iter-02) over **18 specs**, where the suite closed
  at **209 passed / 30 live Playthroughs** — so it prices a suite an order of magnitude smaller and
  must not be quoted as the batch half. Backfilled to `spec-notes.md` + flagged in `overview.md`.
- **C2 — M256 escalated that this suite's timing is not decidable at n=3 on this host** (2.04× spread,
  no trend, six runs) and `D-v28-12` was re-cut for it. That escalation reached neither this plan nor
  the corpus. M258's gate is a **p50 over n=3** whose second half is that suite. Backfilled as
  **evidence, not a re-cut**.

## Phase 0d — pre-flight tooling check

**SKIPPED.** A strategy-authoring iter wires no artifacts through a gate pipeline (the skill's named
skip case). Tik 1 will run it.

## Phase 1 — the strategy

`TOK-01` authored in the **milestone-root** [`decisions.md`](../decisions.md):
**measure the composition before engineering it.**

Four strictly-ordered tiks — re-pin + measure both halves in one campaign → wire the batch-gate at the
tail hook (`up-injected.sh:2810`) under `D-v28-3` semantics → land the world-contract restore leg →
the composed 3× cold campaign. The **order is the strategy**: the milestone has one measured half and
one never-measured half, and wiring first would put its only genuine unknown at the end — the shape
M257 spent three iters inside.

Two decisions the plan required and this iter took (`decisions.md` D1, D2):

- **World contract → (b) restore after**, refuting (a) **on the gate's own text**: the gate requires a
  *"presenter-usable world"* and the overview's own paragraph on (a) ends *"But it is not a presenter
  demo."* M254 left `billion` in exactly that un-restored state; (b) is what stops M258 making that
  swap the outcome of every bring-up.
- **Single-box `--no-public-host` mode**, disclosed as `TOK-01`'s one overturnable assumption — it is
  the only mode in which *"one cold command"* is literally satisfiable, and it proves the composition
  in a mode the presenter never uses. The peer path (`--reset-only`) is unbroken, so the re-cut is
  cheap if the user wants it.

## Phase 3 — signal

No metric delta — a tok does not move the gate.

| | |
|---|---|
| strategy record | `TOK-01` (milestone-root `decisions.md`) |
| gate distance | bring-up half **286.99 s** measured; batch half **UNMEASURED**; **193.01 s** of the 480 s ceiling unaccounted |
| known-context carried | 6 items (`R0`, `C1`, `C2`, `F1`, `F2`, + the SUSPECT-UNROUTED rule) |
| environment | `macmini`, `load1 1.81` at close (contended host; `demo-2` 11 containers + the 5-container dev stack are the user's and were not touched), 175 GiB free |

## Close — 2026-08-12

**Outcome:** `TOK-01` authored — *measure the composition before engineering it*; the world contract
resolved to **(b) restore after** on the gate's own text; the Phase-0b gate cleared **YELLOW** with 13
stale anchors repaired and two never-propagated measurements recorded at the destination.
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n *(bootstrap toks do not exit)* — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n *(0 tiks so far)* — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** D1 (world contract → (b)), D2 (single-box mode, disclosed), D3 (measure before wiring),
D4 (`F1` re-verified against code and survived)
**Side-deliverables:** the Phase-0b repairs — 13 anchors in `overview.md`, `spec-notes.md` written from
stub with the topic→doc→code triple table, and the C1/C2 annotations. Applied under the audit's own
Phase 6, not folded into this iter's close status.
**Routes carried forward:** none new. Six known-context items ride in `TOK-01`; the inherited M257x /
M257 lists stay **SUSPECT-UNROUTED until verified open**, per the rule `D4` demonstrates the value of.
**Lessons:**

- **Verify an inherited item against code before working it *or* discarding it.** `F1` read as
  already-fixed from its commit history and was **open** in the file that mattered. A route list goes
  stale in both directions, and the optimistic direction is the expensive one — it drops real work.
- **An anchor that is in-range and wrong is invisible to lints and costly to readers.** All 13 stale
  anchors passed an out-of-range check; the substance held in every case, which is exactly why a reader
  following one concludes the *mechanism* is missing rather than the *citation*.
- **Pick the instrument, not the grep.** `grep '@pt:'` says 35 Playthroughs; the validator says 30. The
  batch this milestone is gated on is the validator's number.
