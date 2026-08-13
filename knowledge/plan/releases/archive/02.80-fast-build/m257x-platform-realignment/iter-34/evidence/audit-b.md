# iter-34 confirming pass — audit group B

Auditor: group B. Method: every assigned file read top-to-bottom, in full, no grep-scoping.
Verification substrate: `stack-demo/platform` @ `2adcf71`, plus the peer clones it drives
(`stack-demo/app` @ `5ba17044` v1.363.2, `stack-demo/next-web-app` @ `bb3313bc0` v2.133.0,
`stack-demo/graphql-wundergraph` @ `60c229f`, `stack-demo/{cms,jobsimulation,roadrunner,sentinel,storage,messenger,studio-desk}`),
and the rext authoring copy at `.agentspace/rosetta-extensions` for the alignment-tooling claims.

## Positive control

| file | `wc -l` | last line actually read |
|---|---|---|
| `corpus/services/ai-readiness.md` | 597 | **597** |
| `corpus/architecture/alignment_testing.md` | 516 | **516** |
| `corpus/architecture/service_taxonomy.md` | 411 | **411** |
| `corpus/services/README.md` | 79 | **79** |
| `corpus/services/TEMPLATE.md` | 46 | **46** |
| `corpus/services/db-backup.md` | 31 | **31** |
| `corpus/services/intelligence.md` | 18 | **18** |

Total 1698 lines. **No file UNREAD.**

---

## BLOCKERS

### B1 — `corpus/services/ai-readiness.md:402-406`

**Verbatim false text:**

> **⚠ The frozen path is CYCLE-SCOPED; the DEFAULT (`CycleID == nil`) GET does NOT take it.**
> `GetAIReadinessWithOptions` (`ai_readiness.go:283-301`) reaches `buildResponseFromSnapshots` **only**
> when the request carries `opts.CycleID != nil` AND that cycle's `status == "closed"`; the **default
> GET** (line 301) is hardcoded to `buildLiveResponse`.

**What is actually true at HEAD.** The default (`CycleID == nil`) path *does* reach
`buildResponseFromSnapshots`. `GetAIReadinessWithOptions` has a **second, cycle-less** frozen branch:
when the org has **no active cycle** and a latest-closed cycle exists, it returns the frozen response.
The default GET is therefore *not* "hardcoded to `buildLiveResponse`" — that is only the fall-through
after both the cycle-scoped branch and the no-active-cycle branch miss.

**Proof:** `app/internal/aireadiness/readiness.go:289-314` —

```go
// No cycle requested: active cycle -> live (cycle-scoped); no active cycle ->
// last closed cycle snapshot (the last official measurement).
if m.activeCycle(ctx, orgID) == nil {
    if closed := m.latestClosedCycle(ctx, orgID); closed != nil {
        return m.buildResponseFromSnapshots(ctx, orgID, closed, opts)
    }
}
return m.buildLiveResponse(ctx, orgID, opts)
```

**Why it misdirects real work.** This paragraph is the stated premise for two downstream contracts in
the same file: the seed strategy (§ *The CYCLE-STATE contract*) and the M51 iter-08/09 perf analysis of
`buildResponseFromSnapshots → LoadMembers` (`readiness.go:779`, the unbounded whole-org hydration that
the `app-aireadiness-snapshot-loadmembers` demo-patch exists to bound). A reader debugging a
closed-only org — the exact shape M51 shipped — is told the slow frozen path is *unreachable* without
`?cycle=`, when in fact that org's **default** GET takes it every time. It also tells anyone seeing
frozen/stale scores on an unparameterised GET that what they are observing is impossible.

**Grade: BLOCKER.**

**Suggested correction:** replace with — *"`buildResponseFromSnapshots` is reached in TWO cases:
(a) `opts.CycleID != nil` and that cycle is `closed`; (b) `CycleID == nil` AND the org has no active
cycle AND a latest-closed cycle exists (`readiness.go:307-312`). Only when neither holds does the
request fall through to `buildLiveResponse`."*

---

### B2 — `corpus/services/ai-readiness.md:58` + `:66-67`

**Verbatim false text:**

> The feature is off until an org turns it on. Two gates compose (both must be true for the UI to render):
> …
> 2. **PostHog flag** `flag_ai_readiness` — the next-web client also gates the route on this flag
>    before it even queries `aiReadinessEnabled`
>    (`apps/web/.../ai-readiness/AIReadinessClient.tsx`).

