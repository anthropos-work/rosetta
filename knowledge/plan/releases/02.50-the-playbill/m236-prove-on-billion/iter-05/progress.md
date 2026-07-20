# iter-05 — tik: the MANAGER vantage

**Type:** tik

## Step 1 — is it content-specific, or demo-wide?

Probed the manager's drill-down for a **roster hero's** session (Maya Chen), not a content-story one. It
failed **identically**: `undefined undefined's Results for …` + `No data`.

That reading initially looked like a **pre-existing demo-wide defect** — and `coverage.spec.ts`'s manager
sample-rule comment even records the symptom as if it were normal (*"the leaf renders only the dashboard
tab chrome, textLen ~70"*). **That inference was premature**, and step 3 overturned it: the hero URL I
constructed by hand used a **user id**, exactly like the manifest's — so the probe reproduced my own bug,
not a platform one. Recorded because the wrong conclusion was one step away and is the kind that gets
written into a doc as fact (which is precisely how the `skipPaths` premise in iter-04 became folklore).

## Step 2 — name the failing query

Captured the network layer on the failing page. Three identical GraphQL errors, `data: null`:

```
Failed to fetch from Subgraph 'backend'.
  error fetching memberships for InsightsByJobSimulations: ent: membership not found
  path: ["insightsJobSimulationBySessions"]
```

## Step 3 — the resolver's actual contract

`app/internal/web/backend/graphql/graph/resolver_queries.go`:

```go
func (r *queryResolver) InsightsJobSimulationBySessions(
    ctx context.Context, organizationID uuid.UUID, jobSimulationID uuid.UUID,
    membershipsID uuid.UUID, options *modelgen.InsightOptions,
) (*org.InsightsJobSimulationBySessionsResult, error) {
    ...
    membership, err := r.app.OrgManager.GetMembership(ctx, membershipsID)
```

**The route's last segment is a MEMBERSHIP id, not a user id.** M233 projected `owner.UserID` there.
`GetMembership` on a user id returns `ent: membership not found`, which errors the **whole query** to
`data: null`.

That single fact explains **both** symptoms at once, and explains the most dangerous property of the bug:
the page's header comes from a **different** query (the sim definition), so it renders fine. The result is
a page that **looks populated** — real sim name, real "2 skills measured" — while proving nothing. Neither
symptom is a missing row; nothing was ever wrong with the seed.

## Step 4 — prove the fix-shape BEFORE writing code

Queried the real membership id for the same user/org and re-visited the same route:

```
before:  /…/ai-simulations/e0ae482f…/d541c46c…   → "undefined undefined's Results …" + "No data"
after:   /…/ai-simulations/e0ae482f…/28f8a5f1…   → "Clara Romano's Results for Business Development
                                                    Manager: craft a winning client proposal"
                                                   Best | Passed | 100/100 | 30min | Feb 12, 2026
```

## Step 5 — the fix (zero platform edits)

Membership ids are already **deterministic** and derived from the same `(prefix, index)` the projection
holds — `deterministicUUID("<prefix>:membership:<index>")` (`users.go:199`). So `ownerSlot` gained a
`MembershipID` computed the same way, and both projections
(`content_manifest.go`, `content_nonsim.go`) emit it. **No new lookup, no DB read at export time.** The
derived id matched the live-queried one exactly.

**The test was asserting the wrong contract** — *"manager path must end in the owner user-id"* — which is
precisely why this shipped. It now asserts the membership id **and** fails loudly if a user id reappears
there, because the positive assertion alone would not catch a well-meaning revert.

The M233 honesty gate then caught the stale checked-in `presets/content-manifest.json` exactly as designed;
regenerated. **16/16 packages green.** Shipped as rext `playbill-m236-manager-membership-id`, published,
`billion` re-pinned, its manifest re-exported and the cockpit restarted so the **served** manifest is the
corrected one.

## Metric

**16 / 31 → 27 / 31.** All 11 "No data" pairs now land.

```
content-stories: LANDED 27 / 31
  simulation:        24/26
  skill-path-legacy:  3/4
  skill-path-new:     0/1

  x2  no "<player>'s Results for <sim>" header      (the two INTERVIEW manager pairs)
  x1  page.goto timeout                             (one skill-path manager route)
  x1  route rendered a not-found                    (academy)
```

## Close — 2026-07-20

**Outcome:** One root cause, correctly located, fixed in tooling with zero platform edits: **+11 pairs**.
The manager vantage works for every non-interview simulation.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (root cause + why the header masked it), D2 (the test asserted the wrong contract), D3 (the premature "demo-wide defect" conclusion, corrected)
**Side-deliverables:** none.
**Routes carried forward:**
- **The 2 INTERVIEW manager pairs** — fail differently (no header at all), on `/activity-dashboard/interviews/…`. Likely a distinct route/tab; cf. M231's `flag_interview_manager_report`. → **iter-06**, handler `INTERVIEW-M236-iter06-manager-report`.
- **The skill-path manager timeout** (1 pair) — persisted through the fix, so it is independent. → **iter-06**, same sweep, handler `SKILLPATH-M236-iter06-manager-timeout`.
- **Academy** (1 pair) — unchanged; needs the catalog+progress fill. → handler `ACADEMY-M236-iterTBD-catalog-fill`.
**Lessons:**
- **A page that renders its header is not a page that works.** The header came from a different query than the payload, so a 3-line "does it look populated?" check passed it. When a surface is composed of independent queries, assert on the **payload**, never on the frame.
- **Prove the fix-shape against the live system before writing the fix.** Constructing one corrected URL by hand took one query and one probe, and converted a hypothesis into a certainty before any code changed.
- **A test can be the reason a bug ships.** This one asserted the defective contract in plain language, so it passed green for a whole milestone. When a test encodes an interface you did not verify end-to-end, it is a hypothesis wearing a test's clothes.
