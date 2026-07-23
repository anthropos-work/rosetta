# iter-25 — progress

**Type:** tik (run 9, tik 1) — gate (c) final push: recover demo-1 + root-cause & fix the 4 ai-readiness Playthroughs.

## What happened
1. **Recovered demo-1 on billion** (was DOWN, web 502). First attempt failed on `0.0.0.0:18082` bind (stale
   `tailscale serve` held the tailnet-IP port — iter-24's tangle). Fixed with a manual `tailscale serve reset`
   BEFORE the bring-up (up-injected's own pre-reset runs too late, after compose-up) → green (D4).
2. **Root-caused the 4 ai-readiness failures** — LIVE, refuting iter-23's `live_snapshots=0` correlation. The
   dashboard live path doesn't read that table; the pt-world Org C seed is complete (active+closed cycle, 105
   progress, 39 frozen snapshots, 1 interview report). The TELL was in the **billion backend logs**:
   `ERROR: column ai_readiness_cycles.launched_by does not exist` → cycles endpoint 500s → zero-state (D1).
3. **Pinned the true cause: an up-injected BUILD BUG** — the backend `:injected` build-scratch checked out the
   HIGHEST fetched v-tag (`v1.351.0`, which has `launched_by`) instead of the source's PINNED checkout (`v1.341.0`,
   which the schema was migrated from) → binary/schema version skew (D2).
4. **Shipped the durable rext fix** (`up-injected.sh` build-scratch + M217 preflight → source pinned ref;
   +3 regression tests; rext `c755370`, tag moved+pushed, main pushed, billion re-pinned) + rebuilt: app binary
   `launched_by`=0, cycles endpoint 200, autoverify green (D3).
5. **Re-ran the 4 specs**: **member-done PASSES** (fix proven). 3 sub-failures remain (byTeam / interview
   panel 24 chars / member deadline) — distinct render/wiring gaps at v1.341.0, routed to iter-26 (D5).

## Close — 2026-07-23

**Outcome:** landed the durable launched_by version-skew fix (rext c755370); gate (c) 12/16 → **13/16** (member-done now passes + dashboard core renders); 3 ai-readiness sub-renders characterized + routed.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (gate c 13/16; metric stays 7/8 — coarse binary-per-gate, TOK-02/03 pre-registered)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 1/5) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (true root cause: stale-version binary, not live_snapshots), D2 (the up-injected highest-tag build bug + why cold-reset would still hit it), D3 (the durable pinned-ref rext fix), D4 (tailscale serve reset recovery), D5 (fix worked; 3 sub-failures → 13/16, routed)
**Side-deliverables:** none (the rext fix IS the planned scope)
**Routes carried forward:** the 3 ai-readiness sub-render failures → iter-26 (handler FIND-M244-aireadiness-subrenders): (a) manager byTeam "AI Readiness by Tag" per-team breakdown; (b) manager how-we-measure interview-findings panel (report seeded 8396 chars but renders 24); (c) member-progress cycle deadline. Investigate seed team-tags / interview sim_id match / deadline render vs v1.341.0; 0 platform edits.
**Lessons:** always read the LIVE service logs before trusting a routed-forward correlation (iter-23's live_snapshots=0 was a red herring; the backend log named the real cause in one line). A demo's binary can silently diverge from its pinned clone + migrated schema when the build picks the highest fetched tag — the demo-provisioning "not all built as expected" class this release exists to catch.
