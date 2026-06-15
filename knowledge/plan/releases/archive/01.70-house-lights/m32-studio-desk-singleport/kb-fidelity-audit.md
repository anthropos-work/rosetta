---
title: "KB Fidelity Audit — M32 studio-desk single-port / production"
date: 2026-06-15
scope: milestone:M32
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| studio-desk override env (NODE_ENV/single-port) | `corpus/ops/demo/frontend-tier.md` §What `/demo-up` brings up; demo-up SKILL | `stack-injection/gen_injected_override.py` FRONTENDS + `frontend_lines` | PAIRED |
| studio-desk prod vs dev route serving | (none — platform repo, byte-pristine) + `corpus/services/studio-desk.md` | `stack-demo/studio-desk/src/index.ts` 148-272 | PAIRED (code-read) |
| offset-origin CORS allowlist | `corpus/ops/demo/frontend-tier.md` §Offset-origin CORS | `gen_injected_override.py:248-250` | PAIRED |
| single-port regression coverage | `corpus/ops/demo/frontend-tier.md` (port story) | `stack-injection/tests/test_injection.py` ~819-857 | PAIRED |

## Fidelity Findings

1. **frontend-tier.md:21 — studio-desk port story is STALE (doc-lag, = planned M32 deliverable).**
   - Source: `corpus/ops/demo/frontend-tier.md:21` (+ example line 24, §verification line 157-158).
   - Expected (doc): "**9100** (frontend) + **9000** (backend), each + N×10000".
   - Actual (code): `FRONTENDS` studio-desk emits a single port `[(9000, 9000)]`; the comment (lines 51-54) states the production image is single-port 9000 and "no dead 9100".
   - Verdict: STALE. Fix owner: update doc. **This is a planned M32 sweep item, not a blocker** — code is correct, doc lags.

2. **demo-up SKILL — `:9100` references are STALE (doc-lag, = planned M32 deliverable).**
   - Source: `.claude/skills/demo-up/SKILL.md:68-69` ("studio-desk `:9100+`").
   - Same divergence as #1; the code is single-port 9000+offset. Fix owner: update doc (planned sweep).

3. **gen_injected_override.py:249 CORS still lists the un-offset 9100 origin.**
   - The CORS allowlist enumerates `(3000, 3001, 9000, 9100)`. The `9100` origin is dead now studio-desk
     is single-port production. Fix owner: remove (planned M32 sweep + explicit decision M32-D).

4. **test_injection.py:925 — STALE assertion (latent, env-masked).**
   - `test_frontend_blocks_parse_to_valid_compose` asserts `sorted(sd["ports"]) == ["29000:9000", "29100:9100"]`.
     The generator emits studio-desk single-port `["29000:9000"]` only. The test is **skipped** (PyYAML absent
     in this env), masking the inconsistency. Fix owner: update test to `["29000:9000"]`. Folded into the M32 sweep.

## Completeness Gaps

- **Route coverage (the milestone's load-bearing open question), VERIFIED via code-read of
  `stack-demo/studio-desk/src/index.ts`:** the production (`isProduction=true`, else branch lines 207-272)
  path serves every route the dev block (lines 148-206) handled — `/`, `/home`, `/catalog`,
  `/simulation-builder`(+`.html`), `/sim-advanced-builder`, `/sim-guided-builder`, `/builder-skill-path`,
  `/generation`, `/skills` — and ADDS `/simulations`, `/academy`, an `express.static` mount over `dist/public`
  (which serves all the dev block's `.html`-extension targets, e.g. `home.html`/`catalog.html`), and an
  `index.html` SPA `*` fallback. The dev `*` catch-all only forwarded to the Vite dev server (dead in a
  container without Vite); its production analog (SPA `index.html`) is strictly better. **No route gap →
  flipping NODE_ENV=production is safe.** Recorded in decisions.md.

## Applied Fixes
None inline — every finding (1-4) is a planned M32 deliverable; fixing them now would pre-empt the build loop.
Findings are tracked in this report + spec-notes triples + decisions.md.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. All topics PAIRED; the doc/CORS/test staleness items (1-4) ARE the milestone's
planned deliverables (code is already single-port; docs/CORS/the env-skipped test lag). No blind areas, no
load-bearing stale claim the implementer would read as truth (the implementer is the one fixing them).
