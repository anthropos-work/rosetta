# iter-275 — progress

**Type:** tik, under `TOK-08`. Route: `FIX-M257x-274-devops-occupancy-must-stay-at-two`.

## Phase 1 — pre-registrations sealed

Four, sealed before the census. See `overview.md`.

## Phase 2 — the occupancy census (Org A, from the payload already on disk)

| role | incumbents | names | `riskScore` |
|---|---:|---|---:|
| Admin & HR Coordinator | 1 | Liam Costa | 68 |
| Administrative Coordinator | 1 | Chloe Nakamura | 68 |
| Advanced Analytics Specialist | 1 | Chloe Mensah | 68 |
| **Engineering Manager** | **1** | **Morgan Reyes** — *the manager hero, alone* | **68** |
| Adjunct Professor - Strategy, Leadership And People | 2 | Ethan Silva, Omar Moreau | 54 |
| Advanced Analytics and Business Intelligence Manager | 2 | Marcus Lindqvist, Diego Lindqvist | 54 |
| Administration officer | 3 | Clara Okafor, Lucas Larsen, Layla Kovac | 45 |
| Administrative Assistant | 3 | Layla Dubois, Arjun Larsen, Sven Romano | 45 |
| Administrator | 3 | Diego Romano, Priya Haddad, Hassan Larsen | 45 |
| **DevOps Engineer** | **3** | **Pat Ellis**, Noah Novak, Aisha Esposito | **45** |
| Ad Tech Engineer | 4 | Noah Moreau, Nick Silva, Raj Moreau, Bruno Schmidt | 45 |
| Administrative Virtual Assistant | 4 | Nick Schmidt, Mateo Nakamura, Aisha Andersen, Ava Romano | 45 |

**Histogram** (incumbents → roles): `1→4 · 2→2 · 3→4 · 4→2`. **Sum 28 over 12 roles, mean 2.33.**

## Phase 3 — pre-registrations graded

| PR | verdict | evidence |
|---|---|---|
| **PR-1** — the census sums back | **HOLDS** | 28 incumbents = the 28 members `querySuccessionMembers` returns (iter-267). Incumbency and membership are the same population, so iter-274's step function was read against the right denominator |
| **PR-2** — Morgan sole-holds her own role | **HOLDS** | `Engineering Manager`: **1** incumbent, and it is the manager hero |
| **PR-3** — nothing bounds hero-role occupancy | **HOLDS** | `orgRoleSet` adds hero roles to the org's set and nothing more; `memberRoleAt` is `set[hashInt("<prefix>:role:<i>") % 12]`. **No cap, no floor, no reservation, no hero-keyed re-draw** anywhere in the path |
| **PR-4** — the hero roles sit in different tails | **HOLDS** | **1** vs **3**, and the distribution is bimodal (`1→4`, `3→4`) — both heroes at extremes, neither at the 2.33 mean |

## Phase 4 — one mechanism, two opposite failures, both live right now

`orgRoleSet`'s comment states the invariant plainly: *"so no hero is ever the sole holder of her own
title"*. **It is not enforced, and Org A violates it today.**

The two failures are visible **in the same payload**, and they are the two tails of the same unbounded
draw:

| | occupancy | `riskScore` vs 50 | consequence |
|---|---:|---|---|
| **Morgan Reyes** (manager hero) | **1** — *too few* | 68, over | **the stated invariant is violated**; her role reads as critical and she is at-risk — she is row 3 of the at-risk table |
| **Pat Ellis** (employee hero) | **3** — *too many* | 45, under | her role stops reading as critical; **both** her signals vanish and `pt-workforce-succession` fails |

**So iter-274's "the fix is one number" is superseded by its own successor.** Holding `DevOps Engineer` at
2 would green the Playthrough and leave Morgan sole-holding a title the mechanism promises she never will —
fixing the tail that has a test and ignoring the tail that has only a comment. **The requirement is a
bound, not a value: a hero's role must have ≥ 2 and ≤ 2 incumbents — exactly one peer.**

That both tails are wrong simultaneously, in the one org anyone looks at, is the strongest available
evidence that this is unenforced rather than mis-tuned. A mis-tuned constant does not miss in both
directions at once.

