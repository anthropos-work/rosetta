# Alignment Testing

**Status:** canonical · **Last updated:** 2026-06-02 · **Reference implementation:** [`test/alignment/`](../../test/alignment/)

## What this is (and why)

Some of the things we build are **mirrors**: a reimplementation that stands in for an external
engine so a consumer can't tell the difference. The motivating case is **Clerkenstein** (v1.0
milestone M1) — a drop-in mock of Clerk that the platform talks to *thinking it's the real Clerk*.
A mirror is only useful if it's **faithful**, and "faithful" has to be *measured*, not asserted.

**Alignment testing** is how we measure it. An alignment test is a **differential test**: it feeds
the *same* input to two engines — the **source** (the canonical one) and the **mirror** — and checks
that they behave the same. Run across a complete, enumerated set of behaviors, the pass rate becomes
a single **alignment score (0–100%)** that says how faithfully the mirror reproduces the source.

This is a **third class of test**, beside the two everyone already knows:

| Class | Question it answers |
|-------|---------------------|
| **unit** | Does this component behave correctly in isolation? |
| **integration** | Do these components work together? |
| **alignment** | Do *two independent implementations* behave identically for the same input? |

The framework is **engine-agnostic and reusable** — it lives in rosetta and knows nothing about
Clerk. Clerkenstein is just its first consumer.

## Vocabulary

- **Target** — an engine exposing a surface. There are two roles:
  - **Source target** — the canonical engine being mirrored, pinned to a version (e.g. Clerk's
    `clerk-sdk-go/v2 @ v2.6.0`).
  - **Mirror target** — our reimplementation (e.g. Clerkenstein).
- **Capability** — one endpoint/function of the source surface. *(Axis 1.)* e.g. `CreateOrganization`.
- **Variant** — one input/scenario class for a capability: the standard case, corner cases, error
  cases. *(Axis 2.)* e.g. `duplicate-name`, `max-length`, `unicode`.
- **Gene** — one **(capability × variant)** pair: the atomic unit of alignment. Its **id** is
  `<Capability>/<variant>` (e.g. `CreateOrganization/duplicate-name`) and is the join key across the
  DNA, the goldens, the outcomes, the test subtests, and the report.
- **Alignment DNA** — the officially-enumerated **complete set of genes** for a source target at a
  version. It is the score's *denominator*: it defines what "faithful" means. Each gene is one entry.
- **Alignment score** — the weighted percentage of genes whose mirror outcome matches the source
  outcome. 100% = behaviorally indistinguishable across the entire DNA.

The genome metaphor is deliberate: the DNA is the genome, each gene is one inheritable behavior, and
the score is how much of the genome the mirror reproduces.

## The alignment test class

Alignment tests are **marked** so they can be parsed, listed, counted, and run as their own suite —
separate from unit and integration. In the Go reference implementation:

- A **build tag** `//go:build alignment` gates the suite. Plain `go test ./...` skips it (the same
  separation unit/integration tests use); `go test -tags alignment ./...` runs it.
- The suite reports **one subtest per gene**, named by gene id, so `go test -tags alignment -json`
  maps results 1:1 back to the DNA — that's the "countable / parseable" property.
- The test reuses the same comparison core as `alignctl` and asserts a **configurable gate**
  (`critical == 100% AND overall ≥ THRESHOLD`), logging every gene as it goes so a tolerated
  divergence never disappears silently.

There are therefore **two surfaces that compute the same score** over the same core:

1. `go test -tags alignment` — the developer/CI ergonomic (native runner, per-gene subtests).
2. `alignctl run` — the engine-agnostic orchestrator used by skills and automation.

They agree by construction (shared `compare` core + DNA + goldens).

## The DNA format

