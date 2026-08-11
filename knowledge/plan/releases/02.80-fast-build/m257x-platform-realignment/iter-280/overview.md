---
iter: 280
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-11
---

# iter-280 — the disposition registry's reasons are prose; grade them against the platform

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* A disposition
token is the archetype of a mechanically decidable class: it either resolves to a package the platform has
or it does not. No sentence has to be interpreted.

## Step 0 — re-survey (mandatory, and it moved the target)

The orchestrator named `roadrunner:merged-into-app` as false-since-iter-137. **Re-surveyed: it is already
fixed.** Harden pass 71 corrected the token to `deleted-not-merged` and the tree is green
(`service_registry_guard: ALIGNED`, 31 tests OK, rext `9933b6a`). Working the named target verbatim would
have re-done landed work — the route-lists-go-stale rule firing exactly as warned.

**The residual is what pass 71 wrote down while fixing it**, in `services.sh` and in its own test docstring:

> *"Nothing grades this token — `service_registry_guard` reads the NAME and stores the reason as free
> prose — so a corrected fact in the corpus does not reach a reason string here."*

Measured this iter, that is still true of the code: `service_registry_guard.py:296-297` does
`name, why = _split_decl(e, 2)` then `absent[name] = why`, and `why` is never read again. Pass 71 added a
**vocabulary** check (token ∈ the comment block above the array) and **one hardcoded row**
(`test_roadrunner_is_NOT_declared_merged_into_app`). Its own docstring states the ceiling: *"a vocabulary
check cannot catch this, because `merged-into-app` is a perfectly well-documented kind; it is just the
wrong one for this row."*

So the substitution, under the same TOK-08 strategy: **not the roadrunner row — the population it sits in.**

## Cluster / target identified

`SERVICES_NOT_IN_PLATFORM_COMPOSE` — **5 rows, of which 1 is factually graded (by name, by a hardcoded
test) and 4 are graded only against a comment block.** Flip `cms:merged-into-app` to
`cms:deleted-not-merged` today and every fence in the family stays green.

The class is `ROUTE-M257x-h70-corpus-and-code-prose-are-copies-with-no-fence` running in its **tooling→fact**
direction: the token is a factual claim about the platform, checked against prose written next to it.

Pre-measured population (the fence's expected verdict, sealed before the fence exists):

| row | token | predicate | `app/internal/<pkg>/` |
|---|---|---|---|
| `jobsimulation` | `merged-into-app` | must EXIST | PRESENT |
| `cms` | `merged-into-app` | must EXIST | PRESENT |
| `storage` | `merged-into-app` | must EXIST | PRESENT |
| `roadrunner` | `deleted-not-merged` | must NOT exist | ABSENT |
| `directus` | `rext-injected` | no app predicate (never a platform service) | n/a |

**All five are true today.** The deliverable is therefore not a repair — it is the fence that keeps them
true, which is precisely TOK-08's shape (enumerate the class, run it to zero, keep it green).

## Hypothesis

Promoting the disposition token from stored-prose to a **graded predicate over the `app` clone** converts 5
ungraded rows into a census, and makes the roadrunner fact enforceable **by rule** rather than by one
hardcoded per-row test.

## Expected lift

Not a `P` delta — clause 5's semantic reading is not being re-run. The lift is fence reach: **graded
disposition rows 1 → 5**, with the population enumerated and stated.

## Phase plan

1. Arm F in `service_registry_guard.py` — grade every token against the derived `app` clone.
2. The clone is **DERIVED from the platform clone** (its sibling), never a new flag — the family's own
   thrice-stated rule, so the guards cannot end up pointed at different clone sets.
3. **Fail-closed:** clone unreachable → CANNOT-MEASURE (exit 2), named. Proven RED with the precondition
   absent, per the standing "a capability probe that fails OPEN disarms the check it guards" rule.
4. Mutation controls per token kind, both polarities.
5. Fix the false stated reason found in pass 71's own test (below).

## Side-finding to land in-scope

`test_every_reason_is_one_of_the_documented_kinds` docstring says **"Both directions."** The assertion
implements **one** — it computes `undocumented` and asserts nothing about a documented-but-unused kind. It
is currently vacuous-but-true (all 3 kinds are in use). This is the milestone's own recurring class — *a
correct fix with a false stated reason* — and it is in-scope because it is the same test class this iter
extends.

## Escalation conditions

- If arm F fires RED on a row the table above calls true → the pre-measurement was wrong; STOP and report
  the disagreement rather than tuning the fence to agree with it.
- If the `app` clone is unreadable → arm F is CANNOT-MEASURE; that is a reportable state, not a pass.

## Acceptable close-no-lift outcomes

If the census proves the reason strings are already fully graded by some fence not found in this survey,
the iter closes `closed-no-lift` with that falsification recorded — the reach claim would be wrong, and
that is worth more than a redundant arm.
