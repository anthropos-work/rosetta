---
iteration_type: tik
status: in-progress
---

# iter-03 — root-cause + fix the believability leaks (tik)

**Active strategy:** TOK-01 `reprove-corrected-hiring-on-billion`.
**Cluster / target:** the 3 iter-02 findings that block the gate — F2 (hiring-only training leak), F3 (1-sim/candidate
2nd-session leak), F1 (succession/interview_extraction_results FK seed error). Root-cause each in the rext seed
tooling, fix (0 platform edits), then a default cold re-bring-up + warm re-measure on billion.
**Hypothesis:** F1/F2/F3 share a root — Workforce-Intelligence dashboard seeders that write jobsimulation/mirror rows
but were left unguarded by M227 fix#1. Adding the `skipGenericActivityForHiringOrg` guard fixes all three.
**Expected lift:** 5/7 → 7/7 (C2 1-sim/candidate + hiring-only, C4 hiring-only, clean seed) on the fixed tag.
**Phase plan:** root-cause (read the session/interview writers) → fix + RED-proven tests → tag `casting-call-m228-*` →
default cold re-bring-up → warm re-measure C2/C3/C4/C5.
**Escalation conditions:** a fix that needs a platform edit → escalate (none expected — pure seed tooling).
**Acceptable close-no-lift:** n/a — a fix + re-measure iter.

## Root cause (confirmed)
All 3 findings are ONE class: **M227 fix#1 left two WI dashboard seeders unguarded, wrongly assuming they "write no
jobsimulation/mirror rows":**
- **FeedbackSeeder (F2/F3):** since v1.10 M42m it writes `public.local_jobsimulation_sessions` MIRROR rows on GENERIC
  `refs.sims` sims → leaked training sims + a 2nd session per hiring candidate into the recruiter's list.
- **SuccessionSeeder (F1):** writes `jobsimulation.interview_extraction_results` rows FKing each member's first
  population session (written by JobsimSessionsSeeder). Once fix#1 made JobsimSessionsSeeder skip the hiring org, that
  session no longer exists → FK violates (SQLSTATE 23503) → the whole seed reports "failed".

## Fix
Both now consult `skipGenericActivityForHiringOrg` (skip the hiring org). The fix#1 regression test
`TestGenericActivitySeeders_SkipHiringOrg` gains the feedback + succession cases (RED-proven: 42 + 20 leaked rows
without the guard). `hiring_scope.go`'s rule comment corrected. rext tag `casting-call-m228-hiring-scope-fix` (1d97861).
