# M252 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [ ] **AI-key demo-container wiring** — the provisioned studio-desk AI key reaches the demo container
  at runtime (`env_file: <clone>/studio-desk/.env` in `gen_injected_override.py`, OR a
  `bridge_studio_ai_creds()` in `up-injected.sh`); `/api/ai/completion` no longer 500s.
- [ ] **DNA hardening** — a demo-aware assertion that the studio-desk **container** carries a provider
  key (closes the `.env`-vs-container gap).
- [ ] **Builder Playthrough** — `studio-builder-page.ts` page-object + `studioBaseUrl(9000+offset)` +
  studio Clerkenstein hero-login + `manifest/studio-builders.yaml` + an admin/content_creator
  precondition (studio-desk's first entry in the playthroughs manifest).
- [ ] **Talk-to-data double-check** — re-confirm the M239 Bedrock path (expected: COMPLETE, no gap;
  recorded for the audit trail).
- [ ] **Delivers** — `corpus/services/studio-desk.md` + `corpus/ops/secrets-spec.md` (demo-aware
  studio-desk AI note) + `corpus/ops/demo/playthroughs.md` (the builder Playthrough + count).

## Completeness Ledger

### Deferred

### Dropped
