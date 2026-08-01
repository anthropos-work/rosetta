---
iter: 24
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-01
---

# iter-24 — `FIX-M257x-iter23-backend-directus-not-repointed`: point the re-point at the reader

**Active strategy reference:** `TOK-01` ("instrument first, then follow"), step 2 — *"fix the mechanism, not the
symptom"*, and its named mechanism-class: a hand-maintained list that stopped tracking the platform.

## Step 0 — re-survey (before targeting)

Platform origin HEAD re-fetched: **`2adcf71`, unchanged** (re-scope trigger stays at occurrence 1 of 2).
Section baselines re-measured at open: `stack-injection` **OK 286** (1 skip), `stack-core` **14F of 389** —
both match the standing baseline, so any movement is mine.

iter-23's routed finding still holds, and the re-survey **upgraded it from a candidate cause to a proven
one.** Three measurements, taken because a claim that "backend reads prod" is not the same claim as "and that
is why clause 2 fails":

| probe | result |
|---|---|
| `backend`'s own logs on the standing `demo-1` | **96** Directus lines, every one a `403 FORBIDDEN` — `directus_versions` (blocking `publicSkillPaths`, `getSkillPath`, `getOrCreateSkillPathSession`) and `library_categories` (blocking `libraryCategories`) |
| the **local** per-stack Directus, anonymous, from the host | `library_categories` **200** · `skill_paths` **200** · `task_sub_checks` **200** · `/versions` **200** |
| **prod** `content.anthropos.work`, anonymous | `library_categories` **403** · `/versions` **403** |

**The local instance answers the very collection `backend` is refused.** Backend's 403 set matches *prod's*
answers, not the local one's — so this is not a grant that was missed on the replay, it is a client pointed at
the wrong server. (§5 rule 7: the probe must not be able to satisfy itself. Asking only "does the local
Directus serve?" — which iter-18 did, correctly — cannot distinguish these two worlds; asking both ends can.)

This also **subsumes** `FIX-M257x-iter15-directus-versions-403` and explains why iter-19 found it independent
of iter-18's *serving* defect: it is independent, because it is a **pointing** defect, one layer up.

## Cluster / target identified

`DIRECTUS_DATA_CONSUMERS = ("cms",)` — in **both** twins (`stack-injection/gen_injected_override.py:34`, the
demo side; `stack-core/gen_override.py:46`, the dev side). Correct when `cms` was the Directus consumer; since
cms-in-app, `backend` reads Directus in-process and directly. Both files carry the now-false rationale in a
comment (*"`cms` is the only platform service that talks to Directus directly"*), and
`test_only_cms_is_repointed_not_other_services` **asserts `backend` must not carry the re-point** — a test
pinning the pre-merge shape, which is §8 rule 3 and the third occurrence in this milestone of *the suite
arguing for the defect*.

## Hypothesis

Adding `backend` to `DIRECTUS_DATA_CONSUMERS` re-points the actual reader at the per-stack Directus; the 403
class disappears and the skill-path / library reads resolve against the replayed catalog.

## Expected lift

The `directus_versions` 403 class (iter-15's 58 occurrences) clears. Clause 2's `20 live / 10 failing / 1
unimplemented` improves — **by an amount deliberately not predicted.** iter-19's diff proved the ten failures
span at least four causes; this fixes one. Any claim about the other nine is made *after* re-measuring.

## Phase plan

1. Re-point both twins; sweep the sibling leg in the same pass (§5 rule 9 — iter-12 was burned by fixing one
   twin and measuring the other).
2. Correct the tests that assert the pre-merge shape; **watch them RED first**, plus a declared-GREEN no-op
   control (§8 rule 5 — ten REDs cannot distinguish a discriminating fence from a brittle one).
3. Prove **live**, cheaply: regenerate the override + recreate only the `backend` container, then re-probe the
   403 class. §5 rule 15 — diagnose before paying for a full cold cycle.
4. Tag rext, push the tag, re-pin `.agentspace/rext.tag` + the `stack-demo` consumption clone.

## Escalation conditions

- A platform commit landing mid-iter → re-scope occurrence 2 → STOP.
- The re-point not clearing the 403 → a **finding**, not a failure: it would refute the causal chain above and
  that is worth more than the fix.

## Acceptable close-no-lift outcomes

The clause-2 number not moving, with the 403 class proven cleared, is a legitimate close — it would mean the
remaining failures never depended on this cause, which is exactly what iter-19's four-cause split predicts is
possible.
