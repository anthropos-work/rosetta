# M31 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [x] `up-injected.sh` step 3a-bis: mkcert branch (idempotent `-install` + mint `127.0.0.1 localhost ::1`) inside the keep-existing guard
- [x] openssl fallback kept verbatim (mkcert-absent AND mint-failure paths); non-fatal throughout
- [x] `DEMO_NO_MKCERT=1` opt-out parsed + honored (mirrors the DEMO_NO_* flags)
- [x] code comments retired (`up-injected.sh` 3a-bis header, `gen_injected_override.py:295`) + the dev-N shared-helper forward-note
- [x] `corpus/ops/demo/recipe-browser-login.md §B` rewritten manual→automatic (+ security/remote/Firefox/expiry notes)
- [x] `frontend-tier.md` cert one-liner
- [x] demo-up SKILL browser-login note (+ `DEMO_NO_MKCERT` in the argument-hint)
- [x] verified: ZERO change needed to `gen_injected_override.py` / `inject.py` / `fake-fapi/main.go` (path-only consumers — M31-D4)
- [x] README-index guard exit 0
- [x] ext tag `house-lights-m31`
- [x] (close-time verify) browser-trust proven by composition (demo-3 was torn down — see M31-D7): chromium (no `ignoreHTTPSErrors`) trusts a mkcert cert (200, no cert error) vs rejects the old openssl self-signed (`ERR_CERT_AUTHORITY_INVALID`); + this-morning's "cert-trusted → next-web renders `/home`" + the 11 functional plumbing tests

## Notes
- **Ext commit** `6565ef8` (on ext `m31/mkcert-fapi-cert`, from base `868a68a`): the mkcert branch + the
  `gen_openssl_fapi_cert` factored fallback + `DEMO_NO_MKCERT` parse + the comment retire/forward-note +
  the `gen_injected_override.py:295` comment touch-up + a new `FapiCertStep` test class (8 tests).
- **Rosetta doc commit** `359eee4` (on `m31/mkcert-fapi-cert`): recipe-browser-login.md §B rewrite +
  frontend-tier.md one-liner + demo-up SKILL note.