## Phase 5 — the implementation, specified (and deliberately NOT started)

**What it must do:** exclude hero roles from the ordinary supporting draw, and seat exactly **one**
designated peer per hero role, deterministically.

**Why it needs a signature change.** Designating "exactly one" peer requires knowing the population size —
`memberRoleAt(prefix, storyHeroRoles, i)` receives the member index but not the count, so it cannot ask
*"is `i` the peer for this role?"* without it. Every alternative that avoids the parameter is
probabilistic (a second hash accepting ~n/K members) and would restore the lottery this iter is closing.

**Blast radius, measured:** **6 production call sites** (`users.go:181`, `membership_skills.go:145`,
`population_evidence.go:132`, `certificates.go:150`, `profile.go:323`, `target_roles.go:106`) **+ 5 test
references** = **11 sites**, against a function whose own comment records that the previous unification
*"found only FOUR"* of the six, and that a seventh copy inside `target_roles_test.go` went RED the moment
the production six were unified.

**This iter does not start that edit**, per its declared scope. A partially-swept change to the single
derivation six seeders share is precisely the failure that file documents **twice**, and it would land at
the point in the session where it is least likely to be swept completely.

**Verification is cheap and already established:** iter-273's binding suite is **169 s**, so iter-276 can
grade the change against all 30 live Playthroughs — including `negative-controls.spec.ts:429`, the tenancy
control that caught the *first* cut of this same function (M257x iter-31) when it leaked `DevOps Engineer`
into the contrast org's key-role cards. **That control is the one to watch: it is the reason this
mechanism is per-org at all.**

## Close — 2026-08-10

**Outcome:** The requirement is a **bound, not a value**. Hero-role occupancy is an unbounded uniform hash
draw (mean 2.33, bimodal), and Org A shows both tails failing at once: Pat Ellis over-occupied at 3 (her
Playthrough fails) and Morgan Reyes sole-holding at 1 (`orgRoleSet`'s stated invariant violated). iter-274's
one-number framing is superseded — fixing only Pat's tail would green the test and leave the invariant
broken. The implementation is specified with its 11-site blast radius measured, and deliberately not begun.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5

**Decisions:** `D-M257x-275-1` (the invariant is unenforced in both directions; the fix is a bound and the
edit is deferred with its blast radius measured).

**Side-deliverables:** none.

**Routes carried forward:**
- **`FIX-M257x-275-bound-hero-role-occupancy-to-exactly-one-peer`** — supersedes
  `FIX-M257x-274-devops-occupancy-must-stay-at-two` (**CLOSED**: its one-number framing is wrong).
  iter-276 implements the 11-site change and grades it against the 169 s binding suite, **watching
  `negative-controls.spec.ts:429`** — the tenancy control that caught this function's first cut.
- `ROUTE-M257x-274-successor-half-is-uncovered`, `ROUTE-M257x-274-tie-order-is-unstable` → open.
- Gate **clause 5**, and the inherited queue (`FIX-M257x-269`,
  `ROUTE-M257x-270-directus-consumer-cms-key`, `FIX-M257x-266`, `FIX-M257x-265`, `ROUTE-M257x-h59`,
  `ROUTE-M257x-h65`) → open.

**Lessons:**
1. **An invariant stated in a comment and enforced nowhere will be violated in the one org you look at.**
   `orgRoleSet` promises no hero sole-holds her title; Org A's manager hero sole-holds hers. The promise
   read as a guarantee for four milestones because the only consumer that could have noticed was a
   Playthrough asserting the *other* hero.
2. **When a re-survey says the previous iter's fix framing is too small, that is the iter's deliverable.**
   iter-274 closed with "the fix is one number" on good evidence; one layer down the number is a *bound*
   with two failing sides. Re-survey before implementing is what caught it — the same step that caught
   the stale target in iter-271.
3. **Measure the blast radius before the budget decides for you.** The edit is 11 sites in a shared
   derivation whose own comment records a previous half-sweep. Counting them first turned "start and see"
   into a scoped hand-off.
