# TIER-1 ADJUDICATION BATCH 04 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 04-001
- **id**: `B04-001`
- **corpus site**: `corpus/services/gotenberg.md:47-47` (bullet)
- **citation**: `internal/web/backend/coursebuilder/extract.go:77`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/coursebuilder/extract.go`  (86 lines)

**CLAIMING UNIT**

```md
* **Call sites** (two, `app` @ `9d00a313`): `internal/web/backend/coursebuilder/extract.go:77` — course-builder upload text extraction, for the nine MIME types `docconv` can't read directly (`extract.go:17-27` — `.xls`, `.xlsx`, `.ppt`, `.doc`, the three OpenDocument types, and RTF under both its MIME spellings; **DOCX is not among them**, it goes straight to `docconv`); and `internal/worker/tasks/user_import_resume_2d.go:68` — the résumé-import **OCR fallback**, DOCX only, reached only when the document has no readable text
```

**CITED CONTENT**

```
    74  			if err != nil {
    75  				return "", err
    76  			}
    77  			pdf, err := converter.ConvertToPDF(context.Background(), gotenbergURL, data, name)
    78  			if err != nil {
    79  				return "", fmt.Errorf("convert %s via gotenberg: %w", mimeType, err)
    80  			}
```

## 04-002
- **id**: `B04-002`
- **corpus site**: `corpus/services/gotenberg.md:47-47` (bullet)
- **citation**: `internal/worker/tasks/user_import_resume_2d.go:68`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/worker/tasks/user_import_resume_2d.go`  (155 lines)

**CLAIMING UNIT**

```md
* **Call sites** (two, `app` @ `9d00a313`): `internal/web/backend/coursebuilder/extract.go:77` — course-builder upload text extraction, for the nine MIME types `docconv` can't read directly (`extract.go:17-27` — `.xls`, `.xlsx`, `.ppt`, `.doc`, the three OpenDocument types, and RTF under both its MIME spellings; **DOCX is not among them**, it goes straight to `docconv`); and `internal/worker/tasks/user_import_resume_2d.go:68` — the résumé-import **OCR fallback**, DOCX only, reached only when the document has no readable text
```

**CITED CONTENT**

```
    65  			h.logger.Info("document has no readable text, attempting OCR fallback")
    66  			ocrInput := fileBytes
    67  			if mimetype == "application/vnd.openxmlformats-officedocument.wordprocessingml.document" && h.cfg.GotenbergURL != "" {
    68  				pdfBytes, convErr := converter.ConvertToPDF(ctx, h.cfg.GotenbergURL, fileBytes, "resume.docx")
    69  				if convErr != nil {
    70  					h.logger.Error("Gotenberg DOCX-to-PDF conversion failed", "error", convErr)
    71  				} else {
```

