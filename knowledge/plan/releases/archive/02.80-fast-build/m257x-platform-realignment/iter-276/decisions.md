# iter-276 — decisions

## D-M257x-276-1 — the occupancy bound is a RESERVED SLOT, and `st` rides alongside the fenced argument rather than replacing it

**Context.** iter-275 specified the requirement (a hero's role carries exactly one peer: `≥ 2` for the
invariant, `≤ 2` for the `RiskScore ≥ 50` guard) and measured an 11-site radius, deliberately not
beginning the edit.

**Decision — mechanism.** The peer is a **reserved population slot**, computed by `heroPeerSlots` and
skipping hero slots, with the general draw restricted to the non-hero remainder. Rejected alternatives,
each falsified rather than dismissed:

| alternative | why rejected |
|---|---|
| reserve the first `heroCount` indices | heroes ride **hashed** indices, so a reservation lands on a hero ≈ ⅓ of the time in Org A's shape → that role gets **zero** peers (the sole-holder tail) |
| `slot = (i + off) % n` permutation over `n` alone | same collision with hero slots; needs the hero index set regardless |
| bias the draw's modulus toward `n` | probabilistic — restores the lottery this fix exists to remove |
| cap occupancy at ≤ 2 only | greens the test and leaves the invariant broken — **explicitly** what iter-275 forbade |

**Decision — signature, and this is the load-bearing half.** `st` is **added** as a separate argument;
the hero-roles argument is left exactly as it was. Tidier would have been to pass `st` alone and derive
the names inside. That was rejected because **`role_tenancy_fence_test.go` fences the hero-roles
argument's AST shape** (`storyHeroRoleNames(st)` or a forwarded parameter) — the guard written after the
first cut of `orgRoleSet` gave every org every story's hero roles and `negative-controls.spec.ts:429`
caught it live. Replacing that argument would have moved the construct out from under its own fence
**while reading as a cleanup**. The redundancy is commented at the function with a *do not clean this up*
note.

**Consequence worth keeping.** The fence is not merely unbroken — it is the **measurement** that this
sweep reached all six production sites (`6 memberRoleAt call site(s), 1 forwarding helper(s)`), where the
previous unification of this same derivation *"found only FOUR"* and a seventh copy in a test went RED.

## D-M257x-276-2 — a unit test that REQUIRED the defect, superseded in place

`TestOrgRolePool_BoundsTitlesAndKeepsThemInTheirOwnTenant` asserted the hero's title must have **≥ 2
supporting holders**. Against iter-274's measured step function that is `≥ 3` incumbents →
`structuralRisk 53` → `riskScore 45` → **under** the `≥ 50` at-risk guard: the precise state in which
`workforce-intelligence.talent-pool.UC1` failed.

**The test guarding this function encoded the failing state as a requirement.** Its rationale — a
one-incumbent tiebreak deciding the hero's card — had already been closed separately by
`orgRolePoolSize = 12` (12 roles < the view's 25-card render budget → nothing to truncate).

**Decision.** Superseded to `!= 1`, with the supersession *argued in the test body* rather than quietly
rewritten, so the next reader can see that the old assertion was wrong and why. Recorded because
"changing a test so the code passes" and "removing an assertion that demanded the bug" are the same edit
seen from two sides, and only the reasoning distinguishes them.

**The generalisable form:** *a proxy assertion outlives the thing it proxied.* `≥ 2 supporting holders`
was a stand-in for "not a sole holder" written when occupancy had no upper consequence. When iter-274
gave occupancy an upper consequence, the proxy silently became a demand for the defect. **Assertions
written as proxies should name what they proxy**, so a later change of consequence can find them.

## D-M257x-276-3 — the RED proof came before the fix, and the live proof came from origin

Two disciplines applied deliberately, both from this milestone's own standing failure list:

1. **A direct anti-regression test can be GREEN while the bug is live** (iter-270). So
   `TestHeroRoleCarriesExactlyOnePeer` was authored against the **unmodified** signature and run on the
   pre-repair tree first: RED with 3, 3 and 17 supporting holders. Only then was the fix written.
2. **Tagging is not publishing** (M236 lost an iteration to it). The suite was re-run a second time from
   the stack's **own clone**, fetched **from origin** at `fast-build-m257x-iter-276` and rebuilt, with no
   `PT_STACKSEED` override — because run 1 proves only the working copy, and a stack can obtain only what
   origin has. `git ls-remote --tags origin` verified the tag resolves to `0a8674e` before that run.

## D-M257x-276-4 — the 11th hero role reads 1 and that is the specification, not a residual

Live occupancy showed 10 of 11 cockpit hero roles at exactly 2 and `Halcyon Retail | Operations Manager`
at **1**. Two hypotheses were formed and **both falsified by measurement** before the answer was found:
the name fails to resolve (**no** — it resolves to `J-OPERAT-C7F2`), and the org's population comes from
the generated batch rather than `memberRoleAt` (**no** — its two other hero roles sit at 2).

The cause is in the seed declaration: `pt-world.seed.yaml:199` gives Nils Brandt `org_membership: none`
— *"The SOLO user: no organization at all"*. She writes **no membership row**, so she is not an incumbent
of any org, and the single holder counted there **is** her role's reserved peer. Behaviour is exactly as
specified.

**Recorded rather than waived** because "10 of 11" reads as a partial failure, and a green suite would
have absorbed it silently. The residual worth naming is cosmetic and is **not** routed as a defect: the
reservation spends one population slot on a role whose hero is not in the org. It is believable, costs
nothing, and no surface depends on it.
