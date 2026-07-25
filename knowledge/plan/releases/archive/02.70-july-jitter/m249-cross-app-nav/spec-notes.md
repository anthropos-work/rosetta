# M249 — Spec notes

_Working notes for the cross-app navigation build — filled in as each lane lands._

## Pre-flight audits — Section 1 (next-web-back-to-cockpit)
- KB-fidelity (Phase 0b): **YELLOW** — report `kb-fidelity-audit.md`. Both blind areas (additive-UI injection pattern doc + cockpit return-nav section) are planned Delivers; no stale contract the impl reads. Findings KB-1..KB-4 in `decisions.md`.

## Topic → doc → code triples
- demopatch mechanism → `corpus/ops/demo/demopatch-spec.md` → `rext demo-stack/patches/{demopatch,manifest_loader.py}` + `stack-injection/apply-*.sh`
- frontend image build → `corpus/ops/demo/frontend-tier.md` → `rext demo-stack/up-injected.sh` (`build_frontend_next_web|studio_desk|hiring`), `ant-academy.sh`
- presenter cockpit (return-nav target, 7700+offset) → `corpus/ops/demo/cockpit-spec.md` → `rext demo-stack/cockpit.py`
- studio-desk service → `corpus/services/studio-desk.md` → `stack-demo/studio-desk/app/core/scaffold/{userProfile.js,pageWrapper.js}` + `app/services/config.ts`

## Confirmed source facts (verified against stack-demo @ M246-bump)
- **next-web** `packages/ui/src/NavBar/NavbarTop.tsx` — desktop account dropdown = the `if (!hiddenSidebar) { return _compact([mapItem(settingsMenuItem,0), …, mapItem(logOutMenuItem,0)]); }` block (occurs once). `MenuItemProps` + `MenuType` already imported. Outbound items render via `navbarMenuItems`→`MenuWithLink`→`<Link href={item.key}>`.
- **studio-desk** `config.ts`: `WEBAPP_URL = import.meta.env.VITE_WEB_APP_URL || 'http://localhost:3000'`. `Dockerfile.dev` declares ARG VITE_WEB_APP_URL (baked at offset, up-injected.sh:837) — NOT VITE_COCKPIT_URL. `userProfile.js` back-item L148 + logout L302 hardcode `https://app.anthropos.work`; `pageWrapper.js` logo L149 hardcodes it. Neither imports `config`.
- **ant-academy** `code/src/components/UserMenu.jsx` — logout row = `<div className="user-menu-divider" />` + `<div className="user-menu-row user-menu-row--logout">` (the 2-line pair occurs once). Native `next dev`; `NEXT_PUBLIC_COCKPIT_URL` inlined from `code/.env.local` at compile.
- **workspace vars**: `DEMO_WS=$REPO_ROOT/stack-demo`, `DEMO=$DEMO_WS`. Cockpit port = `7700+OFFSET`. `SCHEME=https` on `--public-host`, else `http`.

## Lane A — `next-web-back-to-cockpit` demopatch (web + hiring, shared `packages/ui/NavbarTop`)
_(landed — see decisions.md D1.)_

## Lane B — `studio-desk-back-to-cockpit` demopatch (+ `pageWrapper.js:149` / `userProfile.js:147,302` prod-eject fix)
_(landed — see decisions.md D2/D3. 3 manifests.)_

## Lane C — `ant-academy-back-to-cockpit` demopatch (`UserMenu.jsx:143`)
_(landed — see decisions.md D4. native-run.)_

## Wiring — offset-URL bake + apply/revert into `up-injected.sh` (net-new `build_frontend_studio_desk`) + `ant-academy.sh`, fail-closed
_(landed.)_

## Docs — additive-UI injection pattern + cockpit-spec return-nav section
_(landed.)_

## Lane A — `next-web-back-to-cockpit` demopatch (web + hiring, shared `packages/ui/NavbarTop`)
_(TBD during build.)_

## Lane B — `studio-desk-back-to-cockpit` demopatch (+ `pageWrapper.js:149` / `userProfile.js:147,302` prod-eject fix)
_(TBD during build.)_

## Lane C — `ant-academy-back-to-cockpit` demopatch (`UserMenu.jsx:143`)
_(TBD during build.)_

## Wiring — offset-URL bake + apply/revert into `up-injected.sh` (net-new `build_frontend_studio_desk`) + `ant-academy.sh`, fail-closed
_(TBD during build.)_

## Docs — additive-UI injection pattern + cockpit-spec return-nav section
_(TBD during build.)_
