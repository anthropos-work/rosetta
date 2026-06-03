# M2 — Retro: Clerkenstein browser session + webhook coherence (JS)

## Summary
The final v1.0 milestone, and the **highest-technical-risk** one going in: it closes the last two Clerk
seams so a demo stack is Clerk-free end to end. The defining risk — "can `@clerk/*` be pointed at a fake
FAPI without forking the SDK?" — was **resolved at spike time in the strong direction** (M2-D1): clerk-js
derives its FAPI host *entirely* from the publishable key (`@clerk/shared parsePublishableKey`), so the
override is one env var, **no fork, no fallback exercised**. On that footing, 5 sections delivered: the
fake FAPI server (`fapi/`) + publishable-key codec, the fake BAPI server (`bapi/`) disarming the
platform's networked `orgclient` via an `api.clerk.com` redirect (the M1-D2 Fate-3 pickup, with the store
made concurrency-safe per M2-D2), the svix-signed webhook injector (`webhook/`), and a **second Alignment
DNA** (`clerk-js-5`, 9 genes, runner `cmd/jsfapirun`) scored at 100%/100% — proving the M0 framework is
**surface-generic**. 4 harden passes, then close. Both alignment gates 100%/100% (Go 22/22 + JS 9/9).

## Incidents this cycle
- **P1 (close-review bug): `orgclient.ChangeRole` nil-map panic + phantom-membership divergence (M2-D4).**
  `ChangeRole` checked only `validRole` + org-existence, then assigned `s.members[org][user]`
  unconditionally — so on an org with a nil members map (any org created via `CreateOrganization`) it
  **panicked**, and on a seeded org it **silently minted a phantom membership** for a non-member instead
  of returning `ErrNotMember`. Reachable via the `bapi/` server in a live demo (create org → PATCH a
  not-yet-member's role). **The alignment gate missed it** because the `ChangeRole` gene only targets the
  seeded existing member `o_1`/`u_1` — a sharp reminder that 100% on a thin gene set is not 100% coverage
  of reachable behavior. Fixed at close (membership guard mirroring `DeleteMembership`) + 2 regression
  tests; both gates unchanged.
- **P2 (hygiene): gofmt drift.** 3 test files (`bapi/malformed_test.go`, `cmd/clerkrun/main_test.go`,
  `fapi/fuzz_test.go`) had mis-aligned trailing comments — the harden passes had claimed "gofmt clean"
  without re-verifying after edits. Reformatted; no behavior change.
- **P2 (doc): stale repo README.** The clerkenstein `README.md` was M1-era (Go surface only); the per-unit
  handbook contract caught it at close. Refreshed to cover the full M2 surface.
- **0 flakes** (5/5 `-race` shuffled).

## What went well
- **The spike de-risked the whole milestone before any code.** The "highest-risk" framing assumed an SDK
  fork; an hour reading the *installed* `@clerk/shared`/`@clerk/clerk-js` turned the override into an env
  var. Reading the dependency beat speculating about it.
- **The M0 framework generalized for free.** Authoring a JS/FAPI DNA + a sibling runner and scoring it
  with the *same* `alignctl` confirmed the alignment test class is engine- *and* surface-agnostic — the
  browser-coherence claim ("the browser can't tell the difference") became a number, not an assertion.
- **The end-to-end coherence chain held:** the `fapi/` session-token endpoint mints the *same* HS256
  universal-key JWT M1's authn twin verifies, so the browser session and the backend agree — pinned by
  the `SessionToken/decoded-identity` gene (operator `exact`).

## What didn't
- **A gate is only as strong as its genes.** The ChangeRole bug lived on a path the DNA never exercised
  (non-member / fresh-org role change). The fix is a regression test, but the lesson is broader: the
  alignment gate measures the *enumerated* surface; reachable-but-un-enumerated paths still need ordinary
  adversarial review — which is exactly what caught it at close.
- **Harden "clean" claims need re-verification after the last edit.** gofmt drifted because a late edit
  wasn't re-checked. Cheap to fix, cheaper to avoid.

## Carried forward
- **None at the milestone level.** All findings landed Fate 1 in M2; deferral audit GREEN (0 repeat/
  chronic). The one open thread is **release-scoped, not milestone-scoped**: the
  `feat/demo-environment` → `main` reconciliation (M0-D6 / DEF-M0-01) is owned by
  `/developer-kit:close-release`, which is the immediate next step now that v1.0's last milestone is closed.
- **Live wiring** (the BAPI redirect + publishable-key env var inside a *running* multi-service demo
  stack) is demo-stack work — **M3 (v1.1)**, as scoped in the M2 overview "Out" list. M2 verified the
  fake servers against the SDK request contract + the alignment genes, not a live stack.

## Metrics delta
clerkenstein: **112 Go test/fuzz funcs** (107 tests + 5 fuzz; +84 vs M1's 28) across 7 packages; coverage
authn/orgclient/fapi 100%, clerkrun 97%, bapi 96%, webhook 96%, jsfapirun 94%; flake 5/5; gofmt/vet/
shellcheck clean. Alignment: **both** DNAs at 100% overall / 100% critical (Go 22/22, JS 9/9). Full
figures: [metrics.json](metrics.json).
