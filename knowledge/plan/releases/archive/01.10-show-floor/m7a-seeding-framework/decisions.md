# M7a — Decisions

## M7a-D1 — split the old M7 into M7a/M7b/M7c (design, 2026-06-04, user)
The single `section` M7 "stack-seeding" was reshaped after the user asked to make seeding robust/resilient/
drift-proof/fast/**production-safe**. Redesign research (3 agents over the platform) grounded three findings:
the pollution boundary is small + enumerable; the alignment pattern extends to data (with new operators); the
perf bottleneck is DB-IO not CPU. Decision (user-confirmed): split into **M7a** (framework + safety, section),
**M7b** (data-DNA discipline, section), **M7c** (seeder fleet, iterative). All stay in **v1.1** (user chose
"keep all of it in v1.1" over spinning the robust seeding into v1.2). The old `demo.seed.yaml` design is not
discarded — it folds into M7a (schema/safety/proof) + M7c (the fleet).

## M7a-D2 — the isolation guard is the load-bearing deliverable (design, 2026-06-04)
The single un-guarded shared-write that reaches prod is the failure this milestone exists to prevent. The guard
is therefore a hard, tested boundary (block + audit-log + clean-audit assertion), not a convention. Carries the
M4-D1 hedge: `--validate/--dry-run` fold into the seeder; a standalone validator splits out only if it grows
past a flag.

## M7a-D3 — perf path is DIRECT POSTGRES (COPY/SQL), not ent-linking nor CLI-shelling (build spike, 2026-06-04)
The build-feasibility spike resolved the overview's "confirm the import path early" open question in the
**negative**: `app/internal/bootstrap` (and the ent client) are **`internal/` packages** → Go's internal-import
rule forbids any module outside `github.com/anthropos-work/app` from importing them. So "link the ent client into
the seeder" is impossible without placing seeder code inside the platform tree (violates zero-platform-change).
And shelling the `bootstrap-*` CLIs per-row is the slow bottleneck the research identified. **Decision:** the
seeder is a **host Go module that connects directly to the target stack's Postgres** (exposed on the stack's
offset port — dev `:5432`, demo-N `:5432+N·10000`) and uses **`COPY`/batched SQL** for high-volume surfaces +
direct SQL for side-effecting primitives (the Clerkenstein identity, Sentinel/casbin). The drift risk this
creates (hand-written SQL vs the live schema) is **mitigated by M7b's data-DNA** (schema introspection), not
hand-maintenance — a clean synthesis: the fast path needs the conformance gate, and M7b is that gate.
Spike evidence: `app/internal/bootstrap` is internal; `anthropos-postgresql-1` maps `0.0.0.0:5432`; app data in
schema `public` (`users`/`organizations`/`memberships`), with `auth`/`sentinel`/`skiller`/`skillpath`/`jobsimulation`/`cms` siblings.

## M7a-D4 — the live login→200 proof caught two real bugs (proof, 2026-06-04)
The user chose the FULL proof: bring up an injected `demo-1` stack, seed it, prove a real login→200. The arc is
now reproducible end-to-end (no manual SQL): **no token → "unknown viewer"; token pre-seed → 403; token
post-seed → 200** (`membershipsCount` = 1001 = 1000 bulk + the identity). It surfaced two bugs unit tests could
not:
1. **The casbin `g2` arg order was swapped.** Sentinel's model (`sentinel/internal/authorization/casbin.go`)
   matches `g2(org, sub, role)`, so the stored row is `(org, user, role)` — `v0=org, v1=user`. `identity.go`
   wrote `(user, org, role)`, which made *every* org-feature check silently `403`. Fixed in `seedCasbinGrant` +
   pinned by the regression assertion in `TestIdentitySeeder_ResolvesPluralCasbinTable`. (The M3 Fate-3 note had
   the order backwards — corrected.)
2. **The global Sentinel policy was never bootstrapped on demo stacks.** The platform's `make migrate` pipes
   `sentinel/init_policy.sql` (the ~47 global role→feature `p`/`p2`/`p3`/`p5` rows) into the DB; `migrate-demo.sh`
   did not, so demo stacks had an EMPTY `casbin_rules` and no role mapped to any feature → 403 for everyone. This
   is **platform bootstrap (identical across stacks), not demo data**, so the fix belongs in `migrate-demo.sh`
   (a demo-stack fix made during M7a, Fate-1 — it blocked M7a's proof), applied only when `casbin_rules` is empty
   (init_policy.sql is not idempotent). The seeder owns the per-org/per-user demo data on top of it.

## Open (resolve during build)
- Directus: snapshot-replay into per-stack store vs hard-block-and-skip for the demo MVP.
- Audit log home: per-stack DB table vs sidecar file.
- Backdating fidelity: which timestamps are settable via direct SQL vs ent-Immutable/DB-default.
- The exact per-stack DSN source (platform/.env var name vs derive from the stack's offset port).
