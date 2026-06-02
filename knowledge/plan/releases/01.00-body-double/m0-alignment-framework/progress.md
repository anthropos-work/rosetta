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
