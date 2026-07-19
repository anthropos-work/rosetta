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

- [ ] **S2 — Prod-session sourcing + anonymization mechanism**
  - [ ] Confirm the `/db-query` read path selects interesting real prod sessions per type (honor db-access boundary)
  - [ ] Identify which fields scrub cleanly vs which free-text needs handling
  - [ ] Confirm how to pin a source by prod session-id (deterministic reseed)
  - [ ] Author the sourcing + anonymization contract (mechanism only; the copy is M232)

- [ ] **S3 — Public-sim-by-modality catalog**
  - [ ] Confirm ≥2 voice + 1 code + 1 document-assessment SOURCES exist to pin (modality = LiveKit/Chime voice, Judge0/Roadrunner code, Gotenberg document)
  - [ ] Map each modality to a concrete pinnable public simulation source

- [ ] **S4 — AI-labs feasibility + academy "session" verdict**
  - [ ] Rule AI-labs in/out (labs-api client wired nil?)
  - [ ] Rule the ant-academy content-product section in/out (is there a server session store post-M230?)
  - [ ] Record each verdict with the code-cite that decides it

- [ ] **S5 — Deliverable consolidation**
  - [ ] Assemble the manager-view eligibility matrix (which products HAVE a manager result route)
  - [ ] Finalize `corpus/ops/demo/content-stories-routes.md` (route map + modality catalog + AI-labs verdict + sourcing/anonymization contract)
  - [ ] Wire discoverability (index the new doc from demo/README.md + parent pointers); go/no-go verdict + three-fate routing for un-renderable surfaces
