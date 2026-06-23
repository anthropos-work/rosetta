# M34 — Retro

_Verified-skill chain (vertical slice) — turn the placeholder seeder's hollow world into one real hero (Maya).
The first `section` milestone of v1.9 "storytelling". Two-repo (rosetta doc-half + planning; ext code @ tag
`storytelling-m34`). Closed 2026-06-23._

## Summary

M34 made **one hero (Maya Chen) real, end-to-end** — the verified-skill **spine** the rest of v1.9 builds on.
Before M34 the seeder produced a structurally-correct-but-hollow world: every user "User N", binary 85/35
scores, **zero** verified skills (the core profile surface rendered empty), and — recon finding **G14** — a
session seeder that wrote invalid free-text enum/result values so its sessions were *inserted-but-invisible*
dead rows. M34 shipped five pieces in `rosetta-extensions/stack-seeding`:

- **The `PersonaSeeder`** — the verified-skill **7-table fan-out** per (hero × skill) across three Postgres
  schemas (`jobsimulation.sessions` → `validation_attempt_results` → `validation_attempt_skill_results` →
  `validation_criterion_results` → `public.local_jobsimulation_sessions` → `user_skills` →
  `user_skill_evidences` UPSERT). It honours every constraint the running platform enforces (the
  `user_skills_check_foreign_keys` CHECK via `job_simulation_id`, the partial-UNIQUE distinct-sim-per-row, the
  composite-UNIQUE evidences UPSERT), and writes the two fields the reference `seed.sql` omits — `result_status`
  and the load-bearing **`user_level`** (without which the claimed-vs-verified widget is empty).
- **The `TaxonomyRefs` resolver** — draws every skill node-id from the **real replayed public skiller
  taxonomy** (role-coherent via `skillsByRole`, is_core-first), with the load-bearing **empty-pool skip — never
  fabricate** a node-id.
- **The G14 `jobsim_sessions.go` fix** — valid `status='ended'` / `completion_status∈{passed,failed}` /
  `result_status='completed'` / a `^[a-z0-9]{5,10}$` token / the full `SIMULATION_TYPE_*` proto strings, plus a
  continuous mid-skewed score with a per-user upward growth arc and an ASSESSMENT/HIRING share.
- **The `users.go` patch** — real names / deterministic avatars / org-domain emails; a hero's name + email land
  at the population index the `PersonaSeeder` verifies her chain against (the shared `personaUserIndex` bridge).
- **The seed-side closure gene** (`datadna measure-closure`) — proves 0 dangling seeded skill node-ids against
  the replayed taxonomy (mirrors the M23 cross-surface gene). "Believable" is **measured, not assumed.**

The corpus doc-half graduated the gitignored analysis-of-record into a NEW `corpus/ops/demo/stories-spec.md`
(the verified-skill-chain reference) + `seeding-spec.md` / `safety.md` / `demo/README.md` / `CLAUDE.md` updates.
**Maya proven** via the `//go:build integration` test (every UI query path returns her data) and
orchestrator-verified on the live `--local-content` demo-3 stack.

## Incidents This Cycle

- **P1 (build-caught, fixed in-section) — the missing NOT-NULL `validation_attempt_result_id` FK.** The
  `PersonaSeeder` originally omitted the FK on `validation_attempt_skill_results`. The **hermetic unit tests
  passed** (the fake conn doesn't enforce NOT-NULL) but the **integration test against the real schema rejected
  the insert** — the exact value of an against-real-DB pass per schema-touching seeder. Fixed the column + row
  append; the unit test now asserts the FK is non-empty (`persona_test.go:181-196`) so the unit suite alone
  catches a regression. **Lesson:** a fake-conn unit test can't substitute for one against-real-schema pass per
  schema-touching seeder. (decisions.md "Build-caught bug")
- **P3 (close-caught, doc-accuracy) — two handbook test-count drifts + one comment-vs-behavior mismatch.** The
  `stack-seeding`/`dna` README Status sections quoted stale test counts (62/49/10) that had drifted across many
  milestones; the `taxonomyref.go` `take()` comment claimed a flat-pool top-up the caller doesn't do. All three
  fixed Fate-1 at close (the comment corrected to match the deliberate role-coherence-over-count behavior).
- **0 flakes, 0 regressions.** The suite is deterministic by construction (seeded hashes, no time/random); the
  flake gate ran 5/5 clean.

## What Went Well

- **The "port, don't reinvent" discipline paid off.** The chain was ported from the proven
  `/seed-verified-skill` `seed.sql` rather than re-derived — the only real bug (the FK omission) was a *missing*
  column from the port, caught immediately by the integration test, not a logic error.
- **The vertical-slice-first shape held.** One hero, one role, the full chain — proving the spine end-to-end
  before M35 scales to the multi-org roster. The `personas` blueprint field is purely **additive** (an empty
  list seeds exactly as before), so M34 ships zero regression risk to the existing population path.
- **The integration test earned its keep.** It caught the FK bug the hermetic units missed, and it's the
  automated half of the "Maya renders" acceptance (the live browser half is correctly M37/M38-owned).
- **Closure is measured.** The seed-side closure gene makes "real skill node-ids, never fabricated" a checked
  invariant (0 dangling), not a hope.

## What Didn't

- **Handbook test counts drifted silently across milestones** until the close reconciliation caught them — the
  62→381 / 49→117 gaps had accumulated over several milestones with no per-close reconciliation. Now reconciled;
  the close-milestone Phase 4-step-6 contract is the guard going forward.
- **Two roster-scaling edges were left as M34-benign-by-construction** (the index-collision guard + the
  short-role-pool top-up). Correct for a 1-hero slice, but they're real once M35 adds the trio — routed Fate-3 to
  M35 (D-M34-7) rather than over-building M34.

## Carried Forward

- **D-M34-7 → M35** (annotated in M35's `overview.md` In: list): add a `len(Personas) <= Size` blueprint
  validation + an index-collision warning, and decide whether a short-but-nonempty role pool tops up from the
  flat pool to hit each hero's declared `verified: N` (M35's roster-fidelity product call).
- **The Out-list scope (Fate 2, already-owned):** multi-org + `stories.yaml` + the hero trio → M35; the
  org-aggregate dashboard surfaces → M36; the Clerkenstein multi-identity + presenter cockpit → M37/M38. The
  literal browser-pixels render of Maya's *individual* profile (login-AS-Maya) is M37/M38's "login as a hero".

## Metrics Delta (from metrics.json)

- **stack-seeding Go test funcs:** 259 → **302** (+43; **381** incl. subtests). Module total 1027 → 1070.
- **Coverage:** seeders 93.7% → **96.6%** (+2.9, 2 harden passes); `dna/seed_closure.go` **100%**.
- **Flake:** **0** (flake gate 5/5 sequential shuffled `-race`).
- **Supply-chain:** GREEN (0 new deps). **Alignment:** 100%/100% (untouched).
- **Close review:** 5 findings, 0 blocking, 0 must-fix. **Deferral re-audit:** GREEN (0 deferrals).
