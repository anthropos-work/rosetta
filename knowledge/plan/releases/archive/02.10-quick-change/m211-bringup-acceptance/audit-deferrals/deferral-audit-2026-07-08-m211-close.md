---
title: "Deferral Audit — M211 close (milestone scope)"
date: 2026-07-08
scope: milestone
invoked-by: close-milestone
---

## Verdict
**GREEN**

- No unresolved repeat-deferral; no chronic pattern.
- Every item has a clear, fresh (today-dated) fate decision with a firm owner.
- The one multi-milestone item (TEST-1, rext README test-count drift) is **aged-out → re-fated fresh** this pass:
  Fate-2, structurally owned by `/developer-kit:close-release`'s rext roll (imminent next step). Structural reason,
  not chronic avoidance.

## Summary
- Total deferrals in scope (M211 + inherited v2.1): **7**
- Resolved this milestone: 1 (DEF-M208-01/M25-D9)
- Single deferrals with firm destinations: 4
- Repeat deferrals unresolved: **0**
- Chronic patterns flagged: **0**
- Aged-out re-fated fresh: 1 (TEST-1)

## Deferral Inventory

- **DEF-M208-01 / M25-D9** — dev cold DB-init (extensions-schema bootstrap + PG-readiness before migrate).
  origin M208 → Fate-3 → M211. `reason_recorded`: "clean-bring-up extensions bootstrap; un-editable platform
  `make migrate` doesn't create `extensions`". **RESOLVED in M211 iter-17** — `dev-stack/migrate-dev.sh`
  (mirror of `migrate-demo.sh`), cold-verified on a faithful non-destructive throwaway (extensions +
  `gin_trgm_ops` + 89 public tables + `cms.vector` + casbin, 0 skiller). partial_attempted: no.
- **DEF-M208-02** — `INVITATION_HMAC_SECRET` dev `.env` completeness gap. origin M208 → Fate-2 → M211 /
  `/stack-secrets`. `reason_recorded`: "not merge-caused; owned by the secret-provisioning tooling". Covered:
  the demo secret-DNA already carries `INVITATION_HMAC_SECRET` as a critical values-blind gene (v1.10b M49 #4);
  M211's cold `/demo-up` GREEN exercised it. `/stack-secrets` is the standing owner for the dev target too.
- **TEST-1** — rext `stack-seeding/README` test-count drift (quotes 496 / 8 pkgs; actual ~788 / 13). origin
  M41 (v1.10) → re-noted M209 close → routed v2.1 rext roll / close-release. `reason_recorded`: "rext frozen
  per-milestone at its tag; reconcile at the rext code-of-record roll (close-release)". **AGED_OUT** (mentioned
  across ≥2 milestones).
- **CAVEAT-1** — literal full clean-box destructive all-services `/dev-up` + verify-net not executed on this box.
  origin M211 iter-17. `reason_recorded`: "this box is committed to the user's native-app content-line dev
  (`docker-compose.override.yml` → `backend:host-gateway` + `app-01.10-content-line` worktree); a full
  destructive `/dev-up` would clobber it and can't go green without a v2.1 native backend. The dev-specific gate
  delta (M25-D9 DB-init) was cold-verified on a faithful throwaway + a live docker harness."
- **CAVEAT-2** — pre-existing `dev-stack` CLI unit-test failures (`test_dev_stack.py` ~13, from an incomplete
  local `.agentspace/secrets` source tripping the secret pre-flight). origin: pre-existing environment/secrets
  drift, unrelated to M211. NOT a roadmap deferral (environment condition — the standalone `migrate-dev.sh`
  static+live tests are green regardless).
- **PT-TODO** — 1 declared in-manifest Playthrough TODO (the assign-WRITE half). origin v2.0 (shipped state
  10/11) → Fate-2 reserved manager-write tier (Playthroughs futures M205–M207, `roadmap-vision.md`). Inherited,
  unchanged by v2.1; M211 keeps the suite 10/11 GREEN.
- **PUSH-KEEP** — origin push of `main` + tags + the box-level `.agentspace/rext.tag` re-pin. Administrative,
  user-owned gate; the box-level re-pin is close-release's job. NOT a deferral.

