**Type:** tik (M42e employee coverage — diagnose the skeleton failure). Under TOK-01 move (4).

# iter-09 — tik progress

## Execution log
1. **Re-surveyed the run-2 partial** M42e report: `reachable=7 failingSections=8 personaFailures=2 escapes=0`.
   Viewed the screenshots — the failing pages (library AI-sims, library skill-paths, profile, profile/skills)
   render **perpetual skeleton loaders**, not empty content. Lightweight/SSR data (home greeting, activity
   stats, org name+logo, menu avatar) renders fine; heavy client-side grid/profile queries hang.
2. **Verified the data + backend are healthy:** `public.skills`=42,790, `public.user_skills`=3,884, 601 users,
   5 orgs incl. Cervato Systems; the router serves `publicJobSimulations`/`libraryCategories`/sitemaps directly.
   So NOT a content/seed gap.
3. **HYPOTHESIS (initial): stale-offset next-web image.** The running container's RUNTIME env showed
   `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql` (non-offset). Landed a tooling hardening:
   `up-injected.sh`'s frontend image-reuse guard now validates the cached image's baked offset endpoint before
   reusing (rebuild-on-mismatch) — for both next-web + studio-desk. + 3 unit tests + a contract-fence upgrade
   (all 62 frontend-build tests GREEN; shellcheck clean). Committed rext `846dae9`, tagged `quick-change-m211`,
   re-pinned the stack-demo consumption clone + `.agentspace/rext.tag`.
4. **HYPOTHESIS FALSIFIED.** Rebuilt via the fixed guard → it reported the next-web image was **already at the
   correct offset** (`15050`) and reused it. Confirmed: the built client bundle (`/app/apps/web/.next`) contains
   `http://localhost:15050/graphql` (10×) — the client uses the CORRECT endpoint; the runtime-env 5050 is a
   compose override irrelevant to Next.js's build-time-inlined `NEXT_PUBLIC_*`. **The next-web image was never
   the cause.**
5. **REAL cause found (Playwright diag probe as Maya):** the library grid's `searchSimulations` query returns
   `Failed to fetch from Subgraph 'skiller' at Path 'searchSimulations.simulations.@.skills'`. **The router's
   composed supergraph still lists a dead `skiller` subgraph.** The `demo-1-graphql` router image is dated
   **2026-06-30** — a **pre-merge** build reused by compose `up -d` (no `--build`). The platform composition
   config (`supergraph-config-compose.yaml` + `subgraphs.conf`) correctly lists only 4 subgraphs (no skiller).
6. **Rebuilt the router** from the merged clone (`docker compose build graphql` + recreate). The **skiller
   subgraph error vanished** — `searchSimulations`/`publicJobSimulations.skills` now route to the `backend`
   subgraph (correct, post-merge). BUT the backend then **panics** (`gqlauthz.go:80` nil-ptr) + rejects
   `Skill.name` ("Cannot query field name on type Skill") + returns 422 → the merged app:injected image
   (`14:49`) is **also stale** (rejects a field the current clone defines).
7. **Classified as stale-image, NOT platform:** the app clone (**v1.334.1**, `skills.graphqls:12 name:
   String!`) AND the composition `backend.graphqls:3085` BOTH define `Skill.name`; the composition has **no**
   skiller subgraph. So the clones are correct + consistent — demo-1 is an **inconsistent mishmash of images
   built across the merge** (router 06-30, app 07-08 14:49) that the demo's image-reuse (`up -d` no `--build`
   + injected docker-cache) silently served. **iter-08's "cold /demo-up GREEN" was image-WARM** — it verified
   sub-condition (a) by CONTAINER COUNT only (no skiller container), missing the router's COMPOSED supergraph.

