# M8 — Retro

**Summary:** The consolidation/discoverability milestone — turned the M3 demo stacks + the M7 seeding stack into
a **usable, documented product**: a demo-env corpus family + 3 end-to-end recipes, 3 curated seed presets (2
seed-proven), the `/demo-seed` skill, and the v1.0 **express-gate CI carry-forward** (validated 9/9 locally).
The two M3-deferred injection recipes (the `api.clerk.com` cert-redirect + the browser-login walk-through)
landed as part of the recipe set. The last v1.1 milestone — `/developer-kit:close-release` next.

## What went well
- **Everything was validated, not just authored.** The presets actually seed (mid-500 + large-1k end-to-end,
  isolation clean); the express-gate CI step was run locally with a freshly `npm install`ed `@clerk/express`
  (9/9, exit 0) before being committed; every doc link was machine-checked (8 files, all resolve). No
  "documented-as-designed but untested" claims.
- **The 2 M3 Fate-3 injection recipes found a clean home** in `recipe-browser-login.md` — the cert-redirect and
  the publishable-key walk-through are the two halves of the *interactive* demo, so they belong together.
- **The family-index pattern matched the corpus convention** (`corpus/ops/demo/` parallel to the `staging-*`
  family) — the demo story is now discoverable from corpus/README, root README, and CLAUDE.md.
- **The CI carry-forward closed a real v1.0 loop** — the `@clerk/express` gate was un-wired because it "needs a
  Node step"; now it's wired + proven, so a Clerk-Express bump turns the build red within minutes.

## What didn't / constraints
- **The deployment/injection (4th) gate stays a local gate** — it needs the platform's `colony` via GH_PAT,
  which a public CI runner can't supply cleanly; documented as a local/demo-stack gate rather than forced into
  CI.
- **The cert-redirect is a documented recipe, not a live-wired step** — the fake BAPI is built + alignment-gated,
  but the DNS/cert wiring into a demo stack is the recipe (consistent with M3's "recipe-only" status). The
  backend authn seam (the 403→200 path) works without it; the cert-redirect is for the orgclient reads.

## Carried forward → close-release / v1.2
- **close-release** ships v1.1 "show floor": merge `release/01.10-show-floor` → main, tag `v1.1`, archive the
  release records. A release-level review should re-confirm the four alignment gates + the seeding stack.
- **v1.2 "richer demo worlds"** (the natural successor): the data-DNA `waived` surfaces — taxonomy (skiller
  snapshot) + content (Directus snapshot-replay) — plus AI-generated rich content (transcripts/embeddings).

## Metrics
See [metrics.json](metrics.json). 5 sections, all green. 3 presets (2 seed-proven end-to-end), 4 recipe/family
docs, the `/demo-seed` skill, the express-gate CI carry-forward (9/9 validated), all doc links resolve. No new
test surface (docs + curation); validation was live (seeds + gate + link-check).
