# M28 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [ ] `provision` verb + per-repo target-file map (ant-academy → `code/.env.local` pinned)
- [ ] Alias-mapping per target file (one source value → all its per-file aliases)
- [ ] Idempotency + overwrite policy (copy-if-absent default, `--force`)
- [ ] N=0 main-dev-stack guard (refuse without `--force`)
- [ ] Compose-with-injection-override: never re-arm prod `DIRECTUS_TOKEN` on non-prod / `--local-content` **+ regression test**
- [ ] `PreflightEnv`-passing env emission
- [ ] `check`/`measure`: Overall + Critical (gate==100%) + per-repo rollup
- [ ] Demo-aware coverage (Clerk keys minted-OK)
- [ ] Non-fatal pre-flight wiring into `/dev-up` + `/demo-up` (warn standard, fail critical)
- [ ] Profile-scoping decision settled + implemented
- [ ] Hard safety verified: no verb reads/echoes/logs a value
- [ ] Ext tag `stage-door-m28`

## Notes
_(append build notes here)_
