# M31 — Retro

_mkcert-trusted FAPI cert (the browser-login render fix). The 1st of 2 milestones of v1.7 "house lights".
Two-repo (rosetta doc rewrite + planning artifacts; ext code @ tag `house-lights-m31`). Closed 2026-06-15._

## Summary

M31 fixed the live defect that triggered v1.7: a fresh browser at a demo's next-web rendered a **blank page**
because clerk-js's dev-browser handshake to the fake FAPI (`https://127.0.0.1:35400`) hit an **untrusted openssl
self-signed cert** (`net::ERR_CERT_AUTHORITY_INVALID`) → clerk-js aborted → blank. The fix automates a
locally-trusted **mkcert** FAPI cert into the demo bring-up: **one branch** in `up-injected.sh` step 3a-bis,
inside the existing keep-existing-cert guard — `command -v mkcert && [ "$NO_MKCERT" != 1 ]` → idempotent
`mkcert -install || true` + mint a leaf for `127.0.0.1 localhost ::1`. The historical openssl self-signed path
was factored into `gen_openssl_fapi_cert()` (byte-compatible output, called from BOTH the absent/opted-out branch
AND the mint-failed branch — no two-copy drift), and a `DEMO_NO_MKCERT=1` opt-out was added (mirrors the
`DEMO_NO_*` family) for operators who won't put a dev CA in their trust store. **Non-fatal throughout** (the
never-abort-a-good-bring-up contract). Critically, **ZERO change to the 3 cert-CONSUMING files** (`fake-fapi/main.go`,
`gen_injected_override.py` mount, `inject.py`) — they reference the cert by path only, so a trusted cert at the
same path "just works" (M31-D4, verified at the Phase-0b audit, not assumed).

The docs were rewritten manual→automatic: `recipe-browser-login.md §B` (+ the security/remote-VM/Firefox-`certutil`/
cert-expiry/`DEMO_NO_MKCERT` caveats), a `frontend-tier.md` one-liner, the demo-up SKILL note. The close-time
observable verify was satisfied by **composition** (M31-D7): demo-3 had been torn down, so rather than re-spin a
fresh `/demo-up`, the claim was proven as chromium-trusts-the-mkcert-cert (default context, NO `ignoreHTTPSErrors`:
200 / no cert error) vs chromium-rejects-the-old-openssl-self-signed (`ERR_CERT_AUTHORITY_INVALID` — the exact
blank-page cause) + the earlier cert-trusted→next-web-renders proof + the 11 `FapiCertStep` functional/edge tests.
Necessary + sufficient.

The close review found **2 findings, both Fate-1**: (1) `demo-stack/README.md` quoted "13 unit tests" (×2, stale
since ~M16/M28; actual 50) → reconciled to 50 + a note distinguishing it from GUIDE's guarded "28" curated-subset
count (the Phase-4 handbook-count-reconciliation discipline catching pre-existing drift at the milestone boundary);
(2) an adversarial scenario recorded in `decisions.md` (a zero-byte/truncated cert + the existence-only keep-guard →
NOT a regression, identical to pre-M31 behavior, documented repair path `rm fapi.crt` + re-up).

## Incidents This Cycle

None at close. No P0/P1/P2 incidents, no regressions, no flakes (5/5 randomized sequential, all 50/50). One
**test-harness fidelity** finding during harden (M31-D6, P2-class, not a production bug): the build's extracted-block
test harnesses sourced the cert block under `set -uo pipefail`, but the real bring-up runs `set -euo pipefail`
(`up-injected.sh:13`). Without `-e` the cert step's `mkcert -install … || true` swallow is inert, so the
install-failure harden test would have **false-greened** (passed even with `|| true` deleted). Fixed by switching all
three `FapiCertStep` harnesses to `set -euo pipefail`; **mutation-proven** the test now discriminates (removing
`|| true` → `CalledProcessError`). `up-injected.sh` was byte-identical before and after — the production code was
already correct; only the test harness needed it.

## What Went Well

- **The no-touch guarantee held — and was verified, not assumed.** The whole milestone's leverage came from the
  cert-consuming side referencing the cert by path only. The Phase-0b KB-fidelity audit confirmed `fake-fapi/main.go`
  + the `gen_injected_override.py` mount + `inject.py` are all path-only BEFORE the build started, so M31 swapped
  which tool mints without touching anything downstream. A small, contained, low-risk change by construction.
