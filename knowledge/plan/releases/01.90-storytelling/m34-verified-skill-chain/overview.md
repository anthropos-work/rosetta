---
milestone: M34
slug: verified-skill-chain
version: v1.9 "storytelling"
milestone_shape: section
status: planned
created: 2026-06-22
last_updated: 2026-06-22
complexity: large
delivers: rosetta-extensions/stack-seeding (jobsim_sessions.go G14 fix + TaxonomyRefs resolver + PersonaSeeder 7-table chain + users.go names/avatars + the closure data-DNA gene) + the rosetta corpus doc-half (NEW corpus/ops/demo/stories-spec.md — the verified-skill-chain reference, graduating .agentspace/seeding_gaps.md per spec-D12 — plus seeding-spec.md / safety.md updates)
depends_on: none
spec_ref: .agentspace/seeding_gaps.md (adversarially-verified 2026-06-22; 6-agent recon + 3-agent code review). Reference impl to PORT (not reinvent) = kb-ant-business/.claude/skills/seed-verified-skill/scripts/seed.sql
---

# M34 — Verified-skill chain (vertical slice)

## Goal
One seeded hero's **skill profile + Skill Spotlight chart** render end-to-end on a `--local-content` demo
stack — proving the verified-skill **spine** before anything is scaled. Today the seeder writes **zero**
verified skills (the core product surface is empty), and the session seeder writes invalid enum/result
values so its sessions are *inserted-but-invisible* dead rows (G14). M34 makes one hero (**Maya**) real:
passed sessions → validation rows → `user_skills` + `user_skill_evidences` → a profile that renders, with a
chart that plots and a claimed-vs-verified gap that shows.

## Why section
The path is known and code-verified: the 7-table chain, every column/enum/constraint, and the two
fields the reference `seed.sql` omits are pinned in the spec (`.agentspace/seeding_gaps.md` §3, §6). The
deliverables are a fixed checklist; build with `/developer-kit:build-milestone`.

## Scope
**In:**
- **Fix `jobsim_sessions.go` (G14).** Valid `status='ended'` / `completion_status∈{passed,failed}` /
  `result_status='completed'` / a `[a-z0-9]{5,10}` token / the **full** `SIMULATION_TYPE_*` strings; a
  continuous mid-skewed score (replacing binary 85/35) + a per-user growth arc; the ASSESSMENT/HIRING share
  that feeds verification.
- **`TaxonomyRefs` resolver** — mirror `contentref.go`: read real public `skiller.skills.node_id` from the
  replayed taxonomy + a `skillsByRole(roleName)` query (`job_roles ⋈ job_role_skills ⋈ skills`, is_core-first);
  resolve once/run; **empty-pool fallback** — skip enrichment, **never fabricate** a node_id.
- **`PersonaSeeder`** — the **7-table chain** per (hero × skill): `jobsimulation.sessions` →
  `validation_attempt_results` → `validation_attempt_skill_results` → `validation_criterion_results` →
  `public.local_jobsimulation_sessions` → `public.user_skills` → `public.user_skill_evidences` (UPSERT).
  Port `seed.sql` shapes **plus** the two fields it omits — `user_skill_evidences.user_level` (per the
  hero's `self_eval_bias`, else the claimed-vs-verified widget is empty) and `result_status` on sessions.
- **Patch `users.go`** — real names (name bank), deterministic avatar URLs, org-domain emails.
- **The closure assertion gene** (data-DNA, mirroring the M23 cross-surface gene): after seeding, count
  `user_skills`/`user_skill_evidences`/`validation_attempt_skill_results.skill` node_ids that don't resolve in
  the replayed `skiller` — must be **0**; name a sample on failure.
- **Prove on Maya** — one hero: her profile renders, the Skill Spotlight chart plots (≥2 datapoints, valid
  sim_type/eval/score), the closure gene is green, on a `--local-content` demo stack.

**Out:** multi-org (M35), the `stack.stories.yaml` model (M35), the full trio (M35), the org-aggregate
dashboard surfaces (M36), the cockpit (M37/M38).

## Load-bearing constraints (from the verified spec — do not re-derive)
- Enum/result/sim_type/token columns are **free-text varchar** (no PG enum/CHECK) → a wrong value INSERTs
  but is **filtered out of every query** (the G14 class). Write the exact strings.
- `user_skills` DB CHECK `user_skills_check_foreign_keys` → set `job_simulation_id` (the SIMULATION/Directus
  template UUID, **not** the session UUID). Partial UNIQUE `idx_unique_job_simulation` → a **distinct real
  sim_id per verified row**. `user_skill_evidences` UNIQUE (skill_id, user_id) → **UPSERT**.
- Chart needs **≥2 datapoints**, `sim_type ∈ {SIMULATION_TYPE_ASSESSMENT, SIMULATION_TYPE_HIRING}`,
  `evaluation_status='passed'`, `competency_level_score>0`. Levels stored **0–100** (÷ `maxLevel`, default 5
  — no settings seed needed). The misspelled column `local_jobsimulation_sessions.completition_status` (sic).

## Repo split
- **`rosetta-extensions`** (authoring → tag → consume per-stack): `stack-seeding/seeders/` (the new
  `PersonaSeeder`, the `jobsim_sessions.go` + `users.go` patches, the `TaxonomyRefs` resolver) +
  `stack-seeding/dna/` (the closure gene).
- **`rosetta`** corpus: **NEW `corpus/ops/demo/stories-spec.md`** (the verified-skill-chain reference —
  graduates the gitignored `.agentspace/seeding_gaps.md` per spec-D12) + updates to `seeding-spec.md`
  (the G14 fix + the chain) and `safety.md` (the new write surfaces stay PerStackIsolated).

## Open questions
- **O4** — exact MIGRATED storage-key names (`user_skill_user`, `membership_skill_*`, etc.): one `\d` pass on
  a replayed stack via `/db-query` before COPYing.
- The live render proof needs a `--local-content` demo stack (taxonomy + content replayed — the default).

## Done-when
Maya's seeded profile renders with verified skills; the Skill Spotlight chart plots; the claimed-vs-verified
gap shows; the closure gene reports 0 dangling node_ids; `stories-spec.md` exists; tests green; zero
platform-repo edits.
