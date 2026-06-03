# alignment — the alignment measurement framework (rosetta M0)

`alignctl` measures how faithfully a **mirror** engine reproduces a **source** engine across an
**Alignment DNA** (the enumerated set of `<Capability>/<variant>` *genes*), producing a **0–100%
alignment score**. It is engine-agnostic — it never imports the engine under test; it execs a
pluggable **runner** that speaks the outcomes protocol.

This framework lives in **rosetta** and is reusable. Its first real consumer is **Clerkenstein**
(v1.0 milestone M1), which lives in its **own repo** and ships its own DNA, goldens, and runner —
nothing Clerk-specific lives here.

## Quickstart (the toy reference)

```sh
go test ./...                                  # unit tests
go test -tags alignment ./examples/toy/...     # the alignment test class (the third class)

go run ./cmd/alignctl dna list    --dna examples/toy/dna.json
go run ./cmd/alignctl dna validate --dna examples/toy/dna.json
go run ./cmd/alignctl capture     --dna examples/toy/dna.json \
    --runner "go run ./examples/toy/cmd/toyrun" --golden-dir examples/toy/golden
go run ./cmd/alignctl run         --dna examples/toy/dna.json \
    --runner "go run ./examples/toy/cmd/toyrun" --golden-dir examples/toy/golden
```

The toy mirror carries **one intentional divergence** (`Greet/padded-name`), so `alignctl run`
reports `overall 86.7% / critical 100.0%` and names it — proving the framework *catches*
misalignment, not just reports green.

## How it works

- **DNA** (`dna.json`, JSON): capabilities × variants → genes; each gene has an equivalence
  `operator`, a criticality-derived `weight`, and an `input`.
- **Runner** (engine-supplied): `RUNNER --target {source|mirror} --dna P` prints `{gene id →
  {value, error_class}}`. The toy's is `examples/toy/cmd/toyrun`.
- **Record / replay:** `capture` records the source's outcomes as goldens (run once, commit, replay
  offline — because a real source like Clerk is a live SaaS). `run` compares the mirror against them.
- **Operators:** `exact`, `shape`, `normalized` (ignores listed paths), `error_class`.
- **Score:** weighted % aligned + a separate critical-only %, gateable via `--gate-overall` /
  `--gate-critical`.

Full concepts + reference: [`corpus/architecture/alignment_testing.md`](../../corpus/architecture/alignment_testing.md).

## Layout

```
cmd/alignctl            run | capture | dna list|diff|validate
internal/dna            DNA model, load, validate, weight derivation
internal/outcome        Outcome type, outcomes/golden IO
internal/compare        the 4 operators + weighted score  (divergence detection)
internal/report         human-readable render (JSON report = compare.Report via `alignctl run --report`)
examples/toy            self-contained reference (source, mirror, DNA, runner, golden, alignment test)
```

Stdlib only — no module dependencies, builds and runs offline.
