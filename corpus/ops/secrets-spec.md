# Stack Secret Provisioning — Spec

**The reference for `rosetta-extensions/stack-secrets/`** — how a dev or demo stack gets every repo's
`.env` written from one secret source, the **secret-coverage DNA** that *lists and keeps listed* which
secret each repo needs, and (most importantly) the **values-blind safety contract** that lets a tool move
secret bytes without ever reading, echoing, or logging one.

> **Scope.** This doc covers the v1.6 "stage door" mechanism: the **source-dir/zip ingestion contract**
> (M27), the **secret-coverage DNA** + the **keep-listed gate** (M27), the **provisioning engine** that
> writes each repo's target `.env` (M28), and the **demo-aware coverage check** (M28). The `/stack-secrets`
> skill (the operator entry point) is M29 (this release). The build-from-stack-dev field-bake is M30. The
> code lives in the gitignored `rosetta-extensions` monorepo (its own git), authored + tagged in the
> authoring copy at `.agentspace/rosetta-extensions/`, then consumed per-stack at a pinned tag
> (`stack-<role>/rosetta-extensions @ <tag>`) — **no platform repo is modified**, and **no `.env` ever
> enters git**.

> **This is the read-side family of the snapshot/seeding specs.** Where [`snapshot-spec.md`](snapshot-spec.md)
> set-dresses a stack's *content* and [`seeding-spec.md`](seeding-spec.md) seeds its *data*, this spec is
> about a stack's *secrets* — the third "make the stack actually run" surface. All three are one-sided
> harnesses in the `datadna` mold (gene → criticality weight → two-metric score → keep-listed diff → the
> `0/1/3` exit contract); secret coverage adds an *engine* (`provision`, like `stackseed`) on top of the
> *DNA* (`check`, like `datadna`).

## For PMs — what it does

Standing up a local Anthropos stack means putting the right API keys and tokens into the right `.env`
files across **six repositories** — Clerk keys for three frontends, the GitHub token that pulls private
code at build, AI-provider keys, the Directus content token, the LiveKit voice pair, and a long tail of
config. Today that is a manual, error-prone hand-copy from `platform/.env` (the old `setup_guide.md`
prose). This tool replaces that: you keep your secrets in **one source folder**, and it **provisions
every repo's `.env` from it in one command**.

Two properties make it safe to trust with secrets:

1. **It is values-blind.** No command this tool runs ever *reads*, *prints*, *logs*, or *stores* a secret
   value. You see key *names* and whether each one is present — never the value itself. The one operation
   that necessarily moves secret bytes (writing a repo's `.env`) copies them straight from your source
   folder to the (git-ignored) target file and nowhere else.
2. **It can't re-arm the production-write path.** The one secret that could leak a demo's writes onto the
   live product (the Directus admin token) is deliberately **left blank** on any non-production stack — the
   tool defers to the same strip the demo bring-up enforces, so it can never undo a closed safety hole.

A **coverage scorecard** (the "secret-DNA") tells you, repo by repo, whether your source folder carries
everything a working stack needs — and a CI-style gate keeps that list honest as the platform's required
keys change.

## For engineers

### The source layout contract (the `zEnvs` / stray-`.env` trap defence)

A secret source — a **directory** or a **`.zip`** — is laid out **by repo**, with each repo's keys in the
exact file that repo reads at runtime:

```
<root>/                              # default: .agentspace/secrets
  platform/.env
  app/.env
  sentinel/.env
  studio-desk/.env
  next-web-app/apps/web/.env         # next-web reads apps/web/.env, NOT the repo root
  ant-academy/code/.env.local        # the exact file Next.js precedence reads (.env is absent)
```

Ingestion is **DNA-driven, not glob-driven** (`source.FromDir` / `source.FromZip` in
`stack-secrets/source/source.go`): the reader is handed the set of `(repo, target_file)` pairs the
secret-DNA declares and opens **exactly** `<root>/<repo>/<target_file>` for each — it **never enumerates
the tree**. That is the structural reason a `stack-dev/zEnvs/` backup mirror (not a DNA repo) or a stray
top-level `.env` is **un-ingestable**: a file that isn't at a declared per-repo target path is invisible to
the reader. A repo whose target file is absent is recorded as **Missing** (its genes fail coverage loudly),
never silently substituted from elsewhere. A zip may wrap the layout in one top dir (`secrets/app/.env`) —
the reader matches on the `<repo>/<file>` **suffix**, so both `app/.env` and `secrets/app/.env` resolve;
encrypted zips (age/gpg) are out of v1 scope and surface a read error rather than being silently skipped.

### The secret-coverage DNA (gene = repo × KEY)

The DNA (`stack-secrets/secretdna/secret-dna.json`, parsed by `secretdna.Load` + `Validate`) is a one-sided
harness in the `datadna` mold — it reuses the gene/criticality-weight/two-metric-score/keep-listed-diff
structure of [`seeding-spec.md`](seeding-spec.md#verifying-a-seed--datadna-the-data-dna-cli-m7b)'s data-DNA,
but is identically one-sided ("does the source carry this repo's required key?"), so there is **no**
source-vs-mirror golden machinery (that belongs to the alignment framework, not here).

A **gene** is one `(repo × required-secret KEY)` pair (`secretdna.SecretGene`):

```
repo, key, target_file, scope (shared|service|frontend|config),
criticality (critical|standard|optional → weight 3/2/1),
status (required | optional | waived-<reason>),
operators [key-present (+ nonempty, format:url|jwt|pk|sk)],   # all values-blind
alias (a family id — genes sharing ONE underlying value), source_hint, note
```

The gene id is `<repo>/<KEY>` (e.g. `studio-desk/CLERK_SECRET_KEY`); ids are unique across the DNA.

**The 6-repo / 55-gene map** (the committed `secret-dna.json`, version `stage-door-m27`, profile `graphql`):

| Repo | Target file | Genes | Notable keys |
|---|---|---|---|
| **platform** | `.env` | 28 | `GH_PAT`, the Clerk pair, `OPENAI_KEY`, the Azure variants, `DIRECTUS_TOKEN`, the LiveKit pair, `ENVIRONMENT`, `PUBLIC_HOST` |
| **app** | `.env` | 5 | `GH_TOKEN` (alias), `STRIPE_SECRET_KEY`, `OPENAI_API_KEY` (repo-local backend env, 46 keys) |
| **sentinel** | `.env` | 2 | `DB_CONNECTION` — the **only** Go repo that ships a `.env.example` |
| **studio-desk** | `.env` | 7 | its own Clerk pair, `AI_*`-prefixed AI keys, `DIRECTUS_TOKEN` |
| **next-web-app** | `apps/web/.env` | 7 | Clerk pair, Azure-OpenAI, `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` |
| **ant-academy** | `code/.env.local` | 6 | Clerk pair, `OPENAI_API_KEY` + `ANTHROPIC_API_KEY` (the `/api/ai/chat` route) |

Status split: **40 required · 8 optional · 7 waived**. Of the required genes, **13 are `critical`** (the
gate denominator) and 27 are `standard`. `Validate()` enforces an **anti-vacuous-100 guard** — a DNA with
no required+critical gene is rejected at load (else `Critical` would score a hollow 100% over zero genes),
the same defence the data-DNA + alignment frameworks carry.

### The per-repo target-file map (where `provision` writes)

Each gene's `target_file` is the exact path, **relative to the repo root**, the key lands in. The
non-obvious ones are pinned because the runtime reads a specific file:

- **`ant-academy` → `code/.env.local`.** Next.js env precedence makes `.env.local` win, and the live repo
  ships **no** `code/.env`; the gene targets `code/.env.local` so a provision lands where the app reads.
- **`next-web-app` → `apps/web/.env`.** The web app reads `apps/web/.env`, not the monorepo root.
- everything else → the repo-root `.env`.

`provision` creates parent dirs as needed (`ant-academy/code`, `next-web-app/apps/web`) and writes `0o600`.

### The hybrid `introspect` source + the keep-listed gate (`diff`)

The required-key set is **not** a uniform per-repo `.env.example` — verified on stack-dev, **8 of 9 Go repos
ship none** (only `sentinel` does). So `introspect` (`secretdna.ReadDeclaredKeys` over
`DefaultHybridSources`) rebuilds the required set from the **union** of:

- `platform/.env_example` — the documented backend wishlist baseline (59 keys);
- `sentinel/.env.example` — the lone Go repo declaring keys this way;
- each frontend's `.env.example` (studio-desk, next-web-app, ant-academy);
- a **curated** set of keys docker-compose injects / passes as a build arg (`GH_PAT`, `PUBLIC_HOST`) that no
  `.env.example` declares.

`diff` reconciles the DNA's genes against that hybrid declared set and follows a **two-tier keep-listed
gate** (M27-D2) — scoped to the DNA's own *tracked-secret universe*, not a 1:1 mirror of every example line
(the example files mix curated secrets with config/wiring noise — Sentry DSNs, PostHog keys, feature flags,
ports):

| finding | meaning | gate |
|---|---|---|
| `unlisted-required` | an **already-tracked** secret is declared for another repo with **no gene there** → coverage would be vacuously green | **exit 1** |
| `unlisted-candidate` | a key the DNA has **never** tracked anywhere — a new secret to curate, or config noise | triage (exit 0) |
| `undeclared-gene` | a DNA gene whose key no hybrid source declares (a repo-local-only key, an alias member, or a stale gene) | informational (exit 0) |

The DNA stays **hand-curated**: `diff` surfaces drift, it **never** auto-promotes a candidate into a gene
(`introspect --write` refreshes only the provenance line, never the gene set). This is the device that keeps
the catalog honest as the platform's required keys change — the anti-vacuous-green guard.

### Alias families vs distinct-similar values (the collision rules)

Two key-naming realities the DNA **encodes** (it does not invent them — they are how the repos already name
their keys):

- **Alias families** — one underlying value written under many per-repo keys. The DNA declares them with a
  shared `alias` id; `provision` sources **one** value and writes it under **every** member's key. The one
  shipped family is **`gh-token`**: `platform/GH_PAT` ≡ `platform/GH_ACCESS_TOKEN` ≡ `app/GH_TOKEN` (3
  members — `ValidateAliases` rejects a 1-member "family"). The provisioner resolves the family's value from
  the first member present in the source (`resolveAliasSources`), so a member whose own key is absent from
  the source is still provisioned from a sibling's value.
- **Distinct-similar keys** — keys that *look* like aliases but **may hold different tokens**. These are
  **standalone genes** (no `alias`), each carrying a `do NOT auto-alias` note: `OPENAI_KEY` vs
  `OPENAI_API_KEY`; the Azure variants (`AZURE_OPENAI_KEY` / `AZURE_API_KEY` / `AI_AZURE_KEY`);
  `ANTHROPIC_API_KEY` vs `studio-desk/AI_ANTHROPIC_API_KEY`. **The LiveKit key/secret pair is the sharpest
  case** (M28-D1): `LIVEKIT_API_KEY` and `LIVEKIT_API_SECRET` are a credential **pair** holding **two
  distinct values** — they are *not* an alias family (the alias mechanism means "one underlying value",
  which a key+secret pair is not), so each is sourced by its own key, never alias-copied.

### The waived class (a conscious decision, not a hole)

Seven genes are **waived** — excluded from the coverage denominator with a per-gene rationale in `note`, so
local-vs-prod realities never poison the score:

| Waived class | Genes | Why |
|---|---|---|
| `waived-aws-mount` | `platform/LIVEKIT_RECORDING_AWS_ACCESS_KEY_ID` | AWS recording creds are mounted from `~/.aws/credentials`, never a `.env` secret |
| `waived-profile-gated` | `platform/BREVO_KEY` | only needed under the `messenger` docker-compose profile, not the default `graphql` profile |
| `waived-optional` | `platform/BUNNY_STREAM_API_KEY`, `app/TAILSCALE_AUTH_KEY`, `studio-desk/GCLOUD_SERVICE_ACCOUNT`, `studio-desk/YOUTUBE_API_KEY`, `next-web-app/BUNNY_CDN_TOKEN_KEY` | example-only / absent from live / convenience — a local stack comes up without them |

A waived gene names **no operators** and is never measured (`Validate` enforces this). Because the catalog is
profile-scoped to `graphql` (the DNA's `profile` field), the denominator is honest for the default stack;
a different profile would carry a different waived set.

### The provisioning engine (`provision` — the one place secret bytes move)

`provision` (`stack-secrets/provision/provision.go`) writes each repo's target `.env` from the source. Per
`(repo, target_file)` the DNA declares, for each measurable gene:

1. read the source file's `KEY=VALUE` lines (values carried, never surfaced);
2. read the existing target file's **key NAMES** (values-blind — only for copy-if-absent);
3. resolve the value from the source — directly, or via the gene's alias family — and decide:
   **write** / **skip** (already present, no `--force`) / **blank** (strip-on-non-prod) / **missing**
   (source lacks it);
4. **append** the new lines to the target (existing lines preserved verbatim), `0o600`.

The merge is **append-only**: an existing line is never re-read for its value or rewritten, which is what
makes copy-if-absent honest — `provision` can never corrupt or echo a value already in the target. The
value-carrying boundary is a single file (`provision/io.go`): `sourceValues` is *the one function that reads
a value to write it*, the bytes live only in a local map consumed by `writeTargetFile`, and a hard test
(`provision_safety_test.go`) asserts no value ever surfaces in stdout/stderr/an error/a return.

### The `DIRECTUS_TOKEN` non-rearm safety (the highest-risk interaction)

This is the **blocks-release** safety class — the fix16/fix17 lineage. On a non-prod / `--local-content`
stack, the demo/dev **injection override** (`stack-injection/gen_injected_override.py`,
`stack-core/gen_override.py`) **strips** the prod `DIRECTUS_TOKEN` to `""` at compose-emit time, so a demo
can never write the shared prod Directus (see [`safety.md`](safety.md#23-never-write-shared-directus--prod-s3-the-two-highest-risk-vectors)).
`provision` runs **before** that override and **must defer to the strip** — writing a non-empty prod token
into a non-prod stack's base `.env` would re-arm the closed tenant-data-leak path.

The mechanism (`provision.StripOnNonProdKeys`): the Directus write-token family
(`DIRECTUS_TOKEN` / `DIRECTUS_STATIC_TOKEN` / `DIRECTUS_ADMIN_TOKEN` — the same set
[`safety.md` §2.2](safety.md#22-the-3-layer-isolation-guard)'s `PreflightEnv` rejects) is **never
provisioned with a value on a non-prod target**. It is written **blank** (`KEY=`) — exactly the state the
override would force — so the base `.env` and the override agree and the prod-write path is never re-armed.
This is why the DNA marks `DIRECTUS_TOKEN` as **`key-present` only (no `nonempty`)**: a deliberately-blanked
non-prod value must still pass coverage. A **prod** target (N=0 + `--prod`) is reachable only via the
`--force` N=0 path, so the prod token is never auto-touched either.

### The N=0 guard + idempotency (the run-it-twice contract)

- **N=0 guard** (mirrors `stackseed --reset`, see [`safety.md` §2.5](safety.md#25-the-n0-dev-guards-doubled-in-v13-m13)):
  `provision` **refuses the main dev stack (N=0, `anthropos`)** unless `--force` — N=0 holds the operator's
  real source `.env`; auto-provisioning into it could clobber the developer's working secrets. `--force`
  both overwrites existing keys **and** permits N=0.
- **Idempotency** (the [`idempotency.md`](idempotency.md) run-it-twice contract): default behaviour is
  **copy-if-absent** — a second run with the same source **skips** every already-present key and re-blanks
  the strip-on-non-prod keys to the same blank state, so re-running provisions 0 new keys instead of
  duplicating or clobbering. `--force` is the deliberate overwrite. `--dry-run` runs every guard +
  resolution and prints the per-file plan (write / blank / skip / missing key NAMES) **without writing** —
  an honest preview.

### The demo-aware coverage check (`check` / `measure`)

`check` (`secretdna.MeasureForStack`) scores a source against the DNA and exits 1 if **critical coverage <
100%** — `Overall` = Σ(weight·present)/Σ(weight) (criticality-weighted % provisioned), `Critical` =
present-critical/total-critical (unweighted), gate = `Critical == 1.0` — plus a per-repo rollup ("repo X is
short key Y"). It reuses the data-DNA `ratio()` empty-denominator + anti-vacuous-100 guards.

The check is **stack-type-aware** (`--demo`): on a **demo** stack the Clerk credentials are **not** sourced
from the secret dir — Clerkenstein **mints** them at bring-up (PK_DEMO + an `sk_test_<demo>` secret; see
[`clerkenstein.md`](../services/clerkenstein.md)). So a demo's coverage treats the **minted Clerk family**
as satisfied even when the source lacks them (`secretdna.MintedKeys`: `CLERK_SECRET_KEY`,
`CLERK_PUBLISHABLE_KEY`, `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY`, `VITE_CLERK_PUBLISHABLE_KEY`,
`CLERK_WEBHOOK_SECRET`, `CLERK_JWT_KEY`) — otherwise a perfectly-good demo would false-fail on exactly the
keys it is designed *not* to carry. This is a values-blind overlay on `Measure` (presence by gene NAME, never
a value); a **dev** stack still requires the real Clerk keys in the source. The pre-flight `check` is wired
**non-fatally** into `/dev-up` + `/demo-up` (warn standard / fail critical — the
[`verification.md`](verification.md) convention).

### The values-blind safety statement (the inviolable invariant)

**No verb ever reads, echoes, logs, or persists a secret VALUE** — not in stdout, stderr, an error, or any
committed file. Operators see key NAMES + presence only, at most a value's *shape* (a `url`/`jwt`/`pk_`/`sk_`
structural prefix via `ClassifyShape`, the single function permitted to look at a value, which returns a
shape token, never the value). Extraction from a source is name-only (cut on the first `=`); the value half
is discarded the moment a line is parsed. `provision` **moves** secret bytes source→gitignored-target (its
job) but the bytes never leave the value-carrying boundary (`provision/io.go`) except into the target `.env`.
The `secret-dna.json` file stores NAMES only and is **committable** (unlike a `.env`). This mirrors the
platform's values-blind `Guard.PreflightEnv` discipline — the safety clause is stated authoritatively in
[`safety.md`](safety.md#29-secret-provisioning-is-values-blind-and-never-re-arms-the-prod-write-path-v16-m27m28).

### The CLI — `stacksecrets`

```bash
stacksecrets list       --dna secretdna/secret-dna.json                          # the per-repo catalog (required/optional/waived + alias families)
stacksecrets check      --dna secretdna/secret-dna.json --from <DIR|ZIP> [--demo] # score a source; exit 1 if critical < 100% (alias: measure)
stacksecrets introspect --dna secretdna/secret-dna.json --stack-root <dir>        # rebuild the required set from the hybrid source; reconcile
stacksecrets diff       --dna secretdna/secret-dna.json --stack-root <dir>        # the keep-listed gate; exit 1 on an unlisted-required key
stacksecrets provision  --dna secretdna/secret-dna.json --from <DIR> \
                        --stack-root <dir> --stack <name> [--force] [--prod] [--dry-run]   # write each repo's target .env (values-blind)
```

**Exit codes (the `0/1/3` contract, mirroring `datadna`):** `0` ok / covered / no drift / wrote · `1` a
critical key missing, the keep-listed gate tripped, or a write/guard error · `3` usage error.

The **operator entry point** is the [`/stack-secrets`](../../.claude/skills/stack-secrets/SKILL.md) skill,
which builds this binary from a pinned-tag `rosetta-extensions` clone and runs the right verb against a
non-prod stack, values-blind.

## Status

M27 delivers the framework: the source-dir/zip ingestion + the secret-coverage DNA (the 6-repo/55-gene map)
+ the two-tier keep-listed `diff` gate, **113 Go tests** (hermetic, `-race` clean). M28 adds the `provision`
engine (alias-mapped per-file writes, copy-if-absent + `--force`, N=0-guarded, the `DIRECTUS_TOKEN`
non-rearm regression pinned) + the demo-aware `check`, wired non-fatally into `/dev-up` + `/demo-up`
pre-flight (**160 Go tests** total at tag `stage-door-m28`). M29 (this milestone) authors this spec + the
`/stack-secrets` skill + the corpus wiring; the build-from-stack-dev field-bake (assemble a compliant
`.agentspace/secrets` from stack-dev, prove a full bring-up provisions cleanly with `Critical == 100%`) is
M30. The tooling is **values-blind**, **never commits `.env`**, **never writes prod**, and **zero
platform-repo edits**.
