# M7c — Decisions

## M7c-D1 — iterative shape, gated on data-DNA coverage (design, 2026-06-04)
The seeder fleet is the open-ended part of the seeding redesign: the surfaces' *acceptance* is fixed (M7b's
conformance genes), but which surfaces carry the believable-demo weight + how scenarios compose + where schema
reality fights backdating only emerges from building. So M7c is **iterative** (like M1/M2c), gated on a measured
score — here **data-DNA coverage** — not a fixed checklist of seeder internals. The believable-demo *subset* of
surfaces is the real target, not every one of the ~30 entity types.

## TOK-01 — bootstrap strategy (iter-01, 2026-06-04)
**Initial strategy:** sequence the fleet by **reachability** (surveyed live against demo-1; see iter-01/overview.md).
The **believability core is reachable without solving Directus** — the activity surfaces (`jobsim-sessions`,
`skillpath-sessions`, `assignments`, `activity`) have no DB-enforced content FKs (`sim_id`/`skill_path_id`/
`resource_id` are free values), so they seed now. Order: **jobsim-sessions → activity (depends on sessions) ∥
skillpath-sessions ∥ assignments**, each with the **backdated, time-distributed** generator (the
`seed-verified-skill` pattern). Two surfaces are **snapshot-blocked** and routed to a tail iter / waiver:
`taxonomy` (needs the pre-embedded skiller node-hierarchy snapshot — empty in demo-1) and `content` (the shared
Directus instance — snapshot-replay only; the isolation guard blocks live writes).
**Coverage arithmetic:** seeding the 4 activity surfaces takes data-DNA coverage 4/10 → **8/10 = 80%**. The 90%
gate requires taxonomy or content too; per the milestone's **Re-scope trigger**, if those stay snapshot-blocked
they're **waived with justification** (the hard line: M7c seeds structural data, it does not author the taxonomy
or write shared Directus) and the gate is measured over the **reachable** subset. Confirm the waiver with the
user at close.
**Concrete iter-01 deliverable:** corrected the catalog's guessed table names (skillpath-sessions/assignments/
activity) in `data-dna.json`.

## iter-02 — the believability-core pattern + a harness fix (2026-06-04)
The `jobsim-sessions` seeder establishes the **reusable activity-seeder pattern** for the remaining tractable
surfaces (skillpath-sessions, assignments, activity): bulk-COPY backdated rows whose timestamps are
deterministically distributed across the blueprint's `activity` span, with a pass/fail mix per `pass_rate`,
content references (`sim_id` etc.) as free values (no FK). It promotes the surface planned→seeded in the
data-DNA, then `introspect` captures its shape and `measure` gates conformance — coverage rose 40%→50%.
**Harness fix (iter find):** `datadna introspect` validated the DNA *before* loading, so a freshly-promoted
seeded surface (empty shape, about-to-be-captured) failed to load. Fixed: `introspect` now loads without
Validate (it's the command that makes the DNA valid); `measure`/`diff` still validate. Also updated
`TestShippedManifest` to be M7c-aware (the 4 core surfaces stay critical+seeded; promoted surfaces are
non-critical with a non-empty shape; the stable invariant is the total catalog size, not a fixed planned count).

## iter-03/04/05 — the believability core complete + the re-scope point (2026-06-04)
Three more activity seeders landed (skillpath-sessions, assignments, activity), all on the iter-02 pattern →
**coverage 50%→80% (8/10)**, `measure` Overall/Critical 100%, the full 8-seeder seed runs in **0.69s**. Iter find:
the skillpath table has a UNIQUE `(user_id, skill_path_id, version)` — the seeder collided on its first live run;
fixed by indexing `skill_path_id` by session number (a user's paths are distinct). assignments/activity carry
real FKs (memberships, sessions) all referentially valid.
**Gate status: 3 of 4 met** — (a) login 200 ✓, (c) <2min (0.69s) ✓, (d) zero shared writes ✓; (b) coverage is
**80%**, short of the 90% threshold because the last 2 surfaces (taxonomy, content) are **snapshot-blocked**.
**Re-scope decision (the Re-scope-trigger fires here):** taxonomy needs the pre-embedded skiller snapshot;
content is the shared Directus (snapshot-replay only). Both are the **hard line** — M7c seeds structural data, it
does not author the 60K-skill taxonomy nor write shared Directus. Options surfaced to the user: **(A) waive the
two** and close at 80%/100%-over-reachable (the believable-demo subset is complete), or **(B) keep M7c open** to
solve the snapshot (heavy; arguably v1.2). Awaiting the user's call.

## Open (resolve during iters)
- The must-cover surface subset that makes a demo "believable" (sets which genes are critical).
- Per-surface backdating fidelity (direct-SQL vs ent-Immutable/DB-default).
- Whether the 2–3 seed presets ship in M7c or route to M8 (recipes/discoverability).
