---
title: "Deferral Audit — milestone M50 close"
date: 2026-06-30
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

No repeat-deferral pattern; no chronic pattern; no aged-out items. Every deferral in scope has a clear,
conscious, dated fate. The one item that *arrived* at M50 (the inherited AI-keys policy) is being RESOLVED at
this close (Fate-1 decision: documented-as-absent), not pushed forward again. Calling skill proceeds.

## Summary
- Total deferrals in scope: 4 (1 inherited-and-now-resolved, 1 release-level push-gated KEEP, 2 M50-own carry-forwards)
- Single deferrals: 4
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0 (all recorded 2026-06-29/30, < 3 months; none cycled forward without a decision)

## Deferral Inventory

```yaml
- id: DEF-M49-01            # INHERITED → ARRIVED at M50 → RESOLVED here
  item: "AI-provider-keys policy (which of OPENAI/ANTHROPIC/MISTRAL/ELEVENLABS/LIVEKIT become throwaway/sandbox demo values vs documented-as-absent)"
  origin_milestone: M49
  first_deferred_on: 2026-06-30
  last_seen_in: "m49-bringup-hardening/decisions.md:28 (Fate-2 → M50)"
  destination: "M50 (this milestone) — its consuming milestone"
  reason_recorded: "the bring-up itself needs no AI keys; the policy is M50's call (academy AI chat + content-gen seeding consume the keys)"
  partial_attempted: no
  m50_resolution: "LAND-NOW (Fate 1) — DECIDED: documented-as-absent. Demo content believability needs NO live AI; the AI-provider keys stay ABSENT from the demo secret source; AI-dependent surfaces (sim voice, academy AI chat, M45 batch-gen) are inert-by-design unless an operator supplies sandbox/throwaway keys. Values-blind (no real key provisioned). Recorded → M50 decisions.md D-AI-KEYS + corpus/ops/secrets-spec.md."

- id: DEF-M47-01            # release-level push-gated KEEP (re-affirmed, not re-scoped)
  item: "consumption-clone re-pin (stack-demo/rosetta-extensions @ tag + .agentspace/rext.tag bump on the box)"
  origin_milestone: M47
  first_deferred_on: 2026-06-29
  last_seen_in: "m47 progress.md:15 + decisions.md:55-57 ; M49 close audit ; M50 progress.md:19"
  destination: "release-level pending-pushes (origin push of main + rext tags); authoritative bump at M53 cold-rebuild acceptance"
  reason_recorded: "the re-pin needs the rext tag pushed to origin (tracked with the release's other pending pushes); authoritative bump is M53"
  partial_attempted: no
  m50_status: "push-gated KEEP — at M50 the target tag advances to fit-up-m50 (carries the iter-04/05/06 fixes). Same release-level pending-push class as M47/M49; NOT a repeat-defer. Authoritative bump = M53."

- id: DEF-M50-01            # M50-own carry-forward (USER-DECIDED Fate-2)
  item: "COLD reset-to-seed acceptance — the explicit exit_gate clause: a fresh /demo-up + both-vantage M42 sweeps on the strengthened manifest confirm (0,0) on COLD"
  origin_milestone: M50
  first_deferred_on: 2026-06-30
  last_seen_in: "M50 progress.md:17 ; iter-06/progress.md:76 ; metrics.json gate_met_cold:false"
  destination: "M53 (Cold-rebuild acceptance) — this release"
  reason_recorded: "the warm gate is MET both vantages on the strengthened manifest; all M50 seeders + fixes reproduce from the bring-up tooling on a fresh /demo-up. The COLD environment proof is the M53 cold-rebuild milestone's defining work. USER explicitly decided to defer the cold-environment proof to M53 (Fate-2)."
  partial_attempted: no
  m50_status: "LAND-NEXT (Fate 2) — M53 already owns the cold destroy-and-rebuild acceptance (roadmap M53 overview: 'rebuild from cold ... both-vantage coverage'). User-decided; no plan edit needed; no fresh sign-off required (decision already made by the user)."

- id: DEF-M50-02            # M50-own carry-forward (Fate-3 annotate → M51)
  item: "ant-academy course content + hero academy menu-link + non-anonymous academy session (F6)"
  origin_milestone: M50
  first_deferred_on: 2026-06-30
  last_seen_in: "M50 overview.md:42,56 (candidate fix surface + open question) ; spec-notes.md F6"
  destination: "M51 (AI-readiness showcase org) — this release"
  reason_recorded: "the academy content/wiring touches seeding/content; M51 is the AI-readiness showcase-org milestone that already touches seeding/content. The academy AI chat is documented-as-absent (gated by the AI-keys policy, DEF-M49-01) — NOT an M50 gate blocker (the M42 gate is MET both vantages without it)."
  partial_attempted: no
  m50_status: "LAND-NEXT (Fate 3) — annotate M51's overview In: list to pick up academy course-content + menu-link/non-anonymous-session. Recorded → M50 decisions.md."
```

