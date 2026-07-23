---
title: "KB Fidelity Audit — M239 enterprise surfaces"
date: 2026-07-21
scope: milestone:M239
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| talk-to-data flag gate | `corpus/architecture/ai_architecture.md:51`, `corpus/services/backend.md:19,134` | `next-web-app/apps/web/src/hooks/useTalkToDataAccess.ts`, `.../enterprise/talk-to-data/page.tsx` | PAIRED |
| talk-to-data Bedrock creds | `corpus/ops/staging-bringup.md:188-198` (existing anchor), `corpus/architecture/ai_architecture.md:51` | `app/internal/askengine/bedrock.go` | PAIRED |
| secret-coverage DNA (app Bedrock class) | `corpus/ops/secrets-spec.md` | `rext/stack-secrets/secretdna/secret-dna.json`, `catalog.go` | PAIRED (Delivers extends) |
| secrets posture (AWS in demo) | `corpus/ops/safety.md` §2.9 | `rext/stack-injection/gen_injected_override.py:382` | PAIRED (Delivers adds note) |
| #4 library / :5050 endpoint | `corpus/ops/demo/frontend-tier.md` (endpoint wiring) | `rext/demo-stack/up-injected.sh`, `next-web-ssr-graphql-origin` patch | PAIRED |
| #1 hierarchical manager menu | `corpus/ops/demo/coverage-protocol.md` | next-web nav (manager vantage) | PAIRED (M237 resolved) |

## Fidelity Findings (Phase 2)

1. **talk-to-data flag gate — ALIGNED.** Doc/plan claim: gated by `isAdmin && flag_enable_talk_to_data`, demo escape `NEXT_PUBLIC_DEMO_FLAGS_ALL === 'true'`. Code (`useTalkToDataAccess.ts:9,31,39-42`) matches exactly: `demoBypass` returns `{isEnabled: isAdmin}`; otherwise `isAdmin && effective === true`. **Verdict: ALIGNED.**

2. **Bedrock engine — ALIGNED.** Plan claim: AWS Bedrock, model `eu.anthropic.claude-sonnet-4-6`, region `eu-west-1`, default AWS credential chain, SigV4. Code (`bedrock.go:24-25,68-103`) matches: `DefaultModelID = "eu.anthropic.claude-sonnet-4-6"`, `DefaultRegion = "eu-west-1"`, `config.LoadDefaultConfig` (env→shared→IAM chain), SigV4 via `bedrock.WithConfig`. **Verdict: ALIGNED.** Note: `bedrock.go` reads `AWS_REGION` + `ASK_MODEL_ID`; the AWS *credential* env vars (`AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`/`AWS_SESSION_TOKEN`) are consumed by the AWS SDK's default chain, not by askengine directly. **`CLAUDE_CODE_USE_BEDROCK` is NOT read by `bedrock.go`** (it is a Claude Code CLI convention) — inert-but-harmless for the app; provisioning it is fine, it just does nothing for askengine.

3. **Demo carries no AWS creds; ~/.aws mount is DROPPED — ALIGNED.** Plan claim: the `app` compose service has no AWS creds; only `jobsimulation` mounts `~/.aws`. Code reality (`gen_injected_override.py:382-409`): the injection override **drops** even jobsimulation's `~/.aws` bind (`volumes: !reset null`, M217/DEF-M215-04 — the empty-dir EISDIR bug) and the comment states "a demo carries NO AWS credentials at all (zero AWS_ACCESS_KEY_ID / AWS_SECRET_ACCESS_KEY anywhere in platform/.env)". **Build implication:** wire the app's Bedrock creds as **`AWS_*` env vars** on the app service block (NOT a re-mounted `~/.aws`, which the M217 fix deliberately removes). **Verdict: ALIGNED** (and sharper than the overview: the demo has zero AWS anywhere, not "app lacks what jobsim has").

4. **staging-bringup.md is a strong pre-existing anchor — ALIGNED.** `staging-bringup.md:188-198` already documents talk-to-data needing `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`/`AWS_REGION=eu-west-1`, the "no EC2 IMDS role found" failure mode without them, and that the creds are a prod-inference-profile exception. **Minor discrepancy (build note, not staleness):** staging uses **3** AWS keys (no `AWS_SESSION_TOKEN`, no `CLAUDE_CODE_USE_BEDROCK`); the hyper-studio key set is a **superset** of 5. Both are correct for their env — staging is a permanent IAM user; hyper-studio may use STS session creds (hence the session token). Provision what the source carries.

5. **ai_architecture.md / backend.md — ALIGNED.** `ai_architecture.md:51` ("plus Talk to Data (Bedrock)") and `backend.md:19,134` (askengine, SSE, SQL sandbox, Bedrock) accurately describe the code. **Verdict: ALIGNED.**

## Completeness Gaps (Phase 3)
- **secrets-spec.md** does not yet enumerate an AWS/Bedrock cred class for `app` — this is **expected and covered by the milestone's `Delivers → secrets-spec.md (Bedrock cred class for app)`**. Not a blind area; it is planned knowledge production.
- **safety.md** does not yet carry a note on the demo `app` holding real cloud creds — covered by the milestone's `Delivers → safety.md secrets-posture note`.

## Applied Fixes
- None needed inline. All existing claims ALIGNED. spec-notes.md updated with topic→doc→code triples.

## Open Items (require user decision)
- None at audit time. (The real-cred provisioning itself is a build step; if the hyper-studio creds turn out to be absent/placeholder, that becomes a build-time blocker per the milestone brief — not an audit finding.)

## Gate Result
GREEN — proceed to Phase 1. Every topic PAIRED with accurate docs; the two doc gaps are explicit milestone Delivers (planned knowledge production), not blind areas; no stale load-bearing claims. Notable: more pre-existing anchor than the roadmap assumed (staging-bringup.md documents the AWS-creds-for-talk-to-data pattern), and one code fact sharper than the overview (the demo drops ALL AWS, so wire env vars not a mount).
