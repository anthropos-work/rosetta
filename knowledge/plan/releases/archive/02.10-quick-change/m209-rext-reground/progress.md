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

## M209: Hardening

Harden lands in the **rext repo** (gitignored by rosetta), 3 commits on `75bc4cf..HEAD`
(`42ad600` → `72a5259` → `2f06e78`). The M209 code was already line-covered by the build phase
(the ~111 renamed matchers + `go test ./...` GREEN); harden deepened **edge / boundary / wiring /
regression DEPTH** on the two NEW non-mechanical risk items (MinRows floor, VersionTables/Risk-1)
+ pinned the load-bearing schema-const flip against a revert.

### Pass 1 — 2026-07-08 — MinRows floor + VersionTables digest wiring (`capture/`)
**Scope manifest (milestone-touched, harden-relevant):** `stack-snapshot/{taxonomy,pg,capture,cmd/stacksnap}`
(VersionTables, MinRows, SchemaVersion narrowing — high harden value); `stack-seeding/{seeders,dna}`
(75 `public.`-asserting matchers @97.6%, real-SQL re-pointed — well-covered, spot-check only).

**Coverage delta (touched files):** `capture` 98.0% → 98.0% (lines already hit; +6 tests add edge depth).

**Tests added** (`capture/capture_harden_m209_test.go`, 6):
- MinRows off-by-one boundary (rows==MinRows-1 aborts; pins `<` not `<=`)
- MinRows 0-row wrong-schema capture → floor trips, persists nothing (the real trigger)
- MinRows=0 no-floor contract (an unfloored empty table stays valid)
- floor aborts in-loop *before* later tables are read AND *before* the post-loop leak gate (`AssertCaptured`)
- BuildPlan digests EXACTLY `Surface.VersionTables()` (recording fake: row→own tables, structure→nil)

### Pass 2 — 2026-07-08 — Risk-1 anchors on the REAL surfaces + digest determinism
**Coverage delta:** `taxonomy` 100%, `directus` 99.3%, `pg` 53.1% (SchemaVersionSQL builder 100%;
the uncovered `pg.Conn` remainder is the DB-integration layer — integration-only by design, unchanged stance).

**Tests added** (3 files):
- `taxonomy/taxonomy_harden_m209_test.go`: `Surface().VersionTables()` == exactly its 10 enumerated
  tables, in **lockstep** with the capture `Tables` set (the cache-thrash fix on the real surface)
- `directus/directus_versiontables_harden_test.go`: real structure-bearing surface → `nil` (whole-schema
  digest, so a new dynamic collection still invalidates) — the other half of the gate
- `pg/pg_harden_m209_test.go`: SchemaVersionSQL determinism (`ORDER BY sig` survives narrowing),
  narrowed binds only `$1`+`$2` (no `$3`), empty-slice == nil (whole-schema branch)

### Pass 3 — 2026-07-08 — schema-const revert guard + seeding query-shape pin
**Tests added** (2 files):
- `taxonomy/taxonomy_harden_m209_test.go` (+1): assert `Schema` const is **literally** `"public"` +
  every `TableSpec.Schema` / parent-scope / PublicVia label is `public.`-qualified. The existing identity
  check (`s.Schema != Schema`) is **tautological** — both sides are the same const, so it could not catch
  a flip back to `"skiller"`. This is the real regression guard on THE milestone change.
- `stack-seeding/seeders/taxonomy_snapshot_harden_m209_test.go`: the verifier's build-phase `countConn`
  ignored the SQL; a recorder now asserts it counts `public.skills WHERE organization_id IS NULL`
  (never `skiller.`). Closes a real "matcher didn't assert the re-grounded shape" gap.

**Bugs fixed inline:** none — no production bug surfaced (the M209 re-ground was correct as built).

**Flakes stabilized:** none observed. Flake gate: 3 consecutive clean sequential runs of all new tests.

**Knowledge backfill:** no KB-worthy NEW findings. The invariants hardened (MinRows one-sided floor
`rows>=MinRows`; VersionTables digest-scope==capture-scope lockstep; the `CapturesStructure` gate) are
already recorded in `decisions.md` (Risk 1 / MinRows floor entries) + the code comments. The corpus doc
bodies (`snapshot-spec.md`, `seeding-spec.md`, `safety.md`) are **M210's chartered re-point** — editing
them here would collide with that milestone (Fate-2 boundary respected).

### Stop condition
Stopped after **3 passes**: the Step 2b scan surfaced no new meaningful gap (all three orchestrator
focus areas + sub-edges covered — incl. the "empty/missing table in the set" edge, whose caller-side
`replayProbeExit` translation was already deep in the build-phase harden tests); coverage deltas
negligible **by design** (build phase hit the lines; harden added assertion depth); zero flakes.
The only sizeable uncovered surface is the `pg.Conn` DB-integration layer — deliberately integration-only
and unchanged by M209; a live-DB-mutating test was rejected as inappropriate to the read-only tooling ethos.

## M209: Final Review

Close review (Phases 1–5). Build + harden were already GREEN; this cross-cutting pass surfaced **1 finding**
(nice-to-have, pre-existing, routed) and **0 must-fix / should-fix**.

### Scope
- [x] All 8 section boxes checked; done-bar met (0 `skiller.<table>` queries verified; digest narrowed; column
  list reconciled; rext tagged). Recapture Fate-3→M211, KB-1/2/3 Fate-2→M210 — re-confirmed at deferral audit.

### Code Quality
- [x] `go vet ./...` clean on both touched Go modules (stack-snapshot, stack-seeding); 0 orphan `"skiller"`
  schema literals in production Go (grep-verified); the `skiller.*→public.*` swap is internally consistent.

### Documentation
- [x] Corpus doc bodies (snapshot-spec.md / seeding-spec.md / safety.md) + CLAUDE.md catalog + "5 subgraphs"
  are M210's chartered lockstep re-ground — NOT flipped here (Fate-2, would collide with M210). No new
  top-level unit introduced → no new handbook owed. stack-snapshot/README has no count line (nothing owed).

### Tests & Benchmarks
- [x] 6 Go modules GREEN; 5× sequential flake gate clean on the two touched modules; 1763→ actual test funcs.
- [ ] [nice-to-have · routed] **TEST-1:** `stack-seeding/README.md:106` quotes "496 test funcs across 8
  packages"; authoritative `go test -list` = **788 across 13 packages**. **Pre-existing** multi-release drift
  (README last reconciled at M41 / v1.10; the 496→788 gap accumulated across v1.10b + v2.0 + v2.1). **M209 did
  NOT touch this README**, and it lives in the **rext repo which is mandate-frozen at HEAD `2f06e78`** for this
  close (tag `quick-change-m209`→`2f06e78`, clean tree required) — editing it would break the frozen-HEAD /
  clean-tree invariant. **Disposition:** recorded + routed (D-close-2) to the next rext advance (the v2.1 rext
  roll at close-release, or an M211 rext re-tag). Not fixed in-place; not blocking; not M209-caused.

### Decision Triage
- [x] M209 decisions are implementation-specific (schema-flip mechanics, Risk 1/2 resolutions). The
  knowledge-worthy invariants (MinRows one-sided floor; VersionTables digest-scope==capture-scope lockstep;
  the `CapturesStructure` whole-schema gate) blend into `snapshot-spec.md` — but that body-flip is **M210's**
  chartered deliverable (blending now would collide). Decisions stay as archive; knowledge-blend rides M210.
