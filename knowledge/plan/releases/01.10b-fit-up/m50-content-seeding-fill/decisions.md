# M50 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## TOK-01: seed-fill the genuine empties, sweep-driven, re-seed-to-iterate — 2026-06-30

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Drive M50 by the M42 coverage protocol exactly as written: the primary metric is the
coverage-sweep pair `(failingSections, escapes)` for each vantage (employee=Maya, manager=Dan); the gate is
`(0,0)` BOTH vantages over a frontier-exhausted crawl on a COLD reset-to-seed demo. Each tik: **Phase A** sweep
the live demo-1 as the vantage hero → **Phase B** triage each failing section by the fix-surface routing table →
**Phase C** land the fix in rext `stack-seeding` (new seeder or backfill) + **re-seed demo-1** (the light
re-apply: re-run the seeder from the authoring copy against demo-1's offset DB, or re-pin demo-1's rext clone) →
**Phase D** re-sweep → **Phase E** close on whether the targeted cluster cleared. Author + commit all seeder
code in the rext authoring copy (`.agentspace/rosetta-extensions/stack-seeding`, its own git, tag per the
per-iter convention); the rosetta worktree carries the milestone records + the `Delivers→` doc updates
(`profile-completeness-spec.md`, `stories-spec.md`). Reserve the heavy COLD reset-to-seed (tear-down +
`up-injected.sh` rebuild, ~15-25 min) for the **exit-gate proof only** — be machine-aware (9 GiB Docker VM +
the dev stack co-resident; no concurrent heavy ops).

**Rationale:** The orchestration's critical constraint is to **re-diagnose on the FRESH demo-1** — the
annotation gaps were observed on the OLD stale (pre-M47/M48) demo and several may already render. The iter-01
re-diagnosis (spec-notes) confirms this split: the **genuine seed gaps** are (1) member `location` /
`last_activity_date` / `joined_at` (0/221 each), (2) spoken languages (0 rows across `world_languages` +
`membership_languages` + `user_languages` — needs a NEW `MemberLanguagesSeeder` + the `world_languages`
reference fill), (3) certifications roster-coverage (hero-only → 2/221 — the "Talent really low" gap), (4) Maya
XP (`user_experience_points` 0 DB-wide) + the `/profile/activities` skill-path-completed app-mirror
(`local_skill_path_sessions` 0 for Maya). The **likely-NOT-seed-gaps** (library skill-paths = 22 published
directus rows; 76 target-roles; 114 assignments — all HAVE backing data) are probably federation/frontend/the
demo-up-#7-abort artifact and must be CONFIRMED by the sweep before assuming a seed fix (the protocol's
diagnose-before-assume-fix-surface lesson). So the strategy is sweep-FIRST: let the rendered sweep tell us which
gaps actually surface, then fix the highest-leverage seed cluster per tik. Honour F2: EXTEND `HeroActivitySeeder`,
never duplicate it. The AI-keys policy (F7) is a DECISION deliverable (record in decisions.md + secrets-spec.md;
seeded CONTENT renders without live AI — academy AI chat is documented-as-absent, NOT a gate blocker).

**Strategy class:** new-direction

**Distance-to-gate context:** gate = `(failingSections=0, escapes=0)` both vantages on a COLD demo, frontier
exhausted. Baseline `(failing, escapes)` not yet measured — iter-02 runs the baseline employee + manager sweeps.
The seed-level gap inventory is known (4 genuine clusters above); the rendered-sweep failing-section count is
the real metric and is iter-02's first reading.

**Next-tik direction (iter-02):** Run the **baseline sweeps** — employee (Maya) then manager (Dan) — against
demo-1, frontier-exhausted (raise `COVERAGE_MAX_PAGES` until `cappedAtFrontier===false` before quoting). Record
`(failingSections, escapes)` + the per-page/per-section verdicts. Triage: confirm which of the 4 genuine seed
clusters surface as failing sections, and whether the "has-data" surfaces (library/target-roles/assignments)
render or fail. Then pick the highest-leverage cluster (most sections unblocked per fix) as iter-02's fix target
— the leading candidate is the **member-field backfills** (location/last_activity/joined_at) since they unblock
the `/enterprise/members` manager section directly + likely feed Talent/Growth aggregates, but defer the final
pick to the post-baseline triage.

