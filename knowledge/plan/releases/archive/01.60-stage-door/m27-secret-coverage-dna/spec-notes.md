# M27 — Spec notes

_Technical detail accumulated during build. Stub at scaffold; sections seeded from the overview scope._

## Pre-flight audits — section 1 (scaffold + ingestion + DNA)
- **Phase 0b KB-fidelity (M27): GREEN** — report `kb-fidelity-audit.md`. All three KB-dependency docs
  (`alignment_testing.md`, `seeding-spec.md`, `safety.md`) PAIRED + every load-bearing claim ALIGNED against the
  precedent code (`stack-seeding/dna/` + `cmd/datadna/main.go`). `stack-secrets/` is net-new by design; its corpus
  doc (`secrets-spec.md`) is M29 scope, not a M27 blind area. No fixes needed.
- **Topic→doc→code triples (fast-start for future audits):**
  - DNA framework → `corpus/architecture/alignment_testing.md` → `.agentspace/rosetta-extensions/stack-seeding/dna/{dna,measure,operators}.go`
  - data-DNA precedent (`datadna`) → `corpus/ops/seeding-spec.md` §"Verifying a seed" → `stack-seeding/cmd/datadna/main.go`
  - values-blind / `PreflightEnv` → `corpus/ops/safety.md:156-205` → `stack-seeding/isolation/` (`Guard.PreflightEnv`)

## The `stack-secrets` section layout

Module `github.com/anthropos-work/rosetta-extensions/stack-secrets` (go 1.25, **stdlib-only** — no pgx, no
yaml; the secret-DNA is JSON + the readers are `archive/zip` + `bufio`). Three packages:
- `secretdna/` — the DNA core: `dna.go` (schema + strict Load/Save + Validate), `source.go` (the values-blind
  `Source` interface + `ClassifyShape`), `operators.go` (key-present/nonempty/format ops), `measure.go`
  (two-metric `Score` + per-repo rollup), `introspect.go` (hybrid `ReadDeclaredKeys` + `Reconcile`), `diff.go`
  (the keep-listed gate), `catalog.go` (the `list` renderer), `secret-dna.json` (the committed map).
- `source/` — `source.go`: dir + zip ingestion (`FromDir`/`FromZip`/`Ingest`), `TargetsFromDNA`, `parseEnv`.
- `cmd/stacksecrets/` — the CLI (`list`/`check`(=`measure`)/`introspect`/`diff`), testable `run()` entry.

## Source-ingestion contract (dir + zip)

Default source dir (per the skill in M29): `.agentspace/secrets`. Layout = **by repo**:
`<root>/<repo>/<target_file>` (e.g. `platform/.env`, `ant-academy/code/.env.local`). **DNA-driven** ingestion:
the reader opens EXACTLY the genes' declared targets — it never enumerates the tree, so `zEnvs/` (not a DNA
repo) and stray top-level `.env`/`.env.backup` files are structurally un-ingestable. Zip mode is plain
(unencrypted; age/gpg deferred per the open question) and tolerates one wrapping top dir via suffix match.
`parseEnv` cuts on the first `=`, strips `export `/quotes for the shape probe, and **discards the value**
(local var only). A missing target → recorded `Missing` (loud coverage fail), never substituted.

## The secret-DNA schema

Gene = `repo × KEY`, id `<repo>/<KEY>`. Per gene: `{repo, key, target_file, scope (shared|service|frontend|
config), criticality (critical|standard|optional → weight 3/2/1), status (required|optional|waived-<reason>),
operators [key-present (+nonempty, format:url|jwt|pk|sk)], alias (family id), source_hint, note}`. Key regex
`^[A-Z_][A-Z0-9_]*$`. `Load` rejects unknown fields. `Validate` enforces: version+profile, unique ids, valid
scope/criticality/status, known operators, non-waived names key-present + waived names none, **the
anti-vacuous-100 guard** (≥1 required+critical gene, else Critical scores a vacuous 100% over zero genes), and
**alias-family consistency** (a family needs ≥2 members). Reuses the data-DNA `ratio()` empty-denominator
convention + the two-metric `Score` (Overall = Σ(weight·pass)/Σ(weight); Critical = unweighted over
required+critical) + the `0/1/3` exit-code contract. Committed `secret-dna.json`: **55 genes across 6 repos**
(platform, app, sentinel, studio-desk, next-web-app, ant-academy).

## Hybrid `introspect` source

`DefaultHybridSources(stackRoot)` = `platform/.env_example` (baseline, 59 keys) + `sentinel/.env.example` (the
ONE Go repo that ships one — verified: app/cms/jobsimulation/skiller/skillpath/storage/messenger/roadrunner ship
none) + each frontend's `.env.example` (studio-desk, next-web-app `apps/web`, ant-academy `code`) + a curated
compose/build-arg set (`GH_PAT`, `PUBLIC_HOST`). `ReadDeclaredKeys` reads NAMES only; a missing file is skipped
(partial-clone-tolerant). `Reconcile` produces the two asymmetries (unlisted declared keys, undeclared genes).

## `list` + `diff` "keep-listed" gate — DNA-scoped, two-tier (M27-D2)

`diff`'s gate-fatal `unlisted-required` class fires ONLY when an **already-tracked** secret (a non-waived gene
for some repo) is declared for another repo with no gene there (the real vacuously-green omission). A
never-tracked declared key is `unlisted-candidate` (triage, exit 0 — config noise vs new secret, never
auto-promoted). `undeclared-gene` (a gene's key not in any hybrid source — a repo-local-`.env`-only key, an
alias member, or stale) is informational. Against the real stack-dev: **0 gate-fatal / 101 candidate / 6
undeclared** (the 6 = app's repo-local keys since app ships no `.env.example` + the `GH_ACCESS_TOKEN` alias —
all legit). The build's diff-vs-stack-dev caught 10 real cross-repo omissions in the hand-curated DNA → fixed
Fate-1 by adding the genuinely-needed genes.

## Waived classes + alias families

- **Alias families** (one value → many per-repo keys): `gh-token` (`GH_PAT` ≡ `GH_ACCESS_TOKEN` ≡ app/`GH_TOKEN`),
  `livekit` (API key + secret pair).
- **Distinct-similar, NEVER auto-aliased** (may hold different tokens): `OPENAI_KEY` vs `OPENAI_API_KEY`; the
  Azure variants (`AZURE_OPENAI_KEY` vs `AZURE_API_KEY`); `ANTHROPIC_API_KEY` vs studio-desk's
  `AI_ANTHROPIC_API_KEY`. Each carries a `do NOT auto-alias` note.
- **Waived** (excluded from the denominator, each with a justifying note): `waived-aws-mount` (LiveKit recording
  AWS creds mounted from `~/.aws`), `waived-profile-gated` (BREVO/messenger), `waived-optional` (Bunny CDN,
  GCloud service-account, YouTube, Tailscale).
- **`DIRECTUS_TOKEN` is `key-present`-only** (no `nonempty`): the injection override (fix16/17) strips it to `""`
  on non-prod / `--local-content`, so a blanked value must still pass; M28's `provision` must defer to that strip
  and never re-arm the prod-write path (the highest-risk M28 interaction).
