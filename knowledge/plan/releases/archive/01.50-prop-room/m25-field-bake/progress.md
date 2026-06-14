# M25 — Progress

**Status:** **ALL 5 DONE-BARS GREEN** (resume attempt 2). The operator sanctioned the prod-read DSN; the
first real capture surfaced + **fixed a real safety bug** (M25-D5: the `directus_files` closure over-captured
158 tenant-referenced files; firewall caught it fail-closed) plus two dangling-FK bugs (M25-D6 group / M25-D7
folder+uploaded_by+modified_by). With the fixes the structure-bearing capture **passes the firewall**, the
local Directus **serves** (DB-1 + DB-2 proven by curl on the offset ports), the cold-start capture was
exercised (DB-4), and re-run idempotency + teardown stay GREEN (DB-3, DB-5). **Shape:** section.

## Section checklist
_The 5 live done-bars from `overview.md` § Scope — each a real stack run on this box, binary pass/fail._

- [x] **DB-1 — Fresh `/demo-up`**: **GREEN** — stack UP (`--no-ui`, ~10 GiB VM), taxonomy real (329,859),
  data plane LOCAL (`DIRECTUS_BASE_ADDR=http://directus:8055`) + asset plane PROD (`content.anthropos.work`).
  With the structure-bearing, firewall-clean re-capture (M25-D5) and the three dangling-FK nulls (M25-D6
  group / M25-D7 folder+uploaded_by+modified_by), the directus replay now **exits 0**: structure
  auto-provisions, digest converges, rows load, the per-stack Directus **boots + SERVES** on offset 18055.
  Evidence: `curl` the local Directus REST API → 200 + a real published row; image URL resolves to
  `content.anthropos.work` (asset plane prod). (3 real field bugs found + fixed inline en route.)
- [x] **DB-2 — `/dev-up 2 --local-content`**: **GREEN** — dev-2 (offset 20000, per-stack Directus on 28055)
  brought up with `--local-content`; the **same `dev-setdress.sh` engine** proven on demo-1 EXECUTES the
  per-stack Directus → directus replay exits 0 → it SERVES (curl `/items/simulations` → 200 + real published
  rows on the dev offset port). **N=0 stays prod-read** (verified, M25-D4): dev defaults `LOCAL_CONTENT=0`
  + the hard N=0 set-dress guard — N=0 (`anthropos`) structurally never gets a local Directus from the auto
  flow. (Held the max-2-co-resident line: demo-1 torn down before dev-2.)
- [x] **DB-3 — Re-run idempotency**: **GREEN** — full re-run convergent; migrate no-op ×5, casbin
  "already 248 — skipping", directus bootstrap "already bootstrapped — skipping", taxonomy TRUNCATE-reload →
  329,859 (same), seed idempotent. The directus cache-miss reproduces deterministically (correct idempotency).
- [x] **DB-4 — Cold-start capture**: **EXERCISED + GREEN** (M25-D5) — the sanctioned cold-start fill
  (`stacksnap capture --surface directus --source primary-read --dsn <marco_read>`) ran end-to-end: it
  surfaced the firewall over-capture, the fix closed it, and the **re-run PASSED the firewall**
  (`public-only=true`, `directus_files=1257`, `_structure.sql` present) — the rows-only cache upgraded
  in-place to structure-bearing. This capture IS the DB-4 cold-start path.
- [x] **DB-5 — Clean teardown**: **GREEN** — `/demo-down 1 --purge` reclaims the directus container + frees
  offset port 18055 + releases the registry slot (no stacks registered; 0 demo-1 containers remain).

## Build log

### Session 2026-06-13 (field-bake)
- Phase 0b KB-fidelity: **YELLOW** — 3 parallel audits, all v1.5 behaviors ALIGNED; 2 arg-hint drifts fixed
  inline (`demo-up` fictional `--full`/env-var flags; `dev-up` missing `--local-content`). Report:
  `kb-fidelity-audit.md`.
- Box reconcile: torn down a stale `demo-1` + a crash-looping N=0 sentinel orphan; re-pinned the per-stack
  `stack-demo` ext clone from the v1.3b tag → `prop-room-m25` (the release-under-test tooling).
