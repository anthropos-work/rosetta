# iter-274 — progress

**Type:** tik, under `TOK-08`. Route: `FIX-M257x-272-succession-hero-has-no-qualifying-surface`.

## Phase 1 — pre-registrations sealed

Four, sealed before reading the payloads. See `overview.md`.

## Phase 2 — the 12 roles, ranked

Read from **both** payloads iter-272 captured (no new stack run). Org A, `pt-world`:

| role | `structuralRisk` | × mult | `riskScore` | incumbents | ≥ 50? |
|---|---:|---:|---:|---:|:--:|
| Administrative Coordinator | 80 | 0.85 | **68** | 1 | ✅ |
| Admin & HR Coordinator | 80 | 0.85 | **68** | 1 | ✅ |
| Advanced Analytics Specialist | 80 | 0.85 | **68** | 1 | ✅ |
| Engineering Manager | 80 | 0.85 | **68** | 1 | ✅ |
| Advanced Analytics and Business Intelligence Manager | 63 | 0.85 | **54** | 2 | ✅ |
| Adjunct Professor - Strategy, Leadership And People | 63 | 0.85 | **54** | 2 | ✅ |
| Administrative Virtual Assistant | 53 | 0.85 | 45 | 4 | ❌ |
| Administrator | 53 | 0.85 | 45 | 3 | ❌ |
| Administrative Assistant | 53 | 0.85 | 45 | 3 | ❌ |
| Ad Tech Engineer | 53 | 0.85 | 45 | 4 | ❌ |
| Administration officer | 53 | 0.85 | 45 | 3 | ❌ |
| **DevOps Engineer** *(the hero's role)* | **53** | 0.85 | **45** | **3** | **❌** |

**6 of 12 clear the line. `DevOps Engineer` misses it by 5.**

## Phase 3 — the formula, read rather than inferred

`succession.go:790-822` — and with `ready == 0` / `devel == 0` on every role (iter-272), two of the three
terms are constant across Org A, so the score is a **step function of incumbent count**:

```go
case len(incumbents) == 1: risk += 35
case len(incumbents) == 2: risk += 18
case len(incumbents) <= 4: risk += 8
...
case ready == 0 && devel == 0: risk += 30        // every role
if rareRatio >= 0.5 { risk += 15 }               // every role
multiplier := criticalityMultiplier("unknown", false)   // 0.85, every role — "Phase 3 skipped"
adjustedRisk := int(math.Round(float64(structuralRisk) * multiplier))
```

Which reproduces all three observed tiers exactly: `30 + 15 = 45` base, plus `35 / 18 / 8`:

| incumbents | structuralRisk | `× 0.85` | vs the `≥ 50` guard |
|---:|---:|---:|---|
| 1 | 80 | **68** | over by 18 |
| 2 | 63 | **54** | over by **4** |
| 3–4 | 53 | **45** | under by **5** |

**The guard `RiskScore >= 50` therefore admits roles with ≤ 2 incumbents and excludes roles with ≥ 3.**
The whole question is occupancy, and the boundary is 4 points wide on one side and 5 on the other.

### It reproduces the M256 measurement to the point

The spec's own recorded measurement was `"Pat Ellis / DevOps Engineer / **40** / Rare skill held only by
this person / In fragile role: DevOps Engineer"`. `buildAtRisk` scores rare-skill **+25** and fragile-role
**+15**. **25 + 15 = 40.** Both signals are drawn from the `RiskScore ≥ 50` role sets, so at M256 iter-14
`DevOps Engineer` was **over** the line — i.e. it had **≤ 2** incumbents then and has **3** now.

**Nothing changed in the platform and nothing changed about the hero.** One extra fill-member landed on her
role, occupancy went 2 → 3, `structuralRisk` fell 63 → 53, `riskScore` fell 54 → 45, the role stopped
counting as critical, and **every** member of it — the hero included — lost both role-derived signals at
once. She left the at-risk table as a side effect of a *neighbour's* role assignment.

This is the seed's own recorded hazard, one axis over. Its comment already warns: *"iter-28 D114 paid for
this lesson … appending a hero to Org A displaced Pat Ellis from the member spotlight and turned a
Playthrough RED in all three gate runs"*, and *"a hero's attributes are part of the test SUITE's
contract"* (iter-26 D107, where it was **a role's occupancy perturbing a key-role card** — the same lever,
a different consumer).

### Correcting iter-272

iter-272 wrote that she is excluded from at-risk *because she is well-matched* (`fit 87`). **That is
wrong.** `buildAtRisk` never reads fit for membership. Fit governs whether she can be a **successor**;
occupancy of her role governs whether she can be **at-risk**. Both conclusions iter-272 drew still hold —
she appears on no surface, and the fix is ours — but one of the two reasons was mis-attributed, and it was
mis-attributed in the direction that made the fix look harder than it is.

