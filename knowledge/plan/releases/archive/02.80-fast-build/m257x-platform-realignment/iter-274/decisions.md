# iter-274 — decisions

## D-M257x-274-1 — the at-risk gate is role OCCUPANCY, not hero fit; the fix is one seeded number

**Context.** `pt-workforce-succession` is the single failure standing between the suite and gate clause 2
(iter-273, binding: 30 live / 1 failing / 0 error). iter-272 attributed the hero's absence from the at-risk
table to her being well-matched (`fit 87`).

**Measured** (both of iter-272's captured payloads, no new run) — 12 roles, three tiers, and the tier is set
by incumbent count:

| incumbents | `structuralRisk` | `riskScore` (× 0.85) | vs `RiskScore >= 50` |
|---:|---:|---:|---|
| 1 | 80 | 68 | over |
| 2 | 63 | **54** | over by 4 |
| 3–4 | 53 | **45** | **under by 5** |

`DevOps Engineer` — the hero's role — has **3** incumbents, scores **45**, and misses the guard.
`succession.go:790-822` gives the formula exactly; with `ready == 0` and `devel == 0` on every Org A role
(iter-272), the only varying term is the incumbent-count step.

**Decision — iter-272's account is CORRECTED.** `buildAtRisk` never reads fit for membership. Fit decides
whether she can be a **successor**; her role's **occupancy** decides whether she can be **at-risk**. Both
of iter-272's conclusions survive (she appears on no surface; the fix is ours), but the reason was
mis-attributed — and in the direction that made the fix look like skills work rather than a count.

**Corroboration that closes it.** The spec's recorded M256 measurement is
`"Pat Ellis / DevOps Engineer / 40 / Rare skill held only by this person / In fragile role: DevOps
Engineer"`. `buildAtRisk` scores rare-skill **+25** and fragile-role **+15**; **25 + 15 = 40**, and both
signals come from the `RiskScore ≥ 50` sets. So the role was **over** the line then and is under it now:
occupancy went 2 → 3 and took both signals from every member of the role at once.

**The fix:** hold Org A's `DevOps Engineer` occupancy at **≤ 2**. Seed-side, no platform edit.

**This resolves iter-272's seed-vs-expectation fork by ELIMINATION, not preference.** The expectation route
is foreclosed: the page renders no incumbent names, so there is nowhere satisfiable to re-anchor the hero
without deleting the tenant-specificity the M256 sharpening exists to provide.

**What the fix deliberately does NOT close.** It restores the at-risk anchor; it does not populate
`successors` / `topTalents` / `readyCount`, proven empty on every reset. That gap keeps its own route
(`ROUTE-M257x-274-successor-half-is-uncovered`) so a green cannot absorb it silently.

## D-M257x-274-2 — equal-scoring roles are returned in non-deterministic order

**Context.** PR-3 predicted iter-272's two captured payloads would agree exactly. They agree on every
**value** and disagree on **order**: the four roles tied at `riskScore 68` permute between runs.

**Cause.** The projection sorts roles by score with `sort.Slice`, which is **not stable**, so ties are
returned in arbitrary order across runs of identical data.

**Decision.** Recorded as a **latent** hazard, not a defect to fix: nothing asserts on ordinal position
today, because `keyRoleCard` anchors on a role's own heading text and `talentRow` filters by person name.
Routed as `ROUTE-M257x-274-tie-order-is-unstable`.

**Why it is worth an entry.** The obvious way to sharpen a projection assertion is to name the *first* card
or the *top* row — and on this surface that would be flaky in a way that reproduces only across resets,
which is the most expensive kind of flake to diagnose. The note exists so the next person sharpening this
spec does not have to discover it from a red gate.
