---
milestone: M257x
iter: 59
iteration_type: tok
tok_flavor: triggered
status: closed-no-lift
created: 2026-08-04
active_strategy: TOK-05 (authored by this iter; supersedes TOK-04)
target_clause: "n/a — strategy revision"
refs:
  platform: 0dab54dfac6beacdef54a671e2500d3940fd7329   # origin/main, re-fetched at open (P3) — LEVEL, no re-point owed
  platform_source: stack-demo/platform
  app: v1.366.0 (b948604f)                            # the committed pin, advanced by iter-58
  rext_head: ab81527ae2ebfe4406bc4f1048f6c42056cd90d3 # main, clean
  rext_pin: fast-build-m257x-iter-58                  # .agentspace/rext.tag — level with HEAD's tag
  rosetta: 0f9cd87b1bba2a00b0402c54fd7a73753278a071   # at open
  instrument_platform_alignment_guard_sha256: 2a7862c4bc0bdd44be110d01e9a8cb25977ceb82e04e5ff17838b00ff5e3d232
  stack_core_baseline: 1F/610 (the 1 is the perishable iter-48 answer-key fixture)
  taken_at: 2026-08-04T08:50:46Z
---

# iter-59 — TOK-05: stop repairing claims; fence the predicates under them

**Type:** tok (triggered — by a **direct user directive**, not by the 3-no-prog streak)

## Why this is a tok, and which trigger fired

**The 3-no-prog streak does NOT apply and was checked before this was written.** iters 56, 57 and 58 all
closed `closed-fixed` and each moved a gate clause (1+2 restored, 3 met, 1+2 re-proven at advanced pins).
By the skill's Phase 0 rule 2 the next iter would have been a tik.

**The trigger is the user's own directive**, and it carries new evidence the milestone did not have: the
platform developer's PR shows what he is doing — skiller, cms, graphql, jobsim and skillpath merged into
`app`; storage and messenger **soon but not yet**. The instruction was to use it to scope better what is
needed, done already, or to change: *pause the tik loop, make a tok to reorganize, continue with fresh
insights.* This is the same shape as TOK-04, which was also fired by something other than the streak (the
milestone's `re_scope_trigger` plus a user ruling) and was recorded as `triggered`. Same precedent, same
record shape.

## Step 0 — Re-survey before authoring (mandatory), and it corrected an inherited number

A tok's Step 0 is *confirm the state the revision is premised on*. **INPUT 3 of this run's briefing is that
three inherited numbers failed in a single iteration, one of them the orchestrator's own** — so every
load-bearing denominator below was re-derived from platform artifacts at this open, not inherited.

| premise | inherited | re-derived at open | verdict |
|---|---|---|---|
| platform ref | `0dab54d` | `0dab54d`, and `origin/main` is level (re-fetched) | ✅ P3 satisfied, no re-point owed |
| `PROFILE ?=` default | `core` | `Makefile:10` = `PROFILE ?= core` | ✅ |
| services declaring a `graphql` profile | 0 | **0** (7 `profiles:` lines total, none `graphql`) | ✅ |
| compose services | 8 | **10** — `docker-compose.yml:1` has an `include: [common.yml]`, which adds `postgresql` + `redis` | ⚠️ **corrected** |
| `repos.yml` entries | 6 | **6** (`app sentinel storage messenger next-web-app studio-desk`) | ✅ |
| `migrations: true` repos | 1 | **1** (`app`) | ✅ |
| published port mappings | 13 | **13** | ✅ |
| `*_RPC_ADDR` values | 4, all `backend:8083` | **4**, all `http://backend:8083` | ✅ |
| corpus files asserting a `graphql` profile | ~17 | **17 files / 30 occurrences** at the tight predicate | ✅ |
| `main.go:N` citations in the corpus | 23 | **23** | ✅ |
| `stack-core` baseline | 1F/610 | **1F/610**, `Ran 610 tests`, the 1 = the iter-48 perishable fixture | ✅ |

### The one correction, and it makes the headline defect worse rather than smaller

The briefing states *"`make up PROFILE=graphql` exits 0 and starts nothing."* **Measured, it starts THREE.**

Compose selects every service that declares **no** `profiles:` key regardless of the `--profile` flag, and
after resolving the `include:` those are `postgresql`, `redis` and `sentinel`. So:

```
DERIVED legal profile set (from compose, 8):
  all · backend · core · customerio-sync · frontend · messenger · storage-legacy · studio-desk

PROFILE=core     selects 5 : postgresql redis sentinel backend gotenberg
PROFILE=graphql  selects 3 : postgresql redis sentinel      <- 'graphql' declared by NO service
PROFILE=cms      selects 3 : postgresql redis sentinel      <- 'cms'     declared by NO service
PROFILE=storage  selects 3 : postgresql redis sentinel      <- 'storage' declared by NO service
```

**"Starts nothing" would be an honest failure.** What actually happens is that Postgres comes up, Redis
comes up, sentinel comes up, `docker ps` is non-empty, the database answers — and the *application* is
absent. That is the silent no-op wearing the costume of a partially-working stack, and it is the exact
shape §5 rule 1 exists for: a non-empty result read as success. The corpus promises 11 containers here
(`run_guide.md:88`, `setup_guide.md:441`); `PROFILE=core` gives 5 and the documented `PROFILE=graphql`
gives 3.

**This is the second orchestrator-supplied fact this milestone has re-derived and corrected in two
iterations** (iter-58 corrected *"demo-1 GONE — 0 containers"*, a dead-daemon false absence over 11 live
containers). It is not a criticism of the hand-off; it is the milestone's own thesis applying to its own
inputs.

## Inputs consumed

1. **The PR #14 reconciliation** (`origin/pr-14`, head `e3d4692`, 2026-07-07) — 92 claims already absorbed,
   30 superseded, 5 contradictions standing, **0 refuted, ZERO new information; verdict DO NOT MERGE.** Its
   value is entirely **negative space**: the live defects are where the PR and our corpus *agree*. Not
   re-run here; its 7 ranked findings are the input.
2. **iter-58's finding** — a vetted pin advance moved **22 of 23** `main.go:N` citations and the fence caught
   **1** (4.5%). Schema-safety and citation-safety are unrelated properties and §7 rule 4 only measures the
   first.
3. **The user's scoping instruction** — the fold is a 5-done / 2-pending program, and *storage and messenger
   are not yet done*.

## Phase plan (tok — 2 planned lines)

- **A — re-survey + re-derive every denominator TOK-05 will cite** (done above; one correction landed).
- **B — author TOK-05 in the milestone-root `decisions.md`** with the five directed decisions recorded as
  `D-M257x-59-1..5`, plus the next-tik direction; close, commit, exit `tok-fired`.

## Escalation conditions

- The union-set question (`FIX-M257x-iter53-union-set`, 46 vs 35) is a **PENDING USER DECISION**. TOK-05 may
  state how class-based scoping *subsumes* it; it must not resolve it. Honored — `D-M257x-59-1` says so
  explicitly and leaves the number open.
- Clause 5 must not be re-cut, narrowed, deferred, or read met any other way. The user has ruled three
  times. Honored.
- `CHECK-M257x-iter38-ai-act-classification` needs an owner outside this milestone — not settled here.

## Acceptable close-no-lift outcomes

A tok moves no gate metric by construction. `closed-no-lift` is the correct and expected status: the
deliverable is a strategy plus five recorded decisions, not a metric delta.
