# iter-07 — decisions (local)

- **D1 — evidence-first: run the sweep BEFORE dispositioning.** The resume plan named the 2 cells to drop, but
  there was no per-session presence-only mechanism and no ground truth on WHICH pairs fail + why. Running
  `run-content-stories.sh --host billion` first (foreground, 6.0m) confirmed the 3 residuals precisely — all 3 are
  PLAYER-vantage (managers all landed), and the 2 voice ones fail with "result too short ~127 chars, likely empty".
  This prevented a blind denominator edit and confirmed the disposition is PLAYER-only (keep the manager pair).

- **D2 — a per-session PLAYER-presence-only field, NOT removing the cell.** Presence-only was per-PRODUCT (ai-labs,
  playerLink=false). Removing the 2 cells would delete their seeded activity + their (landing) manager views. The
  honest disposition is per-session: `player_result_unavailable` (+ a REQUIRED disclosed reason) → the projection
  withholds the player CTA (`player_presence_only`) but keeps the manager view. Critically, buildPairs treats a
  non-ai-labs empty player path as a fail-closed DROP (which would FAIL the sweep) — so a real `player_presence_only`
  signal is needed on BOTH the Go projection and the TS buildPairs, or the disposition would read as a failure.

- **D3 — cockpit: partition by VERDICT + disclose the withheld player.** A player-presence-only row has a pass/fail
  verdict + a manager CTA, so it belongs in its verdict column, not lumped with ai-labs "Activity & spend only"
  (the old partition keyed on player_result_path alone). And the withheld as-player is DISCLOSED (the reason travels
  in the manifest and renders as a muted note) — never a silent omission, per the cockpit-guard philosophy.

- **live 47/47 is coupled to gate (h).** billion serves the m243 manifest; the sweep cross-checks the served
  manifest against the local denominator (now 47), so the live 47/47 needs billion re-seeded at the m244 tag. Booked
  to the gate-(h) cold reset-to-seed. gate (b) is NOT counted discharged until that live pass + the intv-ack fix (#3).

- **0 platform edits.** All rext tooling (contentsession + projection + buildPairs + cockpit) + data (denominator +
  fixture + goldens).
