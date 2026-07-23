**Type:** tok (triggered) — 5 consecutive no-metric-delta tiks since iter-06 (the literal primary-metric floor). Run 5, tik-count unaffected (toks don't count toward the cap).

# iter-12 — progress (triggered tok)

## Trigger
Primary metric (gate parts discharged) last moved iter-06 (2→3); flat at **3/8** across iters 07-08-09-10-11 = five consecutive zero-delta tiks. Re-survey confirms 3/8 (not stale — the metric has not moved), so the trigger stands and a tok is authored.

## Consolidation (see overview.md for the full framing)
TOK-01 is SOUND, not stalled — every one of the five tiks landed real, load-bearing work (gate-b sub-residuals resolved in tooling, the gate-(h) cold reset-to-seed DONE green at m244, the load-bearing cross-check fix, gate-b fully root-caused, gates a/e/g re-verified live). The coarse binary-per-gate metric reads flat because iters 07–11 all worked the single hardest bucket (gate b), which only counts at 47/47 and is at 45/47. Strategy class: **more-granular** — drill from "discharge parts" to the specific gate-b fix + a cheap live sweep of the remaining gates on the now-green seed.

## TOK-02 authored
Recorded in the milestone-root `decisions.md`. Sets the run-6 direction: iter-13 lands gate (b) 47/47 (scope the interview container FETCH demopatch to isManagerScope, mirror iter-08; re-bake + re-sweep), then iters 14–15 sweep gates (c)/(f) live, then (h)/(d)/DEF-M239-01.

## Close — 2026-07-22

**Outcome:** triggered tok — TOK-01 consolidated + sharpened (SOUND, not stalled; the 3/8 flat is a coarse-bucket artifact of 5 tiks working the gate-b cluster). TOK-02 written with the precise run-6 direction (gate b fix first, then the cheap live gates). No metric delta (toks don't move the gate). Metric stays **3/8**.
**Type:** tok (triggered)
**Status:** closed-no-lift
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: **y** (this iter is a triggered tok) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: **exit-2 (tok-fired)**
**Decisions:** TOK-02 (milestone-root decisions.md).
**Side-deliverables:** none.
**Routes carried forward:** the run-6 tik sequence (iter-13 gate b fix → iter-14 gate c → iter-15 gate f → then h/d/DEF-M239-01), all named in TOK-02's Next-tik direction.
**Lessons:** a binary-per-gate primary metric under-reports multi-tik work on a single hard bucket — 5 tiks of genuine gate-b progress (presence-only, ack demopatch, cold re-seed, cross-check fix, root-cause) read as a flat 3/8. The triggered tok is the right place to make that visible and re-affirm the strategy rather than let the flat metric masquerade as a stall.
