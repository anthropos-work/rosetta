---
title: "Deferral Audit — release v1.7 \"house lights\" (close-release Phase 1b)"
date: 2026-06-15
scope: release
invoked-by: close-release
---

## Verdict
GREEN

- **No in-release repeat deferrals.** v1.7 has exactly two milestones (M31, M32); both closed GREEN with **zero
  open in-milestone punts**. There is no item deferred across ≥2 in-release milestones — the necessary condition
  for a repeat-deferral pattern — so RED is structurally impossible from in-release deferrals.
- **No AGED_OUT items.** None of the five inherited backlog items lives in an in-release v1.7 milestone, and none
  of v1.7's touched feature areas (mkcert FAPI cert in `demo-stack`; studio-desk single-port in `stack-injection`)
  overlaps any inherited item's area — so the "feature area touched by a later milestone" aging trigger does not
  fire. All five were re-signed at the v1.7 design-roadmap pass (2026-06-15) = fresh authority.
- **No blockers.** Nothing requires a forced Fate-1 conversion.

## Summary
- Total deferrals in scope: **0 in-release milestone punts** + **5 inherited release-vision backlog items** (re-confirmed)
- Single deferrals (in-release): 0
- Repeat deferrals (in-release): 0
- Chronic / drift patterns flagged: 0
- Aged-out items: 0
- Blocking items: 0

## Deferral Inventory

### In-release milestone deferrals — NONE
Independent re-walk of both milestone dirs (`grep -niE "defer|postpone|later|out of scope|future milestone|
tracked for|follow-up|backlog|escape.?hatch|punt"`, excluding the `audit-deferrals/` subfolders) surfaced only
**design-time scope boundaries** and **already-fated items** — no raw punted unit of work:

```yaml
- id: M31/BAPI cert
  classification: DESIGN-TIME SCOPE BOUNDARY (not a deferral)
  evidence: clerkenstein/cmd/fake-bapi is plain http.ListenAndServe (server-to-server); the browser never does a
            BAPI TLS handshake → no cert work exists to defer. (decisions.md:7; spec-notes.md:32; kb-fidelity-audit PAIRED)

- id: M31/M32 ant-academy demo liveness  (overview Out: lines)
  classification: RELEASE-LEVEL ROUTING decided at v1.7 design-roadmap (→ M33, roadmap-vision backlog, repro-first)
  evidence: roadmap-vision.md:42 records "M33 — ant-academy demo liveness (deferred from v1.7 design, 2026-06-15,
            repro-first)". Not an M31/M32 punt — routed by the design pass before coding.

- id: M31-D5  dev-N --local-content UI path wants the same mkcert wiring
  classification: CORRECTLY FATED — Fate 2 (in-code forward-note), net-new scope with NO consumer today
  evidence: decisions.md M31-D5 — no dev-N --local-content UI path exists; building one is net-new scope outside
            v1.7's demo-UI-hardening thesis. Forward-note lives at the exact up-injected.sh code site (candidate:
            extract cert-mint into a shared stack-core helper). No backlog entry warranted (no problem to solve yet).

- id: M32/cert (overview Out: "the cert (M31)")
  classification: SIBLING-MILESTONE BOUNDARY (M31's delivered+merged scope) — nothing to land in M32.
```

### Composition closures (NOT deferrals — work is DONE)
Two close-time "observable behavior" boxes were satisfied **by composition** because the demo-3 stack the live
defect was hit on had been torn down by close time, and re-spinning a fresh `/demo-up` standalone is
disproportionate. Verified each is a real necessary+sufficient chain backed by **landed, merged regression tests**
in ext HEAD `7b17c39` (not a promise):

- **M31-D7** (close-verify): chromium default-context (no `ignoreHTTPSErrors`) trusts the mkcert leaf (`200`) vs
  rejects the openssl self-signed (`ERR_CERT_AUTHORITY_INVALID`, the exact blank-page cause) + the earlier
  cert-trusted→renders proof + the **11 `FapiCertStep` functional/edge tests** (confirmed present:
  `test_func_mkcert_success_mints_trusted_leaf`, `..._falls_back_to_openssl`, `..._demo_no_mkcert_forces_openssl`,
  `..._keep_existing_cert_is_idempotent`, `..._non_fatal` guard, whitespace-quoting, etc.).
