---
title: "KB Fidelity Audit — M256 «playthrough sharpening»"
date: 2026-07-28
scope: milestone:M256
invoked-by: build-mstone-iters (Phase 0b pre-flight gate)
---

## Verdict

**YELLOW — proceed with tracking.**

No topic is left unanchored and **every stale load-bearing claim found was fixed inline in this pass** (6 docs
edited). The three blind areas are blind areas the milestone *exists to fill*, and its `overview.md` already
carries the `Delivers → corpus/ops/demo/playthroughs.md` annotation the gate accepts in lieu of a pre-authored
doc — plus an `iterative` shape whose stated rationale is that the coverage clusters are "unpriced until driven
live."

**Not RED, but read Open Items 1 first.** The audit surfaced one fact that changes what exit-gate clause 2 can
mean: the `entitlement` axis is a **label**, materialized by no seeder, so *"≥ 1 `blocked` outcome"* has no
seedable mechanism today. That is a scope decision for the bootstrap tok, not a missing contract — but it must
be decided before iter-01 commits a strategy.

**Environment:** local only. `billion` / `demo1.anthropos.work` was not contacted (standing sign-off rule);
every finding below is from static reads of `/Users/kirality/Workspace/anthropos/rosetta` and its git-ignored
`.agentspace/rosetta-extensions` clone (HEAD `6ca8764`, tag `fast-build-m255-close`, tree clean).

### ⚠️ Reconciliation — this gate ran CONCURRENTLY with iter-01, and refutes one of its decisions

The bootstrap tok wrote `iter-01/` and `evidence/playthrough-map.md` §8 at **12:52–12:55**, while this audit
was in flight (report written 13:12). The gate is specified to run *before* the tok authors its strategy; it
did not, and the consequence is exactly the failure mode the gate exists to catch:

| iter-01 decision | This audit |
|---|---|
| **D4 — "The `blocked` outcome needs no seed work"** — premised on `seed-worlds.yaml` declaring `pt-free`, the `free` entitlement and the `entitlement-gated` capability, closing with *"`ptvalidate`'s precondition-coverage check will resolve it on day one."* | 🔴 **REFUTED — see F4.** The declaration is real; the **seeded tier is not**. `blueprint.TierMix` is consumed by **no seeder** and never reaches a DB column; `pt-world.seed.yaml` declares no `tier_mix`; `pt-free` is seeded as an ordinary user with a membership, not a tier-gated one. D4's closing clause is *true and is the trap*: `ptvalidate` **will** resolve `entitlement-gated` on day one, because it only checks the string against the world's own declaration — a **fail-open**. A spec written on D4 would assert against a gate that does not exist. **D4 must be re-decided before any `blocked` Playthrough is authored** (Open Item 1). |
| **D1 — the parallelism enabler leaves the critical path** (clause 1 is a per-test median; the Clerkenstein refactor is priced and routed forward, Fate 3) | ✅ **CONFIRMED and strengthened.** D1's pricing of the seat property matches the code exactly (F3). This audit adds what D1 needed and the corpus lacked: the property is now *documented* in `clerkenstein.md` + `playthroughs.md`, including the point neither carried — `storageState` reuse does **not** isolate seats, so L3 (per-seat reuse) must stay within a seat-grouped serial order, as D1/§8.3 already assumes. |
| **D3 — the clause-1 lever is the residual `networkidle`, not parallelism** | ✅ **CONFIRMED and extended.** Same two `goto` overrides (`skill-path-page.ts:31`, `simulation-page.ts:36`) and the same `cockpit-login.ts` default. This audit adds **6 further unbounded `waitForLoadState('networkidle')` sites** D3 does not name (gap 4) — so the widened fence D3 proposes must cover unbounded settles too, not just `goto` + `loginAsHero`. |
| **D2 / §8.5 — onboarding is "one seed question gating up to 5 UCs"** | ✅ **The question is now ANSWERED: no** (F5). No onboarding field exists anywhere in `stack-seeding/` and `UsersSeeder` writes a membership for every user unconditionally. The cost is a **seeder + a `capabilities:` entry + a roster seat**, not a lookup. D2's ordering rule (org-admin first) is *reinforced* by this. |

