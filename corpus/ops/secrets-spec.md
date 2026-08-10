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
2. **It can't re-arm the production-write path.** The secrets that could leak a demo's writes onto the live
   product — the Directus write tokens — are deliberately **left blank** on any non-production stack: the tool
   defers to the same strip the demo bring-up enforces, so it can never undo a closed safety hole. **It is
   TWO genes, not one**, and saying "the one secret" hides the second: the DNA declares
   `platform/DIRECTUS_TOKEN` (`critical`·`required`) **and** `studio-desk/DIRECTUS_TOKEN`
   (`standard`·`required`, noted *"SAFETY: same strip-on-non-prod class as platform/DIRECTUS_TOKEN"*) —
   `stack-secrets/secretdna/secret-dna.json:191`, `:659` @ rext `415240f`. The second one matters on its own
   terms: studio-desk's skill-path builder is the surface that *could have written* prod Directus (see
   §2.2's fix16/fix17 scope note in [`safety.md`](safety.md)). The strip set itself is **three key names** —
   `DIRECTUS_TOKEN`, `DIRECTUS_STATIC_TOKEN`, `DIRECTUS_ADMIN_TOKEN` (`stack-secrets/provision/provision.go:50-54`).

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

**The 6-repo / 64-gene map** (the committed `secret-dna.json`, version `fast-build-m256`; its `profile` field still literally reads the retired `graphql` token — see the waived class below):

| Repo | Target file | Genes | Notable keys |
|---|---|---|---|
| **platform** | `.env` | 32 | `GH_PAT`, the Clerk pair, `OPENAI_KEY`, the Azure variants (incl. the M256 `SKILLER_AZURE_OPENAI_KEY`/`_ENDPOINT_URL` pair), `DIRECTUS_TOKEN`, the LiveKit pair, `INVITATION_HMAC_SECRET`, `ENVIRONMENT`, `PUBLIC_HOST` |
| **app** | `.env` | 10 | `GH_TOKEN` (alias), `STRIPE_SECRET_KEY`, `OPENAI_API_KEY` (repo-local backend env, 46 keys) + **the M239 Bedrock cred class** (`AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY` + `AWS_REGION`/`AWS_SESSION_TOKEN`/`CLAUDE_CODE_USE_BEDROCK` — Talk to Data, see below) |
| **sentinel** | `.env` | 2 | `DB_CONNECTION` (**`waived-config`** — compose-injected, see the waived class), `SENTRY_DSN`; the **only** Go repo that ships a `.env.example` |
| **studio-desk** | `.env` | 7 | its own Clerk pair, `AI_*`-prefixed AI keys, `DIRECTUS_TOKEN` |
| **next-web-app** | `apps/web/.env` | 7 | Clerk pair, Azure-OpenAI, `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` |
| **ant-academy** | `code/.env.local` | 6 | Clerk pair, `OPENAI_API_KEY` + `ANTHROPIC_API_KEY` (the `/api/ai/chat` route) |

Status split: **44 required · 12 optional · 8 waived**. Of the required genes, **13 are `critical`** (the
gate denominator) and 31 are `standard`. `Validate()` enforces an **anti-vacuous-100 guard** — a DNA with
no required+critical gene is rejected at load (else `Critical` would score a hollow 100% over zero genes),
the same defence the data-DNA + alignment frameworks carry. (The M30 field-bake reclassified
`sentinel/DB_CONNECTION` from critical/required to `waived-config` — it is compose-injected, never read from
a `.env`; this shifted the split from 40/8/7 + 13-critical to 39/8/8 + 12-critical. **M49 #4** then added
`platform/INVITATION_HMAC_SECRET` as critical/required — the `app` exits early when it is unset
(`invitations.NewTokenManager` errors and `main` returns: the silent `app Exited (0)` class) — landing the
split at 40/8/8 + 13-critical. **M239** then added the 5-gene **Bedrock cred class** for `app` (2 required-`standard`
+ 3 optional; deliberately **NOT** critical — see below), landing the split at 42/11/8 + 13-critical.
**M256 iter-21** then added 3 `platform` genes — the `SKILLER_AZURE_OPENAI_KEY` / `SKILLER_AZURE_OPENAI_ENDPOINT_URL`
pair (required-`standard`; they gate **every** taxonomy write, and `standard` rather than `critical` was
measured, not assumed) plus optional `SKILLER_OPENAI_KEY` — landing the split at **44/12/8 + 13-critical**
at version `fast-build-m256`; the anti-vacuous guard still holds.)

