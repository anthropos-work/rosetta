# M253 iter-02 — decisions

## D3 — generate the chained manifests programmatically + round-trip-validate
The manifest loader (`manifest_loader.py`) is a strict YAML-subset that dedents block scalars by line-1's lead,
so a multi-line anchor whose first line is indented cannot be hand-authored faithfully. I wrote a generator
(`gen-manifests.py`) that extracts the pristine anchor by line range, left-strips ONLY line 1 (it matches
mid-file-line as a substring), emits 2-space-prefixed block scalars, then VALIDATES by round-tripping through
the real `manifest_loader.load_manifest` (asserting `anchor`/`replacement` are byte-exact) + confirms the anchor
occurs exactly once in the pristine file + computes the CHAINED shas (shell's post == no-thirdparty's pre). Both
patches then passed the real `demopatch check` (G1/G2/G6). This is the demopatch-spec §5-bis discipline applied
to authoring: the anchor is the contract; the sha is a baseline.

## D4 — lib-only rebuild vehicle for the iteration loop
To rebuild ONLY the studio image with the M253 patches without re-running the whole bring-up (seed/snapshot/
verify), I source the authoring `up-injected.sh` with `UP_INJECTED_LIB_ONLY=1` (the L1315 test seam → `return 0`
before the bring-up actions), which loads the REAL `build_frontend_studio_desk` + its env (HERE/STACK/DEMO_WS/
N/OFFSET/SCHEME/HOST), supply `PK_DEMO` from the current image, `docker image rm` to force a fresh build, call
the function (which applies all 5 patches → builds → LIFO-reverts, leaving the clone git-clean), then recreate
just the studio-desk container on the CONSUMED-clone compose (`--force-recreate --no-build --no-deps`). This
exercises the real ladder machinery + real demopatch, so the measured image is exactly what the ladder bakes.

## D5 — the fresh-green autoverify clause is not achievable on this warm demo-2 → M254
Fresh autoverify returned `green:false` (4 warnings, none studio-related — see iter-02/progress.md). The studio
surface + login + shell paint are all verified working by the FCP runner's own probes (5/5 shells, 0 login
bounces — which REQUIRES the fake-FAPI the autoverify probe claims is down). A fully-green verdict requires
re-set-dressing hiring + starting academy, both out of M253 scope. Per coordination rule 9 (overview.md), the
fully-green COLD-p95 confirmation is chartered to M254 (prove-on-billion). So M253 lands the fix + runner +
proves the number on a local demo; M254 confirms it cold + green on billion. Not a deferral against the
three-fate rule — it is the milestone's own planned split (Fate-2: already covered by M254).
