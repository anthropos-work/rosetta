---
iter: 155
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-155 — the three ungradeable services, and what a fence may honestly derive about them

**Active strategy reference:** `TOK-08` — census the mechanical classes; stop sampling them.

**Step 0 — re-survey.** Two routes were candidates.

`SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` was **sized and rejected for this iter,
with the measurement recorded** so the next session does not repeat it: across the 5 test-bearing sections
(109 files) there are **2,854** string-literal `assertIn`/`assertNotIn` calls, of which **766** are
expression-shaped by the naive predicate. That is iter-150's 30-to-1 over-report shape at 30× the volume —
the class needs a **sharper predicate** (the haystack must be a subject file's SOURCE TEXT, not any
string), and finding that predicate is the work, not the sweep. Routed with the numbers attached.

`FIX-M257x-iter153-stack-injected-services-have-no-rows` is taken instead, and its blocker is **partly**
refuted — which is the point of this iter.

**Cluster / target identified.** iter-153 declared `hiring-app`, `fake-fapi` and `fake-bapi` unprobed
because *"a probe row needs a container name, a host port and a health target that have been **observed on
a live stack**"* (`D-M257x-153-6`). Two of those three are **derivable from the generated override** and
one is not:

| ingredient | derivable here? | evidence |
|---|---|---|
| container name | **yes** | the layout constant `container_for_project()` already encodes |
| host port | **yes** | generated at offset 0: `hiring-app:3001`, `fake-fapi:5400`, `fake-bapi:5401` |
| **health target** | **no** | requires knowing what each service answers on, and `fake-bapi` publishes `127.0.0.1:5401 → 443` in-container, i.e. **TLS**, where an `http` probe would fail for the wrong reason |

**Hypothesis.** The registry's `docker` probe kind needs **only** a container name. So a `docker`-kind row
is fully derivable and asserts something real and load-bearing — *the container exists and is running* —
without inventing a health target. That converts three services from **ungradeable** to
**liveness-graded**, and leaves exactly the residual that genuinely needs a live stack.

**Expected lift.** Not an `N` reading. Three services move from declared-ungradeable to liveness-probed;
the declaration array shrinks to the honest residual.

**Phase plan.** A (derive the three rows) → B (add them + re-point the declarations so iter-152's arms
stay honest) → C (fence: the rows are derived from the override, both directions) → D (gates + close).

**Escalation conditions.** If adding the rows makes a **dev** stack or the main `anthropos` project probe
for containers it does not run, stop — the scoping must already exclude them, and if it does not, the
finding is about the scope, not the rows.

**Acceptable close-no-lift outcomes.** If a `docker`-kind row cannot be scoped safely to the stacks that
actually run these services, the iter closes on that falsification and the declaration array stands.