- **M32-D5** (close-smoke): `NODE_ENV=production` pinned by the mutation-checked regression test
  `test_studio_desk_env_pins_node_env_production` → `isProduction=true` → the production `sendFile` block that
  covers every dev-block route (M32-D1 code-read, "NO GAP"); the dead-`:9100` CORS origin dropped + locked by the
  two-membership exact-set assert. A fresh `/demo-up` re-demonstrates both fixes live on demand (operator action).

Both mirror an accepted close pattern; both are **resolved**, not pending — nothing to fate.

### Inherited release-vision backlog (cross-release / unscheduled — re-aged this pass)
The five items named by the orchestrator live in `roadmap-vision.md` §"Unscheduled backlog (not a planned
release)" — cross-release/unscheduled **by design**, NOT in-release v1.7 milestone deferrals:

```yaml
- id: M33  (ant-academy demo liveness)
  status: re-signed at v1.7 design 2026-06-15, repro-first. Native (nohup) surface; root cause UNCONFIRMED (likely
          session-reaping, not a tooling bug); scope only after reproducing. Smallest payoff surface. NOT aged-out.
- id: M26  (self-contained-demo)
  status: ORPHANED ext effort (branch m26/self-contained-demo @ 25ab855, tag prop-room-m26, unmerged+unpushed).
          Status is "unplaced — awaits its OWN design-roadmap pass", NOT "deferred-within-a-release". NOT aged-out.
- id: DEF-M10-01  (cloud SnapshotStore + S3 blob bytes)
  status: gated on an EXTERNAL dependency (eu-west-1 S3 read access landing — verified not wired), not time/scope.
          User-facing sting already removed (v1.5 keeps the asset plane on prod public links). Re-signed v1.5. NOT aged-out.
- id: DEF-M21-01  (replayCmd conn-seam hermetic test)
  status: LANDED at v1.5 close-release 2026-06-14 (survives the branch merge). Resolved-shaped, not pending.
- id: M25-D9  (dev-N taxonomy replay rc=4 migrate-ordering)
  status: dev-only, orthogonal to the content-serve done-bar (DB-2 GREEN). Tracked dev migrate-ordering follow-up. NOT aged-out.
```

## Repeat-Deferral Patterns
**None.** v1.7's two milestones each punt zero work of their own; with no in-release milestone holding a deferred
item, there is no pair to form a repeat. The five inherited items live in roadmap-vision (cross-release/unscheduled),
so they cannot constitute an in-release repeat-deferral pattern. No CHRONIC_DEFER, no DRIFT_DEFER.

## Fate-1 Investigation
- M31/BAPI: Fate-1 N/A — no work exists (plain HTTP, no browser handshake). Correct boundary.
- M31/M32 ant-academy liveness → M33: Fate-1 NO — separate native surface, repro-first, no confirmed repro;
  landing in v1.7 would be net-new release scope. Already routed at design.
- M31-D5 dev-N shared helper: Fate-1 NO — no dev-N `--local-content` UI consumer exists; the helper would be a
  solution without a problem. The in-code forward-note is the correct Fate-2 artifact; no backlog entry warranted.
- M32/cert: Fate-1 N/A — M31's delivered+merged scope.
- M31-D7 / M32-D5 composition closures: LANDED — proven this close, backed by merged regression tests. Nothing to fate.
- Inherited backlog (M33 / M26 / DEF-M10-01 / DEF-M21-01 / M25-D9): Fate-1 NO across the board — external gate /
  orphaned-effort-needing-its-own-design / dev-only-orthogonal / already-landed. All correctly cross-release or unscheduled.

## Recommendations
- All in-release candidates → **no action** (design boundaries, in-code forward-note, or resolved-by-composition; each
  already carries its correct fate in M31/M32 `overview.md`/`decisions.md`/`spec-notes.md` + `roadmap-vision.md`).
- Inherited backlog → **no action** (all correctly cross-release/unscheduled in `roadmap-vision.md`; re-signed at v1.7 design).
- **Release-level git/roadmap-vision HYGIENE (for close-release, not a deferral):** push the unpushed v1.5/v1.6 ext
  tags to origin; the orphaned `m26/self-contained-demo` + `wip/clerkenstein-browser-login` branches await their own
  design-roadmap pass. This is housekeeping the synthesizer/close-release owns — it is NOT a Fate-1-convertible deferral.

## Applied Changes
**None.** Every candidate already carries its correct fate in the milestone docs + `roadmap-vision.md`. No new task
added, no decision rewritten, no roadmap/roadmap-vision edit needed. (Audit is read-only by design; this report is the
sole artifact.)

## Blocking Items (require user decision)
**None.**
