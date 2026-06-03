# M0 — spec notes (frozen contract)

The source of truth every M0 artifact obeys (harness, toy, skills, doc). Frozen 2026-06-02 before
implementation so nothing drifts on field names / formats.

## Pre-flight audits — Contract section

**Phase 0b — KB-fidelity (M0):** verdict **YELLOW (resolved-by-design)**. Alignment testing is a
documented **blind area** (confirmed: corpus has only unit/integration/e2e vocabulary, zero
"alignment" anchor). Per the build-milestone RED→author-the-doc resolution, M0's `Delivers →
corpus/architecture/alignment_testing.md` line *is* the fill: the doc is authored as the last section
(S5), describing the actually-built contract. No load-bearing stale claims block code (greenfield).
Tracked here; no separate KB-{N} debt.

## Vocabulary
- **Target** — an engine exposing a surface. **Source** = canonical, version-pinned. **Mirror** = our reimplementation.
- **Capability** — one endpoint/function of the source surface *(axis 1)*.
- **Variant** — one input/scenario class for a capability: standard + corner + error *(axis 2)*.
- **Gene** — one (capability × variant) pair. The unit of alignment. **Gene id grammar:** `<Capability>/<variant>` where `<Capability>` is the source's function name (verbatim, usually PascalCase) and `<variant>` is kebab-case. Example: `Greet/unicode-name`. The gene id is the join key across DNA, goldens, outcomes, test subtests, and the report.
- **Alignment DNA** — the enumerated complete set of genes for a source@version; the score's denominator.
- **Alignment score** — weighted % of genes whose mirror outcome matches the source outcome.

## DNA format (JSON — chosen over YAML/Go-structs: stdlib-parseable, jq-friendly, git-diffable, zero-dependency)
```json
{
  "schema_version": 1,
  "source": { "name": "toy", "version": "v1", "ref": "examples/toy/source" },
  "mirror": { "name": "toy-mirror", "version": "v1", "ref": "examples/toy/mirror" },
  "capabilities": [
    {
      "id": "Greet",
      "description": "format a greeting for a name",
      "criticality": "standard",            // critical | standard | optional  → default weight 3|2|1
      "variants": [
        {
          "id": "unicode-name",
          "description": "name with combining marks",
          "operator": "exact",              // exact | shape | normalized | error_class
          "input": { "name": "café" },      // inlined; passed to the runner per gene
          "normalize": [],                   // field paths zeroed before compare (operator=normalized only)
          "weight": null                     // null → derive from capability.criticality
        }
      ]
    }
  ]
}
```

## Equivalence operators (compare source outcome vs mirror outcome for a gene)
- **`exact`** — canonical-JSON-equal value AND equal error_class.
- **`shape`** — same JSON structure (keys present + value *types* match), values ignored; error_class must match.
- **`normalized`** — `exact` after zeroing each path in the gene's `normalize` list (for generated ids / timestamps).
- **`error_class`** — compare only `error_class` (both error the same way / both nil); value ignored.

## Outcome + runner protocol (the engine ⇄ alignctl contract — what makes alignctl engine-agnostic)
A **runner** is any executable invoked as `RUNNER --target {source|mirror} --dna PATH`. It runs each
gene's `input` through the chosen target and prints a **normalized outcomes JSON** to stdout:
```json
{ "Greet/unicode-name": { "value": "Hello, café!", "error_class": null },
  "Add/overflow":       { "value": null, "error_class": "overflow" } }
