---
iteration: 03
iteration_type: tik
status: in-progress
created: 2026-06-25
active_strategy: TOK-01 (line 1 — the Studio-link escape, now via the demo-patch mechanism)
---

# iter-03 — demo-patch tool + the Studio urls.ts patch (resolves RESCOPE-1)

## Active strategy reference
**TOK-01 line 1** — the Studio-link escape (the highest-leverage fix: 139 escapes → 0). iter-02 falsified the
env-rewrite hypothesis (next-web `STUDIO_URL` is a `NEXT_PUBLIC_NODE_ENV` ternary with no per-URL override) and
raised RESCOPE-1. The user RESOLVED RESCOPE-1 by choosing a NEW approach: a **demo-patch tool** — a rext-owned
source patch applied to the demo's EPHEMERAL clone (`stack-demo/next-web-app`, gitignored) before the build,
reverted after, so the image carries the fix while the canonical platform repos are NEVER touched. This is the
sanctioned resolution path (c)-pivot of RESCOPE-1. Design: `.agentspace/scratch/work-m42m/demo-patch-design.md`.

## Cluster / target identified
The single baked `studio.anthropos.work` left-nav escape (139 of the manager residual, all the same host),
which iter-02 proved is platform-bound under the zero-CANONICAL-edit line. The demo-patch tool patches the
demo's own ephemeral clone (NOT canonical) — keeping the zero-canonical-edit line intact while making the
demo's Studio link demo-local (`:39000`).

## Hypothesis
A content-anchored source patch to the demo clone's `packages/core-js/src/constants/urls.ts` (mirror
`ACADEMY_URL`: add `process.env.NEXT_PUBLIC_STUDIO_URL ||` before the existing ternary), baked into the demo
next-web image via a build-time `NEXT_PUBLIC_STUDIO_URL=http://localhost:39000` line on the existing
`apps/web/.env.local` overlay, reverted after the build → the running demo's Studio link resolves to the local
studio-desk (`:39000`), not prod. Escapes 139 → 0 for the Studio class.

## Phase plan (per the design's build plan, in order)
1. **Safety guards FIRST** — build the `demopatch` CLI in `demo-stack/patches/` (verbs apply|revert|status|check)
   implementing ALL 6 guards (G1 path-assert / G2 drift-refuse / G3 never-commit/working-tree-only / G4 idempotent
   / G5 self-revert / G6 demo-only) + the G3 never-mutate-git grep unit test + a `check` dry-run verb.
2. **Mechanism + manifest** — content-anchored YAML manifest schema; author the `next-web-studio-url` manifest
   (urls.ts env-aware replacement, `build_env`, scope:demo); compute `pre_sha256` + verify the anchor matches once.
3. **Wire into `up-injected.sh`** `build_frontend_next_web` — append the `NEXT_PUBLIC_STUDIO_URL` line to the
   existing `.env.local` overlay; `demopatch apply` NON-FATAL before the docker build; extend the RETURN trap with
   `demopatch revert`. + R1 pristine-ing pass (ensure-clones.sh head) + R2 push-block. Tag the rext section.
4. **VERIFY on a FRESH demo-up (demo-3)** — the authoritative gate: apply-then-revert (clone git-clean after) /
   the baked bundle carries `localhost:39000` not `studio.anthropos.work` / live click-through as dan-manager
   (left-nav Studio + ManagerHome cards → local studio-desk) / re-up G4 no-op / git status clean in BOTH the
   demo clone AND rosetta / a quick manager re-sweep to confirm the Studio class dropped (139 → 0).

## Expected lift
Escapes 139 → 0 for the Studio class (the manager (d) clause's only remaining escape). NOT gate-met (the
dashboard populate + fan-out exhaustion — TOK-01 lines 2-4 — remain a LATER run).

## Escalation conditions
- A NEW blocker that the demo-patch approach can't resolve (e.g. the cross-origin Clerk-session click-through
  fails for a session reason) → record + surface as re-scope-trigger / user-blocker per Phase 5.
- Any guard that cannot be implemented safely → user-blocker.

## Acceptable close-no-lift outcomes
A falsified verify (e.g. the baked bundle still carries prod, or the click-through ejects) with a documented
root-cause is a complete iter even if escapes don't drop — but the expectation is a clean lift.
