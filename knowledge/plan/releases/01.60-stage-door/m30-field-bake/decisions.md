# M30 — Decisions

_Implementation decisions with rationale, numbered `M30-D1`, `M30-D2`, … . The field-bake is a prove+fix
milestone; these record the two Fate-1 field fixes + the live-bring-up wiring design._

## M30-D1 — `sentinel/DB_CONNECTION` reclassified `critical/required` → `waived-config`

**Decision:** the gene is `waived-config` (criticality `optional`, status `waived-config`, no operators, scope
`config`), not a sourced critical secret.

**Why:** the platform `docker-compose.yml` injects sentinel's DB connection as a hardcoded `environment:`
entry (`postgresql://postgres@postgresql:5432/postgres?search_path=sentinel&sslmode=disable`), and compose
`environment:` always overrides `env_file:`. Sentinel never reads `DB_CONNECTION` from a `.env` at runtime;
the `sentinel/.env.example` is a 26-byte doc stub and no `sentinel/.env` exists on stack-dev. It is a
password-less, in-network wiring DSN identical on every stack — config, not a provisioned secret. Marking it
critical/required made the gate demand a secret the runtime ignores, falsely failing coverage at Critical
84.6%. **Alternatives rejected:** (a) leave it critical and require operators to hand-create `sentinel/.env` —
would demand a secret no runtime reads; (b) drop the gene entirely — loses the documented record of *why* it's
not provisioned. The new `waived-config` class is the honest middle: keep-listed, never measured.

**Guard preserved:** the anti-vacuous-100 guard still holds (12 required+critical genes remain). Pinned by a
regression assertion in `secretdna/secret_dna_json_test.go` (waived-with-no-operators). DNA version bumped
`stage-door-m27` → `stage-door-m30`.

## M30-D2 — the demo bring-up PROVISIONS, then moves only the env-file (topology stays on stack-dev)

**Decision:** after the secret pre-flight passes, `up-injected.sh` runs `stacksecrets provision --force` to
write stack-demo's per-repo `.env` from `.agentspace/secrets`, then repoints the run's `--env-file` base at
the provisioned `stack-demo/platform/.env`. The compose **topology** (`-f docker-compose.yml`) is left on
stack-dev; only the secrets-bearing env-file moves.

**Why:** before M30 the demo bring-up only *checked* coverage, then ran from the operator's live
`stack-dev/platform/.env` — so the assembled source was scored but never used (the field-bake's whole point is
to run *from* the assembled dir). Moving only the env-file (not the compose dir) keeps the proven topology
resolution on the canonical stack-dev clone set (zero drift risk) while sourcing secrets from the curated dir.
**Alternatives rejected:** (a) point `-f` at a stack-demo compose too — drifts topology onto a peer clone for
no benefit; (b) provision unconditionally and hard-fail on any error — would block a good bring-up. **Made
non-fatal** (`verification.md` convention): a missing source / build failure / non-zero exit degrades to the
legacy base with a loud note; `DEMO_NO_PROVISION=1` opts out. The repoint happens only if provision succeeded
AND the provisioned env carries a non-empty `GH_PAT` (values-blind presence check), else it stays on the
proven base.

**Safety:** provision writes the `DIRECTUS_TOKEN` family BLANK on the non-prod target AND the injection
override strips it at compose-emit (defense-in-depth — the fix16/17/M28 non-rearm class). Verified BLANK in
every live demo-3 container.

## M30-D3 — `preflight.sh` source-path resolution corrected to two-levels-up

**Decision:** `REPO_ROOT` resolves as `EXT_ROOT/../..` (was `EXT_ROOT/..`), and `up-injected.sh` passes
`--from "$SECRET_SRC"` explicitly.

**Why:** `rosetta-extensions` lives two dirs deep under the rosetta root in both layouts
(`<root>/.agentspace/rosetta-extensions` authoring; `<root>/stack-<role>/rosetta-extensions` per-stack). The
one-level-up resolution yielded `.agentspace` (or `stack-demo`), so the default `--from` doubled to a
nonexistent `.agentspace/.agentspace/secrets` and the demo-aware coverage gate **always silently skipped (exit
2)** instead of running. Two-levels-up reaches the real root for both layouts (verified). Passing `--from`
explicitly from the caller is belt-and-suspenders so the pre-flight scores the exact source the provision step
writes from, independent of the default-resolution. **Alternative rejected:** rely on the explicit `--from`
alone and leave the default broken — leaves a latent footgun for any other caller relying on the default.

## Adversarial review (Phase 2c)

Scenarios considered for the M30-touched ext code (the live-bring-up wiring + the DNA reclassification):

1. **Provision succeeds but the provisioned env has an empty `GH_PAT`** (e.g. the source omits it). → The
   repoint guard checks for a non-empty `GH_PAT` (values-blind `grep | tail -1 | cut`) and stays on the legacy
   base if absent; the subsequent loud GH_PAT empty-check then fails clearly rather than building with no PAT.
   Handled.
2. **The secret source dir is absent on the box** (no `.agentspace/secrets`). → The provision block is gated
   on `[ -d "$SECRET_SRC" ]`; absent → a loud non-fatal note + legacy base. The bring-up still works (it just
   runs from the operator's dev env, the pre-M30 behavior). Handled.
3. **`go` toolchain or the secret-DNA missing.** → Guarded (`command -v go` + `[ -f "$SECS_DNA" ]`); a build
   failure or missing tool degrades non-fatally to the legacy base. Handled.
4. **A regression silently re-promotes `sentinel/DB_CONNECTION` to critical.** → Pinned by the
   `secret_dna_json_test.go` assertion (waived + no operators); a re-promotion fails the test. Handled.
5. **Provision leaks a secret value into a log / the bring-up output.** → Provision stdout is key NAMES +
   write/blank/skip counts only; `provision_safety_test.go` (M28) reflection-walks the Report/plan/errors for
   any value. Verified live: no len-32 prod token in any container or commit. Handled.
