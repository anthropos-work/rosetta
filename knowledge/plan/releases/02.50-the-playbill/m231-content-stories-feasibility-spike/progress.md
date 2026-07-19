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

- [x] **S5 — Deliverable consolidation**
  - [x] Assemble the manager-view eligibility matrix (§2 + §7; has_manager_view per product)
  - [x] Finalize `corpus/ops/demo/content-stories-routes.md` (route map + modality catalog + AI-labs verdict + sourcing/anonymization contract + go/no-go synthesis) — all cross-links validated
  - [x] Wire discoverability (indexed from demo/README.md + CLAUDE.md); go/no-go verdict + three-fate routing (D3→M232, D4→M234 via overview edits)

## M231: Final Review — Completeness Ledger (section, spike)

**Done (Fate 1):** S1 result-route map + prove-by-render classification · S2 prod-session sourcing + anonymization
contract · S3 public-sim-by-modality catalog (77 voice / 65 code / 30 document) · S4 AI-labs (OUT) + academy (IN)
verdicts · S5 go/no-go synthesis + discoverability wiring. Deliverable `corpus/ops/demo/content-stories-routes.md`
(349 lines, 23 code-citations, 0 broken links). Inline KB-fidelity fixes: hiring.md (intercepting-route→plain Drawer),
skillpath.md (manager mirror), ant-academy.md (backend read/WRITE). All ✅.

**Confirmed-covered (Fate 2):** D5 academy content-story depends on M230 (already landed). D4 AI-labs-out aligns with
the roadmap's conditional framing.

**Annotated (Fate 3) — user-accepted 2026-07-19:** D3 interview PostHog-flag-enablement → **M232** (overview In edited)
· D4 AI-labs presence-only section → **M234** (overview In edited) · D5 academy real-progress form → **M234** (overview
In edited). M232 also gained the public-anchored-sourcing rule + the code/document modality seeders.

**Tracked (Fate-3, arch-doc pass):** KB-2/KB-5/KB-6/KB-8 (jobsimulation.md ports; roadrunner-retired; backend.md labs
path) — tangential doc-accuracy fixes recorded in decisions.md; the deliverable carries the correct facts.

**Dropped:** none. **Escape-hatch:** none.

**Verdict:** Thread B is a GO. All scope delivered; Fate-3 routings user-accepted; deferral audit GREEN (0 repeat-defer,
0 escape-hatch). KB-fidelity YELLOW (the milestone delivers the consolidation doc — no blind area). Clean close.