**What is actually true at HEAD.** `AIReadinessClient.tsx` — the **manager** dashboard at
`/ai-readiness` — contains **no PostHog reference of any kind**. It calls
`useAiReadinessEnabled(true)` unconditionally and derives `featureOn = orgEnabled === true`
(`AIReadinessClient.tsx:133-134`); every query and the whole render gate on that single GraphQL
boolean (`:138`, `:169`, `:176`, `:323`). The **only** consumer of `AI_READINESS_FLAG` in the entire
repo is `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:22`, which feeds the
**employee** `/home` embed (`components/ai-readiness/useAIReadiness.ts:77`) and the
onboarding/reimport redirects. So gate 2 does **not** apply to the route this line cites, and the two
gates do **not** both have to be true "for the UI to render" — they both have to be true for the
*member* surface only.

**Proof:**
- `next-web-app/apps/web/src/app/(authenticated)/(verified)/ai-readiness/AIReadinessClient.tsx:133-134`
  (`const { orgEnabled } = useAiReadinessEnabled(true); const featureOn = orgEnabled === true;`)
- exhaustive repo grep for `AI_READINESS_FLAG` → 4 hits, none in the `/ai-readiness` route directory:
  `aiReadiness.constants.ts:9,26` (definition) and `data/useAiReadinessActive.ts:7,22` (sole use).