## Phase 4 — the fix, now specified rather than debated

**The lever is seed-side, single, and needs no platform edit:** hold `DevOps Engineer` occupancy in Org A
at **≤ 2**. The seed already fixes the hero's role (`pt-employee → DevOps Engineer`); what is unpinned is
how many *fill* members are assigned the same role.

This **supersedes iter-272's recommendation** — which offered "seed vs expectation" as a judgement call and
leaned to the seed for coverage reasons. It is no longer a judgement call:

- the expectation route is **foreclosed** — the page renders no incumbent names, so there is no satisfiable
  place to re-anchor the hero without deleting the tenant-specificity the M256 sharpening exists to provide;
- the seed route is **one number**, and restoring it returns the exact row the spec documents (`40 = 25 + 15`).

**It also leaves the coverage gap honestly open.** Fixing occupancy restores the *at-risk* anchor; it does
**not** populate `successors` / `topTalents` / `readyCount`, which iter-272 proved empty on every reset.
That gap is real, it is separate, and it keeps its own route rather than being quietly closed by a green.

## Phase 5 — pre-registrations graded

| PR | verdict | evidence |
|---|---|---|
| **PR-1** — `DevOps Engineer` below the line | **HOLDS** | `riskScore 45` vs the guard's `50` |
| **PR-2** — some roles above it | **HOLDS** | **6 of 12** at 68 or 54 — which is why 27 members carry role-derived signals while the hero does not |
| **PR-3** — the two payloads agree exactly | **REFUTED — benignly, and worth keeping** | every *value* is identical; the **order of equal-scoring roles differs** between runs (the four 68s permute). `sort.Slice` is not stable, so **tie order is non-deterministic**. Harmless today because `keyRoleCard` anchors on a role *name*; a latent flake for any future assertion anchored on first-card or ordinal position |
| **PR-4** — `riskScore = round(structuralRisk × mult)` | **HOLDS** | exact for all 12; `criticalityMultiplier` is uniformly 0.85 because `succession.go:819` records *"Phase 3 skipped → tier unknown, not key"* |

## Close — 2026-08-10

**Outcome:** The hero's disappearance is explained exactly and the fix is reduced to one seeded number.
`RiskScore ≥ 50` admits roles with ≤ 2 incumbents; `DevOps Engineer` went 2 → 3 occupants, dropped 54 → 45,
stopped counting as critical, and took both of the hero's at-risk signals with it. The M256 score of 40
decomposes as 25 + 15 against the same two signals, confirming the model end-to-end. Seed-side, no platform
edit, and iter-272's fit-based explanation is corrected.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-274-1` (the at-risk gate is role OCCUPANCY, not hero fit — and the fix is one
number), `D-M257x-274-2` (equal-scoring roles are returned in non-deterministic order — a latent flake for
ordinal anchors).

**Side-deliverables:** none.

**Routes carried forward:**
- **`FIX-M257x-274-devops-occupancy-must-stay-at-two`** — supersedes
  `FIX-M257x-272-succession-hero-has-no-qualifying-surface` (**CLOSED**: its fork is resolved by
  elimination, not preference). iter-275 implements + verifies against iter-273's 169 s binding suite.
- **`ROUTE-M257x-274-successor-half-is-uncovered`** — new, and deliberately **not** closed by the fix
  above: `successors` / `topTalents` / `readyCount` are empty on every reset, so no Playthrough exercises
  that computation.
- **`ROUTE-M257x-274-tie-order-is-unstable`** — new, latent; no assertion depends on it today.
- Gate **clause 5**, and the inherited queue (`FIX-M257x-269`, `ROUTE-M257x-270-directus-consumer-cms-key`,
  `FIX-M257x-266`, `FIX-M257x-265`, `ROUTE-M257x-h59`, `ROUTE-M257x-h65`) → open.

**Lessons:**
1. **When a threshold guards a signal, measure the population against the threshold before theorising
   about the subject.** Two iters reasoned about *the hero* — her fit, her skills, which table she belongs
   in. The answer was a property of **her role's other occupants**, and it was one `sort` away in an
   artifact already on disk.
2. **A recorded measurement is a decomposable fact, not a decoration.** The spec's `40` sat in a comment
   for two milestones. Read against `buildAtRisk`'s `+25` / `+15`, it names both missing signals and dates
   the regression — the single most load-bearing number in this investigation was already checked in.
3. **A correction to a prior iter belongs in the iter that finds it, stated plainly.** iter-272's fit-based
   account was wrong and would have sent iter-275 hunting for skills to add. Naming it cost three
   sentences.
