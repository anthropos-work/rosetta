# M27 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [x] `stack-secrets/` section scaffolded (go.mod + cmd/stacksecrets) in the `.agentspace` authoring copy
- [x] Source-ingestion reader: directory mode (default `.agentspace/secrets`), values-blind
- [x] Source-ingestion reader: zip mode
- [x] Source-dir layout contract (zEnvs / per-repo .env never silently ingested)
- [x] secret-DNA schema + `secret-dna.json` (gene = repo×KEY; the per-gene fields; strict load + Validate)
- [x] `introspect` from the hybrid source (platform/.env_example + frontend/sentinel .env.example + compose-injected)
- [x] `list` verb
- [x] `diff` verb — required-key drift exit 1 (the "keep-listed" gate) + undeclared-runtime-required guard
- [x] Waived classes modeled (AWS-mount, profile-gated, optional Bunny/GCloud)
- [x] Alias families encoded (GH_PAT family) vs distinct-similar pairs (OPENAI_KEY/OPENAI_API_KEY — not auto-aliased)
- [x] Hermetic unit tests (no values)
- [x] Ext tag `stage-door-m27`

## Notes
- Built stdlib-only (no pgx/yaml): the secret-DNA is JSON, the readers are `archive/zip` + `bufio`.
- The committed `secret-dna.json` is **55 genes across 6 repos** (platform, app, sentinel, studio-desk,
  next-web-app, ant-academy); 94 tests across the three packages, `-race` + `gofmt` clean.
- A `check`/`measure` verb (score a source; exit 1 if critical < 100%; per-repo rollup) was folded in too —
  the natural pairing with the DNA, exercised end-to-end against the real stack-dev (values-blind).
- **M27-D2** (in `decisions.md`): the keep-listed gate is DNA-scoped two-tier (gate-fatal only on
  already-tracked-secret omissions; never-tracked declared keys → triage candidates), so the gate is usable
  against real `.env.example` files that mix secrets with config noise. The diff-vs-stack-dev caught 10 real
  cross-repo DNA omissions → fixed Fate-1.
- Verified live: `stacksecrets diff --stack-root stack-dev` exits **0** (0 gate-fatal); `check` against
  stack-dev reports real coverage + per-repo shortfalls with **no secret value printed**.
