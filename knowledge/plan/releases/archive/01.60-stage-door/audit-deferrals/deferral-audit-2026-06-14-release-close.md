---
title: "Deferral Audit — v1.6 \"stage door\" (close-release Phase 1b)"
date: 2026-06-14
scope: release
invoked-by: close-release
release: 01.60-stage-door
milestones: [M27, M28, M29, M30]
---

## Verdict
GREEN

- No repeat deferrals: no v1.6 item was deferred across ≥2 milestones. The single within-release item
  (DEF-M27-02 profile-tag) resolved at M28 close (Fate-2 discharged, conditional never triggered); the other
  within-release item (DEF-M27-01 encrypted-zip) was DROPPED at M27 close as a documented v1 scope boundary.
- No aged-out items: the 3 inherited release-level backlog items were last re-decided 2026-06-14 at the v1.5
  `/close-release` re-audit (same calendar day as this audit). None of the four ageing triggers fires — see
  the Aging Check.
- No chronic patterns: no item carries a "not enough time" or identical-reason re-defer chain inside v1.6.
- Every item has a clear fate decision.

## Summary
- Total deferrals in scope: 3 (all inherited release-level backlog; orthogonal to secret provisioning)
- Within-release v1.6 items: 2, both already resolved (DEF-M27-01 DROPPED, DEF-M27-02 Fate-2 discharged)
- New deferrals introduced by v1.6: 0
- Single deferrals: 0 live (the 2 within-release are resolved, not pending)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- Waived-class secret-DNA genes: counted separately (waived-not-deferred — gate denominator excludes by
  design; NOT in the deferral count)

## Deferral Inventory

### Within-release v1.6 items — both RESOLVED (carried for completeness, not pending)
```yaml
- id: DEF-M27-01
  item: "Encrypted-zip (age/gpg) source ingestion support"
  origin_milestone: M27
  status: DROPPED at M27 close
  reason: "complete landing pulls a crypto dependency + key-management surface, contradicting the deliberate
           stdlib-only design (M27-D3.1) whose point is a trivially-auditable values-blind surface; no consumer
           needs it (plain dir/zip fully cover the M30 field-bake); a stub is forbidden by the Fate-1 no-partial
           rule. Recorded as a closed scope boundary in M27-D3.5 — NOT a live ledger entry."
  partial_attempted: no

- id: DEF-M27-02
  item: "Per-gene profile tag (vs default-graphql-profile scoping)"
  origin_milestone: M27
  status: Fate-2 DISCHARGED at M28 close
  reason: "conditional was 'M28 revisits IFF it wires non-default-profile bring-ups'; M28 wired the pre-flight on
           the default graphql profile only, so the condition never triggered. Profile scoping reused unchanged
           (default graphql + the waived-profile-gated class). The per-gene-tag variant remains a consciously-
           rejected, documented v1 option (M27-D3.4)."
  partial_attempted: no
```

### Inherited release-level backlog (orthogonal to secret provisioning — confirmed untouched by all of M27-M30)
```yaml
- id: DEF-M10-01
  item: "Cloud SnapshotStore backend + S3 media blob bytes"
  origin_milestone: M10 (v1.2)
  first_deferred_on: 2026-06-06 (approx, v1.2)
  last_fresh_decision: 2026-06-14 (v1.5 close-release re-audit; user-facing sting removed at v1.5 design 2026-06-11)
  destination: "roadmap-vision.md Unscheduled backlog (gated on eu-west-1 S3 read access actually landing)"
  reason_recorded: "asset plane stays on prod public links → demos show real images without blob-byte work;
                    real blob mirroring + cloud store gated on S3 access (verified not wired)"
  partial_attempted: no

- id: DEF-M21-01
  item: "replayCmd conn-seam hermetic test"
  origin_milestone: M21 (v1.5)
  first_deferred_on: 2026-06-11 (v1.5)
  last_fresh_decision: 2026-06-14 (v1.5 close-release — landed to roadmap-vision backlog so it survives the merge)
  destination: "roadmap-vision.md Unscheduled backlog — pick up in a future stack-snapshot build iter when the
                replay path is next opened"
  reason_recorded: "hermetic replayCmd-wiring test needs an injectable connector seam (>50 lines, touches the
                    load-bearing replay path)"
  partial_attempted: no

- id: M25-D9
  item: "dev-N taxonomy replay rc=4 (\"target schema empty\") migrate-ordering nuance"
  origin_milestone: M25 (v1.5)
  first_deferred_on: 2026-06-13 (v1.5 M25 field-bake)
  last_fresh_decision: 2026-06-14 (v1.5 close-release re-audit)
  destination: "roadmap-vision.md Unscheduled backlog — dev migrate-ordering follow-up"
  reason_recorded: "pre-existing dev-stack migrate-ordering nuance on opt-in dev-N≥1 --local-content stacks;
                    non-fatal, orthogonal to the content-serve path (done-bar DB-2 is GREEN)"
  partial_attempted: no
```

