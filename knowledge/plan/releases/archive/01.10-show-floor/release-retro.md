# Release Retro: v1.1 "show floor"

**Shipped:** 2026-06-05 · **Milestones:** M3 · M4 · M5 · M6 · M7a · M7b · M7c · M8

## What v1.1 delivered
The **platform-operations extension framework**: disposable, Clerk-free, **production-safely-seeded** demo stacks
(and the generalization to dev), in a clean 2-repo constellation. v1.0 made the platform run *without* Clerk;
v1.1 makes it run *as a believable, populated demo world a stakeholder can log into* — reproducibly, on offset
ports, never touching production or the dev stack.

- **Consolidation (M4):** the standalone `clerkenstein` + `rosetta-demo` repos collapsed into one private
  `rosetta-extensions` monorepo via history-preserving `git subtree`; the old repos deleted; rosetta thinned to pointers.
- **The stack tooling (M3/M5/M6):** `demo-N` stacks (Clerkenstein-injected, offset ports), the extracted
  `stack-injection` (demo-ON/dev-OFF), the shared `stack-core` port-offset engine, and `dev-stack`.
- **The seeding stack (M7a/M7b/M7c):** a host Go module that seeds a stack by talking **directly to its Postgres**
  (`COPY`) behind a **3-layer production-isolation guard**; a **data-DNA** discipline that conformance-gates the
  seeders + detects schema drift; and a **fleet** of backdated-activity seeders driven to a coverage gate.
- **The product layer (M8):** the demo-env corpus family + recipes + presets + `/demo-seed` + the express-gate CI.

## P0/P1 incidents across the release (all caught + fixed)
- **The casbin `g2` arg-order bug** (M7a) — the model matches `g2(org, sub, role)`; the seeder wrote
  `(user, org, role)`, silently 403-ing every org-feature check. Caught by the live login→200 proof.
- **The missing global Sentinel policy** on demo stacks (M7a) — `migrate-demo.sh` never ran `init_policy.sql`,
  so `casbin_rules` was empty + nothing authorized. Caught by the same proof; fixed in the bring-up.
- **The ×100 port-offset collision** (M3 extended close) — the base `up` defaulted OFFSET=100, colliding demo-1
  storage with dev jobsimulation. Fixed → 10000.
- **The skillpath `UNIQUE(user, path, version)` collision** (M7c) — caught on the first live seed; fixed by
  indexing the path by session number.
- **A +1-depth path break** (M4) + **registry concurrent-corruption** (M3) + **2 harness bugs** (M7b
  introspect-load, M7c) — all caught by the verify gates / live runs.

## Cross-milestone patterns
- **The live proof was the MVP of the whole release.** Every milestone that touched the running platform
  (M3, M7a, M7b, M7c) caught a real bug that unit tests passed over — because the bugs lived in the seam between
  the seeder and the platform's *actual* schema/authz, which only a live stack exercises. The discipline of
  "prove it against demo-1, not just the fakes" paid for itself repeatedly.
- **Honest scoping held.** M6 (dev-stack scoped to proven value), M7a (perf path revised when ent-linking proved
  impossible), M7c (taxonomy + content waived rather than faked) — each resisted the temptation to over-build or
  fabricate. The data-DNA `waived` status makes the hard line machine-visible.
- **The alignment pattern generalized twice.** From behavioral (Clerk, v1.0) to deployment (M3) to **data**
  (M7b) — the same manifest/score/diff structure, new operators each time.

## Metrics delta
Test funcs 175 → **409** (+134%); 4 alignment gates held **100%/100%** throughout; the seeding stack added the
isolation + data-DNA + fleet gates; flake 0; supply chain clean. See [metrics.json](metrics.json).

## Carry-forward → v1.2 "richer demo worlds"
- **The snapshot mechanism** — lift M7c's `waived` surfaces: the skiller taxonomy node-hierarchy snapshot +
  Directus content snapshot-replay → data-DNA coverage to 100% of the full catalog. (roadmap-vision.md)
- **AI-generated rich content** — transcripts / AI-scored narratives / fresh embeddings, layered on the M7 seeding
  foundation (the v1.1 hard line excludes these by design).
- **External shareability** (Tailscale vs ingress) for customer-facing demos.
- **The deployment/injection CI gate** — currently a local gate (needs colony via GH_PAT); a future runner with
  org-secret access could wire it.
