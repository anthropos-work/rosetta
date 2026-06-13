# M24 — Progress

**Status:** archived (completed 2026-06-13). **Shape:** section.

## Section checklist
_One checkbox per concrete deliverable, ticked as it lands. Sections 1–3 land in `rosetta` (docs);
sections 4–7 land in the `rosetta-extensions` authoring copy (hygiene strand)._

### Rosetta docs (corpus-wide truth-up)
- [x] **§1 — Stale local-Directus corrections** (verified-against-compose): corrected the false local-Directus
  claims in `corpus/architecture/external_services.md` (image 10.10.1 + admin/password + the fictional compose
  snippet + local-container troubleshooting + local-uploads dir — all false; platform compose has no directus
  service, only `cms`'s `DIRECTUS_BASE_ADDR`→prod), `corpus/architecture/service_taxonomy.md` Directus table,
  `corpus/ops/quick_ops.md` ports table. Each now states the prod-read default + points at the v1.5 local tooling.
- [x] **§2 — Known-state / safety / directus-local rewrites**: rewrote the `snapshot-spec.md` known-state block
  (the `--local-content` self-contained path now leads as the converged end-state; the prod-read/exit-4 path is
  the documented fallback; M23 retired from future-tense), finished `corpus/ops/directus-local.md` (status note
  M22→M23, the promised "data-plane cutover (M23)" + referential-closure sections added, M23 moved out of
  future-work, `cms`-only over-claims fixed). `safety.md` §2 verified **already M23-accurate** (landed in M23's
  own close) — investigated, nothing to change (Fate-1: work genuinely complete, not deferred).
- [x] **§3 — Corpus-wide language sweep** (via `/update-knowledge`): swept the "print-only / exit-4 / reads-live-
  from-prod / not-yet-automated" framing across the skills (dev-up SKILL+reference, demo-up, stack-snapshot,
  db-query) + `CLAUDE.md` AND the demo-facing corpus docs the sweep surfaced (demo/README.md, the 3 demo recipes,
  seeding-spec.md) — each now presents `--local-content`-executed (demo default-on / dev opt-in) as the converged
  self-contained state and prod-read/exit-4 as the documented fallback. Surfaced + threaded the real
  `--local-content` flag / `DEMO_NO_LOCAL_CONTENT` env into the skill descriptions (flag↔docs consistency).

### Rosetta-extensions hygiene strand (each small + independently landable)
- [x] **§4 — (a) Go toolchain pin bump** (ext `423bac7`): added an explicit `toolchain go1.25.11` directive to
  all 4 go.mod (alignment, clerkenstein, stack-snapshot, stack-seeding) + tightened the clerkenstein CI workflow
  `setup-go` from floating `"1.25"` to `"1.25.11"`. Lazy — pin only, no rebuild (parse-verified GOTOOLCHAIN=local).
- [x] **§5 — (b) README index-row guard** (ext `d6dd8fc`): `stack-core/corpus_index_guard.py` — for every
  README-bearing corpus dir, every other `*.md` must be referenced (by filename) in that README, else exit 1.
  8 unittest tests (full stack-core suite green, 77 tests). **Dog-food:** the guard surfaced 7 pre-existing gaps
  (5 architecture + 2 ops docs) — all backfilled rosetta-side so the guard passes clean on the live corpus
  (exit 0). Every M24-touched/created doc (directus-local.md etc.) is indexed.
- [x] **§6 — (c) Zero-critical-genes guard** (ext `04de89e`): `dna.Validate` now rejects a DNA with no critical
  gene (the vacuous-100% `pct(0,0)` hole); `compare.Report` gained a `CriticalGenes` count + `GateMet` refuses a
  non-zero critical gate when it's 0 (defence-in-depth). Tested; full alignment suite + vet + gofmt clean; the 5
  shipped DNAs all have a critical gene (live gate unaffected). Corpus: documented in `alignment_testing.md`.
- [x] **§7 — (d) `/project-stats` scope fix** (developer-kit `825cdce`): the `/project-stats` skill is the shared
  developer-kit `stats.sh` (no stats tooling exists in rosetta-extensions). Added `*/stack-*/*` to `PRUNE_PATHS`
  + a general `drop_gitignored` filter on the code-size scan, so the gitignored `stack-*/` platform clones (which
  inflated rosetta's count by ~2M lines / 9,235 files) are no longer scanned. Verified the collapse to the
  gitignore-respecting truth. Landed at the script's real source (the plugin — not a platform repo / not the
  corpus); rationale + a surfaced pre-existing doc-counter limitation in `decisions.md` M24-D3.

