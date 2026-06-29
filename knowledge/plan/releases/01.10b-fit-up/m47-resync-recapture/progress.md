# M47 — progress

## Section checklist

### S1 — Milestone setup ✅
- [x] rext authoring copy cloned to `.agentspace/rosetta-extensions` (on `main`, clean, @ `ce42d1e`)
- [x] milestone branch `m47/resync-recapture` cut from `release/01.10b-fit-up`
- [x] Phase 0b KB-fidelity verdict recorded (see `spec-notes.md`)

### S2 — Cold-start MCP-DSN auto-capture (rext code) — demo-up #2
- [x] normalize `sslmode=no-verify → require` — `pg.NormalizeDSN` in `stack-snapshot/pg/pg.go`, applied at the `Connect` choke point (rext `c5323a1`); 7 table tests + build/vet/test green
- [x] ~~accept the MCP DSN as a primary-read source in `source.go`~~ — **already supported**: `source.go` precedence + `cmd/stacksnap/main.go:204` already make `--dsn` a primary-read candidate; no change needed
- [ ] drop the cold-cache prompt → auto-capture instead *(needs the MCP-DSN-extraction decision — checkpoint)*
- [ ] wire the auto-capture into `demo-stack/up-injected.sh` set-dress (cache-miss → extract MCP DSN → `stacksnap capture --dsn … --source primary-read`)
- [ ] tag `fit-up-m47` + re-pin `stack-demo/rosetta-extensions` *(after the wiring lands)*

### S3 — Re-sync platform clones to current prod — operational ⚠
- [ ] capture before-refs; pull/checkout current per-repo refs across the `stack-demo` repo set
- [ ] rebuild images; re-migrate
- [ ] record after-refs; note any breakage absorbed

### S4 — Recapture snapshot from current prod — operational, uses S2
- [ ] run the auto-capture (taxonomy + Directus) over the sanctioned MCP DSN, public-only firewall (0 tenant rows)
- [ ] bump the capture version; refresh `.agentspace/snapshots/` cache + manifest
- [ ] coordinate the version bump with the M52 batch-cache key

### S5 — Re-validate the M201 false-negatives
- [ ] re-grade member-AI-readiness (+ the other M201 negatives) against current code
- [ ] record which were stale false-negatives (feeds M48 doc + M51 seeder)

### S6 — Doc: snapshot-cold-start.md (Delivers →)
- [ ] document the sanctioned MCP-DSN auto-capture path (supersedes the M20-D4 "MCP is NOT a capture source" stance)

## Notes
- **Section order rationale:** S2 (code) before S4 (recapture) because recapture *uses* the new auto-capture. S3
  (re-sync) is operational + the ⚠ release risk — checkpoint with the user before kicking off the multi-service
  Docker rebuild on the single demo machine.
