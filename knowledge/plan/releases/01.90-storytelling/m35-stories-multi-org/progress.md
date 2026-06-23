# M35 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

- [x] **`stories[]` blueprint** — per-hero `vantage`/`trajectory`/`skills`, per-story `org`/`narrative`; supersedes `stack.seed.yaml` for demo stacks
- [x] **Multi-org parameterization** — per-story `OrgID` + `orgClerkID` threaded through the 4 seeders (org/users/identity/jobsim-sessions/assignments) + Clerkenstein org-claim alignment _(first story keeps the Clerkenstein default org; all 8 seeders threaded, not just 4)_
- [x] **`PersonaSeeder` roster scaling** — the 2-stories × 3-heroes v1 roster (Cervato + Solvantis) _(+ #M34-D7: collision-free declaration-order hero slots + warning, short-pool flat top-up)_
- [x] **Trajectory logic** — thriving (dense/rising/under-claim) vs struggling (sparse/low/over-claim)
- [x] **Supporting-population fidelity** — `job_role_id`+name (real replayed via `jobroleref.go`), ramped `joined_at`, names on non-hero members
- [x] **Single-org default preserved** — existing `dev-min`/preset path still passes (the legacy blueprint resolves to a byte-identical one-story view; full M34/M7 suite green unchanged)
- [x] **Docs** — extended `stories-spec.md` (the model) + `seeding-spec.md` (blueprint supersession) + `/stack-seed` SKILL.md + the rext module README; M35 `decisions.md`
- [x] **Tests** — `stack-seeding` suite green (326 test funcs, +24); closure gene green across all orgs (both integration tests pass live against demo-3); `-race` clean

_Last updated: 2026-06-23 (all sections complete; ext tag `storytelling-m35`)._
