# iter-272 — progress

**Type:** tik, under `TOK-08`. Route: `FIX-M257x-267-capture-the-succession-RESPONSE` (gate clause 2).

## Phase 1 — pre-registrations sealed

Four, sealed in this iter's first commit before any measurement. See `overview.md`.

## Phase 2 — the re-measurement

`./run-playthroughs.sh 2 --reset --grep '@pt:pt-workforce-succession'`, twice, each with a full
reset-to-seed of `pt-world` (FK-ordered TRUNCATE then fresh seed). **Both runs failed identically** —
20:36:27Z→20:36:49Z and 20:39:56Z→20:40:17Z, `PT_RC=1` both times. Deterministic, not a draw.

**It fails at line 96 — the hero-row assertion — and nowhere earlier:**

```
Error: the succession/at-risk projection names the org's seeded hero (Pat Ellis)
expect(received).toBeGreaterThan(expected)   Expected: > 0   Received: 0
Timeout 15000ms exceeded while waiting on the predicate
```

### The inherited characterization is REFUTED

*"Correct chrome, empty projection tables"* is wrong, and it mattered — it is what pointed two iters at the
data layer. Everything before line 96 **passed**: the login, the URL, the `Succession Planning` heading,
`successionSignal()`, `atRiskSignal()`, `roleCandidateRows().first()` visible with `count() > 0`, and
`keyRoleCard("DevOps Engineer")` — the org's **own** seeded role. The page snapshot shows **12 role cards**
and a **27-row** at-risk table paginated across 3 pages.

The tables are not empty. **One row is missing, and it is the hero's.**

## Phase 3 — the capture

### PR-3 is REFUTED, and the reason is the finding

The trace contains **5** client-side GraphQL calls to `:28082` — `billing`, `userPreferences`,
`organizationSettings`, `setUserPreferences`, `userMemberships`. **None is the succession query.** The
projection is fetched **server-side** by next-web, so no `GetSuccession` response body ever crosses the
browser boundary and the prediction as written cannot hold.

Its *conclusion* held anyway, by a route the prediction did not name: the trace retains the **RSC payload**
as a resource, and the payload carries the projection in full. **No new instrument was needed — but not
for the reason predicted**, which is exactly the distinction the milestone's own rules ask to be kept.

### What the wire actually carries (both runs, identical)

| field | value |
|---|---|
| `roles` | **12** |
| `successors` | `[]` — **12 of 12** roles, zero non-empty |
| `readyCount` / `developableCount` | **0 / 0** — 12 of 12 |
| `topTalents` | **empty** (`"topTalents":0`; the page renders *"No ready successors identified"*) |
| `atRisk` | **27 entries** — Hassan Larsen, Chloe Mensah, Morgan Reyes, … Bruno Schmidt |
| **`Pat Ellis` in `atRisk`** | **NO** |
| `Pat Ellis` in the payload | **twice** — as an `incumbents` entry of `DevOps Engineer` (**`fit: 87`**), and in `allMembers` |

