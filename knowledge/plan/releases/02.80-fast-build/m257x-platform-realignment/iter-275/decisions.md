# iter-275 — decisions

## D-M257x-275-1 — hero-role occupancy is UNBOUNDED; the requirement is a bound, and the edit is deferred with its radius measured

**Context.** iter-274 reduced the succession failure to *"hold `DevOps Engineer` occupancy at ≤ 2 — one
number"*. Re-surveying `seeders/jobroleref.go` before implementing shows the framing is too small.

**Measured** (Org A, from the payload already captured; census sums back — 28 incumbents over 12 roles,
mean 2.33, histogram `1→4 · 2→2 · 3→4 · 4→2`):

| hero | role | occupancy | `riskScore` | consequence |
|---|---|---:|---:|---|
| Pat Ellis (employee) | DevOps Engineer | **3** | 45 | under the guard → both at-risk signals lost → **the Playthrough fails** |
| Morgan Reyes (manager) | Engineering Manager | **1** | 68 | **sole holder of her own title** — the state `orgRoleSet`'s comment says the mechanism prevents |

**Read, not inferred:** `orgRoleSet` adds a story's hero roles to that org's set and does nothing further;
`memberRoleAt` is `set[hashInt("<prefix>:role:<i>") % orgRolePoolSize]`. There is **no cap, no floor, no
reservation and no hero-keyed re-draw** on the path. The invariant is documented and **unenforced**.

**Decision.** The requirement is **a bound, not a value**: a hero's role carries **exactly one peer** —
`≥ 2` to satisfy the stated invariant, `≤ 2` to stay over the `RiskScore ≥ 50` guard. Fixing only Pat's
tail would green the test and leave the invariant violated, i.e. fix the tail that has a test and ignore
the tail that has only a comment.

**Deferred deliberately, with the radius measured rather than guessed.** The bound needs the population
size, which `memberRoleAt`'s signature does not carry; every parameter-free alternative is probabilistic
and would restore the lottery. That makes it **11 sites** — 6 production (`users.go:181`,
`membership_skills.go:145`, `population_evidence.go:132`, `certificates.go:150`, `profile.go:323`,
`target_roles.go:106`) + 5 test references — in the **single** derivation six seeders share, and whose own
comment records that the previous unification *"found only FOUR"* of the six while a seventh copy in a
test went RED when they were unified.

**Fate 3, named handler:** `FIX-M257x-275-bound-hero-role-occupancy-to-exactly-one-peer`, iter-276.
Grade it against iter-273's **169 s** binding suite, and watch **`negative-controls.spec.ts:429`** — the
tenancy control that caught this function's first cut at M257x iter-31 when hero roles leaked into every
org, making `DevOps Engineer` a key-role card on the contrast tenant's succession view. A fix for a
tenancy-flavoured Playthrough that breaks the tenancy control is not a fix.

**What would have happened without this iter.** iter-276 would have implemented "≤ 2", greened the
Playthrough, and shipped an org whose manager hero sole-holds her title against a comment promising she
never does — with the suite green and nothing to catch it.
