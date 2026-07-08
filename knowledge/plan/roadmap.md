# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills). This file holds the **active major** only; the retired **v1.x** history (M0 … M46, all
SHIPPED) lives in [`roadmap-legacy.md`](roadmap-legacy.md). Future versions + the unscheduled backlog live in
[`roadmap-vision.md`](roadmap-vision.md). The live source of truth for *current/next* is [`state.md`](state.md).

> **Designed 2026-07-08** via `/developer-kit:design-roadmap`. **v2.1 "quick change"** is the **skiller-in-app
> re-ground** — a field-hardening release (v1.3b "dress rehearsal" / v1.10b "fit-up" lineage) triggered by a
> **landed platform structural change**: the `skiller` service + its DB schema **merged into `app`** (domain → the
> **`public`** schema, table names unchanged `skiller.X → public.X`; RPC → `backend`; the skiller GraphQL subgraph
> gone → **4 subgraphs**; skiller repo/container removed). A colleague's `docs/skiller-in-app-merge` corpus sweep is
> **correct-but-incomplete** and touches no tooling; the **rext** tooling still queries `skiller.<table>` and the
> stacks are half-synced. v2.1 re-fits the tooling, the corpus, and the stacks to the merged platform and **proves
> `dev-up` + `demo-up` still work**. **4 milestones M208 → M209 → M210 → M211** (strictly sequential); tag
> **`v2.1`**; branch `release/02.10-quick-change`. Tooling + docs only — zero platform-repo edits.
>
> **Designed 2026-06-29** via `/developer-kit:design-roadmap`. **v1.10b "fit-up"** is an **interposed
> field-hardening backfill** (the v1.3b "dress rehearsal" lineage): a from-scratch `/demo-up` surfaced 8 bring-up
> issues + a tail of v1.10 content gaps. The M201 close *reported* the `stack-demo` clones ~5 weeks / 115+ commits
> behind prod — but **M47 found the clones actually CURRENT** (next-web @ v2.89.0); the genuinely-stale surface is
> the **corpus** (the shipped AI-readiness feature is undocumented), which **M48** re-grounds. v1.10b recaptures the
> snapshot, re-grounds the corpus, fixes the bring-up + content issues, adds an AI-readiness showcase org, and
> consolidates one auditable manifest — so **v2.0 "opening night" is PAUSED** until it ships. The v1.x flat
> milestone counter **re-opens** at **M47**
> (M47→M53); tag **`v1.10.1`**; branch `release/01.10b-fit-up`.
>
> **Designed 2026-06-28** (prior): **v2.0 "opening night"** opened a **NEW MAJOR** — **Playthroughs** is a new pillar
> (functional-flow *testing*: proving the platform's core user journeys actually work end-to-end), distinct from the
> v1.x demo/seeding lineage. v2+ adopts the **`Mxyy`** scheme (M201, M202, M203, M204). **v2.0 SHIPPED 2026-07-02
> (tag `v2.0`)** — all four milestones closed, 10 live Playthroughs GREEN on cold reset-to-seed; the first v2.x
> release. **v2.1 "quick change"** (the skiller-in-app re-ground) is now **IN DEVELOPMENT** (designed 2026-07-08;
> see below). The pre-v2 v1.x history (M0 … M46) lives in `roadmap-legacy.md`.

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v1.10b** | **fit-up** | Field-hardening backfill — re-ground demo + corpus to current prod, fix the from-scratch `/demo-up` issues + the v1.10 content gaps, add the **AI-readiness showcase org**, and consolidate **one auditable seed+gen manifest** | M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53 | ✅ **SHIPPED 2026-07-01 (tag `v1.10.1`)** (branch `release/01.10b-fit-up`, designed 2026-06-29; all 7 milestones done) |
| **v2.0** | **opening night** | The platform's core user journeys, **proven to actually work** — a new **Playthroughs** pillar: a manifest-driven, deterministic e2e suite that *pretends to be the human* and proves the platform does its job | M201 ✅ ∥ M202 ✅ → { M203 ✅ ∥ M204 ✅ } → ✅ ship | ✅ **SHIPPED 2026-07-02 (tag `v2.0`)** (branch `release/02.00-opening-night`, designed 2026-06-28; all 4 milestones closed-on-gate/complete). **10 live Playthroughs (6 employee + 4 manager) GREEN on cold reset-to-seed, 1 in-manifest TODO.** The first v2.x release. Tooling + docs only, zero platform edits, zero new deps |
| **v2.1** | **quick change** | The **skiller-in-app re-ground** — re-fit the tooling, corpus, and stacks to the merged platform (skiller service + schema folded into `app`/`public`, RPC → `backend`, **4 subgraphs**) and **prove** `dev-up` + `demo-up` still work. Field-hardening lineage (v1.3b/v1.10b), triggered by a landed platform change | M208 → M209 → M210 → M211 (strictly sequential) | 🔨 **IN DEVELOPMENT** (branch `release/02.10-quick-change`, designed 2026-07-08; tag `v2.1`; tooling + docs only, zero platform edits) |

> The complete v1.x version-plan table (v1.0 "body double" … v1.10 "method acting", all ✅ SHIPPED) is preserved
> in [`roadmap-legacy.md`](roadmap-legacy.md) § Version plan.

The Playthroughs capability is governed by the consolidated **capability spec**
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (v0.3 — all in-scope decisions made +
review-hardened). v2.0's milestones build the contract that spec defines. Hard constraints carry over from the
v1.x lineage: **no modification to any platform repo** (the platform stays read-only — a surface that can't be
driven without a platform edit *escalates*, it does not edit), and all stack-operating tooling lives in
**`rosetta-extensions`** (built + tested in the `.agentspace/rosetta-extensions/` authoring copy, tagged, then
consumed per-stack at a pinned tag). Playthroughs reuse the M42 e2e foundation + the seeding machinery — they are
the **functional** sibling of M42's **presence**-only coverage sweep.

---

## In Development — v2.1 "quick change" (designed 2026-07-08)

> **Status (IN DEVELOPMENT):** designed 2026-07-08 via `/developer-kit:design-roadmap`; branch
> `release/02.10-quick-change` cut from `main`; tag `v2.1`; rext tag `v2.1`. **4 milestones M208 → M209 → M210 →
> M211, strictly sequential** (the user's execution choice — single-substrate-safe, no parallel-branch
> coordination). All `planned`.
>
> **Theme:** *quick change — backstage, the actor sheds one costume and re-enters as another, seamless to the
> audience.* skiller's part now folds into `app`; v2.1 re-fits the whole apparatus to the changed stage. A
> **field-hardening re-ground** in the **v1.3b "dress rehearsal" / v1.10b "fit-up"** lineage — but triggered by a
> **landed platform structural change**, not a bring-up defect: the standalone `skiller` Go service **merged into
> the `app` monolith** — its domain (60K-skill taxonomy, embeddings, job-roles, matching) now lives in the
> **`public` schema** (table names unchanged, only the schema prefix changed `skiller.X → public.X`; the old
> `skiller` schema is dropped), its Connect-RPC surface is served by `backend`
> (`SKILLER_RPC_ADDR=http://backend:8083`), its GraphQL subgraph left the federation (**4 subgraphs** now), and
> the skiller repo + container are gone (`make init` no longer clones it). **This is landed upstream** (platform
> `origin/main` @ `0808b92` drops skiller from compose + repos.yml; app @ v1.334 carries the "Deprecate skiller
> schema" merge) — so v2.1 is **tooling + docs + stack-re-sync only; zero platform-repo edits** (the platform
> already did its half).
>
> **Why a release:** three surfaces are stale + mutually inconsistent. (1) **Stacks** — `app` is post-merge but
> `platform` is 2 commits behind (still composes skiller); both stacks hold vestigial `skiller/` clones. (2) **rext
> tooling** — untouched by the docs branch; still queries `skiller.<table>` (seeding taxonomy resolvers, snapshot
> capture/replay) and probes a skiller container/schema (`readiness.sh`, `services.sh`, `up-injected.sh`), so it
> **breaks** on the merged platform. (3) **Corpus** — the colleague's `docs/skiller-in-app-merge` sweep is
> **correct-but-incomplete**: the architecture/subgraph half is solid, but 5–6 rext-facing tooling docs still
> describe `skiller.*` and it **cannot land independently** without contradicting the tooling.
>
> **Designed 2026-07-08** via `/developer-kit:design-roadmap`, from the user's skiller-merge briefing + the
> colleague's `docs/skiller-in-app-merge` branch. A 7-agent research workflow (`wf_08b6bf4a`) mapped the per-module
> blast radius (file:line), **adversarially confirmed** the firewall public-predicate **survives** the merge
> (`organization_id IS NULL` still isolates public taxonomy — no data-leak risk; `skiller_mixins.OrganizationIDMixin`
> ports the tenant boundary 1:1), and confirmed the docs branch **cannot land present-tense** before the rext
> re-ground + stack re-sync. **Phase-0: 🟢 GREEN** deferral audit (M25-D9 opportunistic Fate-1 on the re-sync
> migration path) + clear KB blind-areas (every topic has an anchor).

### Execution graph

```
v2.1 "quick change"   (strictly sequential — single track)
  M208 ─────→ M209 ─────→ M210 ─────→ M211
  re-sync     rext        corpus      bring-up
  +ground     re-ground   re-ground   acceptance (iterative)
```

**Strictly sequential** (the user's choice). **M208** establishes current merged code — everything grades against
it. **M209** re-grounds the rext tooling (+ recapture the snapshot from merged-prod). **M210** completes the corpus
in lockstep with M209's landed schema (the tooling-doc bodies flip to `public.*`). **M211** is the iterative
acceptance closer — bring-up isn't trusted until `/dev-up` + `/demo-up` both go GREEN cold on the merged platform.
*(Scope-flex from design: if the corpus reconcile proves large, M210 can split into land+reconcile / rext-gap-fill;
4 is the clean target.)*

### Milestones

**M208 — Re-sync & merged-schema ground-truth** · `section` · complexity **medium** · depends on: **none** ·
✅ **`done` — closed 2026-07-08** (merged → `release/02.10-quick-change`). The **load-bearing foundation** of
v2.1 (mirrors v1.10b M47) landed: re-synced both stacks to the merged platform (`app` `a848cccb→c3c45e01`
v1.334.1 — the **86-commit merge pull**; `platform` `5e1ae6b→0808b92` — skiller gone from
compose/repos.yml/Make; sibling set current), removed both vestigial `stack-*/skiller/` clones, and **retired
the #1 release risk GREEN** via a live containerized de-risk on stack-dev: cold `make up` rc=0 with a
**4-subgraph** compose and **no skiller container** (`SKILLER_RPC_ADDR=http://backend:8083`); a clean-slate
`make reset-db`+`make migrate` builds the **full `public` taxonomy from scratch** (skills w/ `organization_id`,
job_roles, job_role_skills, skill_embeddings, categories, specializations) with **no `skiller` schema on a clean
DB**; measured prod `public.skills WHERE organization_id IS NULL` = **42,790** (the roadmap's ~42,763 figure).
Pinned the authoritative **merge fact-sheet** in `corpus/services/backend.md` (§ Skiller-in-app merge + banner)
+ a stub banner on `corpus/services/skiller.md`. Committed rosetta diff is **100% documentation** (zero
code/test) → HARDEN N/A; close review found **0 findings**; deferral audit **GREEN** (5 single, 0 repeat).
**Two bring-up findings routed forward (both user-accepted):** Finding 1 (clean-bring-up `extensions`-schema
bootstrap + PG-readiness, the M25-D9 class — did NOT fall out as a trivial Fate-1) → **Fate-3 M211** (+ M209
Risk-2 cross-ref); Finding 2 (`INVITATION_HMAC_SECRET` dev-`.env` completeness gap, not merge-caused) →
**Fate-2 M211 / `/stack-secrets`**. Zero platform-repo edits. **Delivers →** the merge fact-sheet
(`corpus/services/backend.md` + `corpus/services/skiller.md` stub).
**Goal:** bring both stacks (and the snapshot's target reality) current with the merged platform, and pin the
authoritative merge fact-sheet — so every downstream fix grades against current code.
**Scope:**
- **In:** `make pull` stack-dev + stack-demo `platform` to `origin/main` (skiller gone from compose/repos.yml),
  pull `app` to current (v1.334, post-merge domain) + the sibling repo set; capture before/after refs. Remove the
  vestigial `stack-dev/skiller/` + `stack-demo/skiller/` clones. Rebuild images + re-run migrations against the
  merged `public` schema; confirm the 4-subgraph compose (`backend`, jobsimulation, cms, skillpath), no skiller
  container, `SKILLER_RPC_ADDR=http://backend:8083`. Pin the **merge fact-sheet** (moved tables in `public`, the
  confirmed `organization_id IS NULL` public predicate, the ~42,763 public-skill count, the re-pointed RPC).
  Opportunistic **M25-D9** dev migrate-ordering fix (Fate-1 — lives on this path).
- **Out:** rext code (M209); corpus body re-point (M210); live bring-up (M211).
**Delivers → knowledge/corpus:** the merge fact-sheet (anchored in `corpus/services/backend.md` +
`corpus/services/skiller.md` stub).
**Open questions:** does the 86-commit `app` pull + migration re-run surface a schema/migration issue (⚠ the fit-up
M47 risk class)? — bounded; capture before/after.

**M209 — rext tooling re-ground** · `section` · complexity **medium** · depends on: **M208**.
**Status:** `done` (completed 2026-07-08).
**Goal:** re-point every rext tool that queries the old skiller schema or expects the skiller service to the merged
reality, recapture the snapshot, and tag a new rext.
**Scope:**
- **In (snapshot):** flip `stack-snapshot/taxonomy/taxonomy.go:43 const Schema "skiller"→"public"` (re-grounds
  capture *and* replay); update the 2 load-bearing `taxonomy_test.go` PublicVia assertions; **narrow the
  `pg.SchemaVersionSQL` staleness digest** to the surface's enumerated tables (fixes the post-merge whole-monolith
  cache-thrash — Risk 1); **verify the capture SELECT column list** vs merged prod (Risk 2 —
  `embedding→small_embedding3`, `extensions.`-qualified GIN opclasses — the one non-mechanical bit); keep
  `AssertPublicOnly` + add the ~42,763-row post-capture assertion; **recapture** the public taxonomy from
  merged-prod into `.agentspace/snapshots/` (bump the capture version; the batch cache re-keys).
- **In (seeding):** re-point the 5 real-SQL files (`seeders/taxonomyref.go`, `skillref_named.go` [the shared
  `namedSkillSelect` const → also fixes `curated_pools.go` + `ai_readiness_config.go`], `jobroleref.go`,
  `taxonomy_snapshot.go`, `dna/fidelity_probe.go`) `skiller.*→public.*` keeping the `organization_id IS NULL`
  public-pool predicate; drop skiller from `isolation/isolation.go` schema-note + re-ground `dna/data-dna.json`
  golden (schema + FK ref_schema); rename the 111 fake-Conn test string-matchers in lockstep; reword the
  comment/attribution refs (incl. `services/ai`).
- **In (small):** remove skiller from `stack-verify/lib/readiness.sh` `probe_postgres_schemas` + `services.sh`
  container probe; drop skiller from `demo-stack/up-injected.sh` INJECT_SVCS (else it clones/builds a gone repo) +
  the "5 Clerk-token services"→4 note.
- Build + test the authoring copy; **tag a new rext (`v2.1`)**; prepare the per-stack consumption re-pin.
- **Out:** the stack re-sync (M208); doc bodies (M210); live bring-up (M211).
**KB dependencies:** `corpus/ops/snapshot-spec.md`, `corpus/ops/seeding-spec.md`, `corpus/ops/safety.md`.
**Closed 2026-07-08:** re-grounded rext `skiller.*→public.*` across **stack-snapshot** (the `taxonomy.go` const
flip re-grounding capture+replay; the Risk-1 `Surface.VersionTables()` digest-narrowing so taxonomy digests only
its 10 tables while a structure-bearing surface still whole-schema-invalidates; a one-sided `MinRows` under-capture
floor; Risk-2 verified — the capture is names-only/type-agnostic so the `extensions.`-qualified merged columns need
no change), **stack-seeding** (24 files, `organization_id IS NULL` preserved, `data-dna.json` golden, `isolation.go`),
and the **small shell modules** (readiness/services/up-injected/migrate-demo, 5→4 services; stack-verify Python
104/104). **6 Go modules GREEN, `go vet` clean, 5× flake-clean; 0 `skiller.<table>` queries in any production path.**
rext tagged `quick-change-m209` (build `00a3ec5`/`e458acf`/`75bc4cf` + harden `42ad600`/`72a5259`/`2f06e78` — 14
harden test funcs, 0 bugs, 0 flakes; tag re-pointed to the post-harden HEAD `2f06e78` at close). Close: **1
nice-to-have finding** (a pre-existing rext `stack-seeding/README` test-count drift, last reconciled M41 — routed
to the v2.1 rext roll; rext frozen at `2f06e78` this close), **deferral audit GREEN** (10 in scope, 0 repeat/aged/
escape). **Recapture Fate-3→M211** (tooling READY; no local COPY-byte capture source — a data op M211 owns).
The `v2.1` rext roll + `.agentspace/rext.tag` consumption re-pin remain **`/developer-kit:close-release`'s** job.

**M210 — Corpus + skills re-ground** · `section` · complexity **medium** · depends on: **M209** (the tooling-doc
bodies flip to match M209's landed schema).
**Status:** `done` (completed 2026-07-08).
**Goal:** land the colleague's docs sweep as the *complete, internally-consistent* corpus re-ground.
**Scope:**
- **In:** adopt/validate `origin/docs/skiller-in-app-merge` (correct-but-incomplete — the architecture/subgraph
  half is solid); fix the fully-missed `profile-completeness-spec.md` (43/44→44/44); **flip the 5–6 rext-facing
  tooling-doc bodies to `public.*`** (`snapshot-spec.md` [26 mentions — the taxonomy surface enumeration +
  FidelityProbe gene + capture predicate], `safety.md`, `recipe-snapshot-world.md`, `stories-spec.md`,
  `seeding-spec.md`, `coverage-protocol.md`) + delete the interim disclosure notes; reconcile the db-access ↔
  tooling contradiction; sweep the skill files (`dev-up/reference`, `stack-snapshot/SKILL`, `stack-update/reference`,
  `db-query/SKILL`) so container counts / migration lists / RPC addr / subgraph counts match the re-synced stacks;
  update `CLAUDE.md` service catalog.
- **Out:** rext code (M209); live bring-up (M211).
**Closed 2026-07-08** (merged → `release/02.10-quick-change`): made the corpus internally consistent with the merged
platform + M209's landed `public.*` rext code. Adopted the colleague's correct architecture/subgraph/service half
(28 files; each hunk verified vs the M208 fact-sheet + the re-synced `stack-dev/app` clone; kept M208's authoritative
`backend.md`/`skiller.md` fact-sheet — no duplicate merge section); fixed the profile-completeness node-id prose
(**verified NO literal "43/44" exists anywhere** — the design-note count was inaccurate; made the one genuine
merge-sweep fix, did **not** fabricate a phantom count); flipped the **6 rext-facing tooling-doc bodies + directus-local.md**
`skiller.*→public.*` and deleted the interim disclosure notes; reconciled db-access ↔ tooling on `public.*`; swept the
4 skill files to the verified merged compose (no skiller container, **4 subgraphs**, `SKILLER_RPC_ADDR=http://backend:8083`;
superseded the colleague's now-stale "still-targets-skiller/exit-4" note with an accurate M209-done note); updated the
`CLAUDE.md` catalog. Email-asset PNGs excluded. **Verified core outcome:** 0 stale `skiller.<table>` tooling-query refs
corpus-wide, 0 leftover interim notes, 4 subgraphs consistent, 0 broken `.md` links. **RESOLVES the KB-1/2/3 body-flip
deferrals** M208+M209 routed Fate-2 → M210 (7 defers landed at destination). Committed rosetta diff is **100%
documentation** (50 `.md`, 0 code/test) → HARDEN N/A; close review **0 must-fix / 1 nice-to-have no-change-needed**
(the app==backend subgraph dual-naming is the corpus's established convention, not an M210 defect); deferral audit
**GREEN** (11 in scope, 7 resolved, 4 still-open confirm-only → M211/close-release; 0 repeat/aged/escape). Zero
platform-repo edits.
**Delivers → knowledge/corpus:** the completed corpus (`corpus/services/skiller.md` stub + `backend.md` ownership
+ re-pointed tooling docs).

**M211 — Bring-up acceptance: dev-up + demo-up green on the merged platform** · `iterative` (closed-on-gate) ·
complexity **large** · depends on: **M209 + M210**.
**Status:** `planned`.
**Goal:** prove the whole chain works end-to-end on the merged platform with the re-grounded tooling.
**Exit gate:** from a re-synced state, **`/dev-up` AND `/demo-up` both go GREEN cold** — 4-subgraph compose / no
skiller container; snapshot **recapture→replay** loads `public.*` (taxonomy replay exits 0, ~42,763 public skills);
**seed** resolves real public node-ids (closure green); **verify** (`verification.md` net) passes with a
merged-platform assertion (no skiller schema/subgraph/container; `readiness.sh` schema probe green); the M42
coverage sweep + the v2.0 Playthroughs suite stay GREEN; **0 residual skiller-schema references** in any queried
path.
**Iteration protocol:** the fit-up/dress-rehearsal fix→re-measure→re-run bring-up loop
(`corpus/ops/verification.md` + the coverage/playthroughs gates).
**Why iterative:** the merged 4-subgraph platform has never been stood up locally with the re-grounded tooling; it
*will* surface fix-loops (migration ordering, the column-mapping caveat, vestigial container/clone cleanup, cache
behavior).

### Top risks

1. **Cache-key digest regression** (degrades-quality) — post-merge `SchemaVersionSQL` digests the whole app
   monolith → taxonomy cache thrashes on any migration. **Mitigation:** narrow the digest to enumerated surface
   tables (M209). *The single non-obvious regression the merge introduces.*
2. **Capture column-mapping** (blocks-release if wrong) — the SELECT list may not be a pure prefix swap
   (`small_embedding3`, `extensions.`-qualified opclasses). **Mitigation:** verify vs merged-prod in M209; the
   ~42,763-row assertion catches empty/over-broad capture.
3. **86-commit app pull + migration re-run** (M208) — the fit-up M47 ⚠ class. **Mitigation:** bounded; capture
   before/after refs.
4. **Docs lockstep** (would create a self-contradicting corpus) — the branch can't land present-tense before rext +
   re-sync. **Mitigation:** M210 flips bodies in lockstep with M209 (why the design keeps them adjacent + sequential).
5. **Recapture safety** — verified **low**: the firewall predicate HOLDS (`org_id IS NULL` survives the merge) +
   `AssertPublicOnly` runtime net rejects any non-null-org captured row.

---

## Done — v1.10b "fit-up" (SHIPPED 2026-07-01, tag v1.10.1)

> **Theme:** *fit-up — build and rig the set correctly in the venue before opening night.* An **interposed
> field-hardening backfill** in the **v1.3b "dress rehearsal"** lineage. A from-scratch `/demo-up` surfaced 8
> bring-up issues + a tail of v1.10 content gaps. **CORRECTION (M47 finding):** the M201 close *reported* the
> clones ~5 weeks / 115+ commits behind prod, but **M47 found them CURRENT** (next-web @ v2.89.0, every repo ≤2
> behind; the AI-readiness feature present) — the re-sync was a trivial `make pull`. The real stale surface is the
> **corpus** (M48). v1.10b **recaptures** the snapshot from current prod, **re-grounds** the corpus, **fixes** the
> bring-up + content issues, **adds** a curated **AI-readiness showcase org** (redeeming the M201
> member-AI-readiness false-negative), and **consolidates** one auditable **seed+generation manifest**.
> **Tooling + docs only — zero platform-repo edits.** The v1.x flat counter re-opens at **M47** (M47→M53); tag
> **`v1.10.1`**; branch `release/01.10b-fit-up`.
>
> **Designed 2026-06-29** via `/developer-kit:design-roadmap`, from the field review in
> [`.agentspace/annotation.md`](../../.agentspace/annotation.md) + the M201 stale-clone finding. Three research
> agents mapped the fix surfaces (file:line), the content/seeding gaps, and the KB blind-areas (all homed below via
> `Delivers →` lines).

### Execution graph

```
v1.10b "fit-up"   (ONE live demo → verification serializes; only rext authoring parallelizes)
  M47 ──→ ┌ M48  corpus re-ground   (reads code, NO demo) ─┐
          └ M49  bring-up hardening (NEEDS the live demo)  ─┘ ──→ M50 ──→ M51 ──→ M52 ──→ M53
 re-sync                                                       content  AI-ready  manifest  cold-rebuild
```

**The 1-demo-stack constraint shapes the graph.** The user's box hosts **one** demo at a time, so milestone
**verification serializes** on the single live stack — the release is an honest sequential chain ending in a
**cold destroy-and-rebuild acceptance** (M53). The **one** genuine parallel pair is **M48 ∥ M49**: M48 (corpus
re-ground) is pure docs-vs-code and never touches the demo, while M49 (bring-up hardening) monopolizes it —
disjoint file clusters (`architecture`+`services` vs `ops`+rext), additive merge. The "fix-on-live → final cold
rebuild" the user asked for *is* the M47→…→M53 shape.

### Milestones

(compact per-milestone summaries — the full contracts live under
[`releases/archive/01.10b-fit-up/m*/overview.md`](releases/archive/01.10b-fit-up/))

**M47 — Re-sync & recapture** · `section` · depends: **none** (foundation) · ✅ **`done` — closed 2026-06-29**
(merged → `release/01.10b-fit-up`; rext tag `fit-up-m47`). The flagged heavy re-sync was a **no-op** — the clones
were already current (next-web @ v2.89.0); the genuine staleness is the corpus (→ M48). Delivered: `pg.NormalizeDSN`
(sslmode `no-verify→require`) so the wired MCP DSN works as a capture `--dsn` (demo-up #2, proven by a live dry-run +
9 tests); all 3 snapshot surfaces recaptured from current prod (digests unchanged); the **AI-readiness feature
confirmed present** (M201 false-negative resolved); `snapshot-cold-start.md` updated (the MCP-configured-DSN path,
KB-47-01). The `up-injected.sh` auto-capture wiring was descoped per the user (no new entry point — D3); re-pin
deferred (push-gated). **Delivered → `corpus/ops/snapshot-cold-start.md`.**

**M48 — Corpus re-ground** · `section` · depends: **M47** · ∥ **M49** · ✅ **`done` — closed 2026-06-29** (merged →
`release/01.10b-fit-up`). Delivered (3-agent investigation of the current clones): **NEW
[`corpus/services/ai-readiness.md`](../../corpus/services/ai-readiness.md)** documenting the previously-undocumented
member-AI-readiness feature (org enablement gate, 3-step 30/40/30 scoring, the 9 `ai_readiness_*` tables, GraphQL+REST
interface, surfaces, narratives) **+ the M51 seeding contract** (Phase-2c-sharpened: active-cycle ⇒ signals-true,
closed-cycle ⇒ frozen-snapshot-direct — the dashboard recomputes from signals). Reconciled the material drift
(backend.md / next-web-app.md / architecture_overview.md / service_taxonomy.md now name the subsystem) + corrected the
false ant-academy "in repos.yml" claim (CLAUDE.md + ant-academy.md×3; **M49 #5 owns the repos.yml fix**). Docs-only —
never touched the demo. **Delivered → `corpus/services/ai-readiness.md`** + the re-grounded arch/service docs.

**M49 — Bring-up hardening + truth-up** · `section` · depends: **M47** · ∥ **M48** · ✅ **`done` — closed
2026-06-30** (merged → `release/01.10b-fit-up`; rext tag `fit-up-m49` @ `ba586d6` = 7 build fixes + 3 harden
commits, consumed per-stack). Closed the 7 remaining demo-up issues so a from-cold `/demo-up` on a `stack-demo`-only
box completes — **proven end-to-end by the live-verify gate** (a from-cold `/demo-up` on a re-pinned `fit-up-m49`
clone: demo-1 UP, autoverify "verified-working"): #3 `.env`-guard order (provision-then-check), #4
`INVITATION_HMAC_SECRET` (critical secret-DNA gene + values-blind auto-gen + `DemoGeneratedKeys` overlay; the silent
`app Exited (0)` class fixed), #5 ant-academy **explicit `ensure-clones.sh` clone** (NOT `repos.yml` — the ephemeral
gitignored platform clone makes that non-durable + a platform edit; the cms/studio submodule-pattern precedent), #6
disk pre-flight + `demo-down --purge` per-demo image cleanup (~5 GB reclaimed live), #7 *true* non-fatal frontend
(absent image → `--scale svc=0`), #8 demopatch re-anchor to next-web **v2.89.0**, #1 single **`.agentspace/rext.tag`**
source-of-truth (+ CRLF-tolerant `lib/rext_tag.sh` reader; reconciled 4 stale prose pins; doubles as the note-2
reproducible pin). Harden fixed 1 real bug inline (the `rext_tag.sh` CRLF carriage-return leak). Tests: rext Go
1552→**1555** (stack-secrets +3) · demo-stack Python **299** (demopatch 46→47); flake **0** (5/5). Close review: 6
findings all Fate-1 (2 stale `make init` ant-academy claims in `service_taxonomy.md` retired; test-count literals
reconciled); deferral audit **GREEN** (AI-keys policy → M50 Fate-2; consumption-clone re-pin → push-gated KEEP).
**Delivers →** `rosetta_demo.md`, `frontend-tier.md`, `secrets-spec.md`, `ant-academy.md`.

**M50 — Content & seeding fill** · `iterative` · depends: **M49** (+M48). ✅ **`done` — closed-on-gate 2026-06-30**
(merged → `release/01.10b-fit-up`; rext tag `fit-up-m50` @ `f0d984c` + close fix-commit `3c2de27`, consumed
per-stack). **M42 semantic coverage gate MET on BOTH vantages** (employee + manager) on a WARM demo-1, on the
manifest **STRENGTHENED to PROVE the gaps** (frontier-exhausted, (failingSections,escapes)=(0,0), 0 persona,
0 cross-port). 6 iters (1 tok + 5 tiks). Delivered (the sweep-driven seed fill): a NEW **`MemberLanguagesSeeder`**
(ISO-639-1 `world_languages` catalog + per-member `user_languages` → `membership_languages` via the platform's
AFTER-INSERT trigger — the manager Talent-tab "Languages spoken" chart, 0 rows → 747 across all 340 members), the
**`CertificatesSeeder` member-coverage extension** (hero-only → ~45% role-coherent, certs 9→236), the **`UsersSeeder`
member-field backfill** (`memberships.joined_at`/`location`/`last_activity_date` for `/enterprise/members`,
NULL-only idempotent guard), the **`next-web-public-website-url` demopatch** (the `PUBLIC_WEBSITE_URL` JS-constant
escape class) + a NEW **post-replay Directus content-URL rewrite** (the replayed-content escape class — prod hosts
baked into `public_landing_page_url`/`read_more_link`, regex over any `anthropos.work` subdomain → demo host), and
the **manager-manifest strengthening** (new `preAssert` tab-click + `textMatch` OR-assert harness primitives that
ASSERT the members-Location column + the Talent-tab languages/certs charts — the D4/F1 reconciliation: the run-1
gate passed BLIND to two M50-own gaps; the gate now PROVES them). All DATA-DENSITY, `PerStackIsolated` +
closure-GREEN, M17-idempotent, zero platform edits. **AI-keys policy DECIDED: documented-as-absent** (values-blind,
no key provisioned; AI surfaces inert-by-design — resolves the inherited M49 deferral). Tests: rext stack-seeding
719 (seeders pkg 349, +1 at close; 97.4% stmt) · demo-stack Python 108; flake **0** (5/5). Close review: 11
findings all Fate-1 (gofmt 2 files + 2 code pins + 3 docs incl the NEW routing-table escape row + a broken M51
backref); deferral audit **GREEN** (0 repeat/chronic/aged). **Carry-forward (three-fate, none escape-hatch):**
**COLD reset-to-seed acceptance → M53** (Fate-2, user-decided — all M50 seeders + fixes reproduce from the bring-up
tooling on a fresh `/demo-up`); **academy content + menu-link/non-anonymous-session (F6) → M51** (Fate-3);
consumption-clone re-pin to `fit-up-m50` = push-gated KEEP (authoritative at M53). **Delivers →**
`profile-completeness-spec.md`, `seeding-spec.md`, `coverage-protocol.md`, `secrets-spec.md`.

**M51 — AI-readiness showcase org** · `iterative` · depends: **M48** (the feature contract) + **M50**. ✅ **`done` —
closed-on-gate 2026-07-01** (merged → `release/01.10b-fit-up`). **Exit gate:** a curated **200-person 3rd org** with
the AI-readiness manager dashboard **enabled**, **~80%** of members having completed all **3** onboarding/evaluation
steps, **1 hero started + 1 hero completed** — proven by the coverage gate. **MET at iter-09:** manager-vantage
`(failingSections, escapes) = (0, 0)` frontier-exhausted on a fresh demo-up (reachable 70, personaFailures 0);
org **Northwind Aviation** (200 members) ENABLED, **78.4%** all-3-complete, **Ben STARTED** (stage 1) + **Aria
COMPLETED** (stage 3), cycle `closed` + 199 frozen snapshots. **9 iters** (1 bootstrap tok + 8 tiks). The
strategy arc: iter-02/03 landed the 3rd story + the 3 net-new seeders (`OrgSettingsSeeder` +
`AIReadinessConfigSeeder` + `AIReadinessFunnelSeeder`); iters 06/07/08 falsified three read-fast strategies
(active-signals → closed-cycle-snapshot → deep-link) against the platform AI-readiness read-path org-scale perf
wall (**"frozen SCORES ≠ frozen RESPONSE"** — `buildResponseFromSnapshots` re-joins members via an unbounded
whole-org `loadMembers`); **TOK-02** (user-authored, triggered by the 3-tik no-prog streak) pivoted to the
iter-09 **`app-aireadiness-snapshot-loadmembers`** read-path demo-patch (a PURE data-identical bound of that
hydration to the ~199 snapshot users → the frozen `?cycle=` GET 180s-timeout → 19ms). Tests: rext stack-seeding
**749** (seeders pkg 382, **97.6%** stmt, +30 vs M50's 719 across the iters + final harden + close) · e2e TS unit
**33** (+13 at close: the new `section-assert.ts` no-browser spec); flake **0** (5/5 Go + 5/5 TS). Close review:
16 findings all Fate-1 (S1 the 3rd AI-readiness story authored → **`delivers` MET**; C1/C3/C4 code + T1/T2 tests +
D1–D5 docs). Deferral audit **RED→CLEARED**: the academy **F6** repeat-defer (M50→M51) fated **LAND-NEXT → M53**
(Fate-3, user-decided — the cold rebuild is where academy content is seeded + verified). **Carry-forward (three-fate,
none escape-hatch):** academy F6 → M53 (Fate-3); COLD reset-to-seed acceptance → M53 (Fate-2); consumption-clone
re-pin + `.agentspace/rext.tag` bump → M53 (push-gated KEEP). **Delivers →** `demo/stories-spec.md` (the 3rd story),
`services/ai-readiness.md`, `seeding-spec.md`, `CLAUDE.md`; rext tag **`fit-up-m51`**.

**M52 — Single auditable seed+gen manifest** · `section` · depends: **M50 + M51** · ✅ **`done` — closed 2026-07-01**
(merged → `release/01.10b-fit-up`). Extract the Go mother-prompts to
YAML; author **one** checked-in `seed-generation-manifest.yaml` inlining population blueprint (all 3 orgs) +
prompts + batch config + snapshot sources (**cache + generated data excluded**); the cockpit **[Download]** serves
it. **Delivers →** NEW `corpus/ops/demo/seed-manifest-spec.md`.
All 4 sections landed (S1 `go:embed` extraction byte-identical → cache preserved; S2 the NEW `manifest` pkg + the
honesty-gated projection + `--manifest-export` verb; S3 cockpit [Download] repoint, non-fatal fallback; S4 the NEW
spec). Tests: rext stack-seeding **786** (+37 vs M51's 749; NEW `manifest` pkg 100% stmt) · demo-stack Python **313**
(+14: cockpit `--seed-manifest` endpoint + fallback); flake **0** (5/5 Go `-shuffle` + 5/5 Python). Close review: **12
findings all Fate-1** — F1 dedup the 3-way projection helper to one canonical `blueprint` source (removes the
projection-drift the honesty gate can't catch); F3 second cache-key golden fences the `{{else}}(none)` branch;
**F4 `mergeGenerationBatches` now WARNS on an orphan gen-story id** (a story-id typo was silently producing a
generation-less "auditable" manifest — the exact silent drop this milestone exists to prevent); F5 cockpit treats an
empty/blank `--seed-manifest` as absent; F2/F6/F7/F8/F9 (teeth-typo, gen-axis teeth, stale For-PMs prose, strip-helper
fence, vestigial doc-param). Deferral audit **GREEN** (up-injected.sh end-to-end glue = Fate-2 → M53's cold-rebuild;
0 repeats). **Carry-forward (three-fate, none escape-hatch):** up-injected.sh glue → M53 (Fate-2); consumption-clone
re-pin + `.agentspace/rext.tag` bump → M53 (push-gated KEEP). **Delivers →** `demo/seed-manifest-spec.md` (NEW) +
reconciled cross-refs (`cockpit-spec.md`, `ai-generation-spec.md`, `cache-spec.md`, `seeding-spec.md`, `README.md`,
`CLAUDE.md`); rext tag **`fit-up-m52`** (`36d7430`).

**M53 — Cold-rebuild acceptance** · `section` · depends: **M52** · ✅ **`done` — closed 2026-07-01**
(merged → `release/01.10b-fit-up`; **the FINAL v1.10b milestone**). Destroy the live demo + **rebuild from cold** on
a `stack-demo`-only box; assert healthy backends + complete set-dress/seed (all 3 orgs)/verify/cockpit + both-vantage
coverage + the AI-readiness criteria + the complete manifest download. Tag **`v1.10.1`** + bump `.agentspace/rext.tag`.
**All 6 sections landed** (§1 academy F6 seeder/wiring; §2 roll `v1.10.1`; §3 DESTROY via `/demo-down 1 --purge` — 17
containers + network + ALL demo-1 images, M49 #6 verified; §4 COLD REBUILD via a single `/demo-up 1`, no manual steps,
EXIT 0, no #7 abort; §5 ASSERT; §6 acceptance record + rext.tag bump). **Acceptance verdict: GREEN — 6/6 criteria +
academy F6 from cold** (AB1 backends healthy 17-up-0-exited/casbin 1150; AB2 prompt-free replay from the filled 1.4 GB
cache; AB3 set-dress+seed 3 orgs incl. Northwind AI-readiness + cockpit, EXIT 0; **AB4 both-vantage coverage GREEN**;
AB5 AI-readiness dashboard on Northwind — 50/100, 199 members, 3-step funnel, renders fast; AB6 cockpit [Download] =
complete inlined `seed-generation-manifest.yaml` 7593 B; **F6 academy** — real content + 9 cockpit [Academy] links →
authenticated member, Cosmo absent-by-design). **AB4 was RED on first assertion — an M51-owned gate regression
(M51 iter-05's unconditional ai-readiness manager `seedPath` broke the M50 base-org manager gate `dan-manager` @
Cervato); fixed at the acceptance gate** with user approval (an org-conditional manager manifest — `manifestFor(vantage,
expectedOrg)` returns the showcase `MANAGER_MANIFEST` only for Northwind, else `MANAGER_MANIFEST_BASE`; rext `117fe41`,
+3 unit tests; both manager vantages re-verified GREEN). Exactly the late cross-milestone regression M53 exists to catch
— the from-cold both-vantage assertion is the first joint re-measure of the M50 + M51 gates. Tests: rext stack-seeding
**791** (+5 vs M52's 786: F6 academy DeepLink/AcademyDeepLink build + harden single-source tests) · demo-stack Python
**326** (+13: F6 authenticated-session + [Academy] deep-link + harden `_academy_catalog_entry` edge/escape tests) · e2e
TS unit **29** (AB4 org-gating + referential-stability edges, +2 vs the pre-AB4 27); flake **0** (5/5 Go seeders shuffle
+ 5/5 Python cockpit+academy + 5/5 TS manifest). Close review: **2 findings, both Fate-1 docs** — DOC-1 documents the
AB4 org-conditional manager manifest in `coverage-protocol.md` (was undocumented); DOC-2 reconciles the stale
`~80%/≈160` AI-readiness prose to the shipped **78.4%/≈156-of-199** in `ai-readiness.md` (KB-2). Deferral audit **GREEN**
— every carry pointing at M53 LANDED here (up-injected.sh glue Fate-2 via the cold `/demo-up`; COLD acceptance = M53
itself; academy F6 = `e91f004`; box-level re-pin DONE); the historical academy-F6 REPEAT resolved by execution; 0
M53-originated deferrals, 0 escape-hatch. **Delivers →** the release acceptance record (`acceptance-record.md`, feeds
`/developer-kit:close-release`) + `.agentspace/rext.tag` = `v1.10.1`; rext release tag **`v1.10.1`** (`576dbcb` — rolls
up `fit-up-m47..m52` + F6 + AB4 + the M53 harden tests; re-rolled at close, local unpushed annotated re-roll). **Sole
residual = origin push (push-gated KEEP, orchestrator/user).** → **v1.10b is GREEN from cold; the release is complete →
`/developer-kit:close-release`.**

### Top risks

- **clone re-sync (M47) → RETIRED (was flagged the biggest unknown).** M47 found the clones already **current**
  (next-web @ v2.89.0, every repo ≤2 behind) — the re-sync was a trivial `make pull`, no 5-week catch-up, no
  cascading breakage. The snapshot recapture confirmed **both schema digests unchanged** (taxonomy `c75ce94…`,
  directus `ea2e187…`), so it was a clean in-place data refresh. The heavy-rebuild risk did not materialize.
- **content root-causes are hypotheses → degrades-quality.** M50 is iterative; it starts with a fresh observation
  pass on the clean bring-up, not the static guesses. Several "empty" surfaces may be demo-up #7 artifacts.
- **AI-readiness data model unknown → M51 exploratory.** The feature was invisible to the stale clones; M48
  documents it before M51 seeds it.
- **1-demo constraint → sequential chain, longer wall-clock.** No parallel verification; only rext authoring
  parallelizes (worktrees). M53 is the single from-cold acceptance truth.
- **M52 cache integrity.** Extracting prompts to YAML must preserve or deliberately re-key the M45 prompt-hash cache.
- **AI-provider keys → decision (M49/M50).** Which become throwaway/sandbox demo values vs documented-as-absent.

---

## Done — v2.0 "opening night" (SHIPPED 2026-07-02, tag `v2.0`)

> **Status (SHIPPED 2026-07-02, tag `v2.0`):** v2.0 "opening night" — the **Playthroughs** pillar — is **shipped**
> and merged to `main`. All four milestones closed: M201 "Manifest corpus" (`done`, closed-on-gate), **M202
> (Foundation, `section`, `done`** closed-complete 2026-07-01, tag `opening-night-m202`), **M203 (Employee-vantage
> coverage, `iterative`, `done`** closed-on-gate 2026-07-02 — 6/6 employee Playthroughs GREEN on cold reset-to-seed,
> 5/5 deterministic; tag `opening-night-m203`), and **M204 (Manager-vantage coverage, `iterative`, `done`**
> closed-on-gate 2026-07-02 — 4/4 manager Playthroughs GREEN on cold reset-to-seed, 5/5 deterministic; tag
> `opening-night-m204`) — the FINAL v2.0 milestone; it imported M203's shared page-object layer + ran on the
> reset-to-seed lifecycle per `corpus/ops/demo/playthroughs.md`. **Corpus at ship: 10 live Playthroughs (6 employee
> + 4 manager) GREEN on cold reset-to-seed, 1 declared in-manifest TODO** (the assign-WRITE half → Fate-2 → a future
> manager-write tier). Close-release: all 9 review sweeps GREEN, no blockers; rext code-of-record rolled to `v2.0`.
> Records archived under [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
>
> **Theme:** *the platform's core user journeys, proven to actually work.* A **Playthrough** is an automated
> actor that **is the user** — it logs in as a seeded hero, sets out with a goal, plays through a real journey
> across the platform start-to-finish the way a person would, then proves the platform delivered the outcome.
> The capability is the **canonical, living set of these journeys**: the platform's user-facing functionality,
> continuously **proven to actually work** — cleanly decoupled from *"the pixels are identical"* (a Playthrough
> breaks **only when a capability breaks**). It is the **functional** sibling of v1.x's M42 coverage sweep
> (which proves *presence* — every reachable page **shows** real content); Playthroughs prove the hero can **do**
> the things that world is for.
>
> **Designed 2026-06-28** via `/developer-kit:design-roadmap`, from the consolidated capability spec
> [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (v0.3). **A new MAJOR** — Playthroughs is
> a new pillar distinct from the demo/seeding lineage; v2+ uses **`Mxyy`** milestone numbering. **Tooling + docs
> only — zero platform-repo edits** (the read-only platform line carries over; an un-drivable surface escalates
> via the `unimplementable-without-platform-edit` state, it never edits the platform).

### Execution graph

```
v2.0 "opening night"
  M201 ──┐                          (manifest corpus — prose, user-guided)
  M202 ──┼──→ M203 ─┐
                M204 ─┴──→ ship
```

**M201 (the manifest corpus) and M202 (the foundation) open in parallel.** M201 is the **user-guided manifest
curation** — prose-only (the goal-aligned Product → Story → Use Case corpus), so it carries **no code dependency**
and can be authored before / alongside M202. M202 is the **Playthroughs foundation** (the section, the manifest
model + the §5.3 **validator**, the page-object layer, the dedicated seed + reset lifecycle, the runner + 4-state
reporting, one trivial proof Playthrough) — it builds the validator + dedicated seed to **match** the M201 corpus.
Then the two **vantage-coverage** milestones — **M203** (employee) and **M204** (manager) — run **in parallel**,
both `iterative`, implementing Playthroughs against the M201-declared use cases on the M202 foundation; the release
ships when both gates fire.

**M201 ∥ M202 (manifest ∥ foundation).** No hard ordering: M201 produces the **prose contract** (the use-case
manifest); M202 produces the **machinery** (validator + dedicated seed) that validates + seeds against it. They
reconcile when M202's validator runs over the M201 corpus. Where an M201 use case names a **precondition the demo
seed lacks**, that feeds the **M202 dedicated-seed expansion** (M201 records the need; M202 builds the seed).

**Parallelism note (M203 ∥ M204).** The two coverage milestones share an **additive merge surface**: the
per-surface **landmark registry** + the **locator index** (the §5.6 page-object layer every Playthrough imports).
Each vantage adds its own surfaces/anchors to that shared layer — an **additive** merge, not a conflicting one.
Both are `iterative` (the use-cases are *declarable* in the M201 corpus, but getting them green against the real
antd UI + the AI-sim assertion boundary is exploratory, like M42e/M42m), so they advance independently toward
their own exit gates and reconcile the registry additively at merge.

### Milestones

**M201 — Manifest corpus** · `iterative` · **USER-GUIDED** · complexity **large** · depends on: **none** (the
manifest is prose — authorable before/parallel to the M202 foundation).
**Status:** ✅ **`done` — closed-on-gate 2026-06-29.** 9 products · 26 stories · 28 use-cases authored,
**adversarially re-grounded** (11-agent wf `wvpnpvozh` → 15/27 runnable), **user-signed-off**. Records:
[`releases/archive/02.00-opening-night/m201-manifest-corpus/`](releases/archive/02.00-opening-night/m201-manifest-corpus/)
(deliverable: `manifest-draft.yaml`). The close discovered the **stale-clone drift** (next-web 115+ commits behind
prod) → **v2.0 PAUSED for the v1.10b "fit-up" backfill** (re-sync + re-ground + re-validate + fix, user-driven; see
the In-Development section above) before resuming. *(Correction: the v1.10b M47 milestone later found the clones
actually current — next-web @ v2.89.0; the stale surface is the corpus, re-grounded by M48. The pause stands; the
backfill's real work is the corpus + the bring-up/content fixes, not a clone re-sync.)*
**Goal:** top-down, **user-directed** curation of the **full goal-aligned Product → Story → Use Case manifest
corpus** — the build + regression contract every coverage milestone (M203/M204) implements against. The flow per
pass: **outline** (products → stories → use cases) → **validate** (against the real platform surface + the spec's
manifest model) → **write the prose-intent manifest YAML** (spec §5.3, **one file per product**).
**Explicitly NOT bounded by the current minimal/partially-aligned demo stories seed** — it captures what the goal
says must be proven; where a use case needs preconditions the demo lacks, that **feeds the M202 dedicated-seed
expansion** (noted, not resolved here).
**Shape:** `iterative`, **driven by the user** — worked conversationally + via `/developer-kit:work-mstone-iters`,
not autonomously.
**Exit gate:** **the manifest corpus is comprehensively outlined, validated, and written as prose-intent YAML (one
file per product)** — covering the platform's products × their must-work user journeys, each use case carrying
**goal + actor + flow + intermediate/final expectations**, structurally valid (the spec §5.3 validator passes,
**ids unique + both-way**) — **and the USER signs off the corpus as the complete-enough v2.0 coverage contract.**
**iteration_protocol_ref:** the capability spec
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (esp. §2 model, §4 use-case shape, §5.3
manifest format).
**Delivers →** the prose-intent manifest YAML corpus (one file per product); **lands in the rext `playthroughs`
section once M202 exists**, until then drafted under the milestone dir / `spec-drafts/playthroughs/manifest-draft/`.
**Candidate starting outline (the user directs — NOT fixed):** (a) the demo-covered products — **Skill Paths, AI
Simulations, Profile & Skills, Workforce Intelligence, Hiring, Academy**; (b) goal-aligned areas the demo barely
covers (flag *to confirm with the user*) — **Auth & Onboarding, Billing & Entitlements/tier-gates, Org Admin &
Settings, cross-product journeys** (candidate→employee).

**M202 — Playthroughs Foundation** · `section` · complexity **medium** · depends on: **none** (reuses the M42
harness + the seeding machinery; the M201 manifest corpus is its build+regression contract, authorable in parallel).
**Status:** ✅ **`done` — closed-complete 2026-07-01.** All 6 sections + the NEW `corpus/ops/demo/playthroughs.md`
runbook delivered; the trivial proof Playthrough (login → /profile → assert hero identity) **GREEN on demo-1**. The
`playthroughs` rext section: manifest model + light validator (both-way id integrity + precondition-coverage +
datadna closure gate) · per-surface page-object layer (1 surface: `/profile`, re-pin O(surfaces)) · dedicated
decoupled `pt-world` seed (2 private orgs, entitlement tiers + multi-org-private) · reset-to-seed lifecycle +
serial-default runner · 4-state reporting map. **96 Go test/fuzz funcs (98.5% section) + 13 TS** (5/5 flake-clean).
Close surfaced 8 findings, all Fate-1: CQ1 datadna exit-3 diagnosis · CQ2 `PW_WORKERS` serial-safety guard · CQ3
`truncate` totality · CQ4 `ptTagRe` lockstep · DOC1 section-index · DOC2 fixtures wording · M202-D4 anchor-story
landmine blended into `stories-spec.md`. Deferral audit **GREEN** (0 milestone-owned). Tooling + docs only — **zero
platform edits, zero new deps**. rext authoring @ `b1e5528`, tagged **`opening-night-m202`**. The runbook IS the
M203/M204 `iteration_protocol_ref`. Records: [`releases/archive/02.00-opening-night/m202-foundation/`](releases/archive/02.00-opening-night/m202-foundation/).
**Goal:** stand up the **`playthroughs` rext section** on the **shared M42 e2e foundation**, proven by **one
trivial end-to-end Playthrough**.
**Scope — In:**
- the **manifest model + a light validator** — both-way id integrity (use-case ↔ Playthrough, traceable both
  directions) + precondition-coverage (every declared `seed`/`preconditions` resolves to a named seeded world,
  no silent "ideally"), **datadna-gated** (the Playthrough seed is covered by the same `datadna` conformance gate
  as the seeding machinery);
- the **per-surface locator/landmark page-object layer** (the §5.6 shared registry every Playthrough imports —
  a UI/antd/copy shift is absorbed by editing the per-surface registry, not N tests) — **1 surface to start**;
- the **dedicated, decoupled seed** preset (test data ≠ demo data; the demo seed is the *starting point* but
  kept separate) — **spans entitlement tiers + multi-org-private**;
- the **reset-to-seed lifecycle + serial-default runner** (`workers: 1`, `fullyParallel: false`; reset via the
  real `--reset` path honoring its contract — **additive re-seed is FORBIDDEN as a reset**);
- the **4-state reporting map** — `passing` / `failing` / `unimplemented` / `unimplementable-without-platform-edit`
  (the last being the P3 zero-edit escape valve — it escalates, never edits);
- **one trivial proof Playthrough** — **login → /profile → assert hero identity** (the foundation's smoke test).
**Out:** real product coverage (M203+); the AI-sim / integration mirror tier; cross-vantage.
**Delivers →** a corpus runbook that **graduates the playthroughs spec** (e.g.
[`corpus/ops/demo/playthroughs.md`](../../corpus/ops/demo/playthroughs.md)) — becomes the `iteration_protocol_ref`
for M203/M204.
**KB deps (read as contract):** the playthroughs spec-draft
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md), the **M201 manifest corpus**
([`releases/archive/02.00-opening-night/m201-manifest-corpus/`](releases/archive/02.00-opening-night/m201-manifest-corpus/) — the
prose contract the validator + seed implement against),
[`corpus/ops/demo/coverage-protocol.md`](../../corpus/ops/demo/coverage-protocol.md),
[`corpus/ops/seeding-spec.md`](../../corpus/ops/seeding-spec.md),
[`corpus/ops/idempotency.md`](../../corpus/ops/idempotency.md).
**Reuse paths (cite in spec-notes):** `stack-demo/rosetta-extensions/stack-verify/e2e/lib/{cockpit-login,
section-assert,empty-states,coverage-manifest}.ts`, `stack-demo/rosetta-extensions/stack-seeding/`.

**M203 — Employee-vantage coverage** · `iterative` · complexity **large** · depends on: **M202** ·
parallel-with: **M204**. ✅ **`done` — closed-on-gate 2026-07-02** (merged → `release/02.00-opening-night`; rext
tag **`opening-night-m203`** @ `fb94458`). **Exit gate MET at iter-06:** every declared employee-vantage use case
has a PASSING Playthrough on a COLD reset-to-seed demo (**6/6** — Profile identity+verified+growth+timeline · Skill
Paths legacy learn&progress · AI Simulations chat launch §5.8), with **0 false-fails over 5 consecutive reset runs**
(5/5). **6 iters** (1 bootstrap tok + 5 tiks, all closed-fixed). The strategy arc (TOK-01, deterministic-read-first):
iter-02/03 landed the full Profile journey (Spotlight + claimed-vs-verified gap + trajectory + work/education
timeline), iter-04 Skill Paths (browse→open→start→progress; verify-skill composes P7 on the profile side), iter-05
AI Sims chat launch + the **post-seed Sentinel-Reload** fix (a casbin g3 feature grant is only effective after the
enforcer RELOADS — folded into `run-playthroughs.sh`, drift-guarded), iter-06 the 5-run determinism gate (runnable
from the authoring-built `stackseed`). Tests: rext playthroughs Go **103** (+7 vs M202's 96: @pt-tag lockstep ×2
packages + invalid-engine + read-error arms + the iter-05 Sentinel-reload drift guard) · e2e TS **38** unit (+25
vs 13) + **6** browser Playthroughs; flake **0** (Go 5/5 -shuffle + TS unit 5/5; browser 5/5 cold-reset iter-06).
Close review: **11 findings all Fate-1** — F1/F5 (must-fix) the segment-anchor route-shape fix was NOT applied to
two inline `\b`-terminal copies (a load-bearing `waitForURL`) → centralized into `url-shapes.ts` `SKILLS_TAB_URL`;
F3 dropped 4 dead speculative self-eval accessors; TEST-G1 the unenforced `@pt:`-grammar lockstep → twin tests;
TEST-G3/G4 two coverage arms → 100%; DOC README + runbook M203 backfill. Deferral audit **GREEN** (0
repeat/chronic/aged). **Carry-forward (three-fate, none escape-hatch):** the 4 **non-gate** edge UCs
(`ai-simulations.code.UC1` Judge0 · `ai-simulations.interview.UC1` text · the Skill-Paths verify-skill terminal ·
`profile.self-evaluation.UC1`) → **Fate-3 → M206** (roadmap-vision annotated; the gate enumerated the 3 CORE
journeys, all GREEN — these are additional). Academy UC OUT by design; voice sims → M206 by design. Tooling + docs
only — zero platform edits, zero new deps. Records:
[`releases/archive/02.00-opening-night/m203-employee-coverage/`](releases/archive/02.00-opening-night/m203-employee-coverage/).
**Goal (as designed):** **Maya's** core **employee** journeys play green (declared in the M201 manifest corpus) —
Skill Paths (browse → enroll → complete → verify-skill), AI Simulations (chat/code launch → complete →
score-in-range, **NON-voice**), Profile (verified-skill chart + the claimed-vs-verified gap + work/education
timeline). **Why iterative:** the use-cases are *declarable* (in the M201 corpus), but getting them green against
the real antd UI (the landmark layer) + the AI-sim assertion boundary is **exploratory**, like M42e.

**M204 — Manager-vantage coverage** · `iterative` · complexity **large** · depends on: **M202** ·
parallel-with: **M203**. ✅ **`done` — closed-on-gate 2026-07-02** (merged → `release/02.00-opening-night`; rext
tag **`opening-night-m204`** @ `c81c6dd`). **THE FINAL v2.0 milestone.** **Exit gate MET at iter-05:** every
declared manager-vantage use case has a PASSING Playthrough on a COLD reset-to-seed demo (**4/4** — Workforce
funnel + org-scale gap · member roster · per-member activity-dashboard drill-down · succession/at-risk), with **0
false-fails over 5 consecutive reset runs** (5/5). **5 iters** (1 bootstrap tok + 4 tiks, all closed-fixed). The
strategy arc (TOK-01, manager-surface-per-iter): iter-02 the Workforce funnel + roster (the manager landing
surface) + the runner **reporter-override** fix (a stale-JSON gate hazard, drift-guarded), iter-03 the
activity-dashboard per-member drill-down (a SPA-URL race + an out-of-`<main>` table scope fixed), iter-04
succession/at-risk (all 3 declared journeys green — no seed expansion needed, Org A size-40 rendered the M36
aggregates), iter-05 the 5-run determinism gate. All 4 manager UCs are READ/monitoring flows → the risk was
seed-scale render + antd grid ambiguity, not mutation-determinism. Tests: rext playthroughs Go **105** (+2 vs
M203's 103: the reporter-override drift guard + the manager-coverage deliverable-presence pin) · e2e TS **58**
unit (+20 vs 38) + **4** browser Playthroughs; flake **0** (Go 5/5 -shuffle + TS unit 5/5; browser 5/5 cold-reset
iter-05). Close review: **5 findings all Fate-1** — 1 code-quality should-fix (documented the manager
predicate-API's deliberate symmetric shape — kept, not pruned, consistent with the M203 `isOnSkillsTab`
precedent), 3 docs (flip the stale "M204 adds" future-voice → landed; corpus now **10 live Playthroughs, 1 TODO**;
add the M204 manager page-object bullet), 1 decision-triage (D-CLOSE-1). Deferral audit **GREEN** (0
repeat/chronic/aged). **Carry-forward (three-fate, none escape-hatch):** `assignment-monitoring.assign-and-track.UC1`
(the assign-WRITE half — a two-backend org-admin WRITE flow) → **Fate-2**, tracked in-manifest as `playthrough:
TODO` (reports `unimplemented`, presence-pinned; out of M204's declared 3 monitoring journeys; a future
manager-write tier is its home). Tooling + docs only — zero platform edits, zero new deps. Records:
[`releases/archive/02.00-opening-night/m204-manager-coverage/`](releases/archive/02.00-opening-night/m204-manager-coverage/).

**Goal:** **Dan's** core **manager** journeys play green (declared in the M201 manifest corpus) —
- **Workforce funnel** + member roster,
- **member drill-down** (the activity-dashboard),
- **succession / at-risk** (the Growth tab) signals.
**Exit gate:** **same shape as M203, manager-vantage** — every declared manager-vantage use case has a passing
Playthrough on a COLD reset-to-seed demo stack, with 0 false-fails over 5 consecutive reset runs.
**iteration_protocol_ref:** same as M203 (the spec / the M202-delivered runbook).
**Why iterative:** same as M203 — declarable use-cases, exploratory path against the real manager UI + the
assertion boundary.
**Re-scope trigger:** same — `unimplementable-without-platform-edit` → escalate, don't edit.

### Top risks

- **manifest completeness → no auto-gate, user owns "enough".** The M201 manifest is a **build reference** with
  **no introspectable schema for "what user-facing capabilities exist"** (spec §5.9) — an *added* platform
  capability with no use case cannot be auto-detected. The corpus's completeness is a **user judgement** (the M201
  exit gate's sign-off), not a machine check. *Mitigation:* M201 is **user-guided + iterative** (the user directs
  each top-down pass + signs off the complete-enough contract); the cadence-review stance (§5.9) carries forward.
- **antd-a11y → the landmark layer is load-bearing.** zero-platform-edit means we **cannot** add `data-testid`;
  locators bind to the **accessibility tree** (`getByRole`/`getByLabel`/`getByText`) over the **real antd UI**,
  with a Rosetta-side **landmark registry** for ambiguous surfaces. If antd's a11y is thin on a surface, that
  surface's landmark anchors carry the test — the registry's quality is the risk. *Mitigation:* the per-surface
  page-object layer (re-pin is O(surfaces), not O(tests)); start with **1 surface** in M202 to prove the pattern.
- **determinism-under-mutation → M202's reset must be solid.** P6 ("same inputs → same result") holds **only if**
  the world is reset to the known seed between runs, and an *additive* re-seed silently leaves stale state (the
  M42e "green-but-wrong" trap). The whole determinism headline rests on M202's **reset-to-seed lifecycle** being
  correct — it is a **foundation** risk, surfaced and owned in M202 before M203/M204 lean on it.
- **hero-login → demo-N only.** Hero-driven Playthroughs run on **demo-N** (or a Clerkenstein-injected dev-N) —
  a plain dev-N is real Clerk + one identity + `dev-min`, so the hero suite is **not** the same on dev today. The
  target environment is the demo stack; the dev-stack hero-flow generalization is an explicit **later** item
  (spec §5.4), not v2.0 scope.
- **AI-sim mirror tier is future.** The signature voice/recording AI-simulation journey needs a **mirror engine**
  (Clerkenstein mocks **only** Clerk — no LiveKit/Chime/Stripe/Brevo mirror). v2.0 covers the **NON-voice**
  chat/code/document sims (playable as-is, asserted at the launch/completion boundary); voice + recording +
  payments + email are parked as `later — needs a mirror engine` → **M206** ([`roadmap-vision.md`](roadmap-vision.md)).

---

## Shipped releases

- **v2.0 "opening night"** — **2026-07-02**, tag `v2.0`, **4 milestones (M201 … M204)**. The **Playthroughs**
  pillar: a manifest-driven, deterministic e2e suite that plays real user journeys and proves the platform delivers
  the outcome (**function**, vs the M42 coverage sweep's **presence**). Manifest corpus → foundation → employee +
  manager coverage. **10 live Playthroughs GREEN on cold reset-to-seed, 1 in-manifest TODO.** The **first v2.x
  release**. Tooling + docs only, zero platform edits, zero new deps. Records archived under
  [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
- **v1.10.1 "fit-up"** — **2026-07-01**, tag `v1.10.1`, **7 milestones (M47 … M53)**. The interposed
  **field-hardening backfill** (a `.1` patch on v1.10): re-sync + recapture, corpus re-ground, from-cold `/demo-up`
  hardening, content + AI-readiness-showcase-org seeding fill, one auditable seed+gen manifest, then a from-cold
  destroy-and-rebuild acceptance (**6/6 + academy F6 GREEN**). Tooling + docs only, zero platform edits, zero new
  deps. Records archived under
  [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).

The complete earlier shipped history — **v1.0 "body double"** (2026-06-03, tag `v1.0`) through **v1.10 "method acting"**
(2026-06-27, tag `v1.10`), 11 versions / milestones M0 … M46 — is preserved in
[`roadmap-legacy.md`](roadmap-legacy.md) (the retired v1.x major). Records are archived under
[`releases/archive/`](releases/archive/). **v2.0 "opening night" (tag `v2.0`) is the first shipped v2.x release**
(2026-07-02); the next v2.x release awaits `/developer-kit:design-roadmap`.

## Notes

- **Milestone numbering — v2+ uses `Mxyy`** (`M` + major digit + two-digit milestone): **M201, M202, M203, M204**
  for v2.0. This is the major-version scheme `context.md` reserved for *"a future *major* v2+"*; the v1.x flat
  sequential counter (M0 … M46, with the `a`/`b`/`c`/`e`/`m` suffix conventions) lives in
  [`roadmap-legacy.md`](roadmap-legacy.md) § Notes. **It was thought closed at M46, but the interposed v1.10b
  "fit-up" backfill RE-OPENS it at M47** (M47→M53) — the backfill is v1.x-major work (a `.1` patch of v1.10), not a
  v2 milestone, so it keeps the flat counter rather than `Mxyy`.
- **Milestone shapes** mix within v2.0: **M201 is `iterative` + USER-GUIDED** (the manifest corpus — a top-down,
  user-directed prose curation toward a sign-off gate); **M202 is `section`** (a fixed In-scope checklist — the
  foundation is decomposable up front); **M203 + M204 are `iterative`** (a measurable exit gate, exploratory path
  — getting declared use-cases green against the real antd UI + the AI-sim assertion boundary, like the M42e/M42m
  precedent).
- Date format throughout: ISO `YYYY-MM-DD`.
- The Playthroughs capability **graduated from spec-draft to active development** at v2.0 design (2026-06-28); the
  governing spec is [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md), graduated to a corpus
  runbook (`corpus/ops/demo/playthroughs.md`) by M202.

_Last updated: 2026-06-29 (**v1.10b "fit-up" DESIGNED + PROMOTED** via `/developer-kit:design-roadmap` — an
interposed **field-hardening backfill** [the v1.3b "dress rehearsal" lineage]; **7 milestones M47 → { M48 ∥ M49 } →
M50 → M51 → M52 → M53** re-opening the v1.x flat counter; branch `release/01.10b-fit-up` cut from `main`; tag
`v1.10.1`. Designed from the field review `.agentspace/annotation.md` + the M201 stale-clone finding [3 research
agents]. Re-grounds demo + corpus to current prod, fixes the from-scratch `/demo-up` issues + the v1.10 content
gaps, adds the AI-readiness showcase org, consolidates one auditable seed+gen manifest. **v2.0 "opening night"
PAUSED** until it ships. Tooling + docs only — zero platform-repo edits. Prior: 2026-06-28 **v2.0 "opening night"
DESIGNED + PROMOTED** — a NEW MAJOR opening the **Playthroughs** pillar; 4 milestones M201 ∥ M202 → { M203 ∥ M204 };
branch `release/02.00-opening-night`; from `spec-drafts/playthroughs/spec.md` v0.3.)_
