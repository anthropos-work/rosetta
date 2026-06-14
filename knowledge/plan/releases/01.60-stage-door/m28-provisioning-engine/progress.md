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
