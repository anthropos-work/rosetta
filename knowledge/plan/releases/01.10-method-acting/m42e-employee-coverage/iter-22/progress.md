**Type:** tik (P8 acceptance, under TOK-10)

# iter-22 — P8 fresh-demo-up acceptance + the stale-CLI reproducibility fix

## Phase A/B — fresh demo-up + diagnose
Bumped the consumed clone `stack-demo/rosetta-extensions` to `method-acting-m42e-iter21`, then ran a fresh
`demo-down 3 --purge` + `demo-up 3` (DEMO_STORIES on). The bring-up reproduced ZERO-MANUAL: Sentinel reloaded
via `AuthorizationService/Reload` (the P5 FATAL fix — `✓ Sentinel policy reloaded … sim /start renders`),
taxonomy + directus replayed + the per-stack Directus served (`directus_collections=14`), stories seeded,
autoverify GREEN (`casbin_rules=750`). **One reproduction GAP**: the sim-embeddings replay (P6) was SKIPPED —
`stacksnap: unknown surface "sim-embeddings" (known: [directus reference-toy taxonomy])` → `cms.similarities=0`
→ `/library/ai-simulations` empty.

## Phase B — triage (the root cause)
`dev-setdress.sh`'s `build_cli` had `[ -x "$BIN_DIR/$name" ] && return 0` (skip-if-present). The stack's
shared bin dir (`$STACK/bin`, handed via `DEV_SETDRESS_BIN`) is SIBLING to the data dir purged by
`demo-down --purge`, so a STALE `stacksnap` (Jun 24, built before the P6 sim-embeddings surface existed)
SURVIVES the purge and was REUSED on the fresh demo-up → the new surface was invisible. The fresh source
(iter-21) registers the surface; only the binary was stale.

## Phase C — fix (rext)
`build_cli` now ALWAYS rebuilds every CLI from the consumed source (the same ALWAYS-rebuild discipline
up-injected.sh already applies to `stackseed`). Test harnesses that drop argv-recording stubs into
`DEV_SETDRESS_BIN` opt out via `DEV_SETDRESS_USE_STUB_BINS=1` (`run_sd` sets it). 26 DevSetdress tests pass.

## Phase D — re-measure (the authoritative re-run)
Bumped the consumed clone to `iter22`, then a SECOND `demo-down --purge` + fresh `demo-up` (the stale
stacksnap still present — the exact scenario). The iter-22 `build_cli` REBUILT `stacksnap` (binary now
Jun 25 16:06) and **sim-embeddings replayed ZERO-MANUAL**: `replayed "sim-embeddings" into demo-3: 4 table(s)
… 1490 row(s) loaded, reindexed [cms.similarities.small_embedding3]` → `cms.similarities=274` (all public
sims) → `/library/ai-simulations` populated. The fresh roster carries `picture=data-uri` + `org_logo` for
maya-thriving (member) + dan-manager (admin) — the menu identity sources are reproducible.

## Close — 2026-06-25

**Outcome:** the fresh demo-up reproduces zero-manual after the stale-CLI fix — sim-embeddings replays
automatically (cms.similarities=274), Sentinel reloads, content serves, stories seed. The P6 library is now
reproducibly populated on a fresh box.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the gate is graded by the employee semantic sweep — see iter-23, where it flips MET after
the avatar-selector harness fix)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue (into iter-23)
**Decisions:** the build_cli ALWAYS-rebuild + DEV_SETDRESS_USE_STUB_BINS opt-out.
**Side-deliverables:** none.
**Routes carried forward:** the avatar-consistency persona FAIL surfaced by the sweep → iter-23 (a harness
selector bug, not a data gap).
**Lessons:** a stack's `$STACK/bin` survives `demo-down --purge` (sibling to the purged data dir) — any
build helper with a skip-if-present must ALWAYS rebuild from the consumed tag, or a capture-path/surface
change silently no-ops on a fresh demo-up at a new tag (the M40 stale-cache trap's binary analogue).
