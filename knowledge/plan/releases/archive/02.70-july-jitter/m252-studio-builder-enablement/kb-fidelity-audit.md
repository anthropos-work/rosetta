---
title: "KB Fidelity Audit — M252 studio-desk builder enablement"
date: 2026-07-24
scope: milestone:M252
invoked-by: build-mstone-iters (Phase 0b pre-flight) / user
---

## Verdict
**YELLOW** — no blind areas; the two stale load-bearing claims were fixed inline; the remaining
gaps are incidental completeness items that M252 is already chartered to Deliver. Proceed with the
tracked items recorded in `decisions.md` / `spec-notes.md`.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| studio-desk builder pages + AI Copilot `/api/ai/completion` | `corpus/services/studio-desk.md` | `stack-dev/studio-desk/{vite.config.ts, src/index.ts, src/routes/ai.ts, src/services/ai*}` | PAIRED |
| AI-key demo-container wiring (studio-desk service block) | `corpus/ops/demo/frontend-tier.md` (+ secrets-spec Bedrock bridge) | `stack-injection/gen_injected_override.py`, `demo-stack/up-injected.sh` | PAIRED |
| Secret-coverage DNA (studio-desk AI genes) + demo AI-key policy | `corpus/ops/secrets-spec.md` | `stack-secrets/secretdna/secret-dna.json` | PAIRED |
| Builder Playthrough (studio-desk enters manifest) | `corpus/ops/demo/playthroughs.md` | `playthroughs/{manifest, e2e/lib/stack-env.ts}` | PAIRED (doc current; code to be extended) |
| DNA-hardening: container-carries-provider-key assertion | `corpus/ops/secrets-spec.md` (`.env`-vs-container) | (net-new; no existing container-key assert) | DOC-ONLY → milestone Deliverable |

## Fidelity Findings

### F1 — studio-desk.md entry-point list is STALE (load-bearing) — FIXED
- **Source:** `corpus/services/studio-desk.md:34`
- **Expected (doc):** entry points include `designer-sim.html`.
- **Actual (code):** `vite.config.ts:28-42` — **`designer-sim.html` does NOT exist** (0 hits tree-wide). The real
  builder pages the doc OMITS are `simulation-builder.html`, `sim-advanced-builder.html`, `sim-guided-builder.html`
  (each with dev+prod Express routes `src/index.ts:191-250` and prompt dirs `src/prompts/sim-{advanced,guided}-builder/`).
- **Verdict:** STALE. Load-bearing — M252 Lane B's `studio-builder-page.ts` drives exactly `sim-advanced-builder`/`sim-guided-builder`.
- **Fix owner:** doc. **Applied inline** (corrected the list + added a "simulation-builder family" note).

### F2 — studio-desk.md admin gate role list is STALE (load-bearing) — FIXED
- **Source:** `corpus/services/studio-desk.md:288`
- **Expected (doc):** eligible role = `admin` / `org:admin` only.
- **Actual (code):** `src/index.ts:96` — `STUDIO_ACCESS_ROLES = ['admin', 'org:admin', 'content_creator', 'org:content_creator']`;
  gate `checkEnterpriseAndAdmin` (`src/index.ts:99-128`) redirects non-eligible/non-org users to `WEB_APP_URL`.
- **Verdict:** STALE (too narrow). Load-bearing — M252's Playthrough precondition is "admin / **content_creator**", which matches CODE, not the old doc.
- **Fix owner:** doc. **Applied inline** (role list now includes `content_creator` / `org:content_creator`).

### F3 — studio-desk.md AI Copilot provider defaults — ALIGNED
- Default `AI_PROVIDER_CHAIN=azure-openai,openai` (`.env.example:46`), `AI_DEFAULT_TIER=fast` (`.env.example:50`),
  in-code fallback tier `thinking_fast` (`config.ts:182`), 4 tiers per provider (`config.ts:20-39`) — all confirmed. (studio-desk.md:104, 177-179.)

### F4 — studio-desk.md does not name `/api/ai/completion` — INCIDENTAL
- The doc refers to "Copilot" / generic `/api/ai` routes; the failing endpoint is specifically `POST /api/ai/completion`
  (`src/routes/ai.ts:36`), which **500s `{error:'Failed to get AI completion'}`** when no provider is configured
  (`aiService.ts:178` throws `No AI providers available` → route catch `ai.ts:56-60`). Confirms the milestone premise.
  M252 Deliverable adds the demo-aware note. Incidental (not wrong, just less specific).

### F5 — secrets-spec.md DNA facts — ALIGNED
- Version `sound-check-m239`, **61 genes**, split **42 required / 11 optional / 8 waived, 13 critical** — matches
  `secret-dna.json` exactly. studio-desk = **7 genes**; `AI_OPENAI_API_KEY` + `AI_ANTHROPIC_API_KEY` are both
  **required · standard** (confirms the milestone premise: the genes exist and are provisioned — NOT a DNA gap).

### F6 — secrets-spec.md M50 "AI-keys are absent/waived-optional" is imprecise for studio-desk (load-bearing) — TRACKED
- **Source:** `corpus/ops/secrets-spec.md:269-278` (the M50 "AI-provider keys policy — documented-as-absent" box).
- **Expected (doc):** AI-provider keys "remain in the **waived / optional** class for a demo source"; "their absence
  is correct, not a coverage hole."
