# iter-213 — the backlog registry reads a route's NAME as its VERDICT

**Type:** tik — under [`TOK-08`](../decisions.md).

See [`overview.md`](overview.md) for S1–S5 + the stop condition, sealed at `291e079` before any repair.

## What was measured

**S1 — CONFIRMED.** **8 of 367** route ids carrying a disposition match a disposition regex **in their
own name**: five match `REOPEN_RE`, three match `PARTIAL_RE` via `\barm\b` / `\bhalf\b`.

**S2 — CONFIRMED.** Of **1,447** segments carrying at least one route id, **14** change verdict once the
id text is removed, across **3** distinct ids.

**S3 — CONFIRMED. `violations()` is 0 before and 0 after.** LATENT, not a live false-GREEN, and not
reported as one.

**S4 — FALSIFIED, by this iter's own staged control, on its plain-slug leg.** See
[`D-M257x-213-2`](decisions.md). A bare re-listing grades `other` for **both** slugs and never triggered
the rule. The rule fires on `open` after `closed`, so **13 of the 14 changed segments move the census
DISPLAY only** and exactly **1** could ever have moved the rule:
`FIX-M257x-iter177-ledger-carries-a-retracted-retraction` at iter-178 — *"unchanged; owner = the next
harden pass"* — a route carried **unchanged** and booked **re-opened** because its slug says
`retracted-retraction`. It had never been closed, so nothing was in fact suppressed.

**S5 — CONFIRMED against this iter's own first probe.** 40/26 → **14/3**, ~2.9× overstated, because the
first probe re-classified the 240-char truncation stored in the event tuple instead of the segment.

## What was shipped

- **`classify()` strips route ids before grading** — one line, making the function's own docstring true.
- **Five arms** (`ARouteIsNotGradedByItsOwnNAME`): the staged suppression control (two routes identical
  but for the slug, both now RED); its **anti-vacuity twin** proving the only difference is the name,
  across four verdict texts and both poisoned regex families; the `PARTIAL_RE` half (which downgrades a
  **closure** — the more dangerous direction, and exercised by no live segment); the **over-strip**
  control; and the executable record of S4's falsification, paired with the control it corrected.

## Close — 2026-08-09

**Outcome:** the milestone's backlog registry graded a route by its own NAME — `REOPEN_RE` and
`PARTIAL_RE` match ordinary English (`retract`, `refut`, `supersede`, `arm`, `half`), and `classify()`
was handed the segment with the id still in it. **8 of 367** ids are so named; **14 of 1,447** segments
changed verdict once the ids were removed. The live exposure is **1 segment and it is latent**
(`violations()` 0 → 0) — a size this iter established by **falsifying its own pre-registered S4** with
its own first staged control.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-fifth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted, not felt: iters 212, 213 = two tiks this run against a cap of five** —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-213-1` … `D-M257x-213-3` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**, `stack-core` only —
**102 passed / 0 failed** across `test_route_disposition_guard`,
`test_harden_origin_route_visibility_m257x`, `test_fence_registry_population_m257x` and
`test_guard_family` (21 s); **35 passed** on the changed test module alone.
**Live guard, before and after:** `--census` over three milestones byte-identical apart from this
iter's own directory — M257x **421 route ids · 367 with a disposition · 55 closures · 3 ambiguous · 0
malformed · 0 contradictions** both sides. The pre-registered stop condition **did not fire**.
**RED-proof battery, mtime-mitigated (`§5` r77), restore sha-verified:** deleting the single
`ID_RE.sub` line takes **all five** new arms RED and leaves the other 30 green.
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. No
whole-section run — the tree was edited during the iter. No Go, no TypeScript. The four non-`stack-core`
Python sections were read at iter-208 and not re-read since.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter212-a-retraction-does-not-reach-the-code-that-acts-on-it` — **still open.** This
  iter worked one grain BELOW it (the registry's own grammar), not the route itself: nothing here
  connects a retracted route id to the artifacts still citing its claim.
- `SURVEY-M257x-iter213-a-route-id-is-english` — **NEW.** The repair removes ids from the *classified*
  text; it does not stop new ids being named after a verdict word. **8 of 367 already are**, and the
  slug is also read by humans, by `WELL_FORMED_RE`, and by every grep. A naming rule fenced at
  route-creation time is the durable form; sizing it needs its own iter.
- `SURVEY-M257x-iter211-A-and-B-still-spell-their-own-scope` — closed at iter-212.
- All routes from iters 207–210 and 212, unchanged, plus the standing queue.

**Lessons:**
- **A registry that reads its own subjects' NAMES is grading the label, not the record.** The guard's
  docstring already forbade this; only the code disagreed.
- **A control that goes RED on its CONTROL leg has found something.** The plain-slug leg failing is what
  falsified S4 and cut the claimed exposure from 14 to 1 — the arm was right and the hypothesis wrong.
- **Re-classifying a stored truncation is not re-classifying the input.** 2.9× overstatement, same
  class as iter-209's slugger: when the question is about the input, hold the machinery fixed.
- **`reopen` is an EXCUSE, `other` is silence, and they are not interchangeable** — which is why the
  transition table, not the total, is what says how big a classifier defect is.
