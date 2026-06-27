# Release Retro: v1.10 "method acting"

**Shipped:** 2026-06-27 · tag `v1.10` · branch `release/01.10-method-acting` → `main`
**Milestones:** M39 (profile-identity) · M40 (directus-serve-grant) · M41 (profile-depth) · M42e (employee-coverage, gate) · M42m (manager-coverage, gate) · M43 (cockpit-ux) · M44 (profile-completeness) · M45 (generation-engine, gate) · M46 (org-scale-fill, gate)
**Code-of-record:** `rosetta-extensions` @ tags `method-acting-m39` · `m40` · `m41` · `m42e` (`53574ae`) · `m42m-harden-final` · `m43-cockpit-ux-fix1` · `m44-profile-completeness-fix2` · `m45-harden-final` · `m46-servegrant-closure`

## What v1.10 was

The **believable-profile release + the presenter-grade / scalable-generation extension.** v1.9 told the *story*
(the org Workforce dashboard + the verified-skill spine); v1.10 makes each *character* hold up under a close-up.
When a presenter clicks **Login as** a hero, that hero must read as a fully fleshed, believable person — the
individual's **profile** (org name, role + title, work history, education, a real face, deep role-aligned skills)
**and** the content surfaces they land on (**library** + the **activity feed**) populate with real semantic
content, on **every** page a hero of that vantage can reach. The believable-profile half (M39→M42m): M39 profile
identity, M40 the per-stack-Directus serve-grant (the content-surface unblock — the library + activity feed stop
emptying), M41 profile depth, then **two iterative per-vantage Playwright SEMANTIC coverage gates** (M42e employee,
M42m manager). The extension (M43→M46): M43 cockpit UX polish, M44 whole-roster profile completeness (data
density), M45 a cheap-LLM **generation engine** (the first new third-party dep), M46 **org-scale fill** (a whole
~500/735-member org from one supporting-population descriptor) + a gen-batch preview CLI. **Tooling + docs only —
zero platform-repo edits throughout.** Grounded by the live-demo review (`.agentspace/profile_gaps.md` — a hero
logged in as Maya Chen read empty across /profile, /library, the activity feed) + the in-depth root-cause workflow.

## The headline: the M46 org-scale 5-layer activity-dashboard saga (resolved zero-canonical-edit)

The release's hardest moment was M46's **5th gate face**: the M42 Playwright semantic sweep on the **manager**
vantage of a ~500-member org. The manager enterprise grids (`/enterprise/{members,activity-dashboard,settings}`)
refused to hydrate at org scale — and it was NOT one bug. It was **five distinct costs stacked behind one symptom**
(a `…` skeleton / never-resolving query), peeled one layer at a time, each cleared **with zero canonical platform
edits** via the `demopatch` / recapture mechanism:

1. **Sentinel per-row authz fan-out** (members grid) — a per-OBJECT `OrgCheckActionPermission` RPC per row, N+1
   across the federation, 76.7 s @ 500 members. It **can't** be cached object-blind (forbidden-poison correctness
   bug — that attempt was reverted). **FIX B:** drop the read-gate (`roles.go` `checkPermission` short-circuits
   before the RPC; DB roles still render; read-path only — mutations stay enforced) → 0.51 s.
2. **Over-broad fetch + missing FK indexes** (activity-dashboard + settings) — `limit:1000` (ALL members) + two
   unindexed membership joins. **FIX T1:** a next-web pagination demo-patch (`limit:1000→30`) + 2 post-seed
   `CREATE INDEX` → 84 s → ~4 s.
3. **Directus column drift** (cold-only) — the captured per-stack Directus structure had drifted behind prod (cms's
   `SetFields("*", …)` SELECTed `is_interview_validation_enabled`, a column added to prod *after* capture →
   `Directus 500: column does not exist`). **Cache-masked in warm sweeps**; only cold federation exposed it. **FIX
   DD:** a reproducible post-replay `ADD COLUMN IF NOT EXISTS` backfill.
