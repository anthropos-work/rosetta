# iter-02 — progress

**Type:** tik under TOK-01 (Option C).

## What landed (Phase C — the fix mechanism, in the rext authoring copy)
1. **Manifest** `demo-stack/patches/academy-fs-published-fallback/academy-fs-published-fallback.yaml` —
   content-anchored on the `getServerCatalogView` `?? emptyCatalogView()` fallback (the getPublicCatalogView twin
   is a distinct `new Set()` line). `pre_sha256=43977541…`, `post_sha256=e0f48e81…` (computed via manifest_loader
   against the real clone; validated: pristine→post→pristine round-trips, marker present only in patched).
2. **Native apply/revert helper** `stack-injection/apply-academy-fs-published.sh` — drift-refuse,
   single-occurrence, idempotent, post-condition re-check; apply-before-launch / revert-on-stop (mirrors
   `apply-ant-academy-dev-origins.sh`, because ant-academy runs natively via `next dev`).
3. **`ant-academy.sh` wiring** (4 edits) — `FSPUB_PATCH` def, apply-before-launch (default-on,
   `DEMO_NO_ACADEMY_FILL=1` opt-out, non-fatal), `ACADEMY_DEMO_FS_PUBLISHED=1` in the launch env, revert-on-`--stop`.
   `bash -n` clean; 59 existing `test_ant_academy.py` tests still pass (no regression).
4. **Unit test** `demo-stack/tests/test_academy_fs_published.py` — 14 tests (manifest schema/round-trip, apply→
   revert→idempotent→drift-refuse ladder against a temp copy of the real clone, the strip transform's no-chip
   property, launcher-wiring fence). All green, warning-clean.
5. rext committed `76ee1a0` + tagged **`playbill-m230-academy-fs-published`**.

## Phase D — runtime proof (standalone, bounded)
Applied the patch to the real clone → `next dev` on `:3099` (persona bypass, `ACADEMY_DEMO_FS_PUBLISHED=1`, no
real Clerk keys needed) → SSR-curled the home as a signed-in `member` persona → **measured**:

| Signal | Value |
|---|---|
| skill-path cards (`path-face-closed`) | **59** |
| `chapter-card` markers | 440 |
| `path-trigger` | 118 |
| **`draft-ribbon` (must be 0)** | **0** |
| **`data-draft="true"` (must be 0)** | **0** |
| real catalog names (Claude Code / Foundations / Agent / AI Engineering / Business) | all present |
| empty-state markers | none |

Clone **reverted byte-clean** afterward (verified). → the patch renders **real cards through the real
`getServerCatalogView` code path, production-faithful, NO chip** — the gate's substance, proven.

## Side-observation (unrelated, pre-existing)
`test_demopatch.py` has **2 failures** — the `next-web-public-website-url` + `next-web-studio-url` manifests
validated against the LIVE `stack-demo/next-web-app` clone, whose `urls.ts` has DRIFTED from their pinned
`pre_sha256` (`live != pinned` for both). **Not my regression** (my work touched only ant-academy). Signals the
local next-web clone is at a different ref than its manifests pin — relevant to a **cold `/demo-up`** (next-web
patches would drift-refuse) but NOT to the academy patch (which pins the current academy clone). Routed forward.

## Close — 2026-07-19

**Outcome:** Option C fill mechanism BUILT + unit-verified (14 tests) + **runtime-proven** (59 real cards, 0 draft
chips, clone byte-clean). rext tagged `playbill-m230-academy-fs-published`. DELIVERS doc (`frontend-tier.md`)
updated with the corrected F4 attribution + the shipped mechanism.
**Type:** tik
**Status:** closed-fixed  (planned scope — the fill mechanism — landed + runtime-proven)
**Gate:** NOT MET  (the FORMAL gate is the coverage sweep's `ANT_ACADEMY` rendered-card count on a **cold
`/demo-up`**; the standalone runtime proof is a strong proxy but not that specified measurement)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: **y** — (5) cap-reached: n — (6) protocol-stop: n — Outcome: **exit-4** (the formal cold-`/demo-up` gate proof is surfaced as a user-blocker: heavy infra + a concrete unrelated obstacle [next-web clone drift], per the orchestrator's "surface the cold demo-up, don't burn hours" guidance)
**Decisions:** iter-02 D1 (side-observation: next-web clone drift); milestone-root USER-BLOCKER record.
**Side-deliverables:** none (the frontend-tier.md doc update is the milestone's own DELIVERS, not a side-fix).
**Routes carried forward:**
- **The formal gate** — coverage sweep `ANT_ACADEMY` rendered-card count on a cold `/demo-up` — the next build-iter call's focus (needs the user go-ahead for the heavy bring-up).
- **next-web clone re-anchor** (the 2 drifted demopatch manifests) — a demo-up prerequisite; route to demo-up/close.
**Lessons:** The Option C fix is small + fully verifiable off-demo (manifest+ladder unit tests + a pure-functional
transform proof + a cheap standalone SSR-curl runtime proof reaching the exact code path). The heavy, risky part
is exclusively the FORMAL gate (cold `/demo-up` + coverage sweep) — so building + runtime-proving the fix first,
then surfacing the cold-`/demo-up` as its own decision, is the right split (matches the orchestrator's guidance).
