# iter33 — KB-fidelity audit, group 3 (6 files, 1704 lines)

Audited read-only against platform origin HEAD `2adcf71` (`stack-demo/platform`), `app` @ `5ba17044`,
and the `rosetta-extensions` authoring copy @ `b2b46cb` (the subject system for two of the six files).
No file was edited.

## 1. Positive control

| file | `wc -l` | read |
|---|---|---|
| `corpus/architecture/alignment_testing.md` | 516 | read to line 516 (full) |
| `corpus/services/ant-academy.md` | 397 | read to line 397 (full) |
| `corpus/services/clerkenstein.md` | 348 | read to line 348 (full) |
| `corpus/architecture/ai_architecture.md` | 180 | read to line 180 (full) |
| `corpus/architecture/security_compliance.md` | 158 | read to line 158 (full) |
| `corpus/architecture/frontend_architecture.md` | 105 | read to line 105 (full) |

All six read top-to-bottom. None unread.

## 2. Findings

### BLOCKER-1 — `corpus/services/clerkenstein.md:38` + `:47-53` — the headline alignment score is superseded

> `| **Go SDK** (`clerk-2.6.0`, M1) | **97.2% overall · 100% critical** — **26/27 genes**, 14 capabilities | Gate is ≥95 / =100 ⇒ **MET**. The 2.8% is **one deliberately RED gene** (see below). |`
> `**The deliberately RED gene (M218 D16).** `MembershipOrgIdentity/real-org-eid` ships **failing, on purpose**. … Routed forward as `FIX-M219-bapi-org-eid`.`