### The per-repo target-file map (where `provision` writes)

Each gene's `target_file` is the exact path, **relative to the repo root**, the key lands in. The
non-obvious ones are pinned because the runtime reads a specific file:

- **`ant-academy` → `code/.env.local`.** Next.js env precedence makes `.env.local` win, and the live repo
  ships **no** `code/.env`; the gene targets `code/.env.local` so a provision lands where the app reads.
- **`next-web-app` → `apps/web/.env`.** The web app reads `apps/web/.env`, not the monorepo root.
- everything else → the repo-root `.env`.

`provision` creates parent dirs as needed (`ant-academy/code`, `next-web-app/apps/web`) and writes `0o600`.

### The hybrid `introspect` source + the keep-listed gate (`diff`)

The required-key set is **not** a uniform per-repo `.env.example` — verified on stack-dev, **7 of 8 Go repos
ship none** (only `sentinel` does; the count dropped from 8-of-9 when `skillpath` — a Go repo that shipped no `.env.example` — was decommissioned into `app`). So `introspect` (`secretdna.ReadDeclaredKeys` over
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

Eight genes are **waived** — excluded from the coverage denominator with a per-gene rationale in `note`, so
local-vs-prod realities never poison the score:

| Waived class | Genes | Why |
|---|---|---|
| `waived-config` **(M30 field-bake)** | `sentinel/DB_CONNECTION` | docker-compose hardwires it as a sentinel `environment:` entry (`postgresql://postgres@postgresql:5432/postgres?search_path=sentinel&sslmode=disable`), which always overrides `env_file`; sentinel never reads it from `sentinel/.env` at runtime (no `sentinel/.env` exists on stack-dev). An in-network, password-less wiring DSN identical on every stack — config, not a provisioned secret. Was falsely failing the gate at Critical 84.6% before the reclassification |
| `waived-aws-mount` | `platform/LIVEKIT_RECORDING_AWS_ACCESS_KEY_ID` | AWS recording creds are mounted from `~/.aws/credentials`, never a `.env` secret |
| `waived-profile-gated` | `platform/BREVO_KEY` | the class name is historical: the `messenger` profile is gone — `838d907` deleted the container along with the profile. Messenger now runs in-process inside `backend`, gated by `MESSENGER_ENABLED`, which defaults **off** on a developer machine (`ENVIRONMENT=development`; `docker-compose.yml:84-92`), so no default stack ever reads the key |
| `waived-optional` | `platform/BUNNY_STREAM_API_KEY`, `app/TAILSCALE_AUTH_KEY`, `studio-desk/GCLOUD_SERVICE_ACCOUNT`, `studio-desk/YOUTUBE_API_KEY`, `next-web-app/BUNNY_CDN_TOKEN_KEY` | example-only / absent from live / convenience — a local stack comes up without them |

A waived gene names **no operators** and is never measured (`Validate` enforces this). Because the catalog is
scoped to the **default stack's** service set, the denominator is honest for it — `platform/BREVO_KEY` is
waived because the default selection (`core` at platform `0c91421`: `backend` + `gotenberg` + the always-on
`postgresql`/`redis`/`sentinel` floor) sends no mail. Since `838d907`, **no selection does**: there is no
`messenger` container left to start, and `backend`'s in-process messenger stays dormant until
`MESSENGER_ENABLED` is set in `.env`.

> **⚠️ The DNA's `profile` field still literally reads `graphql`** — the token platform `0dab54d` renamed to
> `core`. Nothing mis-selects on it: the field is never resolved against compose, only required non-empty at
> load (`stack-secrets/secretdna/dna.go:233`) and **printed** in the catalog banner
> (`stack-secrets/secretdna/catalog.go:17`), so the staleness is operator-visible rather than behavioural.
> Re-labelling it is a `rosetta-extensions` change, not a doc one.

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

> **⚠️ The demo bring-up ALWAYS passes `--force`, so on that path a re-run is NOT idempotent — it appends a
> full block every time, without bound** (measured M257x iter-269). `up-injected.sh:1538` runs
> `provision … --force` unconditionally, and `--force` skips the copy-if-absent check above while the merge
> stays **append-only** — so each bring-up adds one block. Measured on this box's `stack-demo/platform/.env`
> after 31 bring-ups: **471 lines · 18 distinct keys · 13 of them present 31 times · 0 keys whose value
> varies · `DIRECTUS_TOKEN` blank in all 31.** There is no reaper and no upper bound.
>
> **It is not a bug in either half, and that is why it survived.** Append-only is what makes the tool
> values-blind — `provision/io.go:173-175`: *"an existing line is never re-read for its value or rewritten,
> so provision can never corrupt or echo a value already in the target."* And `--force` is deliberate:
> `up-injected.sh:1522` says it *"overwrites stale keys **AND blanks the `DIRECTUS_TOKEN` family via
> last-wins** (the strip-on-non-prod class)."* **Compose's last-wins resolution is therefore LOAD-BEARING,
> not incidental** — the blank is delivered *by being appended last*.
>
> **Which is why "make the writer replace-or-skip" is the wrong fix, and it was the routed one**
> (`FIX-M257x-262-demo-env-append-is-not-idempotent`). Replace-in-place would either re-read an existing
> value (breaking values-blindness) or drop the trailing blank (re-arming `DIRECTUS_TOKEN` on a demo — the
> fix16/17 class this spec exists to prevent). Any real repair must keep **both** properties and prune
> **older** duplicates rather than stop appending. Re-routed as
> `FIX-M257x-269-force-append-grows-the-demo-env-without-bound`.
>
> **The live hazard to know about:** with N copies of a key, the **last** one wins. Today all 31 agree, so
> nothing is wrong. The moment one writer appends a *differing* value — a stale source, a partial run, a
> hand-edit — the file silently prefers whichever landed last, and the classic symptom is *stack boots,
> catalog empty*. **Diagnose a suspect `.env` by reading the LAST occurrence of a key, never the first.**

### The demo-aware coverage check (`check` / `measure`)

`check` (`secretdna.MeasureForStack`) scores a source against the DNA and exits 1 if **critical coverage <
100%** — `Overall` = Σ(weight·present)/Σ(weight) (criticality-weighted % provisioned), `Critical` =
present-critical/total-critical (unweighted), gate = `Critical == 1.0` — plus a per-repo rollup ("repo X is
short key Y"). It reuses the data-DNA `ratio()` empty-denominator + anti-vacuous-100 guards.

The check is **stack-type-aware** (`--demo`): on a **demo** stack two families of keys are **not** sourced
from the secret dir, yet count as satisfied (`secretdna.demoSatisfied` = `MintedKeys ∪ DemoGeneratedKeys`):

- **The minted Clerk family** — Clerkenstein **mints** them at bring-up (PK_DEMO + an `sk_test_<demo>`
  secret; see [`clerkenstein.md`](../services/clerkenstein.md)) — `secretdna.MintedKeys`: `CLERK_SECRET_KEY`,
  `CLERK_PUBLISHABLE_KEY`, `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY`, `VITE_CLERK_PUBLISHABLE_KEY`,
  `CLERK_WEBHOOK_SECRET`, `CLERK_JWT_KEY`.
- **The demo-auto-generated family (M49 #4)** — `secretdna.DemoGeneratedKeys`: `INVITATION_HMAC_SECRET`. It's
  a **per-deployment** value (not a shared secret), so the source never carries it; a demo is non-prod, so
  `up-injected.sh` generates a **throwaway** value (`openssl rand -hex 32`, values-blind, idempotent) into the
  demo base env at provision. Without the gene + overlay, the `app` exited silently (`NewTokenManager` errors
  when it's unset) — now the pre-flight catches a genuine absence on **dev** while the demo self-provisions.

Otherwise a perfectly-good demo would false-fail on exactly the keys it is designed *not* to carry. This is a
values-blind overlay on `Measure` (presence by gene NAME, never a value); a **dev** stack still requires the
real Clerk keys + the real `INVITATION_HMAC_SECRET` in its source. The pre-flight `check` is wired
**non-fatally** into `/dev-up` + `/demo-up` (warn standard / fail critical — the
[`verification.md`](verification.md) convention).

> **AI-provider keys policy — DECIDED (v1.10b "fit-up" M50): documented-as-absent.** The demo's content
> believability does **not** need live AI — every seeded surface (the heroes, the roster, languages,
> certifications, the Workforce dashboards, the verified-skill chain) renders from **seeded structural data**,
> not a live model call. So the AI-provider keys (`OPENAI` / `ANTHROPIC` / `MISTRAL` / `ELEVENLABS` / the
> `LIVEKIT` voice pair) stay **absent** from the demo secret source — none becomes a throwaway/sandbox demo
> value, and **no real key is ever provisioned** (this decision is itself **values-blind**: it provisions
> nothing). The AI-dependent surfaces are therefore **inert-by-design** unless an operator supplies their own
> sandbox/throwaway keys into the source: the **AI-simulation voice** engine (LiveKit), the **ant-academy
> `/api/ai/chat`** assistant (`OPENAI_API_KEY`/`ANTHROPIC_API_KEY`), and the **M45 AI batch-generation**
> (`ai v1.40.1`) all no-op gracefully — they are not on any demo gate path (the M42 coverage gate is MET on
> both vantages with **zero** AI keys present).
>
> 🔴 **RETRACTED — "these keys remain in the `waived` / optional class … so `check` does not false-fail a demo".
> `waived` is wrong for every one of them, and for two it inverts the gate.** Measured against the committed
> `stack-secrets/secretdna/secret-dna.json` (64 genes, version `fast-build-m256`) @ rext `415240f`: **zero** of
> the keys named above carries a `waived-*` status. The waived class is exactly the **8** genes tabulated under
> "The waived class" above — `sentinel/DB_CONNECTION`, `platform/LIVEKIT_RECORDING_AWS_ACCESS_KEY_ID`,
> `platform/BREVO_KEY`, and 5 `waived-optional` — and **no AI-inference key is among them**. What the manifest
> actually says:
>
> | gene | criticality | status | operators |
> |---|---|---|---|
> | `platform/OPENAI_KEY` | **critical** | **required** | `key-present`, `nonempty` |
> | `platform/AZURE_OPENAI_KEY` | **critical** | **required** | `key-present`, `nonempty` |
> | `platform/ANTHROPIC_API_KEY` · `platform/OPENAI_API_KEY` · `platform/LIVEKIT_API_KEY` · `platform/LIVEKIT_API_SECRET` | standard | **required** | `key-present`, `nonempty` |
> | `app/OPENAI_API_KEY` · `ant-academy/OPENAI_API_KEY` · `ant-academy/ANTHROPIC_API_KEY` · `next-web-app/AZURE_OPENAI_KEY` | standard | **required** | `key-present`, `nonempty` |
> | `platform/ELEVENLABS_API_KEY` · `platform/MISTRAL_API_KEY` · `app/ELEVENLABS_API_KEY` · `platform/SKILLER_OPENAI_KEY` | optional | optional | `key-present` |
>
> **The gate consequence, stated exactly.** `Critical` is the *unweighted* pass ratio over the
> **required + critical** genes and the gate is `Critical == 1.0`; waived genes are excluded from both
> denominators (`stack-secrets/secretdna/measure.go:32-33`, `:44-45`). `platform/OPENAI_KEY` and
> `platform/AZURE_OPENAI_KEY` are **2 of the 13** critical·required genes and both demand `nonempty` — so a
> source carrying neither **cannot** score `Critical == 1.0`. The retracted sentence promised the opposite of
> what the gate does: as written, the waived class it described would have **false-failed its own gate**.
>
> **What survives the retraction:** the *decision* — no real AI key is minted or fabricated for a demo, the
> AI-dependent surfaces are inert-by-design, and the M42 coverage gate is met with zero AI keys. That is a
> statement about the **demo**, not about the **DNA**. If a demo source is meant to pass `check` without AI
> keys, the fix is a **classification change in the manifest** (reclassify or scope those genes, the way
> `sentinel/DB_CONNECTION` was reclassified `waived-config`), not a sentence in this doc asserting a class the
> manifest does not carry.
>
> **The studio-desk AI keys were already caught (v2.7 "july jitter" M252 — the KB-1 correction).**
> `studio-desk/AI_OPENAI_API_KEY` + `studio-desk/AI_ANTHROPIC_API_KEY` are **required · standard** (warn, not
> waived), the **same** posture as the M239 Bedrock class below, because the studio **builder GENERATE**
> (`/api/ai/completion`) is a live-inference surface a demo now drives. See "The studio-desk AI class" below.
> M252 corrected one row; the retraction above is the rest of the same table.

### The Bedrock cred class for app (v2.6 M239, Talk to Data)

**The one AI-provider secret a demo now DOES carry — by operator provision, not by minting.** The M50 policy
above kept AI-provider keys **absent-by-design** because no *believability* surface needs a live model. **Talk
to Data** (`/enterprise/talk-to-data` → `app/internal/askengine`) is the exception the user decided (v2.6,
2026-07-20) to make **FULL**: it is a live-inference feature — natural-language Q&A over the org's data —
that literally cannot answer without a real model call. Its backend (`bedrock.go`) routes through **AWS
Bedrock** (SigV4 over the default AWS credential chain → model `eu.anthropic.claude-sonnet-4-6`, region
`eu-west-1`), so the demo's `app` must hold real AWS creds. This is the **first present-not-absent AWS/cloud
credential class for `app`** — recorded as a secrets-posture shift in [`safety.md` §2.10](safety.md#210-a-demos-app-holds-real-aws-bedrock-creds-v26-m239) below.

**The 5 genes** (all on the `app` repo, target `app/.env`; the `../hyper-studio/.env.example` template):

| Key | Status · criticality | Why |
|---|---|---|
| `AWS_ACCESS_KEY_ID` | **required · standard** | the access-key half the default chain signs SigV4 with |
| `AWS_SECRET_ACCESS_KEY` | **required · standard** | the secret half — a credential **pair** (two distinct values), never an alias family |
| `AWS_REGION` | optional | `bedrock.go` defaults to `eu-west-1` when unset — config with a code default, not a secret |
| `AWS_SESSION_TOKEN` | optional | STS session token — present only with **temporary** creds; absent with permanent IAM keys (the hyper-studio template uses permanent) |
| `CLAUDE_CODE_USE_BEDROCK` | optional | a Claude Code CLI convention — **NOT read** by `askengine/bedrock.go` (which always routes Bedrock); provisioned for parity, inert for the app |

**Why required-`standard`, deliberately NOT `critical` (the R3 decision).** The two real creds are `required`
(the `check` counts them + flags them missing) but **`standard`, so their absence never fails the `Critical
== 100%` gate**. A box without Bedrock creds still brings a demo up cleanly — Talk to Data just stays
**inert** (the SSE stream opens, but the agentic loop fails with `no EC2 IMDS role found`), the same graceful
degradation as any absent AI-provider key. Making them `critical` would break every creds-less demo's gate,
which is exactly the R3 risk the release design flagged. They are also **NOT** in the `demoSatisfied` set —
unlike the minted Clerk family or the auto-generated `INVITATION_HMAC_SECRET`, they are **operator-provided
from the source**, so a creds-less demo legitimately *warns* (standard, non-fatal) rather than being treated
as satisfied.

**The bridge — provision writes `app/.env`, but the containerised backend reads `platform/.env`.** The DNA
targets `app/.env` (where the operator drops the creds, mirroring the native-dev backend env). But the demo's
**backend (`app`) container** reads its env from `env_file: .env` = the demo's **`platform/.env`** — *not*
`app/.env`, which is the repo-local native-dev env the container never mounts. And the M217 override **drops
the `~/.aws` mount** for a demo (the empty-dir `EISDIR` bug), so **env vars — not a mount — are the only
vehicle to the container.** So `up-injected.sh`'s `bridge_bedrock_creds()` copies exactly the Bedrock class
`app/.env → platform/.env` right after provision: **values-blind** (bytes move file→file via `>>`, never
surfaced), **idempotent** (copy-if-absent — a re-run adds nothing), and **non-fatal** (a creds-less
`app/.env` just logs an inert-note). Proven live 2026-07-21: the provisioned creds get a real Bedrock answer
from `eu.anthropic.claude-sonnet-4-6` (`converse` → `pong`, `end_turn`, eu-west-1).

### The studio-desk AI class (v2.7 "july jitter" M252, the studio builders)

**A demo's third live-AI-in-demo surface — after Talk to Data.** M239 (above) made **Talk to Data** a demo's
first present-not-absent live-inference surface. **M252** does the same for **studio-desk's builder GENERATE**:
the advanced + guided simulation builders `POST /api/ai/completion` through the studio Copilot's multi-provider
chain (Azure OpenAI / OpenAI / Anthropic), which cannot answer without a real model call — so a demo's
studio-desk must hold a real provider key. Like the M239 Bedrock pair, it is an **operator-provisioned** AI
credential class (never minted, never demo-generated). The GENERATE surface itself is reached by the
**Clerkenstein-authenticated org-admin hero** (the manager, who passes studio-desk's `checkEnterpriseAndAdmin`
gate) — M252 routes the AI keys into that already-reachable container via `env_file`; it does **not** disarm auth
(there is **no** `MOCK_CLERK`, and the demo studio is **not** a "server is open" surface).

**These keys were ALWAYS required · standard — not the waived class above (the KB-1 correction).** The studio AI
keys have been **DNA genes** since the coverage DNA existed — `studio-desk/AI_OPENAI_API_KEY` +
`studio-desk/AI_ANTHROPIC_API_KEY`, **required · standard**: the `check` counts them and **warns** when a source
omits them, but their absence never fails the `Critical == 100%` gate — the **same R3 posture as the M239 Bedrock
pair**, not the M50 waived/optional class. A key-less demo legitimately *warns* (and the studio GENERATE 500s /
stays inert — the same graceful degradation as any absent AI-provider key), rather than being treated as
satisfied. They are **not** in `demoSatisfied` — operator-provided from the source, not minted like the Clerk
family nor auto-generated like `INVITATION_HMAC_SECRET`.

**Source coverage vs container coverage — the DNA proves one, autoverify the other.** The `stack-secrets` DNA is
**source-vs-DNA only**: it scores whether the *source* provisions the gene into `studio-desk/.env`; it never
inspects a running container. That is necessary but not *sufficient* here, because studio-desk is a
**base-compose** service — in a demo it inherits only `platform/.env` (which carries no AI keys), so the value
must actually be routed **into the container**. M252 does that with the injected-override
**`env_file: <clone>/studio-desk/.env`** (see [`demo/frontend-tier.md`](demo/frontend-tier.md) +
[`../services/studio-desk.md`](../services/studio-desk.md) § Demo AI wiring), and adds the missing
**container-side** proof: a **demo-aware, non-fatal, values-blind** assertion in the **live-verify layer**
(`stack-verify/live/autoverify.sh`) that the studio-desk **container** actually carries a provider key —
mirroring autoverify's existing directus `DB_CONNECTION_STRING` container check. **The DNA proves source
coverage; autoverify proves the container carries it.**

### The values-blind safety statement (the inviolable invariant)

**No verb ever reads, echoes, logs, or persists a secret VALUE** — not in stdout, stderr, an error, or any
committed file. Operators see key NAMES + presence only, at most a value's *shape* (a `url`/`jwt`/`pk_`/`sk_`
structural prefix via `ClassifyShape`, the single function permitted to look at a value, which returns a
shape token, never the value). Extraction from a source is name-only (cut on the first `=`); the value half
is discarded the moment a line is parsed. `provision` **moves** secret bytes source→gitignored-target (its
job) but the bytes never leave the value-carrying boundary (`provision/io.go`) except into the target `.env`.
The `secret-dna.json` file stores NAMES only and is **committable** (unlike a `.env`). This mirrors the
platform's values-blind `Guard.PreflightEnv` discipline — the safety clause is stated authoritatively in
[`safety.md`](safety.md#29-secret-provisioning-is-values-blind-and-never-re-arms-the-prod-write-path-v16-m27m30).

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

M27 delivers the framework: the source-dir/zip ingestion + the secret-coverage DNA (the 6-repo/**64**-gene map
at today's `fast-build-m256`; 55 genes when M27 shipped it — the growth is the split history above)
+ the two-tier keep-listed `diff` gate, **113 Go tests** (hermetic, `-race` clean). M28 adds the `provision`
engine (alias-mapped per-file writes, copy-if-absent + `--force`, N=0-guarded, the `DIRECTUS_TOKEN`
non-rearm regression pinned) + the demo-aware `check`, wired non-fatally into `/dev-up` + `/demo-up`
pre-flight (**160 Go tests**). M29 authors this spec + the `/stack-secrets` skill + the corpus wiring.
**M30 — the field-bake — closes the version:** it assembled a compliant `.agentspace/secrets` dir from
stack-dev (values-blind, the knowns waived), proved `check` scores **Critical == 100%** on both dev and demo
(exit 0), then brought a fresh **demo-3 LIVE from that assembled source** (provision → 17 containers UP,
all liveness+readiness probes pass) — the observable-behavior gate, met. The bake caught + fixed **2 real
release bugs** Fate-1: (1) `sentinel/DB_CONNECTION` was critical/required but is compose-injected config →
reclassified `waived-config` (above) + regression test; (2) the demo bring-up only *checked* coverage but
never *provisioned* (and `preflight.sh` resolved its source path one level too shallow, silently skipping
the demo gate) → added the provision step + fixed the path. Tag `stage-door-m30`. The tooling is
**values-blind**, **never commits `.env`**, **never writes prod**, and **zero platform-repo edits**.
