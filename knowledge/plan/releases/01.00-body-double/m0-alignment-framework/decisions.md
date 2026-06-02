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

## M0-D7 — Gene-id + weight input validation (close-review hardening)
The close-milestone adversarial lens found that gene ids (capability/variant, user-authored in the
DNA) became file paths under `--golden-dir` with no character validation — a crafted id (`..`, `/`,
leading separator) could read/write outside the dir. Separately, an unbounded explicit `weight` could
overflow the score sum.
**Decision:** `dna.Validate` now enforces capability + variant ids match `^[A-Za-z0-9][A-Za-z0-9_-]*$`
(matches the documented PascalCase/kebab grammar AND prevents path traversal) and bounds explicit
weight to `1..1_000_000`. Defense-in-depth: `outcome.goldenPath` independently refuses any path that
resolves outside the golden dir. Both pinned by regression tests (`TestValidateRejectsUnsafeIDs`,
`TestValidateWeightBounds`, `TestGoldenPathContainment`). Also extracted the duplicated canonical-JSON
logic into `internal/canon` (used by both `compare` and the CLI differ).

## Adversarial review (Phase 2c — scenarios considered)
- **Path traversal via gene ids** → fixed (M0-D7): charset validation + golden-dir containment guard.
- **Score integer overflow via huge weights** → fixed (M0-D7): weight upper bound.
- **Malformed runner output / partial JSON / missing genes / empty DNA / duplicate gene ids** → already handled (ParseSet errors surface; missing genes report "no source/mirror outcome"; duplicate ids rejected by Validate).
- **Runner not found / exits non-zero / writes to stderr / empty --runner** → handled (runRunner wraps exec errors; empty runner rejected).
- **Concurrency / purity** → the operators + `Evaluate` are pure (no shared mutable state); safe to call concurrently.
- **Zero critical genes** → `pct(0,0)` returns 100.0 by design (no critical genes ⇒ critical gate vacuously met); intentional, documented behavior.
