# M209 — Progress

Section checklist (closure = all boxes land). **rext code → committed to the rext repo + tagged
`quick-change-m209` (3 commits since `00eef00`); rosetta planning docs → this branch.**

- [x] snapshot: flip `taxonomy/taxonomy.go:43 Schema "skiller"→"public"` + 2 PublicVia test assertions
- [x] snapshot: narrow `pg.SchemaVersionSQL` digest to enumerated surface tables (Risk 1) — via `Surface.VersionTables`; directus keeps whole-schema
- [x] snapshot: verify capture SELECT column list vs merged prod (Risk 2) — VERIFIED unchanged; no code change (tooling is type/opclass-agnostic)
- [x] snapshot: recapture public taxonomy — CODE ready; actual recapture Fate-3 → M211 (no local capture source); ~42,790 MinRows floor added + AssertPublicOnly kept
- [x] seeding: re-point 5 real-SQL files skiller.→public. (24-file dotted swap) + isolation.go + data-dna.json golden
- [x] seeding: rename fake-Conn test string-matchers in lockstep; reword services/ai comments
- [x] small: readiness.sh schema probe + services.sh container probe + up-injected.sh INJECT_SVCS + 5→4 (+ migrate-demo.sh + GUIDE.md); stack-verify Python suite re-grounded 104/104
- [x] build + `go test ./...` GREEN (all 6 Go modules); tag rext `quick-change-m209`; consumption re-pin deferred to M211

## Done-bar status
- rext authoring copy builds + tests GREEN (6 modules): snapshot, seeding, secrets, alignment, clerkenstein, playthroughs. ✓
- **0 `skiller.<table>` queries in any production path** (verified by grep). ✓
- Cache-key digest narrowed (Risk 1). ✓  Capture column list reconciled — no change needed (Risk 2). ✓
- rext tagged `quick-change-m209`. ✓
- Recapture: tooling READY; the DATA op routed Fate-3 → M211 (no local capture source — see decisions.md).
- Consumption re-pin (`.agentspace/rext.tag`) + `v2.1` release roll deferred to close-release/M211 (per plan).