---

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Playthrough count / corpus | `corpus/ops/demo/playthroughs.md` | `playthroughs/manifest/*.yaml`, `e2e/tests/*.spec.ts` | **PAIRED** |
| Manifest schema + 4-state map | `playthroughs.md`, `spec-drafts/playthroughs/spec.md` | `manifest/manifest.go`, `validator.go`, `report/` | **PAIRED** |
| `@pt:` tag registry + grammar | `playthroughs.md` | `cmd/ptvalidate/discover.go`, `report/playwright.go` + twin lockstep tests | **PAIRED** (grammar/lockstep was undocumented → fixed) |
| Runner concurrency / serial default | `playthroughs.md`, `spec.md` §5.7 | `e2e/playwright.config.ts`, `lib/stack-env.ts` | **PAIRED** (rationale was STALE → fixed) |
| Clerkenstein single-global-seat | `cockpit-spec.md` only; `clerkenstein.md` + `playthroughs.md` absent | `clerk-frontend/registry.go`, `server.go` | **CODE-ONLY → now PAIRED** (fixed) |
| Hero-login handshake | `cockpit-spec.md`, `playthroughs.md` | `e2e/lib/hero-login.ts` → `stack-verify/e2e/lib/cockpit-login.ts` | **PAIRED** |
| `pt-world` seed shape | `playthroughs.md` | `seed/pt-world.seed.yaml`, `seed-worlds.yaml`, `manifest/seed_worlds.go` | **PAIRED** (3 STALE claims → fixed) |
| Reset-to-seed lifecycle | `playthroughs.md`, `seeding-spec.md` | `stack-seeding/cmd/stackseed/main.go` §`doReset`, `run-playthroughs.sh` | **PAIRED** (chain STALE → fixed) |
| Seeded hero fan-out (7 tables) | `stories-spec.md` | `stack-seeding/seeders/persona*.go` | **PAIRED** (fully aligned) |
| Presence sweep (reused foundation) | `coverage-protocol.md` | `stack-verify/e2e/` | **PAIRED** (5 STALE claims → fixed) |
| `networkidle` doctrine + fences | `coverage-protocol.md`, `latency-budget.md` | `e2e/lib/page-object.ts`, 2 unit fences | **PAIRED** (8 code violations unfenced → reported) |
| Studio LLM-bound lane | `playthroughs.md` | `e2e/tests/studio-builder.spec.ts` | **PAIRED** |
| **Negative-control contract** | — | — | **BLIND-AREA** (annotated by overview's `Delivers →`) |
| **Per-spec mutation classification** | — | prose comments only, no artifact | **BLIND-AREA** (an overview deliverable) |
| **Onboarding stream (5 UCs)** | — (0 mentions in either doc) | — (no seedable state) | **BLIND-AREA** → Open Item 2 |
| **Org-admin stream (4 UCs)** | — (vantage label only) | read-only crawl descriptors only | **BLIND-AREA** → Open Item 3 |

---

## Fidelity Findings

### The three claims the milestone plan explicitly asked about

**F1 — `playthroughs.md` "18 live Playthroughs, 0 TODO" — ALIGNED.** ✅
18 `playthrough:` keys across the 8 `manifest/*.yaml` files (ai-readiness 4 · profile 4 · workforce 3 ·
assignment-monitoring 2 · studio-builders 2 · ai-simulations 1 · hiring 1 · skill-paths 1) = 18 `test()` blocks
across 17 spec files (`studio-builder.spec.ts` carries 2) = 18 unique `@pt:` ids, 1:1, and **zero** `TODO`
sentinels anywhere in the manifests. The number is exact. *(Fixed adjacent: the doc's own **lead paragraph**
still said "10 live Playthroughs" — the M204-era description — so the count was right at §108 and wrong at §14.)*

**F2 — serial-default config — PARTLY STALE, fixed.**
`fullyParallel: false` ✅ and `retries: 0` ✅ are literal. But the doc claimed literal **`workers: 1`**; the
config actually reads `workers: resolveWorkers()` (`lib/stack-env.ts`), which returns 1 by default and
**fails loud** on a `PW_WORKERS` that is not a positive integer. Semantically the same default, but the
fail-loud guard — a real safety property M256 will lean on when it experiments with `PW_WORKERS` — was
undocumented. The two sanctioned reclaim paths (stack-per-worker · per-worker seed partitions) are stated
correctly in both docs.

**F3 — Clerkenstein single-global-seat — the blind area was REAL, and worse: the doc stated a REFUTED rationale.**
🔴 *(highest-severity finding; fixed)*
- Code truth (all confirmed): `registry.go` holds **one** `activeKey` (`Registry.active()` / `Select()`);
  `server.go`'s `type Server` holds **one** `signedIn`, one `clientID` (`"client_clerkenstein"`, constant) and
  one `sessID` (`"sess_clerkenstein"`, constant). `handleSelectIdentity` re-points the seat **and** sets
  `signedIn = false; sessID = ""` globally. `handleMe`, `handleToken`, `handleClient` and
  `handleMeOrganizationMemberships` all **discard the `*http.Request`** and answer from `activeUserLocked()`;
  `r.Cookie(...)` appears **nowhere** in `clerkenstein/` — so **`storageState` reuse cannot isolate two seats**.
  `handleSignOut` ignores its `{id}` param and logs the whole stack out.
- Doc coverage before this audit: **stated in exactly one place** — `cockpit-spec.md` §*Limitation — one seat per
  stack* (which arrived as a rider on the M255-era deeplink feature). **Absent** from `corpus/services/clerkenstein.md`
  (the canonical service doc — its §Multi-identity sells "server-authoritative" as a *coherence feature*, never
  as a single-tenancy limit) and **absent** from `playthroughs.md` + `spec.md`, which justify the serial default
  **solely by Postgres** — the exact rationale the M256 plan review already refuted. A reader of the corpus would
  conclude per-worker seed partitions suffice. They do not.
- Two in-repo comments already recorded the verdict and neither had reached the corpus:
  `stack-verify/e2e/tests/m224-candidate-heroes.spec.ts` §serial preamble ("*a later `selectSeat` clobbers an
  earlier session*") and `stack-verify/e2e/tests/content-stories.spec.ts` §"SERIAL BY NECESSITY".

### Stale claims — `pt-world` / seeding (all load-bearing for clauses 2 + 3)

**F4 — `actor.entitlement` is DECLARED-ONLY; no seeder materializes a tier.** 🔴 *(fixed in doc; scope decision open)*
`grep -i entitlement` over all of `stack-seeding/` returns **zero** hits. `blueprint.TierMix` is parsed,
defaulted (`blueprint/stories.go` §`DefaultStoryTierMix` = free 0.7 / premium 0.3) and range-validated, then
**consumed by no seeder** — it never reaches a DB column; `pt-world.seed.yaml` declares no `tier_mix` at all.
Yet `playthroughs.md` said the world *"span[s] entitlement tiers"*, `seed-worlds.yaml` declares
`tiers: [anon, free, paying, enterprise, expired]` + the capability `entitlement-gated`, and annotates the
`pt-free` seat *"entitlement-gate use cases — **outcome: blocked**"*. `pt-free` is seeded as a user but is **not**
tier-gated and is referenced by **0 of 18** use cases. Worse, `ptvalidate`'s precondition-coverage check —
the one check whose job is to forbid a silent "ideally" — resolves `entitlement-gated` against the *declared*
list, so it **passes without the gate existing** (fail-open).

**F5 — There is NO pre-onboarding user state.** 🔴 *(answers overview Open Question 1: **no**)*
No `onboarding` / `completed_onboarding` / `is_onboarded` field anywhere in `stack-seeding/`. `UsersSeeder`
writes a `public.memberships` row for **every** seeded user unconditionally; there is no user with no
membership and no completeness gate. Every `pt-world` actor is a *post*-onboarding org member. Clause 3's 5
onboarding UCs therefore need a **net-new seed capability** (seeder + `capabilities:` entry + roster seat), not
just new specs.

**F6 — `--reset` is whole-stack, not org-scoped — and a code comment says the opposite.** 🟠 *(doc fixed; code comment left for M256)*
`doReset` (`stack-seeding/cmd/stackseed/main.go`) takes **no org filter**: it `TRUNCATE … CASCADE`s each of the
~28 `resetTables` — `public.organizations` and `public.users` included — for that stack, guarded only by
`--stack` + the N=0 `--force` rule (it does probe `to_regclass` and skip absent relations). But
`playthroughs/seed/pt-world.seed.yaml`'s header comment claims *"It resets the seeded orgs — NOT the whole DB;
the demo's showcase orgs … not touched by pt-world's reset."* **That is false.** `playthroughs.md`'s own
lifecycle line ("full FK-ordered TRUNCATE, per-stack only") was the accurate one. Operational consequence for
M256: a `--reset` suite run on a shared local demo **destroys the showcase world**.

**F7 — `seeding-spec.md`'s `--reset` chain enumerates 14 relations; `resetTables` is ~28.** 🟡 *(fixed)*
Missing from the doc's authoritative-reading chain: the 8 `ai_readiness_*` tables,
`interview_extraction_results`, `interview_aggregated_reports`, `organization_assignment_sessions`,
`local_skill_path_sessions`, `job_simulation_feedbacks`, `membership_tags`/`tags`,
`organization_target_roles`/`user_target_roles`, `membership_skills`,
`organization_sim_invitation_links`, `jobsimulation.{actors,interactions}`.

**F8 — `seeding-spec.md` mentions `pt-world` / Playthroughs ZERO times.** 🟡 *(fixed with a pointer)*
A declared KB dependency that says nothing about the world it is depended on for.

### Stale claims — `coverage-protocol.md` (the reused sweep foundation)

**F9 — the denominator is stale by two revisions.** 🟠 *(fixed)* Doc said `expected_pairs=49` (and "47/47" as
the live number); `stack-verify/e2e/content-denominator.json` pins **45** (24 player + 21 manager) since the
M254 close moved the 2 Bunny-absent voice cells to **manager**-presence-only as well. Same stale 49 repeated at
the `content-route-contract` paragraph.

**F10 — an exactly-inverted claim.** 🟠 *(fixed)* Doc: *"It now `exec`s the aggregator."* Code
(`run-content-stories.sh`): *"deliberately **NOT** `exec` — exec replaces this shell and the EXIT trap that
removes `$MANIFEST` would never fire. Run it as a child and forward its status verbatim."*

**F11 — the settle ceiling contradicts itself inside one doc.** 🟡 *(fixed)* §"Never wait on networkidle" said a
**~1.5 s** ceiling; `SETTLE_CEILING_MS = 4_000` (`lib/section-assert.ts`), and the doc's own iter-09 note 90
lines later says 4 s and explains why 1.5 s collapsed the BFS frontier.

**F12 — a render shape's name, route and assertion are all superseded.** 🟠 *(fixed)* Doc's `manager-dashboard`
at `/enterprise/activity-dashboard/ai-simulations/<simId>/<membershipId>`, gated on **attempts-table rows**.
Code: shape `manager-scored`, route `/sim/<slug>/<userId>/result/<sessionId>` (`isManagerView=true`) since
**M248**, gated on a persisted **`N/100` score** on a `readable ≥ 400` page — deliberately language-agnostic and
collapse-proof, because the old anchors are absent from a fully-rendered Italian result. (The 6 render shapes
themselves: ALIGNED.)

**F13 — a cited "reusable diagnostic" was never committed.** 🟡 *(fixed)*
`stack-verify/e2e/tests/probe-aireadiness-deeplink.spec.ts` does not exist.

### Milestone-plan anchor drift (`overview.md` itself)

**F14 — three of the overview's own `file:line` anchors have drifted** (all still resolve to the right
*function*, so the claims stand; the numbers do not):
- `server.go:573-586` → `handleSelectIdentity` now starts at **:563**, the global clobber is at **:582-583**.
- `server.go:454-466` → `handleMe` is **:454-464**.
- `registry.go:67,75` → correct for `active()` / `Select()`, but the `activeKey` *field* is at **:23**.
- ✅ Verified exact and unchanged: `skillpath-legacy.spec.ts:21-23`, `skill-path-page.ts:47-48`,
  `studio-builder.spec.ts:45` (`300_000`), `:70-81`, `:91` (`180_000`), `aisim-chat-launch.spec.ts:61`,
  `assignment-assign.spec.ts:43` (`240_000`), `hero-login.ts:44-53`, `m224-candidate-heroes.spec.ts:10-15`,
  `content-stories.spec.ts:128-130`, `spec.md:447-450`, `playthroughs.md:441-443`.
  **Recommendation:** prefer symbol anchors (`server.go::handleSelectIdentity`) — they don't drift.

**F15 — the overview contradicts itself on the mutation count.** §"(a) The count is wrong" says *"Only **10 of
18** carry an explicit '(no mutation)' declaration"* — verified true (9 literal `(no mutation)` comments +
`profile-growth.spec.ts`'s variant) — but its own "Correct statement" two lines later reads *"**17 of 18** are
UNCLASSIFIED for mutation."* The precise state is: **10** declare no-mutation in prose, **2** declare mutation
(`skillpath-legacy` MUTATES, `assignment-assign` WRITE), **6** declare nothing
(`profile-identity`, `profile-timeline`, `profile-verified`, `aisim-chat-launch`, `studio-builder` ×2) — and
**none** of the 18 is classified by a *machine-checked* artifact, which is the point that survives.

### Line-anchor lint (claim type 8)

The six docs in scope cite code **symbolically** (`` `file.ts` §symbol ``) almost without exception — exactly
**one** numeric code anchor exists across all of them (`coverage-protocol.md` → `queries.resolvers.go:70`), and
it is **still correct** (line 70 is `JobSimulationResult`, whose body does plain Ent SELECTs — the M231
persisted-read verdict holds). One path-prefix ambiguity: a `lib/target.sh` citation resolves to
`stack-verify/lib/target.sh`, not `stack-verify/e2e/lib/`.

### Frontmatter freshness (claim type 7)

**None of the six docs has YAML frontmatter** — all six open with an H1, no `title:`/`description:`/
`last_updated:`. There is therefore no machine-readable freshness signal anywhere in this cluster. Graded
against the de-facto lead paragraph, **five of six leads are stale** (they describe the doc as of the milestone
that created it): `playthroughs.md` said "10 live Playthroughs" (**fixed**); `coverage-protocol.md` stops at
M236; `cockpit-spec.md` omits its own Deeplinks / Back-to-Cockpit / content-stories-tab sections;
`stories-spec.md` omits 3 of its 4 stories; `seeding-spec.md` under-claims its own contents.
`spec.md` is internally fresh **as a frozen v0.3 draft** — which is the problem: M256 declares it a KB
dependency and it predates M219/M225/M243/M252.

---

## Completeness Gaps

1. **Critical — the seat-contention property** (F3). Now documented in `clerkenstein.md` + `playthroughs.md`.
2. **Critical — no negative-control contract exists anywhere.** `grep -i "negative control"` over
   `playthroughs.md` + `coverage-protocol.md`: 0 hits. The only prior art in the corpus is
   `LEGACY_AI_READINESS_URL` (one anti-assertion) and `pt-assignment-assign`'s read-back protocol
   ("*the read-back IS the proof — the anti-toothless bar*"). **M256 authors this contract** — it is in the
   overview's `Delivers →` line.
3. **Critical — no per-spec mutation classification artifact.** Prose comments only (F15). **Useful prior art
   the milestone should extend rather than reinvent:** the `@pt:` tag grammar `@pt:([a-z0-9][a-z0-9._-]*)` is
   duplicated across `cmd/ptvalidate/discover.go` and `report/playwright.go` **by necessity** (the Go packages
   don't import each other) and both copies are pinned by **twin lockstep tests** (`pttag_lockstep_test.go` in
   each package) to one canonical literal + a shared match corpus. A new `MUTATES`/`READ-ONLY`/`UNKNOWN` tag
   must follow that shape, not add a third unfenced regex. *(Now documented in `playthroughs.md`.)*
4. **Critical — 8 unfenced `networkidle` violations in `playthroughs/e2e/`**, each a direct cost against gate
   clause 1 (the suite-speed clause) and none covered by either existing fence:
   - `goto` gated on `networkidle` (overriding the compliant base `page-object.ts` `goto`):
     `lib/simulation-page.ts:36`, `lib/skill-path-page.ts:31`
   - unbounded `waitForLoadState('networkidle')` (no timeout — `.catch()` only defers the failure to the
     120 s test budget): `lib/activity-dashboard-page.ts:74`, `lib/assignments-page.ts:63`,
     `lib/profile-page.ts:63`, `lib/profile-page.ts:123`, `tests/skillpath-legacy.spec.ts:73`,
     `tests/aisim-chat-launch.spec.ts:62`
   - the shared `cockpit-login.ts`'s `waitUntil` **default is `'networkidle'`** (preserved for the coverage
     sweep) and `hero-login.ts` forwards it, so any Playthrough omitting `waitUntil` inherits it.
   - Existing fences cover only `/home`-landing `loginAsHero` calls
     (`tests/home-login-networkidle.unit.spec.ts`) and 3 named surfaces
     (`tests/page-object.unit.spec.ts`) — neither `goto` override, and no unbounded settle, is caught.
