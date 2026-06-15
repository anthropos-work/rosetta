---
title: "Deferral Audit — M31 (close)"
date: 2026-06-15
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals (v1.7 "house lights" is in development; M31 is its FIRST milestone — nothing inherited from a prior in-release milestone).
- Every item surfaced has a clear, already-recorded fate. No CHRONIC/DRIFT patterns; no AGED_OUT items.

## Summary
- Total deferrals in scope: 4 candidate items examined
- Single deferrals: 0 (all 4 are design-time scope boundaries or already-fated items, not punted work)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0

## Deferral Inventory

The grep for `defer|postpone|later|out of scope|future milestone|tracked for|follow-up|backlog` across M31's
`overview.md` / `progress.md` / `decisions.md` / `spec-notes.md` surfaced four candidates. On inspection, NONE is a
punted unit of work that erodes scope:

```yaml
- id: DEF-M31-cand-01
  item: "fake BAPI cert / browser TLS"
  origin_milestone: M31
  reason_recorded: "BAPI is out of scope (plain HTTP, no browser TLS handshake)" (decisions.md pre-decided; spec-notes.md §BAPI)
  classification: DESIGN-TIME SCOPE BOUNDARY (not a deferral)
  why: the fake BAPI serves plain `http.ListenAndServe` server-to-server only; the browser never does a BAPI TLS
       handshake, so there is no cert work to defer. Correctly out-of-scope by nature.

- id: DEF-M31-cand-02
  item: "ant-academy demo liveness"
  origin_milestone: M31 (overview.md:45 — listed under `Out:`)
  reason_recorded: "ant-academy liveness (backlog)"
  classification: RELEASE-LEVEL ROUTING decided at design-roadmap (Fate-2-shaped → M33 / roadmap-vision backlog, repro-first)
  why: routed by the v1.7 design-roadmap pass, not by M31. state.md + roadmap.md both record "ant-academy liveness
       (M33) → roadmap-vision backlog (repro-first)". Not M31's to land; not an in-release milestone item.

- id: DEF-M31-cand-03
  item: "close-time live end-to-end browser render"
  origin_milestone: M31 (overview.md:55 — "deferred from design")
  reason_recorded: "end-to-end proof deferred from design (design rested on the Playwright-ignoreHTTPSErrors-renders proof)"
  classification: RESOLVED THIS CLOSE (not pending)
  why: the close-time verify was satisfied by composition — M31-D7. Chromium (default context, no ignoreHTTPSErrors)
       trusts the mkcert cert (200/no error) vs rejects the openssl self-signed (ERR_CERT_AUTHORITY_INVALID) + the
       earlier cert-trusted→renders proof + the 11 FapiCertStep functional/edge tests. Necessary + sufficient.

- id: DEF-M31-cand-04
  item: "dev-N --local-content UI path wants the same mkcert wiring (extract a shared helper)"
  origin_milestone: M31 (decisions.md M31-D5)
  reason_recorded: "Not landed here (no dev-N UI path exists today — it'd be net-new scope). Recorded as a one-line
       forward-note in the code comment … No new backlog entry — the note lives at the exact code site that would consume it."
  classification: CORRECTLY FATED — Fate 2 (forward-note), net-new scope with no consumer today
  why: there is no dev-N --local-content UI path in existence; building one would be net-new scope outside v1.7's
       demo-UI-hardening thesis. The forward-note sits at the exact `up-injected.sh` 3a-bis comment a future builder
       would read, pointing at the shared-helper extraction. This is the three-fate rule's correct outcome for an
       item with no current home, not a punt.
```

## Repeat-Deferral Patterns
None. M31 is v1.7's first milestone; there is no prior in-release milestone to repeat-defer from.

## Fate-1 Investigation
- DEF-M31-cand-01 (BAPI): Fate-1 N/A — no work exists (plain HTTP, no browser handshake). Correct boundary.
- DEF-M31-cand-02 (ant-academy liveness): Fate-1 NO — a separate milestone area (M33), repro-first, no repro yet;
  landing it in M31 would be net-new release scope. Already routed at design (LAND-NEXT-shaped / roadmap-vision).
- DEF-M31-cand-03 (e2e render): LANDED — proven by composition this close (M31-D7). Nothing to fate.
- DEF-M31-cand-04 (dev-N shared helper): Fate-1 NO — no dev-N UI consumer exists; the helper would be a solution
  without a problem. The code-site forward-note is the right artifact (Fate 2). No new backlog entry warranted.

## Recommendations
- DEF-M31-cand-01 → no action (design boundary).
- DEF-M31-cand-02 → no action (already routed to M33/roadmap-vision at design-roadmap time).
- DEF-M31-cand-03 → no action (resolved this close).
- DEF-M31-cand-04 → no action (correctly Fate-2 via the code forward-note; M31-D5 records it).

(Outward-facing release-level git carry-over noted in state.md — push the v1.5/v1.6 ext tags to origin; the orphaned
`m26/self-contained-demo` branch awaiting its own design-roadmap pass — is release/roadmap-vision hygiene, NOT an M31
milestone deferral. Out of this audit's milestone scope; flagged for the orchestrator / a future close-release.)

## Applied Changes
None. All four candidates already carry their correct fate in M31's own docs (overview/decisions/spec-notes) +
state.md/roadmap.md. No new task added, no decision rewritten, no roadmap edit needed.

## Blocking Items (require user decision)
None.
