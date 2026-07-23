# M239 — Spec notes

Topic → doc → code triples + Bedrock-secret / library-race findings accumulate here during build.

## talk-to-data — (a) flag + (b) real Bedrock creds
- (a) flag enablement: `NEXT_PUBLIC_DEMO_FLAGS_ALL` or a flag-gate demopatch (M219/M232 pattern).
- (b) real AWS Bedrock creds via `/stack-secrets`: key set `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` / `AWS_REGION` + `CLAUDE_CODE_USE_BEDROCK` (reference `../hyper-studio/.env.example`).
- Secret-coverage DNA extension for the `app` service; mount/env-wire into the `app` compose service.

## Library empty-first-load (#4)
- Client-fetch race + the open non-offset `:5050` `apps/web` endpoint carry.

## Hierarchical manager menu (#1)
- Confirm it renders for managers via the presence sweep (post-M237 fresh build).

## Secrets posture (safety.md)
- New present-not-absent secret class for `app` (same class as the AI-provider keys) — record the note.

## Pre-flight audits — all sections (2026-07-21)
Verdict **GREEN**. Report: `kb-fidelity-audit.md`. Sha at audit: `06250dd` (release base). Covers all M239 sections (single subsystem: demo-stack tooling + secrets DNA + docs) — reuse across sections per build-milestone §"Audit reuse".

### Verified topic → doc → code triples
- **talk-to-data flag gate** → `ai_architecture.md:51` / `backend.md:19,134` → `next-web-app .../hooks/useTalkToDataAccess.ts` (`FEATURE_FLAG='flag_enable_talk_to_data'`; `isAdmin && effective`; demo bypass `NEXT_PUBLIC_DEMO_FLAGS_ALL==='true'` ⇒ `{isEnabled: isAdmin}`). ALIGNED.
- **Bedrock engine** → `staging-bringup.md:188-198` (existing anchor) → `app/internal/askengine/bedrock.go` (`DefaultModelID="eu.anthropic.claude-sonnet-4-6"`, `DefaultRegion="eu-west-1"`, `config.LoadDefaultConfig` default chain, SigV4). ALIGNED.
- **demo AWS posture** → `safety.md §2.9` → `stack-injection/gen_injected_override.py:382` DROPS jobsim `~/.aws` mount (`!reset null`, M217/DEF-M215-04) — demo has ZERO AWS anywhere. **Wire app creds as `AWS_*` ENV, not a mount.**
- **secret-coverage DNA** → `secrets-spec.md` → `rext/stack-secrets/secretdna/secret-dna.json` (`app` = 5 genes today; Delivers adds a Bedrock class). AI-provider keys = "documented-as-absent" (M50) → Bedrock is a NEW present class.

### Build notes (from audit)
- `CLAUDE_CODE_USE_BEDROCK` is NOT read by `bedrock.go` (Claude Code CLI convention) — inert-but-harmless for app; provisioning it does nothing for askengine. The AWS SDK default chain reads `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`/`AWS_SESSION_TOKEN`/`AWS_REGION`.
- staging uses 3 AWS keys (perm IAM user, no session token); hyper-studio set is a 5-key superset (likely STS session creds). Provision what the source carries.
- Without creds: SSE stream opens but agentic loop fails ("no EC2 IMDS role found" on VMs). This is the "opens but doesn't answer" signature to watch for in live verify.

## Build findings (2026-07-21)

### §1 talk-to-data
- **Bedrock cred source:** hyper-studio/.env carries real permanent IAM creds (AWS_ACCESS_KEY_ID/SECRET/REGION + CLAUDE_CODE_USE_BEDROCK; NO AWS_SESSION_TOKEN). NOT a blocker. `~/.aws/credentials` is empty (0 bytes).
- **Live proof:** `aws bedrock-runtime converse --model-id eu.anthropic.claude-sonnet-4-6 --region eu-west-1` with the wired creds → `pong` / `end_turn` / 20 tokens. Gate (b) creds/region/permission all work.
- **Manager hero IS admin:** `users.go:141,484` — hero role is vantage-faithful (manager→admin). So dan-manager satisfies the talk-to-data `isAdmin` gate; combined with the baked flag, the surface is reachable for the manager hero.
- **Coverage check:** with creds present, `check --demo` → Critical 100.0% exit 0 (Overall 68.5% — the partial local `.agentspace/secrets` lacks STRIPE/OPENAI, pre-existing, non-gate). The 2 AWS creds PASS key-present+nonempty.

### §2 #4 library — the :5050 carry is ALREADY RESOLVED
- apps/web `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` is baked OFFSET by `up-injected.sh:744` (`--build-arg ...=$SCHEME://$HOST:$((5050+OFFSET))/graphql`), overriding `next-web-app/Dockerfile.dev:23` ARG default `http://localhost:5050/graphql` (ENV bakes it). Hiring identical (line 1023).
- Remaining `:5050` refs: the Dockerfile ARG default (overridden) + `packages/graphql/codegen.ts` (dev-authoring schema introspection, not runtime). Neither reaches the running client.
- M218/M220 image-reuse check rebuilds a stale-endpoint image; M237 re-triage found the grid populates (7→29). **No client-fetch race defect remains** → verdict recorded, not a manufactured fix.

### §3 #1 hierarchical menu
- M237-resolved (renders grouped Organization structure for managers on a fresh build). §3 = live presence-verify + a coverage-sweep assertion authored against the live manager render (coverage-manifest.ts SectionDescriptor: region selector + realContent text/count). Manager descriptors need live calibration (manifest note) — hence author against the live render, not blind.

### LIVE VERIFY — demo-1 localhost, cold reset-to-seed (2026-07-21) — ALL GREEN
- Backend container `demo-1-backend-1`: all 4 Bedrock keys PRESENT (values-blind masked `<set>`). Bridge worked live.
- talk-to-data: manager `dan-manager` → `/enterprise/talk-to-data` → asked member count → **"Cervato Systems has 51 members"** (backend `ask.stream`: tool_use → `query_postgres` SELECT COUNT(*) → end_turn, ~7 s). `talk-to-data-m239.spec.ts` GREEN 11.7 s.
- #4 library: `/library/skill-paths` populates first-load (offset endpoint `:15050` baked, not `:5050`). GREEN.
- #1 menu: grouped Organization nav renders for the manager. GREEN. (`enterprise-surfaces-m239.spec.ts`, both GREEN 7.5 s.)
- Infra finding: first bring-up crashed `redis exited(1)` = Docker VM disk FULL (`No space left on device` on redis appendonlydir) from 3× 3.7 GB frontend builds. Recovery: `docker builder prune -af` (~25 GB) → cold rebuild → up. See decisions.md F1.
- rext test specs live at tag `sound-check-m239-live-proof`; the flag+DNA+bridge at `sound-check-m239-enterprise-surfaces`; the bridge log-truthfulness fix at `sound-check-m239-bridge-log-fix`.

_(will accumulate more topic → doc → code triples during build)_
