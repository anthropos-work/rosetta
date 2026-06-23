# M34 — Decisions

Implementation decisions with rationale (recorded during build). Design-time decisions already live in the
spec ([`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md), D1–D16).

## D-M34-1 — The Persona model lands as a blueprint field now, not the full `stack.stories.yaml`
The vertical slice (one hero) needs only `personas: [{id, name, role, verified, self_eval_bias}]` on the
existing `StackSeed`. The full Stories & Heroes model (multi-org, the trio, vantage/trajectory, the cockpit)
is M35–M38 (spec D12/D13). Adding a minimal, validated, *additive* `personas` field (heroes ride on top of
the population; an empty list seeds exactly as before) ships M34 without pre-building M35's schema.

## D-M34-2 — The seed-side closure gene is a standalone measure (`measure-closure`), not a snapshot gene
The M23 snapshot closure runs inside `measure-snapshot`, which requires a captured manifest per surface. The
M34 seed-closure needs NO manifest (it reads only live seeded refs vs the live taxonomy). Folding it into the
snapshot family would force a manifest dependency that doesn't apply. So it's its own `MeasureSeedClosure` +
`datadna measure-closure --stack demo-N` subcommand, reusing the `FidelityProbe`/`PgFidelityProbe` live-DB
plumbing + the M23 closure pattern. Mirrors how `measure-snapshot` is separate from `measure`.

## D-M34-3 — `user_skill_evidences` is an Exec UPSERT, not a CopyRowsIdempotent
The fleet's `CopyRowsIdempotent` does ON CONFLICT on a single `id` column. The evidences UNIQUE is the
composite `(skill_id, user_id)`, so an `id`-based merge can't dedup a re-seed. The PersonaSeeder UPSERTs each
evidence via a per-row `Exec` (`INSERT … ON CONFLICT (skill_id, user_id) DO UPDATE`), mirroring the existing
`seedCasbinGrants` per-row Exec pattern. DO UPDATE keeps the row current on a re-seed (idempotent in effect)
and sets `user_level` (D4) on both paths.

## D-M34-4 — The hero rides on a population user index (bridge: personaUserIndex / personaIndexMap)
A hero IS one of the UsersSeeder rows (spec D2), not a separate user. `personaUserIndex` maps a persona id to
a stable index in [1, Size]; the UsersSeeder writes the hero's real name/email at that index, and the
PersonaSeeder verifies her chain against the SAME derived user uuid. Both seeders MUST agree on the index —
`personaIndexMap` is the shared bridge (first-declared wins a benign hash collision, which the small M34
roster avoids by construction).

## D-M34-5 — Integration test is opt-in (build tag + env), self-cleaning, non-colliding
The "Maya renders" acceptance has an AUTOMATED half (the chain materializes + the UI query paths return data
+ closure green) and a LIVE-browser half (the orchestrator's post-build step). The automated half is a
`//go:build integration` test gated on `STACKSEED_IT_DSN`, so `go test ./...` never touches a DB. It seeds
under a unique `m34-it` stack key (no collision with the orchestrator's demo-3 seed) and DELETES every row it
wrote (incl. the trigger-created `user_basic_info`/`user_params` children) so the stack is left as found.

## Items surfaced + their fate (three-fate rule)

- **`stack.stories.yaml` model, multi-org, hero trio, cockpit** → **Fate 2** (already planned): the M35–M38
  overviews own these (spec §6b phasing, D13). M34's `personas` field is the additive foundation; no new
  deferral written.
- **The dashboard surfaces** (`membership_skills`, tags, target-roles, feedback) → **Fate 2**: M36's scope
  (the org-aggregate dashboard). M34 is the skill-profile vertical slice only (operator-locked priority:
  profile > dashboard).
- **`validation_attempt_result_id` FK omission** → **Fate 1, landed**: caught by the integration test, fixed
  in the same section, unit-test-guarded. (A PR-review-class issue — not a deferral.)

## Build-caught bug (worth recording)

The PersonaSeeder originally omitted the **NOT-NULL `validation_attempt_result_id`** FK on
`validation_attempt_skill_results`. The hermetic unit tests passed (the fake conn doesn't enforce NOT-NULL),
but the **integration test against the real schema rejected the insert** — the exact value of an
against-real-schema test. Fixed the column + row append; the unit test now asserts the FK is non-empty so the
unit suite alone catches a regression. Lesson: a fake-conn unit test can't substitute for one against-real-DB
pass per schema-touching seeder.

## D-M34-6 — Two clamp branches are unreachable-by-construction (harden-recorded, not tested)

Hardening drove the seeders package to 96.6% statements. Two residual uncovered blocks are
**defensive dead-code**, not gaps:
- `seedVerifiedSkill`'s per-session `score > 100` / `comp < 1` / `comp > 100` clamps (`persona.go`
  143/149/152): with `anthropos ∈ [62,85]` and the fixed `verifiedSessionsPerSkill = 3`, the max
  per-session score is `85 + 8 = 93` and max comp is `85 + 6 - 2 = 89` — neither clamp can fire
  without fabricating impossible inputs. They guard a future where the constants change; a test
  that forced them would be a disguised line-bump (a shallow test), explicitly disallowed by the
  three-fate rule's "no shallow tests as a Fate-1 shortcut".
- `users.go` 116 — the casbin-grant error wrapper — is a one-line `return …, fmt.Errorf(…, err)`
  delegating to `seedCasbinGrants`, whose own error returns are covered by its dedicated tests.

Both kept as defensive guards; no compensating test owed (the behavior they'd guard is unreachable
or already tested upstream). Recorded so a future audit doesn't re-flag them as missing coverage.