_(Out-of-scope for this milestone audit: the cross-release standing backlog DEF-M10-01, DEF-M21-01, M314b —
parked in `roadmap-vision.md` with prior sign-off, untouched by v2.1.)_

## Repeat-Deferral Patterns

### AGED_OUT: "rext stack-seeding/README test-count drift" (TEST-1)
- **First noted:** M41 (v1.10), reconciled then; drifted again.
- **Re-noted:** M209 close (2026-07-08), nice-to-have → routed v2.1 rext roll.
- **Current destination:** `/developer-kit:close-release` rext roll (the v2.1 rext code-of-record).
- **Ageing trigger:** mentioned across ≥2 milestones.
- **Pattern:** NOT `CHRONIC_DEFER` — the reason is **structural** (the rext repo is frozen per-milestone at a
  pinned tag; the README test-count reconciliation can only land when the release rolls the rext code-of-record,
  which is close-release's dedicated operation), not "not enough time". No milestone can structurally do it.

No other item appears in ≥2 milestones within v2.1.

## Fate-1 Investigation

- **DEF-M208-01/M25-D9** — Fate-1 feasible: **yes, and DONE** (landed in M211 iter-17). No action.
- **DEF-M208-02** — Fate-1 not needed: already covered by the standing `/stack-secrets` tooling (Fate-2
  confirmed). The dev `.env` is provisioned by `/stack-secrets` from the secret source; the gene exists.
- **TEST-1** — Fate-1 (land now) **infeasible from a milestone close**: the rext repo is frozen at
  `quick-change-m211` for M211; the README reconciliation belongs to the rext code-of-record roll. Fate-2 →
  close-release (imminent). Fresh dated decision made this pass.
- **CAVEAT-1** — Fate-1 (run it now) **deliberately not done**: a full destructive `/dev-up` would clobber the
  user's committed native content-line dev setup and cannot go green without a v2.1 native backend. The gate does
  NOT require it (the sole dev-specific delta was proven). Belt-and-suspenders backlog note (clean box).
- **CAVEAT-2** — nothing to land (environment condition; needs a complete local secrets source, which is a box
  provisioning matter, not tooling scope).
- **PT-TODO** — Fate-2 confirmed (reserved Playthroughs futures); inherited from v2.0, unchanged.

## Recommendations

1. **DEF-M208-01/M25-D9 → LAND-NOW (satisfied).** Mark resolved in `roadmap-vision.md` (migrate-dev.sh, M211).
2. **DEF-M208-02 → LAND-NEXT / confirmed-covered** by `/stack-secrets` (standing tooling owner). No plan edit.
3. **TEST-1 → LAND-NEXT (Fate-2)** → `/developer-kit:close-release` rext roll. Fresh dated decision recorded.
4. **CAVEAT-1 → KEEP (belt-and-suspenders backlog)** — accepted gate-interpretation; add a clean-box full
   `/dev-up` note to `roadmap-vision.md` unscheduled backlog. Not an escape-hatch scope-break (gate met).
5. **CAVEAT-2 → note as pre-existing environment/secrets-source drift.** Not roadmap scope.
6. **PT-TODO → LAND-NEXT / confirmed-covered** (reserved manager-write tier). Inherited; no v2.1 action.

## Applied Changes
- Fresh fate decisions recorded in M211 `decisions.md` (audit re-fate: DEF-M208-01 resolved, DEF-M208-02
  confirmed-covered, TEST-1 Fate-2 → close-release, CAVEAT-1 belt-and-suspenders backlog, CAVEAT-2 env-drift note).
- `roadmap-vision.md`: mark M25-D9 resolved (migrate-dev.sh) + add the CAVEAT-1 clean-box full-`/dev-up` backlog
  note — landed with the Phase 10 roadmap update commit (recorded here as audit-driven).

## Blocking Items (require user decision)
**None.** No unresolved repeat-deferral; the one aged item (TEST-1) is re-fated fresh with a firm imminent owner
(close-release rext roll); CAVEAT-1 is an accepted, gate-met belt-and-suspenders note (not an escape-hatch
scope-break). Gate decision: **GREEN → SEVERITY=clear.**
