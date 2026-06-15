# M31 — Decisions

_Implementation decisions with rationale, numbered `M31-D1`, `M31-D2`, … . Empty at scaffold; filled during build._

_Pre-decided at design (2026-06-15, see `.agentspace/scratch/roadmap-research-2026-06-15.md`):_
- _Fallback is openssl self-signed (NOT fail-loud) — the never-abort-a-good-bring-up contract._
- _BAPI is out of scope (plain HTTP, no browser TLS handshake)._
- _SANs = `127.0.0.1 localhost ::1`; cert CN is a non-issue (the pk validator checks SANs, not CN)._
- _`DEMO_NO_MKCERT=1` opt-out exists (dev-CA-in-trust-store is a real trust expansion)._

## M31-D1 — Factor the openssl fallback into a function (`gen_openssl_fapi_cert`)
The mkcert branch needs the openssl fallback in **two** places (mkcert-absent/opted-out, AND mkcert-mint-failed).
Rather than duplicate the openssl invocation (drift risk: the two copies could diverge on a future SANs/`-days`
edit), factored it into one `gen_openssl_fapi_cert()` defined just before the cert if-block, called from both
fallback branches. Output is byte-compatible with the prior verbatim openssl gen (same two files, same SANs, same
`-days 825`). A static test pins exactly-one-definition + ≥2-call-sites.

## M31-D2 — Test the cert step by extracting-and-running the block (values-blind), not just body-grep
The cert step (3a-bis) is a bring-up ACTION sitting **below** the `UP_INJECTED_LIB_ONLY` early-return, so it is
not sourceable as a lib-only function (unlike the frontend builders). Rather than settle for static `self.BODY`
pins alone (the `test_secret_preflight_*` precedent), the new `FapiCertStep` test ALSO does a **functional** run:
`awk`-extract the `gen_openssl_fapi_cert()` fn + the cert if-block into a tmp script, then source it through all
four branches (mkcert ok / mint-fail / `DEMO_NO_MKCERT` / mkcert-absent) + the keep-existing idempotency guard,
with **stub mkcert binaries** on a constrained PATH (no live docker, no real `mkcert -install`, no real CA touch).
This proves the branch logic, not just its presence. (Harness detail: `bash`/`env` are symlinked into the clean
PATH so the stub's `#!/usr/bin/env bash` shebang resolves — the real bring-up runs with a full PATH.)

## M31-D3 — PATH-based mkcert detection (`command -v mkcert`), not the hard-coded homebrew path
The roadmap research noted mkcert lives at `/opt/homebrew/bin/mkcert` on the dev box, but detection uses
`command -v mkcert` so the branch works regardless of how/where mkcert was installed (Linux, a non-homebrew Mac,
a custom prefix). Hard-coding the path would silently fall back to openssl on any box with mkcert elsewhere.

## M31-D4 — ZERO change to the cert-CONSUMING side (verified, not assumed)
Confirmed at the Phase 0b audit: `fake-fapi/main.go:28-31` reads `FAKE_FAPI_TLS_CERT/KEY` and serves
`ListenAndServeTLS`; `gen_injected_override.py:298,304-307` sets those env + mounts `<stack>/certs:/certs:ro`;
`inject.py` has no cert refs. All reference the cert by **path only** → a browser-trusted cert at the same path
serves identically. Only the `gen_injected_override.py:295` comment was touched (retire the "one-time
mkcert/import" framing). No functional change to any cert consumer.

## M31-D5 — Surfaced + confirmed covered: the dev-N `--local-content` UI path (Fate 2 / forward-note)
A future dev-N `--local-content` UI path would expose the same browser→FAPI TLS handshake and want this exact
mkcert wiring. Not landed here (no dev-N UI path exists today — it'd be net-new scope). Recorded as a one-line
**forward-note in the code comment** (candidate to extract the cert-mint logic into a shared helper, e.g.
stack-core, rather than re-inline) so whoever builds that path finds the pointer. No new backlog entry — the note
lives at the exact code site that would consume it.

## M31-D6 — (harden) The extracted-block test harness must run under production `set -euo pipefail`
The cert step is tested by `awk`-extracting `gen_openssl_fapi_cert()` + the cert if-block into a tmp script and
sourcing it (M31-D2). The build's harness sourced it under `set -uo pipefail` — but the real bring-up runs
`set -euo pipefail` (`up-injected.sh:13`). That `-e` is exactly what makes the cert step's
`mkcert -install >/dev/null 2>&1 || true` swallow **matter**: under `-e`, a failing `-install` would abort the
script before the mint unless the `|| true` catches it; without `-e` the swallow is inert. So a harden test for the
install-failure case (`-install` fails, mint succeeds) would **false-green** under the weaker harness — it passes
even with `|| true` deleted. Fix: switch all three `FapiCertStep` harnesses to `set -euo pipefail` (matching
production). Verified all four real branches still survive under `-e` (rc=0 each), and **mutation-proved** the
install-failure test now discriminates (removing `|| true` → the test FAILS with a `CalledProcessError`).
**General lesson:** when a test extracts-and-runs a bring-up code block, the harness's `set` flags must mirror the
real script's, or strict-mode-dependent guards (`|| true`, `|| pf_rc=$?`, the cmd-sub `|| true` patterns) go
untested. No change to `up-injected.sh` (the production code was already correct; only the test harness needed it).
