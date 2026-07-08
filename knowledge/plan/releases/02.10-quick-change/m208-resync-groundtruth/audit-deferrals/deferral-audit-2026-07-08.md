---
title: "Deferral Audit — milestone:M208"
date: 2026-07-08
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; every item has a clear, in-release fate decision with its plan edit already applied.
- M208 is the **first** milestone of v2.1 — there are no prior-milestone inherited deferrals in scope.

## Summary
- Total deferrals in scope: **5**
- Single deferrals: **5**
- Repeat deferrals: **0**
- Chronic patterns flagged: **0**
- Aged-out items requiring fresh decision: **0** (see M25-D9 note below)

## Deferral Inventory

```yaml
- id: DEF-M208-01
  item: "clean cold-bring-up needs the extensions schema (pgvector + pg_trgm + resolvable gin_trgm_ops) bootstrapped + a PG-readiness wait before make migrate (M25-D9 class)"
  origin_milestone: M208
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:52 / spec-notes.md:81 (Finding 1)
  destination: "M211 (bring-up acceptance) — Fate 3; M209 Risk-2 cross-ref for the extensions.-qualified capture column list"
  reason_recorded: "surfaced on the clean-slate reset-db path; did NOT trivially fall out as a Fate-1 one-liner — it is a bring-up-tooling requirement, which M211 owns. The M208 overview explicitly permitted this ('resolve only if it falls out naturally; do not scope-creep if it doesn't')."
  partial_attempted: no

- id: DEF-M208-02
  item: "dev stack's hand-assembled .env lacks INVITATION_HMAC_SECRET (backend Exited(0) on the containerized cold run)"
  origin_milestone: M208
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:38 / spec-notes.md:99 (Finding 2)
  destination: "M211 / a /stack-secrets follow-up — Fate 2 (the sanctioned values-blind provisioner owns it)"
  reason_recorded: "a per-stack .env completeness gap, NOT merge-caused — the key is already documented (secrets-spec.md:111; a platform .env key; in secretdna.DemoGeneratedKeys). Dev's containerized backend is normally never run (native-worktree dev). A Fate-1 dev-value add was made to confirm the federation serves end-to-end, then reverted per the orchestrator directive; the .env is left in its original gap state."
  partial_attempted: no

- id: DEF-M208-03
  item: "corpus/services/backend.md still carries pre-merge consumer/dependency prose (skiller as a separate downstream service)"
  origin_milestone: M208
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:12 (KB-1) / kb-fidelity-audit.md:23
  destination: "M210 (corpus body-flip) — Fate 2"
  reason_recorded: "the pre-merge corpus staleness is the PREMISE of this release, fully planned. M208 pins the authoritative fact-sheet anchor M210 grades against; M210 flips the bodies. Not read as truth by M208's own implementation."
  partial_attempted: no

- id: DEF-M208-04
  item: "corpus/services/skiller.md still documents a live standalone service"
  origin_milestone: M208
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:15 (KB-2) / kb-fidelity-audit.md:30
  destination: "M210 (corpus body-flip; colleague's origin/docs/skiller-in-app-merge already drafts the stub) — Fate 2"
  reason_recorded: "same premise as DEF-M208-03. M208 added a minimal ⚠ stub banner pointing at the fact-sheet; the full body re-point is M210's chartered scope."
  partial_attempted: no

- id: DEF-M208-05
  item: "CLAUDE.md / corpus/services/graphql-wundergraph.md say '5 subgraphs' (actual = 4 at platform@origin/main)"
  origin_milestone: M208
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:16 (KB-3) / kb-fidelity-audit.md:36
  destination: "M210 (corpus body-flip) — Fate 2"
  reason_recorded: "M208 pins '4 subgraphs' in the fact-sheet; the doc-body corrections are M210's scope."
  partial_attempted: no
```

**Scope-boundary note (not a deferral):** `stack-dev/ant-academy` is 13 commits behind but is **not in
`repos.yml`** (a Clerk-free UI-tier native app, unrelated to the skiller merge) so `make pull` skips it — a
UI-tier concern owned by M211's bring-up acceptance, out of M208's merged-platform de-risk by design. Recorded
for completeness; carries no deferral authority.

## Repeat-Deferral Patterns
None. All 5 records are single (non-repeat), first-seen this release. No item appears in ≥2 milestones; no
CHRONIC_DEFER / DRIFT_DEFER pattern.