5. **Incidental — `stories-spec.md`'s seeder inventory is ~12 seeders behind** the fleet
   (`TaxonomySnapshotSeeder`, `ContentSnapshotSeeder`, `ContentStorySeeder`, `ContentStoryNonSimSeeder`,
   `AIReadinessSimSkillsSeeder`, `GeneratedBatchSeeder`, `MemberLanguagesSeeder`, `CertificatesSeeder`,
   `ProjectsSeeder` are absent by both type name and filename), and its `AIReadinessFunnelSeeder` write list
   omits `jobsimulation.interview_aggregated_reports`. **The doc's core claim is fully ALIGNED**, though: all
   **7** tables of the fan-out are still written by `PersonaSeeder` (`seeders/persona_write.go` flush), every
   named seeder exists, and `datadna measure-closure` exists. Not load-bearing for M256 → left as-is.
6. **Incidental — `coverage-protocol.md` never states Sweep 1's own `workers`/`fullyParallel` posture.** The
   parallel-vs-serial contrast M256's speed clause reasons about exists **only** as a code comment in
   `playthroughs/e2e/playwright.config.ts`. (Verified: `stack-verify/e2e/playwright.config.ts` is
   `fullyParallel: true`, `workers: CI ? 2 : undefined`.)
7. **Code defect found, out of doc scope — reported for routing.** `run-content-stories.sh` re-implements
   `buildPairs()` in inline Python and checks only `player_presence_only`, **not** `manager_presence_only`
   (which `lib/content-pairs.ts` added at the M254 close, and which is `true` for 2 sessions in the canonical
   preset). Its cross-check therefore computes **47** against the external pin of **45** and `sys.exit(2)` —
   **the content-stories sweep would refuse to start.** This is verbatim the "fourth inline copy" hazard
   `content-denominator.json`'s own header warns about. Not M256's suite, but it is in a section M256's docs
   cover; route it to whoever next runs that sweep (M257/M258 compose it).

