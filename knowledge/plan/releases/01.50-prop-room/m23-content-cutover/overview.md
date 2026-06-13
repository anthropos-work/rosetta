---
milestone: M23
slug: content-cutover
version: v1.5 "prop room"
milestone_shape: section
status: planned
created: 2026-06-11
last_updated: 2026-06-13
complexity: medium
delivers: updated safety.md (self-contained-content deltas) + corpus/services/* (env truth) + snapshot-spec.md (cross-surface closure gene) + the env re-point (demo + dev) + closure capture + the closure gene + the directus_files ref capture (Fate-3 from M21)
backlog_refs: NEW-3 (referential-consistency interim options / the close)
---

# M23 — Content cutover + referential closure

## Goal
Point the stack's services at their **own** Directus and guarantee the served catalog is **referentially closed** —
so the content a stack serves never references a taxonomy node-id its captured subset lacks (the empty
Assign-AI-Simulation-picker class disappears).

## Why section
Once M22 boots a verified local Directus, the cutover is concrete: the demo env mechanism exists, the dev env-emission
is one well-scoped addition, the closure problem has two known options. No exploratory uncertainty — it is
data-design care, not path-finding. Build with `/developer-kit:build-milestone`.

## Repo split
- **`rosetta-extensions`** (authoring → tag `prop-room-m23` → consume): the env re-point (demo + dev), the closure
  capture, the cross-surface closure gene, the dev-side prod-token strip.
- **`rosetta`**: `corpus/services/{cms,studio-desk,jobsimulation,next-web-app}.md` (env/dependency truth),
  `corpus/ops/safety.md` (retire the live-prod-read notes), `snapshot-spec.md` (the closure gene).

## Scope
- **In (`rosetta-extensions`):**
  - **Re-point `DIRECTUS_BASE_ADDR`** to the local instance — demo via the ready `gen_injected_override.py` env
    mechanism (one-line per service); **dev via growing `stack-core/gen_override.py` to emit `environment:` blocks**
    (it can only emit ports today — the single genuinely-new bit of plumbing, kept minimal).
  - **Keep the asset plane on prod** — `DIRECTUS_PUBLIC_BASE_ADDR` stays `content.anthropos.work` so browser images
    stay **real** (the user's explicit call), sidestepping the baked next/image host whitelist with no UI rebuild.
  - **studio-desk** gets the local instance (`DIRECTUS_BASE_URL`) + a **locally-minted admin token** so its
    skill-path writes target the per-stack Directus (never prod).
  - **Extend the prod-token strip to opted-in dev stacks** (today demo-only).
  - **Referential closure** — make the taxonomy capture include every node-id the captured content references
    (closure-at-capture; **full-taxonomy capture** as the simple fallback the corpus already names) + a
    **cross-surface fidelity gene** so closure is *measured*, not assumed.
  - **Wire the `directus_files` ref capture** (Fate-3 from M21 — see `m21-structure-capture/audit-deferrals/deferral-audit-2026-06-13-m21-close.md`):
    the dead `media.go` filter/columns need a `directus_files` TableSpec in `directus.Surface()` so captured content
    rows resolve their image-asset UUIDs to the prod-public `<DIRECTUS_PUBLIC_BASE_ADDR>/assets/<uuid>` URLs the asset
    plane serves. Orthogonal to M21's structure-serve gate (which fired without it); it belongs with this milestone's
    asset-plane work (refs only — blob BYTES stay backlog, DEF-M10-01).
- **In (`rosetta`):** update `corpus/services/{cms,studio-desk,jobsimulation,next-web-app}.md` (the env/dependency
  truth — jobsimulation reads Directus via cms RPC, not directly; next-web via cms/router only) + `corpus/ops/safety.md`
  (retire the live-prod-read notes; the token-strip stays as the write-disarm).
- **Out:** M21/M22 mechanics; the cms `PostMultipart` hardcoded-prod-upload-URL **platform bug** (can't fix without
  a platform edit — disarmed by the token strip + documented as a user-owned upstream PR + a verify warning).

## Depends on / parallel with
Depends on: M22 (a booted, verified local Directus to point at). Parallel with: none.

## Open questions
- Closure-at-capture (compute the referenced node-id set, capture exactly those) vs **full-taxonomy capture** (no
  subset — removes the dangling problem wholesale, simpler, slightly bigger snapshot) (lean: full-taxonomy capture
  for simplicity per the maintainability constraint, unless the size is prohibitive).
- Whether N=0 dev re-point gets any tooling or stays a pure documented manual recipe (lean: documented recipe only —
  honor the n=0 guard + zero-platform-edit line).

## KB dependencies
`corpus/ops/snapshot-spec.md` (the referential-closure boundary + the fidelity genes), `corpus/ops/safety.md`,
`corpus/services/*` (the Directus consumers), `corpus/ops/seeding-spec.md` (the content↔seed linkage).

## Risk (correctness — degrades-quality)
A wrong re-point silently sends a stack back to prod (or to a dead local instance). Mitigate: the M22 no-prod-read
verify assert + the `EnvContract` gate catch both; non-fatal degrade keeps a good stack up.
