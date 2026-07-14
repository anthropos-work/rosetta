---
name: align-dna
description: Build or update an Alignment DNA — the enumerated set of (capability × variant) "genes" a mirror engine must reproduce — for a source framework at a pinned version, then diff it across versions and capture source goldens. Authors both BEHAVIOURAL DNAs (the mock matches the source's outcomes) and DEPLOYMENT/INJECTION DNAs (the injection artifact compiles against the consumer's real module + satisfies its contract). Use when starting a new mirror, adding capabilities/variants to measure, adding a deployment surface, or reconciling a mirror after the source library bumps versions. Pairs with /align-run (which scores against the DNA). Full reference: corpus/architecture/alignment_testing.md.
argument-hint: [source framework + version, e.g. "clerk-sdk-go v2.6.0"]
---

# Author the Alignment DNA

The **Alignment DNA** is the genome a mirror must reproduce: the complete, enumerated set of
**genes** (one per `<Capability>/<variant>` pair) for a source engine at a pinned version. It is the
denominator of the alignment score. This skill authors and maintains it. Concepts:
[`corpus/architecture/alignment_testing.md`](../../../corpus/architecture/alignment_testing.md).
Harness: `rosetta-extensions/alignment/` (`alignctl`) — a section of the extensions repo, consumed per-stack at a tag.

> **Where this runs.** The skills + this doc live in **rosetta**; the `alignctl` harness is the
> `rosetta-extensions/alignment/` section, and a specific mirror's DNA, goldens, alignment tests, and
> runner live in **that mirror's own section of rosetta-extensions** (e.g. `clerkenstein/`),
> cloned under the gitignored `stack-demo/`. Author the DNA *in the
> mirror's repo*; pull the source into *its* `.agentspace/`. Note this mirror-repo-LOCAL
> `.agentspace/` (the pinned source-library clone for golden capture) is **distinct** from the
> rosetta-extensions authoring copy at `.agentspace/rosetta-extensions/` — different things, same
> `.agentspace/` parent; leave the mirror-local one intact.

## Two kinds of surface — author BOTH (the M3 lesson)

A mirror has two kinds of fidelity, and each needs its own DNA. **"Your mission" below is the
behavioural path; do the deployment path too for any injected mirror** — skipping it is exactly the gap
that let v1.0 read "100%" while the first real demo (M3) still had to hand-build the injection.

| | **Behavioural DNA** (default) | **Deployment / injection DNA** |
|---|---|---|
| Proves | the mock produces the **same outcomes** as the source, in isolation | the injection **artifact deploys** into the consumer's real shape |
| "Source" | the **live library**, cloned + captured (goldens) | a **hand-authored contract** (what a correctly-injected consumer must do — the M1-D1 hybrid; no live capture) |
| Runner | drives the **mock's** surface | drives the **consumer's REAL interface** at its **pinned version** (the svix / `@clerk/express` / colony pattern) — and *compiling against it IS the contract check* |
| Genes | capability × variant of the source API | the injection contract: the artifact satisfies the consumer's interface + accepts/rejects inputs correctly |
| Example | `clerk-2.6.0`, `clerk-js-5`, `clerk-express-1` | `clerk-deploy-1` (the disarmed `colony/authn/provider/clerk` drop-in vs **real `colony @ v0.34.3`**) |

For a deployment DNA: skip step 1's live-clone (the "source" is your hand-authored expected); in step 8
the runner must **import/build against the consumer's real module at the version the *consumer* pins**
(not whatever the mirror happens to pin — check both); a bump that breaks the interface should fail the
build → RED gate. The heavier end-to-end gene (the real consumer running + accepting the mock's output
over the wire) is **deployment-gated** — document it, run it where the stack exists (cf. the express
gate's Node dependency, the deploy gate's private-module/`GH_PAT` dependency).

## Your mission (behavioural surface; for a deployment DNA see the box above)

1. **Pin the source.** Clone the source library at the exact requested version into the mirror
   repo's gitignored `.agentspace/` (e.g. `clerk-sdk-go @ v2.6.0`). Record name + version + ref —
   they go in the DNA's `source` block.
2. **Enumerate the *consumed* capabilities — not the whole SDK.** Grep the *consumer* (the platform
   code that calls the source) for the functions/endpoints it actually uses; that scoped set is your
   capabilities (axis 1). Enumerating the entire SDK inflates the DNA with genes nobody depends on.
3. **Enumerate variants per capability** (axis 2): the standard case, plus the corner and error
   cases that matter — empty/min, max/oversized, duplicate, missing-required, malformed, the known
   edge cases. Each (capability × variant) is one gene.
