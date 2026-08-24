# MC02 — Spec notes

_Per-iter probe notes accumulate here during the checkpoint pass. **Nothing below is measured yet** — the
headers are the gate's shape, not findings._

> **Rule for this file:** every entry names **which stack** it was measured on (`demo-N` / `dev-N`), the
> **host**, the **date**, and the **artifact** the verdict was read from. A probe note without those four
> is not a measurement — it is a recollection, and this milestone exists because recollections graded
> green.
>
> **Read every doc claim against the RUNNING STACK, never against the diff that changed it.**

## Stacks under test

| role | stack | host | refs (platform / app / rext tag) | brought up | notes |
|---|---|---|---|---|---|
| demo | _not brought up_ | | | | must be the **DEFAULT** `/demo-up` path — clause 4 depends on it |
| dev | _not brought up_ | | | | see open question: `N=0` is never set-dressed |

## Pre-flight

- [ ] All three cluster milestones (M267, M269, M270) closed — MC02 does not run against a partial cluster
- [ ] `rosetta-extensions` tag **on origin** (`git ls-remote --tags origin`) — M236 rung zero
- [ ] Session table name **measured**, not assumed — see clause 1 below

_Not run._

## Clause 1 — start AND complete, three modalities

### Which table actually holds the rows?

**UNRESOLVED AT SCAFFOLD TIME.** The gate text says `public.job_simulation_sessions`; M269 and the M256
probe header say `jobsimulation.sessions`. Both may resolve on a live stack (the legacy `jobsimulation`
schema drop is a pending M810 step). **Measure first — a count against the legacy husk reads as a clean
RED and is a measurement error.**

_Unmeasured._

### Per-modality runs

| stack | modality | sim slug | rows before | rows after | delta | verdict |
|---|---|---|---|---|---|---|
| | chat | | | | | _not run_ |
| | code | | | | | _not run_ |
| | document | | | | | _not run_ |

_Voice is OUT — M271._

### What "COMPLETES" means per modality

_Not settled. A written row proves the hero **entered**; whether it proves she **finished** is the open
question. Record the answer here before grading, not after._

## Clause 2 — the entitlement ceiling

### Seeded-org enumeration

_Not enumerated. The list is the clause's denominator — pin it before probing, per `coverage-protocol.md`._

| stack | org | hero attempted | outcome | `ERROR_JOB_SIMULATION_LIMIT_REACHED`? |
|---|---|---|---|---|
| | | | | _not run_ |

_Design-time context (carried from M267, re-measure before acting): matcher `m6` has **no `default`
escape**, so the `p6` row must name the org id — the failure is **per-org by construction** and one org
does not generalise. UI string `errors.simulationLimitReached` at
`AISimulationStartWithoutSession.tsx:209`._

## Clause 3 — `/library/skill-paths` first paint

| stack | browser | viewport | cache | affordance in first paint? | t(affordance) | t(content) |
|---|---|---|---|---|---|---|
| | | | disabled | _not run_ | | |

_A `curl` or a bundle grep does not satisfy this clause — neither witnesses a paint order. **Verify in a
browser.**_

_An empty region that later fills is RED **even if it fills fast**. The number is recorded for
comparability across passes, not as a way to buy a pass._

## Clause 4 — the batch gate on the DEFAULT `/demo-up` path

### The verdict artifact

_Not obtained. Record the artifact path and quote the verdict line verbatim — the clause turns on the
**literal absence** of the word `skipped`._

### Which cause, if it skips

_Design-time context (carried from M269): **two independent causes, one symptom** — (1) `BIND_HOST` /
`D-M255-7`, `up-injected.sh` binds `0.0.0.0` whenever `STACK_PUBLIC_HOST` is set (line drifted to `:164`
at rext `v2.9.23-rext` / `bfd9835`; predicate unchanged); (2) a connection from the demo host to its own
tailscale IP hits the kernel socket and **bypasses `tailscale serve`**, which terminates TLS. **Fixing the
bind may leave the skip in place.** Clause 4 grades the symptom; this section is where the cause gets
named for the routing back to M269._

_Unmeasured._

## Clause 5 — both stacks

_Two evidence sets, or the clause is RED. Reusing one stack's numbers for the other is RED **by
definition, not by degree** — the two paths build studio-desk and seed differently, which is exactly how
divergence hides._

### Is clause 4 dev-side meaningful?

_Open. If the batch gate is demo-only by construction, that exception must be **written down here**, not
enacted by grading clause 4 once._

_Nothing measured._

## Clause 6 — the four docs, read against the running stacks

| doc | claim checked | observed on stack | verdict |
|---|---|---|---|
| `playthroughs.md` | live-Playthrough count + 4-state reporting map | | _not read_ |
| `verification.md` | the two tail gates + the batch-gate skip contract | | _not read_ |
| `demopatch-spec.md` | every patch the cluster added/re-pinned — **applied or REFUSED?** | | _not read_ |
| `seeding-spec.md` | the `p6` / entitlement seeding contract | | _not read_ |

_A doc that agrees with the commit and disagrees with the stack is **RED**._

## REDs and their routings

_None yet. Every RED routes to the milestone that owns it, and the routing is recorded in
[`decisions.md`](decisions.md) — **a routing is not a routing until the target's own doc says so.**_

| clause | stack | symptom | routed to | target doc updated? |
|---|---|---|---|---|
| _(none yet)_ | | | | |

## Verdict

_Not reached._
