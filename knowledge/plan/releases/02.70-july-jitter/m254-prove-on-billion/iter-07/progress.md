**Type:** tik — TOK-01 cluster 2 (g: host-sensitive demo-stack test-health).

# M254 · iter-07 — progress

## Close — 2026-07-25

**Outcome:** Captured the exact live failures on billion (10 fail + 1 error / 159 across the host-sensitive
files; ~7 unique + 1 error after dedup of the `TestAntAcademyPreBindReap`←`Launcher` inheritance). Root causes
are **diverse**, not one common fix. **FIXED** the clearest high-leverage one — the nvm/node host-robustness of
`test_missing_node_documents` (×2) — verified live on billion (rext `dfdd9bc`, tag
`july-jitter-m254-academy-nonode-hostrobust` on origin). **Characterized + routed** the remaining 6 as a
dedicated test-health batch with precise root causes + named handlers.

**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (overall) — (g) 2/8-class fixed (nvm host-robustness); 6 routed as a test-health batch → coordinator fate. The core demo-proof gates (e/h/c-academy) still pending.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (the routed test-health items are chronic host-sensitive tests historically fated at close; not blockers that change what code lands) — (5) cap-reached: n (tik 3 of the session) — (6) protocol-stop: n — Outcome: continue (→ iter-08).
**Decisions:** D1 live-failure capture; D2 nvm host-robustness fix (verified); D3 root-cause + route the remaining 6 (`FIX-M254-g-testhealth`); D4 (g) disposition (1 fixed / 6 routed → coordinator fate).
**Side-deliverables:** none (the nvm fix is the planned-scope deliverable).
**Routes carried forward (`FIX-M254-g-testhealth` batch):**
- `FIX-M254-g-academy-launch-isolation` — intra-run `:23077` listener leakage + M245 reconcile-message drift (test_second_launch / test_stop_kills / test_a_stale_academy_reaped).
- `FIX-M254-g-devorigins-sha-repin` — next.config.js sha drift (test_apply_revert_round_trip).
- `FIX-M254-g-mutation-meta` — 2 mutation meta-tests' host-nondeterminism (test_mutant_no_term_trap, test_MUTATION_without_the_reap_block).
- `FIX-M254-g-overlay-127` — the write_env_local snippet exit-127 (ERROR test_overlay_has_minted_pk).
- Still pending (core demo-proof): **(c)-academy** durability re-heal; **(e)** builder Playthrough; **(h)-Playthroughs**; **(f)-FCP-p95** coordinator disposition.
**Lessons:** (1) A test harness that builds a "tool absent" PATH from fixed system dirs (`/usr/bin:/bin`) is only host-robust if that tool never lives there — a dev/VM host that installed it system-wide silently defeats the stub. Build a curated allow/deny bindir + a clean HOME instead. (2) A production change to launch semantics (M245 reconcile+render_ok) that isn't mirrored in the fixture (a stub that "runs" but doesn't "render") turns green tests red without a real regression — the stub must model the new contract. (3) Detached (setsid) test daemons on a SHARED port leak across an intra-run suite; tearDown must reap the PORT, not just the recorded pid.
