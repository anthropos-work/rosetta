# iter-27 — progress

**Type:** tik (run 10, tik 1) — land the 3 remaining ai-readiness sub-renders → gate (c) 16/16 = GATE MET.

## What happened
Re-surveyed the 3 remaining ai-readiness Playthrough failures against billion's LIVE DB + the v1.341.0 read
paths + the actual next-web UI, and **falsified iter-26 D1's seed-gap characterization on all three** — every
one is a HARNESS locator mismatch vs the correct, fully-seeded, correctly-rendering v1.341.0 UI (Org C has 13
tags/42 memtags/12 teams; the interview report is rich at 8396 chars keyed by org_id only; the active-cycle
deadline renders "Due Aug 22 · N days left"). Fixed the 3 locators + a 4th (a tailnet pick-timing flake in the
assign-WRITE spec that surfaced only once the full 16-suite ran), all in the Playthroughs e2e lib — 0 seed
changes, 0 platform edits, no billion re-pin/re-seed (the harness runs from the LOCAL authoring copy). Proven
**16/16 Playthroughs GREEN in one clean full run on billion** (96 passed, 5.7m, 100%).

- (a) `byTeam()` `/AI Readiness by Tag/i` → `/AI Readiness by Team/i` (the i18n title; "byTag" is the internal key).
- (b) added `interviewBreakdownPanel()` scoping the >900-char density check to the breakdown CARD, not the
  24-char heading span (findings render as its siblings inside the `openStep===3` card).
- (c) `dueDate()` accepts the year-less "Aug 22" short-date the UI renders, keeping the `due|deadline` anchor.
- (+) `pickFirstSkillPath()` gates the keyboard commit on a real option rendering + retries ×3 until the
  submit enables (rc-virtual-list ArrowDown fired before options painted flaked once over the tailnet).

Combined with the 24 stack-verify discrete specs (green iter-16/18), gate (c) = **40 live-browser specs green**
→ the binary gate-parts metric flips **7/8 → 8/8**.

## Close — 2026-07-23

**Outcome:** all 3 remaining ai-readiness sub-renders + a 4th assign-WRITE flake root-caused to HARNESS
locator/timing issues (data fully seeded + rendering); 4 harness fixes; **16/16 Playthroughs GREEN in one
clean full billion run (96 passed)** → gate (c) ticks → gate parts **8/8 GATE MET**.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET (all 8 exit-gate parts a–h discharged green live on billion; 0 platform edits)
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 1/5) — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (revises iter-26 D1: 3 sub-renders are harness locator mismatches, not seed gaps, each
empirically root-caused), D2 (the 3+1 harness fixes + disposition), D3 (protocol lesson: root-cause the
locator against the live UI+i18n before concluding a seed gap; the empirical order DB→read-path→UI→locator)
**Side-deliverables:** the assign-WRITE tailnet pick-timing flake hardening (`pickFirstSkillPath` option-gated
retry) — surfaced mid-iter when the full 16-suite first ran (15/16, assignment-assign flaked at line 83);
proven a flake (passed on immediate re-run 6.4s), hardened, 0 platform edits. It is part of the gate-(c)
16/16 landing, folded into the same iter (the gate needs all 16 green together).
**Routes carried forward:** DEF-M239-01 (ENOSPC loud-build-fail) remains the sole inherited open carry —
routed to close-milestone / a follow-on; non-blocking, does not gate.
**Lessons:** a "prove on billion" run is the FIRST time these Playthrough locators meet the real target — a
plausible-but-wrong string/scope/format (Tag vs Team, heading vs card, year-bearing vs year-less date) stays
latent until then. When a sub-render "fails" on a demo whose DATA is provably present, walk the rungs
DB-has-data → read-path-produces-it → UI-renders-it → locator-matches-it, and fix at the rung that breaks —
here it was the last rung for all four. (Recorded as D3 for the protocol.)
