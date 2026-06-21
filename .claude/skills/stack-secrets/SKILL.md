---
name: stack-secrets
description: Provision a stack's secrets (dev-N or demo-N) — write every repo's target .env from one secret source, and verify coverage — VALUES-BLIND (no verb ever reads, echoes, or logs a secret value). Drives the stacksecrets CLI (check / provision / status). Use after cloning a stack's repos, when asked to provision / set up / fill in / verify the .env secrets for a stack.
argument-hint: [dev-N|demo-N] [--from DIR|ZIP] [--check|--provision|--status]
---

# Stack Secrets — provision + verify a stack's `.env` secrets, values-blind

Provisions a target stack (`dev-N` or `demo-N`) by writing **every repo's target `.env`** from one secret
source — six repos, alias-mapped per file — and **verifies coverage** against the secret-DNA. It replaces the
manual `platform/.env` hand-copy the old `setup_guide.md` documented. The headline property is **safety**:
**no verb ever reads, echoes, logs, or persists a secret VALUE** — you see key NAMES + presence only.
`provision` moves secret bytes source→(gitignored)target only; it never re-arms the prod-write path. Source
of truth: [`corpus/ops/secrets-spec.md`](../../../corpus/ops/secrets-spec.md).

> **Two `stack-secrets` namespaces, kept distinct.** This **skill** (`/stack-secrets`) drives the
> `stacksecrets` CLI. The extensions **section** named `stack-secrets` (`rosetta-extensions/stack-secrets/`)
> is where that CLI is built. The skill operates the tooling; the section name inside the repo is unchanged.

## Where this sits in the flow
Clone a stack's repos (`make init` / `/dev-up` / `/demo-up`'s `ensure-clones`) → **`/stack-secrets N
--provision`** → `/dev-up N` or `/demo-up N` (build + start). The pre-flight `check` already rides inside
`/dev-up` + `/demo-up` (non-fatal: warn standard / fail critical) — use this skill to **provision** a stack's
`.env` files up front, to **check** coverage on demand, or to inspect the catalog.

> **M26: `/demo-up`'s `ensure-clones` is real.** A demo bring-up now runs `ensure-clones.sh` first — it
> bootstrap-clones `stack-demo/platform` from GitHub + `make init`s the peer clone set, then **seeds**
> `stack-demo/platform/.env` from `stack-dev/platform/.env` copy-if-present (non-fatal if stack-dev is absent).
> `/stack-secrets --provision` then OVERLAYS the per-repo `.env`s values-blind from `.agentspace/secrets` — so
> the seed is a base the provisioner layers over, and a box with no `stack-dev` (only `.agentspace/secrets`) is
> fully supported.

> **M30: `/demo-up` now auto-provisions.** A demo bring-up runs the pre-flight `check`, then PROVISIONS the
> demo's per-repo `.env` from `.agentspace/secrets` (values-blind, `--force`) and runs the demo from that
> assembled-source base env — so a fresh `/demo-up N` is self-sourced from the secret dir without a separate
> `/stack-secrets --provision` step. `DEMO_NO_PROVISION=1` opts out (runs from the ensure-clones-seeded
> `stack-demo/platform/.env` base instead — M26). Use this skill standalone to provision a **dev** stack, to
> provision up front, or to `check`/inspect.

## Mission

1. **Read the spec** — [`corpus/ops/secrets-spec.md`](../../../corpus/ops/secrets-spec.md) (the source-dir/zip
   layout, the 6-repo/55-gene secret-DNA, the per-repo target-file map, the values-blind safety statement, the
   alias/collision rules, the waived class, the `DIRECTUS_TOKEN` non-rearm safety). **Confirm the target is a
   non-prod stack** (`dev-N` / `demo-N`, never production). `provision` hard-refuses the **main dev stack
   (N=0, `anthropos`)** without `--force` — it holds the operator's real source `.env`.

2. **Confirm the secret source** — a **directory** (default `.agentspace/secrets`) or a `.zip`, laid out
   **by repo** at the exact target paths (this is the `zEnvs`/stray-`.env` defence; the reader opens only
   declared `<repo>/<file>` paths, never enumerates the tree):
   ```
   .agentspace/secrets/
     platform/.env
     app/.env
     sentinel/.env
     studio-desk/.env
     next-web-app/apps/web/.env
     ant-academy/code/.env.local
   ```
   The source dir is **gitignored** — never commit a `.env`.

3. **Build the tool** from a **pinned-tag** `rosetta-extensions` clone (the standard two-clone policy — the
   canonical source is the `.agentspace/rosetta-extensions/` authoring copy; a stack consumes it at a pinned
   tag). Use the matching per-stack clone for the target (`stack-dev/` for a dev-N, `stack-demo/` for a
   demo-N), checked out at the latest tag **`stage-door-m30`** (which carries `provision`, the demo-aware
   `check`, the M30 DNA fix, and the field-bake demo bring-up wiring):
   ```bash
   SECDIR=stack-demo/rosetta-extensions   # or stack-dev/... for a dev-N
   git -C "$SECDIR" fetch --tags --quiet && git -C "$SECDIR" checkout --quiet stage-door-m30
   SECS="$SECDIR/stack-secrets"
   DNA="$SECS/secretdna/secret-dna.json"
   go build -o /tmp/stacksecrets "$SECS/cmd/stacksecrets"
   ```

