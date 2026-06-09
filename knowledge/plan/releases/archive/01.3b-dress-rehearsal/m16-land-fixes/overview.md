---
milestone: M16
slug: land-fixes
version: v1.3b "dress rehearsal"
milestone_shape: section
status: archived
created: 2026-06-08
last_updated: 2026-06-08
complexity: small-medium
delivers: rosetta-extensions/knowledge/ KB + GUIDE/README truth-up; a consolidated corpus/ops/ note on the stack-dev layout + back-compat fallback
issues: ISSUE-1, ISSUE-2, ISSUE-3, ISSUE-4, ISSUE-5, ISSUE-7 (push)
---

# M16 — Land the field fixes + restore doc truth

## Goal
Make the two already-applied fixes **durable and public**, finish the `anthropos-dev → stack-dev` rename as the
*documented default*, and clear the stale tooling docs — so the repo tells the truth before more work lands on it.

## Why section
Every deliverable is a known, enumerable edit (a push, a rename migration, a prose sweep, four header-fact fixes,
one invocation fix). No emergent path.

## Repo split
- **`rosetta-extensions`** (authoring copy → tag `dress-rehearsal-m16` → consume per-stack): the push, the stack-core
  rename migration, the prose sweep, the GUIDE/README truth-up, the repo's `knowledge/` KB.
- **`rosetta`** (this corpus): a consolidated `corpus/ops/` note on the `stack-dev` layout + back-compat; sweep any
  residual `anthropos-dev` in `corpus/`.

## Scope
- **In:**
  - **Push the stranded fixes** — the authoring copy is **2 commits ahead of `origin`** (`547de17` devpath +
    `ed72e94` migrate-race), on the **local-only** tag `stack-party-devpath-fix`. Push to `origin`, re-tag as
    `dress-rehearsal-m16` (retire the local-only tag), re-consume per-stack. (ISSUE-1①, ISSUE-7 push.)
  - **Stack-core rename migration** — make `stack-dev` the documented default and demote `anthropos-dev` to a single
    intentional "legacy alias" mention everywhere (the back-compat fallback in the scripts stays as the one alias). (ISSUE-1②.)
  - **Prose sweep** — `demo-stack/README.md:12`, `demo-stack/GUIDE.md:17`, `dev-stack/README.md:73`,
    `stack-core/gen_override.py:4` docstring (`anthropos-dev/` → `stack-dev/`). (ISSUE-2.)
  - **GUIDE.md header truth** (`demo-stack/GUIDE.md:1,5`) — remote **exists**; **13** unit tests not 78; `/stack-list`
    not `/demo-status`; status **v1.3** not "v1.1 'show floor' / M3". (ISSUE-3.)
  - **pytest doc fix** (`demo-stack/GUIDE.md:167`) — recommend `pytest tests/ -v` (the working entrypoint) + a
    3.11/3.12 note (python3.14 has no pytest module). (ISSUE-4.)
  - Refresh the `rosetta-extensions/knowledge/` KB where it repeats any of these; note the expected consumption
    version-jump (v1.2.1 → the fix tag) so it's not a surprise. (ISSUE-5.)
- **Out:** the race/idempotency *behavior* work (M17 — the fixes themselves are already applied; M16 only publishes +
  documents them). Anything frontend/verify/set-dress.

## Depends on
None. **Parallel with:** none (the honest baseline every other milestone builds on).

## Open questions (resolve during build)
- The re-tag/version scheme — lean: per-milestone `dress-rehearsal-mNN` tags + a final `v1.3.1` release tag at close
  (matches the established convention).
- Keep the `anthropos-dev` back-compat fallback in the scripts forever vs sunset it — lean: keep (a one-line
  fallback that costs nothing and protects older layouts).

## KB dependencies (read as contract)
- `corpus/ops/rosetta_demo.md` (the stack layout + the demo lifecycle)
- the `rosetta-extensions` GUIDE/README set + its `knowledge/` KB (the docs being corrected)

## Delivers
- **→ rosetta-extensions:** `knowledge/` KB + GUIDE/README truth-up; the pushed + re-tagged fixes.
- **→ rosetta:** a consolidated `corpus/ops/` note on the `stack-dev` layout + the back-compat fallback (cross-linked from `rosetta_demo.md`).

## Risk
**(blast-radius, nice-to-resolve)** the push is v1.3b's first outward-facing action — the local fixes go public.
Mitigate: re-tag cleanly (retire the local-only `stack-party-devpath-fix`), push, re-consume per-stack, verify.