False at HEAD. The fix landed: `clerkenstein/clerk-backend/store.go:138` (`func (s *Store) SeedOrgIdentity`) and
`:151` (`func (s *Store) LookupOrgEid`) exist, and the DNA itself records it —
`clerkenstein/alignment/dna/clerk-2.6.0.json:131`: *"FIXED in M219 … M219 landed the fix (Store.SeedOrgIdentity /
LookupOrgEid, wired from the roster at cmd/fake-bapi), taking the Go surface 97.2% -> 100%."* The sibling corpus
doc agrees (`alignment_testing.md:268-273`: *"✅ RESOLVED M219 … 97.2% → 100.0% / 100% critical, 27/27, no
divergences"*), so the two docs contradict each other and this one is the wrong side.

**Grade: BLOCKER** — this is the mirror's quoted fidelity number, and the section routes forward a fix that has
already shipped (a reader would re-do `FIX-M219-bapi-org-eid`).

**Correction (one line):** update the Go SDK row to **100% / 100% (27/27)** and convert the RED-gene paragraph to a
past-tense ✅ RESOLVED-M219 note mirroring `alignment_testing.md:268`.

---

### BLOCKER-2 — `corpus/services/clerkenstein.md:42` — wrong `alignctl` exit-code contract

> `| **`@clerk/express`** (`clerk-express-1`, M2c) | **UNMEASURABLE on a box without `@clerk/express` `node_modules`** — the runner cannot build, exits **rc=2, with NO score**. | **Not** a pass. Routed forward as `TEST-M219-expressrun-dep-gate`. |`

False at HEAD. `alignment/cmd/alignctl/run.go:134-135` reads `ExitRegressed = 2` / `ExitUnmeasurable = 3`, and
`run.go:142` prints the `UNMEASURABLE — … THIS IS NOT A PASSING SCORE` banner. The 2/3 split landed at M219 and is
documented correctly in `alignment_testing.md:258` and `:279-283`. The routed-forward item is closed.

**Grade: BLOCKER** — rc=2 now means *a measured regression*. A reader following this doc would read a real
regression as "just the missing Node modules" and dismiss it.

**Correction:** change `rc=2` → `rc=3 (UNMEASURABLE; rc=2 is now REGRESSED)` and mark
`TEST-M219-expressrun-dep-gate` resolved.

---

### BLOCKER-3 — `corpus/services/clerkenstein.md:291-310` — the clerk-js proxy is no longer unbounded or uncached

> `**…and it is UNBOUNDED and UNCACHED — the proxy's real contract (documented in M218; it had never been written down).** `clerk-frontend/server.go:187` fetches the bundle with a bare **`http.Get`**, which is `http.DefaultClient` — i.e. **`Timeout: 0`, no timeout at all**. There is **no server-side cache** …`
> `**Routed forward to M220** (vendor the bundle; serve from disk; keep the CDN proxy only as a *bounded* fallback). Until then, treat a slow/blocked jsdelivr as a **plausible cause of an arbitrarily long demo login**.`

False at HEAD. M220 landed. `clerkenstein/clerk-frontend/server.go:35-67` carries the block header *"M220 (S6/h): the
clerk-js bundle is served FROM DISK; the CDN is a BOUNDED fallback"*, with `clerkJSFetchTimeout = 15 * time.Second`,
`clerkJSMaxBytes = 32 << 20`, `var clerkJSClient = &http.Client{Timeout: clerkJSFetchTimeout}` ("Explicitly NOT
http.DefaultClient"), an on-disk cache at `FAKE_FAPI_CLERKJS_CACHE` (`clerkJSCacheDir()`, `cachePathFor()`, atomic
tmp+rename), and a unit test asserting the package contains no `http.Get(` on the clerk-js path
(`clerk-frontend/clerkjs_cache_test.go`, `clerkjs_proxy_test.go`). Also the cited line is stale: `server.go:187` is
now `mux.HandleFunc("GET /.well-known/jwks.json", s.handleJWKS)`.

**Grade: BLOCKER** — it names a fixed defect as a live "plausible cause of an arbitrarily long demo login", which
sends a latency investigation down a dead path, and re-routes an already-shipped M220 item.

**Correction:** rewrite the block past-tense as ✅ RESOLVED M220 (disk cache + 15 s bounded CDN fallback +
`FAKE_FAPI_CLERKJS_CACHE`), and re-pin the line reference.

---

### BLOCKER-4 — `corpus/architecture/security_compliance.md:65-67` — multi-tenant Layer 1 is over-claimed

> `- Every table has an `organization_id` foreign key`
> `- Ent ORM policies auto-filter all queries by organization`
> `- No cross-tenant data access is possible at the query level`

False at HEAD. In `app/internal/data/ent/schema/` there are **136** entity schemas; only **30** use
`OrganizationMixin{}` (the mixin that actually carries the privacy policy — `mixin.go:126`, `DenyIfNoOrganizationInContext`
+ `FilterOrganizationEdges`), and **7** use `OrganizationIDMixin{}`, which is explicitly *"a plain nullable
organization_id column + index"* with **no** policy (`skill.go:18-19`). **78** schemas never mention `organization`
at all. The platform says so in its own comments:
- `job_simulation_session.go:5` — *"L2: NO Ent privacy Policy; owner/org/tenant are plain fields."* (the central
  jobsim runtime entity, and the whole 23-table run-state fan-out around it)
- `jobrole.go:18`, `category.go:15`, `job_role_embeddings.go:16`, `jobrolecategory.go:14` — *"no
  UserMixin/OrganizationMixin and no Policy(), so the taxonomy stays globally readable"*

**Grade: BLOCKER** — an engineer writing a query against `job_simulation_sessions` (or any of the ~99 policy-less
schemas) while trusting "Ent auto-filters by organization" ships a cross-tenant read. This is exactly the class of
claim that misdirects real work.

**Correction:** state that org isolation is **per-schema and opt-in** — `OrganizationMixin` (30 schemas) carries the
Ent privacy policy; `OrganizationIDMixin` is a bare column with no policy; taxonomy/public tables are deliberately
un-scoped; and jobsim run-state is filtered in application code, not by Ent.

---

### minor-1 — `corpus/services/clerkenstein.md:18` — the monorepo section list is 6 of 11

> `**One monorepo, two clone roles.** `rosetta-extensions` is ONE private monorepo with sections (`clerkenstein`, `demo-stack`, `stack-injection`, `stack-core`, `stack-seeding`, `alignment`).`

At HEAD the repo has 11 sections: the six listed plus `dev-stack`, `stack-secrets`, `stack-snapshot`,
`stack-verify`, `playthroughs`. Reads as a complete enumeration.
**Correction:** add the five missing sections (or say "sections include …").

### minor-2 — `corpus/services/clerkenstein.md:3` — stale "Last updated"

> `**Status:** v0.3 (… v2.3 "cue to cue" M217 … + **M218 the roster-aware fake BAPI**) · **Last updated:** 2026-07-14`

The body carries v2.8 M256 (`:148`, `:170`) and M257x iter-23 (`:260`) content dated 2026-08-01.
**Correction:** bump the status line to v2.8 / 2026-08-01.

### minor-3 — `corpus/architecture/alignment_testing.md:491` — "Clerkenstein ships three" DNAs

> `| | the source's DNA(s) (the genome — e.g. Clerkenstein ships three) |`

Five: `clerkenstein/alignment/dna/{clerk-2.6.0,clerk-js-5,clerk-multi-1,clerk-express-1,clerk-deploy-1}.json`. The
same doc says **five** at `:370`.
**Correction:** "ships five".

### minor-4 — `corpus/architecture/alignment_testing.md:492` — runner list is 3 of 5

> `| | the alignment tests + goldens + the engine's runner(s) (one per surface — `clerkrun`/`jsfapirun`/`expressrun`) |`

`clerkenstein/alignment/cmd/` holds `clerkrun deployrun expressrun jsfapirun multirun`.
**Correction:** add `deployrun`/`multirun` (or defer to `:370`, which lists all five).

### minor-5 — `corpus/architecture/alignment_testing.md:508` — Layout block omits `dna coverage`

> `  cmd/alignctl            run | capture | dna list|diff|validate`

`alignment/cmd/alignctl/dna.go:26` has `case "coverage":`, and the doc documents it at `:245`. The Layout block also
omits `internal/canon` (which exists beside `compare`/`dna`/`outcome`/`report`).
**Correction:** `run | capture | dna list|diff|validate|coverage`.

### minor-6 — `corpus/architecture/alignment_testing.md:193` — wrong line number for the gate's coverage call

> `⚠ **Not the same as the bare command — do not conflate them.** `gate.sh:61` calls `alignctl dna coverage --dna … --if-declared`.`

At HEAD it is `clerkenstein/alignment/scripts/gate.sh:69`. The claim itself is true and verified.
**Correction:** re-pin to `gate.sh:69`.

### minor-7 — `corpus/architecture/alignment_testing.md:254` — the "current scores" table still leads with 97.2%

> `| Go SDK | `clerk-2.6.0` | **97.2% overall · 100% critical** (26/27 genes) | gate ≥95 / =100 ⇒ **MET** |`

Superseded by the ✅ RESOLVED-M219 note 14 lines below (`:268-273` — 27/27, 100%/100%), so the page self-corrects;
but the table is under the heading *"The current scores"* and is what gets quoted. (Not a BLOCKER: the correction is
present and explicit on the same page.)
**Correction:** put 100% / 27/27 in the table and keep the M218→M219 history in the note.

### minor-8 — `corpus/architecture/alignment_testing.md:3` — stale "Last updated"

> `**Status:** canonical · **Last updated:** 2026-06-06`

The body contains the M218 rewrite and the M219 resolutions (July 2026).
**Correction:** bump to the M219 date.

### minor-9 — `corpus/architecture/alignment_testing.md:320` — mislabeled link

> `the bump runbook + exit-code contract are in the repo's own [`knowledge/alignment.md`](../services/clerkenstein.md)`

The link text names a rext file (`clerkenstein/knowledge/alignment.md`, which does exist) but the href resolves to
the rosetta service doc, which is a different document; the very next line links the same target again as
"[Clerkenstein]".
**Correction:** drop the href (it is a repo-internal path, not a corpus link) or point the label at the pointer doc.

### minor-10 — `corpus/architecture/frontend_architecture.md:39` — "~15 sites" undercounts

> `**but there are direct REST/SSE calls**, ~15 sites hitting `NEXT_PUBLIC_BACKEND_API_URL``

At HEAD, 23 code files under `apps/` + `packages/` reference it (35 occurrences; plus 2 `.env.example` and 1 `.md`).
The listed exemplars are all real, but Talk-to-Data, credits, coursebuilder, AI-readiness and workforce are missing
from the list. The claim the number supports ("*GraphQL only* is the wrong mental model") is correct and, if
anything, understated.
**Correction:** "~25 sites" and add `packages/core-js/src/{talkToData,credits,coursebuilder,workforce}/api.ts`.

### minor-11 — `corpus/architecture/ai_architecture.md:141` — the recording-storage sentence omits the Bunny leg

> `Both recordings are stored in S3 and linked to the simulation session.`

True as far as capture goes (`app/internal/jobsimulation/recording/chime.go:36` builds
`arn:aws:s3:::<bucket>` as the media-pipeline sink), but the **playable** artifact is a Bunny.net CDN reference —
`app/internal/data/ent/schema/chime_recording.go:25` `field.String("bunny_video_id")` +
`app/internal/jobsimulation/bunny/bunny.go`, wired into `simulator/manager/manager.go:247`. The corpus's own
`ops/demo/media-substrate-spec.md` calls that leg load-bearing ("the bytes live on Bunny's CDN"). A reader looking
for the video bytes from this page goes to S3 only.
**Correction:** add "…and the playable MP4 is published to a Bunny.net Stream pull-zone (`chime_recordings.bunny_video_id`)".

## 3. Per-file verdicts

- **`corpus/services/clerkenstein.md`** — 3 BLOCKERs (score table, express exit code, the clerk-js proxy block) +
  2 minors. Everything else verified true at HEAD: the `v0.34.3` colony pin vs `app`'s `v0.35.2`
  (`clerkenstein/go.mod:8` vs `app/go.mod:16`) and `clerk-sdk-go/v2 v2.7.0` (`app/go.mod:31`), `sentinel`/`storage`
  still on `v0.34.3`; the BAPI≠FAPI section (`GET /v1/me/organization_memberships` registered at
  `clerk-frontend/server.go:186`, `meorgmemberships_test.go` present, still **no** gene for it in `clerk-js-5.json`);
  the single-tenant seat limitation (`registry.go:24` `activeKey`, `server.go:125` `clientID:"client_clerkenstein"`,
  `:665` `sessID = "sess_clerkenstein"`, **zero** `r.Cookie(` calls anywhere in `clerkenstein/`); the sticky
  `signedOut` flag (`server.go:103-107`, `:380`, `:559`) and all five named pins in `server_test.go` (`:256`, `:286`,
  `:390`, `:427`, `:461`).
- **`corpus/architecture/alignment_testing.md`** — 0 BLOCKERs, 6 minors. The load-bearing M218 corrections are all
  true at HEAD: `dna coverage` exists and `gate.sh` passes `--if-declared`; **only** `clerk-2.6.0.json` declares a
  `consumed_surface` (verified: 1 vs 0/0/0/0); the exit-code split is real (`run.go:134-135`); the inert
  sub-directory workflow exists at `clerkenstein/.github/workflows/alignment.yml` with no repo-root `.github/`;
  `datadna` is at `stack-seeding/cmd/datadna` with `stack-seeding/dna/data-dna.json`; module is
  `anthropos.dev/alignment`; the toy runner is `alignment/examples/toy/cmd/toyrun`.
- **`corpus/architecture/security_compliance.md`** — 1 BLOCKER (Layer 1). **Otherwise clean.** The prod-only content
  (Cosmo Router in the public subnet, `db-backup`, DPA/sub-processors) is legitimately production-scoped and is not
  a finding under the grading rule — ground truth #4 fences the router as still-declared in prod terraform. Layer 2
  verified (`app/internal/authorization/authorization.go:7-8` calls sentinel over proto); the `aiusage` attribution
  at `:115` is correct (`app/internal/aiusage/ai_usage.go`).
- **`corpus/architecture/ai_architecture.md`** — 0 BLOCKERs, 1 minor (recording storage). **Otherwise clean, and
  notably well-pinned:** the `@ 5ba17044` citation is `app`'s exact HEAD; `app/internal/ai/ai.go` really does not
  exist; `internal/jobsimulation/ai/ai.go` and `internal/skillerai/ai.go` do, with `flag_use_azure_us` at the cited
  regions (~:267/:344 and ~:347); `livekitgptrealtime` (`internal/cms/directus/collections/jobsimulation.go:1084`),
  ElevenLabs still live (`internal/jobsimulation/calls/elevenlabs.go`), `flag_use_realtime_openai`
  (`internal/jobsimulation/calls/livekit.go:133`), `anthropic-45-sonnet-aws`/`gpt-5-mini`
  (`.../jobsimulation.go:983,988`), the studio-room-as-subprocess framing, and the embeddings tables all check out.
- **`corpus/architecture/frontend_architecture.md`** — 0 BLOCKERs, 1 minor (the "~15" count). **Clean on ground
  truth #4** — it is the *most* swept file in this group: it states the endpoint is `backend` at `:8082/graphql/query`
  locally since `2adcf71` with the router qualified as prod-only (`:39`, `:44`), the supergraph as the single
  `backend` subgraph (`:32`), and `repos.yml` @ `2adcf71` as exactly 9 entries (verified: app, cms, jobsimulation,
  sentinel, storage, messenger, roadrunner, next-web-app, studio-desk). Apps/ports/versions all verified against the
  clone (`next ^16.2.7`, `pnpm@10.30.3`, `node >=24.0.0`, `!apps/mobile`, Expo `--port 3031`, 8 locales in
  `configs/i18n`).
- **`corpus/services/ant-academy.md`** — **0 findings. Clean.** Checked specifically for the ground-truth-#4 trap:
  the doc never routes the academy through Cosmo/the router — it names
  `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` as the *platform academy subgraph* endpoint served by `app internal/academy`
  (`:100`, `:125`, `:372-375`) and states the supergraph is `backend` alone (`:100`), which is exactly right at
  `2adcf71`. Verified against the clones: `academyCatalogSeries` / `academyCatalogSkillPaths` /
  `upsertChapterProgress` exist in `app/internal/web/backend/graphql/graph/schemas/academy.graphqls:664,672,694`;
  the `academy_*` Ent tables and `app/cmd/academy-seed` exist; ant-academy's own repo confirms port 3077, `node >=22`,
  `/courses` under `(authed)` with **no** `/library/[slug]`, the vendored `code/public/assets/fontawesome/{css,webfonts}`,
  and every named source file (`src/i18n/{locale.js,LocaleSwitch.jsx,translate.js}`,
  `src/lib/{serverChapterBody,backendContent,serverTenant,draftMode,draftCatalog}.js`, `src/graphql/server.js`,
  `proxy.js`, `vercel.json`, `app/not-found.jsx:43` — *"You wandered off the trail."*). The `code/.env` vs
  `code/.env.local` wording differs from the root `CLAUDE.md`, but Next.js loads both, so it is not false.

## 4. Totals

- **BLOCKERs: 4** — 3 in `clerkenstein.md` (`:38`+`:47-53`, `:42`, `:291-310`), 1 in `security_compliance.md` (`:65-67`).
- **minors: 11** — clerkenstein.md ×2, alignment_testing.md ×6, frontend_architecture.md ×1, ai_architecture.md ×1
  (plus minor-3 counted once).
- **Files not fully read: none.** All 6 read to their last line.

**Shape note.** Three of the four blockers are the archetype the ground truth predicted: none of them uses a word a
merged-service sweep would grep for. They are *self-consistent prose describing a system state that a later
milestone in the same programme already changed* — a score, an exit code, and a performance defect, each still
carrying a "routed forward to M219/M220" tail for work that has shipped. The fourth (security Layer 1) is the other
predicted shape: a confident absolute ("Every table…", "No cross-tenant access is possible") that was never true of
the taxonomy tables and stopped being true of the jobsim fan-out when it landed in `public`.