## Build log
- **2026-06-13** — full M24 build in one session. Phase 0b GREEN. §1–§3 (rosetta docs): corrected the false
  local-Directus claims (no directus service in the platform compose), converged the known-state + finished
  `directus-local.md` on the M23 cutover, swept the print-only/exit-4/live-from-prod framing across 5 skills +
  CLAUDE.md + 5 demo-corpus docs. §4–§7 (hygiene): Go toolchain pin → go1.25.11 (ext), the zero-critical-genes
  guard (ext), the README-index-row guard + 7-gap backfill (ext + corpus), the `/project-stats` gitignored-clone
  scope fix (developer-kit). Commits split: rosetta `m24/docs-hygiene`, ext `main` (4 commits), developer-kit
  `main` (1 commit). All 4 hygiene items Fate-1.

## M24: Hardening

### Pass 1 — 2026-06-13 (py corpus-index guard)
**Scope manifest (milestone-touched, hardenable code — ext repo `7e9343a..d6dd8fc`):**
- `stack-core/corpus_index_guard.py` ← `stack-core/tests/test_corpus_index_guard.py` (8 tests)
- `alignment/internal/dna/dna.go` (Validate zero-critical branch) ← `internal/dna/validate_test.go`
- `alignment/internal/compare/compare.go` (GateMet vacuous-critical + `CriticalGenes`) ← `internal/compare/compare_test.go`
- Config-only (no test surface): the 4 `go.mod` `toolchain go1.25.11` pins + `clerkenstein/alignment.yml` CI version (§4); the
  `/project-stats` scope fix (§7) lives in the developer-kit plugin, a separate repo — a small shell script, left as-is per scope.

**Coverage delta (milestone-touched files):** Python — statement coverage already ~100% via stdlib `trace` (no
`coverage`/`pytest` in the py3.14 homebrew env: pip-install of `coverage` fails on a broken `pyexpat` ABI; stdlib
`trace` was the finder). The gaps here were **behavioural, not line-based** — line coverage was green while real
correctness properties went untested.

**Bug fixed inline:** the guard's `name not in readme_text` raw-substring reference test had a **prefix-collision
false-negative** — an unindexed `setup.md` slipped through whenever a different, indexed `dev-setup.md` was present
(`setup.md` is a substring of `dev-setup.md`), i.e. exactly the failure class the guard exists to catch. Fixed with a
**token-bounded** match: the basename must be preceded by a non-(filename-continuation) char, regex-escaped so the
filename's dots are literal. Live corpus still passes clean (exit 0). Commit `191d650 (ext)`.

**Tests added (8 → 16):** prefix-collision regression (fails on pre-fix code, verified), all five legitimate
reference forms still accepted, dot-literal-not-wildcard, `CORPUS_ROOT` env fallback, non-UTF8 README doesn't crash,
deterministic multi-orphan ordering, subdir-without-README scope, empty-README worst case.

### Pass 2 — 2026-06-13 (go zero-critical-genes guard)
**Coverage delta (milestone-touched files):** Go `dna` 98.5% → **100.0%** statements; `compare` 91.1% (steady — its
residual uncovered lines are pre-existing non-M24 paths: `eqValue`/`eqShape` invalid-JSON branches + `truncate`
multibyte, all outside the M24 diff → out of scope per the milestone-touched-only rule). The M24-touched compare code
(`GateMet`, `CriticalGenes`) is fully exercised. gofmt + vet clean; full alignment suite green.

**Bugs fixed inline:** none — pure deepening.

