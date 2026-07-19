# M230 — Spec notes

_(technical details / thresholds accumulate here during build)_

## Pre-flight audits — iter-01

**KB-fidelity (Phase 0b, 2026-07-19): GREEN.** Report: `kb-fidelity-audit.md`. All load-bearing topics
PAIRED + ALIGNED + code-verified; no blind areas, no stale load-bearing claims. The frontend-tier.md F4
correction is the milestone's own `Delivers:` target, not a blind area.

## Topic → doc → code triples (verified iter-01)

- **Content model (empty-grid root cause):** `corpus/services/ant-academy.md` § The Content Model →
  `stack-demo/ant-academy/code/src/lib/{serverTenant,backendContent,draftMode,draftCatalog}.js`.
  The seam: `serverTenant.js::getServerCatalogView()` = `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView(); return draftsEnabled() ? mergeDrafts(view, eids) : view`.
- **Option C vehicle (demo-patch):** `corpus/ops/demo/demopatch-spec.md` →
  `.agentspace/rosetta-extensions/demo-stack/patches/{demopatch, manifest_loader.py}`. Academy-patch
  precedent already present: `patches/ant-academy-dev-origins/`.
- **DELIVERS doc:** `corpus/ops/demo/frontend-tier.md` § ant-academy →
  `.agentspace/rosetta-extensions/demo-stack/ant-academy.sh` (confirmed: sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` 0 times).
- **Gate measurement:** `corpus/ops/demo/coverage-protocol.md` (ANT_ACADEMY descriptor) →
  `.agentspace/rosetta-extensions/stack-verify/e2e/lib/coverage-manifest.ts` + `coverage.spec.ts`.

## Infra feasibility (iter-01 probe)

Docker up (28.5.1, 0 containers), 205Gi free, no ENOSPC. `stack-demo/` fully cloned (ant-academy/code + all
platform repos + rext consumption clone). rext authoring copy present. `~/.pgpass` present (prod DB read
plausibly available → Option B not a priori blocked, but heavier). **demo-1 injected images built 41h ago +
`demo-stack/stacks/demo-1` artifacts present → a cold /demo-up is FEASIBLE on this box** (not a hard blocker).
rext pinned at `casting-call-m228-hiring-scope-fix`.

## The Option C mechanism (iter-02) — shipped artifacts + thresholds

- **Manifest:** `demo-stack/patches/academy-fs-published-fallback/academy-fs-published-fallback.yaml`.
  `repo: ant-academy`, `path: code/src/lib/serverTenant.js`, `scope: demo`.
  `pre_sha256=43977541fcf35352431418930d884c7188be796a5e6cdd95a9a45f7371fb9665`,
  `post_sha256=e0f48e816cb84c5ea372c425645cfad4f36c29224446ef417e712ac55f489dac`,
  `post_marker=ACADEMY_DEMO_FS_PUBLISHED`. Anchor = the `getServerCatalogView` `?? emptyCatalogView()` line
  (occurs 1×; the `getPublicCatalogView` twin reads `new Set()` and is a DISTINCT substring). Replacement =
  env-gated (`ACADEMY_DEMO_FS_PUBLISHED === '1'`) `mergeDrafts(emptyCatalogView(), eids)` with `_draft`/`_origin`
  stripped from chapters/series/skillPaths; else the pristine `emptyCatalogView()` (behavior-identical when unset).
- **Helper:** `stack-injection/apply-academy-fs-published.sh apply|revert` (native, mirrors
  `apply-ant-academy-dev-origins.sh`). Drift-refuse + single-occurrence + idempotent + post-condition re-check.
- **Launcher wiring:** `ant-academy.sh` — apply-before-launch (default-on; `DEMO_NO_ACADEMY_FILL=1` opts out;
  non-fatal) + `ACADEMY_DEMO_FS_PUBLISHED=1` in the launch env (only when the patch applied) + revert-on-`--stop`.
- **Tests:** `demo-stack/tests/test_academy_fs_published.py` — 14, all green.
- **rext tag:** `playbill-m230-academy-fs-published` (authoring-copy commit `76ee1a0`).

## The chip is driven ONLY by `_draft` (code-verified)
`SkillPathCard.jsx:69` `isPathDraft = skillPath?._draft === true` → `.draft-ribbon`; `AcademyClient.jsx:1008`
`isChapterDraft = chapter._draft === true` → the draft badge. So stripping `_draft` (+ `_origin`) removes the
chip entirely — the no-chip property the gate requires.

## Runtime proof (iter-02 Phase D, standalone)
Patched `next dev` on `:3099` (persona bypass, `ACADEMY_DEMO_FS_PUBLISHED=1`, no real Clerk keys) + SSR-curl of
`/` as `e2e_persona=member`: **59 skill-path cards** (`path-face-closed`), 440 `chapter-card`, real names present,
**0 `draft-ribbon` / 0 `data-draft="true"`**, no empty-state; clone reverted byte-clean. → the patch renders real
cards via the real `getServerCatalogView` code path with no chip.

## Observed obstacle for the FORMAL gate (cold /demo-up)
The local `stack-demo/next-web-app` `urls.ts` has DRIFTED from the `next-web-public-website-url` +
`next-web-studio-url` manifests' pinned `pre_sha256` (2 `test_demopatch.py` failures; `live != pinned pre`). A
cold `/demo-up` would drift-refuse those next-web patches → a clone re-sync/re-pin is a prerequisite. Unrelated to
the academy patch (which pins the current academy clone, verified pristine).
