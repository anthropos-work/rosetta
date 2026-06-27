# iter-07 — decisions

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| D1 | `cmd/gen-batch` gets a `--call-timeout` (default 60s) wrapping each `CompleteJSON` | The run loop used `context.Background()` with no deadline — a hung Azure endpoint stalls the whole batch (the watchdog-stall class that killed run-2). A per-call deadline fails fast. | 2026-06-26 |
| D2 | Intra-batch name de-dup = prompt diversity demand + a re-roll "avoid these names" hint (NOT just a bigger reroll budget) | gpt-4o-mini is strongly name-sticky per mother-prompt; a bare seed change re-picks the same name (rerolls=8 still left dups at higher cost). Telling the model which names are taken makes re-rolls actually diverge → 20/20 distinct at the LOWEST cost ($0.0059). | 2026-06-26 |
| D3 | Each generated member gets ONE company + ONE current-role `user_experiences` row; claimed skills tie to it via `user_skill_experience` | The DB CHECK `user_skills_check_foreign_keys` needs ≥1 provenance edge non-NULL; the generated claimed rows had all edges NULL (23514). Reusing the ProfileSeeder's company/experience helpers satisfies the constraint, stays SHALLOW (1 exp = current job, per the M45 scope boundary), and makes the profile believably show a current role. | 2026-06-26 |
| D4 | A reproducible `stackseed --cache-root` flag (threaded doSeed → buildRegistry → `GeneratedBatchSeeder.CacheRoot`) | The demo bring-up must point the generated-batch seeder at the captured prompt-hash cache; an ad-hoc default-path-only seeder isn't reproducible-via-mechanism. | 2026-06-26 |

**Falsification / corrections recorded:**
- The prior run-2's spec-notes carried aspirational N=20 numbers ($0.0033, 20 calls) written before any real
  run completed; this RETRY treated the gate as 0/5 UNPROVEN and produced the REAL measurement ($0.0059, 33
  calls — the extra are name-dedup re-rolls). The spec-notes are corrected to the verified numbers.
- A verification query joining `user_skills.skill_id` on `skiller.skills.id` (uuid) falsely read "47/47
  fabricated"; the resolvers use the `node_id` column (`K-PYTHON-8B21` = "Python"). Re-checked on `node_id`:
  0 fabricated, confirmed by `datadna measure-closure = PASS`.
