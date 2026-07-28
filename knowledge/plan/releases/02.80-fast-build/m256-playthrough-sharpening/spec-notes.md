# M256 — spec notes

_Technical notes accumulate here during the build._

## Pre-flight audits — iter-01

**`/developer-kit:audit-kb-fidelity --milestone=M256` — verdict `YELLOW`** (2026-07-28, `SEVERITY: warning`).
Report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md). Phase 0b proceeds on YELLOW; the gaps become the
iter's known-context and are folded into `TOK-01`'s *Known context* block.

4 blind areas · 13 stale claims (11 fixed inline across `playthroughs.md`, `clerkenstein.md`,
`coverage-protocol.md`, `seeding-spec.md`) · 7 completeness gaps · 6 open items.

**It caught a live error inside the iter that ran it:** iter-01's first-draft `D4` ("the `blocked` outcome
needs no seed work") read `seed-worlds.yaml`'s *declaration* as seeded state. `actor.entitlement` reaches no
DB column — `blueprint.TierMix` is parsed/defaulted/validated and consumed by no seeder — and
`ptvalidate`'s precondition-coverage check **fail-opens** on it. D4 is recorded as refuted, with three
replacement refusal surfaces, in `iter-01/decisions.md`.

`D1` (parallelism off the critical path), `D2` (org-admin first) and `D3` (the `networkidle` lever) all
reconcile — `D3` **strengthened**: the audit found **8** unfenced `networkidle` violations, 6 of which
(unbounded `waitForLoadState` sites) `D3` had not seen.

The confirmed blind area worth naming: the **Clerkenstein single-global-seat** property appeared in exactly
one corpus doc, as a rider on an unrelated feature — while `playthroughs.md` and the capability spec
justified serial-by-default **solely by Postgres**, the rationale the plan review had already refuted. Now
documented in `playthroughs.md` and `corpus/services/clerkenstein.md`.

## Topic → knowledge doc → code triples (from the Phase-0b KB-fidelity audit, 2026-07-28)

Code paths are relative to `.agentspace/rosetta-extensions/` unless noted. Full audit:
[`kb-fidelity-audit.md`](kb-fidelity-audit.md).

| Topic | Knowledge doc | Code |
|---|---|---|
| Playthrough count / corpus | `corpus/ops/demo/playthroughs.md` §"18 live Playthroughs, 0 TODO" | `playthroughs/manifest/*.yaml` (18 `playthrough:` keys) · `playthroughs/e2e/tests/*.spec.ts` (18 `test()` in 17 files) |
| Manifest schema + 4-state map | `playthroughs.md` §The model | `playthroughs/manifest/manifest.go`, `validator.go`, `report/report.go`, `report/unimplementable.yaml` |
| `@pt:` tag registry + grammar lockstep | `playthroughs.md` §Both-way id integrity | `playthroughs/cmd/ptvalidate/discover.go` + `pttag_lockstep_test.go` · `playthroughs/report/playwright.go` + `pttag_lockstep_test.go` |
| Runner concurrency / serial default | `playthroughs.md` §The lifecycle | `playthroughs/e2e/playwright.config.ts` · `playthroughs/e2e/lib/stack-env.ts` §`resolveWorkers` |
| **Seat contention (the real binding surface)** | `playthroughs.md` §The lifecycle (added this audit) · `corpus/services/clerkenstein.md` §Multi-identity (added this audit) · `corpus/ops/demo/cockpit-spec.md` §*Limitation — one seat per stack* | `clerkenstein/clerk-frontend/registry.go` §`activeKey`/`active()`/`Select()` · `clerk-frontend/server.go` §`type Server`, `handleSelectIdentity`, `handleMe`, `handleToken`, `handleSignOut` |
| Hero login handshake | `cockpit-spec.md` · `playthroughs.md` §page-object layer | `playthroughs/e2e/lib/hero-login.ts` → `stack-verify/e2e/lib/cockpit-login.ts` §`loginAs`/`selectSeat` |
| `pt-world` seed + reset-to-seed | `playthroughs.md` §The Playthrough world · `corpus/ops/seeding-spec.md` §Re-run safe | `playthroughs/seed/pt-world.seed.yaml` · `seed/seed-worlds.yaml` · `manifest/seed_worlds.go` · `stack-seeding/cmd/stackseed/main.go` §`resetTables`/`doReset` · `playthroughs/e2e/run-playthroughs.sh` |
| Seeded hero fan-out | `corpus/ops/demo/stories-spec.md` | `stack-seeding/seeders/persona.go`, `persona_write.go` (the 7-table flush) |
| Presence sweep (reused foundation) | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/` (`playwright.config.ts`, `lib/crawl.ts`, `lib/section-assert.ts`, `lib/coverage-manifest.ts`) |
| `networkidle` doctrine | `coverage-protocol.md` §Never wait on networkidle · `corpus/ops/demo/latency-budget.md` | `playthroughs/e2e/lib/page-object.ts` §`goto` · fences: `tests/page-object.unit.spec.ts`, `tests/home-login-networkidle.unit.spec.ts` |
| Studio (LLM-bound) lane | `playthroughs.md` §the `studio` product | `playthroughs/e2e/tests/studio-builder.spec.ts` (`test.setTimeout(300_000)` @ :45, `180_000` @ :91) · `lib/studio-builder-page.ts` |

## KB prerequisites still open (audit Phase 4 — need a user call)

1. **Exit-gate clause 2's `blocked` outcome has no seedable mechanism.** `actor.entitlement` is
   **declared-only**: `blueprint.TierMix` is parsed/defaulted/validated but consumed by **no seeder**, so no
   tier ever reaches a DB column, and `pt-world.seed.yaml` declares no `tier_mix`. The `pt-free` seat exists and
   is annotated *"entitlement-gate use cases — outcome: blocked"* in `seed-worlds.yaml`, but it is not
   tier-gated and is used by 0 of 18 use cases. **Decide before iter-01 commits a strategy:** seed a real tier
   (a `stack-seeding` change), or source the `blocked` from a different refusal surface (cross-org access, an
   RBAC deny, a validation error), or re-cut the clause.
2. **Onboarding (5 UCs) needs a net-new SEED capability, not just tests.** `UsersSeeder` writes a membership
   for every seeded user unconditionally and no onboarding flag exists anywhere in `stack-seeding/`. A
   pre-onboarding actor requires a seeder + a `capabilities:` entry + a roster seat before `ptvalidate`'s
   precondition-coverage check can pass. (This answers `overview.md` Open Question 1: **no**.)
3. **Org-admin (4 UCs) is a documentation blind area too** — no manifest product, no page object, no locator,
   no corpus section for org settings / member management / roles / invites. Only `pt-assignment-assign`'s
   read-back pattern is prior art. Consider extending the overview's `Delivers →` line to name the org-admin +
   onboarding streams so the knowledge production is explicitly in scope.