- **Bug fixed inline (M25-D1):** `/demo-up` build gate aborted deterministically — the offline clerkenstein
  cross-compile (`GOPROXY=off GOSUMDB=off`) couldn't fetch/verify M24's pinned `toolchain go1.25.11`. Added
  `GOTOOLCHAIN=local` (2 sites). Committed authoring `35180c0`; tagged `prop-room-m25`.
- **Resource findings (M25-D2):** the documented 12 GB Docker-VM prereq **fails to boot** on this 16 GB host
  (backed off to 10 GB); the VM **disk filled** (M3 precedent) — pruned 45 GB of build-cache/images to recover.
- **BLOCKER (M25-D3):** the local-Directus **content serve** + the cold-start capture (DB-1/DB-2/DB-4) are
  gated on a **structure-bearing directus capture** that needs an operator-confirmed prod read / staging dump
  — the cache on this box is rows-only (pre-M21). Surfaced to the user. See `decisions.md` M25-D3.

## Open questions
- **RESOLVED.** The blocker (rows-only cache) was cleared by the operator-sanctioned `primary-read` capture.
  That capture surfaced + closed a real safety bug (M25-D5, the `directus_files` tenant-leak the firewall
  caught fail-closed) and two dangling-FK bugs (M25-D6/D7), then **passed the firewall** and made the local
  Directus serve. All 5 done-bars are GREEN.

## M25: Hardening

### Pass 1 — 2026-06-13
Pure test-deepening on the committed fix code (ext repo `stack-snapshot`); no production code
changed. The build's regression tests already covered the central paths (over-capture set-algebra,
require-tenant-filter / closure-must-exclude gates, probe shape, single-col NullColumns render);
this pass deepened the edge/error paths around them without duplication.

**Coverage delta (M25-touched ext packages):**
- `firewall`: 98.0% → **100.0%** statements (the two previously-untested `AssertPlan` reject branches).
- `directus`: **100.0%** (held; the deepening is behavioral — string-level closure composition + tenant-set symmetry).
- `cmd/stacksnap`: 81.6% module total unchanged, but **all M25 fix functions** (`buildPublicSelect`,
  `CopyPublic`, `CountTenantRows`, `buildParentLeakProbe`) are at **100%**; the residual is the CLI
  `main.go` command-wiring, out of M25 scope.
- `capture`: 97.8% unchanged (the residual is non-M25 BuildPlan/Run branches already covered).

**Tests added (ext, commit `1a2fd91`):**
- `firewall/firewall_test.go`: 2 reject-branch tests — a scope-bearing table that also declares a
  referenced-subset tenant filter (rejected via the scope-column branch); a column-less table with a
  tenant filter but NO closure (the tenant filter is meaningless without a closure). Distinct
  diagnostics pinned so the operator fixes the right half.
- `directus/directus_test.go`: 2 behavioral tests — the closure's **string-level** composition (public
  OR-of-INs incl. both resource clauses + the public-sims cover, then `AND NOT (<full tenant predicate>)`,
  so a shared pub+tenant file arrives in the public half then is **subtracted**); tenant-set **symmetry**
  (every public file-ref root has a TENANT-side mirror chased to NOT-public sims; resource the deliberate
  exception). These pin at the real-SQL level what the synthetic set-algebra regression proves abstractly.
- `cmd/stacksnap/adapters_test.go`: 3 unit + 1 fuzz on the `NullColumns` render — the production
  multi-null interleaved case (folder + uploaded_by + modified_by), the **PK-stays-real** contract
  (`id` is never in the null set — a null PK is an unloadable row), nullCols **order-independence**
  (projection follows `cols`), and `FuzzBuildPublicSelect_NullColumns` proving the **COPY load shape**
  (projected-item count == len(cols)) is preserved for arbitrary cols + null-col subsets.

**Bugs fixed inline:** none — the committed fix code held under deepening (the build tested well).

**Flakes stabilized:** none — 3 consecutive clean sequential runs of the 8 new tests; the fuzz
explored ~960k execs with zero crashers written to testdata (the `-fuzztime` `context deadline
exceeded` is the benign worker-shutdown overrun; the seed corpus passes in CI/non-fuzz mode).

