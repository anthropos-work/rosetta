# M7c — Spec notes

Iteration-protocol-specific technical notes accumulate here across iters. Seeded from the 2026-06-04 redesign research.

## The fleet (one module per M7b-catalogued surface)
Users · Orgs/Memberships/Casbin · Features/Tiers · Taxonomy · Content/Library · SkillPath sessions · JobSim
sessions+results · Assignments · Activity. Each module: declares its isolation class (M7a guard), uses its
primitive (bulk-COPY where possible / RPC/CLI where side-effecting), and is measured by its M7b conformance gene.

## The backdated activity generator (the believability core)
Modeled on `seed-verified-skill`: per user, sample the content pack's **real** sim_ids; emit passed/failed
sessions distributed across the activity span; backdate `created_at` where the schema allows. Must cover the
FAILED-session validation-result chain, not just passed (confirmed gap in the old M7 notes).

## Perf at scale (inherited from M7a)
Bulk `COPY` for users/memberships/sessions/activity; goroutine fan-out for `JoinOrg`/Sentinel grants. The < 2 min
gate is the forcing function — profile each seeder; a seeder that loops `go run` per row fails the gate.

## Reusable references (do not reinvent)
- `bootstrap.JoinOrg` (per-user membership+casbin+feature), `seed-verified-skill` (time-distributed activity),
  the pre-embedded skiller snapshot (consume, don't generate).

## To confirm during iters
- The "believable demo" must-cover surface subset (drives which genes are critical).
- Per-surface backdating fidelity (direct-SQL vs ent-Immutable/DB-default).
- Whether 1k `bootstrap-user` via the linked ent client stays under the Clerkenstein mock's state limits.
- Whether the 2–3 presets ship here or route to M8.

## Running notes
_(append per-iter technical findings here)_
