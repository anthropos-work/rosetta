# iter-05 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## D1: content-presence `warmHeavyGrids` vs the networkidle `warmStack` bail — 2026-06-30

**Context:** the manager-vantage authoritative sweep captured skeleton frames on the org-scale `/enterprise/*`
grids → false-fails. The prior warm pushed those grids through `warmStack`, which settles on networkidle ≤4s.

**Options:** (a) raise `warmStack`'s networkidle ceiling; (b) a fixed sleep before the authoritative visit;
(c) a content-presence warm that waits for REAL ROWS up to a generous ceiling.

**Choice:** (c) — a net-new `warmHeavyGrids` (+ `hasRealRows` helper) waiting for real rows (real rows >
skeleton rows, mirroring `isOnlySkeleton`'s selector at page scope) up to a 25s ceiling, polling 1s, manager
vantage only.

**Why:** (a) fails — next-web holds never-idle long-poll connections, so networkidle is the wrong signal
entirely; raising its ceiling just stalls. (b) is brittle (a fixed sleep is either too short at cold or wastes
time when warm). (c) is correct + honest: it primes the COLD federated query (~11.6s at 200 members) so its
result caches before the authoritative visit, AND it's best-effort/bounded — a grid that never paints real rows
within the ceiling is left as-is so the authoritative assert still FAILs honestly (the warm can never mask a
genuinely-empty grid — a STRENGTHENING, not a loosening).

## D2: org-agnostic corporate-email regex replacing the hardcoded `cervato-systems.com` — 2026-06-30

**Context:** the SHARED manager coverage manifest (one manifest, swept against multiple vantage orgs) hardcoded
`cervato-systems.com` (the M42m Cervato calibration org) in `members-roster` + `assign-roster` `mustInclude`.
The M51 sweep runs against **Northwind Aviation** → the literal Cervato domain false-failed the Northwind
roster even though it rendered real member rows.

**Options:** (a) add a Northwind domain alongside Cervato (still org-specific, breaks on the next org);
(b) drop the email assertion entirely (loses the real-email proof → a loosening); (c) assert STRUCTURAL tokens
only + a net-new org-agnostic corporate-email REGEX section.

**Choice:** (c) — `members-roster`/`assign-roster` keep structural tokens (`Members`+`Location`;
`Assign AI Simulation`) + a length floor; a net-new `members-emails` section asserts
`[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}` (`minMatches: 1`).

**Why:** (a) is the same bug deferred to the next org. (b) loosens the gate (a skeleton grid would pass). (c) is
a CORRECTION that PRESERVES the proof org-agnostically: a blank/skeleton grid matches NO email → still FAILs; a
hydrated roster matches every visible member email regardless of org domain. The gate gets stronger and
portable, not weaker.

## D3: the AI-readiness manifest descriptor + `seedPaths` prime — 2026-06-30

**Context:** the M51 seed (config + 200-member funnel, 78.4% all-3, heroes pinned) renders the AI-readiness
dashboard on the live demo, but the route `/enterprise/workforce/ai-readiness` is a `(new)` route group with NO
nav-link (exhaustive grep confirmed) — so the BFS crawl never reaches it, and the gate never ASSERTED the
milestone's headline deliverable (the iter-03/iter-04 "second learning").

**Options:** (a) leave it gate-unproven (rely on a manual eyeball); (b) add it only as a seedPath without
section assertions; (c) add it to `seedPaths` AND author two section descriptors that prove the org-score +
the 3-step funnel render.

**Choice:** (c) — `/enterprise/workforce/ai-readiness` added to `MANAGER_MANIFEST.seedPaths` + two assertions
(`ai-readiness-org-score`: "AI Readiness"/"Overall org readiness"/"Members"; `ai-readiness-funnel`: "Stage
breakdown"/Stage 1/2/3/"Steps completion").

**Why:** seeding the path can't false-inflate (a seed that 404s/redirects is dropped); the two assertions make
the gate PROVE the seeded funnel renders for the showcase org — turning the milestone's deliverable from
eyeballed into gate-asserted. This is the headline gate-strengthening of the milestone.
