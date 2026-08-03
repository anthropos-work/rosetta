---
iter: 57
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-03
active_strategy: TOK-04
target_clause: 3
refs:
  platform: 0dab54dfac6beacdef54a671e2500d3940fd7329   # origin/main, re-fetched at open; clone level
  platform_source: stack-demo/platform
  rext_head: 5833062c98c0aac3e49f8eb4ed4100072b3e337b  # main, clean; tag fast-build-m257x-harden-p15
  rext_pin: fast-build-m257x-iter-56                    # .agentspace/rext.tag — deliberately behind HEAD
  corpus_head: 00f693d826d5996f8e0ef1bc60f8cfe0bb2bae46
  instrument_map_sha256: 8ad8075930ae14d437c28955d3d406c485fea8f9f2ee2f7e9d60988070fcde1a
  instrument_guard_sha256: 1bd6a03f1c8778431c9c217a26c8735c045ea87f5ff9567a54a5b535442b2db2
  taken_at: 2026-08-03T17:07:17Z
---

# iter-57 — clause 3: fence the map's CITATIONS, not just its membership

**Type:** tik
**Active strategy reference:** `TOK-04: pin the target, or stop calling it a measurement` — specifically
**P4** (*derive, else fence, else declare prose-under-review — in that order, and mark which one you used*)
and **P3** (platform ref chosen/recorded/re-checked at open and close).

## Step 0 — Re-survey before targeting (mandatory)

TOK-04's `Next-tik direction` named *"iter-56 = the 81-site sweep as one derived class"*. iter-56 did not
take it — it took the ref baseline instead and **restored clauses 1 and 2** (gate 1 of 5 → 3 of 5). The
named next-target is therefore **stale in sequencing but not in substance**. Re-survey at open:

| clause | reading at open | how confirmed |
|---|---|---|
| 1 | MET | iter-56 evidence, 3 cold cycles at `0dab54d` |
| 2 | MET | `passing=30 failing=0` at `0dab54d` |
| 3 | **NOT MET** | re-run at open: both fences **GREEN**, map prose **FALSE** (below) |
| 4 | MET | 4 guards, 74 OK |
| 5 | NOT MET | untouched; not re-cut |

**Target selected: clause 3**, per the orchestrator's direction and because it is the cheapest restoration
left. This is a **substitution under the same TOK**, not a re-scope: TOK-04 named the 81-site sweep; clause
3's own repair is the same P4 mechanism applied to a 5-claim subset that is *already routed*
(`FIX-M257x-iter55-map-storage-messenger`, `CHECK-M257x-iter55-map-prose-unfenced`).

## Cluster / target identified

Two routed handlers, one root cause.

**Measured at open, and it is the finding:** both clause-3 fences run **GREEN over a map with 5 false
claims**.

- `platform_alignment_guard.py` → `OK — ... agree in both directions`
- `anchor_construct_guard.py` → `OK — every resolvable anchor names a construct` (113 resolved)

Neither is broken. They fence **different properties than the one that failed**:

- `platform_alignment_guard`'s assertion **D checks only that the evidence cell is non-empty**. A cell
  full of citations to lines that no longer exist is, to assertion D, indistinguishable from a correct one.
- `anchor_construct_guard`'s `_QUALIFIED` regex requires a `/` in the path or a `.md` suffix. It was
  scoped that way deliberately (its docstring records the 134-findings-all-ports over-match that forced
  it), but the consequence was never measured **for this file**.

Measuring it now:

| citation class in the map | count | seen by a fence? |
|---|---|---|
| path-qualified (`app/main.go:573`, `storage/terraform/main.tf:19`) | 15 | **yes** — `anchor_construct_guard` |
| bare platform file (`docker-compose.yml:90`, `repos.yml:18-20`, `common.yml:…`) | 25 | **no** |
| bare continuation (`:178`, `:161`) | 12 | **no** |

**37 of 52 citations — 71% of the map's evidence — sit in a class no fence can see.** Both dead
citations iter-55 named are in it. *A fence over membership says nothing about prose* is the routed
finding; the measurement above says **why**, and it is not a wording problem.

## Hypothesis

The map's evidence cells are a **more constrained context than the corpus at large** — inside a row whose
first cell names the subject repo, a bare `docker-compose.yml:90` is unambiguous by construction. So the
over-match that forced `anchor_construct_guard` narrow does not apply here, and the bare class can be
resolved **map-scoped** without touching the corpus-wide guard's calibration.

Add **assertion F** to `platform_alignment_guard.py`: every citation in an evidence cell must resolve
against the platform clone at the recorded ref, and — when the cited **path** does not itself name the
row's subject — the cited **line(s)** must. A citation earns its authority by naming its subject
somewhere; when neither path nor line does, it has drifted off its subject.

## Expected lift

Assertion F goes **RED on the current map**, naming ≥ 2 dead citations (the two iter-55 found by hand)
without being told about them; then **GREEN** after a minimal repair. Clause 3 moves from
*half-fenced with false prose* to *fenced-and-true*, i.e. **MET**.

**Pre-registered, therefore refutable:**
1. Assertion F names **≥ 2** findings on the unrepaired map. *(If it names 0, the fence is calibrated to
   nothing and must not ship.)*
2. It names **more than the 2** iter-55 found by hand — the hand reading had ~43% recall on a 7-seat pass;
   one seat on one file should be worse.
3. The repair is **≤ 3 rows** of the services table plus the §5 narrative line.

## Phase plan (3 planned lines — declared, per the scope-creep tripwire's multi-step carve-out)

- **A — fence.** Assertion F in `platform_alignment_guard.py` + tests (incl. a **no-op control that
  survives** and an **inverted mutant** that must die, §8 rules 4/5). Reach is **named and counted**, and
  **0 resolvable citations refuses (exit 2)** rather than passing (§5 rule 8 / the positive-control rule).
- **B — watch it RED**, on the real unrepaired map. Record what it names vs the pre-registration.
- **C — repair + mark.** Minimal-scoping edits (TOK-03 move 3: deletion > scoping edit > rewrite). Claims
  that are neither derivable nor fenceable get the **prose-under-review** marker P4 says the map must
  carry visibly and currently does not.
- **D — GREEN + regression.** Re-run all four platform-alignment guards + the `stack-core` suite against
  the measured 1F/599 baseline. Re-check the platform ref at close (P3).

## Escalation conditions

- Platform ref moves mid-iter → **re-point in this iteration, as the first act** (P3), not routed forward.
- Assertion F cannot reach the bare class without over-matching → do **not** ship a fence that skips in
  silence; record the reach honestly and route the residue.
- A repair would require rewriting a row wholesale → prefer the prose-under-review marker over a rewrite.

## Acceptable close-no-lift outcomes

- Assertion F ships, runs RED/GREEN, and reveals that iter-55's 5 claims were **not** all falsified —
  a documented falsification of a routed finding is a complete iter.
- Assertion F proves the bare class **cannot** be fenced without unacceptable over-match, measured — that
  refutes this iter's hypothesis and is worth more than a hand-repair that leaves the class invisible.
