# iter-19 — decisions

## D-M257x-19-1: the "downstream of the unserved content layer" attribution is REFUTED

**Decision.** iter-18's close routed `FIX-M257x-iter15-directus-versions-403` and
`FIX-M257x-iter15-library-category-expansion` forward with the caveat that both had been measured through a
Directus that served nothing, and that they might therefore be downstream of it. **They are not.** The
caveat is withdrawn; both are independent defects and the next tik may work on them directly.

**The measurement.** Full suite, `--reset`, on the stack iter-18 proved green (three cold cycles,
`anon GET /items/task_sub_checks` = 200):

    summary: {'passing': 20, 'failing': 10, 'unimplemented': 1}   total 31

and the ten failing ids are **byte-identical** to iter-15's set (`diff` of the sorted lists: no output).

**Why this was worth an iter rather than an assumption in either direction.** The alternative was to spend
the next tik on a 106-occurrence cause that might have evaporated, or to declare it unchanged without
looking. `platform-alignment.md` §5's closing rule — *verify a claim before escalating it* — applies to a
claim this milestone made about its own routed work one iter earlier.

## D-M257x-19-2: the number reproduces, which retires a standing doubt about iter-15's

**Decision.** Record `20 / 10 / 1` as **reproduced**, not merely re-asserted.

iter-15 measured `20/10/1` and then scored `17/31` on a re-run, and attributed the difference to running
without `--reset` after a mutating run (three `onboarding.*` negative controls asserting *"onboarding is
INCOMPLETE"* against a world that was no longer fresh). That explanation was reasoned, not tested. This run
used the real `--reset` path on a differently-built stack and landed on `20/10/1` with the same ten ids —
so the explanation now has a confirming observation, and the reset-vs-additive discipline is load-bearing
rather than merely stated.

## D-M257x-19-3: the largest cause is TWO fields, not one — and the count moved

**Decision.** Re-scope `FIX-M257x-iter15-library-category-expansion` before anyone works it.

**Measured on the green stack**, from the `backend` container's own log:

    119 × cannot unmarshal string into Go struct field JobSimulation.data.library_category of type struct
     11 × cannot unmarshal string into Go struct field JobSimulation.data.job_position    of type simulation.…

iter-15 named `library_category` alone and counted **106**. The class is the same shape — `app` reads an
EXPANDED relation, Directus returns the raw id string — but it spans **at least two fields**, and a fix
aimed at one field name would leave the other. `directus_versions` appears **58** times in the same log.

**Not concluded:** that the two fields share a single root cause (they are both relation-expansion sites, so
it is likely, and that is exactly why it should be measured rather than assumed — this milestone's dominant
defect is a claim reported without being measured).

## D-M257x-19-4: `run-playthroughs.sh` requires `stackseed` on PATH and derives nothing

**Decision.** Route as `FIX-M257x-iter19-playthrough-runner-path`; do not fix mid-measurement.

`run-playthroughs.sh:118` calls a bare `stackseed`, which is **not on PATH on this host** — the bring-up
builds it into the stack's own `stacks/demo-<N>/bin/`. The run died at the reset step with
`stackseed: command not found` and was re-run with the stack's bin dir prepended.

This is the milestone's own §2 shape in the suite that measures gate clause 2: a value that is a property of
the environment (where this stack's tooling was built) is instead assumed of the host. The runner already
derives its ports, its bases and its seed path from `N`; the binary directory is derivable from the same `N`.
Left alone here because changing the instrument during the measurement it is taking is exactly the mistake
this iter exists to avoid.
