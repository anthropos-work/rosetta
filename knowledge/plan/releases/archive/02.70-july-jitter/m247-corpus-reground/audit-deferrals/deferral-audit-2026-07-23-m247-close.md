---
title: "Deferral Audit — M247 corpus re-ground (close)"
date: 2026-07-23
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW

- No unresolved repeat-deferral. One **AGED_OUT** item (the rext-hygiene inert set whose
  originally-fated destination milestone M251 closed without landing it) is re-fated this pass with a
  **fresh decision** dated today — documented-inert standing note. All items carry a clear fate.
- Not GREEN only because this pass had to re-fate an aged-out item + a not-landed inherited hand-off;
  both resolve cleanly (no blocker, no cross-release escape hatch).

## Summary
- Total deferrals in scope: 4 (+ inherited M246/M251 ledger items covered by their own destinations)
- Single deferrals: 3
- Repeat / aged-out deferrals: 1 (DEF-M247-01 — destination-milestone-closed aging trigger)
- Chronic patterns flagged: 0 (the one repeat is a STRUCTURAL re-fate — doc-only M247 cannot touch rext
  files — not a time-pressure chronic pattern; now resolved to a documented-inert standing note)
- Blocking items: 0

## Deferral Inventory

### DEF-M247-01 — rext-hygiene inert items (M247 decisions D0: D-02/D-03/D-04/D-06)
- item: dormant `skillpath` INJECTED key in `gen_injected_override.py` ("4 injected"); `test_injection.py`
  skillpath-injected pins (+ residual skiller `_cfg`); `exposure_claim_guard.py:124` `skillpath:8095`
  fixture; `up-injected.sh:458` historical audit-prose comment ("…skiller, skillpath…").
- origin_milestone: M246 (drift ledger D-02..D-06) → routed at M247 D0
- first_deferred_on: 2026-07-23 (M246 close) → 2026-07-23 (M247 D0)
- last_seen_in: `m247-corpus-reground/decisions.md` § D0
- destination (as-recorded): "M251 / rext-hygiene" (Fate-3)
- reason_recorded: "M247 is DOC-ONLY (no rext) … these stay out and route to already-planned / better-fit siblings"
- partial_attempted: no
- **AGING TRIGGER:** destination milestone **M251 CLOSED (2026-07-23) without landing them** — M251's
  Deferred ledger owns only the optional verification.md anchor (→ M247) + the 8 live-gated failures
  (→ M254). The rext-hygiene set was never in M251's `In:` scope. Fresh decision required.

### DEF-M247-02 — ai-readiness demo-seeder fidelity + D-07 demopatch re-pin + 4 compute line-anchors (M247 D1)
- item: the demo-seeder fidelity counts (Seeding/CYCLE-STATE/FILLED-ness sections), the 4 bare-filename
  compute line-anchors, and the D-07 demopatch re-pin — the demo-side deltas of `ai-readiness.md`.
- origin_milestone: M247 (D1)
- first_deferred_on: 2026-07-23
- last_seen_in: `m247-corpus-reground/decisions.md` § D1
- destination: M250 (AI-readiness fidelity — which also `Delivers → ai-readiness.md`)
- reason_recorded: "the demo-seeder fidelity … M250's exit gate brings the demo to the 31-skill repertoire …"
- partial_attempted: no

### DEF-M247-03 — ops/demo spec-doc reconcile (M247 D-fate-2)
- item: `content-stories-spec.md`, `content-stories-routes.md`, `demopatch-spec.md`, `cockpit-spec.md`,
  `latency-budget.md`, `secrets-spec.md`, and the studio-desk parts of `frontend-tier.md`/`studio-desk.md`.
- origin_milestone: M247 (D-fate-2)
- first_deferred_on: 2026-07-23
- last_seen_in: `m247-corpus-reground/decisions.md` § D-fate-2
- destination: M248/M249/M250/M252/M253 (each updates its own spec docs) + the release-close consistency pass
- reason_recorded: "their not-yet-written deltas belong to M248/M249/M250/M252/M253 … cross-milestone
  consistency pass … is a Fate-2 release-close item"
- partial_attempted: no

### DEF-M247-04 — optional `verification.md` demo-stack-suite + run-unit-roster index anchor (inherited from M251)
- item: an optional `corpus/ops/verification.md` anchor indexing the demo-stack python suite + the run-unit roster.
- origin_milestone: M251 (deferred TO M247, "Fate 3-adjacent", lane-collision avoidance)
- first_deferred_on: 2026-07-23 (M251 close)
- last_seen_in: `m251-test-health/progress.md` § Completeness Ledger → Deferred (line 10)
- destination (as-recorded): M247
- reason_recorded: "M247 owns the corpus reground + `verification.md`; authoring here would collide with
  its concurrent lane … Not a blind area — the code it would index exists + is exercised."
- partial_attempted: no — **M247 did not land it** (not in M247's `Delivers`; `verification.md` untouched
  by the M247 diff).

## Repeat-Deferral Patterns

