**Type:** tik, under `TOK-08`. Route: `ROUTE-M257x-261-succession-projection-is-empty`,
handler `FIX-M257x-261-discriminate-succession-empty`.

# iter-267 — iter-261 offered two causes; the measurement refutes BOTH

## What the iter was asked to do

Discriminate iter-261's two candidate causes for the one failing Playthrough
(`workforce-intelligence.talent-pool.UC1` / `@pt:pt-workforce-succession`) **without bumping the frozen
rext pin**:

1. **Product change** — the advance changed how the succession/at-risk projection is computed, so it no
   longer selects the seeded population.
2. **Seed-contract drift** — `pt-world` at the frozen pin no longer supplies what the projection requires.

The precondition that blocked iter-261 is gone: `stack-dev/app` is at **`3eaadae68`** = current `main`.

## The measurement

`GetSuccession` (`internal/workforce/succession.go:215`) fans out **six** queries under
`errgroup`. Every one was run **verbatim** against `demo-2` for the Playthrough's own org — **Meridian
Labs**, confirmed to be **pt-world Org A** by the pinned manifest
(`playthroughs/manifest/assignment-monitoring.yaml:19`, `studio-builders.yaml:13`):

| `GetSuccession` sub-query | rows for Org A |
|---|---|
| `querySuccessionMembers` | **28** |
| `queryRoleRequirements` | **280** |
| `queryDeclaredSkills` (`public.membership_skills`) | **266** |
| `queryVerifiedSkillsByMember` (validation → session → skills → memberships) | **33** |
| `querySessionActivity` | **89** |
| `queryInterviewSignals` | **12** |

**Every input is populated. Not one query errors. No decommissioned schema exists on the stack.**

## Both candidates are refuted, and the disjunction with them

**Candidate 2 (seed drift) — refuted.** The inputs are present, and the one axis on which drift was most
plausible is closed on both sides: the only commit to touch `succession.go` inside the fold window is
**`65010b59a` (2026-07-23)**, *"repoint workforce + aireadiness reads to public.job_simulation_sessions"* —
it moved three queries off `jobsimulation.*` — and the **pinned** seeder already writes the new location,
saying so in a comment that names the rename explicitly (`stack-seeding/cmd/stackseed/main.go:47-51`:
*"`jobsimulation.sessions` was the ONE rename → `public.job_simulation_sessions`"*). **Reader and writer
agree.** This milestone's founding class does not apply here.

**Candidate 1 (product change in selection) — refuted at the layer it was posed.** The selection predicates
select the seeded population; the numbers above *are* those predicates.

**Therefore the disjunction was false.** iter-261 offered two causes as if exhaustive; the fault lies in
neither, so there is a **third** — above the data layer: the scoring/threshold arithmetic
(`scoreToLevel` / `readinessBucket` / `criticalityMultiplier`), the response caps
(`successionRolesMax = 25`, `successionAtRiskMax = 40`), the API/route, or the frontend.

One plausible mechanism was checked and **excluded**: `g.SetLimit(analyticsQueryConcurrency)` at `:229`
with `analyticsQueryConcurrency = 0` would make `errgroup.Go` block forever — which is exactly what a
**15 s predicate timeout over a page that renders its chrome and never fills** looks like. It is
`manager.go:53` **`= 6`**, equal to the goroutine count. Not a hang.

## Pre-registration grading

| PR | prediction | outcome |
|---|---|---|
| **PR-1** | the projection selects on a verified-skill/assessment artifact | **HELD** — `queryVerifiedSkillsByMember` joins `validation_attempt_skill_results` → `validation_attempt_results` → sessions → `public.skills` |
| **PR-2** | the file was modified inside the window the frozen pin does not cover | **HELD** — `65010b59a`, 2026-07-23, and it is the *only* such commit |
| **PR-3** | the passing siblings do not share succession's predicate | **UNTESTED** — graded honestly. The iter never read `workforce-roster`/`-funnel`/`-org-feedback`'s code paths; claiming it held would be a reconstruction, not a reading (§10, iter-192) |
| **PR-4** | decidable from SOURCE + the seeder alone | **REFUTED** — source and seeder **agree** and jointly predict a *working* projection. A live DB read was needed, and even that did not settle it |
| **PR-5** | the cause is seed-contract drift | **REFUTED** — and so is its alternative. The pre-registration did its job: it made "then it must be a regression" unavailable as a fallback |

## Close — 2026-08-10

**Outcome:** Both candidate causes are **refuted by direct measurement**, and with them the framing that
offered them as exhaustive. The fault is relocated **above the data layer** and the next measurement is
named. No cause is claimed, because none was established — and the frozen-pin control is **unspent**.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-267-1` (both candidates refuted; the disjunction was a framing, not an enumeration).

**Side-deliverables:** none. **Zero writes** — every measurement was a `SELECT` against `demo-2`, which was
not stopped, restarted or re-seeded; `demo-1` untouched throughout.

**Routes carried forward:**
- `ROUTE-M257x-261-succession-projection-is-empty` → **still open, and now correctly aimed.** Handler
  renamed `FIX-M257x-267-capture-the-succession-RESPONSE`. **The next measurement is not more SQL** — it is
  the `GetSuccession` HTTP response for Org A, captured from a logged-in manager session, which
  distinguishes *the backend returned rows and the page dropped them* from *the backend returned an empty
  projection from populated inputs*. Both are now live possibilities and SQL cannot separate them.
- `FIX-M257x-262-dev-path-needs-the-studio-acquisition` (tooling half),
  `FIX-M257x-266-manual-path-drops-gates-the-automated-path-enforces`,
  `FIX-M257x-265-prose-deletion-instructions-are-out-of-D-reach`,
  `ROUTE-M257x-265-stack-demo-carries-six-dead-clones`,
  `FIX-M257x-262-demo-env-append-is-not-idempotent`, `ROUTE-M257x-258-the-pin-is-157-iters-stale` → open.

**Lessons:**
1. **A two-way disjunction offered without a third option is a framing, not an enumeration.** iter-261
   wrote "two candidate causes are named and neither is chosen" — disciplined about not guessing, and
   silently exhaustive about the list. Both were wrong. **Name the residual explicitly**, even as
   *"something else"*, or the next iter spends itself proving a false dichotomy.
2. **A refutation that removes BOTH options is a stronger result than picking one**, and it is only
   available because the pre-registration forbade re-describing a refuted drift hypothesis as a
   regression.
3. **Agreement between source and seeder is evidence the fault is elsewhere — use it that way.** The
   seeder comment naming the exact rename (`main.go:47-51`) is the reader/writer handshake this milestone
   exists to enforce. Finding it *intact* is what let this iter stop looking at the data layer.
