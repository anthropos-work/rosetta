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


## D-CLOSE-1 — AI-provider-keys policy: documented-as-absent (resolves inherited DEF-M49-01) — 2026-06-30

**Context:** M49 deferred the AI-keys policy to M50 (Fate-2) — which of `OPENAI`/`ANTHROPIC`/`MISTRAL`/
`ELEVENLABS`/the `LIVEKIT` pair become throwaway/sandbox demo values vs documented-as-absent. M50 is the
milestone that first touches the AI-consuming surfaces (academy AI chat, M45 batch-gen), so the policy is
decided here at close.

**Options:** (a) provision throwaway/sandbox keys into the demo secret source so the AI surfaces are live;
(b) document the keys as absent and let the AI surfaces no-op gracefully.

**Choice: (b) documented-as-absent.** The demo's content believability needs **no live AI** — every seeded
surface renders from seeded structural data, not a model call. The AI-provider keys stay **absent** from the
demo secret source; **no real key is ever provisioned** (the decision is itself values-blind). The
AI-dependent surfaces (sim voice via LiveKit, ant-academy `/api/ai/chat`, the M45 `ai v1.40.1` batch-gen) are
**inert-by-design** unless an operator supplies their own sandbox/throwaway keys — none is on a demo gate path
(the M42 gate is MET both vantages with zero AI keys). These keys join the **`waived`/optional** class for a
demo source (the `waived-aws-mount` sibling) so the values-blind `check` does not false-fail.

**Why:** the bring-up + the believability gate provably do not need live AI; provisioning real keys would add
a supply-chain + cost surface for zero demo benefit. **Documented → `corpus/ops/secrets-spec.md`** (the
M50-placeholder note converted to the decided policy).

## D-CLOSE-2 — COLD reset-to-seed acceptance → M53 (Fate-2, user-decided) — 2026-06-30

**Context:** the M50 exit_gate names "on a COLD reset-to-seed demo". The warm gate is MET on both vantages on
the STRENGTHENED manifest; all M50 seeders + fixes are baked into the bring-up tooling and reproduce on a
fresh `/demo-up`. The COLD destroy-and-rebuild proof is the v1.10b M53 "cold-rebuild acceptance" milestone's
defining work.

**Choice: Fate-2 carry-forward to M53.** The user explicitly decided to defer the cold-environment proof to
M53 (recorded in the close orchestration). No fresh sign-off required (the decision is already made); no plan
edit needed (M53's overview already lists both-vantage coverage on a from-cold rebuild). The Gate Outcome
Ledger records M50 as **closed-on-gate** for the M42 warm-gate metric, with the COLD clause as a documented
Fate-2 carry-forward — NOT an escape-hatch, NOT closed-incomplete.

**Why:** M53 owns the single from-cold acceptance truth for the whole release (the 1-demo-stack sequential
chain ends there); proving COLD twice (here + M53) would duplicate the heavy ~15-25 min rebuild for no added
assurance — M50's job is the seeders + the warm proof that they render.

## D-CLOSE-3 — Academy content + menu-link/non-anonymous-session (F6) → M51 (Fate-3) — 2026-06-30

**Context:** M50's candidate fix surface listed the hero academy link + a non-anonymous session (F6); the
field review flagged the academy as `0 chapters / 0 skill-paths`, no menu link, anonymous-if-direct.

**Choice: Fate-3 annotate M51.** The M42 gate is MET both vantages WITHOUT the academy (it is not on any
coverage-gate path), so it is not an M50 gate blocker. The academy course-content + menu-link +
non-anonymous-session are a seeding/content surface — M51 (the AI-readiness showcase-org milestone) already
touches seeding/content and is the natural home. **Annotated → M51 `overview.md`** candidate-scope list (with
the M50 provenance). The academy AI chat stays documented-as-absent per D-CLOSE-1.

**Why:** routing the seedable academy surface to the milestone that already owns the seeding/content domain
keeps the work coherent rather than fragmenting it across M50; the AI-chat half is correctly inert by the
keys policy, so only the content/wiring half carries forward.

## Adversarial review (Phase 2c) — 2026-06-30

**Scenario — the load-bearing DB-trigger dependency (`MemberLanguagesSeeder`).** The seeder writes ONLY
`user_languages` and relies on the platform's AFTER-INSERT trigger
`on_insert_user_languages_insert_membership_languages` to fan each row out to `membership_languages` (the
column the manager Talent-tab "Languages spoken" chart reads). **Failure mode:** if the trigger is absent on a
given stack, `membership_languages` stays empty, the chart reads empty, and the seeder reports SUCCESS (it only
counts catalog + `user_languages` rows) — a silent believability failure. **Response:** the scenario is
PROVEN-handled, not just argued: iter-06 observed `membership_languages=747` via the trigger on demo-1, and the
manager M42 gate ASSERTS the rendered Talent-tab languages chart — so the live gate IS the regression test for
the trigger dependency (absent trigger → empty chart → gate fails). The COLD M53 acceptance re-proves it on a
fresh DB. The seeder cannot add the trigger itself (zero-platform-edit; the trigger is the platform's schema),
so write-only-`user_languages` + assert-the-render is the correct design. The dependency is documented in
`member_languages.go:35-46`. This is a fleet-wide convention (`users.go` relies the same way on
`trigger_init_user_tables`), not an M50 regression. No fix required.
