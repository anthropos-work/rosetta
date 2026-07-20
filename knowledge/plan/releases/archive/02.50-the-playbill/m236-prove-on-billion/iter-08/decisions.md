# iter-08 — decisions

## D1 — the academy CTA pointed at a route that does not exist

**Evidence.** ant-academy's full route list contains `app/(public)/library/page.jsx` — the library
**index** — and no `[slug]` child. The per-course route is `app/(authed)/courses/[slug]/page.jsx`.
Separately, the slug `foundation-of-artificial-intelligence` does not appear in the catalog the demo
academy serves (2705 chapter entries / 64 skill paths from the repo's committed FS content). Live:
`/library/foundation-of-artificial-intelligence` → **404**; `/courses/ai-foundations` → **200**, 13 chapter
links, 0 draft chips.

**Decision.** Exhibit `academy-foundation-of-ai` → `academy-ai-foundations`, slug `ai-foundations`, route
`/courses/<slug>`.

**Why the error survived to M236.** The M235 seeder comment asserted the route "resolves,
non-fabricated" — written offline, never driven. And `content_nonsim_test.go` *required* the `/library/`
prefix, so the unit test actively defended the 404. Inverted to guard `/courses/`.

## D2 — `academy-seed` is moot in a demo; the routed-forward plan was unachievable as written

**Evidence.** The academy process (`pid 4034369`, cwd `…/ant-academy/code`) runs with
`ACADEMY_DEMO_FS_PUBLISHED=1` and **no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`**. `getBackendCatalogView`
therefore returns null unconditionally and the catalog is served from the filesystem by the M230 patch.

**Consequence.** Academy progress rows written to the demo DB by `app/cmd/academy-seed` have no reader —
the academy has no backend connection at all. The plan item "wire `academy-seed` into the cold bring-up …
then re-point the CTA to the authed progress-bearing chapter route" cannot be satisfied without first
pointing the academy at the demo's GraphQL endpoint.

**Decision.** Do not wire `academy-seed`. The academy content story is a course page rendering real
content — honest, and what the gate asks for ("the academy grid renders real cards"). Wiring the academy to
the demo backend is a bring-up change; recorded here, not smuggled into this iter.

## D3 — `player-academy`, the sixth render shape

`/courses/…` fell through `shapeFor` to `player-scored`, which demands feedback / evaluated-skills. A
course landing page has neither, so a correct 3744-char render was graded a failure. Added
`player-academy`, selected **by route**, asserting course/chapter structure and a length floor.

**The Draft-chip check lives inside this assertion**, not in a separate sweep: M230's gate is
*production-faithful* rendering, and a Draft chip means the demo filled its grid from the dev draft layer
rather than the published path. Making it part of the same assertion means the pair cannot pass while
looking un-production-faithful.

## D4 — Thread A is met; M230 carry-forward cluster 1 is discharged

Measured live on `billion` in a browser as a member: the academy home grid renders **65 course links / 483
chapter links / 0 Draft chips**. M230 closed `closed-incomplete` with exactly this measurement outstanding
("a rendered-card count on a cold `/demo-up`"), Fate-3'd to M235/M236. It is now taken, on the live VM.

**Cluster 3 is NOT discharged.** The anonymous `/library` index renders 0 cards, because
`getPublicCatalogView` takes the `new Set()` branch that the M230 patch does not cover — a gap the patch
manifest names itself. Out of this milestone's gate (which asks for the academy *grid*, the authed home
surface). Routed forward as `ACADEMY-M236-iter08-public-catalog-twin` with the fix shape already known: a
twin manifest of the same FS-published transform.

## D5 — side-discovery, recorded not fixed: `apps/web` client GraphQL endpoint looks wrong

`apps/web` (pid 4032923) carries `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql` — the
**non-offset** port, which does not exist for demo-1 — while `apps/hiring` carries the correct
`https://billion.taildc510.ts.net:15050/graphql`. SSR uses `WUNDERGRAPH_SSR_ENDPOINT` and is unaffected,
which is why all 26 simulation pairs pass.

Not investigated here (scope-creep tripwire — this iter's planned lines were the academy pair and Thread
A). Flagged to `LATENCY-M236-iterTBD-hero-p95` because a client-side fetch to a dead address is precisely
the *fast-failing* arithmetic signature `latency-budget.md` teaches to read before attributing a leg.
