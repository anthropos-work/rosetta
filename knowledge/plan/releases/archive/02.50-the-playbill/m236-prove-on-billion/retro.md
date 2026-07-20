# M236 "prove-on-billion" — Retro

**Shape:** iterative · **Outcome:** gate MET cold on `billion` · **Iters:** 10 (1 bootstrap tok + 9 tiks, one day)
**Merge:** HELD — the deferral audit returned RED and the blocking item needs a user fate (CLOSE-D2).

## Summary

The final v2.5 milestone re-proved the whole content-vantage feature live on the `billion` Tailscale VM:
**29/29** landable (session × action) pairs rendering real content for both player and manager vantages,
**65** academy course cards with 0 Draft chips, hero p95 click→ACCESS **1.22 s / 1.51 s** against a 5 s
budget — all reproduced on a cold reset-to-seed with no intervention and **0 platform-repo edits**.

It also authored, from scratch and against a live render, the content-stories seat-login sweep that M235
proved could not be written blind.

## The headline: the score was five wrong test assertions to one real product bug

This is the finding worth carrying out of the milestone, and it is not a comment on sloppiness — it is a
structural observation about proving that software works.

Of the defects that cost an iter, the **majority were the test being wrong**, not the product:

- **iter-05** — a manager test that *asserted the defective contract*. That is precisely why the bug
  shipped: the test encoded the wrong URL shape, so the code matching it was "correct".
- **iter-06** — the interview manager view was a **false FAIL**. It had been rendering perfectly all along
  and was being graded against the wrong shape.
- **iter-07** — the skill-path page was graded as a scored simulation, and its sibling was scoring a
  **false PASS** off a definition-only header.
- **iter-08** — the academy CTA named a route that does not exist, and *the unit test required that prefix
  and so defended it*. The test was the thing keeping the bug in place.

Then the final harden found **three more checks passing against a broken subject**:

- an **aggregator reporting success on an empty run** — 0/0 is also arithmetically 100%;
- the **whole e2e suite passing by collecting 0 tests** after a module-scope throw, which silently took
  **61 tests offline for 8 iters**;
- a **grader with no negative tests at all** — nothing established it could ever report failure.

And at the close, a fourth: M236's **own regression test** for the manager-route defect was a
**self-consistent tautology**. It asserted the projected manager path ended in `slot.MembershipID` — but
`slot.MembershipID` was produced by the same expression under test. It passed regardless of whether the key
matched what the users seeder actually wrote. *The fix for the defect contained a test that could not detect
the defect's return.*

**The rule, stated once:** *ask of every layer that reports a number — what does it print when nothing
happened?* A layer that cannot distinguish "everything passed" from "nothing ran" is not a gate. It was
found at five layers here (collection, aggregation, grading, the coverage reading, and the regression test
itself), and the collection layer is the one that can hide all the others at once.

Backfilled into `coverage-protocol.md` (two new subsections), `latency-budget.md`, and
`content-stories-spec.md`; made **executable** at the close via a collection-integrity guard with floors,
a golden-value pin on the membership derivation, and mutation verification on both.

## The second lesson: a lesson written in one file does not propagate to its siblings

This recurred so often it is really the same finding at a different scale:

- The **non-integer-`N` guard** was added to 2 of 4 runners during the milestone; the other two carried the
  identical hazard and were fixed only at the close.
- The **`networkidle` rule** was *already written down* in `latency-budget.md` — and the new sweep inherited
  the default anyway, producing a 180 s "hang" on a page that painted in ~1 s.
- The **doc corrections** followed the same pattern: each iter fixed the doc it touched, leaving three other
  docs asserting the very claim that had just been refuted — including the one that produced the inflated
  denominator.
- The **membership key** was a bare literal at **9 sites**, one of which writes the row and eight of which
  merely hope to match it.

Prose does not propagate. Only a shared definition or an executable fence does. Every instance is now one or
the other.

## Incidents this cycle