**Why it misdirects real work — and it is self-refuting.** This is the doc's *foundational*
"Org enablement (the gate)" section, and it names the wrong component. It is the identical
**wrong-vantage** error this same file spends lines 98-103 diagnosing and lines 292-294 correcting
("*The `flag_ai_readiness` PostHog flag gates the EMPLOYEE side only … The (one) manager dashboard
gates purely on the GraphQL `aiReadinessEnabled` boolean*"). A reader who stops at the gate section —
the natural place to start — concludes the manager dashboard needs a PostHog key or the
`next-web-aireadiness-flag-gate` demo-patch to render. It does not, and chasing that is precisely the
milestone-cost M219 recorded.

**Grade: BLOCKER.**

**Suggested correction:** scope gate 2 to the member surface and re-cite it — *"2. **PostHog flag**
`flag_ai_readiness` — gates the **EMPLOYEE** surface only (`components/ai-readiness/data/
useAiReadinessActive.ts:22`). The manager dashboard at `/ai-readiness` does not consult PostHog at
all; it gates on `aiReadinessEnabled` alone (`AIReadinessClient.tsx:133-134`). Both gates compose for
the member `/home` embed; only gate 1 applies to the manager dashboard."* — and change the lead-in at
`:58` from "both must be true for the UI to render" accordingly.

---

## Minors

### `corpus/services/ai-readiness.md`

| # | line | text | truth at HEAD | proof |
|---|---|---|---|---|
| m1 | `:197` | "the **four** usage KPIs are DERIVED…" | there are **five** (`avg_adoption/avg_transformation/avg_originality/avg_depth/avg_ownership`) — as the same file states at `:185` and `:567-570` | `app/internal/aireadiness/how_we_measure.go:1138-1156` (`usageDimSpecs`, 5 entries) |
| m2 | `:388` | "`keepStartedMembers` **excludes members with no step-1 signal**" | it excludes members with **no `ai_readiness_user_step_progresses` row whose status is past `not_started`** — it never looks at step-1 evidence | `readiness.go:684-698` → `steps.go:907-926` (`SELECT DISTINCT user_id FROM public.ai_readiness_user_step_progresses … AND status <> 'not_started'`) |
| m3 | `:142` | "a member carries a `stage` ∈ {1,2,3} (**0 = none/done**)" | stage 0 means *not started*; **done is stage 3** | `app/internal/aireadiness/types.go:77` (`Stage int // 1/2/3 — how many steps done`); the file's own `:366` calls stage 0 "no evidence, no sessions, no progress row" |
| m4 | `:284` | "consumed by `packages/ui/src/NavBar/useNavbarSections.tsx:253-260`" | the `AI_READINESS_URL` menu item is at **`:398-400`**, consumed at **`:547`**; `:253-260` is `librarySkillPathsMenuItem` | `next-web-app/packages/ui/src/NavBar/useNavbarSections.tsx` |
| m5 | `:297` | "`AIReadinessClient.tsx:69` `const SHOW_SECONDARY_TABS = false;`" | it is at **`:78`** (claim itself is true) | `AIReadinessClient.tsx:78`, used at `:599` |
| m6 | `:415`, `:416` | "`AIReadinessClient.tsx:137-138`" / "gates the data GET on `cyclesQ.isFetched` (`:150-154`)" | `effectiveCycleId` is at **`:153-154`**; the `cyclesQ.isFetched` gates are at **`:169`/`:176`** (claims true) | same file |
| m7 | `:541` | "(`HowWeMeasureTab.tsx:2773-2797`…)" | the file is **1989 lines** — the anchor is past EOF; the tile is at **`:1915-1927`**. The substantive claim (renders `skillsMapped`/`handsOnMinutes`/`interviewMinutes`, never `interviewQuestions`) is **TRUE** | `HowWeMeasureTab.tsx:1915,1921,1927`; repo-wide grep: `interviewQuestions` appears only at `apps/web/src/hooks/useAIReadiness.ts:274` |
| m8 | `:543` | "the FE's TypeScript type, `useAIReadiness.ts:250`" | it is `apps/web/src/hooks/useAIReadiness.ts:**274**` | same |
| m9 | `:186` | "`resolveSessionAuthors` joins (`sessions → memberships`)" | the table is **`public.job_simulation_sessions`** — `sessions` is the dropped name (the exact rot class the brief names) | `how_we_measure.go:1103-1105` |
| m10 | `:218-222` | the M407/M408 email-preview + email-override admin endpoints listed under "**REST/workforce API** (`app/internal/web/backend/api/`)" | they are **not** in `api/server.gen.go`'s route table; they live in `app/internal/web/backend/emailpreview/handler.go` (`New`/`Attach` at `:44`/`:64`) | `grep -n "workforce/ai-readiness" internal/web/backend/api/server.gen.go` → 12 routes, none email |
| m11 | `:531` | "`computeCycleTotals` (`how_we_measure.go:253-261`)" | comment `:256`, func `:260` — range is 3 lines early | `how_we_measure.go:256,260` |

Everything else load-bearing in this file **verified TRUE at HEAD**, including: the whole
`internal/workforce/` → `internal/aireadiness/` rename table (`:25-35`); the D-07 re-anchor
(`m.workforce.LoadMembers(ctx, orgID, "")` in `buildResponseFromSnapshots`, `readiness.go:779`);
`archetypeHighBand = 75` / `archetypeLowCeil = 50` and the None 0-24 / Low 25-50 / Medium 51-74 /
High 75-100 bands (`scoring.go:22-78`); the 30/40/30 split and `(raw/100)×max` tier math
(`readiness.go:29-31, 238-260`); 19-core + 12-enabling = **31** default skills and **3** track-keyed
default sims (`defaults.go:29-79`); the `UNIQUE(org, step_type, track)` index
(`schema/ai_readiness_sim.go:84`); the one-active-cycle partial index and `launched_by`
(`schema/ai_readiness_cycle.go:100,128-129`); **15**-column CSV (`csv.go:13-32`); all 13
`ai_readiness_*` ent schemas incl. `ai_readiness_recommendation`; `interview_aggregated_reports`
columns (`terraform/migrations/20260722081626_jobsim_data_model.sql:47-62`); the `{email, call}`
action-type enum and the `source_id <> target_id` check (`ent/interaction/interaction.go:92-98`,
`schema/interaction.go:48-50`); `organization_settings.go:47`; `queryUserAISkills` reading only
`user_id, skill_id, is_verified` (`readiness.go:87-91`); the full REST surface incl. `/setup` and
`/cycles/{id}/close`; the three deleted orphan components (`AIReadinessContainer/Intro/View` — all
absent) with the five surrounding anchors all still resolving as files; `/enterprise/workforce/
ai-readiness` 404ing (no such dir under `workforce/(new)/`); `WorkforceNewClient.tsx` omitting
readiness (`:125-151`, five tabs, none readiness); `useWorkforceAIReadiness.ts` carrying `tag` and no
`cycle` and never calling `/cycles`; `AIReadinessHero.tsx:88` `if (!air.deadline) return null;`;
`deriveMode` at `useAIReadiness.ts:48-62`; the Analytics-provider double-env PostHog init
(`Analytics.provider.tsx:27-29`); and `app/knowledge/ai-readiness/overview.md`.

### `corpus/architecture/alignment_testing.md`

| # | line | text | truth | proof |
|---|---|---|---|---|
| m12 | `:491-492` | "the source's DNA(s) (the genome — e.g. Clerkenstein ships **three**)" and "the engine's runner(s) (one per surface — `clerkrun`/`jsfapirun`/`expressrun`)" | **five** DNAs and **five** runners — as the same file says at `:370` | `.agentspace/rosetta-extensions/clerkenstein/alignment/dna/` = 5 json; `…/alignment/cmd/` = `clerkrun deployrun expressrun jsfapirun multirun` |
| m13 | `:193` | "`gate.sh:61` calls `alignctl dna coverage --dna … --if-declared`" | it is **`clerkenstein/alignment/scripts/gate.sh:69`** | `grep -n if-declared` → `:65` (comment), `:69` (the call) |
| m14 | `:319` | "Mechanized as `alignment/scripts/{gate,drift-check}.sh`" | those live at **`clerkenstein/alignment/scripts/`**; the reusable `alignment/` section has no `scripts/` dir at all | `ls alignment/` = `README.md cmd examples go.mod internal` |
| m15 | `:250-258` | the table is headed "**The current scores**" and its first row still reads "**97.2% overall** (26/27)" | the note directly beneath it (`:268-273`) records the fix to **100.0% / 100% critical, 27/27** — the table contradicts its own resolution box | internal to the file |
| m16 | `:508`, `:511` | layout block lists `cmd/alignctl run \| capture \| dna list\|diff\|validate` and five `internal/` dirs | omits the `dna coverage` subcommand the same file introduces at `:245`, and `internal/canon` | `ls alignment/internal/` = `canon compare dna outcome report toy` |
| m17 | `:320` | link text `knowledge/alignment.md` targets `../services/clerkenstein.md` | label ≠ target (harmless, but the named file is not what opens) | — |

The one platform-touching claim in this file — `:472`, "the taxonomy surface … formerly the skiller
service's; the domain now lives in `app`'s `public` schema" — is **TRUE** (groundtruth §10, §1).
This file is otherwise entirely about rext tooling and is self-consistent about it.

### `corpus/architecture/service_taxonomy.md`

| # | line | text | truth | proof |
|---|---|---|---|---|
| m18 | `:64`, `:97` | roadrunner "**Orphaned** — nothing calls it" | the (husk) `jobsimulation` compose service still carries `ROADRUNNER_RPC_ADDR=http://roadrunner:10401` — a configured caller, even if it serves no traffic | `platform/docker-compose.yml:118` (inside the `jobsimulation` block, `:83-140`) |
| m19 | `:410` | Summary row "Archived / merged \| Chronos, Intelligence, Skiller …, Skillpath" | omits Jobsimulation, CMS and Roadrunner, which the same file's `:89-97` table lists as merged | internal to the file |
| m20 | `:23-26` | the mermaid puts `Studio-Room` inside the "Studio Services" subgraph as its own box, with `Room --> Desk` | it is embedded in the `app` image and spawned as a subprocess from `app/internal/cms/studio/` — as the same file correctly states at `:162-172` | `app/.github/workflows/build-production.yml:29` (`additional_repo: "anthropos-studio-room:studio"`); `app/internal/cms/studio/` exists |

**Everything else in this file verified TRUE and exact at `2adcf71`** — an unusually clean sweep
result. Checked against the clone: the router banner (`b56d731` 2026-07-31 13:37 + `360efd4` 13:47,
merged `2adcf71` 15:58; `graphql-wundergraph` absent from both `repos.yml` and `docker-compose.yml`;
`terraform/main.tf:20 service_desired_count = 1`; supergraph `backend` alone; `schemas/` =
`backend.graphqls` alone; `subgraphs.conf = BACKEND=v1.360.0`; fold at `915da06`, 2026-07-29 11:24);
`repos.yml:10-13` as the sole migrating repo; `docker-compose.yml:18` `search_path=sentinel` with
`migrations: false`; **every port row** (sentinel 8087, backend 8081-8083 with `PORT=8082`/
`RPC_PORT=8083`/`META_PORT=8084`, jobsim 8400-8401, cms 8090-8091, storage 8300-8301, customerio-sync
8080, messenger 8200-8201, roadrunner 10400-10401, studio-desk 9000/9100, next-web 3000, gotenberg
3200); **every profile row** (sentinel profile-less = always on; `graphql` = backend+jobsim+cms+
storage+roadrunner+gotenberg → six Go services + Gotenberg; the exact per-profile member lists at
`:386-394`); postgres/redis in `common.yml` at 5432/6379 with `bitnamilegacy/redis:latest`; the
husk-container line numbers `:83`/`:144`/`:281`; `045857c`/`fdfa189` (both 2026-04-17); the frontend
endpoint bakes at `docker-compose.yml:352`/`:361` and `:318`/`:334`; `SKILLER_RPC_ADDR=http://
backend:8083` (`:62`,`:121`,`:174`,`:265`); `CMS_RPC_ADDR` on messenger (`:256`);
`DIRECTUS_BASE_ADDR`/`DIRECTUS_PUBLIC_BASE_ADDR` set on `cms` alone (`:164-165`) pointing at
`content.anthropos.work`; no directus service in compose; `JUDGE0_BASE_URL` read by app at
`internal/jobsimwiring/wiring.go:118`; and the shared-library table incl. `taxonomy` = node-id library
only.

### `corpus/services/README.md`

| # | line | text | truth | proof |
|---|---|---|---|---|
| m21 | `:20` | "And **three of the four** (cms, jobsimulation, **roadrunner**) still start CONTAINERS locally" | roadrunner is not one of "the four" — the same banner declares it "**the fifth**" five lines earlier (`:15`). The four are skiller, skillpath, jobsimulation, cms; only **two** of them still start. The substantive claim (those three containers start in the default `graphql` profile) is TRUE | internal to the file; `docker-compose.yml:83,144,281` + `profiles:` at `:140,187,309` |
| m22 | `:15` | "Nothing calls it" (roadrunner) | same as m18 | `docker-compose.yml:118` |

Verified exact: the index enumerates **27** service docs (29 `.md` minus README + TEMPLATE) and the
count matches the directory; `roadrunner/terraform/main.tf:19 = 1`, `cms/terraform/main.tf:39 = 0`,
`jobsimulation/terraform/main.tf:40 = 0` — **all three line numbers and values exact**; the chronos
"NOT archived" correction matches the migration map; every linked target file exists.

### `corpus/services/TEMPLATE.md` · `corpus/services/db-backup.md` · `corpus/services/intelligence.md`

**No findings.** TEMPLATE.md is a pure skeleton with no factual claims. db-backup.md is explicitly
production-only and consistent with the migration map's `db-backup` row (production-only, no compose
service, no `repos.yml` entry at `2adcf71`); nothing in it is a present-tense local-dev instruction.
intelligence.md opens with a standing ⚠ archived banner, cites `fdfa189` correctly, and already
carries the skiller→app note — a correctly-fenced historical doc per the grading rule.

---

## Counts

**2 BLOCKERS · 22 minors.**

Both blockers are in `corpus/services/ai-readiness.md`, and **both are in text the repair pass
touched or left standing in the file it swept** — B2 sits four lines above the M219 correction box
that refutes it, and B1 is the ⚠-flagged premise that the M219 box immediately below it partially
walks back. That matches the addendum's warning exactly: the defects are in the *surrounding prose*,
not the anchors. Every one of the ~40 `file:line` anchors I spot-checked in the swept files
(`service_taxonomy.md` especially) resolved **exact**.

**How the group read:** **mixed, and sharply split by sweep status.** `service_taxonomy.md` (swept)
read as *recently and carefully repaired* — I could not find a single false platform claim in it,
only three cosmetic self-inconsistencies. `ai-readiness.md` (swept) read as the opposite: the newest
text carries both blockers plus 11 minors, and the file now argues with itself in two places about
the PostHog gate and the frozen-read path. The five never-edited files read as **genuinely clean, not
under-detected** — `TEMPLATE`/`db-backup`/`intelligence` have almost no verifiable surface,
`services/README.md` is tightly cited and its three terraform anchors are exact, and
`alignment_testing.md`'s only rot is stale internal counts left behind as the DNA set grew 3 → 5.
