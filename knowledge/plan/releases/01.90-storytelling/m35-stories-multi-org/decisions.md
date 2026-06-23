# M35 — Decisions

Implementation decisions with rationale (recorded during build). Design-time decisions live in the spec
([`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md), esp. D9–D16).

## D-M35-1 — Multi-org via a normalized `EffectiveStories()` view, not a per-seeder branch

The seeder fleet was built against a single hardcoded `OrgID` const consumed by 6 seeders. Rather than add a
`if multiStory { … } else { … }` branch to every seeder, M35 adds **one normalization layer**
(`blueprint.EffectiveStories()`): a multi-story blueprint resolves to its declared stories; a legacy single-org
blueprint synthesizes **one** story from the root `Org`/`Size`/`Personas`. Every seeder iterates the resolved
slice — **one code path** for both modes. This is why the entire existing single-org preset/dev-min suite stays
green unchanged: the legacy blueprint resolves to a one-story view with **byte-identical** ids (the first story
keeps the bare stack key prefix + `LegacyOrgID`). Rationale: a single normalization point is far less
regression-prone than 6 parallel branches, and it's the natural seam for M36–M38 to build on.

## D-M35-2 — The first story keeps the Clerkenstein default org; later stories get deterministic ids

Clerkenstein resolves every session to **one** `DemoUser` (single-identity until M37). So the FIRST story keeps
`LegacyOrgID` (= `DefaultDemoUser().OrgEid`) + `org_clerkenstein` — a per-launch single-identity demo login
lands in it. Every LATER story gets a deterministic per-story org id (`StoryOrgID(storyID)`, a v5-style UUID).
The `IdentitySeeder` seeds the demo identity **only** for the first story (not a second demo user per story —
that's the M37 multi-identity registry). This is M35's "Clerkenstein org-claim alignment" — a **data-side**
requirement (one seeded org == the demo JWT's claim), NOT a Clerkenstein code change.

## D-M35-3 — Per-story deterministic-id namespacing keyed off the first story's bare prefix

Two stories' populations must not collide (story-A user 1 and story-B user 1 must derive distinct uuids).
`storyKeyPrefix(stack, story)` returns the **bare stack name** for the first story (Index 0) and
`stack:story:<id>` for later ones. Keeping the first story bare is what preserves the M34 single-org ids
byte-for-byte; namespacing later stories is what keeps multi-org collision-free.

## D-M35-4 — #M34-D7 resolved: collision-free declaration-order hero slots + the short-pool flat top-up

Routed from M34 (#M34-D7). Two parts, both landed:
- **Index collision.** M34 hashed each hero into `[1, Size]` — with a trio in a Size-30 story this collided
  ~10% of the time, sharing a manager's population row with an employee's name (caught LIVE in the integration
  test: "tom-struggling and dan-manager hash to slot 14"). **Fix: heroes occupy the first `len(heroes)`
  population slots IN DECLARATION ORDER** (collision-free by construction; the blueprint's `len(heroes) <= size`
  validation guarantees they fit). A non-fatal **warning** still fires for the residual Size<heroes clamp case
  (e.g. Size=1, 2 heroes) — never silent. This is the better fix than just warning-on-hash: it makes the trio's
  rows correct, not just visible.
- **Short role-pool top-up.** M34 kept a short role pool (role-coherence over count). M35's product call:
  **top up from the flat public pool** to hit the declared `verified: N` (`resolveHeroSkills` — role-coherent
  skills FIRST, flat padding SECOND). A hero hits her count even when her role defines fewer skills; closure
  stays green (flat skills are real public node-ids).

## D-M35-5 — Roster role names chosen to RESOLVE in the public taxonomy (O6)

The spec's literal labels for some heroes ("Backend Engineer", "Sales Development Rep", "RevOps Lead") have
**0** public `job_role_skills` (O6 enumeration). A non-resolving role still falls back to the flat pool (closure
green) but is **less role-coherent** (D3). So the shipped `stories.seed.yaml` uses resolving role names —
**Backend Developer** / **Engineering Manager** (Cervato) + **Account Executive** / **Sales Manager**
(Solvantis) — all with 10 public role-skills. Display labels are swappable (D16); role-coherence is the
load-bearing `[ENG]` property.

## D-M35-6 — Supporting-population roles drawn from a runtime resolver, never hardcoded

`memberships.job_role_id` stores the role **node-id form** (`J-XXXXXX-XXXX`, NOT a uuid — O4, verified against
live rows). A fabricated `J-…` would dangle like a fabricated skill node-id. So supporting-population roles are
drawn at RUN time from the **replayed** stack's public `job_roles` that have role-skills (`jobroleref.go`,
mirroring the skill-side `TaxonomyRefs`) — never hardcoded prod node-ids. No replayed taxonomy → roles stay
NULL (no fabrication). Heroes carry their own declared role's node-id.

## D-M35-7 — Manager heroes seed no verified-skill chain (they ride the aggregates)

A `vantage: manager` hero demos the org-intelligence surfaces (Workforce dashboard) — she reads the org
AGGREGATES her employee heroes populate, not her own verified skills. The `PersonaSeeder` skips her chain
(`Verified` is 0 by construction, and a manager validates with verified=0). She still rides a population slot
with her real name. The "coherence property" (the two employees ARE the manager's standout high/low rows) is
realized fully in M36's dashboard; M35 seeds the data it reads.

## Note — negative-modulo panic fixed (job-role pool index)

`int(hashInt(...))` can be **negative** (a uint64 high-bit cast), and Go's `negative % len` is negative → an
out-of-range slice index. The job-role pool's `at()` (drawing a supporting member's role) hit this on the first
multi-org test run. Fixed by normalizing the index in `at()` (`idx += n` when negative). A latent crash on any
real seed run with a population — caught by the new tests.
