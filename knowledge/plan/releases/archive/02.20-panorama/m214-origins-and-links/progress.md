# M214 — progress

Section checklist (closure = a MagicDNS-origin browser is admitted by CORS, cross-surface links resolve, the
required patch lands via the sha-pinned mechanism, and the recipe doc exists).

- [x] `CORS_EXTRA_ORIGINS` emission includes the HTTPS MagicDNS origin(s) + injection tests updated
- [x] studio-desk runtime `CLERK_SIGN_IN_URL`/`WEB_APP_URL` host-substituted (https for a public host)
- [x] `VITE_CLERK_SIGN_IN_URL` bake gap resolved (gitignored `.env.production.local` overlay; no Dockerfile ARG)
- [x] ant-academy `allowedDevOrigins` NEW `apply-*.sh` patch (sha-pinned, drift-refuse, idempotent, non-fatal)
- [x] next-web `urls.ts` WEB_APP_URL/HIRING_APP_URL — **documented residual** (evidence: 0-eject sweeps never surfaced them; D-URLS-1)
- [x] shipped demopatches carry the MagicDNS value (the `$SCHEME`/`$HOST` flip in `up-injected.sh`)
- [x] mixed-content check (no browser-facing http under HTTPS-everywhere — the scheme flip covers every surface)
- [x] NEW `corpus/ops/demo/tailscale-serve.md` + cross-ref updates (rosetta_demo, frontend-tier, clerkenstein, demo/README index)

**All sections landed.** rext code + tests: `panorama-m214` (commits `bf3edd1` CORS+redirects, `ca4cb0b`
scheme-flip + VITE bake, `4599a2d` ant-academy patch). rosetta docs + plan on `m214/origins-and-links`.

## M214: Hardening

### Pass 1 — 2026-07-11
Scope manifest (rext `panorama-m213..HEAD`): 5 source files — `stack-injection/gen_injected_override.py`
(`browser_scheme` + CORS/redirect emission; **99% stmt cov**, sole miss = line 485, the `if __name__` entrypoint,
uncoverable), `demo-stack/up-injected.sh` (the `$SCHEME` flip + the studio-desk VITE overlay trap),
`demo-stack/ant-academy.sh` (`$SCHEME` flip + patch apply/revert wiring), `stack-injection/apply-ant-academy-dev-origins.sh`
(NEW patch helper), `demo-stack/patches/ant-academy-dev-origins/ant-academy-dev-origins.yaml` — with tests in
`tests/test_injection.py`, `tests/test_ant_academy.py`, `tests/test_frontend_build.py`, `tests/test_tooling.py`.
The Python surface was already at 99%; the milestone's declared TOP RISKS are shell-behavioral (trap-safety +
drift-refuse), which pytest-cov doesn't instrument — so this pass deepened the shell behavior directly.

**Coverage delta (milestone-touched files):**
- `gen_injected_override.py` statements: 99% → 99% (flat; sole miss is the uncoverable `__main__` entrypoint).
- Shell surfaces (`up-injected.sh` studio-desk builder + `apply-ant-academy-dev-origins.sh`): behavioral coverage
  deepened (+5 tests) on the two trap/error-path surfaces the build-phase tests had left as gaps.

**Tests added** (commit `99c86b7`; demo-stack suite 241 → 246):
- `tests/test_frontend_build.py`: 2 — `test_studio_desk_failed_build_still_reverts_overlay_and_dockerignore`
  (the missing studio-desk analogue of the next-web build-FAILURE trap test: the `.env.production.local` overlay
  AND the transient `!.env.production.local` `.dockerignore` re-include are both trap-reverted even when
  `docker build` fails mid-way) + `test_studio_desk_does_not_clobber_a_preexisting_env_production_local` (the
  never-clobber-a-repo-file skip branch — a repo-shipped overlay is left byte-untouched, no `.dockerignore` edit).