```
- `value` = the capability's normalized return (any JSON), or null on error.
- `error_class` = a stable short string naming the error kind, or null on success.
Each engine ships ONE runner (the toy: `examples/toy/cmd/toyrun`; Clerkenstein in M1: its own). alignctl never imports the engine.

## Record / replay (reproducible offline — Clerk is a live SaaS)
- **`alignctl capture`** execs `RUNNER --target source` → writes per-gene goldens to `golden/<Capability>/<variant>.json` (the recorded source outcome). Run once; commit.
- **`alignctl run`** execs `RUNNER --target mirror` → mirror outcomes; loads goldens as the source side (default) — fully offline. `--source live` re-execs `RUNNER --target source` instead (refresh / cheap-local-source like the toy).

## Score formula
- gene weight = `variant.weight` if set, else `{critical:3, standard:2, optional:1}[capability.criticality]`.
- `aligned(gene)` ∈ {0,1} from the operator compare.
- **overall** = `Σ(weight·aligned) / Σ(weight) × 100`, 1 decimal.
- **critical %** = aligned critical genes / total critical genes (for M1's "100% critical" gate).
- per-capability rollup: aligned/total.

## Alignment test class (the third class, beside unit & integration)
- **Marked** by Go build tag `//go:build alignment` (plain `go test ./...` skips them — same separation as unit/integration).
- **Countable / parseable:** one table-driven `TestAlignment` ranges over DNA genes; each gene is a `t.Run("<Capability>/<variant>", …)` subtest. `go test -tags alignment -json` → test2json → subtests map 1:1 to genes by name.
- The alignment test reuses the same `compare` core + goldens, and asserts the **gate** (configurable): `critical == 100% AND overall ≥ THRESHOLD`. It logs every gene PASS/DIVERGED. Diverged-but-within-gate → green with a logged warning; critical divergence or sub-threshold → red.

## alignctl CLI surface (MUST match the skills' documented invocations)
- `alignctl run --dna P --runner CMD [--golden-dir D] [--source golden|live] [--report out.json] [--gate-overall F] [--gate-critical F]` → prints score + divergence report; exit 0 if gate met, 2 if not.
- `alignctl capture --dna P --runner CMD --golden-dir D` → writes goldens from source.
- `alignctl dna list --dna P [--json]` → list genes, counts, weights (parseable/countable).
- `alignctl dna diff --old P --new P [--json]` → added / removed / changed genes (for `/align-dna` + M1b drift).
- `alignctl dna validate --dna P` → schema + referential checks (every gene has operator; normalized genes have normalize paths).

## File layout (under `test/alignment/`, matching rosetta's `test/` tooling convention)
```
test/alignment/
  go.mod                         module github.com/anthropos-work/rosetta/test/alignment ; go 1.25 ; NO requires (stdlib only)
  README.md
  cmd/alignctl/{main,run,capture,dna}.go
  internal/dna/dna.go  (+ _test.go)        load/validate DNA, weight derivation
  internal/outcome/outcome.go              Outcome type, outcomes-file + golden IO
  internal/compare/compare.go (+ _test.go) operators + score  ← detection logic lives here
  internal/report/report.go (+ _test.go)   human + --json render
  examples/toy/
    dna.json
    surface/surface.go                     the shared interface + types
    source/source.go                       toysource
    mirror/mirror.go                       toymirror — ONE intentional divergence (Greet/unicode-name)
    cmd/toyrun/main.go                     the toy's runner (--target source|mirror)
    golden/<Capability>/<variant>.json     generated by `alignctl capture`, committed
    alignment_test.go                      //go:build alignment ; TestAlignment over genes ; gate
```

## Toy design (proves end-to-end + proves detection)
- **Add** (critical, w3): variants `two-positives` (exact), `with-negative` (exact), `overflow` (error_class). Mirror matches source on all → critical 100%.
- **Greet** (standard, w2): variants `basic` (exact), `empty-name` (exact), `unicode-name` (exact) — **mirror diverges here** (source NFC-normalizes the name; mirror does not), so this gene diverges.
- Expected: total weight 15; diverged weight 2 → **overall ≈ 86.7%**, **critical 100%**. Toy gate = `critical==100% AND overall ≥ 80%` → `go test -tags alignment` GREEN while still logging + reporting the divergence; `alignctl run` prints 86.7% and names `Greet/unicode-name` in the divergence report. (Consumer gates are stricter — M1 uses 100% critical / ≥95% overall.)
- A `compare` unit test asserts the divergent pair is flagged → detection proven in plain `go test ./...` (stays green).