## Re-measurement (gate)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (a) 4-subgraph / no-skiller compose | claimed MET (iter-08, container-count) | **CORRECTED: demo-1 router had a live `skiller` subgraph** (stale image); router-rebuild fixed it on demo-1, but the honest gate needs a truly-cold rebuild |
| (e) M42 coverage + Playthroughs | NOT MET | NOT MET (real cause diagnosed; needs the truly-cold rebuild) |
**Metric:** unchanged (gate still 5/6 nominal) — but iter-09 **corrected a false-green**: the demo needs a
TRULY-cold rebuild (purge the stale pre-merge image mishmash + rebuild all from the merged clones) to be a valid
merged-platform substrate. Diagnosis complete; path proven (router rebuild dropped the skiller error).

## Close — 2026-07-08

**Outcome:** Diagnosed the M42e skeleton failure to its root: demo-1 runs on a **stale-image mishmash** spanning
the merge (pre-merge router with a dead `skiller` subgraph; stale app:injected rejecting `Skill.name`) — NOT a
platform bug (the clones + composition are correct + consistent). iter-08's "cold /demo-up GREEN" was
image-WARM (container-count-only (a) check). The honest fix is a **truly-cold demo-up** (purge + rebuild all).
Landed a real tooling hardening as a side-deliverable (the frontend image offset-validation guard).
**Type:** tik
**Status:** closed-no-lift (initial hypothesis falsified — next-web image was correct; real cause diagnosed + routed; a side-hardening landed)
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (the finding is a Fate-1 tooling/operational fix, not a platform blocker — the clones are correct) — (5) cap-reached: n (1st tik) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (next-web image was NOT the cause — client bundle correct at :15050; runtime-env 5050 is an irrelevant compose override), D2 (real cause = stale pre-merge router + backend images silently reused by `up -d`/docker-cache; a truly-cold rebuild is the honest gate proof), D3 (iter-08's cold demo-up was image-warm; sub-condition (a) must verify the COMPOSED supergraph / a real federated query, not just container count)
**Side-deliverables:** rext `846dae9` (tag `quick-change-m211`) — the frontend image-reuse guard now validates the baked offset endpoint before reusing (rebuild-on-mismatch, next-web + studio-desk) + 3 unit tests + fence upgrade. A genuine hardening against the stale-offset-image trap class (the same class that broke the router, in its own build path), even though demo-1's next-web image was not itself stale.
**Routes carried forward (Fate-3 → iter-10):**
- **iter-10 = the truly-cold demo-up:** `/demo-down --purge` demo-1 → truly-cold `/demo-up` (rebuild ALL images from the merged clones: router 4-subgraph + app:injected with Skill.name + frontends) → re-run M42e (employee) + M42m (manager) coverage → prove GREEN. Handler: `BRINGUP-M211-iter10-cold-demo`.
- **iter-11 = v2.0 Playthroughs** vs the clean demo-1. Handler: `BRINGUP-M211-iter11-playthroughs`.
- **iter-12 = cold `/dev-up`** (the dev half). Handler: `BRINGUP-M211-iter12-cold-dev`.
- **Doc/verify note (route into iter-10 or a cleanup):** `verification.md`/coverage sub-condition (a) should
  assert the COMPOSED supergraph has no skiller subgraph (a real `searchSimulations` federated query, or a
  supergraph-introspection check) — not just "no skiller container". And `rosetta_demo.md`: a re-synced
  (v2.1) platform requires a **cold** demo (purge stale pre-merge images) — `up -d` silently reuses them.
  Handler: `DOC-M211-supergraph-verify + re-sync-cold-purge`.
**Lessons:** (1) A Next.js `NEXT_PUBLIC_*` value is build-time-inlined into the client bundle — the container
RUNTIME env is a red herring; always check the built `.next` bundle, not `docker exec env`. (2) The demo's
image-reuse (`up -d` without `--build`; injected docker-cache) silently serves STALE images across a code
re-sync — a "cold" demo-up is only truly cold if the images are purged first. (3) Sub-condition (a)
"4-subgraph compose" must verify the COMPOSED SUPERGRAPH (a real federated query), not the container count —
a stale router can have a dead subgraph with the container correctly absent.
