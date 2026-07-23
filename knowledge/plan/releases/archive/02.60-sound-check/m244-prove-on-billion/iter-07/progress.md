**Type:** tik — gate (b) evidence + voice cells → presence-only (#2), under TOK-01. Run 4, tik 2.

# iter-07 — progress

## Ground truth (the gate-(b) content-stories sweep vs billion — foreground, 6.0m)
`run-content-stories.sh 1 --host billion.taildc510.ts.net` → **LANDED 46/49**, 3 failing pairs, ALL player-vantage
(every manager pair PASSED):
1. **hire-voice-fail · player** — "result too short (127 readable chars) — likely empty"
2. **asmt-voice-pass-en · player** — "result too short (128 readable chars) — likely empty"
3. **intv-voice-pass · player** — "no interview acknowledgement text" (the priority-#3 ack residual — SEPARATE, iter-08)

The evidence CONFIRMED the decided disposition: (1)+(2) are the 2 Bunny-absent voice cells whose PLAYER result IS
the recorded call playback → renders empty on a demo box with no Bunny keys (DEF-M240-01). Their manager views land.

## What landed (the PLAYER-presence-only disposition — rext, committed + green + pushed 1968fd2, tag → 6f52f0c)
There was NO per-session presence-only mechanism (presence-only was per-PRODUCT, ai-labs). Built the per-session
mirror, player-half only:
- **contentsession** (fixture): new `player_result_unavailable` + `player_unavailable_reason` (reason REQUIRED,
  validated) — set on hire-voice-fail + asmt-voice-pass-en with the disclosed reason.
- **content_manifest projection**: a `player_result_unavailable` session projects `player_presence_only` + the
  reason, NO player path, manager view intact — a legitimate disposition, NOT a fail-closed drop
  (`ValidateContentManifest` unaffected). 2 existing tests updated for the new disposition.
- **buildPairs** (content-pairs.ts): a `player_presence_only` session counts as presence-only (not a landable
  pair, not a drop); the manager pair still forms. + a focused unit test.
- **content-denominator.json**: 49 → 47 (+ arithmetic: 46 → 44 simulation player pairs, managers stay; + a recorded
  M244-iter-07 reason). The route-contract test (canonical manifest yields `expected_pairs`) stays consistent.
- **cockpit**: partition player-presence-only rows by VERDICT (they carry a pass/fail + a manager CTA), not the
  ai-labs presence group; DISCLOSE the withheld as-player ("as-player result unavailable", never silent). + a test.
- both honesty-gated projection goldens regenerated.

## Tests: full stack-seeding Go module green (`go test ./...` exit 0), 21 content TS unit specs green
(incl. the new player-presence-only + the 47-pair route-contract assertion), demo-stack cockpit **159 pass / 6
standing** (the academy+overlay debt, unchanged). 0 platform edits.

## The live 47/47 is coupled to gate (h)
billion still serves the m243 49-pair manifest; the sweep pins EXPECTED_PAIRS from the local denominator (now 47),
so a re-run needs billion re-seeded at the m244 tag. The live 47/47 lands at the gate-(h) cold reset-to-seed
(re-pins billion to `sound-check-m244-content-sweep-robustness` → 6f52f0c). NOT counted as discharged yet.

## Close — 2026-07-22

**Outcome:** gate (b) advanced — ground-truthed the 3 residuals (2 voice player + 1 intv-ack); the 2 Bunny-absent
voice cells dispositioned to PLAYER-presence-only (denominator 49 → 47, honest landable set), full-stack (Go + TS +
cockpit) + tested + pushed. Metric stays **3/8** (gate (b) counted only at the live 47/47 on m244-seeded billion).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (milestone gate a–h not fully met; 3/8; gate (b) sub-progress: denominator corrected + 2 residuals dispositioned, 1 left = intv-ack)
**Phase 5 grading:** (1) gate-met: n (3/8) — (2) triggered-tok: n (iter-05/06/07 all made measurable progress) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 2 of run 4) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (evidence-first: sweep before disposition), D2 (per-session player-presence-only mechanism, not remove-cell), D3 (cockpit partition-by-verdict + disclosure) — iter-07/decisions.md
**Side-deliverables:** none (all in the #2 planned scope; 0 platform edits).
**Routes carried forward:**
- **intv-voice-pass/player ack-page residual (#3)** → iter-08: "no interview acknowledgement text" on the PASSED
  interview player page (`/sim/<slug>/result/<sessionId>`, no reportId) — a sweep-assertion / render-shape mismatch,
  distinct from the extraction REPORT (gate g, done). Named handler: iter-08.
- **live 47/47** (Fate-2, covered) → gate-(h) cold reset-to-seed (re-seeds billion at the m244 tag) + the gate-(b)
  sweep re-drive.
**Lessons:** (1) evidence-first — running the sweep BEFORE dispositioning confirmed exactly which pairs fail + why
(all 3 were player-vantage; managers landed), preventing a blind denominator edit. (2) a per-session presence-only
disposition (player-half only, manager kept) needs a real field + fail-closed-safe plumbing across Go projection +
TS buildPairs + the cockpit partition — an empty player path alone is a DROP (fail-closed), not presence-only.
