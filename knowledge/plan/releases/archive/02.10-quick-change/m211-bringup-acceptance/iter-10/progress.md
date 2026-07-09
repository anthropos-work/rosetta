**Type:** tik (truly-cold demo-up + root-cause the residual federation failure). Under TOK-01 move (4).

# iter-10 — tik progress

## Execution log
1. **Purged demo-1** (`rosetta-demo down 1 --purge`) — removed all demo-1 images (incl. the 7-day pre-merge
   `demo-1-next-web` + a vestigial `demo-1-skiller:injected`). Confirmed 0 demo-1 images + 0 containers.
2. **Truly-cold `/demo-up`** (reap-safe detached+poll, ~10 min): rebuilt all backend Go services + router +
   frontends + clerkenstein. **Autoverify GREEN** (backend /api/health 200, sentinel.casbin_rules=1150,
   directus 21 collections, all liveness+readiness). The demo STANDS UP.
3. **BUT the federation STILL failed** after the clean rebuild: `publicJobSimulations.skills` → backend **422**
   "Internal server error", `searchSimulations` → cms error. Backend logs: `Cannot query field "name" on type
   "Skill"` + a **panic** (nil-ptr `gqlauthz.go:80` — `opCtx.Operation` nil after the schema-validation reject).
   So NOT a stale IMAGE (image purge + rebuild didn't fix it) — something deeper.
4. **Ruled out the obvious:** the app clone (v1.334.1) generated code DOES define `Skill.name` (`graph.go`);
   the Dockerfile uses committed generated code (no codegen step); the authz-skip patch only touches
   `roles.go`. So the SOURCE is correct — yet the running binary rejects `Skill.name` + its build layers were
   CACHED (`docker image rm` doesn't clear the build cache).
5. **ROOT CAUSE (definitive):** the injected **build-scratch** `$STACK/clones/app` (what actually compiles) was
   at **v1.315.0** (pre-merge, **0** `Skill.name` refs) while the source clone is **v1.334.1**. The injection
   loop clones the scratch ONCE (`if [ ! -d $dst/.git ]`) + checks out the highest tag AT THAT TIME, and
   **never re-syncs** it; `rosetta-demo down --purge` removes images but NOT `$STACK/clones`. So every demo-up
   (even "cold") bakes a **stale pre-merge binary** → the merged Skill federation entity is absent → the
   router's `_entities(Skill.name)` 422s + panics → empty library/profile.
6. **FIXED (rext, authoring copy):** the injection loop now re-syncs the scratch to the source's current
   release tag on EVERY bring-up (`git fetch --tags "$src"` + `checkout -f "$tag"`) before re-applying the
   disarmed colony + demo-patches. + a contract fence. All 63 frontend-build tests GREEN; shellcheck clean.
   Committed rext `0593cff`, moved tag `quick-change-m211` onto it, re-pinned the consumption clone.
7. **VERIFIED the fix produces a fresh scratch:** ran the re-sync on the app scratch → **v1.315.0 → v1.334.1**,
   `Skill.name` generated-code refs **0 → 5**. A rebuild from this scratch will bake a correct binary.

## Re-measurement (gate)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| truly-cold demo-up STANDS UP | (image-warm before) | **YES** — autoverify GREEN on a purged+rebuilt stack |
| (e) M42 coverage | NOT MET | NOT MET — root cause (stale build-scratch) FIXED + verified; end-to-end coverage-green re-prove routed to iter-11 |
**Metric:** the gate-blocking root cause is now FIXED in tooling (verified: scratch re-syncs to the merged
version WITH `Skill.name`). Coverage-green requires a re-run of the demo-up under the fixed tooling (rebuild
all injected from the re-synced scratches) — iter-11.

## Close — 2026-07-08

**Outcome:** Root-caused + FIXED the residual federation failure: the injected **build-scratch** was pinned at
a pre-merge tag (**v1.315.0**, no `Skill.name` federation) and never re-synced — surviving even `down --purge`
— so every demo-up baked a stale binary. Landed the re-sync fix (fetch + `checkout -f` the source's current
tag each bring-up) + verified it (scratch v1.315→v1.334, `Skill.name` refs 0→5). The truly-cold demo-up itself
STANDS UP green (autoverify passed).
**Type:** tik
**Status:** closed-fixed-partial (root-cause tooling bug FIXED + verified; the end-to-end M42-coverage-green re-prove routed to iter-11)
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (the root cause is a Fate-1 tooling bug, now fixed — NOT a platform bug: the clones + composition are correct) — (5) cap-reached: n (2nd tik) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (image purge is NOT truly cold — `docker image rm` leaves the build cache + the build-scratch; a "cold" demo-up baked a stale binary from a v1.315.0 scratch), D2 (the injected build-scratch must re-sync to the source's current release tag every bring-up — the fix landed as rext 0593cff), D3 (the demo STANDS UP green on a clean rebuild — federation is the only residual, now root-caused)
**Side-deliverables:** rext `0593cff` (tag `quick-change-m211`, moved from 846dae9) — the build-scratch re-sync fix + contract fence.
**Routes carried forward (Fate-3 → iter-11):**
- **iter-11 = re-prove:** re-run the demo-up under the fixed tooling (re-syncs ALL injected scratches → rebuilds
  app/cms/jobsim/skillpath from v1.334.1) → confirm `searchSimulations`/`publicJobSimulations.skills` resolve →
  M42e (employee) coverage GREEN. Handler: `BRINGUP-M211-iter11-reprove-coverage`.
- **Demopatch drift (surfaced this iter, Fate-3):** the `next-web-studio-url` + `next-web-public-website-url`
  demopatches G2-REFUSED (the re-synced `urls.ts` sha ≠ the manifest's pinned pre/post hash) — non-fatal, but
  the Studio link + PUBLIC_WEBSITE_URL stay prod-baked (an M42m MANAGER prod-eject risk). The authz-skip +
  aireadiness patches (anchored at v1.295.0) will likely also drift now that the scratch re-syncs to v1.334.1.
  Route: re-pin the demopatch manifests to the re-synced (v2.1) source hashes. Handler:
  `TOOLING-M211-demopatch-repin` (needed for the manager gate; employee gate unaffected).
- iter-12+ = M42m manager coverage + v2.0 Playthroughs + cold `/dev-up` (unchanged from iter-09's routes).
**Lessons:** (1) `docker image rm` / `down --purge` does NOT make a build truly cold — the docker BUILD CACHE
and the persistent build-SCRATCH survive. A re-sync release (new source tags) needs the scratch re-synced, not
just the images purged. (2) When a "clean rebuild" still reproduces a bug, check what SOURCE actually compiles
(the build-scratch git tag), not just the top-level clone. (3) A federated `_entities(Field)` 422 with a
`Cannot query field` + an authz nil-panic = the subgraph binary predates the field — a build-provenance
(not schema/data) problem.
