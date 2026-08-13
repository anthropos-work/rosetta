# iter-126 — the guard off exit 2, and the re-pin backlog enumerated

**Type:** tik · **Run 80, tik 3.** Both items routed forward at iter-125's close.

## 1. `platform_alignment_guard` — resolved, not silenced

**The guard was right and its vocabulary was wrong.** It returned exit 2 because the corpus had begun
citing `infrastructure` and `db-backup` — repos no stack clones — and it graded those citations
`unresolvable`, i.e. as **blind spots**. That is a correct verdict on a changed subject, and the subject
changed because iter-123 did the right thing.

Two conditions were sharing one bucket:

| | what it is | whose defect |
|---|---|---|
| **unresolvable** | a path the guard cannot find and cannot account for — typo, unqualified path, non-existent repo | **the citation's** |
| **unclonable** | a repo **the map itself documents**, simply not in any clone set | **the substrate's** |

The 7 split across **both** halves of the directive's own menu, which is why neither alone would have
closed it:

| head | × | resolution |
|---|---|---|
| `terraform` | 3 | `cms`'s row cited `terraform/production/services.tf:64/85/88` **unqualified** — ambiguous across 12 terraform-bearing repos. **Corpus fix: qualified.** |
| *(bare, no preceding path)* | 2 | `graphql-wundergraph`'s row — **iter-124's own bare pins**. **Corpus fix: qualified.** |
| `infrastructure` | 1 | documented by the map, in no clone set. **Guard fix: unclonable + disclosed.** |
| `db-backup` | 1 | same. |

**The disclosure rides IN the verdict sentence**, because `guard_family.run_one` reports `lines[-1]` for a
green member — a qualifier anywhere else is invisible in the family view, which is exactly how an
over-claiming green survives:

```
platform_alignment_guard: OK OVER ITS REACH — platform-migration-status.md and repos.yml agree in
both directions, and 10 citation(s) into 2 repo(s) the map documents but no stack clones
(db-backup, infrastructure) were NOT checked.
```

### The anti-vacuity gate is the decision, not a detail

A head is `unclonable` **only if the map's own services or census table names that repo**. Without it, a
typo would launder itself into the excused bucket and this change would be **a silencer wearing a
disclosure's clothes**. Three controls, each asserting it applied, plus a kill from the guard's side:

| control | proves |
|---|---|
| anti-vacuity — undocumented head | stays `unresolvable`; gap still recorded on the refusal stream |
| mutation — delete the documenting row | the excuse goes with it |
| negative control — nothing unclonable | still gets the **unqualified** OK, so the qualifier carries information |
| **guard-side kill** — remove `head in mapped_repos` | **exactly ONE test goes red.** A named kill, not a smear |

### And the widening immediately bought a finding the exit 2 was masking

With assertion F able to run to completion, `messenger`'s row was caught citing `:622` / `:664` **bound to
`messenger/terraform/main.tf`, a 121-line file**. They are `infrastructure`'s `services.tf`. Repaired.

> **A guard stuck at exit 2 is not a guard being careful. It is a guard that is not looking.**

## 2. Priority 4 — the backlog enumerated, and the substrate trap it walked into first

iter-123 measured *"at most 13, not 74"* **and did not list them** — which is exactly why the item was
still open a run later. **A count with no list cannot be closed.** The list is
[`repin-backlog.md`](repin-backlog.md).

> ### The first re-derivation was wrong, 3.1×, and the cause is this milestone's own rule
>
> | corpus read at | NO-SHA-IN-BLOCK |
> |---|---|
> | **HEAD** (~30 commits after the census) | **22** |
> | **`afe58ac`** — the ref iter-123 named | **7** — reproduces iter-123 exactly |
>
> The backlog's corpus sites are **line pins taken at the census's ref**. `D-M257x-122-4` established
> that a stale substrate *fabricates* defects for platform clones; **it is not about platform clones, it
> is about pins.** Every one of those 15 phantom members would have been an edit to a citation that was
> never broken.

**Disposition of the 7 — 6 closed, 1 refused with a reason:**

| | count | |
|---|---|---|
| **re-pinned** `@ app 2035f9a` | 3 | each **verified at the ref**, not assumed |
| **path-qualified + re-pinned** `@ rext 63ce41a` | 2 | bare basenames; `D-M257x-122-5` refuses proximity resolution |
| **pin REMOVED, not re-pinned** | 1 | `B11-020` — see below |
| **stays, with the reason** | 1 | `B01-021` — an **intra-corpus** citation, where a sha does not apply and two fences already grade the anchor on every run. A **false member** of the class |
| **retired** | **0** | no anchor here has a third-generation history; iter-115's precedent does not fire |

