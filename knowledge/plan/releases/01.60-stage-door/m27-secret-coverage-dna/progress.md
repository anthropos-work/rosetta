# M27 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [ ] `stack-secrets/` section scaffolded (go.mod + cmd/stacksecrets) in the `.agentspace` authoring copy
- [ ] Source-ingestion reader: directory mode (default `.agentspace/secrets`), values-blind
- [ ] Source-ingestion reader: zip mode
- [ ] Source-dir layout contract (zEnvs / per-repo .env never silently ingested)
- [ ] secret-DNA schema + `secret-dna.json` (gene = repo×KEY; the per-gene fields; strict load + Validate)
- [ ] `introspect` from the hybrid source (platform/.env_example + frontend/sentinel .env.example + compose-injected)
- [ ] `list` verb
- [ ] `diff` verb — required-key drift exit 1 (the "keep-listed" gate) + undeclared-runtime-required guard
- [ ] Waived classes modeled (AWS-mount, profile-gated, optional Bunny/GCloud)
- [ ] Alias families encoded (GH_PAT family) vs distinct-similar pairs (OPENAI_KEY/OPENAI_API_KEY — not auto-aliased)
- [ ] Hermetic unit tests (no values)
- [ ] Ext tag `stage-door-m27`

## Notes
_(append build notes here)_
