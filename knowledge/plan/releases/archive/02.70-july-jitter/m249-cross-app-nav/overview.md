---
milestone_shape: section
milestone: M249
title: "cross-app navigation (Back-to-Cockpit + studio prod-eject fix)"
status: archived
release: v2.7 "july jitter"
depends_on: [M246]
parallel_with: [M247, M248, M250, M251, M252]
complexity: medium
created: 2026-07-23
last_updated: 2026-07-24
---

# M249 — cross-app navigation (Back-to-Cockpit + studio prod-eject fix)

## Goal
A "← Back to Cockpit" item in all 4 sub-app menus (next-web, hiring, studio-desk, ant-academy), and the
studio-desk logo / back / logout controls stop ejecting the presenter to production. This milestone **owns the
first-ever studio-desk source demopatch machinery**.

## Shape (why this shape)
`section`. The deliverables are a finite, enumerable set of disjoint demopatch lanes over a **confirmed**
mechanism — the demopatch anchor→replace ladder, which already supports insertion (R4 mitigation) — plus one
wiring-integration lane and one doc lane. There is no exploratory measure→fix loop: the only novelty is the
first studio-desk *source* patch machinery, and that is a known-shape extension of the existing
`build_frontend_*` demopatch ladder (which M253 will later extend again). Build with
`/developer-kit:build-milestone`.

## Scope

### In
- **3 new demopatches** adding a `NEXT_PUBLIC_COCKPIT_URL` / `VITE_COCKPIT_URL` (`7700`+OFFSET) "Back to
  Cockpit" item:
  - **(A)** `next-web-back-to-cockpit` — shared `packages/ui/NavbarTop`, so it covers **both** web + hiring.
  - **(B)** `studio-desk-back-to-cockpit` — **+ fix the `pageWrapper.js:149` logo / `userProfile.js:147,302`
    back+logout prod-ejects** (the same scaffold).
  - **(C)** `ant-academy-back-to-cockpit` — `UserMenu.jsx:143`.
- **Wire the offset-URL bake + apply/revert** into `up-injected.sh` (both next-web overlays + the **net-new**
  `build_frontend_studio_desk` patch machinery) + `ant-academy.sh` (self-contained). **Fail-closed when the env
  is unset.**
- **Author the additive-UI injection pattern doc** + the cockpit-spec return-nav section.

### Out
- Any platform edit.
- The studio blank-page (M253).
- The builder keys (M252).

## Dependencies & parallelism
- **depends_on:** **M246** (the HARD go/no-go re-sync barrier — it re-points the seeder and touches
  `up-injected.sh`; **all fan-out worktrees branch from post-M246 HEAD**).
- **parallel_with:** **M247, M248, M250, M251, M252** — built concurrently as `work-milestone --worktree=<path>`
  agents. **M253 is serial AFTER M249** — it extends the `build_frontend_studio_desk` studio patch ladder this
  milestone creates.
- **Merge/close order:** M251 → { M248, M250 } → **M249** → M253 → M252 → M247-reconcile → M254.
- **Intra-milestone lanes** — parallelism **high (~2×; 3–4× on authoring)**: 3 disjoint patch lanes + 1 doc lane,
  then a serial bottleneck.
  - **Lane A — next-web** (`next-web-back-to-cockpit`, shared `packages/ui/NavbarTop` → web + hiring).
  - **Lane B — studio-desk** (`studio-desk-back-to-cockpit` + the `pageWrapper.js:149` / `userProfile.js:147,302`
    prod-eject fixes, same scaffold).
  - **Lane C — ant-academy** (`ant-academy-back-to-cockpit`, `UserMenu.jsx:143`) — **fully self-contained via
    `ant-academy.sh`.**
  - **Lane D — docs** (the additive-UI injection pattern + the cockpit-spec return-nav section).
  - **Serial bottleneck:** the **`up-injected.sh` integration** — the one shared file (offset-URL bake +
    patch-set fingerprint + studio apply/revert + net-new `build_frontend_studio_desk` machinery). Runs after
    A/B/D land.
- **Recommended subagents:** up to **4 concurrent** (Lane A / B / C / D), then a **single serial agent** for the
  `up-injected.sh` integration.
- **Coordination guardrails:** M249 owns the demo-stack *patch* python tests (M251 owns the *health/inventory*
  tests — no overlap); M249 owns the first studio patch ladder in `up-injected.sh build_frontend_studio_desk`
  (M253 extends it; M252's env wiring lives in the disjoint `gen_injected_override.py`); the `CLAUDE.md` one-line
  bullet **defers to M247** (sole owner); the studio spec docs are **reconciled in the M247-tail**; rung-zero
  (push rext tags to **origin**) before any billion re-pin.

## KB dependencies
- `corpus/ops/demo/demopatch-spec.md` — the demopatch anchor→replace mechanism + the 7 guards + the manifest schema.
- `corpus/ops/demo/cockpit-spec.md` — the presenter cockpit (the return-nav target).
- `corpus/ops/demo/frontend-tier.md` — the demo UI tier (next-web + studio-desk + ant-academy bring-up).

## Delivers
- `corpus/ops/demo/cockpit-spec.md` — the return-nav section.
- `corpus/ops/demo/demopatch-spec.md` — the additive-UI injection pattern + the 3 new patch rows + studio-desk
  registered as the **first *source* patch**.
- `corpus/ops/demo/frontend-tier.md` / `corpus/services/studio-desk.md` — the offset-URL / prod-eject fix (studio
  spec docs authored here in their own subsections, **reconciled in the M247-tail**).

## Open questions
- Rewrite studio's existing hardcoded back-to-prod item vs add a sibling (rewriting also fixes the prod-eject)?
- Add a `DEMO_NO_BACK_TO_COCKPIT` opt-out knob?
- Demo-path only (the cockpit is demo-only)?
