# M249 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

## KB-fidelity (Phase 0b) — YELLOW, tracked

- **KB-1** — demopatch-spec §5 inventory (16 → 21) + `test_patch_inventory.py` fence must move together with the 5 new manifests. In scope (Docs section).
- **KB-2** — frontend-tier.md:277 "studio-desk … no source patch" becomes stale on M249 → update it (Delivers → frontend-tier.md).
- **KB-3** — cockpit-spec.md has no return-nav section → author it (Delivers → cockpit-spec.md).
- **KB-4** — demopatch-spec §4 apply-vehicles table gains a studio-desk (image-baked, `demopatch`-tool) row; §5-bis fingerprint note extends to `build_frontend_studio_desk` (net-new fingerprint).

## Pre-existing test failures surfaced (NOT M249 regressions — confirmed identical on committed rext HEAD)

Two `test_ant_academy*` classes fail on this box independent of M249 (verified by running them against `git show HEAD:demo-stack/ant-academy.sh` — same 6+1 failures). Both are in **M251's** test domain (health/launcher/inventory), not M249's (patch) domain, per the coordination guardrail — so untouched here:
- **`test_ant_academy` — `TestAntAcademyLauncher` + `TestAntAcademyPreBindReap` (6 failures):** flaky integration tests that spawn real fake HTTP servers on demo-2's offset ports (`:23077`) and assert the reap heuristic kills them; a bare `python -c http.server` fake isn't recognized as reap-ownable, and the tests leak the fake on failure (a leftover was found holding `:23077` from 3:33PM, pre-session). My diff adds **no** launcher/reap/bind logic (only a non-fatal `COCKPIT_PATCH apply/revert` — git-diff-confirmed).
- **`test_ant_academy_clerk_wiring.test_overlay_has_minted_pk_and_no_real_secret` (1 error):** `extract_overlay_block()` slices `ENVF=` → `} > "$envlocal"`, a range that spans the pre-existing `. "$HERE/detach.sh"` line — under the test's `set -u` preamble (no `HERE`) that aborts (exit 127) **before** the overlay body runs. The `NEXT_PUBLIC_COCKPIT_URL` line I added to `write_env_local` is inside that range but never reached; my own `TestAntAcademyWiring.test_cockpit_url_baked_in_env_local` covers the addition.

**Fate (close Phase 1b deferral re-audit, 2026-07-24): Fate 2 — confirmed covered by M254, no sibling edit.** These two failures are a strict subset of the **8 live/env/docker-gated demo-stack failures** M251's close already fated **Fate 2 → M254** (`m251-test-health/decisions.md` §"Surfaced but OUT of M251 scope → M254"; M254 gate parts (g) live-box test-health + (h) live-browser re-prove own them). They are NOT M249 deferrals (M249 deferred nothing) and NOT M249 regressions (verified identical on committed rext HEAD); M249 only re-observed them. No new sibling `overview.md` edit — M254 already owns the set. Not repeat, not aged-out, not chronic.

## Design decisions

