---
iter: 12
milestone: M211
iteration_type: tik
status: closed-fixed
created: 2026-07-08
---

# iter-12 — tik: kill the 40 prod-eject escapes (bake the demo-local AI-Academy URL)

**Active strategy reference:** TOK-01 move (4). Clears the highest-impact M42e residual from iter-11.

## Step 0 — Re-survey
iter-11 proved the federation fix (coverage 8→1 failing, 2→0 persona, 7→50 reachable). Two residuals remain:
**escapes=40** (a hardcoded prod `aiacademy.anthropos.work` nav link on every page) and **failingSections=1**
(`/library/ai-simulations` grid empty — `public.simulation_embeddings` doesn't exist; a separate data gap →
iter-13). This tik targets the escapes.

## Cluster / target identified
`urls.ts:17` already reads `process.env.NEXT_PUBLIC_ACADEMY_URL || 'https://aiacademy.anthropos.work'` — a
per-URL override EXISTS (no demopatch needed), but the demo never SUPPLIED the value → the global "AI Academy"
nav link ejected to prod on all 40 reachable pages.

## Hypothesis
Bake `NEXT_PUBLIC_ACADEMY_URL=http://localhost:$((3077+OFFSET))` (the demo's own Clerk-free ant-academy) into
the next-web `.env.local` overlay (mirroring the existing `NEXT_PUBLIC_STUDIO_URL` bake) → the nav link resolves
demo-local → `escapes 40→0`.

## Expected lift
`escapes 40→0`. `failingSections` stays 1 (sim-embeddings, iter-13). So M42e = 1-section-short after this tik.

## Phase plan
1. **Fix (landed):** `up-injected.sh` bakes `NEXT_PUBLIC_ACADEMY_URL` + a contract fence (all 64 tests GREEN,
   shellcheck clean). Commit; move tag `quick-change-m211`; re-pin consumption clone.
2. **Rebuild + prove:** force-rebuild `demo-1-next-web` (the offset-guard validates the WUNDERGRAPH endpoint,
   not ACADEMY_URL, so remove the image first) with the new `.env.local` → recreate the container → re-run the
   M42e sweep → measure `escapes`.

## Escalation conditions
- If baking ACADEMY_URL to :13077 makes the crawler FOLLOW into ant-academy (out of crawl-scope) and surface
  new failures → confirm the crawl treats cross-port links as demo-local-not-followed, else route.

## Acceptable close-no-lift outcomes
If escapes drop but a new escape class surfaces (a different prod link on the now-reachable pages), record +
route it — the ACADEMY_URL fix still landed.
