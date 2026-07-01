# M53 — progress

## Section checklist
_Checked off as each In-scope deliverable lands. Close when all boxes are ticked._

_Ordered per overview.md acceptance flow (a)→(f). Sections gate on prior sections._

- [ ] **§1 — Academy F6 seeder/wiring** (the ONE new-code section; in the rext authoring copy)
  - [ ] Course content seeding present (real chapters/skill-path content, not empty shell)
  - [ ] Hero academy menu-link routes from cockpit/persona into the academy
  - [ ] Non-anonymous academy session (hero lands authenticated)
  - [ ] AI chat stays documented-as-absent (NO `/api/ai/chat` assertion)
  - [ ] Commit on rext authoring `main`
- [ ] **§2 — Roll v1.10.1 rext release tag** (rolls up `fit-up-m47..m52` + academy commit)
  - [ ] Tag `v1.10.1` on rext authoring `main`
  - [ ] Bump `.agentspace/rext.tag` → `v1.10.1`
  - [ ] Re-pin consumption clone `stack-demo/rosetta-extensions` → `v1.10.1` (clean fetch + checkout)
- [ ] **§3 — DESTROY the live demo** (`/demo-down` + image purge — exercises M49 #6 cleanup)
- [ ] **§4 — COLD REBUILD** (single `/demo-up` at v1.10.1 pin, no manual steps)
- [ ] **§5 — ASSERT the acceptance bar** (all 6 criteria + academy F6)
  - [ ] AB1 — all backends healthy (M47/M49; no silent `app Exited`)
  - [ ] AB2 — cold-start MCP-DSN auto-capture filled snapshot with NO prompt (M47)
  - [ ] AB3 — set-dress + seed (all 3 orgs incl. AI-readiness) + verify + cockpit complete (no #7 abort)
  - [ ] AB4 — both-vantage M42 coverage GREEN (employee + manager), 0 escapes (M50)
  - [ ] AB5 — AI-readiness dashboard criteria hold on 3rd org (M51: enabled, ~80%/3-step, 1 started + 1 completed; fast render)
  - [ ] AB6 — cockpit [Download manifest] returns complete inlined `seed-generation-manifest.yaml` (M52)
  - [ ] F6 — academy: course content present + hero menu-link + non-anonymous session
- [ ] **§6 — Acceptance record + rext.tag bump** (corpus acceptance note; feeds close-release)