(M48 D2 seed-strategy → M51 and M48 D3 repos.yml-fix → M49 were Fate-2 at M48 close; D3 LANDED in M49 §4
(explicit ensure-clones.sh academy clone) — both off the live ledger. M49 DEF-M49-01 is the only inherited
*live* item that reaches M50, and it RESOLVES here.)

## Repeat-Deferral Patterns
None. No item appears as a deferral in ≥ 2 distinct milestones with an unresolved forward-push.

- **DEF-M49-01 (AI-keys policy)** is *seen* in M49 (deferred → M50) and M50 (resolved), but that is the
  intended one-hop Fate-2 handoff reaching its destination and being decided — the opposite of a repeat-defer.
- **DEF-M47-01 (re-pin)** is *seen* in M47 / M49 / M50 docs, but it is a single **release-level push-gated**
  item tracked with the release's other pending origin pushes (local-only `main` + rext tags), not a feature
  being re-scoped milestone after milestone. The pin lands once the orchestrator pushes + at the M53 bump.
  Classified push-gated KEEP, consistent with the M47/M49 audits.

## Fate-1 Investigation

### DEF-M49-01 — AI-provider-keys policy
- **Fate-1 (land now, complete) feasible:** YES — and it lands at this close.
- **Landing:** the *complete* policy decision is "documented-as-absent": the demo's content believability does
  not need live AI; the AI-provider keys stay absent from the demo secret source; AI-dependent surfaces
  (sim voice, academy AI chat, M45 batch-gen) are inert-by-design unless an operator supplies their own
  sandbox/throwaway keys. No real key is provisioned (values-blind). This IS the whole policy — there is no
  residual slice. Recorded as a decision in M50 `decisions.md` and documented in `corpus/ops/secrets-spec.md`.

### DEF-M47-01 — consumption-clone re-pin
- **Fate-1 (land now, complete) feasible:** no (release-gated, not milestone-gated).
- **If no:** push-gated KEEP. The authoritative bump depends on the orchestrator pushing the rext tags +
  `main` to origin (a release-level pending action the whole v1.10b release shares); the authoritative re-pin
  is the M53 cold-rebuild acceptance (`bump .agentspace/rext.tag`). At M50 the target tag advances to
  `fit-up-m50`; landing the box-level pin now would re-pin to a tag not yet on origin. Tracked with the
  release's other pending pushes; not a backlog entry, not a repeat-defer.

### DEF-M50-01 — COLD reset-to-seed acceptance
- **Fate-1 (land now, complete) feasible:** no — by user decision + milestone design.
- **If no:** Fate 2 — **M53 already owns it.** The M50 warm gate is MET on both vantages on the strengthened
  manifest; all M50 seeders + fixes are baked into the bring-up tooling and reproduce on a fresh `/demo-up`.
  The COLD destroy-and-rebuild proof is M53's defining work (roadmap M53: "rebuild from cold ... both-vantage
  coverage"). The user explicitly decided to defer the cold-environment proof to M53. No fresh sign-off
  required (the decision is already made); no plan edit needed (M53 already lists it).

### DEF-M50-02 — academy content + menu-link/non-anonymous-session (F6)
- **Fate-1 (land now, complete) feasible:** no — belongs to the seeding/content milestone that follows.
- **If no:** Fate 3 — **annotate M51.** M51 (AI-readiness showcase org) already touches seeding/content and is
  the natural home for academy course-content + the menu-link/non-anonymous-session wiring. The academy AI
  chat is documented-as-absent (gated by DEF-M49-01) and is NOT an M50 gate blocker — the M42 gate is MET both
  vantages without it. M51's `overview.md` In: list is annotated to pick it up.

## Recommendations
- **DEF-M49-01 → LAND-NOW (Fate 1):** decided at M50 close — documented-as-absent. Recorded in M50
  `decisions.md` + `corpus/ops/secrets-spec.md`. Off the ledger.
- **DEF-M47-01 → KEEP (push-gated, release-level):** no fresh sign-off needed; in-release pending-push, target
  tag advances to `fit-up-m50`; authoritative re-pin at M53.
- **DEF-M50-01 → LAND-NEXT (Fate 2):** M53 owns the COLD acceptance. User-decided; no plan edit, no fresh
  sign-off.
- **DEF-M50-02 → LAND-NEXT (Fate 3):** annotate M51's `overview.md` In: list. Recorded in M50 `decisions.md`.

## Applied Changes
The audit confirms the fates; the close-milestone Phase 7 lands the decision records + the M51 annotation:
- M50 `decisions.md` → AI-keys policy (documented-as-absent, Fate-1 resolution of DEF-M49-01); COLD→M53
  (Fate-2, user-decided); academy/F6→M51 (Fate-3) handoff.
- M51 `overview.md` In: list → academy course-content + menu-link/non-anonymous-session (Fate-3 annotation).
- `corpus/ops/secrets-spec.md` → the AI-keys documented-as-absent policy.
This audit records the GREEN verdict for the M50 close.

## Blocking Items (require user decision)
None.