### `B11-020` — qualifying it made it worse, and the fence caught that in the same session

A bare `README.md:21` in a sentence about the shared **`ai`** library. Qualified, the resolver bound it to
**`studio-desk` @ `41ee357`** — an unrelated repo — and landed on a **blank line**;
`anchor_construct_guard` went RED **on the qualification**. `ai` is a private Go module **no stack
clones**, so no `file:line` into it is verifiable from here. **Pin removed; the document named in prose.**

**Two costs recorded rather than tidied** (`D-M257x-126-4`):

1. **A manufactured hedge was written and withdrawn** — the first correction called the anchor
   *"not verifiable from here"* **while leaving the unverifiable pin in place.** The directive forbids
   manufacturing hedges for measurable facts; **the mirror binds too — do not keep a pin you have just
   said nobody can check.**
2. **The retraction had to be rewritten in the fence's vocabulary** (§8, iter-98): explaining the removal
   *by quoting the removed pin in backticks* kept the guard RED, correctly. **A retraction written in the
   vocabulary the fence enumerates is indistinguishable from the claim it retracts.**

## 3. Reach, with denominators

| statement | number | denominator |
|---|---|---|
| unresolvable citations in the map | **7 → 0** | assertion F's citation set |
| unclonable, disclosed | **10** | same — *disclosed, never absorbed* |
| backlog rows enumerated | **89** | 68 `UNRESOLVABLE` + 17 `PARTIAL` + 4 `DOES-NOT-SUPPORT` |
| no-sha class, at the census ref | **7 of 89** | reproduces iter-123 |
| of those, closed | **6** | 3 re-pinned · 2 qualified+pinned · 1 removed |
| refused with a reason | **1** | intra-corpus, a false member |
| **still open of the ≤13** | **6** | the `PIN-DOES-NOT-RESOLVE` class — routed, not attempted |
| guard tests | **51 / 51** | `tests/test_platform_alignment_guard.py` |

## 4. Guards

**22 members · 18 GREEN · 0 RED · 0 could-not-check · 4 not-run** — quoted from the runner's own summary
line (`guard-family: 18 GREEN · 0 RED · 0 could-not-check · 4 not-run`). **The could-not-check is gone.**
**Still not a whole-family green** — 4 members are commit-/input-scoped and were not supplied their
`--range`/`--ledger`, and the family exits **2** to say so. Invocation:

```
cd .agentspace/rosetta-extensions/stack-core
/usr/bin/python3 guard_family.py --repo-root <rosetta> --platform <rosetta>/stack-demo/platform
```

## 5. Routes carried forward

- `FIX-M257x-iter126-pin-does-not-resolve-6` — the other 6 of the ≤13. Each needs a full-history clone
  read **at the block's own named sha**; the procedure is in `repin-backlog.md`. A third line of
  investigation in this iter, so routed rather than opened.
- `FIX-M257x-iter126-nosha-class-excludes-intra-corpus` — narrow the class definition so intra-corpus
  anchors stop being counted as re-pin candidates.

## Close — 2026-08-07

**Outcome:** `platform_alignment_guard` is **off exit 2 with its limitation disclosed in the verdict
sentence**, not silenced — `unresolvable 7 → 0`, `unclonable 10`, backed by three controls plus a
guard-side kill that turns **exactly one** test red; and the widening caught a real mis-bound citation the
exit 2 had been masking. The re-pin backlog is **enumerated** (89 rows → 7 in the no-sha class at the
census ref, reproducing iter-123), **6 closed and 1 refused with its reason**. **No `N` movement is
claimed and no reading was taken.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — 4 of 5, unchanged — (2) triggered-tok: n (**successor strategy
FORBIDDEN by `TOK-08`'s sealed rule; running under the user's directed scope**) — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-126-1` … `D-M257x-126-5` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none — both lines were planned scope, named at iter-125's close.
**Lessons:** **A guard stuck at "cannot check" is not a careful guard; it is a guard that is not
looking.** The exit 2 was masking a genuine mis-bound citation in the same table. When a fence goes
could-not-check on a *changed subject*, the fix is to give the new condition **its own name and its own
disclosure** — never to widen the excuse without an anti-vacuity gate, and never to leave it blind
because blindness feels conservative. → `platform-alignment.md` §5 **rule 56**.
