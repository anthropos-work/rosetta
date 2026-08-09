# iter-178 — decisions

## `D-M257x-178-1` — the NOT-REACHED clause is MEASURED before it is worked, and the measurement is the finding

`derived_count_guard` has printed *"NOT REACHED: the `N of M` prose shape"* on every run since iter-173,
routed as `SURVEY-M257x-iter173-derived-count-guard-reach`, and **nobody had counted the class.**

Measured at corpus `794b167`: **61** occurrences across `corpus/**` + `CLAUDE.md`, **9 over 8 lines** on
the clause-5 surface (`corpus/architecture/**` + `corpus/services/**`).

**Decision: the size goes first, and it changes the ranking.** Nine is enumerable exactly. The class did
not stay unreached for five iters because it was hard — it stayed unreached because an unmeasured hole
looks the same size as every other unmeasured hole. → `§8`.

## `D-M257x-178-2` — Arm D disposes of the claim; it does NOT verify the value

The reason iter-173 declined this shape is still binding: `M` names no source, and attributing it to a
nearby table is the inference `D-M257x-117-2` that iter-119's refutation made unaffordable.

**Decision: split the verification from the disposition and take only the half that is decidable.**
Arm D asserts that every clause-5 `N of M` carries a written `DERIVABLE:` / `OBSERVED:` / `HISTORICAL:`
disposition, **both directions** (an undispositioned claim is RED; a disposition matching no live claim
is RED). It never claims the number is right, and the printed clause says so in the same breath.

This is §9 iter-159 — *grade the instrument at the grain of its claim* — used constructively rather than
as a criticism: the grain at which this class is decidable is *"has somebody said whether this can be
re-derived, and how"*, and that grain is worth fencing.

**Keyed without a line number** (`<path>::<claim text>`), so an edit above a known site is not a fake
finding — `repair_postcondition`'s baseline rule, reused rather than re-invented.

## `D-M257x-178-3` — the arm goes INTO the existing guard, not into a new fence module

A new fence module would drag four registries behind it — the `stack-core/README.md` guard table, a
`derivation_registry` entry, a `guard_family` invocation, and the mechanical-fences battery's seed list.
**Three of those four have been caught rotting in this milestone** (iters 173, 174, 175), and iter-177's
own fence now makes two of them RED-on-omission.

**Decision: extend `derived_count_guard`.** It is the guard whose own output declines this class, so the
disclosure and the check live in one place and cannot drift apart. A new module would have been tidier
to read and strictly worse to maintain.

## `D-M257x-178-4` — the four DERIVABLE claims were re-derived at their cited refs, and all four hold

Substrate stated first (`D-M257x-122-4`: *before believing a defect, read the substrate line*) — read
from `stack-demo/`, `app @ ad9f3c498`, the ref the prose itself cites:

| claim | site | re-derivation | verdict |
|---|---|---|---|
| `1 of 43` | `service_taxonomy.md:300` + `ant-academy.md:77` (twin) | `graph/schemas/*.graphqls` → **43**, `academy.graphqls` present | **TRUE** |
| `31 of 135` | `architecture_overview.md:398` | struct types embedding `ent.Schema` → **135 exactly**; live `OrganizationMixin{}` → **29**; +Membership +Organization = **31** | **TRUE** |
| `6 of 7` | `shared_libraries.md:257` | 7 repos on disk carry a `go.mod`; 6 require `taxonomy`; `roadrunner` the sole exception | **TRUE** |

**The instrument proved itself on the way** (§9 iter-149 — *a census returning ZERO must prove its
instrument*). The schema count surfaced a **fourth** `Policy()` declaration nobody's arithmetic had
mentioned — `user.go` — which had to be adjudicated rather than assumed away. It is correctly OUT:
`User`'s policy is `FilterSameUserRule` / `DenyNotSameUser`, a **per-user** filter, not a per-
organization one. That is an independent corroboration of iter-52's refutation of the `32` reading,
reached from the code rather than from the ledger.

**Decision: record the zero WITH the near-miss.** A census that reports only its zero has thrown away
the evidence that it can distinguish.

## `D-M257x-178-5` — Arm D's fixed surface creates a fail-open, and it is closed by a PAIR

Arm D's subject is the clause-5 roots, not the `--root` argument — the disposition table is a statement
about *that* surface and would mean something else over any other. The cost is that the arm necessarily
no-ops on a tmp tree with no `corpus/architecture/`: **a capability probe that fails OPEN**, §8
iter-174's exact shape, introduced deliberately this time.

**Decision: close it with two halves, neither sufficient alone.**

1. The guard **records the surface state** (`arm D surface: present: … | absent`) on every run and in
   `--json`, so a no-op is visible rather than silent.
2. `test_arm_d_surface_is_PRESENT_on_the_real_tree` asserts the surface exists on the real checkout, so
   the fail-open can never be the state of the tree that matters without a RED.

A fail-open you can see and a fail-open that cannot be the real tree's state are different things, and
only the pair is safe.

## `D-M257x-178-6` — the population SIZE is asserted, not remembered

The class was fenceable *because* it was nine; sixty-one would have been a different decision
(`D-M257x-178-3` would have gone the other way). **Decision: `test_the_arm_d_population_is_the_MEASURED_
nine` pins it.** A move in either direction is a real signal — a claim arrived and needs a disposition,
or a claim left the surface — and either way the iter that moves it has to say which.

This is iter-177's rule applied to this iter's own number: a count about a population that does not name
what produced it is unreadable, so the test names the derivation (Arm D's own site count) and the
docstring names the split (8 dispositions covering 9 occurrences, because one line states its claim
twice).

## `D-M257x-178-7` — the 52 occurrences outside clause 5 are NOT swept into it

`D-M257x-129-2` binds: the clause names `corpus/services/**` + `corpus/architecture/**`, and work
outside it is booked to the user's standing ask, never to the clause. Arm D's surface is the clause-5
roots and the printed clause says `(clause-5 surface only)`.

**Decision: scope it, say so in the instrument's own output, and route the rest.** Widening the arm to
`corpus/ops/**` is a defensible future iter; doing it silently, so that a green over 61 sites reads as a
statement about the gate's surface, is exactly the conflation `F4` was written to forbid.