4. **The serve-grant CLOSURE gap** (the deepest layer, `DEF-M46-01`) — cms's `GetJobSimulation` deep-fetch
   traverses target/junction collections (`knowledge_asset`, `sequences_files`/`_2`, `directus_files`,
   `sim_features`, `sim_translations`, `simulations_translations`, `sim_roles_tasks`) the M40 `servedCollections`
   set never registered → Directus dropped the whole parent `sequences` alias → cms `jobsimulation.go:1097`
   **panicked** → null non-nullable `jobSimulation.title` → the activity-table never hydrated. **FIX SG / Path 2:**
   expand `servedCollections` to the 7 closure collections + a synthesized `directus_files` system read-grant, **+
   RECAPTURE** the prod Directus structure over the sanctioned `marco_read` DSN (firewall `public_only=true`, 0
   tenant rows; metadata captured from prod, **never fabricated**). **Resolves DEF-M46-01.**
5. **cms Redis cache poison** (a self-inflicted mid-stream-restart artifact) — a mid-stream cold-restart raced the
   serve-grant apply → cms cached EMPTY-sequences responses in Redis DB-5 → served the poisoned cache (cache-FIRST,
   24 h TTL) even after a fresh fetch would succeed. **A FRESH `/demo-up` never hits this** (it provisions Directus
   with the grant BEFORE cms first queries → caches CORRECT responses from the start). Root-caused + documented as a
   re-up hazard.

After all 5, the definitive cold manager sweep: `failingSections=0, gateMet=true, personaFailures=0, escapes=0,
notReachedPages=0, frontier=EXHAUSTED`; `/enterprise/activity-dashboard kind=real-content`; 0 cms panics; render-
verified (the dan-manager activity table `rowCount=20`, real per-sim success/score/attempts). The campaign ran **4
adversarial sub-agents** + orchestrator audits + a fresh `--purge /demo-up 3` reproducibility proof.

## Incidents across the release (all build-or-close-caught, none shipped)

- **M34-lesson re-applied (M39).** Each schema-touching seeder gets one against-real-schema integration pass; the
  G2 role-backfill silent-0-rows scenario resolved to the documented per-row trigger invariant. 3 harden passes, 0
  production bugs.
- **M41 — empty-eduIDs round-robin modulo-by-zero** (adversarial AR-1) — already guarded by the in-seeder
  `len()>0` check; the test pins the invariant. The never-clobber-verified UPSERT-guard held.
- **M42e — the `networkidle`-never-settles harness flake** dominated the first baseline `(44,1)` — a *moved number
  with no fix is a flake, not a lift*: swapped `waitUntil` → `domcontentloaded` + a bounded settle, then the TRUE
  residual `(8,1)` decomposed into real content gaps (sim-result deep-links, 2 empty skill-paths, 1 external
  article escape), all routed forward WITHIN M42e and resolved to gate-MET `(0,0,0,0,0)`. **Lesson: the protocol's
  measurement convention — distinguish a harness flake from a content lift before you "fix" it.**
- **M42e — the avatar-menu selector org-monogram collision** (the iter-23 false-fail class) — an un-parenthesized
  `&& / ||` precedence bug in the raster-preference predicate; fixed with explicit grouping + a load-bearing
  comment.
- **M43 — the cockpit meta-line double-escape bug** (build/harden-caught) + the `jump_to` href attribute-injection
  invariant + the SRI-typo regression — all pinned by the 2-pass cockpit.py harden (27→63 tests).
- **M44 — the uint64-modulo bank-index regression + the avatar heal-failure propagation** (harden-caught); the
  close-found `/enterprise/members` reads `memberships.picture_url`, not `users.picture` — render-verified (a
  DB-level acceptance had missed the avatar gap; the **render-verify lesson**).
- **M45 — the `genEmail` separator-only address edge bug** (the 5-pass final harden surfaced + fixed it); the
  hallucination-heavy drop-not-fabricate no-leak fuzz + the on-disk-corruption drop-not-crash held.
- **M46 — the org-scale seeding bugs** (the real ~600-member batch surfaced them): the multi-batch cache-index
  collision (lost a whole story's members), name-distinctness at scale (57.7% → a deterministic seed-time
  disambiguator → 100%), an email-distinctness 23505 collision, and the 998 double-size bug (the curated
  `UsersSeeder` also seeds a `size` body) — contained by sizing the descriptor.

## Cross-milestone patterns

- **The CODE-owns-structure / AI-owns-content boundary is genuinely N-invariant.** Built at M45 (every generated
  role/skill routes through the existing resolvers; non-resolving names DROP, never fabricated, closure stays
  GREEN), it scaled UNCHANGED from a bounded N=20 batch to a ~600-member org at M46 — exactly the bootstrap-tok bet.