DNA is **JSON** (stdlib-parseable, jq-friendly, git-diffable, zero-dependency). One file per
source@version. Example (the toy reference's `Greet` capability):

```json
{
  "schema_version": 1,
  "source": { "name": "toy", "version": "v1", "ref": "examples/toy/source" },
  "mirror": { "name": "toy-mirror", "version": "v1", "ref": "examples/toy/mirror" },
  "capabilities": [
    {
      "id": "Greet",
      "description": "format a greeting, normalizing input whitespace",
      "criticality": "standard",
      "variants": [
        { "id": "basic",       "operator": "exact", "input": { "name": "World" } },
        { "id": "empty-name",  "operator": "exact", "input": { "name": "" } },
        { "id": "padded-name", "operator": "exact", "input": { "name": "  El  Nino  " } }
      ]
    }
  ]
}
```

Field reference:

| Field | Meaning |
|-------|---------|
| `capability.criticality` | `critical` / `standard` / `optional` → default gene weight `3` / `2` / `1`. Also feeds the separate **critical %**. |
| `variant.operator` | The equivalence test (below). |
| `variant.input` | The input passed to the runner for this gene (inlined JSON). |
| `variant.normalize` | Dot-paths zeroed before comparison (only for `operator: normalized`). |
| `variant.weight` | Optional explicit weight; overrides the criticality default. |

`alignctl dna validate` enforces the contract: capability and variant ids must be safe path
segments — `^[A-Za-z0-9][A-Za-z0-9_-]*$` (PascalCase capability, kebab-case variant; no path
separators, so a gene id can never escape the golden dir), gene ids must be unique, `normalized`
genes must list `normalize` paths, and an explicit `weight` must be `1..1000000` (bounding the score
sum). The golden IO independently refuses any path that would resolve outside `--golden-dir`.

## Equivalence operators

Each gene declares how source and mirror outcomes are compared:

| Operator | Aligned when… |
|----------|---------------|
| `exact` | canonical-JSON-equal value **and** equal error class. |
| `shape` | same JSON structure (keys present + value *types* match); values ignored; error class must match. |
| `normalized` | `exact` after deleting the gene's `normalize` paths (for generated ids / timestamps). |
| `error_class` | only the error class is compared (both error the same way / both succeed); value ignored. |

Value comparison is canonical and **precision-preserving** (large integer IDs compare exactly, not
lossily via float).

## The outcomes protocol and the runner

`alignctl` is **engine-agnostic: it never imports the engine under test.** The contract between the
framework and an engine is a small executable, the **runner**:

> A runner is invoked as `RUNNER --target {source|mirror} --dna PATH` and prints, to stdout, a JSON
> map of gene id → outcome:
> ```json
> { "Greet/padded-name": { "value": "Hello, El Nino!", "error_class": null },
>   "Add/overflow":       { "value": null, "error_class": "overflow" } }
> ```
> `value` is the capability's normalized return (any JSON, or `null` on error); `error_class` is a
> stable short string naming the error kind (or `null` on success).

Each engine ships exactly one runner. The toy's is `examples/toy/cmd/toyrun`; **Clerkenstein ships
its own** in its own repo. That's the whole integration surface.

## Record / replay (golden capture)

A real source like Clerk is a **live SaaS** — you can't hit it freely, offline, or deterministically.
So the framework **records once and replays forever**:

- `alignctl capture` runs the runner with `--target source` and writes each gene's outcome to a
  **golden** file under `golden/<Capability>/<variant>.json`. Run once; commit the goldens.
- `alignctl run` runs the runner with `--target mirror`, loads the goldens as the source side
  (default), and compares — fully offline and reproducible. (`--source live` re-runs the source
  instead, for refreshing goldens or when the source is cheap/local.)

## The score

- **gene weight** = `variant.weight` if set, else the capability's criticality default (`3/2/1`).
- **overall** = `Σ(weight · aligned) / Σ(weight) × 100`.
- **critical %** = aligned critical genes / total critical genes — a plain count ratio
  (*unweighted*, unlike the overall) — reported separately so a mirror can be gated on "no critical
  capability may diverge" independently of the overall number.
- **gate** = `--gate-overall` / `--gate-critical`; `alignctl run` exits non-zero when unmet.

> **Honesty caveat:** the score is only as complete as the DNA. 100% on a thin DNA is hollow — it
> just means "matches across the genes we bothered to enumerate." Two things keep the DNA honest:
> `/align-dna`'s capability-coverage check (every consumed endpoint is present) and the
> version-bump DNA diff (M1b) that surfaces newly-added source behavior.

## The two skills

The process is driven by two skills (they orchestrate `alignctl` and own the judgment parts):

- **`/align-dna`** — *build & update alignment targets.* Given a source framework + version: pull the
  pinned source, enumerate the **consumed** capabilities and their variants, emit/update the DNA, diff
  the DNA across source versions, and scaffold test/golden stubs from it. This is where the
  capability × variant enumeration — the genome authoring — happens.
- **`/align-run`** — *measure alignment of two targets.* Given a DNA + a source version + a mirror:
  capture or replay the source goldens, run the mirror, compute the score, and surface the divergence
  report. This is the "how close are we?" loop M1 runs to drive the mirror to its gate.

## `alignctl` reference

The executable harness ([`test/alignment/cmd/alignctl`](../../test/alignment/cmd/alignctl)):

```
alignctl run      --dna P --runner CMD [--golden-dir D] [--source golden|live]
                  [--report out.json] [--gate-overall F] [--gate-critical F]
alignctl capture  --dna P --runner CMD --golden-dir D
alignctl dna list     --dna P [--json]
alignctl dna diff     --old P --new P [--json]    # exit 1 when the DNA moved (the drift signal)
alignctl dna validate --dna P
```

## Worked example: the toy reference

[`test/alignment/examples/toy/`](../../test/alignment/examples/toy/) is a self-contained proof: a
`source` engine and a `mirror` engine that match **except for one intentional divergence**
(`Greet/padded-name` — the source normalizes input whitespace, the mirror forgets to). It exists to
prove the framework **catches misalignment**, not merely that it reports green.

```
$ go run ./cmd/alignctl run --dna examples/toy/dna.json \
    --runner "go run ./examples/toy/cmd/toyrun" --golden-dir examples/toy/golden
Alignment: mirror toy-mirror@v1  vs  source toy@v1
Score: overall 86.7%   critical 100.0%   (5/6 genes aligned)

Per capability:
  Add                          3/3  ok
  Greet                        2/3  DIVERGED

Divergences (1):
  FAIL Greet/padded-name  (exact, w2)
       value differs: source="Hello, El Nino!" mirror="Hello,   El  Nino  !"
```

`go test -tags alignment ./examples/toy/...` passes (the toy gate is 80% overall / 100% critical, and
the divergence is a non-critical gene) while logging the tolerated divergence.

## How M1 and M1b consume this

- **M1 (Clerkenstein backend mirror)** runs the loop: `/align-dna` authors the **Clerk DNA**
  (`clerk@2.6.0` genome), then the build drives `/align-run`'s score up to its **exit gate** (100%
  critical / ≥95% overall) by closing diverging genes. The Clerk DNA, goldens, alignment tests, mirror,
  and runner all live in the **`clerkenstein` repo**, not here.
- **M1b (Clerk drift detection)** reuses the framework wholesale: on a Clerk version bump, `alignctl
  dna diff` shows what changed and `alignctl run` re-scores the existing mirror against the new
  source — a CI gate on the alignment score turns a silent break into a flagged, mechanical update.

## Where things live

| In **rosetta** (this framework — reusable) | In the **mirror's own repo** (e.g. `clerkenstein`) |
|---|---|
| `test/alignment/` — `alignctl` + the toy | the mirror engine itself |
| `/align-dna`, `/align-run` skills | the source's DNA (the genome) |
| this doc | the alignment tests + goldens |
| | the engine's runner |

Rosetta never contains a specific mirror's source — it ships the measuring machinery and a toy that
proves it.

## Layout

```
test/alignment/
  cmd/alignctl            run | capture | dna list|diff|validate
  internal/dna            DNA model, load, validate, weight derivation
  internal/outcome        Outcome type + outcomes/golden IO
  internal/compare        the 4 operators + weighted score (divergence detection)
  internal/report         human-readable render (the JSON report is compare.Report marshaled by `alignctl run --report`)
  examples/toy            the self-contained reference example
```

Stdlib-only Go — builds and runs offline.