**Knowledge backfill:** no KB-worthy findings — the deepened invariants (the closure's
`AND NOT (...)` shape, the PK-never-nulled contract, the complete `NullColumns` set) are already
documented in `decisions.md` M25-D5/D6/D7 and `spec-notes.md`; the tests pin them, they don't reveal
new system truths.

### Stop condition
Stopped after Pass 1: the M25 fix surface is at the coverage ceiling (firewall + all fix funcs 100%),
the Step 2b scan found nothing further worth a non-shallow test, and the flake gate is clean. A second
pass would only add shallow tests. (The brief sanctioned a light harden since the build tested well.)
The tag `prop-room-m25` was moved to the new ext HEAD `1a2fd91`.

## M25: Final Review

_close-milestone consolidation (2026-06-13). Phase 1b deferral re-audit GREEN
(`audit-deferrals/deferral-audit-2026-06-13-m25-close.md`). The rosetta-side M25 surface is 100%
markdown/text (2 skill arg-hint one-liners + milestone-dir artifacts) — the field-bake's CODE fixes
live in the ext repo (`stack-snapshot`/`demo-stack`), already hardened (+8 tests, firewall 98→100%,
0 flakes) and tagged `prop-room-m25 @ 1a2fd91`, so Phases 2/2c/4/8 have no rosetta code/test/lint/bench
surface — the merge is rosetta-side only. Phases 3 (docs) + 5 (decision triage) are the live phases._

**Summary: 3 findings — 1 scope · 0 code-quality · 1 docs · 0 tests · 1-blend+5-archive decision-triage.**
Deferral re-audit GREEN (DEF-M21-02 RESOLVED Fate-1 in its destination milestone; 2 M25 env items fated
fresh).

### Scope
- [x] DB-4 checkbox was `[ ]` while its text read **EXERCISED + GREEN** — checkbox-state nit (the body
  already confirmed done). Flipped to `[x]`. All 5 done-bars now visibly GREEN.

### Code Quality
- [x] No rosetta code surface (the diff is markdown/text only). The ext code fixes were reviewed +
  hardened in their own pass (firewall AssertPlan reject branches, NullColumns render, closure
  composition — all at the coverage ceiling). Nothing to review rosetta-side.

### Documentation
- [x] `corpus/ops/demo/frontend-tier.md` § "The 12 GB Docker-VM prerequisite" lacked the M25 field truth
  that **12 GB fails to boot on a 16 GB host** (~10 GB is the practical ceiling; the full UI tier can't be
  co-resident). Added a field-note callout (#M25-D2). Operator-facing accuracy fix + the Phase-5 blend of M25-D2.
- [x] The 2 skill arg-hint fixes (KB-1/KB-2) verified internally consistent with their SKILL bodies
  (dev-up `--local-content` ×4; demo-up `DEMO_NO_*` env vars ×6). All M25 milestone-doc cross-references resolve.

### Tests & Benchmarks
- [x] No rosetta test/benchmark surface. The ext harden pass (+8 tests, 0 flakes, 3 clean sequential runs,
  ~960k fuzz execs) covers the M25 fix code; out of this merge's scope (separate repo, already tagged).

### Decision Triage
- [x] M25-D2 (16 GB-host VM ceiling) → **BLEND** into `corpus/ops/demo/frontend-tier.md` (the 12 GB-prereq
  section) — done above. A real operator-facing truth.
- [x] M25-D8 / M25-D9 (close re-fates of the full-UI render + dev-2 taxonomy rc=4) → **archive**
  (maintainer-only fate records; the report carries the rationale).
- [x] M25-D1/D3/D4/D5/D6/D7 → **archive** — ext-code / implementation specifics; their *behavioral
  contracts* are already documented in corpus (firewall fail-closed + public-only in `safety.md`;
  structure-bearing capture in `snapshot-spec.md` / `snapshot-cold-start.md`; the exit-0/4 split in
  `directus-local.md`), confirmed ALIGNED by the Phase-0b kb-fidelity audit. No further blend needed.
