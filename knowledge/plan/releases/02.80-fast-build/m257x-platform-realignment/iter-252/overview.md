---
iter: 252
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-252 — close the class: `clone_drift_guard`, the last named existence-grader

**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — census the mechanical classes; stop
sampling them.

## Step 0 — Re-survey before targeting

`ROUTE-M257x-250-the-runtime-bucket-is-one-guard-wide` named four guards. Two are done this run —
`rext_path_guard` (iter-250, 1 instance) and `corpus_citation_guard` (iter-251, 21). Two remain:

- **`fence_command_guard.locate`** — iter-250 measured its whole reach (102 of 103 graded `cd`
  occurrences are workspace-resident) and **already fixed the disclosure half**: the verdict now states
  the tier split on every run. Its refusal buckets (`cd: workspace not provisioned on this host` ×44 on a
  clean tree) mean it does **not** convert an absent workspace into a corpus finding. Re-surveyed: the
  remaining gap there is `ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace`, which needs a
  **text-shaped classifier** — a new instrument, not a repair, and too large to bolt onto this iter.
- **`clone_drift_guard`** — untouched, unmeasured, and named by the route. This is the target.

Substrate confirmed by reading before targeting: `clone_drift_guard.py:147` `if not root.is_dir()`,
`:150` `if d.is_dir() and (d / ".git").exists()`, `:208` `if not gm.is_file()`, `:329`
`if not (repo_root / CORPUS_REL).is_dir()` — four filesystem predicates.

## Cluster / target identified

The last unmeasured member of a class this run has confirmed three times. `TOK-08`'s rule applies exactly:
**a reading samples; a fence censuses** — and the honest close of this route is a *measurement* of the
fourth guard, whichever way it comes out, not another repair assumed by analogy.

## Hypothesis

`clone_drift_guard`'s verdict is operator-dependent in the same way, because its subject is the clone set
and it locates it on the filesystem.

**The competing hypothesis, stated because it is the likelier one and this iter must not assume its
conclusion:** a guard whose *declared subject* IS the clone set is not making a claim about the corpus at
all when the clones are absent — it should REFUSE, and `:147`/`:150` may already be that refusal. In that
case the class ends at three and the route closes with a derived zero.

## Pre-registered numeric claims — sealed in this iter's FIRST commit

| # | claim | prediction |
|---|---|---|
| **PR-1** | on a `rosetta`-only checkout, `clone_drift_guard` **refuses** (exit 2 / could-not-check) rather than reporting corpus findings | **true** — the competing hypothesis is the likelier one |
| **PR-2** | its live and `rosetta`-only exit codes are **equal** | **false** — a refusal is a different code from a green, and that is correct behaviour, not a defect |
| **PR-3** | it emits ≥ 1 finding on the `rosetta`-only checkout that is false about the corpus | **false** |
| **PR-4** | the route `…-runtime-bucket-is-one-guard-wide` closes at **three** guards, not four | **true** |

**Direction check.** iter-249 was 1 of 5 because it guessed; 251 was 5 of 5 because it predicted a
mechanism it had already measured. Here the mechanism is **not** measured, so these are deliberately set
against my own thesis: three of the four predict the class does NOT extend. If they hold, the route closes
honestly; if they fail, there is a fourth instance and the repair is already designed.

## Phase plan

1. **Seal** as `probe(M257x/252)`.
2. **Run** `clone_drift_guard` on the live tree and on the iter-251 `rosetta`-only clone; record both exit
   codes and both verdict lines verbatim.
3. **Classify**: refusal (correct) vs finding-about-the-corpus (the defect).
4. **Repair or close**: if a defect, apply `D-M257x-251-1`; if a refusal, close the route at three with
   the number, and add a regression test pinning the refusal so it cannot silently become a finding.
5. **Grade** PR-1…PR-4 and close.

## Escalation conditions

- If the guard neither refuses nor reports — e.g. it exits 0 over an empty clone set — that is the
  vacuity class (`§9`), worse than either branch, and it is repaired in this iter.

## Acceptable close-no-lift outcomes

- PR-1/PR-3/PR-4 hold: the class ends at three, the route closes with a measurement instead of an
  assumption, and the deliverable is the pinned refusal. That is a complete iter with no repair in it.