- `tests/test_ant_academy.py`: 3 clone-independent helper error paths — `test_unknown_verb_is_refused`,
  `test_apply_missing_target_is_refused`, `test_revert_refuses_a_drifted_file` (refuse leaves the file
  byte-untouched). These pin the helper's refusal contract even on a bare rext checkout, where `REAL_NEXT_CONFIG`
  is absent and the behavioural round-trip / drift tests `skipUnless`-skip.

**Bugs fixed inline:** none — the studio-desk overlay/`.dockerignore` RETURN trap and the helper drift-refuse
guards behaved correctly under the new failure/refusal tests (they were untested, not broken).

**Flakes stabilized:** none observed — new tests deterministic across the 3-run sequential flake gate (5/5 ×3).

**Knowledge backfill:** no KB-worthy findings this pass — no new behavioral invariant surfaced (the trap-revert
contract is already documented in `decisions.md` D-VITE-SIGNIN-1 + the code comments; the helper's refusal exit
codes are self-documenting), and no bug was fixed. Question asked, nothing to blend.

### Stop condition
Stopped after Pass 1: the full Step 2b scan found no further meaningful gap (all three declared top-risk surfaces
— trap-revert, drift-refuse/idempotent re-apply, byte-identical-when-unset — are now pinned; no build-phase
bug-fix commits need regression tests; no parser/perf surface justifies fuzz/benchmark work), the Python coverage
delta is flat at an uncoverable ceiling (99%), and zero flakes. rext code-of-record re-tagged `panorama-m214` @
`99c86b7`.

## M214: Final Review

Close review 2026-07-11. Reviewed the rosetta docs+plan branch (`m214/origins-and-links`) + the FROZEN rext code
diff `b9f41dd..99c86b7` (tag `panorama-m214`). **4 findings, all Fate-1, all rosetta-side — zero rext touch (the
annotated tag stays put; rext README reconcile → close-release, D-CLOSE-3).**

### Scope
- [x] 0 gaps. Every `overview.md` In-item landed Fate-1 except the conditional `urls.ts`, which is the
  evidence-decided documented residual D-URLS-1 (re-confirmed still NOT a deferral in the close deferral audit).

### Code Quality
- [x] 0 must-fix. The frozen rext diff is high quality — one `browser_scheme`/`$SCHEME` predicate mirrored across
  `gen_injected_override.py` + `up-injected.sh` + `ant-academy.sh`, gated on `STACK_PUBLIC_HOST`, byte-identical
  when unset, trap-reverted overlays, drift-refuse + idempotent + non-fatal patch application. (Frozen anyway.)

### Documentation
- [x] [should-fix] `tailscale-serve.md:32` cited clerkenstein.md §"Remote reach over Tailscale" — actual heading
  is "Remote HTTPS over the tailnet (v2.2 M213)". Corrected the prose §-name (the link target already resolved).
- [x] [nice-to-have] Added the `(#M214-D-URLS-1)` reference tag to `tailscale-serve.md`'s "Documented residual"
  section so the residual decision traces back to `decisions.md`.
- [x] All 7 linked docs exist; recipe-browser-login §B + every cross-ref resolves; the recipe is accurate to code.

### Tests & Benchmarks
- [x] 0 gaps. stack-injection 147p/8s (=155) · demo-stack 383p/0s — both exit 0. Reconcile exactly with M213
  (stack-injection 152→155 +3; demo-stack 367→383 = +16 net-new test funcs). rext already hardened at 99% stmt.

### Decision Triage
- [x] Adversarial review — 3 scenarios recorded in `decisions.md` (all handled by existing guards; no code fix).
- [x] D-SCHEME-1 / D-VITE-SIGNIN-1 already blended into `clerkenstein.md` (ref-tagged); D-PATCH-1 / D-URLS-1 into
  `tailscale-serve.md` — no new blend needed.
- [x] D-CLOSE-3 recorded in `decisions.md` (Fate-2 → close-release rext re-tag, bundled with D-CLOSE-1/-2).