---

## Applied Fixes

All in the `rosetta` docs repo, on branch `m256/playthrough-sharpening`, **uncommitted** (the parent agent owns
commits). No file in `.agentspace/rosetta-extensions` was modified.

| File | Fix |
|---|---|
| `corpus/ops/demo/playthroughs.md` | § lifecycle: corrected the serial-default claim (`resolveWorkers()` + fail-loud) and added the **⚠️ "Postgres is NOT the binding shared surface"** box — the full fake-FAPI seat mechanism, the `storageState`-doesn't-isolate consequence, why per-worker seed partitions are insufficient, and the supersession of `spec.md` §5.7 (F2, F3) |
| ″ | § the model: `actor.entitlement` marked **declared-only, never materialized**; `outcome` marked **`blocked`/`error` coverage = 0** (F4) |
| ″ | § the Playthrough world: 3-point correction box — the entitlement declaration, the absent pre-onboarding state, and the whole-stack `--reset` (F4, F5, F6) |
| ″ | § both-way id integrity: documented the `@pt:` grammar + the twin lockstep-test pattern (gap 3) |
| ″ | lead paragraph: 10 → **18 live Playthroughs, 0 TODO** + the frozen-`spec.md` precedence rule (F1, frontmatter) |
| `corpus/services/clerkenstein.md` | § Multi-identity: added the **"server-authoritative means SINGLE-TENANT"** box — one seat / `signedIn` / `clientID` / `sessID` per stack, the destructive select, the no-request-input read path, and that per-client keying is an auth-model + alignment-DNA change (F3) |
| `corpus/ops/demo/coverage-protocol.md` | denominator 49 → **45** with the 49→47→45 history and "read the pin file, never a number in this doc" (F9); the `exec` claim inverted to match the code + its EXIT-trap reason (F10); settle ceiling 1.5 s → **4 s** (F11); `manager-dashboard` → **`manager-scored`** with the M248 route + score-based gate (F12); the non-existent `probe-aireadiness-deeplink.spec.ts` corrected (F13) |
| `corpus/ops/seeding-spec.md` | § Re-run safe: the 14-relation chain marked **illustrative** with the ~28-table truth + the no-org-filter / `to_regclass`-skip properties, and a pointer to the `pt-world` lifecycle (F7, F8) |
| `.../m256-.../spec-notes.md` | the topic → doc → code triple table + the 3 open KB prerequisites |

