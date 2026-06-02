# M0 — progress

Contract frozen in `spec-notes.md` (S1 — done at scaffold). Sections build in dependency order:
harness → toy (these two co-define and must compile/run together) → skills → doc.

## S2 — `alignctl` reference harness (stdlib-only Go) — DONE
- [x] `go.mod` (module `anthropos.dev/alignment`, go 1.25, no requires)
- [x] `internal/dna` — DNA model, load, validate, weight derivation (+ unit tests)
- [x] `internal/outcome` — Outcome type, outcomes-file + golden IO (+ unit tests)
- [x] `internal/compare` — the 4 operators + score formula (+ unit tests incl. divergence-detection + missing-outcome; UseNumber precision-safe)
- [x] `internal/report` — human render (+ unit tests); JSON report = marshaled compare.Report
- [x] `cmd/alignctl` — `run`, `capture`, `dna list|diff|validate`
- [x] `go build ./...` + `go vet ./...` + `gofmt -l` clean + `go test ./...` green
- [x] `README.md`

## S3 — Toy reference mirror (proves end-to-end + proves detection) — DONE
- [x] `examples/toy/surface` — shared interface + types
- [x] `examples/toy/source` + `examples/toy/mirror` (one intentional divergence: `Greet/padded-name` — whitespace normalization; kept stdlib-only, no x/text)
- [x] `examples/toy/dna.json` (6 genes: Add critical×3, Greet standard×3)
- [x] `examples/toy/runner` (shared) + `examples/toy/cmd/toyrun` — the runner (`--target source|mirror`)
- [x] `examples/toy/alignment_test.go` (`//go:build alignment`, `TestAlignment` per-gene subtests, gate 80/100)
- [x] `alignctl capture` generated + committed `golden/` (6 files)
- [x] **VERIFIED:** `alignctl run` → `overall 86.7% / critical 100.0% (5/6)`, names `Greet/padded-name` in divergence report (exit 0, gate met)
- [x] **VERIFIED:** `go test -tags alignment ./examples/toy/...` GREEN (gate met) while logging the tolerated divergence
- [x] **VERIFIED:** `dna diff` fires only on real contract changes (canonical input compare), exit 1 on drift — the M1b signal

## S4 — The two skills — DONE
- [x] `.claude/skills/align-dna/SKILL.md` — build/update DNA, diff across versions, scaffold, capture goldens (drives `alignctl dna …` + capture)
- [x] `.claude/skills/align-run/SKILL.md` — measure alignment, compose score, triage divergence, iterate to gate (drives `alignctl run`)
- [x] both registered (appear in the skill list); flag-consistency VERIFIED both directions (m0-doc-verify lens: ok, 0 findings)

## S5 — Canonical doc + discoverability — DONE
- [x] `corpus/architecture/alignment_testing.md` — vocabulary, test class, DNA format, operators, record/replay, score, the two skills, alignctl ref, the verified toy walkthrough, M1/M1b consumption, repo split, layout
- [x] discoverable: linked from `corpus/architecture/README.md` + CLAUDE.md architecture-docs list
- [x] `last_updated` set; cross-refs valid
- [x] VERIFIED by m0-doc-verify workflow (3 lenses): flag-consistency ok, reproducibility ok (every documented command runs, 86.7% reproduces), fidelity — 1 should-fix + 1 nit (layout-table `internal/report` description + critical-% clarifier) FIXED

## M0: Hardening

### Pass 1 — 2026-06-02
**Scope manifest (milestone-touched Go):** `internal/{dna,outcome,compare,report}` (had unit tests), `cmd/alignctl/{main,run,capture,dna}.go` (NO tests), `examples/toy/{surface,source,mirror,runner,cmd/toyrun}` (only via the tagged alignment test). New unit `test/alignment/` (Go) is documented by its `README.md` (new-unit handbook check: present).

**Coverage delta (per-package, milestone-touched):**
- `internal/dna`: 66.7% → 93.3%  ·  `internal/compare`: 76.9% → 90.8%  ·  `internal/outcome`: 80.6% → 90.3%  ·  `internal/report`: 96.2% (held)
- `examples/toy/runner`: 0% → 83.3%

**Tests added (Pass 1):** dna Load error paths + GeneIDs + Operator.Valid + Criticality.Weight; compare invalid-JSON / shape-arrays / nested-normalize / **large-int exact** (pins the UseNumber precision fix); outcome ParseSet-error + no-slash golden path + bad-golden; runner Invoke/Run + unknown-capability; **3 native fuzz tests** (ParseSet, compareValue, dna.Load — untrusted-JSON boundary, no-panic); the out-of-process CLI integration test (exit codes, gate behavior, error exits).

### Pass 2 — 2026-06-02
Targeted the `cmd/alignctl` gap (out-of-process integration tests aren't coverage-instrumented).
**Tests added:** in-process `dna validate|list|diff` (happy + `--json`) + every subcommand's missing-flag error exit; in-process run/capture e2e against a **temp** golden dir (capture — previously untested; `--source live`; `--source bogus`; gate-unmet; `--report` JSON).
**Coverage delta:** `cmd/alignctl`: 0% → 49% → **68.6%**. Remainder is `main()`'s `os.Exit` dispatch (not in-process testable) + the trivial toy fixtures (source/mirror/surface/toyrun — exercised end-to-end by 4 test files; dedicated tests would be shallow, intentionally skipped).

### Bugs fixed inline
None — no defects surfaced. (The two real bugs of this milestone — the `dna diff` raw-bytes false-positive and the float-precision value compare — were caught + fixed during the S2/S3 build; the new `TestLargeIntExactValue` + `TestGeneChanged/reformatted-input` now pin them as regressions.)

### Flakes stabilized
None — flake gate clean (3/3 consecutive sequential runs of `go test ./...` + `-tags alignment ./...`).

### Knowledge backfill
No KB-worthy findings. Hardening confirmed existing documented behavior (precision-safe canonical comparison, the operators, record/replay) — all already in `corpus/architecture/alignment_testing.md` and `decisions.md`; nothing new to propagate.

### Stop condition
Stopped after Pass 2: core library packages stabilized at 83–96%; the qualitative 6-dimension scan found nothing further worth adding beyond `main()` dispatch + trivial fixtures (not worth shallow tests); 0 flakes. Performance: no SLAs documented for the framework → benchmarks N/A. 28 test/fuzz functions added (3 fuzz); 41 total.
