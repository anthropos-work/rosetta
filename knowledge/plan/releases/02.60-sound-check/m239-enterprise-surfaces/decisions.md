# M239 — Decisions

_(implementation choices with rationale accumulate here during build)_

- **User decision (2026-07-20): talk-to-data → FULL.** Real AWS Bedrock creds via `/stack-secrets` + secret-coverage DNA extension for `app` (reference `../hyper-studio/.env.example`), not just a flag. Recorded at design time; carried here for build traceability.

## §1 — talk-to-data build decisions

- **D1 — flag via env var, not a demopatch.** `NEXT_PUBLIC_DEMO_FLAGS_ALL=true` is baked into the web + hiring `.env.local` overlays (up-injected.sh). The frontend ALREADY reads it (`useTalkToDataAccess.ts:39` / `useCoursebuilderAccess.ts:40` — a demo escape hatch that forces the flag on while still requiring `isAdmin`), so this is demo-env wiring of an existing env var — cleaner than a demopatch (the M219 aireadiness patch was needed only because that gate reads no env var). Only those two admin-gated surfaces read it (verified by grep), so forcing it on unlocks nothing unintended. Folded into `next_web_patchset_fp` so a pre-M239 image rebuilds once (build-inlined value).

- **D2 — creds reach the container via `platform/.env`, bridged from `app/.env` (the key reconciliation).** The user's design said store in `.agentspace/secrets/app/`, DNA on `app`, "wire into the app compose service". Verified against actual code: the demo's **backend (`app`) container reads `env_file: .env` = the demo's `platform/.env`**, NOT `app/.env` (repo-local native-dev env, never mounted by the container). And the M217 override **drops the `~/.aws` mount** for a demo. So honoring the user's `app`-DNA framing correctly required a bridge: `bridge_bedrock_creds()` copies the Bedrock class `app/.env → platform/.env` right after provision (values-blind file→file, idempotent, non-fatal). This fully honors "wire into the app compose service" (the app/backend container now has the AWS env) while being correct for a containerized demo.

- **D3 — required-`standard`, deliberately NOT `critical` (R3).** The 2 real creds (`AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`) are `required` (counted + flagged-if-missing) but `standard`, so their absence never fails the `Critical==100%` gate — a creds-less box still brings a demo up (Talk to Data merely inert, "no IMDS role"). Making them critical would break every creds-less demo (exactly R3). Region/session-token/`CLAUDE_CODE_USE_BEDROCK` are optional (config/STS-only/inert-for-app). NOT in `demoSatisfied` — operator-provided, not minted.

- **D4 — `CLAUDE_CODE_USE_BEDROCK` is inert for the app.** The audit + `bedrock.go` confirm askengine NEVER reads it (it always routes Bedrock unconditionally). It is a Claude Code CLI convention; provisioned for parity with the hyper-studio template, but does nothing for the app. Documented so a future reader doesn't chase it as load-bearing.

- **D5 — live verify: gate (b) Bedrock round-trip PROVEN.** The wired creds get a real answer from the exact app model/region — `aws bedrock-runtime converse --model-id eu.anthropic.claude-sonnet-4-6 --region eu-west-1` → `pong`, `stopReason end_turn`, 20 tokens (2026-07-21). This is the load-bearing, genuinely-uncertain part (the brief's "creds/region/permission" escape) — definitively answered YES. NOT a blocker. Full end-to-end UI click-path verification via a local demo bring-up follows.

## §2 — #4 library empty-first-load — VERDICT: no remaining defect (not a fix to force)

- **D6 — the `:5050` carry is already resolved in the current tooling.** apps/web's `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` is baked to the OFFSET origin by `up-injected.sh` (`build_frontend_next_web` line 744: `--build-arg NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=$SCHEME://$HOST:$((5050+OFFSET))/graphql`), which OVERRIDES the Dockerfile.dev ARG default `http://localhost:5050/graphql` (line 23) that `ENV` bakes. `build_frontend_hiring` (line 1023) does the same. The remaining `localhost:5050` references are the Dockerfile ARG *default* (overridden) + `packages/graphql/codegen.ts` (dev-authoring schema introspection, never in the runtime client). The M218/M220 image-reuse check rebuilds a stale-endpoint image. M237's re-triage already found the grid populates (7→29 cards, 0 errors). **No client-fetch race defect remains** — per the three-fate guard, this is recorded as a verified verdict, not a manufactured fix. Confirmed live on the demo bring-up.

## §3 — #1 hierarchical manager menu — presence-verify + coverage assertion

- **D7 — M237-resolved; confirmed by code + live.** M237 already RESOLVED #1 (the manager nav renders the grouped Organization structure on a fresh build). §3 is a presence-verify (live) + a coverage-sweep assertion. The coverage manifest notes manager descriptors need live calibration; the assertion is authored against the live manager render from this milestone's bring-up (avoids an uncalibrated false-RED).
