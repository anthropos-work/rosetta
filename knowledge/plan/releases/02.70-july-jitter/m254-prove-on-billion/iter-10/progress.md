**Type:** tik — TOK-01 cluster 4 (mutating tail): (e) studio builder Playthrough + (h) the Playthroughs green.

# M254 · iter-10 — progress

## Close — 2026-07-25

**Outcome:** The full Playthrough suite is GREEN on billion (cold reset-to-seed, tailnet peer): **18/18
Playthroughs passing (100 %), 0 failing, DONE_rc=0**. Gate **(e)** MET (both studio builders re-tuned to the
live studio-desk v0.152.1 UX + green) and gate **(h)-Playthroughs** MET (all 16 coordinator-listed + the 2
studio + the specs). Two live fixes shipped (rext `july-jitter-m254-studio-pt-retune` @ 4f1409e, on origin):
the studio-builder re-tune (e) and the `/home` networkidle anti-deadlock (h). 0 platform edits.

**Type:** tik
**Status:** closed-fixed
**Gate:** MET — (e) studio builders green + (h) Playthroughs 18/18 green. With the prior-iter MET parts
(a,b,c,d,f-session-carry,h-latency) + the 3 recorded dispositions, the milestone is at **effective gate-met**.
**Phase 5 grading:** (1) gate-met: **y** — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: **exit-1 (gate-met)**.
**Decisions:** D1 (16/16 coordinator-listed green), D2 (gate-e studio re-tune to v0.152.1 — FIXED), D3 (gate-h
skillpath-legacy networkidle → domcontentloaded on 4 /home logins — FIXED), D4 (definitive 18/18 + rung-zero),
D5 (3 dispositions recorded). See iter-10/decisions.md.
**Side-deliverables:** the /home networkidle anti-deadlock hardening on the 3 currently-passing /home specs
(aireadiness-member-done/-progress + aisim-chat-launch) — same root cause as the skillpath-legacy fix, folded
into D3 (not a separate line; it made gate (h) robustly green vs jitter-dependent + cut the suite 13m→3.8m).
**Routes carried forward (→ close-milestone deferral audit):**
- **(f)-FCP-p95** = ACCEPTED-environmental disposition (recorded milestone decisions.md + carry-forward.md).
- **(c)-academy-durability** = Fate-3 academy-durable follow-up (recorded).
- **(g)-testhealth** = carry-forward `FIX-M254-g-testhealth` (6 host-sensitive tests; recorded).
- **(b)-voice manager_presence_only** flag + content-denominator 47→45 re-seed (prior-iter follow-up; carried).
- **studio-desk billion re-pin to 4f1409e** on the NEXT full cold re-prove (test-only fix → not needed for
  the current demo build; noted so a future re-prove picks it up).
**Lessons:** (1) A "prove-on-billion" live run is where authored-blind locators finally meet the real DOM —
the M252 studio Playthrough was never live-tuned; the studio-desk redesign (v0.152.1, predating the release)
only surfaced here. Live-tuning IS the milestone's job. (2) The redesigned advanced builder gives a *stronger*
gate-e proof: "Design it with AI" drafts a real scenario (characters + tasks) — the designer rendering that
draft proves generation, not just a button. (3) The `/home` polling surface + a `networkidle` login default is
a latent flaky-timeout across 4 Playthroughs; the protocol's "never networkidle" doctrine is load-bearing, not
stylistic — applying `domcontentloaded` fixed the flake AND cut the suite wall-clock ~3.4×.
