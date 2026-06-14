# M28 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [x] `provision` verb + per-repo target-file map (ant-academy → `code/.env.local` pinned) — DNA-driven `(repo,target_file)` grouping; nested dirs MkdirAll'd
- [x] Alias-mapping per target file (one source value → all its per-file aliases) — `resolveAliasSources`; gh-token family incl. cross-repo `app/GH_TOKEN`; distinct-similar never auto-copied
- [x] Idempotency + overwrite policy (copy-if-absent default, `--force`) — append-only write; re-run is a no-op; `TestProvision_Idempotent` / `_CopyIfAbsent` / `_Force`
- [x] N=0 main-dev-stack guard (refuse without `--force`) — mirrors `stackseed --reset`; `TestProvision_RefusesMainDevStackWithoutForce`
- [x] Compose-with-injection-override: never re-arm prod `DIRECTUS_TOKEN` on non-prod / `--local-content` **+ regression test** — `StripOnNonProdKeys` written BLANK on non-prod; **headline `TestProvision_NeverReArmsDirectusTokenOnNonProd`** + the armed-token edge tests (copy-if-absent + --force) (M28-D3)
- [x] `PreflightEnv`-passing env emission — strip-on-non-prod blanks the Directus write tokens (mirrors `PreflightEnv`'s reject set, `safety.md §2.2`); the provisioned env carries no live Directus write token on non-prod
- [x] `check`/`measure`: Overall + Critical (gate==100%) + per-repo rollup — **shipped in M27 (M27-D3.2); reused, not rebuilt**
- [x] Demo-aware coverage (Clerk keys minted-OK) — `MeasureForStack` + `mintedSource` overlay; `secretdna/demo.go`; `TestMeasureForStack_DemoMintsClerkKeys`
- [x] Non-fatal pre-flight wiring into `/dev-up` + `/demo-up` (warn standard, fail critical) — shared `stack-secrets/preflight.sh`; wired into `dev-stack` cmd_up + `demo-stack/up-injected.sh`; `PreflightBehavior` suite + static pins
- [x] Profile-scoping decision settled + implemented — **settled in M27** (default `graphql`; profile-gated keys `waived-profile-gated`); reused unchanged
- [x] Hard safety verified: no verb reads/echoes/logs a value — `provision/io.go::sourceValues` the sole value-carrying boundary; reflection-walk + CLI-output + preflight no-leak tests
- [x] Ext tag `stage-door-m28` — created on the complete green tree (ext `32f258b`)

## Notes
- Code: `rosetta-extensions` @ tag `stage-door-m28` (3 commits a40163d → 32f258b on ext `m28/provisioning-engine`).
- Decisions: M28-D1 (LiveKit de-alias), M28-D2 (append-only values-blind write), M28-D3 (DIRECTUS_TOKEN
  non-rearm = write blank).
- All tests green; gofmt + go vet clean; `-race -shuffle` clean on provision + secretdna; shellcheck-clean on
  preflight.sh + the two wired bring-up scripts; the 73 existing dev-stack subprocess tests still pass (preflight
  skips non-fatally with no source dir).
- The base `check`/`measure` scorer + profile-scoping were M27 deliverables (M27-D3.2) — M28 reused them.

## M28: Hardening

### Pass 1 — 2026-06-14 (ext `9541220`)
**Coverage delta (milestone-touched packages):**
- provision: 87.3% → 93.4% (+6.1)
- secretdna: 98.5% → 99.2% (+0.7)

**Tests added (+13 Go funcs):**
- `provision/provision_harden_test.go`: ParseStackN negative/malformed boundary + Provision propagation;
  source-root-doesn't-exist error; cross-repo alias routing to a frontend repo's conventional file
  (`sourceFileForRepo`); alias-family-with-no-carrier → MISSING; the value-vanished-after-plan no-corrupt-line
  path; MkdirAll/unreadable-target write errors (values-blind error text asserted); the **whole strip-on-non-prod
  family blanked at once** (headline safety, widened) + the prod-target arms-all inverse; `Skipped()` rollup
  accessor + per-file/report consistency; malformed-source values-blind parse (comments/blanks/`export `/`=`-less).
- `secretdna/demo_harden_test.go`: opaque-shaped minted keys (`CLERK_WEBHOOK_SECRET`/`CLERK_JWT_KEY`);
  non-minted passthrough on every overlay method (no false mint); `Keys()` passthrough; `MeasureForStack(demo=false)
  == Measure` boundary.

**Bugs fixed inline:** none in Pass 1.

### Pass 2 — 2026-06-14 (ext `4a30ad4`)
**Risk surface:** the `preflight.sh` non-fatal contract (the warn-standard / fail-critical / skip-non-fatal split)
+ the demo-aware wiring (risk surface 5). The build-phase `PreflightBehavior` suite covered skip/fail/pass/no-leak
but **not the WARN half**, the missing-DNA skip, or the inconclusive-rc non-fatal mapping.

**Tests added (+8 Python funcs, demo-stack `PreflightBehavior` + `TestShellcheck`):**
- warns-on-standard-but-still-passes (rc 0 + a loud ⚠ surface — the missing WARN half);
- skip-when-no-DNA (rc 2); inconclusive-rc-is-non-fatal (rc 2 — a pre-flight bug never blocks a good bring-up);
- unknown-arg-tolerated (forward-compat); the bash-3.2 regression below; `TestShellcheck` now pins `preflight.sh` clean.

**Bugs fixed inline:**
- **`preflight.sh` crashed mid-run under `set -u` on bash 3.2** (the macOS system bash, which `#!/usr/bin/env bash`
  resolves to) — the NON-demo path expands an **empty** `demo_flag` array, and bare `"${demo_flag[@]}"` trips an
  "unbound variable" abort (rc 127). A non-demo `/dev-up` secret pre-flight would CRASH instead of running its
  check — the exact silent-break the non-fatal contract exists to prevent. Fix: the conditional-expansion guard
  `${demo_flag[@]+"${demo_flag[@]}"}` (empty → nothing, set → `--demo`), safe on bash 3.2 AND bash 5. Regression
  `PreflightBehavior.test_non_demo_path_survives_set_u_on_bash32` verified to FAIL pre-fix, PASS post-fix.
  (commit ext `4a30ad4`)

**Doc:** `demo-stack/GUIDE.md` advertised test count 27 → 28 (the `TestGuideDocTruth` drift guard fired when
`TestShellcheck` gained the preflight pin).

### Pass 3 — 2026-06-14 (ext `5f1dfc8`)
**Coverage delta:** provision 93.4% → 94.8% (+1.4).

**Tests added (+5 Go funcs):** ant-academy carrier → cross-repo sibling alias routing; an alias carrier whose
sibling repo has no source file still resolving from the carrier; `sourceValues` + `hasSourceKey` on an absent
file (the documented empty-not-error contract).

**Bugs fixed inline:** none.

**Knowledge backfill:** the bash-3.2 empty-array-under-`set -u` trap was blended into the tooling-safety knowledge
(`corpus/ops/safety.md`) as a shell-portability invariant for the on-every-bring-up scripts — so M29 (docs+skill)
+ future shell-tooling authors don't re-introduce it.

### Stop condition
Stopped after Pass 3: the Step-2b scan found no new *meaningful* behavioral gap, and the residual provision
coverage (~5%) is confirmed-defensive — bufio scanner-error branches (a >1MiB unterminated line), the TOCTOU
value-vanished no-corrupt-line guard, and `WriteString`/`Close` mid-write failures — reachable only via
fault-injection / a concurrency race; covering them would add flake for marginal value (a <2% delta). The flake
gate is clean (3 sequential runs of the Go + preflight/shellcheck suites); `-race -shuffle` clean across all three
Go packages. **Final: provision 94.8%, secretdna 99.2%, cmd 96.4%; +20 Go test funcs (140→160) + 8 Python preflight
funcs; 1 real bug fixed inline.** Note: harden does NOT move the `stage-door-m28` tag — the tag still marks the
build tip (`32f258b`); the harden commits (`9541220`/`4a30ad4`/`5f1dfc8`) live on the `m28/provisioning-engine`
branch ahead of it, to be tagged at close.
