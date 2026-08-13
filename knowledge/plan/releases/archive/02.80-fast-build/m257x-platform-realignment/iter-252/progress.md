**Type:** tik · **Active strategy:** `TOK-08` — census the mechanical classes; stop sampling them.

## Open — 2026-08-10

Sealed PR-1…PR-4 before running the guard (`8ab151d`). Three of the four predict the class does **not**
extend — set deliberately against this run's own thesis.

## What happened

### The measurement

`clone_drift_guard`, run on both trees, verbatim:

```
LIVE            → clone-drift-guard: OK — every cited clone's HEAD is a commit the corpus cites,
                  and all 4 go.mod-cited pin(s) match.            rc=0
ROSETTA-ONLY    → clone-drift-guard: CANNOT RUN — no clone under stack-demo/ or stack-dev/
                                                                   rc=2
```

**Zero findings on the clean tree.** The guard does not convert an absent clone set into a claim about the
corpus — it refuses, names where it looked, and exits 2. **The class ends at three.**

That is the right behaviour and the reason is structural, not lucky: **this guard's declared subject IS the
clone set.** The other three had a corpus subject and reached into the operator's tree to answer it; this
one is asking about the clones themselves, so their absence is a missing instrument, not a missing file.
It is the same distinction `D-M257x-248-3` draws for tests, one level up.

### Not the fourth instance either

`fence_command_guard` was the route's other open name, and it was re-surveyed rather than assumed:
its refusal buckets (`cd: workspace not provisioned on this host` ×44 on a clean tree) mean it, too, does
not manufacture corpus findings from an absent workspace, and iter-250 already shipped its reach
disclosure. What is left there is a different job — a **text-shaped** reach classifier, tracked as
`ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace` — not this class.

### The deliverable: the refusal is now pinned

No repair was needed, so the iter's product is that the measured behaviour cannot silently regress.
Two tests in `tests/test_clone_drift_guard.py`:

- a tree with no clone set **refuses** (exit 2, `CANNOT RUN`) and does not report;
- the refusal **names where it looked** (`stack-demo`, `stack-dev`) — a refusal that does not say what was
  missing sends the reader hunting.

Both fixture bugs they surfaced were mine and were caught before commit: `main()` takes argv **without**
the program name, and `clones()` returns a mapping, not a list.

## Pre-registration grading (sealed at `8ab151d`)

| # | claim | prediction | outcome |
|---|---|---|---|
| **PR-1** | refuses on a `rosetta`-only checkout rather than reporting | true | **HELD — exit 2, `CANNOT RUN`** |
| **PR-2** | live and clean exit codes are equal | **false** | **HELD — 0 vs 2**, and that inequality is the correct behaviour |
| **PR-3** | ≥ 1 finding on the clean tree that is false about the corpus | **false** | **HELD — zero findings** |
| **PR-4** | the route closes at **three** guards, not four | true | **HELD** |

**4 of 4** — and the ones that mattered were the three predicting against my own thesis. Run trend:
**1/5 → 3/5 → 5/5 → 4/4.**

## Close — 2026-08-10

**Outcome:** `ROUTE-M257x-250-the-runtime-bucket-is-one-guard-wide` **CLOSED at three, by measurement
rather than by analogy.** `clone_drift_guard` is not a fourth instance: on a `rosetta`-only clone it exits
**2** with `CANNOT RUN — no clone under stack-demo/ or stack-dev/` and reports **zero** findings, because
its declared subject *is* the clone set. The distinction that decides the class: a guard with a **corpus**
subject that reaches into the operator's tree to answer is broken; a guard whose **subject** is the
operator's tree is not, and must refuse. The refusal is now pinned by two regression tests so it cannot
quietly become a finding.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: y — Outcome: exit-7
**Decisions:** `D-M257x-252-1` (the class boundary is SUBJECT, not mechanism — reaching into the
operator's tree is only a defect when the claim is about the corpus) · `D-M257x-252-2` (a route is closed
by measuring its last member, never by analogy from the first three) · `D-M257x-252-3` (when the
measurement finds nothing to repair, the deliverable is the pin).

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
`test_clone_drift_guard.py` **35 passed (32.66 s)**. Guard family (`--platform stack-demo/platform`, repo
root): **29 GREEN / 0 RED / 0 could-not-check / 5 not-run** — unchanged across all four iters of this run.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-250-the-runtime-bucket-is-one-guard-wide` → **CLOSED** (3 instances, 4th measured and
  refuted).
- `ROUTE-M257x-249-fresh-checkout-hostile-tests` → **open, and the largest thing this run found**: 23
  failures across 13 files, all authored inside this milestone, still being manufactured. `TOK-08` says
  this should become a fence.
- `ROUTE-M257x-251-two-trees-both-called-a-fresh-checkout` · `ROUTE-M257x-249-a-reading-must-name-its-failures` ·
  `ROUTE-M257x-249-anchor-offset-has-three-populations` ·
  `ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace` → open.
- Still open, untouched: `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` ·
  `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-hardcoded-vs-settable` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves`.

**Lessons:**
1. **The class boundary is the SUBJECT, not the mechanism.** All four guards call `is_dir()` / `exists()`
   on the operator's tree. Three were defective and one is correct, and the difference is entirely whether
   the claim being made is about the corpus or about the clones.
2. **Close a route by measuring its last member.** Three confirmed instances made the fourth feel certain.
   It was wrong, and one command settled it — cheaper than the repair it would have justified.
3. **When there is nothing to repair, the deliverable is the pin.** Otherwise the measurement is a memory,
   and the next person re-derives it — which is the waste this milestone exists to end.
