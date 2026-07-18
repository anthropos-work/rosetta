# iter-04 — progress (tik, under TOK-01) — THE FIRST GATE READING

**Reconciled at the TOK-02 boundary** (executed + committed during a driven build-iter leg; record completed here).

## What happened
1. **Baseline render measured:** `min(rows-per-sim) = 0` on the recruiter's comparison surface.
2. **Attribution = SEED-GAP** (the M219 trap, caught by measure-first): the recruiter hero's role
   `"Technical Recruiter"` did not resolve to a real public `job_role` (no `job_role_skills`), so the users-seeder
   fail-fast killed the **entire** hiring seed (Meridian = 0 members; blast-radius to the other orgs).
3. **Fix:** recruiter role `"Technical Recruiter"` → `"Talent Acquisition Specialist"` (a real, resolvable public
   role). rext tag `casting-call-m224-iter04`; `.agentspace/rext.tag` bumped; consumption copy synced.
4. **Cold reset-to-seed:** all 27 seeders GREEN. DB ground truth POST-FIX: **Meridian Talent = 50 members, 259
   sessions; each of the 5 shared hiring sims = 43 comparable candidates; scores 27–100, non-degenerate.**
   **→ the DATA side of the gate is MET.**

## Outcome
**closed-fixed.** Render still 0 — but the residual is NOT seed (data confirmed present). Handed to iter-05 for the
render/reachability diagnosis. No fix-before-attribution violated.
