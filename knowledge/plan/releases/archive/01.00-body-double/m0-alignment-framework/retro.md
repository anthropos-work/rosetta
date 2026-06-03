# M0 — Retro: Alignment measurement framework

## Summary
M0 delivered a reusable, engine-agnostic **alignment measurement framework** — the first real code
in rosetta (which until now was docs + skills only). `alignctl` (stdlib-only Go, builds/runs offline)
measures how faithfully a *mirror* engine reproduces a *source* engine across an **Alignment DNA**
(capability × variant genes) as a weighted **0–100% score**, with record/replay goldens, a third
`//go:build alignment` test class, and `/align-dna` + `/align-run` skills + a canonical corpus doc.
A self-contained toy proves it end-to-end and proves it *catches* misalignment (86.7% / 100% critical,
flags `Greet/padded-name`). Lifecycle: build S1–S5 → harden (2 passes) → close (4-lens review).

## Incidents this cycle
- **P2 (build-phase, self-caught):** `dna diff` compared raw input *bytes* → reformatting/large-ints
  registered as false drift. Caught by my own drift smoke-test during S3; fixed with canonical,
  precision-safe (`UseNumber`) comparison. Pinned by `TestGeneChanged/reformatted-input`.
- **P2 (build-phase, self-caught):** value comparison via `float64` lost large-integer precision.
  Fixed in S2 (`UseNumber`); pinned by `TestLargeIntExactValue`.
- **P1 (close-review):** **path traversal** — gene ids (capability/variant) became golden-dir file
  paths with no validation; a crafted DNA id could read/write outside `--golden-dir`. Found by the
  close adversarial lens; fixed (charset validation + containment guard, M0-D7) before merge.
- **P2 (close-review):** unbounded gene `weight` could overflow the score sum; fixed (bound, M0-D7).
- 0 test flakes (5/5 random-order gate at close).

## What went well
- **Contract-first build.** Freezing the DNA/score/operator/runner contract in `spec-notes.md` before
  writing code kept the harness + toy + skills + doc coherent — zero drift between them.
- **The toy's intentional divergence** turned out to be the highest-value design choice: it proves the
  framework detects misalignment, not just reports green, on every run + in CI.
- **Adversarial review earned its cost.** The path-traversal must-fix was invisible to the author-side
  review (Phase 2) and the doc-verify workflow; only the dedicated adversarial lens surfaced it.
- **Offline-first (stdlib-only, JSON DNA)** matched the reproducibility goal and made every test
  hermetic — no network, no module deps.

## What didn't
- **Out-of-process CLI tests aren't coverage-instrumented** — the integration test exercised
  `cmd/alignctl` fully but read as 9% coverage, prompting a second harden pass to add in-process
  tests. Lesson: for CLI packages, prefer in-process subcommand calls (capturing stdout) over
  exec-the-binary when coverage signal matters.
- **Initial divergence used a non-stdlib import** (`golang.org/x/text` NFC) — caught immediately, but a
  reminder to check the stdlib-only constraint *before* writing, not after.

## Carried forward
- **None.** All scope landed in M0; no Fate-2/3 routing, no escape-hatch deferrals.
- For M1 (next): the Clerk DNA, goldens, alignment tests, mirror, and runner live in the **`clerkenstein`
  repo** — M1 is the first real consumer of this framework and the first exercise of `/align-dna` /
  `/align-run` against a live SaaS source (the record/replay path becomes load-bearing there).

## Metrics delta
First milestone under the lifecycle — baseline, no prior delta. 45 Go test funcs (3 fuzz); core
coverage 83–98% (dna 98.3 · canon 94.7 · report 96.2 · outcome 92.5 · compare 89.7 · runner 83.3 ·
cmd 67.5); 0 flakes. Full figures: [metrics.json](metrics.json).
