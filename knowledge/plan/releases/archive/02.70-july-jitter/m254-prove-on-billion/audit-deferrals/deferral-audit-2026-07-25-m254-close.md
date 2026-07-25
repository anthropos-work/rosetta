---
title: "Deferral Audit — M254 prove-on-billion (close, milestone scope)"
date: 2026-07-25
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW

- Single deferrals only; all have clear, coordinator-accepted fate decisions.
- **No repeat-deferrals** (every M254-originated item was first-surfaced 2026-07-25; nothing deferred across ≥2 milestones).
- **No aged-out items** (nothing ≥3 months old; no destination-milestone-closed-without-landing).
- **No chronic patterns.**
- One item re-fated to the STRONGER outcome: `FIX-M254-h-patch-inventory-drift` moves carry-forward's Fate-3 → **LAND-NOW (Fate-1)** — a RED-at-HEAD test must not ship. Flows into close Phase 7.

## Summary
- Total deferrals in scope: 7 (M254-originated) + 5 inherited-and-discharged
- Single deferrals: 7
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- Blocking items: 0

## Inherited carries routed TO M254 — DISCHARGED by the live gate proof
These were Fate-2 routed to M254 during prior-milestone closes; M254's cold-reset-to-seed live proof on `billion` (gate a–h all MET) is exactly the confirmation they were routed for. All RESOLVED, none open:
- **CARRY-M248-01** (content-stories manager CTA fresh-seed re-confirm) → M254 gate (b)+(h). **DISCHARGED** — gate (b) 45/45 landable ALL LANDED (+4 voice presence-only, disposition below).
- **CARRY-M250-01** (AI-readiness ~150-page manager sweep live) → M254 gate (d)+(h). **DISCHARGED** — gate (d) MET both vantages, 8/8 sections + 3 drift-fixes, 0 escapes.
- **CARRY-M252-01** (5 academy autoverify stub failures, M245 root cause) → M254 gate (g). **DISCHARGED into** `FIX-M254-g-testhealth` (the consolidated host-sensitive test-health membership; see below).
- **CARRY-M252-02** (~10-min async live builder-generate RUN) → M254 gate (e)+(h). **DISCHARGED** — gate (e) MET (studio builders green after v0.152.1 re-tune; 18/18 Playthroughs).
- **CARRY-M253-01** (fully-green COLD-p95 studio first-paint on billion) → M254 gate (f). **DISCHARGED** — gate (f) MET on app-side paint (p50 637–726 ms < 1 s); p95 dispositioned environmental (below).

## Deferral Inventory (M254-originated) + Fate-1 Investigation + Recommendations

### DEF-M254-01 — FIX-M254-h-patch-inventory-drift (demopatch inventory fence RED at HEAD)
- **Origin:** M254 (surfaced by the final harden pass, 2026-07-25). **Root cause:** M253 (`b8969c0`) added `studio-desk-shell-first-paint` + `studio-desk-no-thirdparty` without bumping the fence — latent RED since the M253 tip; NOT an M254 change.
- **partial_attempted:** no.
- **Fate-1 (land now, complete) feasible:** **YES.** The fix is trivial + mechanical, precisely specified, empirically confirmed against the real tree: `EXPECTED_TOTAL 21→23`, `EXPECTED_BY_REPO["studio-desk"] 3→5` in `demo-stack/tests/test_patch_inventory.py` (rext) + the `corpus/ops/demo/demopatch-spec.md §5` inventory table (studio-desk 3→5, total 21→23). Test empirically RED at HEAD: `23 != 21` and `studio-desk:5 != 3`.
- **Recommendation:** **LAND-NOW (Fate 1).** A RED-at-HEAD must not ship. Overrides carry-forward.md's Fate-3 (which pre-dated this close's mandate to land it). Rung-zero: rext commit + new tag + `git push --tags` origin.

### DEF-M254-02 — (f)-FCP-p95 environmental disposition
- **Origin:** M254 iter-06 (D-iter10-2), coordinator-approved. **partial_attempted:** n/a (disposition, not deferral).
- **Fate-1 feasible:** YES — recorded as a disposition. Studio first-paint shell holds on billion (p50 637–726 ms < 1 s, M253 fix holds); the p95 outliers (1443/2014/4943 ms, `reachedShell` always true) are tailnet network-RTT jitter, per `latency-budget.md`'s "state the environment" + the (b) precedent. Gate (f) MET on app-side paint.
- **Recommendation:** **LAND-NOW (Fate 1 — disposition recorded).** No further work; the environmental acceptance IS the complete outcome.