- **Factoring the fallback paid off immediately.** The mkcert branch needs the openssl fallback in two places
  (absent/opted-out + mint-failed). Extracting `gen_openssl_fapi_cert()` (vs duplicating the openssl invocation)
  removed the drift risk a future SANs/`-days` edit would have introduced, and a static test pins exactly-one-
  definition + ≥2-call-sites so the factoring can't silently regress.
- **The composition verify was honest and sufficient.** Rather than claim an un-run end-to-end demo render, the
  close decomposed the observable behavior into its load-bearing link (cert is browser-trusted) + the already-proven
  links (trusted→renders; tool mints at the path) and proved the one that was actually failing. The mkcert-vs-openssl
  chromium contrast IS the proof of the exact defect.
- **The strict-mode harness lesson generalizes.** M31-D6 records a reusable rule: when a test extracts-and-runs a
  bring-up code block, the harness's `set` flags must mirror the real script's, or strict-mode-dependent guards
  (`|| true`, `|| pf_rc=$?`) go untested. A genuine quality improvement that outlives this milestone.

## What Didn't

- **A pre-existing handbook count had drifted unnoticed across many milestones.** `demo-stack/README.md`'s "13 tests"
  was stale since well before M31 (the suite grew through M16/M28/etc. while the README count never followed). It was
  caught only because close-milestone's Phase-4 step-6 reconciliation now requires it — a reminder that README/handbook
  counts not under a guard (unlike GUIDE's `TestGuideDocTruth`-pinned "28") silently rot. Fixed Fate-1; consider a
  guard for the README count too if it drifts again.

## Carried Forward

- **Nothing M31-originated blocks anything.** The dev-N `--local-content` UI path's identical mkcert need (M31-D5) is
  a Fate-2 forward-note at the exact code site — net-new scope with no consumer today (no dev-N UI path exists). If
  one is ever built, the note points at extracting the cert-mint logic into a shared helper (stack-core?) rather than
  re-inlining.
- **ant-academy demo liveness** → M33 / roadmap-vision backlog (repro-first) — routed at the v1.7 design-roadmap, not
  an M31 deferral.
- **M32 studio-desk single-port/production** is the next (and last) v1.7 milestone — sequence after M31 (shared
  `up-injected.sh` + doc cluster). After M32 closes, v1.7 is ready for `/developer-kit:close-release`.
- **Release-level git carry-over** (push the v1.5/v1.6 ext tags to origin; the orphaned `m26/self-contained-demo`
  branch awaiting its own design-roadmap pass) — tracked in state.md, surfaces at v1.7 close-release.

## Metrics Delta

(from `metrics.json`)
- **Findings:** 2 (0 scope · 0 code-quality · 1 docs · 0 tests · 1 adversarial-record) — both Fate-1.
- **Field/production bugs:** 0 (one test-harness fidelity fix at harden — M31-D6, not a production bug).
- **Go tests:** 1027 → **1027** (+0 — M31 touched no Go).
- **Python tests:** the `FapiCertStep` class **+11** (`test_tooling.py` 47→50; demo-stack suite 99→110; the v1.6
  headline 459 → 470). `demo-stack/README.md` count reconciled 13→50.
- **Flake count:** 0 (5/5 randomized sequential, all 50/50; `test_frontend_build` 42/42).
- **Observable-behavior gate:** MET by composition (M31-D7) — chromium trusts the mkcert cert (200) vs rejects the
  openssl self-signed (`ERR_CERT_AUTHORITY_INVALID`) + cert-trusted→renders + 11 `FapiCertStep` tests.
- **Lint:** shellcheck 0.11.0 clean on `up-injected.sh`; `py_compile` clean; README-index guard exit 0.
- **Ext code:** `house-lights-m31` @ `6565ef8` (build); harden `815993f`; close review-fix `5022e72` — all on
  `m31/mkcert-fapi-cert`; the orchestrator finalizes the ext side (ff main + re-point tag + delete branch).
- **Deliverables:** the mkcert bring-up step (ext) + the `recipe-browser-login.md §B` rewrite + the frontend-tier
  one-liner + the demo-up SKILL note + the planning record.
