---
milestone: M32
slug: studio-desk-singleport
version: v1.7 "house lights"
milestone_shape: section
status: archived
created: 2026-06-15
last_updated: 2026-06-15
complexity: small
delivers: rosetta-extensions/stack-injection (the NODE_ENV=production override fix + regression test; ext tag house-lights-m32) + the :9100 doc/CORS sweep
backlog_refs: (none — sibling demo-UI rough edge surfaced by the v1.6 field-bake + the M31 investigation)
---

# M32 — studio-desk single-port / production alignment + the `:9100` sweep

## Goal
A fresh browser at demo-N's studio-desk (e.g. `http://localhost:39000`) lands on a live page instead of a 302 to the
dead `:9100`, by running the container's production code path; and the docs/CORS all agree on single-port `9000`+offset.

## Why section
Verified 1-line root-cause fix with a clear mechanism + low risk; the rest is a doc/CORS sweep + a regression test + a
Playwright smoke. Build with `/developer-kit:build-milestone`.

## Repo split
- **`rosetta-extensions`** (authoring → tag `house-lights-m32` → consume): the `NODE_ENV=production` override in
  `gen_injected_override.py` FRONTENDS + the `test_injection.py` regression assertion + the `:9100` CORS removal.
- **`rosetta`**: the `:9100` doc sweep (demo-up SKILL, `frontend-tier.md`).

## Scope
- **In:** `gen_injected_override.py` FRONTENDS studio-desk dict (~lines 90-96) — add `NODE_ENV=production` (+
  `FRONTEND_PORT=9000` belt-and-suspenders). **Root cause:** the base compose ships `NODE_ENV=development` and the
  override's per-frontend env block is **additive** (deliberately not `!override`), so `development` survives →
  `src/index.ts` `isProduction=false` → the dev block `res.redirect('http://localhost:9100/home')` fires (dead port).
  Production → `sendFile`, no cross-port redirect.
- **In:** a regression assertion in `stack-injection/tests/test_injection.py` near the single-port tests (~820-857) —
  assert `NODE_ENV=production` in the studio-desk env block.
- **In:** a Playwright smoke on `/home` + a couple of studio-desk routes confirming the production `sendFile` path serves
  them (verify the dev block's `.html`-extension redirects aren't load-bearing).
- **In:** the `:9100` sweep — demo-up SKILL (`:9100+`→`:9000+`), `frontend-tier.md:21` (drop the dead `:9100` frontend
  port → single-port `9000`+offset), `gen_injected_override.py:249` CORS (remove the un-offset `9100` origin — explicit
  decision note; dead now that studio-desk is single-port production).
- **Out:** the cert (M31); ant-academy (backlog).
**Depends on:** none functionally; sequence after M31 (shared `up-injected.sh`/doc-cluster surface).
**Parallel with:** none.
**Estimated complexity:** small.
**Open questions:** confirm the production `sendFile` path covers ALL routes the dev block handled (the Playwright smoke proves it).
**KB dependencies:** `corpus/ops/demo/frontend-tier.md` (the studio-desk port story); demo-up SKILL.
**Delivers →** `rosetta-extensions/stack-injection` (the override fix + test; ext tag `house-lights-m32`) + the `:9100` doc/CORS sweep.