### DEF-M254-03 — FIX-M254-c-academy-durable
- **Origin:** M254 iter-05/iter-08 (D-iter10-3), coordinator-approved. A FRESH demo renders Back-to-Cockpit on all 4 apps (gate (c) MET, presenter path); the native academy dev-server reverts the patch only on a long-running demo (a durability edge, not the presenter path).
- **Fate-1 feasible:** NO — the fix belongs to the rext `stack-injection`/`ant-academy.sh` reapply lifecycle (make the patch durable across the native dev-server's lifecycle), which is a fix-iter's tooling work outside this proof milestone's scope. Gate (c) itself is MET.
- **Recommendation:** **LAND-NEXT (Fate 3 → academy-durable follow-up).** Carry-forward; 0 platform edits. Coordinator-approved.

### DEF-M254-04 — FIX-M254-g-testhealth (6 host-sensitive demo-stack test-health tests)
- **Origin:** M254 iter-04/iter-07 (D-iter10-4), coordinator-approved (subsumes the discharged CARRY-M252-01). 2/8-class fixed + verified live (nvm/node host-robustness, rext `dfdd9bc`); the 6 remaining are chronic host-sensitive test-HARNESS issues (intra-run `:23077` port-leak + M245 reconcile-message drift; next.config sha re-pin; 2 mutation-meta; overlay-127) with **0 demo-runtime impact** (real academy serves 200).
- **Fate-1 feasible:** NO — per-test disentangle + fix in rext `demo-stack` is a fix-iter's work; it does not gate the live proof (gate (g) is a test-harness membership, not a demo defect).
- **Recommendation:** **LAND-NEXT (Fate 3 → carry-forward).** 0 platform edits. Coordinator-approved.

### DEF-M254-05 — (b)-voice manager_presence_only (content-stories denominator honesty)
- **Origin:** M254 iter-04 (traces to the accepted vision item DEF-M10-01 / DEF-M240-01 — the Bunny-keyless demo box → voice presence-only, data-controller-dispositioned). Gate (b) = 45/45 landable ALL LANDED + 4 voice presence-only (2 player + 2 manager, symmetric extension of DEF-M240-01).
- **Fate-1 feasible:** NO — the underlying voice-media blocker is an accepted external constraint (Bunny keys, vision-level DEF-M10-01). The M254 follow-up is a small denominator/flag honesty fix (rext `manager_presence_only` flag + content-denominator 47→45 + re-seed), which is tooling work for a fix-iter, not the live proof. The surface renders; no fabricated CTA.
- **Recommendation:** **LAND-NEXT (Fate 3 → follow-up).** Not a repeat/chronic — the media blocker is the accepted vision item; this is a bounded new honesty-fix routed forward.

### DEF-M254-06 — verify.sh stale skillpath default
- **Origin:** M254 iter-04 note. verify.sh's DEFAULT service list (when `--services` omitted) still lists the decommissioned `skillpath`, so a standalone autoverify without `--services` probes a non-existent service (HTTP 000000). NOT gate-blocking (bring-up passes explicit `--services` without skillpath).
- **Fate-1 feasible:** NO (cleanly) — a one-line rext default-list edit, but it belongs to a rext `stack-verify` tidy-iter and does not affect the M254 proof (every M254 autoverify used explicit `--services`). Routing it forward keeps this close scoped to the proof + the RED-at-HEAD.
- **Recommendation:** **LAND-NEXT (Fate 3 → follow-up).** 0 platform edits.

### DEF-M254-07 — studio-desk billion re-pin note (→ july-jitter-m254-studio-pt-retune @ 4f1409e)
- **Origin:** M254 iter-10 note. The iter-10 studio-builder + networkidle fix is **test-only** (runs from the local rext clone; billion serves the unchanged app), so it needs NO demo re-pin for the current proof — a future full cold reset-to-seed should pin billion's rext to 4f1409e (or later) so the whole toolchain is current.
- **Fate-1 feasible:** n/a — a forward-guidance NOTE, not deferred work. No landing needed for M254 (the proof is complete at billion's `dfdd9bc` pin).
- **Recommendation:** **Note (Fate 3).** Carried as guidance for the next cold re-prove.

## Repeat-Deferral Patterns
None. Every M254-originated item was first-surfaced 2026-07-25; the inherited carries were all discharged by the gate proof, not re-deferred.

## Applied Changes
- `FIX-M254-h-patch-inventory-drift` re-fated Fate-3 → **LAND-NOW (Fate 1)** — flows into close Phase 7 (rext test-fence bump + tag + push origin; corpus `demopatch-spec.md §5` table). decisions.md `D-harden-1` + carry-forward.md updated at Phase 7 to record it LANDED.
- The other 6 items confirmed as coordinator-approved Fate-1 dispositions / Fate-3 follow-ups; no plan edits required (all already recorded in decisions.md D-iter10-2/3/4 + carry-forward.md).

## Blocking Items (require user decision)
None. No repeat-deferrals, no aged-out items, no chronic patterns. The single LAND-NOW item is precisely specified and lands in Phase 7.
