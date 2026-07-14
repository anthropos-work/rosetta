# Alignment Testing

**Status:** canonical · **Last updated:** 2026-06-06 · **Reference implementation:** `rosetta-extensions/alignment/` (section of the [extensions repo](https://github.com/anthropos-work/rosetta-extensions); consumed per-stack at a tag)

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

> **The zero-critical-genes guard.** A DNA with capabilities but **no `critical` capability** has a
> *total critical genes* of 0 — and `aligned/total = 0/0` would score a vacuous **100%**, so an
> all-`standard` mirror would clear any `--gate-critical` for free. Two guards close this: `dna.Validate`
> **rejects** a DNA that declares no critical gene (the authoritative load/lint-time gate — every honest DNA
> must name at least one critical capability), and `GateMet` refuses to clear a non-zero critical threshold
> when the report's `critical_genes` count is 0 (the scoring-time defence). The report carries
> `critical_genes` explicitly so a 100% with zero critical genes is self-evidently vacuous rather than
> mistaken for real coverage. (#M24-D2)

> **Honesty caveat:** the score is only as complete as the DNA. 100% on a thin DNA is hollow — it
> just means "matches across the genes we bothered to enumerate." Two things keep the DNA honest:
> the **capability-coverage check** (`alignctl dna coverage` — below) and the version-bump DNA diff (M1b)
> that surfaces newly-added source behavior.

### The capability-coverage check — what it actually guarantees

> ⚠ **This section was rewritten in M218 because the previous version described a check that did not
> exist.** The doc had offered, as *the* named mitigation for a hollow score, "`/align-dna`'s
> capability-coverage check (every consumed endpoint is present)". There was no such check: `alignctl dna`
> was `list | diff | validate`, and the only "coverage" anywhere in the tooling was an **eyeball** step in
> the `/align-dna` skill ("*List to eyeball coverage*"). The cost was exact. `GET /v1/users/{id}` — which
> next-web's server-side `currentUser()` calls on **every authenticated render** — had no capability in any
> of the five DNAs, so Clerkenstein scored **100% critical / 100% overall / 0 divergences while its fake
> BAPI returned the wrong human for every hero**. The goldens ratified the defect: the score was identical
> before and after the fix. A safeguard that exists only in prose is not a safeguard.

The check is now real, and it **binds**:

| | |
|---|---|
| **What declares it** | A DNA may carry a `consumed_surface`: a list of `{endpoint, consumer, capability \| covered_by}` — the endpoints a **real consumer actually calls**, and who calls each one. |
| **What enforces it** | `DNA.Validate()` **rejects** a consumed endpoint that names no capability (or names one that doesn't exist). `alignctl run` calls `Validate` **before it scores anything**, so such a DNA cannot be scored at all — the missing-gene state is *unrepresentable*, not merely undetected. |
| **How you run it** | `alignctl dna coverage --dna PATH` → exit **0** every declared endpoint has a gene · exit **2** an endpoint is uncovered **or the DNA declares no surface at all**. `gate.sh` runs it on every gate, **before** the score — but with **`--if-declared`** (see the row below), which changes what the *gate* enforces. |
| **What the GATE actually enforces** (`--if-declared`) | ⚠ **Not the same as the bare command — do not conflate them.** `gate.sh:61` calls `alignctl dna coverage --dna … --if-declared`. That flag downgrades **exactly one** case — *"this DNA declares no `consumed_surface` at all"* — from **exit 2** to a **loud warning, exit 0**. A DNA that **does** declare a surface and leaves an endpoint **uncovered** still **fails the gate, exit 2, before a single gene is scored.** So: **a declared hole is fenced; an undeclared surface is only warned about.** The flag exists because a deployment/injection DNA has no HTTP surface to declare, and a hard stop there would be noise. |
| **The escape hatch** | `covered_by` names a capability on **another** surface's DNA (e.g. `clerk-express-1:ClerkClientBAPI/get-organization`). It is **not** machine-verified — but it must be *written down*, which is the whole difference from the silence that hid the bug. |

**And what it does NOT guarantee — stated plainly, because over-claiming this is what caused the bug:**

- **It binds the *declared* surface. It cannot discover consumption nobody wrote down.** This is a strictly
  weaker claim than the old "every consumed endpoint is present," and it is the true one. Adding a consumer
  without adding its endpoint to `consumed_surface` re-opens exactly the M218 hole.
- **Only `clerk-2.6.0` (the fake BAPI's HTTP surface — where the failure happened) declares a
  `consumed_surface` today.** The other four DNAs declare none, so the check **does not bind for them**;
  `gate.sh` prints a loud `NO COVERAGE CLAIM` warning on every run for each. An undeclared surface reads as
  **unguarded**, never as clean — that is why the *bare* command exits 2 rather than passing quietly.
  **But the gate passes `--if-declared`, so at the gate that case is a warning, not a stop.** Four of the
  five surfaces are therefore, today, *warned about but not enforced*. Stated plainly because the previous
  version of this very section over-claimed a guarantee the tooling did not provide — and that is the bug
  this whole section exists to document.
- It checks that an endpoint *has a gene*. It cannot check that the gene is a **good** one. A gene whose
  golden encodes the mirror's own bug still passes — see the `universal-user` variants below.

> **The deeper lesson (M218 D15).** Three genes *did* name identity — `ExtractIdentity`, `Me`,
> `DeployIdentity` — and all three assert the variant **`universal-user`**: the stub itself. They stayed
> green *because* the mirror served the stub. **When a golden is captured from a mirror rather than derived
> from the source's contract, it ratifies whatever the mirror does — including its bugs.** The fix is
> two-sided assertions: `GetUser` now asks for hero **A** *and* hero **B** and requires **A ≠ B**, which no
> single stub identity can satisfy. Prefer a gene that *cannot* be satisfied by the failure mode you fear.

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

The executable harness (`rosetta-extensions/alignment/cmd/alignctl`):

```
alignctl run      --dna P --runner CMD [--golden-dir D] [--source golden|live]
                  [--report out.json] [--gate-overall F] [--gate-critical F]
alignctl capture  --dna P --runner CMD --golden-dir D
alignctl dna list     --dna P [--json]
alignctl dna diff     --old P --new P [--json]    # exit 1 when the DNA moved (the drift signal)
alignctl dna validate --dna P
alignctl dna coverage --dna P [--if-declared]     # M218. exit 2 = a consumed endpoint has no gene
                                                  #   (or, WITHOUT --if-declared, the DNA declares no
                                                  #   surface at all). gate.sh passes --if-declared.
```

### The current scores — and the one that is deliberately red

| surface | DNA | score | |
|---|---|---|---|
| Go SDK | `clerk-2.6.0` | **97.2% overall · 100% critical** (26/27 genes) | gate ≥95 / =100 ⇒ **MET** |
| JS/FAPI | `clerk-js-5` | 100% / 100% (9 genes) | |
| multi-identity | `clerk-multi-1` | 100% / 100% (9 genes) | |
| deployment/injection | `clerk-deploy-1` | 100% / 100% (7 genes) | |
| `@clerk/express` | `clerk-express-1` | **UNMEASURABLE** without `@clerk/express` `node_modules` — rc=**2**, **no score** | **not a pass** |

Two things this table is designed to stop you from saying:

1. **"Clerkenstein is at 100%."** The Go surface is at **97.2%**, on purpose. `MembershipOrgIdentity/real-org-eid`
   is a **deliberately RED** `standard` gene (M218 **D16**): the fake BAPI fabricates the org's external id
   instead of returning the roster's real UUID. It could have been made green by **omitting the field from
   the gene** — which is precisely how the *user*-level version of the same stub survived four releases. The
   divergence was therefore printed on **every run** until the fix landed.
   **Prefer a red gene that tells the truth to a green one that doesn't.**
   > ✅ **RESOLVED M219** (`FIX-M219-bapi-org-eid`). The BAPI now reports the roster's real `org_eid`
   > (`Store.SeedOrgIdentity`/`LookupOrgEid`, wired from the roster at `cmd/fake-bapi`), behind a three-tier
   > ladder — roster eid → demo-org eid → the historical stub — so the alignment runner (which mounts no
   > roster) stays byte-identical and **exactly one gene moved**. Go surface: **97.2% → 100.0% / 100%
   > critical, 27/27, no divergences.** The gene **stays in the DNA** as a permanent fence over the whole
   > wiring path (roster JSON → `LoadRoster` → `seedRosterMemberships` → store → `organizationWithEid` → HTTP).
2. **"All five surfaces are measured."** `expressrun` is **dependency-gated**: on a box without the Node
   modules it cannot build and produces **no number at all** — and it exited with the **same code (2)** that a
   real regression uses, so nothing could tell *"we never ran this"* from *"this is fine"*. The express
   surface was recorded as 100% for several releases **having never been run**.
   *Absence of a score is not a passing score.*
   > ✅ **RESOLVED M219** (`TEST-M219-expressrun-dep-gate`). `alignctl run` now splits the codes:
   > **`3` = UNMEASURABLE** (the runner could not execute — **no genes ran, no score exists**) vs
   > **`2` = REGRESSED** (a real, *measured* score below the gate), with a banner that refuses to be mistaken
   > for a pass. `gate.sh` reports the verdict explicitly rather than letting `set -e` blur them. A surface
   > that did not run must be reported as **UNMEASURED** — never carried forward at a stale value.

## Worked example: the toy reference

`rosetta-extensions/alignment/examples/toy/` is a self-contained proof: a
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

## How M1, M1b, M2, and M2c consume this

- **M1 (Clerkenstein backend mirror)** runs the loop: `/align-dna` authors the **Clerk DNA**
  (`clerk@2.6.0` genome), then the build drives `/align-run`'s score up to its **exit gate** (100%
  critical / ≥95% overall) by closing diverging genes. The Clerk DNA, goldens, alignment tests, mirror,
  and runner all live in the **`clerkenstein` repo**, not here.
- **M1b (Clerk drift detection)** reuses the framework wholesale: on a Clerk version bump, `alignctl
  dna diff` shows what changed and `alignctl run` re-scores the existing mirror against the new
  source — turning a silent break into a flagged, mechanical update.
  Mechanized as `alignment/scripts/{gate,drift-check}.sh`; the bump runbook + exit-code contract are in the
  repo's own [`knowledge/alignment.md`](../services/clerkenstein.md) (pointed to from
  [Clerkenstein](../services/clerkenstein.md)).
  > ⚠ **These scripts are run by hand, not by CI (corrected in M218; the correction itself corrected at the
  > M218 close).** This page originally described "a weekly CI workflow in the clerkenstein repo" and said
  > `gate.sh` "CI-gates it alongside the Go DNA." That was false. M218 replaced it with *"rext has no
  > `.github/workflows` at all"* — **which is also false**, and was caught at the close.
  >
  > **The truth, verified:** the workflow **exists** and is **git-tracked**, at
  > `rosetta-extensions/clerkenstein/.github/workflows/alignment.yml`. It **never runs**, because GitHub
  > Actions only reads `.github/workflows` **at the repository root**, and this one sits in a *subdirectory*
  > of the monorepo. The file says so about itself (`:10-11`): *"as a subdir workflow under
  > `clerkenstein/.github` (not at the monorepo root), this is **currently inert** in the rosetta-extensions
  > monorepo; it is kept illustrative of the per-mirror gate shape."*
  >
  > So the gate is a **manual** `/align-run`, and drift is caught only when someone runs it — which is
  > precisely why the alignment blind spot survived four releases undetected. **"Inert" ≠ "absent": a reader
  > who greps for the file finds it, and could reasonably conclude CI covers them.** That it took two
  > attempts to state this correctly is itself the lesson — see *the stale-verdict hazard* in
  > [`../ops/verification.md`](../ops/verification.md).
- **M2 (browser session + webhook)** proves the framework is **surface-generic**: it authors a *second*
  DNA — `clerk-js-5` (the FAPI/browser surface) — with its own runner (`jsfapirun`) and goldens, scored
  by the same `alignctl` to the same gate (100%/100%, 9 genes). Same machinery, a new surface; the
  parameterized `gate.sh` gates it alongside the Go DNA (**by hand — there is no CI**; see the M1b note). See
  [Clerkenstein](../services/clerkenstein.md) (and the repo's `knowledge/architecture.md` for the
  browser↔backend coherence chain).
- **M2c (`@clerk/express` backend session verification)** exercises the framework a **third** time, on the
  Node backend surface: a *third* DNA — `clerk-express-1` (9 genes) — with its own runner (`expressrun`)
  and goldens, scored by the same `alignctl` to the same gate. Its runner drives the **genuine
  `@clerk/express`/`@clerk/backend` SDK** (the *verify-against-the-real-library* discipline, the same one
  `clerk-webhook/` uses with `svix`) rather than a reimplementation — so the score measures whether the real
  SDK accepts Clerkenstein's tokens. It added an **additive RS256/JWKS** path beside the existing HS256
  seams (no migration; M1/M2 gates untouched).
  > ⚠ **This surface is DEPENDENCY-GATED, and today it is frequently UNMEASURED (corrected in M218).** The
  > runner needs `@clerk/express` `node_modules` to build. Without them it exits **rc=2 and produces NO
  > score** — and *nothing in the tooling treats that as a failure*. So on a box that lacks the Node modules,
  > this gate silently contributes **nothing**, while summaries went on reporting "all five surfaces at
  > 100%". **An absent score is not a passing score.** The M218 harden pass could re-measure only **4 of the
  > 5** surfaces for this reason (reproduced identically at the pre-pass baseline ⇒ pre-existing, not a
  > regression). Routed forward as `TEST-M219-expressrun-dep-gate`: a missing dependency must **fail loud**.

- **Deployment / injection (`clerk-deploy-1`, added after M3)** measures a *different kind* of fidelity — see
  the next section. Its runner (`deployrun`) drives the **real platform consumer** (colony) the way `expressrun`
  drives the real `@clerk/express`.
- **Multi-identity seat-switch (`clerk-multi-1`, v1.9 "storytelling" M37)** authors a *fifth* DNA (9 genes,
  runner `multirun`, goldens `golden-multi`) for the multi-session FAPI surface real clerk-js exhibits with
  `single_session_mode=false`: a registry of seeded heroes/orgs + server-authoritative active-seat selection,
  so a demo can present as any seeded hero. Scored by the same `alignctl` to the same gate (100%/100%) — and
  the four existing surfaces stay green through the registry refactor (the single-identity path is
  byte-identical, a one-member registry). See [Clerkenstein](../services/clerkenstein.md) § Multi-identity.

Clerkenstein now drives **five DNAs via five runners** (`clerkrun`, `jsfapirun`, `expressrun`, `deployrun`,
`multirun`) through the one `alignctl` — the clearest evidence the framework is surface-generic.

## What alignment proves — and what it doesn't (the M3 lesson)

The behavioural DNAs (`clerk@2.6.0`, `clerk-js-5`, `clerk-express-1`, `clerk-multi-1`) measure **behavioural
fidelity**: given
an input, the mirror produces the same *outcome* as the source. That is necessary but **not** the same as
**deployability** — that the mock can be *injected into the running platform's exact consumption shape*. v1.0
proved the first and the v1.0 narrative ("a stand-in the platform can't tell apart") implied the second;
**M3 (the first real demo bring-up) was the first test of the second, and it required building things the
alignment never covered** — a vendored `colony` with the disarmed provider in the *concrete* package the
platform builds against (not the standalone `authn.Provider` the Go DNA tested), and runnable fake-server
binaries (the DNAs tested in-process `http.Handler`s). Neither the tool nor a DNA was *buggy*; the **scope**
of the genes didn't cover deployment, and a couple of seams were tested against *idealized interfaces* rather
than the platform's actual ones (the Go authn DNA even pinned a different `colony` version than production).

The fix is a **deployment/injection dimension**, not a patch: the `clerk-deploy-1` DNA + `deployrun` runner
score whether Clerkenstein's injection **artifact** (`deploy/colony-authn`, the disarmed
`colony/authn/provider/clerk` drop-in) **compiles against the platform's real `colony`** and satisfies its
contract (`clerk.NewProvider(...).GetUser` accepts a Clerkenstein token, rejects garbage/expired) — the runner
*compiling* against real colony **is** the contract check, so a colony bump that breaks the seam turns the gate
RED instead of surfacing during a demo. The heavier end-to-end gene (a built platform app accepting a token
over HTTP) is **deployment-gated** — proven once in M3, run where a demo stack exists — like the express gate's
Node dependency. **General principle for any mirror: align both the *behaviour* of the surface and the
*deployability* of the injection artifact, against the consumer's real interface and pinned version.**

## The data dimension (M7b) — alignment applied to *data*, not behaviour

v1.0 exercised the framework on **behaviour** (does the mirror return the same outcome as Clerk?) and M3 added
**deployability** (does the injection artifact compile against the platform?). **M7b adds a third dimension:
DATA** — does a *stack-seeder's output* conform to the platform's *current database schema*? This is the
discipline that keeps the seeder fleet (M7c) from silently breaking when the platform's data model drifts.

The reinterpretation reuses the manifest / score / diff *structure* but flips two things:

| Behavioural DNA (v1.0) | Data-DNA (M7b) |
|---|---|
| Capability = an endpoint/function | Capability = a **seedable surface** (`users`, `memberships`, `casbin-grant`) |
| Source = the live engine; **input → output** (two-sided) | **Source = the platform's live schema, introspected**; **output → schema** (one-sided — no input to compare) |
| Operators = exact / shape / normalized / error_class (**value** semantics) | **Structural operators** — `type-match`, `constraint-satisfied` (NOT-NULL + UNIQUE), `fk-valid`, `row-count` |
| Drift = re-score after a version bump (M1b) | Drift = **schema diff** — an added/removed/changed column or FK flags the catalog stale |

The data-DNA does **two jobs**: (1) it **enumerates the seedable surfaces** — the machine-readable *catalog of
seeders to build* that drives M7c (4 seeded in M7a, ~6 planned); and (2) it is the **conformance gate + drift
detector** — `introspect` captures each seeded surface's shape from a live (migrated) stack as the contract,
`measure` runs the structural operators against a seeded stack (a weighted `Overall` + an unweighted `Critical`,
exactly like the behavioural score), and `diff` re-introspects to flag drift (exit-code contract `0 none / 1
moved / 3 usage`, mirroring `drift-check.sh`).

**Where the analogy holds:** the enumerated genome forces *completeness + drift detection*; the score's two
metrics + the diff machinery carry over; introspection is the record/replay analogue (capture the schema once,
re-introspect to detect change). **Where it breaks:** the operators are structural (FK/constraint/type), not
value-equivalence, and the measurement is one-sided (seeder output vs the schema, not source vs mirror) — which
is why this is a **separate harness** (`rosetta-extensions/stack-seeding/dna/`, the `datadna` CLI), not a new
`alignctl` runner. It connects directly to the stack's Postgres (offset port), never linking the platform.

**Proven live (M7b):** against the M7a-seeded `demo-1`, `measure` reports **100% / Critical 100%** across the 4
seeded surfaces, and `diff` flags an injected column (`added-column`, exit 1) then reads clean again on revert.
The general principle: **a mirror/seeder is faithful only if its *data* conforms to the consumer's *current*
schema — and that conformance is enumerable, measurable, and drift-gated, just like behaviour.**

## The snapshot-fidelity dimension (M9a) — alignment applied to a *replayed snapshot*

M7b measures a **row-by-row seeder's** output against the schema (one-sided). v1.2's snapshot mechanism
([`snapshot-spec.md`](../ops/snapshot-spec.md)) fills the two surfaces M7c **waived** (taxonomy + content) by
**capturing the real public surface from prod once and replaying it per-stack**. That needs a *different* fidelity
question: **does the replayed data reproduce what was captured?** This is **two-sided again** — captured *source*
(recorded in the snapshot manifest) vs replayed *stack* — so it is a genuine alignment reinterpretation, not just
more structural operators.

| Behavioural DNA (v1.0) | Data-DNA (M7b) | Snapshot-DNA (M9a) |
|---|---|---|
| Capability = endpoint | Capability = seedable surface | Capability = a **snapshot surface** (`taxonomy`) |
| live engine; input→output (two-sided) | live schema; output→schema (one-sided) | **captured manifest vs replayed stack (two-sided)** |
| value operators | structural operators | **fidelity operators** (below) |
| drift = re-score on bump | drift = schema diff | staleness = **schema-version mismatch** triggers re-capture |

The snapshot dimension extends the **same** data-DNA harness (`rosetta-extensions/stack-seeding/dna/`, the
`datadna` CLI) rather than spawning a third one — it shares the gene/score/criticality machinery. It adds:

- a new surface **status `snapshot-seeded`** that — unlike `waived` — **counts toward coverage**. A surface M7c
  waived (the snapshot/shared-store hard line) becomes `snapshot-seeded` once a snapshot fills it, so the fleet
  reads **100% coverage with nothing left waived** — the v1.2 thesis (M7c's two waived surfaces lifted to real,
  measured coverage). `Coverage()` counts seeded **OR** snapshot-seeded over the non-waived denominator;
- a **snapshot-fidelity gene class** (`dna/snapshot.go`) — five two-sided operators over a `FidelityProbe` (the
  replayed stack) compared to the captured manifest: **`snapshot-row-count`** (source-vs-replay parity),
  **`snapshot-structural`** (every captured column present after replay), **`snapshot-referential`** (the captured
  surface is referentially closed — every FK's parent table is in the captured set), **`snapshot-embedding-dim`**
  (pgvector columns replayed at the captured dimension — the index was rebuilt, the vectors must carry the same
  width), and **`snapshot-public-only`** (the **provenance gene** — zero tenant-scoped rows after replay, the
  firewall's measured counterpart). A snapshot gene names **snapshot** operators; a structural gene names
  **structural** operators — `Validate` rejects a cross-wire so the two classes never mix.

**Where it breaks from M7b (and why it's a separate gene class, not new structural operators):** the comparison is
captured-vs-replayed (two-sided), the public-only gene asserts a **safety provenance** (no customer data) rather
than a schema property, and the staleness trigger is a **schema-version digest mismatch** (re-capture), not a
column diff (flag-and-fix). The general principle generalizes: **a replayed reference surface is faithful only if
its data reproduces the captured public source — row-for-row, structurally, referentially, at the right embedding
dimension, and with zero tenant leakage — and that fidelity is enumerable and measurable, just like behaviour.**

**Wired to real surfaces (M9b + M10).** The dimension stops being theoretical at M9b: the **taxonomy** surface (the
public skills-taxonomy catalog — formerly the skiller service's; the domain now lives in `app`'s `public` schema) is promoted `waived-m7c → snapshot-seeded-m9b` in `data-dna.json` and carries all five
fidelity operators. **M10** promotes the **content** surface (the public Directus template library)
`waived-m7c → snapshot-seeded-m10`, carrying four operators (no `embedding-dim` — content has no vectors) with the
**public-only gene measured against the per-surface directus predicate** (`private=false AND tenant_id IS NULL AND
status='published'`), not `organization_id`. The two-sided measure is driven by `datadna measure-snapshot`:
`dna.CapturedFromManifest` derives the **source** side from the real snapshot `manifest.json` (per-surface
`PublicFilter` included), `PgFidelityProbe` reads the **replay** side off the live stack, and the gate exits non-zero
if critical fidelity < 100%. With content promoted, **NOTHING is left waived → 100% coverage over the full catalog**.
See [`../ops/snapshot-spec.md`](../ops/snapshot-spec.md#the-directus-content-surface-m10--the-second-real-surface).

## Where things live

rosetta documents the discipline and ships the skills; **all executable machinery — the reusable harness
*and* each mirror — lives in rosetta-extensions** and is consumed per-stack at a tag:

| In **rosetta** (docs + skills, read-only) | In **rosetta-extensions** (executable, consumed per-stack) |
|---|---|
| this doc — the alignment test class + method | `alignment/` — the reusable harness (`alignctl` + the toy) |
| `/align-dna`, `/align-run` skills | each **mirror** section (e.g. `clerkenstein/`) — the mirror engine itself |
| | the source's DNA(s) (the genome — e.g. Clerkenstein ships three) |
| | the alignment tests + goldens + the engine's runner(s) (one per surface — `clerkrun`/`jsfapirun`/`expressrun`) |

rosetta never contains executable alignment code — neither a specific mirror's source nor the reusable
harness. Both are sections of the **rosetta-extensions** monorepo, which carries two clone roles: an
**authoring copy** at `.agentspace/rosetta-extensions/` (spawned on demand — where the `alignment/`
harness, DNAs, goldens, and runners are built, tested, and **tagged**), and **per-stack consumption
copies** `stack-*/rosetta-extensions @ <tag>` (each stack consumes the tooling at a pinned tag). Policy:
all executable stack tooling — the `alignment/` harness, seeders, injection, and each mirror — lives in
rosetta-extensions, built and tagged in the authoring copy, then consumed per-stack; it is never
scattered in the rosetta corpus or authored ad-hoc inside a stack dir. rosetta stays a read-only doc
corpus plus dev-env skills.

## Layout

```
rosetta-extensions/alignment/        (section of the extensions monorepo)
  cmd/alignctl            run | capture | dna list|diff|validate
  internal/dna            DNA model, load, validate, weight derivation
  internal/outcome        Outcome type + outcomes/golden IO
  internal/compare        the 4 operators + weighted score (divergence detection)
  internal/report         human-readable render (the JSON report is compare.Report marshaled by `alignctl run --report`)
  examples/toy            the self-contained reference example
```

Stdlib-only Go (module `anthropos.dev/alignment`) — builds and runs offline.