## Aging — M25-D9 (the one long-history item)
DEF-M208-01 traces to the standing-backlog item **M25-D9** (v1.3b lineage, tracked in `roadmap-vision.md`). It
is **not** a milestone-to-milestone repeat-deferral: it lived as an **unscheduled** standing-backlog item and
was never assigned to a milestone until v2.1 design (Phase-0 deferral audit) flagged it as an *opportunistic
Fate-1 candidate on the M208 re-sync migration path*. At M208 it **surfaced** but did not fall out as a trivial
Fate-1 (it is a bring-up-tooling requirement, not a one-line migrate tweak), so it received its **first concrete
milestone assignment** — Fate-3 into M211, with pinned scope in `m211-bringup-acceptance/overview.md` ("Pre-surfaced
bring-up requirement"). This is forward progress (an unscheduled item finally scheduled with concrete scope), not
scope erosion. The item carries a **fresh, today-dated (2026-07-08) fate decision** that the user has already
accepted (per the M208 orchestration). No stale authority is being rubber-stamped; no AGED_OUT blocker applies.

## Fate-1 Investigation

### DEF-M208-01 — extensions bootstrap + PG-readiness (M25-D9)
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate 3 → M211. It is a **bring-up-tooling** requirement (cold-bring-up ordering: bootstrap the
  `extensions` schema + a PG-readiness wait before `make migrate`), which the iterative bring-up-acceptance
  milestone M211 owns. M208 explicitly scoped live bring-up OUT → M211, and its overview permitted skipping the
  opportunistic item if it didn't fall out naturally. Landing a partial migrate-tweak now would be a disguised
  deferral the three-fate rule rejects. Plan edit already applied (M211 overview + M209 Risk-2 cross-ref).

### DEF-M208-02 — INVITATION_HMAC_SECRET dev .env gap
- **Fate-1 feasible:** no
- **If no:** Fate 2 → M211 / `/stack-secrets`. Provisioning a real secret is the job of the sanctioned
  values-blind provisioner, not an ad-hoc close edit; the orchestrator directive was "do NOT provision secrets
  here." The key is already fully documented in the corpus (no doc change owed). M211's bring-up acceptance runs
  `/stack-secrets`, which owns the dev-stack `.env` completeness.

### DEF-M208-03 / -04 / -05 — pre-merge corpus prose (KB-1/2/3)
- **Fate-1 feasible:** no (by design)
- **If no:** Fate 2 → M210. These are the **release premise** — v2.1 is chartered as M208 ground-truth → M209
  rext → **M210 corpus body-flip** → M211 bring-up. Flipping the doc bodies at M208 would collapse M210 into M208
  and land corpus changes ahead of the M209 tooling they must stay in lockstep with (roadmap Risk 4). M208's
  charter is the minimal authoritative fact-sheet anchor only, which it delivered.

## Recommendations
| Item | Fate | Destination | Verb |
|---|---|---|---|
| DEF-M208-01 | Fate 3 | M211 (+ M209 Risk-2 xref) | LAND-NEXT |
| DEF-M208-02 | Fate 2 | M211 / /stack-secrets | LAND-NEXT |
| DEF-M208-03 | Fate 2 | M210 | LAND-NEXT |
| DEF-M208-04 | Fate 2 | M210 | LAND-NEXT |
| DEF-M208-05 | Fate 2 | M210 | LAND-NEXT |

No LAND-NOW (none feasible as a complete Fate-1 within M208's charter), no DROP, no KEEP-DEFERRED-WITH-SIGNOFF
(zero cross-release / escape-hatch punts).

## Applied Changes
No new plan edits required by this audit — every route was already applied during the M208 build and is confirmed
intact:
- M211 `overview.md` — "Pre-surfaced bring-up requirement" section (DEF-M208-01 Fate-3 landing) — present.
- M209 `overview.md` — Risk-2 cross-ref for the `extensions.`-qualified capture column list — present.
- M208 `decisions.md` — Finding 1, Finding 2, and KB-1/2/3 routing records — present.
This report is the audit artifact; it records the confirmations, it does not re-open settled routes.

## Blocking Items (require user decision)
None. Zero repeat-deferrals, zero aged-out items with revoked authority, zero escape-hatch punts. The close may
proceed.
