# iter-06 — decisions (local)

- **D1 — the assertion lives at the `contentsession` layer, with a checked-in golden of the PLAN section-ids.**
  The seeded interview report DATA lives in the contentsession fixtures; the captured PLAN lives in the
  snapshot (`directus.simulations_extraction`, added to the capture surface iter-05). The alignment is a
  cross-tool invariant, so the assertion pairs the fixture's report keys with a checked-in golden of the
  plan's section-ids per scope (`fixture/interview-plan-sections.json`, grounded read-only from prod — public
  non-PII sim-def metadata). The golden is offline (no live prod at test time) and mirrors what the snapshot
  captures. The stack-snapshot golden (iter-05) already asserts the table IS captured; together they close the
  chain: the demo captures the plan AND the seeded data aligns to it → the report renders.

- **D2 — re-pin intv-voice-fail as the REAL fix, not an allowlisted exception.** The assertion correctly
  flagged intv-voice-fail's v1.3-era report (breadth/context_fit/frequency orphans) against the sim's v1.4
  plan. The honest fix is to re-pin to a session whose report was produced under the current plan, not to
  weaken the assertion. `05dae0f7` is the ideal target: same sole public interview sim, a genuine FAILED
  interview (completion_status=failed, score 0 — matching the exhibit's intent), v1.4-clean (11 in-plan
  manager keys, 0 orphans). content-capture `--only intv-voice-fail --dsn <read-only prod>` regenerated the
  fixture (scrubbed); the offline cleanliness (surviving-token) gate stayed green. Both honesty-gated
  projection goldens were regenerated in lockstep.

- **D3 — DB-level live alignment proof over a flaky authenticated deep-render click.** iter-04/05 found the
  authenticated "Explore Key Moments" deep-render drive flaky (timed out). Instead of another flaky drive, I
  proved alignment on the live demo directly: a query of the two seeded interview clones' `manager_report`
  keys vs the captured plan. It is non-flaky, exact, and STRONGER — it shows the aligned pass clone is
  0-orphans (renders aligned, with iter-05's shell render 0→520) AND that the assertion catches the
  m243-seeded fail clone's real live drift ({breadth,context_fit,frequency}), which the re-pin resolves on the
  next cold reset-to-seed. The visual deep-click is routed to the gate-(b)/(h) live sweeps (Fate-2, covered)
  which re-drive these surfaces and re-seed the re-pinned fixture.

- **0 platform edits.** All work is rext tooling (contentsession assertion + golden + fixture re-pin) + read-only
  prod capture. gate (g) discharged (3/8).
