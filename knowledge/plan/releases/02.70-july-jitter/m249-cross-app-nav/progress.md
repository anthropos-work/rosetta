# M249 — Progress

## Sections
- [x] `next-web-back-to-cockpit` demopatch — shared `packages/ui/NavbarTop` → covers web + hiring
- [x] `studio-desk-back-to-cockpit` demopatch — + `pageWrapper.js:149` logo / `userProfile.js:147,302` back+logout prod-eject fixes (same scaffold) — 3 manifests (chain: back-to-cockpit → logout-url; + logo-url)
- [x] `ant-academy-back-to-cockpit` demopatch — `UserMenu.jsx` (native-run) + `apply-ant-academy-back-to-cockpit.sh` helper
- [x] `up-injected.sh` wiring — offset-URL bake + apply/revert (both next-web overlays + net-new `build_frontend_studio_desk` ladder + patch-set fingerprint) + `ant-academy.sh`, fail-closed when the env is unset
- [x] Docs — additive-UI injection pattern doc (demopatch-spec §8) + cockpit-spec return-nav section + demopatch-spec §5/§4 inventory (16→21) + frontend-tier + studio-desk.md

## Completeness Ledger

### Deferred
- _(none)_

### Dropped
- _(none)_

## Notes
- 5 new manifests (inventory 16→21: next-web-app 10→11, ant-academy 4→5, studio-desk 0→3) + 1 native apply helper.
- `test_back_to_cockpit_m249.py`: shape + fail-closed + LIVE apply/revert (tool + native) + wiring (next-web/studio/academy). M249-relevant suite GREEN (183 in the final run).
- **Live render-confirm (demo-2, LOCAL): DONE + GREEN** — all 4 app menus carry Back-to-Cockpit @ `:27700` (web/hiring/studio bundle-verified in the fresh images + running containers; academy source+env verified on the real clone); studio prod-eject fixed (offset app `:23000` baked, 0 effective ejects). See decisions.md § Live render-confirm.
- **D5 (found at the live-confirm):** studio `.env.production.local` overlay now always-overwritten — a stranded leftover was silently defeating the bake. rext `bcbb779`; regression-fenced.
- 2 pre-existing `test_ant_academy*` failures (launcher/reap flakiness + clerk-wiring extraction bug) — confirmed identical on committed HEAD, M251 domain, NOT M249 regressions (see decisions.md).
- rext consumption tag: `july-jitter-m249-cross-app-nav` → moved to the final code-of-record `bcbb779` (pushed + verified on origin).
