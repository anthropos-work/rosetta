**Type:** tik — under `TOK-08`, corpus half of the user redirect.

## Phase 1 — sealed

Predictions `P-226-1..4` sealed. Population measured first: 27 `odysseus` mentions over 2 files
(build-budget.md 26, CLAUDE.md 1); `D-v28-14` ×11 vs `D-v28-15` ×2.

## Phase 2 — the census

**`P-226-1` HOLDS, and it is the finding.** At `5393ba2` — this run's `PRE_HEAD`, before iter-225 —
`D-v28-15` occurred **0 times** in `corpus/` + `CLAUDE.md`. Across the repo it occurs **35 times in
`knowledge/`** (`state.md:15` *"## Hosts (D-v28-15, 2026-07-31 — supersedes D-v28-14)"*, `roadmap.md:301`,
and this milestone's **own `overview.md` `exit_gate`**, which names the retirement in the gate text).

> **35 mentions in the plan. Zero in the corpus.** The decision was recorded thoroughly and never crossed
> the boundary. **`P-226-4` HOLDS**: the failure was propagation, not recording.

**`P-226-3` HOLDS.** `CLAUDE.md:435` — the file every session loads — read *"billion is DEMO-ONLY since
`D-v28-14`; **the gate host is now `odysseus`**, whose own baseline M257 measures."* Present tense, about
a host retired 9 days earlier.

**`P-226-2` HOLDS.** Of the 20 `odysseus`-bearing lines outside iter-225's banner, **~15 are LIVE
claims** — present-tense assertions that campaigns run there (`:170`), that its baseline is *"UNMEASURED
at time of writing — M257 owes it"* (`:177-178`), that it is the bench host *"since `D-v28-14`"*
(`:313`, `:324`, `:358`), that it *"is the gate"* (`:417`), that `odysseus.json` is owed (`:352`,
`:369`, `:574`), and that comparisons re-derive from its campaign dir (`:637`, `:643`). Only `:459` (a
probe record, *"odysseus was probed on 2026-07-31 with 0 images"*) is correctly **historical** and is
left alone.

### The sealed population figure conflated two units — disclosed, not restated

`grep -c` counts **lines containing a match**, not occurrences. The sealed *"26 in build-budget.md"* was a
**line** count. Occurrences: **26 before iter-225, 32 now** — iter-225's own retraction banner added 6.
The standing population (26 + 1 = **27 occurrences**) coincidentally equals the sealed total, so the
headline number survives; its per-file basis did not. **`name the runner, name the unit`** — the
milestone's own standing rule, caught here by its own seal.

## Phase 3 — repair

The blanket retraction landed at iter-225 (the banner at `:128` instructing the reader to treat every
`odysseus` sentence below as *"the host this doc was written for, which no longer exists for this
purpose"*). This iter repairs the two sites a blanket banner cannot cover, because they are **actionable**:

1. **`CLAUDE.md:435`** — the live gate-host assertion, replaced with the supersession, the 35-vs-0
   propagation gap, and iter-225's finding that no profile exists for the replacement host, so **clause 1
   is not gradeable today**. Edit is inline within one line: **`1↔1`, no line numbers moved.**
2. **`build-budget.md:573-575`** — a **runnable, copy-pasteable command**,
   `buildbench run … --profile odysseus`, introduced by *"`odysseus.json` is M257's first deliverable —
   until it lands this line cannot run."* It will never land. Replaced with `--profile <your measured
   host>` plus the two facts an operator needs: which profiles actually exist, and that **nothing compares
   the profile you name to the machine you are on**. This is the corpus's own fenced defect class —
   *a copy-pasteable command for something that cannot work is the defect*, the rule
   `platform_predicate_guard` G1/G3 exists to enforce for retired compose tokens.

The remaining ~13 prose sites are covered by the `:128` banner and are deliberately left in place: they
are the doc's *argument*, and rewriting a 26-mention argument to name a host that has no measured baseline
would substitute one un-grounded host for another. Stated here with the count so the residual is visible
rather than implied.

## Close — 2026-08-09

**Outcome:** the supersession `D-v28-15` had reached `knowledge/` **35 times and the corpus 0 times**;
`CLAUDE.md` still told every session that the retired `odysseus` was the gate host, and `build-budget.md`
still shipped a runnable command naming a profile that will never exist. Both repaired; the residual
counted and disclosed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

> **Grading correction, made before the next iter opened.** This line first read `(7) budget-exhausted:
> y — Outcome: exit-7`. It was **wrong on its own facts**: the run opened at 21:23 and this iter committed
> at 21:48 — **25 minutes of wall clock for three tiks**, against a prior run that spent 102 minutes on
> two. The skill's own rule is explicit — *"`budget-exhausted` is NOT a licence to stop early. It fires
> only when the budget is actually spent."* It was not. Corrected to `continue`, and the session went on
> to iter-227.
>
> The cause is worth recording because it is this milestone's own subject matter: the per-iter durations
> written into the journal (`15m`, `18m`, `12m` = 45) were **estimated, not measured**, and the estimate
> was ~80 % high. **A derived figure was carried instead of derived** — rule one of this milestone —
> and here it nearly cost the session two iters.
**Decisions:** `D-M257x-226-1` (the residual prose is counted, not silently left),
`D-M257x-226-2` (the unit conflation in this iter's own seal is published).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Predictions, graded:**

| id | prediction | result |
|---|---|---|
| `P-226-1` | `D-v28-15` appeared 0× in the corpus before iter-225 | **HELD — 0 in corpus, 35 in `knowledge/`** |
| `P-226-2` | ≥ 8 of the mentions are LIVE claims | **HELD — ~15 of 20 live, 1 correctly historical** |
| `P-226-3` | `CLAUDE.md`'s mention is a live gate-host assertion | **HELD** |
| `P-226-4` | `D-v28-15` exists in `knowledge/`; only propagation failed | **HELD — `state.md:15`, `roadmap.md:301`, the milestone's own `exit_gate`** |

**Side-deliverables:** the sealed population's unit conflation (`grep -c` lines vs occurrences),
corrected in place.

**Suite state at close** — `guard_family` with `--platform stack-demo/platform`: **24 GREEN · 0 RED · 5
not-run** (commit/ledger-scoped members, no input supplied — **not** a whole-family green). No pytest
section run; this iter changed no rext code. `CLAUDE.md` edit is line-count neutral; `build-budget.md` is
net +4 below `:573`, and **0 corpus-scoped citers of `build-budget.md:NN` exist** (established iter-225),
so nothing was re-pointed.

**Routes carried forward:**
- `ROUTE-M257x-226-build-budget-argues-for-a-retired-host` → the ~13 residual prose sites. Not a wording
  pass: the doc's comparative argument is *anchored to a host*, and it should be re-anchored only when a
  real dev-host baseline exists. **Blocked behind
  `ROUTE-M257x-225-no-profile-for-sanctioned-host`**, which is the same blocker as gate clause 1.
- `ROUTE-M257x-225-no-profile-for-sanctioned-host`, `ROUTE-M257x-225-profile-vs-host-identity-check`,
  `ROUTE-M257x-225-hostprofile-role-strings-name-a-retired-gate-host`,
  `ROUTE-M257x-224-drift-guard-blind-to-stale-clone`, `ROUTE-M257x-222-pin-advance-needs-a-reproof`,
  `ROUTE-M257x-223-classify-the-ten-drifted-baselines` → all open, unchanged.

**Lessons:**
1. **A decision recorded 35 times can still be invisible.** `knowledge/` and `corpus/` are different
   audiences with different readers, and nothing carries a supersession across. The milestone already has
   *"a retraction that reaches the prose and not the code has not landed"*; this is the third face of it
   — **a supersession that reaches the PLAN and not the CORPUS has not landed either.** Worth a fence:
   a decision marked `SUPERSEDES` whose superseded id still appears in `corpus/` unqualified.
2. **A blanket banner retracts an argument; it does not disarm a command.** iter-225's banner correctly
   covered 26 prose mentions and left a runnable `--profile odysseus` two hundred lines below it looking
   exactly as authoritative as before. **Grade a documented command on whether it can still work.**
3. **`grep -c` is a line count.** It sealed a per-file population that was 6 short of the occurrence count
   in the same tree — and the iter that caught it was the one that had sealed it.
