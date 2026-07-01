# M53 — progress

## Section checklist
_Checked off as each In-scope deliverable lands. Close when all boxes are ticked._

_Ordered per overview.md acceptance flow (a)→(f). Sections gate on prior sections._

- [x] **§1 — Academy F6 seeder/wiring** (the ONE new-code section; in the rext authoring copy) — rext `e91f004`
  - [x] Course content present (3250 static-JSON files ship with the clone — verify-only, not seeded)
  - [x] Hero academy menu-link: per-hero cockpit [Academy] deep-link (cockpit.go External entry + cockpit.py + up-injected --academy-base)
  - [x] Non-anonymous academy session: launcher sets BOTH e2e_persona bypass gates; cockpit link sets e2e_persona=member cookie
  - [x] AI chat documented-as-absent (Cosmo flag+key not provisioned; NO `/api/ai/chat` assertion) — D3 + frontend-tier.md
  - [x] Commit on rext authoring `main` (e91f004; tests +13, all green, shellcheck-clean)
- [x] **§2 — Roll v1.10.1 rext release tag** (rolls up `fit-up-m47..m52` + academy commit)
  - [x] Tag `v1.10.1` on rext authoring `main` (annotated; points at e91f004; 46-commit roll-up)
  - [x] Bump `.agentspace/rext.tag` → `v1.10.1` + canonical pin in rosetta_demo.md
  - [x] Re-pin consumption clone `stack-demo/rosetta-extensions` → `v1.10.1` (clean fetch + checkout; tree clean)
- [x] **§3 — DESTROY the live demo** (`/demo-down 1 --purge`) — all 17 containers + network removed, ALL demo-1 images purged (M49 #6 verified working); native academy/cockpit stopped
- [x] **§4 — COLD REBUILD** (single `/demo-up 1` at v1.10.1 pin, no manual steps) — EXIT 0, no #7 abort; 17 containers Up (0 Exited); autoverify GREEN; log: `cold-rebuild.log`
- [~] **§5 — ASSERT the acceptance bar** (5/6 + academy F6 PASS; **AB4 manager-half BLOCKER → routes to M51**)
  - [x] AB1 — all backends healthy: 17 containers Up, 0 Exited; health 200, casbin 1150, all probes passed
  - [x] AB2 — snapshot replay prompt-free from the filled cache (taxonomy/directus/sim-embeddings replayed, no prompt) — KB-1
  - [x] AB3 — set-dress + seed (3 orgs incl. Northwind AI-readiness) + verify + cockpit — EXIT 0, no #7 abort
  - [ ] **AB4 — BLOCKER: employee GREEN (0,0); manager `dan-manager`@Cervato RED (failingSections=2, both ai-readiness) — an M51 unconditional-seedPath regression → routes to M51 (see decisions.md AB4-REGRESSION). Do NOT fix in M53.**
  - [x] AB5 — AI-readiness dashboard GREEN on 3rd org (dana-manager@Northwind): enabled, 50/100, 199 members, 87% functional+, 3-step funnel, renders fast (541 chars, no timeout) — KB-2
  - [x] AB6 — cockpit [Download manifest] serves complete inlined `seed-generation-manifest.yaml` (7593B, 3 orgs + generation prompt + batch $0.3 ceiling + snapshot_sources)
  - [x] F6 — academy: content real (copilot/claude-code/ai-eng chapters), 9 cockpit [Academy] links→:13077, both e2e_persona gates set (authenticated member), Cosmo absent by design
- [ ] **§6 — Acceptance record + rext.tag bump** (corpus acceptance note; feeds close-release)
