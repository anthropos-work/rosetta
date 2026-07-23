---
milestone_shape: section
milestone: M252
title: "studio-desk builder enablement"
status: planned
release: v2.7 "july jitter"
depends_on: [M246]
parallel_with: [M247, M248, M249, M250, M251]
complexity: medium
created: 2026-07-23
---

# M252 — studio-desk builder enablement

## Goal
The studio `sim-advanced-builder` + `sim-guided-builder` work — i.e. the provisioned studio-desk AI
key actually **reaches the demo container** at runtime — proven by a Playthrough. Today the backend
`/api/ai/completion` **500s with no provider** because the key never lands in the container; this is a
demo-**wiring** gap, not a missing secret.

## Shape (why this shape)
`section`. The scope is a small set of bounded, enumerable deliverables — one wiring fix, one DNA
assertion, one net-new Playthrough — each with a clear done-state and a single verify. The root cause
is already understood (a container-wiring gap, not an exploratory failure), so there is no
measure→iterate loop to run; a straight build → verify-once → document pass suffices. (Contrast the
release's iterative siblings M250 AI-readiness fidelity and M253 studio first-paint, which ARE
measurement-driven.)

## Scope
### In
- **AI-key demo-container wiring.** Wire the provisioned studio-desk AI key into the demo container at
  runtime — add `env_file: <clone>/studio-desk/.env` to the studio-desk service block in
  `gen_injected_override.py` (mirrors hiring-app), **OR** a `bridge_studio_ai_creds()` in
  `up-injected.sh` mirroring `bridge_bedrock_creds()`. **Not a DNA gap** — the
  `AI_OPENAI_API_KEY` / `AI_ANTHROPIC_API_KEY` genes exist and are provisioned; the gap is that the
  value never reaches the container, so `/api/ai/completion` 500s.
- **DNA hardening.** A demo-aware assertion that the studio-desk **container** carries a provider key
  (close the `.env`-vs-container gap the wiring fix addresses).
- **Builder Playthrough.** A net-new Playthrough proving the builders generate: a
  `studio-builder-page.ts` page-object + `studioBaseUrl(9000+offset)` + studio Clerkenstein hero-login
  + `manifest/studio-builders.yaml` + an admin / content_creator precondition. (studio-desk is **not**
  in the playthroughs manifest today — this is its first entry.)
- **Talk-to-data double-check.** Re-confirm the M239 Bedrock talk-to-data path — **confirmed COMPLETE,
  no gap** (recorded here for the audit trail; no work owed).

### Out
- The studio blank-page / first-paint perf (M253).
- The studio nav / logo / "Back to Cockpit" + prod-eject fixes (M249).
- Any platform-repo edit.

## Dependencies & parallelism
- **depends_on:** M246 (the HARD go/no-go re-sync barrier). This worktree branches from **post-M246
  HEAD** — M246 touches `up-injected.sh` + the seeder package, and the re-point/pin-bump must be in
  place so the wiring fix and Playthrough are scoped against the consolidated platform.
- **parallel_with:** M247, M248, M249, M250, M251 — a fan-out lane off the M246 barrier. (M253 is
  **not** a peer: it is serial-after-M249 and closes before M252 in the merge order.)
- **Cross-milestone coordination (the "don't clobber each other" rules):**
  - M252's env wiring lives in `gen_injected_override.py` (a **disjoint function**) — no collision with
    the `up-injected.sh build_frontend_studio_desk` studio patch ladder M249 owns / M253 extends.
  - `run-unit.sh` roster is **M251-owned**: if M252 adds a `*.unit.spec.ts` it must be rostered —
    coordinate that one line with M251.
  - Studio spec docs (`studio-desk.md`, `frontend-tier.md`) are written by M249 / M252 / M253 in their
    **own subsections**, reconciled in the **M247-tail** — M247 is sole owner of `CLAUDE.md`, so M252
    defers any one-line index bullet to it.
  - **Rung-zero:** any rext tag M252 cuts is `git push --tags` to **origin** before billion re-pins.
  - **Merge/close order:** … → M253 → **M252** → M247-reconcile → M254 (M252 closes after M253).
- **Intra-milestone LANE decomposition:** **3 disjoint concurrent lanes**, then a **serial verify
  bottleneck**:
  - **Lane A** — wiring-fix + DNA-hardening (`gen_injected_override.py` / `up-injected.sh` +
    the demo-aware container-key assertion).
  - **Lane B** — builder Playthrough (`studio-builder-page.ts` + `manifest/studio-builders.yaml` +
    the studio Clerkenstein hero-login + `studioBaseUrl` + the admin/content_creator precondition).
  - **Lane C** — docs (the studio-desk.md / secrets-spec.md / playthroughs.md deltas).
  - **Serial bottleneck:** verification — the Playthrough proves the fix and needs a live demo
    bring-up (it also makes real LLM calls, R6, so it runs behind a cost ceiling). Recommended:
    **3 concurrent subagents** (A ∥ B ∥ C), then a single serial verify tail on one demo.

## KB dependencies
- `corpus/services/studio-desk.md`
- `corpus/ops/secrets-spec.md`
- `corpus/ops/demo/frontend-tier.md`
- `corpus/ops/demo/playthroughs.md`

## Delivers
- `corpus/services/studio-desk.md` — the demo-aware studio-desk AI note (container-wiring model).
- `corpus/ops/secrets-spec.md` — the demo-aware studio-desk AI note (`.env`-vs-container coverage).
- `corpus/ops/demo/playthroughs.md` — the builder Playthrough + the updated live-Playthrough count.

## Open questions
- `env_file` vs a bridge — mount the studio-desk `.env` directly, or bridge the values out of a mounted
  file the way `bridge_bedrock_creds()` does?
- Which provider (`AI_PROVIDER_CHAIN`) for cost / latency in the demo?
- A real-LLM Playthrough needs a **cost ceiling** (R6) — default-on, or gate it behind a `DEMO_NO_*`
  knob?
