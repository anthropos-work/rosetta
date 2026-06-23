# M34 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

- [x] **G14 fix** — `jobsim_sessions.go`: valid `status`/`completion_status`/`result_status`/token + full `SIMULATION_TYPE_*` + continuous mid-skewed score + per-user growth arc + ASSESSMENT/HIRING share *(rext `c0872be`)*
- [x] **`TaxonomyRefs` resolver** — real `skiller.skills.node_id` + `skillsByRole`, empty-pool fallback (mirror `contentref.go`) *(rext `87e5377`)*
- [x] **`PersonaSeeder`** — the 7-table chain per (hero × skill), incl. `user_level` + `result_status` (the `seed.sql` omissions) *(rext `7817553`; FK fix `dad8c72`)*
- [x] **`users.go` patch** — real names + avatars + org-domain emails *(rext `8e1b57a`)*
- [x] **Closure gene** — data-DNA: 0 dangling node_ids, `datadna measure-closure` (mirror the M23 cross-surface gene) *(rext `1e323e7`)*
- [x] **Maya proof** — integration test (the automated half) seeds the chain against demo-3: profile (18 verified skills) + Spotlight chart (18 datapoints) + the claimed-vs-verified gap render + closure green. *Live browser render = the orchestrator's post-build step.* *(rext `dad8c72`)*
- [x] **Docs** — NEW `corpus/ops/demo/stories-spec.md` + `seeding-spec.md` / `safety.md` / demo `README.md` / `CLAUDE.md` updates *(rosetta worktree)*
- [x] **Tests** — `stack-seeding` suite green (`-race` clean); integration test opt-in (`-tags integration`); zero platform-repo edits

_Last updated: 2026-06-23 (all sections landed; rext tagged `storytelling-m34`)._
