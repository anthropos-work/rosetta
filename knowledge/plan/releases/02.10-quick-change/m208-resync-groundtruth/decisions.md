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
