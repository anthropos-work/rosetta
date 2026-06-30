# iter-04 — decisions (local)

**INCOMPLETE-EXIT-2026-06-30: user-blocker — the demo-1 rext CONSUMPTION clone is in an
untrustworthy partially-hand-modified state that blocks the clean `fit-up-m50` checkout the
perf-wall demo re-up requires; every resolution path is in the FORBIDDEN-OPS list.**

- **What got done (iter-04 Phase A–B + start of C):**
  - Phase A: inherited the iter-03 GATED `(failing=6, escapes=0)` as the pre-iter metric.
  - Phase B (triage): all 6 failing sections = the M46 base-Workforce org-scale PERF-WALL
    (skeleton false-fails, data-in-DB), routed to the demo-UP fix surface (the coverage-protocol
    "Org-scale grid perf wall" row). The fix: re-pin demo-1's consumed rext tag to one carrying
    the perf demo-patches (`next-web-members-pagination` + `app-targetrole-authz-skip` + post-seed
    FK indexes) — all three are wired into `fit-up-m50`'s `up-injected.sh` (verified by
    `git show fit-up-m50:demo-stack/up-injected.sh`).
  - Phase C start: propagated the local-only `fit-up-m50` tag (`bbd8189`) from the AUTHORING copy
    into the consumption clone (`git fetch <authoring> refs/tags/fit-up-m50` — origin has neither
    m49 nor m50, the rext tags live only locally). Re-pinned the SoT `.agentspace/rext.tag` →
    `fit-up-m50` (gitignored; does not dirty rosetta git).

- **What blocked progress:** `git checkout fit-up-m50` in `stack-demo/rosetta-extensions` ABORTED.
  The consumption clone is NOT a pristine tag checkout — it carries leftover hand-modifications
  (almost certainly from the same concurrency incident that left iter-03 uncommitted):
  - `M demo-stack/up-injected.sh` — modified, and DIFFERS FROM BOTH `fit-up-m49` (its nominal HEAD)
    AND `fit-up-m50` (a +28-line partial edit, matching neither tag — so it cannot be trusted to
    wire the perf-patches correctly).
  - `M demo-stack/tests/test_demopatch.py` — modified, byte-IDENTICAL to `fit-up-m50`.
  - `?? demo-stack/patches/next-web-public-website-url/next-web-public-website-url.yaml` — untracked,
    byte-IDENTICAL to `fit-up-m50` — this UNTRACKED file is what aborts the checkout ("would be
    overwritten").
  Someone partially hand-applied the M50 demopatch set into the m49 consumption clone.

- **Why this is a user-blocker (Phase 5 §4 / TOP-OF-PROMPT BAN), not a routine route-forward:**
  the ONLY ways to unblock the `fit-up-m50` checkout are (a) remove/move the untracked
  `next-web-public-website-url.yaml` (= `git clean` / `rm`), and (b) revert the modified
  `up-injected.sh` + `test_demopatch.py` (= `git checkout --` / `git stash`). All four are in the
  build-iter FORBIDDEN-OPS list. The user + orchestrator are the only allowed deciders on this
  dirty consumption-clone state.

- **Recommended resolution (for the user/orchestrator):** in `stack-demo/rosetta-extensions`,
  decide whether the hand-modified `up-injected.sh` carries anything worth keeping (it differs from
  m50 — likely stale/partial). Most likely correct: discard the consumption-clone local mods +
  untracked file, then `git checkout fit-up-m50` cleanly, so demo-1 consumes the canonical m50
  tooling (perf-patches wired). Then resume iter-04 Phase C: `/demo-down 1` + `/demo-up 1` (rebuilds
  injected next-web + app with the perf-patches baked + applies post-seed FK indexes), re-seed the
  AI-readiness config+funnel from the authoring copy (as iter-03 — the iter-03 seeders are committed
  to authoring `main` @ 45a20c0 but NOT yet tagged, so a m50-consumed demo lacks them until the
  M51 tag), re-export roster+cockpit, then the GATED manager-vantage sweep.

- **iter-04 is NOT closed** (no fix landed, mid-Phase-C). No close-section, no `iter(M51/04):`
  commit. The untracked `iter-04/` dir (this file + overview.md + empty progress.md) is left
  uncommitted in the rosetta tree by design (Phase 4 Step 0 budget/blocker-interrupted-iter rule).