**Tests added (3 dedicated guard tests → +3):** `criticalGenes` counts VARIANTS not capabilities (a critical
capability with zero variants contributes zero critical genes and is doubly malformed — a refactor counting
critical *capabilities* would silently regress); `Load`→`Validate` end-to-end (a syntactically valid zero-critical
DNA file loads but Validate rejects it — the guard isn't bypassed via the JSON path); `GateMet` still gates for real
when critical genes ARE present (below/at/zero-threshold boundaries — the `CriticalGenes==0` escape hatch only fires
with nothing to gate).

**Flakes stabilized:** none observed (3 consecutive clean sequential runs per stack).

**Knowledge backfill:** none warranted. The README-index guard is ext tooling — its prefix-collision rationale lives
in its own docstring (the right home, updated in `191d650`); the corpus does not describe the guard's matching
behaviour, so nothing to correct there. The alignment zero-critical-genes guard IS documented
(`corpus/architecture/alignment_testing.md` §160-166) and that text was already accurate — Pass 2 deepened *tests* of
the documented behaviour and surfaced no reader-facing truth that changes the doc (the variant-vs-capability counting
is an implementation detail). Question asked, answer recorded.

### Stop condition
Loop stopped at **2 passes**: the full Step 2b scan across both guards is exhausted (Pass 1 closed the Python guard's
behavioural gaps + fixed the one real bug; Pass 2 closed the Go guard's edge interactions, dna → 100%), no new
meaningful gaps remain, and zero flakes across 3 sequential runs per stack. Coverage delta on further passes would be
negligible. The Go pin (§4) is config (nothing to test); the `/project-stats` fix (§7) is a small cross-repo shell
script left as-is. Legitimately a light harden, as expected for a mostly-docs milestone.

## M24: Final Review

_close-milestone consolidation (2026-06-13). Phase 1b deferral re-audit GREEN (0 new deferrals, 3 standing
unchanged, 0 repeat/aged — `audit-deferrals/deferral-audit-2026-06-13-m24-close.md`). Phases 1–5 below._

### Scope
- [x] All 7 sections checked off; overview `In:` list fully delivered Fate-1 (3 rosetta doc sections + 4 ext/dk
  hygiene items). No silently-dropped scope, no unaccounted TODO/FIXME. Cross-repo §7 landed at its real source
  (developer-kit `825cdce`), already committed; not part of the rosetta merge.

### Code Quality
- [x] [should-fix] `corpus_index_guard.py` silently `errors="replace"`s a non-UTF8 README — correct (no crash,
  doc still flagged) but unobservable. Emit a one-line stderr warning so an operator knows a README has an
  encoding issue worth fixing. (ext)
- [x] [nice-to-have] `compare.go` `GateMet` + `dna.go` `Validate` defence-in-depth comments could state the
  load-time-vs-scoring-time split explicitly (a reader currently cross-references two files). (ext)
- [x] [confirmed-clean] No must-fix. Go vet/gofmt clean; Python py_compile clean; both modules follow their
  package conventions; regex escaping (`re.escape`) + token-boundary correct; no dead code / leaked handles.

### Adversarial review
- [x] Recorded (no fix): the guard treats a filename mentioned ANYWHERE in the README (even inside a URL or code
  fence) as "referenced" — a deliberate lenient-by-filename contract. Probed live: an incidental URL mention of
  `config.md` yields exit 0. This is the intended scope (catch the real recurring miss — a doc with NO mention —
  without over-policing where the mention sits, robust across the corpus's varied README styles). Not a gap.
  See decisions.md `Adversarial review`.

### Documentation
- [x] [confirmed-clean] Corpus truth-up verified: zero remaining stale local-Directus claims (image tag / fake
  creds / fictional compose snippet all gone + marked false), `--local-content` flag + `DEMO_NO_LOCAL_CONTENT`
  env consistent across skills↔docs, M21–M23 in past tense, all relative md cross-refs resolve, live README-index
  guard exit 0.

### Tests & Benchmarks
- [x] [confirmed-clean] Go alignment full suite OK (vet+gofmt clean); py stack-core 85 OK; corpus-index-guard 16
  tests (matches harden claim); live-corpus guard exit 0; flake gate 5/5 both touched suites; no handbook
  test-count drift. No new test gaps — harden's 2 passes already deepened both guards.

### Decision Triage
- [x] M24-D2 (zero-critical guard) → blend tag: the mechanism is already in `alignment_testing.md` §160-167
  (accurate); add the `(#M24-D2)` reference tag so readers can trace back.
- [x] M24-D1 (Go `toolchain` directive) → archive (maintainer-only — internal tooling-build detail; the
  release-level advisory status lives in state.md's Test-health line, not the platform corpus).
- [x] M24-D3 (`/project-stats` scope fix + the surfaced stats.sh doc-counter limitation) → archive (cross-repo
  tooling detail; the observation IS its own audit trail in decisions.md).
