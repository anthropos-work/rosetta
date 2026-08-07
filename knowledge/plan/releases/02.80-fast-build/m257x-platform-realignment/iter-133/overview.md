---
iter: 133
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-07
---

# iter-133 — the fence prints the true module set; the prose still contradicts it

## Active strategy reference

**No successor strategy is authorable** — `TOK-08`'s sealed refutation branch bars one, and the
milestone runs under the user's direct brief (as iters 121–132 did). No `TOK-09` is created.

## Step 0 — re-survey

Route `FIX-M257x-iter131-my-three` is live and unabsorbed: **P5** (`architecture_overview.md:83` still
lists `ai` among the imported modules), **P7** (`platform-migration-status.md` §1 never gained the
`library-unimported` row while assertion C says "nine"), **P19** (*"all three sites are the literal
`curl`"*, false of one of its three). iter-132 touched none of them.

**Substitution under the same brief, recorded per Step 0:** the route names **three** sites. Measuring
the width first (§5 rule 57) shows **P5 and P6 are one predicate with seven more anchors**, so this
iter targets **the predicate**, not the two anchors the reading happened to sample. That is iter-131's
own lesson 1 — *a fence that prints the right answer does not correct the prose beside it* — and its
lesson 2 — *a correction reaches where it was written and nowhere else unless somebody sweeps.*

## Cluster / target identified

**The private-Go-module set a stack builds.** Assertion G (added iter-130) prints the true set on every
run; the prose disagrees with it in multiple files, in **two different wrong directions**:

- sites that still include **`ai`** (folded in-tree at `1e457fa70`, 2026-08-04);
- sites that say **three — colony, proto, taxonomy** — which drops **`analytics-go`** and **`storage`**,
  both *direct* requires of `app`.

**Ground truth, measured this iter, not quoted:**

| repo | ref | org-private requires |
|---|---|---|
| `app` | `ad9f3c498` | `go.mod:14-18` — **analytics-go, colony, proto, storage, taxonomy** (five, all direct) |
| `sentinel` | `f2c4619` | `colony`, `proto` direct; **`taxonomy` `// indirect`** |

Union a stack builds: **five**. Not three, and not four-including-`ai`.

Plus the two singletons the route named: **P7** (a vocabulary change that reached the checker and not
the definition) and **P19** (an over-claim inside the sentence whose purpose was to make a citation
robust).

## Hypothesis

Every site is repaired against `app` `ad9f3c498` / `sentinel` `f2c4619`, and the corpus stops
contradicting a fence it already ships.

## Expected lift

**No `N` reading; no `N` movement claimed** (§9 UNMEASURED rule, guard-rail 1). Success = the predicate
swept with its width stated, P7 and P19 closed, guard family no worse than at open.

## Phase plan

1. Width, two independent searches (done — 16 + 22 candidate lines, triaged below in `progress.md`).
2. Measure ground truth at source (done — the table above).
3. Repair the predicate corpus-wide; then P7; then P19.
4. Guard family + scoped tests.
5. Close.

## Escalation conditions

- If a repaired count disagrees with assertion G's printed set → stop; the fence and the measurement
  cannot both be right, and that is a user-visible contradiction, not a wording choice.

## Acceptable close-no-lift outcomes

- If the width search shows the predicate was already swept, close no-lift with the enumeration as the
  falsification. (It does not — see `progress.md` §2.)