### AGED_OUT: "rext-hygiene inert items (dormant skillpath key + injection test fixtures + audit prose)"
- **First proposed:** M246 drift ledger (D-02..D-06), 2026-07-23
- **Deferred again:** M247 D0, 2026-07-23, reason: "doc-only milestone — rext files out of scope"
- **Ageing trigger:** destination milestone **M251 closed without landing** (+ deferred across ≥2 milestones)
- **Original reason (reference):** M247 is DOC-ONLY; these are rext files.
- **Required action:** fresh user/orchestrator fate-decision this pass — **SUPPLIED** (see Fate below).

No CHRONIC_DEFER: the repeat is structural (doc-only M247 literally cannot edit rext files — 0 of the four
files is tracked in the rosetta corpus repo), not a recurring time-pressure punt. It resolves to a
documented-inert standing note this pass.

## Fate-1 Investigation

### DEF-M247-01 — rext-hygiene inert items
- **Fate-1 (land now, complete) feasible:** no — all four files (`gen_injected_override.py`,
  `test_injection.py`, `exposure_claim_guard.py`, `up-injected.sh`) live in `rosetta-extensions`; **0 are
  tracked in the rosetta corpus repo** (`git ls-files` = 0 hits each). Doc-only M247 has no code surface here.
- **Materiality:** INERT. The load-bearing stale comment (`gen_injected_override.py:16` "3 subgraphs") was
  already corrected in **M246**. What remains is a dormant injection key that is never consumed + test
  fixtures that still pass + one historical audit-prose comment — **0 functional impact**; nothing breaks,
  nothing renders wrong, no test is red because of them.
- **If no:** the honest fate is a **documented-inert standing rext-hygiene note in M247's decisions.md**
  (the orchestrator's constraint forbids editing a sibling milestone's `overview.md`, so a formal Fate-3
  annotate is not taken). To be swept opportunistically by whichever rext milestone next edits those
  injection files (M249 owns `up-injected.sh`; M252 edits `gen_injected_override.py`) — no scope edit made
  to their plans; the sweep is optional cosmetic hygiene, not a promised deliverable.

### DEF-M247-02 — ai-readiness demo-seeder fidelity + D-07 + anchors
- **Fate-1 feasible:** no — the demo-seeder fidelity counts + the 31-skill repertoire + the track-keyed
  named sims + the evaluated-skills directus set-dress are **M250's exit-gate deliverables** (rext seeder +
  live render), out of doc-only M247's scope.
- **If no:** **Fate 2** — M250 already owns it (its overview `Delivers → ai-readiness.md`; its exit gate is
  the 31-skill fidelity). Confirmed covered; no plan edit needed.

### DEF-M247-03 — ops/demo spec-doc reconcile
- **Fate-1 feasible:** no — the deltas those spec docs need are the *outputs* of M248/M249/M250/M252/M253's
  code work, which has not landed yet. Reconciling them now would document behaviour that does not exist.
- **If no:** **Fate 2** — each code milestone updates its own spec doc in-lane; the cross-milestone
  consistency sweep is a release-close item. Confirmed covered; no plan edit needed.

### DEF-M247-04 — optional verification.md anchor
- **Fate-1 feasible:** technically yes (it is a corpus doc + M247 is doc-only) — but it is **out of M247's
  charter** (the consolidation reground: skillpath→app + 3-subgraph + fact sheets + ai-readiness), and it is
  a **test-health indexing** concern (M251's domain). M251 punted it purely to avoid a concurrent-lane
  collision, which is now moot.
- **Materiality:** explicitly **OPTIONAL** + **non-blind** ("the code it would index exists + is exercised").
- **If no (declined for M247):** **Fate 2** — fold into the **release-close consistency pass** (same bucket
  as DEF-M247-03). Not a blind area, no functional gap; a nice-to-have index anchor, not a promised surface.

## Recommendations
- **DEF-M247-01** → **KEEP-as-documented-inert-standing-note** (aged-out, fresh decision this pass; inert,
  0 functional impact, structurally un-landable in doc-only M247). Record in M247 `decisions.md`.
- **DEF-M247-02** → **LAND-NEXT (Fate 2)** — M250 owns it. Confirm, no plan edit.
- **DEF-M247-03** → **LAND-NEXT (Fate 2)** — code milestones + release-close pass own it. Confirm, no plan edit.
- **DEF-M247-04** → **LAND-NEXT (Fate 2)** — release-close consistency pass. Record the re-fate in M247
  `decisions.md` (optional + non-blind; not landed in M247 by charter).

## Applied Changes
- This report written.
- The fresh-decision records for DEF-M247-01 (documented-inert standing note) and DEF-M247-04 (re-fate to
  release-close) are applied to `m247-corpus-reground/decisions.md` in the close-milestone Phase 7 fix pass
  (D-audit entry), to avoid a double edit of the same file within one close.

## Blocking Items (require user decision)
- None. The one aged-out item (DEF-M247-01) carries a fresh, orchestrator-supplied fate (documented-inert
  standing note); no repeat-deferral is left unresolved; no cross-release escape hatch is invoked.