## 04-003
- **id**: `B04-003`
- **corpus site**: `corpus/services/graphql-wundergraph.md:3-28` (paragraph)
- **citation**: `graphql-wundergraph/terraform/main.tf:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/terraform/main.tf`  (63 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ THE ROUTER IS GONE FROM LOCAL DEV — and its two states differ
>
> | | production | a fresh local stack @ platform origin HEAD |
> |---|---|---|
> | the router | **still declared** — `graphql-wundergraph/terraform/main.tf:20` `service_desired_count = 1` | **deleted** — no `graphql` compose service, no `repos.yml` entry |
> | the repo | **ARCHIVED on GitHub, 2026-07-30** (read-only) | not cloned by `make init` |
> | what a frontend talks to | the router | **`backend` directly**, `http://localhost:8082/graphql/query` |
>
> Platform `b56d731` + `360efd4` (merged **`2adcf71`**, 2026-07-31) dropped the router from
> `docker-compose.yml` **and** `repos.yml` and re-pointed local dev at `backend`. **There is no `:5050` on a
> local stack.** **The `graphql` profile is gone too:** `0dab54d` (*"rename graphql -> core"*) renamed it,
> so `Makefile:10` reads `PROFILE ?= core` and the token appears in **no `profiles:` key at all** — the
> **five** that exist are `core`, `backend`, `all`, `studio-desk`, `frontend`. (This list read *"the
> eight"* until platform `838d907`, 2026-08-05: `storage-legacy`, `customerio-sync` and `messenger`
> were deleted along with the three services that declared them, so those three tokens are now
> retired exactly as `graphql` is.) Asking for `graphql` therefore **exits 0** and starts only the
> always-on floor (`postgresql`, `redis`, `sentinel`), which is worse than an error.
>
> The supergraph is **ONE** subgraph: `915da06` (2026-07-29) folded the cms subgraph into `backend`
> (cms-in-app v8.0) and deleted the `jobsimulation` entry in the **same commit** — a **3 → 1** step,
> not 2 → 1. `supergraph-config-prod.yaml` lists `backend` alone and `schemas/` holds
> `backend.graphqls` alone.
>
> Everything below the fold describes the gateway **as it still exists in production and in the archived
> repo**. Read [`../architecture/platform-migration-status.md`](../architecture/platform-migration-status.md)
> — the fenced map — before acting on any local-development instruction on this page.
```

**CITED CONTENT**

```
    17    tags                           = var.tags
    18    aws_region                     = var.aws_region
    19    project                        = local.project
    20    service_desired_count          = 1
    21    service_cpu                    = local.service_cpu
    22    service_memory                 = local.service_memory
    23    service_port                   = local.port
```

## 04-004
- **id**: `B04-004`
- **corpus site**: `corpus/services/graphql-wundergraph.md:172-178` (bullet)
- **citation**: `docker-compose.yml:135`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **Upstream consumers**: every GraphQL client — `next-web-app`, `studio-desk`, mobile. **In production** they
  hit the router; **locally they hit `backend` directly** at `:8082/graphql/query`
  (`docker-compose.yml:135` studio-desk's `VITE_GRAPHQL_ENDPOINT`, `:160` next-web-app's
  `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` — each also baked as a build arg, at `:119` and `:151`), because
  the router service no longer exists in compose. Those line numbers move on every compose
  clean-up — they were `:220`/`:236` at `0dab54d` and `:334`/`:352` at `2adcf71` — so grade the
  construct, not the offset.
```

**CITED CONTENT**

```
   132        - NODE_ENV=development
   133        - PORT=9000
   134        - VITE_ENVIRONMENT=production
   135        - VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query
   136      networks:
   137        - app-network
   138      depends_on:
```

## 04-005
- **id**: `B04-005`
- **corpus site**: `corpus/services/graphql-wundergraph.md:181-181` (bullet)
- **citation**: `ci/update-subgraph.sh:9`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/ci/update-subgraph.sh`  (10 lines)

**CLAIMING UNIT**

```md
* **CI/prod**: GitHub Releases on **`anthropos-work/app` only** (schema artifacts) + `anthropos-work/infrastructure` Terraform + `release-service.yml`. `ci/update-subgraph.sh:9` carries **exactly one** `gh release download`, `-R anthropos-work/app`; the `jobsimulation` and `cms` downloads were **deleted at `915da06`** when those subgraphs folded into `backend`. (This bullet claimed all three until M257x iter-49 — the two bullets above it already carried their historical fence; this one did not.)
```

**CITED CONTENT**

```
     6  
     7  rm -rf schemas && mkdir schemas
     8  # cms-in-app (v8.0): cms folded into backend — the backend subgraph SDL now carries the cms types.
     9  gh release download ${BACKEND?}   --pattern '*.graphqls' -R anthropos-work/app       -O schemas/backend.graphqls
    10  
```

## 04-006
- **id**: `B04-006`
- **corpus site**: `corpus/services/graphql-wundergraph.md:207-215` (paragraph)
- **citation**: `app/internal/authorization/gqlauthz/gqlauthz.go:186`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/authorization/gqlauthz/gqlauthz.go`  (263 lines)

**CLAIMING UNIT**

```md
**That rejection *is* the healthy response — do not read it as a failure.** The endpoint is
default-closed: `app`'s GraphQL authorization middleware rejects any anonymous operation that isn't
`@public`-annotated, a federation query, or (in local dev only) pure schema introspection
(`app/internal/authorization/gqlauthz/gqlauthz.go:186`). `__typename` is **deliberately excluded**
from the introspection exemption in every environment — the app pins that with its own regression
tests, which drive `__typename` specifically (`gqlauthz_test.go`, `"bare __typename is not exempt"`
and `TestAnonymousRejectionLogsAtWarn`). This is stock behaviour, not something a demo patch does.
Note also that the transport is healthy at **HTTP 200**: GraphQL reports the refusal in the `errors`
array, not in the status code.
```

**CITED CONTENT**

```
   183  			// once a minute (GlitchTip #92), so Error here floods the issue
   184  			// tracker with expected traffic.
   185  			l.With("operation", opCtx.OperationName).Warn("viewer is nil")
   186  			return errorResponse(fmt.Errorf("unknown viewer: %w", errForbidden))
   187  		}
   188  		l = l.With("viewer", viewer.ID())
   189  		org := viewer.GetOrganization()
```

## 04-007
- **id**: `B04-007`
- **corpus site**: `corpus/services/graphql-wundergraph.md:274-274` (bullet)
- **citation**: `CLAUDE.md:85`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/CLAUDE.md`  (581 lines)

**CLAIMING UNIT**

```md
* ⚠️ **CORRECTED M257x iter-115 — this sentence had a compound subject and only ONE half survives.** The repo's `CLAUDE.md` *"Version Tracking"* section is **NOT stale**: at `graphql-wundergraph` `60c229f3` it reads *"Service versions are tracked in `subgraphs.conf`. There is exactly **one** pin now: `BACKEND=v1.360.0`. The `CMS`, `JOBSIMULATION`, `SKILLER` and `SKILLPATH` entries were removed as each of those services merged into `app`"* — and `subgraphs.conf` at that same ref is the single line `BACKEND=v1.360.0`, a byte match. It **is** the current form of the very claim this bullet offered as its correction, and `git log -- CLAUDE.md` shows the checkout itself last rewrote it; a reader was being told to distrust an accurate section of the ground-truth repo. **What IS stale is the `-local.yaml` reference** — `CLAUDE.md:85` still says `wgc router compose -i supergraph-config-local.yaml`, while `ls supergraph-config-*.yaml` @ `60c229f3` returns `-compose`, `-dev`, `-prod` and no `-local`. So: `subgraphs.conf` is the version source of truth (as that CLAUDE.md already says), and the config variants are `compose`/`dev`/`prod`.
```

**CITED CONTENT**

```
    82  /stack-update           # the main dev stack
    83  /stack-update dev-2     # a named additional dev stack
    84  ```
    85  
    86  This skill (← the former `update-platform`) executes `corpus/ops/update_guide.md` with:
    87  - Daily/weekly/full update scenarios
    88  - Git conflict handling
```

## 04-008
- **id**: `B04-008`
- **corpus site**: `corpus/services/hiring.md:17-61` (paragraph)
- **citation**: `internal/organization/manager.go:450`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/organization/manager.go`  (1598 lines)

**CLAIMING UNIT**

```md
> **⚠️ RE-GROUNDED — v2.8 M257x iter-23, against platform origin `2adcf71` / `app` @ `5ba17044`.**
> *(The platform-side citations below were re-anchored again to platform `0dab54d` in the M257x sweep —
> `d11a403` had removed the cms / jobsimulation / roadrunner compose services and `repos.yml` entries.)*
> **`5ba17044` is the historical iter-23 re-grounding ref — NOT a governing pin over the anchors below.**
> The `app`-side anchors have been re-derived repeatedly since (M257x iter-49, -52, -98, -100, -102); every
> one re-derived at **iter-102** is measured at `app` **`ad9f3c49`** (= `origin/main` **and** the demo build
> pin `stack-demo/clones.pin.json`, 2026-08-06). They are **not** interchangeable:
> `internal/organization/manager.go:450` / `:453` / `:537` resolve at `ad9f3c49`, and **the offset back to
> `5ba17044` is NOT uniform — do not apply one delta to all three.** Measured line-for-line:
> `:450 → :448` (−2), `:453 → :451` (−2), but **`:537 → :487` (−50)**. This banner said *"each off by −2"*
> until M257x iter-108; on that rule `:537` would resolve to `:535` at `5ba17044`, which is unrelated code. **Read the ref that travels with the anchor; do not read this banner as a pin over the
> whole document.**
> **This doc named a table the platform has since DROPPED — which is the worst possible version of the warning
> directly above.** The score source was `public.local_jobsimulation_sessions`, a `Float32` MIRROR. `app`
> migration `20260729133514.sql:58-62` — *"5. Drop the mirrors."* — **re-points the *referencing* rows** (the
> assignment-session link ids) and then `DROP TABLE "local_jobsimulation_sessions"`. **There is no back-fill:**
> `SET "score"` has **0 hits across the entire migration set**, so no score was copied from the mirror to the
> canonical row — the canonical row already carried it. (This paragraph said "back-fills" until M257x iter-49.)
> Everything below is re-pointed; the three facts that changed:
> 1. **Score source → `public.job_simulation_sessions.score`**, read by
>    `app/internal/organization/intelligence.go:1700` (`m.ent.JobSimulationSession.Query()`). There is no
>    mirror/canonical **pair** any more, so the write-set is **one** session row, not two.
> 2. **Everything `app` writes is in `public`** — which is *not* the same claim as "the `jobsimulation` schema
>    is gone." `20260722104506.sql:79` is `DROP TABLE "sess
```

**CITED CONTENT**

```
   447  	logger := m.logger.With("user_id", user.ID, "organization_id", org.ID)
   448  
   449  	antRole := enum.RoleMember
   450  	switch org.IsHiring {
   451  	case true:
   452  		// if the organization is hiring, add the user as a candidate
   453  		antRole = enum.RoleCandidate
```

## 04-009
- **id**: `B04-009`
- **corpus site**: `corpus/services/hiring.md:17-61` (paragraph)
- **citation**: `app/internal/organization/intelligence.go:1700`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/organization/intelligence.go`  (2296 lines)

**CLAIMING UNIT**

```md
> **⚠️ RE-GROUNDED — v2.8 M257x iter-23, against platform origin `2adcf71` / `app` @ `5ba17044`.**
> *(The platform-side citations below were re-anchored again to platform `0dab54d` in the M257x sweep —
> `d11a403` had removed the cms / jobsimulation / roadrunner compose services and `repos.yml` entries.)*
> **`5ba17044` is the historical iter-23 re-grounding ref — NOT a governing pin over the anchors below.**
> The `app`-side anchors have been re-derived repeatedly since (M257x iter-49, -52, -98, -100, -102); every
> one re-derived at **iter-102** is measured at `app` **`ad9f3c49`** (= `origin/main` **and** the demo build
> pin `stack-demo/clones.pin.json`, 2026-08-06). They are **not** interchangeable:
> `internal/organization/manager.go:450` / `:453` / `:537` resolve at `ad9f3c49`, and **the offset back to
> `5ba17044` is NOT uniform — do not apply one delta to all three.** Measured line-for-line:
> `:450 → :448` (−2), `:453 → :451` (−2), but **`:537 → :487` (−50)**. This banner said *"each off by −2"*
> until M257x iter-108; on that rule `:537` would resolve to `:535` at `5ba17044`, which is unrelated code. **Read the ref that travels with the anchor; do not read this banner as a pin over the
> whole document.**
> **This doc named a table the platform has since DROPPED — which is the worst possible version of the warning
> directly above.** The score source was `public.local_jobsimulation_sessions`, a `Float32` MIRROR. `app`
> migration `20260729133514.sql:58-62` — *"5. Drop the mirrors."* — **re-points the *referencing* rows** (the
> assignment-session link ids) and then `DROP TABLE "local_jobsimulation_sessions"`. **There is no back-fill:**
> `SET "score"` has **0 hits across the entire migration set**, so no score was copied from the mirror to the
> canonical row — the canonical row already carried it. (This paragraph said "back-fills" until M257x iter-49.)
> Everything below is re-pointed; the three facts that changed:
> 1. **Score source → `public.job_simulation_sessions.score`**, read by
>    `app/internal/organization/intelligence.go:1700` (`m.ent.JobSimulationSession.Query()`). There is no
>    mirror/canonical **pair** any more, so the write-set is **one** session row, not two.
> 2. **Everything `app` writes is in `public`** — which is *not* the same claim as "the `jobsimulation` schema
>    is gone." `20260722104506.sql:79` is `DROP TABLE "sess
```

**CITED CONTENT**

```
  1697  	ctx = authorization.NewContextWithTargets(ctx, userIds)
  1698  	ctx = authorization.NewContextWithDecision(ctx, authorization.Allow)
  1699  
  1700  	query := m.ent.JobSimulationSession.Query().
  1701  		Where(
  1702  			jobsimulationsession.SimID(jobSimulationId),
  1703  			jobsimulationsession.OwnerIDIn(userIds...),
```

## 04-010
- **id**: `B04-010`
- **corpus site**: `corpus/services/hiring.md:17-61` (paragraph)
- **citation**: `repos.yml:14-17`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
> **⚠️ RE-GROUNDED — v2.8 M257x iter-23, against platform origin `2adcf71` / `app` @ `5ba17044`.**
> *(The platform-side citations below were re-anchored again to platform `0dab54d` in the M257x sweep —
> `d11a403` had removed the cms / jobsimulation / roadrunner compose services and `repos.yml` entries.)*
> **`5ba17044` is the historical iter-23 re-grounding ref — NOT a governing pin over the anchors below.**
> The `app`-side anchors have been re-derived repeatedly since (M257x iter-49, -52, -98, -100, -102); every
> one re-derived at **iter-102** is measured at `app` **`ad9f3c49`** (= `origin/main` **and** the demo build
> pin `stack-demo/clones.pin.json`, 2026-08-06). They are **not** interchangeable:
> `internal/organization/manager.go:450` / `:453` / `:537` resolve at `ad9f3c49`, and **the offset back to
> `5ba17044` is NOT uniform — do not apply one delta to all three.** Measured line-for-line:
> `:450 → :448` (−2), `:453 → :451` (−2), but **`:537 → :487` (−50)**. This banner said *"each off by −2"*
> until M257x iter-108; on that rule `:537` would resolve to `:535` at `5ba17044`, which is unrelated code. **Read the ref that travels with the anchor; do not read this banner as a pin over the
> whole document.**
> **This doc named a table the platform has since DROPPED — which is the worst possible version of the warning
> directly above.** The score source was `public.local_jobsimulation_sessions`, a `Float32` MIRROR. `app`
> migration `20260729133514.sql:58-62` — *"5. Drop the mirrors."* — **re-points the *referencing* rows** (the
> assignment-session link ids) and then `DROP TABLE "local_jobsimulation_sessions"`. **There is no back-fill:**
> `SET "score"` has **0 hits across the entire migration set**, so no score was copied from the mirror to the
> canonical row — the canonical row already carried it. (This paragraph said "back-fills" until M257x iter-49.)
> Everything below is re-pointed; the three facts that changed:
> 1. **Score source → `public.job_simulation_sessions.score`**, read by
>    `app/internal/organization/intelligence.go:1700` (`m.ent.JobSimulationSession.Query()`). There is no
>    mirror/canonical **pair** any more, so the write-set is **one** session row, not two.
> 2. **Everything `app` writes is in `public`** — which is *not* the same claim as "the `jobsimulation` schema
>    is gone." `20260722104506.sql:79` is `DROP TABLE "sess
```

**CITED CONTENT**

```
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
    14    - name: app
    15      type: go
    16      migrations: true
    17      schema: public
    18    - name: sentinel
    19      type: go
    20      migrations: false
```

## 04-011
- **id**: `B04-011`
- **corpus site**: `corpus/services/hiring.md:17-61` (paragraph)
- **citation**: `dependency_map.md:78`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/architecture/dependency_map.md`  (104 lines)

**CLAIMING UNIT**

```md
> **⚠️ RE-GROUNDED — v2.8 M257x iter-23, against platform origin `2adcf71` / `app` @ `5ba17044`.**
> *(The platform-side citations below were re-anchored again to platform `0dab54d` in the M257x sweep —
> `d11a403` had removed the cms / jobsimulation / roadrunner compose services and `repos.yml` entries.)*
> **`5ba17044` is the historical iter-23 re-grounding ref — NOT a governing pin over the anchors below.**
> The `app`-side anchors have been re-derived repeatedly since (M257x iter-49, -52, -98, -100, -102); every
> one re-derived at **iter-102** is measured at `app` **`ad9f3c49`** (= `origin/main` **and** the demo build
> pin `stack-demo/clones.pin.json`, 2026-08-06). They are **not** interchangeable:
> `internal/organization/manager.go:450` / `:453` / `:537` resolve at `ad9f3c49`, and **the offset back to
> `5ba17044` is NOT uniform — do not apply one delta to all three.** Measured line-for-line:
> `:450 → :448` (−2), `:453 → :451` (−2), but **`:537 → :487` (−50)**. This banner said *"each off by −2"*
> until M257x iter-108; on that rule `:537` would resolve to `:535` at `5ba17044`, which is unrelated code. **Read the ref that travels with the anchor; do not read this banner as a pin over the
> whole document.**
> **This doc named a table the platform has since DROPPED — which is the worst possible version of the warning
> directly above.** The score source was `public.local_jobsimulation_sessions`, a `Float32` MIRROR. `app`
> migration `20260729133514.sql:58-62` — *"5. Drop the mirrors."* — **re-points the *referencing* rows** (the
> assignment-session link ids) and then `DROP TABLE "local_jobsimulation_sessions"`. **There is no back-fill:**
> `SET "score"` has **0 hits across the entire migration set**, so no score was copied from the mirror to the
> canonical row — the canonical row already carried it. (This paragraph said "back-fills" until M257x iter-49.)
> Everything below is re-pointed; the three facts that changed:
> 1. **Score source → `public.job_simulation_sessions.score`**, read by
>    `app/internal/organization/intelligence.go:1700` (`m.ent.JobSimulationSession.Query()`). There is no
>    mirror/canonical **pair** any more, so the write-set is **one** session row, not two.
> 2. **Everything `app` writes is in `public`** — which is *not* the same claim as "the `jobsimulation` schema
>    is gone." `20260722104506.sql:79` is `DROP TABLE "sess
```

**CITED CONTENT**

```
    75  ### 2. Job Simulation
    76  `Frontend` -> `Backend` (`app` — the **jobsimulation domain**, in-process; there is no jobsimulation service to reach)
    77  *   The jobsimulation engine fetches the simulation **definition** (the `simulations` content/blueprint) from the cms domain by ID — in-process since cms-in-app. It owns no content, only the run/session state.
    78  *   The jobsimulation engine stores its session/run **state** (interactions, recordings, validation results, anti-cheat) via the **storage domain inside `app`** (in-process since v9.0 — not the standalone `storage` service, which `838d907` deleted from compose altogether) or directly to the **`public`** schema (the legacy `jobsimulation` schema is non-authoritative).
    79  *   Voice flows go through LiveKit; video recordings via AWS Chime SDK.
    80  
    81  ### 3. Content Delivery
```

## 04-012
- **id**: `B04-012`
- **corpus site**: `corpus/services/hiring.md:89-106` (bullet)
- **citation**: `hiring_config.go:99`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-seeding/seeders/hiring_config.go`  (130 lines)

**CLAIMING UNIT**

```md
1. **Backend — `public.organizations.is_hiring boolean NOT NULL default false`.** The server-side org-type. The
   seeder writes it directly (M222 landed the gate — see § *The seeder-output contract*).
   **The `resolver_queries.go` insights path does NOT read it at all** — `InsightsJobSimulationByMemberships`
   (`:1034-1080`) gates on exactly two things, the `OrgFeatureInsights` Casbin permission (`:1035`) and
   membership status ∈ {active, invited} (`:1053`); `grep -in hiring` over that resolver and over
   `internal/organization/intelligence.go` returns only sim-TYPE filters and nothing, respectively (positive
   controls: `OrgFeatureInsights` ×8, `JobSimulationSession` ×44). **But "no read path reads it" would be a second false claim:** the CONTENT-LIBRARY read path does —
   `PrivateJobSimulations` branches its result set on `GetOrganizationIsHiring`
   (`resolver_cms_queries.go:95,210,258,295` — `isHiring` picks `hiringLibraryTypes()` over
   `workforceLibraryTypes()` at `:99-103`), as do `organization/manager.go:450` (a forced Clerk membership
   is created with role `candidate` instead of `member`, `:453`) and `:537` + `siminvitationlink.go:63` (both
   **hard-error `"organization is not hiring"`** — the latter is `CreateOrganizationSimInvitationLink`. Note
   the `HiringConfigSeeder` does **not** go through that RPC: it writes the 5 positions straight into
   `public.organization_sim_invitation_links` with `CopyRowsIdempotent` (`hiring_config.go:99`), so this
   hard-error never reaches it. This passage claimed the opposite until M257x iter-52). And the client re-skin is **not**
   driven by this column either: it is read from Clerk `publicMetadata.isHiring`
   (`useGetClerkOrganization.tsx:20`, quoted below). So the column gates the content library and the
   org-type surfaces; the *insights* scoreboard is indifferent to it.
```

**CITED CONTENT**

```
    96  	}
    97  
    98  	cols := []string{"id", "created_at", "updated_at", "simulation_id", "token", "options", "organization_id", "created_by"}
    99  	n, err := c.CopyRowsIdempotent(ctx, "public", "organization_sim_invitation_links", cols, rows, "id")
   100  	if err != nil {
   101  		return 0, fmt.Errorf("hiring-config: copy organization_sim_invitation_links: %w", err)
   102  	}
```

## 04-013
- **id**: `B04-013`
- **corpus site**: `corpus/services/hiring.md:131-159` (paragraph)
- **citation**: `hiring_config.go:99`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-seeding/seeders/hiring_config.go`  (130 lines)

**CLAIMING UNIT**

```md
> **Both, or the demo is half-lit — and the two halves fail in opposite directions.** Write each one down
> separately; a single "it doesn't re-skin" covers neither.
>
> - **DB-only** (column `true`, Clerk metadata absent) → **the client half is dead.** `isHiringOrg` is `false`,
>   so the nav keeps the "Activity dashboard" label (`useNavbarSections.tsx:476` @ `next-web-app` `8297c684`;
>   it was `:460` at `bb3313bc`) and the org is *not* filtered
>   out of the workforce list (`useGetClerkOrganization.tsx:16-18`). And the product-boundary hand-off — which
>   reads Clerk and **only** Clerk (`apps/web/src/context/UserStatusContext.tsx:144-149` computes
>   `userHasAllHiringOrgs` from `publicMetadata.isHiring`, then `:168-172` fires
>   `buildSwitchHandoffUrl({targetProduct:'hiring'})`) — **never fires**, so the recruiter is never handed to
>   `apps/hiring`; she sits in a Workforce-skinned `apps/web`. Point the cockpit at the hiring base anyway and
>   the *symmetric* guard bounces her straight back (`apps/hiring/src/context/UserStatusContext.tsx:125,144-145`
>   → `targetProduct:'workforce'`). This is exactly the `billion` spike that produced M222's false `apps/web`
>   premise (§ *The render path*). The server half, meanwhile, is entirely correct.
> - **Clerk-only** (metadata `true`, column `false`) → **the client half is fine and the server half is dead.**
>   The browser *does* re-skin (the re-skin reads Clerk, not the column) and the hand-off *does* route the
>   recruiter to `apps/hiring`. What breaks is server-side: the content library serves the **workforce**
>   type-set instead of the hiring one (`resolver_cms_queries.go:99-103`), and
>   `CreateOrganizationSimInvitationLink` hard-errors `"organization is not hiring"` (`siminvitationlink.go:63`,
>   guarded at `:62`)
>   for any caller that uses it. **The `HiringConfigSeeder` is not such a caller** — it writes the 5 positions
>   directly (`hiring_config.go:99`) and is unaffected. This bullet previously said it *"cannot write the 5
>   positions in the first place"*; that consequence is **refuted** (M257x iter-52).
>
> **Neither half, however, gates the insights scoreboard** — the text here used to say Clerk-only meant *"the
> insights read-path won't treat the cohort as hiring"*, and that sent every empty-scoreboard debug to the wrong
> place. What actually gates it: the `OrgFeatureInsights` Casb
```

**CITED CONTENT**

```
    96  	}
    97  
    98  	cols := []string{"id", "created_at", "updated_at", "simulation_id", "token", "options", "organization_id", "created_by"}
    99  	n, err := c.CopyRowsIdempotent(ctx, "public", "organization_sim_invitation_links", cols, rows, "id")
   100  	if err != nil {
   101  		return 0, fmt.Errorf("hiring-config: copy organization_sim_invitation_links: %w", err)
   102  	}
```

## 04-014
- **id**: `B04-014`
- **corpus site**: `corpus/services/hiring.md:196-209` (paragraph)
- **citation**: `repos.yml:14-17`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
> **History, because a seeder built from the old shape writes to nothing.** Until `app` migration
> `20260729133514.sql` (2026-07-29) the score lived on `public.local_jobsimulation_sessions`, a `Float32`
> **MIRROR** that shadowed a `jobsimulation.sessions` row. That migration **re-pointed the referencing
> assignment-session link ids** and **dropped the mirror** (`:58-62`) — it did *not* back-fill (`SET "score"`
> = 0 hits set-wide). The earlier `20260722104506.sql:79` dropped **`public.sessions`** (a bare
> `DROP TABLE "sessions"` under `search_path=public`) in favour of `public.job_simulation_sessions` (`:2`);
> **`jobsimulation.sessions` itself was NOT dropped** — no `app` migration touches that schema, and in
> **production** it survives frozen until M710. (`askengine/registry.go:192` is cited for the M710 horizon
> only: it is an LLM-facing name-alias map whose `jobsimulation.*` names **resolve to the public tables** —
> it is not evidence that the schema is physically present. On a **local dev/demo stack it is not**:
> `jobsimulation` has had **no `repos.yml` entry at all** since `d11a403` (6 entries @ platform `0dab54d`,
> 4 since `838d907`), and `app` (`repos.yml:14-17`) is the only repo with migrations to run. Qualified M257x iter-52,
> re-anchored in the M257x sweep.) So what is gone is the **mirror half** of the old
> pair, not both halves; there is one row per (candidate × attempt) now, in `public`. Corrected M257x iter-49.
```

**CITED CONTENT**

```
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
    14    - name: app
    15      type: go
    16      migrations: true
    17      schema: public
    18    - name: sentinel
    19      type: go
    20      migrations: false
```

## 04-015
- **id**: `B04-015`
- **corpus site**: `corpus/services/hiring.md:216-216` (table-row)
- **citation**: `packages/graphql/src/query/insights.ts:31-82`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/packages/graphql/src/query/insights.ts`  (328 lines)

**CLAIMING UNIT**

```md
| 2 | `packages/graphql/src/query/insights.ts:31-82` | query `insightsJobSimulationByMemberships` |
```

**CITED CONTENT**

```
    28    }
    29  `);
    30  
    31  export const GET_INSIGHTS_BY_JOB_SIMULATIONS_BY_MEMBERS = gql(`
    32    query insightsJobSimulationByMemberships($organizationId: ID!, $jobSimulationId: ID!, $params: InsightJobSimulationMembershipParams!, $language: ContentLanguage!) {
    33      insightsJobSimulationByMemberships(organizationId: $organizationId, jobSimulationId: $jobSimulationId, params: $params) {
    34        limit
    35        offset
    36        total
    37        rows {
    38          attemptsCount
    39          bestAttemptStartedAt
    40          bestAttemptTimeSpent
    41          interactionsScore
    42          anticheatResult
    43          jobSimulation(language: $language) {
    44            id
    45            slug
    46            type
    47          }
    48          jobSimulationSession {
    49            id
    50            startedAt
    51            endedAt
    52            status
    53          }
    54          membership {
    55            id
    56            firstName
    57            jobRole(language: $language) {
    58              id
    59              name
    60              createdByUser
    61            }
    62            jobRoleTitle
    63            lastName
    64            pictureUrl
    65            userId
    66            email
    67            tags {
    68              id
    69              name
    70            }
    71          }
    72          score
    73          sessionCompletionStatus
    74          skillPath {
    75            id
    76            title
    77            slug
    78          }
    79        }
    80      }
    81    }
    82  `);
    83  
    84  export const GET_INSIGHTS_BY_JOB_SIMULATIONS_BY_MEMBERS_BY_SESSIONS = gql(`
    85    query insightsJobSimulationBySessions($organizationId: ID!, $jobSimulationId: ID!, $membershipsId: ID!, $options: InsightOptions, $language: ContentLanguage!) {
```

## 04-016
- **id**: `B04-016`
- **corpus site**: `corpus/services/hiring.md:218-218` (table-row)
- **citation**: `app/internal/organization/intelligence.go:1700`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/organization/intelligence.go`  (2296 lines)

**CLAIMING UNIT**

```md
| 4 | `app/internal/organization/intelligence.go:1700` | reads `m.ent.JobSimulationSession` (the canonical entity; was `LocalJobsimulationSession` before the mirror drop) |
```

**CITED CONTENT**

```
  1697  	ctx = authorization.NewContextWithTargets(ctx, userIds)
  1698  	ctx = authorization.NewContextWithDecision(ctx, authorization.Allow)
  1699  
  1700  	query := m.ent.JobSimulationSession.Query().
  1701  		Where(
  1702  			jobsimulationsession.SimID(jobSimulationId),
  1703  			jobsimulationsession.OwnerIDIn(userIds...),
```

## 04-017
- **id**: `B04-017`
- **corpus site**: `corpus/services/hiring.md:221-221` (table-row)
- **citation**: `app/internal/data/ent/schema/job_simulation_session.go:45`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/data/ent/schema/job_simulation_session.go`  (79 lines)

**CLAIMING UNIT**

```md
| 7 | `app/internal/data/ent/schema/job_simulation_session.go:45` | Ent table `public.job_simulation_sessions`, `field.Float32("score").Default(0).Min(0).Max(100)` — **the score column, read at `intelligence.go:1820` and assigned at `:1846`. Not a mirror: `local_jobsimulation_session.go` no longer exists** |
```

**CITED CONTENT**

```
    42  		field.Time("started_at").Optional().Nillable(),
    43  		field.Time("ended_at").Optional().Nillable(),
    44  		field.Int("interactions_progress").Default(0).Min(0).Max(100),
    45  		field.Float32("score").Default(0).Min(0).Max(100),
    46  		field.Int("validation_version").Default(2).Min(1).Max(3).Immutable(),
    47  		field.Time("scheduled_timeout").Optional().Nillable(),
    48  		field.UUID("timer_id", uuid.UUID{}).Optional().Nillable(),
```

## 04-018
- **id**: `B04-018`
- **corpus site**: `corpus/services/hiring.md:241-252` (paragraph)
- **citation**: `persona_write.go:69-71`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-seeding/seeders/persona_write.go`  (215 lines)

**CLAIMING UNIT**

```md
**BA-4 — the drill-down is a DIFFERENT set of tables (not the scoreboard).** Clicking a candidate opens the
per-session competency / Job-Fit detail (`[simId]/[userId]`), which reads
`public.validation_attempt_results` / `validation_attempt_skill_results` / `validation_criterion_results` — three
tables (all in **`public`**, `20260722081626_jobsim_data_model.sql:336/355/376`; note the middle one is
`validation_attempt_skill_results`, not `validation_skill_results`) the `PersonaSeeder` also writes (`persona_write.go:69-71,143-167`). These are needed **only for the
drill-down**, NOT for the comparison list. The anticheat badge is a **decorative icon only**, and it is **not a
column on the session row** — it is `summary` on the separate **`public.anticheat_results`** entity
(`ent/schema/anticheat_result.go:24`), whose `session_id` FK was re-pointed at `job_simulation_sessions` by
`20260722104506.sql:53`. So
the open BA-1 question — *"does the list score need a per-session `validation_*`/eval row?"* — is answered **NO**:
the scoreboard scores from the **single** `job_simulation_sessions` row (+ membership + the Casbin gate)
alone — the write-set used to be a PAIR and is now one row, since the mirrors were dropped.
```

**CITED CONTENT**

```
    66  //	jobsimulation.validation_attempt_results        -> public.validation_attempt_results
    67  //	jobsimulation.validation_attempt_skill_results  -> public.validation_attempt_skill_results
    68  //	jobsimulation.validation_criterion_results      -> public.validation_criterion_results
    69  //
    70  // The old `jobsimulation.sessions` step is REMOVED rather than re-pointed: it wrote the SAME
    71  // a.sessions rows under the SAME id as the `public.job_simulation_sessions` step below, so
    72  // re-pointing it would have produced two identical writes. Per platform-alignment.md §7 rule 2 a
    73  // removed write needs its replacement ASSERTED, and the replacement is the very next step — kept,
    74  // and MOVED FIRST because the validation rows FK it (`validation_attempt_results.session_id ->
```

## 04-019
- **id**: `B04-019`
- **corpus site**: `corpus/services/hiring.md:258-294` (bullet)
- **citation**: `app/internal/data/ent/enum/jobsimulation.go:29-35`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/data/ent/enum/jobsimulation.go`  (230 lines)

**CLAIMING UNIT**

```md
1. **`public.job_simulation_sessions`** — the **score source** + row generator, and the only session row there
   is. Non-null `status`, `started_at`, `ended_at`, `owner_id`, `sim_id`, `sim_type`, **`token`**, plus
   `score` (0–100),
   `completion_status` (a closed 5-value enum — **exactly** `pending` / `passed` / `failed` / `discarded` /
   `timedout`, `app/internal/data/ent/enum/jobsimulation.go:29-35`, `Values()` at `:37-43`; **no `SIMULATION…`
   member** — that prefix belongs to the adjacent `sim_type` column, which genuinely is `SIMULATION_TYPE_*`),
   `organization_id`, `tenant_id` (NULL or `=org`), `validation_version`.
   ⚠️ **`token` is the one column that makes the INSERT itself fail, and this contract omitted it until
   M257x iter-49.** It is `NOT NULL` (`20260722104506.sql:13`), `UNIQUE` (`:29`) and carries **no default** —
   one of **four** required-and-undefaulted columns (`owner_id` `:6`, `sim_id` `:7`, `sim_type` `:10`,
   `token` `:13`; every other `NOT NULL` column in the DDL carries a `DEFAULT`) — so an INSERT built from the write-set as it was
   written here does not render wrong, it **errors**. The shipped seeder has always written it
   (`persona_write.go:152-158`); the word `token` simply appeared nowhere in this document. Being UNIQUE, it
   must be generated per row, not reused. (iter-47 read this passage and booked it a MINOR; iter-48's seat
   escalated it after checking the DDL, the Ent schema and the seeder.)
   ⚠️ **Get `completion_status` wrong and NOTHING catches it — the row does not vanish, it renders wrong.**
   The column is a plain `varchar` with **no CHECK** (a rolled-back
   `UPDATE … SET completion_status='SIMULATION_COMPLETION_STATUS_PASSED'` is accepted); Ent's generated
   `assignValues` casts **unconditionally** and cannot error (`ent/jobsimulationsession.go:181-186`:
   `_m.CompletionStatus = enum.SessionCompletionStatus(value.String)`); the read-model re-casts just as blindly
   (`intelligence.go:1844`); and the gqlgen enum marshal is a bare `graphql.MarshalString(string(v))`
   passthrough with **no** membership check (`graphql/graph/graph.go:129546-129554`, and the proto-bound twin at
   `:129392-129400`) — even though the SDL declares only the five lowercase members
   (`graphql/graph/schemas/jobsimulations.graphqls:14` and `:128`). So a raw-SQL seeder writing a `SIMULATION_…`
   value INSERTs cleanly **
```

**CITED CONTENT**

```
    26  
    27  type SessionCompletionStatus string
    28  
    29  const (
    30  	CompletionStatusPending   SessionCompletionStatus = "pending"
    31  	CompletionStatusPassed    SessionCompletionStatus = "passed"
    32  	CompletionStatusFailed    SessionCompletionStatus = "failed"
    33  	CompletionStatusDiscarded SessionCompletionStatus = "discarded"
    34  	CompletionStatusTimedout  SessionCompletionStatus = "timedout"
    35  )
    36  
    37  func (SessionCompletionStatus) Values() []string {
    38  	return []string{
```

## 04-020
- **id**: `B04-020`
- **corpus site**: `corpus/services/hiring.md:258-294` (bullet)
- **citation**: `persona_write.go:152-158`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-seeding/seeders/persona_write.go`  (215 lines)

**CLAIMING UNIT**

```md
1. **`public.job_simulation_sessions`** — the **score source** + row generator, and the only session row there
   is. Non-null `status`, `started_at`, `ended_at`, `owner_id`, `sim_id`, `sim_type`, **`token`**, plus
   `score` (0–100),
   `completion_status` (a closed 5-value enum — **exactly** `pending` / `passed` / `failed` / `discarded` /
   `timedout`, `app/internal/data/ent/enum/jobsimulation.go:29-35`, `Values()` at `:37-43`; **no `SIMULATION…`
   member** — that prefix belongs to the adjacent `sim_type` column, which genuinely is `SIMULATION_TYPE_*`),
   `organization_id`, `tenant_id` (NULL or `=org`), `validation_version`.
   ⚠️ **`token` is the one column that makes the INSERT itself fail, and this contract omitted it until
   M257x iter-49.** It is `NOT NULL` (`20260722104506.sql:13`), `UNIQUE` (`:29`) and carries **no default** —
   one of **four** required-and-undefaulted columns (`owner_id` `:6`, `sim_id` `:7`, `sim_type` `:10`,
   `token` `:13`; every other `NOT NULL` column in the DDL carries a `DEFAULT`) — so an INSERT built from the write-set as it was
   written here does not render wrong, it **errors**. The shipped seeder has always written it
   (`persona_write.go:152-158`); the word `token` simply appeared nowhere in this document. Being UNIQUE, it
   must be generated per row, not reused. (iter-47 read this passage and booked it a MINOR; iter-48's seat
   escalated it after checking the DDL, the Ent schema and the seeder.)
   ⚠️ **Get `completion_status` wrong and NOTHING catches it — the row does not vanish, it renders wrong.**
   The column is a plain `varchar` with **no CHECK** (a rolled-back
   `UPDATE … SET completion_status='SIMULATION_COMPLETION_STATUS_PASSED'` is accepted); Ent's generated
   `assignValues` casts **unconditionally** and cannot error (`ent/jobsimulationsession.go:181-186`:
   `_m.CompletionStatus = enum.SessionCompletionStatus(value.String)`); the read-model re-casts just as blindly
   (`intelligence.go:1844`); and the gqlgen enum marshal is a bare `graphql.MarshalString(string(v))`
   passthrough with **no** membership check (`graphql/graph/graph.go:129546-129554`, and the proto-bound twin at
   `:129392-129400`) — even though the SDL declares only the five lowercase members
   (`graphql/graph/schemas/jobsimulations.graphqls:14` and `:128`). So a raw-SQL seeder writing a `SIMULATION_…`
   value INSERTs cleanly **
```

**CITED CONTENT**

```
   149  // iter-06 removed the legacy half, so there is now exactly ONE target: `public.
   150  // job_simulation_sessions` (created by app migration 20260722104506.sql). All 18 columns
   151  // verified present on the live migrated stack — see TestSeederWriteTargetsExistLive.
   152  func sessionCols() []string {
   153  	return []string{
   154  		"id", "created_at", "updated_at", "owner_id", "sim_id", "sim_type",
   155  		"status", "completion_status", "token", "started_at", "ended_at",
   156  		"score", "result_status", "validation_version", "language",
   157  		"chime_status", "interactions_progress", "organization_id",
   158  	}
   159  }
   160  
   161  func attemptResultCols() []string {
```

## 04-021
- **id**: `B04-021`
- **corpus site**: `corpus/services/hiring.md:311-320` (paragraph)
- **citation**: `rosetta-extensions/stack-seeding/seeders/persona_write.go:91`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-seeding/seeders/persona_write.go`  (215 lines)

**CLAIMING UNIT**

```md
**The machinery already exists — M223 is NOT net-new.** The **`PersonaSeeder` already writes exactly this row** —
`rosetta-extensions/stack-seeding/seeders/persona_write.go:91` writes
`{"public", "job_simulation_sessions", sessionCols(), …}`. (Until M257 it wrote the **pair**, via a second col
builder `localSessionCols()`; that builder was deleted with the mirror and `sessionCols()` at `:152` now serves
the single canonical row.) M223's
candidate-assessment funnel is a **direct generalization** of the same fan-out — each candidate on the **one**
position they applied for (v2.4 "casting call" M227 fix #3; before M227 every candidate took all 5), round-robined
evenly across the 5 shared sims so ~8 candidates rank per position (the M51 `AIReadinessFunnelSeeder` shape, 2 shared
sims → 5) — with the M219 anti-junk discipline (a realistic non-degenerate score DISTRIBUTION, every skill/role ref
through the real resolvers, closure green, never fabricated), **not** a flat score grid.
```

**CITED CONTENT**

```
    88  		cols          []string
    89  		rows          [][]any
    90  	}{
    91  		{"public", "job_simulation_sessions", sessionCols(), a.sessions},
    92  		{"public", "validation_attempt_results", attemptResultCols(), a.attemptResults},
    93  		{"public", "validation_attempt_skill_results", skillResultCols(), a.skillResults},
    94  		{"public", "validation_criterion_results", criterionResultCols(), a.criterionResults},
```

## 04-022
- **id**: `B04-022`
- **corpus site**: `corpus/services/hiring.md:338-354` (bullet)
- **citation**: `packages/ui/src/NavBar/orgGroups.ts:48-65`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/packages/ui/src/NavBar/orgGroups.ts`  (88 lines)

**CLAIMING UNIT**

```md
- **Also under `is_hiring`:** the nav trims the Content Library to **AI-Simulations** alone
  (`packages/ui/src/NavBar/useNavbarSections.tsx:340-343` — `isHiringOrg` selects
  `[librarySimulationsMenuItem]`, label `tNavbar('aiSimulations')` at `:249-256`, instead of the
  simulations + skill-paths + academy + labs set), and hides the member Profile / Skills / Activities
  entries for non-admins (`:329-331`, each `!isHiringOrg || isAdmin`). Both clauses verify.
  ⚠️ **It does NOT gate Workforce Intelligence off — that clause is RETRACTED (M257x iter-102).** Nothing
  gates Workforce Intelligence on `isHiringOrg`. The entry is `enterpriseWorkforceMenuItem`
  (`tNavbar('workforceIntelligence')`, `:391-398`); it sits in the `intelligence` group, whose visibility
  comes from `orgSectionVisibility({ isAdmin, showStudio })` returning `intelligence: isAdmin`
  (`packages/ui/src/NavBar/orgGroups.ts:48-65`, the field at `:61`) — a function that **takes no
  `isHiringOrg` parameter at all** — and the item itself is gated on `showWorkforce`
  (`useNavbarSections.tsx:568`), which **defaults to `true`** (`useNavbarSections.tsx:161`) and is passed
  `false` in exactly **two** places, **both in `apps/hiring`**
  (`apps/hiring/src/app/(authenticated)/(verified)/template.tsx:167` and `:248`). So a recruiter loses
  Workforce Intelligence by being handed off to **`apps/hiring`** (§ *The render path*), not by the
  `is_hiring` flag inside `apps/web`; an `is_hiring` org's admin still browsing `apps/web` keeps it.
  Measured at `next-web-app` `8297c684`. None of these touch the comparison scoreboard.
```

**CITED CONTENT**

```
    45   * role Studio exists for. A content creator therefore sees only
    46   * Organization → Customize → Studio; a plain member sees no Organization at all.
    47   */
    48  export function orgSectionVisibility({
    49    isAdmin,
    50    showStudio,
    51  }: {
    52    isAdmin: boolean;
    53    showStudio: boolean;
    54  }): OrgSectionVisibility {
    55    return {
    56      groups: {
    57        map: isAdmin,
    58        customize: showStudio,
    59        assign: isAdmin,
    60        trackVerify: isAdmin,
    61        intelligence: isAdmin,
    62      },
    63      directLinks: isAdmin,
    64    };
    65  }
    66  
    67  /**
    68   * Assemble the Organization groups from already-gated item arrays.
```

## 04-023
- **id**: `B04-023`
- **corpus site**: `corpus/services/jobsimulation.md:3-82` (paragraph)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"jobsim-in-app"** program, the standalone `jobsimulation` Go microservice has been **merged into
> the `app` monolith** (the service the platform calls "backend"). Jobsimulation no longer runs as a separate
> service **in production** — and since `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository
> (M810)"*) it cannot be *started* there either: **the `module "jobsimulation"` block is deleted**, so
> `service_desired_count` does not appear anywhere in `jobsimulation/terraform/main.tf` (`:15-22`). Its
> subgraph is gone from the supergraph. **M810 has LANDED for the ECS service**; what it has not yet done here
> is drop the legacy `jobsimulation` schema, a deliberately separate step (`:38-40`). **Do not generalise this
> to `cms`, in EITHER direction** — `cms`'s two measured facts point **opposite ways**, and the corpus's flat *"cms has not moved"* is half of them. Measured at `cms` `origin/main` `f38c0c4a` (2026-08-06): the module block has *not* moved — `cms/terraform/main.tf:39` still reads `service_desired_count = 0` in an otherwise-whole 191-line module — **but** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** that repo's `.github/workflows/build-production.yml` under the subject *"the cms ECR repository is decommissioned (M810)"*, because it *"would try to push an image into a registry that no longer exists."* The deciding declaration is in `infrastructure`, which has never been in any clone set: **report both, assert neither** — see [`cms.md`](./cms.md) and the fenced map.
>
> **✅ The husk is GONE locally too (re-measured at platform `0c91421`).** There is no `jobsimulation` compose
> service, no `jobsimulation` entry in `repos.yml` (4 entries: app, sentinel, next-web-app, studio-desk)
> and no `jobsimulation` profile. Platform **`d11a403`** (2026-08-03) deleted
> both in one commit — its `repos.yml` diff removes `- name: cms`, `- name: jobsimulation` **and**
> `- name: roadrunner`. (The entry list read *"6 … storage, messenger"* at `0dab54d`; `838d907` removed
> those two a day later.)
> *This banner used to read "**but locally the husk still starts**", and it was right at `2adcf71`:
> `docker-compose.yml:83` @ that ref defined a `jobsimulation` service with
> `profiles: [graphql, jobsimulation, all]` (`:140`), `graphql` was the default (`Makefile:10`
> `
```

**CITED CONTENT**

```
    36    tags                           = var.tags
    37    aws_region                     = var.aws_region
    38    project                        = local.project
    39    service_desired_count          = 0
    40    service_cpu                    = local.service_cpu
    41    service_memory                 = local.service_memory
    42    health_check_path              = "/_meta"
```

## 04-024
- **id**: `B04-024`
- **corpus site**: `corpus/services/jobsimulation.md:3-82` (paragraph)
- **citation**: `docker-compose.yml:83`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"jobsim-in-app"** program, the standalone `jobsimulation` Go microservice has been **merged into
> the `app` monolith** (the service the platform calls "backend"). Jobsimulation no longer runs as a separate
> service **in production** — and since `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository
> (M810)"*) it cannot be *started* there either: **the `module "jobsimulation"` block is deleted**, so
> `service_desired_count` does not appear anywhere in `jobsimulation/terraform/main.tf` (`:15-22`). Its
> subgraph is gone from the supergraph. **M810 has LANDED for the ECS service**; what it has not yet done here
> is drop the legacy `jobsimulation` schema, a deliberately separate step (`:38-40`). **Do not generalise this
> to `cms`, in EITHER direction** — `cms`'s two measured facts point **opposite ways**, and the corpus's flat *"cms has not moved"* is half of them. Measured at `cms` `origin/main` `f38c0c4a` (2026-08-06): the module block has *not* moved — `cms/terraform/main.tf:39` still reads `service_desired_count = 0` in an otherwise-whole 191-line module — **but** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** that repo's `.github/workflows/build-production.yml` under the subject *"the cms ECR repository is decommissioned (M810)"*, because it *"would try to push an image into a registry that no longer exists."* The deciding declaration is in `infrastructure`, which has never been in any clone set: **report both, assert neither** — see [`cms.md`](./cms.md) and the fenced map.
>
> **✅ The husk is GONE locally too (re-measured at platform `0c91421`).** There is no `jobsimulation` compose
> service, no `jobsimulation` entry in `repos.yml` (4 entries: app, sentinel, next-web-app, studio-desk)
> and no `jobsimulation` profile. Platform **`d11a403`** (2026-08-03) deleted
> both in one commit — its `repos.yml` diff removes `- name: cms`, `- name: jobsimulation` **and**
> `- name: roadrunner`. (The entry list read *"6 … storage, messenger"* at `0dab54d`; `838d907` removed
> those two a day later.)
> *This banner used to read "**but locally the husk still starts**", and it was right at `2adcf71`:
> `docker-compose.yml:83` @ that ref defined a `jobsimulation` service with
> `profiles: [graphql, jobsimulation, all]` (`:140`), `graphql` was the default (`Makefile:10`
> `
```

**CITED CONTENT**

```
    80        - AWS_REGION=eu-west-1
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
    86        # stream or a timer — they send mail and rewrite Brevo contacts — so app gates them
```

## 04-025
- **id**: `B04-025`
- **corpus site**: `corpus/services/jobsimulation.md:3-82` (paragraph)
- **citation**: `repos.yml:17-19`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"jobsim-in-app"** program, the standalone `jobsimulation` Go microservice has been **merged into
> the `app` monolith** (the service the platform calls "backend"). Jobsimulation no longer runs as a separate
> service **in production** — and since `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository
> (M810)"*) it cannot be *started* there either: **the `module "jobsimulation"` block is deleted**, so
> `service_desired_count` does not appear anywhere in `jobsimulation/terraform/main.tf` (`:15-22`). Its
> subgraph is gone from the supergraph. **M810 has LANDED for the ECS service**; what it has not yet done here
> is drop the legacy `jobsimulation` schema, a deliberately separate step (`:38-40`). **Do not generalise this
> to `cms`, in EITHER direction** — `cms`'s two measured facts point **opposite ways**, and the corpus's flat *"cms has not moved"* is half of them. Measured at `cms` `origin/main` `f38c0c4a` (2026-08-06): the module block has *not* moved — `cms/terraform/main.tf:39` still reads `service_desired_count = 0` in an otherwise-whole 191-line module — **but** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** that repo's `.github/workflows/build-production.yml` under the subject *"the cms ECR repository is decommissioned (M810)"*, because it *"would try to push an image into a registry that no longer exists."* The deciding declaration is in `infrastructure`, which has never been in any clone set: **report both, assert neither** — see [`cms.md`](./cms.md) and the fenced map.
>
> **✅ The husk is GONE locally too (re-measured at platform `0c91421`).** There is no `jobsimulation` compose
> service, no `jobsimulation` entry in `repos.yml` (4 entries: app, sentinel, next-web-app, studio-desk)
> and no `jobsimulation` profile. Platform **`d11a403`** (2026-08-03) deleted
> both in one commit — its `repos.yml` diff removes `- name: cms`, `- name: jobsimulation` **and**
> `- name: roadrunner`. (The entry list read *"6 … storage, messenger"* at `0dab54d`; `838d907` removed
> those two a day later.)
> *This banner used to read "**but locally the husk still starts**", and it was right at `2adcf71`:
> `docker-compose.yml:83` @ that ref defined a `jobsimulation` service with
> `profiles: [graphql, jobsimulation, all]` (`:140`), `graphql` was the default (`Makefile:10`
> `
```

**CITED CONTENT**

```
    14    - name: app
    15      type: go
    16      migrations: true
    17      schema: public
    18    - name: sentinel
    19      type: go
    20      migrations: false
    21  
    22    # Frontend
```

## 04-026
- **id**: `B04-026`
- **corpus site**: `corpus/services/jobsimulation.md:3-82` (paragraph)
- **citation**: `app/knowledge/service-dependencies.md:52`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/knowledge/service-dependencies.md`  (122 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"jobsim-in-app"** program, the standalone `jobsimulation` Go microservice has been **merged into
> the `app` monolith** (the service the platform calls "backend"). Jobsimulation no longer runs as a separate
> service **in production** — and since `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository
> (M810)"*) it cannot be *started* there either: **the `module "jobsimulation"` block is deleted**, so
> `service_desired_count` does not appear anywhere in `jobsimulation/terraform/main.tf` (`:15-22`). Its
> subgraph is gone from the supergraph. **M810 has LANDED for the ECS service**; what it has not yet done here
> is drop the legacy `jobsimulation` schema, a deliberately separate step (`:38-40`). **Do not generalise this
> to `cms`, in EITHER direction** — `cms`'s two measured facts point **opposite ways**, and the corpus's flat *"cms has not moved"* is half of them. Measured at `cms` `origin/main` `f38c0c4a` (2026-08-06): the module block has *not* moved — `cms/terraform/main.tf:39` still reads `service_desired_count = 0` in an otherwise-whole 191-line module — **but** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** that repo's `.github/workflows/build-production.yml` under the subject *"the cms ECR repository is decommissioned (M810)"*, because it *"would try to push an image into a registry that no longer exists."* The deciding declaration is in `infrastructure`, which has never been in any clone set: **report both, assert neither** — see [`cms.md`](./cms.md) and the fenced map.
>
> **✅ The husk is GONE locally too (re-measured at platform `0c91421`).** There is no `jobsimulation` compose
> service, no `jobsimulation` entry in `repos.yml` (4 entries: app, sentinel, next-web-app, studio-desk)
> and no `jobsimulation` profile. Platform **`d11a403`** (2026-08-03) deleted
> both in one commit — its `repos.yml` diff removes `- name: cms`, `- name: jobsimulation` **and**
> `- name: roadrunner`. (The entry list read *"6 … storage, messenger"* at `0dab54d`; `838d907` removed
> those two a day later.)
> *This banner used to read "**but locally the husk still starts**", and it was right at `2adcf71`:
> `docker-compose.yml:83` @ that ref defined a `jobsimulation` service with
> `profiles: [graphql, jobsimulation, all]` (`:140`), `graphql` was the default (`Makefile:10`
> `
```

**CITED CONTENT**

```
    49  >
    50  > **There are no external callers of app's RPC mux left.** `messenger` was the last one — it used to
    51  > reach the users, cms, jobsimulation and skiller surfaces at
    52  > `http://backend.internal.anthropos:8081`, and folding it in at v9.0 closed that edge. The mux is
    53  > kept because it is how the in-process domains are wired, not because something outside dials it.
    54  > App also keeps emitting `JOBSIMULATION_STREAM` as an in-process loopback — it feeds the
    55  > real consumers (XP/skills/quota/assignment link/AI Readiness); the `LocalJobsimulationSession`
```

## 04-027
- **id**: `B04-027`
- **corpus site**: `corpus/services/jobsimulation.md:3-82` (paragraph)
- **citation**: `app/main.go:1204`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"jobsim-in-app"** program, the standalone `jobsimulation` Go microservice has been **merged into
> the `app` monolith** (the service the platform calls "backend"). Jobsimulation no longer runs as a separate
> service **in production** — and since `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository
> (M810)"*) it cannot be *started* there either: **the `module "jobsimulation"` block is deleted**, so
> `service_desired_count` does not appear anywhere in `jobsimulation/terraform/main.tf` (`:15-22`). Its
> subgraph is gone from the supergraph. **M810 has LANDED for the ECS service**; what it has not yet done here
> is drop the legacy `jobsimulation` schema, a deliberately separate step (`:38-40`). **Do not generalise this
> to `cms`, in EITHER direction** — `cms`'s two measured facts point **opposite ways**, and the corpus's flat *"cms has not moved"* is half of them. Measured at `cms` `origin/main` `f38c0c4a` (2026-08-06): the module block has *not* moved — `cms/terraform/main.tf:39` still reads `service_desired_count = 0` in an otherwise-whole 191-line module — **but** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** that repo's `.github/workflows/build-production.yml` under the subject *"the cms ECR repository is decommissioned (M810)"*, because it *"would try to push an image into a registry that no longer exists."* The deciding declaration is in `infrastructure`, which has never been in any clone set: **report both, assert neither** — see [`cms.md`](./cms.md) and the fenced map.
>
> **✅ The husk is GONE locally too (re-measured at platform `0c91421`).** There is no `jobsimulation` compose
> service, no `jobsimulation` entry in `repos.yml` (4 entries: app, sentinel, next-web-app, studio-desk)
> and no `jobsimulation` profile. Platform **`d11a403`** (2026-08-03) deleted
> both in one commit — its `repos.yml` diff removes `- name: cms`, `- name: jobsimulation` **and**
> `- name: roadrunner`. (The entry list read *"6 … storage, messenger"* at `0dab54d`; `838d907` removed
> those two a day later.)
> *This banner used to read "**but locally the husk still starts**", and it was right at `2adcf71`:
> `docker-compose.yml:83` @ that ref defined a `jobsimulation` service with
> `profiles: [graphql, jobsimulation, all]` (`:140`), `graphql` was the default (`Makefile:10`
> `
```

**CITED CONTENT**

```
  1201  	// read cms via the in-process RPC server instead of over the wire — no traffic to the
  1202  	// standalone cms. Active whenever the Directus edge is configured (the release sets it);
  1203  	// the external client the switch was seeded with is only the construction-time placeholder.
  1204  	cmsReaderSw.set(cmsRPCServer)
  1205  	// M805: consume the cms studio + ai_video Asynq queue in-process (the app is the sole
  1206  	// consumer post-release — the standalone cms takes no traffic). The consumer polls the SAME
  1207  	// DB index the enqueue client writes to (audit R2). The studio gen.py/postgen.py pipeline
```

## 04-028
- **id**: `B04-028`
- **corpus site**: `corpus/services/jobsimulation.md:3-82` (paragraph)
- **citation**: `jobsimulation/terraform/main.tf:15-22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/terraform/main.tf`  (344 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"jobsim-in-app"** program, the standalone `jobsimulation` Go microservice has been **merged into
> the `app` monolith** (the service the platform calls "backend"). Jobsimulation no longer runs as a separate
> service **in production** — and since `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository
> (M810)"*) it cannot be *started* there either: **the `module "jobsimulation"` block is deleted**, so
> `service_desired_count` does not appear anywhere in `jobsimulation/terraform/main.tf` (`:15-22`). Its
> subgraph is gone from the supergraph. **M810 has LANDED for the ECS service**; what it has not yet done here
> is drop the legacy `jobsimulation` schema, a deliberately separate step (`:38-40`). **Do not generalise this
> to `cms`, in EITHER direction** — `cms`'s two measured facts point **opposite ways**, and the corpus's flat *"cms has not moved"* is half of them. Measured at `cms` `origin/main` `f38c0c4a` (2026-08-06): the module block has *not* moved — `cms/terraform/main.tf:39` still reads `service_desired_count = 0` in an otherwise-whole 191-line module — **but** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** that repo's `.github/workflows/build-production.yml` under the subject *"the cms ECR repository is decommissioned (M810)"*, because it *"would try to push an image into a registry that no longer exists."* The deciding declaration is in `infrastructure`, which has never been in any clone set: **report both, assert neither** — see [`cms.md`](./cms.md) and the fenced map.
>
> **✅ The husk is GONE locally too (re-measured at platform `0c91421`).** There is no `jobsimulation` compose
> service, no `jobsimulation` entry in `repos.yml` (4 entries: app, sentinel, next-web-app, studio-desk)
> and no `jobsimulation` profile. Platform **`d11a403`** (2026-08-03) deleted
> both in one commit — its `repos.yml` diff removes `- name: cms`, `- name: jobsimulation` **and**
> `- name: roadrunner`. (The entry list read *"6 … storage, messenger"* at `0dab54d`; `838d907` removed
> those two a day later.)
> *This banner used to read "**but locally the husk still starts**", and it was right at `2adcf71`:
> `docker-compose.yml:83` @ that ref defined a `jobsimulation` service with
> `profiles: [graphql, jobsimulation, all]` (`:140`), `graphql` was the default (`Makefile:10`
> `
```

**CITED CONTENT**

```
    12    }
    13  }
    14  
    15  // Inspect the target database and load its state.
    16  // This is used to determine which migration to run.
    17  data "atlas_migration" "jobsimulation_migrations" {
    18    dir = "${path.module}/migrations?format=atlas"
    19    url = "${aws_ssm_parameter.db_connection.value}?search_path=jobsimulation"
    20  }
    21  
    22  // Sync the state of the target database with the migrations directory.
    23  resource "atlas_migration" "jobsimulation_migrations" {
    24    dir              = "${path.module}/migrations?format=atlas"
    25    version          = data.atlas_migration.jobsimulation_migrations.latest # Use latest to run all migrations
```

## 04-029
- **id**: `B04-029`
- **corpus site**: `corpus/services/jobsimulation.md:100-100` (bullet)
- **citation**: `cmd/root.go:77`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
* **Ports**: **8080 (GraphQL/HTTP), 8081 (Connect-RPC) — the binary's own defaults**, and now the only ones there are: `cmd/root.go:77` `cmp.Or(os.Getenv("PORT"), "8080")` / `:78` `cmp.Or(os.Getenv("RPC_PORT"), "8081")` (the Dockerfiles `EXPOSE 8080`), which is what the in-repo `CLAUDE.md` documents. The **8400 / 8401** pair quoted all over this corpus was **compose-supplied by a service that no longer exists**: `docker-compose.yml` set `PORT=8400` (`:113`) / `RPC_PORT=8401` (`:119`) and published `8400:8400` / `8401:8401` (`:93-94`) — **at `2adcf71`**. At `0c91421` there is no `jobsimulation` service, so nothing sets those values and nothing is published; **8400/8401 are historical, not an address you can reach**, with or without a `dev-N`/`demo-N` offset. The engine's live HTTP/GraphQL surface is `backend`'s.
```

**CITED CONTENT**

```
    74  		)
    75  		defer colony.FlushLogEvents()
    76  
    77  		port := cmp.Or(os.Getenv("PORT"), "8080")
    78  		rpcPort := cmp.Or(os.Getenv("RPC_PORT"), "8081")
    79  
    80  		authnClerk, err := clerk.NewProvider(os.Getenv("CLERK_SECRET_KEY"))
```

## 04-030
- **id**: `B04-030`
- **corpus site**: `corpus/services/jobsimulation.md:127-128` (bullet)
- **citation**: `app/main.go:1204`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* **RPC**: `internal/rpcsrv` — reached **in-process** by Backend (incl. the in-process skill-path engine), and **over the wire by nobody**. `messenger` was the last reader of `JOBSIMULATION_RPC_ADDR`; the value was `http://backend:8083` at `0dab54d`, set on messenger's block and nowhere else (`d11a403` had already dropped it from `backend`, having verified zero reads in `app`, **and re-pointed it — one of the MIDDLE TWO, with `CMS_RPC_ADDR` — on messenger's own block**; the other two of the four were already at `http://backend:8083` at `d11a403^` and were not touched), and `838d907` then deleted the messenger service — so **no compose file sets it, or any other `*_RPC_ADDR`, today**. There is no husk container left to resolve to either. `app` registers its own in-app `JobSimulationService` handler (`app/main.go:1204` @ `app` `b948604` v1.366.0).
  > **This line used to say the opposite, emphatically — keep the note (M257x iter-60).** Until `2adcf71` it read *"That address is **CURRENT, not stale text**"*, and it was **right at that ref**: only `SKILLER_RPC_ADDR` had been re-pointed then. A refutation is a measurement and expires exactly like the claim it refuted — and anti-repair wording is the kind that survives readings, because it looks already-adjudicated. See [`platform-alignment.md`](../ops/platform-alignment.md) §5 rule 31.
```

**CITED CONTENT**

```
  1201  	// read cms via the in-process RPC server instead of over the wire — no traffic to the
  1202  	// standalone cms. Active whenever the Directus edge is configured (the release sets it);
  1203  	// the external client the switch was seeded with is only the construction-time placeholder.
  1204  	cmsReaderSw.set(cmsRPCServer)
  1205  	// M805: consume the cms studio + ai_video Asynq queue in-process (the app is the sole
  1206  	// consumer post-release — the standalone cms takes no traffic). The consumer polls the SAME
  1207  	// DB index the enqueue client writes to (audit R2). The studio gen.py/postgen.py pipeline
```

## 04-031
- **id**: `B04-031`
- **corpus site**: `corpus/services/jobsimulation.md:130-143` (paragraph)
- **citation**: `internal/graph/queries.resolvers.go:70`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/internal/graph/queries.resolvers.go`  (503 lines)

**CLAIMING UNIT**

```md
> **Session/result READ-MODEL — this doc is not the home for it.** Two things a reader looking for "how does a
> played session render?" will not find here. (1) The **player** result page `/sim/<slug>/result/<sessionId>` is a
> **persisted read**, not a live recompute — `internal/graph/queries.resolvers.go:70` does plain Ent SELECTs over
> `validation_attempt_results`, so a seeded result fan-out renders a full result. (2) The **manager** view
> reads the **same** table — **the mirrors are GONE.** `app/terraform/migrations/20260729133514.sql:58-62`
> (*"5. Drop the mirrors."*) **re-points** `organization_assignment_sessions`' two foreign keys off the mirror
> ids (`:15-23`), NULLs the orphans (`:36-44`), then `DROP TABLE`s both `local_jobsimulation_sessions` and
> `local_skill_path_sessions` (`:62-63`), and `intelligence.go:1700` now reads `m.ent.JobSimulationSession.Query()`.
> **No session row is back-filled — the file contains 0 `INSERT`s** (this said *"back-fills then DROPs"* until
> M257x iter-52).
> **There is one row to seed, not a pair** — the older "seed the mirror or the scoreboard is blank"
> guidance is superseded. Full route-by-route treatment lives in
> [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md); the write side is
> [`../ops/demo/session-clone-spec.md`](../ops/demo/session-clone-spec.md).
```

**CITED CONTENT**

```
    67  func (r *queryResolver) ExportJobSimulations(ctx context.Context, options directus.ListOptions) ([]*simulation.JobSimulation, error) {
    68  	if !r.authManager.CheckIsSuperUser(ctx) {
    69  		return nil, fmt.Errorf("forbidden")
    70  	}
    71  
    72  	return r.jobSimulations.ExportJobSimulations(ctx, nil, options)
    73  }
```

## 04-032
- **id**: `B04-032`
- **corpus site**: `corpus/services/jobsimulation.md:130-143` (paragraph)
- **citation**: `app/terraform/migrations/20260729133514.sql:58-62`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/migrations/20260729133514.sql`  (65 lines)

**CLAIMING UNIT**

```md
> **Session/result READ-MODEL — this doc is not the home for it.** Two things a reader looking for "how does a
> played session render?" will not find here. (1) The **player** result page `/sim/<slug>/result/<sessionId>` is a
> **persisted read**, not a live recompute — `internal/graph/queries.resolvers.go:70` does plain Ent SELECTs over
> `validation_attempt_results`, so a seeded result fan-out renders a full result. (2) The **manager** view
> reads the **same** table — **the mirrors are GONE.** `app/terraform/migrations/20260729133514.sql:58-62`
> (*"5. Drop the mirrors."*) **re-points** `organization_assignment_sessions`' two foreign keys off the mirror
> ids (`:15-23`), NULLs the orphans (`:36-44`), then `DROP TABLE`s both `local_jobsimulation_sessions` and
> `local_skill_path_sessions` (`:62-63`), and `intelligence.go:1700` now reads `m.ent.JobSimulationSession.Query()`.
> **No session row is back-filled — the file contains 0 `INSERT`s** (this said *"back-fills then DROPs"* until
> M257x iter-52).
> **There is one row to seed, not a pair** — the older "seed the mirror or the scoreboard is blank"
> guidance is superseded. Full route-by-route treatment lives in
> [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md); the write side is
> [`../ops/demo/session-clone-spec.md`](../ops/demo/session-clone-spec.md).
```

**CITED CONTENT**

```
    55  ALTER TABLE "personal_assignment_sessions"
    56    ADD CONSTRAINT "personal_assignment_sessions_skill_path_sessions_session" FOREIGN KEY ("session_id") REFERENCES "skill_path_sessions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
    57  
    58  -- 5. Drop the mirrors. Dropping local_jobsimulation_sessions also drops its
    59  --    last-activity trigger; the trigger function (created manually in prod, not via
    60  --    migrations) is cleaned up defensively — the app now maintains
    61  --    memberships.last_activity_date itself on session events.
    62  DROP TABLE "local_jobsimulation_sessions";
    63  DROP TABLE "local_skill_path_sessions";
    64  DROP FUNCTION IF EXISTS on_insert_local_jobsimulation_sessions_update_memberships() CASCADE;
    65  
```

## 04-033
- **id**: `B04-033`
- **corpus site**: `corpus/services/jobsimulation.md:147-150` (paragraph)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> These are the edges the engine has, **not** a reading of a compose block: there is no
> `jobsimulation` service and therefore no `depends_on` list to quote. They are satisfied in-process inside
> `backend` — with one exception on the RPC axis: the only cross-process **Connect-RPC** edge out of
> `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). All at platform `0c91421`; the correctly-scoped model wording is [`architecture_overview.md:335`](../architecture/architecture_overview.md).
```

**CITED CONTENT**

```
    45        - .env
    46      environment:
    47        - AI_USAGE_STREAM=AI
    48        - AUTHORIZATION_ADDRESS=http://sentinel:8087
    49        - AWS_CHIME_SDK_REGION=eu-central-1
    50        - CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo
    51        - CMS_STREAM=cms
```

## 04-034
- **id**: `B04-034`
- **corpus site**: `corpus/services/jobsimulation.md:147-150` (paragraph)
- **citation**: `docker-compose.yml:57`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> These are the edges the engine has, **not** a reading of a compose block: there is no
> `jobsimulation` service and therefore no `depends_on` list to quote. They are satisfied in-process inside
> `backend` — with one exception on the RPC axis: the only cross-process **Connect-RPC** edge out of
> `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). All at platform `0c91421`; the correctly-scoped model wording is [`architecture_overview.md:335`](../architecture/architecture_overview.md).
```

**CITED CONTENT**

```
    54        - ELEVENLABS_EU_TEMPLATE_AGENT_ID=agent_4301k834j6pxfefbgf6bg48g8kpq
    55        - ELEVENLABS_TEMPLATE_AGENT_ID=agent_01k07b5k4ge3f9cvv30rv1d49n
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
    60        - LIVEKIT_AWS_SDK_REGION=eu-central-1
```

## 04-035
- **id**: `B04-035`
- **corpus site**: `corpus/services/jobsimulation.md:147-150` (paragraph)
- **citation**: `docker-compose.yml:183`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> These are the edges the engine has, **not** a reading of a compose block: there is no
> `jobsimulation` service and therefore no `depends_on` list to quote. They are satisfied in-process inside
> `backend` — with one exception on the RPC axis: the only cross-process **Connect-RPC** edge out of
> `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). All at platform `0c91421`; the correctly-scoped model wording is [`architecture_overview.md:335`](../architecture/architecture_overview.md).
```

**CITED CONTENT**

```
   180        - "3200:3200"
   181      networks:
   182        - app-network
   183      profiles: [core, backend, all]
   184  
   185  networks:
   186    app-network:
```

## 04-036
- **id**: `B04-036`
- **corpus site**: `corpus/services/jobsimulation.md:147-150` (paragraph)
- **citation**: `app/internal/converter/gotenberg.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/converter/gotenberg.go`  (54 lines)

**CLAIMING UNIT**

```md
> These are the edges the engine has, **not** a reading of a compose block: there is no
> `jobsimulation` service and therefore no `depends_on` list to quote. They are satisfied in-process inside
> `backend` — with one exception on the RPC axis: the only cross-process **Connect-RPC** edge out of
> `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). All at platform `0c91421`; the correctly-scoped model wording is [`architecture_overview.md:335`](../architecture/architecture_overview.md).
```

**CITED CONTENT**

```
    28  		return nil, fmt.Errorf("gotenberg: can't finalize multipart body: %w", err)
    29  	}
    30  
    31  	req, err := http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)
    32  	if err != nil {
    33  		return nil, fmt.Errorf("gotenberg: can't create request: %w", err)
    34  	}
```

## 04-037
- **id**: `B04-037`
- **corpus site**: `corpus/services/jobsimulation.md:147-150` (paragraph)
- **citation**: `docker-compose.yml:59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> These are the edges the engine has, **not** a reading of a compose block: there is no
> `jobsimulation` service and therefore no `depends_on` list to quote. They are satisfied in-process inside
> `backend` — with one exception on the RPC axis: the only cross-process **Connect-RPC** edge out of
> `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). All at platform `0c91421`; the correctly-scoped model wording is [`architecture_overview.md:335`](../architecture/architecture_overview.md).
```

**CITED CONTENT**

```
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
    60        - LIVEKIT_AWS_SDK_REGION=eu-central-1
    61        - LIVEKIT_HOST_URL=wss://anthropos-pbvktu3v.livekit.cloud
    62        - LIVEKIT_RECORDINGS_BUCKET_NAME=anthropos-livekit-test
```

## 04-038
- **id**: `B04-038`
- **corpus site**: `corpus/services/jobsimulation.md:153-153` (bullet)
- **citation**: `app/main.go:980-982`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* **CMS** — simulation definitions, content, studio entities. **The engine holds no `DIRECTUS_BASE_ADDR`/`DIRECTUS_TOKEN` of its own**; it calls the cms domain **in-process** (same binary, no RPC hop). There is no husk container on either end of that edge any more, and no variable either: `CMS_RPC_ADDR` was read only by `messenger`, pointed at `http://backend:8083` by `d11a403` — **one of the MIDDLE TWO that commit moved, not one of four** (`BACKEND_USERS_RPC_ADDR` and `SKILLER_RPC_ADDR` already held that value at `d11a403^`) — and removed outright with the messenger block at `838d907`. **The M23 content cutover does NOT ride on a `cms` container.** `backend` is the in-process Directus reader (`app/cms_reader_switch.go`; `app/main.go:980-982` @ `app` `b948604` v1.366.0 `log.Fatalf`s without `DIRECTUS_BASE_ADDR`), so re-pointing `cms` alone leaves `backend` reading prod — measured live on `demo-1` at M257x iter-24 as 96 Directus log lines, all 403. rext therefore sets `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` in both twins. No jobsimulation env change is needed, but the cutover must include `backend`.
```

**CITED CONTENT**

```
   977  					cbHandler.SetEnqueuer(workerClient)
   978  				}
   979  				cbHandler.SetUserResolver(authnManager)
   980  				courseBuilderDeps = backend.CourseBuilderDeps{
   981  					Service:       cbSvc,
   982  					Publisher:     cbPublisher,
   983  					AssetUploader: cbAssetSink,
   984  					Notifier:      cbNotifier,
   985  					Handler:       cbHandler,
```

## 04-039
- **id**: `B04-039`
- **corpus site**: `corpus/services/jobsimulation.md:164-164` (bullet)
- **citation**: `app/internal/cms/directus/collections/jobsimulation.go:1594-1597`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/directus/collections/jobsimulation.go`  (1742 lines)

**CLAIMING UNIT**

```md
* **ElevenLabs** — voice agents still used in the call/reply pipeline (`ELEVENLABS_TEMPLATE_AGENT_ID`, `ELEVENLABS_EU_TEMPLATE_AGENT_ID`). Engine choice is per sequence, from the CMS `voice_engine` field, not from a flag; **when that field is nil the default is `gptrealtime`** — `app/internal/cms/directus/collections/jobsimulation.go:1594-1597` **@ `app` `ad9f3c49`** (`func voiceEngineFromDirectus` at `:1594`; `:1595-1596` is the nil branch) — not ElevenLabs. ⚠️ **This cited `cms/directus/collections/…` until M257x iter-115, and that path exists in no clone**: the frozen `cms` repo @ `ca50c817` has no `directus/` at its root (`git ls-files | grep -c '^directus/'` = 0); its own copy is `cms/internal/directus/collections/jobsimulation.go`. The corpus spells this path correctly at **eight** other sites and this was the sole outlier — and it truncated toward the *decommissioned* repo, which [`external_services.md`](../architecture/external_services.md) explicitly warns against (*"`app/internal/cms/directus/` (NOT the frozen cms repo's `internal/directus/`)"*). Its same-fact twin is the *Legacy / Transitioning Engines* paragraph in [`ai_architecture.md`](../architecture/ai_architecture.md), which pinned the same construct as `:1594-1600`; both halves were corrected together, to the same path and the same range
```

**CITED CONTENT**

```
  1591  	return skills
  1592  }
  1593  
  1594  func voiceEngineFromDirectus(directusVoiceEngine *SimulationVoiceEngine) simulation.SimulationVoiceEngine {
  1595  	if directusVoiceEngine == nil {
  1596  		return simulation.SimulationVoiceEngineGptrealtime
  1597  	}
  1598  
  1599  	switch *directusVoiceEngine {
  1600  	case SimulationVoiceEngineGptrealtime:
```

## 04-040
- **id**: `B04-040`
- **corpus site**: `corpus/services/jobsimulation.md:206-215` (paragraph)
- **citation**: `app/main.go:216`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> ⚠️ **The signature did NOT survive the fold — do not go looking for a help block in `backend`.** `app` has
> **no cobra root command**: `app/main.go:216` (@ `app` `7177374` — a **pin**: that was origin/main on
> 2026-08-04 and is 38 commits back now; identical at `9d00a313` v1.367.0; `:212` @ the older `b948604` v1.366.0; and **`:229` at today's origin/main, `ad9f3c49`** — re-derived 2026-08-06, so do not read the pinned line offset as current) is a plain
> `func main()`, and the only `spf13/cobra` import in the whole repo is `cmd/createTaxonomy/main.go`. There is
> no `RunE`, so there is nothing to print `Error: …` and nothing to print a usage block. A failed init in
> `backend` is a single stdlib `log.Fatalf` line — timestamped, no `Error:` prefix, no help — and the container
> exits 1. The jobsim wiring is fatal by design — `jobsimwiring.Wire` at **`app/main.go:721`**, its
> `log.Fatalf("jobsim-in-app: engine wiring failed …")` at **`:723`**, measured @ `app` **`ad9f3c49`**
> (`:614` @ `b948604`). **This line cited an unpinned `:670` until M257x iter-108**; at `ad9f3c49` that
> line is `skillerAzureEndpointEu = &v`, an unrelated Azure-endpoint assignment.
```

**CITED CONTENT**

```
   213  		return dsn
   214  	}
   215  	ms := strconv.FormatInt(d.Milliseconds(), 10)
   216  	if strings.Contains(dsn, "://") {
   217  		u, err := url.Parse(dsn)
   218  		if err != nil {
   219  			return dsn
```

## 04-041
- **id**: `B04-041`
- **corpus site**: `corpus/services/jobsimulation.md:206-215` (paragraph)
- **citation**: `app/main.go:721`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> ⚠️ **The signature did NOT survive the fold — do not go looking for a help block in `backend`.** `app` has
> **no cobra root command**: `app/main.go:216` (@ `app` `7177374` — a **pin**: that was origin/main on
> 2026-08-04 and is 38 commits back now; identical at `9d00a313` v1.367.0; `:212` @ the older `b948604` v1.366.0; and **`:229` at today's origin/main, `ad9f3c49`** — re-derived 2026-08-06, so do not read the pinned line offset as current) is a plain
> `func main()`, and the only `spf13/cobra` import in the whole repo is `cmd/createTaxonomy/main.go`. There is
> no `RunE`, so there is nothing to print `Error: …` and nothing to print a usage block. A failed init in
> `backend` is a single stdlib `log.Fatalf` line — timestamped, no `Error:` prefix, no help — and the container
> exits 1. The jobsim wiring is fatal by design — `jobsimwiring.Wire` at **`app/main.go:721`**, its
> `log.Fatalf("jobsim-in-app: engine wiring failed …")` at **`:723`**, measured @ `app` **`ad9f3c49`**
> (`:614` @ `b948604`). **This line cited an unpinned `:670` until M257x iter-108**; at `ad9f3c49` that
> line is `skillerAzureEndpointEu = &v`, an unrelated Azure-endpoint assignment.
```

**CITED CONTENT**

```
   718  	// jobsim now, so wiring is FATAL, not best-effort: a jobsim-less boot must fail loud, never silently
   719  	// serve a half-wired domain. The GraphQL Session type, the Redis-stream subscribers, and the jobsim
   720  	// Asynq pools are all served by app unconditionally (no dormant gate).
   721  	jobsimDj, err := jobsimwiring.Wire(serverContext, logger, serviceName, ent, pub, redisClientStream, cmsReaderSw, posthogClient, jobsimUsers, jobsimSkiller, copilotDB, authz, storageManager)
   722  	if err != nil {
   723  		log.Fatalf("jobsim-in-app: engine wiring failed (is jobsim env provisioned?): %v", err)
   724  	}
```

## 04-042
- **id**: `B04-042`
- **corpus site**: `corpus/services/jobsimulation.md:232-245` (paragraph)
- **citation**: `docker-compose.yml:91`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
**The bind survived the fold and moved onto `backend`.** `docker-compose.yml:91` binds
`$HOME/.aws/credentials:/root/.aws/credentials:ro` — the **only** AWS bind in the file — under `backend`'s
`volumes:` (`:90`), and compose's own comment says why (`:88-89`: *"jobsim-in-app's Chime/LiveKit recording
managers use the AWS SDK default credential chain — the mount the standalone jobsimulation container had."*).
Measured at platform `0dab54d`. **When the host path does not exist, Docker auto-creates it as an empty
DIRECTORY.** The container then sees a *directory* where a file belongs, and `aws-sdk-go-v2`'s
`config.LoadDefaultConfig()` **opens it successfully** (opening a directory succeeds!) before failing `EISDIR`
on the read — so it is *not* skipped as an unreadable file. In the standalone binary that error propagated out
of `ai.NewAIManager` → the root `RunE` → cobra's usage block → `exit 1`. **The CAUSE is inherited; the
SIGNATURE is not, and the container name is not the only thing that changed.** In `backend` the identical
`config.LoadDefaultConfig` failure comes out of `jsai.NewAIManager` (`app/internal/jobsimulation/ai/ai.go:90`,
`can't load AWS config: %w`), is returned unwrapped by `jobsimwiring.Wire`
(`app/internal/jobsimwiring/wiring.go:147-148`) and dies at `log.Fatalf` in `app/main.go:670` — one timestamped
line, no `Error:` prefix, no usage block (`app` `9d00a313` v1.367.0; the fatal is `:614` @ `b948604`).
```

**CITED CONTENT**

```
    88        # developer machine (ENVIRONMENT=development is what makes unset mean off).
    89        # Pinning them to `false` here would override .env and make opting in impossible
    90        # without editing this file. To exercise either one locally, set it in .env — and
    91        # know that messenger then attaches to the LIVE Redis consumer group and
    92        # customerio-sync writes real Brevo contacts.
    93        - SUPABASE_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    94        - COPILOT_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
```

## 04-043
- **id**: `B04-043`
- **corpus site**: `corpus/services/jobsimulation.md:232-245` (paragraph)
- **citation**: `app/internal/jobsimulation/ai/ai.go:90`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/ai/ai.go`  (355 lines)

**CLAIMING UNIT**

```md
**The bind survived the fold and moved onto `backend`.** `docker-compose.yml:91` binds
`$HOME/.aws/credentials:/root/.aws/credentials:ro` — the **only** AWS bind in the file — under `backend`'s
`volumes:` (`:90`), and compose's own comment says why (`:88-89`: *"jobsim-in-app's Chime/LiveKit recording
managers use the AWS SDK default credential chain — the mount the standalone jobsimulation container had."*).
Measured at platform `0dab54d`. **When the host path does not exist, Docker auto-creates it as an empty
DIRECTORY.** The container then sees a *directory* where a file belongs, and `aws-sdk-go-v2`'s
`config.LoadDefaultConfig()` **opens it successfully** (opening a directory succeeds!) before failing `EISDIR`
on the read — so it is *not* skipped as an unreadable file. In the standalone binary that error propagated out
of `ai.NewAIManager` → the root `RunE` → cobra's usage block → `exit 1`. **The CAUSE is inherited; the
SIGNATURE is not, and the container name is not the only thing that changed.** In `backend` the identical
`config.LoadDefaultConfig` failure comes out of `jsai.NewAIManager` (`app/internal/jobsimulation/ai/ai.go:90`,
`can't load AWS config: %w`), is returned unwrapped by `jobsimwiring.Wire`
(`app/internal/jobsimwiring/wiring.go:147-148`) and dies at `log.Fatalf` in `app/main.go:670` — one timestamped
line, no `Error:` prefix, no usage block (`app` `9d00a313` v1.367.0; the fatal is `:614` @ `b948604`).
```

**CITED CONTENT**

```
    87  		config.WithRegion("eu-west-1"),
    88  	)
    89  	if err != nil {
    90  		return nil, fmt.Errorf("can't load AWS config: %w", err)
    91  	}
    92  	anthropicClient, err := anthropic.NewAnthropic(&cfg, nil)
    93  	if err != nil {
```

## 04-044
- **id**: `B04-044`
- **corpus site**: `corpus/services/jobsimulation.md:232-245` (paragraph)
- **citation**: `app/internal/jobsimwiring/wiring.go:147-148`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
**The bind survived the fold and moved onto `backend`.** `docker-compose.yml:91` binds
`$HOME/.aws/credentials:/root/.aws/credentials:ro` — the **only** AWS bind in the file — under `backend`'s
`volumes:` (`:90`), and compose's own comment says why (`:88-89`: *"jobsim-in-app's Chime/LiveKit recording
managers use the AWS SDK default credential chain — the mount the standalone jobsimulation container had."*).
Measured at platform `0dab54d`. **When the host path does not exist, Docker auto-creates it as an empty
DIRECTORY.** The container then sees a *directory* where a file belongs, and `aws-sdk-go-v2`'s
`config.LoadDefaultConfig()` **opens it successfully** (opening a directory succeeds!) before failing `EISDIR`
on the read — so it is *not* skipped as an unreadable file. In the standalone binary that error propagated out
of `ai.NewAIManager` → the root `RunE` → cobra's usage block → `exit 1`. **The CAUSE is inherited; the
SIGNATURE is not, and the container name is not the only thing that changed.** In `backend` the identical
`config.LoadDefaultConfig` failure comes out of `jsai.NewAIManager` (`app/internal/jobsimulation/ai/ai.go:90`,
`can't load AWS config: %w`), is returned unwrapped by `jobsimwiring.Wire`
(`app/internal/jobsimwiring/wiring.go:147-148`) and dies at `log.Fatalf` in `app/main.go:670` — one timestamped
line, no `Error:` prefix, no usage block (`app` `9d00a313` v1.367.0; the fatal is `:614` @ `b948604`).
```

**CITED CONTENT**

```
   144  		getenv("AZURE_OPENAI_KEY_US"), getenv("AZURE_OPENAI_ENDPOINT_URL_US"),
   145  		getenv("OPENAI_KEY"), azureVoiceKey, azureVoiceEndpoint,
   146  		aiUsagePub, posthogClient)
   147  	if err != nil {
   148  		return nil, err
   149  	}
   150  	aiResultManager, err := jsai.NewAIManager(ctx,
   151  		getenv("AZURE_OPENAI_KEY_RESULTS"), getenv("AZURE_OPENAI_ENDPOINT_URL_RESULTS"),
```