4. **Pick an operator per gene** (see the operator table in the doc): `exact` by default; `shape`
   when only structure is guaranteed; `normalized` (with `normalize` paths) when the response
   carries generated ids/timestamps; `error_class` for genes whose point is the error behavior.
5. **Write `dna.json`** (`schema_version: 1`; `source`/`mirror` refs; `capabilities[].variants[]`
   with `operator`, `input`, and — when needed — `normalize`/`weight`). Set each capability's
   `criticality` (`critical`/`standard`/`optional`); it drives default weight and the critical gate.
6. **Validate:** `alignctl dna validate --dna PATH` — fixes every structural error before proceeding.
7. **Declare the consumed surface, then RUN the coverage check — do not eyeball it.**
   For each endpoint a **real consumer actually calls**, add a `consumed_surface` entry:
   `{endpoint, consumer, capability}` (or `covered_by: "other-dna:Cap/variant"` when another surface's DNA
   owns it). Then:
   ```
   alignctl dna coverage --dna PATH     # exit 0 = every consumed endpoint has a gene; 2 = a hole
   alignctl dna list --dna PATH         # sanity-check gene count / weights
   ```
   `alignctl run` also refuses to score a DNA with an uncovered consumed endpoint, so the hole cannot
   silently become a 100%.

   > ⚠ **This step used to read "List to eyeball coverage."** It was the *named* mitigation in
   > `alignment_testing.md` against a hollow score — and eyeballing is not a check. It missed
   > `GET /v1/users/{id}`, consumed by next-web's `currentUser()` on every authenticated render, and
   > Clerkenstein scored **100% while returning the wrong human for every hero** (M218). Declare the
   > endpoint and let the tool fail; a capability the consumer calls but the DNA omits is a blind spot the
   > score will hide, and a human scanning a list will not reliably catch it.

   **Prefer genes a stub cannot satisfy.** A golden captured from the *mirror* ratifies the mirror's bugs —
   three identity genes all asserted the variant `universal-user` (the stub itself) and stayed green
   throughout. Assert **two-sided**: ask for hero A *and* hero B and require A ≠ B.
8. **Ensure the runner handles every capability.** The mirror repo ships a runner
   (`RUNNER --target {source|mirror} --dna PATH` → outcomes JSON). Each capability id in the DNA must
   have a branch in the runner that invokes it with the gene's `input`.
9. **Capture source goldens:**
   `alignctl capture --dna PATH --runner CMD --golden-dir DIR` — records the source's outcome per
   gene so the mirror can be scored offline. Commit the goldens.

## Reconciling after a source version bump (feeds /developer-kit M1b drift detection)

1. Re-clone the source at the new version into `.agentspace/`.
2. `alignctl dna diff --old OLD.json --new NEW.json` (exit 1 = the DNA moved). It reports:
   - **added** genes → new source behavior; add variants + decide operators, extend the runner.
   - **removed** genes → source dropped a capability/variant; remove them (and any dead mirror code).
   - **changed** genes → an operator/weight/input/normalize change; re-confirm intent.
3. Update `dna.json` accordingly, then **re-capture goldens** against the new source.
4. Hand off to `/align-run` to re-score the mirror against the bumped source.

## `alignctl` commands this skill uses

| Command | Purpose |
|---|---|
| `alignctl dna validate --dna PATH` | structural checks (every gene has a valid operator; `normalized` genes have `normalize` paths; no duplicate gene ids; **every declared consumed endpoint names a real capability**). |
| `alignctl dna coverage --dna PATH [--if-declared]` | **the capability-coverage check** — every consumed endpoint has a gene. Exit 0 covered / 2 uncovered **or no `consumed_surface` declared at all** (an undeclared surface is *unguarded*, not clean). `--if-declared` downgrades only the undeclared case to a warning, for shared gate scripts; a declared hole always fails. |
| `alignctl dna list --dna PATH [--json]` | list genes with operator / weight / criticality; count coverage. |
| `alignctl dna diff --old PATH --new PATH [--json]` | added / removed / changed genes between two DNA versions; exit 1 on drift. |
| `alignctl capture --dna PATH --runner CMD --golden-dir DIR` | record source goldens for offline replay. |

## Done when
- `alignctl dna validate` passes; `alignctl dna list` shows every consumed capability covered;
- goldens captured for every gene (0 missing);
- the runner has a branch for every capability id.
Then run `/align-run` to measure the mirror against the DNA.
