# M28 — Retro

_Provisioning engine + coverage/verify gate. The 2nd milestone of v1.6 "stage door". Closed 2026-06-14._

## Summary

M28 built the engine that WRITES — `stacksecrets provision` takes the secret source and writes each repo's
target `.env` values-blind: grouped by `(repo, target_file)` from the M27 DNA, alias-mapped per file (the
gh-token family incl. cross-repo `app/GH_TOKEN`; distinct-similar pairs like `OPENAI_KEY`/`OPENAI_API_KEY`
never auto-copied), append-only copy-if-absent (M28-D2) with `--force` to overwrite and an N=0 main-dev-stack
refusal mirroring `stackseed --reset`. The headline blocks-release safety landed + is test-pinned: provision
runs BEFORE the demo/dev injection override and NEVER re-arms the stripped prod `DIRECTUS_TOKEN` on a non-prod
stack — it writes the strip-on-non-prod family BLANK (`KEY=`), the exact state the override forces, so the base
`.env` + override agree and the key-present-only gene still passes coverage (M28-D3). The single value-carrying
boundary is `provision/io.go::sourceValues`, proven by a reflection-walk safety test (incl. unexported fields +
map keys) that no value surfaces in the Report / dry-run plan / errors. `check`/`measure` became demo-aware (a
`mintedSource` overlay so Clerkenstein-minted Clerk keys count without the source carrying them) and wired
non-fatally into `/dev-up` + `/demo-up` pre-flight via the shared `stack-secrets/preflight.sh` (warn standard /
fail critical / skip otherwise). The base scorer + profile-scoping were M27 (M27-D3.2) — reused unchanged. Code:
`rosetta-extensions` @ build tip tag `stage-door-m28` (ext head `9742126`).

The close surfaced 1 real code bug (the misplaced demo pre-flight block) — fixed Fate-1; the deferral audit was
GREEN.

## Incidents This Cycle

- **P2 (harden Pass 2) — `preflight.sh` crashed mid-run under `set -u` on bash 3.2.** The non-demo path expanded
  an empty `demo_flag` array bare (`"${demo_flag[@]}"`), which trips an "unbound variable" abort on the macOS
  system bash (`#!/usr/bin/env bash` → bash 3.2). A non-demo `/dev-up` secret pre-flight would CRASH instead of
  running — the exact silent-break the non-fatal contract exists to prevent. Fixed with the conditional-expansion
  guard `${arr[@]+"${arr[@]}"}` + a `/bin/bash`-3.x regression; the shell-portability invariant was backfilled
  into `safety.md §2.8` so future shell-tooling authors don't re-introduce it. (ext `4a30ad4`)
- **P2 (close review) — the demo secret pre-flight block was positioned above the lib-only test seam.** It sat in
  the `up-injected.sh` body, ABOVE the `UP_INJECTED_LIB_ONLY` early-return, so sourcing the script lib-only (the
  `test_frontend_build.py` unit tests, which exercise `build_frontend_*`/`preflight_vm_ram` with a stubbed docker
  in a sandbox lacking the sibling `preflight.sh`) fired the bring-up pre-flight at source time and crashed —
  20 frontend-build tests failed. The pre-flight is a bring-up ACTION, not a function definition; moved below the
  seam alongside the M19 VM pre-flight + pinned with a static positional regression. (ext `9742126`)

Both were caught + fixed inline (harden + close are exactly the nets for this class). No functional gap shipped.

## What Went Well

- **The DIRECTUS_TOKEN handoff from M27 paid off precisely.** The risk was named blocks-release with a
  regression-test requirement at M27 close; M28 landed it cleanly — write-BLANK-and-defer is the simplest correct
  composition with the override, and because the gene is `key-present`-only (no `nonempty`) the safety and the
  coverage score agree by construction. No clever runtime coordination needed.
- **The values-blind invariant scaled to the WRITE path.** Provision necessarily MOVES secret bytes, yet the
  reflection-walk safety test (covering unexported fields + map keys) proves the only place a value lands is the
  gitignored target — the boundary is a single function with a loud doc-comment.
- **M28-D1 (LiveKit de-alias) was a real correctness catch during build.** The M27 DNA had aliased a key+secret
  *pair* as one value; provision's alias-fallback would have written the API key's value under the secret's key
  if the secret were missing — a silent wrong-value provision. De-aliasing reserved the alias mechanism strictly
  for true same-value families.
- **Reuse worked: 0 scorer rebuild.** The base `check`/`measure` + profile-scoping were M27; M28 added only the
  demo-aware overlay + the wiring, exactly the trimmed scope (M27-D3.2), with no scope creep.

## What Didn't

- **Two pre-flight scripts, two empty-array / lib-only-seam traps.** Both incidents are the same root family:
  bring-up wrappers that must run-to-verdict under hostile shells and must not fire when sourced lib-only. The
  harden pass caught one (bash 3.2); the close review caught the other (the seam). The positional-seam class had
  no static guard until this close added one — the wiring lacked a test asserting the block runs after the seam.
- **The misplaced block passed the build + 3 harden passes.** It only surfaced when the *full* `test_frontend_build`
  file was run at close — the harden passes scoped to the M28-touched test files (`test_tooling.py` PreflightBehavior)
  and didn't re-run the sibling frontend-build suite that sources the script lib-only. Phase 4's full-suite mandate
  is what caught it; the lesson is to run the full sibling suites when a wiring change touches a shared script.

## Carried Forward

- **DEF-M27-02 — per-gene profile tag → DISCHARGED.** The conditional ("M28 revisits IFF it wires non-default-profile
  bring-ups") never triggered — M28 wired the default `graphql` profile. No residual work; the per-gene-tag variant
  remains a documented v1 option (M27-D3.4) for any future non-default-profile milestone.
- **Inherited release-level backlog (DEF-M10-01 / DEF-M21-01 / M25-D9)** — KEEP, re-signed at the v1.5 close, all
  orthogonal to secret provisioning.
- **To M29 (docs + skill):** the `secrets-spec.md` must document the provision target-file map, the alias vs
  distinct-similar rules, the DIRECTUS_TOKEN write-blank safety, the demo-aware coverage, and the non-fatal
  pre-flight contract. The `safety.md §2.8` bash-portability invariant is in place for the skill's shell work.

## Metrics Delta

(from `metrics.json`)

- **Go test funcs:** 980 → **1027** (+47, entirely the new M28 `stack-secrets` Go: the provision engine + demo
  overlay + 3-pass harden + the review-fix; `stack-secrets` 113 → 160). Prior 4 sections unchanged.
- **Python tests:** M28 added the demo-stack `PreflightBehavior` suite (8 funcs, harden) + 1 close positional
  regression; demo-stack 99 pass / dev-stack 74 pass at close.
- **Coverage (harden + close):** provision 87.3% → 94.8% (+7.5); secretdna 99.2%; cmd 96.4%; source 96.2%.
- **Flake:** **0** (Go 5/5 `-race -shuffle`, Python 5/5 sequential). `gofmt` + `go vet` + `shellcheck` clean.
- **Review findings:** 1 (0 scope / 1 code / 0 docs / 0 tests / 0 blend), Fate-1; 0 escape-hatch.
- **Field bugs:** 0 (the field-bake is M30). 2 real bugs caught + fixed inline (harden bash-3.2 + close seam).
