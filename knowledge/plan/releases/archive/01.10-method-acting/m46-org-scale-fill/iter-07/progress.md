**Type:** tik (#6, under TOK-01) — the FINAL gate face: seed the populated org + the M42 semantic-coverage sweep

# M46 · iter-07 — the 5th gate face: seed the populated org + the M42 semantic-coverage sweep

Run 1 (iter-06) landed the org-scale engine machinery + 2 scale fixes and proved 4 of the 5 gate faces on a
real 614-member Azure batch. This tik closes the **5th and final gate face** — seed the populated org from the
complete cache, run the **M42 Playwright semantic-coverage sweep** on the generated population, prove closure
GREEN + 0 collisions — and surfaced + fixed **2 more org-scale bugs** the seed + sweep exposed.

## What landed (rext, tag `method-acting-m46-iter07`)

### 1. Seed-time name-distinctness backstop (`seeders/`) — the `$0` cache-hit reseed face
The gen-time disambiguator (`cmd/gen-batch genOne`) only fires on a cache **MISS**; a `$0` cache-hit reseed
read `env.Name` **verbatim**, and the complete 614-cache was only **57.7% distinct** (the disambiguator landed
AFTER the cache completed in iter-06). A `$0` reseed of that cache would seed duplicate identities → FAIL the
persona/believability bar. Fix: apply the **same deterministic disambiguator at SEED time** in
`GeneratedBatchSeeder`, moved into the `seeders` package as the ONE source of truth
(`DisambiguateGeneratedName` + `DisambiguateSurnames`); `cmd/gen-batch`'s gen-time path now delegates to it
(no pool drift). 3 tests + full stack-seeding suite green. Commit `d466f4b`.

### 2. Email-distinctness backstop (`seeders/`) — surfaced by the real seed
The real `$0` reseed then hit `user_basic_info_email_key` (`UNIQUE(email)`, SQLSTATE 23505): name-distinctness
alone doesn't imply email-distinctness — the cache had 71 duplicate `email_local` groups AND 4 cross-name
collisions (distinct names → the same local part, e.g. "Jinwoo Park" / "Jin-woo Park" → `jinwoo.park`). The
colliding email aborted the entire generated-batch users COPY (0 rows). Fix: a SECOND distinctness axis
(`usedEmails` + `uniqueEmail` — index a colliding local part `local+globalIdx@domain`, per-domain). 2 tests +
full suite green. Commit `d5ae926`.

### 3. The org-scale heavy-grid settle fix (`stack-verify/e2e/`) — surfaced by the sweep
The first manager sweep verdict was `failingSections=3` — `/enterprise/members` (skeleton), `/enterprise/
activity-dashboard` + `/enterprise/settings` (empty). The screenshot confirmed the diagnosis: full chrome +
Cervato org name/logo + table headers, but **skeleton rows + a `0 / ∞` count** — the **998-member grid hadn't
hydrated** within the harness's *single* bounded re-assert (calibrated against the pre-M46 221-member org;
~4.5× heavier now). Data IS present (verified in DB) — a settle/timing measurement bug, not a content gap. Two
complementary harness fixes (the protocol's "bound the settle to the heaviest DATA GRID" lesson at org scale):
(a) `section-assert.ts` — the single re-assert became a **bounded re-assert POLL** (up to 6×, return on the
first pass, poll ONLY a skeleton/empty paint-timing kind, genuinely-empty still fails after the budget = no
false-pass); (b) `coverage.spec.ts` — the **manager warm set** now ALSO warms `/enterprise/{members,
activity-dashboard,workforce}` (the cold-cache penalty removal).

## The seed (the populated org — demo-3, `$0` cache-hit reseed) — CORRECTED to ~500 (the 998 double-size fix)
> **Re-seed update (recovery-continuation).** The first iter-07 re-seed produced a **998-member** org (the
> double-size bug: the curated `UsersSeeder` seeds a full `size` body AND the `fill:true` batch adds another
> `size − heroes` generated members → ~2×`size`). The fix sizes the descriptor's `size` to 250 (Cervato) / 120
> (Solvantis), and the re-gen cache is the matching **364-member** batch (247 Cervato + 117 Solvantis). The
> numbers below are the corrected ~500-member org.

- Reset (`stackseed --stack demo-3 --reset`) then seed (`stackseed --stack demo-3 --gen-batches --cache-root
  .agentspace/.batchcache`, a `$0` cache-hit — gen-batch `--preview` confirmed **364/364 cached, $0**). 21
  surfaces ok, audit **49 writes / 60267 rows**, **isolation: clean (no shared/external writes)**.
- **735 users / 735 memberships, 735/735 DISTINCT emails** (the email backstop — UNIQUE(email)-clean), **717
  distinct full-names** (18 realistic dupes, intentional), **0 HERO collisions**. Avatars: every population
  member has a real photo + membership `picture_url`; only the 1 Clerkenstein login identity
  (`demo@anthropos.test`) is photo-less (expected). Org sizes: **Cervato 498** (~250 curated + ~247 generated —
  the ~500 headline gate org) + **Solvantis 237** (~120 + ~117) — both well under the ~1k that walled the
  client render.
- Curated heroes intact post-reseed: **Maya Chen 36 verified / 54 total** (the skill-spotlight chain), Tom
  Becker 6/24, **Dan Rossi 15/18 = Engineering Manager / admin** (the manager vantage), Sara 36/50, Nick 6/22,
  Leah 15/18.
- Sentinel casbin **RELOADED** after the reset+seed (`AuthorizationService/Reload` on :38087 → 2xx; the
  reset deleted 621 g2 grants, the seed re-created g2=371/g3=621, the reload refreshed the in-memory enforcer —
  no silent 403s; the manager/employee sweeps reached `/enterprise/*` proving the reload landed).
- `datadna measure-closure --stack demo-3` = **PASS** (GREEN — every seeded verified-skill node-id resolves).

## The 5 gate faces (on the corrected ~500-member populated org)
| Face | Result | Verdict |
|---|---|---|
| believable role/avatar/skill spread (not 90% hollow) | 735 members, 717 distinct names, real photos, role-coherent skills, persona-assert `ok` | **PASS** |
| hero-name collisions at scale | **0** | **PASS** |
| closure GREEN (`datadna measure-closure`) | PASS (every seeded verified-skill node-id resolves) | **PASS** |
| throughput + cost within budget | 364/364 cached, **$0** reseed (gen across prior capped runs) | **PASS** |
| **M42 semantic-coverage gate PASSES on the generated population** | manager `failingSections=3` (the enterprise grids never hydrate) | **FAIL — platform-bound** |

## Re-measure (the gate's primary metric — the M42 sweep on the populated org)
- **Manager (Dan) sweep:** `reachable=58/150 failingSections=3 personaFailures=0 escapes=0 notReached=0
  crossPortFollowFails=0 frontier=EXHAUSTED → GATE NOT MET`. The 3 fails are `/enterprise/members`
  (`members-roster`), `/enterprise/activity-dashboard` (`activity-table`), `/enterprise/settings`
  (`org-settings`) — all `kind=empty` AFTER the full 6× warm-grid poll. Persona (`role-skills-coherence`,
  `avatar-consistency`, `org-identity`) all `ok`; the studio-desk cross-port follow `ok`.
- **Root cause (decisive, see decisions.md D3):** the three sections show a `…` loading spinner / skeleton +
  `0 / ∞`, i.e. a **never-resolving federated GraphQL query**, NOT a content gap (498 members + 1000 events
  ARE seeded; raw SQL is 31–121 ms). The Cosmo router logged **11 requests at 10 s+** in the sweep window,
  peaking **83.9 s / 80.4 s / 60.0 s** — the per-row resolver fan-out (`jobRole`/`targetRole`/`tags`/
  `lastActivityDate`/`organizationFeatures` × Sentinel authz) is an N+1 across the federation. **The
  998→500 resize barely moved it (10.88 s → 10.5 s)** ⇒ member-count is not the dominant factor. The manager
  gate last PASSED at **221 members** (M42m, `63d3fad`); org-scale fill inherently pushes past that.
- **Employee (Maya) sweep: GATE MET** — `reachable=59/150 failingSections=0 personaFailures=0 escapes=0
  notReached=0 crossPortFollowFails=0 frontier=EXHAUSTED`. `/home`, `/profile`, `/profile/skills`,
  `/profile/activities`, `/library/*` all PASS `[ok]`; PERSONA `role-skills-coherence=ok avatar-consistency=ok
  org-identity=ok`; 2 disclosed presenter-note external refs (allowed). **This is the decisive contrast:** the
  employee vantage does NOT touch the manager enterprise grids, and it PASSES cleanly — so the seeded
  ~500-member org is **fully believable**; the blocker is **isolated to the 3 manager-only enterprise grids**
  whose federated GraphQL queries are platform-bound-slow. The seed is correct, not the cause.

## Close — 2026-06-26 (recovery-continuation)
**Outcome: `closed` (diagnostic iter) + RE-SCOPE-TRIGGER.** Four of the five gate faces PASS on the corrected
~500-member org (believable spread, 0 collisions, closure GREEN, cost/throughput in budget). The 5th face — the
M42 manager sweep — is **platform-bound and unsatisfiable** at org scale: the enterprise org grids
(members/activity-dashboard/settings) do not hydrate org-scale data within any reasonable window (10–84 s
federated-resolver N+1, invariant to the resize), and a platform resolver fix is forbidden (zero-canonical-edit
line) while a `demopatch` of the N+1 is out of scope + would fake the gate.

**What LANDED + is committed (the contained, defensible work):**
- the population-math fix (descriptor sized to ~500, `StoryHasFillBatch` helper) + the seed-time name/email
  distinctness backstops (`d466f4b`, `d5ae926`, prior to this continuation);
- the warm-grid harness fix (`section-assert.ts` bounded re-assert poll + `coverage.spec.ts` vantage-aware
  warm) — correct org-scale extensions of the settle lesson, kept;
- the 3 corpus docs (ai-generation-spec / cache-spec / coverage-protocol) documenting the `$0` distinctness
  backstop + the org-scale settle lesson.

**Routed forward (re-scope):** the M46 exit gate's *"M42 sweep PASSES on a ~500 org"* criterion must be
re-scoped by the roadmap owner — measure org-scale believability on the surfaces that DO render at scale (the
employee profile, the seeded population correctness via DB/closure) and treat the enterprise members/activity
grids as a **documented platform-perf exception at org scale**, OR cap the headline org at the platform's
render threshold (~221) and prove org-scale fill on the generated population's correctness rather than the
enterprise-grid render. The named population-math refactor (heroes-only-`UsersSeeder`) does NOT resolve this —
it fixes org=`size` (not 2×`size`) but not the grid wall. **Not faked.**

---

### Standardized close fields (build-mstone-iters Phase 4)
**Outcome:** 4/5 gate faces MET; the 5th (M42 manager sweep) is platform-bound-blocked — re-scope-trigger raised.
**Type:** tik (#6 under TOK-01)
**Status:** closed-fixed-partial — the planned contained deliverables (the ~500 population-math fix + the
warm-grid harness fix + the re-seed/closure proof) LANDED and are clean; the gate's 5th face could not be
closed because it is platform-bound (federated enterprise-grid resolver perf, forbidden to fix), routed forward
as a gate-criterion re-scope. The employee M42 sweep is GATE MET; the seed is correct.
**Gate:** NOT MET (manager M42 sweep `failingSections=3`, platform-bound; employee M42 sweep MET)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: **y** — (4)
user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: **exit-3 (re-scope-trigger)**
**Decisions:** D1 (contained population-math fix), D2 (warm-grid harness fix landed but didn't close the gate),
D3 (RE-SCOPE-TRIGGER — the 5th gate face is platform-bound).
**Side-deliverables:** the 3 corpus-doc honesty corrections (ai-generation-spec org-composition + platform-limit
caveat; coverage-protocol the never-resolving-query corollary) — committed with the iter, not folded into status.
**Routes carried forward:** the M46 exit-gate re-scope (owner decision) — measure org-scale believability on the
surfaces that render at scale + document the enterprise members/activity grids as an org-scale platform-perf
exception, OR cap the headline org at the platform's render threshold. The named heroes-only-`UsersSeeder`
refactor does NOT unblock the gate (fixes org=size, not the grid wall).
**Lessons:** at org scale a never-resolving federated query is a PLATFORM limit, not a slow-paint — diagnose by
query latency (`docker logs graphql | grep latency`) + raw-SQL timing + org-size-invariance, never widen the
poll to mask it nor shrink the org below the org-scale premise. (Captured in coverage-protocol.md, same release.)
