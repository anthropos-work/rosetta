# iter-24 decisions

## D-M257x-24-1 — `cms` stays in `DIRECTUS_DATA_CONSUMERS`

The obvious edit was `("cms",)` → `("backend",)`: the reader moved, so move the re-point. Rejected.

The `cms` **container** is still started by the default `graphql` profile — a merged-into-`app` husk kept as
the rollback path until platform M810 — and `messenger` still addresses it over `CMS_RPC_ADDR`. Leaving it on
prod would reproduce the exact split this iter is fixing, one service over, and would only surface the day
someone exercised the husk's own Directus path.

The list is now `("cms", "backend")` and will shrink to `("backend",)` at M810, when the husk stops being
started at all. That is a **deletion driven by an observable event**, not a guess about the platform's plans.

## D-M257x-24-2 — proven live on a recreated container, not a cold cycle

Clause 1 is met by three consecutive cold cycles, so the temptation was to prove this the same way. §5 rule 15
says diagnose before paying for a full cycle, and the causal chain here is narrow enough to test directly: the
generator's emission is proven by regenerating from the **pinned consumption clone** and diffing (one added
line), and the runtime effect is proven by recreating **only** `backend` and re-issuing the two queries that
were 403ing.

**What this does NOT prove, stated so nobody quotes it as more than it is:** that a cold `demo-up` produces
the same result end-to-end. It should — the generator is the same code at the same tag, and nothing else in
the bring-up touches this env — but that is reasoning, not a measurement. The next cold cycle run for any
reason should confirm it, and the clause-2 re-measure will exercise it incidentally.

## D-M257x-24-3 — the clause-2 lift is deliberately NOT predicted

This fixes one of at least four causes iter-19's sorted-id `diff` established behind `20 live / 10 failing`.
The overview records an expected lift with no number attached, on purpose.

This milestone has already made the attribution error once in the other direction — iter-15 attributed the
skill-path failures to a permissions problem, iter-18 fixed a serving defect, iter-19 measured and found the
failing set byte-identical. The discipline that caught it was refusing to reason from "the layer I changed
was broken" to "therefore the symptom was downstream of it."

So: fix the runner-path defect first so the suite can reset itself, then run the full suite, then compare
sorted ids — not two summary lines. `20/10/1` twice could still be ten different failures.
