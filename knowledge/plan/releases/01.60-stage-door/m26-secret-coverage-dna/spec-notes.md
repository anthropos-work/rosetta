# M26 — Spec notes

_Technical detail accumulated during build. Stub at scaffold; sections seeded from the overview scope._

## The `stack-secrets` section layout
_(go.mod module path, cmd/stacksecrets entry, package boundaries — mirror stack-snapshot/stack-seeding.)_

## Source-ingestion contract (dir + zip)
_(default `.agentspace/secrets`; the explicit layout contract; values-blind extraction; the zEnvs/per-repo-.env
exclusion rule; zip mode.)_

## The secret-DNA schema
_(gene = repo×KEY; gene id `<repo>/<KEY>`; the per-gene fields; the `secret-dna.json` shape; the id regex; reuse of
Criticality 3/2/1 + the two-metric Score + `ratio()` + the 0/1/3 exit-code contract.)_

## Hybrid `introspect` source
_(platform/.env_example baseline + frontend/sentinel .env.example + docker-compose-injected keys; the rule for the
8 Go repos with no .env.example; the curated platform-env list.)_

## `list` + `diff` "keep-listed" gate
_(diff exits 1 on required-key drift + on runtime-required-but-undeclared keys — the anti-vacuous-green guard.)_

## Waived classes + alias families
_(AWS-via-mount, profile-gated, optional Bunny/GCloud → waived; GH_PAT alias family vs the OPENAI_KEY/OPENAI_API_KEY
distinct-similar pair.)_
