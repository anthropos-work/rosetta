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

## Open (resolve during build)
- Directus: snapshot-replay into per-stack store vs hard-block-and-skip for the demo MVP.
- Link `app/internal/bootstrap`/ent directly vs `go run` CLIs (the perf path) — confirm clean import w/o platform edit.
- Audit log home: per-stack DB table vs sidecar file.
- Backdating fidelity: which timestamps are settable via direct SQL vs ent-Immutable/DB-default.