---

## Open Items (require user decision)

1. **🔴 Exit-gate clause 2's `blocked` outcome has no seedable mechanism (F4) — and iter-01 **D4** already
   decided the opposite.** `iter-01/decisions.md` § D4 ("The `blocked` outcome needs no seed work") is refuted
   above and must be re-decided; the strategy in `playthrough-map.md` §8.6 step 5 inherits it. The `entitlement`
   axis is a label. Pick one: **(a)** make `TierMix` real (a `stack-seeding` change +
   a `pt-world` `tier_mix` + a tier column consumer) and gate `pt-free`; **(b)** source the `blocked` from a
   different refusal surface that *is* real today — cross-org access to a private path, an RBAC/Casbin deny, a
   validation error; or **(c)** re-cut the clause. Also worth deciding: whether to close the `ptvalidate`
   fail-open that lets `entitlement-gated` resolve against a declaration.
2. **🔴 Onboarding (5 UCs) is a SEEDER change, not a test change (F5).** Confirm it stays in clause 3's scope
   at that cost, or route the onboarding stream forward and let clause 3 stand on org-admin + written verdicts.
3. **🟠 Extend the overview's `Delivers →` line** to name the **onboarding** and **org-admin** knowledge, so
   those two blind areas are explicitly milestone deliverables rather than implicitly covered.
4. **🟠 Two code-side items to route** (both outside this audit's doc remit): the false `--reset` comment in
   `playthroughs/seed/pt-world.seed.yaml` (F6), and the `manager_presence_only` omission in
   `run-content-stories.sh` that makes the content-stories sweep `exit 2` before running (gap 7).
5. **🟡 The 8 unfenced `networkidle` violations** (gap 4) are the clearest clause-1 speed lever visible without
   a live run — and the natural place for a new fence. Confirm they are in-scope for iter-01's triage.
6. **🟡 `spec.md` v0.3 is a frozen draft superseded on §5.7.** This audit added the precedence rule to
   `playthroughs.md` rather than edit the pinned draft. Confirm that's the right call, or authorize a
   `Superseded-by` banner in `spec.md` itself.

---

## Gate Result

**YELLOW — proceed with tracking.** `/developer-kit:build-mstone-iters` may enter iter-01 (the bootstrap tok).
Every stale load-bearing claim is fixed; the remaining items are scope decisions, recorded above and in
`spec-notes.md` § KB prerequisites. **The bootstrap tok must resolve Open Item 1 before its strategy commits to
gate clause 2, and Open Item 2 before it prices clause 3.**

`SEVERITY: warning`