### Waived secret-DNA classes (waived-not-deferred — NOT in the deferral count)
Excluded from the coverage-gate denominator by design; each is a documented-correct reason a secret is not
provisioned on a local stack. Verified live in `secretdna/secret-dna.json`:
- `waived-aws-mount` — AWS creds via host mount, not a provisioned `.env` value.
- `waived-profile-gated` — `BREVO_KEY` (messenger profile, off by default).
- `waived-optional` — `BUNNY_STREAM_API_KEY`, `BUNNY_CDN_TOKEN_KEY`, `GCLOUD_SERVICE_ACCOUNT`,
  `TAILSCALE_AUTH_KEY`, `YOUTUBE_API_KEY` (absent from live envs / example-only).
- `waived-config` — `sentinel/DB_CONNECTION` (compose-injected hardcoded `environment:` DSN that always
  overrides `env_file`; in-network password-less wiring config, not a provisioned secret — reclassified by the
  M30 field-bake with a regression test).

These are scope-correct waivers, not punted work. No fate decision applies.

## Repeat-Deferral Patterns
None. No v1.6 item has been deferred across ≥2 milestones:
- DEF-M27-02 resolved (discharged) at M28 rather than repeating.
- DEF-M27-01 was dropped at M27 rather than re-deferred.
- The 3 inherited items were carried **verbatim** through M27→M30 (re-signed once at the v1.5 close, then noted
  as KEEP at each v1.6 milestone close) — this is single-decision carry, not a repeat-defer chain. None
  changed destination or reason inside v1.6.

No `CHRONIC_DEFER` (no repeated "not enough time" reason) and no `DRIFT_DEFER` (no re-scoping inside v1.6).

> Cross-release longevity note (informational, non-blocking): DEF-M10-01 has lived in the backlog since v1.2.
> It is not a v1.6 repeat-defer (v1.6 never re-deferred it; it was re-signed fresh at the v1.5 design with its
> user-facing sting removed) and it does not trip an aging trigger here. Flagged only so the next
> `/developer-kit:design-roadmap` Phase 0 sees it as a long-lived, S3-access-gated vision item — not a v1.6 close blocker.

## Aging Check
Today = 2026-06-14. Triggers: (a) ≥2 milestones, (b) ≥3 months, (c) destination milestone closed without
landing, (d) area touched by a later milestone.