- **Tests:** `FapiCertStep` 8/8 pass (5 static body-pins + functional all-4-branches + idempotency, values-blind,
  stub mkcert, no live docker). Full `test_tooling.py` 47/47 + `test_frontend_build.py` 42/42 (lib-only sourcing
  still works — the cert step's fn sits below the seam). `shellcheck 0.11.0` clean. README-index guard exit 0.
- **Cert-consumer no-touch verified** (M31-D4): `fake-fapi/main.go` + `gen_injected_override.py` mount +
  `inject.py` all reference the cert by path only — a trusted cert at the same path serves identically.
- **Why the last box stays open:** the live "fresh real browser renders next-web signed-in, no proceed-anyway"
  is the **close-time** verify (a live `/demo-up` + browser, not run at build per the milestone contract). The
  static + functional tests prove the cert-mint logic; the render proof is close-milestone's job.

## M31: Hardening

### Pass 1 — 2026-06-15
**Coverage (milestone-touched code):** the surface is one bash branch in `up-injected.sh` step 3a-bis — no
line-coverage tool applies (bash). "Coverage" = branch enumeration. The build's 8 `FapiCertStep` tests covered the
4 main branches + idempotency; this pass closed the genuine remaining edges.

**Tests added (3 regression/edge, in `FapiCertStep`):**
- `test_func_install_failure_still_mints_via_mkcert` — the `mkcert -install … || true` swallow: a FAILING
  `-install` (fresh box / declined trust-store write) must NOT abort the step; the mint still produces the leaf.
  Distinct from the mint-failure→openssl branch. **Mutation-proven** discriminating.
- `test_func_certs_path_with_whitespace_is_quoted` — a `$CERTS` path with a space survives BOTH the mkcert mint
  and the openssl fallback (the `"$CERTS/…"` quoting holds; no word-splitting).
- `test_func_keep_existing_guard_keys_on_crt_only_partial_state` — the outer guard keys on `fapi.crt` only; a
  crt-present/key-absent partial state is skipped (never re-mints over a kept cert, never repairs the missing key).

**Bugs fixed inline:** none in `up-injected.sh` (the cert step is correct). One **test-harness fidelity fix** (see
M31-D6): the build's `_run_block` + the two inline cert harnesses sourced the extracted block under
`set -uo pipefail`, but the real bring-up runs `set -euo pipefail` (`up-injected.sh:13`). Without `-e` the
`|| true` swallow is inert, so the install-failure test would false-green. Switched all three `FapiCertStep`
harnesses to `set -euo pipefail`; verified all four real branches survive under `-e`. `up-injected.sh` unchanged.

**Flakes stabilized:** none — the stub-binary functional tests use per-test `tempfile.mkdtemp()` dirs (no shared
state/ports). Flake gate: 3 consecutive clean sequential runs of `FapiCertStep` (11/11) + the 3 new tests (3/3).

**Values-blind reinforcement (verification, no code change):** the cert step's 5 `log` lines reference the cert by
**path only**, never the key/cert body; no `cat`/`echo`/`printf` of key contents. The tests read only `fapi.crt`'s
**first line** to discriminate which branch fired (`MKCERT-LEAF` / `-----BEGIN CERTIFICATE-----` / `PRE-EXISTING`);
never assert on `fapi.key` body. TLS material stays on disk, off stdout/log/report.

**Doc-vs-code consistency (verification, no new tests):** `recipe-browser-login.md §B` + `frontend-tier.md:26-29` +
the demo-up SKILL `argument-hint` (`DEMO_NO_MKCERT=1`) all match the shipped `up-injected.sh` branch (mkcert
idempotent `-install`, mint SANs `127.0.0.1 localhost ::1`, openssl fallback on absent/`DEMO_NO_MKCERT=1`/mint-fail,
byte-compatible). The "~2.25y mkcert leaf" claim matches mkcert v1.4.4. **No drift.** (GUIDE.md "28 unit tests" is the
guarded curated-subset count — `FapiCertStep`/preflight classes are excluded by the count guard by design — so the
+3 tests need no GUIDE bump and the guard stays green.)

**Knowledge backfill:** recorded the harness-strict-mode lesson as `M31-D6` in `decisions.md` (the test harness for
an extracted bring-up block must reproduce production `set -euo pipefail` or error-swallow guards aren't exercised).
The cert step's behaviors (the 4 branches + the `|| true` swallow + the crt-only guard) are already documented in
`spec-notes.md` and `recipe-browser-login.md §B`; no further corpus edit warranted.

**Result:** `test_tooling.py` 47 → 50 (all pass); `test_frontend_build.py` 42; shellcheck clean; README-index
guard exit 0. Ext harden commit `815993f` on `m31/mkcert-fapi-cert` (tag `house-lights-m31` unmoved @ `6565ef8`).

### Stop condition
Scan clean — all four main branches + the `|| true` swallow + partial-state + whitespace + idempotency now covered;
the remaining theoretical cases (`DEMO_NO_MKCERT` AND mkcert-absent) are logically subsumed and would be shallow.
Single-pass stop (small bash surface; the rule forbids padding).

## M31: Final Review

_close-milestone review of the whole milestone (both repos). Deferral re-audit GREEN
([audit-deferrals/deferral-audit-2026-06-15-m31-close.md](audit-deferrals/deferral-audit-2026-06-15-m31-close.md))._

### Scope
- [x] All `overview.md` In-list items delivered Fate-1 (mkcert branch + openssl fallback factored + `DEMO_NO_MKCERT`
      + comments retired + forward-note + recipe §B rewrite + frontend-tier one-liner + SKILL note + zero-touch on
      the 3 cert-consumers). 0 silent drops; the 4 deferral candidates all correctly fated (audit GREEN).

### Code Quality
- [x] [verify] ext branch well-factored — `gen_openssl_fapi_cert()` removes the two-copy openssl drift risk;
      if-then-else (no SC2015); non-fatal throughout; PATH-based `command -v mkcert` detection. No issues.
- [x] [verify] shellcheck 0.11.0 clean on `up-injected.sh`; `py_compile` clean on `gen_injected_override.py`.

### Documentation
- [x] [must-fix] `demo-stack/README.md` quoted "13 unit tests" (×2) — stale (actual 50). Reconciled to 50 + a
      note distinguishing it from GUIDE's guarded "28" curated-subset count (M31-D8). Pre-existing drift, Fate-1.
- [x] [verify] all 4 rosetta doc diffs match the shipped shell branch; cross-refs resolve; README-index guard exit 0.
- [x] [verify] GUIDE.md "28 unit tests" correct-by-guard (`TestGuideDocTruth`; curated-subset = 28) — no change.

### Tests & Benchmarks
- [x] [verify] `test_tooling.py` 50/50 + `test_frontend_build.py` 42/42 pass; flake gate clean (Phase 8). No new
      test gaps — `FapiCertStep` covers all 4 branches + idempotency + install-failure + whitespace + partial-state.

### Adversarial review (Phase 2c)
- [x] Zero-byte/truncated-cert + existence-only keep-guard scenario examined → NOT a regression (identical to pre-M31
      openssl-only behavior); documented repair path exists (`rm fapi.crt` + re-up). Recorded in `decisions.md`.

### Decision Triage
- [x] M31-D1..D8 reviewed. The cert-step BEHAVIORS (4 branches, `||true` swallow, SANs, fallback, security/remote/
      Firefox/expiry caveats) are already blended into `recipe-browser-login.md §B` + `frontend-tier.md` + the SKILL
      during build/harden. D1/D2/D3/D6 (factoring, test-harness mechanics, strict-mode) are maintainer-only → archive
      in `decisions.md`. D4 (no-touch consumers) already in recipe §B + the code comment. D5 (forward-note) lives at
      the code site. D7 (compose verify) + D8 (count reconciliation) maintainer-only → archive. No further knowledge
      blend needed — the user/developer-facing surface is the recipe, which is current.
