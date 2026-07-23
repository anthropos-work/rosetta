---
iter: 18
milestone: M244
iteration_type: tik
status: closed-fixed
run: 7
created: 2026-07-23
---

# iter-18 — gate (c) discrete-spec remote-capability retrofit (run 7, tik 1)

**Active strategy:** TOK-02 (consolidate + sharpen TOK-01 — sweep the cheap remaining gates on the green m244 seed).

## Step 0 — re-survey (mandatory)
Metric at open: **5/8** (a,b,d,e,g green live). Gate (c) coverage half GREEN (iter-16) + persona-check + m224 green (iter-17). Remaining gate (c) work: the localhost-only discrete stack-verify specs + the 16 Playthroughs (LAST, pt-world reset). TOK-02/iter-17 routed the discrete retrofit as run-7 tik-1.

**Considered substituting a metric-MOVING target** (gate f or h) so iter-18 avoids being a 3rd no-metric-delta tik (16/17/18 → a forced tok at iter-19). Re-survey of the alternatives (this iter's recon):
- **gate (f)** is NOT a clean single-tik tick: `BURNIN-M221-dev-public-host` requires a real `/dev-up --public-host` remote **dev** burn-in (a separate heavy bring-up), not a demo probe.
- **gate (h)** has a **likely-real blocker**: billion's demo has **NO Bedrock/AWS creds** (values-blind check: platform/.env + app/.env carry Azure/OpenAI/Mistral only; source `.agentspace/secrets/app/.env` has no AWS/BEDROCK), so M239's talk-to-data "live via real Bedrock" may not answer live; and the autoverify is **STALE** (19:31:10Z, >4h, run-latency threshold 14400s) so p95 needs an autoverify refresh first.

⇒ **No clean single-tik metric-mover exists** — the coarse binary-per-gate metric stays flat regardless (TOK-02 already characterized this). So the highest-value reliable iter-18 is the **retrofit** (orchestrator's #1): it is foundational, unblocks gate-c discrete AND the gate-h talk-to-data/cockpit spec proofs, and its own spec-run (talk-to-data-m239) **definitively answers the Bedrock question**. Target held = the retrofit.

## Cluster / target identified
Retrofit the tailnet host+scheme drive-path (`COVERAGE_HOST`/`run-coverage.sh` pattern: `https://${HOST}:port` bases + `ignoreHTTPSErrors`) into the remaining localhost-only discrete stack-verify **gate** specs so they drive billion, then run them green on the demo seed (no reset). Lead with **talk-to-data-m239** (also the gate-h Bedrock probe).

## Hypothesis
Most discrete specs already read full-URL env overrides (talk-to-data → `ROSETTA_E2E_BASE_URL`/`ROSETTA_E2E_FAPI_URL`; calibrate → `CALIBRATE_APP`/`CALIBRATE_FAPI`); the retrofit is a per-spec runner/env that composes the tailnet `https://billion…:port` bases (+ a host+scheme var for the truly-hardcoded ones). With `ignoreHTTPSErrors` the cockpit-login lib already logs in over the tailnet https-fapi (coverage/m224/persona prove it). Expect the discrete gate specs to drive billion green.

## Expected lift
Gate-c stack-verify discrete half GREEN on billion. **No binary tick** (gate c still needs the 16 Playthroughs LAST) → metric stays 5/8. This is the coarse-metric artifact TOK-02 flagged; the deliverable is the remote-capability retrofit + live-green discrete specs.

## Phase plan
Retrofit the discrete gate specs (rext, harness-only — no billion re-pin/re-bake) → run each green on billion → re-tag + push the consumption tag → close.

## Escalation conditions
- A spec surfaces a real defect needing a **platform edit** → user-blocker (STOP).
- talk-to-data-m239 fails on Bedrock → route as a gate-(h) blocker finding (a creds/provisioning decision, NOT a platform edit); does not block the retrofit deliverable.

## Acceptable close-no-lift outcomes
Retrofit lands + the discrete gate specs run green modulo a documented deterministic blocker (e.g. talk-to-data Bedrock) → closed-fixed-partial with the blocker routed. A complete retrofit with all discrete gate specs green → closed-fixed.