- **DEF-M27-01 / DEF-M27-02** — resolved this release (DROPPED / discharged); nothing to age.
- **DEF-M10-01 / DEF-M21-01 / M25-D9** — last *fresh* decision 2026-06-14 (v1.5 close, same calendar day):
  - (a) NOT fired — 0 v1.6 milestones re-deferred them; carried verbatim on one v1.5 decision.
  - (b) NOT fired — elapsed since last fresh decision = 0 days (≪ 3 months). (Original M10 origin is irrelevant;
    the policy measures from the last fresh re-decision, which was yesterday's same-day v1.5 close.)
  - (c) NOT fired — destination is the unscheduled backlog (no committed milestone), and no v1.6 milestone was
    chartered to land them.
  - (d) NOT fired — the entire v1.6 diff is the new `stack-secrets/` section + rosetta docs/skills/planning
    (verified: ext diff `fbb8783..29c922b` is `stack-secrets/` only; rosetta diff `main...release/01.60` touches
    no snapshot/replay/dev-taxonomy file). None of the three feature areas was touched.

No item is AGED_OUT.

## Fate-1 Investigation

### DEF-M10-01 — Cloud SnapshotStore + S3 media blob bytes → KEEP-DEFERRED (escape hatch)
- **Fate-1 (land now, complete) feasible:** no.
- **Why not:** a complete landing requires a cloud `SnapshotStore` backend + real blob-byte mirroring, both
  hard-gated on **eu-west-1 S3 read access that is verified not wired**. v1.6 is the secret-provisioning release;
  this is orthogonal (snapshot/asset plane). No partial slice is a valid Fate-1 outcome.
- **Fate:** escape hatch (cross-release) — the user-facing sting is already gone (asset plane on prod public
  links → real images). Destination unchanged: `roadmap-vision.md` Unscheduled backlog.
- **Recommendation: KEEP-DEFERRED-WITH-SIGNOFF.** Re-signed 2026-06-14 at the v1.5 close (same day); carry
  unchanged. Already present in `roadmap-vision.md` with a current reason.

### DEF-M21-01 — replayCmd conn-seam hermetic test → KEEP-DEFERRED (escape hatch)
- **Fate-1 feasible:** no — needs an injectable connector seam (>50 lines) on the load-bearing replay path;
  out of v1.6's secret-provisioning scope and not incidentally unblocked by any v1.6 work.
- **Fate:** escape hatch (cross-release). Destination: `roadmap-vision.md` Unscheduled backlog — pick up in a
  future `stack-snapshot` build iter when the replay path is next opened.
- **Recommendation: KEEP-DEFERRED-WITH-SIGNOFF.** Landed to backlog at the v1.5 close (2026-06-14) so it
  survives the release merge; carry unchanged.

### M25-D9 — dev-N taxonomy replay rc=4 migrate-ordering → KEEP-DEFERRED (escape hatch)
- **Fate-1 feasible:** no — a pre-existing dev-stack migrate-ordering nuance on opt-in `dev-N≥1 --local-content`
  stacks; non-fatal, orthogonal to v1.6 secret provisioning, not unblocked by any v1.6 work.
- **Fate:** escape hatch (cross-release). Destination: `roadmap-vision.md` Unscheduled backlog — dev
  migrate-ordering follow-up.
- **Recommendation: KEEP-DEFERRED-WITH-SIGNOFF.** Carry unchanged from the v1.5 close.

### DEF-M27-01 / DEF-M27-02 — within-release, already resolved
- DEF-M27-01: DROPPED at M27 close (documented v1 boundary in M27-D3.5; no live destination). No re-fate needed.
- DEF-M27-02: Fate-2 discharged at M28 close (conditional never triggered). No re-fate needed.

## Recommendations
| Item | Fate verdict | Action |
|---|---|---|
| DEF-M27-01 encrypted-zip | DROP (done at M27 close) | none — closed boundary |
| DEF-M27-02 per-gene profile tag | LAND-NEXT→discharged (Fate-2, done at M28) | none — discharged |
| DEF-M10-01 cloud store / S3 blob bytes | KEEP-DEFERRED-WITH-SIGNOFF (escape hatch) | carry; already in roadmap-vision |
| DEF-M21-01 replayCmd hermetic test | KEEP-DEFERRED-WITH-SIGNOFF (escape hatch) | carry; already in roadmap-vision |
| M25-D9 dev taxonomy rc=4 | KEEP-DEFERRED-WITH-SIGNOFF (escape hatch) | carry; already in roadmap-vision |
| Waived secret-DNA classes | N/A — waived-not-deferred | none — gate denominator excludes by design |

## Applied Changes
- No plan edits required. The 2 within-release items are already resolved (DEF-M27-01 dropped at M27 close;
  DEF-M27-02 Fate-2 discharged at M28 close). The 3 inherited backlog items are unchanged — already recorded in
  `roadmap-vision.md` with current reasons and re-signed fresh at the v1.5 close (2026-06-14, same calendar day),
  so a re-fate at the v1.6 close would be churn, not diligence. They were verified orthogonal to the entire v1.6
  diff (no aging trigger fired).
- This audit report is the audit trail (the v1.6 release-level deferral re-audit).

## Blocking Items (require user decision)
None.