- **Actual (code):** studio-desk's `AI_OPENAI_API_KEY` / `AI_ANTHROPIC_API_KEY` are **required · standard** (not
  waived/optional) and **not** in `demoSatisfied` — so a creds-less demo **warns** (standard, non-fatal) on them,
  the SAME posture as the M239 Bedrock class documented directly below. The blanket "waived/optional" claim does not
  hold for the studio-desk AI genes, and the M50 box does not yet describe the **studio-builders live-AI surface**
  (the third live-AI-in-demo surface, after Talk-to-Data).
- **Verdict:** STALE/imprecise, but the milestone's premise does NOT rely on the M50 box (the roadmap correctly says
  "the genes exist and are provisioned"), and M252 explicitly **Delivers** "the demo-aware studio-desk AI note
  (`.env`-vs-container coverage)." **Not fixed inline** — deferred to the milestone's own doc deliverable to avoid
  pre-empting the env_file-vs-bridge decision. Tracked as KB-1.
- **Fix owner:** doc — land as part of the M252 secrets-spec.md deliverable (add a studio-builders carve-out mirroring the M239 Bedrock carve-out).

### F7 — secrets-spec.md Bedrock bridge (M239 double-check) — ALIGNED (no gap)
- `bridge_bedrock_creds()` (`up-injected.sh:1244-1271`, call `:1420`) copies the 5-key AWS/Bedrock class
  `app/.env → platform/.env`, values-blind / idempotent / non-fatal — exactly as secrets-spec.md:311-320 describes.
  **The M252 "Talk-to-data double-check" is confirmed COMPLETE, no work owed.**

### F8 — frontend-tier.md studio-desk service-block description is INCOMPLETE (the gap M252 fixes)
- **Source:** `corpus/ops/demo/frontend-tier.md:560-566`.
- **Actual (code):** `gen_injected_override.py:145-181` + `frontend_lines():273-290` emit exactly what the doc lists
  (offset ports, `image: demo-N-*`, DIRECTUS_TOKEN strip→static, CLERK_* minted, CORS on backend) — **and NO
  `env_file:` and NO AI key.** studio-desk **inherits `env_file: .env` (= platform/.env / BASE_ENV)** from the base
  compose service (comment `:103`), which does **not** carry studio-desk's AI keys (those land in `studio-desk/.env`).
  The doc is currently **silent** on the AI-key-to-container path (not stale — it makes no false claim).
- **Verdict:** completeness gap the milestone Deliverable fills. Incidental until then.

### F9 — playthroughs.md counts + studio absence — ALIGNED
- **16 live Playthroughs, 0 TODO** (playthroughs.md:106) — exact. **No `studio` reference anywhere** in
  `playthroughs/manifest/` or `e2e/` → confirms "studio-desk is not in the playthroughs manifest today — this is its
  first entry." `stack-env.ts` has `appBaseUrl`/`fapiBaseUrl`/`hiringAppBaseUrl` (3001+offset) and **no
  `studioBaseUrl`** — the milestone's `studioBaseUrl(9000+offset)` cleanly mirrors the `hiringAppBaseUrl` pattern.

## Completeness Gaps
1. The **studio-builders live-AI surface** (the third live-AI-in-demo surface, after Talk to Data) is undocumented
   across studio-desk.md / secrets-spec.md / frontend-tier.md. **Chartered** as M252 Deliverables. (KB-1)
2. `studioBaseUrl` + studio Clerkenstein hero-login + the studio-desk manifest entry are absent from `playthroughs`
   code (expected — M252 Lane B builds them). playthroughs.md is not stale, only to-be-extended.

## Applied Fixes
- `corpus/services/studio-desk.md:34` — entry-point list corrected: removed the phantom `designer-sim.html`, added
  the real `simulation-builder` / `sim-advanced-builder` / `sim-guided-builder` pages + a "simulation-builder family" note (F1).
- `corpus/services/studio-desk.md:288` — Studio access-role list corrected to the real `STUDIO_ACCESS_ROLES`
  (adds `content_creator` / `org:content_creator`; names the `checkEnterpriseAndAdmin` gate) (F2).

## Open Items (require user decision)
None blocking. One **build-design observation** (not a doc defect) worth the developer's eye:
- **"mirrors hiring-app" is imprecise.** `hiring_lines()` emits `env_file: [<abs> platform/.env]` — pointing at the
  **shared platform `.env`**, which carries **no** AI keys — and hiring-app itself emits no AI key. A *literal* mirror
  would therefore NOT bring studio-desk's AI creds into the container. The provisioned studio AI keys live in
  **`studio-desk/.env`** (per the DNA `target_file`). So the wiring fix must point `env_file` at
  **`<clone>/studio-desk/.env`** (not platform/.env), OR add a `bridge_studio_ai_creds()` that copies the studio AI
  class `studio-desk/.env → platform/.env` the way `bridge_bedrock_creds()` does. The milestone's own open question
  ("env_file vs a bridge") already frames this; recorded here so the "mirrors hiring-app" shorthand isn't taken literally.

## Gate Result
**YELLOW: proceed with tracking.** No blind areas. The two stale load-bearing claims (F1, F2) are FIXED inline and
code-verified. Remaining items (KB-1 + the completeness gaps) are the milestone's own chartered doc Deliverables;
record KB-1 in `decisions.md`. The "mirrors hiring-app" nuance is a build-design note for Lane A, not a doc gap.