4. **Run the requested verb:**

   **`--provision`** (the headline verb) — write each repo's target `.env` from the source, values-blind.
   **Copy-if-absent by default** (idempotent: a re-run skips already-present keys). Always preview with
   `--dry-run` first:
   ```bash
   SRC=.agentspace/secrets                  # the secret source dir (default)
   ROOT=stack-demo                          # the stack workspace root holding the cloned repos
   # preview only — the per-file plan (write / blank / skip / missing key NAMES), nothing written:
   /tmp/stacksecrets provision --dna "$DNA" --from "$SRC" --stack-root "$ROOT" --stack demo-N --dry-run
   # provision for real (writes <ROOT>/<repo>/<target_file>, 0o600, append-only):
   /tmp/stacksecrets provision --dna "$DNA" --from "$SRC" --stack-root "$ROOT" --stack demo-N
   # overwrite existing keys AND permit the main dev stack (N=0) — deliberate:
   /tmp/stacksecrets provision --dna "$DNA" --from "$SRC" --stack-root "$ROOT" --stack dev-N --force
   ```
   On a **non-prod** stack (the default — pass `--prod` only for a real prod target, reachable solely via the
   N=0 `--force` path) `provision` writes the `DIRECTUS_TOKEN` family **blank** (`KEY=`), deferring to the
   injection override — it **never** re-arms the stripped prod-write path (the fix16/17 safety class).

   **`--check`** (= `check` / `measure`) — score the source against the DNA; exit 1 if **critical coverage <
   100%**. On a **demo** stack add `--demo` (Clerkenstein-minted Clerk keys count as satisfied without the
   source):
   ```bash
   /tmp/stacksecrets check --dna "$DNA" --from "$SRC"            # a dev stack — real Clerk keys required
   /tmp/stacksecrets check --dna "$DNA" --from "$SRC" --demo     # a demo stack — minted Clerk keys count
   ```
   It prints per-gene PASS/FAIL, the two metrics (Overall weighted + Critical unweighted), and a per-repo
   rollup ("repo X is short key Y").

   **`--status`** (= `list`) — print the per-repo secret catalog (required / optional / waived + alias
   families). No source needed:
   ```bash
   /tmp/stacksecrets list --dna "$DNA"
   ```

   **(maintenance) `diff`** — the keep-listed gate: reconcile the DNA against the hybrid declared set; exit 1
   on an unlisted-required key (the anti-vacuous-green guard). Run it when the platform's required keys may
   have drifted:
   ```bash
   /tmp/stacksecrets diff --dna "$DNA" --stack-root "$ROOT"
   ```

5. **Verify** — `provision` prints a per-file summary (`N written, N blanked, N skipped`), key NAMES only.
   Follow with `check` to confirm `Critical: 100.0%`. **No secret value ever appears in any output** — if you
   ever see one, that is a bug; stop and report it.

## Safety (the load-bearing part)
- **Values-blind, always.** No verb reads, echoes, logs, or persists a secret VALUE — stdout, stderr, errors,
  and any committed file carry key NAMES + presence/shape only. `provision` MOVES secret bytes
  source→gitignored-target (its job) but the bytes never leave the value-carrying boundary except into the
  target `.env`. **The skill itself NEVER prints a secret value** — do not cat/echo a `.env` or paste a value
  into chat.
- **Never commit a `.env`.** The source dir + every written target are gitignored. The `secret-dna.json` is
  NAMES-only and committable.
- **Never write prod.** `provision` refuses the main dev stack (N=0) without `--force`; the default target is
  non-prod. On non-prod the `DIRECTUS_TOKEN` family is written **blank** — `provision` defers to the injection
  override and never re-arms the prod-write path (the fix16/17 / `DIRECTUS_TOKEN`-non-rearm class — see
  [`corpus/ops/safety.md`](../../../corpus/ops/safety.md#29-secret-provisioning-is-values-blind-and-never-re-arms-the-prod-write-path-v16-m27m28)).
- **Idempotent.** Copy-if-absent by default — re-running is safe (skips present keys, re-blanks the strip
  keys). `--force` is the deliberate overwrite.

## Defaults
- Source: `--from` → default `.agentspace/secrets` (a dir; a `.zip` is auto-detected by extension).
- Tool tag: `stage-door-m30` (the latest; carries `provision` + demo-aware `check` + the M30 field-bake demo
  bring-up wiring). `/demo-up` now AUTO-PROVISIONS a demo's per-repo `.env` from `.agentspace/secrets` by
  default (M30 — `DEMO_NO_PROVISION=1` opts out), then runs the demo from that assembled-source base env.
- DNA: `<clone>/stack-secrets/secretdna/secret-dna.json` (committed, NAMES-only).
- Exit codes (the `0/1/3` contract): `0` ok / covered / no drift / wrote · `1` critical key missing, gate
  tripped, or write/guard error · `3` usage error.

## Related skills

| Skill | Use when |
|-------|----------|
| `/dev-up` · `/demo-up` | Bring up the stack (they run the pre-flight `check` non-fatally); provision first |
| `/stack-list` | List live stacks to pick a target |
| `/stack-snapshot` · `/stack-seed` | Set-dress + seed the stack after it's up and provisioned |