- **P2 — the green gate failed OPEN west of UTC** (iter-09). The age check parsed a UTC `ts` as local time
  on BSD, so a 121-second-old verdict aged as 7321 seconds. West of UTC the same arithmetic reads a **stale**
  verdict as **fresh** — the exact hazard the check exists to prevent, inverted, for half the world. Fixed +
  regression-tested with a full timezone sweep and a shipped-line pin.
- **P2 — 61 tests offline for 8 iters** (found in the final harden). See above.
- **P3 — a false PASS in the gate reading itself** (iter-07). A skill-path manager pair scored green off
  chrome served by a different query than the one that had failed.
- **No regressions.** Flake 0 (milestone-owned).

## What went well

- **The gate denominator was corrected rather than defended.** Driving the surface showed 2 pairs were not
  provable; the target *shrank* 31 → 29, was argued inline with product-source evidence, and the 31 was
  struck through rather than rewritten. It would have been easy — and wrong — to report 31/31.
- **Authoring the harness against a live render, not blind.** M235's USER-BLOCKER-02 insisted on this. It was
  right: iter-04's calibration found 3 shapes, **2 of which a blind author would have gotten wrong in
  opposite directions**. Six shapes exist where the roadmap assumed one.
- **The Phase-0b audit caught a spec that could not be discharged as written**, before any work was done —
  an unprovable gate clause, an unmeasurable metric, a cited page-object that did not exist. Escalating to
  the user rather than working harder was the correct move.
- **Publish-before-prove.** iter-01 found `billion` structurally *could not obtain the feature under test*
  (0 of 13 tags on origin). Naming that as rung zero, and taking a deliberate 0-lift iter to fix it, avoided
  a milestone of measuring the wrong code.

## What didn't

- **Four docs carried the refuted claims to the close.** The iters corrected what they touched. The
  stale-adjacent doc is a systematic gap, not an oversight — worth a mechanical sweep at the point a claim
  is refuted, not at close.
- **Two cross-repo doc-truth guards were red and correct for three milestones**, and were read as noise. The
  cause is instructive: each guard's *own test* hardcoded the superseded fact, so the suite was red for a
  reason that looked like the guard being broken. **A test asserting a stale fact is indistinguishable from
  a real regression** — and it trained readers to ignore a working alarm.
- **The standing carry was briefed as 14 and measured as 19.** Five failures had no ledger entry anywhere.
  A carry that is copied forward rather than re-measured stops being a fact.
- **The harness was modified after the gate was measured** (CLOSE-D3). Unit-proven, not live-re-proven —
  disclosed rather than smoothed over, and routed to the release close.

## Carried forward

| Item | Fate | Destination |
|---|---|---|
| **14 pre-existing demo-stack failures (REPEAT — BLOCKING)** | **needs a USER fate** | v2.5 release close — re-baseline first |
| `ACADEMY-M236-iter08-public-catalog-twin` (anon `/library`+`/free` = 0 cards) | Fate 3 | v2.5 release close |
| `apps/web` client GraphQL on non-offset `:5050` | Fate 3 | v2.5 release close (batch with above) |
| M204 assign-WRITE declared TODO | Fate 3 | v2.5 release close |
| Live re-run of the close-modified harness (CLOSE-D3) | Fate 3 | v2.5 release close |

**Discharged:** M230 carry-forward cluster 1 (rendered-card count) and cluster 2 (clone re-anchor — which
had no closing entry anywhere until this close, and whose diagnosis was falsified).

## Metrics delta

| | M235 | M236 close |
|---|---|---|
| Go test funcs (whole rext) | 1974 | **1976** |
| Go suite | green | **2459 pass / 0 fail**, 6 modules / 58 pkgs |
| Python | — | **1391 pass / 2 fail / 8 skip** |
| stack-verify python | 132 | **141** |
| stack-core | 147 / 5 fail | **150 / 2 fail** (5 fixed) |
| M236 harness specs | 0 | **66** |
| Hero p95 click→ACCESS | 2.11 / 1.31 s (v2.3) | **1.22 / 1.51 s** (cold, billion) |
| Platform-repo edits | 0 | **0** |
| Flake | 0 | **0** |

Full artifact: [`metrics.json`](metrics.json).
