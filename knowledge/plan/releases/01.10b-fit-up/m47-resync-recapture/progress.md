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

### S3 — Re-sync platform clones to current prod ✅ (the heavy re-sync was a no-op — see FINDING)
- [x] measured true lag (fetch + count): **clones already current** — next-web @ v2.89.0 (2 behind, ff'd), all others 0 behind
- [x] `make pull` — trivial fast-forward (next-web +2); NO rebuild needed for the code (per the "build part only" decision)
- [x] recorded: the M201 "115 behind / v2.33.2" premise does NOT hold; the AI-readiness feature is present in `app` v1.315

### S4 — Recapture snapshot from current prod ✅ (content) / ⏳ (taxonomy)
- [x] directus recaptured: 14 tables / 11,986 rows, public-only=true, primary-read (over the wired MCP DSN, sslmode-normalized)
- [x] sim-embeddings recaptured: 4 tables / 1,490 rows
- [⏳] taxonomy recapture running in background (~1.4 GB vectors); both schema digests UNCHANGED → clean in-place refresh
- [x] no version bump needed (digests unchanged) → M52 batch-cache key unaffected

### S5 — Re-validate the M201 false-negatives ✅
- [x] member-AI-readiness CONFIRMED PRESENT in current `app` (v1.315 + the ai-readiness next-web UI) → M201's false-negative was a stale-at-the-time read; resolved
- [x] feeds M48 (document the feature) + M51 (seed the showcase org)

### S6 — Doc: snapshot-cold-start.md (Delivers →) ✅
- [x] documented the M47 update: the MCP's *configured DSN* is a usable `primary-read --dsn` (sslmode auto-normalized) — turnkey cold-start; the MCP *tool* still isn't a COPY source (the M20-D4 nuance preserved). Added Option 2b (values-blind wired-DSN invocation). Resolves KB-47-01.

## Notes
- **Section order rationale:** S2 (code) before S4 (recapture) because recapture *uses* the new auto-capture. S3
  (re-sync) is operational + the ⚠ release risk — checkpoint with the user before kicking off the multi-service
  Docker rebuild on the single demo machine.