- **D1 — Lane A is ONE manifest on `NavbarTop.tsx`, not two files.** The milestone spec names both `NavbarTop.tsx` + `useNavbarSections.tsx`, but demopatch is one-anchor-per-manifest and threading a hook-defined item through the return object + the render site is ≥3 anchors across 2 files. Instead, a single anchor on the desktop `!hiddenSidebar` account-dropdown return block in `NavbarTop.tsx` defines the item **inline** (spreading the in-scope `logOutMenuItem` for a valid `IconDefinition` — no new import) and renders it fail-closed (`process.env.NEXT_PUBLIC_COCKPIT_URL ? … : null`, dropped by the existing `_compact`). Behaviour-identical when the env is unset. Covers **both** apps/web + apps/hiring (shared `packages/ui`). Residual: the responsive **mobile** account menu (`hiddenSidebar` branch) is not covered — a presenter demos on desktop; documented in the manifest DISCLOSED note.
- **D2 — Lane B reads `import.meta.env.VITE_WEB_APP_URL || 'https://app.anthropos.work'`, NOT `config.WEBAPP_URL`.** `config.WEBAPP_URL` = `import.meta.env.VITE_WEB_APP_URL || 'http://localhost:3000'` — reading the env directly (a) needs no `config` import (neither `userProfile.js` nor `pageWrapper.js` imports it → would be a 2nd anchor), and (b) keeps the ORIGINAL `https://app.anthropos.work` fallback so the patch is behaviour-identical when the var is unset (the demopatch design rule). Same underlying env var `config.WEBAPP_URL` reads. Lane B = 3 manifests (`studio-desk-back-to-cockpit` = userProfile menu block: rewrite back-item + add cockpit item; `studio-desk-logout-url` = userProfile `handleLogout`; `studio-desk-logo-url` = pageWrapper logo).
- **D3 — `VITE_COCKPIT_URL` rides the `.env.production.local` overlay, not a `--build-arg`.** studio-desk's `Dockerfile.dev` declares ARGs for VITE_CLERK_PUBLISHABLE_KEY/VITE_GRAPHQL_ENDPOINT/VITE_ENVIRONMENT/VITE_WEB_APP_URL only — NOT VITE_COCKPIT_URL. An undeclared `--build-arg` is dropped by Docker. So (mirroring the existing VITE_CLERK_SIGN_IN_URL bake) it goes into the gitignored `.env.production.local` overlay that `vite build` loads into `import.meta.env`. Zero platform-repo edit.
- **D4 — Lane C is a NATIVE-RUN patch** (ant-academy runs `next dev`, not an image), so a manifest + a `stack-injection/apply-ant-academy-back-to-cockpit.sh` helper (modelled on `apply-ant-academy-dev-origins.sh`), wired into `ant-academy.sh` `reapply_clone_patches` (apply) + `--stop` (revert) + `write_env_local` (the `NEXT_PUBLIC_COCKPIT_URL` line). The item is a plain `<a href={process.env.NEXT_PUBLIC_COCKPIT_URL}>`, rendered fail-closed.
- **D5 — the studio `.env.production.local` overlay is now ALWAYS (over)written (found at the live render-confirm).** The first demo-2 studio rebuild shipped a bundle with the Back-to-Cockpit item fail-closed to nothing: a Jul-17 stranded `.env.production.local` (a crashed prior build's leftover) tripped the M214 `if [ -e "$desk_env" ]` "skip if it exists (repo-shipped?)" branch, so NEITHER `VITE_CLERK_SIGN_IN_URL` NOR `VITE_COCKPIT_URL` was baked. But studio-desk's real repo **gitignores `.env.*.local`** — it can never ship one, so a pre-existing overlay in the demo clone is ALWAYS a leftover. The skip was guarding an impossible case while creating the stranding hazard (the "silently unbaked" trap the fingerprint kills, one level up). Fix: write unconditionally (`>` truncate, the same contract as next-web's `apps/web/.env.local`); the RETURN trap removes it. Regression-fenced (`test_studio_desk_overwrites_a_stranded_overlay`); the obsolete non-clobber test removed. rext `bcbb779`. This also fixes the pre-existing `VITE_CLERK_SIGN_IN_URL` stranding.

## Adversarial review (close Phase 2c)

- **Scenario — the load-bearing inventory number silently drifts from the code.** This milestone's central
  doc claim is a COUNT: `demopatch-spec.md` §5 asserts "21 patches: 11 next-web-app · 2 app · 5 ant-academy ·
  3 studio-desk", and §4/§5-bis/frontend-tier/studio-desk all lean on it. A docs milestone's worst failure is a
  count that reads true but no longer matches the `patches/` directory it describes — a reader trusts a wrong
  number, and nothing catches it. **Verified at close:** enumerated the real rext `demo-stack/patches/`
  directory at the code-of-record tag (`july-jitter-m249-harden` @ 8ab5192) — 21 manifest dirs (the 22nd is
  `__pycache__`), `repo:` breakdown `11 next-web-app · 2 app · 5 ant-academy · 3 studio-desk`, EXACTLY the doc's
  numbers. The drift is additionally **fenced in code** — `test_patch_inventory.py::TestPatchInventory` pins the
  total (21) AND the per-repo breakdown against the §5 table, so the two can only move together. No fix needed;
  the claim is true and self-defending.

## Live render-confirm (demo-2, LOCAL) — Phase 6

Rebuilt the 3 demo-2 frontend images from the M249 authoring `up-injected.sh` (patches applied via the ladder; the **patch-set fingerprint forced** the next-web/hiring rebuild — distinct M249 labels `02cdab…`/`fec26d…`; studio force-removed because the overlay fix is in `up-injected.sh`, not the fingerprinted manifest set), recreated all 3 compose containers (`--no-deps`) on the fresh images, and confirmed:
- **next-web / hiring / studio bundles** all carry `Back to Cockpit` + the cockpit URL `:27700`. Containers Up + serving (web 307 / hiring 307 / studio 302 — auth redirects, expected).
- **studio prod-eject**: the offset app `localhost:23000` is baked (the runtime-winning value); the residual `app.anthropos.work` strings are the inert `|| 'https://app.anthropos.work'` fallback (never executes when `VITE_WEB_APP_URL` is set — the demopatch behaviour-identical-when-unset rule), so **0 effective prod-ejects**.
- **ant-academy** (native `next dev`, not a container): the patch applies to the real demo-2 clone → `UserMenu.jsx` carries the item reading `process.env.NEXT_PUBLIC_COCKPIT_URL`; the env bake value is `NEXT_PUBLIC_COCKPIT_URL=http://localhost:27700` (verified at source+env level, reverted clean — a full native browser render was out of reach because the demo-2 academy uses the pinned CONSUMPTION clone's launcher + a test orphan squatted `:23077`).