The org has **28** members (iter-267's `querySuccessionMembers`). **27 are at-risk. Pat Ellis is the single
exclusion** — and the at-risk rows say why: *"Skills not aligned with role (fit 0%)"*, *"(fit 11%)"*. She
is the only well-matched person in the org.

**PR-4 HOLDS.** The successor half of the projection is empty on the wire, from fully populated inputs, and
the renderer is faithful: `0 ready` on every card is what it was given.

## Phase 4 — the residual, partitioned

`app/internal/workforce/succession.go:741-753` builds each role's candidate set:

```go
if _, isIncumbent := incumbentIDs[m.MembershipID]; isIncumbent { continue }
if m.SkillsCount == 0 { continue }
f := computeFit(m.SkillsMap, effective)
if f == nil || f.Fit < 40 { continue }
```

A successor must be a **non-incumbent** with **fit ≥ 40**; `readinessBucket` then calls **≥ 70** *ready*.

So `successors: []` across all 12 roles **entails** — given the code, it is a derivation and not an
inference — that for every role, every non-incumbent member has `fit < 40` or no skills. Which the at-risk
fit percentages independently corroborate.

**Therefore the assertion is unsatisfiable on this seed, and nothing is broken.**

- Pat Ellis cannot be **at-risk**: she is the one person whose skills match her role (`fit 87`).
- Pat Ellis cannot be a **successor**: incumbents are skipped, and `DevOps Engineer` is her own role.
- No one else can be a successor anywhere: every other member is below the fit-40 floor.

The Playthrough asks for the hero in *"the succession/at-risk table"*. The projection places her in the
**third** structure — `incumbents` — which the assertion does not read.

**Why it used to pass.** The spec's own comment records the M256 iter-14 measurement:
*"Pat Ellis / DevOps Engineer / 40 / Rare skill held only by this person / In fragile role: DevOps
Engineer"*. She was at-risk **at fit 40**, on the boundary. She is now at **87**. Her fit rose — the seed
now gives its thriving hero verified skills that match her role — and **the surface she qualifies for
changed with it.**

### What this means for the fix, and why it matters that it is ours

The fault is **not in the platform**, so v2.8's *0 platform edits* constraint does not bind here — which
was the live risk when the residual still included "the projection arithmetic". The fix surface is one of
two things we own:

1. **the seed** — give at least one non-incumbent a ≥ 40 fit to a role, so the successor projection is
   exercised at all (today the entire `successors`/`topTalents`/`readyCount` half of this surface is
   proven empty on every pt-world reset, so no Playthrough covers it); or
2. **the expectation** — assert the hero where the projection actually puts her, as the **incumbent** on
   her own key-role card.

They are not equivalent, and the choice is a real one: (2) turns the Playthrough green tomorrow; (1) is the
only one that makes the suite *cover* the successor computation. **(2) alone would green a surface that has
never once produced a successor** — the shape of green-but-wrong this milestone keeps finding. That
decision belongs to the next iter, with this evidence in front of it.

## Phase 5 — pre-registrations graded

| PR | verdict | evidence |
|---|---|---|
| **PR-1** — it still fails | **HOLDS** | 2 of 2 runs, `PT_RC=1`, each after a full reset-to-seed |
| **PR-2** — fails on a projection/row assertion, not login/URL/heading | **HOLDS** | fails at spec line 96 (the hero row); login, URL, heading, succession signal, at-risk signal, role rows and the org's own key-role card all **passed** |
| **PR-3** — the trace carries the succession GraphQL response | **REFUTED** | the trace's 5 client GraphQL calls are `billing` / `userPreferences` / `organizationSettings` / `setUserPreferences` / `userMemberships`. The projection is fetched server-side. Its *conclusion* survived by another route — the RSC payload — which is the part worth keeping |
| **PR-4** — the wire is empty | **HOLDS** | `successors: []` and `readyCount: 0` on **12 of 12** roles, `topTalents` empty, from six populated inputs |

Three held, one refuted — and **the refuted one carried the iter.** PR-3's failure is what turned a
missing-response hunt into a payload read, and the payload is what made the answer decidable in a single
artifact.

## Close — 2026-08-10

**Outcome:** The route's measurement landed and **decided**. The failure is not an empty projection, not a
dropped row and not a platform defect — the hero is the org's only well-matched member, so she is excluded
from at-risk *correctly*, and she cannot be her own role's successor *by design*. Gate clause 2's last
failing Playthrough now has a named, in-scope fix surface and a real choice between two fixes.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-272-1` (the failure is a seed/expectation mismatch, not a defect — and the fix
surface is ours), `D-M257x-272-2` (a server-rendered surface is invisible to a browser trace's network log
and visible in its RSC payload — capture the payload, not the calls).

**Side-deliverables:** none.

**Routes carried forward:**
- **`FIX-M257x-272-succession-hero-has-no-qualifying-surface`** — new, replaces
  `FIX-M257x-267-capture-the-succession-RESPONSE` (**CLOSED** by this iter). The fix decision above, with
  the recommendation stated: **prefer (1), the seed**, because (2) alone greens a computation that has
  never produced a non-empty result.
- Gate **clause 5** — the documentation-accuracy reading, unmeasured since iter-131 (`P = 29 / N = 47`).
- `FIX-M257x-269-force-append-grows-the-demo-env-without-bound`,
  `ROUTE-M257x-270-directus-consumer-cms-key-outlived-its-rollback-path`,
  `FIX-M257x-266-manual-path-drops-gates-the-automated-path-enforces`,
  `FIX-M257x-265-prose-deletion-instructions-are-out-of-D-reach`,
  `ROUTE-M257x-h59-rext-edits-fire-no-fence-anywhere`,
  `ROUTE-M257x-h65-fresh-checkout-class-needs-a-scheduled-remeasure` → open.

**Lessons:**
1. **A symptom description is evidence too, and it decays.** *"Empty projection tables"* was carried across
   three iters and sent two of them into the data layer; the tables held 12 role cards and 27 at-risk rows
   the whole time. **Re-derive the symptom before re-deriving the cause** — the first thing this iter did
   was re-run the failure, and that is where the correction came from.
2. **Ask what the assertion is a claim ABOUT, not just whether it passes.** *"The hero appears in the
   succession/at-risk table"* looks like a claim about the projection; it is a claim about **which bucket
   the hero falls into**, and that moved when her fit moved. An assertion pinned to a boundary value (40,
   against a ≥ 40 floor) was one seed change from flipping.
3. **A test going red because the product got BETTER is a real category.** The hero's fit rising from 40 to
   87 is the seed doing its job. Nothing regressed; an expectation aged.
