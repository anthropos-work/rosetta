# M208 — Decisions

_Implementation choices with rationale, logged as they are made._

## KB-fidelity items (Phase 0b, 2026-07-08)

Audit verdict **YELLOW** (report: `kb-fidelity-audit.md`). The corpus currently carries pre-merge
claims — that is the premise of this release, fully tracked. All three are **Fate 2** (owned by a
future milestone of this release, M210's corpus body-flip); M208 pins the authoritative fact-sheet
anchor they grade against.

- **KB-1** — `corpus/services/backend.md` still describes skiller as a separate downstream service
  ("consumed by skiller", "Skiller — taxonomy and matching RPC", "Consumer: … skiller events").
  Stale vs the merged code. → **M210** full body-flip. M208 adds the concise merge fact-sheet section.
- **KB-2** — `corpus/services/skiller.md` still documents a live standalone service. → **M210**
  (colleague's `origin/docs/skiller-in-app-merge` already drafts the "merged into app" stub).
- **KB-3** — `CLAUDE.md` / `corpus/services/graphql-wundergraph.md` say "5 subgraphs". Actual = 4 at
  `platform@origin/main`. → **M210**. M208 pins "4 subgraphs" in the fact-sheet.

Not read as truth by M208's own implementation — M208 authors the correction, grounded in the
verified app clone (`internal/data/ent/schema/skill.go` … + merge commit `1fc00c78`) +
`platform@origin/main` (`SKILLER_RPC_ADDR=http://backend:8083`, no skiller in repos.yml/compose) +
the colleague's docs branch.

## Environment: parked native-dev override for the containerized de-risk (2026-07-08)

`stack-dev/platform/docker-compose.override.yml` exists (untracked, local-only) — the native-worktree
dev override that maps `backend:host-gateway` on the `graphql` (Cosmo router) service so the router
reaches a **natively-run** app on the Mac host instead of the backend container (see MEMORY.md
"dev-native-worktree-topology"). M208's chartered de-risk is an honest **fully-containerized** merged
bring-up (prove migrations apply + the 4-subgraph compose comes up with no skiller container). With the
override present and no native app running (cold state), the router would route the backend subgraph to
a dead host-gateway. **Decision:** temporarily PARK the override (`mv …override.yml
…override.yml.m208-parked`) for the duration of the containerized `make up`/verify, then RESTORE it
verbatim at section close. Fully reversible; the user's native-dev config is returned untouched.
`stack-demo/platform` has no such override.

## Finding 2 — INVITATION_HMAC_SECRET per-stack .env gap → M211 / /stack-secrets (2026-07-08)

The cold containerized `backend` `Exited(0)` at startup on `INVITATION_HMAC_SECRET is not set` (fatal:
`app/main.go:266` → `internal/invitations/token.go:23`; POSTHOG key also empty). **Investigated** (three-fate
guard): a **documented, known** required secret — `corpus/ops/secrets-spec.md:111` states "the app exits
early when it is unset"; one of the platform `.env` 29 keys (line 98); in `secretdna.DemoGeneratedKeys`
(demo auto-generates it; app `docs/invitation-reminders.md` prescribes `openssl rand -hex 32`). **Not a
merge regression** — stack-dev's hand-assembled `.env` lacked it because its containerized `backend` is
normally never run (native-worktree dev). During the existing-volume run I did a Fate-1 dev-value add
(values-blind) to confirm the federation serves end-to-end, then **reverted it** per the orchestrator
directive ("do NOT provision secrets here"); the `.env` is left in its original gap state. **Routed to
M211 / a `/stack-secrets` follow-up** — the sanctioned provisioner, not an ad-hoc edit. No corpus change
(the secrets-spec already documents the key + the exit-early behavior).

## Finding 1 — clean-bring-up extensions bootstrap + PG-readiness (M25-D9 class) → M211 Fate-3 + M209 Risk-2 (2026-07-08)

The M25-D9 opportunistic item **surfaced** — on the **clean-slate** `make reset-db` path (the existing-volume
migrate masked it). A clean reset-db does not bootstrap the `extensions` schema (pgvector + `pg_trgm`) before
migrate, so the merged vector/trigram migrations fail on an empty DB: app `20260518125439`
(`ask_query_examples.embedding extensions.vector(1536)`) and cms `20250116133510`
(`similarities.small_embedding3 …`) → `pq: schema "extensions" does not exist`; app `20260623090000` needs
the `extensions.gin_trgm_ops` opclass. Manual `CREATE SCHEMA extensions; CREATE EXTENSION vector/pg_trgm
SCHEMA extensions;` unblocks table creation. Plus a secondary **PG-readiness race** on reset-db
(`connection reset by peer`; a re-run succeeds). It did **not** trivially fall out as Fate-1 (it's a
bring-up-tooling requirement). **Fate-3: routed to M211** (bring-up acceptance — edited M211/overview.md
In-scope to name the extensions-bootstrap + PG-readiness requirement) with an **M209 Risk-2 cross-ref**
(the capture column list uses `extensions.vector` + `extensions.gin_trgm_ops`). Details in spec-notes
Finding 1.

## Live de-risk result (2026-07-08) — the milestone's load-bearing proof

The #1 release risk (86-commit `app` pull + migration re-run) is **retired GREEN**: the merged image
builds; a clean-slate `make reset-db` + `make migrate` creates the **full public taxonomy from scratch**
(public.skills with `organization_id`, job_roles, job_role_skills, skill_embeddings, categories,
specializations) with **no `skiller` schema on a clean DB** — provided the `extensions` schema is
bootstrapped first (Finding 1); the 4-subgraph compose comes up with **no skiller container** and
`SKILLER_RPC_ADDR=http://backend:8083`; the merged app subscribes to the `skiller` Redis stream in-process
and its DB search_path has no skiller. The existing-volume run additionally confirmed the router SERVES the
absorbed taxonomy subgraph end-to-end (with a dev HMAC secret present; the clean-slate run couldn't, per
Finding 2). Measured prod public-skill count (`public.skills WHERE organization_id IS NULL`) = **42,790**,
confirming the ~42,763 assertion (the figure M209's post-capture assertion will grade against).
M209/M210/M211 grade against this proven state.

## Fact-sheet placement — minimal anchor, not M210's body-flip

M208 delivers a **minimal + authoritative** merge fact-sheet: a self-contained `## Skiller-in-app merge`
section added to `corpus/services/backend.md` (the grading contract for M209/M210/M211) + a top-of-file
pointer banner on `backend.md` and a minimal stub banner on `corpus/services/skiller.md`. The pre-merge
prose bodies are **left untouched** (KB-1/2/3 → M210's full body-flip). Deliberately did NOT adopt the
colleague's `origin/docs/skiller-in-app-merge` full drafts (that branch is M210's to land) — referenced
only to ground facts.

## Adversarial review (Phase 2c, close 2026-07-08)

M208 touched **no code modules** — the diff is 100% documentation (fact-sheet + milestone records + two
sibling-overview routing edits + state.md; `git diff --name-only e319d2f..HEAD` = zero `.go`/`.ts`/`.sh`
files). The classic per-module adversarial angle is therefore N/A. The one artifact carrying downstream
risk is the **fact-sheet as a grading contract**: a wrong fact or a broken anchor would make M209/M210/M211
grade against a falsehood. Scenario examined + neutralized:
- **Broken anchor** (skiller.md → backend.md#…): the two references use
  `#skiller-in-app-merge--fact-sheet-v21-quick-change`, which is exactly the GitHub slug of the
  `## Skiller-in-app merge — fact-sheet (v2.1 "quick change")` heading (verified by slug computation +
  `grep` at close). Resolves.
- **Wrong load-bearing number/fact**: every grading fact — tables moved to `public` (names unchanged), the
  `organization_id IS NULL` public predicate, the **42,790** public-skill count (measured on prod; the
  roadmap's ~42,763 reconciled), the `SKILLER_RPC_ADDR=http://backend:8083` re-point, **4 subgraphs**, and
  **no skiller container/schema on a clean DB** — is grounded in the verified re-synced clone
  (`app@c3c45e01`, `platform@0808b92`), the live containerized `make up`+`reset-db`+`migrate` de-risk, and
  read-only prod. No unverified load-bearing assertion.
- **Residual (non-load-bearing):** the fact-sheet's "`categoryTree`/`fullCategoryTree` were dropped, not
  ported" is a GraphQL-surface detail M210 verifies against `backend.graphqls` during the subgraph
  reconciliation (M210 owns deep subgraph accuracy); it is not read by M209's `public.*` tooling re-ground,
  so it is not load-bearing for the release's tooling path. Noted for M210, not a close-blocker.
