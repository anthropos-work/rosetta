# M0 — decisions

## M0-D1 — DNA format = JSON (not YAML, not Go structs)
The roadmap open question was "YAML vs Go structs." Chose **JSON**: stdlib-parseable (`encoding/json`,
no external dep), jq-friendly, git-diffable, and `/align-dna`-generated (rarely hand-edited, so YAML's
hand-authoring edge doesn't matter). Decisive factor: **zero external dependencies → builds/runs fully
offline**, which is itself a core M0 goal (reproducibility).

## M0-D2 — Reference harness = stdlib-only Go under `test/alignment/`
Language **Go**: matches M1 (Clerkenstein is Go) + the platform, single static binary, and the toy's
alignment tests are then `go test` — directly validating M1's path. **Stdlib-only** (no `require`
lines) so `go build`/`go test` need no network — reproducible offline. Location `test/alignment/`
matches rosetta's existing `test/`-rooted tooling convention (shell + Playwright today; Go alignment
harness joins it). Rosetta thereby gains a small amount of source — justified: M0 is explicitly a
tooling milestone and the framework needs a runnable reference, not just prose.

## M0-D3 — alignctl is engine-agnostic via a pluggable `--runner` command
alignctl never imports the engine under test. The contract is the **outcomes JSON protocol**: a runner
(`RUNNER --target source|mirror --dna P`) emits per-gene `{value, error_class}`. This keeps the
framework generic (the toy ships one runner; Clerkenstein will ship its own) while alignctl owns
orchestration + scoring + reporting + record/replay.

## M0-D4 — Two score surfaces that agree
The score is computable two ways over the same `compare` core + DNA + goldens: (a) `alignctl run` (the
engine-agnostic orchestrator), (b) `go test -tags alignment` (developer ergonomic, per-gene subtests
asserting a configurable gate). Keeping both lets CI use the native test runner while skills/automation
use alignctl; a unit test asserts they don't disagree.

## M0-D5 — Toy carries one intentional, non-critical divergence
`Greet/unicode-name` diverges (source NFC-normalizes, mirror doesn't). Non-critical so critical stays
100% and the toy's lenient gate (≥80% overall) passes — `go test -tags alignment` stays GREEN — while
`alignctl run` and a `compare` unit test both surface the divergence. This proves the framework
**detects + reports** misalignment, not merely that it reports green. (Consumer gates are stricter;
M1 = 100% critical / ≥95% overall.)

## M0-D6 — Release branch cut from `feat/demo-environment`, not `main` (branch-model deviation)
The canonical developer-kit model cuts `release/{version}` from `main`. Here the v1.0 design + planning
docs live on `feat/demo-environment` (2 commits ahead of `main`, not yet merged). Cutting the release
branch from `main` would lose the planning. So: `release/01.00-body-double` ← `feat/demo-environment`,
`m0/alignment-framework` ← release. `feat/demo-environment` → `main` reconciliation is deferred to
release close (or a separate merge), tracked here. No three-fate item — this is branch topology, not scope.
