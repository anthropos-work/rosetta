# M231 — Progress

Spike deliverable: `corpus/ops/demo/content-stories-routes.md`. Sections are the discovery streams;
each records findings in `spec-notes.md` + writes its portion of the deliverable, prove-by-render/code-cited.

## Sections

- [x] **S1 — Per-product result-route map + prove-by-render classification**
  - [x] Enumerate result routes per (product × {player, manager}) for: Simulation {training/assessment/hiring/interview}, Skill-path legacy, Skill-path new (ant-academy), AI-labs
  - [x] Cite the render path in platform code (player route + manager route) per product
  - [x] Answer the central unknown: **RESOLVED — persisted read** (`queries.resolvers.go:70` plain SELECTs, no recompute)
  - [x] Classify each route: renders-from-seed | runtime-computed-blank | needs-demo-patch | no-surface
  - [x] Prove-by-render: code+DB proof (resolver SELECTs + prod persisted score/result_status/fan-out + `seed-verified-skill` precedent). Live billion render corroboration deferred to M235 prove-it-lands (no content-story sessions seeded on billion yet — that's the M232+ build)

- [x] **S2 — Prod-session sourcing + anonymization mechanism**
  - [x] Confirm the `/db-query` read path selects interesting real prod sessions per type (confirmed live; ASSESSMENT 5,172 / TRAINING 1,799 / HIRING 1,679 / INTERVIEW 488 completed)
  - [x] Identify which fields scrub cleanly vs which free-text needs handling (structured IDs/enums/numerics keep; free-text = actor names + LLM feedback + input_data + transcript + interview reports)
  - [x] Confirm how to pin a source by prod session-id (`sessions.id` uuid) + the public-anchoring inner-join rule
  - [x] Author the sourcing + anonymization contract (§3; mechanism only; copy is M232) + resolved the clone-session-subcommand open question

- [x] **S3 — Public-sim-by-modality catalog**
  - [x] Confirm ≥2 voice + 1 code + 1 document-assessment SOURCES: **GO 77 voice / 65 code / 30 document** public sims
  - [x] Map each modality to a concrete pinnable public simulation source (`directus.sim_tasks.task_type` ∈ {call, code, collaborative_doc/send_attachment, chat}; already snapshot-replayable)

- [x] **S4 — AI-labs feasibility + academy "session" verdict**
  - [x] Rule AI-labs: **OUT** — nil client persists a booting row but no VM/grade; grade_result not GraphQL-exposed; /labs/[id] reads live from labs-api. Presence-only in M234.
  - [x] Rule the ant-academy section: **IN** — backend-authoritative since v0.5 M2; academy_chapter_progress seedable via app/cmd/academy-seed. Depends on M230 catalog fill.
  - [x] Record each verdict with the code-cite that decides it (§5, §6)

- [ ] **S5 — Deliverable consolidation**
  - [ ] Assemble the manager-view eligibility matrix (which products HAVE a manager result route)
  - [ ] Finalize `corpus/ops/demo/content-stories-routes.md` (route map + modality catalog + AI-labs verdict + sourcing/anonymization contract)
  - [ ] Wire discoverability (index the new doc from demo/README.md + parent pointers); go/no-go verdict + three-fate routing for un-renderable surfaces
