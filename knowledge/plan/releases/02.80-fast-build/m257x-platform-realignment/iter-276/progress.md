# iter-276 — bound hero-role occupancy to exactly one peer

**Type:** tik

Implements `FIX-M257x-275-bound-hero-role-occupancy-to-exactly-one-peer` — the handler iter-275 named,
with its 11-site blast radius measured and the edit deliberately not begun.

## Phase A — the fix

`memberRoleAt` drew every supporting member's title uniformly over the org's 12-role set, so a hero's
own role landed on an **unbounded** occupancy (measured live pre-fix: 3, 3 and 17 supporting holders in
the property test's three shapes). The bound is now a **reserved slot**, not an outcome of the draw:

- `orgRoleSetSplit` returns the set **plus `heroCount`** — the one fact the slice alone could not carry.
  The set is composed hero-roles-first, so the split point is exactly that count. `orgRoleSet` survives
  as a thin wrapper, so its existing callers and tests are untouched.
- `heroPeerSlots(prefix, st, heroCount)` seats **exactly one supporting member per hero role**, scanning
  from a per-role hash offset (so peers spread through the org rather than clustering on slot 1, which is
  the story's assigner/admin membership) and **skipping hero slots**.
- Everyone else draws from `set[heroCount:]` — the **non-hero remainder** — so the count cannot drift
  back up.

**Why hero slots must be skipped, and why this is the fact that forced the signature.** Heroes ride on
*hashed* population indices (`personaIndexMapForStory`), not low ones. A reservation that ignored them
would seat a peer on a hero's own slot — in Org A's shape roughly a third of the time — and that hero
role would end up with **zero** peers, i.e. the sole-holder tail this bound exists to close. A hero slot
never reaches `memberRoleAt`; UsersSeeder writes her declared role there. **This confirms iter-275's
radius reasoning rather than shortcutting it:** the derivation genuinely needs the population, and no
parameter-free formulation is both exact and deterministic.

### The one design decision worth its own paragraph

`st` is added as a **separate argument** rather than replacing the hero-roles one, which would have been
the tidier signature. `role_tenancy_fence_test.go` **fences that argument's shape** — it must be
`storyHeroRoleNames(st)` at the call site or a forwarded parameter — because the first cut of
`orgRoleSet` handed every org every story's hero roles and `negative-controls.spec.ts:429` caught it
live. Collapsing the two arguments would have moved the construct **out from under its own fence** while
looking like a cleanup. The redundancy is deliberate and is commented as such at the function.

**The fence is untouched and stayed green — and it is also how this sweep is known to have reached all
six sites**, where the previous unification of this same derivation *"found only FOUR"*:

```
role-tenancy fence: scanned 54 seeder sources, 6 memberRoleAt call site(s), 1 forwarding helper(s)
role-tenancy forwarder fence: 1 helper(s), 1 call site(s) checked
```

**Radius as landed:** 6 production call sites (`users.go:181`, `membership_skills.go:145`,
`population_evidence.go:132`, `certificates.go:150` via the `memberRoleName` forwarder, `profile.go:323`,
`target_roles.go:106`) + 3 test references + the derivation itself. iter-275 measured 11; landed 9 edits
across 9 files, the difference being test references that the same `perl` sweep covered in one pass.

## Phase B — the unit gate, RED before GREEN

The standing failure mode is a direct anti-regression test that is **green while the bug is live**
(iter-270). So `TestHeroRoleCarriesExactlyOnePeer` was written **first, against the unmodified
signature**, and run on the pre-repair tree:

```
--- FAIL: TestHeroRoleCarriesExactlyOnePeer
  pt-org-a (size  40): hero role "DevOps Engineer"     held by  3 supporting member(s), want exactly 1
  pt-org-a (size  40): hero role "Engineering Manager" held by  3 supporting member(s), want exactly 1
  pt-org-c (size 200): hero role "Data Analyst"        held by 17 supporting member(s), want exactly 1
```

It proves the property **structurally — any pool, any population** — at both the 40-person story shape
that failed and the 200-person showcase, and asserts the reservation has not collapsed the general draw.

### A test that required the defect

`TestOrgRolePool_BoundsTitlesAndKeepsThemInTheirOwnTenant` asserted
`counts["DevOps Engineer"] < 2` → error, i.e. **the hero's title must have ≥ 2 supporting holders**.
Measured against iter-274's table that is `≥ 2 + the hero herself = ≥ 3 incumbents = structuralRisk 53 =
riskScore 45` — **under** the `≥ 50` at-risk guard. **The unit test guarding this function required the
exact state that failed the Playthrough.** Its stated rationale (a one-incumbent tiebreak deciding the
hero's card) had already been closed separately by `orgRolePoolSize = 12`: 12 roles sit under the view's
25-card render budget, so there is nothing left to truncate. Superseded in place, with the supersession
argued in the test rather than silently rewritten.

Full `stack-seeding` suite green (16 packages).

## Phase C — the live gate

**Two consecutive cold reset-to-seed cycles on `demo-2`, `SUITE_RC=0` both times.** Durations are
**CONTENDED** (12 CPU / 24 GiB, third-party load) and are **not** baselines.

| run | seeder | wall clock | Playwright | ledger |
|---|---|---|---|---|
| 1 | `/tmp/stackseed-iter276` (authoring copy, via `PT_STACKSEED`) | 21:15:18Z → 21:17:41Z = **143 s** | 215 passed | **30 passing / 0 failing / 1 unimplemented / 0 unimplementable** |
| 2 | the stack's **own pinned** `demo-2/bin/stackseed`, no override | 21:21:14Z → 21:23:48Z = **154 s** | 215 passed | **30 passing / 0 failing / 1 unimplemented / 0 unimplementable** |

- `workforce-intelligence.talent-pool.UC1` — **PASS** in both. This is `pt-workforce-succession`, the
  single failure iter-273 measured.
- **`negative-controls.spec.ts:429` — PASS in both** (27 s), the tenancy control iter-275 required be
  watched. *A fix for a tenancy-flavoured Playthrough that breaks the tenancy control is not a fix.*
- The 1 `unimplemented` is the declared TODO carrying M256's machine-checked `will-not-build` verdict —
  the same one `playthroughs.md` counts. **31 total = 30 live + 1 verdicted TODO.**

**Run 2 is the load-bearing one, and it is why the tag was pushed before it was measured.** Run 1 proves
the working copy; only run 2 proves what a stack can actually obtain — the pinned clone was fetched
**from origin** at `fast-build-m257x-iter-276` (`git ls-remote` verified: `0a8674e`) and rebuilt. Tagging
is not publishing, and this milestone has lost an iteration to that distinction before.

### Live occupancy, measured directly rather than inferred

Queried `demo-2` while the fresh seed was live — all 11 cockpit heroes, per org:

| org | hero role | incumbents |
|---|---|---:|
| Meridian Labs | **DevOps Engineer** (Pat Ellis) | **2** — was **3** |
| Meridian Labs | **Engineering Manager** (Morgan Reyes) | **2** — was **1** |
| Vertex Logistics | Data Analyst · Business Operations Analyst · Supply Chain Analyst · Engineering Manager | 2 · 2 · 2 · 2 |
| Kestrel Hiring Group | Account Executive · Talent Acquisition Specialist | 2 · 2 |
| Halcyon Retail | Account Executive · Business Analyst | 2 · 2 |
| Halcyon Retail | Operations Manager | **1 — and correctly** (below) |

**Both tails closed on the org that had both broken.** Pat's role cleared the guard (45 → 54) and
Morgan stopped sole-holding her own title.

**The one row that is not 2 is not an exception — it is the spec.** `Operations Manager` is Nils
Brandt's role, and `pt-world.seed.yaml:199` declares her `org_membership: none` — *"The SOLO user: no
organization at all"*. She writes **no membership row**, so she is not an incumbent, and the single
incumbent counted there **is** her role's reserved peer. The bound behaved exactly as specified. This is
recorded because "10 of 11" would otherwise read as a partial failure; it was chased to the seed
declaration rather than waved through, and the first two hypotheses (name fails to resolve; generated-batch
org) were both **falsified** by measurement — `Operations Manager` resolves to `J-OPERAT-C7F2`.

## Phase D — clause 2

**Gate clause 2 is MET: 30 live / 0 failing / 0 error**, twice, cold, at the shipping pin, with the
tenancy control green — the exact form the exit gate asks for.

## Close — 2026-08-10

**Outcome:** **Gate clause 2 is MET — 30 live / 0 failing / 0 error**, twice, on cold reset-to-seed, the
second run from the stack's own clone pinned at an origin tag. The bound landed on both tails: Pat
Ellis's role 3 → 2 incumbents (riskScore 45 → 54, back over the `≥ 50` at-risk guard, so
`workforce-intelligence.talent-pool.UC1` passes) and Morgan Reyes's 1 → 2 (she no longer sole-holds her
own title, enforcing an invariant documented since iter-31 and enforced nowhere).
`negative-controls.spec.ts:429` green in both runs.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-276-1` (reserved slot; `st` rides alongside the fenced argument rather than
replacing it), `D-M257x-276-2` (a unit test that REQUIRED the defect, superseded in place),
`D-M257x-276-3` (RED-before-GREEN, and the live proof came from origin), `D-M257x-276-4` (the 11th hero
role reads 1 and that is the spec).

**Side-deliverables:** none.

**Routes carried forward:**
- **`FIX-M257x-275-bound-hero-role-occupancy-to-exactly-one-peer` is CLOSED** — landed, proven offline
  and live. (Written in full deliberately: iter-277's census caught the elided form `FIX-M257x-275-…`
  here as a **truncated stem** that `route_disposition_guard` reads as live backlog — `§5` rule 73.)
- **Gate clause 5 remains the last open clause**, unmeasured since iter-131 (`P = 29 / N = 47`, a floor).
  iter-277 measures it under `TOK-08`. **This is now the only thing between the milestone and its gate.**
- Still open, none absorbed into this green: `ROUTE-M257x-274-successor-half-is-uncovered`
  (`successors`/`topTalents`/`readyCount` empty on every reset), `ROUTE-M257x-274-tie-order-is-unstable`
  (**materially reduced** — the hero's card no longer depends on a tiebreak among equal scores, but the
  ordering itself was never made stable, so the route stands), `FIX-M257x-269`,
  `ROUTE-M257x-270-directus-consumer-cms-key`, `FIX-M257x-266`, `FIX-M257x-265`, `ROUTE-M257x-h59`,
  `ROUTE-M257x-h65`.

**Lessons:**
1. **A proxy assertion outlives the thing it proxied.** `counts["DevOps Engineer"] >= 2` stood in for
   "not a sole holder" when occupancy had no upper consequence. iter-274 gave occupancy an upper
   consequence, and the proxy silently became a **demand for the defect** — guarding the function while
   requiring the state that broke it. Assertions written as proxies should name what they proxy.
2. **A fence can make a signature change cheaper, not just safer.** The instinct was to replace the
   hero-roles argument with `st`; the tenancy fence made that visibly wrong, and adding `st` alongside
   left the fence untouched *and* supplied the proof the sweep reached all six sites — where the previous
   unification of this same derivation found only four. The fence answered "did I get them all?" for free.
3. **Chase the one row that does not fit, and be willing to find it is the spec.** 10 of 11 hero roles
   at exactly 2 would have been reported as a partial win; two hypotheses were falsified by measurement
   before the seed declaration explained the 11th exactly. "Mostly" is a signal to keep reading.
