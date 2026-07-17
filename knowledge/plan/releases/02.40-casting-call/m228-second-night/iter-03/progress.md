**Type:** tik (under TOK-01 `reprove-corrected-hiring-on-billion`). Root-cause + fix the 3 iter-02 findings (F1/F2/F3)
in the rext seed tooling, then a default cold re-bring-up + warm re-measure on billion. Protocol: verification.md +
coverage-protocol.md + latency-budget.md.

# iter-03 — work log

1. **Root-caused all 3 findings to ONE class** (see `overview.md § Root cause`): M227 fix#1 left two Workforce-
   Intelligence dashboard seeders unguarded, wrongly assuming they "write no jobsimulation/mirror rows":
   - **FeedbackSeeder (F2/F3)** — writes `public.local_jobsimulation_sessions` MIRROR rows on GENERIC `refs.sims` sims
     (since v1.10 M42m) → leaked training sims + a 2nd session per hiring candidate into the recruiter's list.
   - **SuccessionSeeder (F1)** — writes `jobsimulation.interview_extraction_results` FKing JobsimSessionsSeeder's
     (now-skipped-for-hiring) sessions → FK violation (SQLSTATE 23503) → the whole seed reports "failed".
2. **Fixed:** both now consult `skipGenericActivityForHiringOrg` (2 guard lines). The fix#1 regression test
   `TestGenericActivitySeeders_SkipHiringOrg` gains the feedback + succession cases — **RED-proven** (42 + 20 leaked
   rows without the guard); GREEN with it. `hiring_scope.go`'s rule comment corrected. `go vet` clean; 13 stack-seeding
   packages GREEN.
3. **Committed + tagged:** rext authoring `1d97861`, tag **`casting-call-m228-hiring-scope-fix`**, pushed origin
   (main + tag). rext.tag SoT updated. 0 runtime change beyond the 2 guard lines (harden commit 78a3cb2 confirmed test-only).
4. **Default cold re-bring-up at the fixed tag** — teardown --purge (RC=0) + rext cutover → `casting-call-m228-hiring-scope-fix`
   (1d97861) + `up-injected.sh 1` (NO FLAGS). **IN PROGRESS at wind-down** (Directus-provision/seed phase).

## Re-measurement — PENDING (resume step)
The default cold re-bring-up at the fixed tag was launched; the warm re-measure of C2 (1-sim/candidate + hiring-only)
+ C3 + C4 + C5 + the seed-succeeded (F1) check is the resume step. **This iter is NOT yet closed** (no close section
until the re-measure lands). See the journal RESUME block.

## Expected result (the fix's prediction, to verify on resume)
- **F1:** the seed completes with NO "dev-setdress: seed failed" (succession skips hiring → no FK crash).
- **F2:** the recruiter's list is HIRING-only (no training sims; the 2 leaked training sims gone).
- **F3:** each candidate on exactly 1 sim (~8/position); the 17 extra sessions gone → ~40 candidate sessions (down from 62).
- C1/C3/C5/C6/C7 + fix#2 hold (unchanged by the seed-guard fix).
