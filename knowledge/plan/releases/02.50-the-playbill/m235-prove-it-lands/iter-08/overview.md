---
iter: iter-08
milestone: M235
type: tik
iteration_type: tik
iter_shape: cleanup
status: planned
created: 2026-07-20
active_strategy: TOK-01 (two-track) — Track A step 4 (M230 carry-forward) + the M236 Fate-3 handoff
---

# iter-08 — the M236 Fate-3 handoff + record USER-BLOCKER-M235-02's resolution + corpus doc sync

**Type:** tik (cleanup-shape — process-quality: plan routing + doc sync, no production seeder code).
**Active strategy:** TOK-01 Track A step 4 (M230 carry-forward) + the run-3 ruling's user-authorized Fate-3.

## Step 0 — re-survey
iters 05–07 landed all 3 offline-buildable non-sim sections (3/3). What remains is exclusively LIVE-coupled:
the (session×action) browser proof, the new seat-login coverage/Playthrough plumbing, and the M230
carry-forward live items — ALL routed to M236 per the ruling. No offline production surface remains under
TOK-01. This iter is the routing + recording work (a cleanup-iter: 0 production seeder code).

## Cluster / target identified
The Fate-3 handoff to M236 (user-authorized) + the milestone-level resolution record + the corpus doc sync
that reflects "the non-sim sections are BUILT (offline); the live proof is M236".

## Phase plan (cleanup-iter — planned multi-step)
1. EDIT `m236-prove-on-billion/overview.md` `In:` list — add: (a) AUTHOR the new content-stories seat-login
   coverage/Playthrough plumbing (Playthroughs-style seat-login + exact dynamic `/sim/<slug>/result/<sessionId>`
   URL resolution + the shared result page-object) + CALIBRATE against a live seeded render; (b) prove every
   non-sim (session×action) lands (skill-path player+manager, academy player, ai-labs presence) with the
   per-section M236 live-calibration checklists; (c) the M230 carry-forward live items (ANT_ACADEMY coverage
   descriptor + next-web clone re-anchor + getPublicCatalogView 2nd manifest).
2. RECORD USER-BLOCKER-M235-02's RESOLUTION + the M236 Fate-3 handoff in the M235 milestone `decisions.md`.
3. Corpus doc sync: `content-stories-spec.md §6` (non-sim sections now built) + `content-stories-routes.md` §7
   (per-product disposition: skill-path/ai-labs/academy sections built offline, live proof M236).

## Escalation / close
After this iter the offline-buildable scope is EXHAUSTED (all 3 sections built + unit-proven; the live proof
legitimately routes to M236 per the ruling). EXIT_REASON: protocol-stop (offline clusters exhausted; live
proof routed to M236). The milestone PRAGMATIC-CLOSES per the user's "then close" mandate (the actual
close-milestone merge is a separate step the orchestrator/user drives).
