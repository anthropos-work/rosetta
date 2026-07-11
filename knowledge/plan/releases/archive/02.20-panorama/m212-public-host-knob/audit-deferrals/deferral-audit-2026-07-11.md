---
title: "Deferral Audit — milestone M212 (close)"
date: 2026-07-11
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no chronic/drift patterns; no aged-out items (M212 is the release-opening
  milestone — nothing inherited within v2.2). Every in-scope item is a clean Fate-2 routing to a
  specific milestone of THIS release that already owns it, verified in the target's `overview.md`.

## Summary
- Total deferrals in scope: 2 substantive + the designed §Out decomposition (all Fate 2)
- Single deferrals: 2 (both confirmed-owned, no user decision required)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0 (all records dated 2026-07-11)

## Deferral Inventory

- id: DEF-M212-01
  item: "gen_injected_override CORS + Clerk sign-in/web-app URL EMISSION at the public host"
  origin_milestone: M212
  first_deferred_on: 2026-07-11
  last_seen_in: decisions.md D-IMPL-3; gen_injected_override.py:232-233,304-306 (M212 SEAM comments)
  destination: "M214 (origins-and-links)"
  reason_recorded: "M212 wires the host seam end-to-end (up-injected → gen → build_lines) but the CORS +
    Clerk-URL EMISSION stays localhost; the emission depends on M213's HTTPS/reverse-proxy origin shape
    (a proxied origin is https://<host> with no offset port), so emitting now would just be undone."
  partial_attempted: no  # deliberate wired-but-unemitted seam, pinned byte-identical by a test

- id: DEF-M212-02
  item: "clerkenstein.md dotted-FAPI-host constraint (why FAPI defaults to 127.0.0.1, not localhost)"
  origin_milestone: M212
  first_deferred_on: 2026-07-11
  last_seen_in: decisions.md KB-1; kb-fidelity-audit.md finding 3 / gap 1
  destination: "M214 (origins-and-links)"
  reason_recorded: "Incidental KB completeness gap. The constraint (assertValidPublishableKey rejects a
    dotless host) is enforced + thoroughly commented in code (up-injected.sh:581-586, inject.py:50);
    explaining why a MagicDNS origin validates for the pk is exactly tailscale-serve.md's remit."
  partial_attempted: no

Designed §Out decomposition (not scope erosion — the release's planned milestone split): TLS cert swap +
HTTPS reverse proxy + pk validation → M213; live cross-machine acceptance → M215; dev-path parity +
`/dev-up` flag → optional M216. Each is an `Out:` line in M212's overview and an `In:` line (with its own
overview/progress/spec-notes) in the destination milestone. All Fate 2.

## Repeat-Deferral Patterns
None. M212 opens v2.2; no item has been deferred in a prior milestone of this release.

## Fate-1 Investigation

### DEF-M212-01 — CORS + Clerk-URL emission
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate 2 — already owned by **M214**. M214's `overview.md` §Scope "In (config)" lists exactly
  this: extend `gen_injected_override.py:304-306` so `fe_origins` emits the HTTPS MagicDNS origin(s), and
  thread `$HOST` into runtime `CLERK_SIGN_IN_URL`/`WEB_APP_URL` (`:232-233`); its `progress.md` carries the
  matching checkbox. A complete landing now is genuinely infeasible: the emitted origin string is
  `https://<magicdns>` with NO offset port, a shape M213 (HTTPS/reverse-proxy, parallel-with M214) settles
  first — emitting `http://<host>:<port>` in M212 would be undone. Not a partial slice: the M212 seam is a
  full wire (host flows up-injected → gen → build_lines) pinned byte-identical by a test; M214 flips only the
  terminal emission.

### DEF-M212-02 — clerkenstein.md dotted-host constraint
- **Fate-1 (land now, complete) feasible:** no (by milestone design, not effort)
- **If no:** Fate 2/3 — homed to **M214**, which delivers the NEW `corpus/ops/demo/tailscale-serve.md`
  (the doc whose remit is "why a MagicDNS origin validates for the pk") AND lists `corpus/services/clerkenstein.md`
  under its "Updates". M212's overview explicitly states "No new doc lands here" — the knob's own documentation
  is a declared M214 deliverable. Not load-bearing for M212: the `127.0.0.1` default is preserved byte-identically
  and the constraint is enforced+commented in code. Landing a clerkenstein.md edit in M212 would pre-empt M214's
  doc pass and split the topic across two milestones.

## Recommendations
- DEF-M212-01 → **LAND-NEXT (Fate 2)** — confirmed covered by M214; no plan edit (already in its `In:`).
- DEF-M212-02 → **LAND-NEXT (Fate 2)** — confirmed covered by M214 (`tailscale-serve.md` + clerkenstein.md
  update); no plan edit.
- §Out decomposition → **LAND-NEXT (Fate 2)** across M213/M214/M215/(M216); confirmed in each target's overview.

## Applied Changes
None. All items are pre-existing, correctly-fated Fate-2 routings whose destinations were verified this pass
in `m214-origins-and-links/overview.md` + `progress.md` and `m213-auth-over-tailnet/overview.md`. No
`overview.md` edits were required (Fate 2 = confirm, don't edit). Recorded in M212 `decisions.md` already
(D-IMPL-3, KB-1).

## Blocking Items (require user decision)
None.