- **Fixtures-first front-loads the deterministic work.** M45 and M46 both built every code deliverable unit-proven
  with no key and no cost FIRST; the single capped real Azure run answered only the empirical believability
  question last — minimal real spend (`--max-cost`-capped, values-blind).
- **The `demopatch` / recapture mechanism scales to genuinely hard platform walls with ZERO canonical edits.**
  Introduced at M42m (the Studio→prod eject), it cleared a 5-layer M46 org-scale render wall — every fix
  reproducible on a fresh `/demo-up`.
- **The render-verify discipline catches what DB-level acceptance misses.** M44's avatar gap (a `picture_url` vs
  `users.picture` column) and M46's activity-table hydration both needed cockpit-login + Playwright + screenshot,
  not just a DB column read — for any user-VISIBLE deliverable.
- **Decompose a perf wall and try the DROP before declaring a permanent re-scope.** M46 iter-07 honestly closed
  `exit-3 (re-scope-trigger)` reading the manager grid as platform-bound-unfixable; the close SUPERSEDED that by
  decomposing the wall into 5 distinct costs, four demo-patchable. The re-scope was RESOLVED, not escalated.

## Two honest reproducibility caveats (recorded, not hidden)

1. **`down` needs `--purge` to pick up a regenerated cache.** A plain `down`/`up` keeps the demo's Postgres volume
   and no-ops the structure replay — a regenerated snapshot cache (the SG recapture) only lands on a fresh
   `--purge /demo-up` (proven: a `--purge /demo-up 3` came up with the closure collections + the dashboard
   resolving + 0 panics).
2. **A fresh `/demo-up` seeds the default STORIES preset (~341); the org-scale enterprise org (735 generated) is a
   SEPARATE `/stack-seed --gen-batches`.** The org-scale fill is opt-in — `/demo-up` gives the believable Stories
   trio; the ~500/735-member generated org is the explicit `--gen-batches` seed M46 delivers.

## Carried Forward

- **No open deferrals.** `DEF-M46-01` (the serve-grant CLOSURE + recapture) RESOLVED in-release. The standing
  backlog (DEF-M10-01 cloud SnapshotStore / S3 blob bytes · DEF-M21-01 `replayCmd` hermetic test · M25-D9 dev
  taxonomy `rc=4`) is pre-existing cross-release deferral, tracked in `roadmap-vision.md`, orthogonal to v1.10.
- **CI-wiring for the corpus** — there is no executable corpus suite; rext suites verified per-milestone at their
  tags (the local-3x carry-forward — same as v1.5..v1.9). `/developer-kit:project-stats` was not invoked as a
  sub-invocation at this close (a one-line carry-forward, non-blocking).
- **Pushing the ext tags + `main` + the `v1.10` tag to origin** — the orchestrator's separate post-close step (this
  close is LOCAL-only).

## Metrics Delta

- **Go test funcs (rext):** **1248 → 1551 (+303)** — stack-seeding 444→706 (+262), stack-snapshot 333→363 (+30),
  clerkenstein 259→270 (+11, incl the recorded-vs-grep close-drift reconciled to the 270 ground-truth). alignment
  52 + stack-secrets 160 unchanged. The single biggest driver: **M45 stack-seeding +110** (the generation engine:
  `services/ai` + `blueprint` + `batchcache` + `cmd/gen-batch` + the `GeneratedBatchSeeder`).
- **Python / TS:** the cockpit `cockpit.py` suite 27→63 (+36) + the demopatch suite 18→43 (+25); the FIRST non-Go
  rext dev/test dep — the `@playwright/test ^1.49.0` coverage harness (M42e). No suite decreased.
- **Supply-chain:** **1 NEW DEP — deliberate + sanctioned** (`ai@v1.40.1` at M45, the user-acknowledged in-release
  inflection; breaks the v1.8→v1.9 0-new-deps streak by design). M46 reuses it unchanged. All deps MIT/BSD/Apache.
- **Alignment gates:** 100%/100% on all 5 Clerkenstein surfaces (re-verified M42e).
- **Flake:** 0. **Canonical platform-repo edits across the whole release: 0.**
