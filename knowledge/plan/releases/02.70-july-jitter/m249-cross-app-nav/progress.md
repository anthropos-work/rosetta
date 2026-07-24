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

## M249: Hardening

### Pass 1 — 2026-07-24
No coverage tool for the demo-stack python (unittest, no coverage plugin in this tooling); coverage used as a
manual finder over the 4 harden focus areas — (a) manifest round-trip / apply-revert cleanliness, (b)
fail-closed-when-env-unset, (c) patch-set fingerprint forces a rebuild, (d) the D5 overlay-overwrite fix.
Findings: (c) and (d) are already **deeply + dynamically** covered (`test_patchset_fingerprint.py` proves add /
re-pin / missing / DEMO_NO_PATCH each move the fp AND that every applied manifest — incl. the net-new studio-desk
trio — is in some fp call; `test_frontend_build.py::test_studio_desk_overwrites_a_stranded_overlay` + the
failed-build trap test fence D5). (a)/(b) had **3 genuine gaps**, all closed as pure-unit, clone-independent
tests (no demo/docker needed):

**Tests added (7, all unit):**
- `test_back_to_cockpit_m249.py::TestManifestShape.test_studio_chain_is_sha_linked` — the studio
  back-to-cockpit → logout-url CHAIN: asserts `logout.pre_sha256 == back.post_sha256` (same file). A half re-pin
  was caught only at live apply-time (which SKIPS without a clone); now a loud unit failure.
- `test_back_to_cockpit_m249.py::TestFailClosedMechanism` (2 tests) — the DEEP fail-closed fence. The pre-existing
  `test_fail_closed_env_gate` only checked the env NAME appears; these pin the actual guard (conditional render →
  null/'' for additive items; `env || <original-prod-url>` fallback for prod-eject rewrites), whitespace-normalized.
  Catches a bare unguarded read (dead item off-demo) and a fallback-stripped rewrite (href=""/prod re-eject off-demo).
- `test_back_to_cockpit_m249.py::TestNativeHelperApplyRevert` (4 tests) — the ant-academy native helper
  HAND-RE-IMPLEMENTS the demopatch guard ladder; only its exit-0 happy path was tested (and it SKIPS without the
  real clone). Added clone-independent REFUSE/error-path tests against synthetic targets: apply-drift-refuse
  (exit 2, file untouched), revert-drift-refuse (exit 7, file untouched), unknown-verb (exit 1), missing-target
  (exit 1). The load-bearing invariant — a REFUSE never writes — is asserted byte-for-byte.

**Teeth verified:** each static assertion mutation-proofed (broken chain / stripped fallback / dropped ternary /
bare `${env}` rewrite each fail the check); the helper error-path tests bite by construction (drive the real
helper against drifted/missing inputs → exact REFUSE exit codes).

**Bugs fixed inline:** none — the M249 code was already correct + LIVE-verified GREEN on demo-2; this pass is
pure test-depth (the docstring-vs-assertion gap in fail-closed, the unit gap in chain-linkage, the untested
REFUSE ladder in the native helper).

**Flakes stabilized:** none — 3 consecutive clean sequential runs (32 tests each); deterministic.

**Coverage delta:** `test_back_to_cockpit_m249.py` 25 → 32 tests (+7); full M249-touched suite 131 → 138 GREEN.

**Knowledge backfill:** none warranted — all 3 deepenings fence invariants ALREADY documented in
`demopatch-spec.md` (the chain `pre==post` rule; the behaviour-identical-when-unset rule; the 7-guard ladder the
native helper re-implements) + `decisions.md` D1/D2/D4. No new system truths surfaced.

### Stop condition
Stopped after Pass 1 — the full 6-dimension scan over the 4 focus areas found the 3 gaps above and nothing else;
(c)/(d) already deeply covered; 0 bugs; 0 flakes. A 2nd pass would only pad. Right-sized per the harden brief
(deepen genuine gaps; do NOT force 5 passes on live-verified code).

rext code-of-record: **`july-jitter-m249-harden`** (test-only commit atop `bcbb779`; pushed + verified on origin).

## M249: Final Review

Close review found a near-clean milestone (docs deliverable; code-of-record hardened + LIVE-verified in rext).
Full accounting:

### Scope
- [x] All 5 sections checked; `overview.md` Open questions all resolved (rewrite+add per D2/D5; no
  `DEMO_NO_BACK_TO_COCKPIT` knob — the fail-closed conditional render IS the opt-out; demo-path only). 0 gaps.

### Code Quality
- [x] [verify] 4 touched corpus docs consistent (voice, cross-refs, version-tagging); 0 broken relative links
  (all 4 docs swept). Markdown tables well-formed. No dead cross-refs.
- [x] [adversarial] The load-bearing §5 inventory count (21 = 11+2+5+3) ground-truthed against the real rext
  `patches/` dir at the code-of-record tag — EXACT match; also fenced by `test_patch_inventory.py`. Recorded in
  `decisions.md` § Adversarial review.

### Documentation
- [x] `demopatch-spec.md` §5/§4/§8, `cockpit-spec.md`, `frontend-tier.md`, `studio-desk.md` accurate vs the
  shipped rext code. CLAUDE.md untouched (its demopatch bullet is M247's sole-owner domain per the guardrail).
- [x] [nice-to-have] Added `(#M249-D{K})` decision trace tags at the 3 spots where a doc states a specific
  non-obvious choice (D1 spread-in-scope-symbol · D2 read import.meta.env directly · D3 overlay-not-build-arg).

### Tests & Benchmarks
- [x] 138 M249-touched tests GREEN at code-of-record `july-jitter-m249-harden`; flake gate 5/5 clean
  (varied order). No new gaps beyond the 3 the harden pass already deepened (+7 unit tests). No benchmarks
  apply (docs + demo-patch tooling).

### Decision Triage
- [x] D1-D4 → already blended into `demopatch-spec.md` §8 (additive-UI injection) + `frontend-tier.md` +
  `studio-desk.md` during build; trace tags added at close.
- [x] D5 (studio `.env.production.local` always-overwritten) → archive (maintainer-only rext-internal bug fix;
  fully recorded in `decisions.md` + fenced by `test_studio_desk_overwrites_a_stranded_overlay`).
- [x] Pre-existing test-failure item → Fate-2 confirm → M254 (recorded; deferral audit GREEN).
