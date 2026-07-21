# M239 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **talk-to-data (a) flag enablement** — `NEXT_PUBLIC_DEMO_FLAGS_ALL=true` baked into web+hiring `.env.local` (env var the platform already reads; not a patch) + folded into `next_web_patchset_fp`. rext `443a365` / tag `sound-check-m239-enterprise-surfaces`. Tested (`test_bedrock_bridge_m239.py`).
- [x] **talk-to-data (b) real AWS Bedrock creds** — Bedrock cred class in the secret-DNA (2 req·standard + 3 optional; NOT critical per R3) + values-blind `bridge_bedrock_creds()` (app/.env→platform/.env) + provisioned via the assembled source. Tested (`secret_dna_json_test.go` + bridge tests). **Live Bedrock round-trip PROVEN** (`converse` → `pong`, `eu.anthropic.claude-sonnet-4-6`, eu-west-1). Full-UI live confirm on the demo bring-up.
- [x] **#4 library empty-first-load** — VERDICT: no remaining defect. The `:5050` carry is already resolved (`up-injected.sh:744/1023` bake the offset endpoint for both web+hiring, overriding the Dockerfile ARG default). No client-fetch race remains (M237: grid populates 7→29). Recorded verdict, not a manufactured fix. Live-confirm on the bring-up.
- [x] **#1 hierarchical manager menu** — M237-resolved (grouped Organization nav for managers). Manager hero = org admin (`users.go:141`). Live presence-verify + coverage-sweep assertion (calibrated against the live manager render).
- [x] **Delivers** — `corpus/ops/secrets-spec.md` (the Bedrock cred class subsection + updated 61-gene map) + `corpus/ops/safety.md` §2.10 (secrets-posture shift: demo `app` holds real Bedrock creds; blast-radius + operator-scope caveat).

_Note: all sections implemented + unit-tested + committed (rext) / drafted (docs). The live demo-1 bring-up (in flight) confirms the running-demo wiring (flag baked, bridge landed, backend has creds) + the manager-menu render + calibrates the §3 coverage assertion._
