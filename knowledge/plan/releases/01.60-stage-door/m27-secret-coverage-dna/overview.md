---
milestone: M27
slug: secret-coverage-dna
version: v1.6 "stage door"
milestone_shape: section
status: planned
created: 2026-06-14
last_updated: 2026-06-14
complexity: medium-large
delivers: rosetta-extensions/stack-secrets/ (new section — cmd/stacksecrets + the secret-DNA sub-package); ext tag stage-door-m27
backlog_refs: (none — new feature requested directly by the user 2026-06-14, not from prior backlog)
---

# M27 — Secret-coverage DNA + source ingestion

## Goal
A new `stack-secrets` extension that ingests a secret source (directory **or** zip) and a **secret-coverage
DNA** that *lists and keeps listed* the required secrets per repo — values-blind throughout.

## Why section
The shape is fully known: the source-of-truth files exist on disk (`platform/.env_example`, the frontend +
sentinel `.env.example`, docker-compose), and the DNA/score/diff machinery is a near-copy of the proven
`stack-seeding/dna` (`datadna`) harness. No exploratory uncertainty — concrete construction. Build with
`/developer-kit:build-milestone`.

## Repo split
- **`rosetta-extensions`** (authoring → tag `stage-door-m27` → consume): the new `stack-secrets/` section —
  `go.mod`, `cmd/stacksecrets`, the source-ingestion reader (dir + zip), the secret-DNA sub-package
  (`dna.go`/`measure.go`/`introspect.go`/`diff.go`/`operators.go` + `secret-dna.json`), hermetic tests.
- **`rosetta`**: none in this milestone (docs land in M29).

## Scope
- **In:**
  - New `rosetta-extensions/stack-secrets/` section + `go.mod` + `cmd/stacksecrets`, authored in `.agentspace`, tagged.
  - Source ingestion — directory **and** zip (default `.agentspace/secrets`); values-blind extraction; an explicit
    source-dir **layout contract** so `stack-dev/zEnvs/` and per-repo `.env` files are never silently ingested.
  - The **secret-DNA** sub-package (mirrors `stack-seeding/dna` layout): gene = `repo × KEY`, gene id `<repo>/<KEY>`;
    per gene `{repo, key, scope (shared|service|frontend|config), criticality (critical|standard|optional → weight
    3/2/1), operators [key-present (+optional nonempty, format:url|jwt|pk_*|sk_*)], status (required|optional|
    waived-<reason>), source_hint, note}`. Reuse Criticality 3/2/1, the two-metric Score, `ratio()`, and the
    0/1/3 exit-code contract.
  - `introspect` rebuilds the required set from the **hybrid** source (`platform/.env_example` baseline + each
    frontend's + `sentinel`'s `.env.example` + the keys docker-compose injects/references) — NOT pure per-repo
    `.env.example` (8 of 12 Go repos ship none).
  - `list` + `diff` verbs; `diff` exits 1 on required-key drift (the **"keep-listed" gate**) and on a
    runtime-required-but-undeclared key (the anti-vacuous-green guard).
  - Model the knowns as `waived`: AWS-via-`~/.aws` mount, profile-gated keys (BREVO/messenger, customerio-sync),
    optional Bunny/GCloud. Encode **alias families** (`GH_PAT`≡`GH_TOKEN`≡`GH_ACCESS_TOKEN`) vs distinct-similar
    values (`OPENAI_KEY` vs `OPENAI_API_KEY` — list exact per-repo, do NOT auto-alias).
  - Hermetic unit tests (no values); the DNA file stores NAMES only and is committable.
- **Out:** writing target `.env` files (M28); the coverage gate + bring-up wiring (M28); the `/stack-secrets`
  skill + corpus doc (M29); the build-from-stack-dev validation (M30).

## Depends on
None (first milestone of v1.6).

## Parallel with
None (M28 consumes this milestone's DNA + ingestion reader).

## Estimated complexity
medium-large.

## Open questions
- Zip ingestion mode: extract-to-temp vs in-memory; encrypted-zip (age/gpg) support. **Default:** plain zip + dir
  in v1; encrypted-zip deferred unless requested.
- Profile-tagging on genes vs default-`graphql`-profile scoping — settle here or in M28.
- The canonical declared-required source per repo for keys absent from every `.env.example` but required at runtime
  (curated platform-env list addition + a `diff`-flaggable "undeclared runtime-required" case).

## KB dependencies
- `corpus/architecture/alignment_testing.md` — the DNA framework + the data-DNA precedent (the class to mirror).
- `corpus/ops/seeding-spec.md` — the `datadna`/`stackseed` patterns + the production-isolation boundary.
- `corpus/ops/safety.md` — the values-blind / `PreflightEnv` discipline.

## Delivers →
`rosetta-extensions/stack-secrets/` (the section + the secret-DNA; ext tag `stage-door-m27`).
